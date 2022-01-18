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

package v20191022

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 尚未开通CLS日志服务，请开前往开通。
	FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"

	// 名称重复。
	FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"

	// 操作不允许。
	FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"

	// 存储库有绑定的实例，请先删除绑定的实例。
	FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 名称不合法。
	INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"

	// 请设置日志集、日志主题ID。
	INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"

	// 存储库不存在。
	INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"

	// 实例名称冲突，请更换名称后重试。
	INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"

	// 文件系统作为训练数据源时需要指定VPC配置。
	INVALIDPARAMETERVALUE_FILESYSTEMNEEDVPCCONFIGSUPPORT = "InvalidParameterValue.FileSystemNeedVpcConfigSupport"

	// 文件系统超过限额，最大为2。
	INVALIDPARAMETERVALUE_FILESYSTEMNUMLIMIT = "InvalidParameterValue.FileSystemNumLimit"

	// 文件系统的vpc必须和任务的vpc相同。
	INVALIDPARAMETERVALUE_FILESYSTEMVPCNOTMATCH = "InvalidParameterValue.FileSystemVpcNotMatch"

	// 训练框架对应的版本不支持，请阅读文档查看TIONE目前支持的框架和版本。
	INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"

	// 训练任务镜像不存在。
	INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"

	// 无效的资源类型,支持的资源类型参考：https://cloud.tencent.com/document/product/851/41239。
	INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"

	// 训练任务镜像名称无效。
	INVALIDPARAMETERVALUE_INVALIDTRAININGIMAGENAME = "InvalidParameterValue.InvalidTrainingImageName"

	// KMS密钥不存在。
	INVALIDPARAMETERVALUE_KMSKEYNOTFOUND = "InvalidParameterValue.KmsKeyNotFound"

	// 未找到当前日志集。
	INVALIDPARAMETERVALUE_LOGSETNOTFOUND = "InvalidParameterValue.LogSetNotFound"

	// MPI分布式任务参数ti_mpi_num_of_processes_per_host不能超过GPU卡数。
	INVALIDPARAMETERVALUE_MPIPROCESSESPERHOSTTOOMUCH = "InvalidParameterValue.MpiProcessesPerHostTooMuch"

	// 存储库地址无效。
	INVALIDPARAMETERVALUE_REPOSITORYURLINVALID = "InvalidParameterValue.RepositoryUrlInvalid"

	// 子网不存在。
	INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"

	// 未找到当前日志主题。
	INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"

	// 训练代码不存在或无效。
	INVALIDPARAMETERVALUE_TRAINCODENOTFOUND = "InvalidParameterValue.TrainCodeNotFound"

	// Notebook卷大小只能增加，如需减小容量请重新创建实例。
	INVALIDPARAMETERVALUE_VOLUMESHRINKNOTALLOW = "InvalidParameterValue.VolumeShrinkNotAllow"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未开通该地域服务，请先开通。
	RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"

	// 实例未成功启动。
	RESOURCEUNAVAILABLE_NOTALIVE = "ResourceUnavailable.NotAlive"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
