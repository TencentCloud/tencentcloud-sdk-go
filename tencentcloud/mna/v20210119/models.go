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

package v20210119

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ActivateHardware struct {
	// 厂商名称
	Vendor *string `json:"Vendor,omitnil,omitempty" name:"Vendor"`

	// 设备SN序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 设备密钥
	DataKey *string `json:"DataKey,omitnil,omitempty" name:"DataKey"`

	// 接入环境。0：公有云网关；1：自有网关；2：公有云网关和自有网关。不填默认公有云网关。 具体含义： 公有云网关：即该设备只能接入公有云网关（就近接入） 自有网关：即该设备只能接入已经注册上线的自有网关（就近接入或固定ip接入） 公有云网关和自有网关：即该设备同时可以接入公有云网关和已经注册上线的自有网关（就近接入或固定ip接入）
	AccessScope *int64 `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// 当付费方为租户时，可选择租户license付费方式：
	// 0，月度授权
	// 1，永久授权
	// 若不传则默认为月度授权。
	// 当付费方为厂商时，此参数无效
	LicensePayMode *int64 `json:"LicensePayMode,omitnil,omitempty" name:"LicensePayMode"`

	// 设备分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 设备分组名称，预留参数，需要分组时传入GroupId
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备无流量包处理方式，0: 按量付费，1: 截断加速
	FlowTrunc *int64 `json:"FlowTrunc,omitnil,omitempty" name:"FlowTrunc"`

	// 激活后的设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

// Predefined struct for user
type ActivateHardwareRequestParams struct {
	// 待激活的设备列表
	Hardware []*ActivateHardware `json:"Hardware,omitnil,omitempty" name:"Hardware"`
}

type ActivateHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 待激活的设备列表
	Hardware []*ActivateHardware `json:"Hardware,omitnil,omitempty" name:"Hardware"`
}

func (r *ActivateHardwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateHardwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hardware")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateHardwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateHardwareResponseParams struct {
	// 完成激活的设备信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareInfo []*ActivateHardware `json:"HardwareInfo,omitnil,omitempty" name:"HardwareInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateHardwareResponse struct {
	*tchttp.BaseResponse
	Response *ActivateHardwareResponseParams `json:"Response"`
}

func (r *ActivateHardwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateHardwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDeviceRequestParams struct {
	// 新建设备的名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 新建设备的备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 新建设备的base64密钥字符串，非必选，如果不填写则由系统自动生成
	DataKey *string `json:"DataKey,omitnil,omitempty" name:"DataKey"`

	// 是否设置预置密钥
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// 接入环境。0：公有云网关；1：自有网关；2：公有云网关和自有网关。不填默认公有云网关。
	// 具体含义：
	// 公有云网关：即该设备只能接入公有云网关（就近接入）
	// 自有网关：即该设备只能接入已经注册上线的自有网关（就近接入或固定ip接入）
	// 公有云网关和自有网关：即该设备同时可以接入公有云网关和已经注册上线的自有网关（就近接入或固定ip接入）
	AccessScope *int64 `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// license付费方式： 
	// 0，月度授权 
	// 1，永久授权 
	// 若不传则默认为月度授权，永久授权设备需要调用OrderPerLicense接口支付授权费，否则设备无法使用
	LicensePayMode *int64 `json:"LicensePayMode,omitnil,omitempty" name:"LicensePayMode"`

	// 设备分组名称，非必选，预留参数，需要分组时传入GroupId
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备分组ID，非必选，如果不填写则默认设备无分组
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 设备无流量包处理方式，0: 按量付费，1: 截断加速
	FlowTrunc *int64 `json:"FlowTrunc,omitnil,omitempty" name:"FlowTrunc"`
}

type AddDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 新建设备的名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 新建设备的备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 新建设备的base64密钥字符串，非必选，如果不填写则由系统自动生成
	DataKey *string `json:"DataKey,omitnil,omitempty" name:"DataKey"`

	// 是否设置预置密钥
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// 接入环境。0：公有云网关；1：自有网关；2：公有云网关和自有网关。不填默认公有云网关。
	// 具体含义：
	// 公有云网关：即该设备只能接入公有云网关（就近接入）
	// 自有网关：即该设备只能接入已经注册上线的自有网关（就近接入或固定ip接入）
	// 公有云网关和自有网关：即该设备同时可以接入公有云网关和已经注册上线的自有网关（就近接入或固定ip接入）
	AccessScope *int64 `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// license付费方式： 
	// 0，月度授权 
	// 1，永久授权 
	// 若不传则默认为月度授权，永久授权设备需要调用OrderPerLicense接口支付授权费，否则设备无法使用
	LicensePayMode *int64 `json:"LicensePayMode,omitnil,omitempty" name:"LicensePayMode"`

	// 设备分组名称，非必选，预留参数，需要分组时传入GroupId
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备分组ID，非必选，如果不填写则默认设备无分组
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 设备无流量包处理方式，0: 按量付费，1: 截断加速
	FlowTrunc *int64 `json:"FlowTrunc,omitnil,omitempty" name:"FlowTrunc"`
}

func (r *AddDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceName")
	delete(f, "Remark")
	delete(f, "DataKey")
	delete(f, "Encrypted")
	delete(f, "AccessScope")
	delete(f, "LicensePayMode")
	delete(f, "GroupName")
	delete(f, "GroupId")
	delete(f, "FlowTrunc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDeviceResponseParams struct {
	// 经过加密算法加密后的base64格式密钥
	DataKey *string `json:"DataKey,omitnil,omitempty" name:"DataKey"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 签名字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AddDeviceResponseParams `json:"Response"`
}

func (r *AddDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddGroupRequestParams struct {
	// 分组的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 分组的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AddGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组的名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 分组的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *AddGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddGroupResponseParams struct {
	// 分组的唯一ID，仅做分组唯一区分
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddGroupResponseParams `json:"Response"`
}

func (r *AddGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddHardwareRequestParams struct {
	// 硬件列表
	Hardware []*Hardware `json:"Hardware,omitnil,omitempty" name:"Hardware"`
}

type AddHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 硬件列表
	Hardware []*Hardware `json:"Hardware,omitnil,omitempty" name:"Hardware"`
}

func (r *AddHardwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddHardwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hardware")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddHardwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddHardwareResponseParams struct {
	// 硬件设备
	Hardware []*Hardware `json:"Hardware,omitnil,omitempty" name:"Hardware"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddHardwareResponse struct {
	*tchttp.BaseResponse
	Response *AddHardwareResponseParams `json:"Response"`
}

func (r *AddHardwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddHardwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddL3ConnRequestParams struct {
	// 设置互通网段CIDR1，支持： 10.0.0.0 - 10.255.255.255，172.16.0.0 - 172.31.255.255，192.168.0.0 - 192.168.255.255
	Cidr1 *string `json:"Cidr1,omitnil,omitempty" name:"Cidr1"`

	// 设置互通网段CIDR2，支持： 10.0.0.0 - 10.255.255.255，172.16.0.0 - 172.31.255.255，192.168.0.0 - 192.168.255.255
	Cidr2 *string `json:"Cidr2,omitnil,omitempty" name:"Cidr2"`

	// CIDR1对应的设备ID
	DeviceId1 *string `json:"DeviceId1,omitnil,omitempty" name:"DeviceId1"`

	// CIDR2对应的设备ID
	DeviceId2 *string `json:"DeviceId2,omitnil,omitempty" name:"DeviceId2"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AddL3ConnRequest struct {
	*tchttp.BaseRequest
	
	// 设置互通网段CIDR1，支持： 10.0.0.0 - 10.255.255.255，172.16.0.0 - 172.31.255.255，192.168.0.0 - 192.168.255.255
	Cidr1 *string `json:"Cidr1,omitnil,omitempty" name:"Cidr1"`

	// 设置互通网段CIDR2，支持： 10.0.0.0 - 10.255.255.255，172.16.0.0 - 172.31.255.255，192.168.0.0 - 192.168.255.255
	Cidr2 *string `json:"Cidr2,omitnil,omitempty" name:"Cidr2"`

	// CIDR1对应的设备ID
	DeviceId1 *string `json:"DeviceId1,omitnil,omitempty" name:"DeviceId1"`

	// CIDR2对应的设备ID
	DeviceId2 *string `json:"DeviceId2,omitnil,omitempty" name:"DeviceId2"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *AddL3ConnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddL3ConnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cidr1")
	delete(f, "Cidr2")
	delete(f, "DeviceId1")
	delete(f, "DeviceId2")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddL3ConnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddL3ConnResponseParams struct {
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddL3ConnResponse struct {
	*tchttp.BaseResponse
	Response *AddL3ConnResponseParams `json:"Response"`
}

func (r *AddL3ConnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddL3ConnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Capacity struct {
	// 电信鉴权的Token。要加速的电信手机终端访问 http://qos.189.cn/qos-api/getToken?appid=TencentCloud 页面，获取返回结果中result的值
	CTCCToken *string `json:"CTCCToken,omitnil,omitempty" name:"CTCCToken"`

	// 终端所处在的省份，建议不填写由服务端自动获取，若需填写请填写带有省、市、自治区、特别行政区等后缀的省份中文全称
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`
}

type Context struct {
	// 测速数据
	NetworkData *NetworkData `json:"NetworkData,omitnil,omitempty" name:"NetworkData"`

	// 用户期望最低门限
	ExpectedLowThreshold *ExpectedThreshold `json:"ExpectedLowThreshold,omitnil,omitempty" name:"ExpectedLowThreshold"`

	// 用户期望最高门限
	ExpectedHighThreshold *ExpectedThreshold `json:"ExpectedHighThreshold,omitnil,omitempty" name:"ExpectedHighThreshold"`
}

// Predefined struct for user
type CreateEncryptedKeyRequestParams struct {

}

type CreateEncryptedKeyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateEncryptedKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEncryptedKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEncryptedKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEncryptedKeyResponseParams struct {
	// 预置密钥
	EncryptedKey *string `json:"EncryptedKey,omitnil,omitempty" name:"EncryptedKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEncryptedKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateEncryptedKeyResponseParams `json:"Response"`
}

func (r *CreateEncryptedKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEncryptedKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQosRequestParams struct {
	// 加速业务源地址信息，SrcIpv6和（SrcIpv4+SrcPublicIpv4）二选一，目前Ipv6不可用，全部填写以Ipv4参数为准。
	SrcAddressInfo *SrcAddressInfo `json:"SrcAddressInfo,omitnil,omitempty" name:"SrcAddressInfo"`

	// 加速业务目标地址信息
	DestAddressInfo *DestAddressInfo `json:"DestAddressInfo,omitnil,omitempty" name:"DestAddressInfo"`

	// 加速套餐
	// T100K：时延性保障 + 带宽保障上下行保障 100kbps
	// T200K：时延性保障 + 带宽保障上下行保障 200kbps
	// T400K：时延性保障 + 带宽保障上下行保障  400kbps
	// BD1M：带宽型保障 + 下行带宽保障1Mbps
	// BD2M：带宽型保障 + 下行带宽保障2Mbps
	// BD4M：带宽型保障 + 下行带宽保障4Mbps
	// BU1M：带宽型保障 + 上行带宽保障1Mbps
	// BU2M：带宽型保障 + 上行带宽保障2Mbps
	// BU4M：带宽型保障 + 上行带宽保障4Mbps
	QosMenu *string `json:"QosMenu,omitnil,omitempty" name:"QosMenu"`

	// 申请加速的设备信息，包括运营商，操作系统，设备唯一标识等。
	DeviceInfo *DeviceInfo `json:"DeviceInfo,omitnil,omitempty" name:"DeviceInfo"`

	// 期望加速时长（单位分钟），默认值30分钟
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 接口能力扩展，如果是电信用户，必须填充CTCC Token字段
	Capacity *Capacity `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 应用模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 针对特殊协议进行加速
	// 1. IP （默认值）
	// 2. UDP
	// 3. TCP
	Protocol *uint64 `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 加速策略关键数据
	Context *Context `json:"Context,omitnil,omitempty" name:"Context"`

	// 签名
	Extern *string `json:"Extern,omitnil,omitempty" name:"Extern"`
}

type CreateQosRequest struct {
	*tchttp.BaseRequest
	
	// 加速业务源地址信息，SrcIpv6和（SrcIpv4+SrcPublicIpv4）二选一，目前Ipv6不可用，全部填写以Ipv4参数为准。
	SrcAddressInfo *SrcAddressInfo `json:"SrcAddressInfo,omitnil,omitempty" name:"SrcAddressInfo"`

	// 加速业务目标地址信息
	DestAddressInfo *DestAddressInfo `json:"DestAddressInfo,omitnil,omitempty" name:"DestAddressInfo"`

	// 加速套餐
	// T100K：时延性保障 + 带宽保障上下行保障 100kbps
	// T200K：时延性保障 + 带宽保障上下行保障 200kbps
	// T400K：时延性保障 + 带宽保障上下行保障  400kbps
	// BD1M：带宽型保障 + 下行带宽保障1Mbps
	// BD2M：带宽型保障 + 下行带宽保障2Mbps
	// BD4M：带宽型保障 + 下行带宽保障4Mbps
	// BU1M：带宽型保障 + 上行带宽保障1Mbps
	// BU2M：带宽型保障 + 上行带宽保障2Mbps
	// BU4M：带宽型保障 + 上行带宽保障4Mbps
	QosMenu *string `json:"QosMenu,omitnil,omitempty" name:"QosMenu"`

	// 申请加速的设备信息，包括运营商，操作系统，设备唯一标识等。
	DeviceInfo *DeviceInfo `json:"DeviceInfo,omitnil,omitempty" name:"DeviceInfo"`

	// 期望加速时长（单位分钟），默认值30分钟
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 接口能力扩展，如果是电信用户，必须填充CTCC Token字段
	Capacity *Capacity `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 应用模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 针对特殊协议进行加速
	// 1. IP （默认值）
	// 2. UDP
	// 3. TCP
	Protocol *uint64 `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 加速策略关键数据
	Context *Context `json:"Context,omitnil,omitempty" name:"Context"`

	// 签名
	Extern *string `json:"Extern,omitnil,omitempty" name:"Extern"`
}

func (r *CreateQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcAddressInfo")
	delete(f, "DestAddressInfo")
	delete(f, "QosMenu")
	delete(f, "DeviceInfo")
	delete(f, "Duration")
	delete(f, "Capacity")
	delete(f, "TemplateId")
	delete(f, "Protocol")
	delete(f, "Context")
	delete(f, "Extern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQosResponseParams struct {
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 当前加速剩余时长（单位秒）
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQosResponse struct {
	*tchttp.BaseResponse
	Response *CreateQosResponseParams `json:"Response"`
}

func (r *CreateQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceRequestParams struct {
	// 删除设备的唯一ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 删除设备的唯一ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
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
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceResponseParams struct {
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
type DeleteGroupRequestParams struct {
	// 删除指定分组
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 删除指定分组
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL3ConnRequestParams struct {
	// 互通规则ID列表
	L3ConnIdList []*string `json:"L3ConnIdList,omitnil,omitempty" name:"L3ConnIdList"`
}

type DeleteL3ConnRequest struct {
	*tchttp.BaseRequest
	
	// 互通规则ID列表
	L3ConnIdList []*string `json:"L3ConnIdList,omitnil,omitempty" name:"L3ConnIdList"`
}

func (r *DeleteL3ConnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL3ConnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "L3ConnIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteL3ConnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteL3ConnResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteL3ConnResponse struct {
	*tchttp.BaseResponse
	Response *DeleteL3ConnResponseParams `json:"Response"`
}

func (r *DeleteL3ConnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteL3ConnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQosRequestParams struct {
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DeleteQosRequest struct {
	*tchttp.BaseRequest
	
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DeleteQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQosResponseParams struct {
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 本次加速会话持续时间（单位秒）
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQosResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQosResponseParams `json:"Response"`
}

func (r *DeleteQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQosRequestParams struct {
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeQosRequest struct {
	*tchttp.BaseRequest
	
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQosResponseParams struct {
	// 0：无匹配的加速中会话
	// 1：存在匹配的加速中会话
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 手机公网出口IP，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcPublicIpv4 *string `json:"SrcPublicIpv4,omitnil,omitempty" name:"SrcPublicIpv4"`

	// 业务访问目的IP，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestIpv4 []*string `json:"DestIpv4,omitnil,omitempty" name:"DestIpv4"`

	// 当前加速剩余时长（单位秒）有，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 加速套餐类型，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	QosMenu *string `json:"QosMenu,omitnil,omitempty" name:"QosMenu"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQosResponseParams `json:"Response"`
}

func (r *DescribeQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestAddressInfo struct {
	// 加速业务目标 ip 地址数组
	DestIp []*string `json:"DestIp,omitnil,omitempty" name:"DestIp"`
}

type DeviceBaseInfo struct {
	// 设备唯一ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备创建的时间，单位：ms
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 设备最后在线时间，单位：ms
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// 设备的备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 接入环境。0：公有云网关；1：自有网关；2：公有云网关和自有网关。默认公有云网关。 具体含义： 公有云网关：即该设备只能接入公有云网关（就近接入） 自有网关：即该设备只能接入已经注册上线的自有网关（就近接入或固定ip接入） 公有云网关和自有网关：即该设备同时可以接入公有云网关和已经注册上线的自有网关（就近接入或固定ip接入）
	AccessScope *int64 `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// license授权有效期 0：月度授权 1：永久授权
	LicensePayMode *int64 `json:"LicensePayMode,omitnil,omitempty" name:"LicensePayMode"`

	// 付费方 0：厂商付费 1：客户付费
	Payer *int64 `json:"Payer,omitnil,omitempty" name:"Payer"`

	// 设备分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 设备分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备无流量包处理方式，0: 按量付费，1: 截断加速
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowTrunc *int64 `json:"FlowTrunc,omitnil,omitempty" name:"FlowTrunc"`

	// 设备sn
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sn *string `json:"Sn,omitnil,omitempty" name:"Sn"`

	// 厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vendor *string `json:"Vendor,omitnil,omitempty" name:"Vendor"`
}

type DeviceDetails struct {
	// 设备基本信息
	DeviceBaseInfo *DeviceBaseInfo `json:"DeviceBaseInfo,omitnil,omitempty" name:"DeviceBaseInfo"`

	// 设备网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceNetInfo []*DeviceNetInfo `json:"DeviceNetInfo,omitnil,omitempty" name:"DeviceNetInfo"`

	// 聚合服务器地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewaySite *string `json:"GatewaySite,omitnil,omitempty" name:"GatewaySite"`

	// 业务下行速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessDownRate *float64 `json:"BusinessDownRate,omitnil,omitempty" name:"BusinessDownRate"`

	// 业务上行速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessUpRate *float64 `json:"BusinessUpRate,omitnil,omitempty" name:"BusinessUpRate"`
}

type DeviceInfo struct {
	// 运营商
	// 1：移动 
	// 2：电信
	// 3：联通
	// 4：广电
	// 99：其他
	Vendor *uint64 `json:"Vendor,omitnil,omitempty" name:"Vendor"`

	// 设备操作系统：
	// 1：Android
	// 2： IOS
	// 99：其他
	OS *uint64 `json:"OS,omitnil,omitempty" name:"OS"`

	// 设备唯一标识
	// IOS 填写 IDFV
	// Android 填写 IMEI
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 用户手机号码
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 无线信息
	// 1：4G
	// 2：5G
	// 3：WIFI
	// 99：其他
	Wireless *uint64 `json:"Wireless,omitnil,omitempty" name:"Wireless"`
}

type DeviceNetInfo struct {
	// 网络类型：
	// 0:数据
	// 1:Wi-Fi
	// 2:有线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 启用/禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEnable *bool `json:"DataEnable,omitnil,omitempty" name:"DataEnable"`

	// 上行限速
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadLimit *string `json:"UploadLimit,omitnil,omitempty" name:"UploadLimit"`

	// 下行限速
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadLimit *string `json:"DownloadLimit,omitnil,omitempty" name:"DownloadLimit"`

	// 接收实时速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRx *uint64 `json:"DataRx,omitnil,omitempty" name:"DataRx"`

	// 发送实时速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTx *uint64 `json:"DataTx,omitnil,omitempty" name:"DataTx"`

	// 运营商类型：
	// 1: 中国移动；
	// 2: 中国电信; 
	// 3: 中国联通
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vendor *int64 `json:"Vendor,omitnil,omitempty" name:"Vendor"`

	// 连接状态：
	// 0:无连接
	// 1:连接中
	// 2:已连接
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 信号强度/单位：dbm
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignalStrength *int64 `json:"SignalStrength,omitnil,omitempty" name:"SignalStrength"`

	// 数据网络类型：
	// -1 ：无效值   
	// 2：2G 
	// 3：3G 
	// 4：4G 
	// 5：5G
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rat *int64 `json:"Rat,omitnil,omitempty" name:"Rat"`

	// 网卡名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetInfoName *string `json:"NetInfoName,omitnil,omitempty" name:"NetInfoName"`

	// 下行实时速率（浮点数类型代替上一版本DataRx的整型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownRate *float64 `json:"DownRate,omitnil,omitempty" name:"DownRate"`

	// 上行实时速率（浮点数类型代替上一版本TxRate的整型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpRate *float64 `json:"UpRate,omitnil,omitempty" name:"UpRate"`
}

type DevicePayModeInfo struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 付费模式。
	// 1：预付费流量包
	// 0：按流量后付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 付费模式描述
	PayModeDesc *string `json:"PayModeDesc,omitnil,omitempty" name:"PayModeDesc"`

	// 流量包ID，仅当付费模式为流量包类型时才有。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type ExpectedThreshold struct {
	// 期望发起加速的时延阈值
	RTT *float64 `json:"RTT,omitnil,omitempty" name:"RTT"`

	// 期望发起加速的丢包率阈值
	Loss *float64 `json:"Loss,omitnil,omitempty" name:"Loss"`

	// 期望发起加速的抖动阈值
	Jitter *float64 `json:"Jitter,omitnil,omitempty" name:"Jitter"`
}

type FlowDetails struct {
	// 流量数据点
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetails []*NetDetails `json:"NetDetails,omitnil,omitempty" name:"NetDetails"`

	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 流量最大值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxValue *float64 `json:"MaxValue,omitnil,omitempty" name:"MaxValue"`

	// 流量平均值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgValue *float64 `json:"AvgValue,omitnil,omitempty" name:"AvgValue"`

	// 流量总值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalValue *float64 `json:"TotalValue,omitnil,omitempty" name:"TotalValue"`
}

type FlowPackageInfo struct {
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 流量包所属的用户AppId
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 流量包规格类型。可取值如下：
	// DEVICE_1_FLOW_20G、DEVICE_2_FLOW_50G、
	// DEVICE_3_FLOW_100G、
	// DEVICE_5_FLOW_500G，分别代表20G、50G、100G、500G档位的流量包。
	// 档位也影响流量包可绑定的设备数量上限：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 流量包状态，0：未生效，1：有效期内，2：已过期
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 购买时间，Unix时间戳格式，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 生效时间，Unix时间戳格式，单位：秒
	ActiveTime *int64 `json:"ActiveTime,omitnil,omitempty" name:"ActiveTime"`

	// 过期时间，Unix时间戳格式，单位：秒
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 流量包绑定的设备ID列表
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 流量包总容量，单位：MB
	CapacitySize *uint64 `json:"CapacitySize,omitnil,omitempty" name:"CapacitySize"`

	// 流量包余量，单位：MB
	CapacityRemain *uint64 `json:"CapacityRemain,omitnil,omitempty" name:"CapacityRemain"`

	// 自动续费标识。true代表自动续费，false代表不自动续费
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 资源包变更状态，0：未发生变配；1：变配中；2：已变配或已续费
	ModifyStatus *int64 `json:"ModifyStatus,omitnil,omitempty" name:"ModifyStatus"`

	// 流量截断标识。true代表开启流量截断，false代表不开启流量截断
	TruncFlag *bool `json:"TruncFlag,omitnil,omitempty" name:"TruncFlag"`

	// 流量包精确余量，单位：MB
	CapacityRemainPrecise *uint64 `json:"CapacityRemainPrecise,omitnil,omitempty" name:"CapacityRemainPrecise"`
}

// Predefined struct for user
type GetDevicePayModeRequestParams struct {
	// 设备ID列表
	DeviceIdList []*string `json:"DeviceIdList,omitnil,omitempty" name:"DeviceIdList"`
}

type GetDevicePayModeRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID列表
	DeviceIdList []*string `json:"DeviceIdList,omitnil,omitempty" name:"DeviceIdList"`
}

func (r *GetDevicePayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicePayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDevicePayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicePayModeResponseParams struct {
	// 结果信息
	Result []*DevicePayModeInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDevicePayModeResponse struct {
	*tchttp.BaseResponse
	Response *GetDevicePayModeResponseParams `json:"Response"`
}

func (r *GetDevicePayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicePayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceRequestParams struct {
	// 搜索指定设备的id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type GetDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 搜索指定设备的id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *GetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceResponseParams struct {
	// 设备详细信息
	DeviceDetails *DeviceDetails `json:"DeviceDetails,omitnil,omitempty" name:"DeviceDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceResponseParams `json:"Response"`
}

func (r *GetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicesRequestParams struct {
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索设备的关键字（ID或者设备名），为空时匹配所有设备
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// DeviceType
	// 不传：返回所有设备；
	// 1:自有设备；
	// 2:三方设备
	DeviceType *int64 `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type GetDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索设备的关键字（ID或者设备名），为空时匹配所有设备
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// DeviceType
	// 不传：返回所有设备；
	// 1:自有设备；
	// 2:三方设备
	DeviceType *int64 `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

func (r *GetDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Keyword")
	delete(f, "DeviceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicesResponseParams struct {
	// 设备信息列表
	DeviceInfos []*DeviceBaseInfo `json:"DeviceInfos,omitnil,omitempty" name:"DeviceInfos"`

	// 设备总记录条数
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDevicesResponse struct {
	*tchttp.BaseResponse
	Response *GetDevicesResponseParams `json:"Response"`
}

func (r *GetDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowAlarmInfoRequestParams struct {

}

type GetFlowAlarmInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetFlowAlarmInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowAlarmInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFlowAlarmInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowAlarmInfoResponseParams struct {
	// 流量包的告警阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmValue *int64 `json:"AlarmValue,omitnil,omitempty" name:"AlarmValue"`

	// 告警通知回调url
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyUrl *string `json:"NotifyUrl,omitnil,omitempty" name:"NotifyUrl"`

	// 告警通知回调key
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFlowAlarmInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetFlowAlarmInfoResponseParams `json:"Response"`
}

func (r *GetFlowAlarmInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowAlarmInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowPackagesRequestParams struct {
	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 流量包绑定的设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 流量包状态，0：未生效，1：有效期内，2：已过期
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type GetFlowPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 流量包绑定的设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 流量包状态，0：未生效，1：有效期内，2：已过期
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *GetFlowPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ResourceId")
	delete(f, "DeviceId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFlowPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowPackagesResponseParams struct {
	// 流量包列表
	PackageList []*FlowPackageInfo `json:"PackageList,omitnil,omitempty" name:"PackageList"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFlowPackagesResponse struct {
	*tchttp.BaseResponse
	Response *GetFlowPackagesResponseParams `json:"Response"`
}

func (r *GetFlowPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowStatisticByGroupRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量， 3: 上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`
}

type GetFlowStatisticByGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量， 3: 上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`
}

func (r *GetFlowStatisticByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowStatisticByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "TimeGranularity")
	delete(f, "AccessRegion")
	delete(f, "GatewayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFlowStatisticByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowStatisticByGroupResponseParams struct {
	// 流量详细信息
	NetDetails []*NetDetails `json:"NetDetails,omitnil,omitempty" name:"NetDetails"`

	// 查找时间段流量使用最大值（单位：byte）
	MaxValue *float64 `json:"MaxValue,omitnil,omitempty" name:"MaxValue"`

	// 查找时间段流量使用平均值（单位：byte）
	AvgValue *float64 `json:"AvgValue,omitnil,omitempty" name:"AvgValue"`

	// 查找时间段流量使用总量（单位：byte）
	TotalValue *float64 `json:"TotalValue,omitnil,omitempty" name:"TotalValue"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFlowStatisticByGroupResponse struct {
	*tchttp.BaseResponse
	Response *GetFlowStatisticByGroupResponseParams `json:"Response"`
}

func (r *GetFlowStatisticByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowStatisticByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowStatisticByRegionRequestParams struct {
	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量， 3: 上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 网关类型。0：公有云网关；1：自有网关。 
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`
}

type GetFlowStatisticByRegionRequest struct {
	*tchttp.BaseRequest
	
	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量， 3: 上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 网关类型。0：公有云网关；1：自有网关。 
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`
}

func (r *GetFlowStatisticByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowStatisticByRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "TimeGranularity")
	delete(f, "GatewayType")
	delete(f, "AccessRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFlowStatisticByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowStatisticByRegionResponseParams struct {
	// 流量详细信息
	NetDetails []*NetDetails `json:"NetDetails,omitnil,omitempty" name:"NetDetails"`

	// 查找时间段流量使用最大值（单位：byte）
	MaxValue *float64 `json:"MaxValue,omitnil,omitempty" name:"MaxValue"`

	// 查找时间段流量使用平均值（单位：byte）
	AvgValue *float64 `json:"AvgValue,omitnil,omitempty" name:"AvgValue"`

	// 查找时间段流量使用总量（单位：byte）
	TotalValue *float64 `json:"TotalValue,omitnil,omitempty" name:"TotalValue"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFlowStatisticByRegionResponse struct {
	*tchttp.BaseResponse
	Response *GetFlowStatisticByRegionResponseParams `json:"Response"`
}

func (r *GetFlowStatisticByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowStatisticByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowStatisticRequestParams struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量，3：上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 设备ID列表，用于查询多设备流量，该字段启用时DeviceId可传"-1"
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type GetFlowStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量，3：上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 设备ID列表，用于查询多设备流量，该字段启用时DeviceId可传"-1"
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *GetFlowStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "TimeGranularity")
	delete(f, "AccessRegion")
	delete(f, "GatewayType")
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFlowStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowStatisticResponseParams struct {
	// 流量详细信息
	NetDetails []*NetDetails `json:"NetDetails,omitnil,omitempty" name:"NetDetails"`

	// 查找时间段流量使用最大值（单位：byte）
	MaxValue *float64 `json:"MaxValue,omitnil,omitempty" name:"MaxValue"`

	// 查找时间段流量使用平均值（单位：byte）
	AvgValue *float64 `json:"AvgValue,omitnil,omitempty" name:"AvgValue"`

	// 查找时间段流量使用总量（单位：byte）
	TotalValue *float64 `json:"TotalValue,omitnil,omitempty" name:"TotalValue"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFlowStatisticResponse struct {
	*tchttp.BaseResponse
	Response *GetFlowStatisticResponseParams `json:"Response"`
}

func (r *GetFlowStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFlowStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupDetailRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备	
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备	
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索关键字
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

type GetGroupDetailRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备	
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备	
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索关键字
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

func (r *GetGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "KeyWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupDetailResponseParams struct {
	// 分组基本信息
	GroupInfo *GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

	// 分组中设备列表
	DeviceInfos []*DeviceBaseInfo `json:"DeviceInfos,omitnil,omitempty" name:"DeviceInfos"`

	// 设备总记录条数
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupDetailResponseParams `json:"Response"`
}

func (r *GetGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListRequestParams struct {
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索分组的关键字，为空时匹配所有分组
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索分组的关键字，为空时匹配所有分组
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *GetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListResponseParams struct {
	// 设备信息列表
	GroupInfos []*GroupInfo `json:"GroupInfos,omitnil,omitempty" name:"GroupInfos"`

	// 设备总记录条数
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupListResponseParams `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetHardwareListRequestParams struct {
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页面设备数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type GetHardwareListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页面设备数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *GetHardwareListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetHardwareListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetHardwareListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetHardwareListResponseParams struct {
	// 硬件信息列表
	HardwareInfos []*HardwareInfo `json:"HardwareInfos,omitnil,omitempty" name:"HardwareInfos"`

	// 硬件总数
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetHardwareListResponse struct {
	*tchttp.BaseResponse
	Response *GetHardwareListResponseParams `json:"Response"`
}

func (r *GetHardwareListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetHardwareListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetL3ConnListRequestParams struct {
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索分组的DeviceId，为空时匹配所有分组
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type GetL3ConnListRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 搜索分组的DeviceId，为空时匹配所有分组
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *GetL3ConnListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetL3ConnListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetL3ConnListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetL3ConnListResponseParams struct {
	// 互通规则列表
	L3ConnList []*L3ConnInfo `json:"L3ConnList,omitnil,omitempty" name:"L3ConnList"`

	// 设备总记录条数
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetL3ConnListResponse struct {
	*tchttp.BaseResponse
	Response *GetL3ConnListResponseParams `json:"Response"`
}

func (r *GetL3ConnListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetL3ConnListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMultiFlowStatisticRequestParams struct {
	// 设备id列表，单次最多请求10个设备
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`

	// 1659514436
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 1659515000
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计流量类型（1：上行流量，2：下行流量， 3: 上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 统计时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`
}

type GetMultiFlowStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 设备id列表，单次最多请求10个设备
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`

	// 1659514436
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 1659515000
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计流量类型（1：上行流量，2：下行流量， 3: 上下行总和）
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 统计时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`
}

func (r *GetMultiFlowStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMultiFlowStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "TimeGranularity")
	delete(f, "AccessRegion")
	delete(f, "GatewayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMultiFlowStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMultiFlowStatisticResponseParams struct {
	// 批量设备流量信息
	FlowDetails []*FlowDetails `json:"FlowDetails,omitnil,omitempty" name:"FlowDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMultiFlowStatisticResponse struct {
	*tchttp.BaseResponse
	Response *GetMultiFlowStatisticResponseParams `json:"Response"`
}

func (r *GetMultiFlowStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMultiFlowStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetNetMonitorRequestParams struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计指标（上行速率："TxRate":bit/s，下行速率："RxRate":bit/s，丢包："Loss":%，时延："RTT":ms）
	Metrics *string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`
}

type GetNetMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 统计指标（上行速率："TxRate":bit/s，下行速率："RxRate":bit/s，丢包："Loss":%，时延："RTT":ms）
	Metrics *string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`
}

func (r *GetNetMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetNetMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Metrics")
	delete(f, "GatewayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetNetMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetNetMonitorResponseParams struct {
	// 监控数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorData []*MonitorData `json:"MonitorData,omitnil,omitempty" name:"MonitorData"`

	// 接入区域。取值范围：['MC','AP','EU','AM']
	// MC=中国大陆
	// AP=亚太
	// EU=欧洲
	// AM=美洲
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetNetMonitorResponse struct {
	*tchttp.BaseResponse
	Response *GetNetMonitorResponseParams `json:"Response"`
}

func (r *GetNetMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetNetMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyRequestParams struct {

}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyResponseParams struct {
	// 非对称公钥
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetPublicKeyResponseParams `json:"Response"`
}

func (r *GetPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStatisticDataRequestParams struct {
	// 设备ID。若不指定设备，可传"-1"
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 统计开始时间，单位：s
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 统计结束时间，单位：s
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合粒度：
	// 1:按小时统计
	// 2:按天统计
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 设备ID列表，最多10个设备，下载多个设备流量和时使用，此时DeviceId可传"-1"
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 设备分组ID，若不指定分组则不传，按分组下载数据时使用
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type GetStatisticDataRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID。若不指定设备，可传"-1"
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 统计开始时间，单位：s
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 统计结束时间，单位：s
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合粒度：
	// 1:按小时统计
	// 2:按天统计
	TimeGranularity *int64 `json:"TimeGranularity,omitnil,omitempty" name:"TimeGranularity"`

	// 接入区域。取值范围：['MC','AP','EU','AM'] MC=中国大陆 AP=亚太 EU=欧洲 AM=美洲。不填代表全量区域。
	AccessRegion *string `json:"AccessRegion,omitnil,omitempty" name:"AccessRegion"`

	// 网关类型。0：公有云网关；1：自有网关。不传默认为0。
	GatewayType *int64 `json:"GatewayType,omitnil,omitempty" name:"GatewayType"`

	// 设备ID列表，最多10个设备，下载多个设备流量和时使用，此时DeviceId可传"-1"
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 设备分组ID，若不指定分组则不传，按分组下载数据时使用
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *GetStatisticDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStatisticDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "TimeGranularity")
	delete(f, "AccessRegion")
	delete(f, "GatewayType")
	delete(f, "DeviceList")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetStatisticDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStatisticDataResponseParams struct {
	// 文件地址url
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetStatisticDataResponse struct {
	*tchttp.BaseResponse
	Response *GetStatisticDataResponseParams `json:"Response"`
}

func (r *GetStatisticDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStatisticDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVendorHardwareRequestParams struct {
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 激活状态，
	// 空：全部；
	// 1:待激活；
	// 2:已激活；
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type GetVendorHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 激活状态，
	// 空：全部；
	// 1:待激活；
	// 2:已激活；
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *GetVendorHardwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVendorHardwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVendorHardwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVendorHardwareResponseParams struct {
	// 硬件信息列表
	VendorHardware []*VendorHardware `json:"VendorHardware,omitnil,omitempty" name:"VendorHardware"`

	// 设备总数
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetVendorHardwareResponse struct {
	*tchttp.BaseResponse
	Response *GetVendorHardwareResponseParams `json:"Response"`
}

func (r *GetVendorHardwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVendorHardwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupAddDeviceRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 待添加的设备列表
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type GroupAddDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 待添加的设备列表
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *GroupAddDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupAddDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GroupAddDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupAddDeviceResponseParams struct {
	// 分组中的设备数量
	DeviceNum *int64 `json:"DeviceNum,omitnil,omitempty" name:"DeviceNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GroupAddDeviceResponse struct {
	*tchttp.BaseResponse
	Response *GroupAddDeviceResponseParams `json:"Response"`
}

func (r *GroupAddDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupAddDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupDeleteDeviceRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 待删除的设备列表
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type GroupDeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 待删除的设备列表
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

func (r *GroupDeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupDeleteDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "DeviceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GroupDeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupDeleteDeviceResponseParams struct {
	// 分组中的设备数量
	DeviceNum *int64 `json:"DeviceNum,omitnil,omitempty" name:"DeviceNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GroupDeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *GroupDeleteDeviceResponseParams `json:"Response"`
}

func (r *GroupDeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupDeleteDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 分组创建的时间，单位：ms	
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分组更新的时间，单位：ms	
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 分组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 分组中的设备数量
	DeviceNum *int64 `json:"DeviceNum,omitnil,omitempty" name:"DeviceNum"`
}

type Hardware struct {
	// 硬件序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// license计费模式：
	// 1，租户付费
	// 2，厂商月付费
	// 3，厂商永久授权
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseChargingMode *int64 `json:"LicenseChargingMode,omitnil,omitempty" name:"LicenseChargingMode"`

	// 设备描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 硬件ID，入参无需传递
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareId *string `json:"HardwareId,omitnil,omitempty" name:"HardwareId"`
}

type HardwareInfo struct {
	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitnil,omitempty" name:"ActiveTime"`

	// 最后在线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOnlineTime *string `json:"LastOnlineTime,omitnil,omitempty" name:"LastOnlineTime"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 厂商备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	VendorDescription *string `json:"VendorDescription,omitnil,omitempty" name:"VendorDescription"`

	// license计费模式： 1，租户月付费 2，厂商月付费 3，license永久授权
	// 注：后续将废弃此参数，新接入请使用LicensePayMode和Payer
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseChargingMode *int64 `json:"LicenseChargingMode,omitnil,omitempty" name:"LicenseChargingMode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 硬件序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// license授权有效期 
	// 0：月度授权 
	// 1：永久授权
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicensePayMode *int64 `json:"LicensePayMode,omitnil,omitempty" name:"LicensePayMode"`

	// 付费方 
	// 0：客户付费 
	// 1：厂商付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	Payer *int64 `json:"Payer,omitnil,omitempty" name:"Payer"`

	// 设备分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 设备分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 设备无流量包处理方式，0: 按量付费，1: 截断加速	
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowTrunc *int64 `json:"FlowTrunc,omitnil,omitempty" name:"FlowTrunc"`
}

type L3ConnInfo struct {
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 互通设备ID
	DeviceId1 *string `json:"DeviceId1,omitnil,omitempty" name:"DeviceId1"`

	// 互通规则CIDR
	Cidr1 *string `json:"Cidr1,omitnil,omitempty" name:"Cidr1"`

	// 互通设备ID
	DeviceId2 *string `json:"DeviceId2,omitnil,omitempty" name:"DeviceId2"`

	// 互通规则CIDR
	Cidr2 *string `json:"Cidr2,omitnil,omitempty" name:"Cidr2"`

	// 互通规则启用状态
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 互通规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type ModifyPackageRenewFlagRequestParams struct {
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自动续费标识。true代表自动续费，false代表不自动续费
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyPackageRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自动续费标识。true代表自动续费，false代表不自动续费
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *ModifyPackageRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPackageRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPackageRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPackageRenewFlagResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPackageRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPackageRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyPackageRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPackageRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorData struct {
	// 时间点：s
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 业务指标（bps/ms/%）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessMetrics *float64 `json:"BusinessMetrics,omitnil,omitempty" name:"BusinessMetrics"`

	// 网卡状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotNetInfo []*SlotNetInfo `json:"SlotNetInfo,omitnil,omitempty" name:"SlotNetInfo"`
}

type NetDetails struct {
	// 流量值（byte）
	Current *float64 `json:"Current,omitnil,omitempty" name:"Current"`

	// 时间点，单位：s
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`
}

type NetworkData struct {
	// 时延数组，最大长度30
	RTT []*float64 `json:"RTT,omitnil,omitempty" name:"RTT"`

	// 丢包率
	Loss *float64 `json:"Loss,omitnil,omitempty" name:"Loss"`

	// 抖动
	Jitter *float64 `json:"Jitter,omitnil,omitempty" name:"Jitter"`

	// 10位秒级时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

// Predefined struct for user
type OrderFlowPackageRequestParams struct {
	// 流量包规格类型。可取值如下：
	// DEVICE_1_FLOW_20G、DEVICE_2_FLOW_50G、
	// DEVICE_3_FLOW_100G、
	// DEVICE_5_FLOW_500G，分别代表20G、50G、100G、500G档位的流量包。
	// 档位也影响流量包可绑定的设备数量上限：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 流量包绑定的设备ID列表。捆绑设备个数上限取决于包的规格档位：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 是否自动续费，该选项和流量截断冲突，只能开启一个
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 区域标识，0：国内，1：国外
	PackageRegion *int64 `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 是否开启流量截断功能，该选项和自动续费冲突
	FlowTruncFlag *bool `json:"FlowTruncFlag,omitnil,omitempty" name:"FlowTruncFlag"`

	// 是否自动选择代金券，默认false。
	// 有多张券时的选择策略：按照可支付订单全部金额的券，先到期的券，可抵扣金额最大的券，余额最小的券，现金券 这个优先级进行扣券，且最多只抵扣一张券。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 指定代金券ID。自动选择代金券时此参数无效。目前只允许传入一张代金券。
	// 注：若指定的代金券不符合订单抵扣条件，则正常支付，不扣券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type OrderFlowPackageRequest struct {
	*tchttp.BaseRequest
	
	// 流量包规格类型。可取值如下：
	// DEVICE_1_FLOW_20G、DEVICE_2_FLOW_50G、
	// DEVICE_3_FLOW_100G、
	// DEVICE_5_FLOW_500G，分别代表20G、50G、100G、500G档位的流量包。
	// 档位也影响流量包可绑定的设备数量上限：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 流量包绑定的设备ID列表。捆绑设备个数上限取决于包的规格档位：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	DeviceList []*string `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`

	// 是否自动续费，该选项和流量截断冲突，只能开启一个
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 区域标识，0：国内，1：国外
	PackageRegion *int64 `json:"PackageRegion,omitnil,omitempty" name:"PackageRegion"`

	// 是否开启流量截断功能，该选项和自动续费冲突
	FlowTruncFlag *bool `json:"FlowTruncFlag,omitnil,omitempty" name:"FlowTruncFlag"`

	// 是否自动选择代金券，默认false。
	// 有多张券时的选择策略：按照可支付订单全部金额的券，先到期的券，可抵扣金额最大的券，余额最小的券，现金券 这个优先级进行扣券，且最多只抵扣一张券。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 指定代金券ID。自动选择代金券时此参数无效。目前只允许传入一张代金券。
	// 注：若指定的代金券不符合订单抵扣条件，则正常支付，不扣券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

func (r *OrderFlowPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OrderFlowPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageType")
	delete(f, "DeviceList")
	delete(f, "AutoRenewFlag")
	delete(f, "PackageRegion")
	delete(f, "FlowTruncFlag")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OrderFlowPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OrderFlowPackageResponseParams struct {
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OrderFlowPackageResponse struct {
	*tchttp.BaseResponse
	Response *OrderFlowPackageResponseParams `json:"Response"`
}

func (r *OrderFlowPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OrderFlowPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OrderPerLicenseRequestParams struct {
	// 购买永久授权License的设备ID，如果是厂商未激活设备采用HardwareId
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备类型，0: SDK，1: CPE，作为用户创建或激活设备时传0，作为厂商创建待激活设备时传1
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 购买失败后是否回滚（删除）设备，默认true，如果设备绑定了生效中的流量包则不能回滚。
	RollBack *bool `json:"RollBack,omitnil,omitempty" name:"RollBack"`

	// 是否自动选择代金券，默认false。
	// 有多张券时的选择策略：按照可支付订单全部金额的券，先到期的券，可抵扣金额最大的券，余额最小的券，现金券 这个优先级进行扣券，且最多只抵扣一张券。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 指定代金券ID。自动选择代金券时此参数无效。目前只允许传入一张代金券。
	// 注：若指定的代金券不符合订单抵扣条件，则正常支付，不扣券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type OrderPerLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 购买永久授权License的设备ID，如果是厂商未激活设备采用HardwareId
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备类型，0: SDK，1: CPE，作为用户创建或激活设备时传0，作为厂商创建待激活设备时传1
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 购买失败后是否回滚（删除）设备，默认true，如果设备绑定了生效中的流量包则不能回滚。
	RollBack *bool `json:"RollBack,omitnil,omitempty" name:"RollBack"`

	// 是否自动选择代金券，默认false。
	// 有多张券时的选择策略：按照可支付订单全部金额的券，先到期的券，可抵扣金额最大的券，余额最小的券，现金券 这个优先级进行扣券，且最多只抵扣一张券。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 指定代金券ID。自动选择代金券时此参数无效。目前只允许传入一张代金券。
	// 注：若指定的代金券不符合订单抵扣条件，则正常支付，不扣券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

func (r *OrderPerLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OrderPerLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Type")
	delete(f, "RollBack")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OrderPerLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OrderPerLicenseResponseParams struct {
	// 一次性授权License的资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OrderPerLicenseResponse struct {
	*tchttp.BaseResponse
	Response *OrderPerLicenseResponseParams `json:"Response"`
}

func (r *OrderPerLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OrderPerLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNotifyUrlRequestParams struct {
	// 告警通知回调url
	NotifyUrl *string `json:"NotifyUrl,omitnil,omitempty" name:"NotifyUrl"`

	// 告警通知回调key
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 流量包的告警阈值
	AlarmValue *int64 `json:"AlarmValue,omitnil,omitempty" name:"AlarmValue"`
}

type SetNotifyUrlRequest struct {
	*tchttp.BaseRequest
	
	// 告警通知回调url
	NotifyUrl *string `json:"NotifyUrl,omitnil,omitempty" name:"NotifyUrl"`

	// 告警通知回调key
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 流量包的告警阈值
	AlarmValue *int64 `json:"AlarmValue,omitnil,omitempty" name:"AlarmValue"`
}

func (r *SetNotifyUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNotifyUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotifyUrl")
	delete(f, "CallbackKey")
	delete(f, "AlarmValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNotifyUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNotifyUrlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetNotifyUrlResponse struct {
	*tchttp.BaseResponse
	Response *SetNotifyUrlResponseParams `json:"Response"`
}

func (r *SetNotifyUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNotifyUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlotNetInfo struct {
	// 网卡名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetInfoName *string `json:"NetInfoName,omitnil,omitempty" name:"NetInfoName"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIP *string `json:"PublicIP,omitnil,omitempty" name:"PublicIP"`

	// 指标数据（bps/ms/%）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Current *float64 `json:"Current,omitnil,omitempty" name:"Current"`
}

type SrcAddressInfo struct {
	// 用户私网 ipv4 地址
	SrcIpv4 *string `json:"SrcIpv4,omitnil,omitempty" name:"SrcIpv4"`

	// 用户公网 ipv4 地址
	SrcPublicIpv4 *string `json:"SrcPublicIpv4,omitnil,omitempty" name:"SrcPublicIpv4"`

	// 用户 ipv6 地址
	SrcIpv6 *string `json:"SrcIpv6,omitnil,omitempty" name:"SrcIpv6"`
}

// Predefined struct for user
type UpdateDeviceRequestParams struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 更新设备网络信息
	UpdateNetInfo []*UpdateNetInfo `json:"UpdateNetInfo,omitnil,omitempty" name:"UpdateNetInfo"`

	// 设备无流量包处理方式，0: 按量付费，1: 截断加速
	FlowTrunc *int64 `json:"FlowTrunc,omitnil,omitempty" name:"FlowTrunc"`
}

type UpdateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 更新设备网络信息
	UpdateNetInfo []*UpdateNetInfo `json:"UpdateNetInfo,omitnil,omitempty" name:"UpdateNetInfo"`

	// 设备无流量包处理方式，0: 按量付费，1: 截断加速
	FlowTrunc *int64 `json:"FlowTrunc,omitnil,omitempty" name:"FlowTrunc"`
}

func (r *UpdateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "DeviceName")
	delete(f, "Remark")
	delete(f, "UpdateNetInfo")
	delete(f, "FlowTrunc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceResponseParams `json:"Response"`
}

func (r *UpdateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupRequestParams struct {
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGroupResponseParams `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateHardwareRequestParams struct {
	// 硬件ID
	HardwareId *string `json:"HardwareId,omitnil,omitempty" name:"HardwareId"`

	// 硬件序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 设备备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 硬件ID
	HardwareId *string `json:"HardwareId,omitnil,omitempty" name:"HardwareId"`

	// 硬件序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 设备备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateHardwareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateHardwareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HardwareId")
	delete(f, "SN")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateHardwareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateHardwareResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateHardwareResponse struct {
	*tchttp.BaseResponse
	Response *UpdateHardwareResponseParams `json:"Response"`
}

func (r *UpdateHardwareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateHardwareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateL3CidrRequestParams struct {
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 互通规则CIDR
	Cidr1 *string `json:"Cidr1,omitnil,omitempty" name:"Cidr1"`

	// 互通设备ID
	DeviceId1 *string `json:"DeviceId1,omitnil,omitempty" name:"DeviceId1"`

	// 互通设备ID
	DeviceId2 *string `json:"DeviceId2,omitnil,omitempty" name:"DeviceId2"`

	// 互通规则CIDR
	Cidr2 *string `json:"Cidr2,omitnil,omitempty" name:"Cidr2"`
}

type UpdateL3CidrRequest struct {
	*tchttp.BaseRequest
	
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 互通规则CIDR
	Cidr1 *string `json:"Cidr1,omitnil,omitempty" name:"Cidr1"`

	// 互通设备ID
	DeviceId1 *string `json:"DeviceId1,omitnil,omitempty" name:"DeviceId1"`

	// 互通设备ID
	DeviceId2 *string `json:"DeviceId2,omitnil,omitempty" name:"DeviceId2"`

	// 互通规则CIDR
	Cidr2 *string `json:"Cidr2,omitnil,omitempty" name:"Cidr2"`
}

func (r *UpdateL3CidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateL3CidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "L3ConnId")
	delete(f, "Cidr1")
	delete(f, "DeviceId1")
	delete(f, "DeviceId2")
	delete(f, "Cidr2")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateL3CidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateL3CidrResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateL3CidrResponse struct {
	*tchttp.BaseResponse
	Response *UpdateL3CidrResponseParams `json:"Response"`
}

func (r *UpdateL3CidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateL3CidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateL3ConnRequestParams struct {
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 互通规则备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateL3ConnRequest struct {
	*tchttp.BaseRequest
	
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 互通规则备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateL3ConnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateL3ConnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "L3ConnId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateL3ConnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateL3ConnResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateL3ConnResponse struct {
	*tchttp.BaseResponse
	Response *UpdateL3ConnResponseParams `json:"Response"`
}

func (r *UpdateL3ConnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateL3ConnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateL3SwitchRequestParams struct {
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 互通规则开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type UpdateL3SwitchRequest struct {
	*tchttp.BaseRequest
	
	// 互通规则ID
	L3ConnId *string `json:"L3ConnId,omitnil,omitempty" name:"L3ConnId"`

	// 互通规则开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *UpdateL3SwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateL3SwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "L3ConnId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateL3SwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateL3SwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateL3SwitchResponse struct {
	*tchttp.BaseResponse
	Response *UpdateL3SwitchResponseParams `json:"Response"`
}

func (r *UpdateL3SwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateL3SwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNetInfo struct {
	// 网络类型：
	// 0:数据
	// 1:Wi-Fi
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 启用/禁用
	DataEnable *bool `json:"DataEnable,omitnil,omitempty" name:"DataEnable"`

	// 上行限速：bit
	UploadLimit *uint64 `json:"UploadLimit,omitnil,omitempty" name:"UploadLimit"`

	// 下行限速：bit
	DownloadLimit *uint64 `json:"DownloadLimit,omitnil,omitempty" name:"DownloadLimit"`

	// 网卡名
	NetInfoName *string `json:"NetInfoName,omitnil,omitempty" name:"NetInfoName"`
}

type VendorHardware struct {
	// 硬件id
	HardwareId *string `json:"HardwareId,omitnil,omitempty" name:"HardwareId"`

	// 硬件序列号
	SN *string `json:"SN,omitnil,omitempty" name:"SN"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 激活状态， 空：全部； 1:待激活； 2:已激活
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitnil,omitempty" name:"ActiveTime"`

	// 厂商备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// license计费模式： 1，租户月付费 2，厂商月付费 3，license永久授权
	// 注：设备为租户付费且未激活（未选择月付还是永久付费）时，此参数返回1，仅代表租户付费。后续将废弃此参数，新接入请使用LicensePayMode和Payer
	LicenseChargingMode *int64 `json:"LicenseChargingMode,omitnil,omitempty" name:"LicenseChargingMode"`

	// 最后在线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOnlineTime *string `json:"LastOnlineTime,omitnil,omitempty" name:"LastOnlineTime"`

	// license授权有效期
	// 0：月度授权
	// 1：永久授权
	// -1：未知
	LicensePayMode *int64 `json:"LicensePayMode,omitnil,omitempty" name:"LicensePayMode"`

	// 付费方
	// 0：客户付费
	// 1：厂商付费
	Payer *int64 `json:"Payer,omitnil,omitempty" name:"Payer"`
}