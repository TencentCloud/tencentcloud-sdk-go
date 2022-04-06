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

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Attribute struct {

	// 属性列表
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`
}

type BindProductInfo struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type CLSLogItem struct {

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 请求ID
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 结果
	Result *string `json:"Result,omitempty" name:"Result"`

	// 模块
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 日志时间
	Time *string `json:"Time,omitempty" name:"Time"`

	// 腾讯云账号
	Userid *string `json:"Userid,omitempty" name:"Userid"`
}

type CertInfo struct {

	// 证书名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 证书的序列号，16进制编码
	CertSN *string `json:"CertSN,omitempty" name:"CertSN"`

	// 证书颁发着名称
	IssuerName *string `json:"IssuerName,omitempty" name:"IssuerName"`

	// 证书主题
	Subject *string `json:"Subject,omitempty" name:"Subject"`

	// 证书创建时间，秒级时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 证书生效时间，秒级时间戳
	EffectiveTime *uint64 `json:"EffectiveTime,omitempty" name:"EffectiveTime"`

	// 证书失效时间，秒级时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// X509证书内容
	CertText *string `json:"CertText,omitempty" name:"CertText"`
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品 ID 。创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}。
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备属性
	Attribute *Attribute `json:"Attribute,omitempty" name:"Attribute"`

	// 是否使用自定义PSK，默认不使用
	DefinedPsk *string `json:"DefinedPsk,omitempty" name:"DefinedPsk"`

	// 运营商类型，当产品是NB-IoT产品时，此字段必填。1表示中国电信，2表示中国移动，3表示中国联通
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// IMEI，当产品是NB-IoT产品时，此字段必填
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// LoRa设备的DevEui，当创建LoRa时，此字段必填
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// LoRa设备的MoteType
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// 创建LoRa设备需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`

	// LoRa设备的AppKey
	LoraAppKey *string `json:"LoraAppKey,omitempty" name:"LoraAppKey"`

	// 私有CA创建的设备证书
	TlsCrt *string `json:"TlsCrt,omitempty" name:"TlsCrt"`
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
	delete(f, "Attribute")
	delete(f, "DefinedPsk")
	delete(f, "Isp")
	delete(f, "Imei")
	delete(f, "LoraDevEui")
	delete(f, "LoraMoteType")
	delete(f, "Skey")
	delete(f, "LoraAppKey")
	delete(f, "TlsCrt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备名称
		DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

		// 对称加密密钥，base64编码。采用对称加密时返回该参数
		DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

		// 设备证书，用于 TLS 建立链接时校验客户端身份。采用非对称加密时返回该参数
		DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

		// 设备私钥，用于 TLS 建立链接时校验客户端身份，腾讯云后台不保存，请妥善保管。采用非对称加密时返回该参数
		DevicePrivateKey *string `json:"DevicePrivateKey,omitempty" name:"DevicePrivateKey"`

		// LoRa设备的DevEui，当设备是LoRa设备时，会返回该字段
		LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

		// LoRa设备的MoteType，当设备是LoRa设备时，会返回该字段
		LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

		// LoRa设备的AppKey，当设备是LoRa设备时，会返回该字段
		LoraAppKey *string `json:"LoraAppKey,omitempty" name:"LoraAppKey"`

		// LoRa设备的NwkKey，当设备是LoRa设备时，会返回该字段
		LoraNwkKey *string `json:"LoraNwkKey,omitempty" name:"LoraNwkKey"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreatePrivateCARequest struct {
	*tchttp.BaseRequest

	// CA证书名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// CA证书内容
	CertText *string `json:"CertText,omitempty" name:"CertText"`

	// 校验CA证书的证书内容
	VerifyCertText *string `json:"VerifyCertText,omitempty" name:"VerifyCertText"`
}

func (r *CreatePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	delete(f, "CertText")
	delete(f, "VerifyCertText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest

	// 设备所属的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 需要删除的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 删除LoRa设备以及LoRa网关设备需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
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
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeletePrivateCARequest struct {
	*tchttp.BaseRequest

	// 私有CA证书名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`
}

func (r *DeletePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductPrivateCARequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DeleteProductPrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductPrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProductPrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductPrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductPrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductPrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest

	// 需要删除的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 删除LoRa产品需要skey
	Skey *string `json:"Skey,omitempty" name:"Skey"`
}

func (r *DeleteProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Skey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备名
		DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

		// 设备是否在线，0不在线，1在线
		Online *uint64 `json:"Online,omitempty" name:"Online"`

		// 设备登录时间
		LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

		// 设备固件版本
		Version *string `json:"Version,omitempty" name:"Version"`

		// 设备最后更新时间
		LastUpdateTime *uint64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

		// 设备证书
		DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

		// 设备密钥
		DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

		// 设备属性
		Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`

		// 设备类型
		DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

		// 国际移动设备识别码 IMEI
		Imei *string `json:"Imei,omitempty" name:"Imei"`

		// 运营商类型
		Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

		// IP地址
		ConnIP *uint64 `json:"ConnIP,omitempty" name:"ConnIP"`

		// NB IoT运营商处的DeviceID
		NbiotDeviceID *string `json:"NbiotDeviceID,omitempty" name:"NbiotDeviceID"`

		// Lora设备的dev eui
		LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

		// Lora设备的mote type
		LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

		// 设备的sdk日志等级
	// 注意：此字段可能返回 null，表示取不到有效值。
		LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`

		// 首次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitempty" name:"FirstOnlineTime"`

		// 最近下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		LastOfflineTime *uint64 `json:"LastOfflineTime,omitempty" name:"LastOfflineTime"`

		// 设备创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

		// 设备证书获取状态，0 未获取过设备密钥, 1 已获取过设备密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
		CertState *uint64 `json:"CertState,omitempty" name:"CertState"`

		// 设备启用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
		EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

		// 设备标签
	// 注意：此字段可能返回 null，表示取不到有效值。
		Labels []*DeviceLabel `json:"Labels,omitempty" name:"Labels"`

		// MQTT客户端IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

		// 设备固件更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitempty" name:"FirmwareUpdateTime"`

		// 创建者账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateUserId *uint64 `json:"CreateUserId,omitempty" name:"CreateUserId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest

	// 需要查看设备列表的产品 ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 偏移量，Offset从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页的大小，数值范围 10-250
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 设备固件版本号，若不带此参数会返回所有固件版本的设备。传"None-FirmwareVersion"查询无版本号的设备
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" name:"FirmwareVersion"`

	// 需要过滤的设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备是否启用，0禁用状态1启用状态，默认不区分
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FirmwareVersion")
	delete(f, "DeviceName")
	delete(f, "EnableState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 设备详细信息列表
		Devices []*DeviceInfo `json:"Devices,omitempty" name:"Devices"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateCABindedProductsRequest struct {
	*tchttp.BaseRequest

	// 证书名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询的数据量，默认为20， 最大为200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrivateCABindedProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCABindedProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateCABindedProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateCABindedProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有CA绑定的产品列表
		Products []*BindProductInfo `json:"Products,omitempty" name:"Products"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateCABindedProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCABindedProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateCARequest struct {
	*tchttp.BaseRequest

	// 私有化CA名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`
}

func (r *DescribePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有化CA详情
		CA *CertInfo `json:"CA,omitempty" name:"CA"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateCAsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePrivateCAsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCAsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateCAsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateCAsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 私有CA证书列表
		CAs []*CertInfo `json:"CAs,omitempty" name:"CAs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateCAsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateCAsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductCARequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductCAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CA证书列表
		CAs []*CertInfo `json:"CAs,omitempty" name:"CAs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品ID
		ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

		// 产品名
		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

		// 产品元数据
		ProductMetadata *ProductMetadata `json:"ProductMetadata,omitempty" name:"ProductMetadata"`

		// 产品属性
		ProductProperties *ProductProperties `json:"ProductProperties,omitempty" name:"ProductProperties"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceInfo struct {

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备是否在线，0不在线，1在线
	Online *uint64 `json:"Online,omitempty" name:"Online"`

	// 设备登录时间
	LoginTime *uint64 `json:"LoginTime,omitempty" name:"LoginTime"`

	// 设备版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 设备证书，证书加密的设备返回
	DeviceCert *string `json:"DeviceCert,omitempty" name:"DeviceCert"`

	// 设备密钥，密钥加密的设备返回
	DevicePsk *string `json:"DevicePsk,omitempty" name:"DevicePsk"`

	// 设备属性
	Tags []*DeviceTag `json:"Tags,omitempty" name:"Tags"`

	// 设备类型
	DeviceType *uint64 `json:"DeviceType,omitempty" name:"DeviceType"`

	// 国际移动设备识别码 IMEI
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 运营商类型
	Isp *uint64 `json:"Isp,omitempty" name:"Isp"`

	// NB IOT运营商处的DeviceID
	NbiotDeviceID *string `json:"NbiotDeviceID,omitempty" name:"NbiotDeviceID"`

	// IP地址
	ConnIP *uint64 `json:"ConnIP,omitempty" name:"ConnIP"`

	// 设备最后更新时间
	LastUpdateTime *uint64 `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// LoRa设备的dev eui
	LoraDevEui *string `json:"LoraDevEui,omitempty" name:"LoraDevEui"`

	// LoRa设备的Mote type
	LoraMoteType *uint64 `json:"LoraMoteType,omitempty" name:"LoraMoteType"`

	// 首次上线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOnlineTime *uint64 `json:"FirstOnlineTime,omitempty" name:"FirstOnlineTime"`

	// 最近下线时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOfflineTime *uint64 `json:"LastOfflineTime,omitempty" name:"LastOfflineTime"`

	// 设备创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 设备日志级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`

	// 设备证书获取状态, 1 已获取过设备密钥，0 未获取过设备密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertState *uint64 `json:"CertState,omitempty" name:"CertState"`

	// 设备可用状态，0禁用，1启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableState *uint64 `json:"EnableState,omitempty" name:"EnableState"`

	// 设备标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*DeviceLabel `json:"Labels,omitempty" name:"Labels"`

	// MQTT客户端IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// ota最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirmwareUpdateTime *uint64 `json:"FirmwareUpdateTime,omitempty" name:"FirmwareUpdateTime"`
}

type DeviceLabel struct {

	// 标签标识
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeviceTag struct {

	// 属性名称
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 属性值的类型，1 int，2 string
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 属性的值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 属性描述名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ListLogPayloadRequest struct {
	*tchttp.BaseRequest

	// 日志开始时间
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。键值或文本可以包含多个，以空格隔开。其中可以索引的key比如：RequestID、ProductID、DeviceName等。
	// 一个典型的查询示例：ProductID:ABCDE12345 DeviceName:test publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志最大条数
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

func (r *ListLogPayloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogPayloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "Keywords")
	delete(f, "Context")
	delete(f, "MaxNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLogPayloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListLogPayloadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志上下文
		Context *string `json:"Context,omitempty" name:"Context"`

		// 是否还有日志，如有仍有日志，下次查询的请求带上当前请求返回的Context
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 日志列表
		Results []*PayloadLogItem `json:"Results,omitempty" name:"Results"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogPayloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogPayloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogRequest struct {
	*tchttp.BaseRequest

	// 日志开始时间
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。键值或文本可以包含多个，以空格隔开。其中可以索引的key包括：requestid、productid、devicename、scene、content。
	// 一个典型的查询示例：productid:ABCDE12345 devicename:test scene:SHADOW content:Device%20connect publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询条数
	MaxNum *int64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

func (r *ListLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "Keywords")
	delete(f, "Context")
	delete(f, "MaxNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志上下文
		Context *string `json:"Context,omitempty" name:"Context"`

		// 是否还有日志，如有仍有日志，下次查询的请求带上当前请求返回的Context
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 日志列表
		Results []*CLSLogItem `json:"Results,omitempty" name:"Results"`

		// 日志总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSDKLogRequest struct {
	*tchttp.BaseRequest

	// 日志开始时间
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 日志结束时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 查询关键字，可以同时支持键值查询和文本查询，
	// 例如，查询某key的值为value，并且包含某word的日志，该参数为：key:value word。
	// 键值或文本可以包含多个，以空格隔开。
	// 其中可以索引的key包括：productid、devicename、loglevel
	// 一个典型的查询示例：productid:7JK1G72JNE devicename:name publish loglevel:WARN一个典型的查询示例：productid:ABCDE12345 devicename:test scene:SHADOW publish
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 日志检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询条数
	MaxNum *uint64 `json:"MaxNum,omitempty" name:"MaxNum"`
}

func (r *ListSDKLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSDKLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinTime")
	delete(f, "MaxTime")
	delete(f, "Keywords")
	delete(f, "Context")
	delete(f, "MaxNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSDKLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListSDKLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志检索上下文
		Context *string `json:"Context,omitempty" name:"Context"`

		// 是否还有日志，如有仍有日志，下次查询的请求带上当前请求返回的Context
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 日志列表
		Results []*SDKLogItem `json:"Results,omitempty" name:"Results"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSDKLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSDKLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PayloadLogItem struct {

	// 账号id
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 来源类型
	SrcType *string `json:"SrcType,omitempty" name:"SrcType"`

	// 来源名称
	SrcName *string `json:"SrcName,omitempty" name:"SrcName"`

	// 消息topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 内容格式类型
	PayloadFormatType *string `json:"PayloadFormatType,omitempty" name:"PayloadFormatType"`

	// 内容信息
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 请求ID
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`

	// 日期时间
	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`
}

type ProductMetadata struct {

	// 产品创建时间
	CreationDate *uint64 `json:"CreationDate,omitempty" name:"CreationDate"`
}

type ProductProperties struct {

	// 产品描述
	ProductDescription *string `json:"ProductDescription,omitempty" name:"ProductDescription"`

	// 加密类型，1表示证书认证，2表示签名认证。如不填写，默认值是1
	EncryptionType *string `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 产品所属区域，目前只支持广州（gz）
	Region *string `json:"Region,omitempty" name:"Region"`

	// 产品类型，各个类型值代表的节点-类型如下：
	// 0 普通产品，2 NB-IoT产品，4 LoRa产品，3 LoRa网关产品，5 普通网关产品   默认值是0
	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`

	// 数据格式，取值为json或者custom，默认值是json
	Format *string `json:"Format,omitempty" name:"Format"`

	// 产品所属平台，默认值是0
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// LoRa产品运营侧APPEUI，只有LoRa产品需要填写
	Appeui *string `json:"Appeui,omitempty" name:"Appeui"`

	// 产品绑定的物模型ID，-1表示不绑定
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 产品绑定的物模型名称
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`

	// 产品密钥，suite产品才会有
	ProductKey *string `json:"ProductKey,omitempty" name:"ProductKey"`

	// 动态注册类型 0-关闭, 1-预定义设备名 2-动态定义设备名
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册产品秘钥
	ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

	// RegisterType为2时，设备动态创建的限制数量
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`

	// 划归的产品，展示为源产品ID，其余为空
	OriginProductId *string `json:"OriginProductId,omitempty" name:"OriginProductId"`

	// 私有CA名称
	PrivateCAName *string `json:"PrivateCAName,omitempty" name:"PrivateCAName"`

	// 划归的产品，展示为源用户ID，其余为空
	OriginUserId *uint64 `json:"OriginUserId,omitempty" name:"OriginUserId"`
}

type PublishBroadcastMessageRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 消息内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// 消息质量等级
	Qos *int64 `json:"Qos,omitempty" name:"Qos"`

	// Payload内容的编码格式，取值为base64或空。base64表示云端将收到的请求数据进行base64解码后下发到设备，空则直接将原始内容下发到设备
	PayloadEncoding *string `json:"PayloadEncoding,omitempty" name:"PayloadEncoding"`
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

type PublishBroadcastMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 广播消息任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type SDKLogItem struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 日志等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 日志时间
	DateTime *string `json:"DateTime,omitempty" name:"DateTime"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type SetProductsForbiddenStatusRequest struct {
	*tchttp.BaseRequest

	// 要设置禁用状态的产品列表
	ProductId []*string `json:"ProductId,omitempty" name:"ProductId"`

	// 0启用，1禁用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *SetProductsForbiddenStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetProductsForbiddenStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetProductsForbiddenStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetProductsForbiddenStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetProductsForbiddenStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetProductsForbiddenStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDeviceLogLevelRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 日志级别，0：关闭，1：错误，2：告警，3：信息，4：调试
	LogLevel *uint64 `json:"LogLevel,omitempty" name:"LogLevel"`
}

func (r *UpdateDeviceLogLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceLogLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "LogLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceLogLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDeviceLogLevelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDeviceLogLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceLogLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDevicePSKRequest struct {
	*tchttp.BaseRequest

	// 产品名
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 设备的psk
	Psk *string `json:"Psk,omitempty" name:"Psk"`
}

func (r *UpdateDevicePSKRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicePSKRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "Psk")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDevicePSKRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDevicePSKResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateDevicePSKResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDevicePSKResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDevicesEnableStateRequest struct {
	*tchttp.BaseRequest

	// 设备所属产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称集合
	DeviceNames []*string `json:"DeviceNames,omitempty" name:"DeviceNames"`

	// 要设置的设备状态，1为启用，0为禁用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
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
	delete(f, "ProductId")
	delete(f, "DeviceNames")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDevicesEnableStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateDevicesEnableStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type UpdatePrivateCARequest struct {
	*tchttp.BaseRequest

	// CA证书名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// CA证书内容
	CertText *string `json:"CertText,omitempty" name:"CertText"`

	// 校验CA证书的证书内容
	VerifyCertText *string `json:"VerifyCertText,omitempty" name:"VerifyCertText"`
}

func (r *UpdatePrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertName")
	delete(f, "CertText")
	delete(f, "VerifyCertText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductDynamicRegisterRequest struct {
	*tchttp.BaseRequest

	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
	RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

	// 动态注册设备上限
	RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`
}

func (r *UpdateProductDynamicRegisterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductDynamicRegisterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "RegisterType")
	delete(f, "RegisterLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProductDynamicRegisterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductDynamicRegisterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 动态注册类型，0-关闭 1-预创建设备 2-自动创建设备
		RegisterType *uint64 `json:"RegisterType,omitempty" name:"RegisterType"`

		// 动态注册产品密钥
		ProductSecret *string `json:"ProductSecret,omitempty" name:"ProductSecret"`

		// 动态注册设备上限
		RegisterLimit *uint64 `json:"RegisterLimit,omitempty" name:"RegisterLimit"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductDynamicRegisterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductDynamicRegisterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductPrivateCARequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 私有CA证书名称
	CertName *string `json:"CertName,omitempty" name:"CertName"`
}

func (r *UpdateProductPrivateCARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductPrivateCARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "CertName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProductPrivateCARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateProductPrivateCAResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductPrivateCAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProductPrivateCAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
