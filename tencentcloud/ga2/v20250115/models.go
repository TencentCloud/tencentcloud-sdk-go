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

package v20250115

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateGlobalAcceleratorRequestParams struct {
	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计费模式，PREPAID：表示预付费，即包年包月，POSTPAID：表示后付费，即按量计费。默认：POSTPAID。当前仅支持后付费。</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>跨境类型；HighQuality：精品BGP-IP跨境；Unicom：联通专线跨境。</p>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>此Flag代表签署跨境服务承诺书。当使用跨境服务时候，此字段必传。True：代表签署。</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`

	// <p>标签信息</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateGlobalAcceleratorRequest struct {
	*tchttp.BaseRequest
	
	// <p>名称，最大长度不能超过60个字节。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>计费模式，PREPAID：表示预付费，即包年包月，POSTPAID：表示后付费，即按量计费。默认：POSTPAID。当前仅支持后付费。</p>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// <p>描述信息，最大长度不能超过100个字节。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>跨境类型；HighQuality：精品BGP-IP跨境；Unicom：联通专线跨境。</p>
	CrossBorderType *string `json:"CrossBorderType,omitnil,omitempty" name:"CrossBorderType"`

	// <p>此Flag代表签署跨境服务承诺书。当使用跨境服务时候，此字段必传。True：代表签署。</p>
	CrossBorderPromiseFlag *bool `json:"CrossBorderPromiseFlag,omitnil,omitempty" name:"CrossBorderPromiseFlag"`

	// <p>标签信息</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateGlobalAcceleratorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "InstanceChargeType")
	delete(f, "Description")
	delete(f, "CrossBorderType")
	delete(f, "CrossBorderPromiseFlag")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalAcceleratorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalAcceleratorResponseParams struct {
	// <p>任务ID。</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>全球加速实例ID。</p>
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalAcceleratorResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalAcceleratorResponseParams `json:"Response"`
}

func (r *CreateGlobalAcceleratorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalAcceleratorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderSettlementRequestParams struct {
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域。
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// 终端节点组地域。
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// 账单年月时间。
	SettlementMonth *uint64 `json:"SettlementMonth,omitnil,omitempty" name:"SettlementMonth"`
}

type DescribeCrossBorderSettlementRequest struct {
	*tchttp.BaseRequest
	
	// 全球加速实例ID。
	GlobalAcceleratorId *string `json:"GlobalAcceleratorId,omitnil,omitempty" name:"GlobalAcceleratorId"`

	// 加速地域。
	AccelerateRegion *string `json:"AccelerateRegion,omitnil,omitempty" name:"AccelerateRegion"`

	// 终端节点组地域。
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitnil,omitempty" name:"EndpointGroupRegion"`

	// 账单年月时间。
	SettlementMonth *uint64 `json:"SettlementMonth,omitnil,omitempty" name:"SettlementMonth"`
}

func (r *DescribeCrossBorderSettlementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderSettlementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlobalAcceleratorId")
	delete(f, "AccelerateRegion")
	delete(f, "EndpointGroupRegion")
	delete(f, "SettlementMonth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossBorderSettlementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderSettlementResponseParams struct {
	// 流量用量，单位是GB；精度为保留小数点6位。
	Traffic *float64 `json:"Traffic,omitnil,omitempty" name:"Traffic"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCrossBorderSettlementResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossBorderSettlementResponseParams `json:"Response"`
}

func (r *DescribeCrossBorderSettlementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderSettlementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}