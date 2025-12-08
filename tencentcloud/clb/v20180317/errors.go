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

package v20180317

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// token错误
	AUTHFAILURE_TOKENFAILURE = "AuthFailure.TokenFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 该实例绑定的EIP资源业务带宽超过防误操作检测阈值，执行操作可能存在风险，如仍想继续，请在EIP控制台关闭对应EIP的防误操作检测开关后重试。
	FAILEDOPERATION_EIPTRAFFICCHECKRISK = "FailedOperation.EipTrafficCheckRisk"

	// 删除实例频次校验被判定为高风险，请检查业务或稍后重试。
	FAILEDOPERATION_FREQUENCYCHECKRISK = "FailedOperation.FrequencyCheckRisk"

	// LB状态异常。
	FAILEDOPERATION_INVALIDLBSTATUS = "FailedOperation.InvalidLBStatus"

	// 没有监听器的实例不允许停止。
	FAILEDOPERATION_NOLISTENERINLB = "FailedOperation.NoListenerInLB"

	// 指定的资源正在克隆中，请稍后重试。
	FAILEDOPERATION_RESOURCEINCLONING = "FailedOperation.ResourceInCloning"

	// 正在操作指定的资源，请稍后重试。
	FAILEDOPERATION_RESOURCEINOPERATING = "FailedOperation.ResourceInOperating"

	// 删除实例规则数校验被判定为高风险，如需强行删除请传强制校验参数ForceDelete为true。
	FAILEDOPERATION_TARGETNUMCHECKRISK = "FailedOperation.TargetNumCheckRisk"

	// 删除实例流量校验被判定为高风险，如需强行删除请传强制校验参数ForceDelete为true。
	FAILEDOPERATION_TRAFFICCHECKRISK = "FailedOperation.TrafficCheckRisk"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 为确保资源不泄露，保证创建的资源ID幂等性。通过ClientToken创建资源，当订单流程已结束且发货失败，或订单流程长时间未更新时，提示当前ClientToken已经超时过期。
	INVALIDPARAMETER_CLIENTTOKENLIMITEXCEEDED = "InvalidParameter.ClientTokenLimitExceeded"

	// 参数格式错误。
	INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"

	// 查询参数错误。
	INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"

	// 负载均衡实例ID错误。
	INVALIDPARAMETER_LBIDNOTFOUND = "InvalidParameter.LBIdNotFound"

	// 监听器ID错误。
	INVALIDPARAMETER_LISTENERIDNOTFOUND = "InvalidParameter.ListenerIdNotFound"

	// 查找不到符合条件的转发规则。
	INVALIDPARAMETER_LOCATIONNOTFOUND = "InvalidParameter.LocationNotFound"

	// 监听器端口检查失败，比如端口冲突。
	INVALIDPARAMETER_PORTCHECKFAILED = "InvalidParameter.PortCheckFailed"

	// 监听器协议检查失败，比如相关协议不支持对应操作。
	INVALIDPARAMETER_PROTOCOLCHECKFAILED = "InvalidParameter.ProtocolCheckFailed"

	// 地域无效。
	INVALIDPARAMETER_REGIONNOTFOUND = "InvalidParameter.RegionNotFound"

	// 转发规则已绑定重定向关系。
	INVALIDPARAMETER_REWRITEALREADYEXIST = "InvalidParameter.RewriteAlreadyExist"

	// 一些重定向规则不存在。
	INVALIDPARAMETER_SOMEREWRITENOTFOUND = "InvalidParameter.SomeRewriteNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 参数值有重复。
	INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"

	// Filter参数输入错误。
	INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"

	// 参数长度错误。
	INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"

	// 参数取值范围错误。
	INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"

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

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
