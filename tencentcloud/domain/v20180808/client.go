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

func NewBidDetailPageRequest() (request *BidDetailPageRequest) {
    request = &BidDetailPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "BidDetailPage")
    
    
    return
}

func NewBidDetailPageResponse() (response *BidDetailPageResponse) {
    response = &BidDetailPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BidDetailPage
// 该接口用于用户详情页出价请求
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
func (c *Client) BidDetailPage(request *BidDetailPageRequest) (response *BidDetailPageResponse, err error) {
    return c.BidDetailPageWithContext(context.Background(), request)
}

// BidDetailPage
// 该接口用于用户详情页出价请求
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
func (c *Client) BidDetailPageWithContext(ctx context.Context, request *BidDetailPageRequest) (response *BidDetailPageResponse, err error) {
    if request == nil {
        request = NewBidDetailPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BidDetailPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewBidDetailPageResponse()
    err = c.Send(request, response)
    return
}

func NewBidPreDomainsRequest() (request *BidPreDomainsRequest) {
    request = &BidPreDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "BidPreDomains")
    
    
    return
}

func NewBidPreDomainsResponse() (response *BidPreDomainsResponse) {
    response = &BidPreDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BidPreDomains
// 用户合作商预释放出价
//
// 可能返回的错误码:
//  FAILEDOPERATION_BIDPREDOMAINSPRICEDOWNERR = "FailedOperation.BidPreDomainsPriceDownErr"
//  FAILEDOPERATION_BIDPREDOMAINSTIMEOUTERR = "FailedOperation.BidPreDomainsTimeOutErr"
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) BidPreDomains(request *BidPreDomainsRequest) (response *BidPreDomainsResponse, err error) {
    return c.BidPreDomainsWithContext(context.Background(), request)
}

// BidPreDomains
// 用户合作商预释放出价
//
// 可能返回的错误码:
//  FAILEDOPERATION_BIDPREDOMAINSPRICEDOWNERR = "FailedOperation.BidPreDomainsPriceDownErr"
//  FAILEDOPERATION_BIDPREDOMAINSTIMEOUTERR = "FailedOperation.BidPreDomainsTimeOutErr"
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) BidPreDomainsWithContext(ctx context.Context, request *BidPreDomainsRequest) (response *BidPreDomainsResponse, err error) {
    if request == nil {
        request = NewBidPreDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BidPreDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewBidPreDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewBiddingPreReleaseRequest() (request *BiddingPreReleaseRequest) {
    request = &BiddingPreReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "BiddingPreRelease")
    
    
    return
}

func NewBiddingPreReleaseResponse() (response *BiddingPreReleaseResponse) {
    response = &BiddingPreReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BiddingPreRelease
// 用于出价界面出价请求
//
// 可能返回的错误码:
//  FAILEDOPERATION_BIDCURRENTPRICE = "FailedOperation.BidCurrentPrice"
//  FAILEDOPERATION_BIDPRICEILLEGAL = "FailedOperation.BidPriceIllegal"
//  FAILEDOPERATION_BIDDINGGETPRICEDOING = "FailedOperation.BiddingGetPriceDoing"
//  RESOURCENOTFOUND_BIDPRICECONFIG = "ResourceNotFound.BidPriceConfig"
func (c *Client) BiddingPreRelease(request *BiddingPreReleaseRequest) (response *BiddingPreReleaseResponse, err error) {
    return c.BiddingPreReleaseWithContext(context.Background(), request)
}

// BiddingPreRelease
// 用于出价界面出价请求
//
// 可能返回的错误码:
//  FAILEDOPERATION_BIDCURRENTPRICE = "FailedOperation.BidCurrentPrice"
//  FAILEDOPERATION_BIDPRICEILLEGAL = "FailedOperation.BidPriceIllegal"
//  FAILEDOPERATION_BIDDINGGETPRICEDOING = "FailedOperation.BiddingGetPriceDoing"
//  RESOURCENOTFOUND_BIDPRICECONFIG = "ResourceNotFound.BidPriceConfig"
func (c *Client) BiddingPreReleaseWithContext(ctx context.Context, request *BiddingPreReleaseRequest) (response *BiddingPreReleaseResponse, err error) {
    if request == nil {
        request = NewBiddingPreReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BiddingPreRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewBiddingPreReleaseResponse()
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

func NewCreateCustomDnsHostRequest() (request *CreateCustomDnsHostRequest) {
    request = &CreateCustomDnsHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreateCustomDnsHost")
    
    
    return
}

func NewCreateCustomDnsHostResponse() (response *CreateCustomDnsHostResponse) {
    response = &CreateCustomDnsHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCustomDnsHost
// 创建自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_DUPLICATECUSTOMDNSNAME = "InvalidParameter.DuplicateCustomDnsName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_CUSTOMHOSTOVERLIMIT = "UnsupportedOperation.CustomHostOverLimit"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) CreateCustomDnsHost(request *CreateCustomDnsHostRequest) (response *CreateCustomDnsHostResponse, err error) {
    return c.CreateCustomDnsHostWithContext(context.Background(), request)
}

// CreateCustomDnsHost
// 创建自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  INVALIDPARAMETER_DUPLICATECUSTOMDNSNAME = "InvalidParameter.DuplicateCustomDnsName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_CUSTOMHOSTOVERLIMIT = "UnsupportedOperation.CustomHostOverLimit"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) CreateCustomDnsHostWithContext(ctx context.Context, request *CreateCustomDnsHostRequest) (response *CreateCustomDnsHostResponse, err error) {
    if request == nil {
        request = NewCreateCustomDnsHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomDnsHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomDnsHostResponse()
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

func NewCreateDomainRedemptionRequest() (request *CreateDomainRedemptionRequest) {
    request = &CreateDomainRedemptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "CreateDomainRedemption")
    
    
    return
}

func NewCreateDomainRedemptionResponse() (response *CreateDomainRedemptionResponse) {
    response = &CreateDomainRedemptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDomainRedemption
// 创建赎回订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
func (c *Client) CreateDomainRedemption(request *CreateDomainRedemptionRequest) (response *CreateDomainRedemptionResponse, err error) {
    return c.CreateDomainRedemptionWithContext(context.Background(), request)
}

// CreateDomainRedemption
// 创建赎回订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"
func (c *Client) CreateDomainRedemptionWithContext(ctx context.Context, request *CreateDomainRedemptionRequest) (response *CreateDomainRedemptionResponse, err error) {
    if request == nil {
        request = NewCreateDomainRedemptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDomainRedemption require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDomainRedemptionResponse()
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

func NewDeleteBiddingRequest() (request *DeleteBiddingRequest) {
    request = &DeleteBiddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DeleteBidding")
    
    
    return
}

func NewDeleteBiddingResponse() (response *DeleteBiddingResponse) {
    response = &DeleteBiddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBidding
// 删除记录。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DeleteBidding(request *DeleteBiddingRequest) (response *DeleteBiddingResponse, err error) {
    return c.DeleteBiddingWithContext(context.Background(), request)
}

// DeleteBidding
// 删除记录。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DeleteBiddingWithContext(ctx context.Context, request *DeleteBiddingRequest) (response *DeleteBiddingResponse, err error) {
    if request == nil {
        request = NewDeleteBiddingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBidding require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBiddingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomDnsHostRequest() (request *DeleteCustomDnsHostRequest) {
    request = &DeleteCustomDnsHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DeleteCustomDnsHost")
    
    
    return
}

func NewDeleteCustomDnsHostResponse() (response *DeleteCustomDnsHostResponse) {
    response = &DeleteCustomDnsHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomDnsHost
// 删除自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_DOMAINISINVALID = "InvalidParameter.DomainIsInvalid"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) DeleteCustomDnsHost(request *DeleteCustomDnsHostRequest) (response *DeleteCustomDnsHostResponse, err error) {
    return c.DeleteCustomDnsHostWithContext(context.Background(), request)
}

// DeleteCustomDnsHost
// 删除自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  INVALIDPARAMETER_DOMAINISINVALID = "InvalidParameter.DomainIsInvalid"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) DeleteCustomDnsHostWithContext(ctx context.Context, request *DeleteCustomDnsHostRequest) (response *DeleteCustomDnsHostResponse, err error) {
    if request == nil {
        request = NewDeleteCustomDnsHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomDnsHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomDnsHostResponse()
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

func NewDeleteReservedPreDomainInfoRequest() (request *DeleteReservedPreDomainInfoRequest) {
    request = &DeleteReservedPreDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DeleteReservedPreDomainInfo")
    
    
    return
}

func NewDeleteReservedPreDomainInfoResponse() (response *DeleteReservedPreDomainInfoResponse) {
    response = &DeleteReservedPreDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReservedPreDomainInfo
// 用于清除多余的预定域名信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
func (c *Client) DeleteReservedPreDomainInfo(request *DeleteReservedPreDomainInfoRequest) (response *DeleteReservedPreDomainInfoResponse, err error) {
    return c.DeleteReservedPreDomainInfoWithContext(context.Background(), request)
}

// DeleteReservedPreDomainInfo
// 用于清除多余的预定域名信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMAILISINVALID = "InvalidParameter.EmailIsInvalid"
//  INVALIDPARAMETER_TELEPHONEISINVALID = "InvalidParameter.TelephoneIsInvalid"
func (c *Client) DeleteReservedPreDomainInfoWithContext(ctx context.Context, request *DeleteReservedPreDomainInfoRequest) (response *DeleteReservedPreDomainInfoResponse, err error) {
    if request == nil {
        request = NewDeleteReservedPreDomainInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReservedPreDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReservedPreDomainInfoResponse()
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
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
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
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
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

func NewDescribeAuctionListRequest() (request *DescribeAuctionListRequest) {
    request = &DescribeAuctionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeAuctionList")
    
    
    return
}

func NewDescribeAuctionListResponse() (response *DescribeAuctionListResponse) {
    response = &DescribeAuctionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuctionList
// 用户控制台获取竞价列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETETEMPLATEFAILED = "FailedOperation.DeleteTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) DescribeAuctionList(request *DescribeAuctionListRequest) (response *DescribeAuctionListResponse, err error) {
    return c.DescribeAuctionListWithContext(context.Background(), request)
}

// DescribeAuctionList
// 用户控制台获取竞价列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETETEMPLATEFAILED = "FailedOperation.DeleteTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER_TEMPLATEIDISEMPTY = "MissingParameter.TemplateIdIsEmpty"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) DescribeAuctionListWithContext(ctx context.Context, request *DescribeAuctionListRequest) (response *DescribeAuctionListResponse, err error) {
    if request == nil {
        request = NewDescribeAuctionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuctionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuctionListResponse()
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

func NewDescribeBiddingAppointDetailRequest() (request *DescribeBiddingAppointDetailRequest) {
    request = &DescribeBiddingAppointDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBiddingAppointDetail")
    
    
    return
}

func NewDescribeBiddingAppointDetailResponse() (response *DescribeBiddingAppointDetailResponse) {
    response = &DescribeBiddingAppointDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBiddingAppointDetail
// 我预约的域名-预约详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingAppointDetail(request *DescribeBiddingAppointDetailRequest) (response *DescribeBiddingAppointDetailResponse, err error) {
    return c.DescribeBiddingAppointDetailWithContext(context.Background(), request)
}

// DescribeBiddingAppointDetail
// 我预约的域名-预约详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingAppointDetailWithContext(ctx context.Context, request *DescribeBiddingAppointDetailRequest) (response *DescribeBiddingAppointDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBiddingAppointDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBiddingAppointDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBiddingAppointDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBiddingAppointListRequest() (request *DescribeBiddingAppointListRequest) {
    request = &DescribeBiddingAppointListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBiddingAppointList")
    
    
    return
}

func NewDescribeBiddingAppointListResponse() (response *DescribeBiddingAppointListResponse) {
    response = &DescribeBiddingAppointListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBiddingAppointList
// 我预定的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingAppointList(request *DescribeBiddingAppointListRequest) (response *DescribeBiddingAppointListResponse, err error) {
    return c.DescribeBiddingAppointListWithContext(context.Background(), request)
}

// DescribeBiddingAppointList
// 我预定的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingAppointListWithContext(ctx context.Context, request *DescribeBiddingAppointListRequest) (response *DescribeBiddingAppointListResponse, err error) {
    if request == nil {
        request = NewDescribeBiddingAppointListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBiddingAppointList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBiddingAppointListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBiddingDetailRequest() (request *DescribeBiddingDetailRequest) {
    request = &DescribeBiddingDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBiddingDetail")
    
    
    return
}

func NewDescribeBiddingDetailResponse() (response *DescribeBiddingDetailResponse) {
    response = &DescribeBiddingDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBiddingDetail
// 我竞价的域名-竞价详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingDetail(request *DescribeBiddingDetailRequest) (response *DescribeBiddingDetailResponse, err error) {
    return c.DescribeBiddingDetailWithContext(context.Background(), request)
}

// DescribeBiddingDetail
// 我竞价的域名-竞价详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingDetailWithContext(ctx context.Context, request *DescribeBiddingDetailRequest) (response *DescribeBiddingDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBiddingDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBiddingDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBiddingDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBiddingListRequest() (request *DescribeBiddingListRequest) {
    request = &DescribeBiddingListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBiddingList")
    
    
    return
}

func NewDescribeBiddingListResponse() (response *DescribeBiddingListResponse) {
    response = &DescribeBiddingListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBiddingList
// 我竞价的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingList(request *DescribeBiddingListRequest) (response *DescribeBiddingListResponse, err error) {
    return c.DescribeBiddingListWithContext(context.Background(), request)
}

// DescribeBiddingList
// 我竞价的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingListWithContext(ctx context.Context, request *DescribeBiddingListRequest) (response *DescribeBiddingListResponse, err error) {
    if request == nil {
        request = NewDescribeBiddingListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBiddingList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBiddingListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBiddingSuccessfulDetailRequest() (request *DescribeBiddingSuccessfulDetailRequest) {
    request = &DescribeBiddingSuccessfulDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBiddingSuccessfulDetail")
    
    
    return
}

func NewDescribeBiddingSuccessfulDetailResponse() (response *DescribeBiddingSuccessfulDetailResponse) {
    response = &DescribeBiddingSuccessfulDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBiddingSuccessfulDetail
// 我得标的域名-得标详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingSuccessfulDetail(request *DescribeBiddingSuccessfulDetailRequest) (response *DescribeBiddingSuccessfulDetailResponse, err error) {
    return c.DescribeBiddingSuccessfulDetailWithContext(context.Background(), request)
}

// DescribeBiddingSuccessfulDetail
// 我得标的域名-得标详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingSuccessfulDetailWithContext(ctx context.Context, request *DescribeBiddingSuccessfulDetailRequest) (response *DescribeBiddingSuccessfulDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBiddingSuccessfulDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBiddingSuccessfulDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBiddingSuccessfulDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBiddingSuccessfulListRequest() (request *DescribeBiddingSuccessfulListRequest) {
    request = &DescribeBiddingSuccessfulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBiddingSuccessfulList")
    
    
    return
}

func NewDescribeBiddingSuccessfulListResponse() (response *DescribeBiddingSuccessfulListResponse) {
    response = &DescribeBiddingSuccessfulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBiddingSuccessfulList
// 我得标的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingSuccessfulList(request *DescribeBiddingSuccessfulListRequest) (response *DescribeBiddingSuccessfulListResponse, err error) {
    return c.DescribeBiddingSuccessfulListWithContext(context.Background(), request)
}

// DescribeBiddingSuccessfulList
// 我得标的域名。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESHOPCOLLECTLISTERR = "InternalError.DescribeShopCollectListErr"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) DescribeBiddingSuccessfulListWithContext(ctx context.Context, request *DescribeBiddingSuccessfulListRequest) (response *DescribeBiddingSuccessfulListResponse, err error) {
    if request == nil {
        request = NewDescribeBiddingSuccessfulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBiddingSuccessfulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBiddingSuccessfulListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomDnsHostSetRequest() (request *DescribeCustomDnsHostSetRequest) {
    request = &DescribeCustomDnsHostSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeCustomDnsHostSet")
    
    
    return
}

func NewDescribeCustomDnsHostSetResponse() (response *DescribeCustomDnsHostSetResponse) {
    response = &DescribeCustomDnsHostSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomDnsHostSet
// 查询自定义DNS Host
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribeCustomDnsHostSet(request *DescribeCustomDnsHostSetRequest) (response *DescribeCustomDnsHostSetResponse, err error) {
    return c.DescribeCustomDnsHostSetWithContext(context.Background(), request)
}

// DescribeCustomDnsHostSet
// 查询自定义DNS Host
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  MISSINGPARAMETER_DOMAINISEMPTY = "MissingParameter.DomainIsEmpty"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribeCustomDnsHostSetWithContext(ctx context.Context, request *DescribeCustomDnsHostSetRequest) (response *DescribeCustomDnsHostSetResponse, err error) {
    if request == nil {
        request = NewDescribeCustomDnsHostSetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomDnsHostSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomDnsHostSetResponse()
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

func NewDescribePayWaitDetailRequest() (request *DescribePayWaitDetailRequest) {
    request = &DescribePayWaitDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribePayWaitDetail")
    
    
    return
}

func NewDescribePayWaitDetailResponse() (response *DescribePayWaitDetailResponse) {
    response = &DescribePayWaitDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePayWaitDetail
// 等待支付详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribePayWaitDetail(request *DescribePayWaitDetailRequest) (response *DescribePayWaitDetailResponse, err error) {
    return c.DescribePayWaitDetailWithContext(context.Background(), request)
}

// DescribePayWaitDetail
// 等待支付详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEDOMAINFAILED = "FailedOperation.DescribeDomainFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
func (c *Client) DescribePayWaitDetailWithContext(ctx context.Context, request *DescribePayWaitDetailRequest) (response *DescribePayWaitDetailResponse, err error) {
    if request == nil {
        request = NewDescribePayWaitDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePayWaitDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePayWaitDetailResponse()
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

func NewDescribePreAuctionListRequest() (request *DescribePreAuctionListRequest) {
    request = &DescribePreAuctionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribePreAuctionList")
    
    
    return
}

func NewDescribePreAuctionListResponse() (response *DescribePreAuctionListResponse) {
    response = &DescribePreAuctionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePreAuctionList
// 用于预释放竞价列表数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePreAuctionList(request *DescribePreAuctionListRequest) (response *DescribePreAuctionListResponse, err error) {
    return c.DescribePreAuctionListWithContext(context.Background(), request)
}

// DescribePreAuctionList
// 用于预释放竞价列表数据查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePreAuctionListWithContext(ctx context.Context, request *DescribePreAuctionListRequest) (response *DescribePreAuctionListResponse, err error) {
    if request == nil {
        request = NewDescribePreAuctionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePreAuctionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePreAuctionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePreDomainListRequest() (request *DescribePreDomainListRequest) {
    request = &DescribePreDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribePreDomainList")
    
    
    return
}

func NewDescribePreDomainListResponse() (response *DescribePreDomainListResponse) {
    response = &DescribePreDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePreDomainList
// 用户服务商提前获取预释放域名数据，查询数据根据结束时间进行倒序。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DESCRIBEPREDOMAINLISTNOTBEGIN = "InternalError.DescribePreDomainListNotBegin"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribePreDomainList(request *DescribePreDomainListRequest) (response *DescribePreDomainListResponse, err error) {
    return c.DescribePreDomainListWithContext(context.Background(), request)
}

// DescribePreDomainList
// 用户服务商提前获取预释放域名数据，查询数据根据结束时间进行倒序。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DESCRIBEPREDOMAINLISTNOTBEGIN = "InternalError.DescribePreDomainListNotBegin"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribePreDomainListWithContext(ctx context.Context, request *DescribePreDomainListRequest) (response *DescribePreDomainListResponse, err error) {
    if request == nil {
        request = NewDescribePreDomainListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePreDomainList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePreDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePreReleaseListRequest() (request *DescribePreReleaseListRequest) {
    request = &DescribePreReleaseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribePreReleaseList")
    
    
    return
}

func NewDescribePreReleaseListResponse() (response *DescribePreReleaseListResponse) {
    response = &DescribePreReleaseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePreReleaseList
// 接口用于预释放页面查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DESCRIBEPREDOMAINLISTNOTBEGIN = "InternalError.DescribePreDomainListNotBegin"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribePreReleaseList(request *DescribePreReleaseListRequest) (response *DescribePreReleaseListResponse, err error) {
    return c.DescribePreReleaseListWithContext(context.Background(), request)
}

// DescribePreReleaseList
// 接口用于预释放页面查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DESCRIBEPREDOMAINLISTNOTBEGIN = "InternalError.DescribePreDomainListNotBegin"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribePreReleaseListWithContext(ctx context.Context, request *DescribePreReleaseListRequest) (response *DescribePreReleaseListResponse, err error) {
    if request == nil {
        request = NewDescribePreReleaseListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePreReleaseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePreReleaseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReservedBidInfoRequest() (request *DescribeReservedBidInfoRequest) {
    request = &DescribeReservedBidInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeReservedBidInfo")
    
    
    return
}

func NewDescribeReservedBidInfoResponse() (response *DescribeReservedBidInfoResponse) {
    response = &DescribeReservedBidInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReservedBidInfo
// 接口用于获取合作商竞价过程中竞价详情数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribeReservedBidInfo(request *DescribeReservedBidInfoRequest) (response *DescribeReservedBidInfoResponse, err error) {
    return c.DescribeReservedBidInfoWithContext(context.Background(), request)
}

// DescribeReservedBidInfo
// 接口用于获取合作商竞价过程中竞价详情数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribeReservedBidInfoWithContext(ctx context.Context, request *DescribeReservedBidInfoRequest) (response *DescribeReservedBidInfoResponse, err error) {
    if request == nil {
        request = NewDescribeReservedBidInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReservedBidInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReservedBidInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReservedPreDomainInfoRequest() (request *DescribeReservedPreDomainInfoRequest) {
    request = &DescribeReservedPreDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeReservedPreDomainInfo")
    
    
    return
}

func NewDescribeReservedPreDomainInfoResponse() (response *DescribeReservedPreDomainInfoResponse) {
    response = &DescribeReservedPreDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReservedPreDomainInfo
// 合作商用于查询预约预释放状态信息内容
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribeReservedPreDomainInfo(request *DescribeReservedPreDomainInfoRequest) (response *DescribeReservedPreDomainInfoResponse, err error) {
    return c.DescribeReservedPreDomainInfoWithContext(context.Background(), request)
}

// DescribeReservedPreDomainInfo
// 合作商用于查询预约预释放状态信息内容
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) DescribeReservedPreDomainInfoWithContext(ctx context.Context, request *DescribeReservedPreDomainInfoRequest) (response *DescribeReservedPreDomainInfoResponse, err error) {
    if request == nil {
        request = NewDescribeReservedPreDomainInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReservedPreDomainInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReservedPreDomainInfoResponse()
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

func NewDescribeTldListRequest() (request *DescribeTldListRequest) {
    request = &DescribeTldListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeTldList")
    
    
    return
}

func NewDescribeTldListResponse() (response *DescribeTldListResponse) {
    response = &DescribeTldListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTldList
// 用于获取域名注册当前支持注册的后缀
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTldList(request *DescribeTldListRequest) (response *DescribeTldListResponse, err error) {
    return c.DescribeTldListWithContext(context.Background(), request)
}

// DescribeTldList
// 用于获取域名注册当前支持注册的后缀
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTldListWithContext(ctx context.Context, request *DescribeTldListRequest) (response *DescribeTldListResponse, err error) {
    if request == nil {
        request = NewDescribeTldListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTldList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTldListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnPreDomainDetailRequest() (request *DescribeUnPreDomainDetailRequest) {
    request = &DescribeUnPreDomainDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "DescribeUnPreDomainDetail")
    
    
    return
}

func NewDescribeUnPreDomainDetailResponse() (response *DescribeUnPreDomainDetailResponse) {
    response = &DescribeUnPreDomainDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnPreDomainDetail
// 查询预释放未预约域名详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUnPreDomainDetail(request *DescribeUnPreDomainDetailRequest) (response *DescribeUnPreDomainDetailResponse, err error) {
    return c.DescribeUnPreDomainDetailWithContext(context.Background(), request)
}

// DescribeUnPreDomainDetail
// 查询预释放未预约域名详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeUnPreDomainDetailWithContext(ctx context.Context, request *DescribeUnPreDomainDetailRequest) (response *DescribeUnPreDomainDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUnPreDomainDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnPreDomainDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnPreDomainDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomDnsHostRequest() (request *ModifyCustomDnsHostRequest) {
    request = &ModifyCustomDnsHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "ModifyCustomDnsHost")
    
    
    return
}

func NewModifyCustomDnsHostResponse() (response *ModifyCustomDnsHostResponse) {
    response = &ModifyCustomDnsHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCustomDnsHost
// 修改自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_DOMAINISINVALID = "InvalidParameter.DomainIsInvalid"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) ModifyCustomDnsHost(request *ModifyCustomDnsHostRequest) (response *ModifyCustomDnsHostResponse, err error) {
    return c.ModifyCustomDnsHostWithContext(context.Background(), request)
}

// ModifyCustomDnsHost
// 修改自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_DOMAINISINVALID = "InvalidParameter.DomainIsInvalid"
//  INVALIDPARAMETER_DOMAINNAMEISINVALID = "InvalidParameter.DomainNameIsInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) ModifyCustomDnsHostWithContext(ctx context.Context, request *ModifyCustomDnsHostRequest) (response *ModifyCustomDnsHostResponse, err error) {
    if request == nil {
        request = NewModifyCustomDnsHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomDnsHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomDnsHostResponse()
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

func NewModifyIntlCustomDnsHostRequest() (request *ModifyIntlCustomDnsHostRequest) {
    request = &ModifyIntlCustomDnsHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "ModifyIntlCustomDnsHost")
    
    
    return
}

func NewModifyIntlCustomDnsHostResponse() (response *ModifyIntlCustomDnsHostResponse) {
    response = &ModifyIntlCustomDnsHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIntlCustomDnsHost
// 国际站-修改DNS Host
//
// 可能返回的错误码:
//  INTERNALERROR_DNSHOSTIPCHECKERR = "InternalError.DNSHostIPCheckErr"
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_MODIFYDNSHOSTERR = "InternalError.ModifyDNSHostErr"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) ModifyIntlCustomDnsHost(request *ModifyIntlCustomDnsHostRequest) (response *ModifyIntlCustomDnsHostResponse, err error) {
    return c.ModifyIntlCustomDnsHostWithContext(context.Background(), request)
}

// ModifyIntlCustomDnsHost
// 国际站-修改DNS Host
//
// 可能返回的错误码:
//  INTERNALERROR_DNSHOSTIPCHECKERR = "InternalError.DNSHostIPCheckErr"
//  INTERNALERROR_DESCRIBEDOMAININFOERR = "InternalError.DescribeDomainInfoErr"
//  INTERNALERROR_FORBIDDENREQUEST = "InternalError.ForbiddenRequest"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_METHODNOTMATCH = "InternalError.MethodNotMatch"
//  INTERNALERROR_MODIFYDNSHOSTERR = "InternalError.ModifyDNSHostErr"
//  INTERNALERROR_NEEDLOGIN = "InternalError.NeedLogin"
//  INTERNALERROR_READBODYERROR = "InternalError.ReadBodyError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  MISSINGPARAMETER_ACTIONNOTFOUND = "MissingParameter.ActionNotFound"
func (c *Client) ModifyIntlCustomDnsHostWithContext(ctx context.Context, request *ModifyIntlCustomDnsHostRequest) (response *ModifyIntlCustomDnsHostResponse, err error) {
    if request == nil {
        request = NewModifyIntlCustomDnsHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIntlCustomDnsHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIntlCustomDnsHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTemplateRequest() (request *ModifyTemplateRequest) {
    request = &ModifyTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "ModifyTemplate")
    
    
    return
}

func NewModifyTemplateResponse() (response *ModifyTemplateResponse) {
    response = &ModifyTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTemplate
// 修改模板信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETEMPLATEFAILED = "FailedOperation.CreateTemplateFailed"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  FAILEDOPERATION_PROHIBITPHONEEMAIL = "FailedOperation.ProhibitPhoneEmail"
//  FAILEDOPERATION_TEMPLATECANNOTMODIFY = "FailedOperation.TemplateCanNotModify"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) ModifyTemplate(request *ModifyTemplateRequest) (response *ModifyTemplateResponse, err error) {
    return c.ModifyTemplateWithContext(context.Background(), request)
}

// ModifyTemplate
// 修改模板信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATETEMPLATEFAILED = "FailedOperation.CreateTemplateFailed"
//  FAILEDOPERATION_DESCRIBETEMPLATEFAILED = "FailedOperation.DescribeTemplateFailed"
//  FAILEDOPERATION_PROHIBITPHONEEMAIL = "FailedOperation.ProhibitPhoneEmail"
//  FAILEDOPERATION_TEMPLATECANNOTMODIFY = "FailedOperation.TemplateCanNotModify"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TEMPLATENOTFOUND = "ResourceNotFound.TemplateNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
func (c *Client) ModifyTemplateWithContext(ctx context.Context, request *ModifyTemplateRequest) (response *ModifyTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTemplateResponse()
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

func NewReservedPreDomainsRequest() (request *ReservedPreDomainsRequest) {
    request = &ReservedPreDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "ReservedPreDomains")
    
    
    return
}

func NewReservedPreDomainsResponse() (response *ReservedPreDomainsResponse) {
    response = &ReservedPreDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReservedPreDomains
// 用于合作商对预释放域名进行预留。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) ReservedPreDomains(request *ReservedPreDomainsRequest) (response *ReservedPreDomainsResponse, err error) {
    return c.ReservedPreDomainsWithContext(context.Background(), request)
}

// ReservedPreDomains
// 用于合作商对预释放域名进行预留。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UINNOTWHITELISTERR = "FailedOperation.UinNotWhiteListErr"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
func (c *Client) ReservedPreDomainsWithContext(ctx context.Context, request *ReservedPreDomainsRequest) (response *ReservedPreDomainsResponse, err error) {
    if request == nil {
        request = NewReservedPreDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReservedPreDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewReservedPreDomainsResponse()
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
//  FAILEDOPERATION = "FailedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
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

func NewSyncCustomDnsHostRequest() (request *SyncCustomDnsHostRequest) {
    request = &SyncCustomDnsHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("domain", APIVersion, "SyncCustomDnsHost")
    
    
    return
}

func NewSyncCustomDnsHostResponse() (response *SyncCustomDnsHostResponse) {
    response = &SyncCustomDnsHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncCustomDnsHost
// 同步自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) SyncCustomDnsHost(request *SyncCustomDnsHostRequest) (response *SyncCustomDnsHostResponse, err error) {
    return c.SyncCustomDnsHostWithContext(context.Background(), request)
}

// SyncCustomDnsHost
// 同步自定义DNS Host
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CUSTOMDNSNAMENOTFOUND = "InvalidParameter.CustomDnsNameNotFound"
//  INVALIDPARAMETER_CUSTOMDNSNOTALLOWED = "InvalidParameter.CustomDnsNotAllowed"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOMAINNOTFOUND = "ResourceNotFound.DomainNotFound"
//  UNSUPPORTEDOPERATION_ACCOUNTREALNAME = "UnsupportedOperation.AccountRealName"
//  UNSUPPORTEDOPERATION_DOMAINNOTVERIFIED = "UnsupportedOperation.DomainNotVerified"
func (c *Client) SyncCustomDnsHostWithContext(ctx context.Context, request *SyncCustomDnsHostRequest) (response *SyncCustomDnsHostResponse, err error) {
    if request == nil {
        request = NewSyncCustomDnsHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncCustomDnsHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncCustomDnsHostResponse()
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
//  FAILEDOPERATION_TRANSFERINDOMAINFAILED = "FailedOperation.TransferInDomainFailed"
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
//  FAILEDOPERATION_TRANSFERINDOMAINFAILED = "FailedOperation.TransferInDomainFailed"
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
