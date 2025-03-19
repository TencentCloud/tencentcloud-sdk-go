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

package v20240906

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AssociateTargetGroupsRequestParams struct {
	// 绑定的关系数组。一次请求最多支持20个。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

type AssociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 绑定的关系数组。一次请求最多支持20个。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

func (r *AssociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Associations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateTargetGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateTargetGroupsResponseParams `json:"Response"`
}

func (r *AssociateTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociationItem struct {
	// 关联到的网关负载均衡实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 网关负载均衡实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`
}

// Predefined struct for user
type CreateGatewayLoadBalancerRequestParams struct {
	// 网关负载均衡后端目标设备所属的私有网络 ID，如vpc-azd4dt1c，可以通过 [DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)  接口获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 网关负载均衡后端目标设备所属的私有网络的子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 网关负载均衡实例名称。可支持输入1-60个字符。不填写时默认自动生成。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 创建网关负载均衡的个数，默认值为 1。批量创建数量最大支持10个。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 购买网关负载均衡的同时，给负载均衡打上标签，最大支持20个标签键值对。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 网关负载均衡实例计费类型，当前只支持传POSTPAID_BY_HOUR，默认是POSTPAID_BY_HOUR。
	LBChargeType *string `json:"LBChargeType,omitnil,omitempty" name:"LBChargeType"`
}

type CreateGatewayLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 网关负载均衡后端目标设备所属的私有网络 ID，如vpc-azd4dt1c，可以通过 [DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)  接口获取。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 网关负载均衡后端目标设备所属的私有网络的子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 网关负载均衡实例名称。可支持输入1-60个字符。不填写时默认自动生成。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 创建网关负载均衡的个数，默认值为 1。批量创建数量最大支持10个。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 购买网关负载均衡的同时，给负载均衡打上标签，最大支持20个标签键值对。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 网关负载均衡实例计费类型，当前只支持传POSTPAID_BY_HOUR，默认是POSTPAID_BY_HOUR。
	LBChargeType *string `json:"LBChargeType,omitnil,omitempty" name:"LBChargeType"`
}

func (r *CreateGatewayLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGatewayLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "LoadBalancerName")
	delete(f, "Number")
	delete(f, "Tags")
	delete(f, "LBChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGatewayLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGatewayLoadBalancerResponseParams struct {
	// 由网关负载均衡实例唯一 ID 组成的数组。
	// 存在某些场景，如创建出现延迟时，此字段可能返回为空；此时可以根据接口返回的RequestId或DealName参数，通过DescribeTaskStatus接口查询创建的资源ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGatewayLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CreateGatewayLoadBalancerResponseParams `json:"Response"`
}

func (r *CreateGatewayLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGatewayLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetGroupRequestParams struct {
	// 目标组名称，限定60个字符。
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// 目标组的vpcid属性，不填则使用默认vpc
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 目标组的默认端口， 后续添加服务器时可使用该默认端口。Port和TargetGroupInstances.N中的port二者必填其一。仅支持6081。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 目标组绑定的后端服务器
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`

	// 网关负载均衡目标组协议。
	// - TENCENT_GENEVE ：GENEVE 标准协议
	// - AWS_GENEVE：GENEVE 兼容协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 健康检查设置。
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 均衡算法。
	// - IP_HASH_3_ELASTIC：弹性哈希
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// 是否支持全死全活。默认支持。
	AllDeadToAlive *bool `json:"AllDeadToAlive,omitnil,omitempty" name:"AllDeadToAlive"`
}

type CreateTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// 目标组名称，限定60个字符。
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// 目标组的vpcid属性，不填则使用默认vpc
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 目标组的默认端口， 后续添加服务器时可使用该默认端口。Port和TargetGroupInstances.N中的port二者必填其一。仅支持6081。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 目标组绑定的后端服务器
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`

	// 网关负载均衡目标组协议。
	// - TENCENT_GENEVE ：GENEVE 标准协议
	// - AWS_GENEVE：GENEVE 兼容协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 健康检查设置。
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 均衡算法。
	// - IP_HASH_3_ELASTIC：弹性哈希
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// 是否支持全死全活。默认支持。
	AllDeadToAlive *bool `json:"AllDeadToAlive,omitnil,omitempty" name:"AllDeadToAlive"`
}

func (r *CreateTargetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupName")
	delete(f, "VpcId")
	delete(f, "Port")
	delete(f, "TargetGroupInstances")
	delete(f, "Protocol")
	delete(f, "HealthCheck")
	delete(f, "ScheduleAlgorithm")
	delete(f, "AllDeadToAlive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetGroupResponseParams struct {
	// 创建目标组后生成的id
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTargetGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateTargetGroupResponseParams `json:"Response"`
}

func (r *CreateTargetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGatewayLoadBalancerRequestParams struct {
	// 要删除的网关负载均衡实例 ID数组，数组大小最大支持20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type DeleteGatewayLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的网关负载均衡实例 ID数组，数组大小最大支持20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

func (r *DeleteGatewayLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGatewayLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGatewayLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGatewayLoadBalancerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGatewayLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGatewayLoadBalancerResponseParams `json:"Response"`
}

func (r *DeleteGatewayLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGatewayLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetGroupsRequestParams struct {
	// 目标组ID列表。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
}

type DeleteTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID列表。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
}

func (r *DeleteTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTargetGroupsResponseParams `json:"Response"`
}

func (r *DeleteTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetGroupInstancesRequestParams struct {
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待解绑的服务器信息。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type DeregisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待解绑的服务器信息。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

func (r *DeregisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterTargetGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetGroupInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterTargetGroupInstancesResponseParams `json:"Response"`
}

func (r *DeregisterTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayLoadBalancersRequestParams struct {
	// 网关负载均衡实例ID。支持批量筛选的实例ID数量上限为20个。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 一次批量返回网关负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回网关负载均衡实例列表的起始偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询负载均衡详细信息列表的过滤条件，每次请求的Filters的上限为10，Filter.Values的上限为100。
	// Filter.Name和Filter.Values皆为必填项。详细的过滤条件如下：
	// - VpcId - String - 是否必填：否 - （过滤条件）按照网关负载均衡实例所属的私有网络过滤，如“vpc-bhqk****”。
	// - Vips - String  - 是否必填：否 - （过滤条件）按照网关负载均衡实例所属的私有网络过滤，如“10.1.1.1”
	// - tag:tag-key - String - 是否必填：否 - （过滤条件）按照GWLB标签键值对进行过滤，tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 搜索字段，模糊匹配名称、VIP。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeGatewayLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 网关负载均衡实例ID。支持批量筛选的实例ID数量上限为20个。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 一次批量返回网关负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回网关负载均衡实例列表的起始偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询负载均衡详细信息列表的过滤条件，每次请求的Filters的上限为10，Filter.Values的上限为100。
	// Filter.Name和Filter.Values皆为必填项。详细的过滤条件如下：
	// - VpcId - String - 是否必填：否 - （过滤条件）按照网关负载均衡实例所属的私有网络过滤，如“vpc-bhqk****”。
	// - Vips - String  - 是否必填：否 - （过滤条件）按照网关负载均衡实例所属的私有网络过滤，如“10.1.1.1”
	// - tag:tag-key - String - 是否必填：否 - （过滤条件）按照GWLB标签键值对进行过滤，tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 搜索字段，模糊匹配名称、VIP。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeGatewayLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayLoadBalancersResponseParams struct {
	// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的网关负载均衡实例数组。
	LoadBalancerSet []*GatewayLoadBalancer `json:"LoadBalancerSet,omitnil,omitempty" name:"LoadBalancerSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewayLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayLoadBalancersResponseParams `json:"Response"`
}

func (r *DescribeGatewayLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstanceStatusRequestParams struct {
	// 目标组唯一id
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组绑定的后端服务ip列表
	TargetGroupInstanceIps []*string `json:"TargetGroupInstanceIps,omitnil,omitempty" name:"TargetGroupInstanceIps"`
}

type DescribeTargetGroupInstanceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 目标组唯一id
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组绑定的后端服务ip列表
	TargetGroupInstanceIps []*string `json:"TargetGroupInstanceIps,omitnil,omitempty" name:"TargetGroupInstanceIps"`
}

func (r *DescribeTargetGroupInstanceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstanceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstanceIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupInstanceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstanceStatusResponseParams struct {
	// 健康检查后端rs状态列表
	TargetGroupInstanceSet []*TargetGroupInstanceStatus `json:"TargetGroupInstanceSet,omitnil,omitempty" name:"TargetGroupInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupInstanceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupInstanceStatusResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupInstanceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstanceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstancesRequestParams struct {
	// 过滤条件，当前仅支持TargetGroupId，BindIP，InstanceId过滤。
	// 
	// - TargetGroupId - String - 是否必填：否 - （过滤条件）目标组ID，如“lbtg-5xunivs0”。
	// - BindIP - String - 是否必填：否 - （过滤条件）目标组绑定实例的内网IP地址，如“10.1.1.1”
	// - InstanceId - String - 是否必填：否 - （过滤条件）目标组绑定实例的名称，如“ins-mxzlf9ke”
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 显示数量限制，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，当前仅支持TargetGroupId，BindIP，InstanceId过滤。
	// 
	// - TargetGroupId - String - 是否必填：否 - （过滤条件）目标组ID，如“lbtg-5xunivs0”。
	// - BindIP - String - 是否必填：否 - （过滤条件）目标组绑定实例的内网IP地址，如“10.1.1.1”
	// - InstanceId - String - 是否必填：否 - （过滤条件）目标组绑定实例的名称，如“ins-mxzlf9ke”
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 显示数量限制，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstancesResponseParams struct {
	// 本次查询的结果数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 绑定的服务器信息。
	TargetGroupInstanceSet []*TargetGroupBackend `json:"TargetGroupInstanceSet,omitnil,omitempty" name:"TargetGroupInstanceSet"`

	// 实际统计数量，不受Limit、Offset、CAM的影响。
	RealCount *uint64 `json:"RealCount,omitnil,omitempty" name:"RealCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupInstancesResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupListRequestParams struct {
	// 目标组ID数组。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// 过滤条件数组。
	// 
	// - TargetGroupVpcId - String - 是否必填：否 - （过滤条件）按照目标组所属的私有网络过滤，如“vpc-bhqk****”。
	// - TargetGroupName - String - 是否必填：否 - （过滤条件）按照目标组的名称过滤，如“tg_name”
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTargetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID数组。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// 过滤条件数组。
	// 
	// - TargetGroupVpcId - String - 是否必填：否 - （过滤条件）按照目标组所属的私有网络过滤，如“vpc-bhqk****”。
	// - TargetGroupName - String - 是否必填：否 - （过滤条件）按照目标组的名称过滤，如“tg_name”
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTargetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupListResponseParams struct {
	// 显示的结果数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 显示的目标组信息集合。
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitnil,omitempty" name:"TargetGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupListResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsRequestParams struct {
	// 目标组ID，与Filters互斥。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件数组。
	// 
	// - TargetGroupVpcId - String - 是否必填：否 - （过滤条件）按照目标组所属的私有网络过滤，如“vpc-bhqk****”。
	// - TargetGroupName - String - 是否必填：否 - （过滤条件）按照目标组的名称过滤，如“tg_name”
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID，与Filters互斥。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件数组。
	// 
	// - TargetGroupVpcId - String - 是否必填：否 - （过滤条件）按照目标组所属的私有网络过滤，如“vpc-bhqk****”。
	// - TargetGroupName - String - 是否必填：否 - （过滤条件）按照目标组的名称过滤，如“tg_name”
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsResponseParams struct {
	// 显示的结果数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 显示的目标组信息集合。
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitnil,omitempty" name:"TargetGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupsResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 请求ID，即接口返回的 RequestId 参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 请求ID，即接口返回的 RequestId 参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// 任务的当前状态。 0：成功，1：失败，2：进行中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 由负载均衡实例唯一 ID 组成的数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 辅助描述信息，如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateTargetGroupsRequestParams struct {
	// 待解绑的目标组列表。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

type DisassociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待解绑的目标组列表。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

func (r *DisassociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Associations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateTargetGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateTargetGroupsResponseParams `json:"Response"`
}

func (r *DisassociateTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤器的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤器的值数组
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type GatewayLoadBalancer struct {
	// 网关负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 网关负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 网关负载均衡所属私有网络。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 网关负载均衡所属子网。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 网关负载均衡提供服务的虚拟IP。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 网关负载均衡实例状态。
	// 0：创建中，1：正常运行，3：删除中。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 关联的目标组唯一ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 是否开启删除保护功能。
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`

	// 负载均衡实例的标签信息。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 网关负载均衡实例的计费类型，POSTPAID_BY_HOUR：按量计费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 0：表示未被隔离，1：表示被隔离。
	Isolation *uint64 `json:"Isolation,omitnil,omitempty" name:"Isolation"`

	// 网关负载均衡实例被隔离的时间
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 是否开启配置修改保护功能。
	OperateProtect *bool `json:"OperateProtect,omitnil,omitempty" name:"OperateProtect"`
}

// Predefined struct for user
type InquirePriceCreateGatewayLoadBalancerRequestParams struct {
	// 询价的网关负载均衡实例个数，默认为1
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`
}

type InquirePriceCreateGatewayLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 询价的网关负载均衡实例个数，默认为1
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`
}

func (r *InquirePriceCreateGatewayLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateGatewayLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GoodsNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateGatewayLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateGatewayLoadBalancerResponseParams struct {
	// 该参数表示对应的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceCreateGatewayLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateGatewayLoadBalancerResponseParams `json:"Response"`
}

func (r *InquirePriceCreateGatewayLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateGatewayLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ItemPrice struct {
	// 后付费单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 后付费计价单元，可取值范围： HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 预支费用的原价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 预支费用的折扣价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 后付费的折扣单价，单位:元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 折扣，如20.0代表2折。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`
}

// Predefined struct for user
type ModifyGatewayLoadBalancerAttributeRequestParams struct {
	// 网关负载均衡的唯一ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 网关负载均衡实例名称。可支持输入1-60个字符。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 是否开启删除保护。
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`
}

type ModifyGatewayLoadBalancerAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 网关负载均衡的唯一ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 网关负载均衡实例名称。可支持输入1-60个字符。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 是否开启删除保护。
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`
}

func (r *ModifyGatewayLoadBalancerAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatewayLoadBalancerAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "DeleteProtect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGatewayLoadBalancerAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGatewayLoadBalancerAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGatewayLoadBalancerAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGatewayLoadBalancerAttributeResponseParams `json:"Response"`
}

func (r *ModifyGatewayLoadBalancerAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatewayLoadBalancerAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupAttributeRequestParams struct {
	// 目标组的ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组的新名称。
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// 健康检查详情。
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 是否支持全死全活。
	AllDeadToAlive *bool `json:"AllDeadToAlive,omitnil,omitempty" name:"AllDeadToAlive"`
}

type ModifyTargetGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 目标组的ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组的新名称。
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// 健康检查详情。
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 是否支持全死全活。
	AllDeadToAlive *bool `json:"AllDeadToAlive,omitnil,omitempty" name:"AllDeadToAlive"`
}

func (r *ModifyTargetGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupName")
	delete(f, "HealthCheck")
	delete(f, "AllDeadToAlive")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyTargetGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupInstancesWeightRequestParams struct {
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 实例绑定配置数组。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type ModifyTargetGroupInstancesWeightRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 实例绑定配置数组。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

func (r *ModifyTargetGroupInstancesWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupInstancesWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupInstancesWeightResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupInstancesWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupInstancesWeightResponseParams `json:"Response"`
}

func (r *ModifyTargetGroupInstancesWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Price struct {
	// 描述了实例价格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 描述了GLCU的价格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LcuPrice *ItemPrice `json:"LcuPrice,omitnil,omitempty" name:"LcuPrice"`
}

// Predefined struct for user
type RegisterTargetGroupInstancesRequestParams struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 服务器实例数组
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type RegisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 服务器实例数组
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

func (r *RegisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterTargetGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetGroupInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RegisterTargetGroupInstancesResponseParams `json:"Response"`
}

func (r *RegisterTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TargetGroupAssociation struct {
	// 网关负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`
}

type TargetGroupBackend struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 后端服务的类型，可取：CVM、ENI
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 后端服务的唯一 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 后端服务的监听端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务的转发权重，取值为0或16
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 后端服务的外网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 后端服务的内网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 后端服务的实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 后端服务被绑定的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// 弹性网卡唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`

	// 后端服务的可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type TargetGroupHealthCheck struct {
	// 是否开启健康检查。
	HealthSwitch *bool `json:"HealthSwitch,omitnil,omitempty" name:"HealthSwitch"`

	// 健康检查使用的协议。支持PING和TCP两种方式，默认为PING。
	// 
	// - icmp: 使用PING的方式进行健康检查
	// - tcp: 使用TCP连接的方式进行健康检查
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 健康检查端口，探测协议为tcp时，该参数必填。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 健康检查超时时间。 默认为2秒。 可配置范围：2 - 30秒。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 检测间隔时间。 默认为5秒。 可配置范围：2 - 300秒。
	IntervalTime *int64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// 检测健康阈值。 默认为3秒。 可配置范围：2 - 10次。
	HealthNum *int64 `json:"HealthNum,omitnil,omitempty" name:"HealthNum"`

	// 检测不健康阈值。 默认为3秒。 可配置范围：2 - 10次。
	UnHealthNum *int64 `json:"UnHealthNum,omitnil,omitempty" name:"UnHealthNum"`
}

type TargetGroupInfo struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组的vpcid
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 目标组的名字
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// 目标组的默认端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 目标组的创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 目标组的修改时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 关联到的规则数组。在DescribeTargetGroupList接口调用时无法获取到该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedRule []*AssociationItem `json:"AssociatedRule,omitnil,omitempty" name:"AssociatedRule"`

	// 网关负载均衡目标组协议。
	// - tencent_geneve ：GENEVE 标准协议
	// - aws_geneve：GENEVE 兼容协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 均衡算法。
	// - ip_hash_3_elastic：弹性哈希
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// 健康检查详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 是否支持全死全活。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllDeadToAlive *bool `json:"AllDeadToAlive,omitnil,omitempty" name:"AllDeadToAlive"`

	// 目标组已关联的规则数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedRuleCount *int64 `json:"AssociatedRuleCount,omitnil,omitempty" name:"AssociatedRuleCount"`

	// 目标组内的实例数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisteredInstancesCount *int64 `json:"RegisteredInstancesCount,omitnil,omitempty" name:"RegisteredInstancesCount"`
}

type TargetGroupInstance struct {
	// 目标组实例的内网IP。
	BindIP *string `json:"BindIP,omitnil,omitempty" name:"BindIP"`

	// 目标组实例的端口，只支持6081。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 目标组实例的权重，只支持0或16，非0统一按16处理。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type TargetGroupInstanceStatus struct {
	// 后端rs的IP
	InstanceIp *string `json:"InstanceIp,omitnil,omitempty" name:"InstanceIp"`

	// 健康检查状态，参数值及含义如下：
	// ● on：表示探测中。
	// ● off：表示健康检查关闭。
	// ● health：表示健康。
	// ● unhealth：表示异常。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}