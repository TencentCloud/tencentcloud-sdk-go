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

package v20200915

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ActionAlterCkUserRequestParams struct {
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// api接口类型
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type ActionAlterCkUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// api接口类型
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

func (r *ActionAlterCkUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActionAlterCkUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserInfo")
	delete(f, "ApiType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActionAlterCkUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActionAlterCkUserResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ActionAlterCkUserResponse struct {
	*tchttp.BaseResponse
	Response *ActionAlterCkUserResponseParams `json:"Response"`
}

func (r *ActionAlterCkUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActionAlterCkUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupTableContent struct {
	// 数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// 表
	Table *string `json:"Table,omitempty" name:"Table"`

	// 表总字节数
	TotalBytes *int64 `json:"TotalBytes,omitempty" name:"TotalBytes"`

	// 虚拟cluster
	VCluster *string `json:"VCluster,omitempty" name:"VCluster"`

	// 表ip
	Ips *string `json:"Ips,omitempty" name:"Ips"`
}

type CkUserAlterInfo struct {
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 描述
	Describe *string `json:"Describe,omitempty" name:"Describe"`
}

type ConfigSubmitContext struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 配置文件旧内容，base64编码
	OldConfValue *string `json:"OldConfValue,omitempty" name:"OldConfValue"`

	// 配置文件新内容，base64编码
	NewConfValue *string `json:"NewConfValue,omitempty" name:"NewConfValue"`

	// 保存配置文件的路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
}

// Predefined struct for user
type CreateBackUpScheduleRequestParams struct {
	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔
	WeekDays *string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitempty" name:"BackUpTables"`
}

type CreateBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔
	WeekDays *string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitempty" name:"BackUpTables"`
}

func (r *CreateBackUpScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	delete(f, "WeekDays")
	delete(f, "ExecuteHour")
	delete(f, "BackUpTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackUpScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackUpScheduleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackUpScheduleResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackUpScheduleResponseParams `json:"Response"`
}

func (r *CreateBackUpScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkSqlApisRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// api接口名称
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 集群名称
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 用户名称
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DescribeCkSqlApisRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// api接口名称
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 集群名称
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 用户名称
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DescribeCkSqlApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkSqlApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ApiType")
	delete(f, "Cluster")
	delete(f, "UserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkSqlApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkSqlApisResponseParams struct {
	// 返回的查询数据，大部分情况是list，也可能是bool
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnData *string `json:"ReturnData,omitempty" name:"ReturnData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCkSqlApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCkSqlApisResponseParams `json:"Response"`
}

func (r *DescribeCkSqlApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkSqlApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceShardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsResponseParams struct {
	// 实例shard信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceShardsList *string `json:"InstanceShardsList,omitempty" name:"InstanceShardsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceShardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceShardsResponseParams `json:"Response"`
}

func (r *DescribeInstanceShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecRequestParams struct {
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitempty" name:"IsElastic"`
}

type DescribeSpecRequest struct {
	*tchttp.BaseRequest
	
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitempty" name:"IsElastic"`
}

func (r *DescribeSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "PayMode")
	delete(f, "IsElastic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecResponseParams struct {
	// zookeeper节点规格描述
	CommonSpec []*ResourceSpec `json:"CommonSpec,omitempty" name:"CommonSpec"`

	// 数据节点规格描述
	DataSpec []*ResourceSpec `json:"DataSpec,omitempty" name:"DataSpec"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec []*DiskSpec `json:"AttachCBSSpec,omitempty" name:"AttachCBSSpec"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecResponseParams `json:"Response"`
}

func (r *DescribeSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskSpec struct {
	// 磁盘类型，例如“CLOUD_SSD", "LOCAL_SSD"等
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘类型说明，例如"云SSD", "本地SSD"等
	DiskDesc *string `json:"DiskDesc,omitempty" name:"DiskDesc"`

	// 磁盘最小规格大小，单位G
	MinDiskSize *int64 `json:"MinDiskSize,omitempty" name:"MinDiskSize"`

	// 磁盘最大规格大小，单位G
	MaxDiskSize *int64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`

	// 磁盘数目
	DiskCount *int64 `json:"DiskCount,omitempty" name:"DiskCount"`
}

// Predefined struct for user
type ModifyClusterConfigsRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyClusterConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ModifyConfContext")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterConfigsResponseParams struct {
	// 流程相关信息
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterConfigsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterConfigsResponseParams `json:"Response"`
}

func (r *ModifyClusterConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserNewPrivilegeRequestParams struct {

}

type ModifyUserNewPrivilegeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyUserNewPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserNewPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserNewPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserNewPrivilegeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserNewPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserNewPrivilegeResponseParams `json:"Response"`
}

func (r *ModifyUserNewPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserNewPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenBackUpRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`
}

type OpenBackUpRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`
}

func (r *OpenBackUpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenBackUpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OperationType")
	delete(f, "CosBucketName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenBackUpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenBackUpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenBackUpResponse struct {
	*tchttp.BaseResponse
	Response *OpenBackUpResponseParams `json:"Response"`
}

func (r *OpenBackUpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenBackUpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSpec struct {
	// 规格名称，例如“SCH1"
	Name *string `json:"Name,omitempty" name:"Name"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位G
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 分类标记，STANDARD/BIGDATA/HIGHIO分别表示标准型/大数据型/高IO
	Type *string `json:"Type,omitempty" name:"Type"`

	// 系统盘描述信息
	SystemDisk *DiskSpec `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘描述信息
	DataDisk *DiskSpec `json:"DataDisk,omitempty" name:"DataDisk"`

	// 最大节点数目限制
	MaxNodeSize *int64 `json:"MaxNodeSize,omitempty" name:"MaxNodeSize"`

	// 是否可用，false代表售罄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Available *bool `json:"Available,omitempty" name:"Available"`

	// 规格描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeSpecDesc *string `json:"ComputeSpecDesc,omitempty" name:"ComputeSpecDesc"`

	// 规格名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 库存数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
}