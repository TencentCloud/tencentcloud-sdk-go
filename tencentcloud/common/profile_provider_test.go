package common

import (
	"testing"
)

func TestProfileProvider_GetCredential(t *testing.T) {
	tests := []struct {
		name    string
		want    CredentialIface
		wantErr bool
	}{
		{"valid profile", &Credential{
			SecretId:  "xxxxx",
			SecretKey: "xxxxx",
			Token:     "",
		},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProfileProvider{}
			_, err := p.GetCredential()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
