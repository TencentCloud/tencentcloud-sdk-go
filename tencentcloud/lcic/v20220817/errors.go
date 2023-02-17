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

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 房间状态错误，房间已结束。
	FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"

	// 房间状态错误，房间已过期。
	FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"

	// 课堂状态错误，课堂已开始。
	FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"

	// 课堂时长不能超过5小时。
	FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"

	// 图片参数错误。
	FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"

	// 源账号已存在。
	FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"

	// 房间暂未结束。
	FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 内容包含非法信息（如色情，恐暴，政治等）。
	INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"

	// 结束时间不能早于开始时间。
	INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"

	// 操作的群组成员超过限制
	INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"

	// 参数错误，主讲人与成员都为空
	INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"

	// 主讲人不能同时是群组成员
	INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"

	// 群组主讲老师不存在
	INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"

	// 群组类型错误
	INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"

	// SdkAppId参数错误。
	INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"

	// 开始时间不能早于当前时间。
	INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 存储空间已无剩余，无法使用存储功能。
	RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"

	// 打开课堂失败，请前往控制台查看用量情况。
	RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 文档不存在。
	RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"

	// 群组不存在
	RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"

	// 用户ID不存在
	RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"

	// 房间不存在。
	RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"

	// 用户不存在。
	RESOURCENOTFOUND_USER = "ResourceNotFound.User"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 房间统计结果计算中，请稍候。
	RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
