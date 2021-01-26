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

package v20200513

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AuthInfo struct {

	// 主键
	Id *string `json:"Id,omitempty" name:"Id"`

	// 授权人名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证号/社会信用代码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 授权人类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type BatchDescribeOrderCertificateRequest struct {
	*tchttp.BaseRequest

	// 要下载授权书的订单id
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds" list`
}

func (r *BatchDescribeOrderCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDescribeOrderCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDescribeOrderCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 授权书的下载地址
		CertificateUrls []*string `json:"CertificateUrls,omitempty" name:"CertificateUrls" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDescribeOrderCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDescribeOrderCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDescribeOrderImageRequest struct {
	*tchttp.BaseRequest

	// 要下载图片的订单id
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds" list`
}

func (r *BatchDescribeOrderImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDescribeOrderImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchDescribeOrderImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片的下载地址
		ImageUrls []*string `json:"ImageUrls,omitempty" name:"ImageUrls" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchDescribeOrderImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchDescribeOrderImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOrderAndPayRequest struct {
	*tchttp.BaseRequest

	// 图片ID
	ImageId *uint64 `json:"ImageId,omitempty" name:"ImageId"`

	// 授权人ID
	AuthUserId *string `json:"AuthUserId,omitempty" name:"AuthUserId"`

	// 售卖组合id
	MarshalId *uint64 `json:"MarshalId,omitempty" name:"MarshalId"`
}

func (r *CreateOrderAndPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOrderAndPayRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateOrderAndPayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrderAndPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateOrderAndPayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthUsersRequest struct {
	*tchttp.BaseRequest

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAuthUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthUsersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAuthUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 授权人信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Users []*AuthInfo `json:"Users,omitempty" name:"Users" list`

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 是否是老策略用户
		OldUser *bool `json:"OldUser,omitempty" name:"OldUser"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuthUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAuthUsersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageRequest struct {
	*tchttp.BaseRequest

	// 图片ID
	ImageId *uint64 `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片ID
		ImageId *uint64 `json:"ImageId,omitempty" name:"ImageId"`

		// 图片标题
		Title *string `json:"Title,omitempty" name:"Title"`

		// 图片描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 图片预览链接
		PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

		// 图片缩略图
		ThumbUrl *string `json:"ThumbUrl,omitempty" name:"ThumbUrl"`

		// 图片供应商
		Vendor *string `json:"Vendor,omitempty" name:"Vendor"`

		// 图片售卖组合信息
		Marshals []*ImageMarshal `json:"Marshals,omitempty" name:"Marshals" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest

	// 页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 构图方式，可选以下值：horizontal、vertical、square，分别代表以下含义：横图、竖图、方图
	Orientation *string `json:"Orientation,omitempty" name:"Orientation"`

	// 图片类型，可选以下值：照片、插画
	ImageSenseType *string `json:"ImageSenseType,omitempty" name:"ImageSenseType"`

	// 分层图库id数组，可选以下数值：1(基础)，2(精选)，3(高级)
	LayeredGalleryIds []*int64 `json:"LayeredGalleryIds,omitempty" name:"LayeredGalleryIds" list`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 页偏移量
		Offset *int64 `json:"Offset,omitempty" name:"Offset"`

		// 页大小
		Limit *int64 `json:"Limit,omitempty" name:"Limit"`

		// 总条数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 是否有下一页
		HaveMore *bool `json:"HaveMore,omitempty" name:"HaveMore"`

		// 图片信息数组
		Items []*ImageItem `json:"Items,omitempty" name:"Items" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageItem struct {

	// 图片ID
	ImageId *uint64 `json:"ImageId,omitempty" name:"ImageId"`

	// 图片标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 图片描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 图片预览链接
	PreviewUrl *string `json:"PreviewUrl,omitempty" name:"PreviewUrl"`

	// 图片缩略图
	ThumbUrl *string `json:"ThumbUrl,omitempty" name:"ThumbUrl"`

	// 图片供应商
	Vendor *string `json:"Vendor,omitempty" name:"Vendor"`

	// 图片关键词
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`
}

type ImageMarshal struct {

	// 售卖组合唯一标识
	MarshalId *uint64 `json:"MarshalId,omitempty" name:"MarshalId"`

	// 图片高度
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 图片宽度
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 图片大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 图片格式
	Format *string `json:"Format,omitempty" name:"Format"`

	// 图片价格(单位:分)
	Price *uint64 `json:"Price,omitempty" name:"Price"`

	// 授权范围
	LicenseScope *string `json:"LicenseScope,omitempty" name:"LicenseScope"`

	// 是否支持VIP购买
	IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`
}
