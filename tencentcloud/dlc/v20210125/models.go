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

package v20210125

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddDMSPartitionsRequestParams struct {
	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`
}

type AddDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`
}

func (r *AddDMSPartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDMSPartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Partitions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDMSPartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDMSPartitionsResponseParams struct {
	// 成功数量
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 分区值
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddDMSPartitionsResponse struct {
	*tchttp.BaseResponse
	Response *AddDMSPartitionsResponseParams `json:"Response"`
}

func (r *AddDMSPartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDMSPartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersToWorkGroupRequestParams struct {
	// 要操作的工作组和用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil" name:"AddInfo"`
}

type AddUsersToWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要操作的工作组和用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil" name:"AddInfo"`
}

func (r *AddUsersToWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersToWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUsersToWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersToWorkGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddUsersToWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddUsersToWorkGroupResponseParams `json:"Response"`
}

func (r *AddUsersToWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersToWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSDatabaseRequestParams struct {
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil" name:"CurrentName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 路径
	Location *string `json:"Location,omitnil" name:"Location"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`
}

type AlterDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil" name:"CurrentName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 路径
	Location *string `json:"Location,omitnil" name:"Location"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`
}

func (r *AlterDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrentName")
	delete(f, "SchemaName")
	delete(f, "Location")
	delete(f, "Asset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AlterDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSDatabaseResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AlterDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *AlterDMSDatabaseResponseParams `json:"Response"`
}

func (r *AlterDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSPartitionRequestParams struct {
	// 当前名称，变更前db名称
	CurrentDbName *string `json:"CurrentDbName,omitnil" name:"CurrentDbName"`

	// 当前名称，变更前table名称
	CurrentTableName *string `json:"CurrentTableName,omitnil" name:"CurrentTableName"`

	// 当前名称，变更前Part名称
	CurrentValues *string `json:"CurrentValues,omitnil" name:"CurrentValues"`

	// 分区
	Partition *DMSPartition `json:"Partition,omitnil" name:"Partition"`
}

type AlterDMSPartitionRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称，变更前db名称
	CurrentDbName *string `json:"CurrentDbName,omitnil" name:"CurrentDbName"`

	// 当前名称，变更前table名称
	CurrentTableName *string `json:"CurrentTableName,omitnil" name:"CurrentTableName"`

	// 当前名称，变更前Part名称
	CurrentValues *string `json:"CurrentValues,omitnil" name:"CurrentValues"`

	// 分区
	Partition *DMSPartition `json:"Partition,omitnil" name:"Partition"`
}

func (r *AlterDMSPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrentDbName")
	delete(f, "CurrentTableName")
	delete(f, "CurrentValues")
	delete(f, "Partition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AlterDMSPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSPartitionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AlterDMSPartitionResponse struct {
	*tchttp.BaseResponse
	Response *AlterDMSPartitionResponseParams `json:"Response"`
}

func (r *AlterDMSPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSTableRequestParams struct {
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil" name:"CurrentName"`

	// 当前数据库名称
	CurrentDbName *string `json:"CurrentDbName,omitnil" name:"CurrentDbName"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 当前表名
	Name *string `json:"Name,omitnil" name:"Name"`
}

type AlterDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil" name:"CurrentName"`

	// 当前数据库名称
	CurrentDbName *string `json:"CurrentDbName,omitnil" name:"CurrentDbName"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 当前表名
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *AlterDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrentName")
	delete(f, "CurrentDbName")
	delete(f, "Asset")
	delete(f, "Type")
	delete(f, "DbName")
	delete(f, "StorageSize")
	delete(f, "RecordCount")
	delete(f, "LifeTime")
	delete(f, "DataUpdateTime")
	delete(f, "StructUpdateTime")
	delete(f, "LastAccessTime")
	delete(f, "Sds")
	delete(f, "Columns")
	delete(f, "PartitionKeys")
	delete(f, "ViewOriginalText")
	delete(f, "ViewExpandedText")
	delete(f, "Partitions")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AlterDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSTableResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AlterDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *AlterDMSTableResponseParams `json:"Response"`
}

func (r *AlterDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Asset struct {
	// 主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 对象GUID值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guid *string `json:"Guid,omitnil" name:"Guid"`

	// 数据目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 对象owner
	Owner *string `json:"Owner,omitnil" name:"Owner"`

	// 对象owner账户
	OwnerAccount *string `json:"OwnerAccount,omitnil" name:"OwnerAccount"`

	// 权限
	PermValues []*KVPair `json:"PermValues,omitnil" name:"PermValues"`

	// 附加属性
	Params []*KVPair `json:"Params,omitnil" name:"Params"`

	// 附加业务属性
	BizParams []*KVPair `json:"BizParams,omitnil" name:"BizParams"`

	// 数据版本
	DataVersion *int64 `json:"DataVersion,omitnil" name:"DataVersion"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitnil" name:"ModifiedTime"`

	// 数据源主键
	DatasourceId *int64 `json:"DatasourceId,omitnil" name:"DatasourceId"`
}

// Predefined struct for user
type AttachUserPolicyRequestParams struct {
	// 用户Id，和子用户uin相同，需要先使用CreateUser接口创建用户。可以使用DescribeUsers接口查看。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和子用户uin相同，需要先使用CreateUser接口创建用户。可以使用DescribeUsers接口查看。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

func (r *AttachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachUserPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AttachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachUserPolicyResponseParams `json:"Response"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachWorkGroupPolicyRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 要绑定的策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

type AttachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 要绑定的策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

func (r *AttachWorkGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachWorkGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachWorkGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachWorkGroupPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AttachWorkGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachWorkGroupPolicyResponseParams `json:"Response"`
}

func (r *AttachWorkGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachWorkGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchSqlTask struct {
	// SQL子任务唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 运行SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteSQL *string `json:"ExecuteSQL,omitnil" name:"ExecuteSQL"`

	// 任务信息，成功则返回：Task Success!，失败则返回异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`
}

// Predefined struct for user
type BindWorkGroupsToUserRequestParams struct {
	// 绑定的用户和工作组信息
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil" name:"AddInfo"`
}

type BindWorkGroupsToUserRequest struct {
	*tchttp.BaseRequest
	
	// 绑定的用户和工作组信息
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil" name:"AddInfo"`
}

func (r *BindWorkGroupsToUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindWorkGroupsToUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindWorkGroupsToUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindWorkGroupsToUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BindWorkGroupsToUserResponse struct {
	*tchttp.BaseResponse
	Response *BindWorkGroupsToUserResponseParams `json:"Response"`
}

func (r *BindWorkGroupsToUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindWorkGroupsToUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CSV struct {
	// 压缩格式，["Snappy", "Gzip", "None"选一]。
	CodeCompress *string `json:"CodeCompress,omitnil" name:"CodeCompress"`

	// CSV序列化及反序列化数据结构。
	CSVSerde *CSVSerde `json:"CSVSerde,omitnil" name:"CSVSerde"`

	// 标题行，默认为0。
	HeadLines *int64 `json:"HeadLines,omitnil" name:"HeadLines"`

	// 格式，默认值为CSV
	Format *string `json:"Format,omitnil" name:"Format"`
}

type CSVSerde struct {
	// CSV序列化转义符，默认为"\\"，最长8个字符，如 Escape: "/\"
	Escape *string `json:"Escape,omitnil" name:"Escape"`

	// CSV序列化字段域符，默认为"'"，最长8个字符, 如 Quote: "\""
	Quote *string `json:"Quote,omitnil" name:"Quote"`

	// CSV序列化分隔符，默认为"\t"，最长8个字符, 如 Separator: "\t"
	Separator *string `json:"Separator,omitnil" name:"Separator"`
}

// Predefined struct for user
type CancelNotebookSessionStatementBatchRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

type CancelNotebookSessionStatementBatchRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

func (r *CancelNotebookSessionStatementBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelNotebookSessionStatementBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelNotebookSessionStatementBatchResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelNotebookSessionStatementBatchResponse struct {
	*tchttp.BaseResponse
	Response *CancelNotebookSessionStatementBatchResponseParams `json:"Response"`
}

func (r *CancelNotebookSessionStatementBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelNotebookSessionStatementRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil" name:"StatementId"`
}

type CancelNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil" name:"StatementId"`
}

func (r *CancelNotebookSessionStatementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "StatementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelNotebookSessionStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelNotebookSessionStatementResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelNotebookSessionStatementResponse struct {
	*tchttp.BaseResponse
	Response *CancelNotebookSessionStatementResponseParams `json:"Response"`
}

func (r *CancelNotebookSessionStatementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelSparkSessionBatchSQLRequestParams struct {
	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

type CancelSparkSessionBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

func (r *CancelSparkSessionBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelSparkSessionBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelSparkSessionBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelSparkSessionBatchSQLResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelSparkSessionBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *CancelSparkSessionBatchSQLResponseParams `json:"Response"`
}

func (r *CancelSparkSessionBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelSparkSessionBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTaskRequestParams struct {
	// 任务Id，全局唯一
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id，全局唯一
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelTaskResponseParams `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineConfigPairsValidityRequestParams struct {
	// 引擎小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil" name:"ChildImageVersionId"`

	// 用户自定义参数
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil" name:"DataEngineConfigPairs"`

	// 引擎大版本ID，存在小版本ID时仅需传入小版本ID，不存在时会获取当前大版本下最新的小版本ID。
	ImageVersionId *string `json:"ImageVersionId,omitnil" name:"ImageVersionId"`
}

type CheckDataEngineConfigPairsValidityRequest struct {
	*tchttp.BaseRequest
	
	// 引擎小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil" name:"ChildImageVersionId"`

	// 用户自定义参数
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil" name:"DataEngineConfigPairs"`

	// 引擎大版本ID，存在小版本ID时仅需传入小版本ID，不存在时会获取当前大版本下最新的小版本ID。
	ImageVersionId *string `json:"ImageVersionId,omitnil" name:"ImageVersionId"`
}

func (r *CheckDataEngineConfigPairsValidityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineConfigPairsValidityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChildImageVersionId")
	delete(f, "DataEngineConfigPairs")
	delete(f, "ImageVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDataEngineConfigPairsValidityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineConfigPairsValidityResponseParams struct {
	// 参数有效性：ture:有效，false:至少存在一个无效参数；
	IsAvailable *bool `json:"IsAvailable,omitnil" name:"IsAvailable"`

	// 无效参数集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnavailableConfig []*string `json:"UnavailableConfig,omitnil" name:"UnavailableConfig"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckDataEngineConfigPairsValidityResponse struct {
	*tchttp.BaseResponse
	Response *CheckDataEngineConfigPairsValidityResponseParams `json:"Response"`
}

func (r *CheckDataEngineConfigPairsValidityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineConfigPairsValidityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeRollbackRequestParams struct {
	// 引擎唯一id
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

type CheckDataEngineImageCanBeRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 引擎唯一id
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

func (r *CheckDataEngineImageCanBeRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDataEngineImageCanBeRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeRollbackResponseParams struct {
	// 回滚后日志记录id
	ToRecordId *string `json:"ToRecordId,omitnil" name:"ToRecordId"`

	// 回滚前日志记录id
	FromRecordId *string `json:"FromRecordId,omitnil" name:"FromRecordId"`

	// 是否能够回滚
	IsRollback *bool `json:"IsRollback,omitnil" name:"IsRollback"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckDataEngineImageCanBeRollbackResponse struct {
	*tchttp.BaseResponse
	Response *CheckDataEngineImageCanBeRollbackResponseParams `json:"Response"`
}

func (r *CheckDataEngineImageCanBeRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeUpgradeRequestParams struct {
	// 集群id
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

type CheckDataEngineImageCanBeUpgradeRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

func (r *CheckDataEngineImageCanBeUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeUpgradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDataEngineImageCanBeUpgradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeUpgradeResponseParams struct {
	// 当前大版本下，可升级的集群镜像小版本id
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil" name:"ChildImageVersionId"`

	// 是否能够升级
	IsUpgrade *bool `json:"IsUpgrade,omitnil" name:"IsUpgrade"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckDataEngineImageCanBeUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *CheckDataEngineImageCanBeUpgradeResponseParams `json:"Response"`
}

func (r *CheckDataEngineImageCanBeUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckLockMetaDataRequestParams struct {
	// 锁ID
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil" name:"TxnId"`

	// 过期时间ms
	ElapsedMs *int64 `json:"ElapsedMs,omitnil" name:"ElapsedMs"`
}

type CheckLockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 锁ID
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil" name:"TxnId"`

	// 过期时间ms
	ElapsedMs *int64 `json:"ElapsedMs,omitnil" name:"ElapsedMs"`
}

func (r *CheckLockMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckLockMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LockId")
	delete(f, "DatasourceConnectionName")
	delete(f, "TxnId")
	delete(f, "ElapsedMs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckLockMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckLockMetaDataResponseParams struct {
	// 锁ID
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 锁状态：ACQUIRED、WAITING、ABORT、NOT_ACQUIRED
	LockState *string `json:"LockState,omitnil" name:"LockState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckLockMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *CheckLockMetaDataResponseParams `json:"Response"`
}

func (r *CheckLockMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckLockMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {
	// 列名称，不区分大小写，最大支持25个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 列类型，支持如下类型定义:
	// string|tinyint|smallint|int|bigint|boolean|float|double|decimal|timestamp|date|binary|array<data_type>|map<primitive_type, data_type>|struct<col_name : data_type [COMMENT col_comment], ...>|uniontype<data_type, data_type, ...>。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 对该类的注释。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 表示整个 numeric 的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Precision *int64 `json:"Precision,omitnil" name:"Precision"`

	// 表示小数部分的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`

	// 是否为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nullable *string `json:"Nullable,omitnil" name:"Nullable"`

	// 字段位置，小的在前
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitnil" name:"Position"`

	// 字段创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 字段修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil" name:"ModifiedTime"`

	// 是否为分区字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitnil" name:"IsPartition"`
}

type CommonMetrics struct {
	// 创建任务时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTaskTime *float64 `json:"CreateTaskTime,omitnil" name:"CreateTaskTime"`

	// 预处理总时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessTime *float64 `json:"ProcessTime,omitnil" name:"ProcessTime"`

	// 排队时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueTime *float64 `json:"QueueTime,omitnil" name:"QueueTime"`

	// 执行时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionTime *float64 `json:"ExecutionTime,omitnil" name:"ExecutionTime"`

	// 是否命中结果缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsResultCacheHit *bool `json:"IsResultCacheHit,omitnil" name:"IsResultCacheHit"`

	// 匹配物化视图数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchedMVBytes *int64 `json:"MatchedMVBytes,omitnil" name:"MatchedMVBytes"`

	// 匹配物化视图列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchedMVs *string `json:"MatchedMVs,omitnil" name:"MatchedMVs"`

	// 结果数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedBytes *string `json:"AffectedBytes,omitnil" name:"AffectedBytes"`

	// 	结果行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedRows *int64 `json:"AffectedRows,omitnil" name:"AffectedRows"`

	// 扫描数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedBytes *int64 `json:"ProcessedBytes,omitnil" name:"ProcessedBytes"`

	// 	扫描行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedRows *int64 `json:"ProcessedRows,omitnil" name:"ProcessedRows"`
}

type CosPermission struct {
	// cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil" name:"CosPath"`

	// 权限【"read","write"】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`
}

// Predefined struct for user
type CreateDMSDatabaseRequestParams struct {
	// 基础元数据对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// Schema目录
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// Db存储路径
	Location *string `json:"Location,omitnil" name:"Location"`

	// 数据库名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type CreateDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 基础元数据对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// Schema目录
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// Db存储路径
	Location *string `json:"Location,omitnil" name:"Location"`

	// 数据库名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *CreateDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Asset")
	delete(f, "SchemaName")
	delete(f, "Location")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDMSDatabaseResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDMSDatabaseResponseParams `json:"Response"`
}

func (r *CreateDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDMSTableRequestParams struct {
	// 基础对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type CreateDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 基础对象
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *CreateDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Asset")
	delete(f, "Type")
	delete(f, "DbName")
	delete(f, "StorageSize")
	delete(f, "RecordCount")
	delete(f, "LifeTime")
	delete(f, "DataUpdateTime")
	delete(f, "StructUpdateTime")
	delete(f, "LastAccessTime")
	delete(f, "Sds")
	delete(f, "Columns")
	delete(f, "PartitionKeys")
	delete(f, "ViewOriginalText")
	delete(f, "ViewExpandedText")
	delete(f, "Partitions")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDMSTableResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateDMSTableResponseParams `json:"Response"`
}

func (r *CreateDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataEngineRequestParams struct {
	// 引擎类型spark/presto
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 集群类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 计费模式 0=共享模式 1=按量计费 2=包年包月
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 是否自动启动集群
	AutoResume *bool `json:"AutoResume,omitnil" name:"AutoResume"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil" name:"MaxClusters"`

	// 是否为默认虚拟集群
	//
	// Deprecated: DefaultDataEngine is deprecated.
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitnil" name:"DefaultDataEngine"`

	// VPC网段
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 描述信息
	Message *string `json:"Message,omitnil" name:"Message"`

	// 集群规模
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 计费类型，后付费：0，预付费：1。当前只支持后付费，不填默认为后付费。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 资源使用时长，后付费：固定填3600，预付费：最少填1，代表购买资源一个月，最长不超过120。默认1
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 资源使用时长的单位，后付费：s，预付费：m。默认为s
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 资源的自动续费标志。后付费无需续费，固定填0；预付费下：0表示手动续费、1代表自动续费、2代表不续费，在0下如果是大客户，会自动帮大客户续费。默认为0
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 创建资源的时候需要绑定的标签信息
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，有效值：SQL/BATCH，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitnil" name:"EngineExecType"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil" name:"TolerableQueueTime"`

	// 集群自动挂起时间，默认10分钟
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil" name:"AutoSuspendTime"`

	// 资源类型。Standard_CU：标准型；Memory_CU：内存型
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 集群高级配置
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil" name:"DataEngineConfigPairs"`

	// 集群镜像版本名字。如SuperSQL-P 1.1;SuperSQL-S 3.2等,不传，默认创建最新镜像版本的集群
	ImageVersionName *string `json:"ImageVersionName,omitnil" name:"ImageVersionName"`

	// 主集群名称，创建容灾集群时指定
	MainClusterName *string `json:"MainClusterName,omitnil" name:"MainClusterName"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil" name:"ElasticLimit"`

	// spark作业集群session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`

	// 自动授权
	AutoAuthorization *bool `json:"AutoAuthorization,omitnil" name:"AutoAuthorization"`

	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil" name:"EngineNetworkId"`

	// 引擎世代，SuperSQL：代表supersql引擎，Native：代表标准引擎。默认值为SuperSQL
	EngineGeneration *string `json:"EngineGeneration,omitnil" name:"EngineGeneration"`
}

type CreateDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型spark/presto
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 集群类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 计费模式 0=共享模式 1=按量计费 2=包年包月
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 是否自动启动集群
	AutoResume *bool `json:"AutoResume,omitnil" name:"AutoResume"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil" name:"MaxClusters"`

	// 是否为默认虚拟集群
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitnil" name:"DefaultDataEngine"`

	// VPC网段
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 描述信息
	Message *string `json:"Message,omitnil" name:"Message"`

	// 集群规模
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 计费类型，后付费：0，预付费：1。当前只支持后付费，不填默认为后付费。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 资源使用时长，后付费：固定填3600，预付费：最少填1，代表购买资源一个月，最长不超过120。默认1
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 资源使用时长的单位，后付费：s，预付费：m。默认为s
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 资源的自动续费标志。后付费无需续费，固定填0；预付费下：0表示手动续费、1代表自动续费、2代表不续费，在0下如果是大客户，会自动帮大客户续费。默认为0
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 创建资源的时候需要绑定的标签信息
	Tags []*TagInfo `json:"Tags,omitnil" name:"Tags"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，有效值：SQL/BATCH，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitnil" name:"EngineExecType"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil" name:"TolerableQueueTime"`

	// 集群自动挂起时间，默认10分钟
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil" name:"AutoSuspendTime"`

	// 资源类型。Standard_CU：标准型；Memory_CU：内存型
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 集群高级配置
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil" name:"DataEngineConfigPairs"`

	// 集群镜像版本名字。如SuperSQL-P 1.1;SuperSQL-S 3.2等,不传，默认创建最新镜像版本的集群
	ImageVersionName *string `json:"ImageVersionName,omitnil" name:"ImageVersionName"`

	// 主集群名称，创建容灾集群时指定
	MainClusterName *string `json:"MainClusterName,omitnil" name:"MainClusterName"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil" name:"ElasticLimit"`

	// spark作业集群session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`

	// 自动授权
	AutoAuthorization *bool `json:"AutoAuthorization,omitnil" name:"AutoAuthorization"`

	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil" name:"EngineNetworkId"`

	// 引擎世代，SuperSQL：代表supersql引擎，Native：代表标准引擎。默认值为SuperSQL
	EngineGeneration *string `json:"EngineGeneration,omitnil" name:"EngineGeneration"`
}

func (r *CreateDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	delete(f, "DataEngineName")
	delete(f, "ClusterType")
	delete(f, "Mode")
	delete(f, "AutoResume")
	delete(f, "MinClusters")
	delete(f, "MaxClusters")
	delete(f, "DefaultDataEngine")
	delete(f, "CidrBlock")
	delete(f, "Message")
	delete(f, "Size")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenew")
	delete(f, "Tags")
	delete(f, "AutoSuspend")
	delete(f, "CrontabResumeSuspend")
	delete(f, "CrontabResumeSuspendStrategy")
	delete(f, "EngineExecType")
	delete(f, "MaxConcurrency")
	delete(f, "TolerableQueueTime")
	delete(f, "AutoSuspendTime")
	delete(f, "ResourceType")
	delete(f, "DataEngineConfigPairs")
	delete(f, "ImageVersionName")
	delete(f, "MainClusterName")
	delete(f, "ElasticSwitch")
	delete(f, "ElasticLimit")
	delete(f, "SessionResourceTemplate")
	delete(f, "AutoAuthorization")
	delete(f, "EngineNetworkId")
	delete(f, "EngineGeneration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataEngineResponseParams struct {
	// 虚拟引擎id
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataEngineResponseParams `json:"Response"`
}

func (r *CreateDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseRequestParams struct {
	// 数据库基础信息
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitnil" name:"DatabaseInfo"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库基础信息
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitnil" name:"DatabaseInfo"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseInfo")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseResponseParams struct {
	// 生成的建库执行语句对象。
	Execution *Execution `json:"Execution,omitnil" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatabaseResponseParams `json:"Response"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportTaskRequestParams struct {
	// 数据来源，lakefsStorage、taskResult
	InputType *string `json:"InputType,omitnil" name:"InputType"`

	// 导出任务输入配置
	InputConf []*KVPair `json:"InputConf,omitnil" name:"InputConf"`

	// 导出任务输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil" name:"OutputConf"`

	// 目标数据源的类型，目前支持导出到cos
	OutputType *string `json:"OutputType,omitnil" name:"OutputType"`
}

type CreateExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据来源，lakefsStorage、taskResult
	InputType *string `json:"InputType,omitnil" name:"InputType"`

	// 导出任务输入配置
	InputConf []*KVPair `json:"InputConf,omitnil" name:"InputConf"`

	// 导出任务输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil" name:"OutputConf"`

	// 目标数据源的类型，目前支持导出到cos
	OutputType *string `json:"OutputType,omitnil" name:"OutputType"`
}

func (r *CreateExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputType")
	delete(f, "InputConf")
	delete(f, "OutputConf")
	delete(f, "OutputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateExportTaskResponseParams `json:"Response"`
}

func (r *CreateExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImportTaskRequestParams struct {
	// 数据来源，cos
	InputType *string `json:"InputType,omitnil" name:"InputType"`

	// 输入配置
	InputConf []*KVPair `json:"InputConf,omitnil" name:"InputConf"`

	// 输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil" name:"OutputConf"`

	// 目标数据源的类型，目前支持导入到托管存储，即lakefsStorage
	OutputType *string `json:"OutputType,omitnil" name:"OutputType"`
}

type CreateImportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据来源，cos
	InputType *string `json:"InputType,omitnil" name:"InputType"`

	// 输入配置
	InputConf []*KVPair `json:"InputConf,omitnil" name:"InputConf"`

	// 输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil" name:"OutputConf"`

	// 目标数据源的类型，目前支持导入到托管存储，即lakefsStorage
	OutputType *string `json:"OutputType,omitnil" name:"OutputType"`
}

func (r *CreateImportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputType")
	delete(f, "InputConf")
	delete(f, "OutputConf")
	delete(f, "OutputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImportTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateImportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateImportTaskResponseParams `json:"Response"`
}

func (r *CreateImportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalTableRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`
}

type CreateInternalTableRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`
}

func (r *CreateInternalTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "Partitions")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInternalTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalTableResponseParams struct {
	// 创建托管存储内表sql语句描述
	Execution *string `json:"Execution,omitnil" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInternalTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateInternalTableResponseParams `json:"Response"`
}

func (r *CreateInternalTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionRequestParams struct {
	// Session名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// session文件地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitnil" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitnil" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitnil" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	ProgramArchives []*string `json:"ProgramArchives,omitnil" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil" name:"ExecutorNumbers"`

	// Session相关配置，当前支持：
	// 1. dlc.eni: 用户配置的eni网关信息，可以通过该字段设置；
	// 2. dlc.role.arn: 用户配置的roleArn鉴权策略配置信息，可以通过该字段设置；
	// 3. dlc.sql.set.config: 用户配置的集群配置信息，可以通过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil" name:"Arguments"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitnil" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil" name:"TimeoutInSecond"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil" name:"ExecutorMaxNumbers"`

	// 指定spark版本名称，当前任务使用该spark镜像运行
	SparkImage *string `json:"SparkImage,omitnil" name:"SparkImage"`

	// 是否继承集群的资源类配置：0：自定义（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil" name:"IsInherit"`
}

type CreateNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// session文件地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitnil" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitnil" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitnil" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	ProgramArchives []*string `json:"ProgramArchives,omitnil" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil" name:"ExecutorNumbers"`

	// Session相关配置，当前支持：
	// 1. dlc.eni: 用户配置的eni网关信息，可以通过该字段设置；
	// 2. dlc.role.arn: 用户配置的roleArn鉴权策略配置信息，可以通过该字段设置；
	// 3. dlc.sql.set.config: 用户配置的集群配置信息，可以通过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil" name:"Arguments"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitnil" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil" name:"TimeoutInSecond"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil" name:"ExecutorMaxNumbers"`

	// 指定spark版本名称，当前任务使用该spark镜像运行
	SparkImage *string `json:"SparkImage,omitnil" name:"SparkImage"`

	// 是否继承集群的资源类配置：0：自定义（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil" name:"IsInherit"`
}

func (r *CreateNotebookSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Kind")
	delete(f, "DataEngineName")
	delete(f, "ProgramDependentFiles")
	delete(f, "ProgramDependentJars")
	delete(f, "ProgramDependentPython")
	delete(f, "ProgramArchives")
	delete(f, "DriverSize")
	delete(f, "ExecutorSize")
	delete(f, "ExecutorNumbers")
	delete(f, "Arguments")
	delete(f, "ProxyUser")
	delete(f, "TimeoutInSecond")
	delete(f, "ExecutorMaxNumbers")
	delete(f, "SparkImage")
	delete(f, "IsInherit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionResponseParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// Spark任务返回的AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitnil" name:"State"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNotebookSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookSessionResponseParams `json:"Response"`
}

func (r *CreateNotebookSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`
}

type CreateNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`
}

func (r *CreateNotebookSessionStatementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Code")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookSessionStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementResponseParams struct {
	// Session Statement详情
	NotebookSessionStatement *NotebookSessionStatementInfo `json:"NotebookSessionStatement,omitnil" name:"NotebookSessionStatement"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNotebookSessionStatementResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookSessionStatementResponseParams `json:"Response"`
}

func (r *CreateNotebookSessionStatementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementSupportBatchSQLRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 类型，当前支持：sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 是否保存运行结果
	SaveResult *bool `json:"SaveResult,omitnil" name:"SaveResult"`
}

type CreateNotebookSessionStatementSupportBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 类型，当前支持：sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// 是否保存运行结果
	SaveResult *bool `json:"SaveResult,omitnil" name:"SaveResult"`
}

func (r *CreateNotebookSessionStatementSupportBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementSupportBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Code")
	delete(f, "Kind")
	delete(f, "SaveResult")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookSessionStatementSupportBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementSupportBatchSQLResponseParams struct {
	// Session Statement详情
	NotebookSessionStatementBatches *NotebookSessionStatementBatchInformation `json:"NotebookSessionStatementBatches,omitnil" name:"NotebookSessionStatementBatches"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNotebookSessionStatementSupportBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookSessionStatementSupportBatchSQLResponseParams `json:"Response"`
}

func (r *CreateNotebookSessionStatementSupportBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementSupportBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResultDownloadRequestParams struct {
	// 查询结果任务Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 下载格式
	Format *string `json:"Format,omitnil" name:"Format"`

	// 是否重新生成下载文件，仅当之前任务为 Timout | Error 时有效
	Force *bool `json:"Force,omitnil" name:"Force"`
}

type CreateResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 查询结果任务Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 下载格式
	Format *string `json:"Format,omitnil" name:"Format"`

	// 是否重新生成下载文件，仅当之前任务为 Timout | Error 时有效
	Force *bool `json:"Force,omitnil" name:"Force"`
}

func (r *CreateResultDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResultDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Format")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResultDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResultDownloadResponseParams struct {
	// 下载任务Id
	DownloadId *string `json:"DownloadId,omitnil" name:"DownloadId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateResultDownloadResponse struct {
	*tchttp.BaseResponse
	Response *CreateResultDownloadResponseParams `json:"Response"`
}

func (r *CreateResultDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResultDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScriptRequestParams struct {
	// 脚本名称，最大不能超过255个字符。
	ScriptName *string `json:"ScriptName,omitnil" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitnil" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitnil" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`
}

type CreateScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本名称，最大不能超过255个字符。
	ScriptName *string `json:"ScriptName,omitnil" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitnil" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitnil" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`
}

func (r *CreateScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptName")
	delete(f, "SQLStatement")
	delete(f, "ScriptDesc")
	delete(f, "DatabaseName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScriptResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateScriptResponse struct {
	*tchttp.BaseResponse
	Response *CreateScriptResponseParams `json:"Response"`
}

func (r *CreateScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppRequestParams struct {
	// spark作业名
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil" name:"AppFile"`

	// 数据访问策略，CAM Role arn
	RoleArn *int64 `json:"RoleArn,omitnil" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil" name:"AppExecutorNums"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil" name:"AppFiles"`

	// spark作业程序入参，空格分割
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil" name:"MaxRetries"`

	// 数据源名称
	DataSource *string `json:"DataSource,omitnil" name:"DataSource"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil" name:"AppPythonFiles"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本id
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil" name:"IsSessionStarted"`
}

type CreateSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil" name:"AppFile"`

	// 数据访问策略，CAM Role arn
	RoleArn *int64 `json:"RoleArn,omitnil" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil" name:"AppExecutorNums"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil" name:"AppFiles"`

	// spark作业程序入参，空格分割
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil" name:"MaxRetries"`

	// 数据源名称
	DataSource *string `json:"DataSource,omitnil" name:"DataSource"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil" name:"AppPythonFiles"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本id
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil" name:"IsSessionStarted"`
}

func (r *CreateSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	delete(f, "IsLocalArchives")
	delete(f, "AppArchives")
	delete(f, "SparkImage")
	delete(f, "SparkImageVersion")
	delete(f, "AppExecutorMaxNumbers")
	delete(f, "SessionId")
	delete(f, "IsInherit")
	delete(f, "IsSessionStarted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppResponseParams struct {
	// App唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkAppResponseParams `json:"Response"`
}

func (r *CreateSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppTaskRequestParams struct {
	// spark作业名
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// spark作业程序入参，以空格分隔；一般用于周期性调用使用
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`
}

type CreateSparkAppTaskRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// spark作业程序入参，以空格分隔；一般用于周期性调用使用
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`
}

func (r *CreateSparkAppTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "CmdArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppTaskResponseParams struct {
	// 批Id
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSparkAppTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkAppTaskResponseParams `json:"Response"`
}

func (r *CreateSparkAppTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkSessionBatchSQLRequestParams struct {
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 运行sql，需要base64编码。
	ExecuteSQL *string `json:"ExecuteSQL,omitnil" name:"ExecuteSQL"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil" name:"ExecutorNumbers"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil" name:"ExecutorMaxNumbers"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil" name:"TimeoutInSecond"`

	// Session唯一标识，当指定sessionid，则使用该session运行任务。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 指定要创建的session名称
	SessionName *string `json:"SessionName,omitnil" name:"SessionName"`

	// Session相关配置，当前支持：1.dlc.eni：用户配置的eni网关信息，可以用过该字段设置；
	// 2.dlc.role.arn：用户配置的roleArn鉴权策略配置信息，可以用过该字段设置；
	// 3.dlc.sql.set.config：用户配置的集群配置信息，可以用过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil" name:"Arguments"`

	// 是否继承集群的资源类配置：0：不继承（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil" name:"IsInherit"`
}

type CreateSparkSessionBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 运行sql，需要base64编码。
	ExecuteSQL *string `json:"ExecuteSQL,omitnil" name:"ExecuteSQL"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil" name:"ExecutorNumbers"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil" name:"ExecutorMaxNumbers"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil" name:"TimeoutInSecond"`

	// Session唯一标识，当指定sessionid，则使用该session运行任务。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 指定要创建的session名称
	SessionName *string `json:"SessionName,omitnil" name:"SessionName"`

	// Session相关配置，当前支持：1.dlc.eni：用户配置的eni网关信息，可以用过该字段设置；
	// 2.dlc.role.arn：用户配置的roleArn鉴权策略配置信息，可以用过该字段设置；
	// 3.dlc.sql.set.config：用户配置的集群配置信息，可以用过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil" name:"Arguments"`

	// 是否继承集群的资源类配置：0：不继承（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil" name:"IsInherit"`
}

func (r *CreateSparkSessionBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkSessionBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "ExecuteSQL")
	delete(f, "DriverSize")
	delete(f, "ExecutorSize")
	delete(f, "ExecutorNumbers")
	delete(f, "ExecutorMaxNumbers")
	delete(f, "TimeoutInSecond")
	delete(f, "SessionId")
	delete(f, "SessionName")
	delete(f, "Arguments")
	delete(f, "IsInherit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkSessionBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkSessionBatchSQLResponseParams struct {
	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`

	// Statement任务列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Statements []*StatementInformation `json:"Statements,omitnil" name:"Statements"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSparkSessionBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkSessionBatchSQLResponseParams `json:"Response"`
}

func (r *CreateSparkSessionBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkSessionBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStoreLocationRequestParams struct {
	// 计算结果存储cos路径，如：cosn://bucketname/
	StoreLocation *string `json:"StoreLocation,omitnil" name:"StoreLocation"`
}

type CreateStoreLocationRequest struct {
	*tchttp.BaseRequest
	
	// 计算结果存储cos路径，如：cosn://bucketname/
	StoreLocation *string `json:"StoreLocation,omitnil" name:"StoreLocation"`
}

func (r *CreateStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StoreLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStoreLocationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *CreateStoreLocationResponseParams `json:"Response"`
}

func (r *CreateStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableRequestParams struct {
	// 数据表配置信息
	TableInfo *TableInfo `json:"TableInfo,omitnil" name:"TableInfo"`
}

type CreateTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据表配置信息
	TableInfo *TableInfo `json:"TableInfo,omitnil" name:"TableInfo"`
}

func (r *CreateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableResponseParams struct {
	// 生成的建表执行语句对象。
	Execution *Execution `json:"Execution,omitnil" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateTableResponseParams `json:"Response"`
}

func (r *CreateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 计算任务，该参数中包含任务类型及其相关配置信息
	Task *Task `json:"Task,omitnil" name:"Task"`

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 默认数据源名称。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 数据引擎名称，不填提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 计算任务，该参数中包含任务类型及其相关配置信息
	Task *Task `json:"Task,omitnil" name:"Task"`

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 默认数据源名称。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 数据引擎名称，不填提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Task")
	delete(f, "DatabaseName")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksInOrderRequestParams struct {
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil" name:"Tasks"`

	// 数据源名称，默认为COSDataCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

type CreateTasksInOrderRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil" name:"Tasks"`

	// 数据源名称，默认为COSDataCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

func (r *CreateTasksInOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksInOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Tasks")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTasksInOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksInOrderResponseParams struct {
	// 本批次提交的任务的批次Id
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`

	// 任务Id集合，按照执行顺序排列
	TaskIdSet []*string `json:"TaskIdSet,omitnil" name:"TaskIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTasksInOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateTasksInOrderResponseParams `json:"Response"`
}

func (r *CreateTasksInOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksInOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksRequestParams struct {
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil" name:"Tasks"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 计算引擎名称，不填任务提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

type CreateTasksRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil" name:"Tasks"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 计算引擎名称，不填任务提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

func (r *CreateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Tasks")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksResponseParams struct {
	// 本批次提交的任务的批次Id
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`

	// 任务Id集合，按照执行顺序排列
	TaskIdSet []*string `json:"TaskIdSet,omitnil" name:"TaskIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateTasksResponseParams `json:"Response"`
}

func (r *CreateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 需要授权的子用户uin，可以通过腾讯云控制台右上角 → “账号信息” → “账号ID进行查看”。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitnil" name:"UserDescription"`

	// 绑定到用户的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`

	// 用户类型。ADMIN：管理员 COMMON：一般用户。当用户类型为管理员的时候，不能设置权限集合和绑定的工作组集合，管理员默认拥有所有权限。该参数不填默认为COMMON
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 绑定到用户的工作组ID集合。
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil" name:"WorkGroupIds"`

	// 用户别名，字符长度小50
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 需要授权的子用户uin，可以通过腾讯云控制台右上角 → “账号信息” → “账号ID进行查看”。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitnil" name:"UserDescription"`

	// 绑定到用户的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`

	// 用户类型。ADMIN：管理员 COMMON：一般用户。当用户类型为管理员的时候，不能设置权限集合和绑定的工作组集合，管理员默认拥有所有权限。该参数不填默认为COMMON
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 绑定到用户的工作组ID集合。
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil" name:"WorkGroupIds"`

	// 用户别名，字符长度小50
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserDescription")
	delete(f, "PolicySet")
	delete(f, "UserType")
	delete(f, "WorkGroupIds")
	delete(f, "UserAlias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkGroupRequestParams struct {
	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil" name:"WorkGroupDescription"`

	// 工作组绑定的鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`

	// 需要绑定到工作组的用户Id集合
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

type CreateWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil" name:"WorkGroupDescription"`

	// 工作组绑定的鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`

	// 需要绑定到工作组的用户Id集合
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

func (r *CreateWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupName")
	delete(f, "WorkGroupDescription")
	delete(f, "PolicySet")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkGroupResponseParams struct {
	// 工作组Id，全局唯一
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkGroupResponseParams `json:"Response"`
}

func (r *CreateWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrontabResumeSuspendStrategy struct {
	// 定时拉起时间：如：周一8点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResumeTime *string `json:"ResumeTime,omitnil" name:"ResumeTime"`

	// 定时挂起时间：如：周一20点
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuspendTime *string `json:"SuspendTime,omitnil" name:"SuspendTime"`

	// 挂起配置：0（默认）：等待任务结束后挂起、1：强制挂起
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuspendStrategy *int64 `json:"SuspendStrategy,omitnil" name:"SuspendStrategy"`
}

type DMSColumn struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitnil" name:"Position"`

	// 附加参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*KVPair `json:"Params,omitnil" name:"Params"`

	// 业务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizParams []*KVPair `json:"BizParams,omitnil" name:"BizParams"`

	// 是否分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitnil" name:"IsPartition"`
}

type DMSColumnOrder struct {
	// 列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Col *string `json:"Col,omitnil" name:"Col"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitnil" name:"Order"`
}

type DMSPartition struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 数据目录名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 数据版本
	DataVersion *int64 `json:"DataVersion,omitnil" name:"DataVersion"`

	// 分区名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值列表
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitnil" name:"ModifiedTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil" name:"LastAccessTime"`

	// 附件属性
	Params []*KVPair `json:"Params,omitnil" name:"Params"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil" name:"Sds"`
}

type DMSSds struct {
	// 存储地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil" name:"Location"`

	// 输入格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputFormat *string `json:"InputFormat,omitnil" name:"InputFormat"`

	// 输出格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputFormat *string `json:"OutputFormat,omitnil" name:"OutputFormat"`

	// bucket数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumBuckets *int64 `json:"NumBuckets,omitnil" name:"NumBuckets"`

	// 是是否压缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compressed *bool `json:"Compressed,omitnil" name:"Compressed"`

	// 是否有子目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoredAsSubDirectories *bool `json:"StoredAsSubDirectories,omitnil" name:"StoredAsSubDirectories"`

	// 序列化lib
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeLib *string `json:"SerdeLib,omitnil" name:"SerdeLib"`

	// 序列化名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeName *string `json:"SerdeName,omitnil" name:"SerdeName"`

	// 桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketCols []*string `json:"BucketCols,omitnil" name:"BucketCols"`

	// 序列化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeParams []*KVPair `json:"SerdeParams,omitnil" name:"SerdeParams"`

	// 附加参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*KVPair `json:"Params,omitnil" name:"Params"`

	// 列排序(Expired)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortCols *DMSColumnOrder `json:"SortCols,omitnil" name:"SortCols"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cols []*DMSColumn `json:"Cols,omitnil" name:"Cols"`

	// 列排序字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortColumns []*DMSColumnOrder `json:"SortColumns,omitnil" name:"SortColumns"`
}

type DMSTable struct {
	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewOriginalText *string `json:"ViewOriginalText,omitnil" name:"ViewOriginalText"`

	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewExpandedText *string `json:"ViewExpandedText,omitnil" name:"ViewExpandedText"`

	// hive维护版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitnil" name:"Retention"`

	// 存储对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sds *DMSSds `json:"Sds,omitnil" name:"Sds"`

	// 分区列
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil" name:"PartitionKeys"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeTime *int64 `json:"LifeTime,omitnil" name:"LifeTime"`

	// 最后访问时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAccessTime *string `json:"LastAccessTime,omitnil" name:"LastAccessTime"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUpdateTime *string `json:"DataUpdateTime,omitnil" name:"DataUpdateTime"`

	// 结构更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StructUpdateTime *string `json:"StructUpdateTime,omitnil" name:"StructUpdateTime"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*DMSColumn `json:"Columns,omitnil" name:"Columns"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DMSTableInfo struct {
	// DMS表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *DMSTable `json:"Table,omitnil" name:"Table"`

	// 基础对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`
}

type DataEngineBasicInfo struct {
	// DataEngine名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 数据引擎状态  -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中
	State *int64 `json:"State,omitnil" name:"State"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 引擎类型，有效值：PrestoSQL/SparkSQL/SparkBatch
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineType *string `json:"DataEngineType,omitnil" name:"DataEngineType"`

	// 用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`
}

type DataEngineConfigInstanceInfo struct {
	// 引擎ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 用户自定义配置项集合
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil" name:"DataEngineConfigPairs"`

	// 作业集群资源参数配置模版
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`
}

type DataEngineConfigPair struct {
	// 配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigItem *string `json:"ConfigItem,omitnil" name:"ConfigItem"`

	// 配置值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigValue *string `json:"ConfigValue,omitnil" name:"ConfigValue"`
}

type DataEngineImageVersion struct {
	// 镜像大版本ID
	ImageVersionId *string `json:"ImageVersionId,omitnil" name:"ImageVersionId"`

	// 镜像大版本名称
	ImageVersion *string `json:"ImageVersion,omitnil" name:"ImageVersion"`

	// 镜像大版本描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否为公共版本：1：公共；2：私有
	IsPublic *uint64 `json:"IsPublic,omitnil" name:"IsPublic"`

	// 集群类型：SparkSQL/PrestoSQL/SparkBatch
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 版本状态：1：初始化；2：上线；3：下线
	IsSharedEngine *uint64 `json:"IsSharedEngine,omitnil" name:"IsSharedEngine"`

	// 版本状态：1：初始化；2：上线；3：下线
	State *uint64 `json:"State,omitnil" name:"State"`

	// 插入时间
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type DataEngineInfo struct {
	// DataEngine名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 引擎类型 spark/presto
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 集群资源类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitnil" name:"ClusterType"`

	// 引用ID
	QuotaId *string `json:"QuotaId,omitnil" name:"QuotaId"`

	// 数据引擎状态  -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中
	State *int64 `json:"State,omitnil" name:"State"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 集群规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 计费模式 0共享模式 1按量计费 2包年包月
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 最小集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinClusters *int64 `json:"MinClusters,omitnil" name:"MinClusters"`

	// 最大集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxClusters *int64 `json:"MaxClusters,omitnil" name:"MaxClusters"`

	// 是否自动恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoResume *bool `json:"AutoResume,omitnil" name:"AutoResume"`

	// 自动恢复时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpendAfter *int64 `json:"SpendAfter,omitnil" name:"SpendAfter"`

	// 集群网段
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 是否为默认引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitnil" name:"DefaultDataEngine"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 操作者
	SubAccountUin *string `json:"SubAccountUin,omitnil" name:"SubAccountUin"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 隔离时间
	IsolatedTime *string `json:"IsolatedTime,omitnil" name:"IsolatedTime"`

	// 冲正时间
	ReversalTime *string `json:"ReversalTime,omitnil" name:"ReversalTime"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`

	// 标签对集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagInfo `json:"TagList,omitnil" name:"TagList"`

	// 引擎拥有的权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSuspend *bool `json:"AutoSuspend,omitnil" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，有效值：SQL/BATCH
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineExecType *string `json:"EngineExecType,omitnil" name:"EngineExecType"`

	// 自动续费标志，0，初始状态，默认不自动续费，若用户有预付费不停服特权，自动续费。1：自动续费。2：明确不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 集群自动挂起时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil" name:"AutoSuspendTime"`

	// 网络连接配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionSet []*NetworkConnection `json:"NetworkConnectionSet,omitnil" name:"NetworkConnectionSet"`

	// ui的跳转地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	UiURL *string `json:"UiURL,omitnil" name:"UiURL"`

	// 引擎的资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 集群镜像版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersionId *string `json:"ImageVersionId,omitnil" name:"ImageVersionId"`

	// 集群镜像小版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil" name:"ChildImageVersionId"`

	// 集群镜像版本名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersionName *string `json:"ImageVersionName,omitnil" name:"ImageVersionName"`

	// 是否开启备集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitnil" name:"StartStandbyCluster"`

	// spark jar 包年包月集群是否开启弹性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticLimit *int64 `json:"ElasticLimit,omitnil" name:"ElasticLimit"`

	// 是否为默认引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultHouse *bool `json:"DefaultHouse,omitnil" name:"DefaultHouse"`

	// 单个集群任务最大并发数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil" name:"MaxConcurrency"`

	// 任务排队上限时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil" name:"TolerableQueueTime"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAppId *int64 `json:"UserAppId,omitnil" name:"UserAppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`

	// SessionResourceTemplate
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`

	// 自动授权开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoAuthorization *bool `json:"AutoAuthorization,omitnil" name:"AutoAuthorization"`

	// 引擎版本，支持Native/SuperSQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineGeneration *string `json:"EngineGeneration,omitnil" name:"EngineGeneration"`

	// 引擎详细类型，支持：SparkSQL/SparkBatch/PrestoSQL/Kyuubi
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil" name:"EngineTypeDetail"`
}

type DataFormat struct {
	// 文本格式，TextFile。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextFile *TextFile `json:"TextFile,omitnil" name:"TextFile"`

	// 文本格式，CSV。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CSV *CSV `json:"CSV,omitnil" name:"CSV"`

	// 文本格式，Json。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *Other `json:"Json,omitnil" name:"Json"`

	// Parquet格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parquet *Other `json:"Parquet,omitnil" name:"Parquet"`

	// ORC格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ORC *Other `json:"ORC,omitnil" name:"ORC"`

	// AVRO格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AVRO *Other `json:"AVRO,omitnil" name:"AVRO"`
}

type DataGovernPolicy struct {
	// 治理规则类型，Customize: 自定义；Intelligence: 智能治理
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 治理引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernEngine *string `json:"GovernEngine,omitnil" name:"GovernEngine"`
}

type DataSourceInfo struct {
	// 数据源实例的唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据源的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 数据源的JDBC访问链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	JdbcUrl *string `json:"JdbcUrl,omitnil" name:"JdbcUrl"`

	// 用于访问数据源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 数据源访问密码，需要base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 数据源的VPC和子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *DatasourceConnectionLocation `json:"Location,omitnil" name:"Location"`

	// 默认数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`
}

type DatabaseInfo struct {
	// 数据库名称，长度0~128，支持数字、字母下划线，不允许数字大头，统一转换为小写。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 数据库描述信息，长度 0~500。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 数据库属性列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`

	// 数据库cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil" name:"Location"`
}

type DatabaseResponseInfo struct {
	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 数据库描述信息，长度 0~256。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 允许针对数据库的属性元数据信息进行指定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`

	// 数据库创建时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 数据库更新时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil" name:"ModifiedTime"`

	// cos存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil" name:"Location"`

	// 建库用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`

	// 建库用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSubUin *string `json:"UserSubUin,omitnil" name:"UserSubUin"`

	// 数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernPolicy *DataGovernPolicy `json:"GovernPolicy,omitnil" name:"GovernPolicy"`

	// 数据库ID（无效字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil" name:"DatabaseId"`
}

type DatasourceConnectionConfig struct {
	// Mysql数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mysql *MysqlInfo `json:"Mysql,omitnil" name:"Mysql"`

	// Hive数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hive *HiveInfo `json:"Hive,omitnil" name:"Hive"`

	// Kafka数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kafka *KafkaInfo `json:"Kafka,omitnil" name:"Kafka"`

	// 其他数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherDatasourceConnection *OtherDatasourceConnection `json:"OtherDatasourceConnection,omitnil" name:"OtherDatasourceConnection"`

	// PostgreSQL数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSql *DataSourceInfo `json:"PostgreSql,omitnil" name:"PostgreSql"`

	// SQLServer数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlServer *DataSourceInfo `json:"SqlServer,omitnil" name:"SqlServer"`

	// ClickHouse数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouse *DataSourceInfo `json:"ClickHouse,omitnil" name:"ClickHouse"`

	// Elasticsearch数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elasticsearch *ElasticsearchInfo `json:"Elasticsearch,omitnil" name:"Elasticsearch"`

	// TDSQL-PostgreSQL数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	TDSQLPostgreSql *DataSourceInfo `json:"TDSQLPostgreSql,omitnil" name:"TDSQLPostgreSql"`
}

type DatasourceConnectionInfo struct {
	// 数据源数字Id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 数据源字符串Id
	DatasourceConnectionId *string `json:"DatasourceConnectionId,omitnil" name:"DatasourceConnectionId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 数据源描述
	DatasourceConnectionDesc *string `json:"DatasourceConnectionDesc,omitnil" name:"DatasourceConnectionDesc"`

	// 数据源类型，支持DataLakeCatalog、IcebergCatalog、Result、Mysql、HiveCos、HiveHdfs
	DatasourceConnectionType *string `json:"DatasourceConnectionType,omitnil" name:"DatasourceConnectionType"`

	// 数据源属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionConfig *DatasourceConnectionConfig `json:"DatasourceConnectionConfig,omitnil" name:"DatasourceConnectionConfig"`

	// 数据源状态：0（初始化）、1（成功）、-1（已删除）、-2（失败）、-3（删除中）
	State *int64 `json:"State,omitnil" name:"State"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 用户AppId
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 数据源创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 数据源最近一次更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 数据源同步失败原因
	Message *string `json:"Message,omitnil" name:"Message"`

	// 数据源绑定的计算引擎信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngines []*DataEngineInfo `json:"DataEngines,omitnil" name:"DataEngines"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`

	// 网络配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionSet []*NetworkConnection `json:"NetworkConnectionSet,omitnil" name:"NetworkConnectionSet"`

	// 连通性状态：0（未测试，默认）、1（正常）、2（失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectivityState *uint64 `json:"ConnectivityState,omitnil" name:"ConnectivityState"`

	// 连通性测试提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectivityTips *string `json:"ConnectivityTips,omitnil" name:"ConnectivityTips"`
}

type DatasourceConnectionLocation struct {
	// 数据连接所在Vpc实例Id，如“vpc-azd4dt1c”。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Vpc的IPv4 CIDR
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil" name:"VpcCidrBlock"`

	// 数据连接所在子网的实例Id，如“subnet-bthucmmy”
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Subnet的IPv4 CIDR
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitnil" name:"SubnetCidrBlock"`
}

// Predefined struct for user
type DeleteDataEngineRequestParams struct {
	// 删除虚拟集群的名称数组
	DataEngineNames []*string `json:"DataEngineNames,omitnil" name:"DataEngineNames"`
}

type DeleteDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 删除虚拟集群的名称数组
	DataEngineNames []*string `json:"DataEngineNames,omitnil" name:"DataEngineNames"`
}

func (r *DeleteDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataEngineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataEngineResponseParams `json:"Response"`
}

func (r *DeleteDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookSessionRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

type DeleteNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

func (r *DeleteNotebookSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotebookSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookSessionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNotebookSessionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotebookSessionResponseParams `json:"Response"`
}

func (r *DeleteNotebookSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScriptRequestParams struct {
	// 脚本id，其可以通过DescribeScripts接口提取
	ScriptIds []*string `json:"ScriptIds,omitnil" name:"ScriptIds"`
}

type DeleteScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本id，其可以通过DescribeScripts接口提取
	ScriptIds []*string `json:"ScriptIds,omitnil" name:"ScriptIds"`
}

func (r *DeleteScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScriptResponseParams struct {
	// 删除的脚本数量
	ScriptsAffected *int64 `json:"ScriptsAffected,omitnil" name:"ScriptsAffected"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteScriptResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScriptResponseParams `json:"Response"`
}

func (r *DeleteScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSparkAppRequestParams struct {
	// spark作业名
	AppName *string `json:"AppName,omitnil" name:"AppName"`
}

type DeleteSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	AppName *string `json:"AppName,omitnil" name:"AppName"`
}

func (r *DeleteSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSparkAppResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSparkAppResponseParams `json:"Response"`
}

func (r *DeleteSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// 需要删除的用户的Id
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的用户的Id
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersFromWorkGroupRequestParams struct {
	// 要删除的用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil" name:"AddInfo"`
}

type DeleteUsersFromWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil" name:"AddInfo"`
}

func (r *DeleteUsersFromWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersFromWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersFromWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersFromWorkGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUsersFromWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersFromWorkGroupResponseParams `json:"Response"`
}

func (r *DeleteUsersFromWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersFromWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkGroupRequestParams struct {
	// 要删除的工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil" name:"WorkGroupIds"`
}

type DeleteWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil" name:"WorkGroupIds"`
}

func (r *DeleteWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkGroupResponseParams `json:"Response"`
}

func (r *DeleteWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdvancedStoreLocationRequestParams struct {

}

type DescribeAdvancedStoreLocationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAdvancedStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvancedStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdvancedStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdvancedStoreLocationResponseParams struct {
	// 是否启用高级设置：0-否，1-是
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`

	// 查询结果保存cos路径
	StoreLocation *string `json:"StoreLocation,omitnil" name:"StoreLocation"`

	// 是否有托管存储权限
	HasLakeFs *bool `json:"HasLakeFs,omitnil" name:"HasLakeFs"`

	// 托管存储状态，HasLakeFs等于true时，该值才有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	LakeFsStatus *string `json:"LakeFsStatus,omitnil" name:"LakeFsStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAdvancedStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdvancedStoreLocationResponseParams `json:"Response"`
}

func (r *DescribeAdvancedStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvancedStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSDatabaseRequestParams struct {
	// 数据库名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 匹配规则
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`
}

type DescribeDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 匹配规则
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`
}

func (r *DescribeDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SchemaName")
	delete(f, "Pattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSDatabaseResponseParams struct {
	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 存储地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil" name:"Location"`

	// 数据对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSDatabaseResponseParams `json:"Response"`
}

func (r *DescribeDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSPartitionsRequestParams struct {
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 单个分区名称，精准匹配
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 多个分区名称，精准匹配
	PartitionNames []*string `json:"PartitionNames,omitnil" name:"PartitionNames"`

	// 多个分区字段的匹配，模糊匹配
	PartValues []*string `json:"PartValues,omitnil" name:"PartValues"`

	// 过滤SQL
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 最大分区数量
	MaxParts *int64 `json:"MaxParts,omitnil" name:"MaxParts"`

	// 翻页跳过数量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 页面数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 表达式
	Expression *string `json:"Expression,omitnil" name:"Expression"`
}

type DescribeDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 单个分区名称，精准匹配
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 多个分区名称，精准匹配
	PartitionNames []*string `json:"PartitionNames,omitnil" name:"PartitionNames"`

	// 多个分区字段的匹配，模糊匹配
	PartValues []*string `json:"PartValues,omitnil" name:"PartValues"`

	// 过滤SQL
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 最大分区数量
	MaxParts *int64 `json:"MaxParts,omitnil" name:"MaxParts"`

	// 翻页跳过数量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 页面数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 表达式
	Expression *string `json:"Expression,omitnil" name:"Expression"`
}

func (r *DescribeDMSPartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSPartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "TableName")
	delete(f, "SchemaName")
	delete(f, "Name")
	delete(f, "Values")
	delete(f, "PartitionNames")
	delete(f, "PartValues")
	delete(f, "Filter")
	delete(f, "MaxParts")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Expression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSPartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSPartitionsResponseParams struct {
	// 分区信息
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDMSPartitionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSPartitionsResponseParams `json:"Response"`
}

func (r *DescribeDMSPartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSPartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTableRequestParams struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbName")
	delete(f, "SchemaName")
	delete(f, "Name")
	delete(f, "Catalog")
	delete(f, "Keyword")
	delete(f, "Pattern")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTableResponseParams struct {
	// 基础对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Asset *Asset `json:"Asset,omitnil" name:"Asset"`

	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewOriginalText *string `json:"ViewOriginalText,omitnil" name:"ViewOriginalText"`

	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewExpandedText *string `json:"ViewExpandedText,omitnil" name:"ViewExpandedText"`

	// hive维护版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitnil" name:"Retention"`

	// 存储对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sds *DMSSds `json:"Sds,omitnil" name:"Sds"`

	// 分区列
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil" name:"PartitionKeys"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*DMSPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Schame名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// 生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeTime *int64 `json:"LifeTime,omitnil" name:"LifeTime"`

	// 最后访问时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAccessTime *string `json:"LastAccessTime,omitnil" name:"LastAccessTime"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUpdateTime *string `json:"DataUpdateTime,omitnil" name:"DataUpdateTime"`

	// 结构更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StructUpdateTime *string `json:"StructUpdateTime,omitnil" name:"StructUpdateTime"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*DMSColumn `json:"Columns,omitnil" name:"Columns"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSTableResponseParams `json:"Response"`
}

func (r *DescribeDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTablesRequestParams struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 筛选参数：更新开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 筛选参数：更新结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序字段：create_time：创建时间
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序字段：true：升序（默认），false：降序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`
}

type DescribeDMSTablesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 筛选参数：更新开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 筛选参数：更新结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序字段：create_time：创建时间
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序字段：true：升序（默认），false：降序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`
}

func (r *DescribeDMSTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbName")
	delete(f, "SchemaName")
	delete(f, "Name")
	delete(f, "Catalog")
	delete(f, "Keyword")
	delete(f, "Pattern")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Sort")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTablesResponseParams struct {
	// DMS元数据列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableList []*DMSTableInfo `json:"TableList,omitnil" name:"TableList"`

	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDMSTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSTablesResponseParams `json:"Response"`
}

func (r *DescribeDMSTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineEventsRequestParams struct {
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 返回数量，默认为10，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeDataEngineEventsRequest struct {
	*tchttp.BaseRequest
	
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 返回数量，默认为10，最大为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeDataEngineEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEngineEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineEventsResponseParams struct {
	// 事件详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*HouseEventsInfo `json:"Events,omitnil" name:"Events"`

	// 分页号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPages *int64 `json:"TotalPages,omitnil" name:"TotalPages"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDataEngineEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEngineEventsResponseParams `json:"Response"`
}

func (r *DescribeDataEngineEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineImageVersionsRequestParams struct {
	// 引擎类型：SQL、SparkBatch
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`
}

type DescribeDataEngineImageVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型：SQL、SparkBatch
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`
}

func (r *DescribeDataEngineImageVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineImageVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEngineImageVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineImageVersionsResponseParams struct {
	// 集群大版本镜像信息列表
	ImageParentVersions []*DataEngineImageVersion `json:"ImageParentVersions,omitnil" name:"ImageParentVersions"`

	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDataEngineImageVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEngineImageVersionsResponseParams `json:"Response"`
}

func (r *DescribeDataEngineImageVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineImageVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginePythonSparkImagesRequestParams struct {
	// 集群镜像小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil" name:"ChildImageVersionId"`
}

type DescribeDataEnginePythonSparkImagesRequest struct {
	*tchttp.BaseRequest
	
	// 集群镜像小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil" name:"ChildImageVersionId"`
}

func (r *DescribeDataEnginePythonSparkImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginePythonSparkImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChildImageVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEnginePythonSparkImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginePythonSparkImagesResponseParams struct {
	// PYSPARK镜像信息列表
	PythonSparkImages []*PythonSparkImage `json:"PythonSparkImages,omitnil" name:"PythonSparkImages"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDataEnginePythonSparkImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEnginePythonSparkImagesResponseParams `json:"Response"`
}

func (r *DescribeDataEnginePythonSparkImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginePythonSparkImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineRequestParams struct {
	// House名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

type DescribeDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// House名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

func (r *DescribeDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineResponseParams struct {
	// 数据引擎详细信息
	DataEngine *DataEngineInfo `json:"DataEngine,omitnil" name:"DataEngine"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEngineResponseParams `json:"Response"`
}

func (r *DescribeDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginesRequestParams struct {
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤类型，支持如下的过滤类型，传参Name应为以下其中一个, data-engine-name - String（数据引擎名称）：engine-type - String（引擎类型：spark：spark 引擎，presto：presto引擎），state - String (数据引擎状态 -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中) ， mode - String（计费模式 0共享模式 1按量计费 2包年包月） ， create-time - String（创建时间，10位时间戳） message - String （描述信息），cluster-type - String (集群资源类型 spark_private/presto_private/presto_cu/spark_cu/kyuubi_cu)，engine-id - String（数据引擎ID），key-word - String（数据引擎名称或集群资源类型或描述信息模糊搜索），engine-exec-type - String（引擎执行任务类型，SQL/BATCH），engine-network-id - String（引擎网络Id）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 已废弃，请使用DatasourceConnectionNameSet
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 是否不返回共享引擎，true不返回共享引擎，false可以返回共享引擎
	ExcludePublicEngine *bool `json:"ExcludePublicEngine,omitnil" name:"ExcludePublicEngine"`

	// 参数应该为引擎权限类型，有效类型："USE", "MODIFY", "OPERATE", "MONITOR", "DELETE"
	AccessTypes []*string `json:"AccessTypes,omitnil" name:"AccessTypes"`

	// 引擎执行任务类型，有效值：SQL/BATCH，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitnil" name:"EngineExecType"`

	// 引擎类型，有效值：spark/presto/kyuubi，为空时默认获取非kyuubi引擎（网关引擎）
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 网络配置列表，若传入该参数，则返回网络配置关联的计算引擎
	DatasourceConnectionNameSet []*string `json:"DatasourceConnectionNameSet,omitnil" name:"DatasourceConnectionNameSet"`

	// 引擎版本，有效值：Native/SuperSQL，为空时默认获取SuperSQL引擎
	EngineGeneration *string `json:"EngineGeneration,omitnil" name:"EngineGeneration"`

	// 引擎类型，支持：SparkSQL、SparkBatch、PrestoSQL、Kyuubi
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil" name:"EngineTypeDetail"`
}

type DescribeDataEnginesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤类型，支持如下的过滤类型，传参Name应为以下其中一个, data-engine-name - String（数据引擎名称）：engine-type - String（引擎类型：spark：spark 引擎，presto：presto引擎），state - String (数据引擎状态 -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中) ， mode - String（计费模式 0共享模式 1按量计费 2包年包月） ， create-time - String（创建时间，10位时间戳） message - String （描述信息），cluster-type - String (集群资源类型 spark_private/presto_private/presto_cu/spark_cu/kyuubi_cu)，engine-id - String（数据引擎ID），key-word - String（数据引擎名称或集群资源类型或描述信息模糊搜索），engine-exec-type - String（引擎执行任务类型，SQL/BATCH），engine-network-id - String（引擎网络Id）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 已废弃，请使用DatasourceConnectionNameSet
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 是否不返回共享引擎，true不返回共享引擎，false可以返回共享引擎
	ExcludePublicEngine *bool `json:"ExcludePublicEngine,omitnil" name:"ExcludePublicEngine"`

	// 参数应该为引擎权限类型，有效类型："USE", "MODIFY", "OPERATE", "MONITOR", "DELETE"
	AccessTypes []*string `json:"AccessTypes,omitnil" name:"AccessTypes"`

	// 引擎执行任务类型，有效值：SQL/BATCH，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitnil" name:"EngineExecType"`

	// 引擎类型，有效值：spark/presto/kyuubi，为空时默认获取非kyuubi引擎（网关引擎）
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 网络配置列表，若传入该参数，则返回网络配置关联的计算引擎
	DatasourceConnectionNameSet []*string `json:"DatasourceConnectionNameSet,omitnil" name:"DatasourceConnectionNameSet"`

	// 引擎版本，有效值：Native/SuperSQL，为空时默认获取SuperSQL引擎
	EngineGeneration *string `json:"EngineGeneration,omitnil" name:"EngineGeneration"`

	// 引擎类型，支持：SparkSQL、SparkBatch、PrestoSQL、Kyuubi
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil" name:"EngineTypeDetail"`
}

func (r *DescribeDataEnginesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "DatasourceConnectionName")
	delete(f, "ExcludePublicEngine")
	delete(f, "AccessTypes")
	delete(f, "EngineExecType")
	delete(f, "EngineType")
	delete(f, "DatasourceConnectionNameSet")
	delete(f, "EngineGeneration")
	delete(f, "EngineTypeDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEnginesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginesResponseParams struct {
	// 数据引擎列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngines []*DataEngineInfo `json:"DataEngines,omitnil" name:"DataEngines"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDataEnginesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEnginesResponseParams `json:"Response"`
}

func (r *DescribeDataEnginesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitnil" name:"KeyWord"`

	// 数据源唯名称，该名称可以通过DescribeDatasourceConnection接口查询到。默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 排序字段，CreateTime：创建时间，Name：数据库名称
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序类型：false：降序（默认）、true：升序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitnil" name:"KeyWord"`

	// 数据源唯名称，该名称可以通过DescribeDatasourceConnection接口查询到。默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 排序字段，CreateTime：创建时间，Name：数据库名称
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序类型：false：降序（默认）、true：升序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "KeyWord")
	delete(f, "DatasourceConnectionName")
	delete(f, "Sort")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// 数据库对象列表。
	DatabaseList []*DatabaseResponseInfo `json:"DatabaseList,omitnil" name:"DatabaseList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceConnectionRequestParams struct {
	// 连接ID列表，指定要查询的连接ID
	DatasourceConnectionIds []*string `json:"DatasourceConnectionIds,omitnil" name:"DatasourceConnectionIds"`

	// 过滤条件，当前支持的过滤键为：DatasourceConnectionName（数据源连接名）。
	// DatasourceConnectionType   （数据源连接连接类型）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time（默认，创建时间）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为desc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 筛选字段：起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 筛选字段：截止时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 连接名称列表，指定要查询的连接名称
	DatasourceConnectionNames []*string `json:"DatasourceConnectionNames,omitnil" name:"DatasourceConnectionNames"`

	// 连接类型，支持Mysql/HiveCos/Kafka/DataLakeCatalog
	DatasourceConnectionTypes []*string `json:"DatasourceConnectionTypes,omitnil" name:"DatasourceConnectionTypes"`

	// 返回指定hive版本的数据源，该参数指定后，会过滤掉该参数指定版本以外的hive数据源，非hive数据源正常返回
	HiveVersion []*string `json:"HiveVersion,omitnil" name:"HiveVersion"`
}

type DescribeDatasourceConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接ID列表，指定要查询的连接ID
	DatasourceConnectionIds []*string `json:"DatasourceConnectionIds,omitnil" name:"DatasourceConnectionIds"`

	// 过滤条件，当前支持的过滤键为：DatasourceConnectionName（数据源连接名）。
	// DatasourceConnectionType   （数据源连接连接类型）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time（默认，创建时间）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为desc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 筛选字段：起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 筛选字段：截止时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 连接名称列表，指定要查询的连接名称
	DatasourceConnectionNames []*string `json:"DatasourceConnectionNames,omitnil" name:"DatasourceConnectionNames"`

	// 连接类型，支持Mysql/HiveCos/Kafka/DataLakeCatalog
	DatasourceConnectionTypes []*string `json:"DatasourceConnectionTypes,omitnil" name:"DatasourceConnectionTypes"`

	// 返回指定hive版本的数据源，该参数指定后，会过滤掉该参数指定版本以外的hive数据源，非hive数据源正常返回
	HiveVersion []*string `json:"HiveVersion,omitnil" name:"HiveVersion"`
}

func (r *DescribeDatasourceConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceConnectionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatasourceConnectionNames")
	delete(f, "DatasourceConnectionTypes")
	delete(f, "HiveVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasourceConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceConnectionResponseParams struct {
	// 数据连接总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 数据连接对象集合
	ConnectionSet []*DatasourceConnectionInfo `json:"ConnectionSet,omitnil" name:"ConnectionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatasourceConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasourceConnectionResponseParams `json:"Response"`
}

func (r *DescribeDatasourceConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineUsageInfoRequestParams struct {
	// 数据引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

type DescribeEngineUsageInfoRequest struct {
	*tchttp.BaseRequest
	
	// 数据引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

func (r *DescribeEngineUsageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineUsageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEngineUsageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineUsageInfoResponseParams struct {
	// 集群总规格
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 已占用集群规格
	Used *int64 `json:"Used,omitnil" name:"Used"`

	// 剩余集群规格
	Available *int64 `json:"Available,omitnil" name:"Available"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEngineUsageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEngineUsageInfoResponseParams `json:"Response"`
}

func (r *DescribeEngineUsageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineUsageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForbiddenTableProRequestParams struct {

}

type DescribeForbiddenTableProRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeForbiddenTableProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForbiddenTableProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForbiddenTableProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForbiddenTableProResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeForbiddenTableProResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForbiddenTableProResponseParams `json:"Response"`
}

func (r *DescribeForbiddenTableProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForbiddenTableProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsDirSummaryRequestParams struct {

}

type DescribeLakeFsDirSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLakeFsDirSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsDirSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsDirSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsDirSummaryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLakeFsDirSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsDirSummaryResponseParams `json:"Response"`
}

func (r *DescribeLakeFsDirSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsDirSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsInfoRequestParams struct {

}

type DescribeLakeFsInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLakeFsInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLakeFsInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsInfoResponseParams `json:"Response"`
}

func (r *DescribeLakeFsInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsTaskResultRequestParams struct {
	// 需要访问的任务结果路径
	FsPath *string `json:"FsPath,omitnil" name:"FsPath"`
}

type DescribeLakeFsTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 需要访问的任务结果路径
	FsPath *string `json:"FsPath,omitnil" name:"FsPath"`
}

func (r *DescribeLakeFsTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FsPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsTaskResultResponseParams struct {
	// 路径的访问实例
	AccessToken *LakeFileSystemToken `json:"AccessToken,omitnil" name:"AccessToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLakeFsTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsTaskResultResponseParams `json:"Response"`
}

func (r *DescribeLakeFsTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionLogRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeNotebookSessionLogRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeNotebookSessionLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionLogResponseParams struct {
	// 日志信息，默认获取最新的200条
	Logs []*string `json:"Logs,omitnil" name:"Logs"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNotebookSessionLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionLogResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

type DescribeNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`
}

func (r *DescribeNotebookSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionResponseParams struct {
	// Session详情信息
	Session *NotebookSessionInfo `json:"Session,omitnil" name:"Session"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNotebookSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil" name:"StatementId"`

	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil" name:"StatementId"`

	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeNotebookSessionStatementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "StatementId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementResponseParams struct {
	// Session Statement详情
	NotebookSessionStatement *NotebookSessionStatementInfo `json:"NotebookSessionStatement,omitnil" name:"NotebookSessionStatement"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNotebookSessionStatementResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionStatementResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionStatementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementSqlResultRequestParams struct {
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *uint64 `json:"MaxResults,omitnil" name:"MaxResults"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`
}

type DescribeNotebookSessionStatementSqlResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *uint64 `json:"MaxResults,omitnil" name:"MaxResults"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`
}

func (r *DescribeNotebookSessionStatementSqlResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementSqlResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionStatementSqlResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementSqlResultResponseParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 结果数据
	ResultSet *string `json:"ResultSet,omitnil" name:"ResultSet"`

	// schema
	ResultSchema []*Column `json:"ResultSchema,omitnil" name:"ResultSchema"`

	// 分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// 存储结果地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputPath *string `json:"OutputPath,omitnil" name:"OutputPath"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNotebookSessionStatementSqlResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionStatementSqlResultResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionStatementSqlResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementSqlResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementsRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 批任务id
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

type DescribeNotebookSessionStatementsRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 批任务id
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

func (r *DescribeNotebookSessionStatementsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionStatementsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementsResponseParams struct {
	// Session Statement详情
	NotebookSessionStatements *NotebookSessionStatementBatchInformation `json:"NotebookSessionStatements,omitnil" name:"NotebookSessionStatements"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNotebookSessionStatementsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionStatementsResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionStatementsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionsRequestParams struct {
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State []*string `json:"State,omitnil" name:"State"`

	// 排序字段（默认按创建时间）
	SortFields []*string `json:"SortFields,omitnil" name:"SortFields"`

	// 排序字段：true：升序、false：降序（默认）
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// 分页参数，默认10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeNotebookSessionsRequest struct {
	*tchttp.BaseRequest
	
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State []*string `json:"State,omitnil" name:"State"`

	// 排序字段（默认按创建时间）
	SortFields []*string `json:"SortFields,omitnil" name:"SortFields"`

	// 排序字段：true：升序、false：降序（默认）
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// 分页参数，默认10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeNotebookSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "State")
	delete(f, "SortFields")
	delete(f, "Asc")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionsResponseParams struct {
	// session总数量
	TotalElements *int64 `json:"TotalElements,omitnil" name:"TotalElements"`

	// 总页数
	TotalPages *int64 `json:"TotalPages,omitnil" name:"TotalPages"`

	// 当前页码
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 当前页数量
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// session列表信息
	Sessions []*NotebookSessions `json:"Sessions,omitnil" name:"Sessions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNotebookSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionsResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResultDownloadRequestParams struct {
	// 查询任务Id
	DownloadId *string `json:"DownloadId,omitnil" name:"DownloadId"`
}

type DescribeResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 查询任务Id
	DownloadId *string `json:"DownloadId,omitnil" name:"DownloadId"`
}

func (r *DescribeResultDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResultDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DownloadId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResultDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResultDownloadResponseParams struct {
	// 下载文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 任务状态 init | queue | format | compress | success|  timeout | error
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务异常原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 临时AK
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitnil" name:"SecretId"`

	// 临时SK
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 临时Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeResultDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResultDownloadResponseParams `json:"Response"`
}

func (r *DescribeResultDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResultDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScriptsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序，默认asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeScriptsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序，默认asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScriptsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScriptsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScriptsResponseParams struct {
	// Script列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scripts []*Script `json:"Scripts,omitnil" name:"Scripts"`

	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScriptsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScriptsResponseParams `json:"Response"`
}

func (r *DescribeScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobRequestParams struct {
	// spark作业Id，与JobName同时存在时，JobName无效，JobId与JobName至少存在一个
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitnil" name:"JobName"`
}

type DescribeSparkAppJobRequest struct {
	*tchttp.BaseRequest
	
	// spark作业Id，与JobName同时存在时，JobName无效，JobId与JobName至少存在一个
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitnil" name:"JobName"`
}

func (r *DescribeSparkAppJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobResponseParams struct {
	// spark作业详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *SparkJobInfo `json:"Job,omitnil" name:"Job"`

	// 查询的spark作业是否存在
	IsExists *bool `json:"IsExists,omitnil" name:"IsExists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSparkAppJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppJobResponseParams `json:"Response"`
}

func (r *DescribeSparkAppJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobsRequestParams struct {
	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一:spark-job-name（作业名称），spark-job-id（作业id），spark-app-type（作业类型，1：批任务，2：流任务，4：SQL作业），user-name（创建人），key-word（作业名称或ID关键词模糊搜索）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询列表偏移量, 默认值0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询列表限制数量, 默认值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSparkAppJobsRequest struct {
	*tchttp.BaseRequest
	
	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一:spark-job-name（作业名称），spark-job-id（作业id），spark-app-type（作业类型，1：批任务，2：流任务，4：SQL作业），user-name（创建人），key-word（作业名称或ID关键词模糊搜索）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询列表偏移量, 默认值0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询列表限制数量, 默认值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSparkAppJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobsResponseParams struct {
	// spark作业列表详情
	SparkAppJobs []*SparkJobInfo `json:"SparkAppJobs,omitnil" name:"SparkAppJobs"`

	// spark作业总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSparkAppJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppJobsResponseParams `json:"Response"`
}

func (r *DescribeSparkAppJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppTasksRequestParams struct {
	// spark作业Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询Limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 执行实例id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 按照该参数过滤,支持task-state
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeSparkAppTasksRequest struct {
	*tchttp.BaseRequest
	
	// spark作业Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询Limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 执行实例id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 按照该参数过滤,支持task-state
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeSparkAppTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TaskId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppTasksResponseParams struct {
	// 任务结果（该字段已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks *TaskResponseInfo `json:"Tasks,omitnil" name:"Tasks"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppTasks []*TaskResponseInfo `json:"SparkAppTasks,omitnil" name:"SparkAppTasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSparkAppTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppTasksResponseParams `json:"Response"`
}

func (r *DescribeSparkAppTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSQLRequestParams struct {
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

type DescribeSparkSessionBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

func (r *DescribeSparkSessionBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkSessionBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSQLResponseParams struct {
	// 状态：0：运行中、1：成功、2：失败、3：取消、4：超时；
	State *uint64 `json:"State,omitnil" name:"State"`

	// SQL子任务列表，仅展示运行完成的子任务，若某个任务运行失败，后续其它子任务不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*BatchSqlTask `json:"Tasks,omitnil" name:"Tasks"`

	// 非sql运行的异常事件信息，包含资源创建失败、调度异常，JOB超时等，正常运行下该Event值为空
	Event *string `json:"Event,omitnil" name:"Event"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSparkSessionBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkSessionBatchSQLResponseParams `json:"Response"`
}

func (r *DescribeSparkSessionBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSqlLogRequestParams struct {
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

type DescribeSparkSessionBatchSqlLogRequest struct {
	*tchttp.BaseRequest
	
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

func (r *DescribeSparkSessionBatchSqlLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSqlLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkSessionBatchSqlLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSqlLogResponseParams struct {
	// 状态：0：运行中、1：成功、2：失败、3：取消、4：超时；
	State *uint64 `json:"State,omitnil" name:"State"`

	// 日志信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSet []*SparkSessionBatchLog `json:"LogSet,omitnil" name:"LogSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSparkSessionBatchSqlLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkSessionBatchSqlLogResponseParams `json:"Response"`
}

func (r *DescribeSparkSessionBatchSqlLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSqlLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStoreLocationRequestParams struct {

}

type DescribeStoreLocationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStoreLocationResponseParams struct {
	// 返回用户设置的结果存储位置路径，如果未设置则返回空字符串：""
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoreLocation *string `json:"StoreLocation,omitnil" name:"StoreLocation"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStoreLocationResponseParams `json:"Response"`
}

func (r *DescribeStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableRequestParams struct {
	// 查询对象表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 查询表所在的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest
	
	// 查询对象表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 查询表所在的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

func (r *DescribeTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "DatabaseName")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableResponseParams struct {
	// 数据表对象
	Table *TableResponseInfo `json:"Table,omitnil" name:"Table"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableResponseParams `json:"Response"`
}

func (r *DescribeTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesNameRequestParams struct {
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil" name:"TableFormat"`
}

type DescribeTablesNameRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil" name:"TableFormat"`
}

func (r *DescribeTablesNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "TableType")
	delete(f, "TableFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesNameResponseParams struct {
	// 数据表名称对象列表。
	TableNameList []*string `json:"TableNameList,omitnil" name:"TableNameList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTablesNameResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesNameResponseParams `json:"Response"`
}

func (r *DescribeTablesNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil" name:"TableFormat"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil" name:"TableFormat"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "TableType")
	delete(f, "TableFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// 数据表对象列表。
	TableList []*TableResponseInfo `json:"TableList,omitnil" name:"TableList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultRequestParams struct {
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *int64 `json:"MaxResults,omitnil" name:"MaxResults"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *int64 `json:"MaxResults,omitnil" name:"MaxResults"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// 查询的任务信息，返回为空表示输入任务ID对应的任务不存在。只有当任务状态为成功（2）的时候，才会返回任务的结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfo *TaskResultInfo `json:"TaskInfo,omitnil" name:"TaskInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResultResponseParams `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	// task-kind - string （任务类型过滤）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time（创建时间，默认）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	// task-kind - string （任务类型过滤）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time（创建时间，默认）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 任务对象列表。
	TaskList []*TaskResponseInfo `json:"TaskList,omitnil" name:"TaskList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务概览信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TasksOverview *TasksOverview `json:"TasksOverview,omitnil" name:"TasksOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpdatableDataEnginesRequestParams struct {
	// 引擎配置操作命令，UpdateSparkSQLLakefsPath 更新托管表路径，UpdateSparkSQLResultPath 更新结果桶路径
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil" name:"DataEngineConfigCommand"`
}

type DescribeUpdatableDataEnginesRequest struct {
	*tchttp.BaseRequest
	
	// 引擎配置操作命令，UpdateSparkSQLLakefsPath 更新托管表路径，UpdateSparkSQLResultPath 更新结果桶路径
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil" name:"DataEngineConfigCommand"`
}

func (r *DescribeUpdatableDataEnginesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpdatableDataEnginesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineConfigCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpdatableDataEnginesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpdatableDataEnginesResponseParams struct {
	// 集群基础信息
	DataEngineBasicInfos []*DataEngineBasicInfo `json:"DataEngineBasicInfos,omitnil" name:"DataEngineBasicInfos"`

	// 集群个数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUpdatableDataEnginesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpdatableDataEnginesResponseParams `json:"Response"`
}

func (r *DescribeUpdatableDataEnginesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpdatableDataEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDataEngineConfigRequestParams struct {
	// 排序方式，desc表示倒序，asc表示正序
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,每种过滤参数支持的过滤值不超过5个。
	// app-id - String - （appid过滤）
	// engine-id - String - （引擎ID过滤）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeUserDataEngineConfigRequest struct {
	*tchttp.BaseRequest
	
	// 排序方式，desc表示倒序，asc表示正序
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,每种过滤参数支持的过滤值不超过5个。
	// app-id - String - （appid过滤）
	// engine-id - String - （引擎ID过滤）
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeUserDataEngineConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDataEngineConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDataEngineConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDataEngineConfigResponseParams struct {
	// 用户引擎自定义配置项列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineConfigInstanceInfos []*DataEngineConfigInstanceInfo `json:"DataEngineConfigInstanceInfos,omitnil" name:"DataEngineConfigInstanceInfos"`

	// 配置项总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserDataEngineConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDataEngineConfigResponseParams `json:"Response"`
}

func (r *DescribeUserDataEngineConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDataEngineConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoRequestParams struct {
	// 用户Id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 查询的信息类型，Group：工作组 DataAuth：数据权限 EngineAuth:引擎权限
	Type *string `json:"Type,omitnil" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为Group时，支持Key为workgroup-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为Group时，支持create-time、group-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 查询的信息类型，Group：工作组 DataAuth：数据权限 EngineAuth:引擎权限
	Type *string `json:"Type,omitnil" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为Group时，支持Key为workgroup-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为Group时，支持create-time、group-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "Type")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoResponseParams struct {
	// 用户详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserInfo *UserDetailInfo `json:"UserInfo,omitnil" name:"UserInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInfoResponseParams `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRolesRequestParams struct {
	// 列举的数量限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列举的偏移位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 按照arn模糊列举
	Fuzzy *string `json:"Fuzzy,omitnil" name:"Fuzzy"`

	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`
}

type DescribeUserRolesRequest struct {
	*tchttp.BaseRequest
	
	// 列举的数量限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列举的偏移位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 按照arn模糊列举
	Fuzzy *string `json:"Fuzzy,omitnil" name:"Fuzzy"`

	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`
}

func (r *DescribeUserRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Fuzzy")
	delete(f, "SortBy")
	delete(f, "Sorting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRolesResponseParams struct {
	// 符合列举条件的总数量
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 用户角色信息
	UserRoles []*UserRole `json:"UserRoles,omitnil" name:"UserRoles"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserRolesResponseParams `json:"Response"`
}

func (r *DescribeUserRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserTypeRequestParams struct {
	// 用户ID（UIN），如果不填默认为调用方的子UIN
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type DescribeUserTypeRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID（UIN），如果不填默认为调用方的子UIN
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

func (r *DescribeUserTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserTypeResponseParams struct {
	// 用户类型。ADMIN：管理员 COMMON：普通用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserTypeResponseParams `json:"Response"`
}

func (r *DescribeUserTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersRequestParams struct {
	// 指定查询的子用户uin，用户需要通过CreateUser接口创建。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 过滤条件，支持如下字段类型，user-type：根据用户类型过滤。user-keyword：根据用户名称过滤
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 指定查询的子用户uin，用户需要通过CreateUser接口创建。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 过滤条件，支持如下字段类型，user-type：根据用户类型过滤。user-keyword：根据用户名称过滤
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// 查询到的用户总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 查询到的授权用户信息集合
	UserSet []*UserInfo `json:"UserSet,omitnil" name:"UserSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersResponseParams `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeViewsRequestParams struct {
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 数据库所属的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 排序字段
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序规则，true:升序；false:降序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// 按视图更新时间筛选，开始时间，如2021-11-11 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 按视图更新时间筛选，结束时间，如2021-11-12 00:00:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 数据库所属的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 排序字段
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 排序规则，true:升序；false:降序
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// 按视图更新时间筛选，开始时间，如2021-11-11 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 按视图更新时间筛选，结束时间，如2021-11-12 00:00:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeViewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeViewsResponseParams struct {
	// 视图对象列表。
	ViewList []*ViewResponseInfo `json:"ViewList,omitnil" name:"ViewList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeViewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeViewsResponseParams `json:"Response"`
}

func (r *DescribeViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupInfoRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 查询信息类型：User：用户信息 DataAuth：数据权限 EngineAuth：引擎权限
	Type *string `json:"Type,omitnil" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为User时，支持Key为user-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为User时，支持create-time、user-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeWorkGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 查询信息类型：User：用户信息 DataAuth：数据权限 EngineAuth：引擎权限
	Type *string `json:"Type,omitnil" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为User时，支持Key为user-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为User时，支持create-time、user-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeWorkGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "Type")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupInfoResponseParams struct {
	// 工作组详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupInfo *WorkGroupDetailInfo `json:"WorkGroupInfo,omitnil" name:"WorkGroupInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkGroupInfoResponseParams `json:"Response"`
}

func (r *DescribeWorkGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupsRequestParams struct {
	// 查询的工作组Id，不填或填0表示不过滤。
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 过滤条件，当前仅支持按照工作组名称进行模糊搜索。Key为workgroup-name
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`
}

type DescribeWorkGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的工作组Id，不填或填0表示不过滤。
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 过滤条件，当前仅支持按照工作组名称进行模糊搜索。Key为workgroup-name
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil" name:"Sorting"`
}

func (r *DescribeWorkGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupsResponseParams struct {
	// 工作组总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 工作组信息集合
	WorkGroupSet []*WorkGroupInfo `json:"WorkGroupSet,omitnil" name:"WorkGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWorkGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkGroupsResponseParams `json:"Response"`
}

func (r *DescribeWorkGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyRequestParams struct {
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

func (r *DetachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachUserPolicyResponseParams `json:"Response"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachWorkGroupPolicyRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

type DetachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`
}

func (r *DetachWorkGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachWorkGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachWorkGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachWorkGroupPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetachWorkGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachWorkGroupPolicyResponseParams `json:"Response"`
}

func (r *DetachWorkGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachWorkGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSDatabaseRequestParams struct {
	// 数据库名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil" name:"DeleteData"`

	// 是否级联删除
	Cascade *bool `json:"Cascade,omitnil" name:"Cascade"`
}

type DropDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil" name:"DeleteData"`

	// 是否级联删除
	Cascade *bool `json:"Cascade,omitnil" name:"Cascade"`
}

func (r *DropDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DeleteData")
	delete(f, "Cascade")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSDatabaseResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DropDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DropDMSDatabaseResponseParams `json:"Response"`
}

func (r *DropDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSPartitionsRequestParams struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 数据表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 分区名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 单个分区名称
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 是否删除分区数据
	DeleteData *bool `json:"DeleteData,omitnil" name:"DeleteData"`
}

type DropDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 数据表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 分区名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 单个分区名称
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 是否删除分区数据
	DeleteData *bool `json:"DeleteData,omitnil" name:"DeleteData"`
}

func (r *DropDMSPartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSPartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "SchemaName")
	delete(f, "TableName")
	delete(f, "Name")
	delete(f, "Values")
	delete(f, "DeleteData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropDMSPartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSPartitionsResponseParams struct {
	// 状态
	Status *bool `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DropDMSPartitionsResponse struct {
	*tchttp.BaseResponse
	Response *DropDMSPartitionsResponseParams `json:"Response"`
}

func (r *DropDMSPartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSPartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSTableRequestParams struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil" name:"DeleteData"`

	// 环境属性
	EnvProps *KVPair `json:"EnvProps,omitnil" name:"EnvProps"`
}

type DropDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 表名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil" name:"DeleteData"`

	// 环境属性
	EnvProps *KVPair `json:"EnvProps,omitnil" name:"EnvProps"`
}

func (r *DropDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbName")
	delete(f, "Name")
	delete(f, "DeleteData")
	delete(f, "EnvProps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSTableResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DropDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *DropDMSTableResponseParams `json:"Response"`
}

func (r *DropDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ElasticsearchInfo struct {
	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 密码，需要base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 数据源的VPC和子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *DatasourceConnectionLocation `json:"Location,omitnil" name:"Location"`

	// 默认数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 访问Elasticsearch的ip、端口信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInfo []*IpPortPair `json:"ServiceInfo,omitnil" name:"ServiceInfo"`
}

type Execution struct {
	// 自动生成SQL语句。
	SQL *string `json:"SQL,omitnil" name:"SQL"`
}

type FavorInfo struct {
	// 优先事项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// Catalog名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// DataBase名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataBase *string `json:"DataBase,omitnil" name:"DataBase"`

	// Table名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitnil" name:"Table"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑或（OR）关系。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

// Predefined struct for user
type GenerateCreateMangedTableSqlRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`

	// V2 upsert表 upsert键
	UpsertKeys []*string `json:"UpsertKeys,omitnil" name:"UpsertKeys"`
}

type GenerateCreateMangedTableSqlRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`

	// V2 upsert表 upsert键
	UpsertKeys []*string `json:"UpsertKeys,omitnil" name:"UpsertKeys"`
}

func (r *GenerateCreateMangedTableSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCreateMangedTableSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "Partitions")
	delete(f, "Properties")
	delete(f, "UpsertKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateCreateMangedTableSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateCreateMangedTableSqlResponseParams struct {
	// 创建托管存储内表sql语句描述
	Execution *Execution `json:"Execution,omitnil" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GenerateCreateMangedTableSqlResponse struct {
	*tchttp.BaseResponse
	Response *GenerateCreateMangedTableSqlResponseParams `json:"Response"`
}

func (r *GenerateCreateMangedTableSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCreateMangedTableSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOptimizerPolicyRequestParams struct {
	// 策略描述
	SmartPolicy *SmartPolicy `json:"SmartPolicy,omitnil" name:"SmartPolicy"`
}

type GetOptimizerPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略描述
	SmartPolicy *SmartPolicy `json:"SmartPolicy,omitnil" name:"SmartPolicy"`
}

func (r *GetOptimizerPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOptimizerPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SmartPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOptimizerPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOptimizerPolicyResponseParams struct {
	// 智能优化策略
	SmartOptimizerPolicy *SmartOptimizerPolicy `json:"SmartOptimizerPolicy,omitnil" name:"SmartOptimizerPolicy"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetOptimizerPolicyResponse struct {
	*tchttp.BaseResponse
	Response *GetOptimizerPolicyResponseParams `json:"Response"`
}

func (r *GetOptimizerPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOptimizerPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HiveInfo struct {
	// hive metastore的地址
	MetaStoreUrl *string `json:"MetaStoreUrl,omitnil" name:"MetaStoreUrl"`

	// hive数据源类型，代表数据储存的位置，COS或者HDFS
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据源所在的私有网络信息
	Location *DatasourceConnectionLocation `json:"Location,omitnil" name:"Location"`

	// 如果类型为HDFS，需要传一个用户名
	User *string `json:"User,omitnil" name:"User"`

	// 如果类型为HDFS，需要选择是否高可用
	HighAvailability *bool `json:"HighAvailability,omitnil" name:"HighAvailability"`

	// 如果类型为COS，需要填写COS桶连接
	BucketUrl *string `json:"BucketUrl,omitnil" name:"BucketUrl"`

	// json字符串。如果类型为HDFS，需要填写该字段
	HdfsProperties *string `json:"HdfsProperties,omitnil" name:"HdfsProperties"`

	// Hive的元数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mysql *MysqlInfo `json:"Mysql,omitnil" name:"Mysql"`

	// emr集群Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// emr集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// EMR集群中hive组件的版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	HiveVersion *string `json:"HiveVersion,omitnil" name:"HiveVersion"`

	// Kerberos详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KerberosInfo *KerberosInfo `json:"KerberosInfo,omitnil" name:"KerberosInfo"`

	// 是否开启Kerberos
	// 注意：此字段可能返回 null，表示取不到有效值。
	KerberosEnable *bool `json:"KerberosEnable,omitnil" name:"KerberosEnable"`
}

type HouseEventsInfo struct {
	// 事件时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time []*string `json:"Time,omitnil" name:"Time"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventsAction []*string `json:"EventsAction,omitnil" name:"EventsAction"`

	// 集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterInfo []*string `json:"ClusterInfo,omitnil" name:"ClusterInfo"`
}

type IpPortPair struct {
	// ip信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 端口信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type JobLogResult struct {
	// 日志时间戳，毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitnil" name:"Time"`

	// 日志topic id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 日志topic name
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 日志内容，json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitnil" name:"LogJson"`

	// 日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgLogId *string `json:"PkgLogId,omitnil" name:"PkgLogId"`
}

type KVPair struct {
	// 配置的key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 配置的value值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type KafkaInfo struct {
	// kakfa实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// kakfa数据源的网络信息
	Location *DatasourceConnectionLocation `json:"Location,omitnil" name:"Location"`
}

type KerberosInfo struct {
	// Krb5Conf文件值
	Krb5Conf *string `json:"Krb5Conf,omitnil" name:"Krb5Conf"`

	// KeyTab文件值
	KeyTab *string `json:"KeyTab,omitnil" name:"KeyTab"`

	// 服务主体
	ServicePrincipal *string `json:"ServicePrincipal,omitnil" name:"ServicePrincipal"`
}

type LakeFileSystemToken struct {
	// Token使用的临时秘钥的ID
	SecretId *string `json:"SecretId,omitnil" name:"SecretId"`

	// Token使用的临时秘钥
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// Token信息
	Token *string `json:"Token,omitnil" name:"Token"`

	// 过期时间
	ExpiredTime *int64 `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 颁布时间
	IssueTime *int64 `json:"IssueTime,omitnil" name:"IssueTime"`
}

// Predefined struct for user
type ListTaskJobLogDetailRequestParams struct {
	// 列表返回的Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 下一次分页参数，第一次传空
	Context *string `json:"Context,omitnil" name:"Context"`

	// 最近1000条日志是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// SparkSQL任务唯一ID
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

type ListTaskJobLogDetailRequest struct {
	*tchttp.BaseRequest
	
	// 列表返回的Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 下一次分页参数，第一次传空
	Context *string `json:"Context,omitnil" name:"Context"`

	// 最近1000条日志是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitnil" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// SparkSQL任务唯一ID
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

func (r *ListTaskJobLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskJobLogDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Asc")
	delete(f, "Filters")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskJobLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskJobLogDetailResponseParams struct {
	// 下一次分页参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil" name:"Context"`

	// 是否获取完结
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListOver *bool `json:"ListOver,omitnil" name:"ListOver"`

	// 日志详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*JobLogResult `json:"Results,omitnil" name:"Results"`

	// 日志url
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogUrl *string `json:"LogUrl,omitnil" name:"LogUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTaskJobLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskJobLogDetailResponseParams `json:"Response"`
}

func (r *ListTaskJobLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskJobLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LockComponentInfo struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 分区
	Partition *string `json:"Partition,omitnil" name:"Partition"`

	// 锁类型：SHARED_READ、SHARED_WRITE、EXCLUSIVE
	LockType *string `json:"LockType,omitnil" name:"LockType"`

	// 锁级别：DB、TABLE、PARTITION
	LockLevel *string `json:"LockLevel,omitnil" name:"LockLevel"`

	// 锁操作：SELECT,INSERT,UPDATE,DELETE,UNSET,NO_TXN
	DataOperationType *string `json:"DataOperationType,omitnil" name:"DataOperationType"`

	// 是否保持Acid
	IsAcid *bool `json:"IsAcid,omitnil" name:"IsAcid"`

	// 是否动态分区写
	IsDynamicPartitionWrite *bool `json:"IsDynamicPartitionWrite,omitnil" name:"IsDynamicPartitionWrite"`
}

// Predefined struct for user
type LockMetaDataRequestParams struct {
	// 加锁内容
	LockComponentList []*LockComponentInfo `json:"LockComponentList,omitnil" name:"LockComponentList"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 事务id
	TxnId *int64 `json:"TxnId,omitnil" name:"TxnId"`

	// 客户端信息
	AgentInfo *string `json:"AgentInfo,omitnil" name:"AgentInfo"`

	// 主机名
	Hostname *string `json:"Hostname,omitnil" name:"Hostname"`
}

type LockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 加锁内容
	LockComponentList []*LockComponentInfo `json:"LockComponentList,omitnil" name:"LockComponentList"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 事务id
	TxnId *int64 `json:"TxnId,omitnil" name:"TxnId"`

	// 客户端信息
	AgentInfo *string `json:"AgentInfo,omitnil" name:"AgentInfo"`

	// 主机名
	Hostname *string `json:"Hostname,omitnil" name:"Hostname"`
}

func (r *LockMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LockComponentList")
	delete(f, "DatasourceConnectionName")
	delete(f, "TxnId")
	delete(f, "AgentInfo")
	delete(f, "Hostname")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LockMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockMetaDataResponseParams struct {
	// 锁id
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 锁状态：ACQUIRED、WAITING、ABORT、NOT_ACQUIRED
	LockState *string `json:"LockState,omitnil" name:"LockState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type LockMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *LockMetaDataResponseParams `json:"Response"`
}

func (r *LockMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAdvancedStoreLocationRequestParams struct {
	// 查询结果保存cos路径
	StoreLocation *string `json:"StoreLocation,omitnil" name:"StoreLocation"`

	// 是否启用高级设置：0-否，1-是
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`
}

type ModifyAdvancedStoreLocationRequest struct {
	*tchttp.BaseRequest
	
	// 查询结果保存cos路径
	StoreLocation *string `json:"StoreLocation,omitnil" name:"StoreLocation"`

	// 是否启用高级设置：0-否，1-是
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`
}

func (r *ModifyAdvancedStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdvancedStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StoreLocation")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAdvancedStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAdvancedStoreLocationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAdvancedStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAdvancedStoreLocationResponseParams `json:"Response"`
}

func (r *ModifyAdvancedStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdvancedStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataEngineDescriptionRequestParams struct {
	// 要修改的引擎的名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 引擎的描述信息，最大长度为250
	Message *string `json:"Message,omitnil" name:"Message"`
}

type ModifyDataEngineDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的引擎的名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 引擎的描述信息，最大长度为250
	Message *string `json:"Message,omitnil" name:"Message"`
}

func (r *ModifyDataEngineDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataEngineDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataEngineDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataEngineDescriptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDataEngineDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataEngineDescriptionResponseParams `json:"Response"`
}

func (r *ModifyDataEngineDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataEngineDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernEventRuleRequestParams struct {

}

type ModifyGovernEventRuleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyGovernEventRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernEventRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGovernEventRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernEventRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyGovernEventRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGovernEventRuleResponseParams `json:"Response"`
}

func (r *ModifyGovernEventRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernEventRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppBatchRequestParams struct {
	// 需要批量修改的Spark作业任务ID列表
	SparkAppId []*string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// 引擎ID
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppDriverSize *string `json:"AppDriverSize,omitnil" name:"AppDriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitnil" name:"AppExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	AppExecutorNums *uint64 `json:"AppExecutorNums,omitnil" name:"AppExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	AppExecutorMaxNumbers *uint64 `json:"AppExecutorMaxNumbers,omitnil" name:"AppExecutorMaxNumbers"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil" name:"IsInherit"`
}

type ModifySparkAppBatchRequest struct {
	*tchttp.BaseRequest
	
	// 需要批量修改的Spark作业任务ID列表
	SparkAppId []*string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// 引擎ID
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppDriverSize *string `json:"AppDriverSize,omitnil" name:"AppDriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitnil" name:"AppExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	AppExecutorNums *uint64 `json:"AppExecutorNums,omitnil" name:"AppExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	AppExecutorMaxNumbers *uint64 `json:"AppExecutorMaxNumbers,omitnil" name:"AppExecutorMaxNumbers"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil" name:"IsInherit"`
}

func (r *ModifySparkAppBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SparkAppId")
	delete(f, "DataEngine")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "AppExecutorMaxNumbers")
	delete(f, "IsInherit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySparkAppBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppBatchResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySparkAppBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifySparkAppBatchResponseParams `json:"Response"`
}

func (r *ModifySparkAppBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppRequestParams struct {
	// spark作业名
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil" name:"AppFile"`

	// 数据访问策略，CAM Role arn
	RoleArn *int64 `json:"RoleArn,omitnil" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil" name:"AppExecutorNums"`

	// spark作业Id
	SparkAppId *string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil" name:"AppFiles"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil" name:"AppPythonFiles"`

	// spark作业程序入参
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitnil" name:"DataSource"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 任务资源配置是否继承集群配置模板：0（默认）不继承、1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil" name:"IsSessionStarted"`
}

type ModifySparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil" name:"AppFile"`

	// 数据访问策略，CAM Role arn
	RoleArn *int64 `json:"RoleArn,omitnil" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil" name:"AppExecutorNums"`

	// spark作业Id
	SparkAppId *string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil" name:"AppFiles"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil" name:"AppPythonFiles"`

	// spark作业程序入参
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitnil" name:"DataSource"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 任务资源配置是否继承集群配置模板：0（默认）不继承、1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil" name:"IsSessionStarted"`
}

func (r *ModifySparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "SparkAppId")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	delete(f, "IsLocalArchives")
	delete(f, "AppArchives")
	delete(f, "SparkImage")
	delete(f, "SparkImageVersion")
	delete(f, "AppExecutorMaxNumbers")
	delete(f, "SessionId")
	delete(f, "IsInherit")
	delete(f, "IsSessionStarted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySparkAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifySparkAppResponseParams `json:"Response"`
}

func (r *ModifySparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户描述
	UserDescription *string `json:"UserDescription,omitnil" name:"UserDescription"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户描述
	UserDescription *string `json:"UserDescription,omitnil" name:"UserDescription"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserTypeRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户要修改到的类型，ADMIN：管理员，COMMON：一般用户。
	UserType *string `json:"UserType,omitnil" name:"UserType"`
}

type ModifyUserTypeRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户要修改到的类型，ADMIN：管理员，COMMON：一般用户。
	UserType *string `json:"UserType,omitnil" name:"UserType"`
}

func (r *ModifyUserTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserTypeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserTypeResponseParams `json:"Response"`
}

func (r *ModifyUserTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkGroupRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 工作组描述，最大字符数限制50
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil" name:"WorkGroupDescription"`
}

type ModifyWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 工作组描述，最大字符数限制50
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil" name:"WorkGroupDescription"`
}

func (r *ModifyWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "WorkGroupDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkGroupResponseParams `json:"Response"`
}

func (r *ModifyWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MysqlInfo struct {
	// 连接mysql的jdbc url
	JdbcUrl *string `json:"JdbcUrl,omitnil" name:"JdbcUrl"`

	// 用户名
	User *string `json:"User,omitnil" name:"User"`

	// mysql密码
	Password *string `json:"Password,omitnil" name:"Password"`

	// mysql数据源的网络信息
	Location *DatasourceConnectionLocation `json:"Location,omitnil" name:"Location"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库实例ID，和数据库侧保持一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 数据库实例名称，和数据库侧保持一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type NetworkConnection struct {
	// 网络配置id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 网络配置唯一标志符
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateId *string `json:"AssociateId,omitnil" name:"AssociateId"`

	// 计算引擎id
	// 注意：此字段可能返回 null，表示取不到有效值。
	HouseId *string `json:"HouseId,omitnil" name:"HouseId"`

	// 数据源id(已废弃)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionId *string `json:"DatasourceConnectionId,omitnil" name:"DatasourceConnectionId"`

	// 网络配置状态（0-初始化，1-正常）
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil" name:"State"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 创建用户Appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Appid *int64 `json:"Appid,omitnil" name:"Appid"`

	// 计算引擎名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HouseName *string `json:"HouseName,omitnil" name:"HouseName"`

	// 网络配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 网络配置类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionType *int64 `json:"NetworkConnectionType,omitnil" name:"NetworkConnectionType"`

	// 创建用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 创建用户SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil" name:"SubAccountUin"`

	// 网络配置描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionDesc *string `json:"NetworkConnectionDesc,omitnil" name:"NetworkConnectionDesc"`

	// 数据源vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionVpcId *string `json:"DatasourceConnectionVpcId,omitnil" name:"DatasourceConnectionVpcId"`

	// 数据源SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionSubnetId *string `json:"DatasourceConnectionSubnetId,omitnil" name:"DatasourceConnectionSubnetId"`

	// 数据源SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionCidrBlock *string `json:"DatasourceConnectionCidrBlock,omitnil" name:"DatasourceConnectionCidrBlock"`

	// 数据源SubnetCidrBlock
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionSubnetCidrBlock *string `json:"DatasourceConnectionSubnetCidrBlock,omitnil" name:"DatasourceConnectionSubnetCidrBlock"`
}

type NotebookSessionInfo struct {
	// Session名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// Session相关配置，当前支持：eni、roleArn以及用户指定的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arguments []*KVPair `json:"Arguments,omitnil" name:"Arguments"`

	// 运行程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitnil" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitnil" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitnil" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramArchives []*string `json:"ProgramArchives,omitnil" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DriverSize *string `json:"DriverSize,omitnil" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorSize *string `json:"ExecutorSize,omitnil" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil" name:"ExecutorNumbers"`

	// 代理用户，默认为root
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyUser *string `json:"ProxyUser,omitnil" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil" name:"TimeoutInSecond"`

	// Spark任务返回的AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitnil" name:"State"`

	// Session创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 其它信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppInfo []*KVPair `json:"AppInfo,omitnil" name:"AppInfo"`

	// Spark ui地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkUiUrl *string `json:"SparkUiUrl,omitnil" name:"SparkUiUrl"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil" name:"ExecutorMaxNumbers"`
}

type NotebookSessionStatementBatchInformation struct {
	// 任务详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookSessionStatementBatch []*NotebookSessionStatementInfo `json:"NotebookSessionStatementBatch,omitnil" name:"NotebookSessionStatementBatch"`

	// 当前批任务是否运行完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAvailable *bool `json:"IsAvailable,omitnil" name:"IsAvailable"`

	// Session唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// Batch唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`
}

type NotebookSessionStatementInfo struct {
	// 完成时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Completed *int64 `json:"Completed,omitnil" name:"Completed"`

	// 开始时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Started *int64 `json:"Started,omitnil" name:"Started"`

	// 完成进度，百分制
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitnil" name:"Progress"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil" name:"StatementId"`

	// Session Statement状态，包含：waiting（排队中）、running（运行中）、available（正常）、error（异常）、cancelling（取消中）、cancelled（已取消）
	State *string `json:"State,omitnil" name:"State"`

	// Statement输出信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutPut *StatementOutput `json:"OutPut,omitnil" name:"OutPut"`

	// 批任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil" name:"BatchId"`

	// 运行语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type NotebookSessions struct {
	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// 代理用户，默认为root
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyUser *string `json:"ProxyUser,omitnil" name:"ProxyUser"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitnil" name:"State"`

	// Spark任务返回的AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitnil" name:"SparkAppId"`

	// Session名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// Session创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 最新的运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRunningTime *string `json:"LastRunningTime,omitnil" name:"LastRunningTime"`

	// 创建者
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// spark ui地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkUiUrl *string `json:"SparkUiUrl,omitnil" name:"SparkUiUrl"`
}

type Other struct {
	// 枚举类型，默认值为Json，可选值为[Json, Parquet, ORC, AVRD]之一。
	Format *string `json:"Format,omitnil" name:"Format"`
}

type OtherDatasourceConnection struct {
	// 网络参数
	Location *DatasourceConnectionLocation `json:"Location,omitnil" name:"Location"`
}

type Partition struct {
	// 分区列名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分区类型。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 对分区的描述。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 隐式分区转换策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transform *string `json:"Transform,omitnil" name:"Transform"`

	// 转换策略参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformArgs []*string `json:"TransformArgs,omitnil" name:"TransformArgs"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil" name:"CreateTime"`
}

type Policy struct {
	// 需要授权的数据库名，填*代表当前Catalog下所有数据库。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定数据库。
	Database *string `json:"Database,omitnil" name:"Database"`

	// 需要授权的数据源名称，管理员级别下只支持填*（代表该级别全部资源）；数据源级别和数据库级别鉴权的情况下，只支持填COSDataCatalog或者*；在数据表级别鉴权下可以填写用户自定义数据源。不填情况下默认为DataLakeCatalog。注意：如果是对用户自定义数据源进行鉴权，DLC能够管理的权限是用户接入数据源的时候提供的账户的子集。
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 需要授权的表名，填*代表当前Database下所有表。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定数据表。
	Table *string `json:"Table,omitnil" name:"Table"`

	// 授权的权限操作，对于不同级别的鉴权提供不同操作。管理员权限：ALL，不填默认为ALL；数据连接级鉴权：CREATE；数据库级别鉴权：ALL、CREATE、ALTER、DROP；数据表权限：ALL、SELECT、INSERT、ALTER、DELETE、DROP、UPDATE。注意：在数据表权限下，指定的数据源不为COSDataCatalog的时候，只支持SELECT操作。
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// 授权类型，现在支持八种授权类型：ADMIN:管理员级别鉴权 DATASOURCE：数据连接级别鉴权 DATABASE：数据库级别鉴权 TABLE：表级别鉴权 VIEW：视图级别鉴权 FUNCTION：函数级别鉴权 COLUMN：列级别鉴权 ENGINE：数据引擎鉴权。不填默认为管理员级别鉴权。
	PolicyType *string `json:"PolicyType,omitnil" name:"PolicyType"`

	// 需要授权的函数名，填*代表当前Catalog下所有函数。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定函数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function *string `json:"Function,omitnil" name:"Function"`

	// 需要授权的视图，填*代表当前Database下所有视图。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定视图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	View *string `json:"View,omitnil" name:"View"`

	// 需要授权的列，填*代表当前所有列。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Column *string `json:"Column,omitnil" name:"Column"`

	// 需要授权的数据引擎，填*代表当前所有引擎。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// 用户是否可以进行二次授权。当为true的时候，被授权的用户可以将本次获取的权限再次授权给其他子用户。默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReAuth *bool `json:"ReAuth,omitnil" name:"ReAuth"`

	// 权限来源，入参不填。USER：权限来自用户本身；WORKGROUP：权限来自绑定的工作组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`

	// 授权模式，入参不填。COMMON：普通模式；SENIOR：高级模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 操作者，入参不填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 权限创建的时间，入参不填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 权限所属工作组的ID，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceId *int64 `json:"SourceId,omitnil" name:"SourceId"`

	// 权限所属工作组的名称，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type Policys struct {
	// 策略集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`

	// 策略总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type PrestoMonitorMetrics struct {
	// 	Alluxio本地缓存命中率
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalCacheHitRate *float64 `json:"LocalCacheHitRate,omitnil" name:"LocalCacheHitRate"`

	// Fragment缓存命中率
	// 注意：此字段可能返回 null，表示取不到有效值。
	FragmentCacheHitRate *float64 `json:"FragmentCacheHitRate,omitnil" name:"FragmentCacheHitRate"`
}

type Property struct {
	// 属性key名称。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 属性key对应的value。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type PythonSparkImage struct {
	// spark镜像唯一id
	SparkImageId *string `json:"SparkImageId,omitnil" name:"SparkImageId"`

	// 集群小版本镜像id
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil" name:"ChildImageVersionId"`

	// spark镜像名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil" name:"SparkImageVersion"`

	// spark镜像描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

// Predefined struct for user
type QueryResultRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// lastReadFile为上一次读取的文件，lastReadOffset为上一次读取到的位置
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`
}

type QueryResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// lastReadFile为上一次读取的文件，lastReadOffset为上一次读取到的位置
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`
}

func (r *QueryResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResultResponseParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 结果数据
	ResultSet *string `json:"ResultSet,omitnil" name:"ResultSet"`

	// schema
	ResultSchema []*Column `json:"ResultSchema,omitnil" name:"ResultSchema"`

	// 分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryResultResponse struct {
	*tchttp.BaseResponse
	Response *QueryResultResponseParams `json:"Response"`
}

func (r *QueryResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDataEngineRequestParams struct {
	// CU队列名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 续费时长，单位月，最少续费1一个月
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 付费类型，默认为1，预付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 单位，默认m，仅能填m
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 自动续费标志，0，初始状态，默认不自动续费，若用户有预付费不停服特权，自动续费。1：自动续费。2：明确不自动续费。不传该参数默认为0
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type RenewDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// CU队列名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 续费时长，单位月，最少续费1一个月
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 付费类型，默认为1，预付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 单位，默认m，仅能填m
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 自动续费标志，0，初始状态，默认不自动续费，若用户有预付费不停服特权，自动续费。1：自动续费。2：明确不自动续费。不传该参数默认为0
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

func (r *RenewDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "TimeSpan")
	delete(f, "PayMode")
	delete(f, "TimeUnit")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDataEngineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *RenewDataEngineResponseParams `json:"Response"`
}

func (r *RenewDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportHeartbeatMetaDataRequestParams struct {
	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 锁ID
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil" name:"TxnId"`
}

type ReportHeartbeatMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 锁ID
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil" name:"TxnId"`
}

func (r *ReportHeartbeatMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportHeartbeatMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceConnectionName")
	delete(f, "LockId")
	delete(f, "TxnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportHeartbeatMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportHeartbeatMetaDataResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReportHeartbeatMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *ReportHeartbeatMetaDataResponseParams `json:"Response"`
}

func (r *ReportHeartbeatMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportHeartbeatMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {
	// 归属类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributionType *string `json:"AttributionType,omitnil" name:"AttributionType"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 引擎名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 如资源类型为spark-sql 取值为Name, 如为spark-batch 取值为session app_name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *string `json:"Instance,omitnil" name:"Instance"`

	// 亲和性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Favor []*FavorInfo `json:"Favor,omitnil" name:"Favor"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type RestartDataEngineRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 是否强制重启，忽略任务
	ForcedOperation *bool `json:"ForcedOperation,omitnil" name:"ForcedOperation"`
}

type RestartDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 是否强制重启，忽略任务
	ForcedOperation *bool `json:"ForcedOperation,omitnil" name:"ForcedOperation"`
}

func (r *RestartDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "ForcedOperation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDataEngineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RestartDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *RestartDataEngineResponseParams `json:"Response"`
}

func (r *RestartDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackDataEngineImageRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 检查是否能回滚的接口返回的FromRecordId参数
	FromRecordId *string `json:"FromRecordId,omitnil" name:"FromRecordId"`

	// 检查是否能回滚的接口返回的ToRecordId参数
	ToRecordId *string `json:"ToRecordId,omitnil" name:"ToRecordId"`
}

type RollbackDataEngineImageRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 检查是否能回滚的接口返回的FromRecordId参数
	FromRecordId *string `json:"FromRecordId,omitnil" name:"FromRecordId"`

	// 检查是否能回滚的接口返回的ToRecordId参数
	ToRecordId *string `json:"ToRecordId,omitnil" name:"ToRecordId"`
}

func (r *RollbackDataEngineImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackDataEngineImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "FromRecordId")
	delete(f, "ToRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackDataEngineImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackDataEngineImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RollbackDataEngineImageResponse struct {
	*tchttp.BaseResponse
	Response *RollbackDataEngineImageResponseParams `json:"Response"`
}

func (r *RollbackDataEngineImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackDataEngineImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SQLTask struct {
	// base64加密后的SQL语句
	SQL *string `json:"SQL,omitnil" name:"SQL"`

	// 任务的配置信息
	Config []*KVPair `json:"Config,omitnil" name:"Config"`
}

type Script struct {
	// 脚本Id，长度36字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptId *string `json:"ScriptId,omitnil" name:"ScriptId"`

	// 脚本名称，长度0-25。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitnil" name:"ScriptName"`

	// 脚本描述，长度0-50。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDesc *string `json:"ScriptDesc,omitnil" name:"ScriptDesc"`

	// 默认关联数据库。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// SQL描述，长度0-10000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLStatement *string `json:"SQLStatement,omitnil" name:"SQLStatement"`

	// 更新时间戳， 单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type SessionResourceTemplate struct {
	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	// 注意：此字段可能返回 null，表示取不到有效值。
	DriverSize *string `json:"DriverSize,omitnil" name:"DriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorSize *string `json:"ExecutorSize,omitnil" name:"ExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorNums *uint64 `json:"ExecutorNums,omitnil" name:"ExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil" name:"ExecutorMaxNumbers"`
}

type SmartOptimizerIndexPolicy struct {
	// 开启索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexEnable *string `json:"IndexEnable,omitnil" name:"IndexEnable"`
}

type SmartOptimizerLifecyclePolicy struct {
	// 生命周期启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecycleEnable *string `json:"LifecycleEnable,omitnil" name:"LifecycleEnable"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expiration *int64 `json:"Expiration,omitnil" name:"Expiration"`

	// 是否删表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropTable *bool `json:"DropTable,omitnil" name:"DropTable"`
}

type SmartOptimizerPolicy struct {
	// 是否继承
	// 注意：此字段可能返回 null，表示取不到有效值。
	Inherit *string `json:"Inherit,omitnil" name:"Inherit"`

	// ResourceInfo
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*ResourceInfo `json:"Resources,omitnil" name:"Resources"`

	// SmartOptimizerWrittenPolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	Written *SmartOptimizerWrittenPolicy `json:"Written,omitnil" name:"Written"`

	// SmartOptimizerLifecyclePolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lifecycle *SmartOptimizerLifecyclePolicy `json:"Lifecycle,omitnil" name:"Lifecycle"`

	// SmartOptimizerIndexPolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *SmartOptimizerIndexPolicy `json:"Index,omitnil" name:"Index"`
}

type SmartOptimizerWrittenPolicy struct {

}

type SmartPolicy struct {
	// 基础信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *SmartPolicyBaseInfo `json:"BaseInfo,omitnil" name:"BaseInfo"`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *SmartOptimizerPolicy `json:"Policy,omitnil" name:"Policy"`
}

type SmartPolicyBaseInfo struct {
	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// Catalog/Database/Table
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitnil" name:"PolicyType"`

	// Catalog名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Database *string `json:"Database,omitnil" name:"Database"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitnil" name:"Table"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil" name:"AppId"`
}

type SparkJobInfo struct {
	// spark作业ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// spark作业类型，可去1或者2，1表示batch作业， 2表示streaming作业
	JobType *int64 `json:"JobType,omitnil" name:"JobType"`

	// 引擎名
	DataEngine *string `json:"DataEngine,omitnil" name:"DataEngine"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil" name:"Eni"`

	// 程序包是否本地上传，cos或者lakefs
	IsLocal *string `json:"IsLocal,omitnil" name:"IsLocal"`

	// 程序包路径
	JobFile *string `json:"JobFile,omitnil" name:"JobFile"`

	// 角色ID
	RoleArn *int64 `json:"RoleArn,omitnil" name:"RoleArn"`

	// spark作业运行主类
	MainClass *string `json:"MainClass,omitnil" name:"MainClass"`

	// 命令行参数，spark作业命令行参数，空格分隔
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`

	// spark原生配置，换行符分隔
	JobConf *string `json:"JobConf,omitnil" name:"JobConf"`

	// 依赖jars是否本地上传，cos或者lakefs
	IsLocalJars *string `json:"IsLocalJars,omitnil" name:"IsLocalJars"`

	// spark作业依赖jars，逗号分隔
	JobJars *string `json:"JobJars,omitnil" name:"JobJars"`

	// 依赖文件是否本地上传，cos或者lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitnil" name:"IsLocalFiles"`

	// spark作业依赖文件，逗号分隔
	JobFiles *string `json:"JobFiles,omitnil" name:"JobFiles"`

	// spark作业driver资源大小
	JobDriverSize *string `json:"JobDriverSize,omitnil" name:"JobDriverSize"`

	// spark作业executor资源大小
	JobExecutorSize *string `json:"JobExecutorSize,omitnil" name:"JobExecutorSize"`

	// spark作业executor个数
	JobExecutorNums *int64 `json:"JobExecutorNums,omitnil" name:"JobExecutorNums"`

	// spark流任务最大重试次数
	JobMaxAttempts *int64 `json:"JobMaxAttempts,omitnil" name:"JobMaxAttempts"`

	// spark作业创建者
	JobCreator *string `json:"JobCreator,omitnil" name:"JobCreator"`

	// spark作业创建时间
	JobCreateTime *int64 `json:"JobCreateTime,omitnil" name:"JobCreateTime"`

	// spark作业更新时间
	JobUpdateTime *uint64 `json:"JobUpdateTime,omitnil" name:"JobUpdateTime"`

	// spark作业最近任务ID
	CurrentTaskId *string `json:"CurrentTaskId,omitnil" name:"CurrentTaskId"`

	// spark作业最近运行状态
	JobStatus *int64 `json:"JobStatus,omitnil" name:"JobStatus"`

	// spark流作业统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamingStat *StreamingStatistics `json:"StreamingStat,omitnil" name:"StreamingStat"`

	// 数据源名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitnil" name:"DataSource"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil" name:"IsLocalPythonFiles"`

	// 注：该返回值已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppPythonFiles *string `json:"AppPythonFiles,omitnil" name:"AppPythonFiles"`

	// archives：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLocalArchives *string `json:"IsLocalArchives,omitnil" name:"IsLocalArchives"`

	// archives：依赖资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobArchives *string `json:"JobArchives,omitnil" name:"JobArchives"`

	// Spark Image 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImage *string `json:"SparkImage,omitnil" name:"SparkImage"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobPythonFiles *string `json:"JobPythonFiles,omitnil" name:"JobPythonFiles"`

	// 当前job正在运行或准备运行的任务个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNum *int64 `json:"TaskNum,omitnil" name:"TaskNum"`

	// 引擎状态：-100（默认：未知状态），-2~11：引擎正常状态；
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineStatus *int64 `json:"DataEngineStatus,omitnil" name:"DataEngineStatus"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于JobExecutorNums
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobExecutorMaxNumbers *int64 `json:"JobExecutorMaxNumbers,omitnil" name:"JobExecutorMaxNumbers"`

	// 镜像版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImageVersion *string `json:"SparkImageVersion,omitnil" name:"SparkImageVersion"`

	// 查询脚本关联id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil" name:"SessionId"`

	// spark_emr_livy
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineClusterType *string `json:"DataEngineClusterType,omitnil" name:"DataEngineClusterType"`

	// Spark 3.2-EMR
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineImageVersion *string `json:"DataEngineImageVersion,omitnil" name:"DataEngineImageVersion"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsInherit *uint64 `json:"IsInherit,omitnil" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil" name:"IsSessionStarted"`
}

type SparkMonitorMetrics struct {
	// shuffle写溢出到COS数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShuffleWriteBytesCos *int64 `json:"ShuffleWriteBytesCos,omitnil" name:"ShuffleWriteBytesCos"`

	// shuffle写数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShuffleWriteBytesTotal *int64 `json:"ShuffleWriteBytesTotal,omitnil" name:"ShuffleWriteBytesTotal"`
}

type SparkSessionBatchLog struct {
	// 日志步骤：BEG/CS/DS/DSS/DSF/FINF/RTO/CANCEL/CT/DT/DTS/DTF/FINT/EXCE
	// 注意：此字段可能返回 null，表示取不到有效值。
	Step *string `json:"Step,omitnil" name:"Step"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil" name:"Time"`

	// 日志提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 日志操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operate []*SparkSessionBatchLogOperate `json:"Operate,omitnil" name:"Operate"`
}

type SparkSessionBatchLogOperate struct {
	// 操作提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 操作类型：COPY、LOG、UI、RESULT、List、TAB
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operate *string `json:"Operate,omitnil" name:"Operate"`

	// 补充信息：如：taskid、sessionid、sparkui等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supplement []*KVPair `json:"Supplement,omitnil" name:"Supplement"`
}

type StatementInformation struct {
	// SQL任务唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// SQL内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQL *string `json:"SQL,omitnil" name:"SQL"`
}

type StatementOutput struct {
	// 执行总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionCount *int64 `json:"ExecutionCount,omitnil" name:"ExecutionCount"`

	// Statement数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*KVPair `json:"Data,omitnil" name:"Data"`

	// Statement状态:ok,error
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorName *string `json:"ErrorName,omitnil" name:"ErrorName"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorValue *string `json:"ErrorValue,omitnil" name:"ErrorValue"`

	// 错误堆栈信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage []*string `json:"ErrorMessage,omitnil" name:"ErrorMessage"`

	// SQL类型任务结果返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLResult *string `json:"SQLResult,omitnil" name:"SQLResult"`
}

type StreamingStatistics struct {
	// 任务开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 数据接收器数
	Receivers *int64 `json:"Receivers,omitnil" name:"Receivers"`

	// 运行中的接收器数
	NumActiveReceivers *int64 `json:"NumActiveReceivers,omitnil" name:"NumActiveReceivers"`

	// 不活跃的接收器数
	NumInactiveReceivers *int64 `json:"NumInactiveReceivers,omitnil" name:"NumInactiveReceivers"`

	// 运行中的批数
	NumActiveBatches *int64 `json:"NumActiveBatches,omitnil" name:"NumActiveBatches"`

	// 待处理的批数
	NumRetainedCompletedBatches *int64 `json:"NumRetainedCompletedBatches,omitnil" name:"NumRetainedCompletedBatches"`

	// 已完成的批数
	NumTotalCompletedBatches *int64 `json:"NumTotalCompletedBatches,omitnil" name:"NumTotalCompletedBatches"`

	// 平均输入速率
	AverageInputRate *float64 `json:"AverageInputRate,omitnil" name:"AverageInputRate"`

	// 平均等待时长
	AverageSchedulingDelay *float64 `json:"AverageSchedulingDelay,omitnil" name:"AverageSchedulingDelay"`

	// 平均处理时长
	AverageProcessingTime *float64 `json:"AverageProcessingTime,omitnil" name:"AverageProcessingTime"`

	// 平均延时
	AverageTotalDelay *float64 `json:"AverageTotalDelay,omitnil" name:"AverageTotalDelay"`
}

// Predefined struct for user
type SuspendResumeDataEngineRequestParams struct {
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 操作类型 suspend/resume
	Operate *string `json:"Operate,omitnil" name:"Operate"`
}

type SuspendResumeDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 操作类型 suspend/resume
	Operate *string `json:"Operate,omitnil" name:"Operate"`
}

func (r *SuspendResumeDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendResumeDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Operate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SuspendResumeDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SuspendResumeDataEngineResponseParams struct {
	// 虚拟集群详细信息
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SuspendResumeDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *SuspendResumeDataEngineResponseParams `json:"Response"`
}

func (r *SuspendResumeDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendResumeDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineImageRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 新镜像版本ID
	NewImageVersionId *string `json:"NewImageVersionId,omitnil" name:"NewImageVersionId"`
}

type SwitchDataEngineImageRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 新镜像版本ID
	NewImageVersionId *string `json:"NewImageVersionId,omitnil" name:"NewImageVersionId"`
}

func (r *SwitchDataEngineImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "NewImageVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDataEngineImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchDataEngineImageResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDataEngineImageResponseParams `json:"Response"`
}

func (r *SwitchDataEngineImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineRequestParams struct {
	// 主集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 是否开启备集群
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitnil" name:"StartStandbyCluster"`
}

type SwitchDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 主集群名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 是否开启备集群
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitnil" name:"StartStandbyCluster"`
}

func (r *SwitchDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "StartStandbyCluster")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDataEngineResponseParams `json:"Response"`
}

func (r *SwitchDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TColumn struct {
	// 字段名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 字段描述
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 字段默认值
	Default *string `json:"Default,omitnil" name:"Default"`

	// 字段是否是非空
	NotNull *bool `json:"NotNull,omitnil" name:"NotNull"`

	// 表示整个 numeric 的长度,取值1-38
	Precision *int64 `json:"Precision,omitnil" name:"Precision"`

	// 表示小数部分的长度
	// Scale小于Precision
	Scale *int64 `json:"Scale,omitnil" name:"Scale"`
}

type TPartition struct {
	// 字段名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 字段描述
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 分区类型
	PartitionType *string `json:"PartitionType,omitnil" name:"PartitionType"`

	// 分区格式
	PartitionFormat *string `json:"PartitionFormat,omitnil" name:"PartitionFormat"`

	// 分区分隔数
	PartitionDot *int64 `json:"PartitionDot,omitnil" name:"PartitionDot"`

	// 分区转换策略
	Transform *string `json:"Transform,omitnil" name:"Transform"`

	// 策略参数
	TransformArgs []*string `json:"TransformArgs,omitnil" name:"TransformArgs"`
}

type TableBaseInfo struct {
	// 该数据表所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 数据表名字
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 该数据表所属数据源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 该数据表备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableComment *string `json:"TableComment,omitnil" name:"TableComment"`

	// 具体类型，表or视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据格式类型，hive，iceberg等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableFormat *string `json:"TableFormat,omitnil" name:"TableFormat"`

	// 建表用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`

	// 建表用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSubUin *string `json:"UserSubUin,omitnil" name:"UserSubUin"`

	// 数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernPolicy *DataGovernPolicy `json:"GovernPolicy,omitnil" name:"GovernPolicy"`

	// 库数据治理是否关闭，关闭：true，开启：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbGovernPolicyIsDisable *string `json:"DbGovernPolicyIsDisable,omitnil" name:"DbGovernPolicyIsDisable"`

	// 智能数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartPolicy *SmartPolicy `json:"SmartPolicy,omitnil" name:"SmartPolicy"`
}

type TableInfo struct {
	// 数据表配置信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil" name:"TableBaseInfo"`

	// 数据表格式。每次入参可选如下其一的KV结构，[TextFile，CSV，Json, Parquet, ORC, AVRD]。
	DataFormat *DataFormat `json:"DataFormat,omitnil" name:"DataFormat"`

	// 数据表列信息。
	Columns []*Column `json:"Columns,omitnil" name:"Columns"`

	// 数据表分块信息。
	Partitions []*Partition `json:"Partitions,omitnil" name:"Partitions"`

	// 数据存储路径。当前仅支持cos路径，格式如下：cosn://bucket-name/filepath。
	Location *string `json:"Location,omitnil" name:"Location"`
}

type TableResponseInfo struct {
	// 数据表基本信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil" name:"TableBaseInfo"`

	// 数据表列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil" name:"Columns"`

	// 数据表分块信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*Partition `json:"Partitions,omitnil" name:"Partitions"`

	// 数据存储路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil" name:"Location"`

	// 数据表属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`

	// 数据表更新时间, 单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil" name:"ModifiedTime"`

	// 数据表创建时间,单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 数据格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputFormat *string `json:"InputFormat,omitnil" name:"InputFormat"`

	// 数据表存储大小（单位：Byte）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// 数据表行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitnil" name:"RecordCount"`

	// xxxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapMaterializedViewName *string `json:"MapMaterializedViewName,omitnil" name:"MapMaterializedViewName"`
}

type TagInfo struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type Task struct {
	// SQL查询任务
	SQLTask *SQLTask `json:"SQLTask,omitnil" name:"SQLTask"`

	// Spark SQL查询任务
	SparkSQLTask *SQLTask `json:"SparkSQLTask,omitnil" name:"SparkSQLTask"`
}

type TaskResponseInfo struct {
	// 任务所属Database的名称。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 任务数据量。
	DataAmount *int64 `json:"DataAmount,omitnil" name:"DataAmount"`

	// 任务Id。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 计算耗时，单位： ms
	UsedTime *int64 `json:"UsedTime,omitnil" name:"UsedTime"`

	// 任务输出路径。
	OutputPath *string `json:"OutputPath,omitnil" name:"OutputPath"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务状态：0 初始化， 1 执行中， 2 执行成功，3 数据写入中，4 排队中。-1 执行失败，-3 已取消。
	State *int64 `json:"State,omitnil" name:"State"`

	// 任务SQL类型，DDL|DML等
	SQLType *string `json:"SQLType,omitnil" name:"SQLType"`

	// 任务SQL语句
	SQL *string `json:"SQL,omitnil" name:"SQL"`

	// 结果是否过期。
	ResultExpired *bool `json:"ResultExpired,omitnil" name:"ResultExpired"`

	// 数据影响统计信息。
	RowAffectInfo *string `json:"RowAffectInfo,omitnil" name:"RowAffectInfo"`

	// 任务结果数据表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSet *string `json:"DataSet,omitnil" name:"DataSet"`

	// 失败信息, 例如：errorMessage。该字段已废弃。
	Error *string `json:"Error,omitnil" name:"Error"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitnil" name:"Percentage"`

	// 任务执行输出信息。
	OutputMessage *string `json:"OutputMessage,omitnil" name:"OutputMessage"`

	// 执行SQL的引擎类型
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`

	// 任务进度明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgressDetail *string `json:"ProgressDetail,omitnil" name:"ProgressDetail"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 计算资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 执行sql的子uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil" name:"OperateUin"`

	// 计算资源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 导入类型是本地导入还是cos
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputType *string `json:"InputType,omitnil" name:"InputType"`

	// 导入配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputConf *string `json:"InputConf,omitnil" name:"InputConf"`

	// 数据条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataNumber *int64 `json:"DataNumber,omitnil" name:"DataNumber"`

	// 查询数据能不能下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanDownload *bool `json:"CanDownload,omitnil" name:"CanDownload"`

	// 用户别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`

	// spark应用作业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobName *string `json:"SparkJobName,omitnil" name:"SparkJobName"`

	// spark应用作业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobId *string `json:"SparkJobId,omitnil" name:"SparkJobId"`

	// spark应用入口jar文件
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobFile *string `json:"SparkJobFile,omitnil" name:"SparkJobFile"`

	// spark ui url
	// 注意：此字段可能返回 null，表示取不到有效值。
	UiUrl *string `json:"UiUrl,omitnil" name:"UiUrl"`

	// 任务耗时，单位： ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTime *int64 `json:"TotalTime,omitnil" name:"TotalTime"`

	// spark app job执行task的程序入口参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CmdArgs *string `json:"CmdArgs,omitnil" name:"CmdArgs"`

	// 集群镜像大版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersion *string `json:"ImageVersion,omitnil" name:"ImageVersion"`

	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	// 注意：此字段可能返回 null，表示取不到有效值。
	DriverSize *string `json:"DriverSize,omitnil" name:"DriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorSize *string `json:"ExecutorSize,omitnil" name:"ExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorNums *uint64 `json:"ExecutorNums,omitnil" name:"ExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil" name:"ExecutorMaxNumbers"`

	// 任务公共指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonMetrics *CommonMetrics `json:"CommonMetrics,omitnil" name:"CommonMetrics"`

	// spark任务指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkMonitorMetrics *SparkMonitorMetrics `json:"SparkMonitorMetrics,omitnil" name:"SparkMonitorMetrics"`

	// presto任务指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrestoMonitorMetrics *PrestoMonitorMetrics `json:"PrestoMonitorMetrics,omitnil" name:"PrestoMonitorMetrics"`
}

type TaskResultInfo struct {
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 数据源名称，当前任务执行时候选中的默认数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`

	// 数据库名称，当前任务执行时候选中的默认数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 当前执行的SQL，一个任务包含一个SQL
	SQL *string `json:"SQL,omitnil" name:"SQL"`

	// 执行任务的类型，现在分为DDL、DML、DQL
	SQLType *string `json:"SQLType,omitnil" name:"SQLType"`

	// 任务当前的状态，0：初始化 1：任务运行中 2：任务执行成功  3：数据写入中 4：排队中 -1：任务执行失败 -3：用户手动终止 。只有任务执行成功的情况下，才会返回任务执行的结果
	State *int64 `json:"State,omitnil" name:"State"`

	// 扫描的数据量，单位byte
	DataAmount *int64 `json:"DataAmount,omitnil" name:"DataAmount"`

	// 计算耗时，单位： ms
	UsedTime *int64 `json:"UsedTime,omitnil" name:"UsedTime"`

	// 任务结果输出的COS桶地址
	OutputPath *string `json:"OutputPath,omitnil" name:"OutputPath"`

	// 任务创建时间，时间戳
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务执行信息，成功时返回success，失败时返回失败原因
	OutputMessage *string `json:"OutputMessage,omitnil" name:"OutputMessage"`

	// 被影响的行数
	RowAffectInfo *string `json:"RowAffectInfo,omitnil" name:"RowAffectInfo"`

	// 结果的schema信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSchema []*Column `json:"ResultSchema,omitnil" name:"ResultSchema"`

	// 结果信息，反转义后，外层数组的每个元素为一行数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSet *string `json:"ResultSet,omitnil" name:"ResultSet"`

	// 分页信息，如果没有更多结果数据，nextToken为空
	NextToken *string `json:"NextToken,omitnil" name:"NextToken"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitnil" name:"Percentage"`

	// 任务进度明细
	ProgressDetail *string `json:"ProgressDetail,omitnil" name:"ProgressDetail"`

	// 控制台展示格式。table：表格展示 text：文本展示
	DisplayFormat *string `json:"DisplayFormat,omitnil" name:"DisplayFormat"`

	// 任务耗时，单位： ms
	TotalTime *int64 `json:"TotalTime,omitnil" name:"TotalTime"`
}

type TasksInfo struct {
	// 任务类型，SQLTask：SQL查询任务。SparkSQLTask：Spark SQL查询任务
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`

	// 容错策略。Proceed：前面任务出错/取消后继续执行后面的任务。Terminate：前面的任务出错/取消之后终止后面任务的执行，后面的任务全部标记为已取消。
	FailureTolerance *string `json:"FailureTolerance,omitnil" name:"FailureTolerance"`

	// base64加密后的SQL语句，用";"号分隔每个SQL语句，一次最多提交50个任务。严格按照前后顺序执行
	SQL *string `json:"SQL,omitnil" name:"SQL"`

	// 任务的配置信息，当前仅支持SparkSQLTask任务。
	Config []*KVPair `json:"Config,omitnil" name:"Config"`

	// 任务的用户自定义参数信息
	Params []*KVPair `json:"Params,omitnil" name:"Params"`
}

type TasksOverview struct {
	// 正在排队的任务个数
	TaskQueuedCount *int64 `json:"TaskQueuedCount,omitnil" name:"TaskQueuedCount"`

	// 初始化的任务个数
	TaskInitCount *int64 `json:"TaskInitCount,omitnil" name:"TaskInitCount"`

	// 正在执行的任务个数
	TaskRunningCount *int64 `json:"TaskRunningCount,omitnil" name:"TaskRunningCount"`

	// 当前时间范围的总任务个数
	TotalTaskCount *int64 `json:"TotalTaskCount,omitnil" name:"TotalTaskCount"`
}

type TextFile struct {
	// 文本类型，本参数取值为TextFile。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 处理文本用的正则表达式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitnil" name:"Regex"`
}

// Predefined struct for user
type UnbindWorkGroupsFromUserRequestParams struct {
	// 解绑的工作组Id和用户Id的关联关系
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil" name:"AddInfo"`
}

type UnbindWorkGroupsFromUserRequest struct {
	*tchttp.BaseRequest
	
	// 解绑的工作组Id和用户Id的关联关系
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil" name:"AddInfo"`
}

func (r *UnbindWorkGroupsFromUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkGroupsFromUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindWorkGroupsFromUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindWorkGroupsFromUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnbindWorkGroupsFromUserResponse struct {
	*tchttp.BaseResponse
	Response *UnbindWorkGroupsFromUserResponseParams `json:"Response"`
}

func (r *UnbindWorkGroupsFromUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkGroupsFromUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockMetaDataRequestParams struct {
	// 锁ID
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

type UnlockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 锁ID
	LockId *int64 `json:"LockId,omitnil" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil" name:"DatasourceConnectionName"`
}

func (r *UnlockMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LockId")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlockMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockMetaDataResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnlockMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *UnlockMetaDataResponseParams `json:"Response"`
}

func (r *UnlockMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataEngineConfigRequestParams struct {
	// 引擎ID
	DataEngineIds []*string `json:"DataEngineIds,omitnil" name:"DataEngineIds"`

	// 引擎配置命令，支持UpdateSparkSQLLakefsPath（更新原生表配置）、UpdateSparkSQLResultPath（更新结果路径配置）
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil" name:"DataEngineConfigCommand"`
}

type UpdateDataEngineConfigRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineIds []*string `json:"DataEngineIds,omitnil" name:"DataEngineIds"`

	// 引擎配置命令，支持UpdateSparkSQLLakefsPath（更新原生表配置）、UpdateSparkSQLResultPath（更新结果路径配置）
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil" name:"DataEngineConfigCommand"`
}

func (r *UpdateDataEngineConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineIds")
	delete(f, "DataEngineConfigCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataEngineConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataEngineConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateDataEngineConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataEngineConfigResponseParams `json:"Response"`
}

func (r *UpdateDataEngineConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataEngineRequestParams struct {
	// 资源大小
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil" name:"MaxClusters"`

	// 开启自动刷新：true：开启、false（默认）：关闭
	AutoResume *bool `json:"AutoResume,omitnil" name:"AutoResume"`

	// 数据引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 相关信息
	Message *string `json:"Message,omitnil" name:"Message"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil" name:"CrontabResumeSuspendStrategy"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil" name:"TolerableQueueTime"`

	// 集群自动挂起时间
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil" name:"AutoSuspendTime"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil" name:"ElasticLimit"`

	// Spark批作业集群Session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`
}

type UpdateDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 资源大小
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil" name:"MaxClusters"`

	// 开启自动刷新：true：开启、false（默认）：关闭
	AutoResume *bool `json:"AutoResume,omitnil" name:"AutoResume"`

	// 数据引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil" name:"DataEngineName"`

	// 相关信息
	Message *string `json:"Message,omitnil" name:"Message"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil" name:"CrontabResumeSuspendStrategy"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil" name:"TolerableQueueTime"`

	// 集群自动挂起时间
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil" name:"AutoSuspendTime"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil" name:"ElasticLimit"`

	// Spark批作业集群Session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`
}

func (r *UpdateDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Size")
	delete(f, "MinClusters")
	delete(f, "MaxClusters")
	delete(f, "AutoResume")
	delete(f, "DataEngineName")
	delete(f, "Message")
	delete(f, "AutoSuspend")
	delete(f, "CrontabResumeSuspend")
	delete(f, "CrontabResumeSuspendStrategy")
	delete(f, "MaxConcurrency")
	delete(f, "TolerableQueueTime")
	delete(f, "AutoSuspendTime")
	delete(f, "ElasticSwitch")
	delete(f, "ElasticLimit")
	delete(f, "SessionResourceTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataEngineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataEngineResponseParams `json:"Response"`
}

func (r *UpdateDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRowFilterRequestParams struct {
	// 行过滤策略的id，此值可以通过DescribeUserInfo或者DescribeWorkGroupInfo接口获取
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 新的过滤策略。
	Policy *Policy `json:"Policy,omitnil" name:"Policy"`
}

type UpdateRowFilterRequest struct {
	*tchttp.BaseRequest
	
	// 行过滤策略的id，此值可以通过DescribeUserInfo或者DescribeWorkGroupInfo接口获取
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 新的过滤策略。
	Policy *Policy `json:"Policy,omitnil" name:"Policy"`
}

func (r *UpdateRowFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRowFilterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRowFilterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRowFilterResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateRowFilterResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRowFilterResponseParams `json:"Response"`
}

func (r *UpdateRowFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRowFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDataEngineConfigRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 引擎配置项
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil" name:"DataEngineConfigPairs"`

	// 作业引擎资源配置模版
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`
}

type UpdateUserDataEngineConfigRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`

	// 引擎配置项
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil" name:"DataEngineConfigPairs"`

	// 作业引擎资源配置模版
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil" name:"SessionResourceTemplate"`
}

func (r *UpdateUserDataEngineConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDataEngineConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "DataEngineConfigPairs")
	delete(f, "SessionResourceTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserDataEngineConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDataEngineConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateUserDataEngineConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserDataEngineConfigResponseParams `json:"Response"`
}

func (r *UpdateUserDataEngineConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDataEngineConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDataEngineImageRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

type UpgradeDataEngineImageRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil" name:"DataEngineId"`
}

func (r *UpgradeDataEngineImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDataEngineImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDataEngineImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDataEngineImageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeDataEngineImageResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDataEngineImageResponseParams `json:"Response"`
}

func (r *UpgradeDataEngineImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDataEngineImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDetailInfo struct {
	// 用户Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 返回的信息类型，Group：返回的当前用户的工作组信息；DataAuth：返回的当前用户的数据权限信息；EngineAuth：返回的当前用户的引擎权限信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 用户类型：ADMIN：管理员 COMMON：一般用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 用户描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitnil" name:"UserDescription"`

	// 数据权限信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataPolicyInfo *Policys `json:"DataPolicyInfo,omitnil" name:"DataPolicyInfo"`

	// 引擎权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnginePolicyInfo *Policys `json:"EnginePolicyInfo,omitnil" name:"EnginePolicyInfo"`

	// 绑定到该用户的工作组集合信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupInfo *WorkGroups `json:"WorkGroupInfo,omitnil" name:"WorkGroupInfo"`

	// 用户别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`

	// 行过滤集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowFilterInfo *Policys `json:"RowFilterInfo,omitnil" name:"RowFilterInfo"`
}

type UserIdSetOfWorkGroupId struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 用户Id集合，和CAM侧Uin匹配
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

type UserInfo struct {
	// 用户Id，和子用户uin相同
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitnil" name:"UserDescription"`

	// 单独给用户绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 创建时间，格式如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 关联的工作组集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupSet []*WorkGroupMessage `json:"WorkGroupSet,omitnil" name:"WorkGroupSet"`

	// 是否是主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOwner *bool `json:"IsOwner,omitnil" name:"IsOwner"`

	// 用户类型。ADMIN：管理员 COMMON：普通用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitnil" name:"UserType"`

	// 用户别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`
}

type UserMessage struct {
	// 用户Id，和CAM侧子用户Uin匹配
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitnil" name:"UserDescription"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 当前用户的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 用户别名
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`
}

type UserRole struct {
	// 角色ID
	RoleId *int64 `json:"RoleId,omitnil" name:"RoleId"`

	// 用户app ID
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 用户ID
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 角色权限
	Arn *string `json:"Arn,omitnil" name:"Arn"`

	// 最近修改时间戳
	ModifyTime *int64 `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 角色描述信息
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 创建者UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// cos授权路径列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPermissionList []*CosPermission `json:"CosPermissionList,omitnil" name:"CosPermissionList"`

	// cam策略json
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionJson *string `json:"PermissionJson,omitnil" name:"PermissionJson"`
}

type Users struct {
	// 用户信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSet []*UserMessage `json:"UserSet,omitnil" name:"UserSet"`

	// 用户总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type ViewBaseInfo struct {
	// 该视图所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// 视图名称
	ViewName *string `json:"ViewName,omitnil" name:"ViewName"`

	// 视图创建人昵称
	UserAlias *string `json:"UserAlias,omitnil" name:"UserAlias"`

	// 视图创建人ID
	UserSubUin *string `json:"UserSubUin,omitnil" name:"UserSubUin"`
}

type ViewResponseInfo struct {
	// 视图基本信息。
	ViewBaseInfo *ViewBaseInfo `json:"ViewBaseInfo,omitnil" name:"ViewBaseInfo"`

	// 视图列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil" name:"Columns"`

	// 视图属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil" name:"Properties"`

	// 视图创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 视图更新时间。
	ModifiedTime *string `json:"ModifiedTime,omitnil" name:"ModifiedTime"`
}

type WorkGroupDetailInfo struct {
	// 工作组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 工作组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupName *string `json:"WorkGroupName,omitnil" name:"WorkGroupName"`

	// 包含的信息类型。User：用户信息；DataAuth：数据权限；EngineAuth:引擎权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 工作组上绑定的用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserInfo *Users `json:"UserInfo,omitnil" name:"UserInfo"`

	// 数据权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataPolicyInfo *Policys `json:"DataPolicyInfo,omitnil" name:"DataPolicyInfo"`

	// 引擎权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnginePolicyInfo *Policys `json:"EnginePolicyInfo,omitnil" name:"EnginePolicyInfo"`

	// 工作组描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil" name:"WorkGroupDescription"`

	// 行过滤信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowFilterInfo *Policys `json:"RowFilterInfo,omitnil" name:"RowFilterInfo"`
}

type WorkGroupIdSetOfUserId struct {
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// 工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil" name:"WorkGroupIds"`
}

type WorkGroupInfo struct {
	// 查询到的工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil" name:"WorkGroupName"`

	// 工作组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil" name:"WorkGroupDescription"`

	// 工作组关联的用户数量
	UserNum *int64 `json:"UserNum,omitnil" name:"UserNum"`

	// 工作组关联的用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSet []*UserMessage `json:"UserSet,omitnil" name:"UserSet"`

	// 工作组绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitnil" name:"PolicySet"`

	// 工作组的创建人
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 工作组的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type WorkGroupMessage struct {
	// 工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil" name:"WorkGroupName"`

	// 工作组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil" name:"WorkGroupDescription"`

	// 创建者
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 工作组创建的时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type WorkGroups struct {
	// 工作组信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupSet []*WorkGroupMessage `json:"WorkGroupSet,omitnil" name:"WorkGroupSet"`

	// 工作组总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}