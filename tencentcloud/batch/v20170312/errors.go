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

package v20170312

const (
	// 此产品的特有错误码

	// ComputeEnv 和 EnvId 必须指定一个（且只有一个）参数。
	ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 调用Cpm API返回错误。
	INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"

	// 调用CVM API返回错误。
	INTERNALERROR_CALLCVM = "InternalError.CallCvm"

	// 获取Tag错误。
	INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"

	// 黑石服务返回数据为空。
	INTERNALERROR_CPMRESPONSEDATAEMPTY = "InternalError.CpmResponseDataEmpty"

	// 指定的Filter不被支持。
	INVALIDFILTER = "InvalidFilter"

	// 非法的计算节点ID格式。
	INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"

	// 非法的CVM参数。
	INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"

	// 计算环境描述过长。
	INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"

	// 非法的计算环境ID格式。
	INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"

	// 计算环境名称过长。
	INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"

	// 镜像ID不正确。
	INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"

	// 非法的参数组合。
	INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"

	// 作业描述过长。
	INVALIDPARAMETER_JOBDESCRIPTIONTOOLONG = "InvalidParameter.JobDescriptionTooLong"

	// 非法的作业ID格式。
	INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"

	// 作业名称过长。
	INVALIDPARAMETER_JOBNAMETOOLONG = "InvalidParameter.JobNameTooLong"

	// 重复的消息通知事件名称。
	INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"

	// 非法的主题名称。
	INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"

	// 主题名称过长。
	INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"

	// 非法的任务名称。
	INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"

	// 任务名称过长。
	INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"

	// 任务模板描述过长。
	INVALIDPARAMETER_TASKTEMPLATEDESCRIPTIONTOOLONG = "InvalidParameter.TaskTemplateDescriptionTooLong"

	// 非法的任务模板ID格式。
	INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"

	// 非法的任务模板名称。
	INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"

	// 任务模板名称过长。
	INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"

	// TaskTemplateName、TaskTemplateDescription和TaskTemplateInfo至少包含一项。
	INVALIDPARAMETERATLEASTONEATTRIBUTE = "InvalidParameterAtLeastOneAttribute"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 计算环境参数校验失败。
	INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"

	// 找不到依赖任务定义。
	INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"

	// 禁止环状任务依赖关系。
	INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"

	// 实例ID重复。
	INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"

	// 不支持指定的实例类型。
	INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"

	// 实例类型不能有重复值。
	INVALIDPARAMETERVALUE_INSTANCETYPEDUPLICATE = "InvalidParameterValue.InstanceTypeDuplicate"

	// 实例类型列表不能为空。
	INVALIDPARAMETERVALUE_INSTANCETYPESEMPTY = "InvalidParameterValue.InstanceTypesEmpty"

	// DataTypeAny不合法。
	INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"

	// Filter参数不正确。
	INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"

	// 可用区和地域不匹配。
	INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"

	// Filter参数值数量超过限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 非法的本地存储路径。
	INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"

	// 重试次数过大。
	INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"

	// 非法的负值参数。
	INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"

	// 参数值不是浮点型。
	INVALIDPARAMETERVALUE_NOTFLOAT = "InvalidParameterValue.NotFloat"

	// 指定的OsTypeId不合法。
	INVALIDPARAMETERVALUE_OSTYPEID = "InvalidParameterValue.OsTypeId"

	// 该地域不支持创建黑石计算环境。
	INVALIDPARAMETERVALUE_REGIONNOTSUPPORTCPM = "InvalidParameterValue.RegionNotSupportCpm"

	// 非法的存储路径格式。
	INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"

	// 非法的存储类型。
	INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"

	// 指定的ResourceType不合法。
	INVALIDPARAMETERVALUE_RESOURCETYPE = "InvalidParameterValue.ResourceType"

	// Zone不可用。
	INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"

	// 批量计算不支持的机型付费类型。
	INVALIDPARAMETERVALUE_UNSUPPORTEDBATCHINSTANCECHARGETYPE = "InvalidParameterValue.UnsupportedBatchInstanceChargeType"

	// 指定的zone不存在。
	INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"

	// 计算环境配额不足。
	LIMITEXCEEDED_COMPUTEENVQUOTA = "LimitExceeded.ComputeEnvQuota"

	// CPU配额不足。
	LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"

	// 作业配额不足。
	LIMITEXCEEDED_JOBQUOTA = "LimitExceeded.JobQuota"

	// 任务模板配额不足。
	LIMITEXCEEDED_TASKTEMPLATEQUOTA = "LimitExceeded.TaskTemplateQuota"

	// Job在使用中。
	RESOURCEINUSE_JOB = "ResourceInUse.Job"

	// 指定计算环境不存在。
	RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"

	// 指定计算节点不存在。
	RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"

	// 指定作业不存在。
	RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"

	// 指定作业任务不存在。
	RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"

	// 指定任务实例不存在。
	RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"

	// 指定任务模板ID不存在。
	RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"

	// 禁止使用批量计算服务。
	UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 处理其他请求，禁止删除操作。
	UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"

	// 计算环境正在接受请求。
	UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"

	// 禁止删除操作。
	UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"

	// 禁止终止计算节点。
	UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"

	// 计算节点停止中。
	UNSUPPORTEDOPERATION_COMPUTENODEISTERMINATING = "UnsupportedOperation.ComputeNodeIsTerminating"

	// 不允许将该实例添加到计算环境中。
	UNSUPPORTEDOPERATION_INSTANCESNOTALLOWTOATTACH = "UnsupportedOperation.InstancesNotAllowToAttach"

	// 缩容时，缩容数大于现有计算节点数。
	UNSUPPORTEDOPERATION_NOTENOUGHCOMPUTENODESTOTERMINATE = "UnsupportedOperation.NotEnoughComputeNodesToTerminate"

	// 指定计算环境的任务实例禁止该操作。
	UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
)
