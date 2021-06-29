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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
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
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) BindEnvironment(request *BindEnvironmentRequest) (response *BindEnvironmentResponse, err error) {
    if request == nil {
        request = NewBindEnvironmentRequest()
    }
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
    if request == nil {
        request = NewBindIPStrategyRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_ALREADYBINDUSAGEPLAN = "UnsupportedOperation.AlreadyBindUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDBINDAPIKEY = "UnsupportedOperation.UnsupportedBindApiKey"
func (c *Client) BindSecretIds(request *BindSecretIdsRequest) (response *BindSecretIdsResponse, err error) {
    if request == nil {
        request = NewBindSecretIdsRequest()
    }
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
//  FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"
//  FAILEDOPERATION_DOMAINRESOLVEERROR = "FailedOperation.DomainResolveError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_ISDEFAULTMAPPING = "FailedOperation.IsDefaultMapping"
//  FAILEDOPERATION_NETSUBDOMAINERROR = "FailedOperation.NetSubDomainError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  INTERNALERROR = "InternalError"
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
    if request == nil {
        request = NewBindSubDomainRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) BuildAPIDoc(request *BuildAPIDocRequest) (response *BuildAPIDocResponse, err error) {
    if request == nil {
        request = NewBuildAPIDocRequest()
    }
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
    if request == nil {
        request = NewCreateAPIDocRequest()
    }
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
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SCFERROR = "FailedOperation.ScfError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"
//  INVALIDPARAMETERVALUE_INVALIDAPITYPE = "InvalidParameterValue.InvalidApiType"
//  INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"
//  INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"
//  INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEMOCKRETURNMESSAGE = "InvalidParameterValue.InvalidServiceMockReturnMessage"
//  INVALIDPARAMETERVALUE_INVALIDSERVICEPARAMETERS = "InvalidParameterValue.InvalidServiceParameters"
//  INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
func (c *Client) CreateApi(request *CreateApiRequest) (response *CreateApiResponse, err error) {
    if request == nil {
        request = NewCreateApiRequest()
    }
    response = NewCreateApiResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
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
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateIPStrategy(request *CreateIPStrategyRequest) (response *CreateIPStrategyResponse, err error) {
    if request == nil {
        request = NewCreateIPStrategyRequest()
    }
    response = NewCreateIPStrategyResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
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
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    response = NewCreateServiceResponse()
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
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_USAGEPLANLIMITEXCEEDED = "LimitExceeded.UsagePlanLimitExceeded"
func (c *Client) CreateUsagePlan(request *CreateUsagePlanRequest) (response *CreateUsagePlanResponse, err error) {
    if request == nil {
        request = NewCreateUsagePlanRequest()
    }
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
    if request == nil {
        request = NewDeleteAPIDocRequest()
    }
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
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteApi(request *DeleteApiRequest) (response *DeleteApiResponse, err error) {
    if request == nil {
        request = NewDeleteApiRequest()
    }
    response = NewDeleteApiResponse()
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
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
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
    if request == nil {
        request = NewDeleteIPStrategyRequest()
    }
    response = NewDeleteIPStrategyResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDDELETESERVICE = "UnsupportedOperation.UnsupportedDeleteService"
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
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
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) DeleteServiceSubDomainMapping(request *DeleteServiceSubDomainMappingRequest) (response *DeleteServiceSubDomainMappingResponse, err error) {
    if request == nil {
        request = NewDeleteServiceSubDomainMappingRequest()
    }
    response = NewDeleteServiceSubDomainMappingResponse()
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
    if request == nil {
        request = NewDeleteUsagePlanRequest()
    }
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
//  UNSUPPORTEDOPERATION_NOUSAGEPLANENV = "UnsupportedOperation.NoUsagePlanEnv"
func (c *Client) DemoteServiceUsagePlan(request *DemoteServiceUsagePlanRequest) (response *DemoteServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDemoteServiceUsagePlanRequest()
    }
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
    if request == nil {
        request = NewDescribeAPIDocDetailRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeAPIDocs(request *DescribeAPIDocsRequest) (response *DescribeAPIDocsResponse, err error) {
    if request == nil {
        request = NewDescribeAPIDocsRequest()
    }
    response = NewDescribeAPIDocsResponse()
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
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
    if request == nil {
        request = NewDescribeApiRequest()
    }
    response = NewDescribeApiResponse()
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
    if request == nil {
        request = NewDescribeApiEnvironmentStrategyRequest()
    }
    response = NewDescribeApiEnvironmentStrategyResponse()
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
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
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
// 当用户创建了多个密钥对时，可使用本接口查询一个或多个 API 密钥信息，本接口不会显示密钥 Key。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
func (c *Client) DescribeApiKeysStatus(request *DescribeApiKeysStatusRequest) (response *DescribeApiKeysStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeysStatusRequest()
    }
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
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeApiUsagePlan(request *DescribeApiUsagePlanRequest) (response *DescribeApiUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeApiUsagePlanRequest()
    }
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
func (c *Client) DescribeApisStatus(request *DescribeApisStatusRequest) (response *DescribeApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApisStatusRequest()
    }
    response = NewDescribeApisStatusResponse()
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
//  INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeIPStrategy(request *DescribeIPStrategyRequest) (response *DescribeIPStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"
func (c *Client) DescribeIPStrategyApisStatus(request *DescribeIPStrategyApisStatusRequest) (response *DescribeIPStrategyApisStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIPStrategyApisStatusRequest()
    }
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
    if request == nil {
        request = NewDescribeIPStrategysStatusRequest()
    }
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
func (c *Client) DescribeLogSearch(request *DescribeLogSearchRequest) (response *DescribeLogSearchResponse, err error) {
    if request == nil {
        request = NewDescribeLogSearchRequest()
    }
    response = NewDescribeLogSearchResponse()
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
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
func (c *Client) DescribePlugins(request *DescribePluginsRequest) (response *DescribePluginsResponse, err error) {
    if request == nil {
        request = NewDescribePluginsRequest()
    }
    response = NewDescribePluginsResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
    if request == nil {
        request = NewDescribeServiceRequest()
    }
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
    if request == nil {
        request = NewDescribeServiceEnvironmentListRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentReleaseHistory(request *DescribeServiceEnvironmentReleaseHistoryRequest) (response *DescribeServiceEnvironmentReleaseHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentReleaseHistoryRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceEnvironmentStrategy(request *DescribeServiceEnvironmentStrategyRequest) (response *DescribeServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServiceEnvironmentStrategyRequest()
    }
    response = NewDescribeServiceEnvironmentStrategyResponse()
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceReleaseVersion(request *DescribeServiceReleaseVersionRequest) (response *DescribeServiceReleaseVersionResponse, err error) {
    if request == nil {
        request = NewDescribeServiceReleaseVersionRequest()
    }
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
    if request == nil {
        request = NewDescribeServiceSubDomainMappingsRequest()
    }
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
    if request == nil {
        request = NewDescribeServiceSubDomainsRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) DescribeServiceUsagePlan(request *DescribeServiceUsagePlanRequest) (response *DescribeServiceUsagePlanResponse, err error) {
    if request == nil {
        request = NewDescribeServiceUsagePlanRequest()
    }
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
func (c *Client) DescribeServicesStatus(request *DescribeServicesStatusRequest) (response *DescribeServicesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeServicesStatusRequest()
    }
    response = NewDescribeServicesStatusResponse()
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
    if request == nil {
        request = NewDescribeUsagePlanRequest()
    }
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
    if request == nil {
        request = NewDescribeUsagePlanEnvironmentsRequest()
    }
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
    if request == nil {
        request = NewDescribeUsagePlanSecretIdsRequest()
    }
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
    if request == nil {
        request = NewDescribeUsagePlansStatusRequest()
    }
    response = NewDescribeUsagePlansStatusResponse()
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
    if request == nil {
        request = NewDisableApiKeyRequest()
    }
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
    if request == nil {
        request = NewEnableApiKeyRequest()
    }
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
    if request == nil {
        request = NewGenerateApiDocumentRequest()
    }
    response = NewGenerateApiDocumentResponse()
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
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
func (c *Client) ModifyAPIDoc(request *ModifyAPIDocRequest) (response *ModifyAPIDocResponse, err error) {
    if request == nil {
        request = NewModifyAPIDocRequest()
    }
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
//  FAILEDOPERATION_EBERROR = "FailedOperation.EbError"
//  FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"
//  INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"
//  INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"
//  UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"
//  UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
//  UNSUPPORTEDOPERATION_MODIFYPROTOCOL = "UnsupportedOperation.ModifyProtocol"
//  UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"
func (c *Client) ModifyApi(request *ModifyApiRequest) (response *ModifyApiResponse, err error) {
    if request == nil {
        request = NewModifyApiRequest()
    }
    response = NewModifyApiResponse()
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
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiEnvironmentStrategy(request *ModifyApiEnvironmentStrategyRequest) (response *ModifyApiEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyApiEnvironmentStrategyRequest()
    }
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
//  INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"
//  INVALIDPARAMETERVALUE_NOTHINGMODIFYFOROAUTH = "InvalidParameterValue.NothingModifyForOauth"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
func (c *Client) ModifyApiIncrement(request *ModifyApiIncrementRequest) (response *ModifyApiIncrementResponse, err error) {
    if request == nil {
        request = NewModifyApiIncrementRequest()
    }
    response = NewModifyApiIncrementResponse()
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
    if request == nil {
        request = NewModifyIPStrategyRequest()
    }
    response = NewModifyIPStrategyResponse()
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"
//  UNSUPPORTEDOPERATION_REDUCENETTYPES = "UnsupportedOperation.ReduceNetTypes"
func (c *Client) ModifyService(request *ModifyServiceRequest) (response *ModifyServiceResponse, err error) {
    if request == nil {
        request = NewModifyServiceRequest()
    }
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
//  INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"
func (c *Client) ModifyServiceEnvironmentStrategy(request *ModifyServiceEnvironmentStrategyRequest) (response *ModifyServiceEnvironmentStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServiceEnvironmentStrategyRequest()
    }
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
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SETCUSTOMPATHMAPPINGERROR = "FailedOperation.SetCustomPathMappingError"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  FAILEDOPERATION_UNKNOWNPROTOCOLTYPEERROR = "FailedOperation.UnknownProtocolTypeError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) ModifySubDomain(request *ModifySubDomainRequest) (response *ModifySubDomainResponse, err error) {
    if request == nil {
        request = NewModifySubDomainRequest()
    }
    response = NewModifySubDomainResponse()
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
    if request == nil {
        request = NewModifyUsagePlanRequest()
    }
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReleaseService(request *ReleaseServiceRequest) (response *ReleaseServiceResponse, err error) {
    if request == nil {
        request = NewReleaseServiceRequest()
    }
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
    if request == nil {
        request = NewResetAPIDocPasswordRequest()
    }
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
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindEnvironment(request *UnBindEnvironmentRequest) (response *UnBindEnvironmentResponse, err error) {
    if request == nil {
        request = NewUnBindEnvironmentRequest()
    }
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
//  FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"
//  RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"
func (c *Client) UnBindIPStrategy(request *UnBindIPStrategyRequest) (response *UnBindIPStrategyResponse, err error) {
    if request == nil {
        request = NewUnBindIPStrategyRequest()
    }
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
    if request == nil {
        request = NewUnBindSecretIdsRequest()
    }
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
//  FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"
//  FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"
//  UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"
func (c *Client) UnBindSubDomain(request *UnBindSubDomainRequest) (response *UnBindSubDomainResponse, err error) {
    if request == nil {
        request = NewUnBindSubDomainRequest()
    }
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"
//  LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnReleaseService(request *UnReleaseServiceRequest) (response *UnReleaseServiceResponse, err error) {
    if request == nil {
        request = NewUnReleaseServiceRequest()
    }
    response = NewUnReleaseServiceResponse()
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
//  RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"
//  UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"
func (c *Client) UpdateApiKey(request *UpdateApiKeyRequest) (response *UpdateApiKeyResponse, err error) {
    if request == nil {
        request = NewUpdateApiKeyRequest()
    }
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
//  INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"
//  INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
