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

package v20240404

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

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

	// 任务状态码：
	// JobInit:  "初始化中"
	// JobModerationFailed: "审核失败",
	// JobRunning: "处理中",
	// JobFailed: "处理失败",
	// JobSuccess: "处理完成"。
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 任务状态描述。
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
type SubmitVideoStylizationJobRequestParams struct {
	// 风格ID，取值说明：2d_anime 2D动漫；3d_cartoon 3D卡通；3d_china 3D国潮；pixel_art	像素风。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 输入视频URL。视频要求：
	// - 视频格式：mp4、mov；
	// - 视频时长：1～60秒；
	// - 视频分辨率：540P~2056P，即长宽像素数均在540px～2056px范围内；
	// - 视频大小：不超过200M；
	// - 视频FPS：15～60fps。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`
}

type SubmitVideoStylizationJobRequest struct {
	*tchttp.BaseRequest
	
	// 风格ID，取值说明：2d_anime 2D动漫；3d_cartoon 3D卡通；3d_china 3D国潮；pixel_art	像素风。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 输入视频URL。视频要求：
	// - 视频格式：mp4、mov；
	// - 视频时长：1～60秒；
	// - 视频分辨率：540P~2056P，即长宽像素数均在540px～2056px范围内；
	// - 视频大小：不超过200M；
	// - 视频FPS：15～60fps。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitVideoStylizationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitVideoStylizationJobResponseParams struct {
	// 任务ID
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