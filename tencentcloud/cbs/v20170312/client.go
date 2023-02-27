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

package v20170312

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewApplyDiskBackupRequest() (request *ApplyDiskBackupRequest) {
    request = &ApplyDiskBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ApplyDiskBackup")
    
    
    return
}

func NewApplyDiskBackupResponse() (response *ApplyDiskBackupResponse) {
    response = &ApplyDiskBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyDiskBackup
// 本接口（ApplyDiskBackup）用于回滚备份点到原云硬盘。
//
// 
//
// * 仅支持回滚到原云硬盘上。对于数据盘备份点，如果您需要复制备份点数据到其它云硬盘上，请先使用 CreateSnapshot 将备份点转换为快照，然后使用 CreateDisks 接口创建新的弹性云硬盘，将快照数据复制到新购云硬盘上。
//
// * 用于回滚的备份点必须处于NORMAL状态。备份点状态可以通过DescribeDiskBackups接口查询，见输出参数中BackupState字段解释。
//
// * 如果是弹性云硬盘，则云硬盘必须处于未挂载状态，云硬盘挂载状态可以通过DescribeDisks接口查询，见Attached字段解释；如果是随实例一起购买的非弹性云硬盘，则实例必须处于关机状态，实例状态可以通过DescribeInstancesStatus接口查询。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) ApplyDiskBackup(request *ApplyDiskBackupRequest) (response *ApplyDiskBackupResponse, err error) {
    return c.ApplyDiskBackupWithContext(context.Background(), request)
}

// ApplyDiskBackup
// 本接口（ApplyDiskBackup）用于回滚备份点到原云硬盘。
//
// 
//
// * 仅支持回滚到原云硬盘上。对于数据盘备份点，如果您需要复制备份点数据到其它云硬盘上，请先使用 CreateSnapshot 将备份点转换为快照，然后使用 CreateDisks 接口创建新的弹性云硬盘，将快照数据复制到新购云硬盘上。
//
// * 用于回滚的备份点必须处于NORMAL状态。备份点状态可以通过DescribeDiskBackups接口查询，见输出参数中BackupState字段解释。
//
// * 如果是弹性云硬盘，则云硬盘必须处于未挂载状态，云硬盘挂载状态可以通过DescribeDisks接口查询，见Attached字段解释；如果是随实例一起购买的非弹性云硬盘，则实例必须处于关机状态，实例状态可以通过DescribeInstancesStatus接口查询。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) ApplyDiskBackupWithContext(ctx context.Context, request *ApplyDiskBackupRequest) (response *ApplyDiskBackupResponse, err error) {
    if request == nil {
        request = NewApplyDiskBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyDiskBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyDiskBackupResponse()
    err = c.Send(request, response)
    return
}

func NewApplySnapshotRequest() (request *ApplySnapshotRequest) {
    request = &ApplySnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ApplySnapshot")
    
    
    return
}

func NewApplySnapshotResponse() (response *ApplySnapshotResponse) {
    response = &ApplySnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplySnapshot
// 本接口（ApplySnapshot）用于回滚快照到原云硬盘。
//
// 
//
// * 仅支持回滚到原云硬盘上。对于数据盘快照，如果您需要复制快照数据到其它云硬盘上，请使用[CreateDisks](/document/product/362/16312)接口创建新的弹性云盘，将快照数据复制到新购云盘上。 
//
// * 用于回滚的快照必须处于NORMAL状态。快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
//
// * 如果是弹性云盘，则云盘必须处于未挂载状态，云硬盘挂载状态可以通过[DescribeDisks](/document/product/362/16315)接口查询，见Attached字段解释；如果是随实例一起购买的非弹性云盘，则实例必须处于关机状态，实例状态可以通过[DescribeInstancesStatus](/document/product/213/15738)接口查询。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"
//  INVALIDPARAMETER_SHOULDCONVERTSNAPSHOTTOIMAGE = "InvalidParameter.ShouldConvertSnapshotToImage"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ApplySnapshot(request *ApplySnapshotRequest) (response *ApplySnapshotResponse, err error) {
    return c.ApplySnapshotWithContext(context.Background(), request)
}

// ApplySnapshot
// 本接口（ApplySnapshot）用于回滚快照到原云硬盘。
//
// 
//
// * 仅支持回滚到原云硬盘上。对于数据盘快照，如果您需要复制快照数据到其它云硬盘上，请使用[CreateDisks](/document/product/362/16312)接口创建新的弹性云盘，将快照数据复制到新购云盘上。 
//
// * 用于回滚的快照必须处于NORMAL状态。快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
//
// * 如果是弹性云盘，则云盘必须处于未挂载状态，云硬盘挂载状态可以通过[DescribeDisks](/document/product/362/16315)接口查询，见Attached字段解释；如果是随实例一起购买的非弹性云盘，则实例必须处于关机状态，实例状态可以通过[DescribeInstancesStatus](/document/product/213/15738)接口查询。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"
//  INVALIDPARAMETER_SHOULDCONVERTSNAPSHOTTOIMAGE = "InvalidParameter.ShouldConvertSnapshotToImage"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ApplySnapshotWithContext(ctx context.Context, request *ApplySnapshotRequest) (response *ApplySnapshotResponse, err error) {
    if request == nil {
        request = NewApplySnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplySnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplySnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
    request = &AttachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "AttachDisks")
    
    
    return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
    response = &AttachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachDisks
// 本接口（AttachDisks）用于挂载云硬盘。
//
//  
//
// * 支持批量操作，将多块云盘挂载到同一云主机。如果多个云盘中存在不允许挂载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当挂载云盘的请求成功返回时，表示后台已发起挂载云盘的操作，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHING”变为“ATTACHED”，则为挂载成功。
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  RESOURCEUNAVAILABLE_ZONENOTMATCH = "ResourceUnavailable.ZoneNotMatch"
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    return c.AttachDisksWithContext(context.Background(), request)
}

// AttachDisks
// 本接口（AttachDisks）用于挂载云硬盘。
//
//  
//
// * 支持批量操作，将多块云盘挂载到同一云主机。如果多个云盘中存在不允许挂载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当挂载云盘的请求成功返回时，表示后台已发起挂载云盘的操作，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHING”变为“ATTACHED”，则为挂载成功。
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  RESOURCEUNAVAILABLE_ZONENOTMATCH = "ResourceUnavailable.ZoneNotMatch"
func (c *Client) AttachDisksWithContext(ctx context.Context, request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    if request == nil {
        request = NewAttachDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachDisksResponse()
    err = c.Send(request, response)
    return
}

func NewBindAutoSnapshotPolicyRequest() (request *BindAutoSnapshotPolicyRequest) {
    request = &BindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "BindAutoSnapshotPolicy")
    
    
    return
}

func NewBindAutoSnapshotPolicyResponse() (response *BindAutoSnapshotPolicyResponse) {
    response = &BindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindAutoSnapshotPolicy
// 本接口（BindAutoSnapshotPolicy）用于绑定云硬盘到指定的定期快照策略。
//
// 
//
// * 每个地域下的定期快照策略配额限制请参考文档[定期快照](/document/product/362/8191)。
//
// * 当已绑定定期快照策略的云硬盘处于未使用状态（即弹性云盘未挂载或非弹性云盘的主机处于关机状态）将不会创建定期快照。
//
// 可能返回的错误码:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDDISK_ALREADYBOUND = "InvalidDisk.AlreadyBound"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BINDDISKLIMITEXCEEDED = "InvalidParameterValue.BindDiskLimitExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) BindAutoSnapshotPolicy(request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
    return c.BindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// BindAutoSnapshotPolicy
// 本接口（BindAutoSnapshotPolicy）用于绑定云硬盘到指定的定期快照策略。
//
// 
//
// * 每个地域下的定期快照策略配额限制请参考文档[定期快照](/document/product/362/8191)。
//
// * 当已绑定定期快照策略的云硬盘处于未使用状态（即弹性云盘未挂载或非弹性云盘的主机处于关机状态）将不会创建定期快照。
//
// 可能返回的错误码:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDDISK_ALREADYBOUND = "InvalidDisk.AlreadyBound"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BINDDISKLIMITEXCEEDED = "InvalidParameterValue.BindDiskLimitExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
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

func NewCopySnapshotCrossRegionsRequest() (request *CopySnapshotCrossRegionsRequest) {
    request = &CopySnapshotCrossRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "CopySnapshotCrossRegions")
    
    
    return
}

func NewCopySnapshotCrossRegionsResponse() (response *CopySnapshotCrossRegionsResponse) {
    response = &CopySnapshotCrossRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopySnapshotCrossRegions
// 本接口（CopySnapshotCrossRegions）用于快照跨地域复制。
//
// 
//
// * 本接口为异步接口，当跨地域复制的请求下发成功后会返回一个新的快照ID，此时快照未立即复制到目标地域，可请求目标地域的[DescribeSnapshots](/document/product/362/15647)接口查询新快照的状态，判断是否复制完成。如果快照的状态为“NORMAL”，表示快照复制完成。
//
// * 本接口实现的快照跨地域复制操作将产生跨地域流量，预计2022年第三季度会针对此功能进行商业化计费；请留意后续站内信公告，避免产生预期外扣费。
//
// 可能返回的错误码:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_COPYSNAPSHOTCONFLICT = "ResourceInUse.CopySnapshotConflict"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTSUPPORTCOPY = "UnsupportedOperation.SnapshotNotSupportCopy"
func (c *Client) CopySnapshotCrossRegions(request *CopySnapshotCrossRegionsRequest) (response *CopySnapshotCrossRegionsResponse, err error) {
    return c.CopySnapshotCrossRegionsWithContext(context.Background(), request)
}

// CopySnapshotCrossRegions
// 本接口（CopySnapshotCrossRegions）用于快照跨地域复制。
//
// 
//
// * 本接口为异步接口，当跨地域复制的请求下发成功后会返回一个新的快照ID，此时快照未立即复制到目标地域，可请求目标地域的[DescribeSnapshots](/document/product/362/15647)接口查询新快照的状态，判断是否复制完成。如果快照的状态为“NORMAL”，表示快照复制完成。
//
// * 本接口实现的快照跨地域复制操作将产生跨地域流量，预计2022年第三季度会针对此功能进行商业化计费；请留意后续站内信公告，避免产生预期外扣费。
//
// 可能返回的错误码:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_COPYSNAPSHOTCONFLICT = "ResourceInUse.CopySnapshotConflict"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTSUPPORTCOPY = "UnsupportedOperation.SnapshotNotSupportCopy"
func (c *Client) CopySnapshotCrossRegionsWithContext(ctx context.Context, request *CopySnapshotCrossRegionsRequest) (response *CopySnapshotCrossRegionsResponse, err error) {
    if request == nil {
        request = NewCopySnapshotCrossRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopySnapshotCrossRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopySnapshotCrossRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
    request = &CreateAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "CreateAutoSnapshotPolicy")
    
    
    return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
    response = &CreateAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoSnapshotPolicy
// 本接口（CreateAutoSnapshotPolicy）用于创建定期快照策略。
//
// 
//
// * 每个地域可创建的定期快照策略数量限制请参考文档[定期快照](/document/product/362/8191)。
//
// * 每个地域可创建的快照有数量和容量的限制，具体请见腾讯云控制台快照页面提示，如果快照超配额，定期快照创建会失败。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_AUTOSNAPSHOTPOLICYOUTOFQUOTA = "LimitExceeded.AutoSnapshotPolicyOutOfQuota"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
    return c.CreateAutoSnapshotPolicyWithContext(context.Background(), request)
}

// CreateAutoSnapshotPolicy
// 本接口（CreateAutoSnapshotPolicy）用于创建定期快照策略。
//
// 
//
// * 每个地域可创建的定期快照策略数量限制请参考文档[定期快照](/document/product/362/8191)。
//
// * 每个地域可创建的快照有数量和容量的限制，具体请见腾讯云控制台快照页面提示，如果快照超配额，定期快照创建会失败。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_AUTOSNAPSHOTPOLICYOUTOFQUOTA = "LimitExceeded.AutoSnapshotPolicyOutOfQuota"
//  MISSINGPARAMETER = "MissingParameter"
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

func NewCreateDiskBackupRequest() (request *CreateDiskBackupRequest) {
    request = &CreateDiskBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "CreateDiskBackup")
    
    
    return
}

func NewCreateDiskBackupResponse() (response *CreateDiskBackupResponse) {
    response = &CreateDiskBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDiskBackup
// 为云硬盘创建一个备份点。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
func (c *Client) CreateDiskBackup(request *CreateDiskBackupRequest) (response *CreateDiskBackupResponse, err error) {
    return c.CreateDiskBackupWithContext(context.Background(), request)
}

// CreateDiskBackup
// 为云硬盘创建一个备份点。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
func (c *Client) CreateDiskBackupWithContext(ctx context.Context, request *CreateDiskBackupRequest) (response *CreateDiskBackupResponse, err error) {
    if request == nil {
        request = NewCreateDiskBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDiskBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDiskBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
    request = &CreateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "CreateDisks")
    
    
    return
}

func NewCreateDisksResponse() (response *CreateDisksResponse) {
    response = &CreateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDisks
// 本接口（CreateDisks）用于创建云硬盘。
//
// 
//
// * 预付费云盘的购买会预先扣除本次云盘购买所需金额，在调用本接口前请确保账户余额充足。
//
// * 本接口支持传入数据盘快照来创建云盘，实现将快照数据复制到新购云盘上。
//
// * 本接口为异步接口，当创建请求下发成功后会返回一个新建的云盘ID列表，此时云盘的创建并未立即完成。可以通过调用[DescribeDisks](/document/product/362/16315)接口根据DiskId查询对应云盘，如果能查到云盘，且状态为'UNATTACHED'或'ATTACHED'，则表示创建成功。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNAUTHORIZEDOPERATION_ROLENOTEXISTS = "UnauthorizedOperation.RoleNotExists"
func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
    return c.CreateDisksWithContext(context.Background(), request)
}

// CreateDisks
// 本接口（CreateDisks）用于创建云硬盘。
//
// 
//
// * 预付费云盘的购买会预先扣除本次云盘购买所需金额，在调用本接口前请确保账户余额充足。
//
// * 本接口支持传入数据盘快照来创建云盘，实现将快照数据复制到新购云盘上。
//
// * 本接口为异步接口，当创建请求下发成功后会返回一个新建的云盘ID列表，此时云盘的创建并未立即完成。可以通过调用[DescribeDisks](/document/product/362/16315)接口根据DiskId查询对应云盘，如果能查到云盘，且状态为'UNATTACHED'或'ATTACHED'，则表示创建成功。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNAUTHORIZEDOPERATION_ROLENOTEXISTS = "UnauthorizedOperation.RoleNotExists"
func (c *Client) CreateDisksWithContext(ctx context.Context, request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
    if request == nil {
        request = NewCreateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
    request = &CreateSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "CreateSnapshot")
    
    
    return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
    response = &CreateSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSnapshot
// 本接口（CreateSnapshot）用于对指定云盘创建快照。
//
// 
//
// * 只有具有快照能力的云硬盘才能创建快照。云硬盘是否具有快照能力可由[DescribeDisks](/document/product/362/16315)接口查询，见SnapshotAbility字段。
//
// * 可创建快照数量限制见[产品使用限制](https://cloud.tencent.com/doc/product/362/5145)。
//
// * 当前支持将备份点转化为普通快照，转化之后可能会收取快照使用费用，备份点不保留，其占用的备份点配额也将被释放。
//
// 可能返回的错误码:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTSNAPSHOT = "InvalidDisk.NotSupportSnapshot"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_DISKSNAPSHOTCHAINTOOLARGE = "ResourceUnavailable.DiskSnapshotChainTooLarge"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  RESOURCEUNAVAILABLE_TOOMANYCREATINGSNAPSHOT = "ResourceUnavailable.TooManyCreatingSnapshot"
//  UNSUPPORTEDOPERATION_DISKENCRYPT = "UnsupportedOperation.DiskEncrypt"
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    return c.CreateSnapshotWithContext(context.Background(), request)
}

// CreateSnapshot
// 本接口（CreateSnapshot）用于对指定云盘创建快照。
//
// 
//
// * 只有具有快照能力的云硬盘才能创建快照。云硬盘是否具有快照能力可由[DescribeDisks](/document/product/362/16315)接口查询，见SnapshotAbility字段。
//
// * 可创建快照数量限制见[产品使用限制](https://cloud.tencent.com/doc/product/362/5145)。
//
// * 当前支持将备份点转化为普通快照，转化之后可能会收取快照使用费用，备份点不保留，其占用的备份点配额也将被释放。
//
// 可能返回的错误码:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTSNAPSHOT = "InvalidDisk.NotSupportSnapshot"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_DISKSNAPSHOTCHAINTOOLARGE = "ResourceUnavailable.DiskSnapshotChainTooLarge"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  RESOURCEUNAVAILABLE_TOOMANYCREATINGSNAPSHOT = "ResourceUnavailable.TooManyCreatingSnapshot"
//  UNSUPPORTEDOPERATION_DISKENCRYPT = "UnsupportedOperation.DiskEncrypt"
func (c *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoSnapshotPoliciesRequest() (request *DeleteAutoSnapshotPoliciesRequest) {
    request = &DeleteAutoSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DeleteAutoSnapshotPolicies")
    
    
    return
}

func NewDeleteAutoSnapshotPoliciesResponse() (response *DeleteAutoSnapshotPoliciesResponse) {
    response = &DeleteAutoSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAutoSnapshotPolicies
// 本接口（DeleteAutoSnapshotPolicies）用于删除定期快照策略。
//
// 
//
// *  支持批量操作。如果多个定期快照策略存在无法删除的，则操作不执行，以特定错误码返回。
//
// 可能返回的错误码:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAutoSnapshotPolicies(request *DeleteAutoSnapshotPoliciesRequest) (response *DeleteAutoSnapshotPoliciesResponse, err error) {
    return c.DeleteAutoSnapshotPoliciesWithContext(context.Background(), request)
}

// DeleteAutoSnapshotPolicies
// 本接口（DeleteAutoSnapshotPolicies）用于删除定期快照策略。
//
// 
//
// *  支持批量操作。如果多个定期快照策略存在无法删除的，则操作不执行，以特定错误码返回。
//
// 可能返回的错误码:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAutoSnapshotPoliciesWithContext(ctx context.Context, request *DeleteAutoSnapshotPoliciesRequest) (response *DeleteAutoSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteAutoSnapshotPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoSnapshotPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDiskBackupsRequest() (request *DeleteDiskBackupsRequest) {
    request = &DeleteDiskBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DeleteDiskBackups")
    
    
    return
}

func NewDeleteDiskBackupsResponse() (response *DeleteDiskBackupsResponse) {
    response = &DeleteDiskBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDiskBackups
// 批量删除指定的云硬盘备份点。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DeleteDiskBackups(request *DeleteDiskBackupsRequest) (response *DeleteDiskBackupsResponse, err error) {
    return c.DeleteDiskBackupsWithContext(context.Background(), request)
}

// DeleteDiskBackups
// 批量删除指定的云硬盘备份点。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
func (c *Client) DeleteDiskBackupsWithContext(ctx context.Context, request *DeleteDiskBackupsRequest) (response *DeleteDiskBackupsResponse, err error) {
    if request == nil {
        request = NewDeleteDiskBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDiskBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDiskBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DeleteSnapshots")
    
    
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSnapshots
// 本接口（DeleteSnapshots）用于删除快照。
//
// 
//
// * 快照必须处于NORMAL状态，快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
//
// * 支持批量操作。如果多个快照存在无法删除的快照，则操作不执行，以特定的错误码返回。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    return c.DeleteSnapshotsWithContext(context.Background(), request)
}

// DeleteSnapshots
// 本接口（DeleteSnapshots）用于删除快照。
//
// 
//
// * 快照必须处于NORMAL状态，快照状态可以通过[DescribeSnapshots](/document/product/362/15647)接口查询，见输出参数中SnapshotState字段解释。
//
// * 支持批量操作。如果多个快照存在无法删除的快照，则操作不执行，以特定的错误码返回。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshotsWithContext(ctx context.Context, request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
    request = &DescribeAutoSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeAutoSnapshotPolicies")
    
    
    return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
    response = &DescribeAutoSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoSnapshotPolicies
// 本接口（DescribeAutoSnapshotPolicies）用于查询定期快照策略。
//
// 
//
// * 可以根据定期快照策略ID、名称或者状态等信息来查询定期快照策略的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的定期快照策略表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
    return c.DescribeAutoSnapshotPoliciesWithContext(context.Background(), request)
}

// DescribeAutoSnapshotPolicies
// 本接口（DescribeAutoSnapshotPolicies）用于查询定期快照策略。
//
// 
//
// * 可以根据定期快照策略ID、名称或者状态等信息来查询定期快照策略的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的定期快照策略表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeDiskAssociatedAutoSnapshotPolicyRequest() (request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) {
    request = &DescribeDiskAssociatedAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskAssociatedAutoSnapshotPolicy")
    
    
    return
}

func NewDescribeDiskAssociatedAutoSnapshotPolicyResponse() (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse) {
    response = &DescribeDiskAssociatedAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskAssociatedAutoSnapshotPolicy
// 本接口（DescribeDiskAssociatedAutoSnapshotPolicy）用于查询云盘绑定的定期快照策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDiskAssociatedAutoSnapshotPolicy(request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse, err error) {
    return c.DescribeDiskAssociatedAutoSnapshotPolicyWithContext(context.Background(), request)
}

// DescribeDiskAssociatedAutoSnapshotPolicy
// 本接口（DescribeDiskAssociatedAutoSnapshotPolicy）用于查询云盘绑定的定期快照策略。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDiskAssociatedAutoSnapshotPolicyWithContext(ctx context.Context, request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeDiskAssociatedAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskAssociatedAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskAssociatedAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskBackupsRequest() (request *DescribeDiskBackupsRequest) {
    request = &DescribeDiskBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskBackups")
    
    
    return
}

func NewDescribeDiskBackupsResponse() (response *DescribeDiskBackupsResponse) {
    response = &DescribeDiskBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskBackups
// 本接口（DescribeDiskBackups）用于查询备份点的详细信息。
//
// 
//
// 根据备份点ID、创建备份点的云硬盘ID、创建备份点的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器Filter。
//
// 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的备份点列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDiskBackups(request *DescribeDiskBackupsRequest) (response *DescribeDiskBackupsResponse, err error) {
    return c.DescribeDiskBackupsWithContext(context.Background(), request)
}

// DescribeDiskBackups
// 本接口（DescribeDiskBackups）用于查询备份点的详细信息。
//
// 
//
// 根据备份点ID、创建备份点的云硬盘ID、创建备份点的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器Filter。
//
// 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的备份点列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDiskBackupsWithContext(ctx context.Context, request *DescribeDiskBackupsRequest) (response *DescribeDiskBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeDiskBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskConfigQuotaRequest() (request *DescribeDiskConfigQuotaRequest) {
    request = &DescribeDiskConfigQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskConfigQuota")
    
    
    return
}

func NewDescribeDiskConfigQuotaResponse() (response *DescribeDiskConfigQuotaResponse) {
    response = &DescribeDiskConfigQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskConfigQuota
// 本接口（DescribeDiskConfigQuota）用于查询云硬盘配额。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDiskConfigQuota(request *DescribeDiskConfigQuotaRequest) (response *DescribeDiskConfigQuotaResponse, err error) {
    return c.DescribeDiskConfigQuotaWithContext(context.Background(), request)
}

// DescribeDiskConfigQuota
// 本接口（DescribeDiskConfigQuota）用于查询云硬盘配额。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDiskConfigQuotaWithContext(ctx context.Context, request *DescribeDiskConfigQuotaRequest) (response *DescribeDiskConfigQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeDiskConfigQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskConfigQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskConfigQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskOperationLogsRequest() (request *DescribeDiskOperationLogsRequest) {
    request = &DescribeDiskOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskOperationLogs")
    
    
    return
}

func NewDescribeDiskOperationLogsResponse() (response *DescribeDiskOperationLogsResponse) {
    response = &DescribeDiskOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskOperationLogs
// 接口已废弃，切换至云审计接口。见https://tapd.woa.com/pro/prong/stories/view/1010114221880719007
//
// 
//
// 查询云盘操作日志功能已迁移至LookUpEvents接口（https://cloud.tencent.com/document/product/629/12359），本接口（DescribeDiskOperationLogs）即将下线，后续不再提供调用，请知悉。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskOperationLogs(request *DescribeDiskOperationLogsRequest) (response *DescribeDiskOperationLogsResponse, err error) {
    return c.DescribeDiskOperationLogsWithContext(context.Background(), request)
}

// DescribeDiskOperationLogs
// 接口已废弃，切换至云审计接口。见https://tapd.woa.com/pro/prong/stories/view/1010114221880719007
//
// 
//
// 查询云盘操作日志功能已迁移至LookUpEvents接口（https://cloud.tencent.com/document/product/629/12359），本接口（DescribeDiskOperationLogs）即将下线，后续不再提供调用，请知悉。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskOperationLogsWithContext(ctx context.Context, request *DescribeDiskOperationLogsRequest) (response *DescribeDiskOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDiskOperationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskStoragePoolRequest() (request *DescribeDiskStoragePoolRequest) {
    request = &DescribeDiskStoragePoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskStoragePool")
    
    
    return
}

func NewDescribeDiskStoragePoolResponse() (response *DescribeDiskStoragePoolResponse) {
    response = &DescribeDiskStoragePoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskStoragePool
// 本接口（DescribeDiskStoragePool）查询用户的云硬盘独享集群列表。
//
// 
//
// * 可以根据独享集群ID(CdcId)、可用区(zone)等信息来查询和过滤云硬盘独享集群详细信息，不同的过滤条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘独享集群列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskStoragePool(request *DescribeDiskStoragePoolRequest) (response *DescribeDiskStoragePoolResponse, err error) {
    return c.DescribeDiskStoragePoolWithContext(context.Background(), request)
}

// DescribeDiskStoragePool
// 本接口（DescribeDiskStoragePool）查询用户的云硬盘独享集群列表。
//
// 
//
// * 可以根据独享集群ID(CdcId)、可用区(zone)等信息来查询和过滤云硬盘独享集群详细信息，不同的过滤条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘独享集群列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskStoragePoolWithContext(ctx context.Context, request *DescribeDiskStoragePoolRequest) (response *DescribeDiskStoragePoolResponse, err error) {
    if request == nil {
        request = NewDescribeDiskStoragePoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskStoragePool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskStoragePoolResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
    request = &DescribeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisks")
    
    
    return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
    response = &DescribeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisks
// 本接口（DescribeDisks）用于查询云硬盘列表。
//
// 
//
// * 可以根据云硬盘ID、云硬盘类型或者云硬盘状态等信息来查询云硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    return c.DescribeDisksWithContext(context.Background(), request)
}

// DescribeDisks
// 本接口（DescribeDisks）用于查询云硬盘列表。
//
// 
//
// * 可以根据云硬盘ID、云硬盘类型或者云硬盘状态等信息来查询云硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDisksWithContext(ctx context.Context, request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    if request == nil {
        request = NewDescribeDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDiskNumRequest() (request *DescribeInstancesDiskNumRequest) {
    request = &DescribeInstancesDiskNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeInstancesDiskNum")
    
    
    return
}

func NewDescribeInstancesDiskNumResponse() (response *DescribeInstancesDiskNumResponse) {
    response = &DescribeInstancesDiskNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDiskNum
// 本接口（DescribeInstancesDiskNum）用于查询实例已挂载云硬盘数量。
//
// 
//
// * 支持批量操作，当传入多个云服务器实例ID，返回结果会分别列出每个云服务器挂载的云硬盘数量。
//
// 可能返回的错误码:
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeInstancesDiskNum(request *DescribeInstancesDiskNumRequest) (response *DescribeInstancesDiskNumResponse, err error) {
    return c.DescribeInstancesDiskNumWithContext(context.Background(), request)
}

// DescribeInstancesDiskNum
// 本接口（DescribeInstancesDiskNum）用于查询实例已挂载云硬盘数量。
//
// 
//
// * 支持批量操作，当传入多个云服务器实例ID，返回结果会分别列出每个云服务器挂载的云硬盘数量。
//
// 可能返回的错误码:
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeInstancesDiskNumWithContext(ctx context.Context, request *DescribeInstancesDiskNumRequest) (response *DescribeInstancesDiskNumResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDiskNumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDiskNum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDiskNumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotOperationLogsRequest() (request *DescribeSnapshotOperationLogsRequest) {
    request = &DescribeSnapshotOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotOperationLogs")
    
    
    return
}

func NewDescribeSnapshotOperationLogsResponse() (response *DescribeSnapshotOperationLogsResponse) {
    response = &DescribeSnapshotOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotOperationLogs
// 接口已废弃，切换至云审计接口。见https://tapd.woa.com/pro/prong/stories/view/1010114221880719007
//
// 
//
// 查询快照操作日志功能已迁移至LookUpEvents接口（https://cloud.tencent.com/document/product/629/12359），本接口（DescribeSnapshotOperationLogs）即将下线，后续不再提供调用，请知悉。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotOperationLogs(request *DescribeSnapshotOperationLogsRequest) (response *DescribeSnapshotOperationLogsResponse, err error) {
    return c.DescribeSnapshotOperationLogsWithContext(context.Background(), request)
}

// DescribeSnapshotOperationLogs
// 接口已废弃，切换至云审计接口。见https://tapd.woa.com/pro/prong/stories/view/1010114221880719007
//
// 
//
// 查询快照操作日志功能已迁移至LookUpEvents接口（https://cloud.tencent.com/document/product/629/12359），本接口（DescribeSnapshotOperationLogs）即将下线，后续不再提供调用，请知悉。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
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

func NewDescribeSnapshotSharePermissionRequest() (request *DescribeSnapshotSharePermissionRequest) {
    request = &DescribeSnapshotSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotSharePermission")
    
    
    return
}

func NewDescribeSnapshotSharePermissionResponse() (response *DescribeSnapshotSharePermissionResponse) {
    response = &DescribeSnapshotSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotSharePermission
// 本接口（DescribeSnapshotSharePermission）用于查询快照的分享信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotSharePermission(request *DescribeSnapshotSharePermissionRequest) (response *DescribeSnapshotSharePermissionResponse, err error) {
    return c.DescribeSnapshotSharePermissionWithContext(context.Background(), request)
}

// DescribeSnapshotSharePermission
// 本接口（DescribeSnapshotSharePermission）用于查询快照的分享信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotSharePermissionWithContext(ctx context.Context, request *DescribeSnapshotSharePermissionRequest) (response *DescribeSnapshotSharePermissionResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotSharePermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotSharePermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshots")
    
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshots
// 本接口（DescribeSnapshots）用于查询快照的详细信息。
//
// 
//
// * 根据快照ID、创建快照的云硬盘ID、创建快照的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// *  如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的快照列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

// DescribeSnapshots
// 本接口（DescribeSnapshots）用于查询快照的详细信息。
//
// 
//
// * 根据快照ID、创建快照的云硬盘ID、创建快照的云硬盘类型等对结果进行过滤，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
//
// *  如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的快照列表。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
    request = &DetachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "DetachDisks")
    
    
    return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
    response = &DetachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachDisks
// 本接口（DetachDisks）用于卸载云硬盘。
//
// 
//
// * 支持批量操作，卸载挂载在同一主机上的多块云盘。如果多块云盘中存在不允许卸载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当请求成功返回时，云盘并未立即从主机卸载，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHED”变为“UNATTACHED”，则为卸载成功。
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DETACHPOD = "UnsupportedOperation.DetachPod"
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    return c.DetachDisksWithContext(context.Background(), request)
}

// DetachDisks
// 本接口（DetachDisks）用于卸载云硬盘。
//
// 
//
// * 支持批量操作，卸载挂载在同一主机上的多块云盘。如果多块云盘中存在不允许卸载的云盘，则操作不执行，返回特定的错误码。
//
// * 本接口为异步接口，当请求成功返回时，云盘并未立即从主机卸载，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态由“ATTACHED”变为“UNATTACHED”，则为卸载成功。
//
// 可能返回的错误码:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DETACHPOD = "UnsupportedOperation.DetachPod"
func (c *Client) DetachDisksWithContext(ctx context.Context, request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    if request == nil {
        request = NewDetachDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachDisksResponse()
    err = c.Send(request, response)
    return
}

func NewGetSnapOverviewRequest() (request *GetSnapOverviewRequest) {
    request = &GetSnapOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "GetSnapOverview")
    
    
    return
}

func NewGetSnapOverviewResponse() (response *GetSnapOverviewResponse) {
    response = &GetSnapOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSnapOverview
// 获取快照概览信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetSnapOverview(request *GetSnapOverviewRequest) (response *GetSnapOverviewResponse, err error) {
    return c.GetSnapOverviewWithContext(context.Background(), request)
}

// GetSnapOverview
// 获取快照概览信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetSnapOverviewWithContext(ctx context.Context, request *GetSnapOverviewRequest) (response *GetSnapOverviewResponse, err error) {
    if request == nil {
        request = NewGetSnapOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSnapOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSnapOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewInitializeDisksRequest() (request *InitializeDisksRequest) {
    request = &InitializeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "InitializeDisks")
    
    
    return
}

func NewInitializeDisksResponse() (response *InitializeDisksResponse) {
    response = &InitializeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitializeDisks
// 重新初始化云硬盘至云硬盘初始创建时的状态。使用云硬盘的重新初始化功能时需要注意以下4点：
//
// 1. 如果云硬盘是由快照创建的，则重新初始化会通过此快照重新回滚此云硬盘，即将云硬盘恢复为与快照一致的状态；
//
// 2. 如果云硬盘不是通过快照创建的，则重新初始化会清空此云硬盘的数据；请在重新初始化云硬盘前检查并备份必要的数据；
//
// 3. 当前仅未挂载的、非共享属性的数据盘云硬盘支持重新初始化；
//
// 4. 当创建此云硬盘的原始快照被删除时，不再支持重新初始化此云硬盘。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEBUSY = "ResourceBusy"
func (c *Client) InitializeDisks(request *InitializeDisksRequest) (response *InitializeDisksResponse, err error) {
    return c.InitializeDisksWithContext(context.Background(), request)
}

// InitializeDisks
// 重新初始化云硬盘至云硬盘初始创建时的状态。使用云硬盘的重新初始化功能时需要注意以下4点：
//
// 1. 如果云硬盘是由快照创建的，则重新初始化会通过此快照重新回滚此云硬盘，即将云硬盘恢复为与快照一致的状态；
//
// 2. 如果云硬盘不是通过快照创建的，则重新初始化会清空此云硬盘的数据；请在重新初始化云硬盘前检查并备份必要的数据；
//
// 3. 当前仅未挂载的、非共享属性的数据盘云硬盘支持重新初始化；
//
// 4. 当创建此云硬盘的原始快照被删除时，不再支持重新初始化此云硬盘。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEBUSY = "ResourceBusy"
func (c *Client) InitializeDisksWithContext(ctx context.Context, request *InitializeDisksRequest) (response *InitializeDisksResponse, err error) {
    if request == nil {
        request = NewInitializeDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitializeDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitializeDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyDiskBackupQuotaRequest() (request *InquirePriceModifyDiskBackupQuotaRequest) {
    request = &InquirePriceModifyDiskBackupQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "InquirePriceModifyDiskBackupQuota")
    
    
    return
}

func NewInquirePriceModifyDiskBackupQuotaResponse() (response *InquirePriceModifyDiskBackupQuotaResponse) {
    response = &InquirePriceModifyDiskBackupQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceModifyDiskBackupQuota
// 本接口（InquirePricePriceModifyDiskBackupQuota）用于修改云硬盘备份点配额询价。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) InquirePriceModifyDiskBackupQuota(request *InquirePriceModifyDiskBackupQuotaRequest) (response *InquirePriceModifyDiskBackupQuotaResponse, err error) {
    return c.InquirePriceModifyDiskBackupQuotaWithContext(context.Background(), request)
}

// InquirePriceModifyDiskBackupQuota
// 本接口（InquirePricePriceModifyDiskBackupQuota）用于修改云硬盘备份点配额询价。
//
// 可能返回的错误码:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) InquirePriceModifyDiskBackupQuotaWithContext(ctx context.Context, request *InquirePriceModifyDiskBackupQuotaRequest) (response *InquirePriceModifyDiskBackupQuotaResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyDiskBackupQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModifyDiskBackupQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyDiskBackupQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyDiskExtraPerformanceRequest() (request *InquirePriceModifyDiskExtraPerformanceRequest) {
    request = &InquirePriceModifyDiskExtraPerformanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "InquirePriceModifyDiskExtraPerformance")
    
    
    return
}

func NewInquirePriceModifyDiskExtraPerformanceResponse() (response *InquirePriceModifyDiskExtraPerformanceResponse) {
    response = &InquirePriceModifyDiskExtraPerformanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceModifyDiskExtraPerformance
// 本接口（InquirePriceModifyDiskExtraPerformance）用于调整云硬盘额外性能询价。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) InquirePriceModifyDiskExtraPerformance(request *InquirePriceModifyDiskExtraPerformanceRequest) (response *InquirePriceModifyDiskExtraPerformanceResponse, err error) {
    return c.InquirePriceModifyDiskExtraPerformanceWithContext(context.Background(), request)
}

// InquirePriceModifyDiskExtraPerformance
// 本接口（InquirePriceModifyDiskExtraPerformance）用于调整云硬盘额外性能询价。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) InquirePriceModifyDiskExtraPerformanceWithContext(ctx context.Context, request *InquirePriceModifyDiskExtraPerformanceRequest) (response *InquirePriceModifyDiskExtraPerformanceResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyDiskExtraPerformanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModifyDiskExtraPerformance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyDiskExtraPerformanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateDisksRequest() (request *InquiryPriceCreateDisksRequest) {
    request = &InquiryPriceCreateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceCreateDisks")
    
    
    return
}

func NewInquiryPriceCreateDisksResponse() (response *InquiryPriceCreateDisksResponse) {
    response = &InquiryPriceCreateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateDisks
// 本接口（InquiryPriceCreateDisks）用于创建云硬盘询价。
//
// 
//
// * 支持查询创建多块云硬盘的价格，此时返回结果为总价格。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquiryPriceCreateDisks(request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
    return c.InquiryPriceCreateDisksWithContext(context.Background(), request)
}

// InquiryPriceCreateDisks
// 本接口（InquiryPriceCreateDisks）用于创建云硬盘询价。
//
// 
//
// * 支持查询创建多块云硬盘的价格，此时返回结果为总价格。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquiryPriceCreateDisksWithContext(ctx context.Context, request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewDisksRequest() (request *InquiryPriceRenewDisksRequest) {
    request = &InquiryPriceRenewDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceRenewDisks")
    
    
    return
}

func NewInquiryPriceRenewDisksResponse() (response *InquiryPriceRenewDisksResponse) {
    response = &InquiryPriceRenewDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceRenewDisks
// 本接口（InquiryPriceRenewDisks）用于续费云硬盘询价。
//
// 
//
// * 只支持查询预付费模式的弹性云盘续费价格。
//
// * 支持与挂载实例一起续费的场景，需要在[DiskChargePrepaid](/document/product/362/15669#DiskChargePrepaid)参数中指定CurInstanceDeadline，此时会按对齐到实例续费后的到期时间来续费询价。
//
// * 支持为多块云盘指定不同的续费时长，此时返回的价格为多块云盘续费的总价格。
//
// 可能返回的错误码:
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) InquiryPriceRenewDisks(request *InquiryPriceRenewDisksRequest) (response *InquiryPriceRenewDisksResponse, err error) {
    return c.InquiryPriceRenewDisksWithContext(context.Background(), request)
}

// InquiryPriceRenewDisks
// 本接口（InquiryPriceRenewDisks）用于续费云硬盘询价。
//
// 
//
// * 只支持查询预付费模式的弹性云盘续费价格。
//
// * 支持与挂载实例一起续费的场景，需要在[DiskChargePrepaid](/document/product/362/15669#DiskChargePrepaid)参数中指定CurInstanceDeadline，此时会按对齐到实例续费后的到期时间来续费询价。
//
// * 支持为多块云盘指定不同的续费时长，此时返回的价格为多块云盘续费的总价格。
//
// 可能返回的错误码:
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) InquiryPriceRenewDisksWithContext(ctx context.Context, request *InquiryPriceRenewDisksRequest) (response *InquiryPriceRenewDisksResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResizeDiskRequest() (request *InquiryPriceResizeDiskRequest) {
    request = &InquiryPriceResizeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceResizeDisk")
    
    
    return
}

func NewInquiryPriceResizeDiskResponse() (response *InquiryPriceResizeDiskResponse) {
    response = &InquiryPriceResizeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceResizeDisk
// 本接口（InquiryPriceResizeDisk）用于扩容云硬盘询价。
//
// 可能返回的错误码:
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
func (c *Client) InquiryPriceResizeDisk(request *InquiryPriceResizeDiskRequest) (response *InquiryPriceResizeDiskResponse, err error) {
    return c.InquiryPriceResizeDiskWithContext(context.Background(), request)
}

// InquiryPriceResizeDisk
// 本接口（InquiryPriceResizeDisk）用于扩容云硬盘询价。
//
// 可能返回的错误码:
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
func (c *Client) InquiryPriceResizeDiskWithContext(ctx context.Context, request *InquiryPriceResizeDiskRequest) (response *InquiryPriceResizeDiskResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResizeDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceResizeDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceResizeDiskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoSnapshotPolicyAttributeRequest() (request *ModifyAutoSnapshotPolicyAttributeRequest) {
    request = &ModifyAutoSnapshotPolicyAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyAutoSnapshotPolicyAttribute")
    
    
    return
}

func NewModifyAutoSnapshotPolicyAttributeResponse() (response *ModifyAutoSnapshotPolicyAttributeResponse) {
    response = &ModifyAutoSnapshotPolicyAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoSnapshotPolicyAttribute
// 本接口（ModifyAutoSnapshotPolicyAttribute）用于修改定期快照策略属性。
//
// 
//
// * 可通过该接口修改定期快照策略的执行策略、名称、是否激活等属性。
//
// * 修改保留天数时必须保证不与是否永久保留属性冲突，否则整个操作失败，以特定的错误码返回。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyAutoSnapshotPolicyAttribute(request *ModifyAutoSnapshotPolicyAttributeRequest) (response *ModifyAutoSnapshotPolicyAttributeResponse, err error) {
    return c.ModifyAutoSnapshotPolicyAttributeWithContext(context.Background(), request)
}

// ModifyAutoSnapshotPolicyAttribute
// 本接口（ModifyAutoSnapshotPolicyAttribute）用于修改定期快照策略属性。
//
// 
//
// * 可通过该接口修改定期快照策略的执行策略、名称、是否激活等属性。
//
// * 修改保留天数时必须保证不与是否永久保留属性冲突，否则整个操作失败，以特定的错误码返回。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyAutoSnapshotPolicyAttributeWithContext(ctx context.Context, request *ModifyAutoSnapshotPolicyAttributeRequest) (response *ModifyAutoSnapshotPolicyAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAutoSnapshotPolicyAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoSnapshotPolicyAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoSnapshotPolicyAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiskAttributesRequest() (request *ModifyDiskAttributesRequest) {
    request = &ModifyDiskAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskAttributes")
    
    
    return
}

func NewModifyDiskAttributesResponse() (response *ModifyDiskAttributesResponse) {
    response = &ModifyDiskAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDiskAttributes
// * 只支持修改弹性云盘的项目ID。随云主机创建的云硬盘项目ID与云主机联动。可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中Portable字段解释。
//
// * “云硬盘名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行云盘管理操作的依据。
//
// * 支持批量操作，如果传入多个云盘ID，则所有云盘修改为同一属性。如果存在不允许操作的云盘，则操作不执行，以特定错误码返回。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) ModifyDiskAttributes(request *ModifyDiskAttributesRequest) (response *ModifyDiskAttributesResponse, err error) {
    return c.ModifyDiskAttributesWithContext(context.Background(), request)
}

// ModifyDiskAttributes
// * 只支持修改弹性云盘的项目ID。随云主机创建的云硬盘项目ID与云主机联动。可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中Portable字段解释。
//
// * “云硬盘名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行云盘管理操作的依据。
//
// * 支持批量操作，如果传入多个云盘ID，则所有云盘修改为同一属性。如果存在不允许操作的云盘，则操作不执行，以特定错误码返回。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) ModifyDiskAttributesWithContext(ctx context.Context, request *ModifyDiskAttributesRequest) (response *ModifyDiskAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDiskAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiskAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiskAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiskBackupQuotaRequest() (request *ModifyDiskBackupQuotaRequest) {
    request = &ModifyDiskBackupQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskBackupQuota")
    
    
    return
}

func NewModifyDiskBackupQuotaResponse() (response *ModifyDiskBackupQuotaResponse) {
    response = &ModifyDiskBackupQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDiskBackupQuota
// 此接口 (ModifyDiskBackupQuota) 用于修改云硬盘备份点配额。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
func (c *Client) ModifyDiskBackupQuota(request *ModifyDiskBackupQuotaRequest) (response *ModifyDiskBackupQuotaResponse, err error) {
    return c.ModifyDiskBackupQuotaWithContext(context.Background(), request)
}

// ModifyDiskBackupQuota
// 此接口 (ModifyDiskBackupQuota) 用于修改云硬盘备份点配额。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
func (c *Client) ModifyDiskBackupQuotaWithContext(ctx context.Context, request *ModifyDiskBackupQuotaRequest) (response *ModifyDiskBackupQuotaResponse, err error) {
    if request == nil {
        request = NewModifyDiskBackupQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiskBackupQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiskBackupQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiskExtraPerformanceRequest() (request *ModifyDiskExtraPerformanceRequest) {
    request = &ModifyDiskExtraPerformanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskExtraPerformance")
    
    
    return
}

func NewModifyDiskExtraPerformanceResponse() (response *ModifyDiskExtraPerformanceResponse) {
    response = &ModifyDiskExtraPerformanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDiskExtraPerformance
// 本接口（ModifyDiskExtraPerformance）用于调整云硬盘额外的性能。
//
// 
//
// * 目前仅支持极速型SSD云硬盘（CLOUD_TSSD）和高性能SSD云硬盘(CLOUD_HSSD)。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDiskExtraPerformance(request *ModifyDiskExtraPerformanceRequest) (response *ModifyDiskExtraPerformanceResponse, err error) {
    return c.ModifyDiskExtraPerformanceWithContext(context.Background(), request)
}

// ModifyDiskExtraPerformance
// 本接口（ModifyDiskExtraPerformance）用于调整云硬盘额外的性能。
//
// 
//
// * 目前仅支持极速型SSD云硬盘（CLOUD_TSSD）和高性能SSD云硬盘(CLOUD_HSSD)。
//
// 可能返回的错误码:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDiskExtraPerformanceWithContext(ctx context.Context, request *ModifyDiskExtraPerformanceRequest) (response *ModifyDiskExtraPerformanceResponse, err error) {
    if request == nil {
        request = NewModifyDiskExtraPerformanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiskExtraPerformance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiskExtraPerformanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisksChargeTypeRequest() (request *ModifyDisksChargeTypeRequest) {
    request = &ModifyDisksChargeTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyDisksChargeType")
    
    
    return
}

func NewModifyDisksChargeTypeResponse() (response *ModifyDisksChargeTypeResponse) {
    response = &ModifyDisksChargeTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDisksChargeType
// 接口请求域名： cbs.tencentcloudapi.com 。
//
// 
//
// 本接口 (ModifyDisksChargeType) 用于切换云盘的计费模式。
//
// 
//
// 非弹性云盘不支持此接口，请通过修改实例计费模式接口将实例连同非弹性云盘一起转换。
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
func (c *Client) ModifyDisksChargeType(request *ModifyDisksChargeTypeRequest) (response *ModifyDisksChargeTypeResponse, err error) {
    return c.ModifyDisksChargeTypeWithContext(context.Background(), request)
}

// ModifyDisksChargeType
// 接口请求域名： cbs.tencentcloudapi.com 。
//
// 
//
// 本接口 (ModifyDisksChargeType) 用于切换云盘的计费模式。
//
// 
//
// 非弹性云盘不支持此接口，请通过修改实例计费模式接口将实例连同非弹性云盘一起转换。
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
func (c *Client) ModifyDisksChargeTypeWithContext(ctx context.Context, request *ModifyDisksChargeTypeRequest) (response *ModifyDisksChargeTypeResponse, err error) {
    if request == nil {
        request = NewModifyDisksChargeTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDisksChargeType require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDisksChargeTypeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisksRenewFlagRequest() (request *ModifyDisksRenewFlagRequest) {
    request = &ModifyDisksRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyDisksRenewFlag")
    
    
    return
}

func NewModifyDisksRenewFlagResponse() (response *ModifyDisksRenewFlagResponse) {
    response = &ModifyDisksRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDisksRenewFlag
// 本接口（ModifyDisksRenewFlag）用于修改云硬盘续费标识，支持批量修改。
//
// 可能返回的错误码:
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyDisksRenewFlag(request *ModifyDisksRenewFlagRequest) (response *ModifyDisksRenewFlagResponse, err error) {
    return c.ModifyDisksRenewFlagWithContext(context.Background(), request)
}

// ModifyDisksRenewFlag
// 本接口（ModifyDisksRenewFlag）用于修改云硬盘续费标识，支持批量修改。
//
// 可能返回的错误码:
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
func (c *Client) ModifyDisksRenewFlagWithContext(ctx context.Context, request *ModifyDisksRenewFlagRequest) (response *ModifyDisksRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDisksRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDisksRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDisksRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
    request = &ModifySnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifySnapshotAttribute")
    
    
    return
}

func NewModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
    response = &ModifySnapshotAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySnapshotAttribute
// 本接口（ModifySnapshotAttribute）用于修改指定快照的属性。
//
// 
//
// * 当前仅支持修改快照名称及将非永久快照修改为永久快照。
//
// * “快照名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行快照管理操作的依据。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    return c.ModifySnapshotAttributeWithContext(context.Background(), request)
}

// ModifySnapshotAttribute
// 本接口（ModifySnapshotAttribute）用于修改指定快照的属性。
//
// 
//
// * 当前仅支持修改快照名称及将非永久快照修改为永久快照。
//
// * “快照名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行快照管理操作的依据。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) ModifySnapshotAttributeWithContext(ctx context.Context, request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    if request == nil {
        request = NewModifySnapshotAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotsSharePermissionRequest() (request *ModifySnapshotsSharePermissionRequest) {
    request = &ModifySnapshotsSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ModifySnapshotsSharePermission")
    
    
    return
}

func NewModifySnapshotsSharePermissionResponse() (response *ModifySnapshotsSharePermissionResponse) {
    response = &ModifySnapshotsSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySnapshotsSharePermission
// 本接口（ModifySnapshotsSharePermission）用于修改快照分享信息。
//
// 
//
// 分享快照后，被分享账户可以通过该快照创建云硬盘。
//
// * 每个快照最多可分享给50个账户。
//
// * 分享快照无法更改名称，描述，仅可用于创建云硬盘。
//
// * 只支持分享到对方账户相同地域。
//
// * 仅支持分享数据盘快照。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
func (c *Client) ModifySnapshotsSharePermission(request *ModifySnapshotsSharePermissionRequest) (response *ModifySnapshotsSharePermissionResponse, err error) {
    return c.ModifySnapshotsSharePermissionWithContext(context.Background(), request)
}

// ModifySnapshotsSharePermission
// 本接口（ModifySnapshotsSharePermission）用于修改快照分享信息。
//
// 
//
// 分享快照后，被分享账户可以通过该快照创建云硬盘。
//
// * 每个快照最多可分享给50个账户。
//
// * 分享快照无法更改名称，描述，仅可用于创建云硬盘。
//
// * 只支持分享到对方账户相同地域。
//
// * 仅支持分享数据盘快照。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
func (c *Client) ModifySnapshotsSharePermissionWithContext(ctx context.Context, request *ModifySnapshotsSharePermissionRequest) (response *ModifySnapshotsSharePermissionResponse, err error) {
    if request == nil {
        request = NewModifySnapshotsSharePermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotsSharePermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotsSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDiskRequest() (request *RenewDiskRequest) {
    request = &RenewDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "RenewDisk")
    
    
    return
}

func NewRenewDiskResponse() (response *RenewDiskResponse) {
    response = &RenewDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewDisk
// 本接口（RenewDisk）用于续费云硬盘。
//
// 
//
// * 只支持预付费的云硬盘。云硬盘类型可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中DiskChargeType字段解释。
//
// * 支持与挂载实例一起续费的场景，需要在[DiskChargePrepaid](/document/product/362/15669#DiskChargePrepaid)参数中指定CurInstanceDeadline，此时会按对齐到子机续费后的到期时间来续费。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
func (c *Client) RenewDisk(request *RenewDiskRequest) (response *RenewDiskResponse, err error) {
    return c.RenewDiskWithContext(context.Background(), request)
}

// RenewDisk
// 本接口（RenewDisk）用于续费云硬盘。
//
// 
//
// * 只支持预付费的云硬盘。云硬盘类型可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中DiskChargeType字段解释。
//
// * 支持与挂载实例一起续费的场景，需要在[DiskChargePrepaid](/document/product/362/15669#DiskChargePrepaid)参数中指定CurInstanceDeadline，此时会按对齐到子机续费后的到期时间来续费。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
func (c *Client) RenewDiskWithContext(ctx context.Context, request *RenewDiskRequest) (response *RenewDiskResponse, err error) {
    if request == nil {
        request = NewRenewDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDiskResponse()
    err = c.Send(request, response)
    return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
    request = &ResizeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "ResizeDisk")
    
    
    return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
    response = &ResizeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResizeDisk
// 本接口（ResizeDisk）用于扩容云硬盘。
//
// 
//
// * 只支持扩容弹性云盘。云硬盘类型可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中Portable字段解释。非弹性云硬盘需通过[ResizeInstanceDisks](/document/product/213/15731)接口扩容。
//
// * 本接口为异步接口，接口成功返回时，云盘并未立即扩容到指定大小，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态为“EXPANDING”，表示正在扩容中。 
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNSUPPORTEDOPERATION_INSTANCENOTSTOPPED = "UnsupportedOperation.InstanceNotStopped"
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    return c.ResizeDiskWithContext(context.Background(), request)
}

// ResizeDisk
// 本接口（ResizeDisk）用于扩容云硬盘。
//
// 
//
// * 只支持扩容弹性云盘。云硬盘类型可以通过[DescribeDisks](/document/product/362/16315)接口查询，见输出参数中Portable字段解释。非弹性云硬盘需通过[ResizeInstanceDisks](/document/product/213/15731)接口扩容。
//
// * 本接口为异步接口，接口成功返回时，云盘并未立即扩容到指定大小，可通过接口[DescribeDisks](/document/product/362/16315)来查询对应云盘的状态，如果云盘的状态为“EXPANDING”，表示正在扩容中。 
//
// 可能返回的错误码:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNSUPPORTEDOPERATION_INSTANCENOTSTOPPED = "UnsupportedOperation.InstanceNotStopped"
func (c *Client) ResizeDiskWithContext(ctx context.Context, request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    if request == nil {
        request = NewResizeDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeDiskResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDisksRequest() (request *TerminateDisksRequest) {
    request = &TerminateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "TerminateDisks")
    
    
    return
}

func NewTerminateDisksResponse() (response *TerminateDisksResponse) {
    response = &TerminateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateDisks
// 本接口（TerminateDisks）用于退还云硬盘。
//
// 
//
// * 不再使用的云盘，可通过本接口主动退还。
//
// * 本接口支持退还预付费云盘和按小时后付费云盘。按小时后付费云盘可直接退还，预付费云盘需符合退还规则。
//
// * 支持批量操作，每次请求批量云硬盘的上限为100。如果批量云盘存在不允许操作的，请求会以特定错误码返回。
//
// 可能返回的错误码:
//  INSUFFICIENTREFUNDQUOTA = "InsufficientRefundQuota"
//  INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT_OVERREFUNDQUOTA = "ResourceInsufficient.OverRefundQuota"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTREFUND = "ResourceUnavailable.NotSupportRefund"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TerminateDisks(request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    return c.TerminateDisksWithContext(context.Background(), request)
}

// TerminateDisks
// 本接口（TerminateDisks）用于退还云硬盘。
//
// 
//
// * 不再使用的云盘，可通过本接口主动退还。
//
// * 本接口支持退还预付费云盘和按小时后付费云盘。按小时后付费云盘可直接退还，预付费云盘需符合退还规则。
//
// * 支持批量操作，每次请求批量云硬盘的上限为100。如果批量云盘存在不允许操作的，请求会以特定错误码返回。
//
// 可能返回的错误码:
//  INSUFFICIENTREFUNDQUOTA = "InsufficientRefundQuota"
//  INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT_OVERREFUNDQUOTA = "ResourceInsufficient.OverRefundQuota"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTREFUND = "ResourceUnavailable.NotSupportRefund"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TerminateDisksWithContext(ctx context.Context, request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    if request == nil {
        request = NewTerminateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindAutoSnapshotPolicyRequest() (request *UnbindAutoSnapshotPolicyRequest) {
    request = &UnbindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cbs", APIVersion, "UnbindAutoSnapshotPolicy")
    
    
    return
}

func NewUnbindAutoSnapshotPolicyResponse() (response *UnbindAutoSnapshotPolicyResponse) {
    response = &UnbindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindAutoSnapshotPolicy
// 本接口（UnbindAutoSnapshotPolicy）用于解除云硬盘绑定的定期快照策略。
//
// 
//
// * 支持批量操作，可一次解除多个云盘与同一定期快照策略的绑定。 
//
// * 如果传入的云盘未绑定到当前定期快照策略，接口将自动跳过，仅解绑与当前定期快照策略绑定的云盘。
//
// 可能返回的错误码:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) UnbindAutoSnapshotPolicy(request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
    return c.UnbindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// UnbindAutoSnapshotPolicy
// 本接口（UnbindAutoSnapshotPolicy）用于解除云硬盘绑定的定期快照策略。
//
// 
//
// * 支持批量操作，可一次解除多个云盘与同一定期快照策略的绑定。 
//
// * 如果传入的云盘未绑定到当前定期快照策略，接口将自动跳过，仅解绑与当前定期快照策略绑定的云盘。
//
// 可能返回的错误码:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
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
