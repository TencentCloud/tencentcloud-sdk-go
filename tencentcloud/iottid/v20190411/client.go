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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAuthTestTidRequest() (request *AuthTestTidRequest) {
    request = &AuthTestTidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "AuthTestTid")
    
    
    return
}

func NewAuthTestTidResponse() (response *AuthTestTidResponse) {
    response = &AuthTestTidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AuthTestTid
// 单向认证测试TID 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATA = "InvalidParameterValue.Data"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
func (c *Client) AuthTestTid(request *AuthTestTidRequest) (response *AuthTestTidResponse, err error) {
    return c.AuthTestTidWithContext(context.Background(), request)
}

// AuthTestTid
// 单向认证测试TID 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATA = "InvalidParameterValue.Data"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
func (c *Client) AuthTestTidWithContext(ctx context.Context, request *AuthTestTidRequest) (response *AuthTestTidResponse, err error) {
    if request == nil {
        request = NewAuthTestTidRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuthTestTid require credential")
    }

    request.SetContext(ctx)
    
    response = NewAuthTestTidResponse()
    err = c.Send(request, response)
    return
}

func NewBurnTidNotifyRequest() (request *BurnTidNotifyRequest) {
    request = &BurnTidNotifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "BurnTidNotify")
    
    
    return
}

func NewBurnTidNotifyResponse() (response *BurnTidNotifyResponse) {
    response = &BurnTidNotifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BurnTidNotify
// 安全芯片TID烧录回执 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
func (c *Client) BurnTidNotify(request *BurnTidNotifyRequest) (response *BurnTidNotifyResponse, err error) {
    return c.BurnTidNotifyWithContext(context.Background(), request)
}

// BurnTidNotify
// 安全芯片TID烧录回执 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
func (c *Client) BurnTidNotifyWithContext(ctx context.Context, request *BurnTidNotifyRequest) (response *BurnTidNotifyResponse, err error) {
    if request == nil {
        request = NewBurnTidNotifyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BurnTidNotify require credential")
    }

    request.SetContext(ctx)
    
    response = NewBurnTidNotifyResponse()
    err = c.Send(request, response)
    return
}

func NewDeliverTidNotifyRequest() (request *DeliverTidNotifyRequest) {
    request = &DeliverTidNotifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "DeliverTidNotify")
    
    
    return
}

func NewDeliverTidNotifyResponse() (response *DeliverTidNotifyResponse) {
    response = &DeliverTidNotifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeliverTidNotify
// 安全芯片为载体的TID空发回执，绑定TID与订单号。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_OVERLIMIT = "InvalidParameterValue.OverLimit"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
func (c *Client) DeliverTidNotify(request *DeliverTidNotifyRequest) (response *DeliverTidNotifyResponse, err error) {
    return c.DeliverTidNotifyWithContext(context.Background(), request)
}

// DeliverTidNotify
// 安全芯片为载体的TID空发回执，绑定TID与订单号。 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_OVERLIMIT = "InvalidParameterValue.OverLimit"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_TID = "InvalidParameterValue.Tid"
func (c *Client) DeliverTidNotifyWithContext(ctx context.Context, request *DeliverTidNotifyRequest) (response *DeliverTidNotifyResponse, err error) {
    if request == nil {
        request = NewDeliverTidNotifyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeliverTidNotify require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeliverTidNotifyResponse()
    err = c.Send(request, response)
    return
}

func NewDeliverTidsRequest() (request *DeliverTidsRequest) {
    request = &DeliverTidsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "DeliverTids")
    
    
    return
}

func NewDeliverTidsResponse() (response *DeliverTidsResponse) {
    response = &DeliverTidsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeliverTids
// 设备服务商请求空发产品订单的TID信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
func (c *Client) DeliverTids(request *DeliverTidsRequest) (response *DeliverTidsResponse, err error) {
    return c.DeliverTidsWithContext(context.Background(), request)
}

// DeliverTids
// 设备服务商请求空发产品订单的TID信息 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
func (c *Client) DeliverTidsWithContext(ctx context.Context, request *DeliverTidsRequest) (response *DeliverTidsResponse, err error) {
    if request == nil {
        request = NewDeliverTidsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeliverTids require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeliverTidsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableLibCountRequest() (request *DescribeAvailableLibCountRequest) {
    request = &DescribeAvailableLibCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "DescribeAvailableLibCount")
    
    
    return
}

func NewDescribeAvailableLibCountResponse() (response *DescribeAvailableLibCountResponse) {
    response = &DescribeAvailableLibCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableLibCount
// 查询指定订单的可空发的白盒密钥数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
func (c *Client) DescribeAvailableLibCount(request *DescribeAvailableLibCountRequest) (response *DescribeAvailableLibCountResponse, err error) {
    return c.DescribeAvailableLibCountWithContext(context.Background(), request)
}

// DescribeAvailableLibCount
// 查询指定订单的可空发的白盒密钥数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
func (c *Client) DescribeAvailableLibCountWithContext(ctx context.Context, request *DescribeAvailableLibCountRequest) (response *DescribeAvailableLibCountResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableLibCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableLibCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableLibCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePermissionRequest() (request *DescribePermissionRequest) {
    request = &DescribePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "DescribePermission")
    
    
    return
}

func NewDescribePermissionResponse() (response *DescribePermissionResponse) {
    response = &DescribePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePermission
// 查询企业用户TID平台控制台权限 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
func (c *Client) DescribePermission(request *DescribePermissionRequest) (response *DescribePermissionResponse, err error) {
    return c.DescribePermissionWithContext(context.Background(), request)
}

// DescribePermission
// 查询企业用户TID平台控制台权限 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
func (c *Client) DescribePermissionWithContext(ctx context.Context, request *DescribePermissionRequest) (response *DescribePermissionResponse, err error) {
    if request == nil {
        request = NewDescribePermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadTidsRequest() (request *DownloadTidsRequest) {
    request = &DownloadTidsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "DownloadTids")
    
    
    return
}

func NewDownloadTidsResponse() (response *DownloadTidsResponse) {
    response = &DownloadTidsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadTids
// 下载芯片订单的TID 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
func (c *Client) DownloadTids(request *DownloadTidsRequest) (response *DownloadTidsResponse, err error) {
    return c.DownloadTidsWithContext(context.Background(), request)
}

// DownloadTids
// 下载芯片订单的TID 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
func (c *Client) DownloadTidsWithContext(ctx context.Context, request *DownloadTidsRequest) (response *DownloadTidsResponse, err error) {
    if request == nil {
        request = NewDownloadTidsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadTids require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadTidsResponse()
    err = c.Send(request, response)
    return
}

func NewUploadDeviceUniqueCodeRequest() (request *UploadDeviceUniqueCodeRequest) {
    request = &UploadDeviceUniqueCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "UploadDeviceUniqueCode")
    
    
    return
}

func NewUploadDeviceUniqueCodeResponse() (response *UploadDeviceUniqueCodeResponse) {
    response = &UploadDeviceUniqueCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadDeviceUniqueCode
// 上传硬件唯一标识码，是软加固设备身份参数。本接口如遇到错误数据，则所有当次上传数据失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COUNT = "InvalidParameterValue.Count"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_OVERLIMIT = "InvalidParameterValue.OverLimit"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
func (c *Client) UploadDeviceUniqueCode(request *UploadDeviceUniqueCodeRequest) (response *UploadDeviceUniqueCodeResponse, err error) {
    return c.UploadDeviceUniqueCodeWithContext(context.Background(), request)
}

// UploadDeviceUniqueCode
// 上传硬件唯一标识码，是软加固设备身份参数。本接口如遇到错误数据，则所有当次上传数据失效。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_COUNT = "InvalidParameterValue.Count"
//  INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"
//  INVALIDPARAMETERVALUE_ORDERID = "InvalidParameterValue.OrderId"
//  INVALIDPARAMETERVALUE_OVERLIMIT = "InvalidParameterValue.OverLimit"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
//  INVALIDPARAMETERVALUE_QUANTITY = "InvalidParameterValue.Quantity"
func (c *Client) UploadDeviceUniqueCodeWithContext(ctx context.Context, request *UploadDeviceUniqueCodeRequest) (response *UploadDeviceUniqueCodeResponse, err error) {
    if request == nil {
        request = NewUploadDeviceUniqueCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadDeviceUniqueCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadDeviceUniqueCodeResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyChipBurnInfoRequest() (request *VerifyChipBurnInfoRequest) {
    request = &VerifyChipBurnInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iottid", APIVersion, "VerifyChipBurnInfo")
    
    
    return
}

func NewVerifyChipBurnInfoResponse() (response *VerifyChipBurnInfoResponse) {
    response = &VerifyChipBurnInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyChipBurnInfo
// 下载控制台验证芯片烧录信息，保证TID与中心信息一致 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATA = "InvalidParameterValue.Data"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
func (c *Client) VerifyChipBurnInfo(request *VerifyChipBurnInfoRequest) (response *VerifyChipBurnInfoResponse, err error) {
    return c.VerifyChipBurnInfoWithContext(context.Background(), request)
}

// VerifyChipBurnInfo
// 下载控制台验证芯片烧录信息，保证TID与中心信息一致 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATA = "InvalidParameterValue.Data"
//  INVALIDPARAMETERVALUE_EMPTYSTRING = "InvalidParameterValue.EmptyString"
//  INVALIDPARAMETERVALUE_PERMISSIONDENIED = "InvalidParameterValue.PermissionDenied"
func (c *Client) VerifyChipBurnInfoWithContext(ctx context.Context, request *VerifyChipBurnInfoRequest) (response *VerifyChipBurnInfoResponse, err error) {
    if request == nil {
        request = NewVerifyChipBurnInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyChipBurnInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyChipBurnInfoResponse()
    err = c.Send(request, response)
    return
}
