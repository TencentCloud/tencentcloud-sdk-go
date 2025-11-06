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

package v20230110

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-01-10"

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


func NewBatchApplyAccountBaselinesRequest() (request *BatchApplyAccountBaselinesRequest) {
    request = &BatchApplyAccountBaselinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("controlcenter", APIVersion, "BatchApplyAccountBaselines")
    
    
    return
}

func NewBatchApplyAccountBaselinesResponse() (response *BatchApplyAccountBaselinesResponse) {
    response = &BatchApplyAccountBaselinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchApplyAccountBaselines
// 批量对存量账号应用基线
//
// 可能返回的错误码:
//  FAILEDOPERATION_AFREQUIREDITEMNOTSELECT = "FailedOperation.AFRequiredItemNotSelect"
//  FAILEDOPERATION_ACCOUNTFACTORYCONTACTEMAILNOTVERIFY = "FailedOperation.AccountFactoryContactEmailNotVerify"
//  FAILEDOPERATION_ACCOUNTFACTORYCONTACTPHONENOTVERIFY = "FailedOperation.AccountFactoryContactPhoneNotVerify"
//  FAILEDOPERATION_ACCOUNTFACTORYMEMBERUINNUMEXCEED = "FailedOperation.AccountFactoryMemberUinNumExceed"
//  FAILEDOPERATION_ACCOUNTFACTORYTASKISDEPLOYING = "FailedOperation.AccountFactoryTaskIsDeploying"
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DEPENDONITEMNOTDEPLOY = "FailedOperation.DependOnItemNotDeploy"
//  FAILEDOPERATION_REMOTECALLERROR = "FailedOperation.RemoteCallError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYBASELINENOTEXIST = "ResourceNotFound.AccountFactoryBaselineNotExist"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTCONFIG = "ResourceNotFound.AccountFactoryItemNotConfig"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTEXIST = "ResourceNotFound.AccountFactoryItemNotExist"
func (c *Client) BatchApplyAccountBaselines(request *BatchApplyAccountBaselinesRequest) (response *BatchApplyAccountBaselinesResponse, err error) {
    return c.BatchApplyAccountBaselinesWithContext(context.Background(), request)
}

// BatchApplyAccountBaselines
// 批量对存量账号应用基线
//
// 可能返回的错误码:
//  FAILEDOPERATION_AFREQUIREDITEMNOTSELECT = "FailedOperation.AFRequiredItemNotSelect"
//  FAILEDOPERATION_ACCOUNTFACTORYCONTACTEMAILNOTVERIFY = "FailedOperation.AccountFactoryContactEmailNotVerify"
//  FAILEDOPERATION_ACCOUNTFACTORYCONTACTPHONENOTVERIFY = "FailedOperation.AccountFactoryContactPhoneNotVerify"
//  FAILEDOPERATION_ACCOUNTFACTORYMEMBERUINNUMEXCEED = "FailedOperation.AccountFactoryMemberUinNumExceed"
//  FAILEDOPERATION_ACCOUNTFACTORYTASKISDEPLOYING = "FailedOperation.AccountFactoryTaskIsDeploying"
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DEPENDONITEMNOTDEPLOY = "FailedOperation.DependOnItemNotDeploy"
//  FAILEDOPERATION_REMOTECALLERROR = "FailedOperation.RemoteCallError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYBASELINENOTEXIST = "ResourceNotFound.AccountFactoryBaselineNotExist"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTCONFIG = "ResourceNotFound.AccountFactoryItemNotConfig"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTEXIST = "ResourceNotFound.AccountFactoryItemNotExist"
func (c *Client) BatchApplyAccountBaselinesWithContext(ctx context.Context, request *BatchApplyAccountBaselinesRequest) (response *BatchApplyAccountBaselinesResponse, err error) {
    if request == nil {
        request = NewBatchApplyAccountBaselinesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "controlcenter", APIVersion, "BatchApplyAccountBaselines")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchApplyAccountBaselines require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchApplyAccountBaselinesResponse()
    err = c.Send(request, response)
    return
}

func NewGetAccountFactoryBaselineRequest() (request *GetAccountFactoryBaselineRequest) {
    request = &GetAccountFactoryBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("controlcenter", APIVersion, "GetAccountFactoryBaseline")
    
    
    return
}

func NewGetAccountFactoryBaselineResponse() (response *GetAccountFactoryBaselineResponse) {
    response = &GetAccountFactoryBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAccountFactoryBaseline
// 获取用户基线配置数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYBASELINENOTEXIST = "ResourceNotFound.AccountFactoryBaselineNotExist"
func (c *Client) GetAccountFactoryBaseline(request *GetAccountFactoryBaselineRequest) (response *GetAccountFactoryBaselineResponse, err error) {
    return c.GetAccountFactoryBaselineWithContext(context.Background(), request)
}

// GetAccountFactoryBaseline
// 获取用户基线配置数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYBASELINENOTEXIST = "ResourceNotFound.AccountFactoryBaselineNotExist"
func (c *Client) GetAccountFactoryBaselineWithContext(ctx context.Context, request *GetAccountFactoryBaselineRequest) (response *GetAccountFactoryBaselineResponse, err error) {
    if request == nil {
        request = NewGetAccountFactoryBaselineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "controlcenter", APIVersion, "GetAccountFactoryBaseline")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAccountFactoryBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAccountFactoryBaselineResponse()
    err = c.Send(request, response)
    return
}

func NewListAccountFactoryBaselineItemsRequest() (request *ListAccountFactoryBaselineItemsRequest) {
    request = &ListAccountFactoryBaselineItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("controlcenter", APIVersion, "ListAccountFactoryBaselineItems")
    
    
    return
}

func NewListAccountFactoryBaselineItemsResponse() (response *ListAccountFactoryBaselineItemsResponse) {
    response = &ListAccountFactoryBaselineItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAccountFactoryBaselineItems
// 获取账号工厂系统基线项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
func (c *Client) ListAccountFactoryBaselineItems(request *ListAccountFactoryBaselineItemsRequest) (response *ListAccountFactoryBaselineItemsResponse, err error) {
    return c.ListAccountFactoryBaselineItemsWithContext(context.Background(), request)
}

// ListAccountFactoryBaselineItems
// 获取账号工厂系统基线项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
func (c *Client) ListAccountFactoryBaselineItemsWithContext(ctx context.Context, request *ListAccountFactoryBaselineItemsRequest) (response *ListAccountFactoryBaselineItemsResponse, err error) {
    if request == nil {
        request = NewListAccountFactoryBaselineItemsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "controlcenter", APIVersion, "ListAccountFactoryBaselineItems")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAccountFactoryBaselineItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAccountFactoryBaselineItemsResponse()
    err = c.Send(request, response)
    return
}

func NewListDeployStepTasksRequest() (request *ListDeployStepTasksRequest) {
    request = &ListDeployStepTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("controlcenter", APIVersion, "ListDeployStepTasks")
    
    
    return
}

func NewListDeployStepTasksResponse() (response *ListDeployStepTasksResponse) {
    response = &ListDeployStepTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDeployStepTasks
// 获取某个基线项历史应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTEXIST = "ResourceNotFound.AccountFactoryItemNotExist"
func (c *Client) ListDeployStepTasks(request *ListDeployStepTasksRequest) (response *ListDeployStepTasksResponse, err error) {
    return c.ListDeployStepTasksWithContext(context.Background(), request)
}

// ListDeployStepTasks
// 获取某个基线项历史应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  RESOURCENOTFOUND_ACCOUNTFACTORYITEMNOTEXIST = "ResourceNotFound.AccountFactoryItemNotExist"
func (c *Client) ListDeployStepTasksWithContext(ctx context.Context, request *ListDeployStepTasksRequest) (response *ListDeployStepTasksResponse, err error) {
    if request == nil {
        request = NewListDeployStepTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "controlcenter", APIVersion, "ListDeployStepTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDeployStepTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDeployStepTasksResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAccountFactoryBaselineRequest() (request *UpdateAccountFactoryBaselineRequest) {
    request = &UpdateAccountFactoryBaselineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("controlcenter", APIVersion, "UpdateAccountFactoryBaseline")
    
    
    return
}

func NewUpdateAccountFactoryBaselineResponse() (response *UpdateAccountFactoryBaselineResponse) {
    response = &UpdateAccountFactoryBaselineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAccountFactoryBaseline
// 更新用户当前基线项配置，基线配置会覆盖更新为当前配置。新增基线项时需要将新增的基线配置加到现有配置，删除基线项时需要将删除的基线配置从现有配置移除，然后保存最新基线配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AFREQUIREDITEMNOTSELECT = "FailedOperation.AFRequiredItemNotSelect"
//  FAILEDOPERATION_ACCOUNTFACTORYDEPENDONITEMNOTCONFIG = "FailedOperation.AccountFactoryDependOnItemNotConfig"
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INVALIDPARAMETER_ACCOUNTFACTORYTAGEXCEEDMAXNUM = "InvalidParameter.AccountFactoryTagExceedMaxNum"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) UpdateAccountFactoryBaseline(request *UpdateAccountFactoryBaselineRequest) (response *UpdateAccountFactoryBaselineResponse, err error) {
    return c.UpdateAccountFactoryBaselineWithContext(context.Background(), request)
}

// UpdateAccountFactoryBaseline
// 更新用户当前基线项配置，基线配置会覆盖更新为当前配置。新增基线项时需要将新增的基线配置加到现有配置，删除基线项时需要将删除的基线配置从现有配置移除，然后保存最新基线配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AFREQUIREDITEMNOTSELECT = "FailedOperation.AFRequiredItemNotSelect"
//  FAILEDOPERATION_ACCOUNTFACTORYDEPENDONITEMNOTCONFIG = "FailedOperation.AccountFactoryDependOnItemNotConfig"
//  FAILEDOPERATION_CONTROLCENTERNOTOPEN = "FailedOperation.ControlCenterNotOpen"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INVALIDPARAMETER_ACCOUNTFACTORYTAGEXCEEDMAXNUM = "InvalidParameter.AccountFactoryTagExceedMaxNum"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) UpdateAccountFactoryBaselineWithContext(ctx context.Context, request *UpdateAccountFactoryBaselineRequest) (response *UpdateAccountFactoryBaselineResponse, err error) {
    if request == nil {
        request = NewUpdateAccountFactoryBaselineRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "controlcenter", APIVersion, "UpdateAccountFactoryBaseline")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAccountFactoryBaseline require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAccountFactoryBaselineResponse()
    err = c.Send(request, response)
    return
}
