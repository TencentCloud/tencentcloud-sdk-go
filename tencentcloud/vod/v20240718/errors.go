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

package v20240718

const (
	// 此产品的特有错误码

	// 应用 ID 与存储桶 ID 不匹配。
	FAILEDOPERATION_BUCKETIDNOTMATCH = "FailedOperation.BucketIdNotMatch"

	// 存储桶增量迁移策略正在部署中。
	FAILEDOPERATION_BUCKETINCREMENTALMIGRATIONSTRATEGYDEPLOYING = "FailedOperation.BucketIncrementalMigrationStrategyDeploying"

	// 存储桶增量迁移策略超限。
	FAILEDOPERATION_INCREMENTALMIGRATIONSTRATEGYOVERLIMIT = "FailedOperation.IncrementalMigrationStrategyOverLimit"

	// 存储数量超过限制。
	FAILEDOPERATION_STORAGEREGIONBUCKETOVERLIMIT = "FailedOperation.StorageRegionBucketOverLimit"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 无效的过滤器
	INVALIDFILTER = "InvalidFilter"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 不支持的存储区域。
	INVALIDPARAMETERVALUE_UNSUPPORTEDSTORAGEREGION = "InvalidParameterValue.UnsupportedStorageRegion"

	// 增量迁移策略不存在。
	RESOURCENOTFOUND_INCREMENTALMIGRATIONSTRATEGYNOTFOUND = "ResourceNotFound.IncrementalMigrationStrategyNotFound"

	// 应用未授权。
	UNAUTHORIZEDOPERATION_SUBAPP = "UnauthorizedOperation.SubApp"
)
