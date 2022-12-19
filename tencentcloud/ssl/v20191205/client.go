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

package v20191205

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-12-05"

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


func NewApplyCertificateRequest() (request *ApplyCertificateRequest) {
    request = &ApplyCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ApplyCertificate")
    
    
    return
}

func NewApplyCertificateResponse() (response *ApplyCertificateResponse) {
    response = &ApplyCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyCertificate
// 本接口（ApplyCertificate）用于免费证书申请。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"
//  FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
func (c *Client) ApplyCertificate(request *ApplyCertificateRequest) (response *ApplyCertificateResponse, err error) {
    return c.ApplyCertificateWithContext(context.Background(), request)
}

// ApplyCertificate
// 本接口（ApplyCertificate）用于免费证书申请。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"
//  FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"
func (c *Client) ApplyCertificateWithContext(ctx context.Context, request *ApplyCertificateRequest) (response *ApplyCertificateResponse, err error) {
    if request == nil {
        request = NewApplyCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCertificateOrderRequest() (request *CancelCertificateOrderRequest) {
    request = &CancelCertificateOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CancelCertificateOrder")
    
    
    return
}

func NewCancelCertificateOrderResponse() (response *CancelCertificateOrderResponse) {
    response = &CancelCertificateOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelCertificateOrder
// 取消证书订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) CancelCertificateOrder(request *CancelCertificateOrderRequest) (response *CancelCertificateOrderResponse, err error) {
    return c.CancelCertificateOrderWithContext(context.Background(), request)
}

// CancelCertificateOrder
// 取消证书订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) CancelCertificateOrderWithContext(ctx context.Context, request *CancelCertificateOrderRequest) (response *CancelCertificateOrderResponse, err error) {
    if request == nil {
        request = NewCancelCertificateOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelCertificateOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelCertificateOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCertificateChainRequest() (request *CheckCertificateChainRequest) {
    request = &CheckCertificateChainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CheckCertificateChain")
    
    
    return
}

func NewCheckCertificateChainResponse() (response *CheckCertificateChainResponse) {
    response = &CheckCertificateChainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckCertificateChain
// 本接口（CheckCertificateChain）用于检查证书链是否完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
func (c *Client) CheckCertificateChain(request *CheckCertificateChainRequest) (response *CheckCertificateChainResponse, err error) {
    return c.CheckCertificateChainWithContext(context.Background(), request)
}

// CheckCertificateChain
// 本接口（CheckCertificateChain）用于检查证书链是否完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
func (c *Client) CheckCertificateChainWithContext(ctx context.Context, request *CheckCertificateChainRequest) (response *CheckCertificateChainResponse, err error) {
    if request == nil {
        request = NewCheckCertificateChainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCertificateChain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCertificateChainResponse()
    err = c.Send(request, response)
    return
}

func NewCommitCertificateInformationRequest() (request *CommitCertificateInformationRequest) {
    request = &CommitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CommitCertificateInformation")
    
    
    return
}

func NewCommitCertificateInformationResponse() (response *CommitCertificateInformationResponse) {
    response = &CommitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CommitCertificateInformation
// 提交证书订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CommitCertificateInformation(request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    return c.CommitCertificateInformationWithContext(context.Background(), request)
}

// CommitCertificateInformation
// 提交证书订单。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
func (c *Client) CommitCertificateInformationWithContext(ctx context.Context, request *CommitCertificateInformationRequest) (response *CommitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewCommitCertificateInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CommitCertificateInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCommitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteCertificateRequest() (request *CompleteCertificateRequest) {
    request = &CompleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CompleteCertificate")
    
    
    return
}

func NewCompleteCertificateResponse() (response *CompleteCertificateResponse) {
    response = &CompleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CompleteCertificate
// 本接口（CompleteCertificate）用于主动触发证书验证。仅非DNSPod和Wotrus品牌证书支持使用此接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
func (c *Client) CompleteCertificate(request *CompleteCertificateRequest) (response *CompleteCertificateResponse, err error) {
    return c.CompleteCertificateWithContext(context.Background(), request)
}

// CompleteCertificate
// 本接口（CompleteCertificate）用于主动触发证书验证。仅非DNSPod和Wotrus品牌证书支持使用此接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
func (c *Client) CompleteCertificateWithContext(ctx context.Context, request *CompleteCertificateRequest) (response *CompleteCertificateResponse, err error) {
    if request == nil {
        request = NewCompleteCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCertificateRequest() (request *CreateCertificateRequest) {
    request = &CreateCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "CreateCertificate")
    
    
    return
}

func NewCreateCertificateResponse() (response *CreateCertificateResponse) {
    response = &CreateCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCertificate
// 本接口（CreateCertificate）用于创建付费证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCertificate(request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    return c.CreateCertificateWithContext(context.Background(), request)
}

// CreateCertificate
// 本接口（CreateCertificate）用于创建付费证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCertificateWithContext(ctx context.Context, request *CreateCertificateRequest) (response *CreateCertificateResponse, err error) {
    if request == nil {
        request = NewCreateCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCertificateRequest() (request *DeleteCertificateRequest) {
    request = &DeleteCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeleteCertificate")
    
    
    return
}

func NewDeleteCertificateResponse() (response *DeleteCertificateResponse) {
    response = &DeleteCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCertificate
// 本接口（DeleteCertificate）用于删除证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeleteCertificate(request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    return c.DeleteCertificateWithContext(context.Background(), request)
}

// DeleteCertificate
// 本接口（DeleteCertificate）用于删除证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DeleteCertificateWithContext(ctx context.Context, request *DeleteCertificateRequest) (response *DeleteCertificateResponse, err error) {
    if request == nil {
        request = NewDeleteCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteManagerRequest() (request *DeleteManagerRequest) {
    request = &DeleteManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DeleteManager")
    
    
    return
}

func NewDeleteManagerResponse() (response *DeleteManagerResponse) {
    response = &DeleteManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteManager
// 删除管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DeleteManager(request *DeleteManagerRequest) (response *DeleteManagerResponse, err error) {
    return c.DeleteManagerWithContext(context.Background(), request)
}

// DeleteManager
// 删除管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DeleteManagerWithContext(ctx context.Context, request *DeleteManagerRequest) (response *DeleteManagerResponse, err error) {
    if request == nil {
        request = NewDeleteManagerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteManagerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateRequest() (request *DescribeCertificateRequest) {
    request = &DescribeCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificate")
    
    
    return
}

func NewDescribeCertificateResponse() (response *DescribeCertificateResponse) {
    response = &DescribeCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertificate
// 本接口（DescribeCertificate）用于获取证书信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificate(request *DescribeCertificateRequest) (response *DescribeCertificateResponse, err error) {
    return c.DescribeCertificateWithContext(context.Background(), request)
}

// DescribeCertificate
// 本接口（DescribeCertificate）用于获取证书信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificateWithContext(ctx context.Context, request *DescribeCertificateRequest) (response *DescribeCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateDetailRequest() (request *DescribeCertificateDetailRequest) {
    request = &DescribeCertificateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateDetail")
    
    
    return
}

func NewDescribeCertificateDetailResponse() (response *DescribeCertificateDetailResponse) {
    response = &DescribeCertificateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertificateDetail
// 获取证书详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificateDetail(request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    return c.DescribeCertificateDetailWithContext(context.Background(), request)
}

// DescribeCertificateDetail
// 获取证书详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificateDetailWithContext(ctx context.Context, request *DescribeCertificateDetailRequest) (response *DescribeCertificateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificateOperateLogsRequest() (request *DescribeCertificateOperateLogsRequest) {
    request = &DescribeCertificateOperateLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificateOperateLogs")
    
    
    return
}

func NewDescribeCertificateOperateLogsResponse() (response *DescribeCertificateOperateLogsResponse) {
    response = &DescribeCertificateOperateLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertificateOperateLogs
// 获取用户账号下有关证书的操作日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCertificateOperateLogs(request *DescribeCertificateOperateLogsRequest) (response *DescribeCertificateOperateLogsResponse, err error) {
    return c.DescribeCertificateOperateLogsWithContext(context.Background(), request)
}

// DescribeCertificateOperateLogs
// 获取用户账号下有关证书的操作日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCertificateOperateLogsWithContext(ctx context.Context, request *DescribeCertificateOperateLogsRequest) (response *DescribeCertificateOperateLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCertificateOperateLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificateOperateLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificateOperateLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
    request = &DescribeCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeCertificates")
    
    
    return
}

func NewDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
    response = &DescribeCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertificates
// 本接口（DescribeCertificates）用于获取证书列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    return c.DescribeCertificatesWithContext(context.Background(), request)
}

// DescribeCertificates
// 本接口（DescribeCertificates）用于获取证书列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribeCertificatesWithContext(ctx context.Context, request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployedResourcesRequest() (request *DescribeDeployedResourcesRequest) {
    request = &DescribeDeployedResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeDeployedResources")
    
    
    return
}

func NewDescribeDeployedResourcesResponse() (response *DescribeDeployedResourcesResponse) {
    response = &DescribeDeployedResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeployedResources
// 证书查询关联资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATEIDNUMBERLIMIT = "InvalidParameter.CertificateIdNumberLimit"
//  INVALIDPARAMETER_CONTAINSINVALIDCERTIFICATEID = "InvalidParameter.ContainsInvalidCertificateId"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDeployedResources(request *DescribeDeployedResourcesRequest) (response *DescribeDeployedResourcesResponse, err error) {
    return c.DescribeDeployedResourcesWithContext(context.Background(), request)
}

// DescribeDeployedResources
// 证书查询关联资源
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CERTIFICATEIDNUMBERLIMIT = "InvalidParameter.CertificateIdNumberLimit"
//  INVALIDPARAMETER_CONTAINSINVALIDCERTIFICATEID = "InvalidParameter.ContainsInvalidCertificateId"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDeployedResourcesWithContext(ctx context.Context, request *DescribeDeployedResourcesRequest) (response *DescribeDeployedResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeDeployedResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployedResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeployedResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeManagerDetailRequest() (request *DescribeManagerDetailRequest) {
    request = &DescribeManagerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeManagerDetail")
    
    
    return
}

func NewDescribeManagerDetailResponse() (response *DescribeManagerDetailResponse) {
    response = &DescribeManagerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeManagerDetail
// 查询管理人详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeManagerDetail(request *DescribeManagerDetailRequest) (response *DescribeManagerDetailResponse, err error) {
    return c.DescribeManagerDetailWithContext(context.Background(), request)
}

// DescribeManagerDetail
// 查询管理人详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) DescribeManagerDetailWithContext(ctx context.Context, request *DescribeManagerDetailRequest) (response *DescribeManagerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeManagerDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeManagerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeManagerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeManagersRequest() (request *DescribeManagersRequest) {
    request = &DescribeManagersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribeManagers")
    
    
    return
}

func NewDescribeManagersResponse() (response *DescribeManagersResponse) {
    response = &DescribeManagersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeManagers
// 查询管理人列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeManagers(request *DescribeManagersRequest) (response *DescribeManagersResponse, err error) {
    return c.DescribeManagersWithContext(context.Background(), request)
}

// DescribeManagers
// 查询管理人列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeManagersWithContext(ctx context.Context, request *DescribeManagersRequest) (response *DescribeManagersResponse, err error) {
    if request == nil {
        request = NewDescribeManagersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeManagers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeManagersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackagesRequest() (request *DescribePackagesRequest) {
    request = &DescribePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DescribePackages")
    
    
    return
}

func NewDescribePackagesResponse() (response *DescribePackagesResponse) {
    response = &DescribePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePackages
// 获得权益包列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribePackages(request *DescribePackagesRequest) (response *DescribePackagesResponse, err error) {
    return c.DescribePackagesWithContext(context.Background(), request)
}

// DescribePackages
// 获得权益包列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) DescribePackagesWithContext(ctx context.Context, request *DescribePackagesRequest) (response *DescribePackagesResponse, err error) {
    if request == nil {
        request = NewDescribePackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCertificateRequest() (request *DownloadCertificateRequest) {
    request = &DownloadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "DownloadCertificate")
    
    
    return
}

func NewDownloadCertificateResponse() (response *DownloadCertificateResponse) {
    response = &DownloadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadCertificate
// 本接口（DownloadCertificate）用于下载证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) DownloadCertificate(request *DownloadCertificateRequest) (response *DownloadCertificateResponse, err error) {
    return c.DownloadCertificateWithContext(context.Background(), request)
}

// DownloadCertificate
// 本接口（DownloadCertificate）用于下载证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) DownloadCertificateWithContext(ctx context.Context, request *DownloadCertificateRequest) (response *DownloadCertificateResponse, err error) {
    if request == nil {
        request = NewDownloadCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewHostCertificateRequest() (request *HostCertificateRequest) {
    request = &HostCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "HostCertificate")
    
    
    return
}

func NewHostCertificateResponse() (response *HostCertificateResponse) {
    response = &HostCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// HostCertificate
// 云资源托管
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEHASRENEWED = "FailedOperation.CertificateHasRenewed"
//  FAILEDOPERATION_CERTIFICATEHOSTINGTYPENUMBERLIMIT = "FailedOperation.CertificateHostingTypeNumberLimit"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) HostCertificate(request *HostCertificateRequest) (response *HostCertificateResponse, err error) {
    return c.HostCertificateWithContext(context.Background(), request)
}

// HostCertificate
// 云资源托管
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEHASRENEWED = "FailedOperation.CertificateHasRenewed"
//  FAILEDOPERATION_CERTIFICATEHOSTINGTYPENUMBERLIMIT = "FailedOperation.CertificateHostingTypeNumberLimit"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) HostCertificateWithContext(ctx context.Context, request *HostCertificateRequest) (response *HostCertificateResponse, err error) {
    if request == nil {
        request = NewHostCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HostCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewHostCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateAliasRequest() (request *ModifyCertificateAliasRequest) {
    request = &ModifyCertificateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateAlias")
    
    
    return
}

func NewModifyCertificateAliasResponse() (response *ModifyCertificateAliasResponse) {
    response = &ModifyCertificateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCertificateAlias
// 用户传入证书id和备注来修改证书备注。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCertificateAlias(request *ModifyCertificateAliasRequest) (response *ModifyCertificateAliasResponse, err error) {
    return c.ModifyCertificateAliasWithContext(context.Background(), request)
}

// ModifyCertificateAlias
// 用户传入证书id和备注来修改证书备注。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyCertificateAliasWithContext(ctx context.Context, request *ModifyCertificateAliasRequest) (response *ModifyCertificateAliasResponse, err error) {
    if request == nil {
        request = NewModifyCertificateAliasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateAlias require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificateProjectRequest() (request *ModifyCertificateProjectRequest) {
    request = &ModifyCertificateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificateProject")
    
    
    return
}

func NewModifyCertificateProjectResponse() (response *ModifyCertificateProjectResponse) {
    response = &ModifyCertificateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCertificateProject
// 批量修改证书所属项目。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCertificateProject(request *ModifyCertificateProjectRequest) (response *ModifyCertificateProjectResponse, err error) {
    return c.ModifyCertificateProjectWithContext(context.Background(), request)
}

// ModifyCertificateProject
// 批量修改证书所属项目。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyCertificateProjectWithContext(ctx context.Context, request *ModifyCertificateProjectRequest) (response *ModifyCertificateProjectResponse, err error) {
    if request == nil {
        request = NewModifyCertificateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCertificatesExpiringNotificationSwitchRequest() (request *ModifyCertificatesExpiringNotificationSwitchRequest) {
    request = &ModifyCertificatesExpiringNotificationSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ModifyCertificatesExpiringNotificationSwitch")
    
    
    return
}

func NewModifyCertificatesExpiringNotificationSwitchResponse() (response *ModifyCertificatesExpiringNotificationSwitchResponse) {
    response = &ModifyCertificatesExpiringNotificationSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCertificatesExpiringNotificationSwitch
// 修改忽略证书到期通知。打开或关闭证书到期通知。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"
func (c *Client) ModifyCertificatesExpiringNotificationSwitch(request *ModifyCertificatesExpiringNotificationSwitchRequest) (response *ModifyCertificatesExpiringNotificationSwitchResponse, err error) {
    return c.ModifyCertificatesExpiringNotificationSwitchWithContext(context.Background(), request)
}

// ModifyCertificatesExpiringNotificationSwitch
// 修改忽略证书到期通知。打开或关闭证书到期通知。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"
func (c *Client) ModifyCertificatesExpiringNotificationSwitchWithContext(ctx context.Context, request *ModifyCertificatesExpiringNotificationSwitchRequest) (response *ModifyCertificatesExpiringNotificationSwitchResponse, err error) {
    if request == nil {
        request = NewModifyCertificatesExpiringNotificationSwitchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCertificatesExpiringNotificationSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCertificatesExpiringNotificationSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceCertificateRequest() (request *ReplaceCertificateRequest) {
    request = &ReplaceCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "ReplaceCertificate")
    
    
    return
}

func NewReplaceCertificateResponse() (response *ReplaceCertificateResponse) {
    response = &ReplaceCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReplaceCertificate
// 本接口（ReplaceCertificate）用于重颁发证书。已申请的免费证书仅支持 RSA 算法、密钥对参数为2048的证书重颁发，并且目前仅支持1次重颁发。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) ReplaceCertificate(request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    return c.ReplaceCertificateWithContext(context.Background(), request)
}

// ReplaceCertificate
// 本接口（ReplaceCertificate）用于重颁发证书。已申请的免费证书仅支持 RSA 算法、密钥对参数为2048的证书重颁发，并且目前仅支持1次重颁发。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
func (c *Client) ReplaceCertificateWithContext(ctx context.Context, request *ReplaceCertificateRequest) (response *ReplaceCertificateResponse, err error) {
    if request == nil {
        request = NewReplaceCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeCertificateRequest() (request *RevokeCertificateRequest) {
    request = &RevokeCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "RevokeCertificate")
    
    
    return
}

func NewRevokeCertificateResponse() (response *RevokeCertificateResponse) {
    response = &RevokeCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RevokeCertificate
// 本接口（RevokeCertificate）用于吊销证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_REVOKEFAILED = "FailedOperation.RevokeFailed"
//  FAILEDOPERATION_REVOKERESOURCEFAILED = "FailedOperation.RevokeResourceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) RevokeCertificate(request *RevokeCertificateRequest) (response *RevokeCertificateResponse, err error) {
    return c.RevokeCertificateWithContext(context.Background(), request)
}

// RevokeCertificate
// 本接口（RevokeCertificate）用于吊销证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_REVOKEFAILED = "FailedOperation.RevokeFailed"
//  FAILEDOPERATION_REVOKERESOURCEFAILED = "FailedOperation.RevokeResourceFailed"
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) RevokeCertificateWithContext(ctx context.Context, request *RevokeCertificateRequest) (response *RevokeCertificateResponse, err error) {
    if request == nil {
        request = NewRevokeCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevokeCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitAuditManagerRequest() (request *SubmitAuditManagerRequest) {
    request = &SubmitAuditManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "SubmitAuditManager")
    
    
    return
}

func NewSubmitAuditManagerResponse() (response *SubmitAuditManagerResponse) {
    response = &SubmitAuditManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitAuditManager
// 重新提交审核管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ILLEGALMANAGERSTATUS = "FailedOperation.IllegalManagerStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) SubmitAuditManager(request *SubmitAuditManagerRequest) (response *SubmitAuditManagerResponse, err error) {
    return c.SubmitAuditManagerWithContext(context.Background(), request)
}

// SubmitAuditManager
// 重新提交审核管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ILLEGALMANAGERSTATUS = "FailedOperation.IllegalManagerStatus"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) SubmitAuditManagerWithContext(ctx context.Context, request *SubmitAuditManagerRequest) (response *SubmitAuditManagerResponse, err error) {
    if request == nil {
        request = NewSubmitAuditManagerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitAuditManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitAuditManagerResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitCertificateInformationRequest() (request *SubmitCertificateInformationRequest) {
    request = &SubmitCertificateInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "SubmitCertificateInformation")
    
    
    return
}

func NewSubmitCertificateInformationResponse() (response *SubmitCertificateInformationResponse) {
    response = &SubmitCertificateInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitCertificateInformation
// 提交证书资料。输入参数信息可以分多次提交，但提交的证书资料应最低限度保持完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) SubmitCertificateInformation(request *SubmitCertificateInformationRequest) (response *SubmitCertificateInformationResponse, err error) {
    return c.SubmitCertificateInformationWithContext(context.Background(), request)
}

// SubmitCertificateInformation
// 提交证书资料。输入参数信息可以分多次提交，但提交的证书资料应最低限度保持完整。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  INTERNALERROR = "InternalError"
func (c *Client) SubmitCertificateInformationWithContext(ctx context.Context, request *SubmitCertificateInformationRequest) (response *SubmitCertificateInformationResponse, err error) {
    if request == nil {
        request = NewSubmitCertificateInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitCertificateInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitCertificateInformationResponse()
    err = c.Send(request, response)
    return
}

func NewUploadCertificateRequest() (request *UploadCertificateRequest) {
    request = &UploadCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadCertificate")
    
    
    return
}

func NewUploadCertificateResponse() (response *UploadCertificateResponse) {
    response = &UploadCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadCertificate
// 本接口（UploadCertificate）用于上传证书。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadCertificate(request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    return c.UploadCertificateWithContext(context.Background(), request)
}

// UploadCertificate
// 本接口（UploadCertificate）用于上传证书。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"
//  FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"
//  FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"
//  FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"
//  FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"
//  FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"
//  FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"
//  FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
func (c *Client) UploadCertificateWithContext(ctx context.Context, request *UploadCertificateRequest) (response *UploadCertificateResponse, err error) {
    if request == nil {
        request = NewUploadCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewUploadConfirmLetterRequest() (request *UploadConfirmLetterRequest) {
    request = &UploadConfirmLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadConfirmLetter")
    
    
    return
}

func NewUploadConfirmLetterResponse() (response *UploadConfirmLetterResponse) {
    response = &UploadConfirmLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadConfirmLetter
// 本接口（UploadConfirmLetter）用于上传证书确认函。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  INTERNALERROR = "InternalError"
func (c *Client) UploadConfirmLetter(request *UploadConfirmLetterRequest) (response *UploadConfirmLetterResponse, err error) {
    return c.UploadConfirmLetterWithContext(context.Background(), request)
}

// UploadConfirmLetter
// 本接口（UploadConfirmLetter）用于上传证书确认函。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"
//  FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  INTERNALERROR = "InternalError"
func (c *Client) UploadConfirmLetterWithContext(ctx context.Context, request *UploadConfirmLetterRequest) (response *UploadConfirmLetterResponse, err error) {
    if request == nil {
        request = NewUploadConfirmLetterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadConfirmLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadConfirmLetterResponse()
    err = c.Send(request, response)
    return
}

func NewUploadRevokeLetterRequest() (request *UploadRevokeLetterRequest) {
    request = &UploadRevokeLetterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "UploadRevokeLetter")
    
    
    return
}

func NewUploadRevokeLetterResponse() (response *UploadRevokeLetterResponse) {
    response = &UploadRevokeLetterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadRevokeLetter
// 本接口（UploadRevokeLetter）用于上传证书吊销确认函。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_FILETOOLARGE = "FailedOperation.FileTooLarge"
//  FAILEDOPERATION_FILETOOSMALL = "FailedOperation.FileTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDFILETYPE = "FailedOperation.InvalidFileType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadRevokeLetter(request *UploadRevokeLetterRequest) (response *UploadRevokeLetterResponse, err error) {
    return c.UploadRevokeLetterWithContext(context.Background(), request)
}

// UploadRevokeLetter
// 本接口（UploadRevokeLetter）用于上传证书吊销确认函。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"
//  FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"
//  FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"
//  FAILEDOPERATION_FILETOOLARGE = "FailedOperation.FileTooLarge"
//  FAILEDOPERATION_FILETOOSMALL = "FailedOperation.FileTooSmall"
//  FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"
//  FAILEDOPERATION_INVALIDFILETYPE = "FailedOperation.InvalidFileType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"
//  FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadRevokeLetterWithContext(ctx context.Context, request *UploadRevokeLetterRequest) (response *UploadRevokeLetterResponse, err error) {
    if request == nil {
        request = NewUploadRevokeLetterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadRevokeLetter require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadRevokeLetterResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyManagerRequest() (request *VerifyManagerRequest) {
    request = &VerifyManagerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ssl", APIVersion, "VerifyManager")
    
    
    return
}

func NewVerifyManagerResponse() (response *VerifyManagerResponse) {
    response = &VerifyManagerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyManager
// 重新核验管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) VerifyManager(request *VerifyManagerRequest) (response *VerifyManagerResponse, err error) {
    return c.VerifyManagerWithContext(context.Background(), request)
}

// VerifyManager
// 重新核验管理人
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"
func (c *Client) VerifyManagerWithContext(ctx context.Context, request *VerifyManagerRequest) (response *VerifyManagerResponse, err error) {
    if request == nil {
        request = NewVerifyManagerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyManager require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyManagerResponse()
    err = c.Send(request, response)
    return
}
