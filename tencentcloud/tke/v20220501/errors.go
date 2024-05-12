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

package v20220501

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// db错误。
	INTERNALERROR_DB = "InternalError.Db"

	// DB错误。
	INTERNALERROR_DBAFFECTIVEDROWS = "InternalError.DbAffectivedRows"

	// 记录未找到。
	INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"

	// 初始化master失败。
	INTERNALERROR_INITMASTERFAILED = "InternalError.InitMasterFailed"

	// Param。
	INTERNALERROR_PARAM = "InternalError.Param"

	// 公共集群不支持扩展节点。
	INTERNALERROR_PUBLICCLUSTEROPNOTSUPPORT = "InternalError.PublicClusterOpNotSupport"

	// 内部错误。
	INTERNALERROR_UNEXCEPTEDINTERNAL = "InternalError.UnexceptedInternal"

	// 未知的内部错误。
	INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 集群ID不存在。
	INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"

	// 参数错误。
	INVALIDPARAMETER_PARAM = "InvalidParameter.Param"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 集群不存在。
	RESOURCENOTFOUND_CLUSTERNOTFOUND = "ResourceNotFound.ClusterNotFound"

	// 集群状态不支持该操作。
	RESOURCEUNAVAILABLE_CLUSTERSTATE = "ResourceUnavailable.ClusterState"
)
