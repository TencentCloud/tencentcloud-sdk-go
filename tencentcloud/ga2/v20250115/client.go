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

package v20250115

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-01-15"

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


func NewDescribeCrossBorderSettlementRequest() (request *DescribeCrossBorderSettlementRequest) {
    request = &DescribeCrossBorderSettlementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ga2", APIVersion, "DescribeCrossBorderSettlement")
    
    
    return
}

func NewDescribeCrossBorderSettlementResponse() (response *DescribeCrossBorderSettlementResponse) {
    response = &DescribeCrossBorderSettlementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCrossBorderSettlement
// 查询跨境账单
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCrossBorderSettlement(request *DescribeCrossBorderSettlementRequest) (response *DescribeCrossBorderSettlementResponse, err error) {
    return c.DescribeCrossBorderSettlementWithContext(context.Background(), request)
}

// DescribeCrossBorderSettlement
// 查询跨境账单
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTOUTOFRANGE = "InvalidParameter.InputOutOfRange"
//  INVALIDPARAMETERVALUE_MALFORMED = "InvalidParameterValue.Malformed"
func (c *Client) DescribeCrossBorderSettlementWithContext(ctx context.Context, request *DescribeCrossBorderSettlementRequest) (response *DescribeCrossBorderSettlementResponse, err error) {
    if request == nil {
        request = NewDescribeCrossBorderSettlementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ga2", APIVersion, "DescribeCrossBorderSettlement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossBorderSettlement require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossBorderSettlementResponse()
    err = c.Send(request, response)
    return
}
