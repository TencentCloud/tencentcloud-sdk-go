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

package v20190718

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 文件系统非空。
	FAILEDOPERATION_FILESYSTEMNOTEMPTY = "FailedOperation.FileSystemNotEmpty"

	// 修改的文件系统容量小于当前使用量。
	FAILEDOPERATION_QUOTALESSTHANCURRENTUSED = "FailedOperation.QuotaLessThanCurrentUsed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// AccessGroupId参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"

	// AccessGroupName参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"

	// 权限规则的Address参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"

	// CapacityQuota参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"

	// Description参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"

	// FileSystemId参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"

	// FileSystemName参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"

	// MountPointId参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"

	// MountPointName参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"

	// VpcId参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 权限组不存在。
	RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"

	// 权限规则不存在。
	RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"

	// 文件系统不存在。
	RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"

	// 挂载点不存在。
	RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"

	// VPC网络不存在。
	RESOURCENOTFOUND_VPCNOTEXISTS = "ResourceNotFound.VpcNotExists"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
