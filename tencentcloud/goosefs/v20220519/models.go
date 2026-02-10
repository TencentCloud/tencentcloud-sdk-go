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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 客户端集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type BatchAddClientNodesRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 添加客户端节点列表
	ClientNodes []*LinuxNodeAttribute `json:"ClientNodes,omitnil,omitempty" name:"ClientNodes"`

	// 是否单集群默认是false	
	SingleClusterFlag *bool `json:"SingleClusterFlag,omitnil,omitempty" name:"SingleClusterFlag"`

	// 客户端集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchAddClientNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAddClientNodesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 客户端集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type BatchDeleteClientNodesRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 删除的客户端节点列表
	ClientNodes []*LinuxNodeAttribute `json:"ClientNodes,omitnil,omitempty" name:"ClientNodes"`

	// 是否单集群，默认是false
	SingleClusterFlag *bool `json:"SingleClusterFlag,omitnil,omitempty" name:"SingleClusterFlag"`

	// 客户端集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteClientNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteClientNodesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

// Predefined struct for user
type BuildClientNodeMountCommandRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 自定义挂载目录的绝对路径, 如果未指定, 则会使用默认值, 格式/goosefsx/${fs_id}-proxy. 比如/goosefsx/x-c60-a2b3d4-proxy
	CustomMountDir *string `json:"CustomMountDir,omitnil,omitempty" name:"CustomMountDir"`

	// 客户端集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type BuildClientNodeMountCommandRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 自定义挂载目录的绝对路径, 如果未指定, 则会使用默认值, 格式/goosefsx/${fs_id}-proxy. 比如/goosefsx/x-c60-a2b3d4-proxy
	CustomMountDir *string `json:"CustomMountDir,omitnil,omitempty" name:"CustomMountDir"`

	// 客户端集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *BuildClientNodeMountCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildClientNodeMountCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "CustomMountDir")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BuildClientNodeMountCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildClientNodeMountCommandResponseParams struct {
	// 挂载命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BuildClientNodeMountCommandResponse struct {
	*tchttp.BaseResponse
	Response *BuildClientNodeMountCommandResponseParams `json:"Response"`
}

func (r *BuildClientNodeMountCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildClientNodeMountCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelLoadTaskRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CancelLoadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *CancelLoadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelLoadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelLoadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelLoadTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelLoadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelLoadTaskResponseParams `json:"Response"`
}

func (r *CancelLoadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelLoadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargeAttribute struct {
	// 到期时间
	CurDeadline *string `json:"CurDeadline,omitnil,omitempty" name:"CurDeadline"`

	// 付费方式
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动付费标识：0:默认未设置 1:自动续费 2 不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type ClientClusterManagerNodeInfo struct {
	// 客户端节点IP
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// 节点Instance Id
	NodeInstanceId *string `json:"NodeInstanceId,omitnil,omitempty" name:"NodeInstanceId"`

	// 初始密码
	InitialPassword *string `json:"InitialPassword,omitnil,omitempty" name:"InitialPassword"`

	// 所属集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type ClientNodeAttribute struct {
	// 客户端节点IP
	ClientNodeIp *string `json:"ClientNodeIp,omitnil,omitempty" name:"ClientNodeIp"`

	// 客户端节点服务状态, Active(运行中), Adding(添加中), Destroying(销毁中), Down(已停止)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 客户端节点类型，extend(扩展节点)，manager(管理节点)
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 节点所属vpcid	
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 节点所属子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// cvmId
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义挂载点
	MountPoint *string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`
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

	// 是否开启自定义路径(暂时仅供预热使用)
	EnableDataFlowSubPath *bool `json:"EnableDataFlowSubPath,omitnil,omitempty" name:"EnableDataFlowSubPath"`

	// 自定义路径(暂时仅供预热使用)
	DataFlowSubPath *string `json:"DataFlowSubPath,omitnil,omitempty" name:"DataFlowSubPath"`
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

	// 是否开启自定义路径(暂时仅供预热使用)
	EnableDataFlowSubPath *bool `json:"EnableDataFlowSubPath,omitnil,omitempty" name:"EnableDataFlowSubPath"`

	// 自定义路径(暂时仅供预热使用)
	DataFlowSubPath *string `json:"DataFlowSubPath,omitnil,omitempty" name:"DataFlowSubPath"`
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
	delete(f, "EnableDataFlowSubPath")
	delete(f, "DataFlowSubPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataRepositoryTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataRepositoryTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 文件系统类型, 可填goosefs和goosefsx
	//
	// Deprecated: Type is deprecated.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件系统关联的tag
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// GooseFSx构建时要传递的参数
	GooseFSxBuildElements *GooseFSxBuildElement `json:"GooseFSxBuildElements,omitnil,omitempty" name:"GooseFSxBuildElements"`

	// 客户端集群所属的安全组
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 集群ssh通信端口，默认是22
	ClusterPort *uint64 `json:"ClusterPort,omitnil,omitempty" name:"ClusterPort"`
}

type CreateFileSystemRequest struct {
	*tchttp.BaseRequest
	
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

	// 文件系统类型, 可填goosefs和goosefsx
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件系统关联的tag
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// GooseFSx构建时要传递的参数
	GooseFSxBuildElements *GooseFSxBuildElement `json:"GooseFSxBuildElements,omitnil,omitempty" name:"GooseFSxBuildElements"`

	// 客户端集群所属的安全组
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 集群ssh通信端口，默认是22
	ClusterPort *uint64 `json:"ClusterPort,omitnil,omitempty" name:"ClusterPort"`
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
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Zone")
	delete(f, "Type")
	delete(f, "Tag")
	delete(f, "GooseFSxBuildElements")
	delete(f, "SecurityGroupId")
	delete(f, "ClusterPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSystemResponseParams struct {
	// 创建成功返回的文件系统ID：
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
type CreateFilesetRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Fileset名称
	FsetName *string `json:"FsetName,omitnil,omitempty" name:"FsetName"`

	// Fileset目录
	FsetDir *string `json:"FsetDir,omitnil,omitempty" name:"FsetDir"`

	// Fileset容量配额（需带单位G）
	QuotaSizeLimit *string `json:"QuotaSizeLimit,omitnil,omitempty" name:"QuotaSizeLimit"`

	// Fileset文件数配额
	QuotaFilesLimit *string `json:"QuotaFilesLimit,omitnil,omitempty" name:"QuotaFilesLimit"`

	// Fileset文件删除操作审计
	AuditState *string `json:"AuditState,omitnil,omitempty" name:"AuditState"`
}

type CreateFilesetRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Fileset名称
	FsetName *string `json:"FsetName,omitnil,omitempty" name:"FsetName"`

	// Fileset目录
	FsetDir *string `json:"FsetDir,omitnil,omitempty" name:"FsetDir"`

	// Fileset容量配额（需带单位G）
	QuotaSizeLimit *string `json:"QuotaSizeLimit,omitnil,omitempty" name:"QuotaSizeLimit"`

	// Fileset文件数配额
	QuotaFilesLimit *string `json:"QuotaFilesLimit,omitnil,omitempty" name:"QuotaFilesLimit"`

	// Fileset文件删除操作审计
	AuditState *string `json:"AuditState,omitnil,omitempty" name:"AuditState"`
}

func (r *CreateFilesetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFilesetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FsetName")
	delete(f, "FsetDir")
	delete(f, "QuotaSizeLimit")
	delete(f, "QuotaFilesLimit")
	delete(f, "AuditState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFilesetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFilesetResponseParams struct {
	// Fileset id
	FsetId *string `json:"FsetId,omitnil,omitempty" name:"FsetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFilesetResponse struct {
	*tchttp.BaseResponse
	Response *CreateFilesetResponseParams `json:"Response"`
}

func (r *CreateFilesetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFilesetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadTaskRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 创建预热任务参数
	LoadTaskCreationAttrs *LoadTaskCreationAttrs `json:"LoadTaskCreationAttrs,omitnil,omitempty" name:"LoadTaskCreationAttrs"`
}

type CreateLoadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 创建预热任务参数
	LoadTaskCreationAttrs *LoadTaskCreationAttrs `json:"LoadTaskCreationAttrs,omitnil,omitempty" name:"LoadTaskCreationAttrs"`
}

func (r *CreateLoadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LoadTaskCreationAttrs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadTaskResponseParams struct {
	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadTaskResponseParams `json:"Response"`
}

func (r *CreateLoadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadTaskResponse) FromJsonString(s string) error {
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
type DeleteFilesetRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Fileset id
	FsetId *string `json:"FsetId,omitnil,omitempty" name:"FsetId"`
}

type DeleteFilesetRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Fileset id
	FsetId *string `json:"FsetId,omitnil,omitempty" name:"FsetId"`
}

func (r *DeleteFilesetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFilesetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFilesetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFilesetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFilesetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFilesetResponseParams `json:"Response"`
}

func (r *DeleteFilesetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFilesetResponse) FromJsonString(s string) error {
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
type DescribeFilesetGeneralConfigRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeFilesetGeneralConfigRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeFilesetGeneralConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFilesetGeneralConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFilesetGeneralConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFilesetGeneralConfigResponseParams struct {
	// 配额对root用户生效
	EnforceQuotaOnRoot *string `json:"EnforceQuotaOnRoot,omitnil,omitempty" name:"EnforceQuotaOnRoot"`

	// 配置状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFilesetGeneralConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFilesetGeneralConfigResponseParams `json:"Response"`
}

func (r *DescribeFilesetGeneralConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFilesetGeneralConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFilesetsRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// FsetId列表
	FilesetIds []*string `json:"FilesetIds,omitnil,omitempty" name:"FilesetIds"`

	// FsetDir列表
	FilesetDirs []*string `json:"FilesetDirs,omitnil,omitempty" name:"FilesetDirs"`
}

type DescribeFilesetsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// FsetId列表
	FilesetIds []*string `json:"FilesetIds,omitnil,omitempty" name:"FilesetIds"`

	// FsetDir列表
	FilesetDirs []*string `json:"FilesetDirs,omitnil,omitempty" name:"FilesetDirs"`
}

func (r *DescribeFilesetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFilesetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FilesetIds")
	delete(f, "FilesetDirs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFilesetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFilesetsResponseParams struct {
	// Fileset列表
	FilesetList []*FilesetInfo `json:"FilesetList,omitnil,omitempty" name:"FilesetList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFilesetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFilesetsResponseParams `json:"Response"`
}

func (r *DescribeFilesetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFilesetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadTaskRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeLoadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeLoadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadTaskResponseParams struct {
	// 预热任务参数
	LoadTaskAttrs *LoadTaskAttrs `json:"LoadTaskAttrs,omitnil,omitempty" name:"LoadTaskAttrs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadTaskResponseParams `json:"Response"`
}

func (r *DescribeLoadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadTaskResponse) FromJsonString(s string) error {
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

type DistributedLoadAttrs struct {
	// 预热类型，枚举值 LoadByPath｜LoadByList
	LoadType *string `json:"LoadType,omitnil,omitempty" name:"LoadType"`

	// 是否跳过相同文件，默认为 true
	SkipIfExists *bool `json:"SkipIfExists,omitnil,omitempty" name:"SkipIfExists"`

	// 预热路径，入参单条挂载路径。入参数LoadType为LoadByPath，该参数不应为空
	LoadByPath *string `json:"LoadByPath,omitnil,omitempty" name:"LoadByPath"`

	// 通过文件列表批量预热，入参为 cos://bucket-appid/ 开头的 COS 路径，且仅支持 txt 格式文件，长度不能超过255个字符。入参数LoadType为LoadByList，该参数不应为空
	LoadByList *string `json:"LoadByList,omitnil,omitempty" name:"LoadByList"`

	// 副本数配置，枚举值，可选值 SingleReplica（单副本，默认）｜MaxReplica（最大副本）
	Replica *string `json:"Replica,omitnil,omitempty" name:"Replica"`

	// 同步执行元数据预热，并基于预热后的元数据执行 DistributedLoad。默认为 false
	MetadataSync *bool `json:"MetadataSync,omitnil,omitempty" name:"MetadataSync"`
}

// Predefined struct for user
type ExpandCapacityRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 新增扩容的系统容量
	ExpandedCapacity *uint64 `json:"ExpandedCapacity,omitnil,omitempty" name:"ExpandedCapacity"`

	// 容量修改类型：add/sub
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`
}

type ExpandCapacityRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 新增扩容的系统容量
	ExpandedCapacity *uint64 `json:"ExpandedCapacity,omitnil,omitempty" name:"ExpandedCapacity"`

	// 容量修改类型：add/sub
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`
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
	delete(f, "ModifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandCapacityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 更新属性时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 文件系统付费信息
	ChargeAttribute *ChargeAttribute `json:"ChargeAttribute,omitnil,omitempty" name:"ChargeAttribute"`
}

type FilesetInfo struct {
	// Fileset id
	FsetId *string `json:"FsetId,omitnil,omitempty" name:"FsetId"`

	// Fileset名称
	FsetName *string `json:"FsetName,omitnil,omitempty" name:"FsetName"`

	// Fileset目录
	FsetDir *string `json:"FsetDir,omitnil,omitempty" name:"FsetDir"`

	// Fileset容量配额限定值
	QuotaSizeLimit *string `json:"QuotaSizeLimit,omitnil,omitempty" name:"QuotaSizeLimit"`

	// 已使用容量配额
	QuotaSizeUsed *string `json:"QuotaSizeUsed,omitnil,omitempty" name:"QuotaSizeUsed"`

	// 容量配额使用占比
	QuotaSizeUsedPercent *string `json:"QuotaSizeUsedPercent,omitnil,omitempty" name:"QuotaSizeUsedPercent"`

	// Fileset文件数配额限定值
	QuotaFilesLimit *string `json:"QuotaFilesLimit,omitnil,omitempty" name:"QuotaFilesLimit"`

	// 已使用文件数配额
	QuotaFilesUsed *string `json:"QuotaFilesUsed,omitnil,omitempty" name:"QuotaFilesUsed"`

	// 文件数配额使用占比
	QuotaFilesUsedPercent *string `json:"QuotaFilesUsedPercent,omitnil,omitempty" name:"QuotaFilesUsedPercent"`

	// Fileset审计
	AuditState *string `json:"AuditState,omitnil,omitempty" name:"AuditState"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Fileset状态：creating 配置中 active 已生效 modify 修改中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
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
	//
	// Deprecated: MappedBucketList is deprecated.
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

	// 自定义挂载点
	MountPoint *string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`
}

// Predefined struct for user
type ListLoadTasksRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 偏移量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务创建起始时间戳，默认为3天前：当前时间戳-86400*3
	StartTimestamp *uint64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 任务变更时间戳
	EndTimestamp *uint64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 筛选任务状态，枚举Waiting,Running,Canceled,Completed。默认返回所有任务
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 筛选优先级任务，默认返回所有任务
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type ListLoadTasksRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 偏移量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务创建起始时间戳，默认为3天前：当前时间戳-86400*3
	StartTimestamp *uint64 `json:"StartTimestamp,omitnil,omitempty" name:"StartTimestamp"`

	// 任务变更时间戳
	EndTimestamp *uint64 `json:"EndTimestamp,omitnil,omitempty" name:"EndTimestamp"`

	// 筛选任务状态，枚举Waiting,Running,Canceled,Completed。默认返回所有任务
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 筛选优先级任务，默认返回所有任务
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

func (r *ListLoadTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLoadTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	delete(f, "State")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLoadTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLoadTasksResponseParams struct {
	// 预热任务参数
	LoadTaskList []*LoadTaskAttrs `json:"LoadTaskList,omitnil,omitempty" name:"LoadTaskList"`

	// 任务数总量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListLoadTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListLoadTasksResponseParams `json:"Response"`
}

func (r *ListLoadTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLoadTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoadTaskAttrs struct {
	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 预热任务类型，枚举值，MetadataLoad｜DistributedLoad
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务描述，支持中文
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 任务优先级，数值越高代表优先级越高，边界值 1-9999，默认值为 1
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 元数据预热任务参数，用于仅预热元数据时入参。入参数TaskType为MetadataLoad时，该参数不应为空。
	MetadataLoadAttrs *MetadataLoadAttrs `json:"MetadataLoadAttrs,omitnil,omitempty" name:"MetadataLoadAttrs"`

	// 数据预热任务参数。入参数TaskType为DistributedLoad时，该参数不应为空。
	DistributedLoadAttrs *DistributedLoadAttrs `json:"DistributedLoadAttrs,omitnil,omitempty" name:"DistributedLoadAttrs"`

	// 将任务执行报告写入 COS 的路径，如果不需要报告则入参空
	ReportPath *string `json:"ReportPath,omitnil,omitempty" name:"ReportPath"`

	// 枚举，Completed，Running，Waiting，Cancelled
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 任务执行信息，打印预热文件成功个数，失败个数，预热耗时信息 
	TaskMessage *string `json:"TaskMessage,omitnil,omitempty" name:"TaskMessage"`

	// 预热任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 预热任务变更时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 任务提交账号，子账号或服务角色 ID
	Requester *string `json:"Requester,omitnil,omitempty" name:"Requester"`
}

type LoadTaskCreationAttrs struct {
	// 预热任务类型，枚举值，MetadataLoad｜DistributedLoad。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务优先级，数值越高代表优先级越高，边界值 1-9999，默认值为 1
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 任务描述，支持中文
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 元数据预热任务参数，用于仅预热元数据时入参。入参数TaskType为MetadataLoad时，该参数不应为空。
	MetadataLoadAttrs *MetadataLoadAttrs `json:"MetadataLoadAttrs,omitnil,omitempty" name:"MetadataLoadAttrs"`

	// 数据预热任务参数。入参数TaskType为DistributedLoad时，该参数不应为空。
	DistributedLoadAttrs *DistributedLoadAttrs `json:"DistributedLoadAttrs,omitnil,omitempty" name:"DistributedLoadAttrs"`

	// 将任务执行报告写入 COS 的路径，如果不需要报告则入参空
	ReportPath *string `json:"ReportPath,omitnil,omitempty" name:"ReportPath"`
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
	DataRepositoryTaskAutoStrategy []*string `json:"DataRepositoryTaskAutoStrategy,omitnil,omitempty" name:"DataRepositoryTaskAutoStrategy"`

	// 绑定bucket的数据流动策略ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则备注与描述
	RuleDescription *string `json:"RuleDescription,omitnil,omitempty" name:"RuleDescription"`

	// 桶关联状态 0：关联中 1：关联完成
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否使用全球加速域名
	AccelerateFlag *bool `json:"AccelerateFlag,omitnil,omitempty" name:"AccelerateFlag"`

	// 桶所在的园区
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 自定义Endpoint
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`
}

type MetadataLoadAttrs struct {
	// 预热类型，枚举值 LoadByPath｜LoadByList
	LoadType *string `json:"LoadType,omitnil,omitempty" name:"LoadType"`

	// 是否跳过相同文件，默认为 true
	SkipIfExists *bool `json:"SkipIfExists,omitnil,omitempty" name:"SkipIfExists"`

	// 预热路径，入参单条挂载路径，长度不能超过255个字符。入参数LoadType为LoadByPath，该参数不应为空
	LoadByPath *string `json:"LoadByPath,omitnil,omitempty" name:"LoadByPath"`

	// 通过文件列表批量预热，入参为 cos://bucket-appid/ 开头的 COS 路径，且仅支持 txt 格式文件，长度不能超过255个字符。入参数LoadType为LoadByList，该参数不应为空
	LoadByList *string `json:"LoadByList,omitnil,omitempty" name:"LoadByList"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type SubnetInfo struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 应用的集群；可以是集群id,也可以是All
	UsedCluster *string `json:"UsedCluster,omitnil,omitempty" name:"UsedCluster"`

	// cidr，只有当IsDirectConnect为true时才生效
	CIDR *string `json:"CIDR,omitnil,omitempty" name:"CIDR"`

	// 是否为专线接入场景
	IsDirectConnect *bool `json:"IsDirectConnect,omitnil,omitempty" name:"IsDirectConnect"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type UpdateFilesetGeneralConfigRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 配额对root用户生效
	EnforceQuotaOnRoot *string `json:"EnforceQuotaOnRoot,omitnil,omitempty" name:"EnforceQuotaOnRoot"`
}

type UpdateFilesetGeneralConfigRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 配额对root用户生效
	EnforceQuotaOnRoot *string `json:"EnforceQuotaOnRoot,omitnil,omitempty" name:"EnforceQuotaOnRoot"`
}

func (r *UpdateFilesetGeneralConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFilesetGeneralConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "EnforceQuotaOnRoot")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFilesetGeneralConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFilesetGeneralConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateFilesetGeneralConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFilesetGeneralConfigResponseParams `json:"Response"`
}

func (r *UpdateFilesetGeneralConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFilesetGeneralConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFilesetRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Fileset id
	FsetId *string `json:"FsetId,omitnil,omitempty" name:"FsetId"`

	// 容量配额限定值
	QuotaSizeLimit *string `json:"QuotaSizeLimit,omitnil,omitempty" name:"QuotaSizeLimit"`

	// 文件数配额限定值
	QuotaFilesLimit *string `json:"QuotaFilesLimit,omitnil,omitempty" name:"QuotaFilesLimit"`

	// Fileset文件删除操作审计
	AuditState *string `json:"AuditState,omitnil,omitempty" name:"AuditState"`
}

type UpdateFilesetRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Fileset id
	FsetId *string `json:"FsetId,omitnil,omitempty" name:"FsetId"`

	// 容量配额限定值
	QuotaSizeLimit *string `json:"QuotaSizeLimit,omitnil,omitempty" name:"QuotaSizeLimit"`

	// 文件数配额限定值
	QuotaFilesLimit *string `json:"QuotaFilesLimit,omitnil,omitempty" name:"QuotaFilesLimit"`

	// Fileset文件删除操作审计
	AuditState *string `json:"AuditState,omitnil,omitempty" name:"AuditState"`
}

func (r *UpdateFilesetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFilesetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FsetId")
	delete(f, "QuotaSizeLimit")
	delete(f, "QuotaFilesLimit")
	delete(f, "AuditState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFilesetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFilesetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateFilesetResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFilesetResponseParams `json:"Response"`
}

func (r *UpdateFilesetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFilesetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateLoadTaskPriorityRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务优先级，数值越高代表优先级越高，边界值 1-9999，默认值为 1
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type UpdateLoadTaskPriorityRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务优先级，数值越高代表优先级越高，边界值 1-9999，默认值为 1
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

func (r *UpdateLoadTaskPriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLoadTaskPriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TaskId")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateLoadTaskPriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateLoadTaskPriorityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateLoadTaskPriorityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateLoadTaskPriorityResponseParams `json:"Response"`
}

func (r *UpdateLoadTaskPriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLoadTaskPriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}