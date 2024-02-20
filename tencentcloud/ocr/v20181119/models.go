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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AdvertiseOCRRequestParams struct {
	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type AdvertiseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type AdvertiseOCRResponseParams struct {
	// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
	TextDetections []*AdvertiseTextDetection `json:"TextDetections,omitnil" name:"TextDetections"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AdvertiseOCRResponse struct {
	*tchttp.BaseResponse
	Response *AdvertiseOCRResponseParams `json:"Response"`
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
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// 此字段为扩展字段。
	// GeneralBasicOcr接口返回段落信息Parag，包含ParagNo。
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`
}

type AirTicketInfo struct {
	// 旅客姓名
	PassengerName *string `json:"PassengerName,omitnil" name:"PassengerName"`

	// 有效身份证件号码
	ValidIdNumber *string `json:"ValidIdNumber,omitnil" name:"ValidIdNumber"`

	// 签注
	Endorsement *string `json:"Endorsement,omitnil" name:"Endorsement"`

	// GP单号
	NumberOfGPOrder *string `json:"NumberOfGPOrder,omitnil" name:"NumberOfGPOrder"`

	// 发票号码
	ElectronicInvoiceAirTransportReceiptNumber *string `json:"ElectronicInvoiceAirTransportReceiptNumber,omitnil" name:"ElectronicInvoiceAirTransportReceiptNumber"`

	// 机票详细信息元组
	DetailInformationOfAirTicketTuple []*DetailInformationOfAirTicketTupleList `json:"DetailInformationOfAirTicketTuple,omitnil" name:"DetailInformationOfAirTicketTuple"`

	// 票价
	Fare *string `json:"Fare,omitnil" name:"Fare"`

	// 燃油附加费
	FuelSurcharge *string `json:"FuelSurcharge,omitnil" name:"FuelSurcharge"`

	// 增值税税率
	VatRate *string `json:"VatRate,omitnil" name:"VatRate"`

	// 增值税税额
	VatTaxAmount *string `json:"VatTaxAmount,omitnil" name:"VatTaxAmount"`

	// 民航发展基金
	CivilAviationDevelopmentFund *string `json:"CivilAviationDevelopmentFund,omitnil" name:"CivilAviationDevelopmentFund"`

	// 其他税费
	OtherTaxes *string `json:"OtherTaxes,omitnil" name:"OtherTaxes"`

	// 合计
	TotalAmount *string `json:"TotalAmount,omitnil" name:"TotalAmount"`

	// 电子客票号码
	ElectronicTicketNum *string `json:"ElectronicTicketNum,omitnil" name:"ElectronicTicketNum"`

	// 验证码
	VerificationCode *string `json:"VerificationCode,omitnil" name:"VerificationCode"`

	// 提示信息
	PromptInformation *string `json:"PromptInformation,omitnil" name:"PromptInformation"`

	// 保险费
	Insurance *string `json:"Insurance,omitnil" name:"Insurance"`

	// 销售网点代号
	AgentCode *string `json:"AgentCode,omitnil" name:"AgentCode"`

	// 填开单位
	IssueParty *string `json:"IssueParty,omitnil" name:"IssueParty"`

	// 填开时间
	IssueDate *string `json:"IssueDate,omitnil" name:"IssueDate"`

	// 开具状态
	IssuingStatus *string `json:"IssuingStatus,omitnil" name:"IssuingStatus"`

	// 国内国际标识
	MarkingOfDomesticOrInternational *string `json:"MarkingOfDomesticOrInternational,omitnil" name:"MarkingOfDomesticOrInternational"`

	// 购买方名称
	NameOfPurchaser *string `json:"NameOfPurchaser,omitnil" name:"NameOfPurchaser"`

	// 销售方名称
	NameOfSeller *string `json:"NameOfSeller,omitnil" name:"NameOfSeller"`

	// 统一社会信用代码
	UnifiedSocialCreditCodeOfPurchaser *string `json:"UnifiedSocialCreditCodeOfPurchaser,omitnil" name:"UnifiedSocialCreditCodeOfPurchaser"`
}

type AirTransport struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 电子客票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 校验码
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 印刷序号
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 销售单位代号
	AgentCode *string `json:"AgentCode,omitnil" name:"AgentCode"`

	// 销售单位代号第一行
	AgentCodeFirst *string `json:"AgentCodeFirst,omitnil" name:"AgentCodeFirst"`

	// 销售单位代号第二行
	AgentCodeSecond *string `json:"AgentCodeSecond,omitnil" name:"AgentCodeSecond"`

	// 姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 身份证号
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// 填开单位
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 票价
	Fare *string `json:"Fare,omitnil" name:"Fare"`

	// 合计税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 燃油附加费
	FuelSurcharge *string `json:"FuelSurcharge,omitnil" name:"FuelSurcharge"`

	// 民航发展基金
	AirDevelopmentFund *string `json:"AirDevelopmentFund,omitnil" name:"AirDevelopmentFund"`

	// 保险费
	Insurance *string `json:"Insurance,omitnil" name:"Insurance"`

	// 合计金额（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 国内国际标签
	DomesticInternationalTag *string `json:"DomesticInternationalTag,omitnil" name:"DomesticInternationalTag"`

	// 客票生效日期
	DateStart *string `json:"DateStart,omitnil" name:"DateStart"`

	// 有效截至日期
	DateEnd *string `json:"DateEnd,omitnil" name:"DateEnd"`

	// 签注
	Endorsement *string `json:"Endorsement,omitnil" name:"Endorsement"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 条目
	FlightItems []*FlightItem `json:"FlightItems,omitnil" name:"FlightItems"`
}

// Predefined struct for user
type ArithmeticOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 用于选择是否支持横屏拍摄。打开则支持横屏拍摄图片角度判断，角度信息在返回参数的angle中，默认值为true
	SupportHorizontalImage *bool `json:"SupportHorizontalImage,omitnil" name:"SupportHorizontalImage"`

	// 是否拒绝非速算图，打开则拒绝非速算图(注：非速算图是指风景人物等明显不是速算图片的图片)，默认值为false
	RejectNonArithmeticPic *bool `json:"RejectNonArithmeticPic,omitnil" name:"RejectNonArithmeticPic"`

	// 是否展开耦合算式中的竖式计算，默认值为false
	EnableDispRelatedVertical *bool `json:"EnableDispRelatedVertical,omitnil" name:"EnableDispRelatedVertical"`

	// 是否展示竖式算式的中间结果和格式控制字符，默认值为false
	EnableDispMidResult *bool `json:"EnableDispMidResult,omitnil" name:"EnableDispMidResult"`

	// 是否开启pdf识别，默认值为true
	EnablePdfRecognize *bool `json:"EnablePdfRecognize,omitnil" name:"EnablePdfRecognize"`

	// pdf页码，从0开始，默认为0
	PdfPageIndex *int64 `json:"PdfPageIndex,omitnil" name:"PdfPageIndex"`
}

type ArithmeticOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 用于选择是否支持横屏拍摄。打开则支持横屏拍摄图片角度判断，角度信息在返回参数的angle中，默认值为true
	SupportHorizontalImage *bool `json:"SupportHorizontalImage,omitnil" name:"SupportHorizontalImage"`

	// 是否拒绝非速算图，打开则拒绝非速算图(注：非速算图是指风景人物等明显不是速算图片的图片)，默认值为false
	RejectNonArithmeticPic *bool `json:"RejectNonArithmeticPic,omitnil" name:"RejectNonArithmeticPic"`

	// 是否展开耦合算式中的竖式计算，默认值为false
	EnableDispRelatedVertical *bool `json:"EnableDispRelatedVertical,omitnil" name:"EnableDispRelatedVertical"`

	// 是否展示竖式算式的中间结果和格式控制字符，默认值为false
	EnableDispMidResult *bool `json:"EnableDispMidResult,omitnil" name:"EnableDispMidResult"`

	// 是否开启pdf识别，默认值为true
	EnablePdfRecognize *bool `json:"EnablePdfRecognize,omitnil" name:"EnablePdfRecognize"`

	// pdf页码，从0开始，默认为0
	PdfPageIndex *int64 `json:"PdfPageIndex,omitnil" name:"PdfPageIndex"`
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

// Predefined struct for user
type ArithmeticOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	TextDetections []*TextArithmetic `json:"TextDetections,omitnil" name:"TextDetections"`

	// 图片横屏的角度(90度或270度)
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ArithmeticOCRResponse struct {
	*tchttp.BaseResponse
	Response *ArithmeticOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type BankCardOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回预处理（精确剪裁对齐）后的银行卡图片数据，默认false。
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitnil" name:"RetBorderCutImage"`

	// 是否返回卡号的切图图片数据，默认false。
	RetCardNoImage *bool `json:"RetCardNoImage,omitnil" name:"RetCardNoImage"`

	// 复印件检测开关，如果输入的图片是银行卡复印件图片则返回告警，默认false。
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitnil" name:"EnableCopyCheck"`

	// 翻拍检测开关，如果输入的图片是银行卡翻拍图片则返回告警，默认false。
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitnil" name:"EnableReshootCheck"`

	// 边框遮挡检测开关，如果输入的图片是银行卡边框被遮挡则返回告警，默认false。
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitnil" name:"EnableBorderCheck"`

	// 是否返回图片质量分数（图片质量分数是评价一个图片的模糊程度的标准），默认false。
	EnableQualityValue *bool `json:"EnableQualityValue,omitnil" name:"EnableQualityValue"`
}

type BankCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回预处理（精确剪裁对齐）后的银行卡图片数据，默认false。
	RetBorderCutImage *bool `json:"RetBorderCutImage,omitnil" name:"RetBorderCutImage"`

	// 是否返回卡号的切图图片数据，默认false。
	RetCardNoImage *bool `json:"RetCardNoImage,omitnil" name:"RetCardNoImage"`

	// 复印件检测开关，如果输入的图片是银行卡复印件图片则返回告警，默认false。
	EnableCopyCheck *bool `json:"EnableCopyCheck,omitnil" name:"EnableCopyCheck"`

	// 翻拍检测开关，如果输入的图片是银行卡翻拍图片则返回告警，默认false。
	EnableReshootCheck *bool `json:"EnableReshootCheck,omitnil" name:"EnableReshootCheck"`

	// 边框遮挡检测开关，如果输入的图片是银行卡边框被遮挡则返回告警，默认false。
	EnableBorderCheck *bool `json:"EnableBorderCheck,omitnil" name:"EnableBorderCheck"`

	// 是否返回图片质量分数（图片质量分数是评价一个图片的模糊程度的标准），默认false。
	EnableQualityValue *bool `json:"EnableQualityValue,omitnil" name:"EnableQualityValue"`
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

// Predefined struct for user
type BankCardOCRResponseParams struct {
	// 卡号
	CardNo *string `json:"CardNo,omitnil" name:"CardNo"`

	// 银行信息
	BankInfo *string `json:"BankInfo,omitnil" name:"BankInfo"`

	// 有效期，格式如：07/2023
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// 卡类型
	CardType *string `json:"CardType,omitnil" name:"CardType"`

	// 卡名字
	CardName *string `json:"CardName,omitnil" name:"CardName"`

	// 切片图片数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	BorderCutImage *string `json:"BorderCutImage,omitnil" name:"BorderCutImage"`

	// 卡号图片数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CardNoImage *string `json:"CardNoImage,omitnil" name:"CardNoImage"`

	// WarningCode 告警码列表和释义：
	// -9110:银行卡日期无效; 
	// -9111:银行卡边框不完整; 
	// -9112:银行卡图片反光;
	// -9113:银行卡复印件;
	// -9114:银行卡翻拍件
	// （告警码可以同时存在多个）
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarningCode []*int64 `json:"WarningCode,omitnil" name:"WarningCode"`

	// 图片质量分数，请求EnableQualityValue时返回（取值范围：0-100，分数越低越模糊，建议阈值≥50）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityValue *int64 `json:"QualityValue,omitnil" name:"QualityValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BankCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *BankCardOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type BankSlipOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type BankSlipOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BankSlipOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BankSlipOCRResponseParams struct {
	// 银行回单识别结果，具体内容请点击左侧链接。
	BankSlipInfos []*BankSlipInfo `json:"BankSlipInfos,omitnil" name:"BankSlipInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BankSlipOCRResponse struct {
	*tchttp.BaseResponse
	Response *BankSlipOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type BizLicenseOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回告警码，默认为false
	EnableCopyWarn *bool `json:"EnableCopyWarn,omitnil" name:"EnableCopyWarn"`
}

type BizLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回告警码，默认为false
	EnableCopyWarn *bool `json:"EnableCopyWarn,omitnil" name:"EnableCopyWarn"`
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
	delete(f, "EnableCopyWarn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BizLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BizLicenseOCRResponseParams struct {
	// 统一社会信用代码（三合一之前为注册号）
	RegNum *string `json:"RegNum,omitnil" name:"RegNum"`

	// 公司名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 注册资本
	Capital *string `json:"Capital,omitnil" name:"Capital"`

	// 法定代表人
	Person *string `json:"Person,omitnil" name:"Person"`

	// 地址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 经营范围
	Business *string `json:"Business,omitnil" name:"Business"`

	// 主体类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 营业期限
	Period *string `json:"Period,omitnil" name:"Period"`

	// 组成形式
	ComposingForm *string `json:"ComposingForm,omitnil" name:"ComposingForm"`

	// 成立日期
	SetDate *string `json:"SetDate,omitnil" name:"SetDate"`

	// Code 告警码列表和释义：
	// -9102 黑白复印件告警
	// -9104 翻拍件告警
	RecognizeWarnCode []*int64 `json:"RecognizeWarnCode,omitnil" name:"RecognizeWarnCode"`

	// 告警码说明：
	// WARN_COPY_CARD 黑白复印件告警
	// WARN_RESHOOT_CARD翻拍件告警
	RecognizeWarnMsg []*string `json:"RecognizeWarnMsg,omitnil" name:"RecognizeWarnMsg"`

	// 是否为副本。1为是，-1为不是。
	IsDuplication *int64 `json:"IsDuplication,omitnil" name:"IsDuplication"`

	// 登记日期
	RegistrationDate *string `json:"RegistrationDate,omitnil" name:"RegistrationDate"`

	//  图片旋转角度(角度制)，文本的水平方向为0度；顺时针为正，角度范围是0-360度
	// 
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BizLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *BizLicenseOCRResponseParams `json:"Response"`
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

type BusInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 乘车时间
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// 乘车日期
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// 出发车站
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// 到达车站
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// 票价
	Total *string `json:"Total,omitnil" name:"Total"`

	// 姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 身份证号
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 乘车地点
	PlaceGetOn *string `json:"PlaceGetOn,omitnil" name:"PlaceGetOn"`

	// 检票口
	GateNumber *string `json:"GateNumber,omitnil" name:"GateNumber"`

	// 客票类型
	TicketType *string `json:"TicketType,omitnil" name:"TicketType"`

	// 车型
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// 座位号
	SeatNumber *string `json:"SeatNumber,omitnil" name:"SeatNumber"`

	// 车次
	TrainNumber *string `json:"TrainNumber,omitnil" name:"TrainNumber"`
}

type BusInvoiceInfo struct {
	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、发票号码、日期、票价、始发地、目的地、姓名、时间、发票消费类型、身份证号、省、市、开票日期、乘车地点、检票口、客票类型、车型、座位号、车次。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type BusInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type BusInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BusInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BusInvoiceOCRResponseParams struct {
	// 汽车票识别结果，具体内容请点击左侧链接。
	BusInvoiceInfos []*BusInvoiceInfo `json:"BusInvoiceInfos,omitnil" name:"BusInvoiceInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BusInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *BusInvoiceOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	ItemCoord *ItemCoord `json:"ItemCoord,omitnil" name:"ItemCoord"`
}

// Predefined struct for user
type BusinessCardOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 可选字段，根据需要选择是否请求对应字段。
	// 目前支持的字段为：
	// RetImageType-“PROPROCESS” 图像预处理，string 类型。
	// 图像预处理功能为，检测图片倾斜的角度，将原本倾斜的图片围绕中心点转正，最终输出一张正的名片抠图。
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"RetImageType":"PROPROCESS"})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"RetImageType":"PROPROCESS"}
	Config *string `json:"Config,omitnil" name:"Config"`
}

type BusinessCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 可选字段，根据需要选择是否请求对应字段。
	// 目前支持的字段为：
	// RetImageType-“PROPROCESS” 图像预处理，string 类型。
	// 图像预处理功能为，检测图片倾斜的角度，将原本倾斜的图片围绕中心点转正，最终输出一张正的名片抠图。
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"RetImageType":"PROPROCESS"})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"RetImageType":"PROPROCESS"}
	Config *string `json:"Config,omitnil" name:"Config"`
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

// Predefined struct for user
type BusinessCardOCRResponseParams struct {
	// 名片识别结果，具体内容请点击左侧链接。
	BusinessCardInfos []*BusinessCardInfo `json:"BusinessCardInfos,omitnil" name:"BusinessCardInfos"`

	// 返回图像预处理后的图片，图像预处理未开启时返回内容为空。
	RetImageBase64 *string `json:"RetImageBase64,omitnil" name:"RetImageBase64"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BusinessCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *BusinessCardOCRResponseParams `json:"Response"`
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
	CandWords []*Words `json:"CandWords,omitnil" name:"CandWords"`
}

type CarInvoiceInfo struct {
	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、 机打代码、 发票号码、 发动机号码、 合格证号、 机打号码、 价税合计(小写)、 销货单位名称、 身份证号码/组织机构代码、 购买方名称、 销售方纳税人识别号、 购买方纳税人识别号、主管税务机关、 主管税务机关代码、 开票日期、 不含税价(小写)、 吨位、增值税税率或征收率、 车辆识别代号/车架号码、 增值税税额、 厂牌型号、 省、 市、 发票消费类型、 销售方电话、 销售方账号、 产地、 进口证明书号、 车辆类型、 机器编号、备注、开票人、限乘人数、商检单号、销售方地址、销售方开户银行、价税合计、发票类型。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 字段在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`

	// 字段在原图中的四点坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon *Polygon `json:"Polygon,omitnil" name:"Polygon"`
}

// Predefined struct for user
type CarInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type CarInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CarInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CarInvoiceOCRResponseParams struct {
	// 购车发票识别结果，具体内容请点击左侧链接。
	CarInvoiceInfos []*CarInvoiceInfo `json:"CarInvoiceInfos,omitnil" name:"CarInvoiceInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CarInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *CarInvoiceOCRResponseParams `json:"Response"`
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
	ParagNo *int64 `json:"ParagNo,omitnil" name:"ParagNo"`

	// 字体大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordSize *int64 `json:"WordSize,omitnil" name:"WordSize"`
}

type ClassifyDetectInfo struct {
	// 分类名称，包括：身份证、护照、名片、银行卡、行驶证、驾驶证、港澳台通行证、户口本、港澳台来往内地通行证、港澳台居住证、不动产证、营业执照
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分类类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 位置坐标
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type ClassifyDetectOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

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
	DiscernType []*string `json:"DiscernType,omitnil" name:"DiscernType"`
}

type ClassifyDetectOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

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
	DiscernType []*string `json:"DiscernType,omitnil" name:"DiscernType"`
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

// Predefined struct for user
type ClassifyDetectOCRResponseParams struct {
	// 智能卡证分类结果。当图片类型不支持分类识别或者识别出的类型不在请求参数DiscernType指定的范围内时，返回结果中的Type字段将为空字符串，Name字段将返回"其它"
	ClassifyDetectInfos []*ClassifyDetectInfo `json:"ClassifyDetectInfos,omitnil" name:"ClassifyDetectInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ClassifyDetectOCRResponse struct {
	*tchttp.BaseResponse
	Response *ClassifyDetectOCRResponseParams `json:"Response"`
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
	X *int64 `json:"X,omitnil" name:"X"`

	// 纵坐标
	Y *int64 `json:"Y,omitnil" name:"Y"`
}

// Predefined struct for user
type CreateAIFormTaskRequestParams struct {
	// 多个文件的URL列表
	FileList []*SmartFormFileUrl `json:"FileList,omitnil" name:"FileList"`

	// 备注信息1
	FirstNotes *string `json:"FirstNotes,omitnil" name:"FirstNotes"`

	// 备注信息2
	SecondNotes *string `json:"SecondNotes,omitnil" name:"SecondNotes"`

	// 文件类型
	FileType *uint64 `json:"FileType,omitnil" name:"FileType"`
}

type CreateAIFormTaskRequest struct {
	*tchttp.BaseRequest
	
	// 多个文件的URL列表
	FileList []*SmartFormFileUrl `json:"FileList,omitnil" name:"FileList"`

	// 备注信息1
	FirstNotes *string `json:"FirstNotes,omitnil" name:"FirstNotes"`

	// 备注信息2
	SecondNotes *string `json:"SecondNotes,omitnil" name:"SecondNotes"`

	// 文件类型
	FileType *uint64 `json:"FileType,omitnil" name:"FileType"`
}

func (r *CreateAIFormTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIFormTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileList")
	delete(f, "FirstNotes")
	delete(f, "SecondNotes")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIFormTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIFormTaskResponseParams struct {
	// 本次识别任务的唯一身份ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 本次识别任务的操作URL，有效期自生成之时起共24小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUrl *string `json:"OperateUrl,omitnil" name:"OperateUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAIFormTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIFormTaskResponseParams `json:"Response"`
}

func (r *CreateAIFormTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIFormTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailInformationOfAirTicketTupleList struct {
	// 出发站（自）
	DepartureStation *string `json:"DepartureStation,omitnil" name:"DepartureStation"`

	// 目的地（至）
	DestinationStation *string `json:"DestinationStation,omitnil" name:"DestinationStation"`

	// 航班
	FlightSegment *string `json:"FlightSegment,omitnil" name:"FlightSegment"`

	// 航班
	Carrier *string `json:"Carrier,omitnil" name:"Carrier"`

	// 航班号
	Flight *string `json:"Flight,omitnil" name:"Flight"`

	// 座位等级
	SeatClass *string `json:"SeatClass,omitnil" name:"SeatClass"`

	// 日期
	CarrierDate *string `json:"CarrierDate,omitnil" name:"CarrierDate"`

	// 时间
	DepartureTime *string `json:"DepartureTime,omitnil" name:"DepartureTime"`

	// 客票级别/客票类别
	FareBasis *string `json:"FareBasis,omitnil" name:"FareBasis"`

	// 客票生效日期
	EffectiveDate *string `json:"EffectiveDate,omitnil" name:"EffectiveDate"`

	// 有效截止日期
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 免费行李
	FreeBaggageAllowance *string `json:"FreeBaggageAllowance,omitnil" name:"FreeBaggageAllowance"`
}

type DetectedWordCoordPoint struct {
	// 单字在原图中的坐标，以四个顶点坐标表示，以左上角为起点，顺时针返回。
	WordCoordinate []*Coord `json:"WordCoordinate,omitnil" name:"WordCoordinate"`
}

type DetectedWords struct {
	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 候选字Character
	Character *string `json:"Character,omitnil" name:"Character"`
}

// Predefined struct for user
type DriverLicenseOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT 为驾驶证主页正面（有红色印章的一面），
	// BACK 为驾驶证副页正面（有档案编号的一面）。
	// DOUBLE 支持自动识别驾驶证正副页单面，和正副双面同框识别
	// 默认值为：FRONT。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
}

type DriverLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT 为驾驶证主页正面（有红色印章的一面），
	// BACK 为驾驶证副页正面（有档案编号的一面）。
	// DOUBLE 支持自动识别驾驶证正副页单面，和正副双面同框识别
	// 默认值为：FRONT。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
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

// Predefined struct for user
type DriverLicenseOCRResponseParams struct {
	// 驾驶证正页姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 性别
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 国籍
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// 住址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 出生日期（YYYY-MM-DD）
	DateOfBirth *string `json:"DateOfBirth,omitnil" name:"DateOfBirth"`

	// 初次领证日期（YYYY-MM-DD）
	DateOfFirstIssue *string `json:"DateOfFirstIssue,omitnil" name:"DateOfFirstIssue"`

	// 准驾车型
	Class *string `json:"Class,omitnil" name:"Class"`

	// 有效期开始时间（YYYY-MM-DD）
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 有效期截止时间（新版驾驶证返回 YYYY-MM-DD，
	// 老版驾驶证返回有效期限 X年）
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 驾驶证正页证号
	CardCode *string `json:"CardCode,omitnil" name:"CardCode"`

	// 档案编号
	ArchivesCode *string `json:"ArchivesCode,omitnil" name:"ArchivesCode"`

	// 记录
	Record *string `json:"Record,omitnil" name:"Record"`

	// Code 告警码列表和释义：
	// -9102  复印件告警
	// -9103  翻拍件告警
	// 注：告警码可以同时存在多个
	RecognizeWarnCode []*int64 `json:"RecognizeWarnCode,omitnil" name:"RecognizeWarnCode"`

	// 告警码说明：
	// WARN_DRIVER_LICENSE_COPY_CARD 复印件告警
	// WARN_DRIVER_LICENSE_SCREENED_CARD 翻拍件告警
	// 注：告警信息可以同时存在多个
	RecognizeWarnMsg []*string `json:"RecognizeWarnMsg,omitnil" name:"RecognizeWarnMsg"`

	// 发证单位
	IssuingAuthority *string `json:"IssuingAuthority,omitnil" name:"IssuingAuthority"`

	// 状态（仅电子驾驶证支持返回该字段）
	State *string `json:"State,omitnil" name:"State"`

	// 累积记分（仅电子驾驶证支持返回该字段）
	CumulativeScore *string `json:"CumulativeScore,omitnil" name:"CumulativeScore"`

	// 当前时间（仅电子驾驶证支持返回该字段）
	CurrentTime *string `json:"CurrentTime,omitnil" name:"CurrentTime"`

	// 生成时间（仅电子驾驶证支持返回该字段）
	GenerateTime *string `json:"GenerateTime,omitnil" name:"GenerateTime"`

	// 驾驶证副页姓名
	BackPageName *string `json:"BackPageName,omitnil" name:"BackPageName"`

	// 驾驶证副页证号
	BackPageCardCode *string `json:"BackPageCardCode,omitnil" name:"BackPageCardCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DriverLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *DriverLicenseOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type DutyPaidProofOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type DutyPaidProofOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DutyPaidProofOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DutyPaidProofOCRResponseParams struct {
	// 完税证明识别结果，具体内容请点击左侧链接。
	DutyPaidProofInfos []*DutyPaidProofInfo `json:"DutyPaidProofInfos,omitnil" name:"DutyPaidProofInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DutyPaidProofOCRResponse struct {
	*tchttp.BaseResponse
	Response *DutyPaidProofOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type EduPaperOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 扩展配置信息。
	// 配置格式：{"option1":value1,"option2":value2}
	// 1. task_type：任务类型【0: 关闭版式分析与处理 1: 开启版式分析处理】可选参数，Int32类型，默认值为1
	// 2. is_structuralization：是否结构化输出【true：返回包体同时返回通用和结构化输出  false：返回包体返回通用输出】 可选参数，Bool类型，默认值为true
	// 3. if_readable_format：是否按照版式整合通用文本/公式输出结果 可选参数，Bool类型，默认值为false
	// 示例：
	// {"task_type": 1,"is_structuralization": true,"if_readable_format": true}
	Config *string `json:"Config,omitnil" name:"Config"`
}

type EduPaperOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 扩展配置信息。
	// 配置格式：{"option1":value1,"option2":value2}
	// 1. task_type：任务类型【0: 关闭版式分析与处理 1: 开启版式分析处理】可选参数，Int32类型，默认值为1
	// 2. is_structuralization：是否结构化输出【true：返回包体同时返回通用和结构化输出  false：返回包体返回通用输出】 可选参数，Bool类型，默认值为true
	// 3. if_readable_format：是否按照版式整合通用文本/公式输出结果 可选参数，Bool类型，默认值为false
	// 示例：
	// {"task_type": 1,"is_structuralization": true,"if_readable_format": true}
	Config *string `json:"Config,omitnil" name:"Config"`
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

// Predefined struct for user
type EduPaperOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	EduPaperInfos []*TextEduPaper `json:"EduPaperInfos,omitnil" name:"EduPaperInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。
	Angle *int64 `json:"Angle,omitnil" name:"Angle"`

	// 结构化方式输出，具体内容请点击左侧链接。
	QuestionBlockInfos []*QuestionBlockObj `json:"QuestionBlockInfos,omitnil" name:"QuestionBlockInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EduPaperOCRResponse struct {
	*tchttp.BaseResponse
	Response *EduPaperOCRResponseParams `json:"Response"`
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

type ElectronicAirTransport struct {
	// 发票代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitnil" name:"Date"`

	// 金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *string `json:"Amount,omitnil" name:"Amount"`

	// 校验码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 价税合计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *string `json:"Total,omitnil" name:"Total"`

	// 抵扣标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductionMark *string `json:"DeductionMark,omitnil" name:"DeductionMark"`

	// 发票状态代码，0正常 1 未更新  2作废 3已红冲
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateCode *string `json:"StateCode,omitnil" name:"StateCode"`

	// 购方识别号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuyerTaxCode *string `json:"BuyerTaxCode,omitnil" name:"BuyerTaxCode"`

	// 购方名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuyerName *string `json:"BuyerName,omitnil" name:"BuyerName"`

	// 合计税额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 国内国际标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomesticInternationalMark *string `json:"DomesticInternationalMark,omitnil" name:"DomesticInternationalMark"`

	// 旅客姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassengerName *string `json:"PassengerName,omitnil" name:"PassengerName"`

	// 有效身份证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassengerNo *string `json:"PassengerNo,omitnil" name:"PassengerNo"`

	// 电子客票号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElectronicNumber *string `json:"ElectronicNumber,omitnil" name:"ElectronicNumber"`

	// 全电发票（航空运输电子客票行程单）详细信息
	// 
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElectronicAirTransportDetails []*ElectronicAirTransportDetail `json:"ElectronicAirTransportDetails,omitnil" name:"ElectronicAirTransportDetails"`
}

type ElectronicAirTransportDetail struct {
	// 航段序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlightSegment *string `json:"FlightSegment,omitnil" name:"FlightSegment"`

	// 始发站
	// 注意：此字段可能返回 null，表示取不到有效值。
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// 目的站
	// 注意：此字段可能返回 null，表示取不到有效值。
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// 承运人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Carrier *string `json:"Carrier,omitnil" name:"Carrier"`

	// 航班号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlightNumber *string `json:"FlightNumber,omitnil" name:"FlightNumber"`

	// 座位等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	SeatLevel *string `json:"SeatLevel,omitnil" name:"SeatLevel"`

	// 承运日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlightDate *string `json:"FlightDate,omitnil" name:"FlightDate"`

	// 起飞时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartureTime *string `json:"DepartureTime,omitnil" name:"DepartureTime"`

	// 客票级别/客票类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	FareBasis *string `json:"FareBasis,omitnil" name:"FareBasis"`
}

type ElectronicFlightTicketFull struct {
	// 旅客姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 有效身份证件号码
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// 签注
	Endorsement *string `json:"Endorsement,omitnil" name:"Endorsement"`

	// GP单号
	GPOrder *string `json:"GPOrder,omitnil" name:"GPOrder"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 票价
	Fare *string `json:"Fare,omitnil" name:"Fare"`

	// 燃油附加费
	FuelSurcharge *string `json:"FuelSurcharge,omitnil" name:"FuelSurcharge"`

	// 增值税税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 增值税税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 民航发展基金
	DevelopmentFund *string `json:"DevelopmentFund,omitnil" name:"DevelopmentFund"`

	// 其他税费
	OtherTax *string `json:"OtherTax,omitnil" name:"OtherTax"`

	// 合计
	Total *string `json:"Total,omitnil" name:"Total"`

	// 电子客票号码
	ElectronicTicketNum *string `json:"ElectronicTicketNum,omitnil" name:"ElectronicTicketNum"`

	// 验证码
	VerificationCode *string `json:"VerificationCode,omitnil" name:"VerificationCode"`

	// 提示信息
	PromptInformation *string `json:"PromptInformation,omitnil" name:"PromptInformation"`

	// 保险费
	Insurance *string `json:"Insurance,omitnil" name:"Insurance"`

	// 填开单位
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 填开时间
	Date *string `json:"Date,omitnil" name:"Date"`

	// 国内国际标识
	DomesticInternationalTag *string `json:"DomesticInternationalTag,omitnil" name:"DomesticInternationalTag"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 销售方名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 统一社会信用代码
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 机票详细信息元组
	FlightItems []*FlightItemInfo `json:"FlightItems,omitnil" name:"FlightItems"`
}

type ElectronicTrainTicket struct {
	// 购方名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuyerName *string `json:"BuyerName,omitnil" name:"BuyerName"`

	// 购方识别号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuyerTaxCode *string `json:"BuyerTaxCode,omitnil" name:"BuyerTaxCode"`

	// 发票号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitnil" name:"Date"`

	// 价税合计（中文大写）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCN *string `json:"TotalCN,omitnil" name:"TotalCN"`

	// 税额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 业务类型，0：退票，1:售票
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 出发时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// 车次
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainNumber *string `json:"TrainNumber,omitnil" name:"TrainNumber"`

	// 发票代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 席别
	// 注意：此字段可能返回 null，表示取不到有效值。
	SeatType *string `json:"SeatType,omitnil" name:"SeatType"`

	// 乘车日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// 车厢
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainCabin *string `json:"TrainCabin,omitnil" name:"TrainCabin"`

	// 出发站
	// 注意：此字段可能返回 null，表示取不到有效值。
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// 电子客票号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElectronicNumber *string `json:"ElectronicNumber,omitnil" name:"ElectronicNumber"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassengerName *string `json:"PassengerName,omitnil" name:"PassengerName"`

	// 证件号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassengerNo *string `json:"PassengerNo,omitnil" name:"PassengerNo"`

	// 金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *string `json:"Amount,omitnil" name:"Amount"`

	// 到达站
	// 注意：此字段可能返回 null，表示取不到有效值。
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// 税率
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 席位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// 价税合计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *string `json:"Total,omitnil" name:"Total"`

	// 校验码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 发票状态代码，0正常 1 未更新  2作废 3已红冲
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateCode *string `json:"StateCode,omitnil" name:"StateCode"`
}

type ElectronicTrainTicketFull struct {
	// 电子发票类型
	TypeOfVoucher *string `json:"TypeOfVoucher,omitnil" name:"TypeOfVoucher"`

	// 电子客票号
	ElectronicTicketNum *string `json:"ElectronicTicketNum,omitnil" name:"ElectronicTicketNum"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 始发站
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// 到达站
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// 火车号
	TrainNumber *string `json:"TrainNumber,omitnil" name:"TrainNumber"`

	// 乘车日期
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// 始发时间
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// 座位类型
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// 座位号
	SeatNumber *string `json:"SeatNumber,omitnil" name:"SeatNumber"`

	// 票价
	Fare *string `json:"Fare,omitnil" name:"Fare"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 身份证号
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// 乘车人姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 金额
	Total *string `json:"Total,omitnil" name:"Total"`

	// 税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 统一社会信用代码
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 原发票号码
	OriginalNumber *string `json:"OriginalNumber,omitnil" name:"OriginalNumber"`
}

type Encryption struct {
	// 有加密需求的用户，接入传入kms的CiphertextBlob，关于数据加密可查阅数据加密 文档。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CiphertextBlob *string `json:"CiphertextBlob,omitnil" name:"CiphertextBlob"`

	// 有加密需求的用户，传入CBC加密的初始向量（客户自定义字符串，长度16字符）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Iv *string `json:"Iv,omitnil" name:"Iv"`

	// 加密使用的算法（支持'AES-256-CBC'、'SM4-GCM'），不传默认为'AES-256-CBC'
	// 注意：此字段可能返回 null，表示取不到有效值。
	Algorithm *string `json:"Algorithm,omitnil" name:"Algorithm"`

	// SM4-GCM算法生成的消息摘要（校验消息完整性时使用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*string `json:"TagList,omitnil" name:"TagList"`

	// 在使用加密服务时，指定要被加密的字段。本接口默认为EncryptedBody
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptList []*string `json:"EncryptList,omitnil" name:"EncryptList"`
}

// Predefined struct for user
type EnglishOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。像素须介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。像素须介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 单词四点坐标开关，开启可返回图片中单词的四点坐标。
	// 该参数默认值为false。
	EnableCoordPoint *bool `json:"EnableCoordPoint,omitnil" name:"EnableCoordPoint"`

	// 候选字开关，开启可返回识别时多个可能的候选字（每个候选字对应其置信度）。
	// 该参数默认值为false。
	EnableCandWord *bool `json:"EnableCandWord,omitnil" name:"EnableCandWord"`

	// 预处理开关，功能是检测图片倾斜的角度，将原本倾斜的图片矫正。该参数默认值为true。
	Preprocess *bool `json:"Preprocess,omitnil" name:"Preprocess"`
}

type EnglishOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。像素须介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。像素须介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 单词四点坐标开关，开启可返回图片中单词的四点坐标。
	// 该参数默认值为false。
	EnableCoordPoint *bool `json:"EnableCoordPoint,omitnil" name:"EnableCoordPoint"`

	// 候选字开关，开启可返回识别时多个可能的候选字（每个候选字对应其置信度）。
	// 该参数默认值为false。
	EnableCandWord *bool `json:"EnableCandWord,omitnil" name:"EnableCandWord"`

	// 预处理开关，功能是检测图片倾斜的角度，将原本倾斜的图片矫正。该参数默认值为true。
	Preprocess *bool `json:"Preprocess,omitnil" name:"Preprocess"`
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

// Predefined struct for user
type EnglishOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	TextDetections []*TextDetectionEn `json:"TextDetections,omitnil" name:"TextDetections"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnglishOCRResponse struct {
	*tchttp.BaseResponse
	Response *EnglishOCRResponseParams `json:"Response"`
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
	// 支持以下字段：统一社会信用代码、法定代表人、公司名称、公司地址、注册资金、企业类型、经营范围、成立日期、有效期、开办资金、经费来源、举办单位等；
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type EnterpriseLicenseOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type EnterpriseLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type EnterpriseLicenseOCRResponseParams struct {
	// 企业证照识别结果，具体内容请点击左侧链接。
	EnterpriseLicenseInfos []*EnterpriseLicenseInfo `json:"EnterpriseLicenseInfos,omitnil" name:"EnterpriseLicenseInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnterpriseLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *EnterpriseLicenseOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type EstateCertOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type EstateCertOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type EstateCertOCRResponseParams struct {
	// 权利人
	Obligee *string `json:"Obligee,omitnil" name:"Obligee"`

	// 共有情况
	Ownership *string `json:"Ownership,omitnil" name:"Ownership"`

	// 坐落
	Location *string `json:"Location,omitnil" name:"Location"`

	// 不动产单元号
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 权利类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 权利性质
	Property *string `json:"Property,omitnil" name:"Property"`

	// 用途
	Usage *string `json:"Usage,omitnil" name:"Usage"`

	// 面积
	Area *string `json:"Area,omitnil" name:"Area"`

	// 使用期限
	Term *string `json:"Term,omitnil" name:"Term"`

	// 权利其他状况，多行会用换行符\n连接。
	Other *string `json:"Other,omitnil" name:"Other"`

	// 图片旋转角度
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 不动产权号
	Number *string `json:"Number,omitnil" name:"Number"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EstateCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *EstateCertOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type FinanBillOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type FinanBillOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type FinanBillOCRResponseParams struct {
	// 金融票据整单识别结果，具体内容请点击左侧链接。
	FinanBillInfos []*FinanBillInfo `json:"FinanBillInfos,omitnil" name:"FinanBillInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FinanBillOCRResponse struct {
	*tchttp.BaseResponse
	Response *FinanBillOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type FinanBillSliceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type FinanBillSliceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type FinanBillSliceOCRResponseParams struct {
	// 金融票据切片识别结果，具体内容请点击左侧链接。
	FinanBillSliceInfos []*FinanBillSliceInfo `json:"FinanBillSliceInfos,omitnil" name:"FinanBillSliceInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FinanBillSliceOCRResponse struct {
	*tchttp.BaseResponse
	Response *FinanBillSliceOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段 Name 对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 多个行程的字段所在行号，下标从0开始，非行字段或未能识别行号的该值返回-1。
	Row *int64 `json:"Row,omitnil" name:"Row"`
}

// Predefined struct for user
type FlightInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type FlightInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlightInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlightInvoiceOCRResponseParams struct {
	// 机票行程单识别结果，具体内容请点击左侧链接。
	FlightInvoiceInfos []*FlightInvoiceInfo `json:"FlightInvoiceInfos,omitnil" name:"FlightInvoiceInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FlightInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *FlightInvoiceOCRResponseParams `json:"Response"`
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

type FlightItem struct {
	// 出发航站楼
	TerminalGetOn *string `json:"TerminalGetOn,omitnil" name:"TerminalGetOn"`

	// 到达航站楼
	TerminalGetOff *string `json:"TerminalGetOff,omitnil" name:"TerminalGetOff"`

	// 承运人
	Carrier *string `json:"Carrier,omitnil" name:"Carrier"`

	// 航班号
	FlightNumber *string `json:"FlightNumber,omitnil" name:"FlightNumber"`

	// 座位等级
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// 乘机日期
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// 乘机时间
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// 出发站
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// 到达站
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// 免费行李
	Allow *string `json:"Allow,omitnil" name:"Allow"`

	// 客票级别/客票类别
	FareBasis *string `json:"FareBasis,omitnil" name:"FareBasis"`
}

type FlightItemInfo struct {
	// 出发站
	TerminalGetOn *string `json:"TerminalGetOn,omitnil" name:"TerminalGetOn"`

	// 到达站
	TerminalGetOff *string `json:"TerminalGetOff,omitnil" name:"TerminalGetOff"`

	// 承运人
	Carrier *string `json:"Carrier,omitnil" name:"Carrier"`

	// 航班号
	FlightNumber *string `json:"FlightNumber,omitnil" name:"FlightNumber"`

	// 座位等级
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// 乘机日期
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// 乘机时间
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// 客票级别/客票类别
	FareBasis *string `json:"FareBasis,omitnil" name:"FareBasis"`

	// 免费行李额
	Allow *string `json:"Allow,omitnil" name:"Allow"`
}

// Predefined struct for user
type FormulaOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type FormulaOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type FormulaOCRResponseParams struct {
	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负
	Angle *int64 `json:"Angle,omitnil" name:"Angle"`

	// 检测到的文本信息，具体内容请点击左侧链接。
	FormulaInfos []*TextFormula `json:"FormulaInfos,omitnil" name:"FormulaInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FormulaOCRResponse struct {
	*tchttp.BaseResponse
	Response *FormulaOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type GeneralAccurateOCRRequestParams struct {
	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回单字信息，默认关
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`

	// 是否开启原图切图检测功能，开启后可提升“整图面积大，但单字符占比面积小”（例如：试卷）场景下的识别效果，默认关
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil" name:"EnableDetectSplit"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 文本检测开关，默认为true。设置为false可直接进行单行识别，适用于仅包含正向单行文本的图片场景。
	EnableDetectText *bool `json:"EnableDetectText,omitnil" name:"EnableDetectText"`
}

type GeneralAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回单字信息，默认关
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`

	// 是否开启原图切图检测功能，开启后可提升“整图面积大，但单字符占比面积小”（例如：试卷）场景下的识别效果，默认关
	EnableDetectSplit *bool `json:"EnableDetectSplit,omitnil" name:"EnableDetectSplit"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 文本检测开关，默认为true。设置为false可直接进行单行识别，适用于仅包含正向单行文本的图片场景。
	EnableDetectText *bool `json:"EnableDetectText,omitnil" name:"EnableDetectText"`
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
	delete(f, "EnableDetectText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneralAccurateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GeneralAccurateOCRResponseParams struct {
	// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
	TextDetections []*TextDetection `json:"TextDetections,omitnil" name:"TextDetections"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	//
	// Deprecated: Angel is deprecated.
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GeneralAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *GeneralAccurateOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type GeneralBasicOCRRequestParams struct {
	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 保留字段。
	Scene *string `json:"Scene,omitnil" name:"Scene"`

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
	LanguageType *string `json:"LanguageType,omitnil" name:"LanguageType"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否返回单字信息，默认关
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`
}

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 保留字段。
	Scene *string `json:"Scene,omitnil" name:"Scene"`

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
	LanguageType *string `json:"LanguageType,omitnil" name:"LanguageType"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否返回单字信息，默认关
	IsWords *bool `json:"IsWords,omitnil" name:"IsWords"`
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

// Predefined struct for user
type GeneralBasicOCRResponseParams struct {
	// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
	TextDetections []*TextDetection `json:"TextDetections,omitnil" name:"TextDetections"`

	// 检测到的语言类型，目前支持的语言类型参考入参LanguageType说明。
	Language *string `json:"Language,omitnil" name:"Language"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	//
	// Deprecated: Angel is deprecated.
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// 图片为PDF时，返回PDF的总页数，默认为0
	PdfPageSize *int64 `json:"PdfPageSize,omitnil" name:"PdfPageSize"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GeneralBasicOCRResponse struct {
	*tchttp.BaseResponse
	Response *GeneralBasicOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type GeneralEfficientOCRRequestParams struct {
	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type GeneralEfficientOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 要求图片经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type GeneralEfficientOCRResponseParams struct {
	// 检测到的文本信息，包括文本行内容、置信度、文本行坐标以及文本行旋转纠正后的坐标，具体内容请点击左侧链接。
	TextDetections []*TextDetection `json:"TextDetections,omitnil" name:"TextDetections"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GeneralEfficientOCRResponse struct {
	*tchttp.BaseResponse
	Response *GeneralEfficientOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type GeneralFastOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type GeneralFastOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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

// Predefined struct for user
type GeneralFastOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	TextDetections []*TextDetection `json:"TextDetections,omitnil" name:"TextDetections"`

	// 检测到的语言，目前支持的语种范围为：简体中文、繁体中文、英文、日文、韩文。未来将陆续新增对更多语种的支持。
	// 返回结果含义为：zh - 中英混合，jap - 日文，kor - 韩文。
	Language *string `json:"Language,omitnil" name:"Language"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// 图片为PDF时，返回PDF的总页数，默认为0
	PdfPageSize *int64 `json:"PdfPageSize,omitnil" name:"PdfPageSize"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GeneralFastOCRResponse struct {
	*tchttp.BaseResponse
	Response *GeneralFastOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type GeneralHandwritingOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 场景字段，默认不用填写。
	// 可选值:only_hw  表示只输出手写体识别结果，过滤印刷体。
	Scene *string `json:"Scene,omitnil" name:"Scene"`

	// 是否开启单字的四点定位坐标输出，默认值为false。
	EnableWordPolygon *bool `json:"EnableWordPolygon,omitnil" name:"EnableWordPolygon"`

	// 文本检测开关，默认值为true。
	// 设置为false表示直接进行单行识别，可适用于识别单行手写体签名场景。
	EnableDetectText *bool `json:"EnableDetectText,omitnil" name:"EnableDetectText"`
}

type GeneralHandwritingOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 场景字段，默认不用填写。
	// 可选值:only_hw  表示只输出手写体识别结果，过滤印刷体。
	Scene *string `json:"Scene,omitnil" name:"Scene"`

	// 是否开启单字的四点定位坐标输出，默认值为false。
	EnableWordPolygon *bool `json:"EnableWordPolygon,omitnil" name:"EnableWordPolygon"`

	// 文本检测开关，默认值为true。
	// 设置为false表示直接进行单行识别，可适用于识别单行手写体签名场景。
	EnableDetectText *bool `json:"EnableDetectText,omitnil" name:"EnableDetectText"`
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

// Predefined struct for user
type GeneralHandwritingOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	TextDetections []*TextGeneralHandwriting `json:"TextDetections,omitnil" name:"TextDetections"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	Angel *float64 `json:"Angel,omitnil" name:"Angel"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GeneralHandwritingOCRResponse struct {
	*tchttp.BaseResponse
	Response *GeneralHandwritingOCRResponseParams `json:"Response"`
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

type GeneralMachineItem struct {
	// 项目名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 规格型号
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 单位
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 数量
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// 单价
	Price *string `json:"Price,omitnil" name:"Price"`

	// 金额
	Total *string `json:"Total,omitnil" name:"Total"`

	// 税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`
}

// Predefined struct for user
type GetTaskStateRequestParams struct {
	// 智慧表单任务唯一身份ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type GetTaskStateRequest struct {
	*tchttp.BaseRequest
	
	// 智慧表单任务唯一身份ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *GetTaskStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStateResponseParams struct {
	// 1:任务识别完成，还未提交
	// 2:任务已手动关闭
	// 3:任务已提交
	// 4:任务识别中
	// 5:超时：任务超过了可操作的24H时限
	// 6:任务识别失败
	TaskState *uint64 `json:"TaskState,omitnil" name:"TaskState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTaskStateResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskStateResponseParams `json:"Response"`
}

func (r *GetTaskStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 每一行的元素
	Groups []*LineInfo `json:"Groups,omitnil" name:"Groups"`
}

// Predefined struct for user
type HKIDCardOCRRequestParams struct {
	// 是否鉴伪。
	DetectFake *bool `json:"DetectFake,omitnil" name:"DetectFake"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type HKIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 是否鉴伪。
	DetectFake *bool `json:"DetectFake,omitnil" name:"DetectFake"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type HKIDCardOCRResponseParams struct {
	// 中文姓名
	CnName *string `json:"CnName,omitnil" name:"CnName"`

	// 英文姓名
	EnName *string `json:"EnName,omitnil" name:"EnName"`

	// 中文姓名对应电码
	TelexCode *string `json:"TelexCode,omitnil" name:"TelexCode"`

	// 性别 ：“男M”或“女F”
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 出生日期
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// 永久性居民身份证。
	// 0：非永久；
	// 1：永久；
	// -1：未知。
	Permanent *int64 `json:"Permanent,omitnil" name:"Permanent"`

	// 身份证号码
	IdNum *string `json:"IdNum,omitnil" name:"IdNum"`

	// 证件符号，出生日期下的符号，例如"***AZ"
	Symbol *string `json:"Symbol,omitnil" name:"Symbol"`

	// 首次签发日期
	FirstIssueDate *string `json:"FirstIssueDate,omitnil" name:"FirstIssueDate"`

	// 最近领用日期
	CurrentIssueDate *string `json:"CurrentIssueDate,omitnil" name:"CurrentIssueDate"`

	// 真假判断。
	// 0：无法判断（图像模糊、不完整、反光、过暗等导致无法判断）；
	// 1：假；
	// 2：真。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FakeDetectResult *int64 `json:"FakeDetectResult,omitnil" name:"FakeDetectResult"`

	// 人像照片Base64后的结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadImage *string `json:"HeadImage,omitnil" name:"HeadImage"`

	// 多重告警码，当身份证是翻拍、复印件时返回对应告警码。
	// -9102：证照复印件告警
	// -9103：证照翻拍告警
	WarningCode []*int64 `json:"WarningCode,omitnil" name:"WarningCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type HKIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *HKIDCardOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type HmtResidentPermitOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT：有照片的一面（人像面），
	// BACK：无照片的一面（国徽面），
	// 该参数如果不填或填错，将为您自动判断正反面。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
}

type HmtResidentPermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT：有照片的一面（人像面），
	// BACK：无照片的一面（国徽面），
	// 该参数如果不填或填错，将为您自动判断正反面。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
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

// Predefined struct for user
type HmtResidentPermitOCRResponseParams struct {
	// 证件姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 性别
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 出生日期
	Birth *string `json:"Birth,omitnil" name:"Birth"`

	// 地址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 身份证号
	IdCardNo *string `json:"IdCardNo,omitnil" name:"IdCardNo"`

	// 0-正面
	// 1-反面
	CardType *int64 `json:"CardType,omitnil" name:"CardType"`

	// 证件有效期限
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// 签发机关
	Authority *string `json:"Authority,omitnil" name:"Authority"`

	// 签发次数
	VisaNum *string `json:"VisaNum,omitnil" name:"VisaNum"`

	// 通行证号码
	PassNo *string `json:"PassNo,omitnil" name:"PassNo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type HmtResidentPermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *HmtResidentPermitOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type IDCardOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT：身份证有照片的一面（人像面），
	// BACK：身份证有国徽的一面（国徽面），
	// 该参数如果不填，将为您自动判断身份证正反面。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`

	// 以下可选字段均为bool 类型，默认false：
	// CropIdCard，身份证照片裁剪（去掉证件外多余的边缘、自动矫正拍摄角度）
	// CropPortrait，人像照片裁剪（自动抠取身份证头像区域）
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，疑似存在PS痕迹告警
	// TempIdWarn，临时身份证告警
	// InvalidDateWarn，身份证有效日期不合法告警
	// Quality，图片质量分数（评价图片的模糊程度）
	// MultiCardDetect，是否开启正反面同框识别（仅支持二代身份证正反页同框识别或临时身份证正反页同框识别）
	// ReflectWarn，是否开启反光检测
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitnil" name:"Config"`

	// 默认值为true，打开识别结果纠正开关。开关开启后，身份证号、出生日期、性别，三个字段会进行矫正补齐，统一结果输出；若关闭此开关，以上三个字段不会进行矫正补齐，保持原始识别结果输出，若原图出现篡改情况，这三个字段的识别结果可能会不统一。
	EnableRecognitionRectify *bool `json:"EnableRecognitionRectify,omitnil" name:"EnableRecognitionRectify"`

	// 默认值为false。
	// 
	// 此开关需要在反光检测开关开启下才会生效（即此开关生效的前提是config入参里的"ReflectWarn":true），若EnableReflectDetail设置为true，则会返回反光点覆盖区域详情。反光点覆盖区域详情分为四部分：人像照片位置、国徽位置、识别字段位置、其他位置。一个反光点允许覆盖多个区域，且一张图片可能存在多个反光点。
	EnableReflectDetail *bool `json:"EnableReflectDetail,omitnil" name:"EnableReflectDetail"`
}

type IDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT：身份证有照片的一面（人像面），
	// BACK：身份证有国徽的一面（国徽面），
	// 该参数如果不填，将为您自动判断身份证正反面。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`

	// 以下可选字段均为bool 类型，默认false：
	// CropIdCard，身份证照片裁剪（去掉证件外多余的边缘、自动矫正拍摄角度）
	// CropPortrait，人像照片裁剪（自动抠取身份证头像区域）
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，疑似存在PS痕迹告警
	// TempIdWarn，临时身份证告警
	// InvalidDateWarn，身份证有效日期不合法告警
	// Quality，图片质量分数（评价图片的模糊程度）
	// MultiCardDetect，是否开启正反面同框识别（仅支持二代身份证正反页同框识别或临时身份证正反页同框识别）
	// ReflectWarn，是否开启反光检测
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitnil" name:"Config"`

	// 默认值为true，打开识别结果纠正开关。开关开启后，身份证号、出生日期、性别，三个字段会进行矫正补齐，统一结果输出；若关闭此开关，以上三个字段不会进行矫正补齐，保持原始识别结果输出，若原图出现篡改情况，这三个字段的识别结果可能会不统一。
	EnableRecognitionRectify *bool `json:"EnableRecognitionRectify,omitnil" name:"EnableRecognitionRectify"`

	// 默认值为false。
	// 
	// 此开关需要在反光检测开关开启下才会生效（即此开关生效的前提是config入参里的"ReflectWarn":true），若EnableReflectDetail设置为true，则会返回反光点覆盖区域详情。反光点覆盖区域详情分为四部分：人像照片位置、国徽位置、识别字段位置、其他位置。一个反光点允许覆盖多个区域，且一张图片可能存在多个反光点。
	EnableReflectDetail *bool `json:"EnableReflectDetail,omitnil" name:"EnableReflectDetail"`
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
	delete(f, "EnableRecognitionRectify")
	delete(f, "EnableReflectDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IDCardOCRResponseParams struct {
	// 姓名（人像面）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 性别（人像面）
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 民族（人像面）
	Nation *string `json:"Nation,omitnil" name:"Nation"`

	// 出生日期（人像面）
	Birth *string `json:"Birth,omitnil" name:"Birth"`

	// 地址（人像面）
	Address *string `json:"Address,omitnil" name:"Address"`

	// 身份证号（人像面）
	IdNum *string `json:"IdNum,omitnil" name:"IdNum"`

	// 发证机关（国徽面）
	Authority *string `json:"Authority,omitnil" name:"Authority"`

	// 证件有效期（国徽面）
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

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
	// -9106	身份证疑似存在PS痕迹告警，
	// -9107       身份证反光告警。
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 反光点覆盖区域详情结果，具体内容请点击左侧链接
	ReflectDetailInfos []*ReflectDetailInfo `json:"ReflectDetailInfos,omitnil" name:"ReflectDetailInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *IDCardOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type ImageEnhancementRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 默认为空，ReturnImage的取值以及含义如下：
	// “preprocess”: 返回预处理后的图片数据
	// “origin”：返回原图片数据
	// " ":不返回图片数据
	ReturnImage *string `json:"ReturnImage,omitnil" name:"ReturnImage"`

	// 默认值为1，指定图像增强方法：
	// 1：切边增强
	// 2：弯曲矫正
	// 202：黑白模式
	// 204：提亮模式
	// 205：灰度模式
	// 207：省墨模式
	// 208：文字锐化（适合非彩色图片）
	// 300:自动增强（自动从301～304选择任务类型）
	// 301：去摩尔纹
	// 302：去除阴影
	// 303：去除模糊 
	// 304：去除过曝
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`
}

type ImageEnhancementRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 默认为空，ReturnImage的取值以及含义如下：
	// “preprocess”: 返回预处理后的图片数据
	// “origin”：返回原图片数据
	// " ":不返回图片数据
	ReturnImage *string `json:"ReturnImage,omitnil" name:"ReturnImage"`

	// 默认值为1，指定图像增强方法：
	// 1：切边增强
	// 2：弯曲矫正
	// 202：黑白模式
	// 204：提亮模式
	// 205：灰度模式
	// 207：省墨模式
	// 208：文字锐化（适合非彩色图片）
	// 300:自动增强（自动从301～304选择任务类型）
	// 301：去摩尔纹
	// 302：去除阴影
	// 303：去除模糊 
	// 304：去除过曝
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`
}

func (r *ImageEnhancementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageEnhancementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnImage")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageEnhancementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageEnhancementResponseParams struct {
	// 图片数据标识：
	// “origin”：原图
	// “preprocess”:预处理后的图
	ImageTag *string `json:"ImageTag,omitnil" name:"ImageTag"`

	// 图片数据，返回预处理后图像或原图像base64字符
	Image *string `json:"Image,omitnil" name:"Image"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImageEnhancementResponse struct {
	*tchttp.BaseResponse
	Response *ImageEnhancementResponseParams `json:"Response"`
}

func (r *ImageEnhancementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageEnhancementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstitutionOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type InstitutionOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type InstitutionOCRResponseParams struct {
	// 注册号
	RegId *string `json:"RegId,omitnil" name:"RegId"`

	// 有效期
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// 住所
	Location *string `json:"Location,omitnil" name:"Location"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 法定代表人
	LegalPerson *string `json:"LegalPerson,omitnil" name:"LegalPerson"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InstitutionOCRResponse struct {
	*tchttp.BaseResponse
	Response *InstitutionOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type InsuranceBillOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type InsuranceBillOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type InsuranceBillOCRResponseParams struct {
	// 保险单据识别结果，具体内容请点击左侧链接。
	InsuranceBillInfos []*InsuranceBillInfo `json:"InsuranceBillInfos,omitnil" name:"InsuranceBillInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InsuranceBillOCRResponse struct {
	*tchttp.BaseResponse
	Response *InsuranceBillOCRResponseParams `json:"Response"`
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
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

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
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 识别出的图片在混贴票据图片中的位置信息。与Angel结合可以得出原图位置，组成RotatedRect((X+0.5\*Width,Y+0.5\*Height), (Width, Height), Angle)，详情可参考OpenCV文档。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`

	// 入参 ReturnImage 为 True 时返回 Base64 编码后的图片。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil" name:"Image"`
}

type InvoiceGeneralInfo struct {
	// 识别出的字段名称(关键字)，支持以下字段识别（注：下划线表示一个字段）：
	// 发票代码、发票号码、日期、合计金额(小写)、合计金额(大写)、购买方识别号、销售方识别号、校验码、购买方名称、销售方名称、时间、种类、发票消费类型、省、市、是否有公司印章、发票名称、<span style="text-decoration:underline">购买方地址、电话</span>、<span style="text-decoration:underline">销售方地址、电话</span>、购买方开户行及账号、销售方开户行及账号、经办人取票用户、经办人支付信息、经办人商户号、经办人订单号、<span style="text-decoration:underline">货物或应税劳务、服务名称</span>、数量、单价、税率、税额、金额、单位、规格型号、合计税额、合计金额、备注、收款人、复核、开票人、密码区、行业分类
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type InvoiceGeneralOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type InvoiceGeneralOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvoiceGeneralOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvoiceGeneralOCRResponseParams struct {
	// 通用机打发票识别结果，具体内容请点击左侧链接。
	InvoiceGeneralInfos []*InvoiceGeneralInfo `json:"InvoiceGeneralInfos,omitnil" name:"InvoiceGeneralInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InvoiceGeneralOCRResponse struct {
	*tchttp.BaseResponse
	Response *InvoiceGeneralOCRResponseParams `json:"Response"`
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

type InvoiceItem struct {
	// 识别结果。
	// OK：表示识别成功；FailedOperation.UnsupportedInvoice：表示不支持识别；
	// FailedOperation.UnKnowError：表示识别失败；
	// 其它错误码见各个票据接口的定义。
	Code *string `json:"Code,omitnil" name:"Code"`

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
	// 11：增值税发票（卷票）
	// 12：购车发票
	// 13：过路过桥费发票
	// 15：非税发票
	// 16：全电发票
	// 17：医疗发票
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 该发票在原图片中的四点坐标。
	Polygon *Polygon `json:"Polygon,omitnil" name:"Polygon"`

	// 识别出的图片在混贴票据图片中的旋转角度。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 识别到的内容。
	SingleInvoiceInfos *SingleInvoiceItem `json:"SingleInvoiceInfos,omitnil" name:"SingleInvoiceInfos"`

	// 发票处于识别图片或PDF文件中的页教，默认从1开始。
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 发票详细类型，详见票据识别（高级版）接口文档说明中 SubType 返回值说明
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 发票类型描述，详见票据识别（高级版）接口文档说明中 TypeDescription  返回值说明
	TypeDescription *string `json:"TypeDescription,omitnil" name:"TypeDescription"`

	// 切割单图文件，Base64编码后的切图后的图片文件，开启 EnableCutImage 后进行返回
	CutImage *string `json:"CutImage,omitnil" name:"CutImage"`

	// 发票详细类型描述，详见上方 SubType 返回值说明
	SubTypeDescription *string `json:"SubTypeDescription,omitnil" name:"SubTypeDescription"`

	// 该发票中所有字段坐标信息。包括字段英文名称、字段值所在位置四点坐标、字段所属行号，具体内容请点击左侧链接。
	ItemPolygon []*ItemPolygonInfo `json:"ItemPolygon,omitnil" name:"ItemPolygon"`
}

type ItemCoord struct {
	// 左上角x
	X *int64 `json:"X,omitnil" name:"X"`

	// 左上角y
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 宽width
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 高height
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type ItemInfo struct {
	// key信息组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *Key `json:"Key,omitnil" name:"Key"`

	// Value信息组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *Value `json:"Value,omitnil" name:"Value"`
}

type ItemPolygonInfo struct {
	// 发票的英文字段名称（如Title）
	Key *string `json:"Key,omitnil" name:"Key"`

	// 字段值所在位置的四点坐标
	Polygon *Polygon `json:"Polygon,omitnil" name:"Polygon"`

	// 字段属于第几行，用于相同字段的排版，如发票明细表格项目，普通字段使用默认值为-1，表示无列排版。
	Row *int64 `json:"Row,omitnil" name:"Row"`
}

type Key struct {
	// 自动识别的字段名称
	AutoName *string `json:"AutoName,omitnil" name:"AutoName"`

	// 定义的字段名称（传key的名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitnil" name:"ConfigName"`
}

type LicensePlateInfo struct {
	// 识别出的车牌号码。
	Number *string `json:"Number,omitnil" name:"Number"`

	// 置信度，0 - 100 之间。
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行在原图片中的像素坐标框。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`

	// 识别出的车牌颜色，目前支持颜色包括 “白”、“黑”、“蓝”、“绿“、“黄”、“黄绿”、“临牌”、“喷漆”、“其它”。
	Color *string `json:"Color,omitnil" name:"Color"`
}

// Predefined struct for user
type LicensePlateOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type LicensePlateOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type LicensePlateOCRResponseParams struct {
	// 识别出的车牌号码。
	Number *string `json:"Number,omitnil" name:"Number"`

	// 置信度，0 - 100 之间。
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行在原图片中的像素坐标框。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`

	// 识别出的车牌颜色，目前支持颜色包括 “白”、“黑”、“蓝”、“绿“、“黄”、“黄绿”、“临牌”、“喷漆”、“其它”。
	Color *string `json:"Color,omitnil" name:"Color"`

	// 全部车牌信息。
	LicensePlateInfos []*LicensePlateInfo `json:"LicensePlateInfos,omitnil" name:"LicensePlateInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type LicensePlateOCRResponse struct {
	*tchttp.BaseResponse
	Response *LicensePlateOCRResponseParams `json:"Response"`
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

type LineInfo struct {
	// 每行的一个元素
	Lines []*ItemInfo `json:"Lines,omitnil" name:"Lines"`
}

// Predefined struct for user
type MLIDCardOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。( 中国地区之外不支持这个字段 )
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回图片，默认false
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`
}

type MLIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。( 中国地区之外不支持这个字段 )
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回图片，默认false
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`
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

// Predefined struct for user
type MLIDCardOCRResponseParams struct {
	// 身份证号
	ID *string `json:"ID,omitnil" name:"ID"`

	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 地址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 性别
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 告警码
	// -9103	证照翻拍告警
	// -9102	证照复印件告警
	// -9106       证件遮挡告警
	// -9107       模糊图片告警
	Warn []*int64 `json:"Warn,omitnil" name:"Warn"`

	// 证件图片
	Image *string `json:"Image,omitnil" name:"Image"`

	// 此字段为扩展字段。
	// 返回字段识别结果的置信度，格式如下
	// {
	//   字段名:{
	//     Confidence:0.9999
	//   }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 证件类型
	// MyKad  身份证
	// MyPR    永居证
	// MyTentera   军官证
	// MyKAS    临时身份证
	// POLIS  警察证
	// IKAD   劳工证
	// MyKid 儿童卡
	Type *string `json:"Type,omitnil" name:"Type"`

	// 出生日期（目前该字段仅支持IKAD劳工证、MyKad 身份证）
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MLIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *MLIDCardOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type MLIDPassportOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 是否返回图片，默认false
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 是否返回图片，默认false
	RetImage *bool `json:"RetImage,omitnil" name:"RetImage"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MLIDPassportOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MLIDPassportOCRResponseParams struct {
	// 护照ID（机读码区的解析结果）
	ID *string `json:"ID,omitnil" name:"ID"`

	// 姓名（机读码区的解析结果）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 出生日期（机读码区的解析结果）
	DateOfBirth *string `json:"DateOfBirth,omitnil" name:"DateOfBirth"`

	// 性别（F女，M男）（机读码区的解析结果）
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 有效期（机读码区的解析结果）
	DateOfExpiration *string `json:"DateOfExpiration,omitnil" name:"DateOfExpiration"`

	// 发行国（机读码区的解析结果）
	IssuingCountry *string `json:"IssuingCountry,omitnil" name:"IssuingCountry"`

	// 国家地区代码（机读码区的解析结果）
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// 告警码：
	// -9103	证照翻拍告警
	// -9102	证照复印件告警（包括黑白复印件、彩色复印件）
	// -9106       证件遮挡告警
	Warn []*int64 `json:"Warn,omitnil" name:"Warn"`

	// 证件图片
	Image *string `json:"Image,omitnil" name:"Image"`

	// 扩展字段:
	// {
	//     ID:{
	//         Confidence:0.9999
	//     },
	//     Name:{
	//         Confidence:0.9996
	//     }
	// }
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 最下方第一行 MRZ Code 序列
	CodeSet *string `json:"CodeSet,omitnil" name:"CodeSet"`

	// 最下方第二行 MRZ Code 序列
	CodeCrc *string `json:"CodeCrc,omitnil" name:"CodeCrc"`

	// 姓（机读码区的解析结果）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Surname *string `json:"Surname,omitnil" name:"Surname"`

	// 名（机读码区的解析结果）
	// 注意：此字段可能返回 null，表示取不到有效值。
	GivenName *string `json:"GivenName,omitnil" name:"GivenName"`

	// 类型（机读码区的解析结果）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 信息区证件内容
	PassportRecognizeInfos *PassportRecognizeInfos `json:"PassportRecognizeInfos,omitnil" name:"PassportRecognizeInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MLIDPassportOCRResponse struct {
	*tchttp.BaseResponse
	Response *MLIDPassportOCRResponseParams `json:"Response"`
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

type MachinePrintedInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 校验码
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 密码区
	Ciphertext *string `json:"Ciphertext,omitnil" name:"Ciphertext"`

	// 种类
	Category *string `json:"Category,omitnil" name:"Category"`

	// 税前金额
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 合计税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 行业分类
	IndustryClass *string `json:"IndustryClass,omitnil" name:"IndustryClass"`

	// 销售方名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 销售方纳税人识别号
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// 销售方地址电话
	SellerAddrTel *string `json:"SellerAddrTel,omitnil" name:"SellerAddrTel"`

	// 销售方银行账号
	SellerBankAccount *string `json:"SellerBankAccount,omitnil" name:"SellerBankAccount"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 购买方纳税人识别号
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 购买方地址电话
	BuyerAddrTel *string `json:"BuyerAddrTel,omitnil" name:"BuyerAddrTel"`

	// 购买方银行账号
	BuyerBankAccount *string `json:"BuyerBankAccount,omitnil" name:"BuyerBankAccount"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 是否有公司印章（0：没有，1：有）
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// 是否为浙江/广东通用机打发票（0：没有，1：有）
	ElectronicMark *int64 `json:"ElectronicMark,omitnil" name:"ElectronicMark"`

	// 开票人
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 收款人
	Receiptor *string `json:"Receiptor,omitnil" name:"Receiptor"`

	// 复核人
	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 经办人支付信息
	PaymentInfo *string `json:"PaymentInfo,omitnil" name:"PaymentInfo"`

	// 经办人取票用户
	TicketPickupUser *string `json:"TicketPickupUser,omitnil" name:"TicketPickupUser"`

	// 经办人商户号
	MerchantNumber *string `json:"MerchantNumber,omitnil" name:"MerchantNumber"`

	// 经办人订单号
	OrderNumber *string `json:"OrderNumber,omitnil" name:"OrderNumber"`

	// 条目
	GeneralMachineItems []*GeneralMachineItem `json:"GeneralMachineItems,omitnil" name:"GeneralMachineItems"`
}

// Predefined struct for user
type MainlandPermitOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回头像。默认不返回。
	RetProfile *bool `json:"RetProfile,omitnil" name:"RetProfile"`

	// 图片正反面
	// FRONT：正面、BACK：反面，默认为FRONT
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
}

type MainlandPermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回头像。默认不返回。
	RetProfile *bool `json:"RetProfile,omitnil" name:"RetProfile"`

	// 图片正反面
	// FRONT：正面、BACK：反面，默认为FRONT
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
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
	delete(f, "CardSide")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MainlandPermitOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MainlandPermitOCRResponseParams struct {
	// 中文姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 英文姓名
	EnglishName *string `json:"EnglishName,omitnil" name:"EnglishName"`

	// 性别
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 出生日期
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// 签发机关
	IssueAuthority *string `json:"IssueAuthority,omitnil" name:"IssueAuthority"`

	// 有效期限
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// 证件号
	Number *string `json:"Number,omitnil" name:"Number"`

	// 签发地点
	IssueAddress *string `json:"IssueAddress,omitnil" name:"IssueAddress"`

	// 签发次数
	IssueNumber *string `json:"IssueNumber,omitnil" name:"IssueNumber"`

	// 证件类别， 如：台湾居民来往大陆通行证、港澳居民来往内地通行证。
	Type *string `json:"Type,omitnil" name:"Type"`

	// RetProfile为True时返回头像字段， Base64编码
	Profile *string `json:"Profile,omitnil" name:"Profile"`

	// 背面字段信息
	MainlandTravelPermitBackInfos *MainlandTravelPermitBackInfos `json:"MainlandTravelPermitBackInfos,omitnil" name:"MainlandTravelPermitBackInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MainlandPermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *MainlandPermitOCRResponseParams `json:"Response"`
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

type MainlandTravelPermitBackInfos struct {
	// String	证件类别， 如：台湾居民来往大陆通行证、港澳居民来往内地通行证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 卡证背面的中文姓名	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 卡证背面的身份证号码	
	// 注意：此字段可能返回 null，表示取不到有效值。
	IDNumber *string `json:"IDNumber,omitnil" name:"IDNumber"`

	// 历史通行证号码	
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistoryNumber *string `json:"HistoryNumber,omitnil" name:"HistoryNumber"`
}

type MedicalInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 校验码
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 发票属地
	Place *string `json:"Place,omitnil" name:"Place"`

	// 复核人
	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`
}

type MedicalInvoiceInfo struct {
	// 医疗发票识别结果条目
	MedicalInvoiceItems []*MedicalInvoiceItem `json:"MedicalInvoiceItems,omitnil" name:"MedicalInvoiceItems"`
}

type MedicalInvoiceItem struct {
	// 识别出的字段名称
	// <table><tr><td>分类</td><td>name</td></tr><tr><td>票据基本信息</td><td>发票名称</td></tr><tr><td></td><td>票据代码</td></tr><tr><td></td><td>票据号码</td></tr><tr><td></td><td>电子票据代码</td></tr><tr><td></td><td>电子票据号码</td></tr><tr><td></td><td>交款人统一社会信用代码</td></tr><tr><td></td><td>校验码</td></tr><tr><td></td><td>交款人</td></tr><tr><td></td><td>开票日期</td></tr><tr><td></td><td>收款单位</td></tr><tr><td></td><td>复核人</td></tr><tr><td></td><td>收款人</td></tr><tr><td></td><td>业务流水号</td></tr><tr><td></td><td>门诊号</td></tr><tr><td></td><td>就诊日期</td></tr><tr><td></td><td>医疗机构类型</td></tr><tr><td></td><td>医保类型</td></tr><tr><td></td><td>医保编号</td></tr><tr><td></td><td>性别</td></tr><tr><td></td><td>医保统筹基金支付</td></tr><tr><td></td><td>其他支付</td></tr><tr><td></td><td>个人账户支付</td></tr><tr><td></td><td>个人现金支付</td></tr><tr><td></td><td>个人自付</td></tr><tr><td></td><td>个人自费</td></tr><tr><td></td><td>病历号</td></tr><tr><td></td><td>住院号</td></tr><tr><td></td><td>住院科别</td></tr><tr><td></td><td>住院时间</td></tr><tr><td></td><td>预缴金额</td></tr><tr><td></td><td>补缴金额</td></tr><tr><td></td><td>退费金额</td></tr><tr><td></td><td>发票属地</td></tr><tr><td></td><td>发票类型</td></tr><tr><td>总金额</td><td>总金额大写</td></tr><tr><td></td><td>总金额小写</td></tr><tr><td>收费大项</td><td>大项名称</td></tr><tr><td></td><td>大项金额</td></tr><tr><td>收费细项</td><td>项目名称</td></tr><tr><td></td><td>数量</td></tr><tr><td></td><td>单位</td></tr><tr><td></td><td>金额</td></tr><tr><td></td><td>备注</td></tr><tr><td>票据其他信息</td><td>入院时间</td></tr><tr><td></td><td>出院时间</td></tr><tr><td></td><td>住院天数</td></tr><tr><td></td><td>自付二</td></tr><tr><td></td><td>自付一</td></tr><tr><td></td><td>起付金额</td></tr><tr><td></td><td>超封顶金额</td></tr><tr><td></td><td>自费</td></tr><tr><td></td><td>本次医保范围内金额</td></tr><tr><td></td><td>累计医保内范围金额</td></tr><tr><td></td><td>门诊大额支付</td></tr><tr><td></td><td>残军补助支付</td></tr><tr><td></td><td>年度门诊大额累计支付</td></tr><tr><td></td><td>单位补充险[原公疗]支付</td></tr><tr><td></td><td>社会保障卡号</td></tr><tr><td></td><td>姓名</td></tr><tr><td></td><td>交易流水号</td></tr><tr><td></td><td>本次支付后个人账户余额</td></tr><tr><td></td><td>基金支付</td></tr><tr><td></td><td>现金支付</td></tr><tr><td></td><td>复核</td></tr><tr><td></td><td>自负</td></tr><tr><td></td><td>结算方式</td></tr><tr><td></td><td>医保统筹/公医记账</td></tr><tr><td></td><td>其他</td></tr><tr><td></td><td>个人支付金额</td></tr><tr><td></td><td>欠费</td></tr><tr><td></td><td>退休补充支付</td></tr><tr><td></td><td>医院类型</td></tr><tr><td></td><td>退款</td></tr><tr><td></td><td>补收</td></tr><tr><td></td><td>附加支付</td></tr><tr><td></td><td>分类自负</td></tr><tr><td></td><td>其它</td></tr><tr><td></td><td>预交款</td></tr><tr><td></td><td>个人缴费</td></tr></table>
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果
	Content *string `json:"Content,omitnil" name:"Content"`

	// 识别出的文本行四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vertex *Polygon `json:"Vertex,omitnil" name:"Vertex"`

	// 识别出的文本行在旋转纠正之后的图像中的像素坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coord *Rect `json:"Coord,omitnil" name:"Coord"`
}

// Predefined struct for user
type MixedInvoiceDetectRequestParams struct {
	// 是否需要返回裁剪后的图片。
	ReturnImage *bool `json:"ReturnImage,omitnil" name:"ReturnImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type MixedInvoiceDetectRequest struct {
	*tchttp.BaseRequest
	
	// 是否需要返回裁剪后的图片。
	ReturnImage *bool `json:"ReturnImage,omitnil" name:"ReturnImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MixedInvoiceDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MixedInvoiceDetectResponseParams struct {
	// 检测出的票据类型列表，具体内容请点击左侧链接。
	InvoiceDetectInfos []*InvoiceDetectInfo `json:"InvoiceDetectInfos,omitnil" name:"InvoiceDetectInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MixedInvoiceDetectResponse struct {
	*tchttp.BaseResponse
	Response *MixedInvoiceDetectResponseParams `json:"Response"`
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
	Code *string `json:"Code,omitnil" name:"Code"`

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
	// 11：增值税发票（卷票）
	// 12：购车发票
	// 13：过路过桥费发票
	// 15：非税发票
	// 16：全电发票
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 识别出的图片在混贴票据图片中的位置信息。与Angel结合可以得出原图位置，组成RotatedRect((X+0.5\*Width,Y+0.5\*Height), (Width, Height), Angle)，详情可参考OpenCV文档。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`

	// 识别出的图片在混贴票据图片中的旋转角度。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 识别到的内容。
	SingleInvoiceInfos []*SingleInvoiceInfo `json:"SingleInvoiceInfos,omitnil" name:"SingleInvoiceInfos"`

	// 发票处于识别图片或PDF文件中的页教，默认从1开始。
	Page *int64 `json:"Page,omitnil" name:"Page"`
}

// Predefined struct for user
type MixedInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

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
	// 15：非税发票
	// 16：全电发票
	// ----------------------
	// -1：其他发票,（只传入此类型时，图片均采用其他票类型进行识别）
	Types []*int64 `json:"Types,omitnil" name:"Types"`

	// 是否识别其他类型发票，默认为Yes
	// Yes：识别其他类型发票
	// No：不识别其他类型发票
	ReturnOther *string `json:"ReturnOther,omitnil" name:"ReturnOther"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否开启PDF多页识别，默认值为false，开启后可同时支持多页PDF的识别返回，仅支持返回文件前30页。开启后IsPDF和PdfPageNumber入参不进行控制。
	ReturnMultiplePage *bool `json:"ReturnMultiplePage,omitnil" name:"ReturnMultiplePage"`
}

type MixedInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

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
	// 15：非税发票
	// 16：全电发票
	// ----------------------
	// -1：其他发票,（只传入此类型时，图片均采用其他票类型进行识别）
	Types []*int64 `json:"Types,omitnil" name:"Types"`

	// 是否识别其他类型发票，默认为Yes
	// Yes：识别其他类型发票
	// No：不识别其他类型发票
	ReturnOther *string `json:"ReturnOther,omitnil" name:"ReturnOther"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否开启PDF多页识别，默认值为false，开启后可同时支持多页PDF的识别返回，仅支持返回文件前30页。开启后IsPDF和PdfPageNumber入参不进行控制。
	ReturnMultiplePage *bool `json:"ReturnMultiplePage,omitnil" name:"ReturnMultiplePage"`
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
	delete(f, "ReturnOther")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "ReturnMultiplePage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MixedInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MixedInvoiceOCRResponseParams struct {
	// 混贴票据识别结果，具体内容请点击左侧链接。
	MixedInvoiceItems []*MixedInvoiceItem `json:"MixedInvoiceItems,omitnil" name:"MixedInvoiceItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MixedInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *MixedInvoiceOCRResponseParams `json:"Response"`
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

type MotorVehicleSaleInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 税前金额
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 销售方名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 销售方单位代码
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// 销售方电话
	SellerTel *string `json:"SellerTel,omitnil" name:"SellerTel"`

	// 销售方地址
	SellerAddress *string `json:"SellerAddress,omitnil" name:"SellerAddress"`

	// 销售方开户行
	SellerBank *string `json:"SellerBank,omitnil" name:"SellerBank"`

	// 销售方银行账号
	SellerBankAccount *string `json:"SellerBankAccount,omitnil" name:"SellerBankAccount"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 购买方纳税人识别号
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 购买方身份证号码/组织机构代码
	BuyerID *string `json:"BuyerID,omitnil" name:"BuyerID"`

	// 主管税务机关
	TaxAuthorities *string `json:"TaxAuthorities,omitnil" name:"TaxAuthorities"`

	// 主管税务机关代码
	TaxAuthoritiesCode *string `json:"TaxAuthoritiesCode,omitnil" name:"TaxAuthoritiesCode"`

	// 车辆识别代码
	VIN *string `json:"VIN,omitnil" name:"VIN"`

	// 厂牌型号
	VehicleModel *string `json:"VehicleModel,omitnil" name:"VehicleModel"`

	// 发动机号码
	VehicleEngineCode *string `json:"VehicleEngineCode,omitnil" name:"VehicleEngineCode"`

	// 合格证号
	CertificateNumber *string `json:"CertificateNumber,omitnil" name:"CertificateNumber"`

	// 商检单号
	InspectionNumber *string `json:"InspectionNumber,omitnil" name:"InspectionNumber"`

	// 机器编号
	MachineID *string `json:"MachineID,omitnil" name:"MachineID"`

	// 车辆类型
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 合计税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 是否有公司印章（0：没有，1：有）
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// 吨位
	Tonnage *string `json:"Tonnage,omitnil" name:"Tonnage"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 发票联次
	FormType *string `json:"FormType,omitnil" name:"FormType"`

	// 发票联名
	FormName *string `json:"FormName,omitnil" name:"FormName"`

	// 开票人
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 进口证明书号
	TaxNum *string `json:"TaxNum,omitnil" name:"TaxNum"`

	// 完税凭证号码
	TaxPayNum *string `json:"TaxPayNum,omitnil" name:"TaxPayNum"`

	// 税控码
	TaxCode *string `json:"TaxCode,omitnil" name:"TaxCode"`

	// 限乘人数
	MaxPeopleNum *string `json:"MaxPeopleNum,omitnil" name:"MaxPeopleNum"`

	// 产地
	Origin *string `json:"Origin,omitnil" name:"Origin"`

	// 机打发票代码
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 机打发票号码
	MachineNumber *string `json:"MachineNumber,omitnil" name:"MachineNumber"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`
}

type NonTaxIncomeBill struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 校验码
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 交款人名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 交款人纳税人识别号
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 收款人名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 收款单位名称
	SellerCompany *string `json:"SellerCompany,omitnil" name:"SellerCompany"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 币种
	CurrencyCode *string `json:"CurrencyCode,omitnil" name:"CurrencyCode"`

	// 复核人
	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 其他信息
	OtherInfo *string `json:"OtherInfo,omitnil" name:"OtherInfo"`

	// 缴款码
	PaymentCode *string `json:"PaymentCode,omitnil" name:"PaymentCode"`

	// 执收单位编码
	ReceiveUnitCode *string `json:"ReceiveUnitCode,omitnil" name:"ReceiveUnitCode"`

	// 执收单位名称
	Receiver *string `json:"Receiver,omitnil" name:"Receiver"`

	// 经办人
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 付款人账号
	PayerAccount *string `json:"PayerAccount,omitnil" name:"PayerAccount"`

	// 付款人开户银行
	PayerBank *string `json:"PayerBank,omitnil" name:"PayerBank"`

	// 收款人账号
	ReceiverAccount *string `json:"ReceiverAccount,omitnil" name:"ReceiverAccount"`

	// 收款人开户银行
	ReceiverBank *string `json:"ReceiverBank,omitnil" name:"ReceiverBank"`

	// 条目
	NonTaxItems []*NonTaxItem `json:"NonTaxItems,omitnil" name:"NonTaxItems"`
}

type NonTaxItem struct {
	// 项目编码
	ItemID *string `json:"ItemID,omitnil" name:"ItemID"`

	// 项目名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 单位
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 数量
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// 标准
	Standard *string `json:"Standard,omitnil" name:"Standard"`

	// 金额
	Total *string `json:"Total,omitnil" name:"Total"`
}

type OnlineTaxiItineraryInfo struct {
	// 识别出的字段名称(关键字)，支持以下字段：
	// 发票代码、 机打代码、 发票号码、 发动机号码、 合格证号、 机打号码、 价税合计(小写)、 销货单位名称、 身份证号码/组织机构代码、 购买方名称、 销售方纳税人识别号、 购买方纳税人识别号、主管税务机关、 主管税务机关代码、 开票日期、 不含税价(小写)、 吨位、增值税税率或征收率、 车辆识别代号/车架号码、 增值税税额、 厂牌型号、 省、 市、 发票消费类型、 销售方电话、 销售方账号、 产地、 进口证明书号、 车辆类型、 机器编号、备注、开票人、限乘人数、商检单号、销售方地址、销售方开户银行、价税合计、发票类型。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 字段所在行，下标从0开始，非行字段或未能识别行号的返回-1
	Row *int64 `json:"Row,omitnil" name:"Row"`
}

// Predefined struct for user
type OrgCodeCertOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type OrgCodeCertOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type OrgCodeCertOCRResponseParams struct {
	// 代码
	OrgCode *string `json:"OrgCode,omitnil" name:"OrgCode"`

	// 机构名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 地址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 有效期
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type OrgCodeCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *OrgCodeCertOCRResponseParams `json:"Response"`
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

type OtherInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 金额
	Total *string `json:"Total,omitnil" name:"Total"`

	// 列表
	OtherInvoiceListItems []*OtherInvoiceItem `json:"OtherInvoiceListItems,omitnil" name:"OtherInvoiceListItems"`

	// 表格
	OtherInvoiceTableItems []*OtherInvoiceList `json:"OtherInvoiceTableItems,omitnil" name:"OtherInvoiceTableItems"`

	// 发票日期
	Date *string `json:"Date,omitnil" name:"Date"`
}

type OtherInvoiceItem struct {
	// 票面key值
	Name *string `json:"Name,omitnil" name:"Name"`

	// 票面value值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type OtherInvoiceList struct {
	// 列表
	OtherInvoiceItemList []*OtherInvoiceItem `json:"OtherInvoiceItemList,omitnil" name:"OtherInvoiceItemList"`
}

type PassInvoiceInfo struct {
	// 通行费车牌号
	NumberPlate *string `json:"NumberPlate,omitnil" name:"NumberPlate"`

	// 通行费类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 通行日期起
	PassDateBegin *string `json:"PassDateBegin,omitnil" name:"PassDateBegin"`

	// 通行日期止
	PassDateEnd *string `json:"PassDateEnd,omitnil" name:"PassDateEnd"`

	// 税收分类编码
	TaxClassifyCode *string `json:"TaxClassifyCode,omitnil" name:"TaxClassifyCode"`
}

// Predefined struct for user
type PassportOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 默认填写CN
	// 支持中国大陆地区护照。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type PassportOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 默认填写CN
	// 支持中国大陆地区护照。
	Type *string `json:"Type,omitnil" name:"Type"`
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

// Predefined struct for user
type PassportOCRResponseParams struct {
	// 国家码
	Country *string `json:"Country,omitnil" name:"Country"`

	// 护照号
	PassportNo *string `json:"PassportNo,omitnil" name:"PassportNo"`

	// 性别
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 国籍
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// 出生日期
	BirthDate *string `json:"BirthDate,omitnil" name:"BirthDate"`

	// 出生地点
	BirthPlace *string `json:"BirthPlace,omitnil" name:"BirthPlace"`

	// 签发日期
	IssueDate *string `json:"IssueDate,omitnil" name:"IssueDate"`

	// 签发地点
	IssuePlace *string `json:"IssuePlace,omitnil" name:"IssuePlace"`

	// 有效期
	ExpiryDate *string `json:"ExpiryDate,omitnil" name:"ExpiryDate"`

	// 持证人签名
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 最下方第一行 MRZ Code 序列
	CodeSet *string `json:"CodeSet,omitnil" name:"CodeSet"`

	// 最下方第二行 MRZ Code 序列
	CodeCrc *string `json:"CodeCrc,omitnil" name:"CodeCrc"`

	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 姓
	FamilyName *string `json:"FamilyName,omitnil" name:"FamilyName"`

	// 名
	FirstName *string `json:"FirstName,omitnil" name:"FirstName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PassportOCRResponse struct {
	*tchttp.BaseResponse
	Response *PassportOCRResponseParams `json:"Response"`
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

type PassportRecognizeInfos struct {
	// 证件类型（护照信息页识别结果）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 发行国家（护照信息页识别结果）
	IssuingCountry *string `json:"IssuingCountry,omitnil" name:"IssuingCountry"`

	// 护照号码（护照信息页识别结果）
	PassportID *string `json:"PassportID,omitnil" name:"PassportID"`

	// 姓（护照信息页识别结果）
	Surname *string `json:"Surname,omitnil" name:"Surname"`

	// 名（护照信息页识别结果）
	GivenName *string `json:"GivenName,omitnil" name:"GivenName"`

	// 姓名（护照信息页识别结果）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 国籍信息（护照信息页识别结果）
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// 出生日期（护照信息页识别结果）
	DateOfBirth *string `json:"DateOfBirth,omitnil" name:"DateOfBirth"`

	// 性别（护照信息页识别结果）
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 发行日期（护照信息页识别结果）
	DateOfIssuance *string `json:"DateOfIssuance,omitnil" name:"DateOfIssuance"`

	// 截止日期（护照信息页识别结果）
	DateOfExpiration *string `json:"DateOfExpiration,omitnil" name:"DateOfExpiration"`

	// 持证人签名（护照信息页识别结果）
	// 
	// 仅中国大陆护照支持返回此字段，港澳台及境外护照不支持
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 签发地点（护照信息页识别结果）
	// 
	// 仅中国大陆护照支持返回此字段，港澳台及境外护照不支持
	IssuePlace *string `json:"IssuePlace,omitnil" name:"IssuePlace"`

	// 签发机关（护照信息页识别结果）
	// 
	// 仅中国大陆护照支持返回此字段，港澳台及境外护照不支持
	IssuingAuthority *string `json:"IssuingAuthority,omitnil" name:"IssuingAuthority"`
}

// Predefined struct for user
type PermitOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type PermitOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type PermitOCRResponseParams struct {
	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 英文姓名
	EnglishName *string `json:"EnglishName,omitnil" name:"EnglishName"`

	// 证件号
	Number *string `json:"Number,omitnil" name:"Number"`

	// 性别
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 有效期限
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

	// 签发机关
	IssueAuthority *string `json:"IssueAuthority,omitnil" name:"IssueAuthority"`

	// 签发地点
	IssueAddress *string `json:"IssueAddress,omitnil" name:"IssueAddress"`

	// 出生日期
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PermitOCRResponse struct {
	*tchttp.BaseResponse
	Response *PermitOCRResponseParams `json:"Response"`
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
	LeftTop *Coord `json:"LeftTop,omitnil" name:"LeftTop"`

	// 右上顶点坐标
	RightTop *Coord `json:"RightTop,omitnil" name:"RightTop"`

	// 右下顶点坐标
	RightBottom *Coord `json:"RightBottom,omitnil" name:"RightBottom"`

	// 左下顶点坐标
	LeftBottom *Coord `json:"LeftBottom,omitnil" name:"LeftBottom"`
}

// Predefined struct for user
type PropOwnerCertOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type PropOwnerCertOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type PropOwnerCertOCRResponseParams struct {
	// 房地产权利人
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 共有情况
	Possession *string `json:"Possession,omitnil" name:"Possession"`

	// 登记时间
	RegisterTime *string `json:"RegisterTime,omitnil" name:"RegisterTime"`

	// 规划用途
	Purpose *string `json:"Purpose,omitnil" name:"Purpose"`

	// 房屋性质
	Nature *string `json:"Nature,omitnil" name:"Nature"`

	// 房地坐落
	Location *string `json:"Location,omitnil" name:"Location"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PropOwnerCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *PropOwnerCertOCRResponseParams `json:"Response"`
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
	Wide *int64 `json:"Wide,omitnil" name:"Wide"`

	// 高
	High *int64 `json:"High,omitnil" name:"High"`
}

// Predefined struct for user
type QrcodeOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，支持PNG、JPG、JPEG格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，支持PNG、JPG、JPEG格式。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type QrcodeOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，支持PNG、JPG、JPEG格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，支持PNG、JPG、JPEG格式。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type QrcodeOCRResponseParams struct {
	// 二维码/条形码识别结果信息，具体内容请点击左侧链接。
	CodeResults []*QrcodeResultsInfo `json:"CodeResults,omitnil" name:"CodeResults"`

	// 图片大小，具体内容请点击左侧链接。
	ImgSize *QrcodeImgSize `json:"ImgSize,omitnil" name:"ImgSize"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QrcodeOCRResponse struct {
	*tchttp.BaseResponse
	Response *QrcodeOCRResponseParams `json:"Response"`
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
	LeftTop *Coord `json:"LeftTop,omitnil" name:"LeftTop"`

	// 右上顶点坐标（如果是条形码，X和Y都为-1）
	RightTop *Coord `json:"RightTop,omitnil" name:"RightTop"`

	// 右下顶点坐标（如果是条形码，X和Y都为-1）
	RightBottom *Coord `json:"RightBottom,omitnil" name:"RightBottom"`

	// 左下顶点坐标（如果是条形码，X和Y都为-1）
	LeftBottom *Coord `json:"LeftBottom,omitnil" name:"LeftBottom"`
}

type QrcodeResultsInfo struct {
	// 类型（二维码、条形码）
	TypeName *string `json:"TypeName,omitnil" name:"TypeName"`

	// 二维码/条形码包含的地址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 二维码/条形码坐标
	Position *QrcodePositionObj `json:"Position,omitnil" name:"Position"`
}

type QuestionBlockObj struct {
	// 数学试题识别结构化信息数组
	QuestionArr []*QuestionObj `json:"QuestionArr,omitnil" name:"QuestionArr"`

	// 题目主体区域检测框在图片中的像素坐标
	QuestionBboxCoord *Rect `json:"QuestionBboxCoord,omitnil" name:"QuestionBboxCoord"`
}

type QuestionObj struct {
	// 题号
	QuestionTextNo *string `json:"QuestionTextNo,omitnil" name:"QuestionTextNo"`

	// 题型：
	// 1: "选择题"
	// 2: "填空题"
	// 3: "解答题"
	QuestionTextType *int64 `json:"QuestionTextType,omitnil" name:"QuestionTextType"`

	// 题干
	QuestionText *string `json:"QuestionText,omitnil" name:"QuestionText"`

	// 选择题选项，包含1个或多个option
	QuestionOptions *string `json:"QuestionOptions,omitnil" name:"QuestionOptions"`

	// 所有子题的question属性
	QuestionSubquestion *string `json:"QuestionSubquestion,omitnil" name:"QuestionSubquestion"`

	// 示意图检测框在的图片中的像素坐标
	QuestionImageCoords []*Rect `json:"QuestionImageCoords,omitnil" name:"QuestionImageCoords"`
}

type QuotaInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 是否有公司印章（0：没有，1：有）
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`
}

// Predefined struct for user
type QuotaInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type QuotaInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuotaInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuotaInvoiceOCRResponseParams struct {
	// 发票号码
	InvoiceNum *string `json:"InvoiceNum,omitnil" name:"InvoiceNum"`

	// 发票代码
	InvoiceCode *string `json:"InvoiceCode,omitnil" name:"InvoiceCode"`

	// 大写金额
	Rate *string `json:"Rate,omitnil" name:"Rate"`

	// 小写金额
	RateNum *string `json:"RateNum,omitnil" name:"RateNum"`

	// 发票消费类型
	InvoiceType *string `json:"InvoiceType,omitnil" name:"InvoiceType"`

	// 省
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil" name:"City"`

	// 是否有公司印章（1有 0无 空为识别不出）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasStamp *string `json:"HasStamp,omitnil" name:"HasStamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QuotaInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *QuotaInvoiceOCRResponseParams `json:"Response"`
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

type RailwayTicketInfo struct {
	// 电子发票类型
	TypeOfVoucher *string `json:"TypeOfVoucher,omitnil" name:"TypeOfVoucher"`

	// 电子客票号
	ElectronicTicketNum *string `json:"ElectronicTicketNum,omitnil" name:"ElectronicTicketNum"`

	// 开票日期
	DateOfIssue *string `json:"DateOfIssue,omitnil" name:"DateOfIssue"`

	// 售票或退票类型
	TypeOfBusiness *string `json:"TypeOfBusiness,omitnil" name:"TypeOfBusiness"`

	// 始发站
	DepartureStation *string `json:"DepartureStation,omitnil" name:"DepartureStation"`

	// 始发站英文
	PhonicsOfDepartureStation *string `json:"PhonicsOfDepartureStation,omitnil" name:"PhonicsOfDepartureStation"`

	// 到达站
	DestinationStation *string `json:"DestinationStation,omitnil" name:"DestinationStation"`

	// 到达站英文
	PhonicsOfDestinationStation *string `json:"PhonicsOfDestinationStation,omitnil" name:"PhonicsOfDestinationStation"`

	// 火车号
	TrainNumber *string `json:"TrainNumber,omitnil" name:"TrainNumber"`

	// 火车出发日期
	TravelDate *string `json:"TravelDate,omitnil" name:"TravelDate"`

	// 始发时间
	DepartureTime *string `json:"DepartureTime,omitnil" name:"DepartureTime"`

	// 空调特点
	AirConditioningCharacteristics *string `json:"AirConditioningCharacteristics,omitnil" name:"AirConditioningCharacteristics"`

	// 座位类型
	SeatLevel *string `json:"SeatLevel,omitnil" name:"SeatLevel"`

	// 火车第几车
	Carriage *string `json:"Carriage,omitnil" name:"Carriage"`

	// 座位号
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// 票价
	Fare *string `json:"Fare,omitnil" name:"Fare"`

	// 发票号码
	ElectronicInvoiceRailwayETicketNumber *string `json:"ElectronicInvoiceRailwayETicketNumber,omitnil" name:"ElectronicInvoiceRailwayETicketNumber"`

	// 身份证号
	IdNumber *string `json:"IdNumber,omitnil" name:"IdNumber"`

	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 金额
	TotalAmountExcludingTax *string `json:"TotalAmountExcludingTax,omitnil" name:"TotalAmountExcludingTax"`

	// 税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 税额
	TaxAmount *string `json:"TaxAmount,omitnil" name:"TaxAmount"`

	// 购买方名称
	NameOfPurchaser *string `json:"NameOfPurchaser,omitnil" name:"NameOfPurchaser"`

	// 统一社会信用代码
	UnifiedSocialCreditCodeOfPurchaser *string `json:"UnifiedSocialCreditCodeOfPurchaser,omitnil" name:"UnifiedSocialCreditCodeOfPurchaser"`

	// 原发票号码
	NumberOfOriginalInvoice *string `json:"NumberOfOriginalInvoice,omitnil" name:"NumberOfOriginalInvoice"`
}

// Predefined struct for user
type RecognizeContainerOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizeContainerOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type RecognizeContainerOCRResponseParams struct {
	// 集装箱箱号
	ContainerId *string `json:"ContainerId,omitnil" name:"ContainerId"`

	// 集装箱类型
	ContainerType *string `json:"ContainerType,omitnil" name:"ContainerType"`

	// 集装箱总重量，单位：千克（KG）
	GrossKG *string `json:"GrossKG,omitnil" name:"GrossKG"`

	// 集装箱总重量，单位：磅（LB）
	GrossLB *string `json:"GrossLB,omitnil" name:"GrossLB"`

	// 集装箱有效承重，单位：千克（KG）
	PayloadKG *string `json:"PayloadKG,omitnil" name:"PayloadKG"`

	// 集装箱有效承重，单位：磅（LB）
	PayloadLB *string `json:"PayloadLB,omitnil" name:"PayloadLB"`

	// 集装箱容量，单位：立方米
	CapacityM3 *string `json:"CapacityM3,omitnil" name:"CapacityM3"`

	// 集装箱容量，单位：立英尺
	CapacityFT3 *string `json:"CapacityFT3,omitnil" name:"CapacityFT3"`

	// 告警码
	// -9926	集装箱箱号不完整或者不清晰
	// -9927	集装箱类型不完整或者不清晰
	Warn []*int64 `json:"Warn,omitnil" name:"Warn"`

	// 集装箱自身重量，单位：千克（KG）
	TareKG *string `json:"TareKG,omitnil" name:"TareKG"`

	// 集装箱自身重量，单位：磅（LB）
	TareLB *string `json:"TareLB,omitnil" name:"TareLB"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeContainerOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeContainerOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type RecognizeEncryptedIDCardOCRRequestParams struct {
	// 请求体被加密后的密文，本接口只支持加密传输
	EncryptedBody *string `json:"EncryptedBody,omitnil" name:"EncryptedBody"`

	// 敏感数据加密信息。对传入信息有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil" name:"Encryption"`

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT：身份证有照片的一面（人像面），
	// BACK：身份证有国徽的一面（国徽面），
	// 该参数如果不填，将为您自动判断身份证正反面。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`

	// 以下可选字段均为bool 类型，默认false：
	// CropIdCard，身份证照片裁剪（去掉证件外多余的边缘、自动矫正拍摄角度）
	// CropPortrait，人像照片裁剪（自动抠取身份证头像区域）
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，疑似存在PS痕迹告警
	// TempIdWarn，临时身份证告警
	// InvalidDateWarn，身份证有效日期不合法告警
	// Quality，图片质量分数（评价图片的模糊程度）
	// MultiCardDetect，是否开启正反面同框识别（仅支持二代身份证正反页同框识别或临时身份证正反页同框识别）
	// ReflectWarn，是否开启反光检测
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitnil" name:"Config"`

	// 默认值为true，打开识别结果纠正开关。开关开启后，身份证号、出生日期、性别，三个字段会进行矫正补齐，统一结果输出；若关闭此开关，以上三个字段不会进行矫正补齐，保持原始识别结果输出，若原图出现篡改情况，这三个字段的识别结果可能会不统一。
	EnableRecognitionRectify *bool `json:"EnableRecognitionRectify,omitnil" name:"EnableRecognitionRectify"`

	// 默认值为false。
	// 
	// 此开关需要在反光检测开关开启下才会生效（即此开关生效的前提是config入参里的"ReflectWarn":true），若EnableReflectDetail设置为true，则会返回反光点覆盖区域详情。反光点覆盖区域详情分为四部分：人像照片位置、国徽位置、识别字段位置、其他位置。一个反光点允许覆盖多个区域，且一张图片可能存在多个反光点。
	EnableReflectDetail *bool `json:"EnableReflectDetail,omitnil" name:"EnableReflectDetail"`
}

type RecognizeEncryptedIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 请求体被加密后的密文，本接口只支持加密传输
	EncryptedBody *string `json:"EncryptedBody,omitnil" name:"EncryptedBody"`

	// 敏感数据加密信息。对传入信息有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil" name:"Encryption"`

	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT：身份证有照片的一面（人像面），
	// BACK：身份证有国徽的一面（国徽面），
	// 该参数如果不填，将为您自动判断身份证正反面。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`

	// 以下可选字段均为bool 类型，默认false：
	// CropIdCard，身份证照片裁剪（去掉证件外多余的边缘、自动矫正拍摄角度）
	// CropPortrait，人像照片裁剪（自动抠取身份证头像区域）
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，疑似存在PS痕迹告警
	// TempIdWarn，临时身份证告警
	// InvalidDateWarn，身份证有效日期不合法告警
	// Quality，图片质量分数（评价图片的模糊程度）
	// MultiCardDetect，是否开启正反面同框识别（仅支持二代身份证正反页同框识别或临时身份证正反页同框识别）
	// ReflectWarn，是否开启反光检测
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitnil" name:"Config"`

	// 默认值为true，打开识别结果纠正开关。开关开启后，身份证号、出生日期、性别，三个字段会进行矫正补齐，统一结果输出；若关闭此开关，以上三个字段不会进行矫正补齐，保持原始识别结果输出，若原图出现篡改情况，这三个字段的识别结果可能会不统一。
	EnableRecognitionRectify *bool `json:"EnableRecognitionRectify,omitnil" name:"EnableRecognitionRectify"`

	// 默认值为false。
	// 
	// 此开关需要在反光检测开关开启下才会生效（即此开关生效的前提是config入参里的"ReflectWarn":true），若EnableReflectDetail设置为true，则会返回反光点覆盖区域详情。反光点覆盖区域详情分为四部分：人像照片位置、国徽位置、识别字段位置、其他位置。一个反光点允许覆盖多个区域，且一张图片可能存在多个反光点。
	EnableReflectDetail *bool `json:"EnableReflectDetail,omitnil" name:"EnableReflectDetail"`
}

func (r *RecognizeEncryptedIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeEncryptedIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EncryptedBody")
	delete(f, "Encryption")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "CardSide")
	delete(f, "Config")
	delete(f, "EnableRecognitionRectify")
	delete(f, "EnableReflectDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeEncryptedIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeEncryptedIDCardOCRResponseParams struct {
	// 姓名（人像面）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 性别（人像面）
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 民族（人像面）
	Nation *string `json:"Nation,omitnil" name:"Nation"`

	// 出生日期（人像面）
	Birth *string `json:"Birth,omitnil" name:"Birth"`

	// 地址（人像面）
	Address *string `json:"Address,omitnil" name:"Address"`

	// 身份证号（人像面）
	IdNum *string `json:"IdNum,omitnil" name:"IdNum"`

	// 发证机关（国徽面）
	Authority *string `json:"Authority,omitnil" name:"Authority"`

	// 证件有效期（国徽面）
	ValidDate *string `json:"ValidDate,omitnil" name:"ValidDate"`

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
	// -9106	身份证疑似存在PS痕迹告警，
	// -9107       身份证反光告警。
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 反光点覆盖区域详情结果，具体内容请点击左侧链接
	ReflectDetailInfos []*ReflectDetailInfo `json:"ReflectDetailInfos,omitnil" name:"ReflectDetailInfos"`

	// 加密后的数据
	EncryptedBody *string `json:"EncryptedBody,omitnil" name:"EncryptedBody"`

	// 敏感数据加密信息
	Encryption *Encryption `json:"Encryption,omitnil" name:"Encryption"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeEncryptedIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeEncryptedIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeEncryptedIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeEncryptedIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeForeignPermanentResidentIdCardRequestParams struct {
	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	// 示例值：https://ocr-demo-1254418846.cos.ap-guangzhou.myqcloud.com/docume
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// 需要识别的PDF页面的对应页码，传入时仅支持PDF单页识别，当上传文件为PDF且EnablePdf参数值为true时有效，默认值为1。
	// 示例值：1
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type RecognizeForeignPermanentResidentIdCardRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	// 示例值：https://ocr-demo-1254418846.cos.ap-guangzhou.myqcloud.com/docume
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// 需要识别的PDF页面的对应页码，传入时仅支持PDF单页识别，当上传文件为PDF且EnablePdf参数值为true时有效，默认值为1。
	// 示例值：1
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

func (r *RecognizeForeignPermanentResidentIdCardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeForeignPermanentResidentIdCardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "EnablePdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeForeignPermanentResidentIdCardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeForeignPermanentResidentIdCardResponseParams struct {
	// 中文姓名。
	CnName *string `json:"CnName,omitnil" name:"CnName"`

	// 英文名。
	EnName *string `json:"EnName,omitnil" name:"EnName"`

	// 性别。
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 出生日期。规范格式为 XXXX年XX月XX日。
	DateOfBirth *string `json:"DateOfBirth,omitnil" name:"DateOfBirth"`

	// 国籍。
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// 有效期限。
	PeriodOfValidity *string `json:"PeriodOfValidity,omitnil" name:"PeriodOfValidity"`

	// 证件号码。
	No *string `json:"No,omitnil" name:"No"`

	// 曾持证件号码。
	PreviousNumber *string `json:"PreviousNumber,omitnil" name:"PreviousNumber"`

	// 签发机关。
	IssuedAuthority *string `json:"IssuedAuthority,omitnil" name:"IssuedAuthority"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeForeignPermanentResidentIdCardResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeForeignPermanentResidentIdCardResponseParams `json:"Response"`
}

func (r *RecognizeForeignPermanentResidentIdCardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeForeignPermanentResidentIdCardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeGeneralInvoiceRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 8M。图片下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 8M。图片下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 需要识别的票据类型列表，为空或不填表示识别全部类型。当传入单个类型时，图片均采用该票类型进行处理。
	// 暂不支持多个参数进行局部控制。
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
	// 15：非税发票
	// 16：全电发票
	// 17：医疗发票
	// -1：其他发票
	Types []*int64 `json:"Types,omitnil" name:"Types"`

	// 是否开启其他票识别，默认值为true，开启后可支持其他发票的智能识别。	
	EnableOther *bool `json:"EnableOther,omitnil" name:"EnableOther"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// 需要识别的PDF页面的对应页码，传入时仅支持PDF单页识别，当上传文件为PDF且EnablePdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否开启PDF多页识别，默认值为false，开启后可同时支持多页PDF的识别返回，仅支持返回文件前30页。开启后EnablePdf和PdfPageNumber入参不进行控制。
	EnableMultiplePage *bool `json:"EnableMultiplePage,omitnil" name:"EnableMultiplePage"`

	// 是否返回切割图片base64，默认值为false。
	EnableCutImage *bool `json:"EnableCutImage,omitnil" name:"EnableCutImage"`

	// 是否打开字段坐标返回。默认为false。
	EnableItemPolygon *bool `json:"EnableItemPolygon,omitnil" name:"EnableItemPolygon"`
}

type RecognizeGeneralInvoiceRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 8M。图片下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 8M。图片下载时间不超过 3 秒。
	// 支持的图片像素：单边介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 需要识别的票据类型列表，为空或不填表示识别全部类型。当传入单个类型时，图片均采用该票类型进行处理。
	// 暂不支持多个参数进行局部控制。
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
	// 15：非税发票
	// 16：全电发票
	// 17：医疗发票
	// -1：其他发票
	Types []*int64 `json:"Types,omitnil" name:"Types"`

	// 是否开启其他票识别，默认值为true，开启后可支持其他发票的智能识别。	
	EnableOther *bool `json:"EnableOther,omitnil" name:"EnableOther"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// 需要识别的PDF页面的对应页码，传入时仅支持PDF单页识别，当上传文件为PDF且EnablePdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否开启PDF多页识别，默认值为false，开启后可同时支持多页PDF的识别返回，仅支持返回文件前30页。开启后EnablePdf和PdfPageNumber入参不进行控制。
	EnableMultiplePage *bool `json:"EnableMultiplePage,omitnil" name:"EnableMultiplePage"`

	// 是否返回切割图片base64，默认值为false。
	EnableCutImage *bool `json:"EnableCutImage,omitnil" name:"EnableCutImage"`

	// 是否打开字段坐标返回。默认为false。
	EnableItemPolygon *bool `json:"EnableItemPolygon,omitnil" name:"EnableItemPolygon"`
}

func (r *RecognizeGeneralInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeGeneralInvoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Types")
	delete(f, "EnableOther")
	delete(f, "EnablePdf")
	delete(f, "PdfPageNumber")
	delete(f, "EnableMultiplePage")
	delete(f, "EnableCutImage")
	delete(f, "EnableItemPolygon")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeGeneralInvoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeGeneralInvoiceResponseParams struct {
	// 混贴票据识别结果，具体内容请点击左侧链接。
	MixedInvoiceItems []*InvoiceItem `json:"MixedInvoiceItems,omitnil" name:"MixedInvoiceItems"`

	// PDF文件总页码
	TotalPDFCount *int64 `json:"TotalPDFCount,omitnil" name:"TotalPDFCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeGeneralInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeGeneralInvoiceResponseParams `json:"Response"`
}

func (r *RecognizeGeneralInvoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeGeneralInvoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeHealthCodeOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 需要识别的健康码类型列表，为空或不填表示默认为自动识别。
	// 0:自动识别
	Type *int64 `json:"Type,omitnil" name:"Type"`
}

type RecognizeHealthCodeOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 需要识别的健康码类型列表，为空或不填表示默认为自动识别。
	// 0:自动识别
	Type *int64 `json:"Type,omitnil" name:"Type"`
}

func (r *RecognizeHealthCodeOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeHealthCodeOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeHealthCodeOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeHealthCodeOCRResponseParams struct {
	// 持码人姓名，如：王*（允许返回空值）
	Name *string `json:"Name,omitnil" name:"Name"`

	// 持码人身份证号，如：11**************01（允许返回空值）
	IDNumber *string `json:"IDNumber,omitnil" name:"IDNumber"`

	// 健康码更新时间（允许返回空值）
	Time *string `json:"Time,omitnil" name:"Time"`

	// 健康码颜色：绿色、黄色、红色（允许返回空值）
	Color *string `json:"Color,omitnil" name:"Color"`

	// 核酸检测间隔时长（允许返回空值）
	TestingInterval *string `json:"TestingInterval,omitnil" name:"TestingInterval"`

	// 核酸检测结果：阴性、阳性、暂无核酸检测记录（允许返回空值）
	TestingResult *string `json:"TestingResult,omitnil" name:"TestingResult"`

	// 核酸检测时间（允许返回空值）
	TestingTime *string `json:"TestingTime,omitnil" name:"TestingTime"`

	// 疫苗接种信息，返回接种针数或接种情况（允许返回空值）
	Vaccination *string `json:"Vaccination,omitnil" name:"Vaccination"`

	// 场所名称（允许返回空值）
	SpotName *string `json:"SpotName,omitnil" name:"SpotName"`

	// 疫苗接种时间
	VaccinationTime *string `json:"VaccinationTime,omitnil" name:"VaccinationTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeHealthCodeOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeHealthCodeOCRResponseParams `json:"Response"`
}

func (r *RecognizeHealthCodeOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeHealthCodeOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeIndonesiaIDCardOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 场景参数，默认值为V1
	// 可选值：
	// V1
	// V2
	Scene *string `json:"Scene,omitnil" name:"Scene"`
}

type RecognizeIndonesiaIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 场景参数，默认值为V1
	// 可选值：
	// V1
	// V2
	Scene *string `json:"Scene,omitnil" name:"Scene"`
}

func (r *RecognizeIndonesiaIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeIndonesiaIDCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	delete(f, "Scene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeIndonesiaIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeIndonesiaIDCardOCRResponseParams struct {
	// 证件号码
	NIK *string `json:"NIK,omitnil" name:"NIK"`

	// 姓名
	Nama *string `json:"Nama,omitnil" name:"Nama"`

	// 出生地/出生时间
	TempatTglLahir *string `json:"TempatTglLahir,omitnil" name:"TempatTglLahir"`

	// 性别
	JenisKelamin *string `json:"JenisKelamin,omitnil" name:"JenisKelamin"`

	// 血型
	GolDarah *string `json:"GolDarah,omitnil" name:"GolDarah"`

	// 地址
	Alamat *string `json:"Alamat,omitnil" name:"Alamat"`

	// 街道
	RTRW *string `json:"RTRW,omitnil" name:"RTRW"`

	// 村
	KelDesa *string `json:"KelDesa,omitnil" name:"KelDesa"`

	// 地区
	Kecamatan *string `json:"Kecamatan,omitnil" name:"Kecamatan"`

	// 宗教信仰
	Agama *string `json:"Agama,omitnil" name:"Agama"`

	// 婚姻状况
	StatusPerkawinan *string `json:"StatusPerkawinan,omitnil" name:"StatusPerkawinan"`

	// 职业
	Perkerjaan *string `json:"Perkerjaan,omitnil" name:"Perkerjaan"`

	// 国籍
	KewargaNegaraan *string `json:"KewargaNegaraan,omitnil" name:"KewargaNegaraan"`

	// 身份证有效期限
	BerlakuHingga *string `json:"BerlakuHingga,omitnil" name:"BerlakuHingga"`

	// 发证日期
	IssuedDate *string `json:"IssuedDate,omitnil" name:"IssuedDate"`

	// 人像截图
	Photo *string `json:"Photo,omitnil" name:"Photo"`

	// 省份，Scene为V2时支持识别
	Provinsi *string `json:"Provinsi,omitnil" name:"Provinsi"`

	// 城市，Scene为V2时支持识别
	Kota *string `json:"Kota,omitnil" name:"Kota"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeIndonesiaIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeIndonesiaIDCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeIndonesiaIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeIndonesiaIDCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMedicalInvoiceOCRRequestParams struct {
	// 图片的Base64 值。
	// 支持的文件格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载文件经Base64编码后不超过 7M。文件下载时间不超过 3 秒。
	// 输入参数 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的Url 地址。
	// 支持的文件格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载文件经 Base64 编码后不超过 7M。文件下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否需要返回识别出的文本行在原图上的四点坐标，默认不返回
	ReturnVertex *bool `json:"ReturnVertex,omitnil" name:"ReturnVertex"`

	// 是否需要返回识别出的文本行在旋转纠正之后的图像中的四点坐标，默认不返回
	ReturnCoord *bool `json:"ReturnCoord,omitnil" name:"ReturnCoord"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type RecognizeMedicalInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的Base64 值。
	// 支持的文件格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载文件经Base64编码后不超过 7M。文件下载时间不超过 3 秒。
	// 输入参数 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的Url 地址。
	// 支持的文件格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载文件经 Base64 编码后不超过 7M。文件下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否需要返回识别出的文本行在原图上的四点坐标，默认不返回
	ReturnVertex *bool `json:"ReturnVertex,omitnil" name:"ReturnVertex"`

	// 是否需要返回识别出的文本行在旋转纠正之后的图像中的四点坐标，默认不返回
	ReturnCoord *bool `json:"ReturnCoord,omitnil" name:"ReturnCoord"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

func (r *RecognizeMedicalInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMedicalInvoiceOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnVertex")
	delete(f, "ReturnCoord")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeMedicalInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeMedicalInvoiceOCRResponseParams struct {
	// 识别出的字段信息
	MedicalInvoiceInfos []*MedicalInvoiceInfo `json:"MedicalInvoiceInfos,omitnil" name:"MedicalInvoiceInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeMedicalInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeMedicalInvoiceOCRResponseParams `json:"Response"`
}

func (r *RecognizeMedicalInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeMedicalInvoiceOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeOnlineTaxiItineraryOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type RecognizeOnlineTaxiItineraryOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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

// Predefined struct for user
type RecognizeOnlineTaxiItineraryOCRResponseParams struct {
	// 网约车行程单识别结果，具体内容请点击左侧链接。
	OnlineTaxiItineraryInfos []*OnlineTaxiItineraryInfo `json:"OnlineTaxiItineraryInfos,omitnil" name:"OnlineTaxiItineraryInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeOnlineTaxiItineraryOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeOnlineTaxiItineraryOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type RecognizePhilippinesDrivingLicenseOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

type RecognizePhilippinesDrivingLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

func (r *RecognizePhilippinesDrivingLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesDrivingLicenseOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesDrivingLicenseOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesDrivingLicenseOCRResponseParams struct {
	// 人像照片Base64后的结果
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// 姓名
	Name *TextDetectionResult `json:"Name,omitnil" name:"Name"`

	// 姓氏
	LastName *TextDetectionResult `json:"LastName,omitnil" name:"LastName"`

	// 首姓名
	FirstName *TextDetectionResult `json:"FirstName,omitnil" name:"FirstName"`

	// 中间姓名
	MiddleName *TextDetectionResult `json:"MiddleName,omitnil" name:"MiddleName"`

	// 国籍
	Nationality *TextDetectionResult `json:"Nationality,omitnil" name:"Nationality"`

	// 性别
	Sex *TextDetectionResult `json:"Sex,omitnil" name:"Sex"`

	// 地址
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// 证号
	LicenseNo *TextDetectionResult `json:"LicenseNo,omitnil" name:"LicenseNo"`

	// 有效期
	ExpiresDate *TextDetectionResult `json:"ExpiresDate,omitnil" name:"ExpiresDate"`

	// 机构代码
	AgencyCode *TextDetectionResult `json:"AgencyCode,omitnil" name:"AgencyCode"`

	// 出生日期
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesDrivingLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesDrivingLicenseOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesDrivingLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesDrivingLicenseOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesSssIDOCRRequestParams struct {
	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizePhilippinesSssIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *RecognizePhilippinesSssIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesSssIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesSssIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesSssIDOCRResponseParams struct {
	// 人像照片Base64后的结果
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// 编号
	LicenseNumber *TextDetectionResult `json:"LicenseNumber,omitnil" name:"LicenseNumber"`

	// 姓名
	FullName *TextDetectionResult `json:"FullName,omitnil" name:"FullName"`

	// 生日
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesSssIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesSssIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesSssIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesSssIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesTinIDOCRRequestParams struct {
	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizePhilippinesTinIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *RecognizePhilippinesTinIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesTinIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesTinIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesTinIDOCRResponseParams struct {
	// 人像照片Base64后的结果
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// 编码
	LicenseNumber *TextDetectionResult `json:"LicenseNumber,omitnil" name:"LicenseNumber"`

	// 姓名
	FullName *TextDetectionResult `json:"FullName,omitnil" name:"FullName"`

	// 地址
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// 生日
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// 发证日期
	IssueDate *TextDetectionResult `json:"IssueDate,omitnil" name:"IssueDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesTinIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesTinIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesTinIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesTinIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesUMIDOCRRequestParams struct {
	// 图片的 Base64 值。 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

type RecognizePhilippinesUMIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`
}

func (r *RecognizePhilippinesUMIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesUMIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ReturnHeadImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesUMIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesUMIDOCRResponseParams struct {
	// 姓
	Surname *TextDetectionResult `json:"Surname,omitnil" name:"Surname"`

	// 中间名
	MiddleName *TextDetectionResult `json:"MiddleName,omitnil" name:"MiddleName"`

	// 名
	GivenName *TextDetectionResult `json:"GivenName,omitnil" name:"GivenName"`

	// 地址
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// 生日
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// crn码
	CRN *TextDetectionResult `json:"CRN,omitnil" name:"CRN"`

	// 性别
	Sex *TextDetectionResult `json:"Sex,omitnil" name:"Sex"`

	// 人像照片Base64后的结果
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesUMIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesUMIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesUMIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesUMIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesVoteIDOCRRequestParams struct {
	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizePhilippinesVoteIDOCRRequest struct {
	*tchttp.BaseRequest
	
	// 是否返回人像照片。
	ReturnHeadImage *bool `json:"ReturnHeadImage,omitnil" name:"ReturnHeadImage"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *RecognizePhilippinesVoteIDOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesVoteIDOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReturnHeadImage")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizePhilippinesVoteIDOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizePhilippinesVoteIDOCRResponseParams struct {
	// 人像照片Base64后的结果
	HeadPortrait *TextDetectionResult `json:"HeadPortrait,omitnil" name:"HeadPortrait"`

	// 菲律宾VoteID的VIN
	VIN *TextDetectionResult `json:"VIN,omitnil" name:"VIN"`

	// 姓名
	FirstName *TextDetectionResult `json:"FirstName,omitnil" name:"FirstName"`

	// 姓氏
	LastName *TextDetectionResult `json:"LastName,omitnil" name:"LastName"`

	// 出生日期
	Birthday *TextDetectionResult `json:"Birthday,omitnil" name:"Birthday"`

	// 婚姻状况
	CivilStatus *TextDetectionResult `json:"CivilStatus,omitnil" name:"CivilStatus"`

	// 国籍
	Citizenship *TextDetectionResult `json:"Citizenship,omitnil" name:"Citizenship"`

	// 地址
	Address *TextDetectionResult `json:"Address,omitnil" name:"Address"`

	// 地区
	PrecinctNo *TextDetectionResult `json:"PrecinctNo,omitnil" name:"PrecinctNo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizePhilippinesVoteIDOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizePhilippinesVoteIDOCRResponseParams `json:"Response"`
}

func (r *RecognizePhilippinesVoteIDOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizePhilippinesVoteIDOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTableAccurateOCRRequestParams struct {
	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片支持的像素范围：需介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片支持的像素范围：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定
	// 性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type RecognizeTableAccurateOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片支持的像素范围：需介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片支持的像素范围：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定
	// 性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

func (r *RecognizeTableAccurateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTableAccurateOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeTableAccurateOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTableAccurateOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableDetections []*TableInfo `json:"TableDetections,omitnil" name:"TableDetections"`

	// Base64 编码后的 Excel 数据。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 图片为PDF时，返回PDF的总页数，默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	PdfPageSize *int64 `json:"PdfPageSize,omitnil" name:"PdfPageSize"`

	// 图片旋转角度（角度制），文本的水平方向为0°，统一以逆时针方向旋转，逆时针为负，角度范围为-360°至0°。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeTableAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeTableAccurateOCRResponseParams `json:"Response"`
}

func (r *RecognizeTableAccurateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTableAccurateOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTableOCRRequestParams struct {
	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 语言，zh：中英文（默认）jap：日文
	TableLanguage *string `json:"TableLanguage,omitnil" name:"TableLanguage"`
}

type RecognizeTableOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片/PDF的 Base64 值。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 语言，zh：中英文（默认）jap：日文
	TableLanguage *string `json:"TableLanguage,omitnil" name:"TableLanguage"`
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

// Predefined struct for user
type RecognizeTableOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	TableDetections []*TableDetectInfo `json:"TableDetections,omitnil" name:"TableDetections"`

	// Base64 编码后的 Excel 数据。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 图片为PDF时，返回PDF的总页数，默认为0
	PdfPageSize *int64 `json:"PdfPageSize,omitnil" name:"PdfPageSize"`

	// 图片旋转角度（角度制），文本的水平方向为0°，统一以逆时针方向旋转，逆时针为负，角度范围为-360°至0°。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeTableOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeTableOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type RecognizeThaiIDCardOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片开关。默认为false，不返回泰国身份证头像照片的base64编码。
	// 设置为true时，返回旋转矫正后的泰国身份证头像照片的base64编码
	CropPortrait *bool `json:"CropPortrait,omitnil" name:"CropPortrait"`
}

type RecognizeThaiIDCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片开关。默认为false，不返回泰国身份证头像照片的base64编码。
	// 设置为true时，返回旋转矫正后的泰国身份证头像照片的base64编码
	CropPortrait *bool `json:"CropPortrait,omitnil" name:"CropPortrait"`
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
	delete(f, "CropPortrait")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeThaiIDCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeThaiIDCardOCRResponseParams struct {
	// 身份证号码
	ID *string `json:"ID,omitnil" name:"ID"`

	// 泰文姓名
	ThaiName *string `json:"ThaiName,omitnil" name:"ThaiName"`

	// 英文姓名
	EnFirstName *string `json:"EnFirstName,omitnil" name:"EnFirstName"`

	// 地址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 出生日期
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// 签发日期
	IssueDate *string `json:"IssueDate,omitnil" name:"IssueDate"`

	// 到期日期
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 英文姓名
	EnLastName *string `json:"EnLastName,omitnil" name:"EnLastName"`

	// 证件人像照片抠取
	PortraitImage *string `json:"PortraitImage,omitnil" name:"PortraitImage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeThaiIDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeThaiIDCardOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type RecognizeTravelCardOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RecognizeTravelCardOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

func (r *RecognizeTravelCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTravelCardOCRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeTravelCardOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeTravelCardOCRResponseParams struct {
	// 行程卡更新时间，格式为：XXXX.XX.XX XX:XX:XX
	Time *string `json:"Time,omitnil" name:"Time"`

	// 行程卡颜色：绿色、黄色、红色
	Color *string `json:"Color,omitnil" name:"Color"`

	// 7天内到达或途经的城市（自2022年7月8日起，通信行程卡查询结果的覆盖时间范围由“14天”调整为“7天”）
	ReachedCity []*string `json:"ReachedCity,omitnil" name:"ReachedCity"`

	// 7天内到达或途径存在中高风险地区的城市（自2022年6月29日起，通信行程卡取消“星号”标记，改字段将返回空值）
	RiskArea []*string `json:"RiskArea,omitnil" name:"RiskArea"`

	// 电话号码
	Telephone *string `json:"Telephone,omitnil" name:"Telephone"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeTravelCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeTravelCardOCRResponseParams `json:"Response"`
}

func (r *RecognizeTravelCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeTravelCardOCRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Rect struct {
	// 左上角x
	X *int64 `json:"X,omitnil" name:"X"`

	// 左上角y
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 宽度
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 高度
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type ReflectDetailInfo struct {
	// NationalEmblem 国徽位置
	// Portrait 人像照片位置
	// RecognitionField 识别字段位置
	// Others 其他位置
	Position *string `json:"Position,omitnil" name:"Position"`
}

// Predefined struct for user
type ResidenceBookletOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type ResidenceBookletOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type ResidenceBookletOCRResponseParams struct {
	// 户号
	HouseholdNumber *string `json:"HouseholdNumber,omitnil" name:"HouseholdNumber"`

	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 性别
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 出生地
	BirthPlace *string `json:"BirthPlace,omitnil" name:"BirthPlace"`

	// 民族
	Nation *string `json:"Nation,omitnil" name:"Nation"`

	// 籍贯
	NativePlace *string `json:"NativePlace,omitnil" name:"NativePlace"`

	// 出生日期
	BirthDate *string `json:"BirthDate,omitnil" name:"BirthDate"`

	// 公民身份证件编号
	IdCardNumber *string `json:"IdCardNumber,omitnil" name:"IdCardNumber"`

	// 文化程度
	EducationDegree *string `json:"EducationDegree,omitnil" name:"EducationDegree"`

	// 服务处所
	ServicePlace *string `json:"ServicePlace,omitnil" name:"ServicePlace"`

	// 户别
	Household *string `json:"Household,omitnil" name:"Household"`

	// 住址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 承办人签章文字
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 签发日期
	IssueDate *string `json:"IssueDate,omitnil" name:"IssueDate"`

	// 户主页编号
	HomePageNumber *string `json:"HomePageNumber,omitnil" name:"HomePageNumber"`

	// 户主姓名
	HouseholderName *string `json:"HouseholderName,omitnil" name:"HouseholderName"`

	// 户主或与户主关系
	Relationship *string `json:"Relationship,omitnil" name:"Relationship"`

	// 本市（县）其他住址
	OtherAddresses *string `json:"OtherAddresses,omitnil" name:"OtherAddresses"`

	// 宗教信仰
	ReligiousBelief *string `json:"ReligiousBelief,omitnil" name:"ReligiousBelief"`

	// 身高
	Height *string `json:"Height,omitnil" name:"Height"`

	// 血型
	BloodType *string `json:"BloodType,omitnil" name:"BloodType"`

	// 婚姻状况
	MaritalStatus *string `json:"MaritalStatus,omitnil" name:"MaritalStatus"`

	// 兵役状况
	VeteranStatus *string `json:"VeteranStatus,omitnil" name:"VeteranStatus"`

	// 职业
	Profession *string `json:"Profession,omitnil" name:"Profession"`

	// 何时由何地迁来本市(县)
	MoveToCityInformation *string `json:"MoveToCityInformation,omitnil" name:"MoveToCityInformation"`

	// 何时由何地迁来本址
	MoveToSiteInformation *string `json:"MoveToSiteInformation,omitnil" name:"MoveToSiteInformation"`

	// 登记日期
	RegistrationDate *string `json:"RegistrationDate,omitnil" name:"RegistrationDate"`

	// 曾用名
	FormerName *string `json:"FormerName,omitnil" name:"FormerName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResidenceBookletOCRResponse struct {
	*tchttp.BaseResponse
	Response *ResidenceBookletOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type RideHailingDriverLicenseOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RideHailingDriverLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type RideHailingDriverLicenseOCRResponseParams struct {
	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 证号，对应网约车驾驶证字段：证号/从业资格证号/驾驶员证号/身份证号
	LicenseNumber *string `json:"LicenseNumber,omitnil" name:"LicenseNumber"`

	// 有效起始日期
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 有效期截止时间，对应网约车驾驶证字段：有效期至/营运期限止
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 初始发证日期，对应网约车驾驶证字段：初始领证日期/发证日期
	ReleaseDate *string `json:"ReleaseDate,omitnil" name:"ReleaseDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RideHailingDriverLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *RideHailingDriverLicenseOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type RideHailingTransportLicenseOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type RideHailingTransportLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type RideHailingTransportLicenseOCRResponseParams struct {
	// 交运管许可字号。
	OperationLicenseNumber *string `json:"OperationLicenseNumber,omitnil" name:"OperationLicenseNumber"`

	// 车辆所有人，对应网约车运输证字段：车辆所有人/车主名称/业户名称。
	VehicleOwner *string `json:"VehicleOwner,omitnil" name:"VehicleOwner"`

	// 车牌号码，对应网约车运输证字段：车牌号码/车辆号牌。
	VehicleNumber *string `json:"VehicleNumber,omitnil" name:"VehicleNumber"`

	// 有效起始日期。
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 有效期截止时间，对应网约车运输证字段：有效期至/营运期限止。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 初始发证日期，对应网约车运输证字段：初始领证日期/发证日期。
	ReleaseDate *string `json:"ReleaseDate,omitnil" name:"ReleaseDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RideHailingTransportLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *RideHailingTransportLicenseOCRResponseParams `json:"Response"`
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

type SealInfo struct {
	// 印章主体内容
	SealBody *string `json:"SealBody,omitnil" name:"SealBody"`

	// 印章坐标
	Location *Rect `json:"Location,omitnil" name:"Location"`

	// 印章其它文本内容
	OtherTexts []*string `json:"OtherTexts,omitnil" name:"OtherTexts"`

	// 印章类型，表示为:
	// 圆形印章：0
	// 椭圆形印章：1
	// 方形印章：2
	// 菱形印章：3
	// 三角形印章：4
	SealShape *string `json:"SealShape,omitnil" name:"SealShape"`
}

// Predefined struct for user
type SealOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// 需要识别的PDF页面的对应页码，传入时仅支持PDF单页识别，当上传文件为PDF且EnablePdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type SealOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP、PDF格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	EnablePdf *bool `json:"EnablePdf,omitnil" name:"EnablePdf"`

	// 需要识别的PDF页面的对应页码，传入时仅支持PDF单页识别，当上传文件为PDF且EnablePdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "EnablePdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SealOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SealOCRResponseParams struct {
	// 印章内容
	SealBody *string `json:"SealBody,omitnil" name:"SealBody"`

	// 印章坐标
	Location *Rect `json:"Location,omitnil" name:"Location"`

	// 其它文本内容
	OtherTexts []*string `json:"OtherTexts,omitnil" name:"OtherTexts"`

	// 全部印章信息
	SealInfos []*SealInfo `json:"SealInfos,omitnil" name:"SealInfos"`

	// 印章类型，表示为：
	// 圆形印章：0
	// 椭圆形印章：1
	// 方形印章：2
	// 菱形印章：3
	// 三角形印章：4
	SealShape *string `json:"SealShape,omitnil" name:"SealShape"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SealOCRResponse struct {
	*tchttp.BaseResponse
	Response *SealOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type ShipInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type ShipInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ShipInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ShipInvoiceOCRResponseParams struct {
	// 轮船票识别结果，具体内容请点击左侧链接。
	ShipInvoiceInfos []*ShipInvoiceInfo `json:"ShipInvoiceInfos,omitnil" name:"ShipInvoiceInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ShipInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *ShipInvoiceOCRResponseParams `json:"Response"`
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

type ShippingInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 姓名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 出发车站
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// 到达车站
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// 票价
	Total *string `json:"Total,omitnil" name:"Total"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 币种
	CurrencyCode *string `json:"CurrencyCode,omitnil" name:"CurrencyCode"`
}

type SingleInvoiceInfo struct {
	// 识别出的字段名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 字段属于第几行，用于相同字段的排版，如发票明细表格项目，普通字段使用默认值为-1，表示无列排版。
	Row *int64 `json:"Row,omitnil" name:"Row"`
}

type SingleInvoiceItem struct {
	// 增值税专用发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatSpecialInvoice *VatInvoiceInfo `json:"VatSpecialInvoice,omitnil" name:"VatSpecialInvoice"`

	// 增值税普通发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatCommonInvoice *VatInvoiceInfo `json:"VatCommonInvoice,omitnil" name:"VatCommonInvoice"`

	// 增值税电子普通发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatElectronicCommonInvoice *VatInvoiceInfo `json:"VatElectronicCommonInvoice,omitnil" name:"VatElectronicCommonInvoice"`

	// 增值税电子专用发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatElectronicSpecialInvoice *VatInvoiceInfo `json:"VatElectronicSpecialInvoice,omitnil" name:"VatElectronicSpecialInvoice"`

	// 区块链电子发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatElectronicInvoiceBlockchain *VatInvoiceInfo `json:"VatElectronicInvoiceBlockchain,omitnil" name:"VatElectronicInvoiceBlockchain"`

	// 增值税电子普通发票(通行费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatElectronicInvoiceToll *VatInvoiceInfo `json:"VatElectronicInvoiceToll,omitnil" name:"VatElectronicInvoiceToll"`

	// 电子发票(专用发票)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatElectronicSpecialInvoiceFull *VatElectronicInfo `json:"VatElectronicSpecialInvoiceFull,omitnil" name:"VatElectronicSpecialInvoiceFull"`

	// 电子发票(普通发票)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatElectronicInvoiceFull *VatElectronicInfo `json:"VatElectronicInvoiceFull,omitnil" name:"VatElectronicInvoiceFull"`

	// 通用机打发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachinePrintedInvoice *MachinePrintedInvoice `json:"MachinePrintedInvoice,omitnil" name:"MachinePrintedInvoice"`

	// 汽车票
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusInvoice *BusInvoice `json:"BusInvoice,omitnil" name:"BusInvoice"`

	// 轮船票
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShippingInvoice *ShippingInvoice `json:"ShippingInvoice,omitnil" name:"ShippingInvoice"`

	// 过路过桥费发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	TollInvoice *TollInvoice `json:"TollInvoice,omitnil" name:"TollInvoice"`

	// 其他发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherInvoice *OtherInvoice `json:"OtherInvoice,omitnil" name:"OtherInvoice"`

	// 机动车销售统一发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	MotorVehicleSaleInvoice *MotorVehicleSaleInvoice `json:"MotorVehicleSaleInvoice,omitnil" name:"MotorVehicleSaleInvoice"`

	// 二手车销售统一发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedCarPurchaseInvoice *UsedCarPurchaseInvoice `json:"UsedCarPurchaseInvoice,omitnil" name:"UsedCarPurchaseInvoice"`

	// 增值税普通发票(卷票)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatInvoiceRoll *VatInvoiceRoll `json:"VatInvoiceRoll,omitnil" name:"VatInvoiceRoll"`

	// 出租车发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaxiTicket *TaxiTicket `json:"TaxiTicket,omitnil" name:"TaxiTicket"`

	// 定额发票
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaInvoice *QuotaInvoice `json:"QuotaInvoice,omitnil" name:"QuotaInvoice"`

	// 机票行程单
	// 注意：此字段可能返回 null，表示取不到有效值。
	AirTransport *AirTransport `json:"AirTransport,omitnil" name:"AirTransport"`

	// 非税收入通用票据
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonTaxIncomeGeneralBill *NonTaxIncomeBill `json:"NonTaxIncomeGeneralBill,omitnil" name:"NonTaxIncomeGeneralBill"`

	// 非税收入一般缴款书(电子)
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonTaxIncomeElectronicBill *NonTaxIncomeBill `json:"NonTaxIncomeElectronicBill,omitnil" name:"NonTaxIncomeElectronicBill"`

	// 火车票
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainTicket *TrainTicket `json:"TrainTicket,omitnil" name:"TrainTicket"`

	// 医疗门诊收费票据（电子）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalOutpatientInvoice *MedicalInvoice `json:"MedicalOutpatientInvoice,omitnil" name:"MedicalOutpatientInvoice"`

	// 医疗住院收费票据（电子）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalHospitalizedInvoice *MedicalInvoice `json:"MedicalHospitalizedInvoice,omitnil" name:"MedicalHospitalizedInvoice"`

	// 增值税销货清单
	// 注意：此字段可能返回 null，表示取不到有效值。
	VatSalesList *VatInvoiceInfo `json:"VatSalesList,omitnil" name:"VatSalesList"`

	// 电子发票（火车票）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElectronicTrainTicketFull *ElectronicTrainTicketFull `json:"ElectronicTrainTicketFull,omitnil" name:"ElectronicTrainTicketFull"`

	// 电子发票（机票行程单）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElectronicFlightTicketFull *ElectronicFlightTicketFull `json:"ElectronicFlightTicketFull,omitnil" name:"ElectronicFlightTicketFull"`
}

type SmartFormFileUrl struct {
	// 文件url地址
	FileUrl *string `json:"FileUrl,omitnil" name:"FileUrl"`

	// 文件的顺序，顺序从1开始
	FileOrderNumber *uint64 `json:"FileOrderNumber,omitnil" name:"FileOrderNumber"`
}

// Predefined struct for user
type SmartStructuralOCRRequestParams struct {
	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 自定义结构化功能需返回的字段名称，例：
	// 若客户只想返回姓名、性别两个字段的识别结果，则输入
	// ItemNames=["姓名","性别"]
	ItemNames []*string `json:"ItemNames,omitnil" name:"ItemNames"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否开启全文字段识别，默认值为false，开启后可返回全文字段识别结果。
	ReturnFullText *bool `json:"ReturnFullText,omitnil" name:"ReturnFullText"`
}

type SmartStructuralOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 自定义结构化功能需返回的字段名称，例：
	// 若客户只想返回姓名、性别两个字段的识别结果，则输入
	// ItemNames=["姓名","性别"]
	ItemNames []*string `json:"ItemNames,omitnil" name:"ItemNames"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 是否开启全文字段识别，默认值为false，开启后可返回全文字段识别结果。
	ReturnFullText *bool `json:"ReturnFullText,omitnil" name:"ReturnFullText"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "ReturnFullText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmartStructuralOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SmartStructuralOCRResponseParams struct {
	// 图片旋转角度(角度制)，文本的水平方向
	// 为 0；顺时针为正，逆时针为负
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 识别信息
	StructuralItems []*StructuralItem `json:"StructuralItems,omitnil" name:"StructuralItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SmartStructuralOCRResponse struct {
	*tchttp.BaseResponse
	Response *SmartStructuralOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type SmartStructuralOCRV2RequestParams struct {
	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 自定义结构化功能需返回的字段名称，例：
	// 若客户只想返回姓名、性别两个字段的识别结果，则输入
	// ItemNames=["姓名","性别"]
	ItemNames []*string `json:"ItemNames,omitnil" name:"ItemNames"`

	// 是否开启全文字段识别
	ReturnFullText *bool `json:"ReturnFullText,omitnil" name:"ReturnFullText"`

	// 配置id支持：
	// General -- 通用场景
	// OnlineTaxiItinerary -- 网约车行程单
	// RideHailingDriverLicense -- 网约车驾驶证
	// RideHailingTransportLicense -- 网约车运输证
	// WayBill -- 快递运单
	// AccountOpeningPermit -- 银行开户许可证
	// InvoiceEng -- 海外发票模版
	// Coin --钱币识别模板
	// OnboardingDocuments -- 入职材料识别
	// PropertyOwnershipCertificate -- 房产证识别
	// RealEstateCertificate --不动产权证识别
	// HouseEncumbranceCertificate -- 他权证识别
	ConfigId *string `json:"ConfigId,omitnil" name:"ConfigId"`

	// 是否打开印章识别
	EnableSealRecognize *bool `json:"EnableSealRecognize,omitnil" name:"EnableSealRecognize"`
}

type SmartStructuralOCRV2Request struct {
	*tchttp.BaseRequest
	
	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`

	// 自定义结构化功能需返回的字段名称，例：
	// 若客户只想返回姓名、性别两个字段的识别结果，则输入
	// ItemNames=["姓名","性别"]
	ItemNames []*string `json:"ItemNames,omitnil" name:"ItemNames"`

	// 是否开启全文字段识别
	ReturnFullText *bool `json:"ReturnFullText,omitnil" name:"ReturnFullText"`

	// 配置id支持：
	// General -- 通用场景
	// OnlineTaxiItinerary -- 网约车行程单
	// RideHailingDriverLicense -- 网约车驾驶证
	// RideHailingTransportLicense -- 网约车运输证
	// WayBill -- 快递运单
	// AccountOpeningPermit -- 银行开户许可证
	// InvoiceEng -- 海外发票模版
	// Coin --钱币识别模板
	// OnboardingDocuments -- 入职材料识别
	// PropertyOwnershipCertificate -- 房产证识别
	// RealEstateCertificate --不动产权证识别
	// HouseEncumbranceCertificate -- 他权证识别
	ConfigId *string `json:"ConfigId,omitnil" name:"ConfigId"`

	// 是否打开印章识别
	EnableSealRecognize *bool `json:"EnableSealRecognize,omitnil" name:"EnableSealRecognize"`
}

func (r *SmartStructuralOCRV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralOCRV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	delete(f, "ItemNames")
	delete(f, "ReturnFullText")
	delete(f, "ConfigId")
	delete(f, "EnableSealRecognize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SmartStructuralOCRV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SmartStructuralOCRV2ResponseParams struct {
	// 图片旋转角度(角度制)，文本的水平方向
	// 为 0；顺时针为正，逆时针为负
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 配置结构化文本信息
	StructuralList []*GroupInfo `json:"StructuralList,omitnil" name:"StructuralList"`

	// 还原文本信息
	WordList []*WordItem `json:"WordList,omitnil" name:"WordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SmartStructuralOCRV2Response struct {
	*tchttp.BaseResponse
	Response *SmartStructuralOCRV2ResponseParams `json:"Response"`
}

func (r *SmartStructuralOCRV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmartStructuralOCRV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StructuralItem struct {
	// 识别出的字段名称(关键字)。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 置信度 0 ~100。
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行在旋转纠正之后的图像中的像素
	// 坐标。
	ItemCoord *ItemCoord `json:"ItemCoord,omitnil" name:"ItemCoord"`

	// 字段所在行号，下标从0开始，非行字段或未能识别行号的该值返回-1。
	Row *int64 `json:"Row,omitnil" name:"Row"`
}

type TableCell struct {
	// 单元格左上角的列索引
	ColTl *int64 `json:"ColTl,omitnil" name:"ColTl"`

	// 单元格左上角的行索引
	RowTl *int64 `json:"RowTl,omitnil" name:"RowTl"`

	// 单元格右下角的列索引
	ColBr *int64 `json:"ColBr,omitnil" name:"ColBr"`

	// 单元格右下角的行索引
	RowBr *int64 `json:"RowBr,omitnil" name:"RowBr"`

	// 单元格内识别出的字符串文本，若文本存在多行，以换行符"\n"隔开
	Text *string `json:"Text,omitnil" name:"Text"`

	// 单元格类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 单元格置信度
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 单元格在图像中的四点坐标
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// 此字段为扩展字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 单元格文本属性
	Contents []*CellContent `json:"Contents,omitnil" name:"Contents"`
}

type TableCellInfo struct {
	// 单元格左上角的列索引
	ColTl *int64 `json:"ColTl,omitnil" name:"ColTl"`

	// 单元格左上角的行索引
	RowTl *int64 `json:"RowTl,omitnil" name:"RowTl"`

	// 单元格右下角的列索引
	ColBr *int64 `json:"ColBr,omitnil" name:"ColBr"`

	// 单元格右下角的行索引
	RowBr *int64 `json:"RowBr,omitnil" name:"RowBr"`

	// 单元格内识别出的字符串文本，若文本存在多行，以换行符"\n"隔开
	Text *string `json:"Text,omitnil" name:"Text"`

	// 单元格类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 单元格置信度
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 单元格在图像中的四点坐标
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`
}

type TableDetectInfo struct {
	// 单元格内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cells []*TableCell `json:"Cells,omitnil" name:"Cells"`

	// 表格标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Titles []*TableTitle `json:"Titles,omitnil" name:"Titles"`

	// 图像中的文本块类型，0 为非表格文本，
	// 1 为有线表格，2 为无线表格
	// （接口暂不支持日文无线表格识别，若传入日文无线表格，返回0）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 表格主体四个顶点坐标（依次为左上角，
	// 右上角，右下角，左下角）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableCoordPoint []*Coord `json:"TableCoordPoint,omitnil" name:"TableCoordPoint"`
}

type TableInfo struct {
	// 单元格内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cells []*TableCellInfo `json:"Cells,omitnil" name:"Cells"`

	// 图像中的文本块类型，0 为非表格文本，
	// 1 为有线表格，2 为无线表格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 表格主体四个顶点坐标（依次为左上角，
	// 右上角，右下角，左下角）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableCoordPoint []*Coord `json:"TableCoordPoint,omitnil" name:"TableCoordPoint"`
}

// Predefined struct for user
type TableOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type TableOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type TableOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接
	TextDetections []*TextTable `json:"TextDetections,omitnil" name:"TextDetections"`

	// Base64 编码后的 Excel 数据。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TableOCRResponse struct {
	*tchttp.BaseResponse
	Response *TableOCRResponseParams `json:"Response"`
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
	Text *string `json:"Text,omitnil" name:"Text"`
}

// Predefined struct for user
type TaxiInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type TaxiInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TaxiInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TaxiInvoiceOCRResponseParams struct {
	// 发票代码
	InvoiceNum *string `json:"InvoiceNum,omitnil" name:"InvoiceNum"`

	// 发票号码
	InvoiceCode *string `json:"InvoiceCode,omitnil" name:"InvoiceCode"`

	// 日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 金额
	Fare *string `json:"Fare,omitnil" name:"Fare"`

	// 上车时间
	GetOnTime *string `json:"GetOnTime,omitnil" name:"GetOnTime"`

	// 下车时间
	GetOffTime *string `json:"GetOffTime,omitnil" name:"GetOffTime"`

	// 里程
	Distance *string `json:"Distance,omitnil" name:"Distance"`

	// 发票所在地
	Location *string `json:"Location,omitnil" name:"Location"`

	// 车牌号
	PlateNumber *string `json:"PlateNumber,omitnil" name:"PlateNumber"`

	// 发票消费类型
	InvoiceType *string `json:"InvoiceType,omitnil" name:"InvoiceType"`

	// 省
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil" name:"City"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TaxiInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *TaxiInvoiceOCRResponseParams `json:"Response"`
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

type TaxiTicket struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 上车时间
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// 下车时间
	TimeGetOff *string `json:"TimeGetOff,omitnil" name:"TimeGetOff"`

	// 单价
	Price *string `json:"Price,omitnil" name:"Price"`

	// 里程
	Mileage *string `json:"Mileage,omitnil" name:"Mileage"`

	// 总金额
	Total *string `json:"Total,omitnil" name:"Total"`

	// 发票所在地
	Place *string `json:"Place,omitnil" name:"Place"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 车牌号
	LicensePlate *string `json:"LicensePlate,omitnil" name:"LicensePlate"`

	// 燃油附加费
	FuelFee *string `json:"FuelFee,omitnil" name:"FuelFee"`

	// 预约叫车服务费
	BookingCallFee *string `json:"BookingCallFee,omitnil" name:"BookingCallFee"`

	// 是否有公司印章（0：没有，1：有）
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`
}

type TextArithmetic struct {
	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// 算式运算结果，true-正确   false-错误或非法参数
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 保留字段，暂不支持
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 原图文本行坐标，以四个顶点坐标表示（保留字段，暂不支持）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// 保留字段，暂不支持
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 文本行旋转纠正之后在图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	ItemCoord *ItemCoord `json:"ItemCoord,omitnil" name:"ItemCoord"`

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
	ExpressionType *string `json:"ExpressionType,omitnil" name:"ExpressionType"`

	// 错题推荐答案，算式运算结果正确返回为""，算式运算结果错误返回推荐答案 (注：暂不支持多个关系运算符（如1<10<7）、无关系运算符（如frac(1,2)+frac(2,3)）、单位换算（如1元=100角）错题的推荐答案返回)
	Answer *string `json:"Answer,omitnil" name:"Answer"`
}

// Predefined struct for user
type TextDetectRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type TextDetectRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type TextDetectResponseParams struct {
	// 图片中是否包含文字。
	HasText *bool `json:"HasText,omitnil" name:"HasText"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextDetectResponse struct {
	*tchttp.BaseResponse
	Response *TextDetectResponseParams `json:"Response"`
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
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// 此字段为扩展字段。
	// GeneralBasicOcr接口返回段落信息Parag，包含ParagNo。
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 文本行在旋转纠正之后的图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	ItemPolygon *ItemCoord `json:"ItemPolygon,omitnil" name:"ItemPolygon"`

	// 识别出来的单字信息包括单字（包括单字Character和单字置信度confidence）， 支持识别的接口：GeneralBasicOCR、GeneralAccurateOCR
	Words []*DetectedWords `json:"Words,omitnil" name:"Words"`

	// 单字在原图中的四点坐标， 支持识别的接口：GeneralBasicOCR、GeneralAccurateOCR
	WordCoordPoint []*DetectedWordCoordPoint `json:"WordCoordPoint,omitnil" name:"WordCoordPoint"`
}

type TextDetectionEn struct {
	// 识别出的文本行内容。
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// 置信度 0 ~100。
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行在原图中的四点坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// 此字段为扩展字段。目前EnglishOCR接口返回内容为空。
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 英文单词在原图中的四点坐标。
	WordCoordPoint []*WordCoordPoint `json:"WordCoordPoint,omitnil" name:"WordCoordPoint"`

	// 候选字符集(包含候选字Character以及置信度Confidence)。
	CandWord []*CandWord `json:"CandWord,omitnil" name:"CandWord"`

	// 识别出来的单词信息（包括单词Character和单词置信度confidence）
	Words []*Words `json:"Words,omitnil" name:"Words"`
}

type TextDetectionResult struct {
	// 识别出的文本行内容
	Value *string `json:"Value,omitnil" name:"Value"`

	// 坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`
}

type TextEduPaper struct {
	// 识别出的字段名称（关键字）
	Item *string `json:"Item,omitnil" name:"Item"`

	// 识别出的字段名称对应的值，也就是字段Item对应的字符串结果
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// 文本行在旋转纠正之后的图像中的像素坐标，表示为（左上角x, 左上角y，宽width，高height）
	Itemcoord *ItemCoord `json:"Itemcoord,omitnil" name:"Itemcoord"`
}

type TextFormula struct {
	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`
}

type TextGeneralHandwriting struct {
	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// 置信度 0 - 100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// 此字段为扩展字段。
	// 能返回文本行的段落信息，例如：{\"Parag\":{\"ParagNo\":2}}，
	// 其中ParagNo为段落行，从1开始。
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`

	// 字的坐标数组，以四个顶点坐标表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordPolygon []*Polygon `json:"WordPolygon,omitnil" name:"WordPolygon"`
}

type TextTable struct {
	// 单元格左上角的列索引
	ColTl *int64 `json:"ColTl,omitnil" name:"ColTl"`

	// 单元格左上角的行索引
	RowTl *int64 `json:"RowTl,omitnil" name:"RowTl"`

	// 单元格右下角的列索引
	ColBr *int64 `json:"ColBr,omitnil" name:"ColBr"`

	// 单元格右下角的行索引
	RowBr *int64 `json:"RowBr,omitnil" name:"RowBr"`

	// 单元格文字
	Text *string `json:"Text,omitnil" name:"Text"`

	// 单元格类型，包含body（表格主体）、header（表头）、footer（表尾）三种
	Type *string `json:"Type,omitnil" name:"Type"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitnil" name:"Polygon"`

	// 此字段为扩展字段
	AdvancedInfo *string `json:"AdvancedInfo,omitnil" name:"AdvancedInfo"`
}

type TextVatInvoice struct {
	// 识别出的字段名称（关键字）。支持以下字段的识别：
	// 发票代码、 发票号码、 打印发票代码、 打印发票号码、 开票日期、 购买方识别号、 小写金额、 价税合计(大写)、 销售方识别号、 校验码、 购买方名称、 销售方名称、 税额、 复核、 联次名称、 备注、 联次、 密码区、 开票人、 收款人、 （货物或应税劳务、服务名称）、省、 市、 服务类型、 通行费标志、 是否代开、 是否收购、 合计金额、 是否有公司印章、 发票消费类型、 车船税、 机器编号、 成品油标志、 税率、 合计税额、 （购买方地址、电话）、 （销售方地址、电话）、 单价、 金额、 销售方开户行及账号、 购买方开户行及账号、 规格型号、 发票名称、 单位、 数量、 校验码备选、 校验码后六位备选、发票号码备选、车牌号、类型、通行日期起、通行日期止、发票类型。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 字段在原图中的中的四点坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon *Polygon `json:"Polygon,omitnil" name:"Polygon"`
}

type TextVehicleBack struct {
	// 号牌号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateNo *string `json:"PlateNo,omitnil" name:"PlateNo"`

	// 档案编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNo *string `json:"FileNo,omitnil" name:"FileNo"`

	// 核定人数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowNum *string `json:"AllowNum,omitnil" name:"AllowNum"`

	// 总质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalMass *string `json:"TotalMass,omitnil" name:"TotalMass"`

	// 整备质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurbWeight *string `json:"CurbWeight,omitnil" name:"CurbWeight"`

	// 核定载质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadQuality *string `json:"LoadQuality,omitnil" name:"LoadQuality"`

	// 外廓尺寸
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalSize *string `json:"ExternalSize,omitnil" name:"ExternalSize"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Marks *string `json:"Marks,omitnil" name:"Marks"`

	// 检验记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Record *string `json:"Record,omitnil" name:"Record"`

	// 准牵引总质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalQuasiMass *string `json:"TotalQuasiMass,omitnil" name:"TotalQuasiMass"`

	// 副页编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubPageCode *string `json:"SubPageCode,omitnil" name:"SubPageCode"`

	// 燃料种类
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	FuelType *string `json:"FuelType,omitnil" name:"FuelType"`
}

type TextVehicleFront struct {
	// 号牌号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlateNo *string `json:"PlateNo,omitnil" name:"PlateNo"`

	// 车辆类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// 所有人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 住址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil" name:"Address"`

	// 使用性质
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseCharacter *string `json:"UseCharacter,omitnil" name:"UseCharacter"`

	// 品牌型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitnil" name:"Model"`

	// 车辆识别代号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vin *string `json:"Vin,omitnil" name:"Vin"`

	// 发动机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNo *string `json:"EngineNo,omitnil" name:"EngineNo"`

	// 注册日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterDate *string `json:"RegisterDate,omitnil" name:"RegisterDate"`

	// 发证日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueDate *string `json:"IssueDate,omitnil" name:"IssueDate"`

	// 印章
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seal *string `json:"Seal,omitnil" name:"Seal"`
}

type TextWaybill struct {
	// 收件人姓名
	RecName *WaybillObj `json:"RecName,omitnil" name:"RecName"`

	// 收件人手机号
	RecNum *WaybillObj `json:"RecNum,omitnil" name:"RecNum"`

	// 收件人地址
	RecAddr *WaybillObj `json:"RecAddr,omitnil" name:"RecAddr"`

	// 寄件人姓名
	SenderName *WaybillObj `json:"SenderName,omitnil" name:"SenderName"`

	// 寄件人手机号
	SenderNum *WaybillObj `json:"SenderNum,omitnil" name:"SenderNum"`

	// 寄件人地址
	SenderAddr *WaybillObj `json:"SenderAddr,omitnil" name:"SenderAddr"`

	// 运单号
	WaybillNum *WaybillObj `json:"WaybillNum,omitnil" name:"WaybillNum"`
}

type TollInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 入口
	Entrance *string `json:"Entrance,omitnil" name:"Entrance"`

	// 出口
	Exit *string `json:"Exit,omitnil" name:"Exit"`

	// 高速标志（0：没有，1：有）
	HighwayMark *int64 `json:"HighwayMark,omitnil" name:"HighwayMark"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`
}

type TollInvoiceInfo struct {
	// 识别出的字段名称（关键字）。支持以下字段的识别：
	// 发票代码、发票号码、日期、金额、入口、出口、时间、发票消费类型、高速标志。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type TollInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type TollInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TollInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TollInvoiceOCRResponseParams struct {
	// 过路过桥费发票识别结果，具体内容请点击左侧链接。
	TollInvoiceInfos []*TollInvoiceInfo `json:"TollInvoiceInfos,omitnil" name:"TollInvoiceInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TollInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *TollInvoiceOCRResponseParams `json:"Response"`
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

type TrainTicket struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 乘车日期
	DateGetOn *string `json:"DateGetOn,omitnil" name:"DateGetOn"`

	// 乘车时间
	TimeGetOn *string `json:"TimeGetOn,omitnil" name:"TimeGetOn"`

	// 乘车人姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 出发车站
	StationGetOn *string `json:"StationGetOn,omitnil" name:"StationGetOn"`

	// 到达车站
	StationGetOff *string `json:"StationGetOff,omitnil" name:"StationGetOff"`

	// 座位类型
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// 总金额
	Total *string `json:"Total,omitnil" name:"Total"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 序列号
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`

	// 身份证号
	UserID *string `json:"UserID,omitnil" name:"UserID"`

	// 检票口
	GateNumber *string `json:"GateNumber,omitnil" name:"GateNumber"`

	// 车次
	TrainNumber *string `json:"TrainNumber,omitnil" name:"TrainNumber"`

	// 手续费
	HandlingFee *string `json:"HandlingFee,omitnil" name:"HandlingFee"`

	// 原票价
	OriginalFare *string `json:"OriginalFare,omitnil" name:"OriginalFare"`

	// 大写金额
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 座位号
	SeatNumber *string `json:"SeatNumber,omitnil" name:"SeatNumber"`

	// 取票地址
	PickUpAddress *string `json:"PickUpAddress,omitnil" name:"PickUpAddress"`

	// 是否始发改签
	TicketChange *string `json:"TicketChange,omitnil" name:"TicketChange"`

	// 加收票价
	AdditionalFare *string `json:"AdditionalFare,omitnil" name:"AdditionalFare"`

	// 收据号码
	ReceiptNumber *string `json:"ReceiptNumber,omitnil" name:"ReceiptNumber"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 是否仅供报销使用（0：没有，1：有）
	ReimburseOnlyMark *int64 `json:"ReimburseOnlyMark,omitnil" name:"ReimburseOnlyMark"`

	// 是否有退票费标识（0：没有，1：有）
	RefundMark *int64 `json:"RefundMark,omitnil" name:"RefundMark"`
}

// Predefined struct for user
type TrainTicketOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type TrainTicketOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TrainTicketOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TrainTicketOCRResponseParams struct {
	// 编号
	TicketNum *string `json:"TicketNum,omitnil" name:"TicketNum"`

	// 出发站
	StartStation *string `json:"StartStation,omitnil" name:"StartStation"`

	// 到达站
	DestinationStation *string `json:"DestinationStation,omitnil" name:"DestinationStation"`

	// 出发时间
	Date *string `json:"Date,omitnil" name:"Date"`

	// 车次
	TrainNum *string `json:"TrainNum,omitnil" name:"TrainNum"`

	// 座位号
	Seat *string `json:"Seat,omitnil" name:"Seat"`

	// 姓名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 票价
	Price *string `json:"Price,omitnil" name:"Price"`

	// 席别
	SeatCategory *string `json:"SeatCategory,omitnil" name:"SeatCategory"`

	// 身份证号
	ID *string `json:"ID,omitnil" name:"ID"`

	// 发票消费类型：交通
	InvoiceType *string `json:"InvoiceType,omitnil" name:"InvoiceType"`

	// 序列号
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`

	// 加收票价
	AdditionalCost *string `json:"AdditionalCost,omitnil" name:"AdditionalCost"`

	// 手续费
	HandlingFee *string `json:"HandlingFee,omitnil" name:"HandlingFee"`

	// 大写金额（票面有大写金额该字段才有值）
	LegalAmount *string `json:"LegalAmount,omitnil" name:"LegalAmount"`

	// 售票站
	TicketStation *string `json:"TicketStation,omitnil" name:"TicketStation"`

	// 原票价（一般有手续费的才有原始票价字段）
	OriginalPrice *string `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// 发票类型：火车票、火车票补票、火车票退票凭证
	InvoiceStyle *string `json:"InvoiceStyle,omitnil" name:"InvoiceStyle"`

	// 收据号码
	ReceiptNumber *string `json:"ReceiptNumber,omitnil" name:"ReceiptNumber"`

	// 仅供报销使用：1为是，0为否
	IsReceipt *string `json:"IsReceipt,omitnil" name:"IsReceipt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TrainTicketOCRResponse struct {
	*tchttp.BaseResponse
	Response *TrainTicketOCRResponseParams `json:"Response"`
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

type UsedCarPurchaseInvoice struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 是否存在二维码（0：没有，1：有）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 销货单位名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 销售方电话
	SellerTel *string `json:"SellerTel,omitnil" name:"SellerTel"`

	// 销售方单位代码/个人身份证号
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// 销售方地址
	SellerAddress *string `json:"SellerAddress,omitnil" name:"SellerAddress"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 购买方单位代码/个人身份证号
	BuyerID *string `json:"BuyerID,omitnil" name:"BuyerID"`

	// 购买方地址
	BuyerAddress *string `json:"BuyerAddress,omitnil" name:"BuyerAddress"`

	// 购买方电话
	BuyerTel *string `json:"BuyerTel,omitnil" name:"BuyerTel"`

	// 二手车市场
	CompanyName *string `json:"CompanyName,omitnil" name:"CompanyName"`

	// 二手车市场纳税人识别号
	CompanyTaxID *string `json:"CompanyTaxID,omitnil" name:"CompanyTaxID"`

	// 二手车市场开户银行和账号
	CompanyBankAccount *string `json:"CompanyBankAccount,omitnil" name:"CompanyBankAccount"`

	// 二手车市场电话
	CompanyTel *string `json:"CompanyTel,omitnil" name:"CompanyTel"`

	// 二手车市场地址
	CompanyAddress *string `json:"CompanyAddress,omitnil" name:"CompanyAddress"`

	// 转入地车辆管理所名称
	TransferAdministrationName *string `json:"TransferAdministrationName,omitnil" name:"TransferAdministrationName"`

	// 车牌号
	LicensePlate *string `json:"LicensePlate,omitnil" name:"LicensePlate"`

	// 登记证号
	RegistrationNumber *string `json:"RegistrationNumber,omitnil" name:"RegistrationNumber"`

	// 车辆识别代码
	VIN *string `json:"VIN,omitnil" name:"VIN"`

	// 厂牌型号
	VehicleModel *string `json:"VehicleModel,omitnil" name:"VehicleModel"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 车辆类型
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 发票联次
	FormType *string `json:"FormType,omitnil" name:"FormType"`

	// 发票联名
	FormName *string `json:"FormName,omitnil" name:"FormName"`

	// 是否有公司印章（0：没有，1：有）
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// 经营拍卖单位
	AuctionOrgName *string `json:"AuctionOrgName,omitnil" name:"AuctionOrgName"`

	// 经营拍卖单位地址
	AuctionOrgAddress *string `json:"AuctionOrgAddress,omitnil" name:"AuctionOrgAddress"`

	// 经营拍卖单位纳税人识别号
	AuctionOrgTaxID *string `json:"AuctionOrgTaxID,omitnil" name:"AuctionOrgTaxID"`

	// 经营拍卖单位开户银行账号
	AuctionOrgBankAccount *string `json:"AuctionOrgBankAccount,omitnil" name:"AuctionOrgBankAccount"`

	// 经营拍卖单位电话
	AuctionOrgPhone *string `json:"AuctionOrgPhone,omitnil" name:"AuctionOrgPhone"`

	// 开票人
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 税控码
	TaxCode *string `json:"TaxCode,omitnil" name:"TaxCode"`

	// 机器编号
	MachineSerialNumber *string `json:"MachineSerialNumber,omitnil" name:"MachineSerialNumber"`

	// 机打发票代码
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 机打发票号码
	MachineNumber *string `json:"MachineNumber,omitnil" name:"MachineNumber"`
}

type UsedVehicleInvoiceInfo struct {
	// 所属税局
	TaxBureau *string `json:"TaxBureau,omitnil" name:"TaxBureau"`

	// 买方单位/个人
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 买方单位代码/身份证号码
	BuyerNo *string `json:"BuyerNo,omitnil" name:"BuyerNo"`

	// 买方单位/个人地址
	BuyerAddress *string `json:"BuyerAddress,omitnil" name:"BuyerAddress"`

	// 买方单位电话
	BuyerTel *string `json:"BuyerTel,omitnil" name:"BuyerTel"`

	// 卖方单位/个人
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 卖方单位代码/身份证号码
	SellerNo *string `json:"SellerNo,omitnil" name:"SellerNo"`

	// 卖方单位/个人地址
	SellerAddress *string `json:"SellerAddress,omitnil" name:"SellerAddress"`

	// 卖方单位电话
	SellerTel *string `json:"SellerTel,omitnil" name:"SellerTel"`

	// 车牌照号
	VehicleLicenseNo *string `json:"VehicleLicenseNo,omitnil" name:"VehicleLicenseNo"`

	// 登记证号
	RegisterNo *string `json:"RegisterNo,omitnil" name:"RegisterNo"`

	// 车架号/车辆识别代码
	VehicleIdentifyNo *string `json:"VehicleIdentifyNo,omitnil" name:"VehicleIdentifyNo"`

	// 转入地车辆管理所名称
	ManagementOffice *string `json:"ManagementOffice,omitnil" name:"ManagementOffice"`

	// 车价合计
	VehicleTotalPrice *string `json:"VehicleTotalPrice,omitnil" name:"VehicleTotalPrice"`

	// 经营、拍卖单位
	Auctioneer *string `json:"Auctioneer,omitnil" name:"Auctioneer"`

	// 经营、拍卖单位地址
	AuctioneerAddress *string `json:"AuctioneerAddress,omitnil" name:"AuctioneerAddress"`

	// 经营、拍卖单位纳税人识别号
	AuctioneerTaxpayerNum *string `json:"AuctioneerTaxpayerNum,omitnil" name:"AuctioneerTaxpayerNum"`

	// 经营、拍卖单位开户银行、账号
	AuctioneerBankAccount *string `json:"AuctioneerBankAccount,omitnil" name:"AuctioneerBankAccount"`

	// 经营、拍卖单位电话
	AuctioneerTel *string `json:"AuctioneerTel,omitnil" name:"AuctioneerTel"`

	// 二手车市场
	Market *string `json:"Market,omitnil" name:"Market"`

	// 二手车市场纳税人识别号
	MarketTaxpayerNum *string `json:"MarketTaxpayerNum,omitnil" name:"MarketTaxpayerNum"`

	// 二手车市场地址
	MarketAddress *string `json:"MarketAddress,omitnil" name:"MarketAddress"`

	// 二手车市场开户银行账号
	MarketBankAccount *string `json:"MarketBankAccount,omitnil" name:"MarketBankAccount"`

	// 二手车市场电话
	MarketTel *string `json:"MarketTel,omitnil" name:"MarketTel"`
}

type Value struct {
	// 自动识别的字段内容
	AutoContent *string `json:"AutoContent,omitnil" name:"AutoContent"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coord *Polygon `json:"Coord,omitnil" name:"Coord"`
}

type VatElectronicInfo struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 税前金额
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// 合计税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 销售方名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 销售方纳税人识别号
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 购买方纳税人识别号
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 开票人
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 小计金额
	SubTotal *string `json:"SubTotal,omitnil" name:"SubTotal"`

	// 小计税额
	SubTax *string `json:"SubTax,omitnil" name:"SubTax"`

	// 电子发票详细条目信息
	VatElectronicItems []*VatElectronicItemInfo `json:"VatElectronicItems,omitnil" name:"VatElectronicItems"`
}

type VatElectronicItemInfo struct {
	// 项目名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数量
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// 规格型号
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 单价
	Price *string `json:"Price,omitnil" name:"Price"`

	// 金额
	Total *string `json:"Total,omitnil" name:"Total"`

	// 税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 单位
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 运输工具类型
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// 运输工具牌号
	VehicleBrand *string `json:"VehicleBrand,omitnil" name:"VehicleBrand"`

	// 起始地
	DeparturePlace *string `json:"DeparturePlace,omitnil" name:"DeparturePlace"`

	// 到达地
	ArrivalPlace *string `json:"ArrivalPlace,omitnil" name:"ArrivalPlace"`

	// 运输货物名称，仅货物运输服务发票返回
	TransportItemsName *string `json:"TransportItemsName,omitnil" name:"TransportItemsName"`

	// 建筑服务发生地，仅建筑发票返回
	PlaceOfBuildingService *string `json:"PlaceOfBuildingService,omitnil" name:"PlaceOfBuildingService"`

	// 建筑项目名称，仅建筑发票返回
	BuildingName *string `json:"BuildingName,omitnil" name:"BuildingName"`

	// 产权证书/不动产权证号，仅不动产经营租赁服务发票返回
	EstateNumber *string `json:"EstateNumber,omitnil" name:"EstateNumber"`

	// 面积单位，仅不动产经营租赁服务发票返回
	AreaUnit *string `json:"AreaUnit,omitnil" name:"AreaUnit"`
}

type VatInvoice struct {
	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 购方抬头
	// 通用机打发票类型时不返回
	BuyerName *string `json:"BuyerName,omitnil" name:"BuyerName"`

	// 购方税号
	// 通用机打发票类型时不返回
	BuyerTaxCode *string `json:"BuyerTaxCode,omitnil" name:"BuyerTaxCode"`

	// 购方地址电话
	// 通用机打发票类型做不返回
	BuyerAddressPhone *string `json:"BuyerAddressPhone,omitnil" name:"BuyerAddressPhone"`

	// 购方银行账号
	// 通用机打发票类型时不返回
	BuyerBankAccount *string `json:"BuyerBankAccount,omitnil" name:"BuyerBankAccount"`

	// 销方名称
	SellerName *string `json:"SellerName,omitnil" name:"SellerName"`

	// 销方税号
	SellerTaxCode *string `json:"SellerTaxCode,omitnil" name:"SellerTaxCode"`

	// 销方地址电话
	SellerAddressPhone *string `json:"SellerAddressPhone,omitnil" name:"SellerAddressPhone"`

	// 销方银行账号
	SellerBankAccount *string `json:"SellerBankAccount,omitnil" name:"SellerBankAccount"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 机器编码
	MachineNo *string `json:"MachineNo,omitnil" name:"MachineNo"`

	// 票种类型
	// 01：增值税专用发票，
	// 02：货运运输业增值税专用发票，
	// 03：机动车销售统一发票，
	// 04：增值税普通发票，
	// 08：增值税电子专用发票（含全电，全电仅新版接口支持），
	// 10：增值税电子普通发票（含全电，全电仅新版接口支持），
	// 11：增值税普通发票（卷式），
	// 14：增值税电子（通行费）发票，
	// 15：二手车销售统一发票，
	// 32：深圳区块链发票，
	// 102：通用机打电子发票
	// 61：电子发票（航空运输电子客票行程单）
	// 83：电子发票（铁路电子发票）
	// 0915：全电纸质（二手车统一销售发票）
	// 0903：全电纸质（机动车统一发票）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 具体的全电发票类型：01: 全电专用发票；02：全电普通发票；03：全电火车票；04：全电机票行程单
	ElectronicType *string `json:"ElectronicType,omitnil" name:"ElectronicType"`

	// 检验码
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 是否作废（红冲）是否作废（红冲）
	// Y：已作废，N：未作废，H：红冲，HP：部分红冲，HF：全额红冲
	IsAbandoned *string `json:"IsAbandoned,omitnil" name:"IsAbandoned"`

	// 是否有销货清单 
	// Y: 有清单 N：无清单 
	// 卷票无
	HasSellerList *string `json:"HasSellerList,omitnil" name:"HasSellerList"`

	// 销货清单标题
	SellerListTitle *string `json:"SellerListTitle,omitnil" name:"SellerListTitle"`

	// 销货清单税额
	SellerListTax *string `json:"SellerListTax,omitnil" name:"SellerListTax"`

	// 不含税金额
	AmountWithoutTax *string `json:"AmountWithoutTax,omitnil" name:"AmountWithoutTax"`

	// 税额
	TaxAmount *string `json:"TaxAmount,omitnil" name:"TaxAmount"`

	// 含税金额
	AmountWithTax *string `json:"AmountWithTax,omitnil" name:"AmountWithTax"`

	// 项目明细
	Items []*VatInvoiceItem `json:"Items,omitnil" name:"Items"`

	// 所属税局
	TaxBureau *string `json:"TaxBureau,omitnil" name:"TaxBureau"`

	// 通行费标志:Y、是;N、否
	TrafficFreeFlag *string `json:"TrafficFreeFlag,omitnil" name:"TrafficFreeFlag"`

	// 是否为红票
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedLetterInvoiceMark *bool `json:"RedLetterInvoiceMark,omitnil" name:"RedLetterInvoiceMark"`

	// 开具类型标识（0: 委托代开，1：自开，2：代开，3：代办退税
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssuingTypeMark *int64 `json:"IssuingTypeMark,omitnil" name:"IssuingTypeMark"`

	// 代开销售方名称
	SellerAgentName *string `json:"SellerAgentName,omitnil" name:"SellerAgentName"`

	// 代开销售方税号
	SellerAgentTaxID *string `json:"SellerAgentTaxID,omitnil" name:"SellerAgentTaxID"`
}

type VatInvoiceGoodsInfo struct {
	// 项目名称
	Item *string `json:"Item,omitnil" name:"Item"`

	// 规格型号
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 单位
	MeasurementDimension *string `json:"MeasurementDimension,omitnil" name:"MeasurementDimension"`

	// 价格
	Price *string `json:"Price,omitnil" name:"Price"`

	// 数量
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// 金额
	Amount *string `json:"Amount,omitnil" name:"Amount"`

	// 税率(如6%、免税)
	TaxScheme *string `json:"TaxScheme,omitnil" name:"TaxScheme"`

	// 税额
	TaxAmount *string `json:"TaxAmount,omitnil" name:"TaxAmount"`
}

type VatInvoiceInfo struct {
	// 校验码
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 发票联次
	FormType *string `json:"FormType,omitnil" name:"FormType"`

	// 车船税
	TravelTax *string `json:"TravelTax,omitnil" name:"TravelTax"`

	// 购买方地址电话
	BuyerAddrTel *string `json:"BuyerAddrTel,omitnil" name:"BuyerAddrTel"`

	// 购买方银行账号
	BuyerBankAccount *string `json:"BuyerBankAccount,omitnil" name:"BuyerBankAccount"`

	// 公司印章内容
	CompanySealContent *string `json:"CompanySealContent,omitnil" name:"CompanySealContent"`

	// 税务局章内容
	TaxSealContent *string `json:"TaxSealContent,omitnil" name:"TaxSealContent"`

	// 服务类型
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 是否存在二维码（0：没有，1：有）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 是否有代开标记（0：没有，1：有）
	AgentMark *int64 `json:"AgentMark,omitnil" name:"AgentMark"`

	// 是否有通行费标记（0：没有，1：有）
	TransitMark *int64 `json:"TransitMark,omitnil" name:"TransitMark"`

	// 是否有成品油标记（0：没有，1：有）
	OilMark *int64 `json:"OilMark,omitnil" name:"OilMark"`

	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 机打发票号码
	NumberConfirm *string `json:"NumberConfirm,omitnil" name:"NumberConfirm"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 税前金额
	PretaxAmount *string `json:"PretaxAmount,omitnil" name:"PretaxAmount"`

	// 合计税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 机器编号
	MachineCode *string `json:"MachineCode,omitnil" name:"MachineCode"`

	// 密码区
	Ciphertext *string `json:"Ciphertext,omitnil" name:"Ciphertext"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 销售方名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 销售方纳税人识别号
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// 销售方地址电话
	SellerAddrTel *string `json:"SellerAddrTel,omitnil" name:"SellerAddrTel"`

	// 销售方银行账号
	SellerBankAccount *string `json:"SellerBankAccount,omitnil" name:"SellerBankAccount"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 购买方纳税人识别号
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 是否有公司印章（0：没有，1：有）
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// 开票人
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 复核人
	Reviewer *string `json:"Reviewer,omitnil" name:"Reviewer"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 增值税发票项目信息
	VatInvoiceItemInfos []*VatInvoiceItemInfo `json:"VatInvoiceItemInfos,omitnil" name:"VatInvoiceItemInfos"`

	// 机打发票代码
	CodeConfirm *string `json:"CodeConfirm,omitnil" name:"CodeConfirm"`

	// 收款人
	Receiptor *string `json:"Receiptor,omitnil" name:"Receiptor"`

	// 是否有全电纸质票（0：没有，1：有）
	ElectronicFullMark *int64 `json:"ElectronicFullMark,omitnil" name:"ElectronicFullMark"`

	// 全电号码
	ElectronicFullNumber *string `json:"ElectronicFullNumber,omitnil" name:"ElectronicFullNumber"`

	// 发票联名
	FormName *string `json:"FormName,omitnil" name:"FormName"`

	// 是否有区块链标记（0：没有，1：有）	
	BlockChainMark *int64 `json:"BlockChainMark,omitnil" name:"BlockChainMark"`

	// 是否有收购标记（0：没有，1：有）	
	AcquisitionMark *int64 `json:"AcquisitionMark,omitnil" name:"AcquisitionMark"`

	// 小计金额
	SubTotal *string `json:"SubTotal,omitnil" name:"SubTotal"`

	// 小计税额
	SubTax *string `json:"SubTax,omitnil" name:"SubTax"`
}

type VatInvoiceItem struct {
	// 行号
	LineNo *string `json:"LineNo,omitnil" name:"LineNo"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 规格
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 单位
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 数量
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// 单价
	UnitPrice *string `json:"UnitPrice,omitnil" name:"UnitPrice"`

	// 不含税金额
	AmountWithoutTax *string `json:"AmountWithoutTax,omitnil" name:"AmountWithoutTax"`

	// 税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 税额
	TaxAmount *string `json:"TaxAmount,omitnil" name:"TaxAmount"`

	// 税收分类编码
	TaxClassifyCode *string `json:"TaxClassifyCode,omitnil" name:"TaxClassifyCode"`

	// 运输工具类型
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// 运输工具牌号
	VehicleBrand *string `json:"VehicleBrand,omitnil" name:"VehicleBrand"`

	// 起始地
	DeparturePlace *string `json:"DeparturePlace,omitnil" name:"DeparturePlace"`

	// 到达地
	ArrivalPlace *string `json:"ArrivalPlace,omitnil" name:"ArrivalPlace"`

	// 运输货物名称
	TransportItemsName *string `json:"TransportItemsName,omitnil" name:"TransportItemsName"`

	// 建筑服务发生地
	ConstructionPlace *string `json:"ConstructionPlace,omitnil" name:"ConstructionPlace"`

	// 建筑项目名称
	ConstructionName *string `json:"ConstructionName,omitnil" name:"ConstructionName"`
}

type VatInvoiceItemInfo struct {
	// 项目名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 规格型号
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 单位
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 数量
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// 单价
	Price *string `json:"Price,omitnil" name:"Price"`

	// 金额
	Total *string `json:"Total,omitnil" name:"Total"`

	// 税率
	TaxRate *string `json:"TaxRate,omitnil" name:"TaxRate"`

	// 税额
	Tax *string `json:"Tax,omitnil" name:"Tax"`

	// 通行日期起
	DateStart *string `json:"DateStart,omitnil" name:"DateStart"`

	// 通行日期止
	DateEnd *string `json:"DateEnd,omitnil" name:"DateEnd"`

	// 车牌号
	LicensePlate *string `json:"LicensePlate,omitnil" name:"LicensePlate"`

	// 车辆类型
	VehicleType *string `json:"VehicleType,omitnil" name:"VehicleType"`

	// 序号
	SerialNumber *string `json:"SerialNumber,omitnil" name:"SerialNumber"`
}

// Predefined struct for user
type VatInvoiceOCRRequestParams struct {
	// 图片/PDF的 Base64 值。
	// 支持的文件格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片/PDF大小：所下载文件经Base64编码后不超过 7M。文件下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 输入参数 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 支持的文件格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片/PDF大小：所下载文件经 Base64 编码后不超过 7M。文件下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type VatInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片/PDF的 Base64 值。
	// 支持的文件格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片/PDF大小：所下载文件经Base64编码后不超过 7M。文件下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 输入参数 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片/PDF的 Url 地址。
	// 支持的文件格式：PNG、JPG、JPEG、PDF，暂不支持 GIF 格式。
	// 支持的图片/PDF大小：所下载文件经 Base64 编码后不超过 7M。文件下载时间不超过 3 秒。
	// 支持的图片像素：需介于20-10000px之间。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为false，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *uint64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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

// Predefined struct for user
type VatInvoiceOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	VatInvoiceInfos []*TextVatInvoice `json:"VatInvoiceInfos,omitnil" name:"VatInvoiceInfos"`

	// 明细条目。VatInvoiceInfos中关于明细项的具体条目。
	Items []*VatInvoiceItem `json:"Items,omitnil" name:"Items"`

	// 默认值为0。如果图片为PDF时，返回PDF的总页数。
	PdfPageSize *int64 `json:"PdfPageSize,omitnil" name:"PdfPageSize"`

	// 图片旋转角度（角度制），文本的水平方向为0°；顺时针为正，逆时针为负。点击查看<a href="https://cloud.tencent.com/document/product/866/45139">如何纠正倾斜文本</a>
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VatInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *VatInvoiceOCRResponseParams `json:"Response"`
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

type VatInvoiceRoll struct {
	// 发票名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 发票代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 发票号码
	Number *string `json:"Number,omitnil" name:"Number"`

	// 机打发票号码
	NumberConfirm *string `json:"NumberConfirm,omitnil" name:"NumberConfirm"`

	// 开票日期
	Date *string `json:"Date,omitnil" name:"Date"`

	// 校验码
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 销售方名称
	Seller *string `json:"Seller,omitnil" name:"Seller"`

	// 销售方纳税人识别号
	SellerTaxID *string `json:"SellerTaxID,omitnil" name:"SellerTaxID"`

	// 购买方名称
	Buyer *string `json:"Buyer,omitnil" name:"Buyer"`

	// 购买方纳税人识别号
	BuyerTaxID *string `json:"BuyerTaxID,omitnil" name:"BuyerTaxID"`

	// 种类
	Category *string `json:"Category,omitnil" name:"Category"`

	// 价税合计（小写）
	Total *string `json:"Total,omitnil" name:"Total"`

	// 价税合计（大写）
	TotalCn *string `json:"TotalCn,omitnil" name:"TotalCn"`

	// 发票消费类型
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 省
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市
	City *string `json:"City,omitnil" name:"City"`

	// 是否有公司印章（0：没有，1：有）
	CompanySealMark *int64 `json:"CompanySealMark,omitnil" name:"CompanySealMark"`

	// 是否存在二维码（1：有，0：无）
	QRCodeMark *int64 `json:"QRCodeMark,omitnil" name:"QRCodeMark"`

	// 服务类型
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 公司印章内容
	CompanySealContent *string `json:"CompanySealContent,omitnil" name:"CompanySealContent"`

	// 税务局章内容
	TaxSealContent *string `json:"TaxSealContent,omitnil" name:"TaxSealContent"`

	// 条目
	VatRollItems []*VatRollItem `json:"VatRollItems,omitnil" name:"VatRollItems"`
}

type VatInvoiceUserInfo struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 纳税人识别号
	TaxId *string `json:"TaxId,omitnil" name:"TaxId"`

	// 地 址、电 话
	AddrTel *string `json:"AddrTel,omitnil" name:"AddrTel"`

	// 开户行及账号
	FinancialAccount *string `json:"FinancialAccount,omitnil" name:"FinancialAccount"`
}

// Predefined struct for user
type VatInvoiceVerifyNewRequestParams struct {
	// 发票号码，8位、20位（全电票）
	InvoiceNo *string `json:"InvoiceNo,omitnil" name:"InvoiceNo"`

	// 开票日期（不支持当天发票查询，支持五年以内开具的发票），格式：“YYYY-MM-DD”，如：2019-12-20。
	InvoiceDate *string `json:"InvoiceDate,omitnil" name:"InvoiceDate"`

	// 发票代码（10或12 位），全电发票为空。查验未成功超过5次后当日无法再查。
	InvoiceCode *string `json:"InvoiceCode,omitnil" name:"InvoiceCode"`

	// 票种类型 01:增值税专用发票， 02:货运运输业增值税专用发 票， 03:机动车销售统一发票， 04:增值税普通发票， 08:增值税电子专用发票(含全电)， 10:增值税电子普通发票(含全电)， 11:增值税普通发票(卷式)， 14:增值税电子(通行费)发 票， 15:二手车销售统一发票， 32:深圳区块链发票(云南区块链因业务调整现已下线)。
	InvoiceKind *string `json:"InvoiceKind,omitnil" name:"InvoiceKind"`

	// 校验码后 6 位，增值税普通发票、增值税电子普通发票、增值税普通发票(卷式)、增值税电子普通发票(通行费)、全电纸质发票（增值税普通发票）时必填;
	// 区块链为 5 位
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 不含税金额，增值税专用发票、增值税电子专用发票、机动车销售统一发票、二手车销售统一发票、区块链发票时必填; 全电发票为价税合计(含税金额)
	Amount *string `json:"Amount,omitnil" name:"Amount"`

	// 地区编码，通用机打电子发票时必填。
	// 广东:4400，浙江:3300
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// 销方税号，通用机打电子发票必填，区块链发票时必填
	SellerTaxCode *string `json:"SellerTaxCode,omitnil" name:"SellerTaxCode"`

	// 是否开启通用机打电子发票，默认为关闭。
	EnableCommonElectronic *bool `json:"EnableCommonElectronic,omitnil" name:"EnableCommonElectronic"`
}

type VatInvoiceVerifyNewRequest struct {
	*tchttp.BaseRequest
	
	// 发票号码，8位、20位（全电票）
	InvoiceNo *string `json:"InvoiceNo,omitnil" name:"InvoiceNo"`

	// 开票日期（不支持当天发票查询，支持五年以内开具的发票），格式：“YYYY-MM-DD”，如：2019-12-20。
	InvoiceDate *string `json:"InvoiceDate,omitnil" name:"InvoiceDate"`

	// 发票代码（10或12 位），全电发票为空。查验未成功超过5次后当日无法再查。
	InvoiceCode *string `json:"InvoiceCode,omitnil" name:"InvoiceCode"`

	// 票种类型 01:增值税专用发票， 02:货运运输业增值税专用发 票， 03:机动车销售统一发票， 04:增值税普通发票， 08:增值税电子专用发票(含全电)， 10:增值税电子普通发票(含全电)， 11:增值税普通发票(卷式)， 14:增值税电子(通行费)发 票， 15:二手车销售统一发票， 32:深圳区块链发票(云南区块链因业务调整现已下线)。
	InvoiceKind *string `json:"InvoiceKind,omitnil" name:"InvoiceKind"`

	// 校验码后 6 位，增值税普通发票、增值税电子普通发票、增值税普通发票(卷式)、增值税电子普通发票(通行费)、全电纸质发票（增值税普通发票）时必填;
	// 区块链为 5 位
	CheckCode *string `json:"CheckCode,omitnil" name:"CheckCode"`

	// 不含税金额，增值税专用发票、增值税电子专用发票、机动车销售统一发票、二手车销售统一发票、区块链发票时必填; 全电发票为价税合计(含税金额)
	Amount *string `json:"Amount,omitnil" name:"Amount"`

	// 地区编码，通用机打电子发票时必填。
	// 广东:4400，浙江:3300
	RegionCode *string `json:"RegionCode,omitnil" name:"RegionCode"`

	// 销方税号，通用机打电子发票必填，区块链发票时必填
	SellerTaxCode *string `json:"SellerTaxCode,omitnil" name:"SellerTaxCode"`

	// 是否开启通用机打电子发票，默认为关闭。
	EnableCommonElectronic *bool `json:"EnableCommonElectronic,omitnil" name:"EnableCommonElectronic"`
}

func (r *VatInvoiceVerifyNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatInvoiceVerifyNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoiceNo")
	delete(f, "InvoiceDate")
	delete(f, "InvoiceCode")
	delete(f, "InvoiceKind")
	delete(f, "CheckCode")
	delete(f, "Amount")
	delete(f, "RegionCode")
	delete(f, "SellerTaxCode")
	delete(f, "EnableCommonElectronic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VatInvoiceVerifyNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VatInvoiceVerifyNewResponseParams struct {
	// 增值税发票、购车发票、全电发票的基础要素字段信息。
	Invoice *VatInvoice `json:"Invoice,omitnil" name:"Invoice"`

	// 机动车销售统一发票详细字段信息。
	VehicleInvoiceInfo *VehicleInvoiceInfo `json:"VehicleInvoiceInfo,omitnil" name:"VehicleInvoiceInfo"`

	// 二手车销售统一发票详细字段信息。
	UsedVehicleInvoiceInfo *UsedVehicleInvoiceInfo `json:"UsedVehicleInvoiceInfo,omitnil" name:"UsedVehicleInvoiceInfo"`

	// 通行费发票详细字段信息。
	PassInvoiceInfoList []*PassInvoiceInfo `json:"PassInvoiceInfoList,omitnil" name:"PassInvoiceInfoList"`

	// 全电发票（铁路电子客票）详细字段信息。
	ElectronicTrainTicket *ElectronicTrainTicket `json:"ElectronicTrainTicket,omitnil" name:"ElectronicTrainTicket"`

	// 全电发票（航空运输电子客票行程单）详细字段信息。
	ElectronicAirTransport *ElectronicAirTransport `json:"ElectronicAirTransport,omitnil" name:"ElectronicAirTransport"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VatInvoiceVerifyNewResponse struct {
	*tchttp.BaseResponse
	Response *VatInvoiceVerifyNewResponseParams `json:"Response"`
}

func (r *VatInvoiceVerifyNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VatInvoiceVerifyNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VatInvoiceVerifyRequestParams struct {
	// 发票代码， 一张发票一天只能查询5次。
	InvoiceCode *string `json:"InvoiceCode,omitnil" name:"InvoiceCode"`

	// 发票号码（8位）
	InvoiceNo *string `json:"InvoiceNo,omitnil" name:"InvoiceNo"`

	// 开票日期（不支持当天发票查询，支持五年以内开具的发票），格式：“YYYY-MM-DD”，如：2019-12-20。
	InvoiceDate *string `json:"InvoiceDate,omitnil" name:"InvoiceDate"`

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
	Additional *string `json:"Additional,omitnil" name:"Additional"`
}

type VatInvoiceVerifyRequest struct {
	*tchttp.BaseRequest
	
	// 发票代码， 一张发票一天只能查询5次。
	InvoiceCode *string `json:"InvoiceCode,omitnil" name:"InvoiceCode"`

	// 发票号码（8位）
	InvoiceNo *string `json:"InvoiceNo,omitnil" name:"InvoiceNo"`

	// 开票日期（不支持当天发票查询，支持五年以内开具的发票），格式：“YYYY-MM-DD”，如：2019-12-20。
	InvoiceDate *string `json:"InvoiceDate,omitnil" name:"InvoiceDate"`

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
	Additional *string `json:"Additional,omitnil" name:"Additional"`
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

// Predefined struct for user
type VatInvoiceVerifyResponseParams struct {
	// 增值税发票信息，详情请点击左侧链接。
	Invoice *VatInvoice `json:"Invoice,omitnil" name:"Invoice"`

	// 机动车销售统一发票信息
	VehicleInvoiceInfo *VehicleInvoiceInfo `json:"VehicleInvoiceInfo,omitnil" name:"VehicleInvoiceInfo"`

	// 二手车销售统一发票信息
	UsedVehicleInvoiceInfo *UsedVehicleInvoiceInfo `json:"UsedVehicleInvoiceInfo,omitnil" name:"UsedVehicleInvoiceInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VatInvoiceVerifyResponse struct {
	*tchttp.BaseResponse
	Response *VatInvoiceVerifyResponseParams `json:"Response"`
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
	// 发票代码、合计金额(小写)、合计金额(大写)、开票日期、发票号码、购买方识别号、销售方识别号、校验码、销售方名称、购买方名称、发票消费类型、省、市、是否有公司印章、单价、金额、数量、服务类型、品名、种类。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 文本行在旋转纠正之后的图像中的像素坐标。
	Rect *Rect `json:"Rect,omitnil" name:"Rect"`
}

// Predefined struct for user
type VatRollInvoiceOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type VatRollInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VatRollInvoiceOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VatRollInvoiceOCRResponseParams struct {
	// 增值税发票（卷票）识别结果，具体内容请点击左侧链接。
	VatRollInvoiceInfos []*VatRollInvoiceInfo `json:"VatRollInvoiceInfos,omitnil" name:"VatRollInvoiceInfos"`

	// 图片旋转角度（角度制），文本的水平方向为0°，顺时针为正，逆时针为负。
	Angle *float64 `json:"Angle,omitnil" name:"Angle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VatRollInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *VatRollInvoiceOCRResponseParams `json:"Response"`
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

type VatRollItem struct {
	// 项目名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数量
	Quantity *string `json:"Quantity,omitnil" name:"Quantity"`

	// 单价
	Price *string `json:"Price,omitnil" name:"Price"`

	// 金额
	Total *string `json:"Total,omitnil" name:"Total"`
}

type VehicleInvoiceInfo struct {
	// 车辆类型
	CarType *string `json:"CarType,omitnil" name:"CarType"`

	// 厂牌型号
	PlateModel *string `json:"PlateModel,omitnil" name:"PlateModel"`

	// 产地
	ProduceAddress *string `json:"ProduceAddress,omitnil" name:"ProduceAddress"`

	// 合格证号
	CertificateNo *string `json:"CertificateNo,omitnil" name:"CertificateNo"`

	// 进口证明书号
	ImportNo *string `json:"ImportNo,omitnil" name:"ImportNo"`

	// LSVCA2NP9HN0xxxxx
	VinNo *string `json:"VinNo,omitnil" name:"VinNo"`

	// 完税证书号
	PayTaxesNo *string `json:"PayTaxesNo,omitnil" name:"PayTaxesNo"`

	// 吨位
	Tonnage *string `json:"Tonnage,omitnil" name:"Tonnage"`

	// 限乘人数
	LimitCount *string `json:"LimitCount,omitnil" name:"LimitCount"`

	// 发动机号码
	EngineNo *string `json:"EngineNo,omitnil" name:"EngineNo"`

	// 商检单号
	BizCheckFormNo *string `json:"BizCheckFormNo,omitnil" name:"BizCheckFormNo"`

	// 主管税务机关代码
	TaxtationOrgCode *string `json:"TaxtationOrgCode,omitnil" name:"TaxtationOrgCode"`

	// 主管税务机关名称
	TaxtationOrgName *string `json:"TaxtationOrgName,omitnil" name:"TaxtationOrgName"`

	// 税率
	MotorTaxRate *string `json:"MotorTaxRate,omitnil" name:"MotorTaxRate"`

	// 银行账号
	MotorBankName *string `json:"MotorBankName,omitnil" name:"MotorBankName"`

	// 开户行
	MotorBankAccount *string `json:"MotorBankAccount,omitnil" name:"MotorBankAccount"`

	// 销售地址
	SellerAddress *string `json:"SellerAddress,omitnil" name:"SellerAddress"`

	// 销售电话
	SellerTel *string `json:"SellerTel,omitnil" name:"SellerTel"`

	// 购方身份证
	BuyerNo *string `json:"BuyerNo,omitnil" name:"BuyerNo"`
}

// Predefined struct for user
type VehicleLicenseOCRRequestParams struct {
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT 为行驶证主页正面（有红色印章的一面），
	// BACK 为行驶证副页正面（有号码号牌的一面），
	// DOUBLE 为行驶证主页正面和副页正面。
	// 默认值为：FRONT。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
}

type VehicleLicenseOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。要求图片经Base64编码后不超过 7M，分辨率建议500*800以上，支持PNG、JPG、JPEG、BMP格式。建议卡片部分占据图片2/3以上。图片下载时间不超过 3 秒。
	// 建议图片存储于腾讯云，可保障更高的下载速度和稳定性。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// FRONT 为行驶证主页正面（有红色印章的一面），
	// BACK 为行驶证副页正面（有号码号牌的一面），
	// DOUBLE 为行驶证主页正面和副页正面。
	// 默认值为：FRONT。
	CardSide *string `json:"CardSide,omitnil" name:"CardSide"`
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

// Predefined struct for user
type VehicleLicenseOCRResponseParams struct {
	// 行驶证主页正面的识别结果，CardSide 为 FRONT。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrontInfo *TextVehicleFront `json:"FrontInfo,omitnil" name:"FrontInfo"`

	// 行驶证副页正面的识别结果，CardSide 为 BACK。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackInfo *TextVehicleBack `json:"BackInfo,omitnil" name:"BackInfo"`

	// Code 告警码列表和释义：
	// -9102 复印件告警
	// -9103 翻拍件告警
	// 注：告警码可以同时存在多个
	RecognizeWarnCode []*int64 `json:"RecognizeWarnCode,omitnil" name:"RecognizeWarnCode"`

	// 告警码说明：
	// WARN_DRIVER_LICENSE_COPY_CARD 复印件告警
	// WARN_DRIVER_LICENSE_SCREENED_CARD 翻拍件告警
	// 注：告警信息可以同时存在多个
	RecognizeWarnMsg []*string `json:"RecognizeWarnMsg,omitnil" name:"RecognizeWarnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VehicleLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *VehicleLicenseOCRResponseParams `json:"Response"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type VehicleRegCertOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type VehicleRegCertOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type VehicleRegCertOCRResponseParams struct {
	// 机动车登记证书识别结果，具体内容请点击左侧链接。
	VehicleRegCertInfos []*VehicleRegCertInfo `json:"VehicleRegCertInfos,omitnil" name:"VehicleRegCertInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VehicleRegCertOCRResponse struct {
	*tchttp.BaseResponse
	Response *VehicleRegCertOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type VerifyOfdVatInvoiceOCRRequestParams struct {
	// OFD文件的 Url 地址。
	OfdFileUrl *string `json:"OfdFileUrl,omitnil" name:"OfdFileUrl"`

	// OFD文件的 Base64 值。
	// OfdFileUrl 和 OfdFileBase64 必传其一，若两者都传，只解析OfdFileBase64。
	OfdFileBase64 *string `json:"OfdFileBase64,omitnil" name:"OfdFileBase64"`
}

type VerifyOfdVatInvoiceOCRRequest struct {
	*tchttp.BaseRequest
	
	// OFD文件的 Url 地址。
	OfdFileUrl *string `json:"OfdFileUrl,omitnil" name:"OfdFileUrl"`

	// OFD文件的 Base64 值。
	// OfdFileUrl 和 OfdFileBase64 必传其一，若两者都传，只解析OfdFileBase64。
	OfdFileBase64 *string `json:"OfdFileBase64,omitnil" name:"OfdFileBase64"`
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

// Predefined struct for user
type VerifyOfdVatInvoiceOCRResponseParams struct {
	// 发票类型
	// 026:增值税电子普通发票
	// 028:增值税电子专用发票
	// 010:电子发票（普通发票）
	// 020:电子发票（增值税专用发票）
	// 030:电子发票（铁路电子客票）
	// 040:电子发票（航空运输电子客票行程单）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 发票代码
	InvoiceCode *string `json:"InvoiceCode,omitnil" name:"InvoiceCode"`

	// 发票号码
	InvoiceNumber *string `json:"InvoiceNumber,omitnil" name:"InvoiceNumber"`

	// 开票日期
	IssueDate *string `json:"IssueDate,omitnil" name:"IssueDate"`

	// 验证码
	InvoiceCheckCode *string `json:"InvoiceCheckCode,omitnil" name:"InvoiceCheckCode"`

	// 机器编号
	MachineNumber *string `json:"MachineNumber,omitnil" name:"MachineNumber"`

	// 密码区
	TaxControlCode *string `json:"TaxControlCode,omitnil" name:"TaxControlCode"`

	// 购买方
	Buyer *VatInvoiceUserInfo `json:"Buyer,omitnil" name:"Buyer"`

	// 销售方
	Seller *VatInvoiceUserInfo `json:"Seller,omitnil" name:"Seller"`

	// 价税合计
	TaxInclusiveTotalAmount *string `json:"TaxInclusiveTotalAmount,omitnil" name:"TaxInclusiveTotalAmount"`

	// 开票人
	InvoiceClerk *string `json:"InvoiceClerk,omitnil" name:"InvoiceClerk"`

	// 收款人
	Payee *string `json:"Payee,omitnil" name:"Payee"`

	// 复核人
	Checker *string `json:"Checker,omitnil" name:"Checker"`

	// 税额
	TaxTotalAmount *string `json:"TaxTotalAmount,omitnil" name:"TaxTotalAmount"`

	// 不含税金额
	TaxExclusiveTotalAmount *string `json:"TaxExclusiveTotalAmount,omitnil" name:"TaxExclusiveTotalAmount"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`

	// 货物或服务清单
	GoodsInfos []*VatInvoiceGoodsInfo `json:"GoodsInfos,omitnil" name:"GoodsInfos"`

	// 航空运输电子客票行程单信息
	AirTicketInfo *AirTicketInfo `json:"AirTicketInfo,omitnil" name:"AirTicketInfo"`

	// 铁路电子客票
	RailwayTicketInfo *RailwayTicketInfo `json:"RailwayTicketInfo,omitnil" name:"RailwayTicketInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyOfdVatInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *VerifyOfdVatInvoiceOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type VinOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
}

type VinOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`
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

// Predefined struct for user
type VinOCRResponseParams struct {
	// 检测到的车辆 VIN 码。
	Vin *string `json:"Vin,omitnil" name:"Vin"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VinOCRResponse struct {
	*tchttp.BaseResponse
	Response *VinOCRResponseParams `json:"Response"`
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

// Predefined struct for user
type WaybillOCRRequestParams struct {
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 预检测开关，当待识别运单占整个输入图像的比例较小时，建议打开预检测开关。默认值为false。
	EnablePreDetect *bool `json:"EnablePreDetect,omitnil" name:"EnablePreDetect"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
}

type WaybillOCRRequest struct {
	*tchttp.BaseRequest
	
	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageUrl。
	ImageBase64 *string `json:"ImageBase64,omitnil" name:"ImageBase64"`

	// 图片的 Url 地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 7M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil" name:"ImageUrl"`

	// 预检测开关，当待识别运单占整个输入图像的比例较小时，建议打开预检测开关。默认值为false。
	EnablePreDetect *bool `json:"EnablePreDetect,omitnil" name:"EnablePreDetect"`

	// 是否开启PDF识别，默认值为true，开启后可同时支持图片和PDF的识别。
	IsPdf *bool `json:"IsPdf,omitnil" name:"IsPdf"`

	// 需要识别的PDF页面的对应页码，仅支持PDF单页识别，当上传文件为PDF且IsPdf参数值为true时有效，默认值为1。
	PdfPageNumber *int64 `json:"PdfPageNumber,omitnil" name:"PdfPageNumber"`
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
	delete(f, "IsPdf")
	delete(f, "PdfPageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WaybillOCRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WaybillOCRResponseParams struct {
	// 检测到的文本信息，具体内容请点击左侧链接。
	TextDetections *TextWaybill `json:"TextDetections,omitnil" name:"TextDetections"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type WaybillOCRResponse struct {
	*tchttp.BaseResponse
	Response *WaybillOCRResponseParams `json:"Response"`
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
	Text *string `json:"Text,omitnil" name:"Text"`
}

type WordCoordPoint struct {
	// 英文OCR识别出的每个单词在原图中的四点坐标。
	WordCoordinate []*Coord `json:"WordCoordinate,omitnil" name:"WordCoordinate"`
}

type WordItem struct {
	// 文本块内容
	DetectedText *string `json:"DetectedText,omitnil" name:"DetectedText"`

	// 四点坐标
	Coord *Polygon `json:"Coord,omitnil" name:"Coord"`
}

type Words struct {
	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 候选字Character
	Character *string `json:"Character,omitnil" name:"Character"`
}