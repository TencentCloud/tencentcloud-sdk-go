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

package v20190423

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ActivateDeviceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例类型
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 设备激活信息
	DeviceActivationDetails *DeviceActivationDetail `json:"DeviceActivationDetails,omitnil,omitempty" name:"DeviceActivationDetails"`

	// 已注册设备类型信息
	RegisteredDeviceType *RegisteredDeviceTypeInfo `json:"RegisteredDeviceType,omitnil,omitempty" name:"RegisteredDeviceType"`

	// 已注册设备通信类型信息
	RegisteredDeviceNetType *RegisteredDeviceNetTypeInfo `json:"RegisteredDeviceNetType,omitnil,omitempty" name:"RegisteredDeviceNetType"`
}

// Predefined struct for user
type ActivateTWeCallLicenseRequestParams struct {
	// TWecall类型：0-体验套餐；1-家庭安防场景； 2-穿戴类场景； 3-生活娱乐场景； 4-对讲及其它场景
	PkgType *int64 `json:"PkgType,omitnil,omitempty" name:"PkgType"`

	// 参数已弃用，不用传参
	//
	// Deprecated: MiniProgramAppId is deprecated.
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`

	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type ActivateTWeCallLicenseRequest struct {
	*tchttp.BaseRequest
	
	// TWecall类型：0-体验套餐；1-家庭安防场景； 2-穿戴类场景； 3-生活娱乐场景； 4-对讲及其它场景
	PkgType *int64 `json:"PkgType,omitnil,omitempty" name:"PkgType"`

	// 参数已弃用，不用传参
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`

	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *ActivateTWeCallLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateTWeCallLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PkgType")
	delete(f, "MiniProgramAppId")
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateTWeCallLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateTWeCallLicenseResponseParams struct {
	// 设备激活返回数据
	//
	// Deprecated: DeviceList is deprecated.
	DeviceList []*DeviceActiveResult `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 设备激活失败返回数据
	FailureList []*DeviceActiveResult `json:"FailureList,omitnil,omitempty" name:"FailureList"`

	// 设备激活成功返回数据
	SuccessList []*DeviceActiveResult `json:"SuccessList,omitnil,omitempty" name:"SuccessList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateTWeCallLicenseResponse struct {
	*tchttp.BaseResponse
	Response *ActivateTWeCallLicenseResponseParams `json:"Response"`
}

func (r *ActivateTWeCallLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateTWeCallLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppDeviceInfo struct {
	// 产品ID/设备名
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备别名
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// icon地址
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// 家庭ID
	FamilyId *string `json:"FamilyId,omitnil,omitempty" name:"FamilyId"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 设备类型
	DeviceType *int64 `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AuthMiniProgramAppInfo struct {
	// 小程序APPID
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 小程序名称
	MiniProgramName *string `json:"MiniProgramName,omitnil,omitempty" name:"MiniProgramName"`

	// 激活码数
	LicenseNum *int64 `json:"LicenseNum,omitnil,omitempty" name:"LicenseNum"`

	// 应用ID 
	IotAppId *string `json:"IotAppId,omitnil,omitempty" name:"IotAppId"`

	// 应用名称
	IotAppName *string `json:"IotAppName,omitnil,omitempty" name:"IotAppName"`
}

type BatchProductionInfo struct {
	// 量产ID
	BatchProductionId *string `json:"BatchProductionId,omitnil,omitempty" name:"BatchProductionId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 烧录方式
	BurnMethod *int64 `json:"BurnMethod,omitnil,omitempty" name:"BurnMethod"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

// Predefined struct for user
type BindCloudStorageUserRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type BindCloudStorageUserRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *BindCloudStorageUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudStorageUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindCloudStorageUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudStorageUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindCloudStorageUserResponse struct {
	*tchttp.BaseResponse
	Response *BindCloudStorageUserResponseParams `json:"Response"`
}

func (r *BindCloudStorageUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudStorageUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindDeviceInfo struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

// Predefined struct for user
type BindDevicesRequestParams struct {
	// 网关设备的产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名。
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 被绑定设备的产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 被绑定的多个设备名。
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`
}

type BindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名。
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 被绑定设备的产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 被绑定的多个设备名。
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`
}

func (r *BindDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDevicesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindDevicesResponse struct {
	*tchttp.BaseResponse
	Response *BindDevicesResponseParams `json:"Response"`
}

func (r *BindDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindProductInfo struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品所属项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 物模型类型。
	DataProtocol *int64 `json:"DataProtocol,omitnil,omitempty" name:"DataProtocol"`

	// 产品分组模板ID
	CategoryId *int64 `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 产品类型
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 连接类型
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 状态
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`

	// 产品拥有者名称
	ProductOwnerName *string `json:"ProductOwnerName,omitnil,omitempty" name:"ProductOwnerName"`
}

// Predefined struct for user
type BindProductsRequestParams struct {
	// 网关产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 待绑定的子产品ID数组。
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`
}

type BindProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID。
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 待绑定的子产品ID数组。
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`
}

func (r *BindProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "ProductIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindProductsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindProductsResponse struct {
	*tchttp.BaseResponse
	Response *BindProductsResponseParams `json:"Response"`
}

func (r *BindProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionAsyncRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitnil,omitempty" name:"InputParams"`
}

type CallDeviceActionAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitnil,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ActionId")
	delete(f, "InputParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallDeviceActionAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionAsyncResponseParams struct {
	// 调用Id
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 异步调用状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CallDeviceActionAsyncResponse struct {
	*tchttp.BaseResponse
	Response *CallDeviceActionAsyncResponseParams `json:"Response"`
}

func (r *CallDeviceActionAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionSyncRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitnil,omitempty" name:"InputParams"`
}

type CallDeviceActionSyncRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitnil,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ActionId")
	delete(f, "InputParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallDeviceActionSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionSyncResponseParams struct {
	// 调用Id
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 输出参数，取值设备端上报$thing/up/action method为action_reply 的 response字段，物模型协议参考https://cloud.tencent.com/document/product/1081/34916#.E8.AE.BE.E5.A4.87.E8.A1.8C.E4.B8.BA.E8.B0.83.E7.94.A8
	OutputParams *string `json:"OutputParams,omitnil,omitempty" name:"OutputParams"`

	// 返回状态，取值设备端上报$thing/up/action	method为action_reply 的 status字段，如果不包含status字段，则取默认值，空字符串，物模型协议参考https://cloud.tencent.com/document/product/1081/34916#.E8.AE.BE.E5.A4.87.E8.A1.8C.E4.B8.BA.E8.B0.83.E7.94.A8
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CallDeviceActionSyncResponse struct {
	*tchttp.BaseResponse
	Response *CallDeviceActionSyncResponseParams `json:"Response"`
}

func (r *CallDeviceActionSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CamTag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type CancelAssignTWeCallLicenseRequestParams struct {
	// 订单号
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`
}

type CancelAssignTWeCallLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 订单号
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`
}

func (r *CancelAssignTWeCallLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAssignTWeCallLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PkgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAssignTWeCallLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAssignTWeCallLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelAssignTWeCallLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CancelAssignTWeCallLicenseResponseParams `json:"Response"`
}

func (r *CancelAssignTWeCallLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAssignTWeCallLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeP2PRouteRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// P2P线路
	RouteId *uint64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

type ChangeP2PRouteRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// P2P线路
	RouteId *uint64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`
}

func (r *ChangeP2PRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeP2PRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "RouteId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeP2PRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeP2PRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeP2PRouteResponse struct {
	*tchttp.BaseResponse
	Response *ChangeP2PRouteResponseParams `json:"Response"`
}

func (r *ChangeP2PRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeP2PRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckFirmwareUpdateRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type CheckFirmwareUpdateRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *CheckFirmwareUpdateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFirmwareUpdateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckFirmwareUpdateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckFirmwareUpdateResponseParams struct {
	// 设备当前固件版本。
	CurrentVersion *string `json:"CurrentVersion,omitnil,omitempty" name:"CurrentVersion"`

	// 固件可升级版本。
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckFirmwareUpdateResponse struct {
	*tchttp.BaseResponse
	Response *CheckFirmwareUpdateResponseParams `json:"Response"`
}

func (r *CheckFirmwareUpdateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFirmwareUpdateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudStorageAIServiceTask struct {
	// 云存 AI 服务任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 云存 AI 服务类型。可能取值：
	// 
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 对应云存视频的起始时间（秒级 UNIX 时间戳）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 对应云存视频的起始时间（毫秒级 UNIX 时间戳）
	StartTimeMs *int64 `json:"StartTimeMs,omitnil,omitempty" name:"StartTimeMs"`

	// 对应云存视频的结束时间（秒级 UNIX 时间戳）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 对应云存视频的结束时间（毫秒级 UNIX 时间戳）
	EndTimeMs *int64 `json:"EndTimeMs,omitnil,omitempty" name:"EndTimeMs"`

	// 任务状态（1：失败；2：成功但结果为空；3：成功且结果非空；4：执行中）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 任务输出文件列表
	Files []*string `json:"Files,omitnil,omitempty" name:"Files"`

	// 任务输出文件信息列表
	FilesInfo []*CloudStorageAIServiceTaskFileInfo `json:"FilesInfo,omitnil,omitempty" name:"FilesInfo"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`
}

type CloudStorageAIServiceTaskFileInfo struct {
	// 文件名称（含扩展名）
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小（单位：bytes）
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文件下载 URL
	DownloadURL *string `json:"DownloadURL,omitnil,omitempty" name:"DownloadURL"`

	// 文件的 MIME Type
	MimeType *string `json:"MimeType,omitnil,omitempty" name:"MimeType"`

	// 视频文件元数据（仅当文件为视频类型时包含该字段）
	VideoMetaInfo *CloudStorageAIServiceTaskVideoMetaInfo `json:"VideoMetaInfo,omitnil,omitempty" name:"VideoMetaInfo"`

	// 文件标签
	Labels []*CloudStorageAIServiceTaskFileLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type CloudStorageAIServiceTaskFileLabel struct {
	// key1
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// value1
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type CloudStorageAIServiceTaskVideoMetaInfo struct {
	// 视频对应的缩略图的文件名称（含扩展名）
	ThumbnailFileName *string `json:"ThumbnailFileName,omitnil,omitempty" name:"ThumbnailFileName"`

	// 视频时长（单位：毫秒）
	DurationMilliSeconds *int64 `json:"DurationMilliSeconds,omitnil,omitempty" name:"DurationMilliSeconds"`
}

type CloudStorageEvent struct {
	// 事件起始时间（Unix 时间戳，秒级
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件结束时间（Unix 时间戳，秒级
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件缩略图
	Thumbnail *string `json:"Thumbnail,omitnil,omitempty" name:"Thumbnail"`

	// 事件ID
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 事件录像上传状态，Finished: 全部上传成功 Partial: 部分上传成功 Failed: 上传失败	
	UploadStatus *string `json:"UploadStatus,omitnil,omitempty" name:"UploadStatus"`

	// 事件自定义数据	
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type CloudStorageEventWithAITasks struct {
	// 事件起始时间（Unix 时间戳，秒级
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件结束时间（Unix 时间戳，秒级
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 事件缩略图
	Thumbnail *string `json:"Thumbnail,omitnil,omitempty" name:"Thumbnail"`

	// 事件ID
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 事件录像上传状态，Finished: 全部上传成功 Partial: 部分上传成功 Failed: 上传失败	
	UploadStatus *string `json:"UploadStatus,omitnil,omitempty" name:"UploadStatus"`

	// 事件自定义数据	
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 事件关联的云存 AI 任务列表
	AITasks []*CloudStorageAIServiceTask `json:"AITasks,omitnil,omitempty" name:"AITasks"`
}

type CloudStoragePackageInfo struct {
	// 套餐包id
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 套餐包名字
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 套餐包数量
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 已使用数量
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`
}

type CloudStorageTimeData struct {
	// 云存时间轴信息列表
	TimeList []*CloudStorageTimeInfo `json:"TimeList,omitnil,omitempty" name:"TimeList"`

	// 播放地址
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`
}

type CloudStorageTimeInfo struct {
	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type CloudStorageUserInfo struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

// Predefined struct for user
type ControlDeviceDataRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 请求类型 , 不填该参数或者 desired 表示下发属性给设备,  reported 表示模拟设备上报属性
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName , 通常情况不需要填写
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 上报数据UNIX时间戳(毫秒), 仅对Method:reported有效
	DataTimestamp *int64 `json:"DataTimestamp,omitnil,omitempty" name:"DataTimestamp"`
}

type ControlDeviceDataRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 请求类型 , 不填该参数或者 desired 表示下发属性给设备,  reported 表示模拟设备上报属性
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName , 通常情况不需要填写
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 上报数据UNIX时间戳(毫秒), 仅对Method:reported有效
	DataTimestamp *int64 `json:"DataTimestamp,omitnil,omitempty" name:"DataTimestamp"`
}

func (r *ControlDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Data")
	delete(f, "Method")
	delete(f, "DeviceId")
	delete(f, "DataTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDeviceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceDataResponseParams struct {
	// 返回信息
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// JSON字符串， 返回下发控制的结果信息, 
	// Sent = 1 表示设备已经在线并且订阅了控制下发的mqtt topic.
	// pushResult 是表示发送结果，其中 0 表示成功， 23101 表示设备未在线或没有订阅相关的 MQTT Topic。
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *ControlDeviceDataResponseParams `json:"Response"`
}

func (r *ControlDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountDataInfo struct {
	// 视频上报异常次数
	VideoExceptionNum *uint64 `json:"VideoExceptionNum,omitnil,omitempty" name:"VideoExceptionNum"`

	// 视频上报成功次数
	VideoSuccessNum *uint64 `json:"VideoSuccessNum,omitnil,omitempty" name:"VideoSuccessNum"`

	// 视频上报成功率
	VideoSuccessRate *string `json:"VideoSuccessRate,omitnil,omitempty" name:"VideoSuccessRate"`

	// 事件上报异常次数
	EventExceptionNum *uint64 `json:"EventExceptionNum,omitnil,omitempty" name:"EventExceptionNum"`

	// 事件上报成功次数
	EventSuccessNum *uint64 `json:"EventSuccessNum,omitnil,omitempty" name:"EventSuccessNum"`

	// 事件上报成功率
	EventSuccessRate *string `json:"EventSuccessRate,omitnil,omitempty" name:"EventSuccessRate"`
}

// Predefined struct for user
type CreateBatchProductionRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 烧录方式，0为直接烧录，1为动态注册。
	BurnMethod *int64 `json:"BurnMethod,omitnil,omitempty" name:"BurnMethod"`

	// 生成方式，0为系统生成，1为文件上传。
	GenerationMethod *int64 `json:"GenerationMethod,omitnil,omitempty" name:"GenerationMethod"`

	// 文件上传URL，用于文件上传时填写。
	UploadUrl *string `json:"UploadUrl,omitnil,omitempty" name:"UploadUrl"`

	// 量产数量，用于系统生成时填写。
	BatchCnt *int64 `json:"BatchCnt,omitnil,omitempty" name:"BatchCnt"`

	// 是否生成二维码,0为不生成，1为生成。
	GenerationQRCode *int64 `json:"GenerationQRCode,omitnil,omitempty" name:"GenerationQRCode"`
}

type CreateBatchProductionRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 烧录方式，0为直接烧录，1为动态注册。
	BurnMethod *int64 `json:"BurnMethod,omitnil,omitempty" name:"BurnMethod"`

	// 生成方式，0为系统生成，1为文件上传。
	GenerationMethod *int64 `json:"GenerationMethod,omitnil,omitempty" name:"GenerationMethod"`

	// 文件上传URL，用于文件上传时填写。
	UploadUrl *string `json:"UploadUrl,omitnil,omitempty" name:"UploadUrl"`

	// 量产数量，用于系统生成时填写。
	BatchCnt *int64 `json:"BatchCnt,omitnil,omitempty" name:"BatchCnt"`

	// 是否生成二维码,0为不生成，1为生成。
	GenerationQRCode *int64 `json:"GenerationQRCode,omitnil,omitempty" name:"GenerationQRCode"`
}

func (r *CreateBatchProductionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchProductionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProductId")
	delete(f, "BurnMethod")
	delete(f, "GenerationMethod")
	delete(f, "UploadUrl")
	delete(f, "BatchCnt")
	delete(f, "GenerationQRCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchProductionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchProductionResponseParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 量产id
	BatchProductionId *string `json:"BatchProductionId,omitnil,omitempty" name:"BatchProductionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBatchProductionResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchProductionResponseParams `json:"Response"`
}

func (r *CreateBatchProductionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchProductionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudStorageAIServiceRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 套餐 ID。可选值：
	// 
	// - `1m_low_od`：低功耗目标检测月套餐
	// - `1y_low_od`：低功耗目标检测年套餐
	// - `1m_ev_od`：事件目标检测月套餐
	// - `1y_ev_od`：事件目标检测年套餐
	// - `1m_ft_od`：全时目标检测月套餐
	// - `1y_ft_od`：全时目标检测年套餐
	// - `1m_low_hl`：低功耗视频浓缩月套餐
	// - `1y_low_hl`：低功耗视频浓缩年套餐
	// - `1m_ev_hl`：事件视频浓缩月套餐
	// - `1y_ev_hl`：事件视频浓缩年套餐
	// - `1m_ft_hl`：全时视频浓缩月套餐
	// - `1y_ft_hl`：全时视频浓缩年套餐
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 订单 ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type CreateCloudStorageAIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 套餐 ID。可选值：
	// 
	// - `1m_low_od`：低功耗目标检测月套餐
	// - `1y_low_od`：低功耗目标检测年套餐
	// - `1m_ev_od`：事件目标检测月套餐
	// - `1y_ev_od`：事件目标检测年套餐
	// - `1m_ft_od`：全时目标检测月套餐
	// - `1y_ft_od`：全时目标检测年套餐
	// - `1m_low_hl`：低功耗视频浓缩月套餐
	// - `1y_low_hl`：低功耗视频浓缩年套餐
	// - `1m_ev_hl`：事件视频浓缩月套餐
	// - `1y_ev_hl`：事件视频浓缩年套餐
	// - `1m_ft_hl`：全时视频浓缩月套餐
	// - `1y_ft_hl`：全时视频浓缩年套餐
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 订单 ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *CreateCloudStorageAIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudStorageAIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "PackageId")
	delete(f, "ChannelId")
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudStorageAIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudStorageAIServiceResponseParams struct {
	// 订单 ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudStorageAIServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudStorageAIServiceResponseParams `json:"Response"`
}

func (r *CreateCloudStorageAIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudStorageAIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudStorageAIServiceTaskRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 待分析云存的起始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 待分析云存的结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`
}

type CreateCloudStorageAIServiceTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 待分析云存的起始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 待分析云存的结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`
}

func (r *CreateCloudStorageAIServiceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudStorageAIServiceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ServiceType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ChannelId")
	delete(f, "Config")
	delete(f, "ROI")
	delete(f, "VideoURLs")
	delete(f, "CustomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudStorageAIServiceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudStorageAIServiceTaskResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudStorageAIServiceTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudStorageAIServiceTaskResponseParams `json:"Response"`
}

func (r *CreateCloudStorageAIServiceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudStorageAIServiceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceChannelRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type CreateDeviceChannelRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *CreateDeviceChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceChannelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeviceChannelResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceChannelResponseParams `json:"Response"`
}

func (r *CreateDeviceChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// LoRaWAN 设备地址
	DevAddr *string `json:"DevAddr,omitnil,omitempty" name:"DevAddr"`

	// LoRaWAN 应用密钥
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// LoRaWAN 设备唯一标识
	DevEUI *string `json:"DevEUI,omitnil,omitempty" name:"DevEUI"`

	// LoRaWAN 应用会话密钥
	AppSKey *string `json:"AppSKey,omitnil,omitempty" name:"AppSKey"`

	// LoRaWAN 网络会话密钥
	NwkSKey *string `json:"NwkSKey,omitnil,omitempty" name:"NwkSKey"`

	// 手动指定设备的PSK密钥
	DefinedPsk *string `json:"DefinedPsk,omitnil,omitempty" name:"DefinedPsk"`
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// LoRaWAN 设备地址
	DevAddr *string `json:"DevAddr,omitnil,omitempty" name:"DevAddr"`

	// LoRaWAN 应用密钥
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// LoRaWAN 设备唯一标识
	DevEUI *string `json:"DevEUI,omitnil,omitempty" name:"DevEUI"`

	// LoRaWAN 应用会话密钥
	AppSKey *string `json:"AppSKey,omitnil,omitempty" name:"AppSKey"`

	// LoRaWAN 网络会话密钥
	NwkSKey *string `json:"NwkSKey,omitnil,omitempty" name:"NwkSKey"`

	// 手动指定设备的PSK密钥
	DefinedPsk *string `json:"DefinedPsk,omitnil,omitempty" name:"DefinedPsk"`
}

func (r *CreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "DevAddr")
	delete(f, "AppKey")
	delete(f, "DevEUI")
	delete(f, "AppSKey")
	delete(f, "NwkSKey")
	delete(f, "DefinedPsk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceResponseParams struct {
	// 设备参数描述。
	Data *DeviceData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceResponseParams `json:"Response"`
}

func (r *CreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExternalSourceAIServiceTaskRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`
}

type CreateExternalSourceAIServiceTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`
}

func (r *CreateExternalSourceAIServiceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExternalSourceAIServiceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ServiceType")
	delete(f, "VideoURLs")
	delete(f, "CustomId")
	delete(f, "Config")
	delete(f, "ROI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExternalSourceAIServiceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExternalSourceAIServiceTaskResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExternalSourceAIServiceTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateExternalSourceAIServiceTaskResponseParams `json:"Response"`
}

func (r *CreateExternalSourceAIServiceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExternalSourceAIServiceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFenceBindRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type CreateFenceBindRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitnil,omitempty" name:"Items"`
}

func (r *CreateFenceBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFenceBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFenceBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFenceBindResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFenceBindResponse struct {
	*tchttp.BaseResponse
	Response *CreateFenceBindResponseParams `json:"Response"`
}

func (r *CreateFenceBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFenceBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFreeCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存套餐ID：
	// lye1w3d：低功耗事件3天周套餐。
	// ye1w3d：事件3天周套餐
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 如果当前设备已开启云存套餐，Override=1会使用新套餐覆盖原有套餐。不传此参数则默认为0。
	Override *uint64 `json:"Override,omitnil,omitempty" name:"Override"`

	// 套餐列表顺序：PackageQueue=front会立即使用新购买的套餐，新购套餐结束后，列表中下一个未过期的套餐继续生效；PackageQueue=end会等设备当前所有已购买套餐过期后才会生效新购套餐。与Override参数不能同时使用。
	PackageQueue *string `json:"PackageQueue,omitnil,omitempty" name:"PackageQueue"`

	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 云存视频存储区域，国内默认为ap-guangzhou。海外默认为东南亚ap-singapore，可选美东na-ashburn、欧洲eu-frankfurt。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`
}

type CreateFreeCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存套餐ID：
	// lye1w3d：低功耗事件3天周套餐。
	// ye1w3d：事件3天周套餐
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 如果当前设备已开启云存套餐，Override=1会使用新套餐覆盖原有套餐。不传此参数则默认为0。
	Override *uint64 `json:"Override,omitnil,omitempty" name:"Override"`

	// 套餐列表顺序：PackageQueue=front会立即使用新购买的套餐，新购套餐结束后，列表中下一个未过期的套餐继续生效；PackageQueue=end会等设备当前所有已购买套餐过期后才会生效新购套餐。与Override参数不能同时使用。
	PackageQueue *string `json:"PackageQueue,omitnil,omitempty" name:"PackageQueue"`

	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 云存视频存储区域，国内默认为ap-guangzhou。海外默认为东南亚ap-singapore，可选美东na-ashburn、欧洲eu-frankfurt。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`
}

func (r *CreateFreeCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFreeCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "PackageId")
	delete(f, "Override")
	delete(f, "PackageQueue")
	delete(f, "OrderId")
	delete(f, "ChannelId")
	delete(f, "StorageRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFreeCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFreeCloudStorageResponseParams struct {
	// 订单金额，单位为分
	Price *uint64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 支付金额，单位为分
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFreeCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *CreateFreeCloudStorageResponseParams `json:"Response"`
}

func (r *CreateFreeCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFreeCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotVideoCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存套餐ID：
	// yc1m3d ： 全时3天存储月套餐。
	// yc1m7d ： 全时7天存储月套餐。
	// yc1m30d ：全时30天存储月套餐。
	// yc1y3d ：全时3天存储年套餐。
	// yc1y7d ：全时7天存储年套餐。
	// yc1y30d ：全时30天存储年套餐。
	// ye1m3d ：事件3天存储月套餐。
	// ye1m7d ：事件7天存储月套餐。
	// ye1m30d ：事件30天存储月套餐 。
	// ye1y3d ：事件3天存储年套餐。
	// ye1y7d ：事件7天存储年套餐。
	// ye1y30d ：事件30天存储年套餐。
	// yc1w7d : 全时7天存储周套餐。
	// ye1w7d : 事件7天存储周套餐。
	// lye1m3d：低功耗事件3天月套餐。
	// lye1m7d：低功耗事件7天月套餐。
	// lye1m30d：低功耗事件30天月套餐。
	// lye1y3d：低功耗事件3天年套餐。
	// lye1y7d：低功耗事件7天年套餐。
	// lye1y30d：低功耗事件30天年套餐。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 如果当前设备已开启云存套餐，Override=1会使用新套餐覆盖原有套餐。不传此参数则默认为0。
	Override *uint64 `json:"Override,omitnil,omitempty" name:"Override"`

	// 套餐列表顺序：PackageQueue=front会立即使用新购买的套餐，新购套餐结束后，列表中下一个未过期的套餐继续生效；PackageQueue=end会等设备当前所有已购买套餐过期后才会生效新购套餐。与Override参数不能同时使用。
	PackageQueue *string `json:"PackageQueue,omitnil,omitempty" name:"PackageQueue"`

	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 云存视频存储区域，国内默认为ap-guangzhou。海外默认为东南亚ap-singapore，可选美东na-ashburn、欧洲eu-frankfurt。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`
}

type CreateIotVideoCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存套餐ID：
	// yc1m3d ： 全时3天存储月套餐。
	// yc1m7d ： 全时7天存储月套餐。
	// yc1m30d ：全时30天存储月套餐。
	// yc1y3d ：全时3天存储年套餐。
	// yc1y7d ：全时7天存储年套餐。
	// yc1y30d ：全时30天存储年套餐。
	// ye1m3d ：事件3天存储月套餐。
	// ye1m7d ：事件7天存储月套餐。
	// ye1m30d ：事件30天存储月套餐 。
	// ye1y3d ：事件3天存储年套餐。
	// ye1y7d ：事件7天存储年套餐。
	// ye1y30d ：事件30天存储年套餐。
	// yc1w7d : 全时7天存储周套餐。
	// ye1w7d : 事件7天存储周套餐。
	// lye1m3d：低功耗事件3天月套餐。
	// lye1m7d：低功耗事件7天月套餐。
	// lye1m30d：低功耗事件30天月套餐。
	// lye1y3d：低功耗事件3天年套餐。
	// lye1y7d：低功耗事件7天年套餐。
	// lye1y30d：低功耗事件30天年套餐。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 如果当前设备已开启云存套餐，Override=1会使用新套餐覆盖原有套餐。不传此参数则默认为0。
	Override *uint64 `json:"Override,omitnil,omitempty" name:"Override"`

	// 套餐列表顺序：PackageQueue=front会立即使用新购买的套餐，新购套餐结束后，列表中下一个未过期的套餐继续生效；PackageQueue=end会等设备当前所有已购买套餐过期后才会生效新购套餐。与Override参数不能同时使用。
	PackageQueue *string `json:"PackageQueue,omitnil,omitempty" name:"PackageQueue"`

	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 云存视频存储区域，国内默认为ap-guangzhou。海外默认为东南亚ap-singapore，可选美东na-ashburn、欧洲eu-frankfurt。
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`
}

func (r *CreateIotVideoCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotVideoCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "PackageId")
	delete(f, "Override")
	delete(f, "PackageQueue")
	delete(f, "OrderId")
	delete(f, "ChannelId")
	delete(f, "StorageRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIotVideoCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIotVideoCloudStorageResponseParams struct {
	// 订单金额，单位为分
	Price *uint64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 支付金额，单位为分
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIotVideoCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *CreateIotVideoCloudStorageResponseParams `json:"Response"`
}

func (r *CreateIotVideoCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIotVideoCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaFrequencyRequestParams struct {
	// 频点配置名称
	FreqName *string `json:"FreqName,omitnil,omitempty" name:"FreqName"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitnil,omitempty" name:"ChannelsDataUp"`

	// 数据下行RX1信道
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitnil,omitempty" name:"ChannelsDataRX1"`

	// 数据下行RX2信道
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitnil,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitnil,omitempty" name:"ChannelsJoinUp"`

	// 入网下行RX1信道
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitnil,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行RX2信道
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitnil,omitempty" name:"ChannelsJoinRX2"`

	// 频点配置描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点配置名称
	FreqName *string `json:"FreqName,omitnil,omitempty" name:"FreqName"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitnil,omitempty" name:"ChannelsDataUp"`

	// 数据下行RX1信道
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitnil,omitempty" name:"ChannelsDataRX1"`

	// 数据下行RX2信道
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitnil,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitnil,omitempty" name:"ChannelsJoinUp"`

	// 入网下行RX1信道
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitnil,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行RX2信道
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitnil,omitempty" name:"ChannelsJoinRX2"`

	// 频点配置描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqName")
	delete(f, "ChannelsDataUp")
	delete(f, "ChannelsDataRX1")
	delete(f, "ChannelsDataRX2")
	delete(f, "ChannelsJoinUp")
	delete(f, "ChannelsJoinRX1")
	delete(f, "ChannelsJoinRX2")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaFrequencyResponseParams struct {
	// LoRa频点信息
	Data *LoRaFrequencyEntry `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoRaFrequencyResponseParams `json:"Response"`
}

func (r *CreateLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaGatewayRequestParams struct {
	// LoRa 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 详情描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 位置信息
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitnil,omitempty" name:"PositionDetails"`

	// 是否公开
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitnil,omitempty" name:"FrequencyId"`
}

type CreateLoRaGatewayRequest struct {
	*tchttp.BaseRequest
	
	// LoRa 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 详情描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 位置信息
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitnil,omitempty" name:"PositionDetails"`

	// 是否公开
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitnil,omitempty" name:"FrequencyId"`
}

func (r *CreateLoRaGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Location")
	delete(f, "Position")
	delete(f, "PositionDetails")
	delete(f, "IsPublic")
	delete(f, "FrequencyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoRaGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoRaGatewayResponseParams struct {
	// LoRa 网关信息
	Gateway *LoRaGatewayItem `json:"Gateway,omitnil,omitempty" name:"Gateway"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoRaGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoRaGatewayResponseParams `json:"Response"`
}

func (r *CreateLoRaGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoRaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionFenceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 围栏名称
	FenceName *string `json:"FenceName,omitnil,omitempty" name:"FenceName"`

	// 围栏区域信息，采用 GeoJSON 格式
	FenceArea *string `json:"FenceArea,omitnil,omitempty" name:"FenceArea"`

	// 围栏描述
	FenceDesc *string `json:"FenceDesc,omitnil,omitempty" name:"FenceDesc"`
}

type CreatePositionFenceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 围栏名称
	FenceName *string `json:"FenceName,omitnil,omitempty" name:"FenceName"`

	// 围栏区域信息，采用 GeoJSON 格式
	FenceArea *string `json:"FenceArea,omitnil,omitempty" name:"FenceArea"`

	// 围栏描述
	FenceDesc *string `json:"FenceDesc,omitnil,omitempty" name:"FenceDesc"`
}

func (r *CreatePositionFenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionFenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FenceName")
	delete(f, "FenceArea")
	delete(f, "FenceDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePositionFenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionFenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePositionFenceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePositionFenceResponseParams `json:"Response"`
}

func (r *CreatePositionFenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionFenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionSpaceRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 空间名称
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 授权类型，0：只读 1：读写
	AuthorizeType *int64 `json:"AuthorizeType,omitnil,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitnil,omitempty" name:"ProductIdList"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`
}

type CreatePositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 空间名称
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 授权类型，0：只读 1：读写
	AuthorizeType *int64 `json:"AuthorizeType,omitnil,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitnil,omitempty" name:"ProductIdList"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`
}

func (r *CreatePositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SpaceName")
	delete(f, "AuthorizeType")
	delete(f, "ProductIdList")
	delete(f, "Description")
	delete(f, "Icon")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePositionSpaceResponseParams struct {
	// 空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *CreatePositionSpaceResponseParams `json:"Response"`
}

func (r *CreatePositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitnil,omitempty" name:"ProjectDesc"`

	// 实例ID，不带实例ID，默认为公共实例
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitnil,omitempty" name:"ProjectDesc"`

	// 实例ID，不带实例ID，默认为公共实例
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectName")
	delete(f, "ProjectDesc")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 返回信息
	Project *ProjectEntry `json:"Project,omitnil,omitempty" name:"Project"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStudioProductRequestParams struct {
	// 产品名称，名称不能和已经存在的产品名称重复。命名规则：[a-zA-Z0-9:_-]{1,32}
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品分组模板ID , ( 自定义模板填写1 , 控制台调用会使用预置的其他ID)
	CategoryId *int64 `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 产品类型 填写 ( 0 普通产品 ， 5 网关产品)
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 加密类型 ，1表示证书认证，2表示密钥认证，21表示TID认证-SE方式，22表示TID认证-软加固方式
	EncryptionType *string `json:"EncryptionType,omitnil,omitempty" name:"EncryptionType"`

	// 连接类型 可以填写 wifi、wifi-ble、cellular、5g、lorawan、ble、ethernet、wifi-ethernet、else、sub_zigbee、sub_ble、sub_433mhz、sub_else、sub_blemesh
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 数据协议 (1 使用物模型 2 为自定义)
	DataProtocol *int64 `json:"DataProtocol,omitnil,omitempty" name:"DataProtocol"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitnil,omitempty" name:"ProductDesc"`

	// 产品的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 平均传输速率
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 期限
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

type CreateStudioProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品名称，名称不能和已经存在的产品名称重复。命名规则：[a-zA-Z0-9:_-]{1,32}
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品分组模板ID , ( 自定义模板填写1 , 控制台调用会使用预置的其他ID)
	CategoryId *int64 `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 产品类型 填写 ( 0 普通产品 ， 5 网关产品)
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 加密类型 ，1表示证书认证，2表示密钥认证，21表示TID认证-SE方式，22表示TID认证-软加固方式
	EncryptionType *string `json:"EncryptionType,omitnil,omitempty" name:"EncryptionType"`

	// 连接类型 可以填写 wifi、wifi-ble、cellular、5g、lorawan、ble、ethernet、wifi-ethernet、else、sub_zigbee、sub_ble、sub_433mhz、sub_else、sub_blemesh
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 数据协议 (1 使用物模型 2 为自定义)
	DataProtocol *int64 `json:"DataProtocol,omitnil,omitempty" name:"DataProtocol"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitnil,omitempty" name:"ProductDesc"`

	// 产品的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 平均传输速率
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 期限
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *CreateStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "CategoryId")
	delete(f, "ProductType")
	delete(f, "EncryptionType")
	delete(f, "NetType")
	delete(f, "DataProtocol")
	delete(f, "ProductDesc")
	delete(f, "ProjectId")
	delete(f, "Rate")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStudioProductResponseParams struct {
	// 产品描述
	Product *ProductEntry `json:"Product,omitnil,omitempty" name:"Product"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *CreateStudioProductResponseParams `json:"Response"`
}

func (r *CreateStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTRTCSignaturesWithRoomIdRequestParams struct {
	// TRTC进房间的用户名称数组，数组元素不可重复，最长不超过 10 个。
	TRTCUserIds []*string `json:"TRTCUserIds,omitnil,omitempty" name:"TRTCUserIds"`

	// 房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type CreateTRTCSignaturesWithRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// TRTC进房间的用户名称数组，数组元素不可重复，最长不超过 10 个。
	TRTCUserIds []*string `json:"TRTCUserIds,omitnil,omitempty" name:"TRTCUserIds"`

	// 房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *CreateTRTCSignaturesWithRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTRTCSignaturesWithRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TRTCUserIds")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTRTCSignaturesWithRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTRTCSignaturesWithRoomIdResponseParams struct {
	// 返回参数数组
	TRTCParamList []*TRTCParams `json:"TRTCParamList,omitnil,omitempty" name:"TRTCParamList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTRTCSignaturesWithRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *CreateTRTCSignaturesWithRoomIdResponseParams `json:"Response"`
}

func (r *CreateTRTCSignaturesWithRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTRTCSignaturesWithRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTWeSeeRecognitionTaskRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 输入视频 / 图片的 URL
	InputURL *string `json:"InputURL,omitnil,omitempty" name:"InputURL"`

	// 自定义事件 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 是否保存该事件使其可被搜索
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 事件起始时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	StartTimeMs *uint64 `json:"StartTimeMs,omitnil,omitempty" name:"StartTimeMs"`

	// 事件结束时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	EndTimeMs *uint64 `json:"EndTimeMs,omitnil,omitempty" name:"EndTimeMs"`

	// 算法配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否自定义设备，为 true 时不检查设备存在性，默认为 false
	IsCustomDevice *bool `json:"IsCustomDevice,omitnil,omitempty" name:"IsCustomDevice"`

	// 输入类型。可选值：
	// 
	// - `video`：视频（默认值）
	// - `image`：图片
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 摘要服务质量。可选值：
	// 
	// - `minutely`：分钟级（默认值）
	// - `immediate`：立即
	SummaryQOS *string `json:"SummaryQOS,omitnil,omitempty" name:"SummaryQOS"`
}

type CreateTWeSeeRecognitionTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 输入视频 / 图片的 URL
	InputURL *string `json:"InputURL,omitnil,omitempty" name:"InputURL"`

	// 自定义事件 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 是否保存该事件使其可被搜索
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 事件起始时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	StartTimeMs *uint64 `json:"StartTimeMs,omitnil,omitempty" name:"StartTimeMs"`

	// 事件结束时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	EndTimeMs *uint64 `json:"EndTimeMs,omitnil,omitempty" name:"EndTimeMs"`

	// 算法配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否自定义设备，为 true 时不检查设备存在性，默认为 false
	IsCustomDevice *bool `json:"IsCustomDevice,omitnil,omitempty" name:"IsCustomDevice"`

	// 输入类型。可选值：
	// 
	// - `video`：视频（默认值）
	// - `image`：图片
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 摘要服务质量。可选值：
	// 
	// - `minutely`：分钟级（默认值）
	// - `immediate`：立即
	SummaryQOS *string `json:"SummaryQOS,omitnil,omitempty" name:"SummaryQOS"`
}

func (r *CreateTWeSeeRecognitionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTWeSeeRecognitionTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "InputURL")
	delete(f, "CustomId")
	delete(f, "EnableSearch")
	delete(f, "StartTimeMs")
	delete(f, "EndTimeMs")
	delete(f, "Config")
	delete(f, "IsCustomDevice")
	delete(f, "InputType")
	delete(f, "SummaryQOS")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTWeSeeRecognitionTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTWeSeeRecognitionTaskResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTWeSeeRecognitionTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTWeSeeRecognitionTaskResponseParams `json:"Response"`
}

func (r *CreateTWeSeeRecognitionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTWeSeeRecognitionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic权限，1发布，2订阅，3订阅和发布
	Privilege *uint64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`
}

type CreateTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic权限，1发布，2订阅，3订阅和发布
	Privilege *uint64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`
}

func (r *CreateTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	delete(f, "Privilege")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicPolicyResponseParams `json:"Response"`
}

func (r *CreateTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则内容
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitnil,omitempty" name:"TopicRulePayload"`
}

type CreateTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则内容
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitnil,omitempty" name:"TopicRulePayload"`
}

func (r *CreateTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "TopicRulePayload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicRuleResponseParams `json:"Response"`
}

func (r *CreateTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudStorageEventRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 事件id
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 开始时间，unix时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DeleteCloudStorageEventRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 事件id
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 开始时间，unix时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DeleteCloudStorageEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudStorageEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "EventId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudStorageEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudStorageEventResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudStorageEventResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudStorageEventResponseParams `json:"Response"`
}

func (r *DeleteCloudStorageEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudStorageEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 是否删除绑定设备
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 是否删除绑定设备
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ForceDelete")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResponseParams struct {
	// 删除的结果代码
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 删除的结果信息
	ResultMessage *string `json:"ResultMessage,omitnil,omitempty" name:"ResultMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceResponseParams `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDevicesRequestParams struct {
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitnil,omitempty" name:"DevicesItems"`
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitnil,omitempty" name:"DevicesItems"`
}

func (r *DeleteDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DevicesItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDevicesResponseParams struct {
	// 删除的结果代码
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 删除的结果信息
	ResultMessage *string `json:"ResultMessage,omitnil,omitempty" name:"ResultMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDevicesResponseParams `json:"Response"`
}

func (r *DeleteDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFenceBindRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 围栏绑定的产品信息
	Items []*FenceBindProductItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type DeleteFenceBindRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 围栏绑定的产品信息
	Items []*FenceBindProductItem `json:"Items,omitnil,omitempty" name:"Items"`
}

func (r *DeleteFenceBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFenceBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFenceBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFenceBindResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFenceBindResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFenceBindResponseParams `json:"Response"`
}

func (r *DeleteFenceBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFenceBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaFrequencyRequestParams struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitnil,omitempty" name:"FreqId"`
}

type DeleteLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitnil,omitempty" name:"FreqId"`
}

func (r *DeleteLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaFrequencyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoRaFrequencyResponseParams `json:"Response"`
}

func (r *DeleteLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaGatewayRequestParams struct {
	// LoRa 网关 Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DeleteLoRaGatewayRequest struct {
	*tchttp.BaseRequest
	
	// LoRa 网关 Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

func (r *DeleteLoRaGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoRaGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoRaGatewayResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoRaGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoRaGatewayResponseParams `json:"Response"`
}

func (r *DeleteLoRaGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoRaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionFenceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`
}

type DeletePositionFenceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`
}

func (r *DeletePositionFenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionFenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FenceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePositionFenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionFenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePositionFenceResponse struct {
	*tchttp.BaseResponse
	Response *DeletePositionFenceResponseParams `json:"Response"`
}

func (r *DeletePositionFenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionFenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionSpaceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DeletePositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DeletePositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePositionSpaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DeletePositionSpaceResponseParams `json:"Response"`
}

func (r *DeletePositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStudioProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type DeleteStudioProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

func (r *DeleteStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStudioProductResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStudioProductResponseParams `json:"Response"`
}

func (r *DeleteStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DeleteTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DeleteTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicPolicyResponseParams `json:"Response"`
}

func (r *DeleteTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRuleRequestParams struct {
	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type DeleteTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *DeleteTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicRuleResponseParams `json:"Response"`
}

func (r *DeleteTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivateDeviceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeActivateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeActivateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActivateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivateDeviceResponseParams struct {
	// 设备激活详情信息
	Data *ActivateDeviceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActivateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActivateDeviceResponseParams `json:"Response"`
}

func (r *DescribeActivateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivateLicenseServiceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 激活码类型
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`
}

type DescribeActivateLicenseServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 激活码类型
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`
}

func (r *DescribeActivateLicenseServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivateLicenseServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LicenseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActivateLicenseServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActivateLicenseServiceResponseParams struct {
	// 增值服务激活码信息
	Data []*LicenseServiceNumInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeActivateLicenseServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActivateLicenseServiceResponseParams `json:"Response"`
}

func (r *DescribeActivateLicenseServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActivateLicenseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchProductionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 量产ID
	BatchProductionId *string `json:"BatchProductionId,omitnil,omitempty" name:"BatchProductionId"`
}

type DescribeBatchProductionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 量产ID
	BatchProductionId *string `json:"BatchProductionId,omitnil,omitempty" name:"BatchProductionId"`
}

func (r *DescribeBatchProductionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchProductionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "BatchProductionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchProductionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchProductionResponseParams struct {
	// 量产数量。
	BatchCnt *int64 `json:"BatchCnt,omitnil,omitempty" name:"BatchCnt"`

	// 烧录方式。
	BurnMethod *int64 `json:"BurnMethod,omitnil,omitempty" name:"BurnMethod"`

	// 创建时间。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 下载URL。
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 生成方式。
	GenerationMethod *int64 `json:"GenerationMethod,omitnil,omitempty" name:"GenerationMethod"`

	// 上传URL。
	UploadUrl *string `json:"UploadUrl,omitnil,omitempty" name:"UploadUrl"`

	// 成功数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 量产最后失败原因
	LastFailedReason *string `json:"LastFailedReason,omitnil,omitempty" name:"LastFailedReason"`

	// 量产状态  0：任务创建，未量产；1：处理中；2：量产结束上传结果中；3：任务完成
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchProductionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchProductionResponseParams `json:"Response"`
}

func (r *DescribeBatchProductionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchProductionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindedProductsRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否跨账号绑定产品
	ProductSource *int64 `json:"ProductSource,omitnil,omitempty" name:"ProductSource"`
}

type DescribeBindedProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否跨账号绑定产品
	ProductSource *int64 `json:"ProductSource,omitnil,omitempty" name:"ProductSource"`
}

func (r *DescribeBindedProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindedProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindedProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindedProductsResponseParams struct {
	// 当前分页的子产品数组
	Products []*BindProductInfo `json:"Products,omitnil,omitempty" name:"Products"`

	// 绑定的子产品总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBindedProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindedProductsResponseParams `json:"Response"`
}

func (r *DescribeBindedProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindedProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceCallbackRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type DescribeCloudStorageAIServiceCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

func (r *DescribeCloudStorageAIServiceCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageAIServiceCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceCallbackResponseParams struct {
	// 推送类型。http：HTTP 回调
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// HTTP 回调 URL
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// HTTP 回调鉴权 Token
	CallbackToken *string `json:"CallbackToken,omitnil,omitempty" name:"CallbackToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageAIServiceCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageAIServiceCallbackResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageAIServiceCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

type DescribeCloudStorageAIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

func (r *DescribeCloudStorageAIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageAIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceResponseParams struct {
	// 云存 AI 套餐类型。可能取值：
	// 
	// - `1`：全时套餐
	// - `2`：事件套餐
	// - `3`：低功耗套餐
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 云存 AI 套餐生效状态。可能取值：
	// 
	// - `0`：未开通或已过期
	// - `1`：生效中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 云存 AI 套餐过期时间 UNIX 时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 视频分析启用状态
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageAIServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageAIServiceResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageAIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceTaskRequestParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 下载 URL 的过期时间。
	// 
	// 若传入该参数，则响应中将包含所有文件的下载 URL
	FileURLExpireTime *int64 `json:"FileURLExpireTime,omitnil,omitempty" name:"FileURLExpireTime"`
}

type DescribeCloudStorageAIServiceTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 下载 URL 的过期时间。
	// 
	// 若传入该参数，则响应中将包含所有文件的下载 URL
	FileURLExpireTime *int64 `json:"FileURLExpireTime,omitnil,omitempty" name:"FileURLExpireTime"`
}

func (r *DescribeCloudStorageAIServiceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "FileURLExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageAIServiceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceTaskResponseParams struct {
	// 任务信息
	TaskInfo *CloudStorageAIServiceTask `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageAIServiceTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageAIServiceTaskResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageAIServiceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceTasksRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 任务状态。可选值：
	// - （不传）：查询全部状态的任务
	// - `1`：失败
	// - `2`：成功但结果为空
	// - `3`：成功且结果非空
	// - `4`：执行中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 设备名称列表。
	// 
	// 当需要同时查询多台设备的任务列表时传入，优先级高于参数 `DeviceName`
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`

	// 查询任务时间范围的起始时间（秒级 UNIX 时间戳）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务时间范围的结束时间（秒级 UNIX 时间戳）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 下载 URL 的过期时间。
	// 
	// 若传入该参数，则响应中将包含所有文件的下载 URL
	FileURLExpireTime *int64 `json:"FileURLExpireTime,omitnil,omitempty" name:"FileURLExpireTime"`
}

type DescribeCloudStorageAIServiceTasksRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 任务状态。可选值：
	// - （不传）：查询全部状态的任务
	// - `1`：失败
	// - `2`：成功但结果为空
	// - `3`：成功且结果非空
	// - `4`：执行中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 设备名称列表。
	// 
	// 当需要同时查询多台设备的任务列表时传入，优先级高于参数 `DeviceName`
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`

	// 查询任务时间范围的起始时间（秒级 UNIX 时间戳）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询任务时间范围的结束时间（秒级 UNIX 时间戳）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 下载 URL 的过期时间。
	// 
	// 若传入该参数，则响应中将包含所有文件的下载 URL
	FileURLExpireTime *int64 `json:"FileURLExpireTime,omitnil,omitempty" name:"FileURLExpireTime"`
}

func (r *DescribeCloudStorageAIServiceTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ServiceType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "UserId")
	delete(f, "ChannelId")
	delete(f, "DeviceNames")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FileURLExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageAIServiceTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageAIServiceTasksResponseParams struct {
	// 任务列表
	Tasks []*CloudStorageAIServiceTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务数量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageAIServiceTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageAIServiceTasksResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageAIServiceTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageAIServiceTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageDateRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeCloudStorageDateRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeCloudStorageDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageDateResponseParams struct {
	// 云存日期数组，["2021-01-05","2021-01-06"]
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageDateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageDateResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageEventsRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 查询数据项目的最大数量, 默认为10。假设传Size=10，返回的实际事件数量为N，则 5 <= N <= 10。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeCloudStorageEventsRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 查询数据项目的最大数量, 默认为10。假设传Size=10，返回的实际事件数量为N，则 5 <= N <= 10。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeCloudStorageEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Context")
	delete(f, "Size")
	delete(f, "EventId")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageEventsResponseParams struct {
	// 云存事件列表
	Events []*CloudStorageEvent `json:"Events,omitnil,omitempty" name:"Events"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 拉取结果是否已经结束
	Listover *bool `json:"Listover,omitnil,omitempty" name:"Listover"`

	// 内部结果数量，并不等同于事件总数。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 视频播放URL
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageEventsResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageEventsWithAITasksRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 事件关联的视频 AI 分析服务类型（支持多选）。可选值：
	// 
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceTypes []*string `json:"ServiceTypes,omitnil,omitempty" name:"ServiceTypes"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 查询数据项目的最大数量, 默认为10。假设传Size=10，返回的实际事件数量为N，则 5 <= N <= 10。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeCloudStorageEventsWithAITasksRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 事件关联的视频 AI 分析服务类型（支持多选）。可选值：
	// 
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceTypes []*string `json:"ServiceTypes,omitnil,omitempty" name:"ServiceTypes"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 查询数据项目的最大数量, 默认为10。假设传Size=10，返回的实际事件数量为N，则 5 <= N <= 10。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeCloudStorageEventsWithAITasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageEventsWithAITasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ServiceTypes")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Context")
	delete(f, "Size")
	delete(f, "EventId")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageEventsWithAITasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageEventsWithAITasksResponseParams struct {
	// 云存事件列表
	Events []*CloudStorageEventWithAITasks `json:"Events,omitnil,omitempty" name:"Events"`

	// 请求上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 拉取结果是否已经结束
	Listover *bool `json:"Listover,omitnil,omitempty" name:"Listover"`

	// 内部结果数量，并不等同于事件总数。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 视频播放URL
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageEventsWithAITasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageEventsWithAITasksResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageEventsWithAITasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageEventsWithAITasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageMultiThumbnailRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 多个缩略图文件名根据 | 分割
	MultiThumbnail *string `json:"MultiThumbnail,omitnil,omitempty" name:"MultiThumbnail"`
}

type DescribeCloudStorageMultiThumbnailRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 多个缩略图文件名根据 | 分割
	MultiThumbnail *string `json:"MultiThumbnail,omitnil,omitempty" name:"MultiThumbnail"`
}

func (r *DescribeCloudStorageMultiThumbnailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageMultiThumbnailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "MultiThumbnail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageMultiThumbnailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageMultiThumbnailResponseParams struct {
	// 缩略图访问地址
	ThumbnailURLInfoList []*ThumbnailURLInfoList `json:"ThumbnailURLInfoList,omitnil,omitempty" name:"ThumbnailURLInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageMultiThumbnailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageMultiThumbnailResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageMultiThumbnailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageMultiThumbnailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageOrderRequestParams struct {
	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

type DescribeCloudStorageOrderRequest struct {
	*tchttp.BaseRequest
	
	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`
}

func (r *DescribeCloudStorageOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageOrderResponseParams struct {
	// 云存套餐开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 云存套餐过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 套餐id
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 套餐状态
	// 0：等待生效
	// 1: 已过期
	// 2:生效
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 通道id
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 订单金额，单位为分
	Price *uint64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 支付金额，单位为分
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageOrderResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageOrderResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStoragePackageConsumeDetailsRequestParams struct {
	// 开始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeCloudStoragePackageConsumeDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeCloudStoragePackageConsumeDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStoragePackageConsumeDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStoragePackageConsumeDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStoragePackageConsumeDetailsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStoragePackageConsumeDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStoragePackageConsumeDetailsResponseParams `json:"Response"`
}

func (r *DescribeCloudStoragePackageConsumeDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStoragePackageConsumeDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStoragePackageConsumeStatsRequestParams struct {
	// 开始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期，开始与结束日期间隔不可超过一年
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeCloudStoragePackageConsumeStatsRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期，开始与结束日期间隔不可超过一年
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeCloudStoragePackageConsumeStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStoragePackageConsumeStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStoragePackageConsumeStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStoragePackageConsumeStatsResponseParams struct {
	// 统计列表详情
	Stats []*PackageConsumeStat `json:"Stats,omitnil,omitempty" name:"Stats"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStoragePackageConsumeStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStoragePackageConsumeStatsResponseParams `json:"Response"`
}

func (r *DescribeCloudStoragePackageConsumeStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStoragePackageConsumeStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID 非NVR设备不填 NVR设备必填 默认为无	
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID 非NVR设备不填 NVR设备必填 默认为无	
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageResponseParams struct {
	// 云存开启状态，1为开启，0为未开启或已过期
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 云存类型，1为全时云存，2为事件云存
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 云存套餐过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 云存回看时长
	ShiftDuration *uint64 `json:"ShiftDuration,omitnil,omitempty" name:"ShiftDuration"`

	// 云存用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageStreamDataRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 图片流事件开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type DescribeCloudStorageStreamDataRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 图片流事件开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

func (r *DescribeCloudStorageStreamDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageStreamDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "StartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageStreamDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageStreamDataResponseParams struct {
	// 图片流视频地址
	VideoStream *string `json:"VideoStream,omitnil,omitempty" name:"VideoStream"`

	// 图片流音频地址
	AudioStream *string `json:"AudioStream,omitnil,omitempty" name:"AudioStream"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageStreamDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageStreamDataResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageStreamDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageStreamDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageThumbnailListRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 缩略图文件名列表
	ThumbnailList []*string `json:"ThumbnailList,omitnil,omitempty" name:"ThumbnailList"`
}

type DescribeCloudStorageThumbnailListRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 缩略图文件名列表
	ThumbnailList []*string `json:"ThumbnailList,omitnil,omitempty" name:"ThumbnailList"`
}

func (r *DescribeCloudStorageThumbnailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageThumbnailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ThumbnailList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageThumbnailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageThumbnailListResponseParams struct {
	// 缩略图访问地址
	ThumbnailURLInfoList []*ThumbnailURLInfoList `json:"ThumbnailURLInfoList,omitnil,omitempty" name:"ThumbnailURLInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageThumbnailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageThumbnailListResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageThumbnailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageThumbnailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageThumbnailRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 缩略图文件名
	Thumbnail *string `json:"Thumbnail,omitnil,omitempty" name:"Thumbnail"`
}

type DescribeCloudStorageThumbnailRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 缩略图文件名
	Thumbnail *string `json:"Thumbnail,omitnil,omitempty" name:"Thumbnail"`
}

func (r *DescribeCloudStorageThumbnailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageThumbnailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Thumbnail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageThumbnailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageThumbnailResponseParams struct {
	// 缩略图访问地址
	ThumbnailURL *string `json:"ThumbnailURL,omitnil,omitempty" name:"ThumbnailURL"`

	// 缩略图访问地址的过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageThumbnailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageThumbnailResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageThumbnailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageThumbnailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageTimeRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存日期，例如"2020-01-05"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 开始时间，unix时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeCloudStorageTimeRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存日期，例如"2020-01-05"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 开始时间，unix时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，unix时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeCloudStorageTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Date")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageTimeResponseParams struct {
	// 接口返回数据
	Data *CloudStorageTimeData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageTimeResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageUsersRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCloudStorageUsersRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCloudStorageUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudStorageUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudStorageUsersResponseParams struct {
	// 用户总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 用户信息
	Users []*CloudStorageUserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudStorageUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudStorageUsersResponseParams `json:"Response"`
}

func (r *DescribeCloudStorageUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudStorageUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCsReportCountDataInfoRequestParams struct {
	// 产品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 统计开始时间戳
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 统计结束时间戳
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 设备通道
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeCsReportCountDataInfoRequest struct {
	*tchttp.BaseRequest
	
	// 产品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 统计开始时间戳
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 统计结束时间戳
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 设备通道
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeCsReportCountDataInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCsReportCountDataInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCsReportCountDataInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCsReportCountDataInfoResponseParams struct {
	// 云存上报统计信息
	Data *CountDataInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCsReportCountDataInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCsReportCountDataInfoResponseParams `json:"Response"`
}

func (r *DescribeCsReportCountDataInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCsReportCountDataInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceBindGatewayRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DescribeDeviceBindGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceBindGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceBindGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceBindGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceBindGatewayResponseParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 网关产品名称
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 设备对应产品所属的主账号名称
	GatewayProductOwnerName *string `json:"GatewayProductOwnerName,omitnil,omitempty" name:"GatewayProductOwnerName"`

	// 设备对应产品所属的主账号 UIN
	GatewayProductOwnerUin *string `json:"GatewayProductOwnerUin,omitnil,omitempty" name:"GatewayProductOwnerUin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceBindGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceBindGatewayResponseParams `json:"Response"`
}

func (r *DescribeDeviceBindGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceBindGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataHistoryRequestParams struct {
	// 区间开始时间（Unix 时间戳，毫秒级）
	MinTime *uint64 `json:"MinTime,omitnil,omitempty" name:"MinTime"`

	// 区间结束时间（Unix 时间戳，毫秒级）
	MaxTime *uint64 `json:"MaxTime,omitnil,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 属性字段名称，对应数据模板中功能属性的标识符
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 返回条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type DescribeDeviceDataHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 区间开始时间（Unix 时间戳，毫秒级）
	MinTime *uint64 `json:"MinTime,omitnil,omitempty" name:"MinTime"`

	// 区间结束时间（Unix 时间戳，毫秒级）
	MaxTime *uint64 `json:"MaxTime,omitnil,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 属性字段名称，对应数据模板中功能属性的标识符
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 返回条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

func (r *DescribeDeviceDataHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "FieldName")
	delete(f, "Limit")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceDataHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataHistoryResponseParams struct {
	// 属性字段名称，对应数据模板中功能属性的标识符
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 数据是否已全部返回，true 表示数据全部返回，false 表示还有数据待返回，可将 Context 作为入参，继续查询返回结果。
	Listover *bool `json:"Listover,omitnil,omitempty" name:"Listover"`

	// 检索上下文，当 ListOver 为false时，可以用此上下文，继续读取后续数据
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 历史数据结果数组，返回对应时间点及取值。
	Results []*DeviceDataHistoryItem `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceDataHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceDataHistoryResponseParams `json:"Response"`
}

func (r *DescribeDeviceDataHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DescribeDeviceDataRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *DescribeDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataResponseParams struct {
	// 设备数据
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceDataResponseParams `json:"Response"`
}

func (r *DescribeDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceFirmWareRequestParams struct {
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DescribeDeviceFirmWareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceFirmWareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceFirmWareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceFirmWareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceFirmWareResponseParams struct {
	// 固件信息
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceFirmWareResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceFirmWareResponseParams `json:"Response"`
}

func (r *DescribeDeviceFirmWareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceFirmWareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceFirmwaresRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DescribeDeviceFirmwaresRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceFirmwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceFirmwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceFirmwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceFirmwaresResponseParams struct {
	// 固件信息列表
	Firmwares []*DeviceFirmwareInfo `json:"Firmwares,omitnil,omitempty" name:"Firmwares"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceFirmwaresResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceFirmwaresResponseParams `json:"Response"`
}

func (r *DescribeDeviceFirmwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceFirmwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceLocationSolveRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 定位解析类型，wifi或GNSSNavigation
	LocationType *string `json:"LocationType,omitnil,omitempty" name:"LocationType"`

	// LoRaEdge卫星导航电文
	GNSSNavigation *string `json:"GNSSNavigation,omitnil,omitempty" name:"GNSSNavigation"`

	// wifi信息
	WiFiInfo []*WifiInfo `json:"WiFiInfo,omitnil,omitempty" name:"WiFiInfo"`
}

type DescribeDeviceLocationSolveRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 定位解析类型，wifi或GNSSNavigation
	LocationType *string `json:"LocationType,omitnil,omitempty" name:"LocationType"`

	// LoRaEdge卫星导航电文
	GNSSNavigation *string `json:"GNSSNavigation,omitnil,omitempty" name:"GNSSNavigation"`

	// wifi信息
	WiFiInfo []*WifiInfo `json:"WiFiInfo,omitnil,omitempty" name:"WiFiInfo"`
}

func (r *DescribeDeviceLocationSolveRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceLocationSolveRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "LocationType")
	delete(f, "GNSSNavigation")
	delete(f, "WiFiInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceLocationSolveRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceLocationSolveResponseParams struct {
	// 经度
	Longitude *float64 `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 纬度
	Latitude *float64 `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// 类型
	LocationType *string `json:"LocationType,omitnil,omitempty" name:"LocationType"`

	// 误差精度预估，单位为米
	Accuracy *float64 `json:"Accuracy,omitnil,omitempty" name:"Accuracy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceLocationSolveResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceLocationSolveResponseParams `json:"Response"`
}

func (r *DescribeDeviceLocationSolveResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceLocationSolveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePackagesRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用户id
	CSUserId *string `json:"CSUserId,omitnil,omitempty" name:"CSUserId"`

	// 通道id
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeDevicePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 分页拉取数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页拉取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用户id
	CSUserId *string `json:"CSUserId,omitnil,omitempty" name:"CSUserId"`

	// 通道id
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeDevicePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CSUserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePackagesResponseParams struct {
	// 有效云存套餐数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 有效云存套餐列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Packages []*PackageInfo `json:"Packages,omitnil,omitempty" name:"Packages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDevicePackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePackagesResponseParams `json:"Response"`
}

func (r *DescribeDevicePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePositionListRequestParams struct {
	// 产品标识列表
	ProductIdList []*string `json:"ProductIdList,omitnil,omitempty" name:"ProductIdList"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitnil,omitempty" name:"CoordinateType"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDevicePositionListRequest struct {
	*tchttp.BaseRequest
	
	// 产品标识列表
	ProductIdList []*string `json:"ProductIdList,omitnil,omitempty" name:"ProductIdList"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitnil,omitempty" name:"CoordinateType"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDevicePositionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePositionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductIdList")
	delete(f, "CoordinateType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePositionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePositionListResponseParams struct {
	// 产品设备位置信息列表
	Positions []*ProductDevicesPositionItem `json:"Positions,omitnil,omitempty" name:"Positions"`

	// 产品设备位置信息的数目
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDevicePositionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePositionListResponseParams `json:"Response"`
}

func (r *DescribeDevicePositionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePositionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备ID，该字段有值将代替 ProductId/DeviceName
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *DescribeDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceResponseParams struct {
	// 设备信息
	Device *DeviceInfo `json:"Device,omitnil,omitempty" name:"Device"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceResponseParams `json:"Response"`
}

func (r *DescribeDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceBindListRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeFenceBindListRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeFenceBindListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceBindListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFenceBindListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceBindListResponseParams struct {
	// 围栏绑定的产品设备列表
	List []*FenceBindProductItem `json:"List,omitnil,omitempty" name:"List"`

	// 围栏绑定的设备总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFenceBindListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFenceBindListResponseParams `json:"Response"`
}

func (r *DescribeFenceBindListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceBindListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceEventListRequestParams struct {
	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 告警对应的产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 告警对应的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DescribeFenceEventListRequest struct {
	*tchttp.BaseRequest
	
	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 告警对应的产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 告警对应的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *DescribeFenceEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FenceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFenceEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFenceEventListResponseParams struct {
	// 围栏告警事件列表
	List []*FenceEventItem `json:"List,omitnil,omitempty" name:"List"`

	// 围栏告警事件总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFenceEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFenceEventListResponseParams `json:"Response"`
}

func (r *DescribeFenceEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFenceEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`
}

type DescribeFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`
}

func (r *DescribeFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareResponseParams struct {
	// 固件版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 固件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 固件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 固件Md5值
	Md5sum *string `json:"Md5sum,omitnil,omitempty" name:"Md5sum"`

	// 固件上传的秒级时间戳
	Createtime *uint64 `json:"Createtime,omitnil,omitempty" name:"Createtime"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 固件升级模块
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 固件用户自定义配置信息
	UserDefined *string `json:"UserDefined,omitnil,omitempty" name:"UserDefined"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareResponseParams `json:"Response"`
}

func (r *DescribeFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 固件任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeFirmwareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 固件任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeFirmwareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareTaskResponseParams struct {
	// 固件任务ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 固件任务状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 固件任务创建时间，单位：秒
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 固件任务升级类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 固件任务升级模式。originalVersion（按版本号升级）、filename（提交文件升级）、devicenames（按设备名称升级）
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 原始固件版本号，在UpgradeMode是originalVersion升级模式下会返回
	OriginalVersion *string `json:"OriginalVersion,omitnil,omitempty" name:"OriginalVersion"`

	// 创建账号ID
	CreateUserId *uint64 `json:"CreateUserId,omitnil,omitempty" name:"CreateUserId"`

	// 创建账号ID昵称
	CreatorNickName *string `json:"CreatorNickName,omitnil,omitempty" name:"CreatorNickName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirmwareTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareTaskResponseParams `json:"Response"`
}

func (r *DescribeFirmwareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareUpdateStatusRequestParams struct {
	// 产品 ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DescribeFirmwareUpdateStatusRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID。
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *DescribeFirmwareUpdateStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareUpdateStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFirmwareUpdateStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFirmwareUpdateStatusResponseParams struct {
	// 升级任务源版本。
	OriVersion *string `json:"OriVersion,omitnil,omitempty" name:"OriVersion"`

	// 升级任务目标版本。
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// 升级状态：- 0：设备离线。- 1：待处理。- 2：消息下发成功。- 3：下载中。- 4：烧录中。- 5：失败。- 6：升级完成。- 7：正在处理中。- 8：等待用户确认。- 10：升级超时。- 20：下载完成。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 进度
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFirmwareUpdateStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFirmwareUpdateStatusResponseParams `json:"Response"`
}

func (r *DescribeFirmwareUpdateStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFirmwareUpdateStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFreeCloudStorageNumRequestParams struct {

}

type DescribeFreeCloudStorageNumRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFreeCloudStorageNumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFreeCloudStorageNumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFreeCloudStorageNumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFreeCloudStorageNumResponseParams struct {
	// 套餐包信息
	PackageInfos []*CloudStoragePackageInfo `json:"PackageInfos,omitnil,omitempty" name:"PackageInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFreeCloudStorageNumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFreeCloudStorageNumResponseParams `json:"Response"`
}

func (r *DescribeFreeCloudStorageNumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFreeCloudStorageNumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayBindDevicesRequestParams struct {
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 子产品的ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 分页的偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGatewayBindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 子产品的ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 分页的偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGatewayBindDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayBindDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayBindDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayBindDevicesResponseParams struct {
	// 子设备信息。
	Devices []*BindDeviceInfo `json:"Devices,omitnil,omitempty" name:"Devices"`

	// 子设备总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 子设备所属的产品名。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewayBindDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayBindDevicesResponseParams `json:"Response"`
}

func (r *DescribeGatewayBindDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayBindDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubDeviceListRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGatewaySubDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGatewaySubDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewaySubDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubDeviceListResponseParams struct {
	// 设备的总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 设备列表
	DeviceList []*FamilySubDevice `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewaySubDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewaySubDeviceListResponseParams `json:"Response"`
}

func (r *DescribeGatewaySubDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubProductsRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 分页的偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否跨账号产品
	ProductSource *int64 `json:"ProductSource,omitnil,omitempty" name:"ProductSource"`
}

type DescribeGatewaySubProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 分页的偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否跨账号产品
	ProductSource *int64 `json:"ProductSource,omitnil,omitempty" name:"ProductSource"`
}

func (r *DescribeGatewaySubProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	delete(f, "ProductSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewaySubProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewaySubProductsResponseParams struct {
	// 当前分页的可绑定或解绑的产品信息。
	Products []*BindProductInfo `json:"Products,omitnil,omitempty" name:"Products"`

	// 可绑定或解绑的产品总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewaySubProductsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewaySubProductsResponseParams `json:"Response"`
}

func (r *DescribeGatewaySubProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewaySubProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 附加查询返回包含字段值，不传返回0，有效值 ProductNum、ProjectNum、UsedDeviceNum、TotalDevice、ActivateDevice
	Include []*string `json:"Include,omitnil,omitempty" name:"Include"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品ID，-1 代表全部产品
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 附加查询返回包含字段值，不传返回0，有效值 ProductNum、ProjectNum、UsedDeviceNum、TotalDevice、ActivateDevice
	Include []*string `json:"Include,omitnil,omitempty" name:"Include"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品ID，-1 代表全部产品
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Include")
	delete(f, "ProjectId")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// 实例信息
	Data *InstanceDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoRaFrequencyRequestParams struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitnil,omitempty" name:"FreqId"`
}

type DescribeLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitnil,omitempty" name:"FreqId"`
}

func (r *DescribeLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoRaFrequencyResponseParams struct {
	// 返回详情项
	Data *LoRaFrequencyEntry `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoRaFrequencyResponseParams `json:"Response"`
}

func (r *DescribeLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelDefinitionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type DescribeModelDefinitionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

func (r *DescribeModelDefinitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelDefinitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelDefinitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelDefinitionResponseParams struct {
	// 产品数据模板
	Model *ProductModelDefinition `json:"Model,omitnil,omitempty" name:"Model"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelDefinitionResponseParams `json:"Response"`
}

func (r *DescribeModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeP2PRouteRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type DescribeP2PRouteRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *DescribeP2PRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeP2PRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeP2PRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeP2PRouteResponseParams struct {
	// 当前p2p线路
	RouteId *uint64 `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeP2PRouteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeP2PRouteResponseParams `json:"Response"`
}

func (r *DescribeP2PRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeP2PRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackageConsumeTaskRequestParams struct {
	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribePackageConsumeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribePackageConsumeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageConsumeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackageConsumeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackageConsumeTaskResponseParams struct {
	// 文件下载的url，文件详情是套餐包消耗详情
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePackageConsumeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackageConsumeTaskResponseParams `json:"Response"`
}

func (r *DescribePackageConsumeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageConsumeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackageConsumeTasksRequestParams struct {
	// 分页单页量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，第一页为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribePackageConsumeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 分页单页量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的偏移量，第一页为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribePackageConsumeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageConsumeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackageConsumeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackageConsumeTasksResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表
	List []*PackageConsumeTask `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePackageConsumeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackageConsumeTasksResponseParams `json:"Response"`
}

func (r *DescribePackageConsumeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackageConsumeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionFenceListRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePositionFenceListRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePositionFenceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionFenceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePositionFenceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePositionFenceListResponseParams struct {
	// 围栏列表
	List []*PositionFenceInfo `json:"List,omitnil,omitempty" name:"List"`

	// 围栏数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePositionFenceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePositionFenceListResponseParams `json:"Response"`
}

func (r *DescribePositionFenceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePositionFenceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductCloudStorageAIServiceRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type DescribeProductCloudStorageAIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

func (r *DescribeProductCloudStorageAIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductCloudStorageAIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductCloudStorageAIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductCloudStorageAIServiceResponseParams struct {
	// 开通状态
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 当前账号是否可开通
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductCloudStorageAIServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductCloudStorageAIServiceResponseParams `json:"Response"`
}

func (r *DescribeProductCloudStorageAIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductCloudStorageAIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectResponseParams struct {
	// 返回信息
	Project *ProjectEntryEx `json:"Project,omitnil,omitempty" name:"Project"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectResponseParams `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceFenceEventListRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSpaceFenceEventListRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 围栏告警信息的查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 围栏告警信息的查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSpaceFenceEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceFenceEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceFenceEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceFenceEventListResponseParams struct {
	// 围栏告警事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*FenceEventItem `json:"List,omitnil,omitempty" name:"List"`

	// 围栏告警事件总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceFenceEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceFenceEventListResponseParams `json:"Response"`
}

func (r *DescribeSpaceFenceEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceFenceEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStudioProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type DescribeStudioProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

func (r *DescribeStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStudioProductResponseParams struct {
	// 产品详情
	Product *ProductEntry `json:"Product,omitnil,omitempty" name:"Product"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStudioProductResponseParams `json:"Response"`
}

func (r *DescribeStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTWeSeeConfigRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeTWeSeeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeTWeSeeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTWeSeeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTWeSeeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTWeSeeConfigResponseParams struct {
	// 是否开启视频摘要
	EnableSummary *bool `json:"EnableSummary,omitnil,omitempty" name:"EnableSummary"`

	// 是否开启视频搜索
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTWeSeeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTWeSeeConfigResponseParams `json:"Response"`
}

func (r *DescribeTWeSeeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTWeSeeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Topic名字
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Topic名字
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DescribeTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicPolicyResponseParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic权限
	Privilege *uint64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicPolicyResponseParams `json:"Response"`
}

func (r *DescribeTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRuleRequestParams struct {
	// 规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type DescribeTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *DescribeTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRuleResponseParams struct {
	// 规则描述。
	Rule *TopicRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 规则绑定的标签
	CamTag []*CamTag `json:"CamTag,omitnil,omitempty" name:"CamTag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicRuleResponseParams `json:"Response"`
}

func (r *DescribeTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnbindedDevicesRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUnbindedDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUnbindedDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnbindedDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnbindedDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnbindedDevicesResponseParams struct {
	// 未绑定的设备列表
	UnbindedDevices []*BindDeviceInfo `json:"UnbindedDevices,omitnil,omitempty" name:"UnbindedDevices"`

	// 设备的总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUnbindedDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnbindedDevicesResponseParams `json:"Response"`
}

func (r *DescribeUnbindedDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnbindedDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoLicenseRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeVideoLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeVideoLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoLicenseResponseParams struct {
	// 视频激活码分类概览
	License []*VideoLicenseEntity `json:"License,omitnil,omitempty" name:"License"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoLicenseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoLicenseResponseParams `json:"Response"`
}

func (r *DescribeVideoLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceActivationDetail struct {
	// 可注册设备数
	TotalDeviceNum *int64 `json:"TotalDeviceNum,omitnil,omitempty" name:"TotalDeviceNum"`

	// 已注册设备数
	UsedDeviceNum *int64 `json:"UsedDeviceNum,omitnil,omitempty" name:"UsedDeviceNum"`

	// 设备授权数
	TotalNormalLicense *int64 `json:"TotalNormalLicense,omitnil,omitempty" name:"TotalNormalLicense"`

	// 已使用设备授权数
	UsedNormalLicense *int64 `json:"UsedNormalLicense,omitnil,omitempty" name:"UsedNormalLicense"`

	// 蓝牙授权数
	TotalBluetoothLicense *int64 `json:"TotalBluetoothLicense,omitnil,omitempty" name:"TotalBluetoothLicense"`

	// 已使用蓝牙授权数
	UsedBluetoothLicense *int64 `json:"UsedBluetoothLicense,omitnil,omitempty" name:"UsedBluetoothLicense"`

	// 可免费注册设备数
	TotalFreeLicense *int64 `json:"TotalFreeLicense,omitnil,omitempty" name:"TotalFreeLicense"`

	// 已使用注册设备数
	UsedFreeLicense *int64 `json:"UsedFreeLicense,omitnil,omitempty" name:"UsedFreeLicense"`
}

type DeviceActiveResult struct {
	// 模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModelId is deprecated.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// SN信息
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 设备激活状态，0：激活成功；50011：系统错误；50012：产品不存在；50013：设备不存在；50014：产品无权限；50015：不是音视频产品；50016：SN格式错误；50017：激活码类型错误；50018：激活次数限频；50019：激活码不足；50020：SN已暂停；
	ErrCode *uint64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type DeviceData struct {
	// 设备证书，用于 TLS 建立链接时校验客户端身份。采用非对称加密时返回该参数。
	DeviceCert *string `json:"DeviceCert,omitnil,omitempty" name:"DeviceCert"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备私钥，用于 TLS 建立链接时校验客户端身份，腾讯云后台不保存，请妥善保管。采用非对称加密时返回该参数。
	DevicePrivateKey *string `json:"DevicePrivateKey,omitnil,omitempty" name:"DevicePrivateKey"`

	// 对称加密密钥，base64编码。采用对称加密时返回该参数。
	DevicePsk *string `json:"DevicePsk,omitnil,omitempty" name:"DevicePsk"`
}

type DeviceDataHistoryItem struct {
	// 时间点，毫秒时间戳
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 字段取值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DeviceFirmwareInfo struct {
	// 固件类型
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 固件版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 最后更新时间
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DeviceInfo struct {
	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 0: 离线, 1: 在线, 2: 获取失败, 3 未激活
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 设备密钥，密钥加密的设备返回
	DevicePsk *string `json:"DevicePsk,omitnil,omitempty" name:"DevicePsk"`

	// 首次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOnlineTime *int64 `json:"FirstOnlineTime,omitnil,omitempty" name:"FirstOnlineTime"`

	// 最后一次上线时间
	LoginTime *int64 `json:"LoginTime,omitnil,omitempty" name:"LoginTime"`

	// 设备创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 设备固件版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 设备证书
	DeviceCert *string `json:"DeviceCert,omitnil,omitempty" name:"DeviceCert"`

	// 日志级别
	LogLevel *int64 `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// LoRaWAN 设备地址
	DevAddr *string `json:"DevAddr,omitnil,omitempty" name:"DevAddr"`

	// LoRaWAN 应用密钥
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// LoRaWAN 设备唯一标识
	DevEUI *string `json:"DevEUI,omitnil,omitempty" name:"DevEUI"`

	// LoRaWAN 应用会话密钥
	AppSKey *string `json:"AppSKey,omitnil,omitempty" name:"AppSKey"`

	// LoRaWAN 网络会话密钥
	NwkSKey *string `json:"NwkSKey,omitnil,omitempty" name:"NwkSKey"`

	// 创建人Id
	CreateUserId *int64 `json:"CreateUserId,omitnil,omitempty" name:"CreateUserId"`

	// 创建人昵称
	CreatorNickName *string `json:"CreatorNickName,omitnil,omitempty" name:"CreatorNickName"`

	// 启用/禁用状态
	EnableState *int64 `json:"EnableState,omitnil,omitempty" name:"EnableState"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 设备类型（设备、子设备、网关）
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 是否是 lora 设备
	IsLora *bool `json:"IsLora,omitnil,omitempty" name:"IsLora"`
}

type DevicePositionItem struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 位置信息时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 设备经度信息
	Longitude *float64 `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 设备纬度信息
	Latitude *float64 `json:"Latitude,omitnil,omitempty" name:"Latitude"`
}

type DeviceSignatureInfo struct {
	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备签名
	DeviceSignature *string `json:"DeviceSignature,omitnil,omitempty" name:"DeviceSignature"`
}

type DeviceUser struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户角色 1所有者，0：其他分享者
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 家庭ID，所有者带该参数
	FamilyId *string `json:"FamilyId,omitnil,omitempty" name:"FamilyId"`

	// 家庭名称，所有者带该参数
	FamilyName *string `json:"FamilyName,omitnil,omitempty" name:"FamilyName"`
}

type DevicesItem struct {
	// 产品id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

// Predefined struct for user
type DirectBindDeviceInFamilyRequestParams struct {
	// 小程序appid
	IotAppID *string `json:"IotAppID,omitnil,omitempty" name:"IotAppID"`

	// 用户ID
	UserID *string `json:"UserID,omitnil,omitempty" name:"UserID"`

	// 家庭ID
	FamilyId *string `json:"FamilyId,omitnil,omitempty" name:"FamilyId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DirectBindDeviceInFamilyRequest struct {
	*tchttp.BaseRequest
	
	// 小程序appid
	IotAppID *string `json:"IotAppID,omitnil,omitempty" name:"IotAppID"`

	// 用户ID
	UserID *string `json:"UserID,omitnil,omitempty" name:"UserID"`

	// 家庭ID
	FamilyId *string `json:"FamilyId,omitnil,omitempty" name:"FamilyId"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 房间ID
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DirectBindDeviceInFamilyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DirectBindDeviceInFamilyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IotAppID")
	delete(f, "UserID")
	delete(f, "FamilyId")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DirectBindDeviceInFamilyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DirectBindDeviceInFamilyResponseParams struct {
	// 返回设备信息
	AppDeviceInfo *AppDeviceInfo `json:"AppDeviceInfo,omitnil,omitempty" name:"AppDeviceInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DirectBindDeviceInFamilyResponse struct {
	*tchttp.BaseResponse
	Response *DirectBindDeviceInFamilyResponseParams `json:"Response"`
}

func (r *DirectBindDeviceInFamilyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DirectBindDeviceInFamilyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type DisableTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *DisableTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableTopicRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *DisableTopicRuleResponseParams `json:"Response"`
}

func (r *DisableTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomByStrRoomIdFromTRTCRequestParams struct {
	// 房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

type DismissRoomByStrRoomIdFromTRTCRequest struct {
	*tchttp.BaseRequest
	
	// 房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`
}

func (r *DismissRoomByStrRoomIdFromTRTCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdFromTRTCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomByStrRoomIdFromTRTCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomByStrRoomIdFromTRTCResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DismissRoomByStrRoomIdFromTRTCResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomByStrRoomIdFromTRTCResponseParams `json:"Response"`
}

func (r *DismissRoomByStrRoomIdFromTRTCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdFromTRTCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type EnableTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *EnableTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTopicRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *EnableTopicRuleResponseParams `json:"Response"`
}

func (r *EnableTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventHistoryItem struct {
	// 事件的时间戳
	TimeStamp *int64 `json:"TimeStamp,omitnil,omitempty" name:"TimeStamp"`

	// 事件的产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 事件的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 事件的标识符ID
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 事件的类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件的数据
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type FamilySubDevice struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备别名
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 设备绑定的家庭ID
	FamilyId *string `json:"FamilyId,omitnil,omitempty" name:"FamilyId"`

	// 设备所在的房间ID，默认"0"
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 图标
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// grid图标
	IconUrlGrid *string `json:"IconUrlGrid,omitnil,omitempty" name:"IconUrlGrid"`

	// 设备绑定时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 设备更新时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type FenceAlarmPoint struct {
	// 围栏告警时间
	AlarmTime *int64 `json:"AlarmTime,omitnil,omitempty" name:"AlarmTime"`

	// 围栏告警位置的经度
	Longitude *float64 `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 围栏告警位置的纬度
	Latitude *float64 `json:"Latitude,omitnil,omitempty" name:"Latitude"`
}

type FenceBindDeviceItem struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 告警条件(In，进围栏报警；Out，出围栏报警；InOrOut，进围栏或者出围栏均报警)
	AlertCondition *string `json:"AlertCondition,omitnil,omitempty" name:"AlertCondition"`

	// 是否使能围栏(true，使能；false，禁用)
	FenceEnable *bool `json:"FenceEnable,omitnil,omitempty" name:"FenceEnable"`

	// 告警处理方法
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`
}

type FenceBindProductItem struct {
	// 围栏绑定的设备信息
	Devices []*FenceBindDeviceItem `json:"Devices,omitnil,omitempty" name:"Devices"`

	// 围栏绑定的产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type FenceEventItem struct {
	// 围栏事件的产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 围栏事件的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 围栏事件的告警类型（In，进围栏报警；Out，出围栏报警；InOrOut，进围栏或者出围栏均报警）
	AlertType *string `json:"AlertType,omitnil,omitempty" name:"AlertType"`

	// 围栏事件的设备位置信息
	Data *FenceAlarmPoint `json:"Data,omitnil,omitempty" name:"Data"`
}

type Filter struct {
	// 需要过滤的字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤的一个或多个值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FirmwareInfo struct {
	// 固件版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 固件MD5值
	Md5sum *string `json:"Md5sum,omitnil,omitempty" name:"Md5sum"`

	// 固件创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 固件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 固件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 固件升级模块
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 创建者子 uin
	CreateUserId *int64 `json:"CreateUserId,omitnil,omitempty" name:"CreateUserId"`

	// 创建者昵称
	CreatorNickName *string `json:"CreatorNickName,omitnil,omitempty" name:"CreatorNickName"`

	// 固件用户自定义配置信息
	UserDefined *string `json:"UserDefined,omitnil,omitempty" name:"UserDefined"`
}

// Predefined struct for user
type GenSingleDeviceSignatureOfPublicRequestParams struct {
	// 设备所属的产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 需要绑定的设备
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备绑定签名的有效时间,以秒为单位。取值范围：0 < Expire <= 86400，Expire == -1（十年）
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`
}

type GenSingleDeviceSignatureOfPublicRequest struct {
	*tchttp.BaseRequest
	
	// 设备所属的产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 需要绑定的设备
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备绑定签名的有效时间,以秒为单位。取值范围：0 < Expire <= 86400，Expire == -1（十年）
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`
}

func (r *GenSingleDeviceSignatureOfPublicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenSingleDeviceSignatureOfPublicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Expire")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenSingleDeviceSignatureOfPublicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenSingleDeviceSignatureOfPublicResponseParams struct {
	// 设备签名
	DeviceSignature *DeviceSignatureInfo `json:"DeviceSignature,omitnil,omitempty" name:"DeviceSignature"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenSingleDeviceSignatureOfPublicResponse struct {
	*tchttp.BaseResponse
	Response *GenSingleDeviceSignatureOfPublicResponseParams `json:"Response"`
}

func (r *GenSingleDeviceSignatureOfPublicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenSingleDeviceSignatureOfPublicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateCloudStorageAIServiceTaskFileURLRequestParams struct {
	// 产品 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 过期时间 UNIX 时间戳（默认值为当前时间 1 小时后，最大不超过文件所属任务的过期时间）
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type GenerateCloudStorageAIServiceTaskFileURLRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 过期时间 UNIX 时间戳（默认值为当前时间 1 小时后，最大不超过文件所属任务的过期时间）
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

func (r *GenerateCloudStorageAIServiceTaskFileURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCloudStorageAIServiceTaskFileURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "FileName")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateCloudStorageAIServiceTaskFileURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateCloudStorageAIServiceTaskFileURLResponseParams struct {
	// 文件下载 URL
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// 过期时间 UNIX 时间戳（最大不超过文件所属任务的过期时间）
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateCloudStorageAIServiceTaskFileURLResponse struct {
	*tchttp.BaseResponse
	Response *GenerateCloudStorageAIServiceTaskFileURLResponseParams `json:"Response"`
}

func (r *GenerateCloudStorageAIServiceTaskFileURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCloudStorageAIServiceTaskFileURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateSignedVideoURLRequestParams struct {
	// 视频播放原始URL地址
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`

	// 播放链接过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 通道ID 非NVR设备不填 NVR设备必填 默认为无	
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type GenerateSignedVideoURLRequest struct {
	*tchttp.BaseRequest
	
	// 视频播放原始URL地址
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`

	// 播放链接过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 通道ID 非NVR设备不填 NVR设备必填 默认为无	
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *GenerateSignedVideoURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateSignedVideoURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoURL")
	delete(f, "ExpireTime")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateSignedVideoURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateSignedVideoURLResponseParams struct {
	// 视频防盗链播放URL
	SignedVideoURL *string `json:"SignedVideoURL,omitnil,omitempty" name:"SignedVideoURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateSignedVideoURLResponse struct {
	*tchttp.BaseResponse
	Response *GenerateSignedVideoURLResponseParams `json:"Response"`
}

func (r *GenerateSignedVideoURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateSignedVideoURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthMiniProgramAppListRequestParams struct {
	// appId
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetAuthMiniProgramAppListRequest struct {
	*tchttp.BaseRequest
	
	// appId
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetAuthMiniProgramAppListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthMiniProgramAppListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MiniProgramAppId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAuthMiniProgramAppListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthMiniProgramAppListResponseParams struct {
	// 小程序列表
	MiniProgramList []*AuthMiniProgramAppInfo `json:"MiniProgramList,omitnil,omitempty" name:"MiniProgramList"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAuthMiniProgramAppListResponse struct {
	*tchttp.BaseResponse
	Response *GetAuthMiniProgramAppListResponseParams `json:"Response"`
}

func (r *GetAuthMiniProgramAppListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthMiniProgramAppListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBatchProductionsListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetBatchProductionsListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetBatchProductionsListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBatchProductionsListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBatchProductionsListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBatchProductionsListResponseParams struct {
	// 返回详情信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchProductions []*BatchProductionInfo `json:"BatchProductions,omitnil,omitempty" name:"BatchProductions"`

	// 返回数量。
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetBatchProductionsListResponse struct {
	*tchttp.BaseResponse
	Response *GetBatchProductionsListResponseParams `json:"Response"`
}

func (r *GetBatchProductionsListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBatchProductionsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCOSURLRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 文件大小
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`
}

type GetCOSURLRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 文件大小
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`
}

func (r *GetCOSURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCOSURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "FileSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCOSURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCOSURLResponseParams struct {
	// 固件URL
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCOSURLResponse struct {
	*tchttp.BaseResponse
	Response *GetCOSURLResponseParams `json:"Response"`
}

func (r *GetCOSURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCOSURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceListRequestParams struct {
	// 需要查看设备列表的产品ID, -1代表ProjectId来筛选
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备。传"None-FirmwareVersion"查询无版本号的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 项目ID。产品 ID 为 -1 时，该参数必填
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 每次请求的Filters的上限为10，Filter.Values的上限为1。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type GetDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 需要查看设备列表的产品ID, -1代表ProjectId来筛选
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备。传"None-FirmwareVersion"查询无版本号的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 项目ID。产品 ID 为 -1 时，该参数必填
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 每次请求的Filters的上限为10，Filter.Values的上限为1。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *GetDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FirmwareVersion")
	delete(f, "DeviceName")
	delete(f, "ProjectId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceListResponseParams struct {
	// 返回的设备列表, 注意列表设备的 DevicePsk 为空, 要获取设备的 DevicePsk 请使用 DescribeDevice
	Devices []*DeviceInfo `json:"Devices,omitnil,omitempty" name:"Devices"`

	// 产品下的设备总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceListResponseParams `json:"Response"`
}

func (r *GetDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceLocationHistoryRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitnil,omitempty" name:"CoordinateType"`
}

type GetDeviceLocationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 查询起始时间，Unix时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间，Unix时间，单位为毫秒
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 坐标类型
	CoordinateType *int64 `json:"CoordinateType,omitnil,omitempty" name:"CoordinateType"`
}

func (r *GetDeviceLocationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLocationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "CoordinateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceLocationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceLocationHistoryResponseParams struct {
	// 历史位置列表
	Positions []*PositionItem `json:"Positions,omitnil,omitempty" name:"Positions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDeviceLocationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceLocationHistoryResponseParams `json:"Response"`
}

func (r *GetDeviceLocationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLocationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceSumStatisticsRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品id列表，长度为0则拉取项目内全部产品
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`
}

type GetDeviceSumStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品id列表，长度为0则拉取项目内全部产品
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`
}

func (r *GetDeviceSumStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceSumStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProductIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceSumStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceSumStatisticsResponseParams struct {
	// 激活设备总数
	ActivationCount *int64 `json:"ActivationCount,omitnil,omitempty" name:"ActivationCount"`

	// 在线设备总数
	OnlineCount *int64 `json:"OnlineCount,omitnil,omitempty" name:"OnlineCount"`

	// 前一天激活设备数
	ActivationBeforeDay *int64 `json:"ActivationBeforeDay,omitnil,omitempty" name:"ActivationBeforeDay"`

	// 前一天活跃设备数
	ActiveBeforeDay *int64 `json:"ActiveBeforeDay,omitnil,omitempty" name:"ActiveBeforeDay"`

	// 前一周激活设备数
	ActivationWeekDayCount *int64 `json:"ActivationWeekDayCount,omitnil,omitempty" name:"ActivationWeekDayCount"`

	// 前一周活跃设备数
	ActiveWeekDayCount *int64 `json:"ActiveWeekDayCount,omitnil,omitempty" name:"ActiveWeekDayCount"`

	// 上一周激活设备数
	ActivationBeforeWeekDayCount *int64 `json:"ActivationBeforeWeekDayCount,omitnil,omitempty" name:"ActivationBeforeWeekDayCount"`

	// 上一周活跃设备数
	ActiveBeforeWeekDayCount *int64 `json:"ActiveBeforeWeekDayCount,omitnil,omitempty" name:"ActiveBeforeWeekDayCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDeviceSumStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceSumStatisticsResponseParams `json:"Response"`
}

func (r *GetDeviceSumStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceSumStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFamilyDeviceUserListRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type GetFamilyDeviceUserListRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *GetFamilyDeviceUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFamilyDeviceUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFamilyDeviceUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFamilyDeviceUserListResponseParams struct {
	// 设备的用户列表
	UserList []*DeviceUser `json:"UserList,omitnil,omitempty" name:"UserList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFamilyDeviceUserListResponse struct {
	*tchttp.BaseResponse
	Response *GetFamilyDeviceUserListResponseParams `json:"Response"`
}

func (r *GetFamilyDeviceUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFamilyDeviceUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGatewaySubDeviceListRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetGatewaySubDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备名称
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetGatewaySubDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGatewaySubDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGatewaySubDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGatewaySubDeviceListResponseParams struct {
	// 设备的总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 设备列表
	DeviceList *FamilySubDevice `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetGatewaySubDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *GetGatewaySubDeviceListResponseParams `json:"Response"`
}

func (r *GetGatewaySubDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGatewaySubDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLoRaGatewayListRequestParams struct {
	// 是否是社区网关
	IsCommunity *bool `json:"IsCommunity,omitnil,omitempty" name:"IsCommunity"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetLoRaGatewayListRequest struct {
	*tchttp.BaseRequest
	
	// 是否是社区网关
	IsCommunity *bool `json:"IsCommunity,omitnil,omitempty" name:"IsCommunity"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetLoRaGatewayListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLoRaGatewayListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsCommunity")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLoRaGatewayListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLoRaGatewayListResponseParams struct {
	// 返回总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 返回详情项
	Gateways []*LoRaGatewayItem `json:"Gateways,omitnil,omitempty" name:"Gateways"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLoRaGatewayListResponse struct {
	*tchttp.BaseResponse
	Response *GetLoRaGatewayListResponseParams `json:"Response"`
}

func (r *GetLoRaGatewayListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLoRaGatewayListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPositionSpaceListRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetPositionSpaceListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 翻页偏移量，0起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大返回结果数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetPositionSpaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPositionSpaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPositionSpaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPositionSpaceListResponseParams struct {
	// 位置空间列表
	List []*PositionSpaceInfo `json:"List,omitnil,omitempty" name:"List"`

	// 位置空间数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPositionSpaceListResponse struct {
	*tchttp.BaseResponse
	Response *GetPositionSpaceListResponseParams `json:"Response"`
}

func (r *GetPositionSpaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPositionSpaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProjectListRequestParams struct {
	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按项目ID搜索
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 按产品ID搜索
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 加载 ProductCount、DeviceCount、ApplicationCount，可选值：ProductCount、DeviceCount、ApplicationCount，可多选
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 按项目名称搜索
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type GetProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按项目ID搜索
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 按产品ID搜索
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 加载 ProductCount、DeviceCount、ApplicationCount，可选值：ProductCount、DeviceCount、ApplicationCount，可多选
	Includes []*string `json:"Includes,omitnil,omitempty" name:"Includes"`

	// 按项目名称搜索
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

func (r *GetProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "ProjectId")
	delete(f, "ProductId")
	delete(f, "Includes")
	delete(f, "ProjectName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProjectListResponseParams struct {
	// 项目列表
	Projects []*ProjectEntryEx `json:"Projects,omitnil,omitempty" name:"Projects"`

	// 列表项个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetProjectListResponse struct {
	*tchttp.BaseResponse
	Response *GetProjectListResponseParams `json:"Response"`
}

func (r *GetProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStudioProductListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetStudioProductListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetStudioProductListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStudioProductListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DevStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetStudioProductListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStudioProductListResponseParams struct {
	// 产品列表
	Products []*ProductEntry `json:"Products,omitnil,omitempty" name:"Products"`

	// 产品数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetStudioProductListResponse struct {
	*tchttp.BaseResponse
	Response *GetStudioProductListResponseParams `json:"Response"`
}

func (r *GetStudioProductListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStudioProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTWeCallActiveStatusRequestParams struct {
	// 参数已弃用，不用传参
	//
	// Deprecated: MiniProgramAppId is deprecated.
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`

	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type GetTWeCallActiveStatusRequest struct {
	*tchttp.BaseRequest
	
	// 参数已弃用，不用传参
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`

	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *GetTWeCallActiveStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTWeCallActiveStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MiniProgramAppId")
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTWeCallActiveStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTWeCallActiveStatusResponseParams struct {
	// 激活状态
	TWeCallActiveInfos []*TWeCallActiveInfo `json:"TWeCallActiveInfos,omitnil,omitempty" name:"TWeCallActiveInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTWeCallActiveStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetTWeCallActiveStatusResponseParams `json:"Response"`
}

func (r *GetTWeCallActiveStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTWeCallActiveStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicRuleListRequestParams struct {
	// 请求的页数
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type GetTopicRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 请求的页数
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *GetTopicRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTopicRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicRuleListResponseParams struct {
	// 规则总数量
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 规则列表
	Rules []*TopicRuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTopicRuleListResponse struct {
	*tchttp.BaseResponse
	Response *GetTopicRuleListResponseParams `json:"Response"`
}

func (r *GetTopicRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWechatDeviceTicketRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 是否第三方小程序
	IsThirdApp *int64 `json:"IsThirdApp,omitnil,omitempty" name:"IsThirdApp"`

	// 模板ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 小程序APPID
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`
}

type GetWechatDeviceTicketRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 是否第三方小程序
	IsThirdApp *int64 `json:"IsThirdApp,omitnil,omitempty" name:"IsThirdApp"`

	// 模板ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 小程序APPID
	MiniProgramAppId *string `json:"MiniProgramAppId,omitnil,omitempty" name:"MiniProgramAppId"`
}

func (r *GetWechatDeviceTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWechatDeviceTicketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "IsThirdApp")
	delete(f, "ModelId")
	delete(f, "MiniProgramAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWechatDeviceTicketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWechatDeviceTicketResponseParams struct {
	// 微信设备信息
	WXDeviceInfo *WXDeviceInfo `json:"WXDeviceInfo,omitnil,omitempty" name:"WXDeviceInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWechatDeviceTicketResponse struct {
	*tchttp.BaseResponse
	Response *GetWechatDeviceTicketResponseParams `json:"Response"`
}

func (r *GetWechatDeviceTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWechatDeviceTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InheritCloudStorageUserRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 原始用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 目标用户ID
	ToUserId *string `json:"ToUserId,omitnil,omitempty" name:"ToUserId"`
}

type InheritCloudStorageUserRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 原始用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 目标用户ID
	ToUserId *string `json:"ToUserId,omitnil,omitempty" name:"ToUserId"`
}

func (r *InheritCloudStorageUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InheritCloudStorageUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	delete(f, "ToUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InheritCloudStorageUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InheritCloudStorageUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InheritCloudStorageUserResponse struct {
	*tchttp.BaseResponse
	Response *InheritCloudStorageUserResponseParams `json:"Response"`
}

func (r *InheritCloudStorageUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InheritCloudStorageUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceDetail struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例类型（0 公共实例 1 标准企业实例 2新企业实例3新公共实例）
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 地域字母缩写
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 区域全拼
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 支持设备总数
	TotalDeviceNum *int64 `json:"TotalDeviceNum,omitnil,omitempty" name:"TotalDeviceNum"`

	// 已注册设备数
	UsedDeviceNum *int64 `json:"UsedDeviceNum,omitnil,omitempty" name:"UsedDeviceNum"`

	// 项目数
	ProjectNum *int64 `json:"ProjectNum,omitnil,omitempty" name:"ProjectNum"`

	// 产品数
	ProductNum *int64 `json:"ProductNum,omitnil,omitempty" name:"ProductNum"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 过期时间，公共实例过期时间 0001-01-01T00:00:00Z，公共实例是永久有效
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 总设备数
	TotalDevice *int64 `json:"TotalDevice,omitnil,omitempty" name:"TotalDevice"`

	// 激活设备数
	ActivateDevice *int64 `json:"ActivateDevice,omitnil,omitempty" name:"ActivateDevice"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 实例状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 消息上下行配置TPS
	UpDownTPS *int64 `json:"UpDownTPS,omitnil,omitempty" name:"UpDownTPS"`

	// 当前消息上下行TPS
	UpDownCurrentTPS *int64 `json:"UpDownCurrentTPS,omitnil,omitempty" name:"UpDownCurrentTPS"`

	// 消息转发配置TPS
	ForwardTPS *int64 `json:"ForwardTPS,omitnil,omitempty" name:"ForwardTPS"`

	// 消息转发当前TPS
	ForwardCurrentTPS *int64 `json:"ForwardCurrentTPS,omitnil,omitempty" name:"ForwardCurrentTPS"`

	// 实例单元数
	CellNum *int64 `json:"CellNum,omitnil,omitempty" name:"CellNum"`

	// 实例Tag，企业实例必传
	BillingTag *string `json:"BillingTag,omitnil,omitempty" name:"BillingTag"`

	// 每日消息数
	EverydayFreeMessageCount *int64 `json:"EverydayFreeMessageCount,omitnil,omitempty" name:"EverydayFreeMessageCount"`

	// 最大在线设备数
	MaxDeviceOnlineCount *int64 `json:"MaxDeviceOnlineCount,omitnil,omitempty" name:"MaxDeviceOnlineCount"`
}

// Predefined struct for user
type InvokeAISearchServiceRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 自然语言查询
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 搜索结果总结的语言类型，支持的类型有：en-US、zh-CN、id-ID、th-TH
	SummaryLang *string `json:"SummaryLang,omitnil,omitempty" name:"SummaryLang"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type InvokeAISearchServiceRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 自然语言查询
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 搜索结果总结的语言类型，支持的类型有：en-US、zh-CN、id-ID、th-TH
	SummaryLang *string `json:"SummaryLang,omitnil,omitempty" name:"SummaryLang"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *InvokeAISearchServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeAISearchServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Query")
	delete(f, "SummaryLang")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeAISearchServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeAISearchServiceResponseParams struct {
	// 基于搜索结果的总结
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 视频结果集
	Targets []*TargetInfo `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 视频回放URL
	VideoURL *string `json:"VideoURL,omitnil,omitempty" name:"VideoURL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InvokeAISearchServiceResponse struct {
	*tchttp.BaseResponse
	Response *InvokeAISearchServiceResponseParams `json:"Response"`
}

func (r *InvokeAISearchServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeAISearchServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeCloudStorageAIServiceTaskRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 待分析云存的起始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 待分析云存的结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`
}

type InvokeCloudStorageAIServiceTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 待分析云存的起始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 待分析云存的结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`
}

func (r *InvokeCloudStorageAIServiceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCloudStorageAIServiceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ServiceType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ChannelId")
	delete(f, "Config")
	delete(f, "ROI")
	delete(f, "VideoURLs")
	delete(f, "CustomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeCloudStorageAIServiceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeCloudStorageAIServiceTaskResponseParams struct {
	// 任务是否执行完成
	Completed *bool `json:"Completed,omitnil,omitempty" name:"Completed"`

	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务信息
	TaskInfo *CloudStorageAIServiceTask `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InvokeCloudStorageAIServiceTaskResponse struct {
	*tchttp.BaseResponse
	Response *InvokeCloudStorageAIServiceTaskResponseParams `json:"Response"`
}

func (r *InvokeCloudStorageAIServiceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeCloudStorageAIServiceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeExternalSourceAIServiceTaskRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`
}

type InvokeExternalSourceAIServiceTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	// - `VideoToText`：视频语义理解
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 分析外部传入的视频 URL 列表，支持 HLS 点播（m3u8）及常见视频格式（mp4 等）
	VideoURLs []*string `json:"VideoURLs,omitnil,omitempty" name:"VideoURLs"`

	// 自定义任务 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`
}

func (r *InvokeExternalSourceAIServiceTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeExternalSourceAIServiceTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ServiceType")
	delete(f, "VideoURLs")
	delete(f, "CustomId")
	delete(f, "Config")
	delete(f, "ROI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeExternalSourceAIServiceTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeExternalSourceAIServiceTaskResponseParams struct {
	// 任务是否执行完成
	Completed *bool `json:"Completed,omitnil,omitempty" name:"Completed"`

	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务信息
	TaskInfo *CloudStorageAIServiceTask `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InvokeExternalSourceAIServiceTaskResponse struct {
	*tchttp.BaseResponse
	Response *InvokeExternalSourceAIServiceTaskResponseParams `json:"Response"`
}

func (r *InvokeExternalSourceAIServiceTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeExternalSourceAIServiceTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeTWeSeeRecognitionTaskRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 输入视频 / 图片的 URL
	InputURL *string `json:"InputURL,omitnil,omitempty" name:"InputURL"`

	// 自定义事件 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 是否保存该事件使其可被搜索
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 事件起始时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	StartTimeMs *uint64 `json:"StartTimeMs,omitnil,omitempty" name:"StartTimeMs"`

	// 事件结束时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	EndTimeMs *uint64 `json:"EndTimeMs,omitnil,omitempty" name:"EndTimeMs"`

	// 算法配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否自定义设备，为 true 时不检查设备存在性，默认为 false
	IsCustomDevice *bool `json:"IsCustomDevice,omitnil,omitempty" name:"IsCustomDevice"`

	// 输入类型。可选值：
	// 
	// - `video`：视频（默认值）
	// - `image`：图片
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 摘要服务质量。可选值：
	// 
	// - `minutely`：分钟级（默认值）
	// - `immediate`：立即
	SummaryQOS *string `json:"SummaryQOS,omitnil,omitempty" name:"SummaryQOS"`
}

type InvokeTWeSeeRecognitionTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 输入视频 / 图片的 URL
	InputURL *string `json:"InputURL,omitnil,omitempty" name:"InputURL"`

	// 自定义事件 ID
	CustomId *string `json:"CustomId,omitnil,omitempty" name:"CustomId"`

	// 是否保存该事件使其可被搜索
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 事件起始时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	StartTimeMs *uint64 `json:"StartTimeMs,omitnil,omitempty" name:"StartTimeMs"`

	// 事件结束时间事件起始时间（毫秒级 UNIX 时间戳，若不传则默认为接口调用时间）
	EndTimeMs *uint64 `json:"EndTimeMs,omitnil,omitempty" name:"EndTimeMs"`

	// 算法配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否自定义设备，为 true 时不检查设备存在性，默认为 false
	IsCustomDevice *bool `json:"IsCustomDevice,omitnil,omitempty" name:"IsCustomDevice"`

	// 输入类型。可选值：
	// 
	// - `video`：视频（默认值）
	// - `image`：图片
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 摘要服务质量。可选值：
	// 
	// - `minutely`：分钟级（默认值）
	// - `immediate`：立即
	SummaryQOS *string `json:"SummaryQOS,omitnil,omitempty" name:"SummaryQOS"`
}

func (r *InvokeTWeSeeRecognitionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeTWeSeeRecognitionTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "InputURL")
	delete(f, "CustomId")
	delete(f, "EnableSearch")
	delete(f, "StartTimeMs")
	delete(f, "EndTimeMs")
	delete(f, "Config")
	delete(f, "IsCustomDevice")
	delete(f, "InputType")
	delete(f, "SummaryQOS")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeTWeSeeRecognitionTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeTWeSeeRecognitionTaskResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务是否执行完成
	Completed *bool `json:"Completed,omitnil,omitempty" name:"Completed"`

	// 语义理解任务结果（仅当 Completed 为 true 时包含该出参）
	Result *VisionRecognitionResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InvokeTWeSeeRecognitionTaskResponse struct {
	*tchttp.BaseResponse
	Response *InvokeTWeSeeRecognitionTaskResponseParams `json:"Response"`
}

func (r *InvokeTWeSeeRecognitionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeTWeSeeRecognitionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IotApplication struct {
	// 应用 ID
	IotAppID *string `json:"IotAppID,omitnil,omitempty" name:"IotAppID"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开发模式
	DevMode *int64 `json:"DevMode,omitnil,omitempty" name:"DevMode"`

	// iOS 平台 AppKey
	IOSAppKey *string `json:"IOSAppKey,omitnil,omitempty" name:"IOSAppKey"`

	// iOS 平台 AppSecret
	IOSAppSecret *string `json:"IOSAppSecret,omitnil,omitempty" name:"IOSAppSecret"`

	// Android 平台 AppKey
	AndroidAppKey *string `json:"AndroidAppKey,omitnil,omitempty" name:"AndroidAppKey"`

	// Android 平台 AppSecret
	AndroidAppSecret *string `json:"AndroidAppSecret,omitnil,omitempty" name:"AndroidAppSecret"`

	// 绑定的产品列表，数据为：ProdcutID 数组 JSON 序列化后的字符串
	Products *string `json:"Products,omitnil,omitempty" name:"Products"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 信鸽推送APP ID
	PushSecretID *string `json:"PushSecretID,omitnil,omitempty" name:"PushSecretID"`

	// 信鸽推送SECRET KEY
	PushSecretKey *string `json:"PushSecretKey,omitnil,omitempty" name:"PushSecretKey"`

	// iOS平台推送环境
	PushEnvironment *string `json:"PushEnvironment,omitnil,omitempty" name:"PushEnvironment"`

	// 小程序平台 AppKey
	MiniProgramAppKey *string `json:"MiniProgramAppKey,omitnil,omitempty" name:"MiniProgramAppKey"`

	// 小程序平台 AppSecret
	MiniProgramAppSecret *string `json:"MiniProgramAppSecret,omitnil,omitempty" name:"MiniProgramAppSecret"`

	// TPNS服务iOS应用AccessID，TPNS全称为腾讯移动推送（Tencent Push Notification Service），详见：https://cloud.tencent.com/document/product/548
	TPNSiOSAccessID *string `json:"TPNSiOSAccessID,omitnil,omitempty" name:"TPNSiOSAccessID"`

	// TPNS服务iOS应用SecretKey
	TPNSiOSSecretKey *string `json:"TPNSiOSSecretKey,omitnil,omitempty" name:"TPNSiOSSecretKey"`

	// TPNS服务iOS应用推送环境
	TPNSiOSPushEnvironment *string `json:"TPNSiOSPushEnvironment,omitnil,omitempty" name:"TPNSiOSPushEnvironment"`

	// TPNS服务Android应用AccessID
	TPNSAndroidAccessID *string `json:"TPNSAndroidAccessID,omitnil,omitempty" name:"TPNSAndroidAccessID"`

	// TPNS服务Android应用SecretKey
	TPNSAndroidSecretKey *string `json:"TPNSAndroidSecretKey,omitnil,omitempty" name:"TPNSAndroidSecretKey"`

	// TPNS服务iOS应用所属地域，详细说明参见 ModifyApplication 同名入参。
	TPNSiOSRegion *string `json:"TPNSiOSRegion,omitnil,omitempty" name:"TPNSiOSRegion"`

	// TPNS服务Android应用所属地域，详细说明参见 ModifyApplication 同名入参。
	TPNSAndroidRegion *string `json:"TPNSAndroidRegion,omitnil,omitempty" name:"TPNSAndroidRegion"`

	// 自主短信配置APPID
	SelfSmsAppId *string `json:"SelfSmsAppId,omitnil,omitempty" name:"SelfSmsAppId"`

	// 自主短信配置APPKey
	SelfSmsAppKey *string `json:"SelfSmsAppKey,omitnil,omitempty" name:"SelfSmsAppKey"`

	// 自主短信配置签名
	SelfSmsSign *string `json:"SelfSmsSign,omitnil,omitempty" name:"SelfSmsSign"`

	// 自主短信配置模板ID
	SelfSmsTemplateId *int64 `json:"SelfSmsTemplateId,omitnil,omitempty" name:"SelfSmsTemplateId"`

	// 第三方小程序强提醒开关 0：关闭；1：开启
	WechatNotifyStatus *int64 `json:"WechatNotifyStatus,omitnil,omitempty" name:"WechatNotifyStatus"`

	// 互联互通产品ID列表
	InterconnectionProducts *string `json:"InterconnectionProducts,omitnil,omitempty" name:"InterconnectionProducts"`
}

type LicenseServiceNumInfo struct {
	// 服务类型
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 授权总数
	TotalNum *int64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 已使用授权数
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// TWeCall激活码
	TWeCallLicense []*TWeCallLicenseInfo `json:"TWeCallLicense,omitnil,omitempty" name:"TWeCallLicense"`
}

// Predefined struct for user
type ListEventHistoryRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 搜索的事件类型：alert 表示告警，fault 表示故障，info 表示信息，为空则表示查询上述所有类型事件
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 搜索上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 单次获取的历史数据项目的最大数量, 缺省10
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`
}

type ListEventHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 搜索的事件类型：alert 表示告警，fault 表示故障，info 表示信息，为空则表示查询上述所有类型事件
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 起始时间（Unix 时间戳，秒级）, 为0 表示 当前时间 - 24h
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（Unix 时间戳，秒级）, 为0 表示当前时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 搜索上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 单次获取的历史数据项目的最大数量, 缺省10
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 事件标识符，可以用来指定查询特定的事件，如果不指定，则查询所有事件。
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`
}

func (r *ListEventHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Context")
	delete(f, "Size")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEventHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEventHistoryResponseParams struct {
	// 搜索上下文, 用作查询游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 搜索结果数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 搜索结果是否已经结束
	Listover *bool `json:"Listover,omitnil,omitempty" name:"Listover"`

	// 搜集结果集
	EventHistory []*EventHistoryItem `json:"EventHistory,omitnil,omitempty" name:"EventHistory"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEventHistoryResponse struct {
	*tchttp.BaseResponse
	Response *ListEventHistoryResponseParams `json:"Response"`
}

func (r *ListEventHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEventHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListFirmwaresRequestParams struct {
	// 获取的页数
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ListFirmwaresRequest struct {
	*tchttp.BaseRequest
	
	// 获取的页数
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 分页的大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 搜索过滤条件
	Filters []*SearchKeyword `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ListFirmwaresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFirmwaresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "ProductID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListFirmwaresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListFirmwaresResponseParams struct {
	// 固件总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 固件列表
	Firmwares []*FirmwareInfo `json:"Firmwares,omitnil,omitempty" name:"Firmwares"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListFirmwaresResponse struct {
	*tchttp.BaseResponse
	Response *ListFirmwaresResponseParams `json:"Response"`
}

func (r *ListFirmwaresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListFirmwaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

type ListTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`
}

func (r *ListTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopicPolicyResponseParams struct {
	// Topic列表
	Topics []*TopicItem `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ListTopicPolicyResponseParams `json:"Response"`
}

func (r *ListTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoRaFrequencyEntry struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitnil,omitempty" name:"FreqId"`

	// 频点名称
	FreqName *string `json:"FreqName,omitnil,omitempty" name:"FreqName"`

	// 频点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitnil,omitempty" name:"ChannelsDataUp"`

	// 数据下行信道RX1
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitnil,omitempty" name:"ChannelsDataRX1"`

	// 数据下行信道RX2
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitnil,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitnil,omitempty" name:"ChannelsJoinUp"`

	// 入网下行RX1信道
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitnil,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行RX2信道
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitnil,omitempty" name:"ChannelsJoinRX2"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type LoRaGatewayItem struct {
	// LoRa 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 是否是公开网关
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 网关描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 网关名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 网关位置信息
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 网关位置详情
	PositionDetails *string `json:"PositionDetails,omitnil,omitempty" name:"PositionDetails"`

	// LoRa 网关位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 最后更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 最后上报时间
	LastSeenAt *string `json:"LastSeenAt,omitnil,omitempty" name:"LastSeenAt"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitnil,omitempty" name:"FrequencyId"`
}

type LoRaGatewayLocation struct {
	// 纬度
	Latitude *float64 `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// 精度
	Longitude *float64 `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 准确度
	Accuracy *float64 `json:"Accuracy,omitnil,omitempty" name:"Accuracy"`

	// 海拔
	Altitude *float64 `json:"Altitude,omitnil,omitempty" name:"Altitude"`
}

// Predefined struct for user
type ModifyApplicationRequestParams struct {
	// 应用ID
	IotAppID *string `json:"IotAppID,omitnil,omitempty" name:"IotAppID"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 关联的产品
	Products *string `json:"Products,omitnil,omitempty" name:"Products"`

	// 信鸽推送APP ID
	PushSecretID *string `json:"PushSecretID,omitnil,omitempty" name:"PushSecretID"`

	// 信鸽推送SECRET KEY
	PushSecretKey *string `json:"PushSecretKey,omitnil,omitempty" name:"PushSecretKey"`

	// iOS平台推送环境
	PushEnvironment *string `json:"PushEnvironment,omitnil,omitempty" name:"PushEnvironment"`

	// TPNS服务iOS应用AccessID，TPNS全称为腾讯移动推送（Tencent Push Notification Service），详见：https://cloud.tencent.com/document/product/548
	TPNSiOSAccessID *string `json:"TPNSiOSAccessID,omitnil,omitempty" name:"TPNSiOSAccessID"`

	// TPNS服务iOS应用SecretKey
	TPNSiOSSecretKey *string `json:"TPNSiOSSecretKey,omitnil,omitempty" name:"TPNSiOSSecretKey"`

	// TPNS服务iOS应用推送环境
	TPNSiOSPushEnvironment *string `json:"TPNSiOSPushEnvironment,omitnil,omitempty" name:"TPNSiOSPushEnvironment"`

	// TPNS服务Android应用AccessID
	TPNSAndroidAccessID *string `json:"TPNSAndroidAccessID,omitnil,omitempty" name:"TPNSAndroidAccessID"`

	// TPNS服务Android应用SecretKey
	TPNSAndroidSecretKey *string `json:"TPNSAndroidSecretKey,omitnil,omitempty" name:"TPNSAndroidSecretKey"`

	// TPNS服务iOS应用所属地域，广州：ap-guangzhou，上海：ap-shanghai，中国香港：ap-hongkong，新加坡：ap-singapore。
	TPNSiOSRegion *string `json:"TPNSiOSRegion,omitnil,omitempty" name:"TPNSiOSRegion"`

	// TPNS服务Android应用所属地域，广州：ap-guangzhou，上海：ap-shanghai，中国香港：ap-hongkong，新加坡：ap-singapore。
	TPNSAndroidRegion *string `json:"TPNSAndroidRegion,omitnil,omitempty" name:"TPNSAndroidRegion"`

	// TurnKey小程序托管
	TurnKeySwitch *int64 `json:"TurnKeySwitch,omitnil,omitempty" name:"TurnKeySwitch"`
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	IotAppID *string `json:"IotAppID,omitnil,omitempty" name:"IotAppID"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 关联的产品
	Products *string `json:"Products,omitnil,omitempty" name:"Products"`

	// 信鸽推送APP ID
	PushSecretID *string `json:"PushSecretID,omitnil,omitempty" name:"PushSecretID"`

	// 信鸽推送SECRET KEY
	PushSecretKey *string `json:"PushSecretKey,omitnil,omitempty" name:"PushSecretKey"`

	// iOS平台推送环境
	PushEnvironment *string `json:"PushEnvironment,omitnil,omitempty" name:"PushEnvironment"`

	// TPNS服务iOS应用AccessID，TPNS全称为腾讯移动推送（Tencent Push Notification Service），详见：https://cloud.tencent.com/document/product/548
	TPNSiOSAccessID *string `json:"TPNSiOSAccessID,omitnil,omitempty" name:"TPNSiOSAccessID"`

	// TPNS服务iOS应用SecretKey
	TPNSiOSSecretKey *string `json:"TPNSiOSSecretKey,omitnil,omitempty" name:"TPNSiOSSecretKey"`

	// TPNS服务iOS应用推送环境
	TPNSiOSPushEnvironment *string `json:"TPNSiOSPushEnvironment,omitnil,omitempty" name:"TPNSiOSPushEnvironment"`

	// TPNS服务Android应用AccessID
	TPNSAndroidAccessID *string `json:"TPNSAndroidAccessID,omitnil,omitempty" name:"TPNSAndroidAccessID"`

	// TPNS服务Android应用SecretKey
	TPNSAndroidSecretKey *string `json:"TPNSAndroidSecretKey,omitnil,omitempty" name:"TPNSAndroidSecretKey"`

	// TPNS服务iOS应用所属地域，广州：ap-guangzhou，上海：ap-shanghai，中国香港：ap-hongkong，新加坡：ap-singapore。
	TPNSiOSRegion *string `json:"TPNSiOSRegion,omitnil,omitempty" name:"TPNSiOSRegion"`

	// TPNS服务Android应用所属地域，广州：ap-guangzhou，上海：ap-shanghai，中国香港：ap-hongkong，新加坡：ap-singapore。
	TPNSAndroidRegion *string `json:"TPNSAndroidRegion,omitnil,omitempty" name:"TPNSAndroidRegion"`

	// TurnKey小程序托管
	TurnKeySwitch *int64 `json:"TurnKeySwitch,omitnil,omitempty" name:"TurnKeySwitch"`
}

func (r *ModifyApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IotAppID")
	delete(f, "AppName")
	delete(f, "Description")
	delete(f, "Products")
	delete(f, "PushSecretID")
	delete(f, "PushSecretKey")
	delete(f, "PushEnvironment")
	delete(f, "TPNSiOSAccessID")
	delete(f, "TPNSiOSSecretKey")
	delete(f, "TPNSiOSPushEnvironment")
	delete(f, "TPNSAndroidAccessID")
	delete(f, "TPNSAndroidSecretKey")
	delete(f, "TPNSiOSRegion")
	delete(f, "TPNSAndroidRegion")
	delete(f, "TurnKeySwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationResponseParams struct {
	// 应用信息
	Application *IotApplication `json:"Application,omitnil,omitempty" name:"Application"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationResponseParams `json:"Response"`
}

func (r *ModifyApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudStorageAIServiceCallbackRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 推送类型。可选值：
	// - `http`：HTTP 回调
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// HTTP 回调 URL
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// HTTP 回调鉴权 Token
	CallbackToken *string `json:"CallbackToken,omitnil,omitempty" name:"CallbackToken"`
}

type ModifyCloudStorageAIServiceCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 推送类型。可选值：
	// - `http`：HTTP 回调
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// HTTP 回调 URL
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// HTTP 回调鉴权 Token
	CallbackToken *string `json:"CallbackToken,omitnil,omitempty" name:"CallbackToken"`
}

func (r *ModifyCloudStorageAIServiceCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudStorageAIServiceCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Type")
	delete(f, "CallbackUrl")
	delete(f, "CallbackToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudStorageAIServiceCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudStorageAIServiceCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudStorageAIServiceCallbackResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudStorageAIServiceCallbackResponseParams `json:"Response"`
}

func (r *ModifyCloudStorageAIServiceCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudStorageAIServiceCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudStorageAIServiceRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 视频分析启用状态
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type ModifyCloudStorageAIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 视频分析启用状态
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 视频分析识别区域
	ROI *string `json:"ROI,omitnil,omitempty" name:"ROI"`

	// 视频分析配置参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ModifyCloudStorageAIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudStorageAIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ServiceType")
	delete(f, "Enabled")
	delete(f, "ROI")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudStorageAIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudStorageAIServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudStorageAIServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudStorageAIServiceResponseParams `json:"Response"`
}

func (r *ModifyCloudStorageAIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudStorageAIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFenceBindRequestParams struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type ModifyFenceBindRequest struct {
	*tchttp.BaseRequest
	
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 围栏绑定的产品列表
	Items []*FenceBindProductItem `json:"Items,omitnil,omitempty" name:"Items"`
}

func (r *ModifyFenceBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFenceBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FenceId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFenceBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFenceBindResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFenceBindResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFenceBindResponseParams `json:"Response"`
}

func (r *ModifyFenceBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFenceBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaFrequencyRequestParams struct {
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitnil,omitempty" name:"FreqId"`

	// 频点名称
	FreqName *string `json:"FreqName,omitnil,omitempty" name:"FreqName"`

	// 频点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitnil,omitempty" name:"ChannelsDataUp"`

	// 数据下行信道RX1
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitnil,omitempty" name:"ChannelsDataRX1"`

	// 数据下行信道RX2
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitnil,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitnil,omitempty" name:"ChannelsJoinUp"`

	// 入网下行信道RX1
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitnil,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行信道RX2
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitnil,omitempty" name:"ChannelsJoinRX2"`
}

type ModifyLoRaFrequencyRequest struct {
	*tchttp.BaseRequest
	
	// 频点唯一ID
	FreqId *string `json:"FreqId,omitnil,omitempty" name:"FreqId"`

	// 频点名称
	FreqName *string `json:"FreqName,omitnil,omitempty" name:"FreqName"`

	// 频点描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据上行信道
	ChannelsDataUp []*uint64 `json:"ChannelsDataUp,omitnil,omitempty" name:"ChannelsDataUp"`

	// 数据下行信道RX1
	ChannelsDataRX1 []*uint64 `json:"ChannelsDataRX1,omitnil,omitempty" name:"ChannelsDataRX1"`

	// 数据下行信道RX2
	ChannelsDataRX2 []*uint64 `json:"ChannelsDataRX2,omitnil,omitempty" name:"ChannelsDataRX2"`

	// 入网上行信道
	ChannelsJoinUp []*uint64 `json:"ChannelsJoinUp,omitnil,omitempty" name:"ChannelsJoinUp"`

	// 入网下行信道RX1
	ChannelsJoinRX1 []*uint64 `json:"ChannelsJoinRX1,omitnil,omitempty" name:"ChannelsJoinRX1"`

	// 入网下行信道RX2
	ChannelsJoinRX2 []*uint64 `json:"ChannelsJoinRX2,omitnil,omitempty" name:"ChannelsJoinRX2"`
}

func (r *ModifyLoRaFrequencyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaFrequencyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FreqId")
	delete(f, "FreqName")
	delete(f, "Description")
	delete(f, "ChannelsDataUp")
	delete(f, "ChannelsDataRX1")
	delete(f, "ChannelsDataRX2")
	delete(f, "ChannelsJoinUp")
	delete(f, "ChannelsJoinRX1")
	delete(f, "ChannelsJoinRX2")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoRaFrequencyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaFrequencyResponseParams struct {
	// 频点信息
	Data *LoRaFrequencyEntry `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoRaFrequencyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoRaFrequencyResponseParams `json:"Response"`
}

func (r *ModifyLoRaFrequencyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaFrequencyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaGatewayRequestParams struct {
	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// LoRa网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// LoRa网关位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// LoRa网关名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否公开可见
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 位置信息
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitnil,omitempty" name:"PositionDetails"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitnil,omitempty" name:"FrequencyId"`
}

type ModifyLoRaGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// LoRa网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// LoRa网关位置坐标
	Location *LoRaGatewayLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// LoRa网关名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否公开可见
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 位置信息
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// 位置详情
	PositionDetails *string `json:"PositionDetails,omitnil,omitempty" name:"PositionDetails"`

	// 频点ID
	FrequencyId *string `json:"FrequencyId,omitnil,omitempty" name:"FrequencyId"`
}

func (r *ModifyLoRaGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "GatewayId")
	delete(f, "Location")
	delete(f, "Name")
	delete(f, "IsPublic")
	delete(f, "Position")
	delete(f, "PositionDetails")
	delete(f, "FrequencyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoRaGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoRaGatewayResponseParams struct {
	// 返回网关数据
	Gateway *LoRaGatewayItem `json:"Gateway,omitnil,omitempty" name:"Gateway"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoRaGatewayResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoRaGatewayResponseParams `json:"Response"`
}

func (r *ModifyLoRaGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoRaGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelDefinitionRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 数据模板定义
	ModelSchema *string `json:"ModelSchema,omitnil,omitempty" name:"ModelSchema"`
}

type ModifyModelDefinitionRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 数据模板定义
	ModelSchema *string `json:"ModelSchema,omitnil,omitempty" name:"ModelSchema"`
}

func (r *ModifyModelDefinitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelDefinitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ModelSchema")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelDefinitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelDefinitionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelDefinitionResponseParams `json:"Response"`
}

func (r *ModifyModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelDefinitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionFenceRequestParams struct {

}

type ModifyPositionFenceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyPositionFenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionFenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPositionFenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionFenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPositionFenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPositionFenceResponseParams `json:"Response"`
}

func (r *ModifyPositionFenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionFenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionSpaceRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 位置空间名称
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 授权类型
	AuthorizeType *int64 `json:"AuthorizeType,omitnil,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitnil,omitempty" name:"ProductIdList"`

	// 位置空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`
}

type ModifyPositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 位置空间名称
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 授权类型
	AuthorizeType *int64 `json:"AuthorizeType,omitnil,omitempty" name:"AuthorizeType"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitnil,omitempty" name:"ProductIdList"`

	// 位置空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 缩略图
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`
}

func (r *ModifyPositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "SpaceName")
	delete(f, "AuthorizeType")
	delete(f, "ProductIdList")
	delete(f, "Description")
	delete(f, "Icon")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPositionSpaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPositionSpaceResponseParams `json:"Response"`
}

func (r *ModifyPositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductCloudStorageAIServiceRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 开通状态
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ModifyProductCloudStorageAIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 开通状态
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ModifyProductCloudStorageAIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductCloudStorageAIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProductCloudStorageAIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProductCloudStorageAIServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProductCloudStorageAIServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProductCloudStorageAIServiceResponseParams `json:"Response"`
}

func (r *ModifyProductCloudStorageAIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProductCloudStorageAIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitnil,omitempty" name:"ProjectDesc"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitnil,omitempty" name:"ProjectDesc"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProjectName")
	delete(f, "ProjectDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 项目详情
	Project *ProjectEntry `json:"Project,omitnil,omitempty" name:"Project"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpacePropertyRequestParams struct {
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品属性
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type ModifySpacePropertyRequest struct {
	*tchttp.BaseRequest
	
	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 产品Id
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品属性
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *ModifySpacePropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpacePropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "ProductId")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpacePropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpacePropertyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySpacePropertyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpacePropertyResponseParams `json:"Response"`
}

func (r *ModifySpacePropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpacePropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStudioProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitnil,omitempty" name:"ProductDesc"`

	// 模型ID
	ModuleId *int64 `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 是否打开二进制转Json功能, 取值为字符串 true/false
	EnableProductScript *string `json:"EnableProductScript,omitnil,omitempty" name:"EnableProductScript"`

	// 传1或者2；1代表强踢，2代表非强踢。传其它值不做任何处理
	BindStrategy *uint64 `json:"BindStrategy,omitnil,omitempty" name:"BindStrategy"`
}

type ModifyStudioProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitnil,omitempty" name:"ProductDesc"`

	// 模型ID
	ModuleId *int64 `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 是否打开二进制转Json功能, 取值为字符串 true/false
	EnableProductScript *string `json:"EnableProductScript,omitnil,omitempty" name:"EnableProductScript"`

	// 传1或者2；1代表强踢，2代表非强踢。传其它值不做任何处理
	BindStrategy *uint64 `json:"BindStrategy,omitnil,omitempty" name:"BindStrategy"`
}

func (r *ModifyStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "ProductName")
	delete(f, "ProductDesc")
	delete(f, "ModuleId")
	delete(f, "EnableProductScript")
	delete(f, "BindStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStudioProductResponseParams struct {
	// 产品描述
	Product *ProductEntry `json:"Product,omitnil,omitempty" name:"Product"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStudioProductResponseParams `json:"Response"`
}

func (r *ModifyStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTWeSeeConfigRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 是否开启视频摘要，不传则不修改
	EnableSummary *bool `json:"EnableSummary,omitnil,omitempty" name:"EnableSummary"`

	// 是否开启视频搜索，不传则不修改
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 配置参数，不传则不修改
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

type ModifyTWeSeeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 是否开启视频摘要，不传则不修改
	EnableSummary *bool `json:"EnableSummary,omitnil,omitempty" name:"EnableSummary"`

	// 是否开启视频搜索，不传则不修改
	EnableSearch *bool `json:"EnableSearch,omitnil,omitempty" name:"EnableSearch"`

	// 配置参数，不传则不修改
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ModifyTWeSeeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTWeSeeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	delete(f, "ChannelId")
	delete(f, "EnableSummary")
	delete(f, "EnableSearch")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTWeSeeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTWeSeeConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTWeSeeConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTWeSeeConfigResponseParams `json:"Response"`
}

func (r *ModifyTWeSeeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTWeSeeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicPolicyRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 更新前Topic名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 更新后Topic名
	NewTopicName *string `json:"NewTopicName,omitnil,omitempty" name:"NewTopicName"`

	// Topic权限
	Privilege *uint64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`
}

type ModifyTopicPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 更新前Topic名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 更新后Topic名
	NewTopicName *string `json:"NewTopicName,omitnil,omitempty" name:"NewTopicName"`

	// Topic权限
	Privilege *uint64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`
}

func (r *ModifyTopicPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "TopicName")
	delete(f, "NewTopicName")
	delete(f, "Privilege")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicPolicyResponseParams `json:"Response"`
}

func (r *ModifyTopicPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRuleRequestParams struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 替换的规则包体
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitnil,omitempty" name:"TopicRulePayload"`
}

type ModifyTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 替换的规则包体
	TopicRulePayload *TopicRulePayload `json:"TopicRulePayload,omitnil,omitempty" name:"TopicRulePayload"`
}

func (r *ModifyTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	delete(f, "TopicRulePayload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicRuleResponseParams `json:"Response"`
}

func (r *ModifyTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PackageConsumeStat struct {
	// 云存套餐包id
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 云存套餐包名称
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 消耗个数
	Cnt *uint64 `json:"Cnt,omitnil,omitempty" name:"Cnt"`

	// 套餐包单价，单位分
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 消耗来源，1预付费
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`
}

type PackageConsumeTask struct {
	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务创始时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务状态，1待处理，2处理中，3已完成
	State *int64 `json:"State,omitnil,omitempty" name:"State"`
}

type PackageInfo struct {
	// 云存开启状态，0为未开启，2为正在生效，1为已过期
	// 注：这里只返回状态为0的数据
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 云存类型，1为全时云存，2为事件云存
	CSType *uint64 `json:"CSType,omitnil,omitempty" name:"CSType"`

	// 云存回看时长
	CSShiftDuration *uint64 `json:"CSShiftDuration,omitnil,omitempty" name:"CSShiftDuration"`

	// 云存套餐过期时间
	CSExpiredTime *int64 `json:"CSExpiredTime,omitnil,omitempty" name:"CSExpiredTime"`

	// 云存套餐创建时间
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 云存套餐更新时间
	UpdatedAt *int64 `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 套餐id
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 通道id
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 用户id
	CSUserId *string `json:"CSUserId,omitnil,omitempty" name:"CSUserId"`
}

// Predefined struct for user
type PauseTWeCallDeviceRequestParams struct {
	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type PauseTWeCallDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *PauseTWeCallDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseTWeCallDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseTWeCallDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseTWeCallDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseTWeCallDeviceResponse struct {
	*tchttp.BaseResponse
	Response *PauseTWeCallDeviceResponseParams `json:"Response"`
}

func (r *PauseTWeCallDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseTWeCallDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PositionFenceInfo struct {
	// 围栏信息
	GeoFence *PositionFenceItem `json:"GeoFence,omitnil,omitempty" name:"GeoFence"`

	// 围栏创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 围栏更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type PositionFenceItem struct {
	// 围栏Id
	FenceId *int64 `json:"FenceId,omitnil,omitempty" name:"FenceId"`

	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 围栏名称
	FenceName *string `json:"FenceName,omitnil,omitempty" name:"FenceName"`

	// 围栏描述
	FenceDesc *string `json:"FenceDesc,omitnil,omitempty" name:"FenceDesc"`

	// 围栏区域信息，采用 GeoJSON 格式
	FenceArea *string `json:"FenceArea,omitnil,omitempty" name:"FenceArea"`
}

type PositionItem struct {
	// 位置点的时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 位置点的经度
	Longitude *float64 `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// 位置点的纬度
	Latitude *float64 `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// 位置点的定位类型
	LocationType *string `json:"LocationType,omitnil,omitempty" name:"LocationType"`

	// 位置点的精度预估，单位为米
	Accuracy *float64 `json:"Accuracy,omitnil,omitempty" name:"Accuracy"`
}

type PositionSpaceInfo struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 位置空间Id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 位置空间名称
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 授权类型
	AuthorizeType *int64 `json:"AuthorizeType,omitnil,omitempty" name:"AuthorizeType"`

	// 描述备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 产品列表
	ProductIdList []*string `json:"ProductIdList,omitnil,omitempty" name:"ProductIdList"`

	// 缩略图
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 用户自定义地图缩放
	Zoom *uint64 `json:"Zoom,omitnil,omitempty" name:"Zoom"`
}

type ProductDevicesPositionItem struct {
	// 设备位置列表
	Items []*DevicePositionItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 产品标识
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备位置数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type ProductEntry struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品分组模板ID
	CategoryId *int64 `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 加密类型。1表示证书认证，2表示密钥认证，21表示TID认证-SE方式，22表示TID认证-软加固方式
	EncryptionType *string `json:"EncryptionType,omitnil,omitempty" name:"EncryptionType"`

	// 连接类型。如：
	// wifi、wifi-ble、cellular、5g、lorawan、ble、ethernet、wifi-ethernet、else、sub_zigbee、sub_ble、sub_433mhz、sub_else、sub_blemesh
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 数据协议 (1 使用物模型 2 为自定义类型)
	DataProtocol *int64 `json:"DataProtocol,omitnil,omitempty" name:"DataProtocol"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitnil,omitempty" name:"ProductDesc"`

	// 状态 如：all 全部, dev 开发中, audit 审核中 released 已发布
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 产品类型。如： 0 普通产品 ， 5 网关产品
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品ModuleId
	ModuleId *int64 `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 是否使用脚本进行二进制转json功能 可以取值 true / false
	EnableProductScript *string `json:"EnableProductScript,omitnil,omitempty" name:"EnableProductScript"`

	// 创建人 UinId
	CreateUserId *int64 `json:"CreateUserId,omitnil,omitempty" name:"CreateUserId"`

	// 创建者昵称
	CreatorNickName *string `json:"CreatorNickName,omitnil,omitempty" name:"CreatorNickName"`

	// 绑定策略（1：强踢；2：非强踢；0：表示无意义）
	BindStrategy *uint64 `json:"BindStrategy,omitnil,omitempty" name:"BindStrategy"`

	// 设备数量
	DeviceCount *int64 `json:"DeviceCount,omitnil,omitempty" name:"DeviceCount"`

	// 平均传输速率
	Rate *string `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 有效期
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// 互联互通标识
	IsInterconnection *int64 `json:"IsInterconnection,omitnil,omitempty" name:"IsInterconnection"`
}

type ProductModelDefinition struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 模型定义
	ModelDefine *string `json:"ModelDefine,omitnil,omitempty" name:"ModelDefine"`

	// 更新时间，秒级时间戳
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间，秒级时间戳
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 产品所属分类的模型快照（产品创建时刻的）
	CategoryModel *string `json:"CategoryModel,omitnil,omitempty" name:"CategoryModel"`

	// 产品的连接类型的模型
	NetTypeModel *string `json:"NetTypeModel,omitnil,omitempty" name:"NetTypeModel"`
}

type ProjectEntry struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitnil,omitempty" name:"ProjectDesc"`

	// 创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间，unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type ProjectEntryEx struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitnil,omitempty" name:"ProjectDesc"`

	// 项目创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目更新时间，unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 产品数量
	ProductCount *uint64 `json:"ProductCount,omitnil,omitempty" name:"ProductCount"`

	// NativeApp数量
	NativeAppCount *uint64 `json:"NativeAppCount,omitnil,omitempty" name:"NativeAppCount"`

	// WebApp数量
	WebAppCount *uint64 `json:"WebAppCount,omitnil,omitempty" name:"WebAppCount"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用数量
	ApplicationCount *uint64 `json:"ApplicationCount,omitnil,omitempty" name:"ApplicationCount"`

	// 设备注册总数
	DeviceCount *uint64 `json:"DeviceCount,omitnil,omitempty" name:"DeviceCount"`

	// 是否开通物联使能
	EnableOpenState *uint64 `json:"EnableOpenState,omitnil,omitempty" name:"EnableOpenState"`
}

// Predefined struct for user
type PublishBroadcastMessageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 消息内容
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 消息质量等级
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// ayload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
	PayloadEncoding *string `json:"PayloadEncoding,omitnil,omitempty" name:"PayloadEncoding"`
}

type PublishBroadcastMessageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 消息内容
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 消息质量等级
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// ayload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
	PayloadEncoding *string `json:"PayloadEncoding,omitnil,omitempty" name:"PayloadEncoding"`
}

func (r *PublishBroadcastMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishBroadcastMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Payload")
	delete(f, "Qos")
	delete(f, "PayloadEncoding")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishBroadcastMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishBroadcastMessageResponseParams struct {
	// 广播消息任务Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishBroadcastMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishBroadcastMessageResponseParams `json:"Response"`
}

func (r *PublishBroadcastMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishBroadcastMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishFirmwareUpdateMessageRequestParams struct {
	// 产品 ID。
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

type PublishFirmwareUpdateMessageRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID。
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 设备名称。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`
}

func (r *PublishFirmwareUpdateMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishFirmwareUpdateMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "DeviceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishFirmwareUpdateMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishFirmwareUpdateMessageResponseParams struct {
	// 请求状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishFirmwareUpdateMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishFirmwareUpdateMessageResponseParams `json:"Response"`
}

func (r *PublishFirmwareUpdateMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishFirmwareUpdateMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishMessageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 消息发往的主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 云端下发到设备的控制报文
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 消息服务质量等级，取值为0或1
	Qos *uint64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// Payload的内容编码格式，取值为base64或空。base64表示云端将接收到的base64编码后的报文再转换成二进制报文下发至设备，为空表示不作转换，透传下发至设备
	PayloadEncoding *string `json:"PayloadEncoding,omitnil,omitempty" name:"PayloadEncoding"`
}

type PublishMessageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 消息发往的主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 云端下发到设备的控制报文
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 消息服务质量等级，取值为0或1
	Qos *uint64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// Payload的内容编码格式，取值为base64或空。base64表示云端将接收到的base64编码后的报文再转换成二进制报文下发至设备，为空表示不作转换，透传下发至设备
	PayloadEncoding *string `json:"PayloadEncoding,omitnil,omitempty" name:"PayloadEncoding"`
}

func (r *PublishMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Topic")
	delete(f, "Payload")
	delete(f, "Qos")
	delete(f, "PayloadEncoding")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishMessageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishMessageResponseParams `json:"Response"`
}

func (r *PublishMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishRRPCMessageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 消息内容，utf8编码
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`
}

type PublishRRPCMessageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 消息内容，utf8编码
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`
}

func (r *PublishRRPCMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishRRPCMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Payload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishRRPCMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishRRPCMessageResponseParams struct {
	// RRPC消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageId *int64 `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 设备回复的消息内容，采用base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayloadBase64 *string `json:"PayloadBase64,omitnil,omitempty" name:"PayloadBase64"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishRRPCMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishRRPCMessageResponseParams `json:"Response"`
}

func (r *PublishRRPCMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishRRPCMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisteredDeviceNetTypeInfo struct {
	// 普通设备数
	NormalDeviceNum *int64 `json:"NormalDeviceNum,omitnil,omitempty" name:"NormalDeviceNum"`

	// 蓝牙设备数
	BluetoothDeviceNum *int64 `json:"BluetoothDeviceNum,omitnil,omitempty" name:"BluetoothDeviceNum"`
}

type RegisteredDeviceTypeInfo struct {
	// 已注册设备数
	NormalDeviceNum *int64 `json:"NormalDeviceNum,omitnil,omitempty" name:"NormalDeviceNum"`

	// 已注册网关数
	GatewayDeviceNum *int64 `json:"GatewayDeviceNum,omitnil,omitempty" name:"GatewayDeviceNum"`

	// 已注册子设备数
	SubDeviceNum *int64 `json:"SubDeviceNum,omitnil,omitempty" name:"SubDeviceNum"`

	// 已注册视频设备数
	VideoDeviceNum *int64 `json:"VideoDeviceNum,omitnil,omitempty" name:"VideoDeviceNum"`
}

// Predefined struct for user
type ReleaseStudioProductRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`
}

type ReleaseStudioProductRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`
}

func (r *ReleaseStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DevStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseStudioProductResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseStudioProductResponseParams `json:"Response"`
}

func (r *ReleaseStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserByRoomIdFromTRTCRequestParams struct {
	// 房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用户名称数组，数组元素不可重复，最长不超过 10 个。
	TRTCUserIds []*string `json:"TRTCUserIds,omitnil,omitempty" name:"TRTCUserIds"`
}

type RemoveUserByRoomIdFromTRTCRequest struct {
	*tchttp.BaseRequest
	
	// 房间id
	RoomId *string `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用户名称数组，数组元素不可重复，最长不超过 10 个。
	TRTCUserIds []*string `json:"TRTCUserIds,omitnil,omitempty" name:"TRTCUserIds"`
}

func (r *RemoveUserByRoomIdFromTRTCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByRoomIdFromTRTCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "TRTCUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserByRoomIdFromTRTCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserByRoomIdFromTRTCResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveUserByRoomIdFromTRTCResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserByRoomIdFromTRTCResponseParams `json:"Response"`
}

func (r *RemoveUserByRoomIdFromTRTCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByRoomIdFromTRTCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageAIServiceRequestParams struct {
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ResetCloudStorageAIServiceRequest struct {
	*tchttp.BaseRequest
	
	// 产品 ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 云存 AI 服务类型。可选值：
	// - `RealtimeObjectDetect`：目标检测
	// - `Highlight`：视频浓缩
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 通道 ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ResetCloudStorageAIServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageAIServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ServiceType")
	delete(f, "ChannelId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetCloudStorageAIServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageAIServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetCloudStorageAIServiceResponse struct {
	*tchttp.BaseResponse
	Response *ResetCloudStorageAIServiceResponseParams `json:"Response"`
}

func (r *ResetCloudStorageAIServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageAIServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageEventRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type ResetCloudStorageEventRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *ResetCloudStorageEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "UserId")
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetCloudStorageEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageEventResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetCloudStorageEventResponse struct {
	*tchttp.BaseResponse
	Response *ResetCloudStorageEventResponseParams `json:"Response"`
}

func (r *ResetCloudStorageEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 云存用户Id，为空则为默认云存空间。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ResetCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 通道ID 非NVR设备则不填 NVR设备则必填 默认为无
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 云存用户Id，为空则为默认云存空间。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ResetCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ChannelId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetCloudStorageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *ResetCloudStorageResponseParams `json:"Response"`
}

func (r *ResetCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetTWeCallDeviceRequestParams struct {
	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type ResetTWeCallDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *ResetTWeCallDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetTWeCallDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetTWeCallDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetTWeCallDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetTWeCallDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ResetTWeCallDeviceResponseParams `json:"Response"`
}

func (r *ResetTWeCallDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetTWeCallDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeWeCallDeviceRequestParams struct {
	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type ResumeWeCallDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备列表
	DeviceList []*TWeCallInfo `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *ResumeWeCallDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeWeCallDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeWeCallDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeWeCallDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeWeCallDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeWeCallDeviceResponseParams `json:"Response"`
}

func (r *ResumeWeCallDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeWeCallDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchKeyword struct {
	// 搜索条件的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 搜索条件的值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type SearchPositionSpaceRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 位置空间名字
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大获取数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type SearchPositionSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 位置空间名字
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 偏移量，从0开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大获取数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *SearchPositionSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPositionSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SpaceName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchPositionSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchPositionSpaceResponseParams struct {
	// 位置空间列表
	List []*PositionSpaceInfo `json:"List,omitnil,omitempty" name:"List"`

	// 符合条件的位置空间个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchPositionSpaceResponse struct {
	*tchttp.BaseResponse
	Response *SearchPositionSpaceResponseParams `json:"Response"`
}

func (r *SearchPositionSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchPositionSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchStudioProductRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 列表Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 产品Status
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 每次请求的Filters的上限为10，Filter.Values的上限为1。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type SearchStudioProductRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 列表Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 产品Status
	DevStatus *string `json:"DevStatus,omitnil,omitempty" name:"DevStatus"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 每次请求的Filters的上限为10，Filter.Values的上限为1。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *SearchStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchStudioProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProductName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DevStatus")
	delete(f, "ProductId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchStudioProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchStudioProductResponseParams struct {
	// 产品列表
	Products []*ProductEntry `json:"Products,omitnil,omitempty" name:"Products"`

	// 产品数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *SearchStudioProductResponseParams `json:"Response"`
}

func (r *SearchStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchStudioProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchTopicRuleRequestParams struct {
	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type SearchTopicRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

func (r *SearchTopicRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchTopicRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchTopicRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchTopicRuleResponseParams struct {
	// 搜索到的规则总数
	TotalCnt *int64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 规则信息列表
	Rules []*TopicRuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchTopicRuleResponse struct {
	*tchttp.BaseResponse
	Response *SearchTopicRuleResponseParams `json:"Response"`
}

func (r *SearchTopicRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchTopicRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TRTCParams struct {
	// TRTC入参: TRTC的实例ID
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// TRTC入参: 用户加入房间的ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// TRTC入参: 用户的签名用来鉴权
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`

	// TRTC入参: 加入的TRTC房间名称
	StrRoomId *string `json:"StrRoomId,omitnil,omitempty" name:"StrRoomId"`

	// TRTC入参: 校验TRTC的KEY
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`
}

type TWeCallActiveInfo struct {
	// 小程序ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModelId is deprecated.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Sn信息
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 类型
	PkgType *int64 `json:"PkgType,omitnil,omitempty" name:"PkgType"`
}

type TWeCallInfo struct {
	// Sn信息，SN格式：产品ID_设备名
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 小程序ID，参数已弃用，不用传参
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModelId is deprecated.
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 激活数
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ActiveNum is deprecated.
	ActiveNum *int64 `json:"ActiveNum,omitnil,omitempty" name:"ActiveNum"`
}

type TWeCallLicenseInfo struct {
	// voip类型
	TWeCallType *string `json:"TWeCallType,omitnil,omitempty" name:"TWeCallType"`

	// 总数
	TotalNum *int64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 已使用
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`
}

type TargetInfo struct {
	// 视频唯一ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 视频起始时间（毫秒级Unix时间戳）
	StartTimeMs *int64 `json:"StartTimeMs,omitnil,omitempty" name:"StartTimeMs"`

	// 视频结束时间（毫秒级Unix时间戳）
	EndTimeMs *int64 `json:"EndTimeMs,omitnil,omitempty" name:"EndTimeMs"`

	// 用户自定义事件ID，后续扩展使用
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 视频内容摘要
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 通道ID
	ChannelId *uint64 `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// 缩略图路径
	Thumbnail *string `json:"Thumbnail,omitnil,omitempty" name:"Thumbnail"`
}

type ThumbnailURLInfoList struct {
	// 缩略图访问地址
	ThumbnailURL *string `json:"ThumbnailURL,omitnil,omitempty" name:"ThumbnailURL"`

	// 缩略图访问地址的过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type TopicItem struct {
	// Topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Topic权限 , 1上报  2下发
	Privilege *uint64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`
}

type TopicRule struct {
	// 规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则的SQL语句，如： SELECT * FROM 'pid/dname/event'，然后对其进行base64编码，得：U0VMRUNUICogRlJPTSAncGlkL2RuYW1lL2V2ZW50Jw==
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 行为的JSON字符串。
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 是否禁用规则
	RuleDisabled *bool `json:"RuleDisabled,omitnil,omitempty" name:"RuleDisabled"`
}

type TopicRuleInfo struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 规则是否禁用
	RuleDisabled *bool `json:"RuleDisabled,omitnil,omitempty" name:"RuleDisabled"`
}

type TopicRulePayload struct {
	// 规则的SQL语句，如： SELECT * FROM 'pid/dname/event'，然后对其进行base64编码，得：U0VMRUNUICogRlJPTSAncGlkL2RuYW1lL2V2ZW50Jw==
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 行为的JSON字符串，大部分种类举例如下：
	// [
	// {
	// "republish": {
	// "topic": "TEST/test"
	// }
	// },
	// {
	// "forward": {
	// "api": "http://test.com:8080"
	// }
	// },
	// {
	// "ckafka": {
	// "instance": {
	// "id": "ckafka-test",
	// "name": ""
	// },
	// "topic": {
	// "id": "topic-test",
	// "name": "test"
	// },
	// "region": "gz"
	// }
	// },
	// {
	// "cmqqueue": {
	// "queuename": "queue-test-TEST",
	// "region": "gz"
	// }
	// },
	// {
	// "mysql": {
	// "instanceid": "cdb-test",
	// "region": "gz",
	// "username": "test",
	// "userpwd": "*****",
	// "dbname": "d_mqtt",
	// "tablename": "t_test",
	// "fieldpairs": [
	// {
	// "field": "test",
	// "value": "test"
	// }
	// ],
	// "devicetype": "CUSTOM"
	// }
	// }
	// ]
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否禁用规则
	RuleDisabled *bool `json:"RuleDisabled,omitnil,omitempty" name:"RuleDisabled"`
}

// Predefined struct for user
type TransferCloudStorageRequestParams struct {
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 已开通云存的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 未开通云存的设备名称
	ToDeviceName *string `json:"ToDeviceName,omitnil,omitempty" name:"ToDeviceName"`

	// 未开通云存的设备产品ID
	ToProductId *string `json:"ToProductId,omitnil,omitempty" name:"ToProductId"`
}

type TransferCloudStorageRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 已开通云存的设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 未开通云存的设备名称
	ToDeviceName *string `json:"ToDeviceName,omitnil,omitempty" name:"ToDeviceName"`

	// 未开通云存的设备产品ID
	ToProductId *string `json:"ToProductId,omitnil,omitempty" name:"ToProductId"`
}

func (r *TransferCloudStorageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferCloudStorageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ToDeviceName")
	delete(f, "ToProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferCloudStorageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferCloudStorageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TransferCloudStorageResponse struct {
	*tchttp.BaseResponse
	Response *TransferCloudStorageResponseParams `json:"Response"`
}

func (r *TransferCloudStorageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferCloudStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferTWeCallDeviceRequestParams struct {
	// sn信息，product_deviceName
	TransferInDevice *string `json:"TransferInDevice,omitnil,omitempty" name:"TransferInDevice"`

	// sn信息，product_deviceName
	TransferOutDevice *string `json:"TransferOutDevice,omitnil,omitempty" name:"TransferOutDevice"`
}

type TransferTWeCallDeviceRequest struct {
	*tchttp.BaseRequest
	
	// sn信息，product_deviceName
	TransferInDevice *string `json:"TransferInDevice,omitnil,omitempty" name:"TransferInDevice"`

	// sn信息，product_deviceName
	TransferOutDevice *string `json:"TransferOutDevice,omitnil,omitempty" name:"TransferOutDevice"`
}

func (r *TransferTWeCallDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferTWeCallDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TransferInDevice")
	delete(f, "TransferOutDevice")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferTWeCallDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferTWeCallDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TransferTWeCallDeviceResponse struct {
	*tchttp.BaseResponse
	Response *TransferTWeCallDeviceResponseParams `json:"Response"`
}

func (r *TransferTWeCallDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferTWeCallDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindDevicesRequestParams struct {
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名列表
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`
}

type UnbindDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关设备的产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 网关设备的设备名
	GatewayDeviceName *string `json:"GatewayDeviceName,omitnil,omitempty" name:"GatewayDeviceName"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名列表
	DeviceNames []*string `json:"DeviceNames,omitnil,omitempty" name:"DeviceNames"`
}

func (r *UnbindDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "GatewayDeviceName")
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindDevicesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindDevicesResponse struct {
	*tchttp.BaseResponse
	Response *UnbindDevicesResponseParams `json:"Response"`
}

func (r *UnbindDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindProductsRequestParams struct {
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 待解绑的子产品ID数组
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`
}

type UnbindProductsRequest struct {
	*tchttp.BaseRequest
	
	// 网关产品ID
	GatewayProductId *string `json:"GatewayProductId,omitnil,omitempty" name:"GatewayProductId"`

	// 待解绑的子产品ID数组
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`
}

func (r *UnbindProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayProductId")
	delete(f, "ProductIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindProductsResponseParams struct {
	// 绑定了待解绑的LoRa产品下的设备的网关设备列表
	GatewayDeviceNames []*string `json:"GatewayDeviceNames,omitnil,omitempty" name:"GatewayDeviceNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindProductsResponse struct {
	*tchttp.BaseResponse
	Response *UnbindProductsResponseParams `json:"Response"`
}

func (r *UnbindProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceTWeCallAuthorizeStatusRequestParams struct {
	// TweCall授权状态：0未授权，1已授权
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 微信用户的openId
	WechatOpenId *string `json:"WechatOpenId,omitnil,omitempty" name:"WechatOpenId"`
}

type UpdateDeviceTWeCallAuthorizeStatusRequest struct {
	*tchttp.BaseRequest
	
	// TweCall授权状态：0未授权，1已授权
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 产品ID
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 微信用户的openId
	WechatOpenId *string `json:"WechatOpenId,omitnil,omitempty" name:"WechatOpenId"`
}

func (r *UpdateDeviceTWeCallAuthorizeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceTWeCallAuthorizeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "WechatOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceTWeCallAuthorizeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceTWeCallAuthorizeStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDeviceTWeCallAuthorizeStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceTWeCallAuthorizeStatusResponseParams `json:"Response"`
}

func (r *UpdateDeviceTWeCallAuthorizeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceTWeCallAuthorizeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateRequestParams struct {
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitnil,omitempty" name:"DevicesItems"`

	// 1：启用；0：禁用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateDevicesEnableStateRequest struct {
	*tchttp.BaseRequest
	
	// 多个设备标识
	DevicesItems []*DevicesItem `json:"DevicesItems,omitnil,omitempty" name:"DevicesItems"`

	// 1：启用；0：禁用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *UpdateDevicesEnableStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicesEnableStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DevicesItems")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDevicesEnableStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDevicesEnableStateResponseParams struct {
	// 删除的结果代码
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 删除的结果信息
	ResultMessage *string `json:"ResultMessage,omitnil,omitempty" name:"ResultMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDevicesEnableStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDevicesEnableStateResponseParams `json:"Response"`
}

func (r *UpdateDevicesEnableStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicesEnableStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 固件新的版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 固件原版本号
	FirmwareOriVersion *string `json:"FirmwareOriVersion,omitnil,omitempty" name:"FirmwareOriVersion"`

	// 固件升级方式；0 静默升级 1 用户确认升级   不填默认静默升级
	UpgradeMethod *uint64 `json:"UpgradeMethod,omitnil,omitempty" name:"UpgradeMethod"`
}

type UpdateFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 固件新的版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 固件原版本号
	FirmwareOriVersion *string `json:"FirmwareOriVersion,omitnil,omitempty" name:"FirmwareOriVersion"`

	// 固件升级方式；0 静默升级 1 用户确认升级   不填默认静默升级
	UpgradeMethod *uint64 `json:"UpgradeMethod,omitnil,omitempty" name:"UpgradeMethod"`
}

func (r *UpdateFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "DeviceName")
	delete(f, "FirmwareVersion")
	delete(f, "FirmwareOriVersion")
	delete(f, "UpgradeMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFirmwareResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFirmwareResponseParams `json:"Response"`
}

func (r *UpdateFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFirmwareRequestParams struct {
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 固件的MD5值
	Md5sum *string `json:"Md5sum,omitnil,omitempty" name:"Md5sum"`

	// 固件的大小
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 固件名称
	FirmwareName *string `json:"FirmwareName,omitnil,omitempty" name:"FirmwareName"`

	// 固件描述
	FirmwareDescription *string `json:"FirmwareDescription,omitnil,omitempty" name:"FirmwareDescription"`

	// 固件升级模块；可选值 mcu|moudule
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 固件用户自定义配置信息
	FirmwareUserDefined *string `json:"FirmwareUserDefined,omitnil,omitempty" name:"FirmwareUserDefined"`
}

type UploadFirmwareRequest struct {
	*tchttp.BaseRequest
	
	// 产品ID
	ProductID *string `json:"ProductID,omitnil,omitempty" name:"ProductID"`

	// 固件版本号
	FirmwareVersion *string `json:"FirmwareVersion,omitnil,omitempty" name:"FirmwareVersion"`

	// 固件的MD5值
	Md5sum *string `json:"Md5sum,omitnil,omitempty" name:"Md5sum"`

	// 固件的大小
	FileSize *uint64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 固件名称
	FirmwareName *string `json:"FirmwareName,omitnil,omitempty" name:"FirmwareName"`

	// 固件描述
	FirmwareDescription *string `json:"FirmwareDescription,omitnil,omitempty" name:"FirmwareDescription"`

	// 固件升级模块；可选值 mcu|moudule
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 固件用户自定义配置信息
	FirmwareUserDefined *string `json:"FirmwareUserDefined,omitnil,omitempty" name:"FirmwareUserDefined"`
}

func (r *UploadFirmwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFirmwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductID")
	delete(f, "FirmwareVersion")
	delete(f, "Md5sum")
	delete(f, "FileSize")
	delete(f, "FirmwareName")
	delete(f, "FirmwareDescription")
	delete(f, "FwType")
	delete(f, "FirmwareUserDefined")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFirmwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFirmwareResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadFirmwareResponse struct {
	*tchttp.BaseResponse
	Response *UploadFirmwareResponseParams `json:"Response"`
}

func (r *UploadFirmwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFirmwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoLicenseEntity struct {
	// 激活码类型，取值范围如下：0_5_mbps、1_mbps、1_5_mbps、2_mbps
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 有效激活码总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 待使用的激活码数量
	UsedCount *int64 `json:"UsedCount,omitnil,omitempty" name:"UsedCount"`

	// 即将过期的激活码数量
	ExpiresSoonCount *int64 `json:"ExpiresSoonCount,omitnil,omitempty" name:"ExpiresSoonCount"`
}

type VisionRecognitionResult struct {
	// 任务状态（1：分析失败；2：下载/读取视频/图片失败；3：成功）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 识别到的目标类型。可能取值：
	// 
	// - `person`：人
	// - `vehicle`：车辆
	// - `dog`：狗
	// - `cat`：猫
	// - `fire`：火焰
	// - `smoke`：烟雾
	// - `package`：快递包裹
	// - `license_plate`：车牌
	DetectedClassifications []*string `json:"DetectedClassifications,omitnil,omitempty" name:"DetectedClassifications"`

	// 摘要文本
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 摘要文本（次选语言）
	AlternativeSummary *string `json:"AlternativeSummary,omitnil,omitempty" name:"AlternativeSummary"`
}

type WXDeviceInfo struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备信息
	WXIoTDeviceInfo *WXIoTDeviceInfo `json:"WXIoTDeviceInfo,omitnil,omitempty" name:"WXIoTDeviceInfo"`
}

type WXIoTDeviceInfo struct {
	// sn信息
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 票据
	SNTicket *string `json:"SNTicket,omitnil,omitempty" name:"SNTicket"`

	// 模板ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type WifiInfo struct {
	// mac地址
	MAC *string `json:"MAC,omitnil,omitempty" name:"MAC"`

	// 信号强度
	RSSI *int64 `json:"RSSI,omitnil,omitempty" name:"RSSI"`
}