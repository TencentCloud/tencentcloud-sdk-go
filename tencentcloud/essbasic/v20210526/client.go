// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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

package v20210526

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-26"

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


func NewChannelBatchCancelFlowsRequest() (request *ChannelBatchCancelFlowsRequest) {
    request = &ChannelBatchCancelFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelBatchCancelFlows")
    
    
    return
}

func NewChannelBatchCancelFlowsResponse() (response *ChannelBatchCancelFlowsResponse) {
    response = &ChannelBatchCancelFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelBatchCancelFlows
// 通过合同编号批量撤销合同，单次最多支持撤销100份合同。
//
// 
//
// 适用场景：如果某个合同当前**至少还有一方没有签署**，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 批量撤销结果可以通过接口返回的TaskId关联[批量撤销任务结果回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)或通过接口[查询批量撤销合同结果](https://qian.tencent.com/developers/partnerApis/operateFlows/DescribeCancelFlowsTask)主动查询。
//
// 
//
// 
//
// 注:
//
// - 有对应合同撤销权限的人:  <font color='red'>**合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人**</font>
//
// - 签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>
//
// - <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// - 撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员， 其他参与人员看不到合同的内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelBatchCancelFlows(request *ChannelBatchCancelFlowsRequest) (response *ChannelBatchCancelFlowsResponse, err error) {
    return c.ChannelBatchCancelFlowsWithContext(context.Background(), request)
}

// ChannelBatchCancelFlows
// 通过合同编号批量撤销合同，单次最多支持撤销100份合同。
//
// 
//
// 适用场景：如果某个合同当前**至少还有一方没有签署**，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 批量撤销结果可以通过接口返回的TaskId关联[批量撤销任务结果回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)或通过接口[查询批量撤销合同结果](https://qian.tencent.com/developers/partnerApis/operateFlows/DescribeCancelFlowsTask)主动查询。
//
// 
//
// 
//
// 注:
//
// - 有对应合同撤销权限的人:  <font color='red'>**合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人**</font>
//
// - 签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>
//
// - <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// - 撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员， 其他参与人员看不到合同的内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelBatchCancelFlowsWithContext(ctx context.Context, request *ChannelBatchCancelFlowsRequest) (response *ChannelBatchCancelFlowsResponse, err error) {
    if request == nil {
        request = NewChannelBatchCancelFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelBatchCancelFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelBatchCancelFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCancelFlowRequest() (request *ChannelCancelFlowRequest) {
    request = &ChannelCancelFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCancelFlow")
    
    
    return
}

func NewChannelCancelFlowResponse() (response *ChannelCancelFlowResponse) {
    response = &ChannelCancelFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCancelFlow
// 撤销签署流程接口
//
// 
//
// 适用场景：如果某个合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 注:
//
// - 有对应合同撤销权限的人:  <font color='red'>**合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人**</font>
//
// - 签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>
//
// - <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// - 撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员， 其他参与人员看不到合同的内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCancelFlow(request *ChannelCancelFlowRequest) (response *ChannelCancelFlowResponse, err error) {
    return c.ChannelCancelFlowWithContext(context.Background(), request)
}

// ChannelCancelFlow
// 撤销签署流程接口
//
// 
//
// 适用场景：如果某个合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 注:
//
// - 有对应合同撤销权限的人:  <font color='red'>**合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人**</font>
//
// - 签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>
//
// - <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// - 撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员， 其他参与人员看不到合同的内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCancelFlowWithContext(ctx context.Context, request *ChannelCancelFlowRequest) (response *ChannelCancelFlowResponse, err error) {
    if request == nil {
        request = NewChannelCancelFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCancelFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCancelFlowResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCancelMultiFlowSignQRCodeRequest() (request *ChannelCancelMultiFlowSignQRCodeRequest) {
    request = &ChannelCancelMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCancelMultiFlowSignQRCode")
    
    
    return
}

func NewChannelCancelMultiFlowSignQRCodeResponse() (response *ChannelCancelMultiFlowSignQRCodeResponse) {
    response = &ChannelCancelMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCancelMultiFlowSignQRCode
// 此接口（CancelMultiFlowSignQRCode）用于废除取消一码多签签署码。
//
// 该接口所需的二维码ID，源自[创建一码多签签署码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)生成的。
//
// 如果该签署码尚处于有效期内，可通过本接口将其设置为失效状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCancelMultiFlowSignQRCode(request *ChannelCancelMultiFlowSignQRCodeRequest) (response *ChannelCancelMultiFlowSignQRCodeResponse, err error) {
    return c.ChannelCancelMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// ChannelCancelMultiFlowSignQRCode
// 此接口（CancelMultiFlowSignQRCode）用于废除取消一码多签签署码。
//
// 该接口所需的二维码ID，源自[创建一码多签签署码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)生成的。
//
// 如果该签署码尚处于有效期内，可通过本接口将其设置为失效状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCancelMultiFlowSignQRCodeWithContext(ctx context.Context, request *ChannelCancelMultiFlowSignQRCodeRequest) (response *ChannelCancelMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewChannelCancelMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCancelMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCancelMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCancelUserAutoSignEnableUrlRequest() (request *ChannelCancelUserAutoSignEnableUrlRequest) {
    request = &ChannelCancelUserAutoSignEnableUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCancelUserAutoSignEnableUrl")
    
    
    return
}

func NewChannelCancelUserAutoSignEnableUrlResponse() (response *ChannelCancelUserAutoSignEnableUrlResponse) {
    response = &ChannelCancelUserAutoSignEnableUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCancelUserAutoSignEnableUrl
// 此接口（ChannelCancelUserAutoSignEnableUrl）用来撤销发送给个人用户的自动签开通链接，撤销后对应的个人用户开通链接失效。若个人用户已经完成开通，将无法撤销。（处方单场景专用，使用此接口请与客户经理确认）
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERAUTOSIGNENABLEALREADY = "FailedOperation.UserAutoSignEnableAlready"
//  FAILEDOPERATION_USERAUTOSIGNENABLEURLNOTEXIST = "FailedOperation.UserAutoSignEnableUrlNotExist"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCancelUserAutoSignEnableUrl(request *ChannelCancelUserAutoSignEnableUrlRequest) (response *ChannelCancelUserAutoSignEnableUrlResponse, err error) {
    return c.ChannelCancelUserAutoSignEnableUrlWithContext(context.Background(), request)
}

// ChannelCancelUserAutoSignEnableUrl
// 此接口（ChannelCancelUserAutoSignEnableUrl）用来撤销发送给个人用户的自动签开通链接，撤销后对应的个人用户开通链接失效。若个人用户已经完成开通，将无法撤销。（处方单场景专用，使用此接口请与客户经理确认）
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERAUTOSIGNENABLEALREADY = "FailedOperation.UserAutoSignEnableAlready"
//  FAILEDOPERATION_USERAUTOSIGNENABLEURLNOTEXIST = "FailedOperation.UserAutoSignEnableUrlNotExist"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCancelUserAutoSignEnableUrlWithContext(ctx context.Context, request *ChannelCancelUserAutoSignEnableUrlRequest) (response *ChannelCancelUserAutoSignEnableUrlResponse, err error) {
    if request == nil {
        request = NewChannelCancelUserAutoSignEnableUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCancelUserAutoSignEnableUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCancelUserAutoSignEnableUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateBatchCancelFlowUrlRequest() (request *ChannelCreateBatchCancelFlowUrlRequest) {
    request = &ChannelCreateBatchCancelFlowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateBatchCancelFlowUrl")
    
    
    return
}

func NewChannelCreateBatchCancelFlowUrlResponse() (response *ChannelCreateBatchCancelFlowUrlResponse) {
    response = &ChannelCreateBatchCancelFlowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateBatchCancelFlowUrl
// 通过合同编号生成批量撤销合同的链接，单次最多支持撤销100份合同,   返回的链接需要有此权限的人<font color='red'>**合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人**</font>在<font color='red'>**手机端**</font>打开,  跳转到腾讯电子签小程序输入撤销原因来进行撤销合同
//
// 
//
// 适用场景：如果某个合同当前**至少还有一方没有签署**，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 批量撤销结果可以通过接口返回的TaskId关联[批量撤销任务结果回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)或通过接口[查询批量撤销合同结果](https://qian.tencent.com/developers/partnerApis/operateFlows/DescribeCancelFlowsTask)主动查询。
//
// 
//
// 注:
//
// - 签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>
//
// - <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// - 撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员， 其他参与人员看不到合同的内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateBatchCancelFlowUrl(request *ChannelCreateBatchCancelFlowUrlRequest) (response *ChannelCreateBatchCancelFlowUrlResponse, err error) {
    return c.ChannelCreateBatchCancelFlowUrlWithContext(context.Background(), request)
}

// ChannelCreateBatchCancelFlowUrl
// 通过合同编号生成批量撤销合同的链接，单次最多支持撤销100份合同,   返回的链接需要有此权限的人<font color='red'>**合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人**</font>在<font color='red'>**手机端**</font>打开,  跳转到腾讯电子签小程序输入撤销原因来进行撤销合同
//
// 
//
// 适用场景：如果某个合同当前**至少还有一方没有签署**，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 批量撤销结果可以通过接口返回的TaskId关联[批量撤销任务结果回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)或通过接口[查询批量撤销合同结果](https://qian.tencent.com/developers/partnerApis/operateFlows/DescribeCancelFlowsTask)主动查询。
//
// 
//
// 注:
//
// - 签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>
//
// - <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// - 撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员， 其他参与人员看不到合同的内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateBatchCancelFlowUrlWithContext(ctx context.Context, request *ChannelCreateBatchCancelFlowUrlRequest) (response *ChannelCreateBatchCancelFlowUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateBatchCancelFlowUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateBatchCancelFlowUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateBatchCancelFlowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateBatchQuickSignUrlRequest() (request *ChannelCreateBatchQuickSignUrlRequest) {
    request = &ChannelCreateBatchQuickSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateBatchQuickSignUrl")
    
    
    return
}

func NewChannelCreateBatchQuickSignUrlResponse() (response *ChannelCreateBatchQuickSignUrlResponse) {
    response = &ChannelCreateBatchQuickSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateBatchQuickSignUrl
// 该接口用于发起合同后，生成个人用户的批量待办链接, 暂时不支持企业端签署。
//
// **注意：**
//
// 1. 该接口目前仅支持签署人类型是**个人签署方的批量签署场景**(ApproverType=1)。
//
// 2. 该接口可生成C端签署人的批量签署/查看链接，**签署控件仅支持手写签名(控件类型为SIGN_SIGNATURE)和时间类型的签署控件** 。
//
// 3. 该签署链接**有效期为30分钟**，过期后将失效，如需签署可重新创建批量签署链接 。
//
// 4. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。
//
// 跳转到小程序的实现，参考微信官方文档(分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式)，如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 6. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateBatchQuickSignUrl(request *ChannelCreateBatchQuickSignUrlRequest) (response *ChannelCreateBatchQuickSignUrlResponse, err error) {
    return c.ChannelCreateBatchQuickSignUrlWithContext(context.Background(), request)
}

// ChannelCreateBatchQuickSignUrl
// 该接口用于发起合同后，生成个人用户的批量待办链接, 暂时不支持企业端签署。
//
// **注意：**
//
// 1. 该接口目前仅支持签署人类型是**个人签署方的批量签署场景**(ApproverType=1)。
//
// 2. 该接口可生成C端签署人的批量签署/查看链接，**签署控件仅支持手写签名(控件类型为SIGN_SIGNATURE)和时间类型的签署控件** 。
//
// 3. 该签署链接**有效期为30分钟**，过期后将失效，如需签署可重新创建批量签署链接 。
//
// 4. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。
//
// 跳转到小程序的实现，参考微信官方文档(分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式)，如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 6. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateBatchQuickSignUrlWithContext(ctx context.Context, request *ChannelCreateBatchQuickSignUrlRequest) (response *ChannelCreateBatchQuickSignUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateBatchQuickSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateBatchQuickSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateBatchQuickSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateBatchSignUrlRequest() (request *ChannelCreateBatchSignUrlRequest) {
    request = &ChannelCreateBatchSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateBatchSignUrl")
    
    
    return
}

func NewChannelCreateBatchSignUrlResponse() (response *ChannelCreateBatchSignUrlResponse) {
    response = &ChannelCreateBatchSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateBatchSignUrl
// 通过此接口，创建小程序批量签署链接，个人/企业员工点击此链接即可跳转小程序进行批量签署。
//
// 请确保生成链接时候的身份信息和签署合同参与方的信息保持一致。
//
// 
//
// 注：
//
// - 使用此接口生成链接，需要提前开通 `个人签署方仅校验手机号` 功能，在 `腾讯电子签网页端-企业设置-拓展服务` 中可以找到。
//
// - 个人参与方点击链接后需短信验证码才能查看合同内容。
//
// - 个人用户批量签署，需要传Name，Mobile，IdCardNumber(IdCardType) 参数。
//
// - saas企业员工用户批量签署，在传递了姓名等基本信息参数的情况下，还需要传OrganizationName（参与方所在企业名称）参数生成签署链接，<font color="red">请确保此企业已完成腾讯电子签企业认证</font>。
//
// - 子客企业员工用户批量签署，需要传递员工OpenId和子客企业的OrganizationOpenId。<font color="red">请确保此OrganizationOpenId对应子客已经认证，且OpenId对应员工此子客下已经实名</font>。Name，Mobile, IdCard等信息此时可以不传，系统会查询此OpenId实名信息自动补充。
//
// - 生成批量签署链接时，合同目标参与方状态需为<font color="red">待签署</font>状态。
//
// - 个人批量签署进行的合同的签名区， 全部变成<font color="red">手写签名</font>（不管合同里边设置的签名限制）来进行。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
func (c *Client) ChannelCreateBatchSignUrl(request *ChannelCreateBatchSignUrlRequest) (response *ChannelCreateBatchSignUrlResponse, err error) {
    return c.ChannelCreateBatchSignUrlWithContext(context.Background(), request)
}

// ChannelCreateBatchSignUrl
// 通过此接口，创建小程序批量签署链接，个人/企业员工点击此链接即可跳转小程序进行批量签署。
//
// 请确保生成链接时候的身份信息和签署合同参与方的信息保持一致。
//
// 
//
// 注：
//
// - 使用此接口生成链接，需要提前开通 `个人签署方仅校验手机号` 功能，在 `腾讯电子签网页端-企业设置-拓展服务` 中可以找到。
//
// - 个人参与方点击链接后需短信验证码才能查看合同内容。
//
// - 个人用户批量签署，需要传Name，Mobile，IdCardNumber(IdCardType) 参数。
//
// - saas企业员工用户批量签署，在传递了姓名等基本信息参数的情况下，还需要传OrganizationName（参与方所在企业名称）参数生成签署链接，<font color="red">请确保此企业已完成腾讯电子签企业认证</font>。
//
// - 子客企业员工用户批量签署，需要传递员工OpenId和子客企业的OrganizationOpenId。<font color="red">请确保此OrganizationOpenId对应子客已经认证，且OpenId对应员工此子客下已经实名</font>。Name，Mobile, IdCard等信息此时可以不传，系统会查询此OpenId实名信息自动补充。
//
// - 生成批量签署链接时，合同目标参与方状态需为<font color="red">待签署</font>状态。
//
// - 个人批量签署进行的合同的签名区， 全部变成<font color="red">手写签名</font>（不管合同里边设置的签名限制）来进行。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
func (c *Client) ChannelCreateBatchSignUrlWithContext(ctx context.Context, request *ChannelCreateBatchSignUrlRequest) (response *ChannelCreateBatchSignUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateBatchSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateBatchSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateBatchSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateBoundFlowsRequest() (request *ChannelCreateBoundFlowsRequest) {
    request = &ChannelCreateBoundFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateBoundFlows")
    
    
    return
}

func NewChannelCreateBoundFlowsResponse() (response *ChannelCreateBoundFlowsResponse) {
    response = &ChannelCreateBoundFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateBoundFlows
// 此接口（ChannelCreateBoundFlows）用于子客企业领取未归属给员工的合同，将合同领取给当前员工，合同不能重复领取。
//
// 
//
// 
//
// **未归属合同发起方式**
//
//  指定对应企业的OrganizationOpenId和OrganizationName而不指定具体的参与人(OpenId/名字/手机号等),  则合同进入待领取状态, 示例代码如下
//
// ```
//
// 		FlowApprovers: []*essbasic.FlowApproverInfo{
//
// 			{
//
// 				ApproverType:       common.StringPtr("ORGANIZATION"),
//
// 				OrganizationOpenId: common.StringPtr("org_dianziqian"),
//
// 				OrganizationName:   common.StringPtr("典子谦示例企业"),
//
// 			}
//
// 		},
//
// ```
//
// 
//
// 可以<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录控制台查看带领取的合同
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/a34d0cc56ec871613e94dfc6252bc072.png)
//
// 
//
// 注: 
//
// 1. 支持批量领取,  如果有一个合同流程无法领取会导致接口报错,  使得所有合同都领取失败
//
// 2. 只有企业的<font color="red">超管或者法人</font>才能进行合同的领取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASAUTHORIZED = "FailedOperation.HasAuthorized"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateBoundFlows(request *ChannelCreateBoundFlowsRequest) (response *ChannelCreateBoundFlowsResponse, err error) {
    return c.ChannelCreateBoundFlowsWithContext(context.Background(), request)
}

// ChannelCreateBoundFlows
// 此接口（ChannelCreateBoundFlows）用于子客企业领取未归属给员工的合同，将合同领取给当前员工，合同不能重复领取。
//
// 
//
// 
//
// **未归属合同发起方式**
//
//  指定对应企业的OrganizationOpenId和OrganizationName而不指定具体的参与人(OpenId/名字/手机号等),  则合同进入待领取状态, 示例代码如下
//
// ```
//
// 		FlowApprovers: []*essbasic.FlowApproverInfo{
//
// 			{
//
// 				ApproverType:       common.StringPtr("ORGANIZATION"),
//
// 				OrganizationOpenId: common.StringPtr("org_dianziqian"),
//
// 				OrganizationName:   common.StringPtr("典子谦示例企业"),
//
// 			}
//
// 		},
//
// ```
//
// 
//
// 可以<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>登录控制台查看带领取的合同
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/a34d0cc56ec871613e94dfc6252bc072.png)
//
// 
//
// 注: 
//
// 1. 支持批量领取,  如果有一个合同流程无法领取会导致接口报错,  使得所有合同都领取失败
//
// 2. 只有企业的<font color="red">超管或者法人</font>才能进行合同的领取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASAUTHORIZED = "FailedOperation.HasAuthorized"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateBoundFlowsWithContext(ctx context.Context, request *ChannelCreateBoundFlowsRequest) (response *ChannelCreateBoundFlowsResponse, err error) {
    if request == nil {
        request = NewChannelCreateBoundFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateBoundFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateBoundFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateConvertTaskApiRequest() (request *ChannelCreateConvertTaskApiRequest) {
    request = &ChannelCreateConvertTaskApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateConvertTaskApi")
    
    
    return
}

func NewChannelCreateConvertTaskApiResponse() (response *ChannelCreateConvertTaskApiResponse) {
    response = &ChannelCreateConvertTaskApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateConvertTaskApi
// 此接口（ChannelCreateConvertTaskApi）用来将word、excel、html、图片、txt类型文件转换为PDF文件。<br />
//
// 前提条件：源文件已经通过 <a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">文件上传接口</a>完成上传，并得到了源文件的资源Id。<br />
//
// 适用场景1：已经上传了一个word文件，希望将该word文件转换成pdf文件后发起合同
//
// 适用场景2：已经上传了一个jpg图片文件，希望将该图片文件转换成pdf文件后发起合同<br />
//
// 转换文件是一个耗时操作，若想查看转换任务是否完成，可以通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">查询转换任务状态</a>接口获取任务状态。<br />
//
// 注: 
//
// 1. `支持的文件类型有doc、docx、xls、xlsx、html、jpg、jpeg、png、bmp、txt`
//
// 2. `可通过发起合同时设置预览来检查转换文件是否达到预期效果`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateConvertTaskApi(request *ChannelCreateConvertTaskApiRequest) (response *ChannelCreateConvertTaskApiResponse, err error) {
    return c.ChannelCreateConvertTaskApiWithContext(context.Background(), request)
}

// ChannelCreateConvertTaskApi
// 此接口（ChannelCreateConvertTaskApi）用来将word、excel、html、图片、txt类型文件转换为PDF文件。<br />
//
// 前提条件：源文件已经通过 <a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">文件上传接口</a>完成上传，并得到了源文件的资源Id。<br />
//
// 适用场景1：已经上传了一个word文件，希望将该word文件转换成pdf文件后发起合同
//
// 适用场景2：已经上传了一个jpg图片文件，希望将该图片文件转换成pdf文件后发起合同<br />
//
// 转换文件是一个耗时操作，若想查看转换任务是否完成，可以通过<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelGetTaskResultApi" target="_blank">查询转换任务状态</a>接口获取任务状态。<br />
//
// 注: 
//
// 1. `支持的文件类型有doc、docx、xls、xlsx、html、jpg、jpeg、png、bmp、txt`
//
// 2. `可通过发起合同时设置预览来检查转换文件是否达到预期效果`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateConvertTaskApiWithContext(ctx context.Context, request *ChannelCreateConvertTaskApiRequest) (response *ChannelCreateConvertTaskApiResponse, err error) {
    if request == nil {
        request = NewChannelCreateConvertTaskApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateConvertTaskApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateConvertTaskApiResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateEmbedWebUrlRequest() (request *ChannelCreateEmbedWebUrlRequest) {
    request = &ChannelCreateEmbedWebUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateEmbedWebUrl")
    
    
    return
}

func NewChannelCreateEmbedWebUrlResponse() (response *ChannelCreateEmbedWebUrlResponse) {
    response = &ChannelCreateEmbedWebUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateEmbedWebUrl
// 本接口（ChannelCreateEmbedWebUrl）用于创建可嵌入web页面的URL（此web页面可以通过iframe方式嵌入到贵方系统的网页中），支持以下类型的Web链接创建：
//
// 1. 创建印章
//
// 2. 创建模板
//
// 3. 修改模板
//
// 4. 预览模板
//
// 5. 预览合同流程
//
// 
//
// 预览模板的嵌入页面长相如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/57bdda4a884e3f5b2de12d5a282a3651.png)
//
// 
//
// 预览合同流程的嵌入页面长相如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/dc7af994e2f6da56bdad5975e927de34.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOROLEAUTH = "FailedOperation.NoRoleAuth"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateEmbedWebUrl(request *ChannelCreateEmbedWebUrlRequest) (response *ChannelCreateEmbedWebUrlResponse, err error) {
    return c.ChannelCreateEmbedWebUrlWithContext(context.Background(), request)
}

// ChannelCreateEmbedWebUrl
// 本接口（ChannelCreateEmbedWebUrl）用于创建可嵌入web页面的URL（此web页面可以通过iframe方式嵌入到贵方系统的网页中），支持以下类型的Web链接创建：
//
// 1. 创建印章
//
// 2. 创建模板
//
// 3. 修改模板
//
// 4. 预览模板
//
// 5. 预览合同流程
//
// 
//
// 预览模板的嵌入页面长相如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/57bdda4a884e3f5b2de12d5a282a3651.png)
//
// 
//
// 预览合同流程的嵌入页面长相如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/dc7af994e2f6da56bdad5975e927de34.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOROLEAUTH = "FailedOperation.NoRoleAuth"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateEmbedWebUrlWithContext(ctx context.Context, request *ChannelCreateEmbedWebUrlRequest) (response *ChannelCreateEmbedWebUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateEmbedWebUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateEmbedWebUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateEmbedWebUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowApproversRequest() (request *ChannelCreateFlowApproversRequest) {
    request = &ChannelCreateFlowApproversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowApprovers")
    
    
    return
}

func NewChannelCreateFlowApproversResponse() (response *ChannelCreateFlowApproversResponse) {
    response = &ChannelCreateFlowApproversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateFlowApprovers
// **适用场景**：
//
// 当通过模板或文件发起合同时，若未指定企业签署人信息，则可调用此接口动态补充签署人。同一签署人只允许补充一人，最终实际签署人取决于谁先领取合同完成签署。
//
// 
//
// **接口使用说明**：
//
// 
//
// 1.本接口现已支持批量补充签署人
//
// 
//
// 2.当<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中指定需要补充的FlowId时，可以对指定合同补充签署人；可以指定多个相同发起方的不同合同在完成批量补充
//
// 
//
// 3.当<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中指定需要补充的FlowId时，是对指定的合同补充多个指定的签署人
//
// 
//
// 4.如果同时指定了<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId和<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId，仅使用<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId作为补充的合同
//
// 
//
// 5.如果部分指定了<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId，又指定了<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId；那么<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>存在指定的FlowId，则使用<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId，不存在则使用<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId作为补充的合同
//
// 
//
// 
//
// 6.如果同时未指定了<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId和<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId，则传参错误
//
// 
//
// **限制条件**：
//
// 1. 本企业（发起方企业）企业签署人仅支持通过企业名称+姓名+手机号进行补充。
//
// 2. 个人签署人支持通过姓名+手机号进行补充，补充动态签署人时：若个人用户已完成实名，则可通过姓名+证件号码进行补充。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateFlowApprovers(request *ChannelCreateFlowApproversRequest) (response *ChannelCreateFlowApproversResponse, err error) {
    return c.ChannelCreateFlowApproversWithContext(context.Background(), request)
}

// ChannelCreateFlowApprovers
// **适用场景**：
//
// 当通过模板或文件发起合同时，若未指定企业签署人信息，则可调用此接口动态补充签署人。同一签署人只允许补充一人，最终实际签署人取决于谁先领取合同完成签署。
//
// 
//
// **接口使用说明**：
//
// 
//
// 1.本接口现已支持批量补充签署人
//
// 
//
// 2.当<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中指定需要补充的FlowId时，可以对指定合同补充签署人；可以指定多个相同发起方的不同合同在完成批量补充
//
// 
//
// 3.当<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中指定需要补充的FlowId时，是对指定的合同补充多个指定的签署人
//
// 
//
// 4.如果同时指定了<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId和<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId，仅使用<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId作为补充的合同
//
// 
//
// 5.如果部分指定了<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId，又指定了<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId；那么<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>存在指定的FlowId，则使用<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId，不存在则使用<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId作为补充的合同
//
// 
//
// 
//
// 6.如果同时未指定了<a href="https://qian.tencent.com/developers/partnerApis/dataTypes/#fillapproverinfo/" target="_blank">补充签署人结构体</a>中的FlowId和<a href="https://qian.tencent.com/developers/partnerApis/flows/ChannelCreateFlowApprovers/" target="_blank">补充签署人接口入参</a>中的FlowId，则传参错误
//
// 
//
// **限制条件**：
//
// 1. 本企业（发起方企业）企业签署人仅支持通过企业名称+姓名+手机号进行补充。
//
// 2. 个人签署人支持通过姓名+手机号进行补充，补充动态签署人时：若个人用户已完成实名，则可通过姓名+证件号码进行补充。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateFlowApproversWithContext(ctx context.Context, request *ChannelCreateFlowApproversRequest) (response *ChannelCreateFlowApproversResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowApproversRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowApprovers require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowApproversResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowByFilesRequest() (request *ChannelCreateFlowByFilesRequest) {
    request = &ChannelCreateFlowByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowByFiles")
    
    
    return
}

func NewChannelCreateFlowByFilesResponse() (response *ChannelCreateFlowByFilesResponse) {
    response = &ChannelCreateFlowByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateFlowByFiles
// 接口（ChannelCreateFlowByFiles）用PDF文件创建签署流程。
//
// 
//
// 适用场景：适用非制式的合同文件签署，开发者有每个签署流程的PDF，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。
//
// 
//
// **注**: 
//
// <ul>
//
// <li>此接口静默签(企业自动签)能力为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>此接口需要依赖<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">文件上传接口</a>生成pdf资源编号（FileIds）进行使用。整体的逻辑如下图</li>
//
// </ul>
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/bf86248a2c163228c4e894cf5926af69/ChannelCreateFlowByFiles.png)
//
// 
//
// **可以作为发起方和签署方的角色列表**
//
// <table>     <thead>     <tr>         <th>场景编号</th>         <th>发起方</th>         <th>签署方</th>         <th>补充</th>     </tr>     </thead>     <tbody>     <tr>         <td>场景一</td>         <td>子企业A的员工</td>         <td>子企业A的员工</td>         <td>子企业是通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">CreateConsoleLoginUrl</a>生成子客登录链接注册的企业</td>     </tr>     <tr>         <td>场景二</td>         <td>子企业A的员工</td>         <td>子企业B(不指定经办人走领取逻辑)</td>         <td>领取的逻辑可以参考文档<a href="https://qian.tencent.com/developers/partner/dynamic_signer" target="_blank">动态签署方</a> </td>     </tr>     <tr>         <td>场景三</td>         <td>子企业A的员工</td>         <td>子企业B的员工</td>         <td>-</td>     </tr>     <tr>         <td>场景四</td>         <td>子企业A的员工</td>         <td>个人</td>         <td>就是自然人，不是企业员工</td>     </tr>     <tr>         <td>场景五</td>         <td>子企业A的员工</td>         <td>SaaS平台企业员工</td>         <td>SaaS平台企业是通过<a href="https://qian.tencent.cn/console/company-register" target="_blank">https://qian.tencent.cn/console/company-register</a>链接注册的企业</td>     </tr>     </tbody> </table>
//
// 
//
// 
//
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// 
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
//
// 
//
// `3. 合同发起后就会扣减合同的额度, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（过期，拒签，签署完成，解除完成等状态不会返还额度）`
//
// 
//
// `4. 静默（自动）签署不支持合同签署方存在填写功能`
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-UploadFiles.mp4" target="_blank">【上传文件代码】编写示例</a><br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-ChannelCreateFlowByFiles.mp4" target="_blank">【用PDF文件创建签署流程】编写示例</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  MISSINGPARAMETER_SEALID = "MissingParameter.SealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOAPPROVERMOBILECHECKPERMISSION = "OperationDenied.NoApproverMobileCheckPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowByFiles(request *ChannelCreateFlowByFilesRequest) (response *ChannelCreateFlowByFilesResponse, err error) {
    return c.ChannelCreateFlowByFilesWithContext(context.Background(), request)
}

// ChannelCreateFlowByFiles
// 接口（ChannelCreateFlowByFiles）用PDF文件创建签署流程。
//
// 
//
// 适用场景：适用非制式的合同文件签署，开发者有每个签署流程的PDF，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。
//
// 
//
// **注**: 
//
// <ul>
//
// <li>此接口静默签(企业自动签)能力为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>此接口需要依赖<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">文件上传接口</a>生成pdf资源编号（FileIds）进行使用。整体的逻辑如下图</li>
//
// </ul>
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/bf86248a2c163228c4e894cf5926af69/ChannelCreateFlowByFiles.png)
//
// 
//
// **可以作为发起方和签署方的角色列表**
//
// <table>     <thead>     <tr>         <th>场景编号</th>         <th>发起方</th>         <th>签署方</th>         <th>补充</th>     </tr>     </thead>     <tbody>     <tr>         <td>场景一</td>         <td>子企业A的员工</td>         <td>子企业A的员工</td>         <td>子企业是通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">CreateConsoleLoginUrl</a>生成子客登录链接注册的企业</td>     </tr>     <tr>         <td>场景二</td>         <td>子企业A的员工</td>         <td>子企业B(不指定经办人走领取逻辑)</td>         <td>领取的逻辑可以参考文档<a href="https://qian.tencent.com/developers/partner/dynamic_signer" target="_blank">动态签署方</a> </td>     </tr>     <tr>         <td>场景三</td>         <td>子企业A的员工</td>         <td>子企业B的员工</td>         <td>-</td>     </tr>     <tr>         <td>场景四</td>         <td>子企业A的员工</td>         <td>个人</td>         <td>就是自然人，不是企业员工</td>     </tr>     <tr>         <td>场景五</td>         <td>子企业A的员工</td>         <td>SaaS平台企业员工</td>         <td>SaaS平台企业是通过<a href="https://qian.tencent.cn/console/company-register" target="_blank">https://qian.tencent.cn/console/company-register</a>链接注册的企业</td>     </tr>     </tbody> </table>
//
// 
//
// 
//
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// 
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
//
// 
//
// `3. 合同发起后就会扣减合同的额度, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（过期，拒签，签署完成，解除完成等状态不会返还额度）`
//
// 
//
// `4. 静默（自动）签署不支持合同签署方存在填写功能`
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-UploadFiles.mp4" target="_blank">【上传文件代码】编写示例</a><br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-ChannelCreateFlowByFiles.mp4" target="_blank">【用PDF文件创建签署流程】编写示例</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  MISSINGPARAMETER_SEALID = "MissingParameter.SealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOAPPROVERMOBILECHECKPERMISSION = "OperationDenied.NoApproverMobileCheckPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowByFilesWithContext(ctx context.Context, request *ChannelCreateFlowByFilesRequest) (response *ChannelCreateFlowByFilesResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowByFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowByFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowGroupByFilesRequest() (request *ChannelCreateFlowGroupByFilesRequest) {
    request = &ChannelCreateFlowGroupByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowGroupByFiles")
    
    
    return
}

func NewChannelCreateFlowGroupByFilesResponse() (response *ChannelCreateFlowGroupByFilesResponse) {
    response = &ChannelCreateFlowGroupByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateFlowGroupByFiles
// 接口（ChannelCreateFlowGroupByFiles）用于使用 PDF 文件创建合同组签署流程。
//
// 
//
// - 该接口允许通过选择多个模板一次性创建多个合同，这些合同被组织在一个合同组中。
//
// - 每个签署方将收到一个签署链接，通过这个链接可以访问并签署合同组中的所有合同。
//
// - 合同组中的合同必须作为一个整体进行签署，不能将合同组拆分成单独的合同进行逐一签署。
//
// 
//
// <img src="https://qcloudimg.tencent-cloud.cn/raw/a63074a0293c9ff5bf6c0bb74c0d3b20.png"   width="400" />
//
// 
//
// 
//
// ### 2. 适用场景
//
// 
//
// 该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// ### 3. 发起方要求和签署方实名要求
//
// - **发起方要求**：作为合同发起方的第三方子企业A的员工必须进行实名认证。
//
// - **签署方要求**：签署方可以是多种身份（如第三方子企业的员工、个人、SaaS平台企业员工），其中企业和员工可以不进行实名认证。
//
// 
//
// **可以作为发起方和签署方的角色列表**
//
// 
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>场景编号</th>
//
// <th>可作为发起方类型</th>
//
// <th>可作为签署方的类型</th>
//
// </tr>
//
// </thead>
//
// 
//
// <tbody>
//
// <tr>
//
// <td>场景一</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业A员工</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景二</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业B员工</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景三</td>
//
// <td>第三方子企业A员工</td>
//
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景四</td>
//
// <td>第三方子企业A员工</td>
//
// <td>SaaS平台企业员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 
//
// ### 4. 签署方参数差异
//
// - 根据签署方的不同类型（第三方子企业的员工、个人、SaaS平台企业员工），传递的参数也不同。具体参数的结构和要求可以参考开发者中心提供的 `FlowApproverInfo` 结构体说明。
//
// 
//
// ### 5. 合同额度的扣减与返还
//
// - **扣减时机**：合同一旦发起，相关的合同额度就会被扣减，合同组下面的每个合同都要扣减一个合同额度。
//
// - **返还条件**：只有在合同被撤销且没有任何签署方签署过，或者只有自动签署的情况下，合同额度才会被返还。
//
// - **不返还的情况**：如果合同已过期、被拒签、签署完成或已解除，合同额度将不会被返还。
//
// 
//
// ### 6. 静默（自动）签署的限制
//
// - 在使用静默（自动）签署功能时，合同签署方不能有填写控件。<font color="red">此接口静默签(企业自动签)能力为白名单功能</font>，使用前请联系对接的客户经理沟通。
//
// 
//
// ### 7.合同组暂不支持抄送功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOAPPROVERMOBILECHECKPERMISSION = "OperationDenied.NoApproverMobileCheckPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowGroupByFiles(request *ChannelCreateFlowGroupByFilesRequest) (response *ChannelCreateFlowGroupByFilesResponse, err error) {
    return c.ChannelCreateFlowGroupByFilesWithContext(context.Background(), request)
}

// ChannelCreateFlowGroupByFiles
// 接口（ChannelCreateFlowGroupByFiles）用于使用 PDF 文件创建合同组签署流程。
//
// 
//
// - 该接口允许通过选择多个模板一次性创建多个合同，这些合同被组织在一个合同组中。
//
// - 每个签署方将收到一个签署链接，通过这个链接可以访问并签署合同组中的所有合同。
//
// - 合同组中的合同必须作为一个整体进行签署，不能将合同组拆分成单独的合同进行逐一签署。
//
// 
//
// <img src="https://qcloudimg.tencent-cloud.cn/raw/a63074a0293c9ff5bf6c0bb74c0d3b20.png"   width="400" />
//
// 
//
// 
//
// ### 2. 适用场景
//
// 
//
// 该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// ### 3. 发起方要求和签署方实名要求
//
// - **发起方要求**：作为合同发起方的第三方子企业A的员工必须进行实名认证。
//
// - **签署方要求**：签署方可以是多种身份（如第三方子企业的员工、个人、SaaS平台企业员工），其中企业和员工可以不进行实名认证。
//
// 
//
// **可以作为发起方和签署方的角色列表**
//
// 
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>场景编号</th>
//
// <th>可作为发起方类型</th>
//
// <th>可作为签署方的类型</th>
//
// </tr>
//
// </thead>
//
// 
//
// <tbody>
//
// <tr>
//
// <td>场景一</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业A员工</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景二</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业B员工</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景三</td>
//
// <td>第三方子企业A员工</td>
//
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景四</td>
//
// <td>第三方子企业A员工</td>
//
// <td>SaaS平台企业员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 
//
// ### 4. 签署方参数差异
//
// - 根据签署方的不同类型（第三方子企业的员工、个人、SaaS平台企业员工），传递的参数也不同。具体参数的结构和要求可以参考开发者中心提供的 `FlowApproverInfo` 结构体说明。
//
// 
//
// ### 5. 合同额度的扣减与返还
//
// - **扣减时机**：合同一旦发起，相关的合同额度就会被扣减，合同组下面的每个合同都要扣减一个合同额度。
//
// - **返还条件**：只有在合同被撤销且没有任何签署方签署过，或者只有自动签署的情况下，合同额度才会被返还。
//
// - **不返还的情况**：如果合同已过期、被拒签、签署完成或已解除，合同额度将不会被返还。
//
// 
//
// ### 6. 静默（自动）签署的限制
//
// - 在使用静默（自动）签署功能时，合同签署方不能有填写控件。<font color="red">此接口静默签(企业自动签)能力为白名单功能</font>，使用前请联系对接的客户经理沟通。
//
// 
//
// ### 7.合同组暂不支持抄送功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"
//  INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOAPPROVERMOBILECHECKPERMISSION = "OperationDenied.NoApproverMobileCheckPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_OVERSEAABILITYNOTOPEN = "OperationDenied.OverseaAbilityNotOpen"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowGroupByFilesWithContext(ctx context.Context, request *ChannelCreateFlowGroupByFilesRequest) (response *ChannelCreateFlowGroupByFilesResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowGroupByFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowGroupByFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowGroupByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowGroupByTemplatesRequest() (request *ChannelCreateFlowGroupByTemplatesRequest) {
    request = &ChannelCreateFlowGroupByTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowGroupByTemplates")
    
    
    return
}

func NewChannelCreateFlowGroupByTemplatesResponse() (response *ChannelCreateFlowGroupByTemplatesResponse) {
    response = &ChannelCreateFlowGroupByTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateFlowGroupByTemplates
// 接口（ChannelCreateFlowGroupByTemplates）用于通过多模板创建合同组签署流程。
//
// 
//
// - 该接口允许通过选择多个模板一次性创建多个合同，这些合同被组织在一个合同组中。
//
// - 每个签署方将收到一个签署链接，通过这个链接可以访问并签署合同组中的所有合同。
//
// - 合同组中的合同必须作为一个整体进行签署，不能将合同组拆分成单独的合同进行逐一签署。
//
// 
//
// <img src="https://qcloudimg.tencent-cloud.cn/raw/a63074a0293c9ff5bf6c0bb74c0d3b20.png"   width="400" />
//
// 
//
// ### 2. 适用场景
//
// 
//
// 该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// ### 3. 发起方要求和签署方实名要求
//
// - **发起方要求**：作为合同发起方的第三方子企业A的员工必须进行实名认证。
//
// - **签署方要求**：签署方可以是多种身份（如第三方子企业的员工、个人、SaaS平台企业员工），其中企业和员工可以不进行实名认证。
//
// 
//
// **可以作为发起方和签署方的角色列表**
//
// 
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>场景编号</th>
//
// <th>可作为发起方类型</th>
//
// <th>可作为签署方的类型</th>
//
// </tr>
//
// </thead>
//
// 
//
// <tbody>
//
// <tr>
//
// <td>场景一</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业A员工</td>
//
// </tr>
//
// <tr>
//
// <td>场景二</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业B员工</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景三</td>
//
// <td>第三方子企业A员工</td>
//
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景四</td>
//
// <td>第三方子企业A员工</td>
//
// <td>SaaS平台企业员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 
//
// ### 4. 签署方参数差异
//
// - 根据签署方的不同类型（第三方子企业的员工、个人、SaaS平台企业员工），传递的参数也不同。具体参数的结构和要求可以参考开发者中心提供的 `FlowApproverInfo` 结构体说明。
//
// 
//
// ### 5. 合同额度的扣减与返还
//
// - **扣减时机**：合同一旦发起，相关的合同额度就会被扣减，合同组下面的每个合同都要扣减一个合同额度。
//
// - **返还条件**：只有在合同被撤销且没有任何签署方签署过，或者只有自动签署的情况下，合同额度才会被返还。
//
// - **不返还的情况**：如果合同已过期、被拒签、签署完成或已解除，合同额度将不会被返还。
//
// 
//
// ### 6. 静默（自动）签署的限制
//
// - 在使用静默（自动）签署功能时，合同签署方不能有填写控件。<font color="red">此接口静默签(企业自动签)能力为白名单功能</font>，使用前请联系对接的客户经理沟通。
//
// 
//
// ### 7.合同组暂不支持抄送功能
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowGroupByTemplates(request *ChannelCreateFlowGroupByTemplatesRequest) (response *ChannelCreateFlowGroupByTemplatesResponse, err error) {
    return c.ChannelCreateFlowGroupByTemplatesWithContext(context.Background(), request)
}

// ChannelCreateFlowGroupByTemplates
// 接口（ChannelCreateFlowGroupByTemplates）用于通过多模板创建合同组签署流程。
//
// 
//
// - 该接口允许通过选择多个模板一次性创建多个合同，这些合同被组织在一个合同组中。
//
// - 每个签署方将收到一个签署链接，通过这个链接可以访问并签署合同组中的所有合同。
//
// - 合同组中的合同必须作为一个整体进行签署，不能将合同组拆分成单独的合同进行逐一签署。
//
// 
//
// <img src="https://qcloudimg.tencent-cloud.cn/raw/a63074a0293c9ff5bf6c0bb74c0d3b20.png"   width="400" />
//
// 
//
// ### 2. 适用场景
//
// 
//
// 该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// ### 3. 发起方要求和签署方实名要求
//
// - **发起方要求**：作为合同发起方的第三方子企业A的员工必须进行实名认证。
//
// - **签署方要求**：签署方可以是多种身份（如第三方子企业的员工、个人、SaaS平台企业员工），其中企业和员工可以不进行实名认证。
//
// 
//
// **可以作为发起方和签署方的角色列表**
//
// 
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>场景编号</th>
//
// <th>可作为发起方类型</th>
//
// <th>可作为签署方的类型</th>
//
// </tr>
//
// </thead>
//
// 
//
// <tbody>
//
// <tr>
//
// <td>场景一</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业A员工</td>
//
// </tr>
//
// <tr>
//
// <td>场景二</td>
//
// <td>第三方子企业A员工</td>
//
// <td>第三方子企业B员工</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景三</td>
//
// <td>第三方子企业A员工</td>
//
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景四</td>
//
// <td>第三方子企业A员工</td>
//
// <td>SaaS平台企业员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 
//
// ### 4. 签署方参数差异
//
// - 根据签署方的不同类型（第三方子企业的员工、个人、SaaS平台企业员工），传递的参数也不同。具体参数的结构和要求可以参考开发者中心提供的 `FlowApproverInfo` 结构体说明。
//
// 
//
// ### 5. 合同额度的扣减与返还
//
// - **扣减时机**：合同一旦发起，相关的合同额度就会被扣减，合同组下面的每个合同都要扣减一个合同额度。
//
// - **返还条件**：只有在合同被撤销且没有任何签署方签署过，或者只有自动签署的情况下，合同额度才会被返还。
//
// - **不返还的情况**：如果合同已过期、被拒签、签署完成或已解除，合同额度将不会被返还。
//
// 
//
// ### 6. 静默（自动）签署的限制
//
// - 在使用静默（自动）签署功能时，合同签署方不能有填写控件。<font color="red">此接口静默签(企业自动签)能力为白名单功能</font>，使用前请联系对接的客户经理沟通。
//
// 
//
// ### 7.合同组暂不支持抄送功能
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowGroupByTemplatesWithContext(ctx context.Context, request *ChannelCreateFlowGroupByTemplatesRequest) (response *ChannelCreateFlowGroupByTemplatesResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowGroupByTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowGroupByTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowGroupByTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowRemindsRequest() (request *ChannelCreateFlowRemindsRequest) {
    request = &ChannelCreateFlowRemindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowReminds")
    
    
    return
}

func NewChannelCreateFlowRemindsResponse() (response *ChannelCreateFlowRemindsResponse) {
    response = &ChannelCreateFlowRemindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateFlowReminds
// 指定需要批量催办的签署流程ID，批量催办合同，最多100个。需要符合以下条件的合同才可被催办
//
// 1. 合同中当前状态为 **待签署** 的签署人是催办的对象
//
// 2. **每个合同只能催办一次**
//
// 
//
// **催办的效果**: 对方会收到如下的短信通知
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/3caf94b7f540fa5736270d38528d3a7b.png)
//
// 
//
// 
//
// **注**：`合同催办是白名单功能，请联系客户经理申请开白后使用`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowReminds(request *ChannelCreateFlowRemindsRequest) (response *ChannelCreateFlowRemindsResponse, err error) {
    return c.ChannelCreateFlowRemindsWithContext(context.Background(), request)
}

// ChannelCreateFlowReminds
// 指定需要批量催办的签署流程ID，批量催办合同，最多100个。需要符合以下条件的合同才可被催办
//
// 1. 合同中当前状态为 **待签署** 的签署人是催办的对象
//
// 2. **每个合同只能催办一次**
//
// 
//
// **催办的效果**: 对方会收到如下的短信通知
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/3caf94b7f540fa5736270d38528d3a7b.png)
//
// 
//
// 
//
// **注**：`合同催办是白名单功能，请联系客户经理申请开白后使用`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowRemindsWithContext(ctx context.Context, request *ChannelCreateFlowRemindsRequest) (response *ChannelCreateFlowRemindsResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowRemindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowReminds require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowRemindsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowSignReviewRequest() (request *ChannelCreateFlowSignReviewRequest) {
    request = &ChannelCreateFlowSignReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowSignReview")
    
    
    return
}

func NewChannelCreateFlowSignReviewResponse() (response *ChannelCreateFlowSignReviewResponse) {
    response = &ChannelCreateFlowSignReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateFlowSignReview
// 提交企业流程审批结果 
//
// **当前存在两种审核操作：**
//
// <ul>
//
// <li>签署审核
//
// <ul>
//
// <li>在通过接口<ul><li>CreateFlowsByTemplates</li><li>ChannelCreateFlowByFiles</li><li>ChannelCreateFlowGroupByTemplates</li><li>ChannelCreateFlowGroupByFiles</li><li>ChannelCreatePrepareFlow</li></ul> 发起签署流程时，通过指定NeedSignReview为true，则可以调用此接口，并指定operate=SignReview，以提交企业内部签署审批结果</li>
//
// <li>在通过接口<ul><li>CreateFlowsByTemplates</li><li>ChannelCreateFlowByFiles</li><li>ChannelCreateFlowGroupByTemplates</li><li>ChannelCreateFlowGroupByFiles</li></ul>发起签署流程时，通过指定签署人ApproverNeedSignReview为true，则可以调用此接口，并指定operate=SignReview，并指定RecipientId，以提交企业内部签署审批结果</li>
//
// </ul>
//
// </li>
//
// <li>发起审核
//
//  <ul>
//
// <li>通过接口ChannelCreatePrepareFlow指定发起后需要审核，那么可以调用此接口，并指定operate=CreateReview，以提交企业内部审批结果。可以多次提交审批结果，但一旦审批通过，后续提交的结果将无效
//
// </li>
//
// </ul>
//
// </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateFlowSignReview(request *ChannelCreateFlowSignReviewRequest) (response *ChannelCreateFlowSignReviewResponse, err error) {
    return c.ChannelCreateFlowSignReviewWithContext(context.Background(), request)
}

// ChannelCreateFlowSignReview
// 提交企业流程审批结果 
//
// **当前存在两种审核操作：**
//
// <ul>
//
// <li>签署审核
//
// <ul>
//
// <li>在通过接口<ul><li>CreateFlowsByTemplates</li><li>ChannelCreateFlowByFiles</li><li>ChannelCreateFlowGroupByTemplates</li><li>ChannelCreateFlowGroupByFiles</li><li>ChannelCreatePrepareFlow</li></ul> 发起签署流程时，通过指定NeedSignReview为true，则可以调用此接口，并指定operate=SignReview，以提交企业内部签署审批结果</li>
//
// <li>在通过接口<ul><li>CreateFlowsByTemplates</li><li>ChannelCreateFlowByFiles</li><li>ChannelCreateFlowGroupByTemplates</li><li>ChannelCreateFlowGroupByFiles</li></ul>发起签署流程时，通过指定签署人ApproverNeedSignReview为true，则可以调用此接口，并指定operate=SignReview，并指定RecipientId，以提交企业内部签署审批结果</li>
//
// </ul>
//
// </li>
//
// <li>发起审核
//
//  <ul>
//
// <li>通过接口ChannelCreatePrepareFlow指定发起后需要审核，那么可以调用此接口，并指定operate=CreateReview，以提交企业内部审批结果。可以多次提交审批结果，但一旦审批通过，后续提交的结果将无效
//
// </li>
//
// </ul>
//
// </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateFlowSignReviewWithContext(ctx context.Context, request *ChannelCreateFlowSignReviewRequest) (response *ChannelCreateFlowSignReviewResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowSignReviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowSignReview require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowSignReviewResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateFlowSignUrlRequest() (request *ChannelCreateFlowSignUrlRequest) {
    request = &ChannelCreateFlowSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateFlowSignUrl")
    
    
    return
}

func NewChannelCreateFlowSignUrlResponse() (response *ChannelCreateFlowSignUrlResponse) {
    response = &ChannelCreateFlowSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateFlowSignUrl
// 该接口用于发起合同后，生成用户的签署链接 <br/>
//
// 
//
// **注意**
//
// 1. 该签署**链接有效期为30分钟**，过期后将失效，如需签署可重新创建签署链接 。
//
// 2. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。配置方式请参考：<a href="https://qian.tencent.com/developers/company/openqianh5/">跳转电子签H5</a>。
//
// 如需跳转到小程序的实现，参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 3. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowSignUrl(request *ChannelCreateFlowSignUrlRequest) (response *ChannelCreateFlowSignUrlResponse, err error) {
    return c.ChannelCreateFlowSignUrlWithContext(context.Background(), request)
}

// ChannelCreateFlowSignUrl
// 该接口用于发起合同后，生成用户的签署链接 <br/>
//
// 
//
// **注意**
//
// 1. 该签署**链接有效期为30分钟**，过期后将失效，如需签署可重新创建签署链接 。
//
// 2. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。配置方式请参考：<a href="https://qian.tencent.com/developers/company/openqianh5/">跳转电子签H5</a>。
//
// 如需跳转到小程序的实现，参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 3. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateFlowSignUrlWithContext(ctx context.Context, request *ChannelCreateFlowSignUrlRequest) (response *ChannelCreateFlowSignUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateFlowSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateFlowSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateFlowSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateMultiFlowSignQRCodeRequest() (request *ChannelCreateMultiFlowSignQRCodeRequest) {
    request = &ChannelCreateMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateMultiFlowSignQRCode")
    
    
    return
}

func NewChannelCreateMultiFlowSignQRCodeResponse() (response *ChannelCreateMultiFlowSignQRCodeResponse) {
    response = &ChannelCreateMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateMultiFlowSignQRCode
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多签签署码。 
//
// 
//
// **适用场景**:
//
// 签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 
//
// **注意**:
//
// 1. 本接口适用于**发起方没有填写控件的 B2C或者单C模板**,  若是B2C模板,还要满足以下任意一个条件
//
//     - 模板中配置的签署顺序是无序
//
//     - B端企业的签署方式是静默签署
//
//     - B端企业是非首位签署
//
// 2. 通过扫描一码多签签署码发起的合同，合同涉及到的回调消息可参考文档[合同发起及签署相关回调
//
// ]( https://qian.tencent.com/developers/partner/callback_types_contracts_sign)
//
// 3. 用户通过扫描一码多签签署码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档[签署二维码相关回调](https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83)
//
// 
//
// 签署码的样式如下图:
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/27317cf5aacb094fb1dc6f94179a5148.png )
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateMultiFlowSignQRCode(request *ChannelCreateMultiFlowSignQRCodeRequest) (response *ChannelCreateMultiFlowSignQRCodeResponse, err error) {
    return c.ChannelCreateMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// ChannelCreateMultiFlowSignQRCode
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多签签署码。 
//
// 
//
// **适用场景**:
//
// 签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 
//
// **注意**:
//
// 1. 本接口适用于**发起方没有填写控件的 B2C或者单C模板**,  若是B2C模板,还要满足以下任意一个条件
//
//     - 模板中配置的签署顺序是无序
//
//     - B端企业的签署方式是静默签署
//
//     - B端企业是非首位签署
//
// 2. 通过扫描一码多签签署码发起的合同，合同涉及到的回调消息可参考文档[合同发起及签署相关回调
//
// ]( https://qian.tencent.com/developers/partner/callback_types_contracts_sign)
//
// 3. 用户通过扫描一码多签签署码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档[签署二维码相关回调](https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83)
//
// 
//
// 签署码的样式如下图:
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/27317cf5aacb094fb1dc6f94179a5148.png )
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateMultiFlowSignQRCodeWithContext(ctx context.Context, request *ChannelCreateMultiFlowSignQRCodeRequest) (response *ChannelCreateMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewChannelCreateMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateOrganizationBatchSignUrlRequest() (request *ChannelCreateOrganizationBatchSignUrlRequest) {
    request = &ChannelCreateOrganizationBatchSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateOrganizationBatchSignUrl")
    
    
    return
}

func NewChannelCreateOrganizationBatchSignUrlResponse() (response *ChannelCreateOrganizationBatchSignUrlResponse) {
    response = &ChannelCreateOrganizationBatchSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateOrganizationBatchSignUrl
// 通过此接口，可以创建企业批量签署链接，员工只需点击链接即可跳转至控制台进行批量签署。
//
// 
//
// 注：
//
// - 员工必须在企业下完成实名认证，且需作为批量签署合同的签署方或者领取方。
//
// - 仅支持传入待签署或者待领取的合同，待填写暂不支持。
//
// - 员工批量签署，支持多种签名方式，包括手写签名、临摹签名、系统签名、个人印章、签批控件等。
//
// 
//
// 签署的嵌入页面长相如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/a4754bc835a3f837ddec1e28b02ed9c0.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateOrganizationBatchSignUrl(request *ChannelCreateOrganizationBatchSignUrlRequest) (response *ChannelCreateOrganizationBatchSignUrlResponse, err error) {
    return c.ChannelCreateOrganizationBatchSignUrlWithContext(context.Background(), request)
}

// ChannelCreateOrganizationBatchSignUrl
// 通过此接口，可以创建企业批量签署链接，员工只需点击链接即可跳转至控制台进行批量签署。
//
// 
//
// 注：
//
// - 员工必须在企业下完成实名认证，且需作为批量签署合同的签署方或者领取方。
//
// - 仅支持传入待签署或者待领取的合同，待填写暂不支持。
//
// - 员工批量签署，支持多种签名方式，包括手写签名、临摹签名、系统签名、个人印章、签批控件等。
//
// 
//
// 签署的嵌入页面长相如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/a4754bc835a3f837ddec1e28b02ed9c0.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateOrganizationBatchSignUrlWithContext(ctx context.Context, request *ChannelCreateOrganizationBatchSignUrlRequest) (response *ChannelCreateOrganizationBatchSignUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateOrganizationBatchSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateOrganizationBatchSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateOrganizationBatchSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateOrganizationModifyQrCodeRequest() (request *ChannelCreateOrganizationModifyQrCodeRequest) {
    request = &ChannelCreateOrganizationModifyQrCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateOrganizationModifyQrCode")
    
    
    return
}

func NewChannelCreateOrganizationModifyQrCodeResponse() (response *ChannelCreateOrganizationModifyQrCodeResponse) {
    response = &ChannelCreateOrganizationModifyQrCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateOrganizationModifyQrCode
// 生成渠道子客编辑企业信息二维码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelCreateOrganizationModifyQrCode(request *ChannelCreateOrganizationModifyQrCodeRequest) (response *ChannelCreateOrganizationModifyQrCodeResponse, err error) {
    return c.ChannelCreateOrganizationModifyQrCodeWithContext(context.Background(), request)
}

// ChannelCreateOrganizationModifyQrCode
// 生成渠道子客编辑企业信息二维码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelCreateOrganizationModifyQrCodeWithContext(ctx context.Context, request *ChannelCreateOrganizationModifyQrCodeRequest) (response *ChannelCreateOrganizationModifyQrCodeResponse, err error) {
    if request == nil {
        request = NewChannelCreateOrganizationModifyQrCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateOrganizationModifyQrCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateOrganizationModifyQrCodeResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreatePrepareFlowRequest() (request *ChannelCreatePrepareFlowRequest) {
    request = &ChannelCreatePrepareFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreatePrepareFlow")
    
    
    return
}

func NewChannelCreatePrepareFlowResponse() (response *ChannelCreatePrepareFlowResponse) {
    response = &ChannelCreatePrepareFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreatePrepareFlow
// 通过此接口指定合同、签署人、填写控件等信息，生成嵌入式链接，此链接可以嵌入到其他网页或者直接打开，打开后进入发起页面。在此页面上，合同信息和签署人信息均不可更改。
//
// 
//
// 注意：
//
// 1. <font color="red">仅支持在PC浏览器</font>上进行操作和使用。
//
// 2. 在使用<font color="red">模板发起合同时，需指定RecipientId</font>以明确参与方在模板中所扮演的角色。
//
// 
//
// **嵌入式签署人-各种场景传参说明**:
//
// 
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>场景编号</th>
//
// <th>可作为签署方的类型</th>
//
// <th>签署方传参说明</th>
//
// </tr>
//
// </thead>
//
// 
//
// <tbody>
//
// <tr>
//
// <td>场景一</td>
//
// <td>第三方子企业员工</td>
//
// <td>OpenId、OrganizationName、OrganizationOpenId必传 ,ApproverType设置为0</td>
//
// </tr>
//
// <tr>
//
// <td>场景二</td>
//
// <td>SaaS平台企业员工</td>
//
// <td>Name、Mobile、OrganizationName必传，NotChannelOrganization=True。 ApproverType设置为0</td>
//
// </tr>
//
// <tr>
//
// <td>场景三</td>
//
// <td>个人/自然人</td>
//
// <td>Name、Mobile必传, ApproverType设置为1</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 
//
// 嵌入的页面样式如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/b2ae013fb4d747891dd3815bbe897208.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreatePrepareFlow(request *ChannelCreatePrepareFlowRequest) (response *ChannelCreatePrepareFlowResponse, err error) {
    return c.ChannelCreatePrepareFlowWithContext(context.Background(), request)
}

// ChannelCreatePrepareFlow
// 通过此接口指定合同、签署人、填写控件等信息，生成嵌入式链接，此链接可以嵌入到其他网页或者直接打开，打开后进入发起页面。在此页面上，合同信息和签署人信息均不可更改。
//
// 
//
// 注意：
//
// 1. <font color="red">仅支持在PC浏览器</font>上进行操作和使用。
//
// 2. 在使用<font color="red">模板发起合同时，需指定RecipientId</font>以明确参与方在模板中所扮演的角色。
//
// 
//
// **嵌入式签署人-各种场景传参说明**:
//
// 
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>场景编号</th>
//
// <th>可作为签署方的类型</th>
//
// <th>签署方传参说明</th>
//
// </tr>
//
// </thead>
//
// 
//
// <tbody>
//
// <tr>
//
// <td>场景一</td>
//
// <td>第三方子企业员工</td>
//
// <td>OpenId、OrganizationName、OrganizationOpenId必传 ,ApproverType设置为0</td>
//
// </tr>
//
// <tr>
//
// <td>场景二</td>
//
// <td>SaaS平台企业员工</td>
//
// <td>Name、Mobile、OrganizationName必传，NotChannelOrganization=True。 ApproverType设置为0</td>
//
// </tr>
//
// <tr>
//
// <td>场景三</td>
//
// <td>个人/自然人</td>
//
// <td>Name、Mobile必传, ApproverType设置为1</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 
//
// 嵌入的页面样式如下：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/b2ae013fb4d747891dd3815bbe897208.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreatePrepareFlowWithContext(ctx context.Context, request *ChannelCreatePrepareFlowRequest) (response *ChannelCreatePrepareFlowResponse, err error) {
    if request == nil {
        request = NewChannelCreatePrepareFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreatePrepareFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreatePrepareFlowResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreatePreparedPersonalEsignRequest() (request *ChannelCreatePreparedPersonalEsignRequest) {
    request = &ChannelCreatePreparedPersonalEsignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreatePreparedPersonalEsign")
    
    
    return
}

func NewChannelCreatePreparedPersonalEsignResponse() (response *ChannelCreatePreparedPersonalEsignResponse) {
    response = &ChannelCreatePreparedPersonalEsignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreatePreparedPersonalEsign
// 本接口（ChannelCreatePreparedPersonalEsign）用于创建导入个人印章（处方单场景专用，使用此接口请与客户经理确认）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ChannelCreatePreparedPersonalEsign(request *ChannelCreatePreparedPersonalEsignRequest) (response *ChannelCreatePreparedPersonalEsignResponse, err error) {
    return c.ChannelCreatePreparedPersonalEsignWithContext(context.Background(), request)
}

// ChannelCreatePreparedPersonalEsign
// 本接口（ChannelCreatePreparedPersonalEsign）用于创建导入个人印章（处方单场景专用，使用此接口请与客户经理确认）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ChannelCreatePreparedPersonalEsignWithContext(ctx context.Context, request *ChannelCreatePreparedPersonalEsignRequest) (response *ChannelCreatePreparedPersonalEsignResponse, err error) {
    if request == nil {
        request = NewChannelCreatePreparedPersonalEsignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreatePreparedPersonalEsign require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreatePreparedPersonalEsignResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateReleaseFlowRequest() (request *ChannelCreateReleaseFlowRequest) {
    request = &ChannelCreateReleaseFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateReleaseFlow")
    
    
    return
}

func NewChannelCreateReleaseFlowResponse() (response *ChannelCreateReleaseFlowResponse) {
    response = &ChannelCreateReleaseFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateReleaseFlow
// 发起解除协议的主要应用场景为：基于一份已经签署的合同(签署流程)，进行解除操作。
//
// 解除协议的模板是官方提供，经过提供法务审核，暂不支持自定义。具体用法可以参考文档[合同解除](https://qian.tencent.com/developers/partner/flow_release)。
//
// 
//
// 注意：
//
// <ul><li><code>原合同必须签署完</code>成后才能发起解除协议。</li>
//
// <li>只有原合同企业类型的参与人才能发起解除协议，<code>个人参与方不能发起解除协议</code>。</li>
//
// <li>原合同个人类型参与人必须是解除协议的参与人，<code>不能更换其他第三方个人</code>参与解除协议。</li>
//
// <li>如果原合同企业参与人无法参与解除协议，可以指定同企业具有同等权限的<code>企业员工代为处理</code>。</li>
//
// <li>发起解除协议同发起其他企业合同一样，也会参与合同<code>扣费</code>，扣费标准同其他类型合同。</li>
//
// <li>在解除协议签署完毕后，原合同及解除协议均变为已解除状态。</li>
//
// <li>非原合同企业参与人发起解除协议时，需要有<code>解除合同的权限</code>。</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateReleaseFlow(request *ChannelCreateReleaseFlowRequest) (response *ChannelCreateReleaseFlowResponse, err error) {
    return c.ChannelCreateReleaseFlowWithContext(context.Background(), request)
}

// ChannelCreateReleaseFlow
// 发起解除协议的主要应用场景为：基于一份已经签署的合同(签署流程)，进行解除操作。
//
// 解除协议的模板是官方提供，经过提供法务审核，暂不支持自定义。具体用法可以参考文档[合同解除](https://qian.tencent.com/developers/partner/flow_release)。
//
// 
//
// 注意：
//
// <ul><li><code>原合同必须签署完</code>成后才能发起解除协议。</li>
//
// <li>只有原合同企业类型的参与人才能发起解除协议，<code>个人参与方不能发起解除协议</code>。</li>
//
// <li>原合同个人类型参与人必须是解除协议的参与人，<code>不能更换其他第三方个人</code>参与解除协议。</li>
//
// <li>如果原合同企业参与人无法参与解除协议，可以指定同企业具有同等权限的<code>企业员工代为处理</code>。</li>
//
// <li>发起解除协议同发起其他企业合同一样，也会参与合同<code>扣费</code>，扣费标准同其他类型合同。</li>
//
// <li>在解除协议签署完毕后，原合同及解除协议均变为已解除状态。</li>
//
// <li>非原合同企业参与人发起解除协议时，需要有<code>解除合同的权限</code>。</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateReleaseFlowWithContext(ctx context.Context, request *ChannelCreateReleaseFlowRequest) (response *ChannelCreateReleaseFlowResponse, err error) {
    if request == nil {
        request = NewChannelCreateReleaseFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateReleaseFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateReleaseFlowResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateRoleRequest() (request *ChannelCreateRoleRequest) {
    request = &ChannelCreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateRole")
    
    
    return
}

func NewChannelCreateRoleResponse() (response *ChannelCreateRoleResponse) {
    response = &ChannelCreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateRole
// 此接口（ChannelCreateRole）用来创建企业自定义角色。
//
// 
//
// 适用场景1：创建当前企业的自定义角色，并且创建时不进行权限的设置（PermissionGroups 参数不传），角色中的权限内容可通过接口 ChannelModifyRole 完成更新。
//
// 
//
// 适用场景2：创建当前企业的自定义角色，并且创建时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考[查询角色列表接口](https://qian.tencent.com/developers/partnerApis/accounts/ChannelDescribeRoles) 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ChannelCreateRole(request *ChannelCreateRoleRequest) (response *ChannelCreateRoleResponse, err error) {
    return c.ChannelCreateRoleWithContext(context.Background(), request)
}

// ChannelCreateRole
// 此接口（ChannelCreateRole）用来创建企业自定义角色。
//
// 
//
// 适用场景1：创建当前企业的自定义角色，并且创建时不进行权限的设置（PermissionGroups 参数不传），角色中的权限内容可通过接口 ChannelModifyRole 完成更新。
//
// 
//
// 适用场景2：创建当前企业的自定义角色，并且创建时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考[查询角色列表接口](https://qian.tencent.com/developers/partnerApis/accounts/ChannelDescribeRoles) 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ChannelCreateRoleWithContext(ctx context.Context, request *ChannelCreateRoleRequest) (response *ChannelCreateRoleResponse, err error) {
    if request == nil {
        request = NewChannelCreateRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateRoleResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateSealPolicyRequest() (request *ChannelCreateSealPolicyRequest) {
    request = &ChannelCreateSealPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateSealPolicy")
    
    
    return
}

func NewChannelCreateSealPolicyResponse() (response *ChannelCreateSealPolicyResponse) {
    response = &ChannelCreateSealPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateSealPolicy
// 将指定印章授权给第三方平台子客企业下的某些员工
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelCreateSealPolicy(request *ChannelCreateSealPolicyRequest) (response *ChannelCreateSealPolicyResponse, err error) {
    return c.ChannelCreateSealPolicyWithContext(context.Background(), request)
}

// ChannelCreateSealPolicy
// 将指定印章授权给第三方平台子客企业下的某些员工
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelCreateSealPolicyWithContext(ctx context.Context, request *ChannelCreateSealPolicyRequest) (response *ChannelCreateSealPolicyResponse, err error) {
    if request == nil {
        request = NewChannelCreateSealPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateSealPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateSealPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateUserAutoSignEnableUrlRequest() (request *ChannelCreateUserAutoSignEnableUrlRequest) {
    request = &ChannelCreateUserAutoSignEnableUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateUserAutoSignEnableUrl")
    
    
    return
}

func NewChannelCreateUserAutoSignEnableUrlResponse() (response *ChannelCreateUserAutoSignEnableUrlResponse) {
    response = &ChannelCreateUserAutoSignEnableUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateUserAutoSignEnableUrl
// 获取个人用户自动签的开通链接。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
func (c *Client) ChannelCreateUserAutoSignEnableUrl(request *ChannelCreateUserAutoSignEnableUrlRequest) (response *ChannelCreateUserAutoSignEnableUrlResponse, err error) {
    return c.ChannelCreateUserAutoSignEnableUrlWithContext(context.Background(), request)
}

// ChannelCreateUserAutoSignEnableUrl
// 获取个人用户自动签的开通链接。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
func (c *Client) ChannelCreateUserAutoSignEnableUrlWithContext(ctx context.Context, request *ChannelCreateUserAutoSignEnableUrlRequest) (response *ChannelCreateUserAutoSignEnableUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateUserAutoSignEnableUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateUserAutoSignEnableUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateUserAutoSignEnableUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateUserAutoSignSealUrlRequest() (request *ChannelCreateUserAutoSignSealUrlRequest) {
    request = &ChannelCreateUserAutoSignSealUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateUserAutoSignSealUrl")
    
    
    return
}

func NewChannelCreateUserAutoSignSealUrlResponse() (response *ChannelCreateUserAutoSignSealUrlResponse) {
    response = &ChannelCreateUserAutoSignSealUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateUserAutoSignSealUrl
// 获取设置自动签印章小程序链接。
//
// 
//
// 注意：
//
// <ul><li>需要<code>企业开通自动签</code>后使用。</li>
//
// <li>仅支持<code>已经开通了自动签的个人</code>更换自动签印章。</li>
//
// <li>链接有效期默认7天，<code>最多30天</code>。</li>
//
// <li>该接口的链接适用于<code>小程序</code>端。</li>
//
// <li>该接口不会扣除您的合同套餐，暂不参与计费。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateUserAutoSignSealUrl(request *ChannelCreateUserAutoSignSealUrlRequest) (response *ChannelCreateUserAutoSignSealUrlResponse, err error) {
    return c.ChannelCreateUserAutoSignSealUrlWithContext(context.Background(), request)
}

// ChannelCreateUserAutoSignSealUrl
// 获取设置自动签印章小程序链接。
//
// 
//
// 注意：
//
// <ul><li>需要<code>企业开通自动签</code>后使用。</li>
//
// <li>仅支持<code>已经开通了自动签的个人</code>更换自动签印章。</li>
//
// <li>链接有效期默认7天，<code>最多30天</code>。</li>
//
// <li>该接口的链接适用于<code>小程序</code>端。</li>
//
// <li>该接口不会扣除您的合同套餐，暂不参与计费。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateUserAutoSignSealUrlWithContext(ctx context.Context, request *ChannelCreateUserAutoSignSealUrlRequest) (response *ChannelCreateUserAutoSignSealUrlResponse, err error) {
    if request == nil {
        request = NewChannelCreateUserAutoSignSealUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateUserAutoSignSealUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateUserAutoSignSealUrlResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateUserRolesRequest() (request *ChannelCreateUserRolesRequest) {
    request = &ChannelCreateUserRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateUserRoles")
    
    
    return
}

func NewChannelCreateUserRolesResponse() (response *ChannelCreateUserRolesResponse) {
    response = &ChannelCreateUserRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateUserRoles
// 使用此接口，用来绑定企业实名员工的角色，
//
// 支持以电子签userId、客户系统openId两种方式进行绑定。
//
// 
//
// 对应控制台的操作如下图
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/5b41194d3cb3f2058ec0ba0fb5ebc6a6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateUserRoles(request *ChannelCreateUserRolesRequest) (response *ChannelCreateUserRolesResponse, err error) {
    return c.ChannelCreateUserRolesWithContext(context.Background(), request)
}

// ChannelCreateUserRoles
// 使用此接口，用来绑定企业实名员工的角色，
//
// 支持以电子签userId、客户系统openId两种方式进行绑定。
//
// 
//
// 对应控制台的操作如下图
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/5b41194d3cb3f2058ec0ba0fb5ebc6a6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateUserRolesWithContext(ctx context.Context, request *ChannelCreateUserRolesRequest) (response *ChannelCreateUserRolesResponse, err error) {
    if request == nil {
        request = NewChannelCreateUserRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateUserRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateUserRolesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelCreateWebThemeConfigRequest() (request *ChannelCreateWebThemeConfigRequest) {
    request = &ChannelCreateWebThemeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelCreateWebThemeConfig")
    
    
    return
}

func NewChannelCreateWebThemeConfigResponse() (response *ChannelCreateWebThemeConfigResponse) {
    response = &ChannelCreateWebThemeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelCreateWebThemeConfig
// 用来创建嵌入式页面个性化主题配置（例如是否展示电子签logo、定义主题色等），该接口配合其他所有可嵌入页面接口使用
//
// 创建配置对当前第三方应用全局生效，如果多次调用，会以最后一次的配置为准
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateWebThemeConfig(request *ChannelCreateWebThemeConfigRequest) (response *ChannelCreateWebThemeConfigResponse, err error) {
    return c.ChannelCreateWebThemeConfigWithContext(context.Background(), request)
}

// ChannelCreateWebThemeConfig
// 用来创建嵌入式页面个性化主题配置（例如是否展示电子签logo、定义主题色等），该接口配合其他所有可嵌入页面接口使用
//
// 创建配置对当前第三方应用全局生效，如果多次调用，会以最后一次的配置为准
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateWebThemeConfigWithContext(ctx context.Context, request *ChannelCreateWebThemeConfigRequest) (response *ChannelCreateWebThemeConfigResponse, err error) {
    if request == nil {
        request = NewChannelCreateWebThemeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelCreateWebThemeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelCreateWebThemeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDeleteRoleRequest() (request *ChannelDeleteRoleRequest) {
    request = &ChannelDeleteRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDeleteRole")
    
    
    return
}

func NewChannelDeleteRoleResponse() (response *ChannelDeleteRoleResponse) {
    response = &ChannelDeleteRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDeleteRole
// 此接口（ChannelDeleteRole）用来删除企业自定义角色。
//
// 
//
// 注：**系统角色不可删除。**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ChannelDeleteRole(request *ChannelDeleteRoleRequest) (response *ChannelDeleteRoleResponse, err error) {
    return c.ChannelDeleteRoleWithContext(context.Background(), request)
}

// ChannelDeleteRole
// 此接口（ChannelDeleteRole）用来删除企业自定义角色。
//
// 
//
// 注：**系统角色不可删除。**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ChannelDeleteRoleWithContext(ctx context.Context, request *ChannelDeleteRoleRequest) (response *ChannelDeleteRoleResponse, err error) {
    if request == nil {
        request = NewChannelDeleteRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDeleteRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDeleteRoleResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDeleteRoleUsersRequest() (request *ChannelDeleteRoleUsersRequest) {
    request = &ChannelDeleteRoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDeleteRoleUsers")
    
    
    return
}

func NewChannelDeleteRoleUsersResponse() (response *ChannelDeleteRoleUsersResponse) {
    response = &ChannelDeleteRoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDeleteRoleUsers
// 通过此接口，删除员工绑定的角色，支持以电子签userId、客户系统userId两种方式调用。
//
// 
//
// 对应控制台的操作如下图
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/5b41194d3cb3f2058ec0ba0fb5ebc6a6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLEID = "InvalidParameter.RoleId"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
func (c *Client) ChannelDeleteRoleUsers(request *ChannelDeleteRoleUsersRequest) (response *ChannelDeleteRoleUsersResponse, err error) {
    return c.ChannelDeleteRoleUsersWithContext(context.Background(), request)
}

// ChannelDeleteRoleUsers
// 通过此接口，删除员工绑定的角色，支持以电子签userId、客户系统userId两种方式调用。
//
// 
//
// 对应控制台的操作如下图
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/5b41194d3cb3f2058ec0ba0fb5ebc6a6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLEID = "InvalidParameter.RoleId"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
func (c *Client) ChannelDeleteRoleUsersWithContext(ctx context.Context, request *ChannelDeleteRoleUsersRequest) (response *ChannelDeleteRoleUsersResponse, err error) {
    if request == nil {
        request = NewChannelDeleteRoleUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDeleteRoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDeleteRoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDeleteSealPoliciesRequest() (request *ChannelDeleteSealPoliciesRequest) {
    request = &ChannelDeleteSealPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDeleteSealPolicies")
    
    
    return
}

func NewChannelDeleteSealPoliciesResponse() (response *ChannelDeleteSealPoliciesResponse) {
    response = &ChannelDeleteSealPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDeleteSealPolicies
// 此接口（ChannelDeleteSealPolicies）用于删除已指定员工印章授权信息，删除员工的印章授权后，该员工使用印章进行盖章时，将需要提交印章授权申请且通过审核后才能使用该印章进行签署。
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelDeleteSealPolicies(request *ChannelDeleteSealPoliciesRequest) (response *ChannelDeleteSealPoliciesResponse, err error) {
    return c.ChannelDeleteSealPoliciesWithContext(context.Background(), request)
}

// ChannelDeleteSealPolicies
// 此接口（ChannelDeleteSealPolicies）用于删除已指定员工印章授权信息，删除员工的印章授权后，该员工使用印章进行盖章时，将需要提交印章授权申请且通过审核后才能使用该印章进行签署。
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelDeleteSealPoliciesWithContext(ctx context.Context, request *ChannelDeleteSealPoliciesRequest) (response *ChannelDeleteSealPoliciesResponse, err error) {
    if request == nil {
        request = NewChannelDeleteSealPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDeleteSealPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDeleteSealPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeAccountBillDetailRequest() (request *ChannelDescribeAccountBillDetailRequest) {
    request = &ChannelDescribeAccountBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeAccountBillDetail")
    
    
    return
}

func NewChannelDescribeAccountBillDetailResponse() (response *ChannelDescribeAccountBillDetailResponse) {
    response = &ChannelDescribeAccountBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeAccountBillDetail
// 通过此接口（ChannelDescribeAccountBillDetail）查询该第三方平台子客账号计费详情。
//
// <ul>
//
// <li>对于渠道客户企业的查询，通过指定渠道企业的唯一标识(Agent.ProxyOrganizationId)来查询子客账号消耗详情</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDQUOTATYPE = "InvalidParameterValue.InvalidQuotaType"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  RESOURCENOTFOUND_NOTEXISTAPPLICATION = "ResourceNotFound.NotExistApplication"
func (c *Client) ChannelDescribeAccountBillDetail(request *ChannelDescribeAccountBillDetailRequest) (response *ChannelDescribeAccountBillDetailResponse, err error) {
    return c.ChannelDescribeAccountBillDetailWithContext(context.Background(), request)
}

// ChannelDescribeAccountBillDetail
// 通过此接口（ChannelDescribeAccountBillDetail）查询该第三方平台子客账号计费详情。
//
// <ul>
//
// <li>对于渠道客户企业的查询，通过指定渠道企业的唯一标识(Agent.ProxyOrganizationId)来查询子客账号消耗详情</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDQUOTATYPE = "InvalidParameterValue.InvalidQuotaType"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  RESOURCENOTFOUND_NOTEXISTAPPLICATION = "ResourceNotFound.NotExistApplication"
func (c *Client) ChannelDescribeAccountBillDetailWithContext(ctx context.Context, request *ChannelDescribeAccountBillDetailRequest) (response *ChannelDescribeAccountBillDetailResponse, err error) {
    if request == nil {
        request = NewChannelDescribeAccountBillDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeAccountBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeAccountBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeBillUsageDetailRequest() (request *ChannelDescribeBillUsageDetailRequest) {
    request = &ChannelDescribeBillUsageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeBillUsageDetail")
    
    
    return
}

func NewChannelDescribeBillUsageDetailResponse() (response *ChannelDescribeBillUsageDetailResponse) {
    response = &ChannelDescribeBillUsageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeBillUsageDetail
// 通过此接口（ChannelDescribeBillUsageDetail）查询该第三方平台子客企业的套餐消耗详情。可以支持单个子客和整个应用下所有子客的查询。
//
// <ul>
//
// <li>对于单个子客企业的查询，通过指定子客的唯一标识(Agent.ProxyOrganizationOpenId)来查询该子客消耗详情</li>
//
// <li>对于整个应用下所有企业的查询，不需要指定子客的唯一标识，只需要传入渠道应用标识(Agent.AppId)直接查询整个应用下所有子客企业消耗详情</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDQUOTATYPE = "InvalidParameterValue.InvalidQuotaType"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  RESOURCENOTFOUND_NOTEXISTAPPLICATION = "ResourceNotFound.NotExistApplication"
func (c *Client) ChannelDescribeBillUsageDetail(request *ChannelDescribeBillUsageDetailRequest) (response *ChannelDescribeBillUsageDetailResponse, err error) {
    return c.ChannelDescribeBillUsageDetailWithContext(context.Background(), request)
}

// ChannelDescribeBillUsageDetail
// 通过此接口（ChannelDescribeBillUsageDetail）查询该第三方平台子客企业的套餐消耗详情。可以支持单个子客和整个应用下所有子客的查询。
//
// <ul>
//
// <li>对于单个子客企业的查询，通过指定子客的唯一标识(Agent.ProxyOrganizationOpenId)来查询该子客消耗详情</li>
//
// <li>对于整个应用下所有企业的查询，不需要指定子客的唯一标识，只需要传入渠道应用标识(Agent.AppId)直接查询整个应用下所有子客企业消耗详情</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDQUOTATYPE = "InvalidParameterValue.InvalidQuotaType"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  RESOURCENOTFOUND_NOTEXISTAPPLICATION = "ResourceNotFound.NotExistApplication"
func (c *Client) ChannelDescribeBillUsageDetailWithContext(ctx context.Context, request *ChannelDescribeBillUsageDetailRequest) (response *ChannelDescribeBillUsageDetailResponse, err error) {
    if request == nil {
        request = NewChannelDescribeBillUsageDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeBillUsageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeBillUsageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeEmployeesRequest() (request *ChannelDescribeEmployeesRequest) {
    request = &ChannelDescribeEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeEmployees")
    
    
    return
}

func NewChannelDescribeEmployeesResponse() (response *ChannelDescribeEmployeesResponse) {
    response = &ChannelDescribeEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeEmployees
// 获取企业员工信息, 可以获取员工的名字,OpenId,UserId和简述的角色等信息，支持设置过滤条件以筛选员工查询结果。
//
// 
//
// **注**:通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/SyncProxyOrganizationOperators" target="_blank">企业员工新增或离职</a>接口增加的新员工或者离职的员工也会在列表中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_DEPARTUSERID = "InvalidParameter.DepartUserId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeEmployees(request *ChannelDescribeEmployeesRequest) (response *ChannelDescribeEmployeesResponse, err error) {
    return c.ChannelDescribeEmployeesWithContext(context.Background(), request)
}

// ChannelDescribeEmployees
// 获取企业员工信息, 可以获取员工的名字,OpenId,UserId和简述的角色等信息，支持设置过滤条件以筛选员工查询结果。
//
// 
//
// **注**:通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/SyncProxyOrganizationOperators" target="_blank">企业员工新增或离职</a>接口增加的新员工或者离职的员工也会在列表中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_DEPARTUSERID = "InvalidParameter.DepartUserId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeEmployeesWithContext(ctx context.Context, request *ChannelDescribeEmployeesRequest) (response *ChannelDescribeEmployeesResponse, err error) {
    if request == nil {
        request = NewChannelDescribeEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeFlowComponentsRequest() (request *ChannelDescribeFlowComponentsRequest) {
    request = &ChannelDescribeFlowComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeFlowComponents")
    
    
    return
}

func NewChannelDescribeFlowComponentsResponse() (response *ChannelDescribeFlowComponentsResponse) {
    response = &ChannelDescribeFlowComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeFlowComponents
// 用于获取合同中填写控件填写状态和填写的内容。 
//
// 
//
// **注意**: `附件控件不会出现在结果列表中`
//
// 
//
// 
//
// **授权**:   
//
// 此接口需要授权,  有两种开通权限的途径
//
// 
//
// **第一种**:   需第三方应用的子企业登录控制台进行授权,  授权在**企业中心**的**授权管理**区域,  界面如下图
//
// 授权过程需要**子企业超管**扫描跳转到电子签小程序签署<<渠道端下载渠道子客合同功能授权委托书>>
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/8b483dfebdeafac85051279406944048.png)
//
// 
//
// **第二种**: 第三方应用的配置接口打开全第三个应用下的所有自己起开通, 需要**渠道方企业的超管**扫描二维码跳转到电子签小程序签署 <<渠道端下载渠道子客合同功能开通知情同意书>>
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/238979ef51dd381ccbdbc755a593debc/channel_DescribeResourceUrlsByFlows_appilications2.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
func (c *Client) ChannelDescribeFlowComponents(request *ChannelDescribeFlowComponentsRequest) (response *ChannelDescribeFlowComponentsResponse, err error) {
    return c.ChannelDescribeFlowComponentsWithContext(context.Background(), request)
}

// ChannelDescribeFlowComponents
// 用于获取合同中填写控件填写状态和填写的内容。 
//
// 
//
// **注意**: `附件控件不会出现在结果列表中`
//
// 
//
// 
//
// **授权**:   
//
// 此接口需要授权,  有两种开通权限的途径
//
// 
//
// **第一种**:   需第三方应用的子企业登录控制台进行授权,  授权在**企业中心**的**授权管理**区域,  界面如下图
//
// 授权过程需要**子企业超管**扫描跳转到电子签小程序签署<<渠道端下载渠道子客合同功能授权委托书>>
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/8b483dfebdeafac85051279406944048.png)
//
// 
//
// **第二种**: 第三方应用的配置接口打开全第三个应用下的所有自己起开通, 需要**渠道方企业的超管**扫描二维码跳转到电子签小程序签署 <<渠道端下载渠道子客合同功能开通知情同意书>>
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/238979ef51dd381ccbdbc755a593debc/channel_DescribeResourceUrlsByFlows_appilications2.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
func (c *Client) ChannelDescribeFlowComponentsWithContext(ctx context.Context, request *ChannelDescribeFlowComponentsRequest) (response *ChannelDescribeFlowComponentsResponse, err error) {
    if request == nil {
        request = NewChannelDescribeFlowComponentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeFlowComponents require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeFlowComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeOrganizationSealsRequest() (request *ChannelDescribeOrganizationSealsRequest) {
    request = &ChannelDescribeOrganizationSealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeOrganizationSeals")
    
    
    return
}

func NewChannelDescribeOrganizationSealsResponse() (response *ChannelDescribeOrganizationSealsResponse) {
    response = &ChannelDescribeOrganizationSealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeOrganizationSeals
// 此接口查询子企业电子印章。<br />
//
// 
//
// 注：
//
// 1. 此操作要求操作者具备<b>印章查询权限</b>（若调用者尚无此权限，请联系超级管理员前往Web控制台【组织管理】->【角色管理】添加相应权限）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeOrganizationSeals(request *ChannelDescribeOrganizationSealsRequest) (response *ChannelDescribeOrganizationSealsResponse, err error) {
    return c.ChannelDescribeOrganizationSealsWithContext(context.Background(), request)
}

// ChannelDescribeOrganizationSeals
// 此接口查询子企业电子印章。<br />
//
// 
//
// 注：
//
// 1. 此操作要求操作者具备<b>印章查询权限</b>（若调用者尚无此权限，请联系超级管理员前往Web控制台【组织管理】->【角色管理】添加相应权限）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeOrganizationSealsWithContext(ctx context.Context, request *ChannelDescribeOrganizationSealsRequest) (response *ChannelDescribeOrganizationSealsResponse, err error) {
    if request == nil {
        request = NewChannelDescribeOrganizationSealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeOrganizationSeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeOrganizationSealsResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeRolesRequest() (request *ChannelDescribeRolesRequest) {
    request = &ChannelDescribeRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeRoles")
    
    
    return
}

func NewChannelDescribeRolesResponse() (response *ChannelDescribeRolesResponse) {
    response = &ChannelDescribeRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeRoles
// 分页查询企业角色列表，法人的角色是系统保留角色，不会返回，按照角色创建时间升序排列。
//
// 
//
// 
//
// <font color="red">系统默认角色</font>说明可参考下表
//
// 
//
// | 角色名称| 建议授予对象 | 角色描述 |
//
// | --- | --- | --- |
//
// | **超级管理员** |电子签业务最高权限，可以授权给法务/企业法人/业务负责人等 | 所有功能和数据管理权限，只能设置一位超管。 |
//
// | **业务管理员**|IT 系统负责人，可以授权给CTO等 | 企业合同模块、印章模块、模板模块等全量功能及数据权限。 |
//
// | **经办人**|企业法务负责人等 | 发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看企业所有合同数据。 |
//
// | **业务员**|销售员、采购员 等| 发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看自己相关所有合同数据。 |
//
// 
//
// 附件：<a href="https://dyn.ess.tencent.cn/guide/apivideo/roles.xlsx" target="_blank">点击下载角色对应的权限点的excel文档</a>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeRoles(request *ChannelDescribeRolesRequest) (response *ChannelDescribeRolesResponse, err error) {
    return c.ChannelDescribeRolesWithContext(context.Background(), request)
}

// ChannelDescribeRoles
// 分页查询企业角色列表，法人的角色是系统保留角色，不会返回，按照角色创建时间升序排列。
//
// 
//
// 
//
// <font color="red">系统默认角色</font>说明可参考下表
//
// 
//
// | 角色名称| 建议授予对象 | 角色描述 |
//
// | --- | --- | --- |
//
// | **超级管理员** |电子签业务最高权限，可以授权给法务/企业法人/业务负责人等 | 所有功能和数据管理权限，只能设置一位超管。 |
//
// | **业务管理员**|IT 系统负责人，可以授权给CTO等 | 企业合同模块、印章模块、模板模块等全量功能及数据权限。 |
//
// | **经办人**|企业法务负责人等 | 发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看企业所有合同数据。 |
//
// | **业务员**|销售员、采购员 等| 发起合同、签署合同（含填写、拒签）、撤销合同、持有印章等权限能力，可查看自己相关所有合同数据。 |
//
// 
//
// 附件：<a href="https://dyn.ess.tencent.cn/guide/apivideo/roles.xlsx" target="_blank">点击下载角色对应的权限点的excel文档</a>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelDescribeRolesWithContext(ctx context.Context, request *ChannelDescribeRolesRequest) (response *ChannelDescribeRolesResponse, err error) {
    if request == nil {
        request = NewChannelDescribeRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeRolesResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeSignFaceVideoRequest() (request *ChannelDescribeSignFaceVideoRequest) {
    request = &ChannelDescribeSignFaceVideoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeSignFaceVideo")
    
    
    return
}

func NewChannelDescribeSignFaceVideoResponse() (response *ChannelDescribeSignFaceVideoResponse) {
    response = &ChannelDescribeSignFaceVideoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeSignFaceVideo
// 该接口用于在使用视频认证方式签署合同后，获取用户的签署人脸认证视频。
//
// 
//
// 1. 该接口**仅适用于在H5端签署**的合同，**在通过视频认证后**获取人脸图片。
//
// 2. 该接口**不支持小程序端**的签署人脸图片获取。
//
// 3. 请在**签署完成后的三天内**获取人脸图片，**过期后将无法获取**。
//
// 
//
// **注意：该接口需要开通白名单，请联系客户经理开通后使用。**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelDescribeSignFaceVideo(request *ChannelDescribeSignFaceVideoRequest) (response *ChannelDescribeSignFaceVideoResponse, err error) {
    return c.ChannelDescribeSignFaceVideoWithContext(context.Background(), request)
}

// ChannelDescribeSignFaceVideo
// 该接口用于在使用视频认证方式签署合同后，获取用户的签署人脸认证视频。
//
// 
//
// 1. 该接口**仅适用于在H5端签署**的合同，**在通过视频认证后**获取人脸图片。
//
// 2. 该接口**不支持小程序端**的签署人脸图片获取。
//
// 3. 请在**签署完成后的三天内**获取人脸图片，**过期后将无法获取**。
//
// 
//
// **注意：该接口需要开通白名单，请联系客户经理开通后使用。**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelDescribeSignFaceVideoWithContext(ctx context.Context, request *ChannelDescribeSignFaceVideoRequest) (response *ChannelDescribeSignFaceVideoResponse, err error) {
    if request == nil {
        request = NewChannelDescribeSignFaceVideoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeSignFaceVideo require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeSignFaceVideoResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDescribeUserAutoSignStatusRequest() (request *ChannelDescribeUserAutoSignStatusRequest) {
    request = &ChannelDescribeUserAutoSignStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDescribeUserAutoSignStatus")
    
    
    return
}

func NewChannelDescribeUserAutoSignStatusResponse() (response *ChannelDescribeUserAutoSignStatusResponse) {
    response = &ChannelDescribeUserAutoSignStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDescribeUserAutoSignStatus
// 通过此接口获取个人用户自动签的开通状态。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelDescribeUserAutoSignStatus(request *ChannelDescribeUserAutoSignStatusRequest) (response *ChannelDescribeUserAutoSignStatusResponse, err error) {
    return c.ChannelDescribeUserAutoSignStatusWithContext(context.Background(), request)
}

// ChannelDescribeUserAutoSignStatus
// 通过此接口获取个人用户自动签的开通状态。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelDescribeUserAutoSignStatusWithContext(ctx context.Context, request *ChannelDescribeUserAutoSignStatusRequest) (response *ChannelDescribeUserAutoSignStatusResponse, err error) {
    if request == nil {
        request = NewChannelDescribeUserAutoSignStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDescribeUserAutoSignStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDescribeUserAutoSignStatusResponse()
    err = c.Send(request, response)
    return
}

func NewChannelDisableUserAutoSignRequest() (request *ChannelDisableUserAutoSignRequest) {
    request = &ChannelDisableUserAutoSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelDisableUserAutoSign")
    
    
    return
}

func NewChannelDisableUserAutoSignResponse() (response *ChannelDisableUserAutoSignResponse) {
    response = &ChannelDisableUserAutoSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelDisableUserAutoSign
// 通过此接口可以关闭个人用户自动签功能。
//
// 无需对应的用户刷脸等方式同意即可关闭。
//
// 
//
// 注意: 
//
// 
//
// <ul><li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>如果此用户在开通时候绑定过个人自动签账号许可,  关闭此用户的自动签不会归还个人自动签账号许可的额度。</li></ul>
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelDisableUserAutoSign(request *ChannelDisableUserAutoSignRequest) (response *ChannelDisableUserAutoSignResponse, err error) {
    return c.ChannelDisableUserAutoSignWithContext(context.Background(), request)
}

// ChannelDisableUserAutoSign
// 通过此接口可以关闭个人用户自动签功能。
//
// 无需对应的用户刷脸等方式同意即可关闭。
//
// 
//
// 注意: 
//
// 
//
// <ul><li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>如果此用户在开通时候绑定过个人自动签账号许可,  关闭此用户的自动签不会归还个人自动签账号许可的额度。</li></ul>
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelDisableUserAutoSignWithContext(ctx context.Context, request *ChannelDisableUserAutoSignRequest) (response *ChannelDisableUserAutoSignResponse, err error) {
    if request == nil {
        request = NewChannelDisableUserAutoSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelDisableUserAutoSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelDisableUserAutoSignResponse()
    err = c.Send(request, response)
    return
}

func NewChannelGetTaskResultApiRequest() (request *ChannelGetTaskResultApiRequest) {
    request = &ChannelGetTaskResultApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelGetTaskResultApi")
    
    
    return
}

func NewChannelGetTaskResultApiResponse() (response *ChannelGetTaskResultApiResponse) {
    response = &ChannelGetTaskResultApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelGetTaskResultApi
// 此接口（ChannelGetTaskResultApi）用来查询转换任务的状态。如需发起转换任务，请使用<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行资源文件的转换操作<br />
//
// 前提条件：已调用 <a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行文件转换，并得到了返回的转换任务Id。<br />
//
// 
//
// 适用场景：已创建一个文件转换任务，想查询该文件转换任务的状态，或获取转换后的文件资源ID。<br />
//
// 注：
//
// 1. `大文件转换所需的时间可能会比较长。`
//
// 2. `本接口返回的文件资源ID就是PDF资源ID，可以直接用于【用PDF文件创建签署流程】接口发起合同。`
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelGetTaskResultApi(request *ChannelGetTaskResultApiRequest) (response *ChannelGetTaskResultApiResponse, err error) {
    return c.ChannelGetTaskResultApiWithContext(context.Background(), request)
}

// ChannelGetTaskResultApi
// 此接口（ChannelGetTaskResultApi）用来查询转换任务的状态。如需发起转换任务，请使用<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行资源文件的转换操作<br />
//
// 前提条件：已调用 <a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行文件转换，并得到了返回的转换任务Id。<br />
//
// 
//
// 适用场景：已创建一个文件转换任务，想查询该文件转换任务的状态，或获取转换后的文件资源ID。<br />
//
// 注：
//
// 1. `大文件转换所需的时间可能会比较长。`
//
// 2. `本接口返回的文件资源ID就是PDF资源ID，可以直接用于【用PDF文件创建签署流程】接口发起合同。`
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelGetTaskResultApiWithContext(ctx context.Context, request *ChannelGetTaskResultApiRequest) (response *ChannelGetTaskResultApiResponse, err error) {
    if request == nil {
        request = NewChannelGetTaskResultApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelGetTaskResultApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelGetTaskResultApiResponse()
    err = c.Send(request, response)
    return
}

func NewChannelModifyRoleRequest() (request *ChannelModifyRoleRequest) {
    request = &ChannelModifyRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelModifyRole")
    
    
    return
}

func NewChannelModifyRoleResponse() (response *ChannelModifyRoleResponse) {
    response = &ChannelModifyRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelModifyRole
// 此接口（ChannelModifyRole）用来更新企业自定义角色。
//
// 
//
// 适用场景1：更新当前企业的自定义角色的名称或描述等其他信息，更新时不进行权限的设置（PermissionGroups 参数不传）。
//
// 
//
// 适用场景2：更新当前企业的自定义角色的权限信息，更新时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考[查询角色列表接口](https://qian.tencent.com/developers/partnerApis/accounts/ChannelDescribeRoles) 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ChannelModifyRole(request *ChannelModifyRoleRequest) (response *ChannelModifyRoleResponse, err error) {
    return c.ChannelModifyRoleWithContext(context.Background(), request)
}

// ChannelModifyRole
// 此接口（ChannelModifyRole）用来更新企业自定义角色。
//
// 
//
// 适用场景1：更新当前企业的自定义角色的名称或描述等其他信息，更新时不进行权限的设置（PermissionGroups 参数不传）。
//
// 
//
// 适用场景2：更新当前企业的自定义角色的权限信息，更新时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考[查询角色列表接口](https://qian.tencent.com/developers/partnerApis/accounts/ChannelDescribeRoles) 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ChannelModifyRoleWithContext(ctx context.Context, request *ChannelModifyRoleRequest) (response *ChannelModifyRoleResponse, err error) {
    if request == nil {
        request = NewChannelModifyRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelModifyRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelModifyRoleResponse()
    err = c.Send(request, response)
    return
}

func NewChannelRenewAutoSignLicenseRequest() (request *ChannelRenewAutoSignLicenseRequest) {
    request = &ChannelRenewAutoSignLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelRenewAutoSignLicense")
    
    
    return
}

func NewChannelRenewAutoSignLicenseResponse() (response *ChannelRenewAutoSignLicenseResponse) {
    response = &ChannelRenewAutoSignLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelRenewAutoSignLicense
// 给医疗个人自动签许可续期。续期成功后，可对医疗自动签许可追加一年有效期，只可续期一次。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LICENSENOQUOTA = "FailedOperation.LicenseNoQuota"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelRenewAutoSignLicense(request *ChannelRenewAutoSignLicenseRequest) (response *ChannelRenewAutoSignLicenseResponse, err error) {
    return c.ChannelRenewAutoSignLicenseWithContext(context.Background(), request)
}

// ChannelRenewAutoSignLicense
// 给医疗个人自动签许可续期。续期成功后，可对医疗自动签许可追加一年有效期，只可续期一次。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LICENSENOQUOTA = "FailedOperation.LicenseNoQuota"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelRenewAutoSignLicenseWithContext(ctx context.Context, request *ChannelRenewAutoSignLicenseRequest) (response *ChannelRenewAutoSignLicenseResponse, err error) {
    if request == nil {
        request = NewChannelRenewAutoSignLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelRenewAutoSignLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelRenewAutoSignLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewChannelUpdateSealStatusRequest() (request *ChannelUpdateSealStatusRequest) {
    request = &ChannelUpdateSealStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelUpdateSealStatus")
    
    
    return
}

func NewChannelUpdateSealStatusResponse() (response *ChannelUpdateSealStatusResponse) {
    response = &ChannelUpdateSealStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelUpdateSealStatus
// 此接口（ChannelUpdateSealStatus）用于第三方应用平台为子客企业更新印章状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelUpdateSealStatus(request *ChannelUpdateSealStatusRequest) (response *ChannelUpdateSealStatusResponse, err error) {
    return c.ChannelUpdateSealStatusWithContext(context.Background(), request)
}

// ChannelUpdateSealStatus
// 此接口（ChannelUpdateSealStatus）用于第三方应用平台为子客企业更新印章状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelUpdateSealStatusWithContext(ctx context.Context, request *ChannelUpdateSealStatusRequest) (response *ChannelUpdateSealStatusResponse, err error) {
    if request == nil {
        request = NewChannelUpdateSealStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelUpdateSealStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelUpdateSealStatusResponse()
    err = c.Send(request, response)
    return
}

func NewChannelVerifyPdfRequest() (request *ChannelVerifyPdfRequest) {
    request = &ChannelVerifyPdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ChannelVerifyPdf")
    
    
    return
}

func NewChannelVerifyPdfResponse() (response *ChannelVerifyPdfResponse) {
    response = &ChannelVerifyPdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChannelVerifyPdf
// 对合同流程文件进行数字签名验证，判断数字签名是否有效，合同文件内容是否被篡改。
//
// 
//
// **补充**： 可以到控制台[合同验签](https://qian.tencent.com/verifySign)体验验签功能，界面如下
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/81c333ccb07f0c5fbaf840d9cee61333.png)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelVerifyPdf(request *ChannelVerifyPdfRequest) (response *ChannelVerifyPdfResponse, err error) {
    return c.ChannelVerifyPdfWithContext(context.Background(), request)
}

// ChannelVerifyPdf
// 对合同流程文件进行数字签名验证，判断数字签名是否有效，合同文件内容是否被篡改。
//
// 
//
// **补充**： 可以到控制台[合同验签](https://qian.tencent.com/verifySign)体验验签功能，界面如下
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/81c333ccb07f0c5fbaf840d9cee61333.png)
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelVerifyPdfWithContext(ctx context.Context, request *ChannelVerifyPdfRequest) (response *ChannelVerifyPdfResponse, err error) {
    if request == nil {
        request = NewChannelVerifyPdfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChannelVerifyPdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewChannelVerifyPdfResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchInitOrganizationUrlRequest() (request *CreateBatchInitOrganizationUrlRequest) {
    request = &CreateBatchInitOrganizationUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateBatchInitOrganizationUrl")
    
    
    return
}

func NewCreateBatchInitOrganizationUrlResponse() (response *CreateBatchInitOrganizationUrlResponse) {
    response = &CreateBatchInitOrganizationUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBatchInitOrganizationUrl
// 支持企业进行批量初始化操作：
//
// 
//
// 此接口存在以下限制：
//
// 1. 批量操作的企业需要已经完成电子签的认证流程。
//
// 2. 通过此接口生成的链接在小程序端进行操作时，操作人需要是<font  color="red">所有企业的超管或法人</font>。
//
// 3. 批量操作的企业，需要是本方第三方应用下的企业。
//
// 4. <font  color="red">操作链接过期时间默认为生成链接后7天。</font>
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBatchInitOrganizationUrl(request *CreateBatchInitOrganizationUrlRequest) (response *CreateBatchInitOrganizationUrlResponse, err error) {
    return c.CreateBatchInitOrganizationUrlWithContext(context.Background(), request)
}

// CreateBatchInitOrganizationUrl
// 支持企业进行批量初始化操作：
//
// 
//
// 此接口存在以下限制：
//
// 1. 批量操作的企业需要已经完成电子签的认证流程。
//
// 2. 通过此接口生成的链接在小程序端进行操作时，操作人需要是<font  color="red">所有企业的超管或法人</font>。
//
// 3. 批量操作的企业，需要是本方第三方应用下的企业。
//
// 4. <font  color="red">操作链接过期时间默认为生成链接后7天。</font>
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBatchInitOrganizationUrlWithContext(ctx context.Context, request *CreateBatchInitOrganizationUrlRequest) (response *CreateBatchInitOrganizationUrlResponse, err error) {
    if request == nil {
        request = NewCreateBatchInitOrganizationUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchInitOrganizationUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchInitOrganizationUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchOrganizationRegistrationTasksRequest() (request *CreateBatchOrganizationRegistrationTasksRequest) {
    request = &CreateBatchOrganizationRegistrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateBatchOrganizationRegistrationTasks")
    
    
    return
}

func NewCreateBatchOrganizationRegistrationTasksResponse() (response *CreateBatchOrganizationRegistrationTasksResponse) {
    response = &CreateBatchOrganizationRegistrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBatchOrganizationRegistrationTasks
// 该接口用于批量创建企业认证链接， 可以支持PC浏览器，H5和小程序三种途径。
//
// 此接口为异步提交任务接口，需要与[查询子企业批量认证链接](https://qcloudimg.tencent-cloud.cn/raw/1d3737991b2a3be78002bd78a47d6917.png)配合使用，整体流程如下图。
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/654aa2a72ab7d42f06464ea33c50c3bb.png)
//
// 
//
// 
//
// 
//
// **注意**
//
// 
//
// 1. 单次最多创建10个子企业。
//
// 2. 一天内，同一家企业最多创建8000个子企业。
//
// 3. 同一批创建的子客户不能重复，包括企业名称、企业统一信用代码和子客户经办人openId。
//
// 4. 跳转到小程序的实现，请参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式）。如何配置跳转电子签小程序，可参考：<a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 
//
// 
//
// 
//
// **腾讯电子签小程序的AppID 和 原始Id如下:**
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOROLEAUTH = "FailedOperation.NoRoleAuth"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchOrganizationRegistrationTasks(request *CreateBatchOrganizationRegistrationTasksRequest) (response *CreateBatchOrganizationRegistrationTasksResponse, err error) {
    return c.CreateBatchOrganizationRegistrationTasksWithContext(context.Background(), request)
}

// CreateBatchOrganizationRegistrationTasks
// 该接口用于批量创建企业认证链接， 可以支持PC浏览器，H5和小程序三种途径。
//
// 此接口为异步提交任务接口，需要与[查询子企业批量认证链接](https://qcloudimg.tencent-cloud.cn/raw/1d3737991b2a3be78002bd78a47d6917.png)配合使用，整体流程如下图。
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/654aa2a72ab7d42f06464ea33c50c3bb.png)
//
// 
//
// 
//
// 
//
// **注意**
//
// 
//
// 1. 单次最多创建10个子企业。
//
// 2. 一天内，同一家企业最多创建8000个子企业。
//
// 3. 同一批创建的子客户不能重复，包括企业名称、企业统一信用代码和子客户经办人openId。
//
// 4. 跳转到小程序的实现，请参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式）。如何配置跳转电子签小程序，可参考：<a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 
//
// 
//
// 
//
// **腾讯电子签小程序的AppID 和 原始Id如下:**
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOROLEAUTH = "FailedOperation.NoRoleAuth"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchOrganizationRegistrationTasksWithContext(ctx context.Context, request *CreateBatchOrganizationRegistrationTasksRequest) (response *CreateBatchOrganizationRegistrationTasksResponse, err error) {
    if request == nil {
        request = NewCreateBatchOrganizationRegistrationTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchOrganizationRegistrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchOrganizationRegistrationTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelFlowEvidenceReportRequest() (request *CreateChannelFlowEvidenceReportRequest) {
    request = &CreateChannelFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateChannelFlowEvidenceReport")
    
    
    return
}

func NewCreateChannelFlowEvidenceReportResponse() (response *CreateChannelFlowEvidenceReportResponse) {
    response = &CreateChannelFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateChannelFlowEvidenceReport
// 提交申请出证报告任务并返回报告ID。
//
// 
//
// 注意：
//
// - 使用此功能**需搭配出证套餐**  ，使用前请联系对接的客户经理沟通。
//
// - 操作人必须是**发起方或者签署方企业的(非走授权书认证)法人或者超管**。
//
// - 合同流程必须**所有参与方已经签署完成**。
//
// - 出证过程需一定时间，建议在**提交出证任务后的24小时之后**，通过<a href="https://qian.tencent.com/developers/partnerApis/certificate/DescribeChannelFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>接口进行查询执行结果和出证报告的下载URL。
//
// 
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1b4307ed143a992940c41d61192d3a0f/channel_CreateChannelFlowEvidenceReport.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_PROVENOQUOTA = "OperationDenied.ProveNoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateChannelFlowEvidenceReport(request *CreateChannelFlowEvidenceReportRequest) (response *CreateChannelFlowEvidenceReportResponse, err error) {
    return c.CreateChannelFlowEvidenceReportWithContext(context.Background(), request)
}

// CreateChannelFlowEvidenceReport
// 提交申请出证报告任务并返回报告ID。
//
// 
//
// 注意：
//
// - 使用此功能**需搭配出证套餐**  ，使用前请联系对接的客户经理沟通。
//
// - 操作人必须是**发起方或者签署方企业的(非走授权书认证)法人或者超管**。
//
// - 合同流程必须**所有参与方已经签署完成**。
//
// - 出证过程需一定时间，建议在**提交出证任务后的24小时之后**，通过<a href="https://qian.tencent.com/developers/partnerApis/certificate/DescribeChannelFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>接口进行查询执行结果和出证报告的下载URL。
//
// 
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1b4307ed143a992940c41d61192d3a0f/channel_CreateChannelFlowEvidenceReport.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_PROVENOQUOTA = "OperationDenied.ProveNoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateChannelFlowEvidenceReportWithContext(ctx context.Context, request *CreateChannelFlowEvidenceReportRequest) (response *CreateChannelFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewCreateChannelFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChannelFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChannelFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelOrganizationInfoChangeUrlRequest() (request *CreateChannelOrganizationInfoChangeUrlRequest) {
    request = &CreateChannelOrganizationInfoChangeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateChannelOrganizationInfoChangeUrl")
    
    
    return
}

func NewCreateChannelOrganizationInfoChangeUrlResponse() (response *CreateChannelOrganizationInfoChangeUrlResponse) {
    response = &CreateChannelOrganizationInfoChangeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateChannelOrganizationInfoChangeUrl
// 此接口（CreateChannelOrganizationInfoChangeUrl）用于创建子客企业信息变更链接。
//
// 
//
// <h3 id="1">支持变更链接类型，通过入参 Endpoint 指定，默认为WEIXINAPP。</h3>
//
// 
//
// <h4 id="WEIXINAPP">WEIXINAPP</h4>
//
// <p>创建变更短链。需要在移动端打开，会跳转到微信腾讯电子签小程序进行更换。</p>
//
// 
//
// <h4 id="APP">APP</h4>
//
// <p>创建变更小程序链接，可从第三方App跳转到微信腾讯电子签小程序进行更换。</p>
//
// 
//
// 
//
// <h3 id="2">支持创建企业超管变更链接或企业基础信息变更链接，通过入参 ChangeType 指定。</h3>
//
// 
//
// <h4 id="1-企业超管变更">1. 企业超管变更</h4>
//
// 
//
// <p>换成企业的其他员工来当超管</p>
//
// 
//
// <h4 id="2-企业基础信息变更">2. 企业基础信息变更</h4>
//
// 
//
// <h5 id="可以变动">可以变动</h5>
//
// 
//
// <ul>
//
// <li>企业名称<br>
//
// </li>
//
// <li>法定代表人姓名(新法人有邀请链接)<br>
//
// </li>
//
// <li>企业地址和所在地</li>
//
// </ul>
//
// 
//
// <h5 id="不可变动">不可变动</h5>
//
// 
//
// <ul>
//
// <li>统一社会信用代码<br>
//
// </li>
//
// <li>企业主体类型</li>
//
// </ul>
//
// 
//
// <p>如果企业名称变动会引起下面的变动</p>
//
// 
//
// <ul>
//
// <li>合同:   老合同不做任何处理,   新发起的合同需要用新的企业名字作为签署方, 否则无法签署</li>
//
// <li>印章:   会删除所有的印章所有的机构公章，合同专用章，财务专用章和人事专用章,  然后用新企业名称生成新的机构公章，合同专用章，财务专用章和人事专用章,  而法人章不会处理</li>
//
// <li>证书:   企业证书会重新请求CA机构用新企业名称生成新的证书</li>
//
// </ul>
//
// 
//
// 
//
// 注意： 
//
// 1. 生成的电子签小程序链接<font color='red'>只能由企业的法人或者超管</font>点击后进行操作， 其他员工打开后会提示“无权查看该内容”
//
// 2. 法人可以无需生成链接，直接在电子签小程序中更换本企业的超管
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateChannelOrganizationInfoChangeUrl(request *CreateChannelOrganizationInfoChangeUrlRequest) (response *CreateChannelOrganizationInfoChangeUrlResponse, err error) {
    return c.CreateChannelOrganizationInfoChangeUrlWithContext(context.Background(), request)
}

// CreateChannelOrganizationInfoChangeUrl
// 此接口（CreateChannelOrganizationInfoChangeUrl）用于创建子客企业信息变更链接。
//
// 
//
// <h3 id="1">支持变更链接类型，通过入参 Endpoint 指定，默认为WEIXINAPP。</h3>
//
// 
//
// <h4 id="WEIXINAPP">WEIXINAPP</h4>
//
// <p>创建变更短链。需要在移动端打开，会跳转到微信腾讯电子签小程序进行更换。</p>
//
// 
//
// <h4 id="APP">APP</h4>
//
// <p>创建变更小程序链接，可从第三方App跳转到微信腾讯电子签小程序进行更换。</p>
//
// 
//
// 
//
// <h3 id="2">支持创建企业超管变更链接或企业基础信息变更链接，通过入参 ChangeType 指定。</h3>
//
// 
//
// <h4 id="1-企业超管变更">1. 企业超管变更</h4>
//
// 
//
// <p>换成企业的其他员工来当超管</p>
//
// 
//
// <h4 id="2-企业基础信息变更">2. 企业基础信息变更</h4>
//
// 
//
// <h5 id="可以变动">可以变动</h5>
//
// 
//
// <ul>
//
// <li>企业名称<br>
//
// </li>
//
// <li>法定代表人姓名(新法人有邀请链接)<br>
//
// </li>
//
// <li>企业地址和所在地</li>
//
// </ul>
//
// 
//
// <h5 id="不可变动">不可变动</h5>
//
// 
//
// <ul>
//
// <li>统一社会信用代码<br>
//
// </li>
//
// <li>企业主体类型</li>
//
// </ul>
//
// 
//
// <p>如果企业名称变动会引起下面的变动</p>
//
// 
//
// <ul>
//
// <li>合同:   老合同不做任何处理,   新发起的合同需要用新的企业名字作为签署方, 否则无法签署</li>
//
// <li>印章:   会删除所有的印章所有的机构公章，合同专用章，财务专用章和人事专用章,  然后用新企业名称生成新的机构公章，合同专用章，财务专用章和人事专用章,  而法人章不会处理</li>
//
// <li>证书:   企业证书会重新请求CA机构用新企业名称生成新的证书</li>
//
// </ul>
//
// 
//
// 
//
// 注意： 
//
// 1. 生成的电子签小程序链接<font color='red'>只能由企业的法人或者超管</font>点击后进行操作， 其他员工打开后会提示“无权查看该内容”
//
// 2. 法人可以无需生成链接，直接在电子签小程序中更换本企业的超管
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateChannelOrganizationInfoChangeUrlWithContext(ctx context.Context, request *CreateChannelOrganizationInfoChangeUrlRequest) (response *CreateChannelOrganizationInfoChangeUrlResponse, err error) {
    if request == nil {
        request = NewCreateChannelOrganizationInfoChangeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChannelOrganizationInfoChangeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChannelOrganizationInfoChangeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChannelSubOrganizationActiveRequest() (request *CreateChannelSubOrganizationActiveRequest) {
    request = &CreateChannelSubOrganizationActiveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateChannelSubOrganizationActive")
    
    
    return
}

func NewCreateChannelSubOrganizationActiveResponse() (response *CreateChannelSubOrganizationActiveResponse) {
    response = &CreateChannelSubOrganizationActiveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateChannelSubOrganizationActive
// 激活或续期子客企业， 在激活状态下，企业可以正常发起合同；在非激活状态下，企业仅能查看和签署合同。
//
// 
//
// **1. 激活**：使用一个许可证将子客企业从未激活状态转变为激活状态。**激活状态的有效期为一年，一年后将自动回到未激活状态**。
//
// 
//
// **2. 续期**：使用一个许可证将已激活的子客企业的有效期延长一年。只有处于激活状态的子企业才能进行续期操作（**若处于非激活状态，则需先激活**）。您可以使用多个许可证对同一子客企业进行多次续费。
//
// 
//
// 
//
// 该接口的效果同：**【企业应用管理】 -> 【子客企业管理】 -> 【激活】或者【续期】**
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/cd63761ca6e814c64b4ecf131555b74e.png)
//
// 
//
// 
//
// 如果不想调用此接口或者页面点击进行激活或续期，可以在【应用扩展服务】中打开自动激活或续期，在许可证充足的情况下会自动激活或续期子客企业
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/2ccb37ef6bde463c15c39fdda789216f.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateChannelSubOrganizationActive(request *CreateChannelSubOrganizationActiveRequest) (response *CreateChannelSubOrganizationActiveResponse, err error) {
    return c.CreateChannelSubOrganizationActiveWithContext(context.Background(), request)
}

// CreateChannelSubOrganizationActive
// 激活或续期子客企业， 在激活状态下，企业可以正常发起合同；在非激活状态下，企业仅能查看和签署合同。
//
// 
//
// **1. 激活**：使用一个许可证将子客企业从未激活状态转变为激活状态。**激活状态的有效期为一年，一年后将自动回到未激活状态**。
//
// 
//
// **2. 续期**：使用一个许可证将已激活的子客企业的有效期延长一年。只有处于激活状态的子企业才能进行续期操作（**若处于非激活状态，则需先激活**）。您可以使用多个许可证对同一子客企业进行多次续费。
//
// 
//
// 
//
// 该接口的效果同：**【企业应用管理】 -> 【子客企业管理】 -> 【激活】或者【续期】**
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/cd63761ca6e814c64b4ecf131555b74e.png)
//
// 
//
// 
//
// 如果不想调用此接口或者页面点击进行激活或续期，可以在【应用扩展服务】中打开自动激活或续期，在许可证充足的情况下会自动激活或续期子客企业
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/2ccb37ef6bde463c15c39fdda789216f.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateChannelSubOrganizationActiveWithContext(ctx context.Context, request *CreateChannelSubOrganizationActiveRequest) (response *CreateChannelSubOrganizationActiveResponse, err error) {
    if request == nil {
        request = NewCreateChannelSubOrganizationActiveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChannelSubOrganizationActive require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChannelSubOrganizationActiveResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloseOrganizationUrlRequest() (request *CreateCloseOrganizationUrlRequest) {
    request = &CreateCloseOrganizationUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateCloseOrganizationUrl")
    
    
    return
}

func NewCreateCloseOrganizationUrlResponse() (response *CreateCloseOrganizationUrlResponse) {
    response = &CreateCloseOrganizationUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloseOrganizationUrl
// 创建企业注销链接
//
// 
//
// 系统将返回操作链接。贵方需要主动联系并通知企业的超级管理员（超管）或法人。由他们点击该链接，完成企业的注销操作。
//
// 
//
// 注意： `在调用此接口以管理企业扩展服务时，操作者（ Agent.ProxyOperator.OpenId）必须是企业的超级管理员（超管）或法人。`
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloseOrganizationUrl(request *CreateCloseOrganizationUrlRequest) (response *CreateCloseOrganizationUrlResponse, err error) {
    return c.CreateCloseOrganizationUrlWithContext(context.Background(), request)
}

// CreateCloseOrganizationUrl
// 创建企业注销链接
//
// 
//
// 系统将返回操作链接。贵方需要主动联系并通知企业的超级管理员（超管）或法人。由他们点击该链接，完成企业的注销操作。
//
// 
//
// 注意： `在调用此接口以管理企业扩展服务时，操作者（ Agent.ProxyOperator.OpenId）必须是企业的超级管理员（超管）或法人。`
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloseOrganizationUrlWithContext(ctx context.Context, request *CreateCloseOrganizationUrlRequest) (response *CreateCloseOrganizationUrlResponse, err error) {
    if request == nil {
        request = NewCreateCloseOrganizationUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloseOrganizationUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloseOrganizationUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsoleLoginUrlRequest() (request *CreateConsoleLoginUrlRequest) {
    request = &CreateConsoleLoginUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateConsoleLoginUrl")
    
    
    return
}

func NewCreateConsoleLoginUrlResponse() (response *CreateConsoleLoginUrlResponse) {
    response = &CreateConsoleLoginUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsoleLoginUrl
// 此接口（CreateConsoleLoginUrl）用于创建第三方平台子客企业控制台Web/移动登录链接。支持web控制台、电子签小程序和H5链接。登录链接是进入子客web企业控制台的唯一入口。
//
// 
//
// Web链接访问后，会根据子客企业(**Agent中ProxyOrganizationOpenId表示**)和员工(**Agent中OpenId表示**)的状态，进入不同的流程，主要情况分类如下：
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>子客企业状态</th>
//
// <th>子客企业员工状态</th>
//
// <th>点击链接进入的流程</th>
//
// </tr>
//
// </thead>
//
// <tbody>
//
// <tr>
//
// <td>企业未激活</td>
//
// <td>员工未认证</td>
//
// <td>进入企业激活流程，首次完成企业激活流程的员工会成为超管</td>
//
// </tr>
//
// <tr>
//
// <td>企业已激活</td>
//
// <td>员工未认证</td>
//
// <td>进入员工认证并加入企业流程</td>
//
// </tr>
//
// <tr>
//
// <td>企业已激活</td>
//
// <td>员工已认证</td>
//
// <td>进入子客企业Web控制台</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 如果是企业激活流程，需要注意如下情况：
//
// 
//
// 1. 若在激活过程中，**更换用户OpenID重新生成链接，之前的认证会被清理**。因此不要在企业认证过程生成多个链接给多人同时操作，会导致认证过程互相影响。
//
// 2. 若您认证中发现信息有误需要重新认证，**可通过更换用户OpenID重新生成链接的方式，来清理掉已有的流程**。
//
// 
//
// 系统的渠道企业, 应用, 子客企业, 子客员工的组织形式
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/77677faeea26c9d7f37474597c81fe01.png)
//
// 
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-createconsoleloginin.mp4" target="_blank">【生成子客登录链接】代码编写 &  子企业认证示例</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRNOTSYNCPROXYORGANIZATION = "FailedOperation.ErrNotSyncProxyOrganization"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZATIONTYPE = "InvalidParameter.AuthorizationType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UNIFORMSOCIALCREDITCODE = "InvalidParameter.UniformSocialCreditCode"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_ERRNOTOPENWEAKAUTHORIZATION = "OperationDenied.ErrNotOpenWeakAuthorization"
//  OPERATIONDENIED_NOTSUPPORTORGTYPE = "OperationDenied.NotSupportOrgType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConsoleLoginUrl(request *CreateConsoleLoginUrlRequest) (response *CreateConsoleLoginUrlResponse, err error) {
    return c.CreateConsoleLoginUrlWithContext(context.Background(), request)
}

// CreateConsoleLoginUrl
// 此接口（CreateConsoleLoginUrl）用于创建第三方平台子客企业控制台Web/移动登录链接。支持web控制台、电子签小程序和H5链接。登录链接是进入子客web企业控制台的唯一入口。
//
// 
//
// Web链接访问后，会根据子客企业(**Agent中ProxyOrganizationOpenId表示**)和员工(**Agent中OpenId表示**)的状态，进入不同的流程，主要情况分类如下：
//
// <table>
//
// <thead>
//
// <tr>
//
// <th>子客企业状态</th>
//
// <th>子客企业员工状态</th>
//
// <th>点击链接进入的流程</th>
//
// </tr>
//
// </thead>
//
// <tbody>
//
// <tr>
//
// <td>企业未激活</td>
//
// <td>员工未认证</td>
//
// <td>进入企业激活流程，首次完成企业激活流程的员工会成为超管</td>
//
// </tr>
//
// <tr>
//
// <td>企业已激活</td>
//
// <td>员工未认证</td>
//
// <td>进入员工认证并加入企业流程</td>
//
// </tr>
//
// <tr>
//
// <td>企业已激活</td>
//
// <td>员工已认证</td>
//
// <td>进入子客企业Web控制台</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 如果是企业激活流程，需要注意如下情况：
//
// 
//
// 1. 若在激活过程中，**更换用户OpenID重新生成链接，之前的认证会被清理**。因此不要在企业认证过程生成多个链接给多人同时操作，会导致认证过程互相影响。
//
// 2. 若您认证中发现信息有误需要重新认证，**可通过更换用户OpenID重新生成链接的方式，来清理掉已有的流程**。
//
// 
//
// 系统的渠道企业, 应用, 子客企业, 子客员工的组织形式
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/77677faeea26c9d7f37474597c81fe01.png)
//
// 
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-createconsoleloginin.mp4" target="_blank">【生成子客登录链接】代码编写 &  子企业认证示例</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRNOTSYNCPROXYORGANIZATION = "FailedOperation.ErrNotSyncProxyOrganization"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZATIONTYPE = "InvalidParameter.AuthorizationType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UNIFORMSOCIALCREDITCODE = "InvalidParameter.UniformSocialCreditCode"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_ERRNOTOPENWEAKAUTHORIZATION = "OperationDenied.ErrNotOpenWeakAuthorization"
//  OPERATIONDENIED_NOTSUPPORTORGTYPE = "OperationDenied.NotSupportOrgType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConsoleLoginUrlWithContext(ctx context.Context, request *CreateConsoleLoginUrlRequest) (response *CreateConsoleLoginUrlResponse, err error) {
    if request == nil {
        request = NewCreateConsoleLoginUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsoleLoginUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsoleLoginUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmployeeQualificationSealQrCodeRequest() (request *CreateEmployeeQualificationSealQrCodeRequest) {
    request = &CreateEmployeeQualificationSealQrCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateEmployeeQualificationSealQrCode")
    
    
    return
}

func NewCreateEmployeeQualificationSealQrCodeResponse() (response *CreateEmployeeQualificationSealQrCodeResponse) {
    response = &CreateEmployeeQualificationSealQrCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEmployeeQualificationSealQrCode
// 该接口用于获取个人授权执业章给企业的二维码，需要个人用户通过微信扫码。
//
// 
//
// 扫描后将跳转到腾讯电子签小程序，进入到授权执业章的流程。
//
// 
//
// 个人用户授权成功后，企业印章管理员需对印章进行审核，审核通过后，即可使用个人授权的执业章进行盖章操作。
//
// 
//
// **注意**
//
// 1. 该二维码**有效期为7天**，过期后将失效，可重新创建。
//
//  
//
// 
//
// 整体流程入下图
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/21b6b56dbc796c9d6f402d6ce6febb07.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
func (c *Client) CreateEmployeeQualificationSealQrCode(request *CreateEmployeeQualificationSealQrCodeRequest) (response *CreateEmployeeQualificationSealQrCodeResponse, err error) {
    return c.CreateEmployeeQualificationSealQrCodeWithContext(context.Background(), request)
}

// CreateEmployeeQualificationSealQrCode
// 该接口用于获取个人授权执业章给企业的二维码，需要个人用户通过微信扫码。
//
// 
//
// 扫描后将跳转到腾讯电子签小程序，进入到授权执业章的流程。
//
// 
//
// 个人用户授权成功后，企业印章管理员需对印章进行审核，审核通过后，即可使用个人授权的执业章进行盖章操作。
//
// 
//
// **注意**
//
// 1. 该二维码**有效期为7天**，过期后将失效，可重新创建。
//
//  
//
// 
//
// 整体流程入下图
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/21b6b56dbc796c9d6f402d6ce6febb07.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
func (c *Client) CreateEmployeeQualificationSealQrCodeWithContext(ctx context.Context, request *CreateEmployeeQualificationSealQrCodeRequest) (response *CreateEmployeeQualificationSealQrCodeResponse, err error) {
    if request == nil {
        request = NewCreateEmployeeQualificationSealQrCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmployeeQualificationSealQrCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmployeeQualificationSealQrCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowBlockchainEvidenceUrlRequest() (request *CreateFlowBlockchainEvidenceUrlRequest) {
    request = &CreateFlowBlockchainEvidenceUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateFlowBlockchainEvidenceUrl")
    
    
    return
}

func NewCreateFlowBlockchainEvidenceUrlResponse() (response *CreateFlowBlockchainEvidenceUrlResponse) {
    response = &CreateFlowBlockchainEvidenceUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowBlockchainEvidenceUrl
// 获取区块链存证证书查看链接/二维码接口
//
// 
//
// 适用场景：企业员工可以通过此接口生成合同区块链存证证书的查看链接/二维码，以供他人扫码打开腾讯电子签小程序查看。
//
// 
//
// [点击查看区块链存证证书样式](https://qcloudimg.tencent-cloud.cn/raw/47d5e9c2ffa90ad4e27b3cd14095aa08.jpg)
//
// 
//
// 注：
//
// <ul><li>1. 二维码下载链接过期时间为5分钟，请尽快下载保存。二维码/短链的过期时间为<font color="red">7天</font>，超过有效期则不可用。</li>
//
// <li>2. 合同状态需为<font color="red">签署完成</font> 、<font color="red">已解除</font>才能生成证书查看二维码/短链。</li>
//
// <li>3. 调用接口时，需确保接口调用身份拥有此合同的访问数据权限或为合同参与方。</li>
//
// <li>4. 通过扫码或者点击链接，用户无需登录或者鉴权即可查看对应合同的区块链存证证书，请妥善保管好二维码或链接。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateFlowBlockchainEvidenceUrl(request *CreateFlowBlockchainEvidenceUrlRequest) (response *CreateFlowBlockchainEvidenceUrlResponse, err error) {
    return c.CreateFlowBlockchainEvidenceUrlWithContext(context.Background(), request)
}

// CreateFlowBlockchainEvidenceUrl
// 获取区块链存证证书查看链接/二维码接口
//
// 
//
// 适用场景：企业员工可以通过此接口生成合同区块链存证证书的查看链接/二维码，以供他人扫码打开腾讯电子签小程序查看。
//
// 
//
// [点击查看区块链存证证书样式](https://qcloudimg.tencent-cloud.cn/raw/47d5e9c2ffa90ad4e27b3cd14095aa08.jpg)
//
// 
//
// 注：
//
// <ul><li>1. 二维码下载链接过期时间为5分钟，请尽快下载保存。二维码/短链的过期时间为<font color="red">7天</font>，超过有效期则不可用。</li>
//
// <li>2. 合同状态需为<font color="red">签署完成</font> 、<font color="red">已解除</font>才能生成证书查看二维码/短链。</li>
//
// <li>3. 调用接口时，需确保接口调用身份拥有此合同的访问数据权限或为合同参与方。</li>
//
// <li>4. 通过扫码或者点击链接，用户无需登录或者鉴权即可查看对应合同的区块链存证证书，请妥善保管好二维码或链接。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateFlowBlockchainEvidenceUrlWithContext(ctx context.Context, request *CreateFlowBlockchainEvidenceUrlRequest) (response *CreateFlowBlockchainEvidenceUrlResponse, err error) {
    if request == nil {
        request = NewCreateFlowBlockchainEvidenceUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowBlockchainEvidenceUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowBlockchainEvidenceUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowGroupSignReviewRequest() (request *CreateFlowGroupSignReviewRequest) {
    request = &CreateFlowGroupSignReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateFlowGroupSignReview")
    
    
    return
}

func NewCreateFlowGroupSignReviewResponse() (response *CreateFlowGroupSignReviewResponse) {
    response = &CreateFlowGroupSignReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowGroupSignReview
// 1. 在使用[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles)或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates)创建合同组签署流程时，若指定了参数以下参数为true,则可以调用此接口提交企业内部签署审批结果,即使是自动签署也需要进行审核通过才会进行签署。
//
//   - [FlowInfo.NeedSignReview](https://qian.tencent.com/developers/partnerApis/dataTypes/#flowinfo)
//
//   - [FlowFileInfo.NeedSignReview](https://qian.tencent.com/developers/partnerApis/dataTypes/#flowfileinfo)
//
//   - [FlowApproverInfo.ApproverNeedSignReview](https://qian.tencent.com/developers/partnerApis/dataTypes/#flowapproverinfo) 
//
// 
//
// 2. 同一合同组，同一签署人可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowGroupSignReview(request *CreateFlowGroupSignReviewRequest) (response *CreateFlowGroupSignReviewResponse, err error) {
    return c.CreateFlowGroupSignReviewWithContext(context.Background(), request)
}

// CreateFlowGroupSignReview
// 1. 在使用[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByFiles)或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates)创建合同组签署流程时，若指定了参数以下参数为true,则可以调用此接口提交企业内部签署审批结果,即使是自动签署也需要进行审核通过才会进行签署。
//
//   - [FlowInfo.NeedSignReview](https://qian.tencent.com/developers/partnerApis/dataTypes/#flowinfo)
//
//   - [FlowFileInfo.NeedSignReview](https://qian.tencent.com/developers/partnerApis/dataTypes/#flowfileinfo)
//
//   - [FlowApproverInfo.ApproverNeedSignReview](https://qian.tencent.com/developers/partnerApis/dataTypes/#flowapproverinfo) 
//
// 
//
// 2. 同一合同组，同一签署人可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowGroupSignReviewWithContext(ctx context.Context, request *CreateFlowGroupSignReviewRequest) (response *CreateFlowGroupSignReviewResponse, err error) {
    if request == nil {
        request = NewCreateFlowGroupSignReviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowGroupSignReview require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowGroupSignReviewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowsByTemplatesRequest() (request *CreateFlowsByTemplatesRequest) {
    request = &CreateFlowsByTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateFlowsByTemplates")
    
    
    return
}

func NewCreateFlowsByTemplatesResponse() (response *CreateFlowsByTemplatesResponse) {
    response = &CreateFlowsByTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowsByTemplates
// 接口（CreateFlowsByTemplates）用于使用模板批量创建签署流程。当前可批量发起合同（签署流程）数量为1-20个。
//
// 如若在模板中配置了动态表格, 上传的附件必须为A4大小 
//
// 合同发起人必须在电子签已经进行实名。
//
// 
//
// ### 一. 整体的逻辑如下
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/e193519d4383fa74782a9e19147ef01a/CreateFlowsByTemplates.png)
//
// 
//
// ###  二. 可以作为发起方和签署方的角色列表
//
// 
//
// <table>     <thead>     <tr>         <th>场景编号</th>         <th>发起方</th>         <th>签署方</th>         <th>补充</th>     </tr>     </thead>     <tbody>     <tr>         <td>场景一</td>         <td>子企业A的员工</td>         <td>子企业A的员工</td>         <td>子企业是通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">CreateConsoleLoginUrl</a>生成子客登录链接注册的企业</td>     </tr>     <tr>         <td>场景二</td>         <td>子企业A的员工</td>         <td>子企业B(不指定经办人走领取逻辑)</td>         <td>领取的逻辑可以参考文档<a href="https://qian.tencent.com/developers/partner/dynamic_signer" target="_blank">动态签署方</a> </td>     </tr>     <tr>         <td>场景三</td>         <td>子企业A的员工</td>         <td>子企业B的员工</td>         <td>-</td>     </tr>     <tr>         <td>场景四</td>         <td>子企业A的员工</td>         <td>个人</td>         <td>就是自然人，不是企业员工</td>     </tr>     <tr>         <td>场景五</td>         <td>子企业A的员工</td>         <td>SaaS平台企业员工</td>         <td>SaaS平台企业是通过<a href="https://qian.tencent.cn/console/company-register" target="_blank">https://qian.tencent.cn/console/company-register</a>链接注册的企业</td>     </tr>     </tbody> </table>
//
// 
//
// 
//
// 
//
// 
//
// ### 三. 填充模板中定义的填写控件
//
// 模板中配置的<font color="red">发起人填充控件</font>可以通过本接口的**FormFields数组**字段填充
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/37457e0e450fc221effddfcb8b1bad55.png)
//
// 填充的传参示例如下
//
// 
//
// ```
//
//     request.FormFields = [{
//
//             "ComponentName": "项目的名字",
//
//             "ComponentValue": "休闲山庄"
//
//         }, {
//
//             "ComponentName": "项目的地址",
//
//             "ComponentValue": "凤凰山北侧",
//
//         }, {
//
//             "ComponentName": "范围",
//
//             "ComponentValue": "凤凰山至107国道",
//
//         }, {
//
//             "ComponentName": "面积",
//
//             "ComponentValue": "100亩",
//
//         }, {
//
//             "ComponentName": "基本情况",
//
//             "ComponentValue": "完好",
//
//         }, , {
//
//             "ComponentName": "用途",
//
//             "ComponentValue": "经营农家乐",
//
//         }
//
//     ]
//
// ```
//
// 合成后合同样子示例
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/140a2fb771ac66a185d0a000d37485f6.png)
//
// 
//
// 
//
// 
//
// ### 四. 注意 
//
// 1. 发起合同时候,  作为<font color="red">发起方的第三方子企业A员工的企业和员工必须经过实名</font>, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名
//
// 
//
// 2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明
//
// 
//
// 3. <font color="red">调用接口发起合同成功就会扣减合同的额度</font>, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（过期，拒签，签署完成，解除完成等状态不会返还额度）
//
// 
//
// 4. <font color="red">静默（自动）签署不支持合同签署方存在填写</font>
//
// 
//
// 5.  <font color="red">在下一步创建签署链接前，建议等待DocumentFill </font> <a href="https://qian.tencent.com/developers/partner/callback_types_file_resources">PDF合成完成的回调</a>或者睡眠几秒，尤其是当模板中存在动态表格等复杂填写控件时，因为合成过程可能会耗费秒级别的时间。
//
// 
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-CreateTemplates.mp4" target="_blank">创建模板&设置成本企业自动签署</a><br>
//
// 2. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-CreateFlowsByTemplates.mp4" target="_blank">【用模板创建签署流程】编写示例视频教程</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowsByTemplates(request *CreateFlowsByTemplatesRequest) (response *CreateFlowsByTemplatesResponse, err error) {
    return c.CreateFlowsByTemplatesWithContext(context.Background(), request)
}

// CreateFlowsByTemplates
// 接口（CreateFlowsByTemplates）用于使用模板批量创建签署流程。当前可批量发起合同（签署流程）数量为1-20个。
//
// 如若在模板中配置了动态表格, 上传的附件必须为A4大小 
//
// 合同发起人必须在电子签已经进行实名。
//
// 
//
// ### 一. 整体的逻辑如下
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/e193519d4383fa74782a9e19147ef01a/CreateFlowsByTemplates.png)
//
// 
//
// ###  二. 可以作为发起方和签署方的角色列表
//
// 
//
// <table>     <thead>     <tr>         <th>场景编号</th>         <th>发起方</th>         <th>签署方</th>         <th>补充</th>     </tr>     </thead>     <tbody>     <tr>         <td>场景一</td>         <td>子企业A的员工</td>         <td>子企业A的员工</td>         <td>子企业是通过<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">CreateConsoleLoginUrl</a>生成子客登录链接注册的企业</td>     </tr>     <tr>         <td>场景二</td>         <td>子企业A的员工</td>         <td>子企业B(不指定经办人走领取逻辑)</td>         <td>领取的逻辑可以参考文档<a href="https://qian.tencent.com/developers/partner/dynamic_signer" target="_blank">动态签署方</a> </td>     </tr>     <tr>         <td>场景三</td>         <td>子企业A的员工</td>         <td>子企业B的员工</td>         <td>-</td>     </tr>     <tr>         <td>场景四</td>         <td>子企业A的员工</td>         <td>个人</td>         <td>就是自然人，不是企业员工</td>     </tr>     <tr>         <td>场景五</td>         <td>子企业A的员工</td>         <td>SaaS平台企业员工</td>         <td>SaaS平台企业是通过<a href="https://qian.tencent.cn/console/company-register" target="_blank">https://qian.tencent.cn/console/company-register</a>链接注册的企业</td>     </tr>     </tbody> </table>
//
// 
//
// 
//
// 
//
// 
//
// ### 三. 填充模板中定义的填写控件
//
// 模板中配置的<font color="red">发起人填充控件</font>可以通过本接口的**FormFields数组**字段填充
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/37457e0e450fc221effddfcb8b1bad55.png)
//
// 填充的传参示例如下
//
// 
//
// ```
//
//     request.FormFields = [{
//
//             "ComponentName": "项目的名字",
//
//             "ComponentValue": "休闲山庄"
//
//         }, {
//
//             "ComponentName": "项目的地址",
//
//             "ComponentValue": "凤凰山北侧",
//
//         }, {
//
//             "ComponentName": "范围",
//
//             "ComponentValue": "凤凰山至107国道",
//
//         }, {
//
//             "ComponentName": "面积",
//
//             "ComponentValue": "100亩",
//
//         }, {
//
//             "ComponentName": "基本情况",
//
//             "ComponentValue": "完好",
//
//         }, , {
//
//             "ComponentName": "用途",
//
//             "ComponentValue": "经营农家乐",
//
//         }
//
//     ]
//
// ```
//
// 合成后合同样子示例
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/140a2fb771ac66a185d0a000d37485f6.png)
//
// 
//
// 
//
// 
//
// ### 四. 注意 
//
// 1. 发起合同时候,  作为<font color="red">发起方的第三方子企业A员工的企业和员工必须经过实名</font>, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名
//
// 
//
// 2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明
//
// 
//
// 3. <font color="red">调用接口发起合同成功就会扣减合同的额度</font>, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（过期，拒签，签署完成，解除完成等状态不会返还额度）
//
// 
//
// 4. <font color="red">静默（自动）签署不支持合同签署方存在填写</font>
//
// 
//
// 5.  <font color="red">在下一步创建签署链接前，建议等待DocumentFill </font> <a href="https://qian.tencent.com/developers/partner/callback_types_file_resources">PDF合成完成的回调</a>或者睡眠几秒，尤其是当模板中存在动态表格等复杂填写控件时，因为合成过程可能会耗费秒级别的时间。
//
// 
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-CreateTemplates.mp4" target="_blank">创建模板&设置成本企业自动签署</a><br>
//
// 2. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-CreateFlowsByTemplates.mp4" target="_blank">【用模板创建签署流程】编写示例视频教程</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_APPROVERVERIFYTYPE = "InvalidParameter.ApproverVerifyType"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowsByTemplatesWithContext(ctx context.Context, request *CreateFlowsByTemplatesRequest) (response *CreateFlowsByTemplatesResponse, err error) {
    if request == nil {
        request = NewCreateFlowsByTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowsByTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowsByTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLegalSealQrCodeRequest() (request *CreateLegalSealQrCodeRequest) {
    request = &CreateLegalSealQrCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateLegalSealQrCode")
    
    
    return
}

func NewCreateLegalSealQrCodeResponse() (response *CreateLegalSealQrCodeResponse) {
    response = &CreateLegalSealQrCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLegalSealQrCode
// 该接口用于获取创建法人章的二维码，需要通过微信扫描。扫描后将跳转到腾讯电子签署，进入到创建法人章的流程。
//
// 
//
// **注意**
//
// 1. 该二维码**有效期为7天**，过期后将失效，可重新创建 。
//
// 2. 每个公司**只能有1个法人章**，无法重复创建或者创建多个
//
// 
//
// 法人章的样式可以参考下图索引（也可以自己上传法人印章图片）：
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/36a0a090750c45bb5cac5047ac461b2c.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRNOTSYNCPROXYORGANIZATION = "FailedOperation.ErrNotSyncProxyOrganization"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZATIONTYPE = "InvalidParameter.AuthorizationType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UNIFORMSOCIALCREDITCODE = "InvalidParameter.UniformSocialCreditCode"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_ERRNOTOPENWEAKAUTHORIZATION = "OperationDenied.ErrNotOpenWeakAuthorization"
//  OPERATIONDENIED_NOTSUPPORTORGTYPE = "OperationDenied.NotSupportOrgType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLegalSealQrCode(request *CreateLegalSealQrCodeRequest) (response *CreateLegalSealQrCodeResponse, err error) {
    return c.CreateLegalSealQrCodeWithContext(context.Background(), request)
}

// CreateLegalSealQrCode
// 该接口用于获取创建法人章的二维码，需要通过微信扫描。扫描后将跳转到腾讯电子签署，进入到创建法人章的流程。
//
// 
//
// **注意**
//
// 1. 该二维码**有效期为7天**，过期后将失效，可重新创建 。
//
// 2. 每个公司**只能有1个法人章**，无法重复创建或者创建多个
//
// 
//
// 法人章的样式可以参考下图索引（也可以自己上传法人印章图片）：
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/36a0a090750c45bb5cac5047ac461b2c.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRNOTSYNCPROXYORGANIZATION = "FailedOperation.ErrNotSyncProxyOrganization"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZATIONTYPE = "InvalidParameter.AuthorizationType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UNIFORMSOCIALCREDITCODE = "InvalidParameter.UniformSocialCreditCode"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_ERRNOTOPENWEAKAUTHORIZATION = "OperationDenied.ErrNotOpenWeakAuthorization"
//  OPERATIONDENIED_NOTSUPPORTORGTYPE = "OperationDenied.NotSupportOrgType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLegalSealQrCodeWithContext(ctx context.Context, request *CreateLegalSealQrCodeRequest) (response *CreateLegalSealQrCodeResponse, err error) {
    if request == nil {
        request = NewCreateLegalSealQrCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLegalSealQrCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLegalSealQrCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePartnerAutoSignAuthUrlRequest() (request *CreatePartnerAutoSignAuthUrlRequest) {
    request = &CreatePartnerAutoSignAuthUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreatePartnerAutoSignAuthUrl")
    
    
    return
}

func NewCreatePartnerAutoSignAuthUrlResponse() (response *CreatePartnerAutoSignAuthUrlResponse) {
    response = &CreatePartnerAutoSignAuthUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePartnerAutoSignAuthUrl
// 创建一个用于他方自动签授权的链接（可选择他方授权或我方授权）。通过这个链接，合作方企业可以直接进入小程序，进行自动签授权操作。
//
// 
//
// 如果授权企业尚未开通企业自动签功能，该链接还将引导他们首先开通本企业的自动签服务
//
// 
//
// 
//
// 注: 
//
// 1. <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 2. 已经在授权中或者授权成功的企业，无法重复授权
//
// 3. 授权企业和被授权企业必须都是已认证企业
//
// 4. <font color='red'>需要授权企业或被授权企业的超管或者法人打开链接</font>走开通逻辑。
//
// 
//
// **该接口效果同控制台： 企业设置-> 扩展服务 -> 企业自动签署 -> 合作企业方授权**
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/091823fd4f02af7dda416fa10ca65f2d.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreatePartnerAutoSignAuthUrl(request *CreatePartnerAutoSignAuthUrlRequest) (response *CreatePartnerAutoSignAuthUrlResponse, err error) {
    return c.CreatePartnerAutoSignAuthUrlWithContext(context.Background(), request)
}

// CreatePartnerAutoSignAuthUrl
// 创建一个用于他方自动签授权的链接（可选择他方授权或我方授权）。通过这个链接，合作方企业可以直接进入小程序，进行自动签授权操作。
//
// 
//
// 如果授权企业尚未开通企业自动签功能，该链接还将引导他们首先开通本企业的自动签服务
//
// 
//
// 
//
// 注: 
//
// 1. <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 2. 已经在授权中或者授权成功的企业，无法重复授权
//
// 3. 授权企业和被授权企业必须都是已认证企业
//
// 4. <font color='red'>需要授权企业或被授权企业的超管或者法人打开链接</font>走开通逻辑。
//
// 
//
// **该接口效果同控制台： 企业设置-> 扩展服务 -> 企业自动签署 -> 合作企业方授权**
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/091823fd4f02af7dda416fa10ca65f2d.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreatePartnerAutoSignAuthUrlWithContext(ctx context.Context, request *CreatePartnerAutoSignAuthUrlRequest) (response *CreatePartnerAutoSignAuthUrlResponse, err error) {
    if request == nil {
        request = NewCreatePartnerAutoSignAuthUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePartnerAutoSignAuthUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePartnerAutoSignAuthUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonAuthCertificateImageRequest() (request *CreatePersonAuthCertificateImageRequest) {
    request = &CreatePersonAuthCertificateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreatePersonAuthCertificateImage")
    
    
    return
}

func NewCreatePersonAuthCertificateImageResponse() (response *CreatePersonAuthCertificateImageResponse) {
    response = &CreatePersonAuthCertificateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePersonAuthCertificateImage
// 获取个人用户认证证书图片下载URL
//
// 
//
// 个人用户认证证书图片样式如下图
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/d568bc0f101bef582f7af2cb5ab7a715.png)
//
// 
//
// 注:  
//
// <ul>
//
// <li>只能获取个人用户证明图片, 企业员工的暂不支持</li>
//
// <li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。  </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePersonAuthCertificateImage(request *CreatePersonAuthCertificateImageRequest) (response *CreatePersonAuthCertificateImageResponse, err error) {
    return c.CreatePersonAuthCertificateImageWithContext(context.Background(), request)
}

// CreatePersonAuthCertificateImage
// 获取个人用户认证证书图片下载URL
//
// 
//
// 个人用户认证证书图片样式如下图
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/d568bc0f101bef582f7af2cb5ab7a715.png)
//
// 
//
// 注:  
//
// <ul>
//
// <li>只能获取个人用户证明图片, 企业员工的暂不支持</li>
//
// <li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。  </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePersonAuthCertificateImageWithContext(ctx context.Context, request *CreatePersonAuthCertificateImageRequest) (response *CreatePersonAuthCertificateImageResponse, err error) {
    if request == nil {
        request = NewCreatePersonAuthCertificateImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonAuthCertificateImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonAuthCertificateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSealByImageRequest() (request *CreateSealByImageRequest) {
    request = &CreateSealByImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSealByImage")
    
    
    return
}

func NewCreateSealByImageResponse() (response *CreateSealByImageResponse) {
    response = &CreateSealByImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSealByImage
// 1. 可以**通过图片**为子客企业代创建印章，图片最大5MB
//
// 
//
// 2. 可以**系统创建**子客企业代创建印章, 系统创建的印章样子下图(样式可以调整)
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage.png)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_EXISTSAMESEALNAME = "FailedOperation.ExistSameSealName"
//  INTERNALERROR_SEALUPLOAD = "InternalError.SealUpload"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_IMAGE = "InvalidParameter.Image"
//  INVALIDPARAMETER_LIMITSEALNAME = "InvalidParameter.LimitSealName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  MISSINGPARAMETER_ORGANIZATIONID = "MissingParameter.OrganizationId"
//  MISSINGPARAMETER_SEALIMAGE = "MissingParameter.SealImage"
//  MISSINGPARAMETER_SEALNAME = "MissingParameter.SealName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ALREADYHAS = "OperationDenied.AlreadyHas"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSealByImage(request *CreateSealByImageRequest) (response *CreateSealByImageResponse, err error) {
    return c.CreateSealByImageWithContext(context.Background(), request)
}

// CreateSealByImage
// 1. 可以**通过图片**为子客企业代创建印章，图片最大5MB
//
// 
//
// 2. 可以**系统创建**子客企业代创建印章, 系统创建的印章样子下图(样式可以调整)
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage.png)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_EXISTSAMESEALNAME = "FailedOperation.ExistSameSealName"
//  INTERNALERROR_SEALUPLOAD = "InternalError.SealUpload"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_IMAGE = "InvalidParameter.Image"
//  INVALIDPARAMETER_LIMITSEALNAME = "InvalidParameter.LimitSealName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  MISSINGPARAMETER_ORGANIZATIONID = "MissingParameter.OrganizationId"
//  MISSINGPARAMETER_SEALIMAGE = "MissingParameter.SealImage"
//  MISSINGPARAMETER_SEALNAME = "MissingParameter.SealName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ALREADYHAS = "OperationDenied.AlreadyHas"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSealByImageWithContext(ctx context.Context, request *CreateSealByImageRequest) (response *CreateSealByImageResponse, err error) {
    if request == nil {
        request = NewCreateSealByImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSealByImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSealByImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSignUrlsRequest() (request *CreateSignUrlsRequest) {
    request = &CreateSignUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSignUrls")
    
    
    return
}

func NewCreateSignUrlsResponse() (response *CreateSignUrlsResponse) {
    response = &CreateSignUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSignUrls
// 创建跳转小程序查看或签署的链接
//
// 
//
// **腾讯电子签小程序的AppID 和 原始Id如下:**
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 
//
// **主要使用场景EndPoint分类**
//
// 
//
// |EndPoint| 场景| 说明和示例|
//
// |  ----  | ----  | --- |
//
// |  WEIXINAPP  | 短链跳转腾讯电子签小程序签署场景  |  点击链接打开电子签小程序（与腾讯电子签官方短信提醒用户签署形式一样）<br> 示例: https://essurl.cn/x9nvWU8fTg|
//
// |  LONGURL2WEIXINAPP  | 长链跳转腾讯电子签小程序签署场景  |  点击链接打开电子签小程序, 是WEIXINAPP生成短链代表的那个长链|
//
// |  CHANNEL  | 带有H5引导页的跳转腾讯电子签小程序签署场景 |  点击链接打开一个H5引导页面, 页面中有个"前往小程序"的按钮, 点击后会跳转到腾讯电子签小程序签署场景;  签署完成会回到H5引导页面, 然后跳转到指定创建链接指定的JumpUrl<br>示例: https://res.ess.tencent.cn/cdn/h5-activity-beta/jump-mp.html?use=channel-guide&type=warning&token=uIFKIU8fTd |
//
// |APP| <font color="red">贵方APP</font>跳转腾讯电子签小程序签署场景|  贵方App直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方App的场景<br>跳转到腾讯电子签小程序的实现可以参考微信的官方文档:<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html" target="_blank">开放能力/打开 App</a> <br> 示例: pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3 |
//
// |APP| <font color="red">贵方小程序</font>跳转腾讯电子签小程序签署场景|  贵方小程序直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方小程序的场景<br>跳转到腾讯电子签小程序的实现可以参考微信官方文档<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html" target="_blank">全屏方式</a>和<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html " target="_blank">半屏方式</a><br>此时返回的SignUrl就是官方文档中的path<br> 示例:pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3  |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  INVALIDPARAMETER_GENERATETYPE = "InvalidParameter.GenerateType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWGROUP = "ResourceNotFound.FlowGroup"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSignUrls(request *CreateSignUrlsRequest) (response *CreateSignUrlsResponse, err error) {
    return c.CreateSignUrlsWithContext(context.Background(), request)
}

// CreateSignUrls
// 创建跳转小程序查看或签署的链接
//
// 
//
// **腾讯电子签小程序的AppID 和 原始Id如下:**
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 
//
// **主要使用场景EndPoint分类**
//
// 
//
// |EndPoint| 场景| 说明和示例|
//
// |  ----  | ----  | --- |
//
// |  WEIXINAPP  | 短链跳转腾讯电子签小程序签署场景  |  点击链接打开电子签小程序（与腾讯电子签官方短信提醒用户签署形式一样）<br> 示例: https://essurl.cn/x9nvWU8fTg|
//
// |  LONGURL2WEIXINAPP  | 长链跳转腾讯电子签小程序签署场景  |  点击链接打开电子签小程序, 是WEIXINAPP生成短链代表的那个长链|
//
// |  CHANNEL  | 带有H5引导页的跳转腾讯电子签小程序签署场景 |  点击链接打开一个H5引导页面, 页面中有个"前往小程序"的按钮, 点击后会跳转到腾讯电子签小程序签署场景;  签署完成会回到H5引导页面, 然后跳转到指定创建链接指定的JumpUrl<br>示例: https://res.ess.tencent.cn/cdn/h5-activity-beta/jump-mp.html?use=channel-guide&type=warning&token=uIFKIU8fTd |
//
// |APP| <font color="red">贵方APP</font>跳转腾讯电子签小程序签署场景|  贵方App直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方App的场景<br>跳转到腾讯电子签小程序的实现可以参考微信的官方文档:<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html" target="_blank">开放能力/打开 App</a> <br> 示例: pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3 |
//
// |APP| <font color="red">贵方小程序</font>跳转腾讯电子签小程序签署场景|  贵方小程序直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方小程序的场景<br>跳转到腾讯电子签小程序的实现可以参考微信官方文档<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html" target="_blank">全屏方式</a>和<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html " target="_blank">半屏方式</a><br>此时返回的SignUrl就是官方文档中的path<br> 示例:pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3  |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  INVALIDPARAMETER_GENERATETYPE = "InvalidParameter.GenerateType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWGROUP = "ResourceNotFound.FlowGroup"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSignUrlsWithContext(ctx context.Context, request *CreateSignUrlsRequest) (response *CreateSignUrlsResponse, err error) {
    if request == nil {
        request = NewCreateSignUrlsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSignUrls require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSignUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationAuthorizationsRequest() (request *DeleteOrganizationAuthorizationsRequest) {
    request = &DeleteOrganizationAuthorizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DeleteOrganizationAuthorizations")
    
    
    return
}

func NewDeleteOrganizationAuthorizationsResponse() (response *DeleteOrganizationAuthorizationsResponse) {
    response = &DeleteOrganizationAuthorizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationAuthorizations
// 批量清理未认证的企业认证流程。
//
// 
//
// 此接口用来清除企业方认证信息填写错误，批量清理认证中的认证流信息。
//
// 为接口[提交子企业批量认证链接创建任务](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks) 和[查询子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/DescribeBatchOrganizationRegistrationUrls) 接口的扩展接口。即在批量认证过程中，当发起认证企业发现超管信息错误的时候，可以将当前超管下的所有认证流企业清除。
//
// 
//
// 注意：
//
// **这个接口的操作人必须跟生成批量认证链接接口的应用号一致，才可以调用，否则会返回当前操作人没有认证中的企业认证流**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOROLEAUTH = "FailedOperation.NoRoleAuth"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteOrganizationAuthorizations(request *DeleteOrganizationAuthorizationsRequest) (response *DeleteOrganizationAuthorizationsResponse, err error) {
    return c.DeleteOrganizationAuthorizationsWithContext(context.Background(), request)
}

// DeleteOrganizationAuthorizations
// 批量清理未认证的企业认证流程。
//
// 
//
// 此接口用来清除企业方认证信息填写错误，批量清理认证中的认证流信息。
//
// 为接口[提交子企业批量认证链接创建任务](https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks) 和[查询子企业批量认证链接](https://qian.tencent.com/developers/partnerApis/accounts/DescribeBatchOrganizationRegistrationUrls) 接口的扩展接口。即在批量认证过程中，当发起认证企业发现超管信息错误的时候，可以将当前超管下的所有认证流企业清除。
//
// 
//
// 注意：
//
// **这个接口的操作人必须跟生成批量认证链接接口的应用号一致，才可以调用，否则会返回当前操作人没有认证中的企业认证流**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOROLEAUTH = "FailedOperation.NoRoleAuth"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteOrganizationAuthorizationsWithContext(ctx context.Context, request *DeleteOrganizationAuthorizationsRequest) (response *DeleteOrganizationAuthorizationsResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationAuthorizationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationAuthorizations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationAuthorizationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchOrganizationRegistrationUrlsRequest() (request *DescribeBatchOrganizationRegistrationUrlsRequest) {
    request = &DescribeBatchOrganizationRegistrationUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeBatchOrganizationRegistrationUrls")
    
    
    return
}

func NewDescribeBatchOrganizationRegistrationUrlsResponse() (response *DescribeBatchOrganizationRegistrationUrlsResponse) {
    response = &DescribeBatchOrganizationRegistrationUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBatchOrganizationRegistrationUrls
// 此接口用于获取企业批量认证异步任务的状态及结果。需要先调用接口<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks" target="_blank">提交子企业批量认证链接创建任务</a>获取到任务ID，然后再调用此接口获取到各个子企业的注册认证链接。整体流程如下图。
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/654aa2a72ab7d42f06464ea33c50c3bb.png)
//
// 
//
// 
//
// 注：
//
// `异步任务的处理完成时间视当前已提交的任务量、任务的复杂程度等因素决定，正常情况下 3~5 秒即可完成，但也可能需要更长的时间`
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeBatchOrganizationRegistrationUrls(request *DescribeBatchOrganizationRegistrationUrlsRequest) (response *DescribeBatchOrganizationRegistrationUrlsResponse, err error) {
    return c.DescribeBatchOrganizationRegistrationUrlsWithContext(context.Background(), request)
}

// DescribeBatchOrganizationRegistrationUrls
// 此接口用于获取企业批量认证异步任务的状态及结果。需要先调用接口<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateBatchOrganizationRegistrationTasks" target="_blank">提交子企业批量认证链接创建任务</a>获取到任务ID，然后再调用此接口获取到各个子企业的注册认证链接。整体流程如下图。
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/654aa2a72ab7d42f06464ea33c50c3bb.png)
//
// 
//
// 
//
// 注：
//
// `异步任务的处理完成时间视当前已提交的任务量、任务的复杂程度等因素决定，正常情况下 3~5 秒即可完成，但也可能需要更长的时间`
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeBatchOrganizationRegistrationUrlsWithContext(ctx context.Context, request *DescribeBatchOrganizationRegistrationUrlsRequest) (response *DescribeBatchOrganizationRegistrationUrlsResponse, err error) {
    if request == nil {
        request = NewDescribeBatchOrganizationRegistrationUrlsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchOrganizationRegistrationUrls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchOrganizationRegistrationUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCancelFlowsTaskRequest() (request *DescribeCancelFlowsTaskRequest) {
    request = &DescribeCancelFlowsTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeCancelFlowsTask")
    
    
    return
}

func NewDescribeCancelFlowsTaskResponse() (response *DescribeCancelFlowsTaskResponse) {
    response = &DescribeCancelFlowsTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCancelFlowsTask
// 通过接口[批量撤销合同流程](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelBatchCancelFlows)或者[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchCancelFlowUrl)发起批量撤销任务后，可通过此接口查询批量撤销任务的结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeCancelFlowsTask(request *DescribeCancelFlowsTaskRequest) (response *DescribeCancelFlowsTaskResponse, err error) {
    return c.DescribeCancelFlowsTaskWithContext(context.Background(), request)
}

// DescribeCancelFlowsTask
// 通过接口[批量撤销合同流程](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelBatchCancelFlows)或者[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/partnerApis/operateFlows/ChannelCreateBatchCancelFlowUrl)发起批量撤销任务后，可通过此接口查询批量撤销任务的结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeCancelFlowsTaskWithContext(ctx context.Context, request *DescribeCancelFlowsTaskRequest) (response *DescribeCancelFlowsTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCancelFlowsTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCancelFlowsTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCancelFlowsTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelFlowEvidenceReportRequest() (request *DescribeChannelFlowEvidenceReportRequest) {
    request = &DescribeChannelFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeChannelFlowEvidenceReport")
    
    
    return
}

func NewDescribeChannelFlowEvidenceReportResponse() (response *DescribeChannelFlowEvidenceReportResponse) {
    response = &DescribeChannelFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChannelFlowEvidenceReport
// 获取出证报告任务执行结果，返回报告 URL。
//
// 
//
// 注意：
//
// 
//
// - 使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。
//
// - 需调用创建并返回出证报告接口<a href="https://qian.tencent.com/developers/partnerApis/certificate/CreateChannelFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>获取报告编号后调用当前接口获取报告链接。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1b4307ed143a992940c41d61192d3a0f/channel_CreateChannelFlowEvidenceReport.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeChannelFlowEvidenceReport(request *DescribeChannelFlowEvidenceReportRequest) (response *DescribeChannelFlowEvidenceReportResponse, err error) {
    return c.DescribeChannelFlowEvidenceReportWithContext(context.Background(), request)
}

// DescribeChannelFlowEvidenceReport
// 获取出证报告任务执行结果，返回报告 URL。
//
// 
//
// 注意：
//
// 
//
// - 使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。
//
// - 需调用创建并返回出证报告接口<a href="https://qian.tencent.com/developers/partnerApis/certificate/CreateChannelFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>获取报告编号后调用当前接口获取报告链接。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1b4307ed143a992940c41d61192d3a0f/channel_CreateChannelFlowEvidenceReport.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeChannelFlowEvidenceReportWithContext(ctx context.Context, request *DescribeChannelFlowEvidenceReportRequest) (response *DescribeChannelFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewDescribeChannelFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChannelFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChannelFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelOrganizationsRequest() (request *DescribeChannelOrganizationsRequest) {
    request = &DescribeChannelOrganizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeChannelOrganizations")
    
    
    return
}

func NewDescribeChannelOrganizationsResponse() (response *DescribeChannelOrganizationsResponse) {
    response = &DescribeChannelOrganizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChannelOrganizations
// 查询渠道子客企业信息时，可以支持单个子客和整个应用下所有子客的查询。返回的信息包括超管、法人的信息以及当前企业的认证状态等信息。
//
// 
//
// - 对于单个企业的查询，通过**指定子客的唯一标识**来查询该子客的企业信息
//
// - 对于整个应用下所有企业的查询，**不需要指定子客的唯一标识**，直接查询整个应用下所有子客企业的企业信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeChannelOrganizations(request *DescribeChannelOrganizationsRequest) (response *DescribeChannelOrganizationsResponse, err error) {
    return c.DescribeChannelOrganizationsWithContext(context.Background(), request)
}

// DescribeChannelOrganizations
// 查询渠道子客企业信息时，可以支持单个子客和整个应用下所有子客的查询。返回的信息包括超管、法人的信息以及当前企业的认证状态等信息。
//
// 
//
// - 对于单个企业的查询，通过**指定子客的唯一标识**来查询该子客的企业信息
//
// - 对于整个应用下所有企业的查询，**不需要指定子客的唯一标识**，直接查询整个应用下所有子客企业的企业信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeChannelOrganizationsWithContext(ctx context.Context, request *DescribeChannelOrganizationsRequest) (response *DescribeChannelOrganizationsResponse, err error) {
    if request == nil {
        request = NewDescribeChannelOrganizationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChannelOrganizations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChannelOrganizationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelSealPolicyWorkflowUrlRequest() (request *DescribeChannelSealPolicyWorkflowUrlRequest) {
    request = &DescribeChannelSealPolicyWorkflowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeChannelSealPolicyWorkflowUrl")
    
    
    return
}

func NewDescribeChannelSealPolicyWorkflowUrlResponse() (response *DescribeChannelSealPolicyWorkflowUrlResponse) {
    response = &DescribeChannelSealPolicyWorkflowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChannelSealPolicyWorkflowUrl
// 生成用印申请审批链接，审批人可以通过此链接进入小程序进行审批。
//
//  p.s.
//
// Agent参数中的OpenId 必须为审批者的openId，且链接必须由审批人打开。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeChannelSealPolicyWorkflowUrl(request *DescribeChannelSealPolicyWorkflowUrlRequest) (response *DescribeChannelSealPolicyWorkflowUrlResponse, err error) {
    return c.DescribeChannelSealPolicyWorkflowUrlWithContext(context.Background(), request)
}

// DescribeChannelSealPolicyWorkflowUrl
// 生成用印申请审批链接，审批人可以通过此链接进入小程序进行审批。
//
//  p.s.
//
// Agent参数中的OpenId 必须为审批者的openId，且链接必须由审批人打开。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeChannelSealPolicyWorkflowUrlWithContext(ctx context.Context, request *DescribeChannelSealPolicyWorkflowUrlRequest) (response *DescribeChannelSealPolicyWorkflowUrlResponse, err error) {
    if request == nil {
        request = NewDescribeChannelSealPolicyWorkflowUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChannelSealPolicyWorkflowUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChannelSealPolicyWorkflowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtendedServiceAuthDetailRequest() (request *DescribeExtendedServiceAuthDetailRequest) {
    request = &DescribeExtendedServiceAuthDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeExtendedServiceAuthDetail")
    
    
    return
}

func NewDescribeExtendedServiceAuthDetailResponse() (response *DescribeExtendedServiceAuthDetailResponse) {
    response = &DescribeExtendedServiceAuthDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExtendedServiceAuthDetail
// 查询企业扩展服务的授权详情（列表），当前支持查询以下内容：
//
// 
//
// 1. **企业自动签**
//
// 2. **批量签署**
//
// 
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthDetail(request *DescribeExtendedServiceAuthDetailRequest) (response *DescribeExtendedServiceAuthDetailResponse, err error) {
    return c.DescribeExtendedServiceAuthDetailWithContext(context.Background(), request)
}

// DescribeExtendedServiceAuthDetail
// 查询企业扩展服务的授权详情（列表），当前支持查询以下内容：
//
// 
//
// 1. **企业自动签**
//
// 2. **批量签署**
//
// 
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthDetailWithContext(ctx context.Context, request *DescribeExtendedServiceAuthDetailRequest) (response *DescribeExtendedServiceAuthDetailResponse, err error) {
    if request == nil {
        request = NewDescribeExtendedServiceAuthDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtendedServiceAuthDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtendedServiceAuthDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtendedServiceAuthInfoRequest() (request *DescribeExtendedServiceAuthInfoRequest) {
    request = &DescribeExtendedServiceAuthInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeExtendedServiceAuthInfo")
    
    
    return
}

func NewDescribeExtendedServiceAuthInfoResponse() (response *DescribeExtendedServiceAuthInfoResponse) {
    response = &DescribeExtendedServiceAuthInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExtendedServiceAuthInfo
// 查询企业扩展服务的开通和授权情况，当前支持查询以下内容：
//
// 
//
// 1. **企业自动签**
//
// 2. **企业与港澳台居民签署合同**
//
// 3. **使用手机号验证签署方身份**
//
// 4. **拓宽签署方年龄限制**
//
// 5. **下载企业合同/文件**
//
// 6. **隐藏合同经办人姓名**
//
// 
//
// 对应能力开通页面在子客控制台-企业中心-拓展服务，如下图所示:
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/931a1e02955ab36e5cc69a489af10352.jpg)
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthInfo(request *DescribeExtendedServiceAuthInfoRequest) (response *DescribeExtendedServiceAuthInfoResponse, err error) {
    return c.DescribeExtendedServiceAuthInfoWithContext(context.Background(), request)
}

// DescribeExtendedServiceAuthInfo
// 查询企业扩展服务的开通和授权情况，当前支持查询以下内容：
//
// 
//
// 1. **企业自动签**
//
// 2. **企业与港澳台居民签署合同**
//
// 3. **使用手机号验证签署方身份**
//
// 4. **拓宽签署方年龄限制**
//
// 5. **下载企业合同/文件**
//
// 6. **隐藏合同经办人姓名**
//
// 
//
// 对应能力开通页面在子客控制台-企业中心-拓展服务，如下图所示:
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/931a1e02955ab36e5cc69a489af10352.jpg)
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthInfoWithContext(ctx context.Context, request *DescribeExtendedServiceAuthInfoRequest) (response *DescribeExtendedServiceAuthInfoResponse, err error) {
    if request == nil {
        request = NewDescribeExtendedServiceAuthInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtendedServiceAuthInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtendedServiceAuthInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowDetailInfoRequest() (request *DescribeFlowDetailInfoRequest) {
    request = &DescribeFlowDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFlowDetailInfo")
    
    
    return
}

func NewDescribeFlowDetailInfoResponse() (response *DescribeFlowDetailInfoResponse) {
    response = &DescribeFlowDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlowDetailInfo
// 此接口用于查询合同或者合同组的详情信息，支持查询多个（数量不能超过100）。
//
// 
//
// 适用场景：可用于主动查询某个合同或者合同组的详情信息。
//
// 
//
// 注:  `只能查询本企业创建的合同(创建合同用的Agent和此接口用的Agent数据最好一致) `
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_FFOWIDS = "InvalidParameter.fFowIds"
//  LIMITEXCEEDED_FLOWIDS = "LimitExceeded.FlowIds"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDS = "MissingParameter.FlowIds"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWAPPROVERS = "ResourceNotFound.FlowApprovers"
//  RESOURCENOTFOUND_FLOWGROUP = "ResourceNotFound.FlowGroup"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowDetailInfo(request *DescribeFlowDetailInfoRequest) (response *DescribeFlowDetailInfoResponse, err error) {
    return c.DescribeFlowDetailInfoWithContext(context.Background(), request)
}

// DescribeFlowDetailInfo
// 此接口用于查询合同或者合同组的详情信息，支持查询多个（数量不能超过100）。
//
// 
//
// 适用场景：可用于主动查询某个合同或者合同组的详情信息。
//
// 
//
// 注:  `只能查询本企业创建的合同(创建合同用的Agent和此接口用的Agent数据最好一致) `
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_FFOWIDS = "InvalidParameter.fFowIds"
//  LIMITEXCEEDED_FLOWIDS = "LimitExceeded.FlowIds"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDS = "MissingParameter.FlowIds"
//  MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWAPPROVERS = "ResourceNotFound.FlowApprovers"
//  RESOURCENOTFOUND_FLOWGROUP = "ResourceNotFound.FlowGroup"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowDetailInfoWithContext(ctx context.Context, request *DescribeFlowDetailInfoRequest) (response *DescribeFlowDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFlowDetailInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUrlsByFlowsRequest() (request *DescribeResourceUrlsByFlowsRequest) {
    request = &DescribeResourceUrlsByFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeResourceUrlsByFlows")
    
    
    return
}

func NewDescribeResourceUrlsByFlowsResponse() (response *DescribeResourceUrlsByFlowsResponse) {
    response = &DescribeResourceUrlsByFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceUrlsByFlows
// 获取合同流程PDF的下载链接，可以下载签署中、签署完的此子企业创建的合同。
//
// 
//
// 
//
// 
//
// 
//
// ### 2. 确保合同的PDF已经合成后，再调用本接口。
//
//  用户创建合同或者提交签署动作后，后台需要1~3秒的时间就进行合同PDF合成或者签名，为了确保您下载的是签署完成的完整合同文件，我们建议采取下面两种方式的一种来<font color="red"><b>确保PDF已经合成完成，然后在调用本接口</b></font>。
//
// 
//
// **第一种**：请确保您的系统配置了[接收合同完成通知的回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign)功能。一旦所有参与方签署完毕，我们的系统将自动向您提供的回调地址发送完成通知。
//
// 
//
// **第二种**：通过调用我们的[获取合同信息](https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo)接口来主动检查合同的签署状态。请仅在确认合同状态为“全部签署完成”后，进行文件的下载操作。
//
// 
//
// ### 3.  链接具有有效期限
//
// <font color="red"><b>生成的链接是有时间限制的，过期后将无法访问</b></font>。您可以在接口返回的信息中查看具体的过期时间。为避免错误，请确保在链接过期之前进行下载操作。
//
// 
//
// ### 4. 有两种开通下载权限的途径。
//
// 
//
// **第一种**:   需第三方应用的子企业登录控制台进行授权,  授权在**企业中心**的**授权管理**区域,  界面如下图。
//
// 授权过程需要**子企业超管**扫描跳转到电子签小程序签署<<渠道端下载渠道子客合同功能授权委托书>>
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/8b483dfebdeafac85051279406944048.png)
//
// 
//
// **第二种**: 渠道方企业在**企业应用管理**的配置界面打开需要配置的应用，点击**应用扩展服务**开通此功能，需要**渠道方企业的超管**扫描二维码跳转到电子签小程序签署 <<渠道端下载渠道子客合同功能开通知情同意书>>
//
// 注: 
//
// 1. `请注意如果第三方应用的子客主动关闭了渠道端下载渠道子客合同功能开关，那么渠道方开通了此功能也无法下载子客的合同文件`
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/238979ef51dd381ccbdbc755a593debc/channel_DescribeResourceUrlsByFlows_appilications2.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_USEROPENID = "MissingParameter.UserOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
func (c *Client) DescribeResourceUrlsByFlows(request *DescribeResourceUrlsByFlowsRequest) (response *DescribeResourceUrlsByFlowsResponse, err error) {
    return c.DescribeResourceUrlsByFlowsWithContext(context.Background(), request)
}

// DescribeResourceUrlsByFlows
// 获取合同流程PDF的下载链接，可以下载签署中、签署完的此子企业创建的合同。
//
// 
//
// 
//
// 
//
// 
//
// ### 2. 确保合同的PDF已经合成后，再调用本接口。
//
//  用户创建合同或者提交签署动作后，后台需要1~3秒的时间就进行合同PDF合成或者签名，为了确保您下载的是签署完成的完整合同文件，我们建议采取下面两种方式的一种来<font color="red"><b>确保PDF已经合成完成，然后在调用本接口</b></font>。
//
// 
//
// **第一种**：请确保您的系统配置了[接收合同完成通知的回调](https://qian.tencent.com/developers/partner/callback_types_contracts_sign)功能。一旦所有参与方签署完毕，我们的系统将自动向您提供的回调地址发送完成通知。
//
// 
//
// **第二种**：通过调用我们的[获取合同信息](https://qian.tencent.com/developers/partnerApis/flows/DescribeFlowDetailInfo)接口来主动检查合同的签署状态。请仅在确认合同状态为“全部签署完成”后，进行文件的下载操作。
//
// 
//
// ### 3.  链接具有有效期限
//
// <font color="red"><b>生成的链接是有时间限制的，过期后将无法访问</b></font>。您可以在接口返回的信息中查看具体的过期时间。为避免错误，请确保在链接过期之前进行下载操作。
//
// 
//
// ### 4. 有两种开通下载权限的途径。
//
// 
//
// **第一种**:   需第三方应用的子企业登录控制台进行授权,  授权在**企业中心**的**授权管理**区域,  界面如下图。
//
// 授权过程需要**子企业超管**扫描跳转到电子签小程序签署<<渠道端下载渠道子客合同功能授权委托书>>
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/8b483dfebdeafac85051279406944048.png)
//
// 
//
// **第二种**: 渠道方企业在**企业应用管理**的配置界面打开需要配置的应用，点击**应用扩展服务**开通此功能，需要**渠道方企业的超管**扫描二维码跳转到电子签小程序签署 <<渠道端下载渠道子客合同功能开通知情同意书>>
//
// 注: 
//
// 1. `请注意如果第三方应用的子客主动关闭了渠道端下载渠道子客合同功能开关，那么渠道方开通了此功能也无法下载子客的合同文件`
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/238979ef51dd381ccbdbc755a593debc/channel_DescribeResourceUrlsByFlows_appilications2.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_USEROPENID = "MissingParameter.UserOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
func (c *Client) DescribeResourceUrlsByFlowsWithContext(ctx context.Context, request *DescribeResourceUrlsByFlowsRequest) (response *DescribeResourceUrlsByFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUrlsByFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceUrlsByFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceUrlsByFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
    request = &DescribeTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeTemplates")
    
    
    return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
    response = &DescribeTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTemplates
// 通过此接口（DescribeTemplates）查询该第三方平台子客企业在电子签拥有的有效模板，不包括第三方平台模板。
//
// 
//
// **适用场景**
//
//  该接口常用来配合<a href="https://qian.tencent.com/developers/partnerApis/startFlows/CreateFlowsByTemplates" target="_blank">用模板创建签署流程</a>和<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates" target="_blank">通过多模板创建合同组签署流程</a>接口，作为创建合同的前置接口使用。 
//
// 通过此接口查询到模板信息后，再通过调用创建合同的接口，指定模板ID，指定模板中需要的填写控件内容等，完成合同文档的创建。
//
// 
//
// **模板的来源**
//
// 子客企业的模板有两种途径获取
//
// - 渠道方(平台方)配置完成后, 分发给同应用的各个子企业
//
// - 子客企业通过CreateConsoleLoginUrl创建的链接登录子客控制台自己创建
//
// 
//
// **一个模板通常会包含以下结构信息** 
//
// 
//
// - 模板ID, 模板名字等基本信息
//
// - 发起方参与信息Promoter、签署参与方 Recipients，后者会在模板发起合同时用于指定参与方
//
// - 发起方和签署方的填写控件 Components
//
// - 签署方的签署控件 SignComponents
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/ab81fa948a0a6fea14f48cac91d0e36a/channel_DescribeTemplates.png)
//
// 
//
// 模板中各元素的层级关系, 所有的填写控件和签署控件都归属某一个角色(通过控件的ComponentRecipientId来关联)
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/45c638bd93f9c8024763add9ab47c27f.png)
//
// 
//
// 
//
// **注意**
//
// 
//
// >1. 查询条件TemplateId、TemplateName与ChannelTemplateId可同时存在，即可查询同时满足这些条件的模板。
//
// >2. TemplateId 和TemplateIds互为独立，若两个参数都传入，则以TemplateId为准
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-CreateTemplates.mp4" target="_blank">创建模板&设置成本企业自动签署</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
    return c.DescribeTemplatesWithContext(context.Background(), request)
}

// DescribeTemplates
// 通过此接口（DescribeTemplates）查询该第三方平台子客企业在电子签拥有的有效模板，不包括第三方平台模板。
//
// 
//
// **适用场景**
//
//  该接口常用来配合<a href="https://qian.tencent.com/developers/partnerApis/startFlows/CreateFlowsByTemplates" target="_blank">用模板创建签署流程</a>和<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowGroupByTemplates" target="_blank">通过多模板创建合同组签署流程</a>接口，作为创建合同的前置接口使用。 
//
// 通过此接口查询到模板信息后，再通过调用创建合同的接口，指定模板ID，指定模板中需要的填写控件内容等，完成合同文档的创建。
//
// 
//
// **模板的来源**
//
// 子客企业的模板有两种途径获取
//
// - 渠道方(平台方)配置完成后, 分发给同应用的各个子企业
//
// - 子客企业通过CreateConsoleLoginUrl创建的链接登录子客控制台自己创建
//
// 
//
// **一个模板通常会包含以下结构信息** 
//
// 
//
// - 模板ID, 模板名字等基本信息
//
// - 发起方参与信息Promoter、签署参与方 Recipients，后者会在模板发起合同时用于指定参与方
//
// - 发起方和签署方的填写控件 Components
//
// - 签署方的签署控件 SignComponents
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/ab81fa948a0a6fea14f48cac91d0e36a/channel_DescribeTemplates.png)
//
// 
//
// 模板中各元素的层级关系, 所有的填写控件和签署控件都归属某一个角色(通过控件的ComponentRecipientId来关联)
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/45c638bd93f9c8024763add9ab47c27f.png)
//
// 
//
// 
//
// **注意**
//
// 
//
// >1. 查询条件TemplateId、TemplateName与ChannelTemplateId可同时存在，即可查询同时满足这些条件的模板。
//
// >2. TemplateId 和TemplateIds互为独立，若两个参数都传入，则以TemplateId为准
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-CreateTemplates.mp4" target="_blank">创建模板&设置成本企业自动签署</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeTemplatesWithContext(ctx context.Context, request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageRequest() (request *DescribeUsageRequest) {
    request = &DescribeUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeUsage")
    
    
    return
}

func NewDescribeUsageResponse() (response *DescribeUsageResponse) {
    response = &DescribeUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsage
// 此接口（DescribeUsage）用于获取此应用下子客企业的合同消耗数量。
//
// 
//
// <font color="red">此接口即将下线， 请使用新接口</font>  [查询渠道计费消耗情况](https://qian.tencent.com/developers/partnerApis/fee/ChannelDescribeBillUsageDetail)
//
// 
//
// 注: 此接口**每日限频50次**，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATE = "InvalidParameter.Date"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  LIMITEXCEEDED_CALLTIMES = "LimitExceeded.CallTimes"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DATE = "MissingParameter.Date"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsage(request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    return c.DescribeUsageWithContext(context.Background(), request)
}

// DescribeUsage
// 此接口（DescribeUsage）用于获取此应用下子客企业的合同消耗数量。
//
// 
//
// <font color="red">此接口即将下线， 请使用新接口</font>  [查询渠道计费消耗情况](https://qian.tencent.com/developers/partnerApis/fee/ChannelDescribeBillUsageDetail)
//
// 
//
// 注: 此接口**每日限频50次**，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATE = "InvalidParameter.Date"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  LIMITEXCEEDED_CALLTIMES = "LimitExceeded.CallTimes"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DATE = "MissingParameter.Date"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsageWithContext(ctx context.Context, request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    if request == nil {
        request = NewDescribeUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageResponse()
    err = c.Send(request, response)
    return
}

func NewGetDownloadFlowUrlRequest() (request *GetDownloadFlowUrlRequest) {
    request = &GetDownloadFlowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "GetDownloadFlowUrl")
    
    
    return
}

func NewGetDownloadFlowUrlResponse() (response *GetDownloadFlowUrlResponse) {
    response = &GetDownloadFlowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDownloadFlowUrl
// 此接口（GetDownloadFlowUrl）用户获取合同控制台下载页面链接,  点击链接后会跳转至本企业合同管理控制台(会筛选出传入的合同列表), 点击**下载**按钮后就会下载传入的合同列表, 下载页面如下图
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_GetDownloadFlowUrl.png)
//
// 
//
// 注:
//
// <ul>
//
// <li>仅支持下载 <b>本企业</b> 下合同，链接会 <b>登录企业控制台</b> </li>
//
// <li> <b>链接仅可使用一次</b>，不可重复使用</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWNUMEXCEED = "FailedOperation.FlowNumExceed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWINFO = "MissingParameter.FlowInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_DOWNLOADMORETHANONE = "OperationDenied.DownLoadMoreThanOne"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FILE = "ResourceNotFound.File"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_TEAMWORKORGANIZATION = "ResourceNotFound.TeamWorkOrganization"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDownloadFlowUrl(request *GetDownloadFlowUrlRequest) (response *GetDownloadFlowUrlResponse, err error) {
    return c.GetDownloadFlowUrlWithContext(context.Background(), request)
}

// GetDownloadFlowUrl
// 此接口（GetDownloadFlowUrl）用户获取合同控制台下载页面链接,  点击链接后会跳转至本企业合同管理控制台(会筛选出传入的合同列表), 点击**下载**按钮后就会下载传入的合同列表, 下载页面如下图
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_GetDownloadFlowUrl.png)
//
// 
//
// 注:
//
// <ul>
//
// <li>仅支持下载 <b>本企业</b> 下合同，链接会 <b>登录企业控制台</b> </li>
//
// <li> <b>链接仅可使用一次</b>，不可重复使用</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWNUMEXCEED = "FailedOperation.FlowNumExceed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWINFO = "MissingParameter.FlowInfo"
//  MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_DOWNLOADMORETHANONE = "OperationDenied.DownLoadMoreThanOne"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  RESOURCENOTFOUND_FILE = "ResourceNotFound.File"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_TEAMWORKORGANIZATION = "ResourceNotFound.TeamWorkOrganization"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONDOWNLOAD = "UnauthorizedOperation.UnauthorizedOperationDownload"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONORGANIZATION = "UnauthorizedOperation.UnauthorizedOperationOrganization"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDownloadFlowUrlWithContext(ctx context.Context, request *GetDownloadFlowUrlRequest) (response *GetDownloadFlowUrlResponse, err error) {
    if request == nil {
        request = NewGetDownloadFlowUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDownloadFlowUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDownloadFlowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyExtendedServiceRequest() (request *ModifyExtendedServiceRequest) {
    request = &ModifyExtendedServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifyExtendedService")
    
    
    return
}

func NewModifyExtendedServiceResponse() (response *ModifyExtendedServiceResponse) {
    response = &ModifyExtendedServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyExtendedService
// 管理企业扩展服务
//
// 
//
// - **直接开通的情形：** 若在操作过程中接口没有返回跳转链接，这表明无需进行任何跳转操作。此时，相应的企业拓展服务将会直接被开通或关闭。
//
// 
//
// - **需要法人或者超管签署开通协议的情形：** 当需要开通以下企业拓展服务时， 系统将返回一个操作链接。贵方需要主动联系并通知企业的超级管理员（超管）或法人。由他们点击该链接，完成服务的开通操作。
//
//   - **AUTO_SIGN（企业自动签）**
//
//   - **DOWNLOAD_FLOW（授权渠道下载合同）**
//
// 
//
// 注意： `在调用此接口以管理企业扩展服务时，操作者（ Agent.ProxyOperator.OpenId）必须是企业的超级管理员（超管）或法人`
//
// 
//
// 
//
// 对应的扩展服务能力可以在控制台的【扩展服务】中找到
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/99eebd37883ec55ed1f1df3a57aee60a.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyExtendedService(request *ModifyExtendedServiceRequest) (response *ModifyExtendedServiceResponse, err error) {
    return c.ModifyExtendedServiceWithContext(context.Background(), request)
}

// ModifyExtendedService
// 管理企业扩展服务
//
// 
//
// - **直接开通的情形：** 若在操作过程中接口没有返回跳转链接，这表明无需进行任何跳转操作。此时，相应的企业拓展服务将会直接被开通或关闭。
//
// 
//
// - **需要法人或者超管签署开通协议的情形：** 当需要开通以下企业拓展服务时， 系统将返回一个操作链接。贵方需要主动联系并通知企业的超级管理员（超管）或法人。由他们点击该链接，完成服务的开通操作。
//
//   - **AUTO_SIGN（企业自动签）**
//
//   - **DOWNLOAD_FLOW（授权渠道下载合同）**
//
// 
//
// 注意： `在调用此接口以管理企业扩展服务时，操作者（ Agent.ProxyOperator.OpenId）必须是企业的超级管理员（超管）或法人`
//
// 
//
// 
//
// 对应的扩展服务能力可以在控制台的【扩展服务】中找到
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/99eebd37883ec55ed1f1df3a57aee60a.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyExtendedServiceWithContext(ctx context.Context, request *ModifyExtendedServiceRequest) (response *ModifyExtendedServiceResponse, err error) {
    if request == nil {
        request = NewModifyExtendedServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyExtendedService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyExtendedServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFlowDeadlineRequest() (request *ModifyFlowDeadlineRequest) {
    request = &ModifyFlowDeadlineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifyFlowDeadline")
    
    
    return
}

func NewModifyFlowDeadlineResponse() (response *ModifyFlowDeadlineResponse) {
    response = &ModifyFlowDeadlineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFlowDeadline
// 在已启动的签署流程中，可对签署截止日期进行延期操作，主要分为以下两个层面：
//
// 1. <b> 合同（流程）层面</b>：仅需提供签署流程ID。此操作将对整个签署流程以及未单独设置签署截止时间的签署人进行延期。
//
// 2. <b> 签署人层面</b>  ：需提供流程ID和签署人ID。此操作针对特定签署人进行延期，特别是对于有序合同（流程），签署截止时间不得超过后续签署人的流程截止时间。
//
// 
//
// 此接口存在以下限制：
//
// 1. 执行操作的员工须为<font  color="red">发起方企业的超级管理员、法定代表人或签署流程发起人</font>。
//
// 2. 延长整个签署流程时，<font  color="red">应至少有一方尚未签署</font>（即签署流程不能处于已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等状态）。
//
// 3. 延长整个签署流程时，新的签署截止日期应晚于已设定的签署截止日期和当前日期。
//
// 4. 延长签署方截止时间时，<font  color="red">签署方不能处于流程完结或已终止状态</font>（即签署人不能处于已签署、已拒签、已过期、已撤回、拒绝填写、已解除等状态）。
//
// 5. 延长签署方截止时间时，新的签署截止日期应晚于当前日期和已设定的截止日期。若为有序合同，还需早于或等于下一签署人的截止日期，且早于签署流程整体的截止日期。
//
// 6. <font  color="red">不支持操作合同组合同</font>。
//
// 
//
// 合同（流程）层面截止时间子企业控制台展示的位置：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/f0f88c0eb49a926da9a86e5a6e9efa8b.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyFlowDeadline(request *ModifyFlowDeadlineRequest) (response *ModifyFlowDeadlineResponse, err error) {
    return c.ModifyFlowDeadlineWithContext(context.Background(), request)
}

// ModifyFlowDeadline
// 在已启动的签署流程中，可对签署截止日期进行延期操作，主要分为以下两个层面：
//
// 1. <b> 合同（流程）层面</b>：仅需提供签署流程ID。此操作将对整个签署流程以及未单独设置签署截止时间的签署人进行延期。
//
// 2. <b> 签署人层面</b>  ：需提供流程ID和签署人ID。此操作针对特定签署人进行延期，特别是对于有序合同（流程），签署截止时间不得超过后续签署人的流程截止时间。
//
// 
//
// 此接口存在以下限制：
//
// 1. 执行操作的员工须为<font  color="red">发起方企业的超级管理员、法定代表人或签署流程发起人</font>。
//
// 2. 延长整个签署流程时，<font  color="red">应至少有一方尚未签署</font>（即签署流程不能处于已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等状态）。
//
// 3. 延长整个签署流程时，新的签署截止日期应晚于已设定的签署截止日期和当前日期。
//
// 4. 延长签署方截止时间时，<font  color="red">签署方不能处于流程完结或已终止状态</font>（即签署人不能处于已签署、已拒签、已过期、已撤回、拒绝填写、已解除等状态）。
//
// 5. 延长签署方截止时间时，新的签署截止日期应晚于当前日期和已设定的截止日期。若为有序合同，还需早于或等于下一签署人的截止日期，且早于签署流程整体的截止日期。
//
// 6. <font  color="red">不支持操作合同组合同</font>。
//
// 
//
// 合同（流程）层面截止时间子企业控制台展示的位置：
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/f0f88c0eb49a926da9a86e5a6e9efa8b.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyFlowDeadlineWithContext(ctx context.Context, request *ModifyFlowDeadlineRequest) (response *ModifyFlowDeadlineResponse, err error) {
    if request == nil {
        request = NewModifyFlowDeadlineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFlowDeadline require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFlowDeadlineResponse()
    err = c.Send(request, response)
    return
}

func NewOperateChannelTemplateRequest() (request *OperateChannelTemplateRequest) {
    request = &OperateChannelTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "OperateChannelTemplate")
    
    
    return
}

func NewOperateChannelTemplateResponse() (response *OperateChannelTemplateResponse) {
    response = &OperateChannelTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OperateChannelTemplate
// 此接口（OperateChannelTemplate）用于针对第三方应用平台模板库中的模板对子客企业发布授权的查询和设置。
//
// 平台模板库中的模板的位置在控制台 企业应用管理 中下面的应用模板库管理目录, 可以参照下图位置
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7f2b6c94164b3e931efc9a037e0400f7.png)
//
// 
//
// # 支持的操作
//
// 
//
// ## 1. 查询模板的子客企业授权 (OperateType=SELECT)
//
// - 查询模板的授权子企业列表
//
// 
//
// ## 2. 修改模板的子客企业授权 (OperateType=UPDATE)
//
// - 当模板未发布时，可以修改模板的模板授权范围是**所有第三方应用合作企业**(AuthTag设置为all)或者**指定第三方应用合作企业**(AuthTag设置为part)，**当模板发布后，不可做此修改**
//
// - 如果模板是部分授权,  可通过ProxyOrganizationOpenIds增加子客的授权范围。
//
// 
//
// ## 3. 取消模板的子客企业授权 (OperateType=DELETE)
//
// - 对子客企业进行模板库中模板授权范围的进行删除操作。
//
// - 主要对于手动领取的模板，去除授权后子客在模板库中看不到，就无法再领取了。但是**已经领取过成为自有模板的不会同步删除**。
//
// - 对于自动领取的模板，由于已经下发，更改授权不会影响。
//
// - 如果要同步删除子客自有模板库中的模板，请使用OperateType=UPDATE+Available参数处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATES = "MissingParameter.Templates"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTHTAG = "OperationDenied.AuthTag"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONAUTH = "ResourceNotFound.ApplicationAuth"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OperateChannelTemplate(request *OperateChannelTemplateRequest) (response *OperateChannelTemplateResponse, err error) {
    return c.OperateChannelTemplateWithContext(context.Background(), request)
}

// OperateChannelTemplate
// 此接口（OperateChannelTemplate）用于针对第三方应用平台模板库中的模板对子客企业发布授权的查询和设置。
//
// 平台模板库中的模板的位置在控制台 企业应用管理 中下面的应用模板库管理目录, 可以参照下图位置
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7f2b6c94164b3e931efc9a037e0400f7.png)
//
// 
//
// # 支持的操作
//
// 
//
// ## 1. 查询模板的子客企业授权 (OperateType=SELECT)
//
// - 查询模板的授权子企业列表
//
// 
//
// ## 2. 修改模板的子客企业授权 (OperateType=UPDATE)
//
// - 当模板未发布时，可以修改模板的模板授权范围是**所有第三方应用合作企业**(AuthTag设置为all)或者**指定第三方应用合作企业**(AuthTag设置为part)，**当模板发布后，不可做此修改**
//
// - 如果模板是部分授权,  可通过ProxyOrganizationOpenIds增加子客的授权范围。
//
// 
//
// ## 3. 取消模板的子客企业授权 (OperateType=DELETE)
//
// - 对子客企业进行模板库中模板授权范围的进行删除操作。
//
// - 主要对于手动领取的模板，去除授权后子客在模板库中看不到，就无法再领取了。但是**已经领取过成为自有模板的不会同步删除**。
//
// - 对于自动领取的模板，由于已经下发，更改授权不会影响。
//
// - 如果要同步删除子客自有模板库中的模板，请使用OperateType=UPDATE+Available参数处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATES = "MissingParameter.Templates"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUTHTAG = "OperationDenied.AuthTag"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONAUTH = "ResourceNotFound.ApplicationAuth"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OperateChannelTemplateWithContext(ctx context.Context, request *OperateChannelTemplateRequest) (response *OperateChannelTemplateResponse, err error) {
    if request == nil {
        request = NewOperateChannelTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OperateChannelTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewOperateChannelTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewPrepareFlowsRequest() (request *PrepareFlowsRequest) {
    request = &PrepareFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "PrepareFlows")
    
    
    return
}

func NewPrepareFlowsResponse() (response *PrepareFlowsResponse) {
    response = &PrepareFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PrepareFlows
// 该接口 (PrepareFlows) 用于创建待发起文件
//
// 用户通过该接口进入签署流程发起的确认页面，进行发起信息二次确认， 如果确认则进行正常发起。
//
// 目前该接口只支持B2C，<font color='red'> **不建议使用，将会废弃**</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FLOWINFOS = "LimitExceeded.FlowInfos"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) PrepareFlows(request *PrepareFlowsRequest) (response *PrepareFlowsResponse, err error) {
    return c.PrepareFlowsWithContext(context.Background(), request)
}

// PrepareFlows
// 该接口 (PrepareFlows) 用于创建待发起文件
//
// 用户通过该接口进入签署流程发起的确认页面，进行发起信息二次确认， 如果确认则进行正常发起。
//
// 目前该接口只支持B2C，<font color='red'> **不建议使用，将会废弃**</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FLOWINFOS = "LimitExceeded.FlowInfos"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) PrepareFlowsWithContext(ctx context.Context, request *PrepareFlowsRequest) (response *PrepareFlowsResponse, err error) {
    if request == nil {
        request = NewPrepareFlowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PrepareFlows require credential")
    }

    request.SetContext(ctx)
    
    response = NewPrepareFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewSyncProxyOrganizationRequest() (request *SyncProxyOrganizationRequest) {
    request = &SyncProxyOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "SyncProxyOrganization")
    
    
    return
}

func NewSyncProxyOrganizationResponse() (response *SyncProxyOrganizationResponse) {
    response = &SyncProxyOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncProxyOrganization
// 此接口（SyncProxyOrganization）用于同步第三方平台子客企业信息，包括企业名称、企业营业执照、企业统一社会信用代码和法人姓名等，便于子客企业在企业激活过程中无需手动上传营业执照或补充企业信息。
//
// 
//
// 注意：
//
// 
//
// - **需要在<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>前同步的企业信息**, 否则会出现信息同步没有用的情形
//
// - **企业信息需要和营业执照信息对应**,  否则会出现激活过程验证不通过的问题
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7ec91b79a0a4860e77c9ff9f4a5f13ad/channel_SyncProxyOrganization2.png)
//
// 
//
// 
//
// - **企业统一社会信用代码**: 对应上图中的**1**
//
// - **第三方平台子客企业名称**: 对应上图中的**2**
//
// - **企业法定代表人的名字**:对应上图中的**3**
//
// - **企业详细住所**:对应上图中的**4**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UNIFORMSOCIALCREDITCODE = "InvalidParameter.UniformSocialCreditCode"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganization(request *SyncProxyOrganizationRequest) (response *SyncProxyOrganizationResponse, err error) {
    return c.SyncProxyOrganizationWithContext(context.Background(), request)
}

// SyncProxyOrganization
// 此接口（SyncProxyOrganization）用于同步第三方平台子客企业信息，包括企业名称、企业营业执照、企业统一社会信用代码和法人姓名等，便于子客企业在企业激活过程中无需手动上传营业执照或补充企业信息。
//
// 
//
// 注意：
//
// 
//
// - **需要在<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>前同步的企业信息**, 否则会出现信息同步没有用的情形
//
// - **企业信息需要和营业执照信息对应**,  否则会出现激活过程验证不通过的问题
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7ec91b79a0a4860e77c9ff9f4a5f13ad/channel_SyncProxyOrganization2.png)
//
// 
//
// 
//
// - **企业统一社会信用代码**: 对应上图中的**1**
//
// - **第三方平台子客企业名称**: 对应上图中的**2**
//
// - **企业法定代表人的名字**:对应上图中的**3**
//
// - **企业详细住所**:对应上图中的**4**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UNIFORMSOCIALCREDITCODE = "InvalidParameter.UniformSocialCreditCode"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationWithContext(ctx context.Context, request *SyncProxyOrganizationRequest) (response *SyncProxyOrganizationResponse, err error) {
    if request == nil {
        request = NewSyncProxyOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncProxyOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncProxyOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewSyncProxyOrganizationOperatorsRequest() (request *SyncProxyOrganizationOperatorsRequest) {
    request = &SyncProxyOrganizationOperatorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "SyncProxyOrganizationOperators")
    
    
    return
}

func NewSyncProxyOrganizationOperatorsResponse() (response *SyncProxyOrganizationOperatorsResponse) {
    response = &SyncProxyOrganizationOperatorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncProxyOrganizationOperators
// 此接口（SyncProxyOrganizationOperators）用于同步 第三方平台子客企业经办人列表，主要是同步经办人的离职状态。
//
// 子客Web控制台的组织架构管理，依赖于第三方应用平台的，无法在页面针对员工做新增/更新/离职等操作， 必须通过 API 来操作。 
//
// 
//
// - **新增员工的场景**:    通过本接口提前导入员工列表, 然后调用<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>分享给对应的员工进行实名, 新增员工后员工的状态为**未实名**, 通过链接实名后状态变为**已实名**, 已实名员工就可以参与合同的发起。
//
// 
//
// - **员工离职的场景**: 通过本接口将员工置为离职, 员工无法登录控制台和腾讯电子签小程序进行操作了,   同时给此员工分配的openid会被回收可以给其他新员工使用 (离职后员工数据会被置空,  再次加入公司会从零开始) ,  若员工信息有误可通过离职后在新增来解决,  离职员工状态为**离职**。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7a27a6bb0e4d39c2f6aa2a0b39946181/channel_SyncProxyOrganizationOperators.png)
//
// 
//
// **注**:    
//
// -  新增员工可以配置白名单限制注册使用对应openid的员工必须满足SyncProxyOrganizationOperators导入的(默认生成子客登录链接生成的链接可以任意员工点击注册绑定对应的openid), 此白名单需要咨询接入经理
//
// -  <font color='red'>超管和法人无法通过此接口离职</font>,  需要超管和法人将权限转移给其他人后才可通过此接口离职
//
// - 新增员工的场景同ID不同员工会覆盖掉上一个同ID的员工, 如果上一个员工已经实名则不会被覆盖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_PROXYORGANIZATIONOPERATOR = "LimitExceeded.ProxyOrganizationOperator"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationOperators(request *SyncProxyOrganizationOperatorsRequest) (response *SyncProxyOrganizationOperatorsResponse, err error) {
    return c.SyncProxyOrganizationOperatorsWithContext(context.Background(), request)
}

// SyncProxyOrganizationOperators
// 此接口（SyncProxyOrganizationOperators）用于同步 第三方平台子客企业经办人列表，主要是同步经办人的离职状态。
//
// 子客Web控制台的组织架构管理，依赖于第三方应用平台的，无法在页面针对员工做新增/更新/离职等操作， 必须通过 API 来操作。 
//
// 
//
// - **新增员工的场景**:    通过本接口提前导入员工列表, 然后调用<a href="https://qian.tencent.com/developers/partnerApis/accounts/CreateConsoleLoginUrl" target="_blank">生成子客登录链接</a>分享给对应的员工进行实名, 新增员工后员工的状态为**未实名**, 通过链接实名后状态变为**已实名**, 已实名员工就可以参与合同的发起。
//
// 
//
// - **员工离职的场景**: 通过本接口将员工置为离职, 员工无法登录控制台和腾讯电子签小程序进行操作了,   同时给此员工分配的openid会被回收可以给其他新员工使用 (离职后员工数据会被置空,  再次加入公司会从零开始) ,  若员工信息有误可通过离职后在新增来解决,  离职员工状态为**离职**。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7a27a6bb0e4d39c2f6aa2a0b39946181/channel_SyncProxyOrganizationOperators.png)
//
// 
//
// **注**:    
//
// -  新增员工可以配置白名单限制注册使用对应openid的员工必须满足SyncProxyOrganizationOperators导入的(默认生成子客登录链接生成的链接可以任意员工点击注册绑定对应的openid), 此白名单需要咨询接入经理
//
// -  <font color='red'>超管和法人无法通过此接口离职</font>,  需要超管和法人将权限转移给其他人后才可通过此接口离职
//
// - 新增员工的场景同ID不同员工会覆盖掉上一个同ID的员工, 如果上一个员工已经实名则不会被覆盖
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_PROXYORGANIZATIONOPERATOR = "LimitExceeded.ProxyOrganizationOperator"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationOperatorsWithContext(ctx context.Context, request *SyncProxyOrganizationOperatorsRequest) (response *SyncProxyOrganizationOperatorsResponse, err error) {
    if request == nil {
        request = NewSyncProxyOrganizationOperatorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncProxyOrganizationOperators require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncProxyOrganizationOperatorsResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFilesRequest() (request *UploadFilesRequest) {
    request = &UploadFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("essbasic", APIVersion, "UploadFiles")
    
    
    return
}

func NewUploadFilesResponse() (response *UploadFilesResponse) {
    response = &UploadFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadFiles
// 此接口（UploadFiles）文件上传。<br/>
//
// 
//
// 适用场景：用于合同，印章的文件上传。文件上传以后，
//
// 如果是PDF格式文件可配合<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowByFiles" target="_blank">用PDF文件创建签署流程</a>接口进行合同流程的发起
//
// 如果是其他类型可以配合<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务</a>接口转换成PDF文件
//
// 
//
// 注: 
//
// 1. 图片类型(png/jpg/jpeg)限制大小为5M以下, PDF/word/excel等其他格式限制大小为60M以下
//
// 2. <font color='red'>此接口调用时需要单独设置Domain请求域名 </font>,  联调开发环境为 <font color='red'>file.test.ess.tencent.cn</font>，正式环境需要设置为<font color='red'>file.ess.tencent.cn</font>，代码示例
//
// ```
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// ```
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-UploadFiles.mp4" target="_blank">【上传文件代码】编写示例</a><br>
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    return c.UploadFilesWithContext(context.Background(), request)
}

// UploadFiles
// 此接口（UploadFiles）文件上传。<br/>
//
// 
//
// 适用场景：用于合同，印章的文件上传。文件上传以后，
//
// 如果是PDF格式文件可配合<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateFlowByFiles" target="_blank">用PDF文件创建签署流程</a>接口进行合同流程的发起
//
// 如果是其他类型可以配合<a href="https://qian.tencent.com/developers/partnerApis/files/ChannelCreateConvertTaskApi" target="_blank">创建文件转换任务</a>接口转换成PDF文件
//
// 
//
// 注: 
//
// 1. 图片类型(png/jpg/jpeg)限制大小为5M以下, PDF/word/excel等其他格式限制大小为60M以下
//
// 2. <font color='red'>此接口调用时需要单独设置Domain请求域名 </font>,  联调开发环境为 <font color='red'>file.test.ess.tencent.cn</font>，正式环境需要设置为<font color='red'>file.ess.tencent.cn</font>，代码示例
//
// ```
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// ```
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/essbasic-UploadFiles.mp4" target="_blank">【上传文件代码】编写示例</a><br>
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UploadFilesWithContext(ctx context.Context, request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    if request == nil {
        request = NewUploadFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFilesResponse()
    err = c.Send(request, response)
    return
}
