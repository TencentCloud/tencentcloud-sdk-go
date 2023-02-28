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

package v20180808

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-08"

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


func NewBatchModifyDomainInfoRequest() (request *BatchModifyDomainInfoRequest) {
    request = &BatchModifyDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "BatchModifyDomainInfo")
    
    
    return
}

func NewBatchModifyDomainInfoResponse() (response *BatchModifyDomainInfoResponse) {
    response = &BatchModifyDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyDomainInfo
// 本接口 ( BatchModifyDomainInfo ) 用于批量域名信息修改 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) BatchModifyDomainInfo(request *BatchModifyDomainInfoRequest) (response *BatchModifyDomainInfoResponse, err error) {
    return c.BatchModifyDomainInfoWithContext(context.Background(), request)
}

// BatchModifyDomainInfo
// 本接口 ( BatchModifyDomainInfo ) 用于批量域名信息修改 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) BatchModifyDomainInfoWithContext(ctx context.Context, request *BatchModifyDomainInfoRequest) (response *BatchModifyDomainInfoResponse, err error) {
    if request == nil {
        request = NewBatchModifyDomainInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyDomainInfoResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBatchStatusRequest() (request *CheckBatchStatusRequest) {
    request = &CheckBatchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CheckBatchStatus")
    
    
    return
}

func NewCheckBatchStatusResponse() (response *CheckBatchStatusResponse) {
    response = &CheckBatchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBatchStatus
// 本接口 ( CheckBatchStatus ) 用于查询批量操作日志状态 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CheckBatchStatus(request *CheckBatchStatusRequest) (response *CheckBatchStatusResponse, err error) {
    return c.CheckBatchStatusWithContext(context.Background(), request)
}

// CheckBatchStatus
// 本接口 ( CheckBatchStatus ) 用于查询批量操作日志状态 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CheckBatchStatusWithContext(ctx context.Context, request *CheckBatchStatusRequest) (response *CheckBatchStatusResponse, err error) {
    if request == nil {
        request = NewCheckBatchStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckBatchStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckBatchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDomainRequest() (request *CheckDomainRequest) {
    request = &CheckDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CheckDomain")
    
    
    return
}

func NewCheckDomainResponse() (response *CheckDomainResponse) {
    response = &CheckDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckDomain
// 检查域名是否可以注册。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  FAILEDOPERATION_GETDOMAINPRICEFAILED = "FailedOperation.GetDomainPriceFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
func (c *Client) CheckDomain(request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
    return c.CheckDomainWithContext(context.Background(), request)
}

// CheckDomain
// 检查域名是否可以注册。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  FAILEDOPERATION_GETDOMAINPRICEFAILED = "FailedOperation.GetDomainPriceFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
func (c *Client) CheckDomainWithContext(ctx context.Context, request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
    if request == nil {
        request = NewCheckDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainBatchRequest() (request *CreateDomainBatchRequest) {
    request = &CreateDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreateDomainBatch")
    
    
    return
}

func NewCreateDomainBatchResponse() (response *CreateDomainBatchResponse) {
    response = &CreateDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDomainBatch
// 本接口 ( CreateDomainBatch ) 用于批量域名注册 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGISTERDOMAIN = "FailedOperation.RegisterDomain"
//  FAILEDOPERATION_REGISTERDOMAINFAILED = "FailedOperation.RegisterDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_PACKAGERESOURCEIDINVALID = "InvalidParameter.PackageResourceIdInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) CreateDomainBatch(request *CreateDomainBatchRequest) (response *CreateDomainBatchResponse, err error) {
    return c.CreateDomainBatchWithContext(context.Background(), request)
}

// CreateDomainBatch
// 本接口 ( CreateDomainBatch ) 用于批量域名注册 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REGISTERDOMAIN = "FailedOperation.RegisterDomain"
//  FAILEDOPERATION_REGISTERDOMAINFAILED = "FailedOperation.RegisterDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_PACKAGERESOURCEIDINVALID = "InvalidParameter.PackageResourceIdInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) CreateDomainBatchWithContext(ctx context.Context, request *CreateDomainBatchRequest) (response *CreateDomainBatchResponse, err error) {
    if request == nil {
        request = NewCreateDomainBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePhoneEmailRequest() (request *CreatePhoneEmailRequest) {
    request = &CreatePhoneEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreatePhoneEmail")
    
    
    return
}

func NewCreatePhoneEmailResponse() (response *CreatePhoneEmailResponse) {
    response = &CreatePhoneEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePhoneEmail
// 此接口用于创建有效的手机、邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEPHONEEMAIL = "FailedOperation.DuplicatePhoneEmail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CODETYPEISINVALID = "InvalidParameter.CodeTypeIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  INVALIDPARAMETER_VERIFYCODEISINVALID = "InvalidParameter.VerifyCodeIsInvalid"
func (c *Client) CreatePhoneEmail(request *CreatePhoneEmailRequest) (response *CreatePhoneEmailResponse, err error) {
    return c.CreatePhoneEmailWithContext(context.Background(), request)
}

// CreatePhoneEmail
// 此接口用于创建有效的手机、邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEPHONEEMAIL = "FailedOperation.DuplicatePhoneEmail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CODETYPEISINVALID = "InvalidParameter.CodeTypeIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  INVALIDPARAMETER_VERIFYCODEISINVALID = "InvalidParameter.VerifyCodeIsInvalid"
func (c *Client) CreatePhoneEmailWithContext(ctx context.Context, request *CreatePhoneEmailRequest) (response *CreatePhoneEmailResponse, err error) {
    if request == nil {
        request = NewCreatePhoneEmailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePhoneEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePhoneEmailResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
    request = &CreateTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreateTemplate")
    
    
    return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
    response = &CreateTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTemplate
// 本接口 ( CreateTemplate ) 用于添加域名信息模板 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETEMPLATEFAILED = "FailedOperation.CreateTemplateFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINLISTFAILED = "FailedOperation.DescribeDomainListFailed"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  FAILEDOPERATION_PROHIBITPHONEEMAIL = "FailedOperation.ProhibitPhoneEmail"
//  FAILEDOPERATION_TEMPLATEMAXNUMFAILED = "FailedOperation.TemplateMaxNumFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATECODEISINVALID = "InvalidParameter.CertificateCodeIsInvalid"
//  INVALIDPARAMETER_CERTIFICATEIMAGEISINVALID = "InvalidParameter.CertificateImageIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_EMAILISUNVERIFIED = "InvalidParameter.EmailIsUnverified"
//  INVALIDPARAMETER_NAMEISINVALID = "InvalidParameter.NameIsInvalid"
//  INVALIDPARAMETER_NAMEISKEYWORD = "InvalidParameter.NameIsKeyword"
//  INVALIDPARAMETER_ORGISINVALID = "InvalidParameter.OrgIsInvalid"
//  INVALIDPARAMETER_ORGISKEYWORD = "InvalidParameter.OrgIsKeyword"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
//  INVALIDPARAMETER_STREETISINVALID = "InvalidParameter.StreetIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISUNVERIFIED = "InvalidParameter.TelephoneIsUnverified"
//  INVALIDPARAMETER_USERTYPEISINVALID = "InvalidParameter.UserTypeIsInvalid"
//  INVALIDPARAMETER_ZIPCODEISINVALID = "InvalidParameter.ZipCodeIsInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  MISSINGPARAMETER_TEMPLATEIDISEXIST = "MissingParameter.TemplateIdIsExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
    return c.CreateTemplateWithContext(context.Background(), request)
}

// CreateTemplate
// 本接口 ( CreateTemplate ) 用于添加域名信息模板 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETEMPLATEFAILED = "FailedOperation.CreateTemplateFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINLISTFAILED = "FailedOperation.DescribeDomainListFailed"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  FAILEDOPERATION_PROHIBITPHONEEMAIL = "FailedOperation.ProhibitPhoneEmail"
//  FAILEDOPERATION_TEMPLATEMAXNUMFAILED = "FailedOperation.TemplateMaxNumFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATECODEISINVALID = "InvalidParameter.CertificateCodeIsInvalid"
//  INVALIDPARAMETER_CERTIFICATEIMAGEISINVALID = "InvalidParameter.CertificateImageIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_EMAILISUNVERIFIED = "InvalidParameter.EmailIsUnverified"
//  INVALIDPARAMETER_NAMEISINVALID = "InvalidParameter.NameIsInvalid"
//  INVALIDPARAMETER_NAMEISKEYWORD = "InvalidParameter.NameIsKeyword"
//  INVALIDPARAMETER_ORGISINVALID = "InvalidParameter.OrgIsInvalid"
//  INVALIDPARAMETER_ORGISKEYWORD = "InvalidParameter.OrgIsKeyword"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
//  INVALIDPARAMETER_STREETISINVALID = "InvalidParameter.StreetIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISUNVERIFIED = "InvalidParameter.TelephoneIsUnverified"
//  INVALIDPARAMETER_USERTYPEISINVALID = "InvalidParameter.UserTypeIsInvalid"
//  INVALIDPARAMETER_ZIPCODEISINVALID = "InvalidParameter.ZipCodeIsInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  MISSINGPARAMETER_TEMPLATEIDISEXIST = "MissingParameter.TemplateIdIsExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) CreateTemplateWithContext(ctx context.Context, request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePhoneEmailRequest() (request *DeletePhoneEmailRequest) {
    request = &DeletePhoneEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DeletePhoneEmail")
    
    
    return
}

func NewDeletePhoneEmailResponse() (response *DeletePhoneEmailResponse) {
    response = &DeletePhoneEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePhoneEmail
// 此接口用于删除已验证的手机邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
func (c *Client) DeletePhoneEmail(request *DeletePhoneEmailRequest) (response *DeletePhoneEmailResponse, err error) {
    return c.DeletePhoneEmailWithContext(context.Background(), request)
}

// DeletePhoneEmail
// 此接口用于删除已验证的手机邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
func (c *Client) DeletePhoneEmailWithContext(ctx context.Context, request *DeletePhoneEmailRequest) (response *DeletePhoneEmailResponse, err error) {
    if request == nil {
        request = NewDeletePhoneEmailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePhoneEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePhoneEmailResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
    request = &DeleteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DeleteTemplate")
    
    
    return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
    response = &DeleteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTemplate
// 本接口 ( DeleteTemplate ) 用于删除信息模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETETEMPLATEFAILED = "FailedOperation.DeleteTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
    return c.DeleteTemplateWithContext(context.Background(), request)
}

// DeleteTemplate
// 本接口 ( DeleteTemplate ) 用于删除信息模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETETEMPLATEFAILED = "FailedOperation.DeleteTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) DeleteTemplateWithContext(ctx context.Context, request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchOperationLogDetailsRequest() (request *DescribeBatchOperationLogDetailsRequest) {
    request = &DescribeBatchOperationLogDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBatchOperationLogDetails")
    
    
    return
}

func NewDescribeBatchOperationLogDetailsResponse() (response *DescribeBatchOperationLogDetailsResponse) {
    response = &DescribeBatchOperationLogDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBatchOperationLogDetails
// 本接口 ( DescribeBatchOperationLogDetails ) 用于获取批量操作日志详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBatchOperationLogDetails(request *DescribeBatchOperationLogDetailsRequest) (response *DescribeBatchOperationLogDetailsResponse, err error) {
    return c.DescribeBatchOperationLogDetailsWithContext(context.Background(), request)
}

// DescribeBatchOperationLogDetails
// 本接口 ( DescribeBatchOperationLogDetails ) 用于获取批量操作日志详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBatchOperationLogDetailsWithContext(ctx context.Context, request *DescribeBatchOperationLogDetailsRequest) (response *DescribeBatchOperationLogDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeBatchOperationLogDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchOperationLogDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchOperationLogDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchOperationLogsRequest() (request *DescribeBatchOperationLogsRequest) {
    request = &DescribeBatchOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBatchOperationLogs")
    
    
    return
}

func NewDescribeBatchOperationLogsResponse() (response *DescribeBatchOperationLogsResponse) {
    response = &DescribeBatchOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBatchOperationLogs
// 本接口 ( DescribeBatchOperationLogs ) 用于获取批量操作日志 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBatchOperationLogs(request *DescribeBatchOperationLogsRequest) (response *DescribeBatchOperationLogsResponse, err error) {
    return c.DescribeBatchOperationLogsWithContext(context.Background(), request)
}

// DescribeBatchOperationLogs
// 本接口 ( DescribeBatchOperationLogs ) 用于获取批量操作日志 。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBatchOperationLogsWithContext(ctx context.Context, request *DescribeBatchOperationLogsRequest) (response *DescribeBatchOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeBatchOperationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainBaseInfoRequest() (request *DescribeDomainBaseInfoRequest) {
    request = &DescribeDomainBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainBaseInfo")
    
    
    return
}

func NewDescribeDomainBaseInfoResponse() (response *DescribeDomainBaseInfoResponse) {
    response = &DescribeDomainBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainBaseInfo
// 本接口 (  DescribeDomainBaseInfo) 获取域名基本信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribeDomainBaseInfo(request *DescribeDomainBaseInfoRequest) (response *DescribeDomainBaseInfoResponse, err error) {
    return c.DescribeDomainBaseInfoWithContext(context.Background(), request)
}

// DescribeDomainBaseInfo
// 本接口 (  DescribeDomainBaseInfo) 获取域名基本信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribeDomainBaseInfoWithContext(ctx context.Context, request *DescribeDomainBaseInfoRequest) (response *DescribeDomainBaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainBaseInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainBaseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainNameListRequest() (request *DescribeDomainNameListRequest) {
    request = &DescribeDomainNameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainNameList")
    
    
    return
}

func NewDescribeDomainNameListResponse() (response *DescribeDomainNameListResponse) {
    response = &DescribeDomainNameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainNameList
// 本接口 (  DescribeDomainNameList ) 我的域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEDOMAINLISTFAILED = "FailedOperation.DescribeDomainListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDomainNameList(request *DescribeDomainNameListRequest) (response *DescribeDomainNameListResponse, err error) {
    return c.DescribeDomainNameListWithContext(context.Background(), request)
}

// DescribeDomainNameList
// 本接口 (  DescribeDomainNameList ) 我的域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEDOMAINLISTFAILED = "FailedOperation.DescribeDomainListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDomainNameListWithContext(ctx context.Context, request *DescribeDomainNameListRequest) (response *DescribeDomainNameListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainNameListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainNameList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainNameListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainPriceListRequest() (request *DescribeDomainPriceListRequest) {
    request = &DescribeDomainPriceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainPriceList")
    
    
    return
}

func NewDescribeDomainPriceListResponse() (response *DescribeDomainPriceListResponse) {
    response = &DescribeDomainPriceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainPriceList
// 按照域名后缀获取对应的价格列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINPRICELISTFAILED = "FailedOperation.DomainPriceListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDomainPriceList(request *DescribeDomainPriceListRequest) (response *DescribeDomainPriceListResponse, err error) {
    return c.DescribeDomainPriceListWithContext(context.Background(), request)
}

// DescribeDomainPriceList
// 按照域名后缀获取对应的价格列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINPRICELISTFAILED = "FailedOperation.DomainPriceListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDomainPriceListWithContext(ctx context.Context, request *DescribeDomainPriceListRequest) (response *DescribeDomainPriceListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPriceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainPriceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainPriceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainSimpleInfoRequest() (request *DescribeDomainSimpleInfoRequest) {
    request = &DescribeDomainSimpleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainSimpleInfo")
    
    
    return
}

func NewDescribeDomainSimpleInfoResponse() (response *DescribeDomainSimpleInfoResponse) {
    response = &DescribeDomainSimpleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainSimpleInfo
// 获取域名实名信息详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribeDomainSimpleInfo(request *DescribeDomainSimpleInfoRequest) (response *DescribeDomainSimpleInfoResponse, err error) {
    return c.DescribeDomainSimpleInfoWithContext(context.Background(), request)
}

// DescribeDomainSimpleInfo
// 获取域名实名信息详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribeDomainSimpleInfoWithContext(ctx context.Context, request *DescribeDomainSimpleInfoRequest) (response *DescribeDomainSimpleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainSimpleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainSimpleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDomainSimpleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePhoneEmailListRequest() (request *DescribePhoneEmailListRequest) {
    request = &DescribePhoneEmailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribePhoneEmailList")
    
    
    return
}

func NewDescribePhoneEmailListResponse() (response *DescribePhoneEmailListResponse) {
    response = &DescribePhoneEmailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePhoneEmailList
// 本接口用于获取已验证的手机邮箱列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePhoneEmailList(request *DescribePhoneEmailListRequest) (response *DescribePhoneEmailListResponse, err error) {
    return c.DescribePhoneEmailListWithContext(context.Background(), request)
}

// DescribePhoneEmailList
// 本接口用于获取已验证的手机邮箱列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePhoneEmailListWithContext(ctx context.Context, request *DescribePhoneEmailListRequest) (response *DescribePhoneEmailListResponse, err error) {
    if request == nil {
        request = NewDescribePhoneEmailListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePhoneEmailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePhoneEmailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
    request = &DescribeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeTemplate")
    
    
    return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
    response = &DescribeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTemplate
// 本接口 (DescribeTemplate) 用于获取模板信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
    return c.DescribeTemplateWithContext(context.Background(), request)
}

// DescribeTemplate
// 本接口 (DescribeTemplate) 用于获取模板信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
func (c *Client) DescribeTemplateWithContext(ctx context.Context, request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateListRequest() (request *DescribeTemplateListRequest) {
    request = &DescribeTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeTemplateList")
    
    
    return
}

func NewDescribeTemplateListResponse() (response *DescribeTemplateListResponse) {
    response = &DescribeTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTemplateList
// 本接口 (DescribeTemplateList) 用于获取信息模板列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTemplateList(request *DescribeTemplateListRequest) (response *DescribeTemplateListResponse, err error) {
    return c.DescribeTemplateListWithContext(context.Background(), request)
}

// DescribeTemplateList
// 本接口 (DescribeTemplateList) 用于获取信息模板列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTemplateListWithContext(ctx context.Context, request *DescribeTemplateListRequest) (response *DescribeTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainDNSBatchRequest() (request *ModifyDomainDNSBatchRequest) {
    request = &ModifyDomainDNSBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "ModifyDomainDNSBatch")
    
    
    return
}

func NewModifyDomainDNSBatchResponse() (response *ModifyDomainDNSBatchResponse) {
    response = &ModifyDomainDNSBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainDNSBatch
// 本接口 ( ModifyDomainDNSBatch) 用于批量域名 DNS 修改 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXPIREDUNSUPPORTED = "FailedOperation.DomainExpiredUnsupported"
//  FAILEDOPERATION_SETDOMAINDNSFAILED = "FailedOperation.SetDomainDnsFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_MODIFYDOMAININFOUNSUPPORTED = "UnsupportedOperation.ModifyDomainInfoUnsupported"
//  UNSUPPORTEDOPERATION_MODIFYDOMAINUNSUPPORTED = "UnsupportedOperation.ModifyDomainUnsupported"
func (c *Client) ModifyDomainDNSBatch(request *ModifyDomainDNSBatchRequest) (response *ModifyDomainDNSBatchResponse, err error) {
    return c.ModifyDomainDNSBatchWithContext(context.Background(), request)
}

// ModifyDomainDNSBatch
// 本接口 ( ModifyDomainDNSBatch) 用于批量域名 DNS 修改 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOMAINEXPIREDUNSUPPORTED = "FailedOperation.DomainExpiredUnsupported"
//  FAILEDOPERATION_SETDOMAINDNSFAILED = "FailedOperation.SetDomainDnsFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_MODIFYDOMAININFOUNSUPPORTED = "UnsupportedOperation.ModifyDomainInfoUnsupported"
//  UNSUPPORTEDOPERATION_MODIFYDOMAINUNSUPPORTED = "UnsupportedOperation.ModifyDomainUnsupported"
func (c *Client) ModifyDomainDNSBatchWithContext(ctx context.Context, request *ModifyDomainDNSBatchRequest) (response *ModifyDomainDNSBatchResponse, err error) {
    if request == nil {
        request = NewModifyDomainDNSBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainDNSBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainDNSBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainOwnerBatchRequest() (request *ModifyDomainOwnerBatchRequest) {
    request = &ModifyDomainOwnerBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "ModifyDomainOwnerBatch")
    
    
    return
}

func NewModifyDomainOwnerBatchResponse() (response *ModifyDomainOwnerBatchResponse) {
    response = &ModifyDomainOwnerBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainOwnerBatch
// 本接口 ( ModifyDomainOwnerBatch) 用于域名批量账号间转移 。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYDOMAINOWNERFAILED = "FailedOperation.ModifyDomainOwnerFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) ModifyDomainOwnerBatch(request *ModifyDomainOwnerBatchRequest) (response *ModifyDomainOwnerBatchResponse, err error) {
    return c.ModifyDomainOwnerBatchWithContext(context.Background(), request)
}

// ModifyDomainOwnerBatch
// 本接口 ( ModifyDomainOwnerBatch) 用于域名批量账号间转移 。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYDOMAINOWNERFAILED = "FailedOperation.ModifyDomainOwnerFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) ModifyDomainOwnerBatchWithContext(ctx context.Context, request *ModifyDomainOwnerBatchRequest) (response *ModifyDomainOwnerBatchResponse, err error) {
    if request == nil {
        request = NewModifyDomainOwnerBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainOwnerBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainOwnerBatchResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDomainBatchRequest() (request *RenewDomainBatchRequest) {
    request = &RenewDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "RenewDomainBatch")
    
    
    return
}

func NewRenewDomainBatchResponse() (response *RenewDomainBatchResponse) {
    response = &RenewDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewDomainBatch
// 本接口 ( RenewDomainBatch ) 用于批量续费域名 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) RenewDomainBatch(request *RenewDomainBatchRequest) (response *RenewDomainBatchResponse, err error) {
    return c.RenewDomainBatchWithContext(context.Background(), request)
}

// RenewDomainBatch
// 本接口 ( RenewDomainBatch ) 用于批量续费域名 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) RenewDomainBatchWithContext(ctx context.Context, request *RenewDomainBatchRequest) (response *RenewDomainBatchResponse, err error) {
    if request == nil {
        request = NewRenewDomainBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewSendPhoneEmailCodeRequest() (request *SendPhoneEmailCodeRequest) {
    request = &SendPhoneEmailCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "SendPhoneEmailCode")
    
    
    return
}

func NewSendPhoneEmailCodeResponse() (response *SendPhoneEmailCodeResponse) {
    response = &SendPhoneEmailCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendPhoneEmailCode
// 此接口用于发送手机邮箱验证码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SENDTCBPHONEEMAILCODEFAILED = "FailedOperation.SendTcbPhoneEmailCodeFailed"
//  FAILEDOPERATION_SENDVERIFYCODEISLIMITED = "FailedOperation.SendVerifyCodeIsLimited"
//  FAILEDOPERATION_VERIFYUINISREALNAME = "FailedOperation.VerifyUinIsRealname"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CODETYPEISINVALID = "InvalidParameter.CodeTypeIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  LIMITEXCEEDED_REQUESTLIMIT = "LimitExceeded.RequestLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) SendPhoneEmailCode(request *SendPhoneEmailCodeRequest) (response *SendPhoneEmailCodeResponse, err error) {
    return c.SendPhoneEmailCodeWithContext(context.Background(), request)
}

// SendPhoneEmailCode
// 此接口用于发送手机邮箱验证码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SENDTCBPHONEEMAILCODEFAILED = "FailedOperation.SendTcbPhoneEmailCodeFailed"
//  FAILEDOPERATION_SENDVERIFYCODEISLIMITED = "FailedOperation.SendVerifyCodeIsLimited"
//  FAILEDOPERATION_VERIFYUINISREALNAME = "FailedOperation.VerifyUinIsRealname"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CODETYPEISINVALID = "InvalidParameter.CodeTypeIsInvalid"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
//  LIMITEXCEEDED_REQUESTLIMIT = "LimitExceeded.RequestLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) SendPhoneEmailCodeWithContext(ctx context.Context, request *SendPhoneEmailCodeRequest) (response *SendPhoneEmailCodeResponse, err error) {
    if request == nil {
        request = NewSendPhoneEmailCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendPhoneEmailCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendPhoneEmailCodeResponse()
    err = c.Send(request, response)
    return
}

func NewSetDomainAutoRenewRequest() (request *SetDomainAutoRenewRequest) {
    request = &SetDomainAutoRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "SetDomainAutoRenew")
    
    
    return
}

func NewSetDomainAutoRenewResponse() (response *SetDomainAutoRenewResponse) {
    response = &SetDomainAutoRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetDomainAutoRenew
// 本接口 ( SetDomainAutoRenew ) 用于设置域名自动续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) SetDomainAutoRenew(request *SetDomainAutoRenewRequest) (response *SetDomainAutoRenewResponse, err error) {
    return c.SetDomainAutoRenewWithContext(context.Background(), request)
}

// SetDomainAutoRenew
// 本接口 ( SetDomainAutoRenew ) 用于设置域名自动续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKDOMAINFAILED = "FailedOperation.CheckDomainFailed"
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) SetDomainAutoRenewWithContext(ctx context.Context, request *SetDomainAutoRenewRequest) (response *SetDomainAutoRenewResponse, err error) {
    if request == nil {
        request = NewSetDomainAutoRenewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDomainAutoRenew require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDomainAutoRenewResponse()
    err = c.Send(request, response)
    return
}

func NewTransferInDomainBatchRequest() (request *TransferInDomainBatchRequest) {
    request = &TransferInDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "TransferInDomainBatch")
    
    
    return
}

func NewTransferInDomainBatchResponse() (response *TransferInDomainBatchResponse) {
    response = &TransferInDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransferInDomainBatch
// 本接口 ( TransferInDomainBatch ) 用于批量转入域名 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_DUPLICATEDOMAINEXISTS = "InvalidParameter.DuplicateDomainExists"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) TransferInDomainBatch(request *TransferInDomainBatchRequest) (response *TransferInDomainBatchResponse, err error) {
    return c.TransferInDomainBatchWithContext(context.Background(), request)
}

// TransferInDomainBatch
// 本接口 ( TransferInDomainBatch ) 用于批量转入域名 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_DUPLICATEDOMAINEXISTS = "InvalidParameter.DuplicateDomainExists"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
//  RESOURCENOTFOUND_APPROVEDTEMPLATENOTFOUND = "ResourceNotFound.ApprovedTemplateNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  RESOURCEUNAVAILABLE_DOMAINISMODIFYINGDNS = "ResourceUnavailable.DomainIsModifyingDNS"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) TransferInDomainBatchWithContext(ctx context.Context, request *TransferInDomainBatchRequest) (response *TransferInDomainBatchResponse, err error) {
    if request == nil {
        request = NewTransferInDomainBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransferInDomainBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransferInDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewTransferProhibitionBatchRequest() (request *TransferProhibitionBatchRequest) {
    request = &TransferProhibitionBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "TransferProhibitionBatch")
    
    
    return
}

func NewTransferProhibitionBatchResponse() (response *TransferProhibitionBatchResponse) {
    response = &TransferProhibitionBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransferProhibitionBatch
// 本接口 ( TransferProhibitionBatch ) 用于批量禁止域名转移 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) TransferProhibitionBatch(request *TransferProhibitionBatchRequest) (response *TransferProhibitionBatchResponse, err error) {
    return c.TransferProhibitionBatchWithContext(context.Background(), request)
}

// TransferProhibitionBatch
// 本接口 ( TransferProhibitionBatch ) 用于批量禁止域名转移 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) TransferProhibitionBatchWithContext(ctx context.Context, request *TransferProhibitionBatchRequest) (response *TransferProhibitionBatchResponse, err error) {
    if request == nil {
        request = NewTransferProhibitionBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransferProhibitionBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransferProhibitionBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProhibitionBatchRequest() (request *UpdateProhibitionBatchRequest) {
    request = &UpdateProhibitionBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "UpdateProhibitionBatch")
    
    
    return
}

func NewUpdateProhibitionBatchResponse() (response *UpdateProhibitionBatchResponse) {
    response = &UpdateProhibitionBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProhibitionBatch
// 本接口 ( UpdateProhibitionBatch ) 用于批量禁止更新锁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) UpdateProhibitionBatch(request *UpdateProhibitionBatchRequest) (response *UpdateProhibitionBatchResponse, err error) {
    return c.UpdateProhibitionBatchWithContext(context.Background(), request)
}

// UpdateProhibitionBatch
// 本接口 ( UpdateProhibitionBatch ) 用于批量禁止更新锁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UPTO4000 = "InvalidParameter.UpTo4000"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  MISSINGPARAMETER_REPDATAISNONE = "MissingParameter.RepDataIsNone"
//  RESOURCEINSUFFICIENT_OVERWORK = "ResourceInsufficient.Overwork"
func (c *Client) UpdateProhibitionBatchWithContext(ctx context.Context, request *UpdateProhibitionBatchRequest) (response *UpdateProhibitionBatchResponse, err error) {
    if request == nil {
        request = NewUpdateProhibitionBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProhibitionBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProhibitionBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUploadImageRequest() (request *UploadImageRequest) {
    request = &UploadImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "UploadImage")
    
    
    return
}

func NewUploadImageResponse() (response *UploadImageResponse) {
    response = &UploadImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadImage
// 本接口 ( UploadImage ) 用于证件图片上传 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  FAILEDOPERATION_UPLOADIMAGEFAILED = "FailedOperation.UploadImageFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEEXTINVALID = "InvalidParameter.ImageExtInvalid"
//  INVALIDPARAMETER_IMAGEFILEISINVALID = "InvalidParameter.ImageFileIsInvalid"
//  INVALIDPARAMETER_IMAGEFORMATISINVALID = "InvalidParameter.ImageFormatIsInvalid"
//  INVALIDPARAMETER_IMAGESIZEBELOW = "InvalidParameter.ImageSizeBelow"
//  INVALIDPARAMETER_IMAGESIZEEXCEED = "InvalidParameter.ImageSizeExceed"
//  INVALIDPARAMETER_IMAGESIZELIMIT = "InvalidParameter.ImageSizeLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UploadImage(request *UploadImageRequest) (response *UploadImageResponse, err error) {
    return c.UploadImageWithContext(context.Background(), request)
}

// UploadImage
// 本接口 ( UploadImage ) 用于证件图片上传 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  FAILEDOPERATION_UPLOADIMAGEFAILED = "FailedOperation.UploadImageFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEEXTINVALID = "InvalidParameter.ImageExtInvalid"
//  INVALIDPARAMETER_IMAGEFILEISINVALID = "InvalidParameter.ImageFileIsInvalid"
//  INVALIDPARAMETER_IMAGEFORMATISINVALID = "InvalidParameter.ImageFormatIsInvalid"
//  INVALIDPARAMETER_IMAGESIZEBELOW = "InvalidParameter.ImageSizeBelow"
//  INVALIDPARAMETER_IMAGESIZEEXCEED = "InvalidParameter.ImageSizeExceed"
//  INVALIDPARAMETER_IMAGESIZELIMIT = "InvalidParameter.ImageSizeLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UploadImageWithContext(ctx context.Context, request *UploadImageRequest) (response *UploadImageResponse, err error) {
    if request == nil {
        request = NewUploadImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadImageResponse()
    err = c.Send(request, response)
    return
}
