// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20231024

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-10-24"

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


func NewCreateAddressPoolRequest() (request *CreateAddressPoolRequest) {
    request = &CreateAddressPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "CreateAddressPool")
    
    
    return
}

func NewCreateAddressPoolResponse() (response *CreateAddressPoolResponse) {
    response = &CreateAddressPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAddressPool
// 创建地址池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ADDREXCEEDEDLIMIT = "FailedOperation.AddrExceededLimit"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATEPOOLFAILED = "FailedOperation.CreatePoolFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_POOLEXCEEDEDLIMIT = "FailedOperation.PoolExceededLimit"
//  FAILEDOPERATION_TASKEXCEEDEDLIMIT = "FailedOperation.TaskExceededLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DbErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_POOLADDRDUPLICATED = "InvalidParameterValue.PoolAddrDuplicated"
//  INVALIDPARAMETERVALUE_POOLADDREMPTY = "InvalidParameterValue.PoolAddrEmpty"
//  INVALIDPARAMETERVALUE_POOLADDRINVALID = "InvalidParameterValue.PoolAddrInvalid"
//  INVALIDPARAMETERVALUE_POOLADDRMIXED = "InvalidParameterValue.PoolAddrMixed"
//  INVALIDPARAMETERVALUE_POOLADDRWEIGHTEMPTY = "InvalidParameterValue.PoolAddrWeightEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_POOLNAMEDUPLICATE = "UnsupportedOperation.PoolNameDuplicate"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateAddressPool(request *CreateAddressPoolRequest) (response *CreateAddressPoolResponse, err error) {
    return c.CreateAddressPoolWithContext(context.Background(), request)
}

// CreateAddressPool
// 创建地址池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ADDREXCEEDEDLIMIT = "FailedOperation.AddrExceededLimit"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATEPOOLFAILED = "FailedOperation.CreatePoolFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_POOLEXCEEDEDLIMIT = "FailedOperation.PoolExceededLimit"
//  FAILEDOPERATION_TASKEXCEEDEDLIMIT = "FailedOperation.TaskExceededLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DbErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_POOLADDRDUPLICATED = "InvalidParameterValue.PoolAddrDuplicated"
//  INVALIDPARAMETERVALUE_POOLADDREMPTY = "InvalidParameterValue.PoolAddrEmpty"
//  INVALIDPARAMETERVALUE_POOLADDRINVALID = "InvalidParameterValue.PoolAddrInvalid"
//  INVALIDPARAMETERVALUE_POOLADDRMIXED = "InvalidParameterValue.PoolAddrMixed"
//  INVALIDPARAMETERVALUE_POOLADDRWEIGHTEMPTY = "InvalidParameterValue.PoolAddrWeightEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_POOLNAMEDUPLICATE = "UnsupportedOperation.PoolNameDuplicate"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateAddressPoolWithContext(ctx context.Context, request *CreateAddressPoolRequest) (response *CreateAddressPoolResponse, err error) {
    if request == nil {
        request = NewCreateAddressPoolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "CreateAddressPool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAddressPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAddressPoolResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// 创建实例接口，仅供免费实例使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATEINSTANCEFAILED = "FailedOperation.CreateInstanceFailed"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCESSTYPEINVALID = "InvalidParameter.AccessTypeInvalid"
//  INVALIDPARAMETER_GLOBALACCESSDOMAINCONFLICT = "InvalidParameter.GlobalAccessDomainConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_ACCESSDOMAINUNAUTHORIZED = "UnsupportedOperation.AccessDomainUnauthorized"
//  UNSUPPORTEDOPERATION_DOMAINTTLINVALID = "UnsupportedOperation.DomainTtlInvalid"
//  UNSUPPORTEDOPERATION_INSTANCENAMEDUPLICATE = "UnsupportedOperation.InstanceNameDuplicate"
//  UNSUPPORTEDOPERATION_PACKAGEUNAUTHORIZED = "UnsupportedOperation.PackageUnauthorized"
//  UNSUPPORTEDOPERATION_TTLINVALID = "UnsupportedOperation.TTLInvalid"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 创建实例接口，仅供免费实例使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATEINSTANCEFAILED = "FailedOperation.CreateInstanceFailed"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCESSTYPEINVALID = "InvalidParameter.AccessTypeInvalid"
//  INVALIDPARAMETER_GLOBALACCESSDOMAINCONFLICT = "InvalidParameter.GlobalAccessDomainConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_ACCESSDOMAINUNAUTHORIZED = "UnsupportedOperation.AccessDomainUnauthorized"
//  UNSUPPORTEDOPERATION_DOMAINTTLINVALID = "UnsupportedOperation.DomainTtlInvalid"
//  UNSUPPORTEDOPERATION_INSTANCENAMEDUPLICATE = "UnsupportedOperation.InstanceNameDuplicate"
//  UNSUPPORTEDOPERATION_PACKAGEUNAUTHORIZED = "UnsupportedOperation.PackageUnauthorized"
//  UNSUPPORTEDOPERATION_TTLINVALID = "UnsupportedOperation.TTLInvalid"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "CreateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMonitorRequest() (request *CreateMonitorRequest) {
    request = &CreateMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "CreateMonitor")
    
    
    return
}

func NewCreateMonitorResponse() (response *CreateMonitorResponse) {
    response = &CreateMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMonitor
// 新增监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATEMONITORFAILED = "FailedOperation.CreateMonitorFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateMonitor(request *CreateMonitorRequest) (response *CreateMonitorResponse, err error) {
    return c.CreateMonitorWithContext(context.Background(), request)
}

// CreateMonitor
// 新增监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATEMONITORFAILED = "FailedOperation.CreateMonitorFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateMonitorWithContext(ctx context.Context, request *CreateMonitorRequest) (response *CreateMonitorResponse, err error) {
    if request == nil {
        request = NewCreateMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "CreateMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePackageAndPayRequest() (request *CreatePackageAndPayRequest) {
    request = &CreatePackageAndPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "CreatePackageAndPay")
    
    
    return
}

func NewCreatePackageAndPayResponse() (response *CreatePackageAndPayResponse) {
    response = &CreatePackageAndPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePackageAndPay
// 购买套餐并支付，此接口会在余额扣费，谨慎调用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_DEALANDPAYFAILED = "FailedOperation.DealAndPayFailed"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEALSTANDARDMODIFYFAILED = "InvalidParameter.DealStandardModifyFailed"
//  INVALIDPARAMETER_DEALSTANDARDMODIFYINVALID = "InvalidParameter.DealStandardModifyInvalid"
//  INVALIDPARAMETER_DEALTASKMODIFYINVALID = "InvalidParameter.DealTaskModifyInvalid"
//  INVALIDPARAMETER_DEALTASKMODIFYNUMLESS = "InvalidParameter.DealTaskModifyNumLess"
//  INVALIDPARAMETER_DEALTASKMODIFYNUMSAME = "InvalidParameter.DealTaskModifyNumSame"
//  INVALIDPARAMETER_DEALULTIMATEMODIFYFAILED = "InvalidParameter.DealUltimateModifyFailed"
//  INVALIDPARAMETER_RESOURCEIDEMPTY = "InvalidParameter.ResourceIdEmpty"
//  INVALIDPARAMETER_TIMESPANEMPTY = "InvalidParameter.TimeSpanEmpty"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_RESOURCEIDNOTMATCH = "UnauthorizedOperation.ResourceIdNotMatch"
//  UNAUTHORIZEDOPERATION_RESOURCEIDUNAUTHORIZED = "UnauthorizedOperation.ResourceIdUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreatePackageAndPay(request *CreatePackageAndPayRequest) (response *CreatePackageAndPayResponse, err error) {
    return c.CreatePackageAndPayWithContext(context.Background(), request)
}

// CreatePackageAndPay
// 购买套餐并支付，此接口会在余额扣费，谨慎调用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_DEALANDPAYFAILED = "FailedOperation.DealAndPayFailed"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEALSTANDARDMODIFYFAILED = "InvalidParameter.DealStandardModifyFailed"
//  INVALIDPARAMETER_DEALSTANDARDMODIFYINVALID = "InvalidParameter.DealStandardModifyInvalid"
//  INVALIDPARAMETER_DEALTASKMODIFYINVALID = "InvalidParameter.DealTaskModifyInvalid"
//  INVALIDPARAMETER_DEALTASKMODIFYNUMLESS = "InvalidParameter.DealTaskModifyNumLess"
//  INVALIDPARAMETER_DEALTASKMODIFYNUMSAME = "InvalidParameter.DealTaskModifyNumSame"
//  INVALIDPARAMETER_DEALULTIMATEMODIFYFAILED = "InvalidParameter.DealUltimateModifyFailed"
//  INVALIDPARAMETER_RESOURCEIDEMPTY = "InvalidParameter.ResourceIdEmpty"
//  INVALIDPARAMETER_TIMESPANEMPTY = "InvalidParameter.TimeSpanEmpty"
//  INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_RESOURCEIDNOTMATCH = "UnauthorizedOperation.ResourceIdNotMatch"
//  UNAUTHORIZEDOPERATION_RESOURCEIDUNAUTHORIZED = "UnauthorizedOperation.ResourceIdUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreatePackageAndPayWithContext(ctx context.Context, request *CreatePackageAndPayRequest) (response *CreatePackageAndPayResponse, err error) {
    if request == nil {
        request = NewCreatePackageAndPayRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "CreatePackageAndPay")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePackageAndPay require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePackageAndPayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStrategyRequest() (request *CreateStrategyRequest) {
    request = &CreateStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "CreateStrategy")
    
    
    return
}

func NewCreateStrategyResponse() (response *CreateStrategyResponse) {
    response = &CreateStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStrategy
// 新建策略接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATESTRATEGYFAILED = "FailedOperation.CreateStrategyFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DbErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETER_TRAFFICSTRATEGYINVALID = "InvalidParameter.TrafficStrategyInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TRAFFICSTRATEGYINVALID = "InvalidParameterValue.TrafficStrategyInvalid"
//  INVALIDPARAMETERVALUE_TRAFFICSTRATEGYWEIGHTEMPTY = "InvalidParameterValue.TrafficStrategyWeightEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_NOINSTANCE = "ResourceInsufficient.NoInstance"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_INSTANCEUNAUTHORIZED = "UnauthorizedOperation.InstanceUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINBINDPOOLUNAUTHORIZED = "UnauthorizedOperation.MainBindPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINPOOLUNAUTHORIZED = "UnauthorizedOperation.MainPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_FALLBACKPOOLINVALID = "UnsupportedOperation.FallBackPoolInvalid"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLDUPLICATE = "UnsupportedOperation.MainBindPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLEMPTY = "UnsupportedOperation.MainBindPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLLEVELEXCEED = "UnsupportedOperation.MainBindPoolLevelExceed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLMIXED = "UnsupportedOperation.MainBindPoolMixed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLNUMINVALID = "UnsupportedOperation.MainBindPoolNumInvalid"
//  UNSUPPORTEDOPERATION_MAINPOOLDUPLICATE = "UnsupportedOperation.MainPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINPOOLEMPTY = "UnsupportedOperation.MainPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINPOOLSURVIVENUMINVALID = "UnsupportedOperation.MainPoolSurviveNumInvalid"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_RESOURCEPOOLSOURCENOTEXIST = "UnsupportedOperation.ResourcePoolSourceNotExist"
//  UNSUPPORTEDOPERATION_RESOURCESOURCEDUPLICATED = "UnsupportedOperation.ResourceSourceDuplicated"
//  UNSUPPORTEDOPERATION_STRATEGYNAMEDUPLICATE = "UnsupportedOperation.StrategyNameDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEDUPLICATE = "UnsupportedOperation.StrategySourceDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEINUSED = "UnsupportedOperation.StrategySourceInUsed"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEUNAUTHORIZED = "UnsupportedOperation.StrategySourceUnauthorized"
//  UNSUPPORTEDOPERATION_STRATEGYSWITCHTYPEINVALID = "UnsupportedOperation.StrategySwitchTypeInvalid"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateStrategy(request *CreateStrategyRequest) (response *CreateStrategyResponse, err error) {
    return c.CreateStrategyWithContext(context.Background(), request)
}

// CreateStrategy
// 新建策略接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_CREATESTRATEGYFAILED = "FailedOperation.CreateStrategyFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERR = "InternalError.DbErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETER_TRAFFICSTRATEGYINVALID = "InvalidParameter.TrafficStrategyInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TRAFFICSTRATEGYINVALID = "InvalidParameterValue.TrafficStrategyInvalid"
//  INVALIDPARAMETERVALUE_TRAFFICSTRATEGYWEIGHTEMPTY = "InvalidParameterValue.TrafficStrategyWeightEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_NOINSTANCE = "ResourceInsufficient.NoInstance"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_INSTANCEUNAUTHORIZED = "UnauthorizedOperation.InstanceUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINBINDPOOLUNAUTHORIZED = "UnauthorizedOperation.MainBindPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINPOOLUNAUTHORIZED = "UnauthorizedOperation.MainPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_FALLBACKPOOLINVALID = "UnsupportedOperation.FallBackPoolInvalid"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLDUPLICATE = "UnsupportedOperation.MainBindPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLEMPTY = "UnsupportedOperation.MainBindPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLLEVELEXCEED = "UnsupportedOperation.MainBindPoolLevelExceed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLMIXED = "UnsupportedOperation.MainBindPoolMixed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLNUMINVALID = "UnsupportedOperation.MainBindPoolNumInvalid"
//  UNSUPPORTEDOPERATION_MAINPOOLDUPLICATE = "UnsupportedOperation.MainPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINPOOLEMPTY = "UnsupportedOperation.MainPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINPOOLSURVIVENUMINVALID = "UnsupportedOperation.MainPoolSurviveNumInvalid"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_RESOURCEPOOLSOURCENOTEXIST = "UnsupportedOperation.ResourcePoolSourceNotExist"
//  UNSUPPORTEDOPERATION_RESOURCESOURCEDUPLICATED = "UnsupportedOperation.ResourceSourceDuplicated"
//  UNSUPPORTEDOPERATION_STRATEGYNAMEDUPLICATE = "UnsupportedOperation.StrategyNameDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEDUPLICATE = "UnsupportedOperation.StrategySourceDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEINUSED = "UnsupportedOperation.StrategySourceInUsed"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEUNAUTHORIZED = "UnsupportedOperation.StrategySourceUnauthorized"
//  UNSUPPORTEDOPERATION_STRATEGYSWITCHTYPEINVALID = "UnsupportedOperation.StrategySwitchTypeInvalid"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) CreateStrategyWithContext(ctx context.Context, request *CreateStrategyRequest) (response *CreateStrategyResponse, err error) {
    if request == nil {
        request = NewCreateStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "CreateStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddressPoolRequest() (request *DeleteAddressPoolRequest) {
    request = &DeleteAddressPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DeleteAddressPool")
    
    
    return
}

func NewDeleteAddressPoolResponse() (response *DeleteAddressPoolResponse) {
    response = &DeleteAddressPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAddressPool
// 删除地址池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEERROR = "FailedOperation.DeleteError"
//  FAILEDOPERATION_DELETEPOOLFAILED = "FailedOperation.DeletePoolFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_POOLWASUSINGERROR = "FailedOperation.PoolWasUsingError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_POOLUNAUTHORIZED = "UnauthorizedOperation.PoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DeleteAddressPool(request *DeleteAddressPoolRequest) (response *DeleteAddressPoolResponse, err error) {
    return c.DeleteAddressPoolWithContext(context.Background(), request)
}

// DeleteAddressPool
// 删除地址池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEERROR = "FailedOperation.DeleteError"
//  FAILEDOPERATION_DELETEPOOLFAILED = "FailedOperation.DeletePoolFailed"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_POOLWASUSINGERROR = "FailedOperation.PoolWasUsingError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_POOLUNAUTHORIZED = "UnauthorizedOperation.PoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DeleteAddressPoolWithContext(ctx context.Context, request *DeleteAddressPoolRequest) (response *DeleteAddressPoolResponse, err error) {
    if request == nil {
        request = NewDeleteAddressPoolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DeleteAddressPool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAddressPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAddressPoolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMonitorRequest() (request *DeleteMonitorRequest) {
    request = &DeleteMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DeleteMonitor")
    
    
    return
}

func NewDeleteMonitorResponse() (response *DeleteMonitorResponse) {
    response = &DeleteMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMonitor
// 删除监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEERROR = "FailedOperation.DeleteError"
//  FAILEDOPERATION_DELETEMONITORFAILED = "FailedOperation.DeleteMonitorFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DeleteMonitor(request *DeleteMonitorRequest) (response *DeleteMonitorResponse, err error) {
    return c.DeleteMonitorWithContext(context.Background(), request)
}

// DeleteMonitor
// 删除监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETEERROR = "FailedOperation.DeleteError"
//  FAILEDOPERATION_DELETEMONITORFAILED = "FailedOperation.DeleteMonitorFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DeleteMonitorWithContext(ctx context.Context, request *DeleteMonitorRequest) (response *DeleteMonitorResponse, err error) {
    if request == nil {
        request = NewDeleteMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DeleteMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStrategyRequest() (request *DeleteStrategyRequest) {
    request = &DeleteStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DeleteStrategy")
    
    
    return
}

func NewDeleteStrategyResponse() (response *DeleteStrategyResponse) {
    response = &DeleteStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteStrategy
// 删除策略接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_DELETEERROR = "FailedOperation.DeleteError"
//  FAILEDOPERATION_DELETESTRATEGYFAILED = "FailedOperation.DeleteStrategyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_NOINSTANCE = "ResourceInsufficient.NoInstance"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_INSTANCEUNAUTHORIZED = "UnauthorizedOperation.InstanceUnauthorized"
//  UNAUTHORIZEDOPERATION_STRATEGYUNAUTHORIZED = "UnauthorizedOperation.StrategyUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DeleteStrategy(request *DeleteStrategyRequest) (response *DeleteStrategyResponse, err error) {
    return c.DeleteStrategyWithContext(context.Background(), request)
}

// DeleteStrategy
// 删除策略接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_DELETEERROR = "FailedOperation.DeleteError"
//  FAILEDOPERATION_DELETESTRATEGYFAILED = "FailedOperation.DeleteStrategyFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_NOINSTANCE = "ResourceInsufficient.NoInstance"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_INSTANCEUNAUTHORIZED = "UnauthorizedOperation.InstanceUnauthorized"
//  UNAUTHORIZEDOPERATION_STRATEGYUNAUTHORIZED = "UnauthorizedOperation.StrategyUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DeleteStrategyWithContext(ctx context.Context, request *DeleteStrategyRequest) (response *DeleteStrategyResponse, err error) {
    if request == nil {
        request = NewDeleteStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DeleteStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressLocationRequest() (request *DescribeAddressLocationRequest) {
    request = &DescribeAddressLocationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeAddressLocation")
    
    
    return
}

func NewDescribeAddressLocationResponse() (response *DescribeAddressLocationResponse) {
    response = &DescribeAddressLocationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddressLocation
// 获取地址所属地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeAddressLocation(request *DescribeAddressLocationRequest) (response *DescribeAddressLocationResponse, err error) {
    return c.DescribeAddressLocationWithContext(context.Background(), request)
}

// DescribeAddressLocation
// 获取地址所属地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeAddressLocationWithContext(ctx context.Context, request *DescribeAddressLocationRequest) (response *DescribeAddressLocationResponse, err error) {
    if request == nil {
        request = NewDescribeAddressLocationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeAddressLocation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddressLocation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddressLocationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressPoolDetailRequest() (request *DescribeAddressPoolDetailRequest) {
    request = &DescribeAddressPoolDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeAddressPoolDetail")
    
    
    return
}

func NewDescribeAddressPoolDetailResponse() (response *DescribeAddressPoolDetailResponse) {
    response = &DescribeAddressPoolDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddressPoolDetail
// 地址池详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_POOLUNAUTHORIZED = "UnauthorizedOperation.PoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeAddressPoolDetail(request *DescribeAddressPoolDetailRequest) (response *DescribeAddressPoolDetailResponse, err error) {
    return c.DescribeAddressPoolDetailWithContext(context.Background(), request)
}

// DescribeAddressPoolDetail
// 地址池详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_POOLUNAUTHORIZED = "UnauthorizedOperation.PoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeAddressPoolDetailWithContext(ctx context.Context, request *DescribeAddressPoolDetailRequest) (response *DescribeAddressPoolDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAddressPoolDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeAddressPoolDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddressPoolDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddressPoolDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressPoolListRequest() (request *DescribeAddressPoolListRequest) {
    request = &DescribeAddressPoolListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeAddressPoolList")
    
    
    return
}

func NewDescribeAddressPoolListResponse() (response *DescribeAddressPoolListResponse) {
    response = &DescribeAddressPoolListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAddressPoolList
// 地址池列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_QUERYNOTALLOW = "UnsupportedOperation.QueryNotAllow"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeAddressPoolList(request *DescribeAddressPoolListRequest) (response *DescribeAddressPoolListResponse, err error) {
    return c.DescribeAddressPoolListWithContext(context.Background(), request)
}

// DescribeAddressPoolList
// 地址池列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_QUERYNOTALLOW = "UnsupportedOperation.QueryNotAllow"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeAddressPoolListWithContext(ctx context.Context, request *DescribeAddressPoolListRequest) (response *DescribeAddressPoolListResponse, err error) {
    if request == nil {
        request = NewDescribeAddressPoolListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeAddressPoolList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAddressPoolList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAddressPoolListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetectPackageDetailRequest() (request *DescribeDetectPackageDetailRequest) {
    request = &DescribeDetectPackageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeDetectPackageDetail")
    
    
    return
}

func NewDescribeDetectPackageDetailResponse() (response *DescribeDetectPackageDetailResponse) {
    response = &DescribeDetectPackageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDetectPackageDetail
// 探测任务包详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDetectPackageDetail(request *DescribeDetectPackageDetailRequest) (response *DescribeDetectPackageDetailResponse, err error) {
    return c.DescribeDetectPackageDetailWithContext(context.Background(), request)
}

// DescribeDetectPackageDetail
// 探测任务包详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDetectPackageDetailWithContext(ctx context.Context, request *DescribeDetectPackageDetailRequest) (response *DescribeDetectPackageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDetectPackageDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeDetectPackageDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDetectPackageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDetectPackageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetectTaskPackageListRequest() (request *DescribeDetectTaskPackageListRequest) {
    request = &DescribeDetectTaskPackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeDetectTaskPackageList")
    
    
    return
}

func NewDescribeDetectTaskPackageListResponse() (response *DescribeDetectTaskPackageListResponse) {
    response = &DescribeDetectTaskPackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDetectTaskPackageList
// 探测任务套餐列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDetectTaskPackageList(request *DescribeDetectTaskPackageListRequest) (response *DescribeDetectTaskPackageListResponse, err error) {
    return c.DescribeDetectTaskPackageListWithContext(context.Background(), request)
}

// DescribeDetectTaskPackageList
// 探测任务套餐列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDetectTaskPackageListWithContext(ctx context.Context, request *DescribeDetectTaskPackageListRequest) (response *DescribeDetectTaskPackageListResponse, err error) {
    if request == nil {
        request = NewDescribeDetectTaskPackageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeDetectTaskPackageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDetectTaskPackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDetectTaskPackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetectorsRequest() (request *DescribeDetectorsRequest) {
    request = &DescribeDetectorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeDetectors")
    
    
    return
}

func NewDescribeDetectorsResponse() (response *DescribeDetectorsResponse) {
    response = &DescribeDetectorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDetectors
// 获取探测节点列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDetectors(request *DescribeDetectorsRequest) (response *DescribeDetectorsResponse, err error) {
    return c.DescribeDetectorsWithContext(context.Background(), request)
}

// DescribeDetectors
// 获取探测节点列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDetectorsWithContext(ctx context.Context, request *DescribeDetectorsRequest) (response *DescribeDetectorsResponse, err error) {
    if request == nil {
        request = NewDescribeDetectorsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeDetectors")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDetectors require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDetectorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnsLineListRequest() (request *DescribeDnsLineListRequest) {
    request = &DescribeDnsLineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeDnsLineList")
    
    
    return
}

func NewDescribeDnsLineListResponse() (response *DescribeDnsLineListResponse) {
    response = &DescribeDnsLineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDnsLineList
// 查询分组线路列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDnsLineList(request *DescribeDnsLineListRequest) (response *DescribeDnsLineListResponse, err error) {
    return c.DescribeDnsLineListWithContext(context.Background(), request)
}

// DescribeDnsLineList
// 查询分组线路列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeDnsLineListWithContext(ctx context.Context, request *DescribeDnsLineListRequest) (response *DescribeDnsLineListResponse, err error) {
    if request == nil {
        request = NewDescribeDnsLineListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeDnsLineList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnsLineList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnsLineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDetailRequest() (request *DescribeInstanceDetailRequest) {
    request = &DescribeInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeInstanceDetail")
    
    
    return
}

func NewDescribeInstanceDetailResponse() (response *DescribeInstanceDetailResponse) {
    response = &DescribeInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceDetail
// 实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_INNERREQUESTLIMITEXCEEDED = "RequestLimitExceeded.InnerRequestLimitExceeded"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeInstanceDetail(request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    return c.DescribeInstanceDetailWithContext(context.Background(), request)
}

// DescribeInstanceDetail
// 实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_INNERREQUESTLIMITEXCEEDED = "RequestLimitExceeded.InnerRequestLimitExceeded"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeInstanceDetailWithContext(ctx context.Context, request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
    request = &DescribeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeInstanceList")
    
    
    return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
    response = &DescribeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceList
// 实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_INNERREQUESTLIMITEXCEEDED = "RequestLimitExceeded.InnerRequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    return c.DescribeInstanceListWithContext(context.Background(), request)
}

// DescribeInstanceList
// 实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_INNERREQUESTLIMITEXCEEDED = "RequestLimitExceeded.InnerRequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancePackageListRequest() (request *DescribeInstancePackageListRequest) {
    request = &DescribeInstancePackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeInstancePackageList")
    
    
    return
}

func NewDescribeInstancePackageListResponse() (response *DescribeInstancePackageListResponse) {
    response = &DescribeInstancePackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancePackageList
// 实例套餐列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeInstancePackageList(request *DescribeInstancePackageListRequest) (response *DescribeInstancePackageListResponse, err error) {
    return c.DescribeInstancePackageListWithContext(context.Background(), request)
}

// DescribeInstancePackageList
// 实例套餐列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeInstancePackageListWithContext(ctx context.Context, request *DescribeInstancePackageListRequest) (response *DescribeInstancePackageListResponse, err error) {
    if request == nil {
        request = NewDescribeInstancePackageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeInstancePackageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancePackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancePackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorDetailRequest() (request *DescribeMonitorDetailRequest) {
    request = &DescribeMonitorDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeMonitorDetail")
    
    
    return
}

func NewDescribeMonitorDetailResponse() (response *DescribeMonitorDetailResponse) {
    response = &DescribeMonitorDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonitorDetail
// 查询监控器详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeMonitorDetail(request *DescribeMonitorDetailRequest) (response *DescribeMonitorDetailResponse, err error) {
    return c.DescribeMonitorDetailWithContext(context.Background(), request)
}

// DescribeMonitorDetail
// 查询监控器详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeMonitorDetailWithContext(ctx context.Context, request *DescribeMonitorDetailRequest) (response *DescribeMonitorDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeMonitorDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitorDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorsRequest() (request *DescribeMonitorsRequest) {
    request = &DescribeMonitorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeMonitors")
    
    
    return
}

func NewDescribeMonitorsResponse() (response *DescribeMonitorsResponse) {
    response = &DescribeMonitorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonitors
// 获取所有监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_QUERYNOTALLOW = "UnsupportedOperation.QueryNotAllow"
func (c *Client) DescribeMonitors(request *DescribeMonitorsRequest) (response *DescribeMonitorsResponse, err error) {
    return c.DescribeMonitorsWithContext(context.Background(), request)
}

// DescribeMonitors
// 获取所有监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_QUERYNOTALLOW = "UnsupportedOperation.QueryNotAllow"
func (c *Client) DescribeMonitorsWithContext(ctx context.Context, request *DescribeMonitorsRequest) (response *DescribeMonitorsResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeMonitors")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitors require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotasRequest() (request *DescribeQuotasRequest) {
    request = &DescribeQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeQuotas")
    
    
    return
}

func NewDescribeQuotasResponse() (response *DescribeQuotasResponse) {
    response = &DescribeQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuotas
// 配额查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeQuotas(request *DescribeQuotasRequest) (response *DescribeQuotasResponse, err error) {
    return c.DescribeQuotasWithContext(context.Background(), request)
}

// DescribeQuotas
// 配额查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeQuotasWithContext(ctx context.Context, request *DescribeQuotasRequest) (response *DescribeQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeQuotasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeQuotas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuotas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategyDetailRequest() (request *DescribeStrategyDetailRequest) {
    request = &DescribeStrategyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeStrategyDetail")
    
    
    return
}

func NewDescribeStrategyDetailResponse() (response *DescribeStrategyDetailResponse) {
    response = &DescribeStrategyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStrategyDetail
// 策略详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_STRATEGYUNAUTHORIZED = "UnauthorizedOperation.StrategyUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeStrategyDetail(request *DescribeStrategyDetailRequest) (response *DescribeStrategyDetailResponse, err error) {
    return c.DescribeStrategyDetailWithContext(context.Background(), request)
}

// DescribeStrategyDetail
// 策略详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_STRATEGYUNAUTHORIZED = "UnauthorizedOperation.StrategyUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeStrategyDetailWithContext(ctx context.Context, request *DescribeStrategyDetailRequest) (response *DescribeStrategyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeStrategyDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeStrategyDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStrategyListRequest() (request *DescribeStrategyListRequest) {
    request = &DescribeStrategyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "DescribeStrategyList")
    
    
    return
}

func NewDescribeStrategyListResponse() (response *DescribeStrategyListResponse) {
    response = &DescribeStrategyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStrategyList
// 策略列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_QUERYNOTALLOW = "UnsupportedOperation.QueryNotAllow"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_RESOURCESOURCEDUPLICATED = "UnsupportedOperation.ResourceSourceDuplicated"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeStrategyList(request *DescribeStrategyListRequest) (response *DescribeStrategyListResponse, err error) {
    return c.DescribeStrategyListWithContext(context.Background(), request)
}

// DescribeStrategyList
// 策略列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_QUERYNOTALLOW = "UnsupportedOperation.QueryNotAllow"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_RESOURCESOURCEDUPLICATED = "UnsupportedOperation.ResourceSourceDuplicated"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) DescribeStrategyListWithContext(ctx context.Context, request *DescribeStrategyListRequest) (response *DescribeStrategyListResponse, err error) {
    if request == nil {
        request = NewDescribeStrategyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "DescribeStrategyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategyListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressPoolRequest() (request *ModifyAddressPoolRequest) {
    request = &ModifyAddressPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "ModifyAddressPool")
    
    
    return
}

func NewModifyAddressPoolResponse() (response *ModifyAddressPoolResponse) {
    response = &ModifyAddressPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAddressPool
// 修改地址池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ADDREXCEEDEDLIMIT = "FailedOperation.AddrExceededLimit"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_EXCEEDEDMONITORNUMLIMIT = "FailedOperation.ExceededMonitorNumLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_MODIFYPOOLFAILED = "FailedOperation.ModifyPoolFailed"
//  FAILEDOPERATION_POOLEXCEEDEDLIMIT = "FailedOperation.PoolExceededLimit"
//  FAILEDOPERATION_TASKEXCEEDEDLIMIT = "FailedOperation.TaskExceededLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_POOLADDRDUPLICATED = "InvalidParameterValue.PoolAddrDuplicated"
//  INVALIDPARAMETERVALUE_POOLADDREMPTY = "InvalidParameterValue.PoolAddrEmpty"
//  INVALIDPARAMETERVALUE_POOLADDRINVALID = "InvalidParameterValue.PoolAddrInvalid"
//  INVALIDPARAMETERVALUE_POOLADDRMIXED = "InvalidParameterValue.PoolAddrMixed"
//  INVALIDPARAMETERVALUE_POOLADDRWEIGHTEMPTY = "InvalidParameterValue.PoolAddrWeightEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_POOLADDRESSUNAUTHORIZED = "UnauthorizedOperation.PoolAddressUnauthorized"
//  UNAUTHORIZEDOPERATION_POOLUNAUTHORIZED = "UnauthorizedOperation.PoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_POOLNAMEDUPLICATE = "UnsupportedOperation.PoolNameDuplicate"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyAddressPool(request *ModifyAddressPoolRequest) (response *ModifyAddressPoolResponse, err error) {
    return c.ModifyAddressPoolWithContext(context.Background(), request)
}

// ModifyAddressPool
// 修改地址池
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ADDREXCEEDEDLIMIT = "FailedOperation.AddrExceededLimit"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_EXCEEDEDMONITORNUMLIMIT = "FailedOperation.ExceededMonitorNumLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_MODIFYPOOLFAILED = "FailedOperation.ModifyPoolFailed"
//  FAILEDOPERATION_POOLEXCEEDEDLIMIT = "FailedOperation.PoolExceededLimit"
//  FAILEDOPERATION_TASKEXCEEDEDLIMIT = "FailedOperation.TaskExceededLimit"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_POOLADDRDUPLICATED = "InvalidParameterValue.PoolAddrDuplicated"
//  INVALIDPARAMETERVALUE_POOLADDREMPTY = "InvalidParameterValue.PoolAddrEmpty"
//  INVALIDPARAMETERVALUE_POOLADDRINVALID = "InvalidParameterValue.PoolAddrInvalid"
//  INVALIDPARAMETERVALUE_POOLADDRMIXED = "InvalidParameterValue.PoolAddrMixed"
//  INVALIDPARAMETERVALUE_POOLADDRWEIGHTEMPTY = "InvalidParameterValue.PoolAddrWeightEmpty"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_POOLADDRESSUNAUTHORIZED = "UnauthorizedOperation.PoolAddressUnauthorized"
//  UNAUTHORIZEDOPERATION_POOLUNAUTHORIZED = "UnauthorizedOperation.PoolUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_POOLNAMEDUPLICATE = "UnsupportedOperation.PoolNameDuplicate"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyAddressPoolWithContext(ctx context.Context, request *ModifyAddressPoolRequest) (response *ModifyAddressPoolResponse, err error) {
    if request == nil {
        request = NewModifyAddressPoolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "ModifyAddressPool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAddressPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAddressPoolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceConfigRequest() (request *ModifyInstanceConfigRequest) {
    request = &ModifyInstanceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "ModifyInstanceConfig")
    
    
    return
}

func NewModifyInstanceConfigResponse() (response *ModifyInstanceConfigResponse) {
    response = &ModifyInstanceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceConfig
// 修改实例配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCESSTYPEINVALID = "InvalidParameter.AccessTypeInvalid"
//  INVALIDPARAMETER_GLOBALACCESSDOMAINCONFLICT = "InvalidParameter.GlobalAccessDomainConflict"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_ACCESSDOMAINUNAUTHORIZED = "UnsupportedOperation.AccessDomainUnauthorized"
//  UNSUPPORTEDOPERATION_DOMAINTTLINVALID = "UnsupportedOperation.DomainTtlInvalid"
//  UNSUPPORTEDOPERATION_INSTANCENAMEDUPLICATE = "UnsupportedOperation.InstanceNameDuplicate"
func (c *Client) ModifyInstanceConfig(request *ModifyInstanceConfigRequest) (response *ModifyInstanceConfigResponse, err error) {
    return c.ModifyInstanceConfigWithContext(context.Background(), request)
}

// ModifyInstanceConfig
// 修改实例配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCESSTYPEINVALID = "InvalidParameter.AccessTypeInvalid"
//  INVALIDPARAMETER_GLOBALACCESSDOMAINCONFLICT = "InvalidParameter.GlobalAccessDomainConflict"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_ACCESSDOMAINUNAUTHORIZED = "UnsupportedOperation.AccessDomainUnauthorized"
//  UNSUPPORTEDOPERATION_DOMAINTTLINVALID = "UnsupportedOperation.DomainTtlInvalid"
//  UNSUPPORTEDOPERATION_INSTANCENAMEDUPLICATE = "UnsupportedOperation.InstanceNameDuplicate"
func (c *Client) ModifyInstanceConfigWithContext(ctx context.Context, request *ModifyInstanceConfigRequest) (response *ModifyInstanceConfigResponse, err error) {
    if request == nil {
        request = NewModifyInstanceConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "ModifyInstanceConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMonitorRequest() (request *ModifyMonitorRequest) {
    request = &ModifyMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "ModifyMonitor")
    
    
    return
}

func NewModifyMonitorResponse() (response *ModifyMonitorResponse) {
    response = &ModifyMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMonitor
// 修改监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_EXCEEDEDMONITORNUMLIMIT = "FailedOperation.ExceededMonitorNumLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_MODIFYMONITORFAILED = "FailedOperation.ModifyMonitorFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_MONITORNAMEEXIST = "UnsupportedOperation.MonitorNameExist"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyMonitor(request *ModifyMonitorRequest) (response *ModifyMonitorResponse, err error) {
    return c.ModifyMonitorWithContext(context.Background(), request)
}

// ModifyMonitor
// 修改监控器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_EXCEEDEDMONITORNUMLIMIT = "FailedOperation.ExceededMonitorNumLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_MODIFYMONITORFAILED = "FailedOperation.ModifyMonitorFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_MONITORNAMEEXIST = "UnsupportedOperation.MonitorNameExist"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyMonitorWithContext(ctx context.Context, request *ModifyMonitorRequest) (response *ModifyMonitorResponse, err error) {
    if request == nil {
        request = NewModifyMonitorRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "ModifyMonitor")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMonitorResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPackageAutoRenewRequest() (request *ModifyPackageAutoRenewRequest) {
    request = &ModifyPackageAutoRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "ModifyPackageAutoRenew")
    
    
    return
}

func NewModifyPackageAutoRenewResponse() (response *ModifyPackageAutoRenewResponse) {
    response = &ModifyPackageAutoRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPackageAutoRenew
// 设置自动续费接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTORENEWFAILED = "FailedOperation.AutoRenewFailed"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyPackageAutoRenew(request *ModifyPackageAutoRenewRequest) (response *ModifyPackageAutoRenewResponse, err error) {
    return c.ModifyPackageAutoRenewWithContext(context.Background(), request)
}

// ModifyPackageAutoRenew
// 设置自动续费接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTORENEWFAILED = "FailedOperation.AutoRenewFailed"
//  FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"
//  FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyPackageAutoRenewWithContext(ctx context.Context, request *ModifyPackageAutoRenewRequest) (response *ModifyPackageAutoRenewResponse, err error) {
    if request == nil {
        request = NewModifyPackageAutoRenewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "ModifyPackageAutoRenew")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPackageAutoRenew require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPackageAutoRenewResponse()
    err = c.Send(request, response)
    return
}

func NewModifyStrategyRequest() (request *ModifyStrategyRequest) {
    request = &ModifyStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("igtm", APIVersion, "ModifyStrategy")
    
    
    return
}

func NewModifyStrategyResponse() (response *ModifyStrategyResponse) {
    response = &ModifyStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyStrategy
// 修改策略接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTDELETEDEFAULTLINEERROR = "FailedOperation.CanNotDeleteDefaultLineError"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_MODIFYSTRATEGYFAILED = "FailedOperation.ModifyStrategyFailed"
//  FAILEDOPERATION_POOLWASUSINGERROR = "FailedOperation.PoolWasUsingError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETER_TRAFFICSTRATEGYINVALID = "InvalidParameter.TrafficStrategyInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TRAFFICSTRATEGYWEIGHTEMPTY = "InvalidParameterValue.TrafficStrategyWeightEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_NOINSTANCE = "ResourceInsufficient.NoInstance"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_INSTANCEUNAUTHORIZED = "UnauthorizedOperation.InstanceUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINBINDPOOLUNAUTHORIZED = "UnauthorizedOperation.MainBindPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINPOOLUNAUTHORIZED = "UnauthorizedOperation.MainPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_STRATEGYUNAUTHORIZED = "UnauthorizedOperation.StrategyUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FALLBACKPOOLINVALID = "UnsupportedOperation.FallBackPoolInvalid"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLDUPLICATE = "UnsupportedOperation.MainBindPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLEMPTY = "UnsupportedOperation.MainBindPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLLEVELEXCEED = "UnsupportedOperation.MainBindPoolLevelExceed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLMIXED = "UnsupportedOperation.MainBindPoolMixed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLNUMINVALID = "UnsupportedOperation.MainBindPoolNumInvalid"
//  UNSUPPORTEDOPERATION_MAINPOOLDUPLICATE = "UnsupportedOperation.MainPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINPOOLEMPTY = "UnsupportedOperation.MainPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINPOOLSURVIVENUMINVALID = "UnsupportedOperation.MainPoolSurviveNumInvalid"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_RESOURCEPOOLSOURCENOTEXIST = "UnsupportedOperation.ResourcePoolSourceNotExist"
//  UNSUPPORTEDOPERATION_RESOURCESOURCEDUPLICATED = "UnsupportedOperation.ResourceSourceDuplicated"
//  UNSUPPORTEDOPERATION_STRATEGYNAMEDUPLICATE = "UnsupportedOperation.StrategyNameDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEDUPLICATE = "UnsupportedOperation.StrategySourceDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEINUSED = "UnsupportedOperation.StrategySourceInUsed"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEUNAUTHORIZED = "UnsupportedOperation.StrategySourceUnauthorized"
//  UNSUPPORTEDOPERATION_STRATEGYSWITCHTYPEINVALID = "UnsupportedOperation.StrategySwitchTypeInvalid"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyStrategy(request *ModifyStrategyRequest) (response *ModifyStrategyResponse, err error) {
    return c.ModifyStrategyWithContext(context.Background(), request)
}

// ModifyStrategy
// 修改策略接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CANNOTDELETEDEFAULTLINEERROR = "FailedOperation.CanNotDeleteDefaultLineError"
//  FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"
//  FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"
//  FAILEDOPERATION_MODIFYSTRATEGYFAILED = "FailedOperation.ModifyStrategyFailed"
//  FAILEDOPERATION_POOLWASUSINGERROR = "FailedOperation.PoolWasUsingError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"
//  INVALIDPARAMETER_TRAFFICSTRATEGYINVALID = "InvalidParameter.TrafficStrategyInvalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TRAFFICSTRATEGYWEIGHTEMPTY = "InvalidParameterValue.TrafficStrategyWeightEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_NOINSTANCE = "ResourceInsufficient.NoInstance"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_INSTANCEUNAUTHORIZED = "UnauthorizedOperation.InstanceUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINBINDPOOLUNAUTHORIZED = "UnauthorizedOperation.MainBindPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_MAINPOOLUNAUTHORIZED = "UnauthorizedOperation.MainPoolUnauthorized"
//  UNAUTHORIZEDOPERATION_STRATEGYUNAUTHORIZED = "UnauthorizedOperation.StrategyUnauthorized"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FALLBACKPOOLINVALID = "UnsupportedOperation.FallBackPoolInvalid"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLDUPLICATE = "UnsupportedOperation.MainBindPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLEMPTY = "UnsupportedOperation.MainBindPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLLEVELEXCEED = "UnsupportedOperation.MainBindPoolLevelExceed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLMIXED = "UnsupportedOperation.MainBindPoolMixed"
//  UNSUPPORTEDOPERATION_MAINBINDPOOLNUMINVALID = "UnsupportedOperation.MainBindPoolNumInvalid"
//  UNSUPPORTEDOPERATION_MAINPOOLDUPLICATE = "UnsupportedOperation.MainPoolDuplicate"
//  UNSUPPORTEDOPERATION_MAINPOOLEMPTY = "UnsupportedOperation.MainPoolEmpty"
//  UNSUPPORTEDOPERATION_MAINPOOLSURVIVENUMINVALID = "UnsupportedOperation.MainPoolSurviveNumInvalid"
//  UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"
//  UNSUPPORTEDOPERATION_RESOURCEPOOLSOURCENOTEXIST = "UnsupportedOperation.ResourcePoolSourceNotExist"
//  UNSUPPORTEDOPERATION_RESOURCESOURCEDUPLICATED = "UnsupportedOperation.ResourceSourceDuplicated"
//  UNSUPPORTEDOPERATION_STRATEGYNAMEDUPLICATE = "UnsupportedOperation.StrategyNameDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEDUPLICATE = "UnsupportedOperation.StrategySourceDuplicate"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEINUSED = "UnsupportedOperation.StrategySourceInUsed"
//  UNSUPPORTEDOPERATION_STRATEGYSOURCEUNAUTHORIZED = "UnsupportedOperation.StrategySourceUnauthorized"
//  UNSUPPORTEDOPERATION_STRATEGYSWITCHTYPEINVALID = "UnsupportedOperation.StrategySwitchTypeInvalid"
//  UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
func (c *Client) ModifyStrategyWithContext(ctx context.Context, request *ModifyStrategyRequest) (response *ModifyStrategyResponse, err error) {
    if request == nil {
        request = NewModifyStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "igtm", APIVersion, "ModifyStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyStrategyResponse()
    err = c.Send(request, response)
    return
}
