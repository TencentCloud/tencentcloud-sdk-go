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

package v20200210

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Device struct {
	// 业务入参id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 业务入参类型
	DeviceType *int64 `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`
}

type InputBusinessEncryptData struct {
	// 加密方式；0：AES;1:DES
	EncryptMethod *uint64 `json:"EncryptMethod,omitnil,omitempty" name:"EncryptMethod"`

	// 业务数据加密字符串
	EncryptData *string `json:"EncryptData,omitnil,omitempty" name:"EncryptData"`

	// 加密模式；0：ECB,1:CBC;2:CTR;3:CFB;4:OFB
	EncryptMode *uint64 `json:"EncryptMode,omitnil,omitempty" name:"EncryptMode"`

	// 填充模式;0:ZERO ;1:PKCS5;3:PKCS7
	PaddingType *uint64 `json:"PaddingType,omitnil,omitempty" name:"PaddingType"`
}

type InputRecognizeTargetAudience struct {
	// 模型ID列表
	ModelIdList []*int64 `json:"ModelIdList,omitnil,omitempty" name:"ModelIdList"`

	// 设备ID，AccountType指定的类型
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 设备号类型，1.imei 2.imeiMd5（小写后转MD5转小写）3.idfa， 4.idfaMd5（大写后转MD5转小写），5.手机号,256.其它
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 用户IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 操作系统类型(unknown，android，ios，windows)
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 操作系统版本
	Osv *string `json:"Osv,omitnil,omitempty" name:"Osv"`

	// 纬度
	Lat *string `json:"Lat,omitnil,omitempty" name:"Lat"`

	// 经度
	Lon *string `json:"Lon,omitnil,omitempty" name:"Lon"`

	// 设备型号(MI 6)
	DeviceModel *string `json:"DeviceModel,omitnil,omitempty" name:"DeviceModel"`

	// 竞价底价
	BidFloor *int64 `json:"BidFloor,omitnil,omitempty" name:"BidFloor"`

	// 年龄
	Age *int64 `json:"Age,omitnil,omitempty" name:"Age"`

	// 性别(1.MALE 2.FEMALE)
	Gender *int64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 用户地址
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 投放模式（0=PDB，1=PD，2=RTB，10=其他）
	DeliveryMode *int64 `json:"DeliveryMode,omitnil,omitempty" name:"DeliveryMode"`

	// 广告位类型<br />（0=前贴片，1=开屏广告，2=网页头部广告、3=网页中部广告、4=网页底部广告、5=悬浮广告、10=其它）
	AdvertisingType *int64 `json:"AdvertisingType,omitnil,omitempty" name:"AdvertisingType"`

	// mac地址，建议提供
	Mac *string `json:"Mac,omitnil,omitempty" name:"Mac"`

	// 电话号码
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 浏览器类型
	Ua *string `json:"Ua,omitnil,omitempty" name:"Ua"`

	// 客户端应用
	App *string `json:"App,omitnil,omitempty" name:"App"`

	// 应用包名
	Package *string `json:"Package,omitnil,omitempty" name:"Package"`

	// 设备制造商
	Maker *string `json:"Maker,omitnil,omitempty" name:"Maker"`

	// 设备类型（PHONE,TABLET）
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 入网方式(wifi,4g,3g,2g)
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 运营商(1.移动 2.联通 3.电信等)
	Sp *int64 `json:"Sp,omitnil,omitempty" name:"Sp"`

	// 设备屏幕分辨率宽度像素数
	DeviceW *int64 `json:"DeviceW,omitnil,omitempty" name:"DeviceW"`

	// 设备屏幕分辨率高度像素数
	DeviceH *int64 `json:"DeviceH,omitnil,omitempty" name:"DeviceH"`

	// 是否全屏插广告(0-否，1-是)
	FullScreen *int64 `json:"FullScreen,omitnil,omitempty" name:"FullScreen"`

	// 广告位宽度
	ImpBannerW *int64 `json:"ImpBannerW,omitnil,omitempty" name:"ImpBannerW"`

	// 广告位高度
	ImpBannerH *int64 `json:"ImpBannerH,omitnil,omitempty" name:"ImpBannerH"`

	// 网址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 上下文信息
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 渠道
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// 请求ID
	ReqId *string `json:"ReqId,omitnil,omitempty" name:"ReqId"`

	// 请求ID的md5值
	ReqMd5 *string `json:"ReqMd5,omitnil,omitempty" name:"ReqMd5"`

	// ad_type
	AdType *int64 `json:"AdType,omitnil,omitempty" name:"AdType"`

	// app名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// app版本描述
	AppVer *string `json:"AppVer,omitnil,omitempty" name:"AppVer"`

	// 竞价模式1：rtb 2:pd
	ReqType *int64 `json:"ReqType,omitnil,omitempty" name:"ReqType"`

	// 用户是否授权,1为授权，0为未授权
	IsAuthorized *uint64 `json:"IsAuthorized,omitnil,omitempty" name:"IsAuthorized"`

	// 设备信息
	DeviceList []*Device `json:"DeviceList,omitnil,omitempty" name:"DeviceList"`
}

type ManagePortraitRiskInput struct {
	// 请求时间戳秒
	PostTime *int64 `json:"PostTime,omitnil,omitempty" name:"PostTime"`

	// 用户公网ip（仅支持IPv4）
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 渠道号
	Channel *int64 `json:"Channel,omitnil,omitempty" name:"Channel"`
}

type ManagePortraitRiskOutput struct {
	// 返回码（0，成功，其他失败）
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 返回码对应的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *ManagePortraitRiskValueOutput `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ManagePortraitRiskRequestParams struct {
	// 业务入参
	BusinessSecurityData *ManagePortraitRiskInput `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

type ManagePortraitRiskRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *ManagePortraitRiskInput `json:"BusinessSecurityData,omitnil,omitempty" name:"BusinessSecurityData"`
}

func (r *ManagePortraitRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManagePortraitRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManagePortraitRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManagePortraitRiskResponseParams struct {
	// 业务出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ManagePortraitRiskOutput `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManagePortraitRiskResponse struct {
	*tchttp.BaseResponse
	Response *ManagePortraitRiskResponseParams `json:"Response"`
}

func (r *ManagePortraitRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManagePortraitRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManagePortraitRiskValueOutput struct {
	// 对应的IP
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// 返回风险等级, 0 - 4，0代表无风险，数值越大，风险越高
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`
}

type OutputRecognizeTargetAudience struct {
	// 返回码（0，成功，其他失败）
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 返回码对应的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 返回模型结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputRecognizeTargetAudienceValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputRecognizeTargetAudienceValue struct {
	// 模型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *uint64 `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 是否正常返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFound *int64 `json:"IsFound,omitnil,omitempty" name:"IsFound"`

	// 返回分值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 模型类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelType *uint64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 入参Uid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`
}

// Predefined struct for user
type RecognizeCustomizedAudienceRequestParams struct {
	// 业务入参
	BspData *InputRecognizeTargetAudience `json:"BspData,omitnil,omitempty" name:"BspData"`
}

type RecognizeCustomizedAudienceRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BspData *InputRecognizeTargetAudience `json:"BspData,omitnil,omitempty" name:"BspData"`
}

func (r *RecognizeCustomizedAudienceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeCustomizedAudienceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BspData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeCustomizedAudienceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeCustomizedAudienceResponseParams struct {
	// 业务出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OutputRecognizeTargetAudience `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeCustomizedAudienceResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeCustomizedAudienceResponseParams `json:"Response"`
}

func (r *RecognizeCustomizedAudienceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeCustomizedAudienceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePreciseTargetAudienceRequestParams struct {
	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitnil,omitempty" name:"BspData"`
}

type RecognizePreciseTargetAudienceRequest struct {
	*tchttp.BaseRequest
	
	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitnil,omitempty" name:"BspData"`
}

func (r *RecognizePreciseTargetAudienceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePreciseTargetAudienceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BspData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePreciseTargetAudienceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePreciseTargetAudienceResponseParams struct {
	// 回包数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OutputRecognizeTargetAudience `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizePreciseTargetAudienceResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePreciseTargetAudienceResponseParams `json:"Response"`
}

func (r *RecognizePreciseTargetAudienceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePreciseTargetAudienceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTargetAudienceRequestParams struct {
	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitnil,omitempty" name:"BspData"`

	// 业务加密数据
	BusinessEncryptData *InputBusinessEncryptData `json:"BusinessEncryptData,omitnil,omitempty" name:"BusinessEncryptData"`
}

type RecognizeTargetAudienceRequest struct {
	*tchttp.BaseRequest
	
	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitnil,omitempty" name:"BspData"`

	// 业务加密数据
	BusinessEncryptData *InputBusinessEncryptData `json:"BusinessEncryptData,omitnil,omitempty" name:"BusinessEncryptData"`
}

func (r *RecognizeTargetAudienceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTargetAudienceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BspData")
	delete(f, "BusinessEncryptData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeTargetAudienceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTargetAudienceResponseParams struct {
	// 回包数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OutputRecognizeTargetAudience `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecognizeTargetAudienceResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeTargetAudienceResponseParams `json:"Response"`
}

func (r *RecognizeTargetAudienceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTargetAudienceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}