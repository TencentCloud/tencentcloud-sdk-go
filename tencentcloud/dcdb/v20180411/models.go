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

package v20180411

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Account struct {
	// 账户的名称
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 账户的域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type ActiveHourDCDBInstanceRequestParams struct {
	// 待升级的实例ID列表。形如：["dcdbt-ow728lmc"]，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ActiveHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待升级的实例ID列表。形如：["dcdbt-ow728lmc"]，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *ActiveHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActiveHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActiveHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActiveHourDCDBInstanceResponseParams struct {
	// 解隔离成功的实例id列表
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// 解隔离失败的实例id列表
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActiveHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ActiveHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *ActiveHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActiveHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddShardConfig struct {
	// 新增分片的数量
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 要绑定的安全组ID，类似sg-efil73jd。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 被绑定的实例ID，类似tdsqlshard-lesecurk，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 要绑定的安全组ID，类似sg-efil73jd。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 被绑定的实例ID，类似tdsqlshard-lesecurk，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupConfig struct {
	// 备份策略是否启用。
	EnableBackupPolicy *bool `json:"EnableBackupPolicy,omitnil,omitempty" name:"EnableBackupPolicy"`

	// 超期保留开始日期，早于开始日期的超期备份不保留，格式：yyyy-mm-dd。
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 超期备份保留时长，超出保留时间的超期备份将被删除，可填写1-3650整数。
	MaxRetentionDays *int64 `json:"MaxRetentionDays,omitnil,omitempty" name:"MaxRetentionDays"`

	// 备份模式，可选择按年月周模式保存
	// * 按年：annually
	// * 按月：monthly
	// * 按周：weekly
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// Frequency等于weekly时生效。
	// 表示保留特定工作日备份。可选择周一到周日，支持多选，取星期英文：
	// * 星期一 ：Monday
	// * 星期二 ：Tuesday
	// * 星期三：Wednesday
	// * 星期四：Thursday
	// * 星期五：Friday
	// * 星期六：Saturday
	// * 星期日：Sunday
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 保留备份个数，Frequency等于monthly或weekly时生效。
	// 备份模式选择按月时，可填写1-28整数；
	// 备份模式选择年时，可填写1-336整数。
	BackupCount *int64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`
}

type BriefNodeInfo struct {
	// DB节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// DB节点角色，取值为master或者slave
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 节点所属分片的分片ID
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`
}

// Predefined struct for user
type CancelDcnJobRequestParams struct {
	// 灾备实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CancelDcnJobRequest struct {
	*tchttp.BaseRequest
	
	// 灾备实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CancelDcnJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDcnJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDcnJobResponseParams struct {
	// 流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelDcnJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelDcnJobResponseParams `json:"Response"`
}

func (r *CancelDcnJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneAccountRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 源用户账户名
	SrcUser *string `json:"SrcUser,omitnil,omitempty" name:"SrcUser"`

	// 源用户HOST
	SrcHost *string `json:"SrcHost,omitnil,omitempty" name:"SrcHost"`

	// 目的用户账户名
	DstUser *string `json:"DstUser,omitnil,omitempty" name:"DstUser"`

	// 目的用户HOST
	DstHost *string `json:"DstHost,omitnil,omitempty" name:"DstHost"`

	// 目的用户账户描述
	DstDesc *string `json:"DstDesc,omitnil,omitempty" name:"DstDesc"`
}

type CloneAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 源用户账户名
	SrcUser *string `json:"SrcUser,omitnil,omitempty" name:"SrcUser"`

	// 源用户HOST
	SrcHost *string `json:"SrcHost,omitnil,omitempty" name:"SrcHost"`

	// 目的用户账户名
	DstUser *string `json:"DstUser,omitnil,omitempty" name:"DstUser"`

	// 目的用户HOST
	DstHost *string `json:"DstHost,omitnil,omitempty" name:"DstHost"`

	// 目的用户账户描述
	DstDesc *string `json:"DstDesc,omitnil,omitempty" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUser")
	delete(f, "SrcHost")
	delete(f, "DstUser")
	delete(f, "DstHost")
	delete(f, "DstDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneAccountResponseParams struct {
	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse
	Response *CloneAccountResponseParams `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessRequestParams struct {
	// 待关闭外网访问的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否IPv6，默认0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 待关闭外网访问的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否IPv6，默认0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ipv6Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessResponseParams struct {
	// 异步任务ID，可通过 DescribeFlow 查询任务状态。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *CloseDBExtranetAccessResponseParams `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColumnPrivilege struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 数据库列名
	Column *string `json:"Column,omitnil,omitempty" name:"Column"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type ConfigValue struct {
	// 配置项的名称，支持填写max_user_connections
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 配置值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ConstraintRange struct {
	// 约束类型为section时的最小值
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 约束类型为section时的最大值
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`
}

// Predefined struct for user
type CopyAccountPrivilegesRequestParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 源用户名
	SrcUserName *string `json:"SrcUserName,omitnil,omitempty" name:"SrcUserName"`

	// 源用户允许的访问 host
	SrcHost *string `json:"SrcHost,omitnil,omitempty" name:"SrcHost"`

	// 目的用户名
	DstUserName *string `json:"DstUserName,omitnil,omitempty" name:"DstUserName"`

	// 目的用户允许的访问 host
	DstHost *string `json:"DstHost,omitnil,omitempty" name:"DstHost"`

	// 源账号的 ReadOnly 属性
	SrcReadOnly *string `json:"SrcReadOnly,omitnil,omitempty" name:"SrcReadOnly"`

	// 目的账号的 ReadOnly 属性
	DstReadOnly *string `json:"DstReadOnly,omitnil,omitempty" name:"DstReadOnly"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 源用户名
	SrcUserName *string `json:"SrcUserName,omitnil,omitempty" name:"SrcUserName"`

	// 源用户允许的访问 host
	SrcHost *string `json:"SrcHost,omitnil,omitempty" name:"SrcHost"`

	// 目的用户名
	DstUserName *string `json:"DstUserName,omitnil,omitempty" name:"DstUserName"`

	// 目的用户允许的访问 host
	DstHost *string `json:"DstHost,omitnil,omitempty" name:"DstHost"`

	// 源账号的 ReadOnly 属性
	SrcReadOnly *string `json:"SrcReadOnly,omitnil,omitempty" name:"SrcReadOnly"`

	// 目的账号的 ReadOnly 属性
	DstReadOnly *string `json:"DstReadOnly,omitnil,omitempty" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUserName")
	delete(f, "SrcHost")
	delete(f, "DstUserName")
	delete(f, "DstHost")
	delete(f, "SrcReadOnly")
	delete(f, "DstReadOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAccountPrivilegesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *CopyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// AccountName
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 可以登录的主机，与mysql 账号的 host 格式一致，可以支持通配符，例如 %，10.%，10.20.%。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号密码，密码需要 8-32 个字符，不能以 '/' 开头，并且必须包含小写字母、大写字母、数字和符号()~!@#$%^&*-+=_|{}[]:<>,.?/。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否创建为只读账号，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败，3：只从备机读取。
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 账号备注，可以包含中文、英文字符、常见符号和数字，长度为0~256字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `json:"DelayThresh,omitnil,omitempty" name:"DelayThresh"`

	// 针对只读账号，设置策略是否固定备机，0：不固定备机，即备机不满足条件与客户端不断开连接，Proxy选择其他可用备机，1：备机不满足条件断开连接，确保一个连接固定备机。
	SlaveConst *int64 `json:"SlaveConst,omitnil,omitempty" name:"SlaveConst"`

	// 用户最大连接数限制参数。不传或者传0表示为不限制，对应max_user_connections参数，目前10.1内核版本不支持设置。
	MaxUserConnections *uint64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`

	// 使用GetPublicKey返回的RSA2048公钥加密后的密码
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// AccountName
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 可以登录的主机，与mysql 账号的 host 格式一致，可以支持通配符，例如 %，10.%，10.20.%。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 账号密码，密码需要 8-32 个字符，不能以 '/' 开头，并且必须包含小写字母、大写字母、数字和符号()~!@#$%^&*-+=_|{}[]:<>,.?/。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否创建为只读账号，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败，3：只从备机读取。
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 账号备注，可以包含中文、英文字符、常见符号和数字，长度为0~256字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `json:"DelayThresh,omitnil,omitempty" name:"DelayThresh"`

	// 针对只读账号，设置策略是否固定备机，0：不固定备机，即备机不满足条件与客户端不断开连接，Proxy选择其他可用备机，1：备机不满足条件断开连接，确保一个连接固定备机。
	SlaveConst *int64 `json:"SlaveConst,omitnil,omitempty" name:"SlaveConst"`

	// 用户最大连接数限制参数。不传或者传0表示为不限制，对应max_user_connections参数，目前10.1内核版本不支持设置。
	MaxUserConnections *uint64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`

	// 使用GetPublicKey返回的RSA2048公钥加密后的密码
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "ReadOnly")
	delete(f, "Description")
	delete(f, "DelayThresh")
	delete(f, "SlaveConst")
	delete(f, "MaxUserConnections")
	delete(f, "EncryptedPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// 实例ID，透传入参。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，透传入参。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 允许访问的 host，透传入参。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 透传入参。
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDCDBInstanceRequestParams struct {
	// 分片节点可用区分布，可填写多个可用区。
	// 注意当前可售卖的可用区需要通过DescribeDCDBSaleInfo接口拉取。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil,omitempty" name:"ShardNodeCount"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 欲购买实例的数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 数据库引擎版本，当前可选：8.0，5.7，10.1，10.0。
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例名称， 可以通过该字段自主的设置实例的名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 是否支持IPv6，0:不支持，1:支持
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// 标签键值对数组
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。
	InitParams []*DBParamValue `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// DCN源地域
	DcnRegion *string `json:"DcnRegion,omitnil,omitempty" name:"DcnRegion"`

	// DCN源实例ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil,omitempty" name:"DcnInstanceId"`

	// 自动续费标记，0:默认状态(用户未设置，即初始状态即手动续费，用户开通了预付费不停服特权也会进行自动续费)， 1:自动续费，2:明确不自动续费(用户设置)。若业务无续费概念或无需自动续费，需要设置为0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 安全组ids，安全组可以传数组形式，兼容之前SecurityGroupId参数
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// DCN同步模式，0：异步， 1：强同步 
	DcnSyncMode *int64 `json:"DcnSyncMode,omitnil,omitempty" name:"DcnSyncMode"`
}

type CreateDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 分片节点可用区分布，可填写多个可用区。
	// 注意当前可售卖的可用区需要通过DescribeDCDBSaleInfo接口拉取。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil,omitempty" name:"ShardNodeCount"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 欲购买实例的数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 数据库引擎版本，当前可选：8.0，5.7，10.1，10.0。
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例名称， 可以通过该字段自主的设置实例的名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 是否支持IPv6，0:不支持，1:支持
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// 标签键值对数组
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。
	InitParams []*DBParamValue `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// DCN源地域
	DcnRegion *string `json:"DcnRegion,omitnil,omitempty" name:"DcnRegion"`

	// DCN源实例ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil,omitempty" name:"DcnInstanceId"`

	// 自动续费标记，0:默认状态(用户未设置，即初始状态即手动续费，用户开通了预付费不停服特权也会进行自动续费)， 1:自动续费，2:明确不自动续费(用户设置)。若业务无续费概念或无需自动续费，需要设置为0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 安全组ids，安全组可以传数组形式，兼容之前SecurityGroupId参数
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// DCN同步模式，0：异步， 1：强同步 
	DcnSyncMode *int64 `json:"DcnSyncMode,omitnil,omitempty" name:"DcnSyncMode"`
}

func (r *CreateDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zones")
	delete(f, "Period")
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardNodeCount")
	delete(f, "ShardCount")
	delete(f, "Count")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbVersionId")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceName")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	delete(f, "InitParams")
	delete(f, "DcnRegion")
	delete(f, "DcnInstanceId")
	delete(f, "AutoRenewFlag")
	delete(f, "SecurityGroupIds")
	delete(f, "DcnSyncMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDCDBInstanceResponseParams struct {
	// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDCDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterDCDBInstanceRequestParams struct {
	// 分配实例个数
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 分片数量
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 分片內存大小, 单位GB
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片磁盘大小, 单位GB
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 独享集群集群uuid
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// （废弃）可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// （废弃）cpu大小，单位：核
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// （废弃）分片机型
	ShardMachine *string `json:"ShardMachine,omitnil,omitempty" name:"ShardMachine"`

	// 分片的节点个数
	ShardNodeNum *int64 `json:"ShardNodeNum,omitnil,omitempty" name:"ShardNodeNum"`

	// （废弃）节点cpu核数，单位：1/100核
	ShardNodeCpu *int64 `json:"ShardNodeCpu,omitnil,omitempty" name:"ShardNodeCpu"`

	// （废弃）节点內存大小，单位：GB
	ShardNodeMemory *int64 `json:"ShardNodeMemory,omitnil,omitempty" name:"ShardNodeMemory"`

	// （废弃）节点磁盘大小，单位：GB
	ShardNodeStorage *int64 `json:"ShardNodeStorage,omitnil,omitempty" name:"ShardNodeStorage"`

	// db版本
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组ID列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// DCN源实例ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil,omitempty" name:"DcnInstanceId"`

	// DCN源实例地域名
	DcnRegion *string `json:"DcnRegion,omitnil,omitempty" name:"DcnRegion"`

	// 自定义实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 标签
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 支持IPv6标志：1 支持， 0 不支持
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// （废弃）Pid，可通过获取独享集群售卖配置接口得到
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。
	InitParams []*DBParamValue `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// 指定主节点uuid，不填随机分配
	MasterHostId *string `json:"MasterHostId,omitnil,omitempty" name:"MasterHostId"`

	// 指定从节点uuid，不填随机分配
	SlaveHostIds []*string `json:"SlaveHostIds,omitnil,omitempty" name:"SlaveHostIds"`

	// 需要回档的源实例ID
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil,omitempty" name:"RollbackInstanceId"`

	// 回档时间
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`

	// DCN同步模式，0：异步， 1：强同步
	DcnSyncMode *int64 `json:"DcnSyncMode,omitnil,omitempty" name:"DcnSyncMode"`
}

type CreateDedicatedClusterDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 分配实例个数
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 分片数量
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 分片內存大小, 单位GB
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片磁盘大小, 单位GB
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 独享集群集群uuid
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// （废弃）可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// （废弃）cpu大小，单位：核
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// （废弃）分片机型
	ShardMachine *string `json:"ShardMachine,omitnil,omitempty" name:"ShardMachine"`

	// 分片的节点个数
	ShardNodeNum *int64 `json:"ShardNodeNum,omitnil,omitempty" name:"ShardNodeNum"`

	// （废弃）节点cpu核数，单位：1/100核
	ShardNodeCpu *int64 `json:"ShardNodeCpu,omitnil,omitempty" name:"ShardNodeCpu"`

	// （废弃）节点內存大小，单位：GB
	ShardNodeMemory *int64 `json:"ShardNodeMemory,omitnil,omitempty" name:"ShardNodeMemory"`

	// （废弃）节点磁盘大小，单位：GB
	ShardNodeStorage *int64 `json:"ShardNodeStorage,omitnil,omitempty" name:"ShardNodeStorage"`

	// db版本
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组ID列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// DCN源实例ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil,omitempty" name:"DcnInstanceId"`

	// DCN源实例地域名
	DcnRegion *string `json:"DcnRegion,omitnil,omitempty" name:"DcnRegion"`

	// 自定义实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 标签
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 支持IPv6标志：1 支持， 0 不支持
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// （废弃）Pid，可通过获取独享集群售卖配置接口得到
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。
	InitParams []*DBParamValue `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// 指定主节点uuid，不填随机分配
	MasterHostId *string `json:"MasterHostId,omitnil,omitempty" name:"MasterHostId"`

	// 指定从节点uuid，不填随机分配
	SlaveHostIds []*string `json:"SlaveHostIds,omitnil,omitempty" name:"SlaveHostIds"`

	// 需要回档的源实例ID
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil,omitempty" name:"RollbackInstanceId"`

	// 回档时间
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`

	// DCN同步模式，0：异步， 1：强同步
	DcnSyncMode *int64 `json:"DcnSyncMode,omitnil,omitempty" name:"DcnSyncMode"`
}

func (r *CreateDedicatedClusterDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GoodsNum")
	delete(f, "ShardNum")
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ClusterId")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "Cpu")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ShardMachine")
	delete(f, "ShardNodeNum")
	delete(f, "ShardNodeCpu")
	delete(f, "ShardNodeMemory")
	delete(f, "ShardNodeStorage")
	delete(f, "DbVersionId")
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupIds")
	delete(f, "DcnInstanceId")
	delete(f, "DcnRegion")
	delete(f, "InstanceName")
	delete(f, "ResourceTags")
	delete(f, "Ipv6Flag")
	delete(f, "Pid")
	delete(f, "InitParams")
	delete(f, "MasterHostId")
	delete(f, "SlaveHostIds")
	delete(f, "RollbackInstanceId")
	delete(f, "RollbackTime")
	delete(f, "DcnSyncMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDedicatedClusterDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDedicatedClusterDCDBInstanceResponseParams struct {
	// 分配资源ID数组
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDedicatedClusterDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDedicatedClusterDCDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDedicatedClusterDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDedicatedClusterDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHourDCDBInstanceRequestParams struct {
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil,omitempty" name:"ShardNodeCount"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 欲购买实例的数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 分片cpu大小，单位：核，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardCpu *int64 `json:"ShardCpu,omitnil,omitempty" name:"ShardCpu"`

	// 数据库引擎版本，当前可选：8.0，5.7，10.1，10.0。
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`

	// 分片节点可用区分布，可填写多个可用区。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例名称， 可以通过该字段自主的设置实例的名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 是否支持IPv6，0:不支持，1:支持
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// 标签键值对数组
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// DCN源地域
	DcnRegion *string `json:"DcnRegion,omitnil,omitempty" name:"DcnRegion"`

	// DCN源实例ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil,omitempty" name:"DcnInstanceId"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。
	InitParams []*DBParamValue `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// 需要回档的源实例ID
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil,omitempty" name:"RollbackInstanceId"`

	// 回档时间，例如“2021-11-22 00:00:00”
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`

	// 安全组ids，安全组可以传数组形式，兼容之前SecurityGroupId参数
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// DCN同步模式，0：异步， 1：强同步
	DcnSyncMode *int64 `json:"DcnSyncMode,omitnil,omitempty" name:"DcnSyncMode"`
}

type CreateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 单个分片节点个数，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil,omitempty" name:"ShardNodeCount"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 欲购买实例的数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 分片cpu大小，单位：核，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardCpu *int64 `json:"ShardCpu,omitnil,omitempty" name:"ShardCpu"`

	// 数据库引擎版本，当前可选：8.0，5.7，10.1，10.0。
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`

	// 分片节点可用区分布，可填写多个可用区。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例名称， 可以通过该字段自主的设置实例的名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 是否支持IPv6，0:不支持，1:支持
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// 标签键值对数组
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// DCN源地域
	DcnRegion *string `json:"DcnRegion,omitnil,omitempty" name:"DcnRegion"`

	// DCN源实例ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil,omitempty" name:"DcnInstanceId"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步可退化）。
	InitParams []*DBParamValue `json:"InitParams,omitnil,omitempty" name:"InitParams"`

	// 需要回档的源实例ID
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil,omitempty" name:"RollbackInstanceId"`

	// 回档时间，例如“2021-11-22 00:00:00”
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`

	// 安全组ids，安全组可以传数组形式，兼容之前SecurityGroupId参数
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// DCN同步模式，0：异步， 1：强同步
	DcnSyncMode *int64 `json:"DcnSyncMode,omitnil,omitempty" name:"DcnSyncMode"`
}

func (r *CreateHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardNodeCount")
	delete(f, "ShardCount")
	delete(f, "Count")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ShardCpu")
	delete(f, "DbVersionId")
	delete(f, "Zones")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceName")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	delete(f, "DcnRegion")
	delete(f, "DcnInstanceId")
	delete(f, "InitParams")
	delete(f, "RollbackInstanceId")
	delete(f, "RollbackTime")
	delete(f, "SecurityGroupIds")
	delete(f, "DcnSyncMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHourDCDBInstanceResponseParams struct {
	// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 流程id，可以根据流程id查询创建进度
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *CreateHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTmpDCDBInstanceRequestParams struct {
	// 回档实例的ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 回档时间点
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`
}

type CreateTmpDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 回档实例的ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 回档时间点
	RollbackTime *string `json:"RollbackTime,omitnil,omitempty" name:"RollbackTime"`
}

func (r *CreateTmpDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTmpDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RollbackTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTmpDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTmpDCDBInstanceResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTmpDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateTmpDCDBInstanceResponseParams `json:"Response"`
}

func (r *CreateTmpDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTmpDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户可以从哪台主机登录（对应 MySQL 用户的 host 字段，UserName + Host 唯一标识一个用户，IP形式，IP段以%结尾；支持填入%；为空默认等于%）
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 用户备注信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 只读标记，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败。
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障
	// 建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `json:"DelayThresh,omitnil,omitempty" name:"DelayThresh"`

	// 针对只读账号，设置策略是否固定备机，0：不固定备机，即备机不满足条件与客户端不断开连接，Proxy选择其他可用备机，1：备机不满足条件断开连接，确保一个连接固定备机。
	SlaveConst *int64 `json:"SlaveConst,omitnil,omitempty" name:"SlaveConst"`

	// 用户最大连接数，0代表无限制	
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil,omitempty" name:"MaxUserConnections"`
}

type DBParamValue struct {
	// 参数名称
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DCDBInstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 应用ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC数字ID
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet数字ID
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 状态中文描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 实例状态：0 创建中，1 流程处理中， 2 运行中，3 实例未初始化，-1 实例已隔离，4 实例初始化中，5 实例删除中，6 实例重启中，7 数据迁移中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 内网IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 自动续费标志
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 分片个数
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 隔离时间
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitnil,omitempty" name:"IsolatedTimestamp"`

	// 账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 分片详情
	ShardDetail []*ShardInfo `json:"ShardDetail,omitnil,omitempty" name:"ShardDetail"`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 临时实例标记，0 为非临时实例
	IsTmp *int64 `json:"IsTmp,omitnil,omitempty" name:"IsTmp"`

	// 独享集群ID，为空表示非独享集群实例
	ExclusterId *string `json:"ExclusterId,omitnil,omitempty" name:"ExclusterId"`

	// 字符串型的私有网络ID
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 字符串型的私有网络子网ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 数字实例ID（过时字段，请勿依赖该值）
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 外网访问的域名，公网可解析
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网 IP 地址，公网可访问
	WanVip *string `json:"WanVip,omitnil,omitempty" name:"WanVip"`

	// 外网端口
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 产品类型 ID（过时字段，请勿依赖该值）
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 实例最后更新时间，格式为 2006-01-02 15:04:05
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 数据库引擎
	DbEngine *string `json:"DbEngine,omitnil,omitempty" name:"DbEngine"`

	// 数据库引擎版本
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 付费模式
	Paymode *string `json:"Paymode,omitnil,omitempty" name:"Paymode"`

	// 实例处于异步任务状态时，表示异步任务流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locker *int64 `json:"Locker,omitnil,omitempty" name:"Locker"`

	// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中
	WanStatus *int64 `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 该实例是否支持审计。1-支持；0-不支持
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitnil,omitempty" name:"IsAuditSupported"`

	// Cpu核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例IPv6标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// 内网IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vipv6 *string `json:"Vipv6,omitnil,omitempty" name:"Vipv6"`

	// 外网IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanVipv6 *string `json:"WanVipv6,omitnil,omitempty" name:"WanVipv6"`

	// 外网IPv6端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitnil,omitempty" name:"WanPortIpv6"`

	// 外网IPv6状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitnil,omitempty" name:"WanStatusIpv6"`

	// DCN标志，0-无，1-主实例，2-灾备实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnFlag *int64 `json:"DcnFlag,omitnil,omitempty" name:"DcnFlag"`

	// DCN状态，0-无，1-创建中，2-同步中，3-已断开
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnStatus *int64 `json:"DcnStatus,omitnil,omitempty" name:"DcnStatus"`

	// DCN灾备实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnDstNum *int64 `json:"DcnDstNum,omitnil,omitempty" name:"DcnDstNum"`

	// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 数据库引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`
}

type DCDBShardInfo struct {
	// 所属实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分片SQL透传Id，用于将sql透传到指定分片执行
	ShardSerialId *string `json:"ShardSerialId,omitnil,omitempty" name:"ShardSerialId"`

	// 全局唯一的分片Id
	ShardInstanceId *string `json:"ShardInstanceId,omitnil,omitempty" name:"ShardInstanceId"`

	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态中文描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 字符串格式的私有网络Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 字符串格式的私有网络子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 到期时间
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 存储使用率，单位为 %
	StorageUsage *float64 `json:"StorageUsage,omitnil,omitempty" name:"StorageUsage"`

	// 内存使用率，单位为 %
	MemoryUsage *float64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// 数字分片Id（过时字段，请勿依赖该值）
	ShardId *int64 `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 产品ProductID
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Proxy版本
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// 付费模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paymode *string `json:"Paymode,omitnil,omitempty" name:"Paymode"`

	// 分片的主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardMasterZone *string `json:"ShardMasterZone,omitnil,omitempty" name:"ShardMasterZone"`

	// 分片的从可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardSlaveZones []*string `json:"ShardSlaveZones,omitnil,omitempty" name:"ShardSlaveZones"`

	// CPU核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 分片ShardKey的范围（总共64个哈希值），例如： 0-31，32-63
	Range *string `json:"Range,omitnil,omitempty" name:"Range"`
}

type Database struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DatabaseFunction struct {
	// 函数名称
	Func *string `json:"Func,omitnil,omitempty" name:"Func"`
}

type DatabasePrivilege struct {
	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`
}

type DatabaseProcedure struct {
	// 存储过程名称
	Proc *string `json:"Proc,omitnil,omitempty" name:"Proc"`
}

type DatabaseTable struct {
	// 表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type DatabaseView struct {
	// 视图名称
	View *string `json:"View,omitnil,omitempty" name:"View"`
}

type DcnDetailItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例IP地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 实例IPv6地址
	Vipv6 *string `json:"Vipv6,omitnil,omitempty" name:"Vipv6"`

	// 实例端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 实例DCN标志，1-主，2-备
	DcnFlag *int64 `json:"DcnFlag,omitnil,omitempty" name:"DcnFlag"`

	// 实例DCN状态，0-无，1-创建中，2-同步中，3-已断开
	DcnStatus *int64 `json:"DcnStatus,omitnil,omitempty" name:"DcnStatus"`

	// 实例CPU核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例内存大小，单位 GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例存储大小，单位 GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 付费模式
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例创建时间，格式为 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例到期时间，格式为 2006-01-02 15:04:05
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 是否开启了 kms
	EncryptStatus *int64 `json:"EncryptStatus,omitnil,omitempty" name:"EncryptStatus"`

	// 实例DCN状态描述信息
	DcnStatusDesc *string `json:"DcnStatusDesc,omitnil,omitempty" name:"DcnStatusDesc"`

	// DCN实例绑定的北极星服务所属的北极星实例Id，若未绑定则为空
	PolarisInstanceId *string `json:"PolarisInstanceId,omitnil,omitempty" name:"PolarisInstanceId"`

	// DCN实例绑定的北极星服务所属的北极星实例名，若未绑定则为空
	PolarisInstanceName *string `json:"PolarisInstanceName,omitnil,omitempty" name:"PolarisInstanceName"`

	// DCN实例绑定的北极星服务所属的北极星命名空间，若未绑定则为空
	PolarisNamespace *string `json:"PolarisNamespace,omitnil,omitempty" name:"PolarisNamespace"`

	// DCN实例绑定的北极星服务，若未绑定则为空
	PolarisService *string `json:"PolarisService,omitnil,omitempty" name:"PolarisService"`

	// DCN实例在北极星服务中的状态 0:未开启; 1:已开启; 2:已隔离; 3:切换中
	PolarisServiceStatus *int64 `json:"PolarisServiceStatus,omitnil,omitempty" name:"PolarisServiceStatus"`

	// DCN实例在北极星服务中的状态的描述信息
	PolarisServiceStatusDesc *string `json:"PolarisServiceStatusDesc,omitnil,omitempty" name:"PolarisServiceStatusDesc"`

	// 北极星管控地域
	PolarisRegion *string `json:"PolarisRegion,omitnil,omitempty" name:"PolarisRegion"`

	// 是否支持DCN切换
	IsDcnSwitchSupported *int64 `json:"IsDcnSwitchSupported,omitnil,omitempty" name:"IsDcnSwitchSupported"`
}

type Deal struct {
	// 订单号
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 所属账号
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 商品数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 关联的流程 Id，可用于查询流程执行状态
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 只有创建实例且已完成发货的订单会填充该字段，表示该订单创建的实例的 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 付费模式，0后付费/1预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// 实例ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 类型,可以填入 table 、 view 、 proc 、 func 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示查询该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 权限列表。
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 数据库账号用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 数据库账号Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 实例ID，透传入参。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例用户列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users []*DBAccount `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigsRequestParams struct {
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackupConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupConfigsResponseParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 常规备份存储时长，范围[1, 3650]。
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 每天备份执行的区间的开始时间，格式 mm:ss，形如 22:00。
	StartBackupTime *string `json:"StartBackupTime,omitnil,omitempty" name:"StartBackupTime"`

	// 每天备份执行的区间的结束时间，格式 mm:ss，形如 23:59。
	EndBackupTime *string `json:"EndBackupTime,omitnil,omitempty" name:"EndBackupTime"`

	// 执行备份周期，枚举值：Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 沉降到归档存储时长，-1表示关闭归档设置。
	ArchiveDays *int64 `json:"ArchiveDays,omitnil,omitempty" name:"ArchiveDays"`

	// 超期备份配置。
	BackupConfigSet []*BackupConfig `json:"BackupConfigSet,omitnil,omitempty" name:"BackupConfigSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupConfigsResponseParams `json:"Response"`
}

func (r *DescribeBackupConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesRequestParams struct {
	// 按实例ID查询
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按分片ID查询
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 备份类型，Data:数据备份，Binlog:Binlog备份，Errlog:错误日志，Slowlog:慢日志
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 按开始时间查询
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 按结束时间查询
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序参数，可选值：Time,Size
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序参数，可选值：DESC,ASC
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeBackupFilesRequest struct {
	*tchttp.BaseRequest
	
	// 按实例ID查询
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按分片ID查询
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 备份类型，Data:数据备份，Binlog:Binlog备份，Errlog:错误日志，Slowlog:慢日志
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 按开始时间查询
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 按结束时间查询
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序参数，可选值：Time,Size
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序参数，可选值：DESC,ASC
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *DescribeBackupFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardId")
	delete(f, "BackupType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesResponseParams struct {
	// 备份文件列表
	Files []*InstanceBackupFileItem `json:"Files,omitnil,omitempty" name:"Files"`

	// 总条目数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupFilesResponseParams `json:"Response"`
}

func (r *DescribeBackupFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBEncryptAttributesRequestParams struct {
	// 实例Id，形如：tdsqlshard-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id，形如：tdsqlshard-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBEncryptAttributesResponseParams struct {
	// 是否启用加密，1-已开启；0-未开启。
	EncryptStatus *int64 `json:"EncryptStatus,omitnil,omitempty" name:"EncryptStatus"`

	// DEK密钥
	CipherText *string `json:"CipherText,omitnil,omitempty" name:"CipherText"`

	// DEK密钥过期日期。
	ExpireDate *string `json:"ExpireDate,omitnil,omitempty" name:"ExpireDate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBEncryptAttributesResponseParams `json:"Response"`
}

func (r *DescribeDBEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBLogFilesRequestParams struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分片 ID，形如：shard-7noic7tv
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分片 ID，形如：shard-7noic7tv
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBLogFilesResponseParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 请求日志类型。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 请求日志总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 日志文件列表
	Files []*LogFileInfo `json:"Files,omitnil,omitempty" name:"Files"`

	// 如果是VPC网络的实例，做用本前缀加上URI为下载地址
	VpcPrefix *string `json:"VpcPrefix,omitnil,omitempty" name:"VpcPrefix"`

	// 如果是普通网络的实例，做用本前缀加上URI为下载地址
	NormalPrefix *string `json:"NormalPrefix,omitnil,omitempty" name:"NormalPrefix"`

	// 分片 ID，形如：shard-7noic7tv
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBLogFilesResponseParams `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersRequestParams struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersResponseParams struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 请求DB的当前参数值
	Params []*ParamDesc `json:"Params,omitnil,omitempty" name:"Params"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBParametersResponseParams `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组详情。
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 实例VIP
	// 注意：此字段可能返回 null，表示取不到有效值。
	VIP *string `json:"VIP,omitnil,omitempty" name:"VIP"`

	// 实例端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	VPort *string `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSlowLogsRequestParams struct {
	// 实例 ID，形如：dcdbt-hw0qj6m1
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 从结果的第几条数据开始返回
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回的结果条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的起始时间，形如2016-07-23 14:55:20
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 实例的分片ID，形如shard-53ima8ln
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 查询的结束时间，形如2016-08-22 14:55:20。如果不填，那么查询结束时间就是当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 要查询的具体数据库名称
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 排序指标，取值为query_time_sum或者query_count。不填默认按照query_time_sum排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，desc（降序）或者asc（升序）。不填默认desc排序
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 是否查询从机的慢查询，0-主机; 1-从机。不填默认查询主机慢查询
	Slave *int64 `json:"Slave,omitnil,omitempty" name:"Slave"`
}

type DescribeDBSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-hw0qj6m1
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 从结果的第几条数据开始返回
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回的结果条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询的起始时间，形如2016-07-23 14:55:20
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 实例的分片ID，形如shard-53ima8ln
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 查询的结束时间，形如2016-08-22 14:55:20。如果不填，那么查询结束时间就是当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 要查询的具体数据库名称
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 排序指标，取值为query_time_sum或者query_count。不填默认按照query_time_sum排序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型，desc（降序）或者asc（升序）。不填默认desc排序
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 是否查询从机的慢查询，0-主机; 1-从机。不填默认查询主机慢查询
	Slave *int64 `json:"Slave,omitnil,omitempty" name:"Slave"`
}

func (r *DescribeDBSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "ShardId")
	delete(f, "EndTime")
	delete(f, "Db")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Slave")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSlowLogsResponseParams struct {
	// 所有语句锁时间总和
	LockTimeSum *float64 `json:"LockTimeSum,omitnil,omitempty" name:"LockTimeSum"`

	// 所有语句查询总次数
	QueryCount *int64 `json:"QueryCount,omitnil,omitempty" name:"QueryCount"`

	// 总记录数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 所有语句查询时间总和
	QueryTimeSum *float64 `json:"QueryTimeSum,omitnil,omitempty" name:"QueryTimeSum"`

	// 慢查询日志数据
	Data []*SlowLogData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeDBSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSyncModeRequestParams struct {
	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSyncModeRequest struct {
	*tchttp.BaseRequest
	
	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSyncModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSyncModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSyncModeResponseParams struct {
	// 同步模式：0 异步，1 强同步， 2 强同步可退化
	SyncMode *int64 `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 是否有修改流程在执行中：1 是， 0 否。
	IsModifying *int64 `json:"IsModifying,omitnil,omitempty" name:"IsModifying"`

	// 当前复制方式，0 异步，1 同步
	CurrentSyncMode *int64 `json:"CurrentSyncMode,omitnil,omitempty" name:"CurrentSyncMode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSyncModeResponseParams `json:"Response"`
}

func (r *DescribeDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBTmpInstancesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBTmpInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBTmpInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBTmpInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBTmpInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBTmpInstancesResponseParams struct {
	// 临时实例列表
	TmpInstances []*TmpInstance `json:"TmpInstances,omitnil,omitempty" name:"TmpInstances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBTmpInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBTmpInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBTmpInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBTmpInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBBinlogTimeRequestParams struct {
	// 需要回档的实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDCDBBinlogTimeRequest struct {
	*tchttp.BaseRequest
	
	// 需要回档的实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDCDBBinlogTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBBinlogTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBBinlogTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBBinlogTimeResponseParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBBinlogTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBBinlogTimeResponseParams `json:"Response"`
}

func (r *DescribeDCDBBinlogTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBBinlogTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceDetailRequestParams struct {
	// 实例ID，形如dcdbt-7oaxtcb7
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDCDBInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如dcdbt-7oaxtcb7
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDCDBInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceDetailResponseParams struct {
	// 实例ID，形如dcdbt-7oaxtcb7
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例状态。0-实例创建中；1-异步任务处理中；2-运行中；3-实例未初始化；-1-实例已隔离
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例目前运行状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 实例内网IP地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 实例内网端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例节点数。值为2时表示一主一从，值为3时表示一主二从
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 实例所在地域名称，形如ap-guangzhou
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例私有网络ID，形如vpc-r9jr0de3
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 实例私有网络子网ID，形如subnet-6rqs61o2
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中；4-关闭中
	WanStatus *int64 `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 外网访问的域名，公网可解析
	WanDomain *string `json:"WanDomain,omitnil,omitempty" name:"WanDomain"`

	// 外网IP地址，公网可访问
	WanVip *string `json:"WanVip,omitnil,omitempty" name:"WanVip"`

	// 外网访问端口
	WanPort *int64 `json:"WanPort,omitnil,omitempty" name:"WanPort"`

	// 实例所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例自动续费标志。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 独享集群ID
	ExclusterId *string `json:"ExclusterId,omitnil,omitempty" name:"ExclusterId"`

	// 付费模式。prepaid-预付费；postpaid-按量计费
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例创建时间，格式为 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例到期时间，格式为 2006-01-02 15:04:05
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 数据库版本信息
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 实例是否支持审计。0-不支持；1-支持
	IsAuditSupported *int64 `json:"IsAuditSupported,omitnil,omitempty" name:"IsAuditSupported"`

	// 实例是否支持数据加密。0-不支持；1-支持
	IsEncryptSupported *int64 `json:"IsEncryptSupported,omitnil,omitempty" name:"IsEncryptSupported"`

	// 实例母机机器型号
	Machine *string `json:"Machine,omitnil,omitempty" name:"Machine"`

	// 实例内存大小，单位 GB，各个分片的内存大小的和
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例磁盘存储大小，单位 GB，各个分片的磁盘大小的和
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 实例存储空间使用率，计算方式为：各个分片已经使用的磁盘大小的和/各个分片的磁盘大小的和。
	StorageUsage *float64 `json:"StorageUsage,omitnil,omitempty" name:"StorageUsage"`

	// 日志存储空间大小，单位GB
	LogStorage *int64 `json:"LogStorage,omitnil,omitempty" name:"LogStorage"`

	// 产品类型ID
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 主DB可用区
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 从DB可用区
	SlaveZones []*string `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`

	// 分片信息
	Shards []*ShardBriefInfo `json:"Shards,omitnil,omitempty" name:"Shards"`

	// 内网IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip6 *string `json:"Vip6,omitnil,omitempty" name:"Vip6"`

	// 实例Cpu核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// DB引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbEngine *string `json:"DbEngine,omitnil,omitempty" name:"DbEngine"`

	// 是否支持IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// 外网IPv6地址，公网可访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanVipv6 *string `json:"WanVipv6,omitnil,omitempty" name:"WanVipv6"`

	// 外网状态，0-未开通；1-已开通；2-关闭；3-开通中；4-关闭中
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanStatusIpv6 *int64 `json:"WanStatusIpv6,omitnil,omitempty" name:"WanStatusIpv6"`

	// 外网IPv6端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanPortIpv6 *int64 `json:"WanPortIpv6,omitnil,omitempty" name:"WanPortIpv6"`

	// 标签信息
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// DCN标志，0-无，1-主实例，2-灾备实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnFlag *int64 `json:"DcnFlag,omitnil,omitempty" name:"DcnFlag"`

	// DCN状态，0-无，1-创建中，2-同步中，3-已断开
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnStatus *int64 `json:"DcnStatus,omitnil,omitempty" name:"DcnStatus"`

	// DCN灾备实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnDstNum *int64 `json:"DcnDstNum,omitnil,omitempty" name:"DcnDstNum"`

	// 1： 主实例（独享型）, 2: 主实例, 3： 灾备实例, 4： 灾备实例（独享型）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例是否支持设置用户连接数限制，内核为10.1暂不支持。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMaxUserConnectionsSupported *bool `json:"IsMaxUserConnectionsSupported,omitnil,omitempty" name:"IsMaxUserConnectionsSupported"`

	// 对外显示的数据库版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbVersionId *string `json:"DbVersionId,omitnil,omitempty" name:"DbVersionId"`

	// 加密状态, 0-未开启，1-已开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptStatus *int64 `json:"EncryptStatus,omitnil,omitempty" name:"EncryptStatus"`

	// 独享集群类型，0:公有云, 1:金融围笼, 2:CDC集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExclusterType *int64 `json:"ExclusterType,omitnil,omitempty" name:"ExclusterType"`

	// VPC就近访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitnil,omitempty" name:"RsAccessStrategy"`

	// 尚未回收的网络资源
	ReservedNetResources []*ReservedNetResource `json:"ReservedNetResources,omitnil,omitempty" name:"ReservedNetResources"`

	// 是否支持物理复制
	IsPhysicalReplicationSupported *bool `json:"IsPhysicalReplicationSupported,omitnil,omitempty" name:"IsPhysicalReplicationSupported"`

	// 是否支持强同步DCN
	IsDcnStrongSyncSupported *int64 `json:"IsDcnStrongSyncSupported,omitnil,omitempty" name:"IsDcnStrongSyncSupported"`

	// 是否支持DCN切换
	IsDcnSwitchSupported *int64 `json:"IsDcnSwitchSupported,omitnil,omitempty" name:"IsDcnSwitchSupported"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeDCDBInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceNodeInfoRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 单次最多返回多少条，取值范围为(0-100]，默认为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回数据的偏移值，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDCDBInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 单次最多返回多少条，取值范围为(0-100]，默认为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回数据的偏移值，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDCDBInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBInstanceNodeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstanceNodeInfoResponseParams struct {
	// 节点总个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 节点信息
	NodesInfo []*BriefNodeInfo `json:"NodesInfo,omitnil,omitempty" name:"NodesInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBInstanceNodeInfoResponseParams `json:"Response"`
}

func (r *DescribeDCDBInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstancesRequestParams struct {
	// 按照一个或者多个实例 ID 查询。实例 ID 形如：dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。
	SearchName *string `json:"SearchName,omitnil,omitempty" name:"SearchName"`

	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 按项目 ID 查询
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 是否根据 VPC 网络来搜索
	IsFilterVpc *bool `json:"IsFilterVpc,omitnil,omitempty" name:"IsFilterVpc"`

	// 私有网络 ID， IsFilterVpc 为 1 时有效
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 排序字段， projectId， createtime， instancename 三者之一
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 10，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 1非独享集群，2独享集群， 0全部
	ExclusterType *int64 `json:"ExclusterType,omitnil,omitempty" name:"ExclusterType"`

	// 标识是否使用ExclusterType字段, false不使用，true使用
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitnil,omitempty" name:"IsFilterExcluster"`

	// 独享集群ID
	ExclusterIds []*string `json:"ExclusterIds,omitnil,omitempty" name:"ExclusterIds"`

	// 按标签key查询
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例类型过滤，1-独享实例，2-主实例，3-灾备实例，多个按逗号分隔
	FilterInstanceType *string `json:"FilterInstanceType,omitnil,omitempty" name:"FilterInstanceType"`

	// 按实例状态筛选
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排除实例状态
	ExcludeStatus []*int64 `json:"ExcludeStatus,omitnil,omitempty" name:"ExcludeStatus"`
}

type DescribeDCDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例 ID 查询。实例 ID 形如：dcdbt-2t4cf98d
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。
	SearchName *string `json:"SearchName,omitnil,omitempty" name:"SearchName"`

	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 按项目 ID 查询
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 是否根据 VPC 网络来搜索
	IsFilterVpc *bool `json:"IsFilterVpc,omitnil,omitempty" name:"IsFilterVpc"`

	// 私有网络 ID， IsFilterVpc 为 1 时有效
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 排序字段， projectId， createtime， instancename 三者之一
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 10，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 1非独享集群，2独享集群， 0全部
	ExclusterType *int64 `json:"ExclusterType,omitnil,omitempty" name:"ExclusterType"`

	// 标识是否使用ExclusterType字段, false不使用，true使用
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitnil,omitempty" name:"IsFilterExcluster"`

	// 独享集群ID
	ExclusterIds []*string `json:"ExclusterIds,omitnil,omitempty" name:"ExclusterIds"`

	// 按标签key查询
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例类型过滤，1-独享实例，2-主实例，3-灾备实例，多个按逗号分隔
	FilterInstanceType *string `json:"FilterInstanceType,omitnil,omitempty" name:"FilterInstanceType"`

	// 按实例状态筛选
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排除实例状态
	ExcludeStatus []*int64 `json:"ExcludeStatus,omitnil,omitempty" name:"ExcludeStatus"`
}

func (r *DescribeDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "SearchName")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "IsFilterVpc")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ExclusterType")
	delete(f, "IsFilterExcluster")
	delete(f, "ExclusterIds")
	delete(f, "TagKeys")
	delete(f, "Tags")
	delete(f, "FilterInstanceType")
	delete(f, "Status")
	delete(f, "ExcludeStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBInstancesResponseParams struct {
	// 符合条件的实例数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表
	Instances []*DCDBInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBPriceRequestParams struct {
	// 欲新购实例的可用区ID。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 欲购买实例的数量，目前支持购买1-10个实例
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 单个分片节点个数大小，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil,omitempty" name:"ShardNodeCount"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 付费类型。postpaid：按量付费   prepaid：预付费
	Paymode *string `json:"Paymode,omitnil,omitempty" name:"Paymode"`

	// 价格金额单位，不传默认单位为分，取值：  
	// * pent：分
	// * microPent：微分
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

type DescribeDCDBPriceRequest struct {
	*tchttp.BaseRequest
	
	// 欲新购实例的可用区ID。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 欲购买实例的数量，目前支持购买1-10个实例
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 欲购买的时长，单位：月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 单个分片节点个数大小，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil,omitempty" name:"ShardNodeCount"`

	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec
	//  查询实例规格获得。
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount *int64 `json:"ShardCount,omitnil,omitempty" name:"ShardCount"`

	// 付费类型。postpaid：按量付费   prepaid：预付费
	Paymode *string `json:"Paymode,omitnil,omitempty" name:"Paymode"`

	// 价格金额单位，不传默认单位为分，取值：  
	// * pent：分
	// * microPent：微分
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

func (r *DescribeDCDBPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Count")
	delete(f, "Period")
	delete(f, "ShardNodeCount")
	delete(f, "ShardMemory")
	delete(f, "ShardStorage")
	delete(f, "ShardCount")
	delete(f, "Paymode")
	delete(f, "AmountUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBPriceResponseParams struct {
	// 原价  
	// * 单位：默认为分，若请求参数带有AmountUnit，参考AmountUnit描述
	// * 币种：国内站为人民币，国际站为美元
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 实际价格，受折扣等影响，可能和原价不同
	// * 单位：默认为分，若请求参数带有AmountUnit，参考AmountUnit描述
	// * 币种：国内站人民币，国际站美元
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBPriceResponseParams `json:"Response"`
}

func (r *DescribeDCDBPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBRenewalPriceRequestParams struct {
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。不传则默认为1个月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 价格金额单位，不传默认单位为分，取值：  
	// * pent：分
	// * microPent：微分
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

type DescribeDCDBRenewalPriceRequest struct {
	*tchttp.BaseRequest
	
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。不传则默认为1个月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 价格金额单位，不传默认单位为分，取值：  
	// * pent：分
	// * microPent：微分
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

func (r *DescribeDCDBRenewalPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBRenewalPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Period")
	delete(f, "AmountUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBRenewalPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBRenewalPriceResponseParams struct {
	// 原价  
	// * 单位：默认为分，若请求参数带有AmountUnit，参考AmountUnit描述
	// * 币种：国内站为人民币，国际站为美元
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 实际价格，受折扣等影响，可能和原价不同
	// * 单位：默认为分，若请求参数带有AmountUnit，参考AmountUnit描述
	// * 币种：国内站人民币，国际站美元
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBRenewalPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBRenewalPriceResponseParams `json:"Response"`
}

func (r *DescribeDCDBRenewalPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBRenewalPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBSaleInfoRequestParams struct {

}

type DescribeDCDBSaleInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDCDBSaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBSaleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBSaleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBSaleInfoResponseParams struct {
	// 可售卖地域信息列表
	RegionList []*RegionInfo `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBSaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBSaleInfoResponseParams `json:"Response"`
}

func (r *DescribeDCDBSaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBSaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBShardsRequestParams struct {
	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分片ID列表。
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil,omitempty" name:"ShardInstanceIds"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段， 目前仅支持 createtime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeDCDBShardsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分片ID列表。
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil,omitempty" name:"ShardInstanceIds"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段， 目前仅支持 createtime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序类型， desc 或者 asc
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeDCDBShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardInstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBShardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBShardsResponseParams struct {
	// 符合条件的分片数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分片信息列表
	Shards []*DCDBShardInfo `json:"Shards,omitnil,omitempty" name:"Shards"`

	// 灾备标志，0-无，1-主实例，2-灾备实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	DcnFlag *int64 `json:"DcnFlag,omitnil,omitempty" name:"DcnFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBShardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBShardsResponseParams `json:"Response"`
}

func (r *DescribeDCDBShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBUpgradePriceRequestParams struct {
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 价格金额单位，不传默认单位为分，取值：  
	// * pent：分
	// * microPent：微分
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

type DescribeDCDBUpgradePriceRequest struct {
	*tchttp.BaseRequest
	
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 价格金额单位，不传默认单位为分，取值：  
	// * pent：分
	// * microPent：微分
	AmountUnit *string `json:"AmountUnit,omitnil,omitempty" name:"AmountUnit"`
}

func (r *DescribeDCDBUpgradePriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBUpgradePriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeType")
	delete(f, "AddShardConfig")
	delete(f, "ExpandShardConfig")
	delete(f, "SplitShardConfig")
	delete(f, "AmountUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDCDBUpgradePriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDCDBUpgradePriceResponseParams struct {
	// 原价  
	// * 单位：默认为分，若请求参数带有AmountUnit，参考AmountUnit描述
	// * 币种：国内站为人民币，国际站为美元
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 实际价格，受折扣等影响，可能和原价不同
	// * 单位：默认为分，若请求参数带有AmountUnit，参考AmountUnit描述
	// * 币种：国内站人民币，国际站美元
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 变配明细计算公式
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDCDBUpgradePriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDCDBUpgradePriceResponseParams `json:"Response"`
}

func (r *DescribeDCDBUpgradePriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDCDBUpgradePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsRequestParams struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsResponseParams struct {
	// 透传入参。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表列表。
	Tables []*DatabaseTable `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 视图列表。
	Views []*DatabaseView `json:"Views,omitnil,omitempty" name:"Views"`

	// 存储过程列表。
	Procs []*DatabaseProcedure `json:"Procs,omitnil,omitempty" name:"Procs"`

	// 函数列表。
	Funcs []*DatabaseFunction `json:"Funcs,omitnil,omitempty" name:"Funcs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseObjectsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseTableRequestParams struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名称，通过 DescribeDatabaseObjects 接口获取。
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名称，通过 DescribeDatabaseObjects 接口获取。
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "Table")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseTableResponseParams struct {
	// 实例名称。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名称。
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 列信息。
	Cols []*TableColumn `json:"Cols,omitnil,omitempty" name:"Cols"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseTableResponseParams `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// 该实例上的数据库列表。
	Databases []*Database `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 透传入参。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDcnDetailRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDcnDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDcnDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDcnDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDcnDetailResponseParams struct {
	// DCN同步详情
	DcnDetails []*DcnDetailItem `json:"DcnDetails,omitnil,omitempty" name:"DcnDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDcnDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDcnDetailResponseParams `json:"Response"`
}

func (r *DescribeDcnDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadUrlRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例分片ID
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 不带签名的文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

type DescribeFileDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例分片ID
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 不带签名的文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

func (r *DescribeFileDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ShardId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadUrlResponseParams struct {
	// 带签名的下载连接
	PreSignedUrl *string `json:"PreSignedUrl,omitnil,omitempty" name:"PreSignedUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeFileDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowRequestParams struct {
	// 异步请求接口返回的任务流程号。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// 异步请求接口返回的任务流程号。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowResponseParams struct {
	// 流程状态，0：成功，1：失败，2：运行中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowResponseParams `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogFileRetentionPeriodRequestParams struct {
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeLogFileRetentionPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogFileRetentionPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogFileRetentionPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogFileRetentionPeriodResponseParams struct {
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志备份天数
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogFileRetentionPeriodResponseParams `json:"Response"`
}

func (r *DescribeLogFileRetentionPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogFileRetentionPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersRequestParams struct {
	// 待查询的长订单号列表，创建实例、续费实例、扩容实例接口返回。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// 待查询的长订单号列表，创建实例、续费实例、扩容实例接口返回。
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersResponseParams struct {
	// 返回的订单数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 订单信息列表。
	Deals []*Deal `json:"Deals,omitnil,omitempty" name:"Deals"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrdersResponseParams `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// 安全组详情。
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 安全组个数。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {

}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// 项目列表
	Projects []*Project `json:"Projects,omitnil,omitempty" name:"Projects"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShardSpecRequestParams struct {

}

type DescribeShardSpecRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeShardSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShardSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShardSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShardSpecResponseParams struct {
	// 按机型分类的可售卖规格列表
	SpecConfig []*SpecConfig `json:"SpecConfig,omitnil,omitempty" name:"SpecConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShardSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShardSpecResponseParams `json:"Response"`
}

func (r *DescribeShardSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShardSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserTasksRequestParams struct {
	// 任务的状态列表。0-任务启动中；1-任务运行中；2-任务成功；3-任务失败
	Statuses []*int64 `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 任务类型列表，当前支持的任务类型有：0-回档任务；1-创建实例任务；2-扩容任务；3-迁移任务；4-删除实例任务；5-重启任务
	FlowTypes []*int64 `json:"FlowTypes,omitnil,omitempty" name:"FlowTypes"`

	// 任务的创建时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务ID的数组
	UTaskIds []*int64 `json:"UTaskIds,omitnil,omitempty" name:"UTaskIds"`

	// 每次最多返回多少条任务，取值范围为(0,100]，不传的话默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回任务默认是按照创建时间降序排列，从偏移值Offset处开始返回
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeUserTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务的状态列表。0-任务启动中；1-任务运行中；2-任务成功；3-任务失败
	Statuses []*int64 `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 任务类型列表，当前支持的任务类型有：0-回档任务；1-创建实例任务；2-扩容任务；3-迁移任务；4-删除实例任务；5-重启任务
	FlowTypes []*int64 `json:"FlowTypes,omitnil,omitempty" name:"FlowTypes"`

	// 任务的创建时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务ID的数组
	UTaskIds []*int64 `json:"UTaskIds,omitnil,omitempty" name:"UTaskIds"`

	// 每次最多返回多少条任务，取值范围为(0,100]，不传的话默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回任务默认是按照创建时间降序排列，从偏移值Offset处开始返回
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeUserTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Statuses")
	delete(f, "InstanceIds")
	delete(f, "FlowTypes")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UTaskIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserTasksResponseParams struct {
	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表
	FlowSet []*UserTaskInfo `json:"FlowSet,omitnil,omitempty" name:"FlowSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserTasksResponseParams `json:"Response"`
}

func (r *DescribeUserTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDCDBInstanceRequestParams struct {
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDCDBInstanceResponseParams struct {
	// 实例 ID，与入参InstanceId一致。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异步任务的请求 ID，可使用此 ID [查询异步任务的执行结果](https://cloud.tencent.com/document/product/557/56485)。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyDCDBInstanceResponseParams `json:"Response"`
}

func (r *DestroyDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyHourDCDBInstanceRequestParams struct {
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyHourDCDBInstanceResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID [查询异步任务的执行结果](https://cloud.tencent.com/document/product/557/56485)。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例 ID，与入参InstanceId一致。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *DestroyHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 安全组Id。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例ID列表，一个或者多个实例Id组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 安全组Id。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例ID列表，一个或者多个实例Id组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandShardConfig struct {
	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil,omitempty" name:"ShardInstanceIds"`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`

	// 分片节点数
	ShardNodeCount *int64 `json:"ShardNodeCount,omitnil,omitempty" name:"ShardNodeCount"`
}

// Predefined struct for user
type FlushBinlogRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type FlushBinlogRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *FlushBinlogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushBinlogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlushBinlogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FlushBinlogResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FlushBinlogResponse struct {
	*tchttp.BaseResponse
	Response *FlushBinlogResponseParams `json:"Response"`
}

func (r *FlushBinlogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlushBinlogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantAccountPrivilegesRequestParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 全局权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER，SHOW DATABASES，REPLICATION CLIENT，REPLICATION SLAVE
	// 库权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER 
	// 表权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE VIEW，SHOW VIEW，TRIGGER  
	// 字段权限： INSERT，REFERENCES，SELECT，UPDATE
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 类型,可以填入 table 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示设置该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示对表授权，如果为具体字段名，表示对字段授权
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名。如果为 \*，表示查询全局权限（即 \*.\*），此时忽略 Type 和 Object 参数
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 全局权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER，SHOW DATABASES，REPLICATION CLIENT，REPLICATION SLAVE
	// 库权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER 
	// 表权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE VIEW，SHOW VIEW，TRIGGER  
	// 字段权限： INSERT，REFERENCES，SELECT，UPDATE
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 类型,可以填入 table 和 \*。当 DbName 为具体数据库名，Type为 \* 时，表示设置该数据库权限（即db.\*），此时忽略 Object 参数
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 具体的 Type 的名称，例如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 \* 或者为空
	Object *string `json:"Object,omitnil,omitempty" name:"Object"`

	// 当 Type=table 时，ColName 为 \* 表示对表授权，如果为具体字段名，表示对字段授权
	ColName *string `json:"ColName,omitnil,omitempty" name:"ColName"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Privileges")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantAccountPrivilegesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *GrantAccountPrivilegesResponseParams `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitDCDBInstancesRequestParams struct {
	// 待初始化的实例ID列表，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

type InitDCDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待初始化的实例ID列表，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传，0 - 敏感；1-不敏感），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *InitDCDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDCDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitDCDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitDCDBInstancesResponseParams struct {
	// 异步任务ID，可通过 DescribeFlow 查询任务状态。
	FlowIds []*uint64 `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 透传入参。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InitDCDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InitDCDBInstancesResponseParams `json:"Response"`
}

func (r *InitDCDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDCDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceBackupFileItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例状态
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 分片ID
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 备份类型，Data:数据备份，Binlog:Binlog备份，Errlog:错误日志，Slowlog:慢日志
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 手动备份，0:否，1:是
	ManualBackup *int64 `json:"ManualBackup,omitnil,omitempty" name:"ManualBackup"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 对象的存储类型，枚举值：STANDARD（标准存储）、ARCHIVE（归档存储）。
	StorageClass *string `json:"StorageClass,omitnil,omitempty" name:"StorageClass"`
}

// Predefined struct for user
type IsolateDCDBInstanceRequestParams struct {
	// 实例 ID，格式如：tdsqlshard-avw0207d，与云数据库控制台页面中显示的实例 ID 相同，可使用 查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type IsolateDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsqlshard-avw0207d，与云数据库控制台页面中显示的实例 ID 相同，可使用 查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *IsolateDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDCDBInstanceResponseParams struct {
	// 隔离成功实例ID列表。
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// 隔离失败实例ID列表。
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDCDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDedicatedDBInstanceRequestParams struct {
	// 实例 Id，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type IsolateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *IsolateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDedicatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDedicatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDedicatedDBInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDedicatedDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDedicatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateHourDCDBInstanceRequestParams struct {
	// 待升级的实例ID列表。形如：["dcdbt-ow728lmc"]，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type IsolateHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待升级的实例ID列表。形如：["dcdbt-ow728lmc"]，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *IsolateHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateHourDCDBInstanceResponseParams struct {
	// 隔离成功的实例id列表
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil,omitempty" name:"SuccessInstanceIds"`

	// 隔离失败的实例id列表
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillSessionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID列表
	SessionId []*int64 `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 分片ID，与ShardSerialId设置一个
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 分片序列ID，与ShardId设置一个
	ShardSerialId *string `json:"ShardSerialId,omitnil,omitempty" name:"ShardSerialId"`
}

type KillSessionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 会话ID列表
	SessionId []*int64 `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 分片ID，与ShardSerialId设置一个
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 分片序列ID，与ShardId设置一个
	ShardSerialId *string `json:"ShardSerialId,omitnil,omitempty" name:"ShardSerialId"`
}

func (r *KillSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionId")
	delete(f, "ShardId")
	delete(f, "ShardSerialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillSessionResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillSessionResponse struct {
	*tchttp.BaseResponse
	Response *KillSessionResponseParams `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {
	// Log最后修改时间
	Mtime *uint64 `json:"Mtime,omitnil,omitempty" name:"Mtime"`

	// 文件长度
	Length *uint64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 下载Log时用到的统一资源标识符
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

// Predefined struct for user
type ModifyAccountConfigRequestParams struct {
	// 实例 ID，格式如：tdsqlshard-kpkvq5oj，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 账号的名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号的域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 配置列表，每一个元素是Config和Value的组合
	Configs []*ConfigValue `json:"Configs,omitnil,omitempty" name:"Configs"`
}

type ModifyAccountConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsqlshard-kpkvq5oj，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 账号的名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号的域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 配置列表，每一个元素是Config和Value的组合
	Configs []*ConfigValue `json:"Configs,omitnil,omitempty" name:"Configs"`
}

func (r *ModifyAccountConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Configs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountConfigResponseParams `json:"Response"`
}

func (r *ModifyAccountConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 新的账号备注，长度 0~256。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 新的账号备注，长度 0~256。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountDescriptionResponseParams `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// 实例 ID，格式如：tdsql-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的账号，包括用户名和域名。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，该字段传空数组。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 数据库中表的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// 数据库表中列的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

	// 数据库视图的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	ViewPrivileges []*ViewPrivileges `json:"ViewPrivileges,omitnil,omitempty" name:"ViewPrivileges"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsql-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库的账号，包括用户名和域名。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 全局权限。其中，GlobalPrivileges 中权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "PROCESS", "DROP","REFERENCES","INDEX","ALTER","SHOW DATABASES","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，该字段传空数组。
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE TEMPORARY TABLES","LOCK TABLES","EXECUTE","CREATE VIEW","SHOW VIEW","CREATE ROUTINE","ALTER ROUTINE","EVENT","TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 数据库中表的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`

	// 数据库表中列的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","REFERENCES"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil,omitempty" name:"ColumnPrivileges"`

	// 数据库视图的权限。Privileges 权限的可选值为："SELECT","INSERT","UPDATE","DELETE","CREATE", "DROP","REFERENCES","INDEX","ALTER","CREATE VIEW","SHOW VIEW", "TRIGGER"。
	// 注意，不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组。
	ViewPrivileges []*ViewPrivileges `json:"ViewPrivileges,omitnil,omitempty" name:"ViewPrivileges"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	delete(f, "ColumnPrivileges")
	delete(f, "ViewPrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID [查询异步任务的执行结果](https://cloud.tencent.com/document/product/237/16177)。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigsRequestParams struct {
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 常规备份存储时长，范围[1, 3650]。
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 每天备份执行的区间的开始时间，格式 mm:ss，形如 22:00。
	StartBackupTime *string `json:"StartBackupTime,omitnil,omitempty" name:"StartBackupTime"`

	// 每天备份执行的区间的结束时间，格式 mm:ss，形如 23:59。
	EndBackupTime *string `json:"EndBackupTime,omitnil,omitempty" name:"EndBackupTime"`

	// 执行备份周期，枚举值：Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 沉降到归档存储时长，-1表示关闭归档设置。
	ArchiveDays *int64 `json:"ArchiveDays,omitnil,omitempty" name:"ArchiveDays"`

	// 超期备份配置。
	BackupConfigSet []*NewBackupConfig `json:"BackupConfigSet,omitnil,omitempty" name:"BackupConfigSet"`
}

type ModifyBackupConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 常规备份存储时长，范围[1, 3650]。
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 每天备份执行的区间的开始时间，格式 mm:ss，形如 22:00。
	StartBackupTime *string `json:"StartBackupTime,omitnil,omitempty" name:"StartBackupTime"`

	// 每天备份执行的区间的结束时间，格式 mm:ss，形如 23:59。
	EndBackupTime *string `json:"EndBackupTime,omitnil,omitempty" name:"EndBackupTime"`

	// 执行备份周期，枚举值：Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 沉降到归档存储时长，-1表示关闭归档设置。
	ArchiveDays *int64 `json:"ArchiveDays,omitnil,omitempty" name:"ArchiveDays"`

	// 超期备份配置。
	BackupConfigSet []*NewBackupConfig `json:"BackupConfigSet,omitnil,omitempty" name:"BackupConfigSet"`
}

func (r *ModifyBackupConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Days")
	delete(f, "StartBackupTime")
	delete(f, "EndBackupTime")
	delete(f, "WeekDays")
	delete(f, "ArchiveDays")
	delete(f, "BackupConfigSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupConfigsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupConfigsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupConfigsResponseParams `json:"Response"`
}

func (r *ModifyBackupConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBEncryptAttributesRequestParams struct {
	// 实例Id，形如：tdsqlshard-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否启用数据加密，开启后暂不支持关闭。本接口的可选值为：1-开启数据加密。
	EncryptEnabled *int64 `json:"EncryptEnabled,omitnil,omitempty" name:"EncryptEnabled"`
}

type ModifyDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id，形如：tdsqlshard-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否启用数据加密，开启后暂不支持关闭。本接口的可选值为：1-开启数据加密。
	EncryptEnabled *int64 `json:"EncryptEnabled,omitnil,omitempty" name:"EncryptEnabled"`
}

func (r *ModifyDBEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EncryptEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBEncryptAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBEncryptAttributesResponseParams `json:"Response"`
}

func (r *ModifyDBEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// 实例ID，形如tdsql-hdaprz0v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如tdsql-hdaprz0v
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：dcdb。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectRequestParams struct {
	// 待修改的实例ID列表。实例 ID 形如：dcdbt-ow728lmc。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要分配的项目 ID，可以通过 DescribeProjects 查询项目列表接口获取。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的实例ID列表。实例 ID 形如：dcdbt-ow728lmc。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 要分配的项目 ID，可以通过 DescribeProjects 查询项目列表接口获取。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstancesProjectResponseParams `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersRequestParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersResponseParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 各参数修改结果
	Result []*ParamModifyResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBParametersResponseParams `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSyncModeRequestParams struct {
	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 同步模式：0 异步，1 强同步， 2 强同步可退化
	SyncMode *int64 `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`
}

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest
	
	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 同步模式：0 异步，1 强同步， 2 强同步可退化
	SyncMode *int64 `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SyncMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBSyncModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSyncModeResponseParams struct {
	// 异步任务Id，可通过 DescribeFlow 查询任务状态。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBSyncModeResponseParams `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 希望转到的VPC网络的VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 希望转到的VPC网络的子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 如果需要指定VIP，填上该字段
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 如果需要指定VIPv6，填上该字段
	Vipv6 *string `json:"Vipv6,omitnil,omitempty" name:"Vipv6"`

	// VIP保留时长，单位小时，取值范围（0~168），0表示立即释放，有一分钟释放延迟。不传此参数，默认24小时释放VIP。
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil,omitempty" name:"VipReleaseDelay"`
}

type ModifyInstanceNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 希望转到的VPC网络的VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 希望转到的VPC网络的子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 如果需要指定VIP，填上该字段
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 如果需要指定VIPv6，填上该字段
	Vipv6 *string `json:"Vipv6,omitnil,omitempty" name:"Vipv6"`

	// VIP保留时长，单位小时，取值范围（0~168），0表示立即释放，有一分钟释放延迟。不传此参数，默认24小时释放VIP。
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil,omitempty" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	delete(f, "Vipv6")
	delete(f, "VipReleaseDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkResponseParams struct {
	// 异步任务ID，根据此FlowId通过DescribeFlow接口查询任务进行状态
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNetworkResponseParams `json:"Response"`
}

func (r *ModifyInstanceNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVipRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// IPv6标志
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// VIP保留时长，单位小时，取值范围（0~168），0表示立即释放，有一分钟释放延迟。不传此参数，默认24小时释放VIP。
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil,omitempty" name:"VipReleaseDelay"`
}

type ModifyInstanceVipRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例VIP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// IPv6标志
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`

	// VIP保留时长，单位小时，取值范围（0~168），0表示立即释放，有一分钟释放延迟。不传此参数，默认24小时释放VIP。
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil,omitempty" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Vip")
	delete(f, "Ipv6Flag")
	delete(f, "VipReleaseDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVipResponseParams struct {
	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceVipResponseParams `json:"Response"`
}

func (r *ModifyInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVportRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例VPORT
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type ModifyInstanceVportRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例VPORT
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

func (r *ModifyInstanceVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVportResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceVportResponseParams `json:"Response"`
}

func (r *ModifyInstanceVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRealServerAccessStrategyRequestParams struct {
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// RS就近模式, 0-无策略, 1-可用区就近访问。
	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitnil,omitempty" name:"RsAccessStrategy"`
}

type ModifyRealServerAccessStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：tdsqlshard-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// RS就近模式, 0-无策略, 1-可用区就近访问。
	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitnil,omitempty" name:"RsAccessStrategy"`
}

func (r *ModifyRealServerAccessStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealServerAccessStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RsAccessStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRealServerAccessStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRealServerAccessStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRealServerAccessStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRealServerAccessStrategyResponseParams `json:"Response"`
}

func (r *ModifyRealServerAccessStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRealServerAccessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewBackupConfig struct {
	// 备份策略是否启用。
	EnableBackupPolicy *bool `json:"EnableBackupPolicy,omitnil,omitempty" name:"EnableBackupPolicy"`

	// 超期保留开始日期，早于开始日期的超期备份不保留，格式：yyyy-mm-dd。
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 超期备份保留时长，超出保留时间的超期备份将被删除，可填写1-3650整数。
	MaxRetentionDays *int64 `json:"MaxRetentionDays,omitnil,omitempty" name:"MaxRetentionDays"`

	// 备份模式，可选择按年月周模式保存
	// * 按年：annually
	// * 按月：monthly
	// * 按周：weekly
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// Frequency等于weekly时生效。
	// 表示保留特定工作日备份。可选择周一到周日，支持多选，取星期英文：
	// * 星期一 ：Monday
	// * 星期二 ：Tuesday
	// * 星期三：Wednesday
	// * 星期四：Thursday
	// * 星期五：Friday
	// * 星期六：Saturday
	// * 星期日：Sunday
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 保留备份个数，Frequency等于monthly或weekly时生效。
	// 备份模式选择按月时，可填写1-28整数；
	// 备份模式选择年时，可填写1-336整数。
	BackupCount *int64 `json:"BackupCount,omitnil,omitempty" name:"BackupCount"`
}

type NodeInfo struct {
	// DB节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// DB节点角色，取值为master或者slave
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

// Predefined struct for user
type OpenDBExtranetAccessRequestParams struct {
	// 待开放外网访问的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否IPv6，默认0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 待开放外网访问的实例ID。形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否IPv6，默认0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`
}

func (r *OpenDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ipv6Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBExtranetAccessResponseParams struct {
	// 异步任务ID，可通过 DescribeFlow 查询任务状态。
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenDBExtranetAccessResponseParams `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamConstraint struct {
	// 约束类型,如枚举enum，区间section
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 约束类型为enum时的可选值列表
	Enum *string `json:"Enum,omitnil,omitempty" name:"Enum"`

	// 约束类型为section时的范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *ConstraintRange `json:"Range,omitnil,omitempty" name:"Range"`

	// 约束类型为string时的可选值列表
	String *string `json:"String,omitnil,omitempty" name:"String"`
}

type ParamDesc struct {
	// 参数名字
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// 当前参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 设置过的值，参数生效后，该值和value一样。未设置过就不返回该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetValue *string `json:"SetValue,omitnil,omitempty" name:"SetValue"`

	// 系统默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数限制
	Constraint *ParamConstraint `json:"Constraint,omitnil,omitempty" name:"Constraint"`

	// 是否有设置过值，false:没有设置过值，true:有设置过值。
	HaveSetValue *bool `json:"HaveSetValue,omitnil,omitempty" name:"HaveSetValue"`

	// 是否需要重启生效，false:不需要重启，
	// true:需要重启
	NeedRestart *bool `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`
}

type ParamModifyResult struct {
	// 修改参数名字
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// 参数修改结果。0表示修改成功；-1表示修改失败；-2表示该参数值非法
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`
}

type Project struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源拥有者（主账号）uin
	OwnerUin *int64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 应用Id
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 项目名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建者uin
	CreatorUin *int64 `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 来源平台
	SrcPlat *string `json:"SrcPlat,omitnil,omitempty" name:"SrcPlat"`

	// 来源AppId
	SrcAppId *int64 `json:"SrcAppId,omitnil,omitempty" name:"SrcAppId"`

	// 项目状态,0正常，-1关闭。默认项目为3
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否默认项目，1 是，0 不是
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 描述信息
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`
}

type RegionInfo struct {
	// 地域英文ID
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域数字ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域中文名
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区列表
	ZoneList []*ZonesInfo `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// 可选择的主可用区和从可用区
	AvailableChoice []*ShardZoneChooseInfo `json:"AvailableChoice,omitnil,omitempty" name:"AvailableChoice"`
}

// Predefined struct for user
type RenewDCDBInstanceRequestParams struct {
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type RenewDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费时长，单位：月。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

func (r *RenewDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDCDBInstanceResponseParams struct {
	// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewDCDBInstanceResponseParams `json:"Response"`
}

func (r *RenewDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReservedNetResource struct {
	// 私有网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VpcId,SubnetId下保留的内网ip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Vip下的端口
	Vports []*int64 `json:"Vports,omitnil,omitempty" name:"Vports"`

	// Vip的回收时间	
	RecycleTime *string `json:"RecycleTime,omitnil,omitempty" name:"RecycleTime"`
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 使用GetPublicKey返回的RSA2048公钥加密后的密码，加密算法是PKCS1v15
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 登录用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 使用GetPublicKey返回的RSA2048公钥加密后的密码，加密算法是PKCS1v15
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "EncryptedPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {
	// 标签键key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type SecurityGroup struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// 入站规则
	Inbound []*SecurityGroupBound `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// 出站规则
	Outbound []*SecurityGroupBound `json:"Outbound,omitnil,omitempty" name:"Outbound"`
}

type SecurityGroupBound struct {
	// 来源 IP 或 IP 段，例如192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 端口
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP 等
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`
}

type ShardBriefInfo struct {
	// 分片SerialId
	ShardSerialId *string `json:"ShardSerialId,omitnil,omitempty" name:"ShardSerialId"`

	// 分片ID，形如shard-7vg1o339
	ShardInstanceId *string `json:"ShardInstanceId,omitnil,omitempty" name:"ShardInstanceId"`

	// 分片运行状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 分片运行状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 分片创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分片内存大小，单位GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 分片磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 分片日志磁盘空间大小，单位GB
	LogDisk *int64 `json:"LogDisk,omitnil,omitempty" name:"LogDisk"`

	// 分片节点个数
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 分片磁盘空间使用率
	StorageUsage *float64 `json:"StorageUsage,omitnil,omitempty" name:"StorageUsage"`

	// 分片Proxy版本信息
	ProxyVersion *string `json:"ProxyVersion,omitnil,omitempty" name:"ProxyVersion"`

	// 分片主DB可用区
	ShardMasterZone *string `json:"ShardMasterZone,omitnil,omitempty" name:"ShardMasterZone"`

	// 分片从DB可用区
	ShardSlaveZones []*string `json:"ShardSlaveZones,omitnil,omitempty" name:"ShardSlaveZones"`

	// 分片Cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// DB节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodesInfo []*NodeInfo `json:"NodesInfo,omitnil,omitempty" name:"NodesInfo"`
}

type ShardInfo struct {
	// 分片ID
	ShardInstanceId *string `json:"ShardInstanceId,omitnil,omitempty" name:"ShardInstanceId"`

	// 分片Set ID
	ShardSerialId *string `json:"ShardSerialId,omitnil,omitempty" name:"ShardSerialId"`

	// 状态：0 创建中，1 流程处理中， 2 运行中，3 分片未初始化，-2 分片已删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	Createtime *string `json:"Createtime,omitnil,omitempty" name:"Createtime"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储大小，单位 GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 分片数字ID
	ShardId *int64 `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 节点数，2 为一主一从， 3 为一主二从
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 产品类型 Id（过时字段，请勿依赖该值）
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Cpu核数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

type ShardZoneChooseInfo struct {
	// 主可用区
	MasterZone *ZonesInfo `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 可选的从可用区
	SlaveZones []*ZonesInfo `json:"SlaveZones,omitnil,omitempty" name:"SlaveZones"`
}

type SlowLogData struct {
	// 语句校验和，用于查询详情
	CheckSum *string `json:"CheckSum,omitnil,omitempty" name:"CheckSum"`

	// 数据库名称
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 抽象的SQL语句
	FingerPrint *string `json:"FingerPrint,omitnil,omitempty" name:"FingerPrint"`

	// 平均的锁时间
	LockTimeAvg *string `json:"LockTimeAvg,omitnil,omitempty" name:"LockTimeAvg"`

	// 最大锁时间
	LockTimeMax *string `json:"LockTimeMax,omitnil,omitempty" name:"LockTimeMax"`

	// 最小锁时间
	LockTimeMin *string `json:"LockTimeMin,omitnil,omitempty" name:"LockTimeMin"`

	// 锁时间总和
	LockTimeSum *string `json:"LockTimeSum,omitnil,omitempty" name:"LockTimeSum"`

	// 查询次数
	QueryCount *string `json:"QueryCount,omitnil,omitempty" name:"QueryCount"`

	// 平均查询时间
	QueryTimeAvg *string `json:"QueryTimeAvg,omitnil,omitempty" name:"QueryTimeAvg"`

	// 最大查询时间
	QueryTimeMax *string `json:"QueryTimeMax,omitnil,omitempty" name:"QueryTimeMax"`

	// 最小查询时间
	QueryTimeMin *string `json:"QueryTimeMin,omitnil,omitempty" name:"QueryTimeMin"`

	// 查询时间总和
	QueryTimeSum *string `json:"QueryTimeSum,omitnil,omitempty" name:"QueryTimeSum"`

	// 扫描行数
	RowsExaminedSum *string `json:"RowsExaminedSum,omitnil,omitempty" name:"RowsExaminedSum"`

	// 发送行数
	RowsSentSum *string `json:"RowsSentSum,omitnil,omitempty" name:"RowsSentSum"`

	// 最后执行时间
	TsMax *string `json:"TsMax,omitnil,omitempty" name:"TsMax"`

	// 首次执行时间
	TsMin *string `json:"TsMin,omitnil,omitempty" name:"TsMin"`

	// 账号
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 样例Sql
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExampleSql *string `json:"ExampleSql,omitnil,omitempty" name:"ExampleSql"`

	// 账户的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

type SpecConfig struct {
	// 规格机型
	Machine *string `json:"Machine,omitnil,omitempty" name:"Machine"`

	// 规格列表
	SpecConfigInfos []*SpecConfigInfo `json:"SpecConfigInfos,omitnil,omitempty" name:"SpecConfigInfos"`
}

type SpecConfigInfo struct {
	// 节点个数，2 表示一主一从，3 表示一主二从
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 内存大小，单位 GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 数据盘规格最小值，单位 GB
	MinStorage *int64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// 数据盘规格最大值，单位 GB
	MaxStorage *int64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 推荐的使用场景
	SuitInfo *string `json:"SuitInfo,omitnil,omitempty" name:"SuitInfo"`

	// 产品类型 Id
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 最大 Qps 值
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// CPU核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`
}

type SplitShardConfig struct {
	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil,omitempty" name:"ShardInstanceIds"`

	// 数据切分比例，固定50%
	SplitRate *int64 `json:"SplitRate,omitnil,omitempty" name:"SplitRate"`

	// 分片内存大小，单位 GB
	ShardMemory *int64 `json:"ShardMemory,omitnil,omitempty" name:"ShardMemory"`

	// 分片存储大小，单位 GB
	ShardStorage *int64 `json:"ShardStorage,omitnil,omitempty" name:"ShardStorage"`
}

// Predefined struct for user
type SwitchDBInstanceHARequestParams struct {
	// 实例Id，形如 tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 切换的目标区域，会自动选择该可用区中延迟最低的节点。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 指定分片实例id进行切换
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil,omitempty" name:"ShardInstanceIds"`
}

type SwitchDBInstanceHARequest struct {
	*tchttp.BaseRequest
	
	// 实例Id，形如 tdsql-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 切换的目标区域，会自动选择该可用区中延迟最低的节点。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 指定分片实例id进行切换
	ShardInstanceIds []*string `json:"ShardInstanceIds,omitnil,omitempty" name:"ShardInstanceIds"`
}

func (r *SwitchDBInstanceHARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceHARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Zone")
	delete(f, "ShardInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstanceHARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstanceHAResponseParams struct {
	// 异步流程Id
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchDBInstanceHAResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDBInstanceHAResponseParams `json:"Response"`
}

func (r *SwitchDBInstanceHAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstanceHAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableColumn struct {
	// 列名称
	Col *string `json:"Col,omitnil,omitempty" name:"Col"`

	// 列类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type TablePrivilege struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库表名
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type TerminateDedicatedDBInstanceRequestParams struct {
	// 实例 Id，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TerminateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 Id，形如：dcdbt-ow728lmc。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TerminateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDedicatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDedicatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDedicatedDBInstanceResponseParams struct {
	// 异步流程Id
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDedicatedDBInstanceResponseParams `json:"Response"`
}

func (r *TerminateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDedicatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TmpInstance struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceRemark *string `json:"InstanceRemark,omitnil,omitempty" name:"InstanceRemark"`

	// 0:非临时实例 ,1:无效临时实例, 2:回档成功的有效临时实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	TempType *int64 `json:"TempType,omitnil,omitempty" name:"TempType"`

	// 实例状态,0:待初始化,1:流程处理中,2:有效状态,-1:已隔离，-2：已下线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例 ID，形如：tdsql-ow728lmc。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例虚IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 实例虚端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 有效期结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodEndTime *string `json:"PeriodEndTime,omitnil,omitempty" name:"PeriodEndTime"`

	// 源实例 ID，形如：tdsql-ow728lmc。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// 实例状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 实例所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例虚IPv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vipv6 *string `json:"Vipv6,omitnil,omitempty" name:"Vipv6"`

	// 实例IPv6标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil,omitempty" name:"Ipv6Flag"`
}

// Predefined struct for user
type UpgradeDCDBInstanceRequestParams struct {
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 变更部署时指定的新可用区列表，第1个为主可用区，其余为从可用区
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 切换开始时间，格式如: "2019-12-12 07:00:00"。开始时间必须在当前时间一个小时以后，3天以内。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换结束时间, 格式如: "2019-12-12 07:15:00"，结束时间必须大于开始时间。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 是否自动重试。 0：不自动重试 1：自动重试
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil,omitempty" name:"SwitchAutoRetry"`
}

type UpgradeDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 变更部署时指定的新可用区列表，第1个为主可用区，其余为从可用区
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 切换开始时间，格式如: "2019-12-12 07:00:00"。开始时间必须在当前时间一个小时以后，3天以内。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换结束时间, 格式如: "2019-12-12 07:15:00"，结束时间必须大于开始时间。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 是否自动重试。 0：不自动重试 1：自动重试
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil,omitempty" name:"SwitchAutoRetry"`
}

func (r *UpgradeDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeType")
	delete(f, "AddShardConfig")
	delete(f, "ExpandShardConfig")
	delete(f, "SplitShardConfig")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "Zones")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	delete(f, "SwitchAutoRetry")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDCDBInstanceResponseParams struct {
	// 长订单号。可以据此调用 DescribeOrders
	//  查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDCDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDedicatedDCDBInstanceRequestParams struct {
	// 升级类型，取值为ADD，SPLIT和EXPAND。ADD-添加分片；SPLIT-切分某个分片；EXPAND-垂直扩容某个分片
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 实例ID，形如 dcdbt-mlfjm74h
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当UpgradeType取值为ADD时，添加分片的配置参数
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 当UpgradeType取值为EXPAND时，垂直扩容分片的配置参数
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 当UpgradeType取值为SPLIT时，切分分片的配置参数
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 错过切换时间窗口时，是否自动重试一次，0-否，1-是
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil,omitempty" name:"SwitchAutoRetry"`

	// 切换时间窗口开始时间
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换时间窗口结束时间
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

type UpgradeDedicatedDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 升级类型，取值为ADD，SPLIT和EXPAND。ADD-添加分片；SPLIT-切分某个分片；EXPAND-垂直扩容某个分片
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 实例ID，形如 dcdbt-mlfjm74h
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当UpgradeType取值为ADD时，添加分片的配置参数
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 当UpgradeType取值为EXPAND时，垂直扩容分片的配置参数
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 当UpgradeType取值为SPLIT时，切分分片的配置参数
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 错过切换时间窗口时，是否自动重试一次，0-否，1-是
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil,omitempty" name:"SwitchAutoRetry"`

	// 切换时间窗口开始时间
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换时间窗口结束时间
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

func (r *UpgradeDedicatedDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDedicatedDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UpgradeType")
	delete(f, "InstanceId")
	delete(f, "AddShardConfig")
	delete(f, "ExpandShardConfig")
	delete(f, "SplitShardConfig")
	delete(f, "SwitchAutoRetry")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDedicatedDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDedicatedDCDBInstanceResponseParams struct {
	// 异步任务流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDedicatedDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDedicatedDCDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDedicatedDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDedicatedDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeHourDCDBInstanceRequestParams struct {
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 切换开始时间，格式如: "2019-12-12 07:00:00"。开始时间必须在当前时间一个小时以后，3天以内。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换结束时间,  格式如: "2019-12-12 07:15:00"，结束时间必须大于开始时间。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 是否自动重试。 0：不自动重试  1：自动重试
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil,omitempty" name:"SwitchAutoRetry"`

	// 变更部署时指定的新可用区列表，第1个为主可用区，其余为从可用区
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`
}

type UpgradeHourDCDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级类型，取值范围: 
	// <li> ADD: 新增分片 </li> 
	//  <li> EXPAND: 升级实例中的已有分片 </li> 
	//  <li> SPLIT: 将已有分片中的数据切分到新增分片上</li>
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `json:"AddShardConfig,omitnil,omitempty" name:"AddShardConfig"`

	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `json:"ExpandShardConfig,omitnil,omitempty" name:"ExpandShardConfig"`

	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `json:"SplitShardConfig,omitnil,omitempty" name:"SplitShardConfig"`

	// 切换开始时间，格式如: "2019-12-12 07:00:00"。开始时间必须在当前时间一个小时以后，3天以内。
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// 切换结束时间,  格式如: "2019-12-12 07:15:00"，结束时间必须大于开始时间。
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// 是否自动重试。 0：不自动重试  1：自动重试
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil,omitempty" name:"SwitchAutoRetry"`

	// 变更部署时指定的新可用区列表，第1个为主可用区，其余为从可用区
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`
}

func (r *UpgradeHourDCDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeHourDCDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeType")
	delete(f, "AddShardConfig")
	delete(f, "ExpandShardConfig")
	delete(f, "SplitShardConfig")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	delete(f, "SwitchAutoRetry")
	delete(f, "Zones")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeHourDCDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeHourDCDBInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeHourDCDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeHourDCDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeHourDCDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeHourDCDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserTaskInfo struct {
	// 任务ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户账户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 任务状态，0-任务初始化中；1-任务运行中；2-任务成功；3-任务失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型，0-实例回档；1-实例创建；2-实例扩容；3-实例迁移；4-实例删除；5-实例重启
	UserTaskType *int64 `json:"UserTaskType,omitnil,omitempty" name:"UserTaskType"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务错误信息
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 客户端参数
	InputData *string `json:"InputData,omitnil,omitempty" name:"InputData"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type ViewPrivileges struct {
	// 数据库名
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据库视图名
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 权限信息
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`
}

type ZonesInfo struct {
	// 可用区英文ID
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区数字ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区中文名
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 是否在售
	OnSale *bool `json:"OnSale,omitnil,omitempty" name:"OnSale"`
}