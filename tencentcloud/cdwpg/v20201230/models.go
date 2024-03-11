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

package v20201230

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CBSSpec struct {
	// 盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 个数
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type CBSSpecInfo struct {
	// 盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type ChargeProperties struct {
	// 1-需要自动续期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 订单时间范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时间单位，一般为h和m
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 计费类型0-按量计费，1-包年包月
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// PREPAID、POSTPAID_BY_HOUR
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

// Predefined struct for user
type CreateInstanceByApiRequestParams struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 计费方式
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 集群密码
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// 资源信息
	Resources []*ResourceSpecNew `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 标签列表
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 计费方式
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 集群密码
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// 资源信息
	Resources []*ResourceSpecNew `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 标签列表
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "Zone")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ChargeProperties")
	delete(f, "AdminPassword")
	delete(f, "Resources")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceByApiResponseParams struct {
	// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceByApiResponseParams `json:"Response"`
}

func (r *CreateInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInfoRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInfoResponseParams struct {
	// 1
	SimpleInstanceInfo *SimpleInstanceInfo `json:"SimpleInstanceInfo,omitnil,omitempty" name:"SimpleInstanceInfo"`

	// 1
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// 实例描述信息
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateRequestParams struct {
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateResponseParams struct {
	// 集群状态，例如：Serving
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 集群操作创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// 集群操作名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 集群操作进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 当前步骤的名称，例如：”购买资源中“
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// 集群备份任务开启状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStateResponseParams `json:"Response"`
}

func (r *DescribeInstanceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleInstancesRequestParams struct {
	// 11
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 11
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 11
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 11
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 11
	SearchTags []*string `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeSimpleInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 11
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 11
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 11
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 11
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 11
	SearchTags []*string `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

func (r *DescribeSimpleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSimpleInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleInstancesResponseParams struct {
	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancesList []*InstanceSimpleInfoNew `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSimpleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSimpleInstancesResponseParams `json:"Response"`
}

func (r *DescribeSimpleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiRequestParams struct {
	// 实例名称，例如"cdwpg-xxxx"
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称，例如"cdwpg-xxxx"
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiResponseParams struct {
	// 销毁流程Id
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceByApiResponseParams `json:"Response"`
}

func (r *DestroyInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {
	// ID值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// cdwpg-cn或者其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// cdwpg-cn或者其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Running
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateInfo *InstanceStateInfo `json:"InstanceStateInfo,omitnil,omitempty" name:"InstanceStateInfo"`

	// -
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 2022-09-05 20:00:01
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// ap-chongqing
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// ap
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// region
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// zone
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// v3
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 字符集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// 引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// GTM节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GTMNodes []*InstanceNodeGroup `json:"GTMNodes,omitnil,omitempty" name:"GTMNodes"`

	// CN节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CNNodes []*InstanceNodeGroup `json:"CNNodes,omitnil,omitempty" name:"CNNodes"`

	// DN节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DNNodes []*InstanceNodeGroup `json:"DNNodes,omitnil,omitempty" name:"DNNodes"`

	// 备份存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStorage []*InstanceNodeGroup `json:"BackupStorage,omitnil,omitempty" name:"BackupStorage"`

	// FN节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FNNodes []*InstanceNodeGroup `json:"FNNodes,omitnil,omitempty" name:"FNNodes"`
}

type InstanceNodeGroup struct {

}

type InstanceSimpleInfoNew struct {
	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceStateInfo struct {
	// 集群状态，例如：Serving
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 集群操作创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// 集群操作名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 集群操作进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowProgress *int64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 当前步骤的名称，例如：”购买资源中“
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// 集群是否有备份中任务，有为1,无为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOpenStatus *int64 `json:"BackupOpenStatus,omitnil,omitempty" name:"BackupOpenStatus"`
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {
	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 资源数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 磁盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSpec *CBSSpecInfo `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 资源类型，DATA
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ResourceSpecNew struct {
	// 资源名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 资源数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 磁盘信息
	DiskSpec *CBSSpec `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// 资源类型，DATA
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SearchTags struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1表示只输入标签的键，没有输入值；0表示输入键时且输入值
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SimpleInstanceInfo struct {
	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVPCID *string `json:"UserVPCID,omitnil,omitempty" name:"UserVPCID"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSubnetID *string `json:"UserSubnetID,omitnil,omitempty" name:"UserSubnetID"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}