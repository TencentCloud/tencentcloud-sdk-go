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

package v20220106

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-01-06"

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


func NewCheckCertificateRequest() (request *CheckCertificateRequest) {
    request = &CheckCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CheckCertificate")
    
    
    return
}

func NewCheckCertificateResponse() (response *CheckCertificateResponse) {
    response = &CheckCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckCertificate
// 校验证书
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CheckCertificate(request *CheckCertificateRequest) (response *CheckCertificateResponse, err error) {
    return c.CheckCertificateWithContext(context.Background(), request)
}

// CheckCertificate
// 校验证书
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CheckCertificateWithContext(ctx context.Context, request *CheckCertificateRequest) (response *CheckCertificateResponse, err error) {
    if request == nil {
        request = NewCheckCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRequest() (request *CreateApplicationProxyRequest) {
    request = &CreateApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxy")
    
    
    return
}

func NewCreateApplicationProxyResponse() (response *CreateApplicationProxyResponse) {
    response = &CreateApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxy
// 创建应用代理
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxy(request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    return c.CreateApplicationProxyWithContext(context.Background(), request)
}

// CreateApplicationProxy
// 创建应用代理
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyWithContext(ctx context.Context, request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRuleRequest() (request *CreateApplicationProxyRuleRequest) {
    request = &CreateApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxyRule")
    
    
    return
}

func NewCreateApplicationProxyRuleResponse() (response *CreateApplicationProxyRuleResponse) {
    response = &CreateApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxyRule
// 创建应用代理规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyRule(request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    return c.CreateApplicationProxyRuleWithContext(context.Background(), request)
}

// CreateApplicationProxyRule
// 创建应用代理规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyRuleWithContext(ctx context.Context, request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRulesRequest() (request *CreateApplicationProxyRulesRequest) {
    request = &CreateApplicationProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxyRules")
    
    
    return
}

func NewCreateApplicationProxyRulesResponse() (response *CreateApplicationProxyRulesResponse) {
    response = &CreateApplicationProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxyRules
// 批量创建应用代理规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyRules(request *CreateApplicationProxyRulesRequest) (response *CreateApplicationProxyRulesResponse, err error) {
    return c.CreateApplicationProxyRulesWithContext(context.Background(), request)
}

// CreateApplicationProxyRules
// 批量创建应用代理规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyRulesWithContext(ctx context.Context, request *CreateApplicationProxyRulesRequest) (response *CreateApplicationProxyRulesResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomErrorPageRequest() (request *CreateCustomErrorPageRequest) {
    request = &CreateCustomErrorPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateCustomErrorPage")
    
    
    return
}

func NewCreateCustomErrorPageResponse() (response *CreateCustomErrorPageResponse) {
    response = &CreateCustomErrorPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomErrorPage
// 创建自定义规则的自定义页
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateCustomErrorPage(request *CreateCustomErrorPageRequest) (response *CreateCustomErrorPageResponse, err error) {
    return c.CreateCustomErrorPageWithContext(context.Background(), request)
}

// CreateCustomErrorPage
// 创建自定义规则的自定义页
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateCustomErrorPageWithContext(ctx context.Context, request *CreateCustomErrorPageRequest) (response *CreateCustomErrorPageResponse, err error) {
    if request == nil {
        request = NewCreateCustomErrorPageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomErrorPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomErrorPageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDnsRecordRequest() (request *CreateDnsRecordRequest) {
    request = &CreateDnsRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateDnsRecord")
    
    
    return
}

func NewCreateDnsRecordResponse() (response *CreateDnsRecordResponse) {
    response = &CreateDnsRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDnsRecord
// 创建 DNS 记录
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYNAME = "InvalidParameterValue.InvalidProxyName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateDnsRecord(request *CreateDnsRecordRequest) (response *CreateDnsRecordResponse, err error) {
    return c.CreateDnsRecordWithContext(context.Background(), request)
}

// CreateDnsRecord
// 创建 DNS 记录
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYNAME = "InvalidParameterValue.InvalidProxyName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateDnsRecordWithContext(ctx context.Context, request *CreateDnsRecordRequest) (response *CreateDnsRecordResponse, err error) {
    if request == nil {
        request = NewCreateDnsRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDnsRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDnsRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancingRequest() (request *CreateLoadBalancingRequest) {
    request = &CreateLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateLoadBalancing")
    
    
    return
}

func NewCreateLoadBalancingResponse() (response *CreateLoadBalancingResponse) {
    response = &CreateLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLoadBalancing
// 创建负载均衡
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateLoadBalancing(request *CreateLoadBalancingRequest) (response *CreateLoadBalancingResponse, err error) {
    return c.CreateLoadBalancingWithContext(context.Background(), request)
}

// CreateLoadBalancing
// 创建负载均衡
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateLoadBalancingWithContext(ctx context.Context, request *CreateLoadBalancingRequest) (response *CreateLoadBalancingResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOriginGroupRequest() (request *CreateOriginGroupRequest) {
    request = &CreateOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateOriginGroup")
    
    
    return
}

func NewCreateOriginGroupResponse() (response *CreateOriginGroupResponse) {
    response = &CreateOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOriginGroup
// 源站组创建
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroup(request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    return c.CreateOriginGroupWithContext(context.Background(), request)
}

// CreateOriginGroup
// 源站组创建
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroupWithContext(ctx context.Context, request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    if request == nil {
        request = NewCreateOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePlanForZoneRequest() (request *CreatePlanForZoneRequest) {
    request = &CreatePlanForZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePlanForZone")
    
    
    return
}

func NewCreatePlanForZoneResponse() (response *CreatePlanForZoneResponse) {
    response = &CreatePlanForZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePlanForZone
// 为未购买套餐的站点购买套餐
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreatePlanForZone(request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    return c.CreatePlanForZoneWithContext(context.Background(), request)
}

// CreatePlanForZone
// 为未购买套餐的站点购买套餐
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreatePlanForZoneWithContext(ctx context.Context, request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    if request == nil {
        request = NewCreatePlanForZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlanForZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePlanForZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrefetchTaskRequest() (request *CreatePrefetchTaskRequest) {
    request = &CreatePrefetchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePrefetchTask")
    
    
    return
}

func NewCreatePrefetchTaskResponse() (response *CreatePrefetchTaskResponse) {
    response = &CreatePrefetchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrefetchTask
// 创建预热任务
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreatePrefetchTask(request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    return c.CreatePrefetchTaskWithContext(context.Background(), request)
}

// CreatePrefetchTask
// 创建预热任务
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreatePrefetchTaskWithContext(ctx context.Context, request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    if request == nil {
        request = NewCreatePrefetchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrefetchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrefetchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePurgeTaskRequest() (request *CreatePurgeTaskRequest) {
    request = &CreatePurgeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePurgeTask")
    
    
    return
}

func NewCreatePurgeTaskResponse() (response *CreatePurgeTaskResponse) {
    response = &CreatePurgeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePurgeTask
// 创建清除缓存任务
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreatePurgeTask(request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    return c.CreatePurgeTaskWithContext(context.Background(), request)
}

// CreatePurgeTask
// 创建清除缓存任务
//
// 可能返回的错误码:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreatePurgeTaskWithContext(ctx context.Context, request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    if request == nil {
        request = NewCreatePurgeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePurgeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePurgeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRule
// 规则引擎创建规则。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// 规则引擎创建规则。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
    request = &CreateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateZone")
    
    
    return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
    response = &CreateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateZone
// 用于用户接入新的站点
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    return c.CreateZoneWithContext(context.Background(), request)
}

// CreateZone
// 用于用户接入新的站点
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateZoneWithContext(ctx context.Context, request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    if request == nil {
        request = NewCreateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRequest() (request *DeleteApplicationProxyRequest) {
    request = &DeleteApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxy")
    
    
    return
}

func NewDeleteApplicationProxyResponse() (response *DeleteApplicationProxyResponse) {
    response = &DeleteApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationProxy
// 删除应用代理
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteApplicationProxy(request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    return c.DeleteApplicationProxyWithContext(context.Background(), request)
}

// DeleteApplicationProxy
// 删除应用代理
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteApplicationProxyWithContext(ctx context.Context, request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRuleRequest() (request *DeleteApplicationProxyRuleRequest) {
    request = &DeleteApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxyRule")
    
    
    return
}

func NewDeleteApplicationProxyRuleResponse() (response *DeleteApplicationProxyRuleResponse) {
    response = &DeleteApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationProxyRule
// 删除应用代理规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteApplicationProxyRule(request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    return c.DeleteApplicationProxyRuleWithContext(context.Background(), request)
}

// DeleteApplicationProxyRule
// 删除应用代理规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteApplicationProxyRuleWithContext(ctx context.Context, request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDnsRecordsRequest() (request *DeleteDnsRecordsRequest) {
    request = &DeleteDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteDnsRecords")
    
    
    return
}

func NewDeleteDnsRecordsResponse() (response *DeleteDnsRecordsResponse) {
    response = &DeleteDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDnsRecords
// 批量删除 DNS 记录
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDnsRecords(request *DeleteDnsRecordsRequest) (response *DeleteDnsRecordsResponse, err error) {
    return c.DeleteDnsRecordsWithContext(context.Background(), request)
}

// DeleteDnsRecords
// 批量删除 DNS 记录
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDnsRecordsWithContext(ctx context.Context, request *DeleteDnsRecordsRequest) (response *DeleteDnsRecordsResponse, err error) {
    if request == nil {
        request = NewDeleteDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancingRequest() (request *DeleteLoadBalancingRequest) {
    request = &DeleteLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteLoadBalancing")
    
    
    return
}

func NewDeleteLoadBalancingResponse() (response *DeleteLoadBalancingResponse) {
    response = &DeleteLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoadBalancing
// 删除负载均衡
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteLoadBalancing(request *DeleteLoadBalancingRequest) (response *DeleteLoadBalancingResponse, err error) {
    return c.DeleteLoadBalancingWithContext(context.Background(), request)
}

// DeleteLoadBalancing
// 删除负载均衡
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteLoadBalancingWithContext(ctx context.Context, request *DeleteLoadBalancingRequest) (response *DeleteLoadBalancingResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOriginGroupRequest() (request *DeleteOriginGroupRequest) {
    request = &DeleteOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteOriginGroup")
    
    
    return
}

func NewDeleteOriginGroupResponse() (response *DeleteOriginGroupResponse) {
    response = &DeleteOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOriginGroup
// 源站组删除
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteOriginGroup(request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    return c.DeleteOriginGroupWithContext(context.Background(), request)
}

// DeleteOriginGroup
// 源站组删除
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteOriginGroupWithContext(ctx context.Context, request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    if request == nil {
        request = NewDeleteOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRulesRequest() (request *DeleteRulesRequest) {
    request = &DeleteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteRules")
    
    
    return
}

func NewDeleteRulesResponse() (response *DeleteRulesResponse) {
    response = &DeleteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRules
// 批量删除规则引擎规则。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    return c.DeleteRulesWithContext(context.Background(), request)
}

// DeleteRules
// 批量删除规则引擎规则。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteRulesWithContext(ctx context.Context, request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
    request = &DeleteZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteZone")
    
    
    return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
    response = &DeleteZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteZone
// 删除站点
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    return c.DeleteZoneWithContext(context.Background(), request)
}

// DeleteZone
// 删除站点
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    if request == nil {
        request = NewDeleteZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProxyRequest() (request *DescribeApplicationProxyRequest) {
    request = &DescribeApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeApplicationProxy")
    
    
    return
}

func NewDescribeApplicationProxyResponse() (response *DescribeApplicationProxyResponse) {
    response = &DescribeApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationProxy
// 获取应用代理列表
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeApplicationProxy(request *DescribeApplicationProxyRequest) (response *DescribeApplicationProxyResponse, err error) {
    return c.DescribeApplicationProxyWithContext(context.Background(), request)
}

// DescribeApplicationProxy
// 获取应用代理列表
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeApplicationProxyWithContext(ctx context.Context, request *DescribeApplicationProxyRequest) (response *DescribeApplicationProxyResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProxyDetailRequest() (request *DescribeApplicationProxyDetailRequest) {
    request = &DescribeApplicationProxyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeApplicationProxyDetail")
    
    
    return
}

func NewDescribeApplicationProxyDetailResponse() (response *DescribeApplicationProxyDetailResponse) {
    response = &DescribeApplicationProxyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationProxyDetail
// 获取应用代理详细信息
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeApplicationProxyDetail(request *DescribeApplicationProxyDetailRequest) (response *DescribeApplicationProxyDetailResponse, err error) {
    return c.DescribeApplicationProxyDetailWithContext(context.Background(), request)
}

// DescribeApplicationProxyDetail
// 获取应用代理详细信息
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeApplicationProxyDetailWithContext(ctx context.Context, request *DescribeApplicationProxyDetailRequest) (response *DescribeApplicationProxyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProxyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProxyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProxyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailablePlansRequest() (request *DescribeAvailablePlansRequest) {
    request = &DescribeAvailablePlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAvailablePlans")
    
    
    return
}

func NewDescribeAvailablePlansResponse() (response *DescribeAvailablePlansResponse) {
    response = &DescribeAvailablePlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailablePlans
// 查询当前账户可用套餐信息列表
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAvailablePlans(request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    return c.DescribeAvailablePlansWithContext(context.Background(), request)
}

// DescribeAvailablePlans
// 查询当前账户可用套餐信息列表
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAvailablePlansWithContext(ctx context.Context, request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    if request == nil {
        request = NewDescribeAvailablePlansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailablePlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailablePlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBotLogRequest() (request *DescribeBotLogRequest) {
    request = &DescribeBotLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeBotLog")
    
    
    return
}

func NewDescribeBotLogResponse() (response *DescribeBotLogResponse) {
    response = &DescribeBotLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBotLog
// 查询Bot攻击日志
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBotLog(request *DescribeBotLogRequest) (response *DescribeBotLogResponse, err error) {
    return c.DescribeBotLogWithContext(context.Background(), request)
}

// DescribeBotLog
// 查询Bot攻击日志
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBotLogWithContext(ctx context.Context, request *DescribeBotLogRequest) (response *DescribeBotLogResponse, err error) {
    if request == nil {
        request = NewDescribeBotLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBotLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBotLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBotManagedRulesRequest() (request *DescribeBotManagedRulesRequest) {
    request = &DescribeBotManagedRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeBotManagedRules")
    
    
    return
}

func NewDescribeBotManagedRulesResponse() (response *DescribeBotManagedRulesResponse) {
    response = &DescribeBotManagedRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBotManagedRules
// 分页查询Bot托管规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) DescribeBotManagedRules(request *DescribeBotManagedRulesRequest) (response *DescribeBotManagedRulesResponse, err error) {
    return c.DescribeBotManagedRulesWithContext(context.Background(), request)
}

// DescribeBotManagedRules
// 分页查询Bot托管规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) DescribeBotManagedRulesWithContext(ctx context.Context, request *DescribeBotManagedRulesRequest) (response *DescribeBotManagedRulesResponse, err error) {
    if request == nil {
        request = NewDescribeBotManagedRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBotManagedRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBotManagedRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCnameStatusRequest() (request *DescribeCnameStatusRequest) {
    request = &DescribeCnameStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeCnameStatus")
    
    
    return
}

func NewDescribeCnameStatusResponse() (response *DescribeCnameStatusResponse) {
    response = &DescribeCnameStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCnameStatus
// 查询域名 CNAME 状态
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeCnameStatus(request *DescribeCnameStatusRequest) (response *DescribeCnameStatusResponse, err error) {
    return c.DescribeCnameStatusWithContext(context.Background(), request)
}

// DescribeCnameStatus
// 查询域名 CNAME 状态
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeCnameStatusWithContext(ctx context.Context, request *DescribeCnameStatusRequest) (response *DescribeCnameStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCnameStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCnameStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCnameStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSPolicyRequest() (request *DescribeDDoSPolicyRequest) {
    request = &DescribeDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSPolicy")
    
    
    return
}

func NewDescribeDDoSPolicyResponse() (response *DescribeDDoSPolicyResponse) {
    response = &DescribeDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDoSPolicy
// 查询DDoS防护配置详情
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSPolicy(request *DescribeDDoSPolicyRequest) (response *DescribeDDoSPolicyResponse, err error) {
    return c.DescribeDDoSPolicyWithContext(context.Background(), request)
}

// DescribeDDoSPolicy
// 查询DDoS防护配置详情
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSPolicyWithContext(ctx context.Context, request *DescribeDDoSPolicyRequest) (response *DescribeDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDosAttackDataRequest() (request *DescribeDDosAttackDataRequest) {
    request = &DescribeDDosAttackDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDosAttackData")
    
    
    return
}

func NewDescribeDDosAttackDataResponse() (response *DescribeDDosAttackDataResponse) {
    response = &DescribeDDosAttackDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDosAttackData
// 查询DDos攻击时序数据
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackData(request *DescribeDDosAttackDataRequest) (response *DescribeDDosAttackDataResponse, err error) {
    return c.DescribeDDosAttackDataWithContext(context.Background(), request)
}

// DescribeDDosAttackData
// 查询DDos攻击时序数据
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackDataWithContext(ctx context.Context, request *DescribeDDosAttackDataRequest) (response *DescribeDDosAttackDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDosAttackDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDosAttackData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDosAttackDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDosAttackEventRequest() (request *DescribeDDosAttackEventRequest) {
    request = &DescribeDDosAttackEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDosAttackEvent")
    
    
    return
}

func NewDescribeDDosAttackEventResponse() (response *DescribeDDosAttackEventResponse) {
    response = &DescribeDDosAttackEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDosAttackEvent
// 查询DDos攻击事件
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackEvent(request *DescribeDDosAttackEventRequest) (response *DescribeDDosAttackEventResponse, err error) {
    return c.DescribeDDosAttackEventWithContext(context.Background(), request)
}

// DescribeDDosAttackEvent
// 查询DDos攻击事件
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackEventWithContext(ctx context.Context, request *DescribeDDosAttackEventRequest) (response *DescribeDDosAttackEventResponse, err error) {
    if request == nil {
        request = NewDescribeDDosAttackEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDosAttackEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDosAttackEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDosAttackEventDetailRequest() (request *DescribeDDosAttackEventDetailRequest) {
    request = &DescribeDDosAttackEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDosAttackEventDetail")
    
    
    return
}

func NewDescribeDDosAttackEventDetailResponse() (response *DescribeDDosAttackEventDetailResponse) {
    response = &DescribeDDosAttackEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDosAttackEventDetail
// 查询DDos攻击事件详情
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackEventDetail(request *DescribeDDosAttackEventDetailRequest) (response *DescribeDDosAttackEventDetailResponse, err error) {
    return c.DescribeDDosAttackEventDetailWithContext(context.Background(), request)
}

// DescribeDDosAttackEventDetail
// 查询DDos攻击事件详情
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackEventDetailWithContext(ctx context.Context, request *DescribeDDosAttackEventDetailRequest) (response *DescribeDDosAttackEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDDosAttackEventDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDosAttackEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDosAttackEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDosAttackSourceEventRequest() (request *DescribeDDosAttackSourceEventRequest) {
    request = &DescribeDDosAttackSourceEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDosAttackSourceEvent")
    
    
    return
}

func NewDescribeDDosAttackSourceEventResponse() (response *DescribeDDosAttackSourceEventResponse) {
    response = &DescribeDDosAttackSourceEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDosAttackSourceEvent
// 查询DDos攻击源信息
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackSourceEvent(request *DescribeDDosAttackSourceEventRequest) (response *DescribeDDosAttackSourceEventResponse, err error) {
    return c.DescribeDDosAttackSourceEventWithContext(context.Background(), request)
}

// DescribeDDosAttackSourceEvent
// 查询DDos攻击源信息
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackSourceEventWithContext(ctx context.Context, request *DescribeDDosAttackSourceEventRequest) (response *DescribeDDosAttackSourceEventResponse, err error) {
    if request == nil {
        request = NewDescribeDDosAttackSourceEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDosAttackSourceEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDosAttackSourceEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDosAttackTopDataRequest() (request *DescribeDDosAttackTopDataRequest) {
    request = &DescribeDDosAttackTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDosAttackTopData")
    
    
    return
}

func NewDescribeDDosAttackTopDataResponse() (response *DescribeDDosAttackTopDataResponse) {
    response = &DescribeDDosAttackTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDosAttackTopData
// 查询DDos攻击Top数据
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackTopData(request *DescribeDDosAttackTopDataRequest) (response *DescribeDDosAttackTopDataResponse, err error) {
    return c.DescribeDDosAttackTopDataWithContext(context.Background(), request)
}

// DescribeDDosAttackTopData
// 查询DDos攻击Top数据
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosAttackTopDataWithContext(ctx context.Context, request *DescribeDDosAttackTopDataRequest) (response *DescribeDDosAttackTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDosAttackTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDosAttackTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDosAttackTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDosMajorAttackEventRequest() (request *DescribeDDosMajorAttackEventRequest) {
    request = &DescribeDDosMajorAttackEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDosMajorAttackEvent")
    
    
    return
}

func NewDescribeDDosMajorAttackEventResponse() (response *DescribeDDosMajorAttackEventResponse) {
    response = &DescribeDDosMajorAttackEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDDosMajorAttackEvent
// 此接口已经做不兼容变更，所以此版本的接口已经无法调用，故作下线处理
//
// 
//
// 查询DDos主攻击事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosMajorAttackEvent(request *DescribeDDosMajorAttackEventRequest) (response *DescribeDDosMajorAttackEventResponse, err error) {
    return c.DescribeDDosMajorAttackEventWithContext(context.Background(), request)
}

// DescribeDDosMajorAttackEvent
// 此接口已经做不兼容变更，所以此版本的接口已经无法调用，故作下线处理
//
// 
//
// 查询DDos主攻击事件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDosMajorAttackEventWithContext(ctx context.Context, request *DescribeDDosMajorAttackEventRequest) (response *DescribeDDosMajorAttackEventResponse, err error) {
    if request == nil {
        request = NewDescribeDDosMajorAttackEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDosMajorAttackEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDosMajorAttackEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultCertificatesRequest() (request *DescribeDefaultCertificatesRequest) {
    request = &DescribeDefaultCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDefaultCertificates")
    
    
    return
}

func NewDescribeDefaultCertificatesResponse() (response *DescribeDefaultCertificatesResponse) {
    response = &DescribeDefaultCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultCertificates
// 查询默认证书列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificates(request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    return c.DescribeDefaultCertificatesWithContext(context.Background(), request)
}

// DescribeDefaultCertificates
// 查询默认证书列表
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificatesWithContext(ctx context.Context, request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnsDataRequest() (request *DescribeDnsDataRequest) {
    request = &DescribeDnsDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnsData")
    
    
    return
}

func NewDescribeDnsDataResponse() (response *DescribeDnsDataResponse) {
    response = &DescribeDnsDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDnsData
// 获取DNS请求数统计曲线
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnsData(request *DescribeDnsDataRequest) (response *DescribeDnsDataResponse, err error) {
    return c.DescribeDnsDataWithContext(context.Background(), request)
}

// DescribeDnsData
// 获取DNS请求数统计曲线
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnsDataWithContext(ctx context.Context, request *DescribeDnsDataRequest) (response *DescribeDnsDataResponse, err error) {
    if request == nil {
        request = NewDescribeDnsDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnsData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnsDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnsRecordsRequest() (request *DescribeDnsRecordsRequest) {
    request = &DescribeDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnsRecords")
    
    
    return
}

func NewDescribeDnsRecordsResponse() (response *DescribeDnsRecordsResponse) {
    response = &DescribeDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDnsRecords
// 查询 DNS 记录列表，支持搜索、分页、排序、过滤。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnsRecords(request *DescribeDnsRecordsRequest) (response *DescribeDnsRecordsResponse, err error) {
    return c.DescribeDnsRecordsWithContext(context.Background(), request)
}

// DescribeDnsRecords
// 查询 DNS 记录列表，支持搜索、分页、排序、过滤。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnsRecordsWithContext(ctx context.Context, request *DescribeDnsRecordsRequest) (response *DescribeDnsRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnssecRequest() (request *DescribeDnssecRequest) {
    request = &DescribeDnssecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnssec")
    
    
    return
}

func NewDescribeDnssecResponse() (response *DescribeDnssecResponse) {
    response = &DescribeDnssecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDnssec
// 用于查询 DNSSEC 相关信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnssec(request *DescribeDnssecRequest) (response *DescribeDnssecResponse, err error) {
    return c.DescribeDnssecWithContext(context.Background(), request)
}

// DescribeDnssec
// 用于查询 DNSSEC 相关信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnssecWithContext(ctx context.Context, request *DescribeDnssecRequest) (response *DescribeDnssecResponse, err error) {
    if request == nil {
        request = NewDescribeDnssecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnssec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnssecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsCertificateRequest() (request *DescribeHostsCertificateRequest) {
    request = &DescribeHostsCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeHostsCertificate")
    
    
    return
}

func NewDescribeHostsCertificateResponse() (response *DescribeHostsCertificateResponse) {
    response = &DescribeHostsCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostsCertificate
// 查询域名证书列表，支持搜索、分页、排序、过滤。
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeHostsCertificate(request *DescribeHostsCertificateRequest) (response *DescribeHostsCertificateResponse, err error) {
    return c.DescribeHostsCertificateWithContext(context.Background(), request)
}

// DescribeHostsCertificate
// 查询域名证书列表，支持搜索、分页、排序、过滤。
//
// 可能返回的错误码:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeHostsCertificateWithContext(ctx context.Context, request *DescribeHostsCertificateRequest) (response *DescribeHostsCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeHostsCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostsCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsSettingRequest() (request *DescribeHostsSettingRequest) {
    request = &DescribeHostsSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeHostsSetting")
    
    
    return
}

func NewDescribeHostsSettingResponse() (response *DescribeHostsSettingResponse) {
    response = &DescribeHostsSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostsSetting
// 用于查询域名配置信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeHostsSetting(request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    return c.DescribeHostsSettingWithContext(context.Background(), request)
}

// DescribeHostsSetting
// 用于查询域名配置信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeHostsSettingWithContext(ctx context.Context, request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    if request == nil {
        request = NewDescribeHostsSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostsSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdentificationRequest() (request *DescribeIdentificationRequest) {
    request = &DescribeIdentificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIdentification")
    
    
    return
}

func NewDescribeIdentificationResponse() (response *DescribeIdentificationResponse) {
    response = &DescribeIdentificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIdentification
// 查询验证结果
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIdentification(request *DescribeIdentificationRequest) (response *DescribeIdentificationResponse, err error) {
    return c.DescribeIdentificationWithContext(context.Background(), request)
}

// DescribeIdentification
// 查询验证结果
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIdentificationWithContext(ctx context.Context, request *DescribeIdentificationRequest) (response *DescribeIdentificationResponse, err error) {
    if request == nil {
        request = NewDescribeIdentificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdentification require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdentificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancingRequest() (request *DescribeLoadBalancingRequest) {
    request = &DescribeLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLoadBalancing")
    
    
    return
}

func NewDescribeLoadBalancingResponse() (response *DescribeLoadBalancingResponse) {
    response = &DescribeLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancing
// 获取负载均衡列表
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeLoadBalancing(request *DescribeLoadBalancingRequest) (response *DescribeLoadBalancingResponse, err error) {
    return c.DescribeLoadBalancingWithContext(context.Background(), request)
}

// DescribeLoadBalancing
// 获取负载均衡列表
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeLoadBalancingWithContext(ctx context.Context, request *DescribeLoadBalancingRequest) (response *DescribeLoadBalancingResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancingDetailRequest() (request *DescribeLoadBalancingDetailRequest) {
    request = &DescribeLoadBalancingDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLoadBalancingDetail")
    
    
    return
}

func NewDescribeLoadBalancingDetailResponse() (response *DescribeLoadBalancingDetailResponse) {
    response = &DescribeLoadBalancingDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancingDetail
// 获取负载均衡详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoadBalancingDetail(request *DescribeLoadBalancingDetailRequest) (response *DescribeLoadBalancingDetailResponse, err error) {
    return c.DescribeLoadBalancingDetailWithContext(context.Background(), request)
}

// DescribeLoadBalancingDetail
// 获取负载均衡详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoadBalancingDetailWithContext(ctx context.Context, request *DescribeLoadBalancingDetailRequest) (response *DescribeLoadBalancingDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancingDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancingDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancingDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginGroupRequest() (request *DescribeOriginGroupRequest) {
    request = &DescribeOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginGroup")
    
    
    return
}

func NewDescribeOriginGroupResponse() (response *DescribeOriginGroupResponse) {
    response = &DescribeOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOriginGroup
// 获取源站组信息列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOriginGroup(request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    return c.DescribeOriginGroupWithContext(context.Background(), request)
}

// DescribeOriginGroup
// 获取源站组信息列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOriginGroupWithContext(ctx context.Context, request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    if request == nil {
        request = NewDescribeOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginGroupDetailRequest() (request *DescribeOriginGroupDetailRequest) {
    request = &DescribeOriginGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginGroupDetail")
    
    
    return
}

func NewDescribeOriginGroupDetailResponse() (response *DescribeOriginGroupDetailResponse) {
    response = &DescribeOriginGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOriginGroupDetail
// 获取源站组详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOriginGroupDetail(request *DescribeOriginGroupDetailRequest) (response *DescribeOriginGroupDetailResponse, err error) {
    return c.DescribeOriginGroupDetailWithContext(context.Background(), request)
}

// DescribeOriginGroupDetail
// 获取源站组详细信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeOriginGroupDetailWithContext(ctx context.Context, request *DescribeOriginGroupDetailRequest) (response *DescribeOriginGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeOriginGroupDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginGroupDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewL7DataRequest() (request *DescribeOverviewL7DataRequest) {
    request = &DescribeOverviewL7DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOverviewL7Data")
    
    
    return
}

func NewDescribeOverviewL7DataResponse() (response *DescribeOverviewL7DataResponse) {
    response = &DescribeOverviewL7DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOverviewL7Data
// 接口已经不兼容升级，故作下线处理
//
// 
//
// 本接口（DescribeOverviewL7Data）用于查询七层监控类时序流量数据。
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7Data(request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    return c.DescribeOverviewL7DataWithContext(context.Background(), request)
}

// DescribeOverviewL7Data
// 接口已经不兼容升级，故作下线处理
//
// 
//
// 本接口（DescribeOverviewL7Data）用于查询七层监控类时序流量数据。
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7DataWithContext(ctx context.Context, request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewL7DataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewL7Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewL7DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrefetchTasksRequest() (request *DescribePrefetchTasksRequest) {
    request = &DescribePrefetchTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePrefetchTasks")
    
    
    return
}

func NewDescribePrefetchTasksResponse() (response *DescribePrefetchTasksResponse) {
    response = &DescribePrefetchTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrefetchTasks
// 查询预热任务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasks(request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    return c.DescribePrefetchTasksWithContext(context.Background(), request)
}

// DescribePrefetchTasks
// 查询预热任务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasksWithContext(ctx context.Context, request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    if request == nil {
        request = NewDescribePrefetchTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrefetchTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrefetchTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeTasks
// 查询清除缓存历史记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// 查询清除缓存历史记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRules
// 查询规则引擎规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// 查询规则引擎规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesSettingRequest() (request *DescribeRulesSettingRequest) {
    request = &DescribeRulesSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRulesSetting")
    
    
    return
}

func NewDescribeRulesSettingResponse() (response *DescribeRulesSettingResponse) {
    response = &DescribeRulesSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRulesSetting
// 返回规则引擎可应用匹配请求的设置列表及其详细建议配置信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRulesSetting(request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    return c.DescribeRulesSettingWithContext(context.Background(), request)
}

// DescribeRulesSetting
// 返回规则引擎可应用匹配请求的设置列表及其详细建议配置信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRulesSettingWithContext(ctx context.Context, request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    if request == nil {
        request = NewDescribeRulesSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRulesSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyRequest() (request *DescribeSecurityPolicyRequest) {
    request = &DescribeSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityPolicy")
    
    
    return
}

func NewDescribeSecurityPolicyResponse() (response *DescribeSecurityPolicyResponse) {
    response = &DescribeSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPolicy
// 查询安全防护配置详情
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeSecurityPolicy(request *DescribeSecurityPolicyRequest) (response *DescribeSecurityPolicyResponse, err error) {
    return c.DescribeSecurityPolicyWithContext(context.Background(), request)
}

// DescribeSecurityPolicy
// 查询安全防护配置详情
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeSecurityPolicyWithContext(ctx context.Context, request *DescribeSecurityPolicyRequest) (response *DescribeSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyListRequest() (request *DescribeSecurityPolicyListRequest) {
    request = &DescribeSecurityPolicyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityPolicyList")
    
    
    return
}

func NewDescribeSecurityPolicyListResponse() (response *DescribeSecurityPolicyListResponse) {
    response = &DescribeSecurityPolicyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPolicyList
// 查询全部安全实例
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSecurityPolicyList(request *DescribeSecurityPolicyListRequest) (response *DescribeSecurityPolicyListResponse, err error) {
    return c.DescribeSecurityPolicyListWithContext(context.Background(), request)
}

// DescribeSecurityPolicyList
// 查询全部安全实例
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSecurityPolicyListWithContext(ctx context.Context, request *DescribeSecurityPolicyListRequest) (response *DescribeSecurityPolicyListResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyManagedRulesRequest() (request *DescribeSecurityPolicyManagedRulesRequest) {
    request = &DescribeSecurityPolicyManagedRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityPolicyManagedRules")
    
    
    return
}

func NewDescribeSecurityPolicyManagedRulesResponse() (response *DescribeSecurityPolicyManagedRulesResponse) {
    response = &DescribeSecurityPolicyManagedRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPolicyManagedRules
// 分页查询门神规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) DescribeSecurityPolicyManagedRules(request *DescribeSecurityPolicyManagedRulesRequest) (response *DescribeSecurityPolicyManagedRulesResponse, err error) {
    return c.DescribeSecurityPolicyManagedRulesWithContext(context.Background(), request)
}

// DescribeSecurityPolicyManagedRules
// 分页查询门神规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) DescribeSecurityPolicyManagedRulesWithContext(ctx context.Context, request *DescribeSecurityPolicyManagedRulesRequest) (response *DescribeSecurityPolicyManagedRulesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyManagedRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyManagedRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyManagedRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyManagedRulesIdRequest() (request *DescribeSecurityPolicyManagedRulesIdRequest) {
    request = &DescribeSecurityPolicyManagedRulesIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityPolicyManagedRulesId")
    
    
    return
}

func NewDescribeSecurityPolicyManagedRulesIdResponse() (response *DescribeSecurityPolicyManagedRulesIdResponse) {
    response = &DescribeSecurityPolicyManagedRulesIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPolicyManagedRulesId
// 规则id查询门神规则详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) DescribeSecurityPolicyManagedRulesId(request *DescribeSecurityPolicyManagedRulesIdRequest) (response *DescribeSecurityPolicyManagedRulesIdResponse, err error) {
    return c.DescribeSecurityPolicyManagedRulesIdWithContext(context.Background(), request)
}

// DescribeSecurityPolicyManagedRulesId
// 规则id查询门神规则详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
func (c *Client) DescribeSecurityPolicyManagedRulesIdWithContext(ctx context.Context, request *DescribeSecurityPolicyManagedRulesIdRequest) (response *DescribeSecurityPolicyManagedRulesIdResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyManagedRulesIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyManagedRulesId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyManagedRulesIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPolicyRegionsRequest() (request *DescribeSecurityPolicyRegionsRequest) {
    request = &DescribeSecurityPolicyRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityPolicyRegions")
    
    
    return
}

func NewDescribeSecurityPolicyRegionsResponse() (response *DescribeSecurityPolicyRegionsResponse) {
    response = &DescribeSecurityPolicyRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPolicyRegions
// 查询所有地域信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeSecurityPolicyRegions(request *DescribeSecurityPolicyRegionsRequest) (response *DescribeSecurityPolicyRegionsResponse, err error) {
    return c.DescribeSecurityPolicyRegionsWithContext(context.Background(), request)
}

// DescribeSecurityPolicyRegions
// 查询所有地域信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeSecurityPolicyRegionsWithContext(ctx context.Context, request *DescribeSecurityPolicyRegionsRequest) (response *DescribeSecurityPolicyRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPolicyRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicyRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPolicyRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPortraitRulesRequest() (request *DescribeSecurityPortraitRulesRequest) {
    request = &DescribeSecurityPortraitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityPortraitRules")
    
    
    return
}

func NewDescribeSecurityPortraitRulesResponse() (response *DescribeSecurityPortraitRulesResponse) {
    response = &DescribeSecurityPortraitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPortraitRules
// 查询Bot用户画像规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSecurityPortraitRules(request *DescribeSecurityPortraitRulesRequest) (response *DescribeSecurityPortraitRulesResponse, err error) {
    return c.DescribeSecurityPortraitRulesWithContext(context.Background(), request)
}

// DescribeSecurityPortraitRules
// 查询Bot用户画像规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeSecurityPortraitRulesWithContext(ctx context.Context, request *DescribeSecurityPortraitRulesRequest) (response *DescribeSecurityPortraitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPortraitRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPortraitRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityPortraitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL4DataRequest() (request *DescribeTimingL4DataRequest) {
    request = &DescribeTimingL4DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL4Data")
    
    
    return
}

func NewDescribeTimingL4DataResponse() (response *DescribeTimingL4DataResponse) {
    response = &DescribeTimingL4DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimingL4Data
// 旧版本的接口已经不兼容，故作下线处理
//
// 
//
// 四层时序流量数据查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4Data(request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    return c.DescribeTimingL4DataWithContext(context.Background(), request)
}

// DescribeTimingL4Data
// 旧版本的接口已经不兼容，故作下线处理
//
// 
//
// 四层时序流量数据查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4DataWithContext(ctx context.Context, request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL4DataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL4Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL4DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7AnalysisDataRequest() (request *DescribeTimingL7AnalysisDataRequest) {
    request = &DescribeTimingL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7AnalysisData")
    
    
    return
}

func NewDescribeTimingL7AnalysisDataResponse() (response *DescribeTimingL7AnalysisDataResponse) {
    response = &DescribeTimingL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimingL7AnalysisData
// 查询七层数据分析类时序流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisData(request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    return c.DescribeTimingL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTimingL7AnalysisData
// 查询七层数据分析类时序流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisDataWithContext(ctx context.Context, request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7AnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7CacheDataRequest() (request *DescribeTimingL7CacheDataRequest) {
    request = &DescribeTimingL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7CacheData")
    
    
    return
}

func NewDescribeTimingL7CacheDataResponse() (response *DescribeTimingL7CacheDataResponse) {
    response = &DescribeTimingL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimingL7CacheData
// 七层查询缓存分析时序类流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheData(request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    return c.DescribeTimingL7CacheDataWithContext(context.Background(), request)
}

// DescribeTimingL7CacheData
// 七层查询缓存分析时序类流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheDataWithContext(ctx context.Context, request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7CacheDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7AnalysisDataRequest() (request *DescribeTopL7AnalysisDataRequest) {
    request = &DescribeTopL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7AnalysisData")
    
    
    return
}

func NewDescribeTopL7AnalysisDataResponse() (response *DescribeTopL7AnalysisDataResponse) {
    response = &DescribeTopL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopL7AnalysisData
// 查询top类流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7AnalysisData(request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    return c.DescribeTopL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTopL7AnalysisData
// 查询top类流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7AnalysisDataWithContext(ctx context.Context, request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7AnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7CacheDataRequest() (request *DescribeTopL7CacheDataRequest) {
    request = &DescribeTopL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7CacheData")
    
    
    return
}

func NewDescribeTopL7CacheDataResponse() (response *DescribeTopL7CacheDataResponse) {
    response = &DescribeTopL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopL7CacheData
// 七层查询缓存分析top类流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheData(request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    return c.DescribeTopL7CacheDataWithContext(context.Background(), request)
}

// DescribeTopL7CacheData
// 七层查询缓存分析top类流量数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheDataWithContext(ctx context.Context, request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7CacheDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebManagedRulesAttackEventsRequest() (request *DescribeWebManagedRulesAttackEventsRequest) {
    request = &DescribeWebManagedRulesAttackEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebManagedRulesAttackEvents")
    
    
    return
}

func NewDescribeWebManagedRulesAttackEventsResponse() (response *DescribeWebManagedRulesAttackEventsResponse) {
    response = &DescribeWebManagedRulesAttackEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebManagedRulesAttackEvents
// 查询Web托管攻击事件
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesAttackEvents(request *DescribeWebManagedRulesAttackEventsRequest) (response *DescribeWebManagedRulesAttackEventsResponse, err error) {
    return c.DescribeWebManagedRulesAttackEventsWithContext(context.Background(), request)
}

// DescribeWebManagedRulesAttackEvents
// 查询Web托管攻击事件
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesAttackEventsWithContext(ctx context.Context, request *DescribeWebManagedRulesAttackEventsRequest) (response *DescribeWebManagedRulesAttackEventsResponse, err error) {
    if request == nil {
        request = NewDescribeWebManagedRulesAttackEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebManagedRulesAttackEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebManagedRulesAttackEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebManagedRulesDataRequest() (request *DescribeWebManagedRulesDataRequest) {
    request = &DescribeWebManagedRulesDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebManagedRulesData")
    
    
    return
}

func NewDescribeWebManagedRulesDataResponse() (response *DescribeWebManagedRulesDataResponse) {
    response = &DescribeWebManagedRulesDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebManagedRulesData
// 查询waf攻击时序数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesData(request *DescribeWebManagedRulesDataRequest) (response *DescribeWebManagedRulesDataResponse, err error) {
    return c.DescribeWebManagedRulesDataWithContext(context.Background(), request)
}

// DescribeWebManagedRulesData
// 查询waf攻击时序数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebManagedRulesDataWithContext(ctx context.Context, request *DescribeWebManagedRulesDataRequest) (response *DescribeWebManagedRulesDataResponse, err error) {
    if request == nil {
        request = NewDescribeWebManagedRulesDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebManagedRulesData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebManagedRulesDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebManagedRulesLogRequest() (request *DescribeWebManagedRulesLogRequest) {
    request = &DescribeWebManagedRulesLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebManagedRulesLog")
    
    
    return
}

func NewDescribeWebManagedRulesLogResponse() (response *DescribeWebManagedRulesLogResponse) {
    response = &DescribeWebManagedRulesLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebManagedRulesLog
// 查询waf攻击日志
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebManagedRulesLog(request *DescribeWebManagedRulesLogRequest) (response *DescribeWebManagedRulesLogResponse, err error) {
    return c.DescribeWebManagedRulesLogWithContext(context.Background(), request)
}

// DescribeWebManagedRulesLog
// 查询waf攻击日志
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebManagedRulesLogWithContext(ctx context.Context, request *DescribeWebManagedRulesLogRequest) (response *DescribeWebManagedRulesLogResponse, err error) {
    if request == nil {
        request = NewDescribeWebManagedRulesLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebManagedRulesLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebManagedRulesLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebManagedRulesTopDataRequest() (request *DescribeWebManagedRulesTopDataRequest) {
    request = &DescribeWebManagedRulesTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebManagedRulesTopData")
    
    
    return
}

func NewDescribeWebManagedRulesTopDataResponse() (response *DescribeWebManagedRulesTopDataResponse) {
    response = &DescribeWebManagedRulesTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebManagedRulesTopData
// 此版本的接口已经废弃
//
// 
//
// 查询waf攻击top数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebManagedRulesTopData(request *DescribeWebManagedRulesTopDataRequest) (response *DescribeWebManagedRulesTopDataResponse, err error) {
    return c.DescribeWebManagedRulesTopDataWithContext(context.Background(), request)
}

// DescribeWebManagedRulesTopData
// 此版本的接口已经废弃
//
// 
//
// 查询waf攻击top数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebManagedRulesTopDataWithContext(ctx context.Context, request *DescribeWebManagedRulesTopDataRequest) (response *DescribeWebManagedRulesTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeWebManagedRulesTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebManagedRulesTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebManagedRulesTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebProtectionAttackEventsRequest() (request *DescribeWebProtectionAttackEventsRequest) {
    request = &DescribeWebProtectionAttackEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebProtectionAttackEvents")
    
    
    return
}

func NewDescribeWebProtectionAttackEventsResponse() (response *DescribeWebProtectionAttackEventsResponse) {
    response = &DescribeWebProtectionAttackEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebProtectionAttackEvents
// 查询web防护攻击事件
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionAttackEvents(request *DescribeWebProtectionAttackEventsRequest) (response *DescribeWebProtectionAttackEventsResponse, err error) {
    return c.DescribeWebProtectionAttackEventsWithContext(context.Background(), request)
}

// DescribeWebProtectionAttackEvents
// 查询web防护攻击事件
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionAttackEventsWithContext(ctx context.Context, request *DescribeWebProtectionAttackEventsRequest) (response *DescribeWebProtectionAttackEventsResponse, err error) {
    if request == nil {
        request = NewDescribeWebProtectionAttackEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebProtectionAttackEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebProtectionAttackEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebProtectionDataRequest() (request *DescribeWebProtectionDataRequest) {
    request = &DescribeWebProtectionDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebProtectionData")
    
    
    return
}

func NewDescribeWebProtectionDataResponse() (response *DescribeWebProtectionDataResponse) {
    response = &DescribeWebProtectionDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebProtectionData
// 查询CC防护时序数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionData(request *DescribeWebProtectionDataRequest) (response *DescribeWebProtectionDataResponse, err error) {
    return c.DescribeWebProtectionDataWithContext(context.Background(), request)
}

// DescribeWebProtectionData
// 查询CC防护时序数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeWebProtectionDataWithContext(ctx context.Context, request *DescribeWebProtectionDataRequest) (response *DescribeWebProtectionDataResponse, err error) {
    if request == nil {
        request = NewDescribeWebProtectionDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebProtectionData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebProtectionDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebProtectionLogRequest() (request *DescribeWebProtectionLogRequest) {
    request = &DescribeWebProtectionLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeWebProtectionLog")
    
    
    return
}

func NewDescribeWebProtectionLogResponse() (response *DescribeWebProtectionLogResponse) {
    response = &DescribeWebProtectionLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebProtectionLog
// 查询CC防护日志
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebProtectionLog(request *DescribeWebProtectionLogRequest) (response *DescribeWebProtectionLogResponse, err error) {
    return c.DescribeWebProtectionLogWithContext(context.Background(), request)
}

// DescribeWebProtectionLog
// 查询CC防护日志
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeWebProtectionLogWithContext(ctx context.Context, request *DescribeWebProtectionLogRequest) (response *DescribeWebProtectionLogResponse, err error) {
    if request == nil {
        request = NewDescribeWebProtectionLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebProtectionLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebProtectionLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneDDoSPolicyRequest() (request *DescribeZoneDDoSPolicyRequest) {
    request = &DescribeZoneDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneDDoSPolicy")
    
    
    return
}

func NewDescribeZoneDDoSPolicyResponse() (response *DescribeZoneDDoSPolicyResponse) {
    response = &DescribeZoneDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneDDoSPolicy
// 查询所有DDoS防护分区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeZoneDDoSPolicy(request *DescribeZoneDDoSPolicyRequest) (response *DescribeZoneDDoSPolicyResponse, err error) {
    return c.DescribeZoneDDoSPolicyWithContext(context.Background(), request)
}

// DescribeZoneDDoSPolicy
// 查询所有DDoS防护分区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeZoneDDoSPolicyWithContext(ctx context.Context, request *DescribeZoneDDoSPolicyRequest) (response *DescribeZoneDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeZoneDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneDetailsRequest() (request *DescribeZoneDetailsRequest) {
    request = &DescribeZoneDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneDetails")
    
    
    return
}

func NewDescribeZoneDetailsResponse() (response *DescribeZoneDetailsResponse) {
    response = &DescribeZoneDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneDetails
// 根据站点 ID 查询站点的详细信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeZoneDetails(request *DescribeZoneDetailsRequest) (response *DescribeZoneDetailsResponse, err error) {
    return c.DescribeZoneDetailsWithContext(context.Background(), request)
}

// DescribeZoneDetails
// 根据站点 ID 查询站点的详细信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeZoneDetailsWithContext(ctx context.Context, request *DescribeZoneDetailsRequest) (response *DescribeZoneDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeZoneDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneSettingRequest() (request *DescribeZoneSettingRequest) {
    request = &DescribeZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneSetting")
    
    
    return
}

func NewDescribeZoneSettingResponse() (response *DescribeZoneSettingResponse) {
    response = &DescribeZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneSetting
// 用于查询站点的所有配置信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZoneSetting(request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    return c.DescribeZoneSettingWithContext(context.Background(), request)
}

// DescribeZoneSetting
// 用于查询站点的所有配置信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZoneSettingWithContext(ctx context.Context, request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    if request == nil {
        request = NewDescribeZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// 用户查询用户站点信息列表，支持分页
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 用户查询用户站点信息列表，支持分页
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL7LogsRequest() (request *DownloadL7LogsRequest) {
    request = &DownloadL7LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL7Logs")
    
    
    return
}

func NewDownloadL7LogsResponse() (response *DownloadL7LogsResponse) {
    response = &DownloadL7LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadL7Logs
// 查询七层离线日志
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7Logs(request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    return c.DownloadL7LogsWithContext(context.Background(), request)
}

// DownloadL7Logs
// 查询七层离线日志
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7LogsWithContext(ctx context.Context, request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL7LogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL7Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL7LogsResponse()
    err = c.Send(request, response)
    return
}

func NewIdentifyZoneRequest() (request *IdentifyZoneRequest) {
    request = &IdentifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "IdentifyZone")
    
    
    return
}

func NewIdentifyZoneResponse() (response *IdentifyZoneResponse) {
    response = &IdentifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IdentifyZone
// 用于验证站点所有权
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZone(request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    return c.IdentifyZoneWithContext(context.Background(), request)
}

// IdentifyZone
// 用于验证站点所有权
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZoneWithContext(ctx context.Context, request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    if request == nil {
        request = NewIdentifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IdentifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewIdentifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewImportDnsRecordsRequest() (request *ImportDnsRecordsRequest) {
    request = &ImportDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ImportDnsRecords")
    
    
    return
}

func NewImportDnsRecordsResponse() (response *ImportDnsRecordsResponse) {
    response = &ImportDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportDnsRecords
// 导入 DNS 记录
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDSRVNAME = "InvalidParameterValue.InvalidSRVName"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ImportDnsRecords(request *ImportDnsRecordsRequest) (response *ImportDnsRecordsResponse, err error) {
    return c.ImportDnsRecordsWithContext(context.Background(), request)
}

// ImportDnsRecords
// 导入 DNS 记录
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDSRVNAME = "InvalidParameterValue.InvalidSRVName"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ImportDnsRecordsWithContext(ctx context.Context, request *ImportDnsRecordsRequest) (response *ImportDnsRecordsResponse, err error) {
    if request == nil {
        request = NewImportDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRequest() (request *ModifyApplicationProxyRequest) {
    request = &ModifyApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxy")
    
    
    return
}

func NewModifyApplicationProxyResponse() (response *ModifyApplicationProxyResponse) {
    response = &ModifyApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxy
// 修改应用代理
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDSRVNAME = "InvalidParameterValue.InvalidSRVName"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxy(request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    return c.ModifyApplicationProxyWithContext(context.Background(), request)
}

// ModifyApplicationProxy
// 修改应用代理
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDSRVNAME = "InvalidParameterValue.InvalidSRVName"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyWithContext(ctx context.Context, request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleRequest() (request *ModifyApplicationProxyRuleRequest) {
    request = &ModifyApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRule")
    
    
    return
}

func NewModifyApplicationProxyRuleResponse() (response *ModifyApplicationProxyRuleResponse) {
    response = &ModifyApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyRule
// 修改应用代理规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRule(request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    return c.ModifyApplicationProxyRuleWithContext(context.Background(), request)
}

// ModifyApplicationProxyRule
// 修改应用代理规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRuleWithContext(ctx context.Context, request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleStatusRequest() (request *ModifyApplicationProxyRuleStatusRequest) {
    request = &ModifyApplicationProxyRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRuleStatus")
    
    
    return
}

func NewModifyApplicationProxyRuleStatusResponse() (response *ModifyApplicationProxyRuleStatusResponse) {
    response = &ModifyApplicationProxyRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyRuleStatus
// 修改应用代理规则的状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRuleStatus(request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    return c.ModifyApplicationProxyRuleStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyRuleStatus
// 修改应用代理规则的状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRuleStatusWithContext(ctx context.Context, request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyStatusRequest() (request *ModifyApplicationProxyStatusRequest) {
    request = &ModifyApplicationProxyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyStatus")
    
    
    return
}

func NewModifyApplicationProxyStatusResponse() (response *ModifyApplicationProxyStatusResponse) {
    response = &ModifyApplicationProxyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyStatus
// 修改应用代理的状态
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatus(request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    return c.ModifyApplicationProxyStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyStatus
// 修改应用代理的状态
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatusWithContext(ctx context.Context, request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyRequest() (request *ModifyDDoSPolicyRequest) {
    request = &ModifyDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDDoSPolicy")
    
    
    return
}

func NewModifyDDoSPolicyResponse() (response *ModifyDDoSPolicyResponse) {
    response = &ModifyDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSPolicy
// 修改DDoS防护分区配置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDDoSPolicy(request *ModifyDDoSPolicyRequest) (response *ModifyDDoSPolicyResponse, err error) {
    return c.ModifyDDoSPolicyWithContext(context.Background(), request)
}

// ModifyDDoSPolicy
// 修改DDoS防护分区配置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDDoSPolicyWithContext(ctx context.Context, request *ModifyDDoSPolicyRequest) (response *ModifyDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyHostRequest() (request *ModifyDDoSPolicyHostRequest) {
    request = &ModifyDDoSPolicyHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDDoSPolicyHost")
    
    
    return
}

func NewModifyDDoSPolicyHostResponse() (response *ModifyDDoSPolicyHostResponse) {
    response = &ModifyDDoSPolicyHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSPolicyHost
// 域名DDoS高可用开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDDoSPolicyHost(request *ModifyDDoSPolicyHostRequest) (response *ModifyDDoSPolicyHostResponse, err error) {
    return c.ModifyDDoSPolicyHostWithContext(context.Background(), request)
}

// ModifyDDoSPolicyHost
// 域名DDoS高可用开关
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDDoSPolicyHostWithContext(ctx context.Context, request *ModifyDDoSPolicyHostRequest) (response *ModifyDDoSPolicyHostResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDDoSPolicyHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDDoSPolicyHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDefaultCertificateRequest() (request *ModifyDefaultCertificateRequest) {
    request = &ModifyDefaultCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDefaultCertificate")
    
    
    return
}

func NewModifyDefaultCertificateResponse() (response *ModifyDefaultCertificateResponse) {
    response = &ModifyDefaultCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDefaultCertificate
// 修改默认证书状态
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
func (c *Client) ModifyDefaultCertificate(request *ModifyDefaultCertificateRequest) (response *ModifyDefaultCertificateResponse, err error) {
    return c.ModifyDefaultCertificateWithContext(context.Background(), request)
}

// ModifyDefaultCertificate
// 修改默认证书状态
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
func (c *Client) ModifyDefaultCertificateWithContext(ctx context.Context, request *ModifyDefaultCertificateRequest) (response *ModifyDefaultCertificateResponse, err error) {
    if request == nil {
        request = NewModifyDefaultCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDefaultCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDefaultCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDnsRecordRequest() (request *ModifyDnsRecordRequest) {
    request = &ModifyDnsRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDnsRecord")
    
    
    return
}

func NewModifyDnsRecordResponse() (response *ModifyDnsRecordResponse) {
    response = &ModifyDnsRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDnsRecord
// 修改 DNS 记录
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYNAME = "InvalidParameterValue.InvalidProxyName"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyDnsRecord(request *ModifyDnsRecordRequest) (response *ModifyDnsRecordResponse, err error) {
    return c.ModifyDnsRecordWithContext(context.Background(), request)
}

// ModifyDnsRecord
// 修改 DNS 记录
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYNAME = "InvalidParameterValue.InvalidProxyName"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyDnsRecordWithContext(ctx context.Context, request *ModifyDnsRecordRequest) (response *ModifyDnsRecordResponse, err error) {
    if request == nil {
        request = NewModifyDnsRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDnsRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDnsRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDnssecRequest() (request *ModifyDnssecRequest) {
    request = &ModifyDnssecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDnssec")
    
    
    return
}

func NewModifyDnssecResponse() (response *ModifyDnssecResponse) {
    response = &ModifyDnssecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDnssec
// 修改 DNSSEC 状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDnssec(request *ModifyDnssecRequest) (response *ModifyDnssecResponse, err error) {
    return c.ModifyDnssecWithContext(context.Background(), request)
}

// ModifyDnssec
// 修改 DNSSEC 状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDnssecWithContext(ctx context.Context, request *ModifyDnssecRequest) (response *ModifyDnssecResponse, err error) {
    if request == nil {
        request = NewModifyDnssecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDnssec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDnssecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsCertificateRequest() (request *ModifyHostsCertificateRequest) {
    request = &ModifyHostsCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyHostsCertificate")
    
    
    return
}

func NewModifyHostsCertificateResponse() (response *ModifyHostsCertificateResponse) {
    response = &ModifyHostsCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHostsCertificate
// 用于修改域名证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
func (c *Client) ModifyHostsCertificate(request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    return c.ModifyHostsCertificateWithContext(context.Background(), request)
}

// ModifyHostsCertificate
// 用于修改域名证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
func (c *Client) ModifyHostsCertificateWithContext(ctx context.Context, request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    if request == nil {
        request = NewModifyHostsCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostsCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostsCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancingRequest() (request *ModifyLoadBalancingRequest) {
    request = &ModifyLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyLoadBalancing")
    
    
    return
}

func NewModifyLoadBalancingResponse() (response *ModifyLoadBalancingResponse) {
    response = &ModifyLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancing
// 修改负载均衡
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLoadBalancing(request *ModifyLoadBalancingRequest) (response *ModifyLoadBalancingResponse, err error) {
    return c.ModifyLoadBalancingWithContext(context.Background(), request)
}

// ModifyLoadBalancing
// 修改负载均衡
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLoadBalancingWithContext(ctx context.Context, request *ModifyLoadBalancingRequest) (response *ModifyLoadBalancingResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancingStatusRequest() (request *ModifyLoadBalancingStatusRequest) {
    request = &ModifyLoadBalancingStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyLoadBalancingStatus")
    
    
    return
}

func NewModifyLoadBalancingStatusResponse() (response *ModifyLoadBalancingStatusResponse) {
    response = &ModifyLoadBalancingStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancingStatus
// 修改负载均衡状态
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyLoadBalancingStatus(request *ModifyLoadBalancingStatusRequest) (response *ModifyLoadBalancingStatusResponse, err error) {
    return c.ModifyLoadBalancingStatusWithContext(context.Background(), request)
}

// ModifyLoadBalancingStatus
// 修改负载均衡状态
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyLoadBalancingStatusWithContext(ctx context.Context, request *ModifyLoadBalancingStatusRequest) (response *ModifyLoadBalancingStatusResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancingStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancingStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancingStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOriginGroupRequest() (request *ModifyOriginGroupRequest) {
    request = &ModifyOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyOriginGroup")
    
    
    return
}

func NewModifyOriginGroupResponse() (response *ModifyOriginGroupResponse) {
    response = &ModifyOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyOriginGroup
// 源站组修改
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyOriginGroup(request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    return c.ModifyOriginGroupWithContext(context.Background(), request)
}

// ModifyOriginGroup
// 源站组修改
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyOriginGroupWithContext(ctx context.Context, request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    if request == nil {
        request = NewModifyOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRule")
    
    
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRule
// 修改规则引擎规则。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDREDIRECTURLCAPTURE = "InvalidParameter.InvalidRedirectUrlCapture"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// 修改规则引擎规则。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDREDIRECTURLCAPTURE = "InvalidParameter.InvalidRedirectUrlCapture"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
func (c *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRulePriorityRequest() (request *ModifyRulePriorityRequest) {
    request = &ModifyRulePriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRulePriority")
    
    
    return
}

func NewModifyRulePriorityResponse() (response *ModifyRulePriorityResponse) {
    response = &ModifyRulePriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRulePriority
// 修改规则引擎规则优先级
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDREDIRECTURLCAPTURE = "InvalidParameter.InvalidRedirectUrlCapture"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
func (c *Client) ModifyRulePriority(request *ModifyRulePriorityRequest) (response *ModifyRulePriorityResponse, err error) {
    return c.ModifyRulePriorityWithContext(context.Background(), request)
}

// ModifyRulePriority
// 修改规则引擎规则优先级
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_INVALIDREDIRECTURLCAPTURE = "InvalidParameter.InvalidRedirectUrlCapture"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
func (c *Client) ModifyRulePriorityWithContext(ctx context.Context, request *ModifyRulePriorityRequest) (response *ModifyRulePriorityResponse, err error) {
    if request == nil {
        request = NewModifyRulePriorityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRulePriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRulePriorityResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyRequest() (request *ModifySecurityPolicyRequest) {
    request = &ModifySecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityPolicy")
    
    
    return
}

func NewModifySecurityPolicyResponse() (response *ModifySecurityPolicyResponse) {
    response = &ModifySecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityPolicy
// 修改Web&Bot安全配置
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    return c.ModifySecurityPolicyWithContext(context.Background(), request)
}

// ModifySecurityPolicy
// 修改Web&Bot安全配置
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifySecurityPolicyWithContext(ctx context.Context, request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
    request = &ModifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZone")
    
    
    return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
    response = &ModifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZone
// 用该站点信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    return c.ModifyZoneWithContext(context.Background(), request)
}

// ModifyZone
// 用该站点信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZoneWithContext(ctx context.Context, request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    if request == nil {
        request = NewModifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneCnameSpeedUpRequest() (request *ModifyZoneCnameSpeedUpRequest) {
    request = &ModifyZoneCnameSpeedUpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneCnameSpeedUp")
    
    
    return
}

func NewModifyZoneCnameSpeedUpResponse() (response *ModifyZoneCnameSpeedUpResponse) {
    response = &ModifyZoneCnameSpeedUpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneCnameSpeedUp
// 开启，关闭 CNAME 加速
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyZoneCnameSpeedUp(request *ModifyZoneCnameSpeedUpRequest) (response *ModifyZoneCnameSpeedUpResponse, err error) {
    return c.ModifyZoneCnameSpeedUpWithContext(context.Background(), request)
}

// ModifyZoneCnameSpeedUp
// 开启，关闭 CNAME 加速
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyZoneCnameSpeedUpWithContext(ctx context.Context, request *ModifyZoneCnameSpeedUpRequest) (response *ModifyZoneCnameSpeedUpResponse, err error) {
    if request == nil {
        request = NewModifyZoneCnameSpeedUpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneCnameSpeedUp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneCnameSpeedUpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneSettingRequest() (request *ModifyZoneSettingRequest) {
    request = &ModifyZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneSetting")
    
    
    return
}

func NewModifyZoneSettingResponse() (response *ModifyZoneSettingResponse) {
    response = &ModifyZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneSetting
// 用于修改站点配置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSetting(request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    return c.ModifyZoneSettingWithContext(context.Background(), request)
}

// ModifyZoneSetting
// 用于修改站点配置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSettingWithContext(ctx context.Context, request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    if request == nil {
        request = NewModifyZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneStatusRequest() (request *ModifyZoneStatusRequest) {
    request = &ModifyZoneStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneStatus")
    
    
    return
}

func NewModifyZoneStatusResponse() (response *ModifyZoneStatusResponse) {
    response = &ModifyZoneStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneStatus
// 用于开启，关闭站点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZoneStatus(request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    return c.ModifyZoneStatusWithContext(context.Background(), request)
}

// ModifyZoneStatus
// 用于开启，关闭站点
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZoneStatusWithContext(ctx context.Context, request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    if request == nil {
        request = NewModifyZoneStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneStatusResponse()
    err = c.Send(request, response)
    return
}

func NewReclaimZoneRequest() (request *ReclaimZoneRequest) {
    request = &ReclaimZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ReclaimZone")
    
    
    return
}

func NewReclaimZoneResponse() (response *ReclaimZoneResponse) {
    response = &ReclaimZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReclaimZone
// 站点被其他用户接入后，验证了站点所有权之后，可以找回该站点
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ReclaimZone(request *ReclaimZoneRequest) (response *ReclaimZoneResponse, err error) {
    return c.ReclaimZoneWithContext(context.Background(), request)
}

// ReclaimZone
// 站点被其他用户接入后，验证了站点所有权之后，可以找回该站点
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ReclaimZoneWithContext(ctx context.Context, request *ReclaimZoneRequest) (response *ReclaimZoneResponse, err error) {
    if request == nil {
        request = NewReclaimZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReclaimZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewReclaimZoneResponse()
    err = c.Send(request, response)
    return
}

func NewScanDnsRecordsRequest() (request *ScanDnsRecordsRequest) {
    request = &ScanDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ScanDnsRecords")
    
    
    return
}

func NewScanDnsRecordsResponse() (response *ScanDnsRecordsResponse) {
    response = &ScanDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanDnsRecords
// 扫描站点历史解析记录
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ScanDnsRecords(request *ScanDnsRecordsRequest) (response *ScanDnsRecordsResponse, err error) {
    return c.ScanDnsRecordsWithContext(context.Background(), request)
}

// ScanDnsRecords
// 扫描站点历史解析记录
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ScanDnsRecordsWithContext(ctx context.Context, request *ScanDnsRecordsRequest) (response *ScanDnsRecordsResponse, err error) {
    if request == nil {
        request = NewScanDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanDnsRecordsResponse()
    err = c.Send(request, response)
    return
}
