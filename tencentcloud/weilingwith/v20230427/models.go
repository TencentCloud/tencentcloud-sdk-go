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

package v20230427

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddAlarmProcessRecordRequestParams struct {

}

type AddAlarmProcessRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *AddAlarmProcessRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAlarmProcessRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAlarmProcessRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAlarmProcessRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddAlarmProcessRecordResponse struct {
	*tchttp.BaseResponse
	Response *AddAlarmProcessRecordResponseParams `json:"Response"`
}

func (r *AddAlarmProcessRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAlarmProcessRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateDeviceRequestParams struct {

}

type BatchCreateDeviceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *BatchCreateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchCreateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateDeviceResponseParams `json:"Response"`
}

func (r *BatchCreateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchKillAlarmRequestParams struct {

}

type BatchKillAlarmRequest struct {
	*tchttp.BaseRequest
	
}

func (r *BatchKillAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchKillAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchKillAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchKillAlarmResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchKillAlarmResponse struct {
	*tchttp.BaseResponse
	Response *BatchKillAlarmResponseParams `json:"Response"`
}

func (r *BatchKillAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchKillAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchReportAppMessageRequestParams struct {

}

type BatchReportAppMessageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *BatchReportAppMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchReportAppMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchReportAppMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchReportAppMessageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BatchReportAppMessageResponse struct {
	*tchttp.BaseResponse
	Response *BatchReportAppMessageResponseParams `json:"Response"`
}

func (r *BatchReportAppMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchReportAppMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeAlarmStatusRequestParams struct {

}

type ChangeAlarmStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ChangeAlarmStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAlarmStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeAlarmStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeAlarmStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ChangeAlarmStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChangeAlarmStatusResponseParams `json:"Response"`
}

func (r *ChangeAlarmStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAlarmStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlCameraPTZRequestParams struct {

}

type ControlCameraPTZRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ControlCameraPTZRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlCameraPTZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlCameraPTZRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlCameraPTZResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ControlCameraPTZResponse struct {
	*tchttp.BaseResponse
	Response *ControlCameraPTZResponseParams `json:"Response"`
}

func (r *ControlCameraPTZResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlCameraPTZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceRequestParams struct {

}

type ControlDeviceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ControlDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ControlDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ControlDeviceResponseParams `json:"Response"`
}

func (r *ControlDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationTokenRequestParams struct {

}

type CreateApplicationTokenRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateApplicationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationTokenResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateApplicationTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationTokenResponseParams `json:"Response"`
}

func (r *CreateApplicationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionListRequestParams struct {

}

type DescribeActionListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeActionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeActionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeActionListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeActionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeActionListResponseParams `json:"Response"`
}

func (r *DescribeActionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeActionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdministrationByTagRequestParams struct {

}

type DescribeAdministrationByTagRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAdministrationByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdministrationByTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdministrationByTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdministrationByTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAdministrationByTagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdministrationByTagResponseParams `json:"Response"`
}

func (r *DescribeAdministrationByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdministrationByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmLevelListRequestParams struct {

}

type DescribeAlarmLevelListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAlarmLevelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmLevelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmLevelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmLevelListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAlarmLevelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmLevelListResponseParams `json:"Response"`
}

func (r *DescribeAlarmLevelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmLevelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmListRequestParams struct {

}

type DescribeAlarmListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmListResponseParams `json:"Response"`
}

func (r *DescribeAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmStatusListRequestParams struct {

}

type DescribeAlarmStatusListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAlarmStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmStatusListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAlarmStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmStatusListResponseParams `json:"Response"`
}

func (r *DescribeAlarmStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmTypeListRequestParams struct {

}

type DescribeAlarmTypeListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAlarmTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmTypeListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAlarmTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmTypeListResponseParams `json:"Response"`
}

func (r *DescribeAlarmTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListRequestParams struct {

}

type DescribeApplicationListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeApplicationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeApplicationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationListResponseParams `json:"Response"`
}

func (r *DescribeApplicationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingListRequestParams struct {

}

type DescribeBuildingListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBuildingListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuildingListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBuildingListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuildingListResponseParams `json:"Response"`
}

func (r *DescribeBuildingListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingModelRequestParams struct {

}

type DescribeBuildingModelRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBuildingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuildingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingModelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBuildingModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuildingModelResponseParams `json:"Response"`
}

func (r *DescribeBuildingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingProfileRequestParams struct {

}

type DescribeBuildingProfileRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBuildingProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBuildingProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBuildingProfileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBuildingProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBuildingProfileResponseParams `json:"Response"`
}

func (r *DescribeBuildingProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBuildingProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCameraExtendInfoRequestParams struct {

}

type DescribeCameraExtendInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCameraExtendInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCameraExtendInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCameraExtendInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCameraExtendInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCameraExtendInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCameraExtendInfoResponseParams `json:"Response"`
}

func (r *DescribeCameraExtendInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCameraExtendInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCityWorkspaceListRequestParams struct {

}

type DescribeCityWorkspaceListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCityWorkspaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCityWorkspaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCityWorkspaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCityWorkspaceListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCityWorkspaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCityWorkspaceListResponseParams `json:"Response"`
}

func (r *DescribeCityWorkspaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCityWorkspaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListRequestParams struct {

}

type DescribeDeviceListRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeDeviceShadowListRequestParams struct {

}

type DescribeDeviceShadowListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeviceShadowListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceShadowListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceShadowListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceShadowListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceShadowListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceShadowListResponseParams `json:"Response"`
}

func (r *DescribeDeviceShadowListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceShadowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusListRequestParams struct {

}

type DescribeDeviceStatusListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeviceStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceStatusListResponseParams `json:"Response"`
}

func (r *DescribeDeviceStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusStatRequestParams struct {

}

type DescribeDeviceStatusStatRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeviceStatusStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceStatusStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceStatusStatResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceStatusStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceStatusStatResponseParams `json:"Response"`
}

func (r *DescribeDeviceStatusStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceStatusStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTagListRequestParams struct {

}

type DescribeDeviceTagListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeviceTagListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTagListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceTagListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTagListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceTagListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceTagListResponseParams `json:"Response"`
}

func (r *DescribeDeviceTagListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTagListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTypeListRequestParams struct {

}

type DescribeDeviceTypeListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDeviceTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceTypeListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceTypeListResponseParams `json:"Response"`
}

func (r *DescribeDeviceTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeApplicationTokenRequestParams struct {

}

type DescribeEdgeApplicationTokenRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEdgeApplicationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeApplicationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeApplicationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeApplicationTokenResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgeApplicationTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeApplicationTokenResponseParams `json:"Response"`
}

func (r *DescribeEdgeApplicationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeApplicationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfilePageRequestParams struct {

}

type DescribeElementProfilePageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeElementProfilePageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfilePageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeElementProfilePageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfilePageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeElementProfilePageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeElementProfilePageResponseParams `json:"Response"`
}

func (r *DescribeElementProfilePageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfilePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfileTreeRequestParams struct {

}

type DescribeElementProfileTreeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeElementProfileTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfileTreeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeElementProfileTreeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeElementProfileTreeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeElementProfileTreeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeElementProfileTreeResponseParams `json:"Response"`
}

func (r *DescribeElementProfileTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeElementProfileTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventListRequestParams struct {

}

type DescribeEventListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventListResponseParams `json:"Response"`
}

func (r *DescribeEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadURLRequestParams struct {

}

type DescribeFileDownloadURLRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFileDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadURLResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFileDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileDownloadURLResponseParams `json:"Response"`
}

func (r *DescribeFileDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUploadURLRequestParams struct {

}

type DescribeFileUploadURLRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFileUploadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUploadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileUploadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUploadURLResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFileUploadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileUploadURLResponseParams `json:"Response"`
}

func (r *DescribeFileUploadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUploadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInterfaceListRequestParams struct {

}

type DescribeInterfaceListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInterfaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInterfaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInterfaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInterfaceListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInterfaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInterfaceListResponseParams `json:"Response"`
}

func (r *DescribeInterfaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInterfaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLinkRuleListRequestParams struct {

}

type DescribeLinkRuleListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLinkRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinkRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLinkRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLinkRuleListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLinkRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLinkRuleListResponseParams `json:"Response"`
}

func (r *DescribeLinkRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLinkRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListRequestParams struct {

}

type DescribeModelListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelListResponseParams `json:"Response"`
}

func (r *DescribeModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductListRequestParams struct {

}

type DescribeProductListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProductListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProductListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductListResponseParams `json:"Response"`
}

func (r *DescribeProductListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePropertyListRequestParams struct {

}

type DescribePropertyListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePropertyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePropertyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePropertyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePropertyListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePropertyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePropertyListResponseParams `json:"Response"`
}

func (r *DescribePropertyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePropertyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleDetailRequestParams struct {

}

type DescribeRuleDetailRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleDetailResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSceneListRequestParams struct {

}

type DescribeSceneListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSceneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSceneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSceneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSceneListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSceneListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSceneListResponseParams `json:"Response"`
}

func (r *DescribeSceneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSceneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceIdListRequestParams struct {

}

type DescribeSpaceDeviceIdListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSpaceDeviceIdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceIdListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceDeviceIdListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceIdListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpaceDeviceIdListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceDeviceIdListResponseParams `json:"Response"`
}

func (r *DescribeSpaceDeviceIdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceIdListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceRelationListRequestParams struct {

}

type DescribeSpaceDeviceRelationListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSpaceDeviceRelationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceRelationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceDeviceRelationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceDeviceRelationListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpaceDeviceRelationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceDeviceRelationListResponseParams `json:"Response"`
}

func (r *DescribeSpaceDeviceRelationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceDeviceRelationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceInfoByDeviceIdRequestParams struct {

}

type DescribeSpaceInfoByDeviceIdRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSpaceInfoByDeviceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceInfoByDeviceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceInfoByDeviceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceInfoByDeviceIdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpaceInfoByDeviceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceInfoByDeviceIdResponseParams `json:"Response"`
}

func (r *DescribeSpaceInfoByDeviceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceInfoByDeviceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceRelationByDeviceIdRequestParams struct {

}

type DescribeSpaceRelationByDeviceIdRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSpaceRelationByDeviceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceRelationByDeviceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceRelationByDeviceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceRelationByDeviceIdResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpaceRelationByDeviceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceRelationByDeviceIdResponseParams `json:"Response"`
}

func (r *DescribeSpaceRelationByDeviceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceRelationByDeviceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceTypeListRequestParams struct {

}

type DescribeSpaceTypeListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSpaceTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceTypeListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpaceTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceTypeListResponseParams `json:"Response"`
}

func (r *DescribeSpaceTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantBuildingCountAndAreaRequestParams struct {

}

type DescribeTenantBuildingCountAndAreaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTenantBuildingCountAndAreaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantBuildingCountAndAreaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTenantBuildingCountAndAreaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantBuildingCountAndAreaResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTenantBuildingCountAndAreaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTenantBuildingCountAndAreaResponseParams `json:"Response"`
}

func (r *DescribeTenantBuildingCountAndAreaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantBuildingCountAndAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantDepartmentListRequestParams struct {
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil" name:"Cursor"`
}

type DescribeTenantDepartmentListRequest struct {
	*tchttp.BaseRequest
	
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil" name:"Cursor"`
}

func (r *DescribeTenantDepartmentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantDepartmentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationToken")
	delete(f, "TenantId")
	delete(f, "UpdateAt")
	delete(f, "DepartmentId")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTenantDepartmentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantDepartmentListResponseParams struct {
	// 返回数据
	Result *SsoDepartmentsResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTenantDepartmentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTenantDepartmentListResponseParams `json:"Response"`
}

func (r *DescribeTenantDepartmentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantDepartmentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantUserListRequestParams struct {
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil" name:"Cursor"`

	// 状态，0，获取所有数据，1正常启用，2禁用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 是否递归获取子级数据，0需要，1不需要，默认为0
	NoRecursive *string `json:"NoRecursive,omitnil" name:"NoRecursive"`
}

type DescribeTenantUserListRequest struct {
	*tchttp.BaseRequest
	
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`

	// 部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 用户id
	Cursor *string `json:"Cursor,omitnil" name:"Cursor"`

	// 状态，0，获取所有数据，1正常启用，2禁用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 项目空间id
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 是否递归获取子级数据，0需要，1不需要，默认为0
	NoRecursive *string `json:"NoRecursive,omitnil" name:"NoRecursive"`
}

func (r *DescribeTenantUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ApplicationToken")
	delete(f, "TenantId")
	delete(f, "UpdateAt")
	delete(f, "DepartmentId")
	delete(f, "Cursor")
	delete(f, "Status")
	delete(f, "WorkspaceId")
	delete(f, "Keyword")
	delete(f, "NoRecursive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTenantUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantUserListResponseParams struct {
	// 返回数据
	Result *SsoUserResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTenantUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTenantUserListResponseParams `json:"Response"`
}

func (r *DescribeTenantUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoCloudRecordRequestParams struct {

}

type DescribeVideoCloudRecordRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVideoCloudRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoCloudRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoCloudRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoCloudRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVideoCloudRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoCloudRecordResponseParams `json:"Response"`
}

func (r *DescribeVideoCloudRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoCloudRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoLiveStreamRequestParams struct {

}

type DescribeVideoLiveStreamRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVideoLiveStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoLiveStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoLiveStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoLiveStreamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVideoLiveStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoLiveStreamResponseParams `json:"Response"`
}

func (r *DescribeVideoLiveStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoLiveStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoRecordStreamRequestParams struct {

}

type DescribeVideoRecordStreamRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVideoRecordStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoRecordStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoRecordStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoRecordStreamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVideoRecordStreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoRecordStreamResponseParams `json:"Response"`
}

func (r *DescribeVideoRecordStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoRecordStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkSpaceBuildingCountAndAreaRequestParams struct {

}

type DescribeWorkSpaceBuildingCountAndAreaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWorkSpaceBuildingCountAndAreaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkSpaceBuildingCountAndAreaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkSpaceBuildingCountAndAreaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkSpaceBuildingCountAndAreaResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkSpaceBuildingCountAndAreaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkSpaceBuildingCountAndAreaResponseParams `json:"Response"`
}

func (r *DescribeWorkSpaceBuildingCountAndAreaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkSpaceBuildingCountAndAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceListRequestParams struct {

}

type DescribeWorkspaceListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWorkspaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceUserListRequestParams struct {
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

type DescribeWorkspaceUserListRequest struct {
	*tchttp.BaseRequest
	
	// 翻页页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitnil" name:"WorkspaceId"`

	// token
	ApplicationToken *string `json:"ApplicationToken,omitnil" name:"ApplicationToken"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 更新时间戳，单位秒
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`
}

func (r *DescribeWorkspaceUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "WorkspaceId")
	delete(f, "ApplicationToken")
	delete(f, "TenantId")
	delete(f, "UpdateAt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkspaceUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkspaceUserListResponseParams struct {
	// 返回数据
	Result *SsoTeamUserResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkspaceUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkspaceUserListResponseParams `json:"Response"`
}

func (r *DescribeWorkspaceUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkspaceUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceNameRequestParams struct {

}

type ModifyDeviceNameRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyDeviceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDeviceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceNameResponseParams `json:"Response"`
}

func (r *ModifyDeviceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportAppMessageRequestParams struct {

}

type ReportAppMessageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ReportAppMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAppMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportAppMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportAppMessageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReportAppMessageResponse struct {
	*tchttp.BaseResponse
	Response *ReportAppMessageResponseParams `json:"Response"`
}

func (r *ReportAppMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportAppMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SsoDepartment struct {
	// 部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 父级部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentDepartmentId *string `json:"ParentDepartmentId,omitnil" name:"ParentDepartmentId"`
}

type SsoDepartmentsResult struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 部门列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Departments []*SsoDepartment `json:"Departments,omitnil" name:"Departments"`
}

type SsoTeamUser struct {
	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 用户类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 所属租户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 用户状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *int64 `json:"CreateAt,omitnil" name:"CreateAt"`

	// 部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitnil" name:"DepartmentName"`

	// 是否关联权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkFilter *int64 `json:"LinkFilter,omitnil" name:"LinkFilter"`
}

type SsoTeamUserResult struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 部门用户列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*SsoTeamUser `json:"Users,omitnil" name:"Users"`
}

type SsoUser struct {
	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 用户类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 所属租户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// 所属组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroup *string `json:"UserGroup,omitnil" name:"UserGroup"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 用户状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *int64 `json:"CreateAt,omitnil" name:"CreateAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateAt *int64 `json:"UpdateAt,omitnil" name:"UpdateAt"`

	// 是否属于团队
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongTeam *int64 `json:"BelongTeam,omitnil" name:"BelongTeam"`

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitnil" name:"DepartmentName"`

	// 子账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentUserId *int64 `json:"DepartmentUserId,omitnil" name:"DepartmentUserId"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`
}

type SsoUserResult struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 租户人员数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*SsoUser `json:"Users,omitnil" name:"Users"`
}

// Predefined struct for user
type StopVideoStreamingRequestParams struct {

}

type StopVideoStreamingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *StopVideoStreamingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVideoStreamingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopVideoStreamingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopVideoStreamingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopVideoStreamingResponse struct {
	*tchttp.BaseResponse
	Response *StopVideoStreamingResponseParams `json:"Response"`
}

func (r *StopVideoStreamingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopVideoStreamingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceParkAttributesRequestParams struct {

}

type UpdateWorkspaceParkAttributesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *UpdateWorkspaceParkAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceParkAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkspaceParkAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkspaceParkAttributesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateWorkspaceParkAttributesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkspaceParkAttributesResponseParams `json:"Response"`
}

func (r *UpdateWorkspaceParkAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkspaceParkAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}