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

package v20200727

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-27"

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


func NewDescribeInvocationResultRequest() (request *DescribeInvocationResultRequest) {
    request = &DescribeInvocationResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcex", APIVersion, "DescribeInvocationResult")
    
    
    return
}

func NewDescribeInvocationResultResponse() (response *DescribeInvocationResultResponse) {
    response = &DescribeInvocationResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInvocationResult
// 产品控制台已经下线
//
// 
//
// 获取服务调用结果。和InvokeService接口配置合适，其InvokeId参数为InvokeService接口返回的RequestId。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEFINDERROR = "FailedOperation.DatabaseFindError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETER_REQUESTPARSEERROR = "InvalidParameter.RequestParseError"
func (c *Client) DescribeInvocationResult(request *DescribeInvocationResultRequest) (response *DescribeInvocationResultResponse, err error) {
    return c.DescribeInvocationResultWithContext(context.Background(), request)
}

// DescribeInvocationResult
// 产品控制台已经下线
//
// 
//
// 获取服务调用结果。和InvokeService接口配置合适，其InvokeId参数为InvokeService接口返回的RequestId。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEFINDERROR = "FailedOperation.DatabaseFindError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETER_REQUESTPARSEERROR = "InvalidParameter.RequestParseError"
func (c *Client) DescribeInvocationResultWithContext(ctx context.Context, request *DescribeInvocationResultRequest) (response *DescribeInvocationResultResponse, err error) {
    if request == nil {
        request = NewDescribeInvocationResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInvocationResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInvocationResultResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeServiceRequest() (request *InvokeServiceRequest) {
    request = &InvokeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcex", APIVersion, "InvokeService")
    
    
    return
}

func NewInvokeServiceResponse() (response *InvokeServiceResponse) {
    response = &InvokeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InvokeService
// 产品控制台已经下线
//
// 
//
// 通过传入文档url，测试服务算法。此接口需要和DescribeInvocationResult接口配置使用，该接口使用InvokeService返回的RequestId作为InvokeId参数，用于查询调用结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILENOTEXISTS = "FailedOperation.FileNotExists"
//  FAILEDOPERATION_INSUFFICIENTPRIVILEGE = "FailedOperation.InsufficientPrivilege"
//  FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMETERS = "InvalidParameter.EmptyParameters"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILESIZE = "InvalidParameter.InvalidFileSize"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETER_REQUESTPARSEERROR = "InvalidParameter.RequestParseError"
func (c *Client) InvokeService(request *InvokeServiceRequest) (response *InvokeServiceResponse, err error) {
    return c.InvokeServiceWithContext(context.Background(), request)
}

// InvokeService
// 产品控制台已经下线
//
// 
//
// 通过传入文档url，测试服务算法。此接口需要和DescribeInvocationResult接口配置使用，该接口使用InvokeService返回的RequestId作为InvokeId参数，用于查询调用结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILENOTEXISTS = "FailedOperation.FileNotExists"
//  FAILEDOPERATION_INSUFFICIENTPRIVILEGE = "FailedOperation.InsufficientPrivilege"
//  FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAMETERS = "InvalidParameter.EmptyParameters"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILESIZE = "InvalidParameter.InvalidFileSize"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETER_REQUESTPARSEERROR = "InvalidParameter.RequestParseError"
func (c *Client) InvokeServiceWithContext(ctx context.Context, request *InvokeServiceRequest) (response *InvokeServiceResponse, err error) {
    if request == nil {
        request = NewInvokeServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeService require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeServiceResponse()
    err = c.Send(request, response)
    return
}
