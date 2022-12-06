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

package v20220726

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-07-26"

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


func NewEvaluateUserRiskRequest() (request *EvaluateUserRiskRequest) {
    request = &EvaluateUserRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trdp", APIVersion, "EvaluateUserRisk")
    
    
    return
}

func NewEvaluateUserRiskResponse() (response *EvaluateUserRiskResponse) {
    response = &EvaluateUserRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EvaluateUserRisk
// 用户风险质量接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFINGERPRINTNOTFOUND = "FailedOperation.DeviceFingerprintNotFound"
//  FAILEDOPERATION_UNKNOWNMODELID = "FailedOperation.UnKnownModelId"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
func (c *Client) EvaluateUserRisk(request *EvaluateUserRiskRequest) (response *EvaluateUserRiskResponse, err error) {
    return c.EvaluateUserRiskWithContext(context.Background(), request)
}

// EvaluateUserRisk
// 用户风险质量接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEVICEFINGERPRINTNOTFOUND = "FailedOperation.DeviceFingerprintNotFound"
//  FAILEDOPERATION_UNKNOWNMODELID = "FailedOperation.UnKnownModelId"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  UNAUTHORIZEDOPERATION_TENANTNOTACTIVATED = "UnauthorizedOperation.TenantNotActivated"
func (c *Client) EvaluateUserRiskWithContext(ctx context.Context, request *EvaluateUserRiskRequest) (response *EvaluateUserRiskResponse, err error) {
    if request == nil {
        request = NewEvaluateUserRiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EvaluateUserRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewEvaluateUserRiskResponse()
    err = c.Send(request, response)
    return
}
