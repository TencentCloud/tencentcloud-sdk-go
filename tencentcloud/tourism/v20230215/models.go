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

package v20230215

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeDrawResourceListRequestParams struct {
	// PageNumber
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// PageSize
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeDrawResourceListRequest struct {
	*tchttp.BaseRequest
	
	// PageNumber
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// PageSize
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeDrawResourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrawResourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDrawResourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrawResourceListResponseParams struct {
	// 返回数据条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回数据数组
	ResourceDrawList []*ResourceDrawListType `json:"ResourceDrawList,omitnil,omitempty" name:"ResourceDrawList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDrawResourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDrawResourceListResponseParams `json:"Response"`
}

func (r *DescribeDrawResourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrawResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceDrawListType struct {
	// 记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 资源记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 本订单资源序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// 客户的uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 小订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmallOrderId *string `json:"SmallOrderId,omitnil,omitempty" name:"SmallOrderId"`

	// 资源创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceNewStartTime *string `json:"ResourceNewStartTime,omitnil,omitempty" name:"ResourceNewStartTime"`

	// 资源到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceNewEndTime *string `json:"ResourceNewEndTime,omitnil,omitempty" name:"ResourceNewEndTime"`

	// 资源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *int64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// 本记录状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *uint64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}