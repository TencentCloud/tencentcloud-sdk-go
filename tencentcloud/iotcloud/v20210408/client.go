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

package v20210408

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-08"

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


func NewCreateDeviceRequest() (request *CreateDeviceRequest) {
    request = &CreateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateDevice")
    
    
    return
}

func NewCreateDeviceResponse() (response *CreateDeviceResponse) {
    response = &CreateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDevice
// 本接口（CreateDevice）用于新建一个物联网通信设备。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    return c.CreateDeviceWithContext(context.Background(), request)
}

// CreateDevice
// 本接口（CreateDevice）用于新建一个物联网通信设备。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"
func (c *Client) CreateDeviceWithContext(ctx context.Context, request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateCARequest() (request *CreatePrivateCARequest) {
    request = &CreatePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreatePrivateCA")
    
    
    return
}

func NewCreatePrivateCAResponse() (response *CreatePrivateCAResponse) {
    response = &CreatePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrivateCA
// 创建私有CA证书
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
//  INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) CreatePrivateCA(request *CreatePrivateCARequest) (response *CreatePrivateCAResponse, err error) {
    return c.CreatePrivateCAWithContext(context.Background(), request)
}

// CreatePrivateCA
// 创建私有CA证书
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
//  INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) CreatePrivateCAWithContext(ctx context.Context, request *CreatePrivateCARequest) (response *CreatePrivateCAResponse, err error) {
    if request == nil {
        request = NewCreatePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteDevice")
    
    
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDevice
// 本接口（DeleteDevice）用于删除物联网通信设备。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// 本接口（DeleteDevice）用于删除物联网通信设备。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
func (c *Client) DeleteDeviceWithContext(ctx context.Context, request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateCARequest() (request *DeletePrivateCARequest) {
    request = &DeletePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeletePrivateCA")
    
    
    return
}

func NewDeletePrivateCAResponse() (response *DeletePrivateCAResponse) {
    response = &DeletePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrivateCA
// 删除私有CA证书
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CAALREADYBINDPRODUCT = "LimitExceeded.CAAlreadyBindProduct"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DeletePrivateCA(request *DeletePrivateCARequest) (response *DeletePrivateCAResponse, err error) {
    return c.DeletePrivateCAWithContext(context.Background(), request)
}

// DeletePrivateCA
// 删除私有CA证书
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CAALREADYBINDPRODUCT = "LimitExceeded.CAAlreadyBindProduct"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DeletePrivateCAWithContext(ctx context.Context, request *DeletePrivateCARequest) (response *DeletePrivateCAResponse, err error) {
    if request == nil {
        request = NewDeletePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
    request = &DeleteProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteProduct")
    
    
    return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
    response = &DeleteProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProduct
// 本接口（DeleteProduct）用于删除一个物联网通信产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    return c.DeleteProductWithContext(context.Background(), request)
}

// DeleteProduct
// 本接口（DeleteProduct）用于删除一个物联网通信产品
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
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

func NewDeleteProductPrivateCARequest() (request *DeleteProductPrivateCARequest) {
    request = &DeleteProductPrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteProductPrivateCA")
    
    
    return
}

func NewDeleteProductPrivateCAResponse() (response *DeleteProductPrivateCAResponse) {
    response = &DeleteProductPrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProductPrivateCA
// 删除产品的私有CA证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
func (c *Client) DeleteProductPrivateCA(request *DeleteProductPrivateCARequest) (response *DeleteProductPrivateCAResponse, err error) {
    return c.DeleteProductPrivateCAWithContext(context.Background(), request)
}

// DeleteProductPrivateCA
// 删除产品的私有CA证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
func (c *Client) DeleteProductPrivateCAWithContext(ctx context.Context, request *DeleteProductPrivateCARequest) (response *DeleteProductPrivateCAResponse, err error) {
    if request == nil {
        request = NewDeleteProductPrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProductPrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProductPrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceRequest() (request *DescribeDeviceRequest) {
    request = &DescribeDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDevice")
    
    
    return
}

func NewDescribeDeviceResponse() (response *DescribeDeviceResponse) {
    response = &DescribeDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevice
// 本接口（DescribeDevice）用于查看设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    return c.DescribeDeviceWithContext(context.Background(), request)
}

// DescribeDevice
// 本接口（DescribeDevice）用于查看设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDeviceWithContext(ctx context.Context, request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDevices")
    
    
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevices
// 本接口（DescribeDevices）用于查询物联网通信设备的设备列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    return c.DescribeDevicesWithContext(context.Background(), request)
}

// DescribeDevices
// 本接口（DescribeDevices）用于查询物联网通信设备的设备列表。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevicesWithContext(ctx context.Context, request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCARequest() (request *DescribePrivateCARequest) {
    request = &DescribePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCA")
    
    
    return
}

func NewDescribePrivateCAResponse() (response *DescribePrivateCAResponse) {
    response = &DescribePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCA
// 查询私有化CA信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCA(request *DescribePrivateCARequest) (response *DescribePrivateCAResponse, err error) {
    return c.DescribePrivateCAWithContext(context.Background(), request)
}

// DescribePrivateCA
// 查询私有化CA信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCAWithContext(ctx context.Context, request *DescribePrivateCARequest) (response *DescribePrivateCAResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCABindedProductsRequest() (request *DescribePrivateCABindedProductsRequest) {
    request = &DescribePrivateCABindedProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCABindedProducts")
    
    
    return
}

func NewDescribePrivateCABindedProductsResponse() (response *DescribePrivateCABindedProductsResponse) {
    response = &DescribePrivateCABindedProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCABindedProducts
// 查询私有CA绑定的产品列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCABindedProducts(request *DescribePrivateCABindedProductsRequest) (response *DescribePrivateCABindedProductsResponse, err error) {
    return c.DescribePrivateCABindedProductsWithContext(context.Background(), request)
}

// DescribePrivateCABindedProducts
// 查询私有CA绑定的产品列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCABindedProductsWithContext(ctx context.Context, request *DescribePrivateCABindedProductsRequest) (response *DescribePrivateCABindedProductsResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCABindedProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateCABindedProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateCABindedProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCAsRequest() (request *DescribePrivateCAsRequest) {
    request = &DescribePrivateCAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCAs")
    
    
    return
}

func NewDescribePrivateCAsResponse() (response *DescribePrivateCAsResponse) {
    response = &DescribePrivateCAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCAs
// 查询私有CA证书列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCAs(request *DescribePrivateCAsRequest) (response *DescribePrivateCAsResponse, err error) {
    return c.DescribePrivateCAsWithContext(context.Background(), request)
}

// DescribePrivateCAs
// 查询私有CA证书列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCAsWithContext(ctx context.Context, request *DescribePrivateCAsRequest) (response *DescribePrivateCAsResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrivateCAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrivateCAsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductRequest() (request *DescribeProductRequest) {
    request = &DescribeProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProduct")
    
    
    return
}

func NewDescribeProductResponse() (response *DescribeProductResponse) {
    response = &DescribeProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProduct
// 本接口（DescribeProduct）用于查看产品详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    return c.DescribeProductWithContext(context.Background(), request)
}

// DescribeProduct
// 本接口（DescribeProduct）用于查看产品详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProductWithContext(ctx context.Context, request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    if request == nil {
        request = NewDescribeProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductCARequest() (request *DescribeProductCARequest) {
    request = &DescribeProductCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProductCA")
    
    
    return
}

func NewDescribeProductCAResponse() (response *DescribeProductCAResponse) {
    response = &DescribeProductCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductCA
// 查询产品绑定的CA证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProductCA(request *DescribeProductCARequest) (response *DescribeProductCAResponse, err error) {
    return c.DescribeProductCAWithContext(context.Background(), request)
}

// DescribeProductCA
// 查询产品绑定的CA证书
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProductCAWithContext(ctx context.Context, request *DescribeProductCARequest) (response *DescribeProductCAResponse, err error) {
    if request == nil {
        request = NewDescribeProductCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductCAResponse()
    err = c.Send(request, response)
    return
}

func NewListLogRequest() (request *ListLogRequest) {
    request = &ListLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListLog")
    
    
    return
}

func NewListLogResponse() (response *ListLogResponse) {
    response = &ListLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListLog
// 本接口（ListLog）用于查看日志信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLog(request *ListLogRequest) (response *ListLogResponse, err error) {
    return c.ListLogWithContext(context.Background(), request)
}

// ListLog
// 本接口（ListLog）用于查看日志信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLogWithContext(ctx context.Context, request *ListLogRequest) (response *ListLogResponse, err error) {
    if request == nil {
        request = NewListLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLogResponse()
    err = c.Send(request, response)
    return
}

func NewListLogPayloadRequest() (request *ListLogPayloadRequest) {
    request = &ListLogPayloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListLogPayload")
    
    
    return
}

func NewListLogPayloadResponse() (response *ListLogPayloadResponse) {
    response = &ListLogPayloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListLogPayload
// 获取日志内容列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLogPayload(request *ListLogPayloadRequest) (response *ListLogPayloadResponse, err error) {
    return c.ListLogPayloadWithContext(context.Background(), request)
}

// ListLogPayload
// 获取日志内容列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ListLogPayloadWithContext(ctx context.Context, request *ListLogPayloadRequest) (response *ListLogPayloadResponse, err error) {
    if request == nil {
        request = NewListLogPayloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLogPayload require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLogPayloadResponse()
    err = c.Send(request, response)
    return
}

func NewListSDKLogRequest() (request *ListSDKLogRequest) {
    request = &ListSDKLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "ListSDKLog")
    
    
    return
}

func NewListSDKLogResponse() (response *ListSDKLogResponse) {
    response = &ListSDKLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListSDKLog
// 获取设备上报的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) ListSDKLog(request *ListSDKLogRequest) (response *ListSDKLogResponse, err error) {
    return c.ListSDKLogWithContext(context.Background(), request)
}

// ListSDKLog
// 获取设备上报的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) ListSDKLogWithContext(ctx context.Context, request *ListSDKLogRequest) (response *ListSDKLogResponse, err error) {
    if request == nil {
        request = NewListSDKLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSDKLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSDKLogResponse()
    err = c.Send(request, response)
    return
}

func NewPublishBroadcastMessageRequest() (request *PublishBroadcastMessageRequest) {
    request = &PublishBroadcastMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "PublishBroadcastMessage")
    
    
    return
}

func NewPublishBroadcastMessageResponse() (response *PublishBroadcastMessageResponse) {
    response = &PublishBroadcastMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishBroadcastMessage
// 发布广播消息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) PublishBroadcastMessage(request *PublishBroadcastMessageRequest) (response *PublishBroadcastMessageResponse, err error) {
    return c.PublishBroadcastMessageWithContext(context.Background(), request)
}

// PublishBroadcastMessage
// 发布广播消息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) PublishBroadcastMessageWithContext(ctx context.Context, request *PublishBroadcastMessageRequest) (response *PublishBroadcastMessageResponse, err error) {
    if request == nil {
        request = NewPublishBroadcastMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishBroadcastMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishBroadcastMessageResponse()
    err = c.Send(request, response)
    return
}

func NewSetProductsForbiddenStatusRequest() (request *SetProductsForbiddenStatusRequest) {
    request = &SetProductsForbiddenStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "SetProductsForbiddenStatus")
    
    
    return
}

func NewSetProductsForbiddenStatusResponse() (response *SetProductsForbiddenStatusResponse) {
    response = &SetProductsForbiddenStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetProductsForbiddenStatus
// 批量设置产品禁用状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) SetProductsForbiddenStatus(request *SetProductsForbiddenStatusRequest) (response *SetProductsForbiddenStatusResponse, err error) {
    return c.SetProductsForbiddenStatusWithContext(context.Background(), request)
}

// SetProductsForbiddenStatus
// 批量设置产品禁用状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) SetProductsForbiddenStatusWithContext(ctx context.Context, request *SetProductsForbiddenStatusRequest) (response *SetProductsForbiddenStatusResponse, err error) {
    if request == nil {
        request = NewSetProductsForbiddenStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetProductsForbiddenStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetProductsForbiddenStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceLogLevelRequest() (request *UpdateDeviceLogLevelRequest) {
    request = &UpdateDeviceLogLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDeviceLogLevel")
    
    
    return
}

func NewUpdateDeviceLogLevelResponse() (response *UpdateDeviceLogLevelResponse) {
    response = &UpdateDeviceLogLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceLogLevel
// 设置设备上报的日志级别  
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) UpdateDeviceLogLevel(request *UpdateDeviceLogLevelRequest) (response *UpdateDeviceLogLevelResponse, err error) {
    return c.UpdateDeviceLogLevelWithContext(context.Background(), request)
}

// UpdateDeviceLogLevel
// 设置设备上报的日志级别  
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) UpdateDeviceLogLevelWithContext(ctx context.Context, request *UpdateDeviceLogLevelRequest) (response *UpdateDeviceLogLevelResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceLogLevelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDeviceLogLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceLogLevelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDevicePSKRequest() (request *UpdateDevicePSKRequest) {
    request = &UpdateDevicePSKRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDevicePSK")
    
    
    return
}

func NewUpdateDevicePSKResponse() (response *UpdateDevicePSKResponse) {
    response = &UpdateDevicePSKResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevicePSK
// 本接口（UpdateDevicePSK）用于更新设备的PSK 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
func (c *Client) UpdateDevicePSK(request *UpdateDevicePSKRequest) (response *UpdateDevicePSKResponse, err error) {
    return c.UpdateDevicePSKWithContext(context.Background(), request)
}

// UpdateDevicePSK
// 本接口（UpdateDevicePSK）用于更新设备的PSK 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
func (c *Client) UpdateDevicePSKWithContext(ctx context.Context, request *UpdateDevicePSKRequest) (response *UpdateDevicePSKResponse, err error) {
    if request == nil {
        request = NewUpdateDevicePSKRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDevicePSK require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDevicePSKResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDevicesEnableStateRequest() (request *UpdateDevicesEnableStateRequest) {
    request = &UpdateDevicesEnableStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDevicesEnableState")
    
    
    return
}

func NewUpdateDevicesEnableStateResponse() (response *UpdateDevicesEnableStateResponse) {
    response = &UpdateDevicesEnableStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevicesEnableState
// 批量启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
func (c *Client) UpdateDevicesEnableState(request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    return c.UpdateDevicesEnableStateWithContext(context.Background(), request)
}

// UpdateDevicesEnableState
// 批量启用或者禁用设备 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
func (c *Client) UpdateDevicesEnableStateWithContext(ctx context.Context, request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDevicesEnableStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDevicesEnableState require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDevicesEnableStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrivateCARequest() (request *UpdatePrivateCARequest) {
    request = &UpdatePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdatePrivateCA")
    
    
    return
}

func NewUpdatePrivateCAResponse() (response *UpdatePrivateCAResponse) {
    response = &UpdatePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdatePrivateCA
// 更新私有CA证书
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) UpdatePrivateCA(request *UpdatePrivateCARequest) (response *UpdatePrivateCAResponse, err error) {
    return c.UpdatePrivateCAWithContext(context.Background(), request)
}

// UpdatePrivateCA
// 更新私有CA证书
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) UpdatePrivateCAWithContext(ctx context.Context, request *UpdatePrivateCARequest) (response *UpdatePrivateCAResponse, err error) {
    if request == nil {
        request = NewUpdatePrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProductDynamicRegisterRequest() (request *UpdateProductDynamicRegisterRequest) {
    request = &UpdateProductDynamicRegisterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateProductDynamicRegister")
    
    
    return
}

func NewUpdateProductDynamicRegisterResponse() (response *UpdateProductDynamicRegisterResponse) {
    response = &UpdateProductDynamicRegisterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProductDynamicRegister
// 更新产品动态注册的配置 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateProductDynamicRegister(request *UpdateProductDynamicRegisterRequest) (response *UpdateProductDynamicRegisterResponse, err error) {
    return c.UpdateProductDynamicRegisterWithContext(context.Background(), request)
}

// UpdateProductDynamicRegister
// 更新产品动态注册的配置 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"
func (c *Client) UpdateProductDynamicRegisterWithContext(ctx context.Context, request *UpdateProductDynamicRegisterRequest) (response *UpdateProductDynamicRegisterResponse, err error) {
    if request == nil {
        request = NewUpdateProductDynamicRegisterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProductDynamicRegister require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProductDynamicRegisterResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProductPrivateCARequest() (request *UpdateProductPrivateCARequest) {
    request = &UpdateProductPrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateProductPrivateCA")
    
    
    return
}

func NewUpdateProductPrivateCAResponse() (response *UpdateProductPrivateCAResponse) {
    response = &UpdateProductPrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProductPrivateCA
// 更新产品的私有CA
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CACERTNOTSUPPORT = "LimitExceeded.CACertNotSupport"
func (c *Client) UpdateProductPrivateCA(request *UpdateProductPrivateCARequest) (response *UpdateProductPrivateCAResponse, err error) {
    return c.UpdateProductPrivateCAWithContext(context.Background(), request)
}

// UpdateProductPrivateCA
// 更新产品的私有CA
//
// 可能返回的错误码:
//  LIMITEXCEEDED_CACERTNOTSUPPORT = "LimitExceeded.CACertNotSupport"
func (c *Client) UpdateProductPrivateCAWithContext(ctx context.Context, request *UpdateProductPrivateCARequest) (response *UpdateProductPrivateCAResponse, err error) {
    if request == nil {
        request = NewUpdateProductPrivateCARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProductPrivateCA require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProductPrivateCAResponse()
    err = c.Send(request, response)
    return
}
