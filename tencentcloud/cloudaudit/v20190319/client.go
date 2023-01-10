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

package v20190319

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-19"

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


func NewCreateAuditRequest() (request *CreateAuditRequest) {
    request = &CreateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateAudit")
    
    
    return
}

func NewCreateAuditResponse() (response *CreateAuditResponse) {
    response = &CreateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAudit
// 参数要求：
//
// 1、如果IsCreateNewBucket的值存在的话，cosRegion和cosBucketName都是必填参数。
//
// 2、如果IsEnableCmqNotify的值是1的话，IsCreateNewQueue、CmqRegion和CmqQueueName都是必填参数。
//
// 3、如果IsEnableCmqNotify的值是0的话，IsCreateNewQueue、CmqRegion和CmqQueueName都不能传。
//
// 4、如果IsEnableKmsEncry的值是1的话，KmsRegion和KeyId属于必填项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEBUCKETFAIL = "FailedOperation.CreateBucketFail"
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_CREATEAUDITERROR = "InternalError.CreateAuditError"
//  INVALIDPARAMETERVALUE_AUDITNAMEERROR = "InvalidParameterValue.AuditNameError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_ISCREATENEWBUCKETERROR = "InvalidParameterValue.IsCreateNewBucketError"
//  INVALIDPARAMETERVALUE_ISCREATENEWQUEUEERROR = "InvalidParameterValue.IsCreateNewQueueError"
//  INVALIDPARAMETERVALUE_ISENABLECMQNOTIFYERROR = "InvalidParameterValue.IsEnableCmqNotifyError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  MISSINGPARAMETER_MISSAUDITNAME = "MissingParameter.MissAuditName"
//  MISSINGPARAMETER_MISSCOSBUCKETNAME = "MissingParameter.MissCosBucketName"
//  MISSINGPARAMETER_MISSCOSREGION = "MissingParameter.MissCosRegion"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDIT = "ResourceInUse.AlreadyExistsSameAudit"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCMQCONFIG = "ResourceInUse.AlreadyExistsSameAuditCmqConfig"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCOSCONFIG = "ResourceInUse.AlreadyExistsSameAuditCosConfig"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
//  RESOURCENOTFOUND_ROLENOTEXIST = "ResourceNotFound.RoleNotExist"
func (c *Client) CreateAudit(request *CreateAuditRequest) (response *CreateAuditResponse, err error) {
    return c.CreateAuditWithContext(context.Background(), request)
}

// CreateAudit
// 参数要求：
//
// 1、如果IsCreateNewBucket的值存在的话，cosRegion和cosBucketName都是必填参数。
//
// 2、如果IsEnableCmqNotify的值是1的话，IsCreateNewQueue、CmqRegion和CmqQueueName都是必填参数。
//
// 3、如果IsEnableCmqNotify的值是0的话，IsCreateNewQueue、CmqRegion和CmqQueueName都不能传。
//
// 4、如果IsEnableKmsEncry的值是1的话，KmsRegion和KeyId属于必填项
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEBUCKETFAIL = "FailedOperation.CreateBucketFail"
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_CREATEAUDITERROR = "InternalError.CreateAuditError"
//  INVALIDPARAMETERVALUE_AUDITNAMEERROR = "InvalidParameterValue.AuditNameError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_ISCREATENEWBUCKETERROR = "InvalidParameterValue.IsCreateNewBucketError"
//  INVALIDPARAMETERVALUE_ISCREATENEWQUEUEERROR = "InvalidParameterValue.IsCreateNewQueueError"
//  INVALIDPARAMETERVALUE_ISENABLECMQNOTIFYERROR = "InvalidParameterValue.IsEnableCmqNotifyError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  MISSINGPARAMETER_MISSAUDITNAME = "MissingParameter.MissAuditName"
//  MISSINGPARAMETER_MISSCOSBUCKETNAME = "MissingParameter.MissCosBucketName"
//  MISSINGPARAMETER_MISSCOSREGION = "MissingParameter.MissCosRegion"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDIT = "ResourceInUse.AlreadyExistsSameAudit"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCMQCONFIG = "ResourceInUse.AlreadyExistsSameAuditCmqConfig"
//  RESOURCEINUSE_ALREADYEXISTSSAMEAUDITCOSCONFIG = "ResourceInUse.AlreadyExistsSameAuditCosConfig"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
//  RESOURCENOTFOUND_ROLENOTEXIST = "ResourceNotFound.RoleNotExist"
func (c *Client) CreateAuditWithContext(ctx context.Context, request *CreateAuditRequest) (response *CreateAuditResponse, err error) {
    if request == nil {
        request = NewCreateAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditTrackRequest() (request *CreateAuditTrackRequest) {
    request = &CreateAuditTrackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateAuditTrack")
    
    
    return
}

func NewCreateAuditTrackResponse() (response *CreateAuditTrackResponse) {
    response = &CreateAuditTrackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAuditTrack
// 创建跟踪集
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKCLSTOPICISEXISTFAILED = "FailedOperation.CheckClsTopicIsExistFailed"
//  FAILEDOPERATION_CHECKCOSBUCKETISEXISTFAILED = "FailedOperation.CheckCosBucketIsExistFailed"
//  FAILEDOPERATION_GETCLSTOPICFAILED = "FailedOperation.GetClsTopicFailed"
//  FAILEDOPERATION_GETCOSBUCKETLISTFAILED = "FailedOperation.GetCosBucketListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAuditTrack(request *CreateAuditTrackRequest) (response *CreateAuditTrackResponse, err error) {
    return c.CreateAuditTrackWithContext(context.Background(), request)
}

// CreateAuditTrack
// 创建跟踪集
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CHECKCLSTOPICISEXISTFAILED = "FailedOperation.CheckClsTopicIsExistFailed"
//  FAILEDOPERATION_CHECKCOSBUCKETISEXISTFAILED = "FailedOperation.CheckCosBucketIsExistFailed"
//  FAILEDOPERATION_GETCLSTOPICFAILED = "FailedOperation.GetClsTopicFailed"
//  FAILEDOPERATION_GETCOSBUCKETLISTFAILED = "FailedOperation.GetCosBucketListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAuditTrackWithContext(ctx context.Context, request *CreateAuditTrackRequest) (response *CreateAuditTrackResponse, err error) {
    if request == nil {
        request = NewCreateAuditTrackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditTrack require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditTrackResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRequest() (request *DeleteAuditRequest) {
    request = &DeleteAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteAudit")
    
    
    return
}

func NewDeleteAuditResponse() (response *DeleteAuditResponse) {
    response = &DeleteAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAudit
// 删除跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEAUDITERROR = "InternalError.DeleteAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DeleteAudit(request *DeleteAuditRequest) (response *DeleteAuditResponse, err error) {
    return c.DeleteAuditWithContext(context.Background(), request)
}

// DeleteAudit
// 删除跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR_DELETEAUDITERROR = "InternalError.DeleteAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DeleteAuditWithContext(ctx context.Context, request *DeleteAuditRequest) (response *DeleteAuditResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditTrackRequest() (request *DeleteAuditTrackRequest) {
    request = &DeleteAuditTrackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteAuditTrack")
    
    
    return
}

func NewDeleteAuditTrackResponse() (response *DeleteAuditTrackResponse) {
    response = &DeleteAuditTrackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAuditTrack
// 删除云审计跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DeleteAuditTrack(request *DeleteAuditTrackRequest) (response *DeleteAuditTrackResponse, err error) {
    return c.DeleteAuditTrackWithContext(context.Background(), request)
}

// DeleteAuditTrack
// 删除云审计跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DeleteAuditTrackWithContext(ctx context.Context, request *DeleteAuditTrackRequest) (response *DeleteAuditTrackResponse, err error) {
    if request == nil {
        request = NewDeleteAuditTrackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditTrack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditTrackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRequest() (request *DescribeAuditRequest) {
    request = &DescribeAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAudit")
    
    
    return
}

func NewDescribeAuditResponse() (response *DescribeAuditResponse) {
    response = &DescribeAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAudit
// 查询跟踪集详情
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBEAUDITERROR = "InternalError.DescribeAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DescribeAudit(request *DescribeAuditRequest) (response *DescribeAuditResponse, err error) {
    return c.DescribeAuditWithContext(context.Background(), request)
}

// DescribeAudit
// 查询跟踪集详情
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBEAUDITERROR = "InternalError.DescribeAuditError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DescribeAuditWithContext(ctx context.Context, request *DescribeAuditRequest) (response *DescribeAuditResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditTrackRequest() (request *DescribeAuditTrackRequest) {
    request = &DescribeAuditTrackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAuditTrack")
    
    
    return
}

func NewDescribeAuditTrackResponse() (response *DescribeAuditTrackResponse) {
    response = &DescribeAuditTrackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditTrack
// 查询云审计跟踪集详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DescribeAuditTrack(request *DescribeAuditTrackRequest) (response *DescribeAuditTrackResponse, err error) {
    return c.DescribeAuditTrackWithContext(context.Background(), request)
}

// DescribeAuditTrack
// 查询云审计跟踪集详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) DescribeAuditTrackWithContext(ctx context.Context, request *DescribeAuditTrackRequest) (response *DescribeAuditTrackResponse, err error) {
    if request == nil {
        request = NewDescribeAuditTrackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditTrack require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditTrackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditTracksRequest() (request *DescribeAuditTracksRequest) {
    request = &DescribeAuditTracksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAuditTracks")
    
    
    return
}

func NewDescribeAuditTracksResponse() (response *DescribeAuditTracksResponse) {
    response = &DescribeAuditTracksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditTracks
// 查询云审计跟踪集列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeAuditTracks(request *DescribeAuditTracksRequest) (response *DescribeAuditTracksResponse, err error) {
    return c.DescribeAuditTracksWithContext(context.Background(), request)
}

// DescribeAuditTracks
// 查询云审计跟踪集列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeAuditTracksWithContext(ctx context.Context, request *DescribeAuditTracksRequest) (response *DescribeAuditTracksResponse, err error) {
    if request == nil {
        request = NewDescribeAuditTracksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditTracks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditTracksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeEvents")
    
    
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEvents
// 查询云审计日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    return c.DescribeEventsWithContext(context.Background(), request)
}

// DescribeEvents
// 查询云审计日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
func (c *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttributeKeyRequest() (request *GetAttributeKeyRequest) {
    request = &GetAttributeKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "GetAttributeKey")
    
    
    return
}

func NewGetAttributeKeyResponse() (response *GetAttributeKeyResponse) {
    response = &GetAttributeKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAttributeKey
// 查询AttributeKey的有效取值范围
//
// 可能返回的错误码:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
func (c *Client) GetAttributeKey(request *GetAttributeKeyRequest) (response *GetAttributeKeyResponse, err error) {
    return c.GetAttributeKeyWithContext(context.Background(), request)
}

// GetAttributeKey
// 查询AttributeKey的有效取值范围
//
// 可能返回的错误码:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
func (c *Client) GetAttributeKeyWithContext(ctx context.Context, request *GetAttributeKeyRequest) (response *GetAttributeKeyResponse, err error) {
    if request == nil {
        request = NewGetAttributeKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAttributeKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAttributeKeyResponse()
    err = c.Send(request, response)
    return
}

func NewInquireAuditCreditRequest() (request *InquireAuditCreditRequest) {
    request = &InquireAuditCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "InquireAuditCredit")
    
    
    return
}

func NewInquireAuditCreditResponse() (response *InquireAuditCreditResponse) {
    response = &InquireAuditCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquireAuditCredit
// 查询用户可创建跟踪集的数量
//
// 可能返回的错误码:
//  INTERNALERROR_INQUIREAUDITCREDITERROR = "InternalError.InquireAuditCreditError"
func (c *Client) InquireAuditCredit(request *InquireAuditCreditRequest) (response *InquireAuditCreditResponse, err error) {
    return c.InquireAuditCreditWithContext(context.Background(), request)
}

// InquireAuditCredit
// 查询用户可创建跟踪集的数量
//
// 可能返回的错误码:
//  INTERNALERROR_INQUIREAUDITCREDITERROR = "InternalError.InquireAuditCreditError"
func (c *Client) InquireAuditCreditWithContext(ctx context.Context, request *InquireAuditCreditRequest) (response *InquireAuditCreditResponse, err error) {
    if request == nil {
        request = NewInquireAuditCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquireAuditCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquireAuditCreditResponse()
    err = c.Send(request, response)
    return
}

func NewListAuditsRequest() (request *ListAuditsRequest) {
    request = &ListAuditsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListAudits")
    
    
    return
}

func NewListAuditsResponse() (response *ListAuditsResponse) {
    response = &ListAuditsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAudits
// 查询跟踪集概要
//
// 可能返回的错误码:
//  INTERNALERROR_LISTAUDITSERROR = "InternalError.ListAuditsError"
func (c *Client) ListAudits(request *ListAuditsRequest) (response *ListAuditsResponse, err error) {
    return c.ListAuditsWithContext(context.Background(), request)
}

// ListAudits
// 查询跟踪集概要
//
// 可能返回的错误码:
//  INTERNALERROR_LISTAUDITSERROR = "InternalError.ListAuditsError"
func (c *Client) ListAuditsWithContext(ctx context.Context, request *ListAuditsRequest) (response *ListAuditsResponse, err error) {
    if request == nil {
        request = NewListAuditsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAudits require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAuditsResponse()
    err = c.Send(request, response)
    return
}

func NewListCmqEnableRegionRequest() (request *ListCmqEnableRegionRequest) {
    request = &ListCmqEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCmqEnableRegion")
    
    
    return
}

func NewListCmqEnableRegionResponse() (response *ListCmqEnableRegionResponse) {
    response = &ListCmqEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListCmqEnableRegion
// 查询云审计支持的cmq的可用区
//
// 可能返回的错误码:
//  INTERNALERROR_LISTCMQENABLEREGIONERROR = "InternalError.ListCmqEnableRegionError"
func (c *Client) ListCmqEnableRegion(request *ListCmqEnableRegionRequest) (response *ListCmqEnableRegionResponse, err error) {
    return c.ListCmqEnableRegionWithContext(context.Background(), request)
}

// ListCmqEnableRegion
// 查询云审计支持的cmq的可用区
//
// 可能返回的错误码:
//  INTERNALERROR_LISTCMQENABLEREGIONERROR = "InternalError.ListCmqEnableRegionError"
func (c *Client) ListCmqEnableRegionWithContext(ctx context.Context, request *ListCmqEnableRegionRequest) (response *ListCmqEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCmqEnableRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCmqEnableRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCmqEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewListCosEnableRegionRequest() (request *ListCosEnableRegionRequest) {
    request = &ListCosEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCosEnableRegion")
    
    
    return
}

func NewListCosEnableRegionResponse() (response *ListCosEnableRegionResponse) {
    response = &ListCosEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListCosEnableRegion
// 查询云审计支持的cos可用区
//
// 可能返回的错误码:
//  INTERNALERROR_LISTCOSENABLEREGIONERROR = "InternalError.ListCosEnableRegionError"
func (c *Client) ListCosEnableRegion(request *ListCosEnableRegionRequest) (response *ListCosEnableRegionResponse, err error) {
    return c.ListCosEnableRegionWithContext(context.Background(), request)
}

// ListCosEnableRegion
// 查询云审计支持的cos可用区
//
// 可能返回的错误码:
//  INTERNALERROR_LISTCOSENABLEREGIONERROR = "InternalError.ListCosEnableRegionError"
func (c *Client) ListCosEnableRegionWithContext(ctx context.Context, request *ListCosEnableRegionRequest) (response *ListCosEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCosEnableRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListCosEnableRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListCosEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewListKeyAliasByRegionRequest() (request *ListKeyAliasByRegionRequest) {
    request = &ListKeyAliasByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListKeyAliasByRegion")
    
    
    return
}

func NewListKeyAliasByRegionResponse() (response *ListKeyAliasByRegionResponse) {
    response = &ListKeyAliasByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListKeyAliasByRegion
// 根据地域获取KMS密钥别名
//
// 可能返回的错误码:
//  INTERNALERROR_LISTKEYALIASBYREGIONERROR = "InternalError.ListKeyAliasByRegionError"
//  INVALIDPARAMETERVALUE_KMSREGIONERROR = "InvalidParameterValue.KmsRegionError"
//  RESOURCENOTFOUND_ROLENOTEXIST = "ResourceNotFound.RoleNotExist"
func (c *Client) ListKeyAliasByRegion(request *ListKeyAliasByRegionRequest) (response *ListKeyAliasByRegionResponse, err error) {
    return c.ListKeyAliasByRegionWithContext(context.Background(), request)
}

// ListKeyAliasByRegion
// 根据地域获取KMS密钥别名
//
// 可能返回的错误码:
//  INTERNALERROR_LISTKEYALIASBYREGIONERROR = "InternalError.ListKeyAliasByRegionError"
//  INVALIDPARAMETERVALUE_KMSREGIONERROR = "InvalidParameterValue.KmsRegionError"
//  RESOURCENOTFOUND_ROLENOTEXIST = "ResourceNotFound.RoleNotExist"
func (c *Client) ListKeyAliasByRegionWithContext(ctx context.Context, request *ListKeyAliasByRegionRequest) (response *ListKeyAliasByRegionResponse, err error) {
    if request == nil {
        request = NewListKeyAliasByRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListKeyAliasByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListKeyAliasByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewLookUpEventsRequest() (request *LookUpEventsRequest) {
    request = &LookUpEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "LookUpEvents")
    
    
    return
}

func NewLookUpEventsResponse() (response *LookUpEventsResponse) {
    response = &LookUpEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LookUpEvents
// 用于对操作日志进行检索，便于用户进行查询相关的操作信息。
//
// 可能返回的错误码:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER_TIME = "InvalidParameter.Time"
//  INVALIDPARAMETERVALUE_MAXRESULT = "InvalidParameterValue.MaxResult"
//  INVALIDPARAMETERVALUE_TIME = "InvalidParameterValue.Time"
//  INVALIDPARAMETERVALUE_ATTRIBUTEKEY = "InvalidParameterValue.attributeKey"
//  LIMITEXCEEDED_OVERTIME = "LimitExceeded.OverTime"
func (c *Client) LookUpEvents(request *LookUpEventsRequest) (response *LookUpEventsResponse, err error) {
    return c.LookUpEventsWithContext(context.Background(), request)
}

// LookUpEvents
// 用于对操作日志进行检索，便于用户进行查询相关的操作信息。
//
// 可能返回的错误码:
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER_TIME = "InvalidParameter.Time"
//  INVALIDPARAMETERVALUE_MAXRESULT = "InvalidParameterValue.MaxResult"
//  INVALIDPARAMETERVALUE_TIME = "InvalidParameterValue.Time"
//  INVALIDPARAMETERVALUE_ATTRIBUTEKEY = "InvalidParameterValue.attributeKey"
//  LIMITEXCEEDED_OVERTIME = "LimitExceeded.OverTime"
func (c *Client) LookUpEventsWithContext(ctx context.Context, request *LookUpEventsRequest) (response *LookUpEventsResponse, err error) {
    if request == nil {
        request = NewLookUpEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LookUpEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewLookUpEventsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditTrackRequest() (request *ModifyAuditTrackRequest) {
    request = &ModifyAuditTrackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ModifyAuditTrack")
    
    
    return
}

func NewModifyAuditTrackResponse() (response *ModifyAuditTrackResponse) {
    response = &ModifyAuditTrackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAuditTrack
// 修改云审计跟踪
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKCLSTOPICISEXISTFAILED = "FailedOperation.CheckClsTopicIsExistFailed"
//  FAILEDOPERATION_CHECKCOSBUCKETISEXISTFAILED = "FailedOperation.CheckCosBucketIsExistFailed"
//  FAILEDOPERATION_GETCLSTOPICFAILED = "FailedOperation.GetClsTopicFailed"
//  FAILEDOPERATION_GETCOSBUCKETLISTFAILED = "FailedOperation.GetCosBucketListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_AUDITTRACKNAMENOTSUPPORTMODIFY = "InvalidParameterValue.AuditTrackNameNotSupportModify"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) ModifyAuditTrack(request *ModifyAuditTrackRequest) (response *ModifyAuditTrackResponse, err error) {
    return c.ModifyAuditTrackWithContext(context.Background(), request)
}

// ModifyAuditTrack
// 修改云审计跟踪
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKCLSTOPICISEXISTFAILED = "FailedOperation.CheckClsTopicIsExistFailed"
//  FAILEDOPERATION_CHECKCOSBUCKETISEXISTFAILED = "FailedOperation.CheckCosBucketIsExistFailed"
//  FAILEDOPERATION_GETCLSTOPICFAILED = "FailedOperation.GetClsTopicFailed"
//  FAILEDOPERATION_GETCOSBUCKETLISTFAILED = "FailedOperation.GetCosBucketListFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SEARCHERROR = "InternalError.SearchError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_AUDITTRACKNAMENOTSUPPORTMODIFY = "InvalidParameterValue.AuditTrackNameNotSupportModify"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) ModifyAuditTrackWithContext(ctx context.Context, request *ModifyAuditTrackRequest) (response *ModifyAuditTrackResponse, err error) {
    if request == nil {
        request = NewModifyAuditTrackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditTrack require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditTrackResponse()
    err = c.Send(request, response)
    return
}

func NewStartLoggingRequest() (request *StartLoggingRequest) {
    request = &StartLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StartLogging")
    
    
    return
}

func NewStartLoggingResponse() (response *StartLoggingResponse) {
    response = &StartLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartLogging
// 开启跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR_STARTLOGGINGERROR = "InternalError.StartLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StartLogging(request *StartLoggingRequest) (response *StartLoggingResponse, err error) {
    return c.StartLoggingWithContext(context.Background(), request)
}

// StartLogging
// 开启跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR_STARTLOGGINGERROR = "InternalError.StartLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StartLoggingWithContext(ctx context.Context, request *StartLoggingRequest) (response *StartLoggingResponse, err error) {
    if request == nil {
        request = NewStartLoggingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartLogging require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewStopLoggingRequest() (request *StopLoggingRequest) {
    request = &StopLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StopLogging")
    
    
    return
}

func NewStopLoggingResponse() (response *StopLoggingResponse) {
    response = &StopLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopLogging
// 关闭跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR_STOPLOGGINGERROR = "InternalError.StopLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StopLogging(request *StopLoggingRequest) (response *StopLoggingResponse, err error) {
    return c.StopLoggingWithContext(context.Background(), request)
}

// StopLogging
// 关闭跟踪集
//
// 可能返回的错误码:
//  INTERNALERROR_STOPLOGGINGERROR = "InternalError.StopLoggingError"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) StopLoggingWithContext(ctx context.Context, request *StopLoggingRequest) (response *StopLoggingResponse, err error) {
    if request == nil {
        request = NewStopLoggingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLogging require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAuditRequest() (request *UpdateAuditRequest) {
    request = &UpdateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudaudit", APIVersion, "UpdateAudit")
    
    
    return
}

func NewUpdateAuditResponse() (response *UpdateAuditResponse) {
    response = &UpdateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAudit
// 参数要求：
//
// 1、如果IsCreateNewBucket的值存在的话，cosRegion和cosBucketName都是必填参数。
//
// 2、如果IsEnableCmqNotify的值是1的话，IsCreateNewQueue、CmqRegion和CmqQueueName都是必填参数。
//
// 3、如果IsEnableCmqNotify的值是0的话，IsCreateNewQueue、CmqRegion和CmqQueueName都不能传。
//
// 4、如果IsEnableKmsEncry的值是1的话，KmsRegion和KeyId属于必填项
//
// 可能返回的错误码:
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_UPDATEAUDITERROR = "InternalError.UpdateAuditError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) UpdateAudit(request *UpdateAuditRequest) (response *UpdateAuditResponse, err error) {
    return c.UpdateAuditWithContext(context.Background(), request)
}

// UpdateAudit
// 参数要求：
//
// 1、如果IsCreateNewBucket的值存在的话，cosRegion和cosBucketName都是必填参数。
//
// 2、如果IsEnableCmqNotify的值是1的话，IsCreateNewQueue、CmqRegion和CmqQueueName都是必填参数。
//
// 3、如果IsEnableCmqNotify的值是0的话，IsCreateNewQueue、CmqRegion和CmqQueueName都不能传。
//
// 4、如果IsEnableKmsEncry的值是1的话，KmsRegion和KeyId属于必填项
//
// 可能返回的错误码:
//  INTERNALERROR_CMQERROR = "InternalError.CmqError"
//  INTERNALERROR_UPDATEAUDITERROR = "InternalError.UpdateAuditError"
//  INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"
//  INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"
//  INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"
//  INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"
//  INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"
//  INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"
//  MISSINGPARAMETER_CMQ = "MissingParameter.cmq"
//  RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"
//  RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"
func (c *Client) UpdateAuditWithContext(ctx context.Context, request *UpdateAuditRequest) (response *UpdateAuditResponse, err error) {
    if request == nil {
        request = NewUpdateAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAuditResponse()
    err = c.Send(request, response)
    return
}
