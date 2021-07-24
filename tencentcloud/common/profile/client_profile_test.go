package profile

import (
	"math/rand"
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	expected := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		16 * time.Second,
		32 * time.Second,
		64 * time.Second,
		128 * time.Second,
	}
	for i := 0; i < len(expected); i++ {
		if ExponentialBackoff(i) != expected[i] {
			t.Fatalf("unexpected retry time, %+v expected, got %+v", expected[i], ExponentialBackoff(i))
		}
	}
}

func TestConstantDurationFunc(t *testing.T) {
	wanted := time.Duration(rand.Int()%100) * time.Second
	actual := ConstantDurationFunc(wanted)(rand.Int())
	if actual != wanted {
		t.Fatalf("unexpected retry time, %+v expected, got %+v", wanted, actual)
	}
}
