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

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 要绑定的安全组 ID，类似sg-efil7***。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例 ID，格式如：vdb-c1nl9***，支持指定多个实例
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要绑定的安全组 ID，类似sg-efil7***。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例 ID，格式如：vdb-c1nl9***，支持指定多个实例
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 私有网络 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络 VPC 的子网 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定实例计费方式。
	// - 0：按量付费。
	// - 1：包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 设置实例名称。仅支持长度不超过 60 的中文/英文/数字/-/_。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 安全组 ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 若计费方式为包年包月，指定包年包月续费的时长。
	// - 单位：月。
	// - 取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。默认为1。
	PayPeriod *int64 `json:"PayPeriod,omitnil,omitempty" name:"PayPeriod"`

	// 若为包年包月计费，需指定是否开启自动续费。
	// - 0：不开启自动续费。
	// - 1：开启自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 实例额外参数，通过json提交。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 以数组形式列出标签信息。
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 指定实例所属项目 ID。
	//
	// Deprecated: Project is deprecated.
	Project *string `json:"Project,omitnil,omitempty" name:"Project"`

	// 产品版本，0-标准版，1-容量增强版
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 实例类型。
	// - base：免费测试版。
	// - single：单机版。
	// - cluster：高可用版。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型为高可用版，需指定可用区选项。
	// - two：两可用区。
	// - three：三可用区。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 购买实例数量。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 网络类型。
	// VPC或TCS
	//
	// Deprecated: NetworkType is deprecated.
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 实例所应用的参数模板 ID。
	//
	// Deprecated: TemplateId is deprecated.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 组件具体设置列表。
	//
	// Deprecated: Components is deprecated.
	Components []*CreateInstancesComponent `json:"Components,omitnil,omitempty" name:"Components"`

	// 实例类型为高可用版，通过该参数指定主可用区。
	//
	// Deprecated: Zone is deprecated.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例类型为高可用版，通过该参数指定备可用区。
	//
	// Deprecated: SlaveZones is deprecated.
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// 是否长期有效
	//
	// Deprecated: IsNoExpired is deprecated.
	IsNoExpired *bool `json:"IsNoExpired,omitnil,omitempty" name:"IsNoExpired"`

	// 引擎名称，业务自定义。
	//
	// Deprecated: EngineName is deprecated.
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 引擎版本，业务自定义。
	//
	// Deprecated: EngineVersion is deprecated.
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 实例描述。
	//
	// Deprecated: Brief is deprecated.
	Brief *string `json:"Brief,omitnil,omitempty" name:"Brief"`

	// 负责人信息。
	//
	// Deprecated: Chief is deprecated.
	Chief *string `json:"Chief,omitnil,omitempty" name:"Chief"`

	// DBA人员信息
	//
	// Deprecated: DBA is deprecated.
	DBA *string `json:"DBA,omitnil,omitempty" name:"DBA"`

	// 指定实例的节点类型。具体信息，请参见[选择节点类型](https://cloud.tencent.com/document/product/1709/113399)。
	// - compute：计费型。
	// - normal：标准型。
	// - store：存储型。
	//
	// Deprecated: NodeType is deprecated.
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 指定实例所需的 CPU 核数。实例类型不同，支持的 CPU 核数存在差异。
	// - 计算型： 1、2、4、8、16、24、32。
	// - 标准型： 1、2、4、8、12、16。
	// - 存储型： 1、2、4、6、8。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 指定实例所需的内存大小。单位：GB。选择具体规格，请参见[配置规格（选型）](https://cloud.tencent.com/document/product/1709/113399)。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 指定实例所需的磁盘大小，单位：GB。选择具体规格，请参见[配置规格（选型）](https://cloud.tencent.com/document/product/1709/113399)。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 指定实例所需配置的节点数量。选择方法，请参见[配置规格（选型）](https://cloud.tencent.com/document/product/1709/113399)。
	WorkerNodeNum *uint64 `json:"WorkerNodeNum,omitnil,omitempty" name:"WorkerNodeNum"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 私有网络 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络 VPC 的子网 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定实例计费方式。
	// - 0：按量付费。
	// - 1：包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 设置实例名称。仅支持长度不超过 60 的中文/英文/数字/-/_。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 安全组 ID。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 若计费方式为包年包月，指定包年包月续费的时长。
	// - 单位：月。
	// - 取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。默认为1。
	PayPeriod *int64 `json:"PayPeriod,omitnil,omitempty" name:"PayPeriod"`

	// 若为包年包月计费，需指定是否开启自动续费。
	// - 0：不开启自动续费。
	// - 1：开启自动续费。
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 实例额外参数，通过json提交。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 以数组形式列出标签信息。
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 指定实例所属项目 ID。
	Project *string `json:"Project,omitnil,omitempty" name:"Project"`

	// 产品版本，0-标准版，1-容量增强版
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 实例类型。
	// - base：免费测试版。
	// - single：单机版。
	// - cluster：高可用版。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型为高可用版，需指定可用区选项。
	// - two：两可用区。
	// - three：三可用区。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 购买实例数量。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 网络类型。
	// VPC或TCS
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 实例所应用的参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 组件具体设置列表。
	Components []*CreateInstancesComponent `json:"Components,omitnil,omitempty" name:"Components"`

	// 实例类型为高可用版，通过该参数指定主可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例类型为高可用版，通过该参数指定备可用区。
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// 是否长期有效
	IsNoExpired *bool `json:"IsNoExpired,omitnil,omitempty" name:"IsNoExpired"`

	// 引擎名称，业务自定义。
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 引擎版本，业务自定义。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 实例描述。
	Brief *string `json:"Brief,omitnil,omitempty" name:"Brief"`

	// 负责人信息。
	Chief *string `json:"Chief,omitnil,omitempty" name:"Chief"`

	// DBA人员信息
	DBA *string `json:"DBA,omitnil,omitempty" name:"DBA"`

	// 指定实例的节点类型。具体信息，请参见[选择节点类型](https://cloud.tencent.com/document/product/1709/113399)。
	// - compute：计费型。
	// - normal：标准型。
	// - store：存储型。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 指定实例所需的 CPU 核数。实例类型不同，支持的 CPU 核数存在差异。
	// - 计算型： 1、2、4、8、16、24、32。
	// - 标准型： 1、2、4、8、12、16。
	// - 存储型： 1、2、4、6、8。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 指定实例所需的内存大小。单位：GB。选择具体规格，请参见[配置规格（选型）](https://cloud.tencent.com/document/product/1709/113399)。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 指定实例所需的磁盘大小，单位：GB。选择具体规格，请参见[配置规格（选型）](https://cloud.tencent.com/document/product/1709/113399)。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 指定实例所需配置的节点数量。选择方法，请参见[配置规格（选型）](https://cloud.tencent.com/document/product/1709/113399)。
	WorkerNodeNum *uint64 `json:"WorkerNodeNum,omitnil,omitempty" name:"WorkerNodeNum"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "PayMode")
	delete(f, "InstanceName")
	delete(f, "SecurityGroupIds")
	delete(f, "PayPeriod")
	delete(f, "AutoRenew")
	delete(f, "Params")
	delete(f, "ResourceTags")
	delete(f, "Project")
	delete(f, "ProductType")
	delete(f, "InstanceType")
	delete(f, "Mode")
	delete(f, "GoodsNum")
	delete(f, "NetworkType")
	delete(f, "TemplateId")
	delete(f, "Components")
	delete(f, "Zone")
	delete(f, "SlaveZones")
	delete(f, "IsNoExpired")
	delete(f, "EngineName")
	delete(f, "EngineVersion")
	delete(f, "Brief")
	delete(f, "Chief")
	delete(f, "DBA")
	delete(f, "NodeType")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "DiskSize")
	delete(f, "WorkerNodeNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesComponent struct {
	// 底层组件名，需要和产品模型中的保持一致
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 组件cpu大小
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 组件内存大小
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 业务节点数
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 自定义组件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 组件磁盘大小
	StorageSize *uint64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 磁盘类型
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 组件额外参数，通过JSON提交
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// 实例ID，格式如：vdb-c1nl9***。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：vdb-c1nl9***。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMaintenanceWindowRequestParams struct {
	// 指定查询维护时间窗的具体实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// 指定查询维护时间窗的具体实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMaintenanceWindowResponseParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护时间窗开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 维护时间窗结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 维护时间窗时长。单位：小时。
	TimeSpan *float64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMaintenanceWindowResponseParams `json:"Response"`
}

func (r *DescribeInstanceMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// component
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// component
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Component")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesResponseParams struct {
	// 实例pod列表。
	Items []*NodeInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 查询结果总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 实例ID数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称，支持模糊搜索。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例模糊搜索字段。
	InstanceKeys []*string `json:"InstanceKeys,omitnil,omitempty" name:"InstanceKeys"`

	// 根据状态获取实例， 为空则获取全部非隔离和非下线的实例。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 按照引擎筛选实例。
	//
	// Deprecated: EngineNames is deprecated.
	EngineNames []*string `json:"EngineNames,omitnil,omitempty" name:"EngineNames"`

	// 按照版本筛选实例。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 按照api版本筛选实例
	ApiVersions []*string `json:"ApiVersions,omitnil,omitempty" name:"ApiVersions"`

	// 按照创建时间筛选实例。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 按照可用区筛选实例。
	//
	// Deprecated: Zones is deprecated.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 排序字段。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 查询开始位置。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照标签筛选实例
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 任务状态：1-待执行任务；2-密钥更新中；3-网络变更中；4-参数变更中；5-embedding变更中；6-ai套件变更中；7-滚动升级中；8-纵向扩容中；9-纵向缩容中；10-横向扩容中；11-横向缩容中
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 根据实例vip搜索实例
	Networks []*string `json:"Networks,omitnil,omitempty" name:"Networks"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称，支持模糊搜索。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例模糊搜索字段。
	InstanceKeys []*string `json:"InstanceKeys,omitnil,omitempty" name:"InstanceKeys"`

	// 根据状态获取实例， 为空则获取全部非隔离和非下线的实例。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 按照引擎筛选实例。
	EngineNames []*string `json:"EngineNames,omitnil,omitempty" name:"EngineNames"`

	// 按照版本筛选实例。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 按照api版本筛选实例
	ApiVersions []*string `json:"ApiVersions,omitnil,omitempty" name:"ApiVersions"`

	// 按照创建时间筛选实例。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 按照可用区筛选实例。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 排序字段。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 查询开始位置。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照标签筛选实例
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 任务状态：1-待执行任务；2-密钥更新中；3-网络变更中；4-参数变更中；5-embedding变更中；6-ai套件变更中；7-滚动升级中；8-纵向扩容中；9-纵向缩容中；10-横向扩容中；11-横向缩容中
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 根据实例vip搜索实例
	Networks []*string `json:"Networks,omitnil,omitempty" name:"Networks"`
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
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "InstanceKeys")
	delete(f, "Status")
	delete(f, "EngineNames")
	delete(f, "EngineVersions")
	delete(f, "ApiVersions")
	delete(f, "CreateAt")
	delete(f, "Zones")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourceTags")
	delete(f, "TaskStatus")
	delete(f, "Networks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例列表。
	Items []*InstanceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DestroyInstancesRequestParams struct {
	// 以数组形式指定待销毁下线的实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DestroyInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 以数组形式指定待销毁下线的实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DestroyInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstancesResponseParams `json:"Response"`
}

func (r *DestroyInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 要绑定的安全组 ID，类似sg-efil****。
	SecurityGroupIds *string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例 ID，格式如：vdb-c1nl****，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要绑定的安全组 ID，类似sg-efil****。
	SecurityGroupIds *string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例 ID，格式如：vdb-c1nl****，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Inbound struct {
	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type InstanceInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例自定义名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户APPID。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 产品。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 网络信息。
	Networks []*Network `json:"Networks,omitnil,omitempty" name:"Networks"`

	// 分片信息。
	ShardNum *uint64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 副本数。
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// CPU.
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存。
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘。
	Disk *uint64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 健康得分。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: HealthScore is deprecated.
	HealthScore *float64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 异常告警。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Warning is deprecated.
	Warning *int64 `json:"Warning,omitnil,omitempty" name:"Warning"`

	// 所属项目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Project is deprecated.
	Project *string `json:"Project,omitnil,omitempty" name:"Project"`

	// 所属标签。
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 创建时间。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 资源状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 引擎名称。
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 引擎版本。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// api版本
	ApiVersion *string `json:"ApiVersion,omitnil,omitempty" name:"ApiVersion"`

	// 计费模式。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 差异化扩展信息, json格式。
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`

	// 过期时间。
	ExpiredAt *string `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`

	// 是否不过期(永久)。
	IsNoExpired *bool `json:"IsNoExpired,omitnil,omitempty" name:"IsNoExpired"`

	// 产品版本，0-标准版，1-容量增强版
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 节点类型
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 外网地址。
	WanAddress *string `json:"WanAddress,omitnil,omitempty" name:"WanAddress"`

	// 隔离时间
	IsolateAt *string `json:"IsolateAt,omitnil,omitempty" name:"IsolateAt"`

	// 是否自动续费。0: 不自动续费(可以支持特权不停服)；1:自动续费；2:到期不续费.
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 任务状态：0-无任务；1-待执行任务；2-密钥更新中；3-网络变更中；4-参数变更中；5-embedding变更中；6-ai套件变更中；7-滚动升级中；8-纵向扩容中；9-纵向缩容中；10-横向扩容中；11-横向缩容中
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 绑定的安全组id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

// Predefined struct for user
type IsolateInstanceRequestParams struct {
	// 指定需隔离的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type IsolateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定需隔离的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *IsolateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateInstanceResponseParams `json:"Response"`
}

func (r *IsolateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 要修改的安全组ID列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例ID，格式如：vdb-c9s3****。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的安全组ID列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例ID，格式如：vdb-c9s3****。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceMaintenanceWindowRequestParams struct {
	// 指定需修改维护时间窗的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护时间窗开始时间。取值范围为"00:00-23:00"的任意整点，如01:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 设置维护时间窗的时长。根据指定的维护时间窗开始时间与时长可确定维护时间窗的范围。
	// - 单位：小时。
	// - 取值范围：3、6、8、10、12。
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

type ModifyInstanceMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// 指定需修改维护时间窗的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护时间窗开始时间。取值范围为"00:00-23:00"的任意整点，如01:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 设置维护时间窗的时长。根据指定的维护时间窗开始时间与时长可确定维护时间窗的范围。
	// - 单位：小时。
	// - 取值范围：3、6、8、10、12。
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

func (r *ModifyInstanceMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceMaintenanceWindowResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceMaintenanceWindowResponseParams `json:"Response"`
}

func (r *ModifyInstanceMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Network struct {
	// VpcId(VPC网络下有效)
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id(VPC网络下有效)。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 内网访问IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网访问Port。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 旧 ip 保留时长，单位天
	PreserveDuration *int64 `json:"PreserveDuration,omitnil,omitempty" name:"PreserveDuration"`

	// 旧 ip 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type NodeInfo struct {
	// Pod名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// pod状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type Outbound struct {
	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type RecoverInstanceRequestParams struct {
	// 指定待恢复的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 若实例为包年包月计费，需通过该参数指定续费的时长。
	// - 单位：月。
	// - 取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。默认为1。
	PayPeriod *int64 `json:"PayPeriod,omitnil,omitempty" name:"PayPeriod"`
}

type RecoverInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定待恢复的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 若实例为包年包月计费，需通过该参数指定续费的时长。
	// - 单位：月。
	// - 取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。默认为1。
	PayPeriod *int64 `json:"PayPeriod,omitnil,omitempty" name:"PayPeriod"`
}

func (r *RecoverInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PayPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecoverInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RecoverInstanceResponseParams `json:"Response"`
}

func (r *RecoverInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 指定需扩容节点数量的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定扩容后的节点数量。选项范围起始于当前实例已有的节点数，上限为30个节点。
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 指定水平扩容增加节点数量的时间。
	// - false：默认值，指在下一个维护时间段内执行增加节点数的任务。实例列表中“状态”列将显示“待执行配置变更”，等到维护时间窗内启动扩容任务。维护时间的更多信息，请参维护时间窗。
	// - true：立即执行增加节点数的任务，请确保此时没有重大业务操作。
	RunNow *bool `json:"RunNow,omitnil,omitempty" name:"RunNow"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定需扩容节点数量的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定扩容后的节点数量。选项范围起始于当前实例已有的节点数，上限为30个节点。
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 指定水平扩容增加节点数量的时间。
	// - false：默认值，指在下一个维护时间段内执行增加节点数的任务。实例列表中“状态”列将显示“待执行配置变更”，等到维护时间窗内启动扩容任务。维护时间的更多信息，请参维护时间窗。
	// - true：立即执行增加节点数的任务，请确保此时没有重大业务操作。
	RunNow *bool `json:"RunNow,omitnil,omitempty" name:"RunNow"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReplicaNum")
	delete(f, "RunNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceRequestParams struct {
	// 指定需升级配置的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定升级配置后的 CPU 核数。
	// - 节点规格可选项（CPU & 内存）必须 >= 当前配置。
	// - 可选择的规格信息，请参见[选择节点规格与数量](https://cloud.tencent.com/document/product/1709/113399)。
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 指定升级配置后的内存大小，单位为：GB。
	// - 节点规格可选项（CPU & 内存）必须 >= 当前配置。
	// - 不同实例类型对 CPU 与内存资源的配置比例有不同的要求。例如，计算型实例，CPU 与内存的分配比例要求为 1:2。CPU 被指定为 4 核，那么内存则应被指定为 8GB。节点规格的详细信息，请参见[选择节点规格与数量](https://cloud.tencent.com/document/product/1709/113399)。
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 指定升级配置后的磁盘容量。
	// - 单位：GB。
	// - 取值范围为：[10,1000]。
	// - 取值必须为10的倍数。
	StorageSize *uint64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 指定垂直扩容升级节点配置的时间。
	// - false：默认值，指在下一个维护时间段内执行升配节点规格的任务。实例列表中“状态”列将显示“待执行配置变更”，等到维护时间窗内启动任务。维护时间的更多信息，请参见维护时间窗。
	// - true：立即执行升级配置的任务，请确保此时没有重大业务操作。
	RunNow *bool `json:"RunNow,omitnil,omitempty" name:"RunNow"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定需升级配置的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定升级配置后的 CPU 核数。
	// - 节点规格可选项（CPU & 内存）必须 >= 当前配置。
	// - 可选择的规格信息，请参见[选择节点规格与数量](https://cloud.tencent.com/document/product/1709/113399)。
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 指定升级配置后的内存大小，单位为：GB。
	// - 节点规格可选项（CPU & 内存）必须 >= 当前配置。
	// - 不同实例类型对 CPU 与内存资源的配置比例有不同的要求。例如，计算型实例，CPU 与内存的分配比例要求为 1:2。CPU 被指定为 4 核，那么内存则应被指定为 8GB。节点规格的详细信息，请参见[选择节点规格与数量](https://cloud.tencent.com/document/product/1709/113399)。
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 指定升级配置后的磁盘容量。
	// - 单位：GB。
	// - 取值范围为：[10,1000]。
	// - 取值必须为10的倍数。
	StorageSize *uint64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 指定垂直扩容升级节点配置的时间。
	// - false：默认值，指在下一个维护时间段内执行升配节点规格的任务。实例列表中“状态”列将显示“待执行配置变更”，等到维护时间窗内启动任务。维护时间的更多信息，请参见维护时间窗。
	// - true：立即执行升级配置的任务，请确保此时没有重大业务操作。
	RunNow *bool `json:"RunNow,omitnil,omitempty" name:"RunNow"`
}

func (r *ScaleUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "StorageSize")
	delete(f, "RunNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpInstanceResponseParams `json:"Response"`
}

func (r *ScaleUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称。
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组备注。
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// 出站规则。
	Outbound []*Outbound `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// 入站规则。
	Inbound []*Inbound `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// 修改时间，时间格式：yyyy-mm-dd hh:mm:ss。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}