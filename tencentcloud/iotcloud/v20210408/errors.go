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

	// 白名单校验未开启，用户不可创建设备，平台会根据设备认证时携带的设备名称自动创建设备。
	FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// CA验证证书不匹配。
	INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"

	// 格式错误，DefinedPsk需为Base64格式的字符串。
	INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"

	// 创建的设备名已存在。
	INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"

	// 产品类型不支持。
	INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"

	// 设备数量超过限制。
	LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"

	// 设备不存在。
	RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"

	// 产品不存在。
	RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"

	// 该设备绑定了网关设备，无法删除。
	UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"

	// 设备未启用。
	UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"

	// 删除的产品下还包括未删除的设备。
	UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"

	// 该设备下仍有绑定的设备。
	UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"

	// 该产品类型不能创建LoRa设备。
	UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"

	// NB-IoT产品不允许创建普通设备。
	UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"

	// 该产品类型只能创建LoRa设备。
	UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"

	// 产品不支持密钥认证。
	UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"

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
)
