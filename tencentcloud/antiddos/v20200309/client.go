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

package v20200309

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-09"

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


func NewAssociateDDoSEipAddressRequest() (request *AssociateDDoSEipAddressRequest) {
    request = &AssociateDDoSEipAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "AssociateDDoSEipAddress")
    return
}

func NewAssociateDDoSEipAddressResponse() (response *AssociateDDoSEipAddressResponse) {
    response = &AssociateDDoSEipAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateDDoSEipAddress
// 本接口 (AssociateDDoSEipAddress) 用于将高防弹性公网IP绑定到实例或弹性网卡的指定内网 IP 上。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AssociateDDoSEipAddress(request *AssociateDDoSEipAddressRequest) (response *AssociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewAssociateDDoSEipAddressRequest()
    }
    response = NewAssociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlackWhiteIpListRequest() (request *CreateBlackWhiteIpListRequest) {
    request = &CreateBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateBlackWhiteIpList")
    return
}

func NewCreateBlackWhiteIpListResponse() (response *CreateBlackWhiteIpListResponse) {
    response = &CreateBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBlackWhiteIpList
// 添加DDoS防护的IP黑白名单
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBlackWhiteIpList(request *CreateBlackWhiteIpListRequest) (response *CreateBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewCreateBlackWhiteIpListRequest()
    }
    response = NewCreateBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBoundIPRequest() (request *CreateBoundIPRequest) {
    request = &CreateBoundIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateBoundIP")
    return
}

func NewCreateBoundIPResponse() (response *CreateBoundIPResponse) {
    response = &CreateBoundIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBoundIP
// 绑定IP到高防包实例，支持独享包、共享包；需要注意的是此接口绑定或解绑IP是异步接口，当处于绑定或解绑中时，则不允许再进行绑定或解绑，需要等待当前绑定或解绑完成。
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBoundIP(request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    if request == nil {
        request = NewCreateBoundIPRequest()
    }
    response = NewCreateBoundIPResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSAIRequest() (request *CreateDDoSAIRequest) {
    request = &CreateDDoSAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSAI")
    return
}

func NewCreateDDoSAIResponse() (response *CreateDDoSAIResponse) {
    response = &CreateDDoSAIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSAI
// 设置DDoS防护的AI防护开关
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSAI(request *CreateDDoSAIRequest) (response *CreateDDoSAIResponse, err error) {
    if request == nil {
        request = NewCreateDDoSAIRequest()
    }
    response = NewCreateDDoSAIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSGeoIPBlockConfigRequest() (request *CreateDDoSGeoIPBlockConfigRequest) {
    request = &CreateDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSGeoIPBlockConfig")
    return
}

func NewCreateDDoSGeoIPBlockConfigResponse() (response *CreateDDoSGeoIPBlockConfigResponse) {
    response = &CreateDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSGeoIPBlockConfig
// 添加DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDDoSGeoIPBlockConfig(request *CreateDDoSGeoIPBlockConfigRequest) (response *CreateDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSGeoIPBlockConfigRequest()
    }
    response = NewCreateDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSSpeedLimitConfigRequest() (request *CreateDDoSSpeedLimitConfigRequest) {
    request = &CreateDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDDoSSpeedLimitConfig")
    return
}

func NewCreateDDoSSpeedLimitConfigResponse() (response *CreateDDoSSpeedLimitConfigResponse) {
    response = &CreateDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDDoSSpeedLimitConfig
// 添加DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateDDoSSpeedLimitConfig(request *CreateDDoSSpeedLimitConfigRequest) (response *CreateDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewCreateDDoSSpeedLimitConfigRequest()
    }
    response = NewCreateDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDefaultAlarmThresholdRequest() (request *CreateDefaultAlarmThresholdRequest) {
    request = &CreateDefaultAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateDefaultAlarmThreshold")
    return
}

func NewCreateDefaultAlarmThresholdResponse() (response *CreateDefaultAlarmThresholdResponse) {
    response = &CreateDefaultAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDefaultAlarmThreshold
// 设置单IP默认告警阈值配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateDefaultAlarmThreshold(request *CreateDefaultAlarmThresholdRequest) (response *CreateDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewCreateDefaultAlarmThresholdRequest()
    }
    response = NewCreateDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIPAlarmThresholdConfigRequest() (request *CreateIPAlarmThresholdConfigRequest) {
    request = &CreateIPAlarmThresholdConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateIPAlarmThresholdConfig")
    return
}

func NewCreateIPAlarmThresholdConfigResponse() (response *CreateIPAlarmThresholdConfigResponse) {
    response = &CreateIPAlarmThresholdConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIPAlarmThresholdConfig
// 设置单IP告警阈值配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateIPAlarmThresholdConfig(request *CreateIPAlarmThresholdConfigRequest) (response *CreateIPAlarmThresholdConfigResponse, err error) {
    if request == nil {
        request = NewCreateIPAlarmThresholdConfigRequest()
    }
    response = NewCreateIPAlarmThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RuleCertsRequest() (request *CreateL7RuleCertsRequest) {
    request = &CreateL7RuleCertsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateL7RuleCerts")
    return
}

func NewCreateL7RuleCertsResponse() (response *CreateL7RuleCertsResponse) {
    response = &CreateL7RuleCertsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateL7RuleCerts
// 批量配置L7转发规则的证书供SSL测调用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateL7RuleCerts(request *CreateL7RuleCertsRequest) (response *CreateL7RuleCertsResponse, err error) {
    if request == nil {
        request = NewCreateL7RuleCertsRequest()
    }
    response = NewCreateL7RuleCertsResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePacketFilterConfigRequest() (request *CreatePacketFilterConfigRequest) {
    request = &CreatePacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreatePacketFilterConfig")
    return
}

func NewCreatePacketFilterConfigResponse() (response *CreatePacketFilterConfigResponse) {
    response = &CreatePacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePacketFilterConfig
// 添加DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePacketFilterConfig(request *CreatePacketFilterConfigRequest) (response *CreatePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewCreatePacketFilterConfigRequest()
    }
    response = NewCreatePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProtocolBlockConfigRequest() (request *CreateProtocolBlockConfigRequest) {
    request = &CreateProtocolBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateProtocolBlockConfig")
    return
}

func NewCreateProtocolBlockConfigResponse() (response *CreateProtocolBlockConfigResponse) {
    response = &CreateProtocolBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProtocolBlockConfig
// 设置DDoS防护的协议封禁配置
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProtocolBlockConfig(request *CreateProtocolBlockConfigRequest) (response *CreateProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewCreateProtocolBlockConfigRequest()
    }
    response = NewCreateProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchedulingDomainRequest() (request *CreateSchedulingDomainRequest) {
    request = &CreateSchedulingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateSchedulingDomain")
    return
}

func NewCreateSchedulingDomainResponse() (response *CreateSchedulingDomainResponse) {
    response = &CreateSchedulingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSchedulingDomain
// 创建一个域名，可用于在封堵时调度切换IP
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateSchedulingDomain(request *CreateSchedulingDomainRequest) (response *CreateSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewCreateSchedulingDomainRequest()
    }
    response = NewCreateSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWaterPrintConfigRequest() (request *CreateWaterPrintConfigRequest) {
    request = &CreateWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateWaterPrintConfig")
    return
}

func NewCreateWaterPrintConfigResponse() (response *CreateWaterPrintConfigResponse) {
    response = &CreateWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWaterPrintConfig
// 添加DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateWaterPrintConfig(request *CreateWaterPrintConfigRequest) (response *CreateWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintConfigRequest()
    }
    response = NewCreateWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWaterPrintKeyRequest() (request *CreateWaterPrintKeyRequest) {
    request = &CreateWaterPrintKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "CreateWaterPrintKey")
    return
}

func NewCreateWaterPrintKeyResponse() (response *CreateWaterPrintKeyResponse) {
    response = &CreateWaterPrintKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWaterPrintKey
// 添加DDoS防护的水印防护密钥
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateWaterPrintKey(request *CreateWaterPrintKeyRequest) (response *CreateWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewCreateWaterPrintKeyRequest()
    }
    response = NewCreateWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlackWhiteIpListRequest() (request *DeleteBlackWhiteIpListRequest) {
    request = &DeleteBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteBlackWhiteIpList")
    return
}

func NewDeleteBlackWhiteIpListResponse() (response *DeleteBlackWhiteIpListResponse) {
    response = &DeleteBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBlackWhiteIpList
// 删除DDoS防护的IP黑白名单
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBlackWhiteIpList(request *DeleteBlackWhiteIpListRequest) (response *DeleteBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackWhiteIpListRequest()
    }
    response = NewDeleteBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSGeoIPBlockConfigRequest() (request *DeleteDDoSGeoIPBlockConfigRequest) {
    request = &DeleteDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteDDoSGeoIPBlockConfig")
    return
}

func NewDeleteDDoSGeoIPBlockConfigResponse() (response *DeleteDDoSGeoIPBlockConfigResponse) {
    response = &DeleteDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSGeoIPBlockConfig
// 删除DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSGeoIPBlockConfig(request *DeleteDDoSGeoIPBlockConfigRequest) (response *DeleteDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSGeoIPBlockConfigRequest()
    }
    response = NewDeleteDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSSpeedLimitConfigRequest() (request *DeleteDDoSSpeedLimitConfigRequest) {
    request = &DeleteDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteDDoSSpeedLimitConfig")
    return
}

func NewDeleteDDoSSpeedLimitConfigResponse() (response *DeleteDDoSSpeedLimitConfigResponse) {
    response = &DeleteDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDDoSSpeedLimitConfig
// 删除DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDDoSSpeedLimitConfig(request *DeleteDDoSSpeedLimitConfigRequest) (response *DeleteDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSSpeedLimitConfigRequest()
    }
    response = NewDeleteDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePacketFilterConfigRequest() (request *DeletePacketFilterConfigRequest) {
    request = &DeletePacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeletePacketFilterConfig")
    return
}

func NewDeletePacketFilterConfigResponse() (response *DeletePacketFilterConfigResponse) {
    response = &DeletePacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePacketFilterConfig
// 删除DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePacketFilterConfig(request *DeletePacketFilterConfigRequest) (response *DeletePacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDeletePacketFilterConfigRequest()
    }
    response = NewDeletePacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWaterPrintConfigRequest() (request *DeleteWaterPrintConfigRequest) {
    request = &DeleteWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteWaterPrintConfig")
    return
}

func NewDeleteWaterPrintConfigResponse() (response *DeleteWaterPrintConfigResponse) {
    response = &DeleteWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWaterPrintConfig
// 删除DDoS防护的水印防护配置
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWaterPrintConfig(request *DeleteWaterPrintConfigRequest) (response *DeleteWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintConfigRequest()
    }
    response = NewDeleteWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWaterPrintKeyRequest() (request *DeleteWaterPrintKeyRequest) {
    request = &DeleteWaterPrintKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DeleteWaterPrintKey")
    return
}

func NewDeleteWaterPrintKeyResponse() (response *DeleteWaterPrintKeyResponse) {
    response = &DeleteWaterPrintKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWaterPrintKey
// 删除DDoS防护的水印防护密钥
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteWaterPrintKey(request *DeleteWaterPrintKeyRequest) (response *DeleteWaterPrintKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWaterPrintKeyRequest()
    }
    response = NewDeleteWaterPrintKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlackWhiteIpListRequest() (request *DescribeBlackWhiteIpListRequest) {
    request = &DescribeBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeBlackWhiteIpList")
    return
}

func NewDescribeBlackWhiteIpListResponse() (response *DescribeBlackWhiteIpListResponse) {
    response = &DescribeBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlackWhiteIpList
// 获取DDoS防护的IP黑白名单
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBlackWhiteIpList(request *DescribeBlackWhiteIpListRequest) (response *DescribeBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeBlackWhiteIpListRequest()
    }
    response = NewDescribeBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultAlarmThresholdRequest() (request *DescribeDefaultAlarmThresholdRequest) {
    request = &DescribeDefaultAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDefaultAlarmThreshold")
    return
}

func NewDescribeDefaultAlarmThresholdResponse() (response *DescribeDefaultAlarmThresholdResponse) {
    response = &DescribeDefaultAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultAlarmThreshold
// 获取单IP默认告警阈值配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDefaultAlarmThreshold(request *DescribeDefaultAlarmThresholdRequest) (response *DescribeDefaultAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultAlarmThresholdRequest()
    }
    response = NewDescribeDefaultAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7RulesBySSLCertIdRequest() (request *DescribeL7RulesBySSLCertIdRequest) {
    request = &DescribeL7RulesBySSLCertIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeL7RulesBySSLCertId")
    return
}

func NewDescribeL7RulesBySSLCertIdResponse() (response *DescribeL7RulesBySSLCertIdResponse) {
    response = &DescribeL7RulesBySSLCertIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeL7RulesBySSLCertId
// 查询与证书ID对于域名匹配的七层规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeL7RulesBySSLCertId(request *DescribeL7RulesBySSLCertIdRequest) (response *DescribeL7RulesBySSLCertIdResponse, err error) {
    if request == nil {
        request = NewDescribeL7RulesBySSLCertIdRequest()
    }
    response = NewDescribeL7RulesBySSLCertIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBGPIPInstancesRequest() (request *DescribeListBGPIPInstancesRequest) {
    request = &DescribeListBGPIPInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBGPIPInstances")
    return
}

func NewDescribeListBGPIPInstancesResponse() (response *DescribeListBGPIPInstancesResponse) {
    response = &DescribeListBGPIPInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBGPIPInstances
// 获取高防IP资产实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPIPInstances(request *DescribeListBGPIPInstancesRequest) (response *DescribeListBGPIPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPIPInstancesRequest()
    }
    response = NewDescribeListBGPIPInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBGPInstancesRequest() (request *DescribeListBGPInstancesRequest) {
    request = &DescribeListBGPInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBGPInstances")
    return
}

func NewDescribeListBGPInstancesResponse() (response *DescribeListBGPInstancesResponse) {
    response = &DescribeListBGPInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBGPInstances
// 获取高防包资产实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBGPInstances(request *DescribeListBGPInstancesRequest) (response *DescribeListBGPInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeListBGPInstancesRequest()
    }
    response = NewDescribeListBGPInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListBlackWhiteIpListRequest() (request *DescribeListBlackWhiteIpListRequest) {
    request = &DescribeListBlackWhiteIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListBlackWhiteIpList")
    return
}

func NewDescribeListBlackWhiteIpListResponse() (response *DescribeListBlackWhiteIpListResponse) {
    response = &DescribeListBlackWhiteIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListBlackWhiteIpList
// 获取DDoS防护的IP黑白名单列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListBlackWhiteIpList(request *DescribeListBlackWhiteIpListRequest) (response *DescribeListBlackWhiteIpListResponse, err error) {
    if request == nil {
        request = NewDescribeListBlackWhiteIpListRequest()
    }
    response = NewDescribeListBlackWhiteIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSAIRequest() (request *DescribeListDDoSAIRequest) {
    request = &DescribeListDDoSAIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSAI")
    return
}

func NewDescribeListDDoSAIResponse() (response *DescribeListDDoSAIResponse) {
    response = &DescribeListDDoSAIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSAI
// 获取DDoS防护的AI防护开关列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSAI(request *DescribeListDDoSAIRequest) (response *DescribeListDDoSAIResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSAIRequest()
    }
    response = NewDescribeListDDoSAIResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSGeoIPBlockConfigRequest() (request *DescribeListDDoSGeoIPBlockConfigRequest) {
    request = &DescribeListDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSGeoIPBlockConfig")
    return
}

func NewDescribeListDDoSGeoIPBlockConfigResponse() (response *DescribeListDDoSGeoIPBlockConfigResponse) {
    response = &DescribeListDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSGeoIPBlockConfig
// 获取DDoS防护的区域封禁配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSGeoIPBlockConfig(request *DescribeListDDoSGeoIPBlockConfigRequest) (response *DescribeListDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSGeoIPBlockConfigRequest()
    }
    response = NewDescribeListDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListDDoSSpeedLimitConfigRequest() (request *DescribeListDDoSSpeedLimitConfigRequest) {
    request = &DescribeListDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListDDoSSpeedLimitConfig")
    return
}

func NewDescribeListDDoSSpeedLimitConfigResponse() (response *DescribeListDDoSSpeedLimitConfigResponse) {
    response = &DescribeListDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListDDoSSpeedLimitConfig
// 获取DDoS防护的访问限速配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListDDoSSpeedLimitConfig(request *DescribeListDDoSSpeedLimitConfigRequest) (response *DescribeListDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListDDoSSpeedLimitConfigRequest()
    }
    response = NewDescribeListDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListIPAlarmConfigRequest() (request *DescribeListIPAlarmConfigRequest) {
    request = &DescribeListIPAlarmConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListIPAlarmConfig")
    return
}

func NewDescribeListIPAlarmConfigResponse() (response *DescribeListIPAlarmConfigResponse) {
    response = &DescribeListIPAlarmConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListIPAlarmConfig
// 获取单IP告警阈值配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListIPAlarmConfig(request *DescribeListIPAlarmConfigRequest) (response *DescribeListIPAlarmConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListIPAlarmConfigRequest()
    }
    response = NewDescribeListIPAlarmConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListListenerRequest() (request *DescribeListListenerRequest) {
    request = &DescribeListListenerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListListener")
    return
}

func NewDescribeListListenerResponse() (response *DescribeListListenerResponse) {
    response = &DescribeListListenerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListListener
// 获取转发监听器列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListListener(request *DescribeListListenerRequest) (response *DescribeListListenerResponse, err error) {
    if request == nil {
        request = NewDescribeListListenerRequest()
    }
    response = NewDescribeListListenerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListPacketFilterConfigRequest() (request *DescribeListPacketFilterConfigRequest) {
    request = &DescribeListPacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListPacketFilterConfig")
    return
}

func NewDescribeListPacketFilterConfigResponse() (response *DescribeListPacketFilterConfigResponse) {
    response = &DescribeListPacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListPacketFilterConfig
// 获取DDoS防护的特征过滤规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListPacketFilterConfig(request *DescribeListPacketFilterConfigRequest) (response *DescribeListPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListPacketFilterConfigRequest()
    }
    response = NewDescribeListPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListProtectThresholdConfigRequest() (request *DescribeListProtectThresholdConfigRequest) {
    request = &DescribeListProtectThresholdConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListProtectThresholdConfig")
    return
}

func NewDescribeListProtectThresholdConfigResponse() (response *DescribeListProtectThresholdConfigResponse) {
    response = &DescribeListProtectThresholdConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListProtectThresholdConfig
// 获取防护阈值配置列表，包括DDoS的AI、等级、CC阈值开关等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtectThresholdConfig(request *DescribeListProtectThresholdConfigRequest) (response *DescribeListProtectThresholdConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtectThresholdConfigRequest()
    }
    response = NewDescribeListProtectThresholdConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListProtocolBlockConfigRequest() (request *DescribeListProtocolBlockConfigRequest) {
    request = &DescribeListProtocolBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListProtocolBlockConfig")
    return
}

func NewDescribeListProtocolBlockConfigResponse() (response *DescribeListProtocolBlockConfigResponse) {
    response = &DescribeListProtocolBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListProtocolBlockConfig
// 获取DDoS防护的协议封禁配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListProtocolBlockConfig(request *DescribeListProtocolBlockConfigRequest) (response *DescribeListProtocolBlockConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListProtocolBlockConfigRequest()
    }
    response = NewDescribeListProtocolBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListSchedulingDomainRequest() (request *DescribeListSchedulingDomainRequest) {
    request = &DescribeListSchedulingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListSchedulingDomain")
    return
}

func NewDescribeListSchedulingDomainResponse() (response *DescribeListSchedulingDomainResponse) {
    response = &DescribeListSchedulingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListSchedulingDomain
// 获取智能调度域名列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListSchedulingDomain(request *DescribeListSchedulingDomainRequest) (response *DescribeListSchedulingDomainResponse, err error) {
    if request == nil {
        request = NewDescribeListSchedulingDomainRequest()
    }
    response = NewDescribeListSchedulingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeListWaterPrintConfigRequest() (request *DescribeListWaterPrintConfigRequest) {
    request = &DescribeListWaterPrintConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeListWaterPrintConfig")
    return
}

func NewDescribeListWaterPrintConfigResponse() (response *DescribeListWaterPrintConfigResponse) {
    response = &DescribeListWaterPrintConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeListWaterPrintConfig
// 获取DDoS防护的水印防护配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeListWaterPrintConfig(request *DescribeListWaterPrintConfigRequest) (response *DescribeListWaterPrintConfigResponse, err error) {
    if request == nil {
        request = NewDescribeListWaterPrintConfigRequest()
    }
    response = NewDescribeListWaterPrintConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateDDoSEipAddressRequest() (request *DisassociateDDoSEipAddressRequest) {
    request = &DisassociateDDoSEipAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "DisassociateDDoSEipAddress")
    return
}

func NewDisassociateDDoSEipAddressResponse() (response *DisassociateDDoSEipAddressResponse) {
    response = &DisassociateDDoSEipAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateDDoSEipAddress
// 本接口 (DisassociateDDoSEipAddress) 用于解绑高防弹性公网IP。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisassociateDDoSEipAddress(request *DisassociateDDoSEipAddressRequest) (response *DisassociateDDoSEipAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateDDoSEipAddressRequest()
    }
    response = NewDisassociateDDoSEipAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSGeoIPBlockConfigRequest() (request *ModifyDDoSGeoIPBlockConfigRequest) {
    request = &ModifyDDoSGeoIPBlockConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSGeoIPBlockConfig")
    return
}

func NewModifyDDoSGeoIPBlockConfigResponse() (response *ModifyDDoSGeoIPBlockConfigResponse) {
    response = &ModifyDDoSGeoIPBlockConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSGeoIPBlockConfig
// 修改DDoS防护的区域封禁配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSGeoIPBlockConfig(request *ModifyDDoSGeoIPBlockConfigRequest) (response *ModifyDDoSGeoIPBlockConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSGeoIPBlockConfigRequest()
    }
    response = NewModifyDDoSGeoIPBlockConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSSpeedLimitConfigRequest() (request *ModifyDDoSSpeedLimitConfigRequest) {
    request = &ModifyDDoSSpeedLimitConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDDoSSpeedLimitConfig")
    return
}

func NewModifyDDoSSpeedLimitConfigResponse() (response *ModifyDDoSSpeedLimitConfigResponse) {
    response = &ModifyDDoSSpeedLimitConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDDoSSpeedLimitConfig
// 修改DDoS防护的访问限速配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDDoSSpeedLimitConfig(request *ModifyDDoSSpeedLimitConfigRequest) (response *ModifyDDoSSpeedLimitConfigResponse, err error) {
    if request == nil {
        request = NewModifyDDoSSpeedLimitConfigRequest()
    }
    response = NewModifyDDoSSpeedLimitConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainUsrNameRequest() (request *ModifyDomainUsrNameRequest) {
    request = &ModifyDomainUsrNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyDomainUsrName")
    return
}

func NewModifyDomainUsrNameResponse() (response *ModifyDomainUsrNameResponse) {
    response = &ModifyDomainUsrNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDomainUsrName
// 修改智能解析域名名称
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDomainUsrName(request *ModifyDomainUsrNameRequest) (response *ModifyDomainUsrNameResponse, err error) {
    if request == nil {
        request = NewModifyDomainUsrNameRequest()
    }
    response = NewModifyDomainUsrNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPacketFilterConfigRequest() (request *ModifyPacketFilterConfigRequest) {
    request = &ModifyPacketFilterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("antiddos", APIVersion, "ModifyPacketFilterConfig")
    return
}

func NewModifyPacketFilterConfigResponse() (response *ModifyPacketFilterConfigResponse) {
    response = &ModifyPacketFilterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPacketFilterConfig
// 修改DDoS防护的特征过滤规则
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPacketFilterConfig(request *ModifyPacketFilterConfigRequest) (response *ModifyPacketFilterConfigResponse, err error) {
    if request == nil {
        request = NewModifyPacketFilterConfigRequest()
    }
    response = NewModifyPacketFilterConfigResponse()
    err = c.Send(request, response)
    return
}
