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

package v20201229

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateImageModerationAsyncTaskRequestParams struct {
	// 接收审核信息回调地址，审核过程中产生的所有结果发送至此地址。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 该字段表示策略的具体编号，用于接口调度，在内容安全控制台中可配置。若不传入Biztype参数（留空），则代表采用默认的识别策略；传入则会在审核时根据业务场景采取不同的审核策略。<br>备注：Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段表示待检测图片文件内容的Base64编码，图片**大小不超过10MB**，建议**分辨率不低于256x256**，否则可能会影响识别效果。<br>备注： **该字段与FileUrl必须选择输入其中一个**。
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 该字段表示待检测图片文件的访问链接，图片支持PNG、JPG、JPEG、BMP、GIF、WEBP格式，**大小不超过100MB**，建议**分辨率不低于256x256**；图片下载时间限制为3秒，超过则会返回下载超时；由于网络安全策略，**送审带重定向的链接，可能引起下载失败**，请尽量避免，比如Http返回302状态码的链接，可能导致接口返回ResourceUnavailable.ImageDownloadError。<br>备注：**该字段与FileContent必须选择输入其中一个**。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// **GIF/长图检测专用**，用于表示GIF截帧频率（每隔多少张图片抽取一帧进行检测），长图则按照长边：短边取整计算要切割的总图数；默认值为0，此时只会检测GIF的第一帧或对长图不进行切分处理。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF/长图时，将每间隔2帧检测一次且最多检测400帧。
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// **GIF/长图检测专用**，用于标识最大截帧数量；默认值为1，此时只会检测输入GIF的第一帧或对长图不进行切分处理（可能会造成处理超时）。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF/长图时，将每间隔2帧检测一次且最多检测400帧。
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// 该字段表示待检测对象对应的用户相关信息，若填入则可甄别相应违规风险用户。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段表示待检测对象对应的设备相关信息，若填入则可甄别相应违规风险设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`
}

type CreateImageModerationAsyncTaskRequest struct {
	*tchttp.BaseRequest
	
	// 接收审核信息回调地址，审核过程中产生的所有结果发送至此地址。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 该字段表示策略的具体编号，用于接口调度，在内容安全控制台中可配置。若不传入Biztype参数（留空），则代表采用默认的识别策略；传入则会在审核时根据业务场景采取不同的审核策略。<br>备注：Biztype仅为数字、字母与下划线的组合，长度为3-32个字符；不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段表示待检测图片文件内容的Base64编码，图片**大小不超过10MB**，建议**分辨率不低于256x256**，否则可能会影响识别效果。<br>备注： **该字段与FileUrl必须选择输入其中一个**。
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 该字段表示待检测图片文件的访问链接，图片支持PNG、JPG、JPEG、BMP、GIF、WEBP格式，**大小不超过100MB**，建议**分辨率不低于256x256**；图片下载时间限制为3秒，超过则会返回下载超时；由于网络安全策略，**送审带重定向的链接，可能引起下载失败**，请尽量避免，比如Http返回302状态码的链接，可能导致接口返回ResourceUnavailable.ImageDownloadError。<br>备注：**该字段与FileContent必须选择输入其中一个**。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// **GIF/长图检测专用**，用于表示GIF截帧频率（每隔多少张图片抽取一帧进行检测），长图则按照长边：短边取整计算要切割的总图数；默认值为0，此时只会检测GIF的第一帧或对长图不进行切分处理。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF/长图时，将每间隔2帧检测一次且最多检测400帧。
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// **GIF/长图检测专用**，用于标识最大截帧数量；默认值为1，此时只会检测输入GIF的第一帧或对长图不进行切分处理（可能会造成处理超时）。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF/长图时，将每间隔2帧检测一次且最多检测400帧。
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// 该字段表示待检测对象对应的用户相关信息，若填入则可甄别相应违规风险用户。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段表示待检测对象对应的设备相关信息，若填入则可甄别相应违规风险设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`
}

func (r *CreateImageModerationAsyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageModerationAsyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CallbackUrl")
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "Interval")
	delete(f, "MaxFrames")
	delete(f, "User")
	delete(f, "Device")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageModerationAsyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageModerationAsyncTaskResponseParams struct {
	// 该字段用于返回检测对象对应请求参数中的DataId。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImageModerationAsyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageModerationAsyncTaskResponseParams `json:"Response"`
}

func (r *CreateImageModerationAsyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageModerationAsyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Device struct {
	// 该字段表示业务用户对应设备的IP地址，同时**支持IPv4和IPv6**地址的记录；需要与IpType参数配合使用。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 该字段表示业务用户对应的MAC地址，以方便设备识别与管理；其格式与取值与标准MAC地址一致。
	Mac *string `json:"Mac,omitnil,omitempty" name:"Mac"`

	// *内测中，敬请期待。*
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// *内测中，敬请期待。*
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 该字段表示业务用户对应设备的**IMEI码**（国际移动设备识别码），该识别码可用于识别每一部独立的手机等移动通信设备，方便设备识别与管理。<br>备注：格式为**15-17位纯数字**。
	IMEI *string `json:"IMEI,omitnil,omitempty" name:"IMEI"`

	// **iOS设备专用**，该字段表示业务用户对应的**IDFA**(广告标识符),这是由苹果公司提供的用于标识用户的广告标识符，由一串16进制的32位数字和字母组成。<br>
	// 备注：苹果公司自2021年iOS14更新后允许用户手动关闭或者开启IDFA，故此字符串标记有效性可能有所降低。
	IDFA *string `json:"IDFA,omitnil,omitempty" name:"IDFA"`

	// **iOS设备专用**，该字段表示业务用户对应的**IDFV**(应用开发商标识符),这是由苹果公司提供的用于标注应用开发商的标识符，由一串16进制的32位数字和字母组成，可被用于唯一标识设备。
	IDFV *string `json:"IDFV,omitnil,omitempty" name:"IDFV"`

	// 该字段表示记录的IP地址的类型，取值：**0**（代表IPv4地址）、**1**（代表IPv6地址）；需要与IpType参数配合使用。
	IpType *uint64 `json:"IpType,omitnil,omitempty" name:"IpType"`
}

// Predefined struct for user
type ImageModerationRequestParams struct {
	// 该字段表示使用的策略的具体编号，该字段需要先在[内容安全控制台](https://console.cloud.tencent.com/cms/clouds/manage)中配置。
	// 备注：不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段表示待检测图片文件内容的Base64编码，由于云API对请求包体有大小限制，图片的**Base64编码内容大小不得超过10MB**。<br/>备注：**该字段与FileUrl必须选择输入其中一个**。
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 该字段表示待检测图片文件的访问链接，URL源图**大小不超过30MB**。<br />备注：该字段与FileContent必须选择输入其中一个。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// **GIF检测专用**，用于表示GIF截帧频率（每隔多少张图片抽取一帧进行检测）；默认值为0，此时只会检测GIF的第一帧或不进行切分处理。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF时，将每间隔2帧检测一次且最多检测400帧。
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// **GIF检测专用**，用于标识最大截帧数量；默认值为1，此时只会检测输入GIF的第一帧不进行切分处理（可能会造成处理超时）。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF时，将每间隔2帧检测一次且最多检测400帧。
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// 该字段表示待检测对象对应的用户相关信息，若填入则可甄别相应违规风险用户。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段表示待检测对象对应的设备相关信息，若填入则可甄别相应违规风险设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// 该字段表示输入的图片审核类型，取值含：IMAGE（内容安全）、IMAGE_AIGC（AI生成识别）两种，默认值为IMAGE。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest
	
	// 该字段表示使用的策略的具体编号，该字段需要先在[内容安全控制台](https://console.cloud.tencent.com/cms/clouds/manage)中配置。
	// 备注：不同Biztype关联不同的业务场景与识别能力策略，调用前请确认正确的Biztype。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段表示您为待检测对象分配的数据ID，传入后可方便您对文件进行标识和管理。<br>取值：由英文字母（大小写均可）、数字及四个特殊符号（_，-，@，#）组成，**长度不超过64个字符**。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段表示待检测图片文件内容的Base64编码，由于云API对请求包体有大小限制，图片的**Base64编码内容大小不得超过10MB**。<br/>备注：**该字段与FileUrl必须选择输入其中一个**。
	FileContent *string `json:"FileContent,omitnil,omitempty" name:"FileContent"`

	// 该字段表示待检测图片文件的访问链接，URL源图**大小不超过30MB**。<br />备注：该字段与FileContent必须选择输入其中一个。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// **GIF检测专用**，用于表示GIF截帧频率（每隔多少张图片抽取一帧进行检测）；默认值为0，此时只会检测GIF的第一帧或不进行切分处理。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF时，将每间隔2帧检测一次且最多检测400帧。
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// **GIF检测专用**，用于标识最大截帧数量；默认值为1，此时只会检测输入GIF的第一帧不进行切分处理（可能会造成处理超时）。<br>备注：Interval与MaxFrames参数需要组合使用。例如，Interval=3, MaxFrames=400，则代表在检测GIF时，将每间隔2帧检测一次且最多检测400帧。
	MaxFrames *int64 `json:"MaxFrames,omitnil,omitempty" name:"MaxFrames"`

	// 该字段表示待检测对象对应的用户相关信息，若填入则可甄别相应违规风险用户。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 该字段表示待检测对象对应的设备相关信息，若填入则可甄别相应违规风险设备。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// 该字段表示输入的图片审核类型，取值含：IMAGE（内容安全）、IMAGE_AIGC（AI生成识别）两种，默认值为IMAGE。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "BizType")
	delete(f, "DataId")
	delete(f, "FileContent")
	delete(f, "FileUrl")
	delete(f, "Interval")
	delete(f, "MaxFrames")
	delete(f, "User")
	delete(f, "Device")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageModerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageModerationResponseParams struct {
	// 该字段用于返回Label标签下的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果（LabelResults）中所对应的**优先级最高的恶意标签**，表示模型推荐的审核结果，建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值标签示例：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；（说明：文档仅示例了部分风险类型，更多返回类型请以实际值为准或[提交工单](https://console.cloud.tencent.com/workorder/category)进行咨询）。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回检测结果所命中优先级最高的恶意标签下的子标签名称，如：*色情--性行为*；若未命中任何子标签则返回空字符串。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回当前标签（Label）下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表图片越有可能属于当前返回的标签；如：*色情 99*，则表明该图片非常有可能属于色情内容；*色情 0*，则表明该图片不属于色情内容。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回分类检测模型的详细检测结果；表示模型推荐的审核结果，建议您按照业务所需，对不同违规类型与建议值进行处理。详细返回值信息可参阅对应的数据结构（LabelResults）描述。
	// 
	// 返回值标签示例：Normal:正常，Porn:色情，Abuse:谩骂，Ad:广告（说明：文档仅示例了部分风险类型，更多返回类型请以实际值为准或[提交工单](https://console.cloud.tencent.com/workorder/category)进行咨询）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelResults []*LabelResult `json:"LabelResults,omitnil,omitempty" name:"LabelResults"`

	// 该字段用于返回物体检测模型的详细检测结果；包括：实体、广告台标、二维码等内容命中的标签名称、标签分数、坐标信息、场景识别结果、建议操作等内容审核信息；详细返回值信息可参阅对应的数据结构（ObjectResults）描述。
	ObjectResults []*ObjectResult `json:"ObjectResults,omitnil,omitempty" name:"ObjectResults"`

	// 该字段用于返回OCR文本识别的详细检测结果；包括：文本坐标信息、文本识别结果、建议操作等内容审核信息；详细返回值信息可参阅对应的数据结构（OcrResults）描述。
	OcrResults []*OcrResult `json:"OcrResults,omitnil,omitempty" name:"OcrResults"`

	// 该字段用于返回基于图片风险库（风险黑库与正常白库）识别的结果,详细返回值信息可参阅对应的数据结构（LibResults）描述。<br>备注：图片风险库目前**暂不支持自定义库**。
	LibResults []*LibResult `json:"LibResults,omitnil,omitempty" name:"LibResults"`

	// 该字段用于返回检测对象对应请求参数中的DataId。
	DataId *string `json:"DataId,omitnil,omitempty" name:"DataId"`

	// 该字段用于返回检测对象对应请求参数中的BizType。
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 该字段用于返回根据您的需求配置的额外附加信息（Extra），如未配置则默认返回值为空。<br>备注：不同客户或Biztype下返回信息不同，如需配置该字段请提交工单咨询或联系售后专员处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 该字段用于返回检测对象对应的MD5校验值，以方便校验文件完整性。
	FileMD5 *string `json:"FileMD5,omitnil,omitempty" name:"FileMD5"`

	// 该字段用于返回仅识别图片元素的模型结果；包括：场景模型命中的标签、置信度和位置信息
	RecognitionResults []*RecognitionResult `json:"RecognitionResults,omitnil,omitempty" name:"RecognitionResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type LabelDetailItem struct {
	// 该字段用于返回识别对象的ID以方便识别和区分。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 该字段用于返回时命中的子标签名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 该字段用于返回对应子标签命中的分值，取值为**0-100**，如：*Porn-SexBehavior 99* 则代表相应识别内容命中色情-性行为标签的分值为99。
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type LabelResult struct {
	// 该字段用于返回模型识别出的场景结果，如广告、色情、有害内容等场景。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 该字段用于返回针对当前恶意标签的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回对应恶意标签下的子标签的检测结果，如：*Porn-SexBehavior*等子标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回当前标签（Label）下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表图片越有可能属于当前返回的标签；如：*色情 99*，则表明该图片非常有可能属于色情内容；*色情 0*，则表明该图片不属于色情内容。
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回分类模型命中子标签的详细信息，如：序号、命中标签名称、分数等信息。
	Details []*LabelDetailItem `json:"Details,omitnil,omitempty" name:"Details"`
}

type LibDetail struct {
	// 该字段用于返回识别对象的ID以方便识别和区分。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 该字段用于返回自定义库的ID，以方便自定义库管理和配置。
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 该字段用于返回自定义库的名称,以方便自定义库管理和配置。
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 该字段用于返回识别图像对象的ID以方便文件管理。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回其他自定义标签以满足您的定制化场景需求，若无需求则可略过。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 该字段用于返回对应模型命中的分值，取值为**0-100**，如：*Porn 99* 则代表相应识别内容命中色情标签的分值为99。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type LibResult struct {
	// 该字段表示模型的场景识别结果，默认取值为Similar。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 该字段用于返回后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回恶意标签下对应的子标签的检测结果，如：*Porn-SexBehavior*等子标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回图片检索模型识别的分值，取值为**0-100**，表示该审核图片**与库中样本的相似分值**，得分越高，代表当前内容越有可能命中相似图库内的样本。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回黑白库比对结果的详细信息，如：序号、库名称、恶意标签等信息；详细返回信息敬请参考对应数据结构（[LibDetail](https://cloud.tencent.com/document/product/1125/53274#LibDetail)）的描述文档
	Details []*LibDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type Location struct {
	// 该参数用于返回检测框**左上角位置的横坐标**（x）所在的像素位置，结合剩余参数可唯一确定检测框的大小和位置。
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// 该参数用于返回检测框**左上角位置的纵坐标**（y）所在的像素位置，结合剩余参数可唯一确定检测框的大小和位置。
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 该参数用于返回**检测框的宽度**（由左上角出发在x轴向右延伸的长度），结合剩余参数可唯一确定检测框的大小和位置。
	Width *float64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 该参数用于返回**检测框的高度**（由左上角出发在y轴向下延伸的长度），结合剩余参数可唯一确定检测框的大小和位置。
	Height *float64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 该参数用于返回**检测框的旋转角度**，该参数结合X和Y两个坐标参数可唯一确定检测框的具体位置；取值：**0-360**（**角度制**），方向为**逆时针旋转**。
	Rotate *float64 `json:"Rotate,omitnil,omitempty" name:"Rotate"`
}

type ObjectDetail struct {
	// 该参数用于返回识别对象的ID以方便识别和区分。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 该参数用于返回命中的实体标签。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 该参数用于返回对应实体标签所对应的值或内容。如：当标签为*二维码(QrCode)*时，该字段为识别出的二维码对应的URL地址。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 该参数用于返回对应实体标签命中的分值，取值为**0-100**，如：*QrCode 99* 则代表相应识别内容命中二维码场景标签的概率非常高。
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回实体检测框的坐标位置（左上角xy坐标、长宽、旋转角度）以方便快速定位实体的相关信息。
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// 该参数用于返回命中的实体二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该参数用于返回命中的人脸id
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`
}

type ObjectResult struct {
	// 该字段用于返回实体识别出的实体场景结果，如二维码、logo、图片OCR等场景。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 该字段用于返回针对当前恶意标签的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回检测结果所对应的恶意标签，表示模型推荐的审核结果，建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回当前恶意标签下对应的子标签的检测结果，如：*Porn-SexBehavior* 等子标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回命中当前恶意标签下子标签的分值，取值为**0-100**，如：*Porn-SexBehavior 99* 则代表相应识别内容命中色情-性行为标签的分值为99。
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该标签用于返回所识别出的实体名称。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 该标签用于返回所识别出实体的详细信息，如：序号、命中标签名称、位置坐标等信息，详细返回内容敬请参考相应数据结构（[ObjectDetail
	// ](https://cloud.tencent.com/document/api/1125/53274#ObjectDetail)）。
	Details []*ObjectDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type OcrHitInfo struct {
	// 标识模型命中还是关键词命中
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 命中关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 自定义词库名称
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 位置信息
	Positions []*Positions `json:"Positions,omitnil,omitempty" name:"Positions"`
}

type OcrResult struct {
	// 该字段表示识别场景，取值默认为OCR（图片OCR识别）。
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 该字段用于返回优先级最高的恶意标签对应的后续操作建议。当您获取到判定结果后，返回值表示系统推荐的后续操作；建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Block**：建议屏蔽，**Review** ：建议人工复审，**Pass**：建议通过
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该字段用于返回OCR检测结果所对应的优先级最高的恶意标签，表示模型推荐的审核结果，建议您按照业务所需，对不同违规类型与建议值进行处理。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回当前标签（Label）下对应的子标签的检测结果，如：*Porn-SexBehavior*等子标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 该字段用于返回当前标签（Label）下的置信度，取值范围：0（**置信度最低**）-100（**置信度最高** ），越高代表文本越有可能属于当前返回的标签；如：*色情 99*，则表明该文本非常有可能属于色情内容；*色情 0*，则表明该文本不属于色情内容。
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该字段用于返回OCR识别出的结果的详细内容，如：文本内容、对应标签、识别框位置等信息。
	Details []*OcrTextDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 该字段用于返回OCR识别出的文字信息。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type OcrTextDetail struct {
	// 该字段用于返回OCR识别出的文本内容。<br>备注：OCR文本识别上限在**5000字节内**。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 该字段用于返回检测结果所对应的恶意标签。<br>返回值：**Normal**：正常，**Porn**：色情，**Abuse**：谩骂，**Ad**：广告；以及其他令人反感、不安全或不适宜的内容类型。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 该字段用于返回自定义库的ID，以方便自定义库管理和配置。
	LibId *string `json:"LibId,omitnil,omitempty" name:"LibId"`

	// 该字段用于返回自定义库的名称，以方便自定义库管理和配置。
	LibName *string `json:"LibName,omitnil,omitempty" name:"LibName"`

	// 该参数用于返回在当前label下命中的关键词。
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 该参数用于返回在当前恶意标签下模型命中的分值，取值为**0-100**；分数越高，代表当前场景越符合该恶意标签所对应的场景。
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 该参数用于返回OCR检测框在图片中的位置（左上角xy坐标、长宽、旋转角度），以方便快速定位识别文字的相关信息。
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// 该参数用于返回OCR文本识别结果的置信度，取值在**0**（**置信度最低**）-**100**（**置信度最高**），越高代表对应图像越有可能是识别出的文字；如：*你好 99*，则表明OCR识别框内的文字大概率是”你好“。
	Rate *uint64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 该字段用于返回检测结果所对应的恶意二级标签。
	SubLabel *string `json:"SubLabel,omitnil,omitempty" name:"SubLabel"`

	// 关键词命中位置信息
	HitInfos []*OcrHitInfo `json:"HitInfos,omitnil,omitempty" name:"HitInfos"`
}

type Positions struct {
	// 关键词起始位置
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// 关键词结束位置
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type RecognitionResult struct {
	// 当前可能的取值：Scene（图片场景模型）
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Label对应模型下的识别标签信息
	Tags []*RecognitionTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RecognitionTag struct {
	// 标签名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 置信分：0～100，数值越大表示置信度越高
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 标签位置信息，若模型无位置信息，则可能为零值
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`
}

type User struct {
	// 该字段表示业务用户ID,填写后，系统可根据账号过往违规历史优化审核结果判定，有利于存在可疑违规风险时的辅助判断。<br>
	// 备注：该字段可传入微信openid、QQopenid、字符串等账号信息，与账号类别参数（AccountType）配合使用可确定唯一账号。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 该字段表示业务用户对应的账号昵称信息。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 该字段表示业务用户ID对应的账号类型，取值：**1**-微信uin，**2**-QQ号，**3**-微信群uin，**4**-qq群号，**5**-微信openid，**6**-QQopenid，**7**-其它string。<br>
	// 该字段与账号ID参数（UserId）配合使用可确定唯一账号。
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 该字段表示业务用户对应账号的性别信息。<br>
	// 取值：**0**（默认值，代表性别未知）、**1**（男性）、**2**（女性）。
	Gender *uint64 `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 该字段表示业务用户对应账号的年龄信息。<br>
	// 取值：**0**（默认值，代表年龄未知）-（**自定义年龄上限**）之间的整数。
	Age *uint64 `json:"Age,omitnil,omitempty" name:"Age"`

	// 该字段表示业务用户对应账号的等级信息。<br>
	// 取值：**0**（默认值，代表等级未知）、**1**（等级较低）、**2**（等级中等）、**3**（等级较高），目前**暂不支持自定义等级**。
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 该字段表示业务用户对应账号的手机号信息，支持全球各地区手机号的记录。<br>
	// 备注：请保持手机号格式的统一，如区号格式（086/+86）等。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 该字段表示业务用户的简介信息，支持汉字、英文及特殊符号，**长度不超过5000个汉字字符**。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 该字段表示业务用户头像图片的访问链接(URL)，支持PNG、JPG、JPEG、BMP、GIF、WEBP格式。<br>备注：头像图片**大小不超过5MB**，建议**分辨率不低于256x256**；图片下载时间限制为3秒，超过则会返回下载超时。
	HeadUrl *string `json:"HeadUrl,omitnil,omitempty" name:"HeadUrl"`
}