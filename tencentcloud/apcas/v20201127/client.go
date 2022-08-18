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

package v20201127

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-27"

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


func NewGetTaskDetailRequest() (request *GetTaskDetailRequest) {
    request = &GetTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apcas", APIVersion, "GetTaskDetail")
    
    
    return
}

func NewGetTaskDetailResponse() (response *GetTaskDetailResponse) {
    response = &GetTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTaskDetail
// 查询画像洞察任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTaskDetail(request *GetTaskDetailRequest) (response *GetTaskDetailResponse, err error) {
    return c.GetTaskDetailWithContext(context.Background(), request)
}

// GetTaskDetail
// 查询画像洞察任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTaskDetailWithContext(ctx context.Context, request *GetTaskDetailRequest) (response *GetTaskDetailResponse, err error) {
    if request == nil {
        request = NewGetTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskListRequest() (request *GetTaskListRequest) {
    request = &GetTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apcas", APIVersion, "GetTaskList")
    
    
    return
}

func NewGetTaskListResponse() (response *GetTaskListResponse) {
    response = &GetTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTaskList
// 查询当前账号AppID下的画像洞察任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTaskList(request *GetTaskListRequest) (response *GetTaskListResponse, err error) {
    return c.GetTaskListWithContext(context.Background(), request)
}

// GetTaskList
// 查询当前账号AppID下的画像洞察任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTaskListWithContext(ctx context.Context, request *GetTaskListRequest) (response *GetTaskListResponse, err error) {
    if request == nil {
        request = NewGetTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewPredictRatingRequest() (request *PredictRatingRequest) {
    request = &PredictRatingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apcas", APIVersion, "PredictRating")
    
    
    return
}

func NewPredictRatingResponse() (response *PredictRatingResponse) {
    response = &PredictRatingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PredictRating
// 根据传入的设备号（IMEI、IDFA、手机号、手机号MD5），返回意向评级结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PredictRating(request *PredictRatingRequest) (response *PredictRatingResponse, err error) {
    return c.PredictRatingWithContext(context.Background(), request)
}

// PredictRating
// 根据传入的设备号（IMEI、IDFA、手机号、手机号MD5），返回意向评级结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PredictRatingWithContext(ctx context.Context, request *PredictRatingRequest) (response *PredictRatingResponse, err error) {
    if request == nil {
        request = NewPredictRatingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PredictRating require credential")
    }

    request.SetContext(ctx)
    
    response = NewPredictRatingResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCallDetailsRequest() (request *QueryCallDetailsRequest) {
    request = &QueryCallDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apcas", APIVersion, "QueryCallDetails")
    
    
    return
}

func NewQueryCallDetailsResponse() (response *QueryCallDetailsResponse) {
    response = &QueryCallDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCallDetails
// 查询调用明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCallDetails(request *QueryCallDetailsRequest) (response *QueryCallDetailsResponse, err error) {
    return c.QueryCallDetailsWithContext(context.Background(), request)
}

// QueryCallDetails
// 查询调用明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCallDetailsWithContext(ctx context.Context, request *QueryCallDetailsRequest) (response *QueryCallDetailsResponse, err error) {
    if request == nil {
        request = NewQueryCallDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCallDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCallDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCallStatRequest() (request *QueryCallStatRequest) {
    request = &QueryCallStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apcas", APIVersion, "QueryCallStat")
    
    
    return
}

func NewQueryCallStatResponse() (response *QueryCallStatResponse) {
    response = &QueryCallStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCallStat
// 按时间维度获取调用量统计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCallStat(request *QueryCallStatRequest) (response *QueryCallStatResponse, err error) {
    return c.QueryCallStatWithContext(context.Background(), request)
}

// QueryCallStat
// 按时间维度获取调用量统计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCallStatWithContext(ctx context.Context, request *QueryCallStatRequest) (response *QueryCallStatResponse, err error) {
    if request == nil {
        request = NewQueryCallStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCallStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCallStatResponse()
    err = c.Send(request, response)
    return
}

func NewQueryGeneralStatRequest() (request *QueryGeneralStatRequest) {
    request = &QueryGeneralStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apcas", APIVersion, "QueryGeneralStat")
    
    
    return
}

func NewQueryGeneralStatResponse() (response *QueryGeneralStatResponse) {
    response = &QueryGeneralStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryGeneralStat
// 获取日/月/周/总调用量统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryGeneralStat(request *QueryGeneralStatRequest) (response *QueryGeneralStatResponse, err error) {
    return c.QueryGeneralStatWithContext(context.Background(), request)
}

// QueryGeneralStat
// 获取日/月/周/总调用量统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryGeneralStatWithContext(ctx context.Context, request *QueryGeneralStatRequest) (response *QueryGeneralStatResponse, err error) {
    if request == nil {
        request = NewQueryGeneralStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryGeneralStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryGeneralStatResponse()
    err = c.Send(request, response)
    return
}

func NewUploadIdRequest() (request *UploadIdRequest) {
    request = &UploadIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apcas", APIVersion, "UploadId")
    
    
    return
}

func NewUploadIdResponse() (response *UploadIdResponse) {
    response = &UploadIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadId
// 上传群体画像的ID列表（支持的ID类型：0:imei 7:IDFA 8:MD5(imei)），后台返回生成的画像分析任务ID
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadId(request *UploadIdRequest) (response *UploadIdResponse, err error) {
    return c.UploadIdWithContext(context.Background(), request)
}

// UploadId
// 上传群体画像的ID列表（支持的ID类型：0:imei 7:IDFA 8:MD5(imei)），后台返回生成的画像分析任务ID
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALINVOKEFAILURE = "InternalError.InternalInvokeFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadIdWithContext(ctx context.Context, request *UploadIdRequest) (response *UploadIdResponse, err error) {
    if request == nil {
        request = NewUploadIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadId require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadIdResponse()
    err = c.Send(request, response)
    return
}
