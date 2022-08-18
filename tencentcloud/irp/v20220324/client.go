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

package v20220324

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-03-24"

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


func NewRecommendContentRequest() (request *RecommendContentRequest) {
    request = &RecommendContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "RecommendContent")
    
    
    return
}

func NewRecommendContentResponse() (response *RecommendContentResponse) {
    response = &RecommendContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecommendContent
// 获取推荐结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RecommendContent(request *RecommendContentRequest) (response *RecommendContentResponse, err error) {
    return c.RecommendContentWithContext(context.Background(), request)
}

// RecommendContent
// 获取推荐结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RecommendContentWithContext(ctx context.Context, request *RecommendContentRequest) (response *RecommendContentResponse, err error) {
    if request == nil {
        request = NewRecommendContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecommendContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecommendContentResponse()
    err = c.Send(request, response)
    return
}

func NewReportActionRequest() (request *ReportActionRequest) {
    request = &ReportActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportAction")
    
    
    return
}

func NewReportActionResponse() (response *ReportActionResponse) {
    response = &ReportActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportAction
// 上报行为
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportAction(request *ReportActionRequest) (response *ReportActionResponse, err error) {
    return c.ReportActionWithContext(context.Background(), request)
}

// ReportAction
// 上报行为
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportActionWithContext(ctx context.Context, request *ReportActionRequest) (response *ReportActionResponse, err error) {
    if request == nil {
        request = NewReportActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportActionResponse()
    err = c.Send(request, response)
    return
}

func NewReportMaterialRequest() (request *ReportMaterialRequest) {
    request = &ReportMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportMaterial")
    
    
    return
}

func NewReportMaterialResponse() (response *ReportMaterialResponse) {
    response = &ReportMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportMaterial
// 上报物料
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportMaterial(request *ReportMaterialRequest) (response *ReportMaterialResponse, err error) {
    return c.ReportMaterialWithContext(context.Background(), request)
}

// ReportMaterial
// 上报物料
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportMaterialWithContext(ctx context.Context, request *ReportMaterialRequest) (response *ReportMaterialResponse, err error) {
    if request == nil {
        request = NewReportMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportMaterial require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewReportPortraitRequest() (request *ReportPortraitRequest) {
    request = &ReportPortraitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportPortrait")
    
    
    return
}

func NewReportPortraitResponse() (response *ReportPortraitResponse) {
    response = &ReportPortraitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportPortrait
// 上报用户画像
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportPortrait(request *ReportPortraitRequest) (response *ReportPortraitResponse, err error) {
    return c.ReportPortraitWithContext(context.Background(), request)
}

// ReportPortrait
// 上报用户画像
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportPortraitWithContext(ctx context.Context, request *ReportPortraitRequest) (response *ReportPortraitResponse, err error) {
    if request == nil {
        request = NewReportPortraitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportPortrait require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportPortraitResponse()
    err = c.Send(request, response)
    return
}
