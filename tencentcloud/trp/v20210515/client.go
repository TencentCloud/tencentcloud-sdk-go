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

package v20210515

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-15"

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


func NewCreateCodeBatchRequest() (request *CreateCodeBatchRequest) {
    request = &CreateCodeBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "CreateCodeBatch")
    
    
    return
}

func NewCreateCodeBatchResponse() (response *CreateCodeBatchResponse) {
    response = &CreateCodeBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCodeBatch
// 新增批次
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateCodeBatch(request *CreateCodeBatchRequest) (response *CreateCodeBatchResponse, err error) {
    return c.CreateCodeBatchWithContext(context.Background(), request)
}

// CreateCodeBatch
// 新增批次
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateCodeBatchWithContext(ctx context.Context, request *CreateCodeBatchRequest) (response *CreateCodeBatchResponse, err error) {
    if request == nil {
        request = NewCreateCodeBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodeBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodeBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCodePackRequest() (request *CreateCodePackRequest) {
    request = &CreateCodePackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "CreateCodePack")
    
    
    return
}

func NewCreateCodePackResponse() (response *CreateCodePackResponse) {
    response = &CreateCodePackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCodePack
// 生成码包
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateCodePack(request *CreateCodePackRequest) (response *CreateCodePackResponse, err error) {
    return c.CreateCodePackWithContext(context.Background(), request)
}

// CreateCodePack
// 生成码包
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateCodePackWithContext(ctx context.Context, request *CreateCodePackRequest) (response *CreateCodePackResponse, err error) {
    if request == nil {
        request = NewCreateCodePackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodePack require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodePackResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMerchantRequest() (request *CreateMerchantRequest) {
    request = &CreateMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "CreateMerchant")
    
    
    return
}

func NewCreateMerchantResponse() (response *CreateMerchantResponse) {
    response = &CreateMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMerchant
// 新建商户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMerchant(request *CreateMerchantRequest) (response *CreateMerchantResponse, err error) {
    return c.CreateMerchantWithContext(context.Background(), request)
}

// CreateMerchant
// 新建商户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateMerchantWithContext(ctx context.Context, request *CreateMerchantRequest) (response *CreateMerchantResponse, err error) {
    if request == nil {
        request = NewCreateMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMerchant require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProductRequest() (request *CreateProductRequest) {
    request = &CreateProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "CreateProduct")
    
    
    return
}

func NewCreateProductResponse() (response *CreateProductResponse) {
    response = &CreateProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProduct
// 新建商品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    return c.CreateProductWithContext(context.Background(), request)
}

// CreateProduct
// 新建商品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateProductWithContext(ctx context.Context, request *CreateProductRequest) (response *CreateProductResponse, err error) {
    if request == nil {
        request = NewCreateProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProductResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTraceChainRequest() (request *CreateTraceChainRequest) {
    request = &CreateTraceChainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "CreateTraceChain")
    
    
    return
}

func NewCreateTraceChainResponse() (response *CreateTraceChainResponse) {
    response = &CreateTraceChainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTraceChain
// 上链溯源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTraceChain(request *CreateTraceChainRequest) (response *CreateTraceChainResponse, err error) {
    return c.CreateTraceChainWithContext(context.Background(), request)
}

// CreateTraceChain
// 上链溯源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTraceChainWithContext(ctx context.Context, request *CreateTraceChainRequest) (response *CreateTraceChainResponse, err error) {
    if request == nil {
        request = NewCreateTraceChainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTraceChain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTraceChainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTraceCodesRequest() (request *CreateTraceCodesRequest) {
    request = &CreateTraceCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "CreateTraceCodes")
    
    
    return
}

func NewCreateTraceCodesResponse() (response *CreateTraceCodesResponse) {
    response = &CreateTraceCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTraceCodes
// 批量导入二维码，只支持平台发的码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTraceCodes(request *CreateTraceCodesRequest) (response *CreateTraceCodesResponse, err error) {
    return c.CreateTraceCodesWithContext(context.Background(), request)
}

// CreateTraceCodes
// 批量导入二维码，只支持平台发的码
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTraceCodesWithContext(ctx context.Context, request *CreateTraceCodesRequest) (response *CreateTraceCodesResponse, err error) {
    if request == nil {
        request = NewCreateTraceCodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTraceCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTraceCodesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTraceDataRequest() (request *CreateTraceDataRequest) {
    request = &CreateTraceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "CreateTraceData")
    
    
    return
}

func NewCreateTraceDataResponse() (response *CreateTraceDataResponse) {
    response = &CreateTraceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTraceData
// 新增溯源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTraceData(request *CreateTraceDataRequest) (response *CreateTraceDataResponse, err error) {
    return c.CreateTraceDataWithContext(context.Background(), request)
}

// CreateTraceData
// 新增溯源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateTraceDataWithContext(ctx context.Context, request *CreateTraceDataRequest) (response *CreateTraceDataResponse, err error) {
    if request == nil {
        request = NewCreateTraceDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTraceData require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTraceDataResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCodeBatchRequest() (request *DeleteCodeBatchRequest) {
    request = &DeleteCodeBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DeleteCodeBatch")
    
    
    return
}

func NewDeleteCodeBatchResponse() (response *DeleteCodeBatchResponse) {
    response = &DeleteCodeBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCodeBatch
// 删除批次
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCodeBatch(request *DeleteCodeBatchRequest) (response *DeleteCodeBatchResponse, err error) {
    return c.DeleteCodeBatchWithContext(context.Background(), request)
}

// DeleteCodeBatch
// 删除批次
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteCodeBatchWithContext(ctx context.Context, request *DeleteCodeBatchRequest) (response *DeleteCodeBatchResponse, err error) {
    if request == nil {
        request = NewDeleteCodeBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCodeBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCodeBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMerchantRequest() (request *DeleteMerchantRequest) {
    request = &DeleteMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DeleteMerchant")
    
    
    return
}

func NewDeleteMerchantResponse() (response *DeleteMerchantResponse) {
    response = &DeleteMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMerchant
// 删除商户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteMerchant(request *DeleteMerchantRequest) (response *DeleteMerchantResponse, err error) {
    return c.DeleteMerchantWithContext(context.Background(), request)
}

// DeleteMerchant
// 删除商户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteMerchantWithContext(ctx context.Context, request *DeleteMerchantRequest) (response *DeleteMerchantResponse, err error) {
    if request == nil {
        request = NewDeleteMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMerchant require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
    request = &DeleteProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DeleteProduct")
    
    
    return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
    response = &DeleteProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProduct
// 删除商品，如果商品被使用，则不可删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    return c.DeleteProductWithContext(context.Background(), request)
}

// DeleteProduct
// 删除商品，如果商品被使用，则不可删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteProductWithContext(ctx context.Context, request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    if request == nil {
        request = NewDeleteProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTraceDataRequest() (request *DeleteTraceDataRequest) {
    request = &DeleteTraceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DeleteTraceData")
    
    
    return
}

func NewDeleteTraceDataResponse() (response *DeleteTraceDataResponse) {
    response = &DeleteTraceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTraceData
// 删除溯源信息，如果已经上链则不可删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteTraceData(request *DeleteTraceDataRequest) (response *DeleteTraceDataResponse, err error) {
    return c.DeleteTraceDataWithContext(context.Background(), request)
}

// DeleteTraceData
// 删除溯源信息，如果已经上链则不可删除
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteTraceDataWithContext(ctx context.Context, request *DeleteTraceDataRequest) (response *DeleteTraceDataResponse, err error) {
    if request == nil {
        request = NewDeleteTraceDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTraceData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTraceDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodeBatchByIdRequest() (request *DescribeCodeBatchByIdRequest) {
    request = &DescribeCodeBatchByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeCodeBatchById")
    
    
    return
}

func NewDescribeCodeBatchByIdResponse() (response *DescribeCodeBatchByIdResponse) {
    response = &DescribeCodeBatchByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCodeBatchById
// 查询批次信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodeBatchById(request *DescribeCodeBatchByIdRequest) (response *DescribeCodeBatchByIdResponse, err error) {
    return c.DescribeCodeBatchByIdWithContext(context.Background(), request)
}

// DescribeCodeBatchById
// 查询批次信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodeBatchByIdWithContext(ctx context.Context, request *DescribeCodeBatchByIdRequest) (response *DescribeCodeBatchByIdResponse, err error) {
    if request == nil {
        request = NewDescribeCodeBatchByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCodeBatchById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCodeBatchByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodeBatchsRequest() (request *DescribeCodeBatchsRequest) {
    request = &DescribeCodeBatchsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeCodeBatchs")
    
    
    return
}

func NewDescribeCodeBatchsResponse() (response *DescribeCodeBatchsResponse) {
    response = &DescribeCodeBatchsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCodeBatchs
// 查询批次列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodeBatchs(request *DescribeCodeBatchsRequest) (response *DescribeCodeBatchsResponse, err error) {
    return c.DescribeCodeBatchsWithContext(context.Background(), request)
}

// DescribeCodeBatchs
// 查询批次列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodeBatchsWithContext(ctx context.Context, request *DescribeCodeBatchsRequest) (response *DescribeCodeBatchsResponse, err error) {
    if request == nil {
        request = NewDescribeCodeBatchsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCodeBatchs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCodeBatchsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodePacksRequest() (request *DescribeCodePacksRequest) {
    request = &DescribeCodePacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeCodePacks")
    
    
    return
}

func NewDescribeCodePacksResponse() (response *DescribeCodePacksResponse) {
    response = &DescribeCodePacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCodePacks
// 查询码包列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodePacks(request *DescribeCodePacksRequest) (response *DescribeCodePacksResponse, err error) {
    return c.DescribeCodePacksWithContext(context.Background(), request)
}

// DescribeCodePacks
// 查询码包列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodePacksWithContext(ctx context.Context, request *DescribeCodePacksRequest) (response *DescribeCodePacksResponse, err error) {
    if request == nil {
        request = NewDescribeCodePacksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCodePacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCodePacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodesByPackRequest() (request *DescribeCodesByPackRequest) {
    request = &DescribeCodesByPackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeCodesByPack")
    
    
    return
}

func NewDescribeCodesByPackResponse() (response *DescribeCodesByPackResponse) {
    response = &DescribeCodesByPackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCodesByPack
// 查询码包的二维码列表，上限 3 万
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodesByPack(request *DescribeCodesByPackRequest) (response *DescribeCodesByPackResponse, err error) {
    return c.DescribeCodesByPackWithContext(context.Background(), request)
}

// DescribeCodesByPack
// 查询码包的二维码列表，上限 3 万
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCodesByPackWithContext(ctx context.Context, request *DescribeCodesByPackRequest) (response *DescribeCodesByPackResponse, err error) {
    if request == nil {
        request = NewDescribeCodesByPackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCodesByPack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCodesByPackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMerchantByIdRequest() (request *DescribeMerchantByIdRequest) {
    request = &DescribeMerchantByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeMerchantById")
    
    
    return
}

func NewDescribeMerchantByIdResponse() (response *DescribeMerchantByIdResponse) {
    response = &DescribeMerchantByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMerchantById
// 查询商户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMerchantById(request *DescribeMerchantByIdRequest) (response *DescribeMerchantByIdResponse, err error) {
    return c.DescribeMerchantByIdWithContext(context.Background(), request)
}

// DescribeMerchantById
// 查询商户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMerchantByIdWithContext(ctx context.Context, request *DescribeMerchantByIdRequest) (response *DescribeMerchantByIdResponse, err error) {
    if request == nil {
        request = NewDescribeMerchantByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMerchantById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMerchantByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMerchantsRequest() (request *DescribeMerchantsRequest) {
    request = &DescribeMerchantsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeMerchants")
    
    
    return
}

func NewDescribeMerchantsResponse() (response *DescribeMerchantsResponse) {
    response = &DescribeMerchantsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMerchants
// 查询商户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMerchants(request *DescribeMerchantsRequest) (response *DescribeMerchantsResponse, err error) {
    return c.DescribeMerchantsWithContext(context.Background(), request)
}

// DescribeMerchants
// 查询商户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMerchantsWithContext(ctx context.Context, request *DescribeMerchantsRequest) (response *DescribeMerchantsResponse, err error) {
    if request == nil {
        request = NewDescribeMerchantsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMerchants require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMerchantsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductByIdRequest() (request *DescribeProductByIdRequest) {
    request = &DescribeProductByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeProductById")
    
    
    return
}

func NewDescribeProductByIdResponse() (response *DescribeProductByIdResponse) {
    response = &DescribeProductByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductById
// 查询商品信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProductById(request *DescribeProductByIdRequest) (response *DescribeProductByIdResponse, err error) {
    return c.DescribeProductByIdWithContext(context.Background(), request)
}

// DescribeProductById
// 查询商品信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProductByIdWithContext(ctx context.Context, request *DescribeProductByIdRequest) (response *DescribeProductByIdResponse, err error) {
    if request == nil {
        request = NewDescribeProductByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductsRequest() (request *DescribeProductsRequest) {
    request = &DescribeProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeProducts")
    
    
    return
}

func NewDescribeProductsResponse() (response *DescribeProductsResponse) {
    response = &DescribeProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProducts
// 查询商品列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProducts(request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    return c.DescribeProductsWithContext(context.Background(), request)
}

// DescribeProducts
// 查询商品列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProductsWithContext(ctx context.Context, request *DescribeProductsRequest) (response *DescribeProductsResponse, err error) {
    if request == nil {
        request = NewDescribeProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTraceCodeByIdRequest() (request *DescribeTraceCodeByIdRequest) {
    request = &DescribeTraceCodeByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeTraceCodeById")
    
    
    return
}

func NewDescribeTraceCodeByIdResponse() (response *DescribeTraceCodeByIdResponse) {
    response = &DescribeTraceCodeByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTraceCodeById
// 查询二维码信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTraceCodeById(request *DescribeTraceCodeByIdRequest) (response *DescribeTraceCodeByIdResponse, err error) {
    return c.DescribeTraceCodeByIdWithContext(context.Background(), request)
}

// DescribeTraceCodeById
// 查询二维码信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTraceCodeByIdWithContext(ctx context.Context, request *DescribeTraceCodeByIdRequest) (response *DescribeTraceCodeByIdResponse, err error) {
    if request == nil {
        request = NewDescribeTraceCodeByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTraceCodeById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTraceCodeByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTraceCodesRequest() (request *DescribeTraceCodesRequest) {
    request = &DescribeTraceCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeTraceCodes")
    
    
    return
}

func NewDescribeTraceCodesResponse() (response *DescribeTraceCodesResponse) {
    response = &DescribeTraceCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTraceCodes
// 查询二维码列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTraceCodes(request *DescribeTraceCodesRequest) (response *DescribeTraceCodesResponse, err error) {
    return c.DescribeTraceCodesWithContext(context.Background(), request)
}

// DescribeTraceCodes
// 查询二维码列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTraceCodesWithContext(ctx context.Context, request *DescribeTraceCodesRequest) (response *DescribeTraceCodesResponse, err error) {
    if request == nil {
        request = NewDescribeTraceCodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTraceCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTraceCodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTraceDataListRequest() (request *DescribeTraceDataListRequest) {
    request = &DescribeTraceDataListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "DescribeTraceDataList")
    
    
    return
}

func NewDescribeTraceDataListResponse() (response *DescribeTraceDataListResponse) {
    response = &DescribeTraceDataListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTraceDataList
// 查询溯源信息，通常溯源信息跟生产批次绑定，即一个批次的所有溯源信息都是一样的
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTraceDataList(request *DescribeTraceDataListRequest) (response *DescribeTraceDataListResponse, err error) {
    return c.DescribeTraceDataListWithContext(context.Background(), request)
}

// DescribeTraceDataList
// 查询溯源信息，通常溯源信息跟生产批次绑定，即一个批次的所有溯源信息都是一样的
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTraceDataListWithContext(ctx context.Context, request *DescribeTraceDataListRequest) (response *DescribeTraceDataListResponse, err error) {
    if request == nil {
        request = NewDescribeTraceDataListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTraceDataList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTraceDataListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCodeBatchRequest() (request *ModifyCodeBatchRequest) {
    request = &ModifyCodeBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "ModifyCodeBatch")
    
    
    return
}

func NewModifyCodeBatchResponse() (response *ModifyCodeBatchResponse) {
    response = &ModifyCodeBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCodeBatch
// 修改批次
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCodeBatch(request *ModifyCodeBatchRequest) (response *ModifyCodeBatchResponse, err error) {
    return c.ModifyCodeBatchWithContext(context.Background(), request)
}

// ModifyCodeBatch
// 修改批次
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCodeBatchWithContext(ctx context.Context, request *ModifyCodeBatchRequest) (response *ModifyCodeBatchResponse, err error) {
    if request == nil {
        request = NewModifyCodeBatchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCodeBatch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCodeBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMerchantRequest() (request *ModifyMerchantRequest) {
    request = &ModifyMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "ModifyMerchant")
    
    
    return
}

func NewModifyMerchantResponse() (response *ModifyMerchantResponse) {
    response = &ModifyMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMerchant
// 编辑商户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyMerchant(request *ModifyMerchantRequest) (response *ModifyMerchantResponse, err error) {
    return c.ModifyMerchantWithContext(context.Background(), request)
}

// ModifyMerchant
// 编辑商户
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyMerchantWithContext(ctx context.Context, request *ModifyMerchantRequest) (response *ModifyMerchantResponse, err error) {
    if request == nil {
        request = NewModifyMerchantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMerchant require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProductRequest() (request *ModifyProductRequest) {
    request = &ModifyProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "ModifyProduct")
    
    
    return
}

func NewModifyProductResponse() (response *ModifyProductResponse) {
    response = &ModifyProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyProduct
// 编辑商品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyProduct(request *ModifyProductRequest) (response *ModifyProductResponse, err error) {
    return c.ModifyProductWithContext(context.Background(), request)
}

// ModifyProduct
// 编辑商品
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyProductWithContext(ctx context.Context, request *ModifyProductRequest) (response *ModifyProductResponse, err error) {
    if request == nil {
        request = NewModifyProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProductResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTraceCodeRequest() (request *ModifyTraceCodeRequest) {
    request = &ModifyTraceCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "ModifyTraceCode")
    
    
    return
}

func NewModifyTraceCodeResponse() (response *ModifyTraceCodeResponse) {
    response = &ModifyTraceCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTraceCode
// 冻结或者激活二维码，所属的批次的冻结状态优先级大于单个二维码的状态，即如果批次是冻结的，那么该批次下二维码的状态都是冻结的
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyTraceCode(request *ModifyTraceCodeRequest) (response *ModifyTraceCodeResponse, err error) {
    return c.ModifyTraceCodeWithContext(context.Background(), request)
}

// ModifyTraceCode
// 冻结或者激活二维码，所属的批次的冻结状态优先级大于单个二维码的状态，即如果批次是冻结的，那么该批次下二维码的状态都是冻结的
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyTraceCodeWithContext(ctx context.Context, request *ModifyTraceCodeRequest) (response *ModifyTraceCodeResponse, err error) {
    if request == nil {
        request = NewModifyTraceCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTraceCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTraceCodeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTraceDataRequest() (request *ModifyTraceDataRequest) {
    request = &ModifyTraceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "ModifyTraceData")
    
    
    return
}

func NewModifyTraceDataResponse() (response *ModifyTraceDataResponse) {
    response = &ModifyTraceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTraceData
// 修改溯源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyTraceData(request *ModifyTraceDataRequest) (response *ModifyTraceDataResponse, err error) {
    return c.ModifyTraceDataWithContext(context.Background(), request)
}

// ModifyTraceData
// 修改溯源信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyTraceDataWithContext(ctx context.Context, request *ModifyTraceDataRequest) (response *ModifyTraceDataResponse, err error) {
    if request == nil {
        request = NewModifyTraceDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTraceData require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTraceDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTraceDataRanksRequest() (request *ModifyTraceDataRanksRequest) {
    request = &ModifyTraceDataRanksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("trp", APIVersion, "ModifyTraceDataRanks")
    
    
    return
}

func NewModifyTraceDataRanksResponse() (response *ModifyTraceDataRanksResponse) {
    response = &ModifyTraceDataRanksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTraceDataRanks
// 修改溯源信息的排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyTraceDataRanks(request *ModifyTraceDataRanksRequest) (response *ModifyTraceDataRanksResponse, err error) {
    return c.ModifyTraceDataRanksWithContext(context.Background(), request)
}

// ModifyTraceDataRanks
// 修改溯源信息的排序
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyTraceDataRanksWithContext(ctx context.Context, request *ModifyTraceDataRanksRequest) (response *ModifyTraceDataRanksResponse, err error) {
    if request == nil {
        request = NewModifyTraceDataRanksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTraceDataRanks require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTraceDataRanksResponse()
    err = c.Send(request, response)
    return
}
