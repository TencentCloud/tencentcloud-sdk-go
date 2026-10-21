package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

// tccliSetEnv sets an env var for the duration of the test and restores it afterwards.
func tccliSetEnv(t *testing.T, key, value string) {
	t.Helper()
	old, existed := os.LookupEnv(key)
	if err := os.Setenv(key, value); err != nil {
		t.Fatalf("setenv %s: %v", key, err)
	}
	t.Cleanup(func() {
		if existed {
			os.Setenv(key, old)
		} else {
			os.Unsetenv(key)
		}
	})
}

// tccliUnsetEnv unsets an env var for the duration of the test and restores it afterwards.
func tccliUnsetEnv(t *testing.T, key string) {
	t.Helper()
	old, existed := os.LookupEnv(key)
	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("unsetenv %s: %v", key, err)
	}
	t.Cleanup(func() {
		if existed {
			os.Setenv(key, old)
		}
	})
}

// tccliSetHome points the home directory (HOME / USERPROFILE) to dir.
func tccliSetHome(t *testing.T, dir string) {
	t.Helper()
	tccliSetEnv(t, "HOME", dir)
	tccliSetEnv(t, "USERPROFILE", dir)
}

// tccliTempHome creates a temporary home directory and points HOME to it.
func tccliTempHome(t *testing.T) string {
	t.Helper()
	dir, err := ioutil.TempDir("", "tccli-home-")
	if err != nil {
		t.Fatalf("create temp dir: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	tccliSetHome(t, dir)
	return dir
}

// tccliWriteCredential writes content into {home}/.tccli/{profile}.credential.
func tccliWriteCredential(t *testing.T, home, profile, content string) string {
	t.Helper()
	dir := filepath.Join(home, ".tccli")
	if err := os.MkdirAll(dir, 0700); err != nil {
		t.Fatalf("mkdir %s: %v", dir, err)
	}
	path := filepath.Join(dir, profile+".credential")
	if err := ioutil.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
	return path
}

func assertCredentialSDKError(t *testing.T, err error, msgContains string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	sdkErr, ok := err.(*tcerr.TencentCloudSDKError)
	if !ok {
		t.Fatalf("expected *TencentCloudSDKError, got %T: %v", err, err)
	}
	if sdkErr.GetCode() != creErr {
		t.Errorf("error code = %q, want %q", sdkErr.GetCode(), creErr)
	}
	if msgContains != "" && !strings.Contains(sdkErr.GetMessage(), msgContains) {
		t.Errorf("error message = %q, want contains %q", sdkErr.GetMessage(), msgContains)
	}
}

func TestDefaultTccliProfileProvider(t *testing.T) {
	t.Run("env unset falls back to default", func(t *testing.T) {
		tccliUnsetEnv(t, EnvTCCLIProfile)
		p := DefaultTccliProfileProvider()
		if p.profile != "default" {
			t.Errorf("profile = %q, want %q", p.profile, "default")
		}
	})

	t.Run("env set uses env value", func(t *testing.T) {
		tccliSetEnv(t, EnvTCCLIProfile, "dev")
		p := DefaultTccliProfileProvider()
		if p.profile != "dev" {
			t.Errorf("profile = %q, want %q", p.profile, "dev")
		}
	})

	t.Run("env set to empty keeps empty", func(t *testing.T) {
		tccliSetEnv(t, EnvTCCLIProfile, "")
		p := DefaultTccliProfileProvider()
		if p.profile != "" {
			t.Errorf("profile = %q, want empty", p.profile)
		}
	})
}

func TestNewTccliProfileProvider(t *testing.T) {
	p := NewTccliProfileProvider("prod")
	if p == nil {
		t.Fatal("NewTccliProfileProvider returned nil")
	}
	if p.profile != "prod" {
		t.Errorf("profile = %q, want %q", p.profile, "prod")
	}
}

func TestTccliProfileProvider_ImplementsProvider(t *testing.T) {
	var _ Provider = NewTccliProfileProvider("default")
	var _ Provider = DefaultTccliProfileProvider()
}

func TestTccliProfileProvider_getCredentialFilePath(t *testing.T) {
	t.Run("empty home returns empty path", func(t *testing.T) {
		tccliSetHome(t, "")
		p := NewTccliProfileProvider("default")
		if got := p.getCredentialFilePath(); got != "" {
			t.Errorf("getCredentialFilePath() = %q, want empty", got)
		}
	})

	t.Run("home set returns joined path", func(t *testing.T) {
		home := filepath.Join(string(filepath.Separator)+"tmp", "fake-home")
		tccliSetHome(t, home)
		p := NewTccliProfileProvider("myprofile")
		want := filepath.Join(home, ".tccli", "myprofile.credential")
		if got := p.getCredentialFilePath(); got != want {
			t.Errorf("getCredentialFilePath() = %q, want %q", got, want)
		}
	})
}

func TestTccliProfileProvider_checkCredentialFile(t *testing.T) {
	t.Run("empty home", func(t *testing.T) {
		tccliSetHome(t, "")
		path, err := NewTccliProfileProvider("default").checkCredentialFile()
		if err != nil || path != "" {
			t.Errorf("checkCredentialFile() = (%q, %v), want (\"\", nil)", path, err)
		}
	})

	t.Run("file not exist", func(t *testing.T) {
		tccliTempHome(t)
		path, err := NewTccliProfileProvider("default").checkCredentialFile()
		if err != nil || path != "" {
			t.Errorf("checkCredentialFile() = (%q, %v), want (\"\", nil)", path, err)
		}
	})

	t.Run("file exists", func(t *testing.T) {
		home := tccliTempHome(t)
		want := tccliWriteCredential(t, home, "default", `{}`)
		path, err := NewTccliProfileProvider("default").checkCredentialFile()
		if err != nil {
			t.Fatalf("checkCredentialFile() unexpected error: %v", err)
		}
		if path != want {
			t.Errorf("checkCredentialFile() path = %q, want %q", path, want)
		}
	})

	t.Run("stat error other than not exist", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("ENOTDIR semantics differ on windows")
		}
		// Make HOME a regular file so that stat on HOME/.tccli/x.credential fails with ENOTDIR.
		dir := tccliTempHome(t)
		homeFile := filepath.Join(dir, "not-a-dir")
		if err := ioutil.WriteFile(homeFile, []byte("x"), 0600); err != nil {
			t.Fatalf("write file: %v", err)
		}
		tccliSetHome(t, homeFile)

		path, err := NewTccliProfileProvider("default").checkCredentialFile()
		if err == nil {
			t.Fatalf("checkCredentialFile() expected error, got path %q", path)
		}
		if path != "" {
			t.Errorf("checkCredentialFile() path = %q, want empty", path)
		}
	})
}

func TestTccliProfileProvider_GetCredential(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		content string
		want    CredentialIface
	}{
		{
			name:    "with token",
			profile: "default",
			content: `{"secretId":"AKIDxxx","secretKey":"SKxxx","token":"TOKENxxx"}`,
			want:    &Credential{SecretId: "AKIDxxx", SecretKey: "SKxxx", Token: "TOKENxxx"},
		},
		{
			name:    "without token",
			profile: "default",
			content: `{"secretId":"AKIDxxx","secretKey":"SKxxx"}`,
			want:    &Credential{SecretId: "AKIDxxx", SecretKey: "SKxxx"},
		},
		{
			name:    "extra fields ignored",
			profile: "default",
			content: `{"secretId":"AKIDxxx","secretKey":"SKxxx","region":"ap-guangzhou","type":"secret"}`,
			want:    &Credential{SecretId: "AKIDxxx", SecretKey: "SKxxx"},
		},
		{
			name:    "custom profile",
			profile: "prod",
			content: `{"secretId":"AKIDprod","secretKey":"SKprod"}`,
			want:    &Credential{SecretId: "AKIDprod", SecretKey: "SKprod"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home := tccliTempHome(t)
			tccliWriteCredential(t, home, tt.profile, tt.content)

			got, err := NewTccliProfileProvider(tt.profile).GetCredential()
			if err != nil {
				t.Fatalf("GetCredential() unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredential() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestTccliProfileProvider_GetCredential_DefaultProviderUsesEnvProfile(t *testing.T) {
	home := tccliTempHome(t)
	tccliWriteCredential(t, home, "default", `{"secretId":"AKIDdefault","secretKey":"SKdefault"}`)
	tccliWriteCredential(t, home, "test", `{"secretId":"AKIDtest","secretKey":"SKtest"}`)
	tccliSetEnv(t, EnvTCCLIProfile, "test")

	got, err := DefaultTccliProfileProvider().GetCredential()
	if err != nil {
		t.Fatalf("GetCredential() unexpected error: %v", err)
	}
	want := &Credential{SecretId: "AKIDtest", SecretKey: "SKtest"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetCredential() got = %+v, want %+v", got, want)
	}
}

func TestTccliProfileProvider_GetCredential_Errors(t *testing.T) {
	t.Run("empty home returns fileDoseNotExist", func(t *testing.T) {
		tccliSetHome(t, "")
		got, err := NewTccliProfileProvider("default").GetCredential()
		if got != nil {
			t.Errorf("GetCredential() got = %v, want nil", got)
		}
		if err != fileDoseNotExist {
			t.Errorf("GetCredential() err = %v, want fileDoseNotExist", err)
		}
	})

	t.Run("file not exist returns fileDoseNotExist", func(t *testing.T) {
		home := tccliTempHome(t)
		// a credential for another profile must not be picked up
		tccliWriteCredential(t, home, "other", `{"secretId":"a","secretKey":"b"}`)
		got, err := NewTccliProfileProvider("default").GetCredential()
		if got != nil {
			t.Errorf("GetCredential() got = %v, want nil", got)
		}
		if err != fileDoseNotExist {
			t.Errorf("GetCredential() err = %v, want fileDoseNotExist", err)
		}
	})

	t.Run("stat error wrapped as sdk error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("ENOTDIR semantics differ on windows")
		}
		dir := tccliTempHome(t)
		homeFile := filepath.Join(dir, "not-a-dir")
		if err := ioutil.WriteFile(homeFile, []byte("x"), 0600); err != nil {
			t.Fatalf("write file: %v", err)
		}
		tccliSetHome(t, homeFile)

		got, err := NewTccliProfileProvider("default").GetCredential()
		if got != nil {
			t.Errorf("GetCredential() got = %v, want nil", got)
		}
		assertCredentialSDKError(t, err, "Failed to find TCCLI profile file")
	})

	t.Run("path is directory returns read error", func(t *testing.T) {
		home := tccliTempHome(t)
		if err := os.MkdirAll(filepath.Join(home, ".tccli", "default.credential"), 0700); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		got, err := NewTccliProfileProvider("default").GetCredential()
		if got != nil {
			t.Errorf("GetCredential() got = %v, want nil", got)
		}
		if err == nil {
			t.Fatal("GetCredential() expected error, got nil")
		}
	})

	invalidContentCases := []struct {
		name        string
		content     string
		wantSDKErr  bool
		msgContains string
	}{
		{name: "invalid json", content: `{"secretId":`},
		{name: "empty file", content: ``},
		{name: "json array", content: `[]`},
		{name: "wrong field type", content: `{"secretId":123,"secretKey":"SK"}`},
		{name: "empty object", content: `{}`, wantSDKErr: true, msgContains: "secretId"},
		{name: "missing secretId", content: `{"secretKey":"SKxxx"}`, wantSDKErr: true, msgContains: "secretId"},
		{name: "missing secretKey", content: `{"secretId":"AKIDxxx"}`, wantSDKErr: true, msgContains: "secretKey"},
		{name: "empty secretId", content: `{"secretId":"","secretKey":"SKxxx"}`, wantSDKErr: true, msgContains: "secretId"},
		{name: "empty secretKey", content: `{"secretId":"AKIDxxx","secretKey":""}`, wantSDKErr: true, msgContains: "secretKey"},
		{name: "only token", content: `{"token":"TOKENxxx"}`, wantSDKErr: true, msgContains: "secretId"},
	}
	for _, tc := range invalidContentCases {
		t.Run(tc.name, func(t *testing.T) {
			home := tccliTempHome(t)
			tccliWriteCredential(t, home, "default", tc.content)

			got, err := NewTccliProfileProvider("default").GetCredential()
			if got != nil {
				t.Errorf("GetCredential() got = %v, want nil", got)
			}
			if err == nil {
				t.Fatal("GetCredential() expected error, got nil")
			}
			if tc.wantSDKErr {
				assertCredentialSDKError(t, err, tc.msgContains)
			}
		})
	}
}
