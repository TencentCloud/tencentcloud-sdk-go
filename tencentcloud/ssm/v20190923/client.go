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

package v20190923

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-23"

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


func NewCreateSecretRequest() (request *CreateSecretRequest) {
    request = &CreateSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "CreateSecret")
    return
}

func NewCreateSecretResponse() (response *CreateSecretResponse) {
    response = &CreateSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建新的凭据信息，通过KMS进行加密保护。每个Region最多可创建存储1000个凭据信息。
func (c *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
    if request == nil {
        request = NewCreateSecretRequest()
    }
    response = NewCreateSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecretRequest() (request *DeleteSecretRequest) {
    request = &DeleteSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DeleteSecret")
    return
}

func NewDeleteSecretResponse() (response *DeleteSecretResponse) {
    response = &DeleteSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除指定的凭据信息，可以通过RecoveryWindowInDays参数设置立即删除或者计划删除。对于计划删除的凭据，在删除日期到达之前状态为 PendingDelete，并可以通过RestoreSecret 进行恢复，超出指定删除日期之后会被彻底删除。您必须先通过 DisableSecret 停用凭据后才可以进行（计划）删除操作。
func (c *Client) DeleteSecret(request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
    if request == nil {
        request = NewDeleteSecretRequest()
    }
    response = NewDeleteSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecretVersionRequest() (request *DeleteSecretVersionRequest) {
    request = &DeleteSecretVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DeleteSecretVersion")
    return
}

func NewDeleteSecretVersionResponse() (response *DeleteSecretVersionResponse) {
    response = &DeleteSecretVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于直接删除指定凭据下的单个版本凭据，删除操作立即生效，对所有状态下的凭据版本都可以删除。
func (c *Client) DeleteSecretVersion(request *DeleteSecretVersionRequest) (response *DeleteSecretVersionResponse, err error) {
    if request == nil {
        request = NewDeleteSecretVersionRequest()
    }
    response = NewDeleteSecretVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecretRequest() (request *DescribeSecretRequest) {
    request = &DescribeSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DescribeSecret")
    return
}

func NewDescribeSecretResponse() (response *DescribeSecretResponse) {
    response = &DescribeSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取凭据的详细属性信息。
func (c *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
    if request == nil {
        request = NewDescribeSecretRequest()
    }
    response = NewDescribeSecretResponse()
    err = c.Send(request, response)
    return
}

func NewDisableSecretRequest() (request *DisableSecretRequest) {
    request = &DisableSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "DisableSecret")
    return
}

func NewDisableSecretResponse() (response *DisableSecretResponse) {
    response = &DisableSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 停用指定的凭据，停用后状态为 Disabled，无法通过接口获取该凭据的明文。
func (c *Client) DisableSecret(request *DisableSecretRequest) (response *DisableSecretResponse, err error) {
    if request == nil {
        request = NewDisableSecretRequest()
    }
    response = NewDisableSecretResponse()
    err = c.Send(request, response)
    return
}

func NewEnableSecretRequest() (request *EnableSecretRequest) {
    request = &EnableSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "EnableSecret")
    return
}

func NewEnableSecretResponse() (response *EnableSecretResponse) {
    response = &EnableSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于开启凭据，状态为Enabled。可以通过 GetSecretValue 接口获取凭据明文。处于PendingDelete状态的凭据不能直接开启，需要通过RestoreSecret 恢复后再开启使用。
func (c *Client) EnableSecret(request *EnableSecretRequest) (response *EnableSecretResponse, err error) {
    if request == nil {
        request = NewEnableSecretRequest()
    }
    response = NewEnableSecretResponse()
    err = c.Send(request, response)
    return
}

func NewGetRegionsRequest() (request *GetRegionsRequest) {
    request = &GetRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetRegions")
    return
}

func NewGetRegionsResponse() (response *GetRegionsResponse) {
    response = &GetRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取控制台展示region列表
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
    if request == nil {
        request = NewGetRegionsRequest()
    }
    response = NewGetRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetSecretValueRequest() (request *GetSecretValueRequest) {
    request = &GetSecretValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetSecretValue")
    return
}

func NewGetSecretValueResponse() (response *GetSecretValueResponse) {
    response = &GetSecretValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定凭据名称和版本的凭据明文信息，只能获取启用状态的凭据明文。
func (c *Client) GetSecretValue(request *GetSecretValueRequest) (response *GetSecretValueResponse, err error) {
    if request == nil {
        request = NewGetSecretValueRequest()
    }
    response = NewGetSecretValueResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceStatusRequest() (request *GetServiceStatusRequest) {
    request = &GetServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "GetServiceStatus")
    return
}

func NewGetServiceStatusResponse() (response *GetServiceStatusResponse) {
    response = &GetServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用户获取用户SecretsManager服务开通状态。
func (c *Client) GetServiceStatus(request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceStatusRequest()
    }
    response = NewGetServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewListSecretVersionIdsRequest() (request *ListSecretVersionIdsRequest) {
    request = &ListSecretVersionIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "ListSecretVersionIds")
    return
}

func NewListSecretVersionIdsResponse() (response *ListSecretVersionIdsResponse) {
    response = &ListSecretVersionIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于获取指定凭据下的版本列表信息
func (c *Client) ListSecretVersionIds(request *ListSecretVersionIdsRequest) (response *ListSecretVersionIdsResponse, err error) {
    if request == nil {
        request = NewListSecretVersionIdsRequest()
    }
    response = NewListSecretVersionIdsResponse()
    err = c.Send(request, response)
    return
}

func NewListSecretsRequest() (request *ListSecretsRequest) {
    request = &ListSecretsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "ListSecrets")
    return
}

func NewListSecretsResponse() (response *ListSecretsResponse) {
    response = &ListSecretsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于获取所有凭据的详细列表，可以指定过滤字段、排序方式等。
func (c *Client) ListSecrets(request *ListSecretsRequest) (response *ListSecretsResponse, err error) {
    if request == nil {
        request = NewListSecretsRequest()
    }
    response = NewListSecretsResponse()
    err = c.Send(request, response)
    return
}

func NewPutSecretValueRequest() (request *PutSecretValueRequest) {
    request = &PutSecretValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "PutSecretValue")
    return
}

func NewPutSecretValueResponse() (response *PutSecretValueResponse) {
    response = &PutSecretValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口在指定名称的凭据下增加新版本的凭据内容，一个凭据下最多可以支持10个版本。只能对处于Enabled 和 Disabled 状态的凭据添加新的版本。
func (c *Client) PutSecretValue(request *PutSecretValueRequest) (response *PutSecretValueResponse, err error) {
    if request == nil {
        request = NewPutSecretValueRequest()
    }
    response = NewPutSecretValueResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreSecretRequest() (request *RestoreSecretRequest) {
    request = &RestoreSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "RestoreSecret")
    return
}

func NewRestoreSecretResponse() (response *RestoreSecretResponse) {
    response = &RestoreSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于恢复计划删除（PendingDelete状态）中的凭据，取消计划删除。取消计划删除的凭据将处于Disabled 状态，如需恢复使用，通过EnableSecret 接口开启凭据。
func (c *Client) RestoreSecret(request *RestoreSecretRequest) (response *RestoreSecretResponse, err error) {
    if request == nil {
        request = NewRestoreSecretRequest()
    }
    response = NewRestoreSecretResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDescriptionRequest() (request *UpdateDescriptionRequest) {
    request = &UpdateDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "UpdateDescription")
    return
}

func NewUpdateDescriptionResponse() (response *UpdateDescriptionResponse) {
    response = &UpdateDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于修改指定凭据的描述信息，仅能修改Enabled 和 Disabled 状态的凭据。
func (c *Client) UpdateDescription(request *UpdateDescriptionRequest) (response *UpdateDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateDescriptionRequest()
    }
    response = NewUpdateDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSecretRequest() (request *UpdateSecretRequest) {
    request = &UpdateSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ssm", APIVersion, "UpdateSecret")
    return
}

func NewUpdateSecretResponse() (response *UpdateSecretResponse) {
    response = &UpdateSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于更新指定凭据名称和版本号的内容，调用该接口会对新的凭据内容加密后覆盖旧的内容。仅允许更新Enabled 和 Disabled 状态的凭据。
func (c *Client) UpdateSecret(request *UpdateSecretRequest) (response *UpdateSecretResponse, err error) {
    if request == nil {
        request = NewUpdateSecretRequest()
    }
    response = NewUpdateSecretResponse()
    err = c.Send(request, response)
    return
}
