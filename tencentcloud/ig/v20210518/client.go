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

package v20210518

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-18"

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


func NewDescribeIgOrderListRequest() (request *DescribeIgOrderListRequest) {
    request = &DescribeIgOrderListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ig", APIVersion, "DescribeIgOrderList")
    
    
    return
}

func NewDescribeIgOrderListResponse() (response *DescribeIgOrderListResponse) {
    response = &DescribeIgOrderListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIgOrderList
// 查询智能导诊订单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIgOrderList(request *DescribeIgOrderListRequest) (response *DescribeIgOrderListResponse, err error) {
    return c.DescribeIgOrderListWithContext(context.Background(), request)
}

// DescribeIgOrderList
// 查询智能导诊订单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIgOrderListWithContext(ctx context.Context, request *DescribeIgOrderListRequest) (response *DescribeIgOrderListResponse, err error) {
    if request == nil {
        request = NewDescribeIgOrderListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ig", APIVersion, "DescribeIgOrderList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIgOrderList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIgOrderListResponse()
    err = c.Send(request, response)
    return
}
