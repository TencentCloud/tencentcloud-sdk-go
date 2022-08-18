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

package v20180523

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-05-23"

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


func NewCheckVcodeRequest() (request *CheckVcodeRequest) {
    request = &CheckVcodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "CheckVcode")
    
    
    return
}

func NewCheckVcodeResponse() (response *CheckVcodeResponse) {
    response = &CheckVcodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckVcode
// 检测验证码接口。此接口用于企业电子合同平台通过给用户发送短信验证码，以短信授权方式签署合同。此接口配合发送验证码接口使用。
//
// 
//
// 用户在企业电子合同平台输入收到的验证码后，由企业电子合同平台调用该接口向腾讯云提交确认受托签署合同验证码命令。验证码验证正确时，本次合同签署的授权成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CHECKVCODEERROR = "FailedOperation.CheckVcodeError"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SMSCODEEXPIRED = "FailedOperation.SMSCodeExpired"
//  FAILEDOPERATION_SMSCODELENGTHWRONG = "FailedOperation.SMSCodeLengthWrong"
//  FAILEDOPERATION_WRONGSMSCODE = "FailedOperation.WrongSMSCode"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_CONTRACTPROJECTCODENOTFOUND = "ResourceNotFound.ContractProjectCodeNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) CheckVcode(request *CheckVcodeRequest) (response *CheckVcodeResponse, err error) {
    return c.CheckVcodeWithContext(context.Background(), request)
}

// CheckVcode
// 检测验证码接口。此接口用于企业电子合同平台通过给用户发送短信验证码，以短信授权方式签署合同。此接口配合发送验证码接口使用。
//
// 
//
// 用户在企业电子合同平台输入收到的验证码后，由企业电子合同平台调用该接口向腾讯云提交确认受托签署合同验证码命令。验证码验证正确时，本次合同签署的授权成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CHECKVCODEERROR = "FailedOperation.CheckVcodeError"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SMSCODEEXPIRED = "FailedOperation.SMSCodeExpired"
//  FAILEDOPERATION_SMSCODELENGTHWRONG = "FailedOperation.SMSCodeLengthWrong"
//  FAILEDOPERATION_WRONGSMSCODE = "FailedOperation.WrongSMSCode"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_CONTRACTPROJECTCODENOTFOUND = "ResourceNotFound.ContractProjectCodeNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) CheckVcodeWithContext(ctx context.Context, request *CheckVcodeRequest) (response *CheckVcodeResponse, err error) {
    if request == nil {
        request = NewCheckVcodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckVcode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckVcodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateContractByUploadRequest() (request *CreateContractByUploadRequest) {
    request = &CreateContractByUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "CreateContractByUpload")
    
    
    return
}

func NewCreateContractByUploadResponse() (response *CreateContractByUploadResponse) {
    response = &CreateContractByUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateContractByUpload
// 此接口适用于：客户平台通过上传PDF文件作为合同，以备未来进行签署。接口返回任务号，可调用DescribeTaskStatus接口查看任务执行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHORIZATIONTIMEERROR = "FailedOperation.AuthorizationTimeError"
//  FAILEDOPERATION_CREATECONTRACTERROR = "FailedOperation.CreateContractError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_IPFORMATERROR = "FailedOperation.IPFormatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_CONTRACTFILENAMEERROR = "MissingParameter.ContractFileNameError"
//  MISSINGPARAMETER_CONTRACTFILEPATHERROR = "MissingParameter.ContractFilePathError"
//  MISSINGPARAMETER_LOCATIONNULLERROR = "MissingParameter.LocationNullError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  MISSINGPARAMETER_SIGNERNULLERROR = "MissingParameter.SignerNullError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_INITIATORNOTFOUNDERROR = "ResourceNotFound.InitiatorNotFoundError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_CONTRACTSIGNERUNAVAILABLE = "ResourceUnavailable.ContractSignerUnavailable"
func (c *Client) CreateContractByUpload(request *CreateContractByUploadRequest) (response *CreateContractByUploadResponse, err error) {
    return c.CreateContractByUploadWithContext(context.Background(), request)
}

// CreateContractByUpload
// 此接口适用于：客户平台通过上传PDF文件作为合同，以备未来进行签署。接口返回任务号，可调用DescribeTaskStatus接口查看任务执行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHORIZATIONTIMEERROR = "FailedOperation.AuthorizationTimeError"
//  FAILEDOPERATION_CREATECONTRACTERROR = "FailedOperation.CreateContractError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_IPFORMATERROR = "FailedOperation.IPFormatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_CONTRACTFILENAMEERROR = "MissingParameter.ContractFileNameError"
//  MISSINGPARAMETER_CONTRACTFILEPATHERROR = "MissingParameter.ContractFilePathError"
//  MISSINGPARAMETER_LOCATIONNULLERROR = "MissingParameter.LocationNullError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  MISSINGPARAMETER_SIGNERNULLERROR = "MissingParameter.SignerNullError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_INITIATORNOTFOUNDERROR = "ResourceNotFound.InitiatorNotFoundError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_CONTRACTSIGNERUNAVAILABLE = "ResourceUnavailable.ContractSignerUnavailable"
func (c *Client) CreateContractByUploadWithContext(ctx context.Context, request *CreateContractByUploadRequest) (response *CreateContractByUploadResponse, err error) {
    if request == nil {
        request = NewCreateContractByUploadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateContractByUpload require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateContractByUploadResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEnterpriseAccountRequest() (request *CreateEnterpriseAccountRequest) {
    request = &CreateEnterpriseAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "CreateEnterpriseAccount")
    
    
    return
}

func NewCreateEnterpriseAccountResponse() (response *CreateEnterpriseAccountResponse) {
    response = &CreateEnterpriseAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateEnterpriseAccount
// 为企业电子合同平台的最终企业用户进行开户。在企业电子合同平台进行操作的企业用户，企业电子合同平台向腾讯云发送企业用户的信息，提交开户命令。腾讯云接到请求后，自动为企业电子合同平台的企业用户生成一张数字证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSECONTENTERROR = "FailedOperation.BackendInterfaceResponseContentError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CREATEENTERPRISEACCOUNTERROR = "FailedOperation.CreateEnterpriseAccountError"
//  FAILEDOPERATION_ENTERPRISENAMEFORMATERROR = "FailedOperation.EnterpriseNameFormatError"
//  FAILEDOPERATION_IDENTNOFORMATERROR = "FailedOperation.IdentNoFormatError"
//  FAILEDOPERATION_IDENTTYPEERROR = "FailedOperation.IdentTypeError"
//  FAILEDOPERATION_MESSAGEDATAILLEGAL = "FailedOperation.MessageDataIllegal"
//  FAILEDOPERATION_NAMEISPURENUMBER = "FailedOperation.NameIsPureNumber"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_TRANSACTORNAMEFORMATERROR = "FailedOperation.TransactorNameFormatError"
//  FAILEDOPERATION_TRANSACTORPHONEFORMATERROR = "FailedOperation.TransactorPhoneFormatError"
//  FAILEDOPERATION_WRONGIDENTNOFORMAT = "FailedOperation.WrongIdentNoFormat"
//  FAILEDOPERATION_WRONGIDENTNOSIZE = "FailedOperation.WrongIdentNoSize"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCEINUSE_ACCOUNTEXIST = "ResourceInUse.AccountExist"
//  RESOURCEINUSE_ENTERPRISEACCOUNTALREADYEXIST = "ResourceInUse.EnterpriseAccountAlreadyExist"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) CreateEnterpriseAccount(request *CreateEnterpriseAccountRequest) (response *CreateEnterpriseAccountResponse, err error) {
    return c.CreateEnterpriseAccountWithContext(context.Background(), request)
}

// CreateEnterpriseAccount
// 为企业电子合同平台的最终企业用户进行开户。在企业电子合同平台进行操作的企业用户，企业电子合同平台向腾讯云发送企业用户的信息，提交开户命令。腾讯云接到请求后，自动为企业电子合同平台的企业用户生成一张数字证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSECONTENTERROR = "FailedOperation.BackendInterfaceResponseContentError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CREATEENTERPRISEACCOUNTERROR = "FailedOperation.CreateEnterpriseAccountError"
//  FAILEDOPERATION_ENTERPRISENAMEFORMATERROR = "FailedOperation.EnterpriseNameFormatError"
//  FAILEDOPERATION_IDENTNOFORMATERROR = "FailedOperation.IdentNoFormatError"
//  FAILEDOPERATION_IDENTTYPEERROR = "FailedOperation.IdentTypeError"
//  FAILEDOPERATION_MESSAGEDATAILLEGAL = "FailedOperation.MessageDataIllegal"
//  FAILEDOPERATION_NAMEISPURENUMBER = "FailedOperation.NameIsPureNumber"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_TRANSACTORNAMEFORMATERROR = "FailedOperation.TransactorNameFormatError"
//  FAILEDOPERATION_TRANSACTORPHONEFORMATERROR = "FailedOperation.TransactorPhoneFormatError"
//  FAILEDOPERATION_WRONGIDENTNOFORMAT = "FailedOperation.WrongIdentNoFormat"
//  FAILEDOPERATION_WRONGIDENTNOSIZE = "FailedOperation.WrongIdentNoSize"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCEINUSE_ACCOUNTEXIST = "ResourceInUse.AccountExist"
//  RESOURCEINUSE_ENTERPRISEACCOUNTALREADYEXIST = "ResourceInUse.EnterpriseAccountAlreadyExist"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) CreateEnterpriseAccountWithContext(ctx context.Context, request *CreateEnterpriseAccountRequest) (response *CreateEnterpriseAccountResponse, err error) {
    if request == nil {
        request = NewCreateEnterpriseAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnterpriseAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnterpriseAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonalAccountRequest() (request *CreatePersonalAccountRequest) {
    request = &CreatePersonalAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "CreatePersonalAccount")
    
    
    return
}

func NewCreatePersonalAccountResponse() (response *CreatePersonalAccountResponse) {
    response = &CreatePersonalAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePersonalAccount
// 为企业电子合同平台的最终个人用户进行开户。在企业电子合同平台进行操作的个人用户，企业电子合同平台向腾讯云发送个人用户的信息，提交开户命令。腾讯云接到请求后，自动为企业电子合同平台的个人用户生成一张数字证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSECONTENTERROR = "FailedOperation.BackendInterfaceResponseContentError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CREATEPERSONALACCOUNTERROR = "FailedOperation.CreatePersonalAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_IDENTNOFORMATERROR = "FailedOperation.IdentNoFormatError"
//  FAILEDOPERATION_IDENTTYPEERROR = "FailedOperation.IdentTypeError"
//  FAILEDOPERATION_NAMECONTAINSNUMBER = "FailedOperation.NameContainsNumber"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCEINUSE_ACCOUNTEXIST = "ResourceInUse.AccountExist"
//  RESOURCEINUSE_PERSONACCOUNTALREADYEXIST = "ResourceInUse.PersonAccountAlreadyExist"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) CreatePersonalAccount(request *CreatePersonalAccountRequest) (response *CreatePersonalAccountResponse, err error) {
    return c.CreatePersonalAccountWithContext(context.Background(), request)
}

// CreatePersonalAccount
// 为企业电子合同平台的最终个人用户进行开户。在企业电子合同平台进行操作的个人用户，企业电子合同平台向腾讯云发送个人用户的信息，提交开户命令。腾讯云接到请求后，自动为企业电子合同平台的个人用户生成一张数字证书。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSECONTENTERROR = "FailedOperation.BackendInterfaceResponseContentError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CREATEPERSONALACCOUNTERROR = "FailedOperation.CreatePersonalAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_IDENTNOFORMATERROR = "FailedOperation.IdentNoFormatError"
//  FAILEDOPERATION_IDENTTYPEERROR = "FailedOperation.IdentTypeError"
//  FAILEDOPERATION_NAMECONTAINSNUMBER = "FailedOperation.NameContainsNumber"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCEINUSE_ACCOUNTEXIST = "ResourceInUse.AccountExist"
//  RESOURCEINUSE_PERSONACCOUNTALREADYEXIST = "ResourceInUse.PersonAccountAlreadyExist"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) CreatePersonalAccountWithContext(ctx context.Context, request *CreatePersonalAccountRequest) (response *CreatePersonalAccountResponse, err error) {
    if request == nil {
        request = NewCreatePersonalAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonalAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonalAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSealRequest() (request *CreateSealRequest) {
    request = &CreateSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "CreateSeal")
    
    
    return
}

func NewCreateSealResponse() (response *CreateSealResponse) {
    response = &CreateSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSeal
// 此接口用于客户电子合同平台增加某用户的印章图片。客户平台可以调用此接口增加某用户的印章图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSECONTENTERROR = "FailedOperation.BackendInterfaceResponseContentError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CREATESEALERROR = "FailedOperation.CreateSealError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_DOWNLOADSEALERROR = "FailedOperation.DownloadSealError"
//  FAILEDOPERATION_IMAGENOTPNG = "FailedOperation.ImageNotPNG"
//  FAILEDOPERATION_MESSAGEDATAOVERSIZE = "FailedOperation.MessageDataOverSize"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SEALNUMOVERLIMIT = "FailedOperation.SealNumOverLimit"
//  FAILEDOPERATION_SEALSEXCEED = "FailedOperation.SealsExceed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTUNAVAILABLE = "ResourceUnavailable.AccountUnavailable"
//  RESOURCEUNAVAILABLE_DOWNLOADSEALERROR = "ResourceUnavailable.DownloadSealError"
func (c *Client) CreateSeal(request *CreateSealRequest) (response *CreateSealResponse, err error) {
    return c.CreateSealWithContext(context.Background(), request)
}

// CreateSeal
// 此接口用于客户电子合同平台增加某用户的印章图片。客户平台可以调用此接口增加某用户的印章图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSECONTENTERROR = "FailedOperation.BackendInterfaceResponseContentError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CREATESEALERROR = "FailedOperation.CreateSealError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_DOWNLOADSEALERROR = "FailedOperation.DownloadSealError"
//  FAILEDOPERATION_IMAGENOTPNG = "FailedOperation.ImageNotPNG"
//  FAILEDOPERATION_MESSAGEDATAOVERSIZE = "FailedOperation.MessageDataOverSize"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SEALNUMOVERLIMIT = "FailedOperation.SealNumOverLimit"
//  FAILEDOPERATION_SEALSEXCEED = "FailedOperation.SealsExceed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTUNAVAILABLE = "ResourceUnavailable.AccountUnavailable"
//  RESOURCEUNAVAILABLE_DOWNLOADSEALERROR = "ResourceUnavailable.DownloadSealError"
func (c *Client) CreateSealWithContext(ctx context.Context, request *CreateSealRequest) (response *CreateSealResponse, err error) {
    if request == nil {
        request = NewCreateSealRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSealResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "DeleteAccount")
    
    
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccount
// 删除企业电子合同平台的最终用户。调用该接口后，腾讯云将删除该用户账号。删除账号后，已经签名的合同不受影响。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEACCOUNTERROR = "FailedOperation.DeleteAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_FIRSTENTERPRISEACCOUNTDELETEERROR = "FailedOperation.FirstEnterpriseAccountDeleteError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    return c.DeleteAccountWithContext(context.Background(), request)
}

// DeleteAccount
// 删除企业电子合同平台的最终用户。调用该接口后，腾讯云将删除该用户账号。删除账号后，已经签名的合同不受影响。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEACCOUNTERROR = "FailedOperation.DeleteAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_FIRSTENTERPRISEACCOUNTDELETEERROR = "FailedOperation.FirstEnterpriseAccountDeleteError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSealRequest() (request *DeleteSealRequest) {
    request = &DeleteSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "DeleteSeal")
    
    
    return
}

func NewDeleteSealResponse() (response *DeleteSealResponse) {
    response = &DeleteSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSeal
// 删除印章接口，删除指定账号的某个印章
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_DELETESEALERROR = "FailedOperation.DeleteSealError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_SEALNOTFOUND = "ResourceNotFound.SealNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTUNAVAILABLE = "ResourceUnavailable.AccountUnavailable"
func (c *Client) DeleteSeal(request *DeleteSealRequest) (response *DeleteSealResponse, err error) {
    return c.DeleteSealWithContext(context.Background(), request)
}

// DeleteSeal
// 删除印章接口，删除指定账号的某个印章
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_DELETESEALERROR = "FailedOperation.DeleteSealError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_SEALNOTFOUND = "ResourceNotFound.SealNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTUNAVAILABLE = "ResourceUnavailable.AccountUnavailable"
func (c *Client) DeleteSealWithContext(ctx context.Context, request *DeleteSealRequest) (response *DeleteSealResponse, err error) {
    if request == nil {
        request = NewDeleteSealRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSealResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskStatus
// 接口使用于：客户平台可使用该接口查询任务执行状态或者执行结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// 接口使用于：客户平台可使用该接口查询任务执行状态或者执行结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadContractRequest() (request *DownloadContractRequest) {
    request = &DownloadContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "DownloadContract")
    
    
    return
}

func NewDownloadContractResponse() (response *DownloadContractResponse) {
    response = &DownloadContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadContract
// 下载合同接口。调用该接口可以下载签署中和签署完成的合同。接口返回任务号，可调用DescribeTaskStatus接口查看任务执行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_CONTRACTNOTFOUND = "ResourceNotFound.ContractNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) DownloadContract(request *DownloadContractRequest) (response *DownloadContractResponse, err error) {
    return c.DownloadContractWithContext(context.Background(), request)
}

// DownloadContract
// 下载合同接口。调用该接口可以下载签署中和签署完成的合同。接口返回任务号，可调用DescribeTaskStatus接口查看任务执行结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_CONTRACTNOTFOUND = "ResourceNotFound.ContractNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) DownloadContractWithContext(ctx context.Context, request *DownloadContractRequest) (response *DownloadContractResponse, err error) {
    if request == nil {
        request = NewDownloadContractRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadContract require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadContractResponse()
    err = c.Send(request, response)
    return
}

func NewSendVcodeRequest() (request *SendVcodeRequest) {
    request = &SendVcodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "SendVcode")
    
    
    return
}

func NewSendVcodeResponse() (response *SendVcodeResponse) {
    response = &SendVcodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendVcode
// 发送验证码接口。此接口用于：企业电子合同平台需要腾讯云发送验证码对其用户进行验证时调用，腾讯云将向其用户联系手机(企业电子合同平台为用户开户时通过接口传入)发送验证码，以验证码授权方式签署合同。用户验证工作由企业电子合同平台自身完成。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_COSTACCOUNTERROR = "FailedOperation.CostAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SENDVCODEERROR = "FailedOperation.SendVcodeError"
//  FAILEDOPERATION_SIGNPERMISSIONEXISTED = "FailedOperation.SignPermissionExisted"
//  FAILEDOPERATION_VCODECHECKED = "FailedOperation.VcodeChecked"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_CONTRACTPROJECTCODENOTFOUND = "ResourceNotFound.ContractProjectCodeNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) SendVcode(request *SendVcodeRequest) (response *SendVcodeResponse, err error) {
    return c.SendVcodeWithContext(context.Background(), request)
}

// SendVcode
// 发送验证码接口。此接口用于：企业电子合同平台需要腾讯云发送验证码对其用户进行验证时调用，腾讯云将向其用户联系手机(企业电子合同平台为用户开户时通过接口传入)发送验证码，以验证码授权方式签署合同。用户验证工作由企业电子合同平台自身完成。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_COSTACCOUNTERROR = "FailedOperation.CostAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SENDVCODEERROR = "FailedOperation.SendVcodeError"
//  FAILEDOPERATION_SIGNPERMISSIONEXISTED = "FailedOperation.SignPermissionExisted"
//  FAILEDOPERATION_VCODECHECKED = "FailedOperation.VcodeChecked"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"
//  RESOURCENOTFOUND_CONTRACTPROJECTCODENOTFOUND = "ResourceNotFound.ContractProjectCodeNotFound"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
func (c *Client) SendVcodeWithContext(ctx context.Context, request *SendVcodeRequest) (response *SendVcodeResponse, err error) {
    if request == nil {
        request = NewSendVcodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendVcode require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendVcodeResponse()
    err = c.Send(request, response)
    return
}

func NewSignContractByCoordinateRequest() (request *SignContractByCoordinateRequest) {
    request = &SignContractByCoordinateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "SignContractByCoordinate")
    
    
    return
}

func NewSignContractByCoordinateResponse() (response *SignContractByCoordinateResponse) {
    response = &SignContractByCoordinateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SignContractByCoordinate
// 此接口适用于：客户平台在创建好合同后，由合同签署方对创建的合同内容进行确认，无误后再进行签署。客户平台使用该接口提供详细的PDF文档签名坐标进行签署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_AUTHORIZATIONTIMEERROR = "FailedOperation.AuthorizationTimeError"
//  FAILEDOPERATION_CERTTYPEERROR = "FailedOperation.CertTypeError"
//  FAILEDOPERATION_CONTRACTEXPIRED = "FailedOperation.ContractExpired"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_COORDINATEERROR = "FailedOperation.CoordinateError"
//  FAILEDOPERATION_COORDINATEOUTSIDEPDF = "FailedOperation.CoordinateOutsidePDF"
//  FAILEDOPERATION_COSTACCOUNTERROR = "FailedOperation.CostAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_GETPDFSIZEFAILED = "FailedOperation.GetPDFSizeFailed"
//  FAILEDOPERATION_IPFORMATERROR = "FailedOperation.IPFormatError"
//  FAILEDOPERATION_IMAGENOTBASE = "FailedOperation.ImageNotBase"
//  FAILEDOPERATION_MESSAGEDATAOVERSIZE = "FailedOperation.MessageDataOverSize"
//  FAILEDOPERATION_NOPERMISSIONTOSIGN = "FailedOperation.NoPermissionToSign"
//  FAILEDOPERATION_NOVERIFYERROR = "FailedOperation.NoVerifyError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_REPEATEDCOORDINATE = "FailedOperation.RepeatedCoordinate"
//  FAILEDOPERATION_SEALMISMATCHED = "FailedOperation.SealMismatched"
//  FAILEDOPERATION_SIGNCONTRACTBYCOORDINATEERROR = "FailedOperation.SignContractByCoordinateError"
//  FAILEDOPERATION_SIGNPAGEERROR = "FailedOperation.SignPageError"
//  FAILEDOPERATION_UPDATEFEESTATUSERROR = "FailedOperation.UpdateFeeStatusError"
//  FAILEDOPERATION_WRONGCERTTYPE = "FailedOperation.WrongCertType"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_SUBPLATUNAVAILABLE = "ResourceUnavailable.SubplatUnavailable"
func (c *Client) SignContractByCoordinate(request *SignContractByCoordinateRequest) (response *SignContractByCoordinateResponse, err error) {
    return c.SignContractByCoordinateWithContext(context.Background(), request)
}

// SignContractByCoordinate
// 此接口适用于：客户平台在创建好合同后，由合同签署方对创建的合同内容进行确认，无误后再进行签署。客户平台使用该接口提供详细的PDF文档签名坐标进行签署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_AUTHORIZATIONTIMEERROR = "FailedOperation.AuthorizationTimeError"
//  FAILEDOPERATION_CERTTYPEERROR = "FailedOperation.CertTypeError"
//  FAILEDOPERATION_CONTRACTEXPIRED = "FailedOperation.ContractExpired"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_COORDINATEERROR = "FailedOperation.CoordinateError"
//  FAILEDOPERATION_COORDINATEOUTSIDEPDF = "FailedOperation.CoordinateOutsidePDF"
//  FAILEDOPERATION_COSTACCOUNTERROR = "FailedOperation.CostAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_GETPDFSIZEFAILED = "FailedOperation.GetPDFSizeFailed"
//  FAILEDOPERATION_IPFORMATERROR = "FailedOperation.IPFormatError"
//  FAILEDOPERATION_IMAGENOTBASE = "FailedOperation.ImageNotBase"
//  FAILEDOPERATION_MESSAGEDATAOVERSIZE = "FailedOperation.MessageDataOverSize"
//  FAILEDOPERATION_NOPERMISSIONTOSIGN = "FailedOperation.NoPermissionToSign"
//  FAILEDOPERATION_NOVERIFYERROR = "FailedOperation.NoVerifyError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_REPEATEDCOORDINATE = "FailedOperation.RepeatedCoordinate"
//  FAILEDOPERATION_SEALMISMATCHED = "FailedOperation.SealMismatched"
//  FAILEDOPERATION_SIGNCONTRACTBYCOORDINATEERROR = "FailedOperation.SignContractByCoordinateError"
//  FAILEDOPERATION_SIGNPAGEERROR = "FailedOperation.SignPageError"
//  FAILEDOPERATION_UPDATEFEESTATUSERROR = "FailedOperation.UpdateFeeStatusError"
//  FAILEDOPERATION_WRONGCERTTYPE = "FailedOperation.WrongCertType"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_SUBPLATUNAVAILABLE = "ResourceUnavailable.SubplatUnavailable"
func (c *Client) SignContractByCoordinateWithContext(ctx context.Context, request *SignContractByCoordinateRequest) (response *SignContractByCoordinateResponse, err error) {
    if request == nil {
        request = NewSignContractByCoordinateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SignContractByCoordinate require credential")
    }

    request.SetContext(ctx)
    
    response = NewSignContractByCoordinateResponse()
    err = c.Send(request, response)
    return
}

func NewSignContractByKeywordRequest() (request *SignContractByKeywordRequest) {
    request = &SignContractByKeywordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ds", APIVersion, "SignContractByKeyword")
    
    
    return
}

func NewSignContractByKeywordResponse() (response *SignContractByKeywordResponse) {
    response = &SignContractByKeywordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SignContractByKeyword
// 此接口适用于：客户平台在创建好合同后，由合同签署方对创建的合同内容进行确认，无误后再进行签署。客户平台使用该接口对PDF合同文档按照关键字和坐标进行签署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_AUTHORIZATIONTIMEERROR = "FailedOperation.AuthorizationTimeError"
//  FAILEDOPERATION_CERTTYPEERROR = "FailedOperation.CertTypeError"
//  FAILEDOPERATION_CONTRACTEXPIRED = "FailedOperation.ContractExpired"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_COSTACCOUNTERROR = "FailedOperation.CostAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_IPFORMATERROR = "FailedOperation.IPFormatError"
//  FAILEDOPERATION_IMAGEMEASUREMENTOVERLIMITERROR = "FailedOperation.ImageMeasurementOverLimitError"
//  FAILEDOPERATION_NOPERMISSIONTOSIGN = "FailedOperation.NoPermissionToSign"
//  FAILEDOPERATION_NOVERIFYERROR = "FailedOperation.NoVerifyError"
//  FAILEDOPERATION_OFFSETCOORDOVERLIMITERROR = "FailedOperation.OffsetCoordOverLimitError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SEALMISMATCHED = "FailedOperation.SealMismatched"
//  FAILEDOPERATION_SIGNCONTRACTBYKEYWORDERROR = "FailedOperation.SignContractByKeywordError"
//  FAILEDOPERATION_SIGNFIELDNOTFOUND = "FailedOperation.SignFieldNotFound"
//  FAILEDOPERATION_UPDATEFEESTATUSERROR = "FailedOperation.UpdateFeeStatusError"
//  FAILEDOPERATION_WRONGCERTTYPE = "FailedOperation.WrongCertType"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_IMAGEMEASUREMENTNULLERROR = "MissingParameter.ImageMeasurementNullError"
//  MISSINGPARAMETER_KEYWORDNULLERROR = "MissingParameter.KeywordNullError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  MISSINGPARAMETER_OFFSETCOORDNULLERROR = "MissingParameter.OffsetCoordNullError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_SUBPLATUNAVAILABLE = "ResourceUnavailable.SubplatUnavailable"
func (c *Client) SignContractByKeyword(request *SignContractByKeywordRequest) (response *SignContractByKeywordResponse, err error) {
    return c.SignContractByKeywordWithContext(context.Background(), request)
}

// SignContractByKeyword
// 此接口适用于：客户平台在创建好合同后，由合同签署方对创建的合同内容进行确认，无误后再进行签署。客户平台使用该接口对PDF合同文档按照关键字和坐标进行签署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"
//  FAILEDOPERATION_AUTHORIZATIONTIMEERROR = "FailedOperation.AuthorizationTimeError"
//  FAILEDOPERATION_CERTTYPEERROR = "FailedOperation.CertTypeError"
//  FAILEDOPERATION_CONTRACTEXPIRED = "FailedOperation.ContractExpired"
//  FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"
//  FAILEDOPERATION_COSTACCOUNTERROR = "FailedOperation.CostAccountError"
//  FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"
//  FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"
//  FAILEDOPERATION_IPFORMATERROR = "FailedOperation.IPFormatError"
//  FAILEDOPERATION_IMAGEMEASUREMENTOVERLIMITERROR = "FailedOperation.ImageMeasurementOverLimitError"
//  FAILEDOPERATION_NOPERMISSIONTOSIGN = "FailedOperation.NoPermissionToSign"
//  FAILEDOPERATION_NOVERIFYERROR = "FailedOperation.NoVerifyError"
//  FAILEDOPERATION_OFFSETCOORDOVERLIMITERROR = "FailedOperation.OffsetCoordOverLimitError"
//  FAILEDOPERATION_OTHER = "FailedOperation.Other"
//  FAILEDOPERATION_SEALMISMATCHED = "FailedOperation.SealMismatched"
//  FAILEDOPERATION_SIGNCONTRACTBYKEYWORDERROR = "FailedOperation.SignContractByKeywordError"
//  FAILEDOPERATION_SIGNFIELDNOTFOUND = "FailedOperation.SignFieldNotFound"
//  FAILEDOPERATION_UPDATEFEESTATUSERROR = "FailedOperation.UpdateFeeStatusError"
//  FAILEDOPERATION_WRONGCERTTYPE = "FailedOperation.WrongCertType"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_IMAGEMEASUREMENTNULLERROR = "MissingParameter.ImageMeasurementNullError"
//  MISSINGPARAMETER_KEYWORDNULLERROR = "MissingParameter.KeywordNullError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  MISSINGPARAMETER_OFFSETCOORDNULLERROR = "MissingParameter.OffsetCoordNullError"
//  RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"
//  RESOURCEUNAVAILABLE_SUBPLATUNAVAILABLE = "ResourceUnavailable.SubplatUnavailable"
func (c *Client) SignContractByKeywordWithContext(ctx context.Context, request *SignContractByKeywordRequest) (response *SignContractByKeywordResponse, err error) {
    if request == nil {
        request = NewSignContractByKeywordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SignContractByKeyword require credential")
    }

    request.SetContext(ctx)
    
    response = NewSignContractByKeywordResponse()
    err = c.Send(request, response)
    return
}
