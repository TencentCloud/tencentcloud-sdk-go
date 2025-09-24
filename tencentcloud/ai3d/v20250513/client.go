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

package v20250513

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-05-13"

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


func NewQueryHunyuanTo3DJobRequest() (request *QueryHunyuanTo3DJobRequest) {
    request = &QueryHunyuanTo3DJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "QueryHunyuanTo3DJob")
    
    
    return
}

func NewQueryHunyuanTo3DJobResponse() (response *QueryHunyuanTo3DJobResponse) {
    response = &QueryHunyuanTo3DJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanTo3DJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) QueryHunyuanTo3DJob(request *QueryHunyuanTo3DJobRequest) (response *QueryHunyuanTo3DJobResponse, err error) {
    return c.QueryHunyuanTo3DJobWithContext(context.Background(), request)
}

// QueryHunyuanTo3DJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) QueryHunyuanTo3DJobWithContext(ctx context.Context, request *QueryHunyuanTo3DJobRequest) (response *QueryHunyuanTo3DJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanTo3DJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "QueryHunyuanTo3DJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanTo3DJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanTo3DJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuanTo3DProJobRequest() (request *QueryHunyuanTo3DProJobRequest) {
    request = &QueryHunyuanTo3DProJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "QueryHunyuanTo3DProJob")
    
    
    return
}

func NewQueryHunyuanTo3DProJobResponse() (response *QueryHunyuanTo3DProJobResponse) {
    response = &QueryHunyuanTo3DProJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanTo3DProJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供3个并发，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) QueryHunyuanTo3DProJob(request *QueryHunyuanTo3DProJobRequest) (response *QueryHunyuanTo3DProJobResponse, err error) {
    return c.QueryHunyuanTo3DProJobWithContext(context.Background(), request)
}

// QueryHunyuanTo3DProJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供3个并发，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) QueryHunyuanTo3DProJobWithContext(ctx context.Context, request *QueryHunyuanTo3DProJobRequest) (response *QueryHunyuanTo3DProJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanTo3DProJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "QueryHunyuanTo3DProJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanTo3DProJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanTo3DProJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuanTo3DRapidJobRequest() (request *QueryHunyuanTo3DRapidJobRequest) {
    request = &QueryHunyuanTo3DRapidJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "QueryHunyuanTo3DRapidJob")
    
    
    return
}

func NewQueryHunyuanTo3DRapidJobResponse() (response *QueryHunyuanTo3DRapidJobResponse) {
    response = &QueryHunyuanTo3DRapidJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuanTo3DRapidJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) QueryHunyuanTo3DRapidJob(request *QueryHunyuanTo3DRapidJobRequest) (response *QueryHunyuanTo3DRapidJobResponse, err error) {
    return c.QueryHunyuanTo3DRapidJobWithContext(context.Background(), request)
}

// QueryHunyuanTo3DRapidJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) QueryHunyuanTo3DRapidJobWithContext(ctx context.Context, request *QueryHunyuanTo3DRapidJobRequest) (response *QueryHunyuanTo3DRapidJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuanTo3DRapidJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "QueryHunyuanTo3DRapidJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuanTo3DRapidJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuanTo3DRapidJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuanTo3DJobRequest() (request *SubmitHunyuanTo3DJobRequest) {
    request = &SubmitHunyuanTo3DJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "SubmitHunyuanTo3DJob")
    
    
    return
}

func NewSubmitHunyuanTo3DJobResponse() (response *SubmitHunyuanTo3DJobResponse) {
    response = &SubmitHunyuanTo3DJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanTo3DJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitHunyuanTo3DJob(request *SubmitHunyuanTo3DJobRequest) (response *SubmitHunyuanTo3DJobResponse, err error) {
    return c.SubmitHunyuanTo3DJobWithContext(context.Background(), request)
}

// SubmitHunyuanTo3DJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitHunyuanTo3DJobWithContext(ctx context.Context, request *SubmitHunyuanTo3DJobRequest) (response *SubmitHunyuanTo3DJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanTo3DJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "SubmitHunyuanTo3DJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanTo3DJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanTo3DJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuanTo3DProJobRequest() (request *SubmitHunyuanTo3DProJobRequest) {
    request = &SubmitHunyuanTo3DProJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "SubmitHunyuanTo3DProJob")
    
    
    return
}

func NewSubmitHunyuanTo3DProJobResponse() (response *SubmitHunyuanTo3DProJobResponse) {
    response = &SubmitHunyuanTo3DProJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanTo3DProJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供3个并发，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitHunyuanTo3DProJob(request *SubmitHunyuanTo3DProJobRequest) (response *SubmitHunyuanTo3DProJobResponse, err error) {
    return c.SubmitHunyuanTo3DProJobWithContext(context.Background(), request)
}

// SubmitHunyuanTo3DProJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供3个并发，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitHunyuanTo3DProJobWithContext(ctx context.Context, request *SubmitHunyuanTo3DProJobRequest) (response *SubmitHunyuanTo3DProJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanTo3DProJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "SubmitHunyuanTo3DProJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanTo3DProJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanTo3DProJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuanTo3DRapidJobRequest() (request *SubmitHunyuanTo3DRapidJobRequest) {
    request = &SubmitHunyuanTo3DRapidJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "SubmitHunyuanTo3DRapidJob")
    
    
    return
}

func NewSubmitHunyuanTo3DRapidJobResponse() (response *SubmitHunyuanTo3DRapidJobResponse) {
    response = &SubmitHunyuanTo3DRapidJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanTo3DRapidJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitHunyuanTo3DRapidJob(request *SubmitHunyuanTo3DRapidJobRequest) (response *SubmitHunyuanTo3DRapidJobResponse, err error) {
    return c.SubmitHunyuanTo3DRapidJobWithContext(context.Background(), request)
}

// SubmitHunyuanTo3DRapidJob
// 混元生3D接口，基于混元大模型，根据输入的文本描述/图片智能生成3D。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitHunyuanTo3DRapidJobWithContext(ctx context.Context, request *SubmitHunyuanTo3DRapidJobRequest) (response *SubmitHunyuanTo3DRapidJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanTo3DRapidJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "SubmitHunyuanTo3DRapidJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanTo3DRapidJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanTo3DRapidJobResponse()
    err = c.Send(request, response)
    return
}
