package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type PacketTooLargeError struct {
	size, limit int64
}

func (e PacketTooLargeError) Error() string {
	return fmt.Sprintf("packet too large, limit: %d, size: %d", e.limit, e.size)
}

// The size of a GET request packet is up to 32 KB. The size of a
// POST request is up to 1 MB when the HmacSHA1 or HmacSHA256 signature
// algorithm is used, and up to 10 MB when TC3-HMAC-SHA256 is used.
func checkRequestSize(httpMethod, url, signMethod string, limit int64, body io.Reader) error {
	const (
		KB = 1024
		MB = KB * 1024
	)
	var size int

	switch {
	case httpMethod == http.MethodGet && signMethod == "": // use user specified limitation
		size = len(url)

	case httpMethod == http.MethodGet && signMethod != "":
		if limit > 32*KB {
			limit = 32 * KB
		}
		size = len(url)

	case httpMethod == http.MethodPost && signMethod == "": // use user specified limitation
		bodyBytes, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		size = len(bodyBytes)

	case httpMethod == http.MethodPost && (signMethod == SHA256 || signMethod == SHA1):
		if limit > 1*MB {
			limit = 1 * MB
		}
		bodyBytes, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		size = len(bodyBytes)

	case httpMethod == http.MethodPost && signMethod == "TC3-HMAC-SHA256":
		if limit > 10*MB {
			limit = 10 * MB
		}
		bodyBytes, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		size = len(bodyBytes)
	default:
		return fmt.Errorf("invalid request, method: %s, signMethod: %s", httpMethod, signMethod)
	}

	if int64(size) > limit {
		return PacketTooLargeError{size: int64(size), limit: limit}
	}
	return nil
}
