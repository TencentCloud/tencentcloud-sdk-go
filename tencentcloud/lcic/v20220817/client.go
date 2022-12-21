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

package v20220817

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-08-17"

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


func NewBindDocumentToRoomRequest() (request *BindDocumentToRoomRequest) {
    request = &BindDocumentToRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BindDocumentToRoom")
    
    
    return
}

func NewBindDocumentToRoomResponse() (response *BindDocumentToRoomResponse) {
    response = &BindDocumentToRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindDocumentToRoom
// 绑定文档到房间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoom(request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    return c.BindDocumentToRoomWithContext(context.Background(), request)
}

// BindDocumentToRoom
// 绑定文档到房间
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoomWithContext(ctx context.Context, request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    if request == nil {
        request = NewBindDocumentToRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDocumentToRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDocumentToRoomResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDocumentRequest() (request *CreateDocumentRequest) {
    request = &CreateDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateDocument")
    
    
    return
}

func NewCreateDocumentResponse() (response *CreateDocumentResponse) {
    response = &CreateDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDocument
// 创建房间内可以使用的文档。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateDocument(request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    return c.CreateDocumentWithContext(context.Background(), request)
}

// CreateDocument
// 创建房间内可以使用的文档。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateDocumentWithContext(ctx context.Context, request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    if request == nil {
        request = NewCreateDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoomRequest() (request *CreateRoomRequest) {
    request = &CreateRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateRoom")
    
    
    return
}

func NewCreateRoomResponse() (response *CreateRoomResponse) {
    response = &CreateRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRoom
// 创建房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateRoom(request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    return c.CreateRoomWithContext(context.Background(), request)
}

// CreateRoom
// 创建房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateRoomWithContext(ctx context.Context, request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    if request == nil {
        request = NewCreateRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoomResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSupervisorRequest() (request *CreateSupervisorRequest) {
    request = &CreateSupervisorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateSupervisor")
    
    
    return
}

func NewCreateSupervisorResponse() (response *CreateSupervisorResponse) {
    response = &CreateSupervisorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSupervisor
// 创建巡课
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateSupervisor(request *CreateSupervisorRequest) (response *CreateSupervisorResponse, err error) {
    return c.CreateSupervisorWithContext(context.Background(), request)
}

// CreateSupervisor
// 创建巡课
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateSupervisorWithContext(ctx context.Context, request *CreateSupervisorRequest) (response *CreateSupervisorResponse, err error) {
    if request == nil {
        request = NewCreateSupervisorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSupervisor require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSupervisorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoomRequest() (request *DeleteRoomRequest) {
    request = &DeleteRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteRoom")
    
    
    return
}

func NewDeleteRoomResponse() (response *DeleteRoomResponse) {
    response = &DeleteRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoom
// 删除房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DeleteRoom(request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
    return c.DeleteRoomWithContext(context.Background(), request)
}

// DeleteRoom
// 删除房间
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DeleteRoomWithContext(ctx context.Context, request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
    if request == nil {
        request = NewDeleteRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomRequest() (request *DescribeRoomRequest) {
    request = &DescribeRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRoom")
    
    
    return
}

func NewDescribeRoomResponse() (response *DescribeRoomResponse) {
    response = &DescribeRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoom
// 获取房间信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoom(request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    return c.DescribeRoomWithContext(context.Background(), request)
}

// DescribeRoom
// 获取房间信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoomWithContext(ctx context.Context, request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    if request == nil {
        request = NewDescribeRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomStatisticsRequest() (request *DescribeRoomStatisticsRequest) {
    request = &DescribeRoomStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRoomStatistics")
    
    
    return
}

func NewDescribeRoomStatisticsResponse() (response *DescribeRoomStatisticsResponse) {
    response = &DescribeRoomStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoomStatistics
// 获取房间统计信息，仅可在房间结束后调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
func (c *Client) DescribeRoomStatistics(request *DescribeRoomStatisticsRequest) (response *DescribeRoomStatisticsResponse, err error) {
    return c.DescribeRoomStatisticsWithContext(context.Background(), request)
}

// DescribeRoomStatistics
// 获取房间统计信息，仅可在房间结束后调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
func (c *Client) DescribeRoomStatisticsWithContext(ctx context.Context, request *DescribeRoomStatisticsRequest) (response *DescribeRoomStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRoomStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUser
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// 获取用户信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewLoginOriginIdRequest() (request *LoginOriginIdRequest) {
    request = &LoginOriginIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "LoginOriginId")
    
    
    return
}

func NewLoginOriginIdResponse() (response *LoginOriginIdResponse) {
    response = &LoginOriginIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LoginOriginId
// 使用源账号登录，源账号为注册时填入的originId
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginOriginId(request *LoginOriginIdRequest) (response *LoginOriginIdResponse, err error) {
    return c.LoginOriginIdWithContext(context.Background(), request)
}

// LoginOriginId
// 使用源账号登录，源账号为注册时填入的originId
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginOriginIdWithContext(ctx context.Context, request *LoginOriginIdRequest) (response *LoginOriginIdResponse, err error) {
    if request == nil {
        request = NewLoginOriginIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginOriginId require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginOriginIdResponse()
    err = c.Send(request, response)
    return
}

func NewLoginUserRequest() (request *LoginUserRequest) {
    request = &LoginUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "LoginUser")
    
    
    return
}

func NewLoginUserResponse() (response *LoginUserResponse) {
    response = &LoginUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LoginUser
// 登录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginUser(request *LoginUserRequest) (response *LoginUserResponse, err error) {
    return c.LoginUserWithContext(context.Background(), request)
}

// LoginUser
// 登录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginUserWithContext(ctx context.Context, request *LoginUserRequest) (response *LoginUserResponse, err error) {
    if request == nil {
        request = NewLoginUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApp
// 修改应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// 修改应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterUserRequest() (request *RegisterUserRequest) {
    request = &RegisterUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "RegisterUser")
    
    
    return
}

func NewRegisterUserResponse() (response *RegisterUserResponse) {
    response = &RegisterUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterUser
// 注册用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) RegisterUser(request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
    return c.RegisterUserWithContext(context.Background(), request)
}

// RegisterUser
// 注册用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) RegisterUserWithContext(ctx context.Context, request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
    if request == nil {
        request = NewRegisterUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterUserResponse()
    err = c.Send(request, response)
    return
}

func NewSetAppCustomContentRequest() (request *SetAppCustomContentRequest) {
    request = &SetAppCustomContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SetAppCustomContent")
    
    
    return
}

func NewSetAppCustomContentResponse() (response *SetAppCustomContentResponse) {
    response = &SetAppCustomContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetAppCustomContent
// 设置应用的自定义内容，包括应用图标，自定义的代码等。如果已存在，则为更新。更新js、css内容后，要生效也需要调用该接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContent(request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    return c.SetAppCustomContentWithContext(context.Background(), request)
}

// SetAppCustomContent
// 设置应用的自定义内容，包括应用图标，自定义的代码等。如果已存在，则为更新。更新js、css内容后，要生效也需要调用该接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContentWithContext(ctx context.Context, request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    if request == nil {
        request = NewSetAppCustomContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAppCustomContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAppCustomContentResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindDocumentFromRoomRequest() (request *UnbindDocumentFromRoomRequest) {
    request = &UnbindDocumentFromRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "UnbindDocumentFromRoom")
    
    
    return
}

func NewUnbindDocumentFromRoomResponse() (response *UnbindDocumentFromRoomResponse) {
    response = &UnbindDocumentFromRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindDocumentFromRoom
// 文档从房间解绑
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) UnbindDocumentFromRoom(request *UnbindDocumentFromRoomRequest) (response *UnbindDocumentFromRoomResponse, err error) {
    return c.UnbindDocumentFromRoomWithContext(context.Background(), request)
}

// UnbindDocumentFromRoom
// 文档从房间解绑
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) UnbindDocumentFromRoomWithContext(ctx context.Context, request *UnbindDocumentFromRoomRequest) (response *UnbindDocumentFromRoomResponse, err error) {
    if request == nil {
        request = NewUnbindDocumentFromRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindDocumentFromRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindDocumentFromRoomResponse()
    err = c.Send(request, response)
    return
}
