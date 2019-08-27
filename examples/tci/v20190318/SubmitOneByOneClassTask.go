package main

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tci "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tci/v20190318"
)

func main() {
	credential := common.NewCredential(
		// os.Getenv("TENCENTCLOUD_SECRET_ID"),
		// os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"", "",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 30
	cpf.HttpProfile.Endpoint = "tci.tencentcloudapi.com"
	client, _ := tci.NewClient(credential, "ap-guangzhou", cpf)
	req := tci.NewSubmitOneByOneClassTaskRequest()
	req.FileContent = common.StringPtr("https://edu-test-1253131631.cos.ap-guangzhou.myqcloud.com/aieduautotest/autotest_vedio.mp4")
	req.FileType = common.StringPtr("vod_url")
	req.Lang = common.Int64Ptr(0)
	req.LibrarySet = common.StringPtrs([]string{"library_15603955264181591716"})
	req.VocabLibNameList = common.StringPtrs([]string{"testlib2"})
	req.VoiceEncodeType = common.Int64Ptr(1)
	req.VoiceFileType = common.Int64Ptr(10)
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SubmitOneByOneClassTask(req)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}
