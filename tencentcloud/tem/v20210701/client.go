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

package v20210701

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-07-01"

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


func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateApplication")
    
    
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATESERVICEERROR = "FailedOperation.CreateServiceError"
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_TAGINTERFACEERROR = "InternalError.TagInterfaceError"
//  INVALIDPARAMETERVALUE_INVALIDSERVICENAME = "InvalidParameterValue.InvalidServiceName"
//  INVALIDPARAMETERVALUE_PUBLICREPOTYPEPARAMETERERROR = "InvalidParameterValue.PublicRepoTypeParameterError"
//  INVALIDPARAMETERVALUE_SERVICELOWERCASE = "InvalidParameterValue.ServiceLowerCase"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  INVALIDPARAMETERVALUE_SERVICEREACHMAXIMUM = "InvalidParameterValue.ServiceReachMaximum"
//  INVALIDPARAMETERVALUE_TCRENTINSTANCENAMENOTVALID = "InvalidParameterValue.TcrEntInstanceNameNotValid"
//  MISSINGPARAMETER_TCRENTINSTANCENAMENULL = "MissingParameter.TcrEntInstanceNameNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATESERVICEERROR = "FailedOperation.CreateServiceError"
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_TAGINTERFACEERROR = "InternalError.TagInterfaceError"
//  INVALIDPARAMETERVALUE_INVALIDSERVICENAME = "InvalidParameterValue.InvalidServiceName"
//  INVALIDPARAMETERVALUE_PUBLICREPOTYPEPARAMETERERROR = "InvalidParameterValue.PublicRepoTypeParameterError"
//  INVALIDPARAMETERVALUE_SERVICELOWERCASE = "InvalidParameterValue.ServiceLowerCase"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  INVALIDPARAMETERVALUE_SERVICEREACHMAXIMUM = "InvalidParameterValue.ServiceReachMaximum"
//  INVALIDPARAMETERVALUE_TCRENTINSTANCENAMENOTVALID = "InvalidParameterValue.TcrEntInstanceNameNotValid"
//  MISSINGPARAMETER_TCRENTINSTANCENAMENULL = "MissingParameter.TcrEntInstanceNameNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationAutoscalerRequest() (request *CreateApplicationAutoscalerRequest) {
    request = &CreateApplicationAutoscalerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateApplicationAutoscaler")
    
    
    return
}

func NewCreateApplicationAutoscalerResponse() (response *CreateApplicationAutoscalerResponse) {
    response = &CreateApplicationAutoscalerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationAutoscaler
// 创建弹性伸缩策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_ATLEASTONESCALERRULESHOULDBEAPPLIED = "InvalidParameterValue.AtLeastOneScalerRuleShouldBeApplied"
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_CRONHPAREPLICASINVALID = "InvalidParameterValue.CronHpaReplicasInvalid"
//  INVALIDPARAMETERVALUE_HPAMETRICSINVALID = "InvalidParameterValue.HpaMetricsInvalid"
//  INVALIDPARAMETERVALUE_HPATHRESHOLDINVALID = "InvalidParameterValue.HpaThresholdInvalid"
//  INVALIDPARAMETERVALUE_INVALIDCRONSCALERPERIOD = "InvalidParameterValue.InvalidCronScalerPeriod"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SCALERNAMEDUPLICATED = "InvalidParameterValue.ScalerNameDuplicated"
//  MISSINGPARAMETER_AUTOSCALERNAMENULL = "MissingParameter.AutoScalerNameNull"
//  MISSINGPARAMETER_MINMAXNUMNULL = "MissingParameter.MinMaxNumNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateApplicationAutoscaler(request *CreateApplicationAutoscalerRequest) (response *CreateApplicationAutoscalerResponse, err error) {
    return c.CreateApplicationAutoscalerWithContext(context.Background(), request)
}

// CreateApplicationAutoscaler
// 创建弹性伸缩策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_ATLEASTONESCALERRULESHOULDBEAPPLIED = "InvalidParameterValue.AtLeastOneScalerRuleShouldBeApplied"
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_CRONHPAREPLICASINVALID = "InvalidParameterValue.CronHpaReplicasInvalid"
//  INVALIDPARAMETERVALUE_HPAMETRICSINVALID = "InvalidParameterValue.HpaMetricsInvalid"
//  INVALIDPARAMETERVALUE_HPATHRESHOLDINVALID = "InvalidParameterValue.HpaThresholdInvalid"
//  INVALIDPARAMETERVALUE_INVALIDCRONSCALERPERIOD = "InvalidParameterValue.InvalidCronScalerPeriod"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SCALERNAMEDUPLICATED = "InvalidParameterValue.ScalerNameDuplicated"
//  MISSINGPARAMETER_AUTOSCALERNAMENULL = "MissingParameter.AutoScalerNameNull"
//  MISSINGPARAMETER_MINMAXNUMNULL = "MissingParameter.MinMaxNumNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateApplicationAutoscalerWithContext(ctx context.Context, request *CreateApplicationAutoscalerRequest) (response *CreateApplicationAutoscalerResponse, err error) {
    if request == nil {
        request = NewCreateApplicationAutoscalerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationAutoscaler require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationAutoscalerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationServiceRequest() (request *CreateApplicationServiceRequest) {
    request = &CreateApplicationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateApplicationService")
    
    
    return
}

func NewCreateApplicationServiceResponse() (response *CreateApplicationServiceResponse) {
    response = &CreateApplicationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationService
// 新增访问方式
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETER_APPLICATIONACCESSSERVICEREACHMAXIMUM = "InvalidParameter.ApplicationAccessServiceReachMaximum"
//  INVALIDPARAMETER_LBSERVICECANNOTSUPPORTTCPUDPSAMETIME = "InvalidParameter.LBServiceCannotSupportTcpUdpSameTime"
//  INVALIDPARAMETER_MUSTPROVIDEPORTMAPPINGRULES = "InvalidParameter.MustProvidePortMappingRules"
//  INVALIDPARAMETER_SERVICENAMENOTVALID = "InvalidParameter.ServiceNameNotValid"
//  INVALIDPARAMETER_SERVICEUSERESERVESUFFIX = "InvalidParameter.ServiceUseReserveSuffix"
//  INVALIDPARAMETER_TOOMANYPORTMAPPINGRULES = "InvalidParameter.TooManyPortMappingRules"
//  INVALIDPARAMETERVALUE_APPLICATIONACCESSSERVICEREACHMAXIMUM = "InvalidParameterValue.ApplicationAccessServiceReachMaximum"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICEALREADYEXIST = "InvalidParameterValue.ApplicationServiceAlreadyExist"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICENOTFOUND = "InvalidParameterValue.ApplicationServiceNotFound"
//  INVALIDPARAMETERVALUE_CANNOTOVERWRITEOTHERAPPLICATIONSERVICE = "InvalidParameterValue.CannotOverWriteOtherApplicationService"
//  INVALIDPARAMETERVALUE_CANNOTUPDATESERVICEBYBOTHMETHODS = "InvalidParameterValue.CannotUpdateServiceByBothMethods"
//  INVALIDPARAMETERVALUE_INVALIDEKSSERVICETYPE = "InvalidParameterValue.InvalidEksServiceType"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_PORTDUPLICATEERROR = "InvalidParameterValue.PortDuplicateError"
//  INVALIDPARAMETERVALUE_PORTISRESERVED = "InvalidParameterValue.PortIsReserved"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateApplicationService(request *CreateApplicationServiceRequest) (response *CreateApplicationServiceResponse, err error) {
    return c.CreateApplicationServiceWithContext(context.Background(), request)
}

// CreateApplicationService
// 新增访问方式
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETER_APPLICATIONACCESSSERVICEREACHMAXIMUM = "InvalidParameter.ApplicationAccessServiceReachMaximum"
//  INVALIDPARAMETER_LBSERVICECANNOTSUPPORTTCPUDPSAMETIME = "InvalidParameter.LBServiceCannotSupportTcpUdpSameTime"
//  INVALIDPARAMETER_MUSTPROVIDEPORTMAPPINGRULES = "InvalidParameter.MustProvidePortMappingRules"
//  INVALIDPARAMETER_SERVICENAMENOTVALID = "InvalidParameter.ServiceNameNotValid"
//  INVALIDPARAMETER_SERVICEUSERESERVESUFFIX = "InvalidParameter.ServiceUseReserveSuffix"
//  INVALIDPARAMETER_TOOMANYPORTMAPPINGRULES = "InvalidParameter.TooManyPortMappingRules"
//  INVALIDPARAMETERVALUE_APPLICATIONACCESSSERVICEREACHMAXIMUM = "InvalidParameterValue.ApplicationAccessServiceReachMaximum"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICEALREADYEXIST = "InvalidParameterValue.ApplicationServiceAlreadyExist"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICENOTFOUND = "InvalidParameterValue.ApplicationServiceNotFound"
//  INVALIDPARAMETERVALUE_CANNOTOVERWRITEOTHERAPPLICATIONSERVICE = "InvalidParameterValue.CannotOverWriteOtherApplicationService"
//  INVALIDPARAMETERVALUE_CANNOTUPDATESERVICEBYBOTHMETHODS = "InvalidParameterValue.CannotUpdateServiceByBothMethods"
//  INVALIDPARAMETERVALUE_INVALIDEKSSERVICETYPE = "InvalidParameterValue.InvalidEksServiceType"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_PORTDUPLICATEERROR = "InvalidParameterValue.PortDuplicateError"
//  INVALIDPARAMETERVALUE_PORTISRESERVED = "InvalidParameterValue.PortIsReserved"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) CreateApplicationServiceWithContext(ctx context.Context, request *CreateApplicationServiceRequest) (response *CreateApplicationServiceResponse, err error) {
    if request == nil {
        request = NewCreateApplicationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigDataRequest() (request *CreateConfigDataRequest) {
    request = &CreateConfigDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateConfigData")
    
    
    return
}

func NewCreateConfigDataResponse() (response *CreateConfigDataResponse) {
    response = &CreateConfigDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConfigData
// 创建配置
//
// 可能返回的错误码:
//  INTERNALERROR_CREATECONFIGDATAERROR = "InternalError.CreateConfigDataError"
//  INVALIDPARAMETERVALUE_CONFIGDATAALREADYEXIST = "InvalidParameterValue.ConfigDataAlreadyExist"
//  INVALIDPARAMETERVALUE_CONFIGDATAINVALID = "InvalidParameterValue.ConfigDataInvalid"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateConfigData(request *CreateConfigDataRequest) (response *CreateConfigDataResponse, err error) {
    return c.CreateConfigDataWithContext(context.Background(), request)
}

// CreateConfigData
// 创建配置
//
// 可能返回的错误码:
//  INTERNALERROR_CREATECONFIGDATAERROR = "InternalError.CreateConfigDataError"
//  INVALIDPARAMETERVALUE_CONFIGDATAALREADYEXIST = "InvalidParameterValue.ConfigDataAlreadyExist"
//  INVALIDPARAMETERVALUE_CONFIGDATAINVALID = "InvalidParameterValue.ConfigDataInvalid"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateConfigDataWithContext(ctx context.Context, request *CreateConfigDataRequest) (response *CreateConfigDataResponse, err error) {
    if request == nil {
        request = NewCreateConfigDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigData require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigDataResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosTokenRequest() (request *CreateCosTokenRequest) {
    request = &CreateCosTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateCosToken")
    
    
    return
}

func NewCreateCosTokenResponse() (response *CreateCosTokenResponse) {
    response = &CreateCosTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCosToken
// 生成Cos临时秘钥
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCosToken(request *CreateCosTokenRequest) (response *CreateCosTokenResponse, err error) {
    return c.CreateCosTokenWithContext(context.Background(), request)
}

// CreateCosToken
// 生成Cos临时秘钥
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateCosTokenWithContext(ctx context.Context, request *CreateCosTokenRequest) (response *CreateCosTokenResponse, err error) {
    if request == nil {
        request = NewCreateCosTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
    request = &CreateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateEnvironment")
    
    
    return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
    response = &CreateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEnvironment
// 创建环境
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  INTERNALERROR_CREATEEKSCLUSTERERROR = "InternalError.CreateEksClusterError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETER_UNAUTHORIZEDORMISSINGROLE = "InvalidParameter.UnauthorizedOrMissingRole"
//  INVALIDPARAMETER_VPCOVERQUOTA = "InvalidParameter.VpcOverQuota"
//  INVALIDPARAMETERVALUE_DAILYCREATENAMESPACEREACHMAXIMUM = "InvalidParameterValue.DailyCreateNamespaceReachMaximum"
//  INVALIDPARAMETERVALUE_NAMESPACEDUPLICATEERROR = "InvalidParameterValue.NamespaceDuplicateError"
//  INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"
//  MISSINGPARAMETER_ENVIRONMENTNAMENULL = "MissingParameter.EnvironmentNameNull"
//  OPERATIONDENIED_BALANCENOTENOUGH = "OperationDenied.BalanceNotEnough"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    return c.CreateEnvironmentWithContext(context.Background(), request)
}

// CreateEnvironment
// 创建环境
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  INTERNALERROR_CREATEEKSCLUSTERERROR = "InternalError.CreateEksClusterError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETER_UNAUTHORIZEDORMISSINGROLE = "InvalidParameter.UnauthorizedOrMissingRole"
//  INVALIDPARAMETER_VPCOVERQUOTA = "InvalidParameter.VpcOverQuota"
//  INVALIDPARAMETERVALUE_DAILYCREATENAMESPACEREACHMAXIMUM = "InvalidParameterValue.DailyCreateNamespaceReachMaximum"
//  INVALIDPARAMETERVALUE_NAMESPACEDUPLICATEERROR = "InvalidParameterValue.NamespaceDuplicateError"
//  INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"
//  MISSINGPARAMETER_ENVIRONMENTNAMENULL = "MissingParameter.EnvironmentNameNull"
//  OPERATIONDENIED_BALANCENOTENOUGH = "OperationDenied.BalanceNotEnough"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogConfigRequest() (request *CreateLogConfigRequest) {
    request = &CreateLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateLogConfig")
    
    
    return
}

func NewCreateLogConfigResponse() (response *CreateLogConfigResponse) {
    response = &CreateLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLogConfig
// 创建日志收集配置
//
// 可能返回的错误码:
//  INTERNALERROR_CREATECONFIGDATAERROR = "InternalError.CreateConfigDataError"
//  INTERNALERROR_CREATELOGCONFIGERROR = "InternalError.CreateLogConfigError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_CONFIGDATAALREADYEXIST = "InvalidParameterValue.ConfigDataAlreadyExist"
//  INVALIDPARAMETERVALUE_LOGCONFIGALREADYEXIST = "InvalidParameterValue.LogConfigAlreadyExist"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateLogConfig(request *CreateLogConfigRequest) (response *CreateLogConfigResponse, err error) {
    return c.CreateLogConfigWithContext(context.Background(), request)
}

// CreateLogConfig
// 创建日志收集配置
//
// 可能返回的错误码:
//  INTERNALERROR_CREATECONFIGDATAERROR = "InternalError.CreateConfigDataError"
//  INTERNALERROR_CREATELOGCONFIGERROR = "InternalError.CreateLogConfigError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_CONFIGDATAALREADYEXIST = "InvalidParameterValue.ConfigDataAlreadyExist"
//  INVALIDPARAMETERVALUE_LOGCONFIGALREADYEXIST = "InvalidParameterValue.LogConfigAlreadyExist"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateLogConfigWithContext(ctx context.Context, request *CreateLogConfigRequest) (response *CreateLogConfigResponse, err error) {
    if request == nil {
        request = NewCreateLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
    request = &CreateResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "CreateResource")
    
    
    return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
    response = &CreateResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateResource
// 绑定云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"
//  INVALIDPARAMETERVALUE_NAMESPACERESOURCEREACHMAXIMUM = "InvalidParameterValue.NamespaceResourceReachMaximum"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_VPCINVALID = "InvalidParameterValue.VpcInvalid"
//  RESOURCEINUSE_RESOURCEALREADYUSED = "ResourceInUse.ResourceAlreadyUsed"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    return c.CreateResourceWithContext(context.Background(), request)
}

// CreateResource
// 绑定云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_NAMESPACEREACHMAXIMUM = "InvalidParameterValue.NamespaceReachMaximum"
//  INVALIDPARAMETERVALUE_NAMESPACERESOURCEREACHMAXIMUM = "InvalidParameterValue.NamespaceResourceReachMaximum"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_VPCINVALID = "InvalidParameterValue.VpcInvalid"
//  RESOURCEINUSE_RESOURCEALREADYUSED = "ResourceInUse.ResourceAlreadyUsed"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DeleteApplication")
    
    
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplication
// 服务删除
//
//   - 停止当前运行服务
//
//   - 删除服务相关资源
//
//   - 删除服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DELETESERVICEERROR = "FailedOperation.DeleteServiceError"
//  FAILEDOPERATION_DESCRIBEINGRESSLISTERROR = "FailedOperation.DescribeIngressListError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"
//  INTERNALERROR_DESCRIBEINGRESSLISTERROR = "InternalError.DescribeIngressListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCEINUSE_RESOURCEALREADYLOCKED = "ResourceInUse.ResourceAlreadyLocked"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  RESOURCEUNAVAILABLE_APPLICATIONNOTDELETABLE = "ResourceUnavailable.ApplicationNotDeletable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    return c.DeleteApplicationWithContext(context.Background(), request)
}

// DeleteApplication
// 服务删除
//
//   - 停止当前运行服务
//
//   - 删除服务相关资源
//
//   - 删除服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DELETESERVICEERROR = "FailedOperation.DeleteServiceError"
//  FAILEDOPERATION_DESCRIBEINGRESSLISTERROR = "FailedOperation.DescribeIngressListError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DELETESERVICEERROR = "InternalError.DeleteServiceError"
//  INTERNALERROR_DESCRIBEINGRESSLISTERROR = "InternalError.DescribeIngressListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEFOUNDRUNNINGVERSION = "InvalidParameterValue.ServiceFoundRunningVersion"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_VERSIONROUTERATENOTZERO = "InvalidParameterValue.VersionRouteRateNotZero"
//  RESOURCEINUSE_RESOURCEALREADYLOCKED = "ResourceInUse.ResourceAlreadyLocked"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  RESOURCEUNAVAILABLE_APPLICATIONNOTDELETABLE = "ResourceUnavailable.ApplicationNotDeletable"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationAutoscalerRequest() (request *DeleteApplicationAutoscalerRequest) {
    request = &DeleteApplicationAutoscalerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DeleteApplicationAutoscaler")
    
    
    return
}

func NewDeleteApplicationAutoscalerResponse() (response *DeleteApplicationAutoscalerResponse) {
    response = &DeleteApplicationAutoscalerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationAutoscaler
// 删除应用弹性策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_DISABLESCALERBEFOREDELETE = "InvalidParameterValue.DisableScalerBeforeDelete"
//  MISSINGPARAMETER_SCALERIDNULL = "MissingParameter.ScalerIdNull"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplicationAutoscaler(request *DeleteApplicationAutoscalerRequest) (response *DeleteApplicationAutoscalerResponse, err error) {
    return c.DeleteApplicationAutoscalerWithContext(context.Background(), request)
}

// DeleteApplicationAutoscaler
// 删除应用弹性策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_DISABLESCALERBEFOREDELETE = "InvalidParameterValue.DisableScalerBeforeDelete"
//  MISSINGPARAMETER_SCALERIDNULL = "MissingParameter.ScalerIdNull"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplicationAutoscalerWithContext(ctx context.Context, request *DeleteApplicationAutoscalerRequest) (response *DeleteApplicationAutoscalerResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationAutoscalerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationAutoscaler require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationAutoscalerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationServiceRequest() (request *DeleteApplicationServiceRequest) {
    request = &DeleteApplicationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DeleteApplicationService")
    
    
    return
}

func NewDeleteApplicationServiceResponse() (response *DeleteApplicationServiceResponse) {
    response = &DeleteApplicationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationService
// 删除一条访问方式
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICENOTFOUND = "InvalidParameterValue.ApplicationServiceNotFound"
//  INVALIDPARAMETERVALUE_CANNOTUPDATESERVICEBYBOTHMETHODS = "InvalidParameterValue.CannotUpdateServiceByBothMethods"
//  INVALIDPARAMETERVALUE_INVALIDEKSSERVICETYPE = "InvalidParameterValue.InvalidEksServiceType"
//  INVALIDPARAMETERVALUE_PORTISRESERVED = "InvalidParameterValue.PortIsReserved"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplicationService(request *DeleteApplicationServiceRequest) (response *DeleteApplicationServiceResponse, err error) {
    return c.DeleteApplicationServiceWithContext(context.Background(), request)
}

// DeleteApplicationService
// 删除一条访问方式
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICENOTFOUND = "InvalidParameterValue.ApplicationServiceNotFound"
//  INVALIDPARAMETERVALUE_CANNOTUPDATESERVICEBYBOTHMETHODS = "InvalidParameterValue.CannotUpdateServiceByBothMethods"
//  INVALIDPARAMETERVALUE_INVALIDEKSSERVICETYPE = "InvalidParameterValue.InvalidEksServiceType"
//  INVALIDPARAMETERVALUE_PORTISRESERVED = "InvalidParameterValue.PortIsReserved"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApplicationServiceWithContext(ctx context.Context, request *DeleteApplicationServiceRequest) (response *DeleteApplicationServiceResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIngressRequest() (request *DeleteIngressRequest) {
    request = &DeleteIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DeleteIngress")
    
    
    return
}

func NewDeleteIngressResponse() (response *DeleteIngressResponse) {
    response = &DeleteIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIngress
// 删除 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEINGRESSERROR = "InternalError.DeleteIngressError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteIngress(request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
    return c.DeleteIngressWithContext(context.Background(), request)
}

// DeleteIngress
// 删除 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEINGRESSERROR = "InternalError.DeleteIngressError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteIngressWithContext(ctx context.Context, request *DeleteIngressRequest) (response *DeleteIngressResponse, err error) {
    if request == nil {
        request = NewDeleteIngressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIngress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIngressResponse()
    err = c.Send(request, response)
    return
}

func NewDeployApplicationRequest() (request *DeployApplicationRequest) {
    request = &DeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DeployApplication")
    
    
    return
}

func NewDeployApplicationResponse() (response *DeployApplicationResponse) {
    response = &DeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeployApplication
// 应用部署
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONREADTIMEOUT = "FailedOperation.ActionReadTimeout"
//  FAILEDOPERATION_CREATESERVICEERROR = "FailedOperation.CreateServiceError"
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DESCRIBESERVICELISTERROR = "FailedOperation.DescribeServiceListError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_APMNOTBIND = "InvalidParameterValue.ApmNotBind"
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_CANNOTOVERWRITEOTHERAPPLICATIONSERVICE = "InvalidParameterValue.CannotOverWriteOtherApplicationService"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_INVALIDENVNAME = "InvalidParameterValue.InvalidEnvName"
//  INVALIDPARAMETERVALUE_INVALIDENVVALUE = "InvalidParameterValue.InvalidEnvValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPATH = "InvalidParameterValue.InvalidMountPath"
//  INVALIDPARAMETERVALUE_INVALIDTENANTINFO = "InvalidParameterValue.InvalidTenantInfo"
//  INVALIDPARAMETERVALUE_JDKVERSIONREQUIRED = "InvalidParameterValue.JdkVersionRequired"
//  INVALIDPARAMETERVALUE_MUSTPROVIDEPORTMAPPINGRULES = "InvalidParameterValue.MustProvidePortMappingRules"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_OSNOTSUPPORT = "InvalidParameterValue.OsNotSupport"
//  INVALIDPARAMETERVALUE_POSTSTARTNOTVALID = "InvalidParameterValue.PostStartNotValid"
//  INVALIDPARAMETERVALUE_REGISTRYNOTBIND = "InvalidParameterValue.RegistryNotBind"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  INVALIDPARAMETERVALUE_SERVICENOTBELONGTOAPPID = "InvalidParameterValue.ServiceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEPODREACHMAXIMUM = "InvalidParameterValue.ServicePodReachMaximum"
//  INVALIDPARAMETERVALUE_SERVICEUSERESERVESUFFIX = "InvalidParameterValue.ServiceUseReserveSuffix"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"
//  INVALIDPARAMETERVALUE_VERSIONLENGTHLIMIT = "InvalidParameterValue.VersionLengthLimit"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  MISSINGPARAMETER_DEPLOYMODENULL = "MissingParameter.DeployModeNull"
//  MISSINGPARAMETER_DEPLOYVERSIONNULL = "MissingParameter.DeployVersionNull"
//  MISSINGPARAMETER_IMGREPONULL = "MissingParameter.ImgRepoNull"
//  MISSINGPARAMETER_LOGSETORTOPICNULL = "MissingParameter.LogsetOrTopicNull"
//  MISSINGPARAMETER_PKGNAMENULL = "MissingParameter.PkgNameNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  MISSINGPARAMETER_SVCREPONOTREADY = "MissingParameter.SvcRepoNotReady"
//  MISSINGPARAMETER_VPCSERVICESUBNETNULL = "MissingParameter.VpcServiceSubnetNull"
//  RESOURCEINUSE_SERVICEDEPLOYING = "ResourceInUse.ServiceDeploying"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
//  UNAUTHORIZEDOPERATION_MISSINGEKSLOGROLE = "UnauthorizedOperation.MissingEksLogRole"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeployApplication(request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
    return c.DeployApplicationWithContext(context.Background(), request)
}

// DeployApplication
// 应用部署
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONREADTIMEOUT = "FailedOperation.ActionReadTimeout"
//  FAILEDOPERATION_CREATESERVICEERROR = "FailedOperation.CreateServiceError"
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DESCRIBESERVICELISTERROR = "FailedOperation.DescribeServiceListError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_APMNOTBIND = "InvalidParameterValue.ApmNotBind"
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_CANNOTOVERWRITEOTHERAPPLICATIONSERVICE = "InvalidParameterValue.CannotOverWriteOtherApplicationService"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_INVALIDENVNAME = "InvalidParameterValue.InvalidEnvName"
//  INVALIDPARAMETERVALUE_INVALIDENVVALUE = "InvalidParameterValue.InvalidEnvValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPATH = "InvalidParameterValue.InvalidMountPath"
//  INVALIDPARAMETERVALUE_INVALIDTENANTINFO = "InvalidParameterValue.InvalidTenantInfo"
//  INVALIDPARAMETERVALUE_JDKVERSIONREQUIRED = "InvalidParameterValue.JdkVersionRequired"
//  INVALIDPARAMETERVALUE_MUSTPROVIDEPORTMAPPINGRULES = "InvalidParameterValue.MustProvidePortMappingRules"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_OSNOTSUPPORT = "InvalidParameterValue.OsNotSupport"
//  INVALIDPARAMETERVALUE_POSTSTARTNOTVALID = "InvalidParameterValue.PostStartNotValid"
//  INVALIDPARAMETERVALUE_REGISTRYNOTBIND = "InvalidParameterValue.RegistryNotBind"
//  INVALIDPARAMETERVALUE_SERVICENAMEDUPLICATEERROR = "InvalidParameterValue.ServiceNameDuplicateError"
//  INVALIDPARAMETERVALUE_SERVICENOTBELONGTOAPPID = "InvalidParameterValue.ServiceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEPODREACHMAXIMUM = "InvalidParameterValue.ServicePodReachMaximum"
//  INVALIDPARAMETERVALUE_SERVICEUSERESERVESUFFIX = "InvalidParameterValue.ServiceUseReserveSuffix"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_TRAITSTRACINGNOTSUPPORTED = "InvalidParameterValue.TraitsTracingNotSupported"
//  INVALIDPARAMETERVALUE_VERSIONLENGTHLIMIT = "InvalidParameterValue.VersionLengthLimit"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  MISSINGPARAMETER_DEPLOYMODENULL = "MissingParameter.DeployModeNull"
//  MISSINGPARAMETER_DEPLOYVERSIONNULL = "MissingParameter.DeployVersionNull"
//  MISSINGPARAMETER_IMGREPONULL = "MissingParameter.ImgRepoNull"
//  MISSINGPARAMETER_LOGSETORTOPICNULL = "MissingParameter.LogsetOrTopicNull"
//  MISSINGPARAMETER_PKGNAMENULL = "MissingParameter.PkgNameNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  MISSINGPARAMETER_SVCREPONOTREADY = "MissingParameter.SvcRepoNotReady"
//  MISSINGPARAMETER_VPCSERVICESUBNETNULL = "MissingParameter.VpcServiceSubnetNull"
//  RESOURCEINUSE_SERVICEDEPLOYING = "ResourceInUse.ServiceDeploying"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
//  UNAUTHORIZEDOPERATION_MISSINGEKSLOGROLE = "UnauthorizedOperation.MissingEksLogRole"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeployApplicationWithContext(ctx context.Context, request *DeployApplicationRequest) (response *DeployApplicationResponse, err error) {
    if request == nil {
        request = NewDeployApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationAutoscalerListRequest() (request *DescribeApplicationAutoscalerListRequest) {
    request = &DescribeApplicationAutoscalerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeApplicationAutoscalerList")
    
    
    return
}

func NewDescribeApplicationAutoscalerListResponse() (response *DescribeApplicationAutoscalerListResponse) {
    response = &DescribeApplicationAutoscalerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationAutoscalerList
// 获取应用弹性策略组合
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationAutoscalerList(request *DescribeApplicationAutoscalerListRequest) (response *DescribeApplicationAutoscalerListResponse, err error) {
    return c.DescribeApplicationAutoscalerListWithContext(context.Background(), request)
}

// DescribeApplicationAutoscalerList
// 获取应用弹性策略组合
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationAutoscalerListWithContext(ctx context.Context, request *DescribeApplicationAutoscalerListRequest) (response *DescribeApplicationAutoscalerListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationAutoscalerListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationAutoscalerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationAutoscalerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationInfoRequest() (request *DescribeApplicationInfoRequest) {
    request = &DescribeApplicationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeApplicationInfo")
    
    
    return
}

func NewDescribeApplicationInfoResponse() (response *DescribeApplicationInfoResponse) {
    response = &DescribeApplicationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationInfo
// 服务基本信息查看
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONREADTIMEOUT = "FailedOperation.ActionReadTimeout"
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DESCRIBERUNPODLISTERROR = "FailedOperation.DescribeRunPodListError"
//  FAILEDOPERATION_DESCRIBESERVICEERROR = "FailedOperation.DescribeServiceError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  INTERNALERROR_DESCRIBESERVICEERROR = "InternalError.DescribeServiceError"
//  INVALIDPARAMETERVALUE_SERVICENOTBELONGTOAPPID = "InvalidParameterValue.ServiceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationInfo(request *DescribeApplicationInfoRequest) (response *DescribeApplicationInfoResponse, err error) {
    return c.DescribeApplicationInfoWithContext(context.Background(), request)
}

// DescribeApplicationInfo
// 服务基本信息查看
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONREADTIMEOUT = "FailedOperation.ActionReadTimeout"
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DESCRIBERUNPODLISTERROR = "FailedOperation.DescribeRunPodListError"
//  FAILEDOPERATION_DESCRIBESERVICEERROR = "FailedOperation.DescribeServiceError"
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  INTERNALERROR_DESCRIBESERVICEERROR = "InternalError.DescribeServiceError"
//  INVALIDPARAMETERVALUE_SERVICENOTBELONGTOAPPID = "InvalidParameterValue.ServiceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationInfoWithContext(ctx context.Context, request *DescribeApplicationInfoRequest) (response *DescribeApplicationInfoResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationPodsRequest() (request *DescribeApplicationPodsRequest) {
    request = &DescribeApplicationPodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeApplicationPods")
    
    
    return
}

func NewDescribeApplicationPodsResponse() (response *DescribeApplicationPodsResponse) {
    response = &DescribeApplicationPodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationPods
// 获取应用实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DESCRIBERUNPODLISTERROR = "FailedOperation.DescribeRunPodListError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationPods(request *DescribeApplicationPodsRequest) (response *DescribeApplicationPodsResponse, err error) {
    return c.DescribeApplicationPodsWithContext(context.Background(), request)
}

// DescribeApplicationPods
// 获取应用实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEFAULTINTERNALERROR = "FailedOperation.DefaultInternalError"
//  FAILEDOPERATION_DESCRIBERUNPODLISTERROR = "FailedOperation.DescribeRunPodListError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationPodsWithContext(ctx context.Context, request *DescribeApplicationPodsRequest) (response *DescribeApplicationPodsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationPodsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationPods require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationPodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationServiceListRequest() (request *DescribeApplicationServiceListRequest) {
    request = &DescribeApplicationServiceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeApplicationServiceList")
    
    
    return
}

func NewDescribeApplicationServiceListResponse() (response *DescribeApplicationServiceListResponse) {
    response = &DescribeApplicationServiceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationServiceList
// 查询应用访问方式列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeApplicationServiceList(request *DescribeApplicationServiceListRequest) (response *DescribeApplicationServiceListResponse, err error) {
    return c.DescribeApplicationServiceListWithContext(context.Background(), request)
}

// DescribeApplicationServiceList
// 查询应用访问方式列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_INTERFACENOTFOUND = "ResourceNotFound.InterfaceNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeApplicationServiceListWithContext(ctx context.Context, request *DescribeApplicationServiceListRequest) (response *DescribeApplicationServiceListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationServiceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationServiceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationServiceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeApplications")
    
    
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplications
// 获取运行服务列表
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    return c.DescribeApplicationsWithContext(context.Background(), request)
}

// DescribeApplications
// 获取运行服务列表
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeApplicationsWithContext(ctx context.Context, request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsStatusRequest() (request *DescribeApplicationsStatusRequest) {
    request = &DescribeApplicationsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeApplicationsStatus")
    
    
    return
}

func NewDescribeApplicationsStatusResponse() (response *DescribeApplicationsStatusResponse) {
    response = &DescribeApplicationsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationsStatus
// 单环境下所有应用状态查看
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationsStatus(request *DescribeApplicationsStatusRequest) (response *DescribeApplicationsStatusResponse, err error) {
    return c.DescribeApplicationsStatusWithContext(context.Background(), request)
}

// DescribeApplicationsStatus
// 单环境下所有应用状态查看
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApplicationsStatusWithContext(ctx context.Context, request *DescribeApplicationsStatusRequest) (response *DescribeApplicationsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationsStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigDataRequest() (request *DescribeConfigDataRequest) {
    request = &DescribeConfigDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeConfigData")
    
    
    return
}

func NewDescribeConfigDataResponse() (response *DescribeConfigDataResponse) {
    response = &DescribeConfigDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigData
// 查询配置详情
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBECONFIGDATAERROR = "InternalError.DescribeConfigDataError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_CONFIGDATANOTFOUND = "ResourceNotFound.ConfigDataNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeConfigData(request *DescribeConfigDataRequest) (response *DescribeConfigDataResponse, err error) {
    return c.DescribeConfigDataWithContext(context.Background(), request)
}

// DescribeConfigData
// 查询配置详情
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBECONFIGDATAERROR = "InternalError.DescribeConfigDataError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_CONFIGDATANOTFOUND = "ResourceNotFound.ConfigDataNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeConfigDataWithContext(ctx context.Context, request *DescribeConfigDataRequest) (response *DescribeConfigDataResponse, err error) {
    if request == nil {
        request = NewDescribeConfigDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigDataListRequest() (request *DescribeConfigDataListRequest) {
    request = &DescribeConfigDataListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeConfigDataList")
    
    
    return
}

func NewDescribeConfigDataListResponse() (response *DescribeConfigDataListResponse) {
    response = &DescribeConfigDataListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigDataList
// 查询配置列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBECONFIGDATALISTERROR = "InternalError.DescribeConfigDataListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeConfigDataList(request *DescribeConfigDataListRequest) (response *DescribeConfigDataListResponse, err error) {
    return c.DescribeConfigDataListWithContext(context.Background(), request)
}

// DescribeConfigDataList
// 查询配置列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBECONFIGDATALISTERROR = "InternalError.DescribeConfigDataListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeConfigDataListWithContext(ctx context.Context, request *DescribeConfigDataListRequest) (response *DescribeConfigDataListResponse, err error) {
    if request == nil {
        request = NewDescribeConfigDataListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigDataList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigDataListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployApplicationDetailRequest() (request *DescribeDeployApplicationDetailRequest) {
    request = &DescribeDeployApplicationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeDeployApplicationDetail")
    
    
    return
}

func NewDescribeDeployApplicationDetailResponse() (response *DescribeDeployApplicationDetailResponse) {
    response = &DescribeDeployApplicationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeployApplicationDetail
// 获取分批发布详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBERUNPODLISTERROR = "FailedOperation.DescribeRunPodListError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeDeployApplicationDetail(request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    return c.DescribeDeployApplicationDetailWithContext(context.Background(), request)
}

// DescribeDeployApplicationDetail
// 获取分批发布详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBERUNPODLISTERROR = "FailedOperation.DescribeRunPodListError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBERUNPODLISTERROR = "InternalError.DescribeRunPodListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCENOTFOUND_VERSIONSERVICENOTFOUND = "ResourceNotFound.VersionServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeDeployApplicationDetailWithContext(ctx context.Context, request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDeployApplicationDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployApplicationDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeployApplicationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentRequest() (request *DescribeEnvironmentRequest) {
    request = &DescribeEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeEnvironment")
    
    
    return
}

func NewDescribeEnvironmentResponse() (response *DescribeEnvironmentResponse) {
    response = &DescribeEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvironment
// 获取环境基础信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_NAMESPACENOTFOUND = "InvalidParameterValue.NamespaceNotFound"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEnvironment(request *DescribeEnvironmentRequest) (response *DescribeEnvironmentResponse, err error) {
    return c.DescribeEnvironmentWithContext(context.Background(), request)
}

// DescribeEnvironment
// 获取环境基础信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_NAMESPACENOTFOUND = "InvalidParameterValue.NamespaceNotFound"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEnvironmentWithContext(ctx context.Context, request *DescribeEnvironmentRequest) (response *DescribeEnvironmentResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentStatusRequest() (request *DescribeEnvironmentStatusRequest) {
    request = &DescribeEnvironmentStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeEnvironmentStatus")
    
    
    return
}

func NewDescribeEnvironmentStatusResponse() (response *DescribeEnvironmentStatusResponse) {
    response = &DescribeEnvironmentStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvironmentStatus
// 获取环境状态
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEnvironmentStatus(request *DescribeEnvironmentStatusRequest) (response *DescribeEnvironmentStatusResponse, err error) {
    return c.DescribeEnvironmentStatusWithContext(context.Background(), request)
}

// DescribeEnvironmentStatus
// 获取环境状态
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEnvironmentStatusWithContext(ctx context.Context, request *DescribeEnvironmentStatusRequest) (response *DescribeEnvironmentStatusResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironmentStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvironments
// 获取环境列表
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_INVALIDTENANTINFO = "InvalidParameterValue.InvalidTenantInfo"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// 获取环境列表
//
// 可能返回的错误码:
//  INTERNALERROR_ACTIONREADTIMEOUT = "InternalError.ActionReadTimeout"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_INVALIDTENANTINFO = "InvalidParameterValue.InvalidTenantInfo"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_NAMESPACENOTFOUND = "ResourceNotFound.NamespaceNotFound"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIngressRequest() (request *DescribeIngressRequest) {
    request = &DescribeIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeIngress")
    
    
    return
}

func NewDescribeIngressResponse() (response *DescribeIngressResponse) {
    response = &DescribeIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIngress
// 查询 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBEINGRESSERROR = "InternalError.DescribeIngressError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeIngress(request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    return c.DescribeIngressWithContext(context.Background(), request)
}

// DescribeIngress
// 查询 Ingress 规则
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBEINGRESSERROR = "InternalError.DescribeIngressError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeIngressWithContext(ctx context.Context, request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
    if request == nil {
        request = NewDescribeIngressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIngress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIngressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIngressesRequest() (request *DescribeIngressesRequest) {
    request = &DescribeIngressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeIngresses")
    
    
    return
}

func NewDescribeIngressesResponse() (response *DescribeIngressesResponse) {
    response = &DescribeIngressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIngresses
// 查询 Ingress 规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBESERVICELISTERROR = "FailedOperation.DescribeServiceListError"
//  INTERNALERROR_DESCRIBESERVICELISTERROR = "InternalError.DescribeServiceListError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeIngresses(request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    return c.DescribeIngressesWithContext(context.Background(), request)
}

// DescribeIngresses
// 查询 Ingress 规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBESERVICELISTERROR = "FailedOperation.DescribeServiceListError"
//  INTERNALERROR_DESCRIBESERVICELISTERROR = "InternalError.DescribeServiceListError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeIngressesWithContext(ctx context.Context, request *DescribeIngressesRequest) (response *DescribeIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeIngressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIngresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIngressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogConfigRequest() (request *DescribeLogConfigRequest) {
    request = &DescribeLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeLogConfig")
    
    
    return
}

func NewDescribeLogConfigResponse() (response *DescribeLogConfigResponse) {
    response = &DescribeLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogConfig
// 查询日志收集配置详情
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBELOGCONFIGERROR = "InternalError.DescribeLogConfigError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_LOGCONFIGNOTFOUND = "ResourceNotFound.LogConfigNotFound"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeLogConfig(request *DescribeLogConfigRequest) (response *DescribeLogConfigResponse, err error) {
    return c.DescribeLogConfigWithContext(context.Background(), request)
}

// DescribeLogConfig
// 查询日志收集配置详情
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBELOGCONFIGERROR = "InternalError.DescribeLogConfigError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_LOGCONFIGNOTFOUND = "ResourceNotFound.LogConfigNotFound"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeLogConfigWithContext(ctx context.Context, request *DescribeLogConfigRequest) (response *DescribeLogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePagedLogConfigListRequest() (request *DescribePagedLogConfigListRequest) {
    request = &DescribePagedLogConfigListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribePagedLogConfigList")
    
    
    return
}

func NewDescribePagedLogConfigListResponse() (response *DescribePagedLogConfigListResponse) {
    response = &DescribePagedLogConfigListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePagedLogConfigList
// 查询分页的日志收集配置列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBECONFIGDATALISTERROR = "InternalError.DescribeConfigDataListError"
//  INTERNALERROR_DESCRIBELOGCONFIGLISTERROR = "InternalError.DescribeLogConfigListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribePagedLogConfigList(request *DescribePagedLogConfigListRequest) (response *DescribePagedLogConfigListResponse, err error) {
    return c.DescribePagedLogConfigListWithContext(context.Background(), request)
}

// DescribePagedLogConfigList
// 查询分页的日志收集配置列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DESCRIBECONFIGDATALISTERROR = "InternalError.DescribeConfigDataListError"
//  INTERNALERROR_DESCRIBELOGCONFIGLISTERROR = "InternalError.DescribeLogConfigListError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribePagedLogConfigListWithContext(ctx context.Context, request *DescribePagedLogConfigListRequest) (response *DescribePagedLogConfigListResponse, err error) {
    if request == nil {
        request = NewDescribePagedLogConfigListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePagedLogConfigList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePagedLogConfigListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelatedIngressesRequest() (request *DescribeRelatedIngressesRequest) {
    request = &DescribeRelatedIngressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DescribeRelatedIngresses")
    
    
    return
}

func NewDescribeRelatedIngressesResponse() (response *DescribeRelatedIngressesResponse) {
    response = &DescribeRelatedIngressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRelatedIngresses
// 查询应用关联的 Ingress 规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESERVICEINGRESSERROR = "InternalError.DescribeServiceIngressError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeRelatedIngresses(request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    return c.DescribeRelatedIngressesWithContext(context.Background(), request)
}

// DescribeRelatedIngresses
// 查询应用关联的 Ingress 规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBESERVICEINGRESSERROR = "InternalError.DescribeServiceIngressError"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SERVICEIDNULL = "MissingParameter.ServiceIdNull"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) DescribeRelatedIngressesWithContext(ctx context.Context, request *DescribeRelatedIngressesRequest) (response *DescribeRelatedIngressesResponse, err error) {
    if request == nil {
        request = NewDescribeRelatedIngressesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelatedIngresses require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRelatedIngressesResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyConfigDataRequest() (request *DestroyConfigDataRequest) {
    request = &DestroyConfigDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DestroyConfigData")
    
    
    return
}

func NewDestroyConfigDataResponse() (response *DestroyConfigDataResponse) {
    response = &DestroyConfigDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyConfigData
// 销毁配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DestroyConfigData(request *DestroyConfigDataRequest) (response *DestroyConfigDataResponse, err error) {
    return c.DestroyConfigDataWithContext(context.Background(), request)
}

// DestroyConfigData
// 销毁配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DestroyConfigDataWithContext(ctx context.Context, request *DestroyConfigDataRequest) (response *DestroyConfigDataResponse, err error) {
    if request == nil {
        request = NewDestroyConfigDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyConfigData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyConfigDataResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyEnvironmentRequest() (request *DestroyEnvironmentRequest) {
    request = &DestroyEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DestroyEnvironment")
    
    
    return
}

func NewDestroyEnvironmentResponse() (response *DestroyEnvironmentResponse) {
    response = &DestroyEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyEnvironment
// 销毁环境
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCEINUSE_ENVIRONMENTALREADYLOCKED = "ResourceInUse.EnvironmentAlreadyLocked"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DestroyEnvironment(request *DestroyEnvironmentRequest) (response *DestroyEnvironmentResponse, err error) {
    return c.DestroyEnvironmentWithContext(context.Background(), request)
}

// DestroyEnvironment
// 销毁环境
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCEINUSE_ENVIRONMENTALREADYLOCKED = "ResourceInUse.EnvironmentAlreadyLocked"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DestroyEnvironmentWithContext(ctx context.Context, request *DestroyEnvironmentRequest) (response *DestroyEnvironmentResponse, err error) {
    if request == nil {
        request = NewDestroyEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyLogConfigRequest() (request *DestroyLogConfigRequest) {
    request = &DestroyLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DestroyLogConfig")
    
    
    return
}

func NewDestroyLogConfigResponse() (response *DestroyLogConfigResponse) {
    response = &DestroyLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyLogConfig
// 销毁日志收集配置
//
// 可能返回的错误码:
//  INTERNALERROR_DELETELOGCONFIGERROR = "InternalError.DeleteLogConfigError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DestroyLogConfig(request *DestroyLogConfigRequest) (response *DestroyLogConfigResponse, err error) {
    return c.DestroyLogConfigWithContext(context.Background(), request)
}

// DestroyLogConfig
// 销毁日志收集配置
//
// 可能返回的错误码:
//  INTERNALERROR_DELETELOGCONFIGERROR = "InternalError.DeleteLogConfigError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DestroyLogConfigWithContext(ctx context.Context, request *DestroyLogConfigRequest) (response *DestroyLogConfigResponse, err error) {
    if request == nil {
        request = NewDestroyLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDisableApplicationAutoscalerRequest() (request *DisableApplicationAutoscalerRequest) {
    request = &DisableApplicationAutoscalerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "DisableApplicationAutoscaler")
    
    
    return
}

func NewDisableApplicationAutoscalerResponse() (response *DisableApplicationAutoscalerResponse) {
    response = &DisableApplicationAutoscalerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableApplicationAutoscaler
// 关闭应用弹性策略组合
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DisableApplicationAutoscaler(request *DisableApplicationAutoscalerRequest) (response *DisableApplicationAutoscalerResponse, err error) {
    return c.DisableApplicationAutoscalerWithContext(context.Background(), request)
}

// DisableApplicationAutoscaler
// 关闭应用弹性策略组合
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DisableApplicationAutoscalerWithContext(ctx context.Context, request *DisableApplicationAutoscalerRequest) (response *DisableApplicationAutoscalerResponse, err error) {
    if request == nil {
        request = NewDisableApplicationAutoscalerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableApplicationAutoscaler require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableApplicationAutoscalerResponse()
    err = c.Send(request, response)
    return
}

func NewEnableApplicationAutoscalerRequest() (request *EnableApplicationAutoscalerRequest) {
    request = &EnableApplicationAutoscalerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "EnableApplicationAutoscaler")
    
    
    return
}

func NewEnableApplicationAutoscalerResponse() (response *EnableApplicationAutoscalerResponse) {
    response = &EnableApplicationAutoscalerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableApplicationAutoscaler
// 启用应用弹性策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) EnableApplicationAutoscaler(request *EnableApplicationAutoscalerRequest) (response *EnableApplicationAutoscalerResponse, err error) {
    return c.EnableApplicationAutoscalerWithContext(context.Background(), request)
}

// EnableApplicationAutoscaler
// 启用应用弹性策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) EnableApplicationAutoscalerWithContext(ctx context.Context, request *EnableApplicationAutoscalerRequest) (response *EnableApplicationAutoscalerResponse, err error) {
    if request == nil {
        request = NewEnableApplicationAutoscalerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableApplicationAutoscaler require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableApplicationAutoscalerResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateApplicationPackageDownloadUrlRequest() (request *GenerateApplicationPackageDownloadUrlRequest) {
    request = &GenerateApplicationPackageDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "GenerateApplicationPackageDownloadUrl")
    
    
    return
}

func NewGenerateApplicationPackageDownloadUrlResponse() (response *GenerateApplicationPackageDownloadUrlResponse) {
    response = &GenerateApplicationPackageDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateApplicationPackageDownloadUrl
// 生成应用程序包预签名下载链接
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) GenerateApplicationPackageDownloadUrl(request *GenerateApplicationPackageDownloadUrlRequest) (response *GenerateApplicationPackageDownloadUrlResponse, err error) {
    return c.GenerateApplicationPackageDownloadUrlWithContext(context.Background(), request)
}

// GenerateApplicationPackageDownloadUrl
// 生成应用程序包预签名下载链接
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) GenerateApplicationPackageDownloadUrlWithContext(ctx context.Context, request *GenerateApplicationPackageDownloadUrlRequest) (response *GenerateApplicationPackageDownloadUrlResponse, err error) {
    if request == nil {
        request = NewGenerateApplicationPackageDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateApplicationPackageDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateApplicationPackageDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationAutoscalerRequest() (request *ModifyApplicationAutoscalerRequest) {
    request = &ModifyApplicationAutoscalerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyApplicationAutoscaler")
    
    
    return
}

func NewModifyApplicationAutoscalerResponse() (response *ModifyApplicationAutoscalerResponse) {
    response = &ModifyApplicationAutoscalerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationAutoscaler
// 修改弹性伸缩策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_ATLEASTONESCALERRULESHOULDBEAPPLIED = "InvalidParameterValue.AtLeastOneScalerRuleShouldBeApplied"
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_CRONHPAREPLICASINVALID = "InvalidParameterValue.CronHpaReplicasInvalid"
//  INVALIDPARAMETERVALUE_HPAMINMAXINVALID = "InvalidParameterValue.HpaMinMaxInvalid"
//  MISSINGPARAMETER_MINMAXNUMNULL = "MissingParameter.MinMaxNumNull"
//  MISSINGPARAMETER_SCALERIDNULL = "MissingParameter.ScalerIdNull"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApplicationAutoscaler(request *ModifyApplicationAutoscalerRequest) (response *ModifyApplicationAutoscalerResponse, err error) {
    return c.ModifyApplicationAutoscalerWithContext(context.Background(), request)
}

// ModifyApplicationAutoscaler
// 修改弹性伸缩策略组合
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_ATLEASTONESCALERRULESHOULDBEAPPLIED = "InvalidParameterValue.AtLeastOneScalerRuleShouldBeApplied"
//  INVALIDPARAMETERVALUE_AUTOSCALERLARGERTHANONE = "InvalidParameterValue.AutoScalerLargerThanOne"
//  INVALIDPARAMETERVALUE_CRONHPAREPLICASINVALID = "InvalidParameterValue.CronHpaReplicasInvalid"
//  INVALIDPARAMETERVALUE_HPAMINMAXINVALID = "InvalidParameterValue.HpaMinMaxInvalid"
//  MISSINGPARAMETER_MINMAXNUMNULL = "MissingParameter.MinMaxNumNull"
//  MISSINGPARAMETER_SCALERIDNULL = "MissingParameter.ScalerIdNull"
//  RESOURCEUNAVAILABLE_APPLICATIONSTOPPED = "ResourceUnavailable.ApplicationStopped"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApplicationAutoscalerWithContext(ctx context.Context, request *ModifyApplicationAutoscalerRequest) (response *ModifyApplicationAutoscalerResponse, err error) {
    if request == nil {
        request = NewModifyApplicationAutoscalerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationAutoscaler require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationAutoscalerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationInfoRequest() (request *ModifyApplicationInfoRequest) {
    request = &ModifyApplicationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyApplicationInfo")
    
    
    return
}

func NewModifyApplicationInfoResponse() (response *ModifyApplicationInfoResponse) {
    response = &ModifyApplicationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationInfo
// 修改应用基本信息
//
// 可能返回的错误码:
//  INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApplicationInfo(request *ModifyApplicationInfoRequest) (response *ModifyApplicationInfoResponse, err error) {
    return c.ModifyApplicationInfoWithContext(context.Background(), request)
}

// ModifyApplicationInfo
// 修改应用基本信息
//
// 可能返回的错误码:
//  INTERNALERROR_CREATEAPMRESOURCEERROR = "InternalError.CreateApmResourceError"
//  INTERNALERROR_CREATESERVICEERROR = "InternalError.CreateServiceError"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApplicationInfoWithContext(ctx context.Context, request *ModifyApplicationInfoRequest) (response *ModifyApplicationInfoResponse, err error) {
    if request == nil {
        request = NewModifyApplicationInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationReplicasRequest() (request *ModifyApplicationReplicasRequest) {
    request = &ModifyApplicationReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyApplicationReplicas")
    
    
    return
}

func NewModifyApplicationReplicasResponse() (response *ModifyApplicationReplicasResponse) {
    response = &ModifyApplicationReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationReplicas
// 修改应用实例数量
//
// 可能返回的错误码:
//  OPERATIONDENIED_RESOURCEISOLATED = "OperationDenied.ResourceIsolated"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApplicationReplicas(request *ModifyApplicationReplicasRequest) (response *ModifyApplicationReplicasResponse, err error) {
    return c.ModifyApplicationReplicasWithContext(context.Background(), request)
}

// ModifyApplicationReplicas
// 修改应用实例数量
//
// 可能返回的错误码:
//  OPERATIONDENIED_RESOURCEISOLATED = "OperationDenied.ResourceIsolated"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApplicationReplicasWithContext(ctx context.Context, request *ModifyApplicationReplicasRequest) (response *ModifyApplicationReplicasResponse, err error) {
    if request == nil {
        request = NewModifyApplicationReplicasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationReplicas require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationServiceRequest() (request *ModifyApplicationServiceRequest) {
    request = &ModifyApplicationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyApplicationService")
    
    
    return
}

func NewModifyApplicationServiceResponse() (response *ModifyApplicationServiceResponse) {
    response = &ModifyApplicationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationService
// 修改服务访问方式列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETER_APPLICATIONACCESSSERVICEREACHMAXIMUM = "InvalidParameter.ApplicationAccessServiceReachMaximum"
//  INVALIDPARAMETER_LBSERVICECANNOTSUPPORTTCPUDPSAMETIME = "InvalidParameter.LBServiceCannotSupportTcpUdpSameTime"
//  INVALIDPARAMETER_MUSTPROVIDEPORTMAPPINGRULES = "InvalidParameter.MustProvidePortMappingRules"
//  INVALIDPARAMETER_SERVICENAMENOTVALID = "InvalidParameter.ServiceNameNotValid"
//  INVALIDPARAMETER_SERVICEUSERESERVESUFFIX = "InvalidParameter.ServiceUseReserveSuffix"
//  INVALIDPARAMETER_TOOMANYPORTMAPPINGRULES = "InvalidParameter.TooManyPortMappingRules"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICENOTFOUND = "InvalidParameterValue.ApplicationServiceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDEKSSERVICETYPE = "InvalidParameterValue.InvalidEksServiceType"
//  INVALIDPARAMETERVALUE_PORTISRESERVED = "InvalidParameterValue.PortIsReserved"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) ModifyApplicationService(request *ModifyApplicationServiceRequest) (response *ModifyApplicationServiceResponse, err error) {
    return c.ModifyApplicationServiceWithContext(context.Background(), request)
}

// ModifyApplicationService
// 修改服务访问方式列表
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETER_APPLICATIONACCESSSERVICEREACHMAXIMUM = "InvalidParameter.ApplicationAccessServiceReachMaximum"
//  INVALIDPARAMETER_LBSERVICECANNOTSUPPORTTCPUDPSAMETIME = "InvalidParameter.LBServiceCannotSupportTcpUdpSameTime"
//  INVALIDPARAMETER_MUSTPROVIDEPORTMAPPINGRULES = "InvalidParameter.MustProvidePortMappingRules"
//  INVALIDPARAMETER_SERVICENAMENOTVALID = "InvalidParameter.ServiceNameNotValid"
//  INVALIDPARAMETER_SERVICEUSERESERVESUFFIX = "InvalidParameter.ServiceUseReserveSuffix"
//  INVALIDPARAMETER_TOOMANYPORTMAPPINGRULES = "InvalidParameter.TooManyPortMappingRules"
//  INVALIDPARAMETERVALUE_APPLICATIONSERVICENOTFOUND = "InvalidParameterValue.ApplicationServiceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDEKSSERVICETYPE = "InvalidParameterValue.InvalidEksServiceType"
//  INVALIDPARAMETERVALUE_PORTISRESERVED = "InvalidParameterValue.PortIsReserved"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
func (c *Client) ModifyApplicationServiceWithContext(ctx context.Context, request *ModifyApplicationServiceRequest) (response *ModifyApplicationServiceResponse, err error) {
    if request == nil {
        request = NewModifyApplicationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConfigDataRequest() (request *ModifyConfigDataRequest) {
    request = &ModifyConfigDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyConfigData")
    
    
    return
}

func NewModifyConfigDataResponse() (response *ModifyConfigDataResponse) {
    response = &ModifyConfigDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyConfigData
// 编辑配置
//
// 可能返回的错误码:
//  INTERNALERROR_MODIFYCONFIGDATAERROR = "InternalError.ModifyConfigDataError"
//  INVALIDPARAMETERVALUE_CONFIGDATAINVALID = "InvalidParameterValue.ConfigDataInvalid"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyConfigData(request *ModifyConfigDataRequest) (response *ModifyConfigDataResponse, err error) {
    return c.ModifyConfigDataWithContext(context.Background(), request)
}

// ModifyConfigData
// 编辑配置
//
// 可能返回的错误码:
//  INTERNALERROR_MODIFYCONFIGDATAERROR = "InternalError.ModifyConfigDataError"
//  INVALIDPARAMETERVALUE_CONFIGDATAINVALID = "InvalidParameterValue.ConfigDataInvalid"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyConfigDataWithContext(ctx context.Context, request *ModifyConfigDataRequest) (response *ModifyConfigDataResponse, err error) {
    if request == nil {
        request = NewModifyConfigDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfigData require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvironmentRequest() (request *ModifyEnvironmentRequest) {
    request = &ModifyEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyEnvironment")
    
    
    return
}

func NewModifyEnvironmentResponse() (response *ModifyEnvironmentResponse) {
    response = &ModifyEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyEnvironment
// 编辑环境
//
// 可能返回的错误码:
//  INTERNALERROR_ADDNEWNODEERROR = "InternalError.AddNewNodeError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNAMEIMMUTABLE = "InvalidParameterValue.EnvironmentNameImmutable"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  OPERATIONDENIED_RESOURCEISOLATED = "OperationDenied.ResourceIsolated"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyEnvironment(request *ModifyEnvironmentRequest) (response *ModifyEnvironmentResponse, err error) {
    return c.ModifyEnvironmentWithContext(context.Background(), request)
}

// ModifyEnvironment
// 编辑环境
//
// 可能返回的错误码:
//  INTERNALERROR_ADDNEWNODEERROR = "InternalError.AddNewNodeError"
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNAMEIMMUTABLE = "InvalidParameterValue.EnvironmentNameImmutable"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  OPERATIONDENIED_RESOURCEISOLATED = "OperationDenied.ResourceIsolated"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyEnvironmentWithContext(ctx context.Context, request *ModifyEnvironmentRequest) (response *ModifyEnvironmentResponse, err error) {
    if request == nil {
        request = NewModifyEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIngressRequest() (request *ModifyIngressRequest) {
    request = &ModifyIngressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyIngress")
    
    
    return
}

func NewModifyIngressResponse() (response *ModifyIngressResponse) {
    response = &ModifyIngressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIngress
// 创建或者更新 Ingress 规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPDATEINGRESSERROR = "FailedOperation.UpdateIngressError"
//  INTERNALERROR_UPDATEINGRESSERROR = "InternalError.UpdateIngressError"
//  INVALIDPARAMETERVALUE_INGRESSREWRITEREQUIREDHTTPSENABLE = "InvalidParameterValue.IngressRewriteRequiredHttpsEnable"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyIngress(request *ModifyIngressRequest) (response *ModifyIngressResponse, err error) {
    return c.ModifyIngressWithContext(context.Background(), request)
}

// ModifyIngress
// 创建或者更新 Ingress 规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPDATEINGRESSERROR = "FailedOperation.UpdateIngressError"
//  INTERNALERROR_UPDATEINGRESSERROR = "InternalError.UpdateIngressError"
//  INVALIDPARAMETERVALUE_INGRESSREWRITEREQUIREDHTTPSENABLE = "InvalidParameterValue.IngressRewriteRequiredHttpsEnable"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyIngressWithContext(ctx context.Context, request *ModifyIngressRequest) (response *ModifyIngressResponse, err error) {
    if request == nil {
        request = NewModifyIngressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIngress require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIngressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogConfigRequest() (request *ModifyLogConfigRequest) {
    request = &ModifyLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ModifyLogConfig")
    
    
    return
}

func NewModifyLogConfigResponse() (response *ModifyLogConfigResponse) {
    response = &ModifyLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLogConfig
// 编辑日志收集配置
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_MODIFYLOGCONFIGERROR = "InternalError.ModifyLogConfigError"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyLogConfig(request *ModifyLogConfigRequest) (response *ModifyLogConfigResponse, err error) {
    return c.ModifyLogConfigWithContext(context.Background(), request)
}

// ModifyLogConfig
// 编辑日志收集配置
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_MODIFYLOGCONFIGERROR = "InternalError.ModifyLogConfigError"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyLogConfigWithContext(ctx context.Context, request *ModifyLogConfigRequest) (response *ModifyLogConfigResponse, err error) {
    if request == nil {
        request = NewModifyLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRestartApplicationRequest() (request *RestartApplicationRequest) {
    request = &RestartApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "RestartApplication")
    
    
    return
}

func NewRestartApplicationResponse() (response *RestartApplicationResponse) {
    response = &RestartApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartApplication
// 服务重启
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  OPERATIONDENIED_RESOURCEISOLATED = "OperationDenied.ResourceIsolated"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RestartApplication(request *RestartApplicationRequest) (response *RestartApplicationResponse, err error) {
    return c.RestartApplicationWithContext(context.Background(), request)
}

// RestartApplication
// 服务重启
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_RESTARTAPPLICATIONERROR = "InternalError.RestartApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  OPERATIONDENIED_RESOURCEISOLATED = "OperationDenied.ResourceIsolated"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RestartApplicationWithContext(ctx context.Context, request *RestartApplicationRequest) (response *RestartApplicationResponse, err error) {
    if request == nil {
        request = NewRestartApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewRestartApplicationPodRequest() (request *RestartApplicationPodRequest) {
    request = &RestartApplicationPodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "RestartApplicationPod")
    
    
    return
}

func NewRestartApplicationPodResponse() (response *RestartApplicationPodResponse) {
    response = &RestartApplicationPodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartApplicationPod
// 重启应用实例
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RestartApplicationPod(request *RestartApplicationPodRequest) (response *RestartApplicationPodResponse, err error) {
    return c.RestartApplicationPodWithContext(context.Background(), request)
}

// RestartApplicationPod
// 重启应用实例
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RestartApplicationPodWithContext(ctx context.Context, request *RestartApplicationPodRequest) (response *RestartApplicationPodResponse, err error) {
    if request == nil {
        request = NewRestartApplicationPodRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartApplicationPod require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartApplicationPodResponse()
    err = c.Send(request, response)
    return
}

func NewResumeDeployApplicationRequest() (request *ResumeDeployApplicationRequest) {
    request = &ResumeDeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "ResumeDeployApplication")
    
    
    return
}

func NewResumeDeployApplicationResponse() (response *ResumeDeployApplicationResponse) {
    response = &ResumeDeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeDeployApplication
// 开始下一批次发布
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ResumeDeployApplication(request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    return c.ResumeDeployApplicationWithContext(context.Background(), request)
}

// ResumeDeployApplication
// 开始下一批次发布
//
// 可能返回的错误码:
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ResumeDeployApplicationWithContext(ctx context.Context, request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    if request == nil {
        request = NewResumeDeployApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeDeployApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewRevertDeployApplicationRequest() (request *RevertDeployApplicationRequest) {
    request = &RevertDeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "RevertDeployApplication")
    
    
    return
}

func NewRevertDeployApplicationResponse() (response *RevertDeployApplicationResponse) {
    response = &RevertDeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevertDeployApplication
// 回滚分批发布
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RevertDeployApplication(request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    return c.RevertDeployApplicationWithContext(context.Background(), request)
}

// RevertDeployApplication
// 回滚分批发布
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RevertDeployApplicationWithContext(ctx context.Context, request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    if request == nil {
        request = NewRevertDeployApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevertDeployApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevertDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewRollingUpdateApplicationByVersionRequest() (request *RollingUpdateApplicationByVersionRequest) {
    request = &RollingUpdateApplicationByVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "RollingUpdateApplicationByVersion")
    
    
    return
}

func NewRollingUpdateApplicationByVersionResponse() (response *RollingUpdateApplicationByVersionResponse) {
    response = &RollingUpdateApplicationByVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RollingUpdateApplicationByVersion
// 更新应用部署版本
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_APMNOTBIND = "InvalidParameterValue.ApmNotBind"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_INVALIDENVNAME = "InvalidParameterValue.InvalidEnvName"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPATH = "InvalidParameterValue.InvalidMountPath"
//  INVALIDPARAMETERVALUE_JDKVERSIONREQUIRED = "InvalidParameterValue.JdkVersionRequired"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_OSNOTSUPPORT = "InvalidParameterValue.OsNotSupport"
//  INVALIDPARAMETERVALUE_REGISTRYNOTBIND = "InvalidParameterValue.RegistryNotBind"
//  INVALIDPARAMETERVALUE_SERVICENOTBELONGTOAPPID = "InvalidParameterValue.ServiceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEPODREACHMAXIMUM = "InvalidParameterValue.ServicePodReachMaximum"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_VERSIONLENGTHLIMIT = "InvalidParameterValue.VersionLengthLimit"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  MISSINGPARAMETER_DEPLOYMODENULL = "MissingParameter.DeployModeNull"
//  MISSINGPARAMETER_LOGSETORTOPICNULL = "MissingParameter.LogsetOrTopicNull"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SVCREPONOTREADY = "MissingParameter.SvcRepoNotReady"
//  RESOURCEINUSE_SERVICEDEPLOYING = "ResourceInUse.ServiceDeploying"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
//  UNAUTHORIZEDOPERATION_MISSINGEKSLOGROLE = "UnauthorizedOperation.MissingEksLogRole"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RollingUpdateApplicationByVersion(request *RollingUpdateApplicationByVersionRequest) (response *RollingUpdateApplicationByVersionResponse, err error) {
    return c.RollingUpdateApplicationByVersionWithContext(context.Background(), request)
}

// RollingUpdateApplicationByVersion
// 更新应用部署版本
//
// 可能返回的错误码:
//  INTERNALERROR_DEFAULTINTERNALERROR = "InternalError.DefaultInternalError"
//  INTERNALERROR_DEPLOYVERSIONERROR = "InternalError.DeployVersionError"
//  INVALIDPARAMETERVALUE_APMNOTBIND = "InvalidParameterValue.ApmNotBind"
//  INVALIDPARAMETERVALUE_INVALIDDEPLOYVERSION = "InvalidParameterValue.InvalidDeployVersion"
//  INVALIDPARAMETERVALUE_INVALIDENVNAME = "InvalidParameterValue.InvalidEnvName"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPATH = "InvalidParameterValue.InvalidMountPath"
//  INVALIDPARAMETERVALUE_JDKVERSIONREQUIRED = "InvalidParameterValue.JdkVersionRequired"
//  INVALIDPARAMETERVALUE_NAMESPACENOTBELONGTOAPPID = "InvalidParameterValue.NamespaceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_OSNOTSUPPORT = "InvalidParameterValue.OsNotSupport"
//  INVALIDPARAMETERVALUE_REGISTRYNOTBIND = "InvalidParameterValue.RegistryNotBind"
//  INVALIDPARAMETERVALUE_SERVICENOTBELONGTOAPPID = "InvalidParameterValue.ServiceNotBelongToAppid"
//  INVALIDPARAMETERVALUE_SERVICEPODREACHMAXIMUM = "InvalidParameterValue.ServicePodReachMaximum"
//  INVALIDPARAMETERVALUE_TEMIDINVALID = "InvalidParameterValue.TemIdInvalid"
//  INVALIDPARAMETERVALUE_VERSIONLENGTHLIMIT = "InvalidParameterValue.VersionLengthLimit"
//  INVALIDPARAMETERVALUE_VERSIONLOWERCASE = "InvalidParameterValue.VersionLowerCase"
//  MISSINGPARAMETER_DEPLOYMODENULL = "MissingParameter.DeployModeNull"
//  MISSINGPARAMETER_LOGSETORTOPICNULL = "MissingParameter.LogsetOrTopicNull"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  MISSINGPARAMETER_SVCREPONOTREADY = "MissingParameter.SvcRepoNotReady"
//  RESOURCEINUSE_SERVICEDEPLOYING = "ResourceInUse.ServiceDeploying"
//  RESOURCENOTFOUND_SERVICENOTFOUND = "ResourceNotFound.ServiceNotFound"
//  RESOURCENOTFOUND_VERSIONNAMESPACENOTFOUND = "ResourceNotFound.VersionNamespaceNotFound"
//  RESOURCEUNAVAILABLE_WAITFORKRUISE = "ResourceUnavailable.WaitForKruise"
//  UNAUTHORIZEDOPERATION_MISSINGEKSLOGROLE = "UnauthorizedOperation.MissingEksLogRole"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RollingUpdateApplicationByVersionWithContext(ctx context.Context, request *RollingUpdateApplicationByVersionRequest) (response *RollingUpdateApplicationByVersionResponse, err error) {
    if request == nil {
        request = NewRollingUpdateApplicationByVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollingUpdateApplicationByVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollingUpdateApplicationByVersionResponse()
    err = c.Send(request, response)
    return
}

func NewStopApplicationRequest() (request *StopApplicationRequest) {
    request = &StopApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tem", APIVersion, "StopApplication")
    
    
    return
}

func NewStopApplicationResponse() (response *StopApplicationResponse) {
    response = &StopApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopApplication
// 服务停止
//
// 可能返回的错误码:
//  INTERNALERROR_STOPAPPLICATIONERROR = "InternalError.StopApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) StopApplication(request *StopApplicationRequest) (response *StopApplicationResponse, err error) {
    return c.StopApplicationWithContext(context.Background(), request)
}

// StopApplication
// 服务停止
//
// 可能返回的错误码:
//  INTERNALERROR_STOPAPPLICATIONERROR = "InternalError.StopApplicationError"
//  MISSINGPARAMETER_NAMESPACEIDNULL = "MissingParameter.NamespaceIdNull"
//  RESOURCENOTFOUND_SERVICERUNNINGVERSIONNOTFOUND = "ResourceNotFound.ServiceRunningVersionNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) StopApplicationWithContext(ctx context.Context, request *StopApplicationRequest) (response *StopApplicationResponse, err error) {
    if request == nil {
        request = NewStopApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopApplicationResponse()
    err = c.Send(request, response)
    return
}
