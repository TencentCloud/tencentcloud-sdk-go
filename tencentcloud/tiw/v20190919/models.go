// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20190919

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AuthParam struct {
	// 应用SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户ID对应的签名
	UserSig *string `json:"UserSig,omitnil,omitempty" name:"UserSig"`
}

type Canvas struct {
	// 混流画布宽高配置
	LayoutParams *LayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 背景颜色，默认为黑色，格式为RGB格式，如红色为"#FF0000"
	BackgroundColor *string `json:"BackgroundColor,omitnil,omitempty" name:"BackgroundColor"`
}

type Concat struct {
	// 是否开启拼接功能
	// 在开启了视频拼接功能的情况下，实时录制服务会把同一个用户因为暂停导致的多段视频拼接成一个视频
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 视频拼接时使用的垫片图片下载地址，不填默认用全黑的图片进行视频垫片
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`
}

// Predefined struct for user
type CreatePPTCheckTaskRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 经过URL编码后的PPT文件地址。URL 编码会将字符转换为可通过因特网传输的格式，例如文档地址为http://example.com/测试.pptx，经过URL编码之后为http://example.com/%E6%B5%8B%E8%AF%95.pptx。为了提高URL解析的成功率，请对URL进行编码。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 是否对不支持元素开启自动处理的功能，默认不开启。
	// true -- 开启
	// false -- 不开启
	// 
	// 当设置为`true`时，可配合`AutoHandleUnsupportedElementTypes`参数使用，具体有哪些不兼容元素类型，可参考`AutoHandleUnsupportedElementTypes`参数的说明。
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil,omitempty" name:"AutoHandleUnsupportedElement"`

	// 此参数仅在`AutoHandleUnsupportedElement`参数为`true`的情况下有效。
	// 
	// 指定需要自动处理的不兼容元素类型，默认对所有不兼容的元素进行自动处理。
	// 
	// 目前支持检测的不兼容元素类型及对应的自动处理方式如下：
	// 0: 不支持的墨迹类型
	// -- 自动处理方式：移除墨迹
	// 
	// 1: 自动翻页
	// -- 自动处理方式：移除自动翻页设置，并修改为单击切换
	// 
	// 2: 已损坏音视频
	// -- 自动处理方式：移除对损坏音视频的引用
	// 
	// 3: 不可访问资源
	// -- 自动处理方式：移除对不可访问的资源的引用
	// 
	// 4: 只读文件
	// -- 自动处理方式：移除只读设置
	// 
	// 5: 不支持的元素编辑锁定状态
	// -- 自动处理方式：移除锁定状态
	// 
	// 6: 可能有兼容问题的字体
	// -- 自动处理方式： 不支持处理
	// 
	// 7: 设置了柔化边缘的GIF图片
	// -- 自动处理方式：移除柔化边缘设置
	// 
	// 8: 存在不兼容的空格下划线
	// -- 自动处理方式：通过调整空格下划线前后文本的字体语言体系，保证空格下划线表现正常
	// 
	// 9: 存在设置了分段动画的数学公式和文本混合内容
	// -- 自动处理方式： 不支持处理
	// 
	// 10: 存在设置了分段动画的渐变色文本
	// -- 自动处理方式： 不支持处理
	// 
	// 11: 存在不兼容的分散对齐方式
	// -- 自动处理方式： 不支持处理
	// 
	// 12: 存在不兼容的多倍行距设置
	// -- 自动处理方式： 不支持处理
	// 
	// 13: 存在带有特殊符号内容的datetime类型的a:fld标签元素
	// -- 自动处理方式： a:fld标签替换为普通文本
	AutoHandleUnsupportedElementTypes []*int64 `json:"AutoHandleUnsupportedElementTypes,omitnil,omitempty" name:"AutoHandleUnsupportedElementTypes"`
}

type CreatePPTCheckTaskRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 经过URL编码后的PPT文件地址。URL 编码会将字符转换为可通过因特网传输的格式，例如文档地址为http://example.com/测试.pptx，经过URL编码之后为http://example.com/%E6%B5%8B%E8%AF%95.pptx。为了提高URL解析的成功率，请对URL进行编码。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 是否对不支持元素开启自动处理的功能，默认不开启。
	// true -- 开启
	// false -- 不开启
	// 
	// 当设置为`true`时，可配合`AutoHandleUnsupportedElementTypes`参数使用，具体有哪些不兼容元素类型，可参考`AutoHandleUnsupportedElementTypes`参数的说明。
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil,omitempty" name:"AutoHandleUnsupportedElement"`

	// 此参数仅在`AutoHandleUnsupportedElement`参数为`true`的情况下有效。
	// 
	// 指定需要自动处理的不兼容元素类型，默认对所有不兼容的元素进行自动处理。
	// 
	// 目前支持检测的不兼容元素类型及对应的自动处理方式如下：
	// 0: 不支持的墨迹类型
	// -- 自动处理方式：移除墨迹
	// 
	// 1: 自动翻页
	// -- 自动处理方式：移除自动翻页设置，并修改为单击切换
	// 
	// 2: 已损坏音视频
	// -- 自动处理方式：移除对损坏音视频的引用
	// 
	// 3: 不可访问资源
	// -- 自动处理方式：移除对不可访问的资源的引用
	// 
	// 4: 只读文件
	// -- 自动处理方式：移除只读设置
	// 
	// 5: 不支持的元素编辑锁定状态
	// -- 自动处理方式：移除锁定状态
	// 
	// 6: 可能有兼容问题的字体
	// -- 自动处理方式： 不支持处理
	// 
	// 7: 设置了柔化边缘的GIF图片
	// -- 自动处理方式：移除柔化边缘设置
	// 
	// 8: 存在不兼容的空格下划线
	// -- 自动处理方式：通过调整空格下划线前后文本的字体语言体系，保证空格下划线表现正常
	// 
	// 9: 存在设置了分段动画的数学公式和文本混合内容
	// -- 自动处理方式： 不支持处理
	// 
	// 10: 存在设置了分段动画的渐变色文本
	// -- 自动处理方式： 不支持处理
	// 
	// 11: 存在不兼容的分散对齐方式
	// -- 自动处理方式： 不支持处理
	// 
	// 12: 存在不兼容的多倍行距设置
	// -- 自动处理方式： 不支持处理
	// 
	// 13: 存在带有特殊符号内容的datetime类型的a:fld标签元素
	// -- 自动处理方式： a:fld标签替换为普通文本
	AutoHandleUnsupportedElementTypes []*int64 `json:"AutoHandleUnsupportedElementTypes,omitnil,omitempty" name:"AutoHandleUnsupportedElementTypes"`
}

func (r *CreatePPTCheckTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePPTCheckTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Url")
	delete(f, "AutoHandleUnsupportedElement")
	delete(f, "AutoHandleUnsupportedElementTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePPTCheckTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePPTCheckTaskResponseParams struct {
	// 检测任务的唯一标识Id，用于查询该任务的进度以及检测结果
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePPTCheckTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePPTCheckTaskResponseParams `json:"Response"`
}

func (r *CreatePPTCheckTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePPTCheckTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotTaskRequestParams struct {
	// 白板相关参数
	Whiteboard *SnapshotWhiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 白板房间 `SdkAppId`
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 白板房间号
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 白板板书生成结果通知回调地址
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 白板板书文件 `COS` 存储参数， 不填默认存储在公共存储桶，公共存储桶的数据仅保存3天
	COS *SnapshotCOS `json:"COS,omitnil,omitempty" name:"COS"`

	// 白板板书生成模式，默认为 `AllMarks`。取值说明如下：
	// 
	// `AllMarks` - 全量模式，即对于客户端每一次调用 `addSnapshotMark` 接口打上的白板板书生成标志全部都会生成对应的白板板书图片。
	// 
	// `LatestMarksOnly` - 单页去重模式，即对于客户端在同一页白板上多次调用 `addSnapshotMark` 打上的白板板书生成标志仅保留最新一次标志来生成对应白板页的白板板书图片。
	// 
	// （**注意：`LatestMarksOnly` 模式只有客户端使用v2.6.8及以上版本的白板SDK调用 `addSnapshotMark` 时才生效，否则即使在调用本API是指定了 `LatestMarksOnly` 模式，服务后台会使用默认的 `AllMarks` 模式生成白板板书**）
	SnapshotMode *string `json:"SnapshotMode,omitnil,omitempty" name:"SnapshotMode"`
}

type CreateSnapshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 白板相关参数
	Whiteboard *SnapshotWhiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 白板房间 `SdkAppId`
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 白板房间号
	RoomId *uint64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 白板板书生成结果通知回调地址
	CallbackURL *string `json:"CallbackURL,omitnil,omitempty" name:"CallbackURL"`

	// 白板板书文件 `COS` 存储参数， 不填默认存储在公共存储桶，公共存储桶的数据仅保存3天
	COS *SnapshotCOS `json:"COS,omitnil,omitempty" name:"COS"`

	// 白板板书生成模式，默认为 `AllMarks`。取值说明如下：
	// 
	// `AllMarks` - 全量模式，即对于客户端每一次调用 `addSnapshotMark` 接口打上的白板板书生成标志全部都会生成对应的白板板书图片。
	// 
	// `LatestMarksOnly` - 单页去重模式，即对于客户端在同一页白板上多次调用 `addSnapshotMark` 打上的白板板书生成标志仅保留最新一次标志来生成对应白板页的白板板书图片。
	// 
	// （**注意：`LatestMarksOnly` 模式只有客户端使用v2.6.8及以上版本的白板SDK调用 `addSnapshotMark` 时才生效，否则即使在调用本API是指定了 `LatestMarksOnly` 模式，服务后台会使用默认的 `AllMarks` 模式生成白板板书**）
	SnapshotMode *string `json:"SnapshotMode,omitnil,omitempty" name:"SnapshotMode"`
}

func (r *CreateSnapshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Whiteboard")
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "CallbackURL")
	delete(f, "COS")
	delete(f, "SnapshotMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotTaskResponseParams struct {
	// 白板板书生成任务ID，只有任务创建成功的时候才会返回此字段
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSnapshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotTaskResponseParams `json:"Response"`
}

func (r *CreateSnapshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 经过URL编码后的转码文件地址。URL 编码会将字符转换为可通过因特网传输的格式，比如文档地址为http://example.com/测试.pdf，经过URL编码之后为http://example.com/%E6%B5%8B%E8%AF%95.pdf。为了提高URL解析的成功率，请对URL进行编码。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 是否为静态PPT，默认为False；
	// 如果IsStaticPPT为False，后缀名为.ppt或.pptx的文档会动态转码成HTML5页面，其他格式的文档会静态转码成图片；如果IsStaticPPT为True，所有格式的文档会静态转码成图片；
	IsStaticPPT *bool `json:"IsStaticPPT,omitnil,omitempty" name:"IsStaticPPT"`

	// 注意: 该参数已废弃, 请使用最新的 [云API SDK](https://cloud.tencent.com/document/api/1137/40060#SDK) ，使用 MinScaleResolution字段传递分辨率
	// 
	// 转码后文档的最小分辨率，不传、传空字符串或分辨率格式错误则使用文档原分辨率
	// 
	// 示例：1280x720，注意分辨率宽高中间为英文字母"xyz"的"x"
	MinResolution *string `json:"MinResolution,omitnil,omitempty" name:"MinResolution"`

	// 动态PPT转码可以为文件生成该分辨率的缩略图，不传、传空字符串或分辨率格式错误则不生成缩略图，分辨率格式同MinResolution
	ThumbnailResolution *string `json:"ThumbnailResolution,omitnil,omitempty" name:"ThumbnailResolution"`

	// 转码文件压缩格式，不传、传空字符串或不是指定的格式则不生成压缩文件，目前支持如下压缩格式：
	// 
	// zip： 生成`.zip`压缩包
	// tar.gz： 生成`.tar.gz`压缩包
	CompressFileType *string `json:"CompressFileType,omitnil,omitempty" name:"CompressFileType"`

	// 内部参数
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// 文档转码优先级， 只有对于PPT动态转码生效，支持填入以下值：<br/>
	// - low: 低优先级转码，对于动态转码，能支持500MB（下载超时时间10分钟）以及2000页文档，但资源有限可能会有比较长时间的排队，请酌情使用该功能。<br/>
	// - 不填表示正常优先级转码，支持200MB文件（下载超时时间2分钟），500页以内的文档进行转码
	// <br/>
	// 注意：对于PDF等静态文件转码，无论是正常优先级或者低优先级，最大只能支持200MB
	Priority *string `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 转码后文档的最小分辨率，不传、传空字符串或分辨率格式错误则使用文档原分辨率。
	// 分辨率越高，效果越清晰，转出来的图片资源体积会越大，课件加载耗时会变长，请根据实际使用场景配置此参数。
	// 
	// 示例：1280x720，注意分辨率宽高中间为英文字母"xyz"的"x"
	MinScaleResolution *string `json:"MinScaleResolution,omitnil,omitempty" name:"MinScaleResolution"`

	// 此参数仅对动态转码生效。
	// 
	// 是否对不支持元素开启自动处理的功能，默认不开启。
	// true -- 开启
	// false -- 不开启
	// 
	// 当设置为`true`时，可配合`AutoHandleUnsupportedElementTypes`参数使用，具体有哪些不兼容元素类型，可参考`AutoHandleUnsupportedElementTypes`参数的说明。
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil,omitempty" name:"AutoHandleUnsupportedElement"`

	// 此参数仅在`AutoHandleUnsupportedElement`参数为`true`的情况下有效。
	// 
	// 指定需要自动处理的不兼容元素类型，默认对所有不兼容的元素进行自动处理。
	// 
	// 目前支持检测的不兼容元素类型及对应的自动处理方式如下：
	// 0: 不支持的墨迹类型
	// -- 自动处理方式：移除墨迹
	// 
	// 1: 自动翻页
	// -- 自动处理方式：移除自动翻页设置，并修改为单击切换
	// 
	// 2: 已损坏音视频
	// -- 自动处理方式：移除对损坏音视频的引用
	// 
	// 3: 不可访问资源
	// -- 自动处理方式：移除对不可访问的资源的引用
	// 
	// 4: 只读文件
	// -- 自动处理方式：移除只读设置
	// 
	// 5: 不支持的元素编辑锁定状态
	// -- 自动处理方式：移除锁定状态
	// 
	// 6: 可能有兼容问题的字体
	// -- 自动处理方式： 不支持处理
	// 
	// 7: 设置了柔化边缘的GIF图片
	// -- 自动处理方式：移除柔化边缘设置
	// 
	// 8: 存在不兼容的空格下划线
	// -- 自动处理方式：通过调整空格下划线前后文本的字体语言体系，保证空格下划线表现正常
	// 
	// 9: 存在设置了分段动画的数学公式和文本混合内容
	// -- 自动处理方式： 不支持处理
	// 
	// 10: 存在设置了分段动画的渐变色文本
	// -- 自动处理方式： 不支持处理
	// 
	// 11: 存在不兼容的分散对齐方式
	// -- 自动处理方式： 不支持处理
	// 
	// 12: 存在不兼容的多倍行距设置
	// -- 自动处理方式： 不支持处理
	// 
	// 13: 存在带有特殊符号内容的datetime类型的a:fld标签元素
	// -- 自动处理方式： a:fld标签替换为普通文本
	AutoHandleUnsupportedElementTypes []*int64 `json:"AutoHandleUnsupportedElementTypes,omitnil,omitempty" name:"AutoHandleUnsupportedElementTypes"`

	// Excel表格转码参数，可设置转码时表格纸张大小及纸张方向等参数（仅对转码文件为Excel表格文件的静态转码任务生效）
	ExcelParam *ExcelParam `json:"ExcelParam,omitnil,omitempty" name:"ExcelParam"`
}

type CreateTranscodeRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 经过URL编码后的转码文件地址。URL 编码会将字符转换为可通过因特网传输的格式，比如文档地址为http://example.com/测试.pdf，经过URL编码之后为http://example.com/%E6%B5%8B%E8%AF%95.pdf。为了提高URL解析的成功率，请对URL进行编码。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 是否为静态PPT，默认为False；
	// 如果IsStaticPPT为False，后缀名为.ppt或.pptx的文档会动态转码成HTML5页面，其他格式的文档会静态转码成图片；如果IsStaticPPT为True，所有格式的文档会静态转码成图片；
	IsStaticPPT *bool `json:"IsStaticPPT,omitnil,omitempty" name:"IsStaticPPT"`

	// 注意: 该参数已废弃, 请使用最新的 [云API SDK](https://cloud.tencent.com/document/api/1137/40060#SDK) ，使用 MinScaleResolution字段传递分辨率
	// 
	// 转码后文档的最小分辨率，不传、传空字符串或分辨率格式错误则使用文档原分辨率
	// 
	// 示例：1280x720，注意分辨率宽高中间为英文字母"xyz"的"x"
	MinResolution *string `json:"MinResolution,omitnil,omitempty" name:"MinResolution"`

	// 动态PPT转码可以为文件生成该分辨率的缩略图，不传、传空字符串或分辨率格式错误则不生成缩略图，分辨率格式同MinResolution
	ThumbnailResolution *string `json:"ThumbnailResolution,omitnil,omitempty" name:"ThumbnailResolution"`

	// 转码文件压缩格式，不传、传空字符串或不是指定的格式则不生成压缩文件，目前支持如下压缩格式：
	// 
	// zip： 生成`.zip`压缩包
	// tar.gz： 生成`.tar.gz`压缩包
	CompressFileType *string `json:"CompressFileType,omitnil,omitempty" name:"CompressFileType"`

	// 内部参数
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// 文档转码优先级， 只有对于PPT动态转码生效，支持填入以下值：<br/>
	// - low: 低优先级转码，对于动态转码，能支持500MB（下载超时时间10分钟）以及2000页文档，但资源有限可能会有比较长时间的排队，请酌情使用该功能。<br/>
	// - 不填表示正常优先级转码，支持200MB文件（下载超时时间2分钟），500页以内的文档进行转码
	// <br/>
	// 注意：对于PDF等静态文件转码，无论是正常优先级或者低优先级，最大只能支持200MB
	Priority *string `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 转码后文档的最小分辨率，不传、传空字符串或分辨率格式错误则使用文档原分辨率。
	// 分辨率越高，效果越清晰，转出来的图片资源体积会越大，课件加载耗时会变长，请根据实际使用场景配置此参数。
	// 
	// 示例：1280x720，注意分辨率宽高中间为英文字母"xyz"的"x"
	MinScaleResolution *string `json:"MinScaleResolution,omitnil,omitempty" name:"MinScaleResolution"`

	// 此参数仅对动态转码生效。
	// 
	// 是否对不支持元素开启自动处理的功能，默认不开启。
	// true -- 开启
	// false -- 不开启
	// 
	// 当设置为`true`时，可配合`AutoHandleUnsupportedElementTypes`参数使用，具体有哪些不兼容元素类型，可参考`AutoHandleUnsupportedElementTypes`参数的说明。
	AutoHandleUnsupportedElement *bool `json:"AutoHandleUnsupportedElement,omitnil,omitempty" name:"AutoHandleUnsupportedElement"`

	// 此参数仅在`AutoHandleUnsupportedElement`参数为`true`的情况下有效。
	// 
	// 指定需要自动处理的不兼容元素类型，默认对所有不兼容的元素进行自动处理。
	// 
	// 目前支持检测的不兼容元素类型及对应的自动处理方式如下：
	// 0: 不支持的墨迹类型
	// -- 自动处理方式：移除墨迹
	// 
	// 1: 自动翻页
	// -- 自动处理方式：移除自动翻页设置，并修改为单击切换
	// 
	// 2: 已损坏音视频
	// -- 自动处理方式：移除对损坏音视频的引用
	// 
	// 3: 不可访问资源
	// -- 自动处理方式：移除对不可访问的资源的引用
	// 
	// 4: 只读文件
	// -- 自动处理方式：移除只读设置
	// 
	// 5: 不支持的元素编辑锁定状态
	// -- 自动处理方式：移除锁定状态
	// 
	// 6: 可能有兼容问题的字体
	// -- 自动处理方式： 不支持处理
	// 
	// 7: 设置了柔化边缘的GIF图片
	// -- 自动处理方式：移除柔化边缘设置
	// 
	// 8: 存在不兼容的空格下划线
	// -- 自动处理方式：通过调整空格下划线前后文本的字体语言体系，保证空格下划线表现正常
	// 
	// 9: 存在设置了分段动画的数学公式和文本混合内容
	// -- 自动处理方式： 不支持处理
	// 
	// 10: 存在设置了分段动画的渐变色文本
	// -- 自动处理方式： 不支持处理
	// 
	// 11: 存在不兼容的分散对齐方式
	// -- 自动处理方式： 不支持处理
	// 
	// 12: 存在不兼容的多倍行距设置
	// -- 自动处理方式： 不支持处理
	// 
	// 13: 存在带有特殊符号内容的datetime类型的a:fld标签元素
	// -- 自动处理方式： a:fld标签替换为普通文本
	AutoHandleUnsupportedElementTypes []*int64 `json:"AutoHandleUnsupportedElementTypes,omitnil,omitempty" name:"AutoHandleUnsupportedElementTypes"`

	// Excel表格转码参数，可设置转码时表格纸张大小及纸张方向等参数（仅对转码文件为Excel表格文件的静态转码任务生效）
	ExcelParam *ExcelParam `json:"ExcelParam,omitnil,omitempty" name:"ExcelParam"`
}

func (r *CreateTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Url")
	delete(f, "IsStaticPPT")
	delete(f, "MinResolution")
	delete(f, "ThumbnailResolution")
	delete(f, "CompressFileType")
	delete(f, "ExtraData")
	delete(f, "Priority")
	delete(f, "MinScaleResolution")
	delete(f, "AutoHandleUnsupportedElement")
	delete(f, "AutoHandleUnsupportedElementTypes")
	delete(f, "ExcelParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTranscodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTranscodeResponseParams struct {
	// 文档转码任务的唯一标识Id，用于查询该任务的进度以及转码结果
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateTranscodeResponseParams `json:"Response"`
}

func (r *CreateTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoGenerationTaskRequestParams struct {
	// 实时录制任务的TaskId
	OnlineRecordTaskId *string `json:"OnlineRecordTaskId,omitnil,omitempty" name:"OnlineRecordTaskId"`

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 视频生成的白板参数，例如白板宽高等。
	// 
	// 此参数与开始录制接口提供的Whiteboard参数互斥，在本接口与开始录制接口都提供了Whiteboard参数时，优先使用本接口指定的Whiteboard参数进行视频生成，否则使用开始录制接口提供的Whiteboard参数进行视频生成。
	Whiteboard *Whiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 视频拼接参数
	// 
	// 此参数与开始录制接口提供的Concat参数互斥，在本接口与开始录制接口都提供了Concat参数时，优先使用本接口指定的Concat参数进行视频拼接，否则使用开始录制接口提供的Concat参数进行视频拼接。
	Concat *Concat `json:"Concat,omitnil,omitempty" name:"Concat"`

	// 视频生成混流参数
	// 
	// 此参数与开始录制接口提供的MixStream参数互斥，在本接口与开始录制接口都提供了MixStream参数时，优先使用本接口指定的MixStream参数进行视频混流，否则使用开始录制接口提供的MixStream参数进行视频拼混流。
	MixStream *MixStream `json:"MixStream,omitnil,omitempty" name:"MixStream"`

	// 视频生成控制参数，用于更精细地指定需要生成哪些流，某一路流是否禁用音频，是否只录制小画面等
	// 
	// 此参数与开始录制接口提供的RecordControl参数互斥，在本接口与开始录制接口都提供了RecordControl参数时，优先使用本接口指定的RecordControl参数进行视频生成控制，否则使用开始录制接口提供的RecordControl参数进行视频拼生成控制。
	RecordControl *RecordControl `json:"RecordControl,omitnil,omitempty" name:"RecordControl"`

	// 内部参数
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`
}

type CreateVideoGenerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实时录制任务的TaskId
	OnlineRecordTaskId *string `json:"OnlineRecordTaskId,omitnil,omitempty" name:"OnlineRecordTaskId"`

	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 视频生成的白板参数，例如白板宽高等。
	// 
	// 此参数与开始录制接口提供的Whiteboard参数互斥，在本接口与开始录制接口都提供了Whiteboard参数时，优先使用本接口指定的Whiteboard参数进行视频生成，否则使用开始录制接口提供的Whiteboard参数进行视频生成。
	Whiteboard *Whiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 视频拼接参数
	// 
	// 此参数与开始录制接口提供的Concat参数互斥，在本接口与开始录制接口都提供了Concat参数时，优先使用本接口指定的Concat参数进行视频拼接，否则使用开始录制接口提供的Concat参数进行视频拼接。
	Concat *Concat `json:"Concat,omitnil,omitempty" name:"Concat"`

	// 视频生成混流参数
	// 
	// 此参数与开始录制接口提供的MixStream参数互斥，在本接口与开始录制接口都提供了MixStream参数时，优先使用本接口指定的MixStream参数进行视频混流，否则使用开始录制接口提供的MixStream参数进行视频拼混流。
	MixStream *MixStream `json:"MixStream,omitnil,omitempty" name:"MixStream"`

	// 视频生成控制参数，用于更精细地指定需要生成哪些流，某一路流是否禁用音频，是否只录制小画面等
	// 
	// 此参数与开始录制接口提供的RecordControl参数互斥，在本接口与开始录制接口都提供了RecordControl参数时，优先使用本接口指定的RecordControl参数进行视频生成控制，否则使用开始录制接口提供的RecordControl参数进行视频拼生成控制。
	RecordControl *RecordControl `json:"RecordControl,omitnil,omitempty" name:"RecordControl"`

	// 内部参数
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`
}

func (r *CreateVideoGenerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoGenerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OnlineRecordTaskId")
	delete(f, "SdkAppId")
	delete(f, "Whiteboard")
	delete(f, "Concat")
	delete(f, "MixStream")
	delete(f, "RecordControl")
	delete(f, "ExtraData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoGenerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoGenerationTaskResponseParams struct {
	// 视频生成的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVideoGenerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoGenerationTaskResponseParams `json:"Response"`
}

func (r *CreateVideoGenerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoGenerationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomLayout struct {
	// 混流画布参数
	Canvas *Canvas `json:"Canvas,omitnil,omitempty" name:"Canvas"`

	// 流布局参数，每路流的布局不能超出画布区域
	InputStreamList []*StreamLayout `json:"InputStreamList,omitnil,omitempty" name:"InputStreamList"`
}

// Predefined struct for user
type DescribeOnlineRecordCallbackRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeOnlineRecordCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeOnlineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOnlineRecordCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOnlineRecordCallbackResponseParams struct {
	// 实时录制事件回调地址，如果未设置回调地址，该字段为空字符串
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 实时录制回调鉴权密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOnlineRecordCallbackResponseParams `json:"Response"`
}

func (r *DescribeOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOnlineRecordRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实时录制任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实时录制任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOnlineRecordResponseParams struct {
	// 录制结束原因，
	// - AUTO: 房间内长时间没有音视频上行及白板操作导致自动停止录制
	// - USER_CALL: 主动调用了停止录制接口
	// - EXCEPTION: 录制异常结束
	// - FORCE_STOP: 强制停止录制，一般是因为暂停超过90分钟或者录制总时长超过24小时。
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 需要查询结果的录制任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 录制任务状态
	// - PREPARED: 表示录制正在准备中（进房/启动录制服务等操作）
	// - RECORDING: 表示录制已开始
	// - PAUSED: 表示录制已暂停
	// - STOPPED: 表示录制已停止，正在处理并上传视频
	// - FINISHED: 表示视频处理并上传完成，成功生成录制结果
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 房间号
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 白板的群组 Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 录制用户Id
	RecordUserId *string `json:"RecordUserId,omitnil,omitempty" name:"RecordUserId"`

	// 实际开始录制时间，Unix 时间戳，单位秒
	RecordStartTime *int64 `json:"RecordStartTime,omitnil,omitempty" name:"RecordStartTime"`

	// 实际停止录制时间，Unix 时间戳，单位秒
	RecordStopTime *int64 `json:"RecordStopTime,omitnil,omitempty" name:"RecordStopTime"`

	// 回放视频总时长（单位：毫秒）
	TotalTime *int64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// 录制过程中出现异常的次数
	ExceptionCnt *int64 `json:"ExceptionCnt,omitnil,omitempty" name:"ExceptionCnt"`

	// 拼接视频中被忽略的时间段，只有开启视频拼接功能的时候，这个参数才是有效的
	OmittedDurations []*OmittedDuration `json:"OmittedDurations,omitnil,omitempty" name:"OmittedDurations"`

	// 录制视频列表
	VideoInfos []*VideoInfo `json:"VideoInfos,omitnil,omitempty" name:"VideoInfos"`

	// 回放URL，需配合信令播放器使用。此字段仅适用于`视频生成模式`
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplayUrl *string `json:"ReplayUrl,omitnil,omitempty" name:"ReplayUrl"`

	// 视频流在录制过程中断流次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interrupts []*Interrupt `json:"Interrupts,omitnil,omitempty" name:"Interrupts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOnlineRecordResponseParams `json:"Response"`
}

func (r *DescribeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePPTCheckCallbackRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribePPTCheckCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribePPTCheckCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePPTCheckCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePPTCheckCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePPTCheckCallbackResponseParams struct {
	// 回调地址
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 回调鉴权密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePPTCheckCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribePPTCheckCallbackResponseParams `json:"Response"`
}

func (r *DescribePPTCheckCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePPTCheckCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePPTCheckRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一标识Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribePPTCheckRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 任务的唯一标识Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribePPTCheckRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePPTCheckRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePPTCheckRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePPTCheckResponseParams struct {
	// 任务的唯一标识Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// PPT文件是否正常
	IsOK *bool `json:"IsOK,omitnil,omitempty" name:"IsOK"`

	// 修复后的PPT URL，只有创建任务时参数AutoHandleUnsupportedElement=true，才返回此参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultUrl *string `json:"ResultUrl,omitnil,omitempty" name:"ResultUrl"`

	// 错误PPT页面列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Slides []*PPTErrSlide `json:"Slides,omitnil,omitempty" name:"Slides"`

	// 任务的当前状态 - QUEUED: 正在排队等待 - PROCESSING: 执行中 - FINISHED: 执行完成	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 当前进度,取值范围为0~100
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 错误列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errs []*PPTErr `json:"Errs,omitnil,omitempty" name:"Errs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePPTCheckResponse struct {
	*tchttp.BaseResponse
	Response *DescribePPTCheckResponseParams `json:"Response"`
}

func (r *DescribePPTCheckResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePPTCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunningTasksRequestParams struct {
	// 应用的SdkAppID
	SdkAppID *int64 `json:"SdkAppID,omitnil,omitempty" name:"SdkAppID"`

	// 指定需要获取的任务类型。
	// 有效取值如下：
	// - TranscodeH5: 动态转码任务，文档转HTML5页面
	// - TranscodeJPG: 静态转码任务，文档转图片
	// - WhiteboardPush: 白板推流任务
	// - OnlineRecord: 实时录制任务
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 分页获取时的任务偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每次获取任务列表时最大获取任务数，默认值为100。
	// 有效取值范围：[1, 500]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRunningTasksRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppID
	SdkAppID *int64 `json:"SdkAppID,omitnil,omitempty" name:"SdkAppID"`

	// 指定需要获取的任务类型。
	// 有效取值如下：
	// - TranscodeH5: 动态转码任务，文档转HTML5页面
	// - TranscodeJPG: 静态转码任务，文档转图片
	// - WhiteboardPush: 白板推流任务
	// - OnlineRecord: 实时录制任务
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 分页获取时的任务偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每次获取任务列表时最大获取任务数，默认值为100。
	// 有效取值范围：[1, 500]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRunningTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunningTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppID")
	delete(f, "TaskType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRunningTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRunningTasksResponseParams struct {
	// 当前正在执行中的任务总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 任务信息列表
	Tasks []*RunningTaskItem `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRunningTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRunningTasksResponseParams `json:"Response"`
}

func (r *DescribeRunningTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRunningTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotTaskRequestParams struct {
	// 查询任务ID
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 任务SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeSnapshotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 查询任务ID
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 任务SdkAppId
	SdkAppId *uint64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeSnapshotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotTaskResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 任务状态
	// Running - 任务执行中
	// Finished - 任务已结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务创建时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务完成时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *uint64 `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 任务结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *SnapshotResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotTaskResponseParams `json:"Response"`
}

func (r *DescribeSnapshotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeByUrlRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 经过URL编码后的转码文件地址。URL 编码会将字符转换为可通过因特网传输的格式，比如文档地址为http://example.com/测试.pdf，经过URL编码之后为http://example.com/%E6%B5%8B%E8%AF%95.pdf。为了提高URL解析的成功率，请对URL进行编码。	
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DescribeTranscodeByUrlRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 经过URL编码后的转码文件地址。URL 编码会将字符转换为可通过因特网传输的格式，比如文档地址为http://example.com/测试.pdf，经过URL编码之后为http://example.com/%E6%B5%8B%E8%AF%95.pdf。为了提高URL解析的成功率，请对URL进行编码。	
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *DescribeTranscodeByUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeByUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeByUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeByUrlResponseParams struct {
	// 转码的当前进度,取值范围为0~100
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务的当前状态
	// - QUEUED: 正在排队等待转换
	// - PROCESSING: 转换中
	// - FINISHED: 转换完成
	// - EXCEPTION: 转换异常
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 转码任务的唯一标识Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTranscodeByUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeByUrlResponseParams `json:"Response"`
}

func (r *DescribeTranscodeByUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeByUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeCallbackRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeTranscodeCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeTranscodeCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeCallbackResponseParams struct {
	// 文档转码回调地址
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 文档转码回调鉴权密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeCallbackResponseParams `json:"Response"`
}

func (r *DescribeTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档转码任务的唯一标识Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTranscodeRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档转码任务的唯一标识Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTranscodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTranscodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTranscodeResponseParams struct {
	// 文档的总页数
	Pages *int64 `json:"Pages,omitnil,omitempty" name:"Pages"`

	// 转码的当前进度,取值范围为0~100
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 文档的分辨率
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 转码完成后结果的URL
	// 动态转码：PPT转动态H5的链接
	// 静态转码：文档每一页的图片URL前缀，比如，该URL前缀为`http://example.com/g0jb42ps49vtebjshilb/`，那么文档第1页的图片URL为
	// `http://example.com/g0jb42ps49vtebjshilb/1.jpg`，其它页以此类推
	ResultUrl *string `json:"ResultUrl,omitnil,omitempty" name:"ResultUrl"`

	// 任务的当前状态
	// - QUEUED: 正在排队等待转换
	// - PROCESSING: 转换中
	// - FINISHED: 转换完成
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 转码任务的唯一标识Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 文档的文件名
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 缩略图URL前缀，比如，该URL前缀为`http://example.com/g0jb42ps49vtebjshilb/ `，那么动态PPT第1页的缩略图URL为
	// `http://example.com/g0jb42ps49vtebjshilb/1.jpg`，其它页以此类推
	// 
	// 如果发起文档转码请求参数中带了ThumbnailResolution参数，并且转码类型为动态转码，该参数不为空，其余情况该参数为空字符串
	ThumbnailUrl *string `json:"ThumbnailUrl,omitnil,omitempty" name:"ThumbnailUrl"`

	// 动态转码缩略图生成分辨率
	ThumbnailResolution *string `json:"ThumbnailResolution,omitnil,omitempty" name:"ThumbnailResolution"`

	// 转码压缩文件下载的URL，如果发起文档转码请求参数中`CompressFileType`为空或者不是支持的压缩格式，该参数为空字符串
	CompressFileUrl *string `json:"CompressFileUrl,omitnil,omitempty" name:"CompressFileUrl"`

	// 资源清单文件下载URL(内测体验)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceListUrl *string `json:"ResourceListUrl,omitnil,omitempty" name:"ResourceListUrl"`

	// 文档制作方式(内测体验)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ext *string `json:"Ext,omitnil,omitempty" name:"Ext"`

	// 文档转码任务创建时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文档转码任务分配时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssignTime *uint64 `json:"AssignTime,omitnil,omitempty" name:"AssignTime"`

	// 文档转码任务完成时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedTime *uint64 `json:"FinishedTime,omitnil,omitempty" name:"FinishedTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTranscodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTranscodeResponseParams `json:"Response"`
}

func (r *DescribeTranscodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTranscodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskCallbackRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeVideoGenerationTaskCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeVideoGenerationTaskCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoGenerationTaskCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskCallbackResponseParams struct {
	// 录制视频生成回调地址
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 录制视频生成回调鉴权密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoGenerationTaskCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoGenerationTaskCallbackResponseParams `json:"Response"`
}

func (r *DescribeVideoGenerationTaskCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制视频生成的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeVideoGenerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 录制视频生成的任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeVideoGenerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoGenerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoGenerationTaskResponseParams struct {
	// 任务对应的群组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 任务对应的房间号
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 任务的Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 已废弃
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 录制视频生成任务状态
	// - QUEUED: 正在排队
	// - PROCESSING: 正在生成视频
	// - FINISHED: 生成视频结束（成功完成或失败结束，可以通过错误码和错误信息进一步判断）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 回放视频总时长,单位：毫秒
	TotalTime *int64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// 已废弃，请使用`VideoInfoList`参数
	VideoInfos *VideoInfo `json:"VideoInfos,omitnil,omitempty" name:"VideoInfos"`

	// 录制视频生成视频列表
	VideoInfoList []*VideoInfo `json:"VideoInfoList,omitnil,omitempty" name:"VideoInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoGenerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoGenerationTaskResponseParams `json:"Response"`
}

func (r *DescribeVideoGenerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoGenerationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningCallbackRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeWarningCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeWarningCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWarningCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWarningCallbackResponseParams struct {
	// 告警事件回调地址，如果未设置回调地址，该字段为空字符串
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 告警回调鉴权密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWarningCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWarningCallbackResponseParams `json:"Response"`
}

func (r *DescribeWarningCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWarningCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushCallbackRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

type DescribeWhiteboardPushCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`
}

func (r *DescribeWhiteboardPushCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteboardPushCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushCallbackResponseParams struct {
	// 白板推流事件回调地址，如果未设置回调地址，该字段为空字符串
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 白板推流回调鉴权密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteboardPushCallbackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteboardPushCallbackResponseParams `json:"Response"`
}

func (r *DescribeWhiteboardPushCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 白板推流任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeWhiteboardPushRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 白板推流任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWhiteboardPushRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWhiteboardPushResponseParams struct {
	// 推流结束原因，
	// - AUTO: 房间内长时间没有音视频上行及白板操作导致自动停止推流
	// - USER_CALL: 主动调用了停止推流接口
	// - EXCEPTION: 推流异常结束
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 需要查询结果的白板推流任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 推流任务状态
	// - PREPARED: 表示推流正在准备中（进房/启动推流服务等操作）
	// - PUSHING: 表示推流已开始
	// - STOPPED: 表示推流已停止
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 房间号
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 白板的群组 Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 推流用户Id
	PushUserId *string `json:"PushUserId,omitnil,omitempty" name:"PushUserId"`

	// 实际开始推流时间，Unix 时间戳，单位秒
	PushStartTime *int64 `json:"PushStartTime,omitnil,omitempty" name:"PushStartTime"`

	// 实际停止推流时间，Unix 时间戳，单位秒
	PushStopTime *int64 `json:"PushStopTime,omitnil,omitempty" name:"PushStopTime"`

	// 推流过程中出现异常的次数
	ExceptionCnt *int64 `json:"ExceptionCnt,omitnil,omitempty" name:"ExceptionCnt"`

	// 白板推流首帧对应的IM时间戳，可用于录制回放时IM聊天消息与白板推流视频进行同步对时。
	IMSyncTime *int64 `json:"IMSyncTime,omitnil,omitempty" name:"IMSyncTime"`

	// 备份推流任务结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Backup *string `json:"Backup,omitnil,omitempty" name:"Backup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWhiteboardPushResponseParams `json:"Response"`
}

func (r *DescribeWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWhiteboardPushResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExcelParam struct {
	// 表格转码纸张（画布）大小，默认为0。
	// 0 -- A4
	// 1 -- A2 
	// 2 -- A0
	// 
	// 注：当设置的值超出合法取值范围时，将采用默认值。
	PaperSize *int64 `json:"PaperSize,omitnil,omitempty" name:"PaperSize"`

	// 表格文件转换纸张方向，默认为0。
	// 0 -- 代表垂直方向
	// 非0 -- 代表水平方向
	PaperDirection *int64 `json:"PaperDirection,omitnil,omitempty" name:"PaperDirection"`
}

type Interrupt struct {
	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 视频流断流次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type LayoutParams struct {
	// 流画面宽，取值范围[2,3000]
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 流画面高，取值范围[2,3000]
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 当前画面左上角顶点相对于Canvas左上角顶点的x轴偏移量，默认为0，取值范围[0,3000]
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 当前画面左上角顶点相对于Canvas左上角顶点的y轴偏移量，默认为0， 取值范围[0,3000]
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 画面z轴位置，默认为0
	// z轴确定了重叠画面的遮盖顺序，z轴值大的画面处于顶层
	ZOrder *int64 `json:"ZOrder,omitnil,omitempty" name:"ZOrder"`
}

type MixStream struct {
	// 是否开启混流
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 是否禁用音频混流
	DisableAudio *bool `json:"DisableAudio,omitnil,omitempty" name:"DisableAudio"`

	// 内置混流布局模板ID, 取值 [1, 2], 区别见内置混流布局模板样式示例说明
	// 在没有填Custom字段时候，ModelId是必填的
	ModelId *int64 `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 老师用户ID
	// 此字段只有在ModelId填了的情况下生效
	// 填写TeacherId的效果是把指定为TeacherId的用户视频流显示在内置模板的第一个小画面中
	TeacherId *string `json:"TeacherId,omitnil,omitempty" name:"TeacherId"`

	// 自定义混流布局参数
	// 当此字段存在时，ModelId 及 TeacherId 字段将被忽略
	Custom *CustomLayout `json:"Custom,omitnil,omitempty" name:"Custom"`
}

type OmittedDuration struct {
	// 录制暂停时间戳对应的视频播放时间(单位: 毫秒)
	VideoTime *int64 `json:"VideoTime,omitnil,omitempty" name:"VideoTime"`

	// 录制暂停时间戳(单位: 毫秒)
	PauseTime *int64 `json:"PauseTime,omitnil,omitempty" name:"PauseTime"`

	// 录制恢复时间戳(单位: 毫秒)
	ResumeTime *int64 `json:"ResumeTime,omitnil,omitempty" name:"ResumeTime"`
}

type PPTErr struct {
	// 元素名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 0: 不支持的墨迹类型，1: 不支持自动翻页，2: 存在已损坏音视频，3: 存在不可访问资源，4: 只读文件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type PPTErrSlide struct {
	// 异常元素存在的页面，由页面类型+页码组成，页码类型包括：幻灯片、幻灯片母版、幻灯片布局等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *string `json:"Page,omitnil,omitempty" name:"Page"`

	// 错误元素列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errs []*PPTErr `json:"Errs,omitnil,omitempty" name:"Errs"`
}

// Predefined struct for user
type PauseOnlineRecordRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实时录制任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type PauseOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实时录制任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *PauseOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseOnlineRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *PauseOnlineRecordResponseParams `json:"Response"`
}

func (r *PauseOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordControl struct {
	// 设置是否开启录制控制参数，只有设置为true的时候，录制控制参数才生效。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 设置是否禁用录制的全局控制参数。一般与`StreamControls`参数配合使用。
	// 
	// true - 所有流都不录制。
	// false - 所有流都录制。默认为false。
	// 
	// 这里的设置对所有流都生效，如果同时在 `StreamControls` 列表中针对指定流设置了控制参数，则优先采用`StreamControls`中设置的控制参数。
	DisableRecord *bool `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 设置是否禁用所有流的音频录制的全局控制参数。一般与`StreamControls`参数配合使用。
	// 
	// true - 所有流的录制都不对音频进行录制。
	// false - 所有流的录制都需要对音频进行录制。默认为false。
	// 
	// 这里的设置对所有流都生效，如果同时在 `StreamControls` 列表中针对指定流设置了控制参数，则优先采用`StreamControls`中设置的控制参数。
	DisableAudio *bool `json:"DisableAudio,omitnil,omitempty" name:"DisableAudio"`

	// 设置是否所有流都只录制小画面的全局控制参数。一般与`StreamControls`参数配合使用。
	// 
	// true - 所有流都只录制小画面。设置为true时，请确保上行端在推流的时候同时上行了小画面，否则录制视频可能是黑屏。
	// false - 所有流都录制大画面，默认为false。
	// 
	// 这里的设置对所有流都生效，如果同时在 `StreamControls` 列表中针对指定流设置了控制参数，则优先采用`StreamControls`中设置的控制参数。
	PullSmallVideo *bool `json:"PullSmallVideo,omitnil,omitempty" name:"PullSmallVideo"`

	// 针对具体流指定控制参数，如果列表为空，则所有流采用全局配置的控制参数进行录制。列表不为空，则列表中指定的流将优先按此列表指定的控制参数进行录制。
	StreamControls []*StreamControl `json:"StreamControls,omitnil,omitempty" name:"StreamControls"`
}

// Predefined struct for user
type ResumeOnlineRecordRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 恢复录制的实时录制任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type ResumeOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 恢复录制的实时录制任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *ResumeOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeOnlineRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *ResumeOnlineRecordResponseParams `json:"Response"`
}

func (r *ResumeOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunningTaskItem struct {
	// 应用SdkAppID
	SdkAppID *int64 `json:"SdkAppID,omitnil,omitempty" name:"SdkAppID"`

	// 任务ID
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 任务类型
	// - TranscodeH5: 动态转码任务，文档转HTML5页面
	// - TranscodeJPG: 静态转码任务，文档转图片
	// - WhiteboardPush: 白板推流任务
	// - OnlineRecord: 实时录制任务
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务取消时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancelTime *string `json:"CancelTime,omitnil,omitempty" name:"CancelTime"`

	// 任务状态
	// - QUEUED: 任务正在排队等待执行中
	// - PROCESSING: 任务正在执行中 
	// - FINISHED: 任务已完成
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务当前进度
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 转码任务中转码文件的原始URL
	// 此参数只有任务类型为TranscodeH5、TranscodeJPG类型时才会有有效值。其他任务类型为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileURL *string `json:"FileURL,omitnil,omitempty" name:"FileURL"`

	// 房间号
	// 
	// 当任务类型为TranscodeH5、TranscodeJPG时，房间号为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomID *int64 `json:"RoomID,omitnil,omitempty" name:"RoomID"`
}

// Predefined struct for user
type SetOnlineRecordCallbackKeyRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置实时录制回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥。回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

type SetOnlineRecordCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置实时录制回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥。回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

func (r *SetOnlineRecordCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetOnlineRecordCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOnlineRecordCallbackKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetOnlineRecordCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetOnlineRecordCallbackKeyResponseParams `json:"Response"`
}

func (r *SetOnlineRecordCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOnlineRecordCallbackRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实时录制任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头。回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40258
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type SetOnlineRecordCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 实时录制任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头。回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40258
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *SetOnlineRecordCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetOnlineRecordCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetOnlineRecordCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetOnlineRecordCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetOnlineRecordCallbackResponseParams `json:"Response"`
}

func (r *SetOnlineRecordCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetOnlineRecordCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPPTCheckCallbackKeyRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257	
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

type SetPPTCheckCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257	
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

func (r *SetPPTCheckCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPPTCheckCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPPTCheckCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPPTCheckCallbackKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetPPTCheckCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetPPTCheckCallbackKeyResponseParams `json:"Response"`
}

func (r *SetPPTCheckCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPPTCheckCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPPTCheckCallbackRequestParams struct {
	// 客户的SdkAppId	
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 进度回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头。 回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260#c9cbe05f-fe1a-4410-b4dc-40cc301c7b81	
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type SetPPTCheckCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId	
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 进度回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头。 回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260#c9cbe05f-fe1a-4410-b4dc-40cc301c7b81	
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *SetPPTCheckCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPPTCheckCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPPTCheckCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPPTCheckCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetPPTCheckCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetPPTCheckCallbackResponseParams `json:"Response"`
}

func (r *SetPPTCheckCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPPTCheckCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackKeyRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置文档转码回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

type SetTranscodeCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置文档转码回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

func (r *SetTranscodeCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTranscodeCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetTranscodeCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetTranscodeCallbackKeyResponseParams `json:"Response"`
}

func (r *SetTranscodeCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档转码进度回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头。
	// 回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type SetTranscodeCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 文档转码进度回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头。
	// 回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *SetTranscodeCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTranscodeCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTranscodeCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetTranscodeCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetTranscodeCallbackResponseParams `json:"Response"`
}

func (r *SetTranscodeCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTranscodeCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackKeyRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置视频生成回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

type SetVideoGenerationTaskCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置视频生成回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

func (r *SetVideoGenerationTaskCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVideoGenerationTaskCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetVideoGenerationTaskCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetVideoGenerationTaskCallbackKeyResponseParams `json:"Response"`
}

func (r *SetVideoGenerationTaskCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课后录制任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type SetVideoGenerationTaskCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 课后录制任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *SetVideoGenerationTaskCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVideoGenerationTaskCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVideoGenerationTaskCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetVideoGenerationTaskCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetVideoGenerationTaskCallbackResponseParams `json:"Response"`
}

func (r *SetVideoGenerationTaskCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVideoGenerationTaskCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWarningCallbackRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 告警回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头。
	// 回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/90112
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 设置告警回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

type SetWarningCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 告警回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持http或https协议，即回调地址以http://或https://开头。
	// 回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/90112
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`

	// 设置告警回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

func (r *SetWarningCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWarningCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetWarningCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWarningCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetWarningCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetWarningCallbackResponseParams `json:"Response"`
}

func (r *SetWarningCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWarningCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackKeyRequestParams struct {
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置白板推流回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥。回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

type SetWhiteboardPushCallbackKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 设置白板推流回调鉴权密钥，最长64字符，如果传入空字符串，那么删除现有的鉴权回调密钥。回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	CallbackKey *string `json:"CallbackKey,omitnil,omitempty" name:"CallbackKey"`
}

func (r *SetWhiteboardPushCallbackKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "CallbackKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetWhiteboardPushCallbackKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetWhiteboardPushCallbackKeyResponse struct {
	*tchttp.BaseResponse
	Response *SetWhiteboardPushCallbackKeyResponseParams `json:"Response"`
}

func (r *SetWhiteboardPushCallbackKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 白板推流任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头。回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type SetWhiteboardPushCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 白板推流任务结果回调地址，如果传空字符串会删除原来的回调地址配置，回调地址仅支持 http或https协议，即回调地址以http://或https://开头。回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40257
	Callback *string `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *SetWhiteboardPushCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetWhiteboardPushCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetWhiteboardPushCallbackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetWhiteboardPushCallbackResponse struct {
	*tchttp.BaseResponse
	Response *SetWhiteboardPushCallbackResponseParams `json:"Response"`
}

func (r *SetWhiteboardPushCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetWhiteboardPushCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotCOS struct {
	// cos所在腾讯云账号uin
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// cos所在地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// cos存储桶名称
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 板书文件存储根目录
	TargetDir *string `json:"TargetDir,omitnil,omitempty" name:"TargetDir"`

	// CDN加速域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type SnapshotResult struct {
	// 任务执行错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 任务执行错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 快照生成图片总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 快照图片链接列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Snapshots []*string `json:"Snapshots,omitnil,omitempty" name:"Snapshots"`
}

type SnapshotWhiteboard struct {
	// 白板宽度大小，默认为1280，有效取值范围[0，2560]
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 白板高度大小，默认为720，有效取值范围[0，2560]
	Height *uint64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 白板初始化参数的JSON转义字符串，透传到白板 SDK
	InitParams *string `json:"InitParams,omitnil,omitempty" name:"InitParams"`
}

// Predefined struct for user
type StartOnlineRecordRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要录制的白板房间号，取值范围: (1, 4294967295)。
	// 
	// 1. 在没有指定`GroupId`的情况下，实时录制默认以`RoomId`的字符串表达形式作为同步白板信令的IM群组ID（比如`RoomId`为12358，则IM群组ID为"12358"），并加群进行信令同步，请在开始录制前确保相应IM群组已创建完成，否则会导致录制失败。
	// 2. 在没有指定`TRTCRoomId`和`TRTCRoomIdStr`的情况下，默认会以`RoomId`作为TRTC房间号进房拉流进行录制。
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用于录制服务进房的用户ID，最大长度不能大于60个字节，格式为`tic_record_user_${RoomId}_${Random}`，其中 `${RoomId} `与录制房间号对应，`${Random}`为一个随机字符串。
	// 该ID必须是一个单独的未在SDK中使用的ID，录制服务使用这个用户ID进入房间进行音视频与白板录制，若该ID和SDK中使用的ID重复，会导致SDK和录制服务互踢，影响正常录制。
	RecordUserId *string `json:"RecordUserId,omitnil,omitempty" name:"RecordUserId"`

	// 与`RecordUserId`对应的IM签名
	RecordUserSig *string `json:"RecordUserSig,omitnil,omitempty" name:"RecordUserSig"`

	// 白板进行信令同步的 IM 群组 ID。
	// 在没有指定`GroupId`的情况下，实时录制服务将使用 `RoomId` 的字符串形式作为同步白板信令的IM群组ID。
	// 在指定了`GroupId`的情况下，实时录制将优先使用`GroupId`作为同步白板信令的群组ID。请在开始录制前确保相应的IM群组已创建完成，否则会导致录制失败。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 录制视频拼接参数
	Concat *Concat `json:"Concat,omitnil,omitempty" name:"Concat"`

	// 录制白板参数，例如白板宽高等
	Whiteboard *Whiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 录制混流参数
	// 特别说明：
	// 1. 混流功能需要根据额外开通， 请联系腾讯云互动白板客服人员
	// 2. 使用混流功能，必须提供 Extras 参数，且 Extras 参数中必须包含 "MIX_STREAM"
	MixStream *MixStream `json:"MixStream,omitnil,omitempty" name:"MixStream"`

	// 使用到的高级功能列表
	// 可以选值列表：
	// MIX_STREAM - 混流功能
	Extras []*string `json:"Extras,omitnil,omitempty" name:"Extras"`

	// 是否需要在结果回调中返回各路流的纯音频录制文件，文件格式为mp3
	AudioFileNeeded *bool `json:"AudioFileNeeded,omitnil,omitempty" name:"AudioFileNeeded"`

	// 录制控制参数，用于更精细地指定需要录制哪些流，某一路流是否禁用音频，是否只录制小画面等
	RecordControl *RecordControl `json:"RecordControl,omitnil,omitempty" name:"RecordControl"`

	// 录制模式
	// 
	// REALTIME_MODE - 实时录制模式（默认）
	// VIDEO_GENERATION_MODE - 视频生成模式（内测中，需邮件申请开通）
	RecordMode *string `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 聊天群组ID，此字段仅适用于`视频生成模式`
	// 
	// 在`视频生成模式`下，默认会记录白板群组内的非白板信令消息，如果指定了`ChatGroupId`，则会记录指定群ID的聊天消息。
	ChatGroupId *string `json:"ChatGroupId,omitnil,omitempty" name:"ChatGroupId"`

	// 自动停止录制超时时间，单位秒，取值范围[300, 86400], 默认值为300秒。
	// 
	// 当超过设定时间房间内没有音视频上行且没有白板操作的时候，录制服务会自动停止当前录制任务。
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitnil,omitempty" name:"AutoStopTimeout"`

	// 内部参数，可忽略
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// TRTC数字类型房间号，取值范围: (1, 4294967295)。
	// 
	// 在同时指定了`RoomId`与`TRTCRoomId`的情况下，优先使用`TRTCRoomId`作为实时录制拉TRTC流的TRTC房间号。
	// 
	// 当指定了`TRTCRoomIdStr`的情况下，此字段将被忽略。
	TRTCRoomId *int64 `json:"TRTCRoomId,omitnil,omitempty" name:"TRTCRoomId"`

	// TRTC字符串类型房间号。
	// 
	// 在指定了`TRTCRoomIdStr`的情况下，会优先使用`TRTCRoomIdStr`作为实时录制拉TRTC流的TRTC房间号。
	TRTCRoomIdStr *string `json:"TRTCRoomIdStr,omitnil,omitempty" name:"TRTCRoomIdStr"`
}

type StartOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要录制的白板房间号，取值范围: (1, 4294967295)。
	// 
	// 1. 在没有指定`GroupId`的情况下，实时录制默认以`RoomId`的字符串表达形式作为同步白板信令的IM群组ID（比如`RoomId`为12358，则IM群组ID为"12358"），并加群进行信令同步，请在开始录制前确保相应IM群组已创建完成，否则会导致录制失败。
	// 2. 在没有指定`TRTCRoomId`和`TRTCRoomIdStr`的情况下，默认会以`RoomId`作为TRTC房间号进房拉流进行录制。
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用于录制服务进房的用户ID，最大长度不能大于60个字节，格式为`tic_record_user_${RoomId}_${Random}`，其中 `${RoomId} `与录制房间号对应，`${Random}`为一个随机字符串。
	// 该ID必须是一个单独的未在SDK中使用的ID，录制服务使用这个用户ID进入房间进行音视频与白板录制，若该ID和SDK中使用的ID重复，会导致SDK和录制服务互踢，影响正常录制。
	RecordUserId *string `json:"RecordUserId,omitnil,omitempty" name:"RecordUserId"`

	// 与`RecordUserId`对应的IM签名
	RecordUserSig *string `json:"RecordUserSig,omitnil,omitempty" name:"RecordUserSig"`

	// 白板进行信令同步的 IM 群组 ID。
	// 在没有指定`GroupId`的情况下，实时录制服务将使用 `RoomId` 的字符串形式作为同步白板信令的IM群组ID。
	// 在指定了`GroupId`的情况下，实时录制将优先使用`GroupId`作为同步白板信令的群组ID。请在开始录制前确保相应的IM群组已创建完成，否则会导致录制失败。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 录制视频拼接参数
	Concat *Concat `json:"Concat,omitnil,omitempty" name:"Concat"`

	// 录制白板参数，例如白板宽高等
	Whiteboard *Whiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 录制混流参数
	// 特别说明：
	// 1. 混流功能需要根据额外开通， 请联系腾讯云互动白板客服人员
	// 2. 使用混流功能，必须提供 Extras 参数，且 Extras 参数中必须包含 "MIX_STREAM"
	MixStream *MixStream `json:"MixStream,omitnil,omitempty" name:"MixStream"`

	// 使用到的高级功能列表
	// 可以选值列表：
	// MIX_STREAM - 混流功能
	Extras []*string `json:"Extras,omitnil,omitempty" name:"Extras"`

	// 是否需要在结果回调中返回各路流的纯音频录制文件，文件格式为mp3
	AudioFileNeeded *bool `json:"AudioFileNeeded,omitnil,omitempty" name:"AudioFileNeeded"`

	// 录制控制参数，用于更精细地指定需要录制哪些流，某一路流是否禁用音频，是否只录制小画面等
	RecordControl *RecordControl `json:"RecordControl,omitnil,omitempty" name:"RecordControl"`

	// 录制模式
	// 
	// REALTIME_MODE - 实时录制模式（默认）
	// VIDEO_GENERATION_MODE - 视频生成模式（内测中，需邮件申请开通）
	RecordMode *string `json:"RecordMode,omitnil,omitempty" name:"RecordMode"`

	// 聊天群组ID，此字段仅适用于`视频生成模式`
	// 
	// 在`视频生成模式`下，默认会记录白板群组内的非白板信令消息，如果指定了`ChatGroupId`，则会记录指定群ID的聊天消息。
	ChatGroupId *string `json:"ChatGroupId,omitnil,omitempty" name:"ChatGroupId"`

	// 自动停止录制超时时间，单位秒，取值范围[300, 86400], 默认值为300秒。
	// 
	// 当超过设定时间房间内没有音视频上行且没有白板操作的时候，录制服务会自动停止当前录制任务。
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitnil,omitempty" name:"AutoStopTimeout"`

	// 内部参数，可忽略
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// TRTC数字类型房间号，取值范围: (1, 4294967295)。
	// 
	// 在同时指定了`RoomId`与`TRTCRoomId`的情况下，优先使用`TRTCRoomId`作为实时录制拉TRTC流的TRTC房间号。
	// 
	// 当指定了`TRTCRoomIdStr`的情况下，此字段将被忽略。
	TRTCRoomId *int64 `json:"TRTCRoomId,omitnil,omitempty" name:"TRTCRoomId"`

	// TRTC字符串类型房间号。
	// 
	// 在指定了`TRTCRoomIdStr`的情况下，会优先使用`TRTCRoomIdStr`作为实时录制拉TRTC流的TRTC房间号。
	TRTCRoomIdStr *string `json:"TRTCRoomIdStr,omitnil,omitempty" name:"TRTCRoomIdStr"`
}

func (r *StartOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "RecordUserId")
	delete(f, "RecordUserSig")
	delete(f, "GroupId")
	delete(f, "Concat")
	delete(f, "Whiteboard")
	delete(f, "MixStream")
	delete(f, "Extras")
	delete(f, "AudioFileNeeded")
	delete(f, "RecordControl")
	delete(f, "RecordMode")
	delete(f, "ChatGroupId")
	delete(f, "AutoStopTimeout")
	delete(f, "ExtraData")
	delete(f, "TRTCRoomId")
	delete(f, "TRTCRoomIdStr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartOnlineRecordResponseParams struct {
	// 录制任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *StartOnlineRecordResponseParams `json:"Response"`
}

func (r *StartOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartWhiteboardPushRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要推流的白板房间号，取值范围: (1, 4294967295)。
	// 
	// 1. 在没有指定`GroupId`的情况下，白板推流默认以`RoomId`的字符串表达形式作为IM群组ID（比如RoomId为1234，则IM群组ID为"1234"），并加群进行信令同步，请在开始推流前确保相应IM群组已创建完成，否则会导致推流失败。
	// 2. 在没有指定`TRTCRoomId`和`TRTCRoomIdStr`的情况下，默认会以`RoomId`作为白板流进行推流的TRTC房间号。
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用于白板推流服务进入白板房间的用户ID。在没有额外指定`IMAuthParam`和`TRTCAuthParam`的情况下，这个用户ID同时会用于IM登录、IM加群、TRTC进房推流等操作。
	// 用户ID最大长度不能大于60个字节，该用户ID必须是一个单独的未同时在其他地方使用的用户ID，白板推流服务使用这个用户ID进入房间进行白板音视频推流，若该用户ID和其他地方同时在使用的用户ID重复，会导致白板推流服务与其他使用场景账号互踢，影响正常推流。
	PushUserId *string `json:"PushUserId,omitnil,omitempty" name:"PushUserId"`

	// 与`PushUserId`对应的IM签名(usersig)。
	PushUserSig *string `json:"PushUserSig,omitnil,omitempty" name:"PushUserSig"`

	// 白板参数，例如白板宽高、背景颜色等
	Whiteboard *Whiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 自动停止推流超时时间，单位秒，取值范围[300, 259200], 默认值为1800秒。
	// 
	// 当白板超过设定时间没有操作的时候，白板推流服务会自动停止白板推流。
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitnil,omitempty" name:"AutoStopTimeout"`

	// 对主白板推流任务进行操作时，是否同时同步操作备份任务
	AutoManageBackup *bool `json:"AutoManageBackup,omitnil,omitempty" name:"AutoManageBackup"`

	// 备份白板推流相关参数。
	// 
	// 指定了备份参数的情况下，白板推流服务会在房间内新增一路白板画面视频流，即同一个房间内会有两路白板画面推流。
	Backup *WhiteboardPushBackupParam `json:"Backup,omitnil,omitempty" name:"Backup"`

	// TRTC高级权限控制参数，如果在实时音视频开启了高级权限控制功能，必须提供PrivateMapKey才能保证正常推流。
	PrivateMapKey *string `json:"PrivateMapKey,omitnil,omitempty" name:"PrivateMapKey"`

	// 白板推流视频帧率，取值范围[0, 30]，默认20fps
	VideoFPS *int64 `json:"VideoFPS,omitnil,omitempty" name:"VideoFPS"`

	// 白板推流码率， 取值范围[0, 2000]，默认1200kbps。
	// 
	// 这里的码率设置是一个参考值，实际推流的时候使用的是动态码率，所以真实码率不会固定为指定值，会在指定值附近波动。
	VideoBitrate *int64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 在实时音视频云端录制模式选择为 `指定用户录制` 模式的时候是否自动录制白板推流。
	// 
	// 默认在实时音视频的云端录制模式选择为 `指定用户录制` 模式的情况下，不会自动进行白板推流录制，如果希望进行白板推流录制，请将此参数设置为true。
	// 
	// 如果实时音视频的云端录制模式选择为 `全局自动录制` 模式，可忽略此参数。
	AutoRecord *bool `json:"AutoRecord,omitnil,omitempty" name:"AutoRecord"`

	// 指定白板推流这路流在音视频云端录制中的RecordID，指定的RecordID会用于填充实时音视频云端录制完成后的回调消息中的 "userdefinerecordid" 字段内容，便于您更方便的识别录制回调，以及在点播媒体资源管理中查找相应的录制视频文件。
	// 
	// 限制长度为64字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
	// 
	// 此字段设置后，不管`AutoRecord`字段取值如何，都将自动进行白板推流录制。
	// 
	// 默认RecordId生成规则如下：
	// urlencode(SdkAppID_RoomID_PushUserID)
	// 
	// 例如：
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// 那么：RecordId = 12345678_12345_push_user_1
	UserDefinedRecordId *string `json:"UserDefinedRecordId,omitnil,omitempty" name:"UserDefinedRecordId"`

	// 在实时音视频旁路推流模式选择为`指定用户旁路`模式的时候，是否自动旁路白板推流。
	// 
	// 默认在实时音视频的旁路推流模式选择为 `指定用户旁路` 模式的情况下，不会自动旁路白板推流，如果希望旁路白板推流，请将此参数设置为true。
	// 
	// 如果实时音视频的旁路推流模式选择为 `全局自动旁路` 模式，可忽略此参数。
	AutoPublish *bool `json:"AutoPublish,omitnil,omitempty" name:"AutoPublish"`

	// 指定实时音视频在旁路白板推流这路流时的StreamID，设置之后，您就可以在腾讯云直播 CDN 上通过标准直播方案（FLV或HLS）播放该用户的音视频流。
	// 
	// 限制长度为64字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
	// 
	// 此字段设置后，不管`AutoPublish`字段取值如何，都将自动旁路白板推流。
	// 
	// 默认StreamID生成规则如下：
	// urlencode(SdkAppID_RoomID_PushUserID_main)
	// 
	// 例如：
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// 那么：StreamID = 12345678_12345_push_user_1_main
	UserDefinedStreamId *string `json:"UserDefinedStreamId,omitnil,omitempty" name:"UserDefinedStreamId"`

	// 内部参数，不需要关注此参数
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// TRTC数字类型房间号，取值范围: (1, 4294967295)。
	// 
	// 在同时指定了`RoomId`与`TRTCRoomId`的情况下，优先使用`TRTCRoomId`作为白板流进行推流的TRTC房间号。
	// 
	// 当指定了`TRTCRoomIdStr`的情况下，此字段将被忽略。
	TRTCRoomId *int64 `json:"TRTCRoomId,omitnil,omitempty" name:"TRTCRoomId"`

	// TRTC字符串类型房间号。
	// 
	// 在指定了`TRTCRoomIdStr`的情况下，会优先使用`TRTCRoomIdStr`作为白板流进行推流的TRTC房间号。
	TRTCRoomIdStr *string `json:"TRTCRoomIdStr,omitnil,omitempty" name:"TRTCRoomIdStr"`

	// IM鉴权信息参数，用于IM鉴权。
	// 当白板信令所使用的IM应用与白板应用的SdkAppId不一致时，可以通过此参数提供对应IM应用鉴权信息。
	// 
	// 如果提供了此参数，白板推流服务会优先使用此参数指定的SdkAppId作为白板信令的传输通道，否则使用公共参数中的SdkAppId作为白板信令的传输通道。
	IMAuthParam *AuthParam `json:"IMAuthParam,omitnil,omitempty" name:"IMAuthParam"`

	// TRTC鉴权信息参数，用于TRTC进房推流鉴权。
	// 当需要推流到的TRTC房间所对应的TRTC应用与白板应用的SdkAppId不一致时，可以通过此参数提供对应的TRTC应用鉴权信息。
	// 
	// 如果提供了此参数，白板推流服务会优先使用此参数指定的SdkAppId作为白板推流的目标TRTC应用，否则使用公共参数中的SdkAppId作为白板推流的目标TRTC应用。
	TRTCAuthParam *AuthParam `json:"TRTCAuthParam,omitnil,omitempty" name:"TRTCAuthParam"`

	// 指定白板推流时推流用户进TRTC房间的进房模式。默认为 TRTCAppSceneVideoCall
	// 
	// TRTCAppSceneVideoCall - 视频通话场景，即绝大多数时间都是两人或两人以上视频通话的场景，内部编码器和网络协议优化侧重流畅性，降低通话延迟和卡顿率。
	// TRTCAppSceneLIVE - 直播场景，即绝大多数时间都是一人直播，偶尔有多人视频互动的场景，内部编码器和网络协议优化侧重性能和兼容性，性能和清晰度表现更佳。
	TRTCEnterRoomMode *string `json:"TRTCEnterRoomMode,omitnil,omitempty" name:"TRTCEnterRoomMode"`

	// 白板进行信令同步的 IM 群组 ID。
	// 在没有指定`GroupId`的情况下，白板推流服务将使用 `RoomId` 的字符串形式作为同步白板信令的IM群组ID。
	// 在指定了`GroupId`的情况下，白板推流将优先`GroupId`作为同步白板信令的群组ID。请在开始推流前确保指定的IM群组已创建完成，否则会导致推流失败。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type StartWhiteboardPushRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要推流的白板房间号，取值范围: (1, 4294967295)。
	// 
	// 1. 在没有指定`GroupId`的情况下，白板推流默认以`RoomId`的字符串表达形式作为IM群组ID（比如RoomId为1234，则IM群组ID为"1234"），并加群进行信令同步，请在开始推流前确保相应IM群组已创建完成，否则会导致推流失败。
	// 2. 在没有指定`TRTCRoomId`和`TRTCRoomIdStr`的情况下，默认会以`RoomId`作为白板流进行推流的TRTC房间号。
	RoomId *int64 `json:"RoomId,omitnil,omitempty" name:"RoomId"`

	// 用于白板推流服务进入白板房间的用户ID。在没有额外指定`IMAuthParam`和`TRTCAuthParam`的情况下，这个用户ID同时会用于IM登录、IM加群、TRTC进房推流等操作。
	// 用户ID最大长度不能大于60个字节，该用户ID必须是一个单独的未同时在其他地方使用的用户ID，白板推流服务使用这个用户ID进入房间进行白板音视频推流，若该用户ID和其他地方同时在使用的用户ID重复，会导致白板推流服务与其他使用场景账号互踢，影响正常推流。
	PushUserId *string `json:"PushUserId,omitnil,omitempty" name:"PushUserId"`

	// 与`PushUserId`对应的IM签名(usersig)。
	PushUserSig *string `json:"PushUserSig,omitnil,omitempty" name:"PushUserSig"`

	// 白板参数，例如白板宽高、背景颜色等
	Whiteboard *Whiteboard `json:"Whiteboard,omitnil,omitempty" name:"Whiteboard"`

	// 自动停止推流超时时间，单位秒，取值范围[300, 259200], 默认值为1800秒。
	// 
	// 当白板超过设定时间没有操作的时候，白板推流服务会自动停止白板推流。
	AutoStopTimeout *int64 `json:"AutoStopTimeout,omitnil,omitempty" name:"AutoStopTimeout"`

	// 对主白板推流任务进行操作时，是否同时同步操作备份任务
	AutoManageBackup *bool `json:"AutoManageBackup,omitnil,omitempty" name:"AutoManageBackup"`

	// 备份白板推流相关参数。
	// 
	// 指定了备份参数的情况下，白板推流服务会在房间内新增一路白板画面视频流，即同一个房间内会有两路白板画面推流。
	Backup *WhiteboardPushBackupParam `json:"Backup,omitnil,omitempty" name:"Backup"`

	// TRTC高级权限控制参数，如果在实时音视频开启了高级权限控制功能，必须提供PrivateMapKey才能保证正常推流。
	PrivateMapKey *string `json:"PrivateMapKey,omitnil,omitempty" name:"PrivateMapKey"`

	// 白板推流视频帧率，取值范围[0, 30]，默认20fps
	VideoFPS *int64 `json:"VideoFPS,omitnil,omitempty" name:"VideoFPS"`

	// 白板推流码率， 取值范围[0, 2000]，默认1200kbps。
	// 
	// 这里的码率设置是一个参考值，实际推流的时候使用的是动态码率，所以真实码率不会固定为指定值，会在指定值附近波动。
	VideoBitrate *int64 `json:"VideoBitrate,omitnil,omitempty" name:"VideoBitrate"`

	// 在实时音视频云端录制模式选择为 `指定用户录制` 模式的时候是否自动录制白板推流。
	// 
	// 默认在实时音视频的云端录制模式选择为 `指定用户录制` 模式的情况下，不会自动进行白板推流录制，如果希望进行白板推流录制，请将此参数设置为true。
	// 
	// 如果实时音视频的云端录制模式选择为 `全局自动录制` 模式，可忽略此参数。
	AutoRecord *bool `json:"AutoRecord,omitnil,omitempty" name:"AutoRecord"`

	// 指定白板推流这路流在音视频云端录制中的RecordID，指定的RecordID会用于填充实时音视频云端录制完成后的回调消息中的 "userdefinerecordid" 字段内容，便于您更方便的识别录制回调，以及在点播媒体资源管理中查找相应的录制视频文件。
	// 
	// 限制长度为64字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
	// 
	// 此字段设置后，不管`AutoRecord`字段取值如何，都将自动进行白板推流录制。
	// 
	// 默认RecordId生成规则如下：
	// urlencode(SdkAppID_RoomID_PushUserID)
	// 
	// 例如：
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// 那么：RecordId = 12345678_12345_push_user_1
	UserDefinedRecordId *string `json:"UserDefinedRecordId,omitnil,omitempty" name:"UserDefinedRecordId"`

	// 在实时音视频旁路推流模式选择为`指定用户旁路`模式的时候，是否自动旁路白板推流。
	// 
	// 默认在实时音视频的旁路推流模式选择为 `指定用户旁路` 模式的情况下，不会自动旁路白板推流，如果希望旁路白板推流，请将此参数设置为true。
	// 
	// 如果实时音视频的旁路推流模式选择为 `全局自动旁路` 模式，可忽略此参数。
	AutoPublish *bool `json:"AutoPublish,omitnil,omitempty" name:"AutoPublish"`

	// 指定实时音视频在旁路白板推流这路流时的StreamID，设置之后，您就可以在腾讯云直播 CDN 上通过标准直播方案（FLV或HLS）播放该用户的音视频流。
	// 
	// 限制长度为64字节，只允许包含大小写英文字母（a-zA-Z）、数字（0-9）及下划线和连词符。
	// 
	// 此字段设置后，不管`AutoPublish`字段取值如何，都将自动旁路白板推流。
	// 
	// 默认StreamID生成规则如下：
	// urlencode(SdkAppID_RoomID_PushUserID_main)
	// 
	// 例如：
	// SdkAppID = 12345678，RoomID = 12345，PushUserID = push_user_1
	// 那么：StreamID = 12345678_12345_push_user_1_main
	UserDefinedStreamId *string `json:"UserDefinedStreamId,omitnil,omitempty" name:"UserDefinedStreamId"`

	// 内部参数，不需要关注此参数
	ExtraData *string `json:"ExtraData,omitnil,omitempty" name:"ExtraData"`

	// TRTC数字类型房间号，取值范围: (1, 4294967295)。
	// 
	// 在同时指定了`RoomId`与`TRTCRoomId`的情况下，优先使用`TRTCRoomId`作为白板流进行推流的TRTC房间号。
	// 
	// 当指定了`TRTCRoomIdStr`的情况下，此字段将被忽略。
	TRTCRoomId *int64 `json:"TRTCRoomId,omitnil,omitempty" name:"TRTCRoomId"`

	// TRTC字符串类型房间号。
	// 
	// 在指定了`TRTCRoomIdStr`的情况下，会优先使用`TRTCRoomIdStr`作为白板流进行推流的TRTC房间号。
	TRTCRoomIdStr *string `json:"TRTCRoomIdStr,omitnil,omitempty" name:"TRTCRoomIdStr"`

	// IM鉴权信息参数，用于IM鉴权。
	// 当白板信令所使用的IM应用与白板应用的SdkAppId不一致时，可以通过此参数提供对应IM应用鉴权信息。
	// 
	// 如果提供了此参数，白板推流服务会优先使用此参数指定的SdkAppId作为白板信令的传输通道，否则使用公共参数中的SdkAppId作为白板信令的传输通道。
	IMAuthParam *AuthParam `json:"IMAuthParam,omitnil,omitempty" name:"IMAuthParam"`

	// TRTC鉴权信息参数，用于TRTC进房推流鉴权。
	// 当需要推流到的TRTC房间所对应的TRTC应用与白板应用的SdkAppId不一致时，可以通过此参数提供对应的TRTC应用鉴权信息。
	// 
	// 如果提供了此参数，白板推流服务会优先使用此参数指定的SdkAppId作为白板推流的目标TRTC应用，否则使用公共参数中的SdkAppId作为白板推流的目标TRTC应用。
	TRTCAuthParam *AuthParam `json:"TRTCAuthParam,omitnil,omitempty" name:"TRTCAuthParam"`

	// 指定白板推流时推流用户进TRTC房间的进房模式。默认为 TRTCAppSceneVideoCall
	// 
	// TRTCAppSceneVideoCall - 视频通话场景，即绝大多数时间都是两人或两人以上视频通话的场景，内部编码器和网络协议优化侧重流畅性，降低通话延迟和卡顿率。
	// TRTCAppSceneLIVE - 直播场景，即绝大多数时间都是一人直播，偶尔有多人视频互动的场景，内部编码器和网络协议优化侧重性能和兼容性，性能和清晰度表现更佳。
	TRTCEnterRoomMode *string `json:"TRTCEnterRoomMode,omitnil,omitempty" name:"TRTCEnterRoomMode"`

	// 白板进行信令同步的 IM 群组 ID。
	// 在没有指定`GroupId`的情况下，白板推流服务将使用 `RoomId` 的字符串形式作为同步白板信令的IM群组ID。
	// 在指定了`GroupId`的情况下，白板推流将优先`GroupId`作为同步白板信令的群组ID。请在开始推流前确保指定的IM群组已创建完成，否则会导致推流失败。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *StartWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartWhiteboardPushRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "PushUserId")
	delete(f, "PushUserSig")
	delete(f, "Whiteboard")
	delete(f, "AutoStopTimeout")
	delete(f, "AutoManageBackup")
	delete(f, "Backup")
	delete(f, "PrivateMapKey")
	delete(f, "VideoFPS")
	delete(f, "VideoBitrate")
	delete(f, "AutoRecord")
	delete(f, "UserDefinedRecordId")
	delete(f, "AutoPublish")
	delete(f, "UserDefinedStreamId")
	delete(f, "ExtraData")
	delete(f, "TRTCRoomId")
	delete(f, "TRTCRoomIdStr")
	delete(f, "IMAuthParam")
	delete(f, "TRTCAuthParam")
	delete(f, "TRTCEnterRoomMode")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartWhiteboardPushRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartWhiteboardPushResponseParams struct {
	// 推流任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 备份任务结果参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Backup *string `json:"Backup,omitnil,omitempty" name:"Backup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *StartWhiteboardPushResponseParams `json:"Response"`
}

func (r *StartWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartWhiteboardPushResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopOnlineRecordRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要停止录制的任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopOnlineRecordRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要停止录制的任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopOnlineRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOnlineRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopOnlineRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopOnlineRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopOnlineRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopOnlineRecordResponseParams `json:"Response"`
}

func (r *StopOnlineRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOnlineRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWhiteboardPushRequestParams struct {
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要停止的白板推流任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopWhiteboardPushRequest struct {
	*tchttp.BaseRequest
	
	// 客户的SdkAppId
	SdkAppId *int64 `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 需要停止的白板推流任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopWhiteboardPushRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWhiteboardPushRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWhiteboardPushRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWhiteboardPushResponseParams struct {
	// 备份任务相关参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Backup *string `json:"Backup,omitnil,omitempty" name:"Backup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopWhiteboardPushResponse struct {
	*tchttp.BaseResponse
	Response *StopWhiteboardPushResponseParams `json:"Response"`
}

func (r *StopWhiteboardPushResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWhiteboardPushResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamControl struct {
	// 视频流ID
	// 视频流ID的取值含义如下：
	// 1. tic_record_user - 表示白板视频流
	// 2. tic_substream - 表示辅路视频流
	// 3. 特定用户ID - 表示指定用户的视频流
	// 
	// 在实际录制过程中，视频流ID的匹配规则为前缀匹配，只要真实流ID的前缀与指定的流ID一致就认为匹配成功。
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// 设置是否对此路流开启录制。
	// 
	// true - 表示不对这路流进行录制，录制结果将不包含这路流的视频。
	// false - 表示需要对这路流进行录制，录制结果会包含这路流的视频。
	// 
	// 默认为 false。
	DisableRecord *bool `json:"DisableRecord,omitnil,omitempty" name:"DisableRecord"`

	// 设置是否禁用这路流的音频录制。
	// 
	// true - 表示不对这路流的音频进行录制，录制结果里这路流的视频将会没有声音。
	// false - 录制视频会保留音频，如果设置为true，则录制视频会丢弃这路流的音频。
	// 
	// 默认为 false。
	DisableAudio *bool `json:"DisableAudio,omitnil,omitempty" name:"DisableAudio"`

	// 设置当前流录制视频是否只录制小画面。
	// 
	// true - 录制小画面。设置为true时，请确保上行端同时上行了小画面，否则录制视频可能是黑屏。
	// false - 录制大画面。
	// 
	// 默认为 false。
	PullSmallVideo *bool `json:"PullSmallVideo,omitnil,omitempty" name:"PullSmallVideo"`
}

type StreamLayout struct {
	// 流布局配置参数
	LayoutParams *LayoutParams `json:"LayoutParams,omitnil,omitempty" name:"LayoutParams"`

	// 视频流ID
	// 流ID的取值含义如下：
	// 1. tic_record_user - 表示当前画面用于显示白板视频流
	// 2. tic_substream - 表示当前画面用于显示辅路视频流
	// 3. 特定用户ID - 表示当前画面用于显示指定用户的视频流
	// 4. 不填 - 表示当前画面用于备选，当有新的视频流加入时，会从这些备选的空位中选择一个没有被占用的位置来显示新的视频流画面
	InputStreamId *string `json:"InputStreamId,omitnil,omitempty" name:"InputStreamId"`

	// 背景颜色，默认为黑色，格式为RGB格式，如红色为"#FF0000"
	BackgroundColor *string `json:"BackgroundColor,omitnil,omitempty" name:"BackgroundColor"`

	// 视频画面填充模式。
	// 
	// 0 - 自适应模式，对视频画面进行等比例缩放，在指定区域内显示完整的画面。此模式可能存在黑边。
	// 1 - 全屏模式，对视频画面进行等比例缩放，让画面填充满整个指定区域。此模式不会存在黑边，但会将超出区域的那一部分画面裁剪掉。
	FillMode *int64 `json:"FillMode,omitnil,omitempty" name:"FillMode"`
}

type VideoInfo struct {
	// 视频开始播放的时间（单位：毫秒）
	VideoPlayTime *int64 `json:"VideoPlayTime,omitnil,omitempty" name:"VideoPlayTime"`

	// 视频大小（字节）
	VideoSize *int64 `json:"VideoSize,omitnil,omitempty" name:"VideoSize"`

	// 视频格式
	VideoFormat *string `json:"VideoFormat,omitnil,omitempty" name:"VideoFormat"`

	// 视频播放时长（单位：毫秒）
	VideoDuration *int64 `json:"VideoDuration,omitnil,omitempty" name:"VideoDuration"`

	// 视频文件URL
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 视频文件Id
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// 视频流类型 
	// - 0：摄像头视频 
	// - 1：屏幕分享视频
	// - 2：白板视频 
	// - 3：混流视频
	// - 4：纯音频（mp3)
	VideoType *int64 `json:"VideoType,omitnil,omitempty" name:"VideoType"`

	// 摄像头/屏幕分享视频所属用户的 Id（白板视频为空、混流视频tic_mixstream_房间号_混流布局类型、辅路视频tic_substream_用户Id）
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 视频分辨率的宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 视频分辨率的高
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type Whiteboard struct {
	// 实时录制结果里白板视频宽，取值必须大于等于2，默认为1280
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 实时录制结果里白板视频高，取值必须大于等于2，默认为960
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 白板初始化参数，透传到白板 SDK
	InitParam *string `json:"InitParam,omitnil,omitempty" name:"InitParam"`
}

type WhiteboardPushBackupParam struct {
	// 用于白板推流服务进房的用户ID，
	// 该ID必须是一个单独的未在SDK中使用的ID，白板推流服务将使用这个用户ID进入房间进行白板推流，若该ID和SDK中使用的ID重复，会导致SDK和录制服务互踢，影响正常推流。
	PushUserId *string `json:"PushUserId,omitnil,omitempty" name:"PushUserId"`

	// 与PushUserId对应的签名
	PushUserSig *string `json:"PushUserSig,omitnil,omitempty" name:"PushUserSig"`
}