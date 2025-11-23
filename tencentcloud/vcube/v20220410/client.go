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

package v20220410

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-04-10"

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


func NewCreateActivityLicenseRequest() (request *CreateActivityLicenseRequest) {
    request = &CreateActivityLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateActivityLicense")
    
    
    return
}

func NewCreateActivityLicenseResponse() (response *CreateActivityLicenseResponse) {
    response = &CreateActivityLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateActivityLicense
// 创建活动license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateActivityLicense(request *CreateActivityLicenseRequest) (response *CreateActivityLicenseResponse, err error) {
    return c.CreateActivityLicenseWithContext(context.Background(), request)
}

// CreateActivityLicense
// 创建活动license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateActivityLicenseWithContext(ctx context.Context, request *CreateActivityLicenseRequest) (response *CreateActivityLicenseResponse, err error) {
    if request == nil {
        request = NewCreateActivityLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateActivityLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateActivityLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateActivityLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationAndBindLicenseRequest() (request *CreateApplicationAndBindLicenseRequest) {
    request = &CreateApplicationAndBindLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateApplicationAndBindLicense")
    
    
    return
}

func NewCreateApplicationAndBindLicenseResponse() (response *CreateApplicationAndBindLicenseResponse) {
    response = &CreateApplicationAndBindLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationAndBindLicense
// 创建应用并绑定license或者XMagic
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateApplicationAndBindLicense(request *CreateApplicationAndBindLicenseRequest) (response *CreateApplicationAndBindLicenseResponse, err error) {
    return c.CreateApplicationAndBindLicenseWithContext(context.Background(), request)
}

// CreateApplicationAndBindLicense
// 创建应用并绑定license或者XMagic
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateApplicationAndBindLicenseWithContext(ctx context.Context, request *CreateApplicationAndBindLicenseRequest) (response *CreateApplicationAndBindLicenseResponse, err error) {
    if request == nil {
        request = NewCreateApplicationAndBindLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateApplicationAndBindLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationAndBindLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationAndBindLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationAndVideoRequest() (request *CreateApplicationAndVideoRequest) {
    request = &CreateApplicationAndVideoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateApplicationAndVideo")
    
    
    return
}

func NewCreateApplicationAndVideoResponse() (response *CreateApplicationAndVideoResponse) {
    response = &CreateApplicationAndVideoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationAndVideo
// 创建应用和视频播放license 目前只有国际站可以用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateApplicationAndVideo(request *CreateApplicationAndVideoRequest) (response *CreateApplicationAndVideoResponse, err error) {
    return c.CreateApplicationAndVideoWithContext(context.Background(), request)
}

// CreateApplicationAndVideo
// 创建应用和视频播放license 目前只有国际站可以用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateApplicationAndVideoWithContext(ctx context.Context, request *CreateApplicationAndVideoRequest) (response *CreateApplicationAndVideoResponse, err error) {
    if request == nil {
        request = NewCreateApplicationAndVideoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateApplicationAndVideo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationAndVideo require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationAndVideoResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationAndWebPlayerLicenseRequest() (request *CreateApplicationAndWebPlayerLicenseRequest) {
    request = &CreateApplicationAndWebPlayerLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateApplicationAndWebPlayerLicense")
    
    
    return
}

func NewCreateApplicationAndWebPlayerLicenseResponse() (response *CreateApplicationAndWebPlayerLicenseResponse) {
    response = &CreateApplicationAndWebPlayerLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationAndWebPlayerLicense
// 创建 web 播放器基础版
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateApplicationAndWebPlayerLicense(request *CreateApplicationAndWebPlayerLicenseRequest) (response *CreateApplicationAndWebPlayerLicenseResponse, err error) {
    return c.CreateApplicationAndWebPlayerLicenseWithContext(context.Background(), request)
}

// CreateApplicationAndWebPlayerLicense
// 创建 web 播放器基础版
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateApplicationAndWebPlayerLicenseWithContext(ctx context.Context, request *CreateApplicationAndWebPlayerLicenseRequest) (response *CreateApplicationAndWebPlayerLicenseResponse, err error) {
    if request == nil {
        request = NewCreateApplicationAndWebPlayerLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateApplicationAndWebPlayerLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationAndWebPlayerLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationAndWebPlayerLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLicenseRequest() (request *CreateLicenseRequest) {
    request = &CreateLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateLicense")
    
    
    return
}

func NewCreateLicenseResponse() (response *CreateLicenseResponse) {
    response = &CreateLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLicense
// 绑定license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateLicense(request *CreateLicenseRequest) (response *CreateLicenseResponse, err error) {
    return c.CreateLicenseWithContext(context.Background(), request)
}

// CreateLicense
// 绑定license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateLicenseWithContext(ctx context.Context, request *CreateLicenseRequest) (response *CreateLicenseResponse, err error) {
    if request == nil {
        request = NewCreateLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTestXMagicRequest() (request *CreateTestXMagicRequest) {
    request = &CreateTestXMagicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateTestXMagic")
    
    
    return
}

func NewCreateTestXMagicResponse() (response *CreateTestXMagicResponse) {
    response = &CreateTestXMagicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTestXMagic
// 申请开通测试版腾讯特效
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateTestXMagic(request *CreateTestXMagicRequest) (response *CreateTestXMagicResponse, err error) {
    return c.CreateTestXMagicWithContext(context.Background(), request)
}

// CreateTestXMagic
// 申请开通测试版腾讯特效
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateTestXMagicWithContext(ctx context.Context, request *CreateTestXMagicRequest) (response *CreateTestXMagicResponse, err error) {
    if request == nil {
        request = NewCreateTestXMagicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateTestXMagic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTestXMagic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTestXMagicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrialApplicationAndLicenseRequest() (request *CreateTrialApplicationAndLicenseRequest) {
    request = &CreateTrialApplicationAndLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateTrialApplicationAndLicense")
    
    
    return
}

func NewCreateTrialApplicationAndLicenseResponse() (response *CreateTrialApplicationAndLicenseResponse) {
    response = &CreateTrialApplicationAndLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTrialApplicationAndLicense
// 创建测试应用并开通测试 license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateTrialApplicationAndLicense(request *CreateTrialApplicationAndLicenseRequest) (response *CreateTrialApplicationAndLicenseResponse, err error) {
    return c.CreateTrialApplicationAndLicenseWithContext(context.Background(), request)
}

// CreateTrialApplicationAndLicense
// 创建测试应用并开通测试 license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateTrialApplicationAndLicenseWithContext(ctx context.Context, request *CreateTrialApplicationAndLicenseRequest) (response *CreateTrialApplicationAndLicenseResponse, err error) {
    if request == nil {
        request = NewCreateTrialApplicationAndLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateTrialApplicationAndLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrialApplicationAndLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrialApplicationAndLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrialLicenseRequest() (request *CreateTrialLicenseRequest) {
    request = &CreateTrialLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateTrialLicense")
    
    
    return
}

func NewCreateTrialLicenseResponse() (response *CreateTrialLicenseResponse) {
    response = &CreateTrialLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTrialLicense
// 开通测试license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateTrialLicense(request *CreateTrialLicenseRequest) (response *CreateTrialLicenseResponse, err error) {
    return c.CreateTrialLicenseWithContext(context.Background(), request)
}

// CreateTrialLicense
// 开通测试license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateTrialLicenseWithContext(ctx context.Context, request *CreateTrialLicenseRequest) (response *CreateTrialLicenseResponse, err error) {
    if request == nil {
        request = NewCreateTrialLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateTrialLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrialLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrialLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateXMagicRequest() (request *CreateXMagicRequest) {
    request = &CreateXMagicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "CreateXMagic")
    
    
    return
}

func NewCreateXMagicResponse() (response *CreateXMagicResponse) {
    response = &CreateXMagicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateXMagic
// x08开通正式版优图美视功能，针对已经有Application的情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateXMagic(request *CreateXMagicRequest) (response *CreateXMagicResponse, err error) {
    return c.CreateXMagicWithContext(context.Background(), request)
}

// CreateXMagic
// x08开通正式版优图美视功能，针对已经有Application的情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateXMagicWithContext(ctx context.Context, request *CreateXMagicRequest) (response *CreateXMagicResponse, err error) {
    if request == nil {
        request = NewCreateXMagicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "CreateXMagic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateXMagic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateXMagicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationAndVideoLicenseRequest() (request *DeleteApplicationAndVideoLicenseRequest) {
    request = &DeleteApplicationAndVideoLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DeleteApplicationAndVideoLicense")
    
    
    return
}

func NewDeleteApplicationAndVideoLicenseResponse() (response *DeleteApplicationAndVideoLicenseResponse) {
    response = &DeleteApplicationAndVideoLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationAndVideoLicense
// 删除视频播放器 License 和相关应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteApplicationAndVideoLicense(request *DeleteApplicationAndVideoLicenseRequest) (response *DeleteApplicationAndVideoLicenseResponse, err error) {
    return c.DeleteApplicationAndVideoLicenseWithContext(context.Background(), request)
}

// DeleteApplicationAndVideoLicense
// 删除视频播放器 License 和相关应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteApplicationAndVideoLicenseWithContext(ctx context.Context, request *DeleteApplicationAndVideoLicenseRequest) (response *DeleteApplicationAndVideoLicenseResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationAndVideoLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DeleteApplicationAndVideoLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationAndVideoLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationAndVideoLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationAndWebPlayerLicenseRequest() (request *DeleteApplicationAndWebPlayerLicenseRequest) {
    request = &DeleteApplicationAndWebPlayerLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DeleteApplicationAndWebPlayerLicense")
    
    
    return
}

func NewDeleteApplicationAndWebPlayerLicenseResponse() (response *DeleteApplicationAndWebPlayerLicenseResponse) {
    response = &DeleteApplicationAndWebPlayerLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationAndWebPlayerLicense
// 删除web播放器license和应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteApplicationAndWebPlayerLicense(request *DeleteApplicationAndWebPlayerLicenseRequest) (response *DeleteApplicationAndWebPlayerLicenseResponse, err error) {
    return c.DeleteApplicationAndWebPlayerLicenseWithContext(context.Background(), request)
}

// DeleteApplicationAndWebPlayerLicense
// 删除web播放器license和应用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteApplicationAndWebPlayerLicenseWithContext(ctx context.Context, request *DeleteApplicationAndWebPlayerLicenseRequest) (response *DeleteApplicationAndWebPlayerLicenseResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationAndWebPlayerLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DeleteApplicationAndWebPlayerLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationAndWebPlayerLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationAndWebPlayerLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFeatureListRequest() (request *DescribeFeatureListRequest) {
    request = &DescribeFeatureListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeFeatureList")
    
    
    return
}

func NewDescribeFeatureListResponse() (response *DescribeFeatureListResponse) {
    response = &DescribeFeatureListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFeatureList
// 查询功能列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFeatureList(request *DescribeFeatureListRequest) (response *DescribeFeatureListResponse, err error) {
    return c.DescribeFeatureListWithContext(context.Background(), request)
}

// DescribeFeatureList
// 查询功能列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeFeatureListWithContext(ctx context.Context, request *DescribeFeatureListRequest) (response *DescribeFeatureListResponse, err error) {
    if request == nil {
        request = NewDescribeFeatureListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeFeatureList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFeatureList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFeatureListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLicenseListRequest() (request *DescribeLicenseListRequest) {
    request = &DescribeLicenseListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeLicenseList")
    
    
    return
}

func NewDescribeLicenseListResponse() (response *DescribeLicenseListResponse) {
    response = &DescribeLicenseListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLicenseList
// 总览页查询临期License列表，和统计数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLicenseList(request *DescribeLicenseListRequest) (response *DescribeLicenseListResponse, err error) {
    return c.DescribeLicenseListWithContext(context.Background(), request)
}

// DescribeLicenseList
// 总览页查询临期License列表，和统计数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeLicenseListWithContext(ctx context.Context, request *DescribeLicenseListRequest) (response *DescribeLicenseListResponse, err error) {
    if request == nil {
        request = NewDescribeLicenseListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeLicenseList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLicenseList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLicenseListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNewsRequest() (request *DescribeNewsRequest) {
    request = &DescribeNewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeNews")
    
    
    return
}

func NewDescribeNewsResponse() (response *DescribeNewsResponse) {
    response = &DescribeNewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNews
// 查询产品动态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeNews(request *DescribeNewsRequest) (response *DescribeNewsResponse, err error) {
    return c.DescribeNewsWithContext(context.Background(), request)
}

// DescribeNews
// 查询产品动态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeNewsWithContext(ctx context.Context, request *DescribeNewsRequest) (response *DescribeNewsResponse, err error) {
    if request == nil {
        request = NewDescribeNewsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeNews")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSTSRequest() (request *DescribeSTSRequest) {
    request = &DescribeSTSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeSTS")
    
    
    return
}

func NewDescribeSTSResponse() (response *DescribeSTSResponse) {
    response = &DescribeSTSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSTS
// 获取临时秘钥，用于图片，特效包上传
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSTS(request *DescribeSTSRequest) (response *DescribeSTSResponse, err error) {
    return c.DescribeSTSWithContext(context.Background(), request)
}

// DescribeSTS
// 获取临时秘钥，用于图片，特效包上传
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeSTSWithContext(ctx context.Context, request *DescribeSTSRequest) (response *DescribeSTSResponse, err error) {
    if request == nil {
        request = NewDescribeSTSRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeSTS")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSTS require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSTSResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrialFeatureRequest() (request *DescribeTrialFeatureRequest) {
    request = &DescribeTrialFeatureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeTrialFeature")
    
    
    return
}

func NewDescribeTrialFeatureResponse() (response *DescribeTrialFeatureResponse) {
    response = &DescribeTrialFeatureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrialFeature
// 查询测试应用可以开通的功能
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTrialFeature(request *DescribeTrialFeatureRequest) (response *DescribeTrialFeatureResponse, err error) {
    return c.DescribeTrialFeatureWithContext(context.Background(), request)
}

// DescribeTrialFeature
// 查询测试应用可以开通的功能
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTrialFeatureWithContext(ctx context.Context, request *DescribeTrialFeatureRequest) (response *DescribeTrialFeatureResponse, err error) {
    if request == nil {
        request = NewDescribeTrialFeatureRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeTrialFeature")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrialFeature require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrialFeatureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserConfigRequest() (request *DescribeUserConfigRequest) {
    request = &DescribeUserConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeUserConfig")
    
    
    return
}

func NewDescribeUserConfigResponse() (response *DescribeUserConfigResponse) {
    response = &DescribeUserConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserConfig
// 查询用户个性配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUserConfig(request *DescribeUserConfigRequest) (response *DescribeUserConfigResponse, err error) {
    return c.DescribeUserConfigWithContext(context.Background(), request)
}

// DescribeUserConfig
// 查询用户个性配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUserConfigWithContext(ctx context.Context, request *DescribeUserConfigRequest) (response *DescribeUserConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUserConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeUserConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVcubeApplicationAndLicenseRequest() (request *DescribeVcubeApplicationAndLicenseRequest) {
    request = &DescribeVcubeApplicationAndLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeVcubeApplicationAndLicense")
    
    
    return
}

func NewDescribeVcubeApplicationAndLicenseResponse() (response *DescribeVcubeApplicationAndLicenseResponse) {
    response = &DescribeVcubeApplicationAndLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVcubeApplicationAndLicense
// 查询用户license， 按照应用分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeApplicationAndLicense(request *DescribeVcubeApplicationAndLicenseRequest) (response *DescribeVcubeApplicationAndLicenseResponse, err error) {
    return c.DescribeVcubeApplicationAndLicenseWithContext(context.Background(), request)
}

// DescribeVcubeApplicationAndLicense
// 查询用户license， 按照应用分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeApplicationAndLicenseWithContext(ctx context.Context, request *DescribeVcubeApplicationAndLicenseRequest) (response *DescribeVcubeApplicationAndLicenseResponse, err error) {
    if request == nil {
        request = NewDescribeVcubeApplicationAndLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeVcubeApplicationAndLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVcubeApplicationAndLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVcubeApplicationAndLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVcubeApplicationAndPlayListRequest() (request *DescribeVcubeApplicationAndPlayListRequest) {
    request = &DescribeVcubeApplicationAndPlayListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeVcubeApplicationAndPlayList")
    
    
    return
}

func NewDescribeVcubeApplicationAndPlayListResponse() (response *DescribeVcubeApplicationAndPlayListResponse) {
    response = &DescribeVcubeApplicationAndPlayListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVcubeApplicationAndPlayList
// 查询用户点播直播等license， 按照应用分类,国际站专用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeApplicationAndPlayList(request *DescribeVcubeApplicationAndPlayListRequest) (response *DescribeVcubeApplicationAndPlayListResponse, err error) {
    return c.DescribeVcubeApplicationAndPlayListWithContext(context.Background(), request)
}

// DescribeVcubeApplicationAndPlayList
// 查询用户点播直播等license， 按照应用分类,国际站专用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeApplicationAndPlayListWithContext(ctx context.Context, request *DescribeVcubeApplicationAndPlayListRequest) (response *DescribeVcubeApplicationAndPlayListResponse, err error) {
    if request == nil {
        request = NewDescribeVcubeApplicationAndPlayListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeVcubeApplicationAndPlayList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVcubeApplicationAndPlayList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVcubeApplicationAndPlayListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVcubeApplicationAndXMagicListRequest() (request *DescribeVcubeApplicationAndXMagicListRequest) {
    request = &DescribeVcubeApplicationAndXMagicListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeVcubeApplicationAndXMagicList")
    
    
    return
}

func NewDescribeVcubeApplicationAndXMagicListResponse() (response *DescribeVcubeApplicationAndXMagicListResponse) {
    response = &DescribeVcubeApplicationAndXMagicListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVcubeApplicationAndXMagicList
// 查询用户优图license， 按照应用分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeApplicationAndXMagicList(request *DescribeVcubeApplicationAndXMagicListRequest) (response *DescribeVcubeApplicationAndXMagicListResponse, err error) {
    return c.DescribeVcubeApplicationAndXMagicListWithContext(context.Background(), request)
}

// DescribeVcubeApplicationAndXMagicList
// 查询用户优图license， 按照应用分类
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeApplicationAndXMagicListWithContext(ctx context.Context, request *DescribeVcubeApplicationAndXMagicListRequest) (response *DescribeVcubeApplicationAndXMagicListResponse, err error) {
    if request == nil {
        request = NewDescribeVcubeApplicationAndXMagicListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeVcubeApplicationAndXMagicList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVcubeApplicationAndXMagicList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVcubeApplicationAndXMagicListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVcubeResourcesRequest() (request *DescribeVcubeResourcesRequest) {
    request = &DescribeVcubeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeVcubeResources")
    
    
    return
}

func NewDescribeVcubeResourcesResponse() (response *DescribeVcubeResourcesResponse) {
    response = &DescribeVcubeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVcubeResources
// 查询视立方 license 资源，所有的资源包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeResources(request *DescribeVcubeResourcesRequest) (response *DescribeVcubeResourcesResponse, err error) {
    return c.DescribeVcubeResourcesWithContext(context.Background(), request)
}

// DescribeVcubeResources
// 查询视立方 license 资源，所有的资源包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeResourcesWithContext(ctx context.Context, request *DescribeVcubeResourcesRequest) (response *DescribeVcubeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeVcubeResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeVcubeResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVcubeResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVcubeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVcubeResourcesListRequest() (request *DescribeVcubeResourcesListRequest) {
    request = &DescribeVcubeResourcesListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeVcubeResourcesList")
    
    
    return
}

func NewDescribeVcubeResourcesListResponse() (response *DescribeVcubeResourcesListResponse) {
    response = &DescribeVcubeResourcesListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVcubeResourcesList
// 查询视立方 license 资源，包括资源包赠送和直接购买的资源
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeResourcesList(request *DescribeVcubeResourcesListRequest) (response *DescribeVcubeResourcesListResponse, err error) {
    return c.DescribeVcubeResourcesListWithContext(context.Background(), request)
}

// DescribeVcubeResourcesList
// 查询视立方 license 资源，包括资源包赠送和直接购买的资源
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeVcubeResourcesListWithContext(ctx context.Context, request *DescribeVcubeResourcesListRequest) (response *DescribeVcubeResourcesListResponse, err error) {
    if request == nil {
        request = NewDescribeVcubeResourcesListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeVcubeResourcesList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVcubeResourcesList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVcubeResourcesListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeXMagicResourceRequest() (request *DescribeXMagicResourceRequest) {
    request = &DescribeXMagicResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeXMagicResource")
    
    
    return
}

func NewDescribeXMagicResourceResponse() (response *DescribeXMagicResourceResponse) {
    response = &DescribeXMagicResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeXMagicResource
// 用途美视资源包用于开通正式优图美视
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeXMagicResource(request *DescribeXMagicResourceRequest) (response *DescribeXMagicResourceResponse, err error) {
    return c.DescribeXMagicResourceWithContext(context.Background(), request)
}

// DescribeXMagicResource
// 用途美视资源包用于开通正式优图美视
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeXMagicResourceWithContext(ctx context.Context, request *DescribeXMagicResourceRequest) (response *DescribeXMagicResourceResponse, err error) {
    if request == nil {
        request = NewDescribeXMagicResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeXMagicResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeXMagicResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeXMagicResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeXMagicResourceListRequest() (request *DescribeXMagicResourceListRequest) {
    request = &DescribeXMagicResourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "DescribeXMagicResourceList")
    
    
    return
}

func NewDescribeXMagicResourceListResponse() (response *DescribeXMagicResourceListResponse) {
    response = &DescribeXMagicResourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeXMagicResourceList
// 用于优图美视资源列表展示
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeXMagicResourceList(request *DescribeXMagicResourceListRequest) (response *DescribeXMagicResourceListResponse, err error) {
    return c.DescribeXMagicResourceListWithContext(context.Background(), request)
}

// DescribeXMagicResourceList
// 用于优图美视资源列表展示
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeXMagicResourceListWithContext(ctx context.Context, request *DescribeXMagicResourceListRequest) (response *DescribeXMagicResourceListResponse, err error) {
    if request == nil {
        request = NewDescribeXMagicResourceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "DescribeXMagicResourceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeXMagicResourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeXMagicResourceListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplication
// 更改测试包名信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// 更改测试包名信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "ModifyApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFormalApplicationRequest() (request *ModifyFormalApplicationRequest) {
    request = &ModifyFormalApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "ModifyFormalApplication")
    
    
    return
}

func NewModifyFormalApplicationResponse() (response *ModifyFormalApplicationResponse) {
    response = &ModifyFormalApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFormalApplication
// 修改正式应用的包名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyFormalApplication(request *ModifyFormalApplicationRequest) (response *ModifyFormalApplicationResponse, err error) {
    return c.ModifyFormalApplicationWithContext(context.Background(), request)
}

// ModifyFormalApplication
// 修改正式应用的包名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyFormalApplicationWithContext(ctx context.Context, request *ModifyFormalApplicationRequest) (response *ModifyFormalApplicationResponse, err error) {
    if request == nil {
        request = NewModifyFormalApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "ModifyFormalApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFormalApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFormalApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLicenseRequest() (request *ModifyLicenseRequest) {
    request = &ModifyLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "ModifyLicense")
    
    
    return
}

func NewModifyLicenseResponse() (response *ModifyLicenseResponse) {
    response = &ModifyLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLicense
// 正式license 升降配，点播精简版、基础版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicense(request *ModifyLicenseRequest) (response *ModifyLicenseResponse, err error) {
    return c.ModifyLicenseWithContext(context.Background(), request)
}

// ModifyLicense
// 正式license 升降配，点播精简版、基础版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyLicenseWithContext(ctx context.Context, request *ModifyLicenseRequest) (response *ModifyLicenseResponse, err error) {
    if request == nil {
        request = NewModifyLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "ModifyLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPresetApplicationRequest() (request *ModifyPresetApplicationRequest) {
    request = &ModifyPresetApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "ModifyPresetApplication")
    
    
    return
}

func NewModifyPresetApplicationResponse() (response *ModifyPresetApplicationResponse) {
    response = &ModifyPresetApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPresetApplication
// 修改内置应用包名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPresetApplication(request *ModifyPresetApplicationRequest) (response *ModifyPresetApplicationResponse, err error) {
    return c.ModifyPresetApplicationWithContext(context.Background(), request)
}

// ModifyPresetApplication
// 修改内置应用包名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPresetApplicationWithContext(ctx context.Context, request *ModifyPresetApplicationRequest) (response *ModifyPresetApplicationResponse, err error) {
    if request == nil {
        request = NewModifyPresetApplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "ModifyPresetApplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPresetApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPresetApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTrialLicenseRequest() (request *ModifyTrialLicenseRequest) {
    request = &ModifyTrialLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "ModifyTrialLicense")
    
    
    return
}

func NewModifyTrialLicenseResponse() (response *ModifyTrialLicenseResponse) {
    response = &ModifyTrialLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTrialLicense
// 续期测试license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTrialLicense(request *ModifyTrialLicenseRequest) (response *ModifyTrialLicenseResponse, err error) {
    return c.ModifyTrialLicenseWithContext(context.Background(), request)
}

// ModifyTrialLicense
// 续期测试license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTrialLicenseWithContext(ctx context.Context, request *ModifyTrialLicenseRequest) (response *ModifyTrialLicenseResponse, err error) {
    if request == nil {
        request = NewModifyTrialLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "ModifyTrialLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTrialLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTrialLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyXMagicRequest() (request *ModifyXMagicRequest) {
    request = &ModifyXMagicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "ModifyXMagic")
    
    
    return
}

func NewModifyXMagicResponse() (response *ModifyXMagicResponse) {
    response = &ModifyXMagicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyXMagic
// 使用一个腾讯特效资源，更新现在的腾讯特效license，license功能和到期时间会以新的资源为准，老资源会被替换下来
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyXMagic(request *ModifyXMagicRequest) (response *ModifyXMagicResponse, err error) {
    return c.ModifyXMagicWithContext(context.Background(), request)
}

// ModifyXMagic
// 使用一个腾讯特效资源，更新现在的腾讯特效license，license功能和到期时间会以新的资源为准，老资源会被替换下来
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UPLOADLICENSEFILEFAIL = "FailedOperation.UploadLicenseFileFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyXMagicWithContext(ctx context.Context, request *ModifyXMagicRequest) (response *ModifyXMagicResponse, err error) {
    if request == nil {
        request = NewModifyXMagicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "ModifyXMagic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyXMagic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyXMagicResponse()
    err = c.Send(request, response)
    return
}

func NewRenewLicenseRequest() (request *RenewLicenseRequest) {
    request = &RenewLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "RenewLicense")
    
    
    return
}

func NewRenewLicenseResponse() (response *RenewLicenseResponse) {
    response = &RenewLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewLicense
// 正式license 续期与变更有效期
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RenewLicense(request *RenewLicenseRequest) (response *RenewLicenseResponse, err error) {
    return c.RenewLicenseWithContext(context.Background(), request)
}

// RenewLicense
// 正式license 续期与变更有效期
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RenewLicenseWithContext(ctx context.Context, request *RenewLicenseRequest) (response *RenewLicenseResponse, err error) {
    if request == nil {
        request = NewRenewLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "RenewLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewRenewTestXMagicRequest() (request *RenewTestXMagicRequest) {
    request = &RenewTestXMagicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "RenewTestXMagic")
    
    
    return
}

func NewRenewTestXMagicResponse() (response *RenewTestXMagicResponse) {
    response = &RenewTestXMagicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewTestXMagic
// 续期测试版优图美视
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RenewTestXMagic(request *RenewTestXMagicRequest) (response *RenewTestXMagicResponse, err error) {
    return c.RenewTestXMagicWithContext(context.Background(), request)
}

// RenewTestXMagic
// 续期测试版优图美视
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RenewTestXMagicWithContext(ctx context.Context, request *RenewTestXMagicRequest) (response *RenewTestXMagicResponse, err error) {
    if request == nil {
        request = NewRenewTestXMagicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "RenewTestXMagic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewTestXMagic require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewTestXMagicResponse()
    err = c.Send(request, response)
    return
}

func NewRenewVideoRequest() (request *RenewVideoRequest) {
    request = &RenewVideoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "RenewVideo")
    
    
    return
}

func NewRenewVideoResponse() (response *RenewVideoResponse) {
    response = &RenewVideoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewVideo
// 续期国际站视频播放功能和中国站web基础版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RenewVideo(request *RenewVideoRequest) (response *RenewVideoResponse, err error) {
    return c.RenewVideoWithContext(context.Background(), request)
}

// RenewVideo
// 续期国际站视频播放功能和中国站web基础版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RenewVideoWithContext(ctx context.Context, request *RenewVideoRequest) (response *RenewVideoResponse, err error) {
    if request == nil {
        request = NewRenewVideoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "RenewVideo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewVideo require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewVideoResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTestXMagicRequest() (request *UpdateTestXMagicRequest) {
    request = &UpdateTestXMagicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "UpdateTestXMagic")
    
    
    return
}

func NewUpdateTestXMagicResponse() (response *UpdateTestXMagicResponse) {
    response = &UpdateTestXMagicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTestXMagic
// 将测试xmagic升级到正式版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateTestXMagic(request *UpdateTestXMagicRequest) (response *UpdateTestXMagicResponse, err error) {
    return c.UpdateTestXMagicWithContext(context.Background(), request)
}

// UpdateTestXMagic
// 将测试xmagic升级到正式版
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateTestXMagicWithContext(ctx context.Context, request *UpdateTestXMagicRequest) (response *UpdateTestXMagicResponse, err error) {
    if request == nil {
        request = NewUpdateTestXMagicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "UpdateTestXMagic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTestXMagic require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTestXMagicResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateTrialLicenseRequest() (request *UpdateTrialLicenseRequest) {
    request = &UpdateTrialLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "UpdateTrialLicense")
    
    
    return
}

func NewUpdateTrialLicenseResponse() (response *UpdateTrialLicenseResponse) {
    response = &UpdateTrialLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateTrialLicense
// 测试 license 升级为正式 license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTrialLicense(request *UpdateTrialLicenseRequest) (response *UpdateTrialLicenseResponse, err error) {
    return c.UpdateTrialLicenseWithContext(context.Background(), request)
}

// UpdateTrialLicense
// 测试 license 升级为正式 license
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateTrialLicenseWithContext(ctx context.Context, request *UpdateTrialLicenseRequest) (response *UpdateTrialLicenseResponse, err error) {
    if request == nil {
        request = NewUpdateTrialLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "UpdateTrialLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateTrialLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateTrialLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateXMagicRequest() (request *UpdateXMagicRequest) {
    request = &UpdateXMagicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vcube", APIVersion, "UpdateXMagic")
    
    
    return
}

func NewUpdateXMagicResponse() (response *UpdateXMagicResponse) {
    response = &UpdateXMagicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateXMagic
// 更新优图美视的申请信息 Status 为2，3的时候可用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateXMagic(request *UpdateXMagicRequest) (response *UpdateXMagicResponse, err error) {
    return c.UpdateXMagicWithContext(context.Background(), request)
}

// UpdateXMagic
// 更新优图美视的申请信息 Status 为2，3的时候可用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateXMagicWithContext(ctx context.Context, request *UpdateXMagicRequest) (response *UpdateXMagicResponse, err error) {
    if request == nil {
        request = NewUpdateXMagicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vcube", APIVersion, "UpdateXMagic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateXMagic require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateXMagicResponse()
    err = c.Send(request, response)
    return
}
