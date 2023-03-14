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


func NewAttachPluginRequest() (request *AttachPluginRequest) {
    request = &AttachPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "AttachPlugin")
    
    
    return
}

func NewAttachPluginResponse() (response *AttachPluginResponse) {
    response = &AttachPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachPlugin
// 绑定插件到API上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_DUPLICATEPLUGINCONFIG = "InvalidParameter.DuplicatePluginConfig"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_SERVICECOUNTFORPLUGINLIMITEXCEEDED = "LimitExceeded.ServiceCountForPluginLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_ATTACHPLUGIN = "UnsupportedOperation.AttachPlugin"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
func (c *Client) AttachPlugin(request *AttachPluginRequest) (response *AttachPluginResponse, err error) {
    return c.AttachPluginWithContext(context.Background(), request)
}

// AttachPlugin
// 绑定插件到API上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_DUPLICATEPLUGINCONFIG = "InvalidParameter.DuplicatePluginConfig"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_SERVICECOUNTFORPLUGINLIMITEXCEEDED = "LimitExceeded.ServiceCountForPluginLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_ATTACHPLUGIN = "UnsupportedOperation.AttachPlugin"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
func (c *Client) AttachPluginWithContext(ctx context.Context, request *AttachPluginRequest) (response *AttachPluginResponse, err error) {
    if request == nil {
        request = NewAttachPluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachPluginResponse()
    err = c.Send(request, response)
    return
}

func NewBindApiAppRequest() (request *BindApiAppRequest) {
    request = &BindApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindApiApp")
    
    
    return
}

func NewBindApiAppResponse() (response *BindApiAppResponse) {
    response = &BindApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindApiApp
// 本接口（BindApiApp）用于绑定应用到API。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) BindApiApp(request *BindApiAppRequest) (response *BindApiAppResponse, err error) {
    return c.BindApiAppWithContext(context.Background(), request)
}

// BindApiApp
// 本接口（BindApiApp）用于绑定应用到API。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) BindApiAppWithContext(ctx context.Context, request *BindApiAppRequest) (response *BindApiAppResponse, err error) {
    if request == nil {
        request = NewBindApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewBindEnvironmentRequest() (request *BindEnvironmentRequest) {
    request = &BindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindEnvironment")
    
    
    return
}

func NewBindEnvironmentResponse() (response *BindEnvironmentResponse) {
    response = &BindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindEnvironment
// 本接口（BindEnvironment）用于绑定使用计划到服务或API。
//
// 用户在发布服务到某个环境中后，如果 API 需要鉴权，还需要绑定使用计划才能进行调用，此接口用户将使用计划绑定到特定环境。
//
// 目前支持绑定使用计划到API，但是同一个服务不能同时存在绑定到服务的使用计划和绑定到API的使用计划，所以对已经绑定过服务级别使用计划的环境，请先使用 服务级别使用计划降级 接口进行降级操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) BindEnvironment(request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    return c.BindEnvironmentWithContext(context.Background(), request)
}

// BindEnvironment
// 本接口（BindEnvironment）用于绑定使用计划到服务或API。
//
// 用户在发布服务到某个环境中后，如果 API 需要鉴权，还需要绑定使用计划才能进行调用，此接口用户将使用计划绑定到特定环境。
//
// 目前支持绑定使用计划到API，但是同一个服务不能同时存在绑定到服务的使用计划和绑定到API的使用计划，所以对已经绑定过服务级别使用计划的环境，请先使用 服务级别使用计划降级 接口进行降级操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) BindEnvironmentWithContext(ctx context.Context, request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    if request == nil {
        request = NewBindEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewBindIPStrategyRequest() (request *BindIPStrategyRequest) {
    request = &BindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindIPStrategy")
    
    
    return
}

func NewBindIPStrategyResponse() (response *BindIPStrategyResponse) {
    response = &BindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindIPStrategy
// 本接口（BindIPStrategy）用于API绑定IP策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) BindIPStrategy(request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
    return c.BindIPStrategyWithContext(context.Background(), request)
}

// BindIPStrategy
// 本接口（BindIPStrategy）用于API绑定IP策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) BindIPStrategyWithContext(ctx context.Context, request *BindIPStrategyRequest) (response *BindIPStrategyResponse, err error) {
    if request == nil {
        request = NewBindIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewBindSecretIdsRequest() (request *BindSecretIdsRequest) {
    request = &BindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSecretIds")
    
    
    return
}

func NewBindSecretIdsResponse() (response *BindSecretIdsResponse) {
    response = &BindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindSecretIds
// 本接口（BindSecretIds）用于为使用计划绑定密钥。
//
// 将密钥绑定到某个使用计划，并将此使用计划绑定到某个服务发布的环境上，调用者方可使用此密钥调用这个服务中的 API，可使用本接口为使用计划绑定密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDACCESSKEYIDS = "InvalidParameterValue.InvalidAccessKeyIds"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ACCESSKEYCOUNTINUSAGEPLANLIMITEXCEEDED = "LimitExceeded.AccessKeyCountInUsagePlanLimitExceeded"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_ALREADYBINDUSAGEPLAN = "UnsupportedOperation.AlreadyBindUsagePlan"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDAPIKEY = "UnsupportedOperation.UnsupportedBindApiKey"
func (c *Client) BindSecretIds(request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    return c.BindSecretIdsWithContext(context.Background(), request)
}

// BindSecretIds
// 本接口（BindSecretIds）用于为使用计划绑定密钥。
//
// 将密钥绑定到某个使用计划，并将此使用计划绑定到某个服务发布的环境上，调用者方可使用此密钥调用这个服务中的 API，可使用本接口为使用计划绑定密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDACCESSKEYIDS = "InvalidParameterValue.InvalidAccessKeyIds"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_ACCESSKEYCOUNTINUSAGEPLANLIMITEXCEEDED = "LimitExceeded.AccessKeyCountInUsagePlanLimitExceeded"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_ALREADYBINDUSAGEPLAN = "UnsupportedOperation.AlreadyBindUsagePlan"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDAPIKEY = "UnsupportedOperation.UnsupportedBindApiKey"
func (c *Client) BindSecretIdsWithContext(ctx context.Context, request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    if request == nil {
        request = NewBindSecretIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSecretIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewBindSubDomainRequest() (request *BindSubDomainRequest) {
    request = &BindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BindSubDomain")
    
    
    return
}

func NewBindSubDomainResponse() (response *BindSubDomainResponse) {
    response = &BindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindSubDomain
// 本接口（BindSubDomain）用于绑定自定义域名到服务。
//
// API 网关中每个服务都会提供一个默认的域名供用户调用，但当用户想使用自己的已有域名时，也可以将自定义域名绑定到此服务，在做好备案、与默认域名的 CNAME 后，可直接调用自定义域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEIDENTERPRISEWAITSUBMIT = "FailedOperation.CertificateIdEnterpriseWaitSubmit"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_CERTIFICATEIDEXPIRED = "FailedOperation.CertificateIdExpired"
//  FAILEDOPERATION_CERTIFICATEIDINFOERROR = "FailedOperation.CertificateIdInfoError"
//  FAILEDOPERATION_CERTIFICATEIDUNDERVERIFY = "FailedOperation.CertificateIdUnderVerify"
//  FAILEDOPERATION_CERTIFICATEIDUNKNOWNERROR = "FailedOperation.CertificateIdUnknownError"
//  FAILEDOPERATION_CERTIFICATEIDVERIFYFAIL = "FailedOperation.CertificateIdVerifyFail"
//  FAILEDOPERATION_CERTIFICATEISNULL = "FailedOperation.CertificateIsNull"
//  FAILEDOPERATION_DEFINEMAPPINGENVIRONMENTERROR = "FailedOperation.DefineMappingEnvironmentError"
//  FAILEDOPERATION_DEFINEMAPPINGNOTNULL = "FailedOperation.DefineMappingNotNull"
//  FAILEDOPERATION_DEFINEMAPPINGPARAMREPEAT = "FailedOperation.DefineMappingParamRepeat"
//  FAILEDOPERATION_DEFINEMAPPINGPATHERROR = "FailedOperation.DefineMappingPathError"
//  FAILEDOPERATION_DOMAINALREADYBINDOTHERSERVICE = "FailedOperation.DomainAlreadyBindOtherService"
//  FAILEDOPERATION_DOMAINALREADYBINDSERVICE = "FailedOperation.DomainAlreadyBindService"
//  FAILEDOPERATION_DOMAININBLACKLIST = "FailedOperation.DomainInBlackList"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_DOMAINRESOLVEERROR = "FailedOperation.DomainResolveError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_ISDEFAULTMAPPING = "FailedOperation.IsDefaultMapping"
//  FAILEDOPERATION_NETSUBDOMAINERROR = "FailedOperation.NetSubDomainError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDPROCOTOL = "InvalidParameterValue.InvalidProcotol"
//  LIMITEXCEEDED_EXCEEDEDDEFINEMAPPINGLIMIT = "LimitExceeded.ExceededDefineMappingLimit"
//  LIMITEXCEEDED_EXCEEDEDDOMAINLIMIT = "LimitExceeded.ExceededDomainLimit"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) BindSubDomain(request *BindSubDomainRequest) (response *BindSubDomainResponse, err error) {
    return c.BindSubDomainWithContext(context.Background(), request)
}

// BindSubDomain
// 本接口（BindSubDomain）用于绑定自定义域名到服务。
//
// API 网关中每个服务都会提供一个默认的域名供用户调用，但当用户想使用自己的已有域名时，也可以将自定义域名绑定到此服务，在做好备案、与默认域名的 CNAME 后，可直接调用自定义域名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEIDENTERPRISEWAITSUBMIT = "FailedOperation.CertificateIdEnterpriseWaitSubmit"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_CERTIFICATEIDEXPIRED = "FailedOperation.CertificateIdExpired"
//  FAILEDOPERATION_CERTIFICATEIDINFOERROR = "FailedOperation.CertificateIdInfoError"
//  FAILEDOPERATION_CERTIFICATEIDUNDERVERIFY = "FailedOperation.CertificateIdUnderVerify"
//  FAILEDOPERATION_CERTIFICATEIDUNKNOWNERROR = "FailedOperation.CertificateIdUnknownError"
//  FAILEDOPERATION_CERTIFICATEIDVERIFYFAIL = "FailedOperation.CertificateIdVerifyFail"
//  FAILEDOPERATION_CERTIFICATEISNULL = "FailedOperation.CertificateIsNull"
//  FAILEDOPERATION_DEFINEMAPPINGENVIRONMENTERROR = "FailedOperation.DefineMappingEnvironmentError"
//  FAILEDOPERATION_DEFINEMAPPINGNOTNULL = "FailedOperation.DefineMappingNotNull"
//  FAILEDOPERATION_DEFINEMAPPINGPARAMREPEAT = "FailedOperation.DefineMappingParamRepeat"
//  FAILEDOPERATION_DEFINEMAPPINGPATHERROR = "FailedOperation.DefineMappingPathError"
//  FAILEDOPERATION_DOMAINALREADYBINDOTHERSERVICE = "FailedOperation.DomainAlreadyBindOtherService"
//  FAILEDOPERATION_DOMAINALREADYBINDSERVICE = "FailedOperation.DomainAlreadyBindService"
//  FAILEDOPERATION_DOMAININBLACKLIST = "FailedOperation.DomainInBlackList"
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_DOMAINRESOLVEERROR = "FailedOperation.DomainResolveError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_ISDEFAULTMAPPING = "FailedOperation.IsDefaultMapping"
//  FAILEDOPERATION_NETSUBDOMAINERROR = "FailedOperation.NetSubDomainError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDPROCOTOL = "InvalidParameterValue.InvalidProcotol"
//  LIMITEXCEEDED_EXCEEDEDDEFINEMAPPINGLIMIT = "LimitExceeded.ExceededDefineMappingLimit"
//  LIMITEXCEEDED_EXCEEDEDDOMAINLIMIT = "LimitExceeded.ExceededDomainLimit"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) BindSubDomainWithContext(ctx context.Context, request *BindSubDomainRequest) (response *BindSubDomainResponse, err error) {
    if request == nil {
        request = NewBindSubDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSubDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewBuildAPIDocRequest() (request *BuildAPIDocRequest) {
    request = &BuildAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "BuildAPIDoc")
    
    
    return
}

func NewBuildAPIDocResponse() (response *BuildAPIDocResponse) {
    response = &BuildAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BuildAPIDoc
// 构建 API 文档
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) BuildAPIDoc(request *BuildAPIDocRequest) (response *BuildAPIDocResponse, err error) {
    return c.BuildAPIDocWithContext(context.Background(), request)
}

// BuildAPIDoc
// 构建 API 文档
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) BuildAPIDocWithContext(ctx context.Context, request *BuildAPIDocRequest) (response *BuildAPIDocResponse, err error) {
    if request == nil {
        request = NewBuildAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BuildAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewBuildAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAPIDocRequest() (request *CreateAPIDocRequest) {
    request = &CreateAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateAPIDoc")
    
    
    return
}

func NewCreateAPIDocResponse() (response *CreateAPIDocResponse) {
    response = &CreateAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAPIDoc
// 创建 API 文档
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_APIBINDENVIRONMENT = "FailedOperation.ApiBindEnvironment"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_APIDOCLIMITEXCEEDED = "LimitExceeded.APIDocLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAPIDoc(request *CreateAPIDocRequest) (response *CreateAPIDocResponse, err error) {
    return c.CreateAPIDocWithContext(context.Background(), request)
}

// CreateAPIDoc
// 创建 API 文档
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_APIBINDENVIRONMENT = "FailedOperation.ApiBindEnvironment"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_APIDOCLIMITEXCEEDED = "LimitExceeded.APIDocLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAPIDocWithContext(ctx context.Context, request *CreateAPIDocRequest) (response *CreateAPIDocResponse, err error) {
    if request == nil {
        request = NewCreateAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRequest() (request *CreateApiRequest) {
    request = &CreateApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApi")
    
    
    return
}

func NewCreateApiResponse() (response *CreateApiResponse) {
    response = &CreateApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApi
// 本接口（CreateApi）用于创建 API 接口，创建 API 前，用户需要先创建服务，每个 API 都有自己归属的服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SCFERROR = "FailedOperation.ScfError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CLBEXCEPTION = "InternalError.ClbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPITYPE = "InvalidParameterValue.InvalidApiType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEMOCKRETURNMESSAGE = "InvalidParameterValue.InvalidServiceMockReturnMessage"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAMETERS = "InvalidParameterValue.InvalidServiceParameters"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_BASICSERVICENOMOREAPI = "UnsupportedOperation.BasicServiceNoMoreApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDNETTYPE = "UnsupportedOperation.UnsupportedNetType"
func (c *Client) CreateApi(request *CreateApiRequest) (response *CreateApiResponse, err error) {
    return c.CreateApiWithContext(context.Background(), request)
}

// CreateApi
// 本接口（CreateApi）用于创建 API 接口，创建 API 前，用户需要先创建服务，每个 API 都有自己归属的服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SCFERROR = "FailedOperation.ScfError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CLBEXCEPTION = "InternalError.ClbException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPITYPE = "InvalidParameterValue.InvalidApiType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEMOCKRETURNMESSAGE = "InvalidParameterValue.InvalidServiceMockReturnMessage"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAMETERS = "InvalidParameterValue.InvalidServiceParameters"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_BASICSERVICENOMOREAPI = "UnsupportedOperation.BasicServiceNoMoreApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDNETTYPE = "UnsupportedOperation.UnsupportedNetType"
func (c *Client) CreateApiWithContext(ctx context.Context, request *CreateApiRequest) (response *CreateApiResponse, err error) {
    if request == nil {
        request = NewCreateApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiAppRequest() (request *CreateApiAppRequest) {
    request = &CreateApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiApp")
    
    
    return
}

func NewCreateApiAppResponse() (response *CreateApiAppResponse) {
    response = &CreateApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApiApp
// 本接口（CreateApiApp）用于创建应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
func (c *Client) CreateApiApp(request *CreateApiAppRequest) (response *CreateApiAppResponse, err error) {
    return c.CreateApiAppWithContext(context.Background(), request)
}

// CreateApiApp
// 本接口（CreateApiApp）用于创建应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
func (c *Client) CreateApiAppWithContext(ctx context.Context, request *CreateApiAppRequest) (response *CreateApiAppResponse, err error) {
    if request == nil {
        request = NewCreateApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
    request = &CreateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateApiKey")
    
    
    return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
    response = &CreateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApiKey
// 本接口（CreateApiKey）用于创建一对新的 API 密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    return c.CreateApiKeyWithContext(context.Background(), request)
}

// CreateApiKey
// 本接口（CreateApiKey）用于创建一对新的 API 密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
func (c *Client) CreateApiKeyWithContext(ctx context.Context, request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIPStrategyRequest() (request *CreateIPStrategyRequest) {
    request = &CreateIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateIPStrategy")
    
    
    return
}

func NewCreateIPStrategyResponse() (response *CreateIPStrategyResponse) {
    response = &CreateIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIPStrategy
// 本接口（CreateIPStrategy）用于创建服务IP策略。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateIPStrategy(request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    return c.CreateIPStrategyWithContext(context.Background(), request)
}

// CreateIPStrategy
// 本接口（CreateIPStrategy）用于创建服务IP策略。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateIPStrategyWithContext(ctx context.Context, request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    if request == nil {
        request = NewCreateIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePluginRequest() (request *CreatePluginRequest) {
    request = &CreatePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreatePlugin")
    
    
    return
}

func NewCreatePluginResponse() (response *CreatePluginResponse) {
    response = &CreatePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePlugin
// 创建API网关插件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreatePlugin(request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    return c.CreatePluginWithContext(context.Background(), request)
}

// CreatePlugin
// 创建API网关插件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreatePluginWithContext(ctx context.Context, request *CreatePluginRequest) (response *CreatePluginResponse, err error) {
    if request == nil {
        request = NewCreatePluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePluginResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
    request = &CreateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateService")
    
    
    return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
    response = &CreateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateService
// 本接口（CreateService）用于创建服务。
//
// API 网关使用的最大单元为服务，每个服务中可创建多个 API 接口。每个服务有一个默认域名供客户调用，用户也可绑定自定义域名到此服务中。 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAUTHEXCEPTION = "InternalError.CauthException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  LIMITEXCEEDED_SERVICECOUNTLIMITEXCEEDED = "LimitExceeded.ServiceCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    return c.CreateServiceWithContext(context.Background(), request)
}

// CreateService
// 本接口（CreateService）用于创建服务。
//
// API 网关使用的最大单元为服务，每个服务中可创建多个 API 接口。每个服务有一个默认域名供客户调用，用户也可绑定自定义域名到此服务中。 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_CAUTHEXCEPTION = "InternalError.CauthException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  LIMITEXCEEDED_SERVICECOUNTLIMITEXCEEDED = "LimitExceeded.ServiceCountLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"
//  UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateServiceWithContext(ctx context.Context, request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUpstreamRequest() (request *CreateUpstreamRequest) {
    request = &CreateUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateUpstream")
    
    
    return
}

func NewCreateUpstreamResponse() (response *CreateUpstreamResponse) {
    response = &CreateUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUpstream
// 用于创建创建后端通道
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUpstream(request *CreateUpstreamRequest) (response *CreateUpstreamResponse, err error) {
    return c.CreateUpstreamWithContext(context.Background(), request)
}

// CreateUpstream
// 用于创建创建后端通道
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUpstreamWithContext(ctx context.Context, request *CreateUpstreamRequest) (response *CreateUpstreamResponse, err error) {
    if request == nil {
        request = NewCreateUpstreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUsagePlanRequest() (request *CreateUsagePlanRequest) {
    request = &CreateUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "CreateUsagePlan")
    
    
    return
}

func NewCreateUsagePlanResponse() (response *CreateUsagePlanResponse) {
    response = &CreateUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUsagePlan
// 本接口（CreateUsagePlan）用于创建使用计划。
//
// 用户在使用 API 网关时，需要创建使用计划并将其绑定到服务的环境中使用。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDMAXREQUESTNUM = "InvalidParameterValue.InvalidMaxRequestNum"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_USAGEPLANLIMITEXCEEDED = "LimitExceeded.UsagePlanLimitExceeded"
func (c *Client) CreateUsagePlan(request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    return c.CreateUsagePlanWithContext(context.Background(), request)
}

// CreateUsagePlan
// 本接口（CreateUsagePlan）用于创建使用计划。
//
// 用户在使用 API 网关时，需要创建使用计划并将其绑定到服务的环境中使用。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDMAXREQUESTNUM = "InvalidParameterValue.InvalidMaxRequestNum"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_USAGEPLANLIMITEXCEEDED = "LimitExceeded.UsagePlanLimitExceeded"
func (c *Client) CreateUsagePlanWithContext(ctx context.Context, request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    if request == nil {
        request = NewCreateUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAPIDocRequest() (request *DeleteAPIDocRequest) {
    request = &DeleteAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteAPIDoc")
    
    
    return
}

func NewDeleteAPIDocResponse() (response *DeleteAPIDocResponse) {
    response = &DeleteAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAPIDoc
// 删除 API 文档
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAPIDoc(request *DeleteAPIDocRequest) (response *DeleteAPIDocResponse, err error) {
    return c.DeleteAPIDocWithContext(context.Background(), request)
}

// DeleteAPIDoc
// 删除 API 文档
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAPIDocWithContext(ctx context.Context, request *DeleteAPIDocRequest) (response *DeleteAPIDocResponse, err error) {
    if request == nil {
        request = NewDeleteAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiRequest() (request *DeleteApiRequest) {
    request = &DeleteApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApi")
    
    
    return
}

func NewDeleteApiResponse() (response *DeleteApiResponse) {
    response = &DeleteApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApi
// 本接口（DeleteApi）用于删除已经创建的API。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    return c.DeleteApiWithContext(context.Background(), request)
}

// DeleteApi
// 本接口（DeleteApi）用于删除已经创建的API。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApiWithContext(ctx context.Context, request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    if request == nil {
        request = NewDeleteApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiAppRequest() (request *DeleteApiAppRequest) {
    request = &DeleteApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiApp")
    
    
    return
}

func NewDeleteApiAppResponse() (response *DeleteApiAppResponse) {
    response = &DeleteApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApiApp
// 本接口（DeleteApiApp）用于删除已经创建的应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApiApp(request *DeleteApiAppRequest) (response *DeleteApiAppResponse, err error) {
    return c.DeleteApiAppWithContext(context.Background(), request)
}

// DeleteApiApp
// 本接口（DeleteApiApp）用于删除已经创建的应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) DeleteApiAppWithContext(ctx context.Context, request *DeleteApiAppRequest) (response *DeleteApiAppResponse, err error) {
    if request == nil {
        request = NewDeleteApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
    request = &DeleteApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteApiKey")
    
    
    return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
    response = &DeleteApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApiKey
// 本接口（DeleteApiKey）用于删除一对 API 密钥。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_RESOURCEISINUSE = "UnsupportedOperation.ResourceIsInUse"
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    return c.DeleteApiKeyWithContext(context.Background(), request)
}

// DeleteApiKey
// 本接口（DeleteApiKey）用于删除一对 API 密钥。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
//  UNSUPPORTEDOPERATION_RESOURCEISINUSE = "UnsupportedOperation.ResourceIsInUse"
func (c *Client) DeleteApiKeyWithContext(ctx context.Context, request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIPStrategyRequest() (request *DeleteIPStrategyRequest) {
    request = &DeleteIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteIPStrategy")
    
    
    return
}

func NewDeleteIPStrategyResponse() (response *DeleteIPStrategyResponse) {
    response = &DeleteIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIPStrategy
// 本接口（DeleteIPStrategy）用于删除服务IP策略。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DeleteIPStrategy(request *DeleteIPStrategyRequest) (response *DeleteIPStrategyResponse, err error) {
    return c.DeleteIPStrategyWithContext(context.Background(), request)
}

// DeleteIPStrategy
// 本接口（DeleteIPStrategy）用于删除服务IP策略。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DeleteIPStrategyWithContext(ctx context.Context, request *DeleteIPStrategyRequest) (response *DeleteIPStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePluginRequest() (request *DeletePluginRequest) {
    request = &DeletePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeletePlugin")
    
    
    return
}

func NewDeletePluginResponse() (response *DeletePluginResponse) {
    response = &DeletePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePlugin
// 删除API网关插件
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DeletePlugin(request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    return c.DeletePluginWithContext(context.Background(), request)
}

// DeletePlugin
// 删除API网关插件
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DeletePluginWithContext(ctx context.Context, request *DeletePluginRequest) (response *DeletePluginResponse, err error) {
    if request == nil {
        request = NewDeletePluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
    request = &DeleteServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteService")
    
    
    return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
    response = &DeleteServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteService
// 本接口（DeleteService）用于删除 API 网关中某个服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETESERVICE = "UnsupportedOperation.UnsupportedDeleteService"
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    return c.DeleteServiceWithContext(context.Background(), request)
}

// DeleteService
// 本接口（DeleteService）用于删除 API 网关中某个服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETESERVICE = "UnsupportedOperation.UnsupportedDeleteService"
func (c *Client) DeleteServiceWithContext(ctx context.Context, request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceSubDomainMappingRequest() (request *DeleteServiceSubDomainMappingRequest) {
    request = &DeleteServiceSubDomainMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteServiceSubDomainMapping")
    
    
    return
}

func NewDeleteServiceSubDomainMappingResponse() (response *DeleteServiceSubDomainMappingResponse) {
    response = &DeleteServiceSubDomainMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteServiceSubDomainMapping
// 本接口（DeleteServiceSubDomainMapping）用于删除服务中某个环境的自定义域名映射。
//
// 当用户使用自定义域名，并使用了自定义映射时，可使用此接口。但需注意，若删除了所有环境的映射时，调用此 API 均会返回失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEPATHMAPPINGSETERROR = "FailedOperation.DeletePathMappingSetError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteServiceSubDomainMapping(request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    return c.DeleteServiceSubDomainMappingWithContext(context.Background(), request)
}

// DeleteServiceSubDomainMapping
// 本接口（DeleteServiceSubDomainMapping）用于删除服务中某个环境的自定义域名映射。
//
// 当用户使用自定义域名，并使用了自定义映射时，可使用此接口。但需注意，若删除了所有环境的映射时，调用此 API 均会返回失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEPATHMAPPINGSETERROR = "FailedOperation.DeletePathMappingSetError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteServiceSubDomainMappingWithContext(ctx context.Context, request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    if request == nil {
        request = NewDeleteServiceSubDomainMappingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceSubDomainMapping require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceSubDomainMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUpstreamRequest() (request *DeleteUpstreamRequest) {
    request = &DeleteUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUpstream")
    
    
    return
}

func NewDeleteUpstreamResponse() (response *DeleteUpstreamResponse) {
    response = &DeleteUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUpstream
// 删除后端通道，需要注意有API绑定时，不允许删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEUPSTREAM = "UnsupportedOperation.UnsupportedDeleteUpstream"
func (c *Client) DeleteUpstream(request *DeleteUpstreamRequest) (response *DeleteUpstreamResponse, err error) {
    return c.DeleteUpstreamWithContext(context.Background(), request)
}

// DeleteUpstream
// 删除后端通道，需要注意有API绑定时，不允许删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEUPSTREAM = "UnsupportedOperation.UnsupportedDeleteUpstream"
func (c *Client) DeleteUpstreamWithContext(ctx context.Context, request *DeleteUpstreamRequest) (response *DeleteUpstreamResponse, err error) {
    if request == nil {
        request = NewDeleteUpstreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsagePlanRequest() (request *DeleteUsagePlanRequest) {
    request = &DeleteUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DeleteUsagePlan")
    
    
    return
}

func NewDeleteUsagePlanResponse() (response *DeleteUsagePlanResponse) {
    response = &DeleteUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUsagePlan
// 本接口（DeleteUsagePlan）用于删除使用计划。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_USAGEPLANINUSE = "UnsupportedOperation.UsagePlanInUse"
func (c *Client) DeleteUsagePlan(request *DeleteUsagePlanRequest) (response *DeleteUsagePlanResponse, err error) {
    return c.DeleteUsagePlanWithContext(context.Background(), request)
}

// DeleteUsagePlan
// 本接口（DeleteUsagePlan）用于删除使用计划。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_USAGEPLANINUSE = "UnsupportedOperation.UsagePlanInUse"
func (c *Client) DeleteUsagePlanWithContext(ctx context.Context, request *DeleteUsagePlanRequest) (response *DeleteUsagePlanResponse, err error) {
    if request == nil {
        request = NewDeleteUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDemoteServiceUsagePlanRequest() (request *DemoteServiceUsagePlanRequest) {
    request = &DemoteServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DemoteServiceUsagePlan")
    
    
    return
}

func NewDemoteServiceUsagePlanResponse() (response *DemoteServiceUsagePlanResponse) {
    response = &DemoteServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DemoteServiceUsagePlan
// 本接口（DemoteServiceUsagePlan）用于将某个服务在某个环境的使用计划，降级到API上。
//
// 如果服务内没有API不允许进行此操作。
//
// 如果当前环境没有发布，不允许进行此操作。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_NOUSAGEPLANENV = "UnsupportedOperation.NoUsagePlanEnv"
func (c *Client) DemoteServiceUsagePlan(request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    return c.DemoteServiceUsagePlanWithContext(context.Background(), request)
}

// DemoteServiceUsagePlan
// 本接口（DemoteServiceUsagePlan）用于将某个服务在某个环境的使用计划，降级到API上。
//
// 如果服务内没有API不允许进行此操作。
//
// 如果当前环境没有发布，不允许进行此操作。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_NOUSAGEPLANENV = "UnsupportedOperation.NoUsagePlanEnv"
func (c *Client) DemoteServiceUsagePlanWithContext(ctx context.Context, request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDemoteServiceUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DemoteServiceUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDemoteServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAPIDocDetailRequest() (request *DescribeAPIDocDetailRequest) {
    request = &DescribeAPIDocDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAPIDocDetail")
    
    
    return
}

func NewDescribeAPIDocDetailResponse() (response *DescribeAPIDocDetailResponse) {
    response = &DescribeAPIDocDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAPIDocDetail
// 查询 API 文档详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocDetail(request *DescribeAPIDocDetailRequest) (response *DescribeAPIDocDetailResponse, err error) {
    return c.DescribeAPIDocDetailWithContext(context.Background(), request)
}

// DescribeAPIDocDetail
// 查询 API 文档详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocDetailWithContext(ctx context.Context, request *DescribeAPIDocDetailRequest) (response *DescribeAPIDocDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAPIDocDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIDocDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIDocDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAPIDocsRequest() (request *DescribeAPIDocsRequest) {
    request = &DescribeAPIDocsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAPIDocs")
    
    
    return
}

func NewDescribeAPIDocsResponse() (response *DescribeAPIDocsResponse) {
    response = &DescribeAPIDocsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAPIDocs
// 查询 API 文档列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocs(request *DescribeAPIDocsRequest) (response *DescribeAPIDocsResponse, err error) {
    return c.DescribeAPIDocsWithContext(context.Background(), request)
}

// DescribeAPIDocs
// 查询 API 文档列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocsWithContext(ctx context.Context, request *DescribeAPIDocsRequest) (response *DescribeAPIDocsResponse, err error) {
    if request == nil {
        request = NewDescribeAPIDocsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIDocs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIDocsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllPluginApisRequest() (request *DescribeAllPluginApisRequest) {
    request = &DescribeAllPluginApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeAllPluginApis")
    
    
    return
}

func NewDescribeAllPluginApisResponse() (response *DescribeAllPluginApisResponse) {
    response = &DescribeAllPluginApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllPluginApis
// 展示插件相关的API列表，包括已绑定的和未绑定的API信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeAllPluginApis(request *DescribeAllPluginApisRequest) (response *DescribeAllPluginApisResponse, err error) {
    return c.DescribeAllPluginApisWithContext(context.Background(), request)
}

// DescribeAllPluginApis
// 展示插件相关的API列表，包括已绑定的和未绑定的API信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeAllPluginApisWithContext(ctx context.Context, request *DescribeAllPluginApisRequest) (response *DescribeAllPluginApisResponse, err error) {
    if request == nil {
        request = NewDescribeAllPluginApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllPluginApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllPluginApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRequest() (request *DescribeApiRequest) {
    request = &DescribeApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApi")
    
    
    return
}

func NewDescribeApiResponse() (response *DescribeApiResponse) {
    response = &DescribeApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApi
// 本接口（DescribeApi）用于查询用户 API 网关的 API 接口的详细信息。​
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    return c.DescribeApiWithContext(context.Background(), request)
}

// DescribeApi
// 本接口（DescribeApi）用于查询用户 API 网关的 API 接口的详细信息。​
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApiWithContext(ctx context.Context, request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    if request == nil {
        request = NewDescribeApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiAppRequest() (request *DescribeApiAppRequest) {
    request = &DescribeApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiApp")
    
    
    return
}

func NewDescribeApiAppResponse() (response *DescribeApiAppResponse) {
    response = &DescribeApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiApp
// 本接口（DescribeApiApp）用于根据应用ID搜索应用。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiApp(request *DescribeApiAppRequest) (response *DescribeApiAppResponse, err error) {
    return c.DescribeApiAppWithContext(context.Background(), request)
}

// DescribeApiApp
// 本接口（DescribeApiApp）用于根据应用ID搜索应用。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiAppWithContext(ctx context.Context, request *DescribeApiAppRequest) (response *DescribeApiAppResponse, err error) {
    if request == nil {
        request = NewDescribeApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiAppBindApisStatusRequest() (request *DescribeApiAppBindApisStatusRequest) {
    request = &DescribeApiAppBindApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiAppBindApisStatus")
    
    
    return
}

func NewDescribeApiAppBindApisStatusResponse() (response *DescribeApiAppBindApisStatusResponse) {
    response = &DescribeApiAppBindApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiAppBindApisStatus
// 本接口（DescribeApiAppBindApisStatus）查询应用绑定的Api列表。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeApiAppBindApisStatus(request *DescribeApiAppBindApisStatusRequest) (response *DescribeApiAppBindApisStatusResponse, err error) {
    return c.DescribeApiAppBindApisStatusWithContext(context.Background(), request)
}

// DescribeApiAppBindApisStatus
// 本接口（DescribeApiAppBindApisStatus）查询应用绑定的Api列表。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeApiAppBindApisStatusWithContext(ctx context.Context, request *DescribeApiAppBindApisStatusRequest) (response *DescribeApiAppBindApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiAppBindApisStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiAppBindApisStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiAppBindApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiAppsStatusRequest() (request *DescribeApiAppsStatusRequest) {
    request = &DescribeApiAppsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiAppsStatus")
    
    
    return
}

func NewDescribeApiAppsStatusResponse() (response *DescribeApiAppsStatusResponse) {
    response = &DescribeApiAppsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiAppsStatus
// 本接口（DescribeApiAppsStatus）查询应用列表。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApiAppsStatus(request *DescribeApiAppsStatusRequest) (response *DescribeApiAppsStatusResponse, err error) {
    return c.DescribeApiAppsStatusWithContext(context.Background(), request)
}

// DescribeApiAppsStatus
// 本接口（DescribeApiAppsStatus）查询应用列表。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApiAppsStatusWithContext(ctx context.Context, request *DescribeApiAppsStatusRequest) (response *DescribeApiAppsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiAppsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiAppsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiAppsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiBindApiAppsStatusRequest() (request *DescribeApiBindApiAppsStatusRequest) {
    request = &DescribeApiBindApiAppsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiBindApiAppsStatus")
    
    
    return
}

func NewDescribeApiBindApiAppsStatusResponse() (response *DescribeApiBindApiAppsStatusResponse) {
    response = &DescribeApiBindApiAppsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiBindApiAppsStatus
// 本接口（DescribeApiBindApiAppsStatus）查询Api绑定的应用列表。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiBindApiAppsStatus(request *DescribeApiBindApiAppsStatusRequest) (response *DescribeApiBindApiAppsStatusResponse, err error) {
    return c.DescribeApiBindApiAppsStatusWithContext(context.Background(), request)
}

// DescribeApiBindApiAppsStatus
// 本接口（DescribeApiBindApiAppsStatus）查询Api绑定的应用列表。
//
// 可能返回的错误码:
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeApiBindApiAppsStatusWithContext(ctx context.Context, request *DescribeApiBindApiAppsStatusRequest) (response *DescribeApiBindApiAppsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiBindApiAppsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiBindApiAppsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiBindApiAppsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiEnvironmentStrategyRequest() (request *DescribeApiEnvironmentStrategyRequest) {
    request = &DescribeApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiEnvironmentStrategy")
    
    
    return
}

func NewDescribeApiEnvironmentStrategyResponse() (response *DescribeApiEnvironmentStrategyResponse) {
    response = &DescribeApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiEnvironmentStrategy
// 本接口（DescribeApiEnvironmentStrategy）用于展示API绑定的限流策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiEnvironmentStrategy(request *DescribeApiEnvironmentStrategyRequest) (response *DescribeApiEnvironmentStrategyResponse, err error) {
    return c.DescribeApiEnvironmentStrategyWithContext(context.Background(), request)
}

// DescribeApiEnvironmentStrategy
// 本接口（DescribeApiEnvironmentStrategy）用于展示API绑定的限流策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiEnvironmentStrategyWithContext(ctx context.Context, request *DescribeApiEnvironmentStrategyRequest) (response *DescribeApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeApiEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiForApiAppRequest() (request *DescribeApiForApiAppRequest) {
    request = &DescribeApiForApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiForApiApp")
    
    
    return
}

func NewDescribeApiForApiAppResponse() (response *DescribeApiForApiAppResponse) {
    response = &DescribeApiForApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiForApiApp
// 本接口（DescribeApiForApiApp）用于应用使用者查询部署于 API 网关的 API 接口的详细信息。​
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiForApiApp(request *DescribeApiForApiAppRequest) (response *DescribeApiForApiAppResponse, err error) {
    return c.DescribeApiForApiAppWithContext(context.Background(), request)
}

// DescribeApiForApiApp
// 本接口（DescribeApiForApiApp）用于应用使用者查询部署于 API 网关的 API 接口的详细信息。​
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiForApiAppWithContext(ctx context.Context, request *DescribeApiForApiAppRequest) (response *DescribeApiForApiAppResponse, err error) {
    if request == nil {
        request = NewDescribeApiForApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiForApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiForApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyRequest() (request *DescribeApiKeyRequest) {
    request = &DescribeApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKey")
    
    
    return
}

func NewDescribeApiKeyResponse() (response *DescribeApiKeyResponse) {
    response = &DescribeApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiKey
// 本接口（DescribeApiKey）用于查询密钥详情。
//
// 用户在创建密钥后，可用此接口查询一个 API 密钥的详情，该接口会显示密钥 Key。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    return c.DescribeApiKeyWithContext(context.Background(), request)
}

// DescribeApiKey
// 本接口（DescribeApiKey）用于查询密钥详情。
//
// 用户在创建密钥后，可用此接口查询一个 API 密钥的详情，该接口会显示密钥 Key。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
func (c *Client) DescribeApiKeyWithContext(ctx context.Context, request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeysStatusRequest() (request *DescribeApiKeysStatusRequest) {
    request = &DescribeApiKeysStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiKeysStatus")
    
    
    return
}

func NewDescribeApiKeysStatusResponse() (response *DescribeApiKeysStatusResponse) {
    response = &DescribeApiKeysStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiKeysStatus
// 本接口（DescribeApiKeysStatus）用于查询密钥列表。
//
// 当用户创建了多个密钥对时，可使用本接口查询一个或多个 API 密钥信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeApiKeysStatus(request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
    return c.DescribeApiKeysStatusWithContext(context.Background(), request)
}

// DescribeApiKeysStatus
// 本接口（DescribeApiKeysStatus）用于查询密钥列表。
//
// 当用户创建了多个密钥对时，可使用本接口查询一个或多个 API 密钥信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeApiKeysStatusWithContext(ctx context.Context, request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeysStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKeysStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeysStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUsagePlanRequest() (request *DescribeApiUsagePlanRequest) {
    request = &DescribeApiUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApiUsagePlan")
    
    
    return
}

func NewDescribeApiUsagePlanResponse() (response *DescribeApiUsagePlanResponse) {
    response = &DescribeApiUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiUsagePlan
// 本接口（DescribeApiUsagePlan）用于查询服务中 API 使用计划详情。
//
// 服务若需要鉴权限流生效，则需要绑定使用计划到此服务中，本接口用于查询绑定到一个服务及其中 API 的所有使用计划。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiUsagePlan(request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    return c.DescribeApiUsagePlanWithContext(context.Background(), request)
}

// DescribeApiUsagePlan
// 本接口（DescribeApiUsagePlan）用于查询服务中 API 使用计划详情。
//
// 服务若需要鉴权限流生效，则需要绑定使用计划到此服务中，本接口用于查询绑定到一个服务及其中 API 的所有使用计划。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiUsagePlanWithContext(ctx context.Context, request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeApiUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApisStatusRequest() (request *DescribeApisStatusRequest) {
    request = &DescribeApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeApisStatus")
    
    
    return
}

func NewDescribeApisStatusResponse() (response *DescribeApisStatusResponse) {
    response = &DescribeApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApisStatus
// 本接口（DescribeApisStatus）用于查看一个服务下的某个 API 或所有 API 列表及其相关信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApisStatus(request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
    return c.DescribeApisStatusWithContext(context.Background(), request)
}

// DescribeApisStatus
// 本接口（DescribeApisStatus）用于查看一个服务下的某个 API 或所有 API 列表及其相关信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApisStatusWithContext(ctx context.Context, request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApisStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApisStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExclusiveInstanceDetailRequest() (request *DescribeExclusiveInstanceDetailRequest) {
    request = &DescribeExclusiveInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeExclusiveInstanceDetail")
    
    
    return
}

func NewDescribeExclusiveInstanceDetailResponse() (response *DescribeExclusiveInstanceDetailResponse) {
    response = &DescribeExclusiveInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExclusiveInstanceDetail
// 本接口（DescribeExclusiveInstanceDetail）用于查询独享实例详情信息。​
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSERROR = "FailedOperation.ClsError"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeExclusiveInstanceDetail(request *DescribeExclusiveInstanceDetailRequest) (response *DescribeExclusiveInstanceDetailResponse, err error) {
    return c.DescribeExclusiveInstanceDetailWithContext(context.Background(), request)
}

// DescribeExclusiveInstanceDetail
// 本接口（DescribeExclusiveInstanceDetail）用于查询独享实例详情信息。​
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSERROR = "FailedOperation.ClsError"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeExclusiveInstanceDetailWithContext(ctx context.Context, request *DescribeExclusiveInstanceDetailRequest) (response *DescribeExclusiveInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeExclusiveInstanceDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExclusiveInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExclusiveInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExclusiveInstancesRequest() (request *DescribeExclusiveInstancesRequest) {
    request = &DescribeExclusiveInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeExclusiveInstances")
    
    
    return
}

func NewDescribeExclusiveInstancesResponse() (response *DescribeExclusiveInstancesResponse) {
    response = &DescribeExclusiveInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExclusiveInstances
// 本接口（DescribeExclusiveInstances）用于查询独享实例列表信息。​
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSERROR = "FailedOperation.ClsError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeExclusiveInstances(request *DescribeExclusiveInstancesRequest) (response *DescribeExclusiveInstancesResponse, err error) {
    return c.DescribeExclusiveInstancesWithContext(context.Background(), request)
}

// DescribeExclusiveInstances
// 本接口（DescribeExclusiveInstances）用于查询独享实例列表信息。​
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSERROR = "FailedOperation.ClsError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeExclusiveInstancesWithContext(ctx context.Context, request *DescribeExclusiveInstancesRequest) (response *DescribeExclusiveInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeExclusiveInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExclusiveInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExclusiveInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExclusiveInstancesStatusRequest() (request *DescribeExclusiveInstancesStatusRequest) {
    request = &DescribeExclusiveInstancesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeExclusiveInstancesStatus")
    
    
    return
}

func NewDescribeExclusiveInstancesStatusResponse() (response *DescribeExclusiveInstancesStatusResponse) {
    response = &DescribeExclusiveInstancesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExclusiveInstancesStatus
// 查询专享实例列表（新）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSERROR = "FailedOperation.ClsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeExclusiveInstancesStatus(request *DescribeExclusiveInstancesStatusRequest) (response *DescribeExclusiveInstancesStatusResponse, err error) {
    return c.DescribeExclusiveInstancesStatusWithContext(context.Background(), request)
}

// DescribeExclusiveInstancesStatus
// 查询专享实例列表（新）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSERROR = "FailedOperation.ClsError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeExclusiveInstancesStatusWithContext(ctx context.Context, request *DescribeExclusiveInstancesStatusRequest) (response *DescribeExclusiveInstancesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeExclusiveInstancesStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExclusiveInstancesStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExclusiveInstancesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategyRequest() (request *DescribeIPStrategyRequest) {
    request = &DescribeIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategy")
    
    
    return
}

func NewDescribeIPStrategyResponse() (response *DescribeIPStrategyResponse) {
    response = &DescribeIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIPStrategy
// 本接口（DescribeIPStrategy）用于查询IP策略详情。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategy(request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
    return c.DescribeIPStrategyWithContext(context.Background(), request)
}

// DescribeIPStrategy
// 本接口（DescribeIPStrategy）用于查询IP策略详情。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategyWithContext(ctx context.Context, request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategyApisStatusRequest() (request *DescribeIPStrategyApisStatusRequest) {
    request = &DescribeIPStrategyApisStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategyApisStatus")
    
    
    return
}

func NewDescribeIPStrategyApisStatusResponse() (response *DescribeIPStrategyApisStatusResponse) {
    response = &DescribeIPStrategyApisStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIPStrategyApisStatus
// 本接口（DescribeIPStrategyApisStatus）用于查询IP策略可以绑定的API列表。即服务下所有API和该策略已绑定API的差集。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DescribeIPStrategyApisStatus(request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
    return c.DescribeIPStrategyApisStatusWithContext(context.Background(), request)
}

// DescribeIPStrategyApisStatus
// 本接口（DescribeIPStrategyApisStatus）用于查询IP策略可以绑定的API列表。即服务下所有API和该策略已绑定API的差集。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DescribeIPStrategyApisStatusWithContext(ctx context.Context, request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyApisStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStrategyApisStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStrategyApisStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPStrategysStatusRequest() (request *DescribeIPStrategysStatusRequest) {
    request = &DescribeIPStrategysStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeIPStrategysStatus")
    
    
    return
}

func NewDescribeIPStrategysStatusResponse() (response *DescribeIPStrategysStatusResponse) {
    response = &DescribeIPStrategysStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIPStrategysStatus
// 本接口（DescribeIPStrategysStatus）用于查询服务IP策略列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategysStatus(request *DescribeIPStrategysStatusRequest) (response *DescribeIPStrategysStatusResponse, err error) {
    return c.DescribeIPStrategysStatusWithContext(context.Background(), request)
}

// DescribeIPStrategysStatus
// 本接口（DescribeIPStrategysStatus）用于查询服务IP策略列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategysStatusWithContext(ctx context.Context, request *DescribeIPStrategysStatusRequest) (response *DescribeIPStrategysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategysStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPStrategysStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPStrategysStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogSearchRequest() (request *DescribeLogSearchRequest) {
    request = &DescribeLogSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeLogSearch")
    
    
    return
}

func NewDescribeLogSearchResponse() (response *DescribeLogSearchResponse) {
    response = &DescribeLogSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogSearch
// 本接口DescribeLogSearch用于搜索日志
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNSUPPORTEDOPERATION_CLSSEARCHTIME = "UnsupportedOperation.ClsSearchTime"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    return c.DescribeLogSearchWithContext(context.Background(), request)
}

// DescribeLogSearch
// 本接口DescribeLogSearch用于搜索日志
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNSUPPORTEDOPERATION_CLSSEARCHTIME = "UnsupportedOperation.ClsSearchTime"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeLogSearchWithContext(ctx context.Context, request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    if request == nil {
        request = NewDescribeLogSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginRequest() (request *DescribePluginRequest) {
    request = &DescribePluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribePlugin")
    
    
    return
}

func NewDescribePluginResponse() (response *DescribePluginResponse) {
    response = &DescribePluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlugin
// 展示插件详情，支持按照插件ID进行。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePlugin(request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    return c.DescribePluginWithContext(context.Background(), request)
}

// DescribePlugin
// 展示插件详情，支持按照插件ID进行。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePluginWithContext(ctx context.Context, request *DescribePluginRequest) (response *DescribePluginResponse, err error) {
    if request == nil {
        request = NewDescribePluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginApisRequest() (request *DescribePluginApisRequest) {
    request = &DescribePluginApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribePluginApis")
    
    
    return
}

func NewDescribePluginApisResponse() (response *DescribePluginApisResponse) {
    response = &DescribePluginApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePluginApis
// 查询指定插件下绑定的API信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePluginApis(request *DescribePluginApisRequest) (response *DescribePluginApisResponse, err error) {
    return c.DescribePluginApisWithContext(context.Background(), request)
}

// DescribePluginApis
// 查询指定插件下绑定的API信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribePluginApisWithContext(ctx context.Context, request *DescribePluginApisRequest) (response *DescribePluginApisResponse, err error) {
    if request == nil {
        request = NewDescribePluginApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginsRequest() (request *DescribePluginsRequest) {
    request = &DescribePluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribePlugins")
    
    
    return
}

func NewDescribePluginsResponse() (response *DescribePluginsResponse) {
    response = &DescribePluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlugins
// 展示插件列表和详情，支持分页，支持按照插件类型查询，支持按照插件ID批量查询，支持按照插件名称查询。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribePlugins(request *DescribePluginsRequest) (response *DescribePluginsResponse, err error) {
    return c.DescribePluginsWithContext(context.Background(), request)
}

// DescribePlugins
// 展示插件列表和详情，支持分页，支持按照插件类型查询，支持按照插件ID批量查询，支持按照插件名称查询。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribePluginsWithContext(ctx context.Context, request *DescribePluginsRequest) (response *DescribePluginsResponse, err error) {
    if request == nil {
        request = NewDescribePluginsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginsByApiRequest() (request *DescribePluginsByApiRequest) {
    request = &DescribePluginsByApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribePluginsByApi")
    
    
    return
}

func NewDescribePluginsByApiResponse() (response *DescribePluginsByApiResponse) {
    response = &DescribePluginsByApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePluginsByApi
// 展示API上已绑定的插件列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribePluginsByApi(request *DescribePluginsByApiRequest) (response *DescribePluginsByApiResponse, err error) {
    return c.DescribePluginsByApiWithContext(context.Background(), request)
}

// DescribePluginsByApi
// 展示API上已绑定的插件列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribePluginsByApiWithContext(ctx context.Context, request *DescribePluginsByApiRequest) (response *DescribePluginsByApiResponse, err error) {
    if request == nil {
        request = NewDescribePluginsByApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePluginsByApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePluginsByApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
    request = &DescribeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeService")
    
    
    return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
    response = &DescribeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeService
// 本接口（DescribeService）用于查询一个服务的详细信息、包括服务的描述、域名、协议、创建时间、发布情况等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    return c.DescribeServiceWithContext(context.Background(), request)
}

// DescribeService
// 本接口（DescribeService）用于查询一个服务的详细信息、包括服务的描述、域名、协议、创建时间、发布情况等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeServiceWithContext(ctx context.Context, request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    if request == nil {
        request = NewDescribeServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentListRequest() (request *DescribeServiceEnvironmentListRequest) {
    request = &DescribeServiceEnvironmentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentList")
    
    
    return
}

func NewDescribeServiceEnvironmentListResponse() (response *DescribeServiceEnvironmentListResponse) {
    response = &DescribeServiceEnvironmentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceEnvironmentList
// 本接口（DescribeServiceEnvironmentList）用于查询一个服务的环境列表，可查询到此服务下所有环境及其状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentList(request *DescribeServiceEnvironmentListRequest) (response *DescribeServiceEnvironmentListResponse, err error) {
    return c.DescribeServiceEnvironmentListWithContext(context.Background(), request)
}

// DescribeServiceEnvironmentList
// 本接口（DescribeServiceEnvironmentList）用于查询一个服务的环境列表，可查询到此服务下所有环境及其状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentListWithContext(ctx context.Context, request *DescribeServiceEnvironmentListRequest) (response *DescribeServiceEnvironmentListResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceEnvironmentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceEnvironmentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryRequest() (request *DescribeServiceEnvironmentReleaseHistoryRequest) {
    request = &DescribeServiceEnvironmentReleaseHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentReleaseHistory")
    
    
    return
}

func NewDescribeServiceEnvironmentReleaseHistoryResponse() (response *DescribeServiceEnvironmentReleaseHistoryResponse) {
    response = &DescribeServiceEnvironmentReleaseHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceEnvironmentReleaseHistory
// 本接口（DescribeServiceEnvironmentReleaseHistory）用于查询服务环境的发布历史。
//
// 用户在创建好服务后需要发布到某个环境中才能进行使用，本接口用于查询一个服务某个环境的发布记录。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentReleaseHistory(request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    return c.DescribeServiceEnvironmentReleaseHistoryWithContext(context.Background(), request)
}

// DescribeServiceEnvironmentReleaseHistory
// 本接口（DescribeServiceEnvironmentReleaseHistory）用于查询服务环境的发布历史。
//
// 用户在创建好服务后需要发布到某个环境中才能进行使用，本接口用于查询一个服务某个环境的发布记录。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentReleaseHistoryWithContext(ctx context.Context, request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentReleaseHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceEnvironmentReleaseHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceEnvironmentReleaseHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceEnvironmentStrategyRequest() (request *DescribeServiceEnvironmentStrategyRequest) {
    request = &DescribeServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceEnvironmentStrategy")
    
    
    return
}

func NewDescribeServiceEnvironmentStrategyResponse() (response *DescribeServiceEnvironmentStrategyResponse) {
    response = &DescribeServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceEnvironmentStrategy
// 本接口（DescribeServiceEnvironmentStrategy）用于展示服务限流策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentStrategy(request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    return c.DescribeServiceEnvironmentStrategyWithContext(context.Background(), request)
}

// DescribeServiceEnvironmentStrategy
// 本接口（DescribeServiceEnvironmentStrategy）用于展示服务限流策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentStrategyWithContext(ctx context.Context, request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceForApiAppRequest() (request *DescribeServiceForApiAppRequest) {
    request = &DescribeServiceForApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceForApiApp")
    
    
    return
}

func NewDescribeServiceForApiAppResponse() (response *DescribeServiceForApiAppResponse) {
    response = &DescribeServiceForApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceForApiApp
// 本接口（DescribeServiceForApiApp）用于应用使用者查询一个服务的详细信息、包括服务的描述、域名、协议等信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceForApiApp(request *DescribeServiceForApiAppRequest) (response *DescribeServiceForApiAppResponse, err error) {
    return c.DescribeServiceForApiAppWithContext(context.Background(), request)
}

// DescribeServiceForApiApp
// 本接口（DescribeServiceForApiApp）用于应用使用者查询一个服务的详细信息、包括服务的描述、域名、协议等信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceForApiAppWithContext(ctx context.Context, request *DescribeServiceForApiAppRequest) (response *DescribeServiceForApiAppResponse, err error) {
    if request == nil {
        request = NewDescribeServiceForApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceForApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceForApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceReleaseVersionRequest() (request *DescribeServiceReleaseVersionRequest) {
    request = &DescribeServiceReleaseVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceReleaseVersion")
    
    
    return
}

func NewDescribeServiceReleaseVersionResponse() (response *DescribeServiceReleaseVersionResponse) {
    response = &DescribeServiceReleaseVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceReleaseVersion
// 本接口（DescribeServiceReleaseVersion）查询一个服务下面所有已经发布的版本列表。
//
// 用户在发布服务时，常有多个版本发布，可使用本接口查询已发布的版本。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceReleaseVersion(request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    return c.DescribeServiceReleaseVersionWithContext(context.Background(), request)
}

// DescribeServiceReleaseVersion
// 本接口（DescribeServiceReleaseVersion）查询一个服务下面所有已经发布的版本列表。
//
// 用户在发布服务时，常有多个版本发布，可使用本接口查询已发布的版本。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceReleaseVersionWithContext(ctx context.Context, request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    if request == nil {
        request = NewDescribeServiceReleaseVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceReleaseVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceReleaseVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainMappingsRequest() (request *DescribeServiceSubDomainMappingsRequest) {
    request = &DescribeServiceSubDomainMappingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomainMappings")
    
    
    return
}

func NewDescribeServiceSubDomainMappingsResponse() (response *DescribeServiceSubDomainMappingsResponse) {
    response = &DescribeServiceSubDomainMappingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceSubDomainMappings
// 本接口（DescribeServiceSubDomainMappings）用于查询自定义域名的路径映射。
//
// API 网关可绑定自定义域名到服务，并且可以对自定义域名的路径进行映射，可自定义不同的路径映射到服务中的三个环境，本接口用于查询绑定服务的自定义域名的路径映射列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeServiceSubDomainMappings(request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
    return c.DescribeServiceSubDomainMappingsWithContext(context.Background(), request)
}

// DescribeServiceSubDomainMappings
// 本接口（DescribeServiceSubDomainMappings）用于查询自定义域名的路径映射。
//
// API 网关可绑定自定义域名到服务，并且可以对自定义域名的路径进行映射，可自定义不同的路径映射到服务中的三个环境，本接口用于查询绑定服务的自定义域名的路径映射列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
func (c *Client) DescribeServiceSubDomainMappingsWithContext(ctx context.Context, request *DescribeServiceSubDomainMappingsRequest) (response *DescribeServiceSubDomainMappingsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainMappingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceSubDomainMappings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceSubDomainMappingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceSubDomainsRequest() (request *DescribeServiceSubDomainsRequest) {
    request = &DescribeServiceSubDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceSubDomains")
    
    
    return
}

func NewDescribeServiceSubDomainsResponse() (response *DescribeServiceSubDomainsResponse) {
    response = &DescribeServiceSubDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceSubDomains
// 本接口（DescribeServiceSubDomains）用于查询自定义域名列表。
//
// API 网关可绑定自定义域名到服务，用于服务调用。此接口用于查询用户绑定在服务的自定义域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBESERVICESUBDOMAINSERROR = "FailedOperation.DescribeServiceSubDomainsError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceSubDomains(request *DescribeServiceSubDomainsRequest) (response *DescribeServiceSubDomainsResponse, err error) {
    return c.DescribeServiceSubDomainsWithContext(context.Background(), request)
}

// DescribeServiceSubDomains
// 本接口（DescribeServiceSubDomains）用于查询自定义域名列表。
//
// API 网关可绑定自定义域名到服务，用于服务调用。此接口用于查询用户绑定在服务的自定义域名列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBESERVICESUBDOMAINSERROR = "FailedOperation.DescribeServiceSubDomainsError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceSubDomainsWithContext(ctx context.Context, request *DescribeServiceSubDomainsRequest) (response *DescribeServiceSubDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceSubDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceSubDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceSubDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceUsagePlanRequest() (request *DescribeServiceUsagePlanRequest) {
    request = &DescribeServiceUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServiceUsagePlan")
    
    
    return
}

func NewDescribeServiceUsagePlanResponse() (response *DescribeServiceUsagePlanResponse) {
    response = &DescribeServiceUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceUsagePlan
// 本接口（DescribeServiceUsagePlan）用于查询服务使用计划详情。
//
// 服务若需要鉴权限流生效，则需要绑定使用计划到此服务中，本接口用于查询绑定到一个服务的所有使用计划。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceUsagePlan(request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    return c.DescribeServiceUsagePlanWithContext(context.Background(), request)
}

// DescribeServiceUsagePlan
// 本接口（DescribeServiceUsagePlan）用于查询服务使用计划详情。
//
// 服务若需要鉴权限流生效，则需要绑定使用计划到此服务中，本接口用于查询绑定到一个服务的所有使用计划。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceUsagePlanWithContext(ctx context.Context, request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeServiceUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServicesStatusRequest() (request *DescribeServicesStatusRequest) {
    request = &DescribeServicesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeServicesStatus")
    
    
    return
}

func NewDescribeServicesStatusResponse() (response *DescribeServicesStatusResponse) {
    response = &DescribeServicesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServicesStatus
// 本接口（DescribeServicesStatus）用于搜索查询某一个服务或多个服务的列表，并返回服务相关的域名、时间等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUES = "InvalidParameterValue.InvalidTagValues"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeServicesStatus(request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
    return c.DescribeServicesStatusWithContext(context.Background(), request)
}

// DescribeServicesStatus
// 本接口（DescribeServicesStatus）用于搜索查询某一个服务或多个服务的列表，并返回服务相关的域名、时间等信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUES = "InvalidParameterValue.InvalidTagValues"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeServicesStatusWithContext(ctx context.Context, request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeServicesStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServicesStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServicesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpstreamBindApisRequest() (request *DescribeUpstreamBindApisRequest) {
    request = &DescribeUpstreamBindApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUpstreamBindApis")
    
    
    return
}

func NewDescribeUpstreamBindApisResponse() (response *DescribeUpstreamBindApisResponse) {
    response = &DescribeUpstreamBindApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUpstreamBindApis
// 查询后端通道所绑定的API列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUpstreamBindApis(request *DescribeUpstreamBindApisRequest) (response *DescribeUpstreamBindApisResponse, err error) {
    return c.DescribeUpstreamBindApisWithContext(context.Background(), request)
}

// DescribeUpstreamBindApis
// 查询后端通道所绑定的API列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUpstreamBindApisWithContext(ctx context.Context, request *DescribeUpstreamBindApisRequest) (response *DescribeUpstreamBindApisResponse, err error) {
    if request == nil {
        request = NewDescribeUpstreamBindApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpstreamBindApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpstreamBindApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpstreamsRequest() (request *DescribeUpstreamsRequest) {
    request = &DescribeUpstreamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUpstreams")
    
    
    return
}

func NewDescribeUpstreamsResponse() (response *DescribeUpstreamsResponse) {
    response = &DescribeUpstreamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUpstreams
// 查询后端通道列表详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeUpstreams(request *DescribeUpstreamsRequest) (response *DescribeUpstreamsResponse, err error) {
    return c.DescribeUpstreamsWithContext(context.Background(), request)
}

// DescribeUpstreams
// 查询后端通道列表详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeUpstreamsWithContext(ctx context.Context, request *DescribeUpstreamsRequest) (response *DescribeUpstreamsResponse, err error) {
    if request == nil {
        request = NewDescribeUpstreamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpstreams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpstreamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanRequest() (request *DescribeUsagePlanRequest) {
    request = &DescribeUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlan")
    
    
    return
}

func NewDescribeUsagePlanResponse() (response *DescribeUsagePlanResponse) {
    response = &DescribeUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsagePlan
// 本接口（DescribeUsagePlan）用于查询一个使用计划的详细信息，包括名称、QPS、创建时间绑定的环境等。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlan(request *DescribeUsagePlanRequest) (response *DescribeUsagePlanResponse, err error) {
    return c.DescribeUsagePlanWithContext(context.Background(), request)
}

// DescribeUsagePlan
// 本接口（DescribeUsagePlan）用于查询一个使用计划的详细信息，包括名称、QPS、创建时间绑定的环境等。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanWithContext(ctx context.Context, request *DescribeUsagePlanRequest) (response *DescribeUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanEnvironmentsRequest() (request *DescribeUsagePlanEnvironmentsRequest) {
    request = &DescribeUsagePlanEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanEnvironments")
    
    
    return
}

func NewDescribeUsagePlanEnvironmentsResponse() (response *DescribeUsagePlanEnvironmentsResponse) {
    response = &DescribeUsagePlanEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsagePlanEnvironments
// 本接口（DescribeUsagePlanEnvironments）用于查询使用计划绑定的环境列表。
//
// 用户在绑定了某个使用计划到环境后，可使用本接口查询这个使用计划绑定的所有服务的环境。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanEnvironments(request *DescribeUsagePlanEnvironmentsRequest) (response *DescribeUsagePlanEnvironmentsResponse, err error) {
    return c.DescribeUsagePlanEnvironmentsWithContext(context.Background(), request)
}

// DescribeUsagePlanEnvironments
// 本接口（DescribeUsagePlanEnvironments）用于查询使用计划绑定的环境列表。
//
// 用户在绑定了某个使用计划到环境后，可使用本接口查询这个使用计划绑定的所有服务的环境。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanEnvironmentsWithContext(ctx context.Context, request *DescribeUsagePlanEnvironmentsRequest) (response *DescribeUsagePlanEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlanEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlanEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlanSecretIdsRequest() (request *DescribeUsagePlanSecretIdsRequest) {
    request = &DescribeUsagePlanSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlanSecretIds")
    
    
    return
}

func NewDescribeUsagePlanSecretIdsResponse() (response *DescribeUsagePlanSecretIdsResponse) {
    response = &DescribeUsagePlanSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsagePlanSecretIds
// 本接口（DescribeUsagePlanSecretIds）用于查询使用计划绑定的密钥列表。
//
// 在 API 网关中，一个使用计划可绑定多个密钥对，可使用本接口查询使用计划绑定的密钥列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanSecretIds(request *DescribeUsagePlanSecretIdsRequest) (response *DescribeUsagePlanSecretIdsResponse, err error) {
    return c.DescribeUsagePlanSecretIdsWithContext(context.Background(), request)
}

// DescribeUsagePlanSecretIds
// 本接口（DescribeUsagePlanSecretIds）用于查询使用计划绑定的密钥列表。
//
// 在 API 网关中，一个使用计划可绑定多个密钥对，可使用本接口查询使用计划绑定的密钥列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) DescribeUsagePlanSecretIdsWithContext(ctx context.Context, request *DescribeUsagePlanSecretIdsRequest) (response *DescribeUsagePlanSecretIdsResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlanSecretIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlanSecretIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlanSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsagePlansStatusRequest() (request *DescribeUsagePlansStatusRequest) {
    request = &DescribeUsagePlansStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DescribeUsagePlansStatus")
    
    
    return
}

func NewDescribeUsagePlansStatusResponse() (response *DescribeUsagePlansStatusResponse) {
    response = &DescribeUsagePlansStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsagePlansStatus
// 本接口（DescribeUsagePlanStatus）用于查询使用计划的列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeUsagePlansStatus(request *DescribeUsagePlansStatusRequest) (response *DescribeUsagePlansStatusResponse, err error) {
    return c.DescribeUsagePlansStatusWithContext(context.Background(), request)
}

// DescribeUsagePlansStatus
// 本接口（DescribeUsagePlanStatus）用于查询使用计划的列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeUsagePlansStatusWithContext(ctx context.Context, request *DescribeUsagePlansStatusRequest) (response *DescribeUsagePlansStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUsagePlansStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsagePlansStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsagePlansStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDetachPluginRequest() (request *DetachPluginRequest) {
    request = &DetachPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DetachPlugin")
    
    
    return
}

func NewDetachPluginResponse() (response *DetachPluginResponse) {
    response = &DetachPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachPlugin
// 解除插件与API绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DetachPlugin(request *DetachPluginRequest) (response *DetachPluginResponse, err error) {
    return c.DetachPluginWithContext(context.Background(), request)
}

// DetachPlugin
// 解除插件与API绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"
//  INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DetachPluginWithContext(ctx context.Context, request *DetachPluginRequest) (response *DetachPluginResponse, err error) {
    if request == nil {
        request = NewDetachPluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDisableApiKeyRequest() (request *DisableApiKeyRequest) {
    request = &DisableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "DisableApiKey")
    
    
    return
}

func NewDisableApiKeyResponse() (response *DisableApiKeyResponse) {
    response = &DisableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableApiKey
// 本接口（DisableApiKey）用于禁用一对 API 密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) DisableApiKey(request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
    return c.DisableApiKeyWithContext(context.Background(), request)
}

// DisableApiKey
// 本接口（DisableApiKey）用于禁用一对 API 密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) DisableApiKeyWithContext(ctx context.Context, request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
    if request == nil {
        request = NewDisableApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableApiKeyRequest() (request *EnableApiKeyRequest) {
    request = &EnableApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "EnableApiKey")
    
    
    return
}

func NewEnableApiKeyResponse() (response *EnableApiKeyResponse) {
    response = &EnableApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableApiKey
// 本接口（EnableApiKey）用于启动一对被禁用的 API 密钥。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) EnableApiKey(request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
    return c.EnableApiKeyWithContext(context.Background(), request)
}

// EnableApiKey
// 本接口（EnableApiKey）用于启动一对被禁用的 API 密钥。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"
func (c *Client) EnableApiKeyWithContext(ctx context.Context, request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
    if request == nil {
        request = NewEnableApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateApiDocumentRequest() (request *GenerateApiDocumentRequest) {
    request = &GenerateApiDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "GenerateApiDocument")
    
    
    return
}

func NewGenerateApiDocumentResponse() (response *GenerateApiDocumentResponse) {
    response = &GenerateApiDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateApiDocument
// 本接口（GenerateApiDocument）用于自动生成 API 文档和 SDK，一个服务的一个环境生成一份文档和 SDK。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATEAPIDOCUMENTERROR = "FailedOperation.GenerateApiDocumentError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGENLANGUAGE = "InvalidParameterValue.InvalidGenLanguage"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
func (c *Client) GenerateApiDocument(request *GenerateApiDocumentRequest) (response *GenerateApiDocumentResponse, err error) {
    return c.GenerateApiDocumentWithContext(context.Background(), request)
}

// GenerateApiDocument
// 本接口（GenerateApiDocument）用于自动生成 API 文档和 SDK，一个服务的一个环境生成一份文档和 SDK。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GENERATEAPIDOCUMENTERROR = "FailedOperation.GenerateApiDocumentError"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  INVALIDPARAMETERVALUE_INVALIDGENLANGUAGE = "InvalidParameterValue.InvalidGenLanguage"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
func (c *Client) GenerateApiDocumentWithContext(ctx context.Context, request *GenerateApiDocumentRequest) (response *GenerateApiDocumentResponse, err error) {
    if request == nil {
        request = NewGenerateApiDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateApiDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateApiDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewImportOpenApiRequest() (request *ImportOpenApiRequest) {
    request = &ImportOpenApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ImportOpenApi")
    
    
    return
}

func NewImportOpenApiResponse() (response *ImportOpenApiResponse) {
    response = &ImportOpenApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportOpenApi
// 本接口（ImportOpenApi）用于将OpenAPI规范定义的API导入到API网关。 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPARAMETER = "InvalidParameterValue.UnsupportedParameter"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) ImportOpenApi(request *ImportOpenApiRequest) (response *ImportOpenApiResponse, err error) {
    return c.ImportOpenApiWithContext(context.Background(), request)
}

// ImportOpenApi
// 本接口（ImportOpenApi）用于将OpenAPI规范定义的API导入到API网关。 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDPARAMETER = "InvalidParameterValue.UnsupportedParameter"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) ImportOpenApiWithContext(ctx context.Context, request *ImportOpenApiRequest) (response *ImportOpenApiResponse, err error) {
    if request == nil {
        request = NewImportOpenApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportOpenApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportOpenApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAPIDocRequest() (request *ModifyAPIDocRequest) {
    request = &ModifyAPIDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyAPIDoc")
    
    
    return
}

func NewModifyAPIDocResponse() (response *ModifyAPIDocResponse) {
    response = &ModifyAPIDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAPIDoc
// 修改 API 文档
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyAPIDoc(request *ModifyAPIDocRequest) (response *ModifyAPIDocResponse, err error) {
    return c.ModifyAPIDocWithContext(context.Background(), request)
}

// ModifyAPIDoc
// 修改 API 文档
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyAPIDocWithContext(ctx context.Context, request *ModifyAPIDocRequest) (response *ModifyAPIDocResponse, err error) {
    if request == nil {
        request = NewModifyAPIDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAPIDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAPIDocResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiRequest() (request *ModifyApiRequest) {
    request = &ModifyApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApi")
    
    
    return
}

func NewModifyApiResponse() (response *ModifyApiResponse) {
    response = &ModifyApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApi
// 本接口（ModifyApi）用于修改 API 接口，可调用此接口对已经配置的 API 接口进行编辑修改。修改后的 API 需要重新发布 API 所在的服务到对应环境方能生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_BASICSERVICENOTALLOWATTACHPLUGIN = "InvalidParameter.BasicServiceNotAllowAttachPlugin"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONSTANTPARAMETERS = "InvalidParameterValue.InvalidConstantParameters"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDMETHOD = "InvalidParameterValue.InvalidMethod"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_MODIFYEIAMAUTHAPI = "UnsupportedOperation.ModifyEIAMAuthApi"
//  UNSUPPORTEDOPERATION_MODIFYPROTOCOL = "UnsupportedOperation.ModifyProtocol"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
func (c *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    return c.ModifyApiWithContext(context.Background(), request)
}

// ModifyApi
// 本接口（ModifyApi）用于修改 API 接口，可调用此接口对已经配置的 API 接口进行编辑修改。修改后的 API 需要重新发布 API 所在的服务到对应环境方能生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_BASICSERVICENOTALLOWATTACHPLUGIN = "InvalidParameter.BasicServiceNotAllowAttachPlugin"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONSTANTPARAMETERS = "InvalidParameterValue.InvalidConstantParameters"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDMETHOD = "InvalidParameterValue.InvalidMethod"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_MODIFYEIAMAUTHAPI = "UnsupportedOperation.ModifyEIAMAuthApi"
//  UNSUPPORTEDOPERATION_MODIFYPROTOCOL = "UnsupportedOperation.ModifyProtocol"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
func (c *Client) ModifyApiWithContext(ctx context.Context, request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    if request == nil {
        request = NewModifyApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiAppRequest() (request *ModifyApiAppRequest) {
    request = &ModifyApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiApp")
    
    
    return
}

func NewModifyApiAppResponse() (response *ModifyApiAppResponse) {
    response = &ModifyApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApiApp
// 本接口（ModifyApiApp）用于修改已经创建的应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) ModifyApiApp(request *ModifyApiAppRequest) (response *ModifyApiAppResponse, err error) {
    return c.ModifyApiAppWithContext(context.Background(), request)
}

// ModifyApiApp
// 本接口（ModifyApiApp）用于修改已经创建的应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) ModifyApiAppWithContext(ctx context.Context, request *ModifyApiAppRequest) (response *ModifyApiAppResponse, err error) {
    if request == nil {
        request = NewModifyApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiEnvironmentStrategyRequest() (request *ModifyApiEnvironmentStrategyRequest) {
    request = &ModifyApiEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiEnvironmentStrategy")
    
    
    return
}

func NewModifyApiEnvironmentStrategyResponse() (response *ModifyApiEnvironmentStrategyResponse) {
    response = &ModifyApiEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApiEnvironmentStrategy
// 本接口（ModifyApiEnvironmentStrategy）用于修改API限流策略
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiEnvironmentStrategy(request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    return c.ModifyApiEnvironmentStrategyWithContext(context.Background(), request)
}

// ModifyApiEnvironmentStrategy
// 本接口（ModifyApiEnvironmentStrategy）用于修改API限流策略
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiEnvironmentStrategyWithContext(ctx context.Context, request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyApiEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiIncrementRequest() (request *ModifyApiIncrementRequest) {
    request = &ModifyApiIncrementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyApiIncrement")
    
    
    return
}

func NewModifyApiIncrementResponse() (response *ModifyApiIncrementResponse) {
    response = &ModifyApiIncrementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApiIncrement
// 提供增量更新API能力，主要是给程序调用（区别于ModifyApi，该接口是需要传入API的全量参数，对console使用较友好）
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_NOTHINGMODIFYFOROAUTH = "InvalidParameterValue.NothingModifyForOauth"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    return c.ModifyApiIncrementWithContext(context.Background(), request)
}

// ModifyApiIncrement
// 提供增量更新API能力，主要是给程序调用（区别于ModifyApi，该接口是需要传入API的全量参数，对console使用较友好）
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_NOTHINGMODIFYFOROAUTH = "InvalidParameterValue.NothingModifyForOauth"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyApiIncrementWithContext(ctx context.Context, request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    if request == nil {
        request = NewModifyApiIncrementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiIncrement require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiIncrementResponse()
    err = c.Send(request, response)
    return
}

func NewModifyExclusiveInstanceRequest() (request *ModifyExclusiveInstanceRequest) {
    request = &ModifyExclusiveInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyExclusiveInstance")
    
    
    return
}

func NewModifyExclusiveInstanceResponse() (response *ModifyExclusiveInstanceResponse) {
    response = &ModifyExclusiveInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyExclusiveInstance
// 本接口（ModifyExclusiveInstance）用于修改独享实例信息。​
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) ModifyExclusiveInstance(request *ModifyExclusiveInstanceRequest) (response *ModifyExclusiveInstanceResponse, err error) {
    return c.ModifyExclusiveInstanceWithContext(context.Background(), request)
}

// ModifyExclusiveInstance
// 本接口（ModifyExclusiveInstance）用于修改独享实例信息。​
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"
//  FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"
//  FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) ModifyExclusiveInstanceWithContext(ctx context.Context, request *ModifyExclusiveInstanceRequest) (response *ModifyExclusiveInstanceResponse, err error) {
    if request == nil {
        request = NewModifyExclusiveInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyExclusiveInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyExclusiveInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIPStrategyRequest() (request *ModifyIPStrategyRequest) {
    request = &ModifyIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyIPStrategy")
    
    
    return
}

func NewModifyIPStrategyResponse() (response *ModifyIPStrategyResponse) {
    response = &ModifyIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIPStrategy
// 本接口（ModifyIPStrategy）用于修改服务IP策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) ModifyIPStrategy(request *ModifyIPStrategyRequest) (response *ModifyIPStrategyResponse, err error) {
    return c.ModifyIPStrategyWithContext(context.Background(), request)
}

// ModifyIPStrategy
// 本接口（ModifyIPStrategy）用于修改服务IP策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) ModifyIPStrategyWithContext(ctx context.Context, request *ModifyIPStrategyRequest) (response *ModifyIPStrategyResponse, err error) {
    if request == nil {
        request = NewModifyIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPluginRequest() (request *ModifyPluginRequest) {
    request = &ModifyPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyPlugin")
    
    
    return
}

func NewModifyPluginResponse() (response *ModifyPluginResponse) {
    response = &ModifyPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPlugin
// 修改API网关插件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) ModifyPlugin(request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
    return c.ModifyPluginWithContext(context.Background(), request)
}

// ModifyPlugin
// 修改API网关插件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"
//  INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"
//  INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"
//  INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"
//  INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"
//  MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) ModifyPluginWithContext(ctx context.Context, request *ModifyPluginRequest) (response *ModifyPluginResponse, err error) {
    if request == nil {
        request = NewModifyPluginRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPlugin require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPluginResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceRequest() (request *ModifyServiceRequest) {
    request = &ModifyServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyService")
    
    
    return
}

func NewModifyServiceResponse() (response *ModifyServiceResponse) {
    response = &ModifyServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyService
// 本接口（ModifyService）用于修改服务的相关信息。当服务创建后，服务的名称、描述和服务类型均可被修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_MODIFYNETTYPE = "UnsupportedOperation.ModifyNetType"
//  UNSUPPORTEDOPERATION_REDUCENETTYPES = "UnsupportedOperation.ReduceNetTypes"
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    return c.ModifyServiceWithContext(context.Background(), request)
}

// ModifyService
// 本接口（ModifyService）用于修改服务的相关信息。当服务创建后，服务的名称、描述和服务类型均可被修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_MODIFYNETTYPE = "UnsupportedOperation.ModifyNetType"
//  UNSUPPORTEDOPERATION_REDUCENETTYPES = "UnsupportedOperation.ReduceNetTypes"
func (c *Client) ModifyServiceWithContext(ctx context.Context, request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    if request == nil {
        request = NewModifyServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceEnvironmentStrategyRequest() (request *ModifyServiceEnvironmentStrategyRequest) {
    request = &ModifyServiceEnvironmentStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyServiceEnvironmentStrategy")
    
    
    return
}

func NewModifyServiceEnvironmentStrategyResponse() (response *ModifyServiceEnvironmentStrategyResponse) {
    response = &ModifyServiceEnvironmentStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyServiceEnvironmentStrategy
// 本接口（ModifyServiceEnvironmentStrategy）用于修改服务限流策略
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
func (c *Client) ModifyServiceEnvironmentStrategy(request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    return c.ModifyServiceEnvironmentStrategyWithContext(context.Background(), request)
}

// ModifyServiceEnvironmentStrategy
// 本接口（ModifyServiceEnvironmentStrategy）用于修改服务限流策略
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
func (c *Client) ModifyServiceEnvironmentStrategyWithContext(ctx context.Context, request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServiceEnvironmentStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServiceEnvironmentStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServiceEnvironmentStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubDomainRequest() (request *ModifySubDomainRequest) {
    request = &ModifySubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifySubDomain")
    
    
    return
}

func NewModifySubDomainResponse() (response *ModifySubDomainResponse) {
    response = &ModifySubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubDomain
// 本接口（ModifySubDomain）用于修改服务的自定义域名设置中的路径映射，可以修改绑定自定义域名之前的路径映射规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEIDBINDERROR = "FailedOperation.CertificateIdBindError"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_PATHMAPPINGSETERROR = "FailedOperation.PathMappingSetError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SETCUSTOMPATHMAPPINGERROR = "FailedOperation.SetCustomPathMappingError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  FAILEDOPERATION_UNKNOWNPROTOCOLTYPEERROR = "FailedOperation.UnknownProtocolTypeError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) ModifySubDomain(request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    return c.ModifySubDomainWithContext(context.Background(), request)
}

// ModifySubDomain
// 本接口（ModifySubDomain）用于修改服务的自定义域名设置中的路径映射，可以修改绑定自定义域名之前的路径映射规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEIDBINDERROR = "FailedOperation.CertificateIdBindError"
//  FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"
//  FAILEDOPERATION_PATHMAPPINGSETERROR = "FailedOperation.PathMappingSetError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SETCUSTOMPATHMAPPINGERROR = "FailedOperation.SetCustomPathMappingError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  FAILEDOPERATION_UNKNOWNPROTOCOLTYPEERROR = "FailedOperation.UnknownProtocolTypeError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) ModifySubDomainWithContext(ctx context.Context, request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    if request == nil {
        request = NewModifySubDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUpstreamRequest() (request *ModifyUpstreamRequest) {
    request = &ModifyUpstreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUpstream")
    
    
    return
}

func NewModifyUpstreamResponse() (response *ModifyUpstreamResponse) {
    response = &ModifyUpstreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUpstream
// 修改后端通道
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUpstream(request *ModifyUpstreamRequest) (response *ModifyUpstreamResponse, err error) {
    return c.ModifyUpstreamWithContext(context.Background(), request)
}

// ModifyUpstream
// 修改后端通道
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"
//  INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"
//  INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUpstreamWithContext(ctx context.Context, request *ModifyUpstreamRequest) (response *ModifyUpstreamResponse, err error) {
    if request == nil {
        request = NewModifyUpstreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUpstream require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUpstreamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUsagePlanRequest() (request *ModifyUsagePlanRequest) {
    request = &ModifyUsagePlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ModifyUsagePlan")
    
    
    return
}

func NewModifyUsagePlanResponse() (response *ModifyUsagePlanResponse) {
    response = &ModifyUsagePlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUsagePlan
// 本接口（ModifyUsagePlan）用于修改使用计划的名称，描述及 QPS。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) ModifyUsagePlan(request *ModifyUsagePlanRequest) (response *ModifyUsagePlanResponse, err error) {
    return c.ModifyUsagePlanWithContext(context.Background(), request)
}

// ModifyUsagePlan
// 本接口（ModifyUsagePlan）用于修改使用计划的名称，描述及 QPS。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) ModifyUsagePlanWithContext(ctx context.Context, request *ModifyUsagePlanRequest) (response *ModifyUsagePlanResponse, err error) {
    if request == nil {
        request = NewModifyUsagePlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUsagePlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUsagePlanResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseServiceRequest() (request *ReleaseServiceRequest) {
    request = &ReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ReleaseService")
    
    
    return
}

func NewReleaseServiceResponse() (response *ReleaseServiceResponse) {
    response = &ReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseService
// 本接口（ReleaseService）用于发布服务。
//
// API 网关的服务创建后，需要发布到某个环境方生效后，使用者才能进行调用，此接口用于发布服务到环境，如 release 环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    return c.ReleaseServiceWithContext(context.Background(), request)
}

// ReleaseService
// 本接口（ReleaseService）用于发布服务。
//
// API 网关的服务创建后，需要发布到某个环境方生效后，使用者才能进行调用，此接口用于发布服务到环境，如 release 环境。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReleaseServiceWithContext(ctx context.Context, request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    if request == nil {
        request = NewReleaseServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseService require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAPIDocPasswordRequest() (request *ResetAPIDocPasswordRequest) {
    request = &ResetAPIDocPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "ResetAPIDocPassword")
    
    
    return
}

func NewResetAPIDocPasswordResponse() (response *ResetAPIDocPasswordResponse) {
    response = &ResetAPIDocPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAPIDocPassword
// 重置API文档密码
//
// 可能返回的错误码:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
func (c *Client) ResetAPIDocPassword(request *ResetAPIDocPasswordRequest) (response *ResetAPIDocPasswordResponse, err error) {
    return c.ResetAPIDocPasswordWithContext(context.Background(), request)
}

// ResetAPIDocPassword
// 重置API文档密码
//
// 可能返回的错误码:
//  FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
func (c *Client) ResetAPIDocPasswordWithContext(ctx context.Context, request *ResetAPIDocPasswordRequest) (response *ResetAPIDocPasswordResponse, err error) {
    if request == nil {
        request = NewResetAPIDocPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAPIDocPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAPIDocPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindEnvironmentRequest() (request *UnBindEnvironmentRequest) {
    request = &UnBindEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindEnvironment")
    
    
    return
}

func NewUnBindEnvironmentResponse() (response *UnBindEnvironmentResponse) {
    response = &UnBindEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindEnvironment
// 本接口（UnBindEnvironment）用于将使用计划从特定环境解绑。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindEnvironment(request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    return c.UnBindEnvironmentWithContext(context.Background(), request)
}

// UnBindEnvironment
// 本接口（UnBindEnvironment）用于将使用计划从特定环境解绑。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindEnvironmentWithContext(ctx context.Context, request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    if request == nil {
        request = NewUnBindEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindIPStrategyRequest() (request *UnBindIPStrategyRequest) {
    request = &UnBindIPStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindIPStrategy")
    
    
    return
}

func NewUnBindIPStrategyResponse() (response *UnBindIPStrategyResponse) {
    response = &UnBindIPStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindIPStrategy
// 本接口（UnBindIPStrategy）用于服务解绑IP策略。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
func (c *Client) UnBindIPStrategy(request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    return c.UnBindIPStrategyWithContext(context.Background(), request)
}

// UnBindIPStrategy
// 本接口（UnBindIPStrategy）用于服务解绑IP策略。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
func (c *Client) UnBindIPStrategyWithContext(ctx context.Context, request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    if request == nil {
        request = NewUnBindIPStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindIPStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindIPStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSecretIdsRequest() (request *UnBindSecretIdsRequest) {
    request = &UnBindSecretIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSecretIds")
    
    
    return
}

func NewUnBindSecretIdsResponse() (response *UnBindSecretIdsResponse) {
    response = &UnBindSecretIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindSecretIds
// 本接口（UnBindSecretIds）用于为使用计划解绑密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) UnBindSecretIds(request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
    return c.UnBindSecretIdsWithContext(context.Background(), request)
}

// UnBindSecretIds
// 本接口（UnBindSecretIds）用于为使用计划解绑密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
func (c *Client) UnBindSecretIdsWithContext(ctx context.Context, request *UnBindSecretIdsRequest) (response *UnBindSecretIdsResponse, err error) {
    if request == nil {
        request = NewUnBindSecretIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindSecretIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindSecretIdsResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindSubDomainRequest() (request *UnBindSubDomainRequest) {
    request = &UnBindSubDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnBindSubDomain")
    
    
    return
}

func NewUnBindSubDomainResponse() (response *UnBindSubDomainResponse) {
    response = &UnBindSubDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnBindSubDomain
// 本接口（UnBindSubDomain）用于解绑自定义域名。
//
// 用户使用 API 网关绑定了自定义域名到服务中后，若想要解绑此自定义域名，可使用此接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINNOTBINDSERVICE = "FailedOperation.DomainNotBindService"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
func (c *Client) UnBindSubDomain(request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    return c.UnBindSubDomainWithContext(context.Background(), request)
}

// UnBindSubDomain
// 本接口（UnBindSubDomain）用于解绑自定义域名。
//
// 用户使用 API 网关绑定了自定义域名到服务中后，若想要解绑此自定义域名，可使用此接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOMAINNOTBINDSERVICE = "FailedOperation.DomainNotBindService"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
func (c *Client) UnBindSubDomainWithContext(ctx context.Context, request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    if request == nil {
        request = NewUnBindSubDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindSubDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindSubDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUnReleaseServiceRequest() (request *UnReleaseServiceRequest) {
    request = &UnReleaseServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnReleaseService")
    
    
    return
}

func NewUnReleaseServiceResponse() (response *UnReleaseServiceResponse) {
    response = &UnReleaseServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnReleaseService
// 本接口（UnReleaseService）用于下线服务。
//
// 用户发布服务到某个环境后，此服务中的 API 方可被调用者进行调用，当用户需要将此服务从发布环境中下线时，可调用此 API。下线后的服务不可被调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnReleaseService(request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    return c.UnReleaseServiceWithContext(context.Background(), request)
}

// UnReleaseService
// 本接口（UnReleaseService）用于下线服务。
//
// 用户发布服务到某个环境后，此服务中的 API 方可被调用者进行调用，当用户需要将此服务从发布环境中下线时，可调用此 API。下线后的服务不可被调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnReleaseServiceWithContext(ctx context.Context, request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    if request == nil {
        request = NewUnReleaseServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnReleaseService require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnReleaseServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindApiAppRequest() (request *UnbindApiAppRequest) {
    request = &UnbindApiAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UnbindApiApp")
    
    
    return
}

func NewUnbindApiAppResponse() (response *UnbindApiAppResponse) {
    response = &UnbindApiAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindApiApp
// 本接口（UnbindApiApp）用于解除应用和API绑定。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_RESOURCEUNASSOCIATED = "UnsupportedOperation.ResourceUnassociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) UnbindApiApp(request *UnbindApiAppRequest) (response *UnbindApiAppResponse, err error) {
    return c.UnbindApiAppWithContext(context.Background(), request)
}

// UnbindApiApp
// 本接口（UnbindApiApp）用于解除应用和API绑定。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_RESOURCEUNASSOCIATED = "UnsupportedOperation.ResourceUnassociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
func (c *Client) UnbindApiAppWithContext(ctx context.Context, request *UnbindApiAppRequest) (response *UnbindApiAppResponse, err error) {
    if request == nil {
        request = NewUnbindApiAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindApiApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindApiAppResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiAppKeyRequest() (request *UpdateApiAppKeyRequest) {
    request = &UpdateApiAppKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiAppKey")
    
    
    return
}

func NewUpdateApiAppKeyResponse() (response *UpdateApiAppKeyResponse) {
    response = &UpdateApiAppKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateApiAppKey
// 本接口（UpdateApiAppKey）用于更新应用密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) UpdateApiAppKey(request *UpdateApiAppKeyRequest) (response *UpdateApiAppKeyResponse, err error) {
    return c.UpdateApiAppKeyWithContext(context.Background(), request)
}

// UpdateApiAppKey
// 本接口（UpdateApiAppKey）用于更新应用密钥。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"
//  FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"
func (c *Client) UpdateApiAppKeyWithContext(ctx context.Context, request *UpdateApiAppKeyRequest) (response *UpdateApiAppKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiAppKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiAppKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiAppKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiKeyRequest() (request *UpdateApiKeyRequest) {
    request = &UpdateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateApiKey")
    
    
    return
}

func NewUpdateApiKeyResponse() (response *UpdateApiKeyResponse) {
    response = &UpdateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateApiKey
// 本接口（UpdateApiKey）用于更换用户已创建的一对 API 密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
func (c *Client) UpdateApiKey(request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    return c.UpdateApiKeyWithContext(context.Background(), request)
}

// UpdateApiKey
// 本接口（UpdateApiKey）用于更换用户已创建的一对 API 密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
func (c *Client) UpdateApiKeyWithContext(ctx context.Context, request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
    request = &UpdateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apigateway", APIVersion, "UpdateService")
    
    
    return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
    response = &UpdateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateService
// 本接口（UpdateService）用于从服务发布的环境中运行版本切换到特定版本。用户在使用 API 网关创建服务并发布服务到某个环境后，多因为开发过程会产生多个版本，此时可调用本接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    return c.UpdateServiceWithContext(context.Background(), request)
}

// UpdateService
// 本接口（UpdateService）用于从服务发布的环境中运行版本切换到特定版本。用户在使用 API 网关创建服务并发布服务到某个环境后，多因为开发过程会产生多个版本，此时可调用本接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) UpdateServiceWithContext(ctx context.Context, request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
