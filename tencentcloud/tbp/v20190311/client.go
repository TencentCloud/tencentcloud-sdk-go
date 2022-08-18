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

package v20190311

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-11"

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


func NewCreateBotRequest() (request *CreateBotRequest) {
    request = &CreateBotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbp", APIVersion, "CreateBot")
    
    
    return
}

func NewCreateBotResponse() (response *CreateBotResponse) {
    response = &CreateBotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBot
// 创建机器人
//
// 可能返回的错误码:
//  INTERNALERROR_CALLMMSFAILED = "InternalError.CallMMSFailed"
//  INTERNALERROR_MMSINTERNALERROR = "InternalError.MMSInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateBot(request *CreateBotRequest) (response *CreateBotResponse, err error) {
    return c.CreateBotWithContext(context.Background(), request)
}

// CreateBot
// 创建机器人
//
// 可能返回的错误码:
//  INTERNALERROR_CALLMMSFAILED = "InternalError.CallMMSFailed"
//  INTERNALERROR_MMSINTERNALERROR = "InternalError.MMSInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateBotWithContext(ctx context.Context, request *CreateBotRequest) (response *CreateBotResponse, err error) {
    if request == nil {
        request = NewCreateBotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBotResponse()
    err = c.Send(request, response)
    return
}

func NewResetRequest() (request *ResetRequest) {
    request = &ResetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbp", APIVersion, "Reset")
    
    
    return
}

func NewResetResponse() (response *ResetResponse) {
    response = &ResetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Reset
// 对当前机器人的会话状态进行复位
//
// 可能返回的错误码:
//  INTERNALERROR_ERRORASR = "InternalError.ErrorAsr"
//  INTERNALERROR_ERRORMMS = "InternalError.ErrorMms"
//  INTERNALERROR_ERRORRPC = "InternalError.ErrorRpc"
//  INTERNALERROR_ERRORTTS = "InternalError.ErrorTts"
//  INTERNALERROR_NOAPPPRIVILEGE = "InternalError.NoAppPrivilege"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) Reset(request *ResetRequest) (response *ResetResponse, err error) {
    return c.ResetWithContext(context.Background(), request)
}

// Reset
// 对当前机器人的会话状态进行复位
//
// 可能返回的错误码:
//  INTERNALERROR_ERRORASR = "InternalError.ErrorAsr"
//  INTERNALERROR_ERRORMMS = "InternalError.ErrorMms"
//  INTERNALERROR_ERRORRPC = "InternalError.ErrorRpc"
//  INTERNALERROR_ERRORTTS = "InternalError.ErrorTts"
//  INTERNALERROR_NOAPPPRIVILEGE = "InternalError.NoAppPrivilege"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ResetWithContext(ctx context.Context, request *ResetRequest) (response *ResetResponse, err error) {
    if request == nil {
        request = NewResetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Reset require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetResponse()
    err = c.Send(request, response)
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
// 接收调用侧的文本输入，返回应答文本。已废弃，推荐使用最新版TextProcess接口。
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
// 接收调用侧的文本输入，返回应答文本。已废弃，推荐使用最新版TextProcess接口。
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
// 会话重置接口。已废弃，推荐使用最新版TextReset接口。
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
// 会话重置接口。已废弃，推荐使用最新版TextReset接口。
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
