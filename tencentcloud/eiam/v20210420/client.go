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
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) AddAccountToAccountGroup(request *AddAccountToAccountGroupRequest) (response *AddAccountToAccountGroupResponse, err error) {
    if request == nil {
        request = NewAddAccountToAccountGroupRequest()
    }
    
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
//  FAILEDOPERATION_ADDUSERSTOUSERGROUP = "FailedOperation.AddUsersToUserGroup"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERALREADYEXISTEDINUSERGROUP = "FailedOperation.UserAlreadyExistedInUserGroup"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) AddUserToUserGroup(request *AddUserToUserGroupRequest) (response *AddUserToUserGroupResponse, err error) {
    if request == nil {
        request = NewAddUserToUserGroupRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTGROUPNAMEEXISTED = "FailedOperation.AccountGroupNameExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) CreateAccountGroup(request *CreateAccountGroupRequest) (response *CreateAccountGroupResponse, err error) {
    if request == nil {
        request = NewCreateAccountGroupRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNAMEEXISTED = "FailedOperation.AccountNameExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) CreateAppAccount(request *CreateAppAccountRequest) (response *CreateAppAccountResponse, err error) {
    if request == nil {
        request = NewCreateAppAccountRequest()
    }
    
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
//  FAILEDOPERATION_CHILDORGNODENAMEEXISTS = "FailedOperation.ChildOrgNodeNameExists"
//  FAILEDOPERATION_CREATEORGNODEFAILURE = "FailedOperation.CreateOrgNodeFailure"
//  FAILEDOPERATION_CUSTOMIZEDPARENTORGNODEIDEXISTED = "FailedOperation.CustomizedParentOrgNodeIdExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) CreateOrgNode(request *CreateOrgNodeRequest) (response *CreateOrgNodeResponse, err error) {
    if request == nil {
        request = NewCreateOrgNodeRequest()
    }
    
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
//  FAILEDOPERATION_DESCRIBEORGNODEROOTFAILURE = "FailedOperation.DescribeOrgNodeRootFailure"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  FAILEDOPERATION_USEREXPRIATIONTIMEISILLEGAL = "FailedOperation.UserExpriationTimeIsIllegal"
//  FAILEDOPERATION_USERNAMEEXISTED = "FailedOperation.UserNameExisted"
//  FAILEDOPERATION_USERPHONEEXISTED = "FailedOperation.UserPhoneExisted"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDFAILURE = "InvalidParameter.AttributeValueValidFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETER_PASSWORDISILLEGAL = "InvalidParameter.PasswordIsIllegal"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
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
//  FAILEDOPERATION_CREATEUSERGROUPFAILURE = "FailedOperation.CreateUserGroupFailure"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
    if request == nil {
        request = NewCreateUserGroupRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DeleteAccountGroup(request *DeleteAccountGroupRequest) (response *DeleteAccountGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAccountGroupRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DeleteAppAccount(request *DeleteAppAccountRequest) (response *DeleteAppAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAppAccountRequest()
    }
    
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
//  FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCannotBeDeleted"
//  FAILEDOPERATION_DELETEORGNODEFAILURE = "FailedOperation.DeleteOrgNodeFailure"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCannotBeDeleted"
//  FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCannotBeDeleted"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) DeleteOrgNode(request *DeleteOrgNodeRequest) (response *DeleteOrgNodeResponse, err error) {
    if request == nil {
        request = NewDeleteOrgNodeRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
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
//  FAILEDOPERATION_DELETEUSERGROUPFAILURE = "FailedOperation.DeleteUserGroupFailure"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (response *DeleteUserGroupResponse, err error) {
    if request == nil {
        request = NewDeleteUserGroupRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    
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
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeAccountGroup(request *DescribeAccountGroupRequest) (response *DescribeAccountGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAccountGroupRequest()
    }
    
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
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeAppAccount(request *DescribeAppAccountRequest) (response *DescribeAppAccountResponse, err error) {
    if request == nil {
        request = NewDescribeAppAccountRequest()
    }
    
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
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationRequest()
    }
    
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
//  FAILEDOPERATION_DESCRIBEORGNODEFAILURE = "FailedOperation.DescribeOrgNodeFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTFAILURE = "FailedOperation.DescribeOrgNodeRootFailure"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeOrgNode(request *DescribeOrgNodeRequest) (response *DescribeOrgNodeResponse, err error) {
    if request == nil {
        request = NewDescribeOrgNodeRequest()
    }
    
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
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeOrgResourcesAuthorization(request *DescribeOrgResourcesAuthorizationRequest) (response *DescribeOrgResourcesAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrgResourcesAuthorizationRequest()
    }
    
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
//  FAILEDOPERATION_APPIDISNULL = "FailedOperation.AppIdIsNull"
//  FAILEDOPERATION_APPIDNOTFOUND = "FailedOperation.AppIdNotFound"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) DescribePublicKey(request *DescribePublicKeyRequest) (response *DescribePublicKeyResponse, err error) {
    if request == nil {
        request = NewDescribePublicKeyRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) DescribeUserGroup(request *DescribeUserGroupRequest) (response *DescribeUserGroupResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupRequest()
    }
    
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
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeUserGroupResourcesAuthorization(request *DescribeUserGroupResourcesAuthorizationRequest) (response *DescribeUserGroupResourcesAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeUserGroupResourcesAuthorizationRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    
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
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) DescribeUserResourcesAuthorization(request *DescribeUserResourcesAuthorizationRequest) (response *DescribeUserResourcesAuthorizationResponse, err error) {
    if request == nil {
        request = NewDescribeUserResourcesAuthorizationRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
func (c *Client) DescribeUserThirdPartyAccountInfo(request *DescribeUserThirdPartyAccountInfoRequest) (response *DescribeUserThirdPartyAccountInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserThirdPartyAccountInfoRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ListAccountInAccountGroup(request *ListAccountInAccountGroupRequest) (response *ListAccountInAccountGroupResponse, err error) {
    if request == nil {
        request = NewListAccountInAccountGroupRequest()
    }
    
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
//  FAILEDOPERATION_ENTITYTYPENOTEXISTED = "FailedOperation.EntityTypeNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_SEARCHCRITERIAISILLEGAL = "FailedOperation.SearchCriteriaIsIllegal"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ListApplicationAuthorizations(request *ListApplicationAuthorizationsRequest) (response *ListApplicationAuthorizationsResponse, err error) {
    if request == nil {
        request = NewListApplicationAuthorizationsRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_SEARCHCRITERIAISILLEGAL = "FailedOperation.SearchCriteriaIsIllegal"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ListApplications(request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
    if request == nil {
        request = NewListApplicationsRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) ListAuthorizedApplicationsToOrgNode(request *ListAuthorizedApplicationsToOrgNodeRequest) (response *ListAuthorizedApplicationsToOrgNodeResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToOrgNodeRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_USERAUTHLISTFAILED = "FailedOperation.UserAuthListFailed"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) ListAuthorizedApplicationsToUser(request *ListAuthorizedApplicationsToUserRequest) (response *ListAuthorizedApplicationsToUserResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToUserRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_USERGROUPNOTEXIST = "FailedOperation.UserGroupNotExist"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) ListAuthorizedApplicationsToUserGroup(request *ListAuthorizedApplicationsToUserGroupRequest) (response *ListAuthorizedApplicationsToUserGroupResponse, err error) {
    if request == nil {
        request = NewListAuthorizedApplicationsToUserGroupRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_SEARCHCRITERIAISILLEGAL = "FailedOperation.SearchCriteriaIsIllegal"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ListUserGroups(request *ListUserGroupsRequest) (response *ListUserGroupsResponse, err error) {
    if request == nil {
        request = NewListUserGroupsRequest()
    }
    
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
//  FAILEDOPERATION_LISTUSERGROUPSOFUSERFAILURE = "FailedOperation.ListUserGroupsOfUserFailure"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_USERINFOSORTKEYISILLEGAL = "FailedOperation.UserInfoSortKeyIsIllegal"
//  FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) ListUserGroupsOfUser(request *ListUserGroupsOfUserRequest) (response *ListUserGroupsOfUserResponse, err error) {
    if request == nil {
        request = NewListUserGroupsOfUserRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_SEARCHCRITERIAISILLEGAL = "FailedOperation.SearchCriteriaIsIllegal"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  FAILEDOPERATION_USERINFOSORTKEYISILLEGAL = "FailedOperation.UserInfoSortKeyIsIllegal"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    if request == nil {
        request = NewListUsersRequest()
    }
    
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
//  FAILEDOPERATION_DESCRIBEORGNODEROOTFAILURE = "FailedOperation.DescribeOrgNodeRootFailure"
//  FAILEDOPERATION_LISTUSERSINORGNODEFAILURE = "FailedOperation.ListUsersInOrgNodeFailure"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  FAILEDOPERATION_SEARCHCRITERIAISILLEGAL = "FailedOperation.SearchCriteriaIsIllegal"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ListUsersInOrgNode(request *ListUsersInOrgNodeRequest) (response *ListUsersInOrgNodeResponse, err error) {
    if request == nil {
        request = NewListUsersInOrgNodeRequest()
    }
    
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
//  FAILEDOPERATION_LISTUSERSINUSERGROUPFAILURE = "FailedOperation.ListUsersInUserGroupFailure"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) ListUsersInUserGroup(request *ListUsersInUserGroupRequest) (response *ListUsersInUserGroupResponse, err error) {
    if request == nil {
        request = NewListUsersInUserGroupRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ModifyAccountGroup(request *ModifyAccountGroupRequest) (response *ModifyAccountGroupResponse, err error) {
    if request == nil {
        request = NewModifyAccountGroupRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ModifyAppAccount(request *ModifyAppAccountRequest) (response *ModifyAppAccountResponse, err error) {
    if request == nil {
        request = NewModifyAppAccountRequest()
    }
    
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
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    
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
//  FAILEDOPERATION_DESCRIBEORGNODEFAILURE = "FailedOperation.DescribeOrgNodeFailure"
//  FAILEDOPERATION_DESCRIBEORGNODEROOTFAILURE = "FailedOperation.DescribeOrgNodeRootFailure"
//  FAILEDOPERATION_NEWPASSWORDMUSTNOTBLANK = "FailedOperation.NewPasswordMustNotBlank"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_TIMEFORMATISILLEGAL = "FailedOperation.TimeFormatIsIllegal"
//  FAILEDOPERATION_USEREMAILEXISTED = "FailedOperation.UserEmailExisted"
//  FAILEDOPERATION_USERPHONEEXISTED = "FailedOperation.UserPhoneExisted"
//  FAILEDOPERATION_USERPHONEISEMPTY = "FailedOperation.UserPhoneIsEmpty"
//  INVALIDPARAMETER_ATTRIBUTEVALUEVALIDFAILURE = "InvalidParameter.AttributeValueValidFailure"
//  INVALIDPARAMETER_PARAMETEREXCEEDSLENGTHLIMIT = "InvalidParameter.ParameterExceedsLengthLimit"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
//  INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"
//  INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"
func (c *Client) ModifyUserInfo(request *ModifyUserInfoRequest) (response *ModifyUserInfoResponse, err error) {
    if request == nil {
        request = NewModifyUserInfoRequest()
    }
    
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
//  FAILEDOPERATION_ACCOUNTGROUPNOTEXISTED = "FailedOperation.AccountGroupNotExisted"
//  FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"
//  FAILEDOPERATION_APPNOTEXISTED = "FailedOperation.AppNotExisted"
//  FAILEDOPERATION_ITEMSEXCEEDMAXNUMBER = "FailedOperation.ItemsExceedMaxNumber"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"
func (c *Client) RemoveAccountFromAccountGroup(request *RemoveAccountFromAccountGroupRequest) (response *RemoveAccountFromAccountGroupResponse, err error) {
    if request == nil {
        request = NewRemoveAccountFromAccountGroupRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"
//  FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPFAILURE = "FailedOperation.RemoveUsersFromUserGroupFailure"
//  FAILEDOPERATION_USERGROUPNOTEXISTED = "FailedOperation.UserGroupNotExisted"
//  FAILEDOPERATION_USERNOTEXISTEDINUSERGROUP = "FailedOperation.UserNotExistedInUserGroup"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) RemoveUserFromUserGroup(request *RemoveUserFromUserGroupRequest) (response *RemoveUserFromUserGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromUserGroupRequest()
    }
    
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
//  FAILEDOPERATION_CHILDORGNODENAMEEXISTS = "FailedOperation.ChildOrgNodeNameExists"
//  FAILEDOPERATION_CUSTOMIZEDPARENTORGNODEIDEXISTED = "FailedOperation.CustomizedParentOrgNodeIdExisted"
//  FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"
//  FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"
//  FAILEDOPERATION_PARENTORGNODEISEMPTY = "FailedOperation.ParentOrgNodeIsEmpty"
//  INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"
func (c *Client) UpdateOrgNode(request *UpdateOrgNodeRequest) (response *UpdateOrgNodeResponse, err error) {
    if request == nil {
        request = NewUpdateOrgNodeRequest()
    }
    
    response = NewUpdateOrgNodeResponse()
    err = c.Send(request, response)
    return
}
