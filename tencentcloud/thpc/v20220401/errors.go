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

package v20220401

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// cvm调用失败。
	INTERNALERROR_CALLCVM = "InternalError.CallCvm"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数格式有误。
	INVALIDPARAMETER_MALFORMED = "InvalidParameter.Malformed"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 不支持指定过滤器的键。
	INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"

	// 参数值数量超过限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 不支持该参数值。
	INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"

	// 字段不支持此值。
	INVALIDPARAMETERVALUE_PARAMETERSNOTSUPPORTED = "InvalidParameterValue.ParametersNotSupported"

	// 参数值过大。
	INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"

	// 参数长度过长。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 无效参数值。参数值太短。
	INVALIDPARAMETERVALUE_TOOSHORT = "InvalidParameterValue.TooShort"

	// 参数值过小。
	INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"

	// 参数值重复。不支持此操作。
	INVALIDPARAMETERVALUE_VALUEDUPLICATED = "InvalidParameterValue.ValueDuplicated"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 队列数量达到上限。
	LIMITEXCEEDED_QUEUENUMLIMIT = "LimitExceeded.QueueNumLimit"

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

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 无法找到弹性伸缩组ID。
	RESOURCENOTFOUND_AUTOSCALINGGROUPID = "ResourceNotFound.AutoScalingGroupId"

	// 集群不存在。
	RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"

	// 无法找到镜像ID。
	RESOURCENOTFOUND_IMAGEID = "ResourceNotFound.ImageId"

	// 无法找到ID对应的弹性伸缩启动配置。
	RESOURCENOTFOUND_LAUNCHCONFIGURATIONID = "ResourceNotFound.LaunchConfigurationId"

	// 无法找到ID对应节点。
	RESOURCENOTFOUND_NODEID = "ResourceNotFound.NodeId"

	// 无法找到指定队列。
	RESOURCENOTFOUND_QUEUE = "ResourceNotFound.Queue"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 该伸缩组已绑定集群，请更换伸缩组。
	UNSUPPORTEDOPERATION_AUTOSCALINGGROUPALREADYBINDED = "UnsupportedOperation.AutoScalingGroupAlreadyBinded"

	// 弹性伸缩类型不支持此操作。
	UNSUPPORTEDOPERATION_AUTOSCALINGTYPE = "UnsupportedOperation.AutoScalingType"

	// 该集群当前状态不支持该操作。
	UNSUPPORTEDOPERATION_CLUSTERSTATUSNOTSUPPORT = "UnsupportedOperation.ClusterStatusNotSupport"

	// 类型节点不支持当前操作。
	UNSUPPORTEDOPERATION_INVALIDNODEROLE = "UnsupportedOperation.InvalidNodeRole"

	// 节点状态不支持此操作。
	UNSUPPORTEDOPERATION_NODESTATUSNOTSUPPORT = "UnsupportedOperation.NodeStatusNotSupport"

	// 参数值过大，不支持此操作。
	UNSUPPORTEDOPERATION_PARAMETERTOOLARGE = "UnsupportedOperation.ParameterTooLarge"

	// 参数值过小，不支持此操作。
	UNSUPPORTEDOPERATION_PARAMETERTOOSMALL = "UnsupportedOperation.ParameterTooSmall"

	// 队列内存在节点，不支持此操作。
	UNSUPPORTEDOPERATION_QUEUENOTEMPTY = "UnsupportedOperation.QueueNotEmpty"

	// vpc冲突，不支持当前操作。
	UNSUPPORTEDOPERATION_VPCIDCONFLICT = "UnsupportedOperation.VpcIdConflict"
)
