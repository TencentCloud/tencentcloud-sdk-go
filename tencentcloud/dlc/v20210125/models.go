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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddDMSPartitionsRequestParams struct {
	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`
}

type AddDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`
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
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 分区值
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitempty" name:"AddInfo"`
}

type AddUsersToWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要操作的工作组和用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitempty" name:"AddInfo"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CurrentName *string `json:"CurrentName,omitempty" name:"CurrentName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 路径
	Location *string `json:"Location,omitempty" name:"Location"`

	// 基础对象
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`
}

type AlterDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称
	CurrentName *string `json:"CurrentName,omitempty" name:"CurrentName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 路径
	Location *string `json:"Location,omitempty" name:"Location"`

	// 基础对象
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CurrentDbName *string `json:"CurrentDbName,omitempty" name:"CurrentDbName"`

	// 当前名称，变更前table名称
	CurrentTableName *string `json:"CurrentTableName,omitempty" name:"CurrentTableName"`

	// 当前名称，变更前Part名称
	CurrentValues *string `json:"CurrentValues,omitempty" name:"CurrentValues"`

	// 分区
	Partition *DMSPartition `json:"Partition,omitempty" name:"Partition"`
}

type AlterDMSPartitionRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称，变更前db名称
	CurrentDbName *string `json:"CurrentDbName,omitempty" name:"CurrentDbName"`

	// 当前名称，变更前table名称
	CurrentTableName *string `json:"CurrentTableName,omitempty" name:"CurrentTableName"`

	// 当前名称，变更前Part名称
	CurrentValues *string `json:"CurrentValues,omitempty" name:"CurrentValues"`

	// 分区
	Partition *DMSPartition `json:"Partition,omitempty" name:"Partition"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CurrentName *string `json:"CurrentName,omitempty" name:"CurrentName"`

	// 当前数据库名称
	CurrentDbName *string `json:"CurrentDbName,omitempty" name:"CurrentDbName"`

	// 基础对象
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 当前表名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type AlterDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称
	CurrentName *string `json:"CurrentName,omitempty" name:"CurrentName"`

	// 当前数据库名称
	CurrentDbName *string `json:"CurrentDbName,omitempty" name:"CurrentDbName"`

	// 基础对象
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 当前表名
	Name *string `json:"Name,omitempty" name:"Name"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 对象GUID值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guid *string `json:"Guid,omitempty" name:"Guid"`

	// 数据目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 对象owner
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 对象owner账户
	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`

	// 权限
	PermValues []*KVPair `json:"PermValues,omitempty" name:"PermValues"`

	// 附加属性
	Params []*KVPair `json:"Params,omitempty" name:"Params"`

	// 附加业务属性
	BizParams []*KVPair `json:"BizParams,omitempty" name:"BizParams"`

	// 数据版本
	DataVersion *int64 `json:"DataVersion,omitempty" name:"DataVersion"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 数据源主键
	DatasourceId *int64 `json:"DatasourceId,omitempty" name:"DatasourceId"`
}

// Predefined struct for user
type AttachUserPolicyRequestParams struct {
	// 用户Id，和子用户uin相同，需要先使用CreateUser接口创建用户。可以使用DescribeUsers接口查看。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和子用户uin相同，需要先使用CreateUser接口创建用户。可以使用DescribeUsers接口查看。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 要绑定的策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

type AttachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 要绑定的策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type BindWorkGroupsToUserRequestParams struct {
	// 绑定的用户和工作组信息
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitempty" name:"AddInfo"`
}

type BindWorkGroupsToUserRequest struct {
	*tchttp.BaseRequest
	
	// 绑定的用户和工作组信息
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitempty" name:"AddInfo"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CodeCompress *string `json:"CodeCompress,omitempty" name:"CodeCompress"`

	// CSV序列化及反序列化数据结构。
	CSVSerde *CSVSerde `json:"CSVSerde,omitempty" name:"CSVSerde"`

	// 标题行，默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadLines *int64 `json:"HeadLines,omitempty" name:"HeadLines"`

	// 格式，默认值为CSV
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitempty" name:"Format"`
}

type CSVSerde struct {
	// CSV序列化转义符，默认为"\\"，最长8个字符，如 Escape: "/\"
	Escape *string `json:"Escape,omitempty" name:"Escape"`

	// CSV序列化字段域符，默认为"'"，最长8个字符, 如 Quote: "\""
	Quote *string `json:"Quote,omitempty" name:"Quote"`

	// CSV序列化分隔符，默认为"\t"，最长8个字符, 如 Separator: "\t"
	Separator *string `json:"Separator,omitempty" name:"Separator"`
}

// Predefined struct for user
type CancelNotebookSessionStatementBatchRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type CancelNotebookSessionStatementBatchRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitempty" name:"StatementId"`
}

type CancelNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitempty" name:"StatementId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CancelTaskRequestParams struct {
	// 任务Id，全局唯一
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id，全局唯一
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CheckLockMetaDataRequestParams struct {
	// 锁ID
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitempty" name:"TxnId"`

	// 过期时间ms
	ElapsedMs *int64 `json:"ElapsedMs,omitempty" name:"ElapsedMs"`
}

type CheckLockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 锁ID
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitempty" name:"TxnId"`

	// 过期时间ms
	ElapsedMs *int64 `json:"ElapsedMs,omitempty" name:"ElapsedMs"`
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
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 锁状态：ACQUIRED、WAITING、ABORT、NOT_ACQUIRED
	LockState *string `json:"LockState,omitempty" name:"LockState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 列类型，支持如下类型定义:
	// string|tinyint|smallint|int|bigint|boolean|float|double|decimal|timestamp|date|binary|array<data_type>|map<primitive_type, data_type>|struct<col_name : data_type [COMMENT col_comment], ...>|uniontype<data_type, data_type, ...>。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 对该类的注释。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 表示整个 numeric 的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Precision *int64 `json:"Precision,omitempty" name:"Precision"`

	// 表示小数部分的长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitempty" name:"Scale"`

	// 是否为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nullable *string `json:"Nullable,omitempty" name:"Nullable"`

	// 字段位置，小的在前
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitempty" name:"Position"`

	// 字段创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 字段修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 是否为分区字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitempty" name:"IsPartition"`
}

// Predefined struct for user
type CreateDMSDatabaseRequestParams struct {
	// 基础元数据对象
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// Schema目录
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// Db存储路径
	Location *string `json:"Location,omitempty" name:"Location"`

	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 基础元数据对象
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// Schema目录
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// Db存储路径
	Location *string `json:"Location,omitempty" name:"Location"`

	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 基础对象
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 集群类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 计费模式 0=共享模式 1=按量计费 2=包年包月
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 是否自动启动集群
	AutoResume *bool `json:"AutoResume,omitempty" name:"AutoResume"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitempty" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitempty" name:"MaxClusters"`

	// 是否为默虚拟集群
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitempty" name:"DefaultDataEngine"`

	// VPC网段
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 描述信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 集群规模
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 计费类型，后付费：0，预付费：1。当前只支持后付费，不填默认为后付费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 资源使用时长，后付费：固定填3600，预付费：最少填1，代表购买资源一个月，最长不超过120。默认3600
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源使用时长的单位，后付费：s，预付费：m。默认为s
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 资源的自动续费标志。后付费无需续费，固定填0；预付费下：0表示手动续费、1代表自动续费、2代表不续费，在0下如果是大客户，会自动帮大客户续费。默认为0
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 创建资源的时候需要绑定的标签信息
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitempty" name:"EngineExecType"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitempty" name:"TolerableQueueTime"`

	// 集群自动挂起时间，默认10分钟
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitempty" name:"AutoSuspendTime"`

	// 资源类型。Standard_CU：标准型；Memory_CU：内存型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 集群高级配置
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitempty" name:"DataEngineConfigPairs"`

	// 集群镜像版本名字。如SuperSQL-P 1.1;SuperSQL-S 3.2等,不传，默认创建最新镜像版本的集群
	ImageVersionName *string `json:"ImageVersionName,omitempty" name:"ImageVersionName"`

	// 主集群名称
	MainClusterName *string `json:"MainClusterName,omitempty" name:"MainClusterName"`
}

type CreateDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型spark/presto
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 集群类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 计费模式 0=共享模式 1=按量计费 2=包年包月
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 是否自动启动集群
	AutoResume *bool `json:"AutoResume,omitempty" name:"AutoResume"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitempty" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitempty" name:"MaxClusters"`

	// 是否为默虚拟集群
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitempty" name:"DefaultDataEngine"`

	// VPC网段
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 描述信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 集群规模
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 计费类型，后付费：0，预付费：1。当前只支持后付费，不填默认为后付费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 资源使用时长，后付费：固定填3600，预付费：最少填1，代表购买资源一个月，最长不超过120。默认3600
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源使用时长的单位，后付费：s，预付费：m。默认为s
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 资源的自动续费标志。后付费无需续费，固定填0；预付费下：0表示手动续费、1代表自动续费、2代表不续费，在0下如果是大客户，会自动帮大客户续费。默认为0
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 创建资源的时候需要绑定的标签信息
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitempty" name:"EngineExecType"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitempty" name:"TolerableQueueTime"`

	// 集群自动挂起时间，默认10分钟
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitempty" name:"AutoSuspendTime"`

	// 资源类型。Standard_CU：标准型；Memory_CU：内存型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 集群高级配置
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitempty" name:"DataEngineConfigPairs"`

	// 集群镜像版本名字。如SuperSQL-P 1.1;SuperSQL-S 3.2等,不传，默认创建最新镜像版本的集群
	ImageVersionName *string `json:"ImageVersionName,omitempty" name:"ImageVersionName"`

	// 主集群名称
	MainClusterName *string `json:"MainClusterName,omitempty" name:"MainClusterName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataEngineResponseParams struct {
	// 虚拟引擎id
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库基础信息
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
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
	Execution *Execution `json:"Execution,omitempty" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 导出任务输入配置
	InputConf []*KVPair `json:"InputConf,omitempty" name:"InputConf"`

	// 导出任务输出配置
	OutputConf []*KVPair `json:"OutputConf,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导出到cos
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
}

type CreateExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据来源，lakefsStorage、taskResult
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 导出任务输入配置
	InputConf []*KVPair `json:"InputConf,omitempty" name:"InputConf"`

	// 导出任务输出配置
	OutputConf []*KVPair `json:"OutputConf,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导出到cos
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 输入配置
	InputConf []*KVPair `json:"InputConf,omitempty" name:"InputConf"`

	// 输出配置
	OutputConf []*KVPair `json:"OutputConf,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导入到托管存储，即lakefsStorage
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
}

type CreateImportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据来源，cos
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 输入配置
	InputConf []*KVPair `json:"InputConf,omitempty" name:"InputConf"`

	// 输出配置
	OutputConf []*KVPair `json:"OutputConf,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导入到托管存储，即lakefsStorage
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
}

type CreateInternalTableRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
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
	Execution *string `json:"Execution,omitempty" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// session文件地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitempty" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitempty" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitempty" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	ProgramArchives []*string `json:"ProgramArchives,omitempty" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitempty" name:"ExecutorNumbers"`

	// Session相关配置，当前支持：dlc.eni、dlc.role.arn、dlc.sql.set.config以及用户指定的配置，注：roleArn必填；
	Arguments []*KVPair `json:"Arguments,omitempty" name:"Arguments"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitempty" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitempty" name:"TimeoutInSecond"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitempty" name:"ExecutorMaxNumbers"`
}

type CreateNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// session文件地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitempty" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitempty" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitempty" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	ProgramArchives []*string `json:"ProgramArchives,omitempty" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitempty" name:"ExecutorNumbers"`

	// Session相关配置，当前支持：dlc.eni、dlc.role.arn、dlc.sql.set.config以及用户指定的配置，注：roleArn必填；
	Arguments []*KVPair `json:"Arguments,omitempty" name:"Arguments"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitempty" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitempty" name:"TimeoutInSecond"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitempty" name:"ExecutorMaxNumbers"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionResponseParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Spark任务返回的AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitempty" name:"State"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

type CreateNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`
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
	NotebookSessionStatement *NotebookSessionStatementInfo `json:"NotebookSessionStatement,omitempty" name:"NotebookSessionStatement"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 是否保存运行结果
	SaveResult *bool `json:"SaveResult,omitempty" name:"SaveResult"`
}

type CreateNotebookSessionStatementSupportBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 是否保存运行结果
	SaveResult *bool `json:"SaveResult,omitempty" name:"SaveResult"`
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
	NotebookSessionStatementBatches *NotebookSessionStatementBatchInformation `json:"NotebookSessionStatementBatches,omitempty" name:"NotebookSessionStatementBatches"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 下载格式
	Format *string `json:"Format,omitempty" name:"Format"`

	// 是否重新生成下载文件，仅当之前任务为 Timout | Error 时有效
	Force *bool `json:"Force,omitempty" name:"Force"`
}

type CreateResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 查询结果任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 下载格式
	Format *string `json:"Format,omitempty" name:"Format"`

	// 是否重新生成下载文件，仅当之前任务为 Timout | Error 时有效
	Force *bool `json:"Force,omitempty" name:"Force"`
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
	DownloadId *string `json:"DownloadId,omitempty" name:"DownloadId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitempty" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitempty" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

type CreateScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本名称，最大不能超过255个字符。
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitempty" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitempty" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1代表spark jar应用，2代表spark streaming应用
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// spark应用的执行入口
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// 执行spark作业的角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业driver资源规格大小, 可取small,medium,large,xlarge
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// spark作业executor资源规格大小, 可取small,medium,large,xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 是否本地上传，可去cos,lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// spark jar作业时的主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// 是否本地上传，包含cos,lakefs
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark jar作业依赖jars，以逗号分隔
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// 是否本地上传，包含cos,lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖资源，以逗号分隔
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// spark作业命令行参数
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// 只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// archives：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// archives：依赖资源
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// Spark Image 版本
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
}

type CreateSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1代表spark jar应用，2代表spark streaming应用
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// spark应用的执行入口
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// 执行spark作业的角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业driver资源规格大小, 可取small,medium,large,xlarge
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// spark作业executor资源规格大小, 可取small,medium,large,xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 是否本地上传，可去cos,lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// spark jar作业时的主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// 是否本地上传，包含cos,lakefs
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark jar作业依赖jars，以逗号分隔
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// 是否本地上传，包含cos,lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖资源，以逗号分隔
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// spark作业命令行参数
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// 只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// archives：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// archives：依赖资源
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// Spark Image 版本
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppResponseParams struct {
	// App唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// spark作业的命令行参数，以空格分隔；一般用于周期性调用使用
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`
}

type CreateSparkAppTaskRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// spark作业的命令行参数，以空格分隔；一般用于周期性调用使用
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`
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
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateStoreLocationRequestParams struct {
	// 计算结果存储cos路径，如：cosn://bucketname/
	StoreLocation *string `json:"StoreLocation,omitempty" name:"StoreLocation"`
}

type CreateStoreLocationRequest struct {
	*tchttp.BaseRequest
	
	// 计算结果存储cos路径，如：cosn://bucketname/
	StoreLocation *string `json:"StoreLocation,omitempty" name:"StoreLocation"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TableInfo *TableInfo `json:"TableInfo,omitempty" name:"TableInfo"`
}

type CreateTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据表配置信息
	TableInfo *TableInfo `json:"TableInfo,omitempty" name:"TableInfo"`
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
	Execution *Execution `json:"Execution,omitempty" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Task *Task `json:"Task,omitempty" name:"Task"`

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 默认数据源名称。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 数据引擎名称，不填提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 计算任务，该参数中包含任务类型及其相关配置信息
	Task *Task `json:"Task,omitempty" name:"Task"`

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 默认数据源名称。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 数据引擎名称，不填提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 数据源名称，默认为COSDataCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
}

type CreateTasksInOrderRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 数据源名称，默认为COSDataCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
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
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务Id集合，按照执行顺序排列
	TaskIdSet []*string `json:"TaskIdSet,omitempty" name:"TaskIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 计算引擎名称，不填任务提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

type CreateTasksRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 计算引擎名称，不填任务提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
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
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 任务Id集合，按照执行顺序排列
	TaskIdSet []*string `json:"TaskIdSet,omitempty" name:"TaskIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`

	// 绑定到用户的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 用户类型。ADMIN：管理员 COMMON：一般用户。当用户类型为管理员的时候，不能设置权限集合和绑定的工作组集合，管理员默认拥有所有权限。该参数不填默认为COMMON
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 绑定到用户的工作组ID集合。
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`

	// 用户别名，字符长度小50
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 需要授权的子用户uin，可以通过腾讯云控制台右上角 → “账号信息” → “账号ID进行查看”。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`

	// 绑定到用户的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 用户类型。ADMIN：管理员 COMMON：一般用户。当用户类型为管理员的时候，不能设置权限集合和绑定的工作组集合，管理员默认拥有所有权限。该参数不填默认为COMMON
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 绑定到用户的工作组ID集合。
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`

	// 用户别名，字符长度小50
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	WorkGroupName *string `json:"WorkGroupName,omitempty" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`

	// 工作组绑定的鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 需要绑定到工作组的用户Id集合
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type CreateWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitempty" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`

	// 工作组绑定的鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 需要绑定到工作组的用户Id集合
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
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
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResumeTime *string `json:"ResumeTime,omitempty" name:"ResumeTime"`

	// 定时挂起时间：如：周一20点
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuspendTime *string `json:"SuspendTime,omitempty" name:"SuspendTime"`

	// 挂起配置：0（默认）：等待任务结束后挂起、1：强制挂起
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuspendStrategy *int64 `json:"SuspendStrategy,omitempty" name:"SuspendStrategy"`
}

type DMSColumn struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitempty" name:"Position"`

	// 附加参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*KVPair `json:"Params,omitempty" name:"Params"`

	// 业务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizParams []*KVPair `json:"BizParams,omitempty" name:"BizParams"`

	// 是否分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitempty" name:"IsPartition"`
}

type DMSColumnOrder struct {
	// 列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Col *string `json:"Col,omitempty" name:"Col"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitempty" name:"Order"`
}

type DMSPartition struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据目录名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 数据版本
	DataVersion *int64 `json:"DataVersion,omitempty" name:"DataVersion"`

	// 分区名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值列表
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitempty" name:"LastAccessTime"`

	// 附件属性
	Params []*KVPair `json:"Params,omitempty" name:"Params"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitempty" name:"Sds"`
}

type DMSSds struct {
	// 存储地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 输入格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputFormat *string `json:"InputFormat,omitempty" name:"InputFormat"`

	// 输出格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputFormat *string `json:"OutputFormat,omitempty" name:"OutputFormat"`

	// bucket数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumBuckets *int64 `json:"NumBuckets,omitempty" name:"NumBuckets"`

	// 是是否压缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compressed *bool `json:"Compressed,omitempty" name:"Compressed"`

	// 是否有子目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoredAsSubDirectories *bool `json:"StoredAsSubDirectories,omitempty" name:"StoredAsSubDirectories"`

	// 序列化lib
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeLib *string `json:"SerdeLib,omitempty" name:"SerdeLib"`

	// 序列化名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeName *string `json:"SerdeName,omitempty" name:"SerdeName"`

	// 桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketCols []*string `json:"BucketCols,omitempty" name:"BucketCols"`

	// 序列化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeParams []*KVPair `json:"SerdeParams,omitempty" name:"SerdeParams"`

	// 附加参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*KVPair `json:"Params,omitempty" name:"Params"`

	// 列排序(Expired)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortCols *DMSColumnOrder `json:"SortCols,omitempty" name:"SortCols"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cols []*DMSColumn `json:"Cols,omitempty" name:"Cols"`

	// 列排序字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortColumns []*DMSColumnOrder `json:"SortColumns,omitempty" name:"SortColumns"`
}

type DMSTable struct {
	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" name:"ViewOriginalText"`

	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" name:"ViewExpandedText"`

	// hive维护版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitempty" name:"Retention"`

	// 存储对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sds *DMSSds `json:"Sds,omitempty" name:"Sds"`

	// 分区列
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitempty" name:"PartitionKeys"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeTime *int64 `json:"LifeTime,omitempty" name:"LifeTime"`

	// 最后访问时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAccessTime *string `json:"LastAccessTime,omitempty" name:"LastAccessTime"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUpdateTime *string `json:"DataUpdateTime,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StructUpdateTime *string `json:"StructUpdateTime,omitempty" name:"StructUpdateTime"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*DMSColumn `json:"Columns,omitempty" name:"Columns"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DMSTableInfo struct {
	// DMS表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *DMSTable `json:"Table,omitempty" name:"Table"`

	// 基础对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`
}

type DataEngineConfigPair struct {

}

type DataEngineInfo struct {
	// DataEngine名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 引擎类型 spark/presto
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 集群资源类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 引用ID
	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`

	// 数据引擎状态  -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中
	State *int64 `json:"State,omitempty" name:"State"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 计费模式 0共享模式 1按量计费 2包年包月
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 最小集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinClusters *int64 `json:"MinClusters,omitempty" name:"MinClusters"`

	// 最大集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxClusters *int64 `json:"MaxClusters,omitempty" name:"MaxClusters"`

	// 是否自动恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoResume *bool `json:"AutoResume,omitempty" name:"AutoResume"`

	// 自动恢复时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpendAfter *int64 `json:"SpendAfter,omitempty" name:"SpendAfter"`

	// 集群网段
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 是否为默认引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitempty" name:"DefaultDataEngine"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`

	// 操作者
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 隔离时间
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 冲正时间
	ReversalTime *string `json:"ReversalTime,omitempty" name:"ReversalTime"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// 标签对集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagInfo `json:"TagList,omitempty" name:"TagList"`

	// 引擎拥有的权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSuspend *bool `json:"AutoSuspend,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，有效值：SQL/BATCH
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineExecType *string `json:"EngineExecType,omitempty" name:"EngineExecType"`

	// 自动续费标志，0，初始状态，默认不自动续费，若用户有预付费不停服特权，自动续费。1：自动续费。2：明确不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 集群自动挂起时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitempty" name:"AutoSuspendTime"`

	// 网络连接配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionSet []*NetworkConnection `json:"NetworkConnectionSet,omitempty" name:"NetworkConnectionSet"`

	// ui的跳转地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	UiURL *string `json:"UiURL,omitempty" name:"UiURL"`

	// 引擎的资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 集群镜像版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersionId *string `json:"ImageVersionId,omitempty" name:"ImageVersionId"`

	// 集群镜像小版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildImageVersionId *string `json:"ChildImageVersionId,omitempty" name:"ChildImageVersionId"`

	// 集群镜像版本名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersionName *string `json:"ImageVersionName,omitempty" name:"ImageVersionName"`

	// 是否开启备集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitempty" name:"StartStandbyCluster"`
}

type DataFormat struct {
	// 文本格式，TextFile。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextFile *TextFile `json:"TextFile,omitempty" name:"TextFile"`

	// 文本格式，CSV。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CSV *CSV `json:"CSV,omitempty" name:"CSV"`

	// 文本格式，Json。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *Other `json:"Json,omitempty" name:"Json"`

	// Parquet格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parquet *Other `json:"Parquet,omitempty" name:"Parquet"`

	// ORC格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ORC *Other `json:"ORC,omitempty" name:"ORC"`

	// AVRO格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AVRO *Other `json:"AVRO,omitempty" name:"AVRO"`
}

type DataGovernPolicy struct {

}

type DatabaseInfo struct {
	// 数据库名称，长度0~128，支持数字、字母下划线，不允许数字大头，统一转换为小写。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~500。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 数据库属性列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 数据库cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`
}

type DatabaseResponseInfo struct {
	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~256。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 允许针对数据库的属性元数据信息进行指定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 数据库创建时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据库更新时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// cos存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 建库用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// 建库用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSubUin *string `json:"UserSubUin,omitempty" name:"UserSubUin"`

	// 数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernPolicy *DataGovernPolicy `json:"GovernPolicy,omitempty" name:"GovernPolicy"`

	// 数据库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitempty" name:"DatabaseId"`
}

// Predefined struct for user
type DeleteNotebookSessionRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type DeleteNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ScriptIds []*string `json:"ScriptIds,omitempty" name:"ScriptIds"`
}

type DeleteScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本id，其可以通过DescribeScripts接口提取
	ScriptIds []*string `json:"ScriptIds,omitempty" name:"ScriptIds"`
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
	ScriptsAffected *int64 `json:"ScriptsAffected,omitempty" name:"ScriptsAffected"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

type DeleteSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的用户的Id
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitempty" name:"AddInfo"`
}

type DeleteUsersFromWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitempty" name:"AddInfo"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`
}

type DeleteWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDMSDatabaseRequestParams struct {
	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 匹配规则
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
}

type DescribeDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 匹配规则
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 存储地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 数据对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 单个分区名称，精准匹配
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 多个分区名称，精准匹配
	PartitionNames []*string `json:"PartitionNames,omitempty" name:"PartitionNames"`

	// 多个分区字段的匹配，模糊匹配
	PartValues []*string `json:"PartValues,omitempty" name:"PartValues"`

	// 过滤SQL
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 最大分区数量
	MaxParts *int64 `json:"MaxParts,omitempty" name:"MaxParts"`

	// 翻页跳过数量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页面数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 表达式
	Expression *string `json:"Expression,omitempty" name:"Expression"`
}

type DescribeDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 单个分区名称，精准匹配
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 多个分区名称，精准匹配
	PartitionNames []*string `json:"PartitionNames,omitempty" name:"PartitionNames"`

	// 多个分区字段的匹配，模糊匹配
	PartValues []*string `json:"PartValues,omitempty" name:"PartValues"`

	// 过滤SQL
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 最大分区数量
	MaxParts *int64 `json:"MaxParts,omitempty" name:"MaxParts"`

	// 翻页跳过数量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页面数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 表达式
	Expression *string `json:"Expression,omitempty" name:"Expression"`
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
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`
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
	Asset *Asset `json:"Asset,omitempty" name:"Asset"`

	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" name:"ViewOriginalText"`

	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" name:"ViewExpandedText"`

	// hive维护版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitempty" name:"Retention"`

	// 存储对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sds *DMSSds `json:"Sds,omitempty" name:"Sds"`

	// 分区列
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitempty" name:"PartitionKeys"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*DMSPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// Schame名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// 生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeTime *int64 `json:"LifeTime,omitempty" name:"LifeTime"`

	// 最后访问时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAccessTime *string `json:"LastAccessTime,omitempty" name:"LastAccessTime"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUpdateTime *string `json:"DataUpdateTime,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StructUpdateTime *string `json:"StructUpdateTime,omitempty" name:"StructUpdateTime"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*DMSColumn `json:"Columns,omitempty" name:"Columns"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 筛选参数：更新开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 筛选参数：更新结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段：create_time：创建时间
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序字段：true：升序（默认），false：降序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`
}

type DescribeDMSTablesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据目录
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询模式
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 筛选参数：更新开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 筛选参数：更新结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段：create_time：创建时间
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序字段：true：升序（默认），false：降序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`
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
	TableList []*DMSTableInfo `json:"TableList,omitempty" name:"TableList"`

	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDataEnginesRequestParams struct {
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 滤类型，传参Name应为以下其中一个,
	// data-engine-name - String 
	// engine-type - String
	// state - String 
	// mode - String 
	// create-time - String 
	// message - String
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 已废弃，请使用DatasourceConnectionNameSet
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 是否不返回共享引擎，true不返回共享引擎，false可以返回共享引擎
	ExcludePublicEngine *bool `json:"ExcludePublicEngine,omitempty" name:"ExcludePublicEngine"`

	// 参数应该为引擎权限类型，有效类型："USE", "MODIFY", "OPERATE", "MONITOR", "DELETE"
	AccessTypes []*string `json:"AccessTypes,omitempty" name:"AccessTypes"`

	// 引擎执行任务类型，有效值：SQL/BATCH
	EngineExecType *string `json:"EngineExecType,omitempty" name:"EngineExecType"`

	// 引擎类型，有效值：spark/presto
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 网络配置列表，若传入该参数，则返回网络配置关联的计算引擎
	DatasourceConnectionNameSet []*string `json:"DatasourceConnectionNameSet,omitempty" name:"DatasourceConnectionNameSet"`
}

type DescribeDataEnginesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 滤类型，传参Name应为以下其中一个,
	// data-engine-name - String 
	// engine-type - String
	// state - String 
	// mode - String 
	// create-time - String 
	// message - String
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 已废弃，请使用DatasourceConnectionNameSet
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 是否不返回共享引擎，true不返回共享引擎，false可以返回共享引擎
	ExcludePublicEngine *bool `json:"ExcludePublicEngine,omitempty" name:"ExcludePublicEngine"`

	// 参数应该为引擎权限类型，有效类型："USE", "MODIFY", "OPERATE", "MONITOR", "DELETE"
	AccessTypes []*string `json:"AccessTypes,omitempty" name:"AccessTypes"`

	// 引擎执行任务类型，有效值：SQL/BATCH
	EngineExecType *string `json:"EngineExecType,omitempty" name:"EngineExecType"`

	// 引擎类型，有效值：spark/presto
	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`

	// 网络配置列表，若传入该参数，则返回网络配置关联的计算引擎
	DatasourceConnectionNameSet []*string `json:"DatasourceConnectionNameSet,omitempty" name:"DatasourceConnectionNameSet"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEnginesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginesResponseParams struct {
	// 数据引擎列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngines []*DataEngineInfo `json:"DataEngines,omitempty" name:"DataEngines"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 数据源唯名称，该名称可以通过DescribeDatasourceConnection接口查询到。默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 排序字段，当前版本仅支持按库名排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序类型：false：降序（默认）、true：升序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 数据源唯名称，该名称可以通过DescribeDatasourceConnection接口查询到。默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 排序字段，当前版本仅支持按库名排序
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序类型：false：降序（默认）、true：升序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`
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
	DatabaseList []*DatabaseResponseInfo `json:"DatabaseList,omitempty" name:"DatabaseList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeEngineUsageInfoRequestParams struct {
	// House Id
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`
}

type DescribeEngineUsageInfoRequest struct {
	*tchttp.BaseRequest
	
	// House Id
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`
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
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 已占用集群规格
	Used *int64 `json:"Used,omitempty" name:"Used"`

	// 剩余集群规格
	Available *int64 `json:"Available,omitempty" name:"Available"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeNotebookSessionLogRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeNotebookSessionLogRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	Logs []*string `json:"Logs,omitempty" name:"Logs"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type DescribeNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
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
	Session *NotebookSessionInfo `json:"Session,omitempty" name:"Session"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitempty" name:"StatementId"`
}

type DescribeNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitempty" name:"StatementId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementResponseParams struct {
	// Session Statement详情
	NotebookSessionStatement *NotebookSessionStatementInfo `json:"NotebookSessionStatement,omitempty" name:"NotebookSessionStatement"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

type DescribeNotebookSessionStatementSqlResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 结果数据
	ResultSet *string `json:"ResultSet,omitempty" name:"ResultSet"`

	// schema
	ResultSchema []*Column `json:"ResultSchema,omitempty" name:"ResultSchema"`

	// 分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 存储结果地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 批任务id
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type DescribeNotebookSessionStatementsRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 批任务id
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
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
	NotebookSessionStatements *NotebookSessionStatementBatchInformation `json:"NotebookSessionStatements,omitempty" name:"NotebookSessionStatements"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State []*string `json:"State,omitempty" name:"State"`

	// 排序字段（默认按创建时间）
	SortFields []*string `json:"SortFields,omitempty" name:"SortFields"`

	// 排序字段：true：升序、false：降序（默认）
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// 分页字段
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页字段
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeNotebookSessionsRequest struct {
	*tchttp.BaseRequest
	
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State []*string `json:"State,omitempty" name:"State"`

	// 排序字段（默认按创建时间）
	SortFields []*string `json:"SortFields,omitempty" name:"SortFields"`

	// 排序字段：true：升序、false：降序（默认）
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// 分页字段
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页字段
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalElements *int64 `json:"TotalElements,omitempty" name:"TotalElements"`

	// 总页数
	TotalPages *int64 `json:"TotalPages,omitempty" name:"TotalPages"`

	// 当前页码
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 当前页数量
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// session列表信息
	Sessions []*NotebookSessions `json:"Sessions,omitempty" name:"Sessions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DownloadId *string `json:"DownloadId,omitempty" name:"DownloadId"`
}

type DescribeResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 查询任务Id
	DownloadId *string `json:"DownloadId,omitempty" name:"DownloadId"`
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
	Path *string `json:"Path,omitempty" name:"Path"`

	// 任务状态 init | queue | format | compress | success|  timeout | error
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务异常原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 临时AK
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// 临时SK
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 临时Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序，默认asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeScriptsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序，默认asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	Scripts []*Script `json:"Scripts,omitempty" name:"Scripts"`

	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// spark作业Id，与JobName同时存在时，JobName无效
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

type DescribeSparkAppJobRequest struct {
	*tchttp.BaseRequest
	
	// spark作业Id，与JobName同时存在时，JobName无效
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitempty" name:"JobName"`
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
	Job *SparkJobInfo `json:"Job,omitempty" name:"Job"`

	// 查询的spark作业是否存在
	IsExists *bool `json:"IsExists,omitempty" name:"IsExists"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 按照该参数过滤,支持spark-job-name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 更新时间起始点
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间截止点
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表限制数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSparkAppJobsRequest struct {
	*tchttp.BaseRequest
	
	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 按照该参数过滤,支持spark-job-name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 更新时间起始点
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间截止点
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表限制数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
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
	SparkAppJobs []*SparkJobInfo `json:"SparkAppJobs,omitempty" name:"SparkAppJobs"`

	// spark作业总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 执行实例id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 更新时间起始点
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间截止点
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 按照该参数过滤,支持task-state
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeSparkAppTasksRequest struct {
	*tchttp.BaseRequest
	
	// spark作业Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询Limit
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 执行实例id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 更新时间起始点
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间截止点
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 按照该参数过滤,支持task-state
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	Tasks *TaskResponseInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppTasks []*TaskResponseInfo `json:"SparkAppTasks,omitempty" name:"SparkAppTasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StoreLocation *string `json:"StoreLocation,omitempty" name:"StoreLocation"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 查询表所在的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest
	
	// 查询对象表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 查询表所在的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
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
	Table *TableResponseInfo `json:"Table,omitempty" name:"Table"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTablesRequestParams struct {
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序字段，支持：CreateTime、UpdateTime、StorageSize、RecordCount、Name（不传则默认按name升序）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitempty" name:"TableFormat"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序字段，支持：CreateTime、UpdateTime、StorageSize、RecordCount、Name（不传则默认按name升序）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitempty" name:"TableFormat"`
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
	TableList []*TableResponseInfo `json:"TableList,omitempty" name:"TableList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`
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
	TaskInfo *TaskResultInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	// task-kind - string （任务类型过滤）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time（创建时间，默认）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 支持计算资源名字筛选
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	// task-kind - string （任务类型过滤）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time（创建时间，默认）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 支持计算资源名字筛选
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`
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
	TaskList []*TaskResponseInfo `json:"TaskList,omitempty" name:"TaskList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务概览信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TasksOverview *TasksOverview `json:"TasksOverview,omitempty" name:"TasksOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeUsersRequestParams struct {
	// 指定查询的子用户uin，用户需要通过CreateUser接口创建。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 过滤条件，支持如下字段类型，user-type：根据用户类型过滤。user-keyword：根据用户名称过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 指定查询的子用户uin，用户需要通过CreateUser接口创建。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 过滤条件，支持如下字段类型，user-type：根据用户类型过滤。user-keyword：根据用户名称过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询到的授权用户信息集合
	UserSet []*UserInfo `json:"UserSet,omitempty" name:"UserSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 数据库所属的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 排序字段
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序规则，true:升序；false:降序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// 按视图更新时间筛选，开始时间，如2021-11-11 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 按视图更新时间筛选，结束时间，如2021-11-12 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 数据库所属的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 排序字段
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序规则，true:升序；false:降序
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// 按视图更新时间筛选，开始时间，如2021-11-11 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 按视图更新时间筛选，结束时间，如2021-11-12 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	ViewList []*ViewResponseInfo `json:"ViewList,omitempty" name:"ViewList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeWorkGroupsRequestParams struct {
	// 查询的工作组Id，不填或填0表示不过滤。
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 过滤条件，当前仅支持按照工作组名称进行模糊搜索。Key为workgroup-name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`
}

type DescribeWorkGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的工作组Id，不填或填0表示不过滤。
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 过滤条件，当前仅支持按照工作组名称进行模糊搜索。Key为workgroup-name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 工作组信息集合
	WorkGroupSet []*WorkGroupInfo `json:"WorkGroupSet,omitempty" name:"WorkGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
}

type DetachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitempty" name:"DeleteData"`

	// 是否级联删除
	Cascade *bool `json:"Cascade,omitempty" name:"Cascade"`
}

type DropDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitempty" name:"DeleteData"`

	// 是否级联删除
	Cascade *bool `json:"Cascade,omitempty" name:"Cascade"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 数据表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 分区名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 单个分区名称
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否删除分区数据
	DeleteData *bool `json:"DeleteData,omitempty" name:"DeleteData"`
}

type DropDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 数据表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 分区名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 单个分区名称
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否删除分区数据
	DeleteData *bool `json:"DeleteData,omitempty" name:"DeleteData"`
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
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitempty" name:"DeleteData"`

	// 环境属性
	EnvProps *KVPair `json:"EnvProps,omitempty" name:"EnvProps"`
}

type DropDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitempty" name:"DeleteData"`

	// 环境属性
	EnvProps *KVPair `json:"EnvProps,omitempty" name:"EnvProps"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Execution struct {
	// 自动生成SQL语句。
	SQL *string `json:"SQL,omitempty" name:"SQL"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑或（OR）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type GenerateCreateMangedTableSqlRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
}

type GenerateCreateMangedTableSqlRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateCreateMangedTableSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateCreateMangedTableSqlResponseParams struct {
	// 创建托管存储内表sql语句描述
	Execution *Execution `json:"Execution,omitempty" name:"Execution"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type JobLogResult struct {
	// 日志时间戳，毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 日志topic id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志topic name
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志内容，json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
}

type KVPair struct {
	// 配置的key值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 配置的value值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type ListTaskJobLogDetailRequestParams struct {
	// 列表返回的Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 下一次分页参数，第一次传空
	Context *string `json:"Context,omitempty" name:"Context"`

	// 最近1000条日志是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type ListTaskJobLogDetailRequest struct {
	*tchttp.BaseRequest
	
	// 列表返回的Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 下一次分页参数，第一次传空
	Context *string `json:"Context,omitempty" name:"Context"`

	// 最近1000条日志是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitempty" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskJobLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskJobLogDetailResponseParams struct {
	// 下一次分页参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitempty" name:"Context"`

	// 是否获取完结
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// 日志详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*JobLogResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 分区
	Partition *string `json:"Partition,omitempty" name:"Partition"`

	// 锁类型：SHARED_READ、SHARED_WRITE、EXCLUSIVE
	LockType *string `json:"LockType,omitempty" name:"LockType"`

	// 锁级别：DB、TABLE、PARTITION
	LockLevel *string `json:"LockLevel,omitempty" name:"LockLevel"`

	// 锁操作：SELECT,INSERT,UPDATE,DELETE,UNSET,NO_TXN
	DataOperationType *string `json:"DataOperationType,omitempty" name:"DataOperationType"`

	// 是否保持Acid
	IsAcid *bool `json:"IsAcid,omitempty" name:"IsAcid"`

	// 是否动态分区写
	IsDynamicPartitionWrite *bool `json:"IsDynamicPartitionWrite,omitempty" name:"IsDynamicPartitionWrite"`
}

// Predefined struct for user
type LockMetaDataRequestParams struct {
	// 加锁内容
	LockComponentList []*LockComponentInfo `json:"LockComponentList,omitempty" name:"LockComponentList"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 事务id
	TxnId *int64 `json:"TxnId,omitempty" name:"TxnId"`

	// 客户端信息
	AgentInfo *string `json:"AgentInfo,omitempty" name:"AgentInfo"`

	// 主机名
	Hostname *string `json:"Hostname,omitempty" name:"Hostname"`
}

type LockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 加锁内容
	LockComponentList []*LockComponentInfo `json:"LockComponentList,omitempty" name:"LockComponentList"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 事务id
	TxnId *int64 `json:"TxnId,omitempty" name:"TxnId"`

	// 客户端信息
	AgentInfo *string `json:"AgentInfo,omitempty" name:"AgentInfo"`

	// 主机名
	Hostname *string `json:"Hostname,omitempty" name:"Hostname"`
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
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 锁状态：ACQUIRED、WAITING、ABORT、NOT_ACQUIRED
	LockState *string `json:"LockState,omitempty" name:"LockState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifySparkAppRequestParams struct {
	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1代表spark jar应用，2代表spark streaming应用
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// spark应用的执行入口
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// 执行spark作业的角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业driver资源规格大小, 可取small,medium,large,xlarge
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// spark作业executor资源规格大小, 可取small,medium,large,xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// spark应用Id
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 是否本地上传，可取cos,lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// spark jar作业时的主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// jar资源依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark jar作业依赖jars，以逗号分隔
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// file资源依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖资源，以逗号分隔
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// spark作业命令行参数
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// 只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// archives：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// archives：依赖资源
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// Spark Image 版本
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
}

type ModifySparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 1代表spark jar应用，2代表spark streaming应用
	AppType *int64 `json:"AppType,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// spark应用的执行入口
	AppFile *string `json:"AppFile,omitempty" name:"AppFile"`

	// 执行spark作业的角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业driver资源规格大小, 可取small,medium,large,xlarge
	AppDriverSize *string `json:"AppDriverSize,omitempty" name:"AppDriverSize"`

	// spark作业executor资源规格大小, 可取small,medium,large,xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitempty" name:"AppExecutorNums"`

	// spark应用Id
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 是否本地上传，可取cos,lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// spark jar作业时的主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitempty" name:"AppConf"`

	// jar资源依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark jar作业依赖jars，以逗号分隔
	AppJars *string `json:"AppJars,omitempty" name:"AppJars"`

	// file资源依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖资源，以逗号分隔
	AppFiles *string `json:"AppFiles,omitempty" name:"AppFiles"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// spark作业命令行参数
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// 只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// archives：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// archives：依赖资源
	AppArchives *string `json:"AppArchives,omitempty" name:"AppArchives"`

	// Spark Image 版本
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitempty" name:"AppExecutorMaxNumbers"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyWorkGroupRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`
}

type ModifyWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type NetworkConnection struct {
	// 网络配置id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 网络配置唯一标志符
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateId *string `json:"AssociateId,omitempty" name:"AssociateId"`

	// 计算引擎id
	// 注意：此字段可能返回 null，表示取不到有效值。
	HouseId *string `json:"HouseId,omitempty" name:"HouseId"`

	// 数据源id(已废弃)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionId *string `json:"DatasourceConnectionId,omitempty" name:"DatasourceConnectionId"`

	// 网络配置状态（0-初始化，1-正常）
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitempty" name:"State"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建用户Appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Appid *int64 `json:"Appid,omitempty" name:"Appid"`

	// 计算引擎名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HouseName *string `json:"HouseName,omitempty" name:"HouseName"`

	// 网络配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 网络配置类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionType *int64 `json:"NetworkConnectionType,omitempty" name:"NetworkConnectionType"`

	// 创建用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 创建用户SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`

	// 网络配置描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionDesc *string `json:"NetworkConnectionDesc,omitempty" name:"NetworkConnectionDesc"`

	// 数据源vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionVpcId *string `json:"DatasourceConnectionVpcId,omitempty" name:"DatasourceConnectionVpcId"`

	// 数据源SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionSubnetId *string `json:"DatasourceConnectionSubnetId,omitempty" name:"DatasourceConnectionSubnetId"`

	// 数据源SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionCidrBlock *string `json:"DatasourceConnectionCidrBlock,omitempty" name:"DatasourceConnectionCidrBlock"`

	// 数据源SubnetCidrBlock
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionSubnetCidrBlock *string `json:"DatasourceConnectionSubnetCidrBlock,omitempty" name:"DatasourceConnectionSubnetCidrBlock"`
}

type NotebookSessionInfo struct {
	// Session名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// Session相关配置，当前支持：eni、roleArn以及用户指定的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arguments []*KVPair `json:"Arguments,omitempty" name:"Arguments"`

	// 运行程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitempty" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitempty" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitempty" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramArchives []*string `json:"ProgramArchives,omitempty" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DriverSize *string `json:"DriverSize,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorSize *string `json:"ExecutorSize,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitempty" name:"ExecutorNumbers"`

	// 代理用户，默认为root
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyUser *string `json:"ProxyUser,omitempty" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitempty" name:"TimeoutInSecond"`

	// Spark任务返回的AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitempty" name:"State"`

	// Session创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 其它信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppInfo []*KVPair `json:"AppInfo,omitempty" name:"AppInfo"`

	// Spark ui地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkUiUrl *string `json:"SparkUiUrl,omitempty" name:"SparkUiUrl"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitempty" name:"ExecutorMaxNumbers"`
}

type NotebookSessionStatementBatchInformation struct {
	// 任务详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookSessionStatementBatch []*NotebookSessionStatementInfo `json:"NotebookSessionStatementBatch,omitempty" name:"NotebookSessionStatementBatch"`

	// 当前批任务是否运行完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAvailable *bool `json:"IsAvailable,omitempty" name:"IsAvailable"`

	// Session唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Batch唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type NotebookSessionStatementInfo struct {
	// 完成时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Completed *int64 `json:"Completed,omitempty" name:"Completed"`

	// 开始时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Started *int64 `json:"Started,omitempty" name:"Started"`

	// 完成进度，百分制
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitempty" name:"StatementId"`

	// Session Statement状态，包含：waiting（排队中）、running（运行中）、available（正常）、error（异常）、cancelling（取消中）、cancelled（已取消）
	State *string `json:"State,omitempty" name:"State"`

	// Statement输出信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutPut *StatementOutput `json:"OutPut,omitempty" name:"OutPut"`

	// 批任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 运行语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type NotebookSessions struct {
	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// Session唯一标识
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 代理用户，默认为root
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyUser *string `json:"ProxyUser,omitempty" name:"ProxyUser"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitempty" name:"State"`

	// Spark任务返回的AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppId *string `json:"SparkAppId,omitempty" name:"SparkAppId"`

	// Session名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Session创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 引擎名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 最新的运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRunningTime *string `json:"LastRunningTime,omitempty" name:"LastRunningTime"`

	// 创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// spark ui地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkUiUrl *string `json:"SparkUiUrl,omitempty" name:"SparkUiUrl"`
}

type Other struct {
	// 枚举类型，默认值为Json，可选值为[Json, Parquet, ORC, AVRD]之一。
	Format *string `json:"Format,omitempty" name:"Format"`
}

type Partition struct {
	// 分区列名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 对分区的描述。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 隐式分区转换策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transform *string `json:"Transform,omitempty" name:"Transform"`

	// 转换策略参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformArgs []*string `json:"TransformArgs,omitempty" name:"TransformArgs"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Policy struct {
	// 需要授权的数据库名，填*代表当前Catalog下所有数据库。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定数据库。
	Database *string `json:"Database,omitempty" name:"Database"`

	// 需要授权的数据源名称，管理员级别下只支持填*（代表该级别全部资源）；数据源级别和数据库级别鉴权的情况下，只支持填COSDataCatalog或者*；在数据表级别鉴权下可以填写用户自定义数据源。不填情况下默认为DataLakeCatalog。注意：如果是对用户自定义数据源进行鉴权，DLC能够管理的权限是用户接入数据源的时候提供的账户的子集。
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`

	// 需要授权的表名，填*代表当前Database下所有表。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定数据表。
	Table *string `json:"Table,omitempty" name:"Table"`

	// 授权的权限操作，对于不同级别的鉴权提供不同操作。管理员权限：ALL，不填默认为ALL；数据连接级鉴权：CREATE；数据库级别鉴权：ALL、CREATE、ALTER、DROP；数据表权限：ALL、SELECT、INSERT、ALTER、DELETE、DROP、UPDATE。注意：在数据表权限下，指定的数据源不为COSDataCatalog的时候，只支持SELECT操作。
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 授权类型，现在支持八种授权类型：ADMIN:管理员级别鉴权 DATASOURCE：数据连接级别鉴权 DATABASE：数据库级别鉴权 TABLE：表级别鉴权 VIEW：视图级别鉴权 FUNCTION：函数级别鉴权 COLUMN：列级别鉴权 ENGINE：数据引擎鉴权。不填默认为管理员级别鉴权。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 需要授权的函数名，填*代表当前Catalog下所有函数。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定函数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function *string `json:"Function,omitempty" name:"Function"`

	// 需要授权的视图，填*代表当前Database下所有视图。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定视图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	View *string `json:"View,omitempty" name:"View"`

	// 需要授权的列，填*代表当前所有列。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Column *string `json:"Column,omitempty" name:"Column"`

	// 需要授权的数据引擎，填*代表当前所有引擎。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// 用户是否可以进行二次授权。当为true的时候，被授权的用户可以将本次获取的权限再次授权给其他子用户。默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReAuth *bool `json:"ReAuth,omitempty" name:"ReAuth"`

	// 权限来源，入参不填。USER：权限来自用户本身；WORKGROUP：权限来自绑定的工作组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 授权模式，入参不填。COMMON：普通模式；SENIOR：高级模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 操作者，入参不填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 权限创建的时间，入参不填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 权限所属工作组的ID，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceId *int64 `json:"SourceId,omitempty" name:"SourceId"`

	// 权限所属工作组的名称，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`

	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type Property struct {
	// 属性key名称。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 属性key对应的value。
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type ReportHeartbeatMetaDataRequestParams struct {
	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 锁ID
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitempty" name:"TxnId"`
}

type ReportHeartbeatMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 锁ID
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitempty" name:"TxnId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SQLTask struct {
	// base64加密后的SQL语句
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 任务的配置信息
	Config []*KVPair `json:"Config,omitempty" name:"Config"`
}

type Script struct {
	// 脚本Id，长度36字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptId *string `json:"ScriptId,omitempty" name:"ScriptId"`

	// 脚本名称，长度0-25。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// 脚本描述，长度0-50。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDesc *string `json:"ScriptDesc,omitempty" name:"ScriptDesc"`

	// 默认关联数据库。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL描述，长度0-10000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLStatement *string `json:"SQLStatement,omitempty" name:"SQLStatement"`

	// 更新时间戳， 单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SparkJobInfo struct {
	// spark作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// spark作业类型，可去1或者2，1表示batch作业， 2表示streaming作业
	JobType *int64 `json:"JobType,omitempty" name:"JobType"`

	// 引擎名
	DataEngine *string `json:"DataEngine,omitempty" name:"DataEngine"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitempty" name:"Eni"`

	// 程序包是否本地上传，cos或者lakefs
	IsLocal *string `json:"IsLocal,omitempty" name:"IsLocal"`

	// 程序包路径
	JobFile *string `json:"JobFile,omitempty" name:"JobFile"`

	// 角色ID
	RoleArn *int64 `json:"RoleArn,omitempty" name:"RoleArn"`

	// spark作业运行主类
	MainClass *string `json:"MainClass,omitempty" name:"MainClass"`

	// 命令行参数，spark作业命令行参数，空格分隔
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`

	// spark原生配置，换行符分隔
	JobConf *string `json:"JobConf,omitempty" name:"JobConf"`

	// 依赖jars是否本地上传，cos或者lakefs
	IsLocalJars *string `json:"IsLocalJars,omitempty" name:"IsLocalJars"`

	// spark作业依赖jars，逗号分隔
	JobJars *string `json:"JobJars,omitempty" name:"JobJars"`

	// 依赖文件是否本地上传，cos或者lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitempty" name:"IsLocalFiles"`

	// spark作业依赖文件，逗号分隔
	JobFiles *string `json:"JobFiles,omitempty" name:"JobFiles"`

	// spark作业driver资源大小
	JobDriverSize *string `json:"JobDriverSize,omitempty" name:"JobDriverSize"`

	// spark作业executor资源大小
	JobExecutorSize *string `json:"JobExecutorSize,omitempty" name:"JobExecutorSize"`

	// spark作业executor个数
	JobExecutorNums *int64 `json:"JobExecutorNums,omitempty" name:"JobExecutorNums"`

	// spark流任务最大重试次数
	JobMaxAttempts *int64 `json:"JobMaxAttempts,omitempty" name:"JobMaxAttempts"`

	// spark作业创建者
	JobCreator *string `json:"JobCreator,omitempty" name:"JobCreator"`

	// spark作业创建时间
	JobCreateTime *int64 `json:"JobCreateTime,omitempty" name:"JobCreateTime"`

	// spark作业更新时间
	JobUpdateTime *uint64 `json:"JobUpdateTime,omitempty" name:"JobUpdateTime"`

	// spark作业最近任务ID
	CurrentTaskId *string `json:"CurrentTaskId,omitempty" name:"CurrentTaskId"`

	// spark作业最近运行状态
	JobStatus *int64 `json:"JobStatus,omitempty" name:"JobStatus"`

	// spark流作业统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamingStat *StreamingStatistics `json:"StreamingStat,omitempty" name:"StreamingStat"`

	// 数据源名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitempty" name:"IsLocalPythonFiles"`

	// 注：该返回值已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppPythonFiles *string `json:"AppPythonFiles,omitempty" name:"AppPythonFiles"`

	// archives：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLocalArchives *string `json:"IsLocalArchives,omitempty" name:"IsLocalArchives"`

	// archives：依赖资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobArchives *string `json:"JobArchives,omitempty" name:"JobArchives"`

	// Spark Image 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImage *string `json:"SparkImage,omitempty" name:"SparkImage"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobPythonFiles *string `json:"JobPythonFiles,omitempty" name:"JobPythonFiles"`

	// 当前job正在运行或准备运行的任务个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNum *int64 `json:"TaskNum,omitempty" name:"TaskNum"`

	// 引擎状态：-100（默认：未知状态），-2~11：引擎正常状态；
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineStatus *int64 `json:"DataEngineStatus,omitempty" name:"DataEngineStatus"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于JobExecutorNums
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobExecutorMaxNumbers *int64 `json:"JobExecutorMaxNumbers,omitempty" name:"JobExecutorMaxNumbers"`
}

type StatementOutput struct {
	// 执行总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionCount *int64 `json:"ExecutionCount,omitempty" name:"ExecutionCount"`

	// Statement数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*KVPair `json:"Data,omitempty" name:"Data"`

	// Statement状态:ok,error
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorName *string `json:"ErrorName,omitempty" name:"ErrorName"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorValue *string `json:"ErrorValue,omitempty" name:"ErrorValue"`

	// 错误堆栈信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage []*string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// SQL类型任务结果返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLResult *string `json:"SQLResult,omitempty" name:"SQLResult"`
}

type StreamingStatistics struct {
	// 任务开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 数据接收器数
	Receivers *int64 `json:"Receivers,omitempty" name:"Receivers"`

	// 运行中的接收器数
	NumActiveReceivers *int64 `json:"NumActiveReceivers,omitempty" name:"NumActiveReceivers"`

	// 不活跃的接收器数
	NumInactiveReceivers *int64 `json:"NumInactiveReceivers,omitempty" name:"NumInactiveReceivers"`

	// 运行中的批数
	NumActiveBatches *int64 `json:"NumActiveBatches,omitempty" name:"NumActiveBatches"`

	// 待处理的批数
	NumRetainedCompletedBatches *int64 `json:"NumRetainedCompletedBatches,omitempty" name:"NumRetainedCompletedBatches"`

	// 已完成的批数
	NumTotalCompletedBatches *int64 `json:"NumTotalCompletedBatches,omitempty" name:"NumTotalCompletedBatches"`

	// 平均输入速率
	AverageInputRate *float64 `json:"AverageInputRate,omitempty" name:"AverageInputRate"`

	// 平均等待时长
	AverageSchedulingDelay *float64 `json:"AverageSchedulingDelay,omitempty" name:"AverageSchedulingDelay"`

	// 平均处理时长
	AverageProcessingTime *float64 `json:"AverageProcessingTime,omitempty" name:"AverageProcessingTime"`

	// 平均延时
	AverageTotalDelay *float64 `json:"AverageTotalDelay,omitempty" name:"AverageTotalDelay"`
}

// Predefined struct for user
type SuspendResumeDataEngineRequestParams struct {
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 操作类型 suspend/resume
	Operate *string `json:"Operate,omitempty" name:"Operate"`
}

type SuspendResumeDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 操作类型 suspend/resume
	Operate *string `json:"Operate,omitempty" name:"Operate"`
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
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type SwitchDataEngineRequestParams struct {
	// 主集群名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 是否开启备集群
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitempty" name:"StartStandbyCluster"`
}

type SwitchDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 主集群名称
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 是否开启备集群
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitempty" name:"StartStandbyCluster"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 字段描述
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 字段默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 字段是否是非空
	NotNull *bool `json:"NotNull,omitempty" name:"NotNull"`
}

type TPartition struct {
	// 字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 字段描述
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 分区类型
	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`

	// 分区格式
	PartitionFormat *string `json:"PartitionFormat,omitempty" name:"PartitionFormat"`

	// 分区分隔数
	PartitionDot *int64 `json:"PartitionDot,omitempty" name:"PartitionDot"`

	// 分区转换策略
	Transform *string `json:"Transform,omitempty" name:"Transform"`

	// 策略参数
	TransformArgs []*string `json:"TransformArgs,omitempty" name:"TransformArgs"`
}

type TableBaseInfo struct {
	// 该数据表所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据表名字
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 该数据表所属数据源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 该数据表备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableComment *string `json:"TableComment,omitempty" name:"TableComment"`

	// 具体类型，表or视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据格式类型，hive，iceberg等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableFormat *string `json:"TableFormat,omitempty" name:"TableFormat"`

	// 建表用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// 建表用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSubUin *string `json:"UserSubUin,omitempty" name:"UserSubUin"`

	// 数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernPolicy *DataGovernPolicy `json:"GovernPolicy,omitempty" name:"GovernPolicy"`
}

type TableInfo struct {
	// 数据表配置信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 数据表格式。每次入参可选如下其一的KV结构，[TextFile，CSV，Json, Parquet, ORC, AVRD]。
	DataFormat *DataFormat `json:"DataFormat,omitempty" name:"DataFormat"`

	// 数据表列信息。
	Columns []*Column `json:"Columns,omitempty" name:"Columns"`

	// 数据表分块信息。
	Partitions []*Partition `json:"Partitions,omitempty" name:"Partitions"`

	// 数据存储路径。当前仅支持cos路径，格式如下：cosn://bucket-name/filepath。
	Location *string `json:"Location,omitempty" name:"Location"`
}

type TableResponseInfo struct {
	// 数据表基本信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 数据表列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitempty" name:"Columns"`

	// 数据表分块信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*Partition `json:"Partitions,omitempty" name:"Partitions"`

	// 数据存储路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 数据表属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 数据表更新时间, 单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 数据表创建时间,单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputFormat *string `json:"InputFormat,omitempty" name:"InputFormat"`

	// 数据表存储大小（单位：Byte）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

	// 数据表行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`
}

type TagInfo struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {
	// SQL查询任务
	SQLTask *SQLTask `json:"SQLTask,omitempty" name:"SQLTask"`

	// Spark SQL查询任务
	SparkSQLTask *SQLTask `json:"SparkSQLTask,omitempty" name:"SparkSQLTask"`
}

type TaskResponseInfo struct {
	// 任务所属Database的名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 任务数据量。
	DataAmount *int64 `json:"DataAmount,omitempty" name:"DataAmount"`

	// 任务Id。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 计算耗时，单位： ms
	UsedTime *int64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// 任务输出路径。
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务状态：0 初始化， 1 执行中， 2 执行成功，-1 执行失败，-3 已取消。
	State *int64 `json:"State,omitempty" name:"State"`

	// 任务SQL类型，DDL|DML等
	SQLType *string `json:"SQLType,omitempty" name:"SQLType"`

	// 任务SQL语句
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 结果是否过期。
	ResultExpired *bool `json:"ResultExpired,omitempty" name:"ResultExpired"`

	// 数据影响统计信息。
	RowAffectInfo *string `json:"RowAffectInfo,omitempty" name:"RowAffectInfo"`

	// 任务结果数据表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSet *string `json:"DataSet,omitempty" name:"DataSet"`

	// 失败信息, 例如：errorMessage。该字段已废弃。
	Error *string `json:"Error,omitempty" name:"Error"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`

	// 任务执行输出信息。
	OutputMessage *string `json:"OutputMessage,omitempty" name:"OutputMessage"`

	// 执行SQL的引擎类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务进度明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgressDetail *string `json:"ProgressDetail,omitempty" name:"ProgressDetail"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 计算资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineId *string `json:"DataEngineId,omitempty" name:"DataEngineId"`

	// 执行sql的子uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 计算资源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineName *string `json:"DataEngineName,omitempty" name:"DataEngineName"`

	// 导入类型是本地导入还是cos
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 导入配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputConf *string `json:"InputConf,omitempty" name:"InputConf"`

	// 数据条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataNumber *int64 `json:"DataNumber,omitempty" name:"DataNumber"`

	// 查询数据能不能下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanDownload *bool `json:"CanDownload,omitempty" name:"CanDownload"`

	// 用户别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// spark应用作业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobName *string `json:"SparkJobName,omitempty" name:"SparkJobName"`

	// spark应用作业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobId *string `json:"SparkJobId,omitempty" name:"SparkJobId"`

	// spark应用入口jar文件
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkJobFile *string `json:"SparkJobFile,omitempty" name:"SparkJobFile"`

	// spark ui url
	// 注意：此字段可能返回 null，表示取不到有效值。
	UiUrl *string `json:"UiUrl,omitempty" name:"UiUrl"`

	// 任务耗时，单位： ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`

	// spark app job执行task的程序入口参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CmdArgs *string `json:"CmdArgs,omitempty" name:"CmdArgs"`
}

type TaskResultInfo struct {
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据源名称，当前任务执行时候选中的默认数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`

	// 数据库名称，当前任务执行时候选中的默认数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 当前执行的SQL，一个任务包含一个SQL
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 执行任务的类型，现在分为DDL、DML、DQL
	SQLType *string `json:"SQLType,omitempty" name:"SQLType"`

	// 任务当前的状态，0：初始化 1：任务运行中 2：任务执行成功 -1：任务执行失败 -3：用户手动终止。只有任务执行成功的情况下，才会返回任务执行的结果
	State *int64 `json:"State,omitempty" name:"State"`

	// 扫描的数据量，单位byte
	DataAmount *int64 `json:"DataAmount,omitempty" name:"DataAmount"`

	// 计算耗时，单位： ms
	UsedTime *int64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// 任务结果输出的COS桶地址
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// 任务创建时间，时间戳
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务执行信息，成功时返回success，失败时返回失败原因
	OutputMessage *string `json:"OutputMessage,omitempty" name:"OutputMessage"`

	// 被影响的行数
	RowAffectInfo *string `json:"RowAffectInfo,omitempty" name:"RowAffectInfo"`

	// 结果的schema信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSchema []*Column `json:"ResultSchema,omitempty" name:"ResultSchema"`

	// 结果信息，反转义后，外层数组的每个元素为一行数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSet *string `json:"ResultSet,omitempty" name:"ResultSet"`

	// 分页信息，如果没有更多结果数据，nextToken为空
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`

	// 任务进度明细
	ProgressDetail *string `json:"ProgressDetail,omitempty" name:"ProgressDetail"`

	// 控制台展示格式。table：表格展示 text：文本展示
	DisplayFormat *string `json:"DisplayFormat,omitempty" name:"DisplayFormat"`

	// 任务耗时，单位： ms
	TotalTime *int64 `json:"TotalTime,omitempty" name:"TotalTime"`
}

type TasksInfo struct {
	// 任务类型，SQLTask：SQL查询任务。SparkSQLTask：Spark SQL查询任务
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 容错策略。Proceed：前面任务出错/取消后继续执行后面的任务。Terminate：前面的任务出错/取消之后终止后面任务的执行，后面的任务全部标记为已取消。
	FailureTolerance *string `json:"FailureTolerance,omitempty" name:"FailureTolerance"`

	// base64加密后的SQL语句，用";"号分隔每个SQL语句，一次最多提交50个任务。严格按照前后顺序执行
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 任务的配置信息，当前仅支持SparkSQLTask任务。
	Config []*KVPair `json:"Config,omitempty" name:"Config"`

	// 任务的用户自定义参数信息
	Params []*KVPair `json:"Params,omitempty" name:"Params"`
}

type TasksOverview struct {
	// 正在排队的任务个数
	TaskQueuedCount *int64 `json:"TaskQueuedCount,omitempty" name:"TaskQueuedCount"`

	// 初始化的任务个数
	TaskInitCount *int64 `json:"TaskInitCount,omitempty" name:"TaskInitCount"`

	// 正在执行的任务个数
	TaskRunningCount *int64 `json:"TaskRunningCount,omitempty" name:"TaskRunningCount"`

	// 当前时间范围的总任务个数
	TotalTaskCount *int64 `json:"TotalTaskCount,omitempty" name:"TotalTaskCount"`
}

type TextFile struct {
	// 文本类型，本参数取值为TextFile。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 处理文本用的正则表达式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

// Predefined struct for user
type UnbindWorkGroupsFromUserRequestParams struct {
	// 解绑的工作组Id和用户Id的关联关系
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitempty" name:"AddInfo"`
}

type UnbindWorkGroupsFromUserRequest struct {
	*tchttp.BaseRequest
	
	// 解绑的工作组Id和用户Id的关联关系
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitempty" name:"AddInfo"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
}

type UnlockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 锁ID
	LockId *int64 `json:"LockId,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitempty" name:"DatasourceConnectionName"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type UpdateRowFilterRequestParams struct {
	// 行过滤策略的id，此值可以通过DescribeUserInfo或者DescribeWorkGroupInfo接口获取
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 新的过滤策略。
	Policy *Policy `json:"Policy,omitempty" name:"Policy"`
}

type UpdateRowFilterRequest struct {
	*tchttp.BaseRequest
	
	// 行过滤策略的id，此值可以通过DescribeUserInfo或者DescribeWorkGroupInfo接口获取
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 新的过滤策略。
	Policy *Policy `json:"Policy,omitempty" name:"Policy"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type UserIdSetOfWorkGroupId struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 用户Id集合，和CAM侧Uin匹配
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type UserInfo struct {
	// 用户Id，和子用户uin相同
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`

	// 单独给用户绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 创建时间，格式如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 关联的工作组集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupSet []*WorkGroupMessage `json:"WorkGroupSet,omitempty" name:"WorkGroupSet"`

	// 是否是主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOwner *bool `json:"IsOwner,omitempty" name:"IsOwner"`

	// 用户类型。ADMIN：管理员 COMMON：普通用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 用户别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`
}

type UserMessage struct {
	// 用户Id，和CAM侧子用户Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitempty" name:"UserDescription"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 当前用户的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用户别名
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`
}

type ViewBaseInfo struct {
	// 该视图所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 视图名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 视图创建人昵称
	UserAlias *string `json:"UserAlias,omitempty" name:"UserAlias"`

	// 视图创建人ID
	UserSubUin *string `json:"UserSubUin,omitempty" name:"UserSubUin"`
}

type ViewResponseInfo struct {
	// 视图基本信息。
	ViewBaseInfo *ViewBaseInfo `json:"ViewBaseInfo,omitempty" name:"ViewBaseInfo"`

	// 视图列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitempty" name:"Columns"`

	// 视图属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties"`

	// 视图创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 视图更新时间。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type WorkGroupIdSetOfUserId struct {
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitempty" name:"WorkGroupIds"`
}

type WorkGroupInfo struct {
	// 查询到的工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitempty" name:"WorkGroupName"`

	// 工作组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`

	// 工作组关联的用户数量
	UserNum *int64 `json:"UserNum,omitempty" name:"UserNum"`

	// 工作组关联的用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSet []*UserMessage `json:"UserSet,omitempty" name:"UserSet"`

	// 工作组绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitempty" name:"PolicySet"`

	// 工作组的创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 工作组的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type WorkGroupMessage struct {
	// 工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitempty" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitempty" name:"WorkGroupName"`

	// 工作组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupDescription *string `json:"WorkGroupDescription,omitempty" name:"WorkGroupDescription"`

	// 创建者
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 工作组创建的时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}