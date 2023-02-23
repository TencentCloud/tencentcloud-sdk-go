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

package v20210527

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddUserContactRequestParams struct {
	// 联系人姓名，由中英文、数字、空格、!@#$%^&*()_+-=（）组成，不能以下划线开头，长度在20以内。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 邮箱地址，支持大小写字母、数字、下划线及@字符， 不能以下划线开头，邮箱地址不可重复。
	ContactInfo *string `json:"ContactInfo,omitempty" name:"ContactInfo"`

	// 服务产品类型，固定值："mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type AddUserContactRequest struct {
	*tchttp.BaseRequest
	
	// 联系人姓名，由中英文、数字、空格、!@#$%^&*()_+-=（）组成，不能以下划线开头，长度在20以内。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 邮箱地址，支持大小写字母、数字、下划线及@字符， 不能以下划线开头，邮箱地址不可重复。
	ContactInfo *string `json:"ContactInfo,omitempty" name:"ContactInfo"`

	// 服务产品类型，固定值："mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *AddUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ContactInfo")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserContactResponseParams struct {
	// 添加成功的联系人id。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddUserContactResponse struct {
	*tchttp.BaseResponse
	Response *AddUserContactResponseParams `json:"Response"`
}

func (r *AddUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKillTaskRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CancelKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CancelKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKillTaskResponseParams struct {
	// kill会话任务终止成功返回1。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelKillTaskResponseParams `json:"Response"`
}

func (r *CancelKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContactItem struct {
	// 联系人id。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 联系人姓名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 联系人绑定的邮箱。
	Mail *string `json:"Mail,omitempty" name:"Mail"`
}

// Predefined struct for user
type CreateDBDiagReportTaskRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2020-11-08T14:00:00+08:00”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2020-11-09T14:00:00+08:00”。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否发送邮件: 0 - 否，1 - 是。
	SendMailFlag *int64 `json:"SendMailFlag,omitempty" name:"SendMailFlag"`

	// 接收邮件的联系人ID数组。
	ContactPerson []*int64 `json:"ContactPerson,omitempty" name:"ContactPerson"`

	// 接收邮件的联系组ID数组。
	ContactGroup []*int64 `json:"ContactGroup,omitempty" name:"ContactGroup"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认值为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateDBDiagReportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2020-11-08T14:00:00+08:00”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2020-11-09T14:00:00+08:00”。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否发送邮件: 0 - 否，1 - 是。
	SendMailFlag *int64 `json:"SendMailFlag,omitempty" name:"SendMailFlag"`

	// 接收邮件的联系人ID数组。
	ContactPerson []*int64 `json:"ContactPerson,omitempty" name:"ContactPerson"`

	// 接收邮件的联系组ID数组。
	ContactGroup []*int64 `json:"ContactGroup,omitempty" name:"ContactGroup"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认值为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SendMailFlag")
	delete(f, "ContactPerson")
	delete(f, "ContactGroup")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportTaskResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBDiagReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportTaskResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 健康报告相应的任务ID，可通过DescribeDBDiagReportTasks查询。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateDBDiagReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 健康报告相应的任务ID，可通过DescribeDBDiagReportTasks查询。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlResponseParams struct {
	// 健康报告浏览地址。
	ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

	// 健康报告浏览地址到期时间戳（秒）。
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBDiagReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportUrlResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKillTaskRequestParams struct {
	// kill会话任务的关联实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务持续时间，单位秒，手动关闭任务传-1。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 任务过滤条件，客户端IP。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 任务过滤条件，数据库库名,多个","隔开。
	DB *string `json:"DB,omitempty" name:"DB"`

	// 任务过滤条件，相关命令，多个","隔开。
	Command *string `json:"Command,omitempty" name:"Command"`

	// 任务过滤条件，支持单条件前缀匹配。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 任务过滤条件，用户类型。
	User *string `json:"User,omitempty" name:"User"`

	// 任务过滤条件，会话持续时长，单位秒。
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// kill会话任务的关联实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务持续时间，单位秒，手动关闭任务传-1。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 任务过滤条件，客户端IP。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 任务过滤条件，数据库库名,多个","隔开。
	DB *string `json:"DB,omitempty" name:"DB"`

	// 任务过滤条件，相关命令，多个","隔开。
	Command *string `json:"Command,omitempty" name:"Command"`

	// 任务过滤条件，支持单条件前缀匹配。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 任务过滤条件，用户类型。
	User *string `json:"User,omitempty" name:"User"`

	// 任务过滤条件，会话持续时长，单位秒。
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Duration")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "Command")
	delete(f, "Info")
	delete(f, "User")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKillTaskResponseParams struct {
	// kill会话任务创建成功返回1
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateKillTaskResponseParams `json:"Response"`
}

func (r *CreateKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileRequestParams struct {
	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	ProfileLevel *string `json:"ProfileLevel,omitempty" name:"ProfileLevel"`

	// 配置名称，需要保持唯一性，数据库巡检邮件配置名称自拟；定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 配置绑定的实例ID，当配置级别为"Instance"时需要传入且只能为一个实例；当配置级别为“User”时，此参数不填。
	BindInstanceIds []*string `json:"BindInstanceIds,omitempty" name:"BindInstanceIds"`
}

type CreateMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	ProfileLevel *string `json:"ProfileLevel,omitempty" name:"ProfileLevel"`

	// 配置名称，需要保持唯一性，数据库巡检邮件配置名称自拟；定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 配置绑定的实例ID，当配置级别为"Instance"时需要传入且只能为一个实例；当配置级别为“User”时，此参数不填。
	BindInstanceIds []*string `json:"BindInstanceIds,omitempty" name:"BindInstanceIds"`
}

func (r *CreateMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileInfo")
	delete(f, "ProfileLevel")
	delete(f, "ProfileName")
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "BindInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateMailProfileResponseParams `json:"Response"`
}

func (r *CreateMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxySessionKillTaskRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateProxySessionKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateProxySessionKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxySessionKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxySessionKillTaskResponseParams struct {
	// 创建 kill 会话任务返回的异步任务 id
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateProxySessionKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxySessionKillTaskResponseParams `json:"Response"`
}

func (r *CreateProxySessionKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileRequestParams struct {
	// 取值范围1-7，分别代表周一至周日。
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitempty" name:"WeekConfiguration"`

	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`

	// 配置名称，需要保持唯一性，定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// 配置订阅的实例ID。
	BindInstanceId *string `json:"BindInstanceId,omitempty" name:"BindInstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateSchedulerMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 取值范围1-7，分别代表周一至周日。
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitempty" name:"WeekConfiguration"`

	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`

	// 配置名称，需要保持唯一性，定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"schduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// 配置订阅的实例ID。
	BindInstanceId *string `json:"BindInstanceId,omitempty" name:"BindInstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateSchedulerMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WeekConfiguration")
	delete(f, "ProfileInfo")
	delete(f, "ProfileName")
	delete(f, "BindInstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulerMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSchedulerMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchedulerMailProfileResponseParams `json:"Response"`
}

func (r *CreateSchedulerMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAuditLogExportTaskRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 导出日志开始时间，例如2020-12-28 00:00:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 导出日志结束时间，例如2020-12-28 01:00:00。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 日志风险等级列表，支持值包括：0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*int64 `json:"DangerLevels,omitempty" name:"DangerLevels"`
}

type CreateSecurityAuditLogExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 导出日志开始时间，例如2020-12-28 00:00:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 导出日志结束时间，例如2020-12-28 01:00:00。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 日志风险等级列表，支持值包括：0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*int64 `json:"DangerLevels,omitempty" name:"DangerLevels"`
}

func (r *CreateSecurityAuditLogExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "DangerLevels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityAuditLogExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAuditLogExportTaskResponseParams struct {
	// 日志导出任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSecurityAuditLogExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityAuditLogExportTaskResponseParams `json:"Response"`
}

func (r *CreateSecurityAuditLogExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSqlFilterRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 通过VerifyUserAccount获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// SQL类型，取值包括SELECT, UPDATE, DELETE, INSERT, REPLACE。
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// 关键字，用于筛选SQL语句，多个关键字用英文逗号分隔，逗号不能作为关键词，多个关键词之间的关系为“逻辑与”。
	FilterKey *string `json:"FilterKey,omitempty" name:"FilterKey"`

	// 最大并发度，取值不能小于0，如果该值设为 0，则表示限制所有匹配的SQL执行。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 限流时长，单位秒，支持-1和小于2147483647的正整数，-1表示永不过期。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type CreateSqlFilterRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 通过VerifyUserAccount获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// SQL类型，取值包括SELECT, UPDATE, DELETE, INSERT, REPLACE。
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// 关键字，用于筛选SQL语句，多个关键字用英文逗号分隔，逗号不能作为关键词，多个关键词之间的关系为“逻辑与”。
	FilterKey *string `json:"FilterKey,omitempty" name:"FilterKey"`

	// 最大并发度，取值不能小于0，如果该值设为 0，则表示限制所有匹配的SQL执行。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 限流时长，单位秒，支持-1和小于2147483647的正整数，-1表示永不过期。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *CreateSqlFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSqlFilterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionToken")
	delete(f, "SqlType")
	delete(f, "FilterKey")
	delete(f, "MaxConcurrency")
	delete(f, "Duration")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSqlFilterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSqlFilterResponseParams struct {
	// 限流任务ID。
	FilterId *int64 `json:"FilterId,omitempty" name:"FilterId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSqlFilterResponse struct {
	*tchttp.BaseResponse
	Response *CreateSqlFilterResponseParams `json:"Response"`
}

func (r *CreateSqlFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSqlFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 日志导出任务Id列表，接口会忽略不存在或已删除的任务Id。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值： "mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DeleteSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 日志导出任务Id列表，接口会忽略不存在或已删除的任务Id。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值： "mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DeleteSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityAuditLogExportTasksResponseParams `json:"Response"`
}

func (r *DeleteSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSqlFiltersRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 通过VerifyUserAccount获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 限流任务ID列表。
	FilterIds []*int64 `json:"FilterIds,omitempty" name:"FilterIds"`
}

type DeleteSqlFiltersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 通过VerifyUserAccount获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 限流任务ID列表。
	FilterIds []*int64 `json:"FilterIds,omitempty" name:"FilterIds"`
}

func (r *DeleteSqlFiltersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSqlFiltersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionToken")
	delete(f, "FilterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSqlFiltersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSqlFiltersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSqlFiltersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSqlFiltersResponseParams `json:"Response"`
}

func (r *DeleteSqlFiltersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSqlFiltersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactRequestParams struct {
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 联系人名数组，支持模糊搜索。
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type DescribeAllUserContactRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 联系人名数组，支持模糊搜索。
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DescribeAllUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactResponseParams struct {
	// 联系人的总数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 联系人的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contacts []*ContactItem `json:"Contacts,omitempty" name:"Contacts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllUserContactResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserContactResponseParams `json:"Response"`
}

func (r *DescribeAllUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupRequestParams struct {
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 联系组名称数组，支持模糊搜索。
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type DescribeAllUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 联系组名称数组，支持模糊搜索。
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DescribeAllUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupResponseParams struct {
	// 组总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 组信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GroupItem `json:"Groups,omitempty" name:"Groups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserGroupResponseParams `json:"Response"`
}

func (r *DescribeAllUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 事件 ID 。通过“获取实例诊断历史DescribeDBDiagHistory”获取。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 事件 ID 。通过“获取实例诊断历史DescribeDBDiagHistory”获取。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBDiagEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventResponseParams struct {
	// 诊断项。
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// 诊断类型。
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// 事件 ID 。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 诊断事件详情，若无附加解释信息则输出为空。
	Explanation *string `json:"Explanation,omitempty" name:"Explanation"`

	// 诊断概要。
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// 诊断出的问题。
	Problem *string `json:"Problem,omitempty" name:"Problem"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 诊断建议，若无建议则输出为空。
	Suggestions *string `json:"Suggestions,omitempty" name:"Suggestions"`

	// 保留字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventsRequestParams struct {
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 风险等级列表，取值按影响程度从高至低分别为：1 - 致命、2 -严重、3 - 告警、4 - 提示、5 -健康。
	Severities []*int64 `json:"Severities,omitempty" name:"Severities"`

	// 实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDBDiagEventsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 风险等级列表，取值按影响程度从高至低分别为：1 - 致命、2 -严重、3 - 告警、4 - 提示、5 -健康。
	Severities []*int64 `json:"Severities,omitempty" name:"Severities"`

	// 实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDBDiagEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Severities")
	delete(f, "InstanceIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventsResponseParams struct {
	// 诊断事件的总数目。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 诊断事件的列表。
	Items []*DiagHistoryEventItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventsResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-11 12:13:14”，结束时间与开始时间的间隔最大可为2天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-11 12:13:14”，结束时间与开始时间的间隔最大可为2天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBDiagHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryResponseParams struct {
	// 事件描述。
	Events []*DiagHistoryEventItem `json:"Events,omitempty" name:"Events"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagHistoryResponseParams `json:"Response"`
}

func (r *DescribeDBDiagHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksRequestParams struct {
	// 第一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例ID数组，用于筛选指定实例的任务列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 定时生成；"MANUAL" - 手动触发。
	Sources []*string `json:"Sources,omitempty" name:"Sources"`

	// 报告的健康等级，支持的取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK" - 危险；"HIGH_RISK" - 高危。
	HealthLevels *string `json:"HealthLevels,omitempty" name:"HealthLevels"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatuses *string `json:"TaskStatuses,omitempty" name:"TaskStatuses"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 第一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例ID数组，用于筛选指定实例的任务列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 定时生成；"MANUAL" - 手动触发。
	Sources []*string `json:"Sources,omitempty" name:"Sources"`

	// 报告的健康等级，支持的取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK" - 危险；"HIGH_RISK" - 高危。
	HealthLevels *string `json:"HealthLevels,omitempty" name:"HealthLevels"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatuses *string `json:"TaskStatuses,omitempty" name:"TaskStatuses"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBDiagReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceIds")
	delete(f, "Sources")
	delete(f, "HealthLevels")
	delete(f, "TaskStatuses")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksResponseParams struct {
	// 任务总数目。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务列表。
	Tasks []*HealthReportTask `json:"Tasks,omitempty" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagReportTasksResponseParams `json:"Response"`
}

func (r *DescribeDBDiagReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间段天数，截止日期为当日，默认为7天。
	RangeDays *int64 `json:"RangeDays,omitempty" name:"RangeDays"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间段天数，截止日期为当日，默认为7天。
	RangeDays *int64 `json:"RangeDays,omitempty" name:"RangeDays"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeDBSpaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RangeDays")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSpaceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusResponseParams struct {
	// 磁盘增长量(MB)。
	Growth *int64 `json:"Growth,omitempty" name:"Growth"`

	// 磁盘剩余(MB)。
	Remain *int64 `json:"Remain,omitempty" name:"Remain"`

	// 磁盘总量(MB)。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 预计可用天数。
	AvailableDays *int64 `json:"AvailableDays,omitempty" name:"AvailableDays"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSpaceStatusResponseParams `json:"Response"`
}

func (r *DescribeDBSpaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesRequestParams struct {
	// 是否是DBbrain支持的实例，固定传 true。
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页参数，偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，分页值，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 根据实例名称条件查询。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// 根据实例ID条件查询。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 根据地域条件查询。
	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 是否是DBbrain支持的实例，固定传 true。
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页参数，偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，分页值，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 根据实例名称条件查询。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// 根据实例ID条件查询。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 根据地域条件查询。
	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

func (r *DescribeDiagDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsSupported")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceNames")
	delete(f, "InstanceIds")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesResponseParams struct {
	// 实例总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 全实例巡检状态：0：开启全实例巡检；1：未开启全实例巡检。
	DbScanStatus *int64 `json:"DbScanStatus,omitempty" name:"DbScanStatus"`

	// 实例相关信息。
	Items []*InstanceInfo `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDiagDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreRequestParams struct {
	// 需要获取健康得分的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 获取健康得分的时间，时间格式如：2019-09-10 12:13:14。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeHealthScoreRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取健康得分的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 获取健康得分的时间，时间格式如：2019-09-10 12:13:14。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeHealthScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHealthScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreResponseParams struct {
	// 健康得分以及异常扣分项。
	Data *HealthScoreInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHealthScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHealthScoreResponseParams `json:"Response"`
}

func (r *DescribeHealthScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileRequestParams struct {
	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单位，最大支持50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 根据邮件配置名称查询，定期发送的邮件配置名称遵循："scheduler_"+{instanceId}的规则。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`
}

type DescribeMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单位，最大支持50。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 根据邮件配置名称查询，定期发送的邮件配置名称遵循："scheduler_"+{instanceId}的规则。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`
}

func (r *DescribeMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProfileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileResponseParams struct {
	// 邮件配置详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileList []*UserProfile `json:"ProfileList,omitempty" name:"ProfileList"`

	// 邮件配置总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMailProfileResponseParams `json:"Response"`
}

func (r *DescribeMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 线程的操作账号名，用于筛选线程列表。
	User *string `json:"User,omitempty" name:"User"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表。
	DB *string `json:"DB,omitempty" name:"DB"`

	// 线程的操作状态，用于筛选线程列表。
	State *string `json:"State,omitempty" name:"State"`

	// 线程的执行类型，用于筛选线程列表。
	Command *string `json:"Command,omitempty" name:"Command"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 线程的操作语句，用于筛选线程列表。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeMySqlProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 线程的操作账号名，用于筛选线程列表。
	User *string `json:"User,omitempty" name:"User"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表。
	DB *string `json:"DB,omitempty" name:"DB"`

	// 线程的操作状态，用于筛选线程列表。
	State *string `json:"State,omitempty" name:"State"`

	// 线程的执行类型，用于筛选线程列表。
	Command *string `json:"Command,omitempty" name:"Command"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 线程的操作语句，用于筛选线程列表。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeMySqlProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ID")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "State")
	delete(f, "Command")
	delete(f, "Time")
	delete(f, "Info")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySqlProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListResponseParams struct {
	// 实时线程列表。
	ProcessList []*MySqlProcess `json:"ProcessList,omitempty" name:"ProcessList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMySqlProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySqlProcessListResponseParams `json:"Response"`
}

func (r *DescribeMySqlProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoPrimaryKeyTablesRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早为30天前的日期。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeNoPrimaryKeyTablesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早为30天前的日期。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeNoPrimaryKeyTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoPrimaryKeyTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNoPrimaryKeyTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoPrimaryKeyTablesResponseParams struct {
	// 无主键表总数。
	NoPrimaryKeyTableCount *int64 `json:"NoPrimaryKeyTableCount,omitempty" name:"NoPrimaryKeyTableCount"`

	// 与昨日扫描无主键表的差值，正数为增加，负数为减少，0为无变化。
	NoPrimaryKeyTableCountDiff *int64 `json:"NoPrimaryKeyTableCountDiff,omitempty" name:"NoPrimaryKeyTableCountDiff"`

	// 记录的无主键表总数（不超过无主键表总数），可用于分页查询。
	NoPrimaryKeyTableRecordCount *int64 `json:"NoPrimaryKeyTableRecordCount,omitempty" name:"NoPrimaryKeyTableRecordCount"`

	// 无主键表列表。
	NoPrimaryKeyTables []*Table `json:"NoPrimaryKeyTables,omitempty" name:"NoPrimaryKeyTables"`

	// 采集时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNoPrimaryKeyTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNoPrimaryKeyTablesResponseParams `json:"Response"`
}

func (r *DescribeNoPrimaryKeyTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoPrimaryKeyTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyProcessStatisticsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 该实例下需要查询的某一个 ProxyID 。
	InstanceProxyId *string `json:"InstanceProxyId,omitempty" name:"InstanceProxyId"`

	// 返回数量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照某字断排序。支持值包括："AllConn"，"ActiveConn"，"Ip"。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方向。支持值包括："DESC"，"ASC"。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeProxyProcessStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 该实例下需要查询的某一个 ProxyID 。
	InstanceProxyId *string `json:"InstanceProxyId,omitempty" name:"InstanceProxyId"`

	// 返回数量。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照某字断排序。支持值包括："AllConn"，"ActiveConn"，"Ip"。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方向。支持值包括："DESC"，"ASC"。
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeProxyProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyProcessStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceProxyId")
	delete(f, "Limit")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyProcessStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyProcessStatisticsResponseParams struct {
	// 实时会话统计详情。
	ProcessStatistics *ProcessStatistic `json:"ProcessStatistics,omitempty" name:"ProcessStatistics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxyProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyProcessStatisticsResponseParams `json:"Response"`
}

func (r *DescribeProxyProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// kill 会话异步任务 ID,  接口 CreateProxySessionKillTask 调用成功后获取。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeProxySessionKillTasksRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// kill 会话异步任务 ID,  接口 CreateProxySessionKillTask 调用成功后获取。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeProxySessionKillTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySessionKillTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksResponseParams struct {
	// kill 任务的详情。
	Tasks []*TaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 任务总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProxySessionKillTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySessionKillTasksResponseParams `json:"Response"`
}

func (r *DescribeProxySessionKillTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopBigKeysRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早可为前30天的日期。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 排序字段，取值包括Capacity - 内存，ItemCount - 元素数量，默认为Capacity。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// key类型筛选条件，默认为不进行筛选，取值包括string, list, set, hash, sortedset, stream。
	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRedisTopBigKeysRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早可为前30天的日期。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 排序字段，取值包括Capacity - 内存，ItemCount - 元素数量，默认为Capacity。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// key类型筛选条件，默认为不进行筛选，取值包括string, list, set, hash, sortedset, stream。
	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRedisTopBigKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopBigKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Product")
	delete(f, "SortBy")
	delete(f, "KeyType")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopBigKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopBigKeysResponseParams struct {
	// top key列表。
	TopKeys []*RedisKeySpaceData `json:"TopKeys,omitempty" name:"TopKeys"`

	// 采集时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRedisTopBigKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopBigKeysResponseParams `json:"Response"`
}

func (r *DescribeRedisTopBigKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopBigKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早可为前30天的日期。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRedisTopKeyPrefixListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早可为前30天的日期。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRedisTopKeyPrefixListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Product")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopKeyPrefixListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListResponseParams struct {
	// top key前缀列表。
	Items []*RedisPreKeySpaceData `json:"Items,omitempty" name:"Items"`

	// 采集时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRedisTopKeyPrefixListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopKeyPrefixListResponseParams `json:"Response"`
}

func (r *DescribeRedisTopKeyPrefixListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeSecurityAuditLogDownloadUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogDownloadUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsResponseParams struct {
	// 导出结果的COS链接列表。当结果集很大时，可能会切分为多个url下载。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogDownloadUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogDownloadUrlsResponseParams `json:"Response"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 日志导出任务Id列表。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitempty" name:"SecAuditGroupId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 日志导出任务Id列表。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitempty" name:"AsyncRequestIds"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "Product")
	delete(f, "AsyncRequestIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksResponseParams struct {
	// 安全审计日志导出任务列表。
	Tasks []*SecLogExportTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 安全审计日志导出任务总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogExportTasksResponseParams `json:"Response"`
}

func (r *DescribeSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-10 12:13:14”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-10 12:13:14”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTimeSeriesStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsResponseParams struct {
	// 柱间单位时间间隔，单位为秒。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 单位时间间隔内慢日志数量统计。
	TimeSeries []*TimeSlice `json:"TimeSeries,omitempty" name:"TimeSeries"`

	// 单位时间间隔内的实例 cpu 利用率监控数据。
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTimeSeriesStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键，默认为QueryTime。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序），默认为DESC。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据库名称数组。
	SchemaList []*SchemaItem `json:"SchemaList,omitempty" name:"SchemaList"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键，默认为QueryTime。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序），默认为DESC。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 数据库名称数组。
	SchemaList []*SchemaItem `json:"SchemaList,omitempty" name:"SchemaList"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SchemaList")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTopSqlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 慢日志 top sql 列表
	Rows []*SlowLogTopSqlItem `json:"Rows,omitempty" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTopSqlsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询范围的开始时间，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询范围的结束时间，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// SOL模板的MD5值
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type DescribeSlowLogUserHostStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询范围的开始时间，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询范围的结束时间，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// SOL模板的MD5值
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DescribeSlowLogUserHostStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "Md5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogUserHostStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsResponseParams struct {
	// 来源地址数目。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 各来源地址的慢日志占比详情列表。
	Items []*SlowLogHost `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogUserHostStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogUserHostStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlFiltersRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务ID列表，用于筛选任务列表。
	FilterIds []*int64 `json:"FilterIds,omitempty" name:"FilterIds"`

	// 任务状态列表，用于筛选任务列表，取值包括RUNNING - 运行中, FINISHED - 已完成, TERMINATED - 已终止。
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSqlFiltersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务ID列表，用于筛选任务列表。
	FilterIds []*int64 `json:"FilterIds,omitempty" name:"FilterIds"`

	// 任务状态列表，用于筛选任务列表，取值包括RUNNING - 运行中, FINISHED - 已完成, TERMINATED - 已终止。
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSqlFiltersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlFiltersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilterIds")
	delete(f, "Statuses")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSqlFiltersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlFiltersResponseParams struct {
	// 限流任务总数目。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 限流任务列表。
	Items []*SQLFilter `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSqlFiltersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSqlFiltersResponseParams `json:"Response"`
}

func (r *DescribeSqlFiltersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlFiltersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlTemplateRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名。
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeSqlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名。
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeSqlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Schema")
	delete(f, "SqlText")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSqlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlTemplateResponseParams struct {
	// 数据库名。
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// SQL类型。
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// SQL模版内容。
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// SQL模版ID。
	SqlId *int64 `json:"SqlId,omitempty" name:"SqlId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSqlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSqlTemplateResponseParams `json:"Response"`
}

func (r *DescribeSqlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemaTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemaTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesResponseParams struct {
	// 返回的Top库空间统计信息的时序数据列表。
	TopSpaceSchemaTimeSeries []*SchemaSpaceTimeSeries `json:"TopSpaceSchemaTimeSeries,omitempty" name:"TopSpaceSchemaTimeSeries"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemaTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemaTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemasRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasResponseParams struct {
	// 返回的Top库空间统计信息列表。
	TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitempty" name:"TopSpaceSchemas"`

	// 采集库空间数据的时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemasResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTableTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesResponseParams struct {
	// 返回的Top表空间统计信息的时序数据列表。
	TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitempty" name:"TopSpaceTableTimeSeries"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTableTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesResponseParams struct {
	// 返回的Top表空间统计信息列表。
	TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitempty" name:"TopSpaceTables"`

	// 采集表空间数据的时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTablesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL；"dbbrain-mysql" - 自建 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL；"dbbrain-mysql" - 自建 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeUserSqlAdviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SqlText")
	delete(f, "Schema")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceResponseParams struct {
	// SQL优化建议，可解析为JSON数组，无需优化时输出为空。
	Advices *string `json:"Advices,omitempty" name:"Advices"`

	// SQL优化建议备注，可解析为String数组，无需优化时输出为空。
	Comments *string `json:"Comments,omitempty" name:"Comments"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// 相关表的DDL信息，可解析为JSON数组。
	Tables *string `json:"Tables,omitempty" name:"Tables"`

	// SQL执行计划，可解析为JSON，无需优化时输出为空。
	SqlPlan *string `json:"SqlPlan,omitempty" name:"SqlPlan"`

	// SQL优化后的成本节约详情，可解析为JSON，无需优化时输出为空。
	Cost *string `json:"Cost,omitempty" name:"Cost"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSqlAdviceResponseParams `json:"Response"`
}

func (r *DescribeUserSqlAdviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {
	// 诊断类型。
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事件唯一ID 。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// 诊断概要。
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// 诊断项说明。
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 保留字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 地域。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type EventInfo struct {
	// 事件 ID 。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 诊断类型。
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 概要。
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// 扣分。
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`

	// 保留字段。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 告警数目。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type GroupItem struct {
	// 组id。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 组名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 组成员数量。
	MemberCount *int64 `json:"MemberCount,omitempty" name:"MemberCount"`
}

type HealthReportTask struct {
	// 异步任务请求 ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 定时生成；"MANUAL" - 手动触发。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 任务完成进度，单位%。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务开始执行时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务完成执行时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务所属实例的基础信息。
	InstanceInfo *InstanceBasicInfo `json:"InstanceInfo,omitempty" name:"InstanceInfo"`

	// 健康报告中的健康信息。
	HealthStatus *HealthStatus `json:"HealthStatus,omitempty" name:"HealthStatus"`
}

type HealthScoreInfo struct {
	// 异常详情。
	IssueTypes []*IssueTypeInfo `json:"IssueTypes,omitempty" name:"IssueTypes"`

	// 异常事件总数。
	EventsTotalCount *int64 `json:"EventsTotalCount,omitempty" name:"EventsTotalCount"`

	// 健康得分。
	HealthScore *int64 `json:"HealthScore,omitempty" name:"HealthScore"`

	// 健康等级, 如："HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK"。
	HealthLevel *string `json:"HealthLevel,omitempty" name:"HealthLevel"`
}

type HealthStatus struct {
	// 健康分数，满分100。
	HealthScore *int64 `json:"HealthScore,omitempty" name:"HealthScore"`

	// 健康等级，取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK"- 危险；"HIGH_RISK" - 高危。
	HealthLevel *string `json:"HealthLevel,omitempty" name:"HealthLevel"`

	// 总扣分分数。
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`

	// 扣分详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScoreDetails []*ScoreDetail `json:"ScoreDetails,omitempty" name:"ScoreDetails"`
}

type InstanceBasicInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例内网IP。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例内网Port。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 实例产品。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例引擎版本。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
}

type InstanceConfs struct {
	// 数据库巡检开关, Yes/No。
	DailyInspection *string `json:"DailyInspection,omitempty" name:"DailyInspection"`

	// 实例概览开关，Yes/No。
	OverviewDisplay *string `json:"OverviewDisplay,omitempty" name:"OverviewDisplay"`
}

type InstanceInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例所属地域。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 健康得分。
	HealthScore *int64 `json:"HealthScore,omitempty" name:"HealthScore"`

	// 所属产品。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 异常事件数量。
	EventCount *int64 `json:"EventCount,omitempty" name:"EventCount"`

	// 实例类型：1:MASTER；2:DR，3：RO，4:SDR。
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 核心数。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，单位MB。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 硬盘存储，单位GB。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 数据库版本。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 内网地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 内网端口。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 接入来源。
	Source *string `json:"Source,omitempty" name:"Source"`

	// 分组ID。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组组名。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例状态：0：发货中；1：运行正常；4：销毁中；5：隔离中。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 子网统一ID。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// cdb类型。
	DeployMode *string `json:"DeployMode,omitempty" name:"DeployMode"`

	// cdb实例初始化标志：0：未初始化；1：已初始化。
	InitFlag *int64 `json:"InitFlag,omitempty" name:"InitFlag"`

	// 任务状态。
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 私有网络统一ID。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 实例巡检/概览的状态。
	InstanceConf *InstanceConfs `json:"InstanceConf,omitempty" name:"InstanceConf"`

	// 资源到期时间。
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// 是否是DBbrain支持的实例。
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// 实例安全审计日志开启状态：ON： 安全审计开启；OFF： 未开启安全审计。
	SecAuditStatus *string `json:"SecAuditStatus,omitempty" name:"SecAuditStatus"`

	// 实例审计日志开启状态，ALL_AUDIT： 开启全审计；RULE_AUDIT： 开启规则审计；UNBOUND： 未开启审计。
	AuditPolicyStatus *string `json:"AuditPolicyStatus,omitempty" name:"AuditPolicyStatus"`

	// 实例审计日志运行状态：normal： 运行中； paused： 欠费暂停。
	AuditRunningStatus *string `json:"AuditRunningStatus,omitempty" name:"AuditRunningStatus"`
}

type IssueTypeInfo struct {
	// 指标分类：AVAILABILITY：可用性，MAINTAINABILITY：可维护性，PERFORMANCE，性能，RELIABILITY可靠性。
	IssueType *string `json:"IssueType,omitempty" name:"IssueType"`

	// 异常事件。
	Events []*EventInfo `json:"Events,omitempty" name:"Events"`

	// 异常事件总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type KillMySqlThreadsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// kill会话任务的阶段，取值包括："Prepare"-准备阶段，"Commit"-提交阶段。
	Stage *string `json:"Stage,omitempty" name:"Stage"`

	// 需要kill的sql会话ID列表，此参数用于Prepare阶段。
	Threads []*int64 `json:"Threads,omitempty" name:"Threads"`

	// 执行ID，此参数用于Commit阶段。
	SqlExecId *string `json:"SqlExecId,omitempty" name:"SqlExecId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type KillMySqlThreadsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// kill会话任务的阶段，取值包括："Prepare"-准备阶段，"Commit"-提交阶段。
	Stage *string `json:"Stage,omitempty" name:"Stage"`

	// 需要kill的sql会话ID列表，此参数用于Prepare阶段。
	Threads []*int64 `json:"Threads,omitempty" name:"Threads"`

	// 执行ID，此参数用于Commit阶段。
	SqlExecId *string `json:"SqlExecId,omitempty" name:"SqlExecId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *KillMySqlThreadsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMySqlThreadsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Stage")
	delete(f, "Threads")
	delete(f, "SqlExecId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillMySqlThreadsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillMySqlThreadsResponseParams struct {
	// kill完成的sql会话ID列表。
	Threads []*int64 `json:"Threads,omitempty" name:"Threads"`

	// 执行ID， Prepare阶段的任务输出，用于Commit阶段中指定执行kill操作的会话ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlExecId *string `json:"SqlExecId,omitempty" name:"SqlExecId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type KillMySqlThreadsResponse struct {
	*tchttp.BaseResponse
	Response *KillMySqlThreadsResponseParams `json:"Response"`
}

func (r *KillMySqlThreadsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMySqlThreadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MailConfiguration struct {
	// 是否开启邮件发送: 0, 否; 1, 是。
	SendMail *int64 `json:"SendMail,omitempty" name:"SendMail"`

	// 地域配置, 如["ap-guangzhou", "ap-shanghai"]。巡检的邮件发送模板，配置需要发送巡检邮件的地域；订阅的邮件发送模板，配置当前订阅实例的所属地域。
	Region []*string `json:"Region,omitempty" name:"Region"`

	// 发送指定的健康等级的报告, 如["HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK"]。
	HealthStatus []*string `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// 联系人id, 联系人/联系组不能都为空。
	ContactPerson []*int64 `json:"ContactPerson,omitempty" name:"ContactPerson"`

	// 联系组id, 联系人/联系组不能都为空。
	ContactGroup []*int64 `json:"ContactGroup,omitempty" name:"ContactGroup"`
}

// Predefined struct for user
type ModifyDiagDBInstanceConfRequestParams struct {
	// 实例配置，包括巡检、概览开关等。
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitempty" name:"InstanceConfs"`

	// 生效实例地域，取值为"All"，代表全地域。
	Regions *string `json:"Regions,omitempty" name:"Regions"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 指定更改巡检状态的实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type ModifyDiagDBInstanceConfRequest struct {
	*tchttp.BaseRequest
	
	// 实例配置，包括巡检、概览开关等。
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitempty" name:"InstanceConfs"`

	// 生效实例地域，取值为"All"，代表全地域。
	Regions *string `json:"Regions,omitempty" name:"Regions"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 指定更改巡检状态的实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *ModifyDiagDBInstanceConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceConfs")
	delete(f, "Regions")
	delete(f, "Product")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiagDBInstanceConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiagDBInstanceConfResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDiagDBInstanceConfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiagDBInstanceConfResponseParams `json:"Response"`
}

func (r *ModifyDiagDBInstanceConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySqlFiltersRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 通过VerifyUserAccount获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// SQL限流任务ID列表。
	FilterIds []*int64 `json:"FilterIds,omitempty" name:"FilterIds"`

	// 限流任务状态，取值支持TERMINATED - 终止。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type ModifySqlFiltersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 通过VerifyUserAccount获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// SQL限流任务ID列表。
	FilterIds []*int64 `json:"FilterIds,omitempty" name:"FilterIds"`

	// 限流任务状态，取值支持TERMINATED - 终止。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *ModifySqlFiltersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySqlFiltersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionToken")
	delete(f, "FilterIds")
	delete(f, "Status")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySqlFiltersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySqlFiltersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySqlFiltersResponse struct {
	*tchttp.BaseResponse
	Response *ModifySqlFiltersResponseParams `json:"Response"`
}

func (r *ModifySqlFiltersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySqlFiltersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorFloatMetric struct {
	// 指标名称。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type MonitorFloatMetricSeriesData struct {
	// 监控指标。
	Series []*MonitorFloatMetric `json:"Series,omitempty" name:"Series"`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type MonitorMetric struct {
	// 指标名称。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type MonitorMetricSeriesData struct {
	// 监控指标。
	Series []*MonitorMetric `json:"Series,omitempty" name:"Series"`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type MySqlProcess struct {
	// 线程ID。
	ID *string `json:"ID,omitempty" name:"ID"`

	// 线程的操作账号名。
	User *string `json:"User,omitempty" name:"User"`

	// 线程的操作主机地址。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 线程的操作数据库。
	DB *string `json:"DB,omitempty" name:"DB"`

	// 线程的操作状态。
	State *string `json:"State,omitempty" name:"State"`

	// 线程的执行类型。
	Command *string `json:"Command,omitempty" name:"Command"`

	// 线程的操作时长，单位秒。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 线程的操作语句。
	Info *string `json:"Info,omitempty" name:"Info"`
}

type ProcessStatistic struct {
	// 会话详情数组。
	Items []*SessionItem `json:"Items,omitempty" name:"Items"`

	// 总连接数。
	AllConnSum *int64 `json:"AllConnSum,omitempty" name:"AllConnSum"`

	// 总活跃连接数。
	ActiveConnSum *int64 `json:"ActiveConnSum,omitempty" name:"ActiveConnSum"`
}

type ProfileInfo struct {
	// 语言, 如"zh"。
	Language *string `json:"Language,omitempty" name:"Language"`

	// 邮件模板的内容。
	MailConfiguration *MailConfiguration `json:"MailConfiguration,omitempty" name:"MailConfiguration"`
}

type RedisKeySpaceData struct {
	// key名。
	Key *string `json:"Key,omitempty" name:"Key"`

	// key类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// key编码方式。
	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`

	// key过期时间戳（毫秒），0代表未设置过期时间。
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// key内存大小，单位Byte。
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 元素个数。
	ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`

	// 最大元素长度。
	MaxElementSize *int64 `json:"MaxElementSize,omitempty" name:"MaxElementSize"`
}

type RedisPreKeySpaceData struct {
	// 平均元素长度。
	AveElementSize *int64 `json:"AveElementSize,omitempty" name:"AveElementSize"`

	// 总占用内存（Byte）。
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// key前缀。
	KeyPreIndex *string `json:"KeyPreIndex,omitempty" name:"KeyPreIndex"`

	// 元素数量。
	ItemCount *int64 `json:"ItemCount,omitempty" name:"ItemCount"`

	// key个数。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 最大元素长度。
	MaxElementSize *int64 `json:"MaxElementSize,omitempty" name:"MaxElementSize"`
}

type SQLFilter struct {
	// 任务ID。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 任务状态，取值包括RUNNING - 运行中, FINISHED - 已完成, TERMINATED - 已终止。
	Status *string `json:"Status,omitempty" name:"Status"`

	// SQL类型，取值包括SELECT, UPDATE, DELETE, INSERT, REPLACE。
	SqlType *string `json:"SqlType,omitempty" name:"SqlType"`

	// 筛选SQL的关键词，多个关键词用英文逗号拼接。
	OriginKeys *string `json:"OriginKeys,omitempty" name:"OriginKeys"`

	// 筛选SQL的规则。
	OriginRule *string `json:"OriginRule,omitempty" name:"OriginRule"`

	// 已拒绝SQL数目。
	RejectedSqlCount *int64 `json:"RejectedSqlCount,omitempty" name:"RejectedSqlCount"`

	// 当前并发数。
	CurrentConcurrency *int64 `json:"CurrentConcurrency,omitempty" name:"CurrentConcurrency"`

	// 最大并发数。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 当前时间。
	CurrentTime *string `json:"CurrentTime,omitempty" name:"CurrentTime"`

	// 限流过期时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type SchemaItem struct {
	// 数据库名称
	Schema *string `json:"Schema,omitempty" name:"Schema"`
}

type SchemaSpaceData struct {
	// 库名。
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// 数据空间（MB）。
	DataLength *float64 `json:"DataLength,omitempty" name:"DataLength"`

	// 索引空间（MB）。
	IndexLength *float64 `json:"IndexLength,omitempty" name:"IndexLength"`

	// 碎片空间（MB）。
	DataFree *float64 `json:"DataFree,omitempty" name:"DataFree"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`

	// 碎片率（%）。
	FragRatio *float64 `json:"FragRatio,omitempty" name:"FragRatio"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// 库中所有表对应的独立物理文件大小加和（MB）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitempty" name:"PhysicalFileSize"`
}

type SchemaSpaceTimeSeries struct {
	// 库名
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// 单位时间间隔内的空间指标数据。
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`
}

type ScoreDetail struct {
	// 扣分项分类，取值包括：可用性、可维护性、性能及可靠性。
	IssueType *string `json:"IssueType,omitempty" name:"IssueType"`

	// 扣分总分。
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`

	// 扣分总分上限。
	ScoreLostMax *int64 `json:"ScoreLostMax,omitempty" name:"ScoreLostMax"`

	// 扣分项列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ScoreItem `json:"Items,omitempty" name:"Items"`
}

type ScoreItem struct {
	// 异常诊断项名称。
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// 诊断项分类，取值包括：可用性、可维护性、性能及可靠性。
	IssueType *string `json:"IssueType,omitempty" name:"IssueType"`

	// 健康等级，取值包括：信息、提示、告警、严重、致命。
	TopSeverity *string `json:"TopSeverity,omitempty" name:"TopSeverity"`

	// 该异常诊断项出现次数。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 扣分分数。
	ScoreLost *int64 `json:"ScoreLost,omitempty" name:"ScoreLost"`
}

type SecLogExportTaskInfo struct {
	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 任务开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务执行进度。
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 导出日志开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogStartTime *string `json:"LogStartTime,omitempty" name:"LogStartTime"`

	// 导出日志结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogEndTime *string `json:"LogEndTime,omitempty" name:"LogEndTime"`

	// 日志文件总大小，单位KB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// 风险等级列表。0 无风险；1 低风险；2 中风险；3 高风险。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DangerLevels []*uint64 `json:"DangerLevels,omitempty" name:"DangerLevels"`
}

type SessionItem struct {
	// 访问来源。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 当前访问来源活跃连接数
	ActiveConn *string `json:"ActiveConn,omitempty" name:"ActiveConn"`

	// 当前访问来源总连接数
	AllConn *int64 `json:"AllConn,omitempty" name:"AllConn"`
}

type SlowLogHost struct {
	// 来源地址。
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// 该来源地址的慢日志数目占总数目的比例，单位%。
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`

	// 该来源地址的慢日志数目。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SlowLogTopSqlItem struct {
	// sql总锁等待时间，单位秒
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// 最大锁等待时间，单位秒
	LockTimeMax *float64 `json:"LockTimeMax,omitempty" name:"LockTimeMax"`

	// 最小锁等待时间，单位秒
	LockTimeMin *float64 `json:"LockTimeMin,omitempty" name:"LockTimeMin"`

	// 总扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// 最大扫描行数
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitempty" name:"RowsExaminedMax"`

	// 最小扫描行数
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitempty" name:"RowsExaminedMin"`

	// 总耗时，单位秒
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// 最大执行时间，单位秒
	QueryTimeMax *float64 `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`

	// 最小执行时间，单位秒
	QueryTimeMin *float64 `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`

	// 总返回行数
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// 最大返回行数
	RowsSentMax *int64 `json:"RowsSentMax,omitempty" name:"RowsSentMax"`

	// 最小返回行数
	RowsSentMin *int64 `json:"RowsSentMin,omitempty" name:"RowsSentMin"`

	// 执行次数
	ExecTimes *int64 `json:"ExecTimes,omitempty" name:"ExecTimes"`

	// sql模板
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// 带参数SQL（随机）
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 数据库名
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// 总耗时占比，单位%
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitempty" name:"QueryTimeRatio"`

	// sql总锁等待时间占比，单位%
	LockTimeRatio *float64 `json:"LockTimeRatio,omitempty" name:"LockTimeRatio"`

	// 总扫描行数占比，单位%
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitempty" name:"RowsExaminedRatio"`

	// 总返回行数占比，单位%
	RowsSentRatio *float64 `json:"RowsSentRatio,omitempty" name:"RowsSentRatio"`

	// 平均执行时间，单位秒
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitempty" name:"QueryTimeAvg"`

	// 平均返回行数
	RowsSentAvg *float64 `json:"RowsSentAvg,omitempty" name:"RowsSentAvg"`

	// 平均锁等待时间，单位秒
	LockTimeAvg *float64 `json:"LockTimeAvg,omitempty" name:"LockTimeAvg"`

	// 平均扫描行数
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitempty" name:"RowsExaminedAvg"`

	// SOL模板的MD5值
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type Table struct {
	// 库名。
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// 表名。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`
}

type TableSpaceData struct {
	// 表名。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 数据空间（MB）。
	DataLength *float64 `json:"DataLength,omitempty" name:"DataLength"`

	// 索引空间（MB）。
	IndexLength *float64 `json:"IndexLength,omitempty" name:"IndexLength"`

	// 碎片空间（MB）。
	DataFree *float64 `json:"DataFree,omitempty" name:"DataFree"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`

	// 碎片率（%）。
	FragRatio *float64 `json:"FragRatio,omitempty" name:"FragRatio"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// 表对应的独立物理文件大小（MB）。
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitempty" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {
	// 表名。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 单位时间间隔内的空间指标数据。
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`
}

type TaskInfo struct {
	// 异步任务 ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

	// 当前实例所有 proxy 列表。
	InstProxyList []*string `json:"InstProxyList,omitempty" name:"InstProxyList"`

	// 当前实例所有 proxy 数量。
	InstProxyCount *int64 `json:"InstProxyCount,omitempty" name:"InstProxyCount"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务启动时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 完成 kill 任务的 proxyId。
	FinishedProxyList []*string `json:"FinishedProxyList,omitempty" name:"FinishedProxyList"`

	// kill 任务实行失败的 proxyId。
	FailedProxyList []*string `json:"FailedProxyList,omitempty" name:"FailedProxyList"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务执行进度。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type TimeSlice struct {
	// 总数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 统计开始时间
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type UserProfile struct {
	// 配置的id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileId *string `json:"ProfileId,omitempty" name:"ProfileId"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileLevel *string `json:"ProfileLevel,omitempty" name:"ProfileLevel"`

	// 配置名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// 配置详情。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`
}

// Predefined struct for user
type VerifyUserAccountRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库账号名。
	User *string `json:"User,omitempty" name:"User"`

	// 数据库账号密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

type VerifyUserAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库账号名。
	User *string `json:"User,omitempty" name:"User"`

	// 数据库账号密码。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *VerifyUserAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyUserAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyUserAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyUserAccountResponseParams struct {
	// 会话token，有效期为5分钟。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyUserAccountResponse struct {
	*tchttp.BaseResponse
	Response *VerifyUserAccountResponseParams `json:"Response"`
}

func (r *VerifyUserAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyUserAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}