package common

import (
	"testing"
)

func TestIsRetryable(t *testing.T) {
	examples := map[interface{}]bool{
		new(injectable):   true,
		new(uninjectable): false,
	}
	for msg, expected := range examples {
		if val := isRetryable(msg); val != expected {
			t.Fatalf("retryable failed: expected %+v, got %+v", expected, val)
		}
	}
}
