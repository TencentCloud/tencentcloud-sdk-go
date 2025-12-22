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


func NewConvert3DFormatRequest() (request *Convert3DFormatRequest) {
    request = &Convert3DFormatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "Convert3DFormat")
    
    
    return
}

func NewConvert3DFormatResponse() (response *Convert3DFormatResponse) {
    response = &Convert3DFormatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Convert3DFormat
// 输入3D模型文件后，可进行3D模型文件格式转换。
func (c *Client) Convert3DFormat(request *Convert3DFormatRequest) (response *Convert3DFormatResponse, err error) {
    return c.Convert3DFormatWithContext(context.Background(), request)
}

// Convert3DFormat
// 输入3D模型文件后，可进行3D模型文件格式转换。
func (c *Client) Convert3DFormatWithContext(ctx context.Context, request *Convert3DFormatRequest) (response *Convert3DFormatResponse, err error) {
    if request == nil {
        request = NewConvert3DFormatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "Convert3DFormat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("Convert3DFormat require credential")
    }

    request.SetContext(ctx)
    
    response = NewConvert3DFormatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHunyuanTo3DUVJobRequest() (request *DescribeHunyuanTo3DUVJobRequest) {
    request = &DescribeHunyuanTo3DUVJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "DescribeHunyuanTo3DUVJob")
    
    
    return
}

func NewDescribeHunyuanTo3DUVJobResponse() (response *DescribeHunyuanTo3DUVJobResponse) {
    response = &DescribeHunyuanTo3DUVJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHunyuanTo3DUVJob
// 查询组件拆分任务。
func (c *Client) DescribeHunyuanTo3DUVJob(request *DescribeHunyuanTo3DUVJobRequest) (response *DescribeHunyuanTo3DUVJobResponse, err error) {
    return c.DescribeHunyuanTo3DUVJobWithContext(context.Background(), request)
}

// DescribeHunyuanTo3DUVJob
// 查询组件拆分任务。
func (c *Client) DescribeHunyuanTo3DUVJobWithContext(ctx context.Context, request *DescribeHunyuanTo3DUVJobRequest) (response *DescribeHunyuanTo3DUVJobResponse, err error) {
    if request == nil {
        request = NewDescribeHunyuanTo3DUVJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "DescribeHunyuanTo3DUVJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHunyuanTo3DUVJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHunyuanTo3DUVJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReduceFaceJobRequest() (request *DescribeReduceFaceJobRequest) {
    request = &DescribeReduceFaceJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "DescribeReduceFaceJob")
    
    
    return
}

func NewDescribeReduceFaceJobResponse() (response *DescribeReduceFaceJobResponse) {
    response = &DescribeReduceFaceJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReduceFaceJob
// 混元生3D接口，采用 Polygon 1.5模型，输入3D 高模后，可生成布线规整，较低面数的3D 模型。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) DescribeReduceFaceJob(request *DescribeReduceFaceJobRequest) (response *DescribeReduceFaceJobResponse, err error) {
    return c.DescribeReduceFaceJobWithContext(context.Background(), request)
}

// DescribeReduceFaceJob
// 混元生3D接口，采用 Polygon 1.5模型，输入3D 高模后，可生成布线规整，较低面数的3D 模型。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) DescribeReduceFaceJobWithContext(ctx context.Context, request *DescribeReduceFaceJobRequest) (response *DescribeReduceFaceJobResponse, err error) {
    if request == nil {
        request = NewDescribeReduceFaceJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "DescribeReduceFaceJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReduceFaceJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReduceFaceJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextureTo3DJobRequest() (request *DescribeTextureTo3DJobRequest) {
    request = &DescribeTextureTo3DJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "DescribeTextureTo3DJob")
    
    
    return
}

func NewDescribeTextureTo3DJobResponse() (response *DescribeTextureTo3DJobResponse) {
    response = &DescribeTextureTo3DJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTextureTo3DJob
// 混元生3D接口，输入单几何模型和参考图或文字描述后，可生成对应的纹理贴图。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) DescribeTextureTo3DJob(request *DescribeTextureTo3DJobRequest) (response *DescribeTextureTo3DJobResponse, err error) {
    return c.DescribeTextureTo3DJobWithContext(context.Background(), request)
}

// DescribeTextureTo3DJob
// 混元生3D接口，输入单几何模型和参考图或文字描述后，可生成对应的纹理贴图。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) DescribeTextureTo3DJobWithContext(ctx context.Context, request *DescribeTextureTo3DJobRequest) (response *DescribeTextureTo3DJobResponse, err error) {
    if request == nil {
        request = NewDescribeTextureTo3DJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "DescribeTextureTo3DJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTextureTo3DJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTextureTo3DJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryHunyuan3DPartJobRequest() (request *QueryHunyuan3DPartJobRequest) {
    request = &QueryHunyuan3DPartJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "QueryHunyuan3DPartJob")
    
    
    return
}

func NewQueryHunyuan3DPartJobResponse() (response *QueryHunyuan3DPartJobResponse) {
    response = &QueryHunyuan3DPartJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryHunyuan3DPartJob
// 查询组件生成任务。
func (c *Client) QueryHunyuan3DPartJob(request *QueryHunyuan3DPartJobRequest) (response *QueryHunyuan3DPartJobResponse, err error) {
    return c.QueryHunyuan3DPartJobWithContext(context.Background(), request)
}

// QueryHunyuan3DPartJob
// 查询组件生成任务。
func (c *Client) QueryHunyuan3DPartJobWithContext(ctx context.Context, request *QueryHunyuan3DPartJobRequest) (response *QueryHunyuan3DPartJobResponse, err error) {
    if request == nil {
        request = NewQueryHunyuan3DPartJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "QueryHunyuan3DPartJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryHunyuan3DPartJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryHunyuan3DPartJobResponse()
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

func NewSubmitHunyuan3DPartJobRequest() (request *SubmitHunyuan3DPartJobRequest) {
    request = &SubmitHunyuan3DPartJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "SubmitHunyuan3DPartJob")
    
    
    return
}

func NewSubmitHunyuan3DPartJobResponse() (response *SubmitHunyuan3DPartJobResponse) {
    response = &SubmitHunyuan3DPartJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuan3DPartJob
// 输入3D模型文件后，根据模型结构自动进行组件识别生成。
func (c *Client) SubmitHunyuan3DPartJob(request *SubmitHunyuan3DPartJobRequest) (response *SubmitHunyuan3DPartJobResponse, err error) {
    return c.SubmitHunyuan3DPartJobWithContext(context.Background(), request)
}

// SubmitHunyuan3DPartJob
// 输入3D模型文件后，根据模型结构自动进行组件识别生成。
func (c *Client) SubmitHunyuan3DPartJobWithContext(ctx context.Context, request *SubmitHunyuan3DPartJobRequest) (response *SubmitHunyuan3DPartJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuan3DPartJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "SubmitHunyuan3DPartJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuan3DPartJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuan3DPartJobResponse()
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

func NewSubmitHunyuanTo3DUVJobRequest() (request *SubmitHunyuanTo3DUVJobRequest) {
    request = &SubmitHunyuanTo3DUVJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "SubmitHunyuanTo3DUVJob")
    
    
    return
}

func NewSubmitHunyuanTo3DUVJobResponse() (response *SubmitHunyuanTo3DUVJobResponse) {
    response = &SubmitHunyuanTo3DUVJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanTo3DUVJob
// 输入模型后，可根据模型纹理进行UV展开，输出对应UV贴图。
func (c *Client) SubmitHunyuanTo3DUVJob(request *SubmitHunyuanTo3DUVJobRequest) (response *SubmitHunyuanTo3DUVJobResponse, err error) {
    return c.SubmitHunyuanTo3DUVJobWithContext(context.Background(), request)
}

// SubmitHunyuanTo3DUVJob
// 输入模型后，可根据模型纹理进行UV展开，输出对应UV贴图。
func (c *Client) SubmitHunyuanTo3DUVJobWithContext(ctx context.Context, request *SubmitHunyuanTo3DUVJobRequest) (response *SubmitHunyuanTo3DUVJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanTo3DUVJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "SubmitHunyuanTo3DUVJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanTo3DUVJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanTo3DUVJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitReduceFaceJobRequest() (request *SubmitReduceFaceJobRequest) {
    request = &SubmitReduceFaceJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "SubmitReduceFaceJob")
    
    
    return
}

func NewSubmitReduceFaceJobResponse() (response *SubmitReduceFaceJobResponse) {
    response = &SubmitReduceFaceJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitReduceFaceJob
// 混元生3D接口，采用 Polygon 1.5模型，输入3D 高模后，可生成布线规整，较低面数的3D 模型。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitReduceFaceJob(request *SubmitReduceFaceJobRequest) (response *SubmitReduceFaceJobResponse, err error) {
    return c.SubmitReduceFaceJobWithContext(context.Background(), request)
}

// SubmitReduceFaceJob
// 混元生3D接口，采用 Polygon 1.5模型，输入3D 高模后，可生成布线规整，较低面数的3D 模型。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitReduceFaceJobWithContext(ctx context.Context, request *SubmitReduceFaceJobRequest) (response *SubmitReduceFaceJobResponse, err error) {
    if request == nil {
        request = NewSubmitReduceFaceJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "SubmitReduceFaceJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitReduceFaceJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitReduceFaceJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTextureTo3DJobRequest() (request *SubmitTextureTo3DJobRequest) {
    request = &SubmitTextureTo3DJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ai3d", APIVersion, "SubmitTextureTo3DJob")
    
    
    return
}

func NewSubmitTextureTo3DJobResponse() (response *SubmitTextureTo3DJobResponse) {
    response = &SubmitTextureTo3DJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTextureTo3DJob
// 混元生3D接口，输入单几何模型和参考图或文字描述后，可生成对应的纹理贴图。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitTextureTo3DJob(request *SubmitTextureTo3DJobRequest) (response *SubmitTextureTo3DJobResponse, err error) {
    return c.SubmitTextureTo3DJobWithContext(context.Background(), request)
}

// SubmitTextureTo3DJob
// 混元生3D接口，输入单几何模型和参考图或文字描述后，可生成对应的纹理贴图。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
func (c *Client) SubmitTextureTo3DJobWithContext(ctx context.Context, request *SubmitTextureTo3DJobRequest) (response *SubmitTextureTo3DJobResponse, err error) {
    if request == nil {
        request = NewSubmitTextureTo3DJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ai3d", APIVersion, "SubmitTextureTo3DJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTextureTo3DJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTextureTo3DJobResponse()
    err = c.Send(request, response)
    return
}
