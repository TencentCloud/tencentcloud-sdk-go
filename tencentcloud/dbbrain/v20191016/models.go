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

package v20191016

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddUserContactRequest struct {
	*tchttp.BaseRequest

	// 联系人姓名，大小写字母+数字+下划线，最小 2 位最大 60 位的长度， 不能以"_"开头，且联系人名保持唯一。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 邮箱地址，大小写字母、数字及下划线组成， 不能以"_"开头。
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

type AddUserContactResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加成功的联系人id。
		Id *int64 `json:"Id,omitempty" name:"Id"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ContactItem struct {

	// 联系人id。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 联系人姓名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 联系人绑定的邮箱。
	Mail *string `json:"Mail,omitempty" name:"Mail"`
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

type CreateDBDiagReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AsyncRequestId *int64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateDBDiagReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 健康报告浏览地址。
		ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

		// 健康报告浏览地址到期时间戳（秒）。
		ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateSchedulerMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateSecurityAuditLogExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志导出任务Id。
		AsyncRequestId *uint64 `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAllUserContactResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 联系人的总数量。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 联系人的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Contacts []*ContactItem `json:"Contacts,omitempty" name:"Contacts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAllUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 组总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 组信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Groups []*GroupItem `json:"Groups,omitempty" name:"Groups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 诊断项。
		DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

		// 诊断类型。
		DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

		// 事件 ID 。
		EventId *int64 `json:"EventId,omitempty" name:"EventId"`

		// 事件详情。
		Explanation *string `json:"Explanation,omitempty" name:"Explanation"`

		// 概要。
		Outline *string `json:"Outline,omitempty" name:"Outline"`

		// 诊断出的问题。
		Problem *string `json:"Problem,omitempty" name:"Problem"`

		// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
		Severity *int64 `json:"Severity,omitempty" name:"Severity"`

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 建议。
		Suggestions *string `json:"Suggestions,omitempty" name:"Suggestions"`

		// 保留字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Metric *string `json:"Metric,omitempty" name:"Metric"`

		// 结束时间。
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件描述。
		Events []*DiagHistoryEventItem `json:"Events,omitempty" name:"Events"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 返回数量，默认20。
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

type DescribeDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务总数目。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务列表。
		Tasks []*HealthReportTask `json:"Tasks,omitempty" name:"Tasks"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 是否是DBbrain支持的实例，固定传 true。
	IsSupported *bool `json:"IsSupported,omitempty" name:"IsSupported"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 分页参数，偏移量。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，分页值。
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

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 全实例巡检状态：0：开启全实例巡检；1：未开启全实例巡检。
		DbScanStatus *int64 `json:"DbScanStatus,omitempty" name:"DbScanStatus"`

		// 实例相关信息。
		Items []*InstanceInfo `json:"Items,omitempty" name:"Items"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeHealthScoreRequest struct {
	*tchttp.BaseRequest

	// 需要获取健康得分的实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 获取健康得分的时间。
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

type DescribeHealthScoreResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 健康得分以及异常扣分项。
		Data *HealthScoreInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 邮件配置详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ProfileList []*UserProfile `json:"ProfileList,omitempty" name:"ProfileList"`

		// 邮件模版总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSecurityAuditLogDownloadUrlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 导出结果的COS链接列表。当结果集很大时，可能会切分为多个url下载。
		Urls []*string `json:"Urls,omitempty" name:"Urls"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 返回数量，默认20。
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

type DescribeSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全审计日志导出任务列表。
		Tasks []*SecLogExportTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

		// 安全审计日志导出任务总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 柱间单位时间间隔，单位为秒。
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// 单位时间间隔内慢日志数量统计。
		TimeSeries []*TimeSlice `json:"TimeSeries,omitempty" name:"TimeSeries"`

		// 单位时间间隔内的实例 cpu 利用率监控数据。
		SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-10 12:13:14”，截止时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序）。
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

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 慢日志 top sql 列表
		Rows []*SlowLogTopSqlItem `json:"Rows,omitempty" name:"Rows"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogUserHostStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 来源地址数目。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 各来源地址的慢日志占比详情列表。
		Items []*SlowLogHost `json:"Items,omitempty" name:"Items"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeTopSpaceSchemaTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Top库空间统计信息的时序数据列表。
		TopSpaceSchemaTimeSeries []*SchemaSpaceTimeSeries `json:"TopSpaceSchemaTimeSeries,omitempty" name:"TopSpaceSchemaTimeSeries"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Top库空间统计信息列表。
		TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitempty" name:"TopSpaceSchemas"`

		// 采集库空间数据的时间戳（秒）。
		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Top表空间统计信息的时序数据列表。
		TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitempty" name:"TopSpaceTableTimeSeries"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Top表空间统计信息列表。
		TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitempty" name:"TopSpaceTables"`

		// 采集表空间数据的时间戳（秒）。
		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitempty" name:"Schema"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SQL优化建议，可解析为JSON数组。
		Advices *string `json:"Advices,omitempty" name:"Advices"`

		// SQL优化建议备注，可解析为String数组。
		Comments *string `json:"Comments,omitempty" name:"Comments"`

		// SQL语句。
		SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

		// 库名。
		Schema *string `json:"Schema,omitempty" name:"Schema"`

		// 相关表的DDL信息，可解析为JSON数组。
		Tables *string `json:"Tables,omitempty" name:"Tables"`

		// SQL执行计划，可解析为JSON。
		SqlPlan *string `json:"SqlPlan,omitempty" name:"SqlPlan"`

		// SQL优化后的成本节约详情，可解析为JSON。
		Cost *string `json:"Cost,omitempty" name:"Cost"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 事件 ID 。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// 概要。
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// 诊断项。
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// 实例 ID 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 保留字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type MailConfiguration struct {

	// 是否开启邮件发送: 0, 否; 1, 是。
	SendMail *int64 `json:"SendMail,omitempty" name:"SendMail"`

	// 地域配置, 如["ap-guangzhou", "ap-shanghai"]。巡检的邮件发送模版，配置需要发送巡检邮件的地域；订阅的邮件发送模版，配置当前订阅实例的所属地域。
	Region []*string `json:"Region,omitempty" name:"Region"`

	// 发送指定的健康等级的报告, 如["HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK"]。
	HealthStatus []*string `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// 联系人id, 联系人/联系组不能都为空。
	ContactPerson []*int64 `json:"ContactPerson,omitempty" name:"ContactPerson"`

	// 联系组id, 联系人/联系组不能都为空。
	ContactGroup []*int64 `json:"ContactGroup,omitempty" name:"ContactGroup"`
}

type ModifyDiagDBInstanceConfRequest struct {
	*tchttp.BaseRequest

	// 巡检开关。
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

type ModifyDiagDBInstanceConfResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	Values []*int64 `json:"Values,omitempty" name:"Values"`
}

type MonitorMetricSeriesData struct {

	// 监控指标。
	Series []*MonitorMetric `json:"Series,omitempty" name:"Series"`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type ProfileInfo struct {

	// 语言, 如"zh"。
	Language *string `json:"Language,omitempty" name:"Language"`

	// 邮件模板的内容。
	MailConfiguration *MailConfiguration `json:"MailConfiguration,omitempty" name:"MailConfiguration"`
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

type SlowLogHost struct {

	// 来源地址。
	UserHost *string `json:"UserHost,omitempty" name:"UserHost"`

	// 该来源地址的慢日志数目占总数目的比例，单位%。
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`

	// 该来源地址的慢日志数目。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SlowLogTopSqlItem struct {

	// sql总锁等待时间
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// 最大锁等待时间
	LockTimeMax *float64 `json:"LockTimeMax,omitempty" name:"LockTimeMax"`

	// 最小锁等待时间
	LockTimeMin *float64 `json:"LockTimeMin,omitempty" name:"LockTimeMin"`

	// 总扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// 最大扫描行数
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitempty" name:"RowsExaminedMax"`

	// 最小扫描行数
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitempty" name:"RowsExaminedMin"`

	// 总耗时
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// 最大执行时间
	QueryTimeMax *float64 `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`

	// 最小执行时间
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

	// 总耗时占比
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitempty" name:"QueryTimeRatio"`

	// sql总锁等待时间占比
	LockTimeRatio *float64 `json:"LockTimeRatio,omitempty" name:"LockTimeRatio"`

	// 总扫描行数占比
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitempty" name:"RowsExaminedRatio"`

	// 总返回行数占比
	RowsSentRatio *float64 `json:"RowsSentRatio,omitempty" name:"RowsSentRatio"`

	// 平均执行时间
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitempty" name:"QueryTimeAvg"`

	// 平均返回行数
	RowsSentAvg *float64 `json:"RowsSentAvg,omitempty" name:"RowsSentAvg"`

	// 平均锁等待时间
	LockTimeAvg *float64 `json:"LockTimeAvg,omitempty" name:"LockTimeAvg"`

	// 平均扫描行数
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitempty" name:"RowsExaminedAvg"`
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

	// 配置类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileType *string `json:"ProfileType,omitempty" name:"ProfileType"`

	// 配置级别，"User"或"Instance"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileLevel *string `json:"ProfileLevel,omitempty" name:"ProfileLevel"`

	// 配置名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfileName *string `json:"ProfileName,omitempty" name:"ProfileName"`

	// 配置详情。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitempty" name:"ProfileInfo"`
}
