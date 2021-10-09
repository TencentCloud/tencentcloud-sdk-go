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

package v20190924

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-24"

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


func NewBatchDeleteImagePersonalRequest() (request *BatchDeleteImagePersonalRequest) {
    request = &BatchDeleteImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "BatchDeleteImagePersonal")
    return
}

func NewBatchDeleteImagePersonalResponse() (response *BatchDeleteImagePersonalResponse) {
    response = &BatchDeleteImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDeleteImagePersonal
// 用于在个人版镜像仓库中批量删除Tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) BatchDeleteImagePersonal(request *BatchDeleteImagePersonalRequest) (response *BatchDeleteImagePersonalResponse, err error) {
    if request == nil {
        request = NewBatchDeleteImagePersonalRequest()
    }
    response = NewBatchDeleteImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteRepositoryPersonalRequest() (request *BatchDeleteRepositoryPersonalRequest) {
    request = &BatchDeleteRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "BatchDeleteRepositoryPersonal")
    return
}

func NewBatchDeleteRepositoryPersonalResponse() (response *BatchDeleteRepositoryPersonalResponse) {
    response = &BatchDeleteRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDeleteRepositoryPersonal
// 用于个人版镜像仓库中批量删除镜像仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) BatchDeleteRepositoryPersonal(request *BatchDeleteRepositoryPersonalRequest) (response *BatchDeleteRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewBatchDeleteRepositoryPersonalRequest()
    }
    response = NewBatchDeleteRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCheckInstanceRequest() (request *CheckInstanceRequest) {
    request = &CheckInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CheckInstance")
    return
}

func NewCheckInstanceResponse() (response *CheckInstanceResponse) {
    response = &CheckInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckInstance
// 用于校验企业版实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckInstance(request *CheckInstanceRequest) (response *CheckInstanceResponse, err error) {
    if request == nil {
        request = NewCheckInstanceRequest()
    }
    response = NewCheckInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckInstanceNameRequest() (request *CheckInstanceNameRequest) {
    request = &CheckInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CheckInstanceName")
    return
}

func NewCheckInstanceNameResponse() (response *CheckInstanceNameResponse) {
    response = &CheckInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckInstanceName
// 检查待创建的实例名称是否符合规范
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckInstanceName(request *CheckInstanceNameRequest) (response *CheckInstanceNameResponse, err error) {
    if request == nil {
        request = NewCheckInstanceNameRequest()
    }
    response = NewCheckInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationTriggerPersonalRequest() (request *CreateApplicationTriggerPersonalRequest) {
    request = &CreateApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateApplicationTriggerPersonal")
    return
}

func NewCreateApplicationTriggerPersonalResponse() (response *CreateApplicationTriggerPersonalResponse) {
    response = &CreateApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationTriggerPersonal
// 用于创建应用更新触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  INVALIDPARAMETER_ERRTRIGGEREXIST = "InvalidParameter.ErrTriggerExist"
//  LIMITEXCEEDED_ERRTRIGGERMAXLIMIT = "LimitExceeded.ErrTriggerMaxLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) CreateApplicationTriggerPersonal(request *CreateApplicationTriggerPersonalRequest) (response *CreateApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewCreateApplicationTriggerPersonalRequest()
    }
    response = NewCreateApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageLifecyclePersonalRequest() (request *CreateImageLifecyclePersonalRequest) {
    request = &CreateImageLifecyclePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImageLifecyclePersonal")
    return
}

func NewCreateImageLifecyclePersonalResponse() (response *CreateImageLifecyclePersonalResponse) {
    response = &CreateImageLifecyclePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImageLifecyclePersonal
// 用于在个人版中创建清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) CreateImageLifecyclePersonal(request *CreateImageLifecyclePersonalRequest) (response *CreateImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewCreateImageLifecyclePersonalRequest()
    }
    response = NewCreateImageLifecyclePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImmutableTagRulesRequest() (request *CreateImmutableTagRulesRequest) {
    request = &CreateImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImmutableTagRules")
    return
}

func NewCreateImmutableTagRulesResponse() (response *CreateImmutableTagRulesResponse) {
    response = &CreateImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImmutableTagRules
// 创建镜像不可变规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
func (c *Client) CreateImmutableTagRules(request *CreateImmutableTagRulesRequest) (response *CreateImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewCreateImmutableTagRulesRequest()
    }
    response = NewCreateImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstance")
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstance
// 创建实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORNAMEILLEGAL = "InvalidParameter.ErrorNameIllegal"
//  INVALIDPARAMETER_ERRORNAMERESERVED = "InvalidParameter.ErrorNameReserved"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceTokenRequest() (request *CreateInstanceTokenRequest) {
    request = &CreateInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstanceToken")
    return
}

func NewCreateInstanceTokenResponse() (response *CreateInstanceTokenResponse) {
    response = &CreateInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceToken
// 创建实例的临时或长期访问凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceToken(request *CreateInstanceTokenRequest) (response *CreateInstanceTokenResponse, err error) {
    if request == nil {
        request = NewCreateInstanceTokenRequest()
    }
    response = NewCreateInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInternalEndpointDnsRequest() (request *CreateInternalEndpointDnsRequest) {
    request = &CreateInternalEndpointDnsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInternalEndpointDns")
    return
}

func NewCreateInternalEndpointDnsResponse() (response *CreateInternalEndpointDnsResponse) {
    response = &CreateInternalEndpointDnsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInternalEndpointDns
// 创建tcr内网私有域名解析
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEPRIVATEZONE = "InternalError.CreatePrivateZone"
//  INTERNALERROR_CREATEPRIVATEZONERECORD = "InternalError.CreatePrivateZoneRecord"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_DELETEPRIVATEZONERECORD = "InternalError.DeletePrivateZoneRecord"
//  INTERNALERROR_DESCRIBEINTERNALENDPOINTDNSSTATUS = "InternalError.DescribeInternalEndpointDnsStatus"
//  INTERNALERROR_DESCRIBEPRIVATEZONELIST = "InternalError.DescribePrivateZoneList"
//  INTERNALERROR_DESCRIBEPRIVATEZONERECORDLIST = "InternalError.DescribePrivateZoneRecordList"
//  INTERNALERROR_DESCRIBEPRIVATEZONESERVICELIST = "InternalError.DescribePrivateZoneServiceList"
//  INTERNALERROR_MODIFYPRIVATEZONEVPC = "InternalError.ModifyPrivateZoneVpc"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInternalEndpointDns(request *CreateInternalEndpointDnsRequest) (response *CreateInternalEndpointDnsResponse, err error) {
    if request == nil {
        request = NewCreateInternalEndpointDnsRequest()
    }
    response = NewCreateInternalEndpointDnsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultipleSecurityPolicyRequest() (request *CreateMultipleSecurityPolicyRequest) {
    request = &CreateMultipleSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateMultipleSecurityPolicy")
    return
}

func NewCreateMultipleSecurityPolicyResponse() (response *CreateMultipleSecurityPolicyResponse) {
    response = &CreateMultipleSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMultipleSecurityPolicy
// 用于在TCR实例中，创建多个白名单策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMultipleSecurityPolicy(request *CreateMultipleSecurityPolicyRequest) (response *CreateMultipleSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateMultipleSecurityPolicyRequest()
    }
    response = NewCreateMultipleSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNamespace
// 用于在企业版中创建命名空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespacePersonalRequest() (request *CreateNamespacePersonalRequest) {
    request = &CreateNamespacePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateNamespacePersonal")
    return
}

func NewCreateNamespacePersonalResponse() (response *CreateNamespacePersonalResponse) {
    response = &CreateNamespacePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNamespacePersonal
// 创建个人版镜像仓库命名空间，此命名空间全局唯一
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNAMESPACEEXIST = "InvalidParameter.ErrNamespaceExist"
//  INVALIDPARAMETER_ERRNAMESPACERESERVED = "InvalidParameter.ErrNamespaceReserved"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) CreateNamespacePersonal(request *CreateNamespacePersonalRequest) (response *CreateNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewCreateNamespacePersonalRequest()
    }
    response = NewCreateNamespacePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReplicationInstanceRequest() (request *CreateReplicationInstanceRequest) {
    request = &CreateReplicationInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateReplicationInstance")
    return
}

func NewCreateReplicationInstanceResponse() (response *CreateReplicationInstanceResponse) {
    response = &CreateReplicationInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReplicationInstance
// 创建从实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReplicationInstance(request *CreateReplicationInstanceRequest) (response *CreateReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewCreateReplicationInstanceRequest()
    }
    response = NewCreateReplicationInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRepositoryRequest() (request *CreateRepositoryRequest) {
    request = &CreateRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateRepository")
    return
}

func NewCreateRepositoryResponse() (response *CreateRepositoryResponse) {
    response = &CreateRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRepository
// 用于企业版创建镜像仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryRequest()
    }
    response = NewCreateRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRepositoryPersonalRequest() (request *CreateRepositoryPersonalRequest) {
    request = &CreateRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateRepositoryPersonal")
    return
}

func NewCreateRepositoryPersonalResponse() (response *CreateRepositoryPersonalResponse) {
    response = &CreateRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRepositoryPersonal
// 用于在个人版仓库中创建镜像仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  INVALIDPARAMETER_ERRREPOEXIST = "InvalidParameter.ErrRepoExist"
//  LIMITEXCEEDED_ERRREPOMAXLIMIT = "LimitExceeded.ErrRepoMaxLimit"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateRepositoryPersonal(request *CreateRepositoryPersonalRequest) (response *CreateRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryPersonalRequest()
    }
    response = NewCreateRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityPolicyRequest() (request *CreateSecurityPolicyRequest) {
    request = &CreateSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSecurityPolicy")
    return
}

func NewCreateSecurityPolicyResponse() (response *CreateSecurityPolicyResponse) {
    response = &CreateSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSecurityPolicy
// 创建实例公网访问白名单策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTagRetentionExecutionRequest() (request *CreateTagRetentionExecutionRequest) {
    request = &CreateTagRetentionExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateTagRetentionExecution")
    return
}

func NewCreateTagRetentionExecutionResponse() (response *CreateTagRetentionExecutionResponse) {
    response = &CreateTagRetentionExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTagRetentionExecution
// 手动执行版本保留
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTagRetentionExecution(request *CreateTagRetentionExecutionRequest) (response *CreateTagRetentionExecutionResponse, err error) {
    if request == nil {
        request = NewCreateTagRetentionExecutionRequest()
    }
    response = NewCreateTagRetentionExecutionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTagRetentionRuleRequest() (request *CreateTagRetentionRuleRequest) {
    request = &CreateTagRetentionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateTagRetentionRule")
    return
}

func NewCreateTagRetentionRuleResponse() (response *CreateTagRetentionRuleResponse) {
    response = &CreateTagRetentionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTagRetentionRule
// 创建版本保留规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTagRetentionRule(request *CreateTagRetentionRuleRequest) (response *CreateTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewCreateTagRetentionRuleRequest()
    }
    response = NewCreateTagRetentionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserPersonalRequest() (request *CreateUserPersonalRequest) {
    request = &CreateUserPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateUserPersonal")
    return
}

func NewCreateUserPersonalResponse() (response *CreateUserPersonalResponse) {
    response = &CreateUserPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserPersonal
// 创建个人用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRUSEREXIST = "InvalidParameter.ErrUserExist"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
func (c *Client) CreateUserPersonal(request *CreateUserPersonalRequest) (response *CreateUserPersonalResponse, err error) {
    if request == nil {
        request = NewCreateUserPersonalRequest()
    }
    response = NewCreateUserPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebhookTriggerRequest() (request *CreateWebhookTriggerRequest) {
    request = &CreateWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "CreateWebhookTrigger")
    return
}

func NewCreateWebhookTriggerResponse() (response *CreateWebhookTriggerResponse) {
    response = &CreateWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWebhookTrigger
// 创建触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) CreateWebhookTrigger(request *CreateWebhookTriggerRequest) (response *CreateWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewCreateWebhookTriggerRequest()
    }
    response = NewCreateWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationTriggerPersonalRequest() (request *DeleteApplicationTriggerPersonalRequest) {
    request = &DeleteApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteApplicationTriggerPersonal")
    return
}

func NewDeleteApplicationTriggerPersonalResponse() (response *DeleteApplicationTriggerPersonalResponse) {
    response = &DeleteApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationTriggerPersonal
// 用于删除应用更新触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ERRNOTRIGGER = "ResourceNotFound.ErrNoTrigger"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DeleteApplicationTriggerPersonal(request *DeleteApplicationTriggerPersonalRequest) (response *DeleteApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationTriggerPersonalRequest()
    }
    response = NewDeleteApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageRequest() (request *DeleteImageRequest) {
    request = &DeleteImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImage")
    return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
    response = &DeleteImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImage
// 删除指定镜像
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    response = NewDeleteImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageLifecycleGlobalPersonalRequest() (request *DeleteImageLifecycleGlobalPersonalRequest) {
    request = &DeleteImageLifecycleGlobalPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageLifecycleGlobalPersonal")
    return
}

func NewDeleteImageLifecycleGlobalPersonalResponse() (response *DeleteImageLifecycleGlobalPersonalResponse) {
    response = &DeleteImageLifecycleGlobalPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImageLifecycleGlobalPersonal
// 用于删除个人版全局镜像版本自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteImageLifecycleGlobalPersonal(request *DeleteImageLifecycleGlobalPersonalRequest) (response *DeleteImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageLifecycleGlobalPersonalRequest()
    }
    response = NewDeleteImageLifecycleGlobalPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageLifecyclePersonalRequest() (request *DeleteImageLifecyclePersonalRequest) {
    request = &DeleteImageLifecyclePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageLifecyclePersonal")
    return
}

func NewDeleteImageLifecyclePersonalResponse() (response *DeleteImageLifecyclePersonalResponse) {
    response = &DeleteImageLifecyclePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImageLifecyclePersonal
// 用于在个人版镜像仓库中删除仓库Tag自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DeleteImageLifecyclePersonal(request *DeleteImageLifecyclePersonalRequest) (response *DeleteImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageLifecyclePersonalRequest()
    }
    response = NewDeleteImageLifecyclePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImagePersonalRequest() (request *DeleteImagePersonalRequest) {
    request = &DeleteImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImagePersonal")
    return
}

func NewDeleteImagePersonalResponse() (response *DeleteImagePersonalResponse) {
    response = &DeleteImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImagePersonal
// 用于在个人版中删除tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DeleteImagePersonal(request *DeleteImagePersonalRequest) (response *DeleteImagePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImagePersonalRequest()
    }
    response = NewDeleteImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImmutableTagRulesRequest() (request *DeleteImmutableTagRulesRequest) {
    request = &DeleteImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImmutableTagRules")
    return
}

func NewDeleteImmutableTagRulesResponse() (response *DeleteImmutableTagRulesResponse) {
    response = &DeleteImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImmutableTagRules
//  删除镜像不可变规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteImmutableTagRules(request *DeleteImmutableTagRulesRequest) (response *DeleteImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewDeleteImmutableTagRulesRequest()
    }
    response = NewDeleteImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInstance")
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstance
// 删除镜像仓库企业版实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceTokenRequest() (request *DeleteInstanceTokenRequest) {
    request = &DeleteInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInstanceToken")
    return
}

func NewDeleteInstanceTokenResponse() (response *DeleteInstanceTokenResponse) {
    response = &DeleteInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstanceToken
// 删除长期访问凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceToken(request *DeleteInstanceTokenRequest) (response *DeleteInstanceTokenResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceTokenRequest()
    }
    response = NewDeleteInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInternalEndpointDnsRequest() (request *DeleteInternalEndpointDnsRequest) {
    request = &DeleteInternalEndpointDnsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInternalEndpointDns")
    return
}

func NewDeleteInternalEndpointDnsResponse() (response *DeleteInternalEndpointDnsResponse) {
    response = &DeleteInternalEndpointDnsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInternalEndpointDns
// 删除tcr内网私有域名解析
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEPRIVATEZONE = "InternalError.CreatePrivateZone"
//  INTERNALERROR_CREATEPRIVATEZONERECORD = "InternalError.CreatePrivateZoneRecord"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_DELETEPRIVATEZONERECORD = "InternalError.DeletePrivateZoneRecord"
//  INTERNALERROR_DESCRIBEINTERNALENDPOINTDNSSTATUS = "InternalError.DescribeInternalEndpointDnsStatus"
//  INTERNALERROR_DESCRIBEPRIVATEZONELIST = "InternalError.DescribePrivateZoneList"
//  INTERNALERROR_DESCRIBEPRIVATEZONERECORDLIST = "InternalError.DescribePrivateZoneRecordList"
//  INTERNALERROR_DESCRIBEPRIVATEZONESERVICELIST = "InternalError.DescribePrivateZoneServiceList"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInternalEndpointDns(request *DeleteInternalEndpointDnsRequest) (response *DeleteInternalEndpointDnsResponse, err error) {
    if request == nil {
        request = NewDeleteInternalEndpointDnsRequest()
    }
    response = NewDeleteInternalEndpointDnsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMultipleSecurityPolicyRequest() (request *DeleteMultipleSecurityPolicyRequest) {
    request = &DeleteMultipleSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteMultipleSecurityPolicy")
    return
}

func NewDeleteMultipleSecurityPolicyResponse() (response *DeleteMultipleSecurityPolicyResponse) {
    response = &DeleteMultipleSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMultipleSecurityPolicy
// 用于删除实例多个公网访问白名单策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteMultipleSecurityPolicy(request *DeleteMultipleSecurityPolicyRequest) (response *DeleteMultipleSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteMultipleSecurityPolicyRequest()
    }
    response = NewDeleteMultipleSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteNamespace")
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespacePersonalRequest() (request *DeleteNamespacePersonalRequest) {
    request = &DeleteNamespacePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteNamespacePersonal")
    return
}

func NewDeleteNamespacePersonalResponse() (response *DeleteNamespacePersonalResponse) {
    response = &DeleteNamespacePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNamespacePersonal
// 删除共享版命名空间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNONAMESPACE = "ResourceNotFound.ErrNoNamespace"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DeleteNamespacePersonal(request *DeleteNamespacePersonalRequest) (response *DeleteNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteNamespacePersonalRequest()
    }
    response = NewDeleteNamespacePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
    request = &DeleteRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteRepository")
    return
}

func NewDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
    response = &DeleteRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRepository
// 删除镜像仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryRequest()
    }
    response = NewDeleteRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryPersonalRequest() (request *DeleteRepositoryPersonalRequest) {
    request = &DeleteRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteRepositoryPersonal")
    return
}

func NewDeleteRepositoryPersonalResponse() (response *DeleteRepositoryPersonalResponse) {
    response = &DeleteRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRepositoryPersonal
// 用于个人版镜像仓库中删除
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DeleteRepositoryPersonal(request *DeleteRepositoryPersonalRequest) (response *DeleteRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryPersonalRequest()
    }
    response = NewDeleteRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityPolicyRequest() (request *DeleteSecurityPolicyRequest) {
    request = &DeleteSecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteSecurityPolicy")
    return
}

func NewDeleteSecurityPolicyResponse() (response *DeleteSecurityPolicyResponse) {
    response = &DeleteSecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSecurityPolicy
// 删除实例公网访问白名单策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTagRetentionRuleRequest() (request *DeleteTagRetentionRuleRequest) {
    request = &DeleteTagRetentionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteTagRetentionRule")
    return
}

func NewDeleteTagRetentionRuleResponse() (response *DeleteTagRetentionRuleResponse) {
    response = &DeleteTagRetentionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTagRetentionRule
// 删除版本保留规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteTagRetentionRule(request *DeleteTagRetentionRuleRequest) (response *DeleteTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewDeleteTagRetentionRuleRequest()
    }
    response = NewDeleteTagRetentionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWebhookTriggerRequest() (request *DeleteWebhookTriggerRequest) {
    request = &DeleteWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteWebhookTrigger")
    return
}

func NewDeleteWebhookTriggerResponse() (response *DeleteWebhookTriggerResponse) {
    response = &DeleteWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWebhookTrigger
// 删除触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DeleteWebhookTrigger(request *DeleteWebhookTriggerRequest) (response *DeleteWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteWebhookTriggerRequest()
    }
    response = NewDeleteWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationTriggerLogPersonalRequest() (request *DescribeApplicationTriggerLogPersonalRequest) {
    request = &DescribeApplicationTriggerLogPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeApplicationTriggerLogPersonal")
    return
}

func NewDescribeApplicationTriggerLogPersonalResponse() (response *DescribeApplicationTriggerLogPersonalResponse) {
    response = &DescribeApplicationTriggerLogPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationTriggerLogPersonal
// 用于查询应用更新触发器触发日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeApplicationTriggerLogPersonal(request *DescribeApplicationTriggerLogPersonalRequest) (response *DescribeApplicationTriggerLogPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationTriggerLogPersonalRequest()
    }
    response = NewDescribeApplicationTriggerLogPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationTriggerPersonalRequest() (request *DescribeApplicationTriggerPersonalRequest) {
    request = &DescribeApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeApplicationTriggerPersonal")
    return
}

func NewDescribeApplicationTriggerPersonalResponse() (response *DescribeApplicationTriggerPersonalResponse) {
    response = &DescribeApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationTriggerPersonal
// 用于查询应用更新触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DescribeApplicationTriggerPersonal(request *DescribeApplicationTriggerPersonalRequest) (response *DescribeApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationTriggerPersonalRequest()
    }
    response = NewDescribeApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChartDownloadInfoRequest() (request *DescribeChartDownloadInfoRequest) {
    request = &DescribeChartDownloadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeChartDownloadInfo")
    return
}

func NewDescribeChartDownloadInfoResponse() (response *DescribeChartDownloadInfoResponse) {
    response = &DescribeChartDownloadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChartDownloadInfo
// 用于在企业版中返回Chart的下载信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChartDownloadInfo(request *DescribeChartDownloadInfoRequest) (response *DescribeChartDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeChartDownloadInfoRequest()
    }
    response = NewDescribeChartDownloadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalEndpointStatusRequest() (request *DescribeExternalEndpointStatusRequest) {
    request = &DescribeExternalEndpointStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeExternalEndpointStatus")
    return
}

func NewDescribeExternalEndpointStatusResponse() (response *DescribeExternalEndpointStatusResponse) {
    response = &DescribeExternalEndpointStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeExternalEndpointStatus
// 查询实例公网访问入口状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExternalEndpointStatus(request *DescribeExternalEndpointStatusRequest) (response *DescribeExternalEndpointStatusResponse, err error) {
    if request == nil {
        request = NewDescribeExternalEndpointStatusRequest()
    }
    response = NewDescribeExternalEndpointStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFavorRepositoryPersonalRequest() (request *DescribeFavorRepositoryPersonalRequest) {
    request = &DescribeFavorRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeFavorRepositoryPersonal")
    return
}

func NewDescribeFavorRepositoryPersonalResponse() (response *DescribeFavorRepositoryPersonalResponse) {
    response = &DescribeFavorRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFavorRepositoryPersonal
// 查询个人收藏仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFavorRepositoryPersonal(request *DescribeFavorRepositoryPersonalRequest) (response *DescribeFavorRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeFavorRepositoryPersonalRequest()
    }
    response = NewDescribeFavorRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageFilterPersonalRequest() (request *DescribeImageFilterPersonalRequest) {
    request = &DescribeImageFilterPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageFilterPersonal")
    return
}

func NewDescribeImageFilterPersonalResponse() (response *DescribeImageFilterPersonalResponse) {
    response = &DescribeImageFilterPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageFilterPersonal
// 用于在个人版中查询与指定tag镜像内容相同的tag列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeImageFilterPersonal(request *DescribeImageFilterPersonalRequest) (response *DescribeImageFilterPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageFilterPersonalRequest()
    }
    response = NewDescribeImageFilterPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageLifecycleGlobalPersonalRequest() (request *DescribeImageLifecycleGlobalPersonalRequest) {
    request = &DescribeImageLifecycleGlobalPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageLifecycleGlobalPersonal")
    return
}

func NewDescribeImageLifecycleGlobalPersonalResponse() (response *DescribeImageLifecycleGlobalPersonalResponse) {
    response = &DescribeImageLifecycleGlobalPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageLifecycleGlobalPersonal
// 用于获取个人版全局镜像版本自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeImageLifecycleGlobalPersonal(request *DescribeImageLifecycleGlobalPersonalRequest) (response *DescribeImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageLifecycleGlobalPersonalRequest()
    }
    response = NewDescribeImageLifecycleGlobalPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageLifecyclePersonalRequest() (request *DescribeImageLifecyclePersonalRequest) {
    request = &DescribeImageLifecyclePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageLifecyclePersonal")
    return
}

func NewDescribeImageLifecyclePersonalResponse() (response *DescribeImageLifecyclePersonalResponse) {
    response = &DescribeImageLifecyclePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageLifecyclePersonal
// 用于获取个人版仓库中自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeImageLifecyclePersonal(request *DescribeImageLifecyclePersonalRequest) (response *DescribeImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageLifecyclePersonalRequest()
    }
    response = NewDescribeImageLifecyclePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageManifestsRequest() (request *DescribeImageManifestsRequest) {
    request = &DescribeImageManifestsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageManifests")
    return
}

func NewDescribeImageManifestsResponse() (response *DescribeImageManifestsResponse) {
    response = &DescribeImageManifestsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageManifests
// 查询容器镜像Manifest信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImageManifests(request *DescribeImageManifestsRequest) (response *DescribeImageManifestsResponse, err error) {
    if request == nil {
        request = NewDescribeImageManifestsRequest()
    }
    response = NewDescribeImageManifestsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagePersonalRequest() (request *DescribeImagePersonalRequest) {
    request = &DescribeImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImagePersonal")
    return
}

func NewDescribeImagePersonalResponse() (response *DescribeImagePersonalResponse) {
    response = &DescribeImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImagePersonal
// 用于获取个人版镜像仓库tag列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeImagePersonal(request *DescribeImagePersonalRequest) (response *DescribeImagePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImagePersonalRequest()
    }
    response = NewDescribeImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImages")
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImages
// 查询镜像版本列表或指定容器镜像信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImmutableTagRulesRequest() (request *DescribeImmutableTagRulesRequest) {
    request = &DescribeImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImmutableTagRules")
    return
}

func NewDescribeImmutableTagRulesResponse() (response *DescribeImmutableTagRulesResponse) {
    response = &DescribeImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImmutableTagRules
// 列出镜像不可变规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImmutableTagRules(request *DescribeImmutableTagRulesRequest) (response *DescribeImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewDescribeImmutableTagRulesRequest()
    }
    response = NewDescribeImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceStatusRequest() (request *DescribeInstanceStatusRequest) {
    request = &DescribeInstanceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceStatus")
    return
}

func NewDescribeInstanceStatusResponse() (response *DescribeInstanceStatusResponse) {
    response = &DescribeInstanceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceStatus
// 查询实例当前状态以及过程信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceStatus(request *DescribeInstanceStatusRequest) (response *DescribeInstanceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceStatusRequest()
    }
    response = NewDescribeInstanceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTokenRequest() (request *DescribeInstanceTokenRequest) {
    request = &DescribeInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceToken")
    return
}

func NewDescribeInstanceTokenResponse() (response *DescribeInstanceTokenResponse) {
    response = &DescribeInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceToken
// 查询长期访问凭证信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceToken(request *DescribeInstanceTokenRequest) (response *DescribeInstanceTokenResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTokenRequest()
    }
    response = NewDescribeInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 查询实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternalEndpointDnsStatusRequest() (request *DescribeInternalEndpointDnsStatusRequest) {
    request = &DescribeInternalEndpointDnsStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInternalEndpointDnsStatus")
    return
}

func NewDescribeInternalEndpointDnsStatusResponse() (response *DescribeInternalEndpointDnsStatusResponse) {
    response = &DescribeInternalEndpointDnsStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInternalEndpointDnsStatus
// 批量查询vpc是否已经添加私有域名解析
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEPRIVATEZONE = "InternalError.CreatePrivateZone"
//  INTERNALERROR_CREATEPRIVATEZONERECORD = "InternalError.CreatePrivateZoneRecord"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_DELETEPRIVATEZONERECORD = "InternalError.DeletePrivateZoneRecord"
//  INTERNALERROR_DESCRIBEINTERNALENDPOINTDNSSTATUS = "InternalError.DescribeInternalEndpointDnsStatus"
//  INTERNALERROR_DESCRIBEPRIVATEZONELIST = "InternalError.DescribePrivateZoneList"
//  INTERNALERROR_DESCRIBEPRIVATEZONERECORDLIST = "InternalError.DescribePrivateZoneRecordList"
//  INTERNALERROR_DESCRIBEPRIVATEZONESERVICELIST = "InternalError.DescribePrivateZoneServiceList"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternalEndpointDnsStatus(request *DescribeInternalEndpointDnsStatusRequest) (response *DescribeInternalEndpointDnsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInternalEndpointDnsStatusRequest()
    }
    response = NewDescribeInternalEndpointDnsStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternalEndpointsRequest() (request *DescribeInternalEndpointsRequest) {
    request = &DescribeInternalEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInternalEndpoints")
    return
}

func NewDescribeInternalEndpointsResponse() (response *DescribeInternalEndpointsResponse) {
    response = &DescribeInternalEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInternalEndpoints
// 查询实例内网访问VPC链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInternalEndpoints(request *DescribeInternalEndpointsRequest) (response *DescribeInternalEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeInternalEndpointsRequest()
    }
    response = NewDescribeInternalEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacePersonalRequest() (request *DescribeNamespacePersonalRequest) {
    request = &DescribeNamespacePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeNamespacePersonal")
    return
}

func NewDescribeNamespacePersonalResponse() (response *DescribeNamespacePersonalResponse) {
    response = &DescribeNamespacePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespacePersonal
// 查询个人版命名空间信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DescribeNamespacePersonal(request *DescribeNamespacePersonalRequest) (response *DescribeNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacePersonalRequest()
    }
    response = NewDescribeNamespacePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
    request = &DescribeNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeNamespaces")
    return
}

func NewDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
    response = &DescribeNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNamespaces
// 查询命名空间列表或指定命名空间信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    response = NewDescribeNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationInstanceCreateTasksRequest() (request *DescribeReplicationInstanceCreateTasksRequest) {
    request = &DescribeReplicationInstanceCreateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeReplicationInstanceCreateTasks")
    return
}

func NewDescribeReplicationInstanceCreateTasksResponse() (response *DescribeReplicationInstanceCreateTasksResponse) {
    response = &DescribeReplicationInstanceCreateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReplicationInstanceCreateTasks
// 查询创建从实例任务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstanceCreateTasks(request *DescribeReplicationInstanceCreateTasksRequest) (response *DescribeReplicationInstanceCreateTasksResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstanceCreateTasksRequest()
    }
    response = NewDescribeReplicationInstanceCreateTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationInstanceSyncStatusRequest() (request *DescribeReplicationInstanceSyncStatusRequest) {
    request = &DescribeReplicationInstanceSyncStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeReplicationInstanceSyncStatus")
    return
}

func NewDescribeReplicationInstanceSyncStatusResponse() (response *DescribeReplicationInstanceSyncStatusResponse) {
    response = &DescribeReplicationInstanceSyncStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReplicationInstanceSyncStatus
// 查询从实例同步状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstanceSyncStatus(request *DescribeReplicationInstanceSyncStatusRequest) (response *DescribeReplicationInstanceSyncStatusResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstanceSyncStatusRequest()
    }
    response = NewDescribeReplicationInstanceSyncStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationInstancesRequest() (request *DescribeReplicationInstancesRequest) {
    request = &DescribeReplicationInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeReplicationInstances")
    return
}

func NewDescribeReplicationInstancesResponse() (response *DescribeReplicationInstancesResponse) {
    response = &DescribeReplicationInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReplicationInstances
// 查询从实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeReplicationInstances(request *DescribeReplicationInstancesRequest) (response *DescribeReplicationInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstancesRequest()
    }
    response = NewDescribeReplicationInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoriesRequest() (request *DescribeRepositoriesRequest) {
    request = &DescribeRepositoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositories")
    return
}

func NewDescribeRepositoriesResponse() (response *DescribeRepositoriesResponse) {
    response = &DescribeRepositoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRepositories
// 查询镜像仓库列表或指定镜像仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRepositories(request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoriesRequest()
    }
    response = NewDescribeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryFilterPersonalRequest() (request *DescribeRepositoryFilterPersonalRequest) {
    request = &DescribeRepositoryFilterPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositoryFilterPersonal")
    return
}

func NewDescribeRepositoryFilterPersonalResponse() (response *DescribeRepositoryFilterPersonalResponse) {
    response = &DescribeRepositoryFilterPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRepositoryFilterPersonal
// 用于在个人版镜像仓库中，获取满足输入搜索条件的用户镜像仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DescribeRepositoryFilterPersonal(request *DescribeRepositoryFilterPersonalRequest) (response *DescribeRepositoryFilterPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryFilterPersonalRequest()
    }
    response = NewDescribeRepositoryFilterPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryOwnerPersonalRequest() (request *DescribeRepositoryOwnerPersonalRequest) {
    request = &DescribeRepositoryOwnerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositoryOwnerPersonal")
    return
}

func NewDescribeRepositoryOwnerPersonalResponse() (response *DescribeRepositoryOwnerPersonalResponse) {
    response = &DescribeRepositoryOwnerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRepositoryOwnerPersonal
// 用于在个人版中获取用户全部的镜像仓库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DescribeRepositoryOwnerPersonal(request *DescribeRepositoryOwnerPersonalRequest) (response *DescribeRepositoryOwnerPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryOwnerPersonalRequest()
    }
    response = NewDescribeRepositoryOwnerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepositoryPersonalRequest() (request *DescribeRepositoryPersonalRequest) {
    request = &DescribeRepositoryPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRepositoryPersonal")
    return
}

func NewDescribeRepositoryPersonalResponse() (response *DescribeRepositoryPersonalResponse) {
    response = &DescribeRepositoryPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRepositoryPersonal
// 查询个人版仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeRepositoryPersonal(request *DescribeRepositoryPersonalRequest) (response *DescribeRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryPersonalRequest()
    }
    response = NewDescribeRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityPoliciesRequest() (request *DescribeSecurityPoliciesRequest) {
    request = &DescribeSecurityPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeSecurityPolicies")
    return
}

func NewDescribeSecurityPoliciesResponse() (response *DescribeSecurityPoliciesResponse) {
    response = &DescribeSecurityPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSecurityPolicies
// 查询实例公网访问白名单策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicies(request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPoliciesRequest()
    }
    response = NewDescribeSecurityPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagRetentionExecutionRequest() (request *DescribeTagRetentionExecutionRequest) {
    request = &DescribeTagRetentionExecutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeTagRetentionExecution")
    return
}

func NewDescribeTagRetentionExecutionResponse() (response *DescribeTagRetentionExecutionResponse) {
    response = &DescribeTagRetentionExecutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTagRetentionExecution
// 查询版本保留执行记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionExecution(request *DescribeTagRetentionExecutionRequest) (response *DescribeTagRetentionExecutionResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionExecutionRequest()
    }
    response = NewDescribeTagRetentionExecutionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagRetentionExecutionTaskRequest() (request *DescribeTagRetentionExecutionTaskRequest) {
    request = &DescribeTagRetentionExecutionTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeTagRetentionExecutionTask")
    return
}

func NewDescribeTagRetentionExecutionTaskResponse() (response *DescribeTagRetentionExecutionTaskResponse) {
    response = &DescribeTagRetentionExecutionTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTagRetentionExecutionTask
// 查询版本保留执行任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionExecutionTask(request *DescribeTagRetentionExecutionTaskRequest) (response *DescribeTagRetentionExecutionTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionExecutionTaskRequest()
    }
    response = NewDescribeTagRetentionExecutionTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagRetentionRulesRequest() (request *DescribeTagRetentionRulesRequest) {
    request = &DescribeTagRetentionRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeTagRetentionRules")
    return
}

func NewDescribeTagRetentionRulesResponse() (response *DescribeTagRetentionRulesResponse) {
    response = &DescribeTagRetentionRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTagRetentionRules
// 查询版本保留规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTagRetentionRules(request *DescribeTagRetentionRulesRequest) (response *DescribeTagRetentionRulesResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionRulesRequest()
    }
    response = NewDescribeTagRetentionRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserQuotaPersonalRequest() (request *DescribeUserQuotaPersonalRequest) {
    request = &DescribeUserQuotaPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeUserQuotaPersonal")
    return
}

func NewDescribeUserQuotaPersonalResponse() (response *DescribeUserQuotaPersonalResponse) {
    response = &DescribeUserQuotaPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserQuotaPersonal
// 查询个人用户配额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUserQuotaPersonal(request *DescribeUserQuotaPersonalRequest) (response *DescribeUserQuotaPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotaPersonalRequest()
    }
    response = NewDescribeUserQuotaPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebhookTriggerRequest() (request *DescribeWebhookTriggerRequest) {
    request = &DescribeWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeWebhookTrigger")
    return
}

func NewDescribeWebhookTriggerResponse() (response *DescribeWebhookTriggerResponse) {
    response = &DescribeWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebhookTrigger
// 查询触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTrigger(request *DescribeWebhookTriggerRequest) (response *DescribeWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewDescribeWebhookTriggerRequest()
    }
    response = NewDescribeWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebhookTriggerLogRequest() (request *DescribeWebhookTriggerLogRequest) {
    request = &DescribeWebhookTriggerLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeWebhookTriggerLog")
    return
}

func NewDescribeWebhookTriggerLogResponse() (response *DescribeWebhookTriggerLogResponse) {
    response = &DescribeWebhookTriggerLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWebhookTriggerLog
// 获取触发器日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTriggerLog(request *DescribeWebhookTriggerLogRequest) (response *DescribeWebhookTriggerLogResponse, err error) {
    if request == nil {
        request = NewDescribeWebhookTriggerLogRequest()
    }
    response = NewDescribeWebhookTriggerLogResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadHelmChartRequest() (request *DownloadHelmChartRequest) {
    request = &DownloadHelmChartRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DownloadHelmChart")
    return
}

func NewDownloadHelmChartResponse() (response *DownloadHelmChartResponse) {
    response = &DownloadHelmChartResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadHelmChart
// 用于在TCR中下载helm chart
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DownloadHelmChart(request *DownloadHelmChartRequest) (response *DownloadHelmChartResponse, err error) {
    if request == nil {
        request = NewDownloadHelmChartRequest()
    }
    response = NewDownloadHelmChartResponse()
    err = c.Send(request, response)
    return
}

func NewDuplicateImagePersonalRequest() (request *DuplicateImagePersonalRequest) {
    request = &DuplicateImagePersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "DuplicateImagePersonal")
    return
}

func NewDuplicateImagePersonalResponse() (response *DuplicateImagePersonalResponse) {
    response = &DuplicateImagePersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DuplicateImagePersonal
// 用于在个人版镜像仓库中复制镜像版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOTAG = "ResourceNotFound.ErrNoTag"
func (c *Client) DuplicateImagePersonal(request *DuplicateImagePersonalRequest) (response *DuplicateImagePersonalResponse, err error) {
    if request == nil {
        request = NewDuplicateImagePersonalRequest()
    }
    response = NewDuplicateImagePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewManageExternalEndpointRequest() (request *ManageExternalEndpointRequest) {
    request = &ManageExternalEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageExternalEndpoint")
    return
}

func NewManageExternalEndpointResponse() (response *ManageExternalEndpointResponse) {
    response = &ManageExternalEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageExternalEndpoint
// 管理实例公网访问
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageExternalEndpoint(request *ManageExternalEndpointRequest) (response *ManageExternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageExternalEndpointRequest()
    }
    response = NewManageExternalEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewManageImageLifecycleGlobalPersonalRequest() (request *ManageImageLifecycleGlobalPersonalRequest) {
    request = &ManageImageLifecycleGlobalPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageImageLifecycleGlobalPersonal")
    return
}

func NewManageImageLifecycleGlobalPersonalResponse() (response *ManageImageLifecycleGlobalPersonalResponse) {
    response = &ManageImageLifecycleGlobalPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageImageLifecycleGlobalPersonal
// 用于设置个人版全局镜像版本自动清理策略
//
// 可能返回的错误码:
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ManageImageLifecycleGlobalPersonal(request *ManageImageLifecycleGlobalPersonalRequest) (response *ManageImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewManageImageLifecycleGlobalPersonalRequest()
    }
    response = NewManageImageLifecycleGlobalPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewManageInternalEndpointRequest() (request *ManageInternalEndpointRequest) {
    request = &ManageInternalEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageInternalEndpoint")
    return
}

func NewManageInternalEndpointResponse() (response *ManageInternalEndpointResponse) {
    response = &ManageInternalEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageInternalEndpoint
// 管理实例内网访问VPC链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORVPCDNSSTATUS = "ResourceInsufficient.ErrorVpcDnsStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageInternalEndpoint(request *ManageInternalEndpointRequest) (response *ManageInternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageInternalEndpointRequest()
    }
    response = NewManageInternalEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewManageReplicationRequest() (request *ManageReplicationRequest) {
    request = &ManageReplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ManageReplication")
    return
}

func NewManageReplicationResponse() (response *ManageReplicationResponse) {
    response = &ManageReplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageReplication
// 管理实例同步
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ManageReplication(request *ManageReplicationRequest) (response *ManageReplicationResponse, err error) {
    if request == nil {
        request = NewManageReplicationRequest()
    }
    response = NewManageReplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationTriggerPersonalRequest() (request *ModifyApplicationTriggerPersonalRequest) {
    request = &ModifyApplicationTriggerPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyApplicationTriggerPersonal")
    return
}

func NewModifyApplicationTriggerPersonalResponse() (response *ModifyApplicationTriggerPersonalResponse) {
    response = &ModifyApplicationTriggerPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationTriggerPersonal
// 用于修改应用更新触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
//  RESOURCENOTFOUND_ERRNOTRIGGER = "ResourceNotFound.ErrNoTrigger"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) ModifyApplicationTriggerPersonal(request *ModifyApplicationTriggerPersonalRequest) (response *ModifyApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewModifyApplicationTriggerPersonalRequest()
    }
    response = NewModifyApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImmutableTagRulesRequest() (request *ModifyImmutableTagRulesRequest) {
    request = &ModifyImmutableTagRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyImmutableTagRules")
    return
}

func NewModifyImmutableTagRulesResponse() (response *ModifyImmutableTagRulesResponse) {
    response = &ModifyImmutableTagRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyImmutableTagRules
// 更新镜像不可变规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
func (c *Client) ModifyImmutableTagRules(request *ModifyImmutableTagRulesRequest) (response *ModifyImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewModifyImmutableTagRulesRequest()
    }
    response = NewModifyImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyInstance")
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstance
// 更新实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceTokenRequest() (request *ModifyInstanceTokenRequest) {
    request = &ModifyInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyInstanceToken")
    return
}

func NewModifyInstanceTokenResponse() (response *ModifyInstanceTokenResponse) {
    response = &ModifyInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceToken
// 更新实例内指定长期访问凭证的启用状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceToken(request *ModifyInstanceTokenRequest) (response *ModifyInstanceTokenResponse, err error) {
    if request == nil {
        request = NewModifyInstanceTokenRequest()
    }
    response = NewModifyInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
    request = &ModifyNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyNamespace")
    return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
    response = &ModifyNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNamespace
// 更新命名空间信息，当前仅支持修改命名空间访问级别
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    response = NewModifyNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRepositoryRequest() (request *ModifyRepositoryRequest) {
    request = &ModifyRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyRepository")
    return
}

func NewModifyRepositoryResponse() (response *ModifyRepositoryResponse) {
    response = &ModifyRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRepository
// 更新镜像仓库信息，可修改仓库描述信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRepository(request *ModifyRepositoryRequest) (response *ModifyRepositoryResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryRequest()
    }
    response = NewModifyRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRepositoryAccessPersonalRequest() (request *ModifyRepositoryAccessPersonalRequest) {
    request = &ModifyRepositoryAccessPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyRepositoryAccessPersonal")
    return
}

func NewModifyRepositoryAccessPersonalResponse() (response *ModifyRepositoryAccessPersonalResponse) {
    response = &ModifyRepositoryAccessPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRepositoryAccessPersonal
// 用于更新个人版镜像仓库的访问属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) ModifyRepositoryAccessPersonal(request *ModifyRepositoryAccessPersonalRequest) (response *ModifyRepositoryAccessPersonalResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryAccessPersonalRequest()
    }
    response = NewModifyRepositoryAccessPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRepositoryInfoPersonalRequest() (request *ModifyRepositoryInfoPersonalRequest) {
    request = &ModifyRepositoryInfoPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyRepositoryInfoPersonal")
    return
}

func NewModifyRepositoryInfoPersonalResponse() (response *ModifyRepositoryInfoPersonalResponse) {
    response = &ModifyRepositoryInfoPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRepositoryInfoPersonal
// 用于在个人版镜像仓库中更新容器镜像描述
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) ModifyRepositoryInfoPersonal(request *ModifyRepositoryInfoPersonalRequest) (response *ModifyRepositoryInfoPersonalResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryInfoPersonalRequest()
    }
    response = NewModifyRepositoryInfoPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyRequest() (request *ModifySecurityPolicyRequest) {
    request = &ModifySecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifySecurityPolicy")
    return
}

func NewModifySecurityPolicyResponse() (response *ModifySecurityPolicyResponse) {
    response = &ModifySecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySecurityPolicy
// 更新实例公网访问白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyRequest()
    }
    response = NewModifySecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTagRetentionRuleRequest() (request *ModifyTagRetentionRuleRequest) {
    request = &ModifyTagRetentionRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyTagRetentionRule")
    return
}

func NewModifyTagRetentionRuleResponse() (response *ModifyTagRetentionRuleResponse) {
    response = &ModifyTagRetentionRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTagRetentionRule
// 更新版本保留规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTagRetentionRule(request *ModifyTagRetentionRuleRequest) (response *ModifyTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewModifyTagRetentionRuleRequest()
    }
    response = NewModifyTagRetentionRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserPasswordPersonalRequest() (request *ModifyUserPasswordPersonalRequest) {
    request = &ModifyUserPasswordPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyUserPasswordPersonal")
    return
}

func NewModifyUserPasswordPersonalResponse() (response *ModifyUserPasswordPersonalResponse) {
    response = &ModifyUserPasswordPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUserPasswordPersonal
// 修改个人用户登录密码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) ModifyUserPasswordPersonal(request *ModifyUserPasswordPersonalRequest) (response *ModifyUserPasswordPersonalResponse, err error) {
    if request == nil {
        request = NewModifyUserPasswordPersonalRequest()
    }
    response = NewModifyUserPasswordPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWebhookTriggerRequest() (request *ModifyWebhookTriggerRequest) {
    request = &ModifyWebhookTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ModifyWebhookTrigger")
    return
}

func NewModifyWebhookTriggerResponse() (response *ModifyWebhookTriggerResponse) {
    response = &ModifyWebhookTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyWebhookTrigger
// 更新触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) ModifyWebhookTrigger(request *ModifyWebhookTriggerRequest) (response *ModifyWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewModifyWebhookTriggerRequest()
    }
    response = NewModifyWebhookTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "RenewInstance")
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewInstance
// 预付费实例续费，同时支持按量计费转包年包月
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORNAMEEXISTS = "InvalidParameter.ErrorNameExists"
//  INVALIDPARAMETER_ERRORREGISTRYNAME = "InvalidParameter.ErrorRegistryName"
//  INVALIDPARAMETER_ERRORTAGOVERLIMIT = "InvalidParameter.ErrorTagOverLimit"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewValidateNamespaceExistPersonalRequest() (request *ValidateNamespaceExistPersonalRequest) {
    request = &ValidateNamespaceExistPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ValidateNamespaceExistPersonal")
    return
}

func NewValidateNamespaceExistPersonalResponse() (response *ValidateNamespaceExistPersonalResponse) {
    response = &ValidateNamespaceExistPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ValidateNamespaceExistPersonal
// 查询个人版用户命名空间是否存在
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ValidateNamespaceExistPersonal(request *ValidateNamespaceExistPersonalRequest) (response *ValidateNamespaceExistPersonalResponse, err error) {
    if request == nil {
        request = NewValidateNamespaceExistPersonalRequest()
    }
    response = NewValidateNamespaceExistPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewValidateRepositoryExistPersonalRequest() (request *ValidateRepositoryExistPersonalRequest) {
    request = &ValidateRepositoryExistPersonalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcr", APIVersion, "ValidateRepositoryExistPersonal")
    return
}

func NewValidateRepositoryExistPersonalResponse() (response *ValidateRepositoryExistPersonalResponse) {
    response = &ValidateRepositoryExistPersonalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ValidateRepositoryExistPersonal
// 用于判断个人版仓库是否存在
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ValidateRepositoryExistPersonal(request *ValidateRepositoryExistPersonalRequest) (response *ValidateRepositoryExistPersonalResponse, err error) {
    if request == nil {
        request = NewValidateRepositoryExistPersonalRequest()
    }
    response = NewValidateRepositoryExistPersonalResponse()
    err = c.Send(request, response)
    return
}
