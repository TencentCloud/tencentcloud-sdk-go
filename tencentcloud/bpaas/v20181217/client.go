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

package v20181217

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-12-17"

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


func NewGetBpaasApproveDetailRequest() (request *GetBpaasApproveDetailRequest) {
    request = &GetBpaasApproveDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bpaas", APIVersion, "GetBpaasApproveDetail")
    
    
    return
}

func NewGetBpaasApproveDetailResponse() (response *GetBpaasApproveDetailResponse) {
    response = &GetBpaasApproveDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetBpaasApproveDetail
// 查看审批详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER_IDNOTEXIST = "InvalidParameter.IdNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GetBpaasApproveDetail(request *GetBpaasApproveDetailRequest) (response *GetBpaasApproveDetailResponse, err error) {
    return c.GetBpaasApproveDetailWithContext(context.Background(), request)
}

// GetBpaasApproveDetail
// 查看审批详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER_IDNOTEXIST = "InvalidParameter.IdNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GetBpaasApproveDetailWithContext(ctx context.Context, request *GetBpaasApproveDetailRequest) (response *GetBpaasApproveDetailResponse, err error) {
    if request == nil {
        request = NewGetBpaasApproveDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetBpaasApproveDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetBpaasApproveDetailResponse()
    err = c.Send(request, response)
    return
}

func NewOutApproveBpaasApplicationRequest() (request *OutApproveBpaasApplicationRequest) {
    request = &OutApproveBpaasApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bpaas", APIVersion, "OutApproveBpaasApplication")
    
    
    return
}

func NewOutApproveBpaasApplicationResponse() (response *OutApproveBpaasApplicationResponse) {
    response = &OutApproveBpaasApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OutApproveBpaasApplication
// 外部审批申请单
//
// 可能返回的错误码:
//  FAILEDOPERATION_SENDTOCKAFKA = "FailedOperation.SendToCkafka"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER_HASBEENAPPROVED = "InvalidParameter.HasBeenApproved"
//  INVALIDPARAMETER_ILLEGALNODE = "InvalidParameter.IllegalNode"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OutApproveBpaasApplication(request *OutApproveBpaasApplicationRequest) (response *OutApproveBpaasApplicationResponse, err error) {
    return c.OutApproveBpaasApplicationWithContext(context.Background(), request)
}

// OutApproveBpaasApplication
// 外部审批申请单
//
// 可能返回的错误码:
//  FAILEDOPERATION_SENDTOCKAFKA = "FailedOperation.SendToCkafka"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DBERROR = "InternalError.DbError"
//  INVALIDPARAMETER_HASBEENAPPROVED = "InvalidParameter.HasBeenApproved"
//  INVALIDPARAMETER_ILLEGALNODE = "InvalidParameter.IllegalNode"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OutApproveBpaasApplicationWithContext(ctx context.Context, request *OutApproveBpaasApplicationRequest) (response *OutApproveBpaasApplicationResponse, err error) {
    if request == nil {
        request = NewOutApproveBpaasApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OutApproveBpaasApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewOutApproveBpaasApplicationResponse()
    err = c.Send(request, response)
    return
}
