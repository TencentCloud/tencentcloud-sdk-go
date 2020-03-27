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

package v20180301

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BankCard2EVerificationRequest struct {
	*tchttp.BaseRequest

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`
}

func (r *BankCard2EVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCard2EVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCard2EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码
	// 计费结果码：
	//   '0': '认证通过'
	//   '-1': '认证未通过'
	//  '-4': '持卡人信息有误'
	//   '-5': '未开通无卡支付'
	//   '-6': '此卡被没收'
	//   '-7': '无效卡号'
	//   '-8': '此卡无对应发卡行'
	//   '-9': '该卡未初始化或睡眠卡'
	//   '-10': '作弊卡、吞卡'
	//   '-11': '此卡已挂失'
	//   '-12': '该卡已过期'
	//   '-13': '受限制的卡'
	//   '-14': '密码错误次数超限'
	//   '-15': '发卡行不支持此交易'
	// 不计费结果码：
	//   '-2': '姓名校验不通过'
	//   '-3': '银行卡号码有误'
	//   '-16': '验证中心服务繁忙'
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BankCard2EVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCard2EVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCard4EVerificationRequest struct {
	*tchttp.BaseRequest

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 手机号码
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认：0 身份证，其他证件类型需求可以联系小助手faceid001确认。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`
}

func (r *BankCard4EVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCard4EVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCard4EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码
	// 收费结果码：
	// '0': '认证通过'
	// '-1': '认证未通过'
	// '-6': '持卡人信息有误'
	// '-7': '未开通无卡支付'
	// '-8': '此卡被没收'
	// '-9': '无效卡号'
	// '-10': '此卡无对应发卡行'
	// '-11': '该卡未初始化或睡眠卡'
	// '-12': '作弊卡、吞卡'
	// '-13': '此卡已挂失'
	// '-14': '该卡已过期'
	// '-15': '受限制的卡'
	// '-16': '密码错误次数超限'
	// '-17': '发卡行不支持此交易'
	// 不收费结果码：
	// '-2': '姓名校验不通过'
	// '-3': '身份证号码有误'
	// '-4': '银行卡号码有误'
	// '-5': '手机号码不合法'
	// '-18': '验证中心服务繁忙'
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BankCard4EVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCard4EVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCardVerificationRequest struct {
	*tchttp.BaseRequest

	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认：0 身份证，其他证件类型需求可以联系小助手faceid001确认。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`
}

func (r *BankCardVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCardVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码
	// 收费结果码：
	// '0': '认证通过'
	// '-1': '认证未通过'
	// '-5': '持卡人信息有误'
	// '-6': '未开通无卡支付'
	// '-7': '此卡被没收'
	// '-8': '无效卡号'
	// '-9': '此卡无对应发卡行'
	// '-10': '该卡未初始化或睡眠卡'
	// '-11': '作弊卡、吞卡'
	// '-12': '此卡已挂失'
	// '-13': '该卡已过期'
	// '-14': '受限制的卡'
	// '-15': '密码错误次数超限'
	// '-16': '发卡行不支持此交易'
	// 不收费结果码：
	// '-2': '姓名校验不通过'
	// '-3': '身份证号码有误'
	// '-4': '银行卡号码有误'
	// '-17': '验证中心服务繁忙'
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BankCardVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCardVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectAuthRequest struct {
	*tchttp.BaseRequest

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请加慧眼小助手微信（faceid001）进行咨询。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 本接口不需要传递此参数。
	TerminalType *string `json:"TerminalType,omitempty" name:"TerminalType"`

	// 身份标识（未使用OCR服务时，必须传入）。
	// 规则：a-zA-Z0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。（未使用OCR服务时，必须传入）最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 认证结束后重定向的回调链接地址。最长长度1024位。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// 透传字段，在获取验证结果时返回。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 用于人脸比对的照片，图片的BASE64值；
	// BASE64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectAuthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectAuthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用于发起核身流程的URL，仅微信H5场景使用。
		Url *string `json:"Url,omitempty" name:"Url"`

		// 一次核身流程的标识，有效时间为7,200秒；
	// 完成核身后，可用该标识获取验证结果信息。
		BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectAuthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectDetail struct {

	// 请求时间戳。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqTime *string `json:"ReqTime,omitempty" name:"ReqTime"`

	// 本次活体一比一请求的唯一标记。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seq *string `json:"Seq,omitempty" name:"Seq"`

	// 参与本次活体一比一的身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Idcard *string `json:"Idcard,omitempty" name:"Idcard"`

	// 参与本次活体一比一的姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 本次活体一比一的相似度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sim *string `json:"Sim,omitempty" name:"Sim"`

	// 本次活体一比一是否收费
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNeedCharge *bool `json:"IsNeedCharge,omitempty" name:"IsNeedCharge"`

	// 本次活体一比一最终结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errcode *int64 `json:"Errcode,omitempty" name:"Errcode"`

	// 本次活体一比一最终结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errmsg *string `json:"Errmsg,omitempty" name:"Errmsg"`

	// 本次活体结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Livestatus *int64 `json:"Livestatus,omitempty" name:"Livestatus"`

	// 本次活体结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Livemsg *string `json:"Livemsg,omitempty" name:"Livemsg"`

	// 本次一比一结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparestatus *int64 `json:"Comparestatus,omitempty" name:"Comparestatus"`

	// 本次一比一结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparemsg *string `json:"Comparemsg,omitempty" name:"Comparemsg"`
}

type DetectInfoBestFrame struct {

	// 活体比对最佳帧。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrame *string `json:"BestFrame,omitempty" name:"BestFrame"`

	// 自截帧。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrames []*string `json:"BestFrames,omitempty" name:"BestFrames" list`
}

type DetectInfoIdCardData struct {

	// OCR正面照片的base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFront *string `json:"OcrFront,omitempty" name:"OcrFront"`

	// OCR反面照片的base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrBack *string `json:"OcrBack,omitempty" name:"OcrBack"`

	// 旋转裁边后的正面照片base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedFrontImage *string `json:"ProcessedFrontImage,omitempty" name:"ProcessedFrontImage"`

	// 旋转裁边后的背面照片base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedBackImage *string `json:"ProcessedBackImage,omitempty" name:"ProcessedBackImage"`

	// 身份证正面人像图base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`
}

type DetectInfoText struct {

	// 本次流程最终验证结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 本次流程最终验证结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 本次验证使用的身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 本次验证使用的姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Ocr识别结果。民族。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrNation *string `json:"OcrNation,omitempty" name:"OcrNation"`

	// Ocr识别结果。家庭住址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrAddress *string `json:"OcrAddress,omitempty" name:"OcrAddress"`

	// Ocr识别结果。生日。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrBirth *string `json:"OcrBirth,omitempty" name:"OcrBirth"`

	// Ocr识别结果。签发机关。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrAuthority *string `json:"OcrAuthority,omitempty" name:"OcrAuthority"`

	// Ocr识别结果。有效日期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrValidDate *string `json:"OcrValidDate,omitempty" name:"OcrValidDate"`

	// Ocr识别结果。姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrName *string `json:"OcrName,omitempty" name:"OcrName"`

	// Ocr识别结果。身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrIdCard *string `json:"OcrIdCard,omitempty" name:"OcrIdCard"`

	// Ocr识别结果。性别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrGender *string `json:"OcrGender,omitempty" name:"OcrGender"`

	// 本次流程最终活体结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveStatus *int64 `json:"LiveStatus,omitempty" name:"LiveStatus"`

	// 本次流程最终活体结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveMsg *string `json:"LiveMsg,omitempty" name:"LiveMsg"`

	// 本次流程最终一比一结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparestatus *int64 `json:"Comparestatus,omitempty" name:"Comparestatus"`

	// 本次流程最终一比一结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparemsg *string `json:"Comparemsg,omitempty" name:"Comparemsg"`

	// 本次流程活体一比一的分数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sim *string `json:"Sim,omitempty" name:"Sim"`

	// 地理位置经纬度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// Auth接口带入额外信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 本次流程进行的活体一比一流水。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessDetail []*DetectDetail `json:"LivenessDetail,omitempty" name:"LivenessDetail" list`

	// 手机号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

type DetectInfoVideoData struct {

	// 活体视频的base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessVideo *string `json:"LivenessVideo,omitempty" name:"LivenessVideo"`
}

type GetActionSequenceRequest struct {
	*tchttp.BaseRequest
}

func (r *GetActionSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetActionSequenceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetActionSequenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 动作顺序(2,1 or 1,2) 。1代表张嘴，2代表闭眼。
		ActionSequence *string `json:"ActionSequence,omitempty" name:"ActionSequence"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetActionSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetActionSequenceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDetectInfoEnhancedRequest struct {
	*tchttp.BaseRequest

	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，由腾讯侧在线下对接时分配。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证信息；3：视频最佳截图信息；4：视频信息）。
	// 如 134表示拉取文本类、视频最佳截图信息、视频信息。
	// 默认值：0
	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`

	// 从活体视频中截取一定张数的最佳帧。默认为0，最大为10，超出10的最多只给10张。（InfoType需要包含3）
	BestFramesCount *uint64 `json:"BestFramesCount,omitempty" name:"BestFramesCount"`

	// 是否对身份证照片进行裁边。默认为false。（InfoType需要包含2）
	IsCutIdCardImage *bool `json:"IsCutIdCardImage,omitempty" name:"IsCutIdCardImage"`

	// 是否需要从身份证中抠出头像。默认为false。（InfoType需要包含2）
	IsNeedIdCardAvatar *bool `json:"IsNeedIdCardAvatar,omitempty" name:"IsNeedIdCardAvatar"`
}

func (r *GetDetectInfoEnhancedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDetectInfoEnhancedRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDetectInfoEnhancedResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文本类信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Text *DetectInfoText `json:"Text,omitempty" name:"Text"`

		// 身份证照片信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IdCardData *DetectInfoIdCardData `json:"IdCardData,omitempty" name:"IdCardData"`

		// 最佳帧信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrame *DetectInfoBestFrame `json:"BestFrame,omitempty" name:"BestFrame"`

		// 视频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		VideoData *DetectInfoVideoData `json:"VideoData,omitempty" name:"VideoData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDetectInfoEnhancedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDetectInfoEnhancedResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDetectInfoRequest struct {
	*tchttp.BaseRequest

	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请加慧眼小助手微信（faceid001）进行咨询。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证正反面；3：视频最佳截图照片；4：视频）。
	// 如 134表示拉取文本类、视频最佳截图照片、视频。
	// 默认值：0
	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`
}

func (r *GetDetectInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDetectInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDetectInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// JSON字符串。
	// {
	//   // 文本类信息
	//   "Text": {
	//     "ErrCode": null,      // 本次核身最终结果。0为成功
	//     "ErrMsg": null,       // 本次核身最终结果信息描述。
	//     "IdCard": "",         // 本次核身最终获得的身份证号。
	//     "Name": "",           // 本次核身最终获得的姓名。
	//     "OcrNation": null,    // ocr阶段获取的民族
	//     "OcrAddress": null,   // ocr阶段获取的地址
	//     "OcrBirth": null,     // ocr阶段获取的出生信息
	//     "OcrAuthority": null, // ocr阶段获取的证件签发机关
	//     "OcrValidDate": null, // ocr阶段获取的证件有效期
	//     "OcrName": null,      // ocr阶段获取的姓名
	//     "OcrIdCard": null,    // ocr阶段获取的身份证号
	//     "OcrGender": null,    // ocr阶段获取的性别
	//     "LiveStatus": null,   // 活体检测阶段的错误码。0为成功
	//     "LiveMsg": null,      // 活体检测阶段的错误信息
	//     "Comparestatus": null,// 一比一阶段的错误码。0为成功
	//     "Comparemsg": null,   // 一比一阶段的错误信息
	//     "Location": null, // 地理位置信息
	//     "Extra": "",          // DetectAuth结果传进来的Extra信息
	//     "Detail": {           // 活体一比一信息详情
	//       "LivenessData": []
	//     }
	//   },
	//   // 身份证正反面照片Base64
	//   "IdCardData": {
	//     "OcrFront": null,
	//     "OcrBack": null
	//   },
	//   // 视频最佳帧截图Base64
	//   "BestFrame": {
	//     "BestFrame": null
	//   },
	//   // 活体视频Base64
	//   "VideoData": {
	//     "LivenessVideo": null
	//   }
	// }
		DetectInfo *string `json:"DetectInfo,omitempty" name:"DetectInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDetectInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDetectInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLiveCodeRequest struct {
	*tchttp.BaseRequest
}

func (r *GetLiveCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLiveCodeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetLiveCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数字验证码，如：1234
		LiveCode *string `json:"LiveCode,omitempty" name:"LiveCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLiveCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLiveCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IdCardOCRVerificationRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	// 姓名和身份证号、ImageBase64、ImageUrl三者必须提供其中之一。若都提供了，则按照姓名和身份证号>ImageBase64>ImageUrl的优先级使用参数。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证人像面的 Base64 值
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 身份证人像面的 Url 地址
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *IdCardOCRVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IdCardOCRVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IdCardOCRVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 姓名和身份证号一致
	// -1: 姓名和身份证号不一致
	// 不收费结果码：
	// -2: 非法身份证号（长度、校验位等不正确）
	// -3: 非法姓名（长度、格式等不正确）
	// -4: 证件库服务异常
	// -5: 证件库中无此身份证记录
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 用于验证的姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 用于验证的身份证号
		IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

		// OCR得到的性别
	// 注意：此字段可能返回 null，表示取不到有效值。
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// OCR得到的民族
	// 注意：此字段可能返回 null，表示取不到有效值。
		Nation *string `json:"Nation,omitempty" name:"Nation"`

		// OCR得到的生日
	// 注意：此字段可能返回 null，表示取不到有效值。
		Birth *string `json:"Birth,omitempty" name:"Birth"`

		// OCR得到的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		Address *string `json:"Address,omitempty" name:"Address"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IdCardOCRVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IdCardOCRVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IdCardVerificationRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *IdCardVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IdCardVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IdCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 姓名和身份证号一致
	// -1: 姓名和身份证号不一致
	// 不收费结果码：
	// -2: 非法身份证号（长度、校验位等不正确）
	// -3: 非法姓名（长度、格式等不正确）
	// -4: 证件库服务异常
	// -5: 证件库中无此身份证记录
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IdCardVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IdCardVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageRecognitionRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用于人脸比对的照片，图片的BASE64值；
	// BASE64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitempty" name:"Optional"`
}

func (r *ImageRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LivenessCompareRequest struct {
	*tchttp.BaseRequest

	// 用于人脸比对的照片，图片的BASE64值；
	// BASE64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过5M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitempty" name:"Optional"`
}

func (r *LivenessCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessCompareRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LivenessCompareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）。
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LivenessCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessCompareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LivenessRecognitionRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过5M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitempty" name:"Optional"`
}

func (r *LivenessRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LivenessRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LivenessRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LivenessRequest struct {
	*tchttp.BaseRequest

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过5M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：不需要传递此参数。
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitempty" name:"Optional"`
}

func (r *LivenessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LivenessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LivenessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MinorsVerificationRequest struct {
	*tchttp.BaseRequest

	// 参与校验的参数类型。
	// 0：使用手机号进行校验；
	// 1：使用姓名与身份证号进行校验。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 手机号，11位数字，
	// 特别提示：
	// 手机号验证只限制在腾讯健康守护可信模型覆盖的数据范围内，与手机号本身在运营商是否实名无关联，不在范围会提示“手机号未实名”，建议客户与传入姓名和身份证号信息组合使用。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 身份证号码。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *MinorsVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MinorsVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MinorsVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果码，收费情况如下。
	// 收费结果码：
	// 0: 成年
	// -1: 未成年
	// -2: 手机号未实名
	// -3: 姓名和身份证号不一致
	// 
	// 不收费结果码：
	// -4: 非法身份证号（长度、校验位等不正确）
	// -5: 非法姓名（长度、格式等不正确）
	// -6: 数据源服务异常
	// -7: 数据源中无此身份证记录
	// -8: 公安比对系统升级中，请稍后再试
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 当结果码为0或者-1时，该字段的值为年龄区间。
	// 格式为[a,b)，表示年龄在a岁以上（包括a岁），b岁以下（不包括b岁）。若b为+时表示没有上限。
		AgeRange *string `json:"AgeRange,omitempty" name:"AgeRange"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MinorsVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MinorsVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MobileNetworkTimeVerificationRequest struct {
	*tchttp.BaseRequest

	// 手机号码。不支持电信手机号。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

func (r *MobileNetworkTimeVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MobileNetworkTimeVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MobileNetworkTimeVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 成功
	// -2: 手机号不存在
	// -3: 手机号存在，但无法查询到在网时长
	// 不收费结果码：
	// -1: 手机号格式不正确
	// -4: 验证中心服务繁忙
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 在网时长区间。
	// 格式为(a,b]，表示在网时长在a个月以上，b个月以下。若b为+时表示没有上限。
		Range *string `json:"Range,omitempty" name:"Range"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MobileNetworkTimeVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MobileNetworkTimeVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MobileStatusRequest struct {
	*tchttp.BaseRequest

	// 手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

func (r *MobileStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MobileStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MobileStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0：成功
	// 不收费结果码：
	// -1：未查询到结果
	// -2：手机号格式不正确
	// -3：验证中心服务繁忙
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 状态码：
	// 0：正常
	// 1：停机
	// 2：销号
	// 3：空号
	// 4：不在网
	// 99：未知状态
		StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MobileStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MobileStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PhoneVerificationRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`
}

func (r *PhoneVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PhoneVerificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PhoneVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码:
	// 0: 认证通过
	// -1: 手机号已实名，但是身份证和姓名均与实名信息不一致 
	// -2: 手机号已实名，手机号和证件号一致，姓名不一致
	// -3: 手机号已实名，手机号和姓名一致，身份证不一致
	// -4: 信息不一致
	// -5: 手机号未实名
	// -6: 手机号码不合法
	// -7: 身份证号码有误
	// -8: 姓名校验不通过
	// -9: 没有记录
	// -10: 认证未通过
	// -11: 验证中心服务繁忙
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PhoneVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PhoneVerificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
