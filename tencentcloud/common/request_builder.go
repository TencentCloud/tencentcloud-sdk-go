// Package common
package common

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// contentType enumerates the http Content-Type header values supported by the SDK.
type contentType int

const (
	contentTypeJSON contentType = iota
	contentTypeForm
	contentTypeOctetStream
)

func (c contentType) String() string {
	switch c {
	case contentTypeJSON:
		return "application/json"
	case contentTypeForm:
		return "application/x-www-form-urlencoded"
	case contentTypeOctetStream:
		return "application/octet-stream"
	default:
		return ""
	}
}

// signMethod enumerates the signing schemes supported by the SDK.
type signMethod int

const (
	signMethodV3 signMethod = iota
	signMethodHmacSHA256
	signMethodHmacSHA1
	signMethodSkip
)

// String returns the wire-format name expected by the server and the V1
// SignatureMethod query parameter.
func (s signMethod) String() string {
	switch s {
	case signMethodV3:
		return "TC3-HMAC-SHA256"
	case signMethodHmacSHA256:
		return "HmacSHA256"
	case signMethodHmacSHA1:
		return "HmacSHA1"
	case signMethodSkip:
		return "SKIP"
	default:
		return ""
	}
}

// requestBuilder holds all params for one cloud API call and builds a signed *http.Request.
type requestBuilder struct {
	// —— API identity ——
	service string
	version string
	action  string
	region  string

	// —— Transport ——
	// single source of truth for transport; resign only swaps url.Host
	url *url.URL
	// host signs the canonical Host header; may differ from url.Host
	// (e.g. behind a proxy where transport host != signing host)
	host   string
	method string // GET / POST

	// —— Payload ——
	// payload is one of map[string]interface{}, tchttp.Request, or []byte;
	// dispatched via type assertion at build time.
	payload     interface{}
	contentType contentType

	// —— Caller headers (excluded from signing) ——
	// http.Header to match *http.Request.Header, no conversion needed
	headers http.Header

	// —— Signing ——
	credential      CredentialIface
	signMethod      signMethod
	unsignedPayload bool
	language        string
	requestClient   string

	ctx context.Context

	// func to make timestamp testable and resign auto-refresh
	now func() time.Time

	nonce uint32
}

// build produces a fully-signed *http.Request.
//
// Dispatches on signMethod: V3 (TC3-HMAC-SHA256), V1 (HmacSHA1/HmacSHA256),
// or SKIP (no signature).
func (r *requestBuilder) build() (*http.Request, error) {
	switch r.signMethod {
	case signMethodV3:
		return r.buildV3()
	case signMethodHmacSHA256, signMethodHmacSHA1:
		return r.buildV1()
	case signMethodSkip:
		return r.buildSkip()
	default:
		return nil, fmt.Errorf("unsupported sign method: %v", r.signMethod)
	}
}

// ---------- V3 (TC3-HMAC-SHA256) ----------

func (r *requestBuilder) buildV3() (*http.Request, error) {
	t := r.now()
	timestamp := strconv.FormatInt(t.Unix(), 10)
	date := t.UTC().Format("2006-01-02")

	bodyBytes, canonicalQuery, err := r.v3BodyAndQuery()
	if err != nil {
		return nil, err
	}

	ctStr := r.contentType.String()
	host := r.host
	if host == "" {
		host = r.url.Host
	}

	// canonical headers and signed headers list — order matters for the signature.
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\n", ctStr, host)
	signedHeaders := "content-type;host"

	canonicalURI := r.url.EscapedPath()
	if canonicalURI == "" {
		canonicalURI = "/"
	}

	hashedPayload := sha256hex(string(bodyBytes))
	if r.unsignedPayload {
		hashedPayload = sha256hex("UNSIGNED-PAYLOAD")
	}

	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		r.method, canonicalURI, canonicalQuery,
		canonicalHeaders, signedHeaders, hashedPayload)

	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, r.service)
	stringToSign := fmt.Sprintf("TC3-HMAC-SHA256\n%s\n%s\n%s",
		timestamp, credentialScope, sha256hex(canonicalRequest))

	secId, secKey, token := r.credential.GetCredential()
	secretDate := hmacsha256(date, "TC3"+secKey)
	secretService := hmacsha256(r.service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(stringToSign, secretSigning)))

	authorization := fmt.Sprintf("TC3-HMAC-SHA256 Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		secId, credentialScope, signedHeaders, signature)

	// assemble final headers: signature headers first, caller headers next,
	// so caller cannot accidentally overwrite Authorization/Host/etc.
	hdr := http.Header{}
	hdr.Set("Content-Type", ctStr)
	hdr.Set("Host", host)
	hdr.Set("X-TC-Action", r.action)
	hdr.Set("X-TC-Version", r.version)
	hdr.Set("X-TC-Timestamp", timestamp)
	hdr.Set("X-TC-RequestClient", r.requestClient)
	if r.language != "" {
		hdr.Set("X-TC-Language", r.language)
	}
	if r.region != "" {
		hdr.Set("X-TC-Region", r.region)
	}
	if token != "" {
		hdr.Set("X-TC-Token", token)
	}
	if r.unsignedPayload {
		hdr.Set("X-TC-Content-SHA256", "UNSIGNED-PAYLOAD")
	}
	hdr.Set("Authorization", authorization)
	for k, vs := range r.headers {
		for _, v := range vs {
			hdr.Add(k, v)
		}
	}

	httpReq, err := http.NewRequest(r.method, r.url.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	if r.ctx != nil {
		httpReq = httpReq.WithContext(r.ctx)
	}
	// GET: canonicalQuery is built from params but not in r.url.RawQuery;
	// attach it so the outgoing request actually carries the query string.
	if canonicalQuery != "" {
		httpReq.URL.RawQuery = canonicalQuery
	}
	httpReq.Host = host
	httpReq.Header = hdr
	httpReq.ContentLength = int64(len(bodyBytes))
	return httpReq, nil
}

// v3BodyAndQuery resolves payload to (body, canonicalQuery) per http method.
//
// GET: payload treated as params, encoded into url.RawQuery; body empty.
// POST: payload serialized per contentType into body; query empty.
func (r *requestBuilder) v3BodyAndQuery() (body []byte, canonicalQuery string, err error) {
	if r.method == http.MethodGet {
		params, perr := r.paramsFromPayload()
		if perr != nil {
			return nil, "", perr
		}
		// V3 canonical query excludes the SDK-common params; they ride in headers instead.
		for _, k := range []string{"Action", "Version", "Region", "Timestamp", "Nonce", "RequestClient"} {
			delete(params, k)
		}
		canonicalQuery = tchttp.GetUrlQueriesEncoded(params)
		return nil, canonicalQuery, nil
	}

	// POST
	switch r.contentType {
	case contentTypeJSON:
		switch p := r.payload.(type) {
		case []byte:
			body = p
		case tchttp.Request:
			body, err = json.Marshal(p)
			if err != nil {
				return nil, "", err
			}
		case map[string]interface{}:
			body, err = json.Marshal(p)
			if err != nil {
				return nil, "", err
			}
		default:
			return nil, "", fmt.Errorf("unsupported payload type for json: %T", r.payload)
		}
	case contentTypeForm:
		params, perr := r.paramsFromPayload()
		if perr != nil {
			return nil, "", perr
		}
		body = []byte(tchttp.GetUrlQueriesEncoded(params))
	case contentTypeOctetStream:
		bb, ok := r.payload.([]byte)
		if !ok {
			return nil, "", fmt.Errorf("octet-stream payload must be []byte, got %T", r.payload)
		}
		body = bb
	default:
		return nil, "", fmt.Errorf("unsupported content type for V3 POST: %v", r.contentType)
	}
	return body, "", nil
}

// paramsFromPayload extracts a map[string]string from payload for query/form encoding.
//
// - map[string]interface{}: each value stringified via fmt.Sprintf.
// - tchttp.Request: ConstructParams populates request.GetParams() via reflection.
// - map[string]string: returned as-is.
// - []byte: rejected (must use a structured payload for params).
func (r *requestBuilder) paramsFromPayload() (map[string]string, error) {
	switch p := r.payload.(type) {
	case map[string]string:
		return copyStringMap(p), nil
	case map[string]interface{}:
		out := make(map[string]string, len(p))
		for k, v := range p {
			out[k] = fmt.Sprintf("%v", v)
		}
		return out, nil
	case tchttp.Request:
		if err := tchttp.ConstructParams(p); err != nil {
			return nil, err
		}
		return copyStringMap(p.GetParams()), nil
	default:
		return nil, fmt.Errorf("cannot derive params from payload type %T", r.payload)
	}
}

func copyStringMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// ---------- V1 (HmacSHA1 / HmacSHA256) ----------

func (r *requestBuilder) buildV1() (*http.Request, error) {
	params, err := r.paramsFromPayload()
	if err != nil {
		return nil, err
	}

	// V1 common params ride in the query/form, not headers.
	params["Action"] = r.action
	params["Version"] = r.version
	params["Region"] = r.region
	params["Timestamp"] = strconv.FormatInt(r.now().Unix(), 10)
	params["Nonce"] = strconv.FormatUint(uint64(r.nonce), 10)
	params["RequestClient"] = r.requestClient

	secId, secKey, token := r.credential.GetCredential()
	params["SecretId"] = secId
	params["SignatureMethod"] = r.signMethod.String()
	if token != "" {
		params["Token"] = token
	} else {
		delete(params, "Token")
	}
	delete(params, "Signature")

	host := r.host
	if host == "" {
		host = r.url.Host
	}

	// string-to-sign: method + host + path + "?" + sorted(k=v & ...)
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	buf.WriteString(r.method)
	buf.WriteString(host)
	buf.WriteString(r.url.EscapedPath())
	buf.WriteByte('?')
	for i, k := range keys {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(params[k])
	}
	stringToSign := buf.String()

	signature := Sign(stringToSign, secKey, r.signMethod.String())
	params["Signature"] = signature

	method := r.method
	u := *r.url
	var bodyReader io.Reader
	var bodyBytes []byte
	if method == http.MethodGet {
		q := url.Values{}
		for _, k := range keys {
			q.Add(k, params[k])
		}
		q.Add("Signature", signature)
		u.RawQuery = q.Encode()
	} else {
		// POST form: body is urlencoded params (including Signature).
		form := url.Values{}
		for _, k := range keys {
			form.Add(k, params[k])
		}
		form.Add("Signature", signature)
		bodyBytes = []byte(form.Encode())
		bodyReader = bytes.NewReader(bodyBytes)
	}

	httpReq, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}
	if r.ctx != nil {
		httpReq = httpReq.WithContext(r.ctx)
	}
	httpReq.Host = host
	httpReq.Header.Set("Content-Type", contentTypeForm.String())
	if method == http.MethodPost {
		httpReq.ContentLength = int64(len(bodyBytes))
	}
	for k, vs := range r.headers {
		for _, v := range vs {
			httpReq.Header.Add(k, v)
		}
	}
	return httpReq, nil
}

// ---------- SKIP (no signature) ----------

func (r *requestBuilder) buildSkip() (*http.Request, error) {
	bodyBytes, canonicalQuery, err := r.v3BodyAndQuery()
	if err != nil {
		return nil, err
	}

	u := *r.url
	if canonicalQuery != "" {
		u.RawQuery = canonicalQuery
	}

	method := r.method
	var bodyReader io.Reader
	if len(bodyBytes) > 0 {
		bodyReader = bytes.NewReader(bodyBytes)
	}
	httpReq, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}
	if r.ctx != nil {
		httpReq = httpReq.WithContext(r.ctx)
	}
	host := r.host
	if host == "" {
		host = r.url.Host
	}
	httpReq.Host = host
	httpReq.Header.Set("Content-Type", r.contentType.String())
	httpReq.Header.Set("Host", host)
	httpReq.Header.Set("X-TC-Action", r.action)
	httpReq.Header.Set("X-TC-Version", r.version)
	httpReq.Header.Set("X-TC-Timestamp", strconv.FormatInt(r.now().Unix(), 10))
	httpReq.Header.Set("X-TC-RequestClient", r.requestClient)
	if r.language != "" {
		httpReq.Header.Set("X-TC-Language", r.language)
	}
	if r.region != "" {
		httpReq.Header.Set("X-TC-Region", r.region)
	}
	if r.credential != nil {
		if token := r.credential.GetToken(); token != "" {
			httpReq.Header.Set("X-TC-Token", token)
		}
	}
	httpReq.Header.Set("Authorization", signMethodSkip.String())
	for k, vs := range r.headers {
		for _, v := range vs {
			httpReq.Header.Add(k, v)
		}
	}
	if len(bodyBytes) > 0 {
		httpReq.ContentLength = int64(len(bodyBytes))
	}
	return httpReq, nil
}

// ---------- Context plumbing ----------

type requestBuilderKeyType struct{}

var requestBuilderKey requestBuilderKeyType

// withRequestBuilder stashes r on ctx so downstream pipeline stages (endpoint
// failover) can rebuild a signed request without
// reverse-engineering the serialized *http.Request.
func withRequestBuilder(ctx context.Context, r *requestBuilder) context.Context {
	return context.WithValue(ctx, &requestBuilderKey, r)
}

// requestBuilderFromContext returns the requestBuilder stashed by withRequestBuilder, or nil.
func requestBuilderFromContext(ctx context.Context) *requestBuilder {
	r, _ := ctx.Value(&requestBuilderKey).(*requestBuilder)
	return r
}

// ---------- Sign method parsing ----------

// parseSignMethod maps the legacy string sign method names to the enum.
// Used while Client.signMethod is still a string; once Client migrates to
// the enum directly this can be removed.
func parseSignMethod(s string) signMethod {
	switch s {
	case "TC3-HMAC-SHA256":
		return signMethodV3
	case "HmacSHA256":
		return signMethodHmacSHA256
	case "HmacSHA1":
		return signMethodHmacSHA1
	case "SKIP":
		return signMethodSkip
	default:
		return signMethodV3
	}
}

// newRequestBuilderFromRequest builds a requestBuilder from a tchttp.Request
// and Client config, for stashing on the *http.Request's context so downstream
// pipeline stages (endpoint failover) can resign via resign.
//
// Fields mirror what client.go's sendWithSignatureV3/V1/Skip paths already
// compute inline; this is a bridge until those paths are rewritten to use
// requestBuilder.build() directly.
func newRequestBuilderFromRequest(c *Client, request tchttp.Request) *requestBuilder {
	u := &url.URL{
		Scheme: request.GetScheme(),
		Host:   request.GetDomain(),
		Path:   request.GetPath(),
	}
	if request.GetHttpMethod() == http.MethodGet {
		u.RawQuery = tchttp.GetUrlQueriesEncoded(request.GetParams())
	}

	ct := contentTypeJSON
	if request.GetHttpMethod() == http.MethodGet {
		ct = contentTypeForm
	} else if request.GetContentType() == "application/octet-stream" {
		ct = contentTypeOctetStream
	}

	// payload: prefer the tchttp.Request itself (V3 JSON marshals it);
	// octet-stream carries raw bytes separately.
	payload := interface{}(request)
	if cr, ok := request.(*tchttp.CommonRequest); ok && cr.IsOctetStream() {
		payload = cr.GetOctetStreamBody()
	} else if request.GetContentType() == "application/octet-stream" {
		payload = request.GetBody()
	}

	// caller headers: tchttp.Request.GetHeader() returns map[string]string;
	// convert to http.Header for direct merge in Build.
	hdr := http.Header{}
	for k, v := range request.GetHeader() {
		hdr.Set(k, v)
	}

	sm := parseSignMethod(c.signMethod)
	if request.GetSkipSign() {
		sm = signMethodSkip
	}

	return &requestBuilder{
		service:         request.GetService(),
		version:         request.GetVersion(),
		action:          request.GetAction(),
		region:          c.GetRegion(),
		url:             u,
		host:            request.GetDomain(),
		method:          request.GetHttpMethod(),
		payload:         payload,
		contentType:     ct,
		headers:         hdr,
		credential:      c.credential,
		signMethod:      sm,
		unsignedPayload: c.unsignedPayload,
		language:        c.profile.Language,
		requestClient:   request.GetParams()["RequestClient"],
		ctx:             request.GetContext(),
		now:             time.Now,
	}
}
