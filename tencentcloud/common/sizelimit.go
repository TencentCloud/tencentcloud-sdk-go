package common

import (
	"fmt"
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
func checkRequestSize(httpMethod, url, signMethod string, limit int64, body []byte) error {
	const (
		KB = 1024
		MB = KB * 1024
	)
	var size int

	switch httpMethod {
	case http.MethodGet: // use user specified limitation
		size = len(url)
		if signMethod != "" && limit > 32*KB {
			limit = 32 * KB
		}

	case http.MethodPost: // use user specified limitation
		size = len(body)

		if (signMethod == SHA256 || signMethod == SHA1) && limit > 1*MB {
			limit = 1 * MB
		} else if signMethod == "TC3-HMAC-SHA256" && limit > 10*MB {
			limit = 10 * MB
		}

	default:
		return fmt.Errorf("invalid request, method: %s, signMethod: %s", httpMethod, signMethod)
	}

	if int64(size) > limit {
		return PacketTooLargeError{size: int64(size), limit: limit}
	}
	return nil
}
