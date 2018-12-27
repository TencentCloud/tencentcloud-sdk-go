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

type DetectAuthRequest struct {
	*tchttp.BaseRequest

	// 用于细分客户使用场景，由腾讯侧在线下对接时分配。
	RuleId *string `json:"RuleId" name:"RuleId"`

	// 本接口不需要传递此参数。
	TerminalType *string `json:"TerminalType" name:"TerminalType"`

	// 身份标识（与公安权威库比对时必须是身份证号）。
	// 规则：a-zA-Z0-9组合。最长长度32位。
	IdCard *string `json:"IdCard" name:"IdCard"`

	// 姓名。最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name" name:"Name"`

	// 认证结束后重定向的回调链接地址。最长长度1024位。
	RedirectUrl *string `json:"RedirectUrl" name:"RedirectUrl"`

	// 透传字段，在获取验证结果时返回。
	Extra *string `json:"Extra" name:"Extra"`

	// 用于人脸比对的照片，图片的BASE64值；
	// BASE64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	ImageBase64 *string `json:"ImageBase64" name:"ImageBase64"`
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
		Url *string `json:"Url" name:"Url"`

		// 一次核身流程的标识，有效时间为7,200秒；
	// 完成核身后，可用该标识获取验证结果信息。
		BizToken *string `json:"BizToken" name:"BizToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
		ActionSequence *string `json:"ActionSequence" name:"ActionSequence"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	BizToken *string `json:"BizToken" name:"BizToken"`

	// 用于细分客户使用场景，由腾讯侧在线下对接时分配。
	RuleId *string `json:"RuleId" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证正反面；3：视频最佳截图照片；4：视频）。
	// 如 134表示拉取文本类、视频最佳截图照片、视频。
	InfoType *string `json:"InfoType" name:"InfoType"`
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
	//     "ErrMsg": null,       // 本次核身的错误信息。
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
		DetectInfo *string `json:"DetectInfo" name:"DetectInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
		LiveCode *string `json:"LiveCode" name:"LiveCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLiveCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLiveCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageRecognitionRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name" name:"Name"`

	// 用于人脸比对的照片，图片的BASE64值；
	// BASE64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	ImageBase64 *string `json:"ImageBase64" name:"ImageBase64"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional" name:"Optional"`
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
		Sim *float64 `json:"Sim" name:"Sim"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	ImageBase64 *string `json:"ImageBase64" name:"ImageBase64"`

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过5M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData" name:"ValidateData"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional" name:"Optional"`
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
		BestFrameBase64 *string `json:"BestFrameBase64" name:"BestFrameBase64"`

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）。
		Sim *float64 `json:"Sim" name:"Sim"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
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
	IdCard *string `json:"IdCard" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name" name:"Name"`

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过5M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData" name:"ValidateData"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional" name:"Optional"`
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
		BestFrameBase64 *string `json:"BestFrameBase64" name:"BestFrameBase64"`

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
		Sim *float64 `json:"Sim" name:"Sim"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *LivenessRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
