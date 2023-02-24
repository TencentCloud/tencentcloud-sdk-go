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

package v20210119

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-01-19"

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


func NewAddDeviceRequest() (request *AddDeviceRequest) {
    request = &AddDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddDevice")
    
    
    return
}

func NewAddDeviceResponse() (response *AddDeviceResponse) {
    response = &AddDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddDevice
// 新建设备记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDATAKEY = "InternalError.DuplicateDataKey"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INTERNALERROR_UNDEFINEDENCRYPTEDKEY = "InternalError.UndefinedEncryptedKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddDevice(request *AddDeviceRequest) (response *AddDeviceResponse, err error) {
    return c.AddDeviceWithContext(context.Background(), request)
}

// AddDevice
// 新建设备记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDATAKEY = "InternalError.DuplicateDataKey"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INTERNALERROR_UNDEFINEDENCRYPTEDKEY = "InternalError.UndefinedEncryptedKey"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddDeviceWithContext(ctx context.Context, request *AddDeviceRequest) (response *AddDeviceResponse, err error) {
    if request == nil {
        request = NewAddDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEncryptedKeyRequest() (request *CreateEncryptedKeyRequest) {
    request = &CreateEncryptedKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "CreateEncryptedKey")
    
    
    return
}

func NewCreateEncryptedKeyResponse() (response *CreateEncryptedKeyResponse) {
    response = &CreateEncryptedKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEncryptedKey
// 通过此接口设置和更新预置密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateEncryptedKey(request *CreateEncryptedKeyRequest) (response *CreateEncryptedKeyResponse, err error) {
    return c.CreateEncryptedKeyWithContext(context.Background(), request)
}

// CreateEncryptedKey
// 通过此接口设置和更新预置密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateEncryptedKeyWithContext(ctx context.Context, request *CreateEncryptedKeyRequest) (response *CreateEncryptedKeyResponse, err error) {
    if request == nil {
        request = NewCreateEncryptedKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEncryptedKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEncryptedKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQosRequest() (request *CreateQosRequest) {
    request = &CreateQosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "CreateQos")
    
    
    return
}

func NewCreateQosResponse() (response *CreateQosResponse) {
    response = &CreateQosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateQos
// 移动网络发起Qos加速过程
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VENDORNOTFOUND = "InvalidParameterValue.VendorNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATIONNOTSUGGEST = "OperationDenied.AccelerationNotSuggest"
//  OPERATIONDENIED_CTCCTOKENEXPIRED = "OperationDenied.CTCCTokenExpired"
//  OPERATIONDENIED_CREATEQOSEXCEEDLIMIT = "OperationDenied.CreateQosExceedLimit"
//  OPERATIONDENIED_REQUESTQOSTIMEOUT = "OperationDenied.RequestQosTimeout"
//  OPERATIONDENIED_USEROUTOFCOVERAGE = "OperationDenied.UserOutOfCoverage"
//  OPERATIONDENIED_VENDORRETURNERROR = "OperationDenied.VendorReturnError"
//  OPERATIONDENIED_VENDORSERVERERROR = "OperationDenied.VendorServerError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateQos(request *CreateQosRequest) (response *CreateQosResponse, err error) {
    return c.CreateQosWithContext(context.Background(), request)
}

// CreateQos
// 移动网络发起Qos加速过程
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_VENDORNOTFOUND = "InvalidParameterValue.VendorNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATIONNOTSUGGEST = "OperationDenied.AccelerationNotSuggest"
//  OPERATIONDENIED_CTCCTOKENEXPIRED = "OperationDenied.CTCCTokenExpired"
//  OPERATIONDENIED_CREATEQOSEXCEEDLIMIT = "OperationDenied.CreateQosExceedLimit"
//  OPERATIONDENIED_REQUESTQOSTIMEOUT = "OperationDenied.RequestQosTimeout"
//  OPERATIONDENIED_USEROUTOFCOVERAGE = "OperationDenied.UserOutOfCoverage"
//  OPERATIONDENIED_VENDORRETURNERROR = "OperationDenied.VendorReturnError"
//  OPERATIONDENIED_VENDORSERVERERROR = "OperationDenied.VendorServerError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateQosWithContext(ctx context.Context, request *CreateQosRequest) (response *CreateQosResponse, err error) {
    if request == nil {
        request = NewCreateQosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQos require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQosResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteDevice")
    
    
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDevice
// 删除设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// 删除设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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

func NewDeleteQosRequest() (request *DeleteQosRequest) {
    request = &DeleteQosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteQos")
    
    
    return
}

func NewDeleteQosResponse() (response *DeleteQosResponse) {
    response = &DeleteQosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteQos
// 移动网络停止Qos加速过程
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_REQUESTQOSTIMEOUT = "OperationDenied.RequestQosTimeout"
//  OPERATIONDENIED_USERNONACCELERATED = "OperationDenied.UserNonAccelerated"
//  OPERATIONDENIED_USEROUTOFCOVERAGE = "OperationDenied.UserOutOfCoverage"
//  OPERATIONDENIED_VENDORRETURNERROR = "OperationDenied.VendorReturnError"
//  OPERATIONDENIED_VENDORSERVERERROR = "OperationDenied.VendorServerError"
func (c *Client) DeleteQos(request *DeleteQosRequest) (response *DeleteQosResponse, err error) {
    return c.DeleteQosWithContext(context.Background(), request)
}

// DeleteQos
// 移动网络停止Qos加速过程
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_REQUESTQOSTIMEOUT = "OperationDenied.RequestQosTimeout"
//  OPERATIONDENIED_USERNONACCELERATED = "OperationDenied.UserNonAccelerated"
//  OPERATIONDENIED_USEROUTOFCOVERAGE = "OperationDenied.UserOutOfCoverage"
//  OPERATIONDENIED_VENDORRETURNERROR = "OperationDenied.VendorReturnError"
//  OPERATIONDENIED_VENDORSERVERERROR = "OperationDenied.VendorServerError"
func (c *Client) DeleteQosWithContext(ctx context.Context, request *DeleteQosRequest) (response *DeleteQosResponse, err error) {
    if request == nil {
        request = NewDeleteQosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQosRequest() (request *DescribeQosRequest) {
    request = &DescribeQosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DescribeQos")
    
    
    return
}

func NewDescribeQosResponse() (response *DescribeQosResponse) {
    response = &DescribeQosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeQos
// 获取Qos加速状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQos(request *DescribeQosRequest) (response *DescribeQosResponse, err error) {
    return c.DescribeQosWithContext(context.Background(), request)
}

// DescribeQos
// 获取Qos加速状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQosWithContext(ctx context.Context, request *DescribeQosRequest) (response *DescribeQosResponse, err error) {
    if request == nil {
        request = NewDescribeQosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQosResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceRequest() (request *GetDeviceRequest) {
    request = &GetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetDevice")
    
    
    return
}

func NewGetDeviceResponse() (response *GetDeviceResponse) {
    response = &GetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDevice
// 通过指定设备的ID查找设备详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevice(request *GetDeviceRequest) (response *GetDeviceResponse, err error) {
    return c.GetDeviceWithContext(context.Background(), request)
}

// GetDevice
// 通过指定设备的ID查找设备详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDeviceWithContext(ctx context.Context, request *GetDeviceRequest) (response *GetDeviceResponse, err error) {
    if request == nil {
        request = NewGetDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewGetDevicesRequest() (request *GetDevicesRequest) {
    request = &GetDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetDevices")
    
    
    return
}

func NewGetDevicesResponse() (response *GetDevicesResponse) {
    response = &GetDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDevices
// 获取设备信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevices(request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    return c.GetDevicesWithContext(context.Background(), request)
}

// GetDevices
// 获取设备信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevicesWithContext(ctx context.Context, request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    if request == nil {
        request = NewGetDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowStatisticRequest() (request *GetFlowStatisticRequest) {
    request = &GetFlowStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowStatistic")
    
    
    return
}

func NewGetFlowStatisticResponse() (response *GetFlowStatisticResponse) {
    response = &GetFlowStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFlowStatistic
// 获取指定设备Id，指定时间点数据流量使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatistic(request *GetFlowStatisticRequest) (response *GetFlowStatisticResponse, err error) {
    return c.GetFlowStatisticWithContext(context.Background(), request)
}

// GetFlowStatistic
// 获取指定设备Id，指定时间点数据流量使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticWithContext(ctx context.Context, request *GetFlowStatisticRequest) (response *GetFlowStatisticResponse, err error) {
    if request == nil {
        request = NewGetFlowStatisticRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewGetMultiFlowStatisticRequest() (request *GetMultiFlowStatisticRequest) {
    request = &GetMultiFlowStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetMultiFlowStatistic")
    
    
    return
}

func NewGetMultiFlowStatisticResponse() (response *GetMultiFlowStatisticResponse) {
    response = &GetMultiFlowStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetMultiFlowStatistic
// 批量获取设备流量统计曲线
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMultiFlowStatistic(request *GetMultiFlowStatisticRequest) (response *GetMultiFlowStatisticResponse, err error) {
    return c.GetMultiFlowStatisticWithContext(context.Background(), request)
}

// GetMultiFlowStatistic
// 批量获取设备流量统计曲线
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetMultiFlowStatisticWithContext(ctx context.Context, request *GetMultiFlowStatisticRequest) (response *GetMultiFlowStatisticResponse, err error) {
    if request == nil {
        request = NewGetMultiFlowStatisticRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMultiFlowStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMultiFlowStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewGetPublicKeyRequest() (request *GetPublicKeyRequest) {
    request = &GetPublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetPublicKey")
    
    
    return
}

func NewGetPublicKeyResponse() (response *GetPublicKeyResponse) {
    response = &GetPublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPublicKey
// 获取公钥用于验签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetPublicKey(request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    return c.GetPublicKeyWithContext(context.Background(), request)
}

// GetPublicKey
// 获取公钥用于验签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetPublicKeyWithContext(ctx context.Context, request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    if request == nil {
        request = NewGetPublicKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatisticDataRequest() (request *GetStatisticDataRequest) {
    request = &GetStatisticDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetStatisticData")
    
    
    return
}

func NewGetStatisticDataResponse() (response *GetStatisticDataResponse) {
    response = &GetStatisticDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetStatisticData
// 在用量统计页面下载流量数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetStatisticData(request *GetStatisticDataRequest) (response *GetStatisticDataResponse, err error) {
    return c.GetStatisticDataWithContext(context.Background(), request)
}

// GetStatisticData
// 在用量统计页面下载流量数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_FILEIOERROR = "InternalError.FileIOError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetStatisticDataWithContext(ctx context.Context, request *GetStatisticDataRequest) (response *GetStatisticDataResponse, err error) {
    if request == nil {
        request = NewGetStatisticDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetStatisticData require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetStatisticDataResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceRequest() (request *UpdateDeviceRequest) {
    request = &UpdateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateDevice")
    
    
    return
}

func NewUpdateDeviceResponse() (response *UpdateDeviceResponse) {
    response = &UpdateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevice
// 更新设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateDevice(request *UpdateDeviceRequest) (response *UpdateDeviceResponse, err error) {
    return c.UpdateDeviceWithContext(context.Background(), request)
}

// UpdateDevice
// 更新设备信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateDeviceWithContext(ctx context.Context, request *UpdateDeviceRequest) (response *UpdateDeviceResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDeviceResponse()
    err = c.Send(request, response)
    return
}
