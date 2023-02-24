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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddDeviceRequestParams struct {
	// 新建设备的名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 新建设备的备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 新建设备的base64密钥字符串，非必选，如果不填写则由系统自动生成
	DataKey *string `json:"DataKey,omitempty" name:"DataKey"`

	// 是否设置预置密钥
	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`
}

type AddDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 新建设备的名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 新建设备的备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 新建设备的base64密钥字符串，非必选，如果不填写则由系统自动生成
	DataKey *string `json:"DataKey,omitempty" name:"DataKey"`

	// 是否设置预置密钥
	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`
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
	DataKey *string `json:"DataKey,omitempty" name:"DataKey"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 签名字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Signature *string `json:"Signature,omitempty" name:"Signature"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Capacity struct {
	// 电信鉴权的Token。要加速的电信手机终端访问 http://qos.189.cn/qos-api/getToken?appid=TencentCloud 页面，获取返回结果中result的值
	CTCCToken *string `json:"CTCCToken,omitempty" name:"CTCCToken"`

	// 终端所处在的省份，建议不填写由服务端自动获取，若需填写请填写带有省、市、自治区、特别行政区等后缀的省份中文全称
	Province *string `json:"Province,omitempty" name:"Province"`
}

type Context struct {
	// 测速数据
	NetworkData *NetworkData `json:"NetworkData,omitempty" name:"NetworkData"`

	// 用户期望最低门限
	ExpectedLowThreshold *ExpectedThreshold `json:"ExpectedLowThreshold,omitempty" name:"ExpectedLowThreshold"`

	// 用户期望最高门限
	ExpectedHighThreshold *ExpectedThreshold `json:"ExpectedHighThreshold,omitempty" name:"ExpectedHighThreshold"`
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
	EncryptedKey *string `json:"EncryptedKey,omitempty" name:"EncryptedKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SrcAddressInfo *SrcAddressInfo `json:"SrcAddressInfo,omitempty" name:"SrcAddressInfo"`

	// 加速业务目标地址信息
	DestAddressInfo *DestAddressInfo `json:"DestAddressInfo,omitempty" name:"DestAddressInfo"`

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
	QosMenu *string `json:"QosMenu,omitempty" name:"QosMenu"`

	// 申请加速的设备信息，包括运营商，操作系统，设备唯一标识等。
	DeviceInfo *DeviceInfo `json:"DeviceInfo,omitempty" name:"DeviceInfo"`

	// 期望加速时长（单位分钟），默认值30分钟
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 接口能力扩展，如果是电信用户，必须填充CTCC Token字段
	Capacity *Capacity `json:"Capacity,omitempty" name:"Capacity"`

	// 应用模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 针对特殊协议进行加速
	// 1. IP （默认值）
	// 2. UDP
	// 3. TCP
	Protocol *uint64 `json:"Protocol,omitempty" name:"Protocol"`

	// 加速策略关键数据
	Context *Context `json:"Context,omitempty" name:"Context"`

	// 签名
	Extern *string `json:"Extern,omitempty" name:"Extern"`
}

type CreateQosRequest struct {
	*tchttp.BaseRequest
	
	// 加速业务源地址信息，SrcIpv6和（SrcIpv4+SrcPublicIpv4）二选一，目前Ipv6不可用，全部填写以Ipv4参数为准。
	SrcAddressInfo *SrcAddressInfo `json:"SrcAddressInfo,omitempty" name:"SrcAddressInfo"`

	// 加速业务目标地址信息
	DestAddressInfo *DestAddressInfo `json:"DestAddressInfo,omitempty" name:"DestAddressInfo"`

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
	QosMenu *string `json:"QosMenu,omitempty" name:"QosMenu"`

	// 申请加速的设备信息，包括运营商，操作系统，设备唯一标识等。
	DeviceInfo *DeviceInfo `json:"DeviceInfo,omitempty" name:"DeviceInfo"`

	// 期望加速时长（单位分钟），默认值30分钟
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 接口能力扩展，如果是电信用户，必须填充CTCC Token字段
	Capacity *Capacity `json:"Capacity,omitempty" name:"Capacity"`

	// 应用模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 针对特殊协议进行加速
	// 1. IP （默认值）
	// 2. UDP
	// 3. TCP
	Protocol *uint64 `json:"Protocol,omitempty" name:"Protocol"`

	// 加速策略关键数据
	Context *Context `json:"Context,omitempty" name:"Context"`

	// 签名
	Extern *string `json:"Extern,omitempty" name:"Extern"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 当前加速剩余时长（单位秒）
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 删除设备的唯一ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type DeleteQosRequest struct {
	*tchttp.BaseRequest
	
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 本次加速会话持续时间（单位秒）
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type DescribeQosRequest struct {
	*tchttp.BaseRequest
	
	// 单次加速唯一 Id
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
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
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 手机公网出口IP，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcPublicIpv4 *string `json:"SrcPublicIpv4,omitempty" name:"SrcPublicIpv4"`

	// 业务访问目的IP，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestIpv4 []*string `json:"DestIpv4,omitempty" name:"DestIpv4"`

	// 当前加速剩余时长（单位秒）有，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 加速套餐类型，仅匹配时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	QosMenu *string `json:"QosMenu,omitempty" name:"QosMenu"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DestIp []*string `json:"DestIp,omitempty" name:"DestIp"`
}

type DeviceBaseInfo struct {
	// 设备唯一ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备创建的时间，单位：ms
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备最后在线时间，单位：ms
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 设备的备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DeviceDetails struct {
	// 设备基本信息
	DeviceBaseInfo *DeviceBaseInfo `json:"DeviceBaseInfo,omitempty" name:"DeviceBaseInfo"`

	// 设备网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceNetInfo []*DeviceNetInfo `json:"DeviceNetInfo,omitempty" name:"DeviceNetInfo"`

	// 聚合服务器地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewaySite *string `json:"GatewaySite,omitempty" name:"GatewaySite"`

	// 业务下行速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessDownRate *float64 `json:"BusinessDownRate,omitempty" name:"BusinessDownRate"`

	// 业务上行速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessUpRate *float64 `json:"BusinessUpRate,omitempty" name:"BusinessUpRate"`
}

type DeviceInfo struct {
	// 运营商
	// 1：移动 
	// 2：电信
	// 3：联通
	// 4：广电
	// 99：其他
	Vendor *uint64 `json:"Vendor,omitempty" name:"Vendor"`

	// 设备操作系统：
	// 1：Android
	// 2： IOS
	// 99：其他
	OS *uint64 `json:"OS,omitempty" name:"OS"`

	// 设备唯一标识
	// IOS 填写 IDFV
	// Android 填写 IMEI
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 用户手机号码
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 无线信息
	// 1：4G
	// 2：5G
	// 3：WIFI
	// 99：其他
	Wireless *uint64 `json:"Wireless,omitempty" name:"Wireless"`
}

type DeviceNetInfo struct {
	// 网络类型：
	// 0:数据
	// 1:Wi-Fi
	// 2:有线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 启用/禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEnable *bool `json:"DataEnable,omitempty" name:"DataEnable"`

	// 上行限速
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadLimit *string `json:"UploadLimit,omitempty" name:"UploadLimit"`

	// 下行限速
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadLimit *string `json:"DownloadLimit,omitempty" name:"DownloadLimit"`

	// 接收实时速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRx *uint64 `json:"DataRx,omitempty" name:"DataRx"`

	// 发送实时速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTx *uint64 `json:"DataTx,omitempty" name:"DataTx"`

	// 运营商类型：
	// 1: 中国移动；
	// 2: 中国电信; 
	// 3: 中国联通
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vendor *int64 `json:"Vendor,omitempty" name:"Vendor"`

	// 连接状态：
	// 0:无连接
	// 1:连接中
	// 2:已连接
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitempty" name:"State"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 信号强度/单位：dbm
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignalStrength *int64 `json:"SignalStrength,omitempty" name:"SignalStrength"`

	// 数据网络类型：
	// -1 ：无效值   
	// 2：2G 
	// 3：3G 
	// 4：4G 
	// 5：5G
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rat *int64 `json:"Rat,omitempty" name:"Rat"`

	// 网卡名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetInfoName *string `json:"NetInfoName,omitempty" name:"NetInfoName"`

	// 下行实时速率（浮点数类型代替上一版本DataRx的整型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownRate *float64 `json:"DownRate,omitempty" name:"DownRate"`

	// 上行实时速率（浮点数类型代替上一版本TxRate的整型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpRate *float64 `json:"UpRate,omitempty" name:"UpRate"`
}

type ExpectedThreshold struct {
	// 期望发起加速的时延阈值
	RTT *float64 `json:"RTT,omitempty" name:"RTT"`

	// 期望发起加速的丢包率阈值
	Loss *float64 `json:"Loss,omitempty" name:"Loss"`

	// 期望发起加速的抖动阈值
	Jitter *float64 `json:"Jitter,omitempty" name:"Jitter"`
}

type FlowDetails struct {
	// 流量数据点
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetails []*NetDetails `json:"NetDetails,omitempty" name:"NetDetails"`

	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 流量最大值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxValue *float64 `json:"MaxValue,omitempty" name:"MaxValue"`

	// 流量平均值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgValue *float64 `json:"AvgValue,omitempty" name:"AvgValue"`

	// 流量总值（单位：bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalValue *float64 `json:"TotalValue,omitempty" name:"TotalValue"`
}

// Predefined struct for user
type GetDeviceRequestParams struct {
	// 搜索指定设备的id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
}

type GetDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 搜索指定设备的id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`
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
	DeviceDetails *DeviceDetails `json:"DeviceDetails,omitempty" name:"DeviceDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 搜索设备的关键字（ID或者设备名），为空时匹配所有设备
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type GetDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示记录数，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前查看页码，PageSize、PageNumber值均为-1 时，按照1页无限制条数匹配所有设备
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 搜索设备的关键字（ID或者设备名），为空时匹配所有设备
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicesResponseParams struct {
	// 设备信息列表
	DeviceInfos []*DeviceBaseInfo `json:"DeviceInfos,omitempty" name:"DeviceInfos"`

	// 设备总记录条数
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 总页数
	TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type GetFlowStatisticRequestParams struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`
}

type GetFlowStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 开始查找时间
	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 截止时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 流量种类（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`
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
	NetDetails []*NetDetails `json:"NetDetails,omitempty" name:"NetDetails"`

	// 查找时间段流量使用最大值（单位：bit）
	MaxValue *float64 `json:"MaxValue,omitempty" name:"MaxValue"`

	// 查找时间段流量使用平均值（单位：bit）
	AvgValue *float64 `json:"AvgValue,omitempty" name:"AvgValue"`

	// 查找时间段流量使用总量（单位：bit）
	TotalValue *float64 `json:"TotalValue,omitempty" name:"TotalValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type GetMultiFlowStatisticRequestParams struct {
	// 设备id列表，单次最多请求10个设备
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`

	// 1659514436
	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 1659515000
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 统计流量类型（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 统计时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`
}

type GetMultiFlowStatisticRequest struct {
	*tchttp.BaseRequest
	
	// 设备id列表，单次最多请求10个设备
	DeviceIds []*string `json:"DeviceIds,omitempty" name:"DeviceIds"`

	// 1659514436
	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 1659515000
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 统计流量类型（1：上行流量，2：下行流量）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 统计时间粒度（1：按小时统计，2：按天统计）
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`
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
	FlowDetails []*FlowDetails `json:"FlowDetails,omitempty" name:"FlowDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 统计开始时间，单位：s
	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 统计结束时间，单位：s
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 聚合粒度：
	// 1:按小时统计
	// 2:按天统计
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`
}

type GetStatisticDataRequest struct {
	*tchttp.BaseRequest
	
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 统计开始时间，单位：s
	BeginTime *int64 `json:"BeginTime,omitempty" name:"BeginTime"`

	// 统计结束时间，单位：s
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 聚合粒度：
	// 1:按小时统计
	// 2:按天统计
	TimeGranularity *int64 `json:"TimeGranularity,omitempty" name:"TimeGranularity"`
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
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type NetDetails struct {
	// 流量值（bit）
	Current *float64 `json:"Current,omitempty" name:"Current"`

	// 时间点，单位：s
	Time *string `json:"Time,omitempty" name:"Time"`
}

type NetworkData struct {
	// 时延数组，最大长度30
	RTT []*float64 `json:"RTT,omitempty" name:"RTT"`

	// 丢包率
	Loss *float64 `json:"Loss,omitempty" name:"Loss"`

	// 抖动
	Jitter *float64 `json:"Jitter,omitempty" name:"Jitter"`

	// 10位秒级时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type SrcAddressInfo struct {
	// 用户私网 ipv4 地址
	SrcIpv4 *string `json:"SrcIpv4,omitempty" name:"SrcIpv4"`

	// 用户公网 ipv4 地址
	SrcPublicIpv4 *string `json:"SrcPublicIpv4,omitempty" name:"SrcPublicIpv4"`

	// 用户 ipv6 地址
	SrcIpv6 *string `json:"SrcIpv6,omitempty" name:"SrcIpv6"`
}

// Predefined struct for user
type UpdateDeviceRequestParams struct {
	// 设备id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 更新设备网络信息
	UpdateNetInfo []*UpdateNetInfo `json:"UpdateNetInfo,omitempty" name:"UpdateNetInfo"`
}

type UpdateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 设备id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 更新设备网络信息
	UpdateNetInfo []*UpdateNetInfo `json:"UpdateNetInfo,omitempty" name:"UpdateNetInfo"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type UpdateNetInfo struct {
	// 网络类型：
	// 0:数据
	// 1:Wi-Fi
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 启用/禁用
	DataEnable *bool `json:"DataEnable,omitempty" name:"DataEnable"`

	// 上行限速：bit
	UploadLimit *uint64 `json:"UploadLimit,omitempty" name:"UploadLimit"`

	// 下行限速：bit
	DownloadLimit *uint64 `json:"DownloadLimit,omitempty" name:"DownloadLimit"`

	// 网卡名
	NetInfoName *string `json:"NetInfoName,omitempty" name:"NetInfoName"`
}