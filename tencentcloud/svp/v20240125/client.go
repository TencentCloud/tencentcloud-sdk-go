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

package v20240125

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-01-25"

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


func NewCreateSavingPlanOrderRequest() (request *CreateSavingPlanOrderRequest) {
    request = &CreateSavingPlanOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("svp", APIVersion, "CreateSavingPlanOrder")
    
    
    return
}

func NewCreateSavingPlanOrderResponse() (response *CreateSavingPlanOrderResponse) {
    response = &CreateSavingPlanOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSavingPlanOrder
// 创建节省计划订单
func (c *Client) CreateSavingPlanOrder(request *CreateSavingPlanOrderRequest) (response *CreateSavingPlanOrderResponse, err error) {
    return c.CreateSavingPlanOrderWithContext(context.Background(), request)
}

// CreateSavingPlanOrder
// 创建节省计划订单
func (c *Client) CreateSavingPlanOrderWithContext(ctx context.Context, request *CreateSavingPlanOrderRequest) (response *CreateSavingPlanOrderResponse, err error) {
    if request == nil {
        request = NewCreateSavingPlanOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSavingPlanOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSavingPlanOrderResponse()
    err = c.Send(request, response)
    return
}
