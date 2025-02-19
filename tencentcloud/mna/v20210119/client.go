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


func NewActivateHardwareRequest() (request *ActivateHardwareRequest) {
    request = &ActivateHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "ActivateHardware")
    
    
    return
}

func NewActivateHardwareResponse() (response *ActivateHardwareResponse) {
    response = &ActivateHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActivateHardware
// 激活硬件设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWAREHASACTIVATED = "OperationDenied.HardwareHasActivated"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) ActivateHardware(request *ActivateHardwareRequest) (response *ActivateHardwareResponse, err error) {
    return c.ActivateHardwareWithContext(context.Background(), request)
}

// ActivateHardware
// 激活硬件设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_HARDWAREHASACTIVATED = "OperationDenied.HardwareHasActivated"
//  OPERATIONDENIED_HARDWARENOTEXIST = "OperationDenied.HardwareNotExist"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) ActivateHardwareWithContext(ctx context.Context, request *ActivateHardwareRequest) (response *ActivateHardwareResponse, err error) {
    if request == nil {
        request = NewActivateHardwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateHardwareResponse()
    err = c.Send(request, response)
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

func NewAddGroupRequest() (request *AddGroupRequest) {
    request = &AddGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddGroup")
    
    
    return
}

func NewAddGroupResponse() (response *AddGroupResponse) {
    response = &AddGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddGroup
// 新建分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddGroup(request *AddGroupRequest) (response *AddGroupResponse, err error) {
    return c.AddGroupWithContext(context.Background(), request)
}

// AddGroup
// 新建分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) AddGroupWithContext(ctx context.Context, request *AddGroupRequest) (response *AddGroupResponse, err error) {
    if request == nil {
        request = NewAddGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAddHardwareRequest() (request *AddHardwareRequest) {
    request = &AddHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddHardware")
    
    
    return
}

func NewAddHardwareResponse() (response *AddHardwareResponse) {
    response = &AddHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddHardware
// 添加硬件设备，生成未激活的硬件设备，可支持批量添加
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) AddHardware(request *AddHardwareRequest) (response *AddHardwareResponse, err error) {
    return c.AddHardwareWithContext(context.Background(), request)
}

// AddHardware
// 添加硬件设备，生成未激活的硬件设备，可支持批量添加
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) AddHardwareWithContext(ctx context.Context, request *AddHardwareRequest) (response *AddHardwareResponse, err error) {
    if request == nil {
        request = NewAddHardwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddHardwareResponse()
    err = c.Send(request, response)
    return
}

func NewAddL3ConnRequest() (request *AddL3ConnRequest) {
    request = &AddL3ConnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "AddL3Conn")
    
    
    return
}

func NewAddL3ConnResponse() (response *AddL3ConnResponse) {
    response = &AddL3ConnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddL3Conn
// 新建互通规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_L3CIDROVERLAP = "OperationDenied.L3CidrOverLap"
//  OPERATIONDENIED_L3CONNECTIONOVERSIZE = "OperationDenied.L3ConnectionOverSize"
func (c *Client) AddL3Conn(request *AddL3ConnRequest) (response *AddL3ConnResponse, err error) {
    return c.AddL3ConnWithContext(context.Background(), request)
}

// AddL3Conn
// 新建互通规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_L3CIDROVERLAP = "OperationDenied.L3CidrOverLap"
//  OPERATIONDENIED_L3CONNECTIONOVERSIZE = "OperationDenied.L3ConnectionOverSize"
func (c *Client) AddL3ConnWithContext(ctx context.Context, request *AddL3ConnRequest) (response *AddL3ConnResponse, err error) {
    if request == nil {
        request = NewAddL3ConnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddL3Conn require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddL3ConnResponse()
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
// 产品下线
//
// 
//
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
// 产品下线
//
// 
//
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

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroup
// 删除分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// 删除分组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL3ConnRequest() (request *DeleteL3ConnRequest) {
    request = &DeleteL3ConnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DeleteL3Conn")
    
    
    return
}

func NewDeleteL3ConnResponse() (response *DeleteL3ConnResponse) {
    response = &DeleteL3ConnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteL3Conn
// 删除互通规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteL3Conn(request *DeleteL3ConnRequest) (response *DeleteL3ConnResponse, err error) {
    return c.DeleteL3ConnWithContext(context.Background(), request)
}

// DeleteL3Conn
// 删除互通规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteL3ConnWithContext(ctx context.Context, request *DeleteL3ConnRequest) (response *DeleteL3ConnResponse, err error) {
    if request == nil {
        request = NewDeleteL3ConnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL3Conn require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL3ConnResponse()
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
// 产品下线
//
// 
//
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
// 产品下线
//
// 
//
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
// 产品下线
//
// 
//
// 获取Qos加速状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQos(request *DescribeQosRequest) (response *DescribeQosResponse, err error) {
    return c.DescribeQosWithContext(context.Background(), request)
}

// DescribeQos
// 产品下线
//
// 
//
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

func NewDownloadActiveDeviceCountRequest() (request *DownloadActiveDeviceCountRequest) {
    request = &DownloadActiveDeviceCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "DownloadActiveDeviceCount")
    
    
    return
}

func NewDownloadActiveDeviceCountResponse() (response *DownloadActiveDeviceCountResponse) {
    response = &DownloadActiveDeviceCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadActiveDeviceCount
// 下载活跃设备数量统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DownloadActiveDeviceCount(request *DownloadActiveDeviceCountRequest) (response *DownloadActiveDeviceCountResponse, err error) {
    return c.DownloadActiveDeviceCountWithContext(context.Background(), request)
}

// DownloadActiveDeviceCount
// 下载活跃设备数量统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DownloadActiveDeviceCountWithContext(ctx context.Context, request *DownloadActiveDeviceCountRequest) (response *DownloadActiveDeviceCountResponse, err error) {
    if request == nil {
        request = NewDownloadActiveDeviceCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadActiveDeviceCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadActiveDeviceCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetActiveDeviceCountRequest() (request *GetActiveDeviceCountRequest) {
    request = &GetActiveDeviceCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetActiveDeviceCount")
    
    
    return
}

func NewGetActiveDeviceCountResponse() (response *GetActiveDeviceCountResponse) {
    response = &GetActiveDeviceCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetActiveDeviceCount
// 活跃设备数量统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetActiveDeviceCount(request *GetActiveDeviceCountRequest) (response *GetActiveDeviceCountResponse, err error) {
    return c.GetActiveDeviceCountWithContext(context.Background(), request)
}

// GetActiveDeviceCount
// 活跃设备数量统计
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetActiveDeviceCountWithContext(ctx context.Context, request *GetActiveDeviceCountRequest) (response *GetActiveDeviceCountResponse, err error) {
    if request == nil {
        request = NewGetActiveDeviceCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetActiveDeviceCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetActiveDeviceCountResponse()
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

func NewGetDevicePayModeRequest() (request *GetDevicePayModeRequest) {
    request = &GetDevicePayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetDevicePayMode")
    
    
    return
}

func NewGetDevicePayModeResponse() (response *GetDevicePayModeResponse) {
    response = &GetDevicePayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDevicePayMode
// 获取设备付费模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetDevicePayMode(request *GetDevicePayModeRequest) (response *GetDevicePayModeResponse, err error) {
    return c.GetDevicePayModeWithContext(context.Background(), request)
}

// GetDevicePayMode
// 获取设备付费模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetDevicePayModeWithContext(ctx context.Context, request *GetDevicePayModeRequest) (response *GetDevicePayModeResponse, err error) {
    if request == nil {
        request = NewGetDevicePayModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDevicePayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDevicePayModeResponse()
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

func NewGetFlowAlarmInfoRequest() (request *GetFlowAlarmInfoRequest) {
    request = &GetFlowAlarmInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowAlarmInfo")
    
    
    return
}

func NewGetFlowAlarmInfoResponse() (response *GetFlowAlarmInfoResponse) {
    response = &GetFlowAlarmInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowAlarmInfo
// 根据AppId查询用户设置的流量告警信息，包括阈值，回调url和key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) GetFlowAlarmInfo(request *GetFlowAlarmInfoRequest) (response *GetFlowAlarmInfoResponse, err error) {
    return c.GetFlowAlarmInfoWithContext(context.Background(), request)
}

// GetFlowAlarmInfo
// 根据AppId查询用户设置的流量告警信息，包括阈值，回调url和key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) GetFlowAlarmInfoWithContext(ctx context.Context, request *GetFlowAlarmInfoRequest) (response *GetFlowAlarmInfoResponse, err error) {
    if request == nil {
        request = NewGetFlowAlarmInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowAlarmInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowAlarmInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowPackagesRequest() (request *GetFlowPackagesRequest) {
    request = &GetFlowPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowPackages")
    
    
    return
}

func NewGetFlowPackagesResponse() (response *GetFlowPackagesResponse) {
    response = &GetFlowPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowPackages
// 获取流量包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetFlowPackages(request *GetFlowPackagesRequest) (response *GetFlowPackagesResponse, err error) {
    return c.GetFlowPackagesWithContext(context.Background(), request)
}

// GetFlowPackages
// 获取流量包列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetFlowPackagesWithContext(ctx context.Context, request *GetFlowPackagesRequest) (response *GetFlowPackagesResponse, err error) {
    if request == nil {
        request = NewGetFlowPackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowPackagesResponse()
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
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
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
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
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

func NewGetFlowStatisticByGroupRequest() (request *GetFlowStatisticByGroupRequest) {
    request = &GetFlowStatisticByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowStatisticByGroup")
    
    
    return
}

func NewGetFlowStatisticByGroupResponse() (response *GetFlowStatisticByGroupResponse) {
    response = &GetFlowStatisticByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowStatisticByGroup
// 获取指定分组，指定时间数据流量使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByGroup(request *GetFlowStatisticByGroupRequest) (response *GetFlowStatisticByGroupResponse, err error) {
    return c.GetFlowStatisticByGroupWithContext(context.Background(), request)
}

// GetFlowStatisticByGroup
// 获取指定分组，指定时间数据流量使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByGroupWithContext(ctx context.Context, request *GetFlowStatisticByGroupRequest) (response *GetFlowStatisticByGroupResponse, err error) {
    if request == nil {
        request = NewGetFlowStatisticByGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowStatisticByGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowStatisticByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewGetFlowStatisticByRegionRequest() (request *GetFlowStatisticByRegionRequest) {
    request = &GetFlowStatisticByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetFlowStatisticByRegion")
    
    
    return
}

func NewGetFlowStatisticByRegionResponse() (response *GetFlowStatisticByRegionResponse) {
    response = &GetFlowStatisticByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetFlowStatisticByRegion
// 获取指定区域，指定时间点数据流量使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByRegion(request *GetFlowStatisticByRegionRequest) (response *GetFlowStatisticByRegionResponse, err error) {
    return c.GetFlowStatisticByRegionWithContext(context.Background(), request)
}

// GetFlowStatisticByRegion
// 获取指定区域，指定时间点数据流量使用情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetFlowStatisticByRegionWithContext(ctx context.Context, request *GetFlowStatisticByRegionRequest) (response *GetFlowStatisticByRegionResponse, err error) {
    if request == nil {
        request = NewGetFlowStatisticByRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFlowStatisticByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFlowStatisticByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupDetailRequest() (request *GetGroupDetailRequest) {
    request = &GetGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetGroupDetail")
    
    
    return
}

func NewGetGroupDetailResponse() (response *GetGroupDetailResponse) {
    response = &GetGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetGroupDetail
// 查看分组详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupDetail(request *GetGroupDetailRequest) (response *GetGroupDetailResponse, err error) {
    return c.GetGroupDetailWithContext(context.Background(), request)
}

// GetGroupDetail
// 查看分组详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupDetailWithContext(ctx context.Context, request *GetGroupDetailRequest) (response *GetGroupDetailResponse, err error) {
    if request == nil {
        request = NewGetGroupDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroupDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListRequest() (request *GetGroupListRequest) {
    request = &GetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetGroupList")
    
    
    return
}

func NewGetGroupListResponse() (response *GetGroupListResponse) {
    response = &GetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetGroupList
// 获取分组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    return c.GetGroupListWithContext(context.Background(), request)
}

// GetGroupList
// 获取分组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetGroupListWithContext(ctx context.Context, request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    if request == nil {
        request = NewGetGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewGetHardwareListRequest() (request *GetHardwareListRequest) {
    request = &GetHardwareListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetHardwareList")
    
    
    return
}

func NewGetHardwareListResponse() (response *GetHardwareListResponse) {
    response = &GetHardwareListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetHardwareList
// 获取厂商硬件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetHardwareList(request *GetHardwareListRequest) (response *GetHardwareListResponse, err error) {
    return c.GetHardwareListWithContext(context.Background(), request)
}

// GetHardwareList
// 获取厂商硬件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetHardwareListWithContext(ctx context.Context, request *GetHardwareListRequest) (response *GetHardwareListResponse, err error) {
    if request == nil {
        request = NewGetHardwareListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetHardwareList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetHardwareListResponse()
    err = c.Send(request, response)
    return
}

func NewGetL3ConnListRequest() (request *GetL3ConnListRequest) {
    request = &GetL3ConnListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetL3ConnList")
    
    
    return
}

func NewGetL3ConnListResponse() (response *GetL3ConnListResponse) {
    response = &GetL3ConnListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetL3ConnList
// 获取互通规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetL3ConnList(request *GetL3ConnListRequest) (response *GetL3ConnListResponse, err error) {
    return c.GetL3ConnListWithContext(context.Background(), request)
}

// GetL3ConnList
// 获取互通规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetL3ConnListWithContext(ctx context.Context, request *GetL3ConnListRequest) (response *GetL3ConnListResponse, err error) {
    if request == nil {
        request = NewGetL3ConnListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetL3ConnList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetL3ConnListResponse()
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

func NewGetNetMonitorRequest() (request *GetNetMonitorRequest) {
    request = &GetNetMonitorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetNetMonitor")
    
    
    return
}

func NewGetNetMonitorResponse() (response *GetNetMonitorResponse) {
    response = &GetNetMonitorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetNetMonitor
// 获取单设备的实时流量统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetNetMonitor(request *GetNetMonitorRequest) (response *GetNetMonitorResponse, err error) {
    return c.GetNetMonitorWithContext(context.Background(), request)
}

// GetNetMonitor
// 获取单设备的实时流量统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INTERNALERROR_NETWORKINFOREQUESTERROR = "InternalError.NetworkInfoRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetNetMonitorWithContext(ctx context.Context, request *GetNetMonitorRequest) (response *GetNetMonitorResponse, err error) {
    if request == nil {
        request = NewGetNetMonitorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetNetMonitor require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetNetMonitorResponse()
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
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
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
//  INTERNALERROR_MONITORDATAREQUESTERROR = "InternalError.MonitorDataRequestError"
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

func NewGetVendorHardwareRequest() (request *GetVendorHardwareRequest) {
    request = &GetVendorHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GetVendorHardware")
    
    
    return
}

func NewGetVendorHardwareResponse() (response *GetVendorHardwareResponse) {
    response = &GetVendorHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetVendorHardware
// 获取厂商硬件设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetVendorHardware(request *GetVendorHardwareRequest) (response *GetVendorHardwareResponse, err error) {
    return c.GetVendorHardwareWithContext(context.Background(), request)
}

// GetVendorHardware
// 获取厂商硬件设备列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) GetVendorHardwareWithContext(ctx context.Context, request *GetVendorHardwareRequest) (response *GetVendorHardwareResponse, err error) {
    if request == nil {
        request = NewGetVendorHardwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetVendorHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetVendorHardwareResponse()
    err = c.Send(request, response)
    return
}

func NewGroupAddDeviceRequest() (request *GroupAddDeviceRequest) {
    request = &GroupAddDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GroupAddDevice")
    
    
    return
}

func NewGroupAddDeviceResponse() (response *GroupAddDeviceResponse) {
    response = &GroupAddDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GroupAddDevice
// 向已存在分组中添加设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupAddDevice(request *GroupAddDeviceRequest) (response *GroupAddDeviceResponse, err error) {
    return c.GroupAddDeviceWithContext(context.Background(), request)
}

// GroupAddDevice
// 向已存在分组中添加设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupAddDeviceWithContext(ctx context.Context, request *GroupAddDeviceRequest) (response *GroupAddDeviceResponse, err error) {
    if request == nil {
        request = NewGroupAddDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupAddDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupAddDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewGroupDeleteDeviceRequest() (request *GroupDeleteDeviceRequest) {
    request = &GroupDeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "GroupDeleteDevice")
    
    
    return
}

func NewGroupDeleteDeviceResponse() (response *GroupDeleteDeviceResponse) {
    response = &GroupDeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GroupDeleteDevice
// 删除分组中的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupDeleteDevice(request *GroupDeleteDeviceRequest) (response *GroupDeleteDeviceResponse, err error) {
    return c.GroupDeleteDeviceWithContext(context.Background(), request)
}

// GroupDeleteDevice
// 删除分组中的设备
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GroupDeleteDeviceWithContext(ctx context.Context, request *GroupDeleteDeviceRequest) (response *GroupDeleteDeviceResponse, err error) {
    if request == nil {
        request = NewGroupDeleteDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupDeleteDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPackageRenewFlagRequest() (request *ModifyPackageRenewFlagRequest) {
    request = &ModifyPackageRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "ModifyPackageRenewFlag")
    
    
    return
}

func NewModifyPackageRenewFlagResponse() (response *ModifyPackageRenewFlagResponse) {
    response = &ModifyPackageRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPackageRenewFlag
// 可开启/关闭流量包自动续费，不影响当前周期正在生效的流量包。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
//  OPERATIONDENIED_MODIFIEDORRENEWED = "OperationDenied.ModifiedOrRenewed"
//  OPERATIONDENIED_TRUNCFLAGON = "OperationDenied.TruncFlagOn"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyPackageRenewFlag(request *ModifyPackageRenewFlagRequest) (response *ModifyPackageRenewFlagResponse, err error) {
    return c.ModifyPackageRenewFlagWithContext(context.Background(), request)
}

// ModifyPackageRenewFlag
// 可开启/关闭流量包自动续费，不影响当前周期正在生效的流量包。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
//  OPERATIONDENIED_MODIFIEDORRENEWED = "OperationDenied.ModifiedOrRenewed"
//  OPERATIONDENIED_TRUNCFLAGON = "OperationDenied.TruncFlagOn"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyPackageRenewFlagWithContext(ctx context.Context, request *ModifyPackageRenewFlagRequest) (response *ModifyPackageRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyPackageRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPackageRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPackageRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewOrderFlowPackageRequest() (request *OrderFlowPackageRequest) {
    request = &OrderFlowPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "OrderFlowPackage")
    
    
    return
}

func NewOrderFlowPackageResponse() (response *OrderFlowPackageResponse) {
    response = &OrderFlowPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OrderFlowPackage
// 购买预付费流量包
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNOPENEDLIVESERVICE = "UnauthorizedOperation.UnopenedLiveService"
func (c *Client) OrderFlowPackage(request *OrderFlowPackageRequest) (response *OrderFlowPackageResponse, err error) {
    return c.OrderFlowPackageWithContext(context.Background(), request)
}

// OrderFlowPackage
// 购买预付费流量包
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNOPENEDLIVESERVICE = "UnauthorizedOperation.UnopenedLiveService"
func (c *Client) OrderFlowPackageWithContext(ctx context.Context, request *OrderFlowPackageRequest) (response *OrderFlowPackageResponse, err error) {
    if request == nil {
        request = NewOrderFlowPackageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OrderFlowPackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewOrderFlowPackageResponse()
    err = c.Send(request, response)
    return
}

func NewOrderPerLicenseRequest() (request *OrderPerLicenseRequest) {
    request = &OrderPerLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "OrderPerLicense")
    
    
    return
}

func NewOrderPerLicenseResponse() (response *OrderPerLicenseResponse) {
    response = &OrderPerLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OrderPerLicense
// 购买一次性授权License
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_REPEATPURCHASE = "OperationDenied.RepeatPurchase"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) OrderPerLicense(request *OrderPerLicenseRequest) (response *OrderPerLicenseResponse, err error) {
    return c.OrderPerLicenseWithContext(context.Background(), request)
}

// OrderPerLicense
// 购买一次性授权License
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRANSACTIONEXCEPTION = "FailedOperation.TransactionException"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DEVICENOTFOUND = "OperationDenied.DeviceNotFound"
//  OPERATIONDENIED_INSUFFICIENTBALANCE = "OperationDenied.InsufficientBalance"
//  OPERATIONDENIED_NOTALLOWEDTOPAY = "OperationDenied.NotAllowedToPay"
//  OPERATIONDENIED_REPEATPURCHASE = "OperationDenied.RepeatPurchase"
//  OPERATIONDENIED_UNAUTHORIZEDUSER = "OperationDenied.UnauthorizedUser"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) OrderPerLicenseWithContext(ctx context.Context, request *OrderPerLicenseRequest) (response *OrderPerLicenseResponse, err error) {
    if request == nil {
        request = NewOrderPerLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OrderPerLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewOrderPerLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewSetNotifyUrlRequest() (request *SetNotifyUrlRequest) {
    request = &SetNotifyUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "SetNotifyUrl")
    
    
    return
}

func NewSetNotifyUrlResponse() (response *SetNotifyUrlResponse) {
    response = &SetNotifyUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetNotifyUrl
// 设置用户流量告警信息接口，通过该接口设置流量包告警阈值以及告警时回调的url和key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) SetNotifyUrl(request *SetNotifyUrlRequest) (response *SetNotifyUrlResponse, err error) {
    return c.SetNotifyUrlWithContext(context.Background(), request)
}

// SetNotifyUrl
// 设置用户流量告警信息接口，通过该接口设置流量包告警阈值以及告警时回调的url和key
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ILLEGALREQUEST = "OperationDenied.IllegalRequest"
func (c *Client) SetNotifyUrlWithContext(ctx context.Context, request *SetNotifyUrlRequest) (response *SetNotifyUrlResponse, err error) {
    if request == nil {
        request = NewSetNotifyUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetNotifyUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetNotifyUrlResponse()
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
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
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
//  INTERNALERROR_DUPLICATEDEVICENAME = "InternalError.DuplicateDeviceName"
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

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
    request = &UpdateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateGroup")
    
    
    return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
    response = &UpdateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGroup
// 更新分组备注
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    return c.UpdateGroupWithContext(context.Background(), request)
}

// UpdateGroup
// 更新分组备注
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    if request == nil {
        request = NewUpdateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateHardwareRequest() (request *UpdateHardwareRequest) {
    request = &UpdateHardwareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateHardware")
    
    
    return
}

func NewUpdateHardwareResponse() (response *UpdateHardwareResponse) {
    response = &UpdateHardwareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateHardware
// 更新硬件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) UpdateHardware(request *UpdateHardwareRequest) (response *UpdateHardwareResponse, err error) {
    return c.UpdateHardwareWithContext(context.Background(), request)
}

// UpdateHardware
// 更新硬件信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_DUPLICATESN = "OperationDenied.DuplicateSN"
//  OPERATIONDENIED_VENDORNOTREGISTER = "OperationDenied.VendorNotRegister"
func (c *Client) UpdateHardwareWithContext(ctx context.Context, request *UpdateHardwareRequest) (response *UpdateHardwareResponse, err error) {
    if request == nil {
        request = NewUpdateHardwareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateHardware require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateHardwareResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateL3CidrRequest() (request *UpdateL3CidrRequest) {
    request = &UpdateL3CidrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateL3Cidr")
    
    
    return
}

func NewUpdateL3CidrResponse() (response *UpdateL3CidrResponse) {
    response = &UpdateL3CidrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateL3Cidr
// 更新互通规则CIDR
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3Cidr(request *UpdateL3CidrRequest) (response *UpdateL3CidrResponse, err error) {
    return c.UpdateL3CidrWithContext(context.Background(), request)
}

// UpdateL3Cidr
// 更新互通规则CIDR
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3CidrWithContext(ctx context.Context, request *UpdateL3CidrRequest) (response *UpdateL3CidrResponse, err error) {
    if request == nil {
        request = NewUpdateL3CidrRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateL3Cidr require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateL3CidrResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateL3ConnRequest() (request *UpdateL3ConnRequest) {
    request = &UpdateL3ConnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateL3Conn")
    
    
    return
}

func NewUpdateL3ConnResponse() (response *UpdateL3ConnResponse) {
    response = &UpdateL3ConnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateL3Conn
// 更新互通规则备注
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3Conn(request *UpdateL3ConnRequest) (response *UpdateL3ConnResponse, err error) {
    return c.UpdateL3ConnWithContext(context.Background(), request)
}

// UpdateL3Conn
// 更新互通规则备注
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3ConnWithContext(ctx context.Context, request *UpdateL3ConnRequest) (response *UpdateL3ConnResponse, err error) {
    if request == nil {
        request = NewUpdateL3ConnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateL3Conn require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateL3ConnResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateL3SwitchRequest() (request *UpdateL3SwitchRequest) {
    request = &UpdateL3SwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mna", APIVersion, "UpdateL3Switch")
    
    
    return
}

func NewUpdateL3SwitchResponse() (response *UpdateL3SwitchResponse) {
    response = &UpdateL3SwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateL3Switch
// 更新互通规则开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3Switch(request *UpdateL3SwitchRequest) (response *UpdateL3SwitchResponse, err error) {
    return c.UpdateL3SwitchWithContext(context.Background(), request)
}

// UpdateL3Switch
// 更新互通规则开关
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONTROLREQUESTERROR = "InternalError.ControlRequestError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateL3SwitchWithContext(ctx context.Context, request *UpdateL3SwitchRequest) (response *UpdateL3SwitchResponse, err error) {
    if request == nil {
        request = NewUpdateL3SwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateL3Switch require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateL3SwitchResponse()
    err = c.Send(request, response)
    return
}
