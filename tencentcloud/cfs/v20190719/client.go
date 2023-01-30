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

package v20190719

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-19"

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


func NewBindAutoSnapshotPolicyRequest() (request *BindAutoSnapshotPolicyRequest) {
    request = &BindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "BindAutoSnapshotPolicy")
    
    
    return
}

func NewBindAutoSnapshotPolicyResponse() (response *BindAutoSnapshotPolicyResponse) {
    response = &BindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindAutoSnapshotPolicy
// 文件系统绑定快照策略，可以同时绑定多个fs，一个fs 只能跟一个策略绑定
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) BindAutoSnapshotPolicy(request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
    return c.BindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// BindAutoSnapshotPolicy
// 文件系统绑定快照策略，可以同时绑定多个fs，一个fs 只能跟一个策略绑定
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) BindAutoSnapshotPolicyWithContext(ctx context.Context, request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewBindAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
    request = &CreateAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateAutoSnapshotPolicy")
    
    
    return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
    response = &CreateAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoSnapshotPolicy
// 创建定期快照策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDALIVEDDAYS = "InvalidParameter.InvalidAlivedDays"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
    return c.CreateAutoSnapshotPolicyWithContext(context.Background(), request)
}

// CreateAutoSnapshotPolicy
// 创建定期快照策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDALIVEDDAYS = "InvalidParameter.InvalidAlivedDays"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateAutoSnapshotPolicyWithContext(ctx context.Context, request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCfsFileSystemRequest() (request *CreateCfsFileSystemRequest) {
    request = &CreateCfsFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsFileSystem")
    
    
    return
}

func NewCreateCfsFileSystemResponse() (response *CreateCfsFileSystemResponse) {
    response = &CreateCfsFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCfsFileSystem
// 用于添加新文件系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDRESOURCEPKGFAILED = "FailedOperation.BindResourcePkgFailed"
//  FAILEDOPERATION_CLIENTTOKENINUSE = "FailedOperation.ClientTokenInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFSFAILED = "InternalError.CreateFsFailed"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_CLIENTTOKENLIMITEXCEEDED = "InvalidParameterValue.ClientTokenLimitExceeded"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDCLIENTTOKEN = "InvalidParameterValue.InvalidClientToken"
//  INVALIDPARAMETERVALUE_INVALIDENCRYPTED = "InvalidParameterValue.InvalidEncrypted"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTTARGETIP = "InvalidParameterValue.InvalidMountTargetIp"
//  INVALIDPARAMETERVALUE_INVALIDNETINTERFACE = "InvalidParameterValue.InvalidNetInterface"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCETAGS = "InvalidParameterValue.InvalidResourceTags"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGERESOURCEPKG = "InvalidParameterValue.InvalidStorageResourcePkg"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGETYPE = "InvalidParameterValue.InvalidStorageType"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_INVALIDTAGKEY = "InvalidParameterValue.InvalidTagKey"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_INVALIDTURBOCAPACITY = "InvalidParameterValue.InvalidTurboCapacity"
//  INVALIDPARAMETERVALUE_INVALIDVIP = "InvalidParameterValue.InvalidVip"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  INVALIDPARAMETERVALUE_INVALIDVPCPARAMETER = "InvalidParameterValue.InvalidVpcParameter"
//  INVALIDPARAMETERVALUE_INVALIDZONEID = "InvalidParameterValue.InvalidZoneId"
//  INVALIDPARAMETERVALUE_INVALIDZONEORZONEID = "InvalidParameterValue.InvalidZoneOrZoneId"
//  INVALIDPARAMETERVALUE_MISSINGKMSKEYID = "InvalidParameterValue.MissingKmsKeyId"
//  INVALIDPARAMETERVALUE_MISSINGSTORAGERESOURCEPKG = "InvalidParameterValue.MissingStorageResourcePkg"
//  INVALIDPARAMETERVALUE_MISSINGSUBNETIDORUNSUBNETID = "InvalidParameterValue.MissingSubnetidOrUnsubnetid"
//  INVALIDPARAMETERVALUE_MISSINGVPCPARAMETER = "InvalidParameterValue.MissingVpcParameter"
//  INVALIDPARAMETERVALUE_MISSINGVPCIDORUNVPCID = "InvalidParameterValue.MissingVpcidOrUnvpcid"
//  INVALIDPARAMETERVALUE_MISSINGZONEID = "InvalidParameterValue.MissingZoneId"
//  INVALIDPARAMETERVALUE_MISSINGZONEORZONEID = "InvalidParameterValue.MissingZoneOrZoneId"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_ZONEIDREGIONNOTMATCH = "InvalidParameterValue.ZoneIdRegionNotMatch"
//  RESOURCEINSUFFICIENT_FILESYSTEMLIMITEXCEEDED = "ResourceInsufficient.FileSystemLimitExceeded"
//  RESOURCEINSUFFICIENT_REGIONSOLDOUT = "ResourceInsufficient.RegionSoldOut"
//  RESOURCEINSUFFICIENT_SUBNETIPALLOCCUPIED = "ResourceInsufficient.SubnetIpAllOccupied"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICNETINTERFACENOTSUPPORTED = "UnsupportedOperation.BasicNetInterfaceNotSupported"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsFileSystem(request *CreateCfsFileSystemRequest) (response *CreateCfsFileSystemResponse, err error) {
    return c.CreateCfsFileSystemWithContext(context.Background(), request)
}

// CreateCfsFileSystem
// 用于添加新文件系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDRESOURCEPKGFAILED = "FailedOperation.BindResourcePkgFailed"
//  FAILEDOPERATION_CLIENTTOKENINUSE = "FailedOperation.ClientTokenInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFSFAILED = "InternalError.CreateFsFailed"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_CLIENTTOKENLIMITEXCEEDED = "InvalidParameterValue.ClientTokenLimitExceeded"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDCLIENTTOKEN = "InvalidParameterValue.InvalidClientToken"
//  INVALIDPARAMETERVALUE_INVALIDENCRYPTED = "InvalidParameterValue.InvalidEncrypted"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTTARGETIP = "InvalidParameterValue.InvalidMountTargetIp"
//  INVALIDPARAMETERVALUE_INVALIDNETINTERFACE = "InvalidParameterValue.InvalidNetInterface"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCETAGS = "InvalidParameterValue.InvalidResourceTags"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGERESOURCEPKG = "InvalidParameterValue.InvalidStorageResourcePkg"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGETYPE = "InvalidParameterValue.InvalidStorageType"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_INVALIDTAGKEY = "InvalidParameterValue.InvalidTagKey"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_INVALIDTURBOCAPACITY = "InvalidParameterValue.InvalidTurboCapacity"
//  INVALIDPARAMETERVALUE_INVALIDVIP = "InvalidParameterValue.InvalidVip"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  INVALIDPARAMETERVALUE_INVALIDVPCPARAMETER = "InvalidParameterValue.InvalidVpcParameter"
//  INVALIDPARAMETERVALUE_INVALIDZONEID = "InvalidParameterValue.InvalidZoneId"
//  INVALIDPARAMETERVALUE_INVALIDZONEORZONEID = "InvalidParameterValue.InvalidZoneOrZoneId"
//  INVALIDPARAMETERVALUE_MISSINGKMSKEYID = "InvalidParameterValue.MissingKmsKeyId"
//  INVALIDPARAMETERVALUE_MISSINGSTORAGERESOURCEPKG = "InvalidParameterValue.MissingStorageResourcePkg"
//  INVALIDPARAMETERVALUE_MISSINGSUBNETIDORUNSUBNETID = "InvalidParameterValue.MissingSubnetidOrUnsubnetid"
//  INVALIDPARAMETERVALUE_MISSINGVPCPARAMETER = "InvalidParameterValue.MissingVpcParameter"
//  INVALIDPARAMETERVALUE_MISSINGVPCIDORUNVPCID = "InvalidParameterValue.MissingVpcidOrUnvpcid"
//  INVALIDPARAMETERVALUE_MISSINGZONEID = "InvalidParameterValue.MissingZoneId"
//  INVALIDPARAMETERVALUE_MISSINGZONEORZONEID = "InvalidParameterValue.MissingZoneOrZoneId"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_ZONEIDREGIONNOTMATCH = "InvalidParameterValue.ZoneIdRegionNotMatch"
//  RESOURCEINSUFFICIENT_FILESYSTEMLIMITEXCEEDED = "ResourceInsufficient.FileSystemLimitExceeded"
//  RESOURCEINSUFFICIENT_REGIONSOLDOUT = "ResourceInsufficient.RegionSoldOut"
//  RESOURCEINSUFFICIENT_SUBNETIPALLOCCUPIED = "ResourceInsufficient.SubnetIpAllOccupied"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICNETINTERFACENOTSUPPORTED = "UnsupportedOperation.BasicNetInterfaceNotSupported"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsFileSystemWithContext(ctx context.Context, request *CreateCfsFileSystemRequest) (response *CreateCfsFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateCfsFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCfsPGroupRequest() (request *CreateCfsPGroupRequest) {
    request = &CreateCfsPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsPGroup")
    
    
    return
}

func NewCreateCfsPGroupResponse() (response *CreateCfsPGroupResponse) {
    response = &CreateCfsPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCfsPGroup
// 本接口（CreateCfsPGroup）用于创建权限组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCEINSUFFICIENT_PGROUPNUMBERLIMITEXCEEDED = "ResourceInsufficient.PgroupNumberLimitExceeded"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsPGroup(request *CreateCfsPGroupRequest) (response *CreateCfsPGroupResponse, err error) {
    return c.CreateCfsPGroupWithContext(context.Background(), request)
}

// CreateCfsPGroup
// 本接口（CreateCfsPGroup）用于创建权限组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCEINSUFFICIENT_PGROUPNUMBERLIMITEXCEEDED = "ResourceInsufficient.PgroupNumberLimitExceeded"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsPGroupWithContext(ctx context.Context, request *CreateCfsPGroupRequest) (response *CreateCfsPGroupResponse, err error) {
    if request == nil {
        request = NewCreateCfsPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCfsRuleRequest() (request *CreateCfsRuleRequest) {
    request = &CreateCfsRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsRule")
    
    
    return
}

func NewCreateCfsRuleResponse() (response *CreateCfsRuleResponse) {
    response = &CreateCfsRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCfsRule
// 本接口（CreateCfsRule）用于创建权限组规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCEINSUFFICIENT_RULELIMITEXCEEDED = "ResourceInsufficient.RuleLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsRule(request *CreateCfsRuleRequest) (response *CreateCfsRuleResponse, err error) {
    return c.CreateCfsRuleWithContext(context.Background(), request)
}

// CreateCfsRule
// 本接口（CreateCfsRule）用于创建权限组规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCEINSUFFICIENT_RULELIMITEXCEEDED = "ResourceInsufficient.RuleLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsRuleWithContext(ctx context.Context, request *CreateCfsRuleRequest) (response *CreateCfsRuleResponse, err error) {
    if request == nil {
        request = NewCreateCfsRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCfsSnapshotRequest() (request *CreateCfsSnapshotRequest) {
    request = &CreateCfsSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsSnapshot")
    
    
    return
}

func NewCreateCfsSnapshotResponse() (response *CreateCfsSnapshotResponse) {
    response = &CreateCfsSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCfsSnapshot
// 创建文件系统快照
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsSnapshot(request *CreateCfsSnapshotRequest) (response *CreateCfsSnapshotResponse, err error) {
    return c.CreateCfsSnapshotWithContext(context.Background(), request)
}

// CreateCfsSnapshot
// 创建文件系统快照
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsSnapshotWithContext(ctx context.Context, request *CreateCfsSnapshotRequest) (response *CreateCfsSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateCfsSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoSnapshotPolicyRequest() (request *DeleteAutoSnapshotPolicyRequest) {
    request = &DeleteAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteAutoSnapshotPolicy")
    
    
    return
}

func NewDeleteAutoSnapshotPolicyResponse() (response *DeleteAutoSnapshotPolicyResponse) {
    response = &DeleteAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAutoSnapshotPolicy
// 删除快照定期策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (response *DeleteAutoSnapshotPolicyResponse, err error) {
    return c.DeleteAutoSnapshotPolicyWithContext(context.Background(), request)
}

// DeleteAutoSnapshotPolicy
// 删除快照定期策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteAutoSnapshotPolicyWithContext(ctx context.Context, request *DeleteAutoSnapshotPolicyRequest) (response *DeleteAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsFileSystemRequest() (request *DeleteCfsFileSystemRequest) {
    request = &DeleteCfsFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsFileSystem")
    
    
    return
}

func NewDeleteCfsFileSystemResponse() (response *DeleteCfsFileSystemResponse) {
    response = &DeleteCfsFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCfsFileSystem
// 用于删除文件系统
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  FAILEDOPERATION_UNTAGRESOURCEFAILED = "FailedOperation.UntagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsFileSystem(request *DeleteCfsFileSystemRequest) (response *DeleteCfsFileSystemResponse, err error) {
    return c.DeleteCfsFileSystemWithContext(context.Background(), request)
}

// DeleteCfsFileSystem
// 用于删除文件系统
//
// 可能返回的错误码:
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  FAILEDOPERATION_UNTAGRESOURCEFAILED = "FailedOperation.UntagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsFileSystemWithContext(ctx context.Context, request *DeleteCfsFileSystemRequest) (response *DeleteCfsFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteCfsFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsPGroupRequest() (request *DeleteCfsPGroupRequest) {
    request = &DeleteCfsPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsPGroup")
    
    
    return
}

func NewDeleteCfsPGroupResponse() (response *DeleteCfsPGroupResponse) {
    response = &DeleteCfsPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCfsPGroup
// 本接口（DeleteCfsPGroup）用于删除权限组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsPGroup(request *DeleteCfsPGroupRequest) (response *DeleteCfsPGroupResponse, err error) {
    return c.DeleteCfsPGroupWithContext(context.Background(), request)
}

// DeleteCfsPGroup
// 本接口（DeleteCfsPGroup）用于删除权限组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsPGroupWithContext(ctx context.Context, request *DeleteCfsPGroupRequest) (response *DeleteCfsPGroupResponse, err error) {
    if request == nil {
        request = NewDeleteCfsPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsRuleRequest() (request *DeleteCfsRuleRequest) {
    request = &DeleteCfsRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsRule")
    
    
    return
}

func NewDeleteCfsRuleResponse() (response *DeleteCfsRuleResponse) {
    response = &DeleteCfsRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCfsRule
// 本接口（DeleteCfsRule）用于删除权限组规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  RESOURCENOTFOUND_RULENOTFOUND = "ResourceNotFound.RuleNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsRule(request *DeleteCfsRuleRequest) (response *DeleteCfsRuleResponse, err error) {
    return c.DeleteCfsRuleWithContext(context.Background(), request)
}

// DeleteCfsRule
// 本接口（DeleteCfsRule）用于删除权限组规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  RESOURCENOTFOUND_RULENOTFOUND = "ResourceNotFound.RuleNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsRuleWithContext(ctx context.Context, request *DeleteCfsRuleRequest) (response *DeleteCfsRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCfsRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsSnapshotRequest() (request *DeleteCfsSnapshotRequest) {
    request = &DeleteCfsSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsSnapshot")
    
    
    return
}

func NewDeleteCfsSnapshotResponse() (response *DeleteCfsSnapshotResponse) {
    response = &DeleteCfsSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCfsSnapshot
// 删除文件系统快照
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsSnapshot(request *DeleteCfsSnapshotRequest) (response *DeleteCfsSnapshotResponse, err error) {
    return c.DeleteCfsSnapshotWithContext(context.Background(), request)
}

// DeleteCfsSnapshot
// 删除文件系统快照
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsSnapshotWithContext(ctx context.Context, request *DeleteCfsSnapshotRequest) (response *DeleteCfsSnapshotResponse, err error) {
    if request == nil {
        request = NewDeleteCfsSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMountTargetRequest() (request *DeleteMountTargetRequest) {
    request = &DeleteMountTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteMountTarget")
    
    
    return
}

func NewDeleteMountTargetResponse() (response *DeleteMountTargetResponse) {
    response = &DeleteMountTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMountTarget
// 本接口（DeleteMountTarget）用于删除挂载点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteMountTarget(request *DeleteMountTargetRequest) (response *DeleteMountTargetResponse, err error) {
    return c.DeleteMountTargetWithContext(context.Background(), request)
}

// DeleteMountTarget
// 本接口（DeleteMountTarget）用于删除挂载点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteMountTargetWithContext(ctx context.Context, request *DeleteMountTargetRequest) (response *DeleteMountTargetResponse, err error) {
    if request == nil {
        request = NewDeleteMountTargetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMountTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMountTargetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserQuotaRequest() (request *DeleteUserQuotaRequest) {
    request = &DeleteUserQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteUserQuota")
    
    
    return
}

func NewDeleteUserQuotaResponse() (response *DeleteUserQuotaResponse) {
    response = &DeleteUserQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUserQuota
// 指定条件删除文件系统配额
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUserQuota(request *DeleteUserQuotaRequest) (response *DeleteUserQuotaResponse, err error) {
    return c.DeleteUserQuotaWithContext(context.Background(), request)
}

// DeleteUserQuota
// 指定条件删除文件系统配额
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUserQuotaWithContext(ctx context.Context, request *DeleteUserQuotaRequest) (response *DeleteUserQuotaResponse, err error) {
    if request == nil {
        request = NewDeleteUserQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
    request = &DescribeAutoSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeAutoSnapshotPolicies")
    
    
    return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
    response = &DescribeAutoSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoSnapshotPolicies
// 查询文件系统快照定期策略列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
    return c.DescribeAutoSnapshotPoliciesWithContext(context.Background(), request)
}

// DescribeAutoSnapshotPolicies
// 查询文件系统快照定期策略列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAutoSnapshotPoliciesWithContext(ctx context.Context, request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoSnapshotPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoSnapshotPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableZoneInfoRequest() (request *DescribeAvailableZoneInfoRequest) {
    request = &DescribeAvailableZoneInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeAvailableZoneInfo")
    
    
    return
}

func NewDescribeAvailableZoneInfoResponse() (response *DescribeAvailableZoneInfoResponse) {
    response = &DescribeAvailableZoneInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableZoneInfo
// 本接口（DescribeAvailableZoneInfo）用于查询区域的可用情况。
//
// 可能返回的错误码:
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAvailableZoneInfo(request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
    return c.DescribeAvailableZoneInfoWithContext(context.Background(), request)
}

// DescribeAvailableZoneInfo
// 本接口（DescribeAvailableZoneInfo）用于查询区域的可用情况。
//
// 可能返回的错误码:
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAvailableZoneInfoWithContext(ctx context.Context, request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableZoneInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableZoneInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableZoneInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsFileSystemClientsRequest() (request *DescribeCfsFileSystemClientsRequest) {
    request = &DescribeCfsFileSystemClientsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystemClients")
    
    
    return
}

func NewDescribeCfsFileSystemClientsResponse() (response *DescribeCfsFileSystemClientsResponse) {
    response = &DescribeCfsFileSystemClientsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsFileSystemClients
// 查询挂载该文件系统的客户端。此功能需要客户端安装CFS监控插件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
func (c *Client) DescribeCfsFileSystemClients(request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
    return c.DescribeCfsFileSystemClientsWithContext(context.Background(), request)
}

// DescribeCfsFileSystemClients
// 查询挂载该文件系统的客户端。此功能需要客户端安装CFS监控插件。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
func (c *Client) DescribeCfsFileSystemClientsWithContext(ctx context.Context, request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemClientsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsFileSystemClients require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsFileSystemClientsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsFileSystemsRequest() (request *DescribeCfsFileSystemsRequest) {
    request = &DescribeCfsFileSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystems")
    
    
    return
}

func NewDescribeCfsFileSystemsResponse() (response *DescribeCfsFileSystemsResponse) {
    response = &DescribeCfsFileSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsFileSystems
// 本接口（DescribeCfsFileSystems）用于查询文件系统
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMIDORREGION = "InvalidParameterValue.MissingFileSystemIdOrRegion"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsFileSystems(request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
    return c.DescribeCfsFileSystemsWithContext(context.Background(), request)
}

// DescribeCfsFileSystems
// 本接口（DescribeCfsFileSystems）用于查询文件系统
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMIDORREGION = "InvalidParameterValue.MissingFileSystemIdOrRegion"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsFileSystemsWithContext(ctx context.Context, request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsFileSystems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsFileSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsPGroupsRequest() (request *DescribeCfsPGroupsRequest) {
    request = &DescribeCfsPGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsPGroups")
    
    
    return
}

func NewDescribeCfsPGroupsResponse() (response *DescribeCfsPGroupsResponse) {
    response = &DescribeCfsPGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsPGroups
// 本接口（DescribeCfsPGroups）用于查询权限组列表。
//
// 可能返回的错误码:
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsPGroups(request *DescribeCfsPGroupsRequest) (response *DescribeCfsPGroupsResponse, err error) {
    return c.DescribeCfsPGroupsWithContext(context.Background(), request)
}

// DescribeCfsPGroups
// 本接口（DescribeCfsPGroups）用于查询权限组列表。
//
// 可能返回的错误码:
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsPGroupsWithContext(ctx context.Context, request *DescribeCfsPGroupsRequest) (response *DescribeCfsPGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsPGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsPGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsPGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsRulesRequest() (request *DescribeCfsRulesRequest) {
    request = &DescribeCfsRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsRules")
    
    
    return
}

func NewDescribeCfsRulesResponse() (response *DescribeCfsRulesResponse) {
    response = &DescribeCfsRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsRules
// 本接口（DescribeCfsRules）用于查询权限组规则列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsRules(request *DescribeCfsRulesRequest) (response *DescribeCfsRulesResponse, err error) {
    return c.DescribeCfsRulesWithContext(context.Background(), request)
}

// DescribeCfsRules
// 本接口（DescribeCfsRules）用于查询权限组规则列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsRulesWithContext(ctx context.Context, request *DescribeCfsRulesRequest) (response *DescribeCfsRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCfsRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsServiceStatusRequest() (request *DescribeCfsServiceStatusRequest) {
    request = &DescribeCfsServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsServiceStatus")
    
    
    return
}

func NewDescribeCfsServiceStatusResponse() (response *DescribeCfsServiceStatusResponse) {
    response = &DescribeCfsServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsServiceStatus
// 本接口（DescribeCfsServiceStatus）用于查询用户使用CFS的服务状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsServiceStatus(request *DescribeCfsServiceStatusRequest) (response *DescribeCfsServiceStatusResponse, err error) {
    return c.DescribeCfsServiceStatusWithContext(context.Background(), request)
}

// DescribeCfsServiceStatus
// 本接口（DescribeCfsServiceStatus）用于查询用户使用CFS的服务状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsServiceStatusWithContext(ctx context.Context, request *DescribeCfsServiceStatusRequest) (response *DescribeCfsServiceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCfsServiceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsServiceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsSnapshotOverviewRequest() (request *DescribeCfsSnapshotOverviewRequest) {
    request = &DescribeCfsSnapshotOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsSnapshotOverview")
    
    
    return
}

func NewDescribeCfsSnapshotOverviewResponse() (response *DescribeCfsSnapshotOverviewResponse) {
    response = &DescribeCfsSnapshotOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsSnapshotOverview
// 文件系统快照概览
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshotOverview(request *DescribeCfsSnapshotOverviewRequest) (response *DescribeCfsSnapshotOverviewResponse, err error) {
    return c.DescribeCfsSnapshotOverviewWithContext(context.Background(), request)
}

// DescribeCfsSnapshotOverview
// 文件系统快照概览
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshotOverviewWithContext(ctx context.Context, request *DescribeCfsSnapshotOverviewRequest) (response *DescribeCfsSnapshotOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeCfsSnapshotOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsSnapshotOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsSnapshotOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsSnapshotsRequest() (request *DescribeCfsSnapshotsRequest) {
    request = &DescribeCfsSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsSnapshots")
    
    
    return
}

func NewDescribeCfsSnapshotsResponse() (response *DescribeCfsSnapshotsResponse) {
    response = &DescribeCfsSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsSnapshots
// 查询文件系统快照列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshots(request *DescribeCfsSnapshotsRequest) (response *DescribeCfsSnapshotsResponse, err error) {
    return c.DescribeCfsSnapshotsWithContext(context.Background(), request)
}

// DescribeCfsSnapshots
// 查询文件系统快照列表
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshotsWithContext(ctx context.Context, request *DescribeCfsSnapshotsRequest) (response *DescribeCfsSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
    request = &DescribeMountTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeMountTargets")
    
    
    return
}

func NewDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
    response = &DescribeMountTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMountTargets
// 本接口（DescribeMountTargets）用于查询文件系统挂载点信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
    return c.DescribeMountTargetsWithContext(context.Background(), request)
}

// DescribeMountTargets
// 本接口（DescribeMountTargets）用于查询文件系统挂载点信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeMountTargetsWithContext(ctx context.Context, request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeMountTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMountTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMountTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotOperationLogsRequest() (request *DescribeSnapshotOperationLogsRequest) {
    request = &DescribeSnapshotOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeSnapshotOperationLogs")
    
    
    return
}

func NewDescribeSnapshotOperationLogsResponse() (response *DescribeSnapshotOperationLogsResponse) {
    response = &DescribeSnapshotOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotOperationLogs
// 查询快照操作日志
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeSnapshotOperationLogs(request *DescribeSnapshotOperationLogsRequest) (response *DescribeSnapshotOperationLogsResponse, err error) {
    return c.DescribeSnapshotOperationLogsWithContext(context.Background(), request)
}

// DescribeSnapshotOperationLogs
// 查询快照操作日志
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeSnapshotOperationLogsWithContext(ctx context.Context, request *DescribeSnapshotOperationLogsRequest) (response *DescribeSnapshotOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotOperationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserQuotaRequest() (request *DescribeUserQuotaRequest) {
    request = &DescribeUserQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeUserQuota")
    
    
    return
}

func NewDescribeUserQuotaResponse() (response *DescribeUserQuotaResponse) {
    response = &DescribeUserQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserQuota
// 查询文件系统配额
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_QUOTAUSERTYPEERROR = "InvalidParameterValue.QuotaUserTypeError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserQuota(request *DescribeUserQuotaRequest) (response *DescribeUserQuotaResponse, err error) {
    return c.DescribeUserQuotaWithContext(context.Background(), request)
}

// DescribeUserQuota
// 查询文件系统配额
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_QUOTAUSERTYPEERROR = "InvalidParameterValue.QuotaUserTypeError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserQuotaWithContext(ctx context.Context, request *DescribeUserQuotaRequest) (response *DescribeUserQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewSetUserQuotaRequest() (request *SetUserQuotaRequest) {
    request = &SetUserQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "SetUserQuota")
    
    
    return
}

func NewSetUserQuotaResponse() (response *SetUserQuotaResponse) {
    response = &SetUserQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetUserQuota
// 设置文件系统配额，提供UID/GID的配额设置的接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_QUOTACAPLIMITERROR = "InvalidParameterValue.QuotaCapLimitError"
//  INVALIDPARAMETERVALUE_QUOTAFILELIMITERROR = "InvalidParameterValue.QuotaFileLimitError"
//  INVALIDPARAMETERVALUE_QUOTAUSERIDERROR = "InvalidParameterValue.QuotaUserIdError"
//  INVALIDPARAMETERVALUE_QUOTAUSERTYPEERROR = "InvalidParameterValue.QuotaUserTypeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetUserQuota(request *SetUserQuotaRequest) (response *SetUserQuotaResponse, err error) {
    return c.SetUserQuotaWithContext(context.Background(), request)
}

// SetUserQuota
// 设置文件系统配额，提供UID/GID的配额设置的接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_QUOTACAPLIMITERROR = "InvalidParameterValue.QuotaCapLimitError"
//  INVALIDPARAMETERVALUE_QUOTAFILELIMITERROR = "InvalidParameterValue.QuotaFileLimitError"
//  INVALIDPARAMETERVALUE_QUOTAUSERIDERROR = "InvalidParameterValue.QuotaUserIdError"
//  INVALIDPARAMETERVALUE_QUOTAUSERTYPEERROR = "InvalidParameterValue.QuotaUserTypeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetUserQuotaWithContext(ctx context.Context, request *SetUserQuotaRequest) (response *SetUserQuotaResponse, err error) {
    if request == nil {
        request = NewSetUserQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetUserQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetUserQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewSignUpCfsServiceRequest() (request *SignUpCfsServiceRequest) {
    request = &SignUpCfsServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "SignUpCfsService")
    
    
    return
}

func NewSignUpCfsServiceResponse() (response *SignUpCfsServiceResponse) {
    response = &SignUpCfsServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SignUpCfsService
// 本接口（SignUpCfsService）用于开通CFS服务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) SignUpCfsService(request *SignUpCfsServiceRequest) (response *SignUpCfsServiceResponse, err error) {
    return c.SignUpCfsServiceWithContext(context.Background(), request)
}

// SignUpCfsService
// 本接口（SignUpCfsService）用于开通CFS服务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) SignUpCfsServiceWithContext(ctx context.Context, request *SignUpCfsServiceRequest) (response *SignUpCfsServiceResponse, err error) {
    if request == nil {
        request = NewSignUpCfsServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SignUpCfsService require credential")
    }

    request.SetContext(ctx)
    
    response = NewSignUpCfsServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindAutoSnapshotPolicyRequest() (request *UnbindAutoSnapshotPolicyRequest) {
    request = &UnbindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UnbindAutoSnapshotPolicy")
    
    
    return
}

func NewUnbindAutoSnapshotPolicyResponse() (response *UnbindAutoSnapshotPolicyResponse) {
    response = &UnbindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindAutoSnapshotPolicy
// 解除文件系统绑定的快照策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UnbindAutoSnapshotPolicy(request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
    return c.UnbindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// UnbindAutoSnapshotPolicy
// 解除文件系统绑定的快照策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UnbindAutoSnapshotPolicyWithContext(ctx context.Context, request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewUnbindAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAutoSnapshotPolicyRequest() (request *UpdateAutoSnapshotPolicyRequest) {
    request = &UpdateAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateAutoSnapshotPolicy")
    
    
    return
}

func NewUpdateAutoSnapshotPolicyResponse() (response *UpdateAutoSnapshotPolicyResponse) {
    response = &UpdateAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAutoSnapshotPolicy
// 更新定期自动快照策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateAutoSnapshotPolicy(request *UpdateAutoSnapshotPolicyRequest) (response *UpdateAutoSnapshotPolicyResponse, err error) {
    return c.UpdateAutoSnapshotPolicyWithContext(context.Background(), request)
}

// UpdateAutoSnapshotPolicy
// 更新定期自动快照策略
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateAutoSnapshotPolicyWithContext(ctx context.Context, request *UpdateAutoSnapshotPolicyRequest) (response *UpdateAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewUpdateAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsFileSystemNameRequest() (request *UpdateCfsFileSystemNameRequest) {
    request = &UpdateCfsFileSystemNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemName")
    
    
    return
}

func NewUpdateCfsFileSystemNameResponse() (response *UpdateCfsFileSystemNameResponse) {
    response = &UpdateCfsFileSystemNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsFileSystemName
// 本接口（UpdateCfsFileSystemName）用于更新文件系统名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemName(request *UpdateCfsFileSystemNameRequest) (response *UpdateCfsFileSystemNameResponse, err error) {
    return c.UpdateCfsFileSystemNameWithContext(context.Background(), request)
}

// UpdateCfsFileSystemName
// 本接口（UpdateCfsFileSystemName）用于更新文件系统名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemNameWithContext(ctx context.Context, request *UpdateCfsFileSystemNameRequest) (response *UpdateCfsFileSystemNameResponse, err error) {
    if request == nil {
        request = NewUpdateCfsFileSystemNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsFileSystemName require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsFileSystemNameResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsFileSystemPGroupRequest() (request *UpdateCfsFileSystemPGroupRequest) {
    request = &UpdateCfsFileSystemPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemPGroup")
    
    
    return
}

func NewUpdateCfsFileSystemPGroupResponse() (response *UpdateCfsFileSystemPGroupResponse) {
    response = &UpdateCfsFileSystemPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsFileSystemPGroup
// 本接口（UpdateCfsFileSystemPGroup）用于更新文件系统所使用的权限组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemPGroup(request *UpdateCfsFileSystemPGroupRequest) (response *UpdateCfsFileSystemPGroupResponse, err error) {
    return c.UpdateCfsFileSystemPGroupWithContext(context.Background(), request)
}

// UpdateCfsFileSystemPGroup
// 本接口（UpdateCfsFileSystemPGroup）用于更新文件系统所使用的权限组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemPGroupWithContext(ctx context.Context, request *UpdateCfsFileSystemPGroupRequest) (response *UpdateCfsFileSystemPGroupResponse, err error) {
    if request == nil {
        request = NewUpdateCfsFileSystemPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsFileSystemPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsFileSystemPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsFileSystemSizeLimitRequest() (request *UpdateCfsFileSystemSizeLimitRequest) {
    request = &UpdateCfsFileSystemSizeLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemSizeLimit")
    
    
    return
}

func NewUpdateCfsFileSystemSizeLimitResponse() (response *UpdateCfsFileSystemSizeLimitResponse) {
    response = &UpdateCfsFileSystemSizeLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsFileSystemSizeLimit
// 本接口（UpdateCfsFileSystemSizeLimit）用于更新文件系统存储容量限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSIZELIMIT = "InvalidParameterValue.InvalidFsSizeLimit"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemSizeLimit(request *UpdateCfsFileSystemSizeLimitRequest) (response *UpdateCfsFileSystemSizeLimitResponse, err error) {
    return c.UpdateCfsFileSystemSizeLimitWithContext(context.Background(), request)
}

// UpdateCfsFileSystemSizeLimit
// 本接口（UpdateCfsFileSystemSizeLimit）用于更新文件系统存储容量限制。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSIZELIMIT = "InvalidParameterValue.InvalidFsSizeLimit"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemSizeLimitWithContext(ctx context.Context, request *UpdateCfsFileSystemSizeLimitRequest) (response *UpdateCfsFileSystemSizeLimitResponse, err error) {
    if request == nil {
        request = NewUpdateCfsFileSystemSizeLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsFileSystemSizeLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsFileSystemSizeLimitResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsPGroupRequest() (request *UpdateCfsPGroupRequest) {
    request = &UpdateCfsPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsPGroup")
    
    
    return
}

func NewUpdateCfsPGroupResponse() (response *UpdateCfsPGroupResponse) {
    response = &UpdateCfsPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsPGroup
// 本接口（UpdateCfsPGroup）更新权限组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGNAMEORDESCINFO = "InvalidParameterValue.MissingNameOrDescinfo"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsPGroup(request *UpdateCfsPGroupRequest) (response *UpdateCfsPGroupResponse, err error) {
    return c.UpdateCfsPGroupWithContext(context.Background(), request)
}

// UpdateCfsPGroup
// 本接口（UpdateCfsPGroup）更新权限组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGNAMEORDESCINFO = "InvalidParameterValue.MissingNameOrDescinfo"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsPGroupWithContext(ctx context.Context, request *UpdateCfsPGroupRequest) (response *UpdateCfsPGroupResponse, err error) {
    if request == nil {
        request = NewUpdateCfsPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsRuleRequest() (request *UpdateCfsRuleRequest) {
    request = &UpdateCfsRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsRule")
    
    
    return
}

func NewUpdateCfsRuleResponse() (response *UpdateCfsRuleResponse) {
    response = &UpdateCfsRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsRule
// 本接口（UpdateCfsRule）用于更新权限规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  FAILEDOPERATION_PGROUPLINKCFSV10 = "FailedOperation.PgroupLinkCfsv10"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  INVALIDPARAMETERVALUE_RULENOTMATCHPGROUP = "InvalidParameterValue.RuleNotMatchPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsRule(request *UpdateCfsRuleRequest) (response *UpdateCfsRuleResponse, err error) {
    return c.UpdateCfsRuleWithContext(context.Background(), request)
}

// UpdateCfsRule
// 本接口（UpdateCfsRule）用于更新权限规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  FAILEDOPERATION_PGROUPLINKCFSV10 = "FailedOperation.PgroupLinkCfsv10"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  INVALIDPARAMETERVALUE_RULENOTMATCHPGROUP = "InvalidParameterValue.RuleNotMatchPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsRuleWithContext(ctx context.Context, request *UpdateCfsRuleRequest) (response *UpdateCfsRuleResponse, err error) {
    if request == nil {
        request = NewUpdateCfsRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsSnapshotAttributeRequest() (request *UpdateCfsSnapshotAttributeRequest) {
    request = &UpdateCfsSnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsSnapshotAttribute")
    
    
    return
}

func NewUpdateCfsSnapshotAttributeResponse() (response *UpdateCfsSnapshotAttributeResponse) {
    response = &UpdateCfsSnapshotAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsSnapshotAttribute
// 更新文件系统快照名称及保留时长
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
func (c *Client) UpdateCfsSnapshotAttribute(request *UpdateCfsSnapshotAttributeRequest) (response *UpdateCfsSnapshotAttributeResponse, err error) {
    return c.UpdateCfsSnapshotAttributeWithContext(context.Background(), request)
}

// UpdateCfsSnapshotAttribute
// 更新文件系统快照名称及保留时长
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
func (c *Client) UpdateCfsSnapshotAttributeWithContext(ctx context.Context, request *UpdateCfsSnapshotAttributeRequest) (response *UpdateCfsSnapshotAttributeResponse, err error) {
    if request == nil {
        request = NewUpdateCfsSnapshotAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsSnapshotAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsSnapshotAttributeResponse()
    err = c.Send(request, response)
    return
}
