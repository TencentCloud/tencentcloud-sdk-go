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

package v20220519

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddCrossVpcSubnetSupportForClientNodeRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 子网信息
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitnil,omitempty" name:"SubnetInfo"`
}

type AddCrossVpcSubnetSupportForClientNodeRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 子网信息
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitnil,omitempty" name:"SubnetInfo"`
}

func (r *AddCrossVpcSubnetSupportForClientNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCrossVpcSubnetSupportForClientNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "SubnetInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCrossVpcSubnetSupportForClientNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCrossVpcSubnetSupportForClientNodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCrossVpcSubnetSupportForClientNodeResponse struct {
	*tchttp.BaseResponse
	Response *AddCrossVpcSubnetSupportForClientNodeResponseParams `json:"Response"`
}

func (r *AddCrossVpcSubnetSupportForClientNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCrossVpcSubnetSupportForClientNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachFileSystemBucketRequestParams struct {
	// 无
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 关联新Bucket
	Bucket *MappedBucket `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

type AttachFileSystemBucketRequest struct {
	*tchttp.BaseRequest
	
	// 无
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 关联新Bucket
	Bucket *MappedBucket `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

func (r *AttachFileSystemBucketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachFileSystemBucketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "Bucket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachFileSystemBucketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachFileSystemBucketResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachFileSystemBucketResponse struct {
	*tchttp.BaseResponse
	Response *AttachFileSystemBucketResponseParams `json:"Response"`
}

func (r *AttachFileSystemBucketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachFileSystemBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAddClientNodesRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 添加客户端节点列表
	ClientNodes []*LinuxNodeAttribute `json:"ClientNodes,omitnil,omitempty" name:"ClientNodes"`

	// 是否单集群默认是false	
	SingleClusterFlag *bool `json:"SingleClusterFlag,omitnil,omitempty" name:"SingleClusterFlag"`
}

type BatchAddClientNodesRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 添加客户端节点列表
	ClientNodes []*LinuxNodeAttribute `json:"ClientNodes,omitnil,omitempty" name:"ClientNodes"`

	// 是否单集群默认是false	
	SingleClusterFlag *bool `json:"SingleClusterFlag,omitnil,omitempty" name:"SingleClusterFlag"`
}

func (r *BatchAddClientNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddClientNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "ClientNodes")
	delete(f, "SingleClusterFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchAddClientNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAddClientNodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchAddClientNodesResponse struct {
	*tchttp.BaseResponse
	Response *BatchAddClientNodesResponseParams `json:"Response"`
}

func (r *BatchAddClientNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddClientNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteClientNodesRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 删除的客户端节点列表
	ClientNodes []*LinuxNodeAttribute `json:"ClientNodes,omitnil,omitempty" name:"ClientNodes"`

	// 是否单集群，默认是false
	SingleClusterFlag *bool `json:"SingleClusterFlag,omitnil,omitempty" name:"SingleClusterFlag"`
}

type BatchDeleteClientNodesRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 删除的客户端节点列表
	ClientNodes []*LinuxNodeAttribute `json:"ClientNodes,omitnil,omitempty" name:"ClientNodes"`

	// 是否单集群，默认是false
	SingleClusterFlag *bool `json:"SingleClusterFlag,omitnil,omitempty" name:"SingleClusterFlag"`
}

func (r *BatchDeleteClientNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteClientNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "ClientNodes")
	delete(f, "SingleClusterFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteClientNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteClientNodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeleteClientNodesResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteClientNodesResponseParams `json:"Response"`
}

func (r *BatchDeleteClientNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteClientNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientClusterManagerNodeInfo struct {
	// 客户端节点IP
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// 节点Instance Id
	NodeInstanceId *string `json:"NodeInstanceId,omitnil,omitempty" name:"NodeInstanceId"`

	// 初始密码
	InitialPassword *string `json:"InitialPassword,omitnil,omitempty" name:"InitialPassword"`
}

type ClientNodeAttribute struct {
	// 客户端节点IP
	ClientNodeIp *string `json:"ClientNodeIp,omitnil,omitempty" name:"ClientNodeIp"`

	// 客户端节点服务状态, Active(运行中), Adding(添加中), Destroying(销毁中), Down(已停止)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 客户端节点类型，extend(扩展节点)，manager(管理节点)
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`
}

type ClientToken struct {
	// 节点 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// 挂载点
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDirectory *string `json:"LocalDirectory,omitnil,omitempty" name:"LocalDirectory"`

	// 可以访问的 GooseFS 目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	GooseFSDirectory *string `json:"GooseFSDirectory,omitnil,omitempty" name:"GooseFSDirectory"`

	// token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type ClusterRole struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 目录列表
	DirectoryList []*string `json:"DirectoryList,omitnil,omitempty" name:"DirectoryList"`
}

// Predefined struct for user
type CreateDataRepositoryTaskRequestParams struct {
	// 数据流通任务类型, FS_TO_COS(文件系统到COS Bucket),或者COS_TO_FS(COS Bucket到文件系统)
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// COS存储桶名
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 对于FS_TO_COS, TaskPath是Bucket映射目录的相对路径, 对于COS_TO_FS是COS上的路径。如果置为空, 则表示全部数据
	TaskPath *string `json:"TaskPath,omitnil,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 数据流通方式 MSP_AFM 手动加载  RAW_AFM 按需加载
	RepositoryType *string `json:"RepositoryType,omitnil,omitempty" name:"RepositoryType"`

	// 文件列表下载地址，以http开头
	TextLocation *string `json:"TextLocation,omitnil,omitempty" name:"TextLocation"`
}

type CreateDataRepositoryTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据流通任务类型, FS_TO_COS(文件系统到COS Bucket),或者COS_TO_FS(COS Bucket到文件系统)
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// COS存储桶名
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 对于FS_TO_COS, TaskPath是Bucket映射目录的相对路径, 对于COS_TO_FS是COS上的路径。如果置为空, 则表示全部数据
	TaskPath *string `json:"TaskPath,omitnil,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 数据流通方式 MSP_AFM 手动加载  RAW_AFM 按需加载
	RepositoryType *string `json:"RepositoryType,omitnil,omitempty" name:"RepositoryType"`

	// 文件列表下载地址，以http开头
	TextLocation *string `json:"TextLocation,omitnil,omitempty" name:"TextLocation"`
}

func (r *CreateDataRepositoryTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataRepositoryTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "Bucket")
	delete(f, "FileSystemId")
	delete(f, "TaskPath")
	delete(f, "TaskName")
	delete(f, "RepositoryType")
	delete(f, "TextLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataRepositoryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataRepositoryTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataRepositoryTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataRepositoryTaskResponseParams `json:"Response"`
}

func (r *CreateDataRepositoryTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataRepositoryTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSystemRequestParams struct {
	// 文件系统类型, 可填goosefs和goosefsx
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件系统备注描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// vpc网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网所在的可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 文件系统关联的tag
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// GooseFSx构建时要传递的参数
	GooseFSxBuildElements *GooseFSxBuildElement `json:"GooseFSxBuildElements,omitnil,omitempty" name:"GooseFSxBuildElements"`
}

type CreateFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统类型, 可填goosefs和goosefsx
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件系统备注描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// vpc网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网所在的可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 文件系统关联的tag
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// GooseFSx构建时要传递的参数
	GooseFSxBuildElements *GooseFSxBuildElement `json:"GooseFSxBuildElements,omitnil,omitempty" name:"GooseFSxBuildElements"`
}

func (r *CreateFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Zone")
	delete(f, "Tag")
	delete(f, "GooseFSxBuildElements")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSystemResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileSystemResponseParams `json:"Response"`
}

func (r *CreateFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCrossVpcSubnetSupportForClientNodeRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 子网信息
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitnil,omitempty" name:"SubnetInfo"`
}

type DeleteCrossVpcSubnetSupportForClientNodeRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 子网信息
	SubnetInfo *SubnetInfo `json:"SubnetInfo,omitnil,omitempty" name:"SubnetInfo"`
}

func (r *DeleteCrossVpcSubnetSupportForClientNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCrossVpcSubnetSupportForClientNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "SubnetInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCrossVpcSubnetSupportForClientNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCrossVpcSubnetSupportForClientNodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCrossVpcSubnetSupportForClientNodeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCrossVpcSubnetSupportForClientNodeResponseParams `json:"Response"`
}

func (r *DeleteCrossVpcSubnetSupportForClientNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCrossVpcSubnetSupportForClientNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileSystemRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DeleteFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DeleteFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileSystemResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileSystemResponseParams `json:"Response"`
}

func (r *DeleteFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientNodesRequestParams struct {
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeClientNodesRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeClientNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientNodesResponseParams struct {
	// 客户端节点数组
	ClientNodes []*ClientNodeAttribute `json:"ClientNodes,omitnil,omitempty" name:"ClientNodes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientNodesResponseParams `json:"Response"`
}

func (r *DescribeClientNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterClientTokenRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterClientTokenRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterClientTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterClientTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterClientTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterClientTokenResponseParams struct {
	// 客户端凭证
	ClientTokens []*ClientToken `json:"ClientTokens,omitnil,omitempty" name:"ClientTokens"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterClientTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterClientTokenResponseParams `json:"Response"`
}

func (r *DescribeClusterClientTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterClientTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRoleTokenRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type DescribeClusterRoleTokenRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *DescribeClusterRoleTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRoleTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRoleTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRoleTokenResponseParams struct {
	// 角色凭证
	RoleTokens []*RoleToken `json:"RoleTokens,omitnil,omitempty" name:"RoleTokens"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterRoleTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterRoleTokenResponseParams `json:"Response"`
}

func (r *DescribeClusterRoleTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRoleTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRolesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type DescribeClusterRolesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *DescribeClusterRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRolesResponseParams struct {
	// 集群角色
	ClusterRoles []*ClusterRole `json:"ClusterRoles,omitnil,omitempty" name:"ClusterRoles"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterRolesResponseParams `json:"Response"`
}

func (r *DescribeClusterRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataRepositoryTaskStatusRequestParams struct {
	// task id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// file system id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeDataRepositoryTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// task id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// file system id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeDataRepositoryTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRepositoryTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataRepositoryTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataRepositoryTaskStatusResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态 0(初始化中), 1(运行中), 2(已完成), 3(任务失败)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 已完成的文件数量
	FinishedFileNumber *uint64 `json:"FinishedFileNumber,omitnil,omitempty" name:"FinishedFileNumber"`

	// 已完成的数据量
	FinishedCapacity *uint64 `json:"FinishedCapacity,omitnil,omitempty" name:"FinishedCapacity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataRepositoryTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataRepositoryTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeDataRepositoryTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRepositoryTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemBucketsRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeFileSystemBucketsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeFileSystemBucketsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemBucketsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileSystemBucketsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemBucketsResponseParams struct {
	// bucket列表
	BucketList []*MappedBucket `json:"BucketList,omitnil,omitempty" name:"BucketList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileSystemBucketsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileSystemBucketsResponseParams `json:"Response"`
}

func (r *DescribeFileSystemBucketsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemBucketsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemsRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页的数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeFileSystemsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页的数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeFileSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemsResponseParams struct {
	// 文件系统列表
	FSAttributeList []*FSAttribute `json:"FSAttributeList,omitnil,omitempty" name:"FSAttributeList"`

	// 总共的文件系统数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileSystemsResponseParams `json:"Response"`
}

func (r *DescribeFileSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachFileSystemBucketRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 要解绑的Bucket名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`
}

type DetachFileSystemBucketRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 要解绑的Bucket名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`
}

func (r *DetachFileSystemBucketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachFileSystemBucketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "BucketName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachFileSystemBucketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachFileSystemBucketResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachFileSystemBucketResponse struct {
	*tchttp.BaseResponse
	Response *DetachFileSystemBucketResponseParams `json:"Response"`
}

func (r *DetachFileSystemBucketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachFileSystemBucketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandCapacityRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 新增扩容的系统容量
	ExpandedCapacity *uint64 `json:"ExpandedCapacity,omitnil,omitempty" name:"ExpandedCapacity"`
}

type ExpandCapacityRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 新增扩容的系统容量
	ExpandedCapacity *uint64 `json:"ExpandedCapacity,omitnil,omitempty" name:"ExpandedCapacity"`
}

func (r *ExpandCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "ExpandedCapacity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandCapacityResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExpandCapacityResponse struct {
	*tchttp.BaseResponse
	Response *ExpandCapacityResponseParams `json:"Response"`
}

func (r *ExpandCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FSAttribute struct {
	// 文件系统类型, 可填goosefs和goosefsx
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// GooseFSx文件系统属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	GooseFSxAttribute *GooseFSxAttribute `json:"GooseFSxAttribute,omitnil,omitempty" name:"GooseFSxAttribute"`

	// 文件系统状态 ACTIVE(运行中), CREATING(创建中), DESTROYING(销毁中), FAIL(创建失败),EXPANDING(扩容中),PROBING(容灾中)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件系统备注描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// vpc ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网所在的可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Tag数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 更新属性时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type GooseFSxAttribute struct {
	// GooseFSx的型号
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 容量单位是GB, 例如4608(4.5TB)
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 要关联映射的bucket列表
	MappedBucketList []*MappedBucket `json:"MappedBucketList,omitnil,omitempty" name:"MappedBucketList"`

	// 客户侧管理节点信息
	ClientManagerNodeList []*ClientClusterManagerNodeInfo `json:"ClientManagerNodeList,omitnil,omitempty" name:"ClientManagerNodeList"`
}

type GooseFSxBuildElement struct {
	// GooseFSx的型号
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 容量单位是GB, 例如4608(4.5TB)
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 要关联映射的bucket列表
	MappedBucketList []*MappedBucket `json:"MappedBucketList,omitnil,omitempty" name:"MappedBucketList"`
}

type LinuxNodeAttribute struct {
	// cvmId
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点所属vpcid
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 节点所属子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// linux客户端节点地址
	LinuxClientNodeIp *string `json:"LinuxClientNodeIp,omitnil,omitempty" name:"LinuxClientNodeIp"`
}

type MappedBucket struct {
	// 对象存储Bucket名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 映射到的文件系统路径, 默认为/
	FileSystemPath *string `json:"FileSystemPath,omitnil,omitempty" name:"FileSystemPath"`

	// 数据流动的自动策略, 包含加载与沉降。策略可以是多种的组合
	// 按需加载(OnDemandImport)
	// 自动加载元数据(AutoImportMeta)
	// 自动加载数据(AutoImportData)
	// 周期加载(PeriodImport)
	// 
	// 周期沉降(PeriodExport)
	// 立即沉降(ImmediateExport)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRepositoryTaskAutoStrategy []*string `json:"DataRepositoryTaskAutoStrategy,omitnil,omitempty" name:"DataRepositoryTaskAutoStrategy"`

	// 绑定bucket的数据流动策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则备注与描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleDescription *string `json:"RuleDescription,omitnil,omitempty" name:"RuleDescription"`

	// 桶关联状态 0：关联中 1：关联完成
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否使用全球加速域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateFlag *bool `json:"AccelerateFlag,omitnil,omitempty" name:"AccelerateFlag"`

	// 桶所在的园区
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`
}

// Predefined struct for user
type ModifyDataRepositoryBandwidthRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 带宽, 单位MB/S
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`
}

type ModifyDataRepositoryBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 带宽, 单位MB/S
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`
}

func (r *ModifyDataRepositoryBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataRepositoryBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "Bandwidth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataRepositoryBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataRepositoryBandwidthResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDataRepositoryBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataRepositoryBandwidthResponseParams `json:"Response"`
}

func (r *ModifyDataRepositoryBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataRepositoryBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCrossVpcSubnetSupportForClientNodeRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type QueryCrossVpcSubnetSupportForClientNodeRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *QueryCrossVpcSubnetSupportForClientNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCrossVpcSubnetSupportForClientNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCrossVpcSubnetSupportForClientNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCrossVpcSubnetSupportForClientNodeResponseParams struct {
	// 支持的子网信息集合
	SubnetInfoCollection []*SubnetInfo `json:"SubnetInfoCollection,omitnil,omitempty" name:"SubnetInfoCollection"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCrossVpcSubnetSupportForClientNodeResponse struct {
	*tchttp.BaseResponse
	Response *QueryCrossVpcSubnetSupportForClientNodeResponseParams `json:"Response"`
}

func (r *QueryCrossVpcSubnetSupportForClientNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCrossVpcSubnetSupportForClientNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDataRepositoryBandwidthRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type QueryDataRepositoryBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *QueryDataRepositoryBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDataRepositoryBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDataRepositoryBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDataRepositoryBandwidthResponseParams struct {
	// 数据流动带宽, 单位MB/s
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 带宽状态。1:待扩容;2:运行中;3:扩容中
	BandwidthStatus *uint64 `json:"BandwidthStatus,omitnil,omitempty" name:"BandwidthStatus"`

	// 能设置的最小带宽, 单位MB/s
	MinBandwidth *uint64 `json:"MinBandwidth,omitnil,omitempty" name:"MinBandwidth"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryDataRepositoryBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *QueryDataRepositoryBandwidthResponseParams `json:"Response"`
}

func (r *QueryDataRepositoryBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDataRepositoryBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleToken struct {
	// 角色名
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 用于goosefs client/sdk等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type SubnetInfo struct {
	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}