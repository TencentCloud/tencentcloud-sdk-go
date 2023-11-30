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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vendor *string `json:"Vendor,omitnil" name:"Vendor"`

	// 设备SN序列号
	SN *string `json:"SN,omitnil" name:"SN"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备密钥
	DataKey *string `json:"DataKey,omitnil" name:"DataKey"`
}

// Predefined struct for user
type ActivateHardwareRequestParams struct {
	// 待激活的设备列表
	Hardware []*ActivateHardware `json:"Hardware,omitnil" name:"Hardware"`
}

type ActivateHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 待激活的设备列表
	Hardware []*ActivateHardware `json:"Hardware,omitnil" name:"Hardware"`
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
	HardwareInfo []*ActivateHardware `json:"HardwareInfo,omitnil" name:"HardwareInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 新建设备的备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 新建设备的base64密钥字符串，非必选，如果不填写则由系统自动生成
	DataKey *string `json:"DataKey,omitnil" name:"DataKey"`

	// 是否设置预置密钥
	Encrypted *bool `json:"Encrypted,omitnil" name:"Encrypted"`
}

type AddDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 新建设备的名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 新建设备的备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 新建设备的base64密钥字符串，非必选，如果不填写则由系统自动生成
	DataKey *string `json:"DataKey,omitnil" name:"DataKey"`

	// 是否设置预置密钥
	Encrypted *bool `json:"Encrypted,omitnil" name:"Encrypted"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDeviceResponseParams struct {
	// 经过加密算法加密后的base64格式密钥
	DataKey *string `json:"DataKey,omitnil" name:"DataKey"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 签名字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type AddHardwareRequestParams struct {
	// 硬件列表
	Hardware []*Hardware `json:"Hardware,omitnil" name:"Hardware"`
}

type AddHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 硬件列表
	Hardware []*Hardware `json:"Hardware,omitnil" name:"Hardware"`
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
	Hardware []*Hardware `json:"Hardware,omitnil" name:"Hardware"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type Capacity struct {
	// 电信鉴权的Token。要加速的电信手机终端访问 http://qos.189.cn/qos-api/getToken?appid=TencentCloud 页面，获取返回结果中result的值
	CTCCToken *string `json:"CTCCToken,omitnil" name:"CTCCToken"`

	// 终端所处在的省份，建议不填写由服务端自动获取，若需填写请填写带有省、市、自治区、特别行政区等后缀的省份中文全称
	Province *string `json:"Province,omitnil" name:"Province"`
}

type Context struct {
	// 测速数据
	NetworkData *NetworkData `json:"NetworkData,omitnil" name:"NetworkData"`

	// 用户期望最低门限
	ExpectedLowThreshold *ExpectedThreshold `json:"ExpectedLowThreshold,omitnil" name:"ExpectedLowThreshold"`

	// 用户期望最高门限
	ExpectedHighThreshold *ExpectedThreshold `json:"ExpectedHighThreshold,omitnil" name:"ExpectedHighThreshold"`
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
	EncryptedKey *string `json:"EncryptedKey,omitnil" name:"EncryptedKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SrcAddressInfo *SrcAddressInfo `json:"SrcAddressInfo,omitnil" name:"SrcAddressInfo"`

	// 加速业务目标地址信息
	DestAddressInfo *DestAddressInfo `json:"DestAddressInfo,omitnil" name:"DestAddressInfo"`

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
	QosMenu *string `json:"QosMenu,omitnil" name:"QosMenu"`

	// 申请加速的设备信息，包括运营商，操作系统，设备唯一标识等。
	DeviceInfo *DeviceInfo `json:"DeviceInfo,omitnil" name:"DeviceInfo"`

	// 期望加速时长（单位分钟），默认值30分钟
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 接口能力扩展，如果是电信用户，必须填充CTCC Token字段
	Capacity *Capacity `json:"Capacity,omitnil" name:"Capacity"`

	// 应用模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 针对特殊协议进行加速
	// 1. IP （默认值）
	// 2. UDP
	// 3. TCP
	Protocol *uint64 `json:"Protocol,omitnil" name:"Protocol"`

	// 加速策略关键数据
	Context *Context `json:"Context,omitnil" name:"Context"`

	// 签名
	Extern *string `json:"Extern,omitnil" name:"Extern"`
}

type CreateQosRequest struct {
	*tchttp.BaseRequest
	
	// 加速业务源地址信息，SrcIpv6和（SrcIpv4+SrcPublicIpv4）二选一，目前Ipv6不可用，全部填写以Ipv4参数为准。
	SrcAddressInfo *SrcAddressInfo `json:"SrcAddressInfo,omitnil" name:"SrcAddressInfo"`

	// 加速业务目标地址信息
	DestAddressInfo *DestAddressInfo `json:"DestAddressInfo,omitnil" name:"DestAddressInfo"`

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
	QosMenu *string `json:"QosMenu,omitnil" name:"QosMenu"`

	// 申请加速的设备信息，包括运营商，操作系统，设备唯一标识等。
	DeviceInfo *DeviceInfo `json:"DeviceInfo,omitnil" name:"DeviceInfo"`

	// 期望加速时长（单位分钟），默认值30分钟
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 接口能力扩展，如果是电信用户，必须填充CTCC Token字段
	Capacity *Capacity `json:"Capacity,omitnil" name:"Capacity"`

	// 应用模板ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 针对特殊协议进行加速
	// 1. IP （默认值）
	// 2. UDP
	// 3. TCP
	Protocol *uint64 `json:"Protocol,omitnil" name:"Protocol"`

	// 加速策略关键数据
	Context *Context `json:"Context,omitnil" name:"Context"`

	// 签名
	Extern *string `json:"Extern,omitnil" name:"Extern"`
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
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 当前加速剩余时长（单位秒）
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 删除设备的唯一ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeleteQosRequestParams struct {
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

type DeleteQosRequest struct {
	*tchttp.BaseRequest
	
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
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
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 本次加速会话持续时间（单位秒）
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

type DescribeQosRequest struct {
	*tchttp.BaseRequest
	
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
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
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 手机公网出口IP，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcPublicIpv4 *string `json:"SrcPublicIpv4,omitnil" name:"SrcPublicIpv4"`

	// 业务访问目的IP，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestIpv4 []*string `json:"DestIpv4,omitnil" name:"DestIpv4"`

	// 当前加速剩余时长（单位秒）有，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil" name:"Duration"`

	// 加速套餐类型，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	QosMenu *string `json:"QosMenu,omitnil" name:"QosMenu"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DestIp []*string `json:"DestIp,omitnil" name:"DestIp"`
}

type DeviceBaseInfo struct {
	// 设备唯一ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备创建的时间，单位：ms
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 设备最后在线时间，单位：ms
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 设备的备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type DeviceDetails struct {
	// 设备基本信息
	DeviceBaseInfo *DeviceBaseInfo `json:"DeviceBaseInfo,omitnil" name:"DeviceBaseInfo"`

	// 设备网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceNetInfo []*DeviceNetInfo `json:"DeviceNetInfo,omitnil" name:"DeviceNetInfo"`

	// 聚合服务器地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewaySite *string `json:"GatewaySite,omitnil" name:"GatewaySite"`

	// 业务下行速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessDownRate *float64 `json:"BusinessDownRate,omitnil" name:"BusinessDownRate"`

	// 业务上行速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessUpRate *float64 `json:"BusinessUpRate,omitnil" name:"BusinessUpRate"`
}

type DeviceInfo struct {
	// 运营商
	// 1：移动 
	// 2：电信
	// 3：联通
	// 4：广电
	// 99：其他
	Vendor *uint64 `json:"Vendor,omitnil" name:"Vendor"`

	// 设备操作系统：
	// 1：Android
	// 2： IOS
	// 99：其他
	OS *uint64 `json:"OS,omitnil" name:"OS"`

	// 设备唯一标识
	// IOS 填写 IDFV
	// Android 填写 IMEI
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 用户手机号码
	PhoneNum *string `json:"PhoneNum,omitnil" name:"PhoneNum"`

	// 无线信息
	// 1：4G
	// 2：5G
	// 3：WIFI
	// 99：其他
	Wireless *uint64 `json:"Wireless,omitnil" name:"Wireless"`
}

type DeviceNetInfo struct {
	// 网络类型：
	// 0:数据
	// 1:Wi-Fi
	// 2:有线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 启用/禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEnable *bool `json:"DataEnable,omitnil" name:"DataEnable"`

	// 上行限速
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadLimit *string `json:"UploadLimit,omitnil" name:"UploadLimit"`

	// 下行限速
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadLimit *string `json:"DownloadLimit,omitnil" name:"DownloadLimit"`

	// 接收实时速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRx *uint64 `json:"DataRx,omitnil" name:"DataRx"`

	// 发送实时速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTx *uint64 `json:"DataTx,omitnil" name:"DataTx"`

	// 运营商类型：
	// 1: 中国移动；
	// 2: 中国电信; 
	// 3: 中国联通
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vendor *int64 `json:"Vendor,omitnil" name:"Vendor"`

	// 连接状态：
	// 0:无连接
	// 1:连接中
	// 2:已连接
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil" name:"State"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 信号强度/单位：dbm
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignalStrength *int64 `json:"SignalStrength,omitnil" name:"SignalStrength"`

	// 数据网络类型：
	// -1 ：无效值   
	// 2：2G 
	// 3：3G 
	// 4：4G 
	// 5：5G
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rat *int64 `json:"Rat,omitnil" name:"Rat"`

	// 网卡名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetInfoName *string `json:"NetInfoName,omitnil" name:"NetInfoName"`

	// 下行实时速率（浮点数类型代替上一版本DataRx的整型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownRate *float64 `json:"DownRate,omitnil" name:"DownRate"`

	// 上行实时速率（浮点数类型代替上一版本TxRate的整型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpRate *float64 `json:"UpRate,omitnil" name:"UpRate"`
}

type DevicePayModeInfo struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 付费模式。
	// 1：预付费流量包
	// 0：按流量后付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 付费模式描述
	PayModeDesc *string `json:"PayModeDesc,omitnil" name:"PayModeDesc"`

	// 流量包ID，仅当付费模式为流量包类型时才有。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type ExpectedThreshold struct {
	// 期望发起加速的时延阈值
	RTT *float64 `json:"RTT,omitnil" name:"RTT"`

	// 期望发起加速的丢包率阈值
	Loss *float64 `json:"Loss,omitnil" name:"Loss"`

	// 期望发起加速的抖动阈值
	Jitter *float64 `json:"Jitter,omitnil" name:"Jitter"`
}

type FlowDetails struct {
	// 流量数据点
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetails []*NetDetails `json:"NetDetails,omitnil" name:"NetDetails"`

	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 流量最大值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxValue *float64 `json:"MaxValue,omitnil" name:"MaxValue"`

	// 流量平均值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgValue *float64 `json:"AvgValue,omitnil" name:"AvgValue"`

	// 流量总值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalValue *float64 `json:"TotalValue,omitnil" name:"TotalValue"`
}

type FlowPackageInfo struct {
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流量包所属的用户AppId
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 流量包规格类型。可取值如下：
	// DEVICE_1_FLOW_20G、DEVICE_2_FLOW_50G、
	// DEVICE_3_FLOW_100G、
	// DEVICE_5_FLOW_500G，分别代表20G、50G、100G、500G档位的流量包。
	// 档位也影响流量包可绑定的设备数量上限：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// 流量包状态，0：未生效，1：有效期内，2：已过期
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 生效时间，Unix时间戳格式，单位：秒
	ActiveTime *int64 `json:"ActiveTime,omitnil" name:"ActiveTime"`

	// 过期时间，Unix时间戳格式，单位：秒
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 流量包绑定的设备ID列表
	DeviceList []*string `json:"DeviceList,omitnil" name:"DeviceList"`

	// 流量包总容量，单位：MB
	CapacitySize *uint64 `json:"CapacitySize,omitnil" name:"CapacitySize"`

	// 流量包余量，单位：MB
	CapacityRemain *uint64 `json:"CapacityRemain,omitnil" name:"CapacityRemain"`

	// 自动续费标识。true代表自动续费，false代表不自动续费
	RenewFlag *bool `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

// Predefined struct for user
type GetDevicePayModeRequestParams struct {
	// 设备ID列表
	DeviceIdList []*string `json:"DeviceIdList,omitnil" name:"DeviceIdList"`
}

type GetDevicePayModeRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID列表
	DeviceIdList []*string `json:"DeviceIdList,omitnil" name:"DeviceIdList"`
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
	Result []*DevicePayModeInfo `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
}

type GetDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 搜索指定设备的id
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`
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
	DeviceDetails *DeviceDetails `json:"DeviceDetails,omitnil" name:"DeviceDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 搜索设备的关键字（ID或者设备名），为空时匹配所有设备
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// DeviceType
	// 不传：返回所有设备；
	// 1:自有设备；
	// 2:三方设备
	DeviceType *int64 `json:"DeviceType,omitnil" name:"DeviceType"`
}

type GetDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 搜索设备的关键字（ID或者设备名），为空时匹配所有设备
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// DeviceType
	// 不传：返回所有设备；
	// 1:自有设备；
	// 2:三方设备
	DeviceType *int64 `json:"DeviceType,omitnil" name:"DeviceType"`
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
	DeviceInfos []*DeviceBaseInfo `json:"DeviceInfos,omitnil" name:"DeviceInfos"`

	// 设备总记录条数
	Length *int64 `json:"Length,omitnil" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type GetFlowPackagesRequestParams struct {
	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`

	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流量包绑定的设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 流量包状态，0：未生效，1：有效期内，2：已过期
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type GetFlowPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`

	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流量包绑定的设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 流量包状态，0：未生效，1：有效期内，2：已过期
	Status *int64 `json:"Status,omitnil" name:"Status"`
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
	PackageList []*FlowPackageInfo `json:"PackageList,omitnil" name:"PackageList"`

	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type GetFlowStatisticRequestParams struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil" name:"TimeGranularity"`
}

type GetFlowStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil" name:"TimeGranularity"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFlowStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFlowStatisticResponseParams struct {
	// 流量详细信息
	NetDetails []*NetDetails `json:"NetDetails,omitnil" name:"NetDetails"`

	// 查找时间段流量使用最大值（单位：byte）
	MaxValue *float64 `json:"MaxValue,omitnil" name:"MaxValue"`

	// 查找时间段流量使用平均值（单位：byte）
	AvgValue *float64 `json:"AvgValue,omitnil" name:"AvgValue"`

	// 查找时间段流量使用总量（单位：byte）
	TotalValue *float64 `json:"TotalValue,omitnil" name:"TotalValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type GetHardwareListRequestParams struct {
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 页面设备数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`
}

type GetHardwareListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 页面设备数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`
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
	HardwareInfos []*HardwareInfo `json:"HardwareInfos,omitnil" name:"HardwareInfos"`

	// 硬件总数
	Length *int64 `json:"Length,omitnil" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type GetMultiFlowStatisticRequestParams struct {
	// 设备id列表，单次最多请求10个设备
	DeviceIds []*string `json:"DeviceIds,omitnil" name:"DeviceIds"`

	// 1659514436
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 1659515000
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 统计流量类型（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 统计时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil" name:"TimeGranularity"`
}

type GetMultiFlowStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 设备id列表，单次最多请求10个设备
	DeviceIds []*string `json:"DeviceIds,omitnil" name:"DeviceIds"`

	// 1659514436
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 1659515000
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 统计流量类型（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 统计时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitnil" name:"TimeGranularity"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMultiFlowStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMultiFlowStatisticResponseParams struct {
	// 批量设备流量信息
	FlowDetails []*FlowDetails `json:"FlowDetails,omitnil" name:"FlowDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 统计指标（上行速率："TxRate":bit/s，下行速率："RxRate":bit/s，丢包："Loss":%，时延："RTT":ms）
	Metrics *string `json:"Metrics,omitnil" name:"Metrics"`
}

type GetNetMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 统计指标（上行速率："TxRate":bit/s，下行速率："RxRate":bit/s，丢包："Loss":%，时延："RTT":ms）
	Metrics *string `json:"Metrics,omitnil" name:"Metrics"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetNetMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetNetMonitorResponseParams struct {
	// 监控数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorData []*MonitorData `json:"MonitorData,omitnil" name:"MonitorData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 统计开始时间，单位：s
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 统计结束时间，单位：s
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 聚合粒度：
	// 1:按小时统计
	// 2:按天统计
	TimeGranularity *int64 `json:"TimeGranularity,omitnil" name:"TimeGranularity"`
}

type GetStatisticDataRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 统计开始时间，单位：s
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`

	// 统计结束时间，单位：s
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 聚合粒度：
	// 1:按小时统计
	// 2:按天统计
	TimeGranularity *int64 `json:"TimeGranularity,omitnil" name:"TimeGranularity"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetStatisticDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStatisticDataResponseParams struct {
	// 文件地址url
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 激活状态，
	// 空：全部；
	// 1:待激活；
	// 2:已激活；
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type GetVendorHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 激活状态，
	// 空：全部；
	// 1:待激活；
	// 2:已激活；
	Status *int64 `json:"Status,omitnil" name:"Status"`
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
	VendorHardware []*VendorHardware `json:"VendorHardware,omitnil" name:"VendorHardware"`

	// 设备总数
	Length *int64 `json:"Length,omitnil" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitnil" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type Hardware struct {
	// 硬件序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SN *string `json:"SN,omitnil" name:"SN"`

	// license计费模式：
	// 1，租户月付费
	// 2，厂商月付费
	// 3，license永久授权
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseChargingMode *int64 `json:"LicenseChargingMode,omitnil" name:"LicenseChargingMode"`

	// 设备描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 硬件ID，入参无需传递
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareId *string `json:"HardwareId,omitnil" name:"HardwareId"`
}

type HardwareInfo struct {
	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitnil" name:"ActiveTime"`

	// 最后在线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOnlineTime *string `json:"LastOnlineTime,omitnil" name:"LastOnlineTime"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 厂商备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	VendorDescription *string `json:"VendorDescription,omitnil" name:"VendorDescription"`

	// license计费模式： 1，租户月付费 2，厂商月付费 3，license永久授权
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseChargingMode *int64 `json:"LicenseChargingMode,omitnil" name:"LicenseChargingMode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 硬件序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SN *string `json:"SN,omitnil" name:"SN"`
}

// Predefined struct for user
type ModifyPackageRenewFlagRequestParams struct {
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 自动续费标识。true代表自动续费，false代表不自动续费
	RenewFlag *bool `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type ModifyPackageRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 自动续费标识。true代表自动续费，false代表不自动续费
	RenewFlag *bool `json:"RenewFlag,omitnil" name:"RenewFlag"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Time *string `json:"Time,omitnil" name:"Time"`

	// 业务指标（bps/ms/%）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessMetrics *float64 `json:"BusinessMetrics,omitnil" name:"BusinessMetrics"`

	// 网卡状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotNetInfo []*SlotNetInfo `json:"SlotNetInfo,omitnil" name:"SlotNetInfo"`
}

type NetDetails struct {
	// 流量值（bit）
	Current *float64 `json:"Current,omitnil" name:"Current"`

	// 时间点，单位：s
	Time *string `json:"Time,omitnil" name:"Time"`
}

type NetworkData struct {
	// 时延数组，最大长度30
	RTT []*float64 `json:"RTT,omitnil" name:"RTT"`

	// 丢包率
	Loss *float64 `json:"Loss,omitnil" name:"Loss"`

	// 抖动
	Jitter *float64 `json:"Jitter,omitnil" name:"Jitter"`

	// 10位秒级时间戳
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`
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
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// 流量包绑定的设备ID列表。捆绑设备个数上限取决于包的规格档位：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	DeviceList []*string `json:"DeviceList,omitnil" name:"DeviceList"`

	// 是否自动续费
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 区域标识，0：国内，1：国外
	PackageRegion *int64 `json:"PackageRegion,omitnil" name:"PackageRegion"`
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
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// 流量包绑定的设备ID列表。捆绑设备个数上限取决于包的规格档位：
	// 20G：最多绑定1个设备
	// 50G：最多绑定2个设备
	// 100G：最多绑定3个设备
	// 500G：最多绑定5个设备
	DeviceList []*string `json:"DeviceList,omitnil" name:"DeviceList"`

	// 是否自动续费
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 区域标识，0：国内，1：国外
	PackageRegion *int64 `json:"PackageRegion,omitnil" name:"PackageRegion"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OrderFlowPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OrderFlowPackageResponseParams struct {
	// 流量包的唯一资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type SlotNetInfo struct {
	// 网卡名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetInfoName *string `json:"NetInfoName,omitnil" name:"NetInfoName"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIP *string `json:"PublicIP,omitnil" name:"PublicIP"`

	// 指标数据（bps/ms/%）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Current *float64 `json:"Current,omitnil" name:"Current"`
}

type SrcAddressInfo struct {
	// 用户私网 ipv4 地址
	SrcIpv4 *string `json:"SrcIpv4,omitnil" name:"SrcIpv4"`

	// 用户公网 ipv4 地址
	SrcPublicIpv4 *string `json:"SrcPublicIpv4,omitnil" name:"SrcPublicIpv4"`

	// 用户 ipv6 地址
	SrcIpv6 *string `json:"SrcIpv6,omitnil" name:"SrcIpv6"`
}

// Predefined struct for user
type UpdateDeviceRequestParams struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 更新设备网络信息
	UpdateNetInfo []*UpdateNetInfo `json:"UpdateNetInfo,omitnil" name:"UpdateNetInfo"`
}

type UpdateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 设备备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 更新设备网络信息
	UpdateNetInfo []*UpdateNetInfo `json:"UpdateNetInfo,omitnil" name:"UpdateNetInfo"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type UpdateHardwareRequestParams struct {
	// 硬件ID
	HardwareId *string `json:"HardwareId,omitnil" name:"HardwareId"`

	// 硬件序列号
	SN *string `json:"SN,omitnil" name:"SN"`

	// 设备备注
	Description *string `json:"Description,omitnil" name:"Description"`
}

type UpdateHardwareRequest struct {
	*tchttp.BaseRequest
	
	// 硬件ID
	HardwareId *string `json:"HardwareId,omitnil" name:"HardwareId"`

	// 硬件序列号
	SN *string `json:"SN,omitnil" name:"SN"`

	// 设备备注
	Description *string `json:"Description,omitnil" name:"Description"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type UpdateNetInfo struct {
	// 网络类型：
	// 0:数据
	// 1:Wi-Fi
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 启用/禁用
	DataEnable *bool `json:"DataEnable,omitnil" name:"DataEnable"`

	// 上行限速：bit
	UploadLimit *uint64 `json:"UploadLimit,omitnil" name:"UploadLimit"`

	// 下行限速：bit
	DownloadLimit *uint64 `json:"DownloadLimit,omitnil" name:"DownloadLimit"`

	// 网卡名
	NetInfoName *string `json:"NetInfoName,omitnil" name:"NetInfoName"`
}

type VendorHardware struct {
	// 硬件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardwareId *string `json:"HardwareId,omitnil" name:"HardwareId"`

	// 硬件序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SN *string `json:"SN,omitnil" name:"SN"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 激活状态， 空：全部； 1:待激活； 2:已激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveTime *string `json:"ActiveTime,omitnil" name:"ActiveTime"`

	// 厂商备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 设备id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// license计费模式： 1，租户月付费 2，厂商月付费 3，license永久授权
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseChargingMode *int64 `json:"LicenseChargingMode,omitnil" name:"LicenseChargingMode"`

	// 最后在线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOnlineTime *string `json:"LastOnlineTime,omitnil" name:"LastOnlineTime"`
}