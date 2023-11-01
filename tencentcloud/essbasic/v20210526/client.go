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
// 用于批量撤销合同流程<br/>
//
// 适用场景：
//
// 如果某些合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。<br/>
//
// 通过签署流程编号批量撤销合同，单次最多支持撤销100份合同。只有合同的发起人或者发起方企业的超管/法人才能进行合同撤销。需要注意的是，若合同处于以下已终止状态，则不可撤销：<br/>
//
// - 已全部签署完成
//
// - 已拒签
//
// - 已过期
//
// - 已撤回
//
// - 拒绝填写
//
// - 已解除
//
// 
//
// <br/>
//
// 满足撤销条件的合同会发起异步撤销流程，而不满足撤销条件的合同将返回失败原因。合同撤销成功后，会通过合同状态为 CANCEL 的回调消息通知调用方。具体的回调消息内容可参考 <a href="https://qian.tencent.com/developers/partner/callback_types_contracts_sign" target="_blank">合同状态变更回调消息</a>。
//
// <br/><br/>
//
// 注:
//
// `如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同，`签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程</a>接口。
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
// 用于批量撤销合同流程<br/>
//
// 适用场景：
//
// 如果某些合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。<br/>
//
// 通过签署流程编号批量撤销合同，单次最多支持撤销100份合同。只有合同的发起人或者发起方企业的超管/法人才能进行合同撤销。需要注意的是，若合同处于以下已终止状态，则不可撤销：<br/>
//
// - 已全部签署完成
//
// - 已拒签
//
// - 已过期
//
// - 已撤回
//
// - 拒绝填写
//
// - 已解除
//
// 
//
// <br/>
//
// 满足撤销条件的合同会发起异步撤销流程，而不满足撤销条件的合同将返回失败原因。合同撤销成功后，会通过合同状态为 CANCEL 的回调消息通知调用方。具体的回调消息内容可参考 <a href="https://qian.tencent.com/developers/partner/callback_types_contracts_sign" target="_blank">合同状态变更回调消息</a>。
//
// <br/><br/>
//
// 注:
//
// `如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同，`签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程</a>接口。
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
// <ul><li> `可撤回合同状态` ：未全部签署完成</li>
//
// <li> `不撤回合同状态` ：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。</li></ul>
//
// 注:
//
// <ul><li>能撤回合同的只能是 `合同的发起人或者发起方企业的超管、法人`  </li>
//
// <li>签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>。</li></ul>
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
// <ul><li> `可撤回合同状态` ：未全部签署完成</li>
//
// <li> `不撤回合同状态` ：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。</li></ul>
//
// 注:
//
// <ul><li>能撤回合同的只能是 `合同的发起人或者发起方企业的超管、法人`  </li>
//
// <li>签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/partnerApis/startFlows/ChannelCreateReleaseFlow" target="_blank">发起解除合同流程接口</a>。</li></ul>
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
// 此接口（CancelMultiFlowSignQRCode）用于废除一码多扫流程签署二维码。
//
// 该接口所需的二维码ID，源自[创建一码多扫流程签署二维码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)生成的。
//
// 如果该二维码尚处于有效期内，可通过本接口将其设置为失效状态。
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
// 此接口（CancelMultiFlowSignQRCode）用于废除一码多扫流程签署二维码。
//
// 该接口所需的二维码ID，源自[创建一码多扫流程签署二维码](https://qian.tencent.com/developers/partnerApis/templates/ChannelCreateMultiFlowSignQRCode)生成的。
//
// 如果该二维码尚处于有效期内，可通过本接口将其设置为失效状态。
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
// 指定需要批量撤销的签署流程Id，获取批量撤销链接 - 不建议使用此接口，可使用ChannelBatchCancelFlows
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；
//
// 接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销;
//
// 
//
// 可以撤回：未全部签署完成
//
//  不可以撤回：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
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
// 指定需要批量撤销的签署流程Id，获取批量撤销链接 - 不建议使用此接口，可使用ChannelBatchCancelFlows
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；
//
// 接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销;
//
// 
//
// 可以撤回：未全部签署完成
//
//  不可以撤回：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 注意:
//
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
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
// 该接口用于发起合同后，生成个人用户的批量签署链接, 暂时不支持企业端签署 <br/>
//
// `注意：`<br/>
//
// `1. 该接口目前仅支持签署人类型是个人签署方的批量签署场景(ApproverType=PERSON)。` <br/>
//
// `2. 该接口可生成批量签署链接的C端签署人必须仅有手写签名和时间类型的签署控件，不支持填写控件 。` <br/>
//
// `3. 请确保C端签署人在批量签署合同中为待签署状态，如需顺序签署请待前一位参与人签署完成后，再创建该C端用户的签署链接。` <br/>
//
// `4. 该签署链接有效期为30分钟，过期后将失效，如需签署可重新创建批量签署链接 。` <br/>
//
// `5. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，不支持微信小程序嵌入。`<br/>
//
// 跳转到小程序的实现，参考微信官方文档(分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式)，如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>
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
// 该接口用于发起合同后，生成个人用户的批量签署链接, 暂时不支持企业端签署 <br/>
//
// `注意：`<br/>
//
// `1. 该接口目前仅支持签署人类型是个人签署方的批量签署场景(ApproverType=PERSON)。` <br/>
//
// `2. 该接口可生成批量签署链接的C端签署人必须仅有手写签名和时间类型的签署控件，不支持填写控件 。` <br/>
//
// `3. 请确保C端签署人在批量签署合同中为待签署状态，如需顺序签署请待前一位参与人签署完成后，再创建该C端用户的签署链接。` <br/>
//
// `4. 该签署链接有效期为30分钟，过期后将失效，如需签署可重新创建批量签署链接 。` <br/>
//
// `5. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，不支持微信小程序嵌入。`<br/>
//
// 跳转到小程序的实现，参考微信官方文档(分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式)，如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>
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
// - 参与人点击链接后需短信验证码才能查看合同内容。
//
// - 企业用户批量签署，需要传OrganizationName（参与方所在企业名称）参数生成签署链接，`请确保此企业已完成腾讯电子签企业认证`。
//
// - 个人批量签署，签名区`仅支持手写签名`。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
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
// - 参与人点击链接后需短信验证码才能查看合同内容。
//
// - 企业用户批量签署，需要传OrganizationName（参与方所在企业名称）参数生成签署链接，`请确保此企业已完成腾讯电子签企业认证`。
//
// - 个人批量签署，签名区`仅支持手写签名`。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
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
// 此接口（ChannelCreateBoundFlows）用于子客领取合同，经办人需要有相应的角色，合同不能重复领取。
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
// 此接口（ChannelCreateBoundFlows）用于子客领取合同，经办人需要有相应的角色，合同不能重复领取。
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
// 本接口（ChannelCreateEmbedWebUrl）用于创建常规模块嵌入web的链接
//
// 本接口支持创建：创建印章，创建模板，修改模板，预览模板，预览合同流程的web链接
//
// 进入web连接后与当前控制台操作保持一致
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
// 本接口（ChannelCreateEmbedWebUrl）用于创建常规模块嵌入web的链接
//
// 本接口支持创建：创建印章，创建模板，修改模板，预览模板，预览合同流程的web链接
//
// 进入web连接后与当前控制台操作保持一致
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
// 适用场景：
//
// 当通过模板或文件发起合同时，若未指定企业签署人信息，则可调用此接口动态补充签署人。同一签署人只允许补充一人，最终实际签署人取决于谁先领取合同完成签署。
//
// 
//
// 限制条件：
//
// 1. 本企业（发起方企业）企业签署人仅支持通过企业名称+姓名+手机号进行补充。
//
// 2. 个人签署人仅支持通过姓名+手机号进行补充。
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
// 适用场景：
//
// 当通过模板或文件发起合同时，若未指定企业签署人信息，则可调用此接口动态补充签署人。同一签署人只允许补充一人，最终实际签署人取决于谁先领取合同完成签署。
//
// 
//
// 限制条件：
//
// 1. 本企业（发起方企业）企业签署人仅支持通过企业名称+姓名+手机号进行补充。
//
// 2. 个人签署人仅支持通过姓名+手机号进行补充。
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
// ![image](https://dyn.ess.tencent.cn/guide/capi/ChannelCreateFlowByFiles.png)
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// ![image](https://dyn.ess.tencent.cn/guide/capi/ChannelCreateFlowByFiles.png)
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// 合同组是将多个合同签署流程组织在一起，多个合同同时创建，每个签署方得到一个签署链接，`一次完成合同组中多个合同的签署`。合同组的合同`不能拆分一个一个签署`，只能作为一个整体签署。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// 
//
// 
//
// **注**: 
//
// <ul>
//
// <li>此接口静默签(企业自动签)能力为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>合同组暂不支持抄送功能</li>
//
// <li>此接口需要依赖<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">文件上传接口</a>生成pdf资源编号（FileIds）进行使用。</li>
//
// </ul>
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// 
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// 合同组是将多个合同签署流程组织在一起，多个合同同时创建，每个签署方得到一个签署链接，`一次完成合同组中多个合同的签署`。合同组的合同`不能拆分一个一个签署`，只能作为一个整体签署。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// 
//
// 
//
// **注**: 
//
// <ul>
//
// <li>此接口静默签(企业自动签)能力为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>合同组暂不支持抄送功能</li>
//
// <li>此接口需要依赖<a href="https://qian.tencent.com/developers/partnerApis/files/UploadFiles" target="_blank">文件上传接口</a>生成pdf资源编号（FileIds）进行使用。</li>
//
// </ul>
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// 
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// 合同组是将多个合同签署流程组织在一起，多个合同同时创建，每个签署方得到一个签署链接，`一次完成合同组中多个合同的签署`。合同组的合同`不能拆分一个一个签署`，只能作为一个整体签署。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// **注**: 
//
// <ul>
//
// <li>此接口静默签(企业自动签)能力为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>合同组暂不支持抄送功能</li>
//
// </ul>
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// 
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// 合同组是将多个合同签署流程组织在一起，多个合同同时创建，每个签署方得到一个签署链接，`一次完成合同组中多个合同的签署`。合同组的合同`不能拆分一个一个签署`，只能作为一个整体签署。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// **注**: 
//
// <ul>
//
// <li>此接口静默签(企业自动签)能力为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>合同组暂不支持抄送功能</li>
//
// </ul>
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// 
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// 指定需要批量催办的签署流程ID，批量催办合同，最多100个。需要符合以下条件的合同才可被催办：
//
// 
//
// 1. 合同中当前状态为“待签署”的签署人是催办的对象
//
// 2. 每个合同只能催办一次
//
// 
//
// 注意：该接口无法直接调用，请联系客户经理申请使用。
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
// 指定需要批量催办的签署流程ID，批量催办合同，最多100个。需要符合以下条件的合同才可被催办：
//
// 
//
// 1. 合同中当前状态为“待签署”的签署人是催办的对象
//
// 2. 每个合同只能催办一次
//
// 
//
// 注意：该接口无法直接调用，请联系客户经理申请使用。
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
// 目前存在两种审核操作，签署审核，发起审核
//
// 签署审核：通过接口（CreateFlowsByTemplates或ChannelCreateFlowByFiles或ChannelCreatePrepareFlow）发起签署流程后，若指定了参数 NeedSignReview 为true,则可以调用此接口，指定operate=SignReview，提交企业内部签署审批结果；若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效
//
// 
//
// 发起审核：通过接口ChannelCreatePrepareFlow指定发起后需要审核，则可以通过调用此接口，指定operate=CreateReview，提交企业内部审批结果，可多次提交，当通过后，后续提交结果无效
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
// 目前存在两种审核操作，签署审核，发起审核
//
// 签署审核：通过接口（CreateFlowsByTemplates或ChannelCreateFlowByFiles或ChannelCreatePrepareFlow）发起签署流程后，若指定了参数 NeedSignReview 为true,则可以调用此接口，指定operate=SignReview，提交企业内部签署审批结果；若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效
//
// 
//
// 发起审核：通过接口ChannelCreatePrepareFlow指定发起后需要审核，则可以通过调用此接口，指定operate=CreateReview，提交企业内部审批结果，可多次提交，当通过后，后续提交结果无效
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
// 该接口用于发起合同后，生成C端签署人的签署链接<br/>
//
// 注意：该接口目前签署人类型仅支持个人签署方（PERSON）<br/>
//
// 注意：该接口可生成签署链接的C端签署人必须仅有手写签名和时间类型的签署控件<br/>
//
// 注意：该接口返回的签署链接是用于APP集成的场景，支持APP打开或浏览器直接打开，不支持微信小程序嵌入。微信小程序请使用小程序跳转或半屏弹窗的方式<br/>
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
// 该接口用于发起合同后，生成C端签署人的签署链接<br/>
//
// 注意：该接口目前签署人类型仅支持个人签署方（PERSON）<br/>
//
// 注意：该接口可生成签署链接的C端签署人必须仅有手写签名和时间类型的签署控件<br/>
//
// 注意：该接口返回的签署链接是用于APP集成的场景，支持APP打开或浏览器直接打开，不支持微信小程序嵌入。微信小程序请使用小程序跳转或半屏弹窗的方式<br/>
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
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多扫流程签署二维码。 
//
// 适用场景：无需填写签署人信息，可通过模板id生成签署二维码，签署人可通过扫描二维码补充签署信息进行实名签署。
//
// 常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 
//
// **本接口适用于发起方没有填写控件的 B2C或者单C模板**
//
// 
//
// **若是B2C模板,还要满足以下任意一个条件**
//
// 
//
// - 模板中配置的签署顺序是无序
//
// - B端企业的签署方式是静默签署
//
// - B端企业是非首位签署
//
// 
//
// 通过一码多扫二维码发起的合同，合同涉及到的回调消息可参考文档[合同发起及签署相关回调
//
// ]( https://qian.tencent.com/developers/partner/callback_types_contracts_sign)
//
// 
//
// 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档[签署二维码相关回调](https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83)
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
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多扫流程签署二维码。 
//
// 适用场景：无需填写签署人信息，可通过模板id生成签署二维码，签署人可通过扫描二维码补充签署信息进行实名签署。
//
// 常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 
//
// **本接口适用于发起方没有填写控件的 B2C或者单C模板**
//
// 
//
// **若是B2C模板,还要满足以下任意一个条件**
//
// 
//
// - 模板中配置的签署顺序是无序
//
// - B端企业的签署方式是静默签署
//
// - B端企业是非首位签署
//
// 
//
// 通过一码多扫二维码发起的合同，合同涉及到的回调消息可参考文档[合同发起及签署相关回调
//
// ]( https://qian.tencent.com/developers/partner/callback_types_contracts_sign)
//
// 
//
// 用户通过签署二维码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档[签署二维码相关回调](https://qian.tencent.com/developers/partner/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83)
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
// 通过此接口，创建小程序批量签署链接，可以创建企业批量签署链接，员工只需点击链接即可跳转至控制台进行批量签署。
//
// 
//
// 注：
//
// - 员工必须在企业下完成实名认证，且需作为批量签署合同的签署方或者领取方。
//
// - 仅支持传入待签署或者待领取的合同，待填写暂不支持。
//
// - 员工批量签署，支持多种签名方式，包括手写签名、临摹签名、系统签名、个人印章，暂不支持签批控件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ChannelCreateOrganizationBatchSignUrl(request *ChannelCreateOrganizationBatchSignUrlRequest) (response *ChannelCreateOrganizationBatchSignUrlResponse, err error) {
    return c.ChannelCreateOrganizationBatchSignUrlWithContext(context.Background(), request)
}

// ChannelCreateOrganizationBatchSignUrl
// 通过此接口，创建小程序批量签署链接，可以创建企业批量签署链接，员工只需点击链接即可跳转至控制台进行批量签署。
//
// 
//
// 注：
//
// - 员工必须在企业下完成实名认证，且需作为批量签署合同的签署方或者领取方。
//
// - 仅支持传入待签署或者待领取的合同，待填写暂不支持。
//
// - 员工批量签署，支持多种签名方式，包括手写签名、临摹签名、系统签名、个人印章，暂不支持签批控件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
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
// 创建预发起合同
//
// 通过此接口指定：合同，签署人，填写控件信息，生成预创建合同链接，点击后跳转到web页面完成合同创建并发起
//
// 可指定合同信息不可更改，签署人信息不可更改
//
// 合同发起后，填写及签署流程与现有操作流程一致
//
// 注意：目前仅支持模板发起
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
// 创建预发起合同
//
// 通过此接口指定：合同，签署人，填写控件信息，生成预创建合同链接，点击后跳转到web页面完成合同创建并发起
//
// 可指定合同信息不可更改，签署人信息不可更改
//
// 合同发起后，填写及签署流程与现有操作流程一致
//
// 注意：目前仅支持模板发起
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
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreatePreparedPersonalEsign(request *ChannelCreatePreparedPersonalEsignRequest) (response *ChannelCreatePreparedPersonalEsignResponse, err error) {
    return c.ChannelCreatePreparedPersonalEsignWithContext(context.Background(), request)
}

// ChannelCreatePreparedPersonalEsign
// 本接口（ChannelCreatePreparedPersonalEsign）用于创建导入个人印章（处方单场景专用，使用此接口请与客户经理确认）。
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
// 发起解除协议，主要应用场景为：基于一份已经签署的合同，进行解除操作。
//
// 合同发起人必须在电子签已经进行实名。
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
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelCreateReleaseFlow(request *ChannelCreateReleaseFlowRequest) (response *ChannelCreateReleaseFlowResponse, err error) {
    return c.ChannelCreateReleaseFlowWithContext(context.Background(), request)
}

// ChannelCreateReleaseFlow
// 发起解除协议，主要应用场景为：基于一份已经签署的合同，进行解除操作。
//
// 合同发起人必须在电子签已经进行实名。
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
// 适用场景2：创建当前企业的自定义角色，并且创建时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
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
// 适用场景2：创建当前企业的自定义角色，并且创建时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
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
// 企业方可以通过此接口获取个人用户开启自动签的跳转链接
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
// 企业方可以通过此接口获取个人用户开启自动签的跳转链接
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
// 通过此接口，绑定员工角色，支持以电子签userId、客户系统userId两种方式调用。
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
// 通过此接口，绑定员工角色，支持以电子签userId、客户系统userId两种方式调用。
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
// 注意：系统角色不可删除。
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
// 注意：系统角色不可删除。
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
// 删除指定印章下多个授权信息
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelDeleteSealPolicies(request *ChannelDeleteSealPoliciesRequest) (response *ChannelDeleteSealPoliciesResponse, err error) {
    return c.ChannelDeleteSealPoliciesWithContext(context.Background(), request)
}

// ChannelDeleteSealPolicies
// 删除指定印章下多个授权信息
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
// 用于分页查询企业员工信息列表，支持设置过滤条件以筛选员工查询结果。
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
// 用于分页查询企业员工信息列表，支持设置过滤条件以筛选员工查询结果。
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
// 查询流程填写控件内容，可以根据流程Id查询该流程相关联的填写控件信息和填写内容。 注意：使用此接口前，需要在【企业应用管理】-【应用集成】-【第三方应用管理】中开通【下载应用内全量合同文件及内容数据】功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
func (c *Client) ChannelDescribeFlowComponents(request *ChannelDescribeFlowComponentsRequest) (response *ChannelDescribeFlowComponentsResponse, err error) {
    return c.ChannelDescribeFlowComponentsWithContext(context.Background(), request)
}

// ChannelDescribeFlowComponents
// 查询流程填写控件内容，可以根据流程Id查询该流程相关联的填写控件信息和填写内容。 注意：使用此接口前，需要在【企业应用管理】-【应用集成】-【第三方应用管理】中开通【下载应用内全量合同文件及内容数据】功能。
//
// 可能返回的错误码:
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
// 查询子客企业电子印章，需要操作者具有管理印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数，只返回启用的印章。
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
// 查询子客企业电子印章，需要操作者具有管理印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数，只返回启用的印章。
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
// 相关系统默认角色说明可参考文档：https://cloud.tencent.com/document/product/1323/61355
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
// 相关系统默认角色说明可参考文档：https://cloud.tencent.com/document/product/1323/61355
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
// 适用场景：已创建一个文件转换任务，想查询该文件转换任务的状态，或获取转换后的文件资源Id。<br />
//
// 注：
//
// 1. `大文件转换所需的时间可能会比较长`
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
// 适用场景：已创建一个文件转换任务，想查询该文件转换任务的状态，或获取转换后的文件资源Id。<br />
//
// 注：
//
// 1. `大文件转换所需的时间可能会比较长`
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
// 适用场景2：更新当前企业的自定义角色的权限信息，更新时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
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
// 适用场景2：更新当前企业的自定义角色的权限信息，更新时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 ChannelDescribeRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
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
// 本接口（ChannelUpdateSealStatus）用于第三方应用平台为子客企业更新印章状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ChannelUpdateSealStatus(request *ChannelUpdateSealStatusRequest) (response *ChannelUpdateSealStatusResponse, err error) {
    return c.ChannelUpdateSealStatusWithContext(context.Background(), request)
}

// ChannelUpdateSealStatus
// 本接口（ChannelUpdateSealStatus）用于第三方应用平台为子客企业更新印章状态
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
// <ul><li>使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。</li>
//
// <li>操作人必须是`发起方或者签署方企业的(非走授权书认证)法人或者超管`。</li>
//
// <li>合同流程必须所有参与方`已经签署完成`。</li>
//
// <li>出证过程需一定时间，建议在`提交出证任务后的24小时之后`，通过<a href="https://qian.tencent.com/developers/partnerApis/certificate/DescribeChannelFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>接口进行查询执行结果和出证报告的下载URL。</li></ul>
//
// 
//
// <svg id="SvgjsSvg1006" width="262" height="229" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.com/svgjs"><defs id="SvgjsDefs1007"><pattern patternUnits="userSpaceOnUse" id="pattern_mark_0" width="300" height="300"><text x="150" y="100" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><pattern patternUnits="userSpaceOnUse" id="pattern_mark_1" width="300" height="300"><text x="150" y="200" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><marker id="SvgjsMarker1021" markerWidth="12" markerHeight="8" refX="9" refY="4" viewBox="0 0 12 8" orient="auto" markerUnits="userSpaceOnUse" stroke-dasharray="0,0"><path id="SvgjsPath1022" d="M0,0 L12,4 L0,8 L0,0" fill="#323232" stroke="#323232" stroke-width="1"></path></marker></defs><rect id="svgbackgroundid" width="262" height="229" fill="transparent"></rect><rect id="SvgjsRect1009" width="262" height="229" fill="url(#pattern_mark_0)"></rect><rect id="SvgjsRect1010" width="262" height="229" fill="url(#pattern_mark_1)"></rect><g id="SvgjsG1011" transform="translate(31.75,25)"><path id="SvgjsPath1012" d="M 0 0L 198 0L 198 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1013"><text id="SvgjsText1014" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="178px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1015" dy="16" x="99"><tspan id="SvgjsTspan1016" style="text-decoration:;fill: rgb(28, 30, 33);">CreateFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1017" dy="16" x="99"><tspan id="SvgjsTspan1018" style="text-decoration:;fill: rgb(51, 51, 51);">提交申请出证报告任务</tspan></tspan></text></g></g><g id="SvgjsG1019"><path id="SvgjsPath1020" d="M130.75 84.5L130.75 114.5L130.75 114.5L130.75 143.2" stroke="#323232" stroke-width="1" fill="none" marker-end="url(#SvgjsMarker1021)"></path></g><g id="SvgjsG1023" transform="translate(25,145)"><path id="SvgjsPath1024" d="M 0 0L 211.5 0L 211.5 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1025"><text id="SvgjsText1026" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="192px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1027" dy="16" x="106"><tspan id="SvgjsTspan1028" style="text-decoration:;fill: rgb(28, 30, 33);">DescribeFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1029" dy="16" x="106"><tspan id="SvgjsTspan1030" style="text-decoration:;fill: rgb(51, 51, 51);">获取出证报告任务执行结果</tspan></tspan></text></g></g></svg>
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
// <ul><li>使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。</li>
//
// <li>操作人必须是`发起方或者签署方企业的(非走授权书认证)法人或者超管`。</li>
//
// <li>合同流程必须所有参与方`已经签署完成`。</li>
//
// <li>出证过程需一定时间，建议在`提交出证任务后的24小时之后`，通过<a href="https://qian.tencent.com/developers/partnerApis/certificate/DescribeChannelFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>接口进行查询执行结果和出证报告的下载URL。</li></ul>
//
// 
//
// <svg id="SvgjsSvg1006" width="262" height="229" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.com/svgjs"><defs id="SvgjsDefs1007"><pattern patternUnits="userSpaceOnUse" id="pattern_mark_0" width="300" height="300"><text x="150" y="100" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><pattern patternUnits="userSpaceOnUse" id="pattern_mark_1" width="300" height="300"><text x="150" y="200" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><marker id="SvgjsMarker1021" markerWidth="12" markerHeight="8" refX="9" refY="4" viewBox="0 0 12 8" orient="auto" markerUnits="userSpaceOnUse" stroke-dasharray="0,0"><path id="SvgjsPath1022" d="M0,0 L12,4 L0,8 L0,0" fill="#323232" stroke="#323232" stroke-width="1"></path></marker></defs><rect id="svgbackgroundid" width="262" height="229" fill="transparent"></rect><rect id="SvgjsRect1009" width="262" height="229" fill="url(#pattern_mark_0)"></rect><rect id="SvgjsRect1010" width="262" height="229" fill="url(#pattern_mark_1)"></rect><g id="SvgjsG1011" transform="translate(31.75,25)"><path id="SvgjsPath1012" d="M 0 0L 198 0L 198 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1013"><text id="SvgjsText1014" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="178px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1015" dy="16" x="99"><tspan id="SvgjsTspan1016" style="text-decoration:;fill: rgb(28, 30, 33);">CreateFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1017" dy="16" x="99"><tspan id="SvgjsTspan1018" style="text-decoration:;fill: rgb(51, 51, 51);">提交申请出证报告任务</tspan></tspan></text></g></g><g id="SvgjsG1019"><path id="SvgjsPath1020" d="M130.75 84.5L130.75 114.5L130.75 114.5L130.75 143.2" stroke="#323232" stroke-width="1" fill="none" marker-end="url(#SvgjsMarker1021)"></path></g><g id="SvgjsG1023" transform="translate(25,145)"><path id="SvgjsPath1024" d="M 0 0L 211.5 0L 211.5 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1025"><text id="SvgjsText1026" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="192px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1027" dy="16" x="106"><tspan id="SvgjsTspan1028" style="text-decoration:;fill: rgb(28, 30, 33);">DescribeFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1029" dy="16" x="106"><tspan id="SvgjsTspan1030" style="text-decoration:;fill: rgb(51, 51, 51);">获取出证报告任务执行结果</tspan></tspan></text></g></g></svg>
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
// 此接口（CreateChannelOrganizationInfoChangeUrl）用于创建子客企业信息变更链接（需要在移动端打开，会跳转到微信小程序），支持创建企业超管变更链接或企业基础信息变更链接，通过入参 ChangeType 指定。
//
// 
//
// 
//
// <h3 id="1-企业超管变更">1. 企业超管变更</h3>
//
// 
//
// <p>换成企业的其他员工来当超管</p>
//
// 
//
// <h3 id="2-企业基础信息变更">2. 企业基础信息变更</h3>
//
// 
//
// <h4 id="可以变动">可以变动</h4>
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
// <h4 id="不可变动">不可变动</h4>
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
// <li>印章:   会删除所有的印章所有的机构公章和合同专用章,  然后用新企业名称生成新的机构公章 和合同专用章,  而法人章, 财务专用章和人事专用章不会处理</li>
//
// <li>证书:   企业证书会重新请求CA机构用新企业名称生成新的证书</li>
//
// </ul>
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateChannelOrganizationInfoChangeUrl(request *CreateChannelOrganizationInfoChangeUrlRequest) (response *CreateChannelOrganizationInfoChangeUrlResponse, err error) {
    return c.CreateChannelOrganizationInfoChangeUrlWithContext(context.Background(), request)
}

// CreateChannelOrganizationInfoChangeUrl
// 此接口（CreateChannelOrganizationInfoChangeUrl）用于创建子客企业信息变更链接（需要在移动端打开，会跳转到微信小程序），支持创建企业超管变更链接或企业基础信息变更链接，通过入参 ChangeType 指定。
//
// 
//
// 
//
// <h3 id="1-企业超管变更">1. 企业超管变更</h3>
//
// 
//
// <p>换成企业的其他员工来当超管</p>
//
// 
//
// <h3 id="2-企业基础信息变更">2. 企业基础信息变更</h3>
//
// 
//
// <h4 id="可以变动">可以变动</h4>
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
// <h4 id="不可变动">不可变动</h4>
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
// <li>印章:   会删除所有的印章所有的机构公章和合同专用章,  然后用新企业名称生成新的机构公章 和合同专用章,  而法人章, 财务专用章和人事专用章不会处理</li>
//
// <li>证书:   企业证书会重新请求CA机构用新企业名称生成新的证书</li>
//
// </ul>
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
// <td>进入员认证并加入企业流程</td>
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
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRNOTSYNCPROXYORGANIZATION = "FailedOperation.ErrNotSyncProxyOrganization"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZATIONTYPE = "InvalidParameter.AuthorizationType"
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
// <td>进入员认证并加入企业流程</td>
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
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRNOTSYNCPROXYORGANIZATION = "FailedOperation.ErrNotSyncProxyOrganization"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZATIONTYPE = "InvalidParameter.AuthorizationType"
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
// **整体的逻辑如下**
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateFlowsByTemplates.png)
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// **整体的逻辑如下**
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateFlowsByTemplates.png)
//
// 
//
// **可以作为发起方和签署方的角色列表**
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
// <td>第三方子企业B(不指定经办人)</td>
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
// <td>第三方子企业B员工</td>
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
// <td>个人/自然人</td>
//
// </tr>
//
// 
//
// <tr>
//
// <td>场景五</td>
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
// **注**: 
//
// `1. 发起合同时候,  作为发起方的第三方子企业A员工的企业和员工必须经过实名, 而作为签署方的第三方子企业A员工/个人/自然人/SaaS平台企业员工/第三方子企业B员工企业中的企业和个人/员工可以未实名`
//
// `2. 不同类型的签署方传参不同, 可以参考开发者中心的FlowApproverInfo结构体说明`
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
// **主要使用场景可以更加EndPoint分类如下**
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
// |APP| 贵方App跳转腾讯电子签小程序签署场景|  贵方App直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方App的场景<br>跳转到腾讯电子签小程序的实现可以参考微信的官方文档:<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html" target="_blank">开放能力/打开 App</a> <br> 示例: pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3 |
//
// |APP| 贵方小程序跳转腾讯电子签小程序签署场景|  贵方App直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方小程序的场景<br>跳转到腾讯电子签小程序的实现可以参考微信官方文档<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html" target="_blank">全屏方式</a>和<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html " target="_blank">半屏方式</a><br>此时返回的SignUrl就是官方文档中的path<br> 示例:pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3  |
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
// **主要使用场景可以更加EndPoint分类如下**
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
// |APP| 贵方App跳转腾讯电子签小程序签署场景|  贵方App直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方App的场景<br>跳转到腾讯电子签小程序的实现可以参考微信的官方文档:<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html" target="_blank">开放能力/打开 App</a> <br> 示例: pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3 |
//
// |APP| 贵方小程序跳转腾讯电子签小程序签署场景|  贵方App直接跳转到小程序后, 在腾讯电子签小程序签署完成后返回贵方小程序的场景<br>跳转到腾讯电子签小程序的实现可以参考微信官方文档<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html" target="_blank">全屏方式</a>和<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html " target="_blank">半屏方式</a><br>此时返回的SignUrl就是官方文档中的path<br> 示例:pages/guide?from=default&where=mini& to=CONTRACT_DETAIL& id=yDwiBUUc*duRvquCSX8wd& shortKey=yDwivUA**W1yRsTre3  |
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
// 查询出证报告，返回报告 URL。
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
// 查询出证报告，返回报告 URL。
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
// 1. 企业自动签
//
// 2. 企业与港澳台居民签署合同
//
// 3. 使用手机号验证签署方身份
//
// 4. 骑缝章
//
// 5. 拓宽签署方年龄限制
//
// 6. 下载企业合同/文件
//
// 
//
// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
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
// 1. 企业自动签
//
// 2. 企业与港澳台居民签署合同
//
// 3. 使用手机号验证签署方身份
//
// 4. 骑缝章
//
// 5. 拓宽签署方年龄限制
//
// 6. 下载企业合同/文件
//
// 
//
// 注: 此接口 参数Agent. ProxyOperator.OpenId 需要传递超管或者法人的OpenId
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
// 此接口（DescribeFlowDetailInfo）用于查询合同(签署流程)的详细信息。
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
// 此接口（DescribeFlowDetailInfo）用于查询合同(签署流程)的详细信息。
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
// 获取合同流程PDF的下载链接，可以下载签署中、签署完的此子企业创建的合同
//
// 
//
// **注意**:   
//
// 有两种开通权限的途径
//
// 
//
// **第一种**:   需第三方应用的子企业登录控制台进行授权,  授权在**企业中心**的**授权管理**区域,  界面如下图
//
// 授权过程需要**子企业超管**扫描跳转到电子签小程序签署<<渠道端下载渠道子客合同功能授权委托书>>
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_DescribeResourceUrlsByFlows2.png)
//
// 
//
// **第二种**: 第三方应用的配置接口打开全第三个应用下的所有自己起开通, 需要**渠道方企业的超管**扫描二维码跳转到电子签小程序签署 <<渠道端下载渠道子客合同功能开通知情同意书>>
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_DescribeResourceUrlsByFlows_appilications2.png)
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
// 获取合同流程PDF的下载链接，可以下载签署中、签署完的此子企业创建的合同
//
// 
//
// **注意**:   
//
// 有两种开通权限的途径
//
// 
//
// **第一种**:   需第三方应用的子企业登录控制台进行授权,  授权在**企业中心**的**授权管理**区域,  界面如下图
//
// 授权过程需要**子企业超管**扫描跳转到电子签小程序签署<<渠道端下载渠道子客合同功能授权委托书>>
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_DescribeResourceUrlsByFlows2.png)
//
// 
//
// **第二种**: 第三方应用的配置接口打开全第三个应用下的所有自己起开通, 需要**渠道方企业的超管**扫描二维码跳转到电子签小程序签署 <<渠道端下载渠道子客合同功能开通知情同意书>>
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/channel_DescribeResourceUrlsByFlows_appilications2.png)
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
// > **适用场景** 
//
// >
//
// >  该接口常用来配合“使用模板创建签署流程”接口作为前置的接口使用。 
//
// >  一个模板通常会包含以下结构信息
//
// >- 模板基本信息
//
// >- 发起方参与信息Promoter、签署参与方 Recipients，后者会在模板发起合同时用于指定参与方
//
// >- 填写控件 Components
//
// >- 签署控件 SignComponents
//
// >- 生成模板的文件基础信息 FileInfos
//
// 
//
// 注意：
//
// 
//
// >1. 查询条件TemplateId、TemplateName与ChannelTemplateId可同时存在，即可查询同时满足这些条件的模板。
//
// >2. TemplateId 和TemplateIds互为独立，若两个参数都传入，则以TemplateId为准
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
// > **适用场景** 
//
// >
//
// >  该接口常用来配合“使用模板创建签署流程”接口作为前置的接口使用。 
//
// >  一个模板通常会包含以下结构信息
//
// >- 模板基本信息
//
// >- 发起方参与信息Promoter、签署参与方 Recipients，后者会在模板发起合同时用于指定参与方
//
// >- 填写控件 Components
//
// >- 签署控件 SignComponents
//
// >- 生成模板的文件基础信息 FileInfos
//
// 
//
// 注意：
//
// 
//
// >1. 查询条件TemplateId、TemplateName与ChannelTemplateId可同时存在，即可查询同时满足这些条件的模板。
//
// >2. TemplateId 和TemplateIds互为独立，若两个参数都传入，则以TemplateId为准
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
// 此接口（DescribeUsage）用于获取第三方平台所有合作企业流量消耗情况。
//
//  注: 此接口每日限频50次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
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
// 此接口（DescribeUsage）用于获取第三方平台所有合作企业流量消耗情况。
//
//  注: 此接口每日限频50次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
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
// <li>仅支持下载 **本企业** 下合同，链接会 **登录企业控制台** </li>
//
// <li> **链接仅可使用一次**，不可重复使用</li>
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
// <li>仅支持下载 **本企业** 下合同，链接会 **登录企业控制台** </li>
//
// <li> **链接仅可使用一次**，不可重复使用</li>
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
// 修改（操作）企业扩展服务 ，企业经办人需要是企业超管或者法人。
//
// 
//
// 跳转小程序的几种方式：主要是设置不同的EndPoint
//
// 1. 通过链接Url直接跳转到小程序，不需要返回
//
// 设置EndPoint为WEIXINAPP，得到链接打开即可。
//
// 
//
// 2. 客户App直接跳转到小程序-->腾讯电子签小程序操作完成-->返回App
//
// 跳转到小程序的实现，参考官方文档
//
// https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 4. 客户小程序直接跳到电子签小程序-->腾讯电子签小程序操作完成--->回到客户小程序
//
// 跳转到小程序的实现，参考官方文档（分为全屏、半屏两种方式）
//
// 全屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html）
//
// 半屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html）
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 其中小程序的原始Id如下，或者查看小程序信息自助获取。
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyExtendedService(request *ModifyExtendedServiceRequest) (response *ModifyExtendedServiceResponse, err error) {
    return c.ModifyExtendedServiceWithContext(context.Background(), request)
}

// ModifyExtendedService
// 修改（操作）企业扩展服务 ，企业经办人需要是企业超管或者法人。
//
// 
//
// 跳转小程序的几种方式：主要是设置不同的EndPoint
//
// 1. 通过链接Url直接跳转到小程序，不需要返回
//
// 设置EndPoint为WEIXINAPP，得到链接打开即可。
//
// 
//
// 2. 客户App直接跳转到小程序-->腾讯电子签小程序操作完成-->返回App
//
// 跳转到小程序的实现，参考官方文档
//
// https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 4. 客户小程序直接跳到电子签小程序-->腾讯电子签小程序操作完成--->回到客户小程序
//
// 跳转到小程序的实现，参考官方文档（分为全屏、半屏两种方式）
//
// 全屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html）
//
// 半屏方式：
//
// （https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html）
//
// 其中小程序的原始Id，请联系<对接技术人员>获取，或者查看小程序信息自助获取。
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 其中小程序的原始Id如下，或者查看小程序信息自助获取。
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
//  INTERNALERROR = "InternalError"
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
// 此接口（OperateChannelTemplate）用于针对第三方应用平台模板库中的模板对子客企业可见性的查询和设置。
//
// 
//
// > **使用场景**
//
// >
//
// >  1：查询 OperateType=SELECT
//
// > - 查询第三方应用平台模板库的可见性以及授权的子客列表。
//
// >
//
// >  2：修改部分子客授权 OperateType=UPDATE
//
// > - 对子客企业进行模板库中模板可见性的进行修改操作。
//
// >- 当模板未发布时，可以修改可见性AuthTag（part/all），当模板发布后，不可做此修改
//
// > - 若模板已发布且可见性AuthTag是part，可以通过ProxyOrganizationOpenIds增加子客的授权范围。如果是自动领取的模板，增加授权范围后会自动下发。
//
// >
//
// >  3：取消部分子客授权 OperateType=DELETE
//
// > - 对子客企业进行模板库中模板可见性的进行删除操作。
//
// > - 主要对于手动领取的模板，去除授权后子客在模板库中看不到，就无法再领取了。但是已经领取过成为自有模板的不会同步删除。
//
// > - 对于自动领取的模板，由于已经下发，更改授权不会影响。
//
// > - 如果要同步删除子客自有模板库中的模板，请使用OperateType=UPDATE+Available参数处理。
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
// 此接口（OperateChannelTemplate）用于针对第三方应用平台模板库中的模板对子客企业可见性的查询和设置。
//
// 
//
// > **使用场景**
//
// >
//
// >  1：查询 OperateType=SELECT
//
// > - 查询第三方应用平台模板库的可见性以及授权的子客列表。
//
// >
//
// >  2：修改部分子客授权 OperateType=UPDATE
//
// > - 对子客企业进行模板库中模板可见性的进行修改操作。
//
// >- 当模板未发布时，可以修改可见性AuthTag（part/all），当模板发布后，不可做此修改
//
// > - 若模板已发布且可见性AuthTag是part，可以通过ProxyOrganizationOpenIds增加子客的授权范围。如果是自动领取的模板，增加授权范围后会自动下发。
//
// >
//
// >  3：取消部分子客授权 OperateType=DELETE
//
// > - 对子客企业进行模板库中模板可见性的进行删除操作。
//
// > - 主要对于手动领取的模板，去除授权后子客在模板库中看不到，就无法再领取了。但是已经领取过成为自有模板的不会同步删除。
//
// > - 对于自动领取的模板，由于已经下发，更改授权不会影响。
//
// > - 如果要同步删除子客自有模板库中的模板，请使用OperateType=UPDATE+Available参数处理。
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
// 目前该接口只支持B2C，不建议使用，将会**废弃**。
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
// 目前该接口只支持B2C，不建议使用，将会**废弃**。
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
// 此接口（SyncProxyOrganization）用于同步第三方平台子客企业信息，包括企业名称，企业营业执照，企业统一社会信用代码和法人姓名等，便于子客企业在企业激活过程中无需手动上传营业执照或补充企业信息。注意：
//
// 1. 需要在子客企业激活前调用该接口，如果您的企业已经提交企业信息或者企业已经激活，同步的企业信息将不会生效。
//
// 2. 如果您同时传递了营业执照信息和企业名称等信息，在认证过程中将以营业执照中的企业信息为准，请注意企业信息需要和营业执照信息对应。
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
// 此接口（SyncProxyOrganization）用于同步第三方平台子客企业信息，包括企业名称，企业营业执照，企业统一社会信用代码和法人姓名等，便于子客企业在企业激活过程中无需手动上传营业执照或补充企业信息。注意：
//
// 1. 需要在子客企业激活前调用该接口，如果您的企业已经提交企业信息或者企业已经激活，同步的企业信息将不会生效。
//
// 2. 如果您同时传递了营业执照信息和企业名称等信息，在认证过程中将以营业执照中的企业信息为准，请注意企业信息需要和营业执照信息对应。
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
// 此接口（SyncProxyOrganizationOperators）用于同步 第三方平台子客企业经办人列表，主要是同步经办人的离职状态。子客Web控制台的组织架构管理，是依赖于第三方应用平台的，无法针对员工做新增/更新/离职等操作。 
//
// 若经办人信息有误，或者需要修改，也可以先将之前的经办人做离职操作，然后重新使用控制台链接CreateConsoleLoginUrl让经办人重新实名。
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
// 此接口（SyncProxyOrganizationOperators）用于同步 第三方平台子客企业经办人列表，主要是同步经办人的离职状态。子客Web控制台的组织架构管理，是依赖于第三方应用平台的，无法针对员工做新增/更新/离职等操作。 
//
// 若经办人信息有误，或者需要修改，也可以先将之前的经办人做离职操作，然后重新使用控制台链接CreateConsoleLoginUrl让经办人重新实名。
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
// 1. `图片类型(png/jpg/jpeg)限制大小为5M以下, PDF/word/excel等其他格式限制大小为60M以下`
//
// 2. `联调开发环境调用时需要设置Domain接口请求域名为 file.test.ess.tencent.cn，正式环境需要设置为file.ess.tencent.cn，代码示例`
//
// ```
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// ```
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
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
// 1. `图片类型(png/jpg/jpeg)限制大小为5M以下, PDF/word/excel等其他格式限制大小为60M以下`
//
// 2. `联调开发环境调用时需要设置Domain接口请求域名为 file.test.ess.tencent.cn，正式环境需要设置为file.ess.tencent.cn，代码示例`
//
// ```
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// ```
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
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
