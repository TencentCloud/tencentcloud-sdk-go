// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20251030

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 资源已存在。
	FAILEDOPERATION_RESOURCEALREADYEXISTS = "FailedOperation.ResourceAlreadyExists"

	// 规则优先级已存在。
	FAILEDOPERATION_RULEALREADYEXISTS = "FailedOperation.RuleAlreadyExists"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 无效的过滤器
	INVALIDFILTER = "InvalidFilter"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// ClientToken 格式无效。
	INVALIDPARAMETER_INVALIDCLIENTTOKEN = "InvalidParameter.InvalidClientToken"

	// 字段格式错误。
	INVALIDPARAMETER_INVALIDFIELDFORMAT = "InvalidParameter.InvalidFieldFormat"

	// 字段类型错误。
	INVALIDPARAMETER_INVALIDFIELDTYPE = "InvalidParameter.InvalidFieldType"

	// 字段值错误。
	INVALIDPARAMETER_INVALIDFIELDVALUE = "InvalidParameter.InvalidFieldValue"

	// 配额检查请求无效。
	INVALIDPARAMETER_INVALIDQUOTACHECKREQUEST = "InvalidParameter.InvalidQuotaCheckRequest"

	// 缺少 Action 参数。
	INVALIDPARAMETER_MISSINGACTION = "InvalidParameter.MissingAction"

	// 请填写必要参数。
	INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 业务参数不一致。
	INVALIDPARAMETERVALUE_BUSINESSPARAMETERMISMATCH = "InvalidParameterValue.BusinessParameterMismatch"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 配额超限。
	LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 地域错误
	REGIONERROR = "RegionError"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 实例下还存在监听器，无法删除。
	RESOURCEINUSE_LOADBALANCERHASLISTENERS = "ResourceInUse.LoadBalancerHasListeners"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 证书未找到。
	RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"

	// 监听器未找到。
	RESOURCENOTFOUND_LISTENER = "ResourceNotFound.Listener"

	// 负载均衡资源不存在。
	RESOURCENOTFOUND_LOADBALANCER = "ResourceNotFound.LoadBalancer"

	// 转发规则未找到
	RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"

	// 目标组未找到。
	RESOURCENOTFOUND_TARGETGROUP = "ResourceNotFound.TargetGroup"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 证书处于不可用状态（更新中/删除中）。
	RESOURCEUNAVAILABLE_CERTIFICATE = "ResourceUnavailable.Certificate"

	// 监听器处于不可用状态。
	RESOURCEUNAVAILABLE_LISTENER = "ResourceUnavailable.Listener"

	// ALB状态不可用。
	RESOURCEUNAVAILABLE_LOADBALANCER = "ResourceUnavailable.LoadBalancer"

	// 上个任务未完成。
	RESOURCEUNAVAILABLE_PREVIOUSTASKNOTCOMPLETED = "ResourceUnavailable.PreviousTaskNotCompleted"

	// 规则状态不可用。
	RESOURCEUNAVAILABLE_RULE = "ResourceUnavailable.Rule"

	// 目标组处于不可用状态
	RESOURCEUNAVAILABLE_TARGETGROUP = "ResourceUnavailable.TargetGroup"

	// 上级资源状态阻断。
	RESOURCEUNAVAILABLE_UPSTREAMSTATUSBLOCKED = "ResourceUnavailable.UpstreamStatusBlocked"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 服务角色未授权
	UNAUTHORIZEDOPERATION_SERVICEROLEUNAUTHORIZED = "UnauthorizedOperation.ServiceRoleUnauthorized"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 监听器不允许绑定已绑定的证书。
	UNSUPPORTEDOPERATION_CERTIFICATEALREADYBOUND = "UnsupportedOperation.CertificateAlreadyBound"

	// 监听器不允许解绑未绑定的证书。
	UNSUPPORTEDOPERATION_CERTIFICATENOTBOUND = "UnsupportedOperation.CertificateNotBound"

	// 监听器不允许解绑默认证书。
	UNSUPPORTEDOPERATION_DISASSOCIATEDEFAULTCERTIFICATE = "UnsupportedOperation.DisassociateDefaultCertificate"

	// 无效的状态转换。
	UNSUPPORTEDOPERATION_INVALIDSTATETRANSITION = "UnsupportedOperation.InvalidStateTransition"

	// 监听器协议与目标组后端转发协议不匹配
	UNSUPPORTEDOPERATION_TARGETGROUPPROTOCOLMISMATCH = "UnsupportedOperation.TargetGroupProtocolMismatch"

	// 监听器协议不支持此操作。
	UNSUPPORTEDOPERATION_UNSUPPORTEDPROTOCOL = "UnsupportedOperation.UnsupportedProtocol"
)
