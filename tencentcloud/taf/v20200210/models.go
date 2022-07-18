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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DetectFraudKOLRequestParams struct {
	// 业务数据
	BspData *InputKolBspData `json:"BspData,omitempty" name:"BspData"`

	// 业务加密数据
	BusinessEncryptData *InputBusinessEncryptData `json:"BusinessEncryptData,omitempty" name:"BusinessEncryptData"`
}

type DetectFraudKOLRequest struct {
	*tchttp.BaseRequest
	
	// 业务数据
	BspData *InputKolBspData `json:"BspData,omitempty" name:"BspData"`

	// 业务加密数据
	BusinessEncryptData *InputBusinessEncryptData `json:"BusinessEncryptData,omitempty" name:"BusinessEncryptData"`
}

func (r *DetectFraudKOLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFraudKOLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BspData")
	delete(f, "BusinessEncryptData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectFraudKOLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectFraudKOLResponseParams struct {
	// 回包数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OutputKolData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DetectFraudKOLResponse struct {
	*tchttp.BaseResponse
	Response *DetectFraudKOLResponseParams `json:"Response"`
}

func (r *DetectFraudKOLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectFraudKOLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Device struct {
	// 业务入参id
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 业务入参类型
	DeviceType *int64 `json:"DeviceType,omitempty" name:"DeviceType"`
}

type InputBusinessEncryptData struct {

}

type InputKolBspData struct {
	// BspData
	DataList []*InputKolDataList `json:"DataList,omitempty" name:"DataList"`
}

type InputKolDataList struct {
	// 账号类型[1：微信；2：qq；3：微博]
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// KOL账号ID[比如微信公众号ID]
	Id *string `json:"Id,omitempty" name:"Id"`

	// KOL名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 代理商名称
	AgentInfo *string `json:"AgentInfo,omitempty" name:"AgentInfo"`

	// 是否授权
	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`
}

type InputRecognizeEffectiveFlow struct {

}

type InputRecognizeTargetAudience struct {
	// 模型ID列表
	ModelIdList []*int64 `json:"ModelIdList,omitempty" name:"ModelIdList"`

	// 设备ID，AccountType指定的类型
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 设备号类型，1.imei 2.imeiMd5（小写后转MD5转小写）3.idfa， 4.idfaMd5（大写后转MD5转小写），5.手机号,256.其它
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 用户IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 操作系统类型(unknown，android，ios，windows)
	Os *string `json:"Os,omitempty" name:"Os"`

	// 操作系统版本
	Osv *string `json:"Osv,omitempty" name:"Osv"`

	// 纬度
	Lat *string `json:"Lat,omitempty" name:"Lat"`

	// 经度
	Lon *string `json:"Lon,omitempty" name:"Lon"`

	// 设备型号(MI 6)
	DeviceModel *string `json:"DeviceModel,omitempty" name:"DeviceModel"`

	// 竞价底价
	BidFloor *int64 `json:"BidFloor,omitempty" name:"BidFloor"`

	// 年龄
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 性别(1.MALE 2.FEMALE)
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 用户地址
	Location *string `json:"Location,omitempty" name:"Location"`

	// 投放模式（0=PDB，1=PD，2=RTB，10=其他）
	DeliveryMode *int64 `json:"DeliveryMode,omitempty" name:"DeliveryMode"`

	// 广告位类型<br />（0=前贴片，1=开屏广告，2=网页头部广告、3=网页中部广告、4=网页底部广告、5=悬浮广告、10=其它）
	AdvertisingType *int64 `json:"AdvertisingType,omitempty" name:"AdvertisingType"`

	// mac地址，建议提供
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// 电话号码
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 浏览器类型
	Ua *string `json:"Ua,omitempty" name:"Ua"`

	// 客户端应用
	App *string `json:"App,omitempty" name:"App"`

	// 应用包名
	Package *string `json:"Package,omitempty" name:"Package"`

	// 设备制造商
	Maker *string `json:"Maker,omitempty" name:"Maker"`

	// 设备类型（PHONE,TABLET）
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`

	// 入网方式(wifi,4g,3g,2g)
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`

	// 运营商(1.移动 2.联通 3.电信等)
	Sp *int64 `json:"Sp,omitempty" name:"Sp"`

	// 设备屏幕分辨率宽度像素数
	DeviceW *int64 `json:"DeviceW,omitempty" name:"DeviceW"`

	// 设备屏幕分辨率高度像素数
	DeviceH *int64 `json:"DeviceH,omitempty" name:"DeviceH"`

	// 是否全屏插广告(0-否，1-是)
	FullScreen *int64 `json:"FullScreen,omitempty" name:"FullScreen"`

	// 广告位宽度
	ImpBannerW *int64 `json:"ImpBannerW,omitempty" name:"ImpBannerW"`

	// 广告位高度
	ImpBannerH *int64 `json:"ImpBannerH,omitempty" name:"ImpBannerH"`

	// 网址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 上下文信息
	Context *string `json:"Context,omitempty" name:"Context"`

	// 渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 请求ID
	ReqId *string `json:"ReqId,omitempty" name:"ReqId"`

	// 请求ID的md5值
	ReqMd5 *string `json:"ReqMd5,omitempty" name:"ReqMd5"`

	// ad_type
	AdType *int64 `json:"AdType,omitempty" name:"AdType"`

	// app名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// app版本描述
	AppVer *string `json:"AppVer,omitempty" name:"AppVer"`

	// 竞价模式1：rtb 2:pd
	ReqType *int64 `json:"ReqType,omitempty" name:"ReqType"`

	// 用户是否授权,1为授权，0为未授权
	IsAuthorized *uint64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`

	// 设备信息
	DeviceList []*Device `json:"DeviceList,omitempty" name:"DeviceList"`
}

type InputSendTrafficSecuritySmsMsg struct {
	// 投放任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 手机号码列表（号码量<=200）
	Mobiles []*string `json:"Mobiles,omitempty" name:"Mobiles"`

	// 是否授权，1：已授权
	IsAuthorized *int64 `json:"IsAuthorized,omitempty" name:"IsAuthorized"`

	// 加密方式，0：AES加密；1：DES加密
	EncryptMethod *int64 `json:"EncryptMethod,omitempty" name:"EncryptMethod"`

	// 加密算法中的块处理模式，0：ECB模式；1：CBC模式；2：CTR模式；3：CFB模式；4：OFB模式；
	EncryptMode *int64 `json:"EncryptMode,omitempty" name:"EncryptMode"`

	// 填充模式，0：ZeroPadding；1：PKCS5Padding；2：PKCS7Padding；
	PaddingType *int64 `json:"PaddingType,omitempty" name:"PaddingType"`

	// 加密数据
	EncryptData *string `json:"EncryptData,omitempty" name:"EncryptData"`
}

type OutputKolData struct {
	// 错误码[0:成功；非0：失败的错误码]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 业务返回数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputKolValue `json:"Value,omitempty" name:"Value"`
}

type OutputKolValue struct {
	// KOL账号ID[比如微信公众号ID]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 是否查得[0：未查得；1：查得]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCheck *int64 `json:"IsCheck,omitempty" name:"IsCheck"`

	// 作弊的可能性[0～100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	FraudPScore *int64 `json:"FraudPScore,omitempty" name:"FraudPScore"`

	// 作弊的严重性[0～100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvilPScore *int64 `json:"EvilPScore,omitempty" name:"EvilPScore"`
}

type OutputRecognizeEffectiveFlow struct {
	// 返回码。0表示成功，非0标识失败错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// UTF-8编码，出错消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 业务入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputRecognizeEffectiveFlowValue `json:"Value,omitempty" name:"Value"`
}

type OutputRecognizeEffectiveFlowValue struct {
	// 返回标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lable *string `json:"Lable,omitempty" name:"Lable"`

	// 返回分值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

type OutputRecognizeTargetAudience struct {
	// 返回码（0，成功，其他失败）
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 返回码对应的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 返回模型结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputRecognizeTargetAudienceValue `json:"Value,omitempty" name:"Value"`
}

type OutputRecognizeTargetAudienceValue struct {
	// 模型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *uint64 `json:"ModelId,omitempty" name:"ModelId"`

	// 是否正常返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFound *int64 `json:"IsFound,omitempty" name:"IsFound"`

	// 返回分值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

type OutputSendTrafficSecuritySmsMsg struct {
	// 返回码（0：接口调用成功 非0：接口调用失败）
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 返回码对应的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 发送失败的号码列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type RecognizeCustomizedAudienceRequestParams struct {
	// 业务入参
	BspData *InputRecognizeTargetAudience `json:"BspData,omitempty" name:"BspData"`
}

type RecognizeCustomizedAudienceRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BspData *InputRecognizeTargetAudience `json:"BspData,omitempty" name:"BspData"`
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
	Data *OutputRecognizeTargetAudience `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type RecognizeEffectiveFlowRequestParams struct {
	// 业务入参
	BusinessSecurityData *InputRecognizeEffectiveFlow `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`
}

type RecognizeEffectiveFlowRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BusinessSecurityData *InputRecognizeEffectiveFlow `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`
}

func (r *RecognizeEffectiveFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeEffectiveFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeEffectiveFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeEffectiveFlowResponseParams struct {
	// 业务出参
	Data *OutputRecognizeEffectiveFlow `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecognizeEffectiveFlowResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeEffectiveFlowResponseParams `json:"Response"`
}

func (r *RecognizeEffectiveFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeEffectiveFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePreciseTargetAudienceRequestParams struct {
	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitempty" name:"BspData"`
}

type RecognizePreciseTargetAudienceRequest struct {
	*tchttp.BaseRequest
	
	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitempty" name:"BspData"`
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
	Data *OutputRecognizeTargetAudience `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BspData *InputRecognizeTargetAudience `json:"BspData,omitempty" name:"BspData"`

	// 业务加密数据
	BusinessEncryptData *InputBusinessEncryptData `json:"BusinessEncryptData,omitempty" name:"BusinessEncryptData"`
}

type RecognizeTargetAudienceRequest struct {
	*tchttp.BaseRequest
	
	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitempty" name:"BspData"`

	// 业务加密数据
	BusinessEncryptData *InputBusinessEncryptData `json:"BusinessEncryptData,omitempty" name:"BusinessEncryptData"`
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
	Data *OutputRecognizeTargetAudience `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type SendTrafficSecuritySmsMessageRequestParams struct {
	// 业务入参
	BspData *InputSendTrafficSecuritySmsMsg `json:"BspData,omitempty" name:"BspData"`
}

type SendTrafficSecuritySmsMessageRequest struct {
	*tchttp.BaseRequest
	
	// 业务入参
	BspData *InputSendTrafficSecuritySmsMsg `json:"BspData,omitempty" name:"BspData"`
}

func (r *SendTrafficSecuritySmsMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendTrafficSecuritySmsMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BspData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendTrafficSecuritySmsMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendTrafficSecuritySmsMessageResponseParams struct {
	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OutputSendTrafficSecuritySmsMsg `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendTrafficSecuritySmsMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendTrafficSecuritySmsMessageResponseParams `json:"Response"`
}

func (r *SendTrafficSecuritySmsMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendTrafficSecuritySmsMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}