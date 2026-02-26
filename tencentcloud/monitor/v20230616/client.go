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

package v20230616

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-06-16"

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


func NewCreateNoticeContentTmplRequest() (request *CreateNoticeContentTmplRequest) {
    request = &CreateNoticeContentTmplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateNoticeContentTmpl")
    
    
    return
}

func NewCreateNoticeContentTmplResponse() (response *CreateNoticeContentTmplResponse) {
    response = &CreateNoticeContentTmplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateNoticeContentTmpl
// 创建自定义通知内容模板
func (c *Client) CreateNoticeContentTmpl(request *CreateNoticeContentTmplRequest) (response *CreateNoticeContentTmplResponse, err error) {
    return c.CreateNoticeContentTmplWithContext(context.Background(), request)
}

// CreateNoticeContentTmpl
// 创建自定义通知内容模板
func (c *Client) CreateNoticeContentTmplWithContext(ctx context.Context, request *CreateNoticeContentTmplRequest) (response *CreateNoticeContentTmplResponse, err error) {
    if request == nil {
        request = NewCreateNoticeContentTmplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateNoticeContentTmpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNoticeContentTmpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNoticeContentTmplResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNoticeContentTmplsRequest() (request *DeleteNoticeContentTmplsRequest) {
    request = &DeleteNoticeContentTmplsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteNoticeContentTmpls")
    
    
    return
}

func NewDeleteNoticeContentTmplsResponse() (response *DeleteNoticeContentTmplsResponse) {
    response = &DeleteNoticeContentTmplsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteNoticeContentTmpls
// 删除通知内容模板
func (c *Client) DeleteNoticeContentTmpls(request *DeleteNoticeContentTmplsRequest) (response *DeleteNoticeContentTmplsResponse, err error) {
    return c.DeleteNoticeContentTmplsWithContext(context.Background(), request)
}

// DeleteNoticeContentTmpls
// 删除通知内容模板
func (c *Client) DeleteNoticeContentTmplsWithContext(ctx context.Context, request *DeleteNoticeContentTmplsRequest) (response *DeleteNoticeContentTmplsResponse, err error) {
    if request == nil {
        request = NewDeleteNoticeContentTmplsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteNoticeContentTmpls")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNoticeContentTmpls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNoticeContentTmplsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNotifyHistoriesRequest() (request *DescribeAlarmNotifyHistoriesRequest) {
    request = &DescribeAlarmNotifyHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotifyHistories")
    
    
    return
}

func NewDescribeAlarmNotifyHistoriesResponse() (response *DescribeAlarmNotifyHistoriesResponse) {
    response = &DescribeAlarmNotifyHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotifyHistories
// 按需查询告警的通知历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarmNotifyHistories(request *DescribeAlarmNotifyHistoriesRequest) (response *DescribeAlarmNotifyHistoriesResponse, err error) {
    return c.DescribeAlarmNotifyHistoriesWithContext(context.Background(), request)
}

// DescribeAlarmNotifyHistories
// 按需查询告警的通知历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlarmNotifyHistoriesWithContext(ctx context.Context, request *DescribeAlarmNotifyHistoriesRequest) (response *DescribeAlarmNotifyHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNotifyHistoriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNotifyHistories")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotifyHistories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNotifyHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNoticeContentTmplRequest() (request *DescribeNoticeContentTmplRequest) {
    request = &DescribeNoticeContentTmplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeNoticeContentTmpl")
    
    
    return
}

func NewDescribeNoticeContentTmplResponse() (response *DescribeNoticeContentTmplResponse) {
    response = &DescribeNoticeContentTmplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNoticeContentTmpl
// 根据查询条件获取自定义通知内容模板，若所有查询条件空，则获取账号下所有模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNoticeContentTmpl(request *DescribeNoticeContentTmplRequest) (response *DescribeNoticeContentTmplResponse, err error) {
    return c.DescribeNoticeContentTmplWithContext(context.Background(), request)
}

// DescribeNoticeContentTmpl
// 根据查询条件获取自定义通知内容模板，若所有查询条件空，则获取账号下所有模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNoticeContentTmplWithContext(ctx context.Context, request *DescribeNoticeContentTmplRequest) (response *DescribeNoticeContentTmplResponse, err error) {
    if request == nil {
        request = NewDescribeNoticeContentTmplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeNoticeContentTmpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNoticeContentTmpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNoticeContentTmplResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNoticeContentTmplRequest() (request *ModifyNoticeContentTmplRequest) {
    request = &ModifyNoticeContentTmplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyNoticeContentTmpl")
    
    
    return
}

func NewModifyNoticeContentTmplResponse() (response *ModifyNoticeContentTmplResponse) {
    response = &ModifyNoticeContentTmplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNoticeContentTmpl
// 修改通知内容模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyNoticeContentTmpl(request *ModifyNoticeContentTmplRequest) (response *ModifyNoticeContentTmplResponse, err error) {
    return c.ModifyNoticeContentTmplWithContext(context.Background(), request)
}

// ModifyNoticeContentTmpl
// 修改通知内容模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyNoticeContentTmplWithContext(ctx context.Context, request *ModifyNoticeContentTmplRequest) (response *ModifyNoticeContentTmplResponse, err error) {
    if request == nil {
        request = NewModifyNoticeContentTmplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyNoticeContentTmpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNoticeContentTmpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNoticeContentTmplResponse()
    err = c.Send(request, response)
    return
}
