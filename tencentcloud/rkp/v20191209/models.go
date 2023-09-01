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

package v20191209

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type DevInfoQ struct {
	// devid
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 风险值
	RiskScore *int64 `json:"RiskScore,omitnil" name:"RiskScore"`

	// 风险详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskInfo []*RiskDetail `json:"RiskInfo,omitnil" name:"RiskInfo"`

	// 概率值
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

// Predefined struct for user
type GetOpenIdRequestParams struct {
	// dev临时token，通过sdk接口获取
	DeviceToken *string `json:"DeviceToken,omitnil" name:"DeviceToken"`

	// 业务ID
	BusinessId *int64 `json:"BusinessId,omitnil" name:"BusinessId"`

	// 业务侧账号体系下的用户ID
	BusinessUserId *string `json:"BusinessUserId,omitnil" name:"BusinessUserId"`

	// 平台：0-Android， 1-iOS， 2-web
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 选项
	Option *string `json:"Option,omitnil" name:"Option"`
}

type GetOpenIdRequest struct {
	*tchttp.BaseRequest
	
	// dev临时token，通过sdk接口获取
	DeviceToken *string `json:"DeviceToken,omitnil" name:"DeviceToken"`

	// 业务ID
	BusinessId *int64 `json:"BusinessId,omitnil" name:"BusinessId"`

	// 业务侧账号体系下的用户ID
	BusinessUserId *string `json:"BusinessUserId,omitnil" name:"BusinessUserId"`

	// 平台：0-Android， 1-iOS， 2-web
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 选项
	Option *string `json:"Option,omitnil" name:"Option"`
}

func (r *GetOpenIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpenIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "BusinessId")
	delete(f, "BusinessUserId")
	delete(f, "Platform")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpenIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpenIdResponseParams struct {
	// 设备ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil" name:"OpenId"`

	// 设备风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskInfo []*RiskInfo `json:"RiskInfo,omitnil" name:"RiskInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetOpenIdResponse struct {
	*tchttp.BaseResponse
	Response *GetOpenIdResponseParams `json:"Response"`
}

func (r *GetOpenIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpenIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTokenRequestParams struct {
	// 业务ID
	BusinessId *int64 `json:"BusinessId,omitnil" name:"BusinessId"`

	// 业务子场景
	Scene *int64 `json:"Scene,omitnil" name:"Scene"`

	// 业务侧账号体系下的用户ID
	BusinessUserId *string `json:"BusinessUserId,omitnil" name:"BusinessUserId"`

	// 用户侧的IP
	AppClientIp *string `json:"AppClientIp,omitnil" name:"AppClientIp"`

	// 过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 上一个token
	OldToken *string `json:"OldToken,omitnil" name:"OldToken"`
}

type GetTokenRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID
	BusinessId *int64 `json:"BusinessId,omitnil" name:"BusinessId"`

	// 业务子场景
	Scene *int64 `json:"Scene,omitnil" name:"Scene"`

	// 业务侧账号体系下的用户ID
	BusinessUserId *string `json:"BusinessUserId,omitnil" name:"BusinessUserId"`

	// 用户侧的IP
	AppClientIp *string `json:"AppClientIp,omitnil" name:"AppClientIp"`

	// 过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 上一个token
	OldToken *string `json:"OldToken,omitnil" name:"OldToken"`
}

func (r *GetTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	delete(f, "Scene")
	delete(f, "BusinessUserId")
	delete(f, "AppClientIp")
	delete(f, "ExpireTime")
	delete(f, "OldToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTokenResponseParams struct {
	// 返回token
	Token *string `json:"Token,omitnil" name:"Token"`

	// 过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetTokenResponseParams `json:"Response"`
}

func (r *GetTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDevAndRiskRequestParams struct {
	// 设备类型 0表示Android， 1表示IOS
	DevType *int64 `json:"DevType,omitnil" name:"DevType"`

	// Android Imei号
	Imei *string `json:"Imei,omitnil" name:"Imei"`

	// Mac地址
	Mac *string `json:"Mac,omitnil" name:"Mac"`

	// android  Aid
	Aid *string `json:"Aid,omitnil" name:"Aid"`

	// Android Cid
	Cid *string `json:"Cid,omitnil" name:"Cid"`

	// 手机Imsi
	Imsi *string `json:"Imsi,omitnil" name:"Imsi"`

	// Df 磁盘分区信息
	Df *string `json:"Df,omitnil" name:"Df"`

	// 内核版本
	KernelVer *string `json:"KernelVer,omitnil" name:"KernelVer"`

	// 存储大小
	Storage *string `json:"Storage,omitnil" name:"Storage"`

	// 设备驱动指纹
	Dfp *string `json:"Dfp,omitnil" name:"Dfp"`

	// 启动时间
	BootTime *string `json:"BootTime,omitnil" name:"BootTime"`

	// 分辨率 水平*垂直 格式
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`

	// 铃声列表
	RingList *string `json:"RingList,omitnil" name:"RingList"`

	// 字体列表
	FontList *string `json:"FontList,omitnil" name:"FontList"`

	// 传感器列表
	SensorList *string `json:"SensorList,omitnil" name:"SensorList"`

	// CPU型号
	CpuType *string `json:"CpuType,omitnil" name:"CpuType"`

	// 电池容量
	Battery *string `json:"Battery,omitnil" name:"Battery"`

	// 信通院广告ID
	Oaid *string `json:"Oaid,omitnil" name:"Oaid"`

	// IOS 广告ID
	Idfa *string `json:"Idfa,omitnil" name:"Idfa"`

	// IOS 应用ID
	Idfv *string `json:"Idfv,omitnil" name:"Idfv"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// IOS手机型号
	IphoneModel *string `json:"IphoneModel,omitnil" name:"IphoneModel"`

	// Android 指纹
	Fingerprint *string `json:"Fingerprint,omitnil" name:"Fingerprint"`

	// Android序列号
	SerialId *string `json:"SerialId,omitnil" name:"SerialId"`
}

type QueryDevAndRiskRequest struct {
	*tchttp.BaseRequest
	
	// 设备类型 0表示Android， 1表示IOS
	DevType *int64 `json:"DevType,omitnil" name:"DevType"`

	// Android Imei号
	Imei *string `json:"Imei,omitnil" name:"Imei"`

	// Mac地址
	Mac *string `json:"Mac,omitnil" name:"Mac"`

	// android  Aid
	Aid *string `json:"Aid,omitnil" name:"Aid"`

	// Android Cid
	Cid *string `json:"Cid,omitnil" name:"Cid"`

	// 手机Imsi
	Imsi *string `json:"Imsi,omitnil" name:"Imsi"`

	// Df 磁盘分区信息
	Df *string `json:"Df,omitnil" name:"Df"`

	// 内核版本
	KernelVer *string `json:"KernelVer,omitnil" name:"KernelVer"`

	// 存储大小
	Storage *string `json:"Storage,omitnil" name:"Storage"`

	// 设备驱动指纹
	Dfp *string `json:"Dfp,omitnil" name:"Dfp"`

	// 启动时间
	BootTime *string `json:"BootTime,omitnil" name:"BootTime"`

	// 分辨率 水平*垂直 格式
	Resolution *string `json:"Resolution,omitnil" name:"Resolution"`

	// 铃声列表
	RingList *string `json:"RingList,omitnil" name:"RingList"`

	// 字体列表
	FontList *string `json:"FontList,omitnil" name:"FontList"`

	// 传感器列表
	SensorList *string `json:"SensorList,omitnil" name:"SensorList"`

	// CPU型号
	CpuType *string `json:"CpuType,omitnil" name:"CpuType"`

	// 电池容量
	Battery *string `json:"Battery,omitnil" name:"Battery"`

	// 信通院广告ID
	Oaid *string `json:"Oaid,omitnil" name:"Oaid"`

	// IOS 广告ID
	Idfa *string `json:"Idfa,omitnil" name:"Idfa"`

	// IOS 应用ID
	Idfv *string `json:"Idfv,omitnil" name:"Idfv"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// IOS手机型号
	IphoneModel *string `json:"IphoneModel,omitnil" name:"IphoneModel"`

	// Android 指纹
	Fingerprint *string `json:"Fingerprint,omitnil" name:"Fingerprint"`

	// Android序列号
	SerialId *string `json:"SerialId,omitnil" name:"SerialId"`
}

func (r *QueryDevAndRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDevAndRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DevType")
	delete(f, "Imei")
	delete(f, "Mac")
	delete(f, "Aid")
	delete(f, "Cid")
	delete(f, "Imsi")
	delete(f, "Df")
	delete(f, "KernelVer")
	delete(f, "Storage")
	delete(f, "Dfp")
	delete(f, "BootTime")
	delete(f, "Resolution")
	delete(f, "RingList")
	delete(f, "FontList")
	delete(f, "SensorList")
	delete(f, "CpuType")
	delete(f, "Battery")
	delete(f, "Oaid")
	delete(f, "Idfa")
	delete(f, "Idfv")
	delete(f, "DeviceName")
	delete(f, "IphoneModel")
	delete(f, "Fingerprint")
	delete(f, "SerialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDevAndRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDevAndRiskResponseParams struct {
	// 是否查得
	Found *int64 `json:"Found,omitnil" name:"Found"`

	// 匹配数量级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllCnt *int64 `json:"AllCnt,omitnil" name:"AllCnt"`

	// 匹配到的设备信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Matches []*DevInfoQ `json:"Matches,omitnil" name:"Matches"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryDevAndRiskResponse struct {
	*tchttp.BaseResponse
	Response *QueryDevAndRiskResponseParams `json:"Response"`
}

func (r *QueryDevAndRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDevAndRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskDetail struct {
	// 风险码
	RiskCode *int64 `json:"RiskCode,omitnil" name:"RiskCode"`

	// 风险详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCodeValue *string `json:"RiskCodeValue,omitnil" name:"RiskCodeValue"`
}

type RiskInfo struct {
	// 风险码
	Key *int64 `json:"Key,omitnil" name:"Key"`

	// 风险详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}