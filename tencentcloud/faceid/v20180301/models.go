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

		// 认证结果码。
	//   '0': '认证通过'
	//   '-1': '认证未通过'
	//   '-2': '姓名校验不通过'
	//   '-3': '银行卡号码有误'
	//   '-4': '持卡人信息有误'
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
	//   '-16': '服务繁忙'
		Result *string `json:"Result,omitempty" name:"Result"`

		// 认证结果信息。
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

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。（不填默认0）
	// 0 身份证
	// 1 军官证
	// 2 护照
	// 3 港澳证
	// 4 台胞证
	// 5 警官证
	// 6 士兵证
	// 7 其它证件
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

		// 认证结果码。
	//   '0': '认证通过'
	//   '-1': '认证未通过'
	//   '-2': '姓名校验不通过'
	//   '-3': '身份证号码有误'
	//   '-4': '银行卡号码有误'
	//   '-5': '手机号码不合法'
	//   '-6': '持卡人信息有误'
	//   '-7': '未开通无卡支付'
	//   '-8': '此卡被没收'
	//   '-9': '无效卡号'
	//   '-10': '此卡无对应发卡行'
	//   '-11': '该卡未初始化或睡眠卡'
	//   '-12': '作弊卡、吞卡'
	//   '-13': '此卡已挂失'
	//   '-14': '该卡已过期'
	//   '-15': '受限制的卡'
	//   '-16': '密码错误次数超限'
	//   '-17': '发卡行不支持此交易'
	//   '-18': '服务繁忙'
		Result *string `json:"Result,omitempty" name:"Result"`

		// 认证结果信息。
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

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。（不填默认0）
	// 0 身份证
	// 1 军官证
	// 2 护照
	// 3 港澳证
	// 4 台胞证
	// 5 警官证
	// 6 士兵证
	// 7 其它证件
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

		// 认证结果码。
	// '0': '认证通过'
	// '-1': '认证未通过'
	// '-2': '姓名校验不通过'
	// '-3': '身份证号码有误'
	// '-4': '银行卡号码有误'
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
	// '-17': '服务繁忙'
		Result *string `json:"Result,omitempty" name:"Result"`

		// 认证结果信息。
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

	// 身份标识（与公安权威库比对时必须是身份证号）。
	// 规则：a-zA-Z0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。最长长度32位。中文请使用UTF-8编码。
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

type GetDetectInfoRequest struct {
	*tchttp.BaseRequest

	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请加慧眼小助手微信（faceid001）进行咨询。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证正反面；3：视频最佳截图照片；4：视频）。
	// 如 134表示拉取文本类、视频最佳截图照片、视频。
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

		// 认证结果信息。
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

		// 业务错误描述
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

		// 业务错误描述
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

		// 业务错误描述
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

		// 业务错误描述
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
