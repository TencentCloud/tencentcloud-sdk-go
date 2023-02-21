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

package v20180608

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-08"

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


func NewDescribeAssetDetailRequest() (request *DescribeAssetDetailRequest) {
    request = &DescribeAssetDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeAssetDetail")
    
    
    return
}

func NewDescribeAssetDetailResponse() (response *DescribeAssetDetailResponse) {
    response = &DescribeAssetDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetDetail
// 资产安全页资产详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetDetail(request *DescribeAssetDetailRequest) (response *DescribeAssetDetailResponse, err error) {
    return c.DescribeAssetDetailWithContext(context.Background(), request)
}

// DescribeAssetDetail
// 资产安全页资产详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetDetailWithContext(ctx context.Context, request *DescribeAssetDetailRequest) (response *DescribeAssetDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDetailListRequest() (request *DescribeAssetDetailListRequest) {
    request = &DescribeAssetDetailListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeAssetDetailList")
    
    
    return
}

func NewDescribeAssetDetailListResponse() (response *DescribeAssetDetailListResponse) {
    response = &DescribeAssetDetailListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetDetailList
// 资产条件查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetDetailList(request *DescribeAssetDetailListRequest) (response *DescribeAssetDetailListResponse, err error) {
    return c.DescribeAssetDetailListWithContext(context.Background(), request)
}

// DescribeAssetDetailList
// 资产条件查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetDetailListWithContext(ctx context.Context, request *DescribeAssetDetailListRequest) (response *DescribeAssetDetailListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDetailListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDetailList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDetailListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetListRequest() (request *DescribeAssetListRequest) {
    request = &DescribeAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeAssetList")
    
    
    return
}

func NewDescribeAssetListResponse() (response *DescribeAssetListResponse) {
    response = &DescribeAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetList
// 资产安全资产列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetList(request *DescribeAssetListRequest) (response *DescribeAssetListResponse, err error) {
    return c.DescribeAssetListWithContext(context.Background(), request)
}

// DescribeAssetList
// 资产安全资产列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAssetListWithContext(ctx context.Context, request *DescribeAssetListRequest) (response *DescribeAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetsMappingListRequest() (request *DescribeAssetsMappingListRequest) {
    request = &DescribeAssetsMappingListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeAssetsMappingList")
    
    
    return
}

func NewDescribeAssetsMappingListResponse() (response *DescribeAssetsMappingListResponse) {
    response = &DescribeAssetsMappingListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetsMappingList
// 资产测绘-测绘列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeAssetsMappingList(request *DescribeAssetsMappingListRequest) (response *DescribeAssetsMappingListResponse, err error) {
    return c.DescribeAssetsMappingListWithContext(context.Background(), request)
}

// DescribeAssetsMappingList
// 资产测绘-测绘列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeAssetsMappingListWithContext(ctx context.Context, request *DescribeAssetsMappingListRequest) (response *DescribeAssetsMappingListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetsMappingListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetsMappingList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetsMappingListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckConfigAssetListRequest() (request *DescribeCheckConfigAssetListRequest) {
    request = &DescribeCheckConfigAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeCheckConfigAssetList")
    
    
    return
}

func NewDescribeCheckConfigAssetListResponse() (response *DescribeCheckConfigAssetListResponse) {
    response = &DescribeCheckConfigAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCheckConfigAssetList
// 云安全配置管理资产组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCheckConfigAssetList(request *DescribeCheckConfigAssetListRequest) (response *DescribeCheckConfigAssetListResponse, err error) {
    return c.DescribeCheckConfigAssetListWithContext(context.Background(), request)
}

// DescribeCheckConfigAssetList
// 云安全配置管理资产组列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCheckConfigAssetListWithContext(ctx context.Context, request *DescribeCheckConfigAssetListRequest) (response *DescribeCheckConfigAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeCheckConfigAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckConfigAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckConfigAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckConfigDetailRequest() (request *DescribeCheckConfigDetailRequest) {
    request = &DescribeCheckConfigDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeCheckConfigDetail")
    
    
    return
}

func NewDescribeCheckConfigDetailResponse() (response *DescribeCheckConfigDetailResponse) {
    response = &DescribeCheckConfigDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCheckConfigDetail
// 云安全配置检查项详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCheckConfigDetail(request *DescribeCheckConfigDetailRequest) (response *DescribeCheckConfigDetailResponse, err error) {
    return c.DescribeCheckConfigDetailWithContext(context.Background(), request)
}

// DescribeCheckConfigDetail
// 云安全配置检查项详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCheckConfigDetailWithContext(ctx context.Context, request *DescribeCheckConfigDetailRequest) (response *DescribeCheckConfigDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCheckConfigDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckConfigDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckConfigDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceAssetListRequest() (request *DescribeComplianceAssetListRequest) {
    request = &DescribeComplianceAssetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeComplianceAssetList")
    
    
    return
}

func NewDescribeComplianceAssetListResponse() (response *DescribeComplianceAssetListResponse) {
    response = &DescribeComplianceAssetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceAssetList
// 合规管理-资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceAssetList(request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    return c.DescribeComplianceAssetListWithContext(context.Background(), request)
}

// DescribeComplianceAssetList
// 合规管理-资产列表
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceAssetListWithContext(ctx context.Context, request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceAssetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceAssetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceDetailRequest() (request *DescribeComplianceDetailRequest) {
    request = &DescribeComplianceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeComplianceDetail")
    
    
    return
}

func NewDescribeComplianceDetailResponse() (response *DescribeComplianceDetailResponse) {
    response = &DescribeComplianceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceDetail
// 合规管理检查项详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceDetail(request *DescribeComplianceDetailRequest) (response *DescribeComplianceDetailResponse, err error) {
    return c.DescribeComplianceDetailWithContext(context.Background(), request)
}

// DescribeComplianceDetail
// 合规管理检查项详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceDetailWithContext(ctx context.Context, request *DescribeComplianceDetailRequest) (response *DescribeComplianceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComplianceListRequest() (request *DescribeComplianceListRequest) {
    request = &DescribeComplianceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeComplianceList")
    
    
    return
}

func NewDescribeComplianceListResponse() (response *DescribeComplianceListResponse) {
    response = &DescribeComplianceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComplianceList
// 合规管理总览页检查项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceList(request *DescribeComplianceListRequest) (response *DescribeComplianceListResponse, err error) {
    return c.DescribeComplianceListWithContext(context.Background(), request)
}

// DescribeComplianceList
// 合规管理总览页检查项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceListWithContext(ctx context.Context, request *DescribeComplianceListRequest) (response *DescribeComplianceListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComplianceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComplianceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigListRequest() (request *DescribeConfigListRequest) {
    request = &DescribeConfigListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeConfigList")
    
    
    return
}

func NewDescribeConfigListResponse() (response *DescribeConfigListResponse) {
    response = &DescribeConfigListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConfigList
// 云配置检查项总览页检查项列表
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConfigList(request *DescribeConfigListRequest) (response *DescribeConfigListResponse, err error) {
    return c.DescribeConfigListWithContext(context.Background(), request)
}

// DescribeConfigList
// 云配置检查项总览页检查项列表
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeConfigListWithContext(ctx context.Context, request *DescribeConfigListRequest) (response *DescribeConfigListResponse, err error) {
    if request == nil {
        request = NewDescribeConfigListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventDetailRequest() (request *DescribeEventDetailRequest) {
    request = &DescribeEventDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeEventDetail")
    
    
    return
}

func NewDescribeEventDetailResponse() (response *DescribeEventDetailResponse) {
    response = &DescribeEventDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEventDetail
// 获取安全事件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventDetail(request *DescribeEventDetailRequest) (response *DescribeEventDetailResponse, err error) {
    return c.DescribeEventDetailWithContext(context.Background(), request)
}

// DescribeEventDetail
// 获取安全事件详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventDetailWithContext(ctx context.Context, request *DescribeEventDetailRequest) (response *DescribeEventDetailResponse, err error) {
    if request == nil {
        request = NewDescribeEventDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLeakDetectionListRequest() (request *DescribeLeakDetectionListRequest) {
    request = &DescribeLeakDetectionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeLeakDetectionList")
    
    
    return
}

func NewDescribeLeakDetectionListResponse() (response *DescribeLeakDetectionListResponse) {
    response = &DescribeLeakDetectionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLeakDetectionList
// 获取泄露列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLeakDetectionList(request *DescribeLeakDetectionListRequest) (response *DescribeLeakDetectionListResponse, err error) {
    return c.DescribeLeakDetectionListWithContext(context.Background(), request)
}

// DescribeLeakDetectionList
// 获取泄露列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLeakDetectionListWithContext(ctx context.Context, request *DescribeLeakDetectionListRequest) (response *DescribeLeakDetectionListResponse, err error) {
    if request == nil {
        request = NewDescribeLeakDetectionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLeakDetectionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLeakDetectionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMappingResultsRequest() (request *DescribeMappingResultsRequest) {
    request = &DescribeMappingResultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeMappingResults")
    
    
    return
}

func NewDescribeMappingResultsResponse() (response *DescribeMappingResultsResponse) {
    response = &DescribeMappingResultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMappingResults
// 获取测绘列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeMappingResults(request *DescribeMappingResultsRequest) (response *DescribeMappingResultsResponse, err error) {
    return c.DescribeMappingResultsWithContext(context.Background(), request)
}

// DescribeMappingResults
// 获取测绘列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeMappingResultsWithContext(ctx context.Context, request *DescribeMappingResultsRequest) (response *DescribeMappingResultsResponse, err error) {
    if request == nil {
        request = NewDescribeMappingResultsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMappingResults require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMappingResultsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafetyEventListRequest() (request *DescribeSafetyEventListRequest) {
    request = &DescribeSafetyEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeSafetyEventList")
    
    
    return
}

func NewDescribeSafetyEventListResponse() (response *DescribeSafetyEventListResponse) {
    response = &DescribeSafetyEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSafetyEventList
// 获取安全事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSafetyEventList(request *DescribeSafetyEventListRequest) (response *DescribeSafetyEventListResponse, err error) {
    return c.DescribeSafetyEventListWithContext(context.Background(), request)
}

// DescribeSafetyEventList
// 获取安全事件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSafetyEventListWithContext(ctx context.Context, request *DescribeSafetyEventListRequest) (response *DescribeSafetyEventListResponse, err error) {
    if request == nil {
        request = NewDescribeSafetyEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSafetyEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSafetyEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSocAlertDetailsRequest() (request *DescribeSocAlertDetailsRequest) {
    request = &DescribeSocAlertDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeSocAlertDetails")
    
    
    return
}

func NewDescribeSocAlertDetailsResponse() (response *DescribeSocAlertDetailsResponse) {
    response = &DescribeSocAlertDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSocAlertDetails
// 返回告警详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocAlertDetails(request *DescribeSocAlertDetailsRequest) (response *DescribeSocAlertDetailsResponse, err error) {
    return c.DescribeSocAlertDetailsWithContext(context.Background(), request)
}

// DescribeSocAlertDetails
// 返回告警详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocAlertDetailsWithContext(ctx context.Context, request *DescribeSocAlertDetailsRequest) (response *DescribeSocAlertDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeSocAlertDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSocAlertDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSocAlertDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSocAlertListRequest() (request *DescribeSocAlertListRequest) {
    request = &DescribeSocAlertListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeSocAlertList")
    
    
    return
}

func NewDescribeSocAlertListResponse() (response *DescribeSocAlertListResponse) {
    response = &DescribeSocAlertListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSocAlertList
// 拉取告警列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeSocAlertList(request *DescribeSocAlertListRequest) (response *DescribeSocAlertListResponse, err error) {
    return c.DescribeSocAlertListWithContext(context.Background(), request)
}

// DescribeSocAlertList
// 拉取告警列表
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DescribeSocAlertListWithContext(ctx context.Context, request *DescribeSocAlertListRequest) (response *DescribeSocAlertListResponse, err error) {
    if request == nil {
        request = NewDescribeSocAlertListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSocAlertList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSocAlertListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSocCheckItemListRequest() (request *DescribeSocCheckItemListRequest) {
    request = &DescribeSocCheckItemListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeSocCheckItemList")
    
    
    return
}

func NewDescribeSocCheckItemListResponse() (response *DescribeSocCheckItemListResponse) {
    response = &DescribeSocCheckItemListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSocCheckItemList
// 云安全配置检查项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocCheckItemList(request *DescribeSocCheckItemListRequest) (response *DescribeSocCheckItemListResponse, err error) {
    return c.DescribeSocCheckItemListWithContext(context.Background(), request)
}

// DescribeSocCheckItemList
// 云安全配置检查项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocCheckItemListWithContext(ctx context.Context, request *DescribeSocCheckItemListRequest) (response *DescribeSocCheckItemListResponse, err error) {
    if request == nil {
        request = NewDescribeSocCheckItemListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSocCheckItemList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSocCheckItemListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSocCheckResultListRequest() (request *DescribeSocCheckResultListRequest) {
    request = &DescribeSocCheckResultListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeSocCheckResultList")
    
    
    return
}

func NewDescribeSocCheckResultListResponse() (response *DescribeSocCheckResultListResponse) {
    response = &DescribeSocCheckResultListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSocCheckResultList
// 云安全配置检查项结果列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocCheckResultList(request *DescribeSocCheckResultListRequest) (response *DescribeSocCheckResultListResponse, err error) {
    return c.DescribeSocCheckResultListWithContext(context.Background(), request)
}

// DescribeSocCheckResultList
// 云安全配置检查项结果列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocCheckResultListWithContext(ctx context.Context, request *DescribeSocCheckResultListRequest) (response *DescribeSocCheckResultListResponse, err error) {
    if request == nil {
        request = NewDescribeSocCheckResultListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSocCheckResultList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSocCheckResultListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSocCspmComplianceRequest() (request *DescribeSocCspmComplianceRequest) {
    request = &DescribeSocCspmComplianceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeSocCspmCompliance")
    
    
    return
}

func NewDescribeSocCspmComplianceResponse() (response *DescribeSocCspmComplianceResponse) {
    response = &DescribeSocCspmComplianceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSocCspmCompliance
// 合规详情项
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
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
//  UNAUTHORIZEDOPERATION_SUBACCOUNTUNAUTHORIZED = "UnauthorizedOperation.SubAccountUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSocCspmCompliance(request *DescribeSocCspmComplianceRequest) (response *DescribeSocCspmComplianceResponse, err error) {
    return c.DescribeSocCspmComplianceWithContext(context.Background(), request)
}

// DescribeSocCspmCompliance
// 合规详情项
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHMODULEFAILED = "AuthFailure.AuthModuleFailed"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
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
//  UNAUTHORIZEDOPERATION_SUBACCOUNTUNAUTHORIZED = "UnauthorizedOperation.SubAccountUnauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSocCspmComplianceWithContext(ctx context.Context, request *DescribeSocCspmComplianceRequest) (response *DescribeSocCspmComplianceResponse, err error) {
    if request == nil {
        request = NewDescribeSocCspmComplianceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSocCspmCompliance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSocCspmComplianceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulDetailRequest() (request *DescribeVulDetailRequest) {
    request = &DescribeVulDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeVulDetail")
    
    
    return
}

func NewDescribeVulDetailResponse() (response *DescribeVulDetailResponse) {
    response = &DescribeVulDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulDetail
// 漏洞列表页，获取漏洞详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeVulDetail(request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
    return c.DescribeVulDetailWithContext(context.Background(), request)
}

// DescribeVulDetail
// 漏洞列表页，获取漏洞详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeVulDetailWithContext(ctx context.Context, request *DescribeVulDetailRequest) (response *DescribeVulDetailResponse, err error) {
    if request == nil {
        request = NewDescribeVulDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVulListRequest() (request *DescribeVulListRequest) {
    request = &DescribeVulListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "DescribeVulList")
    
    
    return
}

func NewDescribeVulListResponse() (response *DescribeVulListResponse) {
    response = &DescribeVulListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVulList
// 漏洞管理页，获取漏洞列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVulList(request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
    return c.DescribeVulListWithContext(context.Background(), request)
}

// DescribeVulList
// 漏洞管理页，获取漏洞列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeVulListWithContext(ctx context.Context, request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
    if request == nil {
        request = NewDescribeVulListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVulList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVulListResponse()
    err = c.Send(request, response)
    return
}

func NewSaDivulgeDataQueryPubRequest() (request *SaDivulgeDataQueryPubRequest) {
    request = &SaDivulgeDataQueryPubRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "SaDivulgeDataQueryPub")
    
    
    return
}

func NewSaDivulgeDataQueryPubResponse() (response *SaDivulgeDataQueryPubResponse) {
    response = &SaDivulgeDataQueryPubResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SaDivulgeDataQueryPub
// 查询【通用字段】【泄露监测数据列表】
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) SaDivulgeDataQueryPub(request *SaDivulgeDataQueryPubRequest) (response *SaDivulgeDataQueryPubResponse, err error) {
    return c.SaDivulgeDataQueryPubWithContext(context.Background(), request)
}

// SaDivulgeDataQueryPub
// 查询【通用字段】【泄露监测数据列表】
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) SaDivulgeDataQueryPubWithContext(ctx context.Context, request *SaDivulgeDataQueryPubRequest) (response *SaDivulgeDataQueryPubResponse, err error) {
    if request == nil {
        request = NewSaDivulgeDataQueryPubRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaDivulgeDataQueryPub require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaDivulgeDataQueryPubResponse()
    err = c.Send(request, response)
    return
}

func NewSaEventPubRequest() (request *SaEventPubRequest) {
    request = &SaEventPubRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssa", APIVersion, "SaEventPub")
    
    
    return
}

func NewSaEventPubResponse() (response *SaEventPubResponse) {
    response = &SaEventPubResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SaEventPub
// 安全事件通用字段
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) SaEventPub(request *SaEventPubRequest) (response *SaEventPubResponse, err error) {
    return c.SaEventPubWithContext(context.Background(), request)
}

// SaEventPub
// 安全事件通用字段
//
// 可能返回的错误码:
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) SaEventPubWithContext(ctx context.Context, request *SaEventPubRequest) (response *SaEventPubResponse, err error) {
    if request == nil {
        request = NewSaEventPubRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaEventPub require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaEventPubResponse()
    err = c.Send(request, response)
    return
}
