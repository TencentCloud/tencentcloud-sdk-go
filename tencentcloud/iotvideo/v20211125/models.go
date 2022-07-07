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

package v20211125

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CallDeviceActionAsyncRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

type CallDeviceActionAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ActionId")
	delete(f, "InputParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallDeviceActionAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionAsyncResponseParams struct {
	// 调用Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 异步调用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CallDeviceActionAsyncResponse struct {
	*tchttp.BaseResponse
	Response *CallDeviceActionAsyncResponseParams `json:"Response"`
}

func (r *CallDeviceActionAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionSyncRequestParams struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

type CallDeviceActionSyncRequest struct {
	*tchttp.BaseRequest
	
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	ActionId *string `json:"ActionId,omitempty" name:"ActionId"`

	// 输入参数
	InputParams *string `json:"InputParams,omitempty" name:"InputParams"`
}

func (r *CallDeviceActionSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DeviceName")
	delete(f, "ActionId")
	delete(f, "InputParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallDeviceActionSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallDeviceActionSyncResponseParams struct {
	// 调用Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 输出参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputParams *string `json:"OutputParams,omitempty" name:"OutputParams"`

	// 返回状态，当设备不在线等部分情况，会通过该 Status 返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CallDeviceActionSyncResponse struct {
	*tchttp.BaseResponse
	Response *CallDeviceActionSyncResponseParams `json:"Response"`
}

func (r *CallDeviceActionSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallDeviceActionSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataStatsRequestParams struct {
	// 开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeDeviceDataStatsRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeDeviceDataStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceDataStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceDataStatsResponseParams struct {
	// 累计注册设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterCnt *uint64 `json:"RegisterCnt,omitempty" name:"RegisterCnt"`

	// 设备数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DeviceCntStats `json:"Data,omitempty" name:"Data"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeviceDataStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceDataStatsResponseParams `json:"Response"`
}

func (r *DescribeDeviceDataStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceDataStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageDataStatsRequestParams struct {
	// 统计开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 统计结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type DescribeMessageDataStatsRequest struct {
	*tchttp.BaseRequest
	
	// 统计开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 统计结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 产品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeMessageDataStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageDataStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "ProductId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageDataStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageDataStatsResponseParams struct {
	// 消息数量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*MessageCntStats `json:"Data,omitempty" name:"Data"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMessageDataStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageDataStatsResponseParams `json:"Response"`
}

func (r *DescribeMessageDataStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageDataStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceCntStats struct {
	// 统计日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 新增注册设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewRegisterCnt *uint64 `json:"NewRegisterCnt,omitempty" name:"NewRegisterCnt"`

	// 新增激活设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewActivateCnt *uint64 `json:"NewActivateCnt,omitempty" name:"NewActivateCnt"`

	// 活跃设备数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveCnt *uint64 `json:"ActiveCnt,omitempty" name:"ActiveCnt"`
}

type MessageCntStats struct {
	// 统计日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 物模型上行消息数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpMsgCnt *uint64 `json:"UpMsgCnt,omitempty" name:"UpMsgCnt"`

	// 物模型下行消息数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownMsgCnt *uint64 `json:"DownMsgCnt,omitempty" name:"DownMsgCnt"`

	// ntp消息数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NtpMsgCnt *uint64 `json:"NtpMsgCnt,omitempty" name:"NtpMsgCnt"`
}