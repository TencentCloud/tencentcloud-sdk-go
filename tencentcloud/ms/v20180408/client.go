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

package v20180408

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-08"

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


func NewCreateBindInstanceRequest() (request *CreateBindInstanceRequest) {
    request = &CreateBindInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "CreateBindInstance")
    
    
    return
}

func NewCreateBindInstanceResponse() (response *CreateBindInstanceResponse) {
    response = &CreateBindInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBindInstance
// 将应用和资源进行绑定。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateBindInstance(request *CreateBindInstanceRequest) (response *CreateBindInstanceResponse, err error) {
    return c.CreateBindInstanceWithContext(context.Background(), request)
}

// CreateBindInstance
// 将应用和资源进行绑定。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateBindInstanceWithContext(ctx context.Context, request *CreateBindInstanceRequest) (response *CreateBindInstanceResponse, err error) {
    if request == nil {
        request = NewCreateBindInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBindInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBindInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosSecKeyInstanceRequest() (request *CreateCosSecKeyInstanceRequest) {
    request = &CreateCosSecKeyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "CreateCosSecKeyInstance")
    
    
    return
}

func NewCreateCosSecKeyInstanceResponse() (response *CreateCosSecKeyInstanceResponse) {
    response = &CreateCosSecKeyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCosSecKeyInstance
// 获取云COS文件存储临时密钥，密钥仅限于临时上传文件，有访问限制和时效性，请保管好临时密钥。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateCosSecKeyInstance(request *CreateCosSecKeyInstanceRequest) (response *CreateCosSecKeyInstanceResponse, err error) {
    return c.CreateCosSecKeyInstanceWithContext(context.Background(), request)
}

// CreateCosSecKeyInstance
// 获取云COS文件存储临时密钥，密钥仅限于临时上传文件，有访问限制和时效性，请保管好临时密钥。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateCosSecKeyInstanceWithContext(ctx context.Context, request *CreateCosSecKeyInstanceRequest) (response *CreateCosSecKeyInstanceResponse, err error) {
    if request == nil {
        request = NewCreateCosSecKeyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosSecKeyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosSecKeyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceInstancesRequest() (request *CreateResourceInstancesRequest) {
    request = &CreateResourceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "CreateResourceInstances")
    
    
    return
}

func NewCreateResourceInstancesResponse() (response *CreateResourceInstancesResponse) {
    response = &CreateResourceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateResourceInstances
// 用户可以使用该接口自建资源，只支持白名单用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateResourceInstances(request *CreateResourceInstancesRequest) (response *CreateResourceInstancesResponse, err error) {
    return c.CreateResourceInstancesWithContext(context.Background(), request)
}

// CreateResourceInstances
// 用户可以使用该接口自建资源，只支持白名单用户
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateResourceInstancesWithContext(ctx context.Context, request *CreateResourceInstancesRequest) (response *CreateResourceInstancesResponse, err error) {
    if request == nil {
        request = NewCreateResourceInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourceInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScanInstancesRequest() (request *CreateScanInstancesRequest) {
    request = &CreateScanInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "CreateScanInstances")
    
    
    return
}

func NewCreateScanInstancesResponse() (response *CreateScanInstancesResponse) {
    response = &CreateScanInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScanInstances
// 由于该产品是线上免费使用产品，无企业版用户，升级迭代成本高及人力安排等原因，安全测评产品不再接入新用户，故下线。
//
// 
//
// 用户通过该接口批量提交应用进行应用扫描，扫描后需通过DescribeScanResults接口查询扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGAPPINFO = "MissingParameter.MissingAppInfo"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateScanInstances(request *CreateScanInstancesRequest) (response *CreateScanInstancesResponse, err error) {
    return c.CreateScanInstancesWithContext(context.Background(), request)
}

// CreateScanInstances
// 由于该产品是线上免费使用产品，无企业版用户，升级迭代成本高及人力安排等原因，安全测评产品不再接入新用户，故下线。
//
// 
//
// 用户通过该接口批量提交应用进行应用扫描，扫描后需通过DescribeScanResults接口查询扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGAPPINFO = "MissingParameter.MissingAppInfo"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateScanInstancesWithContext(ctx context.Context, request *CreateScanInstancesRequest) (response *CreateScanInstancesResponse, err error) {
    if request == nil {
        request = NewCreateScanInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScanInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScanInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateShieldInstanceRequest() (request *CreateShieldInstanceRequest) {
    request = &CreateShieldInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "CreateShieldInstance")
    
    
    return
}

func NewCreateShieldInstanceResponse() (response *CreateShieldInstanceResponse) {
    response = &CreateShieldInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateShieldInstance
// 用户通过该接口提交应用进行应用加固，加固后需通过DescribeShieldResult接口查询加固结果。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER_MISSINGSERVICEINFO = "InvalidParameter.MissingServiceInfo"
//  INVALIDPARAMETER_PLANIDNOTFOUND = "InvalidParameter.PlanIdNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGAPPINFO = "MissingParameter.MissingAppInfo"
//  RESOURCENOTFOUND_PLANIDNOTFOUND = "ResourceNotFound.PlanIdNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateShieldInstance(request *CreateShieldInstanceRequest) (response *CreateShieldInstanceResponse, err error) {
    return c.CreateShieldInstanceWithContext(context.Background(), request)
}

// CreateShieldInstance
// 用户通过该接口提交应用进行应用加固，加固后需通过DescribeShieldResult接口查询加固结果。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER_MISSINGSERVICEINFO = "InvalidParameter.MissingServiceInfo"
//  INVALIDPARAMETER_PLANIDNOTFOUND = "InvalidParameter.PlanIdNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGAPPINFO = "MissingParameter.MissingAppInfo"
//  RESOURCENOTFOUND_PLANIDNOTFOUND = "ResourceNotFound.PlanIdNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateShieldInstanceWithContext(ctx context.Context, request *CreateShieldInstanceRequest) (response *CreateShieldInstanceResponse, err error) {
    if request == nil {
        request = NewCreateShieldInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateShieldInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateShieldInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateShieldPlanInstanceRequest() (request *CreateShieldPlanInstanceRequest) {
    request = &CreateShieldPlanInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "CreateShieldPlanInstance")
    
    
    return
}

func NewCreateShieldPlanInstanceResponse() (response *CreateShieldPlanInstanceResponse) {
    response = &CreateShieldPlanInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateShieldPlanInstance
// 对资源进行策略新增。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateShieldPlanInstance(request *CreateShieldPlanInstanceRequest) (response *CreateShieldPlanInstanceResponse, err error) {
    return c.CreateShieldPlanInstanceWithContext(context.Background(), request)
}

// CreateShieldPlanInstance
// 对资源进行策略新增。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) CreateShieldPlanInstanceWithContext(ctx context.Context, request *CreateShieldPlanInstanceRequest) (response *CreateShieldPlanInstanceResponse, err error) {
    if request == nil {
        request = NewCreateShieldPlanInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateShieldPlanInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateShieldPlanInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScanInstancesRequest() (request *DeleteScanInstancesRequest) {
    request = &DeleteScanInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DeleteScanInstances")
    
    
    return
}

func NewDeleteScanInstancesResponse() (response *DeleteScanInstancesResponse) {
    response = &DeleteScanInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScanInstances
// 删除一个或者多个app扫描信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteScanInstances(request *DeleteScanInstancesRequest) (response *DeleteScanInstancesResponse, err error) {
    return c.DeleteScanInstancesWithContext(context.Background(), request)
}

// DeleteScanInstances
// 删除一个或者多个app扫描信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteScanInstancesWithContext(ctx context.Context, request *DeleteScanInstancesRequest) (response *DeleteScanInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteScanInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScanInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScanInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShieldInstancesRequest() (request *DeleteShieldInstancesRequest) {
    request = &DeleteShieldInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DeleteShieldInstances")
    
    
    return
}

func NewDeleteShieldInstancesResponse() (response *DeleteShieldInstancesResponse) {
    response = &DeleteShieldInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteShieldInstances
// 删除一个或者多个app加固信息。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETERVALUE_INVALIDITEMIDS = "InvalidParameterValue.InvalidItemIds"
//  MISSINGPARAMETER_MISSINGITEMIDS = "MissingParameter.MissingItemIds"
//  RESOURCENOTFOUND_ITEMIDNOTFOUND = "ResourceNotFound.ItemIdNotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DeleteShieldInstances(request *DeleteShieldInstancesRequest) (response *DeleteShieldInstancesResponse, err error) {
    return c.DeleteShieldInstancesWithContext(context.Background(), request)
}

// DeleteShieldInstances
// 删除一个或者多个app加固信息。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETERVALUE_INVALIDITEMIDS = "InvalidParameterValue.InvalidItemIds"
//  MISSINGPARAMETER_MISSINGITEMIDS = "MissingParameter.MissingItemIds"
//  RESOURCENOTFOUND_ITEMIDNOTFOUND = "ResourceNotFound.ItemIdNotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DeleteShieldInstancesWithContext(ctx context.Context, request *DeleteShieldInstancesRequest) (response *DeleteShieldInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteShieldInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShieldInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShieldInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApkDetectionResultRequest() (request *DescribeApkDetectionResultRequest) {
    request = &DescribeApkDetectionResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeApkDetectionResult")
    
    
    return
}

func NewDescribeApkDetectionResultResponse() (response *DescribeApkDetectionResultResponse) {
    response = &DescribeApkDetectionResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApkDetectionResult
// 该接口采用同步模式请求腾讯APK云检测服务，即时返回检测数据，需要用户用轮询的方式调用本接口来进行样本送检并获取检测结果(每隔60s发送一次请求，传相同的参数，重试30次)，一般情况下0.5h内会出检测结果，最长时间是3h。当Result为ok并且ResultList数组非空有值时，代表检测完毕，若长时间获取不到检测结果，请联系客服。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APKSERVERERROR = "InternalError.ApkServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApkDetectionResult(request *DescribeApkDetectionResultRequest) (response *DescribeApkDetectionResultResponse, err error) {
    return c.DescribeApkDetectionResultWithContext(context.Background(), request)
}

// DescribeApkDetectionResult
// 该接口采用同步模式请求腾讯APK云检测服务，即时返回检测数据，需要用户用轮询的方式调用本接口来进行样本送检并获取检测结果(每隔60s发送一次请求，传相同的参数，重试30次)，一般情况下0.5h内会出检测结果，最长时间是3h。当Result为ok并且ResultList数组非空有值时，代表检测完毕，若长时间获取不到检测结果，请联系客服。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APKSERVERERROR = "InternalError.ApkServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApkDetectionResultWithContext(ctx context.Context, request *DescribeApkDetectionResultRequest) (response *DescribeApkDetectionResultResponse, err error) {
    if request == nil {
        request = NewDescribeApkDetectionResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApkDetectionResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApkDetectionResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceInstancesRequest() (request *DescribeResourceInstancesRequest) {
    request = &DescribeResourceInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeResourceInstances")
    
    
    return
}

func NewDescribeResourceInstancesResponse() (response *DescribeResourceInstancesResponse) {
    response = &DescribeResourceInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceInstances
// 获取某个用户的所有资源信息。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDOFFSET = "InvalidParameterValue.InvalidOffset"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DescribeResourceInstances(request *DescribeResourceInstancesRequest) (response *DescribeResourceInstancesResponse, err error) {
    return c.DescribeResourceInstancesWithContext(context.Background(), request)
}

// DescribeResourceInstances
// 获取某个用户的所有资源信息。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDOFFSET = "InvalidParameterValue.InvalidOffset"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DescribeResourceInstancesWithContext(ctx context.Context, request *DescribeResourceInstancesRequest) (response *DescribeResourceInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeResourceInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanInstancesRequest() (request *DescribeScanInstancesRequest) {
    request = &DescribeScanInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeScanInstances")
    
    
    return
}

func NewDescribeScanInstancesResponse() (response *DescribeScanInstancesResponse) {
    response = &DescribeScanInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanInstances
// 由于该产品是线上免费使用产品，无企业版用户，升级迭代成本高及人力安排等原因，安全测评产品不再接入新用户，故下线。
//
// 
//
// 本接口用于查看app列表。
//
// 可以通过指定任务唯一标识ItemId来查询指定app的详细信息，或通过设定过滤器来查询满足过滤条件的app的详细信息。 指定偏移(Offset)和限制(Limit)来选择结果中的一部分，默认返回满足条件的前20个app信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDITEMIDS = "InvalidParameterValue.InvalidItemIds"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDOFFSET = "InvalidParameterValue.InvalidOffset"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanInstances(request *DescribeScanInstancesRequest) (response *DescribeScanInstancesResponse, err error) {
    return c.DescribeScanInstancesWithContext(context.Background(), request)
}

// DescribeScanInstances
// 由于该产品是线上免费使用产品，无企业版用户，升级迭代成本高及人力安排等原因，安全测评产品不再接入新用户，故下线。
//
// 
//
// 本接口用于查看app列表。
//
// 可以通过指定任务唯一标识ItemId来查询指定app的详细信息，或通过设定过滤器来查询满足过滤条件的app的详细信息。 指定偏移(Offset)和限制(Limit)来选择结果中的一部分，默认返回满足条件的前20个app信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDITEMIDS = "InvalidParameterValue.InvalidItemIds"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDOFFSET = "InvalidParameterValue.InvalidOffset"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanInstancesWithContext(ctx context.Context, request *DescribeScanInstancesRequest) (response *DescribeScanInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeScanInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanResultsRequest() (request *DescribeScanResultsRequest) {
    request = &DescribeScanResultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeScanResults")
    
    
    return
}

func NewDescribeScanResultsResponse() (response *DescribeScanResultsResponse) {
    response = &DescribeScanResultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanResults
// 由于该产品是线上免费使用产品，无企业版用户，升级迭代成本高及人力安排等原因，安全测评产品不再接入新用户，故下线。
//
// 
//
// 用户通过CreateScanInstances接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞信息，广告信息，插件信息和病毒信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGITEMID = "MissingParameter.MissingItemId"
//  RESOURCENOTFOUND_ITEMIDNOTFOUND = "ResourceNotFound.ItemIdNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResults(request *DescribeScanResultsRequest) (response *DescribeScanResultsResponse, err error) {
    return c.DescribeScanResultsWithContext(context.Background(), request)
}

// DescribeScanResults
// 由于该产品是线上免费使用产品，无企业版用户，升级迭代成本高及人力安排等原因，安全测评产品不再接入新用户，故下线。
//
// 
//
// 用户通过CreateScanInstances接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞信息，广告信息，插件信息和病毒信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGITEMID = "MissingParameter.MissingItemId"
//  RESOURCENOTFOUND_ITEMIDNOTFOUND = "ResourceNotFound.ItemIdNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResultsWithContext(ctx context.Context, request *DescribeScanResultsRequest) (response *DescribeScanResultsResponse, err error) {
    if request == nil {
        request = NewDescribeScanResultsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanResults require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanResultsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShieldInstancesRequest() (request *DescribeShieldInstancesRequest) {
    request = &DescribeShieldInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeShieldInstances")
    
    
    return
}

func NewDescribeShieldInstancesResponse() (response *DescribeShieldInstancesResponse) {
    response = &DescribeShieldInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeShieldInstances
// 本接口用于查看app列表。
//
// 可以通过指定任务唯一标识ItemId来查询指定app的详细信息，或通过设定过滤器来查询满足过滤条件的app的详细信息。 指定偏移(Offset)和限制(Limit)来选择结果中的一部分，默认返回满足条件的前20个app信息。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETERVALUE_INVALIDCOEXISTITEMIDSFILTERS = "InvalidParameterValue.InvalidCoexistItemIdsFilters"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDITEMIDS = "InvalidParameterValue.InvalidItemIds"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDOFFSET = "InvalidParameterValue.InvalidOffset"
//  INVALIDPARAMETERVALUE_INVALIDORDERDIRECTION = "InvalidParameterValue.InvalidOrderDirection"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DescribeShieldInstances(request *DescribeShieldInstancesRequest) (response *DescribeShieldInstancesResponse, err error) {
    return c.DescribeShieldInstancesWithContext(context.Background(), request)
}

// DescribeShieldInstances
// 本接口用于查看app列表。
//
// 可以通过指定任务唯一标识ItemId来查询指定app的详细信息，或通过设定过滤器来查询满足过滤条件的app的详细信息。 指定偏移(Offset)和限制(Limit)来选择结果中的一部分，默认返回满足条件的前20个app信息。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETERVALUE_INVALIDCOEXISTITEMIDSFILTERS = "InvalidParameterValue.InvalidCoexistItemIdsFilters"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDITEMIDS = "InvalidParameterValue.InvalidItemIds"
//  INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"
//  INVALIDPARAMETERVALUE_INVALIDOFFSET = "InvalidParameterValue.InvalidOffset"
//  INVALIDPARAMETERVALUE_INVALIDORDERDIRECTION = "InvalidParameterValue.InvalidOrderDirection"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DescribeShieldInstancesWithContext(ctx context.Context, request *DescribeShieldInstancesRequest) (response *DescribeShieldInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeShieldInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShieldInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShieldInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShieldPlanInstanceRequest() (request *DescribeShieldPlanInstanceRequest) {
    request = &DescribeShieldPlanInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeShieldPlanInstance")
    
    
    return
}

func NewDescribeShieldPlanInstanceResponse() (response *DescribeShieldPlanInstanceResponse) {
    response = &DescribeShieldPlanInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeShieldPlanInstance
// 查询加固策略。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCEUNAVAILABLE_NOTBIND = "ResourceUnavailable.NotBind"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DescribeShieldPlanInstance(request *DescribeShieldPlanInstanceRequest) (response *DescribeShieldPlanInstanceResponse, err error) {
    return c.DescribeShieldPlanInstanceWithContext(context.Background(), request)
}

// DescribeShieldPlanInstance
// 查询加固策略。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  RESOURCEUNAVAILABLE_NOTBIND = "ResourceUnavailable.NotBind"
//  RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
func (c *Client) DescribeShieldPlanInstanceWithContext(ctx context.Context, request *DescribeShieldPlanInstanceRequest) (response *DescribeShieldPlanInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeShieldPlanInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShieldPlanInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShieldPlanInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShieldResultRequest() (request *DescribeShieldResultRequest) {
    request = &DescribeShieldResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeShieldResult")
    
    
    return
}

func NewDescribeShieldResultResponse() (response *DescribeShieldResultResponse) {
    response = &DescribeShieldResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeShieldResult
// 通过唯一标识获取加固的结果。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  MISSINGPARAMETER_MISSINGITEMID = "MissingParameter.MissingItemId"
//  RESOURCENOTFOUND_ITEMIDNOTFOUND = "ResourceNotFound.ItemIdNotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShieldResult(request *DescribeShieldResultRequest) (response *DescribeShieldResultResponse, err error) {
    return c.DescribeShieldResultWithContext(context.Background(), request)
}

// DescribeShieldResult
// 通过唯一标识获取加固的结果。（注意：根据国家互联网用户实名制相关要求，使用该产品前，需先完成实名认证。）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  MISSINGPARAMETER_MISSINGITEMID = "MissingParameter.MissingItemId"
//  RESOURCENOTFOUND_ITEMIDNOTFOUND = "ResourceNotFound.ItemIdNotFound"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShieldResultWithContext(ctx context.Context, request *DescribeShieldResultRequest) (response *DescribeShieldResultResponse, err error) {
    if request == nil {
        request = NewDescribeShieldResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShieldResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShieldResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUrlDetectionResultRequest() (request *DescribeUrlDetectionResultRequest) {
    request = &DescribeUrlDetectionResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeUrlDetectionResult")
    
    
    return
}

func NewDescribeUrlDetectionResultResponse() (response *DescribeUrlDetectionResultResponse) {
    response = &DescribeUrlDetectionResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUrlDetectionResult
// 移动安全-网址检测服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUrlDetectionResult(request *DescribeUrlDetectionResultRequest) (response *DescribeUrlDetectionResultResponse, err error) {
    return c.DescribeUrlDetectionResultWithContext(context.Background(), request)
}

// DescribeUrlDetectionResult
// 移动安全-网址检测服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUrlDetectionResultWithContext(ctx context.Context, request *DescribeUrlDetectionResultRequest) (response *DescribeUrlDetectionResultResponse, err error) {
    if request == nil {
        request = NewDescribeUrlDetectionResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUrlDetectionResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUrlDetectionResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserBaseInfoInstanceRequest() (request *DescribeUserBaseInfoInstanceRequest) {
    request = &DescribeUserBaseInfoInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ms", APIVersion, "DescribeUserBaseInfoInstance")
    
    
    return
}

func NewDescribeUserBaseInfoInstanceResponse() (response *DescribeUserBaseInfoInstanceResponse) {
    response = &DescribeUserBaseInfoInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserBaseInfoInstance
// 获取用户基础信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserBaseInfoInstance(request *DescribeUserBaseInfoInstanceRequest) (response *DescribeUserBaseInfoInstanceResponse, err error) {
    return c.DescribeUserBaseInfoInstanceWithContext(context.Background(), request)
}

// DescribeUserBaseInfoInstance
// 获取用户基础信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserBaseInfoInstanceWithContext(ctx context.Context, request *DescribeUserBaseInfoInstanceRequest) (response *DescribeUserBaseInfoInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeUserBaseInfoInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserBaseInfoInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserBaseInfoInstanceResponse()
    err = c.Send(request, response)
    return
}
