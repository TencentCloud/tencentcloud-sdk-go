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

package v20181011

const (
	// 此产品的特有错误码

	// 生成Token失败。
	FAILEDOPERATION_GENERATETOKENERROR = "FailedOperation.GenerateTokenError"

	// 奖品库存不足。
	FAILEDOPERATION_INSUFFICIENTPRIZESTOCK = "FailedOperation.InsufficientPrizeStock"

	// 内部错误
	INTERNALERROR = "InternalError"

	// 操作失败，请重试。
	INTERNALERROR_THIRDSERVERERROR = "InternalError.ThirdServerError"

	// 渠道不存在。
	RESOURCENOTFOUND_CHANNEL = "ResourceNotFound.Channel"

	// 工单不存在。
	RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"

	// 项目不存在。
	RESOURCENOTFOUND_PROJECT = "ResourceNotFound.Project"

	// 仅编辑中和已下线的项目允许删除!。
	UNSUPPORTEDOPERATION_PROJECTNOTALLOWEDTODELETE = "UnsupportedOperation.ProjectNotAllowedToDelete"

	// 仅运营中项目允许下线。
	UNSUPPORTEDOPERATION_STATUSOFFLINEPROJECT = "UnsupportedOperation.StatusOffLineProject"
)
