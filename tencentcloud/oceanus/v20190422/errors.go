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

package v20190422

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 未授权操作。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 数据库连接失败，请检查参数是否填写正确。
	FAILEDOPERATION_DATASOURCECONNECTIONFAILED = "FailedOperation.DataSourceConnectionFailed"

	// 重复的作业名。
	FAILEDOPERATION_DUPLICATEDJOBNAME = "FailedOperation.DuplicatedJobName"

	// 查询资源关联标签失败。
	FAILEDOPERATION_GETRESOURCETAGSBYRESOURCEIDS = "FailedOperation.GetResourceTagsByResourceIds"

	// 语法检查失败。
	FAILEDOPERATION_GRAMMARCHECKFAILURE = "FailedOperation.GrammarCheckFailure"

	// SQL解析失败。
	FAILEDOPERATION_PARSESQL = "FailedOperation.ParseSql"

	// 用户未实名验证。
	FAILEDOPERATION_USERNOTAUTHENTICATED = "FailedOperation.UserNotAuthenticated"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// CLS接口错误。
	INTERNALERROR_CLS = "InternalError.CLS"

	// COS 服务访问错误。
	INTERNALERROR_COSCLIENT = "InternalError.COSClient"

	// CAM 网关错误。
	INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"

	// 数据库异常。
	INTERNALERROR_DB = "InternalError.DB"

	// 失败的目标资源。
	INTERNALERROR_FAILEDTOBESCRIBERESOURCES = "InternalError.FailedToBescribeResources"

	// 无法更新作业错误。
	INTERNALERROR_FAILEDTOUPDATEJOB = "InternalError.FailedToUpdateJob"

	// 作业实例没找到。
	INTERNALERROR_JOBINSTANCENOTFOUND = "InternalError.JobInstanceNotFound"

	// 内部错误。
	INTERNALERROR_LOGICERROR = "InternalError.LogicError"

	// 资源只有一个版本，无法删除。
	INTERNALERROR_RESOURCECONFIGCANNOTDELETE = "InternalError.ResourceConfigCanNotDelete"

	// 资源不存在。
	INTERNALERROR_RESOURCENOTEXIST = "InternalError.ResourceNotExist"

	// 未找到sqlcode错误。
	INTERNALERROR_SQLCODENOTFOUND = "InternalError.SqlCodeNotFound"

	// 内部错误。
	INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"

	// 系统错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// AppId资源不匹配。
	INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"

	// 非法的 MaxParallelism 参数。
	INVALIDPARAMETER_ILLEGALMAXPARALLELISM = "InvalidParameter.IllegalMaxParallelism"

	// appid错误。
	INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"

	// 无效集群id。
	INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"

	// 名字不符合规范。
	INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"

	// 无效Region。
	INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"

	// ResourceIds非法。
	INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"

	// 创建作业配置错误。
	INVALIDPARAMETER_JOBCONFIGLOGCOLLECTPARAMERROR = "InvalidParameter.JobConfigLogCollectParamError"

	// MaxParallelism 过大。
	INVALIDPARAMETER_MAXPARALLELISMTOOLARGE = "InvalidParameter.MaxParallelismTooLarge"

	// MaxParallelism 不允许小于算子默认并行度。
	INVALIDPARAMETER_MAXPARALLELISMTOOSMALL = "InvalidParameter.MaxParallelismTooSmall"

	// Uin资源不匹配。
	INVALIDPARAMETER_UINRESOURCENOTMATCH = "InvalidParameter.UinResourceNotMatch"

	// Flink参数非法。
	INVALIDPARAMETER_UNSUPPORTEDFLINKCONF = "InvalidParameter.UnsupportedFlinkConf"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 集群Id指定错误，为空或者不存在。
	INVALIDPARAMETERVALUE_CLUSTERID = "InvalidParameterValue.ClusterId"

	// 群集Id报错。
	INVALIDPARAMETERVALUE_CLUSTERIDS = "InvalidParameterValue.ClusterIds"

	// CU内存规格不匹配。
	INVALIDPARAMETERVALUE_CUMEM = "InvalidParameterValue.CuMem"

	// 无效启动模式。
	INVALIDPARAMETERVALUE_INVALIDSTARTMODE = "InvalidParameterValue.InvalidStartMode"

	// 作业id的参数无效。
	INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"

	// Illegal JobName。
	INVALIDPARAMETERVALUE_JOBNAME = "InvalidParameterValue.JobName"

	// 作业名称已存在。
	INVALIDPARAMETERVALUE_JOBNAMEEXISTED = "InvalidParameterValue.JobNameExisted"

	// 集群模式与作业类型不匹配。
	INVALIDPARAMETERVALUE_JOBTYPECOMBINEWITHCLUSTERTYPE = "InvalidParameterValue.JobTypeCombineWithClusterType"

	// SQL作业不能指定EntrypointClass，JAR作业则必须指定。
	INVALIDPARAMETERVALUE_JOBTYPECOMBINEWITHENTRYPOINTCLASS = "InvalidParameterValue.JobTypeCombineWithEntrypointClass"

	// OrderType值错误。
	INVALIDPARAMETERVALUE_ORDERTYPE = "InvalidParameterValue.OrderType"

	// 未找到资源ID。
	INVALIDPARAMETERVALUE_RESOURCEIDSNOTFOUND = "InvalidParameterValue.ResourceIdsNotFound"

	// 批量运行作业个数超过上限。
	INVALIDPARAMETERVALUE_RUNJOBDESCRIPTIONSCOUNT = "InvalidParameterValue.RunJobDescriptionsCount"

	// RunType错误。
	INVALIDPARAMETERVALUE_RUNTYPE = "InvalidParameterValue.RunType"

	// 不支持的复合类型。
	INVALIDPARAMETERVALUE_UNSUPPORTEDCOMPOSITE = "InvalidParameterValue.UnSupportedComposite"

	// 未知停止类型错误。
	INVALIDPARAMETERVALUE_UNKNOWNSTOPTYPE = "InvalidParameterValue.UnknownStopType"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// Job个数超过限额。
	LIMITEXCEEDED_JOB = "LimitExceeded.Job"

	// 作业配置超出限制。
	LIMITEXCEEDED_JOBCONFIG = "LimitExceeded.JobConfig"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 集群有其他操作。
	RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"

	// 资源名称已存在。
	RESOURCEINUSE_RESOURCENAMEALREADYEXISTS = "ResourceInUse.ResourceNameAlreadyExists"

	// CU资源不足。
	RESOURCEINSUFFICIENT_CU = "ResourceInsufficient.CU"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// COS Bucket 未找到或无权限访问。
	RESOURCENOTFOUND_COSBUCKET = "ResourceNotFound.COSBucket"

	// 集群不存在。
	RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"

	// 作业不存在。
	RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"

	// 作业配置版本不存在。
	RESOURCENOTFOUND_JOBCONFIG = "ResourceNotFound.JobConfig"

	// 找不到作业。
	RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"

	// 程序包不存在。
	RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"

	// 程序包版本不存在。
	RESOURCENOTFOUND_RESOURCECONFIG = "ResourceNotFound.ResourceConfig"

	// 资源不存在。
	RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 被某个作业配置使用。
	RESOURCEUNAVAILABLE_BEUSEBYSOMEJOBCONFIG = "ResourceUnavailable.BeUseBySomeJobConfig"

	// 检查资源位置是否存在错误。
	RESOURCEUNAVAILABLE_CHECKRESOURCELOCEXISTS = "ResourceUnavailable.CheckResourceLocExists"

	// 集群处于非运行状态。
	RESOURCEUNAVAILABLE_CLUSTER = "ResourceUnavailable.Cluster"

	// 群集组状态错误。
	RESOURCEUNAVAILABLE_CLUSTERGROUPSTATUS = "ResourceUnavailable.ClusterGroupStatus"

	// 失败的目标资源。
	RESOURCEUNAVAILABLE_FAILEDTOBESCRIBERESOURCES = "ResourceUnavailable.FailedToBescribeResources"

	// 获取发布的作业作业配置错误。
	RESOURCEUNAVAILABLE_GETJOBPUBLISHEDJOBCONFIG = "ResourceUnavailable.GetJobPublishedJobConfig"

	// 作业资源配置未就绪。
	RESOURCEUNAVAILABLE_JOBRESOURCECONFIGNOTREADY = "ResourceUnavailable.JobResourceConfigNotReady"

	// 找不到JobId的正在运行的作业实例。
	RESOURCEUNAVAILABLE_NORUNNINGJOBINSTANCESFOUNDFORJOBID = "ResourceUnavailable.NoRunningJobInstancesFoundForJobId"

	// 不允许删除错误。
	RESOURCEUNAVAILABLE_NOTALLOWEDTOBEDELETED = "ResourceUnavailable.NotAllowedToBeDeleted"

	// 不允许停止或暂停错误。
	RESOURCEUNAVAILABLE_NOTALLOWEDTOBESTOPORPAUSE = "ResourceUnavailable.NotAllowedToBeStopOrPause"

	// 共享群集只允许CuMem=4。
	RESOURCEUNAVAILABLE_REQCUMEM = "ResourceUnavailable.ReqCuMem"

	// cos上的程序包不存在。
	RESOURCEUNAVAILABLE_RESOURCELOCNOTEXISTS = "ResourceUnavailable.ResourceLocNotExists"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// Checkpoint 时间间隔，错误。
	UNSUPPORTEDOPERATION_INVALIDCHECKPOINTINTERVALERROR = "UnsupportedOperation.InvalidCheckpointIntervalError"

	// 权限拦截,没有进入权限。
	UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"

	// 不支持的启动模式。
	UNSUPPORTEDOPERATION_UNSUPPORTEDSTARTMODE = "UnsupportedOperation.UnsupportedStartMode"
)
