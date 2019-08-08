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

package v20190411

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-04-11"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDescribeEntityRequest() (request *DescribeEntityRequest) {
    request = &DescribeEntityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tkgdq", APIVersion, "DescribeEntity")
    return
}

func NewDescribeEntityResponse() (response *DescribeEntityResponse) {
    response = &DescribeEntityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 输入实体名称，返回实体相关的信息如实体别名、实体英文名、实体详细信息、相关实体等
func (c *Client) DescribeEntity(request *DescribeEntityRequest) (response *DescribeEntityResponse, err error) {
    if request == nil {
        request = NewDescribeEntityRequest()
    }
    response = NewDescribeEntityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelationRequest() (request *DescribeRelationRequest) {
    request = &DescribeRelationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tkgdq", APIVersion, "DescribeRelation")
    return
}

func NewDescribeRelationResponse() (response *DescribeRelationResponse) {
    response = &DescribeRelationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 输入两个实体，返回两个实体间的关系，例如马化腾与腾讯公司不仅是相关实体，二者还存在隶属关系（马化腾属于腾讯公司）。
func (c *Client) DescribeRelation(request *DescribeRelationRequest) (response *DescribeRelationResponse, err error) {
    if request == nil {
        request = NewDescribeRelationRequest()
    }
    response = NewDescribeRelationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTripleRequest() (request *DescribeTripleRequest) {
    request = &DescribeTripleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tkgdq", APIVersion, "DescribeTriple")
    return
}

func NewDescribeTripleResponse() (response *DescribeTripleResponse) {
    response = &DescribeTripleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 三元组查询，主要分为两类，SP查询和PO查询。SP查询表示已知主语和谓语查询宾语，PO查询表示已知宾语和谓语查询主语。每一个SP或PO查询都是一个可独立执行的查询，TQL支持SP查询的嵌套查询，即主语可以是一个嵌套的子查询。其他复杂的三元组查询方法，请参考官网API文档示例。
func (c *Client) DescribeTriple(request *DescribeTripleRequest) (response *DescribeTripleResponse, err error) {
    if request == nil {
        request = NewDescribeTripleRequest()
    }
    response = NewDescribeTripleResponse()
    err = c.Send(request, response)
    return
}
