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

package v20181217

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 数据上报失败。
	FAILEDOPERATION_SENDTOCKAFKA = "FailedOperation.SendToCkafka"

	// 鉴权失败。
	INTERNALERROR_CAUTHERROR = "InternalError.CauthError"

	// 插入db失败。
	INTERNALERROR_DBERROR = "InternalError.DbError"

	// 已经审批。
	INVALIDPARAMETER_HASBEENAPPROVED = "InvalidParameter.HasBeenApproved"

	// id不存在。
	INVALIDPARAMETER_IDNOTEXIST = "InvalidParameter.IdNotExist"

	// 非法流程节点。
	INVALIDPARAMETER_ILLEGALNODE = "InvalidParameter.IllegalNode"

	// 没有权限审批。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
)
