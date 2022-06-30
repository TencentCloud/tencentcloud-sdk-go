package v20201111

import (
	"encoding/base64"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20201222"
	"os"
	"testing"
	"time"
)

const (
	EndPoint = "ess.tencentcloudapi.com"

	FileEndPoint = "file.ess.tencent.cn"

	// Ak 腾讯云ak/sk 腾讯云后台CAM
	Ak = "***********************"

	Sk = "***********************"

	// 管理员用户id或者员工用户id
	OperatorId = "***********************"
)

func TestCreateFlow(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewCreateFlowRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")

	// 企业方 静默签署时type为3/非静默签署type为0
	approver1 := ess.FlowCreateApprover{
		ApproverType:     common.Int64Ptr(3),
		OrganizationName: common.StringPtr("***********************"),
		ApproverName:     common.StringPtr("***********************"),
		ApproverMobile:   common.StringPtr("***********************"),
		Required:         common.BoolPtr(true), // 请务必设置
	}
	// 客户个人
	approver2 := ess.FlowCreateApprover{
		ApproverType:   common.Int64Ptr(1),
		ApproverName:   common.StringPtr("***********************"),
		ApproverMobile: common.StringPtr("***********************"),
		Required:       common.BoolPtr(true), // 请务必设置
	}

	// 注：当进行B2B签署时，允许指向未注册的企业，此时签署人可以查看合同并按照指引注册企业
	req.Approvers = []*ess.FlowCreateApprover{&approver1, &approver2}

	// 请设置合理的时间，否则容易造成合同过期
	req.DeadLine = common.Int64Ptr(time.Now().Add(7 * 24 * time.Hour).Unix())
	req.FlowName = common.StringPtr("***********************")

	res, err := client.CreateFlow(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestCreateDocument(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewCreateDocumentRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")

	// 由CreateFlow返回
	req.FlowId = common.StringPtr("***********************")
	// 传入自定义文件名即可
	req.FileNames = []*string{common.StringPtr("************")}
	// 后台配置后查询获取
	req.TemplateId = common.StringPtr("***********************")

	// 控件信息
	req.FormFields = []*ess.FormField{
		{
			// 在模板配置拖入控件的界面可以查询到（ComponentName或者ComponentId选一填写，建议填写ComponentName）
			ComponentName:  common.StringPtr("***********************"),
			ComponentValue: common.StringPtr("***********************"),
		},
	}

	res, err := client.CreateDocument(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestStartFlow(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewStartFlowRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId:   common.StringPtr(OperatorId),
		ClientIp: common.StringPtr("***********************"),
	}
	req.BaseRequest.SetHttpMethod("POST")

	// 由上一步返回
	req.FlowId = common.StringPtr("***********************")

	res, err := client.StartFlow(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestDescribeFileUrls(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewDescribeFileUrlsRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")

	req.BusinessType = common.StringPtr("FLOW")
	req.BusinessIds = []*string{common.StringPtr("***********************")}

	res, err := client.DescribeFileUrls(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestUploadFiles(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = FileEndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := essbasic.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := essbasic.NewUploadFilesRequest()
	// 公共参数
	req.Caller = &essbasic.Caller{
		OperatorId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")
	req.BusinessType = common.StringPtr("FLOW")

	// 读取文件
	data, err := os.ReadFile("/电子签项目/aaa.pdf")
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}

	fileBody := base64.StdEncoding.EncodeToString(data)

	req.FileInfos = []*essbasic.UploadFile{
		{
			FileName: common.StringPtr("test.pdf"),
			FileBody: common.StringPtr(fileBody),
		},
	}

	res, err := client.UploadFiles(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestCreateFlowByFiles(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewCreateFlowByFilesRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")

	// 签署pdf文件的资源编号列表，目前只支持单文件签署
	req.FileIds = []*string{common.StringPtr("***********************")}

	req.Approvers = []*ess.ApproverInfo{
		// 企业方 静默签署时type为3/非静默签署type为0
		{
			ApproverType:     common.Int64Ptr(3),
			ApproverName:     common.StringPtr("***********************"),
			ApproverMobile:   common.StringPtr("***********************"),
			OrganizationName: common.StringPtr("***********************"),
			SignComponents: []*ess.Component{
				{
					ComponentValue:  common.StringPtr("***********************"),
					ComponentPosY:   common.Float64Ptr(497.671875),
					ComponentWidth:  common.Float64Ptr(74),
					FileIndex:       common.Int64Ptr(0),
					ComponentType:   common.StringPtr("SIGN_SEAL"),
					ComponentPage:   common.Int64Ptr(1),
					ComponentPosX:   common.Float64Ptr(417.15625),
					ComponentHeight: common.Float64Ptr(70),
				},
			},
		},
		// 客户个人
		{
			ApproverType:   common.Int64Ptr(1),
			ApproverName:   common.StringPtr("***********************"),
			ApproverMobile: common.StringPtr("***********************"),
			SignComponents: []*ess.Component{
				{
					ComponentPosY:   common.Float64Ptr(472.78125),
					ComponentWidth:  common.Float64Ptr(112),
					FileIndex:       common.Int64Ptr(0),
					ComponentType:   common.StringPtr("SIGN_SIGNATURE"),
					ComponentPage:   common.Int64Ptr(1),
					ComponentPosX:   common.Float64Ptr(146.15625),
					ComponentHeight: common.Float64Ptr(40),
				},
			},
		},
	}

	req.FlowName = common.StringPtr("***********************")
	req.Deadline = common.Int64Ptr(time.Now().Add(7 * 24 * time.Hour).Unix())

	res, err := client.CreateFlowByFiles(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestDescribeFlowBriefs(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewDescribeFlowBriefsRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")

	req.FlowIds = []*string{common.StringPtr("***********************")}

	res, err := client.DescribeFlowBriefs(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestCreateSchemeUrl(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewCreateSchemeUrlRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")

	req.FlowId = common.StringPtr("***********************")
	req.PathType = common.Uint64Ptr(1)

	res, err := client.CreateSchemeUrl(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestDescribeFlowTemplates(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewDescribeFlowTemplatesRequest()
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.BaseRequest.SetHttpMethod("POST")

	res, err := client.DescribeFlowTemplates(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}

func TestCancelFlow(t *testing.T) {
	pf := profile.NewClientProfile()
	pf.HttpProfile.Endpoint = EndPoint
	pf.HttpProfile.Scheme = "HTTPS"
	client, _ := ess.NewClient(common.NewCredential(Ak, Sk), "ap-guangzhou", pf)

	req := ess.NewCancelFlowRequest()
	req.BaseRequest.SetHttpMethod("POST")
	// 公共参数
	req.Operator = &ess.UserInfo{
		UserId: common.StringPtr(OperatorId),
	}
	req.FlowId = common.StringPtr("***********************")
	req.CancelMessage = common.StringPtr("***********************")

	res, err := client.CancelFlow(req)
	if err != nil {
		fmt.Printf("error : %v\n", err)
		return
	}
	fmt.Println(res.ToJsonString())
}
