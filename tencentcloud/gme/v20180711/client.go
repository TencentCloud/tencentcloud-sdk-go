// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180711

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-11"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewControlAIConversationRequest() (request *ControlAIConversationRequest) {
    request = &ControlAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ControlAIConversation")
    
    
    return
}

func NewControlAIConversationResponse() (response *ControlAIConversationResponse) {
    response = &ControlAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlAIConversation
// 提供服务端控制机器人的功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) ControlAIConversation(request *ControlAIConversationRequest) (response *ControlAIConversationResponse, err error) {
    return c.ControlAIConversationWithContext(context.Background(), request)
}

// ControlAIConversation
// 提供服务端控制机器人的功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) ControlAIConversationWithContext(ctx context.Context, request *ControlAIConversationRequest) (response *ControlAIConversationResponse, err error) {
    if request == nil {
        request = NewControlAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "ControlAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgeDetectTaskRequest() (request *CreateAgeDetectTaskRequest) {
    request = &CreateAgeDetectTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "CreateAgeDetectTask")
    
    
    return
}

func NewCreateAgeDetectTaskResponse() (response *CreateAgeDetectTaskResponse) {
    response = &CreateAgeDetectTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgeDetectTask
// 目前该功能底层能力已不具备，不对外提供，目前需要下线，走预下线流程。
//
// 
//
// 用于创建年龄语音识别任务的接口，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// </br>
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音文件进行检测，判断是否为未成年人。</li>
//
// <li>支持批量提交检测子任务。检测子任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：10 M</li>
//
// <li>音频文件时长限制：3分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgeDetectTask(request *CreateAgeDetectTaskRequest) (response *CreateAgeDetectTaskResponse, err error) {
    return c.CreateAgeDetectTaskWithContext(context.Background(), request)
}

// CreateAgeDetectTask
// 目前该功能底层能力已不具备，不对外提供，目前需要下线，走预下线流程。
//
// 
//
// 用于创建年龄语音识别任务的接口，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// </br>
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音文件进行检测，判断是否为未成年人。</li>
//
// <li>支持批量提交检测子任务。检测子任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：10 M</li>
//
// <li>音频文件时长限制：3分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgeDetectTaskWithContext(ctx context.Context, request *CreateAgeDetectTaskRequest) (response *CreateAgeDetectTaskResponse, err error) {
    if request == nil {
        request = NewCreateAgeDetectTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "CreateAgeDetectTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgeDetectTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgeDetectTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApp
// 本接口(CreateApp)用于创建一个GME应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// 本接口(CreateApp)用于创建一个GME应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "CreateApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomizationRequest() (request *CreateCustomizationRequest) {
    request = &CreateCustomizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "CreateCustomization")
    
    
    return
}

func NewCreateCustomizationResponse() (response *CreateCustomizationResponse) {
    response = &CreateCustomizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomization
// 用户使用该接口可以创建语音消息转文本热句模型，以供识别调用
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) CreateCustomization(request *CreateCustomizationRequest) (response *CreateCustomizationResponse, err error) {
    return c.CreateCustomizationWithContext(context.Background(), request)
}

// CreateCustomization
// 用户使用该接口可以创建语音消息转文本热句模型，以供识别调用
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) CreateCustomizationWithContext(ctx context.Context, request *CreateCustomizationRequest) (response *CreateCustomizationResponse, err error) {
    if request == nil {
        request = NewCreateCustomizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "CreateCustomization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScanUserRequest() (request *CreateScanUserRequest) {
    request = &CreateScanUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "CreateScanUser")
    
    
    return
}

func NewCreateScanUserResponse() (response *CreateScanUserResponse) {
    response = &CreateScanUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScanUser
// 新增自定义送检用户。**接口使用前提**：目前 CreateScanUser 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateScanUser(request *CreateScanUserRequest) (response *CreateScanUserResponse, err error) {
    return c.CreateScanUserWithContext(context.Background(), request)
}

// CreateScanUser
// 新增自定义送检用户。**接口使用前提**：目前 CreateScanUser 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateScanUserWithContext(ctx context.Context, request *CreateScanUserRequest) (response *CreateScanUserResponse, err error) {
    if request == nil {
        request = NewCreateScanUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "CreateScanUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScanUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScanUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomizationRequest() (request *DeleteCustomizationRequest) {
    request = &DeleteCustomizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DeleteCustomization")
    
    
    return
}

func NewDeleteCustomizationResponse() (response *DeleteCustomizationResponse) {
    response = &DeleteCustomizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomization
// 用户通过该接口可以删除语音消息转文本热句模型
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DeleteCustomization(request *DeleteCustomizationRequest) (response *DeleteCustomizationResponse, err error) {
    return c.DeleteCustomizationWithContext(context.Background(), request)
}

// DeleteCustomization
// 用户通过该接口可以删除语音消息转文本热句模型
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DeleteCustomizationWithContext(ctx context.Context, request *DeleteCustomizationRequest) (response *DeleteCustomizationResponse, err error) {
    if request == nil {
        request = NewDeleteCustomizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DeleteCustomization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoomMemberRequest() (request *DeleteRoomMemberRequest) {
    request = &DeleteRoomMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DeleteRoomMember")
    
    
    return
}

func NewDeleteRoomMemberResponse() (response *DeleteRoomMemberResponse) {
    response = &DeleteRoomMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoomMember
// 本接口(DeleteRoomMember)用户删除房间或者剔除房间内用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDELETETYPE = "InvalidParameterValue.InvalidDeleteType"
//  INVALIDPARAMETERVALUE_INVALIDSTRUIN = "InvalidParameterValue.InvalidStrUin"
//  INVALIDPARAMETERVALUE_INVALIDUINORSTRUIN = "InvalidParameterValue.InvalidUinOrStrUin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_LIMITDELETEEXCEEDED = "UnsupportedOperation.LimitDeleteExceeded"
func (c *Client) DeleteRoomMember(request *DeleteRoomMemberRequest) (response *DeleteRoomMemberResponse, err error) {
    return c.DeleteRoomMemberWithContext(context.Background(), request)
}

// DeleteRoomMember
// 本接口(DeleteRoomMember)用户删除房间或者剔除房间内用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDELETETYPE = "InvalidParameterValue.InvalidDeleteType"
//  INVALIDPARAMETERVALUE_INVALIDSTRUIN = "InvalidParameterValue.InvalidStrUin"
//  INVALIDPARAMETERVALUE_INVALIDUINORSTRUIN = "InvalidParameterValue.InvalidUinOrStrUin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_LIMITDELETEEXCEEDED = "UnsupportedOperation.LimitDeleteExceeded"
func (c *Client) DeleteRoomMemberWithContext(ctx context.Context, request *DeleteRoomMemberRequest) (response *DeleteRoomMemberResponse, err error) {
    if request == nil {
        request = NewDeleteRoomMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DeleteRoomMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoomMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoomMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScanUserRequest() (request *DeleteScanUserRequest) {
    request = &DeleteScanUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DeleteScanUser")
    
    
    return
}

func NewDeleteScanUserResponse() (response *DeleteScanUserResponse) {
    response = &DeleteScanUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScanUser
// 删除自定义送检用户。**接口使用前提**：目前 DeleteScanUser 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDELETETYPE = "InvalidParameterValue.InvalidDeleteType"
//  INVALIDPARAMETERVALUE_INVALIDSTRUIN = "InvalidParameterValue.InvalidStrUin"
//  INVALIDPARAMETERVALUE_INVALIDUINORSTRUIN = "InvalidParameterValue.InvalidUinOrStrUin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_LIMITDELETEEXCEEDED = "UnsupportedOperation.LimitDeleteExceeded"
func (c *Client) DeleteScanUser(request *DeleteScanUserRequest) (response *DeleteScanUserResponse, err error) {
    return c.DeleteScanUserWithContext(context.Background(), request)
}

// DeleteScanUser
// 删除自定义送检用户。**接口使用前提**：目前 DeleteScanUser 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDELETETYPE = "InvalidParameterValue.InvalidDeleteType"
//  INVALIDPARAMETERVALUE_INVALIDSTRUIN = "InvalidParameterValue.InvalidStrUin"
//  INVALIDPARAMETERVALUE_INVALIDUINORSTRUIN = "InvalidParameterValue.InvalidUinOrStrUin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_LIMITDELETEEXCEEDED = "UnsupportedOperation.LimitDeleteExceeded"
func (c *Client) DeleteScanUserWithContext(ctx context.Context, request *DeleteScanUserRequest) (response *DeleteScanUserResponse, err error) {
    if request == nil {
        request = NewDeleteScanUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DeleteScanUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScanUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScanUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVoicePrintRequest() (request *DeleteVoicePrintRequest) {
    request = &DeleteVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DeleteVoicePrint")
    
    
    return
}

func NewDeleteVoicePrintResponse() (response *DeleteVoicePrintResponse) {
    response = &DeleteVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVoicePrint
// 传入声纹ID，删除之前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
func (c *Client) DeleteVoicePrint(request *DeleteVoicePrintRequest) (response *DeleteVoicePrintResponse, err error) {
    return c.DeleteVoicePrintWithContext(context.Background(), request)
}

// DeleteVoicePrint
// 传入声纹ID，删除之前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
func (c *Client) DeleteVoicePrintWithContext(ctx context.Context, request *DeleteVoicePrintRequest) (response *DeleteVoicePrintResponse, err error) {
    if request == nil {
        request = NewDeleteVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DeleteVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVoicePrintResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIConversationRequest() (request *DescribeAIConversationRequest) {
    request = &DescribeAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeAIConversation")
    
    
    return
}

func NewDescribeAIConversationResponse() (response *DescribeAIConversationResponse) {
    response = &DescribeAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIConversation
// 查询AI对话任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAIConversation(request *DescribeAIConversationRequest) (response *DescribeAIConversationResponse, err error) {
    return c.DescribeAIConversationWithContext(context.Background(), request)
}

// DescribeAIConversation
// 查询AI对话任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAIConversationWithContext(ctx context.Context, request *DescribeAIConversationRequest) (response *DescribeAIConversationResponse, err error) {
    if request == nil {
        request = NewDescribeAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgeDetectTaskRequest() (request *DescribeAgeDetectTaskRequest) {
    request = &DescribeAgeDetectTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeAgeDetectTask")
    
    
    return
}

func NewDescribeAgeDetectTaskResponse() (response *DescribeAgeDetectTaskResponse) {
    response = &DescribeAgeDetectTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgeDetectTask
// 目前该功能底层能力已不具备，不对外提供，目前需要下线，走预下线流程。
//
// 
//
// 查询年龄语音识别任务结果，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgeDetectTask(request *DescribeAgeDetectTaskRequest) (response *DescribeAgeDetectTaskResponse, err error) {
    return c.DescribeAgeDetectTaskWithContext(context.Background(), request)
}

// DescribeAgeDetectTask
// 目前该功能底层能力已不具备，不对外提供，目前需要下线，走预下线流程。
//
// 
//
// 查询年龄语音识别任务结果，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgeDetectTaskWithContext(ctx context.Context, request *DescribeAgeDetectTaskRequest) (response *DescribeAgeDetectTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAgeDetectTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeAgeDetectTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgeDetectTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgeDetectTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppStatisticsRequest() (request *DescribeAppStatisticsRequest) {
    request = &DescribeAppStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeAppStatistics")
    
    
    return
}

func NewDescribeAppStatisticsResponse() (response *DescribeAppStatisticsResponse) {
    response = &DescribeAppStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAppStatistics
// 本接口(DescribeAppStatistics)用于获取某个GME应用的用量数据。包括实时语音，语音消息及转文本，语音分析等。最长查询周期为最近60天。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_DATEOUTOFSIXTYDAYS = "InvalidParameter.DateOutOfSixtyDays"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatistics(request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    return c.DescribeAppStatisticsWithContext(context.Background(), request)
}

// DescribeAppStatistics
// 本接口(DescribeAppStatistics)用于获取某个GME应用的用量数据。包括实时语音，语音消息及转文本，语音分析等。最长查询周期为最近60天。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_DATEOUTOFSIXTYDAYS = "InvalidParameter.DateOutOfSixtyDays"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatisticsWithContext(ctx context.Context, request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAppStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeAppStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationDataRequest() (request *DescribeApplicationDataRequest) {
    request = &DescribeApplicationDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeApplicationData")
    
    
    return
}

func NewDescribeApplicationDataResponse() (response *DescribeApplicationDataResponse) {
    response = &DescribeApplicationDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationData
// 本接口(DescribeApplicationData)用于获取数据详情信息，最多可拉取最近90天的数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationData(request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    return c.DescribeApplicationDataWithContext(context.Background(), request)
}

// DescribeApplicationData
// 本接口(DescribeApplicationData)用于获取数据详情信息，最多可拉取最近90天的数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationDataWithContext(ctx context.Context, request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeApplicationData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationListRequest() (request *DescribeApplicationListRequest) {
    request = &DescribeApplicationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeApplicationList")
    
    
    return
}

func NewDescribeApplicationListResponse() (response *DescribeApplicationListResponse) {
    response = &DescribeApplicationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationList
// 本接口(DescribeApplicationList)用于查询自己账号下的应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationList(request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    return c.DescribeApplicationListWithContext(context.Background(), request)
}

// DescribeApplicationList
// 本接口(DescribeApplicationList)用于查询自己账号下的应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationListWithContext(ctx context.Context, request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeApplicationList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealtimeScanConfigRequest() (request *DescribeRealtimeScanConfigRequest) {
    request = &DescribeRealtimeScanConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeRealtimeScanConfig")
    
    
    return
}

func NewDescribeRealtimeScanConfigResponse() (response *DescribeRealtimeScanConfigResponse) {
    response = &DescribeRealtimeScanConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRealtimeScanConfig
// 获取用户自定义送检信息。**接口使用前提**：目前 DescribeRealtimeScanConfig 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRealtimeScanConfig(request *DescribeRealtimeScanConfigRequest) (response *DescribeRealtimeScanConfigResponse, err error) {
    return c.DescribeRealtimeScanConfigWithContext(context.Background(), request)
}

// DescribeRealtimeScanConfig
// 获取用户自定义送检信息。**接口使用前提**：目前 DescribeRealtimeScanConfig 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRealtimeScanConfigWithContext(ctx context.Context, request *DescribeRealtimeScanConfigRequest) (response *DescribeRealtimeScanConfigResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeScanConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeRealtimeScanConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRealtimeScanConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRealtimeScanConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordInfoRequest() (request *DescribeRecordInfoRequest) {
    request = &DescribeRecordInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeRecordInfo")
    
    
    return
}

func NewDescribeRecordInfoResponse() (response *DescribeRecordInfoResponse) {
    response = &DescribeRecordInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordInfo
// 查询录制任务信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeRecordInfo(request *DescribeRecordInfoRequest) (response *DescribeRecordInfoResponse, err error) {
    return c.DescribeRecordInfoWithContext(context.Background(), request)
}

// DescribeRecordInfo
// 查询录制任务信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeRecordInfoWithContext(ctx context.Context, request *DescribeRecordInfoRequest) (response *DescribeRecordInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRecordInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeRecordInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomInfoRequest() (request *DescribeRoomInfoRequest) {
    request = &DescribeRoomInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeRoomInfo")
    
    
    return
}

func NewDescribeRoomInfoResponse() (response *DescribeRoomInfoResponse) {
    response = &DescribeRoomInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoomInfo
// 获取房间内用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRoomInfo(request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    return c.DescribeRoomInfoWithContext(context.Background(), request)
}

// DescribeRoomInfo
// 获取房间内用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRoomInfoWithContext(ctx context.Context, request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeRoomInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanResultListRequest() (request *DescribeScanResultListRequest) {
    request = &DescribeScanResultListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeScanResultList")
    
    
    return
}

func NewDescribeScanResultListResponse() (response *DescribeScanResultListResponse) {
    response = &DescribeScanResultListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScanResultList
// 本接口(DescribeScanResultList)用于查询语音检测结果，查询任务列表最多支持100个。
//
// <p style="color:red">如果在提交语音检测任务时未设置 Callback 字段，则需要通过本接口获取检测结果</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResultList(request *DescribeScanResultListRequest) (response *DescribeScanResultListResponse, err error) {
    return c.DescribeScanResultListWithContext(context.Background(), request)
}

// DescribeScanResultList
// 本接口(DescribeScanResultList)用于查询语音检测结果，查询任务列表最多支持100个。
//
// <p style="color:red">如果在提交语音检测任务时未设置 Callback 字段，则需要通过本接口获取检测结果</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResultListWithContext(ctx context.Context, request *DescribeScanResultListRequest) (response *DescribeScanResultListResponse, err error) {
    if request == nil {
        request = NewDescribeScanResultListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeScanResultList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanResultList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanResultListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeTaskInfo")
    
    
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskInfo
// 查询房间录制的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    return c.DescribeTaskInfoWithContext(context.Background(), request)
}

// DescribeTaskInfo
// 查询房间录制的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) DescribeTaskInfoWithContext(ctx context.Context, request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeTaskInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInAndOutTimeRequest() (request *DescribeUserInAndOutTimeRequest) {
    request = &DescribeUserInAndOutTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeUserInAndOutTime")
    
    
    return
}

func NewDescribeUserInAndOutTimeResponse() (response *DescribeUserInAndOutTimeResponse) {
    response = &DescribeUserInAndOutTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserInAndOutTime
// 拉取用户在房间得进出时间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserInAndOutTime(request *DescribeUserInAndOutTimeRequest) (response *DescribeUserInAndOutTimeResponse, err error) {
    return c.DescribeUserInAndOutTimeWithContext(context.Background(), request)
}

// DescribeUserInAndOutTime
// 拉取用户在房间得进出时间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserInAndOutTimeWithContext(ctx context.Context, request *DescribeUserInAndOutTimeRequest) (response *DescribeUserInAndOutTimeResponse, err error) {
    if request == nil {
        request = NewDescribeUserInAndOutTimeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeUserInAndOutTime")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInAndOutTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInAndOutTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoicePrintRequest() (request *DescribeVoicePrintRequest) {
    request = &DescribeVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "DescribeVoicePrint")
    
    
    return
}

func NewDescribeVoicePrintResponse() (response *DescribeVoicePrintResponse) {
    response = &DescribeVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVoicePrint
// 查询先前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDVOICEPRINTIDLIST = "InvalidParameter.InvalidVoicePrintIdList"
func (c *Client) DescribeVoicePrint(request *DescribeVoicePrintRequest) (response *DescribeVoicePrintResponse, err error) {
    return c.DescribeVoicePrintWithContext(context.Background(), request)
}

// DescribeVoicePrint
// 查询先前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDVOICEPRINTIDLIST = "InvalidParameter.InvalidVoicePrintIdList"
func (c *Client) DescribeVoicePrintWithContext(ctx context.Context, request *DescribeVoicePrintRequest) (response *DescribeVoicePrintResponse, err error) {
    if request == nil {
        request = NewDescribeVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "DescribeVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoicePrintResponse()
    err = c.Send(request, response)
    return
}

func NewGetCustomizationListRequest() (request *GetCustomizationListRequest) {
    request = &GetCustomizationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "GetCustomizationList")
    
    
    return
}

func NewGetCustomizationListResponse() (response *GetCustomizationListResponse) {
    response = &GetCustomizationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCustomizationList
// 查询语音消息转文本热句模型列表
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) GetCustomizationList(request *GetCustomizationListRequest) (response *GetCustomizationListResponse, err error) {
    return c.GetCustomizationListWithContext(context.Background(), request)
}

// GetCustomizationList
// 查询语音消息转文本热句模型列表
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) GetCustomizationListWithContext(ctx context.Context, request *GetCustomizationListRequest) (response *GetCustomizationListResponse, err error) {
    if request == nil {
        request = NewGetCustomizationListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "GetCustomizationList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCustomizationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCustomizationListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppStatusRequest() (request *ModifyAppStatusRequest) {
    request = &ModifyAppStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ModifyAppStatus")
    
    
    return
}

func NewModifyAppStatusResponse() (response *ModifyAppStatusResponse) {
    response = &ModifyAppStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAppStatus
// 本接口(ModifyAppStatus)用于修改应用总开关状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatus(request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    return c.ModifyAppStatusWithContext(context.Background(), request)
}

// ModifyAppStatus
// 本接口(ModifyAppStatus)用于修改应用总开关状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatusWithContext(ctx context.Context, request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    if request == nil {
        request = NewModifyAppStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "ModifyAppStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAppStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomizationRequest() (request *ModifyCustomizationRequest) {
    request = &ModifyCustomizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ModifyCustomization")
    
    
    return
}

func NewModifyCustomizationResponse() (response *ModifyCustomizationResponse) {
    response = &ModifyCustomizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomization
// 用户通过该接口可以更新语音消息转文本热句模型。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyCustomization(request *ModifyCustomizationRequest) (response *ModifyCustomizationResponse, err error) {
    return c.ModifyCustomizationWithContext(context.Background(), request)
}

// ModifyCustomization
// 用户通过该接口可以更新语音消息转文本热句模型。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyCustomizationWithContext(ctx context.Context, request *ModifyCustomizationRequest) (response *ModifyCustomizationResponse, err error) {
    if request == nil {
        request = NewModifyCustomizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "ModifyCustomization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomizationStateRequest() (request *ModifyCustomizationStateRequest) {
    request = &ModifyCustomizationStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ModifyCustomizationState")
    
    
    return
}

func NewModifyCustomizationStateResponse() (response *ModifyCustomizationStateResponse) {
    response = &ModifyCustomizationStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomizationState
// 通过该接口，用户可以修改语音消息转文本热句模型状态，上下线热句模型
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyCustomizationState(request *ModifyCustomizationStateRequest) (response *ModifyCustomizationStateResponse, err error) {
    return c.ModifyCustomizationStateWithContext(context.Background(), request)
}

// ModifyCustomizationState
// 通过该接口，用户可以修改语音消息转文本热句模型状态，上下线热句模型
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_PTTSWITCHOFF = "UnsupportedOperation.PTTSwitchOff"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyCustomizationStateWithContext(ctx context.Context, request *ModifyCustomizationStateRequest) (response *ModifyCustomizationStateResponse, err error) {
    if request == nil {
        request = NewModifyCustomizationStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "ModifyCustomizationState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomizationState require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomizationStateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordInfoRequest() (request *ModifyRecordInfoRequest) {
    request = &ModifyRecordInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ModifyRecordInfo")
    
    
    return
}

func NewModifyRecordInfoResponse() (response *ModifyRecordInfoResponse) {
    response = &ModifyRecordInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRecordInfo
// 修改录制配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyRecordInfo(request *ModifyRecordInfoRequest) (response *ModifyRecordInfoResponse, err error) {
    return c.ModifyRecordInfoWithContext(context.Background(), request)
}

// ModifyRecordInfo
// 修改录制配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) ModifyRecordInfoWithContext(ctx context.Context, request *ModifyRecordInfoRequest) (response *ModifyRecordInfoResponse, err error) {
    if request == nil {
        request = NewModifyRecordInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "ModifyRecordInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRecordInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRecordInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserMicStatusRequest() (request *ModifyUserMicStatusRequest) {
    request = &ModifyUserMicStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ModifyUserMicStatus")
    
    
    return
}

func NewModifyUserMicStatusResponse() (response *ModifyUserMicStatusResponse) {
    response = &ModifyUserMicStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserMicStatus
// **接口作用**：此接口用于修改房间用户的麦克风状态，例如房间内用户麦克风为打开状态，可调用此接口将该用户麦克风进行关闭，关闭后即使该用户使用客户端接口 EnableMic 打开麦克风，依然无法与房间内成员通话，属于被禁言状态。该状态持续到此用户退房后失效，或者调用该接口重新打开此用户麦克风状态。
//
// **接口应用场景**：此接口多用于游戏业务中台或者风控后台，对一些发表不当言论的玩家进行禁言处理。
//
// **接口使用前提**：目前 ModifyUserMicStatus 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyUserMicStatus(request *ModifyUserMicStatusRequest) (response *ModifyUserMicStatusResponse, err error) {
    return c.ModifyUserMicStatusWithContext(context.Background(), request)
}

// ModifyUserMicStatus
// **接口作用**：此接口用于修改房间用户的麦克风状态，例如房间内用户麦克风为打开状态，可调用此接口将该用户麦克风进行关闭，关闭后即使该用户使用客户端接口 EnableMic 打开麦克风，依然无法与房间内成员通话，属于被禁言状态。该状态持续到此用户退房后失效，或者调用该接口重新打开此用户麦克风状态。
//
// **接口应用场景**：此接口多用于游戏业务中台或者风控后台，对一些发表不当言论的玩家进行禁言处理。
//
// **接口使用前提**：目前 ModifyUserMicStatus 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyUserMicStatusWithContext(ctx context.Context, request *ModifyUserMicStatusRequest) (response *ModifyUserMicStatusResponse, err error) {
    if request == nil {
        request = NewModifyUserMicStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "ModifyUserMicStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserMicStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserMicStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterVoicePrintRequest() (request *RegisterVoicePrintRequest) {
    request = &RegisterVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "RegisterVoicePrint")
    
    
    return
}

func NewRegisterVoicePrintResponse() (response *RegisterVoicePrintResponse) {
    response = &RegisterVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterVoicePrint
// 传入音频base64串，注册声纹信息，返回声纹ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_GENVOICEPRINTIDFAILED = "FailedOperation.GenVoicePrintIdFailed"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  FAILEDOPERATION_VOICEPRINTREGISTRATIONLIMIT = "FailedOperation.VoicePrintRegistrationLimit"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
func (c *Client) RegisterVoicePrint(request *RegisterVoicePrintRequest) (response *RegisterVoicePrintResponse, err error) {
    return c.RegisterVoicePrintWithContext(context.Background(), request)
}

// RegisterVoicePrint
// 传入音频base64串，注册声纹信息，返回声纹ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_GENVOICEPRINTIDFAILED = "FailedOperation.GenVoicePrintIdFailed"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  FAILEDOPERATION_VOICEPRINTREGISTRATIONLIMIT = "FailedOperation.VoicePrintRegistrationLimit"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
func (c *Client) RegisterVoicePrintWithContext(ctx context.Context, request *RegisterVoicePrintRequest) (response *RegisterVoicePrintResponse, err error) {
    if request == nil {
        request = NewRegisterVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "RegisterVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterVoicePrintResponse()
    err = c.Send(request, response)
    return
}

func NewScanVoiceRequest() (request *ScanVoiceRequest) {
    request = &ScanVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "ScanVoice")
    
    
    return
}

func NewScanVoiceResponse() (response *ScanVoiceResponse) {
    response = &ScanVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScanVoice
// 本接口(ScanVoice)用于提交语音检测任务，检测任务列表最多支持100个。使用前请您登录[控制台 - 服务配置](https://console.cloud.tencent.com/gamegme/conf)开启语音内容安全服务。
//
// </br></br>
//
// 
//
// <h4><b>功能试用说明：</b></h4>
//
// <li>打开前往<a href="https://console.cloud.tencent.com/gamegme/tryout">控制台 - 产品试用</a>免费试用语音内容安全服务。</li>
//
// </br>
//
// 
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音流或语音文件进行检测，判断其中是否包含违规内容。</li>
//
// <li>支持设置回调地址 Callback 获取检测结果，同时支持通过接口(查询语音检测结果)主动轮询获取检测结果。</li>
//
// <li>支持场景输入，包括：谩骂、色情等场景</li>
//
// <li>支持批量提交检测任务。检测任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：100 M</li>
//
// <li>音频文件时长限制：30分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// <h4><b>语音流限制说明：</b></h4>
//
// <li>语音流格式支持的类型：.m3u8、.flv</li>
//
// <li>语音流支持的传输协议：RTMP、HTTP、HTTPS</li>
//
// <li>语音流时长限制：4小时</li>
//
// <li>支持音视频流分离并对音频流进行分析</li>
//
// </br>
//
// <h4 id="Label_Value"><b>Scenes 与 Label 参数说明：</b></h4>
//
// <p>提交语音检测任务时，需要指定 Scenes 场景参数，<font color="red">目前要求您设置 Scenes 参数值为：["default"]</font>；而在检测结果中，则包含请求时指定的场景，以及对应类型的检测结果。</p>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>场景</th>
//
// <th>描述</th>
//
// <th>Label</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>语音检测</td>
//
// <td>语音检测的检测类型</td>
//
// <td>
//
// <p>normal:正常文本</p>
//
// <p>porn:色情</p>
//
// <p>abuse:谩骂</p>
//
// <p>ad :广告</p>
//
// <p>illegal :违法</p>
//
// <p>moan :呻吟</p>
//
// <p>customized:自定义词库</p>
//
// </td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// </br>
//
// <h4 id="Callback_Declare"><b>回调相关说明：</b></h4>
//
// <li>如果在请求参数中指定了回调地址参数 Callback，即一个 HTTP(S) 协议接口的 URL，则需要支持 POST 方法，传输数据编码采用 UTF-8。</li>
//
// <li>在推送回调数据后，接收到的 HTTP 状态码为 200 时，表示推送成功。</li>
//
// <li>HTTP 请求参数（query）说明：</li>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>名称</th>
//
// <th>类型</th>
//
// <th>是否必需</th>
//
// <th>描述</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Signatue</td>
//
// <td>string</td>
//
// <td>是</td>
//
// <td>签名，具体见<a href="#Callback_Signatue">签名生成说明</a></td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// <ul  id="Callback_Signatue">
//
// 	<li>签名生成说明：</li>
//
// 	<ul>
//
// 		<li>使用 HMAC-SH1 算法, 最终结果做 BASE64 编码;</li>
//
// 		<li>签名原文串为 POST+body 的整个json内容(长度以 Content-Length 为准);</li>
//
// 		<li>签名key为应用的 SecretKey，可以通过控制台查看。</li>
//
// 	</ul>
//
// </ul>
//
// 
//
// <li>回调示例如下<font color="red">（详细字段说明见结构：
//
// <a href="https://cloud.tencent.com/document/api/607/35375#DescribeScanResult" target="_blank">DescribeScanResult</a>）</font>：</li>
//
// <pre><code>{
//
// 	"Code": 0,
//
// 	"DataId": "1400000000_test_data_id",
//
// 	"ScanFinishTime": 1566720906,
//
// 	"HitFlag": true,
//
// 	"Live": false,
//
// 	"Msg": "",
//
// 	"ScanPiece": [{
//
// 		"DumpUrl": "",
//
// 		"HitFlag": true,
//
// 		"MainType": "abuse",
//
// 		"RoomId": "123",
//
// 		"OpenId": "111",
//
// 		"Info":"",
//
// 		"Offset": 0,
//
// 		"Duration": 3400,
//
// 		"PieceStartTime":1574684231,
//
// 		"ScanDetail": [{
//
// 			"EndTime": 1110,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 1110
//
// 		}, {
//
// 			"EndTime": 1380,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 1560,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 2820,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 2490
//
// 		}]
//
// 	}],
//
// 	"ScanStartTime": 1566720905,
//
// 	"Scenes": [
//
// 		"default"
//
// 	],
//
// 	"Status": "Success",
//
// 	"TaskId": "6330xxxx-9xx7-11ed-98e3-52xxxxe4ac3b",
//
// 	"Url": "https://xxx/xxx.m4a"
//
// }
//
// </code></pre>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ScanVoice(request *ScanVoiceRequest) (response *ScanVoiceResponse, err error) {
    return c.ScanVoiceWithContext(context.Background(), request)
}

// ScanVoice
// 本接口(ScanVoice)用于提交语音检测任务，检测任务列表最多支持100个。使用前请您登录[控制台 - 服务配置](https://console.cloud.tencent.com/gamegme/conf)开启语音内容安全服务。
//
// </br></br>
//
// 
//
// <h4><b>功能试用说明：</b></h4>
//
// <li>打开前往<a href="https://console.cloud.tencent.com/gamegme/tryout">控制台 - 产品试用</a>免费试用语音内容安全服务。</li>
//
// </br>
//
// 
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音流或语音文件进行检测，判断其中是否包含违规内容。</li>
//
// <li>支持设置回调地址 Callback 获取检测结果，同时支持通过接口(查询语音检测结果)主动轮询获取检测结果。</li>
//
// <li>支持场景输入，包括：谩骂、色情等场景</li>
//
// <li>支持批量提交检测任务。检测任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：100 M</li>
//
// <li>音频文件时长限制：30分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// <h4><b>语音流限制说明：</b></h4>
//
// <li>语音流格式支持的类型：.m3u8、.flv</li>
//
// <li>语音流支持的传输协议：RTMP、HTTP、HTTPS</li>
//
// <li>语音流时长限制：4小时</li>
//
// <li>支持音视频流分离并对音频流进行分析</li>
//
// </br>
//
// <h4 id="Label_Value"><b>Scenes 与 Label 参数说明：</b></h4>
//
// <p>提交语音检测任务时，需要指定 Scenes 场景参数，<font color="red">目前要求您设置 Scenes 参数值为：["default"]</font>；而在检测结果中，则包含请求时指定的场景，以及对应类型的检测结果。</p>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>场景</th>
//
// <th>描述</th>
//
// <th>Label</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>语音检测</td>
//
// <td>语音检测的检测类型</td>
//
// <td>
//
// <p>normal:正常文本</p>
//
// <p>porn:色情</p>
//
// <p>abuse:谩骂</p>
//
// <p>ad :广告</p>
//
// <p>illegal :违法</p>
//
// <p>moan :呻吟</p>
//
// <p>customized:自定义词库</p>
//
// </td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// </br>
//
// <h4 id="Callback_Declare"><b>回调相关说明：</b></h4>
//
// <li>如果在请求参数中指定了回调地址参数 Callback，即一个 HTTP(S) 协议接口的 URL，则需要支持 POST 方法，传输数据编码采用 UTF-8。</li>
//
// <li>在推送回调数据后，接收到的 HTTP 状态码为 200 时，表示推送成功。</li>
//
// <li>HTTP 请求参数（query）说明：</li>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>名称</th>
//
// <th>类型</th>
//
// <th>是否必需</th>
//
// <th>描述</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Signatue</td>
//
// <td>string</td>
//
// <td>是</td>
//
// <td>签名，具体见<a href="#Callback_Signatue">签名生成说明</a></td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// <ul  id="Callback_Signatue">
//
// 	<li>签名生成说明：</li>
//
// 	<ul>
//
// 		<li>使用 HMAC-SH1 算法, 最终结果做 BASE64 编码;</li>
//
// 		<li>签名原文串为 POST+body 的整个json内容(长度以 Content-Length 为准);</li>
//
// 		<li>签名key为应用的 SecretKey，可以通过控制台查看。</li>
//
// 	</ul>
//
// </ul>
//
// 
//
// <li>回调示例如下<font color="red">（详细字段说明见结构：
//
// <a href="https://cloud.tencent.com/document/api/607/35375#DescribeScanResult" target="_blank">DescribeScanResult</a>）</font>：</li>
//
// <pre><code>{
//
// 	"Code": 0,
//
// 	"DataId": "1400000000_test_data_id",
//
// 	"ScanFinishTime": 1566720906,
//
// 	"HitFlag": true,
//
// 	"Live": false,
//
// 	"Msg": "",
//
// 	"ScanPiece": [{
//
// 		"DumpUrl": "",
//
// 		"HitFlag": true,
//
// 		"MainType": "abuse",
//
// 		"RoomId": "123",
//
// 		"OpenId": "111",
//
// 		"Info":"",
//
// 		"Offset": 0,
//
// 		"Duration": 3400,
//
// 		"PieceStartTime":1574684231,
//
// 		"ScanDetail": [{
//
// 			"EndTime": 1110,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 1110
//
// 		}, {
//
// 			"EndTime": 1380,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 1560,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 2820,
//
// 			"KeyWord": "违规字",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 2490
//
// 		}]
//
// 	}],
//
// 	"ScanStartTime": 1566720905,
//
// 	"Scenes": [
//
// 		"default"
//
// 	],
//
// 	"Status": "Success",
//
// 	"TaskId": "6330xxxx-9xx7-11ed-98e3-52xxxxe4ac3b",
//
// 	"Url": "https://xxx/xxx.m4a"
//
// }
//
// </code></pre>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ScanVoiceWithContext(ctx context.Context, request *ScanVoiceRequest) (response *ScanVoiceResponse, err error) {
    if request == nil {
        request = NewScanVoiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "ScanVoice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanVoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanVoiceResponse()
    err = c.Send(request, response)
    return
}

func NewStartAIConversationRequest() (request *StartAIConversationRequest) {
    request = &StartAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "StartAIConversation")
    
    
    return
}

func NewStartAIConversationResponse() (response *StartAIConversationResponse) {
    response = &StartAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAIConversation
// 启动AI对话任务，AI通道机器人进入GME房间，与房间内指定的成员进行AI对话，适用于智能客服，AI口语教师等场景
//
// 
//
// GME AI对话功能内置语音转文本能力，同时提供通道服务，即客户可灵活指定第三方AI模型（LLM）服务和文本转音频（TTS)服务，更多[功能说明](https://cloud.tencent.com/document/product/647/108901)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAIConversation(request *StartAIConversationRequest) (response *StartAIConversationResponse, err error) {
    return c.StartAIConversationWithContext(context.Background(), request)
}

// StartAIConversation
// 启动AI对话任务，AI通道机器人进入GME房间，与房间内指定的成员进行AI对话，适用于智能客服，AI口语教师等场景
//
// 
//
// GME AI对话功能内置语音转文本能力，同时提供通道服务，即客户可灵活指定第三方AI模型（LLM）服务和文本转音频（TTS)服务，更多[功能说明](https://cloud.tencent.com/document/product/647/108901)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAIConversationWithContext(ctx context.Context, request *StartAIConversationRequest) (response *StartAIConversationResponse, err error) {
    if request == nil {
        request = NewStartAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "StartAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewStartRecordRequest() (request *StartRecordRequest) {
    request = &StartRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "StartRecord")
    
    
    return
}

func NewStartRecordResponse() (response *StartRecordResponse) {
    response = &StartRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartRecord
// 开启录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCEINUSE_TASKINUSE = "ResourceInUse.TaskInUse"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StartRecord(request *StartRecordRequest) (response *StartRecordResponse, err error) {
    return c.StartRecordWithContext(context.Background(), request)
}

// StartRecord
// 开启录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDRECORDMODE = "InvalidParameterValue.InvalidRecordMode"
//  INVALIDPARAMETERVALUE_INVALIDROOMID = "InvalidParameterValue.InvalidRoomId"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBERECORDUSERIDS = "InvalidParameterValue.InvalidSubscribeRecordUserIds"
//  INVALIDPARAMETERVALUE_INVALIDSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidSubscribeUserIds"
//  INVALIDPARAMETERVALUE_INVALIDUNSUBSCRIBEUSERIDS = "InvalidParameterValue.InvalidUNSubscribeUserIds"
//  RESOURCEINUSE_TASKINUSE = "ResourceInUse.TaskInUse"
//  RESOURCENOTFOUND_ROOMNOTFOUND = "ResourceNotFound.RoomNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StartRecordWithContext(ctx context.Context, request *StartRecordRequest) (response *StartRecordResponse, err error) {
    if request == nil {
        request = NewStartRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "StartRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopAIConversationRequest() (request *StopAIConversationRequest) {
    request = &StopAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "StopAIConversation")
    
    
    return
}

func NewStopAIConversationResponse() (response *StopAIConversationResponse) {
    response = &StopAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAIConversation
// 停止AI对话任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAIConversation(request *StopAIConversationRequest) (response *StopAIConversationResponse, err error) {
    return c.StopAIConversationWithContext(context.Background(), request)
}

// StopAIConversation
// 停止AI对话任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAIConversationWithContext(ctx context.Context, request *StopAIConversationRequest) (response *StopAIConversationResponse, err error) {
    if request == nil {
        request = NewStopAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "StopAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewStopRecordRequest() (request *StopRecordRequest) {
    request = &StopRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "StopRecord")
    
    
    return
}

func NewStopRecordResponse() (response *StopRecordResponse) {
    response = &StopRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopRecord
// 停止录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StopRecord(request *StopRecordRequest) (response *StopRecordResponse, err error) {
    return c.StopRecordWithContext(context.Background(), request)
}

// StopRecord
// 停止录制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDBIZID = "InvalidParameterValue.InvalidBizId"
//  INVALIDPARAMETERVALUE_INVALIDTASKID = "InvalidParameterValue.InvalidTaskId"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SERVICENOTOPENED = "UnsupportedOperation.ServiceNotOpened"
func (c *Client) StopRecordWithContext(ctx context.Context, request *StopRecordRequest) (response *StopRecordResponse, err error) {
    if request == nil {
        request = NewStopRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "StopRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRecordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAIConversationRequest() (request *UpdateAIConversationRequest) {
    request = &UpdateAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "UpdateAIConversation")
    
    
    return
}

func NewUpdateAIConversationResponse() (response *UpdateAIConversationResponse) {
    response = &UpdateAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAIConversation
// 更新AIConversation参数
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
func (c *Client) UpdateAIConversation(request *UpdateAIConversationRequest) (response *UpdateAIConversationResponse, err error) {
    return c.UpdateAIConversationWithContext(context.Background(), request)
}

// UpdateAIConversation
// 更新AIConversation参数
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
func (c *Client) UpdateAIConversationWithContext(ctx context.Context, request *UpdateAIConversationRequest) (response *UpdateAIConversationResponse, err error) {
    if request == nil {
        request = NewUpdateAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "UpdateAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScanRoomsRequest() (request *UpdateScanRoomsRequest) {
    request = &UpdateScanRoomsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "UpdateScanRooms")
    
    
    return
}

func NewUpdateScanRoomsResponse() (response *UpdateScanRoomsResponse) {
    response = &UpdateScanRoomsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateScanRooms
// 更新自定义送检房间号。**接口使用前提**：目前 UpdateScanRooms 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) UpdateScanRooms(request *UpdateScanRoomsRequest) (response *UpdateScanRoomsResponse, err error) {
    return c.UpdateScanRoomsWithContext(context.Background(), request)
}

// UpdateScanRooms
// 更新自定义送检房间号。**接口使用前提**：目前 UpdateScanRooms 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) UpdateScanRoomsWithContext(ctx context.Context, request *UpdateScanRoomsRequest) (response *UpdateScanRoomsResponse, err error) {
    if request == nil {
        request = NewUpdateScanRoomsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "UpdateScanRooms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateScanRooms require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateScanRoomsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScanUsersRequest() (request *UpdateScanUsersRequest) {
    request = &UpdateScanUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "UpdateScanUsers")
    
    
    return
}

func NewUpdateScanUsersResponse() (response *UpdateScanUsersResponse) {
    response = &UpdateScanUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateScanUsers
// 更新自定义送检用户号。
//
// **接口使用前提**：目前 UpdateScanUsers 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) UpdateScanUsers(request *UpdateScanUsersRequest) (response *UpdateScanUsersResponse, err error) {
    return c.UpdateScanUsersWithContext(context.Background(), request)
}

// UpdateScanUsers
// 更新自定义送检用户号。
//
// **接口使用前提**：目前 UpdateScanUsers 接口通过白名单开放，如需使用，需要 [提交工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=438&level2_id=445&source=0&data_title=%E6%B8%B8%E6%88%8F%E5%A4%9A%E5%AA%92%E4%BD%93%E5%BC%95%E6%93%8EGME&step=1)。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) UpdateScanUsersWithContext(ctx context.Context, request *UpdateScanUsersRequest) (response *UpdateScanUsersResponse, err error) {
    if request == nil {
        request = NewUpdateScanUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "UpdateScanUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateScanUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateScanUsersResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateVoicePrintRequest() (request *UpdateVoicePrintRequest) {
    request = &UpdateVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gme", APIVersion, "UpdateVoicePrint")
    
    
    return
}

func NewUpdateVoicePrintResponse() (response *UpdateVoicePrintResponse) {
    response = &UpdateVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateVoicePrint
// 传入声纹ID以及对应音频信息，更新对应声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
//  MISSINGPARAMETER_MISSINGVOICEPRINTUPDATEPARAMS = "MissingParameter.MissingVoicePrintUpdateParams"
func (c *Client) UpdateVoicePrint(request *UpdateVoicePrintRequest) (response *UpdateVoicePrintResponse, err error) {
    return c.UpdateVoicePrintWithContext(context.Background(), request)
}

// UpdateVoicePrint
// 传入声纹ID以及对应音频信息，更新对应声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
//  MISSINGPARAMETER_MISSINGVOICEPRINTUPDATEPARAMS = "MissingParameter.MissingVoicePrintUpdateParams"
func (c *Client) UpdateVoicePrintWithContext(ctx context.Context, request *UpdateVoicePrintRequest) (response *UpdateVoicePrintResponse, err error) {
    if request == nil {
        request = NewUpdateVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "gme", APIVersion, "UpdateVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateVoicePrintResponse()
    err = c.Send(request, response)
    return
}
