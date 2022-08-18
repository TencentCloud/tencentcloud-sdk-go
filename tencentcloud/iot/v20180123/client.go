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

package v20180123

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-01-23"

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


func NewActivateRuleRequest() (request *ActivateRuleRequest) {
    request = &ActivateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "ActivateRule")
    
    
    return
}

func NewActivateRuleResponse() (response *ActivateRuleResponse) {
    response = &ActivateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ActivateRule
// 启用规则
//
// 可能返回的错误码:
//  INTERNALERROR_IOTACTIONSYSTEMERROR = "InternalError.IotActionSystemError"
//  INTERNALERROR_MQRULESYSTEMERROR = "InternalError.MqruleSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_IOTRULEISACTIVE = "ResourceUnavailable.IotRuleIsActive"
//  RESOURCEUNAVAILABLE_IOTRULENOACTION = "ResourceUnavailable.IotRuleNoAction"
//  RESOURCEUNAVAILABLE_IOTRULENOQUERY = "ResourceUnavailable.IotRuleNoQuery"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) ActivateRule(request *ActivateRuleRequest) (response *ActivateRuleResponse, err error) {
    return c.ActivateRuleWithContext(context.Background(), request)
}

// ActivateRule
// 启用规则
//
// 可能返回的错误码:
//  INTERNALERROR_IOTACTIONSYSTEMERROR = "InternalError.IotActionSystemError"
//  INTERNALERROR_MQRULESYSTEMERROR = "InternalError.MqruleSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_IOTRULEISACTIVE = "ResourceUnavailable.IotRuleIsActive"
//  RESOURCEUNAVAILABLE_IOTRULENOACTION = "ResourceUnavailable.IotRuleNoAction"
//  RESOURCEUNAVAILABLE_IOTRULENOQUERY = "ResourceUnavailable.IotRuleNoQuery"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) ActivateRuleWithContext(ctx context.Context, request *ActivateRuleRequest) (response *ActivateRuleResponse, err error) {
    if request == nil {
        request = NewActivateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddDeviceRequest() (request *AddDeviceRequest) {
    request = &AddDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AddDevice")
    
    
    return
}

func NewAddDeviceResponse() (response *AddDeviceResponse) {
    response = &AddDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddDevice
// 提供在指定的产品Id下创建一个设备的能力，生成设备名称与设备秘钥。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTDBERROR = "InternalError.IotDbError"
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTDEVICEOPTOOOFTEN = "LimitExceeded.IotDeviceOpTooOften"
//  RESOURCEINUSE_IOTDEVICEEXISTS = "ResourceInUse.IotDeviceExists"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_MQIOTRESOURCEEXISTS = "ResourceInUse.MqiotResourceExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) AddDevice(request *AddDeviceRequest) (response *AddDeviceResponse, err error) {
    return c.AddDeviceWithContext(context.Background(), request)
}

// AddDevice
// 提供在指定的产品Id下创建一个设备的能力，生成设备名称与设备秘钥。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTDBERROR = "InternalError.IotDbError"
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTDEVICEOPTOOOFTEN = "LimitExceeded.IotDeviceOpTooOften"
//  RESOURCEINUSE_IOTDEVICEEXISTS = "ResourceInUse.IotDeviceExists"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_MQIOTRESOURCEEXISTS = "ResourceInUse.MqiotResourceExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
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

func NewAddProductRequest() (request *AddProductRequest) {
    request = &AddProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AddProduct")
    
    
    return
}

func NewAddProductResponse() (response *AddProductResponse) {
    response = &AddProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddProduct
// 本接口(AddProduct)用于创建、定义某款硬件产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATE = "InvalidParameter.IotProductInvalidDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATERANGE = "InvalidParameter.IotProductInvalidDataTemplateRange"
//  LIMITEXCEEDED_IOTPRODUCTOPTOOOFTEN = "LimitExceeded.IotProductOpTooOften"
//  LIMITEXCEEDED_IOTUSERTOOMANYPRODUCTS = "LimitExceeded.IotUserTooManyProducts"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTPRODUCTEXISTS = "ResourceInUse.IotProductExists"
//  RESOURCEINUSE_MQIOTRESOURCEEXISTS = "ResourceInUse.MqiotResourceExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) AddProduct(request *AddProductRequest) (response *AddProductResponse, err error) {
    return c.AddProductWithContext(context.Background(), request)
}

// AddProduct
// 本接口(AddProduct)用于创建、定义某款硬件产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATE = "InvalidParameter.IotProductInvalidDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATERANGE = "InvalidParameter.IotProductInvalidDataTemplateRange"
//  LIMITEXCEEDED_IOTPRODUCTOPTOOOFTEN = "LimitExceeded.IotProductOpTooOften"
//  LIMITEXCEEDED_IOTUSERTOOMANYPRODUCTS = "LimitExceeded.IotUserTooManyProducts"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTPRODUCTEXISTS = "ResourceInUse.IotProductExists"
//  RESOURCEINUSE_MQIOTRESOURCEEXISTS = "ResourceInUse.MqiotResourceExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) AddProductWithContext(ctx context.Context, request *AddProductRequest) (response *AddProductResponse, err error) {
    if request == nil {
        request = NewAddProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddProductResponse()
    err = c.Send(request, response)
    return
}

func NewAddRuleRequest() (request *AddRuleRequest) {
    request = &AddRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AddRule")
    
    
    return
}

func NewAddRuleResponse() (response *AddRuleResponse) {
    response = &AddRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddRule
// 新增规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTEMPTYDATATEMPLATE = "InvalidParameter.IotProductEmptyDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATAPROTOCOL = "InvalidParameter.IotProductInvalidDataProtocol"
//  LIMITEXCEEDED_IOTRULEOPTOOMANY = "LimitExceeded.IotRuleOpTooMany"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTRULEEXISTS = "ResourceInUse.IotRuleExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) AddRule(request *AddRuleRequest) (response *AddRuleResponse, err error) {
    return c.AddRuleWithContext(context.Background(), request)
}

// AddRule
// 新增规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTEMPTYDATATEMPLATE = "InvalidParameter.IotProductEmptyDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATAPROTOCOL = "InvalidParameter.IotProductInvalidDataProtocol"
//  LIMITEXCEEDED_IOTRULEOPTOOMANY = "LimitExceeded.IotRuleOpTooMany"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTRULEEXISTS = "ResourceInUse.IotRuleExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) AddRuleWithContext(ctx context.Context, request *AddRuleRequest) (response *AddRuleResponse, err error) {
    if request == nil {
        request = NewAddRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddRuleResponse()
    err = c.Send(request, response)
    return
}

func NewAddTopicRequest() (request *AddTopicRequest) {
    request = &AddTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AddTopic")
    
    
    return
}

func NewAddTopicResponse() (response *AddTopicResponse) {
    response = &AddTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddTopic
// 新增Topic，用于设备或应用发布消息至该Topic或订阅该Topic的消息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTPRODUCTTOOMANYTOPICS = "LimitExceeded.IotProductTooManyTopics"
//  LIMITEXCEEDED_IOTTOPICOPTOOOFTEN = "LimitExceeded.IotTopicOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTTOPICEXISTS = "ResourceInUse.IotTopicExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) AddTopic(request *AddTopicRequest) (response *AddTopicResponse, err error) {
    return c.AddTopicWithContext(context.Background(), request)
}

// AddTopic
// 新增Topic，用于设备或应用发布消息至该Topic或订阅该Topic的消息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTPRODUCTTOOMANYTOPICS = "LimitExceeded.IotProductTooManyTopics"
//  LIMITEXCEEDED_IOTTOPICOPTOOOFTEN = "LimitExceeded.IotTopicOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTTOPICEXISTS = "ResourceInUse.IotTopicExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) AddTopicWithContext(ctx context.Context, request *AddTopicRequest) (response *AddTopicResponse, err error) {
    if request == nil {
        request = NewAddTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTopicResponse()
    err = c.Send(request, response)
    return
}

func NewAppAddUserRequest() (request *AppAddUserRequest) {
    request = &AppAddUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppAddUser")
    
    
    return
}

func NewAppAddUserResponse() (response *AppAddUserResponse) {
    response = &AppAddUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppAddUser
// 为APP提供用户注册功能
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCEINUSE_IOTAPPLICATIONUSEREXISTS = "ResourceInUse.IotApplicationUserExists"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
func (c *Client) AppAddUser(request *AppAddUserRequest) (response *AppAddUserResponse, err error) {
    return c.AppAddUserWithContext(context.Background(), request)
}

// AppAddUser
// 为APP提供用户注册功能
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCEINUSE_IOTAPPLICATIONUSEREXISTS = "ResourceInUse.IotApplicationUserExists"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
func (c *Client) AppAddUserWithContext(ctx context.Context, request *AppAddUserRequest) (response *AppAddUserResponse, err error) {
    if request == nil {
        request = NewAppAddUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppAddUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppAddUserResponse()
    err = c.Send(request, response)
    return
}

func NewAppDeleteDeviceRequest() (request *AppDeleteDeviceRequest) {
    request = &AppDeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppDeleteDevice")
    
    
    return
}

func NewAppDeleteDeviceResponse() (response *AppDeleteDeviceResponse) {
    response = &AppDeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppDeleteDevice
// 用户解除与设备的关联关系，解除后APP用户无法控制设备，获取设备数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppDeleteDevice(request *AppDeleteDeviceRequest) (response *AppDeleteDeviceResponse, err error) {
    return c.AppDeleteDeviceWithContext(context.Background(), request)
}

// AppDeleteDevice
// 用户解除与设备的关联关系，解除后APP用户无法控制设备，获取设备数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppDeleteDeviceWithContext(ctx context.Context, request *AppDeleteDeviceRequest) (response *AppDeleteDeviceResponse, err error) {
    if request == nil {
        request = NewAppDeleteDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppDeleteDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewAppGetDeviceRequest() (request *AppGetDeviceRequest) {
    request = &AppGetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppGetDevice")
    
    
    return
}

func NewAppGetDeviceResponse() (response *AppGetDeviceResponse) {
    response = &AppGetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppGetDevice
// 获取绑定设备的基本信息与数据模板定义
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDevice(request *AppGetDeviceRequest) (response *AppGetDeviceResponse, err error) {
    return c.AppGetDeviceWithContext(context.Background(), request)
}

// AppGetDevice
// 获取绑定设备的基本信息与数据模板定义
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDeviceWithContext(ctx context.Context, request *AppGetDeviceRequest) (response *AppGetDeviceResponse, err error) {
    if request == nil {
        request = NewAppGetDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppGetDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppGetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewAppGetDeviceDataRequest() (request *AppGetDeviceDataRequest) {
    request = &AppGetDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppGetDeviceData")
    
    
    return
}

func NewAppGetDeviceDataResponse() (response *AppGetDeviceDataResponse) {
    response = &AppGetDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppGetDeviceData
// 获取绑定设备数据，用于实时展示设备的最新数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDeviceData(request *AppGetDeviceDataRequest) (response *AppGetDeviceDataResponse, err error) {
    return c.AppGetDeviceDataWithContext(context.Background(), request)
}

// AppGetDeviceData
// 获取绑定设备数据，用于实时展示设备的最新数据
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDeviceDataWithContext(ctx context.Context, request *AppGetDeviceDataRequest) (response *AppGetDeviceDataResponse, err error) {
    if request == nil {
        request = NewAppGetDeviceDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppGetDeviceData require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppGetDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewAppGetDeviceStatusesRequest() (request *AppGetDeviceStatusesRequest) {
    request = &AppGetDeviceStatusesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppGetDeviceStatuses")
    
    
    return
}

func NewAppGetDeviceStatusesResponse() (response *AppGetDeviceStatusesResponse) {
    response = &AppGetDeviceStatusesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppGetDeviceStatuses
// 获取绑定设备的上下线状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDeviceStatuses(request *AppGetDeviceStatusesRequest) (response *AppGetDeviceStatusesResponse, err error) {
    return c.AppGetDeviceStatusesWithContext(context.Background(), request)
}

// AppGetDeviceStatuses
// 获取绑定设备的上下线状态
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDeviceStatusesWithContext(ctx context.Context, request *AppGetDeviceStatusesRequest) (response *AppGetDeviceStatusesResponse, err error) {
    if request == nil {
        request = NewAppGetDeviceStatusesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppGetDeviceStatuses require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppGetDeviceStatusesResponse()
    err = c.Send(request, response)
    return
}

func NewAppGetDevicesRequest() (request *AppGetDevicesRequest) {
    request = &AppGetDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppGetDevices")
    
    
    return
}

func NewAppGetDevicesResponse() (response *AppGetDevicesResponse) {
    response = &AppGetDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppGetDevices
// 获取用户的绑定设备列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDevices(request *AppGetDevicesRequest) (response *AppGetDevicesResponse, err error) {
    return c.AppGetDevicesWithContext(context.Background(), request)
}

// AppGetDevices
// 获取用户的绑定设备列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetDevicesWithContext(ctx context.Context, request *AppGetDevicesRequest) (response *AppGetDevicesResponse, err error) {
    if request == nil {
        request = NewAppGetDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppGetDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppGetDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewAppGetTokenRequest() (request *AppGetTokenRequest) {
    request = &AppGetTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppGetToken")
    
    
    return
}

func NewAppGetTokenResponse() (response *AppGetTokenResponse) {
    response = &AppGetTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppGetToken
// 获取用户token
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTAPPLICATIONINVALIDUSERPASSWORD = "InvalidParameter.IotApplicationInvalidUserPassword"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetToken(request *AppGetTokenRequest) (response *AppGetTokenResponse, err error) {
    return c.AppGetTokenWithContext(context.Background(), request)
}

// AppGetToken
// 获取用户token
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTAPPLICATIONINVALIDUSERPASSWORD = "InvalidParameter.IotApplicationInvalidUserPassword"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetTokenWithContext(ctx context.Context, request *AppGetTokenRequest) (response *AppGetTokenResponse, err error) {
    if request == nil {
        request = NewAppGetTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppGetToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppGetTokenResponse()
    err = c.Send(request, response)
    return
}

func NewAppGetUserRequest() (request *AppGetUserRequest) {
    request = &AppGetUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppGetUser")
    
    
    return
}

func NewAppGetUserResponse() (response *AppGetUserResponse) {
    response = &AppGetUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppGetUser
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetUser(request *AppGetUserRequest) (response *AppGetUserResponse, err error) {
    return c.AppGetUserWithContext(context.Background(), request)
}

// AppGetUser
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppGetUserWithContext(ctx context.Context, request *AppGetUserRequest) (response *AppGetUserResponse, err error) {
    if request == nil {
        request = NewAppGetUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppGetUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppGetUserResponse()
    err = c.Send(request, response)
    return
}

func NewAppIssueDeviceControlRequest() (request *AppIssueDeviceControlRequest) {
    request = &AppIssueDeviceControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppIssueDeviceControl")
    
    
    return
}

func NewAppIssueDeviceControlResponse() (response *AppIssueDeviceControlResponse) {
    response = &AppIssueDeviceControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppIssueDeviceControl
// 用户通过APP控制设备
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
func (c *Client) AppIssueDeviceControl(request *AppIssueDeviceControlRequest) (response *AppIssueDeviceControlResponse, err error) {
    return c.AppIssueDeviceControlWithContext(context.Background(), request)
}

// AppIssueDeviceControl
// 用户通过APP控制设备
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
func (c *Client) AppIssueDeviceControlWithContext(ctx context.Context, request *AppIssueDeviceControlRequest) (response *AppIssueDeviceControlResponse, err error) {
    if request == nil {
        request = NewAppIssueDeviceControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppIssueDeviceControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppIssueDeviceControlResponse()
    err = c.Send(request, response)
    return
}

func NewAppResetPasswordRequest() (request *AppResetPasswordRequest) {
    request = &AppResetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppResetPassword")
    
    
    return
}

func NewAppResetPasswordResponse() (response *AppResetPasswordResponse) {
    response = &AppResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppResetPassword
// 重置APP用户密码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTAPPLICATIONINVALIDPASSWORD = "InvalidParameter.IotApplicationInvalidPassword"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppResetPassword(request *AppResetPasswordRequest) (response *AppResetPasswordResponse, err error) {
    return c.AppResetPasswordWithContext(context.Background(), request)
}

// AppResetPassword
// 重置APP用户密码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTAPPLICATIONINVALIDPASSWORD = "InvalidParameter.IotApplicationInvalidPassword"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppResetPasswordWithContext(ctx context.Context, request *AppResetPasswordRequest) (response *AppResetPasswordResponse, err error) {
    if request == nil {
        request = NewAppResetPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppResetPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppResetPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewAppSecureAddDeviceRequest() (request *AppSecureAddDeviceRequest) {
    request = &AppSecureAddDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppSecureAddDevice")
    
    
    return
}

func NewAppSecureAddDeviceResponse() (response *AppSecureAddDeviceResponse) {
    response = &AppSecureAddDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppSecureAddDevice
// 用户绑定设备，绑定后可以在APP端进行控制。绑定设备前需调用“获取设备绑定签名”接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTEXPIREDSIGNATURE = "InvalidParameter.IotExpiredSignature"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTINVALIDSIGNATURE = "InvalidParameter.IotInvalidSignature"
//  RESOURCEINUSE_IOTAPPLICATIONDEVICEEXISTS = "ResourceInUse.IotApplicationDeviceExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
func (c *Client) AppSecureAddDevice(request *AppSecureAddDeviceRequest) (response *AppSecureAddDeviceResponse, err error) {
    return c.AppSecureAddDeviceWithContext(context.Background(), request)
}

// AppSecureAddDevice
// 用户绑定设备，绑定后可以在APP端进行控制。绑定设备前需调用“获取设备绑定签名”接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTEXPIREDSIGNATURE = "InvalidParameter.IotExpiredSignature"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTINVALIDSIGNATURE = "InvalidParameter.IotInvalidSignature"
//  RESOURCEINUSE_IOTAPPLICATIONDEVICEEXISTS = "ResourceInUse.IotApplicationDeviceExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
func (c *Client) AppSecureAddDeviceWithContext(ctx context.Context, request *AppSecureAddDeviceRequest) (response *AppSecureAddDeviceResponse, err error) {
    if request == nil {
        request = NewAppSecureAddDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppSecureAddDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppSecureAddDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewAppUpdateDeviceRequest() (request *AppUpdateDeviceRequest) {
    request = &AppUpdateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppUpdateDevice")
    
    
    return
}

func NewAppUpdateDeviceResponse() (response *AppUpdateDeviceResponse) {
    response = &AppUpdateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppUpdateDevice
// 修改设备别名，便于用户个性化定义设备的名称
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppUpdateDevice(request *AppUpdateDeviceRequest) (response *AppUpdateDeviceResponse, err error) {
    return c.AppUpdateDeviceWithContext(context.Background(), request)
}

// AppUpdateDevice
// 修改设备别名，便于用户个性化定义设备的名称
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"
//  INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppUpdateDeviceWithContext(ctx context.Context, request *AppUpdateDeviceRequest) (response *AppUpdateDeviceResponse, err error) {
    if request == nil {
        request = NewAppUpdateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppUpdateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppUpdateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewAppUpdateUserRequest() (request *AppUpdateUserRequest) {
    request = &AppUpdateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AppUpdateUser")
    
    
    return
}

func NewAppUpdateUserResponse() (response *AppUpdateUserResponse) {
    response = &AppUpdateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AppUpdateUser
// 修改用户信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppUpdateUser(request *AppUpdateUserRequest) (response *AppUpdateUserResponse, err error) {
    return c.AppUpdateUserWithContext(context.Background(), request)
}

// AppUpdateUser
// 修改用户信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"
//  RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"
func (c *Client) AppUpdateUserWithContext(ctx context.Context, request *AppUpdateUserRequest) (response *AppUpdateUserResponse, err error) {
    if request == nil {
        request = NewAppUpdateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppUpdateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppUpdateUserResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSubDeviceToGatewayProductRequest() (request *AssociateSubDeviceToGatewayProductRequest) {
    request = &AssociateSubDeviceToGatewayProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "AssociateSubDeviceToGatewayProduct")
    
    
    return
}

func NewAssociateSubDeviceToGatewayProductResponse() (response *AssociateSubDeviceToGatewayProductResponse) {
    response = &AssociateSubDeviceToGatewayProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSubDeviceToGatewayProduct
// 关联子设备产品和网关产品
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPRODUCTINVALIDGATEWAYPRODUCTID = "InvalidParameter.IotProductInvalidGatewayProductId"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDSUBDEVICEPRODUCTID = "InvalidParameter.IotProductInvalidSubDeviceProductId"
func (c *Client) AssociateSubDeviceToGatewayProduct(request *AssociateSubDeviceToGatewayProductRequest) (response *AssociateSubDeviceToGatewayProductResponse, err error) {
    return c.AssociateSubDeviceToGatewayProductWithContext(context.Background(), request)
}

// AssociateSubDeviceToGatewayProduct
// 关联子设备产品和网关产品
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPRODUCTINVALIDGATEWAYPRODUCTID = "InvalidParameter.IotProductInvalidGatewayProductId"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDSUBDEVICEPRODUCTID = "InvalidParameter.IotProductInvalidSubDeviceProductId"
func (c *Client) AssociateSubDeviceToGatewayProductWithContext(ctx context.Context, request *AssociateSubDeviceToGatewayProductRequest) (response *AssociateSubDeviceToGatewayProductResponse, err error) {
    if request == nil {
        request = NewAssociateSubDeviceToGatewayProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSubDeviceToGatewayProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateSubDeviceToGatewayProductResponse()
    err = c.Send(request, response)
    return
}

func NewDeactivateRuleRequest() (request *DeactivateRuleRequest) {
    request = &DeactivateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "DeactivateRule")
    
    
    return
}

func NewDeactivateRuleResponse() (response *DeactivateRuleResponse) {
    response = &DeactivateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeactivateRule
// 禁用规则
//
// 可能返回的错误码:
//  INTERNALERROR_IOTACTIONSYSTEMERROR = "InternalError.IotActionSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCENOTFOUND_MQRULERULEIDNOTEXISTS = "ResourceNotFound.MqruleRuleIdNotExists"
func (c *Client) DeactivateRule(request *DeactivateRuleRequest) (response *DeactivateRuleResponse, err error) {
    return c.DeactivateRuleWithContext(context.Background(), request)
}

// DeactivateRule
// 禁用规则
//
// 可能返回的错误码:
//  INTERNALERROR_IOTACTIONSYSTEMERROR = "InternalError.IotActionSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCENOTFOUND_MQRULERULEIDNOTEXISTS = "ResourceNotFound.MqruleRuleIdNotExists"
func (c *Client) DeactivateRuleWithContext(ctx context.Context, request *DeactivateRuleRequest) (response *DeactivateRuleResponse, err error) {
    if request == nil {
        request = NewDeactivateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeactivateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeactivateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "DeleteDevice")
    
    
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDevice
// 提供在指定的产品Id下删除一个设备的能力。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTDEVICEOPTOOOFTEN = "LimitExceeded.IotDeviceOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_MQIOTRESOURCENOTEXISTS = "ResourceUnavailable.MqiotResourceNotExists"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    return c.DeleteDeviceWithContext(context.Background(), request)
}

// DeleteDevice
// 提供在指定的产品Id下删除一个设备的能力。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTDEVICEOPTOOOFTEN = "LimitExceeded.IotDeviceOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_MQIOTRESOURCENOTEXISTS = "ResourceUnavailable.MqiotResourceNotExists"
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

func NewDeleteProductRequest() (request *DeleteProductRequest) {
    request = &DeleteProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "DeleteProduct")
    
    
    return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
    response = &DeleteProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProduct
// 删除用户指定的产品Id对应的信息。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTPRODUCTOPTOOOFTEN = "LimitExceeded.IotProductOpTooOften"
//  RESOURCEINUSE_IOTDEVICEEXISTS = "ResourceInUse.IotDeviceExists"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTRULEEXISTS = "ResourceInUse.IotRuleExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    return c.DeleteProductWithContext(context.Background(), request)
}

// DeleteProduct
// 删除用户指定的产品Id对应的信息。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTPRODUCTOPTOOOFTEN = "LimitExceeded.IotProductOpTooOften"
//  RESOURCEINUSE_IOTDEVICEEXISTS = "ResourceInUse.IotDeviceExists"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCEINUSE_IOTRULEEXISTS = "ResourceInUse.IotRuleExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
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

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
    request = &DeleteRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "DeleteRule")
    
    
    return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
    response = &DeleteRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRule
// 删除规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_IOTRULEISACTIVE = "ResourceUnavailable.IotRuleIsActive"
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    return c.DeleteRuleWithContext(context.Background(), request)
}

// DeleteRule
// 删除规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_IOTRULEISACTIVE = "ResourceUnavailable.IotRuleIsActive"
func (c *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
    if request == nil {
        request = NewDeleteRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopic
// 删除Topic
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTTOPICOPTOOOFTEN = "LimitExceeded.IotTopicOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// 删除Topic
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTTOPICOPTOOOFTEN = "LimitExceeded.IotTopicOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataHistoryRequest() (request *GetDataHistoryRequest) {
    request = &GetDataHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDataHistory")
    
    
    return
}

func NewGetDataHistoryResponse() (response *GetDataHistoryResponse) {
    response = &GetDataHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDataHistory
// 批量获取设备某一段时间范围的设备上报数据。该接口适用于使用高级版类型的产品
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDataHistory(request *GetDataHistoryRequest) (response *GetDataHistoryResponse, err error) {
    return c.GetDataHistoryWithContext(context.Background(), request)
}

// GetDataHistory
// 批量获取设备某一段时间范围的设备上报数据。该接口适用于使用高级版类型的产品
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDataHistoryWithContext(ctx context.Context, request *GetDataHistoryRequest) (response *GetDataHistoryResponse, err error) {
    if request == nil {
        request = NewGetDataHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDataHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDataHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewGetDebugLogRequest() (request *GetDebugLogRequest) {
    request = &GetDebugLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDebugLog")
    
    
    return
}

func NewGetDebugLogResponse() (response *GetDebugLogResponse) {
    response = &GetDebugLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDebugLog
// 获取设备的调试日志，用于定位问题
//
// 可能返回的错误码:
//  INTERNALERROR_IOTLOGSYSTEMERROR = "InternalError.IotLogSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDebugLog(request *GetDebugLogRequest) (response *GetDebugLogResponse, err error) {
    return c.GetDebugLogWithContext(context.Background(), request)
}

// GetDebugLog
// 获取设备的调试日志，用于定位问题
//
// 可能返回的错误码:
//  INTERNALERROR_IOTLOGSYSTEMERROR = "InternalError.IotLogSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDebugLogWithContext(ctx context.Context, request *GetDebugLogRequest) (response *GetDebugLogResponse, err error) {
    if request == nil {
        request = NewGetDebugLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDebugLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDebugLogResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceRequest() (request *GetDeviceRequest) {
    request = &GetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDevice")
    
    
    return
}

func NewGetDeviceResponse() (response *GetDeviceResponse) {
    response = &GetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDevice
// 提供查询某个设备详细信息的能力。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDevice(request *GetDeviceRequest) (response *GetDeviceResponse, err error) {
    return c.GetDeviceWithContext(context.Background(), request)
}

// GetDevice
// 提供查询某个设备详细信息的能力。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
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

func NewGetDeviceDataRequest() (request *GetDeviceDataRequest) {
    request = &GetDeviceDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceData")
    
    
    return
}

func NewGetDeviceDataResponse() (response *GetDeviceDataResponse) {
    response = &GetDeviceDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDeviceData
// 获取某个设备当前上报到云端的数据，该接口适用于使用数据模板协议的产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceData(request *GetDeviceDataRequest) (response *GetDeviceDataResponse, err error) {
    return c.GetDeviceDataWithContext(context.Background(), request)
}

// GetDeviceData
// 获取某个设备当前上报到云端的数据，该接口适用于使用数据模板协议的产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceDataWithContext(ctx context.Context, request *GetDeviceDataRequest) (response *GetDeviceDataResponse, err error) {
    if request == nil {
        request = NewGetDeviceDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceData require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceLogRequest() (request *GetDeviceLogRequest) {
    request = &GetDeviceLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceLog")
    
    
    return
}

func NewGetDeviceLogResponse() (response *GetDeviceLogResponse) {
    response = &GetDeviceLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDeviceLog
// 批量获取设备与云端的详细通信日志，该接口适用于使用高级版类型的产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTLOGSYSTEMERROR = "InternalError.IotLogSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceLog(request *GetDeviceLogRequest) (response *GetDeviceLogResponse, err error) {
    return c.GetDeviceLogWithContext(context.Background(), request)
}

// GetDeviceLog
// 批量获取设备与云端的详细通信日志，该接口适用于使用高级版类型的产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTLOGSYSTEMERROR = "InternalError.IotLogSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceLogWithContext(ctx context.Context, request *GetDeviceLogRequest) (response *GetDeviceLogResponse, err error) {
    if request == nil {
        request = NewGetDeviceLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceLogResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceSignaturesRequest() (request *GetDeviceSignaturesRequest) {
    request = &GetDeviceSignaturesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceSignatures")
    
    
    return
}

func NewGetDeviceSignaturesResponse() (response *GetDeviceSignaturesResponse) {
    response = &GetDeviceSignaturesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDeviceSignatures
// 获取设备绑定签名，用于用户绑定某个设备的应用场景
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceSignatures(request *GetDeviceSignaturesRequest) (response *GetDeviceSignaturesResponse, err error) {
    return c.GetDeviceSignaturesWithContext(context.Background(), request)
}

// GetDeviceSignatures
// 获取设备绑定签名，用于用户绑定某个设备的应用场景
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceSignaturesWithContext(ctx context.Context, request *GetDeviceSignaturesRequest) (response *GetDeviceSignaturesResponse, err error) {
    if request == nil {
        request = NewGetDeviceSignaturesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceSignatures require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceSignaturesResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceStatisticsRequest() (request *GetDeviceStatisticsRequest) {
    request = &GetDeviceStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceStatistics")
    
    
    return
}

func NewGetDeviceStatisticsResponse() (response *GetDeviceStatisticsResponse) {
    response = &GetDeviceStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDeviceStatistics
// 查询某段时间范围内产品的在线、激活设备数
//
// 可能返回的错误码:
//  INTERNALERROR_IOTDBERROR = "InternalError.IotDbError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTSTATINVALIDDATE = "InvalidParameter.IotStatInvalidDate"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceStatistics(request *GetDeviceStatisticsRequest) (response *GetDeviceStatisticsResponse, err error) {
    return c.GetDeviceStatisticsWithContext(context.Background(), request)
}

// GetDeviceStatistics
// 查询某段时间范围内产品的在线、激活设备数
//
// 可能返回的错误码:
//  INTERNALERROR_IOTDBERROR = "InternalError.IotDbError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTSTATINVALIDDATE = "InvalidParameter.IotStatInvalidDate"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDeviceStatisticsWithContext(ctx context.Context, request *GetDeviceStatisticsRequest) (response *GetDeviceStatisticsResponse, err error) {
    if request == nil {
        request = NewGetDeviceStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceStatusesRequest() (request *GetDeviceStatusesRequest) {
    request = &GetDeviceStatusesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDeviceStatuses")
    
    
    return
}

func NewGetDeviceStatusesResponse() (response *GetDeviceStatusesResponse) {
    response = &GetDeviceStatusesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDeviceStatuses
// 批量获取设备的当前状态，状态包括在线、离线或未激活状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
func (c *Client) GetDeviceStatuses(request *GetDeviceStatusesRequest) (response *GetDeviceStatusesResponse, err error) {
    return c.GetDeviceStatusesWithContext(context.Background(), request)
}

// GetDeviceStatuses
// 批量获取设备的当前状态，状态包括在线、离线或未激活状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
func (c *Client) GetDeviceStatusesWithContext(ctx context.Context, request *GetDeviceStatusesRequest) (response *GetDeviceStatusesResponse, err error) {
    if request == nil {
        request = NewGetDeviceStatusesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceStatuses require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceStatusesResponse()
    err = c.Send(request, response)
    return
}

func NewGetDevicesRequest() (request *GetDevicesRequest) {
    request = &GetDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetDevices")
    
    
    return
}

func NewGetDevicesResponse() (response *GetDevicesResponse) {
    response = &GetDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDevices
// 提供分页查询某个产品Id下设备信息的能力。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetDevices(request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    return c.GetDevicesWithContext(context.Background(), request)
}

// GetDevices
// 提供分页查询某个产品Id下设备信息的能力。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
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

func NewGetProductRequest() (request *GetProductRequest) {
    request = &GetProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetProduct")
    
    
    return
}

func NewGetProductResponse() (response *GetProductResponse) {
    response = &GetProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetProduct
// 获取产品定义的详细信息，包括产品名称、产品描述，鉴权模式等信息。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetProduct(request *GetProductRequest) (response *GetProductResponse, err error) {
    return c.GetProductWithContext(context.Background(), request)
}

// GetProduct
// 获取产品定义的详细信息，包括产品名称、产品描述，鉴权模式等信息。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetProductWithContext(ctx context.Context, request *GetProductRequest) (response *GetProductResponse, err error) {
    if request == nil {
        request = NewGetProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProductResponse()
    err = c.Send(request, response)
    return
}

func NewGetProductsRequest() (request *GetProductsRequest) {
    request = &GetProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetProducts")
    
    
    return
}

func NewGetProductsResponse() (response *GetProductsResponse) {
    response = &GetProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetProducts
// 获取用户在物联网套件所创建的所有产品信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetProducts(request *GetProductsRequest) (response *GetProductsResponse, err error) {
    return c.GetProductsWithContext(context.Background(), request)
}

// GetProducts
// 获取用户在物联网套件所创建的所有产品信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetProductsWithContext(ctx context.Context, request *GetProductsRequest) (response *GetProductsResponse, err error) {
    if request == nil {
        request = NewGetProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProductsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRuleRequest() (request *GetRuleRequest) {
    request = &GetRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetRule")
    
    
    return
}

func NewGetRuleResponse() (response *GetRuleResponse) {
    response = &GetRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRule
// 获取转发规则信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetRule(request *GetRuleRequest) (response *GetRuleResponse, err error) {
    return c.GetRuleWithContext(context.Background(), request)
}

// GetRule
// 获取转发规则信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetRuleWithContext(ctx context.Context, request *GetRuleRequest) (response *GetRuleResponse, err error) {
    if request == nil {
        request = NewGetRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetRulesRequest() (request *GetRulesRequest) {
    request = &GetRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetRules")
    
    
    return
}

func NewGetRulesResponse() (response *GetRulesResponse) {
    response = &GetRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRules
// 获取转发规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetRules(request *GetRulesRequest) (response *GetRulesResponse, err error) {
    return c.GetRulesWithContext(context.Background(), request)
}

// GetRules
// 获取转发规则列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetRulesWithContext(ctx context.Context, request *GetRulesRequest) (response *GetRulesResponse, err error) {
    if request == nil {
        request = NewGetRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRulesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicRequest() (request *GetTopicRequest) {
    request = &GetTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetTopic")
    
    
    return
}

func NewGetTopicResponse() (response *GetTopicResponse) {
    response = &GetTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTopic
// 获取Topic信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetTopic(request *GetTopicRequest) (response *GetTopicResponse, err error) {
    return c.GetTopicWithContext(context.Background(), request)
}

// GetTopic
// 获取Topic信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetTopicWithContext(ctx context.Context, request *GetTopicRequest) (response *GetTopicResponse, err error) {
    if request == nil {
        request = NewGetTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetTopicsRequest() (request *GetTopicsRequest) {
    request = &GetTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "GetTopics")
    
    
    return
}

func NewGetTopicsResponse() (response *GetTopicsResponse) {
    response = &GetTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTopics
// 获取Topic列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetTopics(request *GetTopicsRequest) (response *GetTopicsResponse, err error) {
    return c.GetTopicsWithContext(context.Background(), request)
}

// GetTopics
// 获取Topic列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
func (c *Client) GetTopicsWithContext(ctx context.Context, request *GetTopicsRequest) (response *GetTopicsResponse, err error) {
    if request == nil {
        request = NewGetTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewIssueDeviceControlRequest() (request *IssueDeviceControlRequest) {
    request = &IssueDeviceControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "IssueDeviceControl")
    
    
    return
}

func NewIssueDeviceControlResponse() (response *IssueDeviceControlResponse) {
    response = &IssueDeviceControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IssueDeviceControl
// 提供下发控制指令到指定设备的能力，该接口适用于使用高级版类型的产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) IssueDeviceControl(request *IssueDeviceControlRequest) (response *IssueDeviceControlResponse, err error) {
    return c.IssueDeviceControlWithContext(context.Background(), request)
}

// IssueDeviceControl
// 提供下发控制指令到指定设备的能力，该接口适用于使用高级版类型的产品。
//
// 可能返回的错误码:
//  INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) IssueDeviceControlWithContext(ctx context.Context, request *IssueDeviceControlRequest) (response *IssueDeviceControlResponse, err error) {
    if request == nil {
        request = NewIssueDeviceControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IssueDeviceControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewIssueDeviceControlResponse()
    err = c.Send(request, response)
    return
}

func NewPublishMsgRequest() (request *PublishMsgRequest) {
    request = &PublishMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "PublishMsg")
    
    
    return
}

func NewPublishMsgResponse() (response *PublishMsgResponse) {
    response = &PublishMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PublishMsg
// 提供向指定的Topic发布消息的能力，常用于向设备下发控制指令。该接口只适用于产品版本为“基础版”类型的产品，使用高级版的产品需使用“下发设备控制指令”接口
//
// 可能返回的错误码:
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDAUTHTYPE = "InvalidParameter.IotProductInvalidAuthType"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_MQIOTRESOURCENOTEXISTS = "ResourceUnavailable.MqiotResourceNotExists"
func (c *Client) PublishMsg(request *PublishMsgRequest) (response *PublishMsgResponse, err error) {
    return c.PublishMsgWithContext(context.Background(), request)
}

// PublishMsg
// 提供向指定的Topic发布消息的能力，常用于向设备下发控制指令。该接口只适用于产品版本为“基础版”类型的产品，使用高级版的产品需使用“下发设备控制指令”接口
//
// 可能返回的错误码:
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDAUTHTYPE = "InvalidParameter.IotProductInvalidAuthType"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_MQIOTRESOURCENOTEXISTS = "ResourceUnavailable.MqiotResourceNotExists"
func (c *Client) PublishMsgWithContext(ctx context.Context, request *PublishMsgRequest) (response *PublishMsgResponse, err error) {
    if request == nil {
        request = NewPublishMsgRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PublishMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewPublishMsgResponse()
    err = c.Send(request, response)
    return
}

func NewResetDeviceRequest() (request *ResetDeviceRequest) {
    request = &ResetDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "ResetDevice")
    
    
    return
}

func NewResetDeviceResponse() (response *ResetDeviceResponse) {
    response = &ResetDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetDevice
// 重置设备操作，将会为设备生成新的证书及清空最新数据，需谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDAUTHTYPE = "InvalidParameter.IotProductInvalidAuthType"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_MQIOTRESOURCENOTEXISTS = "ResourceUnavailable.MqiotResourceNotExists"
func (c *Client) ResetDevice(request *ResetDeviceRequest) (response *ResetDeviceResponse, err error) {
    return c.ResetDeviceWithContext(context.Background(), request)
}

// ResetDevice
// 重置设备操作，将会为设备生成新的证书及清空最新数据，需谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDAUTHTYPE = "InvalidParameter.IotProductInvalidAuthType"
//  RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_MQIOTRESOURCENOTEXISTS = "ResourceUnavailable.MqiotResourceNotExists"
func (c *Client) ResetDeviceWithContext(ctx context.Context, request *ResetDeviceRequest) (response *ResetDeviceResponse, err error) {
    if request == nil {
        request = NewResetDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewUnassociateSubDeviceFromGatewayProductRequest() (request *UnassociateSubDeviceFromGatewayProductRequest) {
    request = &UnassociateSubDeviceFromGatewayProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "UnassociateSubDeviceFromGatewayProduct")
    
    
    return
}

func NewUnassociateSubDeviceFromGatewayProductResponse() (response *UnassociateSubDeviceFromGatewayProductResponse) {
    response = &UnassociateSubDeviceFromGatewayProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnassociateSubDeviceFromGatewayProduct
// 取消子设备产品与网关设备产品的关联
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPRODUCTINVALIDGATEWAYPRODUCTID = "InvalidParameter.IotProductInvalidGatewayProductId"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDSUBDEVICEPRODUCTID = "InvalidParameter.IotProductInvalidSubDeviceProductId"
func (c *Client) UnassociateSubDeviceFromGatewayProduct(request *UnassociateSubDeviceFromGatewayProductRequest) (response *UnassociateSubDeviceFromGatewayProductResponse, err error) {
    return c.UnassociateSubDeviceFromGatewayProductWithContext(context.Background(), request)
}

// UnassociateSubDeviceFromGatewayProduct
// 取消子设备产品与网关设备产品的关联
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPRODUCTINVALIDGATEWAYPRODUCTID = "InvalidParameter.IotProductInvalidGatewayProductId"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDSUBDEVICEPRODUCTID = "InvalidParameter.IotProductInvalidSubDeviceProductId"
func (c *Client) UnassociateSubDeviceFromGatewayProductWithContext(ctx context.Context, request *UnassociateSubDeviceFromGatewayProductRequest) (response *UnassociateSubDeviceFromGatewayProductResponse, err error) {
    if request == nil {
        request = NewUnassociateSubDeviceFromGatewayProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnassociateSubDeviceFromGatewayProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnassociateSubDeviceFromGatewayProductResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProductRequest() (request *UpdateProductRequest) {
    request = &UpdateProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "UpdateProduct")
    
    
    return
}

func NewUpdateProductResponse() (response *UpdateProductResponse) {
    response = &UpdateProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateProduct
// 提供修改产品信息及数据模板的能力。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATE = "InvalidParameter.IotProductInvalidDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATERANGE = "InvalidParameter.IotProductInvalidDataTemplateRange"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) UpdateProduct(request *UpdateProductRequest) (response *UpdateProductResponse, err error) {
    return c.UpdateProductWithContext(context.Background(), request)
}

// UpdateProduct
// 提供修改产品信息及数据模板的能力。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATE = "InvalidParameter.IotProductInvalidDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATERANGE = "InvalidParameter.IotProductInvalidDataTemplateRange"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) UpdateProductWithContext(ctx context.Context, request *UpdateProductRequest) (response *UpdateProductResponse, err error) {
    if request == nil {
        request = NewUpdateProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProductResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRuleRequest() (request *UpdateRuleRequest) {
    request = &UpdateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iot", APIVersion, "UpdateRule")
    
    
    return
}

func NewUpdateRuleResponse() (response *UpdateRuleResponse) {
    response = &UpdateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRule
// 更新规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTEMPTYDATATEMPLATE = "InvalidParameter.IotProductEmptyDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATAPROTOCOL = "InvalidParameter.IotProductInvalidDataProtocol"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTSUBACCOUNTNOTEXISTS = "ResourceNotFound.IotSubAccountNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_IOTRULEISACTIVE = "ResourceUnavailable.IotRuleIsActive"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) UpdateRule(request *UpdateRuleRequest) (response *UpdateRuleResponse, err error) {
    return c.UpdateRuleWithContext(context.Background(), request)
}

// UpdateRule
// 更新规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"
//  INVALIDPARAMETER_IOTPRODUCTEMPTYDATATEMPLATE = "InvalidParameter.IotProductEmptyDataTemplate"
//  INVALIDPARAMETER_IOTPRODUCTINVALIDDATAPROTOCOL = "InvalidParameter.IotProductInvalidDataProtocol"
//  LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"
//  RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"
//  RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"
//  RESOURCENOTFOUND_IOTSUBACCOUNTNOTEXISTS = "ResourceNotFound.IotSubAccountNotExists"
//  RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"
//  RESOURCEUNAVAILABLE_IOTRULEISACTIVE = "ResourceUnavailable.IotRuleIsActive"
//  UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
func (c *Client) UpdateRuleWithContext(ctx context.Context, request *UpdateRuleRequest) (response *UpdateRuleResponse, err error) {
    if request == nil {
        request = NewUpdateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRuleResponse()
    err = c.Send(request, response)
    return
}
