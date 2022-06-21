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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
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

// Predefined struct for user
type BatchDescribeOrderCertificateRequestParams struct {
	// 要下载授权书的订单id
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`
}

type BatchDescribeOrderCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 要下载授权书的订单id
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`
}

func (r *BatchDescribeOrderCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeOrderCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDescribeOrderCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeOrderCertificateResponseParams struct {
	// 授权书的下载地址
	CertificateUrls []*string `json:"CertificateUrls,omitempty" name:"CertificateUrls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDescribeOrderCertificateResponse struct {
	*tchttp.BaseResponse
	Response *BatchDescribeOrderCertificateResponseParams `json:"Response"`
}

func (r *BatchDescribeOrderCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeOrderCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeOrderImageRequestParams struct {
	// 要下载图片的订单id
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`
}

type BatchDescribeOrderImageRequest struct {
	*tchttp.BaseRequest
	
	// 要下载图片的订单id
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`
}

func (r *BatchDescribeOrderImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeOrderImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDescribeOrderImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDescribeOrderImageResponseParams struct {
	// 图片的下载地址
	ImageUrls []*string `json:"ImageUrls,omitempty" name:"ImageUrls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDescribeOrderImageResponse struct {
	*tchttp.BaseResponse
	Response *BatchDescribeOrderImageResponseParams `json:"Response"`
}

func (r *BatchDescribeOrderImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDescribeOrderImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrderAndDownloadsRequestParams struct {
	// ImageId必填，单张购买，所有必填，会员身份可以省略部分参数
	ImageInfos []*ImageInfo `json:"ImageInfos,omitempty" name:"ImageInfos"`
}

type CreateOrderAndDownloadsRequest struct {
	*tchttp.BaseRequest
	
	// ImageId必填，单张购买，所有必填，会员身份可以省略部分参数
	ImageInfos []*ImageInfo `json:"ImageInfos,omitempty" name:"ImageInfos"`
}

func (r *CreateOrderAndDownloadsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderAndDownloadsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrderAndDownloadsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrderAndDownloadsResponseParams struct {
	// 成功核销后可以获取图片基本信息和原图地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadInfos []*DownloadInfo `json:"DownloadInfos,omitempty" name:"DownloadInfos"`

	// 可下载图片数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrderAndDownloadsResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrderAndDownloadsResponseParams `json:"Response"`
}

func (r *CreateOrderAndDownloadsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderAndDownloadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrderAndPayRequestParams struct {
	// 图片ID
	ImageId *uint64 `json:"ImageId,omitempty" name:"ImageId"`

	// 授权人ID
	AuthUserId *string `json:"AuthUserId,omitempty" name:"AuthUserId"`

	// 售卖组合id
	MarshalId *uint64 `json:"MarshalId,omitempty" name:"MarshalId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderAndPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	delete(f, "AuthUserId")
	delete(f, "MarshalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrderAndPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrderAndPayResponseParams struct {
	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrderAndPayResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrderAndPayResponseParams `json:"Response"`
}

func (r *CreateOrderAndPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderAndPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthUsersRequestParams struct {
	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthUsersResponseParams struct {
	// 授权人信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*AuthInfo `json:"Users,omitempty" name:"Users"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 是否是老策略用户
	OldUser *bool `json:"OldUser,omitempty" name:"OldUser"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuthUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthUsersResponseParams `json:"Response"`
}

func (r *DescribeAuthUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDownloadInfosRequestParams struct {
	// 默认10
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间晚于指定时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间早于指定时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 无效值，过滤结果为空
	ImageIds []*int64 `json:"ImageIds,omitempty" name:"ImageIds"`
}

type DescribeDownloadInfosRequest struct {
	*tchttp.BaseRequest
	
	// 默认10
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间晚于指定时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间早于指定时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 无效值，过滤结果为空
	ImageIds []*int64 `json:"ImageIds,omitempty" name:"ImageIds"`
}

func (r *DescribeDownloadInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "ImageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDownloadInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDownloadInfosResponseParams struct {
	// 核销下载记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadInfos []*DownloadInfo `json:"DownloadInfos,omitempty" name:"DownloadInfos"`

	// 总记录数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDownloadInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDownloadInfosResponseParams `json:"Response"`
}

func (r *DescribeDownloadInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageRequestParams struct {
	// 图片ID
	ImageId *uint64 `json:"ImageId,omitempty" name:"ImageId"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageResponseParams struct {
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
	Marshals []*ImageMarshal `json:"Marshals,omitempty" name:"Marshals"`

	// 宽
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 图片格式 jpg/eps/psd/...
	ImageFormat *string `json:"ImageFormat,omitempty" name:"ImageFormat"`

	// 图片类型 摄影图片、插画、漫画、图表、矢量、psd、全景、gif、模板
	ImageSenseType *string `json:"ImageSenseType,omitempty" name:"ImageSenseType"`

	// 关键词，多关键词用空格分隔
	Keywords *string `json:"Keywords,omitempty" name:"Keywords"`

	// 分层图库id
	LayeredGalleryId *int64 `json:"LayeredGalleryId,omitempty" name:"LayeredGalleryId"`

	// 构图方式：horizontal:横图、vertical:竖图、square:方图
	Orientation *string `json:"Orientation,omitempty" name:"Orientation"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageResponseParams `json:"Response"`
}

func (r *DescribeImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
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
	LayeredGalleryIds []*int64 `json:"LayeredGalleryIds,omitempty" name:"LayeredGalleryIds"`
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
	LayeredGalleryIds []*int64 `json:"LayeredGalleryIds,omitempty" name:"LayeredGalleryIds"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Keyword")
	delete(f, "Orientation")
	delete(f, "ImageSenseType")
	delete(f, "LayeredGalleryIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// 页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 总条数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 是否有下一页
	HaveMore *bool `json:"HaveMore,omitempty" name:"HaveMore"`

	// 图片信息数组
	Items []*ImageItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadInfo struct {
	// 图片基础信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 图片原图URL
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片缩略图URL
	ImageThumbUrl *string `json:"ImageThumbUrl,omitempty" name:"ImageThumbUrl"`

	// 订单Id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 订单创建时间
	OrderCreateTime *string `json:"OrderCreateTime,omitempty" name:"OrderCreateTime"`

	// 下载Id
	DownloadId *string `json:"DownloadId,omitempty" name:"DownloadId"`

	// 下载时间
	DownloadTime *string `json:"DownloadTime,omitempty" name:"DownloadTime"`

	// 图片购买类型，单张/会员
	ConsumeType *int64 `json:"ConsumeType,omitempty" name:"ConsumeType"`

	// 是否首次下载
	FirstDownload *bool `json:"FirstDownload,omitempty" name:"FirstDownload"`
}

type ImageInfo struct {
	// 图片Id
	ImageId *int64 `json:"ImageId,omitempty" name:"ImageId"`

	// 授权场景Id
	LicenseScopeId *int64 `json:"LicenseScopeId,omitempty" name:"LicenseScopeId"`

	// 尺寸名称Id
	DimensionsNameId *int64 `json:"DimensionsNameId,omitempty" name:"DimensionsNameId"`

	// 平台用户标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 平台用户下载图片购买的价格(单位:分)
	DownloadPrice *uint64 `json:"DownloadPrice,omitempty" name:"DownloadPrice"`

	// 下载类型。匹配集合中的任意元素：
	// <li>Single: 单张购买下载</li>
	// <li>BasicEnterpriseMember: 企业基础会员下载</li>
	// <li>AdvancedEnterpriseMember: 企业高级会员下载</li>
	// <li>DistinguishedEnterpriseMember: 企业尊享会员下载</li>
	DownloadType *string `json:"DownloadType,omitempty" name:"DownloadType"`
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

	// 宽
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高
	Height *int64 `json:"Height,omitempty" name:"Height"`
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

	// 授权范围id
	LicenseScopeId *int64 `json:"LicenseScopeId,omitempty" name:"LicenseScopeId"`

	// 尺寸
	DimensionsName *string `json:"DimensionsName,omitempty" name:"DimensionsName"`

	// 尺寸id
	DimensionsNameId *int64 `json:"DimensionsNameId,omitempty" name:"DimensionsNameId"`
}