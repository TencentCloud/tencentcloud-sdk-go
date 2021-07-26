package common

import (
	"os"
	"reflect"
	"testing"
)

func TestEnvProvider_GetCredential(t *testing.T) {
	type fields struct {
		secretIdENV  string
		secretKeyENV string
	}
	tests := []struct {
		name    string
		fields  fields
		want    CredentialIface
		wantErr bool
	}{
		{"valid env", fields{
			secretIdENV:  "TENCENTCLOUD_SECRET_ID_test",
			secretKeyENV: "TENCENTCLOUD_SECRET_KEY_test",
		},
			&Credential{
				SecretId:  "xxxxxx",
				SecretKey: "xxxxxx",
				Token:     "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.fields.secretIdENV, tt.want.GetSecretId())
			os.Setenv(tt.fields.secretKeyENV, tt.want.GetSecretKey())
			p := &EnvProvider{
				secretIdENV:  tt.fields.secretIdENV,
				secretKeyENV: tt.fields.secretKeyENV,
			}
			got, err := p.GetCredential()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredential() got = %v, want %v", got, tt.want)
			}
		})
	}
}
