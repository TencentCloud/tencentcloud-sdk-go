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

package v20210811

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-08-11"

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


func NewDescribeKBComponentRequest() (request *DescribeKBComponentRequest) {
    request = &DescribeKBComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bsca", APIVersion, "DescribeKBComponent")
    
    
    return
}

func NewDescribeKBComponentResponse() (response *DescribeKBComponentResponse) {
    response = &DescribeKBComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKBComponent
// 本接口(DescribeKBComponent)用于在知识库中查询开源组件信息。本接口根据用户输入的PURL在知识库中寻找对应的开源组件，其中Name为必填字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeKBComponent(request *DescribeKBComponentRequest) (response *DescribeKBComponentResponse, err error) {
    return c.DescribeKBComponentWithContext(context.Background(), request)
}

// DescribeKBComponent
// 本接口(DescribeKBComponent)用于在知识库中查询开源组件信息。本接口根据用户输入的PURL在知识库中寻找对应的开源组件，其中Name为必填字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeKBComponentWithContext(ctx context.Context, request *DescribeKBComponentRequest) (response *DescribeKBComponentResponse, err error) {
    if request == nil {
        request = NewDescribeKBComponentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bsca", APIVersion, "DescribeKBComponent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKBComponent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKBComponentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKBComponentVersionListRequest() (request *DescribeKBComponentVersionListRequest) {
    request = &DescribeKBComponentVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bsca", APIVersion, "DescribeKBComponentVersionList")
    
    
    return
}

func NewDescribeKBComponentVersionListResponse() (response *DescribeKBComponentVersionListResponse) {
    response = &DescribeKBComponentVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKBComponentVersionList
// 查询特定组件的版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeKBComponentVersionList(request *DescribeKBComponentVersionListRequest) (response *DescribeKBComponentVersionListResponse, err error) {
    return c.DescribeKBComponentVersionListWithContext(context.Background(), request)
}

// DescribeKBComponentVersionList
// 查询特定组件的版本列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeKBComponentVersionListWithContext(ctx context.Context, request *DescribeKBComponentVersionListRequest) (response *DescribeKBComponentVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeKBComponentVersionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bsca", APIVersion, "DescribeKBComponentVersionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKBComponentVersionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKBComponentVersionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKBComponentVulnerabilityRequest() (request *DescribeKBComponentVulnerabilityRequest) {
    request = &DescribeKBComponentVulnerabilityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bsca", APIVersion, "DescribeKBComponentVulnerability")
    
    
    return
}

func NewDescribeKBComponentVulnerabilityResponse() (response *DescribeKBComponentVulnerabilityResponse) {
    response = &DescribeKBComponentVulnerabilityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKBComponentVulnerability
// 本接口(DescribeKBComponentVulnerability)用于在知识库中查询开源组件的漏洞信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeKBComponentVulnerability(request *DescribeKBComponentVulnerabilityRequest) (response *DescribeKBComponentVulnerabilityResponse, err error) {
    return c.DescribeKBComponentVulnerabilityWithContext(context.Background(), request)
}

// DescribeKBComponentVulnerability
// 本接口(DescribeKBComponentVulnerability)用于在知识库中查询开源组件的漏洞信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeKBComponentVulnerabilityWithContext(ctx context.Context, request *DescribeKBComponentVulnerabilityRequest) (response *DescribeKBComponentVulnerabilityResponse, err error) {
    if request == nil {
        request = NewDescribeKBComponentVulnerabilityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bsca", APIVersion, "DescribeKBComponentVulnerability")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKBComponentVulnerability require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKBComponentVulnerabilityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKBLicenseRequest() (request *DescribeKBLicenseRequest) {
    request = &DescribeKBLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bsca", APIVersion, "DescribeKBLicense")
    
    
    return
}

func NewDescribeKBLicenseResponse() (response *DescribeKBLicenseResponse) {
    response = &DescribeKBLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKBLicense
// 本接口(DescribeKBLicense)用于在知识库中查询许可证信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeKBLicense(request *DescribeKBLicenseRequest) (response *DescribeKBLicenseResponse, err error) {
    return c.DescribeKBLicenseWithContext(context.Background(), request)
}

// DescribeKBLicense
// 本接口(DescribeKBLicense)用于在知识库中查询许可证信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeKBLicenseWithContext(ctx context.Context, request *DescribeKBLicenseRequest) (response *DescribeKBLicenseResponse, err error) {
    if request == nil {
        request = NewDescribeKBLicenseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bsca", APIVersion, "DescribeKBLicense")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKBLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKBLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKBVulnerabilityRequest() (request *DescribeKBVulnerabilityRequest) {
    request = &DescribeKBVulnerabilityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bsca", APIVersion, "DescribeKBVulnerability")
    
    
    return
}

func NewDescribeKBVulnerabilityResponse() (response *DescribeKBVulnerabilityResponse) {
    response = &DescribeKBVulnerabilityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKBVulnerability
// 本接口(DescribeKBVulnerability)用于在知识库中查询漏洞详细信息，支持根据CVE、Vul ID、CNVD ID、CNNVD ID查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeKBVulnerability(request *DescribeKBVulnerabilityRequest) (response *DescribeKBVulnerabilityResponse, err error) {
    return c.DescribeKBVulnerabilityWithContext(context.Background(), request)
}

// DescribeKBVulnerability
// 本接口(DescribeKBVulnerability)用于在知识库中查询漏洞详细信息，支持根据CVE、Vul ID、CNVD ID、CNNVD ID查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeKBVulnerabilityWithContext(ctx context.Context, request *DescribeKBVulnerabilityRequest) (response *DescribeKBVulnerabilityResponse, err error) {
    if request == nil {
        request = NewDescribeKBVulnerabilityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bsca", APIVersion, "DescribeKBVulnerability")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKBVulnerability require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKBVulnerabilityResponse()
    err = c.Send(request, response)
    return
}

func NewMatchKBPURLListRequest() (request *MatchKBPURLListRequest) {
    request = &MatchKBPURLListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bsca", APIVersion, "MatchKBPURLList")
    
    
    return
}

func NewMatchKBPURLListResponse() (response *MatchKBPURLListResponse) {
    response = &MatchKBPURLListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MatchKBPURLList
// 本接口(MatchKBPURLList)用于在知识库中匹配与特征对应的开源组件列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) MatchKBPURLList(request *MatchKBPURLListRequest) (response *MatchKBPURLListResponse, err error) {
    return c.MatchKBPURLListWithContext(context.Background(), request)
}

// MatchKBPURLList
// 本接口(MatchKBPURLList)用于在知识库中匹配与特征对应的开源组件列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) MatchKBPURLListWithContext(ctx context.Context, request *MatchKBPURLListRequest) (response *MatchKBPURLListResponse, err error) {
    if request == nil {
        request = NewMatchKBPURLListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bsca", APIVersion, "MatchKBPURLList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("MatchKBPURLList require credential")
    }

    request.SetContext(ctx)
    
    response = NewMatchKBPURLListResponse()
    err = c.Send(request, response)
    return
}

func NewSearchKBComponentRequest() (request *SearchKBComponentRequest) {
    request = &SearchKBComponentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bsca", APIVersion, "SearchKBComponent")
    
    
    return
}

func NewSearchKBComponentResponse() (response *SearchKBComponentResponse) {
    response = &SearchKBComponentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchKBComponent
// 根据输入的组件名、组件类型搜索相应的组件，返回符合条件的组件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SearchKBComponent(request *SearchKBComponentRequest) (response *SearchKBComponentResponse, err error) {
    return c.SearchKBComponentWithContext(context.Background(), request)
}

// SearchKBComponent
// 根据输入的组件名、组件类型搜索相应的组件，返回符合条件的组件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTENOUGH = "FailedOperation.AccountNotEnough"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) SearchKBComponentWithContext(ctx context.Context, request *SearchKBComponentRequest) (response *SearchKBComponentResponse, err error) {
    if request == nil {
        request = NewSearchKBComponentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bsca", APIVersion, "SearchKBComponent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchKBComponent require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchKBComponentResponse()
    err = c.Send(request, response)
    return
}
