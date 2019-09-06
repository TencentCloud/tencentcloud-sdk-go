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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ArithmeticOCRRequest struct {
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

func (r *ArithmeticOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArithmeticOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ArithmeticOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextArithmetic `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArithmeticOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ArithmeticOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCardOCRRequest struct {
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

func (r *BankCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCardOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BankCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 卡号
		CardNo *string `json:"CardNo,omitempty" name:"CardNo"`

		// 银行信息
		BankInfo *string `json:"BankInfo,omitempty" name:"BankInfo"`

		// 有效期
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BankCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BankCardOCRResponse) FromJsonString(s string) error {
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

func (r *BizLicenseOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BizLicenseOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 注册号
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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BizLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BizLicenseOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BusInvoiceInfo struct {

	// 识别出的字段名称（关键字）。
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

func (r *BusInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BusInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BusInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 汽车票识别结果，具体内容请点击左侧链接。
		BusInvoiceInfos []*BusInvoiceInfo `json:"BusInvoiceInfos,omitempty" name:"BusInvoiceInfos" list`

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

func (r *BusInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BusinessCardInfo struct {

	// 识别出的字段名称（关键字，可能重复，比如多个手机），能识别的字段名为：
	// 姓名、英文姓名、英文地址、公司、英文公司、职位、英文职位、部门、英文部门、手机、电话、传真、社交帐号、QQ、MSN、微信、微博、邮箱、邮编、网址、公司账号、其他。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
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

func (r *BusinessCardOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BusinessCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 名片识别结果，具体内容请点击左侧链接。
		BusinessCardInfos []*BusinessCardInfo `json:"BusinessCardInfos,omitempty" name:"BusinessCardInfos" list`

		// 返回图像预处理后的图片，图像预处理未开启时返回内容为空。
		RetImageBase64 *string `json:"RetImageBase64,omitempty" name:"RetImageBase64"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BusinessCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BusinessCardOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CarInvoiceInfo struct {

	// 识别出的字段名称（关键字）。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CarInvoiceOCRRequest struct {
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

func (r *CarInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CarInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CarInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 购车发票识别结果，具体内容请点击左侧链接。
		CarInvoiceInfos []*CarInvoiceInfo `json:"CarInvoiceInfos,omitempty" name:"CarInvoiceInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CarInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CarInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Coord struct {

	// 横坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// 纵坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type DriverLicenseOCRRequest struct {
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

func (r *DriverLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DriverLicenseOCRRequest) FromJsonString(s string) error {
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

		// 出生日期
		DateOfBirth *string `json:"DateOfBirth,omitempty" name:"DateOfBirth"`

		// 初次领证日期
		DateOfFirstIssue *string `json:"DateOfFirstIssue,omitempty" name:"DateOfFirstIssue"`

		// 准驾车型
		Class *string `json:"Class,omitempty" name:"Class"`

		// 有效期开始时间
		StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

		// 有效期截止时间
		EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

		// 证号
		CardCode *string `json:"CardCode,omitempty" name:"CardCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DriverLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DriverLicenseOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DutyPaidProofInfo struct {

	// 识别出的字段名称（关键字）。
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

func (r *DutyPaidProofOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DutyPaidProofOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DutyPaidProofOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 完税证明识别结果，具体内容请点击左侧链接。
		DutyPaidProofInfos []*DutyPaidProofInfo `json:"DutyPaidProofInfos,omitempty" name:"DutyPaidProofInfos" list`

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

func (r *DutyPaidProofOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnglishOCRRequest struct {
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

func (r *EnglishOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnglishOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnglishOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextDetectionEn `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnglishOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnglishOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FlightInvoiceInfo struct {

	// 识别出的字段名称（关键字）。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段 Name 对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type FlightInvoiceOCRRequest struct {
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

func (r *FlightInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlightInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FlightInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机票行程单识别结果，具体内容请点击左侧链接。
		FlightInvoiceInfos []*FlightInvoiceInfo `json:"FlightInvoiceInfos,omitempty" name:"FlightInvoiceInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlightInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlightInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralAccurateOCRRequest struct {
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

func (r *GeneralAccurateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralAccurateOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralAccurateOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralAccurateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralAccurateOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRRequest struct {
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

	// 保留字段。
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 识别语言类型。
	// 支持自动识别语言类型，同时支持自选语言种类，默认中英文混合(zh)。
	// 可选值：
	// zh\auto\jap\kor\
	// spa\fre\ger\por\
	// vie\may\rus\ita\
	// hol\swe\fin\dan\
	// nor\hun\tha\lat
	// 可选值分别表示：
	// 中英文混合、自动识别、日语、韩语、
	// 西班牙语、法语、德语、葡萄牙语、
	// 越南语、马来语、俄语、意大利语、
	// 荷兰语、瑞典语、芬兰语、丹麦语、
	// 挪威语、匈牙利语、泰语、拉丁语系。
	LanguageType *string `json:"LanguageType,omitempty" name:"LanguageType"`
}

func (r *GeneralBasicOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 检测到的语言类型，目前支持的语言类型参考入参LanguageType说明。
		Language *string `json:"Language,omitempty" name:"Language"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralBasicOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralEfficientOCRRequest struct {
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

func (r *GeneralEfficientOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralEfficientOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralEfficientOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralEfficientOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralEfficientOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRRequest struct {
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

func (r *GeneralFastOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralFastOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 检测到的语言，目前支持的语种范围为：简体中文、繁体中文、英文、日文、韩文。未来将陆续新增对更多语种的支持。
	// 返回结果含义为：zh - 中英混合，jap - 日文，kor - 韩文。
		Language *string `json:"Language,omitempty" name:"Language"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralFastOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralFastOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralHandwritingOCRRequest struct {
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

func (r *GeneralHandwritingOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralHandwritingOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralHandwritingOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextGeneralHandwriting `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralHandwritingOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralHandwritingOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IDCardOCRRequest struct {
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

	// FRONT 为身份证有照片的一面（人像面），
	// BACK 为身份证有国徽的一面（国徽面）。
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`

	// 可选字段，根据需要选择是否请求对应字段。
	// 目前包含的字段为：
	// CropIdCard，身份证照片裁剪，bool 类型，默认false，
	// CropPortrait，人像照片裁剪，bool 类型，默认false，
	// CopyWarn，复印件告警，bool 类型，默认false，
	// BorderCheckWarn，边框和框内遮挡告警，bool 类型，默认false，
	// ReshootWarn，翻拍告警，bool 类型，默认false，
	// DetectPsWarn，PS检测告警，bool类型，默认false，
	// TempIdWarn，临时身份证告警，bool类型，默认false，
	// InvalidDateWarn，身份证有效日期不合法告警，bool类型，默认false。
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

func (r *IDCardOCRRequest) FromJsonString(s string) error {
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

		// 扩展信息，根据请求的可选字段返回对应内容，不请求则不返回，具体输入参考示例3和示例4。
	// 
	// 目前支持的扩展字段为：
	// IdCard，身份证照片，请求 CropIdCard 时返回；
	// Portrait，人像照片，请求 CropPortrait 时返回；
	// WarnInfos，告警信息（Code - 告警码），识别出以下告警内容时返回。
	// 
	// Code 告警码列表和释义：
	// -9100	身份证有效日期不合法告警，
	// -9101	身份证边框不完整告警，
	// -9102	身份证复印件告警，
	// -9103	身份证翻拍告警，
	// -9105	身份证框内遮挡告警，
	// -9104	临时身份证告警，
	// -9106	身份证 PS 告警。
		AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IDCardOCRResponse) FromJsonString(s string) error {
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

	// 识别出的图片在混贴票据图片中的位置信息。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`

	// 入参 ReturnImage 为 True 时返回 Base64 编码后的图片。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitempty" name:"Image"`
}

type InvoiceGeneralInfo struct {

	// 识别出的字段名称（关键字）。
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

func (r *InvoiceGeneralOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InvoiceGeneralOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InvoiceGeneralOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通用机打发票识别结果，具体内容请点击左侧链接。
		InvoiceGeneralInfos []*InvoiceGeneralInfo `json:"InvoiceGeneralInfos,omitempty" name:"InvoiceGeneralInfos" list`

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

func (r *InvoiceGeneralOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LicensePlateOCRRequest struct {
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

func (r *LicensePlateOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LicensePlateOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LicensePlateOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别出的车牌号码。
		Number *string `json:"Number,omitempty" name:"Number"`

		// 置信度，0 - 100 之间。
		Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LicensePlateOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LicensePlateOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MLIDCardOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的 Url 地址。( 中国地区之外不支持这个字段 )
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 是否返回图片
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

func (r *MLIDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDCardOCRRequest) FromJsonString(s string) error {
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
		Warn []*int64 `json:"Warn,omitempty" name:"Warn" list`

		// 证件图片
		Image *string `json:"Image,omitempty" name:"Image"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MLIDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDCardOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MLIDPassportOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的 Base64 值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。图片下载时间不超过 3 秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 是否返回图片
	RetImage *bool `json:"RetImage,omitempty" name:"RetImage"`
}

func (r *MLIDPassportOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDPassportOCRRequest) FromJsonString(s string) error {
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
		Warn []*int64 `json:"Warn,omitempty" name:"Warn" list`

		// 证件图片
		Image *string `json:"Image,omitempty" name:"Image"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MLIDPassportOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MLIDPassportOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceDetectRequest struct {
	*tchttp.BaseRequest

	// 是否需要返回裁剪后的图片。
	ReturnImage *bool `json:"ReturnImage,omitempty" name:"ReturnImage"`

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

func (r *MixedInvoiceDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MixedInvoiceDetectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceDetectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测出的票据类型列表，具体内容请点击左侧链接。
		InvoiceDetectInfos []*InvoiceDetectInfo `json:"InvoiceDetectInfos,omitempty" name:"InvoiceDetectInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MixedInvoiceDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MixedInvoiceDetectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceItem struct {

	// 识别结果。
	// OK：表示识别成功；FailedOperation.UnsupportedInvioce：表示不支持识别；
	// FailedOperation.UnKnowError： 表示识别失败；
	// 其它错误码见各个票据接口的定义。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 识别出的图片所属的票据类型。
	// -1：未知类型
	// 0：出租车发票
	// 1：定额发票
	// 2：火车票
	// 3：增值税发票
	// 4：客运限额发票（仅支持类型检测，不支持识别）
	// 5：机票行程单
	// 6：酒店账单（仅支持类型检测，不支持识别）
	// 7：完税证明
	// 8：通用机打发票
	// 9：汽车票
	// 10：轮船票
	// 11：增值税发票（卷票 ）
	// 12：购车发票
	// 13：过路过桥费发票
	// 14：购物小票（仅支持类型检测，不支持识别）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 识别出的图片在混贴票据图片中的位置信息。
	Rect *Rect `json:"Rect,omitempty" name:"Rect"`

	// 识别出的图片在混贴票据图片中的旋转角度。
	Angle *float64 `json:"Angle,omitempty" name:"Angle"`

	// 识别到的内容。
	SingleInvoiceInfos []*SingleInvoiceInfo `json:"SingleInvoiceInfos,omitempty" name:"SingleInvoiceInfos" list`
}

type MixedInvoiceOCRRequest struct {
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

func (r *MixedInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MixedInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MixedInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 混贴票据识别结果，具体内容请点击左侧链接。
		MixedInvoiceItems []*MixedInvoiceItem `json:"MixedInvoiceItems,omitempty" name:"MixedInvoiceItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MixedInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MixedInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PassportOCRRequest struct {
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

	// 护照类型（默认CN）：
	// CN：支持中国大陆居民护照，字段较多，精度更高；
	// HK：支持中国香港护照（部分主要字段）；
	// GENERAL：支持国外护照（部分主要字段）；
	// THAI：支持泰国护照（部分主要字段）。
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *PassportOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PassportOCRRequest) FromJsonString(s string) error {
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

func (r *PassportOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PermitOCRRequest struct {
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

func (r *PermitOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PermitOCRRequest) FromJsonString(s string) error {
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

func (r *PermitOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *QuotaInvoiceOCRRequest) FromJsonString(s string) error {
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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuotaInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuotaInvoiceOCRResponse) FromJsonString(s string) error {
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

type ShipInvoiceInfo struct {

	// 识别出的字段名称（关键字）。
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

func (r *ShipInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShipInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShipInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 轮船票识别结果，具体内容请点击左侧链接。
		ShipInvoiceInfos []*ShipInvoiceInfo `json:"ShipInvoiceInfos,omitempty" name:"ShipInvoiceInfos" list`

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

func (r *ShipInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SingleInvoiceInfo struct {

	// 识别出的字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
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

func (r *TableOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TableOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		TextDetections []*TextTable `json:"TextDetections,omitempty" name:"TextDetections" list`

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

func (r *TableOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TaxiInvoiceOCRRequest struct {
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

func (r *TaxiInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TaxiInvoiceOCRRequest) FromJsonString(s string) error {
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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TaxiInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TaxiInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextArithmetic struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 算式运算结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 保留字段，暂不支持
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示（保留字段，暂不支持）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// 保留字段，暂不支持
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextDetectRequest struct {
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

func (r *TextDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextDetectRequest) FromJsonString(s string) error {
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
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// 此字段为扩展字段。
	// GeneralBasicOcr接口返回段落信息Parag，包含ParagNo。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextDetectionEn struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// 此字段为扩展字段。目前EnglishOCR接口返回内容为空。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextGeneralHandwriting struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 置信度 0 - 100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// 此字段为扩展字段
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
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
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// 此字段为扩展字段
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type TextVatInvoice struct {

	// 识别出的字段名称（关键字）。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 识别出的字段名称对应的值，也就是字段Name对应的字符串结果。
	Value *string `json:"Value,omitempty" name:"Value"`
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

	// 识别出的字段名称（关键字）。
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

func (r *TollInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TollInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TollInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 过路过桥费发票识别结果，具体内容请点击左侧链接。
		TollInvoiceInfos []*TollInvoiceInfo `json:"TollInvoiceInfos,omitempty" name:"TollInvoiceInfos" list`

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

func (r *TollInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TrainTicketOCRRequest struct {
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

func (r *TrainTicketOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrainTicketOCRRequest) FromJsonString(s string) error {
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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TrainTicketOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TrainTicketOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VatInvoiceOCRRequest struct {
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

func (r *VatInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VatInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VatInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接。
		VatInvoiceInfos []*TextVatInvoice `json:"VatInvoiceInfos,omitempty" name:"VatInvoiceInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VatInvoiceOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VatInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VatRollInvoiceInfo struct {

	// 识别出的字段名称（关键字）。
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

func (r *VatRollInvoiceOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VatRollInvoiceOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VatRollInvoiceOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 增值税发票（卷票）识别结果，具体内容请点击左侧链接。
		VatRollInvoiceInfos []*VatRollInvoiceInfo `json:"VatRollInvoiceInfos,omitempty" name:"VatRollInvoiceInfos" list`

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

func (r *VatRollInvoiceOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VehicleLicenseOCRRequest struct {
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

	// FRONT 为行驶证主页正面（有红色印章的一面），
	// BACK 为行驶证副页正面（有号码号牌的一面）。
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`
}

func (r *VehicleLicenseOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VehicleLicenseOCRRequest) FromJsonString(s string) error {
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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VehicleLicenseOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VehicleLicenseOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VinOCRRequest struct {
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

func (r *VinOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VinOCRRequest) FromJsonString(s string) error {
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

func (r *VinOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WaybillOCRRequest struct {
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

func (r *WaybillOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WaybillOCRRequest) FromJsonString(s string) error {
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

func (r *WaybillOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WaybillObj struct {

	// 识别出的文本行内容
	Text *string `json:"Text,omitempty" name:"Text"`
}
