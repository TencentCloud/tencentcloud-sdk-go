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

// Predefined struct for user
type Convert3DFormatRequestParams struct {
	// 3D文件url地址，模型文件大小≤60m，支持fbx，obj，glb格式3D文件输入。
	File3D *string `json:"File3D,omitnil,omitempty" name:"File3D"`

	// 返回的3D文件格式，参考值：STL, USDZ, FBX, MP4, GIF。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type Convert3DFormatRequest struct {
	*tchttp.BaseRequest
	
	// 3D文件url地址，模型文件大小≤60m，支持fbx，obj，glb格式3D文件输入。
	File3D *string `json:"File3D,omitnil,omitempty" name:"File3D"`

	// 返回的3D文件格式，参考值：STL, USDZ, FBX, MP4, GIF。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

func (r *Convert3DFormatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Convert3DFormatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File3D")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "Convert3DFormatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Convert3DFormatResponseParams struct {
	// 3D文件地址
	ResultFile3D *string `json:"ResultFile3D,omitnil,omitempty" name:"ResultFile3D"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type Convert3DFormatResponse struct {
	*tchttp.BaseResponse
	Response *Convert3DFormatResponseParams `json:"Response"`
}

func (r *Convert3DFormatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Convert3DFormatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHunyuanTo3DUVJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeHunyuanTo3DUVJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeHunyuanTo3DUVJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHunyuanTo3DUVJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHunyuanTo3DUVJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHunyuanTo3DUVJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功 示例值：RUN。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 生成文件的URL地址，有效期1天。
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHunyuanTo3DUVJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHunyuanTo3DUVJobResponseParams `json:"Response"`
}

func (r *DescribeHunyuanTo3DUVJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHunyuanTo3DUVJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReduceFaceJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeReduceFaceJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeReduceFaceJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReduceFaceJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReduceFaceJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReduceFaceJobResponseParams struct {
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

type DescribeReduceFaceJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReduceFaceJobResponseParams `json:"Response"`
}

func (r *DescribeReduceFaceJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReduceFaceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextureTo3DJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeTextureTo3DJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeTextureTo3DJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextureTo3DJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTextureTo3DJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTextureTo3DJobResponseParams struct {
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

type DescribeTextureTo3DJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTextureTo3DJobResponseParams `json:"Response"`
}

func (r *DescribeTextureTo3DJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTextureTo3DJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type File3D struct {
	// 文件格式
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件的Url（有效期24小时）
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 预览图片Url
	PreviewImageUrl *string `json:"PreviewImageUrl,omitnil,omitempty" name:"PreviewImageUrl"`
}

type Image struct {
	// 图片base64
	Base64 *string `json:"Base64,omitnil,omitempty" name:"Base64"`

	// 图片url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type InputFile3D struct {
	// 文件的Url（有效期24小时）
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 文件格式
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type QueryHunyuan3DPartJobRequestParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuan3DPartJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuan3DPartJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuan3DPartJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuan3DPartJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuan3DPartJobResponseParams struct {
	// 任务状态。WAIT：等待中，RUN：执行中，FAIL：任务失败，DONE：任务成功 示例值：RUN。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误码。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 错误信息。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 生成文件的URL地址，有效期1天。
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuan3DPartJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuan3DPartJobResponseParams `json:"Response"`
}

func (r *QueryHunyuan3DPartJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuan3DPartJobResponse) FromJsonString(s string) error {
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
type SubmitHunyuan3DPartJobRequestParams struct {
	// 需进行组件生成的3D模型文件，仅支持FBX格式。
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

type SubmitHunyuan3DPartJobRequest struct {
	*tchttp.BaseRequest
	
	// 需进行组件生成的3D模型文件，仅支持FBX格式。
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

func (r *SubmitHunyuan3DPartJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuan3DPartJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuan3DPartJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuan3DPartJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuan3DPartJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuan3DPartJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuan3DPartJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuan3DPartJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DProJobRequestParams struct {
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持1024个 utf-8 字符。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤6m (因base64编码后会大30%左右)
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤8m
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 多视角的模型图片，视角参考值：
	// left：左视图；
	// right：右视图；
	// back：后视图；
	// 
	// 每个视角仅限制一张图片。
	// ●图片大小限制：编码后所有图片大小总和不可超过8M。（base64编码下图片大小总和不超过6M，因base64编码后图片大小会大30%左右）
	// ●图片分辨率限制：单边分辨率小于5000且大于128。
	// ●支持图片格式：支持jpg或png
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// 生成3D模型的面数，默认值为500000。
	// 可支持生成面数范围，参考值：10000-1500000。
	// GenerateType中选择LowPoly时，参考值：3000-1500000。
	FaceCount *int64 `json:"FaceCount,omitnil,omitempty" name:"FaceCount"`

	// 生成任务类型，默认Normal，参考值：
	// Normal：可生成带纹理的几何模型。
	// LowPoly：可生成智能拓扑后的模型，选择此参数时，面数最低可到达3000面。
	// Geometry：可生成不带纹理的几何模型（白模），选择此任务时，EnablePBR参数不生效。
	// Sketch：可输入草图或线稿图生成模型，此模式下prompt和ImageUrl/ImageBase64可一起输入。
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// 该参数仅在GenerateType中选择LowPoly模式可生效。
	// 
	// 多边形类型，表示模型的表面由几边形网格构成，默认为triangle,参考值:
	// triangle: 三角形面。
	// quadrilateral: 四边形面与三角形面混合生成。
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`

	// 生成模型的格式，仅限制生成一种格式； 生成模型文件组默认返回obj、glb格式（开启时Geometry参数时，默认为glb格式）； 可选值：STL，USDZ，FBX；
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`
}

type SubmitHunyuanTo3DProJobRequest struct {
	*tchttp.BaseRequest
	
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持1024个 utf-8 字符。
	// ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤6m (因base64编码后会大30%左右)
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤8m
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 多视角的模型图片，视角参考值：
	// left：左视图；
	// right：右视图；
	// back：后视图；
	// 
	// 每个视角仅限制一张图片。
	// ●图片大小限制：编码后所有图片大小总和不可超过8M。（base64编码下图片大小总和不超过6M，因base64编码后图片大小会大30%左右）
	// ●图片分辨率限制：单边分辨率小于5000且大于128。
	// ●支持图片格式：支持jpg或png
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// 生成3D模型的面数，默认值为500000。
	// 可支持生成面数范围，参考值：10000-1500000。
	// GenerateType中选择LowPoly时，参考值：3000-1500000。
	FaceCount *int64 `json:"FaceCount,omitnil,omitempty" name:"FaceCount"`

	// 生成任务类型，默认Normal，参考值：
	// Normal：可生成带纹理的几何模型。
	// LowPoly：可生成智能拓扑后的模型，选择此参数时，面数最低可到达3000面。
	// Geometry：可生成不带纹理的几何模型（白模），选择此任务时，EnablePBR参数不生效。
	// Sketch：可输入草图或线稿图生成模型，此模式下prompt和ImageUrl/ImageBase64可一起输入。
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// 该参数仅在GenerateType中选择LowPoly模式可生效。
	// 
	// 多边形类型，表示模型的表面由几边形网格构成，默认为triangle,参考值:
	// triangle: 三角形面。
	// quadrilateral: 四边形面与三角形面混合生成。
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`

	// 生成模型的格式，仅限制生成一种格式； 生成模型文件组默认返回obj、glb格式（开启时Geometry参数时，默认为glb格式）； 可选值：STL，USDZ，FBX；
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`
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
	delete(f, "PolygonType")
	delete(f, "ResultFormat")
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
	// 文生3D, ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤6m (因base64编码后会大30%左右)
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤8m
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 生成模型的格式，仅限制生成一种格式。
	// 生成模型文件组默认返回obj格式。
	// 可选值：OBJ，GLB，STL，USDZ，FBX，MP4。
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// 是否开启单几何生成选项，开启后会生成不带纹理的3D模型（白模）； 开启时，生成模型文件不支持OBJ格式，默认生成模型文件为GLB格式。
	EnableGeometry *bool `json:"EnableGeometry,omitnil,omitempty" name:"EnableGeometry"`
}

type SubmitHunyuanTo3DRapidJobRequest struct {
	*tchttp.BaseRequest
	
	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持200个 utf-8 字符。
	// 文生3D, ImageBase64、ImageUrl和 Prompt必填其一，且Prompt和ImageBase64/ImageUrl不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 输入图 Base64 数据。
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤6m (因base64编码后会大30%左右)
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 输入图Url
	// 大小: 单边分辨率要求不小于128，不大于5000，大小≤8m
	// 格式: jpg，png，jpeg，webp.
	// lmageBase64、lmageUr和 Prompt必填其一，且Prompt和lmageBase64/mageUr不能同时存在。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 生成模型的格式，仅限制生成一种格式。
	// 生成模型文件组默认返回obj格式。
	// 可选值：OBJ，GLB，STL，USDZ，FBX，MP4。
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// 是否开启单几何生成选项，开启后会生成不带纹理的3D模型（白模）； 开启时，生成模型文件不支持OBJ格式，默认生成模型文件为GLB格式。
	EnableGeometry *bool `json:"EnableGeometry,omitnil,omitempty" name:"EnableGeometry"`
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
	delete(f, "EnableGeometry")
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

// Predefined struct for user
type SubmitHunyuanTo3DUVJobRequestParams struct {
	// 需进行UV展开的3D文件URL，可支持FBX,OBJ,GLB格式。
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

type SubmitHunyuanTo3DUVJobRequest struct {
	*tchttp.BaseRequest
	
	// 需进行UV展开的3D文件URL，可支持FBX,OBJ,GLB格式。
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

func (r *SubmitHunyuanTo3DUVJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DUVJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DUVJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DUVJobResponseParams struct {
	// 任务ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DUVJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DUVJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DUVJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DUVJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitReduceFaceJobRequestParams struct {
	// 源3D模型文件。其中参数 Type 和 Url 必填，参数 PreviewImageUrl 无意义，可忽略。
	// Type可选值：OBJ，GLB
	File3D *File3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// 多边形类型，表示模型的表面由几边形网格构成，默认为triangle,参考值:
	// triangle:三角形面。
	// quadrilateral：四边形面。
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`

	// 减面后面数档位类型，可选值：high，medium, low。
	FaceLevel *string `json:"FaceLevel,omitnil,omitempty" name:"FaceLevel"`
}

type SubmitReduceFaceJobRequest struct {
	*tchttp.BaseRequest
	
	// 源3D模型文件。其中参数 Type 和 Url 必填，参数 PreviewImageUrl 无意义，可忽略。
	// Type可选值：OBJ，GLB
	File3D *File3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// 多边形类型，表示模型的表面由几边形网格构成，默认为triangle,参考值:
	// triangle:三角形面。
	// quadrilateral：四边形面。
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`

	// 减面后面数档位类型，可选值：high，medium, low。
	FaceLevel *string `json:"FaceLevel,omitnil,omitempty" name:"FaceLevel"`
}

func (r *SubmitReduceFaceJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitReduceFaceJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File3D")
	delete(f, "PolygonType")
	delete(f, "FaceLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitReduceFaceJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitReduceFaceJobResponseParams struct {
	// 任务ID（有效期24小时）
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitReduceFaceJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitReduceFaceJobResponseParams `json:"Response"`
}

func (r *SubmitReduceFaceJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitReduceFaceJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextureTo3DJobRequestParams struct {
	// 源3D模型文件。
	// Type可选值：OBJ，GLB
	File3D *File3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持200个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 3D模型纹理参考图 Base64 数据和参考图图 Url。
	// - Base64 和 Url 必须提供一个，如果都提供以 Url 为准。
	// - 图片限制：单边分辨率小于4096且大于128，转成 Base64 字符串后小于 10MB，格式支持 jpg、jpeg、png。
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

type SubmitTextureTo3DJobRequest struct {
	*tchttp.BaseRequest
	
	// 源3D模型文件。
	// Type可选值：OBJ，GLB
	File3D *File3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// 文生3D，3D内容的描述，中文正向提示词。
	// 最多支持200个 utf-8 字符。
	// 文生3D, image、image_url和 prompt必填其一，且prompt和image/image_url不能同时存在。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 3D模型纹理参考图 Base64 数据和参考图图 Url。
	// - Base64 和 Url 必须提供一个，如果都提供以 Url 为准。
	// - 图片限制：单边分辨率小于4096且大于128，转成 Base64 字符串后小于 10MB，格式支持 jpg、jpeg、png。
	Image *Image `json:"Image,omitnil,omitempty" name:"Image"`

	// 是否开启 PBR材质生成，默认 false。
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

func (r *SubmitTextureTo3DJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextureTo3DJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File3D")
	delete(f, "Prompt")
	delete(f, "Image")
	delete(f, "EnablePBR")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTextureTo3DJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextureTo3DJobResponseParams struct {
	// 任务ID（有效期24小时）
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTextureTo3DJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTextureTo3DJobResponseParams `json:"Response"`
}

func (r *SubmitTextureTo3DJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextureTo3DJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewImage struct {
	// 视角类型。
	// 取值：back、left、right
	ViewType *string `json:"ViewType,omitnil,omitempty" name:"ViewType"`

	// 图片Url地址
	ViewImageUrl *string `json:"ViewImageUrl,omitnil,omitempty" name:"ViewImageUrl"`

	// 图片base64地址
	ViewImageBase64 *string `json:"ViewImageBase64,omitnil,omitempty" name:"ViewImageBase64"`
}