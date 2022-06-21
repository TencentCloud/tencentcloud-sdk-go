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

package v20190321

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CodeDetail struct {
	// 二维码在图片中的位置，由边界点的坐标表示
	CodePosition []*CodePosition `json:"CodePosition,omitempty" name:"CodePosition"`

	// 二维码文本的编码格式
	CodeCharset *string `json:"CodeCharset,omitempty" name:"CodeCharset"`

	// 二维码的文本内容
	CodeText *string `json:"CodeText,omitempty" name:"CodeText"`

	// 二维码的类型：1:ONED_BARCODE，2:QRCOD，3:WXCODE，4:PDF417，5:DATAMATRIX
	CodeType *int64 `json:"CodeType,omitempty" name:"CodeType"`
}

type CodeDetect struct {
	// 从图片中检测到的二维码，可能为多个
	ModerationDetail []*CodeDetail `json:"ModerationDetail,omitempty" name:"ModerationDetail"`

	// 检测是否成功，0：成功，-1：出错
	ModerationCode *int64 `json:"ModerationCode,omitempty" name:"ModerationCode"`
}

type CodePosition struct {
	// 二维码边界点X轴坐标
	FloatX *float64 `json:"FloatX,omitempty" name:"FloatX"`

	// 二维码边界点Y轴坐标
	FloatY *float64 `json:"FloatY,omitempty" name:"FloatY"`
}

type Coordinate struct {
	// 左上角横坐标
	Cx *int64 `json:"Cx,omitempty" name:"Cx"`

	// 左上角纵坐标
	Cy *int64 `json:"Cy,omitempty" name:"Cy"`

	// 高度
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

// Predefined struct for user
type CreateFileSampleRequestParams struct {
	// 文件类型结构数组
	Contents []*FileSample `json:"Contents,omitempty" name:"Contents"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 20105：广告引流
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// image：图片
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 样本类型
	// 1：黑库
	// 2：白库
	Label *uint64 `json:"Label,omitempty" name:"Label"`
}

type CreateFileSampleRequest struct {
	*tchttp.BaseRequest
	
	// 文件类型结构数组
	Contents []*FileSample `json:"Contents,omitempty" name:"Contents"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 20105：广告引流
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// image：图片
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 样本类型
	// 1：黑库
	// 2：白库
	Label *uint64 `json:"Label,omitempty" name:"Label"`
}

func (r *CreateFileSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Contents")
	delete(f, "EvilType")
	delete(f, "FileType")
	delete(f, "Label")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSampleResponseParams struct {
	// 任务状态
	// 1：已完成
	// 2：处理中
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFileSampleResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileSampleResponseParams `json:"Response"`
}

func (r *CreateFileSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTextSampleRequestParams struct {
	// 关键词数组
	Contents []*string `json:"Contents,omitempty" name:"Contents"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 20105：广告引流
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 样本类型
	// 1：黑库
	// 2：白库
	Label *uint64 `json:"Label,omitempty" name:"Label"`

	// 测试修改参数
	Test *string `json:"Test,omitempty" name:"Test"`
}

type CreateTextSampleRequest struct {
	*tchttp.BaseRequest
	
	// 关键词数组
	Contents []*string `json:"Contents,omitempty" name:"Contents"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	// 20105：广告引流
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 样本类型
	// 1：黑库
	// 2：白库
	Label *uint64 `json:"Label,omitempty" name:"Label"`

	// 测试修改参数
	Test *string `json:"Test,omitempty" name:"Test"`
}

func (r *CreateTextSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTextSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Contents")
	delete(f, "EvilType")
	delete(f, "Label")
	delete(f, "Test")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTextSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTextSampleResponseParams struct {
	// 操作样本失败时返回的错误信息示例：  "样本1":错误码，"样本2":错误码
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 任务状态
	// 1：已完成
	// 2：处理中
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTextSampleResponse struct {
	*tchttp.BaseResponse
	Response *CreateTextSampleResponseParams `json:"Response"`
}

func (r *CreateTextSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTextSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomResult struct {
	// 命中的自定义关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 自定义库id
	LibId *string `json:"LibId,omitempty" name:"LibId"`

	// 自定义词库名称
	LibName *string `json:"LibName,omitempty" name:"LibName"`

	// 命中的自定义关键词的类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type DeleteFileSampleRequestParams struct {
	// 唯一标识数组
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

type DeleteFileSampleRequest struct {
	*tchttp.BaseRequest
	
	// 唯一标识数组
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteFileSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileSampleResponseParams struct {
	// 任务状态
	// 1：已完成
	// 2：处理中
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFileSampleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileSampleResponseParams `json:"Response"`
}

func (r *DeleteFileSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTextSampleRequestParams struct {
	// 唯一标识数组，目前暂时只支持单个删除
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

type DeleteTextSampleRequest struct {
	*tchttp.BaseRequest
	
	// 唯一标识数组，目前暂时只支持单个删除
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteTextSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTextSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTextSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTextSampleResponseParams struct {
	// 任务状态
	// 1：已完成
	// 2：处理中
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTextSampleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTextSampleResponseParams `json:"Response"`
}

func (r *DeleteTextSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTextSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSampleRequestParams struct {
	// 支持通过标签值进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 数量限制，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 升序（asc）还是降序（desc），默认：desc
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 按某个字段排序，目前仅支持CreatedAt排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeFileSampleRequest struct {
	*tchttp.BaseRequest
	
	// 支持通过标签值进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 数量限制，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 升序（asc）还是降序（desc），默认：desc
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 按某个字段排序，目前仅支持CreatedAt排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeFileSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderDirection")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSampleResponseParams struct {
	// 符合要求的样本的信息
	FileSampleSet []*FileSampleInfo `json:"FileSampleSet,omitempty" name:"FileSampleSet"`

	// 符合要求的样本的数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileSampleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileSampleResponseParams `json:"Response"`
}

func (r *DescribeFileSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextSampleRequestParams struct {
	// 支持通过标签值进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 数量限制，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 升序（asc）还是降序（desc），默认：desc
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 按某个字段排序，目前仅支持CreatedAt排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeTextSampleRequest struct {
	*tchttp.BaseRequest
	
	// 支持通过标签值进行筛选
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 数量限制，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 升序（asc）还是降序（desc），默认：desc
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 按某个字段排序，目前仅支持CreatedAt排序
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeTextSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextSampleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderDirection")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTextSampleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextSampleResponseParams struct {
	// 符合要求的样本的信息
	TextSampleSet []*TextSample `json:"TextSampleSet,omitempty" name:"TextSampleSet"`

	// 符合要求的样本的数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTextSampleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTextSampleResponseParams `json:"Response"`
}

func (r *DescribeTextSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextSampleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailResult struct {
	// 恶意标签，Normal：正常，Polity：涉政，Porn：色情，Illegal：违法，Abuse：谩骂，Terror：暴恐，Ad：广告，Custom：自定义关键词
	EvilLabel *string `json:"EvilLabel,omitempty" name:"EvilLabel"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂
	// 20105：广告引流 
	// 24001：暴恐
	EvilType *uint64 `json:"EvilType,omitempty" name:"EvilType"`

	// 该标签下命中的关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 该标签模型命中的分值
	Score *uint64 `json:"Score,omitempty" name:"Score"`
}

type Device struct {
	// 设备指纹ID
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// IOS设备，Identifier For Advertising（广告标识符）
	IDFA *string `json:"IDFA,omitempty" name:"IDFA"`

	// IOS设备，IDFV - Identifier For Vendor（应用开发商标识符）
	IDFV *string `json:"IDFV,omitempty" name:"IDFV"`

	// 设备序列号
	IMEI *string `json:"IMEI,omitempty" name:"IMEI"`

	// 用户IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// Mac地址
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// 设备指纹Token
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`
}

type FileSample struct {
	// 文件md5
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件url
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件压缩后云url
	CompressFileUrl *string `json:"CompressFileUrl,omitempty" name:"CompressFileUrl"`
}

type FileSampleInfo struct {
	// 处理错误码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 创建时间戳
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	EvilType *uint64 `json:"EvilType,omitempty" name:"EvilType"`

	// 文件的md5
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 唯一标识
	Id *string `json:"Id,omitempty" name:"Id"`

	// 样本类型
	// 1：黑库
	// 2：白库
	Label *uint64 `json:"Label,omitempty" name:"Label"`

	// 任务状态
	// 1：添加完成
	// 2：添加处理中
	// 3：下载中
	// 4：下载完成
	// 5：上传完成
	// 6：步骤完成
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 文件压缩后云url
	CompressFileUrl *string `json:"CompressFileUrl,omitempty" name:"CompressFileUrl"`

	// 文件的url
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

type Filter struct {
	// 需要过滤的字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 需要过滤字段的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ImageData struct {
	// 是否恶意 0：正常 1：可疑
	EvilFlag *int64 `json:"EvilFlag,omitempty" name:"EvilFlag"`

	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 20103：性感
	// 24001：暴恐
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 图片二维码详情
	CodeDetect *CodeDetect `json:"CodeDetect,omitempty" name:"CodeDetect"`

	// 图片性感详情
	HotDetect *ImageHotDetect `json:"HotDetect,omitempty" name:"HotDetect"`

	// 图片违法详情
	IllegalDetect *ImageIllegalDetect `json:"IllegalDetect,omitempty" name:"IllegalDetect"`

	// logo详情
	LogoDetect *LogoDetail `json:"LogoDetect,omitempty" name:"LogoDetect"`

	// 图片OCR详情
	OCRDetect *OCRDetect `json:"OCRDetect,omitempty" name:"OCRDetect"`

	// 手机检测详情
	PhoneDetect *PhoneDetect `json:"PhoneDetect,omitempty" name:"PhoneDetect"`

	// 图片涉政详情
	PolityDetect *ImagePolityDetect `json:"PolityDetect,omitempty" name:"PolityDetect"`

	// 图片涉黄详情
	PornDetect *ImagePornDetect `json:"PornDetect,omitempty" name:"PornDetect"`

	// 图片相似度详情
	Similar *Similar `json:"Similar,omitempty" name:"Similar"`

	// 图片暴恐详情
	TerrorDetect *ImageTerrorDetect `json:"TerrorDetect,omitempty" name:"TerrorDetect"`
}

type ImageHotDetect struct {
	// 恶意类型
	// 100：正常
	// 20103：性感
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 性感标签：性感特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 性感分：分值范围 0-100，分数越高性感倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImageIllegalDetect struct {
	// 恶意类型
	// 100：正常 
	// 20006：涉毒违法
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 违法标签：返回违法特征中文描述，如赌桌，枪支
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 违法分：分值范围 0-100，分数越高违法倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

// Predefined struct for user
type ImageModerationRequestParams struct {
	// 文件内容 Base64,与FileUrl必须二填一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件MD5值
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 文件地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest
	
	// 文件内容 Base64,与FileUrl必须二填一
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件MD5值
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 文件地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

func (r *ImageModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileContent")
	delete(f, "FileMD5")
	delete(f, "FileUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageModerationResponseParams struct {
	// 识别结果
	Data *ImageData `json:"Data,omitempty" name:"Data"`

	// 业务返回码
	BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImageModerationResponse struct {
	*tchttp.BaseResponse
	Response *ImageModerationResponseParams `json:"Response"`
}

func (r *ImageModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImagePolityDetect struct {
	// 恶意类型
	// 100：正常 
	// 20001：政治
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定  0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 命中的logo标签信息
	PolityLogoDetail []*Logo `json:"PolityLogoDetail,omitempty" name:"PolityLogoDetail"`

	// 命中的人脸名称
	FaceNames []*string `json:"FaceNames,omitempty" name:"FaceNames"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 命中的政治物品名称
	PolityItems []*string `json:"PolityItems,omitempty" name:"PolityItems"`

	// 政治（人脸）分：分值范围 0-100，分数越高可疑程度越高
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImagePornDetect struct {
	// 恶意类型
	// 100：正常
	// 20002：色情
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 色情标签：色情特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 色情分：分值范围 0-100，分数越高色情倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type ImageTerrorDetect struct {
	// 恶意类型
	// 100：正常
	// 24001：暴恐
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 关键词明细
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 暴恐标签：返回暴恐特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 暴恐分：分值范围0--100，分数越高暴恐倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type Logo struct {
	// logo图标坐标信息
	RrectF *RrectF `json:"RrectF,omitempty" name:"RrectF"`

	// logo图标置信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// logo图标名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type LogoDetail struct {
	// 命中的Applogo详情
	AppLogoDetail []*Logo `json:"AppLogoDetail,omitempty" name:"AppLogoDetail"`
}

type ManualReviewContent struct {
	// 审核批次号
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 审核内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 消息Id
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`

	// 审核内容类型 1 图片 2 视频 3 文本 4 音频
	ContentType *int64 `json:"ContentType,omitempty" name:"ContentType"`

	// 用户信息
	UserInfo *User `json:"UserInfo,omitempty" name:"UserInfo"`

	// 机器审核类型，与腾讯机器审核定义一致
	// 100 正常
	// 20001 政治
	// 20002 色情
	// 20006 违法
	// 20007 谩骂
	// 24001 暴恐
	// 20105 广告
	// 20103 性感
	AutoDetailCode *int64 `json:"AutoDetailCode,omitempty" name:"AutoDetailCode"`

	// 机器审核结果 0 放过 1 拦截
	AutoResult *int64 `json:"AutoResult,omitempty" name:"AutoResult"`

	// 回调信息标识，回传数据时原样返回
	CallBackInfo *string `json:"CallBackInfo,omitempty" name:"CallBackInfo"`

	// 创建时间 格式“2020-01-01 00:00:12”
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 审核优先级，可选值 [1,2,3,4]，其中 1 最高，4 最低
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 标题
	Title *string `json:"Title,omitempty" name:"Title"`
}

type ManualReviewData struct {
	// 人审内容批次号
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 人审内容ID
	ContentId *string `json:"ContentId,omitempty" name:"ContentId"`
}

// Predefined struct for user
type ManualReviewRequestParams struct {
	// 人工审核信息
	ReviewContent *ManualReviewContent `json:"ReviewContent,omitempty" name:"ReviewContent"`
}

type ManualReviewRequest struct {
	*tchttp.BaseRequest
	
	// 人工审核信息
	ReviewContent *ManualReviewContent `json:"ReviewContent,omitempty" name:"ReviewContent"`
}

func (r *ManualReviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualReviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReviewContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManualReviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManualReviewResponseParams struct {
	// 人审接口同步响应结果
	Data *ManualReviewData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManualReviewResponse struct {
	*tchttp.BaseResponse
	Response *ManualReviewResponseParams `json:"Response"`
}

func (r *ManualReviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualReviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OCRDetect struct {
	// 识别到的详细信息
	Item []*OCRItem `json:"Item,omitempty" name:"Item"`

	// 识别到的文本信息
	TextInfo *string `json:"TextInfo,omitempty" name:"TextInfo"`
}

type OCRItem struct {
	// 检测到的文本坐标信息
	TextPosition *Coordinate `json:"TextPosition,omitempty" name:"TextPosition"`

	// 文本命中具体标签
	EvilLabel *string `json:"EvilLabel,omitempty" name:"EvilLabel"`

	// 文本命中恶意违规类型
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 文本命中违规的关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 文本涉嫌违规分值
	Rate *int64 `json:"Rate,omitempty" name:"Rate"`

	// 检测到的文本信息
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`
}

type PhoneDetect struct {
	// 恶意类型
	// 100：正常
	// 21000：综合
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：正常 1：可疑
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 特征中文描述
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 分值范围 0-100，分数越高倾向越明显
	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type RiskDetails struct {
	// 预留字段，暂时不使用
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 风险类别，RiskAccount，RiskIP, RiskIMEI
	Label *string `json:"Label,omitempty" name:"Label"`

	// 预留字段，暂时不用
	Lable *string `json:"Lable,omitempty" name:"Lable"`

	// 风险等级，1:疑似，2：恶意
	Level *int64 `json:"Level,omitempty" name:"Level"`
}

type RrectF struct {
	// logo横坐标
	Cx *float64 `json:"Cx,omitempty" name:"Cx"`

	// logo纵坐标
	Cy *float64 `json:"Cy,omitempty" name:"Cy"`

	// logo图标高度
	Height *float64 `json:"Height,omitempty" name:"Height"`

	// logo图标中心旋转度
	Rotate *float64 `json:"Rotate,omitempty" name:"Rotate"`

	// logo图标宽度
	Width *float64 `json:"Width,omitempty" name:"Width"`
}

type Similar struct {
	// 恶意类型
	// 100：正常 
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 24001：暴恐
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 处置判定 0：未匹配到 1：恶意 2：白样本
	HitFlag *int64 `json:"HitFlag,omitempty" name:"HitFlag"`

	// 返回的种子url
	SeedUrl *string `json:"SeedUrl,omitempty" name:"SeedUrl"`
}

type TextData struct {
	// 是否恶意 0：正常 1：可疑
	EvilFlag *int64 `json:"EvilFlag,omitempty" name:"EvilFlag"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂
	// 20105：广告引流 
	// 24001：暴恐
	EvilType *int64 `json:"EvilType,omitempty" name:"EvilType"`

	// 消息类公共相关参数
	Common *TextOutputComm `json:"Common,omitempty" name:"Common"`

	// 返回的自定义词库结果
	CustomResult []*CustomResult `json:"CustomResult,omitempty" name:"CustomResult"`

	// 返回的详细结果
	DetailResult []*DetailResult `json:"DetailResult,omitempty" name:"DetailResult"`

	// 消息类ID信息
	ID *TextOutputID `json:"ID,omitempty" name:"ID"`

	// 消息类输出结果
	Res *TextOutputRes `json:"Res,omitempty" name:"Res"`

	// 账号风险检测结果
	RiskDetails []*RiskDetails `json:"RiskDetails,omitempty" name:"RiskDetails"`

	// 最终使用的BizType
	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`

	// 和请求中的DataId一致，原样返回
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 恶意标签，Normal：正常，Polity：涉政，Porn：色情，Illegal：违法，Abuse：谩骂，Terror：暴恐，Ad：广告，Custom：自定义关键词
	EvilLabel *string `json:"EvilLabel,omitempty" name:"EvilLabel"`

	// 输出的其他信息，不同客户内容不同
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 命中的关键词
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`

	// 命中的模型分值
	Score *uint64 `json:"Score,omitempty" name:"Score"`

	// 建议值,Block：打击,Review：待复审,Normal：正常
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
}

// Predefined struct for user
type TextModerationRequestParams struct {
	// 文本内容Base64编码。原文长度需小于15000字节，即5000个汉字以内。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 设备相关信息
	Device *Device `json:"Device,omitempty" name:"Device"`

	// 用户相关信息
	User *User `json:"User,omitempty" name:"User"`

	// 该字段用于标识业务场景。您可以在内容安全控制台创建对应的ID，配置不同的内容审核策略，通过接口调用，默认不填为0，后端使用默认策略
	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`

	// 数据ID，英文字母、下划线、-组成，不超过64个字符
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 业务应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

type TextModerationRequest struct {
	*tchttp.BaseRequest
	
	// 文本内容Base64编码。原文长度需小于15000字节，即5000个汉字以内。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 设备相关信息
	Device *Device `json:"Device,omitempty" name:"Device"`

	// 用户相关信息
	User *User `json:"User,omitempty" name:"User"`

	// 该字段用于标识业务场景。您可以在内容安全控制台创建对应的ID，配置不同的内容审核策略，通过接口调用，默认不填为0，后端使用默认策略
	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`

	// 数据ID，英文字母、下划线、-组成，不超过64个字符
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 业务应用ID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`
}

func (r *TextModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "Device")
	delete(f, "User")
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextModerationResponseParams struct {
	// 识别结果
	Data *TextData `json:"Data,omitempty" name:"Data"`

	// 业务返回码
	BusinessCode *int64 `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextModerationResponse struct {
	*tchttp.BaseResponse
	Response *TextModerationResponseParams `json:"Response"`
}

func (r *TextModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextModerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextOutputComm struct {
	// 接入业务的唯一ID
	AppID *int64 `json:"AppID,omitempty" name:"AppID"`

	// 接口唯一ID，旁路调用接口返回有该字段，标识唯一接口
	BUCtrlID *int64 `json:"BUCtrlID,omitempty" name:"BUCtrlID"`

	// 消息发送时间
	SendTime *int64 `json:"SendTime,omitempty" name:"SendTime"`

	// 请求字段里的Common.Uin
	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
}

type TextOutputID struct {
	// 接入业务的唯一ID
	MsgID *string `json:"MsgID,omitempty" name:"MsgID"`

	// 用户账号uin，对应请求协议里的Content.User.Uin。旁路结果有回带，串联结果无该字段
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type TextOutputRes struct {
	// 操作人,信安处理人企业微信ID
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 恶意操作码，
	// 删除（1）， 通过（2）， 先审后发（100012）
	ResultCode *int64 `json:"ResultCode,omitempty" name:"ResultCode"`

	// 操作结果备注说明
	ResultMsg *string `json:"ResultMsg,omitempty" name:"ResultMsg"`

	// 恶意类型，广告（10001）， 政治（20001）， 色情（20002）， 社会事件（20004）， 暴力（20011）， 低俗（20012）， 违法犯罪（20006）， 欺诈（20008）， 版权（20013）， 谣言（20104）， 其他（21000）
	ResultType *int64 `json:"ResultType,omitempty" name:"ResultType"`
}

type TextSample struct {
	// 处理错误码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 关键词
	Content *string `json:"Content,omitempty" name:"Content"`

	// 创建时间戳
	CreatedAt *uint64 `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 恶意类型
	// 100：正常
	// 20001：政治
	// 20002：色情 
	// 20006：涉毒违法
	// 20007：谩骂 
	// 20105：广告引流 
	// 24001：暴恐
	EvilType *uint64 `json:"EvilType,omitempty" name:"EvilType"`

	// 唯一标识
	Id *string `json:"Id,omitempty" name:"Id"`

	// 样本类型
	// 1：黑库
	// 2：白库
	Label *uint64 `json:"Label,omitempty" name:"Label"`

	// 任务状态
	// 1：已完成
	// 2：处理中
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type User struct {
	// 账号类别，"1-微信uin 2-QQ号 3-微信群uin 4-qq群号 5-微信openid 6-QQopenid 7-其它string"
	AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

	// 年龄 默认0 未知
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 性别 默认0 未知 1 男性 2 女性
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 用户等级，默认0 未知 1 低 2 中 3 高
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户账号ID，如填写，会根据账号历史恶意情况，判定消息有害结果，特别是有利于可疑恶意情况下的辅助判断。账号可以填写微信uin、QQ号、微信openid、QQopenid、字符串等。该字段和账号类别确定唯一账号。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}