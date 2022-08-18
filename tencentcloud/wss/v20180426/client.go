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

package v20180426

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-26"

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


func NewDeleteCertRequest() (request *DeleteCertRequest) {
    request = &DeleteCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wss", APIVersion, "DeleteCert")
    
    
    return
}

func NewDeleteCertResponse() (response *DeleteCertResponse) {
    response = &DeleteCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCert
// 本接口（DeleteCert）用于删除证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTINVALIDPARAM = "FailedOperation.CertInvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
func (c *Client) DeleteCert(request *DeleteCertRequest) (response *DeleteCertResponse, err error) {
    return c.DeleteCertWithContext(context.Background(), request)
}

// DeleteCert
// 本接口（DeleteCert）用于删除证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTINVALIDPARAM = "FailedOperation.CertInvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
func (c *Client) DeleteCertWithContext(ctx context.Context, request *DeleteCertRequest) (response *DeleteCertResponse, err error) {
    if request == nil {
        request = NewDeleteCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCertResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertListRequest() (request *DescribeCertListRequest) {
    request = &DescribeCertListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wss", APIVersion, "DescribeCertList")
    
    
    return
}

func NewDescribeCertListResponse() (response *DescribeCertListResponse) {
    response = &DescribeCertListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertList
// 本接口(DescribeCertList)用于获取证书列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
func (c *Client) DescribeCertList(request *DescribeCertListRequest) (response *DescribeCertListResponse, err error) {
    return c.DescribeCertListWithContext(context.Background(), request)
}

// DescribeCertList
// 本接口(DescribeCertList)用于获取证书列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
func (c *Client) DescribeCertListWithContext(ctx context.Context, request *DescribeCertListRequest) (response *DescribeCertListResponse, err error) {
    if request == nil {
        request = NewDescribeCertListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertListResponse()
    err = c.Send(request, response)
    return
}

func NewUploadCertRequest() (request *UploadCertRequest) {
    request = &UploadCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wss", APIVersion, "UploadCert")
    
    
    return
}

func NewUploadCertResponse() (response *UploadCertResponse) {
    response = &UploadCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadCert
// 本接口（UploadCert）用于上传证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTINVALIDPARAM = "FailedOperation.CertInvalidParam"
//  FAILEDOPERATION_CERTMISMATCH = "FailedOperation.CertMismatch"
//  FAILEDOPERATION_INVALIDCERT = "FailedOperation.InvalidCert"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
func (c *Client) UploadCert(request *UploadCertRequest) (response *UploadCertResponse, err error) {
    return c.UploadCertWithContext(context.Background(), request)
}

// UploadCert
// 本接口（UploadCert）用于上传证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTINVALIDPARAM = "FailedOperation.CertInvalidParam"
//  FAILEDOPERATION_CERTMISMATCH = "FailedOperation.CertMismatch"
//  FAILEDOPERATION_INVALIDCERT = "FailedOperation.InvalidCert"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DOMAININTERNALERROR = "InternalError.DomainInternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REPTYPEISINVALID = "InvalidParameter.RepTypeIsInvalid"
func (c *Client) UploadCertWithContext(ctx context.Context, request *UploadCertRequest) (response *UploadCertResponse, err error) {
    if request == nil {
        request = NewUploadCertRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadCert require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadCertResponse()
    err = c.Send(request, response)
    return
}
