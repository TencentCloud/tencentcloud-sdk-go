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

package v20190319

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 检查cls日志主题是否存在失败。
	FAILEDOPERATION_CHECKCLSTOPICISEXISTFAILED = "FailedOperation.CheckClsTopicIsExistFailed"

	// 检查cos桶是否存在失败。
	FAILEDOPERATION_CHECKCOSBUCKETISEXISTFAILED = "FailedOperation.CheckCosBucketIsExistFailed"

	// 拉取cls日志主题失败。
	FAILEDOPERATION_GETCLSTOPICFAILED = "FailedOperation.GetClsTopicFailed"

	// 拉取cos存储桶列表失败。
	FAILEDOPERATION_GETCOSBUCKETLISTFAILED = "FailedOperation.GetCosBucketListFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 创建cmq时发生异常，可能您准备创建的cmq队列已经存在，也有可能您没有权限或者欠费。
	INTERNALERROR_CMQERROR = "InternalError.CmqError"

	// 查看跟踪集详情错误，请联系开发人员
	INTERNALERROR_DESCRIBEAUDITERROR = "InternalError.DescribeAuditError"

	// 查询可创建跟踪集的数量错误，请联系开发人员
	INTERNALERROR_INQUIREAUDITCREDITERROR = "InternalError.InquireAuditCreditError"

	// 查询跟踪集概要内部错误，请联系开发人员。
	INTERNALERROR_LISTAUDITSERROR = "InternalError.ListAuditsError"

	// 内部错误，请联系开发人员
	INTERNALERROR_LISTCMQENABLEREGIONERROR = "InternalError.ListCmqEnableRegionError"

	// 内部错误，请联系开发人员
	INTERNALERROR_LISTCOSENABLEREGIONERROR = "InternalError.ListCosEnableRegionError"

	// 内部错误，请联系开发人员
	INTERNALERROR_LISTKEYALIASBYREGIONERROR = "InternalError.ListKeyAliasByRegionError"

	// 内部错误，请联系开发人员
	INTERNALERROR_SEARCHERROR = "InternalError.SearchError"

	// 内部错误，请联系开发人员
	INTERNALERROR_STARTLOGGINGERROR = "InternalError.StartLoggingError"

	// 内部错误，请联系开发人员
	INTERNALERROR_STOPLOGGINGERROR = "InternalError.StopLoggingError"

	// 内部错误，请联系开发人员
	INTERNALERROR_UPDATEAUDITERROR = "InternalError.UpdateAuditError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 必须包含开始时间和结束时间，且必须为整形时间戳（精确到秒）
	INVALIDPARAMETER_TIME = "InvalidParameter.Time"

	// 别名已经存在
	INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"

	// 跟踪集名称不支持修改。
	INVALIDPARAMETERVALUE_AUDITTRACKNAMENOTSUPPORTMODIFY = "InvalidParameterValue.AuditTrackNameNotSupportModify"

	// 操作审计目前不支持输入的cmq地域
	INVALIDPARAMETERVALUE_CMQREGIONERROR = "InvalidParameterValue.CmqRegionError"

	// 输入的cos存储桶名称不符合规范
	INVALIDPARAMETERVALUE_COSNAMEERROR = "InvalidParameterValue.CosNameError"

	// 操作审计目前不支持输入的cos地域
	INVALIDPARAMETERVALUE_COSREGIONERROR = "InvalidParameterValue.CosRegionError"

	// 操作审计暂不支持该地域的KMS
	INVALIDPARAMETERVALUE_KMSREGIONERROR = "InvalidParameterValue.KmsRegionError"

	// 日志前缀格式错误
	INVALIDPARAMETERVALUE_LOGFILEPREFIXERROR = "InvalidParameterValue.LogFilePrefixError"

	// 单次检索支持的最大返回条数是50
	INVALIDPARAMETERVALUE_MAXRESULT = "InvalidParameterValue.MaxResult"

	// 输入的队列名称不符合规范
	INVALIDPARAMETERVALUE_QUEUENAMEERROR = "InvalidParameterValue.QueueNameError"

	// 读写属性值仅支持：1,2,3。1代表只读，2代表只写，3代表全部。
	INVALIDPARAMETERVALUE_READWRITEATTRIBUTEERROR = "InvalidParameterValue.ReadWriteAttributeError"

	// 开始时间不能大于结束时间
	INVALIDPARAMETERVALUE_TIME = "InvalidParameterValue.Time"

	// AttributeKey的有效取值范围是:RequestId、EventName、ReadOnly、Username、ResourceType、ResourceName和AccessKeyId
	INVALIDPARAMETERVALUE_ATTRIBUTEKEY = "InvalidParameterValue.attributeKey"

	// 超过跟踪集的最大值
	LIMITEXCEEDED_OVERAMOUNT = "LimitExceeded.OverAmount"

	// 检索支持的有效时间范围是7天
	LIMITEXCEEDED_OVERTIME = "LimitExceeded.OverTime"

	// IsEnableCmqNotify输入1的话，IsCreateNewQueue、CmqQueueName和CmqRegion都是必须参数。
	MISSINGPARAMETER_CMQ = "MissingParameter.cmq"

	// cos存储桶已经存在
	RESOURCEINUSE_COSBUCKETEXISTS = "ResourceInUse.CosBucketExists"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 跟踪集不存在
	RESOURCENOTFOUND_AUDITNOTEXIST = "ResourceNotFound.AuditNotExist"

	// cos桶不存在。
	RESOURCENOTFOUND_COSNOTEXIST = "ResourceNotFound.CosNotExist"

	// 角色不存在。
	RESOURCENOTFOUND_ROLENOTEXIST = "ResourceNotFound.RoleNotExist"
)
