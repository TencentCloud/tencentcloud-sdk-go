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

package v20220325

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ChargePrepaidConfig struct {
	// 预付费购买周期，单位：月
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 预付费自动续费设置：
	// NOTIFY_AND_MANUAL_RENEW：表示默认状态(用户未设置，即初始状态)， NOTIFY_AND_AUTO_RENEW：表示自动续费，DISABLE_NOTIFY_AND_MANUAL_RENEW：表示明确不自动续费(用户设置)
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

// Predefined struct for user
type CreateEtcdInstanceRequestParams struct {
	// etcd实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// etcd所属vpc
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// etcd对外提供访问ip地址所在子网
	ServiceSubnetId *string `json:"ServiceSubnetId,omitnil,omitempty" name:"ServiceSubnetId"`

	// etcd部署子网
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// etcd版本
	EtcdVersion *string `json:"EtcdVersion,omitnil,omitempty" name:"EtcdVersion"`

	// etcd节点个数，可选范围: 1, 3, 5, 7, 9
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// etcd集群描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 高级设置
	AdvancedSettings *EtcdAdvancedSettings `json:"AdvancedSettings,omitnil,omitempty" name:"AdvancedSettings"`

	// 付费类型：
	// PREPAID 预付费
	// POSTPAID_BY_HOUR 后付费按小时付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 预付费相关配置
	ChargePrepaid *ChargePrepaidConfig `json:"ChargePrepaid,omitnil,omitempty" name:"ChargePrepaid"`

	// 存储大小GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 删除保护，true 删除保护开启；false删除保护关闭
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type CreateEtcdInstanceRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// etcd所属vpc
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// etcd对外提供访问ip地址所在子网
	ServiceSubnetId *string `json:"ServiceSubnetId,omitnil,omitempty" name:"ServiceSubnetId"`

	// etcd部署子网
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// etcd版本
	EtcdVersion *string `json:"EtcdVersion,omitnil,omitempty" name:"EtcdVersion"`

	// etcd节点个数，可选范围: 1, 3, 5, 7, 9
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// etcd集群描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 高级设置
	AdvancedSettings *EtcdAdvancedSettings `json:"AdvancedSettings,omitnil,omitempty" name:"AdvancedSettings"`

	// 付费类型：
	// PREPAID 预付费
	// POSTPAID_BY_HOUR 后付费按小时付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 预付费相关配置
	ChargePrepaid *ChargePrepaidConfig `json:"ChargePrepaid,omitnil,omitempty" name:"ChargePrepaid"`

	// 存储大小GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 删除保护，true 删除保护开启；false删除保护关闭
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *CreateEtcdInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEtcdInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "VpcId")
	delete(f, "ServiceSubnetId")
	delete(f, "SubnetIds")
	delete(f, "EtcdVersion")
	delete(f, "Size")
	delete(f, "Description")
	delete(f, "AdvancedSettings")
	delete(f, "ChargeType")
	delete(f, "ChargePrepaid")
	delete(f, "DiskSize")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEtcdInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEtcdInstanceResponseParams struct {
	// 创建etcd实例的Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEtcdInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateEtcdInstanceResponseParams `json:"Response"`
}

func (r *CreateEtcdInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEtcdInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEtcdSnapshotPolicyRequestParams struct {
	// etcd实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照策略名
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil,omitempty" name:"SnapshotPolicyName"`

	// 备份参数设置
	BackupSettings *EtcdBackupSettings `json:"BackupSettings,omitnil,omitempty" name:"BackupSettings"`
}

type CreateEtcdSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照策略名
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil,omitempty" name:"SnapshotPolicyName"`

	// 备份参数设置
	BackupSettings *EtcdBackupSettings `json:"BackupSettings,omitnil,omitempty" name:"BackupSettings"`
}

func (r *CreateEtcdSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEtcdSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotPolicyName")
	delete(f, "BackupSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEtcdSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEtcdSnapshotPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEtcdSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateEtcdSnapshotPolicyResponseParams `json:"Response"`
}

func (r *CreateEtcdSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEtcdSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEtcdSnapshotRequestParams struct {
	// etcd实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 备份用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 备份密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateEtcdSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 备份用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 备份密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *CreateEtcdSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEtcdSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotName")
	delete(f, "User")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEtcdSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEtcdSnapshotResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEtcdSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateEtcdSnapshotResponseParams `json:"Response"`
}

func (r *CreateEtcdSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEtcdSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdAvailableVersionsRequestParams struct {

}

type DescribeEtcdAvailableVersionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEtcdAvailableVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdAvailableVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdAvailableVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdAvailableVersionsResponseParams struct {
	// 支持的版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Versions []*VersionInstance `json:"Versions,omitnil,omitempty" name:"Versions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdAvailableVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdAvailableVersionsResponseParams `json:"Response"`
}

func (r *DescribeEtcdAvailableVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdAvailableVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdCreatingProgressRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeEtcdCreatingProgressRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeEtcdCreatingProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdCreatingProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdCreatingProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdCreatingProgressResponseParams struct {
	// 创建进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Steps []*TaskStepInfo `json:"Steps,omitnil,omitempty" name:"Steps"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdCreatingProgressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdCreatingProgressResponseParams `json:"Response"`
}

func (r *DescribeEtcdCreatingProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdCreatingProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdCredentialsRequestParams struct {
	// etcd实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeEtcdCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeEtcdCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdCredentialsResponseParams struct {
	// 访问凭证
	// 注意：此字段可能返回 null，表示取不到有效值。
	Credentials []*EtcdCredential `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdCredentialsResponseParams `json:"Response"`
}

func (r *DescribeEtcdCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdInstancesRequestParams struct {
	// 按照一个或者多个实例ID查询。实例ID形如：etcd-xxxxxxxx。参数不支持同时指定InstanceIds和Filters
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 支持按照vpc-id和instance-id过滤。参数不支持同时指定InstanceIds和Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，最大值为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeEtcdInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询。实例ID形如：etcd-xxxxxxxx。参数不支持同时指定InstanceIds和Filters
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 支持按照vpc-id和instance-id过滤。参数不支持同时指定InstanceIds和Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，最大值为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeEtcdInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdInstancesResponseParams struct {
	// etcd实例详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Etcds []*Etcd `json:"Etcds,omitnil,omitempty" name:"Etcds"`

	// 符合条件的实例数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdInstancesResponseParams `json:"Response"`
}

func (r *DescribeEtcdInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdQuotaRequestParams struct {

}

type DescribeEtcdQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEtcdQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdQuotaResponseParams struct {
	// 当前配额限制
	QuotaLimit *uint64 `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdQuotaResponseParams `json:"Response"`
}

func (r *DescribeEtcdQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdRegionsRequestParams struct {

}

type DescribeEtcdRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEtcdRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdRegionsResponseParams struct {
	// 支持的地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions []*RegionInstance `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdRegionsResponseParams `json:"Response"`
}

func (r *DescribeEtcdRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdSnapshotPoliciesRequestParams struct {
	// etcd实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeEtcdSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeEtcdSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdSnapshotPoliciesResponseParams struct {
	// 备份策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotPolicies []*EtcdSnapshotPolicy `json:"SnapshotPolicies,omitnil,omitempty" name:"SnapshotPolicies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DescribeEtcdSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdSnapshotsRequestParams struct {
	// etcd实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeEtcdSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeEtcdSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdSnapshotsResponseParams struct {
	// etcd快照列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Snapshots []*EtcdSnapshot `json:"Snapshots,omitnil,omitempty" name:"Snapshots"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeEtcdSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdTasksRequestParams struct {
	// 任务ID
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// taskType：
	//     按照任务类型过滤，当前支持enable_internet，disable_internet, restore_remote_snapshot
	// resourceId：
	//     按照资源ID过滤
	// lifeState：
	//     按照任务状态过滤，当前支持process， done
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEtcdTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// taskType：
	//     按照任务类型过滤，当前支持enable_internet，disable_internet, restore_remote_snapshot
	// resourceId：
	//     按照资源ID过滤
	// lifeState：
	//     按照任务状态过滤，当前支持process， done
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEtcdTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEtcdTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEtcdTasksResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*EtcdTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEtcdTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEtcdTasksResponseParams `json:"Response"`
}

func (r *DescribeEtcdTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEtcdTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRPCMethodListRequestParams struct {
	// etcd实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd集群pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大长度
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRPCMethodListRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd集群pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大长度
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRPCMethodListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRPCMethodListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PodName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRPCMethodListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRPCMethodListResponseParams struct {
	// gRPC方法列表
	MethodList []*RPCMethod `json:"MethodList,omitnil,omitempty" name:"MethodList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRPCMethodListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRPCMethodListResponseParams `json:"Response"`
}

func (r *DescribeRPCMethodListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRPCMethodListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableEtcdInstanceDeletionProtectionRequestParams struct {
	// etcd实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DisableEtcdInstanceDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DisableEtcdInstanceDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableEtcdInstanceDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableEtcdInstanceDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableEtcdInstanceDeletionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableEtcdInstanceDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DisableEtcdInstanceDeletionProtectionResponseParams `json:"Response"`
}

func (r *DisableEtcdInstanceDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableEtcdInstanceDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableEtcdInstanceDeletionProtectionRequestParams struct {
	// etcd实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type EnableEtcdInstanceDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *EnableEtcdInstanceDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableEtcdInstanceDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableEtcdInstanceDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableEtcdInstanceDeletionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableEtcdInstanceDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *EnableEtcdInstanceDeletionProtectionResponseParams `json:"Response"`
}

func (r *EnableEtcdInstanceDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableEtcdInstanceDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Etcd struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 实例所属vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// etcd版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 实例状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// etcd成员信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Members []*EtcdMember `json:"Members,omitnil,omitempty" name:"Members"`

	// 对外访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 删除保护，true 删除保护开启；false删除保护关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type EtcdAdvancedSettings struct {
	// 安全相关设置
	SecuritySettings *EtcdSecuritySettings `json:"SecuritySettings,omitnil,omitempty" name:"SecuritySettings"`

	// 自动压缩设置
	AutoCompactionSettings *EtcdAutoCompactionSettings `json:"AutoCompactionSettings,omitnil,omitempty" name:"AutoCompactionSettings"`

	// 监控设置
	MonitorSettings *EtcdMonitorSettings `json:"MonitorSettings,omitnil,omitempty" name:"MonitorSettings"`

	// 备份相关设置
	BackupSettings *EtcdBackupSettings `json:"BackupSettings,omitnil,omitempty" name:"BackupSettings"`

	// 自定义域名
	CustomDomains *string `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`

	// 自定义ip
	CustomIPs *string `json:"CustomIPs,omitnil,omitempty" name:"CustomIPs"`
}

type EtcdAutoCompactionSettings struct {
	// 自动压缩模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCompactionMode *string `json:"AutoCompactionMode,omitnil,omitempty" name:"AutoCompactionMode"`

	// 自动压缩保留时间/revison数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCompactionRetention *string `json:"AutoCompactionRetention,omitnil,omitempty" name:"AutoCompactionRetention"`
}

type EtcdBackupSettings struct {
	// 备份间隔(s)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupInterval *int64 `json:"BackupInterval,omitnil,omitempty" name:"BackupInterval"`

	// 最大备份个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxBackupCount *int64 `json:"MaxBackupCount,omitnil,omitempty" name:"MaxBackupCount"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// COS桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`
}

type EtcdCredential struct {
	// CA证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	CACert *string `json:"CACert,omitnil,omitempty" name:"CACert"`

	// 证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// 私钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

type EtcdMember struct {
	// 节点名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点当前版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 所属可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 节点状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type EtcdMonitorSettings struct {
	// Prometheus创建参数
	PrometheusCreationParam *PrometheusCreationParam `json:"PrometheusCreationParam,omitnil,omitempty" name:"PrometheusCreationParam"`

	// Prometheus Id
	ExistedPrometheusInstanceId *string `json:"ExistedPrometheusInstanceId,omitnil,omitempty" name:"ExistedPrometheusInstanceId"`
}

type EtcdSecuritySettings struct {
	// 是否启用https
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *bool `json:"Https,omitnil,omitempty" name:"Https"`

	// 启用客户端证书认证
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertAuth *bool `json:"ClientCertAuth,omitnil,omitempty" name:"ClientCertAuth"`
}

type EtcdSnapshot struct {
	// 快照名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 快照大小
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type EtcdSnapshotPolicy struct {
	// 快照策略名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备份参数
	BackupSettings *EtcdBackupSettings `json:"BackupSettings,omitnil,omitempty" name:"BackupSettings"`
}

type EtcdTaskInfo struct {
	// 任务ID
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资源ID
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`

	// 任务状态
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 任务创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 任务更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type InstanceConfig struct {
	// 核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小Gi
	Mem *uint64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 集群规模
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

// Predefined struct for user
type ModifyEtcdAttributeRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// root账号密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 同步auth状态
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 删除集群关联的cos数据
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`

	// 子网id，数组
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

type ModifyEtcdAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// root账号密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 同步auth状态
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 删除集群关联的cos数据
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`

	// 子网id，数组
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`
}

func (r *ModifyEtcdAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEtcdAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Password")
	delete(f, "EnableAuth")
	delete(f, "EnableDeleteCos")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEtcdAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEtcdAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEtcdAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEtcdAttributeResponseParams `json:"Response"`
}

func (r *ModifyEtcdAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEtcdAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEtcdConfigurationRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动压缩设置
	AutoCompactionSettings *EtcdAutoCompactionSettings `json:"AutoCompactionSettings,omitnil,omitempty" name:"AutoCompactionSettings"`

	// 监控设置
	MonitorSettings *EtcdMonitorSettings `json:"MonitorSettings,omitnil,omitempty" name:"MonitorSettings"`

	// 计费类型
	// PREPAID：预付费
	// POSTPAID_BY_HOUR：后付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 实例配置信息
	InstanceConfig *InstanceConfig `json:"InstanceConfig,omitnil,omitempty" name:"InstanceConfig"`

	// 预付费配置信息
	PrepaidConfig *ChargePrepaidConfig `json:"PrepaidConfig,omitnil,omitempty" name:"PrepaidConfig"`
}

type ModifyEtcdConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自动压缩设置
	AutoCompactionSettings *EtcdAutoCompactionSettings `json:"AutoCompactionSettings,omitnil,omitempty" name:"AutoCompactionSettings"`

	// 监控设置
	MonitorSettings *EtcdMonitorSettings `json:"MonitorSettings,omitnil,omitempty" name:"MonitorSettings"`

	// 计费类型
	// PREPAID：预付费
	// POSTPAID_BY_HOUR：后付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 实例配置信息
	InstanceConfig *InstanceConfig `json:"InstanceConfig,omitnil,omitempty" name:"InstanceConfig"`

	// 预付费配置信息
	PrepaidConfig *ChargePrepaidConfig `json:"PrepaidConfig,omitnil,omitempty" name:"PrepaidConfig"`
}

func (r *ModifyEtcdConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEtcdConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AutoCompactionSettings")
	delete(f, "MonitorSettings")
	delete(f, "ChargeType")
	delete(f, "InstanceConfig")
	delete(f, "PrepaidConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEtcdConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEtcdConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEtcdConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEtcdConfigurationResponseParams `json:"Response"`
}

func (r *ModifyEtcdConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEtcdConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEtcdSnapshotPolicyRequestParams struct {
	// etcd实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照策略名称
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil,omitempty" name:"SnapshotPolicyName"`

	// 备份参数设置
	BackupSettings *EtcdBackupSettings `json:"BackupSettings,omitnil,omitempty" name:"BackupSettings"`
}

type ModifyEtcdSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// etcd实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照策略名称
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil,omitempty" name:"SnapshotPolicyName"`

	// 备份参数设置
	BackupSettings *EtcdBackupSettings `json:"BackupSettings,omitnil,omitempty" name:"BackupSettings"`
}

func (r *ModifyEtcdSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEtcdSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotPolicyName")
	delete(f, "BackupSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEtcdSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEtcdSnapshotPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEtcdSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEtcdSnapshotPolicyResponseParams `json:"Response"`
}

func (r *ModifyEtcdSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEtcdSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrometheusCreationParam struct {
	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 保存时长，只支持天单位
	// 12d = 12天
	DataRetention *int64 `json:"DataRetention,omitnil,omitempty" name:"DataRetention"`

	// grafana用户名
	GrafanaUserName *string `json:"GrafanaUserName,omitnil,omitempty" name:"GrafanaUserName"`

	// grafana密码
	GrafanaPassword *string `json:"GrafanaPassword,omitnil,omitempty" name:"GrafanaPassword"`
}

type RPCMethod struct {
	// 方法名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type RegionInstance struct {
	// 地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域特性开关(按照JSON的形式返回所有属性)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureGates *string `json:"FeatureGates,omitnil,omitempty" name:"FeatureGates"`

	// 地域简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 地域白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type ResizeEtcdInstanceRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd节点个数
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type ResizeEtcdInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd节点个数
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

func (r *ResizeEtcdInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeEtcdInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Size")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeEtcdInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeEtcdInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeEtcdInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResizeEtcdInstanceResponseParams `json:"Response"`
}

func (r *ResizeEtcdInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeEtcdInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskStepInfo struct {
	// 步骤名称
	Step *string `json:"Step,omitnil,omitempty" name:"Step"`

	// 生命周期
	// pending : 步骤未开始
	// running: 步骤执行中
	// success: 步骤成功完成
	// failed: 步骤失败
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 步骤开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 步骤结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// 若步骤生命周期为failed,则此字段显示错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedMsg *string `json:"FailedMsg,omitnil,omitempty" name:"FailedMsg"`
}

// Predefined struct for user
type UpgradeEtcdInstanceRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd版本
	EtcdVersion *string `json:"EtcdVersion,omitnil,omitempty" name:"EtcdVersion"`
}

type UpgradeEtcdInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// etcd版本
	EtcdVersion *string `json:"EtcdVersion,omitnil,omitempty" name:"EtcdVersion"`
}

func (r *UpgradeEtcdInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeEtcdInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EtcdVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeEtcdInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeEtcdInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeEtcdInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeEtcdInstanceResponseParams `json:"Response"`
}

func (r *UpgradeEtcdInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeEtcdInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionInstance struct {
	// 版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Remark信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}