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

package v20220817

const (
	// 此产品的特有错误码

	// 课堂状态错误，课堂已开始。
	FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"

	// 课堂时长不能超过5小时。
	FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"

	// 源账号已存在。
	FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"

	// 房间暂未结束。
	FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 结束时间不能早于开始时间。
	INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"

	// SdkAppId参数错误。
	INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"

	// 开始时间不能早于当前时间。
	INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 存储空间已无剩余，无法使用存储功能。
	RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"

	// 打开课堂失败，请前往控制台查看用量情况。
	RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"

	// 文档不存在。
	RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"

	// 房间不存在。
	RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"

	// 用户不存在。
	RESOURCENOTFOUND_USER = "ResourceNotFound.User"

	// 房间统计结果计算中，请稍候。
	RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
)
