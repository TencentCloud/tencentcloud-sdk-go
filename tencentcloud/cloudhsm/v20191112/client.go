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

package v20191112

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-12"

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


func NewDescribeHSMBySubnetIdRequest() (request *DescribeHSMBySubnetIdRequest) {
    request = &DescribeHSMBySubnetIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeHSMBySubnetId")
    
    
    return
}

func NewDescribeHSMBySubnetIdResponse() (response *DescribeHSMBySubnetIdResponse) {
    response = &DescribeHSMBySubnetIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHSMBySubnetId
// 通过SubnetId获取Hsm资源数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHSMBySubnetId(request *DescribeHSMBySubnetIdRequest) (response *DescribeHSMBySubnetIdResponse, err error) {
    return c.DescribeHSMBySubnetIdWithContext(context.Background(), request)
}

// DescribeHSMBySubnetId
// 通过SubnetId获取Hsm资源数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHSMBySubnetIdWithContext(ctx context.Context, request *DescribeHSMBySubnetIdRequest) (response *DescribeHSMBySubnetIdResponse, err error) {
    if request == nil {
        request = NewDescribeHSMBySubnetIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHSMBySubnetId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHSMBySubnetIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHSMByVpcIdRequest() (request *DescribeHSMByVpcIdRequest) {
    request = &DescribeHSMByVpcIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeHSMByVpcId")
    
    
    return
}

func NewDescribeHSMByVpcIdResponse() (response *DescribeHSMByVpcIdResponse) {
    response = &DescribeHSMByVpcIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHSMByVpcId
// 通过VpcId获取Hsm资源数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHSMByVpcId(request *DescribeHSMByVpcIdRequest) (response *DescribeHSMByVpcIdResponse, err error) {
    return c.DescribeHSMByVpcIdWithContext(context.Background(), request)
}

// DescribeHSMByVpcId
// 通过VpcId获取Hsm资源数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHSMByVpcIdWithContext(ctx context.Context, request *DescribeHSMByVpcIdRequest) (response *DescribeHSMByVpcIdResponse, err error) {
    if request == nil {
        request = NewDescribeHSMByVpcIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHSMByVpcId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHSMByVpcIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetRequest() (request *DescribeSubnetRequest) {
    request = &DescribeSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeSubnet")
    
    
    return
}

func NewDescribeSubnetResponse() (response *DescribeSubnetResponse) {
    response = &DescribeSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubnet
// 查询子网列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSubnet(request *DescribeSubnetRequest) (response *DescribeSubnetResponse, err error) {
    return c.DescribeSubnetWithContext(context.Background(), request)
}

// DescribeSubnet
// 查询子网列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeSubnetWithContext(ctx context.Context, request *DescribeSubnetRequest) (response *DescribeSubnetResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubnet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportedHsmRequest() (request *DescribeSupportedHsmRequest) {
    request = &DescribeSupportedHsmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeSupportedHsm")
    
    
    return
}

func NewDescribeSupportedHsmResponse() (response *DescribeSupportedHsmResponse) {
    response = &DescribeSupportedHsmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSupportedHsm
// 获取当前地域所支持的设备列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSupportedHsm(request *DescribeSupportedHsmRequest) (response *DescribeSupportedHsmResponse, err error) {
    return c.DescribeSupportedHsmWithContext(context.Background(), request)
}

// DescribeSupportedHsm
// 获取当前地域所支持的设备列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeSupportedHsmWithContext(ctx context.Context, request *DescribeSupportedHsmRequest) (response *DescribeSupportedHsmResponse, err error) {
    if request == nil {
        request = NewDescribeSupportedHsmRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportedHsm require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportedHsmResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsgRequest() (request *DescribeUsgRequest) {
    request = &DescribeUsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeUsg")
    
    
    return
}

func NewDescribeUsgResponse() (response *DescribeUsgResponse) {
    response = &DescribeUsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsg
// 根据用户的AppId获取用户安全组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUsg(request *DescribeUsgRequest) (response *DescribeUsgResponse, err error) {
    return c.DescribeUsgWithContext(context.Background(), request)
}

// DescribeUsg
// 根据用户的AppId获取用户安全组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUsgWithContext(ctx context.Context, request *DescribeUsgRequest) (response *DescribeUsgResponse, err error) {
    if request == nil {
        request = NewDescribeUsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsgResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsgRuleRequest() (request *DescribeUsgRuleRequest) {
    request = &DescribeUsgRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeUsgRule")
    
    
    return
}

func NewDescribeUsgRuleResponse() (response *DescribeUsgRuleResponse) {
    response = &DescribeUsgRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsgRule
// 获取安全组详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUsgRule(request *DescribeUsgRuleRequest) (response *DescribeUsgRuleResponse, err error) {
    return c.DescribeUsgRuleWithContext(context.Background(), request)
}

// DescribeUsgRule
// 获取安全组详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUsgRuleWithContext(ctx context.Context, request *DescribeUsgRuleRequest) (response *DescribeUsgRuleResponse, err error) {
    if request == nil {
        request = NewDescribeUsgRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsgRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsgRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcRequest() (request *DescribeVpcRequest) {
    request = &DescribeVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeVpc")
    
    
    return
}

func NewDescribeVpcResponse() (response *DescribeVpcResponse) {
    response = &DescribeVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVpc
// 查询用户的私有网络列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpc(request *DescribeVpcRequest) (response *DescribeVpcResponse, err error) {
    return c.DescribeVpcWithContext(context.Background(), request)
}

// DescribeVpc
// 查询用户的私有网络列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVpcWithContext(ctx context.Context, request *DescribeVpcRequest) (response *DescribeVpcResponse, err error) {
    if request == nil {
        request = NewDescribeVpcRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVsmAttributesRequest() (request *DescribeVsmAttributesRequest) {
    request = &DescribeVsmAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeVsmAttributes")
    
    
    return
}

func NewDescribeVsmAttributesResponse() (response *DescribeVsmAttributesResponse) {
    response = &DescribeVsmAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVsmAttributes
// 获取VSM属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVsmAttributes(request *DescribeVsmAttributesRequest) (response *DescribeVsmAttributesResponse, err error) {
    return c.DescribeVsmAttributesWithContext(context.Background(), request)
}

// DescribeVsmAttributes
// 获取VSM属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVsmAttributesWithContext(ctx context.Context, request *DescribeVsmAttributesRequest) (response *DescribeVsmAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeVsmAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVsmAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVsmAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVsmsRequest() (request *DescribeVsmsRequest) {
    request = &DescribeVsmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "DescribeVsms")
    
    
    return
}

func NewDescribeVsmsResponse() (response *DescribeVsmsResponse) {
    response = &DescribeVsmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVsms
// 获取用户VSM列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVsms(request *DescribeVsmsRequest) (response *DescribeVsmsResponse, err error) {
    return c.DescribeVsmsWithContext(context.Background(), request)
}

// DescribeVsms
// 获取用户VSM列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeVsmsWithContext(ctx context.Context, request *DescribeVsmsRequest) (response *DescribeVsmsResponse, err error) {
    if request == nil {
        request = NewDescribeVsmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVsms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVsmsResponse()
    err = c.Send(request, response)
    return
}

func NewGetAlarmEventRequest() (request *GetAlarmEventRequest) {
    request = &GetAlarmEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "GetAlarmEvent")
    
    
    return
}

func NewGetAlarmEventResponse() (response *GetAlarmEventResponse) {
    response = &GetAlarmEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAlarmEvent
// 获取告警事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetAlarmEvent(request *GetAlarmEventRequest) (response *GetAlarmEventResponse, err error) {
    return c.GetAlarmEventWithContext(context.Background(), request)
}

// GetAlarmEvent
// 获取告警事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetAlarmEventWithContext(ctx context.Context, request *GetAlarmEventRequest) (response *GetAlarmEventResponse, err error) {
    if request == nil {
        request = NewGetAlarmEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAlarmEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAlarmEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetVsmMonitorInfoRequest() (request *GetVsmMonitorInfoRequest) {
    request = &GetVsmMonitorInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "GetVsmMonitorInfo")
    
    
    return
}

func NewGetVsmMonitorInfoResponse() (response *GetVsmMonitorInfoResponse) {
    response = &GetVsmMonitorInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetVsmMonitorInfo
// 获取VSM监控信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetVsmMonitorInfo(request *GetVsmMonitorInfoRequest) (response *GetVsmMonitorInfoResponse, err error) {
    return c.GetVsmMonitorInfoWithContext(context.Background(), request)
}

// GetVsmMonitorInfo
// 获取VSM监控信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetVsmMonitorInfoWithContext(ctx context.Context, request *GetVsmMonitorInfoRequest) (response *GetVsmMonitorInfoResponse, err error) {
    if request == nil {
        request = NewGetVsmMonitorInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetVsmMonitorInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetVsmMonitorInfoResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceBuyVsmRequest() (request *InquiryPriceBuyVsmRequest) {
    request = &InquiryPriceBuyVsmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "InquiryPriceBuyVsm")
    
    
    return
}

func NewInquiryPriceBuyVsmResponse() (response *InquiryPriceBuyVsmResponse) {
    response = &InquiryPriceBuyVsmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceBuyVsm
// 购买询价接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceBuyVsm(request *InquiryPriceBuyVsmRequest) (response *InquiryPriceBuyVsmResponse, err error) {
    return c.InquiryPriceBuyVsmWithContext(context.Background(), request)
}

// InquiryPriceBuyVsm
// 购买询价接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) InquiryPriceBuyVsmWithContext(ctx context.Context, request *InquiryPriceBuyVsmRequest) (response *InquiryPriceBuyVsmResponse, err error) {
    if request == nil {
        request = NewInquiryPriceBuyVsmRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceBuyVsm require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceBuyVsmResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmEventRequest() (request *ModifyAlarmEventRequest) {
    request = &ModifyAlarmEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "ModifyAlarmEvent")
    
    
    return
}

func NewModifyAlarmEventResponse() (response *ModifyAlarmEventResponse) {
    response = &ModifyAlarmEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAlarmEvent
// 修改告警事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAlarmEvent(request *ModifyAlarmEventRequest) (response *ModifyAlarmEventResponse, err error) {
    return c.ModifyAlarmEventWithContext(context.Background(), request)
}

// ModifyAlarmEvent
// 修改告警事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAlarmEventWithContext(ctx context.Context, request *ModifyAlarmEventRequest) (response *ModifyAlarmEventResponse, err error) {
    if request == nil {
        request = NewModifyAlarmEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmEventResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVsmAttributesRequest() (request *ModifyVsmAttributesRequest) {
    request = &ModifyVsmAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudhsm", APIVersion, "ModifyVsmAttributes")
    
    
    return
}

func NewModifyVsmAttributesResponse() (response *ModifyVsmAttributesResponse) {
    response = &ModifyVsmAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyVsmAttributes
// 修改VSM属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyVsmAttributes(request *ModifyVsmAttributesRequest) (response *ModifyVsmAttributesResponse, err error) {
    return c.ModifyVsmAttributesWithContext(context.Background(), request)
}

// ModifyVsmAttributes
// 修改VSM属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyVsmAttributesWithContext(ctx context.Context, request *ModifyVsmAttributesRequest) (response *ModifyVsmAttributesResponse, err error) {
    if request == nil {
        request = NewModifyVsmAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVsmAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVsmAttributesResponse()
    err = c.Send(request, response)
    return
}
