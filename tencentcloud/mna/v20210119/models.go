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

type Capacity struct {

	// 电信鉴权的Token。要加速的电信手机终端访问 http://qos.189.cn/qos-api/getToken?appid=TencentCloud 页面，获取返回结果中result的值
	CTCCToken *string `json:"CTCCToken,omitempty" name:"CTCCToken"`

	// 终端所处在的省份，建议不填写由服务端自动获取，若需填写请填写带有省、市、自治区、特别行政区等后缀的省份中文全称
	Province *string `json:"Province,omitempty" name:"Province"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateQosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 单次加速唯一 Id
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 当前加速剩余时长（单位秒）
		Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteQosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 单次加速唯一 Id
		SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

		// 本次加速会话持续时间（单位秒）
		Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DestAddressInfo struct {

	// 加速业务目标 ip 地址数组
	DestIp []*string `json:"DestIp,omitempty" name:"DestIp"`
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

type SrcAddressInfo struct {

	// 用户私网 ipv4 地址
	SrcIpv4 *string `json:"SrcIpv4,omitempty" name:"SrcIpv4"`

	// 用户公网 ipv4 地址
	SrcPublicIpv4 *string `json:"SrcPublicIpv4,omitempty" name:"SrcPublicIpv4"`

	// 用户 ipv6 地址
	SrcIpv6 *string `json:"SrcIpv6,omitempty" name:"SrcIpv6"`
}
