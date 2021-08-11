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

package v20210701

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-07-01"

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


func NewDescribeDeployApplicationDetailRequest() (request *DescribeDeployApplicationDetailRequest) {
    request = &DescribeDeployApplicationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "DescribeDeployApplicationDetail")
    return
}

func NewDescribeDeployApplicationDetailResponse() (response *DescribeDeployApplicationDetailResponse) {
    response = &DescribeDeployApplicationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeployApplicationDetail
// 获取分批发布详情
func (c *Client) DescribeDeployApplicationDetail(request *DescribeDeployApplicationDetailRequest) (response *DescribeDeployApplicationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDeployApplicationDetailRequest()
    }
    response = NewDescribeDeployApplicationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewResumeDeployApplicationRequest() (request *ResumeDeployApplicationRequest) {
    request = &ResumeDeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "ResumeDeployApplication")
    return
}

func NewResumeDeployApplicationResponse() (response *ResumeDeployApplicationResponse) {
    response = &ResumeDeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeDeployApplication
// 开始下一批次发布
func (c *Client) ResumeDeployApplication(request *ResumeDeployApplicationRequest) (response *ResumeDeployApplicationResponse, err error) {
    if request == nil {
        request = NewResumeDeployApplicationRequest()
    }
    response = NewResumeDeployApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewRevertDeployApplicationRequest() (request *RevertDeployApplicationRequest) {
    request = &RevertDeployApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tem", APIVersion, "RevertDeployApplication")
    return
}

func NewRevertDeployApplicationResponse() (response *RevertDeployApplicationResponse) {
    response = &RevertDeployApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevertDeployApplication
// 回滚分批发布
func (c *Client) RevertDeployApplication(request *RevertDeployApplicationRequest) (response *RevertDeployApplicationResponse, err error) {
    if request == nil {
        request = NewRevertDeployApplicationRequest()
    }
    response = NewRevertDeployApplicationResponse()
    err = c.Send(request, response)
    return
}
