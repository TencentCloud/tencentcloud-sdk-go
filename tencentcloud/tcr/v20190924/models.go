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

package v20190924

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessVpc struct {
	// Vpc的Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 内网接入状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 内网接入Ip
	AccessIp *string `json:"AccessIp,omitempty" name:"AccessIp"`
}

type AutoDelStrategyInfo struct {
	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Value *int64 `json:"Value,omitempty" name:"Value"`

	// Valid
	Valid *int64 `json:"Valid,omitempty" name:"Valid"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`
}

type AutoDelStrategyInfoResp struct {
	// 总数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自动删除策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyInfo []*AutoDelStrategyInfo `json:"StrategyInfo,omitempty" name:"StrategyInfo"`
}

// Predefined struct for user
type BatchDeleteImagePersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag列表
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type BatchDeleteImagePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag列表
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

func (r *BatchDeleteImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteImagePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteImagePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteImagePersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteImagePersonalResponseParams `json:"Response"`
}

func (r *BatchDeleteImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteRepositoryPersonalRequestParams struct {
	// 仓库名称数组
	RepoNames []*string `json:"RepoNames,omitempty" name:"RepoNames"`
}

type BatchDeleteRepositoryPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称数组
	RepoNames []*string `json:"RepoNames,omitempty" name:"RepoNames"`
}

func (r *BatchDeleteRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteRepositoryPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteRepositoryPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteRepositoryPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteRepositoryPersonalResponseParams `json:"Response"`
}

func (r *BatchDeleteRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CVEWhitelistItem struct {
	// 漏洞白名单 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`
}

// Predefined struct for user
type CheckInstanceNameRequestParams struct {
	// 待创建的实例名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`
}

type CheckInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 待创建的实例名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`
}

func (r *CheckInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstanceNameResponseParams struct {
	// 检查结果，true为合法，false为非法
	IsValidated *bool `json:"IsValidated,omitempty" name:"IsValidated"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *CheckInstanceNameResponseParams `json:"Response"`
}

func (r *CheckInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstanceRequestParams struct {
	// 待检测的实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type CheckInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待检测的实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *CheckInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstanceResponseParams struct {
	// 检查结果，true为合法，false为非法
	IsValidated *bool `json:"IsValidated,omitempty" name:"IsValidated"`

	// 实例所在的RegionId
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CheckInstanceResponseParams `json:"Response"`
}

func (r *CheckInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationTriggerPersonalRequestParams struct {
	// 触发器关联的镜像仓库，library/test格式
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 应用所在TKE集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用所在TKE集群命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 应用所在TKE集群工作负载名称
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 应用所在TKE集群工作负载下容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 应用所在TKE集群地域
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`

	// 触发方式对应的表达式
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
}

type CreateApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 触发器关联的镜像仓库，library/test格式
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 应用所在TKE集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用所在TKE集群命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 应用所在TKE集群工作负载名称
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 应用所在TKE集群工作负载下容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 应用所在TKE集群地域
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`

	// 触发方式对应的表达式
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
}

func (r *CreateApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "TriggerName")
	delete(f, "InvokeMethod")
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "WorkloadType")
	delete(f, "WorkloadName")
	delete(f, "ContainerName")
	delete(f, "ClusterRegion")
	delete(f, "InvokeExpr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationTriggerPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationTriggerPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationTriggerPersonalResponseParams `json:"Response"`
}

func (r *CreateApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageAccelerationServiceRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 创建CFS的归属的VPCID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 创建CFS的归属的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 创建CFS的存储类型，其中 SD 为标准型存储， HP为性能存储。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 可用区名称，例如ap-beijing-1，请参考 概览 文档中的地域与可用区列表
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

type CreateImageAccelerationServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 创建CFS的归属的VPCID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 创建CFS的归属的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 创建CFS的存储类型，其中 SD 为标准型存储， HP为性能存储。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 可用区名称，例如ap-beijing-1，请参考 概览 文档中的地域与可用区列表
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

func (r *CreateImageAccelerationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageAccelerationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "StorageType")
	delete(f, "PGroupId")
	delete(f, "Zone")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageAccelerationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageAccelerationServiceResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageAccelerationServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageAccelerationServiceResponseParams `json:"Response"`
}

func (r *CreateImageAccelerationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageAccelerationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageLifecyclePersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// keep_last_days:保留最近几天的数据;keep_last_nums:保留最近多少个
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Val *int64 `json:"Val,omitempty" name:"Val"`
}

type CreateImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// keep_last_days:保留最近几天的数据;keep_last_nums:保留最近多少个
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Val *int64 `json:"Val,omitempty" name:"Val"`
}

func (r *CreateImageLifecyclePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageLifecyclePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Type")
	delete(f, "Val")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageLifecyclePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageLifecyclePersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageLifecyclePersonalResponseParams `json:"Response"`
}

func (r *CreateImageLifecyclePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageLifecyclePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitempty" name:"Rule"`
}

type CreateImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *CreateImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImmutableTagRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateImmutableTagRulesResponseParams `json:"Response"`
}

func (r *CreateImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCustomizedDomainRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type CreateInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *CreateInstanceCustomizedDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "DomainName")
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceCustomizedDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCustomizedDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *CreateInstanceCustomizedDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 企业版实例名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`

	// 企业版实例类型（basic 基础版；standard 标准版；premium 高级版）
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 实例计费类型，0表示按量计费，1表示预付费，默认为按量计费
	RegistryChargeType *int64 `json:"RegistryChargeType,omitempty" name:"RegistryChargeType"`

	// 预付费自动续费标识和购买时长
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitempty" name:"RegistryChargePrepaid"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitempty" name:"SyncTag"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 企业版实例名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`

	// 企业版实例类型（basic 基础版；standard 标准版；premium 高级版）
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 实例计费类型，0表示按量计费，1表示预付费，默认为按量计费
	RegistryChargeType *int64 `json:"RegistryChargeType,omitempty" name:"RegistryChargeType"`

	// 预付费自动续费标识和购买时长
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitempty" name:"RegistryChargePrepaid"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitempty" name:"SyncTag"`
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
	delete(f, "RegistryName")
	delete(f, "RegistryType")
	delete(f, "TagSpecification")
	delete(f, "RegistryChargeType")
	delete(f, "RegistryChargePrepaid")
	delete(f, "SyncTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 企业版实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateInstanceTokenRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 访问凭证类型，longterm 为长期访问凭证，temp 为临时访问凭证，默认是临时访问凭证，有效期1小时
	TokenType *string `json:"TokenType,omitempty" name:"TokenType"`

	// 长期访问凭证描述信息
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

type CreateInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 访问凭证类型，longterm 为长期访问凭证，temp 为临时访问凭证，默认是临时访问凭证，有效期1小时
	TokenType *string `json:"TokenType,omitempty" name:"TokenType"`

	// 长期访问凭证描述信息
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "TokenType")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceTokenResponseParams struct {
	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitempty" name:"Username"`

	// 访问凭证
	Token *string `json:"Token,omitempty" name:"Token"`

	// 访问凭证过期时间戳，是一个时间戳数字，无单位
	ExpTime *int64 `json:"ExpTime,omitempty" name:"ExpTime"`

	// 长期凭证的TokenId，短期凭证没有TokenId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceTokenResponseParams `json:"Response"`
}

func (r *CreateInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalEndpointDnsRequestParams struct {
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitempty" name:"EniLBIp"`

	// true：为默认域名，公网域名一致
	// false: 使用vpc域名
	// 默认为vpc域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

type CreateInternalEndpointDnsRequest struct {
	*tchttp.BaseRequest
	
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitempty" name:"EniLBIp"`

	// true：为默认域名，公网域名一致
	// false: 使用vpc域名
	// 默认为vpc域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *CreateInternalEndpointDnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalEndpointDnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "EniLBIp")
	delete(f, "UsePublicDomain")
	delete(f, "RegionName")
	delete(f, "RegionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInternalEndpointDnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalEndpointDnsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInternalEndpointDnsResponse struct {
	*tchttp.BaseResponse
	Response *CreateInternalEndpointDnsResponseParams `json:"Response"`
}

func (r *CreateInternalEndpointDnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalEndpointDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultipleSecurityPolicyRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type CreateMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateMultipleSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultipleSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultipleSecurityPolicyResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultipleSecurityPolicyResponseParams `json:"Response"`
}

func (r *CreateMultipleSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespacePersonalRequestParams struct {
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type CreateNamespacePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *CreateNamespacePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespacePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNamespacePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespacePersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNamespacePersonalResponse struct {
	*tchttp.BaseResponse
	Response *CreateNamespacePersonalResponseParams `json:"Response"`
}

func (r *CreateNamespacePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespacePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称（长度2-30个字符，只能包含小写字母、数字及分隔符("."、"_"、"-")，且不能以分隔符开头、结尾或连续）
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 是否公开，true为公开，fale为私有
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 自动扫描级别，true为自动，false为手动
	IsAutoScan *bool `json:"IsAutoScan,omitempty" name:"IsAutoScan"`

	// 安全阻断级别，true为自动，false为手动
	IsPreventVUL *bool `json:"IsPreventVUL,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持low、medium、high
	Severity *string `json:"Severity,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitempty" name:"CVEWhitelistItems"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称（长度2-30个字符，只能包含小写字母、数字及分隔符("."、"_"、"-")，且不能以分隔符开头、结尾或连续）
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 是否公开，true为公开，fale为私有
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 云标签描述
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 自动扫描级别，true为自动，false为手动
	IsAutoScan *bool `json:"IsAutoScan,omitempty" name:"IsAutoScan"`

	// 安全阻断级别，true为自动，false为手动
	IsPreventVUL *bool `json:"IsPreventVUL,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持low、medium、high
	Severity *string `json:"Severity,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitempty" name:"CVEWhitelistItems"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "IsPublic")
	delete(f, "TagSpecification")
	delete(f, "IsAutoScan")
	delete(f, "IsPreventVUL")
	delete(f, "Severity")
	delete(f, "CVEWhitelistItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNamespaceResponseParams `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationInstanceRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 复制实例地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// 复制实例地域名称
	ReplicationRegionName *string `json:"ReplicationRegionName,omitempty" name:"ReplicationRegionName"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitempty" name:"SyncTag"`
}

type CreateReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 复制实例地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// 复制实例地域名称
	ReplicationRegionName *string `json:"ReplicationRegionName,omitempty" name:"ReplicationRegionName"`

	// 是否同步TCR云标签至生成的COS Bucket
	SyncTag *bool `json:"SyncTag,omitempty" name:"SyncTag"`
}

func (r *CreateReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegionId")
	delete(f, "ReplicationRegionName")
	delete(f, "SyncTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationInstanceResponseParams struct {
	// 企业版复制实例Id
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateReplicationInstanceResponseParams `json:"Response"`
}

func (r *CreateReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 是否公共,1:公共,0:私有
	Public *uint64 `json:"Public,omitempty" name:"Public"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateRepositoryPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 是否公共,1:公共,0:私有
	Public *uint64 `json:"Public,omitempty" name:"Public"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Public")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRepositoryPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *CreateRepositoryPersonalResponseParams `json:"Response"`
}

func (r *CreateRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "BriefDescription")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateRepositoryResponseParams `json:"Response"`
}

func (r *CreateRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityPolicyRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "CidrBlock")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityPolicyResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityPolicyResponseParams `json:"Response"`
}

func (r *CreateSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignaturePolicyRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// KMS 密钥
	KmsId *string `json:"KmsId,omitempty" name:"KmsId"`

	// KMS 密钥所属地域
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 用户自定义域名，为空时使用 TCR 实例默认域名生成签名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 禁用加签策略，默认为 false
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`
}

type CreateSignaturePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// KMS 密钥
	KmsId *string `json:"KmsId,omitempty" name:"KmsId"`

	// KMS 密钥所属地域
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 用户自定义域名，为空时使用 TCR 实例默认域名生成签名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 禁用加签策略，默认为 false
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`
}

func (r *CreateSignaturePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignaturePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Name")
	delete(f, "NamespaceName")
	delete(f, "KmsId")
	delete(f, "KmsRegion")
	delete(f, "Domain")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignaturePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignaturePolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSignaturePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSignaturePolicyResponseParams `json:"Response"`
}

func (r *CreateSignaturePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignatureRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// Tag名称
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
}

type CreateSignatureRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// Tag名称
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
}

func (r *CreateSignatureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignatureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignatureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignatureResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSignatureResponse struct {
	*tchttp.BaseResponse
	Response *CreateSignatureResponseParams `json:"Response"`
}

func (r *CreateSignatureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionExecutionRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 版本保留规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 是否模拟执行，默认值为false，即非模拟执行
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type CreateTagRetentionExecutionRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 版本保留规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 是否模拟执行，默认值为false，即非模拟执行
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *CreateTagRetentionExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagRetentionExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionExecutionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTagRetentionExecutionResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagRetentionExecutionResponseParams `json:"Response"`
}

func (r *CreateTagRetentionExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionRuleRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的Id
	NamespaceId *int64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitempty" name:"RetentionRule"`

	// 执行周期，当前只能选择： manual;daily;weekly;monthly
	CronSetting *string `json:"CronSetting,omitempty" name:"CronSetting"`

	// 是否禁用规则，默认值为false
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`
}

type CreateTagRetentionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的Id
	NamespaceId *int64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitempty" name:"RetentionRule"`

	// 执行周期，当前只能选择： manual;daily;weekly;monthly
	CronSetting *string `json:"CronSetting,omitempty" name:"CronSetting"`

	// 是否禁用规则，默认值为false
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`
}

func (r *CreateTagRetentionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceId")
	delete(f, "RetentionRule")
	delete(f, "CronSetting")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagRetentionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagRetentionRuleResponseParams `json:"Response"`
}

func (r *CreateTagRetentionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserPersonalRequestParams struct {
	// 用户密码，密码必须为8到16位
	Password *string `json:"Password,omitempty" name:"Password"`
}

type CreateUserPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 用户密码，密码必须为8到16位
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateUserPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserPersonalResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserPersonalResponseParams `json:"Response"`
}

func (r *CreateUserPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebhookTriggerRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type CreateWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *CreateWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Trigger")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebhookTriggerResponseParams struct {
	// 新建的触发器
	Trigger *WebhookTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebhookTriggerResponseParams `json:"Response"`
}

func (r *CreateWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizedDomainInfo struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 域名名称
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 域名创建状态（SUCCESS, FAILURE, CREATING, DELETING）
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteApplicationTriggerPersonalRequestParams struct {
	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
}

type DeleteApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`
}

func (r *DeleteApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TriggerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationTriggerPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationTriggerPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationTriggerPersonalResponseParams `json:"Response"`
}

func (r *DeleteApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageAccelerateServiceRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DeleteImageAccelerateServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DeleteImageAccelerateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageAccelerateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageAccelerateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageAccelerateServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageAccelerateServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageAccelerateServiceResponseParams `json:"Response"`
}

func (r *DeleteImageAccelerateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageAccelerateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageLifecycleGlobalPersonalRequestParams struct {

}

type DeleteImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteImageLifecycleGlobalPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageLifecycleGlobalPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageLifecycleGlobalPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageLifecycleGlobalPersonalResponseParams `json:"Response"`
}

func (r *DeleteImageLifecycleGlobalPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageLifecyclePersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type DeleteImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteImageLifecyclePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageLifecyclePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageLifecyclePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageLifecyclePersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageLifecyclePersonalResponseParams `json:"Response"`
}

func (r *DeleteImageLifecyclePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageLifecyclePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagePersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type DeleteImagePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DeleteImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImagePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImagePersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImagePersonalResponseParams `json:"Response"`
}

func (r *DeleteImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *DeleteImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	delete(f, "NamespaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageResponseParams `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

func (r *DeleteImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImmutableTagRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImmutableTagRulesResponseParams `json:"Response"`
}

func (r *DeleteImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceCustomizedDomainRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type DeleteInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 自定义域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DeleteInstanceCustomizedDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "DomainName")
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceCustomizedDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceCustomizedDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *DeleteInstanceCustomizedDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 实例id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 是否删除存储桶，默认为false
	DeleteBucket *bool `json:"DeleteBucket,omitempty" name:"DeleteBucket"`

	// 是否dryRun模式，缺省值：false
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 是否删除存储桶，默认为false
	DeleteBucket *bool `json:"DeleteBucket,omitempty" name:"DeleteBucket"`

	// 是否dryRun模式，缺省值：false
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "DeleteBucket")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceTokenRequestParams struct {
	// 实例 ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 访问凭证 ID
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`
}

type DeleteInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 访问凭证 ID
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`
}

func (r *DeleteInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "TokenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceTokenResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceTokenResponseParams `json:"Response"`
}

func (r *DeleteInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInternalEndpointDnsRequestParams struct {
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitempty" name:"EniLBIp"`

	// true：使用默认域名
	// false:  使用带有vpc的域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DeleteInternalEndpointDnsRequest struct {
	*tchttp.BaseRequest
	
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitempty" name:"EniLBIp"`

	// true：使用默认域名
	// false:  使用带有vpc的域名
	UsePublicDomain *bool `json:"UsePublicDomain,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *DeleteInternalEndpointDnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInternalEndpointDnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "EniLBIp")
	delete(f, "UsePublicDomain")
	delete(f, "RegionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInternalEndpointDnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInternalEndpointDnsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteInternalEndpointDnsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInternalEndpointDnsResponseParams `json:"Response"`
}

func (r *DeleteInternalEndpointDnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInternalEndpointDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultipleSecurityPolicyRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

type DeleteMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 安全组策略
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteMultipleSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMultipleSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultipleSecurityPolicyResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMultipleSecurityPolicyResponseParams `json:"Response"`
}

func (r *DeleteMultipleSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespacePersonalRequestParams struct {
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DeleteNamespacePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteNamespacePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespacePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNamespacePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespacePersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNamespacePersonalResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNamespacePersonalResponseParams `json:"Response"`
}

func (r *DeleteNamespacePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespacePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNamespaceResponseParams `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceRequestParams struct {
	// 实例id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 复制实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 复制实例地域Id
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`
}

type DeleteReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 复制实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 复制实例地域Id
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`
}

func (r *DeleteReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReplicationInstanceResponseParams `json:"Response"`
}

func (r *DeleteReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type DeleteRepositoryPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRepositoryPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRepositoryPersonalResponseParams `json:"Response"`
}

func (r *DeleteRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库的名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`
}

type DeleteRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库的名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`
}

func (r *DeleteRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRepositoryResponseParams `json:"Response"`
}

func (r *DeleteRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryTagsRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// Tag列表，单次请求Tag数量最大为20
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

type DeleteRepositoryTagsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// Tag列表，单次请求Tag数量最大为20
	Tags []*string `json:"Tags,omitempty" name:"Tags"`
}

func (r *DeleteRepositoryTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRepositoryTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRepositoryTagsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRepositoryTagsResponseParams `json:"Response"`
}

func (r *DeleteRepositoryTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityPolicyRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 白名单Id
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 白名单版本
	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

type DeleteSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 白名单Id
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 白名单版本
	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

func (r *DeleteSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "PolicyIndex")
	delete(f, "PolicyVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityPolicyResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityPolicyResponseParams `json:"Response"`
}

func (r *DeleteSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSignaturePolicyRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type DeleteSignaturePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *DeleteSignaturePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSignaturePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSignaturePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSignaturePolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSignaturePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSignaturePolicyResponseParams `json:"Response"`
}

func (r *DeleteSignaturePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRetentionRuleRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 版本保留规则的Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`
}

type DeleteTagRetentionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 版本保留规则的Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`
}

func (r *DeleteTagRetentionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRetentionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagRetentionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRetentionRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagRetentionRuleResponseParams `json:"Response"`
}

func (r *DeleteTagRetentionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebhookTriggerRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type DeleteWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Namespace")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebhookTriggerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWebhookTriggerResponseParams `json:"Response"`
}

func (r *DeleteWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationTriggerLogPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 升序或降序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 按某列排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

type DescribeApplicationTriggerLogPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 升序或降序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 按某列排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeApplicationTriggerLogPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationTriggerLogPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationTriggerLogPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerLogPersonalResp struct {
	// 返回总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 触发日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogInfo []*TriggerLogResp `json:"LogInfo,omitempty" name:"LogInfo"`
}

// Predefined struct for user
type DescribeApplicationTriggerLogPersonalResponseParams struct {
	// 触发日志返回值
	Data *DescribeApplicationTriggerLogPersonalResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationTriggerLogPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationTriggerLogPersonalResponseParams `json:"Response"`
}

func (r *DescribeApplicationTriggerLogPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationTriggerLogPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationTriggerPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "TriggerName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationTriggerPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationTriggerPersonalResp struct {
	// 返回条目总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 触发器列表
	TriggerInfo []*TriggerResp `json:"TriggerInfo,omitempty" name:"TriggerInfo"`
}

// Predefined struct for user
type DescribeApplicationTriggerPersonalResponseParams struct {
	// 触发器列表返回值
	Data *DescribeApplicationTriggerPersonalResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationTriggerPersonalResponseParams `json:"Response"`
}

func (r *DescribeApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChartDownloadInfoRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Chart包的名称
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// Chart包的版本
	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

type DescribeChartDownloadInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Chart包的名称
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// Chart包的版本
	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

func (r *DescribeChartDownloadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChartDownloadInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "ChartName")
	delete(f, "ChartVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChartDownloadInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChartDownloadInfoResponseParams struct {
	// 用于下载的url的预签名地址
	PreSignedDownloadURL *string `json:"PreSignedDownloadURL,omitempty" name:"PreSignedDownloadURL"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeChartDownloadInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChartDownloadInfoResponseParams `json:"Response"`
}

func (r *DescribeChartDownloadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChartDownloadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalEndpointStatusRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DescribeExternalEndpointStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeExternalEndpointStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalEndpointStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExternalEndpointStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalEndpointStatusResponseParams struct {
	// 开启公网访问状态，开启中（Opening）、已开启（Opened）、关闭（Closed）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExternalEndpointStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExternalEndpointStatusResponseParams `json:"Response"`
}

func (r *DescribeExternalEndpointStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalEndpointStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFavorRepositoryPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset用于分页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeFavorRepositoryPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset用于分页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFavorRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFavorRepositoryPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFavorRepositoryPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFavorRepositoryPersonalResponseParams struct {
	// 个人收藏仓库列表返回信息
	Data *FavorResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFavorRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFavorRepositoryPersonalResponseParams `json:"Response"`
}

func (r *DescribeFavorRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFavorRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGCJobsRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DescribeGCJobsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeGCJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGCJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGCJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGCJobsResponseParams struct {
	// GC Job 列表
	Jobs []*GCJobInfo `json:"Jobs,omitempty" name:"Jobs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGCJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGCJobsResponseParams `json:"Response"`
}

func (r *DescribeGCJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGCJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAccelerateServiceRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DescribeImageAccelerateServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeImageAccelerateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAccelerateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAccelerateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAccelerateServiceResponseParams struct {
	// 镜像加速状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// CFS的VIP
	CFSVIP *string `json:"CFSVIP,omitempty" name:"CFSVIP"`

	// 是否开通
	IsEnable *bool `json:"IsEnable,omitempty" name:"IsEnable"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageAccelerateServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAccelerateServiceResponseParams `json:"Response"`
}

func (r *DescribeImageAccelerateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAccelerateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageFilterPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type DescribeImageFilterPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImageFilterPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageFilterPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageFilterPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageFilterPersonalResponseParams struct {
	// 返回tag镜像内容相同的tag列表
	Data *SameImagesResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageFilterPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageFilterPersonalResponseParams `json:"Response"`
}

func (r *DescribeImageFilterPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageFilterPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageLifecycleGlobalPersonalRequestParams struct {

}

type DescribeImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImageLifecycleGlobalPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageLifecycleGlobalPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageLifecycleGlobalPersonalResponseParams struct {
	// 全局自动删除策略信息
	Data *AutoDelStrategyInfoResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageLifecycleGlobalPersonalResponseParams `json:"Response"`
}

func (r *DescribeImageLifecycleGlobalPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageLifecyclePersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type DescribeImageLifecyclePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeImageLifecyclePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageLifecyclePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageLifecyclePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageLifecyclePersonalResponseParams struct {
	// 自动删除策略信息
	Data *AutoDelStrategyInfoResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageLifecyclePersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageLifecyclePersonalResponseParams `json:"Response"`
}

func (r *DescribeImageLifecyclePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageLifecyclePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageManifestsRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
}

type DescribeImageManifestsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 镜像版本
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
}

func (r *DescribeImageManifestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageManifestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageManifestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageManifestsResponseParams struct {
	// 镜像的Manifest信息
	Manifest *string `json:"Manifest,omitempty" name:"Manifest"`

	// 镜像的配置信息
	Config *string `json:"Config,omitempty" name:"Config"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImageManifestsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageManifestsResponseParams `json:"Response"`
}

func (r *DescribeImageManifestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageManifestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagePersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// tag名称，可根据输入搜索
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

type DescribeImagePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// tag名称，可根据输入搜索
	Tag *string `json:"Tag,omitempty" name:"Tag"`
}

func (r *DescribeImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagePersonalResponseParams struct {
	// 镜像tag信息
	Data *TagInfoResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagePersonalResponseParams `json:"Response"`
}

func (r *DescribeImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 指定镜像版本进行查找，当前为模糊搜索
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 每页个数，用于分页，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页数，默认值为1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定镜像 Digest 进行查找
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// 指定是否为精准匹配，true为精准匹配，不填为模糊匹配
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 指定镜像版本进行查找，当前为模糊搜索
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 每页个数，用于分页，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页数，默认值为1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定镜像 Digest 进行查找
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// 指定是否为精准匹配，true为精准匹配，不填为模糊匹配
	ExactMatch *bool `json:"ExactMatch,omitempty" name:"ExactMatch"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Digest")
	delete(f, "ExactMatch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// 容器镜像信息列表
	ImageInfoList []*TcrImageInfo `json:"ImageInfoList,omitempty" name:"ImageInfoList"`

	// 容器镜像总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DescribeImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImmutableTagRulesResponseParams struct {
	// 规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*ImmutableTagRule `json:"Rules,omitempty" name:"Rules"`

	// 未创建规则的命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyNs []*string `json:"EmptyNs,omitempty" name:"EmptyNs"`

	// 规则总量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImmutableTagRulesResponseParams `json:"Response"`
}

func (r *DescribeImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAllNamespacesRequestParams struct {
	// 每页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstanceAllNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// 每页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceAllNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAllNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAllNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAllNamespacesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceAllNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAllNamespacesResponseParams `json:"Response"`
}

func (r *DescribeInstanceAllNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAllNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCustomizedDomainRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceCustomizedDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceCustomizedDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCustomizedDomainResponseParams struct {
	// 域名信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainInfoList []*CustomizedDomainInfo `json:"DomainInfoList,omitempty" name:"DomainInfoList"`

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *DescribeInstanceCustomizedDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStatusRequestParams struct {
	// 实例ID的数组
	RegistryIds []*string `json:"RegistryIds,omitempty" name:"RegistryIds"`
}

type DescribeInstanceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID的数组
	RegistryIds []*string `json:"RegistryIds,omitempty" name:"RegistryIds"`
}

func (r *DescribeInstanceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStatusResponseParams struct {
	// 实例的状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryStatusSet []*RegistryStatus `json:"RegistryStatusSet,omitempty" name:"RegistryStatusSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStatusResponseParams `json:"Response"`
}

func (r *DescribeInstanceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTokenRequestParams struct {
	// 实例 ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTokenResponseParams struct {
	// 长期访问凭证总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 长期访问凭证列表
	Tokens []*TcrInstanceToken `json:"Tokens,omitempty" name:"Tokens"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTokenResponseParams `json:"Response"`
}

func (r *DescribeInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 实例ID列表(为空时，
	// 表示获取账号下所有实例)
	Registryids []*string `json:"Registryids,omitempty" name:"Registryids"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 获取所有地域的实例，默认为False
	AllRegion *bool `json:"AllRegion,omitempty" name:"AllRegion"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表(为空时，
	// 表示获取账号下所有实例)
	Registryids []*string `json:"Registryids,omitempty" name:"Registryids"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 获取所有地域的实例，默认为False
	AllRegion *bool `json:"AllRegion,omitempty" name:"AllRegion"`
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
	delete(f, "Registryids")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "AllRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 总实例个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Registries []*Registry `json:"Registries,omitempty" name:"Registries"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInternalEndpointDnsStatusRequestParams struct {
	// vpc列表
	VpcSet []*VpcAndDomainInfo `json:"VpcSet,omitempty" name:"VpcSet"`
}

type DescribeInternalEndpointDnsStatusRequest struct {
	*tchttp.BaseRequest
	
	// vpc列表
	VpcSet []*VpcAndDomainInfo `json:"VpcSet,omitempty" name:"VpcSet"`
}

func (r *DescribeInternalEndpointDnsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointDnsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternalEndpointDnsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternalEndpointDnsStatusResponseParams struct {
	// vpc私有域名解析状态列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcSet []*VpcPrivateDomainStatus `json:"VpcSet,omitempty" name:"VpcSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInternalEndpointDnsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternalEndpointDnsStatusResponseParams `json:"Response"`
}

func (r *DescribeInternalEndpointDnsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointDnsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternalEndpointsRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DescribeInternalEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeInternalEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternalEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternalEndpointsResponseParams struct {
	// 内网接入信息的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessVpcSet []*AccessVpc `json:"AccessVpcSet,omitempty" name:"AccessVpcSet"`

	// 内网接入总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInternalEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternalEndpointsResponseParams `json:"Response"`
}

func (r *DescribeInternalEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacePersonalRequestParams struct {
	// 命名空间，支持模糊查询
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeNamespacePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间，支持模糊查询
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNamespacePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespacePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacePersonalResponseParams struct {
	// 用户命名空间返回信息
	Data *NamespaceInfoResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNamespacePersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespacePersonalResponseParams `json:"Response"`
}

func (r *DescribeNamespacePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认查询所有命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 每页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页面偏移（第几页）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 列出所有命名空间
	All *bool `json:"All,omitempty" name:"All"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 仅查询启用了 KMS 镜像签名的空间
	KmsSignPolicy *bool `json:"KmsSignPolicy,omitempty" name:"KmsSignPolicy"`
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认查询所有命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 每页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页面偏移（第几页）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 列出所有命名空间
	All *bool `json:"All,omitempty" name:"All"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 仅查询启用了 KMS 镜像签名的空间
	KmsSignPolicy *bool `json:"KmsSignPolicy,omitempty" name:"KmsSignPolicy"`
}

func (r *DescribeNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "All")
	delete(f, "Filters")
	delete(f, "KmsSignPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesResponseParams struct {
	// 命名空间列表信息
	NamespaceList []*TcrNamespaceInfo `json:"NamespaceList,omitempty" name:"NamespaceList"`

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespacesResponseParams `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 返回的总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 地域信息列表
	Regions []*Region `json:"Regions,omitempty" name:"Regions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceCreateTasksRequestParams struct {
	// 同步实例Id，见实例返回列表中的同步实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 同步实例的地域ID，见实例返回列表中地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`
}

type DescribeReplicationInstanceCreateTasksRequest struct {
	*tchttp.BaseRequest
	
	// 同步实例Id，见实例返回列表中的同步实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 同步实例的地域ID，见实例返回列表中地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`
}

func (r *DescribeReplicationInstanceCreateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceCreateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstanceCreateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceCreateTasksResponseParams struct {
	// 任务详情
	TaskDetail []*TaskDetail `json:"TaskDetail,omitempty" name:"TaskDetail"`

	// 整体任务状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReplicationInstanceCreateTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationInstanceCreateTasksResponseParams `json:"Response"`
}

func (r *DescribeReplicationInstanceCreateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceCreateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceSyncStatusRequestParams struct {
	// 主实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 复制实例Id
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 复制实例的地域Id
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// 是否显示同步日志
	ShowReplicationLog *bool `json:"ShowReplicationLog,omitempty" name:"ShowReplicationLog"`

	// 日志页号, 默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认5，最大为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeReplicationInstanceSyncStatusRequest struct {
	*tchttp.BaseRequest
	
	// 主实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 复制实例Id
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 复制实例的地域Id
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// 是否显示同步日志
	ShowReplicationLog *bool `json:"ShowReplicationLog,omitempty" name:"ShowReplicationLog"`

	// 日志页号, 默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认5，最大为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeReplicationInstanceSyncStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceSyncStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	delete(f, "ShowReplicationLog")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstanceSyncStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceSyncStatusResponseParams struct {
	// 同步状态
	ReplicationStatus *string `json:"ReplicationStatus,omitempty" name:"ReplicationStatus"`

	// 同步完成时间
	ReplicationTime *string `json:"ReplicationTime,omitempty" name:"ReplicationTime"`

	// 同步日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicationLog *ReplicationLog `json:"ReplicationLog,omitempty" name:"ReplicationLog"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReplicationInstanceSyncStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationInstanceSyncStatusResponseParams `json:"Response"`
}

func (r *DescribeReplicationInstanceSyncStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstancesRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeReplicationInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeReplicationInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstancesResponseParams struct {
	// 总实例个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 同步实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicationRegistries []*ReplicationRegistry `json:"ReplicationRegistries,omitempty" name:"ReplicationRegistries"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReplicationInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationInstancesResponseParams `json:"Response"`
}

func (r *DescribeReplicationInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认为查询所有命名空间下镜像仓库
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 指定镜像仓库，不填写默认查询指定命名空间下所有镜像仓库
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 页数，用于分页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页个数，用于分页
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 基于字段排序，支持的值有-creation_time,-name, -update_time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
}

type DescribeRepositoriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 指定命名空间，不填写默认为查询所有命名空间下镜像仓库
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 指定镜像仓库，不填写默认查询指定命名空间下所有镜像仓库
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 页数，用于分页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页个数，用于分页
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 基于字段排序，支持的值有-creation_time,-name, -update_time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeRepositoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesResponseParams struct {
	// 仓库信息列表
	RepositoryList []*TcrRepositoryInfo `json:"RepositoryList,omitempty" name:"RepositoryList"`

	// 总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRepositoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoriesResponseParams `json:"Response"`
}

func (r *DescribeRepositoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryFilterPersonalRequestParams struct {
	// 搜索镜像名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20，最大100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件：1表示public，0表示private
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DescribeRepositoryFilterPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 搜索镜像名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20，最大100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选条件：1表示public，0表示private
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeRepositoryFilterPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryFilterPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Public")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoryFilterPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryFilterPersonalResponseParams struct {
	// 仓库信息
	Data *SearchUserRepositoryResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRepositoryFilterPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoryFilterPersonalResponseParams `json:"Response"`
}

func (r *DescribeRepositoryFilterPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryFilterPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryOwnerPersonalRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type DescribeRepositoryOwnerPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回最大数量，默认 20, 最大值 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeRepositoryOwnerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryOwnerPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RepoName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoryOwnerPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryOwnerPersonalResponseParams struct {
	// 仓库信息
	Data *RepoInfoResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRepositoryOwnerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoryOwnerPersonalResponseParams `json:"Response"`
}

func (r *DescribeRepositoryOwnerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryOwnerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryPersonalRequestParams struct {
	// 仓库名字
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type DescribeRepositoryPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名字
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DescribeRepositoryPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoryPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoryPersonalResponseParams struct {
	// 仓库信息
	Data *RepositoryInfoResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRepositoryPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoryPersonalResponseParams `json:"Response"`
}

func (r *DescribeRepositoryPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoryPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPoliciesRequestParams struct {
	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type DescribeSecurityPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *DescribeSecurityPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPoliciesResponseParams struct {
	// 实例安全策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityPolicySet []*SecurityPolicy `json:"SecurityPolicySet,omitempty" name:"SecurityPolicySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTagRetentionExecutionRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTagRetentionExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagRetentionExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionResponseParams struct {
	// 版本保留执行记录列表
	RetentionExecutionList []*RetentionExecution `json:"RetentionExecutionList,omitempty" name:"RetentionExecutionList"`

	// 版本保留执行记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagRetentionExecutionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionExecutionResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionTaskRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 规则执行Id
	ExecutionId *int64 `json:"ExecutionId,omitempty" name:"ExecutionId"`

	// 分页Page
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTagRetentionExecutionTaskRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 规则执行Id
	ExecutionId *int64 `json:"ExecutionId,omitempty" name:"ExecutionId"`

	// 分页Page
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTagRetentionExecutionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	delete(f, "ExecutionId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagRetentionExecutionTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionTaskResponseParams struct {
	// 版本保留执行任务列表
	RetentionTaskList []*RetentionTask `json:"RetentionTaskList,omitempty" name:"RetentionTaskList"`

	// 版本保留执行任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagRetentionExecutionTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionExecutionTaskResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionExecutionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionRulesRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTagRetentionRulesRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 分页PageSize
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页Page
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTagRetentionRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagRetentionRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionRulesResponseParams struct {
	// 版本保留策略列表
	RetentionPolicyList []*RetentionPolicy `json:"RetentionPolicyList,omitempty" name:"RetentionPolicyList"`

	// 版本保留策略总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTagRetentionRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionRulesResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotaPersonalRequestParams struct {

}

type DescribeUserQuotaPersonalRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserQuotaPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotaPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserQuotaPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotaPersonalResponseParams struct {
	// 配额返回信息
	Data *RespLimit `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserQuotaPersonalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserQuotaPersonalResponseParams `json:"Response"`
}

func (r *DescribeUserQuotaPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotaPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerLogRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeWebhookTriggerLogRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 触发器 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeWebhookTriggerLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Namespace")
	delete(f, "Id")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebhookTriggerLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerLogResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 日志列表
	Logs []*WebhookTriggerLog `json:"Logs,omitempty" name:"Logs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebhookTriggerLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebhookTriggerLogResponseParams `json:"Response"`
}

func (r *DescribeWebhookTriggerLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DescribeWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 分页单页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerResponseParams struct {
	// 触发器总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 触发器列表
	Triggers []*WebhookTrigger `json:"Triggers,omitempty" name:"Triggers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebhookTriggerResponseParams `json:"Response"`
}

func (r *DescribeWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadHelmChartRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Helm chart名称
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// Helm chart版本
	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

type DownloadHelmChartRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Helm chart名称
	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`

	// Helm chart版本
	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

func (r *DownloadHelmChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadHelmChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "ChartName")
	delete(f, "ChartVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadHelmChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadHelmChartResponseParams struct {
	// 临时token
	TmpToken *string `json:"TmpToken,omitempty" name:"TmpToken"`

	// 临时的secretId
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时的secretKey
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 存储桶信息
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 实例ID
	Region *string `json:"Region,omitempty" name:"Region"`

	// chart信息
	Path *string `json:"Path,omitempty" name:"Path"`

	// 开始时间时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// token过期时间时间戳
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadHelmChartResponse struct {
	*tchttp.BaseResponse
	Response *DownloadHelmChartResponseParams `json:"Response"`
}

func (r *DownloadHelmChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadHelmChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DupImageTagResp struct {
	// 镜像Digest值
	Digest *string `json:"Digest,omitempty" name:"Digest"`
}

// Predefined struct for user
type DuplicateImagePersonalRequestParams struct {
	// 源镜像名称，不包含domain。例如： tencentyun/foo:v1
	SrcImage *string `json:"SrcImage,omitempty" name:"SrcImage"`

	// 目的镜像名称，不包含domain。例如： tencentyun/foo:latest
	DestImage *string `json:"DestImage,omitempty" name:"DestImage"`
}

type DuplicateImagePersonalRequest struct {
	*tchttp.BaseRequest
	
	// 源镜像名称，不包含domain。例如： tencentyun/foo:v1
	SrcImage *string `json:"SrcImage,omitempty" name:"SrcImage"`

	// 目的镜像名称，不包含domain。例如： tencentyun/foo:latest
	DestImage *string `json:"DestImage,omitempty" name:"DestImage"`
}

func (r *DuplicateImagePersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DuplicateImagePersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcImage")
	delete(f, "DestImage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DuplicateImagePersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DuplicateImagePersonalResponseParams struct {
	// 复制镜像返回值
	Data *DupImageTagResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DuplicateImagePersonalResponse struct {
	*tchttp.BaseResponse
	Response *DuplicateImagePersonalResponseParams `json:"Response"`
}

func (r *DuplicateImagePersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DuplicateImagePersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FavorResp struct {
	// 收藏仓库的总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoInfo []*Favors `json:"RepoInfo,omitempty" name:"RepoInfo"`
}

type Favors struct {
	// 仓库名字
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// Pull总共的次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 仓库收藏次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 仓库是否公开
	// 注意：此字段可能返回 null，表示取不到有效值。
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 是否为官方所有
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`

	// 仓库Tag的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// Logo
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域的Id
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type GCJobInfo struct {
	// 作业 ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 作业状态
	JobStatus *string `json:"JobStatus,omitempty" name:"JobStatus"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 调度信息
	Schedule *Schedule `json:"Schedule,omitempty" name:"Schedule"`
}

type Header struct {
	// Header Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// Header Values
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ImmutableTagRule struct {
	// 仓库匹配规则
	RepositoryPattern *string `json:"RepositoryPattern,omitempty" name:"RepositoryPattern"`

	// Tag 匹配规则
	TagPattern *string `json:"TagPattern,omitempty" name:"TagPattern"`

	// repoMatches或repoExcludes
	RepositoryDecoration *string `json:"RepositoryDecoration,omitempty" name:"RepositoryDecoration"`

	// matches或excludes
	TagDecoration *string `json:"TagDecoration,omitempty" name:"TagDecoration"`

	// 禁用规则
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 命名空间
	NsName *string `json:"NsName,omitempty" name:"NsName"`
}

type KeyValueString struct {
	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type Limit struct {
	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 配额的类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置的值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type ManageExternalEndpointRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 操作（Create/Delete）
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type ManageExternalEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 操作（Create/Delete）
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *ManageExternalEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageExternalEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageExternalEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageExternalEndpointResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManageExternalEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ManageExternalEndpointResponseParams `json:"Response"`
}

func (r *ManageExternalEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageExternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageImageLifecycleGlobalPersonalRequestParams struct {
	// global_keep_last_days:全局保留最近几天的数据;global_keep_last_nums:全局保留最近多少个
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Val *int64 `json:"Val,omitempty" name:"Val"`
}

type ManageImageLifecycleGlobalPersonalRequest struct {
	*tchttp.BaseRequest
	
	// global_keep_last_days:全局保留最近几天的数据;global_keep_last_nums:全局保留最近多少个
	Type *string `json:"Type,omitempty" name:"Type"`

	// 策略值
	Val *int64 `json:"Val,omitempty" name:"Val"`
}

func (r *ManageImageLifecycleGlobalPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageImageLifecycleGlobalPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Val")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageImageLifecycleGlobalPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageImageLifecycleGlobalPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManageImageLifecycleGlobalPersonalResponse struct {
	*tchttp.BaseResponse
	Response *ManageImageLifecycleGlobalPersonalResponseParams `json:"Response"`
}

func (r *ManageImageLifecycleGlobalPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageImageLifecycleGlobalPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageInternalEndpointRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Create/Delete
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 需要接入的用户vpcid
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 需要接入的用户子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 请求的地域名称，用于实例复制地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type ManageInternalEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// Create/Delete
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 需要接入的用户vpcid
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 需要接入的用户子网id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 请求的地域ID，用于实例复制地域
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 请求的地域名称，用于实例复制地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *ManageInternalEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageInternalEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Operation")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "RegionId")
	delete(f, "RegionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageInternalEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageInternalEndpointResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManageInternalEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ManageInternalEndpointResponseParams `json:"Response"`
}

func (r *ManageInternalEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageInternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageReplicationRequestParams struct {
	// 复制源实例ID
	SourceRegistryId *string `json:"SourceRegistryId,omitempty" name:"SourceRegistryId"`

	// 复制目标实例ID
	DestinationRegistryId *string `json:"DestinationRegistryId,omitempty" name:"DestinationRegistryId"`

	// 同步规则
	Rule *ReplicationRule `json:"Rule,omitempty" name:"Rule"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 目标实例的地域ID，如广州是1
	DestinationRegionId *uint64 `json:"DestinationRegionId,omitempty" name:"DestinationRegionId"`

	// 开启跨主账号实例同步配置项
	PeerReplicationOption *PeerReplicationOption `json:"PeerReplicationOption,omitempty" name:"PeerReplicationOption"`
}

type ManageReplicationRequest struct {
	*tchttp.BaseRequest
	
	// 复制源实例ID
	SourceRegistryId *string `json:"SourceRegistryId,omitempty" name:"SourceRegistryId"`

	// 复制目标实例ID
	DestinationRegistryId *string `json:"DestinationRegistryId,omitempty" name:"DestinationRegistryId"`

	// 同步规则
	Rule *ReplicationRule `json:"Rule,omitempty" name:"Rule"`

	// 规则描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 目标实例的地域ID，如广州是1
	DestinationRegionId *uint64 `json:"DestinationRegionId,omitempty" name:"DestinationRegionId"`

	// 开启跨主账号实例同步配置项
	PeerReplicationOption *PeerReplicationOption `json:"PeerReplicationOption,omitempty" name:"PeerReplicationOption"`
}

func (r *ManageReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceRegistryId")
	delete(f, "DestinationRegistryId")
	delete(f, "Rule")
	delete(f, "Description")
	delete(f, "DestinationRegionId")
	delete(f, "PeerReplicationOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageReplicationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ManageReplicationResponse struct {
	*tchttp.BaseResponse
	Response *ManageReplicationResponseParams `json:"Response"`
}

func (r *ManageReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationTriggerPersonalRequestParams struct {
	// 触发器关联的镜像仓库，library/test格式
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称，必填参数
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 触发方式对应的表达式
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`

	// 应用所在TKE集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用所在TKE集群命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 应用所在TKE集群工作负载名称
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 应用所在TKE集群工作负载下容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 应用所在TKE集群地域数字ID，如1（广州）、16（成都）
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`

	// 新触发器名称
	NewTriggerName *string `json:"NewTriggerName,omitempty" name:"NewTriggerName"`
}

type ModifyApplicationTriggerPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 触发器关联的镜像仓库，library/test格式
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 触发器名称，必填参数
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式，"all"全部触发，"taglist"指定tag触发，"regex"正则触发
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 触发方式对应的表达式
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`

	// 应用所在TKE集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用所在TKE集群命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 应用所在TKE集群工作负载类型,支持Deployment、StatefulSet、DaemonSet、CronJob、Job。
	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`

	// 应用所在TKE集群工作负载名称
	WorkloadName *string `json:"WorkloadName,omitempty" name:"WorkloadName"`

	// 应用所在TKE集群工作负载下容器名称
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// 应用所在TKE集群地域数字ID，如1（广州）、16（成都）
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`

	// 新触发器名称
	NewTriggerName *string `json:"NewTriggerName,omitempty" name:"NewTriggerName"`
}

func (r *ModifyApplicationTriggerPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationTriggerPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "TriggerName")
	delete(f, "InvokeMethod")
	delete(f, "InvokeExpr")
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "WorkloadType")
	delete(f, "WorkloadName")
	delete(f, "ContainerName")
	delete(f, "ClusterRegion")
	delete(f, "NewTriggerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationTriggerPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationTriggerPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyApplicationTriggerPersonalResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationTriggerPersonalResponseParams `json:"Response"`
}

func (r *ModifyApplicationTriggerPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationTriggerPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImmutableTagRulesRequestParams struct {
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitempty" name:"Rule"`
}

type ModifyImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 规则 Id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 规则
	Rule *ImmutableTagRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RuleId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImmutableTagRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImmutableTagRulesResponseParams `json:"Response"`
}

func (r *ModifyImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例的规格
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例的规格
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`
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
	delete(f, "RegistryId")
	delete(f, "RegistryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyInstanceTokenRequestParams struct {
	// 实例长期访问凭证 ID
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`

	// 实例 ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 启用或禁用实例长期访问凭证
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 访问凭证描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 1为修改描述 2为操作启动禁用，默认值为2
	ModifyFlag *int64 `json:"ModifyFlag,omitempty" name:"ModifyFlag"`
}

type ModifyInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// 实例长期访问凭证 ID
	TokenId *string `json:"TokenId,omitempty" name:"TokenId"`

	// 实例 ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 启用或禁用实例长期访问凭证
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 访问凭证描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 1为修改描述 2为操作启动禁用，默认值为2
	ModifyFlag *int64 `json:"ModifyFlag,omitempty" name:"ModifyFlag"`
}

func (r *ModifyInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TokenId")
	delete(f, "RegistryId")
	delete(f, "Enable")
	delete(f, "Desc")
	delete(f, "ModifyFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTokenResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceTokenResponseParams `json:"Response"`
}

func (r *ModifyInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 访问级别，True为公开，False为私有
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 扫描级别，True为自动，False为手动
	IsAutoScan *bool `json:"IsAutoScan,omitempty" name:"IsAutoScan"`

	// 阻断开关，True为开放，False为关闭
	IsPreventVUL *bool `json:"IsPreventVUL,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持 low、medium、high
	Severity *string `json:"Severity,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitempty" name:"CVEWhitelistItems"`
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 访问级别，True为公开，False为私有
	IsPublic *bool `json:"IsPublic,omitempty" name:"IsPublic"`

	// 扫描级别，True为自动，False为手动
	IsAutoScan *bool `json:"IsAutoScan,omitempty" name:"IsAutoScan"`

	// 阻断开关，True为开放，False为关闭
	IsPreventVUL *bool `json:"IsPreventVUL,omitempty" name:"IsPreventVUL"`

	// 阻断漏洞等级，目前仅支持 low、medium、high
	Severity *string `json:"Severity,omitempty" name:"Severity"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitempty" name:"CVEWhitelistItems"`
}

func (r *ModifyNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "IsPublic")
	delete(f, "IsAutoScan")
	delete(f, "IsPreventVUL")
	delete(f, "Severity")
	delete(f, "CVEWhitelistItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNamespaceResponseParams `json:"Response"`
}

func (r *ModifyNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryAccessPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 默认值为0, 1公共，0私有
	Public *int64 `json:"Public,omitempty" name:"Public"`
}

type ModifyRepositoryAccessPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 默认值为0, 1公共，0私有
	Public *int64 `json:"Public,omitempty" name:"Public"`
}

func (r *ModifyRepositoryAccessPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryAccessPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Public")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRepositoryAccessPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryAccessPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRepositoryAccessPersonalResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRepositoryAccessPersonalResponseParams `json:"Response"`
}

func (r *ModifyRepositoryAccessPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryAccessPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryInfoPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyRepositoryInfoPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyRepositoryInfoPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryInfoPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRepositoryInfoPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryInfoPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRepositoryInfoPersonalResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRepositoryInfoPersonalResponseParams `json:"Response"`
}

func (r *ModifyRepositoryInfoPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryInfoPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryRequestParams struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifyRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 镜像仓库名称
	RepositoryName *string `json:"RepositoryName,omitempty" name:"RepositoryName"`

	// 仓库简短描述
	BriefDescription *string `json:"BriefDescription,omitempty" name:"BriefDescription"`

	// 仓库详细描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "BriefDescription")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRepositoryResponseParams `json:"Response"`
}

func (r *ModifyRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyRequestParams struct {
	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// PolicyId
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 192.168.0.0/24 白名单Ip
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// PolicyId
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 192.168.0.0/24 白名单Ip
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifySecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "PolicyIndex")
	delete(f, "CidrBlock")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyResponseParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityPolicyResponseParams `json:"Response"`
}

func (r *ModifySecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagRetentionRuleRequestParams struct {
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的Id，必须填写原有的命名空间id
	NamespaceId *int64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitempty" name:"RetentionRule"`

	// 执行周期，必须填写为原来的设置
	CronSetting *string `json:"CronSetting,omitempty" name:"CronSetting"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 是否禁用规则
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`
}

type ModifyTagRetentionRuleRequest struct {
	*tchttp.BaseRequest
	
	// 主实例iD
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 命名空间的Id，必须填写原有的命名空间id
	NamespaceId *int64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 保留策略
	RetentionRule *RetentionRule `json:"RetentionRule,omitempty" name:"RetentionRule"`

	// 执行周期，必须填写为原来的设置
	CronSetting *string `json:"CronSetting,omitempty" name:"CronSetting"`

	// 规则Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 是否禁用规则
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`
}

func (r *ModifyTagRetentionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagRetentionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceId")
	delete(f, "RetentionRule")
	delete(f, "CronSetting")
	delete(f, "RetentionId")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTagRetentionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagRetentionRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTagRetentionRuleResponseParams `json:"Response"`
}

func (r *ModifyTagRetentionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPasswordPersonalRequestParams struct {
	// 更新后的密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ModifyUserPasswordPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 更新后的密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModifyUserPasswordPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPasswordPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserPasswordPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPasswordPersonalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserPasswordPersonalResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserPasswordPersonalResponseParams `json:"Response"`
}

func (r *ModifyUserPasswordPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPasswordPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebhookTriggerRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type ModifyWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 触发器参数
	Trigger *WebhookTrigger `json:"Trigger,omitempty" name:"Trigger"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ModifyWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Trigger")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebhookTriggerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebhookTriggerResponseParams `json:"Response"`
}

func (r *ModifyWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NamespaceInfo struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 命名空间下仓库数量
	RepoCount *int64 `json:"RepoCount,omitempty" name:"RepoCount"`
}

type NamespaceInfoResp struct {
	// 命名空间数量
	NamespaceCount *int64 `json:"NamespaceCount,omitempty" name:"NamespaceCount"`

	// 命名空间信息
	NamespaceInfo []*NamespaceInfo `json:"NamespaceInfo,omitempty" name:"NamespaceInfo"`
}

type NamespaceIsExistsResp struct {
	// 命名空间是否存在
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`

	// 是否为保留命名空间
	IsPreserved *bool `json:"IsPreserved,omitempty" name:"IsPreserved"`
}

type PeerReplicationOption struct {
	// 待同步实例的uin
	PeerRegistryUin *string `json:"PeerRegistryUin,omitempty" name:"PeerRegistryUin"`

	// 待同步实例的访问永久Token
	PeerRegistryToken *string `json:"PeerRegistryToken,omitempty" name:"PeerRegistryToken"`

	// 是否开启跨主账号实例同步
	EnablePeerReplication *bool `json:"EnablePeerReplication,omitempty" name:"EnablePeerReplication"`
}

type Region struct {
	// gz
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 1
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// ap-guangzhou
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// alluser
	Status *string `json:"Status,omitempty" name:"Status"`

	// remark
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// id
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type Registry struct {
	// 实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例名称
	RegistryName *string `json:"RegistryName,omitempty" name:"RegistryName"`

	// 实例规格
	RegistryType *string `json:"RegistryType,omitempty" name:"RegistryType"`

	// 实例状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例的公共访问地址
	PublicDomain *string `json:"PublicDomain,omitempty" name:"PublicDomain"`

	// 实例创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域Id
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 是否支持匿名
	EnableAnonymous *bool `json:"EnableAnonymous,omitempty" name:"EnableAnonymous"`

	// Token有效时间
	TokenValidTime *uint64 `json:"TokenValidTime,omitempty" name:"TokenValidTime"`

	// 实例内部访问地址
	InternalEndpoint *string `json:"InternalEndpoint,omitempty" name:"InternalEndpoint"`

	// 实例云标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 实例过期时间（预付费）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredAt *string `json:"ExpiredAt,omitempty" name:"ExpiredAt"`

	// 实例付费类型，0表示后付费，1表示预付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMod *int64 `json:"PayMod,omitempty" name:"PayMod"`

	// 预付费续费标识，0表示手动续费，1表示自动续费，2不续费并且不通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type RegistryChargePrepaid struct {
	// 购买实例的时长，单位：月
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识，0：手动续费，1：自动续费，2：不续费并且不通知
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type RegistryCondition struct {
	// 实例创建过程类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 实例创建过程状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 转换到该过程的简明原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type RegistryStatus struct {
	// 实例的Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 实例的状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 附加状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*RegistryCondition `json:"Conditions,omitempty" name:"Conditions"`
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 预付费自动续费标识和购买时长,0：手动续费，1：自动续费，2：不续费并且不通知;单位为月
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitempty" name:"RegistryChargePrepaid"`

	// 0 续费， 1按量转包年包月
	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 预付费自动续费标识和购买时长,0：手动续费，1：自动续费，2：不续费并且不通知;单位为月
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitempty" name:"RegistryChargePrepaid"`

	// 0 续费， 1按量转包年包月
	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RegistryChargePrepaid")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// 企业版实例Id
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstanceResponseParams `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicationFilter struct {
	// 类型（name、tag和resource）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 默认为空
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ReplicationLog struct {
	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 源资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 目的资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destination *string `json:"Destination,omitempty" name:"Destination"`

	// 同步状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ReplicationRegistry struct {
	// 主实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 复制实例ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitempty" name:"ReplicationRegistryId"`

	// 复制实例的地域ID
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitempty" name:"ReplicationRegionId"`

	// 复制实例的地域名称
	ReplicationRegionName *string `json:"ReplicationRegionName,omitempty" name:"ReplicationRegionName"`

	// 复制实例的状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type ReplicationRule struct {
	// 同步规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目标命名空间
	DestNamespace *string `json:"DestNamespace,omitempty" name:"DestNamespace"`

	// 是否覆盖
	Override *bool `json:"Override,omitempty" name:"Override"`

	// 同步过滤条件
	Filters []*ReplicationFilter `json:"Filters,omitempty" name:"Filters"`
}

type RepoInfo struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// Tag数量
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// 是否为公开
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 是否为用户收藏
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

	// 是否为腾讯云官方仓库
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`

	// 被收藏的个数
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 拉取的数量
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 仓库创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 仓库更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type RepoInfoResp struct {
	// 仓库总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库信息列表
	RepoInfo []*RepoInfo `json:"RepoInfo,omitempty" name:"RepoInfo"`

	// Server信息
	Server *string `json:"Server,omitempty" name:"Server"`
}

type RepoIsExistResp struct {
	// 仓库是否存在
	IsExist *bool `json:"IsExist,omitempty" name:"IsExist"`
}

type RepositoryInfoResp struct {
	// 镜像仓库名字
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像仓库类型
	RepoType *string `json:"RepoType,omitempty" name:"RepoType"`

	// 镜像仓库服务地址
	Server *string `json:"Server,omitempty" name:"Server"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 镜像仓库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否为公有镜像
	Public *int64 `json:"Public,omitempty" name:"Public"`

	// 下载次数
	PullCount *int64 `json:"PullCount,omitempty" name:"PullCount"`

	// 收藏次数
	FavorCount *int64 `json:"FavorCount,omitempty" name:"FavorCount"`

	// 是否为用户收藏
	IsUserFavor *bool `json:"IsUserFavor,omitempty" name:"IsUserFavor"`

	// 是否为腾讯云官方镜像
	IsQcloudOfficial *bool `json:"IsQcloudOfficial,omitempty" name:"IsQcloudOfficial"`
}

type RespLimit struct {
	// 配额信息
	LimitInfo []*Limit `json:"LimitInfo,omitempty" name:"LimitInfo"`
}

type RetentionExecution struct {
	// 执行Id
	ExecutionId *int64 `json:"ExecutionId,omitempty" name:"ExecutionId"`

	// 所属规则id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 执行的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 执行的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 执行的状态，Failed, Succeed, Stopped, InProgress
	Status *string `json:"Status,omitempty" name:"Status"`
}

type RetentionPolicy struct {
	// 版本保留策略Id
	RetentionId *int64 `json:"RetentionId,omitempty" name:"RetentionId"`

	// 命名空间的名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 规则列表
	RetentionRuleList []*RetentionRule `json:"RetentionRuleList,omitempty" name:"RetentionRuleList"`

	// 定期执行方式
	CronSetting *string `json:"CronSetting,omitempty" name:"CronSetting"`

	// 是否启用规则
	Disabled *bool `json:"Disabled,omitempty" name:"Disabled"`

	// 基于当前时间根据cronSetting后下一次任务要执行的时间，仅做参考使用
	NextExecutionTime *string `json:"NextExecutionTime,omitempty" name:"NextExecutionTime"`
}

type RetentionRule struct {
	// 支持的策略，可选值为latestPushedK（保留最新推送多少个版本）nDaysSinceLastPush（保留近天内推送）
	Key *string `json:"Key,omitempty" name:"Key"`

	// 规则设置下的对应值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type RetentionTask struct {
	// 任务Id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 所属的规则执行Id
	ExecutionId *int64 `json:"ExecutionId,omitempty" name:"ExecutionId"`

	// 任务开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务的执行状态，Failed, Succeed, Stopped, InProgress
	Status *string `json:"Status,omitempty" name:"Status"`

	// 总tag数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 保留tag数
	Retained *int64 `json:"Retained,omitempty" name:"Retained"`

	// 应用的仓库
	Repository *string `json:"Repository,omitempty" name:"Repository"`
}

type SameImagesResp struct {
	// tag列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SameImages []*string `json:"SameImages,omitempty" name:"SameImages"`
}

type Schedule struct {
	// 类型：Hourly, Daily, Weekly, Custom, Manual, Dryrun, None
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SearchUserRepositoryResp struct {
	// 总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库列表
	RepoInfo []*RepoInfo `json:"RepoInfo,omitempty" name:"RepoInfo"`

	// Server
	Server *string `json:"Server,omitempty" name:"Server"`

	// PrivilegeFiltered
	PrivilegeFiltered *bool `json:"PrivilegeFiltered,omitempty" name:"PrivilegeFiltered"`
}

type SecurityPolicy struct {
	// 策略索引
	PolicyIndex *int64 `json:"PolicyIndex,omitempty" name:"PolicyIndex"`

	// 备注
	Description *string `json:"Description,omitempty" name:"Description"`

	// 运行访问的公网IP地址端
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 安全策略的版本
	PolicyVersion *string `json:"PolicyVersion,omitempty" name:"PolicyVersion"`
}

type Tag struct {
	// 云标签的key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 云标签的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TagInfo struct {
	// Tag名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 制品的 ID
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// docker image 可以看到的id
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 大小
	Size *string `json:"Size,omitempty" name:"Size"`

	// 制品的创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 制品创建至今时间长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationDays *string `json:"DurationDays,omitempty" name:"DurationDays"`

	// 标注的制品作者
	Author *string `json:"Author,omitempty" name:"Author"`

	// 标注的制品平台
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// 创建制品的 Docker 版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// 标注的制品操作系统
	OS *string `json:"OS,omitempty" name:"OS"`

	// 制品大小
	SizeByte *int64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// 序号
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 数据更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 制品更新时间
	PushTime *string `json:"PushTime,omitempty" name:"PushTime"`

	// 制品类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

type TagInfoResp struct {
	// Tag的总数
	TagCount *int64 `json:"TagCount,omitempty" name:"TagCount"`

	// TagInfo列表
	TagInfo []*TagInfo `json:"TagInfo,omitempty" name:"TagInfo"`

	// Server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type TagSpecification struct {
	// 默认值为instance
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 云标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type TaskDetail struct {
	// 任务
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务UUID
	TaskUUID *string `json:"TaskUUID,omitempty" name:"TaskUUID"`

	// 任务状态
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 任务的状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskMessage *string `json:"TaskMessage,omitempty" name:"TaskMessage"`

	// 任务开始时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedTime *string `json:"FinishedTime,omitempty" name:"FinishedTime"`
}

type TcrImageInfo struct {
	// 哈希值
	Digest *string `json:"Digest,omitempty" name:"Digest"`

	// 镜像体积（单位：字节）
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// Tag名称
	ImageVersion *string `json:"ImageVersion,omitempty" name:"ImageVersion"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 制品类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// KMS 签名信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KmsSignature *string `json:"KmsSignature,omitempty" name:"KmsSignature"`
}

type TcrInstanceToken struct {
	// 令牌ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 令牌描述
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 令牌所属实例ID
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`

	// 令牌启用状态
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 令牌创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 令牌过期时间戳
	ExpiredAt *int64 `json:"ExpiredAt,omitempty" name:"ExpiredAt"`
}

type TcrNamespaceInfo struct {
	// 命名空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 访问级别
	Public *bool `json:"Public,omitempty" name:"Public"`

	// 命名空间的Id
	NamespaceId *int64 `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 实例云标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification *TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 命名空间元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata []*KeyValueString `json:"Metadata,omitempty" name:"Metadata"`

	// 漏洞白名单列表
	CVEWhitelistItems []*CVEWhitelistItem `json:"CVEWhitelistItems,omitempty" name:"CVEWhitelistItems"`

	// 扫描级别，true为自动，false为手动
	AutoScan *bool `json:"AutoScan,omitempty" name:"AutoScan"`

	// 安全阻断级别，true为开启，false为关闭
	PreventVUL *bool `json:"PreventVUL,omitempty" name:"PreventVUL"`

	// 阻断漏洞等级，目前仅支持low、medium、high, 为""时表示没有设置
	Severity *string `json:"Severity,omitempty" name:"Severity"`
}

type TcrRepositoryInfo struct {
	// 仓库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 创建时间，格式"2006-01-02 15:04:05.999999999 -0700 MST"
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 是否公开
	Public *bool `json:"Public,omitempty" name:"Public"`

	// 仓库详细描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 简单描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefDescription *string `json:"BriefDescription,omitempty" name:"BriefDescription"`

	// 更新时间，格式"2006-01-02 15:04:05.999999999 -0700 MST"
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TriggerInvokeCondition struct {
	// 触发方式
	InvokeMethod *string `json:"InvokeMethod,omitempty" name:"InvokeMethod"`

	// 触发表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeExpr *string `json:"InvokeExpr,omitempty" name:"InvokeExpr"`
}

type TriggerInvokePara struct {
	// AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// TKE集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// TKE集群命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// TKE集群工作负载名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// TKE集群工作负载中容器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`

	// TKE集群地域数字ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterRegion *int64 `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

type TriggerInvokeResult struct {
	// 请求TKE返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *int64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 请求TKE返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`
}

type TriggerLogResp struct {
	// 仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// Tag名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 触发器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeSource *string `json:"InvokeSource,omitempty" name:"InvokeSource"`

	// 触发动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeAction *string `json:"InvokeAction,omitempty" name:"InvokeAction"`

	// 触发时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeTime *string `json:"InvokeTime,omitempty" name:"InvokeTime"`

	// 触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeCondition *TriggerInvokeCondition `json:"InvokeCondition,omitempty" name:"InvokeCondition"`

	// 触发参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokePara *TriggerInvokePara `json:"InvokePara,omitempty" name:"InvokePara"`

	// 触发结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeResult *TriggerInvokeResult `json:"InvokeResult,omitempty" name:"InvokeResult"`
}

type TriggerResp struct {
	// 触发器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerName *string `json:"TriggerName,omitempty" name:"TriggerName"`

	// 触发来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeSource *string `json:"InvokeSource,omitempty" name:"InvokeSource"`

	// 触发动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeAction *string `json:"InvokeAction,omitempty" name:"InvokeAction"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeCondition *TriggerInvokeCondition `json:"InvokeCondition,omitempty" name:"InvokeCondition"`

	// 触发器参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokePara *TriggerInvokePara `json:"InvokePara,omitempty" name:"InvokePara"`
}

// Predefined struct for user
type ValidateNamespaceExistPersonalRequestParams struct {
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type ValidateNamespaceExistPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *ValidateNamespaceExistPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ValidateNamespaceExistPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ValidateNamespaceExistPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ValidateNamespaceExistPersonalResponseParams struct {
	// 验证命名空间是否存在返回信息
	Data *NamespaceIsExistsResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ValidateNamespaceExistPersonalResponse struct {
	*tchttp.BaseResponse
	Response *ValidateNamespaceExistPersonalResponseParams `json:"Response"`
}

func (r *ValidateNamespaceExistPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ValidateNamespaceExistPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ValidateRepositoryExistPersonalRequestParams struct {
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

type ValidateRepositoryExistPersonalRequest struct {
	*tchttp.BaseRequest
	
	// 仓库名称
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *ValidateRepositoryExistPersonalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ValidateRepositoryExistPersonalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RepoName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ValidateRepositoryExistPersonalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ValidateRepositoryExistPersonalResponseParams struct {
	// 验证个人版仓库是否存在返回信息
	Data *RepoIsExistResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ValidateRepositoryExistPersonalResponse struct {
	*tchttp.BaseResponse
	Response *ValidateRepositoryExistPersonalResponseParams `json:"Response"`
}

func (r *ValidateRepositoryExistPersonalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ValidateRepositoryExistPersonalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcAndDomainInfo struct {
	// tcr实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// tcr内网访问链路ip
	EniLBIp *string `json:"EniLBIp,omitempty" name:"EniLBIp"`

	// true：use instance name as subdomain
	// false: use instancename+"-vpc" as subdomain
	UsePublicDomain *bool `json:"UsePublicDomain,omitempty" name:"UsePublicDomain"`

	// 解析地域，需要保证和vpc处于同一地域，如果不填则默认为主实例地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type VpcPrivateDomainStatus struct {
	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// unique vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// ENABLE代表已经开启，DISABLE代表未开启，ERROR代表查询出错
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type WebhookTarget struct {
	// 目标地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 自定义 Headers
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

type WebhookTrigger struct {
	// 触发器名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 触发器目标
	Targets []*WebhookTarget `json:"Targets,omitempty" name:"Targets"`

	// 触发动作
	EventTypes []*string `json:"EventTypes,omitempty" name:"EventTypes"`

	// 触发规则
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 启用触发器
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 触发器Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 触发器描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 触发器所属命名空间 Id
	NamespaceId *int64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

type WebhookTriggerLog struct {
	// 日志 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 触发器 Id
	TriggerId *int64 `json:"TriggerId,omitempty" name:"TriggerId"`

	// 事件类型
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 通知类型
	NotifyType *string `json:"NotifyType,omitempty" name:"NotifyType"`

	// 详情
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`
}