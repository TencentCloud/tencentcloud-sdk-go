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

package v20190823

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ClearTablesRequest struct {
	*tchttp.BaseRequest

	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待清理表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *ClearTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 清除表结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 清除表结果列表
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClusterInfo struct {

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 集群数据描述语言类型，如：`PROTO`,`TDR`或`MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 集群关联的用户私有网络实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群关联的用户子网实例ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 集群密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密码状态
	PasswordStatus *string `json:"PasswordStatus,omitempty" name:"PasswordStatus"`

	// TcaplusDB SDK连接参数，接入ID
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// TcaplusDB SDK连接参数，接入地址
	ApiAccessIp *string `json:"ApiAccessIp,omitempty" name:"ApiAccessIp"`

	// TcaplusDB SDK连接参数，接入端口
	ApiAccessPort *int64 `json:"ApiAccessPort,omitempty" name:"ApiAccessPort"`

	// 如果PasswordStatus是unmodifiable说明有旧密码还未过期，此字段将显示旧密码过期的时间，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`
}

type CompareIdlFilesRequest struct {
	*tchttp.BaseRequest

	// 待修改表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// 选中的已上传IDL文件列表，与NewIdlFiles必选其一
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles" list`

	// 本次上传IDL文件列表，与ExistingIdlFiles必选其一
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles" list`
}

func (r *CompareIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompareIdlFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompareIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次上传校验所有的IDL文件信息列表
		IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

		// 本次校验合法的表格数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 读取IDL描述文件后,根据用户指示的所选中表格解析校验结果
		TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompareIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompareIdlFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// 待创建备份表所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建备份表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的备份任务ID列表
		TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群数据描述语言类型，如：`PROTO`，`TDR`或`MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// 集群名称，可使用中文或英文字符，最大长度32个字符
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群所绑定的私有网络实例ID，形如：vpc-f49l6u0z
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群所绑定的子网实例ID，形如：subnet-pxir56ns
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群访问密码，必须是a-zA-Z0-9的字符,且必须包含数字和大小写字母
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTableGroupRequest struct {
	*tchttp.BaseRequest

	// 表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组名称，可以采用中文、英文或数字字符，最大长度32个字符
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格组ID，可以由用户指定，但在同一个集群内不能重复，如果不指定则采用自增的模式
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`
}

func (r *CreateTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTableGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功的表格组ID
		TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTableGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTablesRequest struct {
	*tchttp.BaseRequest

	// 待创建表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 用户选定的建表格IDL文件列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// 待创建表格信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *CreateTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量创建表格结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 批量创建表格结果列表
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 待删除的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除集群生成的任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIdlFilesRequest struct {
	*tchttp.BaseRequest

	// IDL所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除的IDL文件信息列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`
}

func (r *DeleteIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIdlFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果记录数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 删除结果
		IdlFileInfos []*IdlFileInfoWithoutContent `json:"IdlFileInfos,omitempty" name:"IdlFileInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteIdlFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTableGroupRequest struct {
	*tchttp.BaseRequest

	// 表格组所属的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`
}

func (r *DeleteTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTableGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除表格组所创建的任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTableGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTablesRequest struct {
	*tchttp.BaseRequest

	// 待删除表所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *DeleteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除表结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 删除表结果详情列表
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// 指定查询的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`

	// 查询过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群实例数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群实例列表
		Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClustersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdlFileInfosRequest struct {
	*tchttp.BaseRequest

	// 文件所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 文件所属表格组ID
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// 指定文件ID列表
	IdlFileIds []*string `json:"IdlFileIds,omitempty" name:"IdlFileIds" list`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIdlFileInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdlFileInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdlFileInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 文件详情列表
		IdlFileInfos []*IdlFileInfo `json:"IdlFileInfos,omitempty" name:"IdlFileInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdlFileInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIdlFileInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用区详情结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 可用区详情结果列表
		RegionInfos []*RegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableGroupsRequest struct {
	*tchttp.BaseRequest

	// 表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// 过滤条件，本接口支持：TableGroupName，TableGroupId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTableGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表格组数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表格组信息列表
		TableGroups []*TableGroupInfo `json:"TableGroups,omitempty" name:"TableGroups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesInRecycleRequest struct {
	*tchttp.BaseRequest

	// 待查询表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询表格所属表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询结果偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果返回记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesInRecycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesInRecycleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesInRecycleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表格数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表格详情结果列表
		TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesInRecycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesInRecycleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest

	// 待查询表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询表格所属表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds" list`

	// 待查询表格信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询结果偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果返回记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表格数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表格详情结果列表
		TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 需要查询任务所属的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds" list`

	// 需要查询的任务ID列表
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`

	// 过滤条件，本接口支持：Content，TaskType, Operator, Time
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询到的任务详情列表
		TaskInfos []*TaskInfoNew `json:"TaskInfos,omitempty" name:"TaskInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUinInWhitelistRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUinInWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUinInWhitelistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUinInWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询结果：`FALSE` 否；`TRUE` 是
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUinInWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUinInWhitelistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ErrorInfo struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type Filter struct {

	// 过滤字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type IdlFileInfo struct {

	// 文件名称，不包含扩展名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 数据描述语言（IDL）类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件扩展名
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// 文件大小（Bytes）
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件ID，对于已上传的文件有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// 文件内容，对于本次新上传的文件有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type IdlFileInfoWithoutContent struct {

	// 文件名称，不包含扩展名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 数据描述语言（IDL）类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件扩展名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// 文件大小（Bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest

	// 需要修改名称的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要修改的集群名称，可使用中文或英文字符，最大长度32个字符
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterPasswordRequest struct {
	*tchttp.BaseRequest

	// 需要修改密码的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 集群旧密码预期失效时间
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// 集群新密码，密码必须是a-zA-Z0-9的字符,且必须包含数字和大小写字母
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 更新模式： `1` 更新密码；`2` 更新旧密码失效时间，默认为`1` 模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyClusterPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClusterPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableGroupNameRequest struct {
	*tchttp.BaseRequest

	// 表格组所属的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改名称的表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 新的表格组名称，可以使用中英文字符和符号
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`
}

func (r *ModifyTableGroupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableGroupNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableGroupNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableGroupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableGroupNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableMemosRequest struct {
	*tchttp.BaseRequest

	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 选定表详情列表
	TableMemos []*SelectedTableInfoNew `json:"TableMemos,omitempty" name:"TableMemos" list`
}

func (r *ModifyTableMemosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableMemosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableMemosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表备注修改结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表备注修改结果列表
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableMemosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableMemosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableQuotasRequest struct {
	*tchttp.BaseRequest

	// 带扩缩容表所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 已选中待修改的表配额列表
	TableQuotas []*SelectedTableInfoNew `json:"TableQuotas,omitempty" name:"TableQuotas" list`
}

func (r *ModifyTableQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 扩缩容结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 扩缩容结果列表
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTableQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTablesRequest struct {
	*tchttp.BaseRequest

	// 待修改表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 选中的改表IDL文件
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// 待改表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *ModifyTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改表结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 修改表结果列表
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ParsedTableInfoNew struct {

	// 表格描述语言类型：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格数据结构类型：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 主键字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyFields *string `json:"KeyFields,omitempty" name:"KeyFields"`

	// 原主键字段信息，改表校验时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldKeyFields *string `json:"OldKeyFields,omitempty" name:"OldKeyFields"`

	// 非主键字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueFields *string `json:"ValueFields,omitempty" name:"ValueFields"`

	// 原非主键字段信息，改表校验时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldValueFields *string `json:"OldValueFields,omitempty" name:"OldValueFields"`

	// 所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 主键字段总大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SumKeyFieldSize *int64 `json:"SumKeyFieldSize,omitempty" name:"SumKeyFieldSize"`

	// 非主键字段总大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SumValueFieldSize *int64 `json:"SumValueFieldSize,omitempty" name:"SumValueFieldSize"`

	// 索引键集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexKeySet *string `json:"IndexKeySet,omitempty" name:"IndexKeySet"`

	// 分表因子集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// TDR版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TdrVersion *int64 `json:"TdrVersion,omitempty" name:"TdrVersion"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// LIST类型表格元素个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// SORTLIST类型表格排序字段个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// SORTLIST类型表格排序顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`
}

type RecoverRecycleTablesRequest struct {
	*tchttp.BaseRequest

	// 表所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待恢复表信息
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`
}

func (r *RecoverRecycleTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverRecycleTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecoverRecycleTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 恢复表结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 恢复表信息列表
		TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecoverRecycleTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecoverRecycleTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域Ap-Code
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域缩写
	RegionAbbr *string `json:"RegionAbbr,omitempty" name:"RegionAbbr"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

type RollbackTablesRequest struct {
	*tchttp.BaseRequest

	// 待回档表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待回档表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// 待回档时间
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// 回档模式，支持：`KEYS`
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *RollbackTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表格回档任务结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表格回档任务结果列表
		TableResults []*TableRollbackResultNew `json:"TableResults,omitempty" name:"TableResults" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SelectedTableInfoNew struct {

	// 表所属表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表实例ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格描述语言类型：`PROTO`或`TDR`
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格数据结构类型：`GENERIC`或`LIST`
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// LIST表元素个数
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// 表格预留容量（GB）
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// 表格预留读QPS
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// 表格预留写QPS
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// 表格备注信息
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// Key回档文件名，回档专用
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Key回档文件扩展名，回档专用
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// Key回档文件大小，回档专用
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Key回档文件内容，回档专用
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type TableGroupInfo struct {

	// 表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格组名称
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格组创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 表格组包含的表格数量
	TableCount *uint64 `json:"TableCount,omitempty" name:"TableCount"`

	// 表格组包含的表格存储总量（MB）
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
}

type TableInfoNew struct {

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表格数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格所属集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 表格所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格所属表格组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格主键字段结构json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyStruct *string `json:"KeyStruct,omitempty" name:"KeyStruct"`

	// 表格非主键字段结构json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueStruct *string `json:"ValueStruct,omitempty" name:"ValueStruct"`

	// 表格分表因子集合，对PROTO类型表格有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// 表格索引键字段集合，对PROTO类型表格有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStruct *string `json:"IndexStruct,omitempty" name:"IndexStruct"`

	// LIST类型表格元素个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListElementNum *uint64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// 表格所关联IDL文件信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// 表格预留容量（GB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// 表格预留读QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// 表格预留写QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// 表格实际数据量大小（MB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableSize *int64 `json:"TableSize,omitempty" name:"TableSize"`

	// 表格状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 表格创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 表格最后一次修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 表格备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// TcaplusDB SDK数据访问接入ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// SORTLIST类型表格排序字段个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// SORTLIST类型表格排序顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`

	// 表格分布式索引信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbClusterInfoStruct *string `json:"DbClusterInfoStruct,omitempty" name:"DbClusterInfoStruct"`
}

type TableResultNew struct {

	// 表格实例ID，形如：tcaplus-3be64cbb
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 任务ID，对于创建单任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// 任务ID列表，对于创建多任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`
}

type TableRollbackResultNew struct {

	// 表格实例ID，形如：tcaplus-3be64cbb
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 任务ID，对于创建单任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表格数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// 任务ID列表，对于创建多任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`

	// 上传的key文件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 校验成功Key数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccKeyNum *uint64 `json:"SuccKeyNum,omitempty" name:"SuccKeyNum"`

	// Key文件中包含总的Key数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalKeyNum *uint64 `json:"TotalKeyNum,omitempty" name:"TotalKeyNum"`
}

type TaskInfoNew struct {

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务所关联的TcaplusDB内部事务ID
	TransId *string `json:"TransId,omitempty" name:"TransId"`

	// 任务所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 任务所属集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 任务进度
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 任务创建时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 操作者
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 任务详情
	Content *string `json:"Content,omitempty" name:"Content"`
}

type VerifyIdlFilesRequest struct {
	*tchttp.BaseRequest

	// 待创建表格的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建表格的表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 曾经上传过的IDL文件信息列表，与NewIdlFiles至少有一者
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles" list`

	// 待上传的IDL文件信息列表，与ExistingIdlFiles至少有一者
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles" list`
}

func (r *VerifyIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyIdlFilesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VerifyIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次上传校验所有的IDL文件信息列表
		IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

		// 读取IDL描述文件后解析出的合法表数量，不包含已经创建的表
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 读取IDL描述文件后解析出的合法表列表，不包含已经创建的表
		TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *VerifyIdlFilesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
