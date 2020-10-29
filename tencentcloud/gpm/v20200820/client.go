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

package v20200820

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-08-20"

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


func NewCancelMatchingRequest() (request *CancelMatchingRequest) {
    request = &CancelMatchingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "CancelMatching")
    return
}

func NewCancelMatchingResponse() (response *CancelMatchingResponse) {
    response = &CancelMatchingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 取消匹配。
func (c *Client) CancelMatching(request *CancelMatchingRequest) (response *CancelMatchingResponse, err error) {
    if request == nil {
        request = NewCancelMatchingRequest()
    }
    response = NewCancelMatchingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMatchRequest() (request *CreateMatchRequest) {
    request = &CreateMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "CreateMatch")
    return
}

func NewCreateMatchResponse() (response *CreateMatchResponse) {
    response = &CreateMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建匹配
func (c *Client) CreateMatch(request *CreateMatchRequest) (response *CreateMatchResponse, err error) {
    if request == nil {
        request = NewCreateMatchRequest()
    }
    response = NewCreateMatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "CreateRule")
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建规则
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMatchRequest() (request *DeleteMatchRequest) {
    request = &DeleteMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DeleteMatch")
    return
}

func NewDeleteMatchResponse() (response *DeleteMatchResponse) {
    response = &DeleteMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除匹配
func (c *Client) DeleteMatch(request *DeleteMatchRequest) (response *DeleteMatchResponse, err error) {
    if request == nil {
        request = NewDeleteMatchRequest()
    }
    response = NewDeleteMatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DeleteRule")
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除规则
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataRequest() (request *DescribeDataRequest) {
    request = &DescribeDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeData")
    return
}

func NewDescribeDataResponse() (response *DescribeDataResponse) {
    response = &DescribeDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 统计数据
func (c *Client) DescribeData(request *DescribeDataRequest) (response *DescribeDataResponse, err error) {
    if request == nil {
        request = NewDescribeDataRequest()
    }
    response = NewDescribeDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchRequest() (request *DescribeMatchRequest) {
    request = &DescribeMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatch")
    return
}

func NewDescribeMatchResponse() (response *DescribeMatchResponse) {
    response = &DescribeMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询匹配详情
func (c *Client) DescribeMatch(request *DescribeMatchRequest) (response *DescribeMatchResponse, err error) {
    if request == nil {
        request = NewDescribeMatchRequest()
    }
    response = NewDescribeMatchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchCodesRequest() (request *DescribeMatchCodesRequest) {
    request = &DescribeMatchCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatchCodes")
    return
}

func NewDescribeMatchCodesResponse() (response *DescribeMatchCodesResponse) {
    response = &DescribeMatchCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分页查询匹配Code
func (c *Client) DescribeMatchCodes(request *DescribeMatchCodesRequest) (response *DescribeMatchCodesResponse, err error) {
    if request == nil {
        request = NewDescribeMatchCodesRequest()
    }
    response = NewDescribeMatchCodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchesRequest() (request *DescribeMatchesRequest) {
    request = &DescribeMatchesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatches")
    return
}

func NewDescribeMatchesResponse() (response *DescribeMatchesResponse) {
    response = &DescribeMatchesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分页查询匹配列表
func (c *Client) DescribeMatches(request *DescribeMatchesRequest) (response *DescribeMatchesResponse, err error) {
    if request == nil {
        request = NewDescribeMatchesRequest()
    }
    response = NewDescribeMatchesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMatchingProgressRequest() (request *DescribeMatchingProgressRequest) {
    request = &DescribeMatchingProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeMatchingProgress")
    return
}

func NewDescribeMatchingProgressResponse() (response *DescribeMatchingProgressResponse) {
    response = &DescribeMatchingProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询匹配进度。
func (c *Client) DescribeMatchingProgress(request *DescribeMatchingProgressRequest) (response *DescribeMatchingProgressResponse, err error) {
    if request == nil {
        request = NewDescribeMatchingProgressRequest()
    }
    response = NewDescribeMatchingProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleRequest() (request *DescribeRuleRequest) {
    request = &DescribeRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeRule")
    return
}

func NewDescribeRuleResponse() (response *DescribeRuleResponse) {
    response = &DescribeRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询规则详情
func (c *Client) DescribeRule(request *DescribeRuleRequest) (response *DescribeRuleResponse, err error) {
    if request == nil {
        request = NewDescribeRuleRequest()
    }
    response = NewDescribeRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "DescribeRules")
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 分页查询规则集列表
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMatchRequest() (request *ModifyMatchRequest) {
    request = &ModifyMatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "ModifyMatch")
    return
}

func NewModifyMatchResponse() (response *ModifyMatchResponse) {
    response = &ModifyMatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改匹配
func (c *Client) ModifyMatch(request *ModifyMatchRequest) (response *ModifyMatchResponse, err error) {
    if request == nil {
        request = NewModifyMatchRequest()
    }
    response = NewModifyMatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "ModifyRule")
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改规则（描述、标签）
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewStartMatchingRequest() (request *StartMatchingRequest) {
    request = &StartMatchingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gpm", APIVersion, "StartMatching")
    return
}

func NewStartMatchingResponse() (response *StartMatchingResponse) {
    response = &StartMatchingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 支持传入一个玩家或多个玩家发起匹配，在同一个请求内的玩家将被分到同一个对局。
func (c *Client) StartMatching(request *StartMatchingRequest) (response *StartMatchingResponse, err error) {
    if request == nil {
        request = NewStartMatchingRequest()
    }
    response = NewStartMatchingResponse()
    err = c.Send(request, response)
    return
}
