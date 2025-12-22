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

// Predefined struct for user
type CheckAnimateImageJobRequestParams struct {
	// 动作模板ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 图片base64数据。
	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 是否对输入图采用加强检测方案。
	// 
	// 默认不加强检测（false），仅对输入图做必要的基础检测。
	// 
	// 开启加强检测（true）有助于提升效果稳定性，将根据选择的动作模板提取建议的人体关键点，并判断输入图中是否包含这些人体关键点。加强检测仅对人像输入图生效，对非人输入图不生效。
	EnableBodyJoins *bool `json:"EnableBodyJoins,omitnil,omitempty" name:"EnableBodyJoins"`

	// 是否开启人脸检测。
	// 
	// 默认开启人脸检测（true），拦截主体为人像但无人脸、人脸不完整或被遮挡的输入图。可选关闭人脸检测（false）。
	EnableFace *bool `json:"EnableFace,omitnil,omitempty" name:"EnableFace"`
}

type CheckAnimateImageJobRequest struct {
	*tchttp.BaseRequest
	
	// 动作模板ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 图片base64数据。
	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 是否对输入图采用加强检测方案。
	// 
	// 默认不加强检测（false），仅对输入图做必要的基础检测。
	// 
	// 开启加强检测（true）有助于提升效果稳定性，将根据选择的动作模板提取建议的人体关键点，并判断输入图中是否包含这些人体关键点。加强检测仅对人像输入图生效，对非人输入图不生效。
	EnableBodyJoins *bool `json:"EnableBodyJoins,omitnil,omitempty" name:"EnableBodyJoins"`

	// 是否开启人脸检测。
	// 
	// 默认开启人脸检测（true），拦截主体为人像但无人脸、人脸不完整或被遮挡的输入图。可选关闭人脸检测（false）。
	EnableFace *bool `json:"EnableFace,omitnil,omitempty" name:"EnableFace"`
}

func (r *CheckAnimateImageJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAnimateImageJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "EnableBodyJoins")
	delete(f, "EnableFace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAnimateImageJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAnimateImageJobResponseParams struct {
	// 输入图是否通过校验。
	CheckPass *bool `json:"CheckPass,omitnil,omitempty" name:"CheckPass"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAnimateImageJobResponse struct {
	*tchttp.BaseResponse
	Response *CheckAnimateImageJobResponseParams `json:"Response"`
}

func (r *CheckAnimateImageJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAnimateImageJobResponse) FromJsonString(s string) error {
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
type DescribeImageAnimateJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeImageAnimateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeImageAnimateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAnimateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAnimateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAnimateJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 结果视频URL。有效期 24 小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 掩码视频链接
	MaskVideoUrl *string `json:"MaskVideoUrl,omitnil,omitempty" name:"MaskVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageAnimateJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAnimateJobResponseParams `json:"Response"`
}

func (r *DescribeImageAnimateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAnimateJobResponse) FromJsonString(s string) error {
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
type DescribeVideoEditJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeVideoEditJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeVideoEditJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoEditJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoEditJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoEditJobResponseParams struct {
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

type DescribeVideoEditJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoEditJobResponseParams `json:"Response"`
}

func (r *DescribeVideoEditJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoEditJobResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type DescribeVideoStylizationJobRequestParams struct {
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeVideoStylizationJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeVideoStylizationJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoStylizationJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoStylizationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoStylizationJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务状态码。取值说明：
	// JobInit:  "初始化中"；
	// JobModerationFailed: "审核失败"；
	// JobRunning: "处理中"；
	// JobFailed: "处理失败"；
	// JobSuccess: "处理完成"。
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 任务状态描述。取值说明：
	// JobInit:  "初始化中"；
	// JobModerationFailed: "审核失败"；
	// JobRunning: "处理中"；
	// JobFailed: "处理失败"；
	// JobSuccess: "处理完成"。
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// 处理结果视频Url。URL有效期为24小时。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil,omitempty" name:"ResultVideoUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoStylizationJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoStylizationJobResponseParams `json:"Response"`
}

func (r *DescribeVideoStylizationJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoStylizationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoVoiceJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeVideoVoiceJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeVideoVoiceJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoVoiceJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoVoiceJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoVoiceJobResponseParams struct {
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

type DescribeVideoVoiceJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoVoiceJobResponseParams `json:"Response"`
}

func (r *DescribeVideoVoiceJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoVoiceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtraParam struct {
	// 预签名的上传url，支持把视频直接传到客户指定的地址。
	UserDesignatedUrl *string `json:"UserDesignatedUrl,omitnil,omitempty" name:"UserDesignatedUrl"`
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
	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 人脸框高度。
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

type LogoParam struct {
	// 水印 Url
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 水印 Base64，Url 和 Base64 二选一传入，如果都提供以 Url 为准
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`

	// 水印图片位于生成结果图中的坐标及宽高，将按照坐标对标识图片进行位置和大小的拉伸匹配。
	LogoRect *LogoRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`
}

type LogoRect struct {
	// 水印图框X坐标值。当值大于0时，坐标轴原点位于原图左侧，方向指右；当值小于0时，坐标轴原点位于原图右侧，方向指左。
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 水印图框Y坐标值。当值大于0时，坐标轴原点位于原图上侧，方向指下；当值小于0时，坐标轴原点位于原图下侧，方向指上。
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 水印图框宽度。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 水印图框高度。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
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
type SubmitImageAnimateJobRequestParams struct {
	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 图片base64数据。
	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 动作模板ID。取值说明：ke3 科目三；tuziwu 兔子舞；huajiangwu 划桨舞。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 结果视频是否保留模板音频。默认为true
	EnableAudio *bool `json:"EnableAudio,omitnil,omitempty" name:"EnableAudio"`

	// 是否对输入图采用加强检测方案。
	// 
	// 默认不加强检测（false），仅对输入图做必要的基础检测。
	// 
	// 开启加强检测（true）有助于提升效果稳定性，将根据选择的动作模板提取建议的人体关键点，并判断输入图中是否包含这些人体关键点。加强检测仅对人像输入图生效，对非人输入图不生效。
	EnableBodyJoins *bool `json:"EnableBodyJoins,omitnil,omitempty" name:"EnableBodyJoins"`

	// 是否对结果视频背景进行分割，默认值为false。
	// true：分割结果视频，结果视频（ResultVideoUrl）将为去除背景的绿幕视频，并返回掩码视频（MaskVideoUrl）；
	// false：不分割结果视频，结果视频（ResultVideoUrl）为带背景的视频，掩码视频（MaskVideoUrl）为空字符串。
	EnableSegment *bool `json:"EnableSegment,omitnil,omitempty" name:"EnableSegment"`

	// 为生成视频添加标识的开关，默认为0。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 是否开启人脸检测。
	// 
	// 默认开启人脸检测（true），拦截主体为人像但无人脸、人脸不完整或被遮挡的输入图。可选关闭人脸检测（false）。
	EnableFace *bool `json:"EnableFace,omitnil,omitempty" name:"EnableFace"`
}

type SubmitImageAnimateJobRequest struct {
	*tchttp.BaseRequest
	
	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 图片base64数据。
	// 图片格式：支持PNG、JPG、JPEG、BMP、WEBP格式；
	// 图片分辨率：长边分辨率范围【192，4096】；
	// 图片大小：不超过10M；
	// 图片宽高比：【宽：高】数值在 1:2 到 1:1.2 范围内
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 动作模板ID。取值说明：ke3 科目三；tuziwu 兔子舞；huajiangwu 划桨舞。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 结果视频是否保留模板音频。默认为true
	EnableAudio *bool `json:"EnableAudio,omitnil,omitempty" name:"EnableAudio"`

	// 是否对输入图采用加强检测方案。
	// 
	// 默认不加强检测（false），仅对输入图做必要的基础检测。
	// 
	// 开启加强检测（true）有助于提升效果稳定性，将根据选择的动作模板提取建议的人体关键点，并判断输入图中是否包含这些人体关键点。加强检测仅对人像输入图生效，对非人输入图不生效。
	EnableBodyJoins *bool `json:"EnableBodyJoins,omitnil,omitempty" name:"EnableBodyJoins"`

	// 是否对结果视频背景进行分割，默认值为false。
	// true：分割结果视频，结果视频（ResultVideoUrl）将为去除背景的绿幕视频，并返回掩码视频（MaskVideoUrl）；
	// false：不分割结果视频，结果视频（ResultVideoUrl）为带背景的视频，掩码视频（MaskVideoUrl）为空字符串。
	EnableSegment *bool `json:"EnableSegment,omitnil,omitempty" name:"EnableSegment"`

	// 为生成视频添加标识的开关，默认为0。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 是否开启人脸检测。
	// 
	// 默认开启人脸检测（true），拦截主体为人像但无人脸、人脸不完整或被遮挡的输入图。可选关闭人脸检测（false）。
	EnableFace *bool `json:"EnableFace,omitnil,omitempty" name:"EnableFace"`
}

func (r *SubmitImageAnimateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageAnimateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	delete(f, "TemplateId")
	delete(f, "EnableAudio")
	delete(f, "EnableBodyJoins")
	delete(f, "EnableSegment")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "EnableFace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitImageAnimateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitImageAnimateJobResponseParams struct {
	// 图片跳舞任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitImageAnimateJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitImageAnimateJobResponseParams `json:"Response"`
}

func (r *SubmitImageAnimateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitImageAnimateJobResponse) FromJsonString(s string) error {
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
type SubmitTemplateToVideoJobRequestParams struct {
	// 特效模板名称。请在 [视频特效模板列表](https://cloud.tencent.com/document/product/1616/119194)  中选择想要生成的特效对应的 template 名称。
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// 参考图像，不同特效输入图片的数量详见： [视频特效模板-图片要求说明](https://cloud.tencent.com/document/product/1616/119194)
	// - 支持传入图片Base64编码或图片URL（确保可访问）
	// - 图片格式：支持png、jpg、jpeg、webp、bmp、tiff
	// - 图片文件：大小不能超过10MB（base64后），图片分辨率不小于300*300px，不大于4096*4096，图片宽高比应在1:4 ~ 4:1之间
	Images []*Image `json:"Images,omitnil,omitempty" name:"Images"`

	// 为生成视频添加标识的开关，默认为1。传0 需前往  [控制台](https://console.cloud.tencent.com/vtc/setting) 申请开启显式标识自主完成后方可生效。
	// 1：添加标识；
	// 0：不添加标识；
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成视频的右下角添加“ AI 生成”或“视频由 AI 生成”字样，如需替换为其他的标识图片，需前往  [控制台](https://console.cloud.tencent.com/vtc/setting) 申请开启显式标识自主完成。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 视频输出分辨率，默认值：360p 。不同特效支持的清晰度及消耗积分数详见：[视频特效模板-单次调用消耗积分数列](https://cloud.tencent.com/document/product/1616/119194 )
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 是否为生成的视频添加背景音乐。默认：false，  传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。
	BGM *bool `json:"BGM,omitnil,omitempty" name:"BGM"`

	// 扩展字段。
	ExtraParam *ExtraParam `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`
}

type SubmitTemplateToVideoJobRequest struct {
	*tchttp.BaseRequest
	
	// 特效模板名称。请在 [视频特效模板列表](https://cloud.tencent.com/document/product/1616/119194)  中选择想要生成的特效对应的 template 名称。
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// 参考图像，不同特效输入图片的数量详见： [视频特效模板-图片要求说明](https://cloud.tencent.com/document/product/1616/119194)
	// - 支持传入图片Base64编码或图片URL（确保可访问）
	// - 图片格式：支持png、jpg、jpeg、webp、bmp、tiff
	// - 图片文件：大小不能超过10MB（base64后），图片分辨率不小于300*300px，不大于4096*4096，图片宽高比应在1:4 ~ 4:1之间
	Images []*Image `json:"Images,omitnil,omitempty" name:"Images"`

	// 为生成视频添加标识的开关，默认为1。传0 需前往  [控制台](https://console.cloud.tencent.com/vtc/setting) 申请开启显式标识自主完成后方可生效。
	// 1：添加标识；
	// 0：不添加标识；
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成视频的右下角添加“ AI 生成”或“视频由 AI 生成”字样，如需替换为其他的标识图片，需前往  [控制台](https://console.cloud.tencent.com/vtc/setting) 申请开启显式标识自主完成。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 视频输出分辨率，默认值：360p 。不同特效支持的清晰度及消耗积分数详见：[视频特效模板-单次调用消耗积分数列](https://cloud.tencent.com/document/product/1616/119194 )
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 是否为生成的视频添加背景音乐。默认：false，  传 true 时系统将从预设 BGM 库中自动挑选合适的音乐并添加；不传或为 false 则不添加 BGM。
	BGM *bool `json:"BGM,omitnil,omitempty" name:"BGM"`

	// 扩展字段。
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
	// 任务ID。
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
type SubmitVideoEditJobRequestParams struct {
	// 输入视频
	// 
	// - 视频格式：MP4
	// - 视频时长：5s以内
	// - 视频分辨率：无限制（待验证是否可以无损输出）
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 视频内容的描述，中文正向提示词。最多支持200个 utf-8 字符（首尾空格不计入字符数）。
	// 支持风格迁移、替换、元素增加、删除控制
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 图片base64或者图片url
	// 
	// - Base64 和 Url 必须提供一个，如果都提供以Url为准。
	// - 上传图url大小不超过 8M
	// - 支持jpg，png，jpeg，webp，bmp，tiff 格式
	// - 单边分辨率不超过5000，不小于50，长宽限制1:4 ~ 4:1
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 为生成视频添加标识的开关，默认为1。 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitVideoEditJobRequest struct {
	*tchttp.BaseRequest
	
	// 输入视频
	// 
	// - 视频格式：MP4
	// - 视频时长：5s以内
	// - 视频分辨率：无限制（待验证是否可以无损输出）
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 视频内容的描述，中文正向提示词。最多支持200个 utf-8 字符（首尾空格不计入字符数）。
	// 支持风格迁移、替换、元素增加、删除控制
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 图片base64或者图片url
	// 
	// - Base64 和 Url 必须提供一个，如果都提供以Url为准。
	// - 上传图url大小不超过 8M
	// - 支持jpg，png，jpeg，webp，bmp，tiff 格式
	// - 单边分辨率不超过5000，不小于50，长宽限制1:4 ~ 4:1
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 为生成视频添加标识的开关，默认为1。 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitVideoEditJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoEditJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoUrl")
	delete(f, "Prompt")
	delete(f, "Image")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoEditJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoEditJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoEditJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoEditJobResponseParams `json:"Response"`
}

func (r *SubmitVideoEditJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoEditJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoFaceFusionJobRequestParams struct {
	// 视频素材下载地址。用户自定义模版视频下载地址，使用前需要先调用视频审核接口进行内容审核。视频限制：分辨率≤4k，fps≤25，视频大小≤1G，时长≤20 秒，支持格式mp4。
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
	
	// 视频素材下载地址。用户自定义模版视频下载地址，使用前需要先调用视频审核接口进行内容审核。视频限制：分辨率≤4k，fps≤25，视频大小≤1G，时长≤20 秒，支持格式mp4。
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

// Predefined struct for user
type SubmitVideoStylizationJobRequestParams struct {
	// 风格ID。取值说明：
	// 2d_anime：2D动漫；
	// 3d_cartoon：3D卡通；
	// 3d_china：3D国潮；
	// pixel_art：像素风。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 输入视频URL。视频要求：
	// - 视频格式：mp4、mov；
	// - 视频时长：1～60秒；
	// - 视频分辨率：540P~2056P，即长宽像素数均在540px～2056px范围内；
	// - 视频大小：不超过200M；
	// - 视频FPS：15～60fps。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 风格化强度。取值说明：
	// low：风格化强度弱；
	// medium：风格化强度中等；
	// high：风格化强度强。
	// 默认值为medium。
	StyleStrength *string `json:"StyleStrength,omitnil,omitempty" name:"StyleStrength"`
}

type SubmitVideoStylizationJobRequest struct {
	*tchttp.BaseRequest
	
	// 风格ID。取值说明：
	// 2d_anime：2D动漫；
	// 3d_cartoon：3D卡通；
	// 3d_china：3D国潮；
	// pixel_art：像素风。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 输入视频URL。视频要求：
	// - 视频格式：mp4、mov；
	// - 视频时长：1～60秒；
	// - 视频分辨率：540P~2056P，即长宽像素数均在540px～2056px范围内；
	// - 视频大小：不超过200M；
	// - 视频FPS：15～60fps。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 风格化强度。取值说明：
	// low：风格化强度弱；
	// medium：风格化强度中等；
	// high：风格化强度强。
	// 默认值为medium。
	StyleStrength *string `json:"StyleStrength,omitnil,omitempty" name:"StyleStrength"`
}

func (r *SubmitVideoStylizationJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoStylizationJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StyleId")
	delete(f, "VideoUrl")
	delete(f, "StyleStrength")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoStylizationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoStylizationJobResponseParams struct {
	// 任务ID。任务有效期为48小时。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoStylizationJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoStylizationJobResponseParams `json:"Response"`
}

func (r *SubmitVideoStylizationJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoStylizationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoVoiceJobRequestParams struct {
	// 输入视频的Url  上传视频时长限制：1-15s 视频格式：MP4，MOV 视频大小：不超过1 GB URL地址中不能包含中文字符。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 描述音效内容的正向提示词。输入上限50个字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 音效内容的原始负向提示词。输入上限50个字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 为生成视频添加标识的开关，默认为1。 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

type SubmitVideoVoiceJobRequest struct {
	*tchttp.BaseRequest
	
	// 输入视频的Url  上传视频时长限制：1-15s 视频格式：MP4，MOV 视频大小：不超过1 GB URL地址中不能包含中文字符。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 描述音效内容的正向提示词。输入上限50个字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 音效内容的原始负向提示词。输入上限50个字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 为生成视频添加标识的开关，默认为1。 1：添加标识。 0：不添加标识。 其他数值：默认按1处理。 建议您使用显著标识来提示，该视频是 AI 生成的视频。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 默认在生成视频的右下角添加“视频由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`
}

func (r *SubmitVideoVoiceJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoVoiceJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoUrl")
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoVoiceJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoVoiceJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitVideoVoiceJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitVideoVoiceJobResponseParams `json:"Response"`
}

func (r *SubmitVideoVoiceJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitVideoVoiceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}