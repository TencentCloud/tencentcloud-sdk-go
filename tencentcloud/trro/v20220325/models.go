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

package v20220325

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type BatchDeleteDevicesRequestParams struct {
	// 目标删除设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标删除设备的设备ID数组
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

type BatchDeleteDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 目标删除设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标删除设备的设备ID数组
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

func (r *BatchDeleteDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteDevicesResponseParams struct {
	// 删除失败的设备ID列表
	FailedDeviceIds []*string `json:"FailedDeviceIds,omitnil,omitempty" name:"FailedDeviceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeleteDevicesResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteDevicesResponseParams `json:"Response"`
}

func (r *BatchDeleteDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeletePolicyRequestParams struct {
	// 删除权限配置的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 删除权限配置的远端设备ID列表
	RemoteDeviceIds []*string `json:"RemoteDeviceIds,omitnil,omitempty" name:"RemoteDeviceIds"`

	// 删除权限配置的权限模式, black为黑名单，white为白名单
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`
}

type BatchDeletePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 删除权限配置的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 删除权限配置的远端设备ID列表
	RemoteDeviceIds []*string `json:"RemoteDeviceIds,omitnil,omitempty" name:"RemoteDeviceIds"`

	// 删除权限配置的权限模式, black为黑名单，white为白名单
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`
}

func (r *BatchDeletePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeletePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RemoteDeviceIds")
	delete(f, "PolicyMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeletePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeletePolicyResponseParams struct {
	// 删除权限配置失败的远端设备ID列表
	FailedRemoteDeviceIds []*string `json:"FailedRemoteDeviceIds,omitnil,omitempty" name:"FailedRemoteDeviceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeletePolicyResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeletePolicyResponseParams `json:"Response"`
}

func (r *BatchDeletePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BoundLicensesRequestParams struct {
	// license数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BoundLicensesRequest struct {
	*tchttp.BaseRequest
	
	// license数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *BoundLicensesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BoundLicensesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Count")
	delete(f, "DeviceId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BoundLicensesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BoundLicensesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BoundLicensesResponse struct {
	*tchttp.BaseResponse
	Response *BoundLicensesResponseParams `json:"Response"`
}

func (r *BoundLicensesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BoundLicensesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudStorage struct {
	// 腾讯云对象存储COS以及第三方云存储账号信息
	// 0：腾讯云对象存储 COS
	// 1：AWS
	// 【注意】目前第三方云存储仅支持AWS，更多第三方云存储陆续支持中
	// 示例值：0
	Vendor *int64 `json:"Vendor,omitnil,omitempty" name:"Vendor"`

	// 腾讯云对象存储的[地域信息]（https://cloud.tencent.com/document/product/436/6224#.E5.9C.B0.E5.9F.9F）。
	// 示例值：cn-shanghai-1
	// 
	// AWS S3[地域信息]（https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-regions）
	// 示例值：ap-shanghai(cos, 具体参考云存储厂商支持的地域)
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 云存储桶名称。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 云存储的access_key账号信息。
	// 若存储至腾讯云对象存储COS，请前往https://console.cloud.tencent.com/cam/capi 查看或创建，对应链接中密钥字段的SecretId值。
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 云存储的secret_key账号信息。
	// 若存储至腾讯云对象存储COS，请前往https://console.cloud.tencent.com/cam/capi 查看或创建，对应链接中密钥字段的SecretKey值。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 云存储bucket 的指定位置，由字符串数组组成。合法的字符串范围az,AZ,0~9,'_'和'-'，举个例子，录制文件xxx.m3u8在 ["prefix1", "prefix2"]作用下，会变成prefix1/prefix2/TaskId/xxx.m3u8。
	FileNamePrefix []*string `json:"FileNamePrefix,omitnil,omitempty" name:"FileNamePrefix"`
}

// Predefined struct for user
type CreateCloudRecordingRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 视频流号
	VideoStreamId *int64 `json:"VideoStreamId,omitnil,omitempty" name:"VideoStreamId"`

	// 腾讯云对象存储COS以及第三方云存储的账号信息
	CloudStorage *CloudStorage `json:"CloudStorage,omitnil,omitempty" name:"CloudStorage"`

	// 如果是aac或者mp4文件格式，超过长度限制后，系统会自动拆分视频文件。单位：分钟。默认为1440min（24h），取值范围为1-1440。【单文件限制最大为2G，满足文件大小 >2G 或录制时长度 > 24h任意一个条件，文件都会自动切分】 Hls 格式录制此参数不生效。
	MaxMediaFileDuration *int64 `json:"MaxMediaFileDuration,omitnil,omitempty" name:"MaxMediaFileDuration"`

	// 输出文件的格式（存储至COS等第三方存储时有效）。0：输出文件为hls格式。1：输出文件格式为hls+mp4。2：输出文件格式为hls+aac 。3：(默认)输出文件格式为mp4。4：输出文件格式为aac。
	OutputFormat *int64 `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// 房间内持续没有主播的状态超过MaxIdleTime的时长，自动停止录制，单位：秒。默认值为 30 秒，该值需大于等于 5秒，且小于等于 86400秒(24小时)。 示例值：30
	MaxIdleTime *int64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`
}

type CreateCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 视频流号
	VideoStreamId *int64 `json:"VideoStreamId,omitnil,omitempty" name:"VideoStreamId"`

	// 腾讯云对象存储COS以及第三方云存储的账号信息
	CloudStorage *CloudStorage `json:"CloudStorage,omitnil,omitempty" name:"CloudStorage"`

	// 如果是aac或者mp4文件格式，超过长度限制后，系统会自动拆分视频文件。单位：分钟。默认为1440min（24h），取值范围为1-1440。【单文件限制最大为2G，满足文件大小 >2G 或录制时长度 > 24h任意一个条件，文件都会自动切分】 Hls 格式录制此参数不生效。
	MaxMediaFileDuration *int64 `json:"MaxMediaFileDuration,omitnil,omitempty" name:"MaxMediaFileDuration"`

	// 输出文件的格式（存储至COS等第三方存储时有效）。0：输出文件为hls格式。1：输出文件格式为hls+mp4。2：输出文件格式为hls+aac 。3：(默认)输出文件格式为mp4。4：输出文件格式为aac。
	OutputFormat *int64 `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// 房间内持续没有主播的状态超过MaxIdleTime的时长，自动停止录制，单位：秒。默认值为 30 秒，该值需大于等于 5秒，且小于等于 86400秒(24小时)。 示例值：30
	MaxIdleTime *int64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`
}

func (r *CreateCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	delete(f, "VideoStreamId")
	delete(f, "CloudStorage")
	delete(f, "MaxMediaFileDuration")
	delete(f, "OutputFormat")
	delete(f, "MaxIdleTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudRecordingResponseParams struct {
	// 云录制服务分配的任务 ID。任务 ID 是对一次录制生命周期过程的唯一标识，结束录制时会失去意义。任务 ID需要业务保存下来，作为下次针对这个录制任务操作的参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudRecordingResponseParams `json:"Response"`
}

func (r *CreateCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceRequestParams struct {
	// 创建设备所归属的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建设备ID，项目内需要唯一，由小写英文字母、数字和下划线构成，长度不超过18
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 创建设备名称，长度小于24, 可包含中文、数字、英文字母和下划线
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备类型，field为现场设备（受控设备），remote为远端设备（操控设备），不填默认为field
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备认证口令，由大小写英文字母和数字构成，须为16位
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
}

type CreateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 创建设备所归属的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建设备ID，项目内需要唯一，由小写英文字母、数字和下划线构成，长度不超过18
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 创建设备名称，长度小于24, 可包含中文、数字、英文字母和下划线
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备类型，field为现场设备（受控设备），remote为远端设备（操控设备），不填默认为field
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备认证口令，由大小写英文字母和数字构成，须为16位
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
}

func (r *CreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	delete(f, "DeviceName")
	delete(f, "DeviceType")
	delete(f, "DeviceToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceResponseParams `json:"Response"`
}

func (r *CreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 项目名称，长度不超过24个字符
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述，长度不超过120个字符，不填默认为空
	ProjectDescription *string `json:"ProjectDescription,omitnil,omitempty" name:"ProjectDescription"`

	// 权限模式，black为黑名单，white为白名单，不填默认为black
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称，长度不超过24个字符
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述，长度不超过120个字符，不填默认为空
	ProjectDescription *string `json:"ProjectDescription,omitnil,omitempty" name:"ProjectDescription"`

	// 权限模式，black为黑名单，white为白名单，不填默认为black
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectName")
	delete(f, "ProjectDescription")
	delete(f, "PolicyMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 项目ID，长度为16位
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRecordingRequestParams struct {
	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// 录制任务的唯一Id，在启动录制成功后会返回。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRecordingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudRecordingResponseParams `json:"Response"`
}

func (r *DeleteCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceInfoRequestParams struct {
	// 目标设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DescribeDeviceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 目标设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *DescribeDeviceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceInfoResponseParams struct {
	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备类型，field为现场设备（被控方），remote为远端设备（操控方）
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备状态，offline为离线，ready为在线准备，connected为会话中
	DeviceStatus *string `json:"DeviceStatus,omitnil,omitempty" name:"DeviceStatus"`

	// 设备状态最后更新时间
	LastReportTime *string `json:"LastReportTime,omitnil,omitempty" name:"LastReportTime"`

	// 设备信息最后修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceInfoResponseParams `json:"Response"`
}

func (r *DescribeDeviceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListRequestParams struct {
	// <p>设备所属项目ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>设备类型筛选，不填默认为全部设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>对设备ID或Name按关键字进行模糊匹配，不填则不进行模糊匹配</p>
	SearchWords *string `json:"SearchWords,omitnil,omitempty" name:"SearchWords"`

	// <p>每页返回的最大设备数，不填默认为10</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>当前页码，不填默认为1（首页）</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>设备状态筛选，不填默认为不过滤。取值：[&quot;ready&quot;,&quot;connected&quot;,&quot;online&quot;]，online代表ready或connected</p>
	DeviceStatus *string `json:"DeviceStatus,omitnil,omitempty" name:"DeviceStatus"`

	// <p>标识查询项目下的设备注册类型，默认不包含免注册登录设备。 若存在免注册登录设备，该参数传&quot;1&quot;</p><p>枚举值：</p><ul><li>0： 项目不包含免注册登录设备</li><li>1： 项目包含免注册登录设备</li></ul><p>默认值：0</p>
	RegisterType *int64 `json:"RegisterType,omitnil,omitempty" name:"RegisterType"`
}

type DescribeDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// <p>设备所属项目ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>设备类型筛选，不填默认为全部设备类型</p>
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// <p>对设备ID或Name按关键字进行模糊匹配，不填则不进行模糊匹配</p>
	SearchWords *string `json:"SearchWords,omitnil,omitempty" name:"SearchWords"`

	// <p>每页返回的最大设备数，不填默认为10</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>当前页码，不填默认为1（首页）</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>设备状态筛选，不填默认为不过滤。取值：[&quot;ready&quot;,&quot;connected&quot;,&quot;online&quot;]，online代表ready或connected</p>
	DeviceStatus *string `json:"DeviceStatus,omitnil,omitempty" name:"DeviceStatus"`

	// <p>标识查询项目下的设备注册类型，默认不包含免注册登录设备。 若存在免注册登录设备，该参数传&quot;1&quot;</p><p>枚举值：</p><ul><li>0： 项目不包含免注册登录设备</li><li>1： 项目包含免注册登录设备</li></ul><p>默认值：0</p>
	RegisterType *int64 `json:"RegisterType,omitnil,omitempty" name:"RegisterType"`
}

func (r *DescribeDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceType")
	delete(f, "SearchWords")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "DeviceStatus")
	delete(f, "RegisterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListResponseParams struct {
	// <p>设备信息列表</p>
	Devices []*DeviceInfo `json:"Devices,omitnil,omitempty" name:"Devices"`

	// <p>设备总数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>本次返回的设备数</p>
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceListResponseParams `json:"Response"`
}

func (r *DescribeDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceSessionDetailsRequestParams struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeDeviceSessionDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeDeviceSessionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceSessionDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceSessionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceSessionDetailsResponseParams struct {
	// 按设备区分的会话详细数据
	Details []*SessionDeviceDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceSessionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceSessionDetailsResponseParams `json:"Response"`
}

func (r *DescribeDeviceSessionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceSessionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceSessionListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDeviceSessionListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeDeviceSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "DeviceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceSessionListResponseParams struct {
	// 总个数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 会话列表
	DeviceSessionList []*SessionInfo `json:"DeviceSessionList,omitnil,omitempty" name:"DeviceSessionList"`

	// 本页数量
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceSessionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceSessionListResponseParams `json:"Response"`
}

func (r *DescribeDeviceSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyRequestParams struct {
	// 查看权限的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查看的权限模式，black为黑名单，white为白名单，不填默认为当前项目生效的权限模式
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`

	// 模糊匹配模式，remoteMatch为远端设备ID匹配，fieldMatch为现场ID匹配，不填默认为remoteMatch
	SearchMode *string `json:"SearchMode,omitnil,omitempty" name:"SearchMode"`

	// 模糊匹配关键字，不填默认不进行模糊匹配
	SearchWords *string `json:"SearchWords,omitnil,omitempty" name:"SearchWords"`

	// 每页返回的最大数量，不填默认为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页码，不填默认为1（首页）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 查看权限的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查看的权限模式，black为黑名单，white为白名单，不填默认为当前项目生效的权限模式
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`

	// 模糊匹配模式，remoteMatch为远端设备ID匹配，fieldMatch为现场ID匹配，不填默认为remoteMatch
	SearchMode *string `json:"SearchMode,omitnil,omitempty" name:"SearchMode"`

	// 模糊匹配关键字，不填默认不进行模糊匹配
	SearchWords *string `json:"SearchWords,omitnil,omitempty" name:"SearchWords"`

	// 每页返回的最大数量，不填默认为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页码，不填默认为1（首页）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PolicyMode")
	delete(f, "SearchMode")
	delete(f, "SearchWords")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyResponseParams struct {
	// 权限模式
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`

	// 返回的权限模式是否为当前生效的权限模式
	PolicyEnabled *bool `json:"PolicyEnabled,omitnil,omitempty" name:"PolicyEnabled"`

	// 权限信息列表
	PolicyInfo []*PolicyInfo `json:"PolicyInfo,omitnil,omitempty" name:"PolicyInfo"`

	// 本次返回的权限信息数量
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 权限信息总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyResponseParams `json:"Response"`
}

func (r *DescribePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectInfoRequestParams struct {
	// 目标项目ID，必填参数
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeProjectInfoRequest struct {
	*tchttp.BaseRequest
	
	// 目标项目ID，必填参数
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectInfoResponseParams struct {
	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDescription *string `json:"ProjectDescription,omitnil,omitempty" name:"ProjectDescription"`

	// 项目权限模式，black为黑名单，white为白名单
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`

	// 项目信息修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectInfoResponseParams `json:"Response"`
}

func (r *DescribeProjectInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectListRequestParams struct {
	// 每页返回的最大项目数量，不填默认为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页码，不填默认为1（首页）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeProjectListRequest struct {
	*tchttp.BaseRequest
	
	// 每页返回的最大项目数量，不填默认为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页码，不填默认为1（首页）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectListResponseParams struct {
	// 项目信息数组
	Projects []*ProjectInfo `json:"Projects,omitnil,omitempty" name:"Projects"`

	// 项目总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 本次返回的项目数
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectListResponseParams `json:"Response"`
}

func (r *DescribeProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecentSessionListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 设备ID，支持过滤远端设备或现场设备
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 时间范围的起始时间。时间范围最大为最近两小时，若不传或超出范围，则起始时间按两小时前计算
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间范围的结束时间。时间范围最大为最近两小时，若不传或超出范围，则结束时间按当前时间计算
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeRecentSessionListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码，从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 设备ID，支持过滤远端设备或现场设备
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 时间范围的起始时间。时间范围最大为最近两小时，若不传或超出范围，则起始时间按两小时前计算
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间范围的结束时间。时间范围最大为最近两小时，若不传或超出范围，则结束时间按当前时间计算
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeRecentSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecentSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "DeviceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecentSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecentSessionListResponseParams struct {
	// 总个数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 会话列表
	RecentSessionList []*RecentSessionInfo `json:"RecentSessionList,omitnil,omitempty" name:"RecentSessionList"`

	// 本页数量
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecentSessionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecentSessionListResponseParams `json:"Response"`
}

func (r *DescribeRecentSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecentSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionStatisticsByIntervalRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 统计时间间隔：hour|day|month
	StatisticInterval *string `json:"StatisticInterval,omitnil,omitempty" name:"StatisticInterval"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 起始时间，单位：秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，单位：秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeSessionStatisticsByIntervalRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 统计时间间隔：hour|day|month
	StatisticInterval *string `json:"StatisticInterval,omitnil,omitempty" name:"StatisticInterval"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 起始时间，单位：秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，单位：秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeSessionStatisticsByIntervalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionStatisticsByIntervalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "StatisticInterval")
	delete(f, "DeviceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionStatisticsByIntervalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionStatisticsByIntervalResponseParams struct {
	// 各时间段的会话统计数据
	SessionStatistics []*SessionIntervalStatistic `json:"SessionStatistics,omitnil,omitempty" name:"SessionStatistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionStatisticsByIntervalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionStatisticsByIntervalResponseParams `json:"Response"`
}

func (r *DescribeSessionStatisticsByIntervalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionStatisticsByIntervalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionStatisticsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 起始时间，单位：秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，单位：秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeSessionStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 起始时间，单位：秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，单位：秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeSessionStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionStatisticsResponseParams struct {
	// 会话数量
	SessionNum *uint64 `json:"SessionNum,omitnil,omitempty" name:"SessionNum"`

	// 通话时长，单位：分钟
	TotalDuration *uint64 `json:"TotalDuration,omitnil,omitempty" name:"TotalDuration"`

	// 活跃现场设备数
	ActiveFieldDeviceNum *uint64 `json:"ActiveFieldDeviceNum,omitnil,omitempty" name:"ActiveFieldDeviceNum"`

	// 活跃远端设备数
	ActiveRemoteDeviceNum *uint64 `json:"ActiveRemoteDeviceNum,omitnil,omitempty" name:"ActiveRemoteDeviceNum"`

	// 优良会话占比，单位：%
	NotBadSessionRatio *uint64 `json:"NotBadSessionRatio,omitnil,omitempty" name:"NotBadSessionRatio"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionStatisticsResponseParams `json:"Response"`
}

func (r *DescribeSessionStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Device struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 已经绑定license数量
	LicenseCount *int64 `json:"LicenseCount,omitnil,omitempty" name:"LicenseCount"`

	// 剩余天数：天
	RemainDay *int64 `json:"RemainDay,omitnil,omitempty" name:"RemainDay"`

	// 过期时间：s
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 服务时长：s
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 已经绑定licenseId列表
	LicenseIds []*string `json:"LicenseIds,omitnil,omitempty" name:"LicenseIds"`

	// 每月license的限定时长
	MonthlyRemainTime *int64 `json:"MonthlyRemainTime,omitnil,omitempty" name:"MonthlyRemainTime"`

	// 月封顶时长（分钟)
	LimitedTime *int64 `json:"LimitedTime,omitnil,omitempty" name:"LimitedTime"`
}

type DeviceInfo struct {
	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 设备状态，offline为离线，ready为在线准备，connected为会话中
	DeviceStatus *string `json:"DeviceStatus,omitnil,omitempty" name:"DeviceStatus"`

	// 设备类型，field为现场设备（受控方），remote为远端设备（操控方）
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 设备信息最近修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 设备状态最近更新时间
	LastReportTime *string `json:"LastReportTime,omitnil,omitempty" name:"LastReportTime"`

	// 设备所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DurationDetails struct {
	// 会话时间
	SessionTime *string `json:"SessionTime,omitnil,omitempty" name:"SessionTime"`

	// 语音:min
	Voice *int64 `json:"Voice,omitnil,omitempty" name:"Voice"`

	// 标清:min
	SD *int64 `json:"SD,omitnil,omitempty" name:"SD"`

	// 高清:min
	HD *int64 `json:"HD,omitnil,omitempty" name:"HD"`

	// 超高清:min
	FHD *int64 `json:"FHD,omitnil,omitempty" name:"FHD"`

	// 2k:min
	TwoK *int64 `json:"TwoK,omitnil,omitempty" name:"TwoK"`

	// 4k:min
	FourK *int64 `json:"FourK,omitnil,omitempty" name:"FourK"`

	// 在线时长:min
	Online *int64 `json:"Online,omitnil,omitempty" name:"Online"`

	// 多网标清:min
	MultiSD *int64 `json:"MultiSD,omitnil,omitempty" name:"MultiSD"`

	// 多网高清:min
	MultiHD *int64 `json:"MultiHD,omitnil,omitempty" name:"MultiHD"`

	// 多网超高清:min
	MultiFHD *int64 `json:"MultiFHD,omitnil,omitempty" name:"MultiFHD"`

	// 多网2k:min
	MultiTwoK *int64 `json:"MultiTwoK,omitnil,omitempty" name:"MultiTwoK"`

	// 多网4k:min
	MultiFourK *int64 `json:"MultiFourK,omitnil,omitempty" name:"MultiFourK"`

	// 多网在线时长:min
	MultiOnline *int64 `json:"MultiOnline,omitnil,omitempty" name:"MultiOnline"`

	// 抵扣时长:min
	DeductDuration *int64 `json:"DeductDuration,omitnil,omitempty" name:"DeductDuration"`
}

// Predefined struct for user
type GetDeviceLicenseRequestParams struct {
	// 目标设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type GetDeviceLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 目标设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *GetDeviceLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeviceLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeviceLicenseResponseParams struct {
	// 指定设备已经绑定的可用license数量
	AvailableCount *int64 `json:"AvailableCount,omitnil,omitempty" name:"AvailableCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDeviceLicenseResponse struct {
	*tchttp.BaseResponse
	Response *GetDeviceLicenseResponseParams `json:"Response"`
}

func (r *GetDeviceLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeviceLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicesRequestParams struct {
	// 页码
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目 ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type GetDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目 ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *GetDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDevicesResponseParams struct {
	// 设备授权列表
	Devices []*Device `json:"Devices,omitnil,omitempty" name:"Devices"`

	// 列表数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDevicesResponse struct {
	*tchttp.BaseResponse
	Response *GetDevicesResponseParams `json:"Response"`
}

func (r *GetDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDurationDetailsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备ID，不传查全部设备
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type GetDurationDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备ID，不传查全部设备
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *GetDurationDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDurationDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDurationDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDurationDetailsResponseParams struct {
	// 列表总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 时长明细列表
	DurationDetails []*DurationDetails `json:"DurationDetails,omitnil,omitempty" name:"DurationDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDurationDetailsResponse struct {
	*tchttp.BaseResponse
	Response *GetDurationDetailsResponseParams `json:"Response"`
}

func (r *GetDurationDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDurationDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLicenseStatRequestParams struct {

}

type GetLicenseStatRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetLicenseStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLicenseStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLicenseStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLicenseStatResponseParams struct {
	// 有效授权
	Valid *int64 `json:"Valid,omitnil,omitempty" name:"Valid"`

	// 已绑定授权
	Bound *int64 `json:"Bound,omitnil,omitempty" name:"Bound"`

	// 未绑定授权
	UnBound *int64 `json:"UnBound,omitnil,omitempty" name:"UnBound"`

	// 过期授权
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`

	// 当月用量超时授权个数
	MonthlyExpire *int64 `json:"MonthlyExpire,omitnil,omitempty" name:"MonthlyExpire"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLicenseStatResponse struct {
	*tchttp.BaseResponse
	Response *GetLicenseStatResponseParams `json:"Response"`
}

func (r *GetLicenseStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLicenseStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLicensesRequestParams struct {
	// 页码
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// projectId
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// DeviceId
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// license状态：0:未绑定；1:已绑定；2:已停服；3:已退费
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type GetLicensesRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 页面数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// projectId
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// DeviceId
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// license状态：0:未绑定；1:已绑定；2:已停服；3:已退费
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *GetLicensesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLicensesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLicensesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLicensesResponseParams struct {
	// license列表
	Licenses []*License `json:"Licenses,omitnil,omitempty" name:"Licenses"`

	// license列表项数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLicensesResponse struct {
	*tchttp.BaseResponse
	Response *GetLicensesResponseParams `json:"Response"`
}

func (r *GetLicensesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLicensesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTotalDurationRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备id，不传查全部
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type GetTotalDurationRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备id，不传查全部
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *GetTotalDurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTotalDurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTotalDurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTotalDurationResponseParams struct {
	// 语音:min
	Voice *int64 `json:"Voice,omitnil,omitempty" name:"Voice"`

	// 标清:min
	SD *int64 `json:"SD,omitnil,omitempty" name:"SD"`

	// 高清:min
	HD *int64 `json:"HD,omitnil,omitempty" name:"HD"`

	// 超高清:min
	FHD *int64 `json:"FHD,omitnil,omitempty" name:"FHD"`

	// 2k:min
	TwoK *int64 `json:"TwoK,omitnil,omitempty" name:"TwoK"`

	// 4k:min
	FourK *int64 `json:"FourK,omitnil,omitempty" name:"FourK"`

	// 在线时长:min 
	Online *int64 `json:"Online,omitnil,omitempty" name:"Online"`

	// 多网标清:min
	MultiSD *int64 `json:"MultiSD,omitnil,omitempty" name:"MultiSD"`

	// 多网高清:min
	MultiHD *int64 `json:"MultiHD,omitnil,omitempty" name:"MultiHD"`

	// 多网超高清:min
	MultiFHD *int64 `json:"MultiFHD,omitnil,omitempty" name:"MultiFHD"`

	// 多网2k:min
	MultiTwoK *int64 `json:"MultiTwoK,omitnil,omitempty" name:"MultiTwoK"`

	// 多网4k:min
	MultiFourK *int64 `json:"MultiFourK,omitnil,omitempty" name:"MultiFourK"`

	// 多网在线时长:min 
	MultiOnline *int64 `json:"MultiOnline,omitnil,omitempty" name:"MultiOnline"`

	// 总抵扣时长:min 
	DeductDuration *int64 `json:"DeductDuration,omitnil,omitempty" name:"DeductDuration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTotalDurationResponse struct {
	*tchttp.BaseResponse
	Response *GetTotalDurationResponseParams `json:"Response"`
}

func (r *GetTotalDurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTotalDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type License struct {
	// 该类型的license个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// license状态：0:未绑定；1:已绑定；2:已停服；3:已退费
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 到期时间戳：s
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 服务时长：s
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 剩余天数：天
	RemainDay *int64 `json:"RemainDay,omitnil,omitempty" name:"RemainDay"`

	// 该类型的licenseId列表
	LicenseIds []*string `json:"LicenseIds,omitnil,omitempty" name:"LicenseIds"`
}

// Predefined struct for user
type ModifyCallbackUrlRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 回调URL
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 回调签名密钥，用于校验回调信息的完整性
	SignKey *string `json:"SignKey,omitnil,omitempty" name:"SignKey"`
}

type ModifyCallbackUrlRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 回调URL
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 回调签名密钥，用于校验回调信息的完整性
	SignKey *string `json:"SignKey,omitnil,omitempty" name:"SignKey"`
}

func (r *ModifyCallbackUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCallbackUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CallbackUrl")
	delete(f, "SignKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCallbackUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCallbackUrlResponseParams struct {
	// 响应码，0：成功，其他：失败
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 响应消息，ok:成功，其他：失败
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCallbackUrlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCallbackUrlResponseParams `json:"Response"`
}

func (r *ModifyCallbackUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCallbackUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceRequestParams struct {
	// 要修改设备归属项目的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 要修改设备的设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 修改后的设备名称，不填则不修改
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 修改后的设备认证口令，不填则不修改
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 要修改设备归属项目的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 要修改设备的设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 修改后的设备名称，不填则不修改
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 修改后的设备认证口令，不填则不修改
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
}

func (r *ModifyDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeviceId")
	delete(f, "DeviceName")
	delete(f, "DeviceToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceResponseParams `json:"Response"`
}

func (r *ModifyDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPolicyRequestParams struct {
	// 修改权限配置的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 修改权限配置的远端设备ID
	RemoteDeviceId *string `json:"RemoteDeviceId,omitnil,omitempty" name:"RemoteDeviceId"`

	// 权限修改涉及的现场设备ID数组
	FieldDeviceIds []*string `json:"FieldDeviceIds,omitnil,omitempty" name:"FieldDeviceIds"`

	// 修改的目标权限模式，black为黑名单，white为白名单
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`

	// 修改模式，add为新增（添加现场设备I关联），remove为删除（解除现场设备关联），set为设置（更新现场设备关联）
	ModifyMode *string `json:"ModifyMode,omitnil,omitempty" name:"ModifyMode"`
}

type ModifyPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 修改权限配置的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 修改权限配置的远端设备ID
	RemoteDeviceId *string `json:"RemoteDeviceId,omitnil,omitempty" name:"RemoteDeviceId"`

	// 权限修改涉及的现场设备ID数组
	FieldDeviceIds []*string `json:"FieldDeviceIds,omitnil,omitempty" name:"FieldDeviceIds"`

	// 修改的目标权限模式，black为黑名单，white为白名单
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`

	// 修改模式，add为新增（添加现场设备I关联），remove为删除（解除现场设备关联），set为设置（更新现场设备关联）
	ModifyMode *string `json:"ModifyMode,omitnil,omitempty" name:"ModifyMode"`
}

func (r *ModifyPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RemoteDeviceId")
	delete(f, "FieldDeviceIds")
	delete(f, "PolicyMode")
	delete(f, "ModifyMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPolicyResponseParams struct {
	// 添加关联失败的现场设备ID列表
	FailedInsertIds []*string `json:"FailedInsertIds,omitnil,omitempty" name:"FailedInsertIds"`

	// 解除关联失败的现场设备ID列表
	FailedDeleteIds []*string `json:"FailedDeleteIds,omitnil,omitempty" name:"FailedDeleteIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPolicyResponseParams `json:"Response"`
}

func (r *ModifyPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// 目标修改项目的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 修改后的项目名称，不填则不修改
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 修改后的项目描述，不填则不修改
	ProjectDescription *string `json:"ProjectDescription,omitnil,omitempty" name:"ProjectDescription"`

	// 修改后的权限模式，black为黑名单，white为白名单,不填则不修改
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目标修改项目的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 修改后的项目名称，不填则不修改
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 修改后的项目描述，不填则不修改
	ProjectDescription *string `json:"ProjectDescription,omitnil,omitempty" name:"ProjectDescription"`

	// 修改后的权限模式，black为黑名单，white为白名单,不填则不修改
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProjectName")
	delete(f, "ProjectDescription")
	delete(f, "PolicyMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectSecModeRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全模式  
	// 0：关闭项目共享密钥 
	// 1：开启项目共享密钥
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 项目密钥 32位 小写英文+数字；  项目密钥模式必填
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 自动注册方式
	// 0：关闭自动注册
	// 1：仅允许现场设备自动注册
	// 2：仅允许远端设备自动注册
	// 3：允许现场和远端设备均自动注册
	AutoRegister *int64 `json:"AutoRegister,omitnil,omitempty" name:"AutoRegister"`

	// 是否允许远端获取现场设备列表（getGwList）
	// 0：不允许
	// 1：允许
	FieldListEnable *int64 `json:"FieldListEnable,omitnil,omitempty" name:"FieldListEnable"`
}

type ModifyProjectSecModeRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全模式  
	// 0：关闭项目共享密钥 
	// 1：开启项目共享密钥
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 项目密钥 32位 小写英文+数字；  项目密钥模式必填
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 自动注册方式
	// 0：关闭自动注册
	// 1：仅允许现场设备自动注册
	// 2：仅允许远端设备自动注册
	// 3：允许现场和远端设备均自动注册
	AutoRegister *int64 `json:"AutoRegister,omitnil,omitempty" name:"AutoRegister"`

	// 是否允许远端获取现场设备列表（getGwList）
	// 0：不允许
	// 1：允许
	FieldListEnable *int64 `json:"FieldListEnable,omitnil,omitempty" name:"FieldListEnable"`
}

func (r *ModifyProjectSecModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectSecModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Mode")
	delete(f, "Key")
	delete(f, "AutoRegister")
	delete(f, "FieldListEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectSecModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectSecModeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectSecModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectSecModeResponseParams `json:"Response"`
}

func (r *ModifyProjectSecModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectSecModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiNet struct {
	// 网卡序号
	NetId *int64 `json:"NetId,omitnil,omitempty" name:"NetId"`

	// 网卡IP
	NetIp *string `json:"NetIp,omitnil,omitempty" name:"NetIp"`

	// 时延，单位ms
	Rtt []*int64 `json:"Rtt,omitnil,omitempty" name:"Rtt"`

	// 丢包率，单位%
	Lost []*int64 `json:"Lost,omitnil,omitempty" name:"Lost"`

	// 发送bps，单位kbps
	SendBps []*int64 `json:"SendBps,omitnil,omitempty" name:"SendBps"`

	// 接收bps，单位kbps
	RecvBps []*int64 `json:"RecvBps,omitnil,omitempty" name:"RecvBps"`
}

type PolicyInfo struct {
	// 远端设备ID
	RemoteDeviceId *string `json:"RemoteDeviceId,omitnil,omitempty" name:"RemoteDeviceId"`

	// 关联的现场设备ID
	FieldDeviceIds []*string `json:"FieldDeviceIds,omitnil,omitempty" name:"FieldDeviceIds"`

	// 最近添加时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type ProjectInfo struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDescription *string `json:"ProjectDescription,omitnil,omitempty" name:"ProjectDescription"`

	// 项目权限模式，black为黑名单，white为白名单
	PolicyMode *string `json:"PolicyMode,omitnil,omitempty" name:"PolicyMode"`

	// 项目信息修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type PublishParams struct {
	// 腾讯云直播推流地址url
	PublishUrl *string `json:"PublishUrl,omitnil,omitempty" name:"PublishUrl"`

	// 是否是腾讯云CDN，0为转推非腾讯云CDN，1为转推腾讯CDN，不携带该参数默认为1。
	IsTencentUrl *int64 `json:"IsTencentUrl,omitnil,omitempty" name:"IsTencentUrl"`
}

type RecentSessionInfo struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 远端设备ID
	RemoteDeviceId *string `json:"RemoteDeviceId,omitnil,omitempty" name:"RemoteDeviceId"`

	// 现场设备ID
	FieldDeviceId *string `json:"FieldDeviceId,omitnil,omitempty" name:"FieldDeviceId"`

	// 分辨率
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 会话开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 最后更新时间
	LatestUpdateTime *uint64 `json:"LatestUpdateTime,omitnil,omitempty" name:"LatestUpdateTime"`
}

type SessionDeviceDetail struct {
	// 设备类型：field或remote
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 起始点位时间，单位：秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束点位时间，单位：秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 码率，单位：kbps
	Rate []*int64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 帧率
	Fps []*int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 丢包率，单位：%
	Lost []*float64 `json:"Lost,omitnil,omitempty" name:"Lost"`

	// 网络时延，单位：ms
	NetworkLatency []*int64 `json:"NetworkLatency,omitnil,omitempty" name:"NetworkLatency"`

	// 视频时延，单位：ms
	VideoLatency []*int64 `json:"VideoLatency,omitnil,omitempty" name:"VideoLatency"`

	// CPU使用率，单位：%
	CpuUsed []*float64 `json:"CpuUsed,omitnil,omitempty" name:"CpuUsed"`

	// 内存使用率，单位：%
	MemUsed []*float64 `json:"MemUsed,omitnil,omitempty" name:"MemUsed"`

	// 时间偏移量，单位：秒
	TimeOffset []*uint64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// sdk版本
	Ver *string `json:"Ver,omitnil,omitempty" name:"Ver"`

	// 模式(p2p/server)
	SdkMode *string `json:"SdkMode,omitnil,omitempty" name:"SdkMode"`

	// 解码耗时，单位：ms
	DecodeCost []*int64 `json:"DecodeCost,omitnil,omitempty" name:"DecodeCost"`

	// 【已废弃，使用RenderCost】
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: RenderConst is deprecated.
	RenderConst []*int64 `json:"RenderConst,omitnil,omitempty" name:"RenderConst"`

	// 卡顿k100
	K100 []*float64 `json:"K100,omitnil,omitempty" name:"K100"`

	// 卡顿k150
	K150 []*float64 `json:"K150,omitnil,omitempty" name:"K150"`

	// nack请求数
	NACK []*int64 `json:"NACK,omitnil,omitempty" name:"NACK"`

	// 服务端调控码率,单位：kbps
	BitRateEstimate []*int64 `json:"BitRateEstimate,omitnil,omitempty" name:"BitRateEstimate"`

	// 宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 编码耗时，单位：ms
	EncodeCost []*int64 `json:"EncodeCost,omitnil,omitempty" name:"EncodeCost"`

	// 采集耗时，单位：ms
	CaptureCost []*int64 `json:"CaptureCost,omitnil,omitempty" name:"CaptureCost"`

	// 渲染耗时，单位：ms
	RenderCost []*int64 `json:"RenderCost,omitnil,omitempty" name:"RenderCost"`

	// 配置宽度
	ConfigWidth *int64 `json:"ConfigWidth,omitnil,omitempty" name:"ConfigWidth"`

	// 配置高度
	ConfigHeight *int64 `json:"ConfigHeight,omitnil,omitempty" name:"ConfigHeight"`

	// 平均帧间隔
	FrameDelta []*int64 `json:"FrameDelta,omitnil,omitempty" name:"FrameDelta"`

	// 最大帧间隔
	MaxFrameDelta []*int64 `json:"MaxFrameDelta,omitnil,omitempty" name:"MaxFrameDelta"`

	// 总码率评估,单位：kbps
	TotalBitrateEstimate []*int64 `json:"TotalBitrateEstimate,omitnil,omitempty" name:"TotalBitrateEstimate"`

	// 帧间隔大于100ms的卡顿时长
	Lag100Duration []*int64 `json:"Lag100Duration,omitnil,omitempty" name:"Lag100Duration"`

	// 帧间隔大于150ms的卡顿时长
	Lag150Duration []*int64 `json:"Lag150Duration,omitnil,omitempty" name:"Lag150Duration"`

	// 是否开启多网：0 单网，1 多网
	MultiMode *int64 `json:"MultiMode,omitnil,omitempty" name:"MultiMode"`

	// 多网卡信息
	MultiNet []*MultiNet `json:"MultiNet,omitnil,omitempty" name:"MultiNet"`
}

type SessionInfo struct {
	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 远端设备ID
	RemoteDeviceId *string `json:"RemoteDeviceId,omitnil,omitempty" name:"RemoteDeviceId"`

	// 现场设备ID
	FieldDeviceId *string `json:"FieldDeviceId,omitnil,omitempty" name:"FieldDeviceId"`

	// 分辨率
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 会话开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 会话结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通话质量：good|normal|bad，对应优良差
	Quality *string `json:"Quality,omitnil,omitempty" name:"Quality"`
}

type SessionIntervalStatistic struct {
	// 活跃现场设备数
	ActiveFieldDeviceNum *uint64 `json:"ActiveFieldDeviceNum,omitnil,omitempty" name:"ActiveFieldDeviceNum"`

	// 活跃远端设备数
	ActiveRemoteDeviceNum *uint64 `json:"ActiveRemoteDeviceNum,omitnil,omitempty" name:"ActiveRemoteDeviceNum"`

	// 会话数量
	SessionNum *uint64 `json:"SessionNum,omitnil,omitempty" name:"SessionNum"`

	// 会话时长，单位：分钟
	TotalDuration *uint64 `json:"TotalDuration,omitnil,omitempty" name:"TotalDuration"`

	// 时间戳，单位：秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间戳，单位：秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 优良会话占比，单位：%
	NotBadSessionRatio *uint64 `json:"NotBadSessionRatio,omitnil,omitempty" name:"NotBadSessionRatio"`
}

// Predefined struct for user
type StartPublishLiveStreamRequestParams struct {
	// 是否转码，0表示无需转码，1表示需要转码。是否收取转码费是由WithTranscoding参数决定的，WithTranscoding为0，表示旁路转推，不会收取转码费用，WithTranscoding为1，表示混流转推，会收取转码费用。 示例值：1
	WithTranscoding *int64 `json:"WithTranscoding,omitnil,omitempty" name:"WithTranscoding"`

	// 所有参与混流转推的主播持续离开TRTC房间或切换成观众超过MaxIdleTime的时长，自动停止转推，单位：秒。默认值为 30 秒，该值需大于等于 5秒，且小于等于 86400秒(24小时)。
	MaxIdleTime *int64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`

	// 转推视频参数
	VideoParams *VideoParams `json:"VideoParams,omitnil,omitempty" name:"VideoParams"`

	// 转推的URL参数，一个任务最多支持10个推流URL
	PublishParams []*PublishParams `json:"PublishParams,omitnil,omitempty" name:"PublishParams"`
}

type StartPublishLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 是否转码，0表示无需转码，1表示需要转码。是否收取转码费是由WithTranscoding参数决定的，WithTranscoding为0，表示旁路转推，不会收取转码费用，WithTranscoding为1，表示混流转推，会收取转码费用。 示例值：1
	WithTranscoding *int64 `json:"WithTranscoding,omitnil,omitempty" name:"WithTranscoding"`

	// 所有参与混流转推的主播持续离开TRTC房间或切换成观众超过MaxIdleTime的时长，自动停止转推，单位：秒。默认值为 30 秒，该值需大于等于 5秒，且小于等于 86400秒(24小时)。
	MaxIdleTime *int64 `json:"MaxIdleTime,omitnil,omitempty" name:"MaxIdleTime"`

	// 转推视频参数
	VideoParams *VideoParams `json:"VideoParams,omitnil,omitempty" name:"VideoParams"`

	// 转推的URL参数，一个任务最多支持10个推流URL
	PublishParams []*PublishParams `json:"PublishParams,omitnil,omitempty" name:"PublishParams"`
}

func (r *StartPublishLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WithTranscoding")
	delete(f, "MaxIdleTime")
	delete(f, "VideoParams")
	delete(f, "PublishParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishLiveStreamResponseParams struct {
	// 用于唯一标识转推任务，由腾讯云服务端生成，后续停止请求需要携带TaskiD参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartPublishLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishLiveStreamResponseParams `json:"Response"`
}

func (r *StartPublishLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishLiveStreamRequestParams struct {
	// 唯一标识转推任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopPublishLiveStreamRequest struct {
	*tchttp.BaseRequest
	
	// 唯一标识转推任务。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopPublishLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishLiveStreamResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopPublishLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishLiveStreamResponseParams `json:"Response"`
}

func (r *StopPublishLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoList struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 流id
	VideoStreamId *int64 `json:"VideoStreamId,omitnil,omitempty" name:"VideoStreamId"`

	// 子画面在输出时的宽度，单位为像素值，不填默认为0。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 子画面在输出时的高度，单位为像素值，不填默认为0。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type VideoParams struct {
	// 输出流宽，音视频输出时必填。取值范围[0,1920]，单位为像素值。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 输出流高，音视频输出时必填。取值范围[0,1080]，单位为像素值。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 输出流帧率，音视频输出时必填。取值范围[1,60]，表示混流的输出帧率可选范围为1到60fps。
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// 输出流码率，音视频输出时必填。取值范围[1,10000]，单位为kbps。
	BitRate *int64 `json:"BitRate,omitnil,omitempty" name:"BitRate"`

	// 输出流gop，音视频输出时必填。取值范围[1,5]，单位为秒。
	Gop *int64 `json:"Gop,omitnil,omitempty" name:"Gop"`

	// 转推视频流列表
	VideoList []*VideoList `json:"VideoList,omitnil,omitempty" name:"VideoList"`
}