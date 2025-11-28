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

package v20190725

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AddNodeList struct {
	// 需要新增的节点角色。
	// - SECONDARY：Mongod 节点。
	// - READONLY：只读节点。
	// - MONGOS：Mongos 节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 节点所对应的可用区。当前支持的可用区，请参见[地域和可用区](https://cloud.tencent.com/document/product/240/3637)。
	// - 单可用区，所有节点在同一可用区。
	// - 多可用区：当前标准规格是三可用区分布，主从节点不在同一可用区，需注意配置新增节点对应的可用区，且新增后必须满足任意2个可用区节点数大于第3个可用区原则。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

// Predefined struct for user
type AssignProjectRequestParams struct {
	// 实例 ID 列表，请登录[MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 项目ID，用户已创建项目的唯一ID。请在控制台账号中心的[项目管理](https://console.cloud.tencent.com/project)中复制项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type AssignProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表，请登录[MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 项目ID，用户已创建项目的唯一ID。请在控制台账号中心的[项目管理](https://console.cloud.tencent.com/project)中复制项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *AssignProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignProjectResponseParams struct {
	// 返回的异步任务ID列表。
	FlowIds []*uint64 `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssignProjectResponse struct {
	*tchttp.BaseResponse
	Response *AssignProjectResponseParams `json:"Response"`
}

func (r *AssignProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuditInstance struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计状态。
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 是否存在审计任务，0：无任务，1：创建中，2：关闭中
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditTask *int64 `json:"AuditTask,omitnil,omitempty" name:"AuditTask"`

	// 审计日志过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighLogExpireDay *int64 `json:"HighLogExpireDay,omitnil,omitempty" name:"HighLogExpireDay"`

	// 低频日志过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowLogExpireDay *int64 `json:"LowLogExpireDay,omitnil,omitempty" name:"LowLogExpireDay"`

	// 费用信息。
	BillingAmount *float64 `json:"BillingAmount,omitnil,omitempty" name:"BillingAmount"`

	// 高频存储容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighRealStorage *float64 `json:"HighRealStorage,omitnil,omitempty" name:"HighRealStorage"`

	// 低频存储容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowRealStorage *float64 `json:"LowRealStorage,omitnil,omitempty" name:"LowRealStorage"`

	// 实例详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 性能分析
	// 注意：此字段可能返回 null，表示取不到有效值。
	PerformancesAnalyse *int64 `json:"PerformancesAnalyse,omitnil,omitempty" name:"PerformancesAnalyse"`

	// true表示全审计，false表示规则审计
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 实例审计最近一次的开通时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 实例绑定的规则模版ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTemplateIds []*string `json:"RuleTemplateIds,omitnil,omitempty" name:"RuleTemplateIds"`

	// 是否开启投递：ON，OFF
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deliver *string `json:"Deliver,omitnil,omitempty" name:"Deliver"`

	// 日志投递信息
	DeliverSummary []*DeliverSummary `json:"DeliverSummary,omitnil,omitempty" name:"DeliverSummary"`

	// 旧规则
	OldRule *bool `json:"OldRule,omitnil,omitempty" name:"OldRule"`

	// 实际存储容量
	RealStorage *float64 `json:"RealStorage,omitnil,omitempty" name:"RealStorage"`
}

type AuditLogFilter struct {
	// 客户端地址。
	Host []*string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户名。
	User []*string `json:"User,omitnil,omitempty" name:"User"`

	// 执行时间。单位为：ms。表示筛选执行时间大于该值的审计日志。
	ExecTime *uint64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// 影响行数。表示筛选影响行数大于该值的审计日志。
	AffectRows *uint64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// 操作类型。
	Atype []*string `json:"Atype,omitnil,omitempty" name:"Atype"`

	// 执行结果。
	Result []*string `json:"Result,omitnil,omitempty" name:"Result"`

	// 根据此关键字过滤日志
	Param []*string `json:"Param,omitnil,omitempty" name:"Param"`
}

type Auth struct {
	// 当前账号具有的权限信息。
	// - 0：无权限。
	// - 1：只读。
	// - 3：读写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mask *int64 `json:"Mask,omitnil,omitempty" name:"Mask"`

	// 指具有当前账号权限的数据库名。
	// - \* ：表示所有数据库。
	// - db.name：表示特定 name 的数据库。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameSpace *string `json:"NameSpace,omitnil,omitempty" name:"NameSpace"`
}

type BackupDownloadTask struct {
	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 备份文件名。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 分片名称。
	ReplicaSetId *string `json:"ReplicaSetId,omitnil,omitempty" name:"ReplicaSetId"`

	// 备份数据大小，单位：字节。
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// 任务状态。
	// - 0：等待执行。
	// - 1：正在下载。
	// - 2：下载完成。
	// - 3：下载失败。
	// - 4：等待重试。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务进度百分比。
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 耗时，单位为秒。
	TimeSpend *int64 `json:"TimeSpend,omitnil,omitempty" name:"TimeSpend"`

	// 备份数据下载链接。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明**:
	// 1. 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// 2. 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 发起备份时指定的备注信息。
	BackupDesc *string `json:"BackupDesc,omitnil,omitempty" name:"BackupDesc"`

	// 地区信息。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Bucket信息。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

type BackupDownloadTaskStatus struct {
	// 分片名。
	ReplicaSetId *string `json:"ReplicaSetId,omitnil,omitempty" name:"ReplicaSetId"`

	// 任务当前状态。
	// - 0：等待执行。
	// - 1：正在下载。
	// - 2：下载完成。
	// - 3：下载失败。
	// - 4：等待重试。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type BackupInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份方式。
	// - 0：自动备份。
	// - 1：手动备份。
	BackupType *uint64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份文件名称。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 备份任务备注信息。
	BackupDesc *string `json:"BackupDesc,omitnil,omitempty" name:"BackupDesc"`

	// 备份文件大小，单位：KB。
	BackupSize *uint64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// 备份开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份状态。
	// - 1：备份中。
	// - 2：备份成功。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明:**
	// - 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// - 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *uint64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份记录 ID。
	BackId *int64 `json:"BackId,omitnil,omitempty" name:"BackId"`

	// 备份删除时间。
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// 异地备份地域。
	BackupRegion *string `json:"BackupRegion,omitnil,omitempty" name:"BackupRegion"`

	// 备份支持的回档时间。
	RestoreTime *string `json:"RestoreTime,omitnil,omitempty" name:"RestoreTime"`
}

type ClientConnection struct {
	// 连接的客户端 IP。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 对应客户端 IP 的连接数。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 是否为内部 IP。
	InternalService *bool `json:"InternalService,omitnil,omitempty" name:"InternalService"`
}

// Predefined struct for user
type CreateAccountUserRequestParams struct {
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新账号名称。其格式要求如下：
	// - 字符范围[1,64]。
	// - 可输入[A,Z]、[a,z]、[1,9]范围的字符以及下划线“\_”与短划线“-”。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 新账号密码。密码复杂度要求如下：
	// - 字符长度范围[8,32]。
	// - 至少包含字母、数字和特殊字符（叹号“!”、at"@"、井号“#”、百分号“%”、插入符“^”、星号“\*”、小括号“()”、下划线“\_”）中的两种。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// mongouser 账号对应的密码。mongouser 为系统默认账号，即为创建实例时，设置的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil,omitempty" name:"MongoUserPassword"`

	// 账号备注信息。
	UserDesc *string `json:"UserDesc,omitnil,omitempty" name:"UserDesc"`

	// 账号的读写权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil,omitempty" name:"AuthRole"`
}

type CreateAccountUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新账号名称。其格式要求如下：
	// - 字符范围[1,64]。
	// - 可输入[A,Z]、[a,z]、[1,9]范围的字符以及下划线“\_”与短划线“-”。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 新账号密码。密码复杂度要求如下：
	// - 字符长度范围[8,32]。
	// - 至少包含字母、数字和特殊字符（叹号“!”、at"@"、井号“#”、百分号“%”、插入符“^”、星号“\*”、小括号“()”、下划线“\_”）中的两种。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// mongouser 账号对应的密码。mongouser 为系统默认账号，即为创建实例时，设置的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil,omitempty" name:"MongoUserPassword"`

	// 账号备注信息。
	UserDesc *string `json:"UserDesc,omitnil,omitempty" name:"UserDesc"`

	// 账号的读写权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil,omitempty" name:"AuthRole"`
}

func (r *CreateAccountUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "MongoUserPassword")
	delete(f, "UserDesc")
	delete(f, "AuthRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountUserResponseParams struct {
	// 创建任务ID。
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccountUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountUserResponseParams `json:"Response"`
}

func (r *CreateAccountUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileRequestParams struct {
	// 实例 ID，格式如：cmgo-xfts****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式为："2021-07-12 10:29:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式为："2021-07-12 10:39:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 审计日志文件的排序方式。
	// <ul><li>ASC：升序。</li><li>DESC：降序。</li></ul>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 审计日志文件的排序字段。当前支持的取值包括：
	// <ul><li>timestamp：时间戳。</li><li>affectRows：影响行数。</li><li>execTime：执行时间。</li></ul>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤条件。可按设置的过滤条件过滤审计日志。
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cmgo-xfts****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式为："2021-07-12 10:29:20"。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式为："2021-07-12 10:39:20"。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 审计日志文件的排序方式。
	// <ul><li>ASC：升序。</li><li>DESC：降序。</li></ul>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 审计日志文件的排序字段。当前支持的取值包括：
	// <ul><li>timestamp：时间戳。</li><li>affectRows：影响行数。</li><li>execTime：执行时间。</li></ul>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤条件。可按设置的过滤条件过滤审计日志。
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "OrderBy")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileResponseParams struct {
	// 审计日志文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditLogFileResponseParams `json:"Response"`
}

func (r *CreateAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDBInstanceRequestParams struct {
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明**:
	// 1. 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// 2. 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份备注信息。
	BackupRemark *string `json:"BackupRemark,omitnil,omitempty" name:"BackupRemark"`
}

type CreateBackupDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明**:
	// 1. 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// 2. 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 备份备注信息。
	BackupRemark *string `json:"BackupRemark,omitnil,omitempty" name:"BackupRemark"`
}

func (r *CreateBackupDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "BackupRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDBInstanceResponseParams struct {
	// 查询备份流程的状态。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupDBInstanceResponseParams `json:"Response"`
}

func (r *CreateBackupDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDownloadTaskRequestParams struct {
	// 实例ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要下载的备份文件名。请通过 [DescribeDBBackups](https://cloud.tencent.com/document/product/240/38574) 接口获取。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 指定要下载的副本集节点 ID 或分片集群的分片节点 ID 列表。
	// - 如副本集实例 ID 为 cmgo-p8vnipr5，示例：BackupSets.0=cmgo-p8vnipr5_0，可下载全量数据。
	// - 如分片集群实例 ID 为 cmgo-p8vnipr5，示例：BackupSets.0=cmgo-p8vnipr5_0&BackupSets.1=cmgo-p8vnipr5_1，即下载分片0和分片1的数据。分片集群如需全量下载，请按示例方式传入全部分片名称。
	BackupSets []*ReplicaSetInfo `json:"BackupSets,omitnil,omitempty" name:"BackupSets"`
}

type CreateBackupDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要下载的备份文件名。请通过 [DescribeDBBackups](https://cloud.tencent.com/document/product/240/38574) 接口获取。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 指定要下载的副本集节点 ID 或分片集群的分片节点 ID 列表。
	// - 如副本集实例 ID 为 cmgo-p8vnipr5，示例：BackupSets.0=cmgo-p8vnipr5_0，可下载全量数据。
	// - 如分片集群实例 ID 为 cmgo-p8vnipr5，示例：BackupSets.0=cmgo-p8vnipr5_0&BackupSets.1=cmgo-p8vnipr5_1，即下载分片0和分片1的数据。分片集群如需全量下载，请按示例方式传入全部分片名称。
	BackupSets []*ReplicaSetInfo `json:"BackupSets,omitnil,omitempty" name:"BackupSets"`
}

func (r *CreateBackupDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	delete(f, "BackupSets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupDownloadTaskResponseParams struct {
	// 下载任务状态。
	Tasks []*BackupDownloadTaskStatus `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupDownloadTaskResponseParams `json:"Response"`
}

func (r *CreateBackupDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourRequestParams struct {
	// 实例内存大小，单位：GB。具体售卖的内存规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。每一个 CPU 规格对应的最大磁盘与最小磁盘范围，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// - 创建副本集实例，指副本集数量，该参数只能为1。
	// - 创建分片集群实例，指分片的数量。请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询分片数量的取值范围，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// - 创建副本集实例，指每个副本集内主从节点数量。每个副本集所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 创建分片集群实例，指每个分片的主从节点数量。每个分片所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 指版本信息。具体支持的版本信息 ，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 产品规格类型。
	// - HIO10G：通用高HIO万兆型。
	// - HCD：云盘版类型。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例数量，最小值1，最大值为30。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 私有网络ID。
	// - 仅支持配置私有网络，必须选择一个与实例同一地域的私有网络。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的私有网络 ID。
	// - 实例创建成功之后，支持更换私有网络。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络 VPC 的子网 ID。
	// - 必须在已选的私有网络内指定一个子网。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的子网 ID。
	// - 实例创建成功之后，支持更换私有网络及子网。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例密码。设置要求如下：
	// - 字符个数为[8,32]。
	// - 可输入[A,Z]、[a,z]、[0,9]范围内的字符。
	// - 可输入的特殊字符包括：感叹号“!”，at“@”，警号“#”、百分号“%”、插入号“^”、星号“\*”、括号“()”、下划线“_”。
	// - 不能设置单一的字母或者数字。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 项目ID。
	// - 若不设置该参数，则为默认项目。
	// - 在 [MongoDB 控制台项目管理](https://console.cloud.tencent.com/project)页面，可获取项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	// - 5：克隆实例。注意：克隆实例 RestoreTime 为必填项。
	Clone *int64 `json:"Clone,omitnil,omitempty" name:"Clone"`

	// 父实例 ID。
	// - 当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制父实例 ID。
	Father *string `json:"Father,omitnil,omitempty" name:"Father"`

	// 安全组 ID。 请登录[安全组控制台](https://console.cloud.tencent.com/vpc/security-group)页面获取与数据库实例同地域的安全组 ID。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 克隆实例回档时间。
	// - 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。
	// - 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil,omitempty" name:"RestoreTime"`

	// 实例名称。仅支持长度为128个字符的中文、英文、数字、下划线\_、分隔符\-。批量购买数据库实例时，支持通过自定义命名模式串与数字后缀自动升序功能，高效设置实例名称。
	// - 基础模式：前缀＋自动升序编号（默认从1开始），**lnstanceName**仅需自定义实例名称前缀，例如设置为：cmgo，设置购买数量为5，则购买后，实例名称依次分别为cmgo1、cmgo2、cmgo3、cmgo4、cmgo5。
	// - 自定义起始序号模式：前缀+｛R:x｝（x为自定义起始序号）。**InstanceName**需填写“前缀｛R:x｝”，例如：cmgo｛R:3｝，设置购买数量为5，则实例名称为cmgo3、cmgo4、cmgo5、cmgo6、cmgo7。
	// - 复合模式串：前缀1{R:x}+前缀2{R:y}+ ⋯+固定后缀，x与y分别为每一段前缀的起始序号。**instanceName**需填写复合模式串，例如：cmgo{R:10}\_node{R:12}\_db，设置批量购买数量为5，则实例名称为 cmgo10\_node12\_db, cmgo11\_node13\_db, cmgo12\_node14\_db, cmgo13\_node15\_db, cluster14\_node16\_db. 
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 若多可用区部署云数据库实例，指定多可用区列表。
	// - 多可用区部署实例，参数 **Zone** 指定实例主可用区信息；**AvailabilityZoneList** 指定所有可用区信息，包含主可用区。输入格式如：[ap-guangzhou-2,ap-guangzhou-3,ap-guangzhou-4]。
	// - 通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 可获取云数据库不同地域规划的可用区信息，以便指定有效的可用区。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil,omitempty" name:"AvailabilityZoneList"`

	// Mongos CPU 核数，支持1、2、4、8、16。购买分片集群时，必须填写。
	MongosCpu *uint64 `json:"MongosCpu,omitnil,omitempty" name:"MongosCpu"`

	// Mongos 内存大小。
	// -  购买分片集群时，必须填写。
	// - 单位：GB，支持1核2GB、2核4GB、4核8GB、8核16GB、16核32GB。
	MongosMemory *uint64 `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// Mongos 数量。购买分片集群时，必须填写。
	// - 单可用区部署实例，其数量范围为[3,32]。
	// - 多可用区部署实例，其数量范围为[6,32]。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil,omitempty" name:"MongosNodeNum"`

	// 只读节点数量，取值范围[0,5]。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil,omitempty" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区数组。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil,omitempty" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil,omitempty" name:"HiddenZone"`

	// 参数模板 ID。
	// - 参数模板是预置了特定参数值的集合，可用于快速配置新的 MongoDB 实例。合理使用参数模板，能有效提升数据库的部署效率与运行性能。
	// - 参数模板 ID 可通过 [DescribeDBInstanceParamTpl ](https://cloud.tencent.com/document/product/240/109155)接口获取。请选择与实例版本与架构所对应的参数模板 ID。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// 实例内存大小，单位：GB。具体售卖的内存规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。每一个 CPU 规格对应的最大磁盘与最小磁盘范围，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// - 创建副本集实例，指副本集数量，该参数只能为1。
	// - 创建分片集群实例，指分片的数量。请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询分片数量的取值范围，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// - 创建副本集实例，指每个副本集内主从节点数量。每个副本集所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 创建分片集群实例，指每个分片的主从节点数量。每个分片所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 指版本信息。具体支持的版本信息 ，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 产品规格类型。
	// - HIO10G：通用高HIO万兆型。
	// - HCD：云盘版类型。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例数量，最小值1，最大值为30。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 私有网络ID。
	// - 仅支持配置私有网络，必须选择一个与实例同一地域的私有网络。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的私有网络 ID。
	// - 实例创建成功之后，支持更换私有网络。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络 VPC 的子网 ID。
	// - 必须在已选的私有网络内指定一个子网。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的子网 ID。
	// - 实例创建成功之后，支持更换私有网络及子网。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例密码。设置要求如下：
	// - 字符个数为[8,32]。
	// - 可输入[A,Z]、[a,z]、[0,9]范围内的字符。
	// - 可输入的特殊字符包括：感叹号“!”，at“@”，警号“#”、百分号“%”、插入号“^”、星号“\*”、括号“()”、下划线“_”。
	// - 不能设置单一的字母或者数字。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 项目ID。
	// - 若不设置该参数，则为默认项目。
	// - 在 [MongoDB 控制台项目管理](https://console.cloud.tencent.com/project)页面，可获取项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	// - 5：克隆实例。注意：克隆实例 RestoreTime 为必填项。
	Clone *int64 `json:"Clone,omitnil,omitempty" name:"Clone"`

	// 父实例 ID。
	// - 当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制父实例 ID。
	Father *string `json:"Father,omitnil,omitempty" name:"Father"`

	// 安全组 ID。 请登录[安全组控制台](https://console.cloud.tencent.com/vpc/security-group)页面获取与数据库实例同地域的安全组 ID。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 克隆实例回档时间。
	// - 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。
	// - 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil,omitempty" name:"RestoreTime"`

	// 实例名称。仅支持长度为128个字符的中文、英文、数字、下划线\_、分隔符\-。批量购买数据库实例时，支持通过自定义命名模式串与数字后缀自动升序功能，高效设置实例名称。
	// - 基础模式：前缀＋自动升序编号（默认从1开始），**lnstanceName**仅需自定义实例名称前缀，例如设置为：cmgo，设置购买数量为5，则购买后，实例名称依次分别为cmgo1、cmgo2、cmgo3、cmgo4、cmgo5。
	// - 自定义起始序号模式：前缀+｛R:x｝（x为自定义起始序号）。**InstanceName**需填写“前缀｛R:x｝”，例如：cmgo｛R:3｝，设置购买数量为5，则实例名称为cmgo3、cmgo4、cmgo5、cmgo6、cmgo7。
	// - 复合模式串：前缀1{R:x}+前缀2{R:y}+ ⋯+固定后缀，x与y分别为每一段前缀的起始序号。**instanceName**需填写复合模式串，例如：cmgo{R:10}\_node{R:12}\_db，设置批量购买数量为5，则实例名称为 cmgo10\_node12\_db, cmgo11\_node13\_db, cmgo12\_node14\_db, cmgo13\_node15\_db, cluster14\_node16\_db. 
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 若多可用区部署云数据库实例，指定多可用区列表。
	// - 多可用区部署实例，参数 **Zone** 指定实例主可用区信息；**AvailabilityZoneList** 指定所有可用区信息，包含主可用区。输入格式如：[ap-guangzhou-2,ap-guangzhou-3,ap-guangzhou-4]。
	// - 通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 可获取云数据库不同地域规划的可用区信息，以便指定有效的可用区。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil,omitempty" name:"AvailabilityZoneList"`

	// Mongos CPU 核数，支持1、2、4、8、16。购买分片集群时，必须填写。
	MongosCpu *uint64 `json:"MongosCpu,omitnil,omitempty" name:"MongosCpu"`

	// Mongos 内存大小。
	// -  购买分片集群时，必须填写。
	// - 单位：GB，支持1核2GB、2核4GB、4核8GB、8核16GB、16核32GB。
	MongosMemory *uint64 `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// Mongos 数量。购买分片集群时，必须填写。
	// - 单可用区部署实例，其数量范围为[3,32]。
	// - 多可用区部署实例，其数量范围为[6,32]。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil,omitempty" name:"MongosNodeNum"`

	// 只读节点数量，取值范围[0,5]。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil,omitempty" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区数组。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil,omitempty" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil,omitempty" name:"HiddenZone"`

	// 参数模板 ID。
	// - 参数模板是预置了特定参数值的集合，可用于快速配置新的 MongoDB 实例。合理使用参数模板，能有效提升数据库的部署效率与运行性能。
	// - 参数模板 ID 可通过 [DescribeDBInstanceParamTpl ](https://cloud.tencent.com/document/product/240/109155)接口获取。请选择与实例版本与架构所对应的参数模板 ID。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "ReplicateSetNum")
	delete(f, "NodeNum")
	delete(f, "MongoVersion")
	delete(f, "MachineCode")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "ClusterType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "ProjectId")
	delete(f, "Tags")
	delete(f, "Clone")
	delete(f, "Father")
	delete(f, "SecurityGroup")
	delete(f, "RestoreTime")
	delete(f, "InstanceName")
	delete(f, "AvailabilityZoneList")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNodeNum")
	delete(f, "ReadonlyNodeNum")
	delete(f, "ReadonlyNodeAvailabilityZoneList")
	delete(f, "HiddenZone")
	delete(f, "ParamTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// 订单ID。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 创建的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceHourResponseParams `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceParamTplRequestParams struct {
	// 参数模板名称。
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 参数模板版本号。当**MirrorTplId**为空时，该字段必填。参数模板支持的售卖版本，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/35767) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本，
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 实例类型。当 MirrorTplId 为空值时，该参数必填。
	// - REPLSET：副本集实例。
	// - SHARD：分片实例。
	// - STANDALONE：单节点实例。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 模板描述信息。
	TplDesc *string `json:"TplDesc,omitnil,omitempty" name:"TplDesc"`

	// 模板参数，若不配置该参数，则以系统默认模板作为新版本参数。
	Params []*ParamType `json:"Params,omitnil,omitempty" name:"Params"`

	// 镜像模板 ID。若指定镜像模板，则以该模板为镜像，克隆出一个新的模板。
	// **注意**：MirrorTplId 不为空值时，MongoVersion 及 ClusterType 将以 MirrorTpl 模板的版本及实例类型为准。
	MirrorTplId *string `json:"MirrorTplId,omitnil,omitempty" name:"MirrorTplId"`
}

type CreateDBInstanceParamTplRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板名称。
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 参数模板版本号。当**MirrorTplId**为空时，该字段必填。参数模板支持的售卖版本，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/35767) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本，
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 实例类型。当 MirrorTplId 为空值时，该参数必填。
	// - REPLSET：副本集实例。
	// - SHARD：分片实例。
	// - STANDALONE：单节点实例。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 模板描述信息。
	TplDesc *string `json:"TplDesc,omitnil,omitempty" name:"TplDesc"`

	// 模板参数，若不配置该参数，则以系统默认模板作为新版本参数。
	Params []*ParamType `json:"Params,omitnil,omitempty" name:"Params"`

	// 镜像模板 ID。若指定镜像模板，则以该模板为镜像，克隆出一个新的模板。
	// **注意**：MirrorTplId 不为空值时，MongoVersion 及 ClusterType 将以 MirrorTpl 模板的版本及实例类型为准。
	MirrorTplId *string `json:"MirrorTplId,omitnil,omitempty" name:"MirrorTplId"`
}

func (r *CreateDBInstanceParamTplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceParamTplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TplName")
	delete(f, "MongoVersion")
	delete(f, "ClusterType")
	delete(f, "TplDesc")
	delete(f, "Params")
	delete(f, "MirrorTplId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceParamTplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceParamTplResponseParams struct {
	// 模板ID
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceParamTplResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceParamTplResponseParams `json:"Response"`
}

func (r *CreateDBInstanceParamTplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceParamTplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceRequestParams struct {
	// - 创建副本集实例，指每个副本集内主从节点数量。每个副本集所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 创建分片集群实例，指每个分片的主从节点数量。每个分片所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例内存大小，单位：GB。具体售卖的内存规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。每一个 CPU 规格对应的最大磁盘与最小磁盘范围，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 指版本信息。具体支持的版本信息 ，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 实例数量, 最小值1，最大值为30。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 指定购买实例的购买时长。取值可选：[1,2,3,4,5,6,7,8,9,10,11,12,24,36]；单位：月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 产品规格类型。
	// - HIO10G：通用高HIO万兆型。
	// - HCD：云盘版类型。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// - 创建副本集实例，指副本集数量，该参数只能为1。
	// - 创建分片集群实例，指分片的数量。请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询分片数量的取值范围，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// 项目ID。
	// - 若不设置该参数，则为默认项目。
	// - 在 [MongoDB 控制台项目管理](https://console.cloud.tencent.com/project)页面，可获取项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 私有网络 ID。
	// - 仅支持配置私有网络，必须选择一个与实例同一地域的私有网络。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的私有网络 ID。
	// - 实例创建成功之后，支持更换私有网络。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络 VPC 的子网 ID。
	// - 必须在已选的私有网络内指定一个子网。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的子网 ID。
	// - 实例创建成功之后，支持更换私有网络及子网。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例密码。设置要求如下：
	// - 字符个数为[8,32]。
	// - 可输入[A,Z]、[a,z]、[0,9]范围内的字符。
	// - 可输入的特殊字符包括：感叹号“!”，at“@”，警号“#”、百分号“%”、插入号“^”、星号“\*”、括号“()”、下划线“\_”。
	// - 不能设置单一的字母或者数字。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自动续费标记。
	// - 0：不自动续费。
	// - 1：自动续费。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券。
	// - 1：是。
	// - 0：否。默认为0。
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	// - 5：克隆实例。注意：克隆实例 RestoreTime 为必填项。
	Clone *int64 `json:"Clone,omitnil,omitempty" name:"Clone"`

	// 父实例 ID。
	// - 当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制父实例 ID。
	Father *string `json:"Father,omitnil,omitempty" name:"Father"`

	// 安全组 ID。 请登录[安全组控制台](https://console.cloud.tencent.com/vpc/security-group)页面获取与数据库实例同地域的安全组 ID。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 克隆实例回档时间，当Clone取值为5或6时为必填。- 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。- 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil,omitempty" name:"RestoreTime"`

	// 实例名称。仅支持长度为128个字符的中文、英文、数字、下划线\_、分隔符\-。批量购买数据库实例时，支持通过自定义命名模式串与数字后缀自动升序功能，高效设置实例名称。
	// - 基础模式：前缀＋自动升序编号（默认从1开始），**lnstanceName**仅需自定义实例名称前缀，例如设置为：cmgo，设置购买数量为5，则购买后，实例名称依次分别为cmgo1、cmgo2、cmgo3、cmgo4、cmgo5。
	// - 自定义起始序号模式：前缀+｛R:x｝（x为自定义起始序号）。**InstanceName**需填写“前缀｛R:x｝”，例如：cmgo｛R:3｝，设置购买数量为5，则实例名称为cmgo3、cmgo4、cmgo5、cmgo6、cmgo7。
	// - 复合模式串：前缀1{R:x}+前缀2{R:y}+ ⋯+固定后缀，x与y分别为每一段前缀的起始序号。**instanceName**需填写复合模式串，例如：cmgo{R:10}\_node{R:12}\_db，设置批量购买数量为5，则实例名称为 cmgo10\_node12\_db, cmgo11\_node13\_db, cmgo12\_node14\_db, cmgo13\_node15\_db, cluster14\_node16\_db. 
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 若多可用区部署云数据库实例，指定多可用区列表。
	// - 多可用区部署实例，参数 **Zone** 指定实例主可用区信息；**AvailabilityZoneList** 指定所有可用区信息，包含主可用区。输入格式如：[ap-guangzhou-2,ap-guangzhou-3,ap-guangzhou-4]。
	// - 通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 可获取云数据库不同地域规划的可用区信息，以便指定有效的可用区。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil,omitempty" name:"AvailabilityZoneList"`

	// Mongos CPU 核数，支持1、2、4、8、16。购买分片集群时，必须填写。
	MongosCpu *uint64 `json:"MongosCpu,omitnil,omitempty" name:"MongosCpu"`

	// Mongos 内存大小。
	// -  购买分片集群时，必须填写。
	// - 单位：GB，支持1核2GB、2核4GB、4核8GB、8核16GB、16核32GB。
	MongosMemory *uint64 `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// Mongos 数量。购买分片集群时，必须填写。
	// - 单可用区部署实例，其数量范围为[3,32]。
	// - 多可用区部署实例，其数量范围为[6,32]。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil,omitempty" name:"MongosNodeNum"`

	// 只读节点数量，取值范围[0,5]。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil,omitempty" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区数组。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil,omitempty" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil,omitempty" name:"HiddenZone"`

	// 参数模板 ID。
	// - 参数模板是预置了特定参数值的集合，可用于快速配置新的 MongoDB 实例。合理使用参数模板，能有效提升数据库的部署效率与运行性能。
	// - 参数模板 ID 可通过 [DescribeDBInstanceParamTpl ](https://cloud.tencent.com/document/product/240/109155)接口获取。请选择与实例版本与架构所对应的参数模板 ID。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// - 创建副本集实例，指每个副本集内主从节点数量。每个副本集所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 创建分片集群实例，指每个分片的主从节点数量。每个分片所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例内存大小，单位：GB。具体售卖的内存规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB。每一个 CPU 规格对应的最大磁盘与最小磁盘范围，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 指版本信息。具体支持的版本信息 ，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 实例数量, 最小值1，最大值为30。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，输入格式如：ap-guangzhou-2。
	// - 具体信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 该参数为主可用区，如果多可用区部署，Zone必须是AvailabilityZoneList中的一个。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 指定购买实例的购买时长。取值可选：[1,2,3,4,5,6,7,8,9,10,11,12,24,36]；单位：月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 产品规格类型。
	// - HIO10G：通用高HIO万兆型。
	// - HCD：云盘版类型。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例架构类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// - 创建副本集实例，指副本集数量，该参数只能为1。
	// - 创建分片集群实例，指分片的数量。请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询分片数量的取值范围，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// 项目ID。
	// - 若不设置该参数，则为默认项目。
	// - 在 [MongoDB 控制台项目管理](https://console.cloud.tencent.com/project)页面，可获取项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 私有网络 ID。
	// - 仅支持配置私有网络，必须选择一个与实例同一地域的私有网络。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的私有网络 ID。
	// - 实例创建成功之后，支持更换私有网络。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络 VPC 的子网 ID。
	// - 必须在已选的私有网络内指定一个子网。请登录[私有网络控制台](https://console.cloud.tencent.com/vpc)获取可使用的子网 ID。
	// - 实例创建成功之后，支持更换私有网络及子网。具体操作，请参见[更换网络](https://cloud.tencent.com/document/product/239/30910)。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例密码。设置要求如下：
	// - 字符个数为[8,32]。
	// - 可输入[A,Z]、[a,z]、[0,9]范围内的字符。
	// - 可输入的特殊字符包括：感叹号“!”，at“@”，警号“#”、百分号“%”、插入号“^”、星号“\*”、括号“()”、下划线“\_”。
	// - 不能设置单一的字母或者数字。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 实例标签信息。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自动续费标记。
	// - 0：不自动续费。
	// - 1：自动续费。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券。
	// - 1：是。
	// - 0：否。默认为0。
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 实例类型。
	// - 1：正式实例。
	// - 3：只读实例。
	// - 4：灾备实例。
	// - 5：克隆实例。注意：克隆实例 RestoreTime 为必填项。
	Clone *int64 `json:"Clone,omitnil,omitempty" name:"Clone"`

	// 父实例 ID。
	// - 当参数**Clone**为3或者4时，即实例为只读或灾备实例时，该参数必须配置。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制父实例 ID。
	Father *string `json:"Father,omitnil,omitempty" name:"Father"`

	// 安全组 ID。 请登录[安全组控制台](https://console.cloud.tencent.com/vpc/security-group)页面获取与数据库实例同地域的安全组 ID。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 克隆实例回档时间，当Clone取值为5或6时为必填。- 若为克隆实例，则必须配置该参数。输入格式示例：2021-08-13 16:30:00。- 回档时间范围：仅能回档7天内时间点的数据。
	RestoreTime *string `json:"RestoreTime,omitnil,omitempty" name:"RestoreTime"`

	// 实例名称。仅支持长度为128个字符的中文、英文、数字、下划线\_、分隔符\-。批量购买数据库实例时，支持通过自定义命名模式串与数字后缀自动升序功能，高效设置实例名称。
	// - 基础模式：前缀＋自动升序编号（默认从1开始），**lnstanceName**仅需自定义实例名称前缀，例如设置为：cmgo，设置购买数量为5，则购买后，实例名称依次分别为cmgo1、cmgo2、cmgo3、cmgo4、cmgo5。
	// - 自定义起始序号模式：前缀+｛R:x｝（x为自定义起始序号）。**InstanceName**需填写“前缀｛R:x｝”，例如：cmgo｛R:3｝，设置购买数量为5，则实例名称为cmgo3、cmgo4、cmgo5、cmgo6、cmgo7。
	// - 复合模式串：前缀1{R:x}+前缀2{R:y}+ ⋯+固定后缀，x与y分别为每一段前缀的起始序号。**instanceName**需填写复合模式串，例如：cmgo{R:10}\_node{R:12}\_db，设置批量购买数量为5，则实例名称为 cmgo10\_node12\_db, cmgo11\_node13\_db, cmgo12\_node14\_db, cmgo13\_node15\_db, cluster14\_node16\_db. 
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 若多可用区部署云数据库实例，指定多可用区列表。
	// - 多可用区部署实例，参数 **Zone** 指定实例主可用区信息；**AvailabilityZoneList** 指定所有可用区信息，包含主可用区。输入格式如：[ap-guangzhou-2,ap-guangzhou-3,ap-guangzhou-4]。
	// - 通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 可获取云数据库不同地域规划的可用区信息，以便指定有效的可用区。
	// - 多可用区部署节点只能部署在3个不同可用区。不支持将集群的大多数节点部署在同一个可用区。例如：3节点集群不支持2个节点部署在同一个区。
	AvailabilityZoneList []*string `json:"AvailabilityZoneList,omitnil,omitempty" name:"AvailabilityZoneList"`

	// Mongos CPU 核数，支持1、2、4、8、16。购买分片集群时，必须填写。
	MongosCpu *uint64 `json:"MongosCpu,omitnil,omitempty" name:"MongosCpu"`

	// Mongos 内存大小。
	// -  购买分片集群时，必须填写。
	// - 单位：GB，支持1核2GB、2核4GB、4核8GB、8核16GB、16核32GB。
	MongosMemory *uint64 `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// Mongos 数量。购买分片集群时，必须填写。
	// - 单可用区部署实例，其数量范围为[3,32]。
	// - 多可用区部署实例，其数量范围为[6,32]。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil,omitempty" name:"MongosNodeNum"`

	// 只读节点数量，取值范围[0,5]。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil,omitempty" name:"ReadonlyNodeNum"`

	// 指只读节点所属可用区数组。跨可用区部署实例，参数**ReadonlyNodeNum**不为**0**时，必须配置该参数。
	ReadonlyNodeAvailabilityZoneList []*string `json:"ReadonlyNodeAvailabilityZoneList,omitnil,omitempty" name:"ReadonlyNodeAvailabilityZoneList"`

	// Hidden节点所属可用区。跨可用区部署实例，必须配置该参数。
	HiddenZone *string `json:"HiddenZone,omitnil,omitempty" name:"HiddenZone"`

	// 参数模板 ID。
	// - 参数模板是预置了特定参数值的集合，可用于快速配置新的 MongoDB 实例。合理使用参数模板，能有效提升数据库的部署效率与运行性能。
	// - 参数模板 ID 可通过 [DescribeDBInstanceParamTpl ](https://cloud.tencent.com/document/product/240/109155)接口获取。请选择与实例版本与架构所对应的参数模板 ID。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "MongoVersion")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "Period")
	delete(f, "MachineCode")
	delete(f, "ClusterType")
	delete(f, "ReplicateSetNum")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "Tags")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "Clone")
	delete(f, "Father")
	delete(f, "SecurityGroup")
	delete(f, "RestoreTime")
	delete(f, "InstanceName")
	delete(f, "AvailabilityZoneList")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNodeNum")
	delete(f, "ReadonlyNodeNum")
	delete(f, "ReadonlyNodeAvailabilityZoneList")
	delete(f, "HiddenZone")
	delete(f, "ParamTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// 订单ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 创建的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogDownloadTaskRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 节点名称
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 日志类别
	LogComponents []*string `json:"LogComponents,omitnil,omitempty" name:"LogComponents"`

	// 日志等级
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 日志ID
	LogIds []*string `json:"LogIds,omitnil,omitempty" name:"LogIds"`

	// 日志连接信息
	LogConnections []*string `json:"LogConnections,omitnil,omitempty" name:"LogConnections"`

	// 日志详情过滤字段
	LogDetailParams []*string `json:"LogDetailParams,omitnil,omitempty" name:"LogDetailParams"`
}

type CreateLogDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 节点名称
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 日志类别
	LogComponents []*string `json:"LogComponents,omitnil,omitempty" name:"LogComponents"`

	// 日志等级
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 日志ID
	LogIds []*string `json:"LogIds,omitnil,omitempty" name:"LogIds"`

	// 日志连接信息
	LogConnections []*string `json:"LogConnections,omitnil,omitempty" name:"LogConnections"`

	// 日志详情过滤字段
	LogDetailParams []*string `json:"LogDetailParams,omitnil,omitempty" name:"LogDetailParams"`
}

func (r *CreateLogDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "NodeNames")
	delete(f, "LogComponents")
	delete(f, "LogLevels")
	delete(f, "LogIds")
	delete(f, "LogConnections")
	delete(f, "LogDetailParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogDownloadTaskResponseParams struct {
	// 任务状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLogDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogDownloadTaskResponseParams `json:"Response"`
}

func (r *CreateLogDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CurrentOp struct {
	// 操作序号。
	OpId *int64 `json:"OpId,omitnil,omitempty" name:"OpId"`

	// 操作所在的命名空间，形式如db.collection。
	Ns *string `json:"Ns,omitnil,omitempty" name:"Ns"`

	// 操作执行语句。
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 操作类型。
	// - none：特殊状态，空闲连接或内部任务等。
	// - update：更新数据。
	// - insert：插入操作。
	// - query：查询操作。
	// - command：命令操作。
	// - getmore：获取更多数据。
	// - remove：删除操作。
	// - killcursors：释放查询游标的操作。
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// 操作所在的分片名称。
	ReplicaSetName *string `json:"ReplicaSetName,omitnil,omitempty" name:"ReplicaSetName"`

	// 操作所在的节点名称。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 操作详细信息。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 节点角色。
	// - primary：主节点。
	// - secondary：从节点。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 操作已执行时间（ms）。
	MicrosecsRunning *uint64 `json:"MicrosecsRunning,omitnil,omitempty" name:"MicrosecsRunning"`

	// 当前操作所在节点信息。
	ExecNode *string `json:"ExecNode,omitnil,omitempty" name:"ExecNode"`
}

type DBInstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type DBInstancePrice struct {
	// 实例单价。单位：元。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 实例原价。单位：元。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 实例折扣价。单位：元。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`
}

type DbURL struct {
	// 指 URI 类别，包括：，
	// - CLUSTER_ALL：指通过该 URI 连接库实例的主节点，可读写。
	// - CLUSTER_READ_READONLY：指通过该 URI 连接实例只读节点。
	// - CLUSTER_READ_SECONDARY：指通过该 URI 连接实例从节点。
	// - CLUSTER_READ_SECONDARY_AND_READONLY：指通过该 URI 连接实例只读从节点。
	// - CLUSTER_PRIMARY_AND_SECONDARY：指通过该 URI 连接实例 主节点与从节点。
	// - MONGOS_ALL：指通过该  URI 连接每个 Mongos 节点，可读写。
	// - MONGOS_READ_READONLY：指通过该 URI 连接 Mongos 的只读节点。
	// - MONGOS_READ_SECONDARY：指通过该 URI 连接 Mongos 的从节点。
	// - MONGOS_READ_PRIMARY_AND_SECONDARY：指通过该URI 连接 Mongos 的主节点与从节点。
	// - MONGOS_READ_SECONDARY_AND_READONLY：指通过该URI 连接 Mongos 的从节点与只读节点。
	URLType *string `json:"URLType,omitnil,omitempty" name:"URLType"`

	// 实例 URI 形式的连接串访问地址示例。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

// Predefined struct for user
type DeleteAccountUserRequestParams struct {
	// 指定待删除账号的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置待删除的账号名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 配置 mongouser 对应的密码。mongouser为系统默认账号，输入其对应的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil,omitempty" name:"MongoUserPassword"`
}

type DeleteAccountUserRequest struct {
	*tchttp.BaseRequest
	
	// 指定待删除账号的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置待删除的账号名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 配置 mongouser 对应的密码。mongouser为系统默认账号，输入其对应的密码。
	MongoUserPassword *string `json:"MongoUserPassword,omitnil,omitempty" name:"MongoUserPassword"`
}

func (r *DeleteAccountUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "MongoUserPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountUserResponseParams struct {
	// 账户删除任务ID。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountUserResponseParams `json:"Response"`
}

func (r *DeleteAccountUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileRequestParams struct {
	// 实例ID，格式如：cmgo-test1234，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志文件名称，须保证文件名的准确性。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-test1234，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志文件名称，须保证文件名的准确性。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

func (r *DeleteAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditLogFileResponseParams `json:"Response"`
}

func (r *DeleteAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogDownloadTaskRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteLogDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteLogDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogDownloadTaskResponseParams struct {
	// 任务状态，0:成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogDownloadTaskResponseParams `json:"Response"`
}

func (r *DeleteLogDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliverSummary struct {
	// 投递类型，store（存储类），mq（消息通道）
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 投递子类型：cls，ckafka。
	DeliverSubType *string `json:"DeliverSubType,omitnil,omitempty" name:"DeliverSubType"`
}

// Predefined struct for user
type DescribeAccountUsersRequestParams struct {
	// 指定待获取账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAccountUsersRequest struct {
	*tchttp.BaseRequest
	
	// 指定待获取账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAccountUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountUsersResponseParams struct {
	// 实例账号列表。
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountUsersResponseParams `json:"Response"`
}

func (r *DescribeAccountUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoRequestParams struct {
	// 指定需查询的异步请求 ID。当接口操作涉及异步流程时（如 [CreateBackupDBInstance](https://cloud.tencent.com/document/product/240/46599)），其返回值中的 AsyncRequestId 即为本参数所需填入的 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest
	
	// 指定需查询的异步请求 ID。当接口操作涉及异步流程时（如 [CreateBackupDBInstance](https://cloud.tencent.com/document/product/240/46599)），其返回值中的 AsyncRequestId 即为本参数所需填入的 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncRequestInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncRequestInfoResponseParams struct {
	// 状态。返回参数有：initial-初始化、running-运行中、paused-任务执行失败，已暂停、undoed-任务执行失败，已回滚、failed-任务执行失败, 已终止、success-成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务执行结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncRequestInfoResponseParams `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListRequestParams struct {
	// 指明待查询的实例为已开通审计或未开通审计。<ul><li>1：已开通审计功能。</li><li>0：未开通审计功能。</li></ul>
	AuditSwitch *uint64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// 筛选条件。
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 审计类型，不传 默认全部，0 全审计，1 规则审计
	AuditMode *uint64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// 每页显示数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAuditInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 指明待查询的实例为已开通审计或未开通审计。<ul><li>1：已开通审计功能。</li><li>0：未开通审计功能。</li></ul>
	AuditSwitch *uint64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// 筛选条件。
	Filters []*Filters `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 审计类型，不传 默认全部，0 全审计，1 规则审计
	AuditMode *uint64 `json:"AuditMode,omitnil,omitempty" name:"AuditMode"`

	// 每页显示数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAuditInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditSwitch")
	delete(f, "Filters")
	delete(f, "AuditMode")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListResponseParams struct {
	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审计实例详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuditInstance `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAuditInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadTaskRequestParams struct {
	// 实例ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定备份文件名，用于过滤指定文件的下载任务。请通过接口 [DescribeDBBackups](https://cloud.tencent.com/document/product/240/38574) 获取备份文件名。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 指定查询时间范围内的任务，StartTime 指定开始时间。若不指定开始时间，则默认不限制开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 指定查询时间范围内的任务，EndTime 指定截止时间。若不指定截止时间，则默认不限制截止时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 此次查询返回的条数，取值范围为1-100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定此次查询返回的页数，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段。
	// - createTime：按照备份下载任务的创建时间排序。默认为 createTime。
	// - finishTime：按照备份下载任务的完成时间排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式。
	// - asc：升序排列。
	// - desc：降序排列。默认为 desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 指定任务状态，用于过滤下载任务。若不配置该参数，则返回所有状态类型的任务。
	// - 0：等待执行。
	// - 1：正在下载。
	// - 2：下载完成。
	// - 3：下载失败。
	// - 4：等待重试。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeBackupDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定备份文件名，用于过滤指定文件的下载任务。请通过接口 [DescribeDBBackups](https://cloud.tencent.com/document/product/240/38574) 获取备份文件名。
	BackupName *string `json:"BackupName,omitnil,omitempty" name:"BackupName"`

	// 指定查询时间范围内的任务，StartTime 指定开始时间。若不指定开始时间，则默认不限制开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 指定查询时间范围内的任务，EndTime 指定截止时间。若不指定截止时间，则默认不限制截止时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 此次查询返回的条数，取值范围为1-100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定此次查询返回的页数，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段。
	// - createTime：按照备份下载任务的创建时间排序。默认为 createTime。
	// - finishTime：按照备份下载任务的完成时间排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式。
	// - asc：升序排列。
	// - desc：降序排列。默认为 desc。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 指定任务状态，用于过滤下载任务。若不配置该参数，则返回所有状态类型的任务。
	// - 0：等待执行。
	// - 1：正在下载。
	// - 2：下载完成。
	// - 3：下载失败。
	// - 4：等待重试。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeBackupDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadTaskResponseParams struct {
	// 满足查询条件的所有条数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 下载任务列表。
	Tasks []*BackupDownloadTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadTaskResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupRulesRequestParams struct {
	// 指定实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackupRulesRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupRulesResponseParams struct {
	// 备份数据保留期限。单位为：天。
	BackupSaveTime *uint64 `json:"BackupSaveTime,omitnil,omitempty" name:"BackupSaveTime"`

	// 自动备份开始时间。
	BackupTime *uint64 `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// 备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	BackupMethod *uint64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupRulesResponseParams `json:"Response"`
}

func (r *DescribeBackupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsRequestParams struct {
	// 指定待查询的实例ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 单次请求返回的数量。最小值为1，最大值为1000，默认值为1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0。Offset=Limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeClientConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 指定待查询的实例ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 单次请求返回的数量。最小值为1，最大值为1000，默认值为1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0。Offset=Limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeClientConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsResponseParams struct {
	// 客户端连接信息，包括客户端 IP 和对应 IP 的连接数量。
	Clients []*ClientConnection `json:"Clients,omitnil,omitempty" name:"Clients"`

	// 满足条件的记录总条数，可用于分页查询。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientConnectionsResponseParams `json:"Response"`
}

func (r *DescribeClientConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentOpRequestParams struct {
	// 指定要查询的实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 操作所属的命名空间 namespace，格式为 db.collection。
	Ns *string `json:"Ns,omitnil,omitempty" name:"Ns"`

	// 设置查询筛选条件为操作任务已经执行的时间。
	// - 默认值为0，取值范围为[0, 3600000]，单位：毫秒。
	// - 结果将返回超过设置时间的操作。
	MillisecondRunning *uint64 `json:"MillisecondRunning,omitnil,omitempty" name:"MillisecondRunning"`

	// 设置查询筛选条件为操作任务类型。取值包括：
	// - none：特殊状态，空闲连接或内部任务等。
	// - update：更新数据。
	// - insert：插入操作。
	// - query：查询操作。
	// - command：命令操作。
	// - getmore：获取更多数据。
	// - remove：删除操作。
	// - killcursors：释放查询游标的操作。
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// 筛选条件，分片名称。
	ReplicaSetName *string `json:"ReplicaSetName,omitnil,omitempty" name:"ReplicaSetName"`

	// 设置查询筛选条件为节点角色。
	// - primary：主节点。
	// - secondary：从节点。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 单次请求返回的数量，默认值为100，取值范围为[0,100]。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0，取值范围为[0,10000]。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回结果集排序的字段，目前支持按照 MicrosecsRunning（操作任务已执行的时间）排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果集排序方式。
	// - ASC：升序。默认为 ASC，按照升序排序。
	// - DESC：降序。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeCurrentOpRequest struct {
	*tchttp.BaseRequest
	
	// 指定要查询的实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 操作所属的命名空间 namespace，格式为 db.collection。
	Ns *string `json:"Ns,omitnil,omitempty" name:"Ns"`

	// 设置查询筛选条件为操作任务已经执行的时间。
	// - 默认值为0，取值范围为[0, 3600000]，单位：毫秒。
	// - 结果将返回超过设置时间的操作。
	MillisecondRunning *uint64 `json:"MillisecondRunning,omitnil,omitempty" name:"MillisecondRunning"`

	// 设置查询筛选条件为操作任务类型。取值包括：
	// - none：特殊状态，空闲连接或内部任务等。
	// - update：更新数据。
	// - insert：插入操作。
	// - query：查询操作。
	// - command：命令操作。
	// - getmore：获取更多数据。
	// - remove：删除操作。
	// - killcursors：释放查询游标的操作。
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// 筛选条件，分片名称。
	ReplicaSetName *string `json:"ReplicaSetName,omitnil,omitempty" name:"ReplicaSetName"`

	// 设置查询筛选条件为节点角色。
	// - primary：主节点。
	// - secondary：从节点。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 单次请求返回的数量，默认值为100，取值范围为[0,100]。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0，取值范围为[0,10000]。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回结果集排序的字段，目前支持按照 MicrosecsRunning（操作任务已执行的时间）排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果集排序方式。
	// - ASC：升序。默认为 ASC，按照升序排序。
	// - DESC：降序。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeCurrentOpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentOpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ns")
	delete(f, "MillisecondRunning")
	delete(f, "Op")
	delete(f, "ReplicaSetName")
	delete(f, "State")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurrentOpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurrentOpResponseParams struct {
	// 符合查询条件的操作总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 当前操作列表。
	CurrentOps []*CurrentOp `json:"CurrentOps,omitnil,omitempty" name:"CurrentOps"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCurrentOpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurrentOpResponseParams `json:"Response"`
}

func (r *DescribeCurrentOpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurrentOpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明**:
	// 1. 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// 2. 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 分页大小，最大值为100，不设置默认查询所有。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，最小值为0，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明**:
	// 1. 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// 2. 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *int64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 分页大小，最大值为100，不设置默认查询所有。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，最小值为0，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsResponseParams struct {
	// 备份列表。
	BackupList []*BackupInfo `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// 备份总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBBackupsResponseParams `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDealRequestParams struct {
	// 订单 ID。
	// - 按量计费实例，请通过 [CreateDBInstanceHour](https://cloud.tencent.com/document/product/240/38570) 接口输出的参数**DealId**获取。。
	// - 包年包月计费实例，请通过 [CreateDBInstance](https://cloud.tencent.com/document/product/240/38571) 接口输出的参数**DealId**获取。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`
}

type DescribeDBInstanceDealRequest struct {
	*tchttp.BaseRequest
	
	// 订单 ID。
	// - 按量计费实例，请通过 [CreateDBInstanceHour](https://cloud.tencent.com/document/product/240/38570) 接口输出的参数**DealId**获取。。
	// - 包年包月计费实例，请通过 [CreateDBInstance](https://cloud.tencent.com/document/product/240/38571) 接口输出的参数**DealId**获取。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`
}

func (r *DescribeDBInstanceDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceDealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDealResponseParams struct {
	// 订单状态。
	// - 1：未支付。
	// - 2：已支付。
	// - 3：发货中。
	// - 4：发货成功。
	// - 5：发货失败。
	// - 6：退款。
	// - 7：订单关闭。
	// - 8：超时未支付关闭。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 订单原价。单位：元。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 订单折扣价格。单位：元。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 订单操作行为。
	// - purchase：新购。
	// - renew：续费。
	// - upgrade：升配.
	// - downgrade：降配.
	// - refund：退货退款。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 当前订单的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceDealResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceDealResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceNamespaceRequestParams struct {
	// 指定查询数据库所属的实例 ID，支持批量查询。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定查询的数据库名。为空时，返回当前实例的全部数据库列表。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DescribeDBInstanceNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 指定查询数据库所属的实例 ID，支持批量查询。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定查询的数据库名。为空时，返回当前实例的全部数据库列表。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

func (r *DescribeDBInstanceNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceNamespaceResponseParams struct {
	// 查询实例的数据库列表。若未使用 DbName 指定具体查询的数据库，则仅返回查询实例所有的数据库列表，而不返回 Collections 集合信息。
	Databases []*string `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 查询的集合信息。指定 DbName 时，则仅返回该数据库下的集合列表。
	Collections []*string `json:"Collections,omitnil,omitempty" name:"Collections"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceNamespaceResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceNodePropertyRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点 ID。请登录 [MongoDB 控制台的节点管理](https://console.cloud.tencent.com/mongodb)复制节点 ID。
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// 节点角色。可选值包括：
	// - PRIMARY：主节点。
	// - SECONDARY：从节点。
	// - READONLY：只读节点。
	// - ARBITER：仲裁节点。
	Roles []*string `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 该参数指定节点是否为 Hidden 节点，默认为 false。
	OnlyHidden *bool `json:"OnlyHidden,omitnil,omitempty" name:"OnlyHidden"`

	// 该参数指定选举新主节点的优先级。其取值范围为[0,100]，数值越高，优先级越高。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 该参数指定节点投票权。
	// - 1：具有投票权。
	// - 0：无投票权。
	Votes *int64 `json:"Votes,omitnil,omitempty" name:"Votes"`

	// 节点标签。
	Tags []*NodeTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeDBInstanceNodePropertyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点 ID。请登录 [MongoDB 控制台的节点管理](https://console.cloud.tencent.com/mongodb)复制节点 ID。
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`

	// 节点角色。可选值包括：
	// - PRIMARY：主节点。
	// - SECONDARY：从节点。
	// - READONLY：只读节点。
	// - ARBITER：仲裁节点。
	Roles []*string `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 该参数指定节点是否为 Hidden 节点，默认为 false。
	OnlyHidden *bool `json:"OnlyHidden,omitnil,omitempty" name:"OnlyHidden"`

	// 该参数指定选举新主节点的优先级。其取值范围为[0,100]，数值越高，优先级越高。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 该参数指定节点投票权。
	// - 1：具有投票权。
	// - 0：无投票权。
	Votes *int64 `json:"Votes,omitnil,omitempty" name:"Votes"`

	// 节点标签。
	Tags []*NodeTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeDBInstanceNodePropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceNodePropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeIds")
	delete(f, "Roles")
	delete(f, "OnlyHidden")
	delete(f, "Priority")
	delete(f, "Votes")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceNodePropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceNodePropertyResponseParams struct {
	// Mongos节点属性。
	Mongos []*NodeProperty `json:"Mongos,omitnil,omitempty" name:"Mongos"`

	// 副本集节点信息。
	ReplicateSets []*ReplicateSetInfo `json:"ReplicateSets,omitnil,omitempty" name:"ReplicateSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceNodePropertyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceNodePropertyResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceNodePropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceNodePropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParamTplDetailRequestParams struct {
	// 参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`

	// 参数名称，传入该值，则只会获取该字段的参数详情。为空时，返回全部参数。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
}

type DescribeDBInstanceParamTplDetailRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`

	// 参数名称，传入该值，则只会获取该字段的参数详情。为空时，返回全部参数。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
}

func (r *DescribeDBInstanceParamTplDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParamTplDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TplId")
	delete(f, "ParamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceParamTplDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParamTplDetailResponseParams struct {
	// 枚举类参数详情列表。
	InstanceEnumParams []*InstanceEnumParam `json:"InstanceEnumParams,omitnil,omitempty" name:"InstanceEnumParams"`

	// 整形参数详情列表。
	InstanceIntegerParams []*InstanceIntegerParam `json:"InstanceIntegerParams,omitnil,omitempty" name:"InstanceIntegerParams"`

	// 文本参数详情列表。
	InstanceTextParams []*InstanceTextParam `json:"InstanceTextParams,omitnil,omitempty" name:"InstanceTextParams"`

	// 多值参数详情列表。
	InstanceMultiParams []*InstanceMultiParam `json:"InstanceMultiParams,omitnil,omitempty" name:"InstanceMultiParams"`

	// 参数总个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 模板适配的实例版本。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本，
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 模板适配集群类型。
	// - REPLSET：副本集实例。
	// - SHARD：分片实例。
	// - STANDALONE：单节点实例。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 参数模板名称。
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceParamTplDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceParamTplDetailResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceParamTplDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParamTplDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParamTplRequestParams struct {
	// 参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplIds []*string `json:"TplIds,omitnil,omitempty" name:"TplIds"`

	// 指定查询的模板名称。
	TplNames []*string `json:"TplNames,omitnil,omitempty" name:"TplNames"`

	// 指定所需查询的参数模板的数据库版本号。具体支持的版本信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本，
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion []*string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 指定查询的模板类型。
	// - DEFAULT：系统默认模板。
	// - CUSTOMIZE：自定义模板。
	TplType *string `json:"TplType,omitnil,omitempty" name:"TplType"`
}

type DescribeDBInstanceParamTplRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplIds []*string `json:"TplIds,omitnil,omitempty" name:"TplIds"`

	// 指定查询的模板名称。
	TplNames []*string `json:"TplNames,omitnil,omitempty" name:"TplNames"`

	// 指定所需查询的参数模板的数据库版本号。具体支持的版本信息，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本，
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion []*string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 指定查询的模板类型。
	// - DEFAULT：系统默认模板。
	// - CUSTOMIZE：自定义模板。
	TplType *string `json:"TplType,omitnil,omitempty" name:"TplType"`
}

func (r *DescribeDBInstanceParamTplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParamTplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TplIds")
	delete(f, "TplNames")
	delete(f, "MongoVersion")
	delete(f, "TplType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceParamTplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParamTplResponseParams struct {
	// 参数模板列表信息。
	ParamTpls []*ParamTpl `json:"ParamTpls,omitnil,omitempty" name:"ParamTpls"`

	// 参数模板总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceParamTplResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceParamTplResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceParamTplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParamTplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceURLRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBInstanceURLRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBInstanceURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceURLResponseParams struct {
	// 实例 URI 形式的连接串访问地址示例。包含：URI 类型及连接串地址。
	Urls []*DbURL `json:"Urls,omitnil,omitempty" name:"Urls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceURLResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 实例 ID 列表。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 指定查询的实例类型。取值范围如下：
	// - 0：所有实例。
	// - 1：正式实例。
	// - 2：临时实例
	// - 3：只读实例。
	// - -1：查询同时包括正式实例、只读实例与灾备实例。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 指定所查询实例的集群类型，取值范围如下：
	// - 0：副本集实例。
	// - 1：分片实例。
	// - -1：副本集与分片实例。
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 指定所查询实例的当前状态，取值范围如下所示：
	// - 0：待初始化。
	// - 1：流程处理中，例如：变更规格、参数修改等。
	// - 2：实例正常运行中。
	// - -2：已隔离（包年包月）。
	// - -3：已隔离（按量计费）。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有网络的 ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其子网 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定所查询实例的付费类型。
	// - 0：查询按量计费实例。
	// - 1：查询包年包月实例。
	// - -1：查询按量计费与包年包月实例。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 单次请求返回的数量。默认值为20，取值范围为(1,100]。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 配置返回结果排序依据的字段。目前支持依据"ProjectId"、"InstanceName"、"CreateTime"排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 配置返回结果的排序方式。
	// - ASC：升序。
	// - DESC：降序。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 项目 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在右上角的账户信息下拉菜单中，选择项目管理查询项目。
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 指定查询搜索的关键词。支持设置为具体的实例ID、实例名称或者内网 IP 地址。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 标签信息，包含标签键与标签值。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 指定查询的实例类型。取值范围如下：
	// - 0：所有实例。
	// - 1：正式实例。
	// - 2：临时实例
	// - 3：只读实例。
	// - -1：查询同时包括正式实例、只读实例与灾备实例。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 指定所查询实例的集群类型，取值范围如下：
	// - 0：副本集实例。
	// - 1：分片实例。
	// - -1：副本集与分片实例。
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 指定所查询实例的当前状态，取值范围如下所示：
	// - 0：待初始化。
	// - 1：流程处理中，例如：变更规格、参数修改等。
	// - 2：实例正常运行中。
	// - -2：已隔离（包年包月）。
	// - -3：已隔离（按量计费）。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有网络的 ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网ID。
	// - 基础网络则无需配置该参数。
	// - 请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表中，单击私有网络名称，在**私有网络**页面获取其子网 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定所查询实例的付费类型。
	// - 0：查询按量计费实例。
	// - 1：查询包年包月实例。
	// - -1：查询按量计费与包年包月实例。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 单次请求返回的数量。默认值为20，取值范围为(1,100]。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 配置返回结果排序依据的字段。目前支持依据"ProjectId"、"InstanceName"、"CreateTime"排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 配置返回结果的排序方式。
	// - ASC：升序。
	// - DESC：降序。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 项目 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在右上角的账户信息下拉菜单中，选择项目管理查询项目。
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 指定查询搜索的关键词。支持设置为具体的实例ID、实例名称或者内网 IP 地址。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 标签信息，包含标签键与标签值。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "ClusterType")
	delete(f, "Status")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "PayMode")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "ProjectIds")
	delete(f, "SearchKey")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 符合查询条件的实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	InstanceDetails []*InstanceDetail `json:"InstanceDetails,omitnil,omitempty" name:"InstanceDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetailedSlowLogsRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定查询慢日志的开始时间。- 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。- 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 指定查询慢日志的结束时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定慢日志查询阈值，即查询执行时间大于此值的慢日志。单位：ms，默认值：100。
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// 指定查询慢日志的命令类型。
	Commands []*string `json:"Commands,omitnil,omitempty" name:"Commands"`

	// 全文搜索关键字，多个关键字间为或关系。
	Texts []*string `json:"Texts,omitnil,omitempty" name:"Texts"`

	// 指定查询慢日志的节点名称。请通过接口 [DescribeDBInstanceNodeProperty](https://cloud.tencent.com/document/product/240/82022) 查询节点名称。
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 指定查询 queryHash 值。
	QueryHash []*string `json:"QueryHash,omitnil,omitempty" name:"QueryHash"`

	// 分页偏移量。默认值：0。取值范围：[0,100]。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回条数。默认值：20。取值范围：[0,10000]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定慢日志排序条件。
	// - StartTime：按照慢日志生成时间排序。
	// - ExecTime：按照慢日志执行时间排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 指定排序方式。
	// - desc：倒序。
	// - asc：顺序。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeDetailedSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定查询慢日志的开始时间。- 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。- 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 指定查询慢日志的结束时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定慢日志查询阈值，即查询执行时间大于此值的慢日志。单位：ms，默认值：100。
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// 指定查询慢日志的命令类型。
	Commands []*string `json:"Commands,omitnil,omitempty" name:"Commands"`

	// 全文搜索关键字，多个关键字间为或关系。
	Texts []*string `json:"Texts,omitnil,omitempty" name:"Texts"`

	// 指定查询慢日志的节点名称。请通过接口 [DescribeDBInstanceNodeProperty](https://cloud.tencent.com/document/product/240/82022) 查询节点名称。
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 指定查询 queryHash 值。
	QueryHash []*string `json:"QueryHash,omitnil,omitempty" name:"QueryHash"`

	// 分页偏移量。默认值：0。取值范围：[0,100]。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回条数。默认值：20。取值范围：[0,10000]
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 指定慢日志排序条件。
	// - StartTime：按照慢日志生成时间排序。
	// - ExecTime：按照慢日志执行时间排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 指定排序方式。
	// - desc：倒序。
	// - asc：顺序。
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeDetailedSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailedSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ExecTime")
	delete(f, "Commands")
	delete(f, "Texts")
	delete(f, "NodeNames")
	delete(f, "QueryHash")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDetailedSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDetailedSlowLogsResponseParams struct {
	// 满足条件的慢日志数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢日志详情。
	DetailedSlowLogs []*SlowLogItem `json:"DetailedSlowLogs,omitnil,omitempty" name:"DetailedSlowLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDetailedSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDetailedSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeDetailedSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDetailedSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// 指定待查询参数列表的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 指定待查询参数列表的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// 参数值为枚举类型的参数集合。
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitnil,omitempty" name:"InstanceEnumParam"`

	// 参数值为 Integer 类型的参数集合。
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitnil,omitempty" name:"InstanceIntegerParam"`

	// 参数值为 Text 类型的参数集合。
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitnil,omitempty" name:"InstanceTextParam"`

	// 参数值为混合类型的参数集合。
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitnil,omitempty" name:"InstanceMultiParam"`

	// 当前实例支持修改的参数数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSSLRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceSSLRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSSLResponseParams struct {
	// SSL开启状态。0为关闭，1为开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书过期时间，格式为2023-05-01 12:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 证书下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertUrl *string `json:"CertUrl,omitnil,omitempty" name:"CertUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSSLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSSLResponseParams `json:"Response"`
}

func (r *DescribeInstanceSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogDownloadTasksRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 下载任务的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 下载任务的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeLogDownloadTasksRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页码
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 下载任务的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 下载任务的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeLogDownloadTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogDownloadTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogDownloadTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogDownloadTasksResponseParams struct {
	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogDownloadTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogDownloadTasksResponseParams `json:"Response"`
}

func (r *DescribeLogDownloadTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogDownloadTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMongodbLogsRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询日志的开启时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询时间范围：仅支持查询最近 7 天内的日志数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询日志的结束时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询时间范围：仅支持查询最近 7 天内的日志数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 节点 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)的**节点管理**页面获取查询的节点 ID。
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 日志类别。
	// - 日志类别包括但不限于 COMMAND、ACCESS、CONTROL、FTDC、INDEX、NETWORK、QUERY、REPL、SHARDING、STORAGE、RECOVERY、JOURNAL 和 WRITE 等。具体支持的类别可能会因 MongoDB 的版本而存在差异。具体信息，请参见[日志消息](https://www.mongodb.com/zh-cn/docs/v5.0/reference/log-messages/#log-message-examples)。
	// - 登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，也可查看**日志类别**。
	LogComponents []*string `json:"LogComponents,omitnil,omitempty" name:"LogComponents"`

	// 日志级别。
	// - 日志级别按严重性从高到低依次为：FATAL、ERROR、WARNING。
	// - 登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，可查看**日志级别**。
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 日志 ID。登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，可查看**日志 ID**。
	LogIds []*string `json:"LogIds,omitnil,omitempty" name:"LogIds"`

	// 日志连接信息。登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，可查看**日志连接信息**。
	LogConnections []*string `json:"LogConnections,omitnil,omitempty" name:"LogConnections"`

	// 指定日志筛选的字段。
	LogDetailParams []*string `json:"LogDetailParams,omitnil,omitempty" name:"LogDetailParams"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMongodbLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询日志的开启时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询时间范围：仅支持查询最近 7 天内的日志数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询日志的结束时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询时间范围：仅支持查询最近 7 天内的日志数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 节点 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)的**节点管理**页面获取查询的节点 ID。
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 日志类别。
	// - 日志类别包括但不限于 COMMAND、ACCESS、CONTROL、FTDC、INDEX、NETWORK、QUERY、REPL、SHARDING、STORAGE、RECOVERY、JOURNAL 和 WRITE 等。具体支持的类别可能会因 MongoDB 的版本而存在差异。具体信息，请参见[日志消息](https://www.mongodb.com/zh-cn/docs/v5.0/reference/log-messages/#log-message-examples)。
	// - 登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，也可查看**日志类别**。
	LogComponents []*string `json:"LogComponents,omitnil,omitempty" name:"LogComponents"`

	// 日志级别。
	// - 日志级别按严重性从高到低依次为：FATAL、ERROR、WARNING。
	// - 登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，可查看**日志级别**。
	LogLevels []*string `json:"LogLevels,omitnil,omitempty" name:"LogLevels"`

	// 日志 ID。登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，可查看**日志 ID**。
	LogIds []*string `json:"LogIds,omitnil,omitempty" name:"LogIds"`

	// 日志连接信息。登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在**日志管理**页面的**错误日志**页签，可查看**日志连接信息**。
	LogConnections []*string `json:"LogConnections,omitnil,omitempty" name:"LogConnections"`

	// 指定日志筛选的字段。
	LogDetailParams []*string `json:"LogDetailParams,omitnil,omitempty" name:"LogDetailParams"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMongodbLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMongodbLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "NodeNames")
	delete(f, "LogComponents")
	delete(f, "LogLevels")
	delete(f, "LogIds")
	delete(f, "LogConnections")
	delete(f, "LogDetailParams")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMongodbLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMongodbLogsResponseParams struct {
	// 日志总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogList []*LogInfo `json:"LogList,omitnil,omitempty" name:"LogList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMongodbLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMongodbLogsResponseParams `json:"Response"`
}

func (r *DescribeMongodbLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMongodbLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupResponseParams struct {
	// 实例绑定的安全组信息。
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogPatternsRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢日志起始时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 慢日志终止时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil,omitempty" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 慢日志返回格式，可设置为json，不传默认返回原生慢日志格式。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type DescribeSlowLogPatternsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢日志起始时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 慢日志终止时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil,omitempty" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 慢日志返回格式，可设置为json，不传默认返回原生慢日志格式。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

func (r *DescribeSlowLogPatternsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogPatternsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SlowMS")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogPatternsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogPatternsResponseParams struct {
	// 慢日志统计信息总数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 慢日志统计信息
	SlowLogPatterns []*SlowLogPattern `json:"SlowLogPatterns,omitnil,omitempty" name:"SlowLogPatterns"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogPatternsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogPatternsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogPatternsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogPatternsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// 实例ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢日志起始时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 慢日志终止时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil,omitempty" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 慢日志返回格式。默认返回原生慢日志格式，4.4及以上版本可设置为json。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢日志起始时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 慢日志终止时间。
	// - 格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。
	// - 查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil,omitempty" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 慢日志返回格式。默认返回原生慢日志格式，4.4及以上版本可设置为json。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SlowMS")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// 慢日志总数。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 慢日志详情。
	SlowLogs []*string `json:"SlowLogs,omitnil,omitempty" name:"SlowLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecInfoRequestParams struct {
	// 待查询可用区。当前支持的可用区，请参见[地域与可用区](https://cloud.tencent.com/document/product/240/3637)。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DescribeSpecInfoRequest struct {
	*tchttp.BaseRequest
	
	// 待查询可用区。当前支持的可用区，请参见[地域与可用区](https://cloud.tencent.com/document/product/240/3637)。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *DescribeSpecInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecInfoResponseParams struct {
	// 实例售卖规格信息列表。
	SpecInfoList []*SpecificationInfo `json:"SpecInfoList,omitnil,omitempty" name:"SpecInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpecInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecInfoResponseParams `json:"Response"`
}

func (r *DescribeSpecInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTransparentDataEncryptionStatusRequestParams struct {
	// 指定实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeTransparentDataEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeTransparentDataEncryptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransparentDataEncryptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTransparentDataEncryptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTransparentDataEncryptionStatusResponseParams struct {
	// 表示是否开启了透明加密。 
	// - close：未开启。
	// - open：已开启。
	TransparentDataEncryptionStatus *string `json:"TransparentDataEncryptionStatus,omitnil,omitempty" name:"TransparentDataEncryptionStatus"`

	// 已绑定的密钥列表，如未绑定，返回null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyInfoList []*KMSInfoDetail `json:"KeyInfoList,omitnil,omitempty" name:"KeyInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTransparentDataEncryptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTransparentDataEncryptionStatusResponseParams `json:"Response"`
}

func (r *DescribeTransparentDataEncryptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransparentDataEncryptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDBInstanceParamTplRequestParams struct {
	// 参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`
}

type DropDBInstanceParamTplRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`
}

func (r *DropDBInstanceParamTplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDBInstanceParamTplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TplId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropDBInstanceParamTplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDBInstanceParamTplResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DropDBInstanceParamTplResponse struct {
	*tchttp.BaseResponse
	Response *DropDBInstanceParamTplResponseParams `json:"Response"`
}

func (r *DropDBInstanceParamTplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDBInstanceParamTplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTransparentDataEncryptionRequestParams struct {
	// 实例 ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。目前支持通用版本包含：4.4、5.0，云盘版暂不支持。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	//  [密钥管理系统（Key Management Service，KMS）](https://cloud.tencent.com/document/product/573/18809)服务所在地域，例如 ap-shanghai。
	KmsRegion *string `json:"KmsRegion,omitnil,omitempty" name:"KmsRegion"`

	// 密钥 ID。若不设置该参数，不指定具体的密钥 ID，由腾讯云自动生成密钥。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type EnableTransparentDataEncryptionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。目前支持通用版本包含：4.4、5.0，云盘版暂不支持。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	//  [密钥管理系统（Key Management Service，KMS）](https://cloud.tencent.com/document/product/573/18809)服务所在地域，例如 ap-shanghai。
	KmsRegion *string `json:"KmsRegion,omitnil,omitempty" name:"KmsRegion"`

	// 密钥 ID。若不设置该参数，不指定具体的密钥 ID，由腾讯云自动生成密钥。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *EnableTransparentDataEncryptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTransparentDataEncryptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KmsRegion")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableTransparentDataEncryptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableTransparentDataEncryptionResponseParams struct {
	// 开启透明加密的异步流程id，用于查询流程状态。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableTransparentDataEncryptionResponse struct {
	*tchttp.BaseResponse
	Response *EnableTransparentDataEncryptionResponseParams `json:"Response"`
}

func (r *EnableTransparentDataEncryptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableTransparentDataEncryptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FBKeyValue struct {
	// 指定按 Key 闪回的目标 Key （键） 。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 指定按 Key 闪回的目标 Key 所对应的 Value（值）。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Filters struct {
	// 搜索字段，目前支持：
	// "InstanceId"：实例Id，例如：cmgo-****）
	// "InstanceName"：实例名称
	// "ClusterId"：实例组Id，例如：cmgo-****
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 筛选值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type FlashBackDBInstanceRequestParams struct {
	// 开启按 Key 回档的实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制需开启按 Key 回档的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定数据回档的具体时间点，即将数据恢复到指定时间点的状态。
	TargetFlashbackTime *string `json:"TargetFlashbackTime,omitnil,omitempty" name:"TargetFlashbackTime"`

	// 指定回档数据的目标库表。
	TargetDatabases []*FlashbackDatabase `json:"TargetDatabases,omitnil,omitempty" name:"TargetDatabases"`

	// 数据回档的目标实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制目标实例 ID。
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`
}

type FlashBackDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 开启按 Key 回档的实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制需开启按 Key 回档的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定数据回档的具体时间点，即将数据恢复到指定时间点的状态。
	TargetFlashbackTime *string `json:"TargetFlashbackTime,omitnil,omitempty" name:"TargetFlashbackTime"`

	// 指定回档数据的目标库表。
	TargetDatabases []*FlashbackDatabase `json:"TargetDatabases,omitnil,omitempty" name:"TargetDatabases"`

	// 数据回档的目标实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制目标实例 ID。
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`
}

func (r *FlashBackDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlashBackDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TargetFlashbackTime")
	delete(f, "TargetDatabases")
	delete(f, "TargetInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlashBackDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlashBackDBInstanceResponseParams struct {
	// 回档数据异步任务 ID。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FlashBackDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *FlashBackDBInstanceResponseParams `json:"Response"`
}

func (r *FlashBackDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlashBackDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FlashbackCollection struct {
	// 指定按 Key 闪回源数据库集合名。
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 指定按 Key 闪回目标数据库集合名。
	TargetResultCollectionName *string `json:"TargetResultCollectionName,omitnil,omitempty" name:"TargetResultCollectionName"`

	// 指定用于过滤按 Key 闪回的 Key（键）。
	FilterKey *string `json:"FilterKey,omitnil,omitempty" name:"FilterKey"`

	// 指定用于按 Key 闪回的键值对。数组元素最大限制为 50000。
	KeyValues []*FBKeyValue `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

type FlashbackDatabase struct {
	// 按 Key 闪回目标数据所在库。
	DBName *string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// 按 Key 闪回的数据库集合。
	Collections []*FlashbackCollection `json:"Collections,omitnil,omitempty" name:"Collections"`
}

// Predefined struct for user
type FlushInstanceRouterConfigRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type FlushInstanceRouterConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *FlushInstanceRouterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushInstanceRouterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlushInstanceRouterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlushInstanceRouterConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FlushInstanceRouterConfigResponse struct {
	*tchttp.BaseResponse
	Response *FlushInstanceRouterConfigResponseParams `json:"Response"`
}

func (r *FlushInstanceRouterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushInstanceRouterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDBInstancesRequestParams struct {
	// 实例所属区域及可用区信息。具体信息，请参见[地域和可用区](https://cloud.tencent.com/document/product/240/3637)。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// - 创建副本集实例，指每个副本集内主从节点数量。每个副本集所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 创建分片集群实例，指每个分片的主从节点数量。每个分片所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例内存大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数CPU与Memory分别对应CPU核数与内存规格。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinStorage与MaxStorage分别对应其最小磁盘规格与最大磁盘规格。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例版本信息。具体支持的版本，请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MongoVersionCode为实例所支持的版本信息。版本信息与版本号对应关系如下：
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 产品规格类型。
	// - HIO10G：通用高HIO万兆型。
	// - HCD：云盘版。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例数量，取值范围为[1,10]。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// - 创建副本集实例，指副本集数量，该参数只能为1。
	// - 创建分片集群实例，指分片的数量。请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询分片数量的取值范围，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// - 选择包年包月计费模式，即 <b>InstanceChargeType </b>设定为<b>PREPAID</b>时，必须设置该参数，指定购买实例的购买时长。取值可选：[1,2,3,4,5,6,7,8,9,10,11,12,24,36]；单位：月。
	// -选择按量计费，即 <b>InstanceChargeType</b> 设定为 **POSTPAID_BY_HOUR** 时，该参数仅可配置为 1。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例付费方式。
	// - PREPAID：包年包月计费。
	// - POSTPAID_BY_HOUR：按量计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Mongos CPU 核数，支持1、2、4、8、16。购买分片集群时，必须填写。注意为空时取默认取值为2C。
	MongosCpu *uint64 `json:"MongosCpu,omitnil,omitempty" name:"MongosCpu"`

	// Mongos 内存大小。-  购买分片集群时，必须填写。- 单位：GB，支持1核2GB、2核4GB、4核8GB、8核16GB、16核32GB。注意为空时取默认取值为4G。
	MongosMemory *uint64 `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// 指 Mongos 个数，取值范围为[3,32]。若为分片集群实例询价，则该参数必须设置。注意为空时取默认取值为3个节点。
	MongosNum *uint64 `json:"MongosNum,omitnil,omitempty" name:"MongosNum"`

	// 指 ConfigServer CPU核数，固定取值为 1，单位：GB，可不配置该参数。
	ConfigServerCpu *uint64 `json:"ConfigServerCpu,omitnil,omitempty" name:"ConfigServerCpu"`

	// 指 ConfigServer 内存大小，固定取值为 2，单位：GB，可不配置该参数。
	ConfigServerMemory *uint64 `json:"ConfigServerMemory,omitnil,omitempty" name:"ConfigServerMemory"`

	// 指 ConfigServer 磁盘大小，固定取值为 20，单位：GB，可不配置该参数。
	ConfigServerVolume *uint64 `json:"ConfigServerVolume,omitnil,omitempty" name:"ConfigServerVolume"`
}

type InquirePriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所属区域及可用区信息。具体信息，请参见[地域和可用区](https://cloud.tencent.com/document/product/240/3637)。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// - 创建副本集实例，指每个副本集内主从节点数量。每个副本集所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 创建分片集群实例，指每个分片的主从节点数量。每个分片所支持的最大节点数与最小节点数，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例内存大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数CPU与Memory分别对应CPU核数与内存规格。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小。
	// - 单位：GB。
	// - 取值范围：请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MinStorage与MaxStorage分别对应其最小磁盘规格与最大磁盘规格。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例版本信息。具体支持的版本，请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询，其返回的数据结构SpecItems中的参数MongoVersionCode为实例所支持的版本信息。版本信息与版本号对应关系如下：
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 产品规格类型。
	// - HIO10G：通用高HIO万兆型。
	// - HCD：云盘版。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例数量，取值范围为[1,10]。
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例类型。
	// - REPLSET：副本集。
	// - SHARD：分片集群。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// - 创建副本集实例，指副本集数量，该参数只能为1。
	// - 创建分片集群实例，指分片的数量。请通过接口[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)查询分片数量的取值范围，其返回的数据结构SpecItems中的参数MinReplicateSetNum与MaxReplicateSetNum分别对应其最小值与最大值。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// - 选择包年包月计费模式，即 <b>InstanceChargeType </b>设定为<b>PREPAID</b>时，必须设置该参数，指定购买实例的购买时长。取值可选：[1,2,3,4,5,6,7,8,9,10,11,12,24,36]；单位：月。
	// -选择按量计费，即 <b>InstanceChargeType</b> 设定为 **POSTPAID_BY_HOUR** 时，该参数仅可配置为 1。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例付费方式。
	// - PREPAID：包年包月计费。
	// - POSTPAID_BY_HOUR：按量计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Mongos CPU 核数，支持1、2、4、8、16。购买分片集群时，必须填写。注意为空时取默认取值为2C。
	MongosCpu *uint64 `json:"MongosCpu,omitnil,omitempty" name:"MongosCpu"`

	// Mongos 内存大小。-  购买分片集群时，必须填写。- 单位：GB，支持1核2GB、2核4GB、4核8GB、8核16GB、16核32GB。注意为空时取默认取值为4G。
	MongosMemory *uint64 `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// 指 Mongos 个数，取值范围为[3,32]。若为分片集群实例询价，则该参数必须设置。注意为空时取默认取值为3个节点。
	MongosNum *uint64 `json:"MongosNum,omitnil,omitempty" name:"MongosNum"`

	// 指 ConfigServer CPU核数，固定取值为 1，单位：GB，可不配置该参数。
	ConfigServerCpu *uint64 `json:"ConfigServerCpu,omitnil,omitempty" name:"ConfigServerCpu"`

	// 指 ConfigServer 内存大小，固定取值为 2，单位：GB，可不配置该参数。
	ConfigServerMemory *uint64 `json:"ConfigServerMemory,omitnil,omitempty" name:"ConfigServerMemory"`

	// 指 ConfigServer 磁盘大小，固定取值为 20，单位：GB，可不配置该参数。
	ConfigServerVolume *uint64 `json:"ConfigServerVolume,omitnil,omitempty" name:"ConfigServerVolume"`
}

func (r *InquirePriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "NodeNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "MongoVersion")
	delete(f, "MachineCode")
	delete(f, "GoodsNum")
	delete(f, "ClusterType")
	delete(f, "ReplicateSetNum")
	delete(f, "Period")
	delete(f, "InstanceChargeType")
	delete(f, "MongosCpu")
	delete(f, "MongosMemory")
	delete(f, "MongosNum")
	delete(f, "ConfigServerCpu")
	delete(f, "ConfigServerMemory")
	delete(f, "ConfigServerVolume")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDBInstancesResponseParams struct {
	// 价格
	Price *DBInstancePrice `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateDBInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDBInstanceSpecRequestParams struct {
	// 实例 ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变更配置后实例内存大小，单位：GB。具体售卖的内存规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 变更配置后实例磁盘大小，单位：GB。每一个 CPU 规格对应的最大磁盘与最小磁盘范围，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例节点数量。请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 副本集实例，指变更配置后实例的主从节点数量。
	// - 分片集群实例，指变更配置后实例每一个分片的主从节点数。
	// **说明**：切勿同时发起调整节点数、调整分片数、调整节点规格的任务。
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 分片集群实例，指变更配置后实例的分片数量。取值范围：[2,36] 。
	// **说明**：变更后的分片数量不能小于当前现有的数量。切勿同时发起调整节点数、调整分片数与调整节点规格的任务。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`
}

type InquirePriceModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 变更配置后实例内存大小，单位：GB。具体售卖的内存规格，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 变更配置后实例磁盘大小，单位：GB。每一个 CPU 规格对应的最大磁盘与最小磁盘范围，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例节点数量。请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。
	// - 副本集实例，指变更配置后实例的主从节点数量。
	// - 分片集群实例，指变更配置后实例每一个分片的主从节点数。
	// **说明**：切勿同时发起调整节点数、调整分片数、调整节点规格的任务。
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 分片集群实例，指变更配置后实例的分片数量。取值范围：[2,36] 。
	// **说明**：变更后的分片数量不能小于当前现有的数量。切勿同时发起调整节点数、调整分片数与调整节点规格的任务。
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`
}

func (r *InquirePriceModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "NodeNum")
	delete(f, "ReplicateSetNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceModifyDBInstanceSpecResponseParams struct {
	// 价格。
	Price *DBInstancePrice `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *InquirePriceModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDBInstancesRequestParams struct {
	// 实例ID。请登录[MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID，且单次最多同时查询5个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式（即包年包月）相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

type InquirePriceRenewDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。请登录[MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID，且单次最多同时查询5个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式（即包年包月）相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquirePriceRenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewDBInstancesResponseParams struct {
	// 价格
	Price *DBInstancePrice `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceRenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewDBInstancesResponseParams `json:"Response"`
}

func (r *InquirePriceRenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。默认为1。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// - NOTIFY_AND_AUTO_RENEW：通知过期且自动续费。在账户余额充足的情况下，实例到期后将按月自动续费。
	// - NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。默认为NOTIFY_AND_MANUAL_RENEW。
	// - DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceDetail struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 付费类型。
	// - 1：包年包月。
	// - 0：按量计费。
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群类型。
	// - 0：副本集实例。
	// - 1：分片实例。
	ClusterType *uint64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 地域信息。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区信息。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 网络类型。
	// - 0：基础网络。
	// - 1：私有网络。
	NetType *uint64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 私有网络的ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例状态。
	// - 0：待初始化。
	// - 1：流程处理中，例如：变更规格、参数修改等。
	// - 2：实例正常运行中。
	// - -2：已隔离（包年包月）。
	// - -3：已隔离（按量计费）。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口号。
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例到期时间。
	DeadLine *string `json:"DeadLine,omitnil,omitempty" name:"DeadLine"`

	// 实例存储引擎版本信息。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 实例内存规格，单位：MB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例磁盘规格，单位：MB。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例 CPU 核心数。
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 实例机器类型。
	// - HIO10G：通用高 HIO 万兆型。
	// - HCD：云盘版类型。
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// 实例从节点数。
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// 实例分片数。
	ReplicationSetNum *uint64 `json:"ReplicationSetNum,omitnil,omitempty" name:"ReplicationSetNum"`

	// 实例自动续费标志。
	// - 0：手动续费。
	// - 1：自动续费。
	// - 2：确认不续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 已用容量，单位：MB。
	UsedVolume *uint64 `json:"UsedVolume,omitnil,omitempty" name:"UsedVolume"`

	// 维护窗口起始时间。
	MaintenanceStart *string `json:"MaintenanceStart,omitnil,omitempty" name:"MaintenanceStart"`

	// 维护窗口结束时间。
	MaintenanceEnd *string `json:"MaintenanceEnd,omitnil,omitempty" name:"MaintenanceEnd"`

	// 分片信息。
	ReplicaSets []*ShardInfo `json:"ReplicaSets,omitnil,omitempty" name:"ReplicaSets"`

	// 只读实例信息。
	ReadonlyInstances []*DBInstanceInfo `json:"ReadonlyInstances,omitnil,omitempty" name:"ReadonlyInstances"`

	// 灾备实例信息。
	StandbyInstances []*DBInstanceInfo `json:"StandbyInstances,omitnil,omitempty" name:"StandbyInstances"`

	// 临时实例信息。
	CloneInstances []*DBInstanceInfo `json:"CloneInstances,omitnil,omitempty" name:"CloneInstances"`

	// 关联实例信息，对于正式实例，该字段表示它的临时实例信息；对于临时实例，则表示它的正式实例信息;如果为只读/灾备实例,则表示他的主实例信息。
	RelatedInstance *DBInstanceInfo `json:"RelatedInstance,omitnil,omitempty" name:"RelatedInstance"`

	// 实例标签信息集合。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例版本标记。
	InstanceVer *uint64 `json:"InstanceVer,omitnil,omitempty" name:"InstanceVer"`

	// 实例版本标记。
	ClusterVer *uint64 `json:"ClusterVer,omitnil,omitempty" name:"ClusterVer"`

	// 协议信息：mongodb。
	Protocol *uint64 `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 实例类型。
	// - 0：所有实例。
	// - 1：正式实例。
	// - 2：临时实例
	// - 3：只读实例。
	// - -1：同时包括正式实例、只读实例与灾备实例。
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态描述。
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil,omitempty" name:"InstanceStatusDesc"`

	// 实例对应的物理实例 ID。回档并替换过的实例有不同的 InstanceId 和 RealInstanceId，从 barad 获取监控数据等场景下需要用物理 ID 获取。
	RealInstanceId *string `json:"RealInstanceId,omitnil,omitempty" name:"RealInstanceId"`

	// 实例当前可用区信息。
	ZoneList []*string `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// mongos 节点个数。
	MongosNodeNum *uint64 `json:"MongosNodeNum,omitnil,omitempty" name:"MongosNodeNum"`

	// mongos 节点内存。单位：MB。
	MongosMemory *uint64 `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// mongos 节点 CPU 核数。
	MongosCpuNum *uint64 `json:"MongosCpuNum,omitnil,omitempty" name:"MongosCpuNum"`

	// Config Server节点个数。
	ConfigServerNodeNum *uint64 `json:"ConfigServerNodeNum,omitnil,omitempty" name:"ConfigServerNodeNum"`

	// Config Server节点内存。单位：MB。
	ConfigServerMemory *uint64 `json:"ConfigServerMemory,omitnil,omitempty" name:"ConfigServerMemory"`

	// Config Server节点磁盘大小。单位：MB。
	ConfigServerVolume *uint64 `json:"ConfigServerVolume,omitnil,omitempty" name:"ConfigServerVolume"`

	// Config Server 节点 CPU 核数。
	ConfigServerCpuNum *uint64 `json:"ConfigServerCpuNum,omitnil,omitempty" name:"ConfigServerCpuNum"`

	// readonly节点个数。
	ReadonlyNodeNum *uint64 `json:"ReadonlyNodeNum,omitnil,omitempty" name:"ReadonlyNodeNum"`
}

// Predefined struct for user
type InstanceEnableSSLRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置是否要开启SSL访问。
	// - true：开启。
	// - false：关闭。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type InstanceEnableSSLRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置是否要开启SSL访问。
	// - true：开启。
	// - false：关闭。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *InstanceEnableSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstanceEnableSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstanceEnableSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstanceEnableSSLResponseParams struct {
	// SSL开启状态。
	// - 0：关闭。
	// - 1：开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书文件过期时间，格式为：2023-05-01 12:00:00。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 证书文件的下载链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertUrl *string `json:"CertUrl,omitnil,omitempty" name:"CertUrl"`

	// 流程id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstanceEnableSSLResponse struct {
	*tchttp.BaseResponse
	Response *InstanceEnableSSLResponseParams `json:"Response"`
}

func (r *InstanceEnableSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstanceEnableSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceEnumParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 枚举值，所有支持的值。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 参数修改之后是否需要重启生效。
	// - 1：需要重启后生效。
	// - 0：无需重启，设置成功即可生效。
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 参数值类型说明。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 是否为运行中参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InstanceInfo struct {
	// 审计日志保存时长。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditLogExpireDay *uint64 `json:"AuditLogExpireDay,omitnil,omitempty" name:"AuditLogExpireDay"`

	// 审计状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例角色。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 数据库版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MongodbVersion *string `json:"MongodbVersion,omitnil,omitempty" name:"MongodbVersion"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否支持审计。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportAudit *bool `json:"SupportAudit,omitnil,omitempty" name:"SupportAudit"`

	// 可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type InstanceIntegerParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 参数最大值。
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// 最小值。
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数修改之后是否需要重启生效。
	// - 1:需要重启后生效。
	// - 0：无需重启，设置成功即可生效。
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 参数类型。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 是否为运行中参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 冗余字段，可忽略。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type InstanceMultiParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 参考值范围。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 参数修改后是否需要重启才会生效。
	// - 1：需要重启后生效。
	// - 0：无需重启，设置成功即可生效。
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 是否为运行中参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 当前值的类型描述，默认为multi。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`
}

type InstanceTextParam struct {
	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 修改参数值之后是否需要重启。
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Text 类型参数对应的值。
	TextValue *string `json:"TextValue,omitnil,omitempty" name:"TextValue"`

	// 参数说明。
	Tips []*string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 参数值类型说明。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 是否为运行中的参数值。
	// - 1：运行中参数值。
	// - 0：非运行中参数值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制需隔离的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制需隔离的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KMSInfoDetail struct {
	// 主密钥 ID。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 主密钥名称。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 实例与密钥绑定时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 密钥状态。
	// - Enabled：开启。
	// - Disabled：不开启。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 密钥用途。
	KeyUsage *string `json:"KeyUsage,omitnil,omitempty" name:"KeyUsage"`

	// 密钥来源。
	KeyOrigin *string `json:"KeyOrigin,omitnil,omitempty" name:"KeyOrigin"`

	// kms所在地域。
	KmsRegion *string `json:"KmsRegion,omitnil,omitempty" name:"KmsRegion"`
}

// Predefined struct for user
type KillOpsRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待终止的操作。
	Operations []*Operation `json:"Operations,omitnil,omitempty" name:"Operations"`
}

type KillOpsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待终止的操作。
	Operations []*Operation `json:"Operations,omitnil,omitempty" name:"Operations"`
}

func (r *KillOpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillOpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Operations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillOpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillOpsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillOpsResponse struct {
	*tchttp.BaseResponse
	Response *KillOpsResponseParams `json:"Response"`
}

func (r *KillOpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillOpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFilter struct {
	// 过滤条件名称
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 过滤条件匹配类型，注意：此参数取值只能等于EQ
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`

	// 过滤条件匹配值
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogInfo struct {
	// 日志类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogComponent *string `json:"LogComponent,omitnil,omitempty" name:"LogComponent"`

	// 日志级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// 日志产生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTime *string `json:"LogTime,omitnil,omitempty" name:"LogTime"`

	// 日志详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogDetail *string `json:"LogDetail,omitnil,omitempty" name:"LogDetail"`

	// 日志连接信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConnection *string `json:"LogConnection,omitnil,omitempty" name:"LogConnection"`

	// 日志id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`
}

// Predefined struct for user
type ModifyAuditServiceRequestParams struct {
	// 实例ID，格式如：cmgo-xfts****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。单位为：天。当前支持的取值包括： 7： 一周。 30： 一个月。 90： 三个月。 180 ： 六个月。 365 ： 一年。 1095 ： 三年。 1825 ： 五年。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// true-全审计，false-规则审计，注意：AuditAll=true 时，RuleFilters 无需填参
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 审计过滤规则，Type的范围【SrcIp、DB、Collection、User、SqlType】，注意：Type=SqlType时，Value必须在这个范围 ["query", "insert", "update", "delete", "command"]
	RuleFilters []*LogFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-xfts****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。单位为：天。当前支持的取值包括： 7： 一周。 30： 一个月。 90： 三个月。 180 ： 六个月。 365 ： 一年。 1095 ： 三年。 1825 ： 五年。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// true-全审计，false-规则审计，注意：AuditAll=true 时，RuleFilters 无需填参
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 审计过滤规则，Type的范围【SrcIp、DB、Collection、User、SqlType】，注意：Type=SqlType时，Value必须在这个范围 ["query", "insert", "update", "delete", "command"]
	RuleFilters []*LogFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

func (r *ModifyAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "AuditAll")
	delete(f, "RuleFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditServiceResponseParams `json:"Response"`
}

func (r *ModifyAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkAddressRequestParams struct {
	// 指定需修改网络的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 原 IP 地址保留时长。
	// - 单位为分钟，0表示立即回收原 IP 地址。
	// - 原 IP 将在约定时间后释放，在释放前原 IP和新 IP均可访问。
	OldIpExpiredTime *uint64 `json:"OldIpExpiredTime,omitnil,omitempty" name:"OldIpExpiredTime"`

	// 切换后的私有网络 ID，若实例当前为基础网络，该字段无需配置。请通过接口 [DescribeDBInstances](https://cloud.tencent.com/document/product/240/38568) 获取私有网络 ID。
	NewUniqVpcId *string `json:"NewUniqVpcId,omitnil,omitempty" name:"NewUniqVpcId"`

	// 切换后私有网络的子网 ID。若实例当前为基础网络，该字段无需配置。请通过接口 [DescribeDBInstances](https://cloud.tencent.com/document/product/240/38568) 获取私有网络的子网 ID。
	NewUniqSubnetId *string `json:"NewUniqSubnetId,omitnil,omitempty" name:"NewUniqSubnetId"`

	// IP 地址信息，包含新 IP 地址与 原 IP 地址。
	NetworkAddresses []*ModifyNetworkAddress `json:"NetworkAddresses,omitnil,omitempty" name:"NetworkAddresses"`
}

type ModifyDBInstanceNetworkAddressRequest struct {
	*tchttp.BaseRequest
	
	// 指定需修改网络的实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 原 IP 地址保留时长。
	// - 单位为分钟，0表示立即回收原 IP 地址。
	// - 原 IP 将在约定时间后释放，在释放前原 IP和新 IP均可访问。
	OldIpExpiredTime *uint64 `json:"OldIpExpiredTime,omitnil,omitempty" name:"OldIpExpiredTime"`

	// 切换后的私有网络 ID，若实例当前为基础网络，该字段无需配置。请通过接口 [DescribeDBInstances](https://cloud.tencent.com/document/product/240/38568) 获取私有网络 ID。
	NewUniqVpcId *string `json:"NewUniqVpcId,omitnil,omitempty" name:"NewUniqVpcId"`

	// 切换后私有网络的子网 ID。若实例当前为基础网络，该字段无需配置。请通过接口 [DescribeDBInstances](https://cloud.tencent.com/document/product/240/38568) 获取私有网络的子网 ID。
	NewUniqSubnetId *string `json:"NewUniqSubnetId,omitnil,omitempty" name:"NewUniqSubnetId"`

	// IP 地址信息，包含新 IP 地址与 原 IP 地址。
	NetworkAddresses []*ModifyNetworkAddress `json:"NetworkAddresses,omitnil,omitempty" name:"NetworkAddresses"`
}

func (r *ModifyDBInstanceNetworkAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldIpExpiredTime")
	delete(f, "NewUniqVpcId")
	delete(f, "NewUniqSubnetId")
	delete(f, "NetworkAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNetworkAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkAddressResponseParams struct {
	// 修改网络异步流程任务ID。
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNetworkAddressResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNetworkAddressResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNetworkAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceParamTplRequestParams struct {
	// 待修改的参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`

	// 待修改参数模板名称，为空时，保持原有名称。
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 待修改参数模板描述，为空时，保持原有描述。
	TplDesc *string `json:"TplDesc,omitnil,omitempty" name:"TplDesc"`

	// 待修改参数名及参数值，为空时，各参数保持原有值，支持单条或批量修改。
	Params []*ParamType `json:"Params,omitnil,omitempty" name:"Params"`
}

type ModifyDBInstanceParamTplRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的参数模板 ID。请通过接口 [DescribeDBInstanceParamTpl](https://cloud.tencent.com/document/product/240/109155) 获取模板 ID。
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`

	// 待修改参数模板名称，为空时，保持原有名称。
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 待修改参数模板描述，为空时，保持原有描述。
	TplDesc *string `json:"TplDesc,omitnil,omitempty" name:"TplDesc"`

	// 待修改参数名及参数值，为空时，各参数保持原有值，支持单条或批量修改。
	Params []*ParamType `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *ModifyDBInstanceParamTplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceParamTplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TplId")
	delete(f, "TplName")
	delete(f, "TplDesc")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceParamTplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceParamTplResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceParamTplResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceParamTplResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceParamTplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceParamTplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标安全组 ID。请登录[安全组控制台页面](https://console.cloud.tencent.com/vpc/security-group)复制目标安全组 ID。
	// **注意**：该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type ModifyDBInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标安全组 ID。请登录[安全组控制台页面](https://console.cloud.tencent.com/vpc/security-group)复制目标安全组 ID。
	// **注意**：该入参会全量替换存量已有集合，非增量更新。修改需传入预期的全量集合。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyDBInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例配置变更后的内存大小。单位：GB。该参数为空值时，默认取实例当前的内存大小。当前所支持的内存规格，请参见[产品规格](https://cloud.tencent.com/document/product/240/64125)。
	// **注意**：内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例配置变更后的硬盘大小，单位：GB。该参数为空值时，默认取当前实例的磁盘大小。当前所支持的磁盘容量，请参见[产品规格](https://cloud.tencent.com/document/product/240/64125)。
	// - 内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。
	// - 降配时，变更后的磁盘容量必须大于已用磁盘容量的1.2倍。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// (已废弃) 请使用ResizeOplog独立接口完成。
	// 
	// 实例配置变更后 Oplog 的大小。
	// - 单位：GB。
	// - 默认 Oplog 占用容量为磁盘空间的10%。系统允许设置的 Oplog 容量范围为磁盘空间的[10%,90%]。
	//
	// Deprecated: OplogSize is deprecated.
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`

	// 实例变更后 mongod 的节点数（不包含 readonly 只读节点数）。
	// -  副本集节点数：请通过 [DescribeSpecInfo ](https://cloud.tencent.com/document/product/240/38567)接口返回的参数 MinNodeNum 与 MaxNodeNum 获取节点数量取值范围。
	// -  分片集群每个分片节点数：请通过 [DescribeSpecInfo ](https://cloud.tencent.com/document/product/240/38567)接口返回的参数 MinReplicateSetNodeNum 与 MaxReplicateSetNodeNum 获取节点数量取值范围。
	// **说明**：变更 mongod 或 mongos 的 CPU 与内存规格时，该参数可以不配置或者输入当前 mongod 或 mongos（不包含readonly）节点数量。
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例变更后的分片数。
	// - 请通过 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 接口返回的参数**MinReplicateSetNum**与**MaxReplicateSetNum**获取实例分片数取值范围。
	// - 实例分片数量只允许增加不允许减少。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// 实例配置变更的切换时间。
	// - 0：调整完成时，立即执行变配任务。默认为0。
	// - 1：在维护时间窗内，执行变配任务。
	// **说明**：调整节点数和分片数不支持在<b>维护时间窗内</b>变更。
	InMaintenance *uint64 `json:"InMaintenance,omitnil,omitempty" name:"InMaintenance"`

	// 分片实例配置变更后的 mongos 内存大小。单位：GB。实例支持的规格，请参见[产品规格](https://cloud.tencent.com/document/product/240/64125)。
	MongosMemory *string `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// 新增节点列表，节点类型及可用区信息。
	AddNodeList []*AddNodeList `json:"AddNodeList,omitnil,omitempty" name:"AddNodeList"`

	// 删除节点列表。
	// **注意**：基于分片实例各片节点的一致性原则，删除分片实例节点时，只需指定0分片对应的节点即可，如：cmgo-9nl1czif_0-node-readonly0 将删除每个分片的第1个只读节点。
	RemoveNodeList []*RemoveNodeList `json:"RemoveNodeList,omitnil,omitempty" name:"RemoveNodeList"`
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例配置变更后的内存大小。单位：GB。该参数为空值时，默认取实例当前的内存大小。当前所支持的内存规格，请参见[产品规格](https://cloud.tencent.com/document/product/240/64125)。
	// **注意**：内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例配置变更后的硬盘大小，单位：GB。该参数为空值时，默认取当前实例的磁盘大小。当前所支持的磁盘容量，请参见[产品规格](https://cloud.tencent.com/document/product/240/64125)。
	// - 内存和磁盘必须同时升配或同时降配，即 Memory 与 Volume 需同时配置变更。
	// - 降配时，变更后的磁盘容量必须大于已用磁盘容量的1.2倍。
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// (已废弃) 请使用ResizeOplog独立接口完成。
	// 
	// 实例配置变更后 Oplog 的大小。
	// - 单位：GB。
	// - 默认 Oplog 占用容量为磁盘空间的10%。系统允许设置的 Oplog 容量范围为磁盘空间的[10%,90%]。
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`

	// 实例变更后 mongod 的节点数（不包含 readonly 只读节点数）。
	// -  副本集节点数：请通过 [DescribeSpecInfo ](https://cloud.tencent.com/document/product/240/38567)接口返回的参数 MinNodeNum 与 MaxNodeNum 获取节点数量取值范围。
	// -  分片集群每个分片节点数：请通过 [DescribeSpecInfo ](https://cloud.tencent.com/document/product/240/38567)接口返回的参数 MinReplicateSetNodeNum 与 MaxReplicateSetNodeNum 获取节点数量取值范围。
	// **说明**：变更 mongod 或 mongos 的 CPU 与内存规格时，该参数可以不配置或者输入当前 mongod 或 mongos（不包含readonly）节点数量。
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例变更后的分片数。
	// - 请通过 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 接口返回的参数**MinReplicateSetNum**与**MaxReplicateSetNum**获取实例分片数取值范围。
	// - 实例分片数量只允许增加不允许减少。
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// 实例配置变更的切换时间。
	// - 0：调整完成时，立即执行变配任务。默认为0。
	// - 1：在维护时间窗内，执行变配任务。
	// **说明**：调整节点数和分片数不支持在<b>维护时间窗内</b>变更。
	InMaintenance *uint64 `json:"InMaintenance,omitnil,omitempty" name:"InMaintenance"`

	// 分片实例配置变更后的 mongos 内存大小。单位：GB。实例支持的规格，请参见[产品规格](https://cloud.tencent.com/document/product/240/64125)。
	MongosMemory *string `json:"MongosMemory,omitnil,omitempty" name:"MongosMemory"`

	// 新增节点列表，节点类型及可用区信息。
	AddNodeList []*AddNodeList `json:"AddNodeList,omitnil,omitempty" name:"AddNodeList"`

	// 删除节点列表。
	// **注意**：基于分片实例各片节点的一致性原则，删除分片实例节点时，只需指定0分片对应的节点即可，如：cmgo-9nl1czif_0-node-readonly0 将删除每个分片的第1个只读节点。
	RemoveNodeList []*RemoveNodeList `json:"RemoveNodeList,omitnil,omitempty" name:"RemoveNodeList"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "OplogSize")
	delete(f, "NodeNum")
	delete(f, "ReplicateSetNum")
	delete(f, "InMaintenance")
	delete(f, "MongosMemory")
	delete(f, "AddNodeList")
	delete(f, "RemoveNodeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecResponseParams struct {
	// 订单 ID。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsRequestParams struct {
	// 指定实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定需修改的参数名及值。当前所支持的参数名及对应取值范围，请通过 [DescribeInstanceParams ](https://cloud.tencent.com/document/product/240/65903)获取。
	InstanceParams []*ModifyMongoDBParamType `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// 操作类型，包括：
	// - IMMEDIATELY：立即调整。
	// - DELAY：延迟调整。可选字段，不配置该参数则默认为立即调整。
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定需修改的参数名及值。当前所支持的参数名及对应取值范围，请通过 [DescribeInstanceParams ](https://cloud.tencent.com/document/product/240/65903)获取。
	InstanceParams []*ModifyMongoDBParamType `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`

	// 操作类型，包括：
	// - IMMEDIATELY：立即调整。
	// - DELAY：延迟调整。可选字段，不配置该参数则默认为立即调整。
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`
}

func (r *ModifyInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceParams")
	delete(f, "ModifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsResponseParams struct {
	// 修改参数配置是否生效。
	// - true：参数修改后的值已生效。
	// - false：执行失败。
	Changed *bool `json:"Changed,omitnil,omitempty" name:"Changed"`

	// 该参数暂时无意义(兼容前端保留)。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamsResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMongoDBParamType struct {
	// 需要修改的参数名称，请严格参考通过 DescribeInstanceParams 获取的当前实例支持的参数名。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 需要修改的参数名称对应的值，请严格参考通过 DescribeInstanceParams 获取的参数对应的值的范围。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ModifyNetworkAddress struct {
	// 新IP地址。
	NewIPAddress *string `json:"NewIPAddress,omitnil,omitempty" name:"NewIPAddress"`

	// 原IP地址。
	OldIpAddress *string `json:"OldIpAddress,omitnil,omitempty" name:"OldIpAddress"`
}

type NodeProperty struct {
	// 节点所在的可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 节点名称。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点访问地址。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 节点公网访问外网地址(IP或域名，示例为IP方式)。
	WanServiceAddress *string `json:"WanServiceAddress,omitnil,omitempty" name:"WanServiceAddress"`

	// 节点角色。
	// - PRIMARY：主节点。
	// - SECONDARY：从节点。
	// - READONLY：只读节点。
	// - ARBITER：仲裁节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 节点是否为 Hidden 节点。
	// - true：Hidden 节点。
	// - false：非 Hidden 节点。
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`

	// 节点状态。
	// - NORMAL：正常运行中。
	// - STARTUP：正在启动。
	// - STARTUP2：正在启动，处理中间数据。
	// - RECOVERING：恢复中，暂不可用。
	// - DOWN：已掉线。
	// - UNKNOWN：未知状态。
	// - ROLLBACK：回滚中。
	// - REMOVED：已移除。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 主从同步延迟时间，单位：秒。
	SlaveDelay *int64 `json:"SlaveDelay,omitnil,omitempty" name:"SlaveDelay"`

	// 节点优先级。其取值范围为[0,100]，数值越高，优先级越高。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 节点投票权。
	// - 1：具有投票权。
	// - 0：无投票权。
	Votes *int64 `json:"Votes,omitnil,omitempty" name:"Votes"`

	// 节点标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*NodeTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 副本集 ID。
	ReplicateSetId *string `json:"ReplicateSetId,omitnil,omitempty" name:"ReplicateSetId"`
}

type NodeTag struct {
	// 节点Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 节点Tag Value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type OfflineIsolatedDBInstanceRequestParams struct {
	// 实例ID。请登录 [MongoDB 控制台回收站](https://console.cloud.tencent.com/mongodb/recycle)在实例列表复制需下线的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OfflineIsolatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。请登录 [MongoDB 控制台回收站](https://console.cloud.tencent.com/mongodb/recycle)在实例列表复制需下线的实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *OfflineIsolatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OfflineIsolatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OfflineIsolatedDBInstanceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OfflineIsolatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *OfflineIsolatedDBInstanceResponseParams `json:"Response"`
}

func (r *OfflineIsolatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OfflineIsolatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// 实例 ID，格式如：cmgo-xfts****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。单位为：天。当前支持的取值包括： 7： 一周。 30： 一个月。 90： 三个月。 180 ： 六个月。 365 ： 一年。 1095 ： 三年。 1825 ： 五年。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// true-全审计，false-规则审计，注意：AuditAll=true 时，RuleFilters 无需填参
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 审计过滤规则，Type的范围【SrcIp、DB、Collection、User、SqlType】，注意：Type=SqlType时，Value必须在这个范围 ["query", "insert", "update", "delete", "command"]
	RuleFilters []*LogFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：cmgo-xfts****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志保存时长。单位为：天。当前支持的取值包括： 7： 一周。 30： 一个月。 90： 三个月。 180 ： 六个月。 365 ： 一年。 1095 ： 三年。 1825 ： 五年。
	LogExpireDay *uint64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// true-全审计，false-规则审计，注意：AuditAll=true 时，RuleFilters 无需填参
	AuditAll *bool `json:"AuditAll,omitnil,omitempty" name:"AuditAll"`

	// 审计过滤规则，Type的范围【SrcIp、DB、Collection、User、SqlType】，注意：Type=SqlType时，Value必须在这个范围 ["query", "insert", "update", "delete", "command"]
	RuleFilters []*LogFilter `json:"RuleFilters,omitnil,omitempty" name:"RuleFilters"`
}

func (r *OpenAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "AuditAll")
	delete(f, "RuleFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenAuditServiceResponseParams `json:"Response"`
}

func (r *OpenAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Operation struct {
	// 操作所在的分片名称。请通过接口 [DescribeCurrentOp](https://cloud.tencent.com/document/product/240/48120) 查询分片名称。
	ReplicaSetName *string `json:"ReplicaSetName,omitnil,omitempty" name:"ReplicaSetName"`

	// 操作所在的节点名。请通过接口 [DescribeCurrentOp](https://cloud.tencent.com/document/product/240/48120) 查询节点名称。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 操作序号。请通过接口 [DescribeCurrentOp](https://cloud.tencent.com/document/product/240/48120) 查询操作序号。
	OpId *int64 `json:"OpId,omitnil,omitempty" name:"OpId"`
}

type ParamTpl struct {
	// 参数模板名称。
	TplName *string `json:"TplName,omitnil,omitempty" name:"TplName"`

	// 参数模板 ID。
	TplId *string `json:"TplId,omitnil,omitempty" name:"TplId"`

	// 参数模板适用的数据库版本。
	// - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本，
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 参数模板适用的数据库类型。
	// - REPLSET：副本集实例。
	// - SHARD：分片实例。
	// - STANDALONE：单节点实例。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 参数模板描述。
	TplDesc *string `json:"TplDesc,omitnil,omitempty" name:"TplDesc"`

	// 模板类型。
	// - DEFAULT：系统默认模板。
	// - CUSTOMIZE：自定义模板。
	TplType *string `json:"TplType,omitnil,omitempty" name:"TplType"`
}

type ParamType struct {
	// 参数
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RemoveNodeList struct {
	// 需要删除的节点角色。
	// - SECONDARY：Mongod 从节点。
	// - READONLY：只读节点。
	// - MONGOS：Mongos 节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 要删除的节点 ID。分片集群须指定一组分片要删除的节点名称即可，其余分片对该组对齐。
	// - 获取方式：登录 [MongoDB控制台](https://console.cloud.tencent.com/mongodb)，在**节点管理**页签，可获取**节点 ID**。
	// - 特别说明：分片集群同一节点上的分片，仅需指定0分片节点 ID 即可。例如：cmgo-6hfk\*\*\*\*\_0-node-primary。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点所对应的可用区。当前支持可用区信息，请参见[地域和可用区](https://cloud.tencent.com/document/product/240/3637)。
	// - 单可用区，所有节点在同一可用区。
	// - 多可用区：当前标准规格是三可用区分布，主从节点不在同一可用区，需注意配置所删除节点对应的可用区，且删除后必须满足任意2个可用区节点数大于第3个可用区原则。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

// Predefined struct for user
type RenameInstanceRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。请登录[MongoDB 控制台](https://console.cloud.tencent.com/mongodb#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义实例名称，要求为1～128 长度的任意字符。
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`
}

type RenameInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。请登录[MongoDB 控制台](https://console.cloud.tencent.com/mongodb#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义实例名称，要求为1～128 长度的任意字符。
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenameInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenameInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenameInstanceResponseParams `json:"Response"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstancesRequestParams struct {
	// 指定续费的一个或多个待操作的实例ID。
	// - 可通过[DescribeDBInstances](https://cloud.tencent.com/document/product/240/38568)接口返回值中的**InstanceId**获取。
	// - 每次续费请求的实例数量上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。包年包月实例该参数为必传参数。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

type RenewDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 指定续费的一个或多个待操作的实例ID。
	// - 可通过[DescribeDBInstances](https://cloud.tencent.com/document/product/240/38568)接口返回值中的**InstanceId**获取。
	// - 每次续费请求的实例数量上限为100。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。包年包月实例该参数为必传参数。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`
}

func (r *RenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewDBInstancesResponseParams `json:"Response"`
}

func (r *RenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicaSetInfo struct {
	// 副本集 ID。
	ReplicaSetId *string `json:"ReplicaSetId,omitnil,omitempty" name:"ReplicaSetId"`
}

type ReplicateSetInfo struct {
	// 节点属性
	Nodes []*NodeProperty `json:"Nodes,omitnil,omitempty" name:"Nodes"`
}

// Predefined struct for user
type ResetDBInstancePasswordRequestParams struct {
	// 指定实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定需修改密码的账号名称。可通过接口 [DescribeAccountUsers](https://cloud.tencent.com/document/product/240/80800) 获取账号列表，复制需修改密码的账号。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 指定账户的新密码。密码复杂度要求：
	// - 8-32个字符长度。
	// - 至少包含字母、数字和字符（!@#%^\*()\_）中的两种。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ResetDBInstancePasswordRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指定需修改密码的账号名称。可通过接口 [DescribeAccountUsers](https://cloud.tencent.com/document/product/240/80800) 获取账号列表，复制需修改密码的账号。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 指定账户的新密码。密码复杂度要求：
	// - 8-32个字符长度。
	// - 至少包含字母、数字和字符（!@#%^\*()\_）中的两种。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ResetDBInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDBInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDBInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDBInstancePasswordResponseParams struct {
	// 任务请求 ID。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetDBInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetDBInstancePasswordResponseParams `json:"Response"`
}

func (r *ResetDBInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDBInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartNodesRequestParams struct {
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要重启的节点 ID 列表。
	// - 支持重启的节点类型包含：mongod节点、只读节点、mongos节点。
	// - 节点 ID，请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在**节点管理**页面复制，或者通过 [DescribeDBInstanceNodeProperty ](https://cloud.tencent.com/document/product/240/82022)接口获取。
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

type RestartNodesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要重启的节点 ID 列表。
	// - 支持重启的节点类型包含：mongod节点、只读节点、mongos节点。
	// - 节点 ID，请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在**节点管理**页面复制，或者通过 [DescribeDBInstanceNodeProperty ](https://cloud.tencent.com/document/product/240/82022)接口获取。
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

func (r *RestartNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartNodesResponseParams struct {
	// 流程 ID。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartNodesResponse struct {
	*tchttp.BaseResponse
	Response *RestartNodesResponseParams `json:"Response"`
}

func (r *RestartNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 所属项目 ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 安全组入站规则。
	Inbound []*SecurityGroupBound `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// 安全组出站规则。
	Outbound []*SecurityGroupBound `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称。
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组备注信息。
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`
}

type SecurityGroupBound struct {
	// 执行策略。
	// - ACCEPT：允许，放行该端口相应的访问请求。
	// - DROP：拒绝，直接丢弃数据包，不返回任何回应信息。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 访问数据库的入站 IP 或 IP 段。
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 访问数据库的端口。
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 传输层协议：tcp。
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 安全组 ID。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// IP 地址或 IP 地址组参数模板 ID。请登录[参数模板控制台](https://console.cloud.tencent.com/vpc/template/ip)获取参数模板 IP 地址详情。
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// 协议端口或协议端口组参数模板 ID。请登录[参数模板控制台](https://console.cloud.tencent.com/vpc/template/protoport)获取参数模板协议端口详情。
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// 安全组描述信息。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

// Predefined struct for user
type SetAccountUserPrivilegeRequestParams struct {
	// 指定待设置账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置访问实例的账号名称。设置要求为：字母开头的1-64个字符，只可输入[A,Z]、[a,z]、[1,9]范围的字符以及下划线“_”与短划线“-”。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 设置权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil,omitempty" name:"AuthRole"`
}

type SetAccountUserPrivilegeRequest struct {
	*tchttp.BaseRequest
	
	// 指定待设置账号的实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置访问实例的账号名称。设置要求为：字母开头的1-64个字符，只可输入[A,Z]、[a,z]、[1,9]范围的字符以及下划线“_”与短划线“-”。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 设置权限信息。
	AuthRole []*Auth `json:"AuthRole,omitnil,omitempty" name:"AuthRole"`
}

func (r *SetAccountUserPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAccountUserPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "AuthRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAccountUserPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAccountUserPrivilegeResponseParams struct {
	// 任务ID。
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetAccountUserPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *SetAccountUserPrivilegeResponseParams `json:"Response"`
}

func (r *SetAccountUserPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAccountUserPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetBackupRulesRequestParams struct {
	// 实例id，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明**:
	// 1. 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// 2. 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *uint64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 设置自动备份开始时间。取值范围为：[0,23]，例如：该参数设置为2，表示02:00开始备份。
	BackupTime *uint64 `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// 自动备份频率，内部展示，默认取值为24h。
	BackupFrequency *uint64 `json:"BackupFrequency,omitnil,omitempty" name:"BackupFrequency"`

	// 设置自动备份发生错误时，是否发送失败告警。
	// - true：发送。
	// - false：不发送。
	Notify *bool `json:"Notify,omitnil,omitempty" name:"Notify"`

	// 指定备份数据保存天数。默认为 7 天，支持设置为7、30、90、180、365。
	BackupRetentionPeriod *uint64 `json:"BackupRetentionPeriod,omitnil,omitempty" name:"BackupRetentionPeriod"`

	// 周几备份，0-6，逗号分割。仅对高级备份生效
	ActiveWeekdays *string `json:"ActiveWeekdays,omitnil,omitempty" name:"ActiveWeekdays"`

	// 长期保留周期，周weekly，月monthly，空不开启
	LongTermUnit *string `json:"LongTermUnit,omitnil,omitempty" name:"LongTermUnit"`

	// 长期保留哪些天的，周0-6，月1-31，逗号分割
	LongTermActiveDays *string `json:"LongTermActiveDays,omitnil,omitempty" name:"LongTermActiveDays"`

	// 长期备份保留多少天
	LongTermExpiredDays *int64 `json:"LongTermExpiredDays,omitnil,omitempty" name:"LongTermExpiredDays"`

	// 增量保留多少天
	OplogExpiredDays *int64 `json:"OplogExpiredDays,omitnil,omitempty" name:"OplogExpiredDays"`

	// 备份版本。旧版本备份为0，高级备份为1。开启高级备份此值传1
	BackupVersion *int64 `json:"BackupVersion,omitnil,omitempty" name:"BackupVersion"`

	// 告警额度。50-300
	AlarmWaterLevel *int64 `json:"AlarmWaterLevel,omitnil,omitempty" name:"AlarmWaterLevel"`
}

type SetBackupRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id，例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份方式。
	// - 0：逻辑备份。
	// - 1：物理备份。
	// - 3：快照备份。
	// **说明**:
	// 1. 通用版实例支持逻辑备份与物理备份。云盘版实例支持物理备份与快照备份，暂不支持逻辑备份。
	// 2. 实例开通存储加密，则备份方式不能为物理备份。
	BackupMethod *uint64 `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// 设置自动备份开始时间。取值范围为：[0,23]，例如：该参数设置为2，表示02:00开始备份。
	BackupTime *uint64 `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// 自动备份频率，内部展示，默认取值为24h。
	BackupFrequency *uint64 `json:"BackupFrequency,omitnil,omitempty" name:"BackupFrequency"`

	// 设置自动备份发生错误时，是否发送失败告警。
	// - true：发送。
	// - false：不发送。
	Notify *bool `json:"Notify,omitnil,omitempty" name:"Notify"`

	// 指定备份数据保存天数。默认为 7 天，支持设置为7、30、90、180、365。
	BackupRetentionPeriod *uint64 `json:"BackupRetentionPeriod,omitnil,omitempty" name:"BackupRetentionPeriod"`

	// 周几备份，0-6，逗号分割。仅对高级备份生效
	ActiveWeekdays *string `json:"ActiveWeekdays,omitnil,omitempty" name:"ActiveWeekdays"`

	// 长期保留周期，周weekly，月monthly，空不开启
	LongTermUnit *string `json:"LongTermUnit,omitnil,omitempty" name:"LongTermUnit"`

	// 长期保留哪些天的，周0-6，月1-31，逗号分割
	LongTermActiveDays *string `json:"LongTermActiveDays,omitnil,omitempty" name:"LongTermActiveDays"`

	// 长期备份保留多少天
	LongTermExpiredDays *int64 `json:"LongTermExpiredDays,omitnil,omitempty" name:"LongTermExpiredDays"`

	// 增量保留多少天
	OplogExpiredDays *int64 `json:"OplogExpiredDays,omitnil,omitempty" name:"OplogExpiredDays"`

	// 备份版本。旧版本备份为0，高级备份为1。开启高级备份此值传1
	BackupVersion *int64 `json:"BackupVersion,omitnil,omitempty" name:"BackupVersion"`

	// 告警额度。50-300
	AlarmWaterLevel *int64 `json:"AlarmWaterLevel,omitnil,omitempty" name:"AlarmWaterLevel"`
}

func (r *SetBackupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBackupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMethod")
	delete(f, "BackupTime")
	delete(f, "BackupFrequency")
	delete(f, "Notify")
	delete(f, "BackupRetentionPeriod")
	delete(f, "ActiveWeekdays")
	delete(f, "LongTermUnit")
	delete(f, "LongTermActiveDays")
	delete(f, "LongTermExpiredDays")
	delete(f, "OplogExpiredDays")
	delete(f, "BackupVersion")
	delete(f, "AlarmWaterLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetBackupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetBackupRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetBackupRulesResponse struct {
	*tchttp.BaseResponse
	Response *SetBackupRulesResponseParams `json:"Response"`
}

func (r *SetBackupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetBackupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDBInstanceDeletionProtectionRequestParams struct {
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例销毁保护选项，取值范围：0-关闭销毁保护，1-开启销毁保护
	EnableDeletionProtection *uint64 `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

type SetDBInstanceDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例销毁保护选项，取值范围：0-关闭销毁保护，1-开启销毁保护
	EnableDeletionProtection *uint64 `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

func (r *SetDBInstanceDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDBInstanceDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "EnableDeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDBInstanceDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDBInstanceDeletionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetDBInstanceDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *SetDBInstanceDeletionProtectionResponseParams `json:"Response"`
}

func (r *SetDBInstanceDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDBInstanceDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetInstanceMaintenanceRequestParams struct {
	// 指定实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护时间窗开始时间。取值范围为"00:00-23:00"的任意整点或半点，如00:00或00:30。
	MaintenanceStart *string `json:"MaintenanceStart,omitnil,omitempty" name:"MaintenanceStart"`

	// 维护时间窗结束时间。
	// - 取值范围为"00:00-23:00"的任意整点或半点，维护时间持续时长最小为30分钟，最大为3小时。
	// - 结束时间务必是基于开始时间向后的时间。
	MaintenanceEnd *string `json:"MaintenanceEnd,omitnil,omitempty" name:"MaintenanceEnd"`
}

type SetInstanceMaintenanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例ID。例如：cmgo-p8vn****。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护时间窗开始时间。取值范围为"00:00-23:00"的任意整点或半点，如00:00或00:30。
	MaintenanceStart *string `json:"MaintenanceStart,omitnil,omitempty" name:"MaintenanceStart"`

	// 维护时间窗结束时间。
	// - 取值范围为"00:00-23:00"的任意整点或半点，维护时间持续时长最小为30分钟，最大为3小时。
	// - 结束时间务必是基于开始时间向后的时间。
	MaintenanceEnd *string `json:"MaintenanceEnd,omitnil,omitempty" name:"MaintenanceEnd"`
}

func (r *SetInstanceMaintenanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetInstanceMaintenanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MaintenanceStart")
	delete(f, "MaintenanceEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetInstanceMaintenanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetInstanceMaintenanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetInstanceMaintenanceResponse struct {
	*tchttp.BaseResponse
	Response *SetInstanceMaintenanceResponseParams `json:"Response"`
}

func (r *SetInstanceMaintenanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetInstanceMaintenanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShardInfo struct {
	// 分片已使用容量
	UsedVolume *float64 `json:"UsedVolume,omitnil,omitempty" name:"UsedVolume"`

	// 分片ID
	ReplicaSetId *string `json:"ReplicaSetId,omitnil,omitempty" name:"ReplicaSetId"`

	// 分片名
	ReplicaSetName *string `json:"ReplicaSetName,omitnil,omitempty" name:"ReplicaSetName"`

	// 分片内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 分片磁盘规格，单位为MB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 分片Oplog大小，单位为MB
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`

	// 分片从节点数
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// 分片物理id
	RealReplicaSetId *string `json:"RealReplicaSetId,omitnil,omitempty" name:"RealReplicaSetId"`
}

type SlowLogItem struct {
	// 慢日志详情。
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 节点名称。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 查询哈希值。
	QueryHash *string `json:"QueryHash,omitnil,omitempty" name:"QueryHash"`
}

type SlowLogPattern struct {
	// 慢日志输出格式：库名.表名.命令。
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 记录慢日志时所带的queryHash 值，标识一类查询。
	QueryHash *string `json:"QueryHash,omitnil,omitempty" name:"QueryHash"`

	// 最大执行时间。单位：毫秒。
	MaxTime *uint64 `json:"MaxTime,omitnil,omitempty" name:"MaxTime"`

	// 平均执行时间。单位：毫秒。
	AverageTime *uint64 `json:"AverageTime,omitnil,omitempty" name:"AverageTime"`

	// 慢日志条数。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type SpecItem struct {
	// 规格信息标识。格式如：mongo.HIO10G.128g。由节点类型、规格类型、内存规格三部分组成。
	// - 节点类型：**mongo**，指 Mongod 节点；**mongos**，指 Mongos 节点；**cfgstr**，指 Configserver 节点。
	// - 规格类型：**HIO10G**，指通用高HIO万兆型；**HCD**：指云盘版类型。
	// - 内存规格：支持4、8、16、32、64、128、240、512。单位g：表示GB。128g 则表示128GB。
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 售卖规格有效标志，取值范围如下：
	// - 0：停止售卖，
	// - 1：开放售卖。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 计算资源规格，CPU核数。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存规格，单位为：MB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 默认磁盘规格，单位为：MB。
	DefaultStorage *uint64 `json:"DefaultStorage,omitnil,omitempty" name:"DefaultStorage"`

	// 最大磁盘规格，单位为：MB。
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 最小磁盘规格，单位为：MB。
	MinStorage *uint64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// 指每秒最大请求次数，单位为：次/秒。
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 规格所支持的最大连接数限制。
	Conns *uint64 `json:"Conns,omitnil,omitempty" name:"Conns"`

	// 实例存储引擎版本信息。
	// - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。
	// - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。
	// - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。
	// - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。
	// - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	// - MONGO_70_WT：MongoDB 7.0 WiredTiger存储引擎版本。
	MongoVersionCode *string `json:"MongoVersionCode,omitnil,omitempty" name:"MongoVersionCode"`

	// 实例版本对应的数字版本。
	MongoVersionValue *uint64 `json:"MongoVersionValue,omitnil,omitempty" name:"MongoVersionValue"`

	// 实例版本信息。支持：4.2、4.4、5.0、6.0、7.0。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 存储引擎。
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 集群类型，取值如下：
	// - 1：分片集群。
	// - 0：副本集集群。
	ClusterType *uint64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 每个副本集最小节点数。
	MinNodeNum *uint64 `json:"MinNodeNum,omitnil,omitempty" name:"MinNodeNum"`

	// 每个副本集最大节点数。
	MaxNodeNum *uint64 `json:"MaxNodeNum,omitnil,omitempty" name:"MaxNodeNum"`

	// 最小分片数。
	MinReplicateSetNum *uint64 `json:"MinReplicateSetNum,omitnil,omitempty" name:"MinReplicateSetNum"`

	// 最大分片数。
	MaxReplicateSetNum *uint64 `json:"MaxReplicateSetNum,omitnil,omitempty" name:"MaxReplicateSetNum"`

	// 每个分片最小节点数。
	MinReplicateSetNodeNum *uint64 `json:"MinReplicateSetNodeNum,omitnil,omitempty" name:"MinReplicateSetNodeNum"`

	// 每个分片最大节点数。
	MaxReplicateSetNodeNum *uint64 `json:"MaxReplicateSetNodeNum,omitnil,omitempty" name:"MaxReplicateSetNodeNum"`

	// 集群的规格类型，取值范围如下：
	// - HIO10G：通用高HIO万兆型。
	// - HCD：云盘版类型。
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`
}

type SpecificationInfo struct {
	// 地域信息。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区信息。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 售卖规格信息。
	SpecItems []*SpecItem `json:"SpecItems,omitnil,omitempty" name:"SpecItems"`

	// 是否支持跨可用区部署。
	// - 1：支持。
	// - 0：不支持。
	SupportMultiAZ *int64 `json:"SupportMultiAZ,omitnil,omitempty" name:"SupportMultiAZ"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Task struct {
	// 下载任务类型，0:慢日志，1:错误日志
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 任务状态，0:初始化，1:运行中，2:成功，3:失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type TerminateDBInstancesRequestParams struct {
	// 指定预隔离实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制预隔离实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TerminateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 指定预隔离实例 ID。请登录 [MongoDB 控制台](https://console.cloud.tencent.com/mongodb)在实例列表复制预隔离实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TerminateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDBInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDBInstancesResponseParams `json:"Response"`
}

func (r *TerminateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceKernelVersionRequestParams struct {
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否维护时间内升级。0-否，1-是
	InMaintenance *int64 `json:"InMaintenance,omitnil,omitempty" name:"InMaintenance"`
}

type UpgradeDBInstanceKernelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否维护时间内升级。0-否，1-是
	InMaintenance *int64 `json:"InMaintenance,omitnil,omitempty" name:"InMaintenance"`
}

func (r *UpgradeDBInstanceKernelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceKernelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InMaintenance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceKernelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceKernelVersionResponseParams struct {
	// 异步流程任务ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceKernelVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceKernelVersionResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceKernelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceKernelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDbInstanceVersionRequestParams struct {
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新升级的数据库版本，当前仅支持MONGO_40_WT（MongoDB 4.0 WiredTiger存储引擎版本）及MONGO_42_WT（MongoDB 4.0 WiredTiger存储引擎版本）。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 是否在维护时间内升级。0-立即升级 1-维护时间内升级
	InMaintenance *int64 `json:"InMaintenance,omitnil,omitempty" name:"InMaintenance"`
}

type UpgradeDbInstanceVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新升级的数据库版本，当前仅支持MONGO_40_WT（MongoDB 4.0 WiredTiger存储引擎版本）及MONGO_42_WT（MongoDB 4.0 WiredTiger存储引擎版本）。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 是否在维护时间内升级。0-立即升级 1-维护时间内升级
	InMaintenance *int64 `json:"InMaintenance,omitnil,omitempty" name:"InMaintenance"`
}

func (r *UpgradeDbInstanceVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDbInstanceVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MongoVersion")
	delete(f, "InMaintenance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDbInstanceVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDbInstanceVersionResponseParams struct {
	// 异步流程任务ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDbInstanceVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDbInstanceVersionResponseParams `json:"Response"`
}

func (r *UpgradeDbInstanceVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDbInstanceVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 账号名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号权限详情。
	AuthRole []*Auth `json:"AuthRole,omitnil,omitempty" name:"AuthRole"`

	// 账号创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 账号更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 备注信息。
	UserDesc *string `json:"UserDesc,omitnil,omitempty" name:"UserDesc"`

	// 控制台密码更新时间
	ConsolePassUpdateTime *string `json:"ConsolePassUpdateTime,omitnil,omitempty" name:"ConsolePassUpdateTime"`
}