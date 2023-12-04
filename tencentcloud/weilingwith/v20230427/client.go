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

package v20230427

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-04-27"

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


func NewAddAlarmProcessRecordRequest() (request *AddAlarmProcessRecordRequest) {
    request = &AddAlarmProcessRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "AddAlarmProcessRecord")
    
    
    return
}

func NewAddAlarmProcessRecordResponse() (response *AddAlarmProcessRecordResponse) {
    response = &AddAlarmProcessRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddAlarmProcessRecord
// 添加告警处理记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDTENANTID = "InvalidParameterValue.InvalidTenantId"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  LIMITEXCEEDED_APILIMITEXCEEDED = "LimitExceeded.ApiLimitExceeded"
//  OPERATIONDENIED_APIPERMISSIONDENIED = "OperationDenied.ApiPermissionDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ALARMIDNOTEXIST = "ResourceNotFound.AlarmIdNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) AddAlarmProcessRecord(request *AddAlarmProcessRecordRequest) (response *AddAlarmProcessRecordResponse, err error) {
    return c.AddAlarmProcessRecordWithContext(context.Background(), request)
}

// AddAlarmProcessRecord
// 添加告警处理记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDTENANTID = "InvalidParameterValue.InvalidTenantId"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  LIMITEXCEEDED_APILIMITEXCEEDED = "LimitExceeded.ApiLimitExceeded"
//  OPERATIONDENIED_APIPERMISSIONDENIED = "OperationDenied.ApiPermissionDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ALARMIDNOTEXIST = "ResourceNotFound.AlarmIdNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) AddAlarmProcessRecordWithContext(ctx context.Context, request *AddAlarmProcessRecordRequest) (response *AddAlarmProcessRecordResponse, err error) {
    if request == nil {
        request = NewAddAlarmProcessRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAlarmProcessRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAlarmProcessRecordResponse()
    err = c.Send(request, response)
    return
}

func NewBatchCreateDeviceRequest() (request *BatchCreateDeviceRequest) {
    request = &BatchCreateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "BatchCreateDevice")
    
    
    return
}

func NewBatchCreateDeviceResponse() (response *BatchCreateDeviceResponse) {
    response = &BatchCreateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchCreateDevice
// 单个/批量新增设备
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_MANAGERSRVFAILED = "InternalError.ManagerSrvFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  LIMITEXCEEDED_DEVICELIMITEXCEEDED = "LimitExceeded.DeviceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) BatchCreateDevice(request *BatchCreateDeviceRequest) (response *BatchCreateDeviceResponse, err error) {
    return c.BatchCreateDeviceWithContext(context.Background(), request)
}

// BatchCreateDevice
// 单个/批量新增设备
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_MANAGERSRVFAILED = "InternalError.ManagerSrvFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  LIMITEXCEEDED_DEVICELIMITEXCEEDED = "LimitExceeded.DeviceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) BatchCreateDeviceWithContext(ctx context.Context, request *BatchCreateDeviceRequest) (response *BatchCreateDeviceResponse, err error) {
    if request == nil {
        request = NewBatchCreateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewBatchKillAlarmRequest() (request *BatchKillAlarmRequest) {
    request = &BatchKillAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "BatchKillAlarm")
    
    
    return
}

func NewBatchKillAlarmResponse() (response *BatchKillAlarmResponse) {
    response = &BatchKillAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchKillAlarm
// 批量消警
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  OPERATIONDENIED_APIPERMISSIONDENIED = "OperationDenied.ApiPermissionDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) BatchKillAlarm(request *BatchKillAlarmRequest) (response *BatchKillAlarmResponse, err error) {
    return c.BatchKillAlarmWithContext(context.Background(), request)
}

// BatchKillAlarm
// 批量消警
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  OPERATIONDENIED_APIPERMISSIONDENIED = "OperationDenied.ApiPermissionDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) BatchKillAlarmWithContext(ctx context.Context, request *BatchKillAlarmRequest) (response *BatchKillAlarmResponse, err error) {
    if request == nil {
        request = NewBatchKillAlarmRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchKillAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchKillAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewBatchReportAppMessageRequest() (request *BatchReportAppMessageRequest) {
    request = &BatchReportAppMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "BatchReportAppMessage")
    
    
    return
}

func NewBatchReportAppMessageResponse() (response *BatchReportAppMessageResponse) {
    response = &BatchReportAppMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchReportAppMessage
// 批量上报应用消息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BatchReportAppMessage(request *BatchReportAppMessageRequest) (response *BatchReportAppMessageResponse, err error) {
    return c.BatchReportAppMessageWithContext(context.Background(), request)
}

// BatchReportAppMessage
// 批量上报应用消息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BatchReportAppMessageWithContext(ctx context.Context, request *BatchReportAppMessageRequest) (response *BatchReportAppMessageResponse, err error) {
    if request == nil {
        request = NewBatchReportAppMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchReportAppMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchReportAppMessageResponse()
    err = c.Send(request, response)
    return
}

func NewChangeAlarmStatusRequest() (request *ChangeAlarmStatusRequest) {
    request = &ChangeAlarmStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "ChangeAlarmStatus")
    
    
    return
}

func NewChangeAlarmStatusResponse() (response *ChangeAlarmStatusResponse) {
    response = &ChangeAlarmStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeAlarmStatus
// 变更告警状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDALARMCONTENT = "InvalidParameter.InvalidAlarmContent"
//  INVALIDPARAMETER_STATUSNOTMATCHPROCESSTYPE = "InvalidParameter.StatusNotMatchProcessType"
//  INVALIDPARAMETERVALUE_INVALIDALARMSTATUS = "InvalidParameterValue.InvalidAlarmStatus"
//  INVALIDPARAMETERVALUE_INVALIDTENANTID = "InvalidParameterValue.InvalidTenantId"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  MISSINGPARAMETER_EMPTYALARMID = "MissingParameter.EmptyAlarmId"
//  OPERATIONDENIED_APIPERMISSIONDENIED = "OperationDenied.ApiPermissionDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ALARMIDNOTFOUND = "ResourceNotFound.AlarmIDNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) ChangeAlarmStatus(request *ChangeAlarmStatusRequest) (response *ChangeAlarmStatusResponse, err error) {
    return c.ChangeAlarmStatusWithContext(context.Background(), request)
}

// ChangeAlarmStatus
// 变更告警状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDALARMCONTENT = "InvalidParameter.InvalidAlarmContent"
//  INVALIDPARAMETER_STATUSNOTMATCHPROCESSTYPE = "InvalidParameter.StatusNotMatchProcessType"
//  INVALIDPARAMETERVALUE_INVALIDALARMSTATUS = "InvalidParameterValue.InvalidAlarmStatus"
//  INVALIDPARAMETERVALUE_INVALIDTENANTID = "InvalidParameterValue.InvalidTenantId"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  MISSINGPARAMETER_EMPTYALARMID = "MissingParameter.EmptyAlarmId"
//  OPERATIONDENIED_APIPERMISSIONDENIED = "OperationDenied.ApiPermissionDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ALARMIDNOTFOUND = "ResourceNotFound.AlarmIDNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) ChangeAlarmStatusWithContext(ctx context.Context, request *ChangeAlarmStatusRequest) (response *ChangeAlarmStatusResponse, err error) {
    if request == nil {
        request = NewChangeAlarmStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeAlarmStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeAlarmStatusResponse()
    err = c.Send(request, response)
    return
}

func NewControlCameraPTZRequest() (request *ControlCameraPTZRequest) {
    request = &ControlCameraPTZRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "ControlCameraPTZ")
    
    
    return
}

func NewControlCameraPTZResponse() (response *ControlCameraPTZResponse) {
    response = &ControlCameraPTZResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlCameraPTZ
// 云台控制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCMD = "UnsupportedOperation.UnsupportedCMD"
func (c *Client) ControlCameraPTZ(request *ControlCameraPTZRequest) (response *ControlCameraPTZResponse, err error) {
    return c.ControlCameraPTZWithContext(context.Background(), request)
}

// ControlCameraPTZ
// 云台控制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCMD = "UnsupportedOperation.UnsupportedCMD"
func (c *Client) ControlCameraPTZWithContext(ctx context.Context, request *ControlCameraPTZRequest) (response *ControlCameraPTZResponse, err error) {
    if request == nil {
        request = NewControlCameraPTZRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlCameraPTZ require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlCameraPTZResponse()
    err = c.Send(request, response)
    return
}

func NewControlDeviceRequest() (request *ControlDeviceRequest) {
    request = &ControlDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "ControlDevice")
    
    
    return
}

func NewControlDeviceResponse() (response *ControlDeviceResponse) {
    response = &ControlDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlDevice
// 设备控制（单个、批量控制）
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  FAILEDOPERATION_SENDMSGERROR = "FailedOperation.SendMsgError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_APIAUTHFAILED = "UnauthorizedOperation.APIAuthFailed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) ControlDevice(request *ControlDeviceRequest) (response *ControlDeviceResponse, err error) {
    return c.ControlDeviceWithContext(context.Background(), request)
}

// ControlDevice
// 设备控制（单个、批量控制）
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  FAILEDOPERATION_SENDMSGERROR = "FailedOperation.SendMsgError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_APIAUTHFAILED = "UnauthorizedOperation.APIAuthFailed"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) ControlDeviceWithContext(ctx context.Context, request *ControlDeviceRequest) (response *ControlDeviceResponse, err error) {
    if request == nil {
        request = NewControlDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationTokenRequest() (request *CreateApplicationTokenRequest) {
    request = &CreateApplicationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "CreateApplicationToken")
    
    
    return
}

func NewCreateApplicationTokenResponse() (response *CreateApplicationTokenResponse) {
    response = &CreateApplicationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationToken
// 调用方应用，创建调用租户API的授权令牌。
//
// 可能返回的错误码:
//  INTERNALERROR_BUSINESSLOGICERROR = "InternalError.BusinessLogicError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_INVALIDREQUESTTIME = "InvalidParameter.InvalidRequestTime"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDNONCE = "InvalidParameterValue.InvalidNonce"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTTIME = "InvalidParameterValue.InvalidRequestTime"
//  INVALIDPARAMETERVALUE_INVALIDSIGNATURE = "InvalidParameterValue.InvalidSignature"
func (c *Client) CreateApplicationToken(request *CreateApplicationTokenRequest) (response *CreateApplicationTokenResponse, err error) {
    return c.CreateApplicationTokenWithContext(context.Background(), request)
}

// CreateApplicationToken
// 调用方应用，创建调用租户API的授权令牌。
//
// 可能返回的错误码:
//  INTERNALERROR_BUSINESSLOGICERROR = "InternalError.BusinessLogicError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER_INVALIDREQUESTTIME = "InvalidParameter.InvalidRequestTime"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
//  INVALIDPARAMETERVALUE_INVALIDNONCE = "InvalidParameterValue.InvalidNonce"
//  INVALIDPARAMETERVALUE_INVALIDREQUESTTIME = "InvalidParameterValue.InvalidRequestTime"
//  INVALIDPARAMETERVALUE_INVALIDSIGNATURE = "InvalidParameterValue.InvalidSignature"
func (c *Client) CreateApplicationTokenWithContext(ctx context.Context, request *CreateApplicationTokenRequest) (response *CreateApplicationTokenResponse, err error) {
    if request == nil {
        request = NewCreateApplicationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActionListRequest() (request *DescribeActionListRequest) {
    request = &DescribeActionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeActionList")
    
    
    return
}

func NewDescribeActionListResponse() (response *DescribeActionListResponse) {
    response = &DescribeActionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeActionList
// 动作列表查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeActionList(request *DescribeActionListRequest) (response *DescribeActionListResponse, err error) {
    return c.DescribeActionListWithContext(context.Background(), request)
}

// DescribeActionList
// 动作列表查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeActionListWithContext(ctx context.Context, request *DescribeActionListRequest) (response *DescribeActionListResponse, err error) {
    if request == nil {
        request = NewDescribeActionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeActionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeActionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAdministrationByTagRequest() (request *DescribeAdministrationByTagRequest) {
    request = &DescribeAdministrationByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeAdministrationByTag")
    
    
    return
}

func NewDescribeAdministrationByTagResponse() (response *DescribeAdministrationByTagResponse) {
    response = &DescribeAdministrationByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAdministrationByTag
// 根据标签获取行政区划列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAdministrationByTag(request *DescribeAdministrationByTagRequest) (response *DescribeAdministrationByTagResponse, err error) {
    return c.DescribeAdministrationByTagWithContext(context.Background(), request)
}

// DescribeAdministrationByTag
// 根据标签获取行政区划列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAdministrationByTagWithContext(ctx context.Context, request *DescribeAdministrationByTagRequest) (response *DescribeAdministrationByTagResponse, err error) {
    if request == nil {
        request = NewDescribeAdministrationByTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAdministrationByTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAdministrationByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmLevelListRequest() (request *DescribeAlarmLevelListRequest) {
    request = &DescribeAlarmLevelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeAlarmLevelList")
    
    
    return
}

func NewDescribeAlarmLevelListResponse() (response *DescribeAlarmLevelListResponse) {
    response = &DescribeAlarmLevelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmLevelList
// 告警级别枚举获取
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmLevelList(request *DescribeAlarmLevelListRequest) (response *DescribeAlarmLevelListResponse, err error) {
    return c.DescribeAlarmLevelListWithContext(context.Background(), request)
}

// DescribeAlarmLevelList
// 告警级别枚举获取
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmLevelListWithContext(ctx context.Context, request *DescribeAlarmLevelListRequest) (response *DescribeAlarmLevelListResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmLevelListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmLevelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmLevelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmListRequest() (request *DescribeAlarmListRequest) {
    request = &DescribeAlarmListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeAlarmList")
    
    
    return
}

func NewDescribeAlarmListResponse() (response *DescribeAlarmListResponse) {
    response = &DescribeAlarmListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmList
// 告警列表查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeAlarmList(request *DescribeAlarmListRequest) (response *DescribeAlarmListResponse, err error) {
    return c.DescribeAlarmListWithContext(context.Background(), request)
}

// DescribeAlarmList
// 告警列表查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeAlarmListWithContext(ctx context.Context, request *DescribeAlarmListRequest) (response *DescribeAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmStatusListRequest() (request *DescribeAlarmStatusListRequest) {
    request = &DescribeAlarmStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeAlarmStatusList")
    
    
    return
}

func NewDescribeAlarmStatusListResponse() (response *DescribeAlarmStatusListResponse) {
    response = &DescribeAlarmStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmStatusList
// 用来查询系统中的告警状态列表
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmStatusList(request *DescribeAlarmStatusListRequest) (response *DescribeAlarmStatusListResponse, err error) {
    return c.DescribeAlarmStatusListWithContext(context.Background(), request)
}

// DescribeAlarmStatusList
// 用来查询系统中的告警状态列表
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmStatusListWithContext(ctx context.Context, request *DescribeAlarmStatusListRequest) (response *DescribeAlarmStatusListResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmStatusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmStatusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmTypeListRequest() (request *DescribeAlarmTypeListRequest) {
    request = &DescribeAlarmTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeAlarmTypeList")
    
    
    return
}

func NewDescribeAlarmTypeListResponse() (response *DescribeAlarmTypeListResponse) {
    response = &DescribeAlarmTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmTypeList
// 告警类型获取
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmTypeList(request *DescribeAlarmTypeListRequest) (response *DescribeAlarmTypeListResponse, err error) {
    return c.DescribeAlarmTypeListWithContext(context.Background(), request)
}

// DescribeAlarmTypeList
// 告警类型获取
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAlarmTypeListWithContext(ctx context.Context, request *DescribeAlarmTypeListRequest) (response *DescribeAlarmTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmTypeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmTypeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationListRequest() (request *DescribeApplicationListRequest) {
    request = &DescribeApplicationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeApplicationList")
    
    
    return
}

func NewDescribeApplicationListResponse() (response *DescribeApplicationListResponse) {
    response = &DescribeApplicationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationList
// 查询指定空间关联的应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
func (c *Client) DescribeApplicationList(request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    return c.DescribeApplicationListWithContext(context.Background(), request)
}

// DescribeApplicationList
// 查询指定空间关联的应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPLICATIONID = "InvalidParameterValue.InvalidApplicationId"
func (c *Client) DescribeApplicationListWithContext(ctx context.Context, request *DescribeApplicationListRequest) (response *DescribeApplicationListResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBuildingListRequest() (request *DescribeBuildingListRequest) {
    request = &DescribeBuildingListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeBuildingList")
    
    
    return
}

func NewDescribeBuildingListResponse() (response *DescribeBuildingListResponse) {
    response = &DescribeBuildingListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBuildingList
// 查询建筑列表
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INTERNALERROR_APIREQUESTPATHPARAMETERERROR = "InternalError.ApiRequestPathParameterError"
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INTERNALERROR_BUSINESSLOGICERROR = "InternalError.BusinessLogicError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeBuildingList(request *DescribeBuildingListRequest) (response *DescribeBuildingListResponse, err error) {
    return c.DescribeBuildingListWithContext(context.Background(), request)
}

// DescribeBuildingList
// 查询建筑列表
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INTERNALERROR_APIREQUESTPATHPARAMETERERROR = "InternalError.ApiRequestPathParameterError"
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INTERNALERROR_BUSINESSLOGICERROR = "InternalError.BusinessLogicError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeBuildingListWithContext(ctx context.Context, request *DescribeBuildingListRequest) (response *DescribeBuildingListResponse, err error) {
    if request == nil {
        request = NewDescribeBuildingListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBuildingList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBuildingListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBuildingModelRequest() (request *DescribeBuildingModelRequest) {
    request = &DescribeBuildingModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeBuildingModel")
    
    
    return
}

func NewDescribeBuildingModelResponse() (response *DescribeBuildingModelResponse) {
    response = &DescribeBuildingModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBuildingModel
// 查询建筑三维模型
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeBuildingModel(request *DescribeBuildingModelRequest) (response *DescribeBuildingModelResponse, err error) {
    return c.DescribeBuildingModelWithContext(context.Background(), request)
}

// DescribeBuildingModel
// 查询建筑三维模型
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeBuildingModelWithContext(ctx context.Context, request *DescribeBuildingModelRequest) (response *DescribeBuildingModelResponse, err error) {
    if request == nil {
        request = NewDescribeBuildingModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBuildingModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBuildingModelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBuildingProfileRequest() (request *DescribeBuildingProfileRequest) {
    request = &DescribeBuildingProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeBuildingProfile")
    
    
    return
}

func NewDescribeBuildingProfileResponse() (response *DescribeBuildingProfileResponse) {
    response = &DescribeBuildingProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBuildingProfile
// 查询建筑信息
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeBuildingProfile(request *DescribeBuildingProfileRequest) (response *DescribeBuildingProfileResponse, err error) {
    return c.DescribeBuildingProfileWithContext(context.Background(), request)
}

// DescribeBuildingProfile
// 查询建筑信息
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeBuildingProfileWithContext(ctx context.Context, request *DescribeBuildingProfileRequest) (response *DescribeBuildingProfileResponse, err error) {
    if request == nil {
        request = NewDescribeBuildingProfileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBuildingProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBuildingProfileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCameraExtendInfoRequest() (request *DescribeCameraExtendInfoRequest) {
    request = &DescribeCameraExtendInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeCameraExtendInfo")
    
    
    return
}

func NewDescribeCameraExtendInfoResponse() (response *DescribeCameraExtendInfoResponse) {
    response = &DescribeCameraExtendInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCameraExtendInfo
// 获取视频扩展信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERAINFONOTEXIST = "ResourceNotFound.CameraInfoNotExist"
func (c *Client) DescribeCameraExtendInfo(request *DescribeCameraExtendInfoRequest) (response *DescribeCameraExtendInfoResponse, err error) {
    return c.DescribeCameraExtendInfoWithContext(context.Background(), request)
}

// DescribeCameraExtendInfo
// 获取视频扩展信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERAINFONOTEXIST = "ResourceNotFound.CameraInfoNotExist"
func (c *Client) DescribeCameraExtendInfoWithContext(ctx context.Context, request *DescribeCameraExtendInfoRequest) (response *DescribeCameraExtendInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCameraExtendInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCameraExtendInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCameraExtendInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCityWorkspaceListRequest() (request *DescribeCityWorkspaceListRequest) {
    request = &DescribeCityWorkspaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeCityWorkspaceList")
    
    
    return
}

func NewDescribeCityWorkspaceListResponse() (response *DescribeCityWorkspaceListResponse) {
    response = &DescribeCityWorkspaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCityWorkspaceList
// 通过城市id查询工作空间列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYADMINISTRATIVECODE = "InvalidParameterValue.EmptyAdministrativeCode"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCityWorkspaceList(request *DescribeCityWorkspaceListRequest) (response *DescribeCityWorkspaceListResponse, err error) {
    return c.DescribeCityWorkspaceListWithContext(context.Background(), request)
}

// DescribeCityWorkspaceList
// 通过城市id查询工作空间列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYADMINISTRATIVECODE = "InvalidParameterValue.EmptyAdministrativeCode"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCityWorkspaceListWithContext(ctx context.Context, request *DescribeCityWorkspaceListRequest) (response *DescribeCityWorkspaceListResponse, err error) {
    if request == nil {
        request = NewDescribeCityWorkspaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCityWorkspaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCityWorkspaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceListRequest() (request *DescribeDeviceListRequest) {
    request = &DescribeDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeDeviceList")
    
    
    return
}

func NewDescribeDeviceListResponse() (response *DescribeDeviceListResponse) {
    response = &DescribeDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceList
// 设备列表查询/单个查询（支持通过筛选条件查询，设备类型、标签、PID、空间）
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INTERNALERROR_BUSINESSLOGICERROR = "InternalError.BusinessLogicError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeDeviceList(request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    return c.DescribeDeviceListWithContext(context.Background(), request)
}

// DescribeDeviceList
// 设备列表查询/单个查询（支持通过筛选条件查询，设备类型、标签、PID、空间）
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INTERNALERROR_BUSINESSLOGICERROR = "InternalError.BusinessLogicError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeDeviceListWithContext(ctx context.Context, request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceShadowListRequest() (request *DescribeDeviceShadowListRequest) {
    request = &DescribeDeviceShadowListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeDeviceShadowList")
    
    
    return
}

func NewDescribeDeviceShadowListResponse() (response *DescribeDeviceShadowListResponse) {
    response = &DescribeDeviceShadowListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceShadowList
// 获取设备影子数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeDeviceShadowList(request *DescribeDeviceShadowListRequest) (response *DescribeDeviceShadowListResponse, err error) {
    return c.DescribeDeviceShadowListWithContext(context.Background(), request)
}

// DescribeDeviceShadowList
// 获取设备影子数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeDeviceShadowListWithContext(ctx context.Context, request *DescribeDeviceShadowListRequest) (response *DescribeDeviceShadowListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceShadowListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceShadowList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceShadowListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceStatusListRequest() (request *DescribeDeviceStatusListRequest) {
    request = &DescribeDeviceStatusListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeDeviceStatusList")
    
    
    return
}

func NewDescribeDeviceStatusListResponse() (response *DescribeDeviceStatusListResponse) {
    response = &DescribeDeviceStatusListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceStatusList
// 设备状态获取
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeDeviceStatusList(request *DescribeDeviceStatusListRequest) (response *DescribeDeviceStatusListResponse, err error) {
    return c.DescribeDeviceStatusListWithContext(context.Background(), request)
}

// DescribeDeviceStatusList
// 设备状态获取
//
// 可能返回的错误码:
//  FAILEDOPERATION_APIRULECONFIGERROR = "FailedOperation.ApiRuleConfigError"
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDAPI = "UnauthorizedOperation.UnauthorizedApi"
func (c *Client) DescribeDeviceStatusListWithContext(ctx context.Context, request *DescribeDeviceStatusListRequest) (response *DescribeDeviceStatusListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceStatusListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceStatusList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceStatusListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceStatusStatRequest() (request *DescribeDeviceStatusStatRequest) {
    request = &DescribeDeviceStatusStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeDeviceStatusStat")
    
    
    return
}

func NewDescribeDeviceStatusStatResponse() (response *DescribeDeviceStatusStatResponse) {
    response = &DescribeDeviceStatusStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceStatusStat
// 设备状态统计
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeDeviceStatusStat(request *DescribeDeviceStatusStatRequest) (response *DescribeDeviceStatusStatResponse, err error) {
    return c.DescribeDeviceStatusStatWithContext(context.Background(), request)
}

// DescribeDeviceStatusStat
// 设备状态统计
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDISOPERATIONFAILED = "FailedOperation.RedisOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeDeviceStatusStatWithContext(ctx context.Context, request *DescribeDeviceStatusStatRequest) (response *DescribeDeviceStatusStatResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceStatusStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceStatusStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceStatusStatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceTagListRequest() (request *DescribeDeviceTagListRequest) {
    request = &DescribeDeviceTagListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeDeviceTagList")
    
    
    return
}

func NewDescribeDeviceTagListResponse() (response *DescribeDeviceTagListResponse) {
    response = &DescribeDeviceTagListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceTagList
// 标签列表查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETTAGSFAILED = "FailedOperation.GetTagsFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeDeviceTagList(request *DescribeDeviceTagListRequest) (response *DescribeDeviceTagListResponse, err error) {
    return c.DescribeDeviceTagListWithContext(context.Background(), request)
}

// DescribeDeviceTagList
// 标签列表查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETTAGSFAILED = "FailedOperation.GetTagsFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeDeviceTagListWithContext(ctx context.Context, request *DescribeDeviceTagListRequest) (response *DescribeDeviceTagListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceTagListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceTagList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceTagListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceTypeListRequest() (request *DescribeDeviceTypeListRequest) {
    request = &DescribeDeviceTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeDeviceTypeList")
    
    
    return
}

func NewDescribeDeviceTypeListResponse() (response *DescribeDeviceTypeListResponse) {
    response = &DescribeDeviceTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceTypeList
// 拉取设备的设备类型列表
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INVALIDPARAMETER_TOKENFIELDNOTFOUND = "InvalidParameter.TokenFieldNotFound"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeDeviceTypeList(request *DescribeDeviceTypeListRequest) (response *DescribeDeviceTypeListResponse, err error) {
    return c.DescribeDeviceTypeListWithContext(context.Background(), request)
}

// DescribeDeviceTypeList
// 拉取设备的设备类型列表
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INVALIDPARAMETER_TOKENFIELDNOTFOUND = "InvalidParameter.TokenFieldNotFound"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeDeviceTypeListWithContext(ctx context.Context, request *DescribeDeviceTypeListRequest) (response *DescribeDeviceTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceTypeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceTypeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEdgeApplicationTokenRequest() (request *DescribeEdgeApplicationTokenRequest) {
    request = &DescribeEdgeApplicationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeEdgeApplicationToken")
    
    
    return
}

func NewDescribeEdgeApplicationTokenResponse() (response *DescribeEdgeApplicationTokenResponse) {
    response = &DescribeEdgeApplicationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEdgeApplicationToken
// 查询边缘应用凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEdgeApplicationToken(request *DescribeEdgeApplicationTokenRequest) (response *DescribeEdgeApplicationTokenResponse, err error) {
    return c.DescribeEdgeApplicationTokenWithContext(context.Background(), request)
}

// DescribeEdgeApplicationToken
// 查询边缘应用凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEdgeApplicationTokenWithContext(ctx context.Context, request *DescribeEdgeApplicationTokenRequest) (response *DescribeEdgeApplicationTokenResponse, err error) {
    if request == nil {
        request = NewDescribeEdgeApplicationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEdgeApplicationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEdgeApplicationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeElementProfilePageRequest() (request *DescribeElementProfilePageRequest) {
    request = &DescribeElementProfilePageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeElementProfilePage")
    
    
    return
}

func NewDescribeElementProfilePageResponse() (response *DescribeElementProfilePageResponse) {
    response = &DescribeElementProfilePageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeElementProfilePage
// 查询分页构件信息
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeElementProfilePage(request *DescribeElementProfilePageRequest) (response *DescribeElementProfilePageResponse, err error) {
    return c.DescribeElementProfilePageWithContext(context.Background(), request)
}

// DescribeElementProfilePage
// 查询分页构件信息
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeElementProfilePageWithContext(ctx context.Context, request *DescribeElementProfilePageRequest) (response *DescribeElementProfilePageResponse, err error) {
    if request == nil {
        request = NewDescribeElementProfilePageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeElementProfilePage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeElementProfilePageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeElementProfileTreeRequest() (request *DescribeElementProfileTreeRequest) {
    request = &DescribeElementProfileTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeElementProfileTree")
    
    
    return
}

func NewDescribeElementProfileTreeResponse() (response *DescribeElementProfileTreeResponse) {
    response = &DescribeElementProfileTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeElementProfileTree
// 查询构件树
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeElementProfileTree(request *DescribeElementProfileTreeRequest) (response *DescribeElementProfileTreeResponse, err error) {
    return c.DescribeElementProfileTreeWithContext(context.Background(), request)
}

// DescribeElementProfileTree
// 查询构件树
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeElementProfileTreeWithContext(ctx context.Context, request *DescribeElementProfileTreeRequest) (response *DescribeElementProfileTreeResponse, err error) {
    if request == nil {
        request = NewDescribeElementProfileTreeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeElementProfileTree require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeElementProfileTreeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventListRequest() (request *DescribeEventListRequest) {
    request = &DescribeEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeEventList")
    
    
    return
}

func NewDescribeEventListResponse() (response *DescribeEventListResponse) {
    response = &DescribeEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventList
// 事件列表查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEventList(request *DescribeEventListRequest) (response *DescribeEventListResponse, err error) {
    return c.DescribeEventListWithContext(context.Background(), request)
}

// DescribeEventList
// 事件列表查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEventListWithContext(ctx context.Context, request *DescribeEventListRequest) (response *DescribeEventListResponse, err error) {
    if request == nil {
        request = NewDescribeEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileDownloadURLRequest() (request *DescribeFileDownloadURLRequest) {
    request = &DescribeFileDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeFileDownloadURL")
    
    
    return
}

func NewDescribeFileDownloadURLResponse() (response *DescribeFileDownloadURLResponse) {
    response = &DescribeFileDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileDownloadURL
// 获取文件下载URL
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_STORAGEINITFAILED = "FailedOperation.StorageInitFailed"
//  FAILEDOPERATION_URLGENERATEFAILED = "FailedOperation.URLGenerateFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeFileDownloadURL(request *DescribeFileDownloadURLRequest) (response *DescribeFileDownloadURLResponse, err error) {
    return c.DescribeFileDownloadURLWithContext(context.Background(), request)
}

// DescribeFileDownloadURL
// 获取文件下载URL
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_STORAGEINITFAILED = "FailedOperation.StorageInitFailed"
//  FAILEDOPERATION_URLGENERATEFAILED = "FailedOperation.URLGenerateFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeFileDownloadURLWithContext(ctx context.Context, request *DescribeFileDownloadURLRequest) (response *DescribeFileDownloadURLResponse, err error) {
    if request == nil {
        request = NewDescribeFileDownloadURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileDownloadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileUploadURLRequest() (request *DescribeFileUploadURLRequest) {
    request = &DescribeFileUploadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeFileUploadURL")
    
    
    return
}

func NewDescribeFileUploadURLResponse() (response *DescribeFileUploadURLResponse) {
    response = &DescribeFileUploadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileUploadURL
// 文件上传接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_STORAGEINITFAILED = "FailedOperation.StorageInitFailed"
//  FAILEDOPERATION_URLGENERATEFAILED = "FailedOperation.URLGenerateFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeFileUploadURL(request *DescribeFileUploadURLRequest) (response *DescribeFileUploadURLResponse, err error) {
    return c.DescribeFileUploadURLWithContext(context.Background(), request)
}

// DescribeFileUploadURL
// 文件上传接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_STORAGEINITFAILED = "FailedOperation.StorageInitFailed"
//  FAILEDOPERATION_URLGENERATEFAILED = "FailedOperation.URLGenerateFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeFileUploadURLWithContext(ctx context.Context, request *DescribeFileUploadURLRequest) (response *DescribeFileUploadURLResponse, err error) {
    if request == nil {
        request = NewDescribeFileUploadURLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileUploadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileUploadURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInterfaceListRequest() (request *DescribeInterfaceListRequest) {
    request = &DescribeInterfaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeInterfaceList")
    
    
    return
}

func NewDescribeInterfaceListResponse() (response *DescribeInterfaceListResponse) {
    response = &DescribeInterfaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInterfaceList
// 查询接口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInterfaceList(request *DescribeInterfaceListRequest) (response *DescribeInterfaceListResponse, err error) {
    return c.DescribeInterfaceListWithContext(context.Background(), request)
}

// DescribeInterfaceList
// 查询接口列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInterfaceListWithContext(ctx context.Context, request *DescribeInterfaceListRequest) (response *DescribeInterfaceListResponse, err error) {
    if request == nil {
        request = NewDescribeInterfaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInterfaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInterfaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLinkRuleListRequest() (request *DescribeLinkRuleListRequest) {
    request = &DescribeLinkRuleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeLinkRuleList")
    
    
    return
}

func NewDescribeLinkRuleListResponse() (response *DescribeLinkRuleListResponse) {
    response = &DescribeLinkRuleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLinkRuleList
// 联动规则列表查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLinkRuleList(request *DescribeLinkRuleListRequest) (response *DescribeLinkRuleListResponse, err error) {
    return c.DescribeLinkRuleListWithContext(context.Background(), request)
}

// DescribeLinkRuleList
// 联动规则列表查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLinkRuleListWithContext(ctx context.Context, request *DescribeLinkRuleListRequest) (response *DescribeLinkRuleListResponse, err error) {
    if request == nil {
        request = NewDescribeLinkRuleListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLinkRuleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLinkRuleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelListRequest() (request *DescribeModelListRequest) {
    request = &DescribeModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeModelList")
    
    
    return
}

func NewDescribeModelListResponse() (response *DescribeModelListResponse) {
    response = &DescribeModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelList
// 模型列表查询/单个查询（产品模型/标准模型）
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENEXPIRED = "AuthFailure.TokenExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelList(request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    return c.DescribeModelListWithContext(context.Background(), request)
}

// DescribeModelList
// 模型列表查询/单个查询（产品模型/标准模型）
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENEXPIRED = "AuthFailure.TokenExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelListWithContext(ctx context.Context, request *DescribeModelListRequest) (response *DescribeModelListResponse, err error) {
    if request == nil {
        request = NewDescribeModelListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductListRequest() (request *DescribeProductListRequest) {
    request = &DescribeProductListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeProductList")
    
    
    return
}

func NewDescribeProductListResponse() (response *DescribeProductListResponse) {
    response = &DescribeProductListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductList
// 产品列表查询
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProductList(request *DescribeProductListRequest) (response *DescribeProductListResponse, err error) {
    return c.DescribeProductListWithContext(context.Background(), request)
}

// DescribeProductList
// 产品列表查询
//
// 可能返回的错误码:
//  AUTHFAILURE_APIAUTHORIZATIONNOTFOUND = "AuthFailure.ApiAuthorizationNotFound"
//  INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeProductListWithContext(ctx context.Context, request *DescribeProductListRequest) (response *DescribeProductListResponse, err error) {
    if request == nil {
        request = NewDescribeProductListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePropertyListRequest() (request *DescribePropertyListRequest) {
    request = &DescribePropertyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribePropertyList")
    
    
    return
}

func NewDescribePropertyListResponse() (response *DescribePropertyListResponse) {
    response = &DescribePropertyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePropertyList
// 查询构件属性（详情）
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribePropertyList(request *DescribePropertyListRequest) (response *DescribePropertyListResponse, err error) {
    return c.DescribePropertyListWithContext(context.Background(), request)
}

// DescribePropertyList
// 查询构件属性（详情）
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribePropertyListWithContext(ctx context.Context, request *DescribePropertyListRequest) (response *DescribePropertyListResponse, err error) {
    if request == nil {
        request = NewDescribePropertyListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePropertyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePropertyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleDetailRequest() (request *DescribeRuleDetailRequest) {
    request = &DescribeRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeRuleDetail")
    
    
    return
}

func NewDescribeRuleDetailResponse() (response *DescribeRuleDetailResponse) {
    response = &DescribeRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuleDetail
// 联动规则详情查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMNOTMATCH = "InvalidParameter.ParamNotMatch"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRuleDetail(request *DescribeRuleDetailRequest) (response *DescribeRuleDetailResponse, err error) {
    return c.DescribeRuleDetailWithContext(context.Background(), request)
}

// DescribeRuleDetail
// 联动规则详情查询
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMNOTMATCH = "InvalidParameter.ParamNotMatch"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRuleDetailWithContext(ctx context.Context, request *DescribeRuleDetailRequest) (response *DescribeRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRuleDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSceneListRequest() (request *DescribeSceneListRequest) {
    request = &DescribeSceneListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeSceneList")
    
    
    return
}

func NewDescribeSceneListResponse() (response *DescribeSceneListResponse) {
    response = &DescribeSceneListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSceneList
// 查询场景列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSceneList(request *DescribeSceneListRequest) (response *DescribeSceneListResponse, err error) {
    return c.DescribeSceneListWithContext(context.Background(), request)
}

// DescribeSceneList
// 查询场景列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSceneListWithContext(ctx context.Context, request *DescribeSceneListRequest) (response *DescribeSceneListResponse, err error) {
    if request == nil {
        request = NewDescribeSceneListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSceneList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSceneListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceDeviceIdListRequest() (request *DescribeSpaceDeviceIdListRequest) {
    request = &DescribeSpaceDeviceIdListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeSpaceDeviceIdList")
    
    
    return
}

func NewDescribeSpaceDeviceIdListResponse() (response *DescribeSpaceDeviceIdListResponse) {
    response = &DescribeSpaceDeviceIdListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceDeviceIdList
// 查询指定空间设备编号列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceDeviceIdList(request *DescribeSpaceDeviceIdListRequest) (response *DescribeSpaceDeviceIdListResponse, err error) {
    return c.DescribeSpaceDeviceIdListWithContext(context.Background(), request)
}

// DescribeSpaceDeviceIdList
// 查询指定空间设备编号列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceDeviceIdListWithContext(ctx context.Context, request *DescribeSpaceDeviceIdListRequest) (response *DescribeSpaceDeviceIdListResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceDeviceIdListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceDeviceIdList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceDeviceIdListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceDeviceRelationListRequest() (request *DescribeSpaceDeviceRelationListRequest) {
    request = &DescribeSpaceDeviceRelationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeSpaceDeviceRelationList")
    
    
    return
}

func NewDescribeSpaceDeviceRelationListResponse() (response *DescribeSpaceDeviceRelationListResponse) {
    response = &DescribeSpaceDeviceRelationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceDeviceRelationList
// 查询指定空间下设备与构件绑定关系列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceDeviceRelationList(request *DescribeSpaceDeviceRelationListRequest) (response *DescribeSpaceDeviceRelationListResponse, err error) {
    return c.DescribeSpaceDeviceRelationListWithContext(context.Background(), request)
}

// DescribeSpaceDeviceRelationList
// 查询指定空间下设备与构件绑定关系列表
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceDeviceRelationListWithContext(ctx context.Context, request *DescribeSpaceDeviceRelationListRequest) (response *DescribeSpaceDeviceRelationListResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceDeviceRelationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceDeviceRelationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceDeviceRelationListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceInfoByDeviceIdRequest() (request *DescribeSpaceInfoByDeviceIdRequest) {
    request = &DescribeSpaceInfoByDeviceIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeSpaceInfoByDeviceId")
    
    
    return
}

func NewDescribeSpaceInfoByDeviceIdResponse() (response *DescribeSpaceInfoByDeviceIdResponse) {
    response = &DescribeSpaceInfoByDeviceIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceInfoByDeviceId
// 查询设备绑定的空间信息
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceInfoByDeviceId(request *DescribeSpaceInfoByDeviceIdRequest) (response *DescribeSpaceInfoByDeviceIdResponse, err error) {
    return c.DescribeSpaceInfoByDeviceIdWithContext(context.Background(), request)
}

// DescribeSpaceInfoByDeviceId
// 查询设备绑定的空间信息
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceInfoByDeviceIdWithContext(ctx context.Context, request *DescribeSpaceInfoByDeviceIdRequest) (response *DescribeSpaceInfoByDeviceIdResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceInfoByDeviceIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceInfoByDeviceId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceInfoByDeviceIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceRelationByDeviceIdRequest() (request *DescribeSpaceRelationByDeviceIdRequest) {
    request = &DescribeSpaceRelationByDeviceIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeSpaceRelationByDeviceId")
    
    
    return
}

func NewDescribeSpaceRelationByDeviceIdResponse() (response *DescribeSpaceRelationByDeviceIdResponse) {
    response = &DescribeSpaceRelationByDeviceIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceRelationByDeviceId
// 查询设备绑定的空间层级关系
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceRelationByDeviceId(request *DescribeSpaceRelationByDeviceIdRequest) (response *DescribeSpaceRelationByDeviceIdResponse, err error) {
    return c.DescribeSpaceRelationByDeviceIdWithContext(context.Background(), request)
}

// DescribeSpaceRelationByDeviceId
// 查询设备绑定的空间层级关系
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceRelationByDeviceIdWithContext(ctx context.Context, request *DescribeSpaceRelationByDeviceIdRequest) (response *DescribeSpaceRelationByDeviceIdResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceRelationByDeviceIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceRelationByDeviceId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceRelationByDeviceIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpaceTypeListRequest() (request *DescribeSpaceTypeListRequest) {
    request = &DescribeSpaceTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeSpaceTypeList")
    
    
    return
}

func NewDescribeSpaceTypeListResponse() (response *DescribeSpaceTypeListResponse) {
    response = &DescribeSpaceTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpaceTypeList
// 查询空间分类
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceTypeList(request *DescribeSpaceTypeListRequest) (response *DescribeSpaceTypeListResponse, err error) {
    return c.DescribeSpaceTypeListWithContext(context.Background(), request)
}

// DescribeSpaceTypeList
// 查询空间分类
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeSpaceTypeListWithContext(ctx context.Context, request *DescribeSpaceTypeListRequest) (response *DescribeSpaceTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeSpaceTypeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpaceTypeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpaceTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTenantBuildingCountAndAreaRequest() (request *DescribeTenantBuildingCountAndAreaRequest) {
    request = &DescribeTenantBuildingCountAndAreaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeTenantBuildingCountAndArea")
    
    
    return
}

func NewDescribeTenantBuildingCountAndAreaResponse() (response *DescribeTenantBuildingCountAndAreaResponse) {
    response = &DescribeTenantBuildingCountAndAreaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTenantBuildingCountAndArea
// 查询租户楼栋数量和楼栋建筑面积
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeTenantBuildingCountAndArea(request *DescribeTenantBuildingCountAndAreaRequest) (response *DescribeTenantBuildingCountAndAreaResponse, err error) {
    return c.DescribeTenantBuildingCountAndAreaWithContext(context.Background(), request)
}

// DescribeTenantBuildingCountAndArea
// 查询租户楼栋数量和楼栋建筑面积
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeTenantBuildingCountAndAreaWithContext(ctx context.Context, request *DescribeTenantBuildingCountAndAreaRequest) (response *DescribeTenantBuildingCountAndAreaResponse, err error) {
    if request == nil {
        request = NewDescribeTenantBuildingCountAndAreaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTenantBuildingCountAndArea require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTenantBuildingCountAndAreaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTenantDepartmentListRequest() (request *DescribeTenantDepartmentListRequest) {
    request = &DescribeTenantDepartmentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeTenantDepartmentList")
    
    
    return
}

func NewDescribeTenantDepartmentListResponse() (response *DescribeTenantDepartmentListResponse) {
    response = &DescribeTenantDepartmentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTenantDepartmentList
// 查询租户组织部门列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTenantDepartmentList(request *DescribeTenantDepartmentListRequest) (response *DescribeTenantDepartmentListResponse, err error) {
    return c.DescribeTenantDepartmentListWithContext(context.Background(), request)
}

// DescribeTenantDepartmentList
// 查询租户组织部门列表
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTenantDepartmentListWithContext(ctx context.Context, request *DescribeTenantDepartmentListRequest) (response *DescribeTenantDepartmentListResponse, err error) {
    if request == nil {
        request = NewDescribeTenantDepartmentListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTenantDepartmentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTenantDepartmentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTenantUserListRequest() (request *DescribeTenantUserListRequest) {
    request = &DescribeTenantUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeTenantUserList")
    
    
    return
}

func NewDescribeTenantUserListResponse() (response *DescribeTenantUserListResponse) {
    response = &DescribeTenantUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTenantUserList
// 查询租户人员列表
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTenantUserList(request *DescribeTenantUserListRequest) (response *DescribeTenantUserListResponse, err error) {
    return c.DescribeTenantUserListWithContext(context.Background(), request)
}

// DescribeTenantUserList
// 查询租户人员列表
//
// 可能返回的错误码:
//  AUTHFAILURE_TOKENNOTFOUND = "AuthFailure.TokenNotFound"
//  INTERNALERROR_APIGATEWAYINTERNALERROR = "InternalError.ApiGatewayInternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTenantUserListWithContext(ctx context.Context, request *DescribeTenantUserListRequest) (response *DescribeTenantUserListResponse, err error) {
    if request == nil {
        request = NewDescribeTenantUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTenantUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTenantUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoCloudRecordRequest() (request *DescribeVideoCloudRecordRequest) {
    request = &DescribeVideoCloudRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeVideoCloudRecord")
    
    
    return
}

func NewDescribeVideoCloudRecordResponse() (response *DescribeVideoCloudRecordResponse) {
    response = &DescribeVideoCloudRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoCloudRecord
// 云录像接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDTIMERANGE = "InvalidParameterValue.InvalidTimeRange"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
func (c *Client) DescribeVideoCloudRecord(request *DescribeVideoCloudRecordRequest) (response *DescribeVideoCloudRecordResponse, err error) {
    return c.DescribeVideoCloudRecordWithContext(context.Background(), request)
}

// DescribeVideoCloudRecord
// 云录像接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDTIMERANGE = "InvalidParameterValue.InvalidTimeRange"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
func (c *Client) DescribeVideoCloudRecordWithContext(ctx context.Context, request *DescribeVideoCloudRecordRequest) (response *DescribeVideoCloudRecordResponse, err error) {
    if request == nil {
        request = NewDescribeVideoCloudRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoCloudRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoCloudRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoLiveStreamRequest() (request *DescribeVideoLiveStreamRequest) {
    request = &DescribeVideoLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeVideoLiveStream")
    
    
    return
}

func NewDescribeVideoLiveStreamResponse() (response *DescribeVideoLiveStreamResponse) {
    response = &DescribeVideoLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoLiveStream
// 实时流接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOCKREDISCACHEFAILED = "FailedOperation.LockRedisCacheFailed"
//  FAILEDOPERATION_SENDMSGTOIOTFAILED = "FailedOperation.SendMsgToIOTFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDMETADATA = "InvalidParameter.InvalidMetaData"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDSTREAMID = "InvalidParameterValue.InvalidStreamId"
//  LIMITEXCEEDED_VIDEOSTREAMTHRESHOLDEXCEEDED = "LimitExceeded.VideoStreamThresholdExceeded"
//  LIMITEXCEEDED_VIDEOTRANSCODE = "LimitExceeded.VideoTranscode"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
//  RESOURCENOTFOUND_CAMERASTATUSNOTEXIST = "ResourceNotFound.CameraStatusNotExist"
//  RESOURCENOTFOUND_SRSHOOKSERVICENODE = "ResourceNotFound.SRSHookServiceNode"
//  RESOURCENOTFOUND_SRSNOTEXIST = "ResourceNotFound.SRSNotExist"
//  RESOURCENOTFOUND_VIDEOPUSHSERVICENODE = "ResourceNotFound.VideoPushServiceNode"
//  RESOURCEUNAVAILABLE_DEVICEOFFLINE = "ResourceUnavailable.DeviceOffline"
func (c *Client) DescribeVideoLiveStream(request *DescribeVideoLiveStreamRequest) (response *DescribeVideoLiveStreamResponse, err error) {
    return c.DescribeVideoLiveStreamWithContext(context.Background(), request)
}

// DescribeVideoLiveStream
// 实时流接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOCKREDISCACHEFAILED = "FailedOperation.LockRedisCacheFailed"
//  FAILEDOPERATION_SENDMSGTOIOTFAILED = "FailedOperation.SendMsgToIOTFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDMETADATA = "InvalidParameter.InvalidMetaData"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDSTREAMID = "InvalidParameterValue.InvalidStreamId"
//  LIMITEXCEEDED_VIDEOSTREAMTHRESHOLDEXCEEDED = "LimitExceeded.VideoStreamThresholdExceeded"
//  LIMITEXCEEDED_VIDEOTRANSCODE = "LimitExceeded.VideoTranscode"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
//  RESOURCENOTFOUND_CAMERASTATUSNOTEXIST = "ResourceNotFound.CameraStatusNotExist"
//  RESOURCENOTFOUND_SRSHOOKSERVICENODE = "ResourceNotFound.SRSHookServiceNode"
//  RESOURCENOTFOUND_SRSNOTEXIST = "ResourceNotFound.SRSNotExist"
//  RESOURCENOTFOUND_VIDEOPUSHSERVICENODE = "ResourceNotFound.VideoPushServiceNode"
//  RESOURCEUNAVAILABLE_DEVICEOFFLINE = "ResourceUnavailable.DeviceOffline"
func (c *Client) DescribeVideoLiveStreamWithContext(ctx context.Context, request *DescribeVideoLiveStreamRequest) (response *DescribeVideoLiveStreamResponse, err error) {
    if request == nil {
        request = NewDescribeVideoLiveStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoLiveStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoRecordStreamRequest() (request *DescribeVideoRecordStreamRequest) {
    request = &DescribeVideoRecordStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeVideoRecordStream")
    
    
    return
}

func NewDescribeVideoRecordStreamResponse() (response *DescribeVideoRecordStreamResponse) {
    response = &DescribeVideoRecordStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoRecordStream
// 历史流接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_SENDMSGTOIOTFAILED = "FailedOperation.SendMsgToIOTFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDMEDIADATA = "InvalidParameter.InvalidMediaData"
//  INVALIDPARAMETERVALUE_INVALIDGB28181CONFIG = "InvalidParameterValue.InvalidGB28181Config"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDSAVETYPE = "InvalidParameterValue.InvalidSaveType"
//  INVALIDPARAMETERVALUE_INVALIDTIMERANGE = "InvalidParameterValue.InvalidTimeRange"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOPLAYRATE = "InvalidParameterValue.InvalidVideoPlayRate"
//  LIMITEXCEEDED_VIDEOSTREAMTHRESHOLDEXCEEDED = "LimitExceeded.VideoStreamThresholdExceeded"
//  LIMITEXCEEDED_VIDEOTRANSCODE = "LimitExceeded.VideoTranscode"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
//  RESOURCENOTFOUND_CAMERAINFONOTEXIST = "ResourceNotFound.CameraInfoNotExist"
//  RESOURCENOTFOUND_NVRORCVRCONFIGNOTEXIST = "ResourceNotFound.NVROrCVRConfigNotExist"
//  RESOURCENOTFOUND_SRSHOOKSERVICENODE = "ResourceNotFound.SRSHookServiceNode"
//  RESOURCENOTFOUND_SRSNOTEXIST = "ResourceNotFound.SRSNotExist"
//  RESOURCENOTFOUND_STREAMNOTEXIST = "ResourceNotFound.StreamNotExist"
//  RESOURCENOTFOUND_VIDEOPUSHSERVICENODE = "ResourceNotFound.VideoPushServiceNode"
func (c *Client) DescribeVideoRecordStream(request *DescribeVideoRecordStreamRequest) (response *DescribeVideoRecordStreamResponse, err error) {
    return c.DescribeVideoRecordStreamWithContext(context.Background(), request)
}

// DescribeVideoRecordStream
// 历史流接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_SENDMSGTOIOTFAILED = "FailedOperation.SendMsgToIOTFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDMEDIADATA = "InvalidParameter.InvalidMediaData"
//  INVALIDPARAMETERVALUE_INVALIDGB28181CONFIG = "InvalidParameterValue.InvalidGB28181Config"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDSAVETYPE = "InvalidParameterValue.InvalidSaveType"
//  INVALIDPARAMETERVALUE_INVALIDTIMERANGE = "InvalidParameterValue.InvalidTimeRange"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOPLAYRATE = "InvalidParameterValue.InvalidVideoPlayRate"
//  LIMITEXCEEDED_VIDEOSTREAMTHRESHOLDEXCEEDED = "LimitExceeded.VideoStreamThresholdExceeded"
//  LIMITEXCEEDED_VIDEOTRANSCODE = "LimitExceeded.VideoTranscode"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
//  RESOURCENOTFOUND_CAMERAINFONOTEXIST = "ResourceNotFound.CameraInfoNotExist"
//  RESOURCENOTFOUND_NVRORCVRCONFIGNOTEXIST = "ResourceNotFound.NVROrCVRConfigNotExist"
//  RESOURCENOTFOUND_SRSHOOKSERVICENODE = "ResourceNotFound.SRSHookServiceNode"
//  RESOURCENOTFOUND_SRSNOTEXIST = "ResourceNotFound.SRSNotExist"
//  RESOURCENOTFOUND_STREAMNOTEXIST = "ResourceNotFound.StreamNotExist"
//  RESOURCENOTFOUND_VIDEOPUSHSERVICENODE = "ResourceNotFound.VideoPushServiceNode"
func (c *Client) DescribeVideoRecordStreamWithContext(ctx context.Context, request *DescribeVideoRecordStreamRequest) (response *DescribeVideoRecordStreamResponse, err error) {
    if request == nil {
        request = NewDescribeVideoRecordStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoRecordStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoRecordStreamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkSpaceBuildingCountAndAreaRequest() (request *DescribeWorkSpaceBuildingCountAndAreaRequest) {
    request = &DescribeWorkSpaceBuildingCountAndAreaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeWorkSpaceBuildingCountAndArea")
    
    
    return
}

func NewDescribeWorkSpaceBuildingCountAndAreaResponse() (response *DescribeWorkSpaceBuildingCountAndAreaResponse) {
    response = &DescribeWorkSpaceBuildingCountAndAreaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkSpaceBuildingCountAndArea
// 查询项目空间楼栋数量与建筑面积
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeWorkSpaceBuildingCountAndArea(request *DescribeWorkSpaceBuildingCountAndAreaRequest) (response *DescribeWorkSpaceBuildingCountAndAreaResponse, err error) {
    return c.DescribeWorkSpaceBuildingCountAndAreaWithContext(context.Background(), request)
}

// DescribeWorkSpaceBuildingCountAndArea
// 查询项目空间楼栋数量与建筑面积
//
// 可能返回的错误码:
//  INTERNALERROR_BIMAPPINTERNALERROR = "InternalError.BIMAppInternalError"
//  INVALIDPARAMETERVALUE_INVALIDELEMENTIDORPROJECTID = "InvalidParameterValue.InvalidElementIdOrProjectId"
//  RESOURCENOTFOUND_ELEMENTNOTFOUND = "ResourceNotFound.ElementNotFound"
func (c *Client) DescribeWorkSpaceBuildingCountAndAreaWithContext(ctx context.Context, request *DescribeWorkSpaceBuildingCountAndAreaRequest) (response *DescribeWorkSpaceBuildingCountAndAreaResponse, err error) {
    if request == nil {
        request = NewDescribeWorkSpaceBuildingCountAndAreaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkSpaceBuildingCountAndArea require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkSpaceBuildingCountAndAreaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceListRequest() (request *DescribeWorkspaceListRequest) {
    request = &DescribeWorkspaceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeWorkspaceList")
    
    
    return
}

func NewDescribeWorkspaceListResponse() (response *DescribeWorkspaceListResponse) {
    response = &DescribeWorkspaceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkspaceList
// 获取租户下的空间列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeWorkspaceList(request *DescribeWorkspaceListRequest) (response *DescribeWorkspaceListResponse, err error) {
    return c.DescribeWorkspaceListWithContext(context.Background(), request)
}

// DescribeWorkspaceList
// 获取租户下的空间列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
func (c *Client) DescribeWorkspaceListWithContext(ctx context.Context, request *DescribeWorkspaceListRequest) (response *DescribeWorkspaceListResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkspaceUserListRequest() (request *DescribeWorkspaceUserListRequest) {
    request = &DescribeWorkspaceUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "DescribeWorkspaceUserList")
    
    
    return
}

func NewDescribeWorkspaceUserListResponse() (response *DescribeWorkspaceUserListResponse) {
    response = &DescribeWorkspaceUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkspaceUserList
// 查询项目空间人员列表
//
// 可能返回的错误码:
//  INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWorkspaceUserList(request *DescribeWorkspaceUserListRequest) (response *DescribeWorkspaceUserListResponse, err error) {
    return c.DescribeWorkspaceUserListWithContext(context.Background(), request)
}

// DescribeWorkspaceUserList
// 查询项目空间人员列表
//
// 可能返回的错误码:
//  INTERNALERROR_APPAPINOSPACEPERMISSION = "InternalError.AppApiNoSpacePermission"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeWorkspaceUserListWithContext(ctx context.Context, request *DescribeWorkspaceUserListRequest) (response *DescribeWorkspaceUserListResponse, err error) {
    if request == nil {
        request = NewDescribeWorkspaceUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkspaceUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkspaceUserListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceNameRequest() (request *ModifyDeviceNameRequest) {
    request = &ModifyDeviceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "ModifyDeviceName")
    
    
    return
}

func NewModifyDeviceNameResponse() (response *ModifyDeviceNameResponse) {
    response = &ModifyDeviceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDeviceName
// 批量修改设备名字
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYDEVICEERROR = "FailedOperation.ModifyDeviceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  RESOURCENOTFOUND_WIDNOTEXIST = "ResourceNotFound.WIDNotExist"
func (c *Client) ModifyDeviceName(request *ModifyDeviceNameRequest) (response *ModifyDeviceNameResponse, err error) {
    return c.ModifyDeviceNameWithContext(context.Background(), request)
}

// ModifyDeviceName
// 批量修改设备名字
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYDEVICEERROR = "FailedOperation.ModifyDeviceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND_EMPTYDEVICELIST = "ResourceNotFound.EmptyDeviceList"
//  RESOURCENOTFOUND_WIDNOTEXIST = "ResourceNotFound.WIDNotExist"
func (c *Client) ModifyDeviceNameWithContext(ctx context.Context, request *ModifyDeviceNameRequest) (response *ModifyDeviceNameResponse, err error) {
    if request == nil {
        request = NewModifyDeviceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceNameResponse()
    err = c.Send(request, response)
    return
}

func NewReportAppMessageRequest() (request *ReportAppMessageRequest) {
    request = &ReportAppMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "ReportAppMessage")
    
    
    return
}

func NewReportAppMessageResponse() (response *ReportAppMessageResponse) {
    response = &ReportAppMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReportAppMessage
// 上报应用消息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReportAppMessage(request *ReportAppMessageRequest) (response *ReportAppMessageResponse, err error) {
    return c.ReportAppMessageWithContext(context.Background(), request)
}

// ReportAppMessage
// 上报应用消息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReportAppMessageWithContext(ctx context.Context, request *ReportAppMessageRequest) (response *ReportAppMessageResponse, err error) {
    if request == nil {
        request = NewReportAppMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportAppMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportAppMessageResponse()
    err = c.Send(request, response)
    return
}

func NewStopVideoStreamingRequest() (request *StopVideoStreamingRequest) {
    request = &StopVideoStreamingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "StopVideoStreaming")
    
    
    return
}

func NewStopVideoStreamingResponse() (response *StopVideoStreamingResponse) {
    response = &StopVideoStreamingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopVideoStreaming
// 断流接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_EMPTYSTREAM = "MissingParameter.EmptyStream"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
//  RESOURCENOTFOUND_CAMERASTATUSNOTEXIST = "ResourceNotFound.CameraStatusNotExist"
//  RESOURCEUNAVAILABLE_DEVICEOFFLINE = "ResourceUnavailable.DeviceOffline"
//  UNSUPPORTEDOPERATION_NOTSTREAMING = "UnsupportedOperation.NotStreaming"
func (c *Client) StopVideoStreaming(request *StopVideoStreamingRequest) (response *StopVideoStreamingResponse, err error) {
    return c.StopVideoStreamingWithContext(context.Background(), request)
}

// StopVideoStreaming
// 断流接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER_EMPTYSTREAM = "MissingParameter.EmptyStream"
//  MISSINGPARAMETER_EMPTYWID = "MissingParameter.EmptyWID"
//  RESOURCENOTFOUND_CAMERABASEINFONOTEXIST = "ResourceNotFound.CameraBaseInfoNotExist"
//  RESOURCENOTFOUND_CAMERASTATUSNOTEXIST = "ResourceNotFound.CameraStatusNotExist"
//  RESOURCEUNAVAILABLE_DEVICEOFFLINE = "ResourceUnavailable.DeviceOffline"
//  UNSUPPORTEDOPERATION_NOTSTREAMING = "UnsupportedOperation.NotStreaming"
func (c *Client) StopVideoStreamingWithContext(ctx context.Context, request *StopVideoStreamingRequest) (response *StopVideoStreamingResponse, err error) {
    if request == nil {
        request = NewStopVideoStreamingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopVideoStreaming require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopVideoStreamingResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkspaceParkAttributesRequest() (request *UpdateWorkspaceParkAttributesRequest) {
    request = &UpdateWorkspaceParkAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("weilingwith", APIVersion, "UpdateWorkspaceParkAttributes")
    
    
    return
}

func NewUpdateWorkspaceParkAttributesResponse() (response *UpdateWorkspaceParkAttributesResponse) {
    response = &UpdateWorkspaceParkAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkspaceParkAttributes
// 修改工作空间园区属性
//
// 可能返回的错误码:
//  INTERNALERROR_UPDATEPARKINFOFAILED = "InternalError.UpdateParkInfoFailed"
//  INVALIDPARAMETER_DUPLICATEPARKCODE = "InvalidParameter.DuplicateParkCode"
//  INVALIDPARAMETER_EXCEEDPARKLENGTHLIMIT = "InvalidParameter.ExceedParkLengthLimit"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateWorkspaceParkAttributes(request *UpdateWorkspaceParkAttributesRequest) (response *UpdateWorkspaceParkAttributesResponse, err error) {
    return c.UpdateWorkspaceParkAttributesWithContext(context.Background(), request)
}

// UpdateWorkspaceParkAttributes
// 修改工作空间园区属性
//
// 可能返回的错误码:
//  INTERNALERROR_UPDATEPARKINFOFAILED = "InternalError.UpdateParkInfoFailed"
//  INVALIDPARAMETER_DUPLICATEPARKCODE = "InvalidParameter.DuplicateParkCode"
//  INVALIDPARAMETER_EXCEEDPARKLENGTHLIMIT = "InvalidParameter.ExceedParkLengthLimit"
//  INVALIDPARAMETERVALUE_INVALIDWORKSPACEID = "InvalidParameterValue.InvalidWorkspaceId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateWorkspaceParkAttributesWithContext(ctx context.Context, request *UpdateWorkspaceParkAttributesRequest) (response *UpdateWorkspaceParkAttributesResponse, err error) {
    if request == nil {
        request = NewUpdateWorkspaceParkAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkspaceParkAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkspaceParkAttributesResponse()
    err = c.Send(request, response)
    return
}
