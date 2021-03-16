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

package v20180416

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

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


func NewCopyFunctionRequest() (request *CopyFunctionRequest) {
    request = &CopyFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CopyFunction")
    return
}

func NewCopyFunctionResponse() (response *CopyFunctionResponse) {
    response = &CopyFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 复制一个函数，您可以选择将复制出的新函数放置在特定的Region和Namespace。
// 注：本接口**不会**复制函数的以下对象或属性：
// 1. 函数的触发器
// 2. 除了$LATEST以外的其它版本
// 3. 函数配置的日志投递到的CLS目标。
// 
// 如有需要，您可以在复制后手动配置新函数。
func (c *Client) CopyFunction(request *CopyFunctionRequest) (response *CopyFunctionResponse, err error) {
    if request == nil {
        request = NewCopyFunctionRequest()
    }
    response = NewCopyFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasRequest() (request *CreateAliasRequest) {
    request = &CreateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateAlias")
    return
}

func NewCreateAliasResponse() (response *CreateAliasResponse) {
    response = &CreateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 为某个函数版本创建一个别名，您可以使用别名来标记特定的函数版本，如DEV/RELEASE版本，也可以随时修改别名指向的版本。
// 一个别名必须指向一个主版本，此外还可以同时指向一个附加版本。调用函数时指定特定的别名，则请求会被发送到别名指向的版本上，您可以配置请求发送到主版本和附加版本的比例。
func (c *Client) CreateAlias(request *CreateAliasRequest) (response *CreateAliasResponse, err error) {
    if request == nil {
        request = NewCreateAliasRequest()
    }
    response = NewCreateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFunctionRequest() (request *CreateFunctionRequest) {
    request = &CreateFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateFunction")
    return
}

func NewCreateFunctionResponse() (response *CreateFunctionResponse) {
    response = &CreateFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入参数创建新的函数。
func (c *Client) CreateFunction(request *CreateFunctionRequest) (response *CreateFunctionResponse, err error) {
    if request == nil {
        request = NewCreateFunctionRequest()
    }
    response = NewCreateFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
    request = &CreateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateNamespace")
    return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
    response = &CreateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入的参数创建命名空间。
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
    if request == nil {
        request = NewCreateNamespaceRequest()
    }
    response = NewCreateNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTriggerRequest() (request *CreateTriggerRequest) {
    request = &CreateTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "CreateTrigger")
    return
}

func NewCreateTriggerResponse() (response *CreateTriggerResponse) {
    response = &CreateTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据参数输入设置新的触发方式。
func (c *Client) CreateTrigger(request *CreateTriggerRequest) (response *CreateTriggerResponse, err error) {
    if request == nil {
        request = NewCreateTriggerRequest()
    }
    response = NewCreateTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasRequest() (request *DeleteAliasRequest) {
    request = &DeleteAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteAlias")
    return
}

func NewDeleteAliasResponse() (response *DeleteAliasResponse) {
    response = &DeleteAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除一个函数版本的别名
func (c *Client) DeleteAlias(request *DeleteAliasRequest) (response *DeleteAliasResponse, err error) {
    if request == nil {
        request = NewDeleteAliasRequest()
    }
    response = NewDeleteAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFunctionRequest() (request *DeleteFunctionRequest) {
    request = &DeleteFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteFunction")
    return
}

func NewDeleteFunctionResponse() (response *DeleteFunctionResponse) {
    response = &DeleteFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入参数删除函数。
func (c *Client) DeleteFunction(request *DeleteFunctionRequest) (response *DeleteFunctionResponse, err error) {
    if request == nil {
        request = NewDeleteFunctionRequest()
    }
    response = NewDeleteFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLayerVersionRequest() (request *DeleteLayerVersionRequest) {
    request = &DeleteLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteLayerVersion")
    return
}

func NewDeleteLayerVersionResponse() (response *DeleteLayerVersionResponse) {
    response = &DeleteLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除指定层的指定版本，被删除的版本无法再关联到函数上，但不会影响正在引用这个层的函数。
func (c *Client) DeleteLayerVersion(request *DeleteLayerVersionRequest) (response *DeleteLayerVersionResponse, err error) {
    if request == nil {
        request = NewDeleteLayerVersionRequest()
    }
    response = NewDeleteLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
    request = &DeleteNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteNamespace")
    return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
    response = &DeleteNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入的参数创建命名空间。
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
    if request == nil {
        request = NewDeleteNamespaceRequest()
    }
    response = NewDeleteNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProvisionedConcurrencyConfigRequest() (request *DeleteProvisionedConcurrencyConfigRequest) {
    request = &DeleteProvisionedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteProvisionedConcurrencyConfig")
    return
}

func NewDeleteProvisionedConcurrencyConfigResponse() (response *DeleteProvisionedConcurrencyConfigResponse) {
    response = &DeleteProvisionedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除函数版本的预置并发配置。
func (c *Client) DeleteProvisionedConcurrencyConfig(request *DeleteProvisionedConcurrencyConfigRequest) (response *DeleteProvisionedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewDeleteProvisionedConcurrencyConfigRequest()
    }
    response = NewDeleteProvisionedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteReservedConcurrencyConfigRequest() (request *DeleteReservedConcurrencyConfigRequest) {
    request = &DeleteReservedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteReservedConcurrencyConfig")
    return
}

func NewDeleteReservedConcurrencyConfigResponse() (response *DeleteReservedConcurrencyConfigResponse) {
    response = &DeleteReservedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除函数的保留并发配置。
func (c *Client) DeleteReservedConcurrencyConfig(request *DeleteReservedConcurrencyConfigRequest) (response *DeleteReservedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewDeleteReservedConcurrencyConfigRequest()
    }
    response = NewDeleteReservedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTriggerRequest() (request *DeleteTriggerRequest) {
    request = &DeleteTriggerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "DeleteTrigger")
    return
}

func NewDeleteTriggerResponse() (response *DeleteTriggerResponse) {
    response = &DeleteTriggerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据参数传入删除已有的触发方式。
func (c *Client) DeleteTrigger(request *DeleteTriggerRequest) (response *DeleteTriggerResponse, err error) {
    if request == nil {
        request = NewDeleteTriggerRequest()
    }
    response = NewDeleteTriggerResponse()
    err = c.Send(request, response)
    return
}

func NewGetAccountRequest() (request *GetAccountRequest) {
    request = &GetAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetAccount")
    return
}

func NewGetAccountResponse() (response *GetAccountResponse) {
    response = &GetAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取账户信息
func (c *Client) GetAccount(request *GetAccountRequest) (response *GetAccountResponse, err error) {
    if request == nil {
        request = NewGetAccountRequest()
    }
    response = NewGetAccountResponse()
    err = c.Send(request, response)
    return
}

func NewGetAliasRequest() (request *GetAliasRequest) {
    request = &GetAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetAlias")
    return
}

func NewGetAliasResponse() (response *GetAliasResponse) {
    response = &GetAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取别名的详细信息，包括名称、描述、版本、路由信息等。
func (c *Client) GetAlias(request *GetAliasRequest) (response *GetAliasResponse, err error) {
    if request == nil {
        request = NewGetAliasRequest()
    }
    response = NewGetAliasResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionRequest() (request *GetFunctionRequest) {
    request = &GetFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunction")
    return
}

func NewGetFunctionResponse() (response *GetFunctionResponse) {
    response = &GetFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口获取某个函数的详细信息，包括名称、代码、处理方法、关联触发器和超时时间等字段。
func (c *Client) GetFunction(request *GetFunctionRequest) (response *GetFunctionResponse, err error) {
    if request == nil {
        request = NewGetFunctionRequest()
    }
    response = NewGetFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionAddressRequest() (request *GetFunctionAddressRequest) {
    request = &GetFunctionAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionAddress")
    return
}

func NewGetFunctionAddressResponse() (response *GetFunctionAddressResponse) {
    response = &GetFunctionAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于获取函数代码包的下载地址。
func (c *Client) GetFunctionAddress(request *GetFunctionAddressRequest) (response *GetFunctionAddressResponse, err error) {
    if request == nil {
        request = NewGetFunctionAddressRequest()
    }
    response = NewGetFunctionAddressResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionEventInvokeConfigRequest() (request *GetFunctionEventInvokeConfigRequest) {
    request = &GetFunctionEventInvokeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionEventInvokeConfig")
    return
}

func NewGetFunctionEventInvokeConfigResponse() (response *GetFunctionEventInvokeConfigResponse) {
    response = &GetFunctionEventInvokeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取函数异步重试配置，包括重试次数和消息保留时间
func (c *Client) GetFunctionEventInvokeConfig(request *GetFunctionEventInvokeConfigRequest) (response *GetFunctionEventInvokeConfigResponse, err error) {
    if request == nil {
        request = NewGetFunctionEventInvokeConfigRequest()
    }
    response = NewGetFunctionEventInvokeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewGetFunctionLogsRequest() (request *GetFunctionLogsRequest) {
    request = &GetFunctionLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetFunctionLogs")
    return
}

func NewGetFunctionLogsResponse() (response *GetFunctionLogsResponse) {
    response = &GetFunctionLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据指定的日志查询条件返回函数运行日志。
func (c *Client) GetFunctionLogs(request *GetFunctionLogsRequest) (response *GetFunctionLogsResponse, err error) {
    if request == nil {
        request = NewGetFunctionLogsRequest()
    }
    response = NewGetFunctionLogsResponse()
    err = c.Send(request, response)
    return
}

func NewGetLayerVersionRequest() (request *GetLayerVersionRequest) {
    request = &GetLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetLayerVersion")
    return
}

func NewGetLayerVersionResponse() (response *GetLayerVersionResponse) {
    response = &GetLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取层版本详细信息，包括用于下载层中文件的链接。
func (c *Client) GetLayerVersion(request *GetLayerVersionRequest) (response *GetLayerVersionResponse, err error) {
    if request == nil {
        request = NewGetLayerVersionRequest()
    }
    response = NewGetLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetProvisionedConcurrencyConfigRequest() (request *GetProvisionedConcurrencyConfigRequest) {
    request = &GetProvisionedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetProvisionedConcurrencyConfig")
    return
}

func NewGetProvisionedConcurrencyConfigResponse() (response *GetProvisionedConcurrencyConfigResponse) {
    response = &GetProvisionedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取函数或函数某一版本的预置并发详情。
func (c *Client) GetProvisionedConcurrencyConfig(request *GetProvisionedConcurrencyConfigRequest) (response *GetProvisionedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewGetProvisionedConcurrencyConfigRequest()
    }
    response = NewGetProvisionedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewGetReservedConcurrencyConfigRequest() (request *GetReservedConcurrencyConfigRequest) {
    request = &GetReservedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "GetReservedConcurrencyConfig")
    return
}

func NewGetReservedConcurrencyConfigResponse() (response *GetReservedConcurrencyConfigResponse) {
    response = &GetReservedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取函数的保留并发详情。
func (c *Client) GetReservedConcurrencyConfig(request *GetReservedConcurrencyConfigRequest) (response *GetReservedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewGetReservedConcurrencyConfigRequest()
    }
    response = NewGetReservedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeRequest() (request *InvokeRequest) {
    request = &InvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "Invoke")
    return
}

func NewInvokeResponse() (response *InvokeResponse) {
    response = &InvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于运行函数。
func (c *Client) Invoke(request *InvokeRequest) (response *InvokeResponse, err error) {
    if request == nil {
        request = NewInvokeRequest()
    }
    response = NewInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewListAliasesRequest() (request *ListAliasesRequest) {
    request = &ListAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListAliases")
    return
}

func NewListAliasesResponse() (response *ListAliasesResponse) {
    response = &ListAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回一个函数下的全部别名，可以根据特定函数版本过滤。
func (c *Client) ListAliases(request *ListAliasesRequest) (response *ListAliasesResponse, err error) {
    if request == nil {
        request = NewListAliasesRequest()
    }
    response = NewListAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewListAsyncEventsRequest() (request *ListAsyncEventsRequest) {
    request = &ListAsyncEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListAsyncEvents")
    return
}

func NewListAsyncEventsResponse() (response *ListAsyncEventsResponse) {
    response = &ListAsyncEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取函数异步事件列表
func (c *Client) ListAsyncEvents(request *ListAsyncEventsRequest) (response *ListAsyncEventsResponse, err error) {
    if request == nil {
        request = NewListAsyncEventsRequest()
    }
    response = NewListAsyncEventsResponse()
    err = c.Send(request, response)
    return
}

func NewListFunctionsRequest() (request *ListFunctionsRequest) {
    request = &ListFunctionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListFunctions")
    return
}

func NewListFunctionsResponse() (response *ListFunctionsResponse) {
    response = &ListFunctionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入的查询参数返回相关函数信息。
func (c *Client) ListFunctions(request *ListFunctionsRequest) (response *ListFunctionsResponse, err error) {
    if request == nil {
        request = NewListFunctionsRequest()
    }
    response = NewListFunctionsResponse()
    err = c.Send(request, response)
    return
}

func NewListLayerVersionsRequest() (request *ListLayerVersionsRequest) {
    request = &ListLayerVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListLayerVersions")
    return
}

func NewListLayerVersionsResponse() (response *ListLayerVersionsResponse) {
    response = &ListLayerVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回指定层的全部版本的信息
func (c *Client) ListLayerVersions(request *ListLayerVersionsRequest) (response *ListLayerVersionsResponse, err error) {
    if request == nil {
        request = NewListLayerVersionsRequest()
    }
    response = NewListLayerVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListLayersRequest() (request *ListLayersRequest) {
    request = &ListLayersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListLayers")
    return
}

func NewListLayersResponse() (response *ListLayersResponse) {
    response = &ListLayersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 返回全部层的列表，其中包含了每个层最新版本的信息，可以通过适配运行时进行过滤。
func (c *Client) ListLayers(request *ListLayersRequest) (response *ListLayersResponse, err error) {
    if request == nil {
        request = NewListLayersRequest()
    }
    response = NewListLayersResponse()
    err = c.Send(request, response)
    return
}

func NewListNamespacesRequest() (request *ListNamespacesRequest) {
    request = &ListNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListNamespaces")
    return
}

func NewListNamespacesResponse() (response *ListNamespacesResponse) {
    response = &ListNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 列出命名空间列表
func (c *Client) ListNamespaces(request *ListNamespacesRequest) (response *ListNamespacesResponse, err error) {
    if request == nil {
        request = NewListNamespacesRequest()
    }
    response = NewListNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewListTriggersRequest() (request *ListTriggersRequest) {
    request = &ListTriggersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListTriggers")
    return
}

func NewListTriggersResponse() (response *ListTriggersResponse) {
    response = &ListTriggersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取函数触发器列表
func (c *Client) ListTriggers(request *ListTriggersRequest) (response *ListTriggersResponse, err error) {
    if request == nil {
        request = NewListTriggersRequest()
    }
    response = NewListTriggersResponse()
    err = c.Send(request, response)
    return
}

func NewListVersionByFunctionRequest() (request *ListVersionByFunctionRequest) {
    request = &ListVersionByFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "ListVersionByFunction")
    return
}

func NewListVersionByFunctionResponse() (response *ListVersionByFunctionResponse) {
    response = &ListVersionByFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入的参数查询函数的版本。
func (c *Client) ListVersionByFunction(request *ListVersionByFunctionRequest) (response *ListVersionByFunctionResponse, err error) {
    if request == nil {
        request = NewListVersionByFunctionRequest()
    }
    response = NewListVersionByFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewPublishLayerVersionRequest() (request *PublishLayerVersionRequest) {
    request = &PublishLayerVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PublishLayerVersion")
    return
}

func NewPublishLayerVersionResponse() (response *PublishLayerVersionResponse) {
    response = &PublishLayerVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用给定的zip文件或cos对象创建一个层的新版本，每次使用相同的层的名称调用本接口，都会生成一个新版本。
func (c *Client) PublishLayerVersion(request *PublishLayerVersionRequest) (response *PublishLayerVersionResponse, err error) {
    if request == nil {
        request = NewPublishLayerVersionRequest()
    }
    response = NewPublishLayerVersionResponse()
    err = c.Send(request, response)
    return
}

func NewPublishVersionRequest() (request *PublishVersionRequest) {
    request = &PublishVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PublishVersion")
    return
}

func NewPublishVersionResponse() (response *PublishVersionResponse) {
    response = &PublishVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口用于用户发布新版本函数。
func (c *Client) PublishVersion(request *PublishVersionRequest) (response *PublishVersionResponse, err error) {
    if request == nil {
        request = NewPublishVersionRequest()
    }
    response = NewPublishVersionResponse()
    err = c.Send(request, response)
    return
}

func NewPutProvisionedConcurrencyConfigRequest() (request *PutProvisionedConcurrencyConfigRequest) {
    request = &PutProvisionedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PutProvisionedConcurrencyConfig")
    return
}

func NewPutProvisionedConcurrencyConfigResponse() (response *PutProvisionedConcurrencyConfigResponse) {
    response = &PutProvisionedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置函数某一非$LATEST版本的预置并发。
func (c *Client) PutProvisionedConcurrencyConfig(request *PutProvisionedConcurrencyConfigRequest) (response *PutProvisionedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewPutProvisionedConcurrencyConfigRequest()
    }
    response = NewPutProvisionedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewPutReservedConcurrencyConfigRequest() (request *PutReservedConcurrencyConfigRequest) {
    request = &PutReservedConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PutReservedConcurrencyConfig")
    return
}

func NewPutReservedConcurrencyConfigResponse() (response *PutReservedConcurrencyConfigResponse) {
    response = &PutReservedConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置函数保留并发
func (c *Client) PutReservedConcurrencyConfig(request *PutReservedConcurrencyConfigRequest) (response *PutReservedConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewPutReservedConcurrencyConfigRequest()
    }
    response = NewPutReservedConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewPutTotalConcurrencyConfigRequest() (request *PutTotalConcurrencyConfigRequest) {
    request = &PutTotalConcurrencyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "PutTotalConcurrencyConfig")
    return
}

func NewPutTotalConcurrencyConfigResponse() (response *PutTotalConcurrencyConfigResponse) {
    response = &PutTotalConcurrencyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改账号并发限制配额
func (c *Client) PutTotalConcurrencyConfig(request *PutTotalConcurrencyConfigRequest) (response *PutTotalConcurrencyConfigResponse, err error) {
    if request == nil {
        request = NewPutTotalConcurrencyConfigRequest()
    }
    response = NewPutTotalConcurrencyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateAsyncEventRequest() (request *TerminateAsyncEventRequest) {
    request = &TerminateAsyncEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "TerminateAsyncEvent")
    return
}

func NewTerminateAsyncEventResponse() (response *TerminateAsyncEventResponse) {
    response = &TerminateAsyncEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 终止正在运行中的函数异步事件
func (c *Client) TerminateAsyncEvent(request *TerminateAsyncEventRequest) (response *TerminateAsyncEventResponse, err error) {
    if request == nil {
        request = NewTerminateAsyncEventRequest()
    }
    response = NewTerminateAsyncEventResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAliasRequest() (request *UpdateAliasRequest) {
    request = &UpdateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateAlias")
    return
}

func NewUpdateAliasResponse() (response *UpdateAliasResponse) {
    response = &UpdateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新别名的配置
func (c *Client) UpdateAlias(request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    response = NewUpdateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionCodeRequest() (request *UpdateFunctionCodeRequest) {
    request = &UpdateFunctionCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionCode")
    return
}

func NewUpdateFunctionCodeResponse() (response *UpdateFunctionCodeResponse) {
    response = &UpdateFunctionCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入参数更新函数代码。
func (c *Client) UpdateFunctionCode(request *UpdateFunctionCodeRequest) (response *UpdateFunctionCodeResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionCodeRequest()
    }
    response = NewUpdateFunctionCodeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionConfigurationRequest() (request *UpdateFunctionConfigurationRequest) {
    request = &UpdateFunctionConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionConfiguration")
    return
}

func NewUpdateFunctionConfigurationResponse() (response *UpdateFunctionConfigurationResponse) {
    response = &UpdateFunctionConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口根据传入参数更新函数配置。
func (c *Client) UpdateFunctionConfiguration(request *UpdateFunctionConfigurationRequest) (response *UpdateFunctionConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionConfigurationRequest()
    }
    response = NewUpdateFunctionConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFunctionEventInvokeConfigRequest() (request *UpdateFunctionEventInvokeConfigRequest) {
    request = &UpdateFunctionEventInvokeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateFunctionEventInvokeConfig")
    return
}

func NewUpdateFunctionEventInvokeConfigResponse() (response *UpdateFunctionEventInvokeConfigResponse) {
    response = &UpdateFunctionEventInvokeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新函数的异步重试配置，包括重试次数和消息保留时间
func (c *Client) UpdateFunctionEventInvokeConfig(request *UpdateFunctionEventInvokeConfigRequest) (response *UpdateFunctionEventInvokeConfigResponse, err error) {
    if request == nil {
        request = NewUpdateFunctionEventInvokeConfigRequest()
    }
    response = NewUpdateFunctionEventInvokeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNamespaceRequest() (request *UpdateNamespaceRequest) {
    request = &UpdateNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("scf", APIVersion, "UpdateNamespace")
    return
}

func NewUpdateNamespaceResponse() (response *UpdateNamespaceResponse) {
    response = &UpdateNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新命名空间
func (c *Client) UpdateNamespace(request *UpdateNamespaceRequest) (response *UpdateNamespaceResponse, err error) {
    if request == nil {
        request = NewUpdateNamespaceRequest()
    }
    response = NewUpdateNamespaceResponse()
    err = c.Send(request, response)
    return
}
