module github.com/tencentcloud/tencentcloud-sdk-go/testing

go 1.14

require (
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v0.0.0
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v0.0.0
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf v0.0.0-00010101000000-000000000000
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v0.0.0-00010101000000-000000000000
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc v0.0.0-00010101000000-000000000000
)

replace (
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common => ../tencentcloud/common
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm => ../tencentcloud/cvm
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf => ../tencentcloud/scf
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms => ../tencentcloud/sms
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc => ../tencentcloud/trtc
)
