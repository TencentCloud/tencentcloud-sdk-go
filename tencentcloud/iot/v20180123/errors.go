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

package v20180123

const (
	// 此产品的特有错误码

	// action系统错误。
	INTERNALERROR_IOTACTIONSYSTEMERROR = "InternalError.IotActionSystemError"

	// DB错误。
	INTERNALERROR_IOTDBERROR = "InternalError.IotDbError"

	// 日志系统错误。
	INTERNALERROR_IOTLOGSYSTEMERROR = "InternalError.IotLogSystemError"

	// shadow系统错误。
	INTERNALERROR_IOTSHADOWSYSTEMERROR = "InternalError.IotShadowSystemError"

	// 系统错误。
	INTERNALERROR_IOTSYSTEMERROR = "InternalError.IotSystemError"

	// Mqiot系统错误。
	INTERNALERROR_MQIOTSYSTEMERROR = "InternalError.MqiotSystemError"

	// rule系统错误。
	INTERNALERROR_MQRULESYSTEMERROR = "InternalError.MqruleSystemError"

	// 非法密码。
	INVALIDPARAMETER_IOTAPPLICATIONINVALIDPASSWORD = "InvalidParameter.IotApplicationInvalidPassword"

	// 用户名或密码错误。
	INVALIDPARAMETER_IOTAPPLICATIONINVALIDUSERPASSWORD = "InvalidParameter.IotApplicationInvalidUserPassword"

	// 过期访问Token。
	INVALIDPARAMETER_IOTEXPIREDACCESSTOKEN = "InvalidParameter.IotExpiredAccessToken"

	// 过期签名。
	INVALIDPARAMETER_IOTEXPIREDSIGNATURE = "InvalidParameter.IotExpiredSignature"

	// 非法访问Token。
	INVALIDPARAMETER_IOTINVALIDACCESSTOKEN = "InvalidParameter.IotInvalidAccessToken"

	// 非法签名。
	INVALIDPARAMETER_IOTINVALIDSIGNATURE = "InvalidParameter.IotInvalidSignature"

	// 非法参数。
	INVALIDPARAMETER_IOTPARAMERROR = "InvalidParameter.IotParamError"

	// 空数据模版。
	INVALIDPARAMETER_IOTPRODUCTEMPTYDATATEMPLATE = "InvalidParameter.IotProductEmptyDataTemplate"

	// 非法产品鉴权类型。
	INVALIDPARAMETER_IOTPRODUCTINVALIDAUTHTYPE = "InvalidParameter.IotProductInvalidAuthType"

	// 产品数据协议错误。
	INVALIDPARAMETER_IOTPRODUCTINVALIDDATAPROTOCOL = "InvalidParameter.IotProductInvalidDataProtocol"

	// 非法数据模版。
	INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATE = "InvalidParameter.IotProductInvalidDataTemplate"

	// 非法数据模版取值范围。
	INVALIDPARAMETER_IOTPRODUCTINVALIDDATATEMPLATERANGE = "InvalidParameter.IotProductInvalidDataTemplateRange"

	// 网关设备产品Id错误。
	INVALIDPARAMETER_IOTPRODUCTINVALIDGATEWAYPRODUCTID = "InvalidParameter.IotProductInvalidGatewayProductId"

	// 子设备产品Id错误。
	INVALIDPARAMETER_IOTPRODUCTINVALIDSUBDEVICEPRODUCTID = "InvalidParameter.IotProductInvalidSubDeviceProductId"

	// 查询日期错误。
	INVALIDPARAMETER_IOTSTATINVALIDDATE = "InvalidParameter.IotStatInvalidDate"

	// 批量操作超限。
	LIMITEXCEEDED_IOTBATCHTOOMANY = "LimitExceeded.IotBatchTooMany"

	// 设备操作太频繁。
	LIMITEXCEEDED_IOTDEVICEOPTOOOFTEN = "LimitExceeded.IotDeviceOpTooOften"

	// 产品操作太频繁。
	LIMITEXCEEDED_IOTPRODUCTOPTOOOFTEN = "LimitExceeded.IotProductOpTooOften"

	// 产品的Topics数量超限。
	LIMITEXCEEDED_IOTPRODUCTTOOMANYTOPICS = "LimitExceeded.IotProductTooManyTopics"

	// 规则批量操作超限。
	LIMITEXCEEDED_IOTRULEOPTOOMANY = "LimitExceeded.IotRuleOpTooMany"

	// 规则操作太频繁。
	LIMITEXCEEDED_IOTRULEOPTOOOFTEN = "LimitExceeded.IotRuleOpTooOften"

	// Topic操作太频繁。
	LIMITEXCEEDED_IOTTOPICOPTOOOFTEN = "LimitExceeded.IotTopicOpTooOften"

	// 用户的产品数超限。
	LIMITEXCEEDED_IOTUSERTOOMANYPRODUCTS = "LimitExceeded.IotUserTooManyProducts"

	// 设备已绑定。
	RESOURCEINUSE_IOTAPPLICATIONDEVICEEXISTS = "ResourceInUse.IotApplicationDeviceExists"

	// 应用用户已存在。
	RESOURCEINUSE_IOTAPPLICATIONUSEREXISTS = "ResourceInUse.IotApplicationUserExists"

	// 设备已存在。
	RESOURCEINUSE_IOTDEVICEEXISTS = "ResourceInUse.IotDeviceExists"

	// 正在处理上一个操作。
	RESOURCEINUSE_IOTOPINPROGRESS = "ResourceInUse.IotOpInProgress"

	// 产品已存在。
	RESOURCEINUSE_IOTPRODUCTEXISTS = "ResourceInUse.IotProductExists"

	// 规则已存在。
	RESOURCEINUSE_IOTRULEEXISTS = "ResourceInUse.IotRuleExists"

	// Topic已存在。
	RESOURCEINUSE_IOTTOPICEXISTS = "ResourceInUse.IotTopicExists"

	// 资源已存在。
	RESOURCEINUSE_MQIOTRESOURCEEXISTS = "ResourceInUse.MqiotResourceExists"

	// 设备未绑定。
	RESOURCENOTFOUND_IOTAPPLICATIONDEVICENOTEXISTS = "ResourceNotFound.IotApplicationDeviceNotExists"

	// 应用不存在。
	RESOURCENOTFOUND_IOTAPPLICATIONNOTEXISTS = "ResourceNotFound.IotApplicationNotExists"

	// 应用用户不存在。
	RESOURCENOTFOUND_IOTAPPLICATIONUSERNOTEXISTS = "ResourceNotFound.IotApplicationUserNotExists"

	// 设备不存在。
	RESOURCENOTFOUND_IOTDEVICENOTEXISTS = "ResourceNotFound.IotDeviceNotExists"

	// 产品不存在。
	RESOURCENOTFOUND_IOTPRODUCTNOTEXISTS = "ResourceNotFound.IotProductNotExists"

	// 规则不存在。
	RESOURCENOTFOUND_IOTRULENOTEXISTS = "ResourceNotFound.IotRuleNotExists"

	// 授权子账号不存在。
	RESOURCENOTFOUND_IOTSUBACCOUNTNOTEXISTS = "ResourceNotFound.IotSubAccountNotExists"

	// Topic不存在。
	RESOURCENOTFOUND_IOTTOPICNOTEXISTS = "ResourceNotFound.IotTopicNotExists"

	// 用户不存在。
	RESOURCENOTFOUND_IOTUSERNOTEXISTS = "ResourceNotFound.IotUserNotExists"

	// 规则config id不存在。
	RESOURCENOTFOUND_MQRULERULEIDNOTEXISTS = "ResourceNotFound.MqruleRuleIdNotExists"

	// 规则已启用。
	RESOURCEUNAVAILABLE_IOTRULEISACTIVE = "ResourceUnavailable.IotRuleIsActive"

	// 规则缺少动作。
	RESOURCEUNAVAILABLE_IOTRULENOACTION = "ResourceUnavailable.IotRuleNoAction"

	// 规则缺少查询。
	RESOURCEUNAVAILABLE_IOTRULENOQUERY = "ResourceUnavailable.IotRuleNoQuery"

	// 资源不存在。
	RESOURCEUNAVAILABLE_MQIOTRESOURCENOTEXISTS = "ResourceUnavailable.MqiotResourceNotExists"

	// 用户被封禁。
	UNAUTHORIZEDOPERATION_IOTUSERISSUSPENDED = "UnauthorizedOperation.IotUserIsSuspended"
)
