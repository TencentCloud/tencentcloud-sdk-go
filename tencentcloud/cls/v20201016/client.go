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

package v20201016

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-10-16"

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


func NewAddMachineGroupInfoRequest() (request *AddMachineGroupInfoRequest) {
    request = &AddMachineGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "AddMachineGroupInfo")
    
    
    return
}

func NewAddMachineGroupInfoResponse() (response *AddMachineGroupInfoResponse) {
    response = &AddMachineGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddMachineGroupInfo
// 用于添加机器组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
func (c *Client) AddMachineGroupInfo(request *AddMachineGroupInfoRequest) (response *AddMachineGroupInfoResponse, err error) {
    return c.AddMachineGroupInfoWithContext(context.Background(), request)
}

// AddMachineGroupInfo
// 用于添加机器组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
func (c *Client) AddMachineGroupInfoWithContext(ctx context.Context, request *AddMachineGroupInfoRequest) (response *AddMachineGroupInfoResponse, err error) {
    if request == nil {
        request = NewAddMachineGroupInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddMachineGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddMachineGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewApplyConfigToMachineGroupRequest() (request *ApplyConfigToMachineGroupRequest) {
    request = &ApplyConfigToMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ApplyConfigToMachineGroup")
    
    
    return
}

func NewApplyConfigToMachineGroupResponse() (response *ApplyConfigToMachineGroupResponse) {
    response = &ApplyConfigToMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyConfigToMachineGroup
// 应用采集配置到指定机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ApplyConfigToMachineGroup(request *ApplyConfigToMachineGroupRequest) (response *ApplyConfigToMachineGroupResponse, err error) {
    return c.ApplyConfigToMachineGroupWithContext(context.Background(), request)
}

// ApplyConfigToMachineGroup
// 应用采集配置到指定机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ApplyConfigToMachineGroupWithContext(ctx context.Context, request *ApplyConfigToMachineGroupRequest) (response *ApplyConfigToMachineGroupResponse, err error) {
    if request == nil {
        request = NewApplyConfigToMachineGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyConfigToMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyConfigToMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFunctionRequest() (request *CheckFunctionRequest) {
    request = &CheckFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CheckFunction")
    
    
    return
}

func NewCheckFunctionResponse() (response *CheckFunctionResponse) {
    response = &CheckFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckFunction
// 本接口用于数据加工DSL函数的语法校验。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CheckFunction(request *CheckFunctionRequest) (response *CheckFunctionResponse, err error) {
    return c.CheckFunctionWithContext(context.Background(), request)
}

// CheckFunction
// 本接口用于数据加工DSL函数的语法校验。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CheckFunctionWithContext(ctx context.Context, request *CheckFunctionRequest) (response *CheckFunctionResponse, err error) {
    if request == nil {
        request = NewCheckFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCheckRechargeKafkaServerRequest() (request *CheckRechargeKafkaServerRequest) {
    request = &CheckRechargeKafkaServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CheckRechargeKafkaServer")
    
    
    return
}

func NewCheckRechargeKafkaServerResponse() (response *CheckRechargeKafkaServerResponse) {
    response = &CheckRechargeKafkaServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckRechargeKafkaServer
// 本接口用于校验Kafka服务集群是否可以正常访问
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckRechargeKafkaServer(request *CheckRechargeKafkaServerRequest) (response *CheckRechargeKafkaServerResponse, err error) {
    return c.CheckRechargeKafkaServerWithContext(context.Background(), request)
}

// CheckRechargeKafkaServer
// 本接口用于校验Kafka服务集群是否可以正常访问
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckRechargeKafkaServerWithContext(ctx context.Context, request *CheckRechargeKafkaServerRequest) (response *CheckRechargeKafkaServerResponse, err error) {
    if request == nil {
        request = NewCheckRechargeKafkaServerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckRechargeKafkaServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckRechargeKafkaServerResponse()
    err = c.Send(request, response)
    return
}

func NewCloseKafkaConsumerRequest() (request *CloseKafkaConsumerRequest) {
    request = &CloseKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CloseKafkaConsumer")
    
    
    return
}

func NewCloseKafkaConsumerResponse() (response *CloseKafkaConsumerResponse) {
    response = &CloseKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseKafkaConsumer
// 关闭Kafka协议消费
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CloseKafkaConsumer(request *CloseKafkaConsumerRequest) (response *CloseKafkaConsumerResponse, err error) {
    return c.CloseKafkaConsumerWithContext(context.Background(), request)
}

// CloseKafkaConsumer
// 关闭Kafka协议消费
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CloseKafkaConsumerWithContext(ctx context.Context, request *CloseKafkaConsumerRequest) (response *CloseKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewCloseKafkaConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmRequest() (request *CreateAlarmRequest) {
    request = &CreateAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateAlarm")
    
    
    return
}

func NewCreateAlarmResponse() (response *CreateAlarmResponse) {
    response = &CreateAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarm
// 本接口用于创建告警策略。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateAlarm(request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
    return c.CreateAlarmWithContext(context.Background(), request)
}

// CreateAlarm
// 本接口用于创建告警策略。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateAlarmWithContext(ctx context.Context, request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
    if request == nil {
        request = NewCreateAlarmRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmNoticeRequest() (request *CreateAlarmNoticeRequest) {
    request = &CreateAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateAlarmNotice")
    
    
    return
}

func NewCreateAlarmNoticeResponse() (response *CreateAlarmNoticeResponse) {
    response = &CreateAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmNotice
// 该接口用于创建通知渠道组，提供两种配置模式，二选一：
//
// 1，简易模式，提供最基本的通知渠道功能。需填写如下参数：
//
// - Type
//
// - NoticeReceivers
//
// - WebCallbacks
//
// 
//
// 2，高级模式，在简易模式基础上，支持设定规则，为不同类型的告警分别设定通知渠道，并支持告警升级功能。需填写如下参数：
//
// - NoticeRules
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmNotice(request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    return c.CreateAlarmNoticeWithContext(context.Background(), request)
}

// CreateAlarmNotice
// 该接口用于创建通知渠道组，提供两种配置模式，二选一：
//
// 1，简易模式，提供最基本的通知渠道功能。需填写如下参数：
//
// - Type
//
// - NoticeReceivers
//
// - WebCallbacks
//
// 
//
// 2，高级模式，在简易模式基础上，支持设定规则，为不同类型的告警分别设定通知渠道，并支持告警升级功能。需填写如下参数：
//
// - NoticeRules
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmNoticeWithContext(ctx context.Context, request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewCreateAlarmNoticeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmShieldRequest() (request *CreateAlarmShieldRequest) {
    request = &CreateAlarmShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateAlarmShield")
    
    
    return
}

func NewCreateAlarmShieldResponse() (response *CreateAlarmShieldResponse) {
    response = &CreateAlarmShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmShield
// 该接口用于创建告警屏蔽规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmShield(request *CreateAlarmShieldRequest) (response *CreateAlarmShieldResponse, err error) {
    return c.CreateAlarmShieldWithContext(context.Background(), request)
}

// CreateAlarmShield
// 该接口用于创建告警屏蔽规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateAlarmShieldWithContext(ctx context.Context, request *CreateAlarmShieldRequest) (response *CreateAlarmShieldResponse, err error) {
    if request == nil {
        request = NewCreateAlarmShieldRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmShield require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmShieldResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudProductLogTaskRequest() (request *CreateCloudProductLogTaskRequest) {
    request = &CreateCloudProductLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateCloudProductLogTask")
    
    
    return
}

func NewCreateCloudProductLogTaskResponse() (response *CreateCloudProductLogTaskResponse) {
    response = &CreateCloudProductLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudProductLogTask
// 内部云产品接入使用相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudProductLogTask(request *CreateCloudProductLogTaskRequest) (response *CreateCloudProductLogTaskResponse, err error) {
    return c.CreateCloudProductLogTaskWithContext(context.Background(), request)
}

// CreateCloudProductLogTask
// 内部云产品接入使用相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudProductLogTaskWithContext(ctx context.Context, request *CreateCloudProductLogTaskRequest) (response *CreateCloudProductLogTaskResponse, err error) {
    if request == nil {
        request = NewCreateCloudProductLogTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudProductLogTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudProductLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
    request = &CreateConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConfig")
    
    
    return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
    response = &CreateConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfig
// 创建采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_CONFIG = "LimitExceeded.Config"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    return c.CreateConfigWithContext(context.Background(), request)
}

// CreateConfig
// 创建采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_CONFIG = "LimitExceeded.Config"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateConfigWithContext(ctx context.Context, request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
    if request == nil {
        request = NewCreateConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigExtraRequest() (request *CreateConfigExtraRequest) {
    request = &CreateConfigExtraRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConfigExtra")
    
    
    return
}

func NewCreateConfigExtraResponse() (response *CreateConfigExtraResponse) {
    response = &CreateConfigExtraResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigExtra
// 本接口用于创建特殊采集配置任务，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateConfigExtra(request *CreateConfigExtraRequest) (response *CreateConfigExtraResponse, err error) {
    return c.CreateConfigExtraWithContext(context.Background(), request)
}

// CreateConfigExtra
// 本接口用于创建特殊采集配置任务，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateConfigExtraWithContext(ctx context.Context, request *CreateConfigExtraRequest) (response *CreateConfigExtraResponse, err error) {
    if request == nil {
        request = NewCreateConfigExtraRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigExtra require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigExtraResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsoleSharingRequest() (request *CreateConsoleSharingRequest) {
    request = &CreateConsoleSharingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConsoleSharing")
    
    
    return
}

func NewCreateConsoleSharingResponse() (response *CreateConsoleSharingResponse) {
    response = &CreateConsoleSharingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsoleSharing
// 创建控制台分享
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateConsoleSharing(request *CreateConsoleSharingRequest) (response *CreateConsoleSharingResponse, err error) {
    return c.CreateConsoleSharingWithContext(context.Background(), request)
}

// CreateConsoleSharing
// 创建控制台分享
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateConsoleSharingWithContext(ctx context.Context, request *CreateConsoleSharingRequest) (response *CreateConsoleSharingResponse, err error) {
    if request == nil {
        request = NewCreateConsoleSharingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsoleSharing require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsoleSharingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerRequest() (request *CreateConsumerRequest) {
    request = &CreateConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateConsumer")
    
    
    return
}

func NewCreateConsumerResponse() (response *CreateConsumerResponse) {
    response = &CreateConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsumer
// 本接口用于创建投递CKafka任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsumer(request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    return c.CreateConsumerWithContext(context.Background(), request)
}

// CreateConsumer
// 本接口用于创建投递CKafka任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateConsumerWithContext(ctx context.Context, request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    if request == nil {
        request = NewCreateConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosRechargeRequest() (request *CreateCosRechargeRequest) {
    request = &CreateCosRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateCosRecharge")
    
    
    return
}

func NewCreateCosRechargeResponse() (response *CreateCosRechargeResponse) {
    response = &CreateCosRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosRecharge
// 本接口用于创建cos导入任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateCosRecharge(request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
    return c.CreateCosRechargeWithContext(context.Background(), request)
}

// CreateCosRecharge
// 本接口用于创建cos导入任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateCosRechargeWithContext(ctx context.Context, request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
    if request == nil {
        request = NewCreateCosRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDashboardSubscribeRequest() (request *CreateDashboardSubscribeRequest) {
    request = &CreateDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDashboardSubscribe")
    
    
    return
}

func NewCreateDashboardSubscribeResponse() (response *CreateDashboardSubscribeResponse) {
    response = &CreateDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDashboardSubscribe
// 此接口用于创建仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) CreateDashboardSubscribe(request *CreateDashboardSubscribeRequest) (response *CreateDashboardSubscribeResponse, err error) {
    return c.CreateDashboardSubscribeWithContext(context.Background(), request)
}

// CreateDashboardSubscribe
// 此接口用于创建仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) CreateDashboardSubscribeWithContext(ctx context.Context, request *CreateDashboardSubscribeRequest) (response *CreateDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateDashboardSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataTransformRequest() (request *CreateDataTransformRequest) {
    request = &CreateDataTransformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDataTransform")
    
    
    return
}

func NewCreateDataTransformResponse() (response *CreateDataTransformResponse) {
    response = &CreateDataTransformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataTransform
// 本接口用于创建数据加工任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateDataTransform(request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
    return c.CreateDataTransformWithContext(context.Background(), request)
}

// CreateDataTransform
// 本接口用于创建数据加工任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateDataTransformWithContext(ctx context.Context, request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
    if request == nil {
        request = NewCreateDataTransformRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataTransformResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeliverCloudFunctionRequest() (request *CreateDeliverCloudFunctionRequest) {
    request = &CreateDeliverCloudFunctionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateDeliverCloudFunction")
    
    
    return
}

func NewCreateDeliverCloudFunctionResponse() (response *CreateDeliverCloudFunctionResponse) {
    response = &CreateDeliverCloudFunctionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeliverCloudFunction
// 本接口用于创建投递SCF任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDeliverCloudFunction(request *CreateDeliverCloudFunctionRequest) (response *CreateDeliverCloudFunctionResponse, err error) {
    return c.CreateDeliverCloudFunctionWithContext(context.Background(), request)
}

// CreateDeliverCloudFunction
// 本接口用于创建投递SCF任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDeliverCloudFunctionWithContext(ctx context.Context, request *CreateDeliverCloudFunctionRequest) (response *CreateDeliverCloudFunctionResponse, err error) {
    if request == nil {
        request = NewCreateDeliverCloudFunctionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeliverCloudFunction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeliverCloudFunctionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExportRequest() (request *CreateExportRequest) {
    request = &CreateExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateExport")
    
    
    return
}

func NewCreateExportResponse() (response *CreateExportResponse) {
    response = &CreateExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExport
// 本接口仅创建下载任务，任务返回的下载地址，请用户调用DescribeExports查看任务列表。其中有下载地址CosPath参数。参考文档https://cloud.tencent.com/document/product/614/56449
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EXPORT = "LimitExceeded.Export"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateExport(request *CreateExportRequest) (response *CreateExportResponse, err error) {
    return c.CreateExportWithContext(context.Background(), request)
}

// CreateExport
// 本接口仅创建下载任务，任务返回的下载地址，请用户调用DescribeExports查看任务列表。其中有下载地址CosPath参数。参考文档https://cloud.tencent.com/document/product/614/56449
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EXPORT = "LimitExceeded.Export"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateExportWithContext(ctx context.Context, request *CreateExportRequest) (response *CreateExportResponse, err error) {
    if request == nil {
        request = NewCreateExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIndexRequest() (request *CreateIndexRequest) {
    request = &CreateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateIndex")
    
    
    return
}

func NewCreateIndexResponse() (response *CreateIndexResponse) {
    response = &CreateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIndex
// 本接口用于创建索引
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndex(request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    return c.CreateIndexWithContext(context.Background(), request)
}

// CreateIndex
// 本接口用于创建索引
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndexWithContext(ctx context.Context, request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    if request == nil {
        request = NewCreateIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKafkaRechargeRequest() (request *CreateKafkaRechargeRequest) {
    request = &CreateKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateKafkaRecharge")
    
    
    return
}

func NewCreateKafkaRechargeResponse() (response *CreateKafkaRechargeResponse) {
    response = &CreateKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKafkaRecharge
// 本接口用于创建Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateKafkaRecharge(request *CreateKafkaRechargeRequest) (response *CreateKafkaRechargeResponse, err error) {
    return c.CreateKafkaRechargeWithContext(context.Background(), request)
}

// CreateKafkaRecharge
// 本接口用于创建Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateKafkaRechargeWithContext(ctx context.Context, request *CreateKafkaRechargeRequest) (response *CreateKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewCreateKafkaRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogsetRequest() (request *CreateLogsetRequest) {
    request = &CreateLogsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateLogset")
    
    
    return
}

func NewCreateLogsetResponse() (response *CreateLogsetResponse) {
    response = &CreateLogsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLogset
// 本接口用于创建日志集，返回新创建的日志集的 ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETCONFLICT = "FailedOperation.LogsetConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateLogset(request *CreateLogsetRequest) (response *CreateLogsetResponse, err error) {
    return c.CreateLogsetWithContext(context.Background(), request)
}

// CreateLogset
// 本接口用于创建日志集，返回新创建的日志集的 ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETCONFLICT = "FailedOperation.LogsetConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateLogsetWithContext(ctx context.Context, request *CreateLogsetRequest) (response *CreateLogsetResponse, err error) {
    if request == nil {
        request = NewCreateLogsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogsetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMachineGroupRequest() (request *CreateMachineGroupRequest) {
    request = &CreateMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateMachineGroup")
    
    
    return
}

func NewCreateMachineGroupResponse() (response *CreateMachineGroupResponse) {
    response = &CreateMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMachineGroup
// 创建机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MACHINEGROUP = "LimitExceeded.MachineGroup"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateMachineGroup(request *CreateMachineGroupRequest) (response *CreateMachineGroupResponse, err error) {
    return c.CreateMachineGroupWithContext(context.Background(), request)
}

// CreateMachineGroup
// 创建机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_MACHINEGROUP = "LimitExceeded.MachineGroup"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateMachineGroupWithContext(ctx context.Context, request *CreateMachineGroupRequest) (response *CreateMachineGroupResponse, err error) {
    if request == nil {
        request = NewCreateMachineGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNoticeContentRequest() (request *CreateNoticeContentRequest) {
    request = &CreateNoticeContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateNoticeContent")
    
    
    return
}

func NewCreateNoticeContentResponse() (response *CreateNoticeContentResponse) {
    response = &CreateNoticeContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNoticeContent
// 该接口用于创建通知内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateNoticeContent(request *CreateNoticeContentRequest) (response *CreateNoticeContentResponse, err error) {
    return c.CreateNoticeContentWithContext(context.Background(), request)
}

// CreateNoticeContent
// 该接口用于创建通知内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) CreateNoticeContentWithContext(ctx context.Context, request *CreateNoticeContentRequest) (response *CreateNoticeContentResponse, err error) {
    if request == nil {
        request = NewCreateNoticeContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNoticeContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNoticeContentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScheduledSqlRequest() (request *CreateScheduledSqlRequest) {
    request = &CreateScheduledSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateScheduledSql")
    
    
    return
}

func NewCreateScheduledSqlResponse() (response *CreateScheduledSqlResponse) {
    response = &CreateScheduledSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScheduledSql
// 本接口用于创建定时SQL分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateScheduledSql(request *CreateScheduledSqlRequest) (response *CreateScheduledSqlResponse, err error) {
    return c.CreateScheduledSqlWithContext(context.Background(), request)
}

// CreateScheduledSql
// 本接口用于创建定时SQL分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateScheduledSqlWithContext(ctx context.Context, request *CreateScheduledSqlRequest) (response *CreateScheduledSqlResponse, err error) {
    if request == nil {
        request = NewCreateScheduledSqlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScheduledSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScheduledSqlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateShipperRequest() (request *CreateShipperRequest) {
    request = &CreateShipperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateShipper")
    
    
    return
}

func NewCreateShipperResponse() (response *CreateShipperResponse) {
    response = &CreateShipperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateShipper
// 新建投递到COS的任务，【！！！注意】使用此接口，需要检查是否配置了投递COS的角色和权限。如果没有配置，请参考文档投递权限查看和配置https://cloud.tencent.com/document/product/614/71623。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SHIPPERCONFLICT = "InvalidParameter.ShipperConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SHIPPER = "LimitExceeded.Shipper"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateShipper(request *CreateShipperRequest) (response *CreateShipperResponse, err error) {
    return c.CreateShipperWithContext(context.Background(), request)
}

// CreateShipper
// 新建投递到COS的任务，【！！！注意】使用此接口，需要检查是否配置了投递COS的角色和权限。如果没有配置，请参考文档投递权限查看和配置https://cloud.tencent.com/document/product/614/71623。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SHIPPERCONFLICT = "InvalidParameter.ShipperConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_SHIPPER = "LimitExceeded.Shipper"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) CreateShipperWithContext(ctx context.Context, request *CreateShipperRequest) (response *CreateShipperResponse, err error) {
    if request == nil {
        request = NewCreateShipperRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateShipperResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "CreateTopic")
    
    
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTopic
// 本接口用于创建日志主题。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TOPICCREATING = "FailedOperation.TopicCreating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TOPIC = "LimitExceeded.Topic"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// 本接口用于创建日志主题。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TOPICCREATING = "FailedOperation.TopicCreating"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"
//  LIMITEXCEEDED_TOPIC = "LimitExceeded.Topic"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmRequest() (request *DeleteAlarmRequest) {
    request = &DeleteAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarm")
    
    
    return
}

func NewDeleteAlarmResponse() (response *DeleteAlarmResponse) {
    response = &DeleteAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarm
// 本接口用于删除告警策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
func (c *Client) DeleteAlarm(request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
    return c.DeleteAlarmWithContext(context.Background(), request)
}

// DeleteAlarm
// 本接口用于删除告警策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
func (c *Client) DeleteAlarmWithContext(ctx context.Context, request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmNoticeRequest() (request *DeleteAlarmNoticeRequest) {
    request = &DeleteAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarmNotice")
    
    
    return
}

func NewDeleteAlarmNoticeResponse() (response *DeleteAlarmNoticeResponse) {
    response = &DeleteAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmNotice
// 该接口用于删除通知渠道组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteAlarmNotice(request *DeleteAlarmNoticeRequest) (response *DeleteAlarmNoticeResponse, err error) {
    return c.DeleteAlarmNoticeWithContext(context.Background(), request)
}

// DeleteAlarmNotice
// 该接口用于删除通知渠道组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteAlarmNoticeWithContext(ctx context.Context, request *DeleteAlarmNoticeRequest) (response *DeleteAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmNoticeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmShieldRequest() (request *DeleteAlarmShieldRequest) {
    request = &DeleteAlarmShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarmShield")
    
    
    return
}

func NewDeleteAlarmShieldResponse() (response *DeleteAlarmShieldResponse) {
    response = &DeleteAlarmShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmShield
// 该接口用于删除告警屏蔽规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteAlarmShield(request *DeleteAlarmShieldRequest) (response *DeleteAlarmShieldResponse, err error) {
    return c.DeleteAlarmShieldWithContext(context.Background(), request)
}

// DeleteAlarmShield
// 该接口用于删除告警屏蔽规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteAlarmShieldWithContext(ctx context.Context, request *DeleteAlarmShieldRequest) (response *DeleteAlarmShieldResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmShieldRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmShield require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmShieldResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudProductLogTaskRequest() (request *DeleteCloudProductLogTaskRequest) {
    request = &DeleteCloudProductLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteCloudProductLogTask")
    
    
    return
}

func NewDeleteCloudProductLogTaskResponse() (response *DeleteCloudProductLogTaskResponse) {
    response = &DeleteCloudProductLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudProductLogTask
// 内部云产品接入使用相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCloudProductLogTask(request *DeleteCloudProductLogTaskRequest) (response *DeleteCloudProductLogTaskResponse, err error) {
    return c.DeleteCloudProductLogTaskWithContext(context.Background(), request)
}

// DeleteCloudProductLogTask
// 内部云产品接入使用相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCloudProductLogTaskWithContext(ctx context.Context, request *DeleteCloudProductLogTaskRequest) (response *DeleteCloudProductLogTaskResponse, err error) {
    if request == nil {
        request = NewDeleteCloudProductLogTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudProductLogTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudProductLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
    request = &DeleteConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConfig")
    
    
    return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
    response = &DeleteConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfig
// 删除采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    return c.DeleteConfigWithContext(context.Background(), request)
}

// DeleteConfig
// 删除采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DeleteConfigWithContext(ctx context.Context, request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
    if request == nil {
        request = NewDeleteConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigExtraRequest() (request *DeleteConfigExtraRequest) {
    request = &DeleteConfigExtraRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigExtra")
    
    
    return
}

func NewDeleteConfigExtraResponse() (response *DeleteConfigExtraResponse) {
    response = &DeleteConfigExtraResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigExtra
// 本接口用于删除特殊采集规则配置，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConfigExtra(request *DeleteConfigExtraRequest) (response *DeleteConfigExtraResponse, err error) {
    return c.DeleteConfigExtraWithContext(context.Background(), request)
}

// DeleteConfigExtra
// 本接口用于删除特殊采集规则配置，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConfigExtraWithContext(ctx context.Context, request *DeleteConfigExtraRequest) (response *DeleteConfigExtraResponse, err error) {
    if request == nil {
        request = NewDeleteConfigExtraRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigExtra require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigExtraResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConfigFromMachineGroupRequest() (request *DeleteConfigFromMachineGroupRequest) {
    request = &DeleteConfigFromMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigFromMachineGroup")
    
    
    return
}

func NewDeleteConfigFromMachineGroupResponse() (response *DeleteConfigFromMachineGroupResponse) {
    response = &DeleteConfigFromMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConfigFromMachineGroup
// 删除应用到机器组的采集配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteConfigFromMachineGroup(request *DeleteConfigFromMachineGroupRequest) (response *DeleteConfigFromMachineGroupResponse, err error) {
    return c.DeleteConfigFromMachineGroupWithContext(context.Background(), request)
}

// DeleteConfigFromMachineGroup
// 删除应用到机器组的采集配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteConfigFromMachineGroupWithContext(ctx context.Context, request *DeleteConfigFromMachineGroupRequest) (response *DeleteConfigFromMachineGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConfigFromMachineGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConfigFromMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConfigFromMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsoleSharingRequest() (request *DeleteConsoleSharingRequest) {
    request = &DeleteConsoleSharingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConsoleSharing")
    
    
    return
}

func NewDeleteConsoleSharingResponse() (response *DeleteConsoleSharingResponse) {
    response = &DeleteConsoleSharingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsoleSharing
// 删除控制台分享
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConsoleSharing(request *DeleteConsoleSharingRequest) (response *DeleteConsoleSharingResponse, err error) {
    return c.DeleteConsoleSharingWithContext(context.Background(), request)
}

// DeleteConsoleSharing
// 删除控制台分享
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteConsoleSharingWithContext(ctx context.Context, request *DeleteConsoleSharingRequest) (response *DeleteConsoleSharingResponse, err error) {
    if request == nil {
        request = NewDeleteConsoleSharingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsoleSharing require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsoleSharingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsumerRequest() (request *DeleteConsumerRequest) {
    request = &DeleteConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteConsumer")
    
    
    return
}

func NewDeleteConsumerResponse() (response *DeleteConsumerResponse) {
    response = &DeleteConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsumer
// 本接口用于删除投递配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConsumer(request *DeleteConsumerRequest) (response *DeleteConsumerResponse, err error) {
    return c.DeleteConsumerWithContext(context.Background(), request)
}

// DeleteConsumer
// 本接口用于删除投递配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DeleteConsumerWithContext(ctx context.Context, request *DeleteConsumerRequest) (response *DeleteConsumerResponse, err error) {
    if request == nil {
        request = NewDeleteConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDashboardSubscribeRequest() (request *DeleteDashboardSubscribeRequest) {
    request = &DeleteDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteDashboardSubscribe")
    
    
    return
}

func NewDeleteDashboardSubscribeResponse() (response *DeleteDashboardSubscribeResponse) {
    response = &DeleteDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDashboardSubscribe
// 此接口用于删除仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) DeleteDashboardSubscribe(request *DeleteDashboardSubscribeRequest) (response *DeleteDashboardSubscribeResponse, err error) {
    return c.DeleteDashboardSubscribeWithContext(context.Background(), request)
}

// DeleteDashboardSubscribe
// 此接口用于删除仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) DeleteDashboardSubscribeWithContext(ctx context.Context, request *DeleteDashboardSubscribeRequest) (response *DeleteDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewDeleteDashboardSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataTransformRequest() (request *DeleteDataTransformRequest) {
    request = &DeleteDataTransformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteDataTransform")
    
    
    return
}

func NewDeleteDataTransformResponse() (response *DeleteDataTransformResponse) {
    response = &DeleteDataTransformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataTransform
// 本接口用于删除数据加工任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteDataTransform(request *DeleteDataTransformRequest) (response *DeleteDataTransformResponse, err error) {
    return c.DeleteDataTransformWithContext(context.Background(), request)
}

// DeleteDataTransform
// 本接口用于删除数据加工任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteDataTransformWithContext(ctx context.Context, request *DeleteDataTransformRequest) (response *DeleteDataTransformResponse, err error) {
    if request == nil {
        request = NewDeleteDataTransformRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataTransformResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteExportRequest() (request *DeleteExportRequest) {
    request = &DeleteExportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteExport")
    
    
    return
}

func NewDeleteExportResponse() (response *DeleteExportResponse) {
    response = &DeleteExportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteExport
// 本接口用于删除日志下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
func (c *Client) DeleteExport(request *DeleteExportRequest) (response *DeleteExportResponse, err error) {
    return c.DeleteExportWithContext(context.Background(), request)
}

// DeleteExport
// 本接口用于删除日志下载任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
func (c *Client) DeleteExportWithContext(ctx context.Context, request *DeleteExportRequest) (response *DeleteExportResponse, err error) {
    if request == nil {
        request = NewDeleteExportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExportResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIndexRequest() (request *DeleteIndexRequest) {
    request = &DeleteIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteIndex")
    
    
    return
}

func NewDeleteIndexResponse() (response *DeleteIndexResponse) {
    response = &DeleteIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIndex
// 本接口用于删除日志主题的索引配置，删除索引配置后将无法检索和查询采集到的日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    return c.DeleteIndexWithContext(context.Background(), request)
}

// DeleteIndex
// 本接口用于删除日志主题的索引配置，删除索引配置后将无法检索和查询采集到的日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteIndexWithContext(ctx context.Context, request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    if request == nil {
        request = NewDeleteIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKafkaRechargeRequest() (request *DeleteKafkaRechargeRequest) {
    request = &DeleteKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteKafkaRecharge")
    
    
    return
}

func NewDeleteKafkaRechargeResponse() (response *DeleteKafkaRechargeResponse) {
    response = &DeleteKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKafkaRecharge
// 本接口用于删除Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteKafkaRecharge(request *DeleteKafkaRechargeRequest) (response *DeleteKafkaRechargeResponse, err error) {
    return c.DeleteKafkaRechargeWithContext(context.Background(), request)
}

// DeleteKafkaRecharge
// 本接口用于删除Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteKafkaRechargeWithContext(ctx context.Context, request *DeleteKafkaRechargeRequest) (response *DeleteKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewDeleteKafkaRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogsetRequest() (request *DeleteLogsetRequest) {
    request = &DeleteLogsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteLogset")
    
    
    return
}

func NewDeleteLogsetResponse() (response *DeleteLogsetResponse) {
    response = &DeleteLogsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogset
// 本接口用于删除日志集。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETNOTEMPTY = "FailedOperation.LogsetNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) DeleteLogset(request *DeleteLogsetRequest) (response *DeleteLogsetResponse, err error) {
    return c.DeleteLogsetWithContext(context.Background(), request)
}

// DeleteLogset
// 本接口用于删除日志集。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGSETNOTEMPTY = "FailedOperation.LogsetNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) DeleteLogsetWithContext(ctx context.Context, request *DeleteLogsetRequest) (response *DeleteLogsetResponse, err error) {
    if request == nil {
        request = NewDeleteLogsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogsetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineGroupRequest() (request *DeleteMachineGroupRequest) {
    request = &DeleteMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroup")
    
    
    return
}

func NewDeleteMachineGroupResponse() (response *DeleteMachineGroupResponse) {
    response = &DeleteMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMachineGroup
// 删除机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteMachineGroup(request *DeleteMachineGroupRequest) (response *DeleteMachineGroupResponse, err error) {
    return c.DeleteMachineGroupWithContext(context.Background(), request)
}

// DeleteMachineGroup
// 删除机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DeleteMachineGroupWithContext(ctx context.Context, request *DeleteMachineGroupRequest) (response *DeleteMachineGroupResponse, err error) {
    if request == nil {
        request = NewDeleteMachineGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMachineGroupInfoRequest() (request *DeleteMachineGroupInfoRequest) {
    request = &DeleteMachineGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroupInfo")
    
    
    return
}

func NewDeleteMachineGroupInfoResponse() (response *DeleteMachineGroupInfoResponse) {
    response = &DeleteMachineGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMachineGroupInfo
// 用于删除机器组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteMachineGroupInfo(request *DeleteMachineGroupInfoRequest) (response *DeleteMachineGroupInfoResponse, err error) {
    return c.DeleteMachineGroupInfoWithContext(context.Background(), request)
}

// DeleteMachineGroupInfo
// 用于删除机器组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteMachineGroupInfoWithContext(ctx context.Context, request *DeleteMachineGroupInfoRequest) (response *DeleteMachineGroupInfoResponse, err error) {
    if request == nil {
        request = NewDeleteMachineGroupInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMachineGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMachineGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNoticeContentRequest() (request *DeleteNoticeContentRequest) {
    request = &DeleteNoticeContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteNoticeContent")
    
    
    return
}

func NewDeleteNoticeContentResponse() (response *DeleteNoticeContentResponse) {
    response = &DeleteNoticeContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNoticeContent
// 该接口用于删除通知内容配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteNoticeContent(request *DeleteNoticeContentRequest) (response *DeleteNoticeContentResponse, err error) {
    return c.DeleteNoticeContentWithContext(context.Background(), request)
}

// DeleteNoticeContent
// 该接口用于删除通知内容配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) DeleteNoticeContentWithContext(ctx context.Context, request *DeleteNoticeContentRequest) (response *DeleteNoticeContentResponse, err error) {
    if request == nil {
        request = NewDeleteNoticeContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNoticeContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNoticeContentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScheduledSqlRequest() (request *DeleteScheduledSqlRequest) {
    request = &DeleteScheduledSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteScheduledSql")
    
    
    return
}

func NewDeleteScheduledSqlResponse() (response *DeleteScheduledSqlResponse) {
    response = &DeleteScheduledSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteScheduledSql
// 本接口用于删除定时SQL分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteScheduledSql(request *DeleteScheduledSqlRequest) (response *DeleteScheduledSqlResponse, err error) {
    return c.DeleteScheduledSqlWithContext(context.Background(), request)
}

// DeleteScheduledSql
// 本接口用于删除定时SQL分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteScheduledSqlWithContext(ctx context.Context, request *DeleteScheduledSqlRequest) (response *DeleteScheduledSqlResponse, err error) {
    if request == nil {
        request = NewDeleteScheduledSqlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteScheduledSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteScheduledSqlResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShipperRequest() (request *DeleteShipperRequest) {
    request = &DeleteShipperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteShipper")
    
    
    return
}

func NewDeleteShipperResponse() (response *DeleteShipperResponse) {
    response = &DeleteShipperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShipper
// 删除投递COS任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DeleteShipper(request *DeleteShipperRequest) (response *DeleteShipperResponse, err error) {
    return c.DeleteShipperWithContext(context.Background(), request)
}

// DeleteShipper
// 删除投递COS任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DeleteShipperWithContext(ctx context.Context, request *DeleteShipperRequest) (response *DeleteShipperResponse, err error) {
    if request == nil {
        request = NewDeleteShipperRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShipperResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopic
// 本接口用于删除日志主题。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASDATAFORMTASK = "OperationDenied.TopicHasDataFormTask"
//  OPERATIONDENIED_TOPICHASDELIVERFUNCTION = "OperationDenied.TopicHasDeliverFunction"
//  OPERATIONDENIED_TOPICHASEXTERNALDATASOURCECONFIG = "OperationDenied.TopicHasExternalDatasourceConfig"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// 本接口用于删除日志主题。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASDATAFORMTASK = "OperationDenied.TopicHasDataFormTask"
//  OPERATIONDENIED_TOPICHASDELIVERFUNCTION = "OperationDenied.TopicHasDeliverFunction"
//  OPERATIONDENIED_TOPICHASEXTERNALDATASOURCECONFIG = "OperationDenied.TopicHasExternalDatasourceConfig"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
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

func NewDescribeAlarmNoticesRequest() (request *DescribeAlarmNoticesRequest) {
    request = &DescribeAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmNotices")
    
    
    return
}

func NewDescribeAlarmNoticesResponse() (response *DescribeAlarmNoticesResponse) {
    response = &DescribeAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotices
// 该接口用于获取通知渠道组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlarmNotices(request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    return c.DescribeAlarmNoticesWithContext(context.Background(), request)
}

// DescribeAlarmNotices
// 该接口用于获取通知渠道组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlarmNoticesWithContext(ctx context.Context, request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmShieldsRequest() (request *DescribeAlarmShieldsRequest) {
    request = &DescribeAlarmShieldsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmShields")
    
    
    return
}

func NewDescribeAlarmShieldsResponse() (response *DescribeAlarmShieldsResponse) {
    response = &DescribeAlarmShieldsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmShields
// 获取告警屏蔽配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmShields(request *DescribeAlarmShieldsRequest) (response *DescribeAlarmShieldsResponse, err error) {
    return c.DescribeAlarmShieldsWithContext(context.Background(), request)
}

// DescribeAlarmShields
// 获取告警屏蔽配置规则
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmShieldsWithContext(ctx context.Context, request *DescribeAlarmShieldsRequest) (response *DescribeAlarmShieldsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmShieldsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmShields require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmShieldsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
    request = &DescribeAlarmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarms")
    
    
    return
}

func NewDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
    response = &DescribeAlarmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarms
// 本接口用于获取告警策略列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    return c.DescribeAlarmsWithContext(context.Background(), request)
}

// DescribeAlarms
// 本接口用于获取告警策略列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeAlarmsWithContext(ctx context.Context, request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertRecordHistoryRequest() (request *DescribeAlertRecordHistoryRequest) {
    request = &DescribeAlertRecordHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeAlertRecordHistory")
    
    
    return
}

func NewDescribeAlertRecordHistoryResponse() (response *DescribeAlertRecordHistoryResponse) {
    response = &DescribeAlertRecordHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertRecordHistory
// 获取告警历史，例如今天未恢复的告警
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlertRecordHistory(request *DescribeAlertRecordHistoryRequest) (response *DescribeAlertRecordHistoryResponse, err error) {
    return c.DescribeAlertRecordHistoryWithContext(context.Background(), request)
}

// DescribeAlertRecordHistory
// 获取告警历史，例如今天未恢复的告警
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeAlertRecordHistoryWithContext(ctx context.Context, request *DescribeAlertRecordHistoryRequest) (response *DescribeAlertRecordHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeAlertRecordHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertRecordHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertRecordHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudProductLogTasksRequest() (request *DescribeCloudProductLogTasksRequest) {
    request = &DescribeCloudProductLogTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeCloudProductLogTasks")
    
    
    return
}

func NewDescribeCloudProductLogTasksResponse() (response *DescribeCloudProductLogTasksResponse) {
    response = &DescribeCloudProductLogTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudProductLogTasks
// 云产品接入使用相关接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCloudProductLogTasks(request *DescribeCloudProductLogTasksRequest) (response *DescribeCloudProductLogTasksResponse, err error) {
    return c.DescribeCloudProductLogTasksWithContext(context.Background(), request)
}

// DescribeCloudProductLogTasks
// 云产品接入使用相关接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCloudProductLogTasksWithContext(ctx context.Context, request *DescribeCloudProductLogTasksRequest) (response *DescribeCloudProductLogTasksResponse, err error) {
    if request == nil {
        request = NewDescribeCloudProductLogTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudProductLogTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudProductLogTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigExtrasRequest() (request *DescribeConfigExtrasRequest) {
    request = &DescribeConfigExtrasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigExtras")
    
    
    return
}

func NewDescribeConfigExtrasResponse() (response *DescribeConfigExtrasResponse) {
    response = &DescribeConfigExtrasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigExtras
// 本接口用于获取特殊采集配置，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConfigExtras(request *DescribeConfigExtrasRequest) (response *DescribeConfigExtrasResponse, err error) {
    return c.DescribeConfigExtrasWithContext(context.Background(), request)
}

// DescribeConfigExtras
// 本接口用于获取特殊采集配置，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConfigExtrasWithContext(ctx context.Context, request *DescribeConfigExtrasRequest) (response *DescribeConfigExtrasResponse, err error) {
    if request == nil {
        request = NewDescribeConfigExtrasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigExtras require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigExtrasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigMachineGroupsRequest() (request *DescribeConfigMachineGroupsRequest) {
    request = &DescribeConfigMachineGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigMachineGroups")
    
    
    return
}

func NewDescribeConfigMachineGroupsResponse() (response *DescribeConfigMachineGroupsResponse) {
    response = &DescribeConfigMachineGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigMachineGroups
// 获取采集规则配置所绑定的机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DescribeConfigMachineGroups(request *DescribeConfigMachineGroupsRequest) (response *DescribeConfigMachineGroupsResponse, err error) {
    return c.DescribeConfigMachineGroupsWithContext(context.Background(), request)
}

// DescribeConfigMachineGroups
// 获取采集规则配置所绑定的机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
func (c *Client) DescribeConfigMachineGroupsWithContext(ctx context.Context, request *DescribeConfigMachineGroupsRequest) (response *DescribeConfigMachineGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigMachineGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigMachineGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigMachineGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
    request = &DescribeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigs")
    
    
    return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
    response = &DescribeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigs
// 获取采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    return c.DescribeConfigsWithContext(context.Background(), request)
}

// DescribeConfigs
// 获取采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConfigsWithContext(ctx context.Context, request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsoleSharingListRequest() (request *DescribeConsoleSharingListRequest) {
    request = &DescribeConsoleSharingListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsoleSharingList")
    
    
    return
}

func NewDescribeConsoleSharingListResponse() (response *DescribeConsoleSharingListResponse) {
    response = &DescribeConsoleSharingListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsoleSharingList
// 批量查询控制台分享列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConsoleSharingList(request *DescribeConsoleSharingListRequest) (response *DescribeConsoleSharingListResponse, err error) {
    return c.DescribeConsoleSharingListWithContext(context.Background(), request)
}

// DescribeConsoleSharingList
// 批量查询控制台分享列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConsoleSharingListWithContext(ctx context.Context, request *DescribeConsoleSharingListRequest) (response *DescribeConsoleSharingListResponse, err error) {
    if request == nil {
        request = NewDescribeConsoleSharingListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsoleSharingList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsoleSharingListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerRequest() (request *DescribeConsumerRequest) {
    request = &DescribeConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumer")
    
    
    return
}

func NewDescribeConsumerResponse() (response *DescribeConsumerResponse) {
    response = &DescribeConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumer
// 本接口用于获取投递配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConsumer(request *DescribeConsumerRequest) (response *DescribeConsumerResponse, err error) {
    return c.DescribeConsumerWithContext(context.Background(), request)
}

// DescribeConsumer
// 本接口用于获取投递配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeConsumerWithContext(ctx context.Context, request *DescribeConsumerRequest) (response *DescribeConsumerResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosRechargesRequest() (request *DescribeCosRechargesRequest) {
    request = &DescribeCosRechargesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeCosRecharges")
    
    
    return
}

func NewDescribeCosRechargesResponse() (response *DescribeCosRechargesResponse) {
    response = &DescribeCosRechargesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosRecharges
// 本接口用于获取cos导入配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeCosRecharges(request *DescribeCosRechargesRequest) (response *DescribeCosRechargesResponse, err error) {
    return c.DescribeCosRechargesWithContext(context.Background(), request)
}

// DescribeCosRecharges
// 本接口用于获取cos导入配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeCosRechargesWithContext(ctx context.Context, request *DescribeCosRechargesRequest) (response *DescribeCosRechargesResponse, err error) {
    if request == nil {
        request = NewDescribeCosRechargesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosRecharges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosRechargesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardSubscribesRequest() (request *DescribeDashboardSubscribesRequest) {
    request = &DescribeDashboardSubscribesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeDashboardSubscribes")
    
    
    return
}

func NewDescribeDashboardSubscribesResponse() (response *DescribeDashboardSubscribesResponse) {
    response = &DescribeDashboardSubscribesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDashboardSubscribes
// 本接口用于获取仪表盘订阅列表，支持分页
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboardSubscribes(request *DescribeDashboardSubscribesRequest) (response *DescribeDashboardSubscribesResponse, err error) {
    return c.DescribeDashboardSubscribesWithContext(context.Background(), request)
}

// DescribeDashboardSubscribes
// 本接口用于获取仪表盘订阅列表，支持分页
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDashboardSubscribesWithContext(ctx context.Context, request *DescribeDashboardSubscribesRequest) (response *DescribeDashboardSubscribesResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardSubscribesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDashboardSubscribes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDashboardSubscribesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardsRequest() (request *DescribeDashboardsRequest) {
    request = &DescribeDashboardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeDashboards")
    
    
    return
}

func NewDescribeDashboardsResponse() (response *DescribeDashboardsResponse) {
    response = &DescribeDashboardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDashboards
// 本接口用于获取仪表盘
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeDashboards(request *DescribeDashboardsRequest) (response *DescribeDashboardsResponse, err error) {
    return c.DescribeDashboardsWithContext(context.Background(), request)
}

// DescribeDashboards
// 本接口用于获取仪表盘
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeDashboardsWithContext(ctx context.Context, request *DescribeDashboardsRequest) (response *DescribeDashboardsResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDashboards require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDashboardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataTransformInfoRequest() (request *DescribeDataTransformInfoRequest) {
    request = &DescribeDataTransformInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformInfo")
    
    
    return
}

func NewDescribeDataTransformInfoResponse() (response *DescribeDataTransformInfoResponse) {
    response = &DescribeDataTransformInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataTransformInfo
// 本接口用于获取数据加工任务列表基本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeDataTransformInfo(request *DescribeDataTransformInfoRequest) (response *DescribeDataTransformInfoResponse, err error) {
    return c.DescribeDataTransformInfoWithContext(context.Background(), request)
}

// DescribeDataTransformInfo
// 本接口用于获取数据加工任务列表基本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeDataTransformInfoWithContext(ctx context.Context, request *DescribeDataTransformInfoRequest) (response *DescribeDataTransformInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDataTransformInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataTransformInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataTransformInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportsRequest() (request *DescribeExportsRequest) {
    request = &DescribeExportsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeExports")
    
    
    return
}

func NewDescribeExportsResponse() (response *DescribeExportsResponse) {
    response = &DescribeExportsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExports
// 本接口用于获取日志下载任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExports(request *DescribeExportsRequest) (response *DescribeExportsResponse, err error) {
    return c.DescribeExportsWithContext(context.Background(), request)
}

// DescribeExports
// 本接口用于获取日志下载任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeExportsWithContext(ctx context.Context, request *DescribeExportsRequest) (response *DescribeExportsResponse, err error) {
    if request == nil {
        request = NewDescribeExportsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExports require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexRequest() (request *DescribeIndexRequest) {
    request = &DescribeIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeIndex")
    
    
    return
}

func NewDescribeIndexResponse() (response *DescribeIndexResponse) {
    response = &DescribeIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIndex
// 本接口用于获取索引配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeIndex(request *DescribeIndexRequest) (response *DescribeIndexResponse, err error) {
    return c.DescribeIndexWithContext(context.Background(), request)
}

// DescribeIndex
// 本接口用于获取索引配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeIndexWithContext(ctx context.Context, request *DescribeIndexRequest) (response *DescribeIndexResponse, err error) {
    if request == nil {
        request = NewDescribeIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaConsumerRequest() (request *DescribeKafkaConsumerRequest) {
    request = &DescribeKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumer")
    
    
    return
}

func NewDescribeKafkaConsumerResponse() (response *DescribeKafkaConsumerResponse) {
    response = &DescribeKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaConsumer
// 获取Kafka协议消费信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumer(request *DescribeKafkaConsumerRequest) (response *DescribeKafkaConsumerResponse, err error) {
    return c.DescribeKafkaConsumerWithContext(context.Background(), request)
}

// DescribeKafkaConsumer
// 获取Kafka协议消费信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaConsumerWithContext(ctx context.Context, request *DescribeKafkaConsumerRequest) (response *DescribeKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKafkaRechargesRequest() (request *DescribeKafkaRechargesRequest) {
    request = &DescribeKafkaRechargesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaRecharges")
    
    
    return
}

func NewDescribeKafkaRechargesResponse() (response *DescribeKafkaRechargesResponse) {
    response = &DescribeKafkaRechargesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKafkaRecharges
// 本接口用于获取Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaRecharges(request *DescribeKafkaRechargesRequest) (response *DescribeKafkaRechargesResponse, err error) {
    return c.DescribeKafkaRechargesWithContext(context.Background(), request)
}

// DescribeKafkaRecharges
// 本接口用于获取Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeKafkaRechargesWithContext(ctx context.Context, request *DescribeKafkaRechargesRequest) (response *DescribeKafkaRechargesResponse, err error) {
    if request == nil {
        request = NewDescribeKafkaRechargesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKafkaRecharges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKafkaRechargesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogContextRequest() (request *DescribeLogContextRequest) {
    request = &DescribeLogContextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeLogContext")
    
    
    return
}

func NewDescribeLogContextResponse() (response *DescribeLogContextResponse) {
    response = &DescribeLogContextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogContext
// 本接口用于搜索日志上下文附近的内容，详情参考[上下文检索](https://cloud.tencent.com/document/product/614/53248)。
//
// API返回数据包最大49MB，建议启用 gzip 压缩（HTTP Request Header Accept-Encoding:gzip）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogContext(request *DescribeLogContextRequest) (response *DescribeLogContextResponse, err error) {
    return c.DescribeLogContextWithContext(context.Background(), request)
}

// DescribeLogContext
// 本接口用于搜索日志上下文附近的内容，详情参考[上下文检索](https://cloud.tencent.com/document/product/614/53248)。
//
// API返回数据包最大49MB，建议启用 gzip 压缩（HTTP Request Header Accept-Encoding:gzip）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogContextWithContext(ctx context.Context, request *DescribeLogContextRequest) (response *DescribeLogContextResponse, err error) {
    if request == nil {
        request = NewDescribeLogContextRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogContext require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogContextResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogHistogramRequest() (request *DescribeLogHistogramRequest) {
    request = &DescribeLogHistogramRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeLogHistogram")
    
    
    return
}

func NewDescribeLogHistogramResponse() (response *DescribeLogHistogramResponse) {
    response = &DescribeLogHistogramResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogHistogram
// 本接口用于构建日志数量直方图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogHistogram(request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
    return c.DescribeLogHistogramWithContext(context.Background(), request)
}

// DescribeLogHistogram
// 本接口用于构建日志数量直方图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeLogHistogramWithContext(ctx context.Context, request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
    if request == nil {
        request = NewDescribeLogHistogramRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogHistogram require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogHistogramResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogsetsRequest() (request *DescribeLogsetsRequest) {
    request = &DescribeLogsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeLogsets")
    
    
    return
}

func NewDescribeLogsetsResponse() (response *DescribeLogsetsResponse) {
    response = &DescribeLogsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogsets
// 本接口用于获取日志集信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogsets(request *DescribeLogsetsRequest) (response *DescribeLogsetsResponse, err error) {
    return c.DescribeLogsetsWithContext(context.Background(), request)
}

// DescribeLogsets
// 本接口用于获取日志集信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLogsetsWithContext(ctx context.Context, request *DescribeLogsetsRequest) (response *DescribeLogsetsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineGroupConfigsRequest() (request *DescribeMachineGroupConfigsRequest) {
    request = &DescribeMachineGroupConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroupConfigs")
    
    
    return
}

func NewDescribeMachineGroupConfigsResponse() (response *DescribeMachineGroupConfigsResponse) {
    response = &DescribeMachineGroupConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMachineGroupConfigs
// 获取机器组绑定的采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachineGroupConfigs(request *DescribeMachineGroupConfigsRequest) (response *DescribeMachineGroupConfigsResponse, err error) {
    return c.DescribeMachineGroupConfigsWithContext(context.Background(), request)
}

// DescribeMachineGroupConfigs
// 获取机器组绑定的采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachineGroupConfigsWithContext(ctx context.Context, request *DescribeMachineGroupConfigsRequest) (response *DescribeMachineGroupConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeMachineGroupConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineGroupConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineGroupConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachineGroupsRequest() (request *DescribeMachineGroupsRequest) {
    request = &DescribeMachineGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroups")
    
    
    return
}

func NewDescribeMachineGroupsResponse() (response *DescribeMachineGroupsResponse) {
    response = &DescribeMachineGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMachineGroups
// 获取机器组信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMachineGroups(request *DescribeMachineGroupsRequest) (response *DescribeMachineGroupsResponse, err error) {
    return c.DescribeMachineGroupsWithContext(context.Background(), request)
}

// DescribeMachineGroups
// 获取机器组信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMachineGroupsWithContext(ctx context.Context, request *DescribeMachineGroupsRequest) (response *DescribeMachineGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeMachineGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachineGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachineGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
    request = &DescribeMachinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeMachines")
    
    
    return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
    response = &DescribeMachinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMachines
// 获取指定机器组下的机器状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_AGENTVERSIONNOTEXIST = "ResourceNotFound.AgentVersionNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    return c.DescribeMachinesWithContext(context.Background(), request)
}

// DescribeMachines
// 获取指定机器组下的机器状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_AGENTVERSIONNOTEXIST = "ResourceNotFound.AgentVersionNotExist"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) DescribeMachinesWithContext(ctx context.Context, request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
    if request == nil {
        request = NewDescribeMachinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMachines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMachinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNoticeContentsRequest() (request *DescribeNoticeContentsRequest) {
    request = &DescribeNoticeContentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeNoticeContents")
    
    
    return
}

func NewDescribeNoticeContentsResponse() (response *DescribeNoticeContentsResponse) {
    response = &DescribeNoticeContentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNoticeContents
// 获取通知内容列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeNoticeContents(request *DescribeNoticeContentsRequest) (response *DescribeNoticeContentsResponse, err error) {
    return c.DescribeNoticeContentsWithContext(context.Background(), request)
}

// DescribeNoticeContents
// 获取通知内容列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) DescribeNoticeContentsWithContext(ctx context.Context, request *DescribeNoticeContentsRequest) (response *DescribeNoticeContentsResponse, err error) {
    if request == nil {
        request = NewDescribeNoticeContentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNoticeContents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNoticeContentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePartitionsRequest() (request *DescribePartitionsRequest) {
    request = &DescribePartitionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribePartitions")
    
    
    return
}

func NewDescribePartitionsResponse() (response *DescribePartitionsResponse) {
    response = &DescribePartitionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePartitions
// 该接口已废弃，如需获取分区数量，请使用DescribeTopics接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribePartitions(request *DescribePartitionsRequest) (response *DescribePartitionsResponse, err error) {
    return c.DescribePartitionsWithContext(context.Background(), request)
}

// DescribePartitions
// 该接口已废弃，如需获取分区数量，请使用DescribeTopics接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribePartitionsWithContext(ctx context.Context, request *DescribePartitionsRequest) (response *DescribePartitionsResponse, err error) {
    if request == nil {
        request = NewDescribePartitionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePartitions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePartitionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScheduledSqlInfoRequest() (request *DescribeScheduledSqlInfoRequest) {
    request = &DescribeScheduledSqlInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeScheduledSqlInfo")
    
    
    return
}

func NewDescribeScheduledSqlInfoResponse() (response *DescribeScheduledSqlInfoResponse) {
    response = &DescribeScheduledSqlInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScheduledSqlInfo
// 本接口用于获取定时SQL分析任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeScheduledSqlInfo(request *DescribeScheduledSqlInfoRequest) (response *DescribeScheduledSqlInfoResponse, err error) {
    return c.DescribeScheduledSqlInfoWithContext(context.Background(), request)
}

// DescribeScheduledSqlInfo
// 本接口用于获取定时SQL分析任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeScheduledSqlInfoWithContext(ctx context.Context, request *DescribeScheduledSqlInfoRequest) (response *DescribeScheduledSqlInfoResponse, err error) {
    if request == nil {
        request = NewDescribeScheduledSqlInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScheduledSqlInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScheduledSqlInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShipperTasksRequest() (request *DescribeShipperTasksRequest) {
    request = &DescribeShipperTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeShipperTasks")
    
    
    return
}

func NewDescribeShipperTasksResponse() (response *DescribeShipperTasksResponse) {
    response = &DescribeShipperTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShipperTasks
// 获取投递任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DescribeShipperTasks(request *DescribeShipperTasksRequest) (response *DescribeShipperTasksResponse, err error) {
    return c.DescribeShipperTasksWithContext(context.Background(), request)
}

// DescribeShipperTasks
// 获取投递任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) DescribeShipperTasksWithContext(ctx context.Context, request *DescribeShipperTasksRequest) (response *DescribeShipperTasksResponse, err error) {
    if request == nil {
        request = NewDescribeShipperTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShipperTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShipperTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShippersRequest() (request *DescribeShippersRequest) {
    request = &DescribeShippersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeShippers")
    
    
    return
}

func NewDescribeShippersResponse() (response *DescribeShippersResponse) {
    response = &DescribeShippersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShippers
// 获取投递到COS的任务配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeShippers(request *DescribeShippersRequest) (response *DescribeShippersResponse, err error) {
    return c.DescribeShippersWithContext(context.Background(), request)
}

// DescribeShippers
// 获取投递到COS的任务配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) DescribeShippersWithContext(ctx context.Context, request *DescribeShippersRequest) (response *DescribeShippersResponse, err error) {
    if request == nil {
        request = NewDescribeShippersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShippers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShippersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
    request = &DescribeTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "DescribeTopics")
    
    
    return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
    response = &DescribeTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopics
// 本接口用于获取日志主题列表，支持分页
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ANALYSISSWITCHCLOSE = "OperationDenied.AnalysisSwitchClose"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    return c.DescribeTopicsWithContext(context.Background(), request)
}

// DescribeTopics
// 本接口用于获取日志主题列表，支持分页
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ANALYSISSWITCHCLOSE = "OperationDenied.AnalysisSwitchClose"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicsWithContext(ctx context.Context, request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewGetAlarmLogRequest() (request *GetAlarmLogRequest) {
    request = &GetAlarmLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "GetAlarmLog")
    
    
    return
}

func NewGetAlarmLogResponse() (response *GetAlarmLogResponse) {
    response = &GetAlarmLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAlarmLog
// 本接口用于获取告警策略执行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAlarmLog(request *GetAlarmLogRequest) (response *GetAlarmLogResponse, err error) {
    return c.GetAlarmLogWithContext(context.Background(), request)
}

// GetAlarmLog
// 本接口用于获取告警策略执行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetAlarmLogWithContext(ctx context.Context, request *GetAlarmLogRequest) (response *GetAlarmLogResponse, err error) {
    if request == nil {
        request = NewGetAlarmLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAlarmLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAlarmLogResponse()
    err = c.Send(request, response)
    return
}

func NewMergePartitionRequest() (request *MergePartitionRequest) {
    request = &MergePartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "MergePartition")
    
    
    return
}

func NewMergePartitionResponse() (response *MergePartitionResponse) {
    response = &MergePartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MergePartition
// 该接口已废弃，如需修改分区数量，请使用ModifyTopic接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MergePartition(request *MergePartitionRequest) (response *MergePartitionResponse, err error) {
    return c.MergePartitionWithContext(context.Background(), request)
}

// MergePartition
// 该接口已废弃，如需修改分区数量，请使用ModifyTopic接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MergePartitionWithContext(ctx context.Context, request *MergePartitionRequest) (response *MergePartitionResponse, err error) {
    if request == nil {
        request = NewMergePartitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MergePartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewMergePartitionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmRequest() (request *ModifyAlarmRequest) {
    request = &ModifyAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarm")
    
    
    return
}

func NewModifyAlarmResponse() (response *ModifyAlarmResponse) {
    response = &ModifyAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarm
// 本接口用于修改告警策略。需要至少修改一项有效内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDALARM = "FailedOperation.InvalidAlarm"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyAlarm(request *ModifyAlarmRequest) (response *ModifyAlarmResponse, err error) {
    return c.ModifyAlarmWithContext(context.Background(), request)
}

// ModifyAlarm
// 本接口用于修改告警策略。需要至少修改一项有效内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDALARM = "FailedOperation.InvalidAlarm"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyAlarmWithContext(ctx context.Context, request *ModifyAlarmRequest) (response *ModifyAlarmResponse, err error) {
    if request == nil {
        request = NewModifyAlarmRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmNoticeRequest() (request *ModifyAlarmNoticeRequest) {
    request = &ModifyAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarmNotice")
    
    
    return
}

func NewModifyAlarmNoticeResponse() (response *ModifyAlarmNoticeResponse) {
    response = &ModifyAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmNotice
// 该接口用于修改通知渠道组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyAlarmNotice(request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    return c.ModifyAlarmNoticeWithContext(context.Background(), request)
}

// ModifyAlarmNotice
// 该接口用于修改通知渠道组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyAlarmNoticeWithContext(ctx context.Context, request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmNoticeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmShieldRequest() (request *ModifyAlarmShieldRequest) {
    request = &ModifyAlarmShieldRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarmShield")
    
    
    return
}

func NewModifyAlarmShieldResponse() (response *ModifyAlarmShieldResponse) {
    response = &ModifyAlarmShieldResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmShield
// 该接口用于修改告警屏蔽规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyAlarmShield(request *ModifyAlarmShieldRequest) (response *ModifyAlarmShieldResponse, err error) {
    return c.ModifyAlarmShieldWithContext(context.Background(), request)
}

// ModifyAlarmShield
// 该接口用于修改告警屏蔽规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyAlarmShieldWithContext(ctx context.Context, request *ModifyAlarmShieldRequest) (response *ModifyAlarmShieldResponse, err error) {
    if request == nil {
        request = NewModifyAlarmShieldRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmShield require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmShieldResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudProductLogTaskRequest() (request *ModifyCloudProductLogTaskRequest) {
    request = &ModifyCloudProductLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyCloudProductLogTask")
    
    
    return
}

func NewModifyCloudProductLogTaskResponse() (response *ModifyCloudProductLogTaskResponse) {
    response = &ModifyCloudProductLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudProductLogTask
// 内部云产品接入使用相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudProductLogTask(request *ModifyCloudProductLogTaskRequest) (response *ModifyCloudProductLogTaskResponse, err error) {
    return c.ModifyCloudProductLogTaskWithContext(context.Background(), request)
}

// ModifyCloudProductLogTask
// 内部云产品接入使用相关接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyCloudProductLogTaskWithContext(ctx context.Context, request *ModifyCloudProductLogTaskRequest) (response *ModifyCloudProductLogTaskResponse, err error) {
    if request == nil {
        request = NewModifyCloudProductLogTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudProductLogTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudProductLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConfigRequest() (request *ModifyConfigRequest) {
    request = &ModifyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConfig")
    
    
    return
}

func NewModifyConfigResponse() (response *ModifyConfigResponse) {
    response = &ModifyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConfig
// 修改采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyConfig(request *ModifyConfigRequest) (response *ModifyConfigResponse, err error) {
    return c.ModifyConfigWithContext(context.Background(), request)
}

// ModifyConfig
// 修改采集规则配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyConfigWithContext(ctx context.Context, request *ModifyConfigRequest) (response *ModifyConfigResponse, err error) {
    if request == nil {
        request = NewModifyConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConfigExtraRequest() (request *ModifyConfigExtraRequest) {
    request = &ModifyConfigExtraRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConfigExtra")
    
    
    return
}

func NewModifyConfigExtraResponse() (response *ModifyConfigExtraResponse) {
    response = &ModifyConfigExtraResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConfigExtra
// 本接口用于修改特殊采集配置任务，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyConfigExtra(request *ModifyConfigExtraRequest) (response *ModifyConfigExtraResponse, err error) {
    return c.ModifyConfigExtraWithContext(context.Background(), request)
}

// ModifyConfigExtra
// 本接口用于修改特殊采集配置任务，特殊采集配置应用于自建K8S环境的采集Agent
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyConfigExtraWithContext(ctx context.Context, request *ModifyConfigExtraRequest) (response *ModifyConfigExtraResponse, err error) {
    if request == nil {
        request = NewModifyConfigExtraRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConfigExtra require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConfigExtraResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsoleSharingRequest() (request *ModifyConsoleSharingRequest) {
    request = &ModifyConsoleSharingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConsoleSharing")
    
    
    return
}

func NewModifyConsoleSharingResponse() (response *ModifyConsoleSharingResponse) {
    response = &ModifyConsoleSharingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsoleSharing
// 修改控制台分享，目前仅允许修改有效期
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConsoleSharing(request *ModifyConsoleSharingRequest) (response *ModifyConsoleSharingResponse, err error) {
    return c.ModifyConsoleSharingWithContext(context.Background(), request)
}

// ModifyConsoleSharing
// 修改控制台分享，目前仅允许修改有效期
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyConsoleSharingWithContext(ctx context.Context, request *ModifyConsoleSharingRequest) (response *ModifyConsoleSharingResponse, err error) {
    if request == nil {
        request = NewModifyConsoleSharingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsoleSharing require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsoleSharingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerRequest() (request *ModifyConsumerRequest) {
    request = &ModifyConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyConsumer")
    
    
    return
}

func NewModifyConsumerResponse() (response *ModifyConsumerResponse) {
    response = &ModifyConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumer
// 本接口用于修改投递Ckafka任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) ModifyConsumer(request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
    return c.ModifyConsumerWithContext(context.Background(), request)
}

// ModifyConsumer
// 本接口用于修改投递Ckafka任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"
func (c *Client) ModifyConsumerWithContext(ctx context.Context, request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
    if request == nil {
        request = NewModifyConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCosRechargeRequest() (request *ModifyCosRechargeRequest) {
    request = &ModifyCosRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyCosRecharge")
    
    
    return
}

func NewModifyCosRechargeResponse() (response *ModifyCosRechargeResponse) {
    response = &ModifyCosRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCosRecharge
// 本接口用于修改cos导入任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyCosRecharge(request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
    return c.ModifyCosRechargeWithContext(context.Background(), request)
}

// ModifyCosRecharge
// 本接口用于修改cos导入任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyCosRechargeWithContext(ctx context.Context, request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
    if request == nil {
        request = NewModifyCosRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCosRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCosRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDashboardSubscribeRequest() (request *ModifyDashboardSubscribeRequest) {
    request = &ModifyDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyDashboardSubscribe")
    
    
    return
}

func NewModifyDashboardSubscribeResponse() (response *ModifyDashboardSubscribeResponse) {
    response = &ModifyDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDashboardSubscribe
// 此接口用于修改仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) ModifyDashboardSubscribe(request *ModifyDashboardSubscribeRequest) (response *ModifyDashboardSubscribeResponse, err error) {
    return c.ModifyDashboardSubscribeWithContext(context.Background(), request)
}

// ModifyDashboardSubscribe
// 此接口用于修改仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
//  RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"
func (c *Client) ModifyDashboardSubscribeWithContext(ctx context.Context, request *ModifyDashboardSubscribeRequest) (response *ModifyDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewModifyDashboardSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataTransformRequest() (request *ModifyDataTransformRequest) {
    request = &ModifyDataTransformRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyDataTransform")
    
    
    return
}

func NewModifyDataTransformResponse() (response *ModifyDataTransformResponse) {
    response = &ModifyDataTransformResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDataTransform
// 本接口用于修改数据加工任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyDataTransform(request *ModifyDataTransformRequest) (response *ModifyDataTransformResponse, err error) {
    return c.ModifyDataTransformWithContext(context.Background(), request)
}

// ModifyDataTransform
// 本接口用于修改数据加工任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyDataTransformWithContext(ctx context.Context, request *ModifyDataTransformRequest) (response *ModifyDataTransformResponse, err error) {
    if request == nil {
        request = NewModifyDataTransformRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataTransform require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataTransformResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIndexRequest() (request *ModifyIndexRequest) {
    request = &ModifyIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyIndex")
    
    
    return
}

func NewModifyIndexResponse() (response *ModifyIndexResponse) {
    response = &ModifyIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIndex
// 本接口用于修改索引配置，该接口除受默认接口请求频率限制外，针对单个日志主题，并发数不能超过1，即同一时间同一个日志主题只能有一个正在执行的索引配置修改操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIndex(request *ModifyIndexRequest) (response *ModifyIndexResponse, err error) {
    return c.ModifyIndexWithContext(context.Background(), request)
}

// ModifyIndex
// 本接口用于修改索引配置，该接口除受默认接口请求频率限制外，针对单个日志主题，并发数不能超过1，即同一时间同一个日志主题只能有一个正在执行的索引配置修改操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"
//  LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyIndexWithContext(ctx context.Context, request *ModifyIndexRequest) (response *ModifyIndexResponse, err error) {
    if request == nil {
        request = NewModifyIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIndexResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKafkaConsumerRequest() (request *ModifyKafkaConsumerRequest) {
    request = &ModifyKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaConsumer")
    
    
    return
}

func NewModifyKafkaConsumerResponse() (response *ModifyKafkaConsumerResponse) {
    response = &ModifyKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKafkaConsumer
// 修改Kafka协议消费信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyKafkaConsumer(request *ModifyKafkaConsumerRequest) (response *ModifyKafkaConsumerResponse, err error) {
    return c.ModifyKafkaConsumerWithContext(context.Background(), request)
}

// ModifyKafkaConsumer
// 修改Kafka协议消费信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyKafkaConsumerWithContext(ctx context.Context, request *ModifyKafkaConsumerRequest) (response *ModifyKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewModifyKafkaConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKafkaRechargeRequest() (request *ModifyKafkaRechargeRequest) {
    request = &ModifyKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaRecharge")
    
    
    return
}

func NewModifyKafkaRechargeResponse() (response *ModifyKafkaRechargeResponse) {
    response = &ModifyKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyKafkaRecharge
// 本接口用于修改Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyKafkaRecharge(request *ModifyKafkaRechargeRequest) (response *ModifyKafkaRechargeResponse, err error) {
    return c.ModifyKafkaRechargeWithContext(context.Background(), request)
}

// ModifyKafkaRecharge
// 本接口用于修改Kafka数据订阅任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) ModifyKafkaRechargeWithContext(ctx context.Context, request *ModifyKafkaRechargeRequest) (response *ModifyKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewModifyKafkaRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLogsetRequest() (request *ModifyLogsetRequest) {
    request = &ModifyLogsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyLogset")
    
    
    return
}

func NewModifyLogsetResponse() (response *ModifyLogsetResponse) {
    response = &ModifyLogsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLogset
// 本接口用于修改日志集信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERIODMODIFYFORBIDDEN = "FailedOperation.PeriodModifyForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) ModifyLogset(request *ModifyLogsetRequest) (response *ModifyLogsetResponse, err error) {
    return c.ModifyLogsetWithContext(context.Background(), request)
}

// ModifyLogset
// 本接口用于修改日志集信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PERIODMODIFYFORBIDDEN = "FailedOperation.PeriodModifyForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"
func (c *Client) ModifyLogsetWithContext(ctx context.Context, request *ModifyLogsetRequest) (response *ModifyLogsetResponse, err error) {
    if request == nil {
        request = NewModifyLogsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLogset require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLogsetResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMachineGroupRequest() (request *ModifyMachineGroupRequest) {
    request = &ModifyMachineGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyMachineGroup")
    
    
    return
}

func NewModifyMachineGroupResponse() (response *ModifyMachineGroupResponse) {
    response = &ModifyMachineGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMachineGroup
// 修改机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ModifyMachineGroup(request *ModifyMachineGroupRequest) (response *ModifyMachineGroupResponse, err error) {
    return c.ModifyMachineGroupWithContext(context.Background(), request)
}

// ModifyMachineGroup
// 修改机器组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"
//  LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"
func (c *Client) ModifyMachineGroupWithContext(ctx context.Context, request *ModifyMachineGroupRequest) (response *ModifyMachineGroupResponse, err error) {
    if request == nil {
        request = NewModifyMachineGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMachineGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMachineGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNoticeContentRequest() (request *ModifyNoticeContentRequest) {
    request = &ModifyNoticeContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyNoticeContent")
    
    
    return
}

func NewModifyNoticeContentResponse() (response *ModifyNoticeContentResponse) {
    response = &ModifyNoticeContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNoticeContent
// 该接口用于修改通知内容配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyNoticeContent(request *ModifyNoticeContentRequest) (response *ModifyNoticeContentResponse, err error) {
    return c.ModifyNoticeContentWithContext(context.Background(), request)
}

// ModifyNoticeContent
// 该接口用于修改通知内容配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"
//  INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"
func (c *Client) ModifyNoticeContentWithContext(ctx context.Context, request *ModifyNoticeContentRequest) (response *ModifyNoticeContentResponse, err error) {
    if request == nil {
        request = NewModifyNoticeContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNoticeContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNoticeContentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScheduledSqlRequest() (request *ModifyScheduledSqlRequest) {
    request = &ModifyScheduledSqlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyScheduledSql")
    
    
    return
}

func NewModifyScheduledSqlResponse() (response *ModifyScheduledSqlResponse) {
    response = &ModifyScheduledSqlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyScheduledSql
// 本接口用于修改定时SQL分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyScheduledSql(request *ModifyScheduledSqlRequest) (response *ModifyScheduledSqlResponse, err error) {
    return c.ModifyScheduledSqlWithContext(context.Background(), request)
}

// ModifyScheduledSql
// 本接口用于修改定时SQL分析任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"
//  INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyScheduledSqlWithContext(ctx context.Context, request *ModifyScheduledSqlRequest) (response *ModifyScheduledSqlResponse, err error) {
    if request == nil {
        request = NewModifyScheduledSqlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyScheduledSql require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyScheduledSqlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyShipperRequest() (request *ModifyShipperRequest) {
    request = &ModifyShipperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyShipper")
    
    
    return
}

func NewModifyShipperResponse() (response *ModifyShipperResponse) {
    response = &ModifyShipperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyShipper
// 修改现有的投递规则，客户如果使用此接口，需要自行处理CLS对指定bucket的写权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) ModifyShipper(request *ModifyShipperRequest) (response *ModifyShipperResponse, err error) {
    return c.ModifyShipperWithContext(context.Background(), request)
}

// ModifyShipper
// 修改现有的投递规则，客户如果使用此接口，需要自行处理CLS对指定bucket的写权限。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
func (c *Client) ModifyShipperWithContext(ctx context.Context, request *ModifyShipperRequest) (response *ModifyShipperResponse, err error) {
    if request == nil {
        request = NewModifyShipperRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyShipper require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyShipperResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
    request = &ModifyTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "ModifyTopic")
    
    
    return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
    response = &ModifyTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTopic
// 本接口用于修改日志主题。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    return c.ModifyTopicWithContext(context.Background(), request)
}

// ModifyTopic
// 本接口用于修改日志主题。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"
//  FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) ModifyTopicWithContext(ctx context.Context, request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    if request == nil {
        request = NewModifyTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicResponse()
    err = c.Send(request, response)
    return
}

func NewOpenKafkaConsumerRequest() (request *OpenKafkaConsumerRequest) {
    request = &OpenKafkaConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "OpenKafkaConsumer")
    
    
    return
}

func NewOpenKafkaConsumerResponse() (response *OpenKafkaConsumerResponse) {
    response = &OpenKafkaConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenKafkaConsumer
// 打开Kafka协议消费功能
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXPORTCONFLICT = "InvalidParameter.ExportConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) OpenKafkaConsumer(request *OpenKafkaConsumerRequest) (response *OpenKafkaConsumerResponse, err error) {
    return c.OpenKafkaConsumerWithContext(context.Background(), request)
}

// OpenKafkaConsumer
// 打开Kafka协议消费功能
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXPORTCONFLICT = "InvalidParameter.ExportConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) OpenKafkaConsumerWithContext(ctx context.Context, request *OpenKafkaConsumerRequest) (response *OpenKafkaConsumerResponse, err error) {
    if request == nil {
        request = NewOpenKafkaConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenKafkaConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenKafkaConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewKafkaRechargeRequest() (request *PreviewKafkaRechargeRequest) {
    request = &PreviewKafkaRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "PreviewKafkaRecharge")
    
    
    return
}

func NewPreviewKafkaRechargeResponse() (response *PreviewKafkaRechargeResponse) {
    response = &PreviewKafkaRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PreviewKafkaRecharge
// 本接口用于预览Kafka数据订阅任务客户日志信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) PreviewKafkaRecharge(request *PreviewKafkaRechargeRequest) (response *PreviewKafkaRechargeResponse, err error) {
    return c.PreviewKafkaRechargeWithContext(context.Background(), request)
}

// PreviewKafkaRecharge
// 本接口用于预览Kafka数据订阅任务客户日志信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
func (c *Client) PreviewKafkaRechargeWithContext(ctx context.Context, request *PreviewKafkaRechargeRequest) (response *PreviewKafkaRechargeResponse, err error) {
    if request == nil {
        request = NewPreviewKafkaRechargeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PreviewKafkaRecharge require credential")
    }

    request.SetContext(ctx)
    
    response = NewPreviewKafkaRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMetricRequest() (request *QueryMetricRequest) {
    request = &QueryMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "QueryMetric")
    
    
    return
}

func NewQueryMetricResponse() (response *QueryMetricResponse) {
    response = &QueryMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryMetric
// 查询指定时刻指标的最新值。
//
// 如果该时刻向前推5分钟内均无指标数据，则无相应的查询结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) QueryMetric(request *QueryMetricRequest) (response *QueryMetricResponse, err error) {
    return c.QueryMetricWithContext(context.Background(), request)
}

// QueryMetric
// 查询指定时刻指标的最新值。
//
// 如果该时刻向前推5分钟内均无指标数据，则无相应的查询结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) QueryMetricWithContext(ctx context.Context, request *QueryMetricRequest) (response *QueryMetricResponse, err error) {
    if request == nil {
        request = NewQueryMetricRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMetric require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMetricResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRangeMetricRequest() (request *QueryRangeMetricRequest) {
    request = &QueryRangeMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "QueryRangeMetric")
    
    
    return
}

func NewQueryRangeMetricResponse() (response *QueryRangeMetricResponse) {
    response = &QueryRangeMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryRangeMetric
// 查询指定时间范围内指标的变化趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryRangeMetric(request *QueryRangeMetricRequest) (response *QueryRangeMetricResponse, err error) {
    return c.QueryRangeMetricWithContext(context.Background(), request)
}

// QueryRangeMetric
// 查询指定时间范围内指标的变化趋势
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryRangeMetricWithContext(ctx context.Context, request *QueryRangeMetricRequest) (response *QueryRangeMetricResponse, err error) {
    if request == nil {
        request = NewQueryRangeMetricRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryRangeMetric require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryRangeMetricResponse()
    err = c.Send(request, response)
    return
}

func NewRetryShipperTaskRequest() (request *RetryShipperTaskRequest) {
    request = &RetryShipperTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "RetryShipperTask")
    
    
    return
}

func NewRetryShipperTaskResponse() (response *RetryShipperTaskResponse) {
    response = &RetryShipperTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryShipperTask
// 重试失败的投递任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHIPPERTASKNOTTORETRY = "FailedOperation.ShipperTaskNotToRetry"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
//  RESOURCENOTFOUND_SHIPPERTASKNOTEXIST = "ResourceNotFound.ShipperTaskNotExist"
func (c *Client) RetryShipperTask(request *RetryShipperTaskRequest) (response *RetryShipperTaskResponse, err error) {
    return c.RetryShipperTaskWithContext(context.Background(), request)
}

// RetryShipperTask
// 重试失败的投递任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHIPPERTASKNOTTORETRY = "FailedOperation.ShipperTaskNotToRetry"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"
//  RESOURCENOTFOUND_SHIPPERTASKNOTEXIST = "ResourceNotFound.ShipperTaskNotExist"
func (c *Client) RetryShipperTaskWithContext(ctx context.Context, request *RetryShipperTaskRequest) (response *RetryShipperTaskResponse, err error) {
    if request == nil {
        request = NewRetryShipperTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryShipperTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryShipperTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSearchCosRechargeInfoRequest() (request *SearchCosRechargeInfoRequest) {
    request = &SearchCosRechargeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SearchCosRechargeInfo")
    
    
    return
}

func NewSearchCosRechargeInfoResponse() (response *SearchCosRechargeInfoResponse) {
    response = &SearchCosRechargeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchCosRechargeInfo
// 本接口用于预览cos导入信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOFILE = "FailedOperation.BucketNoFile"
//  FAILEDOPERATION_DECOMPRESSFILE = "FailedOperation.DecompressFile"
//  FAILEDOPERATION_DOWNLOADFILE = "FailedOperation.DownLoadFile"
//  FAILEDOPERATION_GETLISTFILE = "FailedOperation.GetListFile"
//  FAILEDOPERATION_PREVIEWFILE = "FailedOperation.PreviewFile"
//  FAILEDOPERATION_READFILE = "FailedOperation.ReadFile"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchCosRechargeInfo(request *SearchCosRechargeInfoRequest) (response *SearchCosRechargeInfoResponse, err error) {
    return c.SearchCosRechargeInfoWithContext(context.Background(), request)
}

// SearchCosRechargeInfo
// 本接口用于预览cos导入信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BUCKETNOFILE = "FailedOperation.BucketNoFile"
//  FAILEDOPERATION_DECOMPRESSFILE = "FailedOperation.DecompressFile"
//  FAILEDOPERATION_DOWNLOADFILE = "FailedOperation.DownLoadFile"
//  FAILEDOPERATION_GETLISTFILE = "FailedOperation.GetListFile"
//  FAILEDOPERATION_PREVIEWFILE = "FailedOperation.PreviewFile"
//  FAILEDOPERATION_READFILE = "FailedOperation.ReadFile"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchCosRechargeInfoWithContext(ctx context.Context, request *SearchCosRechargeInfoRequest) (response *SearchCosRechargeInfoResponse, err error) {
    if request == nil {
        request = NewSearchCosRechargeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchCosRechargeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchCosRechargeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSearchDashboardSubscribeRequest() (request *SearchDashboardSubscribeRequest) {
    request = &SearchDashboardSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SearchDashboardSubscribe")
    
    
    return
}

func NewSearchDashboardSubscribeResponse() (response *SearchDashboardSubscribeResponse) {
    response = &SearchDashboardSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchDashboardSubscribe
// 此接口用于预览仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) SearchDashboardSubscribe(request *SearchDashboardSubscribeRequest) (response *SearchDashboardSubscribeResponse, err error) {
    return c.SearchDashboardSubscribeWithContext(context.Background(), request)
}

// SearchDashboardSubscribe
// 此接口用于预览仪表盘订阅
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"
func (c *Client) SearchDashboardSubscribeWithContext(ctx context.Context, request *SearchDashboardSubscribeRequest) (response *SearchDashboardSubscribeResponse, err error) {
    if request == nil {
        request = NewSearchDashboardSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchDashboardSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchDashboardSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
    request = &SearchLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SearchLog")
    
    
    return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
    response = &SearchLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchLog
// 本接口用于检索分析日志，使用该接口时请注意如下事项：
//
// 1. 该接口除受默认接口请求频率限制外，针对单个日志主题，查询并发数不能超过15。
//
// 2. 检索语法建议使用CQL语法规则，请使用SyntaxRule参数，将值设置为1。
//
// 3. API返回数据包最大49MB，建议启用 gzip 压缩（HTTP Request Header Accept-Encoding:gzip）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESOURCES = "LimitExceeded.SearchResources"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  OPERATIONDENIED_OPERATIONNOTSUPPORTINSEARCHLOW = "OperationDenied.OperationNotSupportInSearchLow"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
    return c.SearchLogWithContext(context.Background(), request)
}

// SearchLog
// 本接口用于检索分析日志，使用该接口时请注意如下事项：
//
// 1. 该接口除受默认接口请求频率限制外，针对单个日志主题，查询并发数不能超过15。
//
// 2. 检索语法建议使用CQL语法规则，请使用SyntaxRule参数，将值设置为1。
//
// 3. API返回数据包最大49MB，建议启用 gzip 压缩（HTTP Request Header Accept-Encoding:gzip）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"
//  FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"
//  FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"
//  FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"
//  INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"
//  LIMITEXCEEDED_SEARCHRESOURCES = "LimitExceeded.SearchResources"
//  LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"
//  OPERATIONDENIED_OPERATIONNOTSUPPORTINSEARCHLOW = "OperationDenied.OperationNotSupportInSearchLow"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
func (c *Client) SearchLogWithContext(ctx context.Context, request *SearchLogRequest) (response *SearchLogResponse, err error) {
    if request == nil {
        request = NewSearchLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchLogResponse()
    err = c.Send(request, response)
    return
}

func NewSplitPartitionRequest() (request *SplitPartitionRequest) {
    request = &SplitPartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "SplitPartition")
    
    
    return
}

func NewSplitPartitionResponse() (response *SplitPartitionResponse) {
    response = &SplitPartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SplitPartition
// 该接口已废弃，如需修改分区数量，请使用ModifyTopic接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SplitPartition(request *SplitPartitionRequest) (response *SplitPartitionResponse, err error) {
    return c.SplitPartitionWithContext(context.Background(), request)
}

// SplitPartition
// 该接口已废弃，如需修改分区数量，请使用ModifyTopic接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SplitPartitionWithContext(ctx context.Context, request *SplitPartitionRequest) (response *SplitPartitionResponse, err error) {
    if request == nil {
        request = NewSplitPartitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SplitPartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewSplitPartitionResponse()
    err = c.Send(request, response)
    return
}

func NewUploadLogRequest() (request *UploadLogRequest) {
    request = &UploadLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cls", APIVersion, "UploadLog")
    request.SetContentType("application/octet-stream")
    
    return
}

func NewUploadLogResponse() (response *UploadLogResponse) {
    response = &UploadLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadLog
// ## 提示
//
// 为了保障您日志数据的可靠性以及更高效地使用日志服务，建议您使用CLS优化后的接口[上传结构化日志](https://cloud.tencent.com/document/product/614/16873)。
//
// 
//
// 同时我们给此接口专门优化定制了多个语言版本的SDK供您选择，SDK提供统一的异步发送、资源控制、自动重试、优雅关闭、感知上报等功能，使上报日志功能更完善，详情请参考[SDK采集](https://cloud.tencent.com/document/product/614/67157)。
//
// 
//
// 同时云API上传日志接口也支持同步上传日志数据，如果您选择继续使用此接口请参考下文。
//
// 
//
// ## 功能描述
//
// 
//
// 本接口用于将日志写入到指定的日志主题。
//
// 
//
// #### 输入参数(pb二进制流，位于body中)
//
// 
//
// | 字段名       | 类型    | 位置 | 必须 | 含义                                                         |
//
// | ------------ | ------- | ---- | ---- | ------------------------------------------------------------ |
//
// | logGroupList | message | pb   | 是   | logGroup 列表，封装好的日志组列表内容，建议 logGroup 数量不要超过5个 |
//
// 
//
// LogGroup 说明：
//
// 
//
// | 字段名      | 是否必选 | 含义                                                         |
//
// | ----------- | -------- | ------------------------------------------------------------ |
//
// | logs        | 是       | 日志数组，表示有多个 Log 组成的集合，一个 Log 表示一条日志，一个 LogGroup 中 Log 个数不能超过10000 |
//
// | contextFlow | 否       | LogGroup 的唯一ID，需要使用上下文功能时传入。格式："{上下文ID}-{LogGroupID}"。<br>上下文ID：唯一标识一个上下文（连续滚动的一系列日志文件，或者是需要保序的一系列日志），16进制64位整型字符串。<br>LogGroupID：连续递增的一串整型，16进制64位整型字符串。样例："102700A66102516A-59F59"。                        |
//
// | filename    | 否       | 日志文件名                                                   |
//
// | source      | 否       | 日志来源，一般使用机器 IP 作为标识                           |
//
// | logTags     | 否       | 日志的标签列表                                               |
//
// 
//
// Log 说明：
//
// 
//
// | 字段名   | 是否必选 | 含义                                                         |
//
// | -------- | -------- | ------------------------------------------------------------ |
//
// | time     | 是       | 日志时间（Unix 格式时间戳），支持秒、毫秒，建议采用毫秒      |
//
// | contents | 否       | key-value 格式的日志内容，表示一条日志里的多个 key-value 组合 |
//
// 
//
// Content 说明：
//
// 
//
// | 字段名 | 是否必选 | 含义                                                         |
//
// | ------ | -------- | ------------------------------------------------------------ |
//
// | key    | 是       | 单条日志里某个字段组的 key 值，不能以`_`开头                 |
//
// | value  | 是       | 单条日志某个字段组的 value 值，单条日志 value 不能超过1MB，LogGroup 中所有 value 总和不能超过5MB |
//
// 
//
// LogTag 说明：
//
// 
//
// | 字段名 | 是否必选 | 含义                             |
//
// | ------ | -------- | -------------------------------- |
//
// | key    | 是       | 自定义的标签 key                 |
//
// | value  | 是       | 自定义的标签 key 对应的 value 值 |
//
// 
//
// ## PB 编译示例
//
// 
//
// 本示例将说明如何使用官方 protoc 编译工具将 PB 描述文件 编译生成为 C++ 语言可调用的上传日志接口。
//
// 
//
// > ?目前 protoc 官方支持 Java、C++、Python 等语言的编译，详情请参见 [protoc](https://github.com/protocolbuffers/protobuf)。
//
// 
//
// #### 1. 安装 Protocol Buffer
//
// 
//
// 下载 [Protocol Buffer](https://main.qcloudimg.com/raw/d7810aaf8b3073fbbc9d4049c21532aa/protobuf-2.6.1.tar.gz) ，解压并安装。示例版本为 protobuf 2.6.1，环境为 Centos 7.3 系统。 解压`protobuf-2.6.1.tar.gz`压缩包至`/usr/local`目录并进入该目录，执行命令如下：
//
// 
//
// ```
//
// tar -zxvf protobuf-2.6.1.tar.gz -C /usr/local/ && cd /usr/local/protobuf-2.6.1
//
// ```
//
// 
//
// 开始编译和安装，配置环境变量，执行命令如下：
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# ./configure 
//
// [root@VM_0_8_centos protobuf-2.6.1]# make && make install
//
// [root@VM_0_8_centos protobuf-2.6.1]# export PATH=$PATH:/usr/local/protobuf-2.6.1/bin
//
// ```
//
// 
//
// 编译成功后，您可以使用以下命令查看版本：
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --version
//
// liprotoc 2.6.1
//
// ```
//
// 
//
// #### 2. 创建 PB 描述文件
//
// 
//
// PB 描述文件是通信双方约定的数据交换格式，上传日志时须将规定的协议格式编译成对应语言版本的调用接口，然后添加到工程代码里，详情请参见 [protoc](https://github.com/protocolbuffers/protobuf) 。
//
// 
//
// 以日志服务所规定的 PB 数据格式内容为准， 在本地创建 PB 消息描述文件 cls.proto。
//
// 
//
// > !PB 描述文件内容不可更改，且文件名须以`.proto`结尾。
//
// 
//
// cls.proto 内容（PB 描述文件）如下：
//
// 
//
// ```
//
// package cls;
//
// 
//
// message Log
//
// {
//
//     message Content
//
//     {
//
//         required string key   = 1; // 每组字段的 key
//
//         required string value = 2; // 每组字段的 value
//
//     }
//
//     required int64   time     = 1; // 时间戳，UNIX时间格式
//
//     repeated Content contents = 2; // 一条日志里的多个kv组合
//
// }
//
// 
//
// message LogTag
//
// {
//
//     required string key       = 1;
//
//     required string value     = 2;
//
// }
//
// 
//
// message LogGroup
//
// {
//
//     repeated Log    logs        = 1; // 多条日志合成的日志数组
//
//     optional string contextFlow = 2; // 目前暂无效用
//
//     optional string filename    = 3; // 日志文件名
//
//     optional string source      = 4; // 日志来源，一般使用机器IP
//
//     repeated LogTag logTags     = 5;
//
// }
//
// 
//
// message LogGroupList
//
// {
//
//     repeated LogGroup logGroupList = 1; // 日志组列表
//
// }
//
// ```
//
// 
//
// #### 3. 编译生成
//
// 
//
// 此例中，使用 proto 编译器生成 C++ 语言的文件，在 cls.proto 文件的同一目录下，执行如下编译命令：
//
// 
//
// ```
//
// protoc --cpp_out=./ ./cls.proto 
//
// ```
//
// 
//
// > ?`--cpp_out=./`表示编译成 cpp 格式并输出当前目录下，`./cls.proto`表示位于当前目录下的 cls.proto 描述文件。
//
// 
//
// 编译成功后，会输出对应语言的代码文件。此例会生成 cls.pb.h 头文件和 [cls.pb.cc](http://cls.pb.cc) 代码实现文件，如下所示：
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --cpp_out=./ ./cls.proto
//
// [root@VM_0_8_centos protobuf-2.6.1]# ls
//
// cls.pb.cc cls.pb.h cls.proto
//
// ```
//
// 
//
// #### 4. 调用
//
// 
//
// 将生成的 cls.pb.h 头文件引入代码中，调用接口进行数据格式封装。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MISSINGCONTENT = "FailedOperation.MissingContent"
//  FAILEDOPERATION_READQPSLIMIT = "FailedOperation.ReadQpsLimit"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  FAILEDOPERATION_WRITEQPSLIMIT = "FailedOperation.WriteQpsLimit"
//  FAILEDOPERATION_WRITETRAFFICLIMIT = "FailedOperation.WriteTrafficLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  LIMITEXCEEDED_LOGSIZE = "LimitExceeded.LogSize"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadLog(request *UploadLogRequest, data []byte) (response *UploadLogResponse, err error) {
    return c.UploadLogWithContext(context.Background(), request, data)
}

// UploadLog
// ## 提示
//
// 为了保障您日志数据的可靠性以及更高效地使用日志服务，建议您使用CLS优化后的接口[上传结构化日志](https://cloud.tencent.com/document/product/614/16873)。
//
// 
//
// 同时我们给此接口专门优化定制了多个语言版本的SDK供您选择，SDK提供统一的异步发送、资源控制、自动重试、优雅关闭、感知上报等功能，使上报日志功能更完善，详情请参考[SDK采集](https://cloud.tencent.com/document/product/614/67157)。
//
// 
//
// 同时云API上传日志接口也支持同步上传日志数据，如果您选择继续使用此接口请参考下文。
//
// 
//
// ## 功能描述
//
// 
//
// 本接口用于将日志写入到指定的日志主题。
//
// 
//
// #### 输入参数(pb二进制流，位于body中)
//
// 
//
// | 字段名       | 类型    | 位置 | 必须 | 含义                                                         |
//
// | ------------ | ------- | ---- | ---- | ------------------------------------------------------------ |
//
// | logGroupList | message | pb   | 是   | logGroup 列表，封装好的日志组列表内容，建议 logGroup 数量不要超过5个 |
//
// 
//
// LogGroup 说明：
//
// 
//
// | 字段名      | 是否必选 | 含义                                                         |
//
// | ----------- | -------- | ------------------------------------------------------------ |
//
// | logs        | 是       | 日志数组，表示有多个 Log 组成的集合，一个 Log 表示一条日志，一个 LogGroup 中 Log 个数不能超过10000 |
//
// | contextFlow | 否       | LogGroup 的唯一ID，需要使用上下文功能时传入。格式："{上下文ID}-{LogGroupID}"。<br>上下文ID：唯一标识一个上下文（连续滚动的一系列日志文件，或者是需要保序的一系列日志），16进制64位整型字符串。<br>LogGroupID：连续递增的一串整型，16进制64位整型字符串。样例："102700A66102516A-59F59"。                        |
//
// | filename    | 否       | 日志文件名                                                   |
//
// | source      | 否       | 日志来源，一般使用机器 IP 作为标识                           |
//
// | logTags     | 否       | 日志的标签列表                                               |
//
// 
//
// Log 说明：
//
// 
//
// | 字段名   | 是否必选 | 含义                                                         |
//
// | -------- | -------- | ------------------------------------------------------------ |
//
// | time     | 是       | 日志时间（Unix 格式时间戳），支持秒、毫秒，建议采用毫秒      |
//
// | contents | 否       | key-value 格式的日志内容，表示一条日志里的多个 key-value 组合 |
//
// 
//
// Content 说明：
//
// 
//
// | 字段名 | 是否必选 | 含义                                                         |
//
// | ------ | -------- | ------------------------------------------------------------ |
//
// | key    | 是       | 单条日志里某个字段组的 key 值，不能以`_`开头                 |
//
// | value  | 是       | 单条日志某个字段组的 value 值，单条日志 value 不能超过1MB，LogGroup 中所有 value 总和不能超过5MB |
//
// 
//
// LogTag 说明：
//
// 
//
// | 字段名 | 是否必选 | 含义                             |
//
// | ------ | -------- | -------------------------------- |
//
// | key    | 是       | 自定义的标签 key                 |
//
// | value  | 是       | 自定义的标签 key 对应的 value 值 |
//
// 
//
// ## PB 编译示例
//
// 
//
// 本示例将说明如何使用官方 protoc 编译工具将 PB 描述文件 编译生成为 C++ 语言可调用的上传日志接口。
//
// 
//
// > ?目前 protoc 官方支持 Java、C++、Python 等语言的编译，详情请参见 [protoc](https://github.com/protocolbuffers/protobuf)。
//
// 
//
// #### 1. 安装 Protocol Buffer
//
// 
//
// 下载 [Protocol Buffer](https://main.qcloudimg.com/raw/d7810aaf8b3073fbbc9d4049c21532aa/protobuf-2.6.1.tar.gz) ，解压并安装。示例版本为 protobuf 2.6.1，环境为 Centos 7.3 系统。 解压`protobuf-2.6.1.tar.gz`压缩包至`/usr/local`目录并进入该目录，执行命令如下：
//
// 
//
// ```
//
// tar -zxvf protobuf-2.6.1.tar.gz -C /usr/local/ && cd /usr/local/protobuf-2.6.1
//
// ```
//
// 
//
// 开始编译和安装，配置环境变量，执行命令如下：
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# ./configure 
//
// [root@VM_0_8_centos protobuf-2.6.1]# make && make install
//
// [root@VM_0_8_centos protobuf-2.6.1]# export PATH=$PATH:/usr/local/protobuf-2.6.1/bin
//
// ```
//
// 
//
// 编译成功后，您可以使用以下命令查看版本：
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --version
//
// liprotoc 2.6.1
//
// ```
//
// 
//
// #### 2. 创建 PB 描述文件
//
// 
//
// PB 描述文件是通信双方约定的数据交换格式，上传日志时须将规定的协议格式编译成对应语言版本的调用接口，然后添加到工程代码里，详情请参见 [protoc](https://github.com/protocolbuffers/protobuf) 。
//
// 
//
// 以日志服务所规定的 PB 数据格式内容为准， 在本地创建 PB 消息描述文件 cls.proto。
//
// 
//
// > !PB 描述文件内容不可更改，且文件名须以`.proto`结尾。
//
// 
//
// cls.proto 内容（PB 描述文件）如下：
//
// 
//
// ```
//
// package cls;
//
// 
//
// message Log
//
// {
//
//     message Content
//
//     {
//
//         required string key   = 1; // 每组字段的 key
//
//         required string value = 2; // 每组字段的 value
//
//     }
//
//     required int64   time     = 1; // 时间戳，UNIX时间格式
//
//     repeated Content contents = 2; // 一条日志里的多个kv组合
//
// }
//
// 
//
// message LogTag
//
// {
//
//     required string key       = 1;
//
//     required string value     = 2;
//
// }
//
// 
//
// message LogGroup
//
// {
//
//     repeated Log    logs        = 1; // 多条日志合成的日志数组
//
//     optional string contextFlow = 2; // 目前暂无效用
//
//     optional string filename    = 3; // 日志文件名
//
//     optional string source      = 4; // 日志来源，一般使用机器IP
//
//     repeated LogTag logTags     = 5;
//
// }
//
// 
//
// message LogGroupList
//
// {
//
//     repeated LogGroup logGroupList = 1; // 日志组列表
//
// }
//
// ```
//
// 
//
// #### 3. 编译生成
//
// 
//
// 此例中，使用 proto 编译器生成 C++ 语言的文件，在 cls.proto 文件的同一目录下，执行如下编译命令：
//
// 
//
// ```
//
// protoc --cpp_out=./ ./cls.proto 
//
// ```
//
// 
//
// > ?`--cpp_out=./`表示编译成 cpp 格式并输出当前目录下，`./cls.proto`表示位于当前目录下的 cls.proto 描述文件。
//
// 
//
// 编译成功后，会输出对应语言的代码文件。此例会生成 cls.pb.h 头文件和 [cls.pb.cc](http://cls.pb.cc) 代码实现文件，如下所示：
//
// 
//
// ```
//
// [root@VM_0_8_centos protobuf-2.6.1]# protoc --cpp_out=./ ./cls.proto
//
// [root@VM_0_8_centos protobuf-2.6.1]# ls
//
// cls.pb.cc cls.pb.h cls.proto
//
// ```
//
// 
//
// #### 4. 调用
//
// 
//
// 将生成的 cls.pb.h 头文件引入代码中，调用接口进行数据格式封装。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MISSINGCONTENT = "FailedOperation.MissingContent"
//  FAILEDOPERATION_READQPSLIMIT = "FailedOperation.ReadQpsLimit"
//  FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"
//  FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"
//  FAILEDOPERATION_WRITEQPSLIMIT = "FailedOperation.WriteQpsLimit"
//  FAILEDOPERATION_WRITETRAFFICLIMIT = "FailedOperation.WriteTrafficLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  LIMITEXCEEDED_LOGSIZE = "LimitExceeded.LogSize"
//  LIMITEXCEEDED_TAG = "LimitExceeded.Tag"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"
//  OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"
//  OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"
//  OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"
//  RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"
//  RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadLogWithContext(ctx context.Context, request *UploadLogRequest, data []byte) (response *UploadLogResponse, err error) {
    if request == nil {
        request = NewUploadLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadLog require credential")
    }

    request.SetContext(ctx)
    request.SetBody(data)
    response = NewUploadLogResponse()
    err = c.SendOctetStream(request, response)
    return
}
