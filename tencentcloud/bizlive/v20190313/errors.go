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

package v20190313

const (
	// 此产品的特有错误码

	// 操作失败
	FAILEDOPERATION = "FailedOperation"

	// 带宽不足
	FAILEDOPERATION_LACKBANDWIDTH = "FailedOperation.LackBandwidth"

	// 内部错误
	INTERNALERROR = "InternalError"

	// 调用内部服务错误。
	INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"

	// 配置不存在。
	INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"

	// DB执行错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 获取用户账号错误。
	INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"

	// 获取流信息失败。
	INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"

	// 获取直播源信息错误。
	INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"

	// 无权限操作。
	INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"

	// 流状态异常。
	INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"

	// 更新数据失败。
	INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"

	// 参数错误
	INVALIDPARAMETER = "InvalidParameter"

	// Json解析失败
	INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"

	// 参数取值错误
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 确认插件是否有IM能力
	LIMITEXCEEDED_NOIMABILITY = "LimitExceeded.NoIMAbility"

	// 缺少参数错误
	MISSINGPARAMETER = "MissingParameter"

	// 没有空闲机器
	RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
)
