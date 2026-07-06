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

package v20240523

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CameraControl struct {
	// 枚举值：“simple”, “down_back”, “forward_up”, “right_turn_forward”, “left_turn_forward”
	// simple：简单运镜，此类型下可在"config"中六选一进行运镜
	// down_back：镜头下压并后退 -> 下移拉远，此类型下config参数无需填写
	// forward_up：镜头前进并上仰 -> 推进上移，此类型下config参数无需填写
	// right_turn_forward：先右旋转后前进 -> 右旋推进，此类型下config参数无需填写
	// left_turn_forward：先左旋并前进 -> 左旋推进，此类型下config参数无需填写
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 包含六个字段，用于指定摄像机在不同方向上的运动或变化。
	// - 当运镜类型指定simple时必填，指定其他类型时不填
	// - 参数6选1，即只能有一个参数不为0，其余参数为0
	Config *CameraControlConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type CameraControlConfig struct {
	// 水平运镜，控制摄像机在水平方向上的移动量（沿x轴平移）
	// 
	// 取值范围：[-10, 10]，负值表示向左平移，正值表示向右平移
	Horizontal *float64 `json:"Horizontal,omitnil,omitempty" name:"Horizontal"`

	// 垂直运镜，控制摄像机在垂直方向上的移动量（沿y轴平移）
	// 
	// 取值范围：[-10, 10]，负值表示向下平移，正值表示向上平移
	Vertical *float64 `json:"Vertical,omitnil,omitempty" name:"Vertical"`

	// 水平摇镜，控制摄像机在水平面上的旋转量（绕y轴旋转）
	// 
	// 取值范围：[-10, 10]，负值表示绕y轴向左旋转，正值表示绕y轴向右旋转
	Pan *float64 `json:"Pan,omitnil,omitempty" name:"Pan"`

	// 垂直摇镜，控制摄像机在垂直面上的旋转量（沿x轴旋转）
	// 
	// 取值范围：[-10, 10]，负值表示绕x轴向下旋转，正值表示绕x轴向上旋转
	Tilt *float64 `json:"Tilt,omitnil,omitempty" name:"Tilt"`

	// 旋转运镜，控制摄像机的滚动量（绕z轴旋转）
	// 
	// 取值范围：[-10, 10]，负值表示绕z轴逆时针旋转，正值表示绕z轴顺时针旋转
	Roll *float64 `json:"Roll,omitnil,omitempty" name:"Roll"`

	// 变焦，控制摄像机的焦距变化，影响视野的远近
	// 
	// 取值范围：[-10, 10]，负值表示焦距变长、视野范围变小，正值表示焦距变短、视野范围变大
	Zoom *float64 `json:"Zoom,omitnil,omitempty" name:"Zoom"`
}

// Predefined struct for user
type CreateAigcElementRequestParams struct {
	// <p>主体名称<br>不能超过20个字符</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>主体描述</p><p>不能超过100个字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>主体参考方式</p><p>枚举值：<br>video_refer<br>image_refer<br>video_refer: 视频角色主体，此时将参考element_video_list定义主体外表<br>image_refer: 多图主体，此时将参考element_image_list定义主体外表</p>
	ReferenceType *string `json:"ReferenceType,omitnil,omitempty" name:"ReferenceType"`

	// <p>主体参考图，可通过多张图片设定主体及其细节</p><p>包括正面参考图和其他角度或特写参考图，其中：至少包括1张正面参考图，由frontal_image参数定义；需包括1～3张其他参考图，需与正面参考图有差异，由image_url参数定义<br>支持传入图片Base64编码或图片URL（确保可访问）</p><p>图片格式支持.jpg / .jpeg / .png。图片文件大小不能超过10MB，图片宽高尺寸不小于300px，图片宽高比要在1:2.5 ~ 2.5:1之间</p><p>reference_type参数值为 image_refer 时，当前参数必填</p>
	ElementImageList *ElementImageList `json:"ElementImageList,omitnil,omitempty" name:"ElementImageList"`

	// <p>主体参考视频，可通过视频设定主体及其细节</p><p>可上传有声视频，有声视频包含人声则触发音色定制（定制+入音色库+与主体绑定）</p><p>暂时仅支持通过视频定制写实风格的人形形象</p><p>参考视频时当前参数必填，参考图片时当前参数无效</p><p>用key:value承载。视频格式仅支持MP4/MOV。仅支持时长介于3s～8s之间、宽高比例需为16:9或9:16的1080P视频。至多仅支持上传1段视频，视频大小不超过200MB。video_url参数值不得为空</p>
	VideoList []*string `json:"VideoList,omitnil,omitempty" name:"VideoList"`

	// <p>厂商</p>
	Provider []*string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>为主体配置标签，一个主体可以配置多个标签</p><p>用key:value承载。tag的ID与名称：o_101 热梗, o_102 人物, o_103 动物, o_104 道具, o_105 服饰, o_106 场景, o_107 特效, o_108 其他</p>
	TagList []*TagList `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>主体音色ID，可绑定音色库中已有音色</p><p>当前参数为空时，当前主体不绑定音色<br>为多图主体绑定音色时，仅支持人物形象主体或类人形象主体</p>
	ElementVoiceId *string `json:"ElementVoiceId,omitnil,omitempty" name:"ElementVoiceId"`
}

type CreateAigcElementRequest struct {
	*tchttp.BaseRequest
	
	// <p>主体名称<br>不能超过20个字符</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>主体描述</p><p>不能超过100个字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>主体参考方式</p><p>枚举值：<br>video_refer<br>image_refer<br>video_refer: 视频角色主体，此时将参考element_video_list定义主体外表<br>image_refer: 多图主体，此时将参考element_image_list定义主体外表</p>
	ReferenceType *string `json:"ReferenceType,omitnil,omitempty" name:"ReferenceType"`

	// <p>主体参考图，可通过多张图片设定主体及其细节</p><p>包括正面参考图和其他角度或特写参考图，其中：至少包括1张正面参考图，由frontal_image参数定义；需包括1～3张其他参考图，需与正面参考图有差异，由image_url参数定义<br>支持传入图片Base64编码或图片URL（确保可访问）</p><p>图片格式支持.jpg / .jpeg / .png。图片文件大小不能超过10MB，图片宽高尺寸不小于300px，图片宽高比要在1:2.5 ~ 2.5:1之间</p><p>reference_type参数值为 image_refer 时，当前参数必填</p>
	ElementImageList *ElementImageList `json:"ElementImageList,omitnil,omitempty" name:"ElementImageList"`

	// <p>主体参考视频，可通过视频设定主体及其细节</p><p>可上传有声视频，有声视频包含人声则触发音色定制（定制+入音色库+与主体绑定）</p><p>暂时仅支持通过视频定制写实风格的人形形象</p><p>参考视频时当前参数必填，参考图片时当前参数无效</p><p>用key:value承载。视频格式仅支持MP4/MOV。仅支持时长介于3s～8s之间、宽高比例需为16:9或9:16的1080P视频。至多仅支持上传1段视频，视频大小不超过200MB。video_url参数值不得为空</p>
	VideoList []*string `json:"VideoList,omitnil,omitempty" name:"VideoList"`

	// <p>厂商</p>
	Provider []*string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>为主体配置标签，一个主体可以配置多个标签</p><p>用key:value承载。tag的ID与名称：o_101 热梗, o_102 人物, o_103 动物, o_104 道具, o_105 服饰, o_106 场景, o_107 特效, o_108 其他</p>
	TagList []*TagList `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>主体音色ID，可绑定音色库中已有音色</p><p>当前参数为空时，当前主体不绑定音色<br>为多图主体绑定音色时，仅支持人物形象主体或类人形象主体</p>
	ElementVoiceId *string `json:"ElementVoiceId,omitnil,omitempty" name:"ElementVoiceId"`
}

func (r *CreateAigcElementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAigcElementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ReferenceType")
	delete(f, "ElementImageList")
	delete(f, "VideoList")
	delete(f, "Provider")
	delete(f, "TagList")
	delete(f, "ElementVoiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAigcElementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAigcElementResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// <p>主体Id</p>
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// <p>任务状态</p><p>默认值：任务状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>厂商</p>
	Provider []*string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>任务创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAigcElementResponse struct {
	*tchttp.BaseResponse
	Response *CreateAigcElementResponseParams `json:"Response"`
}

func (r *CreateAigcElementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAigcElementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAigcElementRequestParams struct {

	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`
}

type DeleteAigcElementRequest struct {
	*tchttp.BaseRequest
	
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`
}

func (r *DeleteAigcElementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAigcElementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ElementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAigcElementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAigcElementResponseParams struct {

	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`


	Deleted *bool `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAigcElementResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAigcElementResponseParams `json:"Response"`
}

func (r *DeleteAigcElementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAigcElementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAigcElementRequestParams struct {
	// <p>主体Id</p>
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`
}

type DescribeAigcElementRequest struct {
	*tchttp.BaseRequest
	
	// <p>主体Id</p>
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`
}

func (r *DescribeAigcElementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAigcElementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ElementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAigcElementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAigcElementResponseParams struct {
	// <p>主体名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>主体id</p>
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`

	// <p>主体描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>主体参考方式</p><p>枚举值：</p><ul><li>video_refer： 视频角色主体</li><li>image_refer： 多图主体</li></ul>
	ReferenceType *string `json:"ReferenceType,omitnil,omitempty" name:"ReferenceType"`

	// <p>任务状态</p><p>枚举值：</p><ul><li>pending： 执行中</li><li>failed： 任务失败</li><li>succeed： 任务成功</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>厂商列表</p>
	Provider []*string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>主体参考图，可通过多张图片设定主体及其细节</p>
	ElementImageList *ElementImageList `json:"ElementImageList,omitnil,omitempty" name:"ElementImageList"`

	// <p>主体参考视频，可通过视频设定主体及其细节</p>
	VideoList []*string `json:"VideoList,omitnil,omitempty" name:"VideoList"`

	// <p>为主体配置标签，一个主体可以配置多个标签</p>
	TagList []*TagList `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>厂商聚合字段</p>
	ProviderDetails []*ProviderDetail `json:"ProviderDetails,omitnil,omitempty" name:"ProviderDetails"`

	// <p>创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>更新时间</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// <p>音色Id</p>
	ElementVoiceId *string `json:"ElementVoiceId,omitnil,omitempty" name:"ElementVoiceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAigcElementResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAigcElementResponseParams `json:"Response"`
}

func (r *DescribeAigcElementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAigcElementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHumanActorJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeHumanActorJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeHumanActorJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHumanActorJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHumanActorJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHumanActorJobResponseParams struct {
	// 任务状态。  WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 结果视频URL。有效期 24 小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 任务执行错误码。当任务状态不为 FAIL 时，该值为""。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 任务执行错误信息。当任务状态不为 FAIL 时，该值为""。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHumanActorJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHumanActorJobResponseParams `json:"Response"`
}

func (r *DescribeHumanActorJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHumanActorJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHunyuanToVideoJobRequestParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeHunyuanToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeHunyuanToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHunyuanToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHunyuanToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHunyuanToVideoJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行错误码。当任务状态不为 FAIL 时，该值为""。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 任务执行错误信息。当任务状态不为 FAIL 时，该值为""。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 结果视频 URL。有效期 24 小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHunyuanToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHunyuanToVideoJobResponseParams `json:"Response"`
}

func (r *DescribeHunyuanToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHunyuanToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageToVideoGeneralJobRequestParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeImageToVideoGeneralJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeImageToVideoGeneralJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageToVideoGeneralJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageToVideoGeneralJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageToVideoGeneralJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行错误码。当任务状态不为 FAIL 时，该值为""。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 任务执行错误信息。当任务状态不为 FAIL 时，该值为""。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 结果视频 URL。有效期 24 小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageToVideoGeneralJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageToVideoGeneralJobResponseParams `json:"Response"`
}

func (r *DescribeImageToVideoGeneralJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageToVideoGeneralJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageToVideoJobRequestParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

type DescribeImageToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

func (r *DescribeImageToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ExternalTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageToVideoJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>视频id</p>
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// <p>视频总时长，单位s</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>任务最终扣减积分数值</p>
	FinalUnitDeduction *string `json:"FinalUnitDeduction,omitnil,omitempty" name:"FinalUnitDeduction"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageToVideoJobResponseParams `json:"Response"`
}

func (r *DescribeImageToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageToVideoViduJobRequestParams struct {
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeImageToVideoViduJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeImageToVideoViduJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageToVideoViduJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageToVideoViduJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageToVideoViduJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>该任务消耗的积分数量。</p>
	Credits *float64 `json:"Credits,omitnil,omitempty" name:"Credits"`

	// <p>任务调用时传入的透传参数。</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageToVideoViduJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageToVideoViduJobResponseParams `json:"Response"`
}

func (r *DescribeImageToVideoViduJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageToVideoViduJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMotionControlKlingJobRequestParams struct {
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

type DescribeMotionControlKlingJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

func (r *DescribeMotionControlKlingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMotionControlKlingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ExternalTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMotionControlKlingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMotionControlKlingJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>视频时长。</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>消耗积分数。</p>
	FinalUnitDeduction *string `json:"FinalUnitDeduction,omitnil,omitempty" name:"FinalUnitDeduction"`

	// <p>视频id</p>
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMotionControlKlingJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMotionControlKlingJobResponseParams `json:"Response"`
}

func (r *DescribeMotionControlKlingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMotionControlKlingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortraitSingJobRequestParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribePortraitSingJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribePortraitSingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortraitSingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePortraitSingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortraitSingJobResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务状态码
	// —RUN：处理中
	// —FAIL：处理失败
	// —STOP：处理终止
	// —DONE：处理完成
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 任务状态信息
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// 任务执行错误码。当任务状态不为FAIL时，该值为""。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 任务执行错误信息。当任务状态不为FAIL时，该值为""。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 生成视频的URL地址。有效期24小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePortraitSingJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribePortraitSingJobResponseParams `json:"Response"`
}

func (r *DescribePortraitSingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortraitSingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReferenceToVideoViduJobRequestParams struct {
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeReferenceToVideoViduJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeReferenceToVideoViduJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReferenceToVideoViduJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReferenceToVideoViduJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReferenceToVideoViduJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>该任务消耗的积分数量。</p>
	Credits *float64 `json:"Credits,omitnil,omitempty" name:"Credits"`

	// <p>任务调用时传入的透传参数。</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReferenceToVideoViduJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReferenceToVideoViduJobResponseParams `json:"Response"`
}

func (r *DescribeReferenceToVideoViduJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReferenceToVideoViduJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateToVideoJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeTemplateToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeTemplateToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateToVideoJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行错误码。当任务状态不为 FAIL 时，该值为""。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 任务执行错误信息。当任务状态不为 FAIL 时，该值为""。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 结果视频 URL。有效期 24 小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTemplateToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateToVideoJobResponseParams `json:"Response"`
}

func (r *DescribeTemplateToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextToVideoJobRequestParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

type DescribeTextToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

func (r *DescribeTextToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ExternalTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTextToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextToVideoJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>视频id</p>
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// <p>视频总时长，单位s</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>任务最终扣减积分数值</p>
	FinalUnitDeduction *string `json:"FinalUnitDeduction,omitnil,omitempty" name:"FinalUnitDeduction"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTextToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTextToVideoJobResponseParams `json:"Response"`
}

func (r *DescribeTextToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextToVideoViduJobRequestParams struct {
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeTextToVideoViduJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeTextToVideoViduJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextToVideoViduJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTextToVideoViduJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextToVideoViduJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>该任务消耗的积分数量。</p>
	Credits *float64 `json:"Credits,omitnil,omitempty" name:"Credits"`

	// <p>任务调用时传入的透传参数。</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTextToVideoViduJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTextToVideoViduJobResponseParams `json:"Response"`
}

func (r *DescribeTextToVideoViduJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextToVideoViduJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoEditKlingJobRequestParams struct {
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

type DescribeVideoEditKlingJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

func (r *DescribeVideoEditKlingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoEditKlingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ExternalTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoEditKlingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoEditKlingJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>视频id</p>
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// <p>时长</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>消耗积分数</p>
	FinalUnitDeduction *string `json:"FinalUnitDeduction,omitnil,omitempty" name:"FinalUnitDeduction"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoEditKlingJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoEditKlingJobResponseParams `json:"Response"`
}

func (r *DescribeVideoEditKlingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoEditKlingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoExtendKlingJobRequestParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

type DescribeVideoExtendKlingJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

func (r *DescribeVideoExtendKlingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoExtendKlingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ExternalTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoExtendKlingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoExtendKlingJobResponseParams struct {
	// <p>任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>任务执行错误码。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// <p>任务执行错误信息。当任务状态不为 FAIL 时，该值为&quot;&quot;。</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// <p>结果视频 URL。有效期 24 小时。</p>
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// <p>视频时长。</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>消耗积分数。</p>
	FinalUnitDeduction *string `json:"FinalUnitDeduction,omitnil,omitempty" name:"FinalUnitDeduction"`

	// <p>视频id</p>
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoExtendKlingJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoExtendKlingJobResponseParams `json:"Response"`
}

func (r *DescribeVideoExtendKlingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoExtendKlingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoFaceFusionJobRequestParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeVideoFaceFusionJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeVideoFaceFusionJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoFaceFusionJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoFaceFusionJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoFaceFusionJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行错误码。当任务状态不为 FAIL 时，该值为""。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 任务执行错误信息。当任务状态不为 FAIL 时，该值为""。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 结果视频 URL。有效期 24 小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoFaceFusionJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoFaceFusionJobResponseParams `json:"Response"`
}

func (r *DescribeVideoFaceFusionJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoFaceFusionJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DynamicMask struct {
	// <p>动态笔刷涂抹区域（用户通过运动笔刷涂抹的 mask 图片）</p><p>支持传入图片Base64编码或图片URL（确保可访问，格式要求同 Image 字段）<br>图片格式支持.jpg / .jpeg / .png<br>图片长宽比必须与输入图片相同（即Image字段），否则任务失败<br>StaticMask 和 DynamicMasks.Mask 这两张图片的分辨率必须一致，否则任务失败</p>
	Mask *string `json:"Mask,omitnil,omitempty" name:"Mask"`

	// <p>运动轨迹坐标序列</p><p>生成5s的视频，轨迹长度不超过77，即坐标个数取值范围：[2, 77]<br>轨迹坐标系，以图片左下角为坐标原点<br>注1：坐标点个数越多轨迹刻画越准确，如只有2个轨迹点则为这两点连接的直线<br>注2：轨迹方向以传入顺序为指向，以最先传入的坐标为轨迹起点，依次连接后续坐标形成运动轨迹</p>
	Trajectories []*Trajectory `json:"Trajectories,omitnil,omitempty" name:"Trajectories"`
}

type Element struct {
	// <p>ID配置</p>
	ElementId *string `json:"ElementId,omitnil,omitempty" name:"ElementId"`
}

type ElementImageList struct {
	// <p>主体参考图</p>
	ReferImages []*ReferImageItem `json:"ReferImages,omitnil,omitempty" name:"ReferImages"`

	// <p>主体主图</p>
	FrontalImage *string `json:"FrontalImage,omitnil,omitempty" name:"FrontalImage"`
}

type ExtraParam struct {
	// <p>预签名的上传url，支持把视频直接传到客户指定的地址。</p>
	UserDesignatedUrl *string `json:"UserDesignatedUrl,omitnil,omitempty" name:"UserDesignatedUrl"`

	// <p>回调地址<br>需要您在创建任务时主动设置 CallbackUrl，请求方法为 POST，当视频生成结束时，我们将向此地址发送生成结果。<br>数据格式如下：<br>{<br>    &quot;JobId&quot;: &quot;1397428070633955328&quot;,<br>    &quot;Status&quot;: &quot;DONE&quot;,<br>    &quot;ErrorCode&quot;: &quot;&quot;,<br>    &quot;ErrorMessage&quot;: &quot;&quot;,<br>    &quot;ResultVideoUrl&quot;: &quot;https://vcg.cos.tencentcos.cn/template_to_video/fa80b846-b933-4981-afad-8a39b46ef2ca.mp4&quot;<br>}</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>BGM音频文本。</p>
	BGMText *string `json:"BGMText,omitnil,omitempty" name:"BGMText"`
}

type FaceMergeInfo struct {
	// 融合图片
	MergeFaceImage *Image `json:"MergeFaceImage,omitnil,omitempty" name:"MergeFaceImage"`

	// 上传的图片人脸位置信息（人脸框）
	// Width、Height >= 30。
	MergeFaceRect *FaceRect `json:"MergeFaceRect,omitnil,omitempty" name:"MergeFaceRect"`

	// 素材人脸ID，不填默认取上传图片中最大人脸。
	TemplateFaceID *string `json:"TemplateFaceID,omitnil,omitempty" name:"TemplateFaceID"`
}

type FaceRect struct {
	// <p>人脸框左上角横坐标。</p>
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// <p>人脸框左上角纵坐标。</p>
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// <p>人脸框宽度。<br>单位：px</p>
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// <p>人脸框高度。<br>单位：px</p>
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type FaceTemplateInfo struct {
	// 角色ID。需要与MergeInfos中的TemplateFaceID依次对应。需要填数字，建议填"0"、"1"，依次累加。
	TemplateFaceID *string `json:"TemplateFaceID,omitnil,omitempty" name:"TemplateFaceID"`

	// 视频模板中要替换的人脸图片
	TemplateFaceImage *Image `json:"TemplateFaceImage,omitnil,omitempty" name:"TemplateFaceImage"`

	// 视频模板中要替换的人脸图片的人脸框。不填默认取要替换的人脸图片中最大人脸。
	TemplateFaceRect *FaceRect `json:"TemplateFaceRect,omitnil,omitempty" name:"TemplateFaceRect"`
}

type Image struct {
	// 图片Base64
	Base64 *string `json:"Base64,omitnil,omitempty" name:"Base64"`

	// 图片Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ImageInfo struct {
	// 图片URL
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 首帧：first_frame
	// 尾帧：end_frame
	// 其他：空
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type LogoParam struct {
	// 水印 Url
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 水印 Base64，Url 和 Base64 二选一传入，如果都提供以 Url 为准
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`

	// 水印图片位于生成结果图中的坐标及宽高，将按照坐标对标识图片进行位置和大小的拉伸匹配。
	LogoRect *LogoRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`
}

type LogoRect struct {
	// <p>水印图框X坐标值。当值大于0时，坐标轴原点位于原图左侧，方向指右；当值小于0时，坐标轴原点位于原图右侧，方向指左。</p>
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// <p>水印图框Y坐标值。当值大于0时，坐标轴原点位于原图上侧，方向指下；当值小于0时，坐标轴原点位于原图下侧，方向指上。</p>
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// <p>水印图框宽度。<br>单位：px</p>
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// <p>水印图框高度。<br>单位：px</p>
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type MultiPrompt struct {
	// <p>分镜序号</p>
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>分镜提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>时长</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type ProviderDetail struct {
	// <p>供应商详情</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>错误信息</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type ReferImageItem struct {
	// <p>图片地址</p>
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`
}

type ReferVideoInfo struct {
	// 视频地址
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 视频类型
	// feature为特征参考视频
	// base为待编辑视频
	ReferType *string `json:"ReferType,omitnil,omitempty" name:"ReferType"`

	// 否保留视频原声，yes为保留，no为不保留；
	// 当前参数对特征参考视频（feature）也生效。
	KeepOriginalSound *string `json:"KeepOriginalSound,omitnil,omitempty" name:"KeepOriginalSound"`
}

type ReferenceSubject struct {
	// <p>主体id，后续生成时在提示词中可以通过@主体id的方式使用。</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>该主体对应的图片url，每个主体最多支持3张图片<br>注1：支持传入图片URL（确保可访问）<br>注2：图片支持 png、jpeg、jpg、webp格式<br>注3：图片像素不能小于 128*128，且比例需要小于1:4或者4:1，且大小不超过50M。</p>
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>主体id，后续生成时可以通过@主体id的方式使用</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>主体视频，该主体对应的视频url，与videos必填一个<br>注1： 仅viduq2-pro模型支持使用视频主体<br>注2：每个主体中的图片和视频，共享3个槽位<br>注3：支持1个5秒视频<br>注4：支持传入视频 URL（确保可访问）<br>注5：视频支持 mp4、avi、mov格式<br>注6：视频像素不能小于 128*128，且比例需要小于1:4或者4:1<br>注7：请注意，base64 decode之后的字节长度需要小于20M，且编码必须包含适当的内容类型字符串</p>
	Videos []*string `json:"Videos,omitnil,omitempty" name:"Videos"`

	// <p>音色ID用来决定视频中的声音音色，为空时系统会自动推荐，可选枚举值参考列表：[音色列表](URL https://shengshu.feishu.cn/sheets/EgFvs6DShhiEBStmjzccr5gonOg)</p>
	VoiceId *string `json:"VoiceId,omitnil,omitempty" name:"VoiceId"`
}

// Predefined struct for user
type SubmitHumanActorJobRequestParams struct {
	// 文本提示词，不能超过5000字符。
	// 提示词支持全局和局部控制：
	// 
	// - 全局控制：正常输入提示词即可
	// - 局部控制：可用双井号进行特定时间的提示词约束，例如： "画面中的人物正在对着镜头讲话，偶尔做些手势匹配说话的内容。镜头保持固定。#3#画面中的人物正在对着镜头讲话，同时做出单手做向左方引导的手势。镜头保持固定。"（意思是第三秒的时候让人物做出左方引导手势）
	//  -- 局部控制时间建议整数，最大可读小数点后两位。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 传入音频URL地址，音频要求：
	// - 音频时长：2秒 - 60秒
	// - 音频格式：mp3、wav
	// - 音频大小：10M以内
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 传入图片URL地址，图片要求：
	// - 图片格式：jpg、jpeg、png、bmp、webp
	// - 图片分辨率：192～4096
	// - 图片大小：不超过10M
	// - 图片宽高比：图片【宽：高】在1:4到4:1范围内
	// - 图片内容：避免上传无人脸、无宠物脸或脸部过小、不完整、不清晰、偏转角度过大、嘴部被遮挡的图片。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 传入图片Base64编码，编码后请求体大小不超过10M。
	// 图片Base64编码与URL地址必传其一，如果都传以ImageUrl为准。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 生成视频分辨率
	// 枚举值：720p，1080p
	// 默认1080p
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 生成视频帧数，单位fps。
	// 枚举值：25，50
	// 默认50帧
	FrameRate *int64 `json:"FrameRate,omitnil,omitempty" name:"FrameRate"`

	// 为生成视频添加标识的开关，默认为1。 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitHumanActorJobRequest struct {
	*tchttp.BaseRequest
	
	// 文本提示词，不能超过5000字符。
	// 提示词支持全局和局部控制：
	// 
	// - 全局控制：正常输入提示词即可
	// - 局部控制：可用双井号进行特定时间的提示词约束，例如： "画面中的人物正在对着镜头讲话，偶尔做些手势匹配说话的内容。镜头保持固定。#3#画面中的人物正在对着镜头讲话，同时做出单手做向左方引导的手势。镜头保持固定。"（意思是第三秒的时候让人物做出左方引导手势）
	//  -- 局部控制时间建议整数，最大可读小数点后两位。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 传入音频URL地址，音频要求：
	// - 音频时长：2秒 - 60秒
	// - 音频格式：mp3、wav
	// - 音频大小：10M以内
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 传入图片URL地址，图片要求：
	// - 图片格式：jpg、jpeg、png、bmp、webp
	// - 图片分辨率：192～4096
	// - 图片大小：不超过10M
	// - 图片宽高比：图片【宽：高】在1:4到4:1范围内
	// - 图片内容：避免上传无人脸、无宠物脸或脸部过小、不完整、不清晰、偏转角度过大、嘴部被遮挡的图片。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 传入图片Base64编码，编码后请求体大小不超过10M。
	// 图片Base64编码与URL地址必传其一，如果都传以ImageUrl为准。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 生成视频分辨率
	// 枚举值：720p，1080p
	// 默认1080p
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 生成视频帧数，单位fps。
	// 枚举值：25，50
	// 默认50帧
	FrameRate *int64 `json:"FrameRate,omitnil,omitempty" name:"FrameRate"`

	// 为生成视频添加标识的开关，默认为1。 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitHumanActorJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHumanActorJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "AudioUrl")
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "Resolution")
	delete(f, "FrameRate")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHumanActorJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHumanActorJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHumanActorJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHumanActorJobResponseParams `json:"Response"`
}

func (r *SubmitHumanActorJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHumanActorJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanToVideoJobRequestParams struct {
	// 视频内容的描述，中文正向提示词。最多支持200个 utf-8 字符（首尾空格不计入字符数）。 示例值：一只猫在草原上奔跑，写实风格
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图片
	// 上传图url大小不超过 10M，base64不超过8M。
	// 支持jpg，png，jpeg，webp，bmp，tiff 格式
	// 单边分辨率不超过5000，不小于50，长宽限制1:4 ~ 4:1
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 目前仅支持720p视频分辨率，默认720p。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。
	//  1：添加标识； 0：不添加标识； 其他数值：默认按1处理。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitHunyuanToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频内容的描述，中文正向提示词。最多支持200个 utf-8 字符（首尾空格不计入字符数）。 示例值：一只猫在草原上奔跑，写实风格
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图片
	// 上传图url大小不超过 10M，base64不超过8M。
	// 支持jpg，png，jpeg，webp，bmp，tiff 格式
	// 单边分辨率不超过5000，不小于50，长宽限制1:4 ~ 4:1
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 目前仅支持720p视频分辨率，默认720p。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。
	//  1：添加标识； 0：不添加标识； 其他数值：默认按1处理。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitHunyuanToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "Image")
	delete(f, "Resolution")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanToVideoJobResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanToVideoJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageToVideoGeneralJobRequestParams struct {
	// 输入图片
	// Base64 和 Url 必须提供一个，如果都提供以ImageUrl为准。
	// 上传图url大小不超过 8M
	// 支持jpg，png，jpeg，webp，bmp，tiff 格式
	// 单边分辨率不超过5000，不小于50，长宽限制1:4 ~ 4:1
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 视频内容的描述，中文正向提示词。最多支持200个 utf-8 字符（首尾空格不计入字符数）。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输出视频分辨率。可选择：480p、720p、1080p。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 生成视频的帧率，从16, 24, 30中选择。默认值：30
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。  1：添加标识；  0：不添加标识；  其他数值：默认按1处理。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitImageToVideoGeneralJobRequest struct {
	*tchttp.BaseRequest
	
	// 输入图片
	// Base64 和 Url 必须提供一个，如果都提供以ImageUrl为准。
	// 上传图url大小不超过 8M
	// 支持jpg，png，jpeg，webp，bmp，tiff 格式
	// 单边分辨率不超过5000，不小于50，长宽限制1:4 ~ 4:1
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 视频内容的描述，中文正向提示词。最多支持200个 utf-8 字符（首尾空格不计入字符数）。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输出视频分辨率。可选择：480p、720p、1080p。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 生成视频的帧率，从16, 24, 30中选择。默认值：30
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。  1：添加标识；  0：不添加标识；  其他数值：默认按1处理。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitImageToVideoGeneralJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageToVideoGeneralJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Prompt")
	delete(f, "Resolution")
	delete(f, "Fps")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitImageToVideoGeneralJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageToVideoGeneralJobResponseParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitImageToVideoGeneralJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitImageToVideoGeneralJobResponseParams `json:"Response"`
}

func (r *SubmitImageToVideoGeneralJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageToVideoGeneralJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageToVideoJobRequestParams struct {
	// <p>模型名称。<br>v1.6：Kling-V1-6<br>v2.0：Kling-V2-Master<br>v2.1：Kling-V2-1<br>v2.5：Kling-V2-5-Turbo<br>v2.6：Kling-V2-6<br>V3.0：kling-v3</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>参考图像。</p><ul><li>支持传入图片Base64编码或图片URL（确保可访问）</li><li>图片文件大小不能超过10MB，图片分辨率不小于300*300px，图片宽高比要在1:2.5 ~ 2.5:1之间</li><li>Image 参数与 ImageTail 参数至少二选一，二者不能同时为空</li><li>图片格式支持.jpg / .jpeg / .png</li></ul>
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>参考尾帧图像。</p><ul><li>支持传入图片Base64编码或图片URL（确保可访问）</li><li>图片文件大小不能超过10MB，图片分辨率不小于300*300px，图片宽高比要在1:2.5 ~ 2.5:1之间</li><li>Image 参数与 ImageTail 参数至少二选一，二者不能同时为空</li><li>图片格式支持.jpg / .jpeg / .png</li><li>ImageTail参数、DynamicMasks/StaticMask参数、CameraControl参数三选一，不能同时使用</li></ul>
	ImageTail *Image `json:"ImageTail,omitnil,omitempty" name:"ImageTail"`

	// <p>正向文本提示词。不能超过2500个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>负向文本提示词。不能超过2500个字符</p>
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// <p>生成视频时长，单位s。默认值为5。<br>枚举值：3，4，5，6，7，8，9，10，11，12，13，14，15</p><p>不同模型支持时长不同</p><ul><li>模型v1.6、v2.0、v2.5、v2.6：支持5、10</li><li>模型v3.0：支持3～15s</li></ul>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>生成视频的模式<br>枚举值：std，pro<br>其中std：标准模式（标准），基础模式，性价比高<br>其中pro：专家模式（高品质），高表现模式，生成视频质量更佳</p><p>不同模型版本、视频模式支持范围不同</p><ul><li>v1.6：首尾帧与仅首帧情况下只支持pro</li><li>v2.0、v3.0 ：无需配置</li><li>v2.1、v2.5、v2.6：首尾帧情况下只支持pro</li></ul>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>生成视频的自由度；值越大，模型自由度越小，与用户输入的提示词相关性越强。<br>v2.0、v2.5、v2.6模型的不支持当前参数<br>取值范围：[0, 1]</p>
	CfgScale *float64 `json:"CfgScale,omitnil,omitempty" name:"CfgScale"`

	// <p>生成视频时是否同时生成声音<br>枚举值：on，off<br>不同模型版本、视频模式支持范围不同</p><ul><li>v2.6：有首尾帧的pro模式只能生成无声的视频</li></ul>
	Sound *string `json:"Sound,omitnil,omitempty" name:"Sound"`

	// <p>为生成视频添加标识的开关，默认为1。<br>1：添加标识。<br>0：不添加标识。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>是否生成多镜头视频</p><ul><li>当前参数为true时，Prompt参数无效，且不支持设定首尾帧生视频</li><li>当前参数为false时，ShotType参数及MultiPrompt参数无效</li></ul>
	MultiShot *bool `json:"MultiShot,omitnil,omitempty" name:"MultiShot"`

	// <p>分镜方式</p><ul><li>枚举值：customize，intelligence</li><li>当MultiShot参数为true时，当前参数必填</li></ul>
	ShotType *string `json:"ShotType,omitnil,omitempty" name:"ShotType"`

	// <p>各分镜信息，如提示词、时长等<br>通过Index、Prompt、Duration参数定义分镜序号及相应提示词和时长，其中：</p><ul><li>最多支持6个分镜，最小支持1个分镜</li><li>每个分镜相关内容的最大长度不超过512</li><li>每个分镜的时长不大于当前任务的总时长，不小于1</li><li>所有分镜的时长之和等于当前任务的总时长</li></ul><p>用key:value承载，如下：<br>&quot;MultiPrompt&quot;:[<br>    {<br>      &quot;Index&quot;:int,<br>    &quot;Prompt&quot;: &quot;string&quot;,<br>    &quot;Duration&quot;: &quot;5&quot;<br>  },<br>  {<br>      &quot;Index&quot;:int,<br>    &quot;Prompt&quot;: &quot;string&quot;,<br>    &quot;Duration&quot;: &quot;5&quot;<br>  }<br>]</p><ul><li>当MultiShot参数为true且ShotType参数为customize时，当前参数不得为空</li></ul>
	MultiPrompt []*MultiPrompt `json:"MultiPrompt,omitnil,omitempty" name:"MultiPrompt"`

	// <p>参考主体列表</p><ul><li><p>基于主体库中主体的ID配置，用key:value承载，如下：</p><pre><code>  &quot;ElementList&quot;:[      {         &quot;ElementId&quot;:long    },    {         &quot;ElementId&quot;:long    }  ]</code></pre></li><li><p>最多支持3个参考主体</p></li><li><p>主体分为视频定制主体（简称：视频角色主体）和图片定制主体（简称：多图主体），适用范围不同，请注意区分</p></li><li><p>更多主体信息详见：可灵「主体库 3.0」使用指南</p></li></ul>
	ElementList []*Element `json:"ElementList,omitnil,omitempty" name:"ElementList"`

	// <p>静态笔刷涂抹区域（用户通过运动笔刷涂抹的 mask 图片）<br>“运动笔刷”能力包含“动态笔刷 DynamicMasks”和“静态笔刷 StaticMask”两种<br>支持传入图片Base64编码或图片URL（确保可访问，格式要求同 Image 字段）<br>图片格式支持.jpg / .jpeg / .png<br>图片长宽比必须与输入图片相同（即Image字段），否则任务失败（failed）<br>StaticMask和 DynamicMasks.Mask这两张图片的分辨率必须一致，否则任务失败（failed）</p>
	StaticMask *string `json:"StaticMask,omitnil,omitempty" name:"StaticMask"`

	// <p>动态笔刷配置列表<br>可配置多组（最多6组），每组包含“涂抹区域 Mask”与“运动轨迹 Trajectories”序列</p>
	DynamicMasks []*DynamicMask `json:"DynamicMasks,omitnil,omitempty" name:"DynamicMasks"`

	// <p>控制摄像机运动的协议（如未指定，模型将根据输入的文本/图片进行智能匹配）</p>
	CameraControl *CameraControl `json:"CameraControl,omitnil,omitempty" name:"CameraControl"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>生成视频时所引用的音色的列表</p><p>一次视频生成任务至多引用2个音色<br>当VoiceList参数不为空且Prompt参数中引用音色ID时，视频生成任务按“有指定音色”计量计费<br>VoiceId参数值通过音色定制接口返回，也可使用系统预置音色，详见音色定制相关API；非对口型API的VoiceId<br>ElementList参数与VoiceList参数互斥，不能共存<br>v3模型不支持指定音色<br>用key:value承载，如下：<br>&quot;VoiceList&quot;:[<br>  {&quot;VoiceId&quot;:&quot;VoiceId_1&quot;},<br>  {&quot;VoiceId&quot;:&quot;VoiceId_2&quot;}<br>]</p>
	VoiceList []*Voice `json:"VoiceList,omitnil,omitempty" name:"VoiceList"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

type SubmitImageToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型名称。<br>v1.6：Kling-V1-6<br>v2.0：Kling-V2-Master<br>v2.1：Kling-V2-1<br>v2.5：Kling-V2-5-Turbo<br>v2.6：Kling-V2-6<br>V3.0：kling-v3</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>参考图像。</p><ul><li>支持传入图片Base64编码或图片URL（确保可访问）</li><li>图片文件大小不能超过10MB，图片分辨率不小于300*300px，图片宽高比要在1:2.5 ~ 2.5:1之间</li><li>Image 参数与 ImageTail 参数至少二选一，二者不能同时为空</li><li>图片格式支持.jpg / .jpeg / .png</li></ul>
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>参考尾帧图像。</p><ul><li>支持传入图片Base64编码或图片URL（确保可访问）</li><li>图片文件大小不能超过10MB，图片分辨率不小于300*300px，图片宽高比要在1:2.5 ~ 2.5:1之间</li><li>Image 参数与 ImageTail 参数至少二选一，二者不能同时为空</li><li>图片格式支持.jpg / .jpeg / .png</li><li>ImageTail参数、DynamicMasks/StaticMask参数、CameraControl参数三选一，不能同时使用</li></ul>
	ImageTail *Image `json:"ImageTail,omitnil,omitempty" name:"ImageTail"`

	// <p>正向文本提示词。不能超过2500个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>负向文本提示词。不能超过2500个字符</p>
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// <p>生成视频时长，单位s。默认值为5。<br>枚举值：3，4，5，6，7，8，9，10，11，12，13，14，15</p><p>不同模型支持时长不同</p><ul><li>模型v1.6、v2.0、v2.5、v2.6：支持5、10</li><li>模型v3.0：支持3～15s</li></ul>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>生成视频的模式<br>枚举值：std，pro<br>其中std：标准模式（标准），基础模式，性价比高<br>其中pro：专家模式（高品质），高表现模式，生成视频质量更佳</p><p>不同模型版本、视频模式支持范围不同</p><ul><li>v1.6：首尾帧与仅首帧情况下只支持pro</li><li>v2.0、v3.0 ：无需配置</li><li>v2.1、v2.5、v2.6：首尾帧情况下只支持pro</li></ul>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>生成视频的自由度；值越大，模型自由度越小，与用户输入的提示词相关性越强。<br>v2.0、v2.5、v2.6模型的不支持当前参数<br>取值范围：[0, 1]</p>
	CfgScale *float64 `json:"CfgScale,omitnil,omitempty" name:"CfgScale"`

	// <p>生成视频时是否同时生成声音<br>枚举值：on，off<br>不同模型版本、视频模式支持范围不同</p><ul><li>v2.6：有首尾帧的pro模式只能生成无声的视频</li></ul>
	Sound *string `json:"Sound,omitnil,omitempty" name:"Sound"`

	// <p>为生成视频添加标识的开关，默认为1。<br>1：添加标识。<br>0：不添加标识。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>是否生成多镜头视频</p><ul><li>当前参数为true时，Prompt参数无效，且不支持设定首尾帧生视频</li><li>当前参数为false时，ShotType参数及MultiPrompt参数无效</li></ul>
	MultiShot *bool `json:"MultiShot,omitnil,omitempty" name:"MultiShot"`

	// <p>分镜方式</p><ul><li>枚举值：customize，intelligence</li><li>当MultiShot参数为true时，当前参数必填</li></ul>
	ShotType *string `json:"ShotType,omitnil,omitempty" name:"ShotType"`

	// <p>各分镜信息，如提示词、时长等<br>通过Index、Prompt、Duration参数定义分镜序号及相应提示词和时长，其中：</p><ul><li>最多支持6个分镜，最小支持1个分镜</li><li>每个分镜相关内容的最大长度不超过512</li><li>每个分镜的时长不大于当前任务的总时长，不小于1</li><li>所有分镜的时长之和等于当前任务的总时长</li></ul><p>用key:value承载，如下：<br>&quot;MultiPrompt&quot;:[<br>    {<br>      &quot;Index&quot;:int,<br>    &quot;Prompt&quot;: &quot;string&quot;,<br>    &quot;Duration&quot;: &quot;5&quot;<br>  },<br>  {<br>      &quot;Index&quot;:int,<br>    &quot;Prompt&quot;: &quot;string&quot;,<br>    &quot;Duration&quot;: &quot;5&quot;<br>  }<br>]</p><ul><li>当MultiShot参数为true且ShotType参数为customize时，当前参数不得为空</li></ul>
	MultiPrompt []*MultiPrompt `json:"MultiPrompt,omitnil,omitempty" name:"MultiPrompt"`

	// <p>参考主体列表</p><ul><li><p>基于主体库中主体的ID配置，用key:value承载，如下：</p><pre><code>  &quot;ElementList&quot;:[      {         &quot;ElementId&quot;:long    },    {         &quot;ElementId&quot;:long    }  ]</code></pre></li><li><p>最多支持3个参考主体</p></li><li><p>主体分为视频定制主体（简称：视频角色主体）和图片定制主体（简称：多图主体），适用范围不同，请注意区分</p></li><li><p>更多主体信息详见：可灵「主体库 3.0」使用指南</p></li></ul>
	ElementList []*Element `json:"ElementList,omitnil,omitempty" name:"ElementList"`

	// <p>静态笔刷涂抹区域（用户通过运动笔刷涂抹的 mask 图片）<br>“运动笔刷”能力包含“动态笔刷 DynamicMasks”和“静态笔刷 StaticMask”两种<br>支持传入图片Base64编码或图片URL（确保可访问，格式要求同 Image 字段）<br>图片格式支持.jpg / .jpeg / .png<br>图片长宽比必须与输入图片相同（即Image字段），否则任务失败（failed）<br>StaticMask和 DynamicMasks.Mask这两张图片的分辨率必须一致，否则任务失败（failed）</p>
	StaticMask *string `json:"StaticMask,omitnil,omitempty" name:"StaticMask"`

	// <p>动态笔刷配置列表<br>可配置多组（最多6组），每组包含“涂抹区域 Mask”与“运动轨迹 Trajectories”序列</p>
	DynamicMasks []*DynamicMask `json:"DynamicMasks,omitnil,omitempty" name:"DynamicMasks"`

	// <p>控制摄像机运动的协议（如未指定，模型将根据输入的文本/图片进行智能匹配）</p>
	CameraControl *CameraControl `json:"CameraControl,omitnil,omitempty" name:"CameraControl"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>生成视频时所引用的音色的列表</p><p>一次视频生成任务至多引用2个音色<br>当VoiceList参数不为空且Prompt参数中引用音色ID时，视频生成任务按“有指定音色”计量计费<br>VoiceId参数值通过音色定制接口返回，也可使用系统预置音色，详见音色定制相关API；非对口型API的VoiceId<br>ElementList参数与VoiceList参数互斥，不能共存<br>v3模型不支持指定音色<br>用key:value承载，如下：<br>&quot;VoiceList&quot;:[<br>  {&quot;VoiceId&quot;:&quot;VoiceId_1&quot;},<br>  {&quot;VoiceId&quot;:&quot;VoiceId_2&quot;}<br>]</p>
	VoiceList []*Voice `json:"VoiceList,omitnil,omitempty" name:"VoiceList"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

func (r *SubmitImageToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Image")
	delete(f, "ImageTail")
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Duration")
	delete(f, "Mode")
	delete(f, "CfgScale")
	delete(f, "Sound")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "MultiShot")
	delete(f, "ShotType")
	delete(f, "MultiPrompt")
	delete(f, "ElementList")
	delete(f, "StaticMask")
	delete(f, "DynamicMasks")
	delete(f, "CameraControl")
	delete(f, "CallbackUrl")
	delete(f, "VoiceList")
	delete(f, "ExternalTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitImageToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageToVideoJobResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitImageToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitImageToVideoJobResponseParams `json:"Response"`
}

func (r *SubmitImageToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageToVideoViduJobRequestParams struct {
	// <p>上传单张图时【首帧生视频】：</p><p>模型将以此参数中传入的图片为首帧画面来生成视频。<br>注1：支持传入图片 Base64 编码或图片URL（确保可访问）；<br>注2：支持上传1 张图；<br>注3：图片支持 png、jpeg、jpg、webp格式；<br>注4：图片比例需要小于 1:4 或者 4:1 ；<br>注5：图片大小不超过50M。</p><p>上传两张图时【首尾帧生视频】：<br>上传的第一张图片视作首帧图，第二张图片视作尾帧图，模型将以此参数中传入的图片来生成视频<br>注1: 首尾帧两张输入图的分辨率需相近，首帧图的分辨率/尾帧图的分辨率要在0.8～1.25之间。且图片比例需要小于1:4或者4:1；<br>注2: 支持传入图片 Base64 编码或图片URL（确保可访问）；<br>注3: 图片支持 png、jpeg、jpg、webp格式；<br>注4: 图片大小不超过50M；</p>
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>模型名称，可选值，默认值viduq2-turbo</p><ul><li>viduq2-pro：新模型，效果好，细节丰富</li><li>viduq2-turbo：新模型，效果好，生成快</li><li>viduq3-turbo：对比viduq3-pro，生成速度更快</li><li>viduq3-pro：高效生成优质音视频内容，让视频内容更生动、更形象、更立体，效果更好</li></ul>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>文本提示词<br>生成视频的文本描述。<br>注1：字符长度不能超过 2000 个字符<br>注2：若使用is_rec推荐提示词参数，模型将不考虑此参数所输入的提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>是否使用推荐提示词，默认关闭<br>-true：是，由系统自动推荐提示词，并使用提示词内容生成视频，推荐提示词数量=1<br>-false：否，根据输入的prompt生成视频<br>注意：启用推荐提示词后，每个任务多消耗1积分</p>
	IsRec *bool `json:"IsRec,omitnil,omitempty" name:"IsRec"`

	// <p>【首帧生视频】</p><p>是否使用音视频直出能力，默认关闭，枚举值为：</p><ul><li>false：不需要音视频直出，输出静音视频</li><li>true：需要音视频直出，输出带台词以及背景音的视频<br>注1：该参数为true时，voice_id参数才生效<br>注2：该参数为true时，仅q3模型支持错峰<br>注3：当model 为q3-pro或q3-turbo 时，该参数默认值为true</li></ul><p>【首尾帧生视频】</p><ul><li>false：不需要音视频直出，输出静音视频</li><li>true：需要音画同步，输出声音的视频（包括台词和音效）<br>注1：仅q3系列模型支持该参数</li></ul>
	Audio *bool `json:"Audio,omitnil,omitempty" name:"Audio"`

	// <p>音色id<br>用来决定视频中的声音音色，为空时系统会自动推荐，可选枚举值参考列表：新音色列表<br>暂不支持声音复刻功能</p>
	VoiceId *string `json:"VoiceId,omitnil,omitempty" name:"VoiceId"`

	// <p>视频时长参数：<br>【首帧生视频】：<br>viduq3-pro、viduq3-turboxa0默认为 5，可选：1 - 16<br>viduq2-pro、viduq2-turboxa0默认为 5，可选：1 - 10</p><p>【首尾帧生视频】：<br>viduq3-proxa0、viduq3-turbo默认为 5，可选：1 - 16<br>viduq2-pro、viduq2-turboxa0默认为 5，可选：1 - 8</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>分辨率参数<br>【首帧生视频】：</p><ul><li>viduq3-pro 、viduq3-turbo 1-16秒：默认 720p，可选：540p、720p、1080p</li><li>viduq2-pro、viduq2-turbo 1-10秒：默认 720p，可选：720p、1080p<br>示例值：720p</li></ul><p>【首尾帧生视频】：<br>-xa0viduq3-proxa0、viduq3-turbo1-16秒：默认 720p，可选：540p、720p、1080p<br>-xa0viduq2-proxa0、xa0viduq2-turbo1-8秒：默认 720p，可选：540p、720p、1080p</p>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>运动幅度<br>默认 auto，可选值：auto、small、medium、large<br>注：q2、q3系列模型改参数不生效</p>
	MovementAmplitude *string `json:"MovementAmplitude,omitnil,omitempty" name:"MovementAmplitude"`

	// <p>bgm</p>
	Bgm *bool `json:"Bgm,omitnil,omitempty" name:"Bgm"`

	// <p>音频类型，audio为true时必填，默认为all<br>注：该参数目前仅支持q2、q1、2.0系列模型的音频拆分</p><p>枚举值：</p><ul><li>all： 音效+人声</li><li>speech_only： 仅人声</li><li>sound_effect_only： 仅音效</li></ul>
	AudioType *string `json:"AudioType,omitnil,omitempty" name:"AudioType"`

	// <p>元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：<br>{<br>&quot;Label&quot;: &quot;your_label&quot;,<br>&quot;ContentProducer&quot;: &quot;your_content_producer&quot;,<br>&quot;ContentPropagator&quot;: &quot;your_content_propagator&quot;,<br>&quot;ProduceID&quot;: &quot;your_product_id&quot;,<br>&quot;PropagateID&quot;: &quot;your_propagate_id&quot;,<br>&quot;ReservedCode1&quot;: &quot;your_reserved_code1&quot;,<br>&quot;ReservedCode2&quot;: &quot;your_reserved_code2&quot;<br>}<br>该参数为空时，默认使用vidu生成的元数据标识</p>
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>Callback 协议<br>需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成完成时，将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致<br>回调返回的&quot;status&quot;包括以下状态：</p><ul><li>success 任务完成（如发送失败，回调三次）</li><li>failed 任务失败（如发送失败，回调三次）</li></ul>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>透传参数<br>不做任何处理，仅数据传输<br>注：最多 1048576个字符</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// <p>错峰模式，默认为：false，可选值：</p><ul><li>true：错峰生成视频；</li><li>false：即时生成视频；<br>注1：错峰模式消耗的积分更低<br>注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分</li></ul>
	OffPeak *bool `json:"OffPeak,omitnil,omitempty" name:"OffPeak"`

	// <p>为生成结果图添加显式水印标识的开关，默认为1。<br>1：添加。<br>0：不添加。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。<br>示例值：1</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitImageToVideoViduJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>上传单张图时【首帧生视频】：</p><p>模型将以此参数中传入的图片为首帧画面来生成视频。<br>注1：支持传入图片 Base64 编码或图片URL（确保可访问）；<br>注2：支持上传1 张图；<br>注3：图片支持 png、jpeg、jpg、webp格式；<br>注4：图片比例需要小于 1:4 或者 4:1 ；<br>注5：图片大小不超过50M。</p><p>上传两张图时【首尾帧生视频】：<br>上传的第一张图片视作首帧图，第二张图片视作尾帧图，模型将以此参数中传入的图片来生成视频<br>注1: 首尾帧两张输入图的分辨率需相近，首帧图的分辨率/尾帧图的分辨率要在0.8～1.25之间。且图片比例需要小于1:4或者4:1；<br>注2: 支持传入图片 Base64 编码或图片URL（确保可访问）；<br>注3: 图片支持 png、jpeg、jpg、webp格式；<br>注4: 图片大小不超过50M；</p>
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>模型名称，可选值，默认值viduq2-turbo</p><ul><li>viduq2-pro：新模型，效果好，细节丰富</li><li>viduq2-turbo：新模型，效果好，生成快</li><li>viduq3-turbo：对比viduq3-pro，生成速度更快</li><li>viduq3-pro：高效生成优质音视频内容，让视频内容更生动、更形象、更立体，效果更好</li></ul>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>文本提示词<br>生成视频的文本描述。<br>注1：字符长度不能超过 2000 个字符<br>注2：若使用is_rec推荐提示词参数，模型将不考虑此参数所输入的提示词</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>是否使用推荐提示词，默认关闭<br>-true：是，由系统自动推荐提示词，并使用提示词内容生成视频，推荐提示词数量=1<br>-false：否，根据输入的prompt生成视频<br>注意：启用推荐提示词后，每个任务多消耗1积分</p>
	IsRec *bool `json:"IsRec,omitnil,omitempty" name:"IsRec"`

	// <p>【首帧生视频】</p><p>是否使用音视频直出能力，默认关闭，枚举值为：</p><ul><li>false：不需要音视频直出，输出静音视频</li><li>true：需要音视频直出，输出带台词以及背景音的视频<br>注1：该参数为true时，voice_id参数才生效<br>注2：该参数为true时，仅q3模型支持错峰<br>注3：当model 为q3-pro或q3-turbo 时，该参数默认值为true</li></ul><p>【首尾帧生视频】</p><ul><li>false：不需要音视频直出，输出静音视频</li><li>true：需要音画同步，输出声音的视频（包括台词和音效）<br>注1：仅q3系列模型支持该参数</li></ul>
	Audio *bool `json:"Audio,omitnil,omitempty" name:"Audio"`

	// <p>音色id<br>用来决定视频中的声音音色，为空时系统会自动推荐，可选枚举值参考列表：新音色列表<br>暂不支持声音复刻功能</p>
	VoiceId *string `json:"VoiceId,omitnil,omitempty" name:"VoiceId"`

	// <p>视频时长参数：<br>【首帧生视频】：<br>viduq3-pro、viduq3-turboxa0默认为 5，可选：1 - 16<br>viduq2-pro、viduq2-turboxa0默认为 5，可选：1 - 10</p><p>【首尾帧生视频】：<br>viduq3-proxa0、viduq3-turbo默认为 5，可选：1 - 16<br>viduq2-pro、viduq2-turboxa0默认为 5，可选：1 - 8</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>分辨率参数<br>【首帧生视频】：</p><ul><li>viduq3-pro 、viduq3-turbo 1-16秒：默认 720p，可选：540p、720p、1080p</li><li>viduq2-pro、viduq2-turbo 1-10秒：默认 720p，可选：720p、1080p<br>示例值：720p</li></ul><p>【首尾帧生视频】：<br>-xa0viduq3-proxa0、viduq3-turbo1-16秒：默认 720p，可选：540p、720p、1080p<br>-xa0viduq2-proxa0、xa0viduq2-turbo1-8秒：默认 720p，可选：540p、720p、1080p</p>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>运动幅度<br>默认 auto，可选值：auto、small、medium、large<br>注：q2、q3系列模型改参数不生效</p>
	MovementAmplitude *string `json:"MovementAmplitude,omitnil,omitempty" name:"MovementAmplitude"`

	// <p>bgm</p>
	Bgm *bool `json:"Bgm,omitnil,omitempty" name:"Bgm"`

	// <p>音频类型，audio为true时必填，默认为all<br>注：该参数目前仅支持q2、q1、2.0系列模型的音频拆分</p><p>枚举值：</p><ul><li>all： 音效+人声</li><li>speech_only： 仅人声</li><li>sound_effect_only： 仅音效</li></ul>
	AudioType *string `json:"AudioType,omitnil,omitempty" name:"AudioType"`

	// <p>元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：<br>{<br>&quot;Label&quot;: &quot;your_label&quot;,<br>&quot;ContentProducer&quot;: &quot;your_content_producer&quot;,<br>&quot;ContentPropagator&quot;: &quot;your_content_propagator&quot;,<br>&quot;ProduceID&quot;: &quot;your_product_id&quot;,<br>&quot;PropagateID&quot;: &quot;your_propagate_id&quot;,<br>&quot;ReservedCode1&quot;: &quot;your_reserved_code1&quot;,<br>&quot;ReservedCode2&quot;: &quot;your_reserved_code2&quot;<br>}<br>该参数为空时，默认使用vidu生成的元数据标识</p>
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>Callback 协议<br>需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成完成时，将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致<br>回调返回的&quot;status&quot;包括以下状态：</p><ul><li>success 任务完成（如发送失败，回调三次）</li><li>failed 任务失败（如发送失败，回调三次）</li></ul>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>透传参数<br>不做任何处理，仅数据传输<br>注：最多 1048576个字符</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// <p>错峰模式，默认为：false，可选值：</p><ul><li>true：错峰生成视频；</li><li>false：即时生成视频；<br>注1：错峰模式消耗的积分更低<br>注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分</li></ul>
	OffPeak *bool `json:"OffPeak,omitnil,omitempty" name:"OffPeak"`

	// <p>为生成结果图添加显式水印标识的开关，默认为1。<br>1：添加。<br>0：不添加。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。<br>示例值：1</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitImageToVideoViduJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageToVideoViduJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Images")
	delete(f, "Model")
	delete(f, "Prompt")
	delete(f, "IsRec")
	delete(f, "Audio")
	delete(f, "VoiceId")
	delete(f, "Duration")
	delete(f, "Resolution")
	delete(f, "MovementAmplitude")
	delete(f, "Bgm")
	delete(f, "AudioType")
	delete(f, "MetaData")
	delete(f, "CallbackUrl")
	delete(f, "Payload")
	delete(f, "OffPeak")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitImageToVideoViduJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageToVideoViduJobResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitImageToVideoViduJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitImageToVideoViduJobResponseParams `json:"Response"`
}

func (r *SubmitImageToVideoViduJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageToVideoViduJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitMotionControlKlingJobRequestParams struct {
	// <p>模型名称  枚举值：kling-v2-6, kling-v3。</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>文本提示词，可包含正向描述和负向描述</p><p>可将提示词模板化来满足不同的视频生成需求</p><p>不能超过2500个字</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// <p>参考图像，生成视频中的人物、背景等元素均以参考图为准  视频内容需满足以下要求：  人物比例尽量与参考动作比例一致，尽量避免全身动作驱动半身人物进行生成  人物需要露出清晰的上半身或全身的肢体及头部，避免遮挡  画面中人物避免存在极端朝向，比如倒立、平卧等。人物占画面比例不得太低  支持真实/风格化的角色（包括人物/类人动物/部分纯动物/部分类人肢体比例的角色）通过  包含支持传入图片Base64编码或图片URL（确保可访问）。</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>参考视频的获取链接。生成视频中的人物动作与参考视频一致。  视频内容需满足以下要求：  人物需要漏出清晰的上半身或全身的全部肢体及头部，避免遮挡  建议上传1人动作视频，2人及以上会取画面占比最大的人物动作进行生成  推荐使用真人动作，部分风格化的人物/类人肢体比例可以通过  动作视频一镜到底，角色始终出现在画面中，避免切镜、运镜等。否则会被截取  动作避免过快，相对平稳的动作生成效果更佳  视频文件支持.mp4/.mov，文件大小不超过100MB，仅支持长宽的边长均位于340px~3850px之间，上述校验不通过会返回错误码等信息  视频时长下限不短于3秒，时长上限与人物朝向参考（character_orientation）有关：  当人物朝向与视频中人物一致时，视频时长最长可达30秒；  当人物朝向与图片中人物一致时，视频时长最长可达10秒；  如果您的动作难度比较高、速度比较快，有一定概率生成不足上传视频时长的结果，因为模型只能提取有效动作时长进行生成，最短提取出3s可用连续动作即可生成。请注意，因此消耗的积分将无法退还，建议适当调整动作难度与速度  系统会校验视频内容，如有问题会返回错误码等信息。</p>
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`

	// <p>生成视频的模式</p><p>枚举值：std，pro<br>其中std：标准模式（标准），基础模式，性价比高<br>其中pro：专家模式（高品质），高表现模式，生成视频质量更佳</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>可选择是否保留视频原声  枚举值：yes，no  其中yes：保留视频原声  其中no：不保留视频原声。</p>
	KeepOriginalSound *string `json:"KeepOriginalSound,omitnil,omitempty" name:"KeepOriginalSound"`

	// <p>生成视频中人物的朝向，可选择与图片一致或与视频一致  枚举值：image，video，其中：  其中image：与图片中人物朝向一致；此时参考视频时长不得超过10秒；  其中video：与视频中人物朝向一致；此时参考视频时长不得超过30秒；  引用主体时，生成的视频暂时只能参考视频中的人物朝向。</p>
	CharacterOrientation *string `json:"CharacterOrientation,omitnil,omitempty" name:"CharacterOrientation"`

	// <p>参考主体列表  基于主体库中主体的ID配置，用key:value承载，如下：  &quot;element_list&quot;:[     {        &quot;element_id&quot;:long   },   {        &quot;element_id&quot;:long   } ] 参考主体数量与有无参考视频、参考图片数量有关，其中：  当使用首帧生成视频时，最多支持3个主体；  当使用首尾帧生成视频时，kling-v3-omni模型最多支持3个主体，kling-video-o1模不支持主体；  有参考视频时，参考图片数量和参考主体数量之和不得超过4，且不支持使用视频角色主体；  无参考视频时，参考图片数量和参考主体数量之和不得超过7；</p>
	ElementList []*Element `json:"ElementList,omitnil,omitempty" name:"ElementList"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。 1：添加标识； 0：不添加标识； 其他数值：默认按1处理。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitMotionControlKlingJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型名称  枚举值：kling-v2-6, kling-v3。</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>文本提示词，可包含正向描述和负向描述</p><p>可将提示词模板化来满足不同的视频生成需求</p><p>不能超过2500个字</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// <p>参考图像，生成视频中的人物、背景等元素均以参考图为准  视频内容需满足以下要求：  人物比例尽量与参考动作比例一致，尽量避免全身动作驱动半身人物进行生成  人物需要露出清晰的上半身或全身的肢体及头部，避免遮挡  画面中人物避免存在极端朝向，比如倒立、平卧等。人物占画面比例不得太低  支持真实/风格化的角色（包括人物/类人动物/部分纯动物/部分类人肢体比例的角色）通过  包含支持传入图片Base64编码或图片URL（确保可访问）。</p>
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// <p>参考视频的获取链接。生成视频中的人物动作与参考视频一致。  视频内容需满足以下要求：  人物需要漏出清晰的上半身或全身的全部肢体及头部，避免遮挡  建议上传1人动作视频，2人及以上会取画面占比最大的人物动作进行生成  推荐使用真人动作，部分风格化的人物/类人肢体比例可以通过  动作视频一镜到底，角色始终出现在画面中，避免切镜、运镜等。否则会被截取  动作避免过快，相对平稳的动作生成效果更佳  视频文件支持.mp4/.mov，文件大小不超过100MB，仅支持长宽的边长均位于340px~3850px之间，上述校验不通过会返回错误码等信息  视频时长下限不短于3秒，时长上限与人物朝向参考（character_orientation）有关：  当人物朝向与视频中人物一致时，视频时长最长可达30秒；  当人物朝向与图片中人物一致时，视频时长最长可达10秒；  如果您的动作难度比较高、速度比较快，有一定概率生成不足上传视频时长的结果，因为模型只能提取有效动作时长进行生成，最短提取出3s可用连续动作即可生成。请注意，因此消耗的积分将无法退还，建议适当调整动作难度与速度  系统会校验视频内容，如有问题会返回错误码等信息。</p>
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`

	// <p>生成视频的模式</p><p>枚举值：std，pro<br>其中std：标准模式（标准），基础模式，性价比高<br>其中pro：专家模式（高品质），高表现模式，生成视频质量更佳</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>可选择是否保留视频原声  枚举值：yes，no  其中yes：保留视频原声  其中no：不保留视频原声。</p>
	KeepOriginalSound *string `json:"KeepOriginalSound,omitnil,omitempty" name:"KeepOriginalSound"`

	// <p>生成视频中人物的朝向，可选择与图片一致或与视频一致  枚举值：image，video，其中：  其中image：与图片中人物朝向一致；此时参考视频时长不得超过10秒；  其中video：与视频中人物朝向一致；此时参考视频时长不得超过30秒；  引用主体时，生成的视频暂时只能参考视频中的人物朝向。</p>
	CharacterOrientation *string `json:"CharacterOrientation,omitnil,omitempty" name:"CharacterOrientation"`

	// <p>参考主体列表  基于主体库中主体的ID配置，用key:value承载，如下：  &quot;element_list&quot;:[     {        &quot;element_id&quot;:long   },   {        &quot;element_id&quot;:long   } ] 参考主体数量与有无参考视频、参考图片数量有关，其中：  当使用首帧生成视频时，最多支持3个主体；  当使用首尾帧生成视频时，kling-v3-omni模型最多支持3个主体，kling-video-o1模不支持主体；  有参考视频时，参考图片数量和参考主体数量之和不得超过4，且不支持使用视频角色主体；  无参考视频时，参考图片数量和参考主体数量之和不得超过7；</p>
	ElementList []*Element `json:"ElementList,omitnil,omitempty" name:"ElementList"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。 1：添加标识； 0：不添加标识； 其他数值：默认按1处理。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitMotionControlKlingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitMotionControlKlingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Prompt")
	delete(f, "ExternalTaskId")
	delete(f, "Image")
	delete(f, "Video")
	delete(f, "Mode")
	delete(f, "KeepOriginalSound")
	delete(f, "CharacterOrientation")
	delete(f, "ElementList")
	delete(f, "CallbackUrl")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitMotionControlKlingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitMotionControlKlingJobResponseParams struct {
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitMotionControlKlingJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitMotionControlKlingJobResponseParams `json:"Response"`
}

func (r *SubmitMotionControlKlingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitMotionControlKlingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitPortraitSingJobRequestParams struct {
	// 传入音频URL地址，音频要求：
	// - 音频时长：2秒 - 60秒
	// - 音频格式：mp3、wav、m4a
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 传入图片URL地址，图片要求：
	// - 图片格式：jpg、jpeg、png、bmp、webp
	// - 图片分辨率：192～4096
	// - 图片大小：不超过10M
	// - 图片宽高比：图片【宽：高】在1:2到2:1范围内
	// - 图片内容：避免上传无人脸、无宠物脸或脸部过小、不完整、不清晰、偏转角度过大、嘴部被遮挡的图片。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 传入图片Base64编码，编码后请求体大小不超过10M。
	// 图片Base64编码与URL地址必传其一，如果都传以ImageBase64为准。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 唱演模式，默认使用人像模式。
	// Person：人像模式，仅支持上传人像图片，人像生成效果更好，如果图中未检测到有效人脸将被拦截，生成时会将视频短边分辨率放缩至512。
	// Pet：宠物模式，支持宠物等非人像图片，固定生成512:512分辨率视频。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 生成视频尺寸。可选取值："512:512"。
	// 
	// 人像模式下，如果不传该参数，默认生成视频的短边分辨率为512，长边分辨率不固定、由模型根据生成效果自动适配得到。如需固定生成分辨率可传入512:512。
	// 
	// 宠物模式下，如果不传该参数，默认将脸部唱演视频回贴原图，生成视频分辨率与原图一致。如不需要脸部回贴，仅保留脸部唱演视频，可传入512:512。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成视频添加标识的开关，默认为1。 
	// 1：添加标识；
	//  0：不添加标识；
	// 其他数值：默认按1处理。 
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitPortraitSingJobRequest struct {
	*tchttp.BaseRequest
	
	// 传入音频URL地址，音频要求：
	// - 音频时长：2秒 - 60秒
	// - 音频格式：mp3、wav、m4a
	AudioUrl *string `json:"AudioUrl,omitnil,omitempty" name:"AudioUrl"`

	// 传入图片URL地址，图片要求：
	// - 图片格式：jpg、jpeg、png、bmp、webp
	// - 图片分辨率：192～4096
	// - 图片大小：不超过10M
	// - 图片宽高比：图片【宽：高】在1:2到2:1范围内
	// - 图片内容：避免上传无人脸、无宠物脸或脸部过小、不完整、不清晰、偏转角度过大、嘴部被遮挡的图片。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 传入图片Base64编码，编码后请求体大小不超过10M。
	// 图片Base64编码与URL地址必传其一，如果都传以ImageBase64为准。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 唱演模式，默认使用人像模式。
	// Person：人像模式，仅支持上传人像图片，人像生成效果更好，如果图中未检测到有效人脸将被拦截，生成时会将视频短边分辨率放缩至512。
	// Pet：宠物模式，支持宠物等非人像图片，固定生成512:512分辨率视频。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 生成视频尺寸。可选取值："512:512"。
	// 
	// 人像模式下，如果不传该参数，默认生成视频的短边分辨率为512，长边分辨率不固定、由模型根据生成效果自动适配得到。如需固定生成分辨率可传入512:512。
	// 
	// 宠物模式下，如果不传该参数，默认将脸部唱演视频回贴原图，生成视频分辨率与原图一致。如不需要脸部回贴，仅保留脸部唱演视频，可传入512:512。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成视频添加标识的开关，默认为1。 
	// 1：添加标识；
	//  0：不添加标识；
	// 其他数值：默认按1处理。 
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitPortraitSingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitPortraitSingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AudioUrl")
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "Mode")
	delete(f, "Resolution")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitPortraitSingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitPortraitSingJobResponseParams struct {
	// 任务ID。任务有效期为48小时。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitPortraitSingJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitPortraitSingJobResponseParams `json:"Response"`
}

func (r *SubmitPortraitSingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitPortraitSingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitReferenceToVideoViduJobRequestParams struct {
	// <p>文本提示词<br>生成视频的文本描述。<br>注1：字符长度不能超过 2000 个字符<br>注2：使用Subjects主体参数时，可以通过@主体id 来表示主体内容，例如：&quot;@1 和 @2 在一起吃火锅，并且旁白音说火锅大家都爱吃。&quot;</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>图像参考<br>【非主体调用时传入】<br>支持上传1～7张图片，模型将以此参数中传入的图片中的主题为参考生成具备主体一致的视频。<br>支持传入图片 Base64 编码或图片URL（确保可访问）<br>图片支持 png、jpeg、jpg、webp格式<br>图片像素不能小于 128*128，且比例需要小于1:4或者4:1，且大小不超过50M。</p>
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>参考生图主体信息，支持1-7个主体，主体图片共1 ～ 7张<br>【主体调用时传入】</p>
	Subjects []*ReferenceSubject `json:"Subjects,omitnil,omitempty" name:"Subjects"`

	// <p>视频参考支持上传1～2个视频，模型将以此参数中传入的视频作为参考，生成具备主体一致的视频。<br>注1： 仅viduq2-pro模型支持该参数<br>注2：使用视频参考功能时，最多支持上传 1个8秒 视频 或 2个5秒 视频<br>注3：视频支持 mp4、avi、mov格式<br>注4：视频像素不能小于 128*128，且比例需要小于1:4或者4:1，且大小不超过100M。</p>
	Videos []*string `json:"Videos,omitnil,omitempty" name:"Videos"`

	// <p>模型名称，可选值，当前仅可并且默认viduq2</p><ul><li>viduq2：最新模型</li></ul>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>是否使用音视频直出能力，默认false，可选值 true、false</p><ul><li>true：使用音视频直出能力。</li><li>false：不使用音视频直出能力。<br>仅上传主体时支持此功能</li></ul>
	Audio *bool `json:"Audio,omitnil,omitempty" name:"Audio"`

	// <p>音频类型，audio为true时必填，默认为all</p><p>枚举值：</p><ul><li>all： 音效+人声</li><li>speech_only： 仅人声</li><li>sound_effect_only： 仅音效</li></ul>
	AudioType *string `json:"AudioType,omitnil,omitempty" name:"AudioType"`

	// <p>是否为生成的视频添加背景音乐。<br>默认：false，可选值 true 、false<br>【仅支持非主体调用时生效】</p><ul><li>传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。</li><li>BGM不限制时长，系统根据视频时长自动适配</li><li>BGM参数在q2系列模型的duration为 9秒 或 10秒 时不生效</li></ul>
	Bgm *bool `json:"Bgm,omitnil,omitempty" name:"Bgm"`

	// <p>视频时长参数：<br>默认5秒，可选：1-10（整数）</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>比例<br>默认值：16:9<br>【非主体调用时】<br>可选值如下：16:9、9:16、4:3、3:4、1:1</p><p>【主体调用时】<br>可选值如下：16:9、9:16、1:1<br>注：q2模型 支持任意宽高比</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>分辨率参数<br>默认 720p, 可选：540p、720p、1080p</p>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>运动幅度默认 auto，可选值：auto、small、medium、large<br>注：使用q2系列模型时该参数不生效</p>
	MovementAmplitude *string `json:"MovementAmplitude,omitnil,omitempty" name:"MovementAmplitude"`

	// <p>元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：<br>{<br>&quot;Label&quot;: &quot;your_label&quot;,&quot;ContentProducer&quot;: &quot;yourcontentproducer&quot;,&quot;ContentPropagator&quot;: &quot;your_content_propagator&quot;,&quot;ProduceID&quot;: &quot;yourproductid&quot;, &quot;PropagateID&quot;: &quot;your_propagate_id&quot;,&quot;ReservedCode1&quot;: &quot;yourreservedcode1&quot;, &quot;ReservedCode2&quot;: &quot;your_reserved_code2&quot;<br>}<br>该参数为空时，默认使用vidu生成的元数据标识</p>
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>Callback 协议<br>需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成完成时，将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致<br>回调返回的&quot;status&quot;包括以下状态：</p><ul><li>success 任务完成（如发送失败，回调三次）</li><li>failed 任务失败（如发送失败，回调三次）</li></ul>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>透传参数不做任何处理，仅数据传输注：最多 1048576个字符</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// <p>错峰模式，默认为：false，可选值：</p><ul><li>true：错峰生成视频；</li><li>false：即时生成视频；<br>注1：错峰模式消耗的积分更低<br>注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分</li></ul>
	OffPeak *bool `json:"OffPeak,omitnil,omitempty" name:"OffPeak"`

	// <p>为生成结果图添加显式水印标识的开关，默认为1。<br>1：添加。<br>0：不添加。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。<br>示例值：1</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitReferenceToVideoViduJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>文本提示词<br>生成视频的文本描述。<br>注1：字符长度不能超过 2000 个字符<br>注2：使用Subjects主体参数时，可以通过@主体id 来表示主体内容，例如：&quot;@1 和 @2 在一起吃火锅，并且旁白音说火锅大家都爱吃。&quot;</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>图像参考<br>【非主体调用时传入】<br>支持上传1～7张图片，模型将以此参数中传入的图片中的主题为参考生成具备主体一致的视频。<br>支持传入图片 Base64 编码或图片URL（确保可访问）<br>图片支持 png、jpeg、jpg、webp格式<br>图片像素不能小于 128*128，且比例需要小于1:4或者4:1，且大小不超过50M。</p>
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>参考生图主体信息，支持1-7个主体，主体图片共1 ～ 7张<br>【主体调用时传入】</p>
	Subjects []*ReferenceSubject `json:"Subjects,omitnil,omitempty" name:"Subjects"`

	// <p>视频参考支持上传1～2个视频，模型将以此参数中传入的视频作为参考，生成具备主体一致的视频。<br>注1： 仅viduq2-pro模型支持该参数<br>注2：使用视频参考功能时，最多支持上传 1个8秒 视频 或 2个5秒 视频<br>注3：视频支持 mp4、avi、mov格式<br>注4：视频像素不能小于 128*128，且比例需要小于1:4或者4:1，且大小不超过100M。</p>
	Videos []*string `json:"Videos,omitnil,omitempty" name:"Videos"`

	// <p>模型名称，可选值，当前仅可并且默认viduq2</p><ul><li>viduq2：最新模型</li></ul>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>是否使用音视频直出能力，默认false，可选值 true、false</p><ul><li>true：使用音视频直出能力。</li><li>false：不使用音视频直出能力。<br>仅上传主体时支持此功能</li></ul>
	Audio *bool `json:"Audio,omitnil,omitempty" name:"Audio"`

	// <p>音频类型，audio为true时必填，默认为all</p><p>枚举值：</p><ul><li>all： 音效+人声</li><li>speech_only： 仅人声</li><li>sound_effect_only： 仅音效</li></ul>
	AudioType *string `json:"AudioType,omitnil,omitempty" name:"AudioType"`

	// <p>是否为生成的视频添加背景音乐。<br>默认：false，可选值 true 、false<br>【仅支持非主体调用时生效】</p><ul><li>传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。</li><li>BGM不限制时长，系统根据视频时长自动适配</li><li>BGM参数在q2系列模型的duration为 9秒 或 10秒 时不生效</li></ul>
	Bgm *bool `json:"Bgm,omitnil,omitempty" name:"Bgm"`

	// <p>视频时长参数：<br>默认5秒，可选：1-10（整数）</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>比例<br>默认值：16:9<br>【非主体调用时】<br>可选值如下：16:9、9:16、4:3、3:4、1:1</p><p>【主体调用时】<br>可选值如下：16:9、9:16、1:1<br>注：q2模型 支持任意宽高比</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>分辨率参数<br>默认 720p, 可选：540p、720p、1080p</p>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>运动幅度默认 auto，可选值：auto、small、medium、large<br>注：使用q2系列模型时该参数不生效</p>
	MovementAmplitude *string `json:"MovementAmplitude,omitnil,omitempty" name:"MovementAmplitude"`

	// <p>元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：<br>{<br>&quot;Label&quot;: &quot;your_label&quot;,&quot;ContentProducer&quot;: &quot;yourcontentproducer&quot;,&quot;ContentPropagator&quot;: &quot;your_content_propagator&quot;,&quot;ProduceID&quot;: &quot;yourproductid&quot;, &quot;PropagateID&quot;: &quot;your_propagate_id&quot;,&quot;ReservedCode1&quot;: &quot;yourreservedcode1&quot;, &quot;ReservedCode2&quot;: &quot;your_reserved_code2&quot;<br>}<br>该参数为空时，默认使用vidu生成的元数据标识</p>
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>Callback 协议<br>需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成完成时，将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致<br>回调返回的&quot;status&quot;包括以下状态：</p><ul><li>success 任务完成（如发送失败，回调三次）</li><li>failed 任务失败（如发送失败，回调三次）</li></ul>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>透传参数不做任何处理，仅数据传输注：最多 1048576个字符</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// <p>错峰模式，默认为：false，可选值：</p><ul><li>true：错峰生成视频；</li><li>false：即时生成视频；<br>注1：错峰模式消耗的积分更低<br>注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分</li></ul>
	OffPeak *bool `json:"OffPeak,omitnil,omitempty" name:"OffPeak"`

	// <p>为生成结果图添加显式水印标识的开关，默认为1。<br>1：添加。<br>0：不添加。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。<br>示例值：1</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitReferenceToVideoViduJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitReferenceToVideoViduJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "Images")
	delete(f, "Subjects")
	delete(f, "Videos")
	delete(f, "Model")
	delete(f, "Audio")
	delete(f, "AudioType")
	delete(f, "Bgm")
	delete(f, "Duration")
	delete(f, "AspectRatio")
	delete(f, "Resolution")
	delete(f, "MovementAmplitude")
	delete(f, "MetaData")
	delete(f, "CallbackUrl")
	delete(f, "Payload")
	delete(f, "OffPeak")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitReferenceToVideoViduJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitReferenceToVideoViduJobResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitReferenceToVideoViduJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitReferenceToVideoViduJobResponseParams `json:"Response"`
}

func (r *SubmitReferenceToVideoViduJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitReferenceToVideoViduJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTemplateToVideoJobRequestParams struct {
	// <p>特效模板名称。请在 <a href="https://cloud.tencent.com/document/product/1616/119194">视频特效模板列表</a>  中选择想要生成的特效对应的 template 名称。</p>
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// <p>参考图像，不同特效输入图片的数量详见： <a href="https://cloud.tencent.com/document/product/1616/119194">视频特效模板-图片要求说明</a></p><ul><li>支持传入图片Base64编码或图片URL（确保可访问）</li><li>图片格式：支持png、jpg、jpeg、webp、bmp、tiff</li><li>图片文件：大小不能超过10MB（base64后），图片分辨率不小于300×300px，不大于4096×4096，图片宽高比应在1:4 ~ 4:1之间</li></ul>
	Images []*Image `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>为生成视频添加标识的开关，默认为1。传0 需前往  <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a> 申请开启显式标识自主完成后方可生效。<br>1：添加标识；<br>0：不添加标识；<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“ AI 生成”或“视频由 AI 生成”字样，如需替换为其他的标识图片，需前往  <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a> 申请开启显式标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>视频输出分辨率，默认值：360p 。不同特效支持的清晰度及消耗积分数详见：<a href="https://cloud.tencent.com/document/product/1616/119194">视频特效模板-单次调用消耗积分数列</a></p>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>是否为生成的视频添加背景音乐。默认：false，  传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。</p>
	BGM *bool `json:"BGM,omitnil,omitempty" name:"BGM"`

	// <p>扩展字段。</p>
	ExtraParam *ExtraParam `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`
}

type SubmitTemplateToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>特效模板名称。请在 <a href="https://cloud.tencent.com/document/product/1616/119194">视频特效模板列表</a>  中选择想要生成的特效对应的 template 名称。</p>
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// <p>参考图像，不同特效输入图片的数量详见： <a href="https://cloud.tencent.com/document/product/1616/119194">视频特效模板-图片要求说明</a></p><ul><li>支持传入图片Base64编码或图片URL（确保可访问）</li><li>图片格式：支持png、jpg、jpeg、webp、bmp、tiff</li><li>图片文件：大小不能超过10MB（base64后），图片分辨率不小于300×300px，不大于4096×4096，图片宽高比应在1:4 ~ 4:1之间</li></ul>
	Images []*Image `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>为生成视频添加标识的开关，默认为1。传0 需前往  <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a> 申请开启显式标识自主完成后方可生效。<br>1：添加标识；<br>0：不添加标识；<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“ AI 生成”或“视频由 AI 生成”字样，如需替换为其他的标识图片，需前往  <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a> 申请开启显式标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>视频输出分辨率，默认值：360p 。不同特效支持的清晰度及消耗积分数详见：<a href="https://cloud.tencent.com/document/product/1616/119194">视频特效模板-单次调用消耗积分数列</a></p>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>是否为生成的视频添加背景音乐。默认：false，  传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。</p>
	BGM *bool `json:"BGM,omitnil,omitempty" name:"BGM"`

	// <p>扩展字段。</p>
	ExtraParam *ExtraParam `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`
}

func (r *SubmitTemplateToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTemplateToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Template")
	delete(f, "Images")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "Resolution")
	delete(f, "BGM")
	delete(f, "ExtraParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTemplateToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTemplateToVideoJobResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTemplateToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTemplateToVideoJobResponseParams `json:"Response"`
}

func (r *SubmitTemplateToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTemplateToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextToVideoJobRequestParams struct {
	// <p>正向文本提示词。不能超过2500个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>模型名称。<br>v1.0：Kling-V1<br>v1.5：Kling-V1-5<br>v1.6：Kling-V1-6<br>v2.0：Kling-V2-Master<br>v2.1m：Kling-V2-1-master<br>v2.5：Kling-V2-5-Turbo<br>v2.6：Kling-V2-6<br>v3.0：kling-v3</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>负向文本提示词。不能超过2500个字符</p>
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// <p>生成视频时长，单位s。默认值为5。<br>枚举值：3，4，5，6，7，8，9，10，11，12，13，14，15不同模型支持时长不同<br>●模型v1.0、v1.6、v2.0、v2.1m、v2.5、v2.6：支持5、10<br>●模型v3.0：支持3～15s</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>生成视频的模式；<br>枚举值：std，pro<br>●其中std：标准模式（标准），基础模式，性价比高<br>●其中pro：专家模式（高品质），高表现模式，生成视频质量更佳<br>不同模型版本、视频模式支持范围不同</p><p>●v1.6：std、pro。<br>●v1.0、v1.5：pro<br>●v2.0、v2.1m、v3.0：模型无需配置。<br>●v2.5：首尾帧情况下支持pro。<br>●v2.6：仅支持pro，选择v2.6模型时，默认自动生成高品质pro视频。</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>生成视频的自由度；值越大，模型自由度越小，与用户输入的提示词相关性越强。<br>取值范围：[0, 1]<br>v2.0、v2.5、v2.6 模型不支持当前参数<br>默认值：0.5。</p>
	CfgScale *float64 `json:"CfgScale,omitnil,omitempty" name:"CfgScale"`

	// <p>生成视频的画面纵横比（宽:高）<br>枚举值：16:9, 9:16, 1:1<br>默认值：16:9</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>生成视频时是否同时生成声音</p><ul><li>枚举值：on，off</li></ul><p>仅V2.6及后续版本模型支持当前参数，v2.6模型的std模式只能生成无声的视频</p>
	Sound *string `json:"Sound,omitnil,omitempty" name:"Sound"`

	// <p>为生成视频添加标识的开关，默认为1。<br>1：添加标识。<br>0：不添加标识。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>是否生成多镜头视频</p><ul><li>当前参数为true时，prompt参数无效</li><li>当前参数为false时，shot_type参数及multi_prompt参数无效</li></ul>
	MultiShot *bool `json:"MultiShot,omitnil,omitempty" name:"MultiShot"`

	// <p>分镜方式</p><p>枚举值：customize，intelligence<br>当MultiShot参数为true时，当前参数必填</p>
	ShotType *string `json:"ShotType,omitnil,omitempty" name:"ShotType"`

	// <p>各分镜提示词，可包含正向描述和负向描述</p><p>通过index、prompt、duration参数定义分镜序号及相应提示词和时长，其中：</p><ul><li>最多支持6个分镜，最小支持1个分镜</li><li>每个分镜相关内容的最大长度不超过512</li><li>每个分镜的时长不大于当前任务的总时长，不小于1</li><li>所有分镜的时长之和等于当前任务的总时长</li><li>当MultiShot参数为true且ShotType参数为customize时，当前参数不得为空<br>用key:value承载，如下：<pre><code> &quot;MultiPrompt&quot;:[              {                &quot;Index&quot;:int,              &quot;Prompt&quot;: &quot;string&quot;,              &quot;Duration&quot;: &quot;5&quot;            }  ]</code></pre></li></ul>
	MultiPrompt []*MultiPrompt `json:"MultiPrompt,omitnil,omitempty" name:"MultiPrompt"`

	// <p>控制摄像机运动的协议（如未指定，模型将根据输入的文本/图片进行智能匹配）</p>
	CameraControl *CameraControl `json:"CameraControl,omitnil,omitempty" name:"CameraControl"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

type SubmitTextToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>正向文本提示词。不能超过2500个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>模型名称。<br>v1.0：Kling-V1<br>v1.5：Kling-V1-5<br>v1.6：Kling-V1-6<br>v2.0：Kling-V2-Master<br>v2.1m：Kling-V2-1-master<br>v2.5：Kling-V2-5-Turbo<br>v2.6：Kling-V2-6<br>v3.0：kling-v3</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>负向文本提示词。不能超过2500个字符</p>
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// <p>生成视频时长，单位s。默认值为5。<br>枚举值：3，4，5，6，7，8，9，10，11，12，13，14，15不同模型支持时长不同<br>●模型v1.0、v1.6、v2.0、v2.1m、v2.5、v2.6：支持5、10<br>●模型v3.0：支持3～15s</p>
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>生成视频的模式；<br>枚举值：std，pro<br>●其中std：标准模式（标准），基础模式，性价比高<br>●其中pro：专家模式（高品质），高表现模式，生成视频质量更佳<br>不同模型版本、视频模式支持范围不同</p><p>●v1.6：std、pro。<br>●v1.0、v1.5：pro<br>●v2.0、v2.1m、v3.0：模型无需配置。<br>●v2.5：首尾帧情况下支持pro。<br>●v2.6：仅支持pro，选择v2.6模型时，默认自动生成高品质pro视频。</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>生成视频的自由度；值越大，模型自由度越小，与用户输入的提示词相关性越强。<br>取值范围：[0, 1]<br>v2.0、v2.5、v2.6 模型不支持当前参数<br>默认值：0.5。</p>
	CfgScale *float64 `json:"CfgScale,omitnil,omitempty" name:"CfgScale"`

	// <p>生成视频的画面纵横比（宽:高）<br>枚举值：16:9, 9:16, 1:1<br>默认值：16:9</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>生成视频时是否同时生成声音</p><ul><li>枚举值：on，off</li></ul><p>仅V2.6及后续版本模型支持当前参数，v2.6模型的std模式只能生成无声的视频</p>
	Sound *string `json:"Sound,omitnil,omitempty" name:"Sound"`

	// <p>为生成视频添加标识的开关，默认为1。<br>1：添加标识。<br>0：不添加标识。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>是否生成多镜头视频</p><ul><li>当前参数为true时，prompt参数无效</li><li>当前参数为false时，shot_type参数及multi_prompt参数无效</li></ul>
	MultiShot *bool `json:"MultiShot,omitnil,omitempty" name:"MultiShot"`

	// <p>分镜方式</p><p>枚举值：customize，intelligence<br>当MultiShot参数为true时，当前参数必填</p>
	ShotType *string `json:"ShotType,omitnil,omitempty" name:"ShotType"`

	// <p>各分镜提示词，可包含正向描述和负向描述</p><p>通过index、prompt、duration参数定义分镜序号及相应提示词和时长，其中：</p><ul><li>最多支持6个分镜，最小支持1个分镜</li><li>每个分镜相关内容的最大长度不超过512</li><li>每个分镜的时长不大于当前任务的总时长，不小于1</li><li>所有分镜的时长之和等于当前任务的总时长</li><li>当MultiShot参数为true且ShotType参数为customize时，当前参数不得为空<br>用key:value承载，如下：<pre><code> &quot;MultiPrompt&quot;:[              {                &quot;Index&quot;:int,              &quot;Prompt&quot;: &quot;string&quot;,              &quot;Duration&quot;: &quot;5&quot;            }  ]</code></pre></li></ul>
	MultiPrompt []*MultiPrompt `json:"MultiPrompt,omitnil,omitempty" name:"MultiPrompt"`

	// <p>控制摄像机运动的协议（如未指定，模型将根据输入的文本/图片进行智能匹配）</p>
	CameraControl *CameraControl `json:"CameraControl,omitnil,omitempty" name:"CameraControl"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`
}

func (r *SubmitTextToVideoJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextToVideoJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "Model")
	delete(f, "NegativePrompt")
	delete(f, "Duration")
	delete(f, "Mode")
	delete(f, "CfgScale")
	delete(f, "AspectRatio")
	delete(f, "Sound")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "MultiShot")
	delete(f, "ShotType")
	delete(f, "MultiPrompt")
	delete(f, "CameraControl")
	delete(f, "CallbackUrl")
	delete(f, "ExternalTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTextToVideoJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextToVideoJobResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTextToVideoJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTextToVideoJobResponseParams `json:"Response"`
}

func (r *SubmitTextToVideoJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextToVideoJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextToVideoViduJobRequestParams struct {
	// <p>文本提示词<br>生成视频的文本描述。<br>注：字符长度不能超过 2000 个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>模型名称，可选值，默认viduq2</p><ul><li>viduq2：最新模型</li><li>viduq3-turbo：对比viduq3-pro，生成速度更快</li><li>viduq3-pro：高效生成优质音视频内容，让视频内容更生动、更形象、更立体，效果更好</li></ul>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>视频时长参数，默认值依据模型而定：</p><ul><li>viduq3-pro、viduq3-turbo: 默认5秒，可选：1-16</li><li>viduq2 : 默认5秒，可选：1-10</li></ul>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>是否为生成的视频添加背景音乐。<br>默认：false，可选值 true 、false<br>传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。</p><ul><li>BGM不限制时长，系统根据视频时长自动适配</li><li>BGM参数在q2模型的duration为 9秒 或 10秒 时不生效；该参数在q3系列模型中不生效</li></ul>
	Bgm *bool `json:"Bgm,omitnil,omitempty" name:"Bgm"`

	// <p>比例<br>默认 16:9，可选值如下：16:9、9:16、4:3、3:4、1:1</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>分辨率参数，默认值依据模型和视频时长而定：</p><ul><li>viduq3-pro、viduq3-turbo(1-16秒)：默认 720p，可选：540p、720p、1080p</li><li>viduq2(1-10秒)：默认 720p，可选：540p、720p、1080p</li></ul>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>风格<br>默认 general，可选值：general、anime<br>general：通用风格，可以通过提示词来控制风格<br>anime：动漫风格，仅在动漫风格表现突出，可以通过不同的动漫风格提示词来控制<br>注：使用q2、q3系列模型时该参数不生效</p>
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// <p>运动幅度<br>默认 auto，可选值：auto、small、medium、large<br>注：使用q2、q3系列模型时该参数不生效</p>
	MovementAmplitude *string `json:"MovementAmplitude,omitnil,omitempty" name:"MovementAmplitude"`

	// <p>是否使用音视频直出能力，默认为true，枚举值为：</p><ul><li>false：不需要音视频直出，输出静音视频</li><li>true：需要音画同步，输出声音的视频（包括台词和音效）<br>注1：仅q3系列模型支持该参数</li></ul>
	Audio *bool `json:"Audio,omitnil,omitempty" name:"Audio"`

	// <p>元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：<br>{<br>&quot;Label&quot;: &quot;your_label&quot;,<br>&quot;ContentProducer&quot;: &quot;your_content_producer&quot;,<br>&quot;ContentPropagator&quot;: &quot;your_content_propagator&quot;,<br>&quot;ProduceID&quot;: &quot;your_product_id&quot;,<br>&quot;PropagateID&quot;: &quot;your_propagate_id&quot;,<br>&quot;ReservedCode1&quot;: &quot;your_reserved_code1&quot;,<br>&quot;ReservedCode2&quot;: &quot;your_reserved_code2&quot;<br>}<br>该参数为空时，默认使用vidu生成的元数据标识</p>
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>Callback 协议<br>需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务完成时，将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致<br>回调返回的&quot;status&quot;包括以下状态：</p><ul><li>success 任务完成（如发送失败，回调三次）</li><li>failed 任务失败（如发送失败，回调三次）</li></ul>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>透传参数 不做任何处理，仅数据传输 注：最多 1048576个字符</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// <p>错峰模式，默认为：false，可选值：</p><ul><li>true：错峰生成视频；</li><li>false：即时生成视频；<br>注1：错峰模式消耗的积分更低<br>注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分</li></ul>
	OffPeak *bool `json:"OffPeak,omitnil,omitempty" name:"OffPeak"`

	// <p>为生成结果图添加显式水印标识的开关，默认为1。<br>1：添加。<br>0：不添加。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。<br>示例值：1</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitTextToVideoViduJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>文本提示词<br>生成视频的文本描述。<br>注：字符长度不能超过 2000 个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>模型名称，可选值，默认viduq2</p><ul><li>viduq2：最新模型</li><li>viduq3-turbo：对比viduq3-pro，生成速度更快</li><li>viduq3-pro：高效生成优质音视频内容，让视频内容更生动、更形象、更立体，效果更好</li></ul>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>视频时长参数，默认值依据模型而定：</p><ul><li>viduq3-pro、viduq3-turbo: 默认5秒，可选：1-16</li><li>viduq2 : 默认5秒，可选：1-10</li></ul>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>是否为生成的视频添加背景音乐。<br>默认：false，可选值 true 、false<br>传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。</p><ul><li>BGM不限制时长，系统根据视频时长自动适配</li><li>BGM参数在q2模型的duration为 9秒 或 10秒 时不生效；该参数在q3系列模型中不生效</li></ul>
	Bgm *bool `json:"Bgm,omitnil,omitempty" name:"Bgm"`

	// <p>比例<br>默认 16:9，可选值如下：16:9、9:16、4:3、3:4、1:1</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>分辨率参数，默认值依据模型和视频时长而定：</p><ul><li>viduq3-pro、viduq3-turbo(1-16秒)：默认 720p，可选：540p、720p、1080p</li><li>viduq2(1-10秒)：默认 720p，可选：540p、720p、1080p</li></ul>
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// <p>风格<br>默认 general，可选值：general、anime<br>general：通用风格，可以通过提示词来控制风格<br>anime：动漫风格，仅在动漫风格表现突出，可以通过不同的动漫风格提示词来控制<br>注：使用q2、q3系列模型时该参数不生效</p>
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// <p>运动幅度<br>默认 auto，可选值：auto、small、medium、large<br>注：使用q2、q3系列模型时该参数不生效</p>
	MovementAmplitude *string `json:"MovementAmplitude,omitnil,omitempty" name:"MovementAmplitude"`

	// <p>是否使用音视频直出能力，默认为true，枚举值为：</p><ul><li>false：不需要音视频直出，输出静音视频</li><li>true：需要音画同步，输出声音的视频（包括台词和音效）<br>注1：仅q3系列模型支持该参数</li></ul>
	Audio *bool `json:"Audio,omitnil,omitempty" name:"Audio"`

	// <p>元数据标识，json格式字符串，透传字段，您可以 自定义格式 或使用 示例格式 ，示例如下：<br>{<br>&quot;Label&quot;: &quot;your_label&quot;,<br>&quot;ContentProducer&quot;: &quot;your_content_producer&quot;,<br>&quot;ContentPropagator&quot;: &quot;your_content_propagator&quot;,<br>&quot;ProduceID&quot;: &quot;your_product_id&quot;,<br>&quot;PropagateID&quot;: &quot;your_propagate_id&quot;,<br>&quot;ReservedCode1&quot;: &quot;your_reserved_code1&quot;,<br>&quot;ReservedCode2&quot;: &quot;your_reserved_code2&quot;<br>}<br>该参数为空时，默认使用vidu生成的元数据标识</p>
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>Callback 协议<br>需要您在创建任务时主动设置 callback_url，请求方法为 POST，当视频生成任务完成时，将向此地址发送包含任务最新状态的回调请求。回调请求内容结构与查询任务API的返回体一致<br>回调返回的&quot;status&quot;包括以下状态：</p><ul><li>success 任务完成（如发送失败，回调三次）</li><li>failed 任务失败（如发送失败，回调三次）</li></ul>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>透传参数 不做任何处理，仅数据传输 注：最多 1048576个字符</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// <p>错峰模式，默认为：false，可选值：</p><ul><li>true：错峰生成视频；</li><li>false：即时生成视频；<br>注1：错峰模式消耗的积分更低<br>注2：错峰模式下提交的任务，会在48小时内生成，未能完成的任务会被自动取消，并返还该任务的积分</li></ul>
	OffPeak *bool `json:"OffPeak,omitnil,omitempty" name:"OffPeak"`

	// <p>为生成结果图添加显式水印标识的开关，默认为1。<br>1：添加。<br>0：不添加。<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。<br>示例值：1</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitTextToVideoViduJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextToVideoViduJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "Model")
	delete(f, "Duration")
	delete(f, "Bgm")
	delete(f, "AspectRatio")
	delete(f, "Resolution")
	delete(f, "Style")
	delete(f, "MovementAmplitude")
	delete(f, "Audio")
	delete(f, "MetaData")
	delete(f, "CallbackUrl")
	delete(f, "Payload")
	delete(f, "OffPeak")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTextToVideoViduJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextToVideoViduJobResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTextToVideoViduJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTextToVideoViduJobResponseParams `json:"Response"`
}

func (r *SubmitTextToVideoViduJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextToVideoViduJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoEditKlingJobRequestParams struct {
	// <p>文本提示词，可包含正向描述和负向描述</p><p>可将提示词模板化来满足不同的视频生成需求</p><p>不能超过2500个字</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>模型名称，支持kling-video-o1，kling-v3-omni。默认kling-video-o1。</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// <p>参考图列表</p><p>包括主体、场景、风格等参考图片，也可作为首帧或尾帧生成视频；当作为首帧或尾帧生成视频时：</p><p>通过type参数来定义图片是否为首尾帧：first_frame为首帧，end_frame为尾帧</p><p>暂时不支持仅尾帧，即有尾帧图时必须有首帧图</p><p>首帧或首尾帧生视频时，不能使用视频编辑功能</p><p>用key:value承载，如下：</p><p>&quot;ImageInfo&quot;:[<br>    {<br>      &quot;ImageUrl&quot;:&quot;https://cos.ap-guangzhou.myqcloud.com/test.png&quot;,<br>    &quot;Type&quot;:&quot;first_frame&quot;<br>  },<br>  {<br>      &quot;ImageUrl&quot;:&quot;https://cos.ap-guangzhou.myqcloud.com/test.png&quot;,<br>    &quot;Type&quot;:&quot;end_frame&quot;<br>  }<br>]<br>支持传入图片URL（确保可访问）</p><p>图片格式支持.jpg / .jpeg / .png</p><p>图片文件大小不能超过10MB，图片宽高尺寸不小于300px，不大于8000px，图片宽高比要在1:2.5 ~ 2.5:1之间</p><p>有参考视频时，参考图片数量不得超过4；无参考视频时，参考图片数量不得超过7</p><p>数组中超过2张图片时，不支持设置尾帧</p>
	ImageList []*ImageInfo `json:"ImageList,omitnil,omitempty" name:"ImageList"`

	// <p>生成视频的画面纵横比（宽:高）</p><p>枚举值：16:9, 9:16, 1:1</p><p>未使用首帧参考或视频编辑功能时，当前参数必填</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>生成视频时长，单位s</p><p>枚举值：3，4，5，6，7，8，9，10，其中：使用文生视频、首帧图生视频时，仅支持5和10s<br>使用视频编辑功能（“refer_type”:“base”）时，输出结果与传入视频时长相同，此时当前参数无效；此时，按输入视频时长四舍五入取整计量计费</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。 1：添加标识； 0：不添加标识； 其他数值：默认按1处理。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>生成视频的模式</p><p>枚举值：std，pro<br>其中std：标准模式（标准），基础模式，性价比高<br>其中pro：专家模式（高品质），高表现模式，生成视频质量更佳</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>参考视频，通过URL方式获取</p><p>可作为特征参考视频，也可作为待编辑视频，默认为待编辑视频；可选择性保留视频原声<br>通过ReferType参数区分参考视频类型：feature为特征参考视频，base为待编辑视频<br>参考视频为待编辑视频时，不能定义视频首尾帧<br>通过KeepOriginalSound参数选择是否保留视频原声，yes为保留，no为不保留；当前参数对特征参考视频（feature）也生效</p><ul><li>视频格式仅支持MP4/MOV</li><li>仅支持时长≥3秒且≤10秒的视频</li><li>视频宽高尺寸需介于720px（含）和2160px（含）之间</li><li>视频帧率基于24fps～60fps，生成视频时会输出为24fps</li><li>至多仅支持上传1段视频，视频大小不超过200MB</li></ul>
	VideoList []*ReferVideoInfo `json:"VideoList,omitnil,omitempty" name:"VideoList"`

	// <p>是否生成多镜头视频  当前参数为true时，prompt参数无效，且不支持设定首尾帧生视频  当前参数为false时，shot_type参数及multi_prompt参数无效</p>
	MultiShot *bool `json:"MultiShot,omitnil,omitempty" name:"MultiShot"`

	// <p>分镜方式  枚举值：customize  当multi_shot参数为true时，当前参数必填</p>
	ShotType *string `json:"ShotType,omitnil,omitempty" name:"ShotType"`

	// <p>各分镜信息，如提示词、时长等  通过index、prompt、duration参数定义分镜序号及相应提示词和时长，其中：  最多支持6个分镜，最小支持1个分镜  每个分镜相关内容的最大长度不超过512  每个分镜的时长不大于当前任务的总时长，不小于1  所有分镜的时长之和等于当前任务的总时长</p>
	MultiPrompt []*MultiPrompt `json:"MultiPrompt,omitnil,omitempty" name:"MultiPrompt"`

	// <p>参考主体列表  基于主体库中主体的ID配置，用key:value承载，如下：  &quot;element_list&quot;:[     {        &quot;element_id&quot;:long   },   {        &quot;element_id&quot;:long   } ] 参考主体数量与有无参考视频、参考图片数量有关，其中：  当使用首帧生成视频时，最多支持3个主体；  当使用首尾帧生成视频时，kling-v3-omni模型最多支持3个主体，kling-video-o1模不支持主体；  有参考视频时，参考图片数量和参考主体数量之和不得超过4，且不支持使用视频角色主体；  无参考视频时，参考图片数量和参考主体数量之和不得超过7；</p>
	ElementList []*Element `json:"ElementList,omitnil,omitempty" name:"ElementList"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>是否开启声音</p>
	Sound *string `json:"Sound,omitnil,omitempty" name:"Sound"`
}

type SubmitVideoEditKlingJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>文本提示词，可包含正向描述和负向描述</p><p>可将提示词模板化来满足不同的视频生成需求</p><p>不能超过2500个字</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>模型名称，支持kling-video-o1，kling-v3-omni。默认kling-video-o1。</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// <p>参考图列表</p><p>包括主体、场景、风格等参考图片，也可作为首帧或尾帧生成视频；当作为首帧或尾帧生成视频时：</p><p>通过type参数来定义图片是否为首尾帧：first_frame为首帧，end_frame为尾帧</p><p>暂时不支持仅尾帧，即有尾帧图时必须有首帧图</p><p>首帧或首尾帧生视频时，不能使用视频编辑功能</p><p>用key:value承载，如下：</p><p>&quot;ImageInfo&quot;:[<br>    {<br>      &quot;ImageUrl&quot;:&quot;https://cos.ap-guangzhou.myqcloud.com/test.png&quot;,<br>    &quot;Type&quot;:&quot;first_frame&quot;<br>  },<br>  {<br>      &quot;ImageUrl&quot;:&quot;https://cos.ap-guangzhou.myqcloud.com/test.png&quot;,<br>    &quot;Type&quot;:&quot;end_frame&quot;<br>  }<br>]<br>支持传入图片URL（确保可访问）</p><p>图片格式支持.jpg / .jpeg / .png</p><p>图片文件大小不能超过10MB，图片宽高尺寸不小于300px，不大于8000px，图片宽高比要在1:2.5 ~ 2.5:1之间</p><p>有参考视频时，参考图片数量不得超过4；无参考视频时，参考图片数量不得超过7</p><p>数组中超过2张图片时，不支持设置尾帧</p>
	ImageList []*ImageInfo `json:"ImageList,omitnil,omitempty" name:"ImageList"`

	// <p>生成视频的画面纵横比（宽:高）</p><p>枚举值：16:9, 9:16, 1:1</p><p>未使用首帧参考或视频编辑功能时，当前参数必填</p>
	AspectRatio *string `json:"AspectRatio,omitnil,omitempty" name:"AspectRatio"`

	// <p>生成视频时长，单位s</p><p>枚举值：3，4，5，6，7，8，9，10，其中：使用文生视频、首帧图生视频时，仅支持5和10s<br>使用视频编辑功能（“refer_type”:“base”）时，输出结果与传入视频时长相同，此时当前参数无效；此时，按输入视频时长四舍五入取整计量计费</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>为生成视频添加标识的开关，默认为1，0 需前往 控制台 申请开启显示标识自主完成方可生效。 1：添加标识； 0：不添加标识； 其他数值：默认按1处理。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>默认在生成视频的右下角添加“ AI 生成”字样，如需替换为其他的标识图片，需前往 控制台 申请开启显示标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// <p>生成视频的模式</p><p>枚举值：std，pro<br>其中std：标准模式（标准），基础模式，性价比高<br>其中pro：专家模式（高品质），高表现模式，生成视频质量更佳</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>参考视频，通过URL方式获取</p><p>可作为特征参考视频，也可作为待编辑视频，默认为待编辑视频；可选择性保留视频原声<br>通过ReferType参数区分参考视频类型：feature为特征参考视频，base为待编辑视频<br>参考视频为待编辑视频时，不能定义视频首尾帧<br>通过KeepOriginalSound参数选择是否保留视频原声，yes为保留，no为不保留；当前参数对特征参考视频（feature）也生效</p><ul><li>视频格式仅支持MP4/MOV</li><li>仅支持时长≥3秒且≤10秒的视频</li><li>视频宽高尺寸需介于720px（含）和2160px（含）之间</li><li>视频帧率基于24fps～60fps，生成视频时会输出为24fps</li><li>至多仅支持上传1段视频，视频大小不超过200MB</li></ul>
	VideoList []*ReferVideoInfo `json:"VideoList,omitnil,omitempty" name:"VideoList"`

	// <p>是否生成多镜头视频  当前参数为true时，prompt参数无效，且不支持设定首尾帧生视频  当前参数为false时，shot_type参数及multi_prompt参数无效</p>
	MultiShot *bool `json:"MultiShot,omitnil,omitempty" name:"MultiShot"`

	// <p>分镜方式  枚举值：customize  当multi_shot参数为true时，当前参数必填</p>
	ShotType *string `json:"ShotType,omitnil,omitempty" name:"ShotType"`

	// <p>各分镜信息，如提示词、时长等  通过index、prompt、duration参数定义分镜序号及相应提示词和时长，其中：  最多支持6个分镜，最小支持1个分镜  每个分镜相关内容的最大长度不超过512  每个分镜的时长不大于当前任务的总时长，不小于1  所有分镜的时长之和等于当前任务的总时长</p>
	MultiPrompt []*MultiPrompt `json:"MultiPrompt,omitnil,omitempty" name:"MultiPrompt"`

	// <p>参考主体列表  基于主体库中主体的ID配置，用key:value承载，如下：  &quot;element_list&quot;:[     {        &quot;element_id&quot;:long   },   {        &quot;element_id&quot;:long   } ] 参考主体数量与有无参考视频、参考图片数量有关，其中：  当使用首帧生成视频时，最多支持3个主体；  当使用首尾帧生成视频时，kling-v3-omni模型最多支持3个主体，kling-video-o1模不支持主体；  有参考视频时，参考图片数量和参考主体数量之和不得超过4，且不支持使用视频角色主体；  无参考视频时，参考图片数量和参考主体数量之和不得超过7；</p>
	ElementList []*Element `json:"ElementList,omitnil,omitempty" name:"ElementList"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>是否开启声音</p>
	Sound *string `json:"Sound,omitnil,omitempty" name:"Sound"`
}

func (r *SubmitVideoEditKlingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoEditKlingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "Model")
	delete(f, "ExternalTaskId")
	delete(f, "ImageList")
	delete(f, "AspectRatio")
	delete(f, "Duration")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "Mode")
	delete(f, "VideoList")
	delete(f, "MultiShot")
	delete(f, "ShotType")
	delete(f, "MultiPrompt")
	delete(f, "ElementList")
	delete(f, "CallbackUrl")
	delete(f, "Sound")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoEditKlingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoEditKlingJobResponseParams struct {
	// <p>任务ID</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoEditKlingJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoEditKlingJobResponseParams `json:"Response"`
}

func (r *SubmitVideoEditKlingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoEditKlingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoExtendKlingJobRequestParams struct {
	// <p>视频ID  支持通过文本、图片和视频延长生成的视频的ID（原视频不能超过3分钟）  请注意，基于目前的清理策略、视频生成30天之后会被清理，则无法进行延长</p>
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// <p>正向文本提示词  不能超过2500个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>负向文本提示词  不能超过2500个字符</p>
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// <p>提示词参考强度  取值范围：[0,1]，数值越大参考强度越大</p>
	CfgScale *float64 `json:"CfgScale,omitnil,omitempty" name:"CfgScale"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>为生成视频添加标识的开关，默认为1。传0 需前往  <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a>  申请开启显式标识自主完成后方可生效。<br>1：添加标识；<br>0：不添加标识；<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“ AI 生成”或“视频由 AI 生成”字样，如需替换为其他的标识图片，需前往   <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a>  申请开启显式标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitVideoExtendKlingJobRequest struct {
	*tchttp.BaseRequest
	
	// <p>视频ID  支持通过文本、图片和视频延长生成的视频的ID（原视频不能超过3分钟）  请注意，基于目前的清理策略、视频生成30天之后会被清理，则无法进行延长</p>
	VideoId *string `json:"VideoId,omitnil,omitempty" name:"VideoId"`

	// <p>正向文本提示词  不能超过2500个字符</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>负向文本提示词  不能超过2500个字符</p>
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// <p>提示词参考强度  取值范围：[0,1]，数值越大参考强度越大</p>
	CfgScale *float64 `json:"CfgScale,omitnil,omitempty" name:"CfgScale"`

	// <p>本次任务结果回调通知地址，如果配置，服务端会在任务状态发生变更时主动通知</p>
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// <p>为生成视频添加标识的开关，默认为1。传0 需前往  <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a>  申请开启显式标识自主完成后方可生效。<br>1：添加标识；<br>0：不添加标识；<br>其他数值：默认按1处理。<br>建议您使用显著标识来提示，该视频是 AI 生成的视频。</p>
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// <p>标识内容设置。<br>默认在生成视频的右下角添加“ AI 生成”或“视频由 AI 生成”字样，如需替换为其他的标识图片，需前往   <a href="https://console.cloud.tencent.com/vtc/setting">控制台</a>  申请开启显式标识自主完成。</p>
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitVideoExtendKlingJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoExtendKlingJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoId")
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "ExternalTaskId")
	delete(f, "CfgScale")
	delete(f, "CallbackUrl")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoExtendKlingJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoExtendKlingJobResponseParams struct {
	// <p>任务ID。</p>
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`


	ExternalTaskId *string `json:"ExternalTaskId,omitnil,omitempty" name:"ExternalTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoExtendKlingJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoExtendKlingJobResponseParams `json:"Response"`
}

func (r *SubmitVideoExtendKlingJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoExtendKlingJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoFaceFusionJobRequestParams struct {
	// 视频素材下载地址。用户自定义模板视频下载地址，使用前需要先调用视频审核接口进行内容审核。视频限制：分辨率≤4k，fps≤25，视频大小≤1G，时长≤20 秒，支持格式mp4。
	// 
	// 输入视频建议：
	// 姿态：人脸相对镜头水平方向角度转动不超过 90°,垂直方向角度转动不超过 20°。遮挡：脸部遮挡面积不超过 50%，不要完全遮挡五官，不要有半透明遮挡（强光，玻璃，透明眼镜等）、以及细碎离散的脸部遮挡（如飘落的花瓣）。妆容及光照：避免浓妆、复杂妆容，避免复杂光照、闪烁，这些属性无法完全恢复，并对稳定性有影响。针对特殊表情和微表情，针对局部肌肉控制下的微表情，以及过于夸张的特殊表情等不保证表情效果完全恢复。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 视频素材模板的人脸位置信息。
	// 目前最多支持融合视频素材中的 6 张人脸  
	// 输入图片要求：  
	// 1、用户图限制大小不超过 10MB  
	// 2、图片最大分辨率不超过 4k，建议最小为 128，  人脸框最小为 68
	// 3、支持格式 jpg，png  
	// 4、如果用户图中未指定人脸且有多张人脸，  默认融合最大人脸  
	// 输入图片建议：  包含上述视频中出现的人物的单人照，并且正面、清晰、无遮挡
	TemplateInfos []*FaceTemplateInfo `json:"TemplateInfos,omitnil,omitempty" name:"TemplateInfos"`

	// 用户人脸图片位置信息。
	// 输入图片要求：
	// 1、用户图限制大小不超过 10MB
	// 2、图片最大分辨率不超过 4k，建议最小为 128，人脸框最小为 68
	// 3、支持格式 jpg，png
	// 4、如果未指定人脸且用户图中有多张人脸，
	// 默认融合最大人脸
	// 输入图建议：
	// 正脸无遮挡
	MergeInfos []*FaceMergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 为生成视频添加标识的开关，默认为1。 
	// 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 视频水印Logo参数标识内容设置。   
	// 默认在融合结果图右下角添加“AI生成”类似字样，您可根据自身需要替换为其他的Logo图片。   
	// 输入建议：输入水印图片宽高需小于视频宽高
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitVideoFaceFusionJobRequest struct {
	*tchttp.BaseRequest
	
	// 视频素材下载地址。用户自定义模板视频下载地址，使用前需要先调用视频审核接口进行内容审核。视频限制：分辨率≤4k，fps≤25，视频大小≤1G，时长≤20 秒，支持格式mp4。
	// 
	// 输入视频建议：
	// 姿态：人脸相对镜头水平方向角度转动不超过 90°,垂直方向角度转动不超过 20°。遮挡：脸部遮挡面积不超过 50%，不要完全遮挡五官，不要有半透明遮挡（强光，玻璃，透明眼镜等）、以及细碎离散的脸部遮挡（如飘落的花瓣）。妆容及光照：避免浓妆、复杂妆容，避免复杂光照、闪烁，这些属性无法完全恢复，并对稳定性有影响。针对特殊表情和微表情，针对局部肌肉控制下的微表情，以及过于夸张的特殊表情等不保证表情效果完全恢复。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 视频素材模板的人脸位置信息。
	// 目前最多支持融合视频素材中的 6 张人脸  
	// 输入图片要求：  
	// 1、用户图限制大小不超过 10MB  
	// 2、图片最大分辨率不超过 4k，建议最小为 128，  人脸框最小为 68
	// 3、支持格式 jpg，png  
	// 4、如果用户图中未指定人脸且有多张人脸，  默认融合最大人脸  
	// 输入图片建议：  包含上述视频中出现的人物的单人照，并且正面、清晰、无遮挡
	TemplateInfos []*FaceTemplateInfo `json:"TemplateInfos,omitnil,omitempty" name:"TemplateInfos"`

	// 用户人脸图片位置信息。
	// 输入图片要求：
	// 1、用户图限制大小不超过 10MB
	// 2、图片最大分辨率不超过 4k，建议最小为 128，人脸框最小为 68
	// 3、支持格式 jpg，png
	// 4、如果未指定人脸且用户图中有多张人脸，
	// 默认融合最大人脸
	// 输入图建议：
	// 正脸无遮挡
	MergeInfos []*FaceMergeInfo `json:"MergeInfos,omitnil,omitempty" name:"MergeInfos"`

	// 为生成视频添加标识的开关，默认为1。 
	// 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 视频水印Logo参数标识内容设置。   
	// 默认在融合结果图右下角添加“AI生成”类似字样，您可根据自身需要替换为其他的Logo图片。   
	// 输入建议：输入水印图片宽高需小于视频宽高
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitVideoFaceFusionJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoFaceFusionJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoUrl")
	delete(f, "TemplateInfos")
	delete(f, "MergeInfos")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoFaceFusionJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoFaceFusionJobResponseParams struct {
	// 视频人脸融合任务的job id（job有效期24小时）
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoFaceFusionJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoFaceFusionJobResponseParams `json:"Response"`
}

func (r *SubmitVideoFaceFusionJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoFaceFusionJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagList struct {
	// <p>tag标签</p>
	TagId *string `json:"TagId,omitnil,omitempty" name:"TagId"`
}

type Trajectory struct {
	// 轨迹点横坐标（在像素二维坐标系下，以输入图片image左下为原点的像素坐标）
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 轨迹点纵坐标（在像素二维坐标系下，以输入图片image左下为原点的像素坐标）
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type Voice struct {
	// <p>音色id</p>
	VoiceId *string `json:"VoiceId,omitnil,omitempty" name:"VoiceId"`
}