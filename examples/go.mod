module github.com/tencentcloud/tencentcloud-sdk-go/examples

go 1.17

require (
	github.com/golang/protobuf v1.5.2
	github.com/pierrec/lz4 v2.6.1+incompatible
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecc v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hcm v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tci v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tic v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc v1.0.753
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc v1.0.753
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/frankban/quicktest v1.14.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
)

replace (
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame v1.0.753 => ../tencentcloud/ame
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.753 => ../tencentcloud/common
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.753 => ../tencentcloud/cvm
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecc v1.0.753 => ../tencentcloud/ecc
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess v1.0.753 => ../tencentcloud/ess
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic v1.0.753 => ../tencentcloud/essbasic
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hcm v1.0.753 => ../tencentcloud/hcm
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor v1.0.753 => ../tencentcloud/monitor
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe v1.0.753 => ../tencentcloud/soe
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm v1.0.753 => ../tencentcloud/ssm
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tci v1.0.753 => ../tencentcloud/tci
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tic v1.0.753 => ../tencentcloud/tic
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc v1.0.753 => ../tencentcloud/trtc
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc v1.0.753 => ../tencentcloud/vpc
)
