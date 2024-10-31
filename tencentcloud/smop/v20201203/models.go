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
	// 用户唯一标识，最大长度为64
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 用户设备ID，最大长度为64
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 任务的唯一订单号，只能是数字、大小写字母，且在同一个产品ID下唯一，最大长度为64
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 任务事件Code，在腾讯安心用户运营平台下的任务事件列表中设置并获取
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 任务结果是否异步通知。0表示任务结果在返回信息中同步返回；1表示任务结果通过回调结果异步通知。
	Async *int64 `json:"Async,omitnil,omitempty" name:"Async"`

	// 产品ID，可在腾讯安心用户运营平台的企业管理中获取
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 异步接收任务结果通知的回调地址。在Async为1的时候，会将任务结果通过该回调地址进行通知。
	NotifyURL *string `json:"NotifyURL,omitnil,omitempty" name:"NotifyURL"`
}

type SubmitTaskEventRequest struct {
	*tchttp.BaseRequest
	
	// 用户唯一标识，最大长度为64
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 用户设备ID，最大长度为64
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 任务的唯一订单号，只能是数字、大小写字母，且在同一个产品ID下唯一，最大长度为64
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 任务事件Code，在腾讯安心用户运营平台下的任务事件列表中设置并获取
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 任务结果是否异步通知。0表示任务结果在返回信息中同步返回；1表示任务结果通过回调结果异步通知。
	Async *int64 `json:"Async,omitnil,omitempty" name:"Async"`

	// 产品ID，可在腾讯安心用户运营平台的企业管理中获取
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 异步接收任务结果通知的回调地址。在Async为1的时候，会将任务结果通过该回调地址进行通知。
	NotifyURL *string `json:"NotifyURL,omitnil,omitempty" name:"NotifyURL"`
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
	// 任务的唯一订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 信息码。0表示成功，-1标识失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务处理结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskEventData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 状态码，0为成功，-1为失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 当前完成或正在完成的安心用户运营平台的任务订单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskOrderId *string `json:"TaskOrderId,omitnil,omitempty" name:"TaskOrderId"`

	// 当前任务订单状态码。1代表未完成；2代表已完成但未提交任务；3表示已完成，且已提交获得积分任务；4表示过期任务，提交后不获得积分。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCode *int64 `json:"TaskCode,omitnil,omitempty" name:"TaskCode"`

	// 获得积分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCoinNumber *int64 `json:"TaskCoinNumber,omitnil,omitempty" name:"TaskCoinNumber"`

	// 任务类型后台代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 用户的当前积分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCoin *int64 `json:"TotalCoin,omitnil,omitempty" name:"TotalCoin"`

	// 用户透传的附加数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attach *string `json:"Attach,omitnil,omitempty" name:"Attach"`

	// 计次任务当前完成次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoneTimes *int64 `json:"DoneTimes,omitnil,omitempty" name:"DoneTimes"`

	// 计次任务当前所需完成次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTimes *int64 `json:"TotalTimes,omitnil,omitempty" name:"TotalTimes"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 用户当前成长值
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrowScore *int64 `json:"GrowScore,omitnil,omitempty" name:"GrowScore"`
}