package common

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const (
	octetStream = "application/octet-stream"
)

// DefaultHttpClient is the default HTTP client used by the SDK.
// It can be overridden for custom HTTP client configurations.
var DefaultHttpClient *http.Client

// Client encapsulates the core functionalities for interacting with Tencent Cloud services.
type Client struct {
	// The region to which the client is connected.
	region string

	// The HTTP client used for making requests.
	httpClient *http.Client

	// The HTTP profile containing endpoint and method settings.
	httpProfile *profile.HttpProfile

	// The client profile containing retry and sign method settings.
	profile *profile.ClientProfile

	// The credential used for authentication.
	credential CredentialIface

	// The signature method used for signing requests (e.g., HmacSHA256, HmacSHA1).
	signMethod string

	// Indicates whether the payload should be unsigned.
	unsignedPayload bool

	// Enables debug logging.
	debug bool

	// The circuit breaker for handling service unavailability.
	rb *circuitBreaker

	// The logger for logging messages.
	logger Logger

	// The client identifier sent in the request header.
	requestClient string
}

// Send sends the request and parses the response.
// It handles request completion, signature, and circuit breaking.
func (c *Client) Send(request tchttp.Request, response tchttp.Response) (err error) {
	c.completeRequest(request)

	tchttp.CompleteCommonParams(request, c.GetRegion(), c.requestClient)

	// reflect to inject client if field ClientToken exists and retry feature is enabled
	if c.profile.NetworkFailureMaxRetries > 0 || c.profile.RateLimitExceededMaxRetries > 0 {
		safeInjectClientToken(request)
	}

	if request.GetSkipSign() {
		// Some APIs allow skipping the signature process.
		return c.sendWithoutSignature(request, response)
	} else if c.profile.DisableRegionBreaker == true || c.rb == nil {
		return c.sendWithSignature(request, response)
	} else {
		return c.sendWithRegionBreaker(request, response)
	}
}

// completeRequest fills in any missing request parameters with default values
// from the client's configuration.
func (c *Client) completeRequest(request tchttp.Request) {
	if request.GetScheme() == "" {
		request.SetScheme(c.httpProfile.Scheme)
	}

	if request.GetRootDomain() == "" {
		request.SetRootDomain(c.httpProfile.RootDomain)
	}

	if request.GetDomain() == "" {
		domain := c.httpProfile.Endpoint
		if domain == "" {
			domain = request.GetServiceDomain(request.GetService())
		}
		request.SetDomain(domain)
	}

	if request.GetHttpMethod() == "" {
		request.SetHttpMethod(c.httpProfile.ReqMethod)
	}

	// Add idempotency key for unsafe retries.
	if c.profile.UnsafeRetryOnConnectionFailure {
		header := request.GetHeader()
		if header == nil {
			header = map[string]string{}
		}
		// http.Transport will automatically retry the request that considered Idempotent
		// see http.Request.isReplayable
		header["X-Idempotency-Key"] = "x"
		request.SetHeader(header)
	}
}

func (c *Client) sendWithRegionBreaker(request tchttp.Request, response tchttp.Response) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			msg := fmt.Sprintf("%s", e)
			err = tcerr.NewTencentCloudSDKError("ClientError.CircuitBreakerError", msg, "")
		}
	}()

	ge, err := c.rb.beforeRequest()

	if err == errOpenState {
		newEndpoint := request.GetService() + "." + c.rb.backupEndpoint
		request.SetDomain(newEndpoint)
	}
	err = c.sendWithSignature(request, response)
	isSuccess := false
	// Success is considered only when the server returns an effective response (have requestId and the code is not InternalError )
	if e, ok := err.(*tcerr.TencentCloudSDKError); ok {
		if e.GetRequestId() != "" && e.GetCode() != "InternalError" {
			isSuccess = true
		}
	}
	c.rb.afterRequest(ge, isSuccess)
	return err
}

func (c *Client) sendWithSignature(request tchttp.Request, response tchttp.Response) (err error) {
	if c.signMethod == "HmacSHA1" || c.signMethod == "HmacSHA256" {
		return c.sendWithSignatureV1(request, response)
	} else {
		return c.sendWithSignatureV3(request, response)
	}
}

func (c *Client) sendWithoutSignature(request tchttp.Request, response tchttp.Response) error {
	headers := map[string]string{
		"Host":               request.GetDomain(),
		"X-TC-Action":        request.GetAction(),
		"X-TC-Version":       request.GetVersion(),
		"X-TC-Timestamp":     request.GetParams()["Timestamp"],
		"X-TC-RequestClient": request.GetParams()["RequestClient"],
		"X-TC-Language":      c.profile.Language,
		"Authorization":      "SKIP",
	}
	if c.region != "" {
		headers["X-TC-Region"] = c.region
	}
	if c.credential != nil {
		credToken := c.credential.GetToken()
		if credToken != "" {
			headers["X-TC-Token"] = credToken
		}
	}
	if request.GetHttpMethod() == "GET" {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	} else {
		headers["Content-Type"] = "application/json"
	}
	isOctetStream := false
	cr := &tchttp.CommonRequest{}
	ok := false
	var octetStreamBody []byte
	if cr, ok = request.(*tchttp.CommonRequest); ok {
		if cr.IsOctetStream() {
			isOctetStream = true
			// custom headers must contain Content-Type : application/octet-stream
			// todo:the custom header may overwrite headers
			for k, v := range cr.GetHeader() {
				headers[k] = v
			}
			octetStreamBody = cr.GetOctetStreamBody()
		}
	}

	for k, v := range request.GetHeader() {
		switch k {
		case "X-TC-Action", "X-TC-Version", "X-TC-Timestamp", "X-TC-RequestClient",
			"X-TC-Language", "Content-Type", "X-TC-Region", "X-TC-Token":
			c.logger.Printf("Skip header \"%s\": can not specify built-in header", k)
		default:
			headers[k] = v
		}
	}

	if !isOctetStream && request.GetContentType() == octetStream {
		isOctetStream = true
		b, _ := json.Marshal(request)
		var m map[string]string
		_ = json.Unmarshal(b, &m)
		for k, v := range m {
			key := "X-" + strings.ToUpper(request.GetService()) + "-" + k
			headers[key] = v
		}

		headers["Content-Type"] = octetStream
		octetStreamBody = request.GetBody()
	}
	// start signature v3 process

	// build canonical request string
	httpRequestMethod := request.GetHttpMethod()
	canonicalQueryString := ""
	if httpRequestMethod == "GET" {
		err := tchttp.ConstructParams(request)
		if err != nil {
			return err
		}
		params := make(map[string]string)
		for key, value := range request.GetParams() {
			params[key] = value
		}
		delete(params, "Action")
		delete(params, "Version")
		delete(params, "Nonce")
		delete(params, "Region")
		delete(params, "RequestClient")
		delete(params, "Timestamp")
		canonicalQueryString = tchttp.GetUrlQueriesEncoded(params)
	}
	requestPayload := ""
	if httpRequestMethod == "POST" {
		if isOctetStream {
			// todo Conversion comparison between string and []byte affects performance much
			requestPayload = string(octetStreamBody)
		} else {
			b, err := json.Marshal(request)
			if err != nil {
				return err
			}
			requestPayload = string(b)
		}
	}
	if c.unsignedPayload {
		headers["X-TC-Content-SHA256"] = "UNSIGNED-PAYLOAD"
	}

	url := request.GetScheme() + "://" + request.GetDomain() + request.GetPath()
	if canonicalQueryString != "" {
		url = url + "?" + canonicalQueryString
	}
	httpRequest, err := http.NewRequest(httpRequestMethod, url, strings.NewReader(requestPayload))
	if err != nil {
		return err
	}
	httpRequest = httpRequest.WithContext(request.GetContext())
	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}
	httpResponse, err := c.sendWithRateLimitRetry(httpRequest, isRetryable(request))
	if err != nil {
		return err
	}
	err = tchttp.ParseFromHttpResponse(httpResponse, response)
	return err
}

func (c *Client) sendWithSignatureV1(request tchttp.Request, response tchttp.Response) (err error) {
	// TODO: not an elegant way, it should be done in common params, but finally it need to refactor
	request.GetParams()["Language"] = c.profile.Language
	err = tchttp.ConstructParams(request)
	if err != nil {
		return err
	}
	err = signRequest(request, c.credential, c.signMethod)
	if err != nil {
		return err
	}
	httpRequest, err := http.NewRequest(request.GetHttpMethod(), request.GetUrl(), request.GetBodyReader())
	if err != nil {
		return err
	}
	httpRequest = httpRequest.WithContext(request.GetContext())
	if request.GetHttpMethod() == "POST" {
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	for k, v := range request.GetHeader() {
		httpRequest.Header.Set(k, v)
	}

	httpResponse, err := c.sendWithRateLimitRetry(httpRequest, isRetryable(request))
	if err != nil {
		return err
	}
	err = tchttp.ParseFromHttpResponse(httpResponse, response)
	return err
}

func (c *Client) sendWithSignatureV3(request tchttp.Request, response tchttp.Response) (err error) {
	// Prepare the headers for the request, including essential information
	// and the necessary components for Signature Version 3 authentication.
	headers := map[string]string{
		"Host":               request.GetDomain(),
		"X-TC-Action":        request.GetAction(),
		"X-TC-Version":       request.GetVersion(),
		"X-TC-Timestamp":     request.GetParams()["Timestamp"],
		"X-TC-RequestClient": request.GetParams()["RequestClient"],
		"X-TC-Language":      c.profile.Language,
	}

	// Include the region if specified.
	if c.region != "" {
		headers["X-TC-Region"] = c.region
	}

	// Retrieve the secret ID, secret key, and security token from the credentials.
	secId, secKey, token := c.credential.GetCredential()
	if token != "" {
		headers["X-TC-Token"] = token
	}

	// Set the Content-Type header based on the HTTP method.
	if request.GetHttpMethod() == "GET" {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	} else {
		headers["Content-Type"] = "application/json"
	}

	// Handle octet-stream (binary data) requests.
	isOctetStream := false
	cr := &tchttp.CommonRequest{}
	ok := false
	var octetStreamBody []byte
	if cr, ok = request.(*tchttp.CommonRequest); ok {
		if cr.IsOctetStream() {
			isOctetStream = true
			// custom headers must contain Content-Type : application/octet-stream
			// todo:the custom header may overwrite headers
			for k, v := range cr.GetHeader() {
				headers[k] = v
			}
			octetStreamBody = cr.GetOctetStreamBody()
		}
	}

	// Merge any additional headers from the request, but skip built-in headers
	// to prevent them from being overridden.
	for k, v := range request.GetHeader() {
		switch k {
		case "X-TC-Action", "X-TC-Version", "X-TC-Timestamp", "X-TC-RequestClient",
			"X-TC-Language", "X-TC-Region", "X-TC-Token":
			c.logger.Printf("Skip header \"%s\": can not specify built-in header", k)
		default:
			headers[k] = v
		}
	}

	// Handle the case where the request content type is explicitly set to octet-stream,
	// but it's not already handled as an OctetStream CommonRequest.
	if !isOctetStream && request.GetContentType() == octetStream {
		isOctetStream = true
		b, _ := json.Marshal(request)
		var m map[string]string
		_ = json.Unmarshal(b, &m)
		for k, v := range m {
			key := "X-" + strings.ToUpper(request.GetService()) + "-" + k
			headers[key] = v
		}

		headers["Content-Type"] = octetStream
		octetStreamBody = request.GetBody()
	}
	// --- Begin Signature Version 3 (TC3-HMAC-SHA256) Signing Process ---

	// 1. Construct the Canonical Request

	// HTTP Method (e.g., "GET", "POST").
	httpRequestMethod := request.GetHttpMethod()

	// Canonical URI (always "/").
	canonicalURI := "/"

	// Canonical Query String (for GET requests).
	canonicalQueryString := ""
	if httpRequestMethod == "GET" {
		err = tchttp.ConstructParams(request)
		if err != nil {
			return err
		}
		params := make(map[string]string)
		for key, value := range request.GetParams() {
			params[key] = value
		}
		// Remove standard parameters that are not part of the canonical query string.
		delete(params, "Action")
		delete(params, "Version")
		delete(params, "Nonce")
		delete(params, "Region")
		delete(params, "RequestClient")
		delete(params, "Timestamp")
		canonicalQueryString = tchttp.GetUrlQueriesEncoded(params)
	}

	// Canonical Headers (sorted and formatted).
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\n", headers["Content-Type"], headers["Host"])

	// Signed Headers (list of headers included in the signature).
	signedHeaders := "content-type;host"

	// Request Payload (for POST requests).
	requestPayload := ""
	if httpRequestMethod == "POST" {
		if isOctetStream {
			// todo Conversion comparison between string and []byte affects performance much
			requestPayload = string(octetStreamBody)
		} else {
			b, err := json.Marshal(request)
			if err != nil {
				return err
			}
			requestPayload = string(b)
		}
	}

	// Hashed Request Payload.
	hashedRequestPayload := ""
	if c.unsignedPayload {
		hashedRequestPayload = sha256hex("UNSIGNED-PAYLOAD")
		headers["X-TC-Content-SHA256"] = "UNSIGNED-PAYLOAD"
	} else {
		hashedRequestPayload = sha256hex(requestPayload)
	}

	// Construct the complete Canonical Request String.
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	//log.Println("canonicalRequest:", canonicalRequest)

	// 2. Construct the String to Sign

	// Algorithm.
	algorithm := "TC3-HMAC-SHA256"

	// Request Timestamp.
	requestTimestamp := headers["X-TC-Timestamp"]

	// Credential Scope.
	timestamp, _ := strconv.ParseInt(requestTimestamp, 10, 64)
	t := time.Unix(timestamp, 0).UTC()
	// must be the format 2006-01-02, ref to package time for more info
	date := t.Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, request.GetService())

	// Hashed Canonical Request.
	hashedCanonicalRequest := sha256hex(canonicalRequest)

	// Construct the String to Sign.
	string2sign := fmt.Sprintf("%s\n%s\n%s\n%s",
		algorithm,
		requestTimestamp,
		credentialScope,
		hashedCanonicalRequest)
	//log.Println("string2sign", string2sign)

	// 3. Calculate the Signature

	// Secret Date.
	secretDate := hmacsha256(date, "TC3"+secKey)

	// Secret Service.
	secretService := hmacsha256(request.GetService(), secretDate)

	// Secret Key.
	secretKey := hmacsha256("tc3_request", secretService)

	// Signature.
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretKey)))
	//log.Println("signature", signature)

	// 4. Construct the Authorization Header

	// Authorization Header.
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secId,
		credentialScope,
		signedHeaders,
		signature)
	//log.Println("authorization", authorization)

	// Add the Authorization header to the request headers.
	headers["Authorization"] = authorization

	// --- End Signature Version 3 Signing Process ---

	// Construct the full URL.
	url := request.GetScheme() + "://" + request.GetDomain() + request.GetPath()
	if canonicalQueryString != "" {
		url = url + "?" + canonicalQueryString
	}

	// Create the HTTP request.
	httpRequest, err := http.NewRequest(httpRequestMethod, url, strings.NewReader(requestPayload))
	if err != nil {
		return err
	}
	httpRequest = httpRequest.WithContext(request.GetContext())

	// Set all the headers on the request.
	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}

	// Send the HTTP request with rate limit retry logic.
	httpResponse, err := c.sendWithRateLimitRetry(httpRequest, isRetryable(request))
	if err != nil {
		return err
	}

	// Parse the HTTP response into the specified response object.
	return tchttp.ParseFromHttpResponse(httpResponse, response)
}

// send http request
func (c *Client) sendHttp(request *http.Request) (response *http.Response, err error) {
	if len(c.httpProfile.ApigwEndpoint) > 0 {
		request.URL.Host = c.httpProfile.ApigwEndpoint
		request.Host = c.httpProfile.ApigwEndpoint
	}

	if c.debug && request != nil {
		outBytes, err := httputil.DumpRequest(request, true)
		if err != nil {
			c.logger.Printf("[ERROR] dump request failed: %s", err)
		} else {
			c.logger.Printf("[DEBUG] http request: %s", outBytes)
		}
	}

	response, err = c.httpClient.Do(request)

	if c.debug && response != nil {
		dumpBody := true
		switch response.Header.Get("Content-Type") {
		case "text/event-stream", "application/octet-stream":
			dumpBody = false
		}

		out, err := httputil.DumpResponse(response, dumpBody)
		if err != nil {
			c.logger.Printf("[ERROR] dump response failed: %s", err)
		} else {
			c.logger.Printf("[DEBUG] http response: %s", out)
		}
	}

	return response, err
}

func (c *Client) GetRegion() string {
	return c.region
}

func (c *Client) Init(region string) *Client {
	const defaultIdleConnTimeout = 30 * time.Second

	if DefaultHttpClient == nil {
		// try not to modify http.DefaultTransport if possible
		// since we could possibly modify Transport.Proxy
		transport := http.DefaultTransport
		if _, ok := transport.(*http.Transport); ok {
			// http.Transport.Clone is only available after go1.12
			if cloneMethod, hasClone := reflect.TypeOf(transport).MethodByName("Clone"); hasClone {
				cloned := cloneMethod.Func.Call([]reflect.Value{reflect.ValueOf(transport)})[0].Interface().(http.RoundTripper)
				if clonedTransport, ok := cloned.(*http.Transport); ok {
					clonedTransport.IdleConnTimeout = defaultIdleConnTimeout
					transport = clonedTransport
				}
			}
		}

		c.httpClient = &http.Client{Transport: transport}
	} else {
		c.httpClient = DefaultHttpClient
	}

	c.region = region
	c.signMethod = "TC3-HMAC-SHA256"
	c.debug = false
	c.logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	return c
}

func (c *Client) WithSecretId(secretId, secretKey string) *Client {
	c.credential = NewCredential(secretId, secretKey)
	return c
}

func (c *Client) WithCredential(cred CredentialIface) *Client {
	c.credential = cred
	return c
}

func (c *Client) WithRequestClient(rc string) *Client {
	const reRequestClient = "^[0-9a-zA-Z-_ ,;.]+$"

	if len(rc) > 128 {
		c.logger.Printf("the length of RequestClient should be within 128 characters, it will be truncated")
		rc = rc[:128]
	}

	match, err := regexp.MatchString(reRequestClient, rc)
	if err != nil {
		c.logger.Printf("regexp is wrong: %s", reRequestClient)
		return c
	}
	if !match {
		c.logger.Printf("RequestClient not match the regexp: %s, ignored", reRequestClient)
		return c
	}

	c.requestClient = rc
	return c
}

func (c *Client) GetCredential() CredentialIface {
	return c.credential
}

func (c *Client) WithProfile(clientProfile *profile.ClientProfile) *Client {
	c.profile = clientProfile
	if c.profile.DisableRegionBreaker == false {
		c.withRegionBreaker()
	}
	c.signMethod = clientProfile.SignMethod
	c.unsignedPayload = clientProfile.UnsignedPayload
	c.httpProfile = clientProfile.HttpProfile
	c.debug = clientProfile.Debug
	c.httpClient.Timeout = time.Duration(c.httpProfile.ReqTimeout) * time.Second
	if c.httpProfile.Proxy != "" {
		u, err := url.Parse(c.httpProfile.Proxy)
		if err != nil {
			panic(err)
		}

		if c.httpClient.Transport == nil {
			c.logger.Printf("trying to set proxy when httpClient.Transport is nil")
		}

		if _, ok := c.httpClient.Transport.(*http.Transport); ok {
			c.httpClient.Transport.(*http.Transport).Proxy = http.ProxyURL(u)
		} else {
			c.logger.Printf("setting proxy while httpClient.Transport is not a http.Transport is not supported")
		}
	}
	return c
}

func (c *Client) WithSignatureMethod(method string) *Client {
	c.signMethod = method
	return c
}

func (c *Client) WithHttpTransport(transport http.RoundTripper) *Client {
	c.httpClient.Transport = transport
	return c
}

func (c *Client) WithDebug(flag bool) *Client {
	c.debug = flag
	return c
}

// WithProvider use specify provider to get a credential and use it to build a client
func (c *Client) WithProvider(provider Provider) (*Client, error) {
	cred, err := provider.GetCredential()
	if err != nil {
		return nil, err
	}
	return c.WithCredential(cred), nil
}

func (c *Client) withRegionBreaker() *Client {
	rb := defaultRegionBreaker()
	if c.profile.BackupEndpoint != "" {
		rb.backupEndpoint = c.profile.BackupEndpoint
	} else if c.profile.BackupEndPoint != "" {
		rb.backupEndpoint = c.profile.BackupEndPoint
	}
	c.rb = rb
	return c
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey)
	return
}

// NewClientWithProviders build client with your custom providers;
// If you don't specify the providers, it will use the DefaultProviderChain to find credential
func NewClientWithProviders(region string, providers ...Provider) (client *Client, err error) {
	client = (&Client{}).Init(region)
	var pc Provider
	if len(providers) == 0 {
		pc = DefaultProviderChain()
	} else {
		pc = NewProviderChain(providers)
	}
	return client.WithProvider(pc)
}

func (c *Client) InitBaseRequest(req **tchttp.BaseRequest, mod, ver, act string) {
	if *req == nil {
		*req = &tchttp.BaseRequest{}
	}

	if (*req).GetParams() == nil {
		(*req).Init()
	}

	if (*req).GetAction() == "" {
		(*req).WithApiInfo(mod, ver, act)
	}
}
