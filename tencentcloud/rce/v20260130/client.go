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

package v20260130

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-01-30"

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


func NewAssessDeviceRiskPremiumProRequest() (request *AssessDeviceRiskPremiumProRequest) {
    request = &AssessDeviceRiskPremiumProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "AssessDeviceRiskPremiumPro")
    
    
    return
}

func NewAssessDeviceRiskPremiumProResponse() (response *AssessDeviceRiskPremiumProResponse) {
    response = &AssessDeviceRiskPremiumProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssessDeviceRiskPremiumPro
// 设备风险评估-高级版
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMEXCEPTION = "InternalError.SystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_LENGTHEXCEED = "InvalidParameter.LengthExceed"
//  INVALIDPARAMETERVALUE_EVENTNOTEXIST = "InvalidParameterValue.EventNotExist"
//  INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssessDeviceRiskPremiumPro(request *AssessDeviceRiskPremiumProRequest) (response *AssessDeviceRiskPremiumProResponse, err error) {
    return c.AssessDeviceRiskPremiumProWithContext(context.Background(), request)
}

// AssessDeviceRiskPremiumPro
// 设备风险评估-高级版
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMEXCEPTION = "InternalError.SystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_LENGTHEXCEED = "InvalidParameter.LengthExceed"
//  INVALIDPARAMETERVALUE_EVENTNOTEXIST = "InvalidParameterValue.EventNotExist"
//  INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssessDeviceRiskPremiumProWithContext(ctx context.Context, request *AssessDeviceRiskPremiumProRequest) (response *AssessDeviceRiskPremiumProResponse, err error) {
    if request == nil {
        request = NewAssessDeviceRiskPremiumProRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rce", APIVersion, "AssessDeviceRiskPremiumPro")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssessDeviceRiskPremiumPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssessDeviceRiskPremiumProResponse()
    err = c.Send(request, response)
    return
}

func NewAssessDeviceRiskProRequest() (request *AssessDeviceRiskProRequest) {
    request = &AssessDeviceRiskProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "AssessDeviceRiskPro")
    
    
    return
}

func NewAssessDeviceRiskProResponse() (response *AssessDeviceRiskProResponse) {
    response = &AssessDeviceRiskProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssessDeviceRiskPro
// 设备风险评估-基础版
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMEXCEPTION = "InternalError.SystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_LENGTHEXCEED = "InvalidParameter.LengthExceed"
//  INVALIDPARAMETERVALUE_EVENTNOTEXIST = "InvalidParameterValue.EventNotExist"
//  INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssessDeviceRiskPro(request *AssessDeviceRiskProRequest) (response *AssessDeviceRiskProResponse, err error) {
    return c.AssessDeviceRiskProWithContext(context.Background(), request)
}

// AssessDeviceRiskPro
// 设备风险评估-基础版
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMEXCEPTION = "InternalError.SystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_LENGTHEXCEED = "InvalidParameter.LengthExceed"
//  INVALIDPARAMETERVALUE_EVENTNOTEXIST = "InvalidParameterValue.EventNotExist"
//  INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssessDeviceRiskProWithContext(ctx context.Context, request *AssessDeviceRiskProRequest) (response *AssessDeviceRiskProResponse, err error) {
    if request == nil {
        request = NewAssessDeviceRiskProRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rce", APIVersion, "AssessDeviceRiskPro")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssessDeviceRiskPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssessDeviceRiskProResponse()
    err = c.Send(request, response)
    return
}

func NewAssessEnvironmentRiskRequest() (request *AssessEnvironmentRiskRequest) {
    request = &AssessEnvironmentRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("rce", APIVersion, "AssessEnvironmentRisk")
    
    
    return
}

func NewAssessEnvironmentRiskResponse() (response *AssessEnvironmentRiskResponse) {
    response = &AssessEnvironmentRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssessEnvironmentRisk
// 环境风险评估
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMEXCEPTION = "InternalError.SystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_LENGTHEXCEED = "InvalidParameter.LengthExceed"
//  INVALIDPARAMETERVALUE_EVENTNOTEXIST = "InvalidParameterValue.EventNotExist"
//  INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssessEnvironmentRisk(request *AssessEnvironmentRiskRequest) (response *AssessEnvironmentRiskResponse, err error) {
    return c.AssessEnvironmentRiskWithContext(context.Background(), request)
}

// AssessEnvironmentRisk
// 环境风险评估
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMEXCEPTION = "InternalError.SystemException"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_LENGTHEXCEED = "InvalidParameter.LengthExceed"
//  INVALIDPARAMETERVALUE_EVENTNOTEXIST = "InvalidParameterValue.EventNotExist"
//  INVALIDPARAMETERVALUE_TENANTNOTEXIST = "InvalidParameterValue.TenantNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FIELDMISSED = "MissingParameter.FieldMissed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssessEnvironmentRiskWithContext(ctx context.Context, request *AssessEnvironmentRiskRequest) (response *AssessEnvironmentRiskResponse, err error) {
    if request == nil {
        request = NewAssessEnvironmentRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "rce", APIVersion, "AssessEnvironmentRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssessEnvironmentRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssessEnvironmentRiskResponse()
    err = c.Send(request, response)
    return
}
