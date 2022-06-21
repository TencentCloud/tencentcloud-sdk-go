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

package v20191010

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type FlowProductRemindRequestParams struct {
	// 服务商uin
	ProviderUin *string `json:"ProviderUin,omitempty" name:"ProviderUin"`

	// 服务商实例ID
	SignId *string `json:"SignId,omitempty" name:"SignId"`

	// 云市场实例ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 实例总流量
	TotalFlow *string `json:"TotalFlow,omitempty" name:"TotalFlow"`

	// 剩余流量
	LeftFlow *string `json:"LeftFlow,omitempty" name:"LeftFlow"`

	// 流量单位
	FlowUnit *string `json:"FlowUnit,omitempty" name:"FlowUnit"`
}

type FlowProductRemindRequest struct {
	*tchttp.BaseRequest
	
	// 服务商uin
	ProviderUin *string `json:"ProviderUin,omitempty" name:"ProviderUin"`

	// 服务商实例ID
	SignId *string `json:"SignId,omitempty" name:"SignId"`

	// 云市场实例ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 实例总流量
	TotalFlow *string `json:"TotalFlow,omitempty" name:"TotalFlow"`

	// 剩余流量
	LeftFlow *string `json:"LeftFlow,omitempty" name:"LeftFlow"`

	// 流量单位
	FlowUnit *string `json:"FlowUnit,omitempty" name:"FlowUnit"`
}

func (r *FlowProductRemindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlowProductRemindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProviderUin")
	delete(f, "SignId")
	delete(f, "ResourceId")
	delete(f, "TotalFlow")
	delete(f, "LeftFlow")
	delete(f, "FlowUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlowProductRemindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlowProductRemindResponseParams struct {
	// 是否成功
	Success *string `json:"Success,omitempty" name:"Success"`

	// 流水号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FlowProductRemindResponse struct {
	*tchttp.BaseResponse
	Response *FlowProductRemindResponseParams `json:"Response"`
}

func (r *FlowProductRemindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlowProductRemindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUsagePlanUsageAmountRequestParams struct {
	// 用于查询实例的Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type GetUsagePlanUsageAmountRequest struct {
	*tchttp.BaseRequest
	
	// 用于查询实例的Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetUsagePlanUsageAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUsagePlanUsageAmountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUsagePlanUsageAmountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUsagePlanUsageAmountResponseParams struct {
	// 最大调用量
	MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

	// 已经调用量
	InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

	// 剩余调用量
	RemainingRequestNum *int64 `json:"RemainingRequestNum,omitempty" name:"RemainingRequestNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetUsagePlanUsageAmountResponse struct {
	*tchttp.BaseResponse
	Response *GetUsagePlanUsageAmountResponseParams `json:"Response"`
}

func (r *GetUsagePlanUsageAmountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUsagePlanUsageAmountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}