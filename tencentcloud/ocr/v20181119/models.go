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

package v20181119

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AdvertiseOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *AdvertiseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdvertiseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdvertiseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AdvertiseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
		TextDetections []*AdvertiseTextDetection `json:"TextDetections,omitempty" name:"TextDetections"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AdvertiseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdvertiseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvertiseTextDetection struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// 此字段为扩展字段。
	// GeneralBasicOcr接口返回段落信息Parag，包含ParagNo。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type ArithmeticOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 用于选择是否支持横屏拍摄。打开则支持横屏拍摄图片角度判断，角度信息在返回参数的angle中，默认值为true
	SupportHorizontalImage *bool `json:"SupportHorizontalImage,omitempty" name:"SupportHorizontalImage"`

	// 是否拒绝非速算图，打开则拒绝非速算图(注：非速算图是指风景人物等明显不是速算图片的图片)，默认值为false
	RejectNonArithmeticPic *bool `json:"RejectNonArithmeticPic,omitempty" name:"RejectNonArithmeticPic"`

	// 是否展开耦合算式中的竖式计算，默认值为false
	EnableDispRelatedVertical *bool `json:"EnableDispRelatedVertical,omitempty" name:"EnableDispRelatedVertical"`

	// 是否展示竖式算式的中间结果和格式控制字符，默认值为false
	EnableDispMidResult *bool `json:"EnableDispMidResult,omitempty" name:"EnableDispMidResult"`

	// 是否开启pdf识别，默认值为true
	EnablePdfRecognize *bool `json:"EnablePdfRecognize,omitempty" name:"EnablePdfRecognize"`

	// pdf页码，从0开始，默认为0
	PdfPageIndex *int64 `json:"PdfPageIndex,omitempty" name:"PdfPageIndex"`
}

func (r *ArithmeticOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArithmeticOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "SupportHorizontalImage")
	delete(f, "RejectNonArithmeticPic")
	delete(f, "EnableDispRelatedVertical")
	delete(f, "EnableDispMidResult")
	delete(f, "EnablePdfRecognize")
	delete(f, "PdfPageIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ArithmeticOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ArithmeticOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextArithmetic `json:"TextDetections,omitempty" name:"TextDetections"`

		// 图片横屏的角度(90度或270度)
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArithmeticOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArithmeticOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BankCardOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否返回预处理（精确剪裁对齐）后的银行卡图片数据，默认false。
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitempty" name:"RetBorderCutImage"`

	// 是否返回卡号的切图图片数据，默认false。
	RetCardNoImage *bool `json:"RetCardNoImage,omitempty" name:"RetCardNoImage"`

	// 复印件检测开关，如果输入的图片是银行卡复印件图片则返回告警，默认false。
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitempty" name:"EnableCopyCheck"`

	// 翻拍检测开关，如果输入的图片是银行卡翻拍图片则返回告警，默认false。
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitempty" name:"EnableReshootCheck"`

	// 边框遮挡检测开关，如果输入的图片是银行卡边框被遮挡则返回告警，默认false。
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitempty" name:"EnableBorderCheck"`

	// 是否返回图片质量分数（图片质量分数是评价一个图片的模糊程度的标准），默认false。
	EnableQualityValue *bool `json:"EnableQualityValue,omitempty" name:"EnableQualityValue"`
}

func (r *BankCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "RetBorderCutImage")
	delete(f, "RetCardNoImage")
	delete(f, "EnableCopyCheck")
	delete(f, "EnableReshootCheck")
	delete(f, "EnableBorderCheck")
	delete(f, "EnableQualityValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BankCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BankCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 卡号
		CardNo *string `json:"CardNo,omitempty" name:"CardNo"`

		// 银行信息
		BankInfo *string `json:"BankInfo,omitempty" name:"BankInfo"`

		// 有效期，格式如：07/2023
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 卡类型
		CardType *string `json:"CardType,omitempty" name:"CardType"`

		// 卡名字
		CardName *string `json:"CardName,omitempty" name:"CardName"`

		// 切片图片数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		BorderCutImage *string `json:"BorderCutImage,omitempty" name:"BorderCutImage"`

		// 卡号图片数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		CardNoImage *string `json:"CardNoImage,omitempty" name:"CardNoImage"`

		// WarningCode 告警码列表和释义：
	// -9110:银行卡日期无效; 
	// -9111:银行卡边框不完整; 
	// -9112:银行卡图片反光;
	// -9113:银行卡复印件;
	// -9114:银行卡翻拍件
	// （告警码可以同时存在多个）
	// 注意：此字段可能返回 null，表示取不到有效值。
		WarningCode []*int64 `json:"WarningCode,omitempty" name:"WarningCode"`

		// 图片质量分数，请求EnableQualityValue时返回（取值范围：0-100，分数越低越模糊，建议阈值≥50）。
	// 注意：此字段可能返回 null，表示取不到有效值。
		QualityValue *int64 `json:"QualityValue,omitempty" name:"QualityValue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BankCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BankSlipInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 付款开户行、收款开户行、付款账号、收款账号、回单类型、回单编号、币种、流水号、凭证号码、交易机构、交易金额、手续费、日期等字段信息。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type BankSlipOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *BankSlipOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankSlipOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BankSlipOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BankSlipOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 银行回单识别结果，具体内容请点击左侧链接。
		BankSlipInfos []*BankSlipInfo `json:"BankSlipInfos,omitempty" name:"BankSlipInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BankSlipOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankSlipOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BizLicenseOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *BizLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BizLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BizLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BizLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统一社会信用代码（三合一之前为注册号）
		RegNum *string `json:"RegNum,omitempty" name:"RegNum"`

		// 公司名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 注册资本
		Capital *string `json:"Capital,omitempty" name:"Capital"`

		// 法定代表人
		Person *string `json:"Person,omitempty" name:"Person"`

		// 地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 经营范围
		Business *string `json:"Business,omitempty" name:"Business"`

		// 主体类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 营业期限
		Period *string `json:"Period,omitempty" name:"Period"`

		// 组成形式
		ComposingForm *string `json:"ComposingForm,omitempty" name:"ComposingForm"`

		// 成立日期
		SetDate *string `json:"SetDate,omitempty" name:"SetDate"`

		// Code 告警码列表和释义：
	// -20001 非营业执照
	// 注：告警码可以同时存在多个
		RecognizeWarnCode []*int64 `json:"RecognizeWarnCode,omitempty" name:"RecognizeWarnCode"`

		// 告警码说明：
	// OCR_WARNING_TPYE_NOT_MATCH 非营业执照
	// 注：告警信息可以同时存在多个
		RecognizeWarnMsg []*string `json:"RecognizeWarnMsg,omitempty" name:"RecognizeWarnMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BizLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BizLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BizLicenseVerifyResult struct {

	// “0“：一致
	// “-1”：此号未查询到结果
	RegNum *string `json:"RegNum,omitempty" name:"RegNum"`

	// “0“：一致
	// “-1”：不一致
	// “”：不验真
	Name *string `json:"Name,omitempty" name:"Name"`

	// “0“：一致
	// “-1”：不一致
	// “”：不验真
	Address *string `json:"Address,omitempty" name:"Address"`
}

type BusInvoiceInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、发票号码、日期、票价、始发地、目的地、姓名、时间、发票消费类型、身份证号、省、市、开票日期、乘车地点、检票口、客票类型、车型、座位号、车次。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type BusInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *BusInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BusInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BusInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BusInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 汽车票识别结果，具体内容请点击左侧链接。
		BusInvoiceInfos []*BusInvoiceInfo `json:"BusInvoiceInfos,omitempty" name:"BusInvoiceInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BusInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BusInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BusinessCardInfo struct {

	// 识别出的字段名称（关键字，可能重复，比如多个手机），能识别的字段名为：
	// 姓名、英文姓名、英文地址、公司、英文公司、职位、英文职位、部门、英文部门、手机、电话、传真、社交帐号、QQ、MSN、微信、微博、邮箱、邮编、网址、公司账号、其他。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	ItemCoord *ItemCoord `json:"ItemCoord,omitempty" name:"ItemCoord"`
}

type BusinessCardOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 可选字段，根据需要选择是否请求对应字段。
	// 目前支持的字段为：
	// RetImageType-“PROPROCESS” 图像预处理，string 类型。
	// 图像预处理功能为，检测图片倾斜的角度，将原本倾斜的图片围绕中心点转正，最终输出一张正的名片抠图。
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"RetImageType":"PROPROCESS"})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"RetImageType":"PROPROCESS"}
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *BusinessCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BusinessCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BusinessCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BusinessCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 名片识别结果，具体内容请点击左侧链接。
		BusinessCardInfos []*BusinessCardInfo `json:"BusinessCardInfos,omitempty" name:"BusinessCardInfos"`

		// 返回图像预处理后的图片，图像预处理未开启时返回内容为空。
		RetImageBase64 *string `json:"RetImageBase64,omitempty" name:"RetImageBase64"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BusinessCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BusinessCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CandWord struct {

	// 候选字符集的单词信息（包括单词Character和单词置信度confidence）
	CandWords []*Words `json:"CandWords,omitempty" name:"CandWords"`
}

type CarInvoiceInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、 机打代码、 发票号码、 发动机号码、 合格证号、 机打号码、 价税合计(小写)、 销货单位名称、 身份证号码/组织机构代码、 购买方名称、 销售方纳税人识别号、 购买方纳税人识别号、主管税务机关、 主管税务机关代码、 开票日期、 不含税价(小写)、 吨位、增值税税率或征收率、 车辆识别代号/车架号码、 增值税税额、 厂牌型号、 省、 市、 发票消费类型、 销售方电话、 销售方账号、 产地、 进口证明书号、 车辆类型、 机器编号、备注、开票人、限乘人数、商检单号、销售方地址、销售方开户银行、价税合计、发票类型。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 字段在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`

	// 字段在原图中的中的四点坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon *Polygon `json:"Polygon,omitempty" name:"Polygon"`
}

type CarInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *CarInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CarInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CarInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CarInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 购车发票识别结果，具体内容请点击左侧链接。
		CarInvoiceInfos []*CarInvoiceInfo `json:"CarInvoiceInfos,omitempty" name:"CarInvoiceInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CarInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CarInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CellContent struct {

	// 段落编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParagNo *int64 `json:"ParagNo,omitempty" name:"ParagNo"`

	// 字体大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordSize *int64 `json:"WordSize,omitempty" name:"WordSize"`
}

type ClassifyDetectInfo struct {

	// 分类名称，包括：身份证、护照、名片、银行卡、行驶证、驾驶证、港澳台通行证、户口本、港澳台来往内地通行证、港澳台居住证、不动产证、营业执照
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分类类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 位置坐标
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type ClassifyDetectOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 可以指定要识别的票证类型,指定后不出现在此列表的票证将不返回类型。不指定时默认返回所有支持类别票证的识别信息。
	// 
	// 以下是当前支持的类型：
	// IDCardFront: 身份证正面识别
	// IDCardBack: 身份证背面识别
	// Passport: 护照
	// BusinessCard: 名片识别
	// BankCard: 银行卡识别
	// VehicleLicenseFront: 行驶证主页识别
	// VehicleLicenseBack: 行驶证副页识别
	// DriverLicenseFront: 驾驶证主页识别
	// DriverLicenseBack: 驾驶证副页识别
	// PermitFront: 港澳台通行证正面
	// ResidenceBooklet: 户口本资料页
	// MainlandPermitFront: 港澳台来往内地通行证正面
	// HmtResidentPermitFront: 港澳台居住证正面
	// HmtResidentPermitBack: 港澳台居住证背面
	// EstateCert: 不动产证
	// BizLicense: 营业执照
	DiscernType []*string `json:"DiscernType,omitempty" name:"DiscernType"`
}

func (r *ClassifyDetectOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClassifyDetectOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "DiscernType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClassifyDetectOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClassifyDetectOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 智能卡证分类结果。当图片类型不支持分类识别或者识别出的类型不在请求参数DiscernType指定的范围内时，返回结果中的Type字段将为空字符串，Name字段将返回"其它"
		ClassifyDetectInfos []*ClassifyDetectInfo `json:"ClassifyDetectInfos,omitempty" name:"ClassifyDetectInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClassifyDetectOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClassifyDetectOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Coord struct {

	// 横坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// 纵坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type Detail struct {

	// 企业四要素核验结果状态码
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 企业四要素核验结果描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type DetectedWordCoordPoint struct {

	// 单字在原图中的坐标，以四个顶点坐标表示，以左上角为起点，顺时针返回。
	WordCoordinate []*Coord `json:"WordCoordinate,omitempty" name:"WordCoordinate"`
}

type DetectedWords struct {

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 候选字Character
	Character *string `json:"Character,omitempty" name:"Character"`
}

type DriverLicenseOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// FRONT 为驾驶证主页正面（有红色印章的一面），
	// BACK 为驾驶证副页正面（有档案编号的一面）。
	// 默认值为：FRONT。
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`
}

func (r *DriverLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DriverLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DriverLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DriverLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 国籍
		Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

		// 住址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 出生日期（YYYY-MM-DD）
		DateOfBirth *string `json:"DateOfBirth,omitempty" name:"DateOfBirth"`

		// 初次领证日期（YYYY-MM-DD）
		DateOfFirstIssue *string `json:"DateOfFirstIssue,omitempty" name:"DateOfFirstIssue"`

		// 准驾车型
		Class *string `json:"Class,omitempty" name:"Class"`

		// 有效期开始时间（YYYY-MM-DD）
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 有效期截止时间（YYYY-MM-DD）
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// 证号
		CardCode *string `json:"CardCode,omitempty" name:"CardCode"`

		// 档案编号
		ArchivesCode *string `json:"ArchivesCode,omitempty" name:"ArchivesCode"`

		// 记录
		Record *string `json:"Record,omitempty" name:"Record"`

		// Code 告警码列表和释义：
	// -9102  复印件告警
	// -9103  翻拍件告警
	// -9106  ps告警
	// 注：告警码可以同时存在多个
		RecognizeWarnCode []*int64 `json:"RecognizeWarnCode,omitempty" name:"RecognizeWarnCode"`

		// 告警码说明：
	// WARN_DRIVER_LICENSE_COPY_CARD 复印件告警
	// WARN_DRIVER_LICENSE_SCREENED_CARD 翻拍件告警
	// WARN_DRIVER_LICENSE_PS_CARD ps告警
	// 注：告警信息可以同时存在多个
		RecognizeWarnMsg []*string `json:"RecognizeWarnMsg,omitempty" name:"RecognizeWarnMsg"`

		// 发证单位
		IssuingAuthority *string `json:"IssuingAuthority,omitempty" name:"IssuingAuthority"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DriverLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DriverLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DutyPaidProofInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 税号 、纳税人识别号 、纳税人名称 、金额合计大写 、金额合计小写 、填发日期 、税务机关 、填票人。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type DutyPaidProofOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *DutyPaidProofOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DutyPaidProofOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DutyPaidProofOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DutyPaidProofOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 完税证明识别结果，具体内容请点击左侧链接。
		DutyPaidProofInfos []*DutyPaidProofInfo `json:"DutyPaidProofInfos,omitempty" name:"DutyPaidProofInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DutyPaidProofOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DutyPaidProofOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EduPaperOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 扩展配置信息。
	// 配置格式：{"option1":value1,"option2":value2}
	// 1. task_type：任务类型【0: 关闭版式分析与处理 1: 开启版式分析处理】可选参数，Int32类型，默认值为1
	// 2. is_structuralization：是否结构化输出【true：返回包体同时返回通用和结构化输出  false：返回包体返回通用输出】 可选参数，Bool类型，默认值为true
	// 3. if_readable_format：是否按照版式整合通用文本/公式输出结果 可选参数，Bool类型，默认值为false
	// 示例：
	// {"task_type": 1,"is_structuralization": true,"if_readable_format": true}
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *EduPaperOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EduPaperOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EduPaperOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EduPaperOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		EduPaperInfos []*TextEduPaper `json:"EduPaperInfos,omitempty" name:"EduPaperInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。
		Angle *int64 `json:"Angle,omitempty" name:"Angle"`

		// 结构化方式输出，具体内容请点击左侧链接。
		QuestionBlockInfos []*QuestionBlockObj `json:"QuestionBlockInfos,omitempty" name:"QuestionBlockInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EduPaperOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EduPaperOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnglishOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 单词四点坐标开关，开启可返回图片中单词的四点坐标。
	// 该参数默认值为false。
	EnableCoordPoint *bool `json:"EnableCoordPoint,omitempty" name:"EnableCoordPoint"`

	// 候选字开关，开启可返回识别时多个可能的候选字（每个候选字对应其置信度）。
	// 该参数默认值为false。
	EnableCandWord *bool `json:"EnableCandWord,omitempty" name:"EnableCandWord"`

	// 预处理开关，功能是检测图片倾斜的角度，将原本倾斜的图片矫正。该参数默认值为true。
	Preprocess *bool `json:"Preprocess,omitempty" name:"Preprocess"`
}

func (r *EnglishOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnglishOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "EnableCoordPoint")
	delete(f, "EnableCandWord")
	delete(f, "Preprocess")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnglishOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnglishOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextDetectionEn `json:"TextDetections,omitempty" name:"TextDetections"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnglishOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnglishOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnterpriseLicenseInfo struct {

	// 识别出的字段名称（关键字），不同证件类型可能不同，证件类型包含企业登记证书、许可证书、企业执照、三证合一类证书；
	// 支持以下字段：统一社会信用代码、法定代表人、公司名称、公司地址、注册资金、企业关型、经营范围、成立日期、有效期、开办资金、经费来源、举办单位等；
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type EnterpriseLicenseOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *EnterpriseLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnterpriseLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnterpriseLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EnterpriseLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业证照识别结果，具体内容请点击左侧链接。
		EnterpriseLicenseInfos []*EnterpriseLicenseInfo `json:"EnterpriseLicenseInfos,omitempty" name:"EnterpriseLicenseInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnterpriseLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnterpriseLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EstateCertOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *EstateCertOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EstateCertOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EstateCertOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type EstateCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权利人
		Obligee *string `json:"Obligee,omitempty" name:"Obligee"`

		// 共有情况
		Ownership *string `json:"Ownership,omitempty" name:"Ownership"`

		// 坐落
		Location *string `json:"Location,omitempty" name:"Location"`

		// 不动产单元号
		Unit *string `json:"Unit,omitempty" name:"Unit"`

		// 权利类型
		Type *string `json:"Type,omitempty" name:"Type"`

		// 权利性质
		Property *string `json:"Property,omitempty" name:"Property"`

		// 用途
		Usage *string `json:"Usage,omitempty" name:"Usage"`

		// 面积
		Area *string `json:"Area,omitempty" name:"Area"`

		// 使用期限
		Term *string `json:"Term,omitempty" name:"Term"`

		// 权利其他状况，多行会用换行符\n连接。
		Other *string `json:"Other,omitempty" name:"Other"`

		// 图片旋转角度
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 不动产权号
		Number *string `json:"Number,omitempty" name:"Number"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EstateCertOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EstateCertOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FinanBillInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 【进账单】
	// 日期、出票全称、出票账号、出票开户行、收款人全称、收款人账号、收款开户行、大写金额、小写金额、票据种类、票据张数、票据号码；
	// 【支票】
	// 开户银行、支票种类、凭证号码2、日期、大写金额、小写金额、付款行编号、密码、凭证号码1；
	// 【银行承兑汇票】或【商业承兑汇票】
	// 出票日期、行号1、行号2、出票人全称、出票人账号、付款行全称、收款人全称、收款人账号、收款人开户行、出票金额大写、出票金额小写、汇票到期日、付款行行号、付款行地址。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type FinanBillOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *FinanBillOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FinanBillOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FinanBillOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FinanBillOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 金融票据整单识别结果，具体内容请点击左侧链接。
		FinanBillInfos []*FinanBillInfo `json:"FinanBillInfos,omitempty" name:"FinanBillInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FinanBillOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FinanBillOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FinanBillSliceInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 大写金额、小写金额、账号、票号1、票号2、收款人、大写日期、同城交换号、地址-省份、地址-城市、付款行全称、支票密码、支票用途。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type FinanBillSliceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *FinanBillSliceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FinanBillSliceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FinanBillSliceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FinanBillSliceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 金融票据切片识别结果，具体内容请点击左侧链接。
		FinanBillSliceInfos []*FinanBillSliceInfo `json:"FinanBillSliceInfos,omitempty" name:"FinanBillSliceInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FinanBillSliceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FinanBillSliceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlightInvoiceInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 票价、合计金额、填开日期、有效身份证件号码、电子客票号码、验证码、旅客姓名、填开单位、其他税费、燃油附加费、民航发展基金、保险费、销售单位代号、始发地、目的地、航班号、时间、日期、座位等级、承运人、发票消费类型、国内国际标签、印刷序号、客票级别/类别、客票生效日期、有效期截止日期、免费行李。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段 Name 对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 多个行程的字段所在行号，下标从0开始，非行字段或未能识别行号的该值返回-1。
	Row *int64 `json:"Row,omitempty" name:"Row"`
}

type FlightInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *FlightInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlightInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlightInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FlightInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机票行程单识别结果，具体内容请点击左侧链接。
		FlightInvoiceInfos []*FlightInvoiceInfo `json:"FlightInvoiceInfos,omitempty" name:"FlightInvoiceInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlightInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlightInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FormulaOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *FormulaOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FormulaOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FormulaOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FormulaOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负
		Angle *int64 `json:"Angle,omitempty" name:"Angle"`

		// 检测到的文本信息，具体内容请点击左侧链接。
		FormulaInfos []*TextFormula `json:"FormulaInfos,omitempty" name:"FormulaInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FormulaOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FormulaOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralAccurateOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否返回单字信息，默认关
	IsWords *bool `json:"IsWords,omitempty" name:"IsWords"`

	// 是否开启原图切图检测功能，开启后可提升“整图面积大，但单字符占比面积小”（例如：试卷）场景下的识别效果，默认关
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitempty" name:"EnableDetectSplit"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`
}

func (r *GeneralAccurateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralAccurateOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "IsWords")
	delete(f, "EnableDetectSplit")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralAccurateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GeneralAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralAccurateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralAccurateOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest

	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 保留字段。
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 识别语言类型。
	// 支持自动识别语言类型，同时支持自选语言种类，默认中英文混合(zh)，各种语言均支持与英文混合的文字识别。
	// 可选值：
	// zh：中英混合
	// zh_rare：支持英文、数字、中文生僻字、繁体字，特殊符号等
	// auto：自动
	// mix：混合语种
	// jap：日语
	// kor：韩语
	// spa：西班牙语
	// fre：法语
	// ger：德语
	// por：葡萄牙语
	// vie：越语
	// may：马来语
	// rus：俄语
	// ita：意大利语
	// hol：荷兰语
	// swe：瑞典语
	// fin：芬兰语
	// dan：丹麦语
	// nor：挪威语
	// hun：匈牙利语
	// tha：泰语
	// hi：印地语
	// ara：阿拉伯语
	LanguageType *string `json:"LanguageType,omitempty" name:"LanguageType"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`

	// 是否返回单字信息，默认关
	IsWords *bool `json:"IsWords,omitempty" name:"IsWords"`
}

func (r *GeneralBasicOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralBasicOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Scene")
	delete(f, "LanguageType")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "IsWords")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralBasicOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections"`

		// 检测到的语言类型，目前支持的语言类型参考入参LanguageType说明。
		Language *string `json:"Language,omitempty" name:"Language"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// 图片为PDF时，返回PDF的总页数，默认为0
		PdfPageSize *int64 `json:"PdfPageSize,omitempty" name:"PdfPageSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralBasicOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralBasicOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralEfficientOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *GeneralEfficientOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralEfficientOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralEfficientOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GeneralEfficientOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralEfficientOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralEfficientOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`
}

func (r *GeneralFastOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralFastOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralFastOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections"`

		// 检测到的语言，目前支持的语种范围为：简体中文、繁体中文、英文、日文、韩文。未来将陆续新增对更多语种的支持。
	// 返回结果含义为：zh - 中英混合，jap - 日文，kor - 韩文。
		Language *string `json:"Language,omitempty" name:"Language"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// 图片为PDF时，返回PDF的总页数，默认为0
		PdfPageSize *int64 `json:"PdfPageSize,omitempty" name:"PdfPageSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralFastOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralFastOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralHandwritingOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 场景字段，默认不用填写。
	// 可选值:only_hw  表示只输出手写体识别结果，过滤印刷体。
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 是否开启单字的四点定位坐标输出，默认值为false。
	EnableWordPolygon *bool `json:"EnableWordPolygon,omitempty" name:"EnableWordPolygon"`

	// 文本检测开关，默认值为true。
	// 设置为false表示直接进行单行识别，可适用于识别单行手写体签名场景。
	EnableDetectText *bool `json:"EnableDetectText,omitempty" name:"EnableDetectText"`
}

func (r *GeneralHandwritingOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralHandwritingOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Scene")
	delete(f, "EnableWordPolygon")
	delete(f, "EnableDetectText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralHandwritingOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GeneralHandwritingOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextGeneralHandwriting `json:"TextDetections,omitempty" name:"TextDetections"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
		Angel *float64 `json:"Angel,omitempty" name:"Angel"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralHandwritingOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneralHandwritingOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HKIDCardOCRRequest struct {
	*tchttp.BaseRequest

	// 是否鉴伪。
	DetectFake *bool `json:"DetectFake,omitempty" name:"DetectFake"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitempty" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *HKIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HKIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DetectFake")
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HKIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type HKIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 中文姓名
		CnName *string `json:"CnName,omitempty" name:"CnName"`

		// 英文姓名
		EnName *string `json:"EnName,omitempty" name:"EnName"`

		// 中文姓名对应电码
		TelexCode *string `json:"TelexCode,omitempty" name:"TelexCode"`

		// 性别 ：“男M”或“女F”
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 出生日期
		Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

		// 永久性居民身份证。
	// 0：非永久；
	// 1：永久；
	// -1：未知。
		Permanent *int64 `json:"Permanent,omitempty" name:"Permanent"`

		// 身份证号码
		IdNum *string `json:"IdNum,omitempty" name:"IdNum"`

		// 证件符号，出生日期下的符号，例如"***AZ"
		Symbol *string `json:"Symbol,omitempty" name:"Symbol"`

		// 首次签发日期
		FirstIssueDate *string `json:"FirstIssueDate,omitempty" name:"FirstIssueDate"`

		// 最近领用日期
		CurrentIssueDate *string `json:"CurrentIssueDate,omitempty" name:"CurrentIssueDate"`

		// 真假判断。
	// 0：无法判断（图像模糊、不完整、反光、过暗等导致无法判断）；
	// 1：假；
	// 2：真。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FakeDetectResult *int64 `json:"FakeDetectResult,omitempty" name:"FakeDetectResult"`

		// 人像照片Base64后的结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		HeadImage *string `json:"HeadImage,omitempty" name:"HeadImage"`

		// 多重告警码，当身份证是翻拍、复印、PS件时返回对应告警码。
	// -9102：证照复印件告警
	// -9103：证照翻拍告警
	// -9104：证照PS告警
	// -9105：证照防伪告警
		WarningCode []*int64 `json:"WarningCode,omitempty" name:"WarningCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HKIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HKIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HmtResidentPermitOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// FRONT：有照片的一面（人像面），
	// BACK：无照片的一面（国徽面），
	// 该参数如果不填或填错，将为您自动判断正反面。
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`
}

func (r *HmtResidentPermitOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HmtResidentPermitOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HmtResidentPermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type HmtResidentPermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证件姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 出生日期
		Birth *string `json:"Birth,omitempty" name:"Birth"`

		// 地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 身份证号
		IdCardNo *string `json:"IdCardNo,omitempty" name:"IdCardNo"`

		// 0-正面
	// 1-反面
		CardType *int64 `json:"CardType,omitempty" name:"CardType"`

		// 证件有效期限
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 签发机关
		Authority *string `json:"Authority,omitempty" name:"Authority"`

		// 签发次数
		VisaNum *string `json:"VisaNum,omitempty" name:"VisaNum"`

		// 通行证号码
		PassNo *string `json:"PassNo,omitempty" name:"PassNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HmtResidentPermitOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HmtResidentPermitOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IDCardOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// FRONT：身份证有照片的一面（人像面），
	// BACK：身份证有国徽的一面（国徽面），
	// 该参数如果不填，将为您自动判断身份证正反面。
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`

	// 以下可选字段均为bool 类型，默认false：
	// CropIdCard，身份证照片裁剪（去掉证件外多余的边缘、自动矫正拍摄角度）
	// CropPortrait，人像照片裁剪（自动抠取身份证头像区域）
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，PS检测告警
	// TempIdWarn，临时身份证告警
	// InvalidDateWarn，身份证有效日期不合法告警
	// Quality，图片质量分数（评价图片的模糊程度）
	// MultiCardDetect，是否开启多卡证检测
	// ReflectWarn，是否开启反光检测
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *IDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type IDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 姓名（人像面）
		Name *string `json:"Name,omitempty" name:"Name"`

		// 性别（人像面）
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 民族（人像面）
		Nation *string `json:"Nation,omitempty" name:"Nation"`

		// 出生日期（人像面）
		Birth *string `json:"Birth,omitempty" name:"Birth"`

		// 地址（人像面）
		Address *string `json:"Address,omitempty" name:"Address"`

		// 身份证号（人像面）
		IdNum *string `json:"IdNum,omitempty" name:"IdNum"`

		// 发证机关（国徽面）
		Authority *string `json:"Authority,omitempty" name:"Authority"`

		// 证件有效期（国徽面）
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 扩展信息，不请求则不返回，具体输入参考示例3和示例4。
	// IdCard，裁剪后身份证照片的base64编码，请求 Config.CropIdCard 时返回；
	// Portrait，身份证头像照片的base64编码，请求 Config.CropPortrait 时返回；
	// 
	// Quality，图片质量分数，请求 Config.Quality 时返回（取值范围：0 ~ 100，分数越低越模糊，建议阈值≥50）;
	// BorderCodeValue，身份证边框不完整告警阈值分数，请求 Config.BorderCheckWarn时返回（取值范围：0 ~ 100，分数越低边框遮挡可能性越低，建议阈值≤50）;
	// 
	// WarnInfos，告警信息，Code 告警码列表和释义：
	// -9100	身份证有效日期不合法告警，
	// -9101	身份证边框不完整告警，
	// -9102	身份证复印件告警，
	// -9103	身份证翻拍告警，
	// -9105	身份证框内遮挡告警，
	// -9104	临时身份证告警，
	// -9106	身份证 PS 告警，
	// -9107       身份证反光告警。
		AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstitutionOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *InstitutionOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstitutionOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstitutionOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InstitutionOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 注册号
		RegId *string `json:"RegId,omitempty" name:"RegId"`

		// 有效期
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 住所
		Location *string `json:"Location,omitempty" name:"Location"`

		// 名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 法定代表人
		LegalPerson *string `json:"LegalPerson,omitempty" name:"LegalPerson"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstitutionOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstitutionOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsuranceBillInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 【病案首页】
	// 姓名、性别、出生日期、出院诊断、疾病编码、入院病情等。
	// 【费用清单】
	// 医疗参保人员类别、身份证号、入院方式、结账日期、项目、金额等。
	// 【结算单】
	// 名称、单价、数量、金额、医保内、医保外等。
	// 【医疗发票】
	// 姓名、性别、住院时间、收费项目、金额、合计等。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InsuranceBillOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *InsuranceBillOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InsuranceBillOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InsuranceBillOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InsuranceBillOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 保险单据识别结果，具体内容请点击左侧链接。
		InsuranceBillInfos []*InsuranceBillInfo `json:"InsuranceBillInfos,omitempty" name:"InsuranceBillInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InsuranceBillOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InsuranceBillOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InvoiceDetectInfo struct {

	// 识别出的图片在混贴票据图片中的旋转角度。
	Angle *float64 `json:"Angle,omitempty" name:"Angle"`

	// 识别出的图片所属的票据类型。
	// -1：未知类型
	// 0：出租车发票
	// 1：定额发票
	// 2：火车票
	// 3：增值税发票
	// 4：客运限额发票
	// 5：机票行程单
	// 6：酒店账单
	// 7：完税证明
	// 8：通用机打发票
	// 9：汽车票
	// 10：轮船票
	// 11：增值税发票（卷票 ）
	// 12：购车发票
	// 13：过路过桥费发票
	// 14：购物小票
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 识别出的图片在混贴票据图片中的位置信息。与Angel结合可以得出原图位置，组成RotatedRect((X+0.5\*Width,Y+0.5\*Height), (Width, Height), Angle)，详情可参考OpenCV文档。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`

	// 入参 ReturnImage 为 True 时返回 Base64 编码后的图片。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitempty" name:"Image"`
}

type InvoiceGeneralInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、发票号码、日期、合计金额(小写)、合计金额(大写)、购买方识别号、销售方识别号、校验码、购买方名称、销售方名称、时间、种类、发票消费类型、省、市、是否有公司印章。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type InvoiceGeneralOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *InvoiceGeneralOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvoiceGeneralOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvoiceGeneralOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InvoiceGeneralOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通用机打发票识别结果，具体内容请点击左侧链接。
		InvoiceGeneralInfos []*InvoiceGeneralInfo `json:"InvoiceGeneralInfos,omitempty" name:"InvoiceGeneralInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InvoiceGeneralOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvoiceGeneralOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemCoord struct {

	// 左上角x
	X *int64 `json:"X,omitempty" name:"X"`

	// 左上角y
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 宽width
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高height
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type LicensePlateOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *LicensePlateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LicensePlateOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LicensePlateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type LicensePlateOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别出的车牌号码。
		Number *string `json:"Number,omitempty" name:"Number"`

		// 置信度，0 - 100 之间。
		Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

		// 文本行在原图片中的像素坐标框。
		Rect *Rect `json:"Rect,omitempty" name:"Rect"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LicensePlateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LicensePlateOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MLIDCardOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。( 中国地区之外不支持这个字段 )
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否返回图片，默认false
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

func (r *MLIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "RetImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MLIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MLIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 身份证号
		ID *string `json:"ID,omitempty" name:"ID"`

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 告警码
	// -9103	证照翻拍告警
	// -9102	证照复印件告警
	// -9106       证件遮挡告警
		Warn []*int64 `json:"Warn,omitempty" name:"Warn"`

		// 证件图片
		Image *string `json:"Image,omitempty" name:"Image"`

		// 扩展字段：
	// {
	//     ID:{
	//         Confidence:0.9999
	//     },
	//     Name:{
	//         Confidence:0.9996
	//     }
	// }
		AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

		// 证件类型
	// MyKad  身份证
	// MyPR    永居证
	// MyTentera   军官证
	// MyKAS    临时身份证
	// POLIS  警察证
	// IKAD   劳工证
	// MyKid 儿童卡
		Type *string `json:"Type,omitempty" name:"Type"`

		// 出生日期（目前该字段仅支持IKAD劳工证）
		Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MLIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 是否返回图片，默认false
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

func (r *MLIDPassportOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDPassportOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "RetImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MLIDPassportOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MLIDPassportOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 护照ID
		ID *string `json:"ID,omitempty" name:"ID"`

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 出生日期
		DateOfBirth *string `json:"DateOfBirth,omitempty" name:"DateOfBirth"`

		// 性别（F女，M男）
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 有效期
		DateOfExpiration *string `json:"DateOfExpiration,omitempty" name:"DateOfExpiration"`

		// 发行国
		IssuingCountry *string `json:"IssuingCountry,omitempty" name:"IssuingCountry"`

		// 国籍
		Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

		// 告警码
	// -9103	证照翻拍告警
	// -9102	证照复印件告警
	// -9106       证件遮挡告警
		Warn []*int64 `json:"Warn,omitempty" name:"Warn"`

		// 证件图片
		Image *string `json:"Image,omitempty" name:"Image"`

		// 扩展字段:
	// {
	//     ID:{
	//         Confidence:0.9999
	//     },
	//     Name:{
	//         Confidence:0.9996
	//     }
	// }
		AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MLIDPassportOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MLIDPassportOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MainlandPermitOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否返回头像。默认不返回。
	RetProfile *bool `json:"RetProfile,omitempty" name:"RetProfile"`
}

func (r *MainlandPermitOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MainlandPermitOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "RetProfile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MainlandPermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MainlandPermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 中文姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 英文姓名
		EnglishName *string `json:"EnglishName,omitempty" name:"EnglishName"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 出生日期
		Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

		// 签发机关
		IssueAuthority *string `json:"IssueAuthority,omitempty" name:"IssueAuthority"`

		// 有效期限
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 证件号
		Number *string `json:"Number,omitempty" name:"Number"`

		// 签发地点
		IssueAddress *string `json:"IssueAddress,omitempty" name:"IssueAddress"`

		// 签发次数
		IssueNumber *string `json:"IssueNumber,omitempty" name:"IssueNumber"`

		// 证件类别， 如：台湾居民来往大陆通行证、港澳居民来往内地通行证。
		Type *string `json:"Type,omitempty" name:"Type"`

		// RetProfile为True时返回头像字段， Base64编码
		Profile *string `json:"Profile,omitempty" name:"Profile"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MainlandPermitOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MainlandPermitOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceDetectRequest struct {
	*tchttp.BaseRequest

	// 是否需要返回裁剪后的图片。
	ReturnImage *bool `json:"ReturnImage,omitempty" name:"ReturnImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *MixedInvoiceDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MixedInvoiceDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReturnImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MixedInvoiceDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测出的票据类型列表，具体内容请点击左侧链接。
		InvoiceDetectInfos []*InvoiceDetectInfo `json:"InvoiceDetectInfos,omitempty" name:"InvoiceDetectInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MixedInvoiceDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MixedInvoiceDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceItem struct {

	// 识别结果。
	// OK：表示识别成功；FailedOperation.UnsupportedInvioce：表示不支持识别；
	// FailedOperation.UnKnowError：表示识别失败；
	// 其它错误码见各个票据接口的定义。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 识别出的图片所属的票据类型。
	// -1：未知类型
	// 0：出租车发票
	// 1：定额发票
	// 2：火车票
	// 3：增值税发票
	// 5：机票行程单
	// 8：通用机打发票
	// 9：汽车票
	// 10：轮船票
	// 11：增值税发票（卷票 ）
	// 12：购车发票
	// 13：过路过桥费发票
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 识别出的图片在混贴票据图片中的位置信息。与Angel结合可以得出原图位置，组成RotatedRect((X+0.5\*Width,Y+0.5\*Height), (Width, Height), Angle)，详情可参考OpenCV文档。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`

	// 识别出的图片在混贴票据图片中的旋转角度。
	Angle *float64 `json:"Angle,omitempty" name:"Angle"`

	// 识别到的内容。
	SingleInvoiceInfos []*SingleInvoiceInfo `json:"SingleInvoiceInfos,omitempty" name:"SingleInvoiceInfos"`
}

type MixedInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 需要识别的票据类型列表，为空或不填表示识别全部类型。
	// 0：出租车发票
	// 1：定额发票
	// 2：火车票
	// 3：增值税发票
	// 5：机票行程单
	// 8：通用机打发票
	// 9：汽车票
	// 10：轮船票
	// 11：增值税发票（卷票 ）
	// 12：购车发票
	// 13：过路过桥费发票
	Types []*int64 `json:"Types,omitempty" name:"Types"`
}

func (r *MixedInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MixedInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Types")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MixedInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 混贴票据识别结果，具体内容请点击左侧链接。
		MixedInvoiceItems []*MixedInvoiceItem `json:"MixedInvoiceItems,omitempty" name:"MixedInvoiceItems"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MixedInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MixedInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OnlineTaxiItineraryInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、 机打代码、 发票号码、 发动机号码、 合格证号、 机打号码、 价税合计(小写)、 销货单位名称、 身份证号码/组织机构代码、 购买方名称、 销售方纳税人识别号、 购买方纳税人识别号、主管税务机关、 主管税务机关代码、 开票日期、 不含税价(小写)、 吨位、增值税税率或征收率、 车辆识别代号/车架号码、 增值税税额、 厂牌型号、 省、 市、 发票消费类型、 销售方电话、 销售方账号、 产地、 进口证明书号、 车辆类型、 机器编号、备注、开票人、限乘人数、商检单号、销售方地址、销售方开户银行、价税合计、发票类型。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type OrgCodeCertOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *OrgCodeCertOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OrgCodeCertOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OrgCodeCertOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type OrgCodeCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代码
		OrgCode *string `json:"OrgCode,omitempty" name:"OrgCode"`

		// 机构名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// 地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 有效期
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OrgCodeCertOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OrgCodeCertOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PassportOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 默认填写CN
	// 支持中国大陆地区护照。
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *PassportOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PassportOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PassportOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PassportOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 国家码
		Country *string `json:"Country,omitempty" name:"Country"`

		// 护照号
		PassportNo *string `json:"PassportNo,omitempty" name:"PassportNo"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 国籍
		Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

		// 出生日期
		BirthDate *string `json:"BirthDate,omitempty" name:"BirthDate"`

		// 出生地点
		BirthPlace *string `json:"BirthPlace,omitempty" name:"BirthPlace"`

		// 签发日期
		IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`

		// 签发地点
		IssuePlace *string `json:"IssuePlace,omitempty" name:"IssuePlace"`

		// 有效期
		ExpiryDate *string `json:"ExpiryDate,omitempty" name:"ExpiryDate"`

		// 持证人签名
		Signature *string `json:"Signature,omitempty" name:"Signature"`

		// 最下方第一行 MRZ Code 序列
		CodeSet *string `json:"CodeSet,omitempty" name:"CodeSet"`

		// 最下方第二行 MRZ Code 序列
		CodeCrc *string `json:"CodeCrc,omitempty" name:"CodeCrc"`

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 姓
		FamilyName *string `json:"FamilyName,omitempty" name:"FamilyName"`

		// 名
		FirstName *string `json:"FirstName,omitempty" name:"FirstName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PassportOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PassportOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PermitOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *PermitOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PermitOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 英文姓名
		EnglishName *string `json:"EnglishName,omitempty" name:"EnglishName"`

		// 证件号
		Number *string `json:"Number,omitempty" name:"Number"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 有效期限
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 签发机关
		IssueAuthority *string `json:"IssueAuthority,omitempty" name:"IssueAuthority"`

		// 签发地点
		IssueAddress *string `json:"IssueAddress,omitempty" name:"IssueAddress"`

		// 出生日期
		Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PermitOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PermitOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Polygon struct {

	// 左上顶点坐标
	LeftTop *Coord `json:"LeftTop,omitempty" name:"LeftTop"`

	// 右上顶点坐标
	RightTop *Coord `json:"RightTop,omitempty" name:"RightTop"`

	// 右下顶点坐标
	RightBottom *Coord `json:"RightBottom,omitempty" name:"RightBottom"`

	// 左下顶点坐标
	LeftBottom *Coord `json:"LeftBottom,omitempty" name:"LeftBottom"`
}

type ProductDataRecord struct {

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品名称(英文)
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// 品牌名称
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// 规格型号
	Type *string `json:"Type,omitempty" name:"Type"`

	// 宽度，单位毫米
	Width *string `json:"Width,omitempty" name:"Width"`

	// 高度，单位毫米
	Height *string `json:"Height,omitempty" name:"Height"`

	// 深度，单位毫米
	Depth *string `json:"Depth,omitempty" name:"Depth"`

	// 关键字
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 简短描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 图片链接
	ImageLink []*string `json:"ImageLink,omitempty" name:"ImageLink"`

	// 厂家名称
	ManufacturerName *string `json:"ManufacturerName,omitempty" name:"ManufacturerName"`

	// 厂家地址
	ManufacturerAddress *string `json:"ManufacturerAddress,omitempty" name:"ManufacturerAddress"`

	// 企业社会信用代码
	FirmCode *string `json:"FirmCode,omitempty" name:"FirmCode"`

	// 表示数据查询状态
	// checkResult	状态说明
	// 1	 经查，该商品条码已在中国物品编码中心注册
	// 2	经查，该厂商识别代码已在中国物品编码中心注册，但编码信息未按规定通报。
	// 3	经查，该厂商识别代码已于xxxxx注销，请关注产品生产日期。
	// 4	经查，该企业以及条码未经条码中心注册，属于违法使用
	// -1	经查，该商品条码被冒用
	// -2	经查，该厂商识别代码已在中国物品编码中心注册，但该产品已经下市
	// S001                未找到该厂商识别代码的注册信息。
	// S002		该厂商识别代码已经在GS1注册，但编码信息未通报
	// S003		该商品条码已在GS1通报
	// S004		该商品条码已注销
	// S005		数字不正确。GS1前缀（3位国家/地区代码）用于特殊用途。
	// E001		完整性失败：此GTIN的长度无效。
	// E002		完整性失败：校验位不正确。
	// E003		完整性失败：字符串包含字母数字字符。
	// E004		数字不正确。GS1前缀（3位国家/地区代码）不存在。
	// E005		数字不正确。GS1前缀（3位国家/地区代码）用于特殊用途。
	// E006		数字不正确。尚未分配该GS1公司前缀。
	// E008	        经查，该企业厂商识别代码以及条码尚未通报
	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`

	// UNSPSC分类码
	CategoryCode *string `json:"CategoryCode,omitempty" name:"CategoryCode"`
}

type PropOwnerCertOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *PropOwnerCertOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PropOwnerCertOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PropOwnerCertOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PropOwnerCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 房地产权利人
		Owner *string `json:"Owner,omitempty" name:"Owner"`

		// 共有情况
		Possession *string `json:"Possession,omitempty" name:"Possession"`

		// 登记时间
		RegisterTime *string `json:"RegisterTime,omitempty" name:"RegisterTime"`

		// 规划用途
		Purpose *string `json:"Purpose,omitempty" name:"Purpose"`

		// 房屋性质
		Nature *string `json:"Nature,omitempty" name:"Nature"`

		// 房地坐落
		Location *string `json:"Location,omitempty" name:"Location"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PropOwnerCertOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PropOwnerCertOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QrcodeImgSize struct {

	// 宽
	Wide *int64 `json:"Wide,omitempty" name:"Wide"`

	// 高
	High *int64 `json:"High,omitempty" name:"High"`
}

type QrcodeOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，支持PNG、JPG、JPEG格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，支持PNG、JPG、JPEG格式。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *QrcodeOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QrcodeOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QrcodeOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QrcodeOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 二维码/条形码识别结果信息，具体内容请点击左侧链接。
		CodeResults []*QrcodeResultsInfo `json:"CodeResults,omitempty" name:"CodeResults"`

		// 图片大小，具体内容请点击左侧链接。
		ImgSize *QrcodeImgSize `json:"ImgSize,omitempty" name:"ImgSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QrcodeOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QrcodeOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QrcodePositionObj struct {

	// 左上顶点坐标（如果是条形码，X和Y都为-1）
	LeftTop *Coord `json:"LeftTop,omitempty" name:"LeftTop"`

	// 右上顶点坐标（如果是条形码，X和Y都为-1）
	RightTop *Coord `json:"RightTop,omitempty" name:"RightTop"`

	// 右下顶点坐标（如果是条形码，X和Y都为-1）
	RightBottom *Coord `json:"RightBottom,omitempty" name:"RightBottom"`

	// 左下顶点坐标（如果是条形码，X和Y都为-1）
	LeftBottom *Coord `json:"LeftBottom,omitempty" name:"LeftBottom"`
}

type QrcodeResultsInfo struct {

	// 类型（二维码、条形码）
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 二维码/条形码包含的地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 二维码/条形码坐标（二维码会返回位置坐标，条形码暂不返回位置坐标，因此默认X和Y返回值均为-1）
	Position *QrcodePositionObj `json:"Position,omitempty" name:"Position"`
}

type QueryBarCodeRequest struct {
	*tchttp.BaseRequest

	// 条形码
	BarCode *string `json:"BarCode,omitempty" name:"BarCode"`
}

func (r *QueryBarCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBarCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BarCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBarCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryBarCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 条码
		BarCode *string `json:"BarCode,omitempty" name:"BarCode"`

		// 条码信息数组
		ProductDataRecords []*ProductDataRecord `json:"ProductDataRecords,omitempty" name:"ProductDataRecords"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBarCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBarCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuestionBlockObj struct {

	// 数学试题识别结构化信息数组
	QuestionArr []*QuestionObj `json:"QuestionArr,omitempty" name:"QuestionArr"`

	// 题目主体区域检测框在图片中的像素坐标
	QuestionBboxCoord *Rect `json:"QuestionBboxCoord,omitempty" name:"QuestionBboxCoord"`
}

type QuestionObj struct {

	// 题号
	QuestionTextNo *string `json:"QuestionTextNo,omitempty" name:"QuestionTextNo"`

	// 题型：
	// 1: "选择题"
	// 2: "填空题"
	// 3: "解答题"
	QuestionTextType *int64 `json:"QuestionTextType,omitempty" name:"QuestionTextType"`

	// 题干
	QuestionText *string `json:"QuestionText,omitempty" name:"QuestionText"`

	// 选择题选项，包含1个或多个option
	QuestionOptions *string `json:"QuestionOptions,omitempty" name:"QuestionOptions"`

	// 所有子题的question属性
	QuestionSubquestion *string `json:"QuestionSubquestion,omitempty" name:"QuestionSubquestion"`

	// 示意图检测框在的图片中的像素坐标
	QuestionImageCoords []*Rect `json:"QuestionImageCoords,omitempty" name:"QuestionImageCoords"`
}

type QuotaInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *QuotaInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuotaInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuotaInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QuotaInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票号码
		InvoiceNum *string `json:"InvoiceNum,omitempty" name:"InvoiceNum"`

		// 发票代码
		InvoiceCode *string `json:"InvoiceCode,omitempty" name:"InvoiceCode"`

		// 大写金额
		Rate *string `json:"Rate,omitempty" name:"Rate"`

		// 小写金额
		RateNum *string `json:"RateNum,omitempty" name:"RateNum"`

		// 发票消费类型
		InvoiceType *string `json:"InvoiceType,omitempty" name:"InvoiceType"`

		// 省
	// 注意：此字段可能返回 null，表示取不到有效值。
		Province *string `json:"Province,omitempty" name:"Province"`

		// 市
	// 注意：此字段可能返回 null，表示取不到有效值。
		City *string `json:"City,omitempty" name:"City"`

		// 是否有公司印章（1有 0无 空为识别不出）
	// 注意：此字段可能返回 null，表示取不到有效值。
		HasStamp *string `json:"HasStamp,omitempty" name:"HasStamp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuotaInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuotaInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeContainerOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *RecognizeContainerOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeContainerOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeContainerOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeContainerOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集装箱箱号
		ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`

		// 集装箱类型
		ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

		// 集装箱总重量，单位：千克（KG）
		GrossKG *string `json:"GrossKG,omitempty" name:"GrossKG"`

		// 集装箱总重量，单位：磅（LB）
		GrossLB *string `json:"GrossLB,omitempty" name:"GrossLB"`

		// 集装箱有效承重，单位：千克（KG）
		PayloadKG *string `json:"PayloadKG,omitempty" name:"PayloadKG"`

		// 集装箱有效承重，单位：磅（LB）
		PayloadLB *string `json:"PayloadLB,omitempty" name:"PayloadLB"`

		// 集装箱容量，单位：立方米
		CapacityM3 *string `json:"CapacityM3,omitempty" name:"CapacityM3"`

		// 集装箱容量，单位：立英尺
		CapacityFT3 *string `json:"CapacityFT3,omitempty" name:"CapacityFT3"`

		// 告警码
	// -9926	集装箱箱号不完整或者不清晰
	// -9927	集装箱类型不完整或者不清晰
		Warn []*int64 `json:"Warn,omitempty" name:"Warn"`

		// 集装箱自身重量，单位：千克（KG）
		TareKG *string `json:"TareKG,omitempty" name:"TareKG"`

		// 集装箱自身重量，单位：磅（LB）
		TareLB *string `json:"TareLB,omitempty" name:"TareLB"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizeContainerOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeContainerOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeOnlineTaxiItineraryOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`
}

func (r *RecognizeOnlineTaxiItineraryOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeOnlineTaxiItineraryOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeOnlineTaxiItineraryOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeOnlineTaxiItineraryOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 网约车行程单识别结果，具体内容请点击左侧链接。
		OnlineTaxiItineraryInfos []*OnlineTaxiItineraryInfo `json:"OnlineTaxiItineraryInfos,omitempty" name:"OnlineTaxiItineraryInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizeOnlineTaxiItineraryOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeOnlineTaxiItineraryOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeTableOCRRequest struct {
	*tchttp.BaseRequest

	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`

	// 语言，zh：中英文（默认）jap：日文
	TableLanguage *string `json:"TableLanguage,omitempty" name:"TableLanguage"`
}

func (r *RecognizeTableOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTableOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "TableLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeTableOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeTableOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TableDetections []*TableDetectInfo `json:"TableDetections,omitempty" name:"TableDetections"`

		// Base64 编码后的 Excel 数据。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 图片为PDF时，返回PDF的总页数，默认为0
		PdfPageSize *int64 `json:"PdfPageSize,omitempty" name:"PdfPageSize"`

		// 图片旋转角度（角度制），文本的水平方向为0°，统一以逆时针方向旋转，逆时针为负，角度范围为-360°至0°。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizeTableOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTableOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeThaiIDCardOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *RecognizeThaiIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeThaiIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeThaiIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RecognizeThaiIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 身份证号码
		ID *string `json:"ID,omitempty" name:"ID"`

		// 泰文姓名
		ThaiName *string `json:"ThaiName,omitempty" name:"ThaiName"`

		// 英文姓名
		EnFirstName *string `json:"EnFirstName,omitempty" name:"EnFirstName"`

		// 地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 出生日期
		Birthday *string `json:"Birthday,omitempty" name:"Birthday"`

		// 首次领用日期
		IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`

		// 签发日期
		ExpirationDate *string `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

		// 英文姓名
		EnLastName *string `json:"EnLastName,omitempty" name:"EnLastName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizeThaiIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeThaiIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rect struct {

	// 左上角x
	X *int64 `json:"X,omitempty" name:"X"`

	// 左上角y
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高度
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type ResidenceBookletOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *ResidenceBookletOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResidenceBookletOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResidenceBookletOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResidenceBookletOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 户号
		HouseholdNumber *string `json:"HouseholdNumber,omitempty" name:"HouseholdNumber"`

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 出生地
		BirthPlace *string `json:"BirthPlace,omitempty" name:"BirthPlace"`

		// 民族
		Nation *string `json:"Nation,omitempty" name:"Nation"`

		// 籍贯
		NativePlace *string `json:"NativePlace,omitempty" name:"NativePlace"`

		// 出生日期
		BirthDate *string `json:"BirthDate,omitempty" name:"BirthDate"`

		// 公民身份证件编号
		IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

		// 文化程度
		EducationDegree *string `json:"EducationDegree,omitempty" name:"EducationDegree"`

		// 服务处所
		ServicePlace *string `json:"ServicePlace,omitempty" name:"ServicePlace"`

		// 户别
		Household *string `json:"Household,omitempty" name:"Household"`

		// 住址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 承办人签章文字
		Signature *string `json:"Signature,omitempty" name:"Signature"`

		// 签发日期
		IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`

		// 户主页编号
		HomePageNumber *string `json:"HomePageNumber,omitempty" name:"HomePageNumber"`

		// 户主姓名
		HouseholderName *string `json:"HouseholderName,omitempty" name:"HouseholderName"`

		// 户主或与户主关系
		Relationship *string `json:"Relationship,omitempty" name:"Relationship"`

		// 本市（县）其他住址
		OtherAddresses *string `json:"OtherAddresses,omitempty" name:"OtherAddresses"`

		// 宗教信仰
		ReligiousBelief *string `json:"ReligiousBelief,omitempty" name:"ReligiousBelief"`

		// 身高
		Height *string `json:"Height,omitempty" name:"Height"`

		// 血型
		BloodType *string `json:"BloodType,omitempty" name:"BloodType"`

		// 婚姻状况
		MaritalStatus *string `json:"MaritalStatus,omitempty" name:"MaritalStatus"`

		// 兵役状况
		VeteranStatus *string `json:"VeteranStatus,omitempty" name:"VeteranStatus"`

		// 职业
		Profession *string `json:"Profession,omitempty" name:"Profession"`

		// 何时由何地迁来本市(县)
		MoveToCityInformation *string `json:"MoveToCityInformation,omitempty" name:"MoveToCityInformation"`

		// 何时由何地迁来本址
		MoveToSiteInformation *string `json:"MoveToSiteInformation,omitempty" name:"MoveToSiteInformation"`

		// 登记日期
		RegistrationDate *string `json:"RegistrationDate,omitempty" name:"RegistrationDate"`

		// 曾用名
		FormerName *string `json:"FormerName,omitempty" name:"FormerName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResidenceBookletOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResidenceBookletOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RideHailingDriverLicenseOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *RideHailingDriverLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RideHailingDriverLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RideHailingDriverLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RideHailingDriverLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 证号，对应网约车驾驶证字段：证号/从业资格证号/驾驶员证号/身份证号
		LicenseNumber *string `json:"LicenseNumber,omitempty" name:"LicenseNumber"`

		// 有效起始日期
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 有效期截止时间，对应网约车驾驶证字段：有效期至/营运期限止
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// 初始发证日期，对应网约车驾驶证字段：初始领证日期/发证日期
		ReleaseDate *string `json:"ReleaseDate,omitempty" name:"ReleaseDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RideHailingDriverLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RideHailingDriverLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RideHailingTransportLicenseOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *RideHailingTransportLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RideHailingTransportLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RideHailingTransportLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RideHailingTransportLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交运管许可字号。
		OperationLicenseNumber *string `json:"OperationLicenseNumber,omitempty" name:"OperationLicenseNumber"`

		// 车辆所有人，对应网约车运输证字段：车辆所有人/车主名称/业户名称。
		VehicleOwner *string `json:"VehicleOwner,omitempty" name:"VehicleOwner"`

		// 车牌号码，对应网约车运输证字段：车牌号码/车辆号牌。
		VehicleNumber *string `json:"VehicleNumber,omitempty" name:"VehicleNumber"`

		// 有效起始日期。
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 有效期截止时间，对应网约车运输证字段：有效期至/营运期限止。
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// 初始发证日期，对应网约车运输证字段：初始领证日期/发证日期。
		ReleaseDate *string `json:"ReleaseDate,omitempty" name:"ReleaseDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RideHailingTransportLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RideHailingTransportLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SealOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *SealOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SealOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SealOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SealOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 印章内容
		SealBody *string `json:"SealBody,omitempty" name:"SealBody"`

		// 印章坐标
		Location *Rect `json:"Location,omitempty" name:"Location"`

		// 其它文本内容
		OtherTexts []*string `json:"OtherTexts,omitempty" name:"OtherTexts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SealOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SealOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShipInvoiceInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、发票号码、日期、票价、始发地、目的地、姓名、时间、发票消费类型、省、市、币种。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type ShipInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *ShipInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShipInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ShipInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ShipInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 轮船票识别结果，具体内容请点击左侧链接。
		ShipInvoiceInfos []*ShipInvoiceInfo `json:"ShipInvoiceInfos,omitempty" name:"ShipInvoiceInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShipInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ShipInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SingleInvoiceInfo struct {

	// 识别出的字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SmartStructuralOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 需返回的字段名称，例：
	// 若客户只想返回姓名、性别两个字段的识别结果，则输入
	// ItemNames=["姓名","性别"]
	ItemNames []*string `json:"ItemNames,omitempty" name:"ItemNames"`
}

func (r *SmartStructuralOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "ItemNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmartStructuralOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SmartStructuralOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片旋转角度(角度制)，文本的水平方向
	// 为 0；顺时针为正，逆时针为负
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 识别信息
		StructuralItems []*StructuralItem `json:"StructuralItems,omitempty" name:"StructuralItems"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SmartStructuralOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StructuralItem struct {

	// 识别出的字段名称(关键字)。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 置信度 0 ~100。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行在旋转纠正之后的图像中的像素
	// 坐标。
	ItemCoord *ItemCoord `json:"ItemCoord,omitempty" name:"ItemCoord"`
}

type TableCell struct {

	// 单元格左上角的列索引
	ColTl *int64 `json:"ColTl,omitempty" name:"ColTl"`

	// 单元格左上角的行索引
	RowTl *int64 `json:"RowTl,omitempty" name:"RowTl"`

	// 单元格右下角的列索引
	ColBr *int64 `json:"ColBr,omitempty" name:"ColBr"`

	// 单元格右下角的行索引
	RowBr *int64 `json:"RowBr,omitempty" name:"RowBr"`

	// 单元格内识别出的字符串文本，若文本存在多行，以换行符"\n"隔开
	Text *string `json:"Text,omitempty" name:"Text"`

	// 单元格类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 单元格置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 单元格在图像中的四点坐标
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// 此字段为扩展字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 单元格文本属性
	Contents []*CellContent `json:"Contents,omitempty" name:"Contents"`
}

type TableDetectInfo struct {

	// 单元格内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cells []*TableCell `json:"Cells,omitempty" name:"Cells"`

	// 表格标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Titles []*TableTitle `json:"Titles,omitempty" name:"Titles"`

	// 图像中的文本块类型，0 为非表格文本，
	// 1 为有线表格，2 为无线表格
	// （接口暂不支持日文无线表格识别，若传入日文无线表格，返回0）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 表格主体四个顶点坐标（依次为左上角，
	// 右上角，右下角，左下角）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableCoordPoint []*Coord `json:"TableCoordPoint,omitempty" name:"TableCoordPoint"`
}

type TableOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *TableOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TableOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TableOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TableOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接
		TextDetections []*TextTable `json:"TextDetections,omitempty" name:"TextDetections"`

		// Base64 编码后的 Excel 数据。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TableOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TableOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableTitle struct {

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type TaxiInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *TaxiInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TaxiInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TaxiInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TaxiInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票代码
		InvoiceNum *string `json:"InvoiceNum,omitempty" name:"InvoiceNum"`

		// 发票号码
		InvoiceCode *string `json:"InvoiceCode,omitempty" name:"InvoiceCode"`

		// 日期
		Date *string `json:"Date,omitempty" name:"Date"`

		// 金额
		Fare *string `json:"Fare,omitempty" name:"Fare"`

		// 上车时间
		GetOnTime *string `json:"GetOnTime,omitempty" name:"GetOnTime"`

		// 下车时间
		GetOffTime *string `json:"GetOffTime,omitempty" name:"GetOffTime"`

		// 里程
		Distance *string `json:"Distance,omitempty" name:"Distance"`

		// 发票所在地
		Location *string `json:"Location,omitempty" name:"Location"`

		// 车牌号
		PlateNumber *string `json:"PlateNumber,omitempty" name:"PlateNumber"`

		// 发票消费类型
		InvoiceType *string `json:"InvoiceType,omitempty" name:"InvoiceType"`

		// 省
	// 注意：此字段可能返回 null，表示取不到有效值。
		Province *string `json:"Province,omitempty" name:"Province"`

		// 市
	// 注意：此字段可能返回 null，表示取不到有效值。
		City *string `json:"City,omitempty" name:"City"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TaxiInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TaxiInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextArithmetic struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 算式运算结果，true-正确   false-错误或非法参数
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 保留字段，暂不支持
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 原图文本行坐标，以四个顶点坐标表示（保留字段，暂不支持）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// 保留字段，暂不支持
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 文本行旋转纠正之后在图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	ItemCoord *ItemCoord `json:"ItemCoord,omitempty" name:"ItemCoord"`

	// 算式题型编号：
	// ‘1’: 加减乘除四则
	// ‘2’: 加减乘除已知结果求运算因子
	// ‘3’: 判断大小
	// ‘4’: 约等于估算
	// ‘5’: 带余数除法
	// ‘6’: 分数四则运算
	// ‘7’: 单位换算
	// ‘8’: 竖式加减法
	// ‘9’: 竖式乘除法
	// ‘10’: 脱式计算
	// ‘11’: 解方程
	ExpressionType *string `json:"ExpressionType,omitempty" name:"ExpressionType"`

	// 错题推荐答案，算式运算结果正确返回为""，算式运算结果错误返回推荐答案 (注：暂不支持多个关系运算符（如1<10<7）、无关系运算符（如frac(1,2)+frac(2,3)）、单位换算（如1元=100角）错题的推荐答案返回)
	Answer *string `json:"Answer,omitempty" name:"Answer"`
}

type TextDetectRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *TextDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TextDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片中是否包含文字。
		HasText *bool `json:"HasText,omitempty" name:"HasText"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextDetection struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// 此字段为扩展字段。
	// GeneralBasicOcr接口返回段落信息Parag，包含ParagNo。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 文本行在旋转纠正之后的图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	ItemPolygon *ItemCoord `json:"ItemPolygon,omitempty" name:"ItemPolygon"`

	// 识别出来的单字信息包括单字（包括单字Character和单字置信度confidence）， 支持识别的接口：GeneralBasicOCR、GeneralAccurateOCR
	Words []*DetectedWords `json:"Words,omitempty" name:"Words"`

	// 单字在原图中的四点坐标， 支持识别的接口：GeneralBasicOCR、GeneralAccurateOCR
	WordCoordPoint []*DetectedWordCoordPoint `json:"WordCoordPoint,omitempty" name:"WordCoordPoint"`
}

type TextDetectionEn struct {

	// 识别出的文本行内容。
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 置信度 0 ~100。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行在原图中的四点坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// 此字段为扩展字段。目前EnglishOCR接口返回内容为空。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 英文单词在原图中的四点坐标。
	WordCoordPoint []*WordCoordPoint `json:"WordCoordPoint,omitempty" name:"WordCoordPoint"`

	// 候选字符集(包含候选字Character以及置信度Confidence)。
	CandWord []*CandWord `json:"CandWord,omitempty" name:"CandWord"`

	// 识别出来的单词信息（包括单词Character和单词置信度confidence）
	Words []*Words `json:"Words,omitempty" name:"Words"`
}

type TextEduPaper struct {

	// 识别出的字段名称（关键字）
	Item *string `json:"Item,omitempty" name:"Item"`

	// 识别出的字段名称对应的值，也就是字段Item对应的字符串结果
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 文本行在旋转纠正之后的图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	Itemcoord *ItemCoord `json:"Itemcoord,omitempty" name:"Itemcoord"`
}

type TextFormula struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`
}

type TextGeneralHandwriting struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 置信度 0 - 100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// 此字段为扩展字段。
	// 能返回文本行的段落信息，例如：{\"Parag\":{\"ParagNo\":2}}，
	// 其中ParagNo为段落行，从1开始。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 字的坐标数组，以四个顶点坐标表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordPolygon []*Polygon `json:"WordPolygon,omitempty" name:"WordPolygon"`
}

type TextTable struct {

	// 单元格左上角的列索引
	ColTl *int64 `json:"ColTl,omitempty" name:"ColTl"`

	// 单元格左上角的行索引
	RowTl *int64 `json:"RowTl,omitempty" name:"RowTl"`

	// 单元格右下角的列索引
	ColBr *int64 `json:"ColBr,omitempty" name:"ColBr"`

	// 单元格右下角的行索引
	RowBr *int64 `json:"RowBr,omitempty" name:"RowBr"`

	// 单元格文字
	Text *string `json:"Text,omitempty" name:"Text"`

	// 单元格类型，包含body（表格主体）、header（表头）、footer（表尾）三种
	Type *string `json:"Type,omitempty" name:"Type"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon"`

	// 此字段为扩展字段
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextVatInvoice struct {

	// 识别出的字段名称（关键字）。支持以下字段的识别：
	// 发票代码、 发票号码、 打印发票代码、 打印发票号码、 开票日期、 购买方识别号、 小写金额、 价税合计(大写)、 销售方识别号、 校验码、 购买方名称、 销售方名称、 税额、 复核、 联次名称、 备注、 联次、 密码区、 开票人、 收款人、 （货物或应税劳务、服务名称）、省、 市、 服务类型、 通行费标志、 是否代开、 是否收购、 合计金额、 是否有公司印章、 发票消费类型、 车船税、 机器编号、 成品油标志、 税率、 合计税额、 （购买方地址、电话）、 （销售方地址、电话）、 单价、 金额、 销售方开户行及账号、 购买方开户行及账号、 规格型号、 发票名称、 单位、 数量、 校验码备选、 校验码后六位备选、发票号码备选、车牌号、类型、通行日期起、通行日期止、发票类型。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 字段在原图中的中的四点坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon *Polygon `json:"Polygon,omitempty" name:"Polygon"`
}

type TextVehicleBack struct {

	// 号牌号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateNo *string `json:"PlateNo,omitempty" name:"PlateNo"`

	// 档案编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNo *string `json:"FileNo,omitempty" name:"FileNo"`

	// 核定人数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowNum *string `json:"AllowNum,omitempty" name:"AllowNum"`

	// 总质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalMass *string `json:"TotalMass,omitempty" name:"TotalMass"`

	// 整备质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurbWeight *string `json:"CurbWeight,omitempty" name:"CurbWeight"`

	// 核定载质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadQuality *string `json:"LoadQuality,omitempty" name:"LoadQuality"`

	// 外廓尺寸
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalSize *string `json:"ExternalSize,omitempty" name:"ExternalSize"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Marks *string `json:"Marks,omitempty" name:"Marks"`

	// 检验记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Record *string `json:"Record,omitempty" name:"Record"`

	// 准牵引总质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalQuasiMass *string `json:"TotalQuasiMass,omitempty" name:"TotalQuasiMass"`
}

type TextVehicleFront struct {

	// 号牌号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateNo *string `json:"PlateNo,omitempty" name:"PlateNo"`

	// 车辆类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VehicleType *string `json:"VehicleType,omitempty" name:"VehicleType"`

	// 所有人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 住址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 使用性质
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseCharacter *string `json:"UseCharacter,omitempty" name:"UseCharacter"`

	// 品牌型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitempty" name:"Model"`

	// 车辆识别代号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vin *string `json:"Vin,omitempty" name:"Vin"`

	// 发动机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNo *string `json:"EngineNo,omitempty" name:"EngineNo"`

	// 注册日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterDate *string `json:"RegisterDate,omitempty" name:"RegisterDate"`

	// 发证日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`

	// 印章
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seal *string `json:"Seal,omitempty" name:"Seal"`
}

type TextWaybill struct {

	// 收件人姓名
	RecName *WaybillObj `json:"RecName,omitempty" name:"RecName"`

	// 收件人手机号
	RecNum *WaybillObj `json:"RecNum,omitempty" name:"RecNum"`

	// 收件人地址
	RecAddr *WaybillObj `json:"RecAddr,omitempty" name:"RecAddr"`

	// 寄件人姓名
	SenderName *WaybillObj `json:"SenderName,omitempty" name:"SenderName"`

	// 寄件人手机号
	SenderNum *WaybillObj `json:"SenderNum,omitempty" name:"SenderNum"`

	// 寄件人地址
	SenderAddr *WaybillObj `json:"SenderAddr,omitempty" name:"SenderAddr"`

	// 运单号
	WaybillNum *WaybillObj `json:"WaybillNum,omitempty" name:"WaybillNum"`
}

type TollInvoiceInfo struct {

	// 识别出的字段名称（关键字）。支持以下字段的识别：
	// 发票代码、发票号码、日期、金额、入口、出口、时间、发票消费类型、高速标志。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type TollInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *TollInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TollInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TollInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TollInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 过路过桥费发票识别结果，具体内容请点击左侧链接。
		TollInvoiceInfos []*TollInvoiceInfo `json:"TollInvoiceInfos,omitempty" name:"TollInvoiceInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TollInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TollInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TrainTicketOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *TrainTicketOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrainTicketOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrainTicketOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TrainTicketOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编号
		TicketNum *string `json:"TicketNum,omitempty" name:"TicketNum"`

		// 出发站
		StartStation *string `json:"StartStation,omitempty" name:"StartStation"`

		// 到达站
		DestinationStation *string `json:"DestinationStation,omitempty" name:"DestinationStation"`

		// 出发时间
		Date *string `json:"Date,omitempty" name:"Date"`

		// 车次
		TrainNum *string `json:"TrainNum,omitempty" name:"TrainNum"`

		// 座位号
		Seat *string `json:"Seat,omitempty" name:"Seat"`

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 票价
		Price *string `json:"Price,omitempty" name:"Price"`

		// 席别
		SeatCategory *string `json:"SeatCategory,omitempty" name:"SeatCategory"`

		// 身份证号
		ID *string `json:"ID,omitempty" name:"ID"`

		// 发票消费类型：交通
		InvoiceType *string `json:"InvoiceType,omitempty" name:"InvoiceType"`

		// 序列号
		SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`

		// 加收票价
		AdditionalCost *string `json:"AdditionalCost,omitempty" name:"AdditionalCost"`

		// 手续费
		HandlingFee *string `json:"HandlingFee,omitempty" name:"HandlingFee"`

		// 大写金额（票面有大写金额该字段才有值）
		LegalAmount *string `json:"LegalAmount,omitempty" name:"LegalAmount"`

		// 售票站
		TicketStation *string `json:"TicketStation,omitempty" name:"TicketStation"`

		// 原票价（一般有手续费的才有原始票价字段）
		OriginalPrice *string `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 发票类型：火车票、火车票补票、火车票退票凭证
		InvoiceStyle *string `json:"InvoiceStyle,omitempty" name:"InvoiceStyle"`

		// 收据号码
		ReceiptNumber *string `json:"ReceiptNumber,omitempty" name:"ReceiptNumber"`

		// 仅供报销使用：1为是，0为否
		IsReceipt *string `json:"IsReceipt,omitempty" name:"IsReceipt"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrainTicketOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TrainTicketOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsedVehicleInvoiceInfo struct {

	// 所属税局
	TaxBureau *string `json:"TaxBureau,omitempty" name:"TaxBureau"`

	// 买方单位/个人
	Buyer *string `json:"Buyer,omitempty" name:"Buyer"`

	// 买方单位代码/身份证号码
	BuyerNo *string `json:"BuyerNo,omitempty" name:"BuyerNo"`

	// 买方单位/个人地址
	BuyerAddress *string `json:"BuyerAddress,omitempty" name:"BuyerAddress"`

	// 买方单位电话
	BuyerTel *string `json:"BuyerTel,omitempty" name:"BuyerTel"`

	// 卖方单位/个人
	Seller *string `json:"Seller,omitempty" name:"Seller"`

	// 卖方单位代码/身份证号码
	SellerNo *string `json:"SellerNo,omitempty" name:"SellerNo"`

	// 卖方单位/个人地址
	SellerAddress *string `json:"SellerAddress,omitempty" name:"SellerAddress"`

	// 卖方单位电话
	SellerTel *string `json:"SellerTel,omitempty" name:"SellerTel"`

	// 车牌照号
	VehicleLicenseNo *string `json:"VehicleLicenseNo,omitempty" name:"VehicleLicenseNo"`

	// 登记证号
	RegisterNo *string `json:"RegisterNo,omitempty" name:"RegisterNo"`

	// 车架号/车辆识别代码
	VehicleIdentifyNo *string `json:"VehicleIdentifyNo,omitempty" name:"VehicleIdentifyNo"`

	// 转入地车辆管理所名称
	ManagementOffice *string `json:"ManagementOffice,omitempty" name:"ManagementOffice"`

	// 车价合计
	VehicleTotalPrice *string `json:"VehicleTotalPrice,omitempty" name:"VehicleTotalPrice"`

	// 经营、拍卖单位
	Auctioneer *string `json:"Auctioneer,omitempty" name:"Auctioneer"`

	// 经营、拍卖单位地址
	AuctioneerAddress *string `json:"AuctioneerAddress,omitempty" name:"AuctioneerAddress"`

	// 经营、拍卖单位纳税人识别号
	AuctioneerTaxpayerNum *string `json:"AuctioneerTaxpayerNum,omitempty" name:"AuctioneerTaxpayerNum"`

	// 经营、拍卖单位开户银行、账号
	AuctioneerBankAccount *string `json:"AuctioneerBankAccount,omitempty" name:"AuctioneerBankAccount"`

	// 经营、拍卖单位电话
	AuctioneerTel *string `json:"AuctioneerTel,omitempty" name:"AuctioneerTel"`

	// 二手车市场
	Market *string `json:"Market,omitempty" name:"Market"`

	// 二手车市场纳税人识别号
	MarketTaxpayerNum *string `json:"MarketTaxpayerNum,omitempty" name:"MarketTaxpayerNum"`

	// 二手车市场地址
	MarketAddress *string `json:"MarketAddress,omitempty" name:"MarketAddress"`

	// 二手车市场开户银行账号
	MarketBankAccount *string `json:"MarketBankAccount,omitempty" name:"MarketBankAccount"`

	// 二手车市场电话
	MarketTel *string `json:"MarketTel,omitempty" name:"MarketTel"`
}

type VatInvoice struct {

	// 发票代码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitempty" name:"Number"`

	// 开票日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 购方抬头
	BuyerName *string `json:"BuyerName,omitempty" name:"BuyerName"`

	// 购方税号
	BuyerTaxCode *string `json:"BuyerTaxCode,omitempty" name:"BuyerTaxCode"`

	// 购方地址电话
	BuyerAddressPhone *string `json:"BuyerAddressPhone,omitempty" name:"BuyerAddressPhone"`

	// 购方银行账号
	BuyerBankAccount *string `json:"BuyerBankAccount,omitempty" name:"BuyerBankAccount"`

	// 销方名称
	SellerName *string `json:"SellerName,omitempty" name:"SellerName"`

	// 销方税号
	SellerTaxCode *string `json:"SellerTaxCode,omitempty" name:"SellerTaxCode"`

	// 销方地址电话
	SellerAddressPhone *string `json:"SellerAddressPhone,omitempty" name:"SellerAddressPhone"`

	// 销方银行账号
	SellerBankAccount *string `json:"SellerBankAccount,omitempty" name:"SellerBankAccount"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 机器编码
	MachineNo *string `json:"MachineNo,omitempty" name:"MachineNo"`

	// 发票类型
	// 01：专用发票 
	// 02：货运发票
	// 03：机动车发票 
	// 04：普通发票 
	// 10：电子发票 
	// 11：卷式发票 
	// 14：通行费发票 
	// 15：二手车发票
	Type *string `json:"Type,omitempty" name:"Type"`

	// 检验码
	CheckCode *string `json:"CheckCode,omitempty" name:"CheckCode"`

	// 是否作废（红冲）是否作废（红冲）
	// Y: 已作废 N：未作废 H：红冲
	IsAbandoned *string `json:"IsAbandoned,omitempty" name:"IsAbandoned"`

	// 是否有销货清单 
	// Y: 有清单 N：无清单 
	// 卷票无
	HasSellerList *string `json:"HasSellerList,omitempty" name:"HasSellerList"`

	// 销货清单标题
	SellerListTitle *string `json:"SellerListTitle,omitempty" name:"SellerListTitle"`

	// 销货清单税额
	SellerListTax *string `json:"SellerListTax,omitempty" name:"SellerListTax"`

	// 不含税金额
	AmountWithoutTax *string `json:"AmountWithoutTax,omitempty" name:"AmountWithoutTax"`

	// 税额
	TaxAmount *string `json:"TaxAmount,omitempty" name:"TaxAmount"`

	// 含税金额
	AmountWithTax *string `json:"AmountWithTax,omitempty" name:"AmountWithTax"`

	// 项目明细
	Items []*VatInvoiceItem `json:"Items,omitempty" name:"Items"`
}

type VatInvoiceGoodsInfo struct {

	// 项目名称
	Item *string `json:"Item,omitempty" name:"Item"`

	// 规格型号
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 单位
	MeasurementDimension *string `json:"MeasurementDimension,omitempty" name:"MeasurementDimension"`

	// 价格
	Price *string `json:"Price,omitempty" name:"Price"`

	// 数量
	Quantity *string `json:"Quantity,omitempty" name:"Quantity"`

	// 金额
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 税率(如6%、免税)
	TaxScheme *string `json:"TaxScheme,omitempty" name:"TaxScheme"`

	// 税额
	TaxAmount *string `json:"TaxAmount,omitempty" name:"TaxAmount"`
}

type VatInvoiceItem struct {

	// 行号
	LineNo *string `json:"LineNo,omitempty" name:"LineNo"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规格
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 数量
	Quantity *string `json:"Quantity,omitempty" name:"Quantity"`

	// 单价
	UnitPrice *string `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 不含税金额
	AmountWithoutTax *string `json:"AmountWithoutTax,omitempty" name:"AmountWithoutTax"`

	// 税率
	TaxRate *string `json:"TaxRate,omitempty" name:"TaxRate"`

	// 税额
	TaxAmount *string `json:"TaxAmount,omitempty" name:"TaxAmount"`
}

type VatInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片/PDF的 Base64 值。
	// 支持的文件格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片/PDF大小：所下载文件经Base64编码后不超过 7M。文件下载时间不超过 3 秒。
	// 输入参数 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 支持的文件格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片/PDF大小：所下载文件经 Base64 编码后不超过 7M。文件下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitempty" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitempty" name:"PdfPageNumber"`
}

func (r *VatInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VatInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VatInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		VatInvoiceInfos []*TextVatInvoice `json:"VatInvoiceInfos,omitempty" name:"VatInvoiceInfos"`

		// 明细条目。VatInvoiceInfos中关于明细项的具体条目。
		Items []*VatInvoiceItem `json:"Items,omitempty" name:"Items"`

		// 默认值为0。如果图片为PDF时，返回PDF的总页数。
		PdfPageSize *int64 `json:"PdfPageSize,omitempty" name:"PdfPageSize"`

		// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VatInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VatInvoiceUserInfo struct {

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 纳税人识别号
	TaxId *string `json:"TaxId,omitempty" name:"TaxId"`

	// 地 址、电 话
	AddrTel *string `json:"AddrTel,omitempty" name:"AddrTel"`

	// 开户行及账号
	FinancialAccount *string `json:"FinancialAccount,omitempty" name:"FinancialAccount"`
}

type VatInvoiceVerifyRequest struct {
	*tchttp.BaseRequest

	// 发票代码， 一张发票一天只能查询5次。
	InvoiceCode *string `json:"InvoiceCode,omitempty" name:"InvoiceCode"`

	// 发票号码（8位）。
	InvoiceNo *string `json:"InvoiceNo,omitempty" name:"InvoiceNo"`

	// 开票日期（不支持当天发票查询，只支持一年以内），如：2019-12-20。
	InvoiceDate *string `json:"InvoiceDate,omitempty" name:"InvoiceDate"`

	// 根据票种传递对应值，如果报参数错误，请仔细检查每个票种对应的值
	// 
	// 增值税专用发票：开具金额（不含税）
	// 
	// 增值税普通发票、增值税电子普通发票（含通行费发票）、增值税普通发票（卷票）：校验码后6位
	// 
	// 区块链发票：不含税金额/校验码，例如：“285.01/856ab”
	// 
	// 机动车销售统一发票：不含税价
	// 
	// 货物运输业增值税专用发票：合计金额
	// 
	// 二手车销售统一发票：车价合计
	Additional *string `json:"Additional,omitempty" name:"Additional"`
}

func (r *VatInvoiceVerifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatInvoiceVerifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoiceCode")
	delete(f, "InvoiceNo")
	delete(f, "InvoiceDate")
	delete(f, "Additional")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VatInvoiceVerifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VatInvoiceVerifyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 增值税发票信息，详情请点击左侧链接。
		Invoice *VatInvoice `json:"Invoice,omitempty" name:"Invoice"`

		// 机动车销售统一发票信息
		VehicleInvoiceInfo *VehicleInvoiceInfo `json:"VehicleInvoiceInfo,omitempty" name:"VehicleInvoiceInfo"`

		// 二手车销售统一发票信息
		UsedVehicleInvoiceInfo *UsedVehicleInvoiceInfo `json:"UsedVehicleInvoiceInfo,omitempty" name:"UsedVehicleInvoiceInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VatInvoiceVerifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatInvoiceVerifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VatRollInvoiceInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、合计金额(小写)、合计金额(大写)、开票日期、发票号码、购买方识别号、销售方识别号、校验码、销售方名称、购买方名称、发票消费类型、省、市、是否有公司印章、服务类型、品名、种类。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`
}

type VatRollInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *VatRollInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatRollInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VatRollInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VatRollInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 增值税发票（卷票）识别结果，具体内容请点击左侧链接。
		VatRollInvoiceInfos []*VatRollInvoiceInfo `json:"VatRollInvoiceInfos,omitempty" name:"VatRollInvoiceInfos"`

		// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
		Angle *float64 `json:"Angle,omitempty" name:"Angle"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VatRollInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatRollInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VehicleInvoiceInfo struct {

	// 车辆类型
	CarType *string `json:"CarType,omitempty" name:"CarType"`

	// 厂牌型号
	PlateModel *string `json:"PlateModel,omitempty" name:"PlateModel"`

	// 产地
	ProduceAddress *string `json:"ProduceAddress,omitempty" name:"ProduceAddress"`

	// 合格证号
	CertificateNo *string `json:"CertificateNo,omitempty" name:"CertificateNo"`

	// 进口证明书号
	ImportNo *string `json:"ImportNo,omitempty" name:"ImportNo"`

	// LSVCA2NP9HN0xxxxx
	VinNo *string `json:"VinNo,omitempty" name:"VinNo"`

	// 完税证书号
	PayTaxesNo *string `json:"PayTaxesNo,omitempty" name:"PayTaxesNo"`

	// 吨位
	Tonnage *string `json:"Tonnage,omitempty" name:"Tonnage"`

	// 限乘人数
	LimitCount *string `json:"LimitCount,omitempty" name:"LimitCount"`

	// 发动机号码
	EngineNo *string `json:"EngineNo,omitempty" name:"EngineNo"`

	// 商检单号
	BizCheckFormNo *string `json:"BizCheckFormNo,omitempty" name:"BizCheckFormNo"`

	// 主管税务机关代码
	TaxtationOrgCode *string `json:"TaxtationOrgCode,omitempty" name:"TaxtationOrgCode"`

	// 主管税务机关名称
	TaxtationOrgName *string `json:"TaxtationOrgName,omitempty" name:"TaxtationOrgName"`

	// 税率
	MotorTaxRate *string `json:"MotorTaxRate,omitempty" name:"MotorTaxRate"`

	// 开户行
	MotorBankName *string `json:"MotorBankName,omitempty" name:"MotorBankName"`

	// 账号
	MotorBankAccount *string `json:"MotorBankAccount,omitempty" name:"MotorBankAccount"`
}

type VehicleLicenseOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// FRONT 为行驶证主页正面（有红色印章的一面），
	// BACK 为行驶证副页正面（有号码号牌的一面）。
	// 默认值为：FRONT。
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`
}

func (r *VehicleLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VehicleLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VehicleLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VehicleLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 行驶证主页正面的识别结果，CardSide 为 FRONT。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontInfo *TextVehicleFront `json:"FrontInfo,omitempty" name:"FrontInfo"`

		// 行驶证副页正面的识别结果，CardSide 为 BACK。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BackInfo *TextVehicleBack `json:"BackInfo,omitempty" name:"BackInfo"`

		// Code 告警码列表和释义：
	// -9102 复印件告警
	// -9103 翻拍件告警
	// -9106 ps告警
	// 注：告警码可以同时存在多个
		RecognizeWarnCode []*int64 `json:"RecognizeWarnCode,omitempty" name:"RecognizeWarnCode"`

		// 告警码说明：
	// WARN_DRIVER_LICENSE_COPY_CARD 复印件告警
	// WARN_DRIVER_LICENSE_SCREENED_CARD 翻拍件告警
	// WARN_DRIVER_LICENSE_PS_CARD ps告警
	// 注：告警信息可以同时存在多个
		RecognizeWarnMsg []*string `json:"RecognizeWarnMsg,omitempty" name:"RecognizeWarnMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VehicleLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VehicleLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VehicleRegCertInfo struct {

	// 识别出的字段名称(关键字)，支持以下字段：
	// 【注册登记页】
	// 车辆型号、车辆识别代号/车架号、发动机号、制造厂名称、轴距、轮胎数、总质量、外廓尺寸、轴数、车辆出厂日期、发证日期、使用性质、车辆获得方式、车辆类型、国产/进口、燃料种类、车身颜色、发动机型号、车辆品牌、编号、转向形式、
	// 机动车所有人1、身份证明名称1、号码1、登记机关1、登记日期1
	// 机动车所有人2、身份证明名称2、号码2、登记机关2、登记日期2
	// 机动车所有人3、身份证明名称3、号码3、登记机关3、登记日期3
	// 机动车所有人4、身份证明名称4、号码4、登记机关4、登记日期4
	// 机动车所有人5、身份证明名称5、号码5、登记机关5、登记日期5
	// 机动车所有人6、身份证明名称6、号码6、登记机关6、登记日期6
	// 机动车所有人7、身份证明名称7、号码7、登记机关7、登记日期7
	// 【抵押登记页】
	// 机动车登记证书编号、身份证明名称/号码、抵押权人姓名/名称、抵押登记日期。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type VehicleRegCertOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *VehicleRegCertOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VehicleRegCertOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VehicleRegCertOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VehicleRegCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机动车登记证书识别结果，具体内容请点击左侧链接。
		VehicleRegCertInfos []*VehicleRegCertInfo `json:"VehicleRegCertInfos,omitempty" name:"VehicleRegCertInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VehicleRegCertOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VehicleRegCertOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyBasicBizLicenseRequest struct {
	*tchttp.BaseRequest

	// 用于入参是营业执照图片的场景，ImageBase64和ImageUrl必须提供一个，如果都提供，只使用 ImageUrl。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 用于入参是营业执照图片的场景，ImageBase64和ImageUrl必须提供一个，如果都提供，只使用 ImageUrl。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 用于入参是营业执照图片的场景，表示需要校验的参数：RegNum（注册号或者统一社会信用代码），Name（企业名称），Address（经营地址）。选择后会返回相关参数校验结果。RegNum为必选，Name和Address可选。
	// 格式为{RegNum: true, Name:true/false, Address:true/false}
	// 
	// 设置方式参考：
	// Config = Json.stringify({"Name":true,"Address":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"Name":true,"Address":true}
	ImageConfig *string `json:"ImageConfig,omitempty" name:"ImageConfig"`

	// 用于入参是文本的场景，RegNum表示注册号或者统一社会信用代码。若没有传入营业执照图片，则RegNum为必选项，若图片和RegNum都传入，则只使用RegNum。
	RegNum *string `json:"RegNum,omitempty" name:"RegNum"`

	// 用于入参是文本的场景，Name表示企业名称。Name为可选项，填写后会返回Name的校验结果。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用于入参是文本的场景，Address表示经营地址。Address为可选项，填写后会返回Address的校验结果。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 1表示输出注册资本字段（单位：万元），其他值表示不输出。默认不输出。
	RegCapital *int64 `json:"RegCapital,omitempty" name:"RegCapital"`

	// true表示展示成立/注册日期
	EstablishTime *bool `json:"EstablishTime,omitempty" name:"EstablishTime"`
}

func (r *VerifyBasicBizLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyBasicBizLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ImageConfig")
	delete(f, "RegNum")
	delete(f, "Name")
	delete(f, "Address")
	delete(f, "RegCapital")
	delete(f, "EstablishTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyBasicBizLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VerifyBasicBizLicenseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态码，成功时返回0
		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// 统一社会信用代码
		CreditCode *string `json:"CreditCode,omitempty" name:"CreditCode"`

		// 经营期限自（YYYY-MM-DD）
		Opfrom *string `json:"Opfrom,omitempty" name:"Opfrom"`

		// 经营期限至（YYYY-MM-DD）
		Opto *string `json:"Opto,omitempty" name:"Opto"`

		// 法人姓名
		Frname *string `json:"Frname,omitempty" name:"Frname"`

		// 经营状态，包括：成立、筹建、存续、在营、开业、在册、正常经营、开业登记中、登记成立、撤销、撤销登记、非正常户、告解、个体暂时吊销、个体转企业、吊销（未注销）、拟注销、已注销、（待）迁入、（待）迁出、停业、歇业、清算等。
		Entstatus *string `json:"Entstatus,omitempty" name:"Entstatus"`

		// 经营业务范围
		Zsopscope *string `json:"Zsopscope,omitempty" name:"Zsopscope"`

		// 查询的状态信息
		Reason *string `json:"Reason,omitempty" name:"Reason"`

		// 原注册号
		Oriregno *string `json:"Oriregno,omitempty" name:"Oriregno"`

		// 要核验的工商注册号
		VerifyRegno *string `json:"VerifyRegno,omitempty" name:"VerifyRegno"`

		// 工商注册号
		Regno *string `json:"Regno,omitempty" name:"Regno"`

		// 要核验的企业名称
		VerifyEntname *string `json:"VerifyEntname,omitempty" name:"VerifyEntname"`

		// 企业名称
		Entname *string `json:"Entname,omitempty" name:"Entname"`

		// 要核验的住址
		VerifyDom *string `json:"VerifyDom,omitempty" name:"VerifyDom"`

		// 住址
		Dom *string `json:"Dom,omitempty" name:"Dom"`

		// 验证结果
		RegNumResult *BizLicenseVerifyResult `json:"RegNumResult,omitempty" name:"RegNumResult"`

		// 注册资本（单位：万元）,只有输入参数regCapital为1的时候才输出
		RegCapital *string `json:"RegCapital,omitempty" name:"RegCapital"`

		// 成立/注册日期，只有输入参数EstablishTime为true时展示，默认为空
		EstablishTime *string `json:"EstablishTime,omitempty" name:"EstablishTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyBasicBizLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyBasicBizLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyBizLicenseRequest struct {
	*tchttp.BaseRequest

	// 用于入参是营业执照图片的场景，ImageBase64和ImageUrl必须提供一个，如果都提供，只使用 ImageUrl。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 用于入参是营业执照图片的场景，ImageBase64和ImageUrl必须提供一个，如果都提供，只使用 ImageUrl。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 用于入参是营业执照图片的场景，表示需要校验的参数：RegNum（注册号或者统一社会信用代码），Name（企业名称），Address（经营地址）。选择后会返回相关参数校验结果。RegNum为必选，Name和Address可选。
	// 格式为{RegNum: true, Name:true/false, Address:true/false}
	// 
	// 设置方式参考：
	// Config = Json.stringify({"Name":true,"Address":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"Name":true,"Address":true}
	ImageConfig *string `json:"ImageConfig,omitempty" name:"ImageConfig"`

	// 用于入参是文本的场景，RegNum表示注册号或者统一社会信用代码。若没有传入营业执照图片，则RegNum为必选项，若图片和RegNum都传入，则只使用RegNum。
	RegNum *string `json:"RegNum,omitempty" name:"RegNum"`

	// 用于入参是文本的场景，Name表示企业名称。Name为可选项，填写后会返回Name的校验结果。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用于入参是文本的场景，Address表示经营地址，填写后会返回Address的校验结果。
	Address *string `json:"Address,omitempty" name:"Address"`
}

func (r *VerifyBizLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyBizLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ImageConfig")
	delete(f, "RegNum")
	delete(f, "Name")
	delete(f, "Address")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyBizLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VerifyBizLicenseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态码
		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

		// 统一社会信用代码
		CreditCode *string `json:"CreditCode,omitempty" name:"CreditCode"`

		// 组织机构代码
		OrgCode *string `json:"OrgCode,omitempty" name:"OrgCode"`

		// 经营期限自（YYYY-MM-DD）
		OpenFrom *string `json:"OpenFrom,omitempty" name:"OpenFrom"`

		// 经营期限至（YYYY-MM-DD）
		OpenTo *string `json:"OpenTo,omitempty" name:"OpenTo"`

		// 法人姓名
		FrName *string `json:"FrName,omitempty" name:"FrName"`

		// 经营状态，包括：成立、筹建、存续、在营、开业、在册、正常经营、开业登记中、登记成立、撤销、撤销登记、非正常户、告解、个体暂时吊销、个体转企业、吊销（未注销）、拟注销、已注销、（待）迁入、（待）迁出、停业、歇业、清算等。
		EnterpriseStatus *string `json:"EnterpriseStatus,omitempty" name:"EnterpriseStatus"`

		// 经营（业务）范围及方式
		OperateScopeAndForm *string `json:"OperateScopeAndForm,omitempty" name:"OperateScopeAndForm"`

		// 注册资金（单位:万元）
		RegCap *string `json:"RegCap,omitempty" name:"RegCap"`

		// 注册币种
		RegCapCur *string `json:"RegCapCur,omitempty" name:"RegCapCur"`

		// 登记机关
		RegOrg *string `json:"RegOrg,omitempty" name:"RegOrg"`

		// 开业日期（YYYY-MM-DD）
		EsDate *string `json:"EsDate,omitempty" name:"EsDate"`

		// 企业（机构）类型
		EnterpriseType *string `json:"EnterpriseType,omitempty" name:"EnterpriseType"`

		// 注销日期
		CancelDate *string `json:"CancelDate,omitempty" name:"CancelDate"`

		// 吊销日期
		RevokeDate *string `json:"RevokeDate,omitempty" name:"RevokeDate"`

		// 许可经营项目
		AbuItem *string `json:"AbuItem,omitempty" name:"AbuItem"`

		// 一般经营项目
		CbuItem *string `json:"CbuItem,omitempty" name:"CbuItem"`

		// 核准时间
		ApprDate *string `json:"ApprDate,omitempty" name:"ApprDate"`

		// 省（返回空值）
		Province *string `json:"Province,omitempty" name:"Province"`

		// 地级市（返回空值）
		City *string `json:"City,omitempty" name:"City"`

		// 区\县（返回空值）
		County *string `json:"County,omitempty" name:"County"`

		// 住所所在行政区划代码（返回空值）
		AreaCode *string `json:"AreaCode,omitempty" name:"AreaCode"`

		// 行业门类代码（返回空值）
		IndustryPhyCode *string `json:"IndustryPhyCode,omitempty" name:"IndustryPhyCode"`

		// 行业门类名称（返回空值）
		IndustryPhyName *string `json:"IndustryPhyName,omitempty" name:"IndustryPhyName"`

		// 国民经济行业代码（返回空值）
		IndustryCode *string `json:"IndustryCode,omitempty" name:"IndustryCode"`

		// 国民经济行业名称（返回空值）
		IndustryName *string `json:"IndustryName,omitempty" name:"IndustryName"`

		// 经营（业务）范围
		OperateScope *string `json:"OperateScope,omitempty" name:"OperateScope"`

		// 要核验的工商注册号
		VerifyRegNo *string `json:"VerifyRegNo,omitempty" name:"VerifyRegNo"`

		// 工商注册号
		RegNo *string `json:"RegNo,omitempty" name:"RegNo"`

		// 要核验的企业名称
		VerifyEnterpriseName *string `json:"VerifyEnterpriseName,omitempty" name:"VerifyEnterpriseName"`

		// 企业名称
		EnterpriseName *string `json:"EnterpriseName,omitempty" name:"EnterpriseName"`

		// 要核验的注册地址
		VerifyAddress *string `json:"VerifyAddress,omitempty" name:"VerifyAddress"`

		// 注册地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 验证结果
		RegNumResult *BizLicenseVerifyResult `json:"RegNumResult,omitempty" name:"RegNumResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyBizLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyBizLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyEnterpriseFourFactorsRequest struct {
	*tchttp.BaseRequest

	// 姓名
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 证件号码（公司注册证件号）
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 企业全称
	EnterpriseName *string `json:"EnterpriseName,omitempty" name:"EnterpriseName"`

	// 企业标识（注册号，统一社会信用代码）
	EnterpriseMark *string `json:"EnterpriseMark,omitempty" name:"EnterpriseMark"`
}

func (r *VerifyEnterpriseFourFactorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyEnterpriseFourFactorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RealName")
	delete(f, "IdCard")
	delete(f, "EnterpriseName")
	delete(f, "EnterpriseMark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyEnterpriseFourFactorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VerifyEnterpriseFourFactorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 核验一致性（1:一致，2:不一致，3:查询无记录）
		State *int64 `json:"State,omitempty" name:"State"`

		// 核验结果明细，7：企业法人/负责人，6：企业股东，5：企
	// 业管理人员，-21：企业名称与企业标识不符，-22：姓名不一致，-23：证件号码不一致，-24：企业名称不一致，-25：企业标识不一致
	// 注意：此字段可能返回 null，表示取不到有效值。
		Detail *Detail `json:"Detail,omitempty" name:"Detail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyEnterpriseFourFactorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyEnterpriseFourFactorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyOfdVatInvoiceOCRRequest struct {
	*tchttp.BaseRequest

	// OFD文件的 Url 地址。
	OfdFileUrl *string `json:"OfdFileUrl,omitempty" name:"OfdFileUrl"`

	// OFD文件的 Base64 值。
	// OfdFileUrl 和 OfdFileBase64 必传其一，若两者都传，只解析OfdFileBase64。
	OfdFileBase64 *string `json:"OfdFileBase64,omitempty" name:"OfdFileBase64"`
}

func (r *VerifyOfdVatInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyOfdVatInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OfdFileUrl")
	delete(f, "OfdFileBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyOfdVatInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VerifyOfdVatInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票类型
	// 026:增值税电子普通发票
	// 028:增值税电子专用发票
		Type *string `json:"Type,omitempty" name:"Type"`

		// 发票代码
		InvoiceCode *string `json:"InvoiceCode,omitempty" name:"InvoiceCode"`

		// 发票号码
		InvoiceNumber *string `json:"InvoiceNumber,omitempty" name:"InvoiceNumber"`

		// 开票日期
		IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`

		// 验证码
		InvoiceCheckCode *string `json:"InvoiceCheckCode,omitempty" name:"InvoiceCheckCode"`

		// 机器编号
		MachineNumber *string `json:"MachineNumber,omitempty" name:"MachineNumber"`

		// 密码区
		TaxControlCode *string `json:"TaxControlCode,omitempty" name:"TaxControlCode"`

		// 购买方
		Buyer *VatInvoiceUserInfo `json:"Buyer,omitempty" name:"Buyer"`

		// 销售方
		Seller *VatInvoiceUserInfo `json:"Seller,omitempty" name:"Seller"`

		// 价税合计
		TaxInclusiveTotalAmount *string `json:"TaxInclusiveTotalAmount,omitempty" name:"TaxInclusiveTotalAmount"`

		// 开票人
		InvoiceClerk *string `json:"InvoiceClerk,omitempty" name:"InvoiceClerk"`

		// 收款人
		Payee *string `json:"Payee,omitempty" name:"Payee"`

		// 复核人
		Checker *string `json:"Checker,omitempty" name:"Checker"`

		// 税额
		TaxTotalAmount *string `json:"TaxTotalAmount,omitempty" name:"TaxTotalAmount"`

		// 不含税金额
		TaxExclusiveTotalAmount *string `json:"TaxExclusiveTotalAmount,omitempty" name:"TaxExclusiveTotalAmount"`

		// 备注
		Note *string `json:"Note,omitempty" name:"Note"`

		// 货物或服务清单
		GoodsInfos []*VatInvoiceGoodsInfo `json:"GoodsInfos,omitempty" name:"GoodsInfos"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyOfdVatInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyOfdVatInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VinOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *VinOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VinOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VinOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type VinOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的车辆 VIN 码。
		Vin *string `json:"Vin,omitempty" name:"Vin"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VinOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VinOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WaybillOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 预检测开关，当待识别运单占整个输入图像的比例较小时，建议打开预检测开关。默认值为false。
	EnablePreDetect *bool `json:"EnablePreDetect,omitempty" name:"EnablePreDetect"`
}

func (r *WaybillOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WaybillOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "EnablePreDetect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WaybillOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type WaybillOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections *TextWaybill `json:"TextDetections,omitempty" name:"TextDetections"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WaybillOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WaybillOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WaybillObj struct {

	// 识别出的文本行内容
	Text *string `json:"Text,omitempty" name:"Text"`
}

type WordCoordPoint struct {

	// 英文OCR识别出的每个单词在原图中的四点坐标。
	WordCoordinate []*Coord `json:"WordCoordinate,omitempty" name:"WordCoordinate"`
}

type Words struct {

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 候选字Character
	Character *string `json:"Character,omitempty" name:"Character"`
}
