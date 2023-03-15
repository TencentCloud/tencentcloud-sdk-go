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

package v20200324

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-24"

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
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ApplyDiskBackup")
    
    
    return
}

func NewApplyDiskBackupResponse() (response *ApplyDiskBackupResponse) {
    response = &ApplyDiskBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyDiskBackup
// 本接口（ApplyDiskBackup）用于回滚指定云硬盘的备份点。
//
// * 仅支持回滚到原云硬盘。
//
// * 用于回滚的云硬盘备份点必须处于 NORMAL 状态。
//
//   云硬盘备份点状态可以通过  [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379) 接口查询。
//
// * 回滚云硬盘备份点时，云硬盘的状态必须为 UNATTACHED或ATTACHED。
//
//   云硬盘状态可通过 [DescribeDisks](https://cloud.tencent.com/document/api/1207/66093) 接口查询。
//
// * 如果云硬盘处于 ATTACHED状态，相关RUNNING 状态的实例会强制关机，然后回滚云硬盘备份点。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  OPERATIONDENIED_DISKBACKUPBUSY = "OperationDenied.DiskBackupBusy"
//  OPERATIONDENIED_DISKBACKUPOPERATIONINPROGRESS = "OperationDenied.DiskBackupOperationInProgress"
//  OPERATIONDENIED_DISKBUSYFORBACKUPOPERATION = "OperationDenied.DiskBusyForBackupOperation"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKATTACHEDHASNOINSTANCEID = "ResourceNotFound.DiskAttachedHasNoInstanceId"
//  RESOURCENOTFOUND_DISKBACKUPNOTEXISTS = "ResourceNotFound.DiskBackupNotExists"
//  RESOURCENOTFOUND_DISKNOTEXISTS = "ResourceNotFound.DiskNotExists"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_APPLYDISKBACKUPTOANOTHERDISK = "UnsupportedOperation.ApplyDiskBackupToAnotherDisk"
//  UNSUPPORTEDOPERATION_INVALIDDISKBACKUPSTATE = "UnsupportedOperation.InvalidDiskBackupState"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ApplyDiskBackup(request *ApplyDiskBackupRequest) (response *ApplyDiskBackupResponse, err error) {
    return c.ApplyDiskBackupWithContext(context.Background(), request)
}

// ApplyDiskBackup
// 本接口（ApplyDiskBackup）用于回滚指定云硬盘的备份点。
//
// * 仅支持回滚到原云硬盘。
//
// * 用于回滚的云硬盘备份点必须处于 NORMAL 状态。
//
//   云硬盘备份点状态可以通过  [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379) 接口查询。
//
// * 回滚云硬盘备份点时，云硬盘的状态必须为 UNATTACHED或ATTACHED。
//
//   云硬盘状态可通过 [DescribeDisks](https://cloud.tencent.com/document/api/1207/66093) 接口查询。
//
// * 如果云硬盘处于 ATTACHED状态，相关RUNNING 状态的实例会强制关机，然后回滚云硬盘备份点。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  OPERATIONDENIED_DISKBACKUPBUSY = "OperationDenied.DiskBackupBusy"
//  OPERATIONDENIED_DISKBACKUPOPERATIONINPROGRESS = "OperationDenied.DiskBackupOperationInProgress"
//  OPERATIONDENIED_DISKBUSYFORBACKUPOPERATION = "OperationDenied.DiskBusyForBackupOperation"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKATTACHEDHASNOINSTANCEID = "ResourceNotFound.DiskAttachedHasNoInstanceId"
//  RESOURCENOTFOUND_DISKBACKUPNOTEXISTS = "ResourceNotFound.DiskBackupNotExists"
//  RESOURCENOTFOUND_DISKNOTEXISTS = "ResourceNotFound.DiskNotExists"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_APPLYDISKBACKUPTOANOTHERDISK = "UnsupportedOperation.ApplyDiskBackupToAnotherDisk"
//  UNSUPPORTEDOPERATION_INVALIDDISKBACKUPSTATE = "UnsupportedOperation.InvalidDiskBackupState"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
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

func NewApplyInstanceSnapshotRequest() (request *ApplyInstanceSnapshotRequest) {
    request = &ApplyInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ApplyInstanceSnapshot")
    
    
    return
}

func NewApplyInstanceSnapshotResponse() (response *ApplyInstanceSnapshotResponse) {
    response = &ApplyInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyInstanceSnapshot
// 本接口（ApplyInstanceSnapshot）用于回滚指定实例的系统盘快照。
//
// <li>仅支持回滚到原系统盘。</li>
//
// <li>用于回滚的快照必须处于 NORMAL 状态。快照状态可以通 DescribeSnapshots 接口查询，见输出参数中 SnapshotState 字段解释。</li>
//
// <li>回滚快照时，实例的状态必须为 STOPPED 或 RUNNING，可通过 DescribeInstances 接口查询实例状态。处于 RUNNING 状态的实例会强制关机，然后回滚快照。</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DISKSIZENOTMATCH = "InvalidParameterValue.DiskSizeNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ApplyInstanceSnapshot(request *ApplyInstanceSnapshotRequest) (response *ApplyInstanceSnapshotResponse, err error) {
    return c.ApplyInstanceSnapshotWithContext(context.Background(), request)
}

// ApplyInstanceSnapshot
// 本接口（ApplyInstanceSnapshot）用于回滚指定实例的系统盘快照。
//
// <li>仅支持回滚到原系统盘。</li>
//
// <li>用于回滚的快照必须处于 NORMAL 状态。快照状态可以通 DescribeSnapshots 接口查询，见输出参数中 SnapshotState 字段解释。</li>
//
// <li>回滚快照时，实例的状态必须为 STOPPED 或 RUNNING，可通过 DescribeInstances 接口查询实例状态。处于 RUNNING 状态的实例会强制关机，然后回滚快照。</li>
//
// 可能返回的错误码:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DISKSIZENOTMATCH = "InvalidParameterValue.DiskSizeNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ApplyInstanceSnapshotWithContext(ctx context.Context, request *ApplyInstanceSnapshotRequest) (response *ApplyInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewApplyInstanceSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyInstanceSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateInstancesKeyPairsRequest() (request *AssociateInstancesKeyPairsRequest) {
    request = &AssociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "AssociateInstancesKeyPairs")
    
    
    return
}

func NewAssociateInstancesKeyPairsResponse() (response *AssociateInstancesKeyPairsResponse) {
    response = &AssociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateInstancesKeyPairs
// 本接口（AssociateInstancesKeyPairs）用于绑定用户指定密钥对到实例。
//
// * 只支持 [RUNNING, STOPPED] 状态的 LINUX_UNIX 操作系统的实例。处于 RUNNING 状态的实例会强制关机，然后绑定。
//
// * 将密钥的公钥写入到实例的 SSH 配置当中，用户就可以通过该密钥的私钥来登录实例。
//
// * 如果实例原来绑定过密钥，那么原来的密钥将失效。
//
// * 如果实例原来是通过密码登录，绑定密钥后无法使用密码登录。
//
// * 支持批量操作。每次请求批量实例的上限为 100。如果批量实例存在不允许操作的实例，操作会以特定错误码返回。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDDUPLICATE = "UnsupportedOperation.KeyPairBindDuplicate"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_WINDOWSNOTALLOWTOASSOCIATEKEYPAIR = "UnsupportedOperation.WindowsNotAllowToAssociateKeyPair"
//  UNSUPPORTEDOPERATION_WINDOWSNOTSUPPORTKEYPAIR = "UnsupportedOperation.WindowsNotSupportKeyPair"
func (c *Client) AssociateInstancesKeyPairs(request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    return c.AssociateInstancesKeyPairsWithContext(context.Background(), request)
}

// AssociateInstancesKeyPairs
// 本接口（AssociateInstancesKeyPairs）用于绑定用户指定密钥对到实例。
//
// * 只支持 [RUNNING, STOPPED] 状态的 LINUX_UNIX 操作系统的实例。处于 RUNNING 状态的实例会强制关机，然后绑定。
//
// * 将密钥的公钥写入到实例的 SSH 配置当中，用户就可以通过该密钥的私钥来登录实例。
//
// * 如果实例原来绑定过密钥，那么原来的密钥将失效。
//
// * 如果实例原来是通过密码登录，绑定密钥后无法使用密码登录。
//
// * 支持批量操作。每次请求批量实例的上限为 100。如果批量实例存在不允许操作的实例，操作会以特定错误码返回。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDDUPLICATE = "UnsupportedOperation.KeyPairBindDuplicate"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_WINDOWSNOTALLOWTOASSOCIATEKEYPAIR = "UnsupportedOperation.WindowsNotAllowToAssociateKeyPair"
//  UNSUPPORTEDOPERATION_WINDOWSNOTSUPPORTKEYPAIR = "UnsupportedOperation.WindowsNotSupportKeyPair"
func (c *Client) AssociateInstancesKeyPairsWithContext(ctx context.Context, request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewAssociateInstancesKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateInstancesKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewAttachCcnRequest() (request *AttachCcnRequest) {
    request = &AttachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "AttachCcn")
    
    
    return
}

func NewAttachCcnResponse() (response *AttachCcnResponse) {
    response = &AttachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachCcn
// 本接口 (AttachCcn) 用于建立与云联网的关联。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHCCNCONDITIONUNSATISFIED = "UnsupportedOperation.AttachCcnConditionUnsatisfied"
//  UNSUPPORTEDOPERATION_ATTACHCCNFAILED = "UnsupportedOperation.AttachCcnFailed"
//  UNSUPPORTEDOPERATION_CCNALREADYATTACHED = "UnsupportedOperation.CcnAlreadyAttached"
func (c *Client) AttachCcn(request *AttachCcnRequest) (response *AttachCcnResponse, err error) {
    return c.AttachCcnWithContext(context.Background(), request)
}

// AttachCcn
// 本接口 (AttachCcn) 用于建立与云联网的关联。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHCCNCONDITIONUNSATISFIED = "UnsupportedOperation.AttachCcnConditionUnsatisfied"
//  UNSUPPORTEDOPERATION_ATTACHCCNFAILED = "UnsupportedOperation.AttachCcnFailed"
//  UNSUPPORTEDOPERATION_CCNALREADYATTACHED = "UnsupportedOperation.CcnAlreadyAttached"
func (c *Client) AttachCcnWithContext(ctx context.Context, request *AttachCcnRequest) (response *AttachCcnResponse, err error) {
    if request == nil {
        request = NewAttachCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachCcn require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
    request = &AttachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "AttachDisks")
    
    
    return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
    response = &AttachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachDisks
// 本接口（AttachDisks）用于挂载一个或多个云硬盘。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  LIMITEXCEEDED_ATTACHDATADISKQUOTALIMITEXCEEDED = "LimitExceeded.AttachDataDiskQuotaLimitExceeded"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    return c.AttachDisksWithContext(context.Background(), request)
}

// AttachDisks
// 本接口（AttachDisks）用于挂载一个或多个云硬盘。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  LIMITEXCEEDED_ATTACHDATADISKQUOTALIMITEXCEEDED = "LimitExceeded.AttachDataDiskQuotaLimitExceeded"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
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

func NewCreateBlueprintRequest() (request *CreateBlueprintRequest) {
    request = &CreateBlueprintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateBlueprint")
    
    
    return
}

func NewCreateBlueprintResponse() (response *CreateBlueprintResponse) {
    response = &CreateBlueprintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBlueprint
// 本接口 (CreateBlueprint) 用于创建镜像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEBLUEPRINTFAILED = "FailedOperation.CreateBlueprintFailed"
//  FAILEDOPERATION_UNABLETOCREATEBLUEPRINT = "FailedOperation.UnableToCreateBlueprint"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateBlueprint(request *CreateBlueprintRequest) (response *CreateBlueprintResponse, err error) {
    return c.CreateBlueprintWithContext(context.Background(), request)
}

// CreateBlueprint
// 本接口 (CreateBlueprint) 用于创建镜像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEBLUEPRINTFAILED = "FailedOperation.CreateBlueprintFailed"
//  FAILEDOPERATION_UNABLETOCREATEBLUEPRINT = "FailedOperation.UnableToCreateBlueprint"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateBlueprintWithContext(ctx context.Context, request *CreateBlueprintRequest) (response *CreateBlueprintResponse, err error) {
    if request == nil {
        request = NewCreateBlueprintRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlueprint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBlueprintResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDiskBackupRequest() (request *CreateDiskBackupRequest) {
    request = &CreateDiskBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateDiskBackup")
    
    
    return
}

func NewCreateDiskBackupResponse() (response *CreateDiskBackupResponse) {
    response = &CreateDiskBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDiskBackup
// 本接口 ( CreateDiskBackup  ) 用于创建指定云硬盘（当前只支持数据盘）的备份点。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  LIMITEXCEEDED_DISKBACKUPQUOTALIMITEXCEEDED = "LimitExceeded.DiskBackupQuotaLimitExceeded"
//  OPERATIONDENIED_DISKBUSYFORBACKUPOPERATION = "OperationDenied.DiskBusyForBackupOperation"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKATTACHEDHASNOINSTANCEID = "ResourceNotFound.DiskAttachedHasNoInstanceId"
//  RESOURCENOTFOUND_DISKNOTEXISTS = "ResourceNotFound.DiskNotExists"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateDiskBackup(request *CreateDiskBackupRequest) (response *CreateDiskBackupResponse, err error) {
    return c.CreateDiskBackupWithContext(context.Background(), request)
}

// CreateDiskBackup
// 本接口 ( CreateDiskBackup  ) 用于创建指定云硬盘（当前只支持数据盘）的备份点。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  LIMITEXCEEDED_DISKBACKUPQUOTALIMITEXCEEDED = "LimitExceeded.DiskBackupQuotaLimitExceeded"
//  OPERATIONDENIED_DISKBUSYFORBACKUPOPERATION = "OperationDenied.DiskBusyForBackupOperation"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKATTACHEDHASNOINSTANCEID = "ResourceNotFound.DiskAttachedHasNoInstanceId"
//  RESOURCENOTFOUND_DISKNOTEXISTS = "ResourceNotFound.DiskNotExists"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
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

func NewCreateFirewallRulesRequest() (request *CreateFirewallRulesRequest) {
    request = &CreateFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateFirewallRules")
    
    
    return
}

func NewCreateFirewallRulesResponse() (response *CreateFirewallRulesResponse) {
    response = &CreateFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFirewallRules
// 本接口（CreateFirewallRules）用于在实例上添加防火墙规则。
//
// 
//
// 
//
// * FirewallVersion 为防火墙版本号，用户每次更新防火墙规则版本会自动加1，防止您更新的规则已过期，不填不考虑冲突。
//
// 
//
// 在 FirewallRules 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETER_FIREWALLRULESEXIST = "InvalidParameter.FirewallRulesExist"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_FIREWALLRULESLIMITEXCEEDED = "LimitExceeded.FirewallRulesLimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) CreateFirewallRules(request *CreateFirewallRulesRequest) (response *CreateFirewallRulesResponse, err error) {
    return c.CreateFirewallRulesWithContext(context.Background(), request)
}

// CreateFirewallRules
// 本接口（CreateFirewallRules）用于在实例上添加防火墙规则。
//
// 
//
// 
//
// * FirewallVersion 为防火墙版本号，用户每次更新防火墙规则版本会自动加1，防止您更新的规则已过期，不填不考虑冲突。
//
// 
//
// 在 FirewallRules 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETER_FIREWALLRULESEXIST = "InvalidParameter.FirewallRulesExist"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_FIREWALLRULESLIMITEXCEEDED = "LimitExceeded.FirewallRulesLimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) CreateFirewallRulesWithContext(ctx context.Context, request *CreateFirewallRulesRequest) (response *CreateFirewallRulesResponse, err error) {
    if request == nil {
        request = NewCreateFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceSnapshotRequest() (request *CreateInstanceSnapshotRequest) {
    request = &CreateInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateInstanceSnapshot")
    
    
    return
}

func NewCreateInstanceSnapshotResponse() (response *CreateInstanceSnapshotResponse) {
    response = &CreateInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceSnapshot
// 本接口（CreateInstanceSnapshot）用于创建指定实例的系统盘快照。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_GETSNAPSHOTALLOCQUOTALOCKERROR = "InternalError.GetSnapshotAllocQuotaLockError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  LIMITEXCEEDED_SNAPSHOTQUOTALIMITEXCEEDED = "LimitExceeded.SnapshotQuotaLimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOTFORSTORAGEBUNDLE = "OperationDenied.OperationDeniedCreateSnapshotForStorageBundle"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateInstanceSnapshot(request *CreateInstanceSnapshotRequest) (response *CreateInstanceSnapshotResponse, err error) {
    return c.CreateInstanceSnapshotWithContext(context.Background(), request)
}

// CreateInstanceSnapshot
// 本接口（CreateInstanceSnapshot）用于创建指定实例的系统盘快照。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_GETSNAPSHOTALLOCQUOTALOCKERROR = "InternalError.GetSnapshotAllocQuotaLockError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  LIMITEXCEEDED_SNAPSHOTQUOTALIMITEXCEEDED = "LimitExceeded.SnapshotQuotaLimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOTFORSTORAGEBUNDLE = "OperationDenied.OperationDeniedCreateSnapshotForStorageBundle"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateInstanceSnapshotWithContext(ctx context.Context, request *CreateInstanceSnapshotRequest) (response *CreateInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateInstanceSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstances
// 本接口(CreateInstances)用于创建一个或多个指定套餐的轻量应用服务器实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEINSTANCESFAILED = "FailedOperation.CreateInstancesFailed"
//  FAILEDOPERATION_UNABLETOCREATEINSTANCES = "FailedOperation.UnableToCreateInstances"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDACTIONNOTFOUND = "InternalError.InvalidActionNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameter.BundleAndBlueprintNotMatch"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameterValue.BundleAndBlueprintNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  INVALIDPARAMETERVALUE_INVALIDBUNDLE = "InvalidParameterValue.InvalidBundle"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERCOMBINATION = "InvalidParameterValue.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INSTANCEQUOTALIMITEXCEEDED = "LimitExceeded.InstanceQuotaLimitExceeded"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE_BUNDLEUNAVAILABLE = "ResourceUnavailable.BundleUnavailable"
//  RESOURCESSOLDOUT_PURCHASESOURCEHASNOBUNDLECONFIGS = "ResourcesSoldOut.PurchaseSourceHasNoBundleConfigs"
//  RESOURCESSOLDOUT_ZONESHASNOBUNDLECONFIGS = "ResourcesSoldOut.ZonesHasNoBundleConfigs"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INSTANCELINUXUNIXCREATINGNOTSUPPORTPASSWORD = "UnsupportedOperation.InstanceLinuxUnixCreatingNotSupportPassword"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// 本接口(CreateInstances)用于创建一个或多个指定套餐的轻量应用服务器实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATEINSTANCESFAILED = "FailedOperation.CreateInstancesFailed"
//  FAILEDOPERATION_UNABLETOCREATEINSTANCES = "FailedOperation.UnableToCreateInstances"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDACTIONNOTFOUND = "InternalError.InvalidActionNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameter.BundleAndBlueprintNotMatch"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameterValue.BundleAndBlueprintNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  INVALIDPARAMETERVALUE_INVALIDBUNDLE = "InvalidParameterValue.InvalidBundle"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERCOMBINATION = "InvalidParameterValue.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INSTANCEQUOTALIMITEXCEEDED = "LimitExceeded.InstanceQuotaLimitExceeded"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE_BUNDLEUNAVAILABLE = "ResourceUnavailable.BundleUnavailable"
//  RESOURCESSOLDOUT_PURCHASESOURCEHASNOBUNDLECONFIGS = "ResourcesSoldOut.PurchaseSourceHasNoBundleConfigs"
//  RESOURCESSOLDOUT_ZONESHASNOBUNDLECONFIGS = "ResourcesSoldOut.ZonesHasNoBundleConfigs"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INSTANCELINUXUNIXCREATINGNOTSUPPORTPASSWORD = "UnsupportedOperation.InstanceLinuxUnixCreatingNotSupportPassword"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
    request = &CreateKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateKeyPair")
    
    
    return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
    response = &CreateKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKeyPair
// 本接口（CreateKeyPair）用于创建一个密钥对。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    return c.CreateKeyPairWithContext(context.Background(), request)
}

// CreateKeyPair
// 本接口（CreateKeyPair）用于创建一个密钥对。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateKeyPairWithContext(ctx context.Context, request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKeyPair require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlueprintsRequest() (request *DeleteBlueprintsRequest) {
    request = &DeleteBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteBlueprints")
    
    
    return
}

func NewDeleteBlueprintsResponse() (response *DeleteBlueprintsResponse) {
    response = &DeleteBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBlueprints
// 本接口 (DeleteBlueprints) 用于删除镜像。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_BLUEPRINTOCCUPIED = "UnsupportedOperation.BlueprintOccupied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
//  UNSUPPORTEDOPERATION_POSTDESTROYRESOURCEFAILED = "UnsupportedOperation.PostDestroyResourceFailed"
func (c *Client) DeleteBlueprints(request *DeleteBlueprintsRequest) (response *DeleteBlueprintsResponse, err error) {
    return c.DeleteBlueprintsWithContext(context.Background(), request)
}

// DeleteBlueprints
// 本接口 (DeleteBlueprints) 用于删除镜像。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_BLUEPRINTOCCUPIED = "UnsupportedOperation.BlueprintOccupied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
//  UNSUPPORTEDOPERATION_POSTDESTROYRESOURCEFAILED = "UnsupportedOperation.PostDestroyResourceFailed"
func (c *Client) DeleteBlueprintsWithContext(ctx context.Context, request *DeleteBlueprintsRequest) (response *DeleteBlueprintsResponse, err error) {
    if request == nil {
        request = NewDeleteBlueprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBlueprints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDiskBackupsRequest() (request *DeleteDiskBackupsRequest) {
    request = &DeleteDiskBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteDiskBackups")
    
    
    return
}

func NewDeleteDiskBackupsResponse() (response *DeleteDiskBackupsResponse) {
    response = &DeleteDiskBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDiskBackups
// 本接口（DeleteDiskBackups）用于删除云硬盘备份点。
//
// 云硬盘备份点必须处于 NORMAL 状态，云硬盘备份点状态可以通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询，见输出参数中 DiskBackupState 字段解释。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  OPERATIONDENIED_DISKBACKUPOPERATIONINPROGRESS = "OperationDenied.DiskBackupOperationInProgress"
//  RESOURCEINUSE_DISKBACKUPINUSE = "ResourceInUse.DiskBackupInUse"
//  RESOURCENOTFOUND_DISKBACKUPIDNOTFOUND = "ResourceNotFound.DiskBackupIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_DISKBACKUPLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskBackupLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKBACKUPSTATE = "UnsupportedOperation.InvalidDiskBackupState"
func (c *Client) DeleteDiskBackups(request *DeleteDiskBackupsRequest) (response *DeleteDiskBackupsResponse, err error) {
    return c.DeleteDiskBackupsWithContext(context.Background(), request)
}

// DeleteDiskBackups
// 本接口（DeleteDiskBackups）用于删除云硬盘备份点。
//
// 云硬盘备份点必须处于 NORMAL 状态，云硬盘备份点状态可以通过 [DescribeDiskBackups](https://cloud.tencent.com/document/api/1207/84379)接口查询，见输出参数中 DiskBackupState 字段解释。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  OPERATIONDENIED_DISKBACKUPOPERATIONINPROGRESS = "OperationDenied.DiskBackupOperationInProgress"
//  RESOURCEINUSE_DISKBACKUPINUSE = "ResourceInUse.DiskBackupInUse"
//  RESOURCENOTFOUND_DISKBACKUPIDNOTFOUND = "ResourceNotFound.DiskBackupIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_DISKBACKUPLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskBackupLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKBACKUPSTATE = "UnsupportedOperation.InvalidDiskBackupState"
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

func NewDeleteFirewallRulesRequest() (request *DeleteFirewallRulesRequest) {
    request = &DeleteFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteFirewallRules")
    
    
    return
}

func NewDeleteFirewallRulesResponse() (response *DeleteFirewallRulesResponse) {
    response = &DeleteFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFirewallRules
// 本接口（DeleteFirewallRules）用于删除实例的防火墙规则。
//
// 
//
// * FirewallVersion 用于指定要操作的防火墙的版本。传入 FirewallVersion 版本号若不等于当前防火墙的最新版本，将返回失败；若不传 FirewallVersion 则直接删除指定的规则。
//
// 
//
// 在 FirewallRules 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) DeleteFirewallRules(request *DeleteFirewallRulesRequest) (response *DeleteFirewallRulesResponse, err error) {
    return c.DeleteFirewallRulesWithContext(context.Background(), request)
}

// DeleteFirewallRules
// 本接口（DeleteFirewallRules）用于删除实例的防火墙规则。
//
// 
//
// * FirewallVersion 用于指定要操作的防火墙的版本。传入 FirewallVersion 版本号若不等于当前防火墙的最新版本，将返回失败；若不传 FirewallVersion 则直接删除指定的规则。
//
// 
//
// 在 FirewallRules 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) DeleteFirewallRulesWithContext(ctx context.Context, request *DeleteFirewallRulesRequest) (response *DeleteFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDeleteFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKeyPairsRequest() (request *DeleteKeyPairsRequest) {
    request = &DeleteKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteKeyPairs")
    
    
    return
}

func NewDeleteKeyPairsResponse() (response *DeleteKeyPairsResponse) {
    response = &DeleteKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteKeyPairs
// 本接口（DeleteKeyPairs）用于删除密钥对。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  RESOURCEINUSE_KEYPAIRINUSE = "ResourceInUse.KeyPairInUse"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDTOBLUEPRINTS = "UnsupportedOperation.KeyPairBindToBlueprints"
func (c *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    return c.DeleteKeyPairsWithContext(context.Background(), request)
}

// DeleteKeyPairs
// 本接口（DeleteKeyPairs）用于删除密钥对。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  RESOURCEINUSE_KEYPAIRINUSE = "ResourceInUse.KeyPairInUse"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDTOBLUEPRINTS = "UnsupportedOperation.KeyPairBindToBlueprints"
func (c *Client) DeleteKeyPairsWithContext(ctx context.Context, request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    if request == nil {
        request = NewDeleteKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteSnapshots")
    
    
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
// 快照必须处于 NORMAL 状态，快照状态可以通过 DescribeSnapshots 接口查询，见输出参数中 SnapshotState 字段解释。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    return c.DeleteSnapshotsWithContext(context.Background(), request)
}

// DeleteSnapshots
// 本接口（DeleteSnapshots）用于删除快照。
//
// 快照必须处于 NORMAL 状态，快照状态可以通过 DescribeSnapshots 接口查询，见输出参数中 SnapshotState 字段解释。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
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

func NewDescribeAllScenesRequest() (request *DescribeAllScenesRequest) {
    request = &DescribeAllScenesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeAllScenes")
    
    
    return
}

func NewDescribeAllScenesResponse() (response *DescribeAllScenesResponse) {
    response = &DescribeAllScenesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAllScenes
// 本接口(DescribeAllScenes)用于查询全地域使用场景列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAllScenes(request *DescribeAllScenesRequest) (response *DescribeAllScenesResponse, err error) {
    return c.DescribeAllScenesWithContext(context.Background(), request)
}

// DescribeAllScenes
// 本接口(DescribeAllScenes)用于查询全地域使用场景列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAllScenesWithContext(ctx context.Context, request *DescribeAllScenesRequest) (response *DescribeAllScenesResponse, err error) {
    if request == nil {
        request = NewDescribeAllScenesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllScenes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllScenesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlueprintInstancesRequest() (request *DescribeBlueprintInstancesRequest) {
    request = &DescribeBlueprintInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBlueprintInstances")
    
    
    return
}

func NewDescribeBlueprintInstancesResponse() (response *DescribeBlueprintInstancesResponse) {
    response = &DescribeBlueprintInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlueprintInstances
// 本接口（DescribeBlueprintInstances）用于查询镜像实例信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprintInstances(request *DescribeBlueprintInstancesRequest) (response *DescribeBlueprintInstancesResponse, err error) {
    return c.DescribeBlueprintInstancesWithContext(context.Background(), request)
}

// DescribeBlueprintInstances
// 本接口（DescribeBlueprintInstances）用于查询镜像实例信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprintInstancesWithContext(ctx context.Context, request *DescribeBlueprintInstancesRequest) (response *DescribeBlueprintInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeBlueprintInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlueprintInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlueprintInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlueprintsRequest() (request *DescribeBlueprintsRequest) {
    request = &DescribeBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBlueprints")
    
    
    return
}

func NewDescribeBlueprintsResponse() (response *DescribeBlueprintsResponse) {
    response = &DescribeBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlueprints
// 本接口（DescribeBlueprints）用于查询镜像信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEBLUEPRINTSFAILED = "FailedOperation.DescribeBlueprintsFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  RESOURCENOTFOUND_SCENEIDNOTFOUND = "ResourceNotFound.SceneIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprints(request *DescribeBlueprintsRequest) (response *DescribeBlueprintsResponse, err error) {
    return c.DescribeBlueprintsWithContext(context.Background(), request)
}

// DescribeBlueprints
// 本接口（DescribeBlueprints）用于查询镜像信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEBLUEPRINTSFAILED = "FailedOperation.DescribeBlueprintsFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"
//  RESOURCENOTFOUND_SCENEIDNOTFOUND = "ResourceNotFound.SceneIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprintsWithContext(ctx context.Context, request *DescribeBlueprintsRequest) (response *DescribeBlueprintsResponse, err error) {
    if request == nil {
        request = NewDescribeBlueprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlueprints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBundleDiscountRequest() (request *DescribeBundleDiscountRequest) {
    request = &DescribeBundleDiscountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBundleDiscount")
    
    
    return
}

func NewDescribeBundleDiscountResponse() (response *DescribeBundleDiscountResponse) {
    response = &DescribeBundleDiscountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBundleDiscount
// 本接口（DescribeBundleDiscount）用于查询套餐折扣信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
func (c *Client) DescribeBundleDiscount(request *DescribeBundleDiscountRequest) (response *DescribeBundleDiscountResponse, err error) {
    return c.DescribeBundleDiscountWithContext(context.Background(), request)
}

// DescribeBundleDiscount
// 本接口（DescribeBundleDiscount）用于查询套餐折扣信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
func (c *Client) DescribeBundleDiscountWithContext(ctx context.Context, request *DescribeBundleDiscountRequest) (response *DescribeBundleDiscountResponse, err error) {
    if request == nil {
        request = NewDescribeBundleDiscountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBundleDiscount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBundleDiscountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBundlesRequest() (request *DescribeBundlesRequest) {
    request = &DescribeBundlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBundles")
    
    
    return
}

func NewDescribeBundlesResponse() (response *DescribeBundlesResponse) {
    response = &DescribeBundlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBundles
// 本接口（DescribeBundles）用于查询套餐信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
func (c *Client) DescribeBundles(request *DescribeBundlesRequest) (response *DescribeBundlesResponse, err error) {
    return c.DescribeBundlesWithContext(context.Background(), request)
}

// DescribeBundles
// 本接口（DescribeBundles）用于查询套餐信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
func (c *Client) DescribeBundlesWithContext(ctx context.Context, request *DescribeBundlesRequest) (response *DescribeBundlesResponse, err error) {
    if request == nil {
        request = NewDescribeBundlesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBundles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBundlesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnAttachedInstancesRequest() (request *DescribeCcnAttachedInstancesRequest) {
    request = &DescribeCcnAttachedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeCcnAttachedInstances")
    
    
    return
}

func NewDescribeCcnAttachedInstancesResponse() (response *DescribeCcnAttachedInstancesResponse) {
    response = &DescribeCcnAttachedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcnAttachedInstances
// 本接口 (DescribeCcnAttachedInstances) 用于查询云联网关联的实例信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DESCRIBECCNATTACHEDINSTANCESFAILED = "UnsupportedOperation.DescribeCcnAttachedInstancesFailed"
func (c *Client) DescribeCcnAttachedInstances(request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    return c.DescribeCcnAttachedInstancesWithContext(context.Background(), request)
}

// DescribeCcnAttachedInstances
// 本接口 (DescribeCcnAttachedInstances) 用于查询云联网关联的实例信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DESCRIBECCNATTACHEDINSTANCESFAILED = "UnsupportedOperation.DescribeCcnAttachedInstancesFailed"
func (c *Client) DescribeCcnAttachedInstancesWithContext(ctx context.Context, request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnAttachedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcnAttachedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCcnAttachedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskBackupsRequest() (request *DescribeDiskBackupsRequest) {
    request = &DescribeDiskBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDiskBackups")
    
    
    return
}

func NewDescribeDiskBackupsResponse() (response *DescribeDiskBackupsResponse) {
    response = &DescribeDiskBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskBackups
// 本接口（DescribeDiskBackups）用于查询云硬盘备份点的详细信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskBackups(request *DescribeDiskBackupsRequest) (response *DescribeDiskBackupsResponse, err error) {
    return c.DescribeDiskBackupsWithContext(context.Background(), request)
}

// DescribeDiskBackups
// 本接口（DescribeDiskBackups）用于查询云硬盘备份点的详细信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
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

func NewDescribeDiskBackupsDeniedActionsRequest() (request *DescribeDiskBackupsDeniedActionsRequest) {
    request = &DescribeDiskBackupsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDiskBackupsDeniedActions")
    
    
    return
}

func NewDescribeDiskBackupsDeniedActionsResponse() (response *DescribeDiskBackupsDeniedActionsResponse) {
    response = &DescribeDiskBackupsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskBackupsDeniedActions
// 本接口（DescribeDiskBackupsDeniedActions）用于查询一个或多个云硬盘备份点的操作限制列表信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskBackupsDeniedActions(request *DescribeDiskBackupsDeniedActionsRequest) (response *DescribeDiskBackupsDeniedActionsResponse, err error) {
    return c.DescribeDiskBackupsDeniedActionsWithContext(context.Background(), request)
}

// DescribeDiskBackupsDeniedActions
// 本接口（DescribeDiskBackupsDeniedActions）用于查询一个或多个云硬盘备份点的操作限制列表信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskBackupsDeniedActionsWithContext(ctx context.Context, request *DescribeDiskBackupsDeniedActionsRequest) (response *DescribeDiskBackupsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeDiskBackupsDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskBackupsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskBackupsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskConfigsRequest() (request *DescribeDiskConfigsRequest) {
    request = &DescribeDiskConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDiskConfigs")
    
    
    return
}

func NewDescribeDiskConfigsResponse() (response *DescribeDiskConfigsResponse) {
    response = &DescribeDiskConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskConfigs
// 本接口（DescribeDiskConfigs）用于查询云硬盘配置。
//
// 可能返回的错误码:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskConfigs(request *DescribeDiskConfigsRequest) (response *DescribeDiskConfigsResponse, err error) {
    return c.DescribeDiskConfigsWithContext(context.Background(), request)
}

// DescribeDiskConfigs
// 本接口（DescribeDiskConfigs）用于查询云硬盘配置。
//
// 可能返回的错误码:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskConfigsWithContext(ctx context.Context, request *DescribeDiskConfigsRequest) (response *DescribeDiskConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeDiskConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskDiscountRequest() (request *DescribeDiskDiscountRequest) {
    request = &DescribeDiskDiscountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDiskDiscount")
    
    
    return
}

func NewDescribeDiskDiscountResponse() (response *DescribeDiskDiscountResponse) {
    response = &DescribeDiskDiscountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskDiscount
// 本接口(DescribeDiskDiscount)用于查询云硬盘折扣信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskDiscount(request *DescribeDiskDiscountRequest) (response *DescribeDiskDiscountResponse, err error) {
    return c.DescribeDiskDiscountWithContext(context.Background(), request)
}

// DescribeDiskDiscount
// 本接口(DescribeDiskDiscount)用于查询云硬盘折扣信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskDiscountWithContext(ctx context.Context, request *DescribeDiskDiscountRequest) (response *DescribeDiskDiscountResponse, err error) {
    if request == nil {
        request = NewDescribeDiskDiscountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskDiscount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskDiscountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
    request = &DescribeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDisks")
    
    
    return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
    response = &DescribeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisks
// 本接口（DescribeDisks）用于查询云硬盘信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISKNAMETOOLONG = "InvalidParameterValue.DiskNameTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_REGIONNOTFOUND = "InvalidParameterValue.RegionNotFound"
//  INVALIDPARAMETERVALUE_REGIONNOTMATCH = "InvalidParameterValue.RegionNotMatch"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_REGIONUNAVAILABLE = "InvalidParameterValue.RegionUnavailable"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    return c.DescribeDisksWithContext(context.Background(), request)
}

// DescribeDisks
// 本接口（DescribeDisks）用于查询云硬盘信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISKNAMETOOLONG = "InvalidParameterValue.DiskNameTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_REGIONNOTFOUND = "InvalidParameterValue.RegionNotFound"
//  INVALIDPARAMETERVALUE_REGIONNOTMATCH = "InvalidParameterValue.RegionNotMatch"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"
//  INVALIDPARAMETERVALUE_REGIONUNAVAILABLE = "InvalidParameterValue.RegionUnavailable"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
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

func NewDescribeDisksDeniedActionsRequest() (request *DescribeDisksDeniedActionsRequest) {
    request = &DescribeDisksDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDisksDeniedActions")
    
    
    return
}

func NewDescribeDisksDeniedActionsResponse() (response *DescribeDisksDeniedActionsResponse) {
    response = &DescribeDisksDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisksDeniedActions
// 本接口（DescribeDisksDeniedActions）用于查询一个或多个云硬盘的操作限制列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDisksDeniedActions(request *DescribeDisksDeniedActionsRequest) (response *DescribeDisksDeniedActionsResponse, err error) {
    return c.DescribeDisksDeniedActionsWithContext(context.Background(), request)
}

// DescribeDisksDeniedActions
// 本接口（DescribeDisksDeniedActions）用于查询一个或多个云硬盘的操作限制列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDisksDeniedActionsWithContext(ctx context.Context, request *DescribeDisksDeniedActionsRequest) (response *DescribeDisksDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeDisksDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisksDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisksDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksReturnableRequest() (request *DescribeDisksReturnableRequest) {
    request = &DescribeDisksReturnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDisksReturnable")
    
    
    return
}

func NewDescribeDisksReturnableResponse() (response *DescribeDisksReturnableResponse) {
    response = &DescribeDisksReturnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisksReturnable
// 本接口（DescribeDisksReturnable）用于查询云硬盘是否可退还。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBEDISKSRETURNABLEERROR = "InternalError.DescribeDisksReturnableError"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDisksReturnable(request *DescribeDisksReturnableRequest) (response *DescribeDisksReturnableResponse, err error) {
    return c.DescribeDisksReturnableWithContext(context.Background(), request)
}

// DescribeDisksReturnable
// 本接口（DescribeDisksReturnable）用于查询云硬盘是否可退还。
//
// 可能返回的错误码:
//  INTERNALERROR_DESCRIBEDISKSRETURNABLEERROR = "InternalError.DescribeDisksReturnableError"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDisksReturnableWithContext(ctx context.Context, request *DescribeDisksReturnableRequest) (response *DescribeDisksReturnableResponse, err error) {
    if request == nil {
        request = NewDescribeDisksReturnableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisksReturnable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisksReturnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirewallRulesRequest() (request *DescribeFirewallRulesRequest) {
    request = &DescribeFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeFirewallRules")
    
    
    return
}

func NewDescribeFirewallRulesResponse() (response *DescribeFirewallRulesResponse) {
    response = &DescribeFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirewallRules
// 本接口（DescribeFirewallRules）用于查询实例的防火墙规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeFirewallRules(request *DescribeFirewallRulesRequest) (response *DescribeFirewallRulesResponse, err error) {
    return c.DescribeFirewallRulesWithContext(context.Background(), request)
}

// DescribeFirewallRules
// 本接口（DescribeFirewallRules）用于查询实例的防火墙规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeFirewallRulesWithContext(ctx context.Context, request *DescribeFirewallRulesRequest) (response *DescribeFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDescribeFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirewallRulesTemplateRequest() (request *DescribeFirewallRulesTemplateRequest) {
    request = &DescribeFirewallRulesTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeFirewallRulesTemplate")
    
    
    return
}

func NewDescribeFirewallRulesTemplateResponse() (response *DescribeFirewallRulesTemplateResponse) {
    response = &DescribeFirewallRulesTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirewallRulesTemplate
// 本接口（DescribeFirewallRulesTemplate）用于查询防火墙规则模板。
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeFirewallRulesTemplate(request *DescribeFirewallRulesTemplateRequest) (response *DescribeFirewallRulesTemplateResponse, err error) {
    return c.DescribeFirewallRulesTemplateWithContext(context.Background(), request)
}

// DescribeFirewallRulesTemplate
// 本接口（DescribeFirewallRulesTemplate）用于查询防火墙规则模板。
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeFirewallRulesTemplateWithContext(ctx context.Context, request *DescribeFirewallRulesTemplateRequest) (response *DescribeFirewallRulesTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeFirewallRulesTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirewallRulesTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFirewallRulesTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralResourceQuotasRequest() (request *DescribeGeneralResourceQuotasRequest) {
    request = &DescribeGeneralResourceQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeGeneralResourceQuotas")
    
    
    return
}

func NewDescribeGeneralResourceQuotasResponse() (response *DescribeGeneralResourceQuotasResponse) {
    response = &DescribeGeneralResourceQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGeneralResourceQuotas
// 本接口（DescribeGeneralResourceQuotas）用于查询通用资源配额信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDRESOURCEQUOTARESOURCENAME = "InvalidParameterValue.InvalidResourceQuotaResourceName"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeGeneralResourceQuotas(request *DescribeGeneralResourceQuotasRequest) (response *DescribeGeneralResourceQuotasResponse, err error) {
    return c.DescribeGeneralResourceQuotasWithContext(context.Background(), request)
}

// DescribeGeneralResourceQuotas
// 本接口（DescribeGeneralResourceQuotas）用于查询通用资源配额信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDRESOURCEQUOTARESOURCENAME = "InvalidParameterValue.InvalidResourceQuotaResourceName"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeGeneralResourceQuotasWithContext(ctx context.Context, request *DescribeGeneralResourceQuotasRequest) (response *DescribeGeneralResourceQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralResourceQuotasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralResourceQuotas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralResourceQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLoginKeyPairAttributeRequest() (request *DescribeInstanceLoginKeyPairAttributeRequest) {
    request = &DescribeInstanceLoginKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstanceLoginKeyPairAttribute")
    
    
    return
}

func NewDescribeInstanceLoginKeyPairAttributeResponse() (response *DescribeInstanceLoginKeyPairAttributeResponse) {
    response = &DescribeInstanceLoginKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceLoginKeyPairAttribute
// 本接口用于查询实例默认登录密钥属性。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstanceLoginKeyPairAttribute(request *DescribeInstanceLoginKeyPairAttributeRequest) (response *DescribeInstanceLoginKeyPairAttributeResponse, err error) {
    return c.DescribeInstanceLoginKeyPairAttributeWithContext(context.Background(), request)
}

// DescribeInstanceLoginKeyPairAttribute
// 本接口用于查询实例默认登录密钥属性。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstanceLoginKeyPairAttributeWithContext(ctx context.Context, request *DescribeInstanceLoginKeyPairAttributeRequest) (response *DescribeInstanceLoginKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLoginKeyPairAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLoginKeyPairAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLoginKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
    request = &DescribeInstanceVncUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstanceVncUrl")
    
    
    return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
    response = &DescribeInstanceVncUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceVncUrl
// 本接口 ( DescribeInstanceVncUrl ) 用于查询实例管理终端地址，获取的地址可用于实例的 VNC 登录。
//
// 
//
// * 处于 `STOPPED` 状态的机器无法使用此功能。
//
// * 管理终端地址的有效期为 15 秒，调用接口成功后如果 15 秒内不使用该链接进行访问，管理终端地址自动失效，您需要重新查询。
//
// * 管理终端地址一旦被访问，将自动失效，您需要重新查询。
//
// * 如果连接断开，每分钟内重新连接的次数不能超过 30 次。
//
// * 参数 `InstanceVncUrl` ：调用接口成功后会返回的 `InstanceVncUrl` 的值。
//
// 获取到 `InstanceVncUrl` 后，您需要在链接 `https://img.qcloud.com/qcloud/app/active_vnc/index.html?` 末尾加上参数 `InstanceVncUrl=xxxx`。
//
//  最后组成的 URL 格式如下：
//
// 
//
// ```
//
// https://img.qcloud.com/qcloud/app/active_vnc/index.html?InstanceVncUrl=wss%3A%2F%2Fbjvnc.qcloud.com%3A26789%2Fvnc%3Fs%3DaHpjWnRVMFNhYmxKdDM5MjRHNlVTSVQwajNUSW0wb2tBbmFtREFCTmFrcy8vUUNPMG0wSHZNOUUxRm5PMmUzWmFDcWlOdDJIbUJxSTZDL0RXcHZxYnZZMmRkWWZWcEZia2lyb09XMzdKNmM9
//
// ```
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    return c.DescribeInstanceVncUrlWithContext(context.Background(), request)
}

// DescribeInstanceVncUrl
// 本接口 ( DescribeInstanceVncUrl ) 用于查询实例管理终端地址，获取的地址可用于实例的 VNC 登录。
//
// 
//
// * 处于 `STOPPED` 状态的机器无法使用此功能。
//
// * 管理终端地址的有效期为 15 秒，调用接口成功后如果 15 秒内不使用该链接进行访问，管理终端地址自动失效，您需要重新查询。
//
// * 管理终端地址一旦被访问，将自动失效，您需要重新查询。
//
// * 如果连接断开，每分钟内重新连接的次数不能超过 30 次。
//
// * 参数 `InstanceVncUrl` ：调用接口成功后会返回的 `InstanceVncUrl` 的值。
//
// 获取到 `InstanceVncUrl` 后，您需要在链接 `https://img.qcloud.com/qcloud/app/active_vnc/index.html?` 末尾加上参数 `InstanceVncUrl=xxxx`。
//
//  最后组成的 URL 格式如下：
//
// 
//
// ```
//
// https://img.qcloud.com/qcloud/app/active_vnc/index.html?InstanceVncUrl=wss%3A%2F%2Fbjvnc.qcloud.com%3A26789%2Fvnc%3Fs%3DaHpjWnRVMFNhYmxKdDM5MjRHNlVTSVQwajNUSW0wb2tBbmFtREFCTmFrcy8vUUNPMG0wSHZNOUUxRm5PMmUzWmFDcWlOdDJIbUJxSTZDL0RXcHZxYnZZMmRkWWZWcEZia2lyb09XMzdKNmM9
//
// ```
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DescribeInstanceVncUrlWithContext(ctx context.Context, request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceVncUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 本接口（DescribeInstances）用于查询一个或多个实例的详细信息。
//
// 
//
// * 可以根据实例 ID、实例名称或者实例的内网 IP 查询实例的详细信息。
//
// * 过滤信息详细请见过滤器 [Filters](https://cloud.tencent.com/document/product/1207/47576#Filter) 。
//
// * 如果参数为空，返回当前用户一定数量（Limit 所指定的数量，默认为 20）的实例。
//
// * 支持查询实例的最新操作（LatestOperation）以及最新操作状态（LatestOperationState）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 本接口（DescribeInstances）用于查询一个或多个实例的详细信息。
//
// 
//
// * 可以根据实例 ID、实例名称或者实例的内网 IP 查询实例的详细信息。
//
// * 过滤信息详细请见过滤器 [Filters](https://cloud.tencent.com/document/product/1207/47576#Filter) 。
//
// * 如果参数为空，返回当前用户一定数量（Limit 所指定的数量，默认为 20）的实例。
//
// * 支持查询实例的最新操作（LatestOperation）以及最新操作状态（LatestOperationState）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDeniedActionsRequest() (request *DescribeInstancesDeniedActionsRequest) {
    request = &DescribeInstancesDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesDeniedActions")
    
    
    return
}

func NewDescribeInstancesDeniedActionsResponse() (response *DescribeInstancesDeniedActionsResponse) {
    response = &DescribeInstancesDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDeniedActions
// 本接口（DescribeInstancesDeniedActions）用于查询一个或多个实例的操作限制列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesDeniedActions(request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    return c.DescribeInstancesDeniedActionsWithContext(context.Background(), request)
}

// DescribeInstancesDeniedActions
// 本接口（DescribeInstancesDeniedActions）用于查询一个或多个实例的操作限制列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesDeniedActionsWithContext(ctx context.Context, request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDiskNumRequest() (request *DescribeInstancesDiskNumRequest) {
    request = &DescribeInstancesDiskNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesDiskNum")
    
    
    return
}

func NewDescribeInstancesDiskNumResponse() (response *DescribeInstancesDiskNumResponse) {
    response = &DescribeInstancesDiskNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDiskNum
// 本接口(DescribeInstancesDiskNum)用于查询实例挂载云硬盘数量。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeInstancesDiskNum(request *DescribeInstancesDiskNumRequest) (response *DescribeInstancesDiskNumResponse, err error) {
    return c.DescribeInstancesDiskNumWithContext(context.Background(), request)
}

// DescribeInstancesDiskNum
// 本接口(DescribeInstancesDiskNum)用于查询实例挂载云硬盘数量。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewDescribeInstancesReturnableRequest() (request *DescribeInstancesReturnableRequest) {
    request = &DescribeInstancesReturnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesReturnable")
    
    
    return
}

func NewDescribeInstancesReturnableResponse() (response *DescribeInstancesReturnableResponse) {
    response = &DescribeInstancesReturnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesReturnable
// 本接口（DescribeInstancesReturnable）用于查询实例是否可退还。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_DESCRIBEINSTANCESRETURNABLEERROR = "FailedOperation.DescribeInstancesReturnableError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESRETURNABLEERROR = "InternalError.DescribeInstancesReturnableError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesReturnable(request *DescribeInstancesReturnableRequest) (response *DescribeInstancesReturnableResponse, err error) {
    return c.DescribeInstancesReturnableWithContext(context.Background(), request)
}

// DescribeInstancesReturnable
// 本接口（DescribeInstancesReturnable）用于查询实例是否可退还。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_DESCRIBEINSTANCESRETURNABLEERROR = "FailedOperation.DescribeInstancesReturnableError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESRETURNABLEERROR = "InternalError.DescribeInstancesReturnableError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesReturnableWithContext(ctx context.Context, request *DescribeInstancesReturnableRequest) (response *DescribeInstancesReturnableResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesReturnableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesReturnable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesReturnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesTrafficPackagesRequest() (request *DescribeInstancesTrafficPackagesRequest) {
    request = &DescribeInstancesTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesTrafficPackages")
    
    
    return
}

func NewDescribeInstancesTrafficPackagesResponse() (response *DescribeInstancesTrafficPackagesResponse) {
    response = &DescribeInstancesTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesTrafficPackages
// 本接口（DescribeInstancesTrafficPackages）用于查询一个或多个实例的流量包详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "FailedOperation.DescribeInstancesTrafficPackagesFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "InternalError.DescribeInstancesTrafficPackagesFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesTrafficPackages(request *DescribeInstancesTrafficPackagesRequest) (response *DescribeInstancesTrafficPackagesResponse, err error) {
    return c.DescribeInstancesTrafficPackagesWithContext(context.Background(), request)
}

// DescribeInstancesTrafficPackages
// 本接口（DescribeInstancesTrafficPackages）用于查询一个或多个实例的流量包详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "FailedOperation.DescribeInstancesTrafficPackagesFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "InternalError.DescribeInstancesTrafficPackagesFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesTrafficPackagesWithContext(ctx context.Context, request *DescribeInstancesTrafficPackagesRequest) (response *DescribeInstancesTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesTrafficPackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesTrafficPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
    request = &DescribeKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeKeyPairs")
    
    
    return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
    response = &DescribeKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKeyPairs
// 本接口 (DescribeKeyPairs) 用于查询用户密钥对信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    return c.DescribeKeyPairsWithContext(context.Background(), request)
}

// DescribeKeyPairs
// 本接口 (DescribeKeyPairs) 用于查询用户密钥对信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeKeyPairsWithContext(ctx context.Context, request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModifyInstanceBundlesRequest() (request *DescribeModifyInstanceBundlesRequest) {
    request = &DescribeModifyInstanceBundlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeModifyInstanceBundles")
    
    
    return
}

func NewDescribeModifyInstanceBundlesResponse() (response *DescribeModifyInstanceBundlesResponse) {
    response = &DescribeModifyInstanceBundlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModifyInstanceBundles
// 本接口（DescribeModifyInstanceBundles）用于查询实例可变更套餐列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESMODIFICATIONERROR = "FailedOperation.DescribeInstancesModificationError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeModifyInstanceBundles(request *DescribeModifyInstanceBundlesRequest) (response *DescribeModifyInstanceBundlesResponse, err error) {
    return c.DescribeModifyInstanceBundlesWithContext(context.Background(), request)
}

// DescribeModifyInstanceBundles
// 本接口（DescribeModifyInstanceBundles）用于查询实例可变更套餐列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESMODIFICATIONERROR = "FailedOperation.DescribeInstancesModificationError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeModifyInstanceBundlesWithContext(ctx context.Context, request *DescribeModifyInstanceBundlesRequest) (response *DescribeModifyInstanceBundlesResponse, err error) {
    if request == nil {
        request = NewDescribeModifyInstanceBundlesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModifyInstanceBundles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModifyInstanceBundlesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// 本接口（DescribeRegions）用于查询地域信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 本接口（DescribeRegions）用于查询地域信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResetInstanceBlueprintsRequest() (request *DescribeResetInstanceBlueprintsRequest) {
    request = &DescribeResetInstanceBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeResetInstanceBlueprints")
    
    
    return
}

func NewDescribeResetInstanceBlueprintsResponse() (response *DescribeResetInstanceBlueprintsResponse) {
    response = &DescribeResetInstanceBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResetInstanceBlueprints
// 本接口（DescribeResetInstanceBlueprints）查询重置实例的镜像信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeResetInstanceBlueprints(request *DescribeResetInstanceBlueprintsRequest) (response *DescribeResetInstanceBlueprintsResponse, err error) {
    return c.DescribeResetInstanceBlueprintsWithContext(context.Background(), request)
}

// DescribeResetInstanceBlueprints
// 本接口（DescribeResetInstanceBlueprints）查询重置实例的镜像信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeResetInstanceBlueprintsWithContext(ctx context.Context, request *DescribeResetInstanceBlueprintsRequest) (response *DescribeResetInstanceBlueprintsResponse, err error) {
    if request == nil {
        request = NewDescribeResetInstanceBlueprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResetInstanceBlueprints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResetInstanceBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScenesRequest() (request *DescribeScenesRequest) {
    request = &DescribeScenesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeScenes")
    
    
    return
}

func NewDescribeScenesResponse() (response *DescribeScenesResponse) {
    response = &DescribeScenesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScenes
// 本接口(DescribeScenes)用于查看使用场景列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScenes(request *DescribeScenesRequest) (response *DescribeScenesResponse, err error) {
    return c.DescribeScenesWithContext(context.Background(), request)
}

// DescribeScenes
// 本接口(DescribeScenes)用于查看使用场景列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeScenesWithContext(ctx context.Context, request *DescribeScenesRequest) (response *DescribeScenesResponse, err error) {
    if request == nil {
        request = NewDescribeScenesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScenes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScenesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeSnapshots")
    
    
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
// 可能返回的错误码:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

// DescribeSnapshots
// 本接口（DescribeSnapshots）用于查询快照的详细信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
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

func NewDescribeSnapshotsDeniedActionsRequest() (request *DescribeSnapshotsDeniedActionsRequest) {
    request = &DescribeSnapshotsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeSnapshotsDeniedActions")
    
    
    return
}

func NewDescribeSnapshotsDeniedActionsResponse() (response *DescribeSnapshotsDeniedActionsResponse) {
    response = &DescribeSnapshotsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotsDeniedActions
// 本接口（DescribeSnapshotsDeniedActions）用于查询一个或多个快照的操作限制列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeSnapshotsDeniedActions(request *DescribeSnapshotsDeniedActionsRequest) (response *DescribeSnapshotsDeniedActionsResponse, err error) {
    return c.DescribeSnapshotsDeniedActionsWithContext(context.Background(), request)
}

// DescribeSnapshotsDeniedActions
// 本接口（DescribeSnapshotsDeniedActions）用于查询一个或多个快照的操作限制列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeSnapshotsDeniedActionsWithContext(ctx context.Context, request *DescribeSnapshotsDeniedActionsRequest) (response *DescribeSnapshotsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// 查询地域下可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 查询地域下可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachCcnRequest() (request *DetachCcnRequest) {
    request = &DetachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DetachCcn")
    
    
    return
}

func NewDetachCcnResponse() (response *DetachCcnResponse) {
    response = &DetachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachCcn
// 本接口 (AttachCcn) 用于解除与云联网的关联。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_DETACHCCNFAILED = "UnsupportedOperation.DetachCcnFailed"
func (c *Client) DetachCcn(request *DetachCcnRequest) (response *DetachCcnResponse, err error) {
    return c.DetachCcnWithContext(context.Background(), request)
}

// DetachCcn
// 本接口 (AttachCcn) 用于解除与云联网的关联。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_DETACHCCNFAILED = "UnsupportedOperation.DetachCcnFailed"
func (c *Client) DetachCcnWithContext(ctx context.Context, request *DetachCcnRequest) (response *DetachCcnResponse, err error) {
    if request == nil {
        request = NewDetachCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachCcn require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
    request = &DetachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DetachDisks")
    
    
    return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
    response = &DetachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachDisks
// 本接口（DetachDisks）用于卸载一个或多个云硬盘。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    return c.DetachDisksWithContext(context.Background(), request)
}

// DetachDisks
// 本接口（DetachDisks）用于卸载一个或多个云硬盘。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
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

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
    request = &DisassociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "DisassociateInstancesKeyPairs")
    
    
    return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
    response = &DisassociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateInstancesKeyPairs
// 本接口（DisassociateInstancesKeyPairs）用于解除实例与指定密钥对的绑定关系。
//
// 
//
// * 只支持 [RUNNING, STOPPED] 状态的 LINUX_UNIX 操作系统的实例。处于 RUNNING 状态的实例会强制关机，然后解绑。
//
// * 解绑密钥后，实例可以通过原来设置的密码登录。
//
// * 如果原来没有设置密码，解绑后将无法使用 SSH 登录。可以调用 ResetInstancesPassword 接口来设置登录密码。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRNOTBOUNDTOINSTANCE = "UnsupportedOperation.KeyPairNotBoundToInstance"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    return c.DisassociateInstancesKeyPairsWithContext(context.Background(), request)
}

// DisassociateInstancesKeyPairs
// 本接口（DisassociateInstancesKeyPairs）用于解除实例与指定密钥对的绑定关系。
//
// 
//
// * 只支持 [RUNNING, STOPPED] 状态的 LINUX_UNIX 操作系统的实例。处于 RUNNING 状态的实例会强制关机，然后解绑。
//
// * 解绑密钥后，实例可以通过原来设置的密码登录。
//
// * 如果原来没有设置密码，解绑后将无法使用 SSH 登录。可以调用 ResetInstancesPassword 接口来设置登录密码。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRNOTBOUNDTOINSTANCE = "UnsupportedOperation.KeyPairNotBoundToInstance"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DisassociateInstancesKeyPairsWithContext(ctx context.Context, request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateInstancesKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewImportKeyPairRequest() (request *ImportKeyPairRequest) {
    request = &ImportKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ImportKeyPair")
    
    
    return
}

func NewImportKeyPairResponse() (response *ImportKeyPairResponse) {
    response = &ImportKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportKeyPair
// 本接口（ImportKeyPair）用于导入用户指定密钥对。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  FAILEDOPERATION_IMPORTKEYPAIRFAILED = "FailedOperation.ImportKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYDUPLICATED = "InvalidParameterValue.KeyPairPublicKeyDuplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYMALFORMED = "InvalidParameterValue.KeyPairPublicKeyMalformed"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    return c.ImportKeyPairWithContext(context.Background(), request)
}

// ImportKeyPair
// 本接口（ImportKeyPair）用于导入用户指定密钥对。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  FAILEDOPERATION_IMPORTKEYPAIRFAILED = "FailedOperation.ImportKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYDUPLICATED = "InvalidParameterValue.KeyPairPublicKeyDuplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYMALFORMED = "InvalidParameterValue.KeyPairPublicKeyMalformed"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ImportKeyPairWithContext(ctx context.Context, request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    if request == nil {
        request = NewImportKeyPairRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportKeyPair require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateBlueprintRequest() (request *InquirePriceCreateBlueprintRequest) {
    request = &InquirePriceCreateBlueprintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateBlueprint")
    
    
    return
}

func NewInquirePriceCreateBlueprintResponse() (response *InquirePriceCreateBlueprintResponse) {
    response = &InquirePriceCreateBlueprintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateBlueprint
// 本接口 (InquirePriceCreateBlueprint) 用于创建镜像询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateBlueprint(request *InquirePriceCreateBlueprintRequest) (response *InquirePriceCreateBlueprintResponse, err error) {
    return c.InquirePriceCreateBlueprintWithContext(context.Background(), request)
}

// InquirePriceCreateBlueprint
// 本接口 (InquirePriceCreateBlueprint) 用于创建镜像询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateBlueprintWithContext(ctx context.Context, request *InquirePriceCreateBlueprintRequest) (response *InquirePriceCreateBlueprintResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateBlueprintRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateBlueprint require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateBlueprintResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateDisksRequest() (request *InquirePriceCreateDisksRequest) {
    request = &InquirePriceCreateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateDisks")
    
    
    return
}

func NewInquirePriceCreateDisksResponse() (response *InquirePriceCreateDisksResponse) {
    response = &InquirePriceCreateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateDisks
// 本接口（InquirePriceCreateDisks）用于新购云硬盘询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
func (c *Client) InquirePriceCreateDisks(request *InquirePriceCreateDisksRequest) (response *InquirePriceCreateDisksResponse, err error) {
    return c.InquirePriceCreateDisksWithContext(context.Background(), request)
}

// InquirePriceCreateDisks
// 本接口（InquirePriceCreateDisks）用于新购云硬盘询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
func (c *Client) InquirePriceCreateDisksWithContext(ctx context.Context, request *InquirePriceCreateDisksRequest) (response *InquirePriceCreateDisksResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateInstancesRequest() (request *InquirePriceCreateInstancesRequest) {
    request = &InquirePriceCreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateInstances")
    
    
    return
}

func NewInquirePriceCreateInstancesResponse() (response *InquirePriceCreateInstancesResponse) {
    response = &InquirePriceCreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateInstances
// 本接口（InquiryPriceCreateInstances）用于创建实例询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateInstances(request *InquirePriceCreateInstancesRequest) (response *InquirePriceCreateInstancesResponse, err error) {
    return c.InquirePriceCreateInstancesWithContext(context.Background(), request)
}

// InquirePriceCreateInstances
// 本接口（InquiryPriceCreateInstances）用于创建实例询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateInstancesWithContext(ctx context.Context, request *InquirePriceCreateInstancesRequest) (response *InquirePriceCreateInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewDisksRequest() (request *InquirePriceRenewDisksRequest) {
    request = &InquirePriceRenewDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceRenewDisks")
    
    
    return
}

func NewInquirePriceRenewDisksResponse() (response *InquirePriceRenewDisksResponse) {
    response = &InquirePriceRenewDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceRenewDisks
// 本接口（InquirePriceRenewDisks）用于续费云硬盘询价。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
func (c *Client) InquirePriceRenewDisks(request *InquirePriceRenewDisksRequest) (response *InquirePriceRenewDisksResponse, err error) {
    return c.InquirePriceRenewDisksWithContext(context.Background(), request)
}

// InquirePriceRenewDisks
// 本接口（InquirePriceRenewDisks）用于续费云硬盘询价。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
func (c *Client) InquirePriceRenewDisksWithContext(ctx context.Context, request *InquirePriceRenewDisksRequest) (response *InquirePriceRenewDisksResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewInstancesRequest() (request *InquirePriceRenewInstancesRequest) {
    request = &InquirePriceRenewInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceRenewInstances")
    
    
    return
}

func NewInquirePriceRenewInstancesResponse() (response *InquirePriceRenewInstancesResponse) {
    response = &InquirePriceRenewInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceRenewInstances
// 本接口（InquirePriceRenewInstances）用于续费实例询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEDATADISKNOTFOUND = "ResourceNotFound.InstanceDataDiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceRenewInstances(request *InquirePriceRenewInstancesRequest) (response *InquirePriceRenewInstancesResponse, err error) {
    return c.InquirePriceRenewInstancesWithContext(context.Background(), request)
}

// InquirePriceRenewInstances
// 本接口（InquirePriceRenewInstances）用于续费实例询价。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEDATADISKNOTFOUND = "ResourceNotFound.InstanceDataDiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceRenewInstancesWithContext(ctx context.Context, request *InquirePriceRenewInstancesRequest) (response *InquirePriceRenewInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateInstancesRequest() (request *IsolateInstancesRequest) {
    request = &IsolateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "IsolateInstances")
    
    
    return
}

func NewIsolateInstancesResponse() (response *IsolateInstancesResponse) {
    response = &IsolateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateInstances
// 本接口(IsolateInstances)用于退还一个或多个轻量应用服务器实例。
//
// * 只有状态为 RUNNING 或 STOPPED 的实例才可以进行此操作。
//
// * 接口调用成功后，实例会进入SHUTDOWN 状态。
//
// * 支持批量操作。每次请求批量资源（包括实例与数据盘）的上限为 20。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ISOLATERESOURCESFAILED = "FailedOperation.IsolateResourcesFailed"
//  INTERNALERROR_DESCRIBERESOURCESRETURNABLEERROR = "InternalError.DescribeResourcesReturnableError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_ISOLATERESOURCESLIMITEXCEEDED = "LimitExceeded.IsolateResourcesLimitExceeded"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_RESOURCENOTRETURNABLE = "UnsupportedOperation.ResourceNotReturnable"
func (c *Client) IsolateInstances(request *IsolateInstancesRequest) (response *IsolateInstancesResponse, err error) {
    return c.IsolateInstancesWithContext(context.Background(), request)
}

// IsolateInstances
// 本接口(IsolateInstances)用于退还一个或多个轻量应用服务器实例。
//
// * 只有状态为 RUNNING 或 STOPPED 的实例才可以进行此操作。
//
// * 接口调用成功后，实例会进入SHUTDOWN 状态。
//
// * 支持批量操作。每次请求批量资源（包括实例与数据盘）的上限为 20。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ISOLATERESOURCESFAILED = "FailedOperation.IsolateResourcesFailed"
//  INTERNALERROR_DESCRIBERESOURCESRETURNABLEERROR = "InternalError.DescribeResourcesReturnableError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_ISOLATERESOURCESLIMITEXCEEDED = "LimitExceeded.IsolateResourcesLimitExceeded"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_RESOURCENOTRETURNABLE = "UnsupportedOperation.ResourceNotReturnable"
func (c *Client) IsolateInstancesWithContext(ctx context.Context, request *IsolateInstancesRequest) (response *IsolateInstancesResponse, err error) {
    if request == nil {
        request = NewIsolateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlueprintAttributeRequest() (request *ModifyBlueprintAttributeRequest) {
    request = &ModifyBlueprintAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyBlueprintAttribute")
    
    
    return
}

func NewModifyBlueprintAttributeResponse() (response *ModifyBlueprintAttributeResponse) {
    response = &ModifyBlueprintAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBlueprintAttribute
// 本接口 (ModifyBlueprintAttribute) 用于修改镜像属性。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_PRIVATEBLUEPRINTNOTFOUND = "ResourceNotFound.PrivateBlueprintNotFound"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
func (c *Client) ModifyBlueprintAttribute(request *ModifyBlueprintAttributeRequest) (response *ModifyBlueprintAttributeResponse, err error) {
    return c.ModifyBlueprintAttributeWithContext(context.Background(), request)
}

// ModifyBlueprintAttribute
// 本接口 (ModifyBlueprintAttribute) 用于修改镜像属性。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_PRIVATEBLUEPRINTNOTFOUND = "ResourceNotFound.PrivateBlueprintNotFound"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
func (c *Client) ModifyBlueprintAttributeWithContext(ctx context.Context, request *ModifyBlueprintAttributeRequest) (response *ModifyBlueprintAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBlueprintAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlueprintAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlueprintAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiskBackupsAttributeRequest() (request *ModifyDiskBackupsAttributeRequest) {
    request = &ModifyDiskBackupsAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyDiskBackupsAttribute")
    
    
    return
}

func NewModifyDiskBackupsAttributeResponse() (response *ModifyDiskBackupsAttributeResponse) {
    response = &ModifyDiskBackupsAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDiskBackupsAttribute
// 本接口 (ModifyDiskBackupsAttribute) 用于修改云硬盘备份点属性。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  RESOURCENOTFOUND_DISKBACKUPIDNOTFOUND = "ResourceNotFound.DiskBackupIdNotFound"
//  RESOURCENOTFOUND_DISKBACKUPNOTFOUND = "ResourceNotFound.DiskBackupNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyDiskBackupsAttribute(request *ModifyDiskBackupsAttributeRequest) (response *ModifyDiskBackupsAttributeResponse, err error) {
    return c.ModifyDiskBackupsAttributeWithContext(context.Background(), request)
}

// ModifyDiskBackupsAttribute
// 本接口 (ModifyDiskBackupsAttribute) 用于修改云硬盘备份点属性。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"
//  RESOURCENOTFOUND_DISKBACKUPIDNOTFOUND = "ResourceNotFound.DiskBackupIdNotFound"
//  RESOURCENOTFOUND_DISKBACKUPNOTFOUND = "ResourceNotFound.DiskBackupNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyDiskBackupsAttributeWithContext(ctx context.Context, request *ModifyDiskBackupsAttributeRequest) (response *ModifyDiskBackupsAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDiskBackupsAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiskBackupsAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiskBackupsAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisksAttributeRequest() (request *ModifyDisksAttributeRequest) {
    request = &ModifyDisksAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyDisksAttribute")
    
    
    return
}

func NewModifyDisksAttributeResponse() (response *ModifyDisksAttributeResponse) {
    response = &ModifyDisksAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDisksAttribute
// 本接口(ModifyDisksAttribute)用于修改云硬盘属性。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
func (c *Client) ModifyDisksAttribute(request *ModifyDisksAttributeRequest) (response *ModifyDisksAttributeResponse, err error) {
    return c.ModifyDisksAttributeWithContext(context.Background(), request)
}

// ModifyDisksAttribute
// 本接口(ModifyDisksAttribute)用于修改云硬盘属性。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
func (c *Client) ModifyDisksAttributeWithContext(ctx context.Context, request *ModifyDisksAttributeRequest) (response *ModifyDisksAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDisksAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDisksAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDisksAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisksRenewFlagRequest() (request *ModifyDisksRenewFlagRequest) {
    request = &ModifyDisksRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyDisksRenewFlag")
    
    
    return
}

func NewModifyDisksRenewFlagResponse() (response *ModifyDisksRenewFlagResponse) {
    response = &ModifyDisksRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDisksRenewFlag
// 本接口（ModifyDisksRenewFlag）用于修改云硬盘续费标识。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
func (c *Client) ModifyDisksRenewFlag(request *ModifyDisksRenewFlagRequest) (response *ModifyDisksRenewFlagResponse, err error) {
    return c.ModifyDisksRenewFlagWithContext(context.Background(), request)
}

// ModifyDisksRenewFlag
// 本接口（ModifyDisksRenewFlag）用于修改云硬盘续费标识。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
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

func NewModifyFirewallRuleDescriptionRequest() (request *ModifyFirewallRuleDescriptionRequest) {
    request = &ModifyFirewallRuleDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyFirewallRuleDescription")
    
    
    return
}

func NewModifyFirewallRuleDescriptionResponse() (response *ModifyFirewallRuleDescriptionResponse) {
    response = &ModifyFirewallRuleDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFirewallRuleDescription
// 本接口（ModifyFirewallRuleDescription）用于修改单条防火墙规则描述。
//
// 
//
// * FirewallVersion 用于指定要操作的防火墙的版本。传入 FirewallVersion 版本号若不等于当前防火墙的最新版本，将返回失败；若不传 FirewallVersion 则直接修改防火墙规则备注。
//
// 
//
// 在 FirewallRule 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRuleDescription(request *ModifyFirewallRuleDescriptionRequest) (response *ModifyFirewallRuleDescriptionResponse, err error) {
    return c.ModifyFirewallRuleDescriptionWithContext(context.Background(), request)
}

// ModifyFirewallRuleDescription
// 本接口（ModifyFirewallRuleDescription）用于修改单条防火墙规则描述。
//
// 
//
// * FirewallVersion 用于指定要操作的防火墙的版本。传入 FirewallVersion 版本号若不等于当前防火墙的最新版本，将返回失败；若不传 FirewallVersion 则直接修改防火墙规则备注。
//
// 
//
// 在 FirewallRule 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRuleDescriptionWithContext(ctx context.Context, request *ModifyFirewallRuleDescriptionRequest) (response *ModifyFirewallRuleDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyFirewallRuleDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFirewallRuleDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFirewallRuleDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFirewallRulesRequest() (request *ModifyFirewallRulesRequest) {
    request = &ModifyFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyFirewallRules")
    
    
    return
}

func NewModifyFirewallRulesResponse() (response *ModifyFirewallRulesResponse) {
    response = &ModifyFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFirewallRules
// 本接口（ModifyFirewallRules）用于重置实例防火墙规则。
//
// 
//
// 本接口先删除当前实例的所有防火墙规则，然后添加新的规则。
//
// 
//
// * FirewallVersion 用于指定要操作的防火墙的版本。传入 FirewallVersion 版本号若不等于当前防火墙的最新版本，将返回失败；若不传 FirewallVersion 则直接重置防火墙规则。
//
// 
//
// 在 FirewallRules 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRules(request *ModifyFirewallRulesRequest) (response *ModifyFirewallRulesResponse, err error) {
    return c.ModifyFirewallRulesWithContext(context.Background(), request)
}

// ModifyFirewallRules
// 本接口（ModifyFirewallRules）用于重置实例防火墙规则。
//
// 
//
// 本接口先删除当前实例的所有防火墙规则，然后添加新的规则。
//
// 
//
// * FirewallVersion 用于指定要操作的防火墙的版本。传入 FirewallVersion 版本号若不等于当前防火墙的最新版本，将返回失败；若不传 FirewallVersion 则直接重置防火墙规则。
//
// 
//
// 在 FirewallRules 参数中：
//
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
//
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
//
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
//
// * Action 字段只允许输入 ACCEPT 或 DROP。
//
// * FirewallRuleDescription 字段长度不得超过 64。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRulesWithContext(ctx context.Context, request *ModifyFirewallRulesRequest) (response *ModifyFirewallRulesResponse, err error) {
    if request == nil {
        request = NewModifyFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesAttribute")
    
    
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesAttribute
// 本接口（ModifyInstancesAttribute）用于修改实例的属性。
//
// * “实例名称”仅为方便用户自己管理之用。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER_ONLYALLOWMODIFYONEATTRIBUTE = "InvalidParameter.OnlyAllowModifyOneAttribute"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    return c.ModifyInstancesAttributeWithContext(context.Background(), request)
}

// ModifyInstancesAttribute
// 本接口（ModifyInstancesAttribute）用于修改实例的属性。
//
// * “实例名称”仅为方便用户自己管理之用。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER_ONLYALLOWMODIFYONEATTRIBUTE = "InvalidParameter.OnlyAllowModifyOneAttribute"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesAttributeWithContext(ctx context.Context, request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesBundleRequest() (request *ModifyInstancesBundleRequest) {
    request = &ModifyInstancesBundleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesBundle")
    
    
    return
}

func NewModifyInstancesBundleResponse() (response *ModifyInstancesBundleResponse) {
    response = &ModifyInstancesBundleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesBundle
// 本接口(ModifyInstancesBundle)用于变更一个或多个轻量应用服务器实例套餐。
//
// * 只有状态为 RUNNING，STOPPED的实例才可以进行此操作。
//
// * 支持批量操作。每次请求批量实例的上限为 30。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  FAILEDOPERATION_MODIFYINSTANCESBUNDLEFAILED = "FailedOperation.ModifyInstancesBundleFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) ModifyInstancesBundle(request *ModifyInstancesBundleRequest) (response *ModifyInstancesBundleResponse, err error) {
    return c.ModifyInstancesBundleWithContext(context.Background(), request)
}

// ModifyInstancesBundle
// 本接口(ModifyInstancesBundle)用于变更一个或多个轻量应用服务器实例套餐。
//
// * 只有状态为 RUNNING，STOPPED的实例才可以进行此操作。
//
// * 支持批量操作。每次请求批量实例的上限为 30。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  FAILEDOPERATION_MODIFYINSTANCESBUNDLEFAILED = "FailedOperation.ModifyInstancesBundleFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) ModifyInstancesBundleWithContext(ctx context.Context, request *ModifyInstancesBundleRequest) (response *ModifyInstancesBundleResponse, err error) {
    if request == nil {
        request = NewModifyInstancesBundleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesBundle require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesBundleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesLoginKeyPairAttributeRequest() (request *ModifyInstancesLoginKeyPairAttributeRequest) {
    request = &ModifyInstancesLoginKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesLoginKeyPairAttribute")
    
    
    return
}

func NewModifyInstancesLoginKeyPairAttributeResponse() (response *ModifyInstancesLoginKeyPairAttributeResponse) {
    response = &ModifyInstancesLoginKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesLoginKeyPairAttribute
// 本接口用于设置实例默认登录密钥对属性。
//
// 
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCELOGINKEYPAIRPERMITLOGIN = "InvalidParameterValue.InvalidInstanceLoginKeyPairPermitLogin"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesLoginKeyPairAttribute(request *ModifyInstancesLoginKeyPairAttributeRequest) (response *ModifyInstancesLoginKeyPairAttributeResponse, err error) {
    return c.ModifyInstancesLoginKeyPairAttributeWithContext(context.Background(), request)
}

// ModifyInstancesLoginKeyPairAttribute
// 本接口用于设置实例默认登录密钥对属性。
//
// 
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCELOGINKEYPAIRPERMITLOGIN = "InvalidParameterValue.InvalidInstanceLoginKeyPairPermitLogin"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesLoginKeyPairAttributeWithContext(ctx context.Context, request *ModifyInstancesLoginKeyPairAttributeRequest) (response *ModifyInstancesLoginKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesLoginKeyPairAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesLoginKeyPairAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesLoginKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesRenewFlagRequest() (request *ModifyInstancesRenewFlagRequest) {
    request = &ModifyInstancesRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesRenewFlag")
    
    
    return
}

func NewModifyInstancesRenewFlagResponse() (response *ModifyInstancesRenewFlagResponse) {
    response = &ModifyInstancesRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesRenewFlag
// 本接口 (ModifyInstancesRenewFlag) 用于修改包年包月实例续费标识。
//
// 
//
// * 实例被标识为自动续费后，每次在实例到期时，会自动续费一个月。
//
// * 支持批量操作。每次请求批量实例的上限为100。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesRenewFlag(request *ModifyInstancesRenewFlagRequest) (response *ModifyInstancesRenewFlagResponse, err error) {
    return c.ModifyInstancesRenewFlagWithContext(context.Background(), request)
}

// ModifyInstancesRenewFlag
// 本接口 (ModifyInstancesRenewFlag) 用于修改包年包月实例续费标识。
//
// 
//
// * 实例被标识为自动续费后，每次在实例到期时，会自动续费一个月。
//
// * 支持批量操作。每次请求批量实例的上限为100。
//
// 可能返回的错误码:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesRenewFlagWithContext(ctx context.Context, request *ModifyInstancesRenewFlagRequest) (response *ModifyInstancesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyInstancesRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
    request = &ModifySnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifySnapshotAttribute")
    
    
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
// <li>“快照名称”仅为方便用户自己管理之用。</li>
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    return c.ModifySnapshotAttributeWithContext(context.Background(), request)
}

// ModifySnapshotAttribute
// 本接口（ModifySnapshotAttribute）用于修改指定快照的属性。
//
// <li>“快照名称”仅为方便用户自己管理之用。</li>
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
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

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "RebootInstances")
    
    
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RebootInstances
// 本接口（RebootInstances）用于重启实例。
//
// 
//
// * 只有状态为 RUNNING 的实例才可以进行此操作。
//
// * 接口调用成功时，实例会进入 REBOOTING 状态；重启实例成功时，实例会进入 RUNNING 状态。
//
// * 支持批量操作，每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    return c.RebootInstancesWithContext(context.Background(), request)
}

// RebootInstances
// 本接口（RebootInstances）用于重启实例。
//
// 
//
// * 只有状态为 RUNNING 的实例才可以进行此操作。
//
// * 接口调用成功时，实例会进入 REBOOTING 状态；重启实例成功时，实例会进入 RUNNING 状态。
//
// * 支持批量操作，每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstancesRequest() (request *RenewInstancesRequest) {
    request = &RenewInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "RenewInstances")
    
    
    return
}

func NewRenewInstancesResponse() (response *RenewInstancesResponse) {
    response = &RenewInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewInstances
// 本接口(RenewInstances)用于续费一个或多个轻量应用服务器实例。
//
// * 只有状态为 RUNNING，STOPPED 或 SHUTDOWN 的实例才可以进行此操作。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RENEWINSTANCESFAILED = "FailedOperation.RenewInstancesFailed"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
func (c *Client) RenewInstances(request *RenewInstancesRequest) (response *RenewInstancesResponse, err error) {
    return c.RenewInstancesWithContext(context.Background(), request)
}

// RenewInstances
// 本接口(RenewInstances)用于续费一个或多个轻量应用服务器实例。
//
// * 只有状态为 RUNNING，STOPPED 或 SHUTDOWN 的实例才可以进行此操作。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RENEWINSTANCESFAILED = "FailedOperation.RenewInstancesFailed"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
func (c *Client) RenewInstancesWithContext(ctx context.Context, request *RenewInstancesRequest) (response *RenewInstancesResponse, err error) {
    if request == nil {
        request = NewRenewInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetAttachCcnRequest() (request *ResetAttachCcnRequest) {
    request = &ResetAttachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetAttachCcn")
    
    
    return
}

func NewResetAttachCcnResponse() (response *ResetAttachCcnResponse) {
    response = &ResetAttachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAttachCcn
// 本接口 (ResetAttachCcn) 用于关联云联网实例申请过期时，重新申请关联操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_RESETATTACHCCNFAILED = "UnsupportedOperation.ResetAttachCcnFailed"
func (c *Client) ResetAttachCcn(request *ResetAttachCcnRequest) (response *ResetAttachCcnResponse, err error) {
    return c.ResetAttachCcnWithContext(context.Background(), request)
}

// ResetAttachCcn
// 本接口 (ResetAttachCcn) 用于关联云联网实例申请过期时，重新申请关联操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_RESETATTACHCCNFAILED = "UnsupportedOperation.ResetAttachCcnFailed"
func (c *Client) ResetAttachCcnWithContext(ctx context.Context, request *ResetAttachCcnRequest) (response *ResetAttachCcnResponse, err error) {
    if request == nil {
        request = NewResetAttachCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAttachCcn require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAttachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
    request = &ResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetInstance")
    
    
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstance
// 本接口（ResetInstance）用于重装指定实例上的镜像。
//
// 
//
// * 如果指定了 BlueprintId 参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装。
//
// * 系统盘将会被格式化，并重置；请确保系统盘中无重要文件。
//
// * 目前不支持实例使用该接口实现 LINUX_UNIX 和 WINDOWS 操作系统切换。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTCONFIGNOTMATCH = "InvalidParameterValue.BlueprintConfigNotMatch"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_BLUEPRINTUNAVAILABLE = "ResourceUnavailable.BlueprintUnavailable"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    return c.ResetInstanceWithContext(context.Background(), request)
}

// ResetInstance
// 本接口（ResetInstance）用于重装指定实例上的镜像。
//
// 
//
// * 如果指定了 BlueprintId 参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装。
//
// * 系统盘将会被格式化，并重置；请确保系统盘中无重要文件。
//
// * 目前不支持实例使用该接口实现 LINUX_UNIX 和 WINDOWS 操作系统切换。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTCONFIGNOTMATCH = "InvalidParameterValue.BlueprintConfigNotMatch"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_BLUEPRINTUNAVAILABLE = "ResourceUnavailable.BlueprintUnavailable"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ResetInstanceWithContext(ctx context.Context, request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetInstancesPassword")
    
    
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesPassword
// 本接口（ResetInstancesPassword）用于将实例操作系统的密码重置为用户指定的密码。
//
// * 只修改管理员帐号的密码。实例的操作系统不同，管理员帐号也会不一样（Windows 为 Administrator，Ubuntu 为 ubuntu ，其它系统为 root）。
//
// * 支持批量操作。将多个实例操作系统的密码重置为相同的密码。每次请求批量实例的上限为 100。
//
// * 建议对运行中的实例先手动关机，然后再进行密码重置。如实例处于运行中状态，本接口操作过程中会对实例进行关机操作，尝试正常关机失败后进行强制关机。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 注意：强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    return c.ResetInstancesPasswordWithContext(context.Background(), request)
}

// ResetInstancesPassword
// 本接口（ResetInstancesPassword）用于将实例操作系统的密码重置为用户指定的密码。
//
// * 只修改管理员帐号的密码。实例的操作系统不同，管理员帐号也会不一样（Windows 为 Administrator，Ubuntu 为 ubuntu ，其它系统为 root）。
//
// * 支持批量操作。将多个实例操作系统的密码重置为相同的密码。每次请求批量实例的上限为 100。
//
// * 建议对运行中的实例先手动关机，然后再进行密码重置。如实例处于运行中状态，本接口操作过程中会对实例进行关机操作，尝试正常关机失败后进行强制关机。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 注意：强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ResetInstancesPasswordWithContext(ctx context.Context, request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstancesPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "StartInstances")
    
    
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartInstances
// 本接口（StartInstances）用于启动一个或多个实例。
//
// 
//
// * 只有状态为 STOPPED 的实例才可以进行此操作。
//
// * 接口调用成功时，实例会进入 STARTING 状态；启动实例成功时，实例会进入 RUNNING 状态。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    return c.StartInstancesWithContext(context.Background(), request)
}

// StartInstances
// 本接口（StartInstances）用于启动一个或多个实例。
//
// 
//
// * 只有状态为 STOPPED 的实例才可以进行此操作。
//
// * 接口调用成功时，实例会进入 STARTING 状态；启动实例成功时，实例会进入 RUNNING 状态。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "StopInstances")
    
    
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopInstances
// 本接口（StopInstances）用于关闭一个或多个实例。
//
// * 只有状态为 RUNNING 的实例才可以进行此操作。
//
// * 接口调用成功时，实例会进入 STOPPING 状态；关闭实例成功时，实例会进入 STOPPED 状态。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    return c.StopInstancesWithContext(context.Background(), request)
}

// StopInstances
// 本接口（StopInstances）用于关闭一个或多个实例。
//
// * 只有状态为 RUNNING 的实例才可以进行此操作。
//
// * 接口调用成功时，实例会进入 STOPPING 状态；关闭实例成功时，实例会进入 STOPPED 状态。
//
// * 支持批量操作。每次请求批量实例的上限为 100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDisksRequest() (request *TerminateDisksRequest) {
    request = &TerminateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "TerminateDisks")
    
    
    return
}

func NewTerminateDisksResponse() (response *TerminateDisksResponse) {
    response = &TerminateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateDisks
// 本接口（TerminateDisks）用于销毁一个或多个云硬盘。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
func (c *Client) TerminateDisks(request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    return c.TerminateDisksWithContext(context.Background(), request)
}

// TerminateDisks
// 本接口（TerminateDisks）用于销毁一个或多个云硬盘。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
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

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lighthouse", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateInstances
// 本接口 (TerminateInstances) 用于销毁实例。
//
// 
//
// * 处于 SHUTDOWN 状态的实例，可通过本接口销毁，且不可恢复。
//
// * 支持批量操作，每次请求批量实例的上限为100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态 (LatestOperationState) 为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// 本接口 (TerminateInstances) 用于销毁实例。
//
// 
//
// * 处于 SHUTDOWN 状态的实例，可通过本接口销毁，且不可恢复。
//
// * 支持批量操作，每次请求批量实例的上限为100。
//
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态 (LatestOperationState) 为“SUCCESS”，则代表操作成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
