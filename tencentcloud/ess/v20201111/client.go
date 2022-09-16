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
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
func (c *Client) CreateConvertTaskApi(request *CreateConvertTaskApiRequest) (response *CreateConvertTaskApiResponse, err error) {
    return c.CreateConvertTaskApiWithContext(context.Background(), request)
}

// CreateConvertTaskApi
// 创建文件转换任务
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
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
// 适用场景：见创建签署流程接口。
//
// 
//
// 注：该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。是“发起流程”接口的前置接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
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
func (c *Client) CreateDocument(request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    return c.CreateDocumentWithContext(context.Background(), request)
}

// CreateDocument
// 创建签署流程电子文档
//
// 适用场景：见创建签署流程接口。
//
// 
//
// 注：该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。是“发起流程”接口的前置接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
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
// 适用场景：在通过模版或者文件发起合同时，若未指定本企业签署人信息，则流程发起后，可以调用此接口补充签署人。
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
func (c *Client) CreateFlowApprovers(request *CreateFlowApproversRequest) (response *CreateFlowApproversResponse, err error) {
    return c.CreateFlowApproversWithContext(context.Background(), request)
}

// CreateFlowApprovers
// 补充签署流程本企业签署人信息
//
// 适用场景：在通过模版或者文件发起合同时，若未指定本企业签署人信息，则流程发起后，可以调用此接口补充签署人。
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_PDF = "InternalError.Pdf"
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
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
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
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_PDF = "InternalError.Pdf"
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
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
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
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
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
// 【描述】：创建出证报告，返回报告 URL
//
// 【注意】：此接口需要通过添加白名单获取调用权限，请联系运营人员加白
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
func (c *Client) CreateFlowEvidenceReport(request *CreateFlowEvidenceReportRequest) (response *CreateFlowEvidenceReportResponse, err error) {
    return c.CreateFlowEvidenceReportWithContext(context.Background(), request)
}

// CreateFlowEvidenceReport
// 【描述】：创建出证报告，返回报告 URL
//
// 【注意】：此接口需要通过添加白名单获取调用权限，请联系运营人员加白
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
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
// 在通过接口(CreateFlow 或者CreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
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
func (c *Client) CreateFlowSignReview(request *CreateFlowSignReviewRequest) (response *CreateFlowSignReviewResponse, err error) {
    return c.CreateFlowSignReviewWithContext(context.Background(), request)
}

// CreateFlowSignReview
// 提交企业签署流程审批结果
//
// 适用场景: 
//
// 在通过接口(CreateFlow 或者CreateFlowByFiles)创建签署流程时，若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
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
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
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
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
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
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  MISSINGPARAMETER = "MissingParameter"
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
//  MISSINGPARAMETER = "MissingParameter"
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
// 二期接口-查询模板
//
// 适用场景：当模板较多或模板中的控件较多时，可以通过查询模板接口更方便的获取自己主体下的模板列表，以及每个模板内的控件信息。该接口常用来配合“创建电子文档”接口作为前置的接口使用。
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
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFlowTemplates(request *DescribeFlowTemplatesRequest) (response *DescribeFlowTemplatesResponse, err error) {
    return c.DescribeFlowTemplatesWithContext(context.Background(), request)
}

// DescribeFlowTemplates
// 二期接口-查询模板
//
// 适用场景：当模板较多或模板中的控件较多时，可以通过查询模板接口更方便的获取自己主体下的模板列表，以及每个模板内的控件信息。该接口常用来配合“创建电子文档”接口作为前置的接口使用。
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
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
// 调用时需要设置Domain 为 file.ess.tencent.cn，设置Version为2020-12-22
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    return c.UploadFilesWithContext(context.Background(), request)
}

// UploadFiles
// 此接口（UploadFiles）用于文件上传。
//
// 适用场景：用于生成pdf资源编号（FileIds）来配合“用PDF创建流程”接口使用，使用场景可详见“用PDF创建流程”接口说明。
//
// 调用时需要设置Domain 为 file.ess.tencent.cn，设置Version为2020-12-22
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
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
