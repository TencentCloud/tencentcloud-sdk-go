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

package v20190627

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-27"

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


func NewTextProcessRequest() (request *TextProcessRequest) {
    request = &TextProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbp", APIVersion, "TextProcess")
    
    
    return
}

func NewTextProcessResponse() (response *TextProcessResponse) {
    response = &TextProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextProcess
// 接收调用侧的文本输入，返回应答文本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORMMS = "InternalError.ErrorMms"
//  INTERNALERROR_ERRORNLU = "InternalError.ErrorNlu"
//  INTERNALERROR_ERRORRPC = "InternalError.ErrorRpc"
//  INTERNALERROR_ERRORWEBHOOK = "InternalError.ErrorWebHook"
//  INTERNALERROR_NOAPPPRIVILEGE = "InternalError.NoAppPrivilege"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextProcess(request *TextProcessRequest) (response *TextProcessResponse, err error) {
    return c.TextProcessWithContext(context.Background(), request)
}

// TextProcess
// 接收调用侧的文本输入，返回应答文本。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORMMS = "InternalError.ErrorMms"
//  INTERNALERROR_ERRORNLU = "InternalError.ErrorNlu"
//  INTERNALERROR_ERRORRPC = "InternalError.ErrorRpc"
//  INTERNALERROR_ERRORWEBHOOK = "InternalError.ErrorWebHook"
//  INTERNALERROR_NOAPPPRIVILEGE = "InternalError.NoAppPrivilege"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextProcessWithContext(ctx context.Context, request *TextProcessRequest) (response *TextProcessResponse, err error) {
    if request == nil {
        request = NewTextProcessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextProcess require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextProcessResponse()
    err = c.Send(request, response)
    return
}

func NewTextResetRequest() (request *TextResetRequest) {
    request = &TextResetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbp", APIVersion, "TextReset")
    
    
    return
}

func NewTextResetResponse() (response *TextResetResponse) {
    response = &TextResetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextReset
// 会话重置接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORMMS = "InternalError.ErrorMms"
//  INTERNALERROR_ERRORNLU = "InternalError.ErrorNlu"
//  INTERNALERROR_ERRORRPC = "InternalError.ErrorRpc"
//  INTERNALERROR_ERRORWEBHOOK = "InternalError.ErrorWebHook"
//  INTERNALERROR_NOAPPPRIVILEGE = "InternalError.NoAppPrivilege"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextReset(request *TextResetRequest) (response *TextResetResponse, err error) {
    return c.TextResetWithContext(context.Background(), request)
}

// TextReset
// 会话重置接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORMMS = "InternalError.ErrorMms"
//  INTERNALERROR_ERRORNLU = "InternalError.ErrorNlu"
//  INTERNALERROR_ERRORRPC = "InternalError.ErrorRpc"
//  INTERNALERROR_ERRORWEBHOOK = "InternalError.ErrorWebHook"
//  INTERNALERROR_NOAPPPRIVILEGE = "InternalError.NoAppPrivilege"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TextResetWithContext(ctx context.Context, request *TextResetRequest) (response *TextResetResponse, err error) {
    if request == nil {
        request = NewTextResetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextReset require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextResetResponse()
    err = c.Send(request, response)
    return
}
