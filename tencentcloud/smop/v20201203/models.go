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

package v20201203

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type SubmitTaskEventRequestParams struct {
	// 用户ID
	AccountId *string `json:"AccountId,omitnil" name:"AccountId"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 订单ID
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 任务事件Code
	Code *string `json:"Code,omitnil" name:"Code"`

	// 同步异步方式：0为同步、1位异步
	Async *int64 `json:"Async,omitnil" name:"Async"`

	// 产品ID
	ProductId *int64 `json:"ProductId,omitnil" name:"ProductId"`

	// 回调地址
	NotifyURL *string `json:"NotifyURL,omitnil" name:"NotifyURL"`
}

type SubmitTaskEventRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	AccountId *string `json:"AccountId,omitnil" name:"AccountId"`

	// 设备ID
	DeviceId *string `json:"DeviceId,omitnil" name:"DeviceId"`

	// 订单ID
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 任务事件Code
	Code *string `json:"Code,omitnil" name:"Code"`

	// 同步异步方式：0为同步、1位异步
	Async *int64 `json:"Async,omitnil" name:"Async"`

	// 产品ID
	ProductId *int64 `json:"ProductId,omitnil" name:"ProductId"`

	// 回调地址
	NotifyURL *string `json:"NotifyURL,omitnil" name:"NotifyURL"`
}

func (r *SubmitTaskEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountId")
	delete(f, "DeviceId")
	delete(f, "OrderId")
	delete(f, "Code")
	delete(f, "Async")
	delete(f, "ProductId")
	delete(f, "NotifyURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTaskEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskEventResponseParams struct {
	// 订单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 信息码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil" name:"Code"`

	// success
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 任务处理结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskEventData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SubmitTaskEventResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTaskEventResponseParams `json:"Response"`
}

func (r *SubmitTaskEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskEventData struct {
	// 状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil" name:"Code"`

	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// 当前完成或正在完成的任务订单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskOrderId *string `json:"TaskOrderId,omitnil" name:"TaskOrderId"`

	// 当前任务订单状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCode *int64 `json:"TaskCode,omitnil" name:"TaskCode"`

	// 获得积分数/成长值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCoinNumber *int64 `json:"TaskCoinNumber,omitnil" name:"TaskCoinNumber"`

	// 任务类型后台代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 当前积分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCoin *int64 `json:"TotalCoin,omitnil" name:"TotalCoin"`

	// 用户透传的代码块
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attach *string `json:"Attach,omitnil" name:"Attach"`

	// 计次任务当前完成次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoneTimes *int64 `json:"DoneTimes,omitnil" name:"DoneTimes"`

	// 计次任务当前所需完成次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTimes *int64 `json:"TotalTimes,omitnil" name:"TotalTimes"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 当前成长值
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrowScore *int64 `json:"GrowScore,omitnil" name:"GrowScore"`
}