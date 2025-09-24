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

package v20250513

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type File3D struct {
	// 文件格式
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件的Url（有效期24小时）
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 预览图片Url
	PreviewImageUrl *string `json:"PreviewImageUrl,omitnil,omitempty" name:"PreviewImageUrl"`
}

// Predefined struct for user
type QueryHunyuanTo3DJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanTo3DJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanTo3DJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanTo3DJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 生成的3D文件数组。
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanTo3DJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanTo3DJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanTo3DJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DProJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanTo3DProJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanTo3DProJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DProJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanTo3DProJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DProJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 生成的3D文件数组。
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanTo3DProJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanTo3DProJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanTo3DProJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DProJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DRapidJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanTo3DRapidJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanTo3DRapidJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DRapidJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanTo3DRapidJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DRapidJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 生成的3D文件数组。
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanTo3DRapidJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanTo3DRapidJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanTo3DRapidJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DRapidJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DJobRequestParams struct {
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持1024个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64/ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 多视角的模型图片，视角参考值：
	// left：左视图；
	// right：右视图；
	// back：后视图；
	// 
	// 每个视角仅限制一张图片。
	// ●图片大小限制：编码后大小不可超过8M。
	// ●图片分辨率限制：单边分辨率小于5000且大于128。
	// ●支持图片格式：支持jpg或png
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// 生成模型的格式，仅限制生成一种格式。
	// 生成模型文件组默认返回obj格式。
	// 可选值：OBJ，GLB，STL，USDZ，FBX，MP4。
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

type SubmitHunyuanTo3DJobRequest struct {
	*tchttp.BaseRequest
	
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持1024个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64/ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 多视角的模型图片，视角参考值：
	// left：左视图；
	// right：右视图；
	// back：后视图；
	// 
	// 每个视角仅限制一张图片。
	// ●图片大小限制：编码后大小不可超过8M。
	// ●图片分辨率限制：单边分辨率小于5000且大于128。
	// ●支持图片格式：支持jpg或png
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// 生成模型的格式，仅限制生成一种格式。
	// 生成模型文件组默认返回obj格式。
	// 可选值：OBJ，GLB，STL，USDZ，FBX，MP4。
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

func (r *SubmitHunyuanTo3DJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "MultiViewImages")
	delete(f, "ResultFormat")
	delete(f, "EnablePBR")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DJobResponseParams struct {
	// 任务ID（有效期24小时）
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DProJobRequestParams struct {
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持1024个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64/ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 多视角的模型图片，视角参考值：
	// left：左视图；
	// right：右视图；
	// back：后视图；
	// 
	// 每个视角仅限制一张图片。
	// ●图片大小限制：编码后大小不可超过8M。
	// ●图片分辨率限制：单边分辨率小于5000且大于128。
	// ●支持图片格式：支持jpg或png
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// 生成3D模型的面数，默认值为500000。
	// 可支持生成面数范围，参考值：40000-500000。
	FaceCount *int64 `json:"FaceCount,omitnil,omitempty" name:"FaceCount"`

	// 生成任务类型，默认Normal，参考值：
	// Normal：可生成带纹理的几何模型。
	// LowPoly：可生成智能减面后的模型。
	// Geometry：可生成不带纹理的几何模型（白模），选择此任务时，EnablePBR参数不生效。
	// Sketch：可输入草图或线稿图生成模型，此模式下prompt和ImageUrl/ImageBase64可一起输入。
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`
}

type SubmitHunyuanTo3DProJobRequest struct {
	*tchttp.BaseRequest
	
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持1024个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64/ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 多视角的模型图片，视角参考值：
	// left：左视图；
	// right：右视图；
	// back：后视图；
	// 
	// 每个视角仅限制一张图片。
	// ●图片大小限制：编码后大小不可超过8M。
	// ●图片分辨率限制：单边分辨率小于5000且大于128。
	// ●支持图片格式：支持jpg或png
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// 生成3D模型的面数，默认值为500000。
	// 可支持生成面数范围，参考值：40000-500000。
	FaceCount *int64 `json:"FaceCount,omitnil,omitempty" name:"FaceCount"`

	// 生成任务类型，默认Normal，参考值：
	// Normal：可生成带纹理的几何模型。
	// LowPoly：可生成智能减面后的模型。
	// Geometry：可生成不带纹理的几何模型（白模），选择此任务时，EnablePBR参数不生效。
	// Sketch：可输入草图或线稿图生成模型，此模式下prompt和ImageUrl/ImageBase64可一起输入。
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`
}

func (r *SubmitHunyuanTo3DProJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DProJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "MultiViewImages")
	delete(f, "EnablePBR")
	delete(f, "FaceCount")
	delete(f, "GenerateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DProJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DProJobResponseParams struct {
	// 任务ID（有效期24小时）
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DProJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DProJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DProJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DProJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DRapidJobRequestParams struct {
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持200个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64/ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 生成模型的格式，仅限制生成一种格式。
	// 生成模型文件组默认返回obj格式。
	// 可选值：OBJ，GLB，STL，USDZ，FBX，MP4。
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

type SubmitHunyuanTo3DRapidJobRequest struct {
	*tchttp.BaseRequest
	
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持200个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url。
	// 大小：单边分辨率要求不小于128，不大于5000。大小不超过8m（base64编码后会大30%左右，建议实际输入图片不超过6m）
	// 格式：jpg，png，jpeg，webp。
	// ImageBase64/ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 生成模型的格式，仅限制生成一种格式。
	// 生成模型文件组默认返回obj格式。
	// 可选值：OBJ，GLB，STL，USDZ，FBX，MP4。
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

func (r *SubmitHunyuanTo3DRapidJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DRapidJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ResultFormat")
	delete(f, "EnablePBR")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DRapidJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DRapidJobResponseParams struct {
	// 任务ID（有效期24小时）
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DRapidJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DRapidJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DRapidJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DRapidJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewImage struct {
	// 视角类型。
	// 取值：back、left、right
	ViewType *string `json:"ViewType,omitnil,omitempty" name:"ViewType"`

	// 图片Url地址
	ViewImageUrl *string `json:"ViewImageUrl,omitnil,omitempty" name:"ViewImageUrl"`
}