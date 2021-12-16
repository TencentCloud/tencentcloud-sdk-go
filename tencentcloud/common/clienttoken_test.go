package common

import (
	"testing"
)

type injectable struct {
	ClientToken *string
}

type uninjectable struct {
}

func TestSafeInjectClientToken(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			t.Fatalf("panic on injecting client token: %+v", e)
		}
	}()

	injectableVal := new(injectable)
	safeInjectClientToken(injectableVal)
	if injectableVal.ClientToken == nil || len(*injectableVal.ClientToken) == 0 {
		t.Fatalf("no client token injected: %+v", injectableVal)
	}

	uninjectableVal := new(uninjectable)
	safeInjectClientToken(uninjectableVal)
}

var (
	exists = make(map[string]struct{})
)

func BenchmarkGenerateClientToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		token := randomClientToken()
		if _, conflict := exists[token]; conflict {
			b.Fatalf("conflict with generated token: %s, %d", token, i)
		}
		exists[token] = struct{}{}
	}
}
