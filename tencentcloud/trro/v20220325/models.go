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
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备类型筛选，不填默认为全部设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 对设备ID或Name按关键字进行模糊匹配，不填则不进行模糊匹配
	SearchWords *string `json:"SearchWords,omitnil,omitempty" name:"SearchWords"`

	// 每页返回的最大设备数，不填默认为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页码，不填默认为1（首页）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// 设备所属项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 设备类型筛选，不填默认为全部设备类型
	DeviceType *string `json:"DeviceType,omitnil,omitempty" name:"DeviceType"`

	// 对设备ID或Name按关键字进行模糊匹配，不填则不进行模糊匹配
	SearchWords *string `json:"SearchWords,omitnil,omitempty" name:"SearchWords"`

	// 每页返回的最大设备数，不填默认为10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页码，不填默认为1（首页）
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListResponseParams struct {
	// 设备信息列表
	Devices []*DeviceInfo `json:"Devices,omitnil,omitempty" name:"Devices"`

	// 设备总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 本次返回的设备数
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 设备名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 已经绑定license数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseCount *int64 `json:"LicenseCount,omitnil,omitempty" name:"LicenseCount"`

	// 剩余天数：天
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainDay *int64 `json:"RemainDay,omitnil,omitempty" name:"RemainDay"`

	// 过期时间：s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 服务时长：s
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 已经绑定licenseId列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseIds []*string `json:"LicenseIds,omitnil,omitempty" name:"LicenseIds"`

	// 每月license的限定时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonthlyRemainTime *int64 `json:"MonthlyRemainTime,omitnil,omitempty" name:"MonthlyRemainTime"`
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

type License struct {
	// 该类型的license个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// license状态：0:未绑定；1:已绑定；2:已停服；3:已退费
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 到期时间戳：s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 服务时长：s
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 剩余天数：天
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainDay *int64 `json:"RemainDay,omitnil,omitempty" name:"RemainDay"`

	// 该类型的licenseId列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseIds []*string `json:"LicenseIds,omitnil,omitempty" name:"LicenseIds"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInsertIds []*string `json:"FailedInsertIds,omitnil,omitempty" name:"FailedInsertIds"`

	// 解除关联失败的现场设备ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ver *string `json:"Ver,omitnil,omitempty" name:"Ver"`

	// 模式(p2p/server)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkMode *string `json:"SdkMode,omitnil,omitempty" name:"SdkMode"`

	// 解码耗时，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	DecodeCost []*int64 `json:"DecodeCost,omitnil,omitempty" name:"DecodeCost"`

	// 渲染耗时，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenderConst []*int64 `json:"RenderConst,omitnil,omitempty" name:"RenderConst"`

	// 卡顿k100
	// 注意：此字段可能返回 null，表示取不到有效值。
	K100 []*float64 `json:"K100,omitnil,omitempty" name:"K100"`

	// 卡顿k150
	// 注意：此字段可能返回 null，表示取不到有效值。
	K150 []*float64 `json:"K150,omitnil,omitempty" name:"K150"`

	// nack请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NACK []*int64 `json:"NACK,omitnil,omitempty" name:"NACK"`

	// 服务端调控码率,单位：kbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	BitRateEstimate []*int64 `json:"BitRateEstimate,omitnil,omitempty" name:"BitRateEstimate"`

	// 宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`

	// 编码耗时，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncodeCost []*int64 `json:"EncodeCost,omitnil,omitempty" name:"EncodeCost"`

	// 采集耗时，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptureCost []*int64 `json:"CaptureCost,omitnil,omitempty" name:"CaptureCost"`
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