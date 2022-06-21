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

package v20180319

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DeregisterMigrationTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DeregisterMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeregisterMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterMigrationTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeregisterMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterMigrationTaskResponseParams `json:"Response"`
}

func (r *DeregisterMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTaskRequestParams struct {
	// 任务ID，例如msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，例如msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTaskResponseParams struct {
	// 迁移详情列表
	TaskStatus []*TaskStatus `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationTaskResponseParams `json:"Response"`
}

func (r *DescribeMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DstInfo struct {
	// 迁移目的地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 迁移目的Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 迁移目的端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 迁移目的实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type ListMigrationProjectRequestParams struct {
	// 记录起始数，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回条数，默认值为500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type ListMigrationProjectRequest struct {
	*tchttp.BaseRequest
	
	// 记录起始数，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回条数，默认值为500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListMigrationProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListMigrationProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListMigrationProjectResponseParams struct {
	// 项目列表
	Projects []*Project `json:"Projects,omitempty" name:"Projects"`

	// 项目总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListMigrationProjectResponse struct {
	*tchttp.BaseResponse
	Response *ListMigrationProjectResponseParams `json:"Response"`
}

func (r *ListMigrationProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListMigrationTaskRequestParams struct {
	// 记录起始数，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 记录条数，默认值为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 项目ID，默认值为空
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ListMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 记录起始数，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 记录条数，默认值为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 项目ID，默认值为空
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ListMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListMigrationTaskResponseParams struct {
	// 记录总条数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 迁移任务列表
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *ListMigrationTaskResponseParams `json:"Response"`
}

func (r *ListMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationTaskBelongToProjectRequestParams struct {
	// 任务ID，例如msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目ID，例如10005
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ModifyMigrationTaskBelongToProjectRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，例如msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目ID，例如10005
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyMigrationTaskBelongToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskBelongToProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationTaskBelongToProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationTaskBelongToProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMigrationTaskBelongToProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrationTaskBelongToProjectResponseParams `json:"Response"`
}

func (r *ModifyMigrationTaskBelongToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskBelongToProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationTaskStatusRequestParams struct {
	// 任务状态，取值为unstart，migrating，finish，fail之一，分别代表该迁移任务状态为迁移未开始，迁移中，迁移完成，迁移失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务ID，例如msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type ModifyMigrationTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务状态，取值为unstart，migrating，finish，fail之一，分别代表该迁移任务状态为迁移未开始，迁移中，迁移完成，迁移失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务ID，例如msp-jitoh33n
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *ModifyMigrationTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationTaskStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMigrationTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrationTaskStatusResponseParams `json:"Response"`
}

func (r *ModifyMigrationTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {
	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

// Predefined struct for user
type RegisterMigrationTaskRequestParams struct {
	// 任务类型，取值database（数据库迁移）、file（文件迁移）、host（主机迁移）
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 服务提供商名称
	ServiceSupplier *string `json:"ServiceSupplier,omitempty" name:"ServiceSupplier"`

	// 迁移任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 迁移任务更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 迁移类别，如数据库迁移中mysql:mysql代表从mysql迁移到mysql，文件迁移中oss:cos代表从阿里云oss迁移到腾讯云cos
	MigrateClass *string `json:"MigrateClass,omitempty" name:"MigrateClass"`

	// 迁移任务源信息
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 迁移任务目的信息
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 源实例接入类型，数据库迁移时填写值为：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例数据库类型，数据库迁移时填写，取值为mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb 之一
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 目标实例接入类型，数据库迁移时填写值为：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例数据库类型,数据库迁移时填写，取值为mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb 之一
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
}

type RegisterMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型，取值database（数据库迁移）、file（文件迁移）、host（主机迁移）
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 服务提供商名称
	ServiceSupplier *string `json:"ServiceSupplier,omitempty" name:"ServiceSupplier"`

	// 迁移任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 迁移任务更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 迁移类别，如数据库迁移中mysql:mysql代表从mysql迁移到mysql，文件迁移中oss:cos代表从阿里云oss迁移到腾讯云cos
	MigrateClass *string `json:"MigrateClass,omitempty" name:"MigrateClass"`

	// 迁移任务源信息
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 迁移任务目的信息
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 源实例接入类型，数据库迁移时填写值为：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源实例数据库类型，数据库迁移时填写，取值为mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb 之一
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 目标实例接入类型，数据库迁移时填写值为：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标实例数据库类型,数据库迁移时填写，取值为mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb 之一
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`
}

func (r *RegisterMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "TaskName")
	delete(f, "ServiceSupplier")
	delete(f, "CreateTime")
	delete(f, "UpdateTime")
	delete(f, "MigrateClass")
	delete(f, "SrcInfo")
	delete(f, "DstInfo")
	delete(f, "SrcAccessType")
	delete(f, "SrcDatabaseType")
	delete(f, "DstAccessType")
	delete(f, "DstDatabaseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterMigrationTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *RegisterMigrationTaskResponseParams `json:"Response"`
}

func (r *RegisterMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SrcInfo struct {
	// 迁移源地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 迁移源Ip
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 迁移源端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 迁移源实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type Task struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 迁移类型
	MigrationType *string `json:"MigrationType,omitempty" name:"MigrationType"`

	// 迁移状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 项目Id
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 迁移源信息
	SrcInfo *SrcInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 迁移时间信息
	MigrationTimeLine *TimeObj `json:"MigrationTimeLine,omitempty" name:"MigrationTimeLine"`

	// 状态更新时间
	Updated *string `json:"Updated,omitempty" name:"Updated"`

	// 迁移目的信息
	DstInfo *DstInfo `json:"DstInfo,omitempty" name:"DstInfo"`
}

type TaskStatus struct {
	// 迁移状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 迁移进度
	Progress *string `json:"Progress,omitempty" name:"Progress"`

	// 迁移日期
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TimeObj struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}