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

package v20201111

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-11"

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


func NewCancelFlowRequest() (request *CancelFlowRequest) {
    request = &CancelFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CancelFlow")
    
    
    return
}

func NewCancelFlowResponse() (response *CancelFlowResponse) {
    response = &CancelFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelFlow
// 用于撤销签署流程
//
// 适用场景：如果某个合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 注：如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER_CANCELREASON = "MissingParameter.CancelReason"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelFlow(request *CancelFlowRequest) (response *CancelFlowResponse, err error) {
    return c.CancelFlowWithContext(context.Background(), request)
}

// CancelFlow
// 用于撤销签署流程
//
// 适用场景：如果某个合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。
//
// 注：如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER_CANCELREASON = "MissingParameter.CancelReason"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelFlowWithContext(ctx context.Context, request *CancelFlowRequest) (response *CancelFlowResponse, err error) {
    if request == nil {
        request = NewCancelFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCancelMultiFlowSignQRCodeRequest() (request *CancelMultiFlowSignQRCodeRequest) {
    request = &CancelMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CancelMultiFlowSignQRCode")
    
    
    return
}

func NewCancelMultiFlowSignQRCodeResponse() (response *CancelMultiFlowSignQRCodeResponse) {
    response = &CancelMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelMultiFlowSignQRCode
// 此接口（CancelMultiFlowSignQRCode）用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_QRCODEID = "MissingParameter.QrCodeId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QRHASEXPIRE = "OperationDenied.QrHasExpire"
//  OPERATIONDENIED_QRINVALID = "OperationDenied.QrInvalid"
//  RESOURCENOTFOUND_QRINFO = "ResourceNotFound.QrInfo"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelMultiFlowSignQRCode(request *CancelMultiFlowSignQRCodeRequest) (response *CancelMultiFlowSignQRCodeResponse, err error) {
    return c.CancelMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// CancelMultiFlowSignQRCode
// 此接口（CancelMultiFlowSignQRCode）用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_QRCODEID = "MissingParameter.QrCodeId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QRHASEXPIRE = "OperationDenied.QrHasExpire"
//  OPERATIONDENIED_QRINVALID = "OperationDenied.QrInvalid"
//  RESOURCENOTFOUND_QRINFO = "ResourceNotFound.QrInfo"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelMultiFlowSignQRCodeWithContext(ctx context.Context, request *CancelMultiFlowSignQRCodeRequest) (response *CancelMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewCancelMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchCancelFlowUrlRequest() (request *CreateBatchCancelFlowUrlRequest) {
    request = &CreateBatchCancelFlowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateBatchCancelFlowUrl")
    
    
    return
}

func NewCreateBatchCancelFlowUrlResponse() (response *CreateBatchCancelFlowUrlResponse) {
    response = &CreateBatchCancelFlowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBatchCancelFlowUrl
// 电子签企业版：指定需要批量撤回的签署流程Id，获取批量撤销链接
//
// 客户指定需要撤回的签署流程Id，最多100个，超过100不处理；接口调用成功返回批量撤回合同的链接，通过链接跳转到电子签小程序完成批量撤回
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_BATCHCANCELFORBID = "OperationDenied.BatchCancelForbid"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchCancelFlowUrl(request *CreateBatchCancelFlowUrlRequest) (response *CreateBatchCancelFlowUrlResponse, err error) {
    return c.CreateBatchCancelFlowUrlWithContext(context.Background(), request)
}

// CreateBatchCancelFlowUrl
// 电子签企业版：指定需要批量撤回的签署流程Id，获取批量撤销链接
//
// 客户指定需要撤回的签署流程Id，最多100个，超过100不处理；接口调用成功返回批量撤回合同的链接，通过链接跳转到电子签小程序完成批量撤回
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_BATCHCANCELFORBID = "OperationDenied.BatchCancelForbid"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchCancelFlowUrlWithContext(ctx context.Context, request *CreateBatchCancelFlowUrlRequest) (response *CreateBatchCancelFlowUrlResponse, err error) {
    if request == nil {
        request = NewCreateBatchCancelFlowUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchCancelFlowUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchCancelFlowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConvertTaskApiRequest() (request *CreateConvertTaskApiRequest) {
    request = &CreateConvertTaskApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateConvertTaskApi")
    
    
    return
}

func NewCreateConvertTaskApiResponse() (response *CreateConvertTaskApiResponse) {
    response = &CreateConvertTaskApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConvertTaskApi
// 创建文件转换任务
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateConvertTaskApi(request *CreateConvertTaskApiRequest) (response *CreateConvertTaskApiResponse, err error) {
    return c.CreateConvertTaskApiWithContext(context.Background(), request)
}

// CreateConvertTaskApi
// 创建文件转换任务
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateConvertTaskApiWithContext(ctx context.Context, request *CreateConvertTaskApiRequest) (response *CreateConvertTaskApiResponse, err error) {
    if request == nil {
        request = NewCreateConvertTaskApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConvertTaskApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConvertTaskApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDocumentRequest() (request *CreateDocumentRequest) {
    request = &CreateDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateDocument")
    
    
    return
}

func NewCreateDocumentResponse() (response *CreateDocumentResponse) {
    response = &CreateDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDocument
// 创建签署流程电子文档
//
// 适用场景：见创建签署流程接口。x0b
//
// 注：该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。是“发起流程”接口的前置接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateDocument(request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    return c.CreateDocumentWithContext(context.Background(), request)
}

// CreateDocument
// 创建签署流程电子文档
//
// 适用场景：见创建签署流程接口。x0b
//
// 注：该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。是“发起流程”接口的前置接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateDocumentWithContext(ctx context.Context, request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    if request == nil {
        request = NewCreateDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowRequest() (request *CreateFlowRequest) {
    request = &CreateFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlow")
    
    
    return
}

func NewCreateFlowResponse() (response *CreateFlowResponse) {
    response = &CreateFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlow
// 创建签署流程
//
// 适用场景：在标准制式的合同场景中，可通过提前预制好模板文件，每次调用模板文件的id，补充合同内容信息及签署信息生成电子合同。
//
// 注：该接口是通过模板生成合同流程的前置接口，先创建一个不包含签署文件的流程。配合“创建电子文档”接口和“发起流程”接口使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlow(request *CreateFlowRequest) (response *CreateFlowResponse, err error) {
    return c.CreateFlowWithContext(context.Background(), request)
}

// CreateFlow
// 创建签署流程
//
// 适用场景：在标准制式的合同场景中，可通过提前预制好模板文件，每次调用模板文件的id，补充合同内容信息及签署信息生成电子合同。
//
// 注：该接口是通过模板生成合同流程的前置接口，先创建一个不包含签署文件的流程。配合“创建电子文档”接口和“发起流程”接口使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowWithContext(ctx context.Context, request *CreateFlowRequest) (response *CreateFlowResponse, err error) {
    if request == nil {
        request = NewCreateFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowApproversRequest() (request *CreateFlowApproversRequest) {
    request = &CreateFlowApproversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowApprovers")
    
    
    return
}

func NewCreateFlowApproversResponse() (response *CreateFlowApproversResponse) {
    response = &CreateFlowApproversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowApprovers
// 补充签署流程本企业签署人信息
//
// 适用场景：在通过模板或者文件发起合同时，若未指定本企业签署人信息，则流程发起后，可以调用此接口补充签署人。
//
// 同一签署人可以补充多个员工作为候选签署人,最终签署人取决于谁先领取合同完成签署。
//
// 
//
// 注：目前暂时只支持补充来源于企业微信的员工作为候选签署人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SENSITIVE = "InvalidParameter.Sensitive"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowApprovers(request *CreateFlowApproversRequest) (response *CreateFlowApproversResponse, err error) {
    return c.CreateFlowApproversWithContext(context.Background(), request)
}

// CreateFlowApprovers
// 补充签署流程本企业签署人信息
//
// 适用场景：在通过模板或者文件发起合同时，若未指定本企业签署人信息，则流程发起后，可以调用此接口补充签署人。
//
// 同一签署人可以补充多个员工作为候选签署人,最终签署人取决于谁先领取合同完成签署。
//
// 
//
// 注：目前暂时只支持补充来源于企业微信的员工作为候选签署人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SENSITIVE = "InvalidParameter.Sensitive"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowApproversWithContext(ctx context.Context, request *CreateFlowApproversRequest) (response *CreateFlowApproversResponse, err error) {
    if request == nil {
        request = NewCreateFlowApproversRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowApprovers require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowApproversResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowByFilesRequest() (request *CreateFlowByFilesRequest) {
    request = &CreateFlowByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowByFiles")
    
    
    return
}

func NewCreateFlowByFilesResponse() (response *CreateFlowByFilesResponse) {
    response = &CreateFlowByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowByFiles
// 此接口（CreateFlowByFiles）用来通过上传后的pdf资源编号来创建待签署的合同流程。
//
// 适用场景1：适用非制式的合同文件签署。一般开发者自己有完整的签署文件，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。
//
// 适用场景2：可通过该接口传入制式合同文件，同时在指定位置添加签署控件。可以起到接口创建临时模板的效果。如果是标准的制式文件，建议使用模板功能生成模板ID进行合同流程的生成。
//
// 注意事项：该接口需要依赖“多文件上传”接口生成pdf资源编号（FileIds）进行使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowByFiles(request *CreateFlowByFilesRequest) (response *CreateFlowByFilesResponse, err error) {
    return c.CreateFlowByFilesWithContext(context.Background(), request)
}

// CreateFlowByFiles
// 此接口（CreateFlowByFiles）用来通过上传后的pdf资源编号来创建待签署的合同流程。
//
// 适用场景1：适用非制式的合同文件签署。一般开发者自己有完整的签署文件，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。
//
// 适用场景2：可通过该接口传入制式合同文件，同时在指定位置添加签署控件。可以起到接口创建临时模板的效果。如果是标准的制式文件，建议使用模板功能生成模板ID进行合同流程的生成。
//
// 注意事项：该接口需要依赖“多文件上传”接口生成pdf资源编号（FileIds）进行使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowByFilesWithContext(ctx context.Context, request *CreateFlowByFilesRequest) (response *CreateFlowByFilesResponse, err error) {
    if request == nil {
        request = NewCreateFlowByFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowByFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowEvidenceReportRequest() (request *CreateFlowEvidenceReportRequest) {
    request = &CreateFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowEvidenceReport")
    
    
    return
}

func NewCreateFlowEvidenceReportResponse() (response *CreateFlowEvidenceReportResponse) {
    response = &CreateFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowEvidenceReport
// 创建出证报告，返回报告 ID。需要配合出证套餐才能调用。
//
// 出证需要一定时间，建议调用创建出证24小时之后再通过DescribeFlowEvidenceReport进行查询。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowEvidenceReport(request *CreateFlowEvidenceReportRequest) (response *CreateFlowEvidenceReportResponse, err error) {
    return c.CreateFlowEvidenceReportWithContext(context.Background(), request)
}

// CreateFlowEvidenceReport
// 创建出证报告，返回报告 ID。需要配合出证套餐才能调用。
//
// 出证需要一定时间，建议调用创建出证24小时之后再通过DescribeFlowEvidenceReport进行查询。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowEvidenceReportWithContext(ctx context.Context, request *CreateFlowEvidenceReportRequest) (response *CreateFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewCreateFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowRemindsRequest() (request *CreateFlowRemindsRequest) {
    request = &CreateFlowRemindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowReminds")
    
    
    return
}

func NewCreateFlowRemindsResponse() (response *CreateFlowRemindsResponse) {
    response = &CreateFlowRemindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowReminds
// 指定需要批量催办的签署流程Id，批量催办合同，最多100个; 接口失败后返回错误信息
//
// 注意:
//
// 该接口不可直接调用，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowReminds(request *CreateFlowRemindsRequest) (response *CreateFlowRemindsResponse, err error) {
    return c.CreateFlowRemindsWithContext(context.Background(), request)
}

// CreateFlowReminds
// 指定需要批量催办的签署流程Id，批量催办合同，最多100个; 接口失败后返回错误信息
//
// 注意:
//
// 该接口不可直接调用，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowRemindsWithContext(ctx context.Context, request *CreateFlowRemindsRequest) (response *CreateFlowRemindsResponse, err error) {
    if request == nil {
        request = NewCreateFlowRemindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowReminds require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowRemindsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowSignReviewRequest() (request *CreateFlowSignReviewRequest) {
    request = &CreateFlowSignReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowSignReview")
    
    
    return
}

func NewCreateFlowSignReviewResponse() (response *CreateFlowSignReviewResponse) {
    response = &CreateFlowSignReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowSignReview
// 提交企业签署流程审批结果
//
// 适用场景: 
//
// 在通过接口(CreateFlow 或者CreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true，且发起方企业作为签署方参与了流程签署，则可以调用此接口提交企业内部签署审批结果。
//
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignReview(request *CreateFlowSignReviewRequest) (response *CreateFlowSignReviewResponse, err error) {
    return c.CreateFlowSignReviewWithContext(context.Background(), request)
}

// CreateFlowSignReview
// 提交企业签署流程审批结果
//
// 适用场景: 
//
// 在通过接口(CreateFlow 或者CreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true，且发起方企业作为签署方参与了流程签署，则可以调用此接口提交企业内部签署审批结果。
//
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignReviewWithContext(ctx context.Context, request *CreateFlowSignReviewRequest) (response *CreateFlowSignReviewResponse, err error) {
    if request == nil {
        request = NewCreateFlowSignReviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowSignReview require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowSignReviewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowSignUrlRequest() (request *CreateFlowSignUrlRequest) {
    request = &CreateFlowSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowSignUrl")
    
    
    return
}

func NewCreateFlowSignUrlResponse() (response *CreateFlowSignUrlResponse) {
    response = &CreateFlowSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowSignUrl
// 创建签署链接，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignUrl(request *CreateFlowSignUrlRequest) (response *CreateFlowSignUrlResponse, err error) {
    return c.CreateFlowSignUrlWithContext(context.Background(), request)
}

// CreateFlowSignUrl
// 创建签署链接，请联系客户经理申请使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignUrlWithContext(ctx context.Context, request *CreateFlowSignUrlRequest) (response *CreateFlowSignUrlResponse, err error) {
    if request == nil {
        request = NewCreateFlowSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrationEmployeesRequest() (request *CreateIntegrationEmployeesRequest) {
    request = &CreateIntegrationEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateIntegrationEmployees")
    
    
    return
}

func NewCreateIntegrationEmployeesResponse() (response *CreateIntegrationEmployeesResponse) {
    response = &CreateIntegrationEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIntegrationEmployees
// 创建员工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationEmployees(request *CreateIntegrationEmployeesRequest) (response *CreateIntegrationEmployeesResponse, err error) {
    return c.CreateIntegrationEmployeesWithContext(context.Background(), request)
}

// CreateIntegrationEmployees
// 创建员工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationEmployeesWithContext(ctx context.Context, request *CreateIntegrationEmployeesRequest) (response *CreateIntegrationEmployeesResponse, err error) {
    if request == nil {
        request = NewCreateIntegrationEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrationEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrationEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultiFlowSignQRCodeRequest() (request *CreateMultiFlowSignQRCodeRequest) {
    request = &CreateMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateMultiFlowSignQRCode")
    
    
    return
}

func NewCreateMultiFlowSignQRCodeResponse() (response *CreateMultiFlowSignQRCodeResponse) {
    response = &CreateMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMultiFlowSignQRCode
// 此接口（CreateMultiFlowSignQRCode）用于创建一码多扫流程签署二维码。
//
// 适用场景：无需填写签署人信息，可通过模板id生成签署二维码，签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 适用的模板仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模板，且模板中发起方没有填写控件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_QREFFECTDAY = "InvalidParameter.QrEffectDay"
//  INVALIDPARAMETER_QRFLOWEFFECTDAY = "InvalidParameter.QrFlowEffectDay"
//  INVALIDPARAMETER_TEMPLATEID = "InvalidParameter.TemplateId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateMultiFlowSignQRCode(request *CreateMultiFlowSignQRCodeRequest) (response *CreateMultiFlowSignQRCodeResponse, err error) {
    return c.CreateMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// CreateMultiFlowSignQRCode
// 此接口（CreateMultiFlowSignQRCode）用于创建一码多扫流程签署二维码。
//
// 适用场景：无需填写签署人信息，可通过模板id生成签署二维码，签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 适用的模板仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模板，且模板中发起方没有填写控件。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_QREFFECTDAY = "InvalidParameter.QrEffectDay"
//  INVALIDPARAMETER_QRFLOWEFFECTDAY = "InvalidParameter.QrFlowEffectDay"
//  INVALIDPARAMETER_TEMPLATEID = "InvalidParameter.TemplateId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateMultiFlowSignQRCodeWithContext(ctx context.Context, request *CreateMultiFlowSignQRCodeRequest) (response *CreateMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewCreateMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrepareFlowRequest() (request *CreatePrepareFlowRequest) {
    request = &CreatePrepareFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreatePrepareFlow")
    
    
    return
}

func NewCreatePrepareFlowResponse() (response *CreatePrepareFlowResponse) {
    response = &CreatePrepareFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrepareFlow
// 创建快速发起流程
//
// 适用场景：用户通过API 合同文件及签署信息，并可通过我们返回的URL在页面完成签署控件等信息的编辑与确认，快速发起合同.
//
// 注：该接口文件的resourceId 是通过上传文件之后获取的。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreatePrepareFlow(request *CreatePrepareFlowRequest) (response *CreatePrepareFlowResponse, err error) {
    return c.CreatePrepareFlowWithContext(context.Background(), request)
}

// CreatePrepareFlow
// 创建快速发起流程
//
// 适用场景：用户通过API 合同文件及签署信息，并可通过我们返回的URL在页面完成签署控件等信息的编辑与确认，快速发起合同.
//
// 注：该接口文件的resourceId 是通过上传文件之后获取的。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreatePrepareFlowWithContext(ctx context.Context, request *CreatePrepareFlowRequest) (response *CreatePrepareFlowResponse, err error) {
    if request == nil {
        request = NewCreatePrepareFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrepareFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrepareFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseFlowRequest() (request *CreateReleaseFlowRequest) {
    request = &CreateReleaseFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateReleaseFlow")
    
    
    return
}

func NewCreateReleaseFlowResponse() (response *CreateReleaseFlowResponse) {
    response = &CreateReleaseFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReleaseFlow
// 发起解除协议，主要应用场景为：基于一份已经签署的合同(签署流程)，进行解除操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateReleaseFlow(request *CreateReleaseFlowRequest) (response *CreateReleaseFlowResponse, err error) {
    return c.CreateReleaseFlowWithContext(context.Background(), request)
}

// CreateReleaseFlow
// 发起解除协议，主要应用场景为：基于一份已经签署的合同(签署流程)，进行解除操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateReleaseFlowWithContext(ctx context.Context, request *CreateReleaseFlowRequest) (response *CreateReleaseFlowResponse, err error) {
    if request == nil {
        request = NewCreateReleaseFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReleaseFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchemeUrlRequest() (request *CreateSchemeUrlRequest) {
    request = &CreateSchemeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateSchemeUrl")
    
    
    return
}

func NewCreateSchemeUrlResponse() (response *CreateSchemeUrlResponse) {
    response = &CreateSchemeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSchemeUrl
// 获取小程序跳转链接
//
// 
//
// 适用场景：如果需要签署人在自己的APP、小程序、H5应用中签署，可以通过此接口获取跳转腾讯电子签小程序的签署跳转链接。
//
// 
//
// 注：如果签署人是在PC端扫码签署，可以通过生成跳转链接自主转换成二维码，让签署人在PC端扫码签署。
//
// 
//
// 
//
// 跳转到小程序的实现，参考官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式）
//
// 
//
// 
//
// 如您需要自主配置小程序跳转链接，请参考: <a href="https://cloud.tencent.com/document/product/1323/74774">跳转小程序链接配置说明</a>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSchemeUrl(request *CreateSchemeUrlRequest) (response *CreateSchemeUrlResponse, err error) {
    return c.CreateSchemeUrlWithContext(context.Background(), request)
}

// CreateSchemeUrl
// 获取小程序跳转链接
//
// 
//
// 适用场景：如果需要签署人在自己的APP、小程序、H5应用中签署，可以通过此接口获取跳转腾讯电子签小程序的签署跳转链接。
//
// 
//
// 注：如果签署人是在PC端扫码签署，可以通过生成跳转链接自主转换成二维码，让签署人在PC端扫码签署。
//
// 
//
// 
//
// 跳转到小程序的实现，参考官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式）
//
// 
//
// 
//
// 如您需要自主配置小程序跳转链接，请参考: <a href="https://cloud.tencent.com/document/product/1323/74774">跳转小程序链接配置说明</a>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSchemeUrlWithContext(ctx context.Context, request *CreateSchemeUrlRequest) (response *CreateSchemeUrlResponse, err error) {
    if request == nil {
        request = NewCreateSchemeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchemeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSchemeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSealPolicyRequest() (request *CreateSealPolicyRequest) {
    request = &CreateSealPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateSealPolicy")
    
    
    return
}

func NewCreateSealPolicyResponse() (response *CreateSealPolicyResponse) {
    response = &CreateSealPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSealPolicy
// 对企业员工进行印章授权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSealPolicy(request *CreateSealPolicyRequest) (response *CreateSealPolicyResponse, err error) {
    return c.CreateSealPolicyWithContext(context.Background(), request)
}

// CreateSealPolicy
// 对企业员工进行印章授权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSealPolicyWithContext(ctx context.Context, request *CreateSealPolicyRequest) (response *CreateSealPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSealPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSealPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSealPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserAutoSignEnableUrlRequest() (request *CreateUserAutoSignEnableUrlRequest) {
    request = &CreateUserAutoSignEnableUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateUserAutoSignEnableUrl")
    
    
    return
}

func NewCreateUserAutoSignEnableUrlResponse() (response *CreateUserAutoSignEnableUrlResponse) {
    response = &CreateUserAutoSignEnableUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserAutoSignEnableUrl
// 企业方可以通过此接口获取个人用户开启自动签的跳转链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
func (c *Client) CreateUserAutoSignEnableUrl(request *CreateUserAutoSignEnableUrlRequest) (response *CreateUserAutoSignEnableUrlResponse, err error) {
    return c.CreateUserAutoSignEnableUrlWithContext(context.Background(), request)
}

// CreateUserAutoSignEnableUrl
// 企业方可以通过此接口获取个人用户开启自动签的跳转链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
func (c *Client) CreateUserAutoSignEnableUrlWithContext(ctx context.Context, request *CreateUserAutoSignEnableUrlRequest) (response *CreateUserAutoSignEnableUrlResponse, err error) {
    if request == nil {
        request = NewCreateUserAutoSignEnableUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserAutoSignEnableUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserAutoSignEnableUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntegrationEmployeesRequest() (request *DeleteIntegrationEmployeesRequest) {
    request = &DeleteIntegrationEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DeleteIntegrationEmployees")
    
    
    return
}

func NewDeleteIntegrationEmployeesResponse() (response *DeleteIntegrationEmployeesResponse) {
    response = &DeleteIntegrationEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIntegrationEmployees
// 移除员工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteIntegrationEmployees(request *DeleteIntegrationEmployeesRequest) (response *DeleteIntegrationEmployeesResponse, err error) {
    return c.DeleteIntegrationEmployeesWithContext(context.Background(), request)
}

// DeleteIntegrationEmployees
// 移除员工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteIntegrationEmployeesWithContext(ctx context.Context, request *DeleteIntegrationEmployeesRequest) (response *DeleteIntegrationEmployeesResponse, err error) {
    if request == nil {
        request = NewDeleteIntegrationEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntegrationEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntegrationEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSealPoliciesRequest() (request *DeleteSealPoliciesRequest) {
    request = &DeleteSealPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DeleteSealPolicies")
    
    
    return
}

func NewDeleteSealPoliciesResponse() (response *DeleteSealPoliciesResponse) {
    response = &DeleteSealPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSealPolicies
// 撤销员工持有的印章权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteSealPolicies(request *DeleteSealPoliciesRequest) (response *DeleteSealPoliciesResponse, err error) {
    return c.DeleteSealPoliciesWithContext(context.Background(), request)
}

// DeleteSealPolicies
// 撤销员工持有的印章权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteSealPoliciesWithContext(ctx context.Context, request *DeleteSealPoliciesRequest) (response *DeleteSealPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSealPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSealPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSealPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileUrlsRequest() (request *DescribeFileUrlsRequest) {
    request = &DescribeFileUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFileUrls")
    
    
    return
}

func NewDescribeFileUrlsResponse() (response *DescribeFileUrlsResponse) {
    response = &DescribeFileUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileUrls
// 查询文件下载URL
//
// 适用场景：通过传参合同流程编号，下载对应的合同PDF文件流到本地。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSID = "InvalidParameter.BusinessId"
//  INVALIDPARAMETER_BUSINESSTYPE = "InvalidParameter.BusinessType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFileUrls(request *DescribeFileUrlsRequest) (response *DescribeFileUrlsResponse, err error) {
    return c.DescribeFileUrlsWithContext(context.Background(), request)
}

// DescribeFileUrls
// 查询文件下载URL
//
// 适用场景：通过传参合同流程编号，下载对应的合同PDF文件流到本地。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSID = "InvalidParameter.BusinessId"
//  INVALIDPARAMETER_BUSINESSTYPE = "InvalidParameter.BusinessType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFileUrlsWithContext(ctx context.Context, request *DescribeFileUrlsRequest) (response *DescribeFileUrlsResponse, err error) {
    if request == nil {
        request = NewDescribeFileUrlsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileUrls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowBriefsRequest() (request *DescribeFlowBriefsRequest) {
    request = &DescribeFlowBriefsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowBriefs")
    
    
    return
}

func NewDescribeFlowBriefsResponse() (response *DescribeFlowBriefsResponse) {
    response = &DescribeFlowBriefsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowBriefs
// 查询流程摘要
//
// 适用场景：可用于主动查询某个合同流程的签署状态信息。可以配合回调通知使用。
//
// 日调用量默认10W
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPENID = "InvalidParameter.InvalidOpenId"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDROLEID = "InvalidParameter.InvalidRoleId"
//  INVALIDPARAMETER_INVALIDROLENAME = "InvalidParameter.InvalidRoleName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_INVALIDVERIFYCODE = "InvalidParameter.InvalidVerifyCode"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowBriefs(request *DescribeFlowBriefsRequest) (response *DescribeFlowBriefsResponse, err error) {
    return c.DescribeFlowBriefsWithContext(context.Background(), request)
}

// DescribeFlowBriefs
// 查询流程摘要
//
// 适用场景：可用于主动查询某个合同流程的签署状态信息。可以配合回调通知使用。
//
// 日调用量默认10W
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPENID = "InvalidParameter.InvalidOpenId"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDROLEID = "InvalidParameter.InvalidRoleId"
//  INVALIDPARAMETER_INVALIDROLENAME = "InvalidParameter.InvalidRoleName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_INVALIDVERIFYCODE = "InvalidParameter.InvalidVerifyCode"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowBriefsWithContext(ctx context.Context, request *DescribeFlowBriefsRequest) (response *DescribeFlowBriefsResponse, err error) {
    if request == nil {
        request = NewDescribeFlowBriefsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowBriefs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowBriefsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowEvidenceReportRequest() (request *DescribeFlowEvidenceReportRequest) {
    request = &DescribeFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowEvidenceReport")
    
    
    return
}

func NewDescribeFlowEvidenceReportResponse() (response *DescribeFlowEvidenceReportResponse) {
    response = &DescribeFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowEvidenceReport
// 查询出证报告，返回报告 URL。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowEvidenceReport(request *DescribeFlowEvidenceReportRequest) (response *DescribeFlowEvidenceReportResponse, err error) {
    return c.DescribeFlowEvidenceReportWithContext(context.Background(), request)
}

// DescribeFlowEvidenceReport
// 查询出证报告，返回报告 URL。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowEvidenceReportWithContext(ctx context.Context, request *DescribeFlowEvidenceReportRequest) (response *DescribeFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewDescribeFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowInfoRequest() (request *DescribeFlowInfoRequest) {
    request = &DescribeFlowInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowInfo")
    
    
    return
}

func NewDescribeFlowInfoResponse() (response *DescribeFlowInfoResponse) {
    response = &DescribeFlowInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowInfo
// 查询合同详情
//
// 适用场景：可用于主动查询某个合同详情信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowInfo(request *DescribeFlowInfoRequest) (response *DescribeFlowInfoResponse, err error) {
    return c.DescribeFlowInfoWithContext(context.Background(), request)
}

// DescribeFlowInfo
// 查询合同详情
//
// 适用场景：可用于主动查询某个合同详情信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowInfoWithContext(ctx context.Context, request *DescribeFlowInfoRequest) (response *DescribeFlowInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFlowInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowTemplatesRequest() (request *DescribeFlowTemplatesRequest) {
    request = &DescribeFlowTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowTemplates")
    
    
    return
}

func NewDescribeFlowTemplatesResponse() (response *DescribeFlowTemplatesResponse) {
    response = &DescribeFlowTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowTemplates
// 当模板较多或模板中的控件较多时，可以通过查询模板接口更方便的获取模板列表，以及每个模板内的控件信息。该接口常用来配合“创建电子文档”接口作为前置的接口使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFlowTemplates(request *DescribeFlowTemplatesRequest) (response *DescribeFlowTemplatesResponse, err error) {
    return c.DescribeFlowTemplatesWithContext(context.Background(), request)
}

// DescribeFlowTemplates
// 当模板较多或模板中的控件较多时，可以通过查询模板接口更方便的获取模板列表，以及每个模板内的控件信息。该接口常用来配合“创建电子文档”接口作为前置的接口使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFlowTemplatesWithContext(ctx context.Context, request *DescribeFlowTemplatesRequest) (response *DescribeFlowTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeFlowTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationEmployeesRequest() (request *DescribeIntegrationEmployeesRequest) {
    request = &DescribeIntegrationEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeIntegrationEmployees")
    
    
    return
}

func NewDescribeIntegrationEmployeesResponse() (response *DescribeIntegrationEmployeesResponse) {
    response = &DescribeIntegrationEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntegrationEmployees
// 查询员工信息，每次返回的数据量最大为20
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeIntegrationEmployees(request *DescribeIntegrationEmployeesRequest) (response *DescribeIntegrationEmployeesResponse, err error) {
    return c.DescribeIntegrationEmployeesWithContext(context.Background(), request)
}

// DescribeIntegrationEmployees
// 查询员工信息，每次返回的数据量最大为20
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeIntegrationEmployeesWithContext(ctx context.Context, request *DescribeIntegrationEmployeesRequest) (response *DescribeIntegrationEmployeesResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationMainOrganizationUserRequest() (request *DescribeIntegrationMainOrganizationUserRequest) {
    request = &DescribeIntegrationMainOrganizationUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeIntegrationMainOrganizationUser")
    
    
    return
}

func NewDescribeIntegrationMainOrganizationUserResponse() (response *DescribeIntegrationMainOrganizationUserResponse) {
    response = &DescribeIntegrationMainOrganizationUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIntegrationMainOrganizationUser
// 通过子企业影子账号查询主企业员工账号
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIntegrationMainOrganizationUser(request *DescribeIntegrationMainOrganizationUserRequest) (response *DescribeIntegrationMainOrganizationUserResponse, err error) {
    return c.DescribeIntegrationMainOrganizationUserWithContext(context.Background(), request)
}

// DescribeIntegrationMainOrganizationUser
// 通过子企业影子账号查询主企业员工账号
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeIntegrationMainOrganizationUserWithContext(ctx context.Context, request *DescribeIntegrationMainOrganizationUserRequest) (response *DescribeIntegrationMainOrganizationUserResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationMainOrganizationUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationMainOrganizationUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationMainOrganizationUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationGroupOrganizationsRequest() (request *DescribeOrganizationGroupOrganizationsRequest) {
    request = &DescribeOrganizationGroupOrganizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeOrganizationGroupOrganizations")
    
    
    return
}

func NewDescribeOrganizationGroupOrganizationsResponse() (response *DescribeOrganizationGroupOrganizationsResponse) {
    response = &DescribeOrganizationGroupOrganizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrganizationGroupOrganizations
// 此API接口用户查询加入集团的成员企业
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeOrganizationGroupOrganizations(request *DescribeOrganizationGroupOrganizationsRequest) (response *DescribeOrganizationGroupOrganizationsResponse, err error) {
    return c.DescribeOrganizationGroupOrganizationsWithContext(context.Background(), request)
}

// DescribeOrganizationGroupOrganizations
// 此API接口用户查询加入集团的成员企业
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeOrganizationGroupOrganizationsWithContext(ctx context.Context, request *DescribeOrganizationGroupOrganizationsRequest) (response *DescribeOrganizationGroupOrganizationsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationGroupOrganizationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationGroupOrganizations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationGroupOrganizationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationSealsRequest() (request *DescribeOrganizationSealsRequest) {
    request = &DescribeOrganizationSealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeOrganizationSeals")
    
    
    return
}

func NewDescribeOrganizationSealsResponse() (response *DescribeOrganizationSealsResponse) {
    response = &DescribeOrganizationSealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrganizationSeals
// 查询企业印章的列表，需要操作者具有查询印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeOrganizationSeals(request *DescribeOrganizationSealsRequest) (response *DescribeOrganizationSealsResponse, err error) {
    return c.DescribeOrganizationSealsWithContext(context.Background(), request)
}

// DescribeOrganizationSeals
// 查询企业印章的列表，需要操作者具有查询印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeOrganizationSealsWithContext(ctx context.Context, request *DescribeOrganizationSealsRequest) (response *DescribeOrganizationSealsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationSealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationSeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationSealsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeThirdPartyAuthCodeRequest() (request *DescribeThirdPartyAuthCodeRequest) {
    request = &DescribeThirdPartyAuthCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeThirdPartyAuthCode")
    
    
    return
}

func NewDescribeThirdPartyAuthCodeResponse() (response *DescribeThirdPartyAuthCodeResponse) {
    response = &DescribeThirdPartyAuthCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeThirdPartyAuthCode
// 通过AuthCode查询用户是否实名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_AUTHCODE = "MissingParameter.AuthCode"
//  OPERATIONDENIED_AUTHCODEINVALID = "OperationDenied.AuthCodeInvalid"
func (c *Client) DescribeThirdPartyAuthCode(request *DescribeThirdPartyAuthCodeRequest) (response *DescribeThirdPartyAuthCodeResponse, err error) {
    return c.DescribeThirdPartyAuthCodeWithContext(context.Background(), request)
}

// DescribeThirdPartyAuthCode
// 通过AuthCode查询用户是否实名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_AUTHCODE = "MissingParameter.AuthCode"
//  OPERATIONDENIED_AUTHCODEINVALID = "OperationDenied.AuthCodeInvalid"
func (c *Client) DescribeThirdPartyAuthCodeWithContext(ctx context.Context, request *DescribeThirdPartyAuthCodeRequest) (response *DescribeThirdPartyAuthCodeResponse, err error) {
    if request == nil {
        request = NewDescribeThirdPartyAuthCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeThirdPartyAuthCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeThirdPartyAuthCodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserAutoSignStatusRequest() (request *DescribeUserAutoSignStatusRequest) {
    request = &DescribeUserAutoSignStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeUserAutoSignStatus")
    
    
    return
}

func NewDescribeUserAutoSignStatusResponse() (response *DescribeUserAutoSignStatusResponse) {
    response = &DescribeUserAutoSignStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserAutoSignStatus
// 企业方可以通过此接口查询个人用户自动签开启状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeUserAutoSignStatus(request *DescribeUserAutoSignStatusRequest) (response *DescribeUserAutoSignStatusResponse, err error) {
    return c.DescribeUserAutoSignStatusWithContext(context.Background(), request)
}

// DescribeUserAutoSignStatus
// 企业方可以通过此接口查询个人用户自动签开启状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeUserAutoSignStatusWithContext(ctx context.Context, request *DescribeUserAutoSignStatusRequest) (response *DescribeUserAutoSignStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUserAutoSignStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserAutoSignStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserAutoSignStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisableUserAutoSignRequest() (request *DisableUserAutoSignRequest) {
    request = &DisableUserAutoSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DisableUserAutoSign")
    
    
    return
}

func NewDisableUserAutoSignResponse() (response *DisableUserAutoSignResponse) {
    response = &DisableUserAutoSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableUserAutoSign
// 企业方可以通过此接口关闭个人的自动签功能
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DisableUserAutoSign(request *DisableUserAutoSignRequest) (response *DisableUserAutoSignResponse, err error) {
    return c.DisableUserAutoSignWithContext(context.Background(), request)
}

// DisableUserAutoSign
// 企业方可以通过此接口关闭个人的自动签功能
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DisableUserAutoSignWithContext(ctx context.Context, request *DisableUserAutoSignRequest) (response *DisableUserAutoSignResponse, err error) {
    if request == nil {
        request = NewDisableUserAutoSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableUserAutoSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableUserAutoSignResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskResultApiRequest() (request *GetTaskResultApiRequest) {
    request = &GetTaskResultApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "GetTaskResultApi")
    
    
    return
}

func NewGetTaskResultApiResponse() (response *GetTaskResultApiResponse) {
    response = &GetTaskResultApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTaskResultApi
// 查询转换任务状态
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) GetTaskResultApi(request *GetTaskResultApiRequest) (response *GetTaskResultApiResponse, err error) {
    return c.GetTaskResultApiWithContext(context.Background(), request)
}

// GetTaskResultApi
// 查询转换任务状态
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) GetTaskResultApiWithContext(ctx context.Context, request *GetTaskResultApiRequest) (response *GetTaskResultApiResponse, err error) {
    if request == nil {
        request = NewGetTaskResultApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskResultApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskResultApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationCallbackInfoRequest() (request *ModifyApplicationCallbackInfoRequest) {
    request = &ModifyApplicationCallbackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "ModifyApplicationCallbackInfo")
    
    
    return
}

func NewModifyApplicationCallbackInfoResponse() (response *ModifyApplicationCallbackInfoResponse) {
    response = &ModifyApplicationCallbackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationCallbackInfo
// 新增/删除应用callbackinfo
//
// callbackinfo包含： 回调地址和签名key
//
// 操作：新增/删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyApplicationCallbackInfo(request *ModifyApplicationCallbackInfoRequest) (response *ModifyApplicationCallbackInfoResponse, err error) {
    return c.ModifyApplicationCallbackInfoWithContext(context.Background(), request)
}

// ModifyApplicationCallbackInfo
// 新增/删除应用callbackinfo
//
// callbackinfo包含： 回调地址和签名key
//
// 操作：新增/删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyApplicationCallbackInfoWithContext(ctx context.Context, request *ModifyApplicationCallbackInfoRequest) (response *ModifyApplicationCallbackInfoResponse, err error) {
    if request == nil {
        request = NewModifyApplicationCallbackInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationCallbackInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationCallbackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewStartFlowRequest() (request *StartFlowRequest) {
    request = &StartFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "StartFlow")
    
    
    return
}

func NewStartFlowResponse() (response *StartFlowResponse) {
    response = &StartFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartFlow
// 此接口用于发起流程
//
// 适用场景：见创建签署流程接口。
//
// 注：该接口是“创建电子文档”接口的后置接口，用于激活包含完整合同信息（模板及内容信息）的流程。激活后的流程就是一份待签署的电子合同。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOCUMENTNOAVAILABLE = "OperationDenied.DocumentNoAvailable"
//  OPERATIONDENIED_FLOWCANCELFORBID = "OperationDenied.FlowCancelForbid"
//  OPERATIONDENIED_FLOWHASSTARTED = "OperationDenied.FlowHasStarted"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FLOWNONEEDREVIEW = "OperationDenied.FlowNoNeedReview"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartFlow(request *StartFlowRequest) (response *StartFlowResponse, err error) {
    return c.StartFlowWithContext(context.Background(), request)
}

// StartFlow
// 此接口用于发起流程
//
// 适用场景：见创建签署流程接口。
//
// 注：该接口是“创建电子文档”接口的后置接口，用于激活包含完整合同信息（模板及内容信息）的流程。激活后的流程就是一份待签署的电子合同。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOCUMENTNOAVAILABLE = "OperationDenied.DocumentNoAvailable"
//  OPERATIONDENIED_FLOWCANCELFORBID = "OperationDenied.FlowCancelForbid"
//  OPERATIONDENIED_FLOWHASSTARTED = "OperationDenied.FlowHasStarted"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FLOWNONEEDREVIEW = "OperationDenied.FlowNoNeedReview"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartFlowWithContext(ctx context.Context, request *StartFlowRequest) (response *StartFlowResponse, err error) {
    if request == nil {
        request = NewStartFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartFlowResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFilesRequest() (request *UploadFilesRequest) {
    request = &UploadFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "UploadFiles")
    
    
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
// 适用场景：用于生成pdf资源编号（FileIds）来配合“用PDF创建流程”接口使用，使用场景可详见“用PDF创建流程”接口说明。
//
// 其中上传的文件，图片类型(png/jpg/jpeg)大小限制为5M，其他大小限制为60M。
//
// 调用时需要设置Domain/接口请求域名为 file.ess.tencent.cn，并设置参数Version/版本号为2020-12-22
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    return c.UploadFilesWithContext(context.Background(), request)
}

// UploadFiles
// 此接口（UploadFiles）用于文件上传。
//
// 适用场景：用于生成pdf资源编号（FileIds）来配合“用PDF创建流程”接口使用，使用场景可详见“用PDF创建流程”接口说明。
//
// 其中上传的文件，图片类型(png/jpg/jpeg)大小限制为5M，其他大小限制为60M。
//
// 调用时需要设置Domain/接口请求域名为 file.ess.tencent.cn，并设置参数Version/版本号为2020-12-22
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
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

func NewVerifyPdfRequest() (request *VerifyPdfRequest) {
    request = &VerifyPdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "VerifyPdf")
    
    
    return
}

func NewVerifyPdfResponse() (response *VerifyPdfResponse) {
    response = &VerifyPdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyPdf
// 验证合同文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) VerifyPdf(request *VerifyPdfRequest) (response *VerifyPdfResponse, err error) {
    return c.VerifyPdfWithContext(context.Background(), request)
}

// VerifyPdf
// 验证合同文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) VerifyPdfWithContext(ctx context.Context, request *VerifyPdfRequest) (response *VerifyPdfResponse, err error) {
    if request == nil {
        request = NewVerifyPdfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyPdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyPdfResponse()
    err = c.Send(request, response)
    return
}
