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

package v20210420

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-20"

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


func NewAddAccountToAccountGroupRequest() (request *AddAccountToAccountGroupRequest) {
    request = &AddAccountToAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "AddAccountToAccountGroup")
    
    
    return
}

func NewAddAccountToAccountGroupResponse() (response *AddAccountToAccountGroupResponse) {
    response = &AddAccountToAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddAccountToAccountGroup
// 账号组添加账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTALREADYEXISTEDINACCOUNTGROUP = "FailedOperation.AccountAlreadyExistedInAccountGroup"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_ACCOUNTIDSCANNOTBEEMPTY = "InvalidParameterValue.AccountIdsCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) AddAccountToAccountGroup(request *AddAccountToAccountGroupRequest) (response *AddAccountToAccountGroupResponse, err error) {
    return c.AddAccountToAccountGroupWithContext(context.Background(), request)
}

// AddAccountToAccountGroup
// 账号组添加账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTALREADYEXISTEDINACCOUNTGROUP = "FailedOperation.AccountAlreadyExistedInAccountGroup"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_ACCOUNTIDSCANNOTBEEMPTY = "InvalidParameterValue.AccountIdsCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) AddAccountToAccountGroupWithContext(ctx context.Context, request *AddAccountToAccountGroupRequest) (response *AddAccountToAccountGroupResponse, err error) {
    if request == nil {
        request = NewAddAccountToAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddAccountToAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddAccountToAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserToUserGroupRequest() (request *AddUserToUserGroupRequest) {
    request = &AddUserToUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "AddUserToUserGroup")
    
    
    return
}

func NewAddUserToUserGroupResponse() (response *AddUserToUserGroupResponse) {
    response = &AddUserToUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddUserToUserGroup
// 加入用户到用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) AddUserToUserGroup(request *AddUserToUserGroupRequest) (response *AddUserToUserGroupResponse, err error) {
    return c.AddUserToUserGroupWithContext(context.Background(), request)
}

// AddUserToUserGroup
// 加入用户到用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) AddUserToUserGroupWithContext(ctx context.Context, request *AddUserToUserGroupRequest) (response *AddUserToUserGroupResponse, err error) {
    if request == nil {
        request = NewAddUserToUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserToUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserToUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountGroupRequest() (request *CreateAccountGroupRequest) {
    request = &CreateAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateAccountGroup")
    
    
    return
}

func NewCreateAccountGroupResponse() (response *CreateAccountGroupResponse) {
    response = &CreateAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccountGroup
// 创建账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEALREADYEXISTS = "FailedOperation.AccountGroupNameAlreadyExists"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAccountGroup(request *CreateAccountGroupRequest) (response *CreateAccountGroupResponse, err error) {
    return c.CreateAccountGroupWithContext(context.Background(), request)
}

// CreateAccountGroup
// 创建账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEALREADYEXISTS = "FailedOperation.AccountGroupNameAlreadyExists"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAccountGroupWithContext(ctx context.Context, request *CreateAccountGroupRequest) (response *CreateAccountGroupResponse, err error) {
    if request == nil {
        request = NewCreateAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppAccountRequest() (request *CreateAppAccountRequest) {
    request = &CreateAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateAppAccount")
    
    
    return
}

func NewCreateAppAccountResponse() (response *CreateAppAccountResponse) {
    response = &CreateAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAppAccount
// 创建应用账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNAMEALREADYEXISTS = "FailedOperation.AccountNameAlreadyExists"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAppAccount(request *CreateAppAccountRequest) (response *CreateAppAccountResponse, err error) {
    return c.CreateAppAccountWithContext(context.Background(), request)
}

// CreateAppAccount
// 创建应用账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNAMEALREADYEXISTS = "FailedOperation.AccountNameAlreadyExists"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateAppAccountWithContext(ctx context.Context, request *CreateAppAccountRequest) (response *CreateAppAccountResponse, err error) {
    if request == nil {
        request = NewCreateAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrgNodeRequest() (request *CreateOrgNodeRequest) {
    request = &CreateOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateOrgNode")
    
    
    return
}

func NewCreateOrgNodeResponse() (response *CreateOrgNodeResponse) {
    response = &CreateOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOrgNode
// 新建一个机构节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CREATEORGNODEERROR = "FailedOperation.CreateOrgNodeError"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) CreateOrgNode(request *CreateOrgNodeRequest) (response *CreateOrgNodeResponse, err error) {
    return c.CreateOrgNodeWithContext(context.Background(), request)
}

// CreateOrgNode
// 新建一个机构节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CREATEORGNODEERROR = "FailedOperation.CreateOrgNodeError"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) CreateOrgNodeWithContext(ctx context.Context, request *CreateOrgNodeRequest) (response *CreateOrgNodeResponse, err error) {
    if request == nil {
        request = NewCreateOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUser
// 新建一个用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEUSERFAILURE = "FailedOperation.CreateUserFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LIMITQUOTANOTENOUGH = "FailedOperation.LimitQuotaNotEnough"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_USEREMAILALREADYEXISTS = "FailedOperation.UserEmailAlreadyExists"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 新建一个用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEUSERFAILURE = "FailedOperation.CreateUserFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LIMITQUOTANOTENOUGH = "FailedOperation.LimitQuotaNotEnough"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_USEREMAILALREADYEXISTS = "FailedOperation.UserEmailAlreadyExists"
//  FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserGroupRequest() (request *CreateUserGroupRequest) {
    request = &CreateUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "CreateUserGroup")
    
    
    return
}

func NewCreateUserGroupResponse() (response *CreateUserGroupResponse) {
    response = &CreateUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserGroup
// 新建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEUSERGROUPERROR = "FailedOperation.CreateUserGroupError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    return c.CreateUserGroupWithContext(context.Background(), request)
}

// CreateUserGroup
// 新建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEUSERGROUPERROR = "FailedOperation.CreateUserGroupError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) CreateUserGroupWithContext(ctx context.Context, request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    if request == nil {
        request = NewCreateUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountGroupRequest() (request *DeleteAccountGroupRequest) {
    request = &DeleteAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteAccountGroup")
    
    
    return
}

func NewDeleteAccountGroupResponse() (response *DeleteAccountGroupResponse) {
    response = &DeleteAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccountGroup
// 删除账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) DeleteAccountGroup(request *DeleteAccountGroupRequest) (response *DeleteAccountGroupResponse, err error) {
    return c.DeleteAccountGroupWithContext(context.Background(), request)
}

// DeleteAccountGroup
// 删除账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) DeleteAccountGroupWithContext(ctx context.Context, request *DeleteAccountGroupRequest) (response *DeleteAccountGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppAccountRequest() (request *DeleteAppAccountRequest) {
    request = &DeleteAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteAppAccount")
    
    
    return
}

func NewDeleteAppAccountResponse() (response *DeleteAppAccountResponse) {
    response = &DeleteAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAppAccount
// 删除应用账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_ACCOUNTIDSCANNOTBEEMPTY = "InvalidParameterValue.AccountIdsCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) DeleteAppAccount(request *DeleteAppAccountRequest) (response *DeleteAppAccountResponse, err error) {
    return c.DeleteAppAccountWithContext(context.Background(), request)
}

// DeleteAppAccount
// 删除应用账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_ACCOUNTIDSCANNOTBEEMPTY = "InvalidParameterValue.AccountIdsCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) DeleteAppAccountWithContext(ctx context.Context, request *DeleteAppAccountRequest) (response *DeleteAppAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrgNodeRequest() (request *DeleteOrgNodeRequest) {
    request = &DeleteOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteOrgNode")
    
    
    return
}

func NewDeleteOrgNodeResponse() (response *DeleteOrgNodeResponse) {
    response = &DeleteOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteOrgNode
// 删除一个机构节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCanNotBeDeleted"
//  FAILEDOPERATION_DEFAULTORGNODECANNOTBEDELETED = "FailedOperation.DefaultOrgNodeCanNotBeDeleted"
//  FAILEDOPERATION_DELETEORGNODEERROR = "FailedOperation.DeleteOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCanNotBeDeleted"
//  FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCanNotBeDeleted"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteOrgNode(request *DeleteOrgNodeRequest) (response *DeleteOrgNodeResponse, err error) {
    return c.DeleteOrgNodeWithContext(context.Background(), request)
}

// DeleteOrgNode
// 删除一个机构节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCanNotBeDeleted"
//  FAILEDOPERATION_DEFAULTORGNODECANNOTBEDELETED = "FailedOperation.DefaultOrgNodeCanNotBeDeleted"
//  FAILEDOPERATION_DELETEORGNODEERROR = "FailedOperation.DeleteOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCanNotBeDeleted"
//  FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCanNotBeDeleted"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteOrgNodeWithContext(ctx context.Context, request *DeleteOrgNodeRequest) (response *DeleteOrgNodeResponse, err error) {
    if request == nil {
        request = NewDeleteOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUser
// 通过用户名或用户 id 删除用户。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERERROR = "FailedOperation.DeleteUserError"
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// 通过用户名或用户 id 删除用户。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERERROR = "FailedOperation.DeleteUserError"
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserGroupRequest() (request *DeleteUserGroupRequest) {
    request = &DeleteUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteUserGroup")
    
    
    return
}

func NewDeleteUserGroupResponse() (response *DeleteUserGroupResponse) {
    response = &DeleteUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUserGroup
// 删除一个用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERGROUPERROR = "FailedOperation.DeleteUserGroupError"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (response *DeleteUserGroupResponse, err error) {
    return c.DeleteUserGroupWithContext(context.Background(), request)
}

// DeleteUserGroup
// 删除一个用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERGROUPERROR = "FailedOperation.DeleteUserGroupError"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUserGroupWithContext(ctx context.Context, request *DeleteUserGroupRequest) (response *DeleteUserGroupResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUsers
// 批量删除当前节点下的用户。如果出现个别用户删除错误，将不影响其余被勾选用户被删除的操作，同时提示未被删除的用户名称/用户ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERERROR = "FailedOperation.DeleteUserError"
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// 批量删除当前节点下的用户。如果出现个别用户删除错误，将不影响其余被勾选用户被删除的操作，同时提示未被删除的用户名称/用户ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERERROR = "FailedOperation.DeleteUserError"
//  FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountGroupRequest() (request *DescribeAccountGroupRequest) {
    request = &DescribeAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeAccountGroup")
    
    
    return
}

func NewDescribeAccountGroupResponse() (response *DescribeAccountGroupResponse) {
    response = &DescribeAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountGroup
// 查询账号组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAccountGroup(request *DescribeAccountGroupRequest) (response *DescribeAccountGroupResponse, err error) {
    return c.DescribeAccountGroupWithContext(context.Background(), request)
}

// DescribeAccountGroup
// 查询账号组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAccountGroupWithContext(ctx context.Context, request *DescribeAccountGroupRequest) (response *DescribeAccountGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppAccountRequest() (request *DescribeAppAccountRequest) {
    request = &DescribeAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeAppAccount")
    
    
    return
}

func NewDescribeAppAccountResponse() (response *DescribeAppAccountResponse) {
    response = &DescribeAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppAccount
// 查询应用账号列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAppAccount(request *DescribeAppAccountRequest) (response *DescribeAppAccountResponse, err error) {
    return c.DescribeAppAccountWithContext(context.Background(), request)
}

// DescribeAppAccount
// 查询应用账号列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeAppAccountWithContext(ctx context.Context, request *DescribeAppAccountRequest) (response *DescribeAppAccountResponse, err error) {
    if request == nil {
        request = NewDescribeAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
    request = &DescribeApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeApplication")
    
    
    return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
    response = &DescribeApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplication
// 获取一个应用的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    return c.DescribeApplicationWithContext(context.Background(), request)
}

// DescribeApplication
// 获取一个应用的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeApplicationWithContext(ctx context.Context, request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrgNodeRequest() (request *DescribeOrgNodeRequest) {
    request = &DescribeOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeOrgNode")
    
    
    return
}

func NewDescribeOrgNodeResponse() (response *DescribeOrgNodeResponse) {
    response = &DescribeOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrgNode
// 根据机构节点ID读取机构节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeOrgNode(request *DescribeOrgNodeRequest) (response *DescribeOrgNodeResponse, err error) {
    return c.DescribeOrgNodeWithContext(context.Background(), request)
}

// DescribeOrgNode
// 根据机构节点ID读取机构节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeOrgNodeWithContext(ctx context.Context, request *DescribeOrgNodeRequest) (response *DescribeOrgNodeResponse, err error) {
    if request == nil {
        request = NewDescribeOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrgResourcesAuthorizationRequest() (request *DescribeOrgResourcesAuthorizationRequest) {
    request = &DescribeOrgResourcesAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeOrgResourcesAuthorization")
    
    
    return
}

func NewDescribeOrgResourcesAuthorizationResponse() (response *DescribeOrgResourcesAuthorizationResponse) {
    response = &DescribeOrgResourcesAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrgResourcesAuthorization
// 查询指定机构下的资源授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeOrgResourcesAuthorization(request *DescribeOrgResourcesAuthorizationRequest) (response *DescribeOrgResourcesAuthorizationResponse, err error) {
    return c.DescribeOrgResourcesAuthorizationWithContext(context.Background(), request)
}

// DescribeOrgResourcesAuthorization
// 查询指定机构下的资源授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeOrgResourcesAuthorizationWithContext(ctx context.Context, request *DescribeOrgResourcesAuthorizationRequest) (response *DescribeOrgResourcesAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrgResourcesAuthorizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrgResourcesAuthorization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrgResourcesAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublicKeyRequest() (request *DescribePublicKeyRequest) {
    request = &DescribePublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribePublicKey")
    
    
    return
}

func NewDescribePublicKeyResponse() (response *DescribePublicKeyResponse) {
    response = &DescribePublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublicKey
// 获取JWT公钥信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribePublicKey(request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
    return c.DescribePublicKeyWithContext(context.Background(), request)
}

// DescribePublicKey
// 获取JWT公钥信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribePublicKeyWithContext(ctx context.Context, request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
    if request == nil {
        request = NewDescribePublicKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublicKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupRequest() (request *DescribeUserGroupRequest) {
    request = &DescribeUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserGroup")
    
    
    return
}

func NewDescribeUserGroupResponse() (response *DescribeUserGroupResponse) {
    response = &DescribeUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserGroup
// 获取用户组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserGroup(request *DescribeUserGroupRequest) (response *DescribeUserGroupResponse, err error) {
    return c.DescribeUserGroupWithContext(context.Background(), request)
}

// DescribeUserGroup
// 获取用户组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserGroupWithContext(ctx context.Context, request *DescribeUserGroupRequest) (response *DescribeUserGroupResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserGroupResourcesAuthorizationRequest() (request *DescribeUserGroupResourcesAuthorizationRequest) {
    request = &DescribeUserGroupResourcesAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserGroupResourcesAuthorization")
    
    
    return
}

func NewDescribeUserGroupResourcesAuthorizationResponse() (response *DescribeUserGroupResourcesAuthorizationResponse) {
    response = &DescribeUserGroupResourcesAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserGroupResourcesAuthorization
// 查询指定用户组下的资源授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserGroupResourcesAuthorization(request *DescribeUserGroupResourcesAuthorizationRequest) (response *DescribeUserGroupResourcesAuthorizationResponse, err error) {
    return c.DescribeUserGroupResourcesAuthorizationWithContext(context.Background(), request)
}

// DescribeUserGroupResourcesAuthorization
// 查询指定用户组下的资源授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserGroupResourcesAuthorizationWithContext(ctx context.Context, request *DescribeUserGroupResourcesAuthorizationRequest) (response *DescribeUserGroupResourcesAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupResourcesAuthorizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserGroupResourcesAuthorization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserGroupResourcesAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
    request = &DescribeUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserInfo")
    
    
    return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
    response = &DescribeUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserInfo
// 通过用户名或用户 id 搜索用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_USERIDCANNOTBEEMPTY = "InvalidParameterValue.UserIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERNAMECANNOTBEEMPTY = "InvalidParameterValue.UserNameCanNotBeEmpty"
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    return c.DescribeUserInfoWithContext(context.Background(), request)
}

// DescribeUserInfo
// 通过用户名或用户 id 搜索用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_USERIDCANNOTBEEMPTY = "InvalidParameterValue.UserIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERNAMECANNOTBEEMPTY = "InvalidParameterValue.UserNameCanNotBeEmpty"
func (c *Client) DescribeUserInfoWithContext(ctx context.Context, request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserResourcesAuthorizationRequest() (request *DescribeUserResourcesAuthorizationRequest) {
    request = &DescribeUserResourcesAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserResourcesAuthorization")
    
    
    return
}

func NewDescribeUserResourcesAuthorizationResponse() (response *DescribeUserResourcesAuthorizationResponse) {
    response = &DescribeUserResourcesAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserResourcesAuthorization
// 查询指定用户下的资源授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserResourcesAuthorization(request *DescribeUserResourcesAuthorizationRequest) (response *DescribeUserResourcesAuthorizationResponse, err error) {
    return c.DescribeUserResourcesAuthorizationWithContext(context.Background(), request)
}

// DescribeUserResourcesAuthorization
// 查询指定用户下的资源授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserResourcesAuthorizationWithContext(ctx context.Context, request *DescribeUserResourcesAuthorizationRequest) (response *DescribeUserResourcesAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeUserResourcesAuthorizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserResourcesAuthorization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResourcesAuthorizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserThirdPartyAccountInfoRequest() (request *DescribeUserThirdPartyAccountInfoRequest) {
    request = &DescribeUserThirdPartyAccountInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "DescribeUserThirdPartyAccountInfo")
    
    
    return
}

func NewDescribeUserThirdPartyAccountInfoResponse() (response *DescribeUserThirdPartyAccountInfoResponse) {
    response = &DescribeUserThirdPartyAccountInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserThirdPartyAccountInfo
// 通过用户名或用户 id 获取用户的第三方账号绑定信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserThirdPartyAccountInfo(request *DescribeUserThirdPartyAccountInfoRequest) (response *DescribeUserThirdPartyAccountInfoResponse, err error) {
    return c.DescribeUserThirdPartyAccountInfoWithContext(context.Background(), request)
}

// DescribeUserThirdPartyAccountInfo
// 通过用户名或用户 id 获取用户的第三方账号绑定信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) DescribeUserThirdPartyAccountInfoWithContext(ctx context.Context, request *DescribeUserThirdPartyAccountInfoRequest) (response *DescribeUserThirdPartyAccountInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserThirdPartyAccountInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserThirdPartyAccountInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserThirdPartyAccountInfoResponse()
    err = c.Send(request, response)
    return
}

func NewListAccountInAccountGroupRequest() (request *ListAccountInAccountGroupRequest) {
    request = &ListAccountInAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAccountInAccountGroup")
    
    
    return
}

func NewListAccountInAccountGroupResponse() (response *ListAccountInAccountGroupResponse) {
    response = &ListAccountInAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAccountInAccountGroup
//  获取账号组中的账号列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAccountInAccountGroup(request *ListAccountInAccountGroupRequest) (response *ListAccountInAccountGroupResponse, err error) {
    return c.ListAccountInAccountGroupWithContext(context.Background(), request)
}

// ListAccountInAccountGroup
//  获取账号组中的账号列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAccountInAccountGroupWithContext(ctx context.Context, request *ListAccountInAccountGroupRequest) (response *ListAccountInAccountGroupResponse, err error) {
    if request == nil {
        request = NewListAccountInAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAccountInAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAccountInAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewListApplicationAuthorizationsRequest() (request *ListApplicationAuthorizationsRequest) {
    request = &ListApplicationAuthorizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListApplicationAuthorizations")
    
    
    return
}

func NewListApplicationAuthorizationsResponse() (response *ListApplicationAuthorizationsResponse) {
    response = &ListApplicationAuthorizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListApplicationAuthorizations
// 应用授权关系列表（含搜索条件匹配）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_ENTITYTYPENOTMATCH = "InvalidParameterValue.EntityTypeNotMatch"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  OPERATIONDENIED_ACTIONPERMISSIONDENY = "OperationDenied.ActionPermissionDeny"
func (c *Client) ListApplicationAuthorizations(request *ListApplicationAuthorizationsRequest) (response *ListApplicationAuthorizationsResponse, err error) {
    return c.ListApplicationAuthorizationsWithContext(context.Background(), request)
}

// ListApplicationAuthorizations
// 应用授权关系列表（含搜索条件匹配）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_ENTITYTYPENOTMATCH = "InvalidParameterValue.EntityTypeNotMatch"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  OPERATIONDENIED_ACTIONPERMISSIONDENY = "OperationDenied.ActionPermissionDeny"
func (c *Client) ListApplicationAuthorizationsWithContext(ctx context.Context, request *ListApplicationAuthorizationsRequest) (response *ListApplicationAuthorizationsResponse, err error) {
    if request == nil {
        request = NewListApplicationAuthorizationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListApplicationAuthorizations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListApplicationAuthorizationsResponse()
    err = c.Send(request, response)
    return
}

func NewListApplicationsRequest() (request *ListApplicationsRequest) {
    request = &ListApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListApplications")
    
    
    return
}

func NewListApplicationsResponse() (response *ListApplicationsResponse) {
    response = &ListApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListApplications
// 获取应用列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_APPLICATIONINFOSORTKEYILLEGAL = "InvalidParameterValue.ApplicationInfoSortKeyIllegal"
//  OPERATIONDENIED_ACTIONPERMISSIONDENY = "OperationDenied.ActionPermissionDeny"
func (c *Client) ListApplications(request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
    return c.ListApplicationsWithContext(context.Background(), request)
}

// ListApplications
// 获取应用列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_APPLICATIONINFOSORTKEYILLEGAL = "InvalidParameterValue.ApplicationInfoSortKeyIllegal"
//  OPERATIONDENIED_ACTIONPERMISSIONDENY = "OperationDenied.ActionPermissionDeny"
func (c *Client) ListApplicationsWithContext(ctx context.Context, request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
    if request == nil {
        request = NewListApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewListApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewListAuthorizedApplicationsToOrgNodeRequest() (request *ListAuthorizedApplicationsToOrgNodeRequest) {
    request = &ListAuthorizedApplicationsToOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAuthorizedApplicationsToOrgNode")
    
    
    return
}

func NewListAuthorizedApplicationsToOrgNodeResponse() (response *ListAuthorizedApplicationsToOrgNodeResponse) {
    response = &ListAuthorizedApplicationsToOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAuthorizedApplicationsToOrgNode
// 通过机构节点ID获得被授权访问的应用列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToOrgNode(request *ListAuthorizedApplicationsToOrgNodeRequest) (response *ListAuthorizedApplicationsToOrgNodeResponse, err error) {
    return c.ListAuthorizedApplicationsToOrgNodeWithContext(context.Background(), request)
}

// ListAuthorizedApplicationsToOrgNode
// 通过机构节点ID获得被授权访问的应用列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToOrgNodeWithContext(ctx context.Context, request *ListAuthorizedApplicationsToOrgNodeRequest) (response *ListAuthorizedApplicationsToOrgNodeResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAuthorizedApplicationsToOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuthorizedApplicationsToOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewListAuthorizedApplicationsToUserRequest() (request *ListAuthorizedApplicationsToUserRequest) {
    request = &ListAuthorizedApplicationsToUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAuthorizedApplicationsToUser")
    
    
    return
}

func NewListAuthorizedApplicationsToUserResponse() (response *ListAuthorizedApplicationsToUserResponse) {
    response = &ListAuthorizedApplicationsToUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAuthorizedApplicationsToUser
// 通过用户ID获得被授权访问的应用列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDTOCODEERROR = "FailedOperation.IdToCodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERAUTHLISTERROR = "FailedOperation.UserAuthListError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUser(request *ListAuthorizedApplicationsToUserRequest) (response *ListAuthorizedApplicationsToUserResponse, err error) {
    return c.ListAuthorizedApplicationsToUserWithContext(context.Background(), request)
}

// ListAuthorizedApplicationsToUser
// 通过用户ID获得被授权访问的应用列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDTOCODEERROR = "FailedOperation.IdToCodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERAUTHLISTERROR = "FailedOperation.UserAuthListError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUserWithContext(ctx context.Context, request *ListAuthorizedApplicationsToUserRequest) (response *ListAuthorizedApplicationsToUserResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAuthorizedApplicationsToUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuthorizedApplicationsToUserResponse()
    err = c.Send(request, response)
    return
}

func NewListAuthorizedApplicationsToUserGroupRequest() (request *ListAuthorizedApplicationsToUserGroupRequest) {
    request = &ListAuthorizedApplicationsToUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListAuthorizedApplicationsToUserGroup")
    
    
    return
}

func NewListAuthorizedApplicationsToUserGroupResponse() (response *ListAuthorizedApplicationsToUserGroupResponse) {
    response = &ListAuthorizedApplicationsToUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAuthorizedApplicationsToUserGroup
// 通过用户组ID获得被授权访问的应用列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERGROUPNOTEXIST = "FailedOperation.UserGroupNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUserGroup(request *ListAuthorizedApplicationsToUserGroupRequest) (response *ListAuthorizedApplicationsToUserGroupResponse, err error) {
    return c.ListAuthorizedApplicationsToUserGroupWithContext(context.Background(), request)
}

// ListAuthorizedApplicationsToUserGroup
// 通过用户组ID获得被授权访问的应用列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERGROUPNOTEXIST = "FailedOperation.UserGroupNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListAuthorizedApplicationsToUserGroupWithContext(ctx context.Context, request *ListAuthorizedApplicationsToUserGroupRequest) (response *ListAuthorizedApplicationsToUserGroupResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAuthorizedApplicationsToUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuthorizedApplicationsToUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewListUserGroupsRequest() (request *ListUserGroupsRequest) {
    request = &ListUserGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUserGroups")
    
    
    return
}

func NewListUserGroupsResponse() (response *ListUserGroupsResponse) {
    response = &ListUserGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUserGroups
// 获取用户组列表信息（包含查询条件）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUserGroups(request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    return c.ListUserGroupsWithContext(context.Background(), request)
}

// ListUserGroups
// 获取用户组列表信息（包含查询条件）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUserGroupsWithContext(ctx context.Context, request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    if request == nil {
        request = NewListUserGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListUserGroupsOfUserRequest() (request *ListUserGroupsOfUserRequest) {
    request = &ListUserGroupsOfUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUserGroupsOfUser")
    
    
    return
}

func NewListUserGroupsOfUserResponse() (response *ListUserGroupsOfUserResponse) {
    response = &ListUserGroupsOfUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUserGroupsOfUser
// 获取用户所在的用户组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LISTUSERGROUPSOFUSERERROR = "FailedOperation.ListUserGroupsOfUserError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_SORTKEYILLEGAL = "InvalidParameterValue.SortKeyIllegal"
func (c *Client) ListUserGroupsOfUser(request *ListUserGroupsOfUserRequest) (response *ListUserGroupsOfUserResponse, err error) {
    return c.ListUserGroupsOfUserWithContext(context.Background(), request)
}

// ListUserGroupsOfUser
// 获取用户所在的用户组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_LISTUSERGROUPSOFUSERERROR = "FailedOperation.ListUserGroupsOfUserError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_SORTKEYILLEGAL = "InvalidParameterValue.SortKeyIllegal"
func (c *Client) ListUserGroupsOfUserWithContext(ctx context.Context, request *ListUserGroupsOfUserRequest) (response *ListUserGroupsOfUserResponse, err error) {
    if request == nil {
        request = NewListUserGroupsOfUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserGroupsOfUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserGroupsOfUserResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersRequest() (request *ListUsersRequest) {
    request = &ListUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUsers")
    
    
    return
}

func NewListUsersResponse() (response *ListUsersResponse) {
    response = &ListUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUsers
// 获取用户列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPECTFIELDSNOTFOUND = "FailedOperation.ExpectFieldsNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_SORTKEYILLEGAL = "InvalidParameterValue.SortKeyIllegal"
func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    return c.ListUsersWithContext(context.Background(), request)
}

// ListUsers
// 获取用户列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPECTFIELDSNOTFOUND = "FailedOperation.ExpectFieldsNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_SORTKEYILLEGAL = "InvalidParameterValue.SortKeyIllegal"
func (c *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest) (response *ListUsersResponse, err error) {
    if request == nil {
        request = NewListUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersInOrgNodeRequest() (request *ListUsersInOrgNodeRequest) {
    request = &ListUsersInOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUsersInOrgNode")
    
    
    return
}

func NewListUsersInOrgNodeResponse() (response *ListUsersInOrgNodeResponse) {
    response = &ListUsersInOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUsersInOrgNode
// 根据机构节点ID读取节点下用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LISTUSERSINORGNODEERROR = "FailedOperation.ListUsersInOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInOrgNode(request *ListUsersInOrgNodeRequest) (response *ListUsersInOrgNodeResponse, err error) {
    return c.ListUsersInOrgNodeWithContext(context.Background(), request)
}

// ListUsersInOrgNode
// 根据机构节点ID读取节点下用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_LISTUSERSINORGNODEERROR = "FailedOperation.ListUsersInOrgNodeError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInOrgNodeWithContext(ctx context.Context, request *ListUsersInOrgNodeRequest) (response *ListUsersInOrgNodeResponse, err error) {
    if request == nil {
        request = NewListUsersInOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsersInOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersInOrgNodeResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersInUserGroupRequest() (request *ListUsersInUserGroupRequest) {
    request = &ListUsersInUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ListUsersInUserGroup")
    
    
    return
}

func NewListUsersInUserGroupResponse() (response *ListUsersInUserGroupResponse) {
    response = &ListUsersInUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListUsersInUserGroup
// 获取用户组中的用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_LISTUSERSINUSERGROUPERROR = "FailedOperation.ListUsersInUserGroupError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInUserGroup(request *ListUsersInUserGroupRequest) (response *ListUsersInUserGroupResponse, err error) {
    return c.ListUsersInUserGroupWithContext(context.Background(), request)
}

// ListUsersInUserGroup
// 获取用户组中的用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_LISTUSERSINUSERGROUPERROR = "FailedOperation.ListUsersInUserGroupError"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ListUsersInUserGroupWithContext(ctx context.Context, request *ListUsersInUserGroupRequest) (response *ListUsersInUserGroupResponse, err error) {
    if request == nil {
        request = NewListUsersInUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsersInUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersInUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountGroupRequest() (request *ModifyAccountGroupRequest) {
    request = &ModifyAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyAccountGroup")
    
    
    return
}

func NewModifyAccountGroupResponse() (response *ModifyAccountGroupResponse) {
    response = &ModifyAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountGroup
// 修改账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEALREADYEXISTS = "FailedOperation.AccountGroupNameAlreadyExists"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAccountGroup(request *ModifyAccountGroupRequest) (response *ModifyAccountGroupResponse, err error) {
    return c.ModifyAccountGroupWithContext(context.Background(), request)
}

// ModifyAccountGroup
// 修改账号组
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNAMEALREADYEXISTS = "FailedOperation.AccountGroupNameAlreadyExists"
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAccountGroupWithContext(ctx context.Context, request *ModifyAccountGroupRequest) (response *ModifyAccountGroupResponse, err error) {
    if request == nil {
        request = NewModifyAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppAccountRequest() (request *ModifyAppAccountRequest) {
    request = &ModifyAppAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyAppAccount")
    
    
    return
}

func NewModifyAppAccountResponse() (response *ModifyAppAccountResponse) {
    response = &ModifyAppAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAppAccount
// 修改应用账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNAMEALREADYEXISTS = "FailedOperation.AccountNameAlreadyExists"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAppAccount(request *ModifyAppAccountRequest) (response *ModifyAppAccountResponse, err error) {
    return c.ModifyAppAccountWithContext(context.Background(), request)
}

// ModifyAppAccount
// 修改应用账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNAMEALREADYEXISTS = "FailedOperation.AccountNameAlreadyExists"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyAppAccountWithContext(ctx context.Context, request *ModifyAppAccountRequest) (response *ModifyAppAccountResponse, err error) {
    if request == nil {
        request = NewModifyAppAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAppAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplication
// 更新一个应用的信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPDISPLAYNAMEALREADYEXISTS = "FailedOperation.AppDisplayNameAlreadyExists"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// 更新一个应用的信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPDISPLAYNAMEALREADYEXISTS = "FailedOperation.AppDisplayNameAlreadyExists"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserInfoRequest() (request *ModifyUserInfoRequest) {
    request = &ModifyUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "ModifyUserInfo")
    
    
    return
}

func NewModifyUserInfoResponse() (response *ModifyUserInfoResponse) {
    response = &ModifyUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUserInfo
// 通过用户名或用户 id 冻结用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_UPDATELDAPUSERORGEXCEEDSRANGE = "FailedOperation.UpdateLDAPUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATELDAPUSERORGNOTINSAMECROP = "FailedOperation.UpdateLDAPUserOrgNotInSameCrop"
//  FAILEDOPERATION_UPDATEUSEREXCEEDSRANGE = "FailedOperation.UpdateUserExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGEXCEEDSRANGE = "FailedOperation.UpdateWeComUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGNOTINSAMECROP = "FailedOperation.UpdateWeComUserOrgNotInSameCrop"
//  FAILEDOPERATION_USEREMAILALREADYEXISTS = "FailedOperation.UserEmailAlreadyExists"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETERVALUE_NEWPASSWORDCANNOTBEEMPTY = "InvalidParameterValue.NewPasswordCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_USERIDCANNOTBEEMPTY = "InvalidParameterValue.UserIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERNAMECANNOTBEEMPTY = "InvalidParameterValue.UserNameCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERPHONECANNOTBEEMPTY = "InvalidParameterValue.UserPhoneCanNotBeEmpty"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
func (c *Client) ModifyUserInfo(request *ModifyUserInfoRequest) (response *ModifyUserInfoResponse, err error) {
    return c.ModifyUserInfoWithContext(context.Background(), request)
}

// ModifyUserInfo
// 通过用户名或用户 id 冻结用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"
//  FAILEDOPERATION_UPDATELDAPUSERORGEXCEEDSRANGE = "FailedOperation.UpdateLDAPUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATELDAPUSERORGNOTINSAMECROP = "FailedOperation.UpdateLDAPUserOrgNotInSameCrop"
//  FAILEDOPERATION_UPDATEUSEREXCEEDSRANGE = "FailedOperation.UpdateUserExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGEXCEEDSRANGE = "FailedOperation.UpdateWeComUserOrgExceedsRange"
//  FAILEDOPERATION_UPDATEWECOMUSERORGNOTINSAMECROP = "FailedOperation.UpdateWeComUserOrgNotInSameCrop"
//  FAILEDOPERATION_USEREMAILALREADYEXISTS = "FailedOperation.UserEmailAlreadyExists"
//  FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"
//  INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"
//  INVALIDPARAMETERVALUE_NEWPASSWORDCANNOTBEEMPTY = "InvalidParameterValue.NewPasswordCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  INVALIDPARAMETERVALUE_USERIDCANNOTBEEMPTY = "InvalidParameterValue.UserIdCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERNAMECANNOTBEEMPTY = "InvalidParameterValue.UserNameCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_USERPHONECANNOTBEEMPTY = "InvalidParameterValue.UserPhoneCanNotBeEmpty"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
//  LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"
func (c *Client) ModifyUserInfoWithContext(ctx context.Context, request *ModifyUserInfoRequest) (response *ModifyUserInfoResponse, err error) {
    if request == nil {
        request = NewModifyUserInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveAccountFromAccountGroupRequest() (request *RemoveAccountFromAccountGroupRequest) {
    request = &RemoveAccountFromAccountGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "RemoveAccountFromAccountGroup")
    
    
    return
}

func NewRemoveAccountFromAccountGroupResponse() (response *RemoveAccountFromAccountGroupResponse) {
    response = &RemoveAccountFromAccountGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveAccountFromAccountGroup
// 从账号组中移除账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_ACCOUNTIDSCANNOTBEEMPTY = "InvalidParameterValue.AccountIdsCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) RemoveAccountFromAccountGroup(request *RemoveAccountFromAccountGroupRequest) (response *RemoveAccountFromAccountGroupResponse, err error) {
    return c.RemoveAccountFromAccountGroupWithContext(context.Background(), request)
}

// RemoveAccountFromAccountGroup
// 从账号组中移除账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"
//  FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"
//  FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETERVALUE_ACCOUNTIDSCANNOTBEEMPTY = "InvalidParameterValue.AccountIdsCanNotBeEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"
func (c *Client) RemoveAccountFromAccountGroupWithContext(ctx context.Context, request *RemoveAccountFromAccountGroupRequest) (response *RemoveAccountFromAccountGroupResponse, err error) {
    if request == nil {
        request = NewRemoveAccountFromAccountGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveAccountFromAccountGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveAccountFromAccountGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserFromUserGroupRequest() (request *RemoveUserFromUserGroupRequest) {
    request = &RemoveUserFromUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "RemoveUserFromUserGroup")
    
    
    return
}

func NewRemoveUserFromUserGroupResponse() (response *RemoveUserFromUserGroupResponse) {
    response = &RemoveUserFromUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveUserFromUserGroup
// 从用户组中移除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPERROR = "FailedOperation.RemoveUsersFromUserGroupError"
//  FAILEDOPERATION_USERNOTEXISTINUSERGROUP = "FailedOperation.UserNotExistInUserGroup"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) RemoveUserFromUserGroup(request *RemoveUserFromUserGroupRequest) (response *RemoveUserFromUserGroupResponse, err error) {
    return c.RemoveUserFromUserGroupWithContext(context.Background(), request)
}

// RemoveUserFromUserGroup
// 从用户组中移除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPERROR = "FailedOperation.RemoveUsersFromUserGroupError"
//  FAILEDOPERATION_USERNOTEXISTINUSERGROUP = "FailedOperation.UserNotExistInUserGroup"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
func (c *Client) RemoveUserFromUserGroupWithContext(ctx context.Context, request *RemoveUserFromUserGroupRequest) (response *RemoveUserFromUserGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromUserGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserFromUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserFromUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrgNodeRequest() (request *UpdateOrgNodeRequest) {
    request = &UpdateOrgNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("eiam", APIVersion, "UpdateOrgNode")
    
    
    return
}

func NewUpdateOrgNodeResponse() (response *UpdateOrgNodeResponse) {
    response = &UpdateOrgNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateOrgNode
// 新建一个机构节点，
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) UpdateOrgNode(request *UpdateOrgNodeRequest) (response *UpdateOrgNodeResponse, err error) {
    return c.UpdateOrgNodeWithContext(context.Background(), request)
}

// UpdateOrgNode
// 新建一个机构节点，
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"
//  FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"
//  FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"
//  INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"
//  INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"
//  LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"
func (c *Client) UpdateOrgNodeWithContext(ctx context.Context, request *UpdateOrgNodeRequest) (response *UpdateOrgNodeResponse, err error) {
    if request == nil {
        request = NewUpdateOrgNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrgNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrgNodeResponse()
    err = c.Send(request, response)
    return
}
