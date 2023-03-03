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
    "context"
    "errors"
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
    return c.BatchDeleteImagePersonalWithContext(context.Background(), request)
}

// BatchDeleteImagePersonal
// 用于在个人版镜像仓库中批量删除Tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) BatchDeleteImagePersonalWithContext(ctx context.Context, request *BatchDeleteImagePersonalRequest) (response *BatchDeleteImagePersonalResponse, err error) {
    if request == nil {
        request = NewBatchDeleteImagePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteImagePersonal require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_ERRTOOLARGE = "InvalidParameter.ErrTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) BatchDeleteRepositoryPersonal(request *BatchDeleteRepositoryPersonalRequest) (response *BatchDeleteRepositoryPersonalResponse, err error) {
    return c.BatchDeleteRepositoryPersonalWithContext(context.Background(), request)
}

// BatchDeleteRepositoryPersonal
// 用于个人版镜像仓库中批量删除镜像仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  INVALIDPARAMETER_ERRTOOLARGE = "InvalidParameter.ErrTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) BatchDeleteRepositoryPersonalWithContext(ctx context.Context, request *BatchDeleteRepositoryPersonalRequest) (response *BatchDeleteRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewBatchDeleteRepositoryPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteRepositoryPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CheckInstanceWithContext(context.Background(), request)
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
func (c *Client) CheckInstanceWithContext(ctx context.Context, request *CheckInstanceRequest) (response *CheckInstanceResponse, err error) {
    if request == nil {
        request = NewCheckInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckInstance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CheckInstanceNameWithContext(context.Background(), request)
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
func (c *Client) CheckInstanceNameWithContext(ctx context.Context, request *CheckInstanceNameRequest) (response *CheckInstanceNameResponse, err error) {
    if request == nil {
        request = NewCheckInstanceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckInstanceName require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateApplicationTriggerPersonalWithContext(context.Background(), request)
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
func (c *Client) CreateApplicationTriggerPersonalWithContext(ctx context.Context, request *CreateApplicationTriggerPersonalRequest) (response *CreateApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewCreateApplicationTriggerPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationTriggerPersonal require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationTriggerPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageAccelerationServiceRequest() (request *CreateImageAccelerationServiceRequest) {
    request = &CreateImageAccelerationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateImageAccelerationService")
    
    
    return
}

func NewCreateImageAccelerationServiceResponse() (response *CreateImageAccelerationServiceResponse) {
    response = &CreateImageAccelerationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImageAccelerationService
// 创建镜像加速服务
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
func (c *Client) CreateImageAccelerationService(request *CreateImageAccelerationServiceRequest) (response *CreateImageAccelerationServiceResponse, err error) {
    return c.CreateImageAccelerationServiceWithContext(context.Background(), request)
}

// CreateImageAccelerationService
// 创建镜像加速服务
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
func (c *Client) CreateImageAccelerationServiceWithContext(ctx context.Context, request *CreateImageAccelerationServiceRequest) (response *CreateImageAccelerationServiceResponse, err error) {
    if request == nil {
        request = NewCreateImageAccelerationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageAccelerationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageAccelerationServiceResponse()
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
    return c.CreateImageLifecyclePersonalWithContext(context.Background(), request)
}

// CreateImageLifecyclePersonal
// 用于在个人版中创建清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) CreateImageLifecyclePersonalWithContext(ctx context.Context, request *CreateImageLifecyclePersonalRequest) (response *CreateImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewCreateImageLifecyclePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageLifecyclePersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateImmutableTagRulesWithContext(context.Background(), request)
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
func (c *Client) CreateImmutableTagRulesWithContext(ctx context.Context, request *CreateImmutableTagRulesRequest) (response *CreateImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewCreateImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_TRADEFAILED = "FailedOperation.TradeFailed"
//  FAILEDOPERATION_VALIDATEREGISTRYNAMEFAIL = "FailedOperation.ValidateRegistryNameFail"
//  FAILEDOPERATION_VALIDATESUPPORTEDREGIONFAIL = "FailedOperation.ValidateSupportedRegionFail"
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
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 创建实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_TRADEFAILED = "FailedOperation.TradeFailed"
//  FAILEDOPERATION_VALIDATEREGISTRYNAMEFAIL = "FailedOperation.ValidateRegistryNameFail"
//  FAILEDOPERATION_VALIDATESUPPORTEDREGIONFAIL = "FailedOperation.ValidateSupportedRegionFail"
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
//  OPERATIONDENIED_QUOTAOVERLIMIT = "OperationDenied.QuotaOverLimit"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceCustomizedDomainRequest() (request *CreateInstanceCustomizedDomainRequest) {
    request = &CreateInstanceCustomizedDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateInstanceCustomizedDomain")
    
    
    return
}

func NewCreateInstanceCustomizedDomainResponse() (response *CreateInstanceCustomizedDomainResponse) {
    response = &CreateInstanceCustomizedDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceCustomizedDomain
// 创建自定义域名
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
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceCustomizedDomain(request *CreateInstanceCustomizedDomainRequest) (response *CreateInstanceCustomizedDomainResponse, err error) {
    return c.CreateInstanceCustomizedDomainWithContext(context.Background(), request)
}

// CreateInstanceCustomizedDomain
// 创建自定义域名
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
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstanceCustomizedDomainWithContext(ctx context.Context, request *CreateInstanceCustomizedDomainRequest) (response *CreateInstanceCustomizedDomainResponse, err error) {
    if request == nil {
        request = NewCreateInstanceCustomizedDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceCustomizedDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceCustomizedDomainResponse()
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
    return c.CreateInstanceTokenWithContext(context.Background(), request)
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
func (c *Client) CreateInstanceTokenWithContext(ctx context.Context, request *CreateInstanceTokenRequest) (response *CreateInstanceTokenResponse, err error) {
    if request == nil {
        request = NewCreateInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceToken require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_CREATEPRIVATEZONE = "UnsupportedOperation.CreatePrivateZone"
//  UNSUPPORTEDOPERATION_CREATEPRIVATEZONERECORD = "UnsupportedOperation.CreatePrivateZoneRecord"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONELIST = "UnsupportedOperation.DescribePrivateZoneList"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONERECORDLIST = "UnsupportedOperation.DescribePrivateZoneRecordList"
//  UNSUPPORTEDOPERATION_MODIFYPRIVATEZONERECORD = "UnsupportedOperation.ModifyPrivateZoneRecord"
//  UNSUPPORTEDOPERATION_MODIFYPRIVATEZONEVPC = "UnsupportedOperation.ModifyPrivateZoneVpc"
func (c *Client) CreateInternalEndpointDns(request *CreateInternalEndpointDnsRequest) (response *CreateInternalEndpointDnsResponse, err error) {
    return c.CreateInternalEndpointDnsWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_CREATEPRIVATEZONE = "UnsupportedOperation.CreatePrivateZone"
//  UNSUPPORTEDOPERATION_CREATEPRIVATEZONERECORD = "UnsupportedOperation.CreatePrivateZoneRecord"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONELIST = "UnsupportedOperation.DescribePrivateZoneList"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONERECORDLIST = "UnsupportedOperation.DescribePrivateZoneRecordList"
//  UNSUPPORTEDOPERATION_MODIFYPRIVATEZONERECORD = "UnsupportedOperation.ModifyPrivateZoneRecord"
//  UNSUPPORTEDOPERATION_MODIFYPRIVATEZONEVPC = "UnsupportedOperation.ModifyPrivateZoneVpc"
func (c *Client) CreateInternalEndpointDnsWithContext(ctx context.Context, request *CreateInternalEndpointDnsRequest) (response *CreateInternalEndpointDnsResponse, err error) {
    if request == nil {
        request = NewCreateInternalEndpointDnsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInternalEndpointDns require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMultipleSecurityPolicy(request *CreateMultipleSecurityPolicyRequest) (response *CreateMultipleSecurityPolicyResponse, err error) {
    return c.CreateMultipleSecurityPolicyWithContext(context.Background(), request)
}

// CreateMultipleSecurityPolicy
// 用于在TCR实例中，创建多个白名单策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMultipleSecurityPolicyWithContext(ctx context.Context, request *CreateMultipleSecurityPolicyRequest) (response *CreateMultipleSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateMultipleSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultipleSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORTCRINVALIDMEDIATYPE = "FailedOperation.ErrorTcrInvalidMediaType"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
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
    return c.CreateNamespaceWithContext(context.Background(), request)
}

// CreateNamespace
// 用于在企业版中创建命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORTCRINVALIDMEDIATYPE = "FailedOperation.ErrorTcrInvalidMediaType"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
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
func (c *Client) CreateNamespaceWithContext(ctx context.Context, request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNamespace require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNAMESPACEEXIST = "InvalidParameter.ErrNamespaceExist"
//  INVALIDPARAMETER_ERRNAMESPACERESERVED = "InvalidParameter.ErrNamespaceReserved"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
//  UNSUPPORTEDOPERATION_ERRNOUSERINITIALIZED = "UnsupportedOperation.ErrNoUserInitialized"
func (c *Client) CreateNamespacePersonal(request *CreateNamespacePersonalRequest) (response *CreateNamespacePersonalResponse, err error) {
    return c.CreateNamespacePersonalWithContext(context.Background(), request)
}

// CreateNamespacePersonal
// 创建个人版镜像仓库命名空间，此命名空间全局唯一
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNAMESPACEEXIST = "InvalidParameter.ErrNamespaceExist"
//  INVALIDPARAMETER_ERRNAMESPACERESERVED = "InvalidParameter.ErrNamespaceReserved"
//  LIMITEXCEEDED_ERRNAMESPACEMAXLIMIT = "LimitExceeded.ErrNamespaceMaxLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
//  UNSUPPORTEDOPERATION_ERRNOUSERINITIALIZED = "UnsupportedOperation.ErrNoUserInitialized"
func (c *Client) CreateNamespacePersonalWithContext(ctx context.Context, request *CreateNamespacePersonalRequest) (response *CreateNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewCreateNamespacePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNamespacePersonal require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPLICATIONEXISTS = "InvalidParameter.ReplicationExists"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReplicationInstance(request *CreateReplicationInstanceRequest) (response *CreateReplicationInstanceResponse, err error) {
    return c.CreateReplicationInstanceWithContext(context.Background(), request)
}

// CreateReplicationInstance
// 创建从实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRCONFLICT = "InternalError.ErrConflict"
//  INTERNALERROR_ERRNOTEXIST = "InternalError.ErrNotExist"
//  INTERNALERROR_ERRORCONFLICT = "InternalError.ErrorConflict"
//  INTERNALERROR_ERROROVERLIMIT = "InternalError.ErrorOverLimit"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPLICATIONEXISTS = "InvalidParameter.ReplicationExists"
//  INVALIDPARAMETER_UNSUPPORTEDREGION = "InvalidParameter.UnsupportedRegion"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateReplicationInstanceWithContext(ctx context.Context, request *CreateReplicationInstanceRequest) (response *CreateReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewCreateReplicationInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReplicationInstance require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
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
    return c.CreateRepositoryWithContext(context.Background(), request)
}

// CreateRepository
// 用于企业版创建镜像仓库
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
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
func (c *Client) CreateRepositoryWithContext(ctx context.Context, request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRepository require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateRepositoryPersonalWithContext(context.Background(), request)
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
func (c *Client) CreateRepositoryPersonalWithContext(ctx context.Context, request *CreateRepositoryPersonalRequest) (response *CreateRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewCreateRepositoryPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRepositoryPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateSecurityPolicyWithContext(context.Background(), request)
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
func (c *Client) CreateSecurityPolicyWithContext(ctx context.Context, request *CreateSecurityPolicyRequest) (response *CreateSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSignatureRequest() (request *CreateSignatureRequest) {
    request = &CreateSignatureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSignature")
    
    
    return
}

func NewCreateSignatureResponse() (response *CreateSignatureResponse) {
    response = &CreateSignatureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSignature
// 为一个镜像版本创建签名
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignature(request *CreateSignatureRequest) (response *CreateSignatureResponse, err error) {
    return c.CreateSignatureWithContext(context.Background(), request)
}

// CreateSignature
// 为一个镜像版本创建签名
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRINVALIDMEDIATYPE = "InternalError.ErrorTcrInvalidMediaType"
//  INTERNALERROR_ERRORTCRRESOURCECONFLICT = "InternalError.ErrorTcrResourceConflict"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignatureWithContext(ctx context.Context, request *CreateSignatureRequest) (response *CreateSignatureResponse, err error) {
    if request == nil {
        request = NewCreateSignatureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSignature require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSignatureResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSignaturePolicyRequest() (request *CreateSignaturePolicyRequest) {
    request = &CreateSignaturePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "CreateSignaturePolicy")
    
    
    return
}

func NewCreateSignaturePolicyResponse() (response *CreateSignaturePolicyResponse) {
    response = &CreateSignaturePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSignaturePolicy
// 创建镜像签名策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignaturePolicy(request *CreateSignaturePolicyRequest) (response *CreateSignaturePolicyResponse, err error) {
    return c.CreateSignaturePolicyWithContext(context.Background(), request)
}

// CreateSignaturePolicy
// 创建镜像签名策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignaturePolicyWithContext(ctx context.Context, request *CreateSignaturePolicyRequest) (response *CreateSignaturePolicyResponse, err error) {
    if request == nil {
        request = NewCreateSignaturePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSignaturePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSignaturePolicyResponse()
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
    return c.CreateTagRetentionExecutionWithContext(context.Background(), request)
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
func (c *Client) CreateTagRetentionExecutionWithContext(ctx context.Context, request *CreateTagRetentionExecutionRequest) (response *CreateTagRetentionExecutionResponse, err error) {
    if request == nil {
        request = NewCreateTagRetentionExecutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTagRetentionExecution require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateTagRetentionRuleWithContext(context.Background(), request)
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
func (c *Client) CreateTagRetentionRuleWithContext(ctx context.Context, request *CreateTagRetentionRuleRequest) (response *CreateTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewCreateTagRetentionRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTagRetentionRule require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateUserPersonalWithContext(context.Background(), request)
}

// CreateUserPersonal
// 创建个人用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRUSEREXIST = "InvalidParameter.ErrUserExist"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
func (c *Client) CreateUserPersonalWithContext(ctx context.Context, request *CreateUserPersonalRequest) (response *CreateUserPersonalResponse, err error) {
    if request == nil {
        request = NewCreateUserPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateWebhookTriggerWithContext(context.Background(), request)
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
func (c *Client) CreateWebhookTriggerWithContext(ctx context.Context, request *CreateWebhookTriggerRequest) (response *CreateWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewCreateWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteApplicationTriggerPersonalWithContext(context.Background(), request)
}

// DeleteApplicationTriggerPersonal
// 用于删除应用更新触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ERRNOTRIGGER = "ResourceNotFound.ErrNoTrigger"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DeleteApplicationTriggerPersonalWithContext(ctx context.Context, request *DeleteApplicationTriggerPersonalRequest) (response *DeleteApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationTriggerPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationTriggerPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteImageWithContext(context.Background(), request)
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
func (c *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageAccelerateServiceRequest() (request *DeleteImageAccelerateServiceRequest) {
    request = &DeleteImageAccelerateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteImageAccelerateService")
    
    
    return
}

func NewDeleteImageAccelerateServiceResponse() (response *DeleteImageAccelerateServiceResponse) {
    response = &DeleteImageAccelerateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImageAccelerateService
// 删除镜像加速服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteImageAccelerateService(request *DeleteImageAccelerateServiceRequest) (response *DeleteImageAccelerateServiceResponse, err error) {
    return c.DeleteImageAccelerateServiceWithContext(context.Background(), request)
}

// DeleteImageAccelerateService
// 删除镜像加速服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteImageAccelerateServiceWithContext(ctx context.Context, request *DeleteImageAccelerateServiceRequest) (response *DeleteImageAccelerateServiceResponse, err error) {
    if request == nil {
        request = NewDeleteImageAccelerateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageAccelerateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteImageAccelerateServiceResponse()
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
    return c.DeleteImageLifecycleGlobalPersonalWithContext(context.Background(), request)
}

// DeleteImageLifecycleGlobalPersonal
// 用于删除个人版全局镜像版本自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteImageLifecycleGlobalPersonalWithContext(ctx context.Context, request *DeleteImageLifecycleGlobalPersonalRequest) (response *DeleteImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageLifecycleGlobalPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageLifecycleGlobalPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteImageLifecyclePersonalWithContext(context.Background(), request)
}

// DeleteImageLifecyclePersonal
// 用于在个人版镜像仓库中删除仓库Tag自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DeleteImageLifecyclePersonalWithContext(ctx context.Context, request *DeleteImageLifecyclePersonalRequest) (response *DeleteImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImageLifecyclePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImageLifecyclePersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteImagePersonalWithContext(context.Background(), request)
}

// DeleteImagePersonal
// 用于在个人版中删除tag
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DeleteImagePersonalWithContext(ctx context.Context, request *DeleteImagePersonalRequest) (response *DeleteImagePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteImagePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImagePersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteImmutableTagRulesWithContext(context.Background(), request)
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
func (c *Client) DeleteImmutableTagRulesWithContext(ctx context.Context, request *DeleteImmutableTagRulesRequest) (response *DeleteImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewDeleteImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteInstanceWithContext(context.Background(), request)
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
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceCustomizedDomainRequest() (request *DeleteInstanceCustomizedDomainRequest) {
    request = &DeleteInstanceCustomizedDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteInstanceCustomizedDomain")
    
    
    return
}

func NewDeleteInstanceCustomizedDomainResponse() (response *DeleteInstanceCustomizedDomainResponse) {
    response = &DeleteInstanceCustomizedDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstanceCustomizedDomain
// 删除自定义域名
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceCustomizedDomain(request *DeleteInstanceCustomizedDomainRequest) (response *DeleteInstanceCustomizedDomainResponse, err error) {
    return c.DeleteInstanceCustomizedDomainWithContext(context.Background(), request)
}

// DeleteInstanceCustomizedDomain
// 删除自定义域名
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceCustomizedDomainWithContext(ctx context.Context, request *DeleteInstanceCustomizedDomainRequest) (response *DeleteInstanceCustomizedDomainResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceCustomizedDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstanceCustomizedDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceCustomizedDomainResponse()
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
    return c.DeleteInstanceTokenWithContext(context.Background(), request)
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
func (c *Client) DeleteInstanceTokenWithContext(ctx context.Context, request *DeleteInstanceTokenRequest) (response *DeleteInstanceTokenResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstanceToken require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION_DELETEPRIVATEZONERECORD = "UnsupportedOperation.DeletePrivateZoneRecord"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONELIST = "UnsupportedOperation.DescribePrivateZoneList"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONERECORDLIST = "UnsupportedOperation.DescribePrivateZoneRecordList"
func (c *Client) DeleteInternalEndpointDns(request *DeleteInternalEndpointDnsRequest) (response *DeleteInternalEndpointDnsResponse, err error) {
    return c.DeleteInternalEndpointDnsWithContext(context.Background(), request)
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
//  UNSUPPORTEDOPERATION_DELETEPRIVATEZONERECORD = "UnsupportedOperation.DeletePrivateZoneRecord"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONELIST = "UnsupportedOperation.DescribePrivateZoneList"
//  UNSUPPORTEDOPERATION_DESCRIBEPRIVATEZONERECORDLIST = "UnsupportedOperation.DescribePrivateZoneRecordList"
func (c *Client) DeleteInternalEndpointDnsWithContext(ctx context.Context, request *DeleteInternalEndpointDnsRequest) (response *DeleteInternalEndpointDnsResponse, err error) {
    if request == nil {
        request = NewDeleteInternalEndpointDnsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInternalEndpointDns require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteMultipleSecurityPolicyWithContext(context.Background(), request)
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
func (c *Client) DeleteMultipleSecurityPolicyWithContext(ctx context.Context, request *DeleteMultipleSecurityPolicyRequest) (response *DeleteMultipleSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteMultipleSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMultipleSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
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
    return c.DeleteNamespaceWithContext(context.Background(), request)
}

// DeleteNamespace
// 删除命名空间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
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
func (c *Client) DeleteNamespaceWithContext(ctx context.Context, request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNamespace require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteNamespacePersonalWithContext(context.Background(), request)
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
func (c *Client) DeleteNamespacePersonalWithContext(ctx context.Context, request *DeleteNamespacePersonalRequest) (response *DeleteNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewDeleteNamespacePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNamespacePersonal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNamespacePersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReplicationInstanceRequest() (request *DeleteReplicationInstanceRequest) {
    request = &DeleteReplicationInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteReplicationInstance")
    
    
    return
}

func NewDeleteReplicationInstanceResponse() (response *DeleteReplicationInstanceResponse) {
    response = &DeleteReplicationInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteReplicationInstance
// 删除从实例
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
func (c *Client) DeleteReplicationInstance(request *DeleteReplicationInstanceRequest) (response *DeleteReplicationInstanceResponse, err error) {
    return c.DeleteReplicationInstanceWithContext(context.Background(), request)
}

// DeleteReplicationInstance
// 删除从实例
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
func (c *Client) DeleteReplicationInstanceWithContext(ctx context.Context, request *DeleteReplicationInstanceRequest) (response *DeleteReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteReplicationInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReplicationInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReplicationInstanceResponse()
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
    return c.DeleteRepositoryWithContext(context.Background(), request)
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
func (c *Client) DeleteRepositoryWithContext(ctx context.Context, request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRepository require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteRepositoryPersonalWithContext(context.Background(), request)
}

// DeleteRepositoryPersonal
// 用于个人版镜像仓库中删除
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DeleteRepositoryPersonalWithContext(ctx context.Context, request *DeleteRepositoryPersonalRequest) (response *DeleteRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRepositoryPersonal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRepositoryTagsRequest() (request *DeleteRepositoryTagsRequest) {
    request = &DeleteRepositoryTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteRepositoryTags")
    
    
    return
}

func NewDeleteRepositoryTagsResponse() (response *DeleteRepositoryTagsResponse) {
    response = &DeleteRepositoryTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRepositoryTags
// 用于企业版批量删除Repository Tag
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
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRepositoryTags(request *DeleteRepositoryTagsRequest) (response *DeleteRepositoryTagsResponse, err error) {
    return c.DeleteRepositoryTagsWithContext(context.Background(), request)
}

// DeleteRepositoryTags
// 用于企业版批量删除Repository Tag
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
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRepositoryTagsWithContext(ctx context.Context, request *DeleteRepositoryTagsRequest) (response *DeleteRepositoryTagsResponse, err error) {
    if request == nil {
        request = NewDeleteRepositoryTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRepositoryTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRepositoryTagsResponse()
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
    return c.DeleteSecurityPolicyWithContext(context.Background(), request)
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
func (c *Client) DeleteSecurityPolicyWithContext(ctx context.Context, request *DeleteSecurityPolicyRequest) (response *DeleteSecurityPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSignaturePolicyRequest() (request *DeleteSignaturePolicyRequest) {
    request = &DeleteSignaturePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DeleteSignaturePolicy")
    
    
    return
}

func NewDeleteSignaturePolicyResponse() (response *DeleteSignaturePolicyResponse) {
    response = &DeleteSignaturePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSignaturePolicy
// 删除命名空间加签策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
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
func (c *Client) DeleteSignaturePolicy(request *DeleteSignaturePolicyRequest) (response *DeleteSignaturePolicyResponse, err error) {
    return c.DeleteSignaturePolicyWithContext(context.Background(), request)
}

// DeleteSignaturePolicy
// 删除命名空间加签策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
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
func (c *Client) DeleteSignaturePolicyWithContext(ctx context.Context, request *DeleteSignaturePolicyRequest) (response *DeleteSignaturePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteSignaturePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSignaturePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSignaturePolicyResponse()
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
    return c.DeleteTagRetentionRuleWithContext(context.Background(), request)
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
func (c *Client) DeleteTagRetentionRuleWithContext(ctx context.Context, request *DeleteTagRetentionRuleRequest) (response *DeleteTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewDeleteTagRetentionRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTagRetentionRule require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DeleteWebhookTrigger(request *DeleteWebhookTriggerRequest) (response *DeleteWebhookTriggerResponse, err error) {
    return c.DeleteWebhookTriggerWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DeleteWebhookTriggerWithContext(ctx context.Context, request *DeleteWebhookTriggerRequest) (response *DeleteWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeApplicationTriggerLogPersonalWithContext(context.Background(), request)
}

// DescribeApplicationTriggerLogPersonal
// 用于查询应用更新触发器触发日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeApplicationTriggerLogPersonalWithContext(ctx context.Context, request *DescribeApplicationTriggerLogPersonalRequest) (response *DescribeApplicationTriggerLogPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationTriggerLogPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationTriggerLogPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeApplicationTriggerPersonalWithContext(context.Background(), request)
}

// DescribeApplicationTriggerPersonal
// 用于查询应用更新触发器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DescribeApplicationTriggerPersonalWithContext(ctx context.Context, request *DescribeApplicationTriggerPersonalRequest) (response *DescribeApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationTriggerPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationTriggerPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeChartDownloadInfoWithContext(context.Background(), request)
}

// DescribeChartDownloadInfo
// 用于在企业版中返回Chart的下载信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChartDownloadInfoWithContext(ctx context.Context, request *DescribeChartDownloadInfoRequest) (response *DescribeChartDownloadInfoResponse, err error) {
    if request == nil {
        request = NewDescribeChartDownloadInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChartDownloadInfo require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExternalEndpointStatus(request *DescribeExternalEndpointStatusRequest) (response *DescribeExternalEndpointStatusResponse, err error) {
    return c.DescribeExternalEndpointStatusWithContext(context.Background(), request)
}

// DescribeExternalEndpointStatus
// 查询实例公网访问入口状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExternalEndpointStatusWithContext(ctx context.Context, request *DescribeExternalEndpointStatusRequest) (response *DescribeExternalEndpointStatusResponse, err error) {
    if request == nil {
        request = NewDescribeExternalEndpointStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalEndpointStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeFavorRepositoryPersonalWithContext(context.Background(), request)
}

// DescribeFavorRepositoryPersonal
// 查询个人收藏仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFavorRepositoryPersonalWithContext(ctx context.Context, request *DescribeFavorRepositoryPersonalRequest) (response *DescribeFavorRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeFavorRepositoryPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFavorRepositoryPersonal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFavorRepositoryPersonalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGCJobsRequest() (request *DescribeGCJobsRequest) {
    request = &DescribeGCJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeGCJobs")
    
    
    return
}

func NewDescribeGCJobsResponse() (response *DescribeGCJobsResponse) {
    response = &DescribeGCJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGCJobs
// GC 最近10条历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGCJobs(request *DescribeGCJobsRequest) (response *DescribeGCJobsResponse, err error) {
    return c.DescribeGCJobsWithContext(context.Background(), request)
}

// DescribeGCJobs
// GC 最近10条历史
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGCJobsWithContext(ctx context.Context, request *DescribeGCJobsRequest) (response *DescribeGCJobsResponse, err error) {
    if request == nil {
        request = NewDescribeGCJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGCJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGCJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAccelerateServiceRequest() (request *DescribeImageAccelerateServiceRequest) {
    request = &DescribeImageAccelerateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeImageAccelerateService")
    
    
    return
}

func NewDescribeImageAccelerateServiceResponse() (response *DescribeImageAccelerateServiceResponse) {
    response = &DescribeImageAccelerateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageAccelerateService
// 查询镜像加速服务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImageAccelerateService(request *DescribeImageAccelerateServiceRequest) (response *DescribeImageAccelerateServiceResponse, err error) {
    return c.DescribeImageAccelerateServiceWithContext(context.Background(), request)
}

// DescribeImageAccelerateService
// 查询镜像加速服务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeImageAccelerateServiceWithContext(ctx context.Context, request *DescribeImageAccelerateServiceRequest) (response *DescribeImageAccelerateServiceResponse, err error) {
    if request == nil {
        request = NewDescribeImageAccelerateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAccelerateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAccelerateServiceResponse()
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
    return c.DescribeImageFilterPersonalWithContext(context.Background(), request)
}

// DescribeImageFilterPersonal
// 用于在个人版中查询与指定tag镜像内容相同的tag列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeImageFilterPersonalWithContext(ctx context.Context, request *DescribeImageFilterPersonalRequest) (response *DescribeImageFilterPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageFilterPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageFilterPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeImageLifecycleGlobalPersonalWithContext(context.Background(), request)
}

// DescribeImageLifecycleGlobalPersonal
// 用于获取个人版全局镜像版本自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeImageLifecycleGlobalPersonalWithContext(ctx context.Context, request *DescribeImageLifecycleGlobalPersonalRequest) (response *DescribeImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageLifecycleGlobalPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageLifecycleGlobalPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeImageLifecyclePersonalWithContext(context.Background(), request)
}

// DescribeImageLifecyclePersonal
// 用于获取个人版仓库中自动清理策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeImageLifecyclePersonalWithContext(ctx context.Context, request *DescribeImageLifecyclePersonalRequest) (response *DescribeImageLifecyclePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImageLifecyclePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageLifecyclePersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeImageManifestsWithContext(context.Background(), request)
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
func (c *Client) DescribeImageManifestsWithContext(ctx context.Context, request *DescribeImageManifestsRequest) (response *DescribeImageManifestsResponse, err error) {
    if request == nil {
        request = NewDescribeImageManifestsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageManifests require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeImagePersonalWithContext(context.Background(), request)
}

// DescribeImagePersonal
// 用于获取个人版镜像仓库tag列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRUNAUTHORIZED = "InternalError.ErrUnauthorized"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeImagePersonalWithContext(ctx context.Context, request *DescribeImagePersonalRequest) (response *DescribeImagePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeImagePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImagePersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeImagesWithContext(context.Background(), request)
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
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImages require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImmutableTagRules(request *DescribeImmutableTagRulesRequest) (response *DescribeImmutableTagRulesResponse, err error) {
    return c.DescribeImmutableTagRulesWithContext(context.Background(), request)
}

// DescribeImmutableTagRules
// 列出镜像不可变规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORTCRINTERNAL = "InternalError.ErrorTcrInternal"
//  INTERNALERROR_ERRORTCRUNAUTHORIZED = "InternalError.ErrorTcrUnauthorized"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeImmutableTagRulesWithContext(ctx context.Context, request *DescribeImmutableTagRulesRequest) (response *DescribeImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewDescribeImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImmutableTagRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAllNamespacesRequest() (request *DescribeInstanceAllNamespacesRequest) {
    request = &DescribeInstanceAllNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceAllNamespaces")
    
    
    return
}

func NewDescribeInstanceAllNamespacesResponse() (response *DescribeInstanceAllNamespacesResponse) {
    response = &DescribeInstanceAllNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceAllNamespaces
// 查询所有实例命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
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
func (c *Client) DescribeInstanceAllNamespaces(request *DescribeInstanceAllNamespacesRequest) (response *DescribeInstanceAllNamespacesResponse, err error) {
    return c.DescribeInstanceAllNamespacesWithContext(context.Background(), request)
}

// DescribeInstanceAllNamespaces
// 查询所有实例命名空间列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
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
func (c *Client) DescribeInstanceAllNamespacesWithContext(ctx context.Context, request *DescribeInstanceAllNamespacesRequest) (response *DescribeInstanceAllNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAllNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAllNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAllNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCustomizedDomainRequest() (request *DescribeInstanceCustomizedDomainRequest) {
    request = &DescribeInstanceCustomizedDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeInstanceCustomizedDomain")
    
    
    return
}

func NewDescribeInstanceCustomizedDomainResponse() (response *DescribeInstanceCustomizedDomainResponse) {
    response = &DescribeInstanceCustomizedDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceCustomizedDomain
// 查询实例自定义域名列表
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
func (c *Client) DescribeInstanceCustomizedDomain(request *DescribeInstanceCustomizedDomainRequest) (response *DescribeInstanceCustomizedDomainResponse, err error) {
    return c.DescribeInstanceCustomizedDomainWithContext(context.Background(), request)
}

// DescribeInstanceCustomizedDomain
// 查询实例自定义域名列表
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
func (c *Client) DescribeInstanceCustomizedDomainWithContext(ctx context.Context, request *DescribeInstanceCustomizedDomainRequest) (response *DescribeInstanceCustomizedDomainResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCustomizedDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceCustomizedDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceCustomizedDomainResponse()
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
    return c.DescribeInstanceStatusWithContext(context.Background(), request)
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
func (c *Client) DescribeInstanceStatusWithContext(ctx context.Context, request *DescribeInstanceStatusRequest) (response *DescribeInstanceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeInstanceTokenWithContext(context.Background(), request)
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
func (c *Client) DescribeInstanceTokenWithContext(ctx context.Context, request *DescribeInstanceTokenRequest) (response *DescribeInstanceTokenResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceToken require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
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
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 查询实例信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT_ERRORINSTANCENOTRUNNING = "ResourceInsufficient.ErrorInstanceNotRunning"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeInternalEndpointDnsStatusWithContext(context.Background(), request)
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
func (c *Client) DescribeInternalEndpointDnsStatusWithContext(ctx context.Context, request *DescribeInternalEndpointDnsStatusRequest) (response *DescribeInternalEndpointDnsStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInternalEndpointDnsStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInternalEndpointDnsStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeInternalEndpointsWithContext(context.Background(), request)
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
func (c *Client) DescribeInternalEndpointsWithContext(ctx context.Context, request *DescribeInternalEndpointsRequest) (response *DescribeInternalEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeInternalEndpointsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInternalEndpoints require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeNamespacePersonalWithContext(context.Background(), request)
}

// DescribeNamespacePersonal
// 查询个人版命名空间信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DescribeNamespacePersonalWithContext(ctx context.Context, request *DescribeNamespacePersonalRequest) (response *DescribeNamespacePersonalResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNamespacePersonal require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
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
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    return c.DescribeNamespacesWithContext(context.Background(), request)
}

// DescribeNamespaces
// 查询命名空间列表或指定命名空间信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DEPENDENCEERROR = "FailedOperation.DependenceError"
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_ERRORTCRRESOURCECONFLICT = "FailedOperation.ErrorTcrResourceConflict"
//  FAILEDOPERATION_ERRORTCRUNAUTHORIZED = "FailedOperation.ErrorTcrUnauthorized"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_OPERATIONCANCEL = "FailedOperation.OperationCancel"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
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
func (c *Client) DescribeNamespacesWithContext(ctx context.Context, request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeNamespacesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcr", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// 用于在TCR中获取可用区域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 用于在TCR中获取可用区域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
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
    return c.DescribeReplicationInstanceCreateTasksWithContext(context.Background(), request)
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
func (c *Client) DescribeReplicationInstanceCreateTasksWithContext(ctx context.Context, request *DescribeReplicationInstanceCreateTasksRequest) (response *DescribeReplicationInstanceCreateTasksResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstanceCreateTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationInstanceCreateTasks require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeReplicationInstanceSyncStatusWithContext(context.Background(), request)
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
func (c *Client) DescribeReplicationInstanceSyncStatusWithContext(ctx context.Context, request *DescribeReplicationInstanceSyncStatusRequest) (response *DescribeReplicationInstanceSyncStatusResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstanceSyncStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationInstanceSyncStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeReplicationInstancesWithContext(context.Background(), request)
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
func (c *Client) DescribeReplicationInstancesWithContext(ctx context.Context, request *DescribeReplicationInstancesRequest) (response *DescribeReplicationInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationInstances require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
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
func (c *Client) DescribeRepositories(request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    return c.DescribeRepositoriesWithContext(context.Background(), request)
}

// DescribeRepositories
// 查询镜像仓库列表或指定镜像仓库信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
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
func (c *Client) DescribeRepositoriesWithContext(ctx context.Context, request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositories require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeRepositoryFilterPersonalWithContext(context.Background(), request)
}

// DescribeRepositoryFilterPersonal
// 用于在个人版镜像仓库中，获取满足输入搜索条件的用户镜像仓库
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) DescribeRepositoryFilterPersonalWithContext(ctx context.Context, request *DescribeRepositoryFilterPersonalRequest) (response *DescribeRepositoryFilterPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryFilterPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositoryFilterPersonal require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRepositoryOwnerPersonal(request *DescribeRepositoryOwnerPersonalRequest) (response *DescribeRepositoryOwnerPersonalResponse, err error) {
    return c.DescribeRepositoryOwnerPersonalWithContext(context.Background(), request)
}

// DescribeRepositoryOwnerPersonal
// 用于在个人版中获取用户全部的镜像仓库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRepositoryOwnerPersonalWithContext(ctx context.Context, request *DescribeRepositoryOwnerPersonalRequest) (response *DescribeRepositoryOwnerPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryOwnerPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositoryOwnerPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeRepositoryPersonalWithContext(context.Background(), request)
}

// DescribeRepositoryPersonal
// 查询个人版仓库信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) DescribeRepositoryPersonalWithContext(ctx context.Context, request *DescribeRepositoryPersonalRequest) (response *DescribeRepositoryPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeRepositoryPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepositoryPersonal require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETSECURITYPOLICYFAIL = "FailedOperation.GetSecurityPolicyFail"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPolicies(request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    return c.DescribeSecurityPoliciesWithContext(context.Background(), request)
}

// DescribeSecurityPolicies
// 查询实例公网访问白名单策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORGETDBDATAERROR = "FailedOperation.ErrorGetDBDataError"
//  FAILEDOPERATION_GETDBDATAERROR = "FailedOperation.GetDBDataError"
//  FAILEDOPERATION_GETSECURITYPOLICYFAIL = "FailedOperation.GetSecurityPolicyFail"
//  FAILEDOPERATION_GETTCRCLIENT = "FailedOperation.GetTcrClient"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityPoliciesWithContext(ctx context.Context, request *DescribeSecurityPoliciesRequest) (response *DescribeSecurityPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityPolicies require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeTagRetentionExecutionWithContext(context.Background(), request)
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
func (c *Client) DescribeTagRetentionExecutionWithContext(ctx context.Context, request *DescribeTagRetentionExecutionRequest) (response *DescribeTagRetentionExecutionResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionExecutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagRetentionExecution require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeTagRetentionExecutionTaskWithContext(context.Background(), request)
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
func (c *Client) DescribeTagRetentionExecutionTaskWithContext(ctx context.Context, request *DescribeTagRetentionExecutionTaskRequest) (response *DescribeTagRetentionExecutionTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionExecutionTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagRetentionExecutionTask require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
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
    return c.DescribeTagRetentionRulesWithContext(context.Background(), request)
}

// DescribeTagRetentionRules
// 查询版本保留规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYCOREBODY = "FailedOperation.EmptyCoreBody"
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
func (c *Client) DescribeTagRetentionRulesWithContext(ctx context.Context, request *DescribeTagRetentionRulesRequest) (response *DescribeTagRetentionRulesResponse, err error) {
    if request == nil {
        request = NewDescribeTagRetentionRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagRetentionRules require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeUserQuotaPersonalWithContext(context.Background(), request)
}

// DescribeUserQuotaPersonal
// 查询个人用户配额
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUserQuotaPersonalWithContext(ctx context.Context, request *DescribeUserQuotaPersonalRequest) (response *DescribeUserQuotaPersonalResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotaPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserQuotaPersonal require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTrigger(request *DescribeWebhookTriggerRequest) (response *DescribeWebhookTriggerResponse, err error) {
    return c.DescribeWebhookTriggerWithContext(context.Background(), request)
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
//  INVALIDPARAMETER_ERRORTCRINVALIDPARAMETER = "InvalidParameter.ErrorTcrInvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
func (c *Client) DescribeWebhookTriggerWithContext(ctx context.Context, request *DescribeWebhookTriggerRequest) (response *DescribeWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewDescribeWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWebhookTriggerLogWithContext(context.Background(), request)
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
func (c *Client) DescribeWebhookTriggerLogWithContext(ctx context.Context, request *DescribeWebhookTriggerLogRequest) (response *DescribeWebhookTriggerLogResponse, err error) {
    if request == nil {
        request = NewDescribeWebhookTriggerLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebhookTriggerLog require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DownloadHelmChart(request *DownloadHelmChartRequest) (response *DownloadHelmChartResponse, err error) {
    return c.DownloadHelmChartWithContext(context.Background(), request)
}

// DownloadHelmChart
// 用于在TCR中下载helm chart
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DownloadHelmChartWithContext(ctx context.Context, request *DownloadHelmChartRequest) (response *DownloadHelmChartResponse, err error) {
    if request == nil {
        request = NewDownloadHelmChartRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadHelmChart require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DuplicateImagePersonalWithContext(context.Background(), request)
}

// DuplicateImagePersonal
// 用于在个人版镜像仓库中复制镜像版本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOTAG = "ResourceNotFound.ErrNoTag"
func (c *Client) DuplicateImagePersonalWithContext(ctx context.Context, request *DuplicateImagePersonalRequest) (response *DuplicateImagePersonalResponse, err error) {
    if request == nil {
        request = NewDuplicateImagePersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DuplicateImagePersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ManageExternalEndpointWithContext(context.Background(), request)
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
func (c *Client) ManageExternalEndpointWithContext(ctx context.Context, request *ManageExternalEndpointRequest) (response *ManageExternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageExternalEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageExternalEndpoint require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ManageImageLifecycleGlobalPersonalWithContext(context.Background(), request)
}

// ManageImageLifecycleGlobalPersonal
// 用于设置个人版全局镜像版本自动清理策略
//
// 可能返回的错误码:
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ManageImageLifecycleGlobalPersonalWithContext(ctx context.Context, request *ManageImageLifecycleGlobalPersonalRequest) (response *ManageImageLifecycleGlobalPersonalResponse, err error) {
    if request == nil {
        request = NewManageImageLifecycleGlobalPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageImageLifecycleGlobalPersonal require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
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
    return c.ManageInternalEndpointWithContext(context.Background(), request)
}

// ManageInternalEndpoint
// 管理实例内网访问VPC链接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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
func (c *Client) ManageInternalEndpointWithContext(ctx context.Context, request *ManageInternalEndpointRequest) (response *ManageInternalEndpointResponse, err error) {
    if request == nil {
        request = NewManageInternalEndpointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageInternalEndpoint require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ManageReplicationWithContext(context.Background(), request)
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
func (c *Client) ManageReplicationWithContext(ctx context.Context, request *ManageReplicationRequest) (response *ManageReplicationResponse, err error) {
    if request == nil {
        request = NewManageReplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageReplication require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyApplicationTriggerPersonalWithContext(context.Background(), request)
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
func (c *Client) ModifyApplicationTriggerPersonalWithContext(ctx context.Context, request *ModifyApplicationTriggerPersonalRequest) (response *ModifyApplicationTriggerPersonalResponse, err error) {
    if request == nil {
        request = NewModifyApplicationTriggerPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationTriggerPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyImmutableTagRulesWithContext(context.Background(), request)
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
func (c *Client) ModifyImmutableTagRulesWithContext(ctx context.Context, request *ModifyImmutableTagRulesRequest) (response *ModifyImmutableTagRulesResponse, err error) {
    if request == nil {
        request = NewModifyImmutableTagRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyImmutableTagRules require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyInstanceWithContext(context.Background(), request)
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
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyInstanceTokenWithContext(context.Background(), request)
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
func (c *Client) ModifyInstanceTokenWithContext(ctx context.Context, request *ModifyInstanceTokenRequest) (response *ModifyInstanceTokenResponse, err error) {
    if request == nil {
        request = NewModifyInstanceTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceToken require credential")
    }

    request.SetContext(ctx)
    
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
// 更新命名空间信息
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
    return c.ModifyNamespaceWithContext(context.Background(), request)
}

// ModifyNamespace
// 更新命名空间信息
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
func (c *Client) ModifyNamespaceWithContext(ctx context.Context, request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
    if request == nil {
        request = NewModifyNamespaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNamespace require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRepository(request *ModifyRepositoryRequest) (response *ModifyRepositoryResponse, err error) {
    return c.ModifyRepositoryWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TCRRESOURCENOTFOUND = "ResourceNotFound.TcrResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyRepositoryWithContext(ctx context.Context, request *ModifyRepositoryRequest) (response *ModifyRepositoryResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRepository require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyRepositoryAccessPersonalWithContext(context.Background(), request)
}

// ModifyRepositoryAccessPersonal
// 用于更新个人版镜像仓库的访问属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) ModifyRepositoryAccessPersonalWithContext(ctx context.Context, request *ModifyRepositoryAccessPersonalRequest) (response *ModifyRepositoryAccessPersonalResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryAccessPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRepositoryAccessPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyRepositoryInfoPersonalWithContext(context.Background(), request)
}

// ModifyRepositoryInfoPersonal
// 用于在个人版镜像仓库中更新容器镜像描述
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOREPO = "ResourceNotFound.ErrNoRepo"
func (c *Client) ModifyRepositoryInfoPersonalWithContext(ctx context.Context, request *ModifyRepositoryInfoPersonalRequest) (response *ModifyRepositoryInfoPersonalResponse, err error) {
    if request == nil {
        request = NewModifyRepositoryInfoPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRepositoryInfoPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifySecurityPolicyWithContext(context.Background(), request)
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
    return c.ModifyTagRetentionRuleWithContext(context.Background(), request)
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
func (c *Client) ModifyTagRetentionRuleWithContext(ctx context.Context, request *ModifyTagRetentionRuleRequest) (response *ModifyTagRetentionRuleResponse, err error) {
    if request == nil {
        request = NewModifyTagRetentionRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTagRetentionRule require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyUserPasswordPersonalWithContext(context.Background(), request)
}

// ModifyUserPasswordPersonal
// 修改个人用户登录密码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ERRNOUSER = "ResourceNotFound.ErrNoUser"
func (c *Client) ModifyUserPasswordPersonalWithContext(ctx context.Context, request *ModifyUserPasswordPersonalRequest) (response *ModifyUserPasswordPersonalResponse, err error) {
    if request == nil {
        request = NewModifyUserPasswordPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserPasswordPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifyWebhookTriggerWithContext(context.Background(), request)
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
func (c *Client) ModifyWebhookTriggerWithContext(ctx context.Context, request *ModifyWebhookTriggerRequest) (response *ModifyWebhookTriggerResponse, err error) {
    if request == nil {
        request = NewModifyWebhookTriggerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWebhookTrigger require credential")
    }

    request.SetContext(ctx)
    
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
    return c.RenewInstanceWithContext(context.Background(), request)
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
func (c *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ValidateNamespaceExistPersonalWithContext(context.Background(), request)
}

// ValidateNamespaceExistPersonal
// 查询个人版用户命名空间是否存在
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ValidateNamespaceExistPersonalWithContext(ctx context.Context, request *ValidateNamespaceExistPersonalRequest) (response *ValidateNamespaceExistPersonalResponse, err error) {
    if request == nil {
        request = NewValidateNamespaceExistPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ValidateNamespaceExistPersonal require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ValidateRepositoryExistPersonalWithContext(context.Background(), request)
}

// ValidateRepositoryExistPersonal
// 用于判断个人版仓库是否存在
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ERRNSMISMATCH = "InvalidParameter.ErrNSMisMatch"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ValidateRepositoryExistPersonalWithContext(ctx context.Context, request *ValidateRepositoryExistPersonalRequest) (response *ValidateRepositoryExistPersonalResponse, err error) {
    if request == nil {
        request = NewValidateRepositoryExistPersonalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ValidateRepositoryExistPersonal require credential")
    }

    request.SetContext(ctx)
    
    response = NewValidateRepositoryExistPersonalResponse()
    err = c.Send(request, response)
    return
}
