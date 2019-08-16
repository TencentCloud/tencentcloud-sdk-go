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

	// 结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 保留字段，暂无意义
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示（预留字段，目前不支持）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// 此字段为扩展字段
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
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
