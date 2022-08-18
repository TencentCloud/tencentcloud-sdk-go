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

package v20180228

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-02-28"

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


func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccount
// 创建集团门店管理员账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// 创建集团门店管理员账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFacePictureRequest() (request *CreateFacePictureRequest) {
    request = &CreateFacePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "CreateFacePicture")
    
    
    return
}

func NewCreateFacePictureResponse() (response *CreateFacePictureResponse) {
    response = &CreateFacePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFacePicture
// 通过上传指定规格的人脸图片，创建黑名单用户或者白名单用户。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFacePicture(request *CreateFacePictureRequest) (response *CreateFacePictureResponse, err error) {
    return c.CreateFacePictureWithContext(context.Background(), request)
}

// CreateFacePicture
// 通过上传指定规格的人脸图片，创建黑名单用户或者白名单用户。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFacePictureWithContext(ctx context.Context, request *CreateFacePictureRequest) (response *CreateFacePictureResponse, err error) {
    if request == nil {
        request = NewCreateFacePictureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFacePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFacePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonFeatureRequest() (request *DeletePersonFeatureRequest) {
    request = &DeletePersonFeatureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DeletePersonFeature")
    
    
    return
}

func NewDeletePersonFeatureResponse() (response *DeletePersonFeatureResponse) {
    response = &DeletePersonFeatureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePersonFeature
// 删除顾客特征，仅支持删除黑名单或者白名单用户特征。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NOPERSON = "FailedOperation.NoPerson"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePersonFeature(request *DeletePersonFeatureRequest) (response *DeletePersonFeatureResponse, err error) {
    return c.DeletePersonFeatureWithContext(context.Background(), request)
}

// DeletePersonFeature
// 删除顾客特征，仅支持删除黑名单或者白名单用户特征。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NOPERSON = "FailedOperation.NoPerson"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePersonFeatureWithContext(ctx context.Context, request *DeletePersonFeatureRequest) (response *DeletePersonFeatureResponse, err error) {
    if request == nil {
        request = NewDeletePersonFeatureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePersonFeature require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePersonFeatureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCameraPersonRequest() (request *DescribeCameraPersonRequest) {
    request = &DescribeCameraPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeCameraPerson")
    
    
    return
}

func NewDescribeCameraPersonResponse() (response *DescribeCameraPersonResponse) {
    response = &DescribeCameraPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCameraPerson
// 通过指定设备ID和指定时段，获取该时段内中收银台摄像设备抓取到顾客头像及身份ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
func (c *Client) DescribeCameraPerson(request *DescribeCameraPersonRequest) (response *DescribeCameraPersonResponse, err error) {
    return c.DescribeCameraPersonWithContext(context.Background(), request)
}

// DescribeCameraPerson
// 通过指定设备ID和指定时段，获取该时段内中收银台摄像设备抓取到顾客头像及身份ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
func (c *Client) DescribeCameraPersonWithContext(ctx context.Context, request *DescribeCameraPersonRequest) (response *DescribeCameraPersonResponse, err error) {
    if request == nil {
        request = NewDescribeCameraPersonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCameraPerson require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCameraPersonResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterPersonArrivedMallRequest() (request *DescribeClusterPersonArrivedMallRequest) {
    request = &DescribeClusterPersonArrivedMallRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeClusterPersonArrivedMall")
    
    
    return
}

func NewDescribeClusterPersonArrivedMallResponse() (response *DescribeClusterPersonArrivedMallResponse) {
    response = &DescribeClusterPersonArrivedMallResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterPersonArrivedMall
// 输出开始时间到结束时间段内的进出场数据。按天聚合的情况下，每天多次进出场算一次，以最初进场时间为进场时间，最后离场时间为离场时间。停留时间为多次进出场的停留时间之和。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterPersonArrivedMall(request *DescribeClusterPersonArrivedMallRequest) (response *DescribeClusterPersonArrivedMallResponse, err error) {
    return c.DescribeClusterPersonArrivedMallWithContext(context.Background(), request)
}

// DescribeClusterPersonArrivedMall
// 输出开始时间到结束时间段内的进出场数据。按天聚合的情况下，每天多次进出场算一次，以最初进场时间为进场时间，最后离场时间为离场时间。停留时间为多次进出场的停留时间之和。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterPersonArrivedMallWithContext(ctx context.Context, request *DescribeClusterPersonArrivedMallRequest) (response *DescribeClusterPersonArrivedMallResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPersonArrivedMallRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPersonArrivedMall require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPersonArrivedMallResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterPersonTraceRequest() (request *DescribeClusterPersonTraceRequest) {
    request = &DescribeClusterPersonTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeClusterPersonTrace")
    
    
    return
}

func NewDescribeClusterPersonTraceResponse() (response *DescribeClusterPersonTraceResponse) {
    response = &DescribeClusterPersonTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterPersonTrace
// 输出开始时间到结束时间段内的进出场数据。按天聚合的情况下，每天多次进出场算一次，以最初进场时间为进场时间，最后离场时间为离场时间。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterPersonTrace(request *DescribeClusterPersonTraceRequest) (response *DescribeClusterPersonTraceResponse, err error) {
    return c.DescribeClusterPersonTraceWithContext(context.Background(), request)
}

// DescribeClusterPersonTrace
// 输出开始时间到结束时间段内的进出场数据。按天聚合的情况下，每天多次进出场算一次，以最初进场时间为进场时间，最后离场时间为离场时间。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterPersonTraceWithContext(ctx context.Context, request *DescribeClusterPersonTraceRequest) (response *DescribeClusterPersonTraceResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPersonTraceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPersonTrace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPersonTraceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFaceIdByTempIdRequest() (request *DescribeFaceIdByTempIdRequest) {
    request = &DescribeFaceIdByTempIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeFaceIdByTempId")
    
    
    return
}

func NewDescribeFaceIdByTempIdResponse() (response *DescribeFaceIdByTempIdResponse) {
    response = &DescribeFaceIdByTempIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFaceIdByTempId
// 通过DescribeCameraPerson接口上报的收银台身份ID查询顾客的FaceID。查询最佳时间为收银台上报的次日1点后。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFaceIdByTempId(request *DescribeFaceIdByTempIdRequest) (response *DescribeFaceIdByTempIdResponse, err error) {
    return c.DescribeFaceIdByTempIdWithContext(context.Background(), request)
}

// DescribeFaceIdByTempId
// 通过DescribeCameraPerson接口上报的收银台身份ID查询顾客的FaceID。查询最佳时间为收银台上报的次日1点后。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFaceIdByTempIdWithContext(ctx context.Context, request *DescribeFaceIdByTempIdRequest) (response *DescribeFaceIdByTempIdResponse, err error) {
    if request == nil {
        request = NewDescribeFaceIdByTempIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFaceIdByTempId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFaceIdByTempIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHistoryNetworkInfoRequest() (request *DescribeHistoryNetworkInfoRequest) {
    request = &DescribeHistoryNetworkInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeHistoryNetworkInfo")
    
    
    return
}

func NewDescribeHistoryNetworkInfoResponse() (response *DescribeHistoryNetworkInfoResponse) {
    response = &DescribeHistoryNetworkInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHistoryNetworkInfo
// 返回当前门店历史网络状态数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHistoryNetworkInfo(request *DescribeHistoryNetworkInfoRequest) (response *DescribeHistoryNetworkInfoResponse, err error) {
    return c.DescribeHistoryNetworkInfoWithContext(context.Background(), request)
}

// DescribeHistoryNetworkInfo
// 返回当前门店历史网络状态数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeHistoryNetworkInfoWithContext(ctx context.Context, request *DescribeHistoryNetworkInfoRequest) (response *DescribeHistoryNetworkInfoResponse, err error) {
    if request == nil {
        request = NewDescribeHistoryNetworkInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHistoryNetworkInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHistoryNetworkInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkInfoRequest() (request *DescribeNetworkInfoRequest) {
    request = &DescribeNetworkInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeNetworkInfo")
    
    
    return
}

func NewDescribeNetworkInfoResponse() (response *DescribeNetworkInfoResponse) {
    response = &DescribeNetworkInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNetworkInfo
// 返回当前门店最新网络状态数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNetworkInfo(request *DescribeNetworkInfoRequest) (response *DescribeNetworkInfoResponse, err error) {
    return c.DescribeNetworkInfoWithContext(context.Background(), request)
}

// DescribeNetworkInfo
// 返回当前门店最新网络状态数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeNetworkInfoWithContext(ctx context.Context, request *DescribeNetworkInfoRequest) (response *DescribeNetworkInfoResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNetworkInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNetworkInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonRequest() (request *DescribePersonRequest) {
    request = &DescribePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePerson")
    
    
    return
}

func NewDescribePersonResponse() (response *DescribePersonResponse) {
    response = &DescribePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePerson
// 查询指定某一卖场的用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePerson(request *DescribePersonRequest) (response *DescribePersonResponse, err error) {
    return c.DescribePersonWithContext(context.Background(), request)
}

// DescribePerson
// 查询指定某一卖场的用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonWithContext(ctx context.Context, request *DescribePersonRequest) (response *DescribePersonResponse, err error) {
    if request == nil {
        request = NewDescribePersonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePerson require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonArrivedMallRequest() (request *DescribePersonArrivedMallRequest) {
    request = &DescribePersonArrivedMallRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonArrivedMall")
    
    
    return
}

func NewDescribePersonArrivedMallResponse() (response *DescribePersonArrivedMallResponse) {
    response = &DescribePersonArrivedMallResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePersonArrivedMall
// 输出开始时间到结束时间段内的进出场数据。不做按天聚合的情况下，每次进出场，产生一条进出场数据。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonArrivedMall(request *DescribePersonArrivedMallRequest) (response *DescribePersonArrivedMallResponse, err error) {
    return c.DescribePersonArrivedMallWithContext(context.Background(), request)
}

// DescribePersonArrivedMall
// 输出开始时间到结束时间段内的进出场数据。不做按天聚合的情况下，每次进出场，产生一条进出场数据。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonArrivedMallWithContext(ctx context.Context, request *DescribePersonArrivedMallRequest) (response *DescribePersonArrivedMallResponse, err error) {
    if request == nil {
        request = NewDescribePersonArrivedMallRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonArrivedMall require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonArrivedMallResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonInfoRequest() (request *DescribePersonInfoRequest) {
    request = &DescribePersonInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonInfo")
    
    
    return
}

func NewDescribePersonInfoResponse() (response *DescribePersonInfoResponse) {
    response = &DescribePersonInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePersonInfo
// 指定门店获取所有顾客详情列表，包含客户ID、图片、年龄、性别
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePersonInfo(request *DescribePersonInfoRequest) (response *DescribePersonInfoResponse, err error) {
    return c.DescribePersonInfoWithContext(context.Background(), request)
}

// DescribePersonInfo
// 指定门店获取所有顾客详情列表，包含客户ID、图片、年龄、性别
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePersonInfoWithContext(ctx context.Context, request *DescribePersonInfoRequest) (response *DescribePersonInfoResponse, err error) {
    if request == nil {
        request = NewDescribePersonInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonInfoByFacePictureRequest() (request *DescribePersonInfoByFacePictureRequest) {
    request = &DescribePersonInfoByFacePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonInfoByFacePicture")
    
    
    return
}

func NewDescribePersonInfoByFacePictureResponse() (response *DescribePersonInfoByFacePictureResponse) {
    response = &DescribePersonInfoByFacePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePersonInfoByFacePicture
// 通过上传人脸图片检索系统face id、顾客身份信息及底图
//
// 可能返回的错误码:
//  FAILEDOPERATION_BADFACEQUALITY = "FailedOperation.BadFaceQuality"
//  FAILEDOPERATION_FACENOTFOUND = "FailedOperation.FaceNotFound"
//  FAILEDOPERATION_FACESIZEERROR = "FailedOperation.FaceSizeError"
//  FAILEDOPERATION_MULTIFACEDETECTED = "FailedOperation.MultiFaceDetected"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NOPERSON = "FailedOperation.NoPerson"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  FAILEDOPERATION_OTHERS = "FailedOperation.Others"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonInfoByFacePicture(request *DescribePersonInfoByFacePictureRequest) (response *DescribePersonInfoByFacePictureResponse, err error) {
    return c.DescribePersonInfoByFacePictureWithContext(context.Background(), request)
}

// DescribePersonInfoByFacePicture
// 通过上传人脸图片检索系统face id、顾客身份信息及底图
//
// 可能返回的错误码:
//  FAILEDOPERATION_BADFACEQUALITY = "FailedOperation.BadFaceQuality"
//  FAILEDOPERATION_FACENOTFOUND = "FailedOperation.FaceNotFound"
//  FAILEDOPERATION_FACESIZEERROR = "FailedOperation.FaceSizeError"
//  FAILEDOPERATION_MULTIFACEDETECTED = "FailedOperation.MultiFaceDetected"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NOPERSON = "FailedOperation.NoPerson"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  FAILEDOPERATION_OTHERS = "FailedOperation.Others"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonInfoByFacePictureWithContext(ctx context.Context, request *DescribePersonInfoByFacePictureRequest) (response *DescribePersonInfoByFacePictureResponse, err error) {
    if request == nil {
        request = NewDescribePersonInfoByFacePictureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonInfoByFacePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonInfoByFacePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonTraceRequest() (request *DescribePersonTraceRequest) {
    request = &DescribePersonTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonTrace")
    
    
    return
}

func NewDescribePersonTraceResponse() (response *DescribePersonTraceResponse) {
    response = &DescribePersonTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePersonTrace
// 输出开始时间到结束时间段内的进出场数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonTrace(request *DescribePersonTraceRequest) (response *DescribePersonTraceResponse, err error) {
    return c.DescribePersonTraceWithContext(context.Background(), request)
}

// DescribePersonTrace
// 输出开始时间到结束时间段内的进出场数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonTraceWithContext(ctx context.Context, request *DescribePersonTraceRequest) (response *DescribePersonTraceResponse, err error) {
    if request == nil {
        request = NewDescribePersonTraceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonTrace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonTraceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonTraceDetailRequest() (request *DescribePersonTraceDetailRequest) {
    request = &DescribePersonTraceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonTraceDetail")
    
    
    return
}

func NewDescribePersonTraceDetailResponse() (response *DescribePersonTraceDetailResponse) {
    response = &DescribePersonTraceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePersonTraceDetail
// 查询客户单次到场轨迹明细
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonTraceDetail(request *DescribePersonTraceDetailRequest) (response *DescribePersonTraceDetailResponse, err error) {
    return c.DescribePersonTraceDetailWithContext(context.Background(), request)
}

// DescribePersonTraceDetail
// 查询客户单次到场轨迹明细
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePersonTraceDetailWithContext(ctx context.Context, request *DescribePersonTraceDetailRequest) (response *DescribePersonTraceDetailResponse, err error) {
    if request == nil {
        request = NewDescribePersonTraceDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonTraceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonTraceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonVisitInfoRequest() (request *DescribePersonVisitInfoRequest) {
    request = &DescribePersonVisitInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribePersonVisitInfo")
    
    
    return
}

func NewDescribePersonVisitInfoResponse() (response *DescribePersonVisitInfoResponse) {
    response = &DescribePersonVisitInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePersonVisitInfo
// 获取门店指定时间范围内的所有用户到访信息记录，支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePersonVisitInfo(request *DescribePersonVisitInfoRequest) (response *DescribePersonVisitInfoResponse, err error) {
    return c.DescribePersonVisitInfoWithContext(context.Background(), request)
}

// DescribePersonVisitInfo
// 获取门店指定时间范围内的所有用户到访信息记录，支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePersonVisitInfoWithContext(ctx context.Context, request *DescribePersonVisitInfoRequest) (response *DescribePersonVisitInfoResponse, err error) {
    if request == nil {
        request = NewDescribePersonVisitInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonVisitInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonVisitInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShopHourTrafficInfoRequest() (request *DescribeShopHourTrafficInfoRequest) {
    request = &DescribeShopHourTrafficInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeShopHourTrafficInfo")
    
    
    return
}

func NewDescribeShopHourTrafficInfoResponse() (response *DescribeShopHourTrafficInfoResponse) {
    response = &DescribeShopHourTrafficInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeShopHourTrafficInfo
// 按小时提供查询日期范围内门店的每天每小时累计客流人数数据，支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeShopHourTrafficInfo(request *DescribeShopHourTrafficInfoRequest) (response *DescribeShopHourTrafficInfoResponse, err error) {
    return c.DescribeShopHourTrafficInfoWithContext(context.Background(), request)
}

// DescribeShopHourTrafficInfo
// 按小时提供查询日期范围内门店的每天每小时累计客流人数数据，支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeShopHourTrafficInfoWithContext(ctx context.Context, request *DescribeShopHourTrafficInfoRequest) (response *DescribeShopHourTrafficInfoResponse, err error) {
    if request == nil {
        request = NewDescribeShopHourTrafficInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShopHourTrafficInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShopHourTrafficInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShopInfoRequest() (request *DescribeShopInfoRequest) {
    request = &DescribeShopInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeShopInfo")
    
    
    return
}

func NewDescribeShopInfoResponse() (response *DescribeShopInfoResponse) {
    response = &DescribeShopInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeShopInfo
// 根据客户身份标识获取客户下所有的门店信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeShopInfo(request *DescribeShopInfoRequest) (response *DescribeShopInfoResponse, err error) {
    return c.DescribeShopInfoWithContext(context.Background(), request)
}

// DescribeShopInfo
// 根据客户身份标识获取客户下所有的门店信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeShopInfoWithContext(ctx context.Context, request *DescribeShopInfoRequest) (response *DescribeShopInfoResponse, err error) {
    if request == nil {
        request = NewDescribeShopInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShopInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShopInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShopTrafficInfoRequest() (request *DescribeShopTrafficInfoRequest) {
    request = &DescribeShopTrafficInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeShopTrafficInfo")
    
    
    return
}

func NewDescribeShopTrafficInfoResponse() (response *DescribeShopTrafficInfoResponse) {
    response = &DescribeShopTrafficInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeShopTrafficInfo
// 按天提供查询日期范围内门店的单日累计客流人数，支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeShopTrafficInfo(request *DescribeShopTrafficInfoRequest) (response *DescribeShopTrafficInfoResponse, err error) {
    return c.DescribeShopTrafficInfoWithContext(context.Background(), request)
}

// DescribeShopTrafficInfo
// 按天提供查询日期范围内门店的单日累计客流人数，支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeShopTrafficInfoWithContext(ctx context.Context, request *DescribeShopTrafficInfoRequest) (response *DescribeShopTrafficInfoResponse, err error) {
    if request == nil {
        request = NewDescribeShopTrafficInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShopTrafficInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShopTrafficInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrajectoryDataRequest() (request *DescribeTrajectoryDataRequest) {
    request = &DescribeTrajectoryDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeTrajectoryData")
    
    
    return
}

func NewDescribeTrajectoryDataResponse() (response *DescribeTrajectoryDataResponse) {
    response = &DescribeTrajectoryDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrajectoryData
// 获取动线轨迹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTrajectoryData(request *DescribeTrajectoryDataRequest) (response *DescribeTrajectoryDataResponse, err error) {
    return c.DescribeTrajectoryDataWithContext(context.Background(), request)
}

// DescribeTrajectoryData
// 获取动线轨迹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeTrajectoryDataWithContext(ctx context.Context, request *DescribeTrajectoryDataRequest) (response *DescribeTrajectoryDataResponse, err error) {
    if request == nil {
        request = NewDescribeTrajectoryDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrajectoryData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrajectoryDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneFlowAgeInfoByZoneIdRequest() (request *DescribeZoneFlowAgeInfoByZoneIdRequest) {
    request = &DescribeZoneFlowAgeInfoByZoneIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneFlowAgeInfoByZoneId")
    
    
    return
}

func NewDescribeZoneFlowAgeInfoByZoneIdResponse() (response *DescribeZoneFlowAgeInfoByZoneIdResponse) {
    response = &DescribeZoneFlowAgeInfoByZoneIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneFlowAgeInfoByZoneId
// 获取指定区域人流各年龄占比
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowAgeInfoByZoneId(request *DescribeZoneFlowAgeInfoByZoneIdRequest) (response *DescribeZoneFlowAgeInfoByZoneIdResponse, err error) {
    return c.DescribeZoneFlowAgeInfoByZoneIdWithContext(context.Background(), request)
}

// DescribeZoneFlowAgeInfoByZoneId
// 获取指定区域人流各年龄占比
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowAgeInfoByZoneIdWithContext(ctx context.Context, request *DescribeZoneFlowAgeInfoByZoneIdRequest) (response *DescribeZoneFlowAgeInfoByZoneIdResponse, err error) {
    if request == nil {
        request = NewDescribeZoneFlowAgeInfoByZoneIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneFlowAgeInfoByZoneId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneFlowAgeInfoByZoneIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneFlowAndStayTimeRequest() (request *DescribeZoneFlowAndStayTimeRequest) {
    request = &DescribeZoneFlowAndStayTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneFlowAndStayTime")
    
    
    return
}

func NewDescribeZoneFlowAndStayTimeResponse() (response *DescribeZoneFlowAndStayTimeResponse) {
    response = &DescribeZoneFlowAndStayTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneFlowAndStayTime
// 获取区域人流和停留时间
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowAndStayTime(request *DescribeZoneFlowAndStayTimeRequest) (response *DescribeZoneFlowAndStayTimeResponse, err error) {
    return c.DescribeZoneFlowAndStayTimeWithContext(context.Background(), request)
}

// DescribeZoneFlowAndStayTime
// 获取区域人流和停留时间
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowAndStayTimeWithContext(ctx context.Context, request *DescribeZoneFlowAndStayTimeRequest) (response *DescribeZoneFlowAndStayTimeResponse, err error) {
    if request == nil {
        request = NewDescribeZoneFlowAndStayTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneFlowAndStayTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneFlowAndStayTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneFlowDailyByZoneIdRequest() (request *DescribeZoneFlowDailyByZoneIdRequest) {
    request = &DescribeZoneFlowDailyByZoneIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneFlowDailyByZoneId")
    
    
    return
}

func NewDescribeZoneFlowDailyByZoneIdResponse() (response *DescribeZoneFlowDailyByZoneIdResponse) {
    response = &DescribeZoneFlowDailyByZoneIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneFlowDailyByZoneId
// 获取指定区域每日客流量
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowDailyByZoneId(request *DescribeZoneFlowDailyByZoneIdRequest) (response *DescribeZoneFlowDailyByZoneIdResponse, err error) {
    return c.DescribeZoneFlowDailyByZoneIdWithContext(context.Background(), request)
}

// DescribeZoneFlowDailyByZoneId
// 获取指定区域每日客流量
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowDailyByZoneIdWithContext(ctx context.Context, request *DescribeZoneFlowDailyByZoneIdRequest) (response *DescribeZoneFlowDailyByZoneIdResponse, err error) {
    if request == nil {
        request = NewDescribeZoneFlowDailyByZoneIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneFlowDailyByZoneId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneFlowDailyByZoneIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneFlowGenderAvrStayTimeByZoneIdRequest() (request *DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest) {
    request = &DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneFlowGenderAvrStayTimeByZoneId")
    
    
    return
}

func NewDescribeZoneFlowGenderAvrStayTimeByZoneIdResponse() (response *DescribeZoneFlowGenderAvrStayTimeByZoneIdResponse) {
    response = &DescribeZoneFlowGenderAvrStayTimeByZoneIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneFlowGenderAvrStayTimeByZoneId
// 获取指定区域不同年龄段男女平均停留时间
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowGenderAvrStayTimeByZoneId(request *DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest) (response *DescribeZoneFlowGenderAvrStayTimeByZoneIdResponse, err error) {
    return c.DescribeZoneFlowGenderAvrStayTimeByZoneIdWithContext(context.Background(), request)
}

// DescribeZoneFlowGenderAvrStayTimeByZoneId
// 获取指定区域不同年龄段男女平均停留时间
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowGenderAvrStayTimeByZoneIdWithContext(ctx context.Context, request *DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest) (response *DescribeZoneFlowGenderAvrStayTimeByZoneIdResponse, err error) {
    if request == nil {
        request = NewDescribeZoneFlowGenderAvrStayTimeByZoneIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneFlowGenderAvrStayTimeByZoneId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneFlowGenderAvrStayTimeByZoneIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneFlowGenderInfoByZoneIdRequest() (request *DescribeZoneFlowGenderInfoByZoneIdRequest) {
    request = &DescribeZoneFlowGenderInfoByZoneIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneFlowGenderInfoByZoneId")
    
    
    return
}

func NewDescribeZoneFlowGenderInfoByZoneIdResponse() (response *DescribeZoneFlowGenderInfoByZoneIdResponse) {
    response = &DescribeZoneFlowGenderInfoByZoneIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneFlowGenderInfoByZoneId
// 获取指定区域性别占比
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowGenderInfoByZoneId(request *DescribeZoneFlowGenderInfoByZoneIdRequest) (response *DescribeZoneFlowGenderInfoByZoneIdResponse, err error) {
    return c.DescribeZoneFlowGenderInfoByZoneIdWithContext(context.Background(), request)
}

// DescribeZoneFlowGenderInfoByZoneId
// 获取指定区域性别占比
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowGenderInfoByZoneIdWithContext(ctx context.Context, request *DescribeZoneFlowGenderInfoByZoneIdRequest) (response *DescribeZoneFlowGenderInfoByZoneIdResponse, err error) {
    if request == nil {
        request = NewDescribeZoneFlowGenderInfoByZoneIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneFlowGenderInfoByZoneId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneFlowGenderInfoByZoneIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneFlowHourlyByZoneIdRequest() (request *DescribeZoneFlowHourlyByZoneIdRequest) {
    request = &DescribeZoneFlowHourlyByZoneIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneFlowHourlyByZoneId")
    
    
    return
}

func NewDescribeZoneFlowHourlyByZoneIdResponse() (response *DescribeZoneFlowHourlyByZoneIdResponse) {
    response = &DescribeZoneFlowHourlyByZoneIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneFlowHourlyByZoneId
// 获取指定区域分时客流量
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowHourlyByZoneId(request *DescribeZoneFlowHourlyByZoneIdRequest) (response *DescribeZoneFlowHourlyByZoneIdResponse, err error) {
    return c.DescribeZoneFlowHourlyByZoneIdWithContext(context.Background(), request)
}

// DescribeZoneFlowHourlyByZoneId
// 获取指定区域分时客流量
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneFlowHourlyByZoneIdWithContext(ctx context.Context, request *DescribeZoneFlowHourlyByZoneIdRequest) (response *DescribeZoneFlowHourlyByZoneIdResponse, err error) {
    if request == nil {
        request = NewDescribeZoneFlowHourlyByZoneIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneFlowHourlyByZoneId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneFlowHourlyByZoneIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneTrafficInfoRequest() (request *DescribeZoneTrafficInfoRequest) {
    request = &DescribeZoneTrafficInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "DescribeZoneTrafficInfo")
    
    
    return
}

func NewDescribeZoneTrafficInfoResponse() (response *DescribeZoneTrafficInfoResponse) {
    response = &DescribeZoneTrafficInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneTrafficInfo
// 按天提供查询日期范围内，客户指定门店下的所有区域（优Mall部署时已配置区域）的累计客流人次和平均停留时间。支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneTrafficInfo(request *DescribeZoneTrafficInfoRequest) (response *DescribeZoneTrafficInfoResponse, err error) {
    return c.DescribeZoneTrafficInfoWithContext(context.Background(), request)
}

// DescribeZoneTrafficInfo
// 按天提供查询日期范围内，客户指定门店下的所有区域（优Mall部署时已配置区域）的累计客流人次和平均停留时间。支持的时间范围：过去365天，含当天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeZoneTrafficInfoWithContext(ctx context.Context, request *DescribeZoneTrafficInfoRequest) (response *DescribeZoneTrafficInfoResponse, err error) {
    if request == nil {
        request = NewDescribeZoneTrafficInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneTrafficInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneTrafficInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonFeatureInfoRequest() (request *ModifyPersonFeatureInfoRequest) {
    request = &ModifyPersonFeatureInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "ModifyPersonFeatureInfo")
    
    
    return
}

func NewModifyPersonFeatureInfoResponse() (response *ModifyPersonFeatureInfoResponse) {
    response = &ModifyPersonFeatureInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPersonFeatureInfo
// 支持修改黑白名单类型的顾客特征
//
// 可能返回的错误码:
//  FAILEDOPERATION_BADFACEQUALITY = "FailedOperation.BadFaceQuality"
//  FAILEDOPERATION_EXTRACTFEATUREERROR = "FailedOperation.ExtractFeatureError"
//  FAILEDOPERATION_FACENOTFOUND = "FailedOperation.FaceNotFound"
//  FAILEDOPERATION_FACESIZEERROR = "FailedOperation.FaceSizeError"
//  FAILEDOPERATION_HASEXISTPERSON = "FailedOperation.HasExistPerson"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MULTIFACEDETECTED = "FailedOperation.MultiFaceDetected"
//  FAILEDOPERATION_NOPERSON = "FailedOperation.NoPerson"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPersonFeatureInfo(request *ModifyPersonFeatureInfoRequest) (response *ModifyPersonFeatureInfoResponse, err error) {
    return c.ModifyPersonFeatureInfoWithContext(context.Background(), request)
}

// ModifyPersonFeatureInfo
// 支持修改黑白名单类型的顾客特征
//
// 可能返回的错误码:
//  FAILEDOPERATION_BADFACEQUALITY = "FailedOperation.BadFaceQuality"
//  FAILEDOPERATION_EXTRACTFEATUREERROR = "FailedOperation.ExtractFeatureError"
//  FAILEDOPERATION_FACENOTFOUND = "FailedOperation.FaceNotFound"
//  FAILEDOPERATION_FACESIZEERROR = "FailedOperation.FaceSizeError"
//  FAILEDOPERATION_HASEXISTPERSON = "FailedOperation.HasExistPerson"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MULTIFACEDETECTED = "FailedOperation.MultiFaceDetected"
//  FAILEDOPERATION_NOPERSON = "FailedOperation.NoPerson"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPersonFeatureInfoWithContext(ctx context.Context, request *ModifyPersonFeatureInfoRequest) (response *ModifyPersonFeatureInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonFeatureInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonFeatureInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonFeatureInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonTagInfoRequest() (request *ModifyPersonTagInfoRequest) {
    request = &ModifyPersonTagInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "ModifyPersonTagInfo")
    
    
    return
}

func NewModifyPersonTagInfoResponse() (response *ModifyPersonTagInfoResponse) {
    response = &ModifyPersonTagInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPersonTagInfo
// 标记到店顾客的身份类型，例如黑名单、白名单等
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPersonTagInfo(request *ModifyPersonTagInfoRequest) (response *ModifyPersonTagInfoResponse, err error) {
    return c.ModifyPersonTagInfoWithContext(context.Background(), request)
}

// ModifyPersonTagInfo
// 标记到店顾客的身份类型，例如黑名单、白名单等
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyPersonTagInfoWithContext(ctx context.Context, request *ModifyPersonTagInfoRequest) (response *ModifyPersonTagInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonTagInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonTagInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonTagInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonTypeRequest() (request *ModifyPersonTypeRequest) {
    request = &ModifyPersonTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "ModifyPersonType")
    
    
    return
}

func NewModifyPersonTypeResponse() (response *ModifyPersonTypeResponse) {
    response = &ModifyPersonTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPersonType
// 修改顾客身份类型接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyPersonType(request *ModifyPersonTypeRequest) (response *ModifyPersonTypeResponse, err error) {
    return c.ModifyPersonTypeWithContext(context.Background(), request)
}

// ModifyPersonType
// 修改顾客身份类型接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTEXIST = "FailedOperation.AccountExist"
//  FAILEDOPERATION_NEEDGRANTROLEFIRST = "FailedOperation.NeedGrantRoleFirst"
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_NORIGHT = "FailedOperation.NoRight"
//  FAILEDOPERATION_NOTMATCHSHOPCODE = "FailedOperation.NotMatchShopCode"
//  FAILEDOPERATION_PARAMETERERROR = "FailedOperation.ParameterError"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATAERROR = "InternalError.DataError"
//  INTERNALERROR_DATAHASEXISTS = "InternalError.DataHasExists"
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INTERNALERROR_USERNOTEXIST = "InternalError.UserNotExist"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyPersonTypeWithContext(ctx context.Context, request *ModifyPersonTypeRequest) (response *ModifyPersonTypeResponse, err error) {
    if request == nil {
        request = NewModifyPersonTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPersonTypeResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterCallbackRequest() (request *RegisterCallbackRequest) {
    request = &RegisterCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("youmall", APIVersion, "RegisterCallback")
    
    
    return
}

func NewRegisterCallbackResponse() (response *RegisterCallbackResponse) {
    response = &RegisterCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RegisterCallback
// 调用本接口在优Mall中注册自己集团的到店通知回调接口地址，接口协议为HTTP或HTTPS。注册后，若集团有特殊身份（例如老客）到店通知，优Mall后台将主动将到店信息push给该接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
func (c *Client) RegisterCallback(request *RegisterCallbackRequest) (response *RegisterCallbackResponse, err error) {
    return c.RegisterCallbackWithContext(context.Background(), request)
}

// RegisterCallback
// 调用本接口在优Mall中注册自己集团的到店通知回调接口地址，接口协议为HTTP或HTTPS。注册后，若集团有特殊身份（例如老客）到店通知，优Mall后台将主动将到店信息push给该接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODATA = "FailedOperation.NoData"
//  FAILEDOPERATION_PROCESSFAIL = "FailedOperation.ProcessFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_JSONPARSEERR = "InvalidParameterValue.JsonParseErr"
func (c *Client) RegisterCallbackWithContext(ctx context.Context, request *RegisterCallbackRequest) (response *RegisterCallbackResponse, err error) {
    if request == nil {
        request = NewRegisterCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterCallbackResponse()
    err = c.Send(request, response)
    return
}
