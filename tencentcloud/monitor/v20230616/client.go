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

package v20230616

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-06-16"

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


func NewDescribeAlarmNotifyHistoriesRequest() (request *DescribeAlarmNotifyHistoriesRequest) {
    request = &DescribeAlarmNotifyHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotifyHistories")
    
    
    return
}

func NewDescribeAlarmNotifyHistoriesResponse() (response *DescribeAlarmNotifyHistoriesResponse) {
    response = &DescribeAlarmNotifyHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotifyHistories
// 按需查询告警的通知历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarmNotifyHistories(request *DescribeAlarmNotifyHistoriesRequest) (response *DescribeAlarmNotifyHistoriesResponse, err error) {
    return c.DescribeAlarmNotifyHistoriesWithContext(context.Background(), request)
}

// DescribeAlarmNotifyHistories
// 按需查询告警的通知历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarmNotifyHistoriesWithContext(ctx context.Context, request *DescribeAlarmNotifyHistoriesRequest) (response *DescribeAlarmNotifyHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNotifyHistoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotifyHistories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNotifyHistoriesResponse()
    err = c.Send(request, response)
    return
}
