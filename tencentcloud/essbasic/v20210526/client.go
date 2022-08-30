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
// 此接口（ChannelCancelMultiFlowSignQRCode）用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
func (c *Client) ChannelCancelMultiFlowSignQRCode(request *ChannelCancelMultiFlowSignQRCodeRequest) (response *ChannelCancelMultiFlowSignQRCodeResponse, err error) {
    return c.ChannelCancelMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// ChannelCancelMultiFlowSignQRCode
// 此接口（ChannelCancelMultiFlowSignQRCode）用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
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
// 指定需要批量撤销的签署流程Id，获取批量撤销链接
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChannelCreateBatchCancelFlowUrl(request *ChannelCreateBatchCancelFlowUrlRequest) (response *ChannelCreateBatchCancelFlowUrlResponse, err error) {
    return c.ChannelCreateBatchCancelFlowUrlWithContext(context.Background(), request)
}

// ChannelCreateBatchCancelFlowUrl
// 指定需要批量撤销的签署流程Id，获取批量撤销链接
//
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
// 渠道创建文件转换任务
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelCreateConvertTaskApi(request *ChannelCreateConvertTaskApiRequest) (response *ChannelCreateConvertTaskApiResponse, err error) {
    return c.ChannelCreateConvertTaskApiWithContext(context.Background(), request)
}

// ChannelCreateConvertTaskApi
// 渠道创建文件转换任务
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
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
// 接口（ChannelCreateFlowByFiles）用于渠道版通过文件创建签署流程。此接口不可直接使用，需要运营申请
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
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
func (c *Client) ChannelCreateFlowByFiles(request *ChannelCreateFlowByFilesRequest) (response *ChannelCreateFlowByFilesResponse, err error) {
    return c.ChannelCreateFlowByFilesWithContext(context.Background(), request)
}

// ChannelCreateFlowByFiles
// 接口（ChannelCreateFlowByFiles）用于渠道版通过文件创建签署流程。此接口不可直接使用，需要运营申请
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
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"
//  INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"
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
// 提交企业签署流程审批结果
//
// 
//
// 在通过接口(CreateFlowsByTemplates 或者ChannelCreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
//
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ChannelCreateFlowSignReview(request *ChannelCreateFlowSignReviewRequest) (response *ChannelCreateFlowSignReviewResponse, err error) {
    return c.ChannelCreateFlowSignReviewWithContext(context.Background(), request)
}

// ChannelCreateFlowSignReview
// 提交企业签署流程审批结果
//
// 
//
// 在通过接口(CreateFlowsByTemplates 或者ChannelCreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
//
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
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
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多扫签署流程二维码。
//
// 适用的模版仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模版，且模版中发起方没有填写控件。
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
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
func (c *Client) ChannelCreateMultiFlowSignQRCode(request *ChannelCreateMultiFlowSignQRCodeRequest) (response *ChannelCreateMultiFlowSignQRCodeResponse, err error) {
    return c.ChannelCreateMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// ChannelCreateMultiFlowSignQRCode
// 此接口（ChannelCreateMultiFlowSignQRCode）用于创建一码多扫签署流程二维码。
//
// 适用的模版仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模版，且模版中发起方没有填写控件。
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
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
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
// 渠道版查询转换任务状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ChannelGetTaskResultApi(request *ChannelGetTaskResultApiRequest) (response *ChannelGetTaskResultApiResponse, err error) {
    return c.ChannelGetTaskResultApiWithContext(context.Background(), request)
}

// ChannelGetTaskResultApi
// 渠道版查询转换任务状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
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
// 创建出证报告，返回报告 URL
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
func (c *Client) CreateChannelFlowEvidenceReport(request *CreateChannelFlowEvidenceReportRequest) (response *CreateChannelFlowEvidenceReportResponse, err error) {
    return c.CreateChannelFlowEvidenceReportWithContext(context.Background(), request)
}

// CreateChannelFlowEvidenceReport
// 创建出证报告，返回报告 URL
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.URL"
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
// 此接口（CreateConsoleLoginUrl）用于创建电子签控制台登录链接。若企业未激活，调用同步企业信息、同步经办人信息
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
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConsoleLoginUrl(request *CreateConsoleLoginUrlRequest) (response *CreateConsoleLoginUrlResponse, err error) {
    return c.CreateConsoleLoginUrlWithContext(context.Background(), request)
}

// CreateConsoleLoginUrl
// 此接口（CreateConsoleLoginUrl）用于创建电子签控制台登录链接。若企业未激活，调用同步企业信息、同步经办人信息
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
// 接口（CreateFlowsByTemplates）用于使用多个模板批量创建签署流程。当前可批量发起合同（签署流程）数量最大为20个。
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
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFlowsByTemplates(request *CreateFlowsByTemplatesRequest) (response *CreateFlowsByTemplatesResponse, err error) {
    return c.CreateFlowsByTemplatesWithContext(context.Background(), request)
}

// CreateFlowsByTemplates
// 接口（CreateFlowsByTemplates）用于使用多个模板批量创建签署流程。当前可批量发起合同（签署流程）数量最大为20个。
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
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"
//  INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
// 渠道通过图片为子客代创建印章，图片最大5m；此接口不可直接使用，需要运营申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_EXISTSAMESEALNAME = "FailedOperation.ExistSameSealName"
//  INTERNALERROR_SEALUPLOAD = "InternalError.SealUpload"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_IMAGE = "InvalidParameter.Image"
//  INVALIDPARAMETER_LIMITSEALNAME = "InvalidParameter.LimitSealName"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  MISSINGPARAMETER_ORGANIZATIONID = "MissingParameter.OrganizationId"
//  MISSINGPARAMETER_SEALIMAGE = "MissingParameter.SealImage"
//  MISSINGPARAMETER_SEALNAME = "MissingParameter.SealName"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
func (c *Client) CreateSealByImage(request *CreateSealByImageRequest) (response *CreateSealByImageResponse, err error) {
    return c.CreateSealByImageWithContext(context.Background(), request)
}

// CreateSealByImage
// 渠道通过图片为子客代创建印章，图片最大5m；此接口不可直接使用，需要运营申请
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_EXISTSAMESEALNAME = "FailedOperation.ExistSameSealName"
//  INTERNALERROR_SEALUPLOAD = "InternalError.SealUpload"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_IMAGE = "InvalidParameter.Image"
//  INVALIDPARAMETER_LIMITSEALNAME = "InvalidParameter.LimitSealName"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  MISSINGPARAMETER_ORGANIZATIONID = "MissingParameter.OrganizationId"
//  MISSINGPARAMETER_SEALIMAGE = "MissingParameter.SealImage"
//  MISSINGPARAMETER_SEALNAME = "MissingParameter.SealName"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
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
// 创建跳转小程序查看或签署的链接；自动签署的签署方不创建签署链接；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_GENERATETYPE = "InvalidParameter.GenerateType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateSignUrls(request *CreateSignUrlsRequest) (response *CreateSignUrlsResponse, err error) {
    return c.CreateSignUrlsWithContext(context.Background(), request)
}

// CreateSignUrls
// 创建跳转小程序查看或签署的链接；自动签署的签署方不创建签署链接；
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"
//  INVALIDPARAMETER_GENERATETYPE = "InvalidParameter.GenerateType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
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
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  LIMITEXCEEDED_FLOWIDS = "LimitExceeded.FlowIds"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDS = "MissingParameter.FlowIds"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWAPPROVERS = "ResourceNotFound.FlowApprovers"
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
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"
//  LIMITEXCEEDED_FLOWIDS = "LimitExceeded.FlowIds"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWIDS = "MissingParameter.FlowIds"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOWAPPROVERS = "ResourceNotFound.FlowApprovers"
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
// 根据签署流程信息批量获取资源下载链接，需合作企业先进行授权
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceUrlsByFlows(request *DescribeResourceUrlsByFlowsRequest) (response *DescribeResourceUrlsByFlowsResponse, err error) {
    return c.DescribeResourceUrlsByFlowsWithContext(context.Background(), request)
}

// DescribeResourceUrlsByFlows
// 根据签署流程信息批量获取资源下载链接，需合作企业先进行授权
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
// 通过此接口（DescribeTemplates）查询该企业在电子签渠道版中配置的有效模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
    return c.DescribeTemplatesWithContext(context.Background(), request)
}

// DescribeTemplates
// 通过此接口（DescribeTemplates）查询该企业在电子签渠道版中配置的有效模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
// 此接口（DescribeUsage）用于获取渠道所有合作企业流量消耗情况。
//
//  注: 此接口每日限频2次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATE = "InvalidParameter.Date"
//  LIMITEXCEEDED_CALLTIMES = "LimitExceeded.CallTimes"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DATE = "MissingParameter.Date"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUsage(request *DescribeUsageRequest) (response *DescribeUsageResponse, err error) {
    return c.DescribeUsageWithContext(context.Background(), request)
}

// DescribeUsage
// 此接口（DescribeUsage）用于获取渠道所有合作企业流量消耗情况。
//
//  注: 此接口每日限频2次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATE = "InvalidParameter.Date"
//  LIMITEXCEEDED_CALLTIMES = "LimitExceeded.CallTimes"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DATE = "MissingParameter.Date"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
// 此接口（GetDownloadFlowUrl）用于创建电子签批量下载地址，让合作企业进入控制台直接下载，支持客户合同（流程）按照自定义文件夹形式 分类下载。
//
// 当前接口限制最多合同（流程）50个.
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
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetDownloadFlowUrl(request *GetDownloadFlowUrlRequest) (response *GetDownloadFlowUrlResponse, err error) {
    return c.GetDownloadFlowUrlWithContext(context.Background(), request)
}

// GetDownloadFlowUrl
// 此接口（GetDownloadFlowUrl）用于创建电子签批量下载地址，让合作企业进入控制台直接下载，支持客户合同（流程）按照自定义文件夹形式 分类下载。
//
// 当前接口限制最多合同（流程）50个.
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
// 此接口（OperateChannelTemplate）用于渠道侧将模板库中的模板对合作企业进行查询和设置, 其中包括可见性的修改以及对合作企业的设置.
//
// 1、同步标识=select时：
//
// 返回渠道侧模板库当前模板的属性.
//
// 2、同步标识=update或者delete时：
//
// 对渠道子客进行模板库中模板授权,修改操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
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
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OperateChannelTemplate(request *OperateChannelTemplateRequest) (response *OperateChannelTemplateResponse, err error) {
    return c.OperateChannelTemplateWithContext(context.Background(), request)
}

// OperateChannelTemplate
// 此接口（OperateChannelTemplate）用于渠道侧将模板库中的模板对合作企业进行查询和设置, 其中包括可见性的修改以及对合作企业的设置.
//
// 1、同步标识=select时：
//
// 返回渠道侧模板库当前模板的属性.
//
// 2、同步标识=update或者delete时：
//
// 对渠道子客进行模板库中模板授权,修改操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
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
// 目前该接口只支持B2C，不建议使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FLOWINFOS = "LimitExceeded.FlowInfos"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
func (c *Client) PrepareFlows(request *PrepareFlowsRequest) (response *PrepareFlowsResponse, err error) {
    return c.PrepareFlowsWithContext(context.Background(), request)
}

// PrepareFlows
// 该接口 (PrepareFlows) 用于创建待发起文件
//
// 用户通过该接口进入签署流程发起的确认页面，进行发起信息二次确认， 如果确认则进行正常发起。
//
// 目前该接口只支持B2C，不建议使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FLOWINFOS = "LimitExceeded.FlowInfos"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
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
// 此接口（SyncProxyOrganization）用于同步渠道侧企业信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganization(request *SyncProxyOrganizationRequest) (response *SyncProxyOrganizationResponse, err error) {
    return c.SyncProxyOrganizationWithContext(context.Background(), request)
}

// SyncProxyOrganization
// 此接口（SyncProxyOrganization）用于同步渠道侧企业信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"
//  INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"
//  INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
// 此接口（SyncProxyOrganizationOperators）用于同步渠道合作企业经办人列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
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
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncProxyOrganizationOperators(request *SyncProxyOrganizationOperatorsRequest) (response *SyncProxyOrganizationOperatorsResponse, err error) {
    return c.SyncProxyOrganizationOperatorsWithContext(context.Background(), request)
}

// SyncProxyOrganizationOperators
// 此接口（SyncProxyOrganizationOperators）用于同步渠道合作企业经办人列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"
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
// 此接口（UploadFiles）用于文件上传。
//
// 调用时需要设置Domain 为 file.ess.tencent.cn
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    return c.UploadFilesWithContext(context.Background(), request)
}

// UploadFiles
// 此接口（UploadFiles）用于文件上传。
//
// 调用时需要设置Domain 为 file.ess.tencent.cn
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
