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

package v20230812

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-08-12"

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


func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "RunInstances")
    
    
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunInstances
// 本接口 (RunInstances) 用于创建一个或多个指定配置的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSACCOUNTCANNOTRUNINSTANCES = "FailedOperation.ArrearsAccountCannotRunInstances"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"
//  INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    return c.RunInstancesWithContext(context.Background(), request)
}

// RunInstances
// 本接口 (RunInstances) 用于创建一个或多个指定配置的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSACCOUNTCANNOTRUNINSTANCES = "FailedOperation.ArrearsAccountCannotRunInstances"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"
//  INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"
//  RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"
func (c *Client) RunInstancesWithContext(ctx context.Context, request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("hai", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateInstances
// 本接口 (TerminateInstances) 用于主动退还实例。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// 本接口 (TerminateInstances) 用于主动退还实例。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
