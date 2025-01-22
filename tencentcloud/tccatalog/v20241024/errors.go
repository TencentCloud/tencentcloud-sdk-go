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

package v20241024

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 时间格式异常，仅支持：YYYY-mm-dd HH:MM:SS
	INVALIDPARAMETER_INVALIDTIMEFORMAT = "InvalidParameter.InvalidTimeFormat"

	// 无效的时间，结束时间应大于起始时间
	INVALIDPARAMETER_INVALIDTIMEPARAMETER = "InvalidParameter.InvalidTimeParameter"

	// 集群未处于运行状态
	RESOURCEUNAVAILABLE_CLUSTERUNAVAILABLE = "ResourceUnavailable.ClusterUnavailable"

	// 该用户禁止当前操作
	UNAUTHORIZEDOPERATION_USERNOTALLOWOPERATION = "UnauthorizedOperation.UserNotAllowOperation"
)
