module github.com/tencentcloud/tencentcloud-sdk-go/testing

go 1.14

require (
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v0.0.0
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v0.0.0
)

replace (
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common => ../tencentcloud/common
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm => ../tencentcloud/cvm
)
