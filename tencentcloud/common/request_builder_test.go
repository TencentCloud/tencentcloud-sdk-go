package common

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

// fixedTime is the canonical test timestamp: 2024-01-01 00:00:00 UTC.
const (
	fixedUnix    = 1704067200
	fixedSecID   = "AKIDTEST"
	fixedSecKey  = "SKTEST"
	fixedToken   = ""
	fixedService = "cvm"
	fixedVersion = "2017-03-12"
	fixedAction  = "DescribeInstances"
	fixedRegion  = "ap-guangzhou"
	fixedHost    = "cvm.tencentcloudapi.com"
)

var fixedTime = time.Unix(fixedUnix, 0).UTC()

// newTestBuilder returns a requestBuilder with all common fields filled, ready
// to be specialized per-test (override payload/method/signMethod/etc).
func newTestBuilder() *requestBuilder {
	return &requestBuilder{
		service:       fixedService,
		version:       fixedVersion,
		action:        fixedAction,
		region:        fixedRegion,
		url:           &url.URL{Scheme: "https", Host: fixedHost, Path: "/"},
		host:          fixedHost,
		method:        http.MethodPost,
		contentType:   contentTypeJSON,
		credential:    NewCredential(fixedSecID, fixedSecKey),
		signMethod:    signMethodV3,
		language:      "zh-CN",
		requestClient: "SDK_GO_TEST",
		now:           func() time.Time { return fixedTime },
		nonce:         12345,
	}
}

// expectedV3Auth recomputes the TC3-HMAC-SHA256 Authorization header
// independently from requestBuilder.buildV3, so tests catch regressions in
// canonical-request assembly, header ordering, and body serialization — not
// just in the crypto primitives (which are covered by sign_test.go).
func expectedV3Auth(t *testing.T, method, host, ctStr, canonicalQuery, body, service string) string {
	t.Helper()
	return expectedV3AuthWithPayloadHash(t, method, host, ctStr, canonicalQuery, sha256hex(body), service)
}

func expectedV3AuthWithPayloadHash(t *testing.T, method, host, ctStr, canonicalQuery, hashedPayload, service string) string {
	t.Helper()
	ts := strconv.FormatInt(int64(fixedUnix), 10)
	date := fixedTime.Format("2006-01-02")

	canonicalURI := "/"
	canonicalHeaders := "content-type:" + ctStr + "\nhost:" + host + "\n"
	signedHeaders := "content-type;host"
	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQuery + "\n" +
		canonicalHeaders + "\n" + signedHeaders + "\n" + hashedPayload

	credentialScope := date + "/" + service + "/tc3_request"
	stringToSign := "TC3-HMAC-SHA256\n" + ts + "\n" + credentialScope + "\n" +
		sha256hex(canonicalRequest)

	secretDate := hmacsha256(date, "TC3"+fixedSecKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(stringToSign, secretSigning)))

	return "TC3-HMAC-SHA256 Credential=" + fixedSecID + "/" + credentialScope +
		", SignedHeaders=" + signedHeaders + ", Signature=" + signature
}

// =====================================================================
//  V3 (TC3-HMAC-SHA256)
// =====================================================================

func Test_requestBuilder_buildV3PostJSON(t *testing.T) {
	rb := newTestBuilder()
	rb.payload = map[string]interface{}{"Limit": 10, "Offset": 0}

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}

	if req.Method != http.MethodPost {
		t.Fatalf("Method = %q, want POST", req.Method)
	}
	if req.Host != fixedHost {
		t.Fatalf("req.Host = %q, want %q", req.Host, fixedHost)
	}
	if req.URL.Host != fixedHost {
		t.Fatalf("URL.Host = %q, want %q", req.URL.Host, fixedHost)
	}
	if got := req.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", got)
	}
	if got := req.Header.Get("X-TC-Action"); got != fixedAction {
		t.Fatalf("X-TC-Action = %q, want %q", got, fixedAction)
	}
	if got := req.Header.Get("X-TC-Version"); got != fixedVersion {
		t.Fatalf("X-TC-Version = %q, want %q", got, fixedVersion)
	}
	if got := req.Header.Get("X-TC-Region"); got != fixedRegion {
		t.Fatalf("X-TC-Region = %q, want %q", got, fixedRegion)
	}
	if got := req.Header.Get("X-TC-Timestamp"); got != strconv.FormatInt(int64(fixedUnix), 10) {
		t.Fatalf("X-TC-Timestamp = %q, want %d", got, fixedUnix)
	}
	if got := req.Header.Get("X-TC-Language"); got != "zh-CN" {
		t.Fatalf("X-TC-Language = %q, want zh-CN", got)
	}

	body, _ := ioutil.ReadAll(req.Body)
	wantAuth := expectedV3Auth(t, http.MethodPost, fixedHost, "application/json", "", string(body), fixedService)
	if got := req.Header.Get("Authorization"); got != wantAuth {
		t.Fatalf("Authorization mismatch:\n got: %s\nwant: %s", got, wantAuth)
	}
}

func Test_requestBuilder_buildV3Get(t *testing.T) {
	rb := newTestBuilder()
	rb.method = http.MethodGet
	rb.contentType = contentTypeForm
	rb.payload = map[string]interface{}{"Limit": 10}

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}

	if req.Method != http.MethodGet {
		t.Fatalf("Method = %q, want GET", req.Method)
	}
	if got := req.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
		t.Fatalf("Content-Type = %q, want form", got)
	}
	// Body must be empty for GET.
	if req.Body != nil {
		b, _ := ioutil.ReadAll(req.Body)
		if len(b) != 0 {
			t.Fatalf("GET body = %q, want empty", string(b))
		}
	}
	// Query carries the params (Limit), not the SDK-common ones.
	if !strings.Contains(req.URL.RawQuery, "Limit=10") {
		t.Fatalf("RawQuery = %q, want it to contain Limit=10", req.URL.RawQuery)
	}
	for _, banned := range []string{"Action", "Version", "Region", "Timestamp", "Nonce", "RequestClient"} {
		if strings.Contains(req.URL.RawQuery, banned+"=") {
			t.Fatalf("RawQuery = %q, must not contain %s", req.URL.RawQuery, banned)
		}
	}

	wantQuery := url.Values{"Limit": {"10"}}.Encode()
	wantAuth := expectedV3Auth(t, http.MethodGet, fixedHost, "application/x-www-form-urlencoded", wantQuery, "", fixedService)
	if got := req.Header.Get("Authorization"); got != wantAuth {
		t.Fatalf("Authorization mismatch:\n got: %s\nwant: %s", got, wantAuth)
	}
}

func Test_requestBuilder_buildV3OctetStream(t *testing.T) {
	rb := newTestBuilder()
	rb.contentType = contentTypeOctetStream
	rb.payload = []byte("raw-bytes-payload")

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	if got := req.Header.Get("Content-Type"); got != "application/octet-stream" {
		t.Fatalf("Content-Type = %q, want octet-stream", got)
	}
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != "raw-bytes-payload" {
		t.Fatalf("body = %q, want raw-bytes-payload", string(body))
	}
}

func Test_requestBuilder_buildV3BytesJSONPayload(t *testing.T) {
	// Pre-serialized JSON []byte should be used as-is.
	preJSON := `{"Limit":10}`
	rb := newTestBuilder()
	rb.payload = []byte(preJSON)

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != preJSON {
		t.Fatalf("body = %q, want %s", string(body), preJSON)
	}
	wantAuth := expectedV3Auth(t, http.MethodPost, fixedHost, "application/json", "", preJSON, fixedService)
	if got := req.Header.Get("Authorization"); got != wantAuth {
		t.Fatalf("Authorization mismatch:\n got: %s\nwant: %s", got, wantAuth)
	}
}

// =====================================================================
//  V1 (HmacSHA1 / HmacSHA256)
// =====================================================================

func Test_requestBuilder_buildV1Post(t *testing.T) {
	for _, sm := range []signMethod{signMethodHmacSHA1, signMethodHmacSHA256} {
		t.Run(sm.String(), func(t *testing.T) {
			rb := newTestBuilder()
			rb.signMethod = sm
			rb.contentType = contentTypeForm
			rb.payload = map[string]interface{}{"Limit": 10}

			req, err := rb.build()
			if err != nil {
				t.Fatalf("build() error: %v", err)
			}
			if got := req.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
				t.Fatalf("Content-Type = %q, want form", got)
			}
			body, _ := ioutil.ReadAll(req.Body)
			if !strings.Contains(string(body), "Signature=") {
				t.Fatalf("body = %q, want it to contain Signature=", string(body))
			}
			if !strings.Contains(string(body), "SignatureMethod="+sm.String()) {
				t.Fatalf("body = %q, want SignatureMethod=%s", string(body), sm.String())
			}
			if !strings.Contains(string(body), "SecretId="+fixedSecID) {
				t.Fatalf("body = %q, want SecretId=%s", string(body), fixedSecID)
			}
			if !strings.Contains(string(body), "Action="+fixedAction) {
				t.Fatalf("body = %q, want Action=%s", string(body), fixedAction)
			}
			// Authorization header must NOT be set for V1.
			if got := req.Header.Get("Authorization"); got != "" {
				t.Fatalf("Authorization = %q, want empty for V1", got)
			}
		})
	}
}

func Test_requestBuilder_buildV1Get(t *testing.T) {
	rb := newTestBuilder()
	rb.signMethod = signMethodHmacSHA256
	rb.method = http.MethodGet
	rb.contentType = contentTypeForm
	rb.payload = map[string]interface{}{"Limit": 10}

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	if !strings.Contains(req.URL.RawQuery, "Signature=") {
		t.Fatalf("RawQuery = %q, want Signature=", req.URL.RawQuery)
	}
	if req.Body != nil {
		b, _ := ioutil.ReadAll(req.Body)
		if len(b) != 0 {
			t.Fatalf("GET body = %q, want empty", string(b))
		}
	}
}

// =====================================================================
//  SKIP
// =====================================================================

func Test_requestBuilder_buildSkip(t *testing.T) {
	rb := newTestBuilder()
	rb.signMethod = signMethodSkip
	rb.payload = map[string]interface{}{"Limit": 10}

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	if got := req.Header.Get("Authorization"); got != "SKIP" {
		t.Fatalf("Authorization = %q, want SKIP", got)
	}
	if got := req.Header.Get("X-TC-Action"); got != fixedAction {
		t.Fatalf("X-TC-Action = %q, want %q", got, fixedAction)
	}
	if got := req.Header.Get("X-TC-Timestamp"); got != strconv.FormatInt(int64(fixedUnix), 10) {
		t.Fatalf("X-TC-Timestamp = %q, want %d", got, fixedUnix)
	}
	if got := req.Header.Get("X-TC-RequestClient"); got != "SDK_GO_TEST" {
		t.Fatalf("X-TC-RequestClient = %q, want SDK_GO_TEST", got)
	}
	if got := req.Header.Get("X-TC-Language"); got != "zh-CN" {
		t.Fatalf("X-TC-Language = %q, want zh-CN", got)
	}
}

// =====================================================================
//  Caller headers merge
// =====================================================================

func Test_requestBuilder_callerHeadersMerged(t *testing.T) {
	rb := newTestBuilder()
	rb.headers = http.Header{}
	rb.headers.Set("X-Custom", "custom-val")
	rb.payload = map[string]interface{}{"Limit": 1}

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	if got := req.Header.Get("X-Custom"); got != "custom-val" {
		t.Fatalf("X-Custom = %q, want custom-val", got)
	}
}

func Test_requestBuilder_signatureHeadersNotOverwrittenByCaller(t *testing.T) {
	// Caller tries to overwrite Authorization — must be ignored (signature wins).
	rb := newTestBuilder()
	rb.headers = http.Header{}
	rb.headers.Set("Authorization", "EVIL")
	rb.payload = map[string]interface{}{"Limit": 1}

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	if got := req.Header.Get("Authorization"); got == "EVIL" || !strings.HasPrefix(got, "TC3-HMAC-SHA256 ") {
		t.Fatalf("Authorization = %q, want a real V3 signature (caller overwrite must be ignored)", got)
	}
}

// =====================================================================
//  host vs url.Host separation (proxy scenario)
// =====================================================================

func Test_requestBuilder_hostDiffersFromURLHost(t *testing.T) {
	rb := newTestBuilder()
	rb.url.Host = "proxy.internal:8080"
	rb.host = fixedHost
	rb.payload = map[string]interface{}{"Limit": 1}

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	// Transport target stays the proxy.
	if req.URL.Host != "proxy.internal:8080" {
		t.Fatalf("URL.Host = %q, want proxy.internal:8080", req.URL.Host)
	}
	// Signing Host header is the real service host.
	if req.Host != fixedHost {
		t.Fatalf("req.Host = %q, want %q", req.Host, fixedHost)
	}
	// Authorization signs against the real host, not the proxy.
	body, _ := ioutil.ReadAll(req.Body)
	wantAuth := expectedV3Auth(t, http.MethodPost, fixedHost, "application/json", "", string(body), fixedService)
	if got := req.Header.Get("Authorization"); got != wantAuth {
		t.Fatalf("Authorization must sign against %q, not the proxy:\n got: %s\nwant: %s",
			fixedHost, got, wantAuth)
	}
}

// =====================================================================
//  Context plumbing
// =====================================================================

func Test_withRequestBuilderRoundTrip(t *testing.T) {
	rb := newTestBuilder()
	ctx := context.Background()
	ctx2 := withRequestBuilder(ctx, rb)
	if got := requestBuilderFromContext(ctx2); got != rb {
		t.Fatalf("round-trip: got %p, want %p", got, rb)
	}
	// Empty context returns nil.
	if got := requestBuilderFromContext(context.Background()); got != nil {
		t.Fatalf("empty ctx: got %p, want nil", got)
	}
}

// =====================================================================
//  parseSignMethod
// =====================================================================

func Test_parseSignMethod(t *testing.T) {
	cases := []struct {
		in   string
		want signMethod
	}{
		{"TC3-HMAC-SHA256", signMethodV3},
		{"HmacSHA256", signMethodHmacSHA256},
		{"HmacSHA1", signMethodHmacSHA1},
		{"SKIP", signMethodSkip},
		{"", signMethodV3},      // unknown defaults to V3
		{"bogus", signMethodV3}, // unknown defaults to V3
	}
	for _, c := range cases {
		if got := parseSignMethod(c.in); got != c.want {
			t.Fatalf("parseSignMethod(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// =====================================================================
//  contentType.String
// =====================================================================

func Test_contentTypeString(t *testing.T) {
	cases := []struct {
		ct   contentType
		want string
	}{
		{contentTypeJSON, "application/json"},
		{contentTypeForm, "application/x-www-form-urlencoded"},
		{contentTypeOctetStream, "application/octet-stream"},
	}
	for _, c := range cases {
		if got := c.ct.String(); got != c.want {
			t.Fatalf("%v.String() = %q, want %q", c.ct, got, c.want)
		}
	}
}

// =====================================================================
//  now() freshness across rebuilds
// =====================================================================

func Test_requestBuilder_nowReEvaluatedOnEachBuild(t *testing.T) {
	// First build at t0, second at t0+60s — timestamps must differ.
	t0 := time.Unix(fixedUnix, 0).UTC()
	calls := 0
	rb := newTestBuilder()
	rb.payload = map[string]interface{}{"Limit": 1}
	rb.now = func() time.Time {
		calls++
		return t0.Add(time.Duration(calls-1) * 60 * time.Second)
	}

	req1, _ := rb.build()
	ts1 := req1.Header.Get("X-TC-Timestamp")
	req2, _ := rb.build()
	ts2 := req2.Header.Get("X-TC-Timestamp")

	if ts1 == ts2 {
		t.Fatalf("timestamp not refreshed: ts1=%s ts2=%s (want different)", ts1, ts2)
	}
	want1 := strconv.FormatInt(int64(fixedUnix), 10)
	want2 := strconv.FormatInt(int64(fixedUnix+60), 10)
	if ts1 != want1 || ts2 != want2 {
		t.Fatalf("timestamps = (%s, %s), want (%s, %s)", ts1, ts2, want1, want2)
	}
}

// =====================================================================
//  newRequestBuilderFromRequest bridge
// =====================================================================

// minimal tchttp.Request for testing the bridge. Uses CommonRequest directly.
func newTchttpRequestForBridge() tchttp.Request {
	r := tchttp.NewCommonRequest(fixedService, fixedVersion, fixedAction)
	r.SetDomain(fixedHost)
	r.SetScheme("https")
	r.SetPath("/")
	r.SetHttpMethod("POST")
	r.GetParams()["RequestClient"] = "SDK_GO_TEST"
	return r
}

func Test_newRequestBuilderFromRequest(t *testing.T) {
	cpf := profile.NewClientProfile()
	c := NewCommonClient(NewCredential(fixedSecID, fixedSecKey), fixedRegion, cpf)
	tcr := newTchttpRequestForBridge()

	rb := newRequestBuilderFromRequest(c, tcr)

	if rb.service != fixedService {
		t.Fatalf("service = %q, want %q", rb.service, fixedService)
	}
	if rb.action != fixedAction {
		t.Fatalf("action = %q, want %q", rb.action, fixedAction)
	}
	if rb.host != fixedHost {
		t.Fatalf("host = %q, want %q", rb.host, fixedHost)
	}
	if rb.url.Host != fixedHost {
		t.Fatalf("url.Host = %q, want %q", rb.url.Host, fixedHost)
	}
	if rb.url.Scheme != "https" {
		t.Fatalf("url.Scheme = %q, want https", rb.url.Scheme)
	}
	if rb.method != http.MethodPost {
		t.Fatalf("method = %q, want POST", rb.method)
	}
	if rb.region != fixedRegion {
		t.Fatalf("region = %q, want %q", rb.region, fixedRegion)
	}
	if rb.signMethod != signMethodV3 {
		t.Fatalf("signMethod = %v, want V3", rb.signMethod)
	}
	// Bridge must produce a buildable request.
	if _, err := rb.build(); err != nil {
		t.Fatalf("bridge-built requestBuilder.build() error: %v", err)
	}
}

// =====================================================================
//  tchttp.Request payload (json.Marshal)
// =====================================================================

func Test_requestBuilder_buildV3WithTchttpRequestPayload(t *testing.T) {
	tcr := tchttp.NewCommonRequest(fixedService, fixedVersion, fixedAction)
	// CommonRequest has no user-set fields here, so Marshal yields {}.
	rb := newTestBuilder()
	rb.payload = tcr

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	body, _ := ioutil.ReadAll(req.Body)
	// Body must be valid JSON (whatever Marshal produced).
	if !json.Valid(body) {
		t.Fatalf("body = %q, want valid JSON", string(body))
	}
	wantAuth := expectedV3Auth(t, http.MethodPost, fixedHost, "application/json", "", string(body), fixedService)
	if got := req.Header.Get("Authorization"); got != wantAuth {
		t.Fatalf("Authorization mismatch:\n got: %s\nwant: %s", got, wantAuth)
	}
}

func Test_requestBuilder_buildV3UnsignedPayload(t *testing.T) {
	rb := newTestBuilder()
	rb.payload = map[string]interface{}{"Limit": 10}
	rb.unsignedPayload = true

	req, err := rb.build()
	if err != nil {
		t.Fatalf("build() error: %v", err)
	}
	if got := req.Header.Get("X-TC-Content-SHA256"); got != "UNSIGNED-PAYLOAD" {
		t.Fatalf("X-TC-Content-SHA256 = %q, want UNSIGNED-PAYLOAD", got)
	}
	wantAuth := expectedV3AuthWithPayloadHash(t, http.MethodPost, fixedHost, "application/json", "", sha256hex("UNSIGNED-PAYLOAD"), fixedService)
	if got := req.Header.Get("Authorization"); got != wantAuth {
		t.Fatalf("Authorization mismatch:\n got: %s\nwant: %s", got, wantAuth)
	}
}

func Test_requestBuilder_buildPreservesContext(t *testing.T) {
	type ctxKey struct{}
	ctx := context.WithValue(context.Background(), ctxKey{}, "marker")
	cases := []struct {
		name string
		sm   signMethod
	}{
		{"v3", signMethodV3},
		{"v1", signMethodHmacSHA256},
		{"skip", signMethodSkip},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rb := newTestBuilder()
			rb.signMethod = tc.sm
			rb.ctx = ctx
			if tc.sm == signMethodSkip {
				rb.credential = nil
			}
			rb.payload = map[string]interface{}{"Limit": 10}

			req, err := rb.build()
			if err != nil {
				t.Fatalf("build() error: %v", err)
			}
			if got := req.Context().Value(ctxKey{}); got != "marker" {
				t.Fatalf("context value = %v, want marker", got)
			}
		})
	}
}

func Test_requestBuilder_paramsFromPayloadCopiesSourceMaps(t *testing.T) {
	rb := newTestBuilder()
	rb.method = http.MethodGet
	rb.contentType = contentTypeForm
	params := map[string]string{
		"Action":        fixedAction,
		"Version":       fixedVersion,
		"RequestClient": "SDK_GO_TEST",
		"Limit":         "10",
	}
	rb.payload = params

	if _, _, err := rb.v3BodyAndQuery(); err != nil {
		t.Fatalf("v3BodyAndQuery() error: %v", err)
	}
	for _, k := range []string{"Action", "Version", "RequestClient"} {
		if _, ok := params[k]; !ok {
			t.Fatalf("params[%q] was deleted from source map", k)
		}
	}

	tcr := tchttp.NewCommonRequest(fixedService, fixedVersion, fixedAction)
	tcr.SetScheme("https")
	tcr.SetDomain(fixedHost)
	tcr.SetPath("/")
	tcr.SetHttpMethod(http.MethodGet)
	tcr.GetParams()["Action"] = fixedAction
	tcr.GetParams()["Version"] = fixedVersion
	tcr.GetParams()["RequestClient"] = "SDK_GO_TEST"
	tcr.GetParams()["Limit"] = "10"
	rb.payload = tcr
	if _, _, err := rb.v3BodyAndQuery(); err != nil {
		t.Fatalf("v3BodyAndQuery() with tchttp.Request error: %v", err)
	}
	for _, k := range []string{"Action", "Version", "RequestClient"} {
		if _, ok := tcr.GetParams()[k]; !ok {
			t.Fatalf("tchttp request params[%q] was deleted", k)
		}
	}
}

// =====================================================================
//  Error paths
// =====================================================================

func Test_requestBuilder_unsupportedSignMethod(t *testing.T) {
	rb := newTestBuilder()
	rb.signMethod = signMethod(99) // bogus
	rb.payload = map[string]interface{}{"Limit": 1}

	if _, err := rb.build(); err == nil {
		t.Fatalf("build() with bogus signMethod: want error, got nil")
	}
}

func Test_requestBuilder_v3PostUnsupportedPayloadType(t *testing.T) {
	rb := newTestBuilder()
	rb.payload = 12345 // int, not a supported payload type

	if _, err := rb.build(); err == nil {
		t.Fatalf("build() with int payload: want error, got nil")
	}
}

func Test_requestBuilder_octetStreamNonBytes(t *testing.T) {
	rb := newTestBuilder()
	rb.contentType = contentTypeOctetStream
	rb.payload = "not-bytes"

	if _, err := rb.build(); err == nil {
		t.Fatalf("build() with string octet-stream payload: want error, got nil")
	}
}

// avoid unused import warnings (bytes used conditionally in some builds).
var _ = bytes.NewReader
