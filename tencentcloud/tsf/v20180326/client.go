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

package v20180326

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-26"

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


func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
    request = &AddClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AddClusterInstances")
    
    
    return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
    response = &AddClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddClusterInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCERESETERROR = "FailedOperation.InstanceResetError"
//  FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"
//  RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
    if request == nil {
        request = NewAddClusterInstancesRequest()
    }
    
    response = NewAddClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// AddClusterInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCERESETERROR = "FailedOperation.InstanceResetError"
//  FAILEDOPERATION_INSTANCERESETTIMEOUT = "FailedOperation.InstanceResetTimeout"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TKECLUSTERQUERYFAILED = "FailedOperation.TkeClusterQueryFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_INSTANCEINVALIDIMAGE = "InvalidParameterValue.InstanceInvalidImage"
//  RESOURCEINUSE_INSTANCEHASBEENUSED = "ResourceInUse.InstanceHasBeenUsed"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCENOTFOUND_TKECLUSTERNOTEXISTS = "ResourceNotFound.TkeClusterNotExists"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) AddClusterInstancesWithContext(ctx context.Context, request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
    if request == nil {
        request = NewAddClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewAddClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAddInstancesRequest() (request *AddInstancesRequest) {
    request = &AddInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "AddInstances")
    
    
    return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
    response = &AddInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

// AddInstances
// 添加云主机节点至TSF集群
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
func (c *Client) AddInstancesWithContext(ctx context.Context, request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewBindApiGroupRequest() (request *BindApiGroupRequest) {
    request = &BindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "BindApiGroup")
    
    
    return
}

func NewBindApiGroupResponse() (response *BindApiGroupResponse) {
    response = &BindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindApiGroup
// 网关与API分组批量绑定
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) BindApiGroup(request *BindApiGroupRequest) (response *BindApiGroupResponse, err error) {
    if request == nil {
        request = NewBindApiGroupRequest()
    }
    
    response = NewBindApiGroupResponse()
    err = c.Send(request, response)
    return
}

// BindApiGroup
// 网关与API分组批量绑定
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) BindApiGroupWithContext(ctx context.Context, request *BindApiGroupRequest) (response *BindApiGroupResponse, err error) {
    if request == nil {
        request = NewBindApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindPluginRequest() (request *BindPluginRequest) {
    request = &BindPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "BindPlugin")
    
    
    return
}

func NewBindPluginResponse() (response *BindPluginResponse) {
    response = &BindPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindPlugin
// 插件与网关分组/API批量绑定
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) BindPlugin(request *BindPluginRequest) (response *BindPluginResponse, err error) {
    if request == nil {
        request = NewBindPluginRequest()
    }
    
    response = NewBindPluginResponse()
    err = c.Send(request, response)
    return
}

// BindPlugin
// 插件与网关分组/API批量绑定
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) BindPluginWithContext(ctx context.Context, request *BindPluginRequest) (response *BindPluginResponse, err error) {
    if request == nil {
        request = NewBindPluginRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindPluginResponse()
    err = c.Send(request, response)
    return
}

func NewChangeApiUsableStatusRequest() (request *ChangeApiUsableStatusRequest) {
    request = &ChangeApiUsableStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ChangeApiUsableStatus")
    
    
    return
}

func NewChangeApiUsableStatusResponse() (response *ChangeApiUsableStatusResponse) {
    response = &ChangeApiUsableStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeApiUsableStatus
// 启用或禁用API
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) ChangeApiUsableStatus(request *ChangeApiUsableStatusRequest) (response *ChangeApiUsableStatusResponse, err error) {
    if request == nil {
        request = NewChangeApiUsableStatusRequest()
    }
    
    response = NewChangeApiUsableStatusResponse()
    err = c.Send(request, response)
    return
}

// ChangeApiUsableStatus
// 启用或禁用API
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCOMMONERROR = "InternalError.GatewayCommonError"
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) ChangeApiUsableStatusWithContext(ctx context.Context, request *ChangeApiUsableStatusRequest) (response *ChangeApiUsableStatusResponse, err error) {
    if request == nil {
        request = NewChangeApiUsableStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewChangeApiUsableStatusResponse()
    err = c.Send(request, response)
    return
}

func NewContinueRunFailedTaskBatchRequest() (request *ContinueRunFailedTaskBatchRequest) {
    request = &ContinueRunFailedTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ContinueRunFailedTaskBatch")
    
    
    return
}

func NewContinueRunFailedTaskBatchResponse() (response *ContinueRunFailedTaskBatchResponse) {
    response = &ContinueRunFailedTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ContinueRunFailedTaskBatch
// 对执行失败的任务批次执行续跑
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ContinueRunFailedTaskBatch(request *ContinueRunFailedTaskBatchRequest) (response *ContinueRunFailedTaskBatchResponse, err error) {
    if request == nil {
        request = NewContinueRunFailedTaskBatchRequest()
    }
    
    response = NewContinueRunFailedTaskBatchResponse()
    err = c.Send(request, response)
    return
}

// ContinueRunFailedTaskBatch
// 对执行失败的任务批次执行续跑
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ContinueRunFailedTaskBatchWithContext(ctx context.Context, request *ContinueRunFailedTaskBatchRequest) (response *ContinueRunFailedTaskBatchResponse, err error) {
    if request == nil {
        request = NewContinueRunFailedTaskBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewContinueRunFailedTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAllGatewayApiAsyncRequest() (request *CreateAllGatewayApiAsyncRequest) {
    request = &CreateAllGatewayApiAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateAllGatewayApiAsync")
    
    
    return
}

func NewCreateAllGatewayApiAsyncResponse() (response *CreateAllGatewayApiAsyncResponse) {
    response = &CreateAllGatewayApiAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAllGatewayApiAsync
// 一键导入API分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) CreateAllGatewayApiAsync(request *CreateAllGatewayApiAsyncRequest) (response *CreateAllGatewayApiAsyncResponse, err error) {
    if request == nil {
        request = NewCreateAllGatewayApiAsyncRequest()
    }
    
    response = NewCreateAllGatewayApiAsyncResponse()
    err = c.Send(request, response)
    return
}

// CreateAllGatewayApiAsync
// 一键导入API分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) CreateAllGatewayApiAsyncWithContext(ctx context.Context, request *CreateAllGatewayApiAsyncRequest) (response *CreateAllGatewayApiAsyncResponse, err error) {
    if request == nil {
        request = NewCreateAllGatewayApiAsyncRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAllGatewayApiAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiGroupRequest() (request *CreateApiGroupRequest) {
    request = &CreateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiGroup")
    
    
    return
}

func NewCreateApiGroupResponse() (response *CreateApiGroupResponse) {
    response = &CreateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApiGroup
// 创建API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApiGroup(request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
    if request == nil {
        request = NewCreateApiGroupRequest()
    }
    
    response = NewCreateApiGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateApiGroup
// 创建API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApiGroupWithContext(ctx context.Context, request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
    if request == nil {
        request = NewCreateApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApiRateLimitRuleRequest() (request *CreateApiRateLimitRuleRequest) {
    request = &CreateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApiRateLimitRule")
    
    
    return
}

func NewCreateApiRateLimitRuleResponse() (response *CreateApiRateLimitRuleResponse) {
    response = &CreateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApiRateLimitRule
// 创建API限流规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateApiRateLimitRule(request *CreateApiRateLimitRuleRequest) (response *CreateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewCreateApiRateLimitRuleRequest()
    }
    
    response = NewCreateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

// CreateApiRateLimitRule
// 创建API限流规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateApiRateLimitRuleWithContext(ctx context.Context, request *CreateApiRateLimitRuleRequest) (response *CreateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewCreateApiRateLimitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateApplication")
    
    
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
//  FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  INTERNALERROR_APPLICATIONMASTERFEIGNERROR = "InternalError.ApplicationMasterFeignError"
//  INTERNALERROR_APPLICATIONMASTERNUKNOWNERROR = "InternalError.ApplicationMasterNuknownError"
//  INTERNALERROR_APPLICATIONSCALABLEINITERROR = "InternalError.ApplicationScalableInitError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_APPLICATIONMICROTYPEINVALID = "InvalidParameterValue.ApplicationMicroTypeInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEEXIST = "InvalidParameterValue.ApplicationNameExist"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMENULL = "InvalidParameterValue.ApplicationNameNull"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEREGXINVALID = "InvalidParameterValue.ApplicationNameRegxInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONTYPEINVALID = "InvalidParameterValue.ApplicationTypeInvalid"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  MISSINGPARAMETER_APPLICATIONTYPENULL = "MissingParameter.ApplicationTypeNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

// CreateApplication
// 创建应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONCREATEESATUHERROR = "FailedOperation.ApplicationCreateEsAtuhError"
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  INTERNALERROR_APPLICATIONMASTERFEIGNERROR = "InternalError.ApplicationMasterFeignError"
//  INTERNALERROR_APPLICATIONMASTERNUKNOWNERROR = "InternalError.ApplicationMasterNuknownError"
//  INTERNALERROR_APPLICATIONSCALABLEINITERROR = "InternalError.ApplicationScalableInitError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_APPLICATIONMICROTYPEINVALID = "InvalidParameterValue.ApplicationMicroTypeInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEEXIST = "InvalidParameterValue.ApplicationNameExist"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMENULL = "InvalidParameterValue.ApplicationNameNull"
//  INVALIDPARAMETERVALUE_APPLICATIONNAMEREGXINVALID = "InvalidParameterValue.ApplicationNameRegxInvalid"
//  INVALIDPARAMETERVALUE_APPLICATIONTYPEINVALID = "InvalidParameterValue.ApplicationTypeInvalid"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  MISSINGPARAMETER_APPLICATIONTYPENULL = "MissingParameter.ApplicationTypeNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_TKECLUSTERCREATEFAILED = "FailedOperation.TkeClusterCreateFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERCOMMONERROR = "InternalError.ClusterCommonError"
//  INTERNALERROR_CLUSTERMASTERFEIGNERROR = "InternalError.ClusterMasterFeignError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CLUSTERCIDRCONFLICT = "InvalidParameterValue.ClusterCidrConflict"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEEXIST = "InvalidParameterValue.ClusterNameExist"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEREQUIRED = "InvalidParameterValue.ClusterNameRequired"
//  INVALIDPARAMETERVALUE_CLUSTERREGIONINVALID = "InvalidParameterValue.ClusterRegionInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEINVALID = "InvalidParameterValue.ClusterTypeInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERZONEINVALID = "InvalidParameterValue.ClusterZoneInvalid"
//  LIMITEXCEEDED_TKECLUSTERNUMBEREXCEEDLIMIT = "LimitExceeded.TkeClusterNumberExceedLimit"
//  MISSINGPARAMETER_CLUSTERSUBNETREQUIRED = "MissingParameter.ClusterSubnetRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CLUSTERVPCNOTEXIST = "ResourceNotFound.ClusterVpcNotExist"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERCREATEVPCFAIL = "FailedOperation.ClusterCreateVpcFail"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_TKECLUSTERCREATEFAILED = "FailedOperation.TkeClusterCreateFailed"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERCOMMONERROR = "InternalError.ClusterCommonError"
//  INTERNALERROR_CLUSTERMASTERFEIGNERROR = "InternalError.ClusterMasterFeignError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CLUSTERCIDRCONFLICT = "InvalidParameterValue.ClusterCidrConflict"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEEXIST = "InvalidParameterValue.ClusterNameExist"
//  INVALIDPARAMETERVALUE_CLUSTERNAMEREQUIRED = "InvalidParameterValue.ClusterNameRequired"
//  INVALIDPARAMETERVALUE_CLUSTERREGIONINVALID = "InvalidParameterValue.ClusterRegionInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEINVALID = "InvalidParameterValue.ClusterTypeInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERZONEINVALID = "InvalidParameterValue.ClusterZoneInvalid"
//  LIMITEXCEEDED_TKECLUSTERNUMBEREXCEEDLIMIT = "LimitExceeded.TkeClusterNumberExceedLimit"
//  MISSINGPARAMETER_CLUSTERSUBNETREQUIRED = "MissingParameter.ClusterSubnetRequired"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CLUSTERVPCNOTEXIST = "ResourceNotFound.ClusterVpcNotExist"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
    request = &CreateConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateConfig")
    
    
    return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
    response = &CreateConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConfig
// 创建配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"
//  MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    if request == nil {
        request = NewCreateConfigRequest()
    }
    
    response = NewCreateConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateConfig
// 创建配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGNAMEINVALID = "InvalidParameterValue.ConfigNameInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONDESCINVALID = "InvalidParameterValue.ConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_APPLICATIONIDREQUIRED = "MissingParameter.ApplicationIdRequired"
//  MISSINGPARAMETER_CONFIGNAMEREQUIRED = "MissingParameter.ConfigNameRequired"
//  MISSINGPARAMETER_CONFIGTYPEREQUIRED = "MissingParameter.ConfigTypeRequired"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  MISSINGPARAMETER_CONFIGVERSIONREQUIRED = "MissingParameter.ConfigVersionRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateConfigWithContext(ctx context.Context, request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    if request == nil {
        request = NewCreateConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContainGroupRequest() (request *CreateContainGroupRequest) {
    request = &CreateContainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateContainGroup")
    
    
    return
}

func NewCreateContainGroupResponse() (response *CreateContainGroupResponse) {
    response = &CreateContainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateContainGroup
// 创建容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMELEGNTH = "InvalidParameterValue.ContainergroupGroupnameLegnth"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMEREGEXMATCHFALSE = "InvalidParameterValue.ContainergroupGroupnameRegexMatchFalse"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTINVALID = "InvalidParameterValue.ContainergroupPortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
func (c *Client) CreateContainGroup(request *CreateContainGroupRequest) (response *CreateContainGroupResponse, err error) {
    if request == nil {
        request = NewCreateContainGroupRequest()
    }
    
    response = NewCreateContainGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateContainGroup
// 创建容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMELEGNTH = "InvalidParameterValue.ContainergroupGroupnameLegnth"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPNAMEREGEXMATCHFALSE = "InvalidParameterValue.ContainergroupGroupnameRegexMatchFalse"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTINVALID = "InvalidParameterValue.ContainergroupPortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLINVALID = "InvalidParameterValue.ContainergroupProtocolInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
func (c *Client) CreateContainGroupWithContext(ctx context.Context, request *CreateContainGroupRequest) (response *CreateContainGroupResponse, err error) {
    if request == nil {
        request = NewCreateContainGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateContainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileConfigRequest() (request *CreateFileConfigRequest) {
    request = &CreateFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateFileConfig")
    
    
    return
}

func NewCreateFileConfigResponse() (response *CreateFileConfigResponse) {
    response = &CreateFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFileConfig
// 创建文件配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"
//  INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"
func (c *Client) CreateFileConfig(request *CreateFileConfigRequest) (response *CreateFileConfigResponse, err error) {
    if request == nil {
        request = NewCreateFileConfigRequest()
    }
    
    response = NewCreateFileConfigResponse()
    err = c.Send(request, response)
    return
}

// CreateFileConfig
// 创建文件配置项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGCREATEFAILED = "FailedOperation.ConfigCreateFailed"
//  INVALIDPARAMETERVALUE_APPLICATIONNOTEXISTS = "InvalidParameterValue.ApplicationNotExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTS = "InvalidParameterValue.FileConfigExists"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATH = "InvalidParameterValue.FileConfigExistsPath"
//  INVALIDPARAMETERVALUE_FILECONFIGEXISTSPATHOTHER = "InvalidParameterValue.FileConfigExistsPathOther"
//  INVALIDPARAMETERVALUE_FILECONFIGFILEPATHINVALID = "InvalidParameterValue.FileConfigFilePathInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGNAMEINVALID = "InvalidParameterValue.FileConfigNameInvalid"
//  INVALIDPARAMETERVALUE_FILECONFIGVERSIONDESCINVALID = "InvalidParameterValue.FileConfigVersionDescInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_FILECONFIGFILEVALUEREQUIRED = "MissingParameter.FileConfigFileValueRequired"
func (c *Client) CreateFileConfigWithContext(ctx context.Context, request *CreateFileConfigRequest) (response *CreateFileConfigResponse, err error) {
    if request == nil {
        request = NewCreateFileConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatewayApiRequest() (request *CreateGatewayApiRequest) {
    request = &CreateGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayApi")
    
    
    return
}

func NewCreateGatewayApiResponse() (response *CreateGatewayApiResponse) {
    response = &CreateGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGatewayApi
// 批量导入API至api分组(也支持新建API到分组)
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNSUPPORTEDOPERATION_GATEWAYTOOMANYREQUESTPARAMETER = "UnsupportedOperation.GatewayTooManyRequestParameter"
func (c *Client) CreateGatewayApi(request *CreateGatewayApiRequest) (response *CreateGatewayApiResponse, err error) {
    if request == nil {
        request = NewCreateGatewayApiRequest()
    }
    
    response = NewCreateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

// CreateGatewayApi
// 批量导入API至api分组(也支持新建API到分组)
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNSUPPORTEDOPERATION_GATEWAYTOOMANYREQUESTPARAMETER = "UnsupportedOperation.GatewayTooManyRequestParameter"
func (c *Client) CreateGatewayApiWithContext(ctx context.Context, request *CreateGatewayApiRequest) (response *CreateGatewayApiResponse, err error) {
    if request == nil {
        request = NewCreateGatewayApiRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateGroup")
    
    
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGroup
// 创建虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"
//  INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"
//  MISSINGPARAMETER_GROUPAPPLICATIONNULL = "MissingParameter.GroupApplicationNull"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_CVMCAEMASTERRESOURCENOTFOUND = "ResourceNotFound.CvmcaeMasterResourceNotFound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateGroup
// 创建虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_GROUPMASTERNUKNOWNERROR = "InternalError.GroupMasterNuknownError"
//  INVALIDPARAMETERVALUE_GROUPCLUSTERNAMESPACENOTBOUND = "InvalidParameterValue.GroupClusterNamespaceNotBound"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"
//  INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"
//  MISSINGPARAMETER_GROUPAPPLICATIONNULL = "MissingParameter.GroupApplicationNull"
//  RESOURCENOTFOUND_CONTAINERGROUPCLUSTERNOTFOUND = "ResourceNotFound.ContainergroupClusterNotfound"
//  RESOURCENOTFOUND_CVMCAEMASTERRESOURCENOTFOUND = "ResourceNotFound.CvmcaeMasterResourceNotFound"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaneRequest() (request *CreateLaneRequest) {
    request = &CreateLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateLane")
    
    
    return
}

func NewCreateLaneResponse() (response *CreateLaneResponse) {
    response = &CreateLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLane
// 创建泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLane(request *CreateLaneRequest) (response *CreateLaneResponse, err error) {
    if request == nil {
        request = NewCreateLaneRequest()
    }
    
    response = NewCreateLaneResponse()
    err = c.Send(request, response)
    return
}

// CreateLane
// 创建泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANEINFORELEASEMESHFAILED = "FailedOperation.LaneInfoReleaseMeshFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLaneWithContext(ctx context.Context, request *CreateLaneRequest) (response *CreateLaneResponse, err error) {
    if request == nil {
        request = NewCreateLaneRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLaneResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaneRuleRequest() (request *CreateLaneRuleRequest) {
    request = &CreateLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateLaneRule")
    
    
    return
}

func NewCreateLaneRuleResponse() (response *CreateLaneRuleResponse) {
    response = &CreateLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLaneRule
// 创建泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLaneRule(request *CreateLaneRuleRequest) (response *CreateLaneRuleResponse, err error) {
    if request == nil {
        request = NewCreateLaneRuleRequest()
    }
    
    response = NewCreateLaneRuleResponse()
    err = c.Send(request, response)
    return
}

// CreateLaneRule
// 创建泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) CreateLaneRuleWithContext(ctx context.Context, request *CreateLaneRuleRequest) (response *CreateLaneRuleResponse, err error) {
    if request == nil {
        request = NewCreateLaneRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMicroserviceRequest() (request *CreateMicroserviceRequest) {
    request = &CreateMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroservice")
    
    
    return
}

func NewCreateMicroserviceResponse() (response *CreateMicroserviceResponse) {
    response = &CreateMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMicroservice
// 新增微服务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SERVICEDESCLENGTH = "InvalidParameterValue.ServiceDescLength"
//  INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateMicroservice(request *CreateMicroserviceRequest) (response *CreateMicroserviceResponse, err error) {
    if request == nil {
        request = NewCreateMicroserviceRequest()
    }
    
    response = NewCreateMicroserviceResponse()
    err = c.Send(request, response)
    return
}

// CreateMicroservice
// 新增微服务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SERVICEDESCLENGTH = "InvalidParameterValue.ServiceDescLength"
//  INVALIDPARAMETERVALUE_SERVICENAMEREPEATED = "InvalidParameterValue.ServiceNameRepeated"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateMicroserviceWithContext(ctx context.Context, request *CreateMicroserviceRequest) (response *CreateMicroserviceResponse, err error) {
    if request == nil {
        request = NewCreateMicroserviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateNamespace")
    
    
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNamespace
// 创建命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_KUBERNETESAPICREATENAMESPACESERROR = "InternalError.KubernetesApiCreateNamespacesError"
//  INTERNALERROR_KUBERNETESAPICREATESECRETERROR = "InternalError.KubernetesApiCreateSecretError"
//  INVALIDPARAMETERVALUE_GLOBALNAMESPACENAMEEXIST = "InvalidParameterValue.GlobalNamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACEALREADYBINDCLUSTER = "InvalidParameterValue.NamespaceAlreadyBindCluster"
//  INVALIDPARAMETERVALUE_NAMESPACEDESCINVALID = "InvalidParameterValue.NamespaceDescInvalid"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEEXIST = "InvalidParameterValue.NamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEINVALID = "InvalidParameterValue.NamespaceNameInvalid"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

// CreateNamespace
// 创建命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACECREATEFAILED = "FailedOperation.NamespaceCreateFailed"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CPCLUSTERUNAVAILABLE = "InternalError.CpClusterUnavailable"
//  INTERNALERROR_KUBERNETESAPICREATENAMESPACESERROR = "InternalError.KubernetesApiCreateNamespacesError"
//  INTERNALERROR_KUBERNETESAPICREATESECRETERROR = "InternalError.KubernetesApiCreateSecretError"
//  INVALIDPARAMETERVALUE_GLOBALNAMESPACENAMEEXIST = "InvalidParameterValue.GlobalNamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACEALREADYBINDCLUSTER = "InvalidParameterValue.NamespaceAlreadyBindCluster"
//  INVALIDPARAMETERVALUE_NAMESPACEDESCINVALID = "InvalidParameterValue.NamespaceDescInvalid"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEEXIST = "InvalidParameterValue.NamespaceNameExist"
//  INVALIDPARAMETERVALUE_NAMESPACENAMEINVALID = "InvalidParameterValue.NamespaceNameInvalid"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateNamespaceWithContext(ctx context.Context, request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePathRewritesRequest() (request *CreatePathRewritesRequest) {
    request = &CreatePathRewritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePathRewrites")
    
    
    return
}

func NewCreatePathRewritesResponse() (response *CreatePathRewritesResponse) {
    response = &CreatePathRewritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePathRewrites
// 创建路径重写
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreatePathRewrites(request *CreatePathRewritesRequest) (response *CreatePathRewritesResponse, err error) {
    if request == nil {
        request = NewCreatePathRewritesRequest()
    }
    
    response = NewCreatePathRewritesResponse()
    err = c.Send(request, response)
    return
}

// CreatePathRewrites
// 创建路径重写
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreatePathRewritesWithContext(ctx context.Context, request *CreatePathRewritesRequest) (response *CreatePathRewritesResponse, err error) {
    if request == nil {
        request = NewCreatePathRewritesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePathRewritesResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublicConfigRequest() (request *CreatePublicConfigRequest) {
    request = &CreatePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreatePublicConfig")
    
    
    return
}

func NewCreatePublicConfigResponse() (response *CreatePublicConfigResponse) {
    response = &CreatePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePublicConfig
// 创建公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreatePublicConfig(request *CreatePublicConfigRequest) (response *CreatePublicConfigResponse, err error) {
    if request == nil {
        request = NewCreatePublicConfigRequest()
    }
    
    response = NewCreatePublicConfigResponse()
    err = c.Send(request, response)
    return
}

// CreatePublicConfig
// 创建公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGEXISTS = "InvalidParameterValue.ConfigExists"
//  INVALIDPARAMETERVALUE_CONFIGVALUEFORMATINVALID = "InvalidParameterValue.ConfigValueFormatInvalid"
//  INVALIDPARAMETERVALUE_CONFIGVERSIONINVALID = "InvalidParameterValue.ConfigVersionInvalid"
//  MISSINGPARAMETER_CONFIGVALUEREQUIRED = "MissingParameter.ConfigValueRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreatePublicConfigWithContext(ctx context.Context, request *CreatePublicConfigRequest) (response *CreatePublicConfigResponse, err error) {
    if request == nil {
        request = NewCreatePublicConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreatePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRepositoryRequest() (request *CreateRepositoryRequest) {
    request = &CreateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateRepository")
    
    
    return
}

func NewCreateRepositoryResponse() (response *CreateRepositoryResponse) {
    response = &CreateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRepository
// 创建仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
func (c *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryRequest()
    }
    
    response = NewCreateRepositoryResponse()
    err = c.Send(request, response)
    return
}

// CreateRepository
// 创建仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
func (c *Client) CreateRepositoryWithContext(ctx context.Context, request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerlessGroupRequest() (request *CreateServerlessGroupRequest) {
    request = &CreateServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateServerlessGroup")
    
    
    return
}

func NewCreateServerlessGroupResponse() (response *CreateServerlessGroupResponse) {
    response = &CreateServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServerlessGroup
// 创建Serverless部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"
//  INVALIDPARAMETERVALUE_GROUPNAMENULL = "InvalidParameterValue.GroupNameNull"
//  INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"
//  INVALIDPARAMETERVALUE_GROUPPKGNULL = "InvalidParameterValue.GroupPkgNull"
func (c *Client) CreateServerlessGroup(request *CreateServerlessGroupRequest) (response *CreateServerlessGroupResponse, err error) {
    if request == nil {
        request = NewCreateServerlessGroupRequest()
    }
    
    response = NewCreateServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateServerlessGroup
// 创建Serverless部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPNAMEEXIST = "InvalidParameterValue.GroupNameExist"
//  INVALIDPARAMETERVALUE_GROUPNAMELENGTH = "InvalidParameterValue.GroupNameLength"
//  INVALIDPARAMETERVALUE_GROUPNAMENULL = "InvalidParameterValue.GroupNameNull"
//  INVALIDPARAMETERVALUE_GROUPNAMEREGXMISMATCH = "InvalidParameterValue.GroupNameRegxMismatch"
//  INVALIDPARAMETERVALUE_GROUPPKGNULL = "InvalidParameterValue.GroupPkgNull"
func (c *Client) CreateServerlessGroupWithContext(ctx context.Context, request *CreateServerlessGroupRequest) (response *CreateServerlessGroupResponse, err error) {
    if request == nil {
        request = NewCreateServerlessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTask
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

// CreateTask
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFlowRequest() (request *CreateTaskFlowRequest) {
    request = &CreateTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateTaskFlow")
    
    
    return
}

func NewCreateTaskFlowResponse() (response *CreateTaskFlowResponse) {
    response = &CreateTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTaskFlow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTaskFlow(request *CreateTaskFlowRequest) (response *CreateTaskFlowResponse, err error) {
    if request == nil {
        request = NewCreateTaskFlowRequest()
    }
    
    response = NewCreateTaskFlowResponse()
    err = c.Send(request, response)
    return
}

// CreateTaskFlow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKCREATEERROR = "FailedOperation.TaskCreateError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) CreateTaskFlowWithContext(ctx context.Context, request *CreateTaskFlowRequest) (response *CreateTaskFlowResponse, err error) {
    if request == nil {
        request = NewCreateTaskFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnitRuleRequest() (request *CreateUnitRuleRequest) {
    request = &CreateUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "CreateUnitRule")
    
    
    return
}

func NewCreateUnitRuleResponse() (response *CreateUnitRuleResponse) {
    response = &CreateUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUnitRule
// 创建单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitRule(request *CreateUnitRuleRequest) (response *CreateUnitRuleResponse, err error) {
    if request == nil {
        request = NewCreateUnitRuleRequest()
    }
    
    response = NewCreateUnitRuleResponse()
    err = c.Send(request, response)
    return
}

// CreateUnitRule
// 创建单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) CreateUnitRuleWithContext(ctx context.Context, request *CreateUnitRuleRequest) (response *CreateUnitRuleResponse, err error) {
    if request == nil {
        request = NewCreateUnitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiGroupRequest() (request *DeleteApiGroupRequest) {
    request = &DeleteApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiGroup")
    
    
    return
}

func NewDeleteApiGroupResponse() (response *DeleteApiGroupResponse) {
    response = &DeleteApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApiGroup
// 删除Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
    if request == nil {
        request = NewDeleteApiGroupRequest()
    }
    
    response = NewDeleteApiGroupResponse()
    err = c.Send(request, response)
    return
}

// DeleteApiGroup
// 删除Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DeleteApiGroupWithContext(ctx context.Context, request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
    if request == nil {
        request = NewDeleteApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
    request = &DeleteApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteApplication")
    
    
    return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
    response = &DeleteApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplication
// 删除应用
//
// 可能返回的错误码:
//  INTERNALERROR_APPLICATIONREPODELETEPKG = "InternalError.ApplicationRepoDeletePkg"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"
//  RESOURCEINUSE_APPLICATIONCANNOTDELETE = "ResourceInUse.ApplicationCannotDelete"
//  RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

// DeleteApplication
// 删除应用
//
// 可能返回的错误码:
//  INTERNALERROR_APPLICATIONREPODELETEPKG = "InternalError.ApplicationRepoDeletePkg"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"
//  RESOURCEINUSE_APPLICATIONCANNOTDELETE = "ResourceInUse.ApplicationCannotDelete"
//  RESOURCEINUSE_CVMCAEMASTERCANNOTDELETE = "ResourceInUse.CvmcaeMasterCannotDelete"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
    request = &DeleteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfig")
    
    
    return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
    response = &DeleteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteConfig
// 删除配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRequest()
    }
    
    response = NewDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

// DeleteConfig
// 删除配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteConfigWithContext(ctx context.Context, request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContainerGroupRequest() (request *DeleteContainerGroupRequest) {
    request = &DeleteContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteContainerGroup")
    
    
    return
}

func NewDeleteContainerGroupResponse() (response *DeleteContainerGroupResponse) {
    response = &DeleteContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteContainerGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (response *DeleteContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteContainerGroupRequest()
    }
    
    response = NewDeleteContainerGroupResponse()
    err = c.Send(request, response)
    return
}

// DeleteContainerGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteContainerGroupWithContext(ctx context.Context, request *DeleteContainerGroupRequest) (response *DeleteContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteContainerGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPDELETECLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupDeleteClusterTypeMismatch"
//  RESOURCEINUSE_GROUPCANNOTDELETE = "ResourceInUse.GroupCannotDelete"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

// DeleteGroup
// 删除容器部署组
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPDELETECLUSTERTYPEMISMATCH = "InvalidParameterValue.GroupDeleteClusterTypeMismatch"
//  RESOURCEINUSE_GROUPCANNOTDELETE = "ResourceInUse.GroupCannotDelete"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageTagsRequest() (request *DeleteImageTagsRequest) {
    request = &DeleteImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteImageTags")
    
    
    return
}

func NewDeleteImageTagsResponse() (response *DeleteImageTagsResponse) {
    response = &DeleteImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImageTags
// 批量删除镜像版本
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPIMAGETAGISINUSE = "InvalidParameterValue.ContainerGroupImageTagIsInUse"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
func (c *Client) DeleteImageTags(request *DeleteImageTagsRequest) (response *DeleteImageTagsResponse, err error) {
    if request == nil {
        request = NewDeleteImageTagsRequest()
    }
    
    response = NewDeleteImageTagsResponse()
    err = c.Send(request, response)
    return
}

// DeleteImageTags
// 批量删除镜像版本
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPIMAGETAGISINUSE = "InvalidParameterValue.ContainerGroupImageTagIsInUse"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
func (c *Client) DeleteImageTagsWithContext(ctx context.Context, request *DeleteImageTagsRequest) (response *DeleteImageTagsResponse, err error) {
    if request == nil {
        request = NewDeleteImageTagsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaneRequest() (request *DeleteLaneRequest) {
    request = &DeleteLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteLane")
    
    
    return
}

func NewDeleteLaneResponse() (response *DeleteLaneResponse) {
    response = &DeleteLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLane
// 删除泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLane(request *DeleteLaneRequest) (response *DeleteLaneResponse, err error) {
    if request == nil {
        request = NewDeleteLaneRequest()
    }
    
    response = NewDeleteLaneResponse()
    err = c.Send(request, response)
    return
}

// DeleteLane
// 删除泳道
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLaneWithContext(ctx context.Context, request *DeleteLaneRequest) (response *DeleteLaneResponse, err error) {
    if request == nil {
        request = NewDeleteLaneRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLaneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaneRuleRequest() (request *DeleteLaneRuleRequest) {
    request = &DeleteLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteLaneRule")
    
    
    return
}

func NewDeleteLaneRuleResponse() (response *DeleteLaneRuleResponse) {
    response = &DeleteLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLaneRule
// 删除泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLaneRule(request *DeleteLaneRuleRequest) (response *DeleteLaneRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLaneRuleRequest()
    }
    
    response = NewDeleteLaneRuleResponse()
    err = c.Send(request, response)
    return
}

// DeleteLaneRule
// 删除泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) DeleteLaneRuleWithContext(ctx context.Context, request *DeleteLaneRuleRequest) (response *DeleteLaneRuleResponse, err error) {
    if request == nil {
        request = NewDeleteLaneRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMicroserviceRequest() (request *DeleteMicroserviceRequest) {
    request = &DeleteMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteMicroservice")
    
    
    return
}

func NewDeleteMicroserviceResponse() (response *DeleteMicroserviceResponse) {
    response = &DeleteMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMicroservice
// 删除微服务
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteMicroservice(request *DeleteMicroserviceRequest) (response *DeleteMicroserviceResponse, err error) {
    if request == nil {
        request = NewDeleteMicroserviceRequest()
    }
    
    response = NewDeleteMicroserviceResponse()
    err = c.Send(request, response)
    return
}

// DeleteMicroservice
// 删除微服务
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteMicroserviceWithContext(ctx context.Context, request *DeleteMicroserviceRequest) (response *DeleteMicroserviceResponse, err error) {
    if request == nil {
        request = NewDeleteMicroserviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteNamespace")
    
    
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNamespace
// 删除命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPEXISTS = "FailedOperation.GroupExists"
//  INTERNALERROR_KUBERNETESCALLERROR = "InternalError.KubernetesCallError"
//  RESOURCEINUSE_DEFAULTNAMEPSACECANNOTBEDELETED = "ResourceInUse.DefaultNamepsaceCannotBeDeleted"
//  RESOURCEINUSE_NAMESPACECANNOTDELETE = "ResourceInUse.NamespaceCannotDelete"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

// DeleteNamespace
// 删除命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPEXISTS = "FailedOperation.GroupExists"
//  INTERNALERROR_KUBERNETESCALLERROR = "InternalError.KubernetesCallError"
//  RESOURCEINUSE_DEFAULTNAMEPSACECANNOTBEDELETED = "ResourceInUse.DefaultNamepsaceCannotBeDeleted"
//  RESOURCEINUSE_NAMESPACECANNOTDELETE = "ResourceInUse.NamespaceCannotDelete"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_NAMESPACENOTEXIST = "ResourceNotFound.NamespaceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePathRewritesRequest() (request *DeletePathRewritesRequest) {
    request = &DeletePathRewritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePathRewrites")
    
    
    return
}

func NewDeletePathRewritesResponse() (response *DeletePathRewritesResponse) {
    response = &DeletePathRewritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePathRewrites
// 删除路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeletePathRewrites(request *DeletePathRewritesRequest) (response *DeletePathRewritesResponse, err error) {
    if request == nil {
        request = NewDeletePathRewritesRequest()
    }
    
    response = NewDeletePathRewritesResponse()
    err = c.Send(request, response)
    return
}

// DeletePathRewrites
// 删除路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeletePathRewritesWithContext(ctx context.Context, request *DeletePathRewritesRequest) (response *DeletePathRewritesResponse, err error) {
    if request == nil {
        request = NewDeletePathRewritesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePathRewritesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePkgsRequest() (request *DeletePkgsRequest) {
    request = &DeletePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePkgs")
    
    
    return
}

func NewDeletePkgsResponse() (response *DeletePkgsResponse) {
    response = &DeletePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePkgs
// 从软件仓库批量删除程序包。
//
// 一次最多支持删除1000个包，数量超过1000，返回UpperDeleteLimit错误。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PACKAGEINUSE = "InvalidParameter.PackageInUse"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UPPERDELETELIMIT = "InvalidParameter.UpperDeleteLimit"
func (c *Client) DeletePkgs(request *DeletePkgsRequest) (response *DeletePkgsResponse, err error) {
    if request == nil {
        request = NewDeletePkgsRequest()
    }
    
    response = NewDeletePkgsResponse()
    err = c.Send(request, response)
    return
}

// DeletePkgs
// 从软件仓库批量删除程序包。
//
// 一次最多支持删除1000个包，数量超过1000，返回UpperDeleteLimit错误。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PACKAGEINUSE = "InvalidParameter.PackageInUse"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_UPPERDELETELIMIT = "InvalidParameter.UpperDeleteLimit"
func (c *Client) DeletePkgsWithContext(ctx context.Context, request *DeletePkgsRequest) (response *DeletePkgsResponse, err error) {
    if request == nil {
        request = NewDeletePkgsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublicConfigRequest() (request *DeletePublicConfigRequest) {
    request = &DeletePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeletePublicConfig")
    
    
    return
}

func NewDeletePublicConfigResponse() (response *DeletePublicConfigResponse) {
    response = &DeletePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePublicConfig
// 删除公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
func (c *Client) DeletePublicConfig(request *DeletePublicConfigRequest) (response *DeletePublicConfigResponse, err error) {
    if request == nil {
        request = NewDeletePublicConfigRequest()
    }
    
    response = NewDeletePublicConfigResponse()
    err = c.Send(request, response)
    return
}

// DeletePublicConfig
// 删除公共配置项
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_RELEASEDCONFIGCANNOTBEDELETED = "InvalidParameterValue.ReleasedConfigCanNotBeDeleted"
//  MISSINGPARAMETER_CONFIGIDREQUIRED = "MissingParameter.ConfigIdRequired"
func (c *Client) DeletePublicConfigWithContext(ctx context.Context, request *DeletePublicConfigRequest) (response *DeletePublicConfigResponse, err error) {
    if request == nil {
        request = NewDeletePublicConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeletePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
    request = &DeleteRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteRepository")
    
    
    return
}

func NewDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
    response = &DeleteRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRepository
// 删除仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOSITORYNOTEMPTY = "InvalidParameter.RepositoryNotEmpty"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryRequest()
    }
    
    response = NewDeleteRepositoryResponse()
    err = c.Send(request, response)
    return
}

// DeleteRepository
// 删除仓库
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOSITORYNOTEMPTY = "InvalidParameter.RepositoryNotEmpty"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DeleteRepositoryWithContext(ctx context.Context, request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServerlessGroupRequest() (request *DeleteServerlessGroupRequest) {
    request = &DeleteServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteServerlessGroup")
    
    
    return
}

func NewDeleteServerlessGroupResponse() (response *DeleteServerlessGroupResponse) {
    response = &DeleteServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteServerlessGroup
// 删除Serverless部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteServerlessGroup(request *DeleteServerlessGroupRequest) (response *DeleteServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessGroupRequest()
    }
    
    response = NewDeleteServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

// DeleteServerlessGroup
// 删除Serverless部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeleteServerlessGroupWithContext(ctx context.Context, request *DeleteServerlessGroupRequest) (response *DeleteServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServerlessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteTask")
    
    
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTask
// 删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKDELETEERROR = "FailedOperation.TaskDeleteError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

// DeleteTask
// 删除任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKDELETEERROR = "FailedOperation.TaskDeleteError"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeleteTaskWithContext(ctx context.Context, request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUnitNamespacesRequest() (request *DeleteUnitNamespacesRequest) {
    request = &DeleteUnitNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteUnitNamespaces")
    
    
    return
}

func NewDeleteUnitNamespacesResponse() (response *DeleteUnitNamespacesResponse) {
    response = &DeleteUnitNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUnitNamespaces
// 删除单元化命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitNamespaces(request *DeleteUnitNamespacesRequest) (response *DeleteUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDeleteUnitNamespacesRequest()
    }
    
    response = NewDeleteUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

// DeleteUnitNamespaces
// 删除单元化命名空间
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitNamespacesWithContext(ctx context.Context, request *DeleteUnitNamespacesRequest) (response *DeleteUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDeleteUnitNamespacesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUnitRuleRequest() (request *DeleteUnitRuleRequest) {
    request = &DeleteUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeleteUnitRule")
    
    
    return
}

func NewDeleteUnitRuleResponse() (response *DeleteUnitRuleResponse) {
    response = &DeleteUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUnitRule
// 删除单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitRule(request *DeleteUnitRuleRequest) (response *DeleteUnitRuleResponse, err error) {
    if request == nil {
        request = NewDeleteUnitRuleRequest()
    }
    
    response = NewDeleteUnitRuleResponse()
    err = c.Send(request, response)
    return
}

// DeleteUnitRule
// 删除单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DeleteUnitRuleWithContext(ctx context.Context, request *DeleteUnitRuleRequest) (response *DeleteUnitRuleResponse, err error) {
    if request == nil {
        request = NewDeleteUnitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeployContainerGroupRequest() (request *DeployContainerGroupRequest) {
    request = &DeployContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerGroup")
    
    
    return
}

func NewDeployContainerGroupResponse() (response *DeployContainerGroupResponse) {
    response = &DeployContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeployContainerGroup
// 部署容器应用-更新
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPENVVALUENOTSET = "InvalidParameterValue.ContainergroupEnvValueNotSet"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATETYPEINVALID = "InvalidParameterValue.ContainergroupUpdatetypeInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoReponameNull"
//  INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployContainerGroup(request *DeployContainerGroupRequest) (response *DeployContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeployContainerGroupRequest()
    }
    
    response = NewDeployContainerGroupResponse()
    err = c.Send(request, response)
    return
}

// DeployContainerGroup
// 部署容器应用-更新
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPENVVALUENOTSET = "InvalidParameterValue.ContainergroupEnvValueNotSet"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDCPUINFO = "InvalidParameterValue.ContainergroupInvalidCpuInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINVALIDMEMINFO = "InvalidParameterValue.ContainergroupInvalidMemInfo"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPNODEPORTINVALID = "InvalidParameterValue.ContainergroupNodePortInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPORTSREPEAT = "InvalidParameterValue.ContainergroupPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPPROTOCOLPORTSNULL = "InvalidParameterValue.ContainergroupProtocolPortsNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPREPONAMEINVALID = "InvalidParameterValue.ContainergroupReponameInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPRESOURCEAGENTVALUEINVALID = "InvalidParameterValue.ContainergroupResourceAgentValueInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTSREPEAT = "InvalidParameterValue.ContainergroupTargetPortsRepeat"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPTARGETPORTNULL = "InvalidParameterValue.ContainergroupTargetportNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATEIVLINVALID = "InvalidParameterValue.ContainergroupUpdateivlInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPUPDATETYPEINVALID = "InvalidParameterValue.ContainergroupUpdatetypeInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPYAMLUSERCONTAINERNOTFOUND = "InvalidParameterValue.ContainergroupYamlUserContainerNotFound"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMEINVALID = "InvalidParameterValue.ImagerepoReponameInvalid"
//  INVALIDPARAMETERVALUE_IMAGEREPOREPONAMENULL = "InvalidParameterValue.ImagerepoReponameNull"
//  INVALIDPARAMETERVALUE_IMAGEREPOTAGNAMENULL = "InvalidParameterValue.ImagerepoTagnameNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployContainerGroupWithContext(ctx context.Context, request *DeployContainerGroupRequest) (response *DeployContainerGroupResponse, err error) {
    if request == nil {
        request = NewDeployContainerGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeployContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployGroupRequest() (request *DeployGroupRequest) {
    request = &DeployGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployGroup")
    
    
    return
}

func NewDeployGroupResponse() (response *DeployGroupResponse) {
    response = &DeployGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeployGroup
// 部署虚拟机部署组应用
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERINTERNALERROR = "InternalError.CvmCaeMasterInternalError"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"
//  INVALIDPARAMETERVALUE_GROUPBATCHPARAMETERINVALID = "InvalidParameterValue.GroupBatchParameterInvalid"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployGroup(request *DeployGroupRequest) (response *DeployGroupResponse, err error) {
    if request == nil {
        request = NewDeployGroupRequest()
    }
    
    response = NewDeployGroupResponse()
    err = c.Send(request, response)
    return
}

// DeployGroup
// 部署虚拟机部署组应用
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERINTERNALERROR = "InternalError.CvmCaeMasterInternalError"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"
//  INVALIDPARAMETERVALUE_GROUPBATCHPARAMETERINVALID = "InvalidParameterValue.GroupBatchParameterInvalid"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DeployGroupWithContext(ctx context.Context, request *DeployGroupRequest) (response *DeployGroupResponse, err error) {
    if request == nil {
        request = NewDeployGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeployGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeployServerlessGroupRequest() (request *DeployServerlessGroupRequest) {
    request = &DeployServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DeployServerlessGroup")
    
    
    return
}

func NewDeployServerlessGroupResponse() (response *DeployServerlessGroupResponse) {
    response = &DeployServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeployServerlessGroup
// 部署Serverless应用
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCEINUSE_GROUPINOPERATION = "ResourceInUse.GroupInOperation"
//  RESOURCEINSUFFICIENT_INSTANCEEXCESSLIMIT = "ResourceInsufficient.InstanceExcessLimit"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeployServerlessGroup(request *DeployServerlessGroupRequest) (response *DeployServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeployServerlessGroupRequest()
    }
    
    response = NewDeployServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

// DeployServerlessGroup
// 部署Serverless应用
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTNOTFOUND = "InvalidParameterValue.CvmCaeMasterAgentNotFound"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCEINUSE_GROUPINOPERATION = "ResourceInUse.GroupInOperation"
//  RESOURCEINSUFFICIENT_INSTANCEEXCESSLIMIT = "ResourceInsufficient.InstanceExcessLimit"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DeployServerlessGroupWithContext(ctx context.Context, request *DeployServerlessGroupRequest) (response *DeployServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDeployServerlessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeployServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiDetailRequest() (request *DescribeApiDetailRequest) {
    request = &DescribeApiDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiDetail")
    
    
    return
}

func NewDescribeApiDetailResponse() (response *DescribeApiDetailResponse) {
    response = &DescribeApiDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiDetail
// 查询API详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIMETAPARSEFAILED = "FailedOperation.ApiMetaParseFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiDetail(request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiDetailRequest()
    }
    
    response = NewDescribeApiDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeApiDetail
// 查询API详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIMETAPARSEFAILED = "FailedOperation.ApiMetaParseFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiDetailWithContext(ctx context.Context, request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApiDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupRequest() (request *DescribeApiGroupRequest) {
    request = &DescribeApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroup")
    
    
    return
}

func NewDescribeApiGroupResponse() (response *DescribeApiGroupResponse) {
    response = &DescribeApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiGroup
// 查询API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupRequest()
    }
    
    response = NewDescribeApiGroupResponse()
    err = c.Send(request, response)
    return
}

// DescribeApiGroup
// 查询API分组
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiGroupWithContext(ctx context.Context, request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiGroupsRequest() (request *DescribeApiGroupsRequest) {
    request = &DescribeApiGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroups")
    
    
    return
}

func NewDescribeApiGroupsResponse() (response *DescribeApiGroupsResponse) {
    response = &DescribeApiGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiGroups
// 查询API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (response *DescribeApiGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupsRequest()
    }
    
    response = NewDescribeApiGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeApiGroups
// 查询API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApiGroupsWithContext(ctx context.Context, request *DescribeApiGroupsRequest) (response *DescribeApiGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeApiGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApiGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiRateLimitRulesRequest() (request *DescribeApiRateLimitRulesRequest) {
    request = &DescribeApiRateLimitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiRateLimitRules")
    
    
    return
}

func NewDescribeApiRateLimitRulesResponse() (response *DescribeApiRateLimitRulesResponse) {
    response = &DescribeApiRateLimitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiRateLimitRules
// 查询API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeApiRateLimitRules(request *DescribeApiRateLimitRulesRequest) (response *DescribeApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeApiRateLimitRulesRequest()
    }
    
    response = NewDescribeApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

// DescribeApiRateLimitRules
// 查询API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeApiRateLimitRulesWithContext(ctx context.Context, request *DescribeApiRateLimitRulesRequest) (response *DescribeApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeApiRateLimitRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiUseDetailRequest() (request *DescribeApiUseDetailRequest) {
    request = &DescribeApiUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiUseDetail")
    
    
    return
}

func NewDescribeApiUseDetailResponse() (response *DescribeApiUseDetailResponse) {
    response = &DescribeApiUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiUseDetail
// 查询网关API监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiUseDetail(request *DescribeApiUseDetailRequest) (response *DescribeApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiUseDetailRequest()
    }
    
    response = NewDescribeApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeApiUseDetail
// 查询网关API监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeApiUseDetailWithContext(ctx context.Context, request *DescribeApiUseDetailRequest) (response *DescribeApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApiUseDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiVersionsRequest() (request *DescribeApiVersionsRequest) {
    request = &DescribeApiVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiVersions")
    
    
    return
}

func NewDescribeApiVersionsResponse() (response *DescribeApiVersionsResponse) {
    response = &DescribeApiVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApiVersions
// 查询API 版本
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiVersions(request *DescribeApiVersionsRequest) (response *DescribeApiVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeApiVersionsRequest()
    }
    
    response = NewDescribeApiVersionsResponse()
    err = c.Send(request, response)
    return
}

// DescribeApiVersions
// 查询API 版本
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeApiVersionsWithContext(ctx context.Context, request *DescribeApiVersionsRequest) (response *DescribeApiVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeApiVersionsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApiVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
    request = &DescribeApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplication")
    
    
    return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
    response = &DescribeApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplication
// 获取应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

// DescribeApplication
// 获取应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_APPLICATIONIDNULL = "MissingParameter.ApplicationIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationWithContext(ctx context.Context, request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationAttributeRequest() (request *DescribeApplicationAttributeRequest) {
    request = &DescribeApplicationAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationAttribute")
    
    
    return
}

func NewDescribeApplicationAttributeResponse() (response *DescribeApplicationAttributeResponse) {
    response = &DescribeApplicationAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationAttribute
// 获取应用列表其它字段，如实例数量信息等
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_APPLICATIONPROJECTNOTMATCH = "ResourceNotFound.ApplicationProjectNotMatch"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationAttribute(request *DescribeApplicationAttributeRequest) (response *DescribeApplicationAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationAttributeRequest()
    }
    
    response = NewDescribeApplicationAttributeResponse()
    err = c.Send(request, response)
    return
}

// DescribeApplicationAttribute
// 获取应用列表其它字段，如实例数量信息等
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_APPLICATIONPROJECTNOTMATCH = "ResourceNotFound.ApplicationProjectNotMatch"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationAttributeWithContext(ctx context.Context, request *DescribeApplicationAttributeRequest) (response *DescribeApplicationAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApplicationAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
    request = &DescribeApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplications")
    
    
    return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
    response = &DescribeApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplications
// 获取应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_APPLICATIONPAGELIMITINVALID = "InvalidParameterValue.ApplicationPageLimitInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

// DescribeApplications
// 获取应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_APPLICATIONPAGELIMITINVALID = "InvalidParameterValue.ApplicationPageLimitInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeApplicationsWithContext(ctx context.Context, request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicResourceUsageRequest() (request *DescribeBasicResourceUsageRequest) {
    request = &DescribeBasicResourceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeBasicResourceUsage")
    
    
    return
}

func NewDescribeBasicResourceUsageResponse() (response *DescribeBasicResourceUsageResponse) {
    response = &DescribeBasicResourceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBasicResourceUsage
// TSF基本资源信息概览接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeBasicResourceUsage(request *DescribeBasicResourceUsageRequest) (response *DescribeBasicResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeBasicResourceUsageRequest()
    }
    
    response = NewDescribeBasicResourceUsageResponse()
    err = c.Send(request, response)
    return
}

// DescribeBasicResourceUsage
// TSF基本资源信息概览接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeBasicResourceUsageWithContext(ctx context.Context, request *DescribeBasicResourceUsageRequest) (response *DescribeBasicResourceUsageResponse, err error) {
    if request == nil {
        request = NewDescribeBasicResourceUsageRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeBasicResourceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterInstances")
    
    
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterInstances
// 查询集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERNOTEXISTORPRIVILEGEERROR = "InternalError.ClusterNotExistOrPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeClusterInstances
// 查询集群实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_INSTANCEUPDATEFAILED = "FailedOperation.InstanceUpdateFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_CLUSTERNOTEXISTORPRIVILEGEERROR = "InternalError.ClusterNotExistOrPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeClusterInstancesWithContext(ctx context.Context, request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfig")
    
    
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfig
// 查询配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeConfig
// 查询配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
func (c *Client) DescribeConfigWithContext(ctx context.Context, request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleaseLogsRequest() (request *DescribeConfigReleaseLogsRequest) {
    request = &DescribeConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleaseLogs")
    
    
    return
}

func NewDescribeConfigReleaseLogsResponse() (response *DescribeConfigReleaseLogsResponse) {
    response = &DescribeConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigReleaseLogs
// 查询配置发布历史
//
// 可能返回的错误码:
//  INTERNALERROR_SQLTOOMANYINITEM = "InternalError.SqlTooManyInItem"
func (c *Client) DescribeConfigReleaseLogs(request *DescribeConfigReleaseLogsRequest) (response *DescribeConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleaseLogsRequest()
    }
    
    response = NewDescribeConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

// DescribeConfigReleaseLogs
// 查询配置发布历史
//
// 可能返回的错误码:
//  INTERNALERROR_SQLTOOMANYINITEM = "InternalError.SqlTooManyInItem"
func (c *Client) DescribeConfigReleaseLogsWithContext(ctx context.Context, request *DescribeConfigReleaseLogsRequest) (response *DescribeConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleaseLogsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigReleasesRequest() (request *DescribeConfigReleasesRequest) {
    request = &DescribeConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleases")
    
    
    return
}

func NewDescribeConfigReleasesResponse() (response *DescribeConfigReleasesResponse) {
    response = &DescribeConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigReleases
// 查询配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigReleases(request *DescribeConfigReleasesRequest) (response *DescribeConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleasesRequest()
    }
    
    response = NewDescribeConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

// DescribeConfigReleases
// 查询配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigReleasesWithContext(ctx context.Context, request *DescribeConfigReleasesRequest) (response *DescribeConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribeConfigReleasesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigSummaryRequest() (request *DescribeConfigSummaryRequest) {
    request = &DescribeConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigSummary")
    
    
    return
}

func NewDescribeConfigSummaryResponse() (response *DescribeConfigSummaryResponse) {
    response = &DescribeConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigSummary
// 查询配置汇总列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigSummary(request *DescribeConfigSummaryRequest) (response *DescribeConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeConfigSummaryRequest()
    }
    
    response = NewDescribeConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

// DescribeConfigSummary
// 查询配置汇总列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigSummaryWithContext(ctx context.Context, request *DescribeConfigSummaryRequest) (response *DescribeConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeConfigSummaryRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
    request = &DescribeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigs")
    
    
    return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
    response = &DescribeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigs
// 查询配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_CONFIGQUERYFAILED = "FailedOperation.ConfigQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

// DescribeConfigs
// 查询配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
//  FAILEDOPERATION_CONFIGQUERYFAILED = "FailedOperation.ConfigQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeConfigsWithContext(ctx context.Context, request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerEventsRequest() (request *DescribeContainerEventsRequest) {
    request = &DescribeContainerEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerEvents")
    
    
    return
}

func NewDescribeContainerEventsResponse() (response *DescribeContainerEventsResponse) {
    response = &DescribeContainerEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerEvents
// 获取容器事件列表
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_RESOURCETYPEERROR = "InvalidParameterValue.ResourceTypeError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
func (c *Client) DescribeContainerEvents(request *DescribeContainerEventsRequest) (response *DescribeContainerEventsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerEventsRequest()
    }
    
    response = NewDescribeContainerEventsResponse()
    err = c.Send(request, response)
    return
}

// DescribeContainerEvents
// 获取容器事件列表
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_RESOURCETYPEERROR = "InvalidParameterValue.ResourceTypeError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
func (c *Client) DescribeContainerEventsWithContext(ctx context.Context, request *DescribeContainerEventsRequest) (response *DescribeContainerEventsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerEventsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeContainerEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupDeployInfoRequest() (request *DescribeContainerGroupDeployInfoRequest) {
    request = &DescribeContainerGroupDeployInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDeployInfo")
    
    
    return
}

func NewDescribeContainerGroupDeployInfoResponse() (response *DescribeContainerGroupDeployInfoResponse) {
    response = &DescribeContainerGroupDeployInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerGroupDeployInfo
//  获取部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDeployInfo(request *DescribeContainerGroupDeployInfoRequest) (response *DescribeContainerGroupDeployInfoResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDeployInfoRequest()
    }
    
    response = NewDescribeContainerGroupDeployInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeContainerGroupDeployInfo
//  获取部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDeployInfoWithContext(ctx context.Context, request *DescribeContainerGroupDeployInfoRequest) (response *DescribeContainerGroupDeployInfoResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDeployInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeContainerGroupDeployInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupDetailRequest() (request *DescribeContainerGroupDetailRequest) {
    request = &DescribeContainerGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDetail")
    
    
    return
}

func NewDescribeContainerGroupDetailResponse() (response *DescribeContainerGroupDetailResponse) {
    response = &DescribeContainerGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerGroupDetail
//  容器部署组详情（已废弃，请使用  DescribeContainerGroupDeployInfo）
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDetail(request *DescribeContainerGroupDetailRequest) (response *DescribeContainerGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDetailRequest()
    }
    
    response = NewDescribeContainerGroupDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeContainerGroupDetail
//  容器部署组详情（已废弃，请使用  DescribeContainerGroupDeployInfo）
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupDetailWithContext(ctx context.Context, request *DescribeContainerGroupDetailRequest) (response *DescribeContainerGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeContainerGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerGroupsRequest() (request *DescribeContainerGroupsRequest) {
    request = &DescribeContainerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroups")
    
    
    return
}

func NewDescribeContainerGroupsResponse() (response *DescribeContainerGroupsResponse) {
    response = &DescribeContainerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerGroups
// 容器部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupsRequest()
    }
    
    response = NewDescribeContainerGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeContainerGroups
// 容器部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPSQLFAILED = "InternalError.ContainergroupSqlFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPAPPLICATIONIDNULL = "InvalidParameterValue.ContainergroupApplicationIdNull"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeContainerGroupsWithContext(ctx context.Context, request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeContainerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCreateGatewayApiStatusRequest() (request *DescribeCreateGatewayApiStatusRequest) {
    request = &DescribeCreateGatewayApiStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeCreateGatewayApiStatus")
    
    
    return
}

func NewDescribeCreateGatewayApiStatusResponse() (response *DescribeCreateGatewayApiStatusResponse) {
    response = &DescribeCreateGatewayApiStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCreateGatewayApiStatus
// 查询一键导入API分组任务的状态
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeCreateGatewayApiStatus(request *DescribeCreateGatewayApiStatusRequest) (response *DescribeCreateGatewayApiStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCreateGatewayApiStatusRequest()
    }
    
    response = NewDescribeCreateGatewayApiStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeCreateGatewayApiStatus
// 查询一键导入API分组任务的状态
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
func (c *Client) DescribeCreateGatewayApiStatusWithContext(ctx context.Context, request *DescribeCreateGatewayApiStatusRequest) (response *DescribeCreateGatewayApiStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCreateGatewayApiStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCreateGatewayApiStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadInfoRequest() (request *DescribeDownloadInfoRequest) {
    request = &DescribeDownloadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeDownloadInfo")
    
    
    return
}

func NewDescribeDownloadInfoResponse() (response *DescribeDownloadInfoResponse) {
    response = &DescribeDownloadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDownloadInfo
// TSF上传的程序包存放在腾讯云对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DescribeDownloadInfo(request *DescribeDownloadInfoRequest) (response *DescribeDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadInfoRequest()
    }
    
    response = NewDescribeDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeDownloadInfo
// TSF上传的程序包存放在腾讯云对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DescribeDownloadInfoWithContext(ctx context.Context, request *DescribeDownloadInfoRequest) (response *DescribeDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnabledUnitRuleRequest() (request *DescribeEnabledUnitRuleRequest) {
    request = &DescribeEnabledUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeEnabledUnitRule")
    
    
    return
}

func NewDescribeEnabledUnitRuleResponse() (response *DescribeEnabledUnitRuleResponse) {
    response = &DescribeEnabledUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnabledUnitRule
// 查询生效的单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeEnabledUnitRule(request *DescribeEnabledUnitRuleRequest) (response *DescribeEnabledUnitRuleResponse, err error) {
    if request == nil {
        request = NewDescribeEnabledUnitRuleRequest()
    }
    
    response = NewDescribeEnabledUnitRuleResponse()
    err = c.Send(request, response)
    return
}

// DescribeEnabledUnitRule
// 查询生效的单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeEnabledUnitRuleWithContext(ctx context.Context, request *DescribeEnabledUnitRuleRequest) (response *DescribeEnabledUnitRuleResponse, err error) {
    if request == nil {
        request = NewDescribeEnabledUnitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEnabledUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileConfigsRequest() (request *DescribeFileConfigsRequest) {
    request = &DescribeFileConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigs")
    
    
    return
}

func NewDescribeFileConfigsResponse() (response *DescribeFileConfigsResponse) {
    response = &DescribeFileConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileConfigs
// 查询文件配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
func (c *Client) DescribeFileConfigs(request *DescribeFileConfigsRequest) (response *DescribeFileConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigsRequest()
    }
    
    response = NewDescribeFileConfigsResponse()
    err = c.Send(request, response)
    return
}

// DescribeFileConfigs
// 查询文件配置项列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGAPPLICATIONQUERYFAILED = "FailedOperation.ConfigApplicationQueryFailed"
func (c *Client) DescribeFileConfigsWithContext(ctx context.Context, request *DescribeFileConfigsRequest) (response *DescribeFileConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeFileConfigsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFileConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowLastBatchStateRequest() (request *DescribeFlowLastBatchStateRequest) {
    request = &DescribeFlowLastBatchStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowLastBatchState")
    
    
    return
}

func NewDescribeFlowLastBatchStateResponse() (response *DescribeFlowLastBatchStateResponse) {
    response = &DescribeFlowLastBatchStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowLastBatchState
// 查询工作流最新一个批次的状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeFlowLastBatchState(request *DescribeFlowLastBatchStateRequest) (response *DescribeFlowLastBatchStateResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLastBatchStateRequest()
    }
    
    response = NewDescribeFlowLastBatchStateResponse()
    err = c.Send(request, response)
    return
}

// DescribeFlowLastBatchState
// 查询工作流最新一个批次的状态信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeFlowLastBatchStateWithContext(ctx context.Context, request *DescribeFlowLastBatchStateRequest) (response *DescribeFlowLastBatchStateResponse, err error) {
    if request == nil {
        request = NewDescribeFlowLastBatchStateRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFlowLastBatchStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayAllGroupApisRequest() (request *DescribeGatewayAllGroupApisRequest) {
    request = &DescribeGatewayAllGroupApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayAllGroupApis")
    
    
    return
}

func NewDescribeGatewayAllGroupApisResponse() (response *DescribeGatewayAllGroupApisResponse) {
    response = &DescribeGatewayAllGroupApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayAllGroupApis
// 查询网关所有分组下Api列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEPLOYGROUPNOTEXISTS = "InvalidParameterValue.DeployGroupNotExists"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGatewayAllGroupApis(request *DescribeGatewayAllGroupApisRequest) (response *DescribeGatewayAllGroupApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAllGroupApisRequest()
    }
    
    response = NewDescribeGatewayAllGroupApisResponse()
    err = c.Send(request, response)
    return
}

// DescribeGatewayAllGroupApis
// 查询网关所有分组下Api列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DEPLOYGROUPNOTEXISTS = "InvalidParameterValue.DeployGroupNotExists"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGatewayAllGroupApisWithContext(ctx context.Context, request *DescribeGatewayAllGroupApisRequest) (response *DescribeGatewayAllGroupApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayAllGroupApisRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGatewayAllGroupApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayApisRequest() (request *DescribeGatewayApisRequest) {
    request = &DescribeGatewayApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayApis")
    
    
    return
}

func NewDescribeGatewayApisResponse() (response *DescribeGatewayApisResponse) {
    response = &DescribeGatewayApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayApis
// 查询API分组下的Api列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribeGatewayApis(request *DescribeGatewayApisRequest) (response *DescribeGatewayApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayApisRequest()
    }
    
    response = NewDescribeGatewayApisResponse()
    err = c.Send(request, response)
    return
}

// DescribeGatewayApis
// 查询API分组下的Api列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GATEWAYDBERROR = "InternalError.GatewayDbError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribeGatewayApisWithContext(ctx context.Context, request *DescribeGatewayApisRequest) (response *DescribeGatewayApisResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayApisRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGatewayApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatewayMonitorOverviewRequest() (request *DescribeGatewayMonitorOverviewRequest) {
    request = &DescribeGatewayMonitorOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayMonitorOverview")
    
    
    return
}

func NewDescribeGatewayMonitorOverviewResponse() (response *DescribeGatewayMonitorOverviewResponse) {
    response = &DescribeGatewayMonitorOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGatewayMonitorOverview
// 查询网关监控概览
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGatewayMonitorOverview(request *DescribeGatewayMonitorOverviewRequest) (response *DescribeGatewayMonitorOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayMonitorOverviewRequest()
    }
    
    response = NewDescribeGatewayMonitorOverviewResponse()
    err = c.Send(request, response)
    return
}

// DescribeGatewayMonitorOverview
// 查询网关监控概览
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGatewayMonitorOverviewWithContext(ctx context.Context, request *DescribeGatewayMonitorOverviewRequest) (response *DescribeGatewayMonitorOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeGatewayMonitorOverviewRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGatewayMonitorOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroup")
    
    
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroup
// 查询虚拟机部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroup
// 查询虚拟机部署组详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupWithContext(ctx context.Context, request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupAttributeRequest() (request *DescribeGroupAttributeRequest) {
    request = &DescribeGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupAttribute")
    
    
    return
}

func NewDescribeGroupAttributeResponse() (response *DescribeGroupAttributeResponse) {
    response = &DescribeGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupAttribute
// 获取部署组其他属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupAttribute(request *DescribeGroupAttributeRequest) (response *DescribeGroupAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAttributeRequest()
    }
    
    response = NewDescribeGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupAttribute
// 获取部署组其他属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_GROUPCOMMONERROR = "InternalError.GroupCommonError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_REPOPACKAGEPARAMERROR = "InvalidParameter.RepoPackageParamError"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeGroupAttributeWithContext(ctx context.Context, request *DescribeGroupAttributeRequest) (response *DescribeGroupAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeGroupAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupBindedGatewaysRequest() (request *DescribeGroupBindedGatewaysRequest) {
    request = &DescribeGroupBindedGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBindedGateways")
    
    
    return
}

func NewDescribeGroupBindedGatewaysResponse() (response *DescribeGroupBindedGatewaysResponse) {
    response = &DescribeGroupBindedGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupBindedGateways
// 查询某个API分组已绑定的网关部署组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupBindedGateways(request *DescribeGroupBindedGatewaysRequest) (response *DescribeGroupBindedGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupBindedGatewaysRequest()
    }
    
    response = NewDescribeGroupBindedGatewaysResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupBindedGateways
// 查询某个API分组已绑定的网关部署组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupBindedGatewaysWithContext(ctx context.Context, request *DescribeGroupBindedGatewaysRequest) (response *DescribeGroupBindedGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupBindedGatewaysRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupBindedGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupGatewaysRequest() (request *DescribeGroupGatewaysRequest) {
    request = &DescribeGroupGatewaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupGateways")
    
    
    return
}

func NewDescribeGroupGatewaysResponse() (response *DescribeGroupGatewaysResponse) {
    response = &DescribeGroupGatewaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupGateways
// 查询某个网关绑定的API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeGroupGateways(request *DescribeGroupGatewaysRequest) (response *DescribeGroupGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupGatewaysRequest()
    }
    
    response = NewDescribeGroupGatewaysResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupGateways
// 查询某个网关绑定的API 分组信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeGroupGatewaysWithContext(ctx context.Context, request *DescribeGroupGatewaysRequest) (response *DescribeGroupGatewaysResponse, err error) {
    if request == nil {
        request = NewDescribeGroupGatewaysRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupGatewaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupInstancesRequest() (request *DescribeGroupInstancesRequest) {
    request = &DescribeGroupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupInstances")
    
    
    return
}

func NewDescribeGroupInstancesResponse() (response *DescribeGroupInstancesResponse) {
    response = &DescribeGroupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupInstances
// 查询虚拟机部署组云主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupInstances(request *DescribeGroupInstancesRequest) (response *DescribeGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInstancesRequest()
    }
    
    response = NewDescribeGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupInstances
// 查询虚拟机部署组云主机列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_INSTANCECOMMONERROR = "InternalError.InstanceCommonError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_GROUPIDNULL = "MissingParameter.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupInstancesWithContext(ctx context.Context, request *DescribeGroupInstancesRequest) (response *DescribeGroupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupReleaseRequest() (request *DescribeGroupReleaseRequest) {
    request = &DescribeGroupReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupRelease")
    
    
    return
}

func NewDescribeGroupReleaseResponse() (response *DescribeGroupReleaseResponse) {
    response = &DescribeGroupReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupRelease
// 查询部署组相关的发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeGroupRelease(request *DescribeGroupReleaseRequest) (response *DescribeGroupReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeGroupReleaseRequest()
    }
    
    response = NewDescribeGroupReleaseResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupRelease
// 查询部署组相关的发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) DescribeGroupReleaseWithContext(ctx context.Context, request *DescribeGroupReleaseRequest) (response *DescribeGroupReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeGroupReleaseRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupUseDetailRequest() (request *DescribeGroupUseDetailRequest) {
    request = &DescribeGroupUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupUseDetail")
    
    
    return
}

func NewDescribeGroupUseDetailResponse() (response *DescribeGroupUseDetailResponse) {
    response = &DescribeGroupUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupUseDetail
// 查询网关分组监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupUseDetail(request *DescribeGroupUseDetailRequest) (response *DescribeGroupUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGroupUseDetailRequest()
    }
    
    response = NewDescribeGroupUseDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupUseDetail
// 查询网关分组监控明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeGroupUseDetailWithContext(ctx context.Context, request *DescribeGroupUseDetailRequest) (response *DescribeGroupUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGroupUseDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
    request = &DescribeGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroups")
    
    
    return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
    response = &DescribeGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroups
// 获取虚拟机部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroups
// 获取虚拟机部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupsWithContext(ctx context.Context, request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsWithPluginRequest() (request *DescribeGroupsWithPluginRequest) {
    request = &DescribeGroupsWithPluginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupsWithPlugin")
    
    
    return
}

func NewDescribeGroupsWithPluginResponse() (response *DescribeGroupsWithPluginResponse) {
    response = &DescribeGroupsWithPluginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupsWithPlugin
// 查询某个插件下绑定或未绑定的API分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupsWithPlugin(request *DescribeGroupsWithPluginRequest) (response *DescribeGroupsWithPluginResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsWithPluginRequest()
    }
    
    response = NewDescribeGroupsWithPluginResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroupsWithPlugin
// 查询某个插件下绑定或未绑定的API分组
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_GROUPPAGELIMITINVALID = "InvalidParameterValue.GroupPageLimitInvalid"
//  INVALIDPARAMETERVALUE_GROUPSTATUSINVALID = "InvalidParameterValue.GroupStatusInvalid"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeGroupsWithPluginWithContext(ctx context.Context, request *DescribeGroupsWithPluginRequest) (response *DescribeGroupsWithPluginResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsWithPluginRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupsWithPluginResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRepositoryRequest() (request *DescribeImageRepositoryRequest) {
    request = &DescribeImageRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageRepository")
    
    
    return
}

func NewDescribeImageRepositoryResponse() (response *DescribeImageRepositoryResponse) {
    response = &DescribeImageRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageRepository
// 镜像仓库列表 
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageRepository(request *DescribeImageRepositoryRequest) (response *DescribeImageRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeImageRepositoryRequest()
    }
    
    response = NewDescribeImageRepositoryResponse()
    err = c.Send(request, response)
    return
}

// DescribeImageRepository
// 镜像仓库列表 
//
// 可能返回的错误码:
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageRepositoryWithContext(ctx context.Context, request *DescribeImageRepositoryRequest) (response *DescribeImageRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeImageRepositoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImageRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageTagsRequest() (request *DescribeImageTagsRequest) {
    request = &DescribeImageTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageTags")
    
    
    return
}

func NewDescribeImageTagsResponse() (response *DescribeImageTagsResponse) {
    response = &DescribeImageTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageTags
// 镜像版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageTags(request *DescribeImageTagsRequest) (response *DescribeImageTagsResponse, err error) {
    if request == nil {
        request = NewDescribeImageTagsRequest()
    }
    
    response = NewDescribeImageTagsResponse()
    err = c.Send(request, response)
    return
}

// DescribeImageTags
// 镜像版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  INTERNALERROR_CAMROLEREQUESTERROR = "InternalError.CamRoleRequestError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
//  RESOURCENOTFOUND_GROUPAPPLICATIONNOTEXIST = "ResourceNotFound.GroupApplicationNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMGENERALERROR = "UnauthorizedOperation.CamGeneralError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeImageTagsWithContext(ctx context.Context, request *DescribeImageTagsRequest) (response *DescribeImageTagsResponse, err error) {
    if request == nil {
        request = NewDescribeImageTagsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImageTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLaneRulesRequest() (request *DescribeLaneRulesRequest) {
    request = &DescribeLaneRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeLaneRules")
    
    
    return
}

func NewDescribeLaneRulesResponse() (response *DescribeLaneRulesResponse) {
    response = &DescribeLaneRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLaneRules
// 查询泳道规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLaneRules(request *DescribeLaneRulesRequest) (response *DescribeLaneRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLaneRulesRequest()
    }
    
    response = NewDescribeLaneRulesResponse()
    err = c.Send(request, response)
    return
}

// DescribeLaneRules
// 查询泳道规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_LICENSESERVERNOTFOUND = "ResourceNotFound.LicenseServerNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLaneRulesWithContext(ctx context.Context, request *DescribeLaneRulesRequest) (response *DescribeLaneRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLaneRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLaneRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLanesRequest() (request *DescribeLanesRequest) {
    request = &DescribeLanesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeLanes")
    
    
    return
}

func NewDescribeLanesResponse() (response *DescribeLanesResponse) {
    response = &DescribeLanesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLanes
// 查询泳道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLanes(request *DescribeLanesRequest) (response *DescribeLanesResponse, err error) {
    if request == nil {
        request = NewDescribeLanesRequest()
    }
    
    response = NewDescribeLanesResponse()
    err = c.Send(request, response)
    return
}

// DescribeLanes
// 查询泳道列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeLanesWithContext(ctx context.Context, request *DescribeLanesRequest) (response *DescribeLanesResponse, err error) {
    if request == nil {
        request = NewDescribeLanesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLanesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroserviceRequest() (request *DescribeMicroserviceRequest) {
    request = &DescribeMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservice")
    
    
    return
}

func NewDescribeMicroserviceResponse() (response *DescribeMicroserviceResponse) {
    response = &DescribeMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMicroservice
// 查询微服务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservice(request *DescribeMicroserviceRequest) (response *DescribeMicroserviceResponse, err error) {
    if request == nil {
        request = NewDescribeMicroserviceRequest()
    }
    
    response = NewDescribeMicroserviceResponse()
    err = c.Send(request, response)
    return
}

// DescribeMicroservice
// 查询微服务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CLOUDAPIPROXYERROR = "InternalError.CloudApiProxyError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_SERVICEIDREQUIRED = "MissingParameter.ServiceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroserviceWithContext(ctx context.Context, request *DescribeMicroserviceRequest) (response *DescribeMicroserviceResponse, err error) {
    if request == nil {
        request = NewDescribeMicroserviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMicroservicesRequest() (request *DescribeMicroservicesRequest) {
    request = &DescribeMicroservicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservices")
    
    
    return
}

func NewDescribeMicroservicesResponse() (response *DescribeMicroservicesResponse) {
    response = &DescribeMicroservicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMicroservices
// 获取微服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_SERVICENOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ServiceNotExistsOrPermissionDenied"
//  MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservices(request *DescribeMicroservicesRequest) (response *DescribeMicroservicesResponse, err error) {
    if request == nil {
        request = NewDescribeMicroservicesRequest()
    }
    
    response = NewDescribeMicroservicesResponse()
    err = c.Send(request, response)
    return
}

// DescribeMicroservices
// 获取微服务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_SERVICENOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ServiceNotExistsOrPermissionDenied"
//  MISSINGPARAMETER_NAMESPACEIDREQUIRED = "MissingParameter.NamespaceIdRequired"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeMicroservicesWithContext(ctx context.Context, request *DescribeMicroservicesRequest) (response *DescribeMicroservicesResponse, err error) {
    if request == nil {
        request = NewDescribeMicroservicesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMicroservicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMsApiListRequest() (request *DescribeMsApiListRequest) {
    request = &DescribeMsApiListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsApiList")
    
    
    return
}

func NewDescribeMsApiListResponse() (response *DescribeMsApiListResponse) {
    response = &DescribeMsApiListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMsApiList
// 查询服务API列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeMsApiList(request *DescribeMsApiListRequest) (response *DescribeMsApiListResponse, err error) {
    if request == nil {
        request = NewDescribeMsApiListRequest()
    }
    
    response = NewDescribeMsApiListResponse()
    err = c.Send(request, response)
    return
}

// DescribeMsApiList
// 查询服务API列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEINSERTFAILED = "FailedOperation.ServiceInsertFailed"
//  FAILEDOPERATION_SERVICEQUERYFAILED = "FailedOperation.ServiceQueryFailed"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_REQUIREDPARAMETERMISSING = "MissingParameter.RequiredParameterMissing"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) DescribeMsApiListWithContext(ctx context.Context, request *DescribeMsApiListRequest) (response *DescribeMsApiListResponse, err error) {
    if request == nil {
        request = NewDescribeMsApiListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMsApiListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePathRewriteRequest() (request *DescribePathRewriteRequest) {
    request = &DescribePathRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePathRewrite")
    
    
    return
}

func NewDescribePathRewriteResponse() (response *DescribePathRewriteResponse) {
    response = &DescribePathRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePathRewrite
// 查询路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribePathRewrite(request *DescribePathRewriteRequest) (response *DescribePathRewriteResponse, err error) {
    if request == nil {
        request = NewDescribePathRewriteRequest()
    }
    
    response = NewDescribePathRewriteResponse()
    err = c.Send(request, response)
    return
}

// DescribePathRewrite
// 查询路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribePathRewriteWithContext(ctx context.Context, request *DescribePathRewriteRequest) (response *DescribePathRewriteResponse, err error) {
    if request == nil {
        request = NewDescribePathRewriteRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePathRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePathRewritesRequest() (request *DescribePathRewritesRequest) {
    request = &DescribePathRewritesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePathRewrites")
    
    
    return
}

func NewDescribePathRewritesResponse() (response *DescribePathRewritesResponse) {
    response = &DescribePathRewritesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePathRewrites
// 查询路径重写列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePathRewrites(request *DescribePathRewritesRequest) (response *DescribePathRewritesResponse, err error) {
    if request == nil {
        request = NewDescribePathRewritesRequest()
    }
    
    response = NewDescribePathRewritesResponse()
    err = c.Send(request, response)
    return
}

// DescribePathRewrites
// 查询路径重写列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePathRewritesWithContext(ctx context.Context, request *DescribePathRewritesRequest) (response *DescribePathRewritesResponse, err error) {
    if request == nil {
        request = NewDescribePathRewritesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePathRewritesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePkgsRequest() (request *DescribePkgsRequest) {
    request = &DescribePkgsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePkgs")
    
    
    return
}

func NewDescribePkgsResponse() (response *DescribePkgsResponse) {
    response = &DescribePkgsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePkgs
// 无
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribePkgs(request *DescribePkgsRequest) (response *DescribePkgsResponse, err error) {
    if request == nil {
        request = NewDescribePkgsRequest()
    }
    
    response = NewDescribePkgsResponse()
    err = c.Send(request, response)
    return
}

// DescribePkgs
// 无
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribePkgsWithContext(ctx context.Context, request *DescribePkgsRequest) (response *DescribePkgsResponse, err error) {
    if request == nil {
        request = NewDescribePkgsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePkgsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePluginInstancesRequest() (request *DescribePluginInstancesRequest) {
    request = &DescribePluginInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePluginInstances")
    
    
    return
}

func NewDescribePluginInstancesResponse() (response *DescribePluginInstancesResponse) {
    response = &DescribePluginInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePluginInstances
// 分页查询网关分组/API绑定（或未绑定）的插件列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribePluginInstances(request *DescribePluginInstancesRequest) (response *DescribePluginInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePluginInstancesRequest()
    }
    
    response = NewDescribePluginInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribePluginInstances
// 分页查询网关分组/API绑定（或未绑定）的插件列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribePluginInstancesWithContext(ctx context.Context, request *DescribePluginInstancesRequest) (response *DescribePluginInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePluginInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePluginInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePodInstancesRequest() (request *DescribePodInstancesRequest) {
    request = &DescribePodInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePodInstances")
    
    
    return
}

func NewDescribePodInstancesResponse() (response *DescribePodInstancesResponse) {
    response = &DescribePodInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePodInstances
// 获取部署组实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePodInstances(request *DescribePodInstancesRequest) (response *DescribePodInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePodInstancesRequest()
    }
    
    response = NewDescribePodInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribePodInstances
// 获取部署组实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_INSTANCEQUERYFAILED = "FailedOperation.InstanceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNAMESPACECLUSTERNOTFOUND = "ResourceNotFound.ContainergroupGroupNamespaceClusterNotFound"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePodInstancesWithContext(ctx context.Context, request *DescribePodInstancesRequest) (response *DescribePodInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePodInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePodInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigRequest() (request *DescribePublicConfigRequest) {
    request = &DescribePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfig")
    
    
    return
}

func NewDescribePublicConfigResponse() (response *DescribePublicConfigResponse) {
    response = &DescribePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublicConfig
// 查询公共配置（单条）
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfig(request *DescribePublicConfigRequest) (response *DescribePublicConfigResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigRequest()
    }
    
    response = NewDescribePublicConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribePublicConfig
// 查询公共配置（单条）
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigWithContext(ctx context.Context, request *DescribePublicConfigRequest) (response *DescribePublicConfigResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleaseLogsRequest() (request *DescribePublicConfigReleaseLogsRequest) {
    request = &DescribePublicConfigReleaseLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleaseLogs")
    
    
    return
}

func NewDescribePublicConfigReleaseLogsResponse() (response *DescribePublicConfigReleaseLogsResponse) {
    response = &DescribePublicConfigReleaseLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublicConfigReleaseLogs
// 查询公共配置发布历史
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleaseLogs(request *DescribePublicConfigReleaseLogsRequest) (response *DescribePublicConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleaseLogsRequest()
    }
    
    response = NewDescribePublicConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

// DescribePublicConfigReleaseLogs
// 查询公共配置发布历史
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleaseLogsWithContext(ctx context.Context, request *DescribePublicConfigReleaseLogsRequest) (response *DescribePublicConfigReleaseLogsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleaseLogsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePublicConfigReleaseLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigReleasesRequest() (request *DescribePublicConfigReleasesRequest) {
    request = &DescribePublicConfigReleasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleases")
    
    
    return
}

func NewDescribePublicConfigReleasesResponse() (response *DescribePublicConfigReleasesResponse) {
    response = &DescribePublicConfigReleasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublicConfigReleases
// 查询公共配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleases(request *DescribePublicConfigReleasesRequest) (response *DescribePublicConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleasesRequest()
    }
    
    response = NewDescribePublicConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

// DescribePublicConfigReleases
// 查询公共配置发布信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  FAILEDOPERATION_CONFIGRELEASEQUERYFAILED = "FailedOperation.ConfigReleaseQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
func (c *Client) DescribePublicConfigReleasesWithContext(ctx context.Context, request *DescribePublicConfigReleasesRequest) (response *DescribePublicConfigReleasesResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigReleasesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePublicConfigReleasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigSummaryRequest() (request *DescribePublicConfigSummaryRequest) {
    request = &DescribePublicConfigSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigSummary")
    
    
    return
}

func NewDescribePublicConfigSummaryResponse() (response *DescribePublicConfigSummaryResponse) {
    response = &DescribePublicConfigSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublicConfigSummary
// 查询公共配置汇总列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigSummary(request *DescribePublicConfigSummaryRequest) (response *DescribePublicConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigSummaryRequest()
    }
    
    response = NewDescribePublicConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

// DescribePublicConfigSummary
// 查询公共配置汇总列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigSummaryWithContext(ctx context.Context, request *DescribePublicConfigSummaryRequest) (response *DescribePublicConfigSummaryResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigSummaryRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePublicConfigSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicConfigsRequest() (request *DescribePublicConfigsRequest) {
    request = &DescribePublicConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigs")
    
    
    return
}

func NewDescribePublicConfigsResponse() (response *DescribePublicConfigsResponse) {
    response = &DescribePublicConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublicConfigs
// 查询公共配置项列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigs(request *DescribePublicConfigsRequest) (response *DescribePublicConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigsRequest()
    }
    
    response = NewDescribePublicConfigsResponse()
    err = c.Send(request, response)
    return
}

// DescribePublicConfigs
// 查询公共配置项列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribePublicConfigsWithContext(ctx context.Context, request *DescribePublicConfigsRequest) (response *DescribePublicConfigsResponse, err error) {
    if request == nil {
        request = NewDescribePublicConfigsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePublicConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleasedConfigRequest() (request *DescribeReleasedConfigRequest) {
    request = &DescribeReleasedConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasedConfig")
    
    
    return
}

func NewDescribeReleasedConfigResponse() (response *DescribeReleasedConfigResponse) {
    response = &DescribeReleasedConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReleasedConfig
// 查询group发布的配置
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeReleasedConfig(request *DescribeReleasedConfigRequest) (response *DescribeReleasedConfigResponse, err error) {
    if request == nil {
        request = NewDescribeReleasedConfigRequest()
    }
    
    response = NewDescribeReleasedConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeReleasedConfig
// 查询group发布的配置
//
// 可能返回的错误码:
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeReleasedConfigWithContext(ctx context.Context, request *DescribeReleasedConfigRequest) (response *DescribeReleasedConfigResponse, err error) {
    if request == nil {
        request = NewDescribeReleasedConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeReleasedConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoriesRequest() (request *DescribeRepositoriesRequest) {
    request = &DescribeRepositoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepositories")
    
    
    return
}

func NewDescribeRepositoriesResponse() (response *DescribeRepositoriesResponse) {
    response = &DescribeRepositoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRepositories
// 查询仓库列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeRepositories(request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoriesRequest()
    }
    
    response = NewDescribeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

// DescribeRepositories
// 查询仓库列表
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) DescribeRepositoriesWithContext(ctx context.Context, request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoriesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryRequest() (request *DescribeRepositoryRequest) {
    request = &DescribeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepository")
    
    
    return
}

func NewDescribeRepositoryResponse() (response *DescribeRepositoryResponse) {
    response = &DescribeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRepository
// 查询仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DescribeRepository(request *DescribeRepositoryRequest) (response *DescribeRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryRequest()
    }
    
    response = NewDescribeRepositoryResponse()
    err = c.Send(request, response)
    return
}

// DescribeRepository
// 查询仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DescribeRepositoryWithContext(ctx context.Context, request *DescribeRepositoryRequest) (response *DescribeRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessGroupRequest() (request *DescribeServerlessGroupRequest) {
    request = &DescribeServerlessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroup")
    
    
    return
}

func NewDescribeServerlessGroupResponse() (response *DescribeServerlessGroupResponse) {
    response = &DescribeServerlessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServerlessGroup
// 查询Serverless部署组明细
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFASRESOURCESERVERERROR = "FailedOperation.TsfAsResourceServerError"
//  INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeServerlessGroup(request *DescribeServerlessGroupRequest) (response *DescribeServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupRequest()
    }
    
    response = NewDescribeServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

// DescribeServerlessGroup
// 查询Serverless部署组明细
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFASRESOURCESERVERERROR = "FailedOperation.TsfAsResourceServerError"
//  INTERNALERROR_REMOTESERVICECALLERROR = "InternalError.RemoteServiceCallError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GROUPIDNULL = "InvalidParameterValue.GroupIdNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeServerlessGroupWithContext(ctx context.Context, request *DescribeServerlessGroupRequest) (response *DescribeServerlessGroupResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeServerlessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessGroupsRequest() (request *DescribeServerlessGroupsRequest) {
    request = &DescribeServerlessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroups")
    
    
    return
}

func NewDescribeServerlessGroupsResponse() (response *DescribeServerlessGroupsResponse) {
    response = &DescribeServerlessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServerlessGroups
// 查询Serverless部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNULL = "InvalidParameterValue.ApplicationIdNull"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeServerlessGroups(request *DescribeServerlessGroupsRequest) (response *DescribeServerlessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupsRequest()
    }
    
    response = NewDescribeServerlessGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeServerlessGroups
// 查询Serverless部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNULL = "InvalidParameterValue.ApplicationIdNull"
//  RESOURCENOTFOUND_GROUPNAMESPACENOTEXIST = "ResourceNotFound.GroupNamespaceNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) DescribeServerlessGroupsWithContext(ctx context.Context, request *DescribeServerlessGroupsRequest) (response *DescribeServerlessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeServerlessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleApplicationsRequest() (request *DescribeSimpleApplicationsRequest) {
    request = &DescribeSimpleApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleApplications")
    
    
    return
}

func NewDescribeSimpleApplicationsResponse() (response *DescribeSimpleApplicationsResponse) {
    response = &DescribeSimpleApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSimpleApplications
// 查询简单应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleApplications(request *DescribeSimpleApplicationsRequest) (response *DescribeSimpleApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleApplicationsRequest()
    }
    
    response = NewDescribeSimpleApplicationsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSimpleApplications
// 查询简单应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONQUERYFAILED = "FailedOperation.ApplicationQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleApplicationsWithContext(ctx context.Context, request *DescribeSimpleApplicationsRequest) (response *DescribeSimpleApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleApplicationsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSimpleApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleClustersRequest() (request *DescribeSimpleClustersRequest) {
    request = &DescribeSimpleClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleClusters")
    
    
    return
}

func NewDescribeSimpleClustersResponse() (response *DescribeSimpleClustersResponse) {
    response = &DescribeSimpleClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSimpleClusters
// 查询简单集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleClusters(request *DescribeSimpleClustersRequest) (response *DescribeSimpleClustersResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleClustersRequest()
    }
    
    response = NewDescribeSimpleClustersResponse()
    err = c.Send(request, response)
    return
}

// DescribeSimpleClusters
// 查询简单集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleClustersWithContext(ctx context.Context, request *DescribeSimpleClustersRequest) (response *DescribeSimpleClustersResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleClustersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSimpleClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleGroupsRequest() (request *DescribeSimpleGroupsRequest) {
    request = &DescribeSimpleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleGroups")
    
    
    return
}

func NewDescribeSimpleGroupsResponse() (response *DescribeSimpleGroupsResponse) {
    response = &DescribeSimpleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSimpleGroups
// 查询简单部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleGroups(request *DescribeSimpleGroupsRequest) (response *DescribeSimpleGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleGroupsRequest()
    }
    
    response = NewDescribeSimpleGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeSimpleGroups
// 查询简单部署组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPQUERYFAILD = "FailedOperation.GroupQueryFaild"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleGroupsWithContext(ctx context.Context, request *DescribeSimpleGroupsRequest) (response *DescribeSimpleGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSimpleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleNamespacesRequest() (request *DescribeSimpleNamespacesRequest) {
    request = &DescribeSimpleNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleNamespaces")
    
    
    return
}

func NewDescribeSimpleNamespacesResponse() (response *DescribeSimpleNamespacesResponse) {
    response = &DescribeSimpleNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSimpleNamespaces
// 查询简单命名空间列表 
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleNamespaces(request *DescribeSimpleNamespacesRequest) (response *DescribeSimpleNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleNamespacesRequest()
    }
    
    response = NewDescribeSimpleNamespacesResponse()
    err = c.Send(request, response)
    return
}

// DescribeSimpleNamespaces
// 查询简单命名空间列表 
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_NAMESPACEQUERYFAILED = "FailedOperation.NamespaceQueryFailed"
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOLICENSE = "UnauthorizedOperation.NoLicense"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeSimpleNamespacesWithContext(ctx context.Context, request *DescribeSimpleNamespacesRequest) (response *DescribeSimpleNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleNamespacesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSimpleNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
    request = &DescribeTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskDetail")
    
    
    return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
    response = &DescribeTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskDetail
// 查询任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeTaskDetail
// 查询任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskDetailWithContext(ctx context.Context, request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTaskDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLastStatusRequest() (request *DescribeTaskLastStatusRequest) {
    request = &DescribeTaskLastStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskLastStatus")
    
    
    return
}

func NewDescribeTaskLastStatusResponse() (response *DescribeTaskLastStatusResponse) {
    response = &DescribeTaskLastStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskLastStatus
// 查询任务最近一次执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskLastStatus(request *DescribeTaskLastStatusRequest) (response *DescribeTaskLastStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLastStatusRequest()
    }
    
    response = NewDescribeTaskLastStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeTaskLastStatus
// 查询任务最近一次执行状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskLastStatusWithContext(ctx context.Context, request *DescribeTaskLastStatusRequest) (response *DescribeTaskLastStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLastStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskLastStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRecordsRequest() (request *DescribeTaskRecordsRequest) {
    request = &DescribeTaskRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskRecords")
    
    
    return
}

func NewDescribeTaskRecordsResponse() (response *DescribeTaskRecordsResponse) {
    response = &DescribeTaskRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskRecords
// 翻页查询任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskRecords(request *DescribeTaskRecordsRequest) (response *DescribeTaskRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRecordsRequest()
    }
    
    response = NewDescribeTaskRecordsResponse()
    err = c.Send(request, response)
    return
}

// DescribeTaskRecords
// 翻页查询任务列表
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DescribeTaskRecordsWithContext(ctx context.Context, request *DescribeTaskRecordsRequest) (response *DescribeTaskRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRecordsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTaskRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitApiUseDetailRequest() (request *DescribeUnitApiUseDetailRequest) {
    request = &DescribeUnitApiUseDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitApiUseDetail")
    
    
    return
}

func NewDescribeUnitApiUseDetailResponse() (response *DescribeUnitApiUseDetailResponse) {
    response = &DescribeUnitApiUseDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnitApiUseDetail
// 查询网关API监控明细数据（仅单元化网关），非单元化网关使用DescribeApiUseDetail
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitApiUseDetail(request *DescribeUnitApiUseDetailRequest) (response *DescribeUnitApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUnitApiUseDetailRequest()
    }
    
    response = NewDescribeUnitApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeUnitApiUseDetail
// 查询网关API监控明细数据（仅单元化网关），非单元化网关使用DescribeApiUseDetail
//
// 可能返回的错误码:
//  FAILEDOPERATION_GATEWAYREMOTECALLERROR = "FailedOperation.GatewayRemoteCallError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitApiUseDetailWithContext(ctx context.Context, request *DescribeUnitApiUseDetailRequest) (response *DescribeUnitApiUseDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUnitApiUseDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUnitApiUseDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitNamespacesRequest() (request *DescribeUnitNamespacesRequest) {
    request = &DescribeUnitNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitNamespaces")
    
    
    return
}

func NewDescribeUnitNamespacesResponse() (response *DescribeUnitNamespacesResponse) {
    response = &DescribeUnitNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnitNamespaces
// 查询单元化命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitNamespaces(request *DescribeUnitNamespacesRequest) (response *DescribeUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeUnitNamespacesRequest()
    }
    
    response = NewDescribeUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

// DescribeUnitNamespaces
// 查询单元化命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitNamespacesWithContext(ctx context.Context, request *DescribeUnitNamespacesRequest) (response *DescribeUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeUnitNamespacesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitRuleRequest() (request *DescribeUnitRuleRequest) {
    request = &DescribeUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRule")
    
    
    return
}

func NewDescribeUnitRuleResponse() (response *DescribeUnitRuleResponse) {
    response = &DescribeUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnitRule
// 查询单元化规则详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRule(request *DescribeUnitRuleRequest) (response *DescribeUnitRuleResponse, err error) {
    if request == nil {
        request = NewDescribeUnitRuleRequest()
    }
    
    response = NewDescribeUnitRuleResponse()
    err = c.Send(request, response)
    return
}

// DescribeUnitRule
// 查询单元化规则详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRuleWithContext(ctx context.Context, request *DescribeUnitRuleRequest) (response *DescribeUnitRuleResponse, err error) {
    if request == nil {
        request = NewDescribeUnitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnitRulesRequest() (request *DescribeUnitRulesRequest) {
    request = &DescribeUnitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRules")
    
    
    return
}

func NewDescribeUnitRulesResponse() (response *DescribeUnitRulesResponse) {
    response = &DescribeUnitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUnitRules
// 查询单元化规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRules(request *DescribeUnitRulesRequest) (response *DescribeUnitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeUnitRulesRequest()
    }
    
    response = NewDescribeUnitRulesResponse()
    err = c.Send(request, response)
    return
}

// DescribeUnitRules
// 查询单元化规则列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUnitRulesWithContext(ctx context.Context, request *DescribeUnitRulesRequest) (response *DescribeUnitRulesResponse, err error) {
    if request == nil {
        request = NewDescribeUnitRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUnitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadInfoRequest() (request *DescribeUploadInfoRequest) {
    request = &DescribeUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUploadInfo")
    
    
    return
}

func NewDescribeUploadInfoResponse() (response *DescribeUploadInfoResponse) {
    response = &DescribeUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUploadInfo
// TSF会将软件包上传到腾讯云对象存储（COS）。调用此接口获取上传信息，如目标地域，桶，包Id，存储路径，鉴权信息等，之后请使用COS API（或SDK）进行上传。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
//  RESOURCEINSUFFICIENT_PACKAGESPACEFULL = "ResourceInsufficient.PackageSpaceFull"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest) (response *DescribeUploadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadInfoRequest()
    }
    
    response = NewDescribeUploadInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeUploadInfo
// TSF会将软件包上传到腾讯云对象存储（COS）。调用此接口获取上传信息，如目标地域，桶，包Id，存储路径，鉴权信息等，之后请使用COS API（或SDK）进行上传。
//
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCEINUSE_OBJECTEXIST = "ResourceInUse.ObjectExist"
//  RESOURCEINSUFFICIENT_PACKAGESPACEFULL = "ResourceInsufficient.PackageSpaceFull"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) DescribeUploadInfoWithContext(ctx context.Context, request *DescribeUploadInfoRequest) (response *DescribeUploadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsableUnitNamespacesRequest() (request *DescribeUsableUnitNamespacesRequest) {
    request = &DescribeUsableUnitNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DescribeUsableUnitNamespaces")
    
    
    return
}

func NewDescribeUsableUnitNamespacesResponse() (response *DescribeUsableUnitNamespacesResponse) {
    response = &DescribeUsableUnitNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsableUnitNamespaces
// 查询可用于被导入的命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUsableUnitNamespaces(request *DescribeUsableUnitNamespacesRequest) (response *DescribeUsableUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsableUnitNamespacesRequest()
    }
    
    response = NewDescribeUsableUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

// DescribeUsableUnitNamespaces
// 查询可用于被导入的命名空间列表
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DescribeUsableUnitNamespacesWithContext(ctx context.Context, request *DescribeUsableUnitNamespacesRequest) (response *DescribeUsableUnitNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeUsableUnitNamespacesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUsableUnitNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTaskRequest() (request *DisableTaskRequest) {
    request = &DisableTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableTask")
    
    
    return
}

func NewDisableTaskResponse() (response *DisableTaskResponse) {
    response = &DisableTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableTask
// 停用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTask(request *DisableTaskRequest) (response *DisableTaskResponse, err error) {
    if request == nil {
        request = NewDisableTaskRequest()
    }
    
    response = NewDisableTaskResponse()
    err = c.Send(request, response)
    return
}

// DisableTask
// 停用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTaskWithContext(ctx context.Context, request *DisableTaskRequest) (response *DisableTaskResponse, err error) {
    if request == nil {
        request = NewDisableTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisableTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDisableTaskFlowRequest() (request *DisableTaskFlowRequest) {
    request = &DisableTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableTaskFlow")
    
    
    return
}

func NewDisableTaskFlowResponse() (response *DisableTaskFlowResponse) {
    response = &DisableTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableTaskFlow
// 停用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTaskFlow(request *DisableTaskFlowRequest) (response *DisableTaskFlowResponse, err error) {
    if request == nil {
        request = NewDisableTaskFlowRequest()
    }
    
    response = NewDisableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

// DisableTaskFlow
// 停用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) DisableTaskFlowWithContext(ctx context.Context, request *DisableTaskFlowRequest) (response *DisableTaskFlowResponse, err error) {
    if request == nil {
        request = NewDisableTaskFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDisableUnitRouteRequest() (request *DisableUnitRouteRequest) {
    request = &DisableUnitRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableUnitRoute")
    
    
    return
}

func NewDisableUnitRouteResponse() (response *DisableUnitRouteResponse) {
    response = &DisableUnitRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableUnitRoute
// 禁用单元化路由
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRoute(request *DisableUnitRouteRequest) (response *DisableUnitRouteResponse, err error) {
    if request == nil {
        request = NewDisableUnitRouteRequest()
    }
    
    response = NewDisableUnitRouteResponse()
    err = c.Send(request, response)
    return
}

// DisableUnitRoute
// 禁用单元化路由
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRouteWithContext(ctx context.Context, request *DisableUnitRouteRequest) (response *DisableUnitRouteResponse, err error) {
    if request == nil {
        request = NewDisableUnitRouteRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisableUnitRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDisableUnitRuleRequest() (request *DisableUnitRuleRequest) {
    request = &DisableUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DisableUnitRule")
    
    
    return
}

func NewDisableUnitRuleResponse() (response *DisableUnitRuleResponse) {
    response = &DisableUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableUnitRule
// 禁用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRule(request *DisableUnitRuleRequest) (response *DisableUnitRuleResponse, err error) {
    if request == nil {
        request = NewDisableUnitRuleRequest()
    }
    
    response = NewDisableUnitRuleResponse()
    err = c.Send(request, response)
    return
}

// DisableUnitRule
// 禁用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) DisableUnitRuleWithContext(ctx context.Context, request *DisableUnitRuleRequest) (response *DisableUnitRuleResponse, err error) {
    if request == nil {
        request = NewDisableUnitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisableUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDraftApiGroupRequest() (request *DraftApiGroupRequest) {
    request = &DraftApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "DraftApiGroup")
    
    
    return
}

func NewDraftApiGroupResponse() (response *DraftApiGroupResponse) {
    response = &DraftApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DraftApiGroup
// 下线Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DraftApiGroup(request *DraftApiGroupRequest) (response *DraftApiGroupResponse, err error) {
    if request == nil {
        request = NewDraftApiGroupRequest()
    }
    
    response = NewDraftApiGroupResponse()
    err = c.Send(request, response)
    return
}

// DraftApiGroup
// 下线Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) DraftApiGroupWithContext(ctx context.Context, request *DraftApiGroupRequest) (response *DraftApiGroupResponse, err error) {
    if request == nil {
        request = NewDraftApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDraftApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTaskRequest() (request *EnableTaskRequest) {
    request = &EnableTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableTask")
    
    
    return
}

func NewEnableTaskResponse() (response *EnableTaskResponse) {
    response = &EnableTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableTask
// 启用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTask(request *EnableTaskRequest) (response *EnableTaskResponse, err error) {
    if request == nil {
        request = NewEnableTaskRequest()
    }
    
    response = NewEnableTaskResponse()
    err = c.Send(request, response)
    return
}

// EnableTask
// 启用任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTaskWithContext(ctx context.Context, request *EnableTaskRequest) (response *EnableTaskResponse, err error) {
    if request == nil {
        request = NewEnableTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewEnableTaskResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTaskFlowRequest() (request *EnableTaskFlowRequest) {
    request = &EnableTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableTaskFlow")
    
    
    return
}

func NewEnableTaskFlowResponse() (response *EnableTaskFlowResponse) {
    response = &EnableTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableTaskFlow
// 启用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTaskFlow(request *EnableTaskFlowRequest) (response *EnableTaskFlowResponse, err error) {
    if request == nil {
        request = NewEnableTaskFlowRequest()
    }
    
    response = NewEnableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

// EnableTaskFlow
// 启用工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) EnableTaskFlowWithContext(ctx context.Context, request *EnableTaskFlowRequest) (response *EnableTaskFlowResponse, err error) {
    if request == nil {
        request = NewEnableTaskFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewEnableTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewEnableUnitRouteRequest() (request *EnableUnitRouteRequest) {
    request = &EnableUnitRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableUnitRoute")
    
    
    return
}

func NewEnableUnitRouteResponse() (response *EnableUnitRouteResponse) {
    response = &EnableUnitRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableUnitRoute
// 启用单元化路由
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRoute(request *EnableUnitRouteRequest) (response *EnableUnitRouteResponse, err error) {
    if request == nil {
        request = NewEnableUnitRouteRequest()
    }
    
    response = NewEnableUnitRouteResponse()
    err = c.Send(request, response)
    return
}

// EnableUnitRoute
// 启用单元化路由
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRouteWithContext(ctx context.Context, request *EnableUnitRouteRequest) (response *EnableUnitRouteResponse, err error) {
    if request == nil {
        request = NewEnableUnitRouteRequest()
    }
    request.SetContext(ctx)
    
    response = NewEnableUnitRouteResponse()
    err = c.Send(request, response)
    return
}

func NewEnableUnitRuleRequest() (request *EnableUnitRuleRequest) {
    request = &EnableUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "EnableUnitRule")
    
    
    return
}

func NewEnableUnitRuleResponse() (response *EnableUnitRuleResponse) {
    response = &EnableUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableUnitRule
// 启用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRule(request *EnableUnitRuleRequest) (response *EnableUnitRuleResponse, err error) {
    if request == nil {
        request = NewEnableUnitRuleRequest()
    }
    
    response = NewEnableUnitRuleResponse()
    err = c.Send(request, response)
    return
}

// EnableUnitRule
// 启用单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) EnableUnitRuleWithContext(ctx context.Context, request *EnableUnitRuleRequest) (response *EnableUnitRuleResponse, err error) {
    if request == nil {
        request = NewEnableUnitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewEnableUnitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskRequest() (request *ExecuteTaskRequest) {
    request = &ExecuteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTask")
    
    
    return
}

func NewExecuteTaskResponse() (response *ExecuteTaskResponse) {
    response = &ExecuteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExecuteTask
// 手动执行一次任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ExecuteTask(request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
    if request == nil {
        request = NewExecuteTaskRequest()
    }
    
    response = NewExecuteTaskResponse()
    err = c.Send(request, response)
    return
}

// ExecuteTask
// 手动执行一次任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_UNSUPPORTACTION = "UnsupportedOperation.UnsupportAction"
func (c *Client) ExecuteTaskWithContext(ctx context.Context, request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
    if request == nil {
        request = NewExecuteTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewExecuteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskFlowRequest() (request *ExecuteTaskFlowRequest) {
    request = &ExecuteTaskFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTaskFlow")
    
    
    return
}

func NewExecuteTaskFlowResponse() (response *ExecuteTaskFlowResponse) {
    response = &ExecuteTaskFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExecuteTaskFlow
// 执行一次工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKPUSHERROR = "FailedOperation.TaskPushError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ExecuteTaskFlow(request *ExecuteTaskFlowRequest) (response *ExecuteTaskFlowResponse, err error) {
    if request == nil {
        request = NewExecuteTaskFlowRequest()
    }
    
    response = NewExecuteTaskFlowResponse()
    err = c.Send(request, response)
    return
}

// ExecuteTaskFlow
// 执行一次工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKPUSHERROR = "FailedOperation.TaskPushError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ExecuteTaskFlowWithContext(ctx context.Context, request *ExecuteTaskFlowRequest) (response *ExecuteTaskFlowResponse, err error) {
    if request == nil {
        request = NewExecuteTaskFlowRequest()
    }
    request.SetContext(ctx)
    
    response = NewExecuteTaskFlowResponse()
    err = c.Send(request, response)
    return
}

func NewExpandGroupRequest() (request *ExpandGroupRequest) {
    request = &ExpandGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ExpandGroup")
    
    
    return
}

func NewExpandGroupResponse() (response *ExpandGroupResponse) {
    response = &ExpandGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExpandGroup
// 虚拟机部署组添加实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_CVMCAEMASTERUNKNOWNINSTANCESTATUS = "InvalidParameter.CvmCaeMasterUnknownInstanceStatus"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ExpandGroup(request *ExpandGroupRequest) (response *ExpandGroupResponse, err error) {
    if request == nil {
        request = NewExpandGroupRequest()
    }
    
    response = NewExpandGroupResponse()
    err = c.Send(request, response)
    return
}

// ExpandGroup
// 虚拟机部署组添加实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_CVMCAEMASTERUNKNOWNINSTANCESTATUS = "InvalidParameter.CvmCaeMasterUnknownInstanceStatus"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ExpandGroupWithContext(ctx context.Context, request *ExpandGroupRequest) (response *ExpandGroupResponse, err error) {
    if request == nil {
        request = NewExpandGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewExpandGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerGroupRequest() (request *ModifyContainerGroupRequest) {
    request = &ModifyContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerGroup")
    
    
    return
}

func NewModifyContainerGroupResponse() (response *ModifyContainerGroupResponse) {
    response = &ModifyContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyContainerGroup
// 修改容器部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
func (c *Client) ModifyContainerGroup(request *ModifyContainerGroupRequest) (response *ModifyContainerGroupResponse, err error) {
    if request == nil {
        request = NewModifyContainerGroupRequest()
    }
    
    response = NewModifyContainerGroupResponse()
    err = c.Send(request, response)
    return
}

// ModifyContainerGroup
// 修改容器部署组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONTAINERGROUPACCESSTYPENULL = "InvalidParameterValue.ContainergroupAccesstypeNull"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPGROUPIDNULL = "InvalidParameterValue.ContainergroupGroupidNull"
func (c *Client) ModifyContainerGroupWithContext(ctx context.Context, request *ModifyContainerGroupRequest) (response *ModifyContainerGroupResponse, err error) {
    if request == nil {
        request = NewModifyContainerGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyContainerReplicasRequest() (request *ModifyContainerReplicasRequest) {
    request = &ModifyContainerReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerReplicas")
    
    
    return
}

func NewModifyContainerReplicasResponse() (response *ModifyContainerReplicasResponse) {
    response = &ModifyContainerReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyContainerReplicas
// 修改容器部署组实例数
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINSTANCENUMINVALID = "InvalidParameterValue.ContainergroupInstanceNumInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyContainerReplicas(request *ModifyContainerReplicasRequest) (response *ModifyContainerReplicasResponse, err error) {
    if request == nil {
        request = NewModifyContainerReplicasRequest()
    }
    
    response = NewModifyContainerReplicasResponse()
    err = c.Send(request, response)
    return
}

// ModifyContainerReplicas
// 修改容器部署组实例数
//
// 可能返回的错误码:
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPCPULIMITOVER = "InvalidParameterValue.ContainergroupCpulimitOver"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPINSTANCENUMINVALID = "InvalidParameterValue.ContainergroupInstanceNumInvalid"
//  INVALIDPARAMETERVALUE_CONTAINERGROUPMEMLIMITOVER = "InvalidParameterValue.ContainergroupMemlimitOver"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyContainerReplicasWithContext(ctx context.Context, request *ModifyContainerReplicasRequest) (response *ModifyContainerReplicasResponse, err error) {
    if request == nil {
        request = NewModifyContainerReplicasRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyContainerReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaneRequest() (request *ModifyLaneRequest) {
    request = &ModifyLaneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyLane")
    
    
    return
}

func NewModifyLaneResponse() (response *ModifyLaneResponse) {
    response = &ModifyLaneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLane
// 更新泳道信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) ModifyLane(request *ModifyLaneRequest) (response *ModifyLaneResponse, err error) {
    if request == nil {
        request = NewModifyLaneRequest()
    }
    
    response = NewModifyLaneResponse()
    err = c.Send(request, response)
    return
}

// ModifyLane
// 更新泳道信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) ModifyLaneWithContext(ctx context.Context, request *ModifyLaneRequest) (response *ModifyLaneResponse, err error) {
    if request == nil {
        request = NewModifyLaneRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyLaneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaneRuleRequest() (request *ModifyLaneRuleRequest) {
    request = &ModifyLaneRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyLaneRule")
    
    
    return
}

func NewModifyLaneRuleResponse() (response *ModifyLaneRuleResponse) {
    response = &ModifyLaneRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLaneRule
// 更新泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) ModifyLaneRule(request *ModifyLaneRuleRequest) (response *ModifyLaneRuleResponse, err error) {
    if request == nil {
        request = NewModifyLaneRuleRequest()
    }
    
    response = NewModifyLaneRuleResponse()
    err = c.Send(request, response)
    return
}

// ModifyLaneRule
// 更新泳道规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_LANEINFODELETECONSULFAILED = "FailedOperation.LaneInfoDeleteConsulFailed"
//  FAILEDOPERATION_LANEINFOGROUPNOTEMPTY = "FailedOperation.LaneInfoGroupNotEmpty"
//  FAILEDOPERATION_LANEINFORELEASECONSULFAILED = "FailedOperation.LaneInfoReleaseConsulFailed"
//  FAILEDOPERATION_LANERULEENABLECONSULFAILED = "FailedOperation.LaneRuleEnableConsulFailed"
//  FAILEDOPERATION_LANERULEMAXLIMIT = "FailedOperation.LaneRuleMaxLimit"
//  INVALIDPARAMETER_LANEINFOALREADYUSED = "InvalidParameter.LaneInfoAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEALREADYUSED = "InvalidParameter.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETER_LANEINFONAMEINVALID = "InvalidParameter.LaneInfoNameInvalid"
//  INVALIDPARAMETER_LANEINFONAMENOTEMPTY = "InvalidParameter.LaneInfoNameNotEmpty"
//  INVALIDPARAMETER_LANEINFONAMETOOLONG = "InvalidParameter.LaneInfoNameTooLong"
//  INVALIDPARAMETER_LANEINFONOTEXIST = "InvalidParameter.LaneInfoNotExist"
//  INVALIDPARAMETER_LANEINFONOTEXISTENTRANCE = "InvalidParameter.LaneInfoNotExistEntrance"
//  INVALIDPARAMETER_LANEINFOREMARKTOOLONG = "InvalidParameter.LaneInfoRemarkTooLong"
//  INVALIDPARAMETER_LANERULEINFONOTEXIST = "InvalidParameter.LaneRuleInfoNotExist"
//  INVALIDPARAMETER_LANERULENAMEALREADYUSED = "InvalidParameter.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETER_LANERULENAMEINVALID = "InvalidParameter.LaneRuleNameInvalid"
//  INVALIDPARAMETER_LANERULENAMENOTEMPTY = "InvalidParameter.LaneRuleNameNotEmpty"
//  INVALIDPARAMETER_LANERULENAMETOOLONG = "InvalidParameter.LaneRuleNameTooLong"
//  INVALIDPARAMETER_LANERULENOTEXIST = "InvalidParameter.LaneRuleNotExist"
//  INVALIDPARAMETER_LANERULEREMARKTOOLONG = "InvalidParameter.LaneRuleRemarkTooLong"
//  INVALIDPARAMETER_LANERULETAGNAMENOTEMPTY = "InvalidParameter.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETER_LANERULETAGNAMETOOLONG = "InvalidParameter.LaneRuleTagNameTooLong"
//  INVALIDPARAMETER_LANERULETAGNOTEMPTY = "InvalidParameter.LaneRuleTagNotEmpty"
//  INVALIDPARAMETER_LANERULETAGVALUETOOLONG = "InvalidParameter.LaneRuleTagValueTooLong"
//  INVALIDPARAMETER_LANERULETAGVALUETOTALTOOLONG = "InvalidParameter.LaneRuleTagValueTotalTooLong"
//  INVALIDPARAMETERVALUE_LANEINFOALREADYUSED = "InvalidParameterValue.LaneInfoAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEALREADYUSED = "InvalidParameterValue.LaneInfoNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANEINFONAMEINVALID = "InvalidParameterValue.LaneInfoNameInvalid"
//  INVALIDPARAMETERVALUE_LANEINFONAMENOTEMPTY = "InvalidParameterValue.LaneInfoNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANEINFONAMETOOLONG = "InvalidParameterValue.LaneInfoNameTooLong"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXIST = "InvalidParameterValue.LaneInfoNotExist"
//  INVALIDPARAMETERVALUE_LANEINFONOTEXISTENTRANCE = "InvalidParameterValue.LaneInfoNotExistEntrance"
//  INVALIDPARAMETERVALUE_LANEINFOREMARKTOOLONG = "InvalidParameterValue.LaneInfoRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULEINFONOTEXIST = "InvalidParameterValue.LaneRuleInfoNotExist"
//  INVALIDPARAMETERVALUE_LANERULENAMEALREADYUSED = "InvalidParameterValue.LaneRuleNameAlreadyUsed"
//  INVALIDPARAMETERVALUE_LANERULENAMEINVALID = "InvalidParameterValue.LaneRuleNameInvalid"
//  INVALIDPARAMETERVALUE_LANERULENAMENOTEMPTY = "InvalidParameterValue.LaneRuleNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULENAMETOOLONG = "InvalidParameterValue.LaneRuleNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULENOTEXIST = "InvalidParameterValue.LaneRuleNotExist"
//  INVALIDPARAMETERVALUE_LANERULEREMARKTOOLONG = "InvalidParameterValue.LaneRuleRemarkTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMENOTEMPTY = "InvalidParameterValue.LaneRuleTagNameNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGNAMETOOLONG = "InvalidParameterValue.LaneRuleTagNameTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGNOTEMPTY = "InvalidParameterValue.LaneRuleTagNotEmpty"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOOLONG = "InvalidParameterValue.LaneRuleTagValueTooLong"
//  INVALIDPARAMETERVALUE_LANERULETAGVALUETOTALTOOLONG = "InvalidParameterValue.LaneRuleTagValueTotalTooLong"
func (c *Client) ModifyLaneRuleWithContext(ctx context.Context, request *ModifyLaneRuleRequest) (response *ModifyLaneRuleResponse, err error) {
    if request == nil {
        request = NewModifyLaneRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyLaneRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMicroserviceRequest() (request *ModifyMicroserviceRequest) {
    request = &ModifyMicroserviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyMicroservice")
    
    
    return
}

func NewModifyMicroserviceResponse() (response *ModifyMicroserviceResponse) {
    response = &ModifyMicroserviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMicroservice
// 修改微服务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) ModifyMicroservice(request *ModifyMicroserviceRequest) (response *ModifyMicroserviceResponse, err error) {
    if request == nil {
        request = NewModifyMicroserviceRequest()
    }
    
    response = NewModifyMicroserviceResponse()
    err = c.Send(request, response)
    return
}

// ModifyMicroservice
// 修改微服务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SERVICENOTEXIST = "ResourceNotFound.ServiceNotExist"
func (c *Client) ModifyMicroserviceWithContext(ctx context.Context, request *ModifyMicroserviceRequest) (response *ModifyMicroserviceResponse, err error) {
    if request == nil {
        request = NewModifyMicroserviceRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyMicroserviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPathRewriteRequest() (request *ModifyPathRewriteRequest) {
    request = &ModifyPathRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyPathRewrite")
    
    
    return
}

func NewModifyPathRewriteResponse() (response *ModifyPathRewriteResponse) {
    response = &ModifyPathRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPathRewrite
// 修改路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) ModifyPathRewrite(request *ModifyPathRewriteRequest) (response *ModifyPathRewriteResponse, err error) {
    if request == nil {
        request = NewModifyPathRewriteRequest()
    }
    
    response = NewModifyPathRewriteResponse()
    err = c.Send(request, response)
    return
}

// ModifyPathRewrite
// 修改路径重写
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) ModifyPathRewriteWithContext(ctx context.Context, request *ModifyPathRewriteRequest) (response *ModifyPathRewriteResponse, err error) {
    if request == nil {
        request = NewModifyPathRewriteRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyPathRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskRequest() (request *ModifyTaskRequest) {
    request = &ModifyTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyTask")
    
    
    return
}

func NewModifyTaskResponse() (response *ModifyTaskResponse) {
    response = &ModifyTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTask
// 修改任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyTask(request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
    if request == nil {
        request = NewModifyTaskRequest()
    }
    
    response = NewModifyTaskResponse()
    err = c.Send(request, response)
    return
}

// ModifyTask
// 修改任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKUPDATEERROR = "FailedOperation.TaskUpdateError"
//  INVALIDPARAMETER_BADREQUEST = "InvalidParameter.BadRequest"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERFORMAT = "InvalidParameterValue.InvalidParameterFormat"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ModifyTaskWithContext(ctx context.Context, request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
    if request == nil {
        request = NewModifyTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUploadInfoRequest() (request *ModifyUploadInfoRequest) {
    request = &ModifyUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ModifyUploadInfo")
    
    
    return
}

func NewModifyUploadInfoResponse() (response *ModifyUploadInfoResponse) {
    response = &ModifyUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUploadInfo
// 调用该接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
//
// 调用此接口完成后，才标志上传包流程结束。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) ModifyUploadInfo(request *ModifyUploadInfoRequest) (response *ModifyUploadInfoResponse, err error) {
    if request == nil {
        request = NewModifyUploadInfoRequest()
    }
    
    response = NewModifyUploadInfoResponse()
    err = c.Send(request, response)
    return
}

// ModifyUploadInfo
// 调用该接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
//
// 调用此接口完成后，才标志上传包流程结束。
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND_OBJECTNOEXIST = "ResourceNotFound.ObjectNoExist"
func (c *Client) ModifyUploadInfoWithContext(ctx context.Context, request *ModifyUploadInfoRequest) (response *ModifyUploadInfoResponse, err error) {
    if request == nil {
        request = NewModifyUploadInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewOperateApplicationTcrBindingRequest() (request *OperateApplicationTcrBindingRequest) {
    request = &OperateApplicationTcrBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "OperateApplicationTcrBinding")
    
    
    return
}

func NewOperateApplicationTcrBindingResponse() (response *OperateApplicationTcrBindingResponse) {
    response = &OperateApplicationTcrBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperateApplicationTcrBinding
// 绑定解绑tcr仓库
//
// 可能返回的错误码:
//  INTERNALERROR_IMAGEREPOTCRBINDERROR = "InternalError.ImagerepoTcrBindError"
func (c *Client) OperateApplicationTcrBinding(request *OperateApplicationTcrBindingRequest) (response *OperateApplicationTcrBindingResponse, err error) {
    if request == nil {
        request = NewOperateApplicationTcrBindingRequest()
    }
    
    response = NewOperateApplicationTcrBindingResponse()
    err = c.Send(request, response)
    return
}

// OperateApplicationTcrBinding
// 绑定解绑tcr仓库
//
// 可能返回的错误码:
//  INTERNALERROR_IMAGEREPOTCRBINDERROR = "InternalError.ImagerepoTcrBindError"
func (c *Client) OperateApplicationTcrBindingWithContext(ctx context.Context, request *OperateApplicationTcrBindingRequest) (response *OperateApplicationTcrBindingResponse, err error) {
    if request == nil {
        request = NewOperateApplicationTcrBindingRequest()
    }
    request.SetContext(ctx)
    
    response = NewOperateApplicationTcrBindingResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskRequest() (request *RedoTaskRequest) {
    request = &RedoTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTask")
    
    
    return
}

func NewRedoTaskResponse() (response *RedoTaskResponse) {
    response = &RedoTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RedoTask
// 重新执行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_TASKNOTSUPPORTED = "UnsupportedOperation.TaskNotSupported"
func (c *Client) RedoTask(request *RedoTaskRequest) (response *RedoTaskResponse, err error) {
    if request == nil {
        request = NewRedoTaskRequest()
    }
    
    response = NewRedoTaskResponse()
    err = c.Send(request, response)
    return
}

// RedoTask
// 重新执行任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
//  UNSUPPORTEDOPERATION_TASKNOTSUPPORTED = "UnsupportedOperation.TaskNotSupported"
func (c *Client) RedoTaskWithContext(ctx context.Context, request *RedoTaskRequest) (response *RedoTaskResponse, err error) {
    if request == nil {
        request = NewRedoTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewRedoTaskResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskBatchRequest() (request *RedoTaskBatchRequest) {
    request = &RedoTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskBatch")
    
    
    return
}

func NewRedoTaskBatchResponse() (response *RedoTaskBatchResponse) {
    response = &RedoTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RedoTaskBatch
// 重新执行任务批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskBatch(request *RedoTaskBatchRequest) (response *RedoTaskBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskBatchRequest()
    }
    
    response = NewRedoTaskBatchResponse()
    err = c.Send(request, response)
    return
}

// RedoTaskBatch
// 重新执行任务批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskBatchWithContext(ctx context.Context, request *RedoTaskBatchRequest) (response *RedoTaskBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewRedoTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskExecuteRequest() (request *RedoTaskExecuteRequest) {
    request = &RedoTaskExecuteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskExecute")
    
    
    return
}

func NewRedoTaskExecuteResponse() (response *RedoTaskExecuteResponse) {
    response = &RedoTaskExecuteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RedoTaskExecute
// 重新执行在某个节点上执行任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskExecute(request *RedoTaskExecuteRequest) (response *RedoTaskExecuteResponse, err error) {
    if request == nil {
        request = NewRedoTaskExecuteRequest()
    }
    
    response = NewRedoTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

// RedoTaskExecute
// 重新执行在某个节点上执行任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskExecuteWithContext(ctx context.Context, request *RedoTaskExecuteRequest) (response *RedoTaskExecuteResponse, err error) {
    if request == nil {
        request = NewRedoTaskExecuteRequest()
    }
    request.SetContext(ctx)
    
    response = NewRedoTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

func NewRedoTaskFlowBatchRequest() (request *RedoTaskFlowBatchRequest) {
    request = &RedoTaskFlowBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskFlowBatch")
    
    
    return
}

func NewRedoTaskFlowBatchResponse() (response *RedoTaskFlowBatchResponse) {
    response = &RedoTaskFlowBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RedoTaskFlowBatch
// 重新执行工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskFlowBatch(request *RedoTaskFlowBatchRequest) (response *RedoTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskFlowBatchRequest()
    }
    
    response = NewRedoTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

// RedoTaskFlowBatch
// 重新执行工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RedoTaskFlowBatchWithContext(ctx context.Context, request *RedoTaskFlowBatchRequest) (response *RedoTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewRedoTaskFlowBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewRedoTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseApiGroupRequest() (request *ReleaseApiGroupRequest) {
    request = &ReleaseApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseApiGroup")
    
    
    return
}

func NewReleaseApiGroupResponse() (response *ReleaseApiGroupResponse) {
    response = &ReleaseApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseApiGroup
// 发布Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseApiGroup(request *ReleaseApiGroupRequest) (response *ReleaseApiGroupResponse, err error) {
    if request == nil {
        request = NewReleaseApiGroupRequest()
    }
    
    response = NewReleaseApiGroupResponse()
    err = c.Send(request, response)
    return
}

// ReleaseApiGroup
// 发布Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseApiGroupWithContext(ctx context.Context, request *ReleaseApiGroupRequest) (response *ReleaseApiGroupResponse, err error) {
    if request == nil {
        request = NewReleaseApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewReleaseApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseConfigRequest() (request *ReleaseConfigRequest) {
    request = &ReleaseConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseConfig")
    
    
    return
}

func NewReleaseConfigResponse() (response *ReleaseConfigResponse) {
    response = &ReleaseConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseConfig
// 发布配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseConfig(request *ReleaseConfigRequest) (response *ReleaseConfigResponse, err error) {
    if request == nil {
        request = NewReleaseConfigRequest()
    }
    
    response = NewReleaseConfigResponse()
    err = c.Send(request, response)
    return
}

// ReleaseConfig
// 发布配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGGROUPQUERYFAILED = "FailedOperation.ConfigGroupQueryFailed"
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGGROUPAPPLICATIONIDNOTMATCH = "InvalidParameterValue.ConfigGroupApplicationIdNotMatch"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) ReleaseConfigWithContext(ctx context.Context, request *ReleaseConfigRequest) (response *ReleaseConfigResponse, err error) {
    if request == nil {
        request = NewReleaseConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewReleaseConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseFileConfigRequest() (request *ReleaseFileConfigRequest) {
    request = &ReleaseFileConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleaseFileConfig")
    
    
    return
}

func NewReleaseFileConfigResponse() (response *ReleaseFileConfigResponse) {
    response = &ReleaseFileConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseFileConfig
// 发布文件配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_FILECONFIGALREADYRELEASED = "InvalidParameterValue.FileConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_FILECONFIGPATHEXISTS = "InvalidParameterValue.FileConfigPathExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleaseFileConfig(request *ReleaseFileConfigRequest) (response *ReleaseFileConfigResponse, err error) {
    if request == nil {
        request = NewReleaseFileConfigRequest()
    }
    
    response = NewReleaseFileConfigResponse()
    err = c.Send(request, response)
    return
}

// ReleaseFileConfig
// 发布文件配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_FILECONFIGALREADYRELEASED = "InvalidParameterValue.FileConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_FILECONFIGPATHEXISTS = "InvalidParameterValue.FileConfigPathExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleaseFileConfigWithContext(ctx context.Context, request *ReleaseFileConfigRequest) (response *ReleaseFileConfigResponse, err error) {
    if request == nil {
        request = NewReleaseFileConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewReleaseFileConfigResponse()
    err = c.Send(request, response)
    return
}

func NewReleasePublicConfigRequest() (request *ReleasePublicConfigRequest) {
    request = &ReleasePublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ReleasePublicConfig")
    
    
    return
}

func NewReleasePublicConfigResponse() (response *ReleasePublicConfigResponse) {
    response = &ReleasePublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleasePublicConfig
// 发布公共配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_NAMESPACENOTEXISTS = "InvalidParameterValue.NamespaceNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleasePublicConfig(request *ReleasePublicConfigRequest) (response *ReleasePublicConfigResponse, err error) {
    if request == nil {
        request = NewReleasePublicConfigRequest()
    }
    
    response = NewReleasePublicConfigResponse()
    err = c.Send(request, response)
    return
}

// ReleasePublicConfig
// 发布公共配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGNAMESPACEQUERYFAILED = "FailedOperation.ConfigNamespaceQueryFailed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CONFIGALREADYRELEASED = "InvalidParameterValue.ConfigAlreadyReleased"
//  INVALIDPARAMETERVALUE_CONFIGNOTEXISTSORPERMISSIONDENIED = "InvalidParameterValue.ConfigNotExistsOrPermissionDenied"
//  INVALIDPARAMETERVALUE_CONFIGVALUETOOLONG = "InvalidParameterValue.ConfigValueTooLong"
//  INVALIDPARAMETERVALUE_NAMESPACENOTEXISTS = "InvalidParameterValue.NamespaceNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
func (c *Client) ReleasePublicConfigWithContext(ctx context.Context, request *ReleasePublicConfigRequest) (response *ReleasePublicConfigResponse, err error) {
    if request == nil {
        request = NewReleasePublicConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewReleasePublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveInstancesRequest() (request *RemoveInstancesRequest) {
    request = &RemoveInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RemoveInstances")
    
    
    return
}

func NewRemoveInstancesResponse() (response *RemoveInstancesResponse) {
    response = &RemoveInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveInstances
// 从 TSF 集群中批量移除云主机节点
//
// 可能返回的错误码:
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
    
    response = NewRemoveInstancesResponse()
    err = c.Send(request, response)
    return
}

// RemoveInstances
// 从 TSF 集群中批量移除云主机节点
//
// 可能返回的错误码:
//  INTERNALERROR_TKEAPIFAILEDOPERATION = "InternalError.TkeApiFailedOperation"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_CLUSTERNOTEXIST = "ResourceNotFound.ClusterNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) RemoveInstancesWithContext(ctx context.Context, request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRemoveInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationConfigRequest() (request *RevocationConfigRequest) {
    request = &RevocationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationConfig")
    
    
    return
}

func NewRevocationConfigResponse() (response *RevocationConfigResponse) {
    response = &RevocationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevocationConfig
// 撤回已发布的配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGRELEASEIDREQUIRED = "MissingParameter.ConfigReleaseIdRequired"
func (c *Client) RevocationConfig(request *RevocationConfigRequest) (response *RevocationConfigResponse, err error) {
    if request == nil {
        request = NewRevocationConfigRequest()
    }
    
    response = NewRevocationConfigResponse()
    err = c.Send(request, response)
    return
}

// RevocationConfig
// 撤回已发布的配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
//  INVALIDPARAMETERVALUE_RESOURCEPERMISSIONDENIED = "InvalidParameterValue.ResourcePermissionDenied"
//  MISSINGPARAMETER_CONFIGRELEASEIDREQUIRED = "MissingParameter.ConfigReleaseIdRequired"
func (c *Client) RevocationConfigWithContext(ctx context.Context, request *RevocationConfigRequest) (response *RevocationConfigResponse, err error) {
    if request == nil {
        request = NewRevocationConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewRevocationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRevocationPublicConfigRequest() (request *RevocationPublicConfigRequest) {
    request = &RevocationPublicConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RevocationPublicConfig")
    
    
    return
}

func NewRevocationPublicConfigResponse() (response *RevocationPublicConfigResponse) {
    response = &RevocationPublicConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevocationPublicConfig
// 撤回已发布的公共配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
func (c *Client) RevocationPublicConfig(request *RevocationPublicConfigRequest) (response *RevocationPublicConfigResponse, err error) {
    if request == nil {
        request = NewRevocationPublicConfigRequest()
    }
    
    response = NewRevocationPublicConfigResponse()
    err = c.Send(request, response)
    return
}

// RevocationPublicConfig
// 撤回已发布的公共配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CONFIGRELEASENOTEXISTS = "InvalidParameterValue.ConfigReleaseNotExists"
func (c *Client) RevocationPublicConfigWithContext(ctx context.Context, request *RevocationPublicConfigRequest) (response *RevocationPublicConfigResponse, err error) {
    if request == nil {
        request = NewRevocationPublicConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewRevocationPublicConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackConfigRequest() (request *RollbackConfigRequest) {
    request = &RollbackConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "RollbackConfig")
    
    
    return
}

func NewRollbackConfigResponse() (response *RollbackConfigResponse) {
    response = &RollbackConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RollbackConfig
// 回滚配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
func (c *Client) RollbackConfig(request *RollbackConfigRequest) (response *RollbackConfigResponse, err error) {
    if request == nil {
        request = NewRollbackConfigRequest()
    }
    
    response = NewRollbackConfigResponse()
    err = c.Send(request, response)
    return
}

// RollbackConfig
// 回滚配置
//
// 可能返回的错误码:
//  INTERNALERROR_CONSULSERVERERROR = "InternalError.ConsulServerError"
//  INVALIDPARAMETERVALUE_GROUPNOTEXISTS = "InvalidParameterValue.GroupNotExists"
func (c *Client) RollbackConfigWithContext(ctx context.Context, request *RollbackConfigRequest) (response *RollbackConfigResponse, err error) {
    if request == nil {
        request = NewRollbackConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewRollbackConfigResponse()
    err = c.Send(request, response)
    return
}

func NewSearchBusinessLogRequest() (request *SearchBusinessLogRequest) {
    request = &SearchBusinessLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchBusinessLog")
    
    
    return
}

func NewSearchBusinessLogResponse() (response *SearchBusinessLogResponse) {
    response = &SearchBusinessLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchBusinessLog
// 业务日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmBusiLogSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) SearchBusinessLog(request *SearchBusinessLogRequest) (response *SearchBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchBusinessLogRequest()
    }
    
    response = NewSearchBusinessLogResponse()
    err = c.Send(request, response)
    return
}

// SearchBusinessLog
// 业务日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMBUSILOGCFGAPPRELATIONPARAMERROR = "InvalidParameter.TsfApmBusiLogCfgAppRelationParamError"
//  INVALIDPARAMETER_TSFAPMBUSILOGSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmBusiLogSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) SearchBusinessLogWithContext(ctx context.Context, request *SearchBusinessLogRequest) (response *SearchBusinessLogResponse, err error) {
    if request == nil {
        request = NewSearchBusinessLogRequest()
    }
    request.SetContext(ctx)
    
    response = NewSearchBusinessLogResponse()
    err = c.Send(request, response)
    return
}

func NewSearchStdoutLogRequest() (request *SearchStdoutLogRequest) {
    request = &SearchStdoutLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "SearchStdoutLog")
    
    
    return
}

func NewSearchStdoutLogResponse() (response *SearchStdoutLogResponse) {
    response = &SearchStdoutLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchStdoutLog
// 标准输出日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) SearchStdoutLog(request *SearchStdoutLogRequest) (response *SearchStdoutLogResponse, err error) {
    if request == nil {
        request = NewSearchStdoutLogRequest()
    }
    
    response = NewSearchStdoutLogResponse()
    err = c.Send(request, response)
    return
}

// SearchStdoutLog
// 标准输出日志搜索
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) SearchStdoutLogWithContext(ctx context.Context, request *SearchStdoutLogRequest) (response *SearchStdoutLogResponse, err error) {
    if request == nil {
        request = NewSearchStdoutLogRequest()
    }
    request.SetContext(ctx)
    
    response = NewSearchStdoutLogResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkGroupRequest() (request *ShrinkGroupRequest) {
    request = &ShrinkGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkGroup")
    
    
    return
}

func NewShrinkGroupResponse() (response *ShrinkGroupResponse) {
    response = &ShrinkGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ShrinkGroup
// 下线部署组所有机器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) ShrinkGroup(request *ShrinkGroupRequest) (response *ShrinkGroupResponse, err error) {
    if request == nil {
        request = NewShrinkGroupRequest()
    }
    
    response = NewShrinkGroupResponse()
    err = c.Send(request, response)
    return
}

// ShrinkGroup
// 下线部署组所有机器实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_TSFPRIVILEGEERROR = "FailedOperation.TsfPrivilegeError"
//  INTERNALERROR_TSFAPMESRESPONSESTATUSEXCEPTION = "InternalError.TsfApmEsResponseStatusException"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_TSFAPMSTDOUTSEARCHREQUESTPARAMERROR = "InvalidParameter.TsfApmStdoutSearchRequestParamError"
//  RESOURCENOTFOUND_MICROSERVICEOFFLINE = "ResourceNotFound.MicroserviceOffline"
func (c *Client) ShrinkGroupWithContext(ctx context.Context, request *ShrinkGroupRequest) (response *ShrinkGroupResponse, err error) {
    if request == nil {
        request = NewShrinkGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewShrinkGroupResponse()
    err = c.Send(request, response)
    return
}

func NewShrinkInstancesRequest() (request *ShrinkInstancesRequest) {
    request = &ShrinkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "ShrinkInstances")
    
    
    return
}

func NewShrinkInstancesResponse() (response *ShrinkInstancesResponse) {
    response = &ShrinkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ShrinkInstances
// 虚拟机部署组下线实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ShrinkInstances(request *ShrinkInstancesRequest) (response *ShrinkInstancesResponse, err error) {
    if request == nil {
        request = NewShrinkInstancesRequest()
    }
    
    response = NewShrinkInstancesResponse()
    err = c.Send(request, response)
    return
}

// ShrinkInstances
// 虚拟机部署组下线实例
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_GROUPVALIDINSTANCENULL = "InvalidParameterValue.GroupValidInstanceNull"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) ShrinkInstancesWithContext(ctx context.Context, request *ShrinkInstancesRequest) (response *ShrinkInstancesResponse, err error) {
    if request == nil {
        request = NewShrinkInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewShrinkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartContainerGroupRequest() (request *StartContainerGroupRequest) {
    request = &StartContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartContainerGroup")
    
    
    return
}

func NewStartContainerGroupResponse() (response *StartContainerGroupResponse) {
    response = &StartContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartContainerGroup
// 启动容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASRUN = "FailedOperation.ContainergroupGroupHasrun"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StartContainerGroup(request *StartContainerGroupRequest) (response *StartContainerGroupResponse, err error) {
    if request == nil {
        request = NewStartContainerGroupRequest()
    }
    
    response = NewStartContainerGroupResponse()
    err = c.Send(request, response)
    return
}

// StartContainerGroup
// 启动容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASRUN = "FailedOperation.ContainergroupGroupHasrun"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  INVALIDPARAMETER_KUBERNETESPARAMERROR = "InvalidParameter.KubernetesParamError"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StartContainerGroupWithContext(ctx context.Context, request *StartContainerGroupRequest) (response *StartContainerGroupResponse, err error) {
    if request == nil {
        request = NewStartContainerGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStartGroupRequest() (request *StartGroupRequest) {
    request = &StartGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StartGroup")
    
    
    return
}

func NewStartGroupResponse() (response *StartGroupResponse) {
    response = &StartGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartGroup
// 启动分组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) StartGroup(request *StartGroupRequest) (response *StartGroupResponse, err error) {
    if request == nil {
        request = NewStartGroupRequest()
    }
    
    response = NewStartGroupResponse()
    err = c.Send(request, response)
    return
}

// StartGroup
// 启动分组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERDISPATCHERROR = "InternalError.CvmCaeMasterDispatchError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) StartGroupWithContext(ctx context.Context, request *StartGroupRequest) (response *StartGroupResponse, err error) {
    if request == nil {
        request = NewStartGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopContainerGroupRequest() (request *StopContainerGroupRequest) {
    request = &StopContainerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopContainerGroup")
    
    
    return
}

func NewStopContainerGroupResponse() (response *StopContainerGroupResponse) {
    response = &StopContainerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopContainerGroup
// 停止容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASSTOP = "FailedOperation.ContainergroupGroupHasstop"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopContainerGroup(request *StopContainerGroupRequest) (response *StopContainerGroupResponse, err error) {
    if request == nil {
        request = NewStopContainerGroupRequest()
    }
    
    response = NewStopContainerGroupResponse()
    err = c.Send(request, response)
    return
}

// StopContainerGroup
// 停止容器部署组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTAINERGROUPGROUPHASSTOP = "FailedOperation.ContainergroupGroupHasstop"
//  INTERNALERROR_CONTAINERGROUPKUBERNETEAPIINVOKEERROR = "InternalError.ContainergroupKuberneteApiInvokeError"
//  INTERNALERROR_CONTAINERGROUPKUBERNETECONNECTERROR = "InternalError.ContainergroupKuberneteConnectError"
//  RESOURCENOTFOUND_CONTAINERGROUPGROUPNOTFOUND = "ResourceNotFound.ContainergroupGroupNotFound"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopContainerGroupWithContext(ctx context.Context, request *StopContainerGroupRequest) (response *StopContainerGroupResponse, err error) {
    if request == nil {
        request = NewStopContainerGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopContainerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopGroupRequest() (request *StopGroupRequest) {
    request = &StopGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopGroup")
    
    
    return
}

func NewStopGroupResponse() (response *StopGroupResponse) {
    response = &StopGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopGroup
// 停止虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERGROUPNOAGENT = "InvalidParameterValue.CvmCaeMasterGroupNoAgent"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopGroup(request *StopGroupRequest) (response *StopGroupResponse, err error) {
    if request == nil {
        request = NewStopGroupRequest()
    }
    
    response = NewStopGroupResponse()
    err = c.Send(request, response)
    return
}

// StopGroup
// 停止虚拟机部署组
//
// 可能返回的错误码:
//  INTERNALERROR_CVMCAEMASTERNONALIVE = "InternalError.CvmCaeMasterNonAlive"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERAGENTBUSY = "InvalidParameterValue.CvmCaeMasterAgentBusy"
//  INVALIDPARAMETERVALUE_CVMCAEMASTERGROUPNOAGENT = "InvalidParameterValue.CvmCaeMasterGroupNoAgent"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopGroupWithContext(ctx context.Context, request *StopGroupRequest) (response *StopGroupResponse, err error) {
    if request == nil {
        request = NewStopGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopTaskBatchRequest() (request *StopTaskBatchRequest) {
    request = &StopTaskBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopTaskBatch")
    
    
    return
}

func NewStopTaskBatchResponse() (response *StopTaskBatchResponse) {
    response = &StopTaskBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopTaskBatch
// 停止执行中的任务批次， 非运行中的任务不可调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskBatch(request *StopTaskBatchRequest) (response *StopTaskBatchResponse, err error) {
    if request == nil {
        request = NewStopTaskBatchRequest()
    }
    
    response = NewStopTaskBatchResponse()
    err = c.Send(request, response)
    return
}

// StopTaskBatch
// 停止执行中的任务批次， 非运行中的任务不可调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskBatchWithContext(ctx context.Context, request *StopTaskBatchRequest) (response *StopTaskBatchResponse, err error) {
    if request == nil {
        request = NewStopTaskBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopTaskBatchResponse()
    err = c.Send(request, response)
    return
}

func NewStopTaskExecuteRequest() (request *StopTaskExecuteRequest) {
    request = &StopTaskExecuteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "StopTaskExecute")
    
    
    return
}

func NewStopTaskExecuteResponse() (response *StopTaskExecuteResponse) {
    response = &StopTaskExecuteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopTaskExecute
// 停止正在某个节点上执行的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskExecute(request *StopTaskExecuteRequest) (response *StopTaskExecuteResponse, err error) {
    if request == nil {
        request = NewStopTaskExecuteRequest()
    }
    
    response = NewStopTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

// StopTaskExecute
// 停止正在某个节点上执行的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFAILED = "FailedOperation.TaskOperationFailed"
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  FAILEDOPERATION_TASKQUERYERROR = "FailedOperation.TaskQueryError"
//  FAILEDOPERATION_TASKTERMINATEFAILED = "FailedOperation.TaskTerminateFailed"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INTERNALERROR_UNHANDLEDEXCEPTION = "InternalError.UnhandledException"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) StopTaskExecuteWithContext(ctx context.Context, request *StopTaskExecuteRequest) (response *StopTaskExecuteResponse, err error) {
    if request == nil {
        request = NewStopTaskExecuteRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopTaskExecuteResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTaskFlowBatchRequest() (request *TerminateTaskFlowBatchRequest) {
    request = &TerminateTaskFlowBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "TerminateTaskFlowBatch")
    
    
    return
}

func NewTerminateTaskFlowBatchResponse() (response *TerminateTaskFlowBatchResponse) {
    response = &TerminateTaskFlowBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateTaskFlowBatch
// 停止一个工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) TerminateTaskFlowBatch(request *TerminateTaskFlowBatchRequest) (response *TerminateTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewTerminateTaskFlowBatchRequest()
    }
    
    response = NewTerminateTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

// TerminateTaskFlowBatch
// 停止一个工作流批次
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKOPERATIONFORBIDDEN = "FailedOperation.TaskOperationForbidden"
//  INTERNALERROR_TASKINTERNALERROR = "InternalError.TaskInternalError"
//  INVALIDPARAMETERVALUE_TASKPARAMETERINVALID = "InvalidParameterValue.TaskParameterInvalid"
//  MISSINGPARAMETER_TASKPARAMETERMISSED = "MissingParameter.TaskParameterMissed"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOPERMISSION = "UnauthorizedOperation.CamTsfRoleNoPermission"
//  UNAUTHORIZEDOPERATION_CAMTSFROLENOTEXIST = "UnauthorizedOperation.CamTsfRoleNotExist"
//  UNAUTHORIZEDOPERATION_LICENSEINACTIVE = "UnauthorizedOperation.LicenseInactive"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPRIVILEGE = "UnauthorizedOperation.NoPrivilege"
func (c *Client) TerminateTaskFlowBatchWithContext(ctx context.Context, request *TerminateTaskFlowBatchRequest) (response *TerminateTaskFlowBatchResponse, err error) {
    if request == nil {
        request = NewTerminateTaskFlowBatchRequest()
    }
    request.SetContext(ctx)
    
    response = NewTerminateTaskFlowBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindApiGroupRequest() (request *UnbindApiGroupRequest) {
    request = &UnbindApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UnbindApiGroup")
    
    
    return
}

func NewUnbindApiGroupResponse() (response *UnbindApiGroupResponse) {
    response = &UnbindApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindApiGroup
// API分组批量与网关解绑
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
func (c *Client) UnbindApiGroup(request *UnbindApiGroupRequest) (response *UnbindApiGroupResponse, err error) {
    if request == nil {
        request = NewUnbindApiGroupRequest()
    }
    
    response = NewUnbindApiGroupResponse()
    err = c.Send(request, response)
    return
}

// UnbindApiGroup
// API分组批量与网关解绑
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
func (c *Client) UnbindApiGroupWithContext(ctx context.Context, request *UnbindApiGroupRequest) (response *UnbindApiGroupResponse, err error) {
    if request == nil {
        request = NewUnbindApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewUnbindApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiGroupRequest() (request *UpdateApiGroupRequest) {
    request = &UpdateApiGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiGroup")
    
    
    return
}

func NewUpdateApiGroupResponse() (response *UpdateApiGroupResponse) {
    response = &UpdateApiGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateApiGroup
// 更新Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) UpdateApiGroup(request *UpdateApiGroupRequest) (response *UpdateApiGroupResponse, err error) {
    if request == nil {
        request = NewUpdateApiGroupRequest()
    }
    
    response = NewUpdateApiGroupResponse()
    err = c.Send(request, response)
    return
}

// UpdateApiGroup
// 更新Api分组
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
func (c *Client) UpdateApiGroupWithContext(ctx context.Context, request *UpdateApiGroupRequest) (response *UpdateApiGroupResponse, err error) {
    if request == nil {
        request = NewUpdateApiGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateApiGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiRateLimitRuleRequest() (request *UpdateApiRateLimitRuleRequest) {
    request = &UpdateApiRateLimitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRule")
    
    
    return
}

func NewUpdateApiRateLimitRuleResponse() (response *UpdateApiRateLimitRuleResponse) {
    response = &UpdateApiRateLimitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateApiRateLimitRule
// 更新API限流规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"
func (c *Client) UpdateApiRateLimitRule(request *UpdateApiRateLimitRuleRequest) (response *UpdateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRuleRequest()
    }
    
    response = NewUpdateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

// UpdateApiRateLimitRule
// 更新API限流规则
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERERROR = "InvalidParameterValue.GatewayParameterError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  RESOURCEINUSE_RATELIMITRULEEXISTERROR = "ResourceInUse.RatelimitRuleExistError"
func (c *Client) UpdateApiRateLimitRuleWithContext(ctx context.Context, request *UpdateApiRateLimitRuleRequest) (response *UpdateApiRateLimitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateApiRateLimitRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiRateLimitRulesRequest() (request *UpdateApiRateLimitRulesRequest) {
    request = &UpdateApiRateLimitRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRules")
    
    
    return
}

func NewUpdateApiRateLimitRulesResponse() (response *UpdateApiRateLimitRulesResponse) {
    response = &UpdateApiRateLimitRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateApiRateLimitRules
// 批量更新API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiRateLimitRules(request *UpdateApiRateLimitRulesRequest) (response *UpdateApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRulesRequest()
    }
    
    response = NewUpdateApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

// UpdateApiRateLimitRules
// 批量更新API限流规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_RATELIMITCONSULERROR = "FailedOperation.RatelimitConsulError"
//  INTERNALERROR_GATEWAYCONSISTENCYERROR = "InternalError.GatewayConsistencyError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
//  UNAUTHORIZEDOPERATION_LICENSEUNAUTHORIZED = "UnauthorizedOperation.LicenseUnauthorized"
func (c *Client) UpdateApiRateLimitRulesWithContext(ctx context.Context, request *UpdateApiRateLimitRulesRequest) (response *UpdateApiRateLimitRulesResponse, err error) {
    if request == nil {
        request = NewUpdateApiRateLimitRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateApiRateLimitRulesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateApiTimeoutsRequest() (request *UpdateApiTimeoutsRequest) {
    request = &UpdateApiTimeoutsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiTimeouts")
    
    
    return
}

func NewUpdateApiTimeoutsResponse() (response *UpdateApiTimeoutsResponse) {
    response = &UpdateApiTimeoutsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateApiTimeouts
// 批量更新API超时
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateApiTimeouts(request *UpdateApiTimeoutsRequest) (response *UpdateApiTimeoutsResponse, err error) {
    if request == nil {
        request = NewUpdateApiTimeoutsRequest()
    }
    
    response = NewUpdateApiTimeoutsResponse()
    err = c.Send(request, response)
    return
}

// UpdateApiTimeouts
// 批量更新API超时
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateApiTimeoutsWithContext(ctx context.Context, request *UpdateApiTimeoutsRequest) (response *UpdateApiTimeoutsResponse, err error) {
    if request == nil {
        request = NewUpdateApiTimeoutsRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateApiTimeoutsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGatewayApiRequest() (request *UpdateGatewayApiRequest) {
    request = &UpdateGatewayApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayApi")
    
    
    return
}

func NewUpdateGatewayApiResponse() (response *UpdateGatewayApiResponse) {
    response = &UpdateGatewayApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateGatewayApi
// 更新API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateGatewayApi(request *UpdateGatewayApiRequest) (response *UpdateGatewayApiResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayApiRequest()
    }
    
    response = NewUpdateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

// UpdateGatewayApi
// 更新API
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateGatewayApiWithContext(ctx context.Context, request *UpdateGatewayApiRequest) (response *UpdateGatewayApiResponse, err error) {
    if request == nil {
        request = NewUpdateGatewayApiRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateGatewayApiResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateHealthCheckSettingsRequest() (request *UpdateHealthCheckSettingsRequest) {
    request = &UpdateHealthCheckSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateHealthCheckSettings")
    
    
    return
}

func NewUpdateHealthCheckSettingsResponse() (response *UpdateHealthCheckSettingsResponse) {
    response = &UpdateHealthCheckSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateHealthCheckSettings
// 更新健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMCAEMASTERHEALTHCHECKCONFIGERROR = "FailedOperation.CvmCaeMasterHealthCheckConfigError"
func (c *Client) UpdateHealthCheckSettings(request *UpdateHealthCheckSettingsRequest) (response *UpdateHealthCheckSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateHealthCheckSettingsRequest()
    }
    
    response = NewUpdateHealthCheckSettingsResponse()
    err = c.Send(request, response)
    return
}

// UpdateHealthCheckSettings
// 更新健康检查配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CVMCAEMASTERHEALTHCHECKCONFIGERROR = "FailedOperation.CvmCaeMasterHealthCheckConfigError"
func (c *Client) UpdateHealthCheckSettingsWithContext(ctx context.Context, request *UpdateHealthCheckSettingsRequest) (response *UpdateHealthCheckSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateHealthCheckSettingsRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateHealthCheckSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRepositoryRequest() (request *UpdateRepositoryRequest) {
    request = &UpdateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateRepository")
    
    
    return
}

func NewUpdateRepositoryResponse() (response *UpdateRepositoryResponse) {
    response = &UpdateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRepository
// 更新仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) UpdateRepository(request *UpdateRepositoryRequest) (response *UpdateRepositoryResponse, err error) {
    if request == nil {
        request = NewUpdateRepositoryRequest()
    }
    
    response = NewUpdateRepositoryResponse()
    err = c.Send(request, response)
    return
}

// UpdateRepository
// 更新仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR_RUNTIMEERROR = "InternalError.RuntimeError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) UpdateRepositoryWithContext(ctx context.Context, request *UpdateRepositoryRequest) (response *UpdateRepositoryResponse, err error) {
    if request == nil {
        request = NewUpdateRepositoryRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUnitRuleRequest() (request *UpdateUnitRuleRequest) {
    request = &UpdateUnitRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tsf", APIVersion, "UpdateUnitRule")
    
    
    return
}

func NewUpdateUnitRuleResponse() (response *UpdateUnitRuleResponse) {
    response = &UpdateUnitRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateUnitRule
// 更新单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateUnitRule(request *UpdateUnitRuleRequest) (response *UpdateUnitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateUnitRuleRequest()
    }
    
    response = NewUpdateUnitRuleResponse()
    err = c.Send(request, response)
    return
}

// UpdateUnitRule
// 更新单元化规则
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYCONSULERROR = "InternalError.GatewayConsulError"
//  INVALIDPARAMETERVALUE_GATEWAYPARAMETERINVALID = "InvalidParameterValue.GatewayParameterInvalid"
//  MISSINGPARAMETER_GATEWAYPARAMETERREQUIRED = "MissingParameter.GatewayParameterRequired"
func (c *Client) UpdateUnitRuleWithContext(ctx context.Context, request *UpdateUnitRuleRequest) (response *UpdateUnitRuleResponse, err error) {
    if request == nil {
        request = NewUpdateUnitRuleRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateUnitRuleResponse()
    err = c.Send(request, response)
    return
}
