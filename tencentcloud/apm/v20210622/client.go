// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20210622

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-06-22"

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


func NewCreateApmInstanceRequest() (request *CreateApmInstanceRequest) {
    request = &CreateApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmInstance")
    
    
    return
}

func NewCreateApmInstanceResponse() (response *CreateApmInstanceResponse) {
    response = &CreateApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApmInstance
// 业务购买 APM 业务系统，调用该接口创建
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_REGIONNOTSUPPORT = "FailedOperation.RegionNotSupport"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) CreateApmInstance(request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
    return c.CreateApmInstanceWithContext(context.Background(), request)
}

// CreateApmInstance
// 业务购买 APM 业务系统，调用该接口创建
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_REGIONNOTSUPPORT = "FailedOperation.RegionNotSupport"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) CreateApmInstanceWithContext(ctx context.Context, request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
    if request == nil {
        request = NewCreateApmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateApmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApmPrometheusRuleRequest() (request *CreateApmPrometheusRuleRequest) {
    request = &CreateApmPrometheusRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmPrometheusRule")
    
    
    return
}

func NewCreateApmPrometheusRuleResponse() (response *CreateApmPrometheusRuleResponse) {
    response = &CreateApmPrometheusRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApmPrometheusRule
// 用于创建apm业务系统与Prometheus实例的指标匹配规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULENAMECONFLICT = "FailedOperation.PromRuleNameConflict"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) CreateApmPrometheusRule(request *CreateApmPrometheusRuleRequest) (response *CreateApmPrometheusRuleResponse, err error) {
    return c.CreateApmPrometheusRuleWithContext(context.Background(), request)
}

// CreateApmPrometheusRule
// 用于创建apm业务系统与Prometheus实例的指标匹配规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULENAMECONFLICT = "FailedOperation.PromRuleNameConflict"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) CreateApmPrometheusRuleWithContext(ctx context.Context, request *CreateApmPrometheusRuleRequest) (response *CreateApmPrometheusRuleResponse, err error) {
    if request == nil {
        request = NewCreateApmPrometheusRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateApmPrometheusRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApmPrometheusRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApmPrometheusRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApmSampleConfigRequest() (request *CreateApmSampleConfigRequest) {
    request = &CreateApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmSampleConfig")
    
    
    return
}

func NewCreateApmSampleConfigResponse() (response *CreateApmSampleConfigResponse) {
    response = &CreateApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApmSampleConfig
// 创建采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLENAMECONFLICT = "FailedOperation.SampleNameConflict"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) CreateApmSampleConfig(request *CreateApmSampleConfigRequest) (response *CreateApmSampleConfigResponse, err error) {
    return c.CreateApmSampleConfigWithContext(context.Background(), request)
}

// CreateApmSampleConfig
// 创建采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLENAMECONFLICT = "FailedOperation.SampleNameConflict"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) CreateApmSampleConfigWithContext(ctx context.Context, request *CreateApmSampleConfigRequest) (response *CreateApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewCreateApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProfileTaskRequest() (request *CreateProfileTaskRequest) {
    request = &CreateProfileTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateProfileTask")
    
    
    return
}

func NewCreateProfileTaskResponse() (response *CreateProfileTaskResponse) {
    response = &CreateProfileTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProfileTask
// 创建事件任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTNOTONLINEERROR = "FailedOperation.AgentNotOnlineError"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTERROR = "FailedOperation.AgentVersionNotSupportError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
func (c *Client) CreateProfileTask(request *CreateProfileTaskRequest) (response *CreateProfileTaskResponse, err error) {
    return c.CreateProfileTaskWithContext(context.Background(), request)
}

// CreateProfileTask
// 创建事件任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTNOTONLINEERROR = "FailedOperation.AgentNotOnlineError"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTERROR = "FailedOperation.AgentVersionNotSupportError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
func (c *Client) CreateProfileTaskWithContext(ctx context.Context, request *CreateProfileTaskRequest) (response *CreateProfileTaskResponse, err error) {
    if request == nil {
        request = NewCreateProfileTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateProfileTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProfileTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProfileTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApmSampleConfigRequest() (request *DeleteApmSampleConfigRequest) {
    request = &DeleteApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DeleteApmSampleConfig")
    
    
    return
}

func NewDeleteApmSampleConfigResponse() (response *DeleteApmSampleConfigResponse) {
    response = &DeleteApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApmSampleConfig
// 删除采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DeleteApmSampleConfig(request *DeleteApmSampleConfigRequest) (response *DeleteApmSampleConfigResponse, err error) {
    return c.DeleteApmSampleConfigWithContext(context.Background(), request)
}

// DeleteApmSampleConfig
// 删除采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DeleteApmSampleConfigWithContext(ctx context.Context, request *DeleteApmSampleConfigRequest) (response *DeleteApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewDeleteApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DeleteApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmAgentRequest() (request *DescribeApmAgentRequest) {
    request = &DescribeApmAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAgent")
    
    
    return
}

func NewDescribeApmAgentResponse() (response *DescribeApmAgentResponse) {
    response = &DescribeApmAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmAgent
// 获取 APM 接入点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmAgent(request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
    return c.DescribeApmAgentWithContext(context.Background(), request)
}

// DescribeApmAgent
// 获取 APM 接入点
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmAgentWithContext(ctx context.Context, request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
    if request == nil {
        request = NewDescribeApmAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmApplicationConfigRequest() (request *DescribeApmApplicationConfigRequest) {
    request = &DescribeApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmApplicationConfig")
    
    
    return
}

func NewDescribeApmApplicationConfigResponse() (response *DescribeApmApplicationConfigResponse) {
    response = &DescribeApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmApplicationConfig
// 查询应用配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeApmApplicationConfig(request *DescribeApmApplicationConfigRequest) (response *DescribeApmApplicationConfigResponse, err error) {
    return c.DescribeApmApplicationConfigWithContext(context.Background(), request)
}

// DescribeApmApplicationConfig
// 查询应用配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeApmApplicationConfigWithContext(ctx context.Context, request *DescribeApmApplicationConfigRequest) (response *DescribeApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmAssociationRequest() (request *DescribeApmAssociationRequest) {
    request = &DescribeApmAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAssociation")
    
    
    return
}

func NewDescribeApmAssociationResponse() (response *DescribeApmAssociationResponse) {
    response = &DescribeApmAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmAssociation
// 用于查询apm业务系统与其他产品的关联关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
func (c *Client) DescribeApmAssociation(request *DescribeApmAssociationRequest) (response *DescribeApmAssociationResponse, err error) {
    return c.DescribeApmAssociationWithContext(context.Background(), request)
}

// DescribeApmAssociation
// 用于查询apm业务系统与其他产品的关联关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
func (c *Client) DescribeApmAssociationWithContext(ctx context.Context, request *DescribeApmAssociationRequest) (response *DescribeApmAssociationResponse, err error) {
    if request == nil {
        request = NewDescribeApmAssociationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmAssociation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmAssociation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmAssociationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmInstancesRequest() (request *DescribeApmInstancesRequest) {
    request = &DescribeApmInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInstances")
    
    
    return
}

func NewDescribeApmInstancesResponse() (response *DescribeApmInstancesResponse) {
    response = &DescribeApmInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmInstances
// 获取 APM 业务系统列表
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmInstances(request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
    return c.DescribeApmInstancesWithContext(context.Background(), request)
}

// DescribeApmInstances
// 获取 APM 业务系统列表
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmInstancesWithContext(ctx context.Context, request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeApmInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmPrometheusRuleRequest() (request *DescribeApmPrometheusRuleRequest) {
    request = &DescribeApmPrometheusRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmPrometheusRule")
    
    
    return
}

func NewDescribeApmPrometheusRuleResponse() (response *DescribeApmPrometheusRuleResponse) {
    response = &DescribeApmPrometheusRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmPrometheusRule
// 用于查询apm业务系统与Prometheus实例的指标匹配规则
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmPrometheusRule(request *DescribeApmPrometheusRuleRequest) (response *DescribeApmPrometheusRuleResponse, err error) {
    return c.DescribeApmPrometheusRuleWithContext(context.Background(), request)
}

// DescribeApmPrometheusRule
// 用于查询apm业务系统与Prometheus实例的指标匹配规则
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmPrometheusRuleWithContext(ctx context.Context, request *DescribeApmPrometheusRuleRequest) (response *DescribeApmPrometheusRuleResponse, err error) {
    if request == nil {
        request = NewDescribeApmPrometheusRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmPrometheusRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmPrometheusRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmPrometheusRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmSampleConfigRequest() (request *DescribeApmSampleConfigRequest) {
    request = &DescribeApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmSampleConfig")
    
    
    return
}

func NewDescribeApmSampleConfigResponse() (response *DescribeApmSampleConfigResponse) {
    response = &DescribeApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmSampleConfig
// 查询采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_SAMPLENOTFOUND = "FailedOperation.SampleNotFound"
func (c *Client) DescribeApmSampleConfig(request *DescribeApmSampleConfigRequest) (response *DescribeApmSampleConfigResponse, err error) {
    return c.DescribeApmSampleConfigWithContext(context.Background(), request)
}

// DescribeApmSampleConfig
// 查询采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_SAMPLENOTFOUND = "FailedOperation.SampleNotFound"
func (c *Client) DescribeApmSampleConfigWithContext(ctx context.Context, request *DescribeApmSampleConfigRequest) (response *DescribeApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewDescribeApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmServiceMetricRequest() (request *DescribeApmServiceMetricRequest) {
    request = &DescribeApmServiceMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmServiceMetric")
    
    
    return
}

func NewDescribeApmServiceMetricResponse() (response *DescribeApmServiceMetricResponse) {
    response = &DescribeApmServiceMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmServiceMetric
// 获取 APM 应用指标列表
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"
//  FAILEDOPERATION_SERVICENOTMATCHAPPIDERR = "FailedOperation.ServiceNotMatchAppIdErr"
func (c *Client) DescribeApmServiceMetric(request *DescribeApmServiceMetricRequest) (response *DescribeApmServiceMetricResponse, err error) {
    return c.DescribeApmServiceMetricWithContext(context.Background(), request)
}

// DescribeApmServiceMetric
// 获取 APM 应用指标列表
//
// 可能返回的错误码:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"
//  FAILEDOPERATION_SERVICENOTMATCHAPPIDERR = "FailedOperation.ServiceNotMatchAppIdErr"
func (c *Client) DescribeApmServiceMetricWithContext(ctx context.Context, request *DescribeApmServiceMetricRequest) (response *DescribeApmServiceMetricResponse, err error) {
    if request == nil {
        request = NewDescribeApmServiceMetricRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmServiceMetric")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmServiceMetric require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmServiceMetricResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralApmApplicationConfigRequest() (request *DescribeGeneralApmApplicationConfigRequest) {
    request = &DescribeGeneralApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralApmApplicationConfig")
    
    
    return
}

func NewDescribeGeneralApmApplicationConfigResponse() (response *DescribeGeneralApmApplicationConfigResponse) {
    response = &DescribeGeneralApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralApmApplicationConfig
// 查询应用配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeGeneralApmApplicationConfig(request *DescribeGeneralApmApplicationConfigRequest) (response *DescribeGeneralApmApplicationConfigResponse, err error) {
    return c.DescribeGeneralApmApplicationConfigWithContext(context.Background(), request)
}

// DescribeGeneralApmApplicationConfig
// 查询应用配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeGeneralApmApplicationConfigWithContext(ctx context.Context, request *DescribeGeneralApmApplicationConfigRequest) (response *DescribeGeneralApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralMetricDataRequest() (request *DescribeGeneralMetricDataRequest) {
    request = &DescribeGeneralMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralMetricData")
    
    
    return
}

func NewDescribeGeneralMetricDataResponse() (response *DescribeGeneralMetricDataResponse) {
    response = &DescribeGeneralMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralMetricData
// 获取指标数据通用接口。用户根据需要上送请求参数，返回对应的指标数据。
//
// 接口调用频率限制为：20次/秒，1200次/分钟。单请求的数据点数限制为1440个。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_METRICFILTERSLACKPARAMS = "FailedOperation.MetricFiltersLackParams"
//  FAILEDOPERATION_QUERYTIMEINTERVALISNOTSUPPORTED = "FailedOperation.QueryTimeIntervalIsNotSupported"
//  FAILEDOPERATION_VIEWNAMENOTEXISTORILLEGAL = "FailedOperation.ViewNameNotExistOrIllegal"
//  INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"
//  INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"
//  INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"
//  INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"
//  INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"
//  INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
func (c *Client) DescribeGeneralMetricData(request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
    return c.DescribeGeneralMetricDataWithContext(context.Background(), request)
}

// DescribeGeneralMetricData
// 获取指标数据通用接口。用户根据需要上送请求参数，返回对应的指标数据。
//
// 接口调用频率限制为：20次/秒，1200次/分钟。单请求的数据点数限制为1440个。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_METRICFILTERSLACKPARAMS = "FailedOperation.MetricFiltersLackParams"
//  FAILEDOPERATION_QUERYTIMEINTERVALISNOTSUPPORTED = "FailedOperation.QueryTimeIntervalIsNotSupported"
//  FAILEDOPERATION_VIEWNAMENOTEXISTORILLEGAL = "FailedOperation.ViewNameNotExistOrIllegal"
//  INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"
//  INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"
//  INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"
//  INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"
//  INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"
//  INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
func (c *Client) DescribeGeneralMetricDataWithContext(ctx context.Context, request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralMetricDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralMetricData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralOTSpanListRequest() (request *DescribeGeneralOTSpanListRequest) {
    request = &DescribeGeneralOTSpanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralOTSpanList")
    
    
    return
}

func NewDescribeGeneralOTSpanListResponse() (response *DescribeGeneralOTSpanListResponse) {
    response = &DescribeGeneralOTSpanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralOTSpanList
// 通用查询 OpenTelemetry 调用链列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralOTSpanList(request *DescribeGeneralOTSpanListRequest) (response *DescribeGeneralOTSpanListResponse, err error) {
    return c.DescribeGeneralOTSpanListWithContext(context.Background(), request)
}

// DescribeGeneralOTSpanList
// 通用查询 OpenTelemetry 调用链列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralOTSpanListWithContext(ctx context.Context, request *DescribeGeneralOTSpanListRequest) (response *DescribeGeneralOTSpanListResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralOTSpanListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralOTSpanList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralOTSpanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralOTSpanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralSpanListRequest() (request *DescribeGeneralSpanListRequest) {
    request = &DescribeGeneralSpanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralSpanList")
    
    
    return
}

func NewDescribeGeneralSpanListResponse() (response *DescribeGeneralSpanListResponse) {
    response = &DescribeGeneralSpanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralSpanList
// 通用查询调用链列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralSpanList(request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
    return c.DescribeGeneralSpanListWithContext(context.Background(), request)
}

// DescribeGeneralSpanList
// 通用查询调用链列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralSpanListWithContext(ctx context.Context, request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralSpanListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralSpanList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralSpanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralSpanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricRecordsRequest() (request *DescribeMetricRecordsRequest) {
    request = &DescribeMetricRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricRecords")
    
    
    return
}

func NewDescribeMetricRecordsResponse() (response *DescribeMetricRecordsResponse) {
    response = &DescribeMetricRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricRecords
// 查询指标列表接口，查询指标更推荐使用DescribeGeneralMetricData接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricRecords(request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
    return c.DescribeMetricRecordsWithContext(context.Background(), request)
}

// DescribeMetricRecords
// 查询指标列表接口，查询指标更推荐使用DescribeGeneralMetricData接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricRecordsWithContext(ctx context.Context, request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeMetricRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeMetricRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceOverviewRequest() (request *DescribeServiceOverviewRequest) {
    request = &DescribeServiceOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceOverview")
    
    
    return
}

func NewDescribeServiceOverviewResponse() (response *DescribeServiceOverviewResponse) {
    response = &DescribeServiceOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceOverview
// 应用概览数据拉取
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServiceOverview(request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
    return c.DescribeServiceOverviewWithContext(context.Background(), request)
}

// DescribeServiceOverview
// 应用概览数据拉取
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServiceOverviewWithContext(ctx context.Context, request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeServiceOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeServiceOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagValuesRequest() (request *DescribeTagValuesRequest) {
    request = &DescribeTagValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeTagValues")
    
    
    return
}

func NewDescribeTagValuesResponse() (response *DescribeTagValuesResponse) {
    response = &DescribeTagValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagValues
// 根据维度名和过滤条件，查询维度数据.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DescribeTagValues(request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    return c.DescribeTagValuesWithContext(context.Background(), request)
}

// DescribeTagValues
// 根据维度名和过滤条件，查询维度数据.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DescribeTagValuesWithContext(ctx context.Context, request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    if request == nil {
        request = NewDescribeTagValuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeTagValues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagValuesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmApplicationConfigRequest() (request *ModifyApmApplicationConfigRequest) {
    request = &ModifyApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmApplicationConfig")
    
    
    return
}

func NewModifyApmApplicationConfigResponse() (response *ModifyApmApplicationConfigResponse) {
    response = &ModifyApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmApplicationConfig
// 修改应用配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTOPERATIONCONFIGINVALID = "FailedOperation.AgentOperationConfigInvalid"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
func (c *Client) ModifyApmApplicationConfig(request *ModifyApmApplicationConfigRequest) (response *ModifyApmApplicationConfigResponse, err error) {
    return c.ModifyApmApplicationConfigWithContext(context.Background(), request)
}

// ModifyApmApplicationConfig
// 修改应用配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTOPERATIONCONFIGINVALID = "FailedOperation.AgentOperationConfigInvalid"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
func (c *Client) ModifyApmApplicationConfigWithContext(ctx context.Context, request *ModifyApmApplicationConfigRequest) (response *ModifyApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewModifyApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmAssociationRequest() (request *ModifyApmAssociationRequest) {
    request = &ModifyApmAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmAssociation")
    
    
    return
}

func NewModifyApmAssociationResponse() (response *ModifyApmAssociationResponse) {
    response = &ModifyApmAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmAssociation
// 用于修改apm业务系统与其他产品的关联关系（包括创建和删除）
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASSOCIATIONMODIFYREQUESTNOTVALIDERROR = "FailedOperation.AssociationModifyRequestNotValidError"
//  FAILEDOPERATION_CKAFKADIFFASSOCIATIONERROR = "FailedOperation.CKafkaDiffAssociationError"
//  FAILEDOPERATION_CKAFKAEMPTYTOPICERROR = "FailedOperation.CKafkaEmptyTopicError"
//  FAILEDOPERATION_CKAFKAGETROUTEIDFAILEDERROR = "FailedOperation.CKafkaGetRouteIDFailedError"
//  FAILEDOPERATION_CKAFKAGETROUTETIMEOUTERROR = "FailedOperation.CKafkaGetRouteTimeoutError"
//  FAILEDOPERATION_CKAFKANOTAVAILABLEERROR = "FailedOperation.CKafkaNotAvailableError"
//  FAILEDOPERATION_PEERIDNOTAVAILABLE = "FailedOperation.PeerIdNotAvailable"
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
//  FAILEDOPERATION_PROMINSTANCENOTAVAILABLEERROR = "FailedOperation.PromInstanceNotAvailableError"
//  FAILEDOPERATION_ROUTENOTAVAILABLEERROR = "FailedOperation.RouteNotAvailableError"
//  FAILEDOPERATION_TOPICNOTAVAILABLEERROR = "FailedOperation.TopicNotAvailableError"
func (c *Client) ModifyApmAssociation(request *ModifyApmAssociationRequest) (response *ModifyApmAssociationResponse, err error) {
    return c.ModifyApmAssociationWithContext(context.Background(), request)
}

// ModifyApmAssociation
// 用于修改apm业务系统与其他产品的关联关系（包括创建和删除）
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASSOCIATIONMODIFYREQUESTNOTVALIDERROR = "FailedOperation.AssociationModifyRequestNotValidError"
//  FAILEDOPERATION_CKAFKADIFFASSOCIATIONERROR = "FailedOperation.CKafkaDiffAssociationError"
//  FAILEDOPERATION_CKAFKAEMPTYTOPICERROR = "FailedOperation.CKafkaEmptyTopicError"
//  FAILEDOPERATION_CKAFKAGETROUTEIDFAILEDERROR = "FailedOperation.CKafkaGetRouteIDFailedError"
//  FAILEDOPERATION_CKAFKAGETROUTETIMEOUTERROR = "FailedOperation.CKafkaGetRouteTimeoutError"
//  FAILEDOPERATION_CKAFKANOTAVAILABLEERROR = "FailedOperation.CKafkaNotAvailableError"
//  FAILEDOPERATION_PEERIDNOTAVAILABLE = "FailedOperation.PeerIdNotAvailable"
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
//  FAILEDOPERATION_PROMINSTANCENOTAVAILABLEERROR = "FailedOperation.PromInstanceNotAvailableError"
//  FAILEDOPERATION_ROUTENOTAVAILABLEERROR = "FailedOperation.RouteNotAvailableError"
//  FAILEDOPERATION_TOPICNOTAVAILABLEERROR = "FailedOperation.TopicNotAvailableError"
func (c *Client) ModifyApmAssociationWithContext(ctx context.Context, request *ModifyApmAssociationRequest) (response *ModifyApmAssociationResponse, err error) {
    if request == nil {
        request = NewModifyApmAssociationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmAssociation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmAssociation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmAssociationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmInstanceRequest() (request *ModifyApmInstanceRequest) {
    request = &ModifyApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmInstance")
    
    
    return
}

func NewModifyApmInstanceResponse() (response *ModifyApmInstanceResponse) {
    response = &ModifyApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmInstance
// 修改APM业务系统接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTMODIFY = "FailedOperation.InstanceCannotModify"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApmInstance(request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
    return c.ModifyApmInstanceWithContext(context.Background(), request)
}

// ModifyApmInstance
// 修改APM业务系统接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTMODIFY = "FailedOperation.InstanceCannotModify"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApmInstanceWithContext(ctx context.Context, request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
    if request == nil {
        request = NewModifyApmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmPrometheusRuleRequest() (request *ModifyApmPrometheusRuleRequest) {
    request = &ModifyApmPrometheusRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmPrometheusRule")
    
    
    return
}

func NewModifyApmPrometheusRuleResponse() (response *ModifyApmPrometheusRuleResponse) {
    response = &ModifyApmPrometheusRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmPrometheusRule
// 用于修改apm业务系统与Prometheus实例的指标匹配规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) ModifyApmPrometheusRule(request *ModifyApmPrometheusRuleRequest) (response *ModifyApmPrometheusRuleResponse, err error) {
    return c.ModifyApmPrometheusRuleWithContext(context.Background(), request)
}

// ModifyApmPrometheusRule
// 用于修改apm业务系统与Prometheus实例的指标匹配规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) ModifyApmPrometheusRuleWithContext(ctx context.Context, request *ModifyApmPrometheusRuleRequest) (response *ModifyApmPrometheusRuleResponse, err error) {
    if request == nil {
        request = NewModifyApmPrometheusRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmPrometheusRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmPrometheusRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmPrometheusRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmSampleConfigRequest() (request *ModifyApmSampleConfigRequest) {
    request = &ModifyApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmSampleConfig")
    
    
    return
}

func NewModifyApmSampleConfigResponse() (response *ModifyApmSampleConfigResponse) {
    response = &ModifyApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmSampleConfig
// 修改采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) ModifyApmSampleConfig(request *ModifyApmSampleConfigRequest) (response *ModifyApmSampleConfigResponse, err error) {
    return c.ModifyApmSampleConfigWithContext(context.Background(), request)
}

// ModifyApmSampleConfig
// 修改采样配置接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) ModifyApmSampleConfigWithContext(ctx context.Context, request *ModifyApmSampleConfigRequest) (response *ModifyApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewModifyApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGeneralApmApplicationConfigRequest() (request *ModifyGeneralApmApplicationConfigRequest) {
    request = &ModifyGeneralApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyGeneralApmApplicationConfig")
    
    
    return
}

func NewModifyGeneralApmApplicationConfigResponse() (response *ModifyGeneralApmApplicationConfigResponse) {
    response = &ModifyGeneralApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGeneralApmApplicationConfig
// 对外开放的openApi，客户可以灵活的指定需要修改的字段，再加入需要修改的服务列表.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APMCREDENTIALNOTEXIST = "FailedOperation.ApmCredentialNotExist"
//  FAILEDOPERATION_DUPLICATESERVICE = "FailedOperation.DuplicateService"
//  FAILEDOPERATION_DUPLICATETAGFIELD = "FailedOperation.DuplicateTagField"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
//  FAILEDOPERATION_INVALIDTAGFIELD = "FailedOperation.InvalidTagField"
//  FAILEDOPERATION_INVALIDTOKEN = "FailedOperation.InvalidToken"
//  FAILEDOPERATION_SERVICELISTEXCEEDINGLIMITNUMBER = "FailedOperation.ServiceListExceedingLimitNumber"
//  FAILEDOPERATION_SERVICELISTNULL = "FailedOperation.ServiceListNull"
func (c *Client) ModifyGeneralApmApplicationConfig(request *ModifyGeneralApmApplicationConfigRequest) (response *ModifyGeneralApmApplicationConfigResponse, err error) {
    return c.ModifyGeneralApmApplicationConfigWithContext(context.Background(), request)
}

// ModifyGeneralApmApplicationConfig
// 对外开放的openApi，客户可以灵活的指定需要修改的字段，再加入需要修改的服务列表.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APMCREDENTIALNOTEXIST = "FailedOperation.ApmCredentialNotExist"
//  FAILEDOPERATION_DUPLICATESERVICE = "FailedOperation.DuplicateService"
//  FAILEDOPERATION_DUPLICATETAGFIELD = "FailedOperation.DuplicateTagField"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
//  FAILEDOPERATION_INVALIDTAGFIELD = "FailedOperation.InvalidTagField"
//  FAILEDOPERATION_INVALIDTOKEN = "FailedOperation.InvalidToken"
//  FAILEDOPERATION_SERVICELISTEXCEEDINGLIMITNUMBER = "FailedOperation.ServiceListExceedingLimitNumber"
//  FAILEDOPERATION_SERVICELISTNULL = "FailedOperation.ServiceListNull"
func (c *Client) ModifyGeneralApmApplicationConfigWithContext(ctx context.Context, request *ModifyGeneralApmApplicationConfigRequest) (response *ModifyGeneralApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewModifyGeneralApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyGeneralApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGeneralApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGeneralApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateApmInstanceRequest() (request *TerminateApmInstanceRequest) {
    request = &TerminateApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "TerminateApmInstance")
    
    
    return
}

func NewTerminateApmInstanceResponse() (response *TerminateApmInstanceResponse) {
    response = &TerminateApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateApmInstance
// 销毁 APM 业务系统
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTTERMINATE = "FailedOperation.InstanceCannotTerminate"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TerminateApmInstance(request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
    return c.TerminateApmInstanceWithContext(context.Background(), request)
}

// TerminateApmInstance
// 销毁 APM 业务系统
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTTERMINATE = "FailedOperation.InstanceCannotTerminate"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TerminateApmInstanceWithContext(ctx context.Context, request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateApmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "TerminateApmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateApmInstanceResponse()
    err = c.Send(request, response)
    return
}
