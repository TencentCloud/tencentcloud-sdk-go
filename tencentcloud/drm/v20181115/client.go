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

package v20181115

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-11-15"

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


func NewAddFairPlayPemRequest() (request *AddFairPlayPemRequest) {
    request = &AddFairPlayPemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("drm", APIVersion, "AddFairPlayPem")
    return
}

func NewAddFairPlayPemResponse() (response *AddFairPlayPemResponse) {
    response = &AddFairPlayPemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用来设置fairplay方案所需的私钥、私钥密钥、ask等信息。
// 如需使用fairplay方案，请务必先设置私钥。
func (c *Client) AddFairPlayPem(request *AddFairPlayPemRequest) (response *AddFairPlayPemResponse, err error) {
    if request == nil {
        request = NewAddFairPlayPemRequest()
    }
    response = NewAddFairPlayPemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLicenseRequest() (request *CreateLicenseRequest) {
    request = &CreateLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("drm", APIVersion, "CreateLicense")
    return
}

func NewCreateLicenseResponse() (response *CreateLicenseResponse) {
    response = &CreateLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用来生成DRM方案对应的播放许可证，开发者需提供DRM方案类型、内容类型参数，后台将生成许可证后返回许可证数据
// 开发者需要转发终端设备发出的许可证请求信息。
func (c *Client) CreateLicense(request *CreateLicenseRequest) (response *CreateLicenseResponse, err error) {
    if request == nil {
        request = NewCreateLicenseRequest()
    }
    response = NewCreateLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFairPlayPemRequest() (request *DeleteFairPlayPemRequest) {
    request = &DeleteFairPlayPemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("drm", APIVersion, "DeleteFairPlayPem")
    return
}

func NewDeleteFairPlayPemResponse() (response *DeleteFairPlayPemResponse) {
    response = &DeleteFairPlayPemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用来删除fairplay方案的私钥、ask等信息
// 注：高风险操作，删除后，您将无法使用腾讯云DRM提供的fairplay服务。
// 由于缓存，删除操作需要约半小时生效
func (c *Client) DeleteFairPlayPem(request *DeleteFairPlayPemRequest) (response *DeleteFairPlayPemResponse, err error) {
    if request == nil {
        request = NewDeleteFairPlayPemRequest()
    }
    response = NewDeleteFairPlayPemResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFairPlayPemRequest() (request *DescribeFairPlayPemRequest) {
    request = &DescribeFairPlayPemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("drm", APIVersion, "DescribeFairPlayPem")
    return
}

func NewDescribeFairPlayPemResponse() (response *DescribeFairPlayPemResponse) {
    response = &DescribeFairPlayPemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用来查询设置的FairPlay私钥校验信息。可用该接口校验设置的私钥与本身的私钥是否一致。
func (c *Client) DescribeFairPlayPem(request *DescribeFairPlayPemRequest) (response *DescribeFairPlayPemResponse, err error) {
    if request == nil {
        request = NewDescribeFairPlayPemRequest()
    }
    response = NewDescribeFairPlayPemResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeysRequest() (request *DescribeKeysRequest) {
    request = &DescribeKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("drm", APIVersion, "DescribeKeys")
    return
}

func NewDescribeKeysResponse() (response *DescribeKeysResponse) {
    response = &DescribeKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开发者需要指定使用的DRM类型、和需要加密的Track类型，后台返回加密使用的密钥
// 如果加密使用的ContentID没有关联的密钥信息，后台会自动生成新的密钥返回
func (c *Client) DescribeKeys(request *DescribeKeysRequest) (response *DescribeKeysResponse, err error) {
    if request == nil {
        request = NewDescribeKeysRequest()
    }
    response = NewDescribeKeysResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFairPlayPemRequest() (request *ModifyFairPlayPemRequest) {
    request = &ModifyFairPlayPemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("drm", APIVersion, "ModifyFairPlayPem")
    return
}

func NewModifyFairPlayPemResponse() (response *ModifyFairPlayPemResponse) {
    response = &ModifyFairPlayPemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用来设置fairplay方案所需的私钥、私钥密钥、ask等信息。
// 如需使用fairplay方案，请务必先设置私钥。
func (c *Client) ModifyFairPlayPem(request *ModifyFairPlayPemRequest) (response *ModifyFairPlayPemResponse, err error) {
    if request == nil {
        request = NewModifyFairPlayPemRequest()
    }
    response = NewModifyFairPlayPemResponse()
    err = c.Send(request, response)
    return
}

func NewStartEncryptionRequest() (request *StartEncryptionRequest) {
    request = &StartEncryptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("drm", APIVersion, "StartEncryption")
    return
}

func NewStartEncryptionResponse() (response *StartEncryptionResponse) {
    response = &StartEncryptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开发者调用该接口，启动一次内容文件的DRM加密工作流
func (c *Client) StartEncryption(request *StartEncryptionRequest) (response *StartEncryptionResponse, err error) {
    if request == nil {
        request = NewStartEncryptionRequest()
    }
    response = NewStartEncryptionResponse()
    err = c.Send(request, response)
    return
}
