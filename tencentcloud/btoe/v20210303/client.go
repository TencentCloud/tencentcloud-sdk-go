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

package v20210303

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-03-03"

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


func NewCreateAudioDepositRequest() (request *CreateAudioDepositRequest) {
    request = &CreateAudioDepositRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateAudioDeposit")
    
    
    return
}

func NewCreateAudioDepositResponse() (response *CreateAudioDepositResponse) {
    response = &CreateAudioDepositResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAudioDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的音频原文件或下载URL，BTOE对音频原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。音频类型支持格式：mp3、wav、wma、midi、flac；原文件上传大小不超过5 MB，下载URL文件大小不超过25 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateAudioDeposit(request *CreateAudioDepositRequest) (response *CreateAudioDepositResponse, err error) {
    return c.CreateAudioDepositWithContext(context.Background(), request)
}

// CreateAudioDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的音频原文件或下载URL，BTOE对音频原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。音频类型支持格式：mp3、wav、wma、midi、flac；原文件上传大小不超过5 MB，下载URL文件大小不超过25 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateAudioDepositWithContext(ctx context.Context, request *CreateAudioDepositRequest) (response *CreateAudioDepositResponse, err error) {
    if request == nil {
        request = NewCreateAudioDepositRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAudioDeposit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAudioDepositResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataDepositRequest() (request *CreateDataDepositRequest) {
    request = &CreateDataDepositRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateDataDeposit")
    
    
    return
}

func NewCreateDataDepositResponse() (response *CreateDataDepositResponse) {
    response = &CreateDataDepositResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDataDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的业务数据明文，业务数据明文存证写入后不可修改，BTOE对业务数据明文存证生成含有电子签章的区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_SENSITIVEDATA = "FailedOperation.SensitiveData"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATAINFOTOOLONG = "InvalidParameterValue.DataInfoTooLong"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateDataDeposit(request *CreateDataDepositRequest) (response *CreateDataDepositResponse, err error) {
    return c.CreateDataDepositWithContext(context.Background(), request)
}

// CreateDataDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的业务数据明文，业务数据明文存证写入后不可修改，BTOE对业务数据明文存证生成含有电子签章的区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_SENSITIVEDATA = "FailedOperation.SensitiveData"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATAINFOTOOLONG = "InvalidParameterValue.DataInfoTooLong"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateDataDepositWithContext(ctx context.Context, request *CreateDataDepositRequest) (response *CreateDataDepositResponse, err error) {
    if request == nil {
        request = NewCreateDataDepositRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataDeposit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataDepositResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDocDepositRequest() (request *CreateDocDepositRequest) {
    request = &CreateDocDepositRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateDocDeposit")
    
    
    return
}

func NewCreateDocDepositResponse() (response *CreateDocDepositResponse) {
    response = &CreateDocDepositResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDocDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的文档原文件或下载URL，BTOE对文档原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。文档类型支持格式：doc、docx、xls、xlsx、ppt、pptx、 pdf、html、txt、md、csv；原文件上传大小不超过5 MB，下载URL文件大小不超过10 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEREADFAILED = "FailedOperation.FileReadFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILESUFFIX = "InvalidParameter.InvalidFileSuffix"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateDocDeposit(request *CreateDocDepositRequest) (response *CreateDocDepositResponse, err error) {
    return c.CreateDocDepositWithContext(context.Background(), request)
}

// CreateDocDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的文档原文件或下载URL，BTOE对文档原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。文档类型支持格式：doc、docx、xls、xlsx、ppt、pptx、 pdf、html、txt、md、csv；原文件上传大小不超过5 MB，下载URL文件大小不超过10 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEREADFAILED = "FailedOperation.FileReadFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILESUFFIX = "InvalidParameter.InvalidFileSuffix"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateDocDepositWithContext(ctx context.Context, request *CreateDocDepositRequest) (response *CreateDocDepositResponse, err error) {
    if request == nil {
        request = NewCreateDocDepositRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocDeposit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocDepositResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHashDepositRequest() (request *CreateHashDepositRequest) {
    request = &CreateHashDepositRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateHashDeposit")
    
    
    return
}

func NewCreateHashDepositResponse() (response *CreateHashDepositResponse) {
    response = &CreateHashDepositResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHashDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的原文数据Hash值，BTOE对业务数据Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateHashDeposit(request *CreateHashDepositRequest) (response *CreateHashDepositResponse, err error) {
    return c.CreateHashDepositWithContext(context.Background(), request)
}

// CreateHashDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的原文数据Hash值，BTOE对业务数据Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateHashDepositWithContext(ctx context.Context, request *CreateHashDepositRequest) (response *CreateHashDepositResponse, err error) {
    if request == nil {
        request = NewCreateHashDepositRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHashDeposit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHashDepositResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHashDepositNoCertRequest() (request *CreateHashDepositNoCertRequest) {
    request = &CreateHashDepositNoCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateHashDepositNoCert")
    
    
    return
}

func NewCreateHashDepositNoCertResponse() (response *CreateHashDepositNoCertResponse) {
    response = &CreateHashDepositNoCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHashDepositNoCert
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的原文数据Hash值，BTOE对业务数据Hash值存证上链，本接口不生成区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateHashDepositNoCert(request *CreateHashDepositNoCertRequest) (response *CreateHashDepositNoCertResponse, err error) {
    return c.CreateHashDepositNoCertWithContext(context.Background(), request)
}

// CreateHashDepositNoCert
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的原文数据Hash值，BTOE对业务数据Hash值存证上链，本接口不生成区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateHashDepositNoCertWithContext(ctx context.Context, request *CreateHashDepositNoCertRequest) (response *CreateHashDepositNoCertResponse, err error) {
    if request == nil {
        request = NewCreateHashDepositNoCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHashDepositNoCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHashDepositNoCertResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHashDepositNoSealRequest() (request *CreateHashDepositNoSealRequest) {
    request = &CreateHashDepositNoSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateHashDepositNoSeal")
    
    
    return
}

func NewCreateHashDepositNoSealResponse() (response *CreateHashDepositNoSealResponse) {
    response = &CreateHashDepositNoSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHashDepositNoSeal
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的原文数据Hash值，BTOE对业务数据Hash值存证上链，并生成无电子签章的区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateHashDepositNoSeal(request *CreateHashDepositNoSealRequest) (response *CreateHashDepositNoSealResponse, err error) {
    return c.CreateHashDepositNoSealWithContext(context.Background(), request)
}

// CreateHashDepositNoSeal
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的原文数据Hash值，BTOE对业务数据Hash值存证上链，并生成无电子签章的区块链存证电子凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateHashDepositNoSealWithContext(ctx context.Context, request *CreateHashDepositNoSealRequest) (response *CreateHashDepositNoSealResponse, err error) {
    if request == nil {
        request = NewCreateHashDepositNoSealRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHashDepositNoSeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHashDepositNoSealResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageDepositRequest() (request *CreateImageDepositRequest) {
    request = &CreateImageDepositRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateImageDeposit")
    
    
    return
}

func NewCreateImageDepositResponse() (response *CreateImageDepositResponse) {
    response = &CreateImageDepositResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImageDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的图片原文件或下载URL，BTOE对图片原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。图片类型支持格式：png、jpg、jpeg、bmp、gif、svg；原文件上传大小不超过5 MB，下载URL文件大小不超过10 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEENCODINDFORMATERROR = "FailedOperation.FileEncodindFormatError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCEINSUFFICIENT_LOWBALANCE = "ResourceInsufficient.LowBalance"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateImageDeposit(request *CreateImageDepositRequest) (response *CreateImageDepositResponse, err error) {
    return c.CreateImageDepositWithContext(context.Background(), request)
}

// CreateImageDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的图片原文件或下载URL，BTOE对图片原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。图片类型支持格式：png、jpg、jpeg、bmp、gif、svg；原文件上传大小不超过5 MB，下载URL文件大小不超过10 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEENCODINDFORMATERROR = "FailedOperation.FileEncodindFormatError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCEINSUFFICIENT_LOWBALANCE = "ResourceInsufficient.LowBalance"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateImageDepositWithContext(ctx context.Context, request *CreateImageDepositRequest) (response *CreateImageDepositResponse, err error) {
    if request == nil {
        request = NewCreateImageDepositRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageDeposit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageDepositResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoDepositRequest() (request *CreateVideoDepositRequest) {
    request = &CreateVideoDepositRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateVideoDeposit")
    
    
    return
}

func NewCreateVideoDepositResponse() (response *CreateVideoDepositResponse) {
    response = &CreateVideoDepositResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVideoDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的视频的原文件或下载URL，BTOE对视频原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。视频文件支持格式：mp4、avi、mkv、mov、flv,wmv,rmvb,3gp；文件大小限制：直接上传原文件不大于5MB，下载URL文件大小不大于200 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEENCODINDFORMATERROR = "FailedOperation.FileEncodindFormatError"
//  FAILEDOPERATION_FILEREADFAILED = "FailedOperation.FileReadFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILESUFFIX = "InvalidParameter.InvalidFileSuffix"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateVideoDeposit(request *CreateVideoDepositRequest) (response *CreateVideoDepositResponse, err error) {
    return c.CreateVideoDepositWithContext(context.Background(), request)
}

// CreateVideoDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE写入待存证的视频的原文件或下载URL，BTOE对视频原文件存储后，将其Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。视频文件支持格式：mp4、avi、mkv、mov、flv,wmv,rmvb,3gp；文件大小限制：直接上传原文件不大于5MB，下载URL文件大小不大于200 MB。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEENCODINDFORMATERROR = "FailedOperation.FileEncodindFormatError"
//  FAILEDOPERATION_FILEREADFAILED = "FailedOperation.FileReadFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDFILESUFFIX = "InvalidParameter.InvalidFileSuffix"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_HASHNOMATCH = "InvalidParameterValue.HashNoMatch"
//  INVALIDPARAMETERVALUE_INVALIDFILESUFFIX = "InvalidParameterValue.InvalidFileSuffix"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  INVALIDPARAMETERVALUE_TOOLARGEFILEERROR = "InvalidParameterValue.TooLargeFileError"
//  RESOURCENOTFOUND_DOWNLOADERROR = "ResourceNotFound.DownLoadError"
func (c *Client) CreateVideoDepositWithContext(ctx context.Context, request *CreateVideoDepositRequest) (response *CreateVideoDepositResponse, err error) {
    if request == nil {
        request = NewCreateVideoDepositRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoDeposit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoDepositResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebpageDepositRequest() (request *CreateWebpageDepositRequest) {
    request = &CreateWebpageDepositRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "CreateWebpageDeposit")
    
    
    return
}

func NewCreateWebpageDepositResponse() (response *CreateWebpageDepositResponse) {
    response = &CreateWebpageDepositResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWebpageDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE提交待存证网页的URL，BTOE对URL进行网页快照，并将快照图片存储，将网页快照Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。URL格式必须以http、https开头。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  RESOURCEINSUFFICIENT_LOWBALANCE = "ResourceInsufficient.LowBalance"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateWebpageDeposit(request *CreateWebpageDepositRequest) (response *CreateWebpageDepositResponse, err error) {
    return c.CreateWebpageDepositWithContext(context.Background(), request)
}

// CreateWebpageDeposit
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过本接口向BTOE提交待存证网页的URL，BTOE对URL进行网页快照，并将快照图片存储，将网页快照Hash值存证上链，并生成含有电子签章的区块链存证电子凭证。URL格式必须以http、https开头。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETER_INVALIDURL = "InvalidParameter.InvalidURL"
//  INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidURL"
//  RESOURCEINSUFFICIENT_LOWBALANCE = "ResourceInsufficient.LowBalance"
//  RESOURCEUNAVAILABLE_RESOURCENOTOPENED = "ResourceUnavailable.ResourceNotOpened"
func (c *Client) CreateWebpageDepositWithContext(ctx context.Context, request *CreateWebpageDepositRequest) (response *CreateWebpageDepositResponse, err error) {
    if request == nil {
        request = NewCreateWebpageDepositRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebpageDeposit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebpageDepositResponse()
    err = c.Send(request, response)
    return
}

func NewGetDepositCertRequest() (request *GetDepositCertRequest) {
    request = &GetDepositCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "GetDepositCert")
    
    
    return
}

func NewGetDepositCertResponse() (response *GetDepositCertResponse) {
    response = &GetDepositCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDepositCert
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过存证编码向BTOE查询存证电子凭证信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDepositCert(request *GetDepositCertRequest) (response *GetDepositCertResponse, err error) {
    return c.GetDepositCertWithContext(context.Background(), request)
}

// GetDepositCert
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过存证编码向BTOE查询存证电子凭证信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDepositCertWithContext(ctx context.Context, request *GetDepositCertRequest) (response *GetDepositCertResponse, err error) {
    if request == nil {
        request = NewGetDepositCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDepositCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDepositCertResponse()
    err = c.Send(request, response)
    return
}

func NewGetDepositFileRequest() (request *GetDepositFileRequest) {
    request = &GetDepositFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "GetDepositFile")
    
    
    return
}

func NewGetDepositFileResponse() (response *GetDepositFileResponse) {
    response = &GetDepositFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDepositFile
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过存证编码向BTOE获取存证文件的下载URL。
//
// -注：Hash类存证、业务数据明文存证不产生存证文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"
//  FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"
//  FAILEDOPERATION_DATAINFOTOOLONG = "FailedOperation.DataInfoTooLong"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_HASHNOMATCH = "FailedOperation.HashNoMatch"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_SENSITIVEDATA = "FailedOperation.SensitiveData"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTINFOINVALID = "InvalidParameter.AccountInfoInvalid"
//  INVALIDPARAMETER_INVALIDFILESUFFIX = "InvalidParameter.InvalidFileSuffix"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEINSUFFICIENT_LOWBALANCE = "ResourceInsufficient.LowBalance"
func (c *Client) GetDepositFile(request *GetDepositFileRequest) (response *GetDepositFileResponse, err error) {
    return c.GetDepositFileWithContext(context.Background(), request)
}

// GetDepositFile
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过存证编码向BTOE获取存证文件的下载URL。
//
// -注：Hash类存证、业务数据明文存证不产生存证文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"
//  FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"
//  FAILEDOPERATION_DATAINFOTOOLONG = "FailedOperation.DataInfoTooLong"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_HASHNOMATCH = "FailedOperation.HashNoMatch"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_SENSITIVEDATA = "FailedOperation.SensitiveData"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACCOUNTINFOINVALID = "InvalidParameter.AccountInfoInvalid"
//  INVALIDPARAMETER_INVALIDFILESUFFIX = "InvalidParameter.InvalidFileSuffix"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEINSUFFICIENT_LOWBALANCE = "ResourceInsufficient.LowBalance"
func (c *Client) GetDepositFileWithContext(ctx context.Context, request *GetDepositFileRequest) (response *GetDepositFileResponse, err error) {
    if request == nil {
        request = NewGetDepositFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDepositFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDepositFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetDepositInfoRequest() (request *GetDepositInfoRequest) {
    request = &GetDepositInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("btoe", APIVersion, "GetDepositInfo")
    
    
    return
}

func NewGetDepositInfoResponse() (response *GetDepositInfoResponse) {
    response = &GetDepositInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDepositInfo
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过存证编码向BTOE查询存证基本信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEREADFAILED = "FailedOperation.FileReadFailed"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_SENSITIVEDATA = "FailedOperation.SensitiveData"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDepositInfo(request *GetDepositInfoRequest) (response *GetDepositInfoResponse, err error) {
    return c.GetDepositInfoWithContext(context.Background(), request)
}

// GetDepositInfo
// 功能迭代，已上线更高版本的接口2021-05-14
//
// 
//
// 用户通过存证编码向BTOE查询存证基本信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEREADFAILED = "FailedOperation.FileReadFailed"
//  FAILEDOPERATION_ONCHAINFAILURE = "FailedOperation.OnChainFailure"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_SENSITIVEDATA = "FailedOperation.SensitiveData"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERVALUE = "InvalidParameter.InvalidParameterValue"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDepositInfoWithContext(ctx context.Context, request *GetDepositInfoRequest) (response *GetDepositInfoResponse, err error) {
    if request == nil {
        request = NewGetDepositInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDepositInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDepositInfoResponse()
    err = c.Send(request, response)
    return
}
