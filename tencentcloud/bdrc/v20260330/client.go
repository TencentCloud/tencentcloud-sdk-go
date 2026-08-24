// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20260330

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-03-30"

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


func NewApplyBackupGroupRequest() (request *ApplyBackupGroupRequest) {
    request = &ApplyBackupGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ApplyBackupGroup")
    
    
    return
}

func NewApplyBackupGroupResponse() (response *ApplyBackupGroupResponse) {
    response = &ApplyBackupGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyBackupGroup
// 回滚备份组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_SHOULDCONVERTBACKUPTOIMAGE = "InvalidParameter.ShouldConvertBackupToImage"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BACKUPGROUPDISKATTACHMULTIINSTANCE = "UnsupportedOperation.BackupGroupDiskAttachMultiInstance"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ApplyBackupGroup(request *ApplyBackupGroupRequest) (response *ApplyBackupGroupResponse, err error) {
    return c.ApplyBackupGroupWithContext(context.Background(), request)
}

// ApplyBackupGroup
// 回滚备份组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_SHOULDCONVERTBACKUPTOIMAGE = "InvalidParameter.ShouldConvertBackupToImage"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BACKUPGROUPDISKATTACHMULTIINSTANCE = "UnsupportedOperation.BackupGroupDiskAttachMultiInstance"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ApplyBackupGroupWithContext(ctx context.Context, request *ApplyBackupGroupRequest) (response *ApplyBackupGroupResponse, err error) {
    if request == nil {
        request = NewApplyBackupGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ApplyBackupGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyBackupGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyBackupGroupResponse()
    err = c.Send(request, response)
    return
}

func NewBindAutoBackupPolicyRequest() (request *BindAutoBackupPolicyRequest) {
    request = &BindAutoBackupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "BindAutoBackupPolicy")
    
    
    return
}

func NewBindAutoBackupPolicyResponse() (response *BindAutoBackupPolicyResponse) {
    response = &BindAutoBackupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindAutoBackupPolicy
// 将实例绑定到备份策略上
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_AUTOBACKUPPOLICYNOTFOUND = "ResourceNotFound.AutoBackupPolicyNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) BindAutoBackupPolicy(request *BindAutoBackupPolicyRequest) (response *BindAutoBackupPolicyResponse, err error) {
    return c.BindAutoBackupPolicyWithContext(context.Background(), request)
}

// BindAutoBackupPolicy
// 将实例绑定到备份策略上
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_AUTOBACKUPPOLICYNOTFOUND = "ResourceNotFound.AutoBackupPolicyNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) BindAutoBackupPolicyWithContext(ctx context.Context, request *BindAutoBackupPolicyRequest) (response *BindAutoBackupPolicyResponse, err error) {
    if request == nil {
        request = NewBindAutoBackupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "BindAutoBackupPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAutoBackupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAutoBackupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoBackupPolicyRequest() (request *CreateAutoBackupPolicyRequest) {
    request = &CreateAutoBackupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateAutoBackupPolicy")
    
    
    return
}

func NewCreateAutoBackupPolicyResponse() (response *CreateAutoBackupPolicyResponse) {
    response = &CreateAutoBackupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAutoBackupPolicy
// 创建备份策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAutoBackupPolicy(request *CreateAutoBackupPolicyRequest) (response *CreateAutoBackupPolicyResponse, err error) {
    return c.CreateAutoBackupPolicyWithContext(context.Background(), request)
}

// CreateAutoBackupPolicy
// 创建备份策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAutoBackupPolicyWithContext(ctx context.Context, request *CreateAutoBackupPolicyRequest) (response *CreateAutoBackupPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAutoBackupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateAutoBackupPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoBackupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoBackupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupGroupRequest() (request *CreateBackupGroupRequest) {
    request = &CreateBackupGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateBackupGroup")
    
    
    return
}

func NewCreateBackupGroupResponse() (response *CreateBackupGroupResponse) {
    response = &CreateBackupGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackupGroup
// 创建备份组
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_BACKUPCREATING = "ResourceUnavailable.BackupCreating"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) CreateBackupGroup(request *CreateBackupGroupRequest) (response *CreateBackupGroupResponse, err error) {
    return c.CreateBackupGroupWithContext(context.Background(), request)
}

// CreateBackupGroup
// 创建备份组
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_BACKUPCREATING = "ResourceUnavailable.BackupCreating"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) CreateBackupGroupWithContext(ctx context.Context, request *CreateBackupGroupRequest) (response *CreateBackupGroupResponse, err error) {
    if request == nil {
        request = NewCreateBackupGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateBackupGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupVaultRequest() (request *CreateBackupVaultRequest) {
    request = &CreateBackupVaultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateBackupVault")
    
    
    return
}

func NewCreateBackupVaultResponse() (response *CreateBackupVaultResponse) {
    response = &CreateBackupVaultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackupVault
// 创建备份库
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INTERNALERROR_COSCONNECTIVITYERROR = "InternalError.CosConnectivityError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBackupVault(request *CreateBackupVaultRequest) (response *CreateBackupVaultResponse, err error) {
    return c.CreateBackupVaultWithContext(context.Background(), request)
}

// CreateBackupVault
// 创建备份库
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INTERNALERROR_COSCONNECTIVITYERROR = "InternalError.CosConnectivityError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBackupVaultWithContext(ctx context.Context, request *CreateBackupVaultRequest) (response *CreateBackupVaultResponse, err error) {
    if request == nil {
        request = NewCreateBackupVaultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateBackupVault")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupVault require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupVaultResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisasterRecoveryProtectGroupRequest() (request *CreateDisasterRecoveryProtectGroupRequest) {
    request = &CreateDisasterRecoveryProtectGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateDisasterRecoveryProtectGroup")
    
    
    return
}

func NewCreateDisasterRecoveryProtectGroupResponse() (response *CreateDisasterRecoveryProtectGroupResponse) {
    response = &CreateDisasterRecoveryProtectGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDisasterRecoveryProtectGroup
// 本接口用于创建容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) CreateDisasterRecoveryProtectGroup(request *CreateDisasterRecoveryProtectGroupRequest) (response *CreateDisasterRecoveryProtectGroupResponse, err error) {
    return c.CreateDisasterRecoveryProtectGroupWithContext(context.Background(), request)
}

// CreateDisasterRecoveryProtectGroup
// 本接口用于创建容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) CreateDisasterRecoveryProtectGroupWithContext(ctx context.Context, request *CreateDisasterRecoveryProtectGroupRequest) (response *CreateDisasterRecoveryProtectGroupResponse, err error) {
    if request == nil {
        request = NewCreateDisasterRecoveryProtectGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateDisasterRecoveryProtectGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDisasterRecoveryProtectGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDisasterRecoveryProtectGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisasterRecoverySitePairRequest() (request *CreateDisasterRecoverySitePairRequest) {
    request = &CreateDisasterRecoverySitePairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateDisasterRecoverySitePair")
    
    
    return
}

func NewCreateDisasterRecoverySitePairResponse() (response *CreateDisasterRecoverySitePairResponse) {
    response = &CreateDisasterRecoverySitePairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDisasterRecoverySitePair
// 创建容灾站点对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) CreateDisasterRecoverySitePair(request *CreateDisasterRecoverySitePairRequest) (response *CreateDisasterRecoverySitePairResponse, err error) {
    return c.CreateDisasterRecoverySitePairWithContext(context.Background(), request)
}

// CreateDisasterRecoverySitePair
// 创建容灾站点对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) CreateDisasterRecoverySitePairWithContext(ctx context.Context, request *CreateDisasterRecoverySitePairRequest) (response *CreateDisasterRecoverySitePairResponse, err error) {
    if request == nil {
        request = NewCreateDisasterRecoverySitePairRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateDisasterRecoverySitePair")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDisasterRecoverySitePair require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDisasterRecoverySitePairResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisasterRecoveryVpcMappingRequest() (request *CreateDisasterRecoveryVpcMappingRequest) {
    request = &CreateDisasterRecoveryVpcMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateDisasterRecoveryVpcMapping")
    
    
    return
}

func NewCreateDisasterRecoveryVpcMappingResponse() (response *CreateDisasterRecoveryVpcMappingResponse) {
    response = &CreateDisasterRecoveryVpcMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDisasterRecoveryVpcMapping
// 本接口用于创建容灾站点VPC网络映射
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) CreateDisasterRecoveryVpcMapping(request *CreateDisasterRecoveryVpcMappingRequest) (response *CreateDisasterRecoveryVpcMappingResponse, err error) {
    return c.CreateDisasterRecoveryVpcMappingWithContext(context.Background(), request)
}

// CreateDisasterRecoveryVpcMapping
// 本接口用于创建容灾站点VPC网络映射
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) CreateDisasterRecoveryVpcMappingWithContext(ctx context.Context, request *CreateDisasterRecoveryVpcMappingRequest) (response *CreateDisasterRecoveryVpcMappingResponse, err error) {
    if request == nil {
        request = NewCreateDisasterRecoveryVpcMappingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateDisasterRecoveryVpcMapping")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDisasterRecoveryVpcMapping require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDisasterRecoveryVpcMappingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileBackupRequest() (request *CreateFileBackupRequest) {
    request = &CreateFileBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateFileBackup")
    
    
    return
}

func NewCreateFileBackupResponse() (response *CreateFileBackupResponse) {
    response = &CreateFileBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileBackup
// 本接口用于创建文件备份点
//
// 可能返回的错误码:
//  INTERNALERROR_CVMQUERYFAILED = "InternalError.CvmQueryFailed"
//  INTERNALERROR_INITREPOSITORYFAILED = "InternalError.InitRepositoryFailed"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER_PATHNESTED = "InvalidParameter.PathNested"
//  INVALIDPARAMETERVALUE_VAULTSTATEERROR = "InvalidParameterValue.VaultStateError"
//  RESOURCENOTFOUND_AGENTNOTINSTALLED = "ResourceNotFound.AgentNotInstalled"
//  RESOURCENOTFOUND_VAULTNOTEXIST = "ResourceNotFound.VaultNotExist"
//  UNAUTHORIZEDOPERATION_LACKOFQUALIFICATION = "UnauthorizedOperation.LackOfQualification"
//  UNSUPPORTEDOPERATION_AGENTNOTACTIVE = "UnsupportedOperation.AgentNotActive"
//  UNSUPPORTEDOPERATION_FILEBACKUPRESTORECONFLICT = "UnsupportedOperation.FileBackupRestoreConflict"
//  UNSUPPORTEDOPERATION_FILEBACKUPTASKDUPLICATE = "UnsupportedOperation.FileBackupTaskDuplicate"
func (c *Client) CreateFileBackup(request *CreateFileBackupRequest) (response *CreateFileBackupResponse, err error) {
    return c.CreateFileBackupWithContext(context.Background(), request)
}

// CreateFileBackup
// 本接口用于创建文件备份点
//
// 可能返回的错误码:
//  INTERNALERROR_CVMQUERYFAILED = "InternalError.CvmQueryFailed"
//  INTERNALERROR_INITREPOSITORYFAILED = "InternalError.InitRepositoryFailed"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER_PATHNESTED = "InvalidParameter.PathNested"
//  INVALIDPARAMETERVALUE_VAULTSTATEERROR = "InvalidParameterValue.VaultStateError"
//  RESOURCENOTFOUND_AGENTNOTINSTALLED = "ResourceNotFound.AgentNotInstalled"
//  RESOURCENOTFOUND_VAULTNOTEXIST = "ResourceNotFound.VaultNotExist"
//  UNAUTHORIZEDOPERATION_LACKOFQUALIFICATION = "UnauthorizedOperation.LackOfQualification"
//  UNSUPPORTEDOPERATION_AGENTNOTACTIVE = "UnsupportedOperation.AgentNotActive"
//  UNSUPPORTEDOPERATION_FILEBACKUPRESTORECONFLICT = "UnsupportedOperation.FileBackupRestoreConflict"
//  UNSUPPORTEDOPERATION_FILEBACKUPTASKDUPLICATE = "UnsupportedOperation.FileBackupTaskDuplicate"
func (c *Client) CreateFileBackupWithContext(ctx context.Context, request *CreateFileBackupRequest) (response *CreateFileBackupResponse, err error) {
    if request == nil {
        request = NewCreateFileBackupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateFileBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileBackupPlanRequest() (request *CreateFileBackupPlanRequest) {
    request = &CreateFileBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateFileBackupPlan")
    
    
    return
}

func NewCreateFileBackupPlanResponse() (response *CreateFileBackupPlanResponse) {
    response = &CreateFileBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileBackupPlan
// 本接口用于创建备份计划
//
// 可能返回的错误码:
//  INTERNALERROR_CVMQUERYFAILED = "InternalError.CvmQueryFailed"
//  INVALIDPARAMETER_PATHNESTED = "InvalidParameter.PathNested"
//  INVALIDPARAMETER_POLICYNOTSUPPORTVAULT = "InvalidParameter.PolicyNotSupportVault"
//  INVALIDPARAMETER_POLICYTYPEMISMATCH = "InvalidParameter.PolicyTypeMismatch"
//  INVALIDPARAMETER_POLICYVAULTCONFLICT = "InvalidParameter.PolicyVaultConflict"
//  INVALIDPARAMETERVALUE_CVMINSTANCENOTEXIST = "InvalidParameterValue.CvmInstanceNotExist"
//  INVALIDPARAMETERVALUE_POLICYNOTAVAILABLE = "InvalidParameterValue.PolicyNotAvailable"
//  INVALIDPARAMETERVALUE_VAULTSTATEERROR = "InvalidParameterValue.VaultStateError"
//  RESOURCENOTFOUND_AGENTNOTINSTALLED = "ResourceNotFound.AgentNotInstalled"
//  RESOURCENOTFOUND_VAULTNOTEXIST = "ResourceNotFound.VaultNotExist"
//  UNSUPPORTEDOPERATION_AGENTNOTACTIVE = "UnsupportedOperation.AgentNotActive"
func (c *Client) CreateFileBackupPlan(request *CreateFileBackupPlanRequest) (response *CreateFileBackupPlanResponse, err error) {
    return c.CreateFileBackupPlanWithContext(context.Background(), request)
}

// CreateFileBackupPlan
// 本接口用于创建备份计划
//
// 可能返回的错误码:
//  INTERNALERROR_CVMQUERYFAILED = "InternalError.CvmQueryFailed"
//  INVALIDPARAMETER_PATHNESTED = "InvalidParameter.PathNested"
//  INVALIDPARAMETER_POLICYNOTSUPPORTVAULT = "InvalidParameter.PolicyNotSupportVault"
//  INVALIDPARAMETER_POLICYTYPEMISMATCH = "InvalidParameter.PolicyTypeMismatch"
//  INVALIDPARAMETER_POLICYVAULTCONFLICT = "InvalidParameter.PolicyVaultConflict"
//  INVALIDPARAMETERVALUE_CVMINSTANCENOTEXIST = "InvalidParameterValue.CvmInstanceNotExist"
//  INVALIDPARAMETERVALUE_POLICYNOTAVAILABLE = "InvalidParameterValue.PolicyNotAvailable"
//  INVALIDPARAMETERVALUE_VAULTSTATEERROR = "InvalidParameterValue.VaultStateError"
//  RESOURCENOTFOUND_AGENTNOTINSTALLED = "ResourceNotFound.AgentNotInstalled"
//  RESOURCENOTFOUND_VAULTNOTEXIST = "ResourceNotFound.VaultNotExist"
//  UNSUPPORTEDOPERATION_AGENTNOTACTIVE = "UnsupportedOperation.AgentNotActive"
func (c *Client) CreateFileBackupPlanWithContext(ctx context.Context, request *CreateFileBackupPlanRequest) (response *CreateFileBackupPlanResponse, err error) {
    if request == nil {
        request = NewCreateFileBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateFileBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileRestoreTaskRequest() (request *CreateFileRestoreTaskRequest) {
    request = &CreateFileRestoreTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateFileRestoreTask")
    
    
    return
}

func NewCreateFileRestoreTaskResponse() (response *CreateFileRestoreTaskResponse) {
    response = &CreateFileRestoreTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileRestoreTask
// 创建恢复任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDCONFLICTSTRATEGY = "InvalidParameterValue.InvalidConflictStrategy"
//  INVALIDPARAMETERVALUE_PATHTOOLONG = "InvalidParameterValue.PathTooLong"
//  RESOURCENOTFOUND_AGENTNOTINSTALLED = "ResourceNotFound.AgentNotInstalled"
//  RESOURCENOTFOUND_FILEBACKUPNOTEXIST = "ResourceNotFound.FileBackupNotExist"
//  UNSUPPORTEDOPERATION_AGENTNOTACTIVE = "UnsupportedOperation.AgentNotActive"
//  UNSUPPORTEDOPERATION_FILEBACKUPRESTORECONFLICT = "UnsupportedOperation.FileBackupRestoreConflict"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
//  UNSUPPORTEDOPERATION_INSTANCESNAPSHOTROLLBACKING = "UnsupportedOperation.InstanceSnapshotRollbacking"
func (c *Client) CreateFileRestoreTask(request *CreateFileRestoreTaskRequest) (response *CreateFileRestoreTaskResponse, err error) {
    return c.CreateFileRestoreTaskWithContext(context.Background(), request)
}

// CreateFileRestoreTask
// 创建恢复任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDCONFLICTSTRATEGY = "InvalidParameterValue.InvalidConflictStrategy"
//  INVALIDPARAMETERVALUE_PATHTOOLONG = "InvalidParameterValue.PathTooLong"
//  RESOURCENOTFOUND_AGENTNOTINSTALLED = "ResourceNotFound.AgentNotInstalled"
//  RESOURCENOTFOUND_FILEBACKUPNOTEXIST = "ResourceNotFound.FileBackupNotExist"
//  UNSUPPORTEDOPERATION_AGENTNOTACTIVE = "UnsupportedOperation.AgentNotActive"
//  UNSUPPORTEDOPERATION_FILEBACKUPRESTORECONFLICT = "UnsupportedOperation.FileBackupRestoreConflict"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
//  UNSUPPORTEDOPERATION_INSTANCESNAPSHOTROLLBACKING = "UnsupportedOperation.InstanceSnapshotRollbacking"
func (c *Client) CreateFileRestoreTaskWithContext(ctx context.Context, request *CreateFileRestoreTaskRequest) (response *CreateFileRestoreTaskResponse, err error) {
    if request == nil {
        request = NewCreateFileRestoreTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateFileRestoreTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileRestoreTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileRestoreTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceCopyPairRequest() (request *CreateInstanceCopyPairRequest) {
    request = &CreateInstanceCopyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateInstanceCopyPair")
    
    
    return
}

func NewCreateInstanceCopyPairResponse() (response *CreateInstanceCopyPairResponse) {
    response = &CreateInstanceCopyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceCopyPair
// 本接口用于创建CVM复制对
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_SOURCECVMNOTFOUND = "ResourceNotFound.SourceCVMNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
func (c *Client) CreateInstanceCopyPair(request *CreateInstanceCopyPairRequest) (response *CreateInstanceCopyPairResponse, err error) {
    return c.CreateInstanceCopyPairWithContext(context.Background(), request)
}

// CreateInstanceCopyPair
// 本接口用于创建CVM复制对
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_SOURCECVMNOTFOUND = "ResourceNotFound.SourceCVMNotFound"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
func (c *Client) CreateInstanceCopyPairWithContext(ctx context.Context, request *CreateInstanceCopyPairRequest) (response *CreateInstanceCopyPairResponse, err error) {
    if request == nil {
        request = NewCreateInstanceCopyPairRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateInstanceCopyPair")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceCopyPair require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceCopyPairResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceDrillPairsRequest() (request *CreateInstanceDrillPairsRequest) {
    request = &CreateInstanceDrillPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateInstanceDrillPairs")
    
    
    return
}

func NewCreateInstanceDrillPairsResponse() (response *CreateInstanceDrillPairsResponse) {
    response = &CreateInstanceDrillPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceDrillPairs
// 创建cvm演练
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPBINDRESOURCE = "UnsupportedOperation.DisasterRecoveryProtectGroupBindResource"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPSTATEERROR = "UnsupportedOperation.DisasterRecoveryProtectGroupStateError"
//  UNSUPPORTEDOPERATION_HAVENOPROTECTIONPOINT = "UnsupportedOperation.HaveNoProtectionPoint"
func (c *Client) CreateInstanceDrillPairs(request *CreateInstanceDrillPairsRequest) (response *CreateInstanceDrillPairsResponse, err error) {
    return c.CreateInstanceDrillPairsWithContext(context.Background(), request)
}

// CreateInstanceDrillPairs
// 创建cvm演练
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPBINDRESOURCE = "UnsupportedOperation.DisasterRecoveryProtectGroupBindResource"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPSTATEERROR = "UnsupportedOperation.DisasterRecoveryProtectGroupStateError"
//  UNSUPPORTEDOPERATION_HAVENOPROTECTIONPOINT = "UnsupportedOperation.HaveNoProtectionPoint"
func (c *Client) CreateInstanceDrillPairsWithContext(ctx context.Context, request *CreateInstanceDrillPairsRequest) (response *CreateInstanceDrillPairsResponse, err error) {
    if request == nil {
        request = NewCreateInstanceDrillPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateInstanceDrillPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceDrillPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceDrillPairsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupMappingRequest() (request *CreateSecurityGroupMappingRequest) {
    request = &CreateSecurityGroupMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "CreateSecurityGroupMapping")
    
    
    return
}

func NewCreateSecurityGroupMappingResponse() (response *CreateSecurityGroupMappingResponse) {
    response = &CreateSecurityGroupMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityGroupMapping
// 本接口用于为站点对新增安全组映射，生产端实例绑定的安全组为源端，需要为每个生产端实例绑定的安全组建立映射，在创建复制对时，会自动以映射后的目标安全组作为容灾端实例绑定的安全组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) CreateSecurityGroupMapping(request *CreateSecurityGroupMappingRequest) (response *CreateSecurityGroupMappingResponse, err error) {
    return c.CreateSecurityGroupMappingWithContext(context.Background(), request)
}

// CreateSecurityGroupMapping
// 本接口用于为站点对新增安全组映射，生产端实例绑定的安全组为源端，需要为每个生产端实例绑定的安全组建立映射，在创建复制对时，会自动以映射后的目标安全组作为容灾端实例绑定的安全组。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) CreateSecurityGroupMappingWithContext(ctx context.Context, request *CreateSecurityGroupMappingRequest) (response *CreateSecurityGroupMappingResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupMappingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "CreateSecurityGroupMapping")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityGroupMapping require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityGroupMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoBackupPoliciesRequest() (request *DeleteAutoBackupPoliciesRequest) {
    request = &DeleteAutoBackupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteAutoBackupPolicies")
    
    
    return
}

func NewDeleteAutoBackupPoliciesResponse() (response *DeleteAutoBackupPoliciesResponse) {
    response = &DeleteAutoBackupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAutoBackupPolicies
// 删除备份策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  RESOURCENOTFOUND_AUTOBACKUPPOLICYNOTFOUND = "ResourceNotFound.AutoBackupPolicyNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) DeleteAutoBackupPolicies(request *DeleteAutoBackupPoliciesRequest) (response *DeleteAutoBackupPoliciesResponse, err error) {
    return c.DeleteAutoBackupPoliciesWithContext(context.Background(), request)
}

// DeleteAutoBackupPolicies
// 删除备份策略
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  RESOURCENOTFOUND_AUTOBACKUPPOLICYNOTFOUND = "ResourceNotFound.AutoBackupPolicyNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) DeleteAutoBackupPoliciesWithContext(ctx context.Context, request *DeleteAutoBackupPoliciesRequest) (response *DeleteAutoBackupPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteAutoBackupPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteAutoBackupPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoBackupPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoBackupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupGroupsRequest() (request *DeleteBackupGroupsRequest) {
    request = &DeleteBackupGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteBackupGroups")
    
    
    return
}

func NewDeleteBackupGroupsResponse() (response *DeleteBackupGroupsResponse) {
    response = &DeleteBackupGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBackupGroups
// 删除备份组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) DeleteBackupGroups(request *DeleteBackupGroupsRequest) (response *DeleteBackupGroupsResponse, err error) {
    return c.DeleteBackupGroupsWithContext(context.Background(), request)
}

// DeleteBackupGroups
// 删除备份组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) DeleteBackupGroupsWithContext(ctx context.Context, request *DeleteBackupGroupsRequest) (response *DeleteBackupGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteBackupGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteBackupGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackupGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupVaultsRequest() (request *DeleteBackupVaultsRequest) {
    request = &DeleteBackupVaultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteBackupVaults")
    
    
    return
}

func NewDeleteBackupVaultsResponse() (response *DeleteBackupVaultsResponse) {
    response = &DeleteBackupVaultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBackupVaults
// 删除备份库
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DeleteBackupVaults(request *DeleteBackupVaultsRequest) (response *DeleteBackupVaultsResponse, err error) {
    return c.DeleteBackupVaultsWithContext(context.Background(), request)
}

// DeleteBackupVaults
// 删除备份库
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DeleteBackupVaultsWithContext(ctx context.Context, request *DeleteBackupVaultsRequest) (response *DeleteBackupVaultsResponse, err error) {
    if request == nil {
        request = NewDeleteBackupVaultsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteBackupVaults")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackupVaults require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupVaultsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCopyPairsRequest() (request *DeleteCopyPairsRequest) {
    request = &DeleteCopyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteCopyPairs")
    
    
    return
}

func NewDeleteCopyPairsResponse() (response *DeleteCopyPairsResponse) {
    response = &DeleteCopyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCopyPairs
// 本接口用于删除容灾复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
//  UNSUPPORTEDOPERATION_TARGETRESOURCEROLLBACKING = "UnsupportedOperation.TargetResourceRollbacking"
func (c *Client) DeleteCopyPairs(request *DeleteCopyPairsRequest) (response *DeleteCopyPairsResponse, err error) {
    return c.DeleteCopyPairsWithContext(context.Background(), request)
}

// DeleteCopyPairs
// 本接口用于删除容灾复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
//  UNSUPPORTEDOPERATION_TARGETRESOURCEROLLBACKING = "UnsupportedOperation.TargetResourceRollbacking"
func (c *Client) DeleteCopyPairsWithContext(ctx context.Context, request *DeleteCopyPairsRequest) (response *DeleteCopyPairsResponse, err error) {
    if request == nil {
        request = NewDeleteCopyPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteCopyPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCopyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCopyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDisasterRecoveryProtectGroupsRequest() (request *DeleteDisasterRecoveryProtectGroupsRequest) {
    request = &DeleteDisasterRecoveryProtectGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteDisasterRecoveryProtectGroups")
    
    
    return
}

func NewDeleteDisasterRecoveryProtectGroupsResponse() (response *DeleteDisasterRecoveryProtectGroupsResponse) {
    response = &DeleteDisasterRecoveryProtectGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDisasterRecoveryProtectGroups
// 本接口用于删除容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPBINDRESOURCE = "UnsupportedOperation.DisasterRecoveryProtectGroupBindResource"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPSTATEERROR = "UnsupportedOperation.DisasterRecoveryProtectGroupStateError"
func (c *Client) DeleteDisasterRecoveryProtectGroups(request *DeleteDisasterRecoveryProtectGroupsRequest) (response *DeleteDisasterRecoveryProtectGroupsResponse, err error) {
    return c.DeleteDisasterRecoveryProtectGroupsWithContext(context.Background(), request)
}

// DeleteDisasterRecoveryProtectGroups
// 本接口用于删除容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPBINDRESOURCE = "UnsupportedOperation.DisasterRecoveryProtectGroupBindResource"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPSTATEERROR = "UnsupportedOperation.DisasterRecoveryProtectGroupStateError"
func (c *Client) DeleteDisasterRecoveryProtectGroupsWithContext(ctx context.Context, request *DeleteDisasterRecoveryProtectGroupsRequest) (response *DeleteDisasterRecoveryProtectGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDisasterRecoveryProtectGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteDisasterRecoveryProtectGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDisasterRecoveryProtectGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDisasterRecoveryProtectGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDisasterRecoverySitePairsRequest() (request *DeleteDisasterRecoverySitePairsRequest) {
    request = &DeleteDisasterRecoverySitePairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteDisasterRecoverySitePairs")
    
    
    return
}

func NewDeleteDisasterRecoverySitePairsResponse() (response *DeleteDisasterRecoverySitePairsResponse) {
    response = &DeleteDisasterRecoverySitePairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDisasterRecoverySitePairs
// 删除容灾站点对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRBINDRESOURCE = "UnsupportedOperation.DisasterRecoverySitePairBindResource"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) DeleteDisasterRecoverySitePairs(request *DeleteDisasterRecoverySitePairsRequest) (response *DeleteDisasterRecoverySitePairsResponse, err error) {
    return c.DeleteDisasterRecoverySitePairsWithContext(context.Background(), request)
}

// DeleteDisasterRecoverySitePairs
// 删除容灾站点对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRBINDRESOURCE = "UnsupportedOperation.DisasterRecoverySitePairBindResource"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) DeleteDisasterRecoverySitePairsWithContext(ctx context.Context, request *DeleteDisasterRecoverySitePairsRequest) (response *DeleteDisasterRecoverySitePairsResponse, err error) {
    if request == nil {
        request = NewDeleteDisasterRecoverySitePairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteDisasterRecoverySitePairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDisasterRecoverySitePairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDisasterRecoverySitePairsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDisasterRecoveryVpcMappingRequest() (request *DeleteDisasterRecoveryVpcMappingRequest) {
    request = &DeleteDisasterRecoveryVpcMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteDisasterRecoveryVpcMapping")
    
    
    return
}

func NewDeleteDisasterRecoveryVpcMappingResponse() (response *DeleteDisasterRecoveryVpcMappingResponse) {
    response = &DeleteDisasterRecoveryVpcMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDisasterRecoveryVpcMapping
// 本接口用于删除容灾站点对vpc映射信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCENOTFOUND_VPCMAPPINGNOTEXIST = "ResourceNotFound.VpcMappingNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DeleteDisasterRecoveryVpcMapping(request *DeleteDisasterRecoveryVpcMappingRequest) (response *DeleteDisasterRecoveryVpcMappingResponse, err error) {
    return c.DeleteDisasterRecoveryVpcMappingWithContext(context.Background(), request)
}

// DeleteDisasterRecoveryVpcMapping
// 本接口用于删除容灾站点对vpc映射信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCENOTFOUND_VPCMAPPINGNOTEXIST = "ResourceNotFound.VpcMappingNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DeleteDisasterRecoveryVpcMappingWithContext(ctx context.Context, request *DeleteDisasterRecoveryVpcMappingRequest) (response *DeleteDisasterRecoveryVpcMappingResponse, err error) {
    if request == nil {
        request = NewDeleteDisasterRecoveryVpcMappingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteDisasterRecoveryVpcMapping")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDisasterRecoveryVpcMapping require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDisasterRecoveryVpcMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDrillPairsRequest() (request *DeleteDrillPairsRequest) {
    request = &DeleteDrillPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteDrillPairs")
    
    
    return
}

func NewDeleteDrillPairsResponse() (response *DeleteDrillPairsResponse) {
    response = &DeleteDrillPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDrillPairs
// 删除演练对/演练组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DRILLNOTEXIST = "ResourceNotFound.DrillNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DRILLPAIRSTATEERROR = "UnsupportedOperation.DrillPairStateError"
//  UNSUPPORTEDOPERATION_TARGETRESOURCEROLLBACKING = "UnsupportedOperation.TargetResourceRollbacking"
func (c *Client) DeleteDrillPairs(request *DeleteDrillPairsRequest) (response *DeleteDrillPairsResponse, err error) {
    return c.DeleteDrillPairsWithContext(context.Background(), request)
}

// DeleteDrillPairs
// 删除演练对/演练组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DRILLNOTEXIST = "ResourceNotFound.DrillNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DRILLPAIRSTATEERROR = "UnsupportedOperation.DrillPairStateError"
//  UNSUPPORTEDOPERATION_TARGETRESOURCEROLLBACKING = "UnsupportedOperation.TargetResourceRollbacking"
func (c *Client) DeleteDrillPairsWithContext(ctx context.Context, request *DeleteDrillPairsRequest) (response *DeleteDrillPairsResponse, err error) {
    if request == nil {
        request = NewDeleteDrillPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteDrillPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDrillPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDrillPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileBackupPlansRequest() (request *DeleteFileBackupPlansRequest) {
    request = &DeleteFileBackupPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteFileBackupPlans")
    
    
    return
}

func NewDeleteFileBackupPlansResponse() (response *DeleteFileBackupPlansResponse) {
    response = &DeleteFileBackupPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFileBackupPlans
// 删除备份计划
//
// 可能返回的错误码:
//  INTERNALERROR_PARTIALDELETEFAILED = "InternalError.PartialDeleteFailed"
func (c *Client) DeleteFileBackupPlans(request *DeleteFileBackupPlansRequest) (response *DeleteFileBackupPlansResponse, err error) {
    return c.DeleteFileBackupPlansWithContext(context.Background(), request)
}

// DeleteFileBackupPlans
// 删除备份计划
//
// 可能返回的错误码:
//  INTERNALERROR_PARTIALDELETEFAILED = "InternalError.PartialDeleteFailed"
func (c *Client) DeleteFileBackupPlansWithContext(ctx context.Context, request *DeleteFileBackupPlansRequest) (response *DeleteFileBackupPlansResponse, err error) {
    if request == nil {
        request = NewDeleteFileBackupPlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteFileBackupPlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFileBackupPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileBackupPlansResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileBackupsRequest() (request *DeleteFileBackupsRequest) {
    request = &DeleteFileBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteFileBackups")
    
    
    return
}

func NewDeleteFileBackupsResponse() (response *DeleteFileBackupsResponse) {
    response = &DeleteFileBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFileBackups
// 删除文件备份点
//
// 可能返回的错误码:
//  INTERNALERROR_PARTIALDELETEFAILED = "InternalError.PartialDeleteFailed"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
func (c *Client) DeleteFileBackups(request *DeleteFileBackupsRequest) (response *DeleteFileBackupsResponse, err error) {
    return c.DeleteFileBackupsWithContext(context.Background(), request)
}

// DeleteFileBackups
// 删除文件备份点
//
// 可能返回的错误码:
//  INTERNALERROR_PARTIALDELETEFAILED = "InternalError.PartialDeleteFailed"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
func (c *Client) DeleteFileBackupsWithContext(ctx context.Context, request *DeleteFileBackupsRequest) (response *DeleteFileBackupsResponse, err error) {
    if request == nil {
        request = NewDeleteFileBackupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteFileBackups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFileBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupMappingRequest() (request *DeleteSecurityGroupMappingRequest) {
    request = &DeleteSecurityGroupMappingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DeleteSecurityGroupMapping")
    
    
    return
}

func NewDeleteSecurityGroupMappingResponse() (response *DeleteSecurityGroupMappingResponse) {
    response = &DeleteSecurityGroupMappingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityGroupMapping
// 本接口用于删除站点对已添加的安全组映射
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
func (c *Client) DeleteSecurityGroupMapping(request *DeleteSecurityGroupMappingRequest) (response *DeleteSecurityGroupMappingResponse, err error) {
    return c.DeleteSecurityGroupMappingWithContext(context.Background(), request)
}

// DeleteSecurityGroupMapping
// 本接口用于删除站点对已添加的安全组映射
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
func (c *Client) DeleteSecurityGroupMappingWithContext(ctx context.Context, request *DeleteSecurityGroupMappingRequest) (response *DeleteSecurityGroupMappingResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupMappingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DeleteSecurityGroupMapping")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityGroupMapping require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityGroupMappingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoBackupPoliciesRequest() (request *DescribeAutoBackupPoliciesRequest) {
    request = &DescribeAutoBackupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeAutoBackupPolicies")
    
    
    return
}

func NewDescribeAutoBackupPoliciesResponse() (response *DescribeAutoBackupPoliciesResponse) {
    response = &DescribeAutoBackupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoBackupPolicies
// 查询定期备份策略列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAutoBackupPolicies(request *DescribeAutoBackupPoliciesRequest) (response *DescribeAutoBackupPoliciesResponse, err error) {
    return c.DescribeAutoBackupPoliciesWithContext(context.Background(), request)
}

// DescribeAutoBackupPolicies
// 查询定期备份策略列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAutoBackupPoliciesWithContext(ctx context.Context, request *DescribeAutoBackupPoliciesRequest) (response *DescribeAutoBackupPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoBackupPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeAutoBackupPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoBackupPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoBackupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupGroupRollbackTasksRequest() (request *DescribeBackupGroupRollbackTasksRequest) {
    request = &DescribeBackupGroupRollbackTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupGroupRollbackTasks")
    
    
    return
}

func NewDescribeBackupGroupRollbackTasksResponse() (response *DescribeBackupGroupRollbackTasksResponse) {
    response = &DescribeBackupGroupRollbackTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupGroupRollbackTasks
// 查询备份组恢复任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupGroupRollbackTasks(request *DescribeBackupGroupRollbackTasksRequest) (response *DescribeBackupGroupRollbackTasksResponse, err error) {
    return c.DescribeBackupGroupRollbackTasksWithContext(context.Background(), request)
}

// DescribeBackupGroupRollbackTasks
// 查询备份组恢复任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupGroupRollbackTasksWithContext(ctx context.Context, request *DescribeBackupGroupRollbackTasksRequest) (response *DescribeBackupGroupRollbackTasksResponse, err error) {
    if request == nil {
        request = NewDescribeBackupGroupRollbackTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupGroupRollbackTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupGroupRollbackTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupGroupRollbackTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupGroupsRequest() (request *DescribeBackupGroupsRequest) {
    request = &DescribeBackupGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupGroups")
    
    
    return
}

func NewDescribeBackupGroupsResponse() (response *DescribeBackupGroupsResponse) {
    response = &DescribeBackupGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupGroups
// 查询备份组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupGroups(request *DescribeBackupGroupsRequest) (response *DescribeBackupGroupsResponse, err error) {
    return c.DescribeBackupGroupsWithContext(context.Background(), request)
}

// DescribeBackupGroups
// 查询备份组列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupGroupsWithContext(ctx context.Context, request *DescribeBackupGroupsRequest) (response *DescribeBackupGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupGroupsDeniedActionsRequest() (request *DescribeBackupGroupsDeniedActionsRequest) {
    request = &DescribeBackupGroupsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupGroupsDeniedActions")
    
    
    return
}

func NewDescribeBackupGroupsDeniedActionsResponse() (response *DescribeBackupGroupsDeniedActionsResponse) {
    response = &DescribeBackupGroupsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupGroupsDeniedActions
// 查询操作掩码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupGroupsDeniedActions(request *DescribeBackupGroupsDeniedActionsRequest) (response *DescribeBackupGroupsDeniedActionsResponse, err error) {
    return c.DescribeBackupGroupsDeniedActionsWithContext(context.Background(), request)
}

// DescribeBackupGroupsDeniedActions
// 查询操作掩码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupGroupsDeniedActionsWithContext(ctx context.Context, request *DescribeBackupGroupsDeniedActionsRequest) (response *DescribeBackupGroupsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupGroupsDeniedActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupGroupsDeniedActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupGroupsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupGroupsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupInstancesRequest() (request *DescribeBackupInstancesRequest) {
    request = &DescribeBackupInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupInstances")
    
    
    return
}

func NewDescribeBackupInstancesResponse() (response *DescribeBackupInstancesResponse) {
    response = &DescribeBackupInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupInstances
// 本接口用来浏览已有受保护实例列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) DescribeBackupInstances(request *DescribeBackupInstancesRequest) (response *DescribeBackupInstancesResponse, err error) {
    return c.DescribeBackupInstancesWithContext(context.Background(), request)
}

// DescribeBackupInstances
// 本接口用来浏览已有受保护实例列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) DescribeBackupInstancesWithContext(ctx context.Context, request *DescribeBackupInstancesRequest) (response *DescribeBackupInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupOverviewGeneralRequest() (request *DescribeBackupOverviewGeneralRequest) {
    request = &DescribeBackupOverviewGeneralRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupOverviewGeneral")
    
    
    return
}

func NewDescribeBackupOverviewGeneralResponse() (response *DescribeBackupOverviewGeneralResponse) {
    response = &DescribeBackupOverviewGeneralResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupOverviewGeneral
// 查询备份概览信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) DescribeBackupOverviewGeneral(request *DescribeBackupOverviewGeneralRequest) (response *DescribeBackupOverviewGeneralResponse, err error) {
    return c.DescribeBackupOverviewGeneralWithContext(context.Background(), request)
}

// DescribeBackupOverviewGeneral
// 查询备份概览信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) DescribeBackupOverviewGeneralWithContext(ctx context.Context, request *DescribeBackupOverviewGeneralRequest) (response *DescribeBackupOverviewGeneralResponse, err error) {
    if request == nil {
        request = NewDescribeBackupOverviewGeneralRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupOverviewGeneral")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupOverviewGeneral require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupOverviewGeneralResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupPlansRequest() (request *DescribeBackupPlansRequest) {
    request = &DescribeBackupPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupPlans")
    
    
    return
}

func NewDescribeBackupPlansResponse() (response *DescribeBackupPlansResponse) {
    response = &DescribeBackupPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupPlans
// 查询整机备份计划
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupPlans(request *DescribeBackupPlansRequest) (response *DescribeBackupPlansResponse, err error) {
    return c.DescribeBackupPlansWithContext(context.Background(), request)
}

// DescribeBackupPlans
// 查询整机备份计划
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBackupPlansWithContext(ctx context.Context, request *DescribeBackupPlansRequest) (response *DescribeBackupPlansResponse, err error) {
    if request == nil {
        request = NewDescribeBackupPlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupPlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupPlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupVaultsRequest() (request *DescribeBackupVaultsRequest) {
    request = &DescribeBackupVaultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupVaults")
    
    
    return
}

func NewDescribeBackupVaultsResponse() (response *DescribeBackupVaultsResponse) {
    response = &DescribeBackupVaultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupVaults
// 查询备份库信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeBackupVaults(request *DescribeBackupVaultsRequest) (response *DescribeBackupVaultsResponse, err error) {
    return c.DescribeBackupVaultsWithContext(context.Background(), request)
}

// DescribeBackupVaults
// 查询备份库信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeBackupVaultsWithContext(ctx context.Context, request *DescribeBackupVaultsRequest) (response *DescribeBackupVaultsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupVaultsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupVaults")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupVaults require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupVaultsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupVaultsDeniedActionsRequest() (request *DescribeBackupVaultsDeniedActionsRequest) {
    request = &DescribeBackupVaultsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeBackupVaultsDeniedActions")
    
    
    return
}

func NewDescribeBackupVaultsDeniedActionsResponse() (response *DescribeBackupVaultsDeniedActionsResponse) {
    response = &DescribeBackupVaultsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupVaultsDeniedActions
// 查询备份库操作掩码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeBackupVaultsDeniedActions(request *DescribeBackupVaultsDeniedActionsRequest) (response *DescribeBackupVaultsDeniedActionsResponse, err error) {
    return c.DescribeBackupVaultsDeniedActionsWithContext(context.Background(), request)
}

// DescribeBackupVaultsDeniedActions
// 查询备份库操作掩码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeBackupVaultsDeniedActionsWithContext(ctx context.Context, request *DescribeBackupVaultsDeniedActionsRequest) (response *DescribeBackupVaultsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupVaultsDeniedActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeBackupVaultsDeniedActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupVaultsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupVaultsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommonBackupPointsRequest() (request *DescribeCommonBackupPointsRequest) {
    request = &DescribeCommonBackupPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeCommonBackupPoints")
    
    
    return
}

func NewDescribeCommonBackupPointsResponse() (response *DescribeCommonBackupPointsResponse) {
    response = &DescribeCommonBackupPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCommonBackupPoints
// 查询共同备份点信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCommonBackupPoints(request *DescribeCommonBackupPointsRequest) (response *DescribeCommonBackupPointsResponse, err error) {
    return c.DescribeCommonBackupPointsWithContext(context.Background(), request)
}

// DescribeCommonBackupPoints
// 查询共同备份点信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCommonBackupPointsWithContext(ctx context.Context, request *DescribeCommonBackupPointsRequest) (response *DescribeCommonBackupPointsResponse, err error) {
    if request == nil {
        request = NewDescribeCommonBackupPointsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeCommonBackupPoints")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCommonBackupPoints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCommonBackupPointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCopyPairsRequest() (request *DescribeCopyPairsRequest) {
    request = &DescribeCopyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeCopyPairs")
    
    
    return
}

func NewDescribeCopyPairsResponse() (response *DescribeCopyPairsResponse) {
    response = &DescribeCopyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCopyPairs
// 本接口用来查询容灾复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeCopyPairs(request *DescribeCopyPairsRequest) (response *DescribeCopyPairsResponse, err error) {
    return c.DescribeCopyPairsWithContext(context.Background(), request)
}

// DescribeCopyPairs
// 本接口用来查询容灾复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeCopyPairsWithContext(ctx context.Context, request *DescribeCopyPairsRequest) (response *DescribeCopyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeCopyPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeCopyPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCopyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCopyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCopyPairsDeniedActionsRequest() (request *DescribeCopyPairsDeniedActionsRequest) {
    request = &DescribeCopyPairsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeCopyPairsDeniedActions")
    
    
    return
}

func NewDescribeCopyPairsDeniedActionsResponse() (response *DescribeCopyPairsDeniedActionsResponse) {
    response = &DescribeCopyPairsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCopyPairsDeniedActions
// 查询复制对掩码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeCopyPairsDeniedActions(request *DescribeCopyPairsDeniedActionsRequest) (response *DescribeCopyPairsDeniedActionsResponse, err error) {
    return c.DescribeCopyPairsDeniedActionsWithContext(context.Background(), request)
}

// DescribeCopyPairsDeniedActions
// 查询复制对掩码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeCopyPairsDeniedActionsWithContext(ctx context.Context, request *DescribeCopyPairsDeniedActionsRequest) (response *DescribeCopyPairsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeCopyPairsDeniedActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeCopyPairsDeniedActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCopyPairsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCopyPairsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoveryDrillGroupsRequest() (request *DescribeDisasterRecoveryDrillGroupsRequest) {
    request = &DescribeDisasterRecoveryDrillGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDisasterRecoveryDrillGroups")
    
    
    return
}

func NewDescribeDisasterRecoveryDrillGroupsResponse() (response *DescribeDisasterRecoveryDrillGroupsResponse) {
    response = &DescribeDisasterRecoveryDrillGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDisasterRecoveryDrillGroups
// 本接口用来查询容灾复制对
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoveryDrillGroups(request *DescribeDisasterRecoveryDrillGroupsRequest) (response *DescribeDisasterRecoveryDrillGroupsResponse, err error) {
    return c.DescribeDisasterRecoveryDrillGroupsWithContext(context.Background(), request)
}

// DescribeDisasterRecoveryDrillGroups
// 本接口用来查询容灾复制对
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoveryDrillGroupsWithContext(ctx context.Context, request *DescribeDisasterRecoveryDrillGroupsRequest) (response *DescribeDisasterRecoveryDrillGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoveryDrillGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDisasterRecoveryDrillGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisasterRecoveryDrillGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoveryDrillGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoveryOverviewRequest() (request *DescribeDisasterRecoveryOverviewRequest) {
    request = &DescribeDisasterRecoveryOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDisasterRecoveryOverview")
    
    
    return
}

func NewDescribeDisasterRecoveryOverviewResponse() (response *DescribeDisasterRecoveryOverviewResponse) {
    response = &DescribeDisasterRecoveryOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDisasterRecoveryOverview
// 查询容灾资源概览
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeDisasterRecoveryOverview(request *DescribeDisasterRecoveryOverviewRequest) (response *DescribeDisasterRecoveryOverviewResponse, err error) {
    return c.DescribeDisasterRecoveryOverviewWithContext(context.Background(), request)
}

// DescribeDisasterRecoveryOverview
// 查询容灾资源概览
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeDisasterRecoveryOverviewWithContext(ctx context.Context, request *DescribeDisasterRecoveryOverviewRequest) (response *DescribeDisasterRecoveryOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoveryOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDisasterRecoveryOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisasterRecoveryOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoveryOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoveryProtectGroupsRequest() (request *DescribeDisasterRecoveryProtectGroupsRequest) {
    request = &DescribeDisasterRecoveryProtectGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDisasterRecoveryProtectGroups")
    
    
    return
}

func NewDescribeDisasterRecoveryProtectGroupsResponse() (response *DescribeDisasterRecoveryProtectGroupsResponse) {
    response = &DescribeDisasterRecoveryProtectGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDisasterRecoveryProtectGroups
// 本接口用来查询容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoveryProtectGroups(request *DescribeDisasterRecoveryProtectGroupsRequest) (response *DescribeDisasterRecoveryProtectGroupsResponse, err error) {
    return c.DescribeDisasterRecoveryProtectGroupsWithContext(context.Background(), request)
}

// DescribeDisasterRecoveryProtectGroups
// 本接口用来查询容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoveryProtectGroupsWithContext(ctx context.Context, request *DescribeDisasterRecoveryProtectGroupsRequest) (response *DescribeDisasterRecoveryProtectGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoveryProtectGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDisasterRecoveryProtectGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisasterRecoveryProtectGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoveryProtectGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoverySitePairsRequest() (request *DescribeDisasterRecoverySitePairsRequest) {
    request = &DescribeDisasterRecoverySitePairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDisasterRecoverySitePairs")
    
    
    return
}

func NewDescribeDisasterRecoverySitePairsResponse() (response *DescribeDisasterRecoverySitePairsResponse) {
    response = &DescribeDisasterRecoverySitePairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDisasterRecoverySitePairs
// 本接口用来查询容灾站点对
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
func (c *Client) DescribeDisasterRecoverySitePairs(request *DescribeDisasterRecoverySitePairsRequest) (response *DescribeDisasterRecoverySitePairsResponse, err error) {
    return c.DescribeDisasterRecoverySitePairsWithContext(context.Background(), request)
}

// DescribeDisasterRecoverySitePairs
// 本接口用来查询容灾站点对
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
func (c *Client) DescribeDisasterRecoverySitePairsWithContext(ctx context.Context, request *DescribeDisasterRecoverySitePairsRequest) (response *DescribeDisasterRecoverySitePairsResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverySitePairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDisasterRecoverySitePairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisasterRecoverySitePairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoverySitePairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoverySitePairsDeniedActionsRequest() (request *DescribeDisasterRecoverySitePairsDeniedActionsRequest) {
    request = &DescribeDisasterRecoverySitePairsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDisasterRecoverySitePairsDeniedActions")
    
    
    return
}

func NewDescribeDisasterRecoverySitePairsDeniedActionsResponse() (response *DescribeDisasterRecoverySitePairsDeniedActionsResponse) {
    response = &DescribeDisasterRecoverySitePairsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDisasterRecoverySitePairsDeniedActions
// 查询指定容灾站点对当前不允许执行的操作列表（操作掩码）。前端在展示容灾策略操作菜单时，可基于该接口返回结果灰化或屏蔽相应入口，并向用户提示原因（错误码 + 错误信息）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoverySitePairsDeniedActions(request *DescribeDisasterRecoverySitePairsDeniedActionsRequest) (response *DescribeDisasterRecoverySitePairsDeniedActionsResponse, err error) {
    return c.DescribeDisasterRecoverySitePairsDeniedActionsWithContext(context.Background(), request)
}

// DescribeDisasterRecoverySitePairsDeniedActions
// 查询指定容灾站点对当前不允许执行的操作列表（操作掩码）。前端在展示容灾策略操作菜单时，可基于该接口返回结果灰化或屏蔽相应入口，并向用户提示原因（错误码 + 错误信息）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoverySitePairsDeniedActionsWithContext(ctx context.Context, request *DescribeDisasterRecoverySitePairsDeniedActionsRequest) (response *DescribeDisasterRecoverySitePairsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverySitePairsDeniedActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDisasterRecoverySitePairsDeniedActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisasterRecoverySitePairsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoverySitePairsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoverySupportRegionRequest() (request *DescribeDisasterRecoverySupportRegionRequest) {
    request = &DescribeDisasterRecoverySupportRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDisasterRecoverySupportRegion")
    
    
    return
}

func NewDescribeDisasterRecoverySupportRegionResponse() (response *DescribeDisasterRecoverySupportRegionResponse) {
    response = &DescribeDisasterRecoverySupportRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDisasterRecoverySupportRegion
// 查询当前地域支持容灾的生产地域配置列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoverySupportRegion(request *DescribeDisasterRecoverySupportRegionRequest) (response *DescribeDisasterRecoverySupportRegionResponse, err error) {
    return c.DescribeDisasterRecoverySupportRegionWithContext(context.Background(), request)
}

// DescribeDisasterRecoverySupportRegion
// 查询当前地域支持容灾的生产地域配置列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisasterRecoverySupportRegionWithContext(ctx context.Context, request *DescribeDisasterRecoverySupportRegionRequest) (response *DescribeDisasterRecoverySupportRegionResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverySupportRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDisasterRecoverySupportRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisasterRecoverySupportRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoverySupportRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
    request = &DescribeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDisks")
    
    
    return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
    response = &DescribeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDisks
// 本接口用来查询容灾云硬盘的详情，如系统盘的镜像格式。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    return c.DescribeDisksWithContext(context.Background(), request)
}

// DescribeDisks
// 本接口用来查询容灾云硬盘的详情，如系统盘的镜像格式。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDisksWithContext(ctx context.Context, request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    if request == nil {
        request = NewDescribeDisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDrillPairsRequest() (request *DescribeDrillPairsRequest) {
    request = &DescribeDrillPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDrillPairs")
    
    
    return
}

func NewDescribeDrillPairsResponse() (response *DescribeDrillPairsResponse) {
    response = &DescribeDrillPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDrillPairs
// 查询演练对列表
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDrillPairs(request *DescribeDrillPairsRequest) (response *DescribeDrillPairsResponse, err error) {
    return c.DescribeDrillPairsWithContext(context.Background(), request)
}

// DescribeDrillPairs
// 查询演练对列表
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeDrillPairsWithContext(ctx context.Context, request *DescribeDrillPairsRequest) (response *DescribeDrillPairsResponse, err error) {
    if request == nil {
        request = NewDescribeDrillPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDrillPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDrillPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDrillPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDrillPairsDeniedActionsRequest() (request *DescribeDrillPairsDeniedActionsRequest) {
    request = &DescribeDrillPairsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeDrillPairsDeniedActions")
    
    
    return
}

func NewDescribeDrillPairsDeniedActionsResponse() (response *DescribeDrillPairsDeniedActionsResponse) {
    response = &DescribeDrillPairsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDrillPairsDeniedActions
// 查询演练操作掩码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDrillPairsDeniedActions(request *DescribeDrillPairsDeniedActionsRequest) (response *DescribeDrillPairsDeniedActionsResponse, err error) {
    return c.DescribeDrillPairsDeniedActionsWithContext(context.Background(), request)
}

// DescribeDrillPairsDeniedActions
// 查询演练操作掩码
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDrillPairsDeniedActionsWithContext(ctx context.Context, request *DescribeDrillPairsDeniedActionsRequest) (response *DescribeDrillPairsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeDrillPairsDeniedActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeDrillPairsDeniedActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDrillPairsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDrillPairsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileBackupObjectsRequest() (request *DescribeFileBackupObjectsRequest) {
    request = &DescribeFileBackupObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeFileBackupObjects")
    
    
    return
}

func NewDescribeFileBackupObjectsResponse() (response *DescribeFileBackupObjectsResponse) {
    response = &DescribeFileBackupObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileBackupObjects
// 本接口用来浏览已有备份目录/文件内容
//
// 可能返回的错误码:
//  INTERNALERROR_SNAPSHOTTREEQUERYFAILED = "InternalError.SnapshotTreeQueryFailed"
//  INVALIDPARAMETERVALUE_INVALIDORDERDIRECTION = "InvalidParameterValue.InvalidOrderDirection"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
//  UNSUPPORTEDOPERATION_SNAPSHOTUNAVAILABLE = "UnsupportedOperation.SnapshotUnavailable"
func (c *Client) DescribeFileBackupObjects(request *DescribeFileBackupObjectsRequest) (response *DescribeFileBackupObjectsResponse, err error) {
    return c.DescribeFileBackupObjectsWithContext(context.Background(), request)
}

// DescribeFileBackupObjects
// 本接口用来浏览已有备份目录/文件内容
//
// 可能返回的错误码:
//  INTERNALERROR_SNAPSHOTTREEQUERYFAILED = "InternalError.SnapshotTreeQueryFailed"
//  INVALIDPARAMETERVALUE_INVALIDORDERDIRECTION = "InvalidParameterValue.InvalidOrderDirection"
//  INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
//  UNSUPPORTEDOPERATION_SNAPSHOTUNAVAILABLE = "UnsupportedOperation.SnapshotUnavailable"
func (c *Client) DescribeFileBackupObjectsWithContext(ctx context.Context, request *DescribeFileBackupObjectsRequest) (response *DescribeFileBackupObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeFileBackupObjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeFileBackupObjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileBackupObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileBackupObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileBackupPlansRequest() (request *DescribeFileBackupPlansRequest) {
    request = &DescribeFileBackupPlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeFileBackupPlans")
    
    
    return
}

func NewDescribeFileBackupPlansResponse() (response *DescribeFileBackupPlansResponse) {
    response = &DescribeFileBackupPlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileBackupPlans
// 本接口用来浏览已有备份计划内容
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileBackupPlans(request *DescribeFileBackupPlansRequest) (response *DescribeFileBackupPlansResponse, err error) {
    return c.DescribeFileBackupPlansWithContext(context.Background(), request)
}

// DescribeFileBackupPlans
// 本接口用来浏览已有备份计划内容
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileBackupPlansWithContext(ctx context.Context, request *DescribeFileBackupPlansRequest) (response *DescribeFileBackupPlansResponse, err error) {
    if request == nil {
        request = NewDescribeFileBackupPlansRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeFileBackupPlans")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileBackupPlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileBackupPlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileBackupsRequest() (request *DescribeFileBackupsRequest) {
    request = &DescribeFileBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeFileBackups")
    
    
    return
}

func NewDescribeFileBackupsResponse() (response *DescribeFileBackupsResponse) {
    response = &DescribeFileBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileBackups
// 本接口用来浏览已有备份点详情
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileBackups(request *DescribeFileBackupsRequest) (response *DescribeFileBackupsResponse, err error) {
    return c.DescribeFileBackupsWithContext(context.Background(), request)
}

// DescribeFileBackups
// 本接口用来浏览已有备份点详情
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileBackupsWithContext(ctx context.Context, request *DescribeFileBackupsRequest) (response *DescribeFileBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeFileBackupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeFileBackups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileBackupsDeniedActionsRequest() (request *DescribeFileBackupsDeniedActionsRequest) {
    request = &DescribeFileBackupsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeFileBackupsDeniedActions")
    
    
    return
}

func NewDescribeFileBackupsDeniedActionsResponse() (response *DescribeFileBackupsDeniedActionsResponse) {
    response = &DescribeFileBackupsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileBackupsDeniedActions
// 本接口用来查询备份操作掩码
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileBackupsDeniedActions(request *DescribeFileBackupsDeniedActionsRequest) (response *DescribeFileBackupsDeniedActionsResponse, err error) {
    return c.DescribeFileBackupsDeniedActionsWithContext(context.Background(), request)
}

// DescribeFileBackupsDeniedActions
// 本接口用来查询备份操作掩码
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileBackupsDeniedActionsWithContext(ctx context.Context, request *DescribeFileBackupsDeniedActionsRequest) (response *DescribeFileBackupsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeFileBackupsDeniedActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeFileBackupsDeniedActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileBackupsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileBackupsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileRestoreTasksRequest() (request *DescribeFileRestoreTasksRequest) {
    request = &DescribeFileRestoreTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeFileRestoreTasks")
    
    
    return
}

func NewDescribeFileRestoreTasksResponse() (response *DescribeFileRestoreTasksResponse) {
    response = &DescribeFileRestoreTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileRestoreTasks
// 查询备份恢复任务列表
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileRestoreTasks(request *DescribeFileRestoreTasksRequest) (response *DescribeFileRestoreTasksResponse, err error) {
    return c.DescribeFileRestoreTasksWithContext(context.Background(), request)
}

// DescribeFileRestoreTasks
// 查询备份恢复任务列表
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeFileRestoreTasksWithContext(ctx context.Context, request *DescribeFileRestoreTasksRequest) (response *DescribeFileRestoreTasksResponse, err error) {
    if request == nil {
        request = NewDescribeFileRestoreTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeFileRestoreTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileRestoreTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileRestoreTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeJobs")
    
    
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeJobs
// 本接口用于Agent查询相关Agent任务信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// 本接口用于Agent查询相关Agent任务信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
func (c *Client) DescribeJobsWithContext(ctx context.Context, request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePriceCreateCopyPairsRequest() (request *DescribePriceCreateCopyPairsRequest) {
    request = &DescribePriceCreateCopyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribePriceCreateCopyPairs")
    
    
    return
}

func NewDescribePriceCreateCopyPairsResponse() (response *DescribePriceCreateCopyPairsResponse) {
    response = &DescribePriceCreateCopyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePriceCreateCopyPairs
// 本接口（DescribePriceCreateCopyPairs）用于查询创建容灾复制对的价格。支持批量询价，入参为每个复制对的盘容量数组，返回与入参一一对应的后付费每小时价格。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribePriceCreateCopyPairs(request *DescribePriceCreateCopyPairsRequest) (response *DescribePriceCreateCopyPairsResponse, err error) {
    return c.DescribePriceCreateCopyPairsWithContext(context.Background(), request)
}

// DescribePriceCreateCopyPairs
// 本接口（DescribePriceCreateCopyPairs）用于查询创建容灾复制对的价格。支持批量询价，入参为每个复制对的盘容量数组，返回与入参一一对应的后付费每小时价格。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribePriceCreateCopyPairsWithContext(ctx context.Context, request *DescribePriceCreateCopyPairsRequest) (response *DescribePriceCreateCopyPairsResponse, err error) {
    if request == nil {
        request = NewDescribePriceCreateCopyPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribePriceCreateCopyPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePriceCreateCopyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePriceCreateCopyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProtectGroupsDeniedActionsRequest() (request *DescribeProtectGroupsDeniedActionsRequest) {
    request = &DescribeProtectGroupsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeProtectGroupsDeniedActions")
    
    
    return
}

func NewDescribeProtectGroupsDeniedActionsResponse() (response *DescribeProtectGroupsDeniedActionsResponse) {
    response = &DescribeProtectGroupsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProtectGroupsDeniedActions
// 查询保护组操作掩码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeProtectGroupsDeniedActions(request *DescribeProtectGroupsDeniedActionsRequest) (response *DescribeProtectGroupsDeniedActionsResponse, err error) {
    return c.DescribeProtectGroupsDeniedActionsWithContext(context.Background(), request)
}

// DescribeProtectGroupsDeniedActions
// 查询保护组操作掩码
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeProtectGroupsDeniedActionsWithContext(ctx context.Context, request *DescribeProtectGroupsDeniedActionsRequest) (response *DescribeProtectGroupsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeProtectGroupsDeniedActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeProtectGroupsDeniedActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProtectGroupsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProtectGroupsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProtectedInstancesRequest() (request *DescribeProtectedInstancesRequest) {
    request = &DescribeProtectedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeProtectedInstances")
    
    
    return
}

func NewDescribeProtectedInstancesResponse() (response *DescribeProtectedInstancesResponse) {
    response = &DescribeProtectedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProtectedInstances
// 本接口用来浏览已有受保护实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeProtectedInstances(request *DescribeProtectedInstancesRequest) (response *DescribeProtectedInstancesResponse, err error) {
    return c.DescribeProtectedInstancesWithContext(context.Background(), request)
}

// DescribeProtectedInstances
// 本接口用来浏览已有受保护实例列表
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
func (c *Client) DescribeProtectedInstancesWithContext(ctx context.Context, request *DescribeProtectedInstancesRequest) (response *DescribeProtectedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeProtectedInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeProtectedInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProtectedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProtectedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupMappingsRequest() (request *DescribeSecurityGroupMappingsRequest) {
    request = &DescribeSecurityGroupMappingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeSecurityGroupMappings")
    
    
    return
}

func NewDescribeSecurityGroupMappingsResponse() (response *DescribeSecurityGroupMappingsResponse) {
    response = &DescribeSecurityGroupMappingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroupMappings
// 本接口用于查询安全组映射列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeSecurityGroupMappings(request *DescribeSecurityGroupMappingsRequest) (response *DescribeSecurityGroupMappingsResponse, err error) {
    return c.DescribeSecurityGroupMappingsWithContext(context.Background(), request)
}

// DescribeSecurityGroupMappings
// 本接口用于查询安全组映射列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeSecurityGroupMappingsWithContext(ctx context.Context, request *DescribeSecurityGroupMappingsRequest) (response *DescribeSecurityGroupMappingsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupMappingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeSecurityGroupMappings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroupMappings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupMappingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcMappingsRequest() (request *DescribeVpcMappingsRequest) {
    request = &DescribeVpcMappingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "DescribeVpcMappings")
    
    
    return
}

func NewDescribeVpcMappingsResponse() (response *DescribeVpcMappingsResponse) {
    response = &DescribeVpcMappingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVpcMappings
// 本接口用来查询站点对的vpc映射信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeVpcMappings(request *DescribeVpcMappingsRequest) (response *DescribeVpcMappingsResponse, err error) {
    return c.DescribeVpcMappingsWithContext(context.Background(), request)
}

// DescribeVpcMappings
// 本接口用来查询站点对的vpc映射信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) DescribeVpcMappingsWithContext(ctx context.Context, request *DescribeVpcMappingsRequest) (response *DescribeVpcMappingsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcMappingsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "DescribeVpcMappings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVpcMappings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVpcMappingsResponse()
    err = c.Send(request, response)
    return
}

func NewFinishFailoverCopyPairsRequest() (request *FinishFailoverCopyPairsRequest) {
    request = &FinishFailoverCopyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "FinishFailoverCopyPairs")
    
    
    return
}

func NewFinishFailoverCopyPairsResponse() (response *FinishFailoverCopyPairsResponse) {
    response = &FinishFailoverCopyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FinishFailoverCopyPairs
// 完成切换
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
func (c *Client) FinishFailoverCopyPairs(request *FinishFailoverCopyPairsRequest) (response *FinishFailoverCopyPairsResponse, err error) {
    return c.FinishFailoverCopyPairsWithContext(context.Background(), request)
}

// FinishFailoverCopyPairs
// 完成切换
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
func (c *Client) FinishFailoverCopyPairsWithContext(ctx context.Context, request *FinishFailoverCopyPairsRequest) (response *FinishFailoverCopyPairsResponse, err error) {
    if request == nil {
        request = NewFinishFailoverCopyPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "FinishFailoverCopyPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FinishFailoverCopyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewFinishFailoverCopyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoBackupPolicyAttributeRequest() (request *ModifyAutoBackupPolicyAttributeRequest) {
    request = &ModifyAutoBackupPolicyAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyAutoBackupPolicyAttribute")
    
    
    return
}

func NewModifyAutoBackupPolicyAttributeResponse() (response *ModifyAutoBackupPolicyAttributeResponse) {
    response = &ModifyAutoBackupPolicyAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoBackupPolicyAttribute
// 修改备份策略
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_POLICYNOTAVAILABLE = "InvalidParameterValue.PolicyNotAvailable"
//  RESOURCENOTFOUND_AUTOBACKUPPOLICYNOTFOUND = "ResourceNotFound.AutoBackupPolicyNotFound"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyAutoBackupPolicyAttribute(request *ModifyAutoBackupPolicyAttributeRequest) (response *ModifyAutoBackupPolicyAttributeResponse, err error) {
    return c.ModifyAutoBackupPolicyAttributeWithContext(context.Background(), request)
}

// ModifyAutoBackupPolicyAttribute
// 修改备份策略
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_POLICYNOTAVAILABLE = "InvalidParameterValue.PolicyNotAvailable"
//  RESOURCENOTFOUND_AUTOBACKUPPOLICYNOTFOUND = "ResourceNotFound.AutoBackupPolicyNotFound"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyAutoBackupPolicyAttributeWithContext(ctx context.Context, request *ModifyAutoBackupPolicyAttributeRequest) (response *ModifyAutoBackupPolicyAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAutoBackupPolicyAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyAutoBackupPolicyAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoBackupPolicyAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoBackupPolicyAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupAttributeRequest() (request *ModifyBackupAttributeRequest) {
    request = &ModifyBackupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyBackupAttribute")
    
    
    return
}

func NewModifyBackupAttributeResponse() (response *ModifyBackupAttributeResponse) {
    response = &ModifyBackupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupAttribute
// 删除备份组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyBackupAttribute(request *ModifyBackupAttributeRequest) (response *ModifyBackupAttributeResponse, err error) {
    return c.ModifyBackupAttributeWithContext(context.Background(), request)
}

// ModifyBackupAttribute
// 删除备份组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyBackupAttributeWithContext(ctx context.Context, request *ModifyBackupAttributeRequest) (response *ModifyBackupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBackupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyBackupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupVaultAttributeRequest() (request *ModifyBackupVaultAttributeRequest) {
    request = &ModifyBackupVaultAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyBackupVaultAttribute")
    
    
    return
}

func NewModifyBackupVaultAttributeResponse() (response *ModifyBackupVaultAttributeResponse) {
    response = &ModifyBackupVaultAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupVaultAttribute
// 修改备份库信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyBackupVaultAttribute(request *ModifyBackupVaultAttributeRequest) (response *ModifyBackupVaultAttributeResponse, err error) {
    return c.ModifyBackupVaultAttributeWithContext(context.Background(), request)
}

// ModifyBackupVaultAttribute
// 修改备份库信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyBackupVaultAttributeWithContext(ctx context.Context, request *ModifyBackupVaultAttributeRequest) (response *ModifyBackupVaultAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBackupVaultAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyBackupVaultAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupVaultAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupVaultAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCopyPairAttributeRequest() (request *ModifyCopyPairAttributeRequest) {
    request = &ModifyCopyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyCopyPairAttribute")
    
    
    return
}

func NewModifyCopyPairAttributeResponse() (response *ModifyCopyPairAttributeResponse) {
    response = &ModifyCopyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCopyPairAttribute
// 修改容灾复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyCopyPairAttribute(request *ModifyCopyPairAttributeRequest) (response *ModifyCopyPairAttributeResponse, err error) {
    return c.ModifyCopyPairAttributeWithContext(context.Background(), request)
}

// ModifyCopyPairAttribute
// 修改容灾复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyCopyPairAttributeWithContext(ctx context.Context, request *ModifyCopyPairAttributeRequest) (response *ModifyCopyPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCopyPairAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyCopyPairAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCopyPairAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCopyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDrillGroupAttributeRequest() (request *ModifyDrillGroupAttributeRequest) {
    request = &ModifyDrillGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyDrillGroupAttribute")
    
    
    return
}

func NewModifyDrillGroupAttributeResponse() (response *ModifyDrillGroupAttributeResponse) {
    response = &ModifyDrillGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDrillGroupAttribute
// 修改演练组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DRILLNOTEXIST = "ResourceNotFound.DrillNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyDrillGroupAttribute(request *ModifyDrillGroupAttributeRequest) (response *ModifyDrillGroupAttributeResponse, err error) {
    return c.ModifyDrillGroupAttributeWithContext(context.Background(), request)
}

// ModifyDrillGroupAttribute
// 修改演练组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DRILLNOTEXIST = "ResourceNotFound.DrillNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyDrillGroupAttributeWithContext(ctx context.Context, request *ModifyDrillGroupAttributeRequest) (response *ModifyDrillGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDrillGroupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyDrillGroupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDrillGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDrillGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDrillPairAttributeRequest() (request *ModifyDrillPairAttributeRequest) {
    request = &ModifyDrillPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyDrillPairAttribute")
    
    
    return
}

func NewModifyDrillPairAttributeResponse() (response *ModifyDrillPairAttributeResponse) {
    response = &ModifyDrillPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDrillPairAttribute
// 修改演练
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DRILLNOTEXIST = "ResourceNotFound.DrillNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyDrillPairAttribute(request *ModifyDrillPairAttributeRequest) (response *ModifyDrillPairAttributeResponse, err error) {
    return c.ModifyDrillPairAttributeWithContext(context.Background(), request)
}

// ModifyDrillPairAttribute
// 修改演练
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_DRILLNOTEXIST = "ResourceNotFound.DrillNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyDrillPairAttributeWithContext(ctx context.Context, request *ModifyDrillPairAttributeRequest) (response *ModifyDrillPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDrillPairAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyDrillPairAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDrillPairAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDrillPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileBackupAttributeRequest() (request *ModifyFileBackupAttributeRequest) {
    request = &ModifyFileBackupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyFileBackupAttribute")
    
    
    return
}

func NewModifyFileBackupAttributeResponse() (response *ModifyFileBackupAttributeResponse) {
    response = &ModifyFileBackupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFileBackupAttribute
// 修改文件备份信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILEBACKUPNOTEXIST = "ResourceNotFound.FileBackupNotExist"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
func (c *Client) ModifyFileBackupAttribute(request *ModifyFileBackupAttributeRequest) (response *ModifyFileBackupAttributeResponse, err error) {
    return c.ModifyFileBackupAttributeWithContext(context.Background(), request)
}

// ModifyFileBackupAttribute
// 修改文件备份信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILEBACKUPNOTEXIST = "ResourceNotFound.FileBackupNotExist"
//  UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"
func (c *Client) ModifyFileBackupAttributeWithContext(ctx context.Context, request *ModifyFileBackupAttributeRequest) (response *ModifyFileBackupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyFileBackupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyFileBackupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFileBackupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFileBackupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileBackupPlanRequest() (request *ModifyFileBackupPlanRequest) {
    request = &ModifyFileBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyFileBackupPlan")
    
    
    return
}

func NewModifyFileBackupPlanResponse() (response *ModifyFileBackupPlanResponse) {
    response = &ModifyFileBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFileBackupPlan
// 本接口用于修改已有的备份计划配置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  RESOURCENOTFOUND_BACKUPPLANNOTEXIST = "ResourceNotFound.BackupPlanNotExist"
func (c *Client) ModifyFileBackupPlan(request *ModifyFileBackupPlanRequest) (response *ModifyFileBackupPlanResponse, err error) {
    return c.ModifyFileBackupPlanWithContext(context.Background(), request)
}

// ModifyFileBackupPlan
// 本接口用于修改已有的备份计划配置
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"
//  RESOURCENOTFOUND_BACKUPPLANNOTEXIST = "ResourceNotFound.BackupPlanNotExist"
func (c *Client) ModifyFileBackupPlanWithContext(ctx context.Context, request *ModifyFileBackupPlanRequest) (response *ModifyFileBackupPlanResponse, err error) {
    if request == nil {
        request = NewModifyFileBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyFileBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFileBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFileBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProtectGroupAttributeRequest() (request *ModifyProtectGroupAttributeRequest) {
    request = &ModifyProtectGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifyProtectGroupAttribute")
    
    
    return
}

func NewModifyProtectGroupAttributeResponse() (response *ModifyProtectGroupAttributeResponse) {
    response = &ModifyProtectGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProtectGroupAttribute
// 修改容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyProtectGroupAttribute(request *ModifyProtectGroupAttributeRequest) (response *ModifyProtectGroupAttributeResponse, err error) {
    return c.ModifyProtectGroupAttributeWithContext(context.Background(), request)
}

// ModifyProtectGroupAttribute
// 修改容灾保护组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyProtectGroupAttributeWithContext(ctx context.Context, request *ModifyProtectGroupAttributeRequest) (response *ModifyProtectGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyProtectGroupAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifyProtectGroupAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProtectGroupAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProtectGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySitePairAttributeRequest() (request *ModifySitePairAttributeRequest) {
    request = &ModifySitePairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ModifySitePairAttribute")
    
    
    return
}

func NewModifySitePairAttributeResponse() (response *ModifySitePairAttributeResponse) {
    response = &ModifySitePairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySitePairAttribute
// 修改容灾站点对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) ModifySitePairAttribute(request *ModifySitePairAttributeRequest) (response *ModifySitePairAttributeResponse, err error) {
    return c.ModifySitePairAttributeWithContext(context.Background(), request)
}

// ModifySitePairAttribute
// 修改容灾站点对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"
func (c *Client) ModifySitePairAttributeWithContext(ctx context.Context, request *ModifySitePairAttributeRequest) (response *ModifySitePairAttributeResponse, err error) {
    if request == nil {
        request = NewModifySitePairAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ModifySitePairAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySitePairAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySitePairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewReportAgentMetricsRequest() (request *ReportAgentMetricsRequest) {
    request = &ReportAgentMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ReportAgentMetrics")
    
    
    return
}

func NewReportAgentMetricsResponse() (response *ReportAgentMetricsResponse) {
    response = &ReportAgentMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReportAgentMetrics
// 本接口用于上报Agent指标信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GATEWAY = "ResourceNotFound.Gateway"
func (c *Client) ReportAgentMetrics(request *ReportAgentMetricsRequest) (response *ReportAgentMetricsResponse, err error) {
    return c.ReportAgentMetricsWithContext(context.Background(), request)
}

// ReportAgentMetrics
// 本接口用于上报Agent指标信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GATEWAY = "ResourceNotFound.Gateway"
func (c *Client) ReportAgentMetricsWithContext(ctx context.Context, request *ReportAgentMetricsRequest) (response *ReportAgentMetricsResponse, err error) {
    if request == nil {
        request = NewReportAgentMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ReportAgentMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportAgentMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportAgentMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewReportGatewayHeartbeatRequest() (request *ReportGatewayHeartbeatRequest) {
    request = &ReportGatewayHeartbeatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ReportGatewayHeartbeat")
    
    
    return
}

func NewReportGatewayHeartbeatResponse() (response *ReportGatewayHeartbeatResponse) {
    response = &ReportGatewayHeartbeatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReportGatewayHeartbeat
// 本接口用于Agent心跳上报
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GATEWAY = "ResourceNotFound.Gateway"
func (c *Client) ReportGatewayHeartbeat(request *ReportGatewayHeartbeatRequest) (response *ReportGatewayHeartbeatResponse, err error) {
    return c.ReportGatewayHeartbeatWithContext(context.Background(), request)
}

// ReportGatewayHeartbeat
// 本接口用于Agent心跳上报
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GATEWAY = "ResourceNotFound.Gateway"
func (c *Client) ReportGatewayHeartbeatWithContext(ctx context.Context, request *ReportGatewayHeartbeatRequest) (response *ReportGatewayHeartbeatResponse, err error) {
    if request == nil {
        request = NewReportGatewayHeartbeatRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ReportGatewayHeartbeat")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportGatewayHeartbeat require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportGatewayHeartbeatResponse()
    err = c.Send(request, response)
    return
}

func NewReportJobProgressRequest() (request *ReportJobProgressRequest) {
    request = &ReportJobProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "ReportJobProgress")
    
    
    return
}

func NewReportJobProgressResponse() (response *ReportJobProgressResponse) {
    response = &ReportJobProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReportJobProgress
// 本接口用于上报Agent任务信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  INVALIDPARAMETERVALUE_INVALIDPROGRESS = "InvalidParameterValue.InvalidProgress"
//  INVALIDPARAMETERVALUE_INVALIDSTATUSTRANSITION = "InvalidParameterValue.InvalidStatusTransition"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) ReportJobProgress(request *ReportJobProgressRequest) (response *ReportJobProgressResponse, err error) {
    return c.ReportJobProgressWithContext(context.Background(), request)
}

// ReportJobProgress
// 本接口用于上报Agent任务信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"
//  INVALIDPARAMETERVALUE_INVALIDPROGRESS = "InvalidParameterValue.InvalidProgress"
//  INVALIDPARAMETERVALUE_INVALIDSTATUSTRANSITION = "InvalidParameterValue.InvalidStatusTransition"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) ReportJobProgressWithContext(ctx context.Context, request *ReportJobProgressRequest) (response *ReportJobProgressResponse, err error) {
    if request == nil {
        request = NewReportJobProgressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "ReportJobProgress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportJobProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportJobProgressResponse()
    err = c.Send(request, response)
    return
}

func NewRunCopyPairTasksRequest() (request *RunCopyPairTasksRequest) {
    request = &RunCopyPairTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "RunCopyPairTasks")
    
    
    return
}

func NewRunCopyPairTasksResponse() (response *RunCopyPairTasksResponse) {
    response = &RunCopyPairTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunCopyPairTasks
// 启动复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
func (c *Client) RunCopyPairTasks(request *RunCopyPairTasksRequest) (response *RunCopyPairTasksResponse, err error) {
    return c.RunCopyPairTasksWithContext(context.Background(), request)
}

// RunCopyPairTasks
// 启动复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
func (c *Client) RunCopyPairTasksWithContext(ctx context.Context, request *RunCopyPairTasksRequest) (response *RunCopyPairTasksResponse, err error) {
    if request == nil {
        request = NewRunCopyPairTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "RunCopyPairTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunCopyPairTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunCopyPairTasksResponse()
    err = c.Send(request, response)
    return
}

func NewRunFailoverCopyPairsRequest() (request *RunFailoverCopyPairsRequest) {
    request = &RunFailoverCopyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "RunFailoverCopyPairs")
    
    
    return
}

func NewRunFailoverCopyPairsResponse() (response *RunFailoverCopyPairsResponse) {
    response = &RunFailoverCopyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunFailoverCopyPairs
// 故障切换
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
//  UNSUPPORTEDOPERATION_HAVENOPROTECTIONPOINT = "UnsupportedOperation.HaveNoProtectionPoint"
func (c *Client) RunFailoverCopyPairs(request *RunFailoverCopyPairsRequest) (response *RunFailoverCopyPairsResponse, err error) {
    return c.RunFailoverCopyPairsWithContext(context.Background(), request)
}

// RunFailoverCopyPairs
// 故障切换
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
//  UNSUPPORTEDOPERATION_HAVENOPROTECTIONPOINT = "UnsupportedOperation.HaveNoProtectionPoint"
func (c *Client) RunFailoverCopyPairsWithContext(ctx context.Context, request *RunFailoverCopyPairsRequest) (response *RunFailoverCopyPairsResponse, err error) {
    if request == nil {
        request = NewRunFailoverCopyPairsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "RunFailoverCopyPairs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunFailoverCopyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunFailoverCopyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstancesWithBackupGroupRequest() (request *RunInstancesWithBackupGroupRequest) {
    request = &RunInstancesWithBackupGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "RunInstancesWithBackupGroup")
    
    
    return
}

func NewRunInstancesWithBackupGroupResponse() (response *RunInstancesWithBackupGroupResponse) {
    response = &RunInstancesWithBackupGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunInstancesWithBackupGroup
// 备份组新建云服务器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_SHOULDCONVERTBACKUPTOIMAGE = "InvalidParameter.ShouldConvertBackupToImage"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) RunInstancesWithBackupGroup(request *RunInstancesWithBackupGroupRequest) (response *RunInstancesWithBackupGroupResponse, err error) {
    return c.RunInstancesWithBackupGroupWithContext(context.Background(), request)
}

// RunInstancesWithBackupGroup
// 备份组新建云服务器
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_SHOULDCONVERTBACKUPTOIMAGE = "InvalidParameter.ShouldConvertBackupToImage"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) RunInstancesWithBackupGroupWithContext(ctx context.Context, request *RunInstancesWithBackupGroupRequest) (response *RunInstancesWithBackupGroupResponse, err error) {
    if request == nil {
        request = NewRunInstancesWithBackupGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "RunInstancesWithBackupGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunInstancesWithBackupGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunInstancesWithBackupGroupResponse()
    err = c.Send(request, response)
    return
}

func NewStopCopyPairTasksRequest() (request *StopCopyPairTasksRequest) {
    request = &StopCopyPairTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "StopCopyPairTasks")
    
    
    return
}

func NewStopCopyPairTasksResponse() (response *StopCopyPairTasksResponse) {
    response = &StopCopyPairTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCopyPairTasks
// 停止复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
func (c *Client) StopCopyPairTasks(request *StopCopyPairTasksRequest) (response *StopCopyPairTasksResponse, err error) {
    return c.StopCopyPairTasksWithContext(context.Background(), request)
}

// StopCopyPairTasks
// 停止复制对
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"
//  RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"
//  UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"
func (c *Client) StopCopyPairTasksWithContext(ctx context.Context, request *StopCopyPairTasksRequest) (response *StopCopyPairTasksResponse, err error) {
    if request == nil {
        request = NewStopCopyPairTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "StopCopyPairTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCopyPairTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCopyPairTasksResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindAutoBackupPolicyRequest() (request *UnbindAutoBackupPolicyRequest) {
    request = &UnbindAutoBackupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bdrc", APIVersion, "UnbindAutoBackupPolicy")
    
    
    return
}

func NewUnbindAutoBackupPolicyResponse() (response *UnbindAutoBackupPolicyResponse) {
    response = &UnbindAutoBackupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindAutoBackupPolicy
// 将实例从备份策略上解绑
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) UnbindAutoBackupPolicy(request *UnbindAutoBackupPolicyRequest) (response *UnbindAutoBackupPolicyResponse, err error) {
    return c.UnbindAutoBackupPolicyWithContext(context.Background(), request)
}

// UnbindAutoBackupPolicy
// 将实例从备份策略上解绑
//
// 可能返回的错误码:
//  INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) UnbindAutoBackupPolicyWithContext(ctx context.Context, request *UnbindAutoBackupPolicyRequest) (response *UnbindAutoBackupPolicyResponse, err error) {
    if request == nil {
        request = NewUnbindAutoBackupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bdrc", APIVersion, "UnbindAutoBackupPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindAutoBackupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindAutoBackupPolicyResponse()
    err = c.Send(request, response)
    return
}
