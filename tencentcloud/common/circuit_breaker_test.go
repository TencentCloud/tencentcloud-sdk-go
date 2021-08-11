package common

import "testing"

func Test_checkDomain(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid endpoint", args{endpoint: "cvm.tencentcloudapi.com"}, true},
		{"valid endpoint", args{endpoint: "cvm.ap-beijing.tencentcloudapi.com"}, true},
		{"invalid endpoint", args{endpoint: "cvm.tencentcloud.com"}, false},
		{"invalid endpoint", args{endpoint: "cvm.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEndpoint(tt.args.endpoint); got != tt.want {
				t.Errorf("checkEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renewUrl(t *testing.T) {
	type args struct {
		oldDomain string
		region    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success3", args{
			oldDomain: "cvm.tencentcloudapi.com",
			region:    "ap-beijing",
		}, "cvm.ap-beijing.tencentcloudapi.com"},
		{"success4", args{
			oldDomain: "cvm.ap-beijing.tencentcloudapi.com",
			region:    "ap-shanghai",
		}, "cvm.ap-shanghai.tencentcloudapi.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renewUrl(tt.args.oldDomain, tt.args.region); got != tt.want {
				t.Errorf("renewUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
