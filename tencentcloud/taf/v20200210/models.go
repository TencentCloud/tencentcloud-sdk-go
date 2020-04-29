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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DetectAccountActivityRequest struct {
	*tchttp.BaseRequest

	// 业务入参
	BusinessSecurityData *InputDetectAccountActivity `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`
}

func (r *DetectAccountActivityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectAccountActivityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectAccountActivityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回包数据
		Data *OutputDetectAccountActivity `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectAccountActivityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectAccountActivityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectFraudKOLRequest struct {
	*tchttp.BaseRequest

	// 业务数据
	BspData *InputKolBspData `json:"BspData,omitempty" name:"BspData"`
}

func (r *DetectFraudKOLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectFraudKOLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectFraudKOLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回包数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *OutputKolData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectFraudKOLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectFraudKOLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhanceTaDegreeRequest struct {
	*tchttp.BaseRequest

	// 业务数据
	BspData *InputTaBspData `json:"BspData,omitempty" name:"BspData"`
}

func (r *EnhanceTaDegreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnhanceTaDegreeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhanceTaDegreeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回包数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *OutputTaData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnhanceTaDegreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnhanceTaDegreeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InputDetectAccountActivity struct {

	// 用户ID值，如微信/QQ openid，或 手机号等（如15912345687）
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 用户账号类型 
	// 1：QQ开放帐号 
	// 2：微信开放账号 
	// 4：手机号 （暂仅支持国内手机号）
	// 10004： 手机号MD5
	AccountType *uint64 `json:"AccountType,omitempty" name:"AccountType"`

	// 用户真实外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 用户操作时间戳，单位秒（格林威治时间精确到秒，如1501590972）
	PostTime *uint64 `json:"PostTime,omitempty" name:"PostTime"`

	// accountType是QQ或微信开放账号时，该参数必填，表示QQ或微信分配给网站或应用的appId，用来唯一标识网站或应用
	AppIdU *string `json:"AppIdU,omitempty" name:"AppIdU"`

	// 昵称，UTF-8 编码
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// 手机号。若 accountType 选4（手机号）、或10004（手机号 MD5），则无需重复填写。否则填入对应的手机号（如15912345687）
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 用户邮箱地址（非系统自动生成）
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 用户 HTTP 请求中的 cookie 进行2次 hash 的值，只要保证相同 cookie 的 hash 值一致即可
	CookieHash *float64 `json:"CookieHash,omitempty" name:"CookieHash"`

	// 用户HTTP请求的 userAgent
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// 用户HTTP请求中的 x_forward_for
	XForwardedFor *string `json:"XForwardedFor,omitempty" name:"XForwardedFor"`

	// Mac地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机制造商ID，如果手机注册，请带上此信息
	VendorId *string `json:"VendorId,omitempty" name:"VendorId"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`
}

type InputKolBspData struct {

	// BspData
	DataList []*InputKolDataList `json:"DataList,omitempty" name:"DataList" list`
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
}

type InputRecognizeTargetAudience struct {

	// 设备ID，AccountType指定的类型
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 设备号类型，1.imei 2.imeiMd5（小写后转MD5转小写）3.idfa， 4.idfaMd5（大写后转MD5转小写），5.手机号,256.其它
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 模型ID列表
	ModelIdList []*int64 `json:"ModelIdList,omitempty" name:"ModelIdList" list`

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

	// app name
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// appVer
	AppVer *string `json:"AppVer,omitempty" name:"AppVer"`

	// 竞价模式1：rtb 2:pd
	ReqType *int64 `json:"ReqType,omitempty" name:"ReqType"`
}

type InputSendTrafficSecuritySmsMsg struct {

	// 投放任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 手机号码列表（号码量<=200）
	Mobiles []*string `json:"Mobiles,omitempty" name:"Mobiles" list`
}

type InputTaBspData struct {

	// 请求序列号
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 操作系统类型[0：未知；1：android；2：ios；3：windows]
	OsType *string `json:"OsType,omitempty" name:"OsType"`

	// 年龄下限
	AgeFloor *int64 `json:"AgeFloor,omitempty" name:"AgeFloor"`

	// 年龄上限
	AgeCeil *int64 `json:"AgeCeil,omitempty" name:"AgeCeil"`

	// 性别[1：男；2：女]
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 用户操作时间
	UserTime *int64 `json:"UserTime,omitempty" name:"UserTime"`

	// Imei [在(Imei|ImeiMd5|Idfa|IdfaMd5)里面4选1]
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// Imei小写后加密Md5 [在(Imei|ImeiMd5|Idfa|IdfaMd5)里面4选1]
	ImeiMd5 *string `json:"ImeiMd5,omitempty" name:"ImeiMd5"`

	// Idfa [在(Imei|ImeiMd5|Idfa|IdfaMd5)里面4选1]
	Idfa *string `json:"Idfa,omitempty" name:"Idfa"`

	// Idfa大写后加密Md5 [在(Imei|ImeiMd5|Idfa|IdfaMd5)里面4选1]
	IdfaMd5 *string `json:"IdfaMd5,omitempty" name:"IdfaMd5"`

	// 用户IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// MAC地址[建议提供]
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// 手机号码[中国大陆]
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 浏览器
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// APP名称
	App *string `json:"App,omitempty" name:"App"`

	// 应用安装包名称
	Package *string `json:"Package,omitempty" name:"Package"`

	// 设备制造商
	DeviceMaker *string `json:"DeviceMaker,omitempty" name:"DeviceMaker"`

	// 设备型号
	DeviceModule *string `json:"DeviceModule,omitempty" name:"DeviceModule"`

	// 入网方式[1：WIFI；2：4G；3：3G；4：2G；5：其它]
	AccessMode *string `json:"AccessMode,omitempty" name:"AccessMode"`

	// 运营商[1：移动；2：联通；3：电信；4：其它]
	Sp *string `json:"Sp,omitempty" name:"Sp"`

	// 网址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 用户地址
	Location *string `json:"Location,omitempty" name:"Location"`

	// 纬度
	Latitude *string `json:"Latitude,omitempty" name:"Latitude"`

	// 精度
	Longitude *string `json:"Longitude,omitempty" name:"Longitude"`

	// 辅助区分信息
	Context *string `json:"Context,omitempty" name:"Context"`
}

type OutputDetectAccountActivity struct {

	// 返回码（0，成功，其他失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 返回码对应的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 返回活跃度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputDetectAccountActivityValue `json:"Value,omitempty" name:"Value"`
}

type OutputDetectAccountActivityValue struct {

	// 用户 ID accountType 不同对应不同的用户 ID。如是 QQ 或微信用户则填入对应的 openId
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 操作时间戳，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostTime *int64 `json:"PostTime,omitempty" name:"PostTime"`

	// 用户操作的真实外网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 0：表示不活跃
	// 1 - 4：活跃等级由低到高
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 账号标签：
	// 3，无效账号，送检账号参数无法成功解析，请检查微信openid是否有误 ，QQopenid是否与appidU对应，手机号是否有误。
	// 4，黑名单，该账号在业务侧有过拉黑记录
	// 5，白名单，该账号在业务侧有过加白名单记录
	// 
	// 环境标签：
	// 205，非公网有效ip，传进来的IP地址为内网ip地址或者ip保留地址；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type []*int64 `json:"Type,omitempty" name:"Type" list`
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
	Value []*OutputKolValue `json:"Value,omitempty" name:"Value" list`
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

type OutputRecognizeTargetAudience struct {

	// 返回码（0，成功，其他失败）
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 返回码对应的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 返回模型结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*OutputRecognizeTargetAudienceValue `json:"Value,omitempty" name:"Value" list`
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
	Value []*string `json:"Value,omitempty" name:"Value" list`
}

type OutputTaData struct {

	// 错误码[0:成功；非0：失败的错误码]
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 结果数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *OutputTaValue `json:"Value,omitempty" name:"Value"`
}

type OutputTaValue struct {

	// 是否查得[0：未查得；1：查得]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCheck *int64 `json:"IsCheck,omitempty" name:"IsCheck"`

	// 是否符合[0：不符合；1：符合]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMatch *int64 `json:"IsMatch,omitempty" name:"IsMatch"`
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

func (r *RecognizeCustomizedAudienceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecognizeCustomizedAudienceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 业务出参
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *OutputRecognizeTargetAudience `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizeCustomizedAudienceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecognizeCustomizedAudienceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *RecognizePreciseTargetAudienceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecognizePreciseTargetAudienceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回包数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *OutputRecognizeTargetAudience `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizePreciseTargetAudienceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecognizePreciseTargetAudienceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecognizeTargetAudienceRequest struct {
	*tchttp.BaseRequest

	// 业务数据
	BspData *InputRecognizeTargetAudience `json:"BspData,omitempty" name:"BspData"`
}

func (r *RecognizeTargetAudienceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecognizeTargetAudienceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecognizeTargetAudienceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 回包数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *OutputRecognizeTargetAudience `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizeTargetAudienceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecognizeTargetAudienceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *SendTrafficSecuritySmsMessageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendTrafficSecuritySmsMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *OutputSendTrafficSecuritySmsMsg `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendTrafficSecuritySmsMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendTrafficSecuritySmsMessageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
