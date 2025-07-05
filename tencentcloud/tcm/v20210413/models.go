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

package v20210413

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type APM struct {
	// 是否启用
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// APM 实例，如果创建时传入的参数为空，则表示自动创建 APM 实例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否要删除APM实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedDelete *bool `json:"NeedDelete,omitnil,omitempty" name:"NeedDelete"`
}

type AccessLogConfig struct {
	// 是否启用
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 采用的模板，可选值：istio（默认）、trace
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// 选中的范围
	SelectedRange *SelectedRange `json:"SelectedRange,omitnil,omitempty" name:"SelectedRange"`

	// 腾讯云日志服务相关参数
	CLS *CLS `json:"CLS,omitnil,omitempty" name:"CLS"`

	// 编码格式，可选值：TEXT、JSON
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 日志格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// GRPC第三方服务器地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 是否启用GRPC第三方服务器
	EnableServer *bool `json:"EnableServer,omitnil,omitempty" name:"EnableServer"`

	// 是否启用标准输出
	EnableStdout *bool `json:"EnableStdout,omitnil,omitempty" name:"EnableStdout"`
}

type ActiveOperation struct {
	// 操作Id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 操作类型，取值范围：
	// - LINK_CLUSTERS: 关联集群
	// - RELINK_CLUSTERS: 重新关联集群
	// - UNLINK_CLUSTERS: 解关联集群
	// - INSTALL_MESH: 安装网格
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type AutoInjectionNamespaceState struct {
	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 注入状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type CLS struct {
	// 是否启用
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 日志集
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 日志主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 是否删除
	NeedDelete *bool `json:"NeedDelete,omitnil,omitempty" name:"NeedDelete"`

	// cls 主题创建的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type Cluster struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 集群角色，取值范围：
	// - MASTER：控制面所在的主集群
	// - REMOTE：主集群管理的远端集群
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 名称，只读
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 状态，只读
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 关联时间，只读
	LinkedTime *string `json:"LinkedTime,omitnil,omitempty" name:"LinkedTime"`

	// 集群配置
	Config *ClusterConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 详细状态，只读
	Status *ClusterStatus `json:"Status,omitnil,omitempty" name:"Status"`

	// 类型，取值范围：
	// - TKE
	// - EKS
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 集群关联的 Namespace 列表
	HostedNamespaces []*string `json:"HostedNamespaces,omitnil,omitempty" name:"HostedNamespaces"`
}

type ClusterConfig struct {
	// 自动注入SideCar的NameSpace
	AutoInjectionNamespaceList []*string `json:"AutoInjectionNamespaceList,omitnil,omitempty" name:"AutoInjectionNamespaceList"`

	// Ingress配置列表
	IngressGatewayList []*IngressGateway `json:"IngressGatewayList,omitnil,omitempty" name:"IngressGatewayList"`

	// Egress配置列表
	EgressGatewayList []*EgressGateway `json:"EgressGatewayList,omitnil,omitempty" name:"EgressGatewayList"`

	// Istiod配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Istiod *IstiodConfig `json:"Istiod,omitnil,omitempty" name:"Istiod"`

	// 部署配置
	DeployConfig *DeployConfig `json:"DeployConfig,omitnil,omitempty" name:"DeployConfig"`

	// 自动注入命名空间状态列表
	AutoInjectionNamespaceStateList []*AutoInjectionNamespaceState `json:"AutoInjectionNamespaceStateList,omitnil,omitempty" name:"AutoInjectionNamespaceStateList"`
}

type ClusterStatus struct {
	// 关联状态，取值范围：
	// - LINKING: 关联中
	// - LINKED: 已关联
	// - UNLINKING: 解关联中
	// - LINK_FAILED: 关联失败
	// - UNLINK_FAILED: 解关联失败
	LinkState *string `json:"LinkState,omitnil,omitempty" name:"LinkState"`

	// 关联错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkErrorDetail *string `json:"LinkErrorDetail,omitnil,omitempty" name:"LinkErrorDetail"`
}

// Predefined struct for user
type CreateMeshRequestParams struct {
	// Mesh名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// Mesh版本
	MeshVersion *string `json:"MeshVersion,omitnil,omitempty" name:"MeshVersion"`

	// Mesh类型，取值范围：
	// - HOSTED：托管网格
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Mesh配置
	Config *MeshConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 关联集群
	ClusterList []*Cluster `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type CreateMeshRequest struct {
	*tchttp.BaseRequest
	
	// Mesh名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// Mesh版本
	MeshVersion *string `json:"MeshVersion,omitnil,omitempty" name:"MeshVersion"`

	// Mesh类型，取值范围：
	// - HOSTED：托管网格
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Mesh配置
	Config *MeshConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 关联集群
	ClusterList []*Cluster `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *CreateMeshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMeshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayName")
	delete(f, "MeshVersion")
	delete(f, "Type")
	delete(f, "Config")
	delete(f, "ClusterList")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMeshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMeshResponseParams struct {
	// 创建的Mesh的Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMeshResponse struct {
	*tchttp.BaseResponse
	Response *CreateMeshResponseParams `json:"Response"`
}

func (r *CreateMeshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMeshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossRegionConfig struct {

}

type CustomPromConfig struct {
	// Prometheus 访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 认证方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 是否公网地址，缺省为 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublicAddr *bool `json:"IsPublicAddr,omitnil,omitempty" name:"IsPublicAddr"`

	// 虚拟网络id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Prometheus 用户名（用于 basic 认证方式）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Prometheus 密码（用于 basic 认证方式）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

// Predefined struct for user
type DeleteMeshRequestParams struct {
	// 需要删除的MeshId
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// CLS组件是否被删除
	NeedDeleteCLS *bool `json:"NeedDeleteCLS,omitnil,omitempty" name:"NeedDeleteCLS"`

	// TMP组件是否被删除
	NeedDeleteTMP *bool `json:"NeedDeleteTMP,omitnil,omitempty" name:"NeedDeleteTMP"`

	// APM组件是否被删除
	NeedDeleteAPM *bool `json:"NeedDeleteAPM,omitnil,omitempty" name:"NeedDeleteAPM"`

	// Grafana组件是否被删除
	NeedDeleteGrafana *bool `json:"NeedDeleteGrafana,omitnil,omitempty" name:"NeedDeleteGrafana"`
}

type DeleteMeshRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的MeshId
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// CLS组件是否被删除
	NeedDeleteCLS *bool `json:"NeedDeleteCLS,omitnil,omitempty" name:"NeedDeleteCLS"`

	// TMP组件是否被删除
	NeedDeleteTMP *bool `json:"NeedDeleteTMP,omitnil,omitempty" name:"NeedDeleteTMP"`

	// APM组件是否被删除
	NeedDeleteAPM *bool `json:"NeedDeleteAPM,omitnil,omitempty" name:"NeedDeleteAPM"`

	// Grafana组件是否被删除
	NeedDeleteGrafana *bool `json:"NeedDeleteGrafana,omitnil,omitempty" name:"NeedDeleteGrafana"`
}

func (r *DeleteMeshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMeshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	delete(f, "NeedDeleteCLS")
	delete(f, "NeedDeleteTMP")
	delete(f, "NeedDeleteAPM")
	delete(f, "NeedDeleteGrafana")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMeshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMeshResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMeshResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMeshResponseParams `json:"Response"`
}

func (r *DeleteMeshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMeshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployConfig struct {
	// 部署类型，取值范围：
	// - SPECIFIC：专有模式
	// - AUTO：普通模式
	NodeSelectType *string `json:"NodeSelectType,omitnil,omitempty" name:"NodeSelectType"`

	// 指定的节点
	Nodes []*string `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

// Predefined struct for user
type DescribeAccessLogConfigRequestParams struct {
	// mesh名字
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`
}

type DescribeAccessLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// mesh名字
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`
}

func (r *DescribeAccessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessLogConfigResponseParams struct {
	// 访问日志输出路径。默认 /dev/stdout
	File *string `json:"File,omitnil,omitempty" name:"File"`

	// 访问日志的格式。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 访问日志输出编码，可取值为 "TEXT" 或 "JSON"，默认 TEXT"
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 选中的范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectedRange *SelectedRange `json:"SelectedRange,omitnil,omitempty" name:"SelectedRange"`

	// 采用的模板，可取值为"istio" 或 "trace"，默认为"istio"
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// 腾讯云日志服务相关参数
	CLS *CLS `json:"CLS,omitnil,omitempty" name:"CLS"`

	// GRPC第三方服务器地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 是否启用GRPC第三方服务器
	EnableServer *bool `json:"EnableServer,omitnil,omitempty" name:"EnableServer"`

	// 是否启用标准输出
	EnableStdout *bool `json:"EnableStdout,omitnil,omitempty" name:"EnableStdout"`

	// 是否启用访问日志采集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessLogConfigResponseParams `json:"Response"`
}

func (r *DescribeAccessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMeshListRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeMeshListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeMeshListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMeshListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMeshListResponseParams struct {
	// 查询到的网格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MeshList []*Mesh `json:"MeshList,omitnil,omitempty" name:"MeshList"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMeshListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMeshListResponseParams `json:"Response"`
}

func (r *DescribeMeshListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMeshRequestParams struct {
	// 需要查询的网格 Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`
}

type DescribeMeshRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的网格 Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`
}

func (r *DescribeMeshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMeshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMeshResponseParams struct {
	// Mesh详细信息
	Mesh *Mesh `json:"Mesh,omitnil,omitempty" name:"Mesh"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMeshResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMeshResponseParams `json:"Response"`
}

func (r *DescribeMeshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EgressGateway struct {
	// Egress名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所在的Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 工作负载配置
	Workload *WorkloadConfig `json:"Workload,omitnil,omitempty" name:"Workload"`

	// 工作负载的状态
	Status *EgressGatewayStatus `json:"Status,omitnil,omitempty" name:"Status"`
}

type EgressGatewayStatus struct {
	// egress gateway的当前版本
	CurrentVersion *string `json:"CurrentVersion,omitnil,omitempty" name:"CurrentVersion"`

	// egress gateway的目标版本
	DesiredVersion *string `json:"DesiredVersion,omitnil,omitempty" name:"DesiredVersion"`

	// egress gateway的状态，取值：running，upgrading，rollbacking
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type ExtensiveCluster struct {
	// Cluster ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ExtensiveClusters struct {
	// 4层集群配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	L4Clusters []*ExtensiveCluster `json:"L4Clusters,omitnil,omitempty" name:"L4Clusters"`

	// 7层集群配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	L7Clusters []*ExtensiveCluster `json:"L7Clusters,omitnil,omitempty" name:"L7Clusters"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type GrafanaInfo struct {
	// 是否开启
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 内网地址
	InternalURL *string `json:"InternalURL,omitnil,omitempty" name:"InternalURL"`

	// 公网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicURL *string `json:"PublicURL,omitnil,omitempty" name:"PublicURL"`

	// 公网失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicFailedReason *string `json:"PublicFailedReason,omitnil,omitempty" name:"PublicFailedReason"`

	// 公网失败详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicFailedMessage *string `json:"PublicFailedMessage,omitnil,omitempty" name:"PublicFailedMessage"`
}

type HorizontalPodAutoscalerSpec struct {
	// 最小副本数
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// 最大副本数
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 用于计算副本数的指标
	Metrics []*MetricSpec `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

type IngressGateway struct {
	// IngressGateway 实例名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Service 配置
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// Workload 配置
	Workload *WorkloadConfig `json:"Workload,omitnil,omitempty" name:"Workload"`

	// 负载均衡配置，自动创建 CLB 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitnil,omitempty" name:"LoadBalancer"`

	// IngressGateway 状态信息，只读
	Status *IngressGatewayStatus `json:"Status,omitnil,omitempty" name:"Status"`

	// 负载均衡实例ID，使用已有 CLB 时返回
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type IngressGatewayStatus struct {
	// 负载均衡实例状态
	LoadBalancer *LoadBalancerStatus `json:"LoadBalancer,omitnil,omitempty" name:"LoadBalancer"`

	// ingress gateway 当前的版本
	CurrentVersion *string `json:"CurrentVersion,omitnil,omitempty" name:"CurrentVersion"`

	// ingress gateway 目标的版本
	DesiredVersion *string `json:"DesiredVersion,omitnil,omitempty" name:"DesiredVersion"`

	// ingress gateway的状态，取值running, upgrading, rollbacking
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type InjectConfig struct {
	// 不需要进行代理的 ip 地址范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeIPRanges []*string `json:"ExcludeIPRanges,omitnil,omitempty" name:"ExcludeIPRanges"`

	// 是否等待sidecar启动
	// 注意：此字段可能返回 null，表示取不到有效值。
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitnil,omitempty" name:"HoldApplicationUntilProxyStarts"`

	// 是否允许sidecar等待
	// 注意：此字段可能返回 null，表示取不到有效值。
	HoldProxyUntilApplicationEnds *bool `json:"HoldProxyUntilApplicationEnds,omitnil,omitempty" name:"HoldProxyUntilApplicationEnds"`
}

type IstioConfig struct {
	// 外部流量策略
	OutboundTrafficPolicy *string `json:"OutboundTrafficPolicy,omitnil,omitempty" name:"OutboundTrafficPolicy"`

	// 调用链配置（Deprecated，请使用 MeshConfig.Tracing 进行配置）
	Tracing *TracingConfig `json:"Tracing,omitnil,omitempty" name:"Tracing"`

	// 禁用策略检查功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisablePolicyChecks *bool `json:"DisablePolicyChecks,omitnil,omitempty" name:"DisablePolicyChecks"`

	// 支持HTTP1.0协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnablePilotHTTP *bool `json:"EnablePilotHTTP,omitnil,omitempty" name:"EnablePilotHTTP"`

	// 禁用HTTP重试策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableHTTPRetry *bool `json:"DisableHTTPRetry,omitnil,omitempty" name:"DisableHTTPRetry"`

	// SmartDNS策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartDNS *SmartDNSConfig `json:"SmartDNS,omitnil,omitempty" name:"SmartDNS"`
}

type IstiodConfig struct {
	// 工作负载配置
	Workload *WorkloadConfig `json:"Workload,omitnil,omitempty" name:"Workload"`
}

// Predefined struct for user
type LinkClusterListRequestParams struct {
	// 网格Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 关联集群
	ClusterList []*Cluster `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`
}

type LinkClusterListRequest struct {
	*tchttp.BaseRequest
	
	// 网格Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 关联集群
	ClusterList []*Cluster `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`
}

func (r *LinkClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	delete(f, "ClusterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LinkClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LinkClusterListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LinkClusterListResponse struct {
	*tchttp.BaseResponse
	Response *LinkClusterListResponseParams `json:"Response"`
}

func (r *LinkClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LinkPrometheusRequestParams struct {
	// 网格ID
	MeshID *string `json:"MeshID,omitnil,omitempty" name:"MeshID"`

	// 配置
	Prometheus *PrometheusConfig `json:"Prometheus,omitnil,omitempty" name:"Prometheus"`
}

type LinkPrometheusRequest struct {
	*tchttp.BaseRequest
	
	// 网格ID
	MeshID *string `json:"MeshID,omitnil,omitempty" name:"MeshID"`

	// 配置
	Prometheus *PrometheusConfig `json:"Prometheus,omitnil,omitempty" name:"Prometheus"`
}

func (r *LinkPrometheusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkPrometheusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshID")
	delete(f, "Prometheus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LinkPrometheusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LinkPrometheusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LinkPrometheusResponse struct {
	*tchttp.BaseResponse
	Response *LinkPrometheusResponseParams `json:"Response"`
}

func (r *LinkPrometheusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkPrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadBalancer struct {
	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	// 只读。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例所在的子网（仅对内网VPC型LB有意义），只读。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// TRAFFIC_POSTPAID_BY_HOUR 按流量按小时后计费 ; BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费;只读。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 最大出带宽，单位Mbps，仅对公网属性的LB生效，默认值 10
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 可用区 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneID *string `json:"ZoneID,omitnil,omitempty" name:"ZoneID"`

	// 运营商类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// TGW Group 名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TgwGroupName *string `json:"TgwGroupName,omitnil,omitempty" name:"TgwGroupName"`

	// IP 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 内网独占集群配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtensiveClusters *ExtensiveClusters `json:"ExtensiveClusters,omitnil,omitempty" name:"ExtensiveClusters"`

	// 负载均衡跨地域配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossRegionConfig *CrossRegionConfig `json:"CrossRegionConfig,omitnil,omitempty" name:"CrossRegionConfig"`

	// 设置跨可用区容灾时的主可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZoneID *string `json:"MasterZoneID,omitnil,omitempty" name:"MasterZoneID"`

	// 设置跨可用区容灾时的备可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZoneID *string `json:"SlaveZoneID,omitnil,omitempty" name:"SlaveZoneID"`
}

type LoadBalancerStatus struct {
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名字
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例 VIP
	LoadBalancerVip *string `json:"LoadBalancerVip,omitnil,omitempty" name:"LoadBalancerVip"`

	// 负载均衡实例 Hostname
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerHostname *string `json:"LoadBalancerHostname,omitnil,omitempty" name:"LoadBalancerHostname"`
}

type Mesh struct {
	// Mesh实例Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// Mesh名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// Mesh类型，取值范围：
	// - STANDALONE：独立网格
	// - HOSTED：托管网格
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Mesh状态，取值范围：
	// - PENDING：等待中
	// - CREATING：创建中
	// - RUNNING：运行中
	// - ABNORMAL：异常
	// - UPGRADING：升级中
	// - CANARY_UPGRADED：升级灰度完成
	// - ROLLBACKING：升级回滚
	// - DELETING：删除中
	// - CREATE_FAILED：安装失败
	// - DELETE_FAILED：删除失败
	// - UPGRADE_FAILED：升级失败
	// - ROLLBACK_FAILED：回滚失败
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 集群列表
	ClusterList []*Cluster `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`

	// Mesh配置
	Config *MeshConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Mesh详细状态
	Status *MeshStatus `json:"Status,omitnil,omitempty" name:"Status"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type MeshConfig struct {
	// Istio配置
	Istio *IstioConfig `json:"Istio,omitnil,omitempty" name:"Istio"`

	// AccessLog配置
	AccessLog *AccessLogConfig `json:"AccessLog,omitnil,omitempty" name:"AccessLog"`

	// Prometheus配置
	Prometheus *PrometheusConfig `json:"Prometheus,omitnil,omitempty" name:"Prometheus"`

	// 自动注入配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Inject *InjectConfig `json:"Inject,omitnil,omitempty" name:"Inject"`

	// 调用跟踪配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tracing *TracingConfig `json:"Tracing,omitnil,omitempty" name:"Tracing"`

	// Sidecar自定义资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SidecarResources *ResourceRequirements `json:"SidecarResources,omitnil,omitempty" name:"SidecarResources"`
}

type MeshStatus struct {
	// 服务数量
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// 灰度升级的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanaryVersion *string `json:"CanaryVersion,omitnil,omitempty" name:"CanaryVersion"`

	// 已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prometheus []*PrometheusStatus `json:"Prometheus,omitnil,omitempty" name:"Prometheus"`

	// 状态附带信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateMessage *string `json:"StateMessage,omitnil,omitempty" name:"StateMessage"`

	// 正在执行的异步操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveOperationList []*ActiveOperation `json:"ActiveOperationList,omitnil,omitempty" name:"ActiveOperationList"`

	// 获取TPS信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TPS *PrometheusStatus `json:"TPS,omitnil,omitempty" name:"TPS"`
}

type MetricSpec struct {
	// 指标来源类型，支持 Pods/Resource
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 使用自定义指标扩进行自动扩缩容
	Pods *PodsMetricSource `json:"Pods,omitnil,omitempty" name:"Pods"`

	// 使用资源指标扩进行自动扩缩容
	Resource *ResourceMetricSource `json:"Resource,omitnil,omitempty" name:"Resource"`
}

// Predefined struct for user
type ModifyAccessLogConfigRequestParams struct {
	// mesh ID
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 选中的范围
	SelectedRange *SelectedRange `json:"SelectedRange,omitnil,omitempty" name:"SelectedRange"`

	// 采用的模板，可选值：istio（默认）、trace、custom
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// 是否启用
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 腾讯云日志服务相关参数
	CLS *CLS `json:"CLS,omitnil,omitempty" name:"CLS"`

	// 编码格式，可选值：TEXT、JSON
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 日志格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否启用标准输出
	EnableStdout *bool `json:"EnableStdout,omitnil,omitempty" name:"EnableStdout"`

	// 是否启动GRPC第三方服务器
	EnableServer *bool `json:"EnableServer,omitnil,omitempty" name:"EnableServer"`

	// GRPC第三方服务器地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

type ModifyAccessLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// mesh ID
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 选中的范围
	SelectedRange *SelectedRange `json:"SelectedRange,omitnil,omitempty" name:"SelectedRange"`

	// 采用的模板，可选值：istio（默认）、trace、custom
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// 是否启用
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 腾讯云日志服务相关参数
	CLS *CLS `json:"CLS,omitnil,omitempty" name:"CLS"`

	// 编码格式，可选值：TEXT、JSON
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 日志格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否启用标准输出
	EnableStdout *bool `json:"EnableStdout,omitnil,omitempty" name:"EnableStdout"`

	// 是否启动GRPC第三方服务器
	EnableServer *bool `json:"EnableServer,omitnil,omitempty" name:"EnableServer"`

	// GRPC第三方服务器地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

func (r *ModifyAccessLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	delete(f, "SelectedRange")
	delete(f, "Template")
	delete(f, "Enable")
	delete(f, "CLS")
	delete(f, "Encoding")
	delete(f, "Format")
	delete(f, "EnableStdout")
	delete(f, "EnableServer")
	delete(f, "Address")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessLogConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccessLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessLogConfigResponseParams `json:"Response"`
}

func (r *ModifyAccessLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMeshRequestParams struct {
	// 需要修改的网格Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 修改的网格名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 修改的网格配置
	Config *MeshConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 修改的集群配置
	ClusterList []*Cluster `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`
}

type ModifyMeshRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的网格Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 修改的网格名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 修改的网格配置
	Config *MeshConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 修改的集群配置
	ClusterList []*Cluster `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`
}

func (r *ModifyMeshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMeshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	delete(f, "DisplayName")
	delete(f, "Config")
	delete(f, "ClusterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMeshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMeshResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMeshResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMeshResponseParams `json:"Response"`
}

func (r *ModifyMeshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMeshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTracingConfigRequestParams struct {
	// mesh名字
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 是否启用调用跟踪
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 腾讯云 APM 服务相关参数
	APM *APM `json:"APM,omitnil,omitempty" name:"APM"`

	// 调用跟踪采样值
	Sampling *float64 `json:"Sampling,omitnil,omitempty" name:"Sampling"`

	// 调用追踪Zipkin相关配置
	Zipkin *TracingZipkin `json:"Zipkin,omitnil,omitempty" name:"Zipkin"`
}

type ModifyTracingConfigRequest struct {
	*tchttp.BaseRequest
	
	// mesh名字
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 是否启用调用跟踪
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 腾讯云 APM 服务相关参数
	APM *APM `json:"APM,omitnil,omitempty" name:"APM"`

	// 调用跟踪采样值
	Sampling *float64 `json:"Sampling,omitnil,omitempty" name:"Sampling"`

	// 调用追踪Zipkin相关配置
	Zipkin *TracingZipkin `json:"Zipkin,omitnil,omitempty" name:"Zipkin"`
}

func (r *ModifyTracingConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTracingConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	delete(f, "Enable")
	delete(f, "APM")
	delete(f, "Sampling")
	delete(f, "Zipkin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTracingConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTracingConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTracingConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTracingConfigResponseParams `json:"Response"`
}

func (r *ModifyTracingConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTracingConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodsMetricSource struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 目标值
	TargetAverageValue *string `json:"TargetAverageValue,omitnil,omitempty" name:"TargetAverageValue"`
}

type PrometheusConfig struct {
	// 虚拟网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 关联已存在实例Id，不填则默认创建
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 第三方 Prometheus
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomProm *CustomPromConfig `json:"CustomProm,omitnil,omitempty" name:"CustomProm"`
}

type PrometheusStatus struct {
	// Prometheus Id
	PrometheusId *string `json:"PrometheusId,omitnil,omitempty" name:"PrometheusId"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 虚拟网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Grafana信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grafana *GrafanaInfo `json:"Grafana,omitnil,omitempty" name:"Grafana"`

	// Prometheus 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Resource struct {
	// 资源类型 cpu/memory
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源数量
	Quantity *string `json:"Quantity,omitnil,omitempty" name:"Quantity"`
}

type ResourceMetricSource struct {
	// 资源名称 cpu/memory
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 目标平均利用率
	TargetAverageUtilization *int64 `json:"TargetAverageUtilization,omitnil,omitempty" name:"TargetAverageUtilization"`

	// 目标平均值
	TargetAverageValue *string `json:"TargetAverageValue,omitnil,omitempty" name:"TargetAverageValue"`
}

type ResourceRequirements struct {
	// Limits 描述了允许的最大计算资源量。
	Limits []*Resource `json:"Limits,omitnil,omitempty" name:"Limits"`

	// Requests 描述所需的最小计算资源量。
	Requests []*Resource `json:"Requests,omitnil,omitempty" name:"Requests"`
}

type SelectedItems struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 选中项目名字
	ItemName *string `json:"ItemName,omitnil,omitempty" name:"ItemName"`

	// ingress gw的名称列表
	Gateways []*string `json:"Gateways,omitnil,omitempty" name:"Gateways"`
}

type SelectedRange struct {
	// 选中的项目详细内容
	Items []*SelectedItems `json:"Items,omitnil,omitempty" name:"Items"`

	// 是否全选
	All *bool `json:"All,omitnil,omitempty" name:"All"`
}

type Service struct {
	// ClusterIP/NodePort/LoadBalancer
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否开启LB直通Pod
	CLBDirectAccess *bool `json:"CLBDirectAccess,omitnil,omitempty" name:"CLBDirectAccess"`

	// 服务是否希望将外部流量路由到节点本地或集群范围的端点。 有两个可用选项：Cluster（默认）和 Local。Cluster 隐藏了客户端源 IP，可能导致第二跳到另一个节点；Local 保留客户端源 IP 并避免 LoadBalancer 和 NodePort 类型服务的第二跳。
	ExternalTrafficPolicy *string `json:"ExternalTrafficPolicy,omitnil,omitempty" name:"ExternalTrafficPolicy"`
}

type SmartDNSConfig struct {
	// 开启DNS代理
	// 注意：此字段可能返回 null，表示取不到有效值。
	IstioMetaDNSCapture *bool `json:"IstioMetaDNSCapture,omitnil,omitempty" name:"IstioMetaDNSCapture"`

	// 开启自动地址分配
	// 注意：此字段可能返回 null，表示取不到有效值。
	IstioMetaDNSAutoAllocate *bool `json:"IstioMetaDNSAutoAllocate,omitnil,omitempty" name:"IstioMetaDNSAutoAllocate"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 是否透传给其他关联产品
	Passthrough *bool `json:"Passthrough,omitnil,omitempty" name:"Passthrough"`
}

type TracingConfig struct {
	// 调用链采样率，百分比
	Sampling *float64 `json:"Sampling,omitnil,omitempty" name:"Sampling"`

	// 是否启用调用跟踪
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 腾讯云 APM 服务相关参数
	APM *APM `json:"APM,omitnil,omitempty" name:"APM"`

	// 启动第三方服务器的地址
	Zipkin *TracingZipkin `json:"Zipkin,omitnil,omitempty" name:"Zipkin"`
}

type TracingZipkin struct {
	// Zipkin调用地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

// Predefined struct for user
type UnlinkClusterRequestParams struct {
	// 网格Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 取消关联的集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type UnlinkClusterRequest struct {
	*tchttp.BaseRequest
	
	// 网格Id
	MeshId *string `json:"MeshId,omitnil,omitempty" name:"MeshId"`

	// 取消关联的集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *UnlinkClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlinkClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlinkClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlinkClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnlinkClusterResponse struct {
	*tchttp.BaseResponse
	Response *UnlinkClusterResponseParams `json:"Response"`
}

func (r *UnlinkClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlinkClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlinkPrometheusRequestParams struct {
	// 网格ID
	MeshID *string `json:"MeshID,omitnil,omitempty" name:"MeshID"`
}

type UnlinkPrometheusRequest struct {
	*tchttp.BaseRequest
	
	// 网格ID
	MeshID *string `json:"MeshID,omitnil,omitempty" name:"MeshID"`
}

func (r *UnlinkPrometheusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlinkPrometheusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlinkPrometheusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlinkPrometheusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnlinkPrometheusResponse struct {
	*tchttp.BaseResponse
	Response *UnlinkPrometheusResponseParams `json:"Response"`
}

func (r *UnlinkPrometheusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlinkPrometheusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WorkloadConfig struct {
	// 工作副本数
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 资源配置
	Resources *ResourceRequirements `json:"Resources,omitnil,omitempty" name:"Resources"`

	// HPA策略
	HorizontalPodAutoscaler *HorizontalPodAutoscalerSpec `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 部署到指定节点
	SelectedNodeList []*string `json:"SelectedNodeList,omitnil,omitempty" name:"SelectedNodeList"`

	// 组件的部署模式，取值说明：
	// IN_GENERAL_NODE：常规节点
	// IN_EKLET：eklet 节点
	// IN_SHARED_NODE_POOL：共享节电池
	// IN_EXCLUSIVE_NODE_POOL：独占节点池
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`
}