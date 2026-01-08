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

package v20250806

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-08-06"

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


func NewAddCalcEnginesToProjectRequest() (request *AddCalcEnginesToProjectRequest) {
    request = &AddCalcEnginesToProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "AddCalcEnginesToProject")
    
    
    return
}

func NewAddCalcEnginesToProjectResponse() (response *AddCalcEnginesToProjectResponse) {
    response = &AddCalcEnginesToProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCalcEnginesToProject
// 关联项目集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) AddCalcEnginesToProject(request *AddCalcEnginesToProjectRequest) (response *AddCalcEnginesToProjectResponse, err error) {
    return c.AddCalcEnginesToProjectWithContext(context.Background(), request)
}

// AddCalcEnginesToProject
// 关联项目集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) AddCalcEnginesToProjectWithContext(ctx context.Context, request *AddCalcEnginesToProjectRequest) (response *AddCalcEnginesToProjectResponse, err error) {
    if request == nil {
        request = NewAddCalcEnginesToProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "AddCalcEnginesToProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCalcEnginesToProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCalcEnginesToProjectResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateResourceGroupToProjectRequest() (request *AssociateResourceGroupToProjectRequest) {
    request = &AssociateResourceGroupToProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "AssociateResourceGroupToProject")
    
    
    return
}

func NewAssociateResourceGroupToProjectResponse() (response *AssociateResourceGroupToProjectResponse) {
    response = &AssociateResourceGroupToProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateResourceGroupToProject
// 该接口用于将指定执行资源组绑定到项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECUTORCLUSTERSTATUSERROR = "FailedOperation.ExecutorClusterStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AssociateResourceGroupToProject(request *AssociateResourceGroupToProjectRequest) (response *AssociateResourceGroupToProjectResponse, err error) {
    return c.AssociateResourceGroupToProjectWithContext(context.Background(), request)
}

// AssociateResourceGroupToProject
// 该接口用于将指定执行资源组绑定到项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECUTORCLUSTERSTATUSERROR = "FailedOperation.ExecutorClusterStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) AssociateResourceGroupToProjectWithContext(ctx context.Context, request *AssociateResourceGroupToProjectRequest) (response *AssociateResourceGroupToProjectResponse, err error) {
    if request == nil {
        request = NewAssociateResourceGroupToProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "AssociateResourceGroupToProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateResourceGroupToProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateResourceGroupToProjectResponse()
    err = c.Send(request, response)
    return
}

func NewAuthorizeDataSourceRequest() (request *AuthorizeDataSourceRequest) {
    request = &AuthorizeDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "AuthorizeDataSource")
    
    
    return
}

func NewAuthorizeDataSourceResponse() (response *AuthorizeDataSourceResponse) {
    response = &AuthorizeDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AuthorizeDataSource
// 授权数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) AuthorizeDataSource(request *AuthorizeDataSourceRequest) (response *AuthorizeDataSourceResponse, err error) {
    return c.AuthorizeDataSourceWithContext(context.Background(), request)
}

// AuthorizeDataSource
// 授权数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) AuthorizeDataSourceWithContext(ctx context.Context, request *AuthorizeDataSourceRequest) (response *AuthorizeDataSourceResponse, err error) {
    if request == nil {
        request = NewAuthorizeDataSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "AuthorizeDataSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuthorizeDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewAuthorizeDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCodeFileRequest() (request *CreateCodeFileRequest) {
    request = &CreateCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateCodeFile")
    
    
    return
}

func NewCreateCodeFileResponse() (response *CreateCodeFileResponse) {
    response = &CreateCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCodeFile
// 新建代码文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFile(request *CreateCodeFileRequest) (response *CreateCodeFileResponse, err error) {
    return c.CreateCodeFileWithContext(context.Background(), request)
}

// CreateCodeFile
// 新建代码文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFileWithContext(ctx context.Context, request *CreateCodeFileRequest) (response *CreateCodeFileResponse, err error) {
    if request == nil {
        request = NewCreateCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCodeFolderRequest() (request *CreateCodeFolderRequest) {
    request = &CreateCodeFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateCodeFolder")
    
    
    return
}

func NewCreateCodeFolderResponse() (response *CreateCodeFolderResponse) {
    response = &CreateCodeFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCodeFolder
// 新建代码文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFolder(request *CreateCodeFolderRequest) (response *CreateCodeFolderResponse, err error) {
    return c.CreateCodeFolderWithContext(context.Background(), request)
}

// CreateCodeFolder
// 新建代码文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodeFolderWithContext(ctx context.Context, request *CreateCodeFolderRequest) (response *CreateCodeFolderResponse, err error) {
    if request == nil {
        request = NewCreateCodeFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateCodeFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodeFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodeFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCodePermissionsRequest() (request *CreateCodePermissionsRequest) {
    request = &CreateCodePermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateCodePermissions")
    
    
    return
}

func NewCreateCodePermissionsResponse() (response *CreateCodePermissionsResponse) {
    response = &CreateCodePermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCodePermissions
// 配置CodeStudio实体权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodePermissions(request *CreateCodePermissionsRequest) (response *CreateCodePermissionsResponse, err error) {
    return c.CreateCodePermissionsWithContext(context.Background(), request)
}

// CreateCodePermissions
// 配置CodeStudio实体权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateCodePermissionsWithContext(ctx context.Context, request *CreateCodePermissionsRequest) (response *CreateCodePermissionsResponse, err error) {
    if request == nil {
        request = NewCreateCodePermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateCodePermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodePermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodePermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataBackfillPlanRequest() (request *CreateDataBackfillPlanRequest) {
    request = &CreateDataBackfillPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateDataBackfillPlan")
    
    
    return
}

func NewCreateDataBackfillPlanResponse() (response *CreateDataBackfillPlanResponse) {
    response = &CreateDataBackfillPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataBackfillPlan
// 创建数据补录计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDataBackfillPlan(request *CreateDataBackfillPlanRequest) (response *CreateDataBackfillPlanResponse, err error) {
    return c.CreateDataBackfillPlanWithContext(context.Background(), request)
}

// CreateDataBackfillPlan
// 创建数据补录计划
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateDataBackfillPlanWithContext(ctx context.Context, request *CreateDataBackfillPlanRequest) (response *CreateDataBackfillPlanResponse, err error) {
    if request == nil {
        request = NewCreateDataBackfillPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateDataBackfillPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataBackfillPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataBackfillPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataSourceRequest() (request *CreateDataSourceRequest) {
    request = &CreateDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateDataSource")
    
    
    return
}

func NewCreateDataSourceResponse() (response *CreateDataSourceResponse) {
    response = &CreateDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataSource
// 该接口用于在指定项目中创建数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDataSource(request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
    return c.CreateDataSourceWithContext(context.Background(), request)
}

// CreateDataSource
// 该接口用于在指定项目中创建数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateDataSourceWithContext(ctx context.Context, request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
    if request == nil {
        request = NewCreateDataSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateDataSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOpsAlarmRuleRequest() (request *CreateOpsAlarmRuleRequest) {
    request = &CreateOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateOpsAlarmRule")
    
    
    return
}

func NewCreateOpsAlarmRuleResponse() (response *CreateOpsAlarmRuleResponse) {
    response = &CreateOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOpsAlarmRule
// 设置告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) CreateOpsAlarmRule(request *CreateOpsAlarmRuleRequest) (response *CreateOpsAlarmRuleResponse, err error) {
    return c.CreateOpsAlarmRuleWithContext(context.Background(), request)
}

// CreateOpsAlarmRule
// 设置告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) CreateOpsAlarmRuleWithContext(ctx context.Context, request *CreateOpsAlarmRuleRequest) (response *CreateOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewCreateOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProject
// 创建项目，创建时包含集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_QUOTAEXCEEDERROR = "InvalidParameter.QuotaExceedError"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 创建项目，创建时包含集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_QUOTAEXCEEDERROR = "InvalidParameter.QuotaExceedError"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectMemberRequest() (request *CreateProjectMemberRequest) {
    request = &CreateProjectMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateProjectMember")
    
    
    return
}

func NewCreateProjectMemberResponse() (response *CreateProjectMemberResponse) {
    response = &CreateProjectMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProjectMember
// 添加项目用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateProjectMember(request *CreateProjectMemberRequest) (response *CreateProjectMemberResponse, err error) {
    return c.CreateProjectMemberWithContext(context.Background(), request)
}

// CreateProjectMember
// 添加项目用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateProjectMemberWithContext(ctx context.Context, request *CreateProjectMemberRequest) (response *CreateProjectMemberResponse, err error) {
    if request == nil {
        request = NewCreateProjectMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateProjectMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProjectMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQualityRuleRequest() (request *CreateQualityRuleRequest) {
    request = &CreateQualityRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateQualityRule")
    
    
    return
}

func NewCreateQualityRuleResponse() (response *CreateQualityRuleResponse) {
    response = &CreateQualityRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQualityRule
// 创建质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) CreateQualityRule(request *CreateQualityRuleRequest) (response *CreateQualityRuleResponse, err error) {
    return c.CreateQualityRuleWithContext(context.Background(), request)
}

// CreateQualityRule
// 创建质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) CreateQualityRuleWithContext(ctx context.Context, request *CreateQualityRuleRequest) (response *CreateQualityRuleResponse, err error) {
    if request == nil {
        request = NewCreateQualityRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateQualityRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQualityRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQualityRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQualityRuleGroupRequest() (request *CreateQualityRuleGroupRequest) {
    request = &CreateQualityRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateQualityRuleGroup")
    
    
    return
}

func NewCreateQualityRuleGroupResponse() (response *CreateQualityRuleGroupResponse) {
    response = &CreateQualityRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQualityRuleGroup
// 创建监控任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) CreateQualityRuleGroup(request *CreateQualityRuleGroupRequest) (response *CreateQualityRuleGroupResponse, err error) {
    return c.CreateQualityRuleGroupWithContext(context.Background(), request)
}

// CreateQualityRuleGroup
// 创建监控任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"
//  INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"
func (c *Client) CreateQualityRuleGroupWithContext(ctx context.Context, request *CreateQualityRuleGroupRequest) (response *CreateQualityRuleGroupResponse, err error) {
    if request == nil {
        request = NewCreateQualityRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateQualityRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQualityRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQualityRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceFileRequest() (request *CreateResourceFileRequest) {
    request = &CreateResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateResourceFile")
    
    
    return
}

func NewCreateResourceFileResponse() (response *CreateResourceFileResponse) {
    response = &CreateResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourceFile
// 创建资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFile(request *CreateResourceFileRequest) (response *CreateResourceFileResponse, err error) {
    return c.CreateResourceFileWithContext(context.Background(), request)
}

// CreateResourceFile
// 创建资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFileWithContext(ctx context.Context, request *CreateResourceFileRequest) (response *CreateResourceFileResponse, err error) {
    if request == nil {
        request = NewCreateResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceFolderRequest() (request *CreateResourceFolderRequest) {
    request = &CreateResourceFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateResourceFolder")
    
    
    return
}

func NewCreateResourceFolderResponse() (response *CreateResourceFolderResponse) {
    response = &CreateResourceFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourceFolder
// 创建资源文件文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFolder(request *CreateResourceFolderRequest) (response *CreateResourceFolderResponse, err error) {
    return c.CreateResourceFolderWithContext(context.Background(), request)
}

// CreateResourceFolder
// 创建资源文件文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceFolderWithContext(ctx context.Context, request *CreateResourceFolderRequest) (response *CreateResourceFolderResponse, err error) {
    if request == nil {
        request = NewCreateResourceFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateResourceFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourceFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceGroupRequest() (request *CreateResourceGroupRequest) {
    request = &CreateResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateResourceGroup")
    
    
    return
}

func NewCreateResourceGroupResponse() (response *CreateResourceGroupResponse) {
    response = &CreateResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourceGroup
// 该接口用于购买资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (response *CreateResourceGroupResponse, err error) {
    return c.CreateResourceGroupWithContext(context.Background(), request)
}

// CreateResourceGroup
// 该接口用于购买资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateResourceGroupWithContext(ctx context.Context, request *CreateResourceGroupRequest) (response *CreateResourceGroupResponse, err error) {
    if request == nil {
        request = NewCreateResourceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateResourceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSQLFolderRequest() (request *CreateSQLFolderRequest) {
    request = &CreateSQLFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateSQLFolder")
    
    
    return
}

func NewCreateSQLFolderResponse() (response *CreateSQLFolderResponse) {
    response = &CreateSQLFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSQLFolder
// 创建数据探索脚本文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLFolder(request *CreateSQLFolderRequest) (response *CreateSQLFolderResponse, err error) {
    return c.CreateSQLFolderWithContext(context.Background(), request)
}

// CreateSQLFolder
// 创建数据探索脚本文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLFolderWithContext(ctx context.Context, request *CreateSQLFolderRequest) (response *CreateSQLFolderResponse, err error) {
    if request == nil {
        request = NewCreateSQLFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateSQLFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSQLFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSQLFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSQLScriptRequest() (request *CreateSQLScriptRequest) {
    request = &CreateSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateSQLScript")
    
    
    return
}

func NewCreateSQLScriptResponse() (response *CreateSQLScriptResponse) {
    response = &CreateSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSQLScript
// 新增SQL脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLScript(request *CreateSQLScriptRequest) (response *CreateSQLScriptResponse, err error) {
    return c.CreateSQLScriptWithContext(context.Background(), request)
}

// CreateSQLScript
// 新增SQL脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateSQLScriptWithContext(ctx context.Context, request *CreateSQLScriptRequest) (response *CreateSQLScriptResponse, err error) {
    if request == nil {
        request = NewCreateSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
    request = &CreateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTask")
    
    
    return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
    response = &CreateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTask
// 创建任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    return c.CreateTaskWithContext(context.Background(), request)
}

// CreateTask
// 创建任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
    if request == nil {
        request = NewCreateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFolderRequest() (request *CreateTaskFolderRequest) {
    request = &CreateTaskFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTaskFolder")
    
    
    return
}

func NewCreateTaskFolderResponse() (response *CreateTaskFolderResponse) {
    response = &CreateTaskFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskFolder
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskFolder(request *CreateTaskFolderRequest) (response *CreateTaskFolderResponse, err error) {
    return c.CreateTaskFolderWithContext(context.Background(), request)
}

// CreateTaskFolder
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTaskFolderWithContext(ctx context.Context, request *CreateTaskFolderRequest) (response *CreateTaskFolderResponse, err error) {
    if request == nil {
        request = NewCreateTaskFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateTaskFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTriggerTaskRequest() (request *CreateTriggerTaskRequest) {
    request = &CreateTriggerTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTriggerTask")
    
    
    return
}

func NewCreateTriggerTaskResponse() (response *CreateTriggerTaskResponse) {
    response = &CreateTriggerTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTriggerTask
// 创建任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTriggerTask(request *CreateTriggerTaskRequest) (response *CreateTriggerTaskResponse, err error) {
    return c.CreateTriggerTaskWithContext(context.Background(), request)
}

// CreateTriggerTask
// 创建任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateTriggerTaskWithContext(ctx context.Context, request *CreateTriggerTaskRequest) (response *CreateTriggerTaskResponse, err error) {
    if request == nil {
        request = NewCreateTriggerTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateTriggerTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTriggerTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTriggerTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTriggerWorkflowRequest() (request *CreateTriggerWorkflowRequest) {
    request = &CreateTriggerWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateTriggerWorkflow")
    
    
    return
}

func NewCreateTriggerWorkflowResponse() (response *CreateTriggerWorkflowResponse) {
    response = &CreateTriggerWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTriggerWorkflow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateTriggerWorkflow(request *CreateTriggerWorkflowRequest) (response *CreateTriggerWorkflowResponse, err error) {
    return c.CreateTriggerWorkflowWithContext(context.Background(), request)
}

// CreateTriggerWorkflow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateTriggerWorkflowWithContext(ctx context.Context, request *CreateTriggerWorkflowRequest) (response *CreateTriggerWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateTriggerWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateTriggerWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTriggerWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTriggerWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowRequest() (request *CreateWorkflowRequest) {
    request = &CreateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateWorkflow")
    
    
    return
}

func NewCreateWorkflowResponse() (response *CreateWorkflowResponse) {
    response = &CreateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflow(request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    return c.CreateWorkflowWithContext(context.Background(), request)
}

// CreateWorkflow
// 创建工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowWithContext(ctx context.Context, request *CreateWorkflowRequest) (response *CreateWorkflowResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowFolderRequest() (request *CreateWorkflowFolderRequest) {
    request = &CreateWorkflowFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateWorkflowFolder")
    
    
    return
}

func NewCreateWorkflowFolderResponse() (response *CreateWorkflowFolderResponse) {
    response = &CreateWorkflowFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflowFolder
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateWorkflowFolder(request *CreateWorkflowFolderRequest) (response *CreateWorkflowFolderResponse, err error) {
    return c.CreateWorkflowFolderWithContext(context.Background(), request)
}

// CreateWorkflowFolder
// 创建文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateWorkflowFolderWithContext(ctx context.Context, request *CreateWorkflowFolderRequest) (response *CreateWorkflowFolderResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateWorkflowFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflowFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowFolderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWorkflowPermissionsRequest() (request *CreateWorkflowPermissionsRequest) {
    request = &CreateWorkflowPermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "CreateWorkflowPermissions")
    
    
    return
}

func NewCreateWorkflowPermissionsResponse() (response *CreateWorkflowPermissionsResponse) {
    response = &CreateWorkflowPermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkflowPermissions
// 数据开发配置权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowPermissions(request *CreateWorkflowPermissionsRequest) (response *CreateWorkflowPermissionsResponse, err error) {
    return c.CreateWorkflowPermissionsWithContext(context.Background(), request)
}

// CreateWorkflowPermissions
// 数据开发配置权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateWorkflowPermissionsWithContext(ctx context.Context, request *CreateWorkflowPermissionsRequest) (response *CreateWorkflowPermissionsResponse, err error) {
    if request == nil {
        request = NewCreateWorkflowPermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "CreateWorkflowPermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkflowPermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkflowPermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCodeFileRequest() (request *DeleteCodeFileRequest) {
    request = &DeleteCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteCodeFile")
    
    
    return
}

func NewDeleteCodeFileResponse() (response *DeleteCodeFileResponse) {
    response = &DeleteCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCodeFile
// 删除代码文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFile(request *DeleteCodeFileRequest) (response *DeleteCodeFileResponse, err error) {
    return c.DeleteCodeFileWithContext(context.Background(), request)
}

// DeleteCodeFile
// 删除代码文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFileWithContext(ctx context.Context, request *DeleteCodeFileRequest) (response *DeleteCodeFileResponse, err error) {
    if request == nil {
        request = NewDeleteCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCodeFolderRequest() (request *DeleteCodeFolderRequest) {
    request = &DeleteCodeFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteCodeFolder")
    
    
    return
}

func NewDeleteCodeFolderResponse() (response *DeleteCodeFolderResponse) {
    response = &DeleteCodeFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCodeFolder
// 数据探索删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFolder(request *DeleteCodeFolderRequest) (response *DeleteCodeFolderResponse, err error) {
    return c.DeleteCodeFolderWithContext(context.Background(), request)
}

// DeleteCodeFolder
// 数据探索删除文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodeFolderWithContext(ctx context.Context, request *DeleteCodeFolderRequest) (response *DeleteCodeFolderResponse, err error) {
    if request == nil {
        request = NewDeleteCodeFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteCodeFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCodeFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCodeFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCodePermissionsRequest() (request *DeleteCodePermissionsRequest) {
    request = &DeleteCodePermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteCodePermissions")
    
    
    return
}

func NewDeleteCodePermissionsResponse() (response *DeleteCodePermissionsResponse) {
    response = &DeleteCodePermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCodePermissions
// 删除CodeStudio实体权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodePermissions(request *DeleteCodePermissionsRequest) (response *DeleteCodePermissionsResponse, err error) {
    return c.DeleteCodePermissionsWithContext(context.Background(), request)
}

// DeleteCodePermissions
// 删除CodeStudio实体权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteCodePermissionsWithContext(ctx context.Context, request *DeleteCodePermissionsRequest) (response *DeleteCodePermissionsResponse, err error) {
    if request == nil {
        request = NewDeleteCodePermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteCodePermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCodePermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCodePermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataBackfillPlanAsyncRequest() (request *DeleteDataBackfillPlanAsyncRequest) {
    request = &DeleteDataBackfillPlanAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteDataBackfillPlanAsync")
    
    
    return
}

func NewDeleteDataBackfillPlanAsyncResponse() (response *DeleteDataBackfillPlanAsyncResponse) {
    response = &DeleteDataBackfillPlanAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataBackfillPlanAsync
// 删除补录计划
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteDataBackfillPlanAsync(request *DeleteDataBackfillPlanAsyncRequest) (response *DeleteDataBackfillPlanAsyncResponse, err error) {
    return c.DeleteDataBackfillPlanAsyncWithContext(context.Background(), request)
}

// DeleteDataBackfillPlanAsync
// 删除补录计划
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteDataBackfillPlanAsyncWithContext(ctx context.Context, request *DeleteDataBackfillPlanAsyncRequest) (response *DeleteDataBackfillPlanAsyncResponse, err error) {
    if request == nil {
        request = NewDeleteDataBackfillPlanAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteDataBackfillPlanAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataBackfillPlanAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataBackfillPlanAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataSourceRequest() (request *DeleteDataSourceRequest) {
    request = &DeleteDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteDataSource")
    
    
    return
}

func NewDeleteDataSourceResponse() (response *DeleteDataSourceResponse) {
    response = &DeleteDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataSource
// 该接口用于删除数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDataSource(request *DeleteDataSourceRequest) (response *DeleteDataSourceResponse, err error) {
    return c.DeleteDataSourceWithContext(context.Background(), request)
}

// DeleteDataSource
// 该接口用于删除数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDataSourceWithContext(ctx context.Context, request *DeleteDataSourceRequest) (response *DeleteDataSourceResponse, err error) {
    if request == nil {
        request = NewDeleteDataSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteDataSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLineageRequest() (request *DeleteLineageRequest) {
    request = &DeleteLineageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteLineage")
    
    
    return
}

func NewDeleteLineageResponse() (response *DeleteLineageResponse) {
    response = &DeleteLineageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLineage
// RegisterLineage
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteLineage(request *DeleteLineageRequest) (response *DeleteLineageResponse, err error) {
    return c.DeleteLineageWithContext(context.Background(), request)
}

// DeleteLineage
// RegisterLineage
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteLineageWithContext(ctx context.Context, request *DeleteLineageRequest) (response *DeleteLineageResponse, err error) {
    if request == nil {
        request = NewDeleteLineageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteLineage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLineage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLineageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOpsAlarmRuleRequest() (request *DeleteOpsAlarmRuleRequest) {
    request = &DeleteOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteOpsAlarmRule")
    
    
    return
}

func NewDeleteOpsAlarmRuleResponse() (response *DeleteOpsAlarmRuleResponse) {
    response = &DeleteOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOpsAlarmRule
// 删除告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteOpsAlarmRule(request *DeleteOpsAlarmRuleRequest) (response *DeleteOpsAlarmRuleResponse, err error) {
    return c.DeleteOpsAlarmRuleWithContext(context.Background(), request)
}

// DeleteOpsAlarmRule
// 删除告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteOpsAlarmRuleWithContext(ctx context.Context, request *DeleteOpsAlarmRuleRequest) (response *DeleteOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewDeleteOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProject
// 删除项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 删除项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectMemberRequest() (request *DeleteProjectMemberRequest) {
    request = &DeleteProjectMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteProjectMember")
    
    
    return
}

func NewDeleteProjectMemberResponse() (response *DeleteProjectMemberResponse) {
    response = &DeleteProjectMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjectMember
// 删除项目用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteProjectMember(request *DeleteProjectMemberRequest) (response *DeleteProjectMemberResponse, err error) {
    return c.DeleteProjectMemberWithContext(context.Background(), request)
}

// DeleteProjectMember
// 删除项目用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteProjectMemberWithContext(ctx context.Context, request *DeleteProjectMemberRequest) (response *DeleteProjectMemberResponse, err error) {
    if request == nil {
        request = NewDeleteProjectMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteProjectMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjectMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQualityRuleRequest() (request *DeleteQualityRuleRequest) {
    request = &DeleteQualityRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteQualityRule")
    
    
    return
}

func NewDeleteQualityRuleResponse() (response *DeleteQualityRuleResponse) {
    response = &DeleteQualityRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQualityRule
// 删除质量规则接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteQualityRule(request *DeleteQualityRuleRequest) (response *DeleteQualityRuleResponse, err error) {
    return c.DeleteQualityRuleWithContext(context.Background(), request)
}

// DeleteQualityRule
// 删除质量规则接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteQualityRuleWithContext(ctx context.Context, request *DeleteQualityRuleRequest) (response *DeleteQualityRuleResponse, err error) {
    if request == nil {
        request = NewDeleteQualityRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteQualityRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQualityRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQualityRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQualityRuleGroupRequest() (request *DeleteQualityRuleGroupRequest) {
    request = &DeleteQualityRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteQualityRuleGroup")
    
    
    return
}

func NewDeleteQualityRuleGroupResponse() (response *DeleteQualityRuleGroupResponse) {
    response = &DeleteQualityRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQualityRuleGroup
// 删除监控任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteQualityRuleGroup(request *DeleteQualityRuleGroupRequest) (response *DeleteQualityRuleGroupResponse, err error) {
    return c.DeleteQualityRuleGroupWithContext(context.Background(), request)
}

// DeleteQualityRuleGroup
// 删除监控任务
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteQualityRuleGroupWithContext(ctx context.Context, request *DeleteQualityRuleGroupRequest) (response *DeleteQualityRuleGroupResponse, err error) {
    if request == nil {
        request = NewDeleteQualityRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteQualityRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQualityRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQualityRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceFileRequest() (request *DeleteResourceFileRequest) {
    request = &DeleteResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResourceFile")
    
    
    return
}

func NewDeleteResourceFileResponse() (response *DeleteResourceFileResponse) {
    response = &DeleteResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceFile
// 资源管理-删除资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFile(request *DeleteResourceFileRequest) (response *DeleteResourceFileResponse, err error) {
    return c.DeleteResourceFileWithContext(context.Background(), request)
}

// DeleteResourceFile
// 资源管理-删除资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFileWithContext(ctx context.Context, request *DeleteResourceFileRequest) (response *DeleteResourceFileResponse, err error) {
    if request == nil {
        request = NewDeleteResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceFolderRequest() (request *DeleteResourceFolderRequest) {
    request = &DeleteResourceFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResourceFolder")
    
    
    return
}

func NewDeleteResourceFolderResponse() (response *DeleteResourceFolderResponse) {
    response = &DeleteResourceFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceFolder
// 删除资源文件文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFolder(request *DeleteResourceFolderRequest) (response *DeleteResourceFolderResponse, err error) {
    return c.DeleteResourceFolderWithContext(context.Background(), request)
}

// DeleteResourceFolder
// 删除资源文件文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceFolderWithContext(ctx context.Context, request *DeleteResourceFolderRequest) (response *DeleteResourceFolderResponse, err error) {
    if request == nil {
        request = NewDeleteResourceFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteResourceFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceGroupRequest() (request *DeleteResourceGroupRequest) {
    request = &DeleteResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteResourceGroup")
    
    
    return
}

func NewDeleteResourceGroupResponse() (response *DeleteResourceGroupResponse) {
    response = &DeleteResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceGroup
// 该接口用于销毁资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    return c.DeleteResourceGroupWithContext(context.Background(), request)
}

// DeleteResourceGroup
// 该接口用于销毁资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteResourceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteResourceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSQLFolderRequest() (request *DeleteSQLFolderRequest) {
    request = &DeleteSQLFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteSQLFolder")
    
    
    return
}

func NewDeleteSQLFolderResponse() (response *DeleteSQLFolderResponse) {
    response = &DeleteSQLFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSQLFolder
// 删除SQL文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLFolder(request *DeleteSQLFolderRequest) (response *DeleteSQLFolderResponse, err error) {
    return c.DeleteSQLFolderWithContext(context.Background(), request)
}

// DeleteSQLFolder
// 删除SQL文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLFolderWithContext(ctx context.Context, request *DeleteSQLFolderRequest) (response *DeleteSQLFolderResponse, err error) {
    if request == nil {
        request = NewDeleteSQLFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteSQLFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSQLFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSQLFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSQLScriptRequest() (request *DeleteSQLScriptRequest) {
    request = &DeleteSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteSQLScript")
    
    
    return
}

func NewDeleteSQLScriptResponse() (response *DeleteSQLScriptResponse) {
    response = &DeleteSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSQLScript
// 删除探索脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLScript(request *DeleteSQLScriptRequest) (response *DeleteSQLScriptResponse, err error) {
    return c.DeleteSQLScriptWithContext(context.Background(), request)
}

// DeleteSQLScript
// 删除探索脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteSQLScriptWithContext(ctx context.Context, request *DeleteSQLScriptRequest) (response *DeleteSQLScriptResponse, err error) {
    if request == nil {
        request = NewDeleteSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteTask")
    
    
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTask
// 删除编排空间任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    return c.DeleteTaskWithContext(context.Background(), request)
}

// DeleteTask
// 删除编排空间任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteTaskWithContext(ctx context.Context, request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskFolderRequest() (request *DeleteTaskFolderRequest) {
    request = &DeleteTaskFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteTaskFolder")
    
    
    return
}

func NewDeleteTaskFolderResponse() (response *DeleteTaskFolderResponse) {
    response = &DeleteTaskFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTaskFolder
// 删除数据开发任务文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTaskFolder(request *DeleteTaskFolderRequest) (response *DeleteTaskFolderResponse, err error) {
    return c.DeleteTaskFolderWithContext(context.Background(), request)
}

// DeleteTaskFolder
// 删除数据开发任务文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTaskFolderWithContext(ctx context.Context, request *DeleteTaskFolderRequest) (response *DeleteTaskFolderResponse, err error) {
    if request == nil {
        request = NewDeleteTaskFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteTaskFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTaskFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTriggerTaskRequest() (request *DeleteTriggerTaskRequest) {
    request = &DeleteTriggerTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteTriggerTask")
    
    
    return
}

func NewDeleteTriggerTaskResponse() (response *DeleteTriggerTaskResponse) {
    response = &DeleteTriggerTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTriggerTask
// 删除工作流调度任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTriggerTask(request *DeleteTriggerTaskRequest) (response *DeleteTriggerTaskResponse, err error) {
    return c.DeleteTriggerTaskWithContext(context.Background(), request)
}

// DeleteTriggerTask
// 删除工作流调度任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteTriggerTaskWithContext(ctx context.Context, request *DeleteTriggerTaskRequest) (response *DeleteTriggerTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTriggerTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteTriggerTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTriggerTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTriggerTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTriggerWorkflowRequest() (request *DeleteTriggerWorkflowRequest) {
    request = &DeleteTriggerWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteTriggerWorkflow")
    
    
    return
}

func NewDeleteTriggerWorkflowResponse() (response *DeleteTriggerWorkflowResponse) {
    response = &DeleteTriggerWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTriggerWorkflow
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteTriggerWorkflow(request *DeleteTriggerWorkflowRequest) (response *DeleteTriggerWorkflowResponse, err error) {
    return c.DeleteTriggerWorkflowWithContext(context.Background(), request)
}

// DeleteTriggerWorkflow
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteTriggerWorkflowWithContext(ctx context.Context, request *DeleteTriggerWorkflowRequest) (response *DeleteTriggerWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteTriggerWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteTriggerWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTriggerWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTriggerWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowRequest() (request *DeleteWorkflowRequest) {
    request = &DeleteWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflow")
    
    
    return
}

func NewDeleteWorkflowResponse() (response *DeleteWorkflowResponse) {
    response = &DeleteWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflow
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    return c.DeleteWorkflowWithContext(context.Background(), request)
}

// DeleteWorkflow
// 删除工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest) (response *DeleteWorkflowResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowFolderRequest() (request *DeleteWorkflowFolderRequest) {
    request = &DeleteWorkflowFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflowFolder")
    
    
    return
}

func NewDeleteWorkflowFolderResponse() (response *DeleteWorkflowFolderResponse) {
    response = &DeleteWorkflowFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflowFolder
// 删除数据开发文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWorkflowFolder(request *DeleteWorkflowFolderRequest) (response *DeleteWorkflowFolderResponse, err error) {
    return c.DeleteWorkflowFolderWithContext(context.Background(), request)
}

// DeleteWorkflowFolder
// 删除数据开发文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteWorkflowFolderWithContext(ctx context.Context, request *DeleteWorkflowFolderRequest) (response *DeleteWorkflowFolderResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteWorkflowFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflowFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowFolderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWorkflowPermissionsRequest() (request *DeleteWorkflowPermissionsRequest) {
    request = &DeleteWorkflowPermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DeleteWorkflowPermissions")
    
    
    return
}

func NewDeleteWorkflowPermissionsResponse() (response *DeleteWorkflowPermissionsResponse) {
    response = &DeleteWorkflowPermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkflowPermissions
// 删除工作流文件夹权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWorkflowPermissions(request *DeleteWorkflowPermissionsRequest) (response *DeleteWorkflowPermissionsResponse, err error) {
    return c.DeleteWorkflowPermissionsWithContext(context.Background(), request)
}

// DeleteWorkflowPermissions
// 删除工作流文件夹权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteWorkflowPermissionsWithContext(ctx context.Context, request *DeleteWorkflowPermissionsRequest) (response *DeleteWorkflowPermissionsResponse, err error) {
    if request == nil {
        request = NewDeleteWorkflowPermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DeleteWorkflowPermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkflowPermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkflowPermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSourceAuthorityRequest() (request *DescribeDataSourceAuthorityRequest) {
    request = &DescribeDataSourceAuthorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DescribeDataSourceAuthority")
    
    
    return
}

func NewDescribeDataSourceAuthorityResponse() (response *DescribeDataSourceAuthorityResponse) {
    response = &DescribeDataSourceAuthorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataSourceAuthority
// 查看数据源权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDataSourceAuthority(request *DescribeDataSourceAuthorityRequest) (response *DescribeDataSourceAuthorityResponse, err error) {
    return c.DescribeDataSourceAuthorityWithContext(context.Background(), request)
}

// DescribeDataSourceAuthority
// 查看数据源权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDataSourceAuthorityWithContext(ctx context.Context, request *DescribeDataSourceAuthorityRequest) (response *DescribeDataSourceAuthorityResponse, err error) {
    if request == nil {
        request = NewDescribeDataSourceAuthorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DescribeDataSourceAuthority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSourceAuthority require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSourceAuthorityResponse()
    err = c.Send(request, response)
    return
}

func NewDisableProjectRequest() (request *DisableProjectRequest) {
    request = &DisableProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DisableProject")
    
    
    return
}

func NewDisableProjectResponse() (response *DisableProjectResponse) {
    response = &DisableProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableProject
// 禁用项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DisableProject(request *DisableProjectRequest) (response *DisableProjectResponse, err error) {
    return c.DisableProjectWithContext(context.Background(), request)
}

// DisableProject
// 禁用项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DisableProjectWithContext(ctx context.Context, request *DisableProjectRequest) (response *DisableProjectResponse, err error) {
    if request == nil {
        request = NewDisableProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DisableProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDissociateResourceGroupFromProjectRequest() (request *DissociateResourceGroupFromProjectRequest) {
    request = &DissociateResourceGroupFromProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "DissociateResourceGroupFromProject")
    
    
    return
}

func NewDissociateResourceGroupFromProjectResponse() (response *DissociateResourceGroupFromProjectResponse) {
    response = &DissociateResourceGroupFromProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DissociateResourceGroupFromProject
// 该接口用于将指定执行资源组解除与项目的绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DissociateResourceGroupFromProject(request *DissociateResourceGroupFromProjectRequest) (response *DissociateResourceGroupFromProjectResponse, err error) {
    return c.DissociateResourceGroupFromProjectWithContext(context.Background(), request)
}

// DissociateResourceGroupFromProject
// 该接口用于将指定执行资源组解除与项目的绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DissociateResourceGroupFromProjectWithContext(ctx context.Context, request *DissociateResourceGroupFromProjectRequest) (response *DissociateResourceGroupFromProjectResponse, err error) {
    if request == nil {
        request = NewDissociateResourceGroupFromProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "DissociateResourceGroupFromProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DissociateResourceGroupFromProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDissociateResourceGroupFromProjectResponse()
    err = c.Send(request, response)
    return
}

func NewEnableProjectRequest() (request *EnableProjectRequest) {
    request = &EnableProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "EnableProject")
    
    
    return
}

func NewEnableProjectResponse() (response *EnableProjectResponse) {
    response = &EnableProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableProject
// 启用项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EnableProject(request *EnableProjectRequest) (response *EnableProjectResponse, err error) {
    return c.EnableProjectWithContext(context.Background(), request)
}

// EnableProject
// 启用项目
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) EnableProjectWithContext(ctx context.Context, request *EnableProjectRequest) (response *EnableProjectResponse, err error) {
    if request == nil {
        request = NewEnableProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "EnableProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableProjectResponse()
    err = c.Send(request, response)
    return
}

func NewGetAlarmMessageRequest() (request *GetAlarmMessageRequest) {
    request = &GetAlarmMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetAlarmMessage")
    
    
    return
}

func NewGetAlarmMessageResponse() (response *GetAlarmMessageResponse) {
    response = &GetAlarmMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAlarmMessage
// 查询告警信息详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAlarmMessage(request *GetAlarmMessageRequest) (response *GetAlarmMessageResponse, err error) {
    return c.GetAlarmMessageWithContext(context.Background(), request)
}

// GetAlarmMessage
// 查询告警信息详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAlarmMessageWithContext(ctx context.Context, request *GetAlarmMessageRequest) (response *GetAlarmMessageResponse, err error) {
    if request == nil {
        request = NewGetAlarmMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetAlarmMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAlarmMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAlarmMessageResponse()
    err = c.Send(request, response)
    return
}

func NewGetCodeFileRequest() (request *GetCodeFileRequest) {
    request = &GetCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetCodeFile")
    
    
    return
}

func NewGetCodeFileResponse() (response *GetCodeFileResponse) {
    response = &GetCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCodeFile
// 查看代码文件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetCodeFile(request *GetCodeFileRequest) (response *GetCodeFileResponse, err error) {
    return c.GetCodeFileWithContext(context.Background(), request)
}

// GetCodeFile
// 查看代码文件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetCodeFileWithContext(ctx context.Context, request *GetCodeFileRequest) (response *GetCodeFileResponse, err error) {
    if request == nil {
        request = NewGetCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetCodeFolderRequest() (request *GetCodeFolderRequest) {
    request = &GetCodeFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetCodeFolder")
    
    
    return
}

func NewGetCodeFolderResponse() (response *GetCodeFolderResponse) {
    response = &GetCodeFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCodeFolder
// 获取sql文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetCodeFolder(request *GetCodeFolderRequest) (response *GetCodeFolderResponse, err error) {
    return c.GetCodeFolderWithContext(context.Background(), request)
}

// GetCodeFolder
// 获取sql文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetCodeFolderWithContext(ctx context.Context, request *GetCodeFolderRequest) (response *GetCodeFolderResponse, err error) {
    if request == nil {
        request = NewGetCodeFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetCodeFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCodeFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCodeFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataBackfillPlanRequest() (request *GetDataBackfillPlanRequest) {
    request = &GetDataBackfillPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetDataBackfillPlan")
    
    
    return
}

func NewGetDataBackfillPlanResponse() (response *GetDataBackfillPlanResponse) {
    response = &GetDataBackfillPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDataBackfillPlan
// 获取补录计划详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetDataBackfillPlan(request *GetDataBackfillPlanRequest) (response *GetDataBackfillPlanResponse, err error) {
    return c.GetDataBackfillPlanWithContext(context.Background(), request)
}

// GetDataBackfillPlan
// 获取补录计划详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetDataBackfillPlanWithContext(ctx context.Context, request *GetDataBackfillPlanRequest) (response *GetDataBackfillPlanResponse, err error) {
    if request == nil {
        request = NewGetDataBackfillPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetDataBackfillPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDataBackfillPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDataBackfillPlanResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataSourceRequest() (request *GetDataSourceRequest) {
    request = &GetDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetDataSource")
    
    
    return
}

func NewGetDataSourceResponse() (response *GetDataSourceResponse) {
    response = &GetDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDataSource
// 该接口用于查看指定数据源的详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) GetDataSource(request *GetDataSourceRequest) (response *GetDataSourceResponse, err error) {
    return c.GetDataSourceWithContext(context.Background(), request)
}

// GetDataSource
// 该接口用于查看指定数据源的详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) GetDataSourceWithContext(ctx context.Context, request *GetDataSourceRequest) (response *GetDataSourceResponse, err error) {
    if request == nil {
        request = NewGetDataSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetDataSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewGetDataSourceRelatedTasksRequest() (request *GetDataSourceRelatedTasksRequest) {
    request = &GetDataSourceRelatedTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetDataSourceRelatedTasks")
    
    
    return
}

func NewGetDataSourceRelatedTasksResponse() (response *GetDataSourceRelatedTasksResponse) {
    response = &GetDataSourceRelatedTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDataSourceRelatedTasks
// 数据源关联任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetDataSourceRelatedTasks(request *GetDataSourceRelatedTasksRequest) (response *GetDataSourceRelatedTasksResponse, err error) {
    return c.GetDataSourceRelatedTasksWithContext(context.Background(), request)
}

// GetDataSourceRelatedTasks
// 数据源关联任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetDataSourceRelatedTasksWithContext(ctx context.Context, request *GetDataSourceRelatedTasksRequest) (response *GetDataSourceRelatedTasksResponse, err error) {
    if request == nil {
        request = NewGetDataSourceRelatedTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetDataSourceRelatedTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDataSourceRelatedTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDataSourceRelatedTasksResponse()
    err = c.Send(request, response)
    return
}

func NewGetMyCodeMaxPermissionRequest() (request *GetMyCodeMaxPermissionRequest) {
    request = &GetMyCodeMaxPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetMyCodeMaxPermission")
    
    
    return
}

func NewGetMyCodeMaxPermissionResponse() (response *GetMyCodeMaxPermissionResponse) {
    response = &GetMyCodeMaxPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMyCodeMaxPermission
// 查看当前用户对CodeStudio实体的最大权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetMyCodeMaxPermission(request *GetMyCodeMaxPermissionRequest) (response *GetMyCodeMaxPermissionResponse, err error) {
    return c.GetMyCodeMaxPermissionWithContext(context.Background(), request)
}

// GetMyCodeMaxPermission
// 查看当前用户对CodeStudio实体的最大权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetMyCodeMaxPermissionWithContext(ctx context.Context, request *GetMyCodeMaxPermissionRequest) (response *GetMyCodeMaxPermissionResponse, err error) {
    if request == nil {
        request = NewGetMyCodeMaxPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetMyCodeMaxPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMyCodeMaxPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMyCodeMaxPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewGetMyWorkflowMaxPermissionRequest() (request *GetMyWorkflowMaxPermissionRequest) {
    request = &GetMyWorkflowMaxPermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetMyWorkflowMaxPermission")
    
    
    return
}

func NewGetMyWorkflowMaxPermissionResponse() (response *GetMyWorkflowMaxPermissionResponse) {
    response = &GetMyWorkflowMaxPermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMyWorkflowMaxPermission
// 查询当前用户对工作流文件夹的递归最大权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetMyWorkflowMaxPermission(request *GetMyWorkflowMaxPermissionRequest) (response *GetMyWorkflowMaxPermissionResponse, err error) {
    return c.GetMyWorkflowMaxPermissionWithContext(context.Background(), request)
}

// GetMyWorkflowMaxPermission
// 查询当前用户对工作流文件夹的递归最大权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetMyWorkflowMaxPermissionWithContext(ctx context.Context, request *GetMyWorkflowMaxPermissionRequest) (response *GetMyWorkflowMaxPermissionResponse, err error) {
    if request == nil {
        request = NewGetMyWorkflowMaxPermissionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetMyWorkflowMaxPermission")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMyWorkflowMaxPermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMyWorkflowMaxPermissionResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsAlarmRuleRequest() (request *GetOpsAlarmRuleRequest) {
    request = &GetOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsAlarmRule")
    
    
    return
}

func NewGetOpsAlarmRuleResponse() (response *GetOpsAlarmRuleResponse) {
    response = &GetOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsAlarmRule
// 根据告警规则id/名称查询单个告警规则信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsAlarmRule(request *GetOpsAlarmRuleRequest) (response *GetOpsAlarmRuleResponse, err error) {
    return c.GetOpsAlarmRuleWithContext(context.Background(), request)
}

// GetOpsAlarmRule
// 根据告警规则id/名称查询单个告警规则信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsAlarmRuleWithContext(ctx context.Context, request *GetOpsAlarmRuleRequest) (response *GetOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewGetOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsAsyncJobRequest() (request *GetOpsAsyncJobRequest) {
    request = &GetOpsAsyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsAsyncJob")
    
    
    return
}

func NewGetOpsAsyncJobResponse() (response *GetOpsAsyncJobResponse) {
    response = &GetOpsAsyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsAsyncJob
// 查询运维中心异步操作详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetOpsAsyncJob(request *GetOpsAsyncJobRequest) (response *GetOpsAsyncJobResponse, err error) {
    return c.GetOpsAsyncJobWithContext(context.Background(), request)
}

// GetOpsAsyncJob
// 查询运维中心异步操作详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetOpsAsyncJobWithContext(ctx context.Context, request *GetOpsAsyncJobRequest) (response *GetOpsAsyncJobResponse, err error) {
    if request == nil {
        request = NewGetOpsAsyncJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsAsyncJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsAsyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsAsyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsTaskRequest() (request *GetOpsTaskRequest) {
    request = &GetOpsTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsTask")
    
    
    return
}

func NewGetOpsTaskResponse() (response *GetOpsTaskResponse) {
    response = &GetOpsTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsTask
// 获取任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetOpsTask(request *GetOpsTaskRequest) (response *GetOpsTaskResponse, err error) {
    return c.GetOpsTaskWithContext(context.Background(), request)
}

// GetOpsTask
// 获取任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetOpsTaskWithContext(ctx context.Context, request *GetOpsTaskRequest) (response *GetOpsTaskResponse, err error) {
    if request == nil {
        request = NewGetOpsTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsTaskCodeRequest() (request *GetOpsTaskCodeRequest) {
    request = &GetOpsTaskCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsTaskCode")
    
    
    return
}

func NewGetOpsTaskCodeResponse() (response *GetOpsTaskCodeResponse) {
    response = &GetOpsTaskCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsTaskCode
// 获取任务代码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetOpsTaskCode(request *GetOpsTaskCodeRequest) (response *GetOpsTaskCodeResponse, err error) {
    return c.GetOpsTaskCodeWithContext(context.Background(), request)
}

// GetOpsTaskCode
// 获取任务代码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetOpsTaskCodeWithContext(ctx context.Context, request *GetOpsTaskCodeRequest) (response *GetOpsTaskCodeResponse, err error) {
    if request == nil {
        request = NewGetOpsTaskCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsTaskCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsTaskCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsTaskCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsTriggerWorkflowRequest() (request *GetOpsTriggerWorkflowRequest) {
    request = &GetOpsTriggerWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsTriggerWorkflow")
    
    
    return
}

func NewGetOpsTriggerWorkflowResponse() (response *GetOpsTriggerWorkflowResponse) {
    response = &GetOpsTriggerWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsTriggerWorkflow
// 查询工作流任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetOpsTriggerWorkflow(request *GetOpsTriggerWorkflowRequest) (response *GetOpsTriggerWorkflowResponse, err error) {
    return c.GetOpsTriggerWorkflowWithContext(context.Background(), request)
}

// GetOpsTriggerWorkflow
// 查询工作流任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetOpsTriggerWorkflowWithContext(ctx context.Context, request *GetOpsTriggerWorkflowRequest) (response *GetOpsTriggerWorkflowResponse, err error) {
    if request == nil {
        request = NewGetOpsTriggerWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsTriggerWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsTriggerWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsTriggerWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetOpsWorkflowRequest() (request *GetOpsWorkflowRequest) {
    request = &GetOpsWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetOpsWorkflow")
    
    
    return
}

func NewGetOpsWorkflowResponse() (response *GetOpsWorkflowResponse) {
    response = &GetOpsWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetOpsWorkflow
// 根据工作流id，获取工作流调度详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsWorkflow(request *GetOpsWorkflowRequest) (response *GetOpsWorkflowResponse, err error) {
    return c.GetOpsWorkflowWithContext(context.Background(), request)
}

// GetOpsWorkflow
// 根据工作流id，获取工作流调度详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetOpsWorkflowWithContext(ctx context.Context, request *GetOpsWorkflowRequest) (response *GetOpsWorkflowResponse, err error) {
    if request == nil {
        request = NewGetOpsWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetOpsWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetOpsWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetOpsWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetProjectRequest() (request *GetProjectRequest) {
    request = &GetProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetProject")
    
    
    return
}

func NewGetProjectResponse() (response *GetProjectResponse) {
    response = &GetProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) GetProject(request *GetProjectRequest) (response *GetProjectResponse, err error) {
    return c.GetProjectWithContext(context.Background(), request)
}

// GetProject
// 获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"
func (c *Client) GetProjectWithContext(ctx context.Context, request *GetProjectRequest) (response *GetProjectResponse, err error) {
    if request == nil {
        request = NewGetProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProjectResponse()
    err = c.Send(request, response)
    return
}

func NewGetResourceFileRequest() (request *GetResourceFileRequest) {
    request = &GetResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetResourceFile")
    
    
    return
}

func NewGetResourceFileResponse() (response *GetResourceFileResponse) {
    response = &GetResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetResourceFile
// 获取资源文件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceFile(request *GetResourceFileRequest) (response *GetResourceFileResponse, err error) {
    return c.GetResourceFileWithContext(context.Background(), request)
}

// GetResourceFile
// 获取资源文件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceFileWithContext(ctx context.Context, request *GetResourceFileRequest) (response *GetResourceFileResponse, err error) {
    if request == nil {
        request = NewGetResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetResourceFolderRequest() (request *GetResourceFolderRequest) {
    request = &GetResourceFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetResourceFolder")
    
    
    return
}

func NewGetResourceFolderResponse() (response *GetResourceFolderResponse) {
    response = &GetResourceFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetResourceFolder
// 查询资源文件文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceFolder(request *GetResourceFolderRequest) (response *GetResourceFolderResponse, err error) {
    return c.GetResourceFolderWithContext(context.Background(), request)
}

// GetResourceFolder
// 查询资源文件文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceFolderWithContext(ctx context.Context, request *GetResourceFolderRequest) (response *GetResourceFolderResponse, err error) {
    if request == nil {
        request = NewGetResourceFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetResourceFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetResourceFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetResourceFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGetResourceGroupMetricsRequest() (request *GetResourceGroupMetricsRequest) {
    request = &GetResourceGroupMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetResourceGroupMetrics")
    
    
    return
}

func NewGetResourceGroupMetricsResponse() (response *GetResourceGroupMetricsResponse) {
    response = &GetResourceGroupMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetResourceGroupMetrics
// 该接口用于查看指定执行资源组的监控指标
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetResourceGroupMetrics(request *GetResourceGroupMetricsRequest) (response *GetResourceGroupMetricsResponse, err error) {
    return c.GetResourceGroupMetricsWithContext(context.Background(), request)
}

// GetResourceGroupMetrics
// 该接口用于查看指定执行资源组的监控指标
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetResourceGroupMetricsWithContext(ctx context.Context, request *GetResourceGroupMetricsRequest) (response *GetResourceGroupMetricsResponse, err error) {
    if request == nil {
        request = NewGetResourceGroupMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetResourceGroupMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetResourceGroupMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetResourceGroupMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewGetSQLFolderRequest() (request *GetSQLFolderRequest) {
    request = &GetSQLFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetSQLFolder")
    
    
    return
}

func NewGetSQLFolderResponse() (response *GetSQLFolderResponse) {
    response = &GetSQLFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSQLFolder
// 获取sql文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetSQLFolder(request *GetSQLFolderRequest) (response *GetSQLFolderResponse, err error) {
    return c.GetSQLFolderWithContext(context.Background(), request)
}

// GetSQLFolder
// 获取sql文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetSQLFolderWithContext(ctx context.Context, request *GetSQLFolderRequest) (response *GetSQLFolderResponse, err error) {
    if request == nil {
        request = NewGetSQLFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetSQLFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSQLFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSQLFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGetSQLScriptRequest() (request *GetSQLScriptRequest) {
    request = &GetSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetSQLScript")
    
    
    return
}

func NewGetSQLScriptResponse() (response *GetSQLScriptResponse) {
    response = &GetSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSQLScript
// 查询脚本详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetSQLScript(request *GetSQLScriptRequest) (response *GetSQLScriptResponse, err error) {
    return c.GetSQLScriptWithContext(context.Background(), request)
}

// GetSQLScript
// 查询脚本详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetSQLScriptWithContext(ctx context.Context, request *GetSQLScriptRequest) (response *GetSQLScriptResponse, err error) {
    if request == nil {
        request = NewGetSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewGetTableRequest() (request *GetTableRequest) {
    request = &GetTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTable")
    
    
    return
}

func NewGetTableResponse() (response *GetTableResponse) {
    response = &GetTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTable
// 查询表详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetTable(request *GetTableRequest) (response *GetTableResponse, err error) {
    return c.GetTableWithContext(context.Background(), request)
}

// GetTable
// 查询表详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetTableWithContext(ctx context.Context, request *GetTableRequest) (response *GetTableResponse, err error) {
    if request == nil {
        request = NewGetTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTableResponse()
    err = c.Send(request, response)
    return
}

func NewGetTableColumnsRequest() (request *GetTableColumnsRequest) {
    request = &GetTableColumnsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTableColumns")
    
    
    return
}

func NewGetTableColumnsResponse() (response *GetTableColumnsResponse) {
    response = &GetTableColumnsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTableColumns
// 查询表所有字段列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetTableColumns(request *GetTableColumnsRequest) (response *GetTableColumnsResponse, err error) {
    return c.GetTableColumnsWithContext(context.Background(), request)
}

// GetTableColumns
// 查询表所有字段列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetTableColumnsWithContext(ctx context.Context, request *GetTableColumnsRequest) (response *GetTableColumnsResponse, err error) {
    if request == nil {
        request = NewGetTableColumnsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTableColumns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTableColumns require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTableColumnsResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskRequest() (request *GetTaskRequest) {
    request = &GetTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTask")
    
    
    return
}

func NewGetTaskResponse() (response *GetTaskResponse) {
    response = &GetTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTask
// 获取任务详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTask(request *GetTaskRequest) (response *GetTaskResponse, err error) {
    return c.GetTaskWithContext(context.Background(), request)
}

// GetTask
// 获取任务详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskWithContext(ctx context.Context, request *GetTaskRequest) (response *GetTaskResponse, err error) {
    if request == nil {
        request = NewGetTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskCodeRequest() (request *GetTaskCodeRequest) {
    request = &GetTaskCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskCode")
    
    
    return
}

func NewGetTaskCodeResponse() (response *GetTaskCodeResponse) {
    response = &GetTaskCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskCode
// 获取任务代码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskCode(request *GetTaskCodeRequest) (response *GetTaskCodeResponse, err error) {
    return c.GetTaskCodeWithContext(context.Background(), request)
}

// GetTaskCode
// 获取任务代码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskCodeWithContext(ctx context.Context, request *GetTaskCodeRequest) (response *GetTaskCodeResponse, err error) {
    if request == nil {
        request = NewGetTaskCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskFolderRequest() (request *GetTaskFolderRequest) {
    request = &GetTaskFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskFolder")
    
    
    return
}

func NewGetTaskFolderResponse() (response *GetTaskFolderResponse) {
    response = &GetTaskFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskFolder
// 查询任务文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskFolder(request *GetTaskFolderRequest) (response *GetTaskFolderResponse, err error) {
    return c.GetTaskFolderWithContext(context.Background(), request)
}

// GetTaskFolder
// 查询任务文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskFolderWithContext(ctx context.Context, request *GetTaskFolderRequest) (response *GetTaskFolderResponse, err error) {
    if request == nil {
        request = NewGetTaskFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskInstanceRequest() (request *GetTaskInstanceRequest) {
    request = &GetTaskInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskInstance")
    
    
    return
}

func NewGetTaskInstanceResponse() (response *GetTaskInstanceResponse) {
    response = &GetTaskInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskInstance
// 调度实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstance(request *GetTaskInstanceRequest) (response *GetTaskInstanceResponse, err error) {
    return c.GetTaskInstanceWithContext(context.Background(), request)
}

// GetTaskInstance
// 调度实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstanceWithContext(ctx context.Context, request *GetTaskInstanceRequest) (response *GetTaskInstanceResponse, err error) {
    if request == nil {
        request = NewGetTaskInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskInstanceLogRequest() (request *GetTaskInstanceLogRequest) {
    request = &GetTaskInstanceLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskInstanceLog")
    
    
    return
}

func NewGetTaskInstanceLogResponse() (response *GetTaskInstanceLogResponse) {
    response = &GetTaskInstanceLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskInstanceLog
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstanceLog(request *GetTaskInstanceLogRequest) (response *GetTaskInstanceLogResponse, err error) {
    return c.GetTaskInstanceLogWithContext(context.Background(), request)
}

// GetTaskInstanceLog
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTaskInstanceLogWithContext(ctx context.Context, request *GetTaskInstanceLogRequest) (response *GetTaskInstanceLogResponse, err error) {
    if request == nil {
        request = NewGetTaskInstanceLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskInstanceLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskInstanceLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskInstanceLogResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskVersionRequest() (request *GetTaskVersionRequest) {
    request = &GetTaskVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTaskVersion")
    
    
    return
}

func NewGetTaskVersionResponse() (response *GetTaskVersionResponse) {
    response = &GetTaskVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskVersion
// 拉取任务版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskVersion(request *GetTaskVersionRequest) (response *GetTaskVersionResponse, err error) {
    return c.GetTaskVersionWithContext(context.Background(), request)
}

// GetTaskVersion
// 拉取任务版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTaskVersionWithContext(ctx context.Context, request *GetTaskVersionRequest) (response *GetTaskVersionResponse, err error) {
    if request == nil {
        request = NewGetTaskVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTaskVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetTriggerTaskRequest() (request *GetTriggerTaskRequest) {
    request = &GetTriggerTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTriggerTask")
    
    
    return
}

func NewGetTriggerTaskResponse() (response *GetTriggerTaskResponse) {
    response = &GetTriggerTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTriggerTask
// 获取任务详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTask(request *GetTriggerTaskRequest) (response *GetTriggerTaskResponse, err error) {
    return c.GetTriggerTaskWithContext(context.Background(), request)
}

// GetTriggerTask
// 获取任务详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTaskWithContext(ctx context.Context, request *GetTriggerTaskRequest) (response *GetTriggerTaskResponse, err error) {
    if request == nil {
        request = NewGetTriggerTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTriggerTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTriggerTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTriggerTaskResponse()
    err = c.Send(request, response)
    return
}

func NewGetTriggerTaskCodeRequest() (request *GetTriggerTaskCodeRequest) {
    request = &GetTriggerTaskCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTriggerTaskCode")
    
    
    return
}

func NewGetTriggerTaskCodeResponse() (response *GetTriggerTaskCodeResponse) {
    response = &GetTriggerTaskCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTriggerTaskCode
// 获取工作流调度任务代码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTaskCode(request *GetTriggerTaskCodeRequest) (response *GetTriggerTaskCodeResponse, err error) {
    return c.GetTriggerTaskCodeWithContext(context.Background(), request)
}

// GetTriggerTaskCode
// 获取工作流调度任务代码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTaskCodeWithContext(ctx context.Context, request *GetTriggerTaskCodeRequest) (response *GetTriggerTaskCodeResponse, err error) {
    if request == nil {
        request = NewGetTriggerTaskCodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTriggerTaskCode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTriggerTaskCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTriggerTaskCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetTriggerTaskRunRequest() (request *GetTriggerTaskRunRequest) {
    request = &GetTriggerTaskRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTriggerTaskRun")
    
    
    return
}

func NewGetTriggerTaskRunResponse() (response *GetTriggerTaskRunResponse) {
    response = &GetTriggerTaskRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTriggerTaskRun
// 查询任务执行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTaskRun(request *GetTriggerTaskRunRequest) (response *GetTriggerTaskRunResponse, err error) {
    return c.GetTriggerTaskRunWithContext(context.Background(), request)
}

// GetTriggerTaskRun
// 查询任务执行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTaskRunWithContext(ctx context.Context, request *GetTriggerTaskRunRequest) (response *GetTriggerTaskRunResponse, err error) {
    if request == nil {
        request = NewGetTriggerTaskRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTriggerTaskRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTriggerTaskRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTriggerTaskRunResponse()
    err = c.Send(request, response)
    return
}

func NewGetTriggerTaskVersionRequest() (request *GetTriggerTaskVersionRequest) {
    request = &GetTriggerTaskVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTriggerTaskVersion")
    
    
    return
}

func NewGetTriggerTaskVersionResponse() (response *GetTriggerTaskVersionResponse) {
    response = &GetTriggerTaskVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTriggerTaskVersion
// 拉取任务版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTaskVersion(request *GetTriggerTaskVersionRequest) (response *GetTriggerTaskVersionResponse, err error) {
    return c.GetTriggerTaskVersionWithContext(context.Background(), request)
}

// GetTriggerTaskVersion
// 拉取任务版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTriggerTaskVersionWithContext(ctx context.Context, request *GetTriggerTaskVersionRequest) (response *GetTriggerTaskVersionResponse, err error) {
    if request == nil {
        request = NewGetTriggerTaskVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTriggerTaskVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTriggerTaskVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTriggerTaskVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetTriggerWorkflowRequest() (request *GetTriggerWorkflowRequest) {
    request = &GetTriggerWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTriggerWorkflow")
    
    
    return
}

func NewGetTriggerWorkflowResponse() (response *GetTriggerWorkflowResponse) {
    response = &GetTriggerWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTriggerWorkflow
// 获取工作流信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTriggerWorkflow(request *GetTriggerWorkflowRequest) (response *GetTriggerWorkflowResponse, err error) {
    return c.GetTriggerWorkflowWithContext(context.Background(), request)
}

// GetTriggerWorkflow
// 获取工作流信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTriggerWorkflowWithContext(ctx context.Context, request *GetTriggerWorkflowRequest) (response *GetTriggerWorkflowResponse, err error) {
    if request == nil {
        request = NewGetTriggerWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTriggerWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTriggerWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTriggerWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetTriggerWorkflowRunRequest() (request *GetTriggerWorkflowRunRequest) {
    request = &GetTriggerWorkflowRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetTriggerWorkflowRun")
    
    
    return
}

func NewGetTriggerWorkflowRunResponse() (response *GetTriggerWorkflowRunResponse) {
    response = &GetTriggerWorkflowRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTriggerWorkflowRun
// 查询工作流任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetTriggerWorkflowRun(request *GetTriggerWorkflowRunRequest) (response *GetTriggerWorkflowRunResponse, err error) {
    return c.GetTriggerWorkflowRunWithContext(context.Background(), request)
}

// GetTriggerWorkflowRun
// 查询工作流任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetTriggerWorkflowRunWithContext(ctx context.Context, request *GetTriggerWorkflowRunRequest) (response *GetTriggerWorkflowRunResponse, err error) {
    if request == nil {
        request = NewGetTriggerWorkflowRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetTriggerWorkflowRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTriggerWorkflowRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTriggerWorkflowRunResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowRequest() (request *GetWorkflowRequest) {
    request = &GetWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetWorkflow")
    
    
    return
}

func NewGetWorkflowResponse() (response *GetWorkflowResponse) {
    response = &GetWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflow
// 获取工作流信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetWorkflow(request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    return c.GetWorkflowWithContext(context.Background(), request)
}

// GetWorkflow
// 获取工作流信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetWorkflowWithContext(ctx context.Context, request *GetWorkflowRequest) (response *GetWorkflowResponse, err error) {
    if request == nil {
        request = NewGetWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewGetWorkflowFolderRequest() (request *GetWorkflowFolderRequest) {
    request = &GetWorkflowFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GetWorkflowFolder")
    
    
    return
}

func NewGetWorkflowFolderResponse() (response *GetWorkflowFolderResponse) {
    response = &GetWorkflowFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWorkflowFolder
// 查询文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetWorkflowFolder(request *GetWorkflowFolderRequest) (response *GetWorkflowFolderResponse, err error) {
    return c.GetWorkflowFolderWithContext(context.Background(), request)
}

// GetWorkflowFolder
// 查询文件夹详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetWorkflowFolderWithContext(ctx context.Context, request *GetWorkflowFolderRequest) (response *GetWorkflowFolderResponse, err error) {
    if request == nil {
        request = NewGetWorkflowFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GetWorkflowFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWorkflowFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWorkflowFolderResponse()
    err = c.Send(request, response)
    return
}

func NewGrantMemberProjectRoleRequest() (request *GrantMemberProjectRoleRequest) {
    request = &GrantMemberProjectRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "GrantMemberProjectRole")
    
    
    return
}

func NewGrantMemberProjectRoleResponse() (response *GrantMemberProjectRoleResponse) {
    response = &GrantMemberProjectRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GrantMemberProjectRole
// 修改项目用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GrantMemberProjectRole(request *GrantMemberProjectRoleRequest) (response *GrantMemberProjectRoleResponse, err error) {
    return c.GrantMemberProjectRoleWithContext(context.Background(), request)
}

// GrantMemberProjectRole
// 修改项目用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GrantMemberProjectRoleWithContext(ctx context.Context, request *GrantMemberProjectRoleRequest) (response *GrantMemberProjectRoleResponse, err error) {
    if request == nil {
        request = NewGrantMemberProjectRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "GrantMemberProjectRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GrantMemberProjectRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewGrantMemberProjectRoleResponse()
    err = c.Send(request, response)
    return
}

func NewKillTaskInstancesAsyncRequest() (request *KillTaskInstancesAsyncRequest) {
    request = &KillTaskInstancesAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "KillTaskInstancesAsync")
    
    
    return
}

func NewKillTaskInstancesAsyncResponse() (response *KillTaskInstancesAsyncResponse) {
    response = &KillTaskInstancesAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillTaskInstancesAsync
// 实例批量终止操作-异步操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) KillTaskInstancesAsync(request *KillTaskInstancesAsyncRequest) (response *KillTaskInstancesAsyncResponse, err error) {
    return c.KillTaskInstancesAsyncWithContext(context.Background(), request)
}

// KillTaskInstancesAsync
// 实例批量终止操作-异步操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) KillTaskInstancesAsyncWithContext(ctx context.Context, request *KillTaskInstancesAsyncRequest) (response *KillTaskInstancesAsyncResponse, err error) {
    if request == nil {
        request = NewKillTaskInstancesAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "KillTaskInstancesAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillTaskInstancesAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillTaskInstancesAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewKillTriggerWorkflowRunsRequest() (request *KillTriggerWorkflowRunsRequest) {
    request = &KillTriggerWorkflowRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "KillTriggerWorkflowRuns")
    
    
    return
}

func NewKillTriggerWorkflowRunsResponse() (response *KillTriggerWorkflowRunsResponse) {
    response = &KillTriggerWorkflowRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillTriggerWorkflowRuns
// 终止运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) KillTriggerWorkflowRuns(request *KillTriggerWorkflowRunsRequest) (response *KillTriggerWorkflowRunsResponse, err error) {
    return c.KillTriggerWorkflowRunsWithContext(context.Background(), request)
}

// KillTriggerWorkflowRuns
// 终止运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) KillTriggerWorkflowRunsWithContext(ctx context.Context, request *KillTriggerWorkflowRunsRequest) (response *KillTriggerWorkflowRunsResponse, err error) {
    if request == nil {
        request = NewKillTriggerWorkflowRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "KillTriggerWorkflowRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillTriggerWorkflowRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillTriggerWorkflowRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListAlarmMessagesRequest() (request *ListAlarmMessagesRequest) {
    request = &ListAlarmMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListAlarmMessages")
    
    
    return
}

func NewListAlarmMessagesResponse() (response *ListAlarmMessagesResponse) {
    response = &ListAlarmMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAlarmMessages
// 获取告警信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAlarmMessages(request *ListAlarmMessagesRequest) (response *ListAlarmMessagesResponse, err error) {
    return c.ListAlarmMessagesWithContext(context.Background(), request)
}

// ListAlarmMessages
// 获取告警信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAlarmMessagesWithContext(ctx context.Context, request *ListAlarmMessagesRequest) (response *ListAlarmMessagesResponse, err error) {
    if request == nil {
        request = NewListAlarmMessagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListAlarmMessages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAlarmMessages require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAlarmMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewListCatalogRequest() (request *ListCatalogRequest) {
    request = &ListCatalogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListCatalog")
    
    
    return
}

func NewListCatalogResponse() (response *ListCatalogResponse) {
    response = &ListCatalogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListCatalog
// 获取资产目录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCatalog(request *ListCatalogRequest) (response *ListCatalogResponse, err error) {
    return c.ListCatalogWithContext(context.Background(), request)
}

// ListCatalog
// 获取资产目录信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCatalogWithContext(ctx context.Context, request *ListCatalogRequest) (response *ListCatalogResponse, err error) {
    if request == nil {
        request = NewListCatalogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListCatalog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCatalog require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCatalogResponse()
    err = c.Send(request, response)
    return
}

func NewListCodeFolderContentsRequest() (request *ListCodeFolderContentsRequest) {
    request = &ListCodeFolderContentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListCodeFolderContents")
    
    
    return
}

func NewListCodeFolderContentsResponse() (response *ListCodeFolderContentsResponse) {
    response = &ListCodeFolderContentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListCodeFolderContents
// 获取文件夹内容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCodeFolderContents(request *ListCodeFolderContentsRequest) (response *ListCodeFolderContentsResponse, err error) {
    return c.ListCodeFolderContentsWithContext(context.Background(), request)
}

// ListCodeFolderContents
// 获取文件夹内容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCodeFolderContentsWithContext(ctx context.Context, request *ListCodeFolderContentsRequest) (response *ListCodeFolderContentsResponse, err error) {
    if request == nil {
        request = NewListCodeFolderContentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListCodeFolderContents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCodeFolderContents require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCodeFolderContentsResponse()
    err = c.Send(request, response)
    return
}

func NewListCodePermissionsRequest() (request *ListCodePermissionsRequest) {
    request = &ListCodePermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListCodePermissions")
    
    
    return
}

func NewListCodePermissionsResponse() (response *ListCodePermissionsResponse) {
    response = &ListCodePermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListCodePermissions
// 查看CodeStudio实体权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCodePermissions(request *ListCodePermissionsRequest) (response *ListCodePermissionsResponse, err error) {
    return c.ListCodePermissionsWithContext(context.Background(), request)
}

// ListCodePermissions
// 查看CodeStudio实体权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListCodePermissionsWithContext(ctx context.Context, request *ListCodePermissionsRequest) (response *ListCodePermissionsResponse, err error) {
    if request == nil {
        request = NewListCodePermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListCodePermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCodePermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCodePermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewListColumnLineageRequest() (request *ListColumnLineageRequest) {
    request = &ListColumnLineageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListColumnLineage")
    
    
    return
}

func NewListColumnLineageResponse() (response *ListColumnLineageResponse) {
    response = &ListColumnLineageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListColumnLineage
// 获取表字段血缘信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListColumnLineage(request *ListColumnLineageRequest) (response *ListColumnLineageResponse, err error) {
    return c.ListColumnLineageWithContext(context.Background(), request)
}

// ListColumnLineage
// 获取表字段血缘信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListColumnLineageWithContext(ctx context.Context, request *ListColumnLineageRequest) (response *ListColumnLineageResponse, err error) {
    if request == nil {
        request = NewListColumnLineageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListColumnLineage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListColumnLineage require credential")
    }

    request.SetContext(ctx)
    
    response = NewListColumnLineageResponse()
    err = c.Send(request, response)
    return
}

func NewListDataBackfillInstancesRequest() (request *ListDataBackfillInstancesRequest) {
    request = &ListDataBackfillInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDataBackfillInstances")
    
    
    return
}

func NewListDataBackfillInstancesResponse() (response *ListDataBackfillInstancesResponse) {
    response = &ListDataBackfillInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDataBackfillInstances
// 获取单次补录的所有实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDataBackfillInstances(request *ListDataBackfillInstancesRequest) (response *ListDataBackfillInstancesResponse, err error) {
    return c.ListDataBackfillInstancesWithContext(context.Background(), request)
}

// ListDataBackfillInstances
// 获取单次补录的所有实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDataBackfillInstancesWithContext(ctx context.Context, request *ListDataBackfillInstancesRequest) (response *ListDataBackfillInstancesResponse, err error) {
    if request == nil {
        request = NewListDataBackfillInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDataBackfillInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDataBackfillInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDataBackfillInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListDataSourcesRequest() (request *ListDataSourcesRequest) {
    request = &ListDataSourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDataSources")
    
    
    return
}

func NewListDataSourcesResponse() (response *ListDataSourcesResponse) {
    response = &ListDataSourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDataSources
// 该接口用于查询指定项目中的数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListDataSources(request *ListDataSourcesRequest) (response *ListDataSourcesResponse, err error) {
    return c.ListDataSourcesWithContext(context.Background(), request)
}

// ListDataSources
// 该接口用于查询指定项目中的数据源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListDataSourcesWithContext(ctx context.Context, request *ListDataSourcesRequest) (response *ListDataSourcesResponse, err error) {
    if request == nil {
        request = NewListDataSourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDataSources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDataSources require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDataSourcesResponse()
    err = c.Send(request, response)
    return
}

func NewListDatabaseRequest() (request *ListDatabaseRequest) {
    request = &ListDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDatabase")
    
    
    return
}

func NewListDatabaseResponse() (response *ListDatabaseResponse) {
    response = &ListDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDatabase
// 获取资产数据库信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListDatabase(request *ListDatabaseRequest) (response *ListDatabaseResponse, err error) {
    return c.ListDatabaseWithContext(context.Background(), request)
}

// ListDatabase
// 获取资产数据库信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListDatabaseWithContext(ctx context.Context, request *ListDatabaseRequest) (response *ListDatabaseResponse, err error) {
    if request == nil {
        request = NewListDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewListDownstreamOpsTasksRequest() (request *ListDownstreamOpsTasksRequest) {
    request = &ListDownstreamOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDownstreamOpsTasks")
    
    
    return
}

func NewListDownstreamOpsTasksResponse() (response *ListDownstreamOpsTasksResponse) {
    response = &ListDownstreamOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDownstreamOpsTasks
// 获取任务直接下游详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamOpsTasks(request *ListDownstreamOpsTasksRequest) (response *ListDownstreamOpsTasksResponse, err error) {
    return c.ListDownstreamOpsTasksWithContext(context.Background(), request)
}

// ListDownstreamOpsTasks
// 获取任务直接下游详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamOpsTasksWithContext(ctx context.Context, request *ListDownstreamOpsTasksRequest) (response *ListDownstreamOpsTasksResponse, err error) {
    if request == nil {
        request = NewListDownstreamOpsTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDownstreamOpsTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDownstreamOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDownstreamOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListDownstreamTaskInstancesRequest() (request *ListDownstreamTaskInstancesRequest) {
    request = &ListDownstreamTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDownstreamTaskInstances")
    
    
    return
}

func NewListDownstreamTaskInstancesResponse() (response *ListDownstreamTaskInstancesResponse) {
    response = &ListDownstreamTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDownstreamTaskInstances
// 获取实例直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTaskInstances(request *ListDownstreamTaskInstancesRequest) (response *ListDownstreamTaskInstancesResponse, err error) {
    return c.ListDownstreamTaskInstancesWithContext(context.Background(), request)
}

// ListDownstreamTaskInstances
// 获取实例直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTaskInstancesWithContext(ctx context.Context, request *ListDownstreamTaskInstancesRequest) (response *ListDownstreamTaskInstancesResponse, err error) {
    if request == nil {
        request = NewListDownstreamTaskInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDownstreamTaskInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDownstreamTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDownstreamTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListDownstreamTasksRequest() (request *ListDownstreamTasksRequest) {
    request = &ListDownstreamTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDownstreamTasks")
    
    
    return
}

func NewListDownstreamTasksResponse() (response *ListDownstreamTasksResponse) {
    response = &ListDownstreamTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDownstreamTasks
// 获取任务直接下游详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTasks(request *ListDownstreamTasksRequest) (response *ListDownstreamTasksResponse, err error) {
    return c.ListDownstreamTasksWithContext(context.Background(), request)
}

// ListDownstreamTasks
// 获取任务直接下游详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListDownstreamTasksWithContext(ctx context.Context, request *ListDownstreamTasksRequest) (response *ListDownstreamTasksResponse, err error) {
    if request == nil {
        request = NewListDownstreamTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDownstreamTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDownstreamTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDownstreamTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListDownstreamTriggerTasksRequest() (request *ListDownstreamTriggerTasksRequest) {
    request = &ListDownstreamTriggerTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListDownstreamTriggerTasks")
    
    
    return
}

func NewListDownstreamTriggerTasksResponse() (response *ListDownstreamTriggerTasksResponse) {
    response = &ListDownstreamTriggerTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDownstreamTriggerTasks
// 获取任务直接下游详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDownstreamTriggerTasks(request *ListDownstreamTriggerTasksRequest) (response *ListDownstreamTriggerTasksResponse, err error) {
    return c.ListDownstreamTriggerTasksWithContext(context.Background(), request)
}

// ListDownstreamTriggerTasks
// 获取任务直接下游详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDownstreamTriggerTasksWithContext(ctx context.Context, request *ListDownstreamTriggerTasksRequest) (response *ListDownstreamTriggerTasksResponse, err error) {
    if request == nil {
        request = NewListDownstreamTriggerTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListDownstreamTriggerTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDownstreamTriggerTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDownstreamTriggerTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListLineageRequest() (request *ListLineageRequest) {
    request = &ListLineageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListLineage")
    
    
    return
}

func NewListLineageResponse() (response *ListLineageResponse) {
    response = &ListLineageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListLineage
// 获取资产血缘信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListLineage(request *ListLineageRequest) (response *ListLineageResponse, err error) {
    return c.ListLineageWithContext(context.Background(), request)
}

// ListLineage
// 获取资产血缘信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListLineageWithContext(ctx context.Context, request *ListLineageRequest) (response *ListLineageResponse, err error) {
    if request == nil {
        request = NewListLineageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListLineage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLineage require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLineageResponse()
    err = c.Send(request, response)
    return
}

func NewListOpsAlarmRulesRequest() (request *ListOpsAlarmRulesRequest) {
    request = &ListOpsAlarmRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListOpsAlarmRules")
    
    
    return
}

func NewListOpsAlarmRulesResponse() (response *ListOpsAlarmRulesResponse) {
    response = &ListOpsAlarmRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOpsAlarmRules
// 查询告警规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsAlarmRules(request *ListOpsAlarmRulesRequest) (response *ListOpsAlarmRulesResponse, err error) {
    return c.ListOpsAlarmRulesWithContext(context.Background(), request)
}

// ListOpsAlarmRules
// 查询告警规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsAlarmRulesWithContext(ctx context.Context, request *ListOpsAlarmRulesRequest) (response *ListOpsAlarmRulesResponse, err error) {
    if request == nil {
        request = NewListOpsAlarmRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListOpsAlarmRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOpsAlarmRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOpsAlarmRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListOpsTasksRequest() (request *ListOpsTasksRequest) {
    request = &ListOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListOpsTasks")
    
    
    return
}

func NewListOpsTasksResponse() (response *ListOpsTasksResponse) {
    response = &ListOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOpsTasks
// 根据项目id获取任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsTasks(request *ListOpsTasksRequest) (response *ListOpsTasksResponse, err error) {
    return c.ListOpsTasksWithContext(context.Background(), request)
}

// ListOpsTasks
// 根据项目id获取任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListOpsTasksWithContext(ctx context.Context, request *ListOpsTasksRequest) (response *ListOpsTasksResponse, err error) {
    if request == nil {
        request = NewListOpsTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListOpsTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListOpsTriggerWorkflowsRequest() (request *ListOpsTriggerWorkflowsRequest) {
    request = &ListOpsTriggerWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListOpsTriggerWorkflows")
    
    
    return
}

func NewListOpsTriggerWorkflowsResponse() (response *ListOpsTriggerWorkflowsResponse) {
    response = &ListOpsTriggerWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOpsTriggerWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListOpsTriggerWorkflows(request *ListOpsTriggerWorkflowsRequest) (response *ListOpsTriggerWorkflowsResponse, err error) {
    return c.ListOpsTriggerWorkflowsWithContext(context.Background(), request)
}

// ListOpsTriggerWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListOpsTriggerWorkflowsWithContext(ctx context.Context, request *ListOpsTriggerWorkflowsRequest) (response *ListOpsTriggerWorkflowsResponse, err error) {
    if request == nil {
        request = NewListOpsTriggerWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListOpsTriggerWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOpsTriggerWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOpsTriggerWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewListOpsWorkflowsRequest() (request *ListOpsWorkflowsRequest) {
    request = &ListOpsWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListOpsWorkflows")
    
    
    return
}

func NewListOpsWorkflowsResponse() (response *ListOpsWorkflowsResponse) {
    response = &ListOpsWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOpsWorkflows
// 根据项目ID获取项目下工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListOpsWorkflows(request *ListOpsWorkflowsRequest) (response *ListOpsWorkflowsResponse, err error) {
    return c.ListOpsWorkflowsWithContext(context.Background(), request)
}

// ListOpsWorkflows
// 根据项目ID获取项目下工作流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListOpsWorkflowsWithContext(ctx context.Context, request *ListOpsWorkflowsRequest) (response *ListOpsWorkflowsResponse, err error) {
    if request == nil {
        request = NewListOpsWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListOpsWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOpsWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOpsWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewListProcessLineageRequest() (request *ListProcessLineageRequest) {
    request = &ListProcessLineageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListProcessLineage")
    
    
    return
}

func NewListProcessLineageResponse() (response *ListProcessLineageResponse) {
    response = &ListProcessLineageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListProcessLineage
// 获取资产血缘信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListProcessLineage(request *ListProcessLineageRequest) (response *ListProcessLineageResponse, err error) {
    return c.ListProcessLineageWithContext(context.Background(), request)
}

// ListProcessLineage
// 获取资产血缘信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListProcessLineageWithContext(ctx context.Context, request *ListProcessLineageRequest) (response *ListProcessLineageResponse, err error) {
    if request == nil {
        request = NewListProcessLineageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListProcessLineage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListProcessLineage require credential")
    }

    request.SetContext(ctx)
    
    response = NewListProcessLineageResponse()
    err = c.Send(request, response)
    return
}

func NewListProjectMembersRequest() (request *ListProjectMembersRequest) {
    request = &ListProjectMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListProjectMembers")
    
    
    return
}

func NewListProjectMembersResponse() (response *ListProjectMembersResponse) {
    response = &ListProjectMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListProjectMembers
// 获取项目下的用户，分页返回
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListProjectMembers(request *ListProjectMembersRequest) (response *ListProjectMembersResponse, err error) {
    return c.ListProjectMembersWithContext(context.Background(), request)
}

// ListProjectMembers
// 获取项目下的用户，分页返回
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListProjectMembersWithContext(ctx context.Context, request *ListProjectMembersRequest) (response *ListProjectMembersResponse, err error) {
    if request == nil {
        request = NewListProjectMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListProjectMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListProjectMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListProjectMembersResponse()
    err = c.Send(request, response)
    return
}

func NewListProjectRolesRequest() (request *ListProjectRolesRequest) {
    request = &ListProjectRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListProjectRoles")
    
    
    return
}

func NewListProjectRolesResponse() (response *ListProjectRolesResponse) {
    response = &ListProjectRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListProjectRoles
// 获取角色列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListProjectRoles(request *ListProjectRolesRequest) (response *ListProjectRolesResponse, err error) {
    return c.ListProjectRolesWithContext(context.Background(), request)
}

// ListProjectRoles
// 获取角色列表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListProjectRolesWithContext(ctx context.Context, request *ListProjectRolesRequest) (response *ListProjectRolesResponse, err error) {
    if request == nil {
        request = NewListProjectRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListProjectRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListProjectRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListProjectRolesResponse()
    err = c.Send(request, response)
    return
}

func NewListProjectsRequest() (request *ListProjectsRequest) {
    request = &ListProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListProjects")
    
    
    return
}

func NewListProjectsResponse() (response *ListProjectsResponse) {
    response = &ListProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListProjects
// 租户全局范围的项目列表，与用户查看范围无关.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListProjects(request *ListProjectsRequest) (response *ListProjectsResponse, err error) {
    return c.ListProjectsWithContext(context.Background(), request)
}

// ListProjects
// 租户全局范围的项目列表，与用户查看范围无关.
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"
func (c *Client) ListProjectsWithContext(ctx context.Context, request *ListProjectsRequest) (response *ListProjectsResponse, err error) {
    if request == nil {
        request = NewListProjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListProjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewListProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewListQualityRuleGroupExecResultsByPageRequest() (request *ListQualityRuleGroupExecResultsByPageRequest) {
    request = &ListQualityRuleGroupExecResultsByPageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListQualityRuleGroupExecResultsByPage")
    
    
    return
}

func NewListQualityRuleGroupExecResultsByPageResponse() (response *ListQualityRuleGroupExecResultsByPageResponse) {
    response = &ListQualityRuleGroupExecResultsByPageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQualityRuleGroupExecResultsByPage
// 规则组执行结果分页查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) ListQualityRuleGroupExecResultsByPage(request *ListQualityRuleGroupExecResultsByPageRequest) (response *ListQualityRuleGroupExecResultsByPageResponse, err error) {
    return c.ListQualityRuleGroupExecResultsByPageWithContext(context.Background(), request)
}

// ListQualityRuleGroupExecResultsByPage
// 规则组执行结果分页查询接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
func (c *Client) ListQualityRuleGroupExecResultsByPageWithContext(ctx context.Context, request *ListQualityRuleGroupExecResultsByPageRequest) (response *ListQualityRuleGroupExecResultsByPageResponse, err error) {
    if request == nil {
        request = NewListQualityRuleGroupExecResultsByPageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListQualityRuleGroupExecResultsByPage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQualityRuleGroupExecResultsByPage require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQualityRuleGroupExecResultsByPageResponse()
    err = c.Send(request, response)
    return
}

func NewListQualityRuleGroupsRequest() (request *ListQualityRuleGroupsRequest) {
    request = &ListQualityRuleGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListQualityRuleGroups")
    
    
    return
}

func NewListQualityRuleGroupsResponse() (response *ListQualityRuleGroupsResponse) {
    response = &ListQualityRuleGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQualityRuleGroups
// 【过滤条件】
//
// {表idTableId,支持匹配}       {任务负责人Operator,支持匹配}    
//
// 【必要字段】
//
// {数据来源TableId}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRuleGroups(request *ListQualityRuleGroupsRequest) (response *ListQualityRuleGroupsResponse, err error) {
    return c.ListQualityRuleGroupsWithContext(context.Background(), request)
}

// ListQualityRuleGroups
// 【过滤条件】
//
// {表idTableId,支持匹配}       {任务负责人Operator,支持匹配}    
//
// 【必要字段】
//
// {数据来源TableId}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRuleGroupsWithContext(ctx context.Context, request *ListQualityRuleGroupsRequest) (response *ListQualityRuleGroupsResponse, err error) {
    if request == nil {
        request = NewListQualityRuleGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListQualityRuleGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQualityRuleGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQualityRuleGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListQualityRuleGroupsTableRequest() (request *ListQualityRuleGroupsTableRequest) {
    request = &ListQualityRuleGroupsTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListQualityRuleGroupsTable")
    
    
    return
}

func NewListQualityRuleGroupsTableResponse() (response *ListQualityRuleGroupsTableResponse) {
    response = &ListQualityRuleGroupsTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQualityRuleGroupsTable
// 查询监控列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRuleGroupsTable(request *ListQualityRuleGroupsTableRequest) (response *ListQualityRuleGroupsTableResponse, err error) {
    return c.ListQualityRuleGroupsTableWithContext(context.Background(), request)
}

// ListQualityRuleGroupsTable
// 查询监控列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRuleGroupsTableWithContext(ctx context.Context, request *ListQualityRuleGroupsTableRequest) (response *ListQualityRuleGroupsTableResponse, err error) {
    if request == nil {
        request = NewListQualityRuleGroupsTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListQualityRuleGroupsTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQualityRuleGroupsTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQualityRuleGroupsTableResponse()
    err = c.Send(request, response)
    return
}

func NewListQualityRuleTemplatesRequest() (request *ListQualityRuleTemplatesRequest) {
    request = &ListQualityRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListQualityRuleTemplates")
    
    
    return
}

func NewListQualityRuleTemplatesResponse() (response *ListQualityRuleTemplatesResponse) {
    response = &ListQualityRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQualityRuleTemplates
// 【过滤条件】 {模板名称Name，查询使用Keyword进行模糊匹配} {模板类型type，1.系统模板 2.自定义模板} {质量检测维度QualityDims, 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性} 【排序字段】 { 引用数排序类型CitationOrderType，根据引用数量排序 ASC DESC}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRuleTemplates(request *ListQualityRuleTemplatesRequest) (response *ListQualityRuleTemplatesResponse, err error) {
    return c.ListQualityRuleTemplatesWithContext(context.Background(), request)
}

// ListQualityRuleTemplates
// 【过滤条件】 {模板名称Name，查询使用Keyword进行模糊匹配} {模板类型type，1.系统模板 2.自定义模板} {质量检测维度QualityDims, 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性} 【排序字段】 { 引用数排序类型CitationOrderType，根据引用数量排序 ASC DESC}
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRuleTemplatesWithContext(ctx context.Context, request *ListQualityRuleTemplatesRequest) (response *ListQualityRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewListQualityRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListQualityRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQualityRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQualityRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewListQualityRulesRequest() (request *ListQualityRulesRequest) {
    request = &ListQualityRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListQualityRules")
    
    
    return
}

func NewListQualityRulesResponse() (response *ListQualityRulesResponse) {
    response = &ListQualityRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQualityRules
// 分页查询质量规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRules(request *ListQualityRulesRequest) (response *ListQualityRulesResponse, err error) {
    return c.ListQualityRulesWithContext(context.Background(), request)
}

// ListQualityRules
// 分页查询质量规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"
//  UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"
func (c *Client) ListQualityRulesWithContext(ctx context.Context, request *ListQualityRulesRequest) (response *ListQualityRulesResponse, err error) {
    if request == nil {
        request = NewListQualityRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListQualityRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQualityRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQualityRulesResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceFilesRequest() (request *ListResourceFilesRequest) {
    request = &ListResourceFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListResourceFiles")
    
    
    return
}

func NewListResourceFilesResponse() (response *ListResourceFilesResponse) {
    response = &ListResourceFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListResourceFiles
// 获取资源文件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFiles(request *ListResourceFilesRequest) (response *ListResourceFilesResponse, err error) {
    return c.ListResourceFilesWithContext(context.Background(), request)
}

// ListResourceFiles
// 获取资源文件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFilesWithContext(ctx context.Context, request *ListResourceFilesRequest) (response *ListResourceFilesResponse, err error) {
    if request == nil {
        request = NewListResourceFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListResourceFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListResourceFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListResourceFilesResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceFoldersRequest() (request *ListResourceFoldersRequest) {
    request = &ListResourceFoldersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListResourceFolders")
    
    
    return
}

func NewListResourceFoldersResponse() (response *ListResourceFoldersResponse) {
    response = &ListResourceFoldersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListResourceFolders
// 查询资源文件文件夹列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFolders(request *ListResourceFoldersRequest) (response *ListResourceFoldersResponse, err error) {
    return c.ListResourceFoldersWithContext(context.Background(), request)
}

// ListResourceFolders
// 查询资源文件文件夹列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListResourceFoldersWithContext(ctx context.Context, request *ListResourceFoldersRequest) (response *ListResourceFoldersResponse, err error) {
    if request == nil {
        request = NewListResourceFoldersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListResourceFolders")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListResourceFolders require credential")
    }

    request.SetContext(ctx)
    
    response = NewListResourceFoldersResponse()
    err = c.Send(request, response)
    return
}

func NewListResourceGroupsRequest() (request *ListResourceGroupsRequest) {
    request = &ListResourceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListResourceGroups")
    
    
    return
}

func NewListResourceGroupsResponse() (response *ListResourceGroupsResponse) {
    response = &ListResourceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListResourceGroups
// 该接口用于查询执行资源组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListResourceGroups(request *ListResourceGroupsRequest) (response *ListResourceGroupsResponse, err error) {
    return c.ListResourceGroupsWithContext(context.Background(), request)
}

// ListResourceGroups
// 该接口用于查询执行资源组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListResourceGroupsWithContext(ctx context.Context, request *ListResourceGroupsRequest) (response *ListResourceGroupsResponse, err error) {
    if request == nil {
        request = NewListResourceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListResourceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListResourceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListResourceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListSQLFolderContentsRequest() (request *ListSQLFolderContentsRequest) {
    request = &ListSQLFolderContentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListSQLFolderContents")
    
    
    return
}

func NewListSQLFolderContentsResponse() (response *ListSQLFolderContentsResponse) {
    response = &ListSQLFolderContentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSQLFolderContents
// 查询数据探索文件夹树，包括文件夹下的脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLFolderContents(request *ListSQLFolderContentsRequest) (response *ListSQLFolderContentsResponse, err error) {
    return c.ListSQLFolderContentsWithContext(context.Background(), request)
}

// ListSQLFolderContents
// 查询数据探索文件夹树，包括文件夹下的脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLFolderContentsWithContext(ctx context.Context, request *ListSQLFolderContentsRequest) (response *ListSQLFolderContentsResponse, err error) {
    if request == nil {
        request = NewListSQLFolderContentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListSQLFolderContents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSQLFolderContents require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSQLFolderContentsResponse()
    err = c.Send(request, response)
    return
}

func NewListSQLScriptRunsRequest() (request *ListSQLScriptRunsRequest) {
    request = &ListSQLScriptRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListSQLScriptRuns")
    
    
    return
}

func NewListSQLScriptRunsResponse() (response *ListSQLScriptRunsResponse) {
    response = &ListSQLScriptRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSQLScriptRuns
// 查询SQL运行记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLScriptRuns(request *ListSQLScriptRunsRequest) (response *ListSQLScriptRunsResponse, err error) {
    return c.ListSQLScriptRunsWithContext(context.Background(), request)
}

// ListSQLScriptRuns
// 查询SQL运行记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSQLScriptRunsWithContext(ctx context.Context, request *ListSQLScriptRunsRequest) (response *ListSQLScriptRunsResponse, err error) {
    if request == nil {
        request = NewListSQLScriptRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListSQLScriptRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSQLScriptRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSQLScriptRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListSchemaRequest() (request *ListSchemaRequest) {
    request = &ListSchemaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListSchema")
    
    
    return
}

func NewListSchemaResponse() (response *ListSchemaResponse) {
    response = &ListSchemaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSchema
// 获取资产数据库Schema信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSchema(request *ListSchemaRequest) (response *ListSchemaResponse, err error) {
    return c.ListSchemaWithContext(context.Background(), request)
}

// ListSchema
// 获取资产数据库Schema信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSchemaWithContext(ctx context.Context, request *ListSchemaRequest) (response *ListSchemaResponse, err error) {
    if request == nil {
        request = NewListSchemaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListSchema")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSchema require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSchemaResponse()
    err = c.Send(request, response)
    return
}

func NewListTableRequest() (request *ListTableRequest) {
    request = &ListTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTable")
    
    
    return
}

func NewListTableResponse() (response *ListTableResponse) {
    response = &ListTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTable
// 获取资产表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListTable(request *ListTableRequest) (response *ListTableResponse, err error) {
    return c.ListTableWithContext(context.Background(), request)
}

// ListTable
// 获取资产表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListTableWithContext(ctx context.Context, request *ListTableRequest) (response *ListTableResponse, err error) {
    if request == nil {
        request = NewListTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTableResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskFoldersRequest() (request *ListTaskFoldersRequest) {
    request = &ListTaskFoldersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTaskFolders")
    
    
    return
}

func NewListTaskFoldersResponse() (response *ListTaskFoldersResponse) {
    response = &ListTaskFoldersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTaskFolders
// 查询任务文件夹列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListTaskFolders(request *ListTaskFoldersRequest) (response *ListTaskFoldersResponse, err error) {
    return c.ListTaskFoldersWithContext(context.Background(), request)
}

// ListTaskFolders
// 查询任务文件夹列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListTaskFoldersWithContext(ctx context.Context, request *ListTaskFoldersRequest) (response *ListTaskFoldersResponse, err error) {
    if request == nil {
        request = NewListTaskFoldersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTaskFolders")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskFolders require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskFoldersResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskInstanceExecutionsRequest() (request *ListTaskInstanceExecutionsRequest) {
    request = &ListTaskInstanceExecutionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTaskInstanceExecutions")
    
    
    return
}

func NewListTaskInstanceExecutionsResponse() (response *ListTaskInstanceExecutionsResponse) {
    response = &ListTaskInstanceExecutionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTaskInstanceExecutions
// 调度实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstanceExecutions(request *ListTaskInstanceExecutionsRequest) (response *ListTaskInstanceExecutionsResponse, err error) {
    return c.ListTaskInstanceExecutionsWithContext(context.Background(), request)
}

// ListTaskInstanceExecutions
// 调度实例详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstanceExecutionsWithContext(ctx context.Context, request *ListTaskInstanceExecutionsRequest) (response *ListTaskInstanceExecutionsResponse, err error) {
    if request == nil {
        request = NewListTaskInstanceExecutionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTaskInstanceExecutions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskInstanceExecutions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskInstanceExecutionsResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskInstancesRequest() (request *ListTaskInstancesRequest) {
    request = &ListTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTaskInstances")
    
    
    return
}

func NewListTaskInstancesResponse() (response *ListTaskInstancesResponse) {
    response = &ListTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTaskInstances
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstances(request *ListTaskInstancesRequest) (response *ListTaskInstancesResponse, err error) {
    return c.ListTaskInstancesWithContext(context.Background(), request)
}

// ListTaskInstances
// 获取实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskInstancesWithContext(ctx context.Context, request *ListTaskInstancesRequest) (response *ListTaskInstancesResponse, err error) {
    if request == nil {
        request = NewListTaskInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTaskInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListTaskVersionsRequest() (request *ListTaskVersionsRequest) {
    request = &ListTaskVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTaskVersions")
    
    
    return
}

func NewListTaskVersionsResponse() (response *ListTaskVersionsResponse) {
    response = &ListTaskVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTaskVersions
// 任务保存版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskVersions(request *ListTaskVersionsRequest) (response *ListTaskVersionsResponse, err error) {
    return c.ListTaskVersionsWithContext(context.Background(), request)
}

// ListTaskVersions
// 任务保存版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTaskVersionsWithContext(ctx context.Context, request *ListTaskVersionsRequest) (response *ListTaskVersionsResponse, err error) {
    if request == nil {
        request = NewListTaskVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTaskVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTaskVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTaskVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListTasksRequest() (request *ListTasksRequest) {
    request = &ListTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTasks")
    
    
    return
}

func NewListTasksResponse() (response *ListTasksResponse) {
    response = &ListTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTasks
// 查询任务分页信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
    return c.ListTasksWithContext(context.Background(), request)
}

// ListTasks
// 查询任务分页信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) (response *ListTasksResponse, err error) {
    if request == nil {
        request = NewListTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListTenantRolesRequest() (request *ListTenantRolesRequest) {
    request = &ListTenantRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTenantRoles")
    
    
    return
}

func NewListTenantRolesResponse() (response *ListTenantRolesResponse) {
    response = &ListTenantRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTenantRoles
// 获取所有主账号角色列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTenantRoles(request *ListTenantRolesRequest) (response *ListTenantRolesResponse, err error) {
    return c.ListTenantRolesWithContext(context.Background(), request)
}

// ListTenantRoles
// 获取所有主账号角色列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTenantRolesWithContext(ctx context.Context, request *ListTenantRolesRequest) (response *ListTenantRolesResponse, err error) {
    if request == nil {
        request = NewListTenantRolesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTenantRoles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTenantRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTenantRolesResponse()
    err = c.Send(request, response)
    return
}

func NewListTriggerTaskVersionsRequest() (request *ListTriggerTaskVersionsRequest) {
    request = &ListTriggerTaskVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTriggerTaskVersions")
    
    
    return
}

func NewListTriggerTaskVersionsResponse() (response *ListTriggerTaskVersionsResponse) {
    response = &ListTriggerTaskVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTriggerTaskVersions
// 任务保存版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTriggerTaskVersions(request *ListTriggerTaskVersionsRequest) (response *ListTriggerTaskVersionsResponse, err error) {
    return c.ListTriggerTaskVersionsWithContext(context.Background(), request)
}

// ListTriggerTaskVersions
// 任务保存版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTriggerTaskVersionsWithContext(ctx context.Context, request *ListTriggerTaskVersionsRequest) (response *ListTriggerTaskVersionsResponse, err error) {
    if request == nil {
        request = NewListTriggerTaskVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTriggerTaskVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTriggerTaskVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTriggerTaskVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListTriggerTasksRequest() (request *ListTriggerTasksRequest) {
    request = &ListTriggerTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTriggerTasks")
    
    
    return
}

func NewListTriggerTasksResponse() (response *ListTriggerTasksResponse) {
    response = &ListTriggerTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTriggerTasks
// 查询任务分页信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTriggerTasks(request *ListTriggerTasksRequest) (response *ListTriggerTasksResponse, err error) {
    return c.ListTriggerTasksWithContext(context.Background(), request)
}

// ListTriggerTasks
// 查询任务分页信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTriggerTasksWithContext(ctx context.Context, request *ListTriggerTasksRequest) (response *ListTriggerTasksResponse, err error) {
    if request == nil {
        request = NewListTriggerTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTriggerTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTriggerTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTriggerTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListTriggerWorkflowRunsRequest() (request *ListTriggerWorkflowRunsRequest) {
    request = &ListTriggerWorkflowRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTriggerWorkflowRuns")
    
    
    return
}

func NewListTriggerWorkflowRunsResponse() (response *ListTriggerWorkflowRunsResponse) {
    response = &ListTriggerWorkflowRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTriggerWorkflowRuns
// 查询工作流运行
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListTriggerWorkflowRuns(request *ListTriggerWorkflowRunsRequest) (response *ListTriggerWorkflowRunsResponse, err error) {
    return c.ListTriggerWorkflowRunsWithContext(context.Background(), request)
}

// ListTriggerWorkflowRuns
// 查询工作流运行
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ListTriggerWorkflowRunsWithContext(ctx context.Context, request *ListTriggerWorkflowRunsRequest) (response *ListTriggerWorkflowRunsResponse, err error) {
    if request == nil {
        request = NewListTriggerWorkflowRunsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTriggerWorkflowRuns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTriggerWorkflowRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTriggerWorkflowRunsResponse()
    err = c.Send(request, response)
    return
}

func NewListTriggerWorkflowsRequest() (request *ListTriggerWorkflowsRequest) {
    request = &ListTriggerWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListTriggerWorkflows")
    
    
    return
}

func NewListTriggerWorkflowsResponse() (response *ListTriggerWorkflowsResponse) {
    response = &ListTriggerWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTriggerWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTriggerWorkflows(request *ListTriggerWorkflowsRequest) (response *ListTriggerWorkflowsResponse, err error) {
    return c.ListTriggerWorkflowsWithContext(context.Background(), request)
}

// ListTriggerWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListTriggerWorkflowsWithContext(ctx context.Context, request *ListTriggerWorkflowsRequest) (response *ListTriggerWorkflowsResponse, err error) {
    if request == nil {
        request = NewListTriggerWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListTriggerWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTriggerWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTriggerWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewListUpstreamOpsTasksRequest() (request *ListUpstreamOpsTasksRequest) {
    request = &ListUpstreamOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListUpstreamOpsTasks")
    
    
    return
}

func NewListUpstreamOpsTasksResponse() (response *ListUpstreamOpsTasksResponse) {
    response = &ListUpstreamOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUpstreamOpsTasks
// 获取任务直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamOpsTasks(request *ListUpstreamOpsTasksRequest) (response *ListUpstreamOpsTasksResponse, err error) {
    return c.ListUpstreamOpsTasksWithContext(context.Background(), request)
}

// ListUpstreamOpsTasks
// 获取任务直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamOpsTasksWithContext(ctx context.Context, request *ListUpstreamOpsTasksRequest) (response *ListUpstreamOpsTasksResponse, err error) {
    if request == nil {
        request = NewListUpstreamOpsTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListUpstreamOpsTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUpstreamOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUpstreamOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListUpstreamTaskInstancesRequest() (request *ListUpstreamTaskInstancesRequest) {
    request = &ListUpstreamTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListUpstreamTaskInstances")
    
    
    return
}

func NewListUpstreamTaskInstancesResponse() (response *ListUpstreamTaskInstancesResponse) {
    response = &ListUpstreamTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUpstreamTaskInstances
// 获取实例直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTaskInstances(request *ListUpstreamTaskInstancesRequest) (response *ListUpstreamTaskInstancesResponse, err error) {
    return c.ListUpstreamTaskInstancesWithContext(context.Background(), request)
}

// ListUpstreamTaskInstances
// 获取实例直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTaskInstancesWithContext(ctx context.Context, request *ListUpstreamTaskInstancesRequest) (response *ListUpstreamTaskInstancesResponse, err error) {
    if request == nil {
        request = NewListUpstreamTaskInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListUpstreamTaskInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUpstreamTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUpstreamTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewListUpstreamTasksRequest() (request *ListUpstreamTasksRequest) {
    request = &ListUpstreamTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListUpstreamTasks")
    
    
    return
}

func NewListUpstreamTasksResponse() (response *ListUpstreamTasksResponse) {
    response = &ListUpstreamTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUpstreamTasks
// 获取任务直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTasks(request *ListUpstreamTasksRequest) (response *ListUpstreamTasksResponse, err error) {
    return c.ListUpstreamTasksWithContext(context.Background(), request)
}

// ListUpstreamTasks
// 获取任务直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListUpstreamTasksWithContext(ctx context.Context, request *ListUpstreamTasksRequest) (response *ListUpstreamTasksResponse, err error) {
    if request == nil {
        request = NewListUpstreamTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListUpstreamTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUpstreamTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUpstreamTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListUpstreamTriggerTasksRequest() (request *ListUpstreamTriggerTasksRequest) {
    request = &ListUpstreamTriggerTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListUpstreamTriggerTasks")
    
    
    return
}

func NewListUpstreamTriggerTasksResponse() (response *ListUpstreamTriggerTasksResponse) {
    response = &ListUpstreamTriggerTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUpstreamTriggerTasks
// 获取任务直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListUpstreamTriggerTasks(request *ListUpstreamTriggerTasksRequest) (response *ListUpstreamTriggerTasksResponse, err error) {
    return c.ListUpstreamTriggerTasksWithContext(context.Background(), request)
}

// ListUpstreamTriggerTasks
// 获取任务直接上游
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListUpstreamTriggerTasksWithContext(ctx context.Context, request *ListUpstreamTriggerTasksRequest) (response *ListUpstreamTriggerTasksResponse, err error) {
    if request == nil {
        request = NewListUpstreamTriggerTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListUpstreamTriggerTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUpstreamTriggerTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUpstreamTriggerTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowFoldersRequest() (request *ListWorkflowFoldersRequest) {
    request = &ListWorkflowFoldersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListWorkflowFolders")
    
    
    return
}

func NewListWorkflowFoldersResponse() (response *ListWorkflowFoldersResponse) {
    response = &ListWorkflowFoldersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowFolders
// 查询文件夹列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowFolders(request *ListWorkflowFoldersRequest) (response *ListWorkflowFoldersResponse, err error) {
    return c.ListWorkflowFoldersWithContext(context.Background(), request)
}

// ListWorkflowFolders
// 查询文件夹列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListWorkflowFoldersWithContext(ctx context.Context, request *ListWorkflowFoldersRequest) (response *ListWorkflowFoldersResponse, err error) {
    if request == nil {
        request = NewListWorkflowFoldersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListWorkflowFolders")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowFolders require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowFoldersResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowPermissionsRequest() (request *ListWorkflowPermissionsRequest) {
    request = &ListWorkflowPermissionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListWorkflowPermissions")
    
    
    return
}

func NewListWorkflowPermissionsResponse() (response *ListWorkflowPermissionsResponse) {
    response = &ListWorkflowPermissionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflowPermissions
// 查询工作流授权权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowPermissions(request *ListWorkflowPermissionsRequest) (response *ListWorkflowPermissionsResponse, err error) {
    return c.ListWorkflowPermissionsWithContext(context.Background(), request)
}

// ListWorkflowPermissions
// 查询工作流授权权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowPermissionsWithContext(ctx context.Context, request *ListWorkflowPermissionsRequest) (response *ListWorkflowPermissionsResponse, err error) {
    if request == nil {
        request = NewListWorkflowPermissionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListWorkflowPermissions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflowPermissions require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowPermissionsResponse()
    err = c.Send(request, response)
    return
}

func NewListWorkflowsRequest() (request *ListWorkflowsRequest) {
    request = &ListWorkflowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ListWorkflows")
    
    
    return
}

func NewListWorkflowsResponse() (response *ListWorkflowsResponse) {
    response = &ListWorkflowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflows(request *ListWorkflowsRequest) (response *ListWorkflowsResponse, err error) {
    return c.ListWorkflowsWithContext(context.Background(), request)
}

// ListWorkflows
// 查询工作流列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ListWorkflowsWithContext(ctx context.Context, request *ListWorkflowsRequest) (response *ListWorkflowsResponse, err error) {
    if request == nil {
        request = NewListWorkflowsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ListWorkflows")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListWorkflows require credential")
    }

    request.SetContext(ctx)
    
    response = NewListWorkflowsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQualityRuleRequest() (request *ModifyQualityRuleRequest) {
    request = &ModifyQualityRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyQualityRule")
    
    
    return
}

func NewModifyQualityRuleResponse() (response *ModifyQualityRuleResponse) {
    response = &ModifyQualityRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQualityRule
// 更新质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) ModifyQualityRule(request *ModifyQualityRuleRequest) (response *ModifyQualityRuleResponse, err error) {
    return c.ModifyQualityRuleWithContext(context.Background(), request)
}

// ModifyQualityRule
// 更新质量规则接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) ModifyQualityRuleWithContext(ctx context.Context, request *ModifyQualityRuleRequest) (response *ModifyQualityRuleResponse, err error) {
    if request == nil {
        request = NewModifyQualityRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ModifyQualityRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQualityRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQualityRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQualityRuleGroupRequest() (request *ModifyQualityRuleGroupRequest) {
    request = &ModifyQualityRuleGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "ModifyQualityRuleGroup")
    
    
    return
}

func NewModifyQualityRuleGroupResponse() (response *ModifyQualityRuleGroupResponse) {
    response = &ModifyQualityRuleGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQualityRuleGroup
// 创建监控任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) ModifyQualityRuleGroup(request *ModifyQualityRuleGroupRequest) (response *ModifyQualityRuleGroupResponse, err error) {
    return c.ModifyQualityRuleGroupWithContext(context.Background(), request)
}

// ModifyQualityRuleGroup
// 创建监控任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"
func (c *Client) ModifyQualityRuleGroupWithContext(ctx context.Context, request *ModifyQualityRuleGroupRequest) (response *ModifyQualityRuleGroupResponse, err error) {
    if request == nil {
        request = NewModifyQualityRuleGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "ModifyQualityRuleGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQualityRuleGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQualityRuleGroupResponse()
    err = c.Send(request, response)
    return
}

func NewPauseOpsTasksAsyncRequest() (request *PauseOpsTasksAsyncRequest) {
    request = &PauseOpsTasksAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "PauseOpsTasksAsync")
    
    
    return
}

func NewPauseOpsTasksAsyncResponse() (response *PauseOpsTasksAsyncResponse) {
    response = &PauseOpsTasksAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseOpsTasksAsync
// 异步批量暂停任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseOpsTasksAsync(request *PauseOpsTasksAsyncRequest) (response *PauseOpsTasksAsyncResponse, err error) {
    return c.PauseOpsTasksAsyncWithContext(context.Background(), request)
}

// PauseOpsTasksAsync
// 异步批量暂停任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) PauseOpsTasksAsyncWithContext(ctx context.Context, request *PauseOpsTasksAsyncRequest) (response *PauseOpsTasksAsyncResponse, err error) {
    if request == nil {
        request = NewPauseOpsTasksAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "PauseOpsTasksAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseOpsTasksAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseOpsTasksAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterLineageRequest() (request *RegisterLineageRequest) {
    request = &RegisterLineageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RegisterLineage")
    
    
    return
}

func NewRegisterLineageResponse() (response *RegisterLineageResponse) {
    response = &RegisterLineageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterLineage
// RegisterLineage
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RegisterLineage(request *RegisterLineageRequest) (response *RegisterLineageResponse, err error) {
    return c.RegisterLineageWithContext(context.Background(), request)
}

// RegisterLineage
// RegisterLineage
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RegisterLineageWithContext(ctx context.Context, request *RegisterLineageRequest) (response *RegisterLineageResponse, err error) {
    if request == nil {
        request = NewRegisterLineageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RegisterLineage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterLineage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterLineageResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveMemberProjectRoleRequest() (request *RemoveMemberProjectRoleRequest) {
    request = &RemoveMemberProjectRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RemoveMemberProjectRole")
    
    
    return
}

func NewRemoveMemberProjectRoleResponse() (response *RemoveMemberProjectRoleResponse) {
    response = &RemoveMemberProjectRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveMemberProjectRole
// 删除项目用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemoveMemberProjectRole(request *RemoveMemberProjectRoleRequest) (response *RemoveMemberProjectRoleResponse, err error) {
    return c.RemoveMemberProjectRoleWithContext(context.Background(), request)
}

// RemoveMemberProjectRole
// 删除项目用户角色
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemoveMemberProjectRoleWithContext(ctx context.Context, request *RemoveMemberProjectRoleRequest) (response *RemoveMemberProjectRoleResponse, err error) {
    if request == nil {
        request = NewRemoveMemberProjectRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RemoveMemberProjectRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveMemberProjectRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveMemberProjectRoleResponse()
    err = c.Send(request, response)
    return
}

func NewRerunTaskInstancesAsyncRequest() (request *RerunTaskInstancesAsyncRequest) {
    request = &RerunTaskInstancesAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RerunTaskInstancesAsync")
    
    
    return
}

func NewRerunTaskInstancesAsyncResponse() (response *RerunTaskInstancesAsyncResponse) {
    response = &RerunTaskInstancesAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunTaskInstancesAsync
// 实例批量重跑-异步
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunTaskInstancesAsync(request *RerunTaskInstancesAsyncRequest) (response *RerunTaskInstancesAsyncResponse, err error) {
    return c.RerunTaskInstancesAsyncWithContext(context.Background(), request)
}

// RerunTaskInstancesAsync
// 实例批量重跑-异步
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RerunTaskInstancesAsyncWithContext(ctx context.Context, request *RerunTaskInstancesAsyncRequest) (response *RerunTaskInstancesAsyncResponse, err error) {
    if request == nil {
        request = NewRerunTaskInstancesAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RerunTaskInstancesAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunTaskInstancesAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunTaskInstancesAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewRerunTriggerWorkflowRunAsyncRequest() (request *RerunTriggerWorkflowRunAsyncRequest) {
    request = &RerunTriggerWorkflowRunAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RerunTriggerWorkflowRunAsync")
    
    
    return
}

func NewRerunTriggerWorkflowRunAsyncResponse() (response *RerunTriggerWorkflowRunAsyncResponse) {
    response = &RerunTriggerWorkflowRunAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RerunTriggerWorkflowRunAsync
// 重跑运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RerunTriggerWorkflowRunAsync(request *RerunTriggerWorkflowRunAsyncRequest) (response *RerunTriggerWorkflowRunAsyncResponse, err error) {
    return c.RerunTriggerWorkflowRunAsyncWithContext(context.Background(), request)
}

// RerunTriggerWorkflowRunAsync
// 重跑运行
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RerunTriggerWorkflowRunAsyncWithContext(ctx context.Context, request *RerunTriggerWorkflowRunAsyncRequest) (response *RerunTriggerWorkflowRunAsyncResponse, err error) {
    if request == nil {
        request = NewRerunTriggerWorkflowRunAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RerunTriggerWorkflowRunAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RerunTriggerWorkflowRunAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewRerunTriggerWorkflowRunAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeDataSourceAuthorizationRequest() (request *RevokeDataSourceAuthorizationRequest) {
    request = &RevokeDataSourceAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RevokeDataSourceAuthorization")
    
    
    return
}

func NewRevokeDataSourceAuthorizationResponse() (response *RevokeDataSourceAuthorizationResponse) {
    response = &RevokeDataSourceAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevokeDataSourceAuthorization
// 回收数据源权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RevokeDataSourceAuthorization(request *RevokeDataSourceAuthorizationRequest) (response *RevokeDataSourceAuthorizationResponse, err error) {
    return c.RevokeDataSourceAuthorizationWithContext(context.Background(), request)
}

// RevokeDataSourceAuthorization
// 回收数据源权限
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RevokeDataSourceAuthorizationWithContext(ctx context.Context, request *RevokeDataSourceAuthorizationRequest) (response *RevokeDataSourceAuthorizationResponse, err error) {
    if request == nil {
        request = NewRevokeDataSourceAuthorizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RevokeDataSourceAuthorization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeDataSourceAuthorization require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevokeDataSourceAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewRunSQLScriptRequest() (request *RunSQLScriptRequest) {
    request = &RunSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "RunSQLScript")
    
    
    return
}

func NewRunSQLScriptResponse() (response *RunSQLScriptResponse) {
    response = &RunSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunSQLScript
// 运行SQL脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunSQLScript(request *RunSQLScriptRequest) (response *RunSQLScriptResponse, err error) {
    return c.RunSQLScriptWithContext(context.Background(), request)
}

// RunSQLScript
// 运行SQL脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunSQLScriptWithContext(ctx context.Context, request *RunSQLScriptRequest) (response *RunSQLScriptResponse, err error) {
    if request == nil {
        request = NewRunSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "RunSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewSetSuccessTaskInstancesAsyncRequest() (request *SetSuccessTaskInstancesAsyncRequest) {
    request = &SetSuccessTaskInstancesAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SetSuccessTaskInstancesAsync")
    
    
    return
}

func NewSetSuccessTaskInstancesAsyncResponse() (response *SetSuccessTaskInstancesAsyncResponse) {
    response = &SetSuccessTaskInstancesAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetSuccessTaskInstancesAsync
// 实例批量置成功-异步
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetSuccessTaskInstancesAsync(request *SetSuccessTaskInstancesAsyncRequest) (response *SetSuccessTaskInstancesAsyncResponse, err error) {
    return c.SetSuccessTaskInstancesAsyncWithContext(context.Background(), request)
}

// SetSuccessTaskInstancesAsync
// 实例批量置成功-异步
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SetSuccessTaskInstancesAsyncWithContext(ctx context.Context, request *SetSuccessTaskInstancesAsyncRequest) (response *SetSuccessTaskInstancesAsyncResponse, err error) {
    if request == nil {
        request = NewSetSuccessTaskInstancesAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "SetSuccessTaskInstancesAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetSuccessTaskInstancesAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetSuccessTaskInstancesAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewStartOpsTasksRequest() (request *StartOpsTasksRequest) {
    request = &StartOpsTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StartOpsTasks")
    
    
    return
}

func NewStartOpsTasksResponse() (response *StartOpsTasksResponse) {
    response = &StartOpsTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartOpsTasks
// 异步批量启动任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StartOpsTasks(request *StartOpsTasksRequest) (response *StartOpsTasksResponse, err error) {
    return c.StartOpsTasksWithContext(context.Background(), request)
}

// StartOpsTasks
// 异步批量启动任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StartOpsTasksWithContext(ctx context.Context, request *StartOpsTasksRequest) (response *StartOpsTasksResponse, err error) {
    if request == nil {
        request = NewStartOpsTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "StartOpsTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartOpsTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartOpsTasksResponse()
    err = c.Send(request, response)
    return
}

func NewStopOpsTasksAsyncRequest() (request *StopOpsTasksAsyncRequest) {
    request = &StopOpsTasksAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StopOpsTasksAsync")
    
    
    return
}

func NewStopOpsTasksAsyncResponse() (response *StopOpsTasksAsyncResponse) {
    response = &StopOpsTasksAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopOpsTasksAsync
// 异步批量下线任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopOpsTasksAsync(request *StopOpsTasksAsyncRequest) (response *StopOpsTasksAsyncResponse, err error) {
    return c.StopOpsTasksAsyncWithContext(context.Background(), request)
}

// StopOpsTasksAsync
// 异步批量下线任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopOpsTasksAsyncWithContext(ctx context.Context, request *StopOpsTasksAsyncRequest) (response *StopOpsTasksAsyncResponse, err error) {
    if request == nil {
        request = NewStopOpsTasksAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "StopOpsTasksAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopOpsTasksAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopOpsTasksAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewStopSQLScriptRunRequest() (request *StopSQLScriptRunRequest) {
    request = &StopSQLScriptRunRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "StopSQLScriptRun")
    
    
    return
}

func NewStopSQLScriptRunResponse() (response *StopSQLScriptRunResponse) {
    response = &StopSQLScriptRunResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSQLScriptRun
// 停止运行SQL脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopSQLScriptRun(request *StopSQLScriptRunRequest) (response *StopSQLScriptRunResponse, err error) {
    return c.StopSQLScriptRunWithContext(context.Background(), request)
}

// StopSQLScriptRun
// 停止运行SQL脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopSQLScriptRunWithContext(ctx context.Context, request *StopSQLScriptRunRequest) (response *StopSQLScriptRunResponse, err error) {
    if request == nil {
        request = NewStopSQLScriptRunRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "StopSQLScriptRun")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSQLScriptRun require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSQLScriptRunResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTaskRequest() (request *SubmitTaskRequest) {
    request = &SubmitTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitTask")
    
    
    return
}

func NewSubmitTaskResponse() (response *SubmitTaskResponse) {
    response = &SubmitTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTask
// 提交任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitTask(request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    return c.SubmitTaskWithContext(context.Background(), request)
}

// SubmitTask
// 提交任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitTaskWithContext(ctx context.Context, request *SubmitTaskRequest) (response *SubmitTaskResponse, err error) {
    if request == nil {
        request = NewSubmitTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "SubmitTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTriggerTaskRequest() (request *SubmitTriggerTaskRequest) {
    request = &SubmitTriggerTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "SubmitTriggerTask")
    
    
    return
}

func NewSubmitTriggerTaskResponse() (response *SubmitTriggerTaskResponse) {
    response = &SubmitTriggerTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTriggerTask
// 提交工作流调度任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitTriggerTask(request *SubmitTriggerTaskRequest) (response *SubmitTriggerTaskResponse, err error) {
    return c.SubmitTriggerTaskWithContext(context.Background(), request)
}

// SubmitTriggerTask
// 提交工作流调度任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SubmitTriggerTaskWithContext(ctx context.Context, request *SubmitTriggerTaskRequest) (response *SubmitTriggerTaskResponse, err error) {
    if request == nil {
        request = NewSubmitTriggerTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "SubmitTriggerTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTriggerTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTriggerTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCodeFileRequest() (request *UpdateCodeFileRequest) {
    request = &UpdateCodeFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateCodeFile")
    
    
    return
}

func NewUpdateCodeFileResponse() (response *UpdateCodeFileResponse) {
    response = &UpdateCodeFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCodeFile
// 更新代码文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFile(request *UpdateCodeFileRequest) (response *UpdateCodeFileResponse, err error) {
    return c.UpdateCodeFileWithContext(context.Background(), request)
}

// UpdateCodeFile
// 更新代码文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFileWithContext(ctx context.Context, request *UpdateCodeFileRequest) (response *UpdateCodeFileResponse, err error) {
    if request == nil {
        request = NewUpdateCodeFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateCodeFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCodeFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCodeFileResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCodeFolderRequest() (request *UpdateCodeFolderRequest) {
    request = &UpdateCodeFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateCodeFolder")
    
    
    return
}

func NewUpdateCodeFolderResponse() (response *UpdateCodeFolderResponse) {
    response = &UpdateCodeFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCodeFolder
// 重命名代码文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFolder(request *UpdateCodeFolderRequest) (response *UpdateCodeFolderResponse, err error) {
    return c.UpdateCodeFolderWithContext(context.Background(), request)
}

// UpdateCodeFolder
// 重命名代码文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateCodeFolderWithContext(ctx context.Context, request *UpdateCodeFolderRequest) (response *UpdateCodeFolderResponse, err error) {
    if request == nil {
        request = NewUpdateCodeFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateCodeFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCodeFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCodeFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDataSourceRequest() (request *UpdateDataSourceRequest) {
    request = &UpdateDataSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateDataSource")
    
    
    return
}

func NewUpdateDataSourceResponse() (response *UpdateDataSourceResponse) {
    response = &UpdateDataSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDataSource
// 该接口用于更新数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateDataSource(request *UpdateDataSourceRequest) (response *UpdateDataSourceResponse, err error) {
    return c.UpdateDataSourceWithContext(context.Background(), request)
}

// UpdateDataSource
// 该接口用于更新数据源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateDataSourceWithContext(ctx context.Context, request *UpdateDataSourceRequest) (response *UpdateDataSourceResponse, err error) {
    if request == nil {
        request = NewUpdateDataSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateDataSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDataSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDataSourceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOpsAlarmRuleRequest() (request *UpdateOpsAlarmRuleRequest) {
    request = &UpdateOpsAlarmRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateOpsAlarmRule")
    
    
    return
}

func NewUpdateOpsAlarmRuleResponse() (response *UpdateOpsAlarmRuleResponse) {
    response = &UpdateOpsAlarmRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOpsAlarmRule
// 修改告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsAlarmRule(request *UpdateOpsAlarmRuleRequest) (response *UpdateOpsAlarmRuleResponse, err error) {
    return c.UpdateOpsAlarmRuleWithContext(context.Background(), request)
}

// UpdateOpsAlarmRule
// 修改告警规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsAlarmRuleWithContext(ctx context.Context, request *UpdateOpsAlarmRuleRequest) (response *UpdateOpsAlarmRuleResponse, err error) {
    if request == nil {
        request = NewUpdateOpsAlarmRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateOpsAlarmRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOpsAlarmRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOpsAlarmRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOpsTasksOwnerRequest() (request *UpdateOpsTasksOwnerRequest) {
    request = &UpdateOpsTasksOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateOpsTasksOwner")
    
    
    return
}

func NewUpdateOpsTasksOwnerResponse() (response *UpdateOpsTasksOwnerResponse) {
    response = &UpdateOpsTasksOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOpsTasksOwner
// 修改任务负责人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsTasksOwner(request *UpdateOpsTasksOwnerRequest) (response *UpdateOpsTasksOwnerResponse, err error) {
    return c.UpdateOpsTasksOwnerWithContext(context.Background(), request)
}

// UpdateOpsTasksOwner
// 修改任务负责人
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsTasksOwnerWithContext(ctx context.Context, request *UpdateOpsTasksOwnerRequest) (response *UpdateOpsTasksOwnerResponse, err error) {
    if request == nil {
        request = NewUpdateOpsTasksOwnerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateOpsTasksOwner")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOpsTasksOwner require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOpsTasksOwnerResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOpsTriggerTasksOwnerRequest() (request *UpdateOpsTriggerTasksOwnerRequest) {
    request = &UpdateOpsTriggerTasksOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateOpsTriggerTasksOwner")
    
    
    return
}

func NewUpdateOpsTriggerTasksOwnerResponse() (response *UpdateOpsTriggerTasksOwnerResponse) {
    response = &UpdateOpsTriggerTasksOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOpsTriggerTasksOwner
// 查询任务执行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsTriggerTasksOwner(request *UpdateOpsTriggerTasksOwnerRequest) (response *UpdateOpsTriggerTasksOwnerResponse, err error) {
    return c.UpdateOpsTriggerTasksOwnerWithContext(context.Background(), request)
}

// UpdateOpsTriggerTasksOwner
// 查询任务执行详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"
//  INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"
//  INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"
func (c *Client) UpdateOpsTriggerTasksOwnerWithContext(ctx context.Context, request *UpdateOpsTriggerTasksOwnerRequest) (response *UpdateOpsTriggerTasksOwnerResponse, err error) {
    if request == nil {
        request = NewUpdateOpsTriggerTasksOwnerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateOpsTriggerTasksOwner")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOpsTriggerTasksOwner require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOpsTriggerTasksOwnerResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProjectRequest() (request *UpdateProjectRequest) {
    request = &UpdateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateProject")
    
    
    return
}

func NewUpdateProjectResponse() (response *UpdateProjectResponse) {
    response = &UpdateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateProject
// 修改项目基础信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENAMEDUPLICATION = "InvalidParameter.WorkspaceNameDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateProject(request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
    return c.UpdateProjectWithContext(context.Background(), request)
}

// UpdateProject
// 修改项目基础信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_WORKSPACENAMEDUPLICATION = "InvalidParameter.WorkspaceNameDuplication"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) UpdateProjectWithContext(ctx context.Context, request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
    if request == nil {
        request = NewUpdateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateResourceFileRequest() (request *UpdateResourceFileRequest) {
    request = &UpdateResourceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateResourceFile")
    
    
    return
}

func NewUpdateResourceFileResponse() (response *UpdateResourceFileResponse) {
    response = &UpdateResourceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateResourceFile
// 更新资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFile(request *UpdateResourceFileRequest) (response *UpdateResourceFileResponse, err error) {
    return c.UpdateResourceFileWithContext(context.Background(), request)
}

// UpdateResourceFile
// 更新资源文件
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFileWithContext(ctx context.Context, request *UpdateResourceFileRequest) (response *UpdateResourceFileResponse, err error) {
    if request == nil {
        request = NewUpdateResourceFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateResourceFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateResourceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateResourceFileResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateResourceFolderRequest() (request *UpdateResourceFolderRequest) {
    request = &UpdateResourceFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateResourceFolder")
    
    
    return
}

func NewUpdateResourceFolderResponse() (response *UpdateResourceFolderResponse) {
    response = &UpdateResourceFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateResourceFolder
// 更新资源文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFolder(request *UpdateResourceFolderRequest) (response *UpdateResourceFolderResponse, err error) {
    return c.UpdateResourceFolderWithContext(context.Background(), request)
}

// UpdateResourceFolder
// 更新资源文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceFolderWithContext(ctx context.Context, request *UpdateResourceFolderRequest) (response *UpdateResourceFolderResponse, err error) {
    if request == nil {
        request = NewUpdateResourceFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateResourceFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateResourceFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateResourceFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateResourceGroupRequest() (request *UpdateResourceGroupRequest) {
    request = &UpdateResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateResourceGroup")
    
    
    return
}

func NewUpdateResourceGroupResponse() (response *UpdateResourceGroupResponse) {
    response = &UpdateResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateResourceGroup
// 该接口用于变配/续费资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceGroup(request *UpdateResourceGroupRequest) (response *UpdateResourceGroupResponse, err error) {
    return c.UpdateResourceGroupWithContext(context.Background(), request)
}

// UpdateResourceGroup
// 该接口用于变配/续费资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateResourceGroupWithContext(ctx context.Context, request *UpdateResourceGroupRequest) (response *UpdateResourceGroupResponse, err error) {
    if request == nil {
        request = NewUpdateResourceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateResourceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSQLFolderRequest() (request *UpdateSQLFolderRequest) {
    request = &UpdateSQLFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateSQLFolder")
    
    
    return
}

func NewUpdateSQLFolderResponse() (response *UpdateSQLFolderResponse) {
    response = &UpdateSQLFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSQLFolder
// 重命名SQL文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLFolder(request *UpdateSQLFolderRequest) (response *UpdateSQLFolderResponse, err error) {
    return c.UpdateSQLFolderWithContext(context.Background(), request)
}

// UpdateSQLFolder
// 重命名SQL文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLFolderWithContext(ctx context.Context, request *UpdateSQLFolderRequest) (response *UpdateSQLFolderResponse, err error) {
    if request == nil {
        request = NewUpdateSQLFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateSQLFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSQLFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSQLFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSQLScriptRequest() (request *UpdateSQLScriptRequest) {
    request = &UpdateSQLScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateSQLScript")
    
    
    return
}

func NewUpdateSQLScriptResponse() (response *UpdateSQLScriptResponse) {
    response = &UpdateSQLScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSQLScript
// 保存探索脚本内容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLScript(request *UpdateSQLScriptRequest) (response *UpdateSQLScriptResponse, err error) {
    return c.UpdateSQLScriptWithContext(context.Background(), request)
}

// UpdateSQLScript
// 保存探索脚本内容
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateSQLScriptWithContext(ctx context.Context, request *UpdateSQLScriptRequest) (response *UpdateSQLScriptResponse, err error) {
    if request == nil {
        request = NewUpdateSQLScriptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateSQLScript")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSQLScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSQLScriptResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTaskRequest() (request *UpdateTaskRequest) {
    request = &UpdateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTask")
    
    
    return
}

func NewUpdateTaskResponse() (response *UpdateTaskResponse) {
    response = &UpdateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTask
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTask(request *UpdateTaskRequest) (response *UpdateTaskResponse, err error) {
    return c.UpdateTaskWithContext(context.Background(), request)
}

// UpdateTask
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTaskWithContext(ctx context.Context, request *UpdateTaskRequest) (response *UpdateTaskResponse, err error) {
    if request == nil {
        request = NewUpdateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTaskFolderRequest() (request *UpdateTaskFolderRequest) {
    request = &UpdateTaskFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTaskFolder")
    
    
    return
}

func NewUpdateTaskFolderResponse() (response *UpdateTaskFolderResponse) {
    response = &UpdateTaskFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTaskFolder
// 更新任务文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTaskFolder(request *UpdateTaskFolderRequest) (response *UpdateTaskFolderResponse, err error) {
    return c.UpdateTaskFolderWithContext(context.Background(), request)
}

// UpdateTaskFolder
// 更新任务文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTaskFolderWithContext(ctx context.Context, request *UpdateTaskFolderRequest) (response *UpdateTaskFolderResponse, err error) {
    if request == nil {
        request = NewUpdateTaskFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTaskFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTaskFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTaskFolderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTaskPartiallyRequest() (request *UpdateTaskPartiallyRequest) {
    request = &UpdateTaskPartiallyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTaskPartially")
    
    
    return
}

func NewUpdateTaskPartiallyResponse() (response *UpdateTaskPartiallyResponse) {
    response = &UpdateTaskPartiallyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTaskPartially
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTaskPartially(request *UpdateTaskPartiallyRequest) (response *UpdateTaskPartiallyResponse, err error) {
    return c.UpdateTaskPartiallyWithContext(context.Background(), request)
}

// UpdateTaskPartially
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTaskPartiallyWithContext(ctx context.Context, request *UpdateTaskPartiallyRequest) (response *UpdateTaskPartiallyResponse, err error) {
    if request == nil {
        request = NewUpdateTaskPartiallyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTaskPartially")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTaskPartially require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTaskPartiallyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTriggerTaskRequest() (request *UpdateTriggerTaskRequest) {
    request = &UpdateTriggerTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTriggerTask")
    
    
    return
}

func NewUpdateTriggerTaskResponse() (response *UpdateTriggerTaskResponse) {
    response = &UpdateTriggerTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTriggerTask
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTriggerTask(request *UpdateTriggerTaskRequest) (response *UpdateTriggerTaskResponse, err error) {
    return c.UpdateTriggerTaskWithContext(context.Background(), request)
}

// UpdateTriggerTask
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTriggerTaskWithContext(ctx context.Context, request *UpdateTriggerTaskRequest) (response *UpdateTriggerTaskResponse, err error) {
    if request == nil {
        request = NewUpdateTriggerTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTriggerTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTriggerTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTriggerTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTriggerTaskPartiallyRequest() (request *UpdateTriggerTaskPartiallyRequest) {
    request = &UpdateTriggerTaskPartiallyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTriggerTaskPartially")
    
    
    return
}

func NewUpdateTriggerTaskPartiallyResponse() (response *UpdateTriggerTaskPartiallyResponse) {
    response = &UpdateTriggerTaskPartiallyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTriggerTaskPartially
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTriggerTaskPartially(request *UpdateTriggerTaskPartiallyRequest) (response *UpdateTriggerTaskPartiallyResponse, err error) {
    return c.UpdateTriggerTaskPartiallyWithContext(context.Background(), request)
}

// UpdateTriggerTaskPartially
// 更新任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTriggerTaskPartiallyWithContext(ctx context.Context, request *UpdateTriggerTaskPartiallyRequest) (response *UpdateTriggerTaskPartiallyResponse, err error) {
    if request == nil {
        request = NewUpdateTriggerTaskPartiallyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTriggerTaskPartially")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTriggerTaskPartially require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTriggerTaskPartiallyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTriggerWorkflowRequest() (request *UpdateTriggerWorkflowRequest) {
    request = &UpdateTriggerWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTriggerWorkflow")
    
    
    return
}

func NewUpdateTriggerWorkflowResponse() (response *UpdateTriggerWorkflowResponse) {
    response = &UpdateTriggerWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTriggerWorkflow
// 更新工作流（包括工作流基本信息与工作流参数）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateTriggerWorkflow(request *UpdateTriggerWorkflowRequest) (response *UpdateTriggerWorkflowResponse, err error) {
    return c.UpdateTriggerWorkflowWithContext(context.Background(), request)
}

// UpdateTriggerWorkflow
// 更新工作流（包括工作流基本信息与工作流参数）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateTriggerWorkflowWithContext(ctx context.Context, request *UpdateTriggerWorkflowRequest) (response *UpdateTriggerWorkflowResponse, err error) {
    if request == nil {
        request = NewUpdateTriggerWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTriggerWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTriggerWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTriggerWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTriggerWorkflowPartiallyRequest() (request *UpdateTriggerWorkflowPartiallyRequest) {
    request = &UpdateTriggerWorkflowPartiallyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateTriggerWorkflowPartially")
    
    
    return
}

func NewUpdateTriggerWorkflowPartiallyResponse() (response *UpdateTriggerWorkflowPartiallyResponse) {
    response = &UpdateTriggerWorkflowPartiallyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTriggerWorkflowPartially
// 更新工作流（包括工作流基本信息与工作流参数）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateTriggerWorkflowPartially(request *UpdateTriggerWorkflowPartiallyRequest) (response *UpdateTriggerWorkflowPartiallyResponse, err error) {
    return c.UpdateTriggerWorkflowPartiallyWithContext(context.Background(), request)
}

// UpdateTriggerWorkflowPartially
// 更新工作流（包括工作流基本信息与工作流参数）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateTriggerWorkflowPartiallyWithContext(ctx context.Context, request *UpdateTriggerWorkflowPartiallyRequest) (response *UpdateTriggerWorkflowPartiallyResponse, err error) {
    if request == nil {
        request = NewUpdateTriggerWorkflowPartiallyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateTriggerWorkflowPartially")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTriggerWorkflowPartially require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTriggerWorkflowPartiallyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkflowRequest() (request *UpdateWorkflowRequest) {
    request = &UpdateWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateWorkflow")
    
    
    return
}

func NewUpdateWorkflowResponse() (response *UpdateWorkflowResponse) {
    response = &UpdateWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkflow
// 更新工作流（包括工作流基本信息与工作流参数）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    return c.UpdateWorkflowWithContext(context.Background(), request)
}

// UpdateWorkflow
// 更新工作流（包括工作流基本信息与工作流参数）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UpdateWorkflowWithContext(ctx context.Context, request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
    if request == nil {
        request = NewUpdateWorkflowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateWorkflow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateWorkflowFolderRequest() (request *UpdateWorkflowFolderRequest) {
    request = &UpdateWorkflowFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wedata", APIVersion, "UpdateWorkflowFolder")
    
    
    return
}

func NewUpdateWorkflowFolderResponse() (response *UpdateWorkflowFolderResponse) {
    response = &UpdateWorkflowFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateWorkflowFolder
// 更新工作流文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateWorkflowFolder(request *UpdateWorkflowFolderRequest) (response *UpdateWorkflowFolderResponse, err error) {
    return c.UpdateWorkflowFolderWithContext(context.Background(), request)
}

// UpdateWorkflowFolder
// 更新工作流文件夹
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateWorkflowFolderWithContext(ctx context.Context, request *UpdateWorkflowFolderRequest) (response *UpdateWorkflowFolderResponse, err error) {
    if request == nil {
        request = NewUpdateWorkflowFolderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wedata", APIVersion, "UpdateWorkflowFolder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateWorkflowFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateWorkflowFolderResponse()
    err = c.Send(request, response)
    return
}
