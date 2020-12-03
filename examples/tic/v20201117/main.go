package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tic/v20201117"
)

var client *tic.Client

// checkError 异常处理
func checkError(err error) {
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Panicf("An API error has returned: %+v", err)
	}
	// unexpected errors
	if err != nil {
		log.Panicf("%+v", err)
	}
}

// dumpResponse 打印返回的json字符串
func dumpResponse(msg string, resp interface{}) {
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Panicf("json marshal: %+v", err)
	}
	log.Printf("%s: %s", msg, b)
}

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	templateURL := os.Getenv("TIC_TEMPLATE_URL")
	if templateURL == "" {
		log.Fatalf("please provide template url as environment variable TIC_TEMPLATE_URL, a zip file uploaded to COS")
	}
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// 实例化要请求产品的client对象
	ticClient, err := tic.NewClient(credential, "" /* 不需要考虑地域，在CreateStack接口里指定 */, cpf)
	if err != nil {
		log.Panicf("tic NewClient: %+v", err)
	}

	client = ticClient

	// 创建资源栈，如果已有资源栈，可以调用DescribeStacks获得
	createStackReq := tic.NewCreateStackRequest()
	createStackReq.StackName = common.StringPtr("my stack name")
	createStackReq.Description = common.StringPtr("my awesome stack descripion")
	createStackReq.StackRegion = common.StringPtr("ap-guangzhou") // 指定资源栈所属的地域，这里需要换成你的
	// 这里需要先将terraform的文件打包成zip上传到腾讯云cos
	createStackReq.TemplateUrl = common.StringPtr(templateURL)
	createStackResp, err := client.CreateStack(createStackReq)
	checkError(err)
	dumpResponse("CreateStack", createStackResp)

	// 此时会得到一个资源栈ID和一个版本ID，可以调用CreateStackVersion给资源栈增加一个版本
	// 或者UpdateStackVersion更新版本的内容，这里略过，直接使用创建资源栈时使用的版本

	// 执行plan事件
	planStackReq := tic.NewPlanStackRequest()
	planStackReq.StackId = createStackResp.Response.StackId
	planStackReq.VersionId = createStackResp.Response.VersionId
	planStackResp, err := client.PlanStack(planStackReq)
	checkError(err)
	dumpResponse("PlanStack", planStackResp)
	waitEventFinished(*planStackResp.Response.EventId)

	// 执行apply事件
	applyStackReq := tic.NewApplyStackRequest()
	applyStackReq.StackId = createStackResp.Response.StackId
	applyStackReq.VersionId = createStackResp.Response.VersionId
	applyStackResp, err := client.ApplyStack(applyStackReq)
	checkError(err)
	dumpResponse("ApplyStack", applyStackResp)
	waitEventFinished(*planStackResp.Response.EventId)

	// 执行destroy事件，销毁创建的资源
	destroyStackReq := tic.NewDestroyStackRequest()
	destroyStackReq.StackId = createStackResp.Response.StackId
	destroyStackReq.VersionId = createStackResp.Response.VersionId
	destroyStackResp, err := client.DestroyStack(destroyStackReq)
	checkError(err)
	dumpResponse("DestroyStack", destroyStackResp)
	waitEventFinished(*planStackResp.Response.EventId)
}

// 查看并等待（如果事件仍在执行中）事件执行的日志
func waitEventFinished(eventID string) {
	describeStackEventReq := tic.NewDescribeStackEventRequest()
	describeStackEventReq.EventId = &eventID
	describeStackEventResp, err := client.DescribeStackEvent(describeStackEventReq)
	if err, ok := err.(*errors.TencentCloudSDKError); ok {
		// 版本可能处于plan，apply进行中状态，需要等待状态变成成功或失败
		if err.Code == "UnsupportedOperation.ForbidOp" {
			// 可能事件仍在执行中，这里简单的10秒钟后再查询一次，实际代码需要更完善的重试逻辑
			log.Printf("waiting event finished...")
			time.Sleep(time.Second * 10)
			waitEventFinished(eventID)
			return
		}
		log.Panicf("An API error has returned: %+v", err)
	}
	// unexpected errors
	if err != nil {
		log.Panicf("%+v", err)
	}
	dumpResponse("DescribeStackEvent", describeStackEventResp)

	if *describeStackEventResp.Response.Status == "running" { // 等待running状态结束
		log.Printf("waiting event finished...")
		time.Sleep(time.Second * 10)
		waitEventFinished(eventID)
		return
	}
}
