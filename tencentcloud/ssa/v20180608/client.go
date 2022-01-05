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
    if request == nil {
        request = NewDescribeAssetDetailRequest()
    }
    
    response = NewDescribeAssetDetailResponse()
    err = c.Send(request, response)
    return
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
    request.SetContext(ctx)
    
    response = NewDescribeAssetDetailResponse()
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
    if request == nil {
        request = NewDescribeAssetListRequest()
    }
    
    response = NewDescribeAssetListResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewDescribeAssetsMappingListRequest()
    }
    
    response = NewDescribeAssetsMappingListResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeAssetsMappingListWithContext(ctx context.Context, request *DescribeAssetsMappingListRequest) (response *DescribeAssetsMappingListResponse, err error) {
    if request == nil {
        request = NewDescribeAssetsMappingListRequest()
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
    if request == nil {
        request = NewDescribeCheckConfigAssetListRequest()
    }
    
    response = NewDescribeCheckConfigAssetListResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeCheckConfigAssetListWithContext(ctx context.Context, request *DescribeCheckConfigAssetListRequest) (response *DescribeCheckConfigAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeCheckConfigAssetListRequest()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCheckConfigDetail(request *DescribeCheckConfigDetailRequest) (response *DescribeCheckConfigDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCheckConfigDetailRequest()
    }
    
    response = NewDescribeCheckConfigDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeCheckConfigDetail
// 云安全配置检查项详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeCheckConfigDetailWithContext(ctx context.Context, request *DescribeCheckConfigDetailRequest) (response *DescribeCheckConfigDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCheckConfigDetailRequest()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceAssetList(request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetListRequest()
    }
    
    response = NewDescribeComplianceAssetListResponse()
    err = c.Send(request, response)
    return
}

// DescribeComplianceAssetList
// 合规管理-资产列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceAssetListWithContext(ctx context.Context, request *DescribeComplianceAssetListRequest) (response *DescribeComplianceAssetListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceAssetListRequest()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceDetail(request *DescribeComplianceDetailRequest) (response *DescribeComplianceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceDetailRequest()
    }
    
    response = NewDescribeComplianceDetailResponse()
    err = c.Send(request, response)
    return
}

// DescribeComplianceDetail
// 合规管理检查项详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeComplianceDetailWithContext(ctx context.Context, request *DescribeComplianceDetailRequest) (response *DescribeComplianceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceDetailRequest()
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
    if request == nil {
        request = NewDescribeComplianceListRequest()
    }
    
    response = NewDescribeComplianceListResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeComplianceListWithContext(ctx context.Context, request *DescribeComplianceListRequest) (response *DescribeComplianceListResponse, err error) {
    if request == nil {
        request = NewDescribeComplianceListRequest()
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
    if request == nil {
        request = NewDescribeConfigListRequest()
    }
    
    response = NewDescribeConfigListResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewDescribeEventDetailRequest()
    }
    
    response = NewDescribeEventDetailResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewDescribeLeakDetectionListRequest()
    }
    
    response = NewDescribeLeakDetectionListResponse()
    err = c.Send(request, response)
    return
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
    request.SetContext(ctx)
    
    response = NewDescribeLeakDetectionListResponse()
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
    if request == nil {
        request = NewDescribeSafetyEventListRequest()
    }
    
    response = NewDescribeSafetyEventListResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewDescribeSocAlertDetailsRequest()
    }
    
    response = NewDescribeSocAlertDetailsResponse()
    err = c.Send(request, response)
    return
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
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocAlertList(request *DescribeSocAlertListRequest) (response *DescribeSocAlertListResponse, err error) {
    if request == nil {
        request = NewDescribeSocAlertListRequest()
    }
    
    response = NewDescribeSocAlertListResponse()
    err = c.Send(request, response)
    return
}

// DescribeSocAlertList
// 拉取告警列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
func (c *Client) DescribeSocAlertListWithContext(ctx context.Context, request *DescribeSocAlertListRequest) (response *DescribeSocAlertListResponse, err error) {
    if request == nil {
        request = NewDescribeSocAlertListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSocAlertListResponse()
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
    if request == nil {
        request = NewDescribeSocCspmComplianceRequest()
    }
    
    response = NewDescribeSocCspmComplianceResponse()
    err = c.Send(request, response)
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
func (c *Client) DescribeSocCspmComplianceWithContext(ctx context.Context, request *DescribeSocCspmComplianceRequest) (response *DescribeSocCspmComplianceResponse, err error) {
    if request == nil {
        request = NewDescribeSocCspmComplianceRequest()
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
    if request == nil {
        request = NewDescribeVulDetailRequest()
    }
    
    response = NewDescribeVulDetailResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewDescribeVulListRequest()
    }
    
    response = NewDescribeVulListResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewSaDivulgeDataQueryPubRequest()
    }
    
    response = NewSaDivulgeDataQueryPubResponse()
    err = c.Send(request, response)
    return
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
    request.SetContext(ctx)
    
    response = NewSaDivulgeDataQueryPubResponse()
    err = c.Send(request, response)
    return
}
