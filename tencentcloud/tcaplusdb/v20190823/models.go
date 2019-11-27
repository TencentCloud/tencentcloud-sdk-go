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

type AppInfo struct {

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 数据描述语言类型，如：`PROTO`,`TDR`或`MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 关联的用户私有网络实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 关联的用户子网实例ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 应用密码
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

type ClearTablesRequest struct {
	*tchttp.BaseRequest

	// 表所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待清理表信息列表
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`
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
		TableResults []*TableResult `json:"TableResults,omitempty" name:"TableResults" list`

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

type CompareIdlFilesRequest struct {
	*tchttp.BaseRequest

	// 待修改表所在应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待修改表列表
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`

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

		// 本次上传校验所有的Idl文件信息列表
		IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

		// 本次校验合法的表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 读取IDL描述文件后,根据用户指示的所选中表解析校验结果
		TableInfos []*ParsedTableInfo `json:"TableInfos,omitempty" name:"TableInfos" list`

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

type CreateAppRequest struct {
	*tchttp.BaseRequest

	// 应用数据描述语言类型，如：`PROTO`，`TDR`或`MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// 应用名称，可使用中文或英文字符，长度在30个字符以内
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 应用所绑定的私有网络实例ID，形如：vpc-f49l6u0z
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 应用所绑定的子网实例ID，形如：subnet-pxir56ns
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 应用访问密码，可使用英文和数字字符，长度为12~16个字符
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAppRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用ID
		ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAppResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTablesRequest struct {
	*tchttp.BaseRequest

	// 待创建表所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 用户选定的建表IDL文件列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// 待创建表信息列表
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`
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

		// 批量创建表结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 批量创建表结果列表
		TableResults []*TableResult `json:"TableResults,omitempty" name:"TableResults" list`

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

type CreateZoneRequest struct {
	*tchttp.BaseRequest

	// 大区所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 大区名称，可以采用中文、英文或数字字符，长度不能超过30
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 大区ID，可以由用户指定，但在同一个App内不能重复，如果不指定则采用自增的模式
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`
}

func (r *CreateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建的大区ID
		LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest

	// 待删除的应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DeleteAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAppRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除应用生成的任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAppResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteIdlFilesRequest struct {
	*tchttp.BaseRequest

	// 应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

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

type DeleteTablesRequest struct {
	*tchttp.BaseRequest

	// 待删除表所在应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待删除表信息列表
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`
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
		TableResults []*TableResult `json:"TableResults,omitempty" name:"TableResults" list`

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

type DeleteZoneRequest struct {
	*tchttp.BaseRequest

	// 大区所属的应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 大区ID
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`
}

func (r *DeleteZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteZoneRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除大区所创建的任务ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteZoneResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppsRequest struct {
	*tchttp.BaseRequest

	// 指定查询的应用ID
	ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds" list`

	// 查询过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 应用列表的大小，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用实例数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 应用实例列表
		Apps []*AppInfo `json:"Apps,omitempty" name:"Apps" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIdlFileInfosRequest struct {
	*tchttp.BaseRequest

	// 文件所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 文件所属大区ID
	LogicZoneIds []*string `json:"LogicZoneIds,omitempty" name:"LogicZoneIds" list`

	// 指定文件ID
	IdlFileIds []*string `json:"IdlFileIds,omitempty" name:"IdlFileIds" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 文件列表大小
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

type DescribeTablesInRecycleRequest struct {
	*tchttp.BaseRequest

	// 待查询表所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待查询表所属大区列表
	LogicZoneIds []*string `json:"LogicZoneIds,omitempty" name:"LogicZoneIds" list`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果列表数量
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

		// 表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表详情结果列表
		TableInfos []*TableInfo `json:"TableInfos,omitempty" name:"TableInfos" list`

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

	// 待查询表所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待查询表所属大区列表
	LogicZoneIds []*string `json:"LogicZoneIds,omitempty" name:"LogicZoneIds" list`

	// 待查询表信息列表
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果列表数量
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

		// 表数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表详情结果列表
		TableInfos []*TableInfo `json:"TableInfos,omitempty" name:"TableInfos" list`

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

	// 需要查询任务所属的应用ID列表
	ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds" list`

	// 需要查询的任务ID列表
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`

	// 过滤条件，本接口支持：Content，TaskType, Operator, Time
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 任务列表大小
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
		TaskInfos []*TaskInfo `json:"TaskInfos,omitempty" name:"TaskInfos" list`

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

type DescribeZonesRequest struct {
	*tchttp.BaseRequest

	// 大区所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 大区ID
	LogicZoneIds []*string `json:"LogicZoneIds,omitempty" name:"LogicZoneIds" list`

	// 过滤条件，本接口支持：ZoneName，ZoneId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 大区列表大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 大区数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 大区信息列表
		Zones []*ZoneInfo `json:"Zones,omitempty" name:"Zones" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
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

type ModifyAppNameRequest struct {
	*tchttp.BaseRequest

	// 需要修改名称的应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要修改的应用名称，需要URLEncode
	AppName *string `json:"AppName,omitempty" name:"AppName"`
}

func (r *ModifyAppNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAppNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAppNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAppNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAppPasswordRequest struct {
	*tchttp.BaseRequest

	// 需要修改密码的应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 应用旧密码预期失效时间
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// 应用新密码
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 更新模式： `1` 更新密码；`2` 更新旧密码失效时间，默认为`1` 模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyAppPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAppPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAppPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAppPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAppPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableMemosRequest struct {
	*tchttp.BaseRequest

	// 表所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 选定表详情列表
	TableMemos []*SelectedTableInfo `json:"TableMemos,omitempty" name:"TableMemos" list`
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
		TableResults []*TableResult `json:"TableResults,omitempty" name:"TableResults" list`

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

	// 带扩缩容表所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 已选中待修改的表配额列表
	TableQuotas []*SelectedTableInfo `json:"TableQuotas,omitempty" name:"TableQuotas" list`
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
		TableResults []*TableResult `json:"TableResults,omitempty" name:"TableResults" list`

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

	// 待修改表所在应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 选中的改表IDL文件
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// 待改表列表
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`
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
		TableResults []*TableResult `json:"TableResults,omitempty" name:"TableResults" list`

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

type ModifyZoneNameRequest struct {
	*tchttp.BaseRequest

	// 大区所属的应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待修改名称的大区ID
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

	// 新的大区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

func (r *ModifyZoneNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyZoneNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyZoneNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyZoneNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyZoneNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ParsedTableInfo struct {

	// 表描述语言类型：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表数据类型：`GENERIC`或`TDR`
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

	// 所属大区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

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
}

type RecoverRecycleTablesRequest struct {
	*tchttp.BaseRequest

	// 表所在应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待恢复表信息
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`
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
		TableResults []*TableResult `json:"TableResults,omitempty" name:"TableResults" list`

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

	// 待回档表所在应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待回档表列表
	SelectedTables []*SelectedTableInfo `json:"SelectedTables,omitempty" name:"SelectedTables" list`

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

		// 表回档任务结果数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 表回档任务结果列表
		TableResults []*TableRollbackResult `json:"TableResults,omitempty" name:"TableResults" list`

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

type SelectedTableInfo struct {

	// 表所属大区ID
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表实例ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表描述语言类型：`PROTO`或`TDR`
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表数据结构类型：`GENERIC`或`LIST`
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// LIST表元素个数
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// 表预留容量（GB）
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// 表预留读QPS
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// 表预留写QPS
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// 表备注信息
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

type TableInfo struct {

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表所属应用实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 表所属应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 表所属大区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

	// 表所属大区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 表主键结构json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyStruct *string `json:"KeyStruct,omitempty" name:"KeyStruct"`

	// 表非主键结构json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueStruct *string `json:"ValueStruct,omitempty" name:"ValueStruct"`

	// 表分表因子集合，PROTO表有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// 表索引键集合，PROTO表有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStruct *string `json:"IndexStruct,omitempty" name:"IndexStruct"`

	// LIST表元素个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListElementNum *uint64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// 表所关联IDL文件信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

	// 表预留容量（GB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// 表预留读QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// 表预留写QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// 表实际数据量大小（MB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableSize *int64 `json:"TableSize,omitempty" name:"TableSize"`

	// 表状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 表创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 最后一次更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 表备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// Api接入ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`
}

type TableResult struct {

	// 表实例ID，形如：tcaplus-3be64cbb
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 任务ID，对于创建单任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表所属大区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// 任务ID列表，对于创建多任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds" list`
}

type TableRollbackResult struct {

	// 表实例ID，形如：tcaplus-3be64cbb
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 任务ID，对于创建单任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表所属大区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

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

type TaskInfo struct {

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务所关联的TcaplusDB内部事务ID
	TransId *string `json:"TransId,omitempty" name:"TransId"`

	// 任务所属应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 任务所属应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

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

	// 待加表的应用实例ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 待加表的大区ID
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

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

		// 本次上传校验所有的Idl文件信息列表
		IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles" list`

		// 读取Idl描述文件后解析出的合法表数量，不包含已经创建的表
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 读取Idl描述文件后解析出的合法表列表，不包含已经创建的表
		TableInfos []*ParsedTableInfo `json:"TableInfos,omitempty" name:"TableInfos" list`

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

type ZoneInfo struct {

	// 大区ID
	LogicZoneId *string `json:"LogicZoneId,omitempty" name:"LogicZoneId"`

	// 大区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 大区创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 大区表格数量
	TableCount *uint64 `json:"TableCount,omitempty" name:"TableCount"`

	// 大区表格存储总量（MB）
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
}
