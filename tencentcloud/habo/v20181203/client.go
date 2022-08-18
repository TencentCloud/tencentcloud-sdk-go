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

package v20181203

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-12-03"

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


func NewDescribeStatusRequest() (request *DescribeStatusRequest) {
    request = &DescribeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("habo", APIVersion, "DescribeStatus")
    
    
    return
}

func NewDescribeStatusResponse() (response *DescribeStatusResponse) {
    response = &DescribeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStatus
// 查询指定md5样本是否分析完成，并获取分析日志下载地址。
func (c *Client) DescribeStatus(request *DescribeStatusRequest) (response *DescribeStatusResponse, err error) {
    return c.DescribeStatusWithContext(context.Background(), request)
}

// DescribeStatus
// 查询指定md5样本是否分析完成，并获取分析日志下载地址。
func (c *Client) DescribeStatusWithContext(ctx context.Context, request *DescribeStatusRequest) (response *DescribeStatusResponse, err error) {
    if request == nil {
        request = NewDescribeStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStatusResponse()
    err = c.Send(request, response)
    return
}

func NewStartAnalyseRequest() (request *StartAnalyseRequest) {
    request = &StartAnalyseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("habo", APIVersion, "StartAnalyse")
    
    
    return
}

func NewStartAnalyseResponse() (response *StartAnalyseResponse) {
    response = &StartAnalyseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartAnalyse
// 上传样本到哈勃进行分析，异步生成分析日志。
func (c *Client) StartAnalyse(request *StartAnalyseRequest) (response *StartAnalyseResponse, err error) {
    return c.StartAnalyseWithContext(context.Background(), request)
}

// StartAnalyse
// 上传样本到哈勃进行分析，异步生成分析日志。
func (c *Client) StartAnalyseWithContext(ctx context.Context, request *StartAnalyseRequest) (response *StartAnalyseResponse, err error) {
    if request == nil {
        request = NewStartAnalyseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAnalyse require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAnalyseResponse()
    err = c.Send(request, response)
    return
}
