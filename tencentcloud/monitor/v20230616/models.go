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

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIWorkbenchSREDigitalTwinTask struct {
	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务配置
	TaskConfig *string `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`

	// 唯一标识
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 所属数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`
}

type AIWorkbenchSREDigitalTwinTaskList struct {
	// 任务列表
	Tasks []*AIWorkbenchSREDigitalTwinTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type AIWorkbenchSREDigitalTwinWorkLog struct {
	// 唯一标识符
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 所属数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 所属数字分身任务ID
	TaskID *uint64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 分析时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 分析状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 分析结果摘要
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 所属任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 所属任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type AIWorkbenchSREDigitalTwinWorkLogDetail struct {
	// 工作日志详细内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 工作日志任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 工作日志相关对话ID
	DialogID *int64 `json:"DialogID,omitnil,omitempty" name:"DialogID"`
}

type AIWorkbenchSREDigitalTwinWorkLogList struct {
	// 工作日志列表
	WorkLogs []*AIWorkbenchSREDigitalTwinWorkLog `json:"WorkLogs,omitnil,omitempty" name:"WorkLogs"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

// Predefined struct for user
type CreateNoticeContentTmplRequestParams struct {
	// 模板名称
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 模板内容
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// 模板语言 en/zh
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`
}

type CreateNoticeContentTmplRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 模板内容
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// 模板语言 en/zh
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`
}

func (r *CreateNoticeContentTmplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNoticeContentTmplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TmplName")
	delete(f, "MonitorType")
	delete(f, "TmplContents")
	delete(f, "TmplLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNoticeContentTmplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNoticeContentTmplResponseParams struct {
	// 自定义内容模板ID
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNoticeContentTmplResponse struct {
	*tchttp.BaseResponse
	Response *CreateNoticeContentTmplResponseParams `json:"Response"`
}

func (r *CreateNoticeContentTmplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNoticeContentTmplsRequestParams struct {
	// 要删除的模板id
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`
}

type DeleteNoticeContentTmplsRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的模板id
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`
}

func (r *DeleteNoticeContentTmplsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNoticeContentTmplsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TmplIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNoticeContentTmplsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNoticeContentTmplsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNoticeContentTmplsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNoticeContentTmplsResponseParams `json:"Response"`
}

func (r *DeleteNoticeContentTmplsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNoticeContentTmplsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinTaskListRequestParams struct {
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAIWorkbenchSREDigitalTwinTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TwinID")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSREDigitalTwinTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinTaskListResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身任务列表
	Data *AIWorkbenchSREDigitalTwinTaskList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSREDigitalTwinTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSREDigitalTwinTaskListResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequestParams struct {
	// 工作日志ID
	WorkLogID *int64 `json:"WorkLogID,omitnil,omitempty" name:"WorkLogID"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest struct {
	*tchttp.BaseRequest
	
	// 工作日志ID
	WorkLogID *int64 `json:"WorkLogID,omitnil,omitempty" name:"WorkLogID"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkLogID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身详细信息
	Data *AIWorkbenchSREDigitalTwinWorkLogDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogListRequestParams struct {
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest struct {
	*tchttp.BaseRequest
	
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TwinID")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogListResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身工作日志列表
	Data *AIWorkbenchSREDigitalTwinWorkLogList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSREDigitalTwinWorkLogListResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesRequestParams struct {
	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 起始时间点，unix秒级时间戳
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// 从 QueryBaseTime 开始，需要查询往前多久的时间，单位秒
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// 分页参数
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// 当监控类型为 MT_QCE 时候需要填写，归属的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 当监控类型为 MT_QCE 时候需要填写， 告警策略类型
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询某个策略的通知历史
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeAlarmNotifyHistoriesRequest struct {
	*tchttp.BaseRequest
	
	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 起始时间点，unix秒级时间戳
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// 从 QueryBaseTime 开始，需要查询往前多久的时间，单位秒
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// 分页参数
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// 当监控类型为 MT_QCE 时候需要填写，归属的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 当监控类型为 MT_QCE 时候需要填写， 告警策略类型
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询某个策略的通知历史
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmNotifyHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorType")
	delete(f, "QueryBaseTime")
	delete(f, "QueryBeforeSeconds")
	delete(f, "PageParams")
	delete(f, "Namespace")
	delete(f, "ModelName")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNotifyHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNotifyHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNotifyHistoriesResponseParams `json:"Response"`
}

func (r *DescribeAlarmNotifyHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeContentTmplRequestParams struct {
	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 指定模板ID查询，查询参数都为空则默认查询账号下所有模板
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`

	// 指定模板名称查询，查询参数都为空则默认查询账号下所有模板
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 指定通知模板ID查询，查询参数都为空则默认查询账号下所有模板
	NoticeID *string `json:"NoticeID,omitnil,omitempty" name:"NoticeID"`

	// 模板语言 en/zh 缺省不过滤
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

type DescribeNoticeContentTmplRequest struct {
	*tchttp.BaseRequest
	
	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 指定模板ID查询，查询参数都为空则默认查询账号下所有模板
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`

	// 指定模板名称查询，查询参数都为空则默认查询账号下所有模板
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 指定通知模板ID查询，查询参数都为空则默认查询账号下所有模板
	NoticeID *string `json:"NoticeID,omitnil,omitempty" name:"NoticeID"`

	// 模板语言 en/zh 缺省不过滤
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

func (r *DescribeNoticeContentTmplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeContentTmplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "TmplIDs")
	delete(f, "TmplName")
	delete(f, "NoticeID")
	delete(f, "TmplLanguage")
	delete(f, "MonitorType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNoticeContentTmplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeContentTmplResponseParams struct {
	// 自定义通知内容模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeContentTmpls []*NoticeContentTmpl `json:"NoticeContentTmpls,omitnil,omitempty" name:"NoticeContentTmpls"`

	// 通知内容模板绑定的告警策略数量
	NoticeContentTmplBindPolicyCounts []*NoticeContentTmplBindPolicyCount `json:"NoticeContentTmplBindPolicyCounts,omitnil,omitempty" name:"NoticeContentTmplBindPolicyCounts"`

	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNoticeContentTmplResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNoticeContentTmplResponseParams `json:"Response"`
}

func (r *DescribeNoticeContentTmplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DingDingRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`

	// 标题模板
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type DingDingRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *DingDingRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type FeiShuRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`

	// 标题模板
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type FeiShuRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *FeiShuRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type GoogleChatRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`
}

type GoogleChatRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *GoogleChatRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

// Predefined struct for user
type ModifyNoticeContentTmplRequestParams struct {
	// 模板名称
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 模板内容
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// 需要修改的模板ID
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`
}

type ModifyNoticeContentTmplRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 模板内容
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// 需要修改的模板ID
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`
}

func (r *ModifyNoticeContentTmplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNoticeContentTmplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TmplName")
	delete(f, "TmplContents")
	delete(f, "TmplID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNoticeContentTmplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNoticeContentTmplResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNoticeContentTmplResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNoticeContentTmplResponseParams `json:"Response"`
}

func (r *ModifyNoticeContentTmplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeContentTmpl struct {
	// 自定义通知内容模板id，唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`

	// 自定义通知内容模板名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 通知内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// Unix时间戳，秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Unix时间戳，秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最后修改人
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifier *string `json:"LastModifier,omitnil,omitempty" name:"LastModifier"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 模板语言 en/zh
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`
}

type NoticeContentTmplBindPolicyCount struct {
	// 通知内容模板ID
	NoticeContentTmplID *string `json:"NoticeContentTmplID,omitnil,omitempty" name:"NoticeContentTmplID"`

	// 绑定告警策略数量
	BindCount *uint64 `json:"BindCount,omitnil,omitempty" name:"BindCount"`
}

type NoticeContentTmplItem struct {
	// 官网通知渠道配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	QCloudYehe []*QCloudYeheNoticeTmplMatcher `json:"QCloudYehe,omitnil,omitempty" name:"QCloudYehe"`

	// 企业微信机器人通知渠道配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeWorkRobot []*WeWorkRobotNoticeTmplMatcher `json:"WeWorkRobot,omitnil,omitempty" name:"WeWorkRobot"`

	// 钉钉机器人通知渠道配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DingDingRobot []*DingDingRobotNoticeTmplMatcher `json:"DingDingRobot,omitnil,omitempty" name:"DingDingRobot"`

	// 飞书机器人通知渠道配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeiShuRobot []*FeiShuRobotNoticeTmplMatcher `json:"FeiShuRobot,omitnil,omitempty" name:"FeiShuRobot"`

	// 自定义Webhook通知渠道配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Webhook []*WebhookNoticeTmplMatcher `json:"Webhook,omitnil,omitempty" name:"Webhook"`

	// Teams机器人通知渠道配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeamsRobot []*TeamsRobotNoticeTmplMatcher `json:"TeamsRobot,omitnil,omitempty" name:"TeamsRobot"`

	// PagerDutyRobot机器人通知渠道配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PagerDutyRobot []*PagerDutyRobotNoticeTmplMatcher `json:"PagerDutyRobot,omitnil,omitempty" name:"PagerDutyRobot"`

	// GoogleChat
	GoogleChatRobot []*GoogleChatRobotNoticeTmplMatcher `json:"GoogleChatRobot,omitnil,omitempty" name:"GoogleChatRobot"`
}

type PageByNoParams struct {
	// 每个分页的数量是多少
	// 注意：此字段可能返回 null，表示取不到有效值。
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// 第几个分页，从1开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

type PagerDutyRobotNoticeTmpl struct {
	// 请求体模板 仅支持json
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 请求头 暂时未支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*PagerDutyRobotNoticeTmplHeader `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 标题模板
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type PagerDutyRobotNoticeTmplHeader struct {
	// http请求中header的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// http请求中header的value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type PagerDutyRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid; Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 自定义PagerDutyRobot内容模板
	Template *PagerDutyRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type QCloudYeheNoticeTmpl struct {
	// 邮件通知渠道
	Email *QCloudYeheNoticeTmplItem `json:"Email,omitnil,omitempty" name:"Email"`

	// 企业微信通知渠道
	QYWX *QCloudYeheNoticeTmplItem `json:"QYWX,omitnil,omitempty" name:"QYWX"`

	// 短信通知渠道
	SMS *QCloudYeheNoticeTmplItem `json:"SMS,omitnil,omitempty" name:"SMS"`

	// 语音通知渠道
	Voice *QCloudYeheNoticeTmplItem `json:"Voice,omitnil,omitempty" name:"Voice"`

	// 微信通知渠道
	WeChat *QCloudYeheWeChatNoticeTmplItem `json:"WeChat,omitnil,omitempty" name:"WeChat"`

	// 站内信通知渠道
	Site *QCloudYeheNoticeTmplItem `json:"Site,omitnil,omitempty" name:"Site"`

	// 安灯通知渠道
	Andon *QCloudYeheNoticeTmplItem `json:"Andon,omitnil,omitempty" name:"Andon"`
}

type QCloudYeheNoticeTmplItem struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`

	// 标题
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type QCloudYeheNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *QCloudYeheNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type QCloudYeheWeChatNoticeTmplItem struct {
	// 告警内容模板
	AlarmContentTmpl *string `json:"AlarmContentTmpl,omitnil,omitempty" name:"AlarmContentTmpl"`

	// 告警对象模板
	AlarmObjectTmpl *string `json:"AlarmObjectTmpl,omitnil,omitempty" name:"AlarmObjectTmpl"`

	// 告警地域模板
	AlarmRegionTmpl *string `json:"AlarmRegionTmpl,omitnil,omitempty" name:"AlarmRegionTmpl"`

	// 告警时间模板
	AlarmTimeTmpl *string `json:"AlarmTimeTmpl,omitnil,omitempty" name:"AlarmTimeTmpl"`
}

type TeamsRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`
}

type TeamsRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *TeamsRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

// Predefined struct for user
type TriggerAIWorkbenchSREDigitalTwinTaskRequestParams struct {
	// 数字分身任务ID
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type TriggerAIWorkbenchSREDigitalTwinTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数字分身任务ID
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *TriggerAIWorkbenchSREDigitalTwinTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchSREDigitalTwinTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerAIWorkbenchSREDigitalTwinTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerAIWorkbenchSREDigitalTwinTaskResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身任务信息
	Data *TriggerDigitalTwinTaskResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerAIWorkbenchSREDigitalTwinTaskResponse struct {
	*tchttp.BaseResponse
	Response *TriggerAIWorkbenchSREDigitalTwinTaskResponseParams `json:"Response"`
}

func (r *TriggerAIWorkbenchSREDigitalTwinTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchSREDigitalTwinTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TriggerDigitalTwinTaskResp struct {
	// 数字分身任务ID
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type WeWorkRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`
}

type WeWorkRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *WeWorkRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type WebhookNoticeTmpl struct {
	// 请求体
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 请求体的类型，非必填、默认为JSON
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyContentType *string `json:"BodyContentType,omitnil,omitempty" name:"BodyContentType"`

	// 请求头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*WebhookNoticeTmplHeader `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type WebhookNoticeTmplHeader struct {
	// http请求中header的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// http请求中header的value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type WebhookNoticeTmplMatcher struct {
	// 匹配状态 Invalid; Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 自定义Webhook内容模板
	Template *WebhookNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}