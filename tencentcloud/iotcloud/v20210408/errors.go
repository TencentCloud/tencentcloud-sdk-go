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

package v20210408

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 操作失败，账号已欠费隔离。
	FAILEDOPERATION_ACCOUNTISOLATED = "FailedOperation.AccountIsolated"

	// 已分发设备，不能再次创建。
	FAILEDOPERATION_ALREADYDISTRIBUTIONDEVICE = "FailedOperation.AlreadyDistributionDevice"

	// 绑定设备超过限制。
	FAILEDOPERATION_BINDDEVICEOVERLIMIT = "FailedOperation.BindDeviceOverLimit"

	// 单次绑定的设备数量超过限制。
	FAILEDOPERATION_BINDDEVICEPERONCEOVERLIMIT = "FailedOperation.BindDevicePerOnceOverLimit"

	// 广播任务正在执行。
	FAILEDOPERATION_BROADCASTTASKISRUNNING = "FailedOperation.BroadcastTaskIsRunning"

	// 设备已经被禁用。
	FAILEDOPERATION_DEVICEALREADYDISABLED = "FailedOperation.DeviceAlreadyDisabled"

	// 设备固件升级任务已经完成。
	FAILEDOPERATION_DEVICEFIRMWARETASKALREADDONE = "FailedOperation.DeviceFirmwareTaskAlreadDone"

	// 设备正在升级中。
	FAILEDOPERATION_DEVICEISUPDATING = "FailedOperation.DeviceIsUpdating"

	// 设备没有订阅相应的topic。
	FAILEDOPERATION_DEVICENOSUBSCRIPTION = "FailedOperation.DeviceNoSubscription"

	// 设备离线。
	FAILEDOPERATION_DEVICEOFFLINE = "FailedOperation.DeviceOffline"

	// 设备已经运行其他ota升级任务。
	FAILEDOPERATION_DEVICERUNNINGOTHEROTATASK = "FailedOperation.DeviceRunningOtherOtaTask"

	// 不能创建重复的函数。
	FAILEDOPERATION_DUPLICATIONOFFUNCTIONITEM = "FailedOperation.DuplicationOfFunctionItem"

	// 函数文件不存在。
	FAILEDOPERATION_FUNCTIONFILENOTEXIST = "FailedOperation.FunctionFileNotExist"

	// 消息长度非法。
	FAILEDOPERATION_INVALIDMSGLEN = "FailedOperation.InvalidMsgLen"

	// 消息topic非法。
	FAILEDOPERATION_INVALIDTOPICNAME = "FailedOperation.InvalidTopicName"

	// 产品未绑定，无法代理订阅。
	FAILEDOPERATION_PRODUCTNOTBIND = "FailedOperation.ProductNotBind"

	// 同名产品资源已存在。
	FAILEDOPERATION_PRODUCTRESOURCEDUPLICATE = "FailedOperation.ProductResourceDuplicate"

	// 代理ip或端口资源不足。
	FAILEDOPERATION_PROXYIPISNOTENOUGH = "FailedOperation.ProxyIPIsNotEnough"

	// RRPC接口未收到设备端响应。
	FAILEDOPERATION_RRPCTIMEOUT = "FailedOperation.RRPCTimeout"

	// 资源文件MD5或者大小不一致。
	FAILEDOPERATION_RESOURCEFILENOTMATCH = "FailedOperation.ResourceFileNotMatch"

	// 该规则引擎已经是禁用状态，不需要再被禁用。
	FAILEDOPERATION_RULEALREADYDISABLED = "FailedOperation.RuleAlreadyDisabled"

	// 规则已经是启用状态。
	FAILEDOPERATION_RULEALREADYENABLED = "FailedOperation.RuleAlreadyEnabled"

	// 白名单校验未开启，用户不可创建设备，平台会根据设备认证时携带的设备名称自动创建设备。
	FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"

	// 更新版本不匹配。
	FAILEDOPERATION_UPDATEVERSIONNOTMATCH = "FailedOperation.UpdateVersionNotMatch"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库内部错误。
	INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 规则行为未配置。
	INVALIDPARAMETERVALUE_ACTIONNIL = "InvalidParameterValue.ActionNil"

	// CA证书内容错误。
	INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"

	// CA验证证书不匹配。
	INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"

	// 检查第三方URL超时或失败。
	INVALIDPARAMETERVALUE_CHECKFORWARDURLFAIL = "InvalidParameterValue.CheckForwardURLFail"

	// 保存失败，行为操作和转发错误行为数据目标不可一致。
	INVALIDPARAMETERVALUE_CLOUDCOMPONENTALREADYEXIST = "InvalidParameterValue.CloudComponentAlreadyExist"

	// 格式错误，DefinedPsk需为Base64格式的字符串。
	INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"

	// 创建的设备名已存在。
	INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"

	// 设备不是网关类型。
	INVALIDPARAMETERVALUE_DEVICEISNOTGATEWAY = "InvalidParameterValue.DeviceIsNotGateway"

	// 存失败，行为操作和转发错误行为数据目标不可为同一设备。
	INVALIDPARAMETERVALUE_FAILACTIONHASSAMEDEVICE = "InvalidParameterValue.FailActionHasSameDevice"

	// 固件已存在。
	INVALIDPARAMETERVALUE_FIRMWAREALREADYEXIST = "InvalidParameterValue.FirmwareAlreadyExist"

	// 不允许转发重定向。
	INVALIDPARAMETERVALUE_FORWARDREDIRECTDENIED = "InvalidParameterValue.ForwardRedirectDenied"

	// JSON参数非法。
	INVALIDPARAMETERVALUE_INVALIDJSON = "InvalidParameterValue.InvalidJSON"

	// SQL语句含有非法字符。
	INVALIDPARAMETERVALUE_INVALIDSQL = "InvalidParameterValue.InvalidSQL"

	// State JSON对象中包含非法节点。
	INVALIDPARAMETERVALUE_JSONHASINVALIDNODE = "InvalidParameterValue.JSONHasInvalidNode"

	// State JSON对象超过大小限制，最大为 8k。
	INVALIDPARAMETERVALUE_JSONSIZEEXCEEDLIMIT = "InvalidParameterValue.JSONSizeExceedLimit"

	// 不可合并。
	INVALIDPARAMETERVALUE_NOTMERGEABLE = "InvalidParameterValue.NotMergeAble"

	// 修改规则的操作被禁止。
	INVALIDPARAMETERVALUE_OPERATIONDENIED = "InvalidParameterValue.OperationDenied"

	// 请求中缺少关键字段信息。
	INVALIDPARAMETERVALUE_PARAMINCOMPLETE = "InvalidParameterValue.ParamIncomplete"

	// 消息Payload超出限制。
	INVALIDPARAMETERVALUE_PAYLOADOVERLIMIT = "InvalidParameterValue.PayloadOverLimit"

	// prefix不合法。
	INVALIDPARAMETERVALUE_PREFIXINVALID = "InvalidParameterValue.PrefixInvalid"

	// 创建的产品名已存在。
	INVALIDPARAMETERVALUE_PRODUCTALREADYEXIST = "InvalidParameterValue.ProductAlreadyExist"

	// 产品类型不支持。
	INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"

	// 转发的topic格式错误。
	INVALIDPARAMETERVALUE_REPUBLISHTOPICFORMATERROR = "InvalidParameterValue.RepublishTopicFormatError"

	// 规则数量超过限制。
	INVALIDPARAMETERVALUE_RULENUMBERBEYONDLIMIT = "InvalidParameterValue.RuleNumberBeyondLimit"

	// 该TID产品已存在。
	INVALIDPARAMETERVALUE_TIDPRODUCTALREADYEXIST = "InvalidParameterValue.TidProductAlreadyExist"

	// Topic已存在。
	INVALIDPARAMETERVALUE_TOPICPOLICYALREADYEXIST = "InvalidParameterValue.TopicPolicyAlreadyExist"

	// 规则已存在。
	INVALIDPARAMETERVALUE_TOPICRULEALREADYEXIST = "InvalidParameterValue.TopicRuleAlreadyExist"

	// 规则sql未编辑。
	INVALIDPARAMETERVALUE_TOPICRULESQLNOTEDITED = "InvalidParameterValue.TopicRuleSqlNotEdited"

	// 请确认规则相关数据是否有更新。
	INVALIDPARAMETERVALUE_UPDATETOPICRULEDBFAIL = "InvalidParameterValue.UpdateTopicRuleDBFail"

	// CA证书已经绑定了产品，无法操作。
	LIMITEXCEEDED_CAALREADYBINDPRODUCT = "LimitExceeded.CAAlreadyBindProduct"

	// CA证书达到上限。
	LIMITEXCEEDED_CACERTLIMIT = "LimitExceeded.CACertLimit"

	// CA证书名称重复。
	LIMITEXCEEDED_CACERTNAMEREPEAT = "LimitExceeded.CACertNameRepeat"

	// 不支持私有证书操作。
	LIMITEXCEEDED_CACERTNOTSUPPORT = "LimitExceeded.CACertNotSupport"

	// CA证书重复。
	LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"

	// 设备数量超过限制。
	LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"

	// 固件数量超出限制。
	LIMITEXCEEDED_FIRMWAREEXCEEDLIMIT = "LimitExceeded.FirmwareExceedLimit"

	// 消息已经保存到离线队列。
	LIMITEXCEEDED_MESSAGESAVED = "LimitExceeded.MessageSaved"

	// qos为1的离线消息超过数量限制。
	LIMITEXCEEDED_OFFLINEMESSAGEEXCEEDLIMIT = "LimitExceeded.OfflineMessageExceedLimit"

	// 超过产品数量限制。
	LIMITEXCEEDED_PRODUCTEXCEEDLIMIT = "LimitExceeded.ProductExceedLimit"

	// Topic数量超出限制。
	LIMITEXCEEDED_TOPICPOLICYEXCEEDLIMIT = "LimitExceeded.TopicPolicyExceedLimit"

	// CA证书不存在。
	RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"

	// 批量创建设备任务不存在。
	RESOURCENOTFOUND_CREATEMULTIDEVICETASKNOTEXIST = "ResourceNotFound.CreateMultiDeviceTaskNotExist"

	// 设备固件升级任务不存在。
	RESOURCENOTFOUND_DEVICEFIRMWARETASKNOTEXIST = "ResourceNotFound.DeviceFirmwareTaskNotExist"

	// 设备无固件版本。
	RESOURCENOTFOUND_DEVICEHASNOFIRMWARE = "ResourceNotFound.DeviceHasNoFirmware"

	// 设备不存在。
	RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"

	// 设备资源不存在。
	RESOURCENOTFOUND_DEVICERESOURCENOTEXIST = "ResourceNotFound.DeviceResourceNotExist"

	// 设备影子不存在。
	RESOURCENOTFOUND_DEVICESHADOWNOTEXIST = "ResourceNotFound.DeviceShadowNotExist"

	// 固件不存在。
	RESOURCENOTFOUND_FIRMWARENOTEXIST = "ResourceNotFound.FirmwareNotExist"

	// 固件升级任务不存在。
	RESOURCENOTFOUND_FIRMWARETASKNOTEXIST = "ResourceNotFound.FirmwareTaskNotExist"

	// 产品不存在。
	RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"

	// 用户不存在此产品或设备。
	RESOURCENOTFOUND_PRODUCTORDEVICENOTEXIST = "ResourceNotFound.ProductOrDeviceNotExist"

	// 产品资源不存在。
	RESOURCENOTFOUND_PRODUCTRESOURCENOTEXIST = "ResourceNotFound.ProductResourceNotExist"

	// 资源文件不存在。
	RESOURCENOTFOUND_RESOURCEFILENOTEXIST = "ResourceNotFound.ResourceFileNotExist"

	// 任务不存在。
	RESOURCENOTFOUND_TASKNOTEXIST = "ResourceNotFound.TaskNotExist"

	// 物模型不存在。
	RESOURCENOTFOUND_THINGMODELNOTEXIST = "ResourceNotFound.ThingModelNotExist"

	// Topic不存在。
	RESOURCENOTFOUND_TOPICPOLICYNOTEXIST = "ResourceNotFound.TopicPolicyNotExist"

	// 规则不存在。
	RESOURCENOTFOUND_TOPICRULENOTEXIST = "ResourceNotFound.TopicRuleNotExist"

	// 该产品已存在TID申请，禁止删除。
	UNAUTHORIZEDOPERATION_DELETETIDFAIL = "UnauthorizedOperation.DeleteTidFail"

	// 该设备绑定了网关设备，无法删除。
	UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"

	// 设备未启用。
	UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"

	// 删除的产品下还包括未删除的设备。
	UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"

	// 该设备下仍有绑定的设备。
	UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"

	// 没有权限。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"

	// 该产品类型不能创建LoRa设备。
	UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"

	// NB-IoT产品不允许创建普通设备。
	UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"

	// 该产品类型只能创建LoRa设备。
	UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"

	// 产品禁用了该功能。
	UNAUTHORIZEDOPERATION_PRODUCTISFORBIDDEN = "UnauthorizedOperation.ProductIsForbidden"

	// 产品不支持密钥认证。
	UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"

	// 用户未通过实名认证。
	UNAUTHORIZEDOPERATION_USERNOTAUTHENTICAED = "UnauthorizedOperation.UserNotAuthenticaed"

	// 设备私钥已被获取。
	UNSUPPORTEDOPERATION_CLIENTCERTALREADYGOT = "UnsupportedOperation.ClientCertAlreadyGot"

	// 设备ota升级中。
	UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"

	// 网关产品下存在绑定的子产品，无法删除。
	UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"

	// 存在网关设备绑定当前产品，无法删除。
	UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"

	// 产品存在绑定的网关产品，无法删除。
	UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"

	// 产品为Suite token类型，无法创建新设备。
	UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"

	// 不支持的认证类型。
	UNSUPPORTEDOPERATION_WRONGPRODUCTAUTHTYPE = "UnsupportedOperation.WrongProductAuthType"
)
