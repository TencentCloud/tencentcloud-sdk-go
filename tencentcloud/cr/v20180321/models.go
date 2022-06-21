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

package v20180321

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ApplyBlackListDataRequestParams struct {
	// 模块名，AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，ApplyBlackListData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 黑名单列表
	BlackList []*BlackListData `json:"BlackList,omitempty" name:"BlackList"`
}

type ApplyBlackListDataRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，ApplyBlackListData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 黑名单列表
	BlackList []*BlackListData `json:"BlackList,omitempty" name:"BlackList"`
}

func (r *ApplyBlackListDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyBlackListDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BlackList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyBlackListDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyBlackListDataResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyBlackListDataResponse struct {
	*tchttp.BaseResponse
	Response *ApplyBlackListDataResponseParams `json:"Response"`
}

func (r *ApplyBlackListDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyBlackListDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyBlackListRequestParams struct {
	// 模块名，本接口取值：account
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：ApplyBlackList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 黑名单列表
	BlackList []*SingleBlackApply `json:"BlackList,omitempty" name:"BlackList"`

	// 实例ID，不传默认为系统分配的初始实例
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

type ApplyBlackListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：account
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：ApplyBlackList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 黑名单列表
	BlackList []*SingleBlackApply `json:"BlackList,omitempty" name:"BlackList"`

	// 实例ID，不传默认为系统分配的初始实例
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *ApplyBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyBlackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BlackList")
	delete(f, "InstId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyBlackListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyBlackListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyBlackListResponse struct {
	*tchttp.BaseResponse
	Response *ApplyBlackListResponseParams `json:"Response"`
}

func (r *ApplyBlackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCreditAuditRequestParams struct {
	// 模块名，本接口取值：Credit
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Apply
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`

	// 产品ID，形如P******。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 信审任务ID，同一天内，同一InstId下，同一CaseId只能调用一次。
	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`

	// 回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// JSON格式的业务字段。
	Data *string `json:"Data,omitempty" name:"Data"`
}

type ApplyCreditAuditRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Credit
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Apply
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`

	// 产品ID，形如P******。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 信审任务ID，同一天内，同一InstId下，同一CaseId只能调用一次。
	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`

	// 回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// JSON格式的业务字段。
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *ApplyCreditAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCreditAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "InstId")
	delete(f, "ProductId")
	delete(f, "CaseId")
	delete(f, "CallbackUrl")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyCreditAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCreditAuditResponseParams struct {
	// 请求日期
	RequestDate *string `json:"RequestDate,omitempty" name:"RequestDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyCreditAuditResponse struct {
	*tchttp.BaseResponse
	Response *ApplyCreditAuditResponseParams `json:"Response"`
}

func (r *ApplyCreditAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCreditAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlackListData struct {
	// 黑名单类型，01代表手机号码。
	BlackType *string `json:"BlackType,omitempty" name:"BlackType"`

	// 操作类型，A为新增，D为删除。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperType *string `json:"OperType,omitempty" name:"OperType"`

	// 黑名单值，BlackType为01时，填写11位手机号码。
	BlackValue *string `json:"BlackValue,omitempty" name:"BlackValue"`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlackDescription *string `json:"BlackDescription,omitempty" name:"BlackDescription"`

	// 黑名单生效截止日期，格式为YYYY-MM-DD，不填默认为永久。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlackValidDate *string `json:"BlackValidDate,omitempty" name:"BlackValidDate"`

	// 黑名单加入日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlackAddDate *string `json:"BlackAddDate,omitempty" name:"BlackAddDate"`

	// 0-生效 1-失效
	BlackStatus *string `json:"BlackStatus,omitempty" name:"BlackStatus"`
}

type BotFileData struct {
	// 文件类型 A 拨打结果 T 记录详情
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件地址
	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`
}

type BotFlow struct {
	// 对话流ID
	BotFlowId *string `json:"BotFlowId,omitempty" name:"BotFlowId"`

	// 对话流名称
	BotFlowName *string `json:"BotFlowName,omitempty" name:"BotFlowName"`

	// 号码组信息列表
	PhonePoolList []*PhonePool `json:"PhonePoolList,omitempty" name:"PhonePoolList"`
}

type BotInfo struct {
	// 机器人ID
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 机器人名称
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 机器人状态。0-停用 1-启用 2-待审核
	BotStatus *string `json:"BotStatus,omitempty" name:"BotStatus"`
}

type CallInfo struct {
	// 业务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 状态 WAIT：待执行；DOING：执行中；ERROR：执行错误；DONE：已完成；
	Status *string `json:"Status,omitempty" name:"Status"`

	// 成功总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型 I：呼叫文件 R：停拨文件
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 作业唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallId *string `json:"CallId,omitempty" name:"CallId"`
}

type CallTimeDict struct {
	// 周一
	Monday *CallTimeInfo `json:"Monday,omitempty" name:"Monday"`

	// 周二
	Tuesday *CallTimeInfo `json:"Tuesday,omitempty" name:"Tuesday"`

	// 周三
	Wednesday *CallTimeInfo `json:"Wednesday,omitempty" name:"Wednesday"`

	// 周四
	Thursday *CallTimeInfo `json:"Thursday,omitempty" name:"Thursday"`

	// 周五
	Friday *CallTimeInfo `json:"Friday,omitempty" name:"Friday"`

	// 周六
	Saturday *CallTimeInfo `json:"Saturday,omitempty" name:"Saturday"`

	// 周日
	Sunday *CallTimeInfo `json:"Sunday,omitempty" name:"Sunday"`
}

type CallTimeInfo struct {
	// 产品开始拨打时间，HHmmss格式,默认090000
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 产品结束拨打时间，HHmmss格式.默认200000
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

// Predefined struct for user
type ChangeBotCallStatusRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：ChangeBotCallStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 作业变更状态
	// SUSPEND：暂停；EXECUTE：恢复；
	Status *string `json:"Status,omitempty" name:"Status"`

	// 作业唯一标识
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 业务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

type ChangeBotCallStatusRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：ChangeBotCallStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 作业变更状态
	// SUSPEND：暂停；EXECUTE：恢复；
	Status *string `json:"Status,omitempty" name:"Status"`

	// 作业唯一标识
	CallId *string `json:"CallId,omitempty" name:"CallId"`

	// 业务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

func (r *ChangeBotCallStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeBotCallStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Status")
	delete(f, "CallId")
	delete(f, "BizDate")
	delete(f, "BotId")
	delete(f, "BotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeBotCallStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeBotCallStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeBotCallStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChangeBotCallStatusResponseParams `json:"Response"`
}

func (r *ChangeBotCallStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeBotCallStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeBotTaskStatusRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：ChangeBotTaskStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 作业变更状态
	// SUSPEND：暂停；EXECUTE：恢复；
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

type ChangeBotTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：ChangeBotTaskStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 作业变更状态
	// SUSPEND：暂停；EXECUTE：恢复；
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

func (r *ChangeBotTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeBotTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Status")
	delete(f, "BotId")
	delete(f, "BotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeBotTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeBotTaskStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeBotTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *ChangeBotTaskStatusResponseParams `json:"Response"`
}

func (r *ChangeBotTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeBotTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBotTaskRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：CreateTask
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务名称
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 对话流ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 是否禁止拨打，默认Y
	BanCall *string `json:"BanCall,omitempty" name:"BanCall"`

	// 拨打线路集合
	PhoneCollection *string `json:"PhoneCollection,omitempty" name:"PhoneCollection"`

	// 产品拨打时间集合
	CallTimeCollection *CallTimeDict `json:"CallTimeCollection,omitempty" name:"CallTimeCollection"`

	// 禁止拨打起始时间。默认130000
	StartTimeBan *string `json:"StartTimeBan,omitempty" name:"StartTimeBan"`

	// 禁止拨打结束时间。默认140000
	EndTimeBan *string `json:"EndTimeBan,omitempty" name:"EndTimeBan"`

	// 重播方式，NON：未接通、LABEL：意向分级，可多选，用竖线分隔：NON|LABEL
	CodeType *string `json:"CodeType,omitempty" name:"CodeType"`

	// 重播值集合，A：强意向、B：中意向、C：低意向、D：无意向、E：在忙、F：未接通、G：无效号码，可多选，用竖线分隔：A|B|C|D|E|F|G
	CodeCollection *string `json:"CodeCollection,omitempty" name:"CodeCollection"`

	// 继续拨打次数
	CallCount *int64 `json:"CallCount,omitempty" name:"CallCount"`

	// 拨打间隔
	CallInterval *int64 `json:"CallInterval,omitempty" name:"CallInterval"`

	// 未接通引用短信签名ID
	SmsSignId *string `json:"SmsSignId,omitempty" name:"SmsSignId"`

	// 未接通引用短信模板ID
	SmsTemplateId *string `json:"SmsTemplateId,omitempty" name:"SmsTemplateId"`

	// 拨打方式。NORMAL - 正常拨打；TIMER - 定时拨打
	CallType *string `json:"CallType,omitempty" name:"CallType"`

	// 拨打开始日期。CallType=TIMER时有值，yyyy-MM-dd
	CallStartDate *string `json:"CallStartDate,omitempty" name:"CallStartDate"`

	// 拨打结束日期。CallType=PERIOD 时有值，yyyy-MM-dd
	CallEndDate *string `json:"CallEndDate,omitempty" name:"CallEndDate"`
}

type CreateBotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：CreateTask
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务名称
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 对话流ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 是否禁止拨打，默认Y
	BanCall *string `json:"BanCall,omitempty" name:"BanCall"`

	// 拨打线路集合
	PhoneCollection *string `json:"PhoneCollection,omitempty" name:"PhoneCollection"`

	// 产品拨打时间集合
	CallTimeCollection *CallTimeDict `json:"CallTimeCollection,omitempty" name:"CallTimeCollection"`

	// 禁止拨打起始时间。默认130000
	StartTimeBan *string `json:"StartTimeBan,omitempty" name:"StartTimeBan"`

	// 禁止拨打结束时间。默认140000
	EndTimeBan *string `json:"EndTimeBan,omitempty" name:"EndTimeBan"`

	// 重播方式，NON：未接通、LABEL：意向分级，可多选，用竖线分隔：NON|LABEL
	CodeType *string `json:"CodeType,omitempty" name:"CodeType"`

	// 重播值集合，A：强意向、B：中意向、C：低意向、D：无意向、E：在忙、F：未接通、G：无效号码，可多选，用竖线分隔：A|B|C|D|E|F|G
	CodeCollection *string `json:"CodeCollection,omitempty" name:"CodeCollection"`

	// 继续拨打次数
	CallCount *int64 `json:"CallCount,omitempty" name:"CallCount"`

	// 拨打间隔
	CallInterval *int64 `json:"CallInterval,omitempty" name:"CallInterval"`

	// 未接通引用短信签名ID
	SmsSignId *string `json:"SmsSignId,omitempty" name:"SmsSignId"`

	// 未接通引用短信模板ID
	SmsTemplateId *string `json:"SmsTemplateId,omitempty" name:"SmsTemplateId"`

	// 拨打方式。NORMAL - 正常拨打；TIMER - 定时拨打
	CallType *string `json:"CallType,omitempty" name:"CallType"`

	// 拨打开始日期。CallType=TIMER时有值，yyyy-MM-dd
	CallStartDate *string `json:"CallStartDate,omitempty" name:"CallStartDate"`

	// 拨打结束日期。CallType=PERIOD 时有值，yyyy-MM-dd
	CallEndDate *string `json:"CallEndDate,omitempty" name:"CallEndDate"`
}

func (r *CreateBotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BotName")
	delete(f, "FlowId")
	delete(f, "BanCall")
	delete(f, "PhoneCollection")
	delete(f, "CallTimeCollection")
	delete(f, "StartTimeBan")
	delete(f, "EndTimeBan")
	delete(f, "CodeType")
	delete(f, "CodeCollection")
	delete(f, "CallCount")
	delete(f, "CallInterval")
	delete(f, "SmsSignId")
	delete(f, "SmsTemplateId")
	delete(f, "CallType")
	delete(f, "CallStartDate")
	delete(f, "CallEndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBotTaskResponseParams struct {
	// 机器人任务Id
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBotTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateBotTaskResponseParams `json:"Response"`
}

func (r *CreateBotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotFlowRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：GetFlow
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type DescribeBotFlowRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：GetFlow
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DescribeBotFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBotFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBotFlowResponseParams struct {
	// 机器人对话流列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotFlowList []*BotFlow `json:"BotFlowList,omitempty" name:"BotFlowList"`

	// 短信签名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsSignList []*SmsSign `json:"SmsSignList,omitempty" name:"SmsSignList"`

	// 短信模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsTemplateList []*SmsTemplate `json:"SmsTemplateList,omitempty" name:"SmsTemplateList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBotFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBotFlowResponseParams `json:"Response"`
}

func (r *DescribeBotFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBotFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCreditResultRequestParams struct {
	// 模块名，本接口取值：Credit
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Get
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`

	// 产品ID，形如P******。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 信审任务ID
	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`

	// 请求日期，格式为YYYY-MM-DD
	RequestDate *string `json:"RequestDate,omitempty" name:"RequestDate"`
}

type DescribeCreditResultRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Credit
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Get
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`

	// 产品ID，形如P******。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 信审任务ID
	CaseId *string `json:"CaseId,omitempty" name:"CaseId"`

	// 请求日期，格式为YYYY-MM-DD
	RequestDate *string `json:"RequestDate,omitempty" name:"RequestDate"`
}

func (r *DescribeCreditResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreditResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "InstId")
	delete(f, "ProductId")
	delete(f, "CaseId")
	delete(f, "RequestDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCreditResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCreditResultResponseParams struct {
	// <p>呼叫结果，取值范围：</p><ul style="margin-bottom:0px;"><li>NON：接通</li><li>DBU：号码忙</li><li>DRF：不在服务区</li><li>ANA：欠费未接听</li><li>REJ：拒接</li><li>SHU：关机</li><li>NAN：空号</li><li>HAL：停机</li><li>DAD：未接听</li><li>EXE：其他异常</li></ul>
	ResultCode *string `json:"ResultCode,omitempty" name:"ResultCode"`

	// 客户标识代码，多个标识码以英文逗号分隔，ResultCode为NON时才有。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCode *string `json:"ClientCode,omitempty" name:"ClientCode"`

	// 开始振铃时间，ResultCode为NON或DAD时才有此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RingStartTime *string `json:"RingStartTime,omitempty" name:"RingStartTime"`

	// 振铃时长
	RingDuration *int64 `json:"RingDuration,omitempty" name:"RingDuration"`

	// 接通时长
	AnswerDuration *int64 `json:"AnswerDuration,omitempty" name:"AnswerDuration"`

	// JSON格式的扩展信息字段，ResultCode为NON时才有。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextValue *string `json:"ContextValue,omitempty" name:"ContextValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCreditResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCreditResultResponseParams `json:"Response"`
}

func (r *DescribeCreditResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreditResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileModelRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：DescribeFileModel
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 模板文件类型，输入input，停拨stop
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

type DescribeFileModelRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：DescribeFileModel
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 模板文件类型，输入input，停拨stop
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

func (r *DescribeFileModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "FileType")
	delete(f, "BotId")
	delete(f, "BotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileModelResponseParams struct {
	// 模板下载链接
	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileModelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileModelResponseParams `json:"Response"`
}

func (r *DescribeFileModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordsRequestParams struct {
	// 模块名，本接口取值：Record
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：List
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 案件编号
	AccountNum *string `json:"AccountNum,omitempty" name:"AccountNum"`

	// 被叫号码
	CalledPhone *string `json:"CalledPhone,omitempty" name:"CalledPhone"`

	// 查询起始日期，格式为YYYY-MM-DD
	StartBizDate *string `json:"StartBizDate,omitempty" name:"StartBizDate"`

	// 查询结束日期，格式为YYYY-MM-DD
	EndBizDate *string `json:"EndBizDate,omitempty" name:"EndBizDate"`

	// 分页参数，索引，默认为0
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，页长，默认为20
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 实例ID，不传默认为系统分配的初始实例
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

type DescribeRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Record
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：List
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 案件编号
	AccountNum *string `json:"AccountNum,omitempty" name:"AccountNum"`

	// 被叫号码
	CalledPhone *string `json:"CalledPhone,omitempty" name:"CalledPhone"`

	// 查询起始日期，格式为YYYY-MM-DD
	StartBizDate *string `json:"StartBizDate,omitempty" name:"StartBizDate"`

	// 查询结束日期，格式为YYYY-MM-DD
	EndBizDate *string `json:"EndBizDate,omitempty" name:"EndBizDate"`

	// 分页参数，索引，默认为0
	Offset *string `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，页长，默认为20
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 实例ID，不传默认为系统分配的初始实例
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *DescribeRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ProductId")
	delete(f, "AccountNum")
	delete(f, "CalledPhone")
	delete(f, "StartBizDate")
	delete(f, "EndBizDate")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordsResponseParams struct {
	// 录音列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordList []*SingleRecord `json:"RecordList,omitempty" name:"RecordList"`

	// 录音总量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordsResponseParams `json:"Response"`
}

func (r *DescribeRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 模块名，本接口取值：Task
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DescribeTaskStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务ID，"上传文件"接口返回的DataResId，形如abc-xyz123
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Task
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DescribeTaskStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务ID，"上传文件"接口返回的DataResId，形如abc-xyz123
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "TaskId")
	delete(f, "InstId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// <p>任务结果：</p><ul style="margin-bottom:0px;"><li>处理中："Uploading Data."</li><li>上传成功："File Uploading Task Success."</li><li>上传失败：具体失败原因</li></ul>
	TaskResult *string `json:"TaskResult,omitempty" name:"TaskResult"`

	// <p>任务类型：</p><ul style="margin-bottom:0px;"><li>到期/逾期提醒数据上传：002</li><li>到期/逾期提醒停拨数据上传：003</li><li>回访数据上传：004</li><li>回访停拨数据上传：005</li></ul>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 过滤文件下载链接，有过滤数据时才存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFileUrl *string `json:"TaskFileUrl,omitempty" name:"TaskFileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadBotRecordRequestParams struct {
	// 模块：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作：DownloadRecord
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 业务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`
}

type DownloadBotRecordRequest struct {
	*tchttp.BaseRequest
	
	// 模块：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作：DownloadRecord
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 业务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`
}

func (r *DownloadBotRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadBotRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BizDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadBotRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadBotRecordResponseParams struct {
	// 录音地址。请求后30分钟内有效
	RecordCosUrl *string `json:"RecordCosUrl,omitempty" name:"RecordCosUrl"`

	// 文本地址。请求后30分钟内有效
	TextCosUrl *string `json:"TextCosUrl,omitempty" name:"TextCosUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadBotRecordResponse struct {
	*tchttp.BaseResponse
	Response *DownloadBotRecordResponseParams `json:"Response"`
}

func (r *DownloadBotRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadBotRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadDialogueTextRequestParams struct {
	// 模块名，本接口取值：Report
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DownloadTextReport
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 报告日期，格式为YYYY-MM-DD
	ReportDate *string `json:"ReportDate,omitempty" name:"ReportDate"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

type DownloadDialogueTextRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Report
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DownloadTextReport
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 报告日期，格式为YYYY-MM-DD
	ReportDate *string `json:"ReportDate,omitempty" name:"ReportDate"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *DownloadDialogueTextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadDialogueTextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ReportDate")
	delete(f, "InstId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadDialogueTextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadDialogueTextResponseParams struct {
	// 对话文本下载地址
	TextReportUrl *string `json:"TextReportUrl,omitempty" name:"TextReportUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadDialogueTextResponse struct {
	*tchttp.BaseResponse
	Response *DownloadDialogueTextResponseParams `json:"Response"`
}

func (r *DownloadDialogueTextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadDialogueTextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadRecordListRequestParams struct {
	// 模块名，本接口取值：Record
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DownloadList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 录音日期，格式为YYYY-MM-DD
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

type DownloadRecordListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Record
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DownloadList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 录音日期，格式为YYYY-MM-DD
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 实例ID
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *DownloadRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BizDate")
	delete(f, "InstId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadRecordListResponseParams struct {
	// 录音列表下载地址
	RecordListUrl *string `json:"RecordListUrl,omitempty" name:"RecordListUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadRecordListResponse struct {
	*tchttp.BaseResponse
	Response *DownloadRecordListResponseParams `json:"Response"`
}

func (r *DownloadRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadReportRequestParams struct {
	// 模块名，本接口取值：Report
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DownloadReport
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 报告日期，格式为YYYY-MM-DD
	ReportDate *string `json:"ReportDate,omitempty" name:"ReportDate"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

type DownloadReportRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Report
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：DownloadReport
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 报告日期，格式为YYYY-MM-DD
	ReportDate *string `json:"ReportDate,omitempty" name:"ReportDate"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *DownloadReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ReportDate")
	delete(f, "InstId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadReportResponseParams struct {
	// 到期/逾期提醒日报下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DailyReportUrl *string `json:"DailyReportUrl,omitempty" name:"DailyReportUrl"`

	// 到期/逾期提醒结果下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultReportUrl *string `json:"ResultReportUrl,omitempty" name:"ResultReportUrl"`

	// 到期/逾期提醒明细下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailReportUrl *string `json:"DetailReportUrl,omitempty" name:"DetailReportUrl"`

	// 回访日报下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackDailyReportUrl *string `json:"CallbackDailyReportUrl,omitempty" name:"CallbackDailyReportUrl"`

	// 回访结果下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackResultReportUrl *string `json:"CallbackResultReportUrl,omitempty" name:"CallbackResultReportUrl"`

	// 回访明细下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackDetailReportUrl *string `json:"CallbackDetailReportUrl,omitempty" name:"CallbackDetailReportUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadReportResponse struct {
	*tchttp.BaseResponse
	Response *DownloadReportResponseParams `json:"Response"`
}

func (r *DownloadReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBotDataRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：ExportBotData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 业务日期。YYYY-MM-DD
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

type ExportBotDataRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：ExportBotData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 业务日期。YYYY-MM-DD
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

func (r *ExportBotDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBotDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BizDate")
	delete(f, "BotId")
	delete(f, "BotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportBotDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportBotDataResponseParams struct {
	// 导出文件列表
	Data []*BotFileData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExportBotDataResponse struct {
	*tchttp.BaseResponse
	Response *ExportBotDataResponseParams `json:"Response"`
}

func (r *ExportBotDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportBotDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhonePool struct {
	// 号码组ID
	PoolId *string `json:"PoolId,omitempty" name:"PoolId"`

	// 号码组名称
	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`
}

type ProductQueryInfo struct {
	// 产品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 产品状态 0 禁用 1 启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductStatus *int64 `json:"ProductStatus,omitempty" name:"ProductStatus"`

	// 场景类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneType *string `json:"SceneType,omitempty" name:"SceneType"`
}

// Predefined struct for user
type QueryBlackListDataRequestParams struct {
	// 模块:AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作:QueryBlackListData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 开始日期
	StartBizDate *string `json:"StartBizDate,omitempty" name:"StartBizDate"`

	// 结束日期
	EndBizDate *string `json:"EndBizDate,omitempty" name:"EndBizDate"`

	// 电话号码、手机
	BlackValue *string `json:"BlackValue,omitempty" name:"BlackValue"`
}

type QueryBlackListDataRequest struct {
	*tchttp.BaseRequest
	
	// 模块:AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作:QueryBlackListData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 开始日期
	StartBizDate *string `json:"StartBizDate,omitempty" name:"StartBizDate"`

	// 结束日期
	EndBizDate *string `json:"EndBizDate,omitempty" name:"EndBizDate"`

	// 电话号码、手机
	BlackValue *string `json:"BlackValue,omitempty" name:"BlackValue"`
}

func (r *QueryBlackListDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBlackListDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartBizDate")
	delete(f, "EndBizDate")
	delete(f, "BlackValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBlackListDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryBlackListDataResponseParams struct {
	// 总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 黑名单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*BlackListData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryBlackListDataResponse struct {
	*tchttp.BaseResponse
	Response *QueryBlackListDataResponseParams `json:"Response"`
}

func (r *QueryBlackListDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBlackListDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryBotListRequestParams struct {
	// 模块名：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名：QueryBotList
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type QueryBotListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名：QueryBotList
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *QueryBotListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBotListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBotListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryBotListResponseParams struct {
	// 任务列表。
	BotList []*BotInfo `json:"BotList,omitempty" name:"BotList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryBotListResponse struct {
	*tchttp.BaseResponse
	Response *QueryBotListResponseParams `json:"Response"`
}

func (r *QueryBotListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBotListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCallListRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：QueryCallList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 业务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 通过API或平台上传的文件完整名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

type QueryCallListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：QueryCallList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 业务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 通过API或平台上传的文件完整名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

func (r *QueryCallListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BizDate")
	delete(f, "BotId")
	delete(f, "BotName")
	delete(f, "FileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCallListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCallListResponseParams struct {
	// 任务作业状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallList []*CallInfo `json:"CallList,omitempty" name:"CallList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCallListResponse struct {
	*tchttp.BaseResponse
	Response *QueryCallListResponseParams `json:"Response"`
}

func (r *QueryCallListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCallListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryInstantDataRequestParams struct {
	// 模块名，本接口取值：Data
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Query
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 实例ID，不传默认为系统分配的初始实例
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询类型：callRecord 通话记录
	QueryModel *string `json:"QueryModel,omitempty" name:"QueryModel"`

	// 查询参数
	Data *string `json:"Data,omitempty" name:"Data"`
}

type QueryInstantDataRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Data
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Query
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 实例ID，不传默认为系统分配的初始实例
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询类型：callRecord 通话记录
	QueryModel *string `json:"QueryModel,omitempty" name:"QueryModel"`

	// 查询参数
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *QueryInstantDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInstantDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ProductId")
	delete(f, "InstanceId")
	delete(f, "QueryModel")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryInstantDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryInstantDataResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryInstantDataResponse struct {
	*tchttp.BaseResponse
	Response *QueryInstantDataResponseParams `json:"Response"`
}

func (r *QueryInstantDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInstantDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryProductsRequestParams struct {
	// 模块名。默认值（固定）：Product
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：QueryProducts
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例Id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type QueryProductsRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：Product
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：QueryProducts
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例Id。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *QueryProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryProductsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryProductsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryProductsResponseParams struct {
	// 产品信息。
	ProductList []*ProductQueryInfo `json:"ProductList,omitempty" name:"ProductList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryProductsResponse struct {
	*tchttp.BaseResponse
	Response *QueryProductsResponseParams `json:"Response"`
}

func (r *QueryProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRecordListRequestParams struct {
	// 模块名。AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。QueryRecordList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 偏移位移，最大20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 被叫号码
	CalledPhone *string `json:"CalledPhone,omitempty" name:"CalledPhone"`

	// 开始日期
	StartBizDate *string `json:"StartBizDate,omitempty" name:"StartBizDate"`

	// 结束日期
	EndBizDate *string `json:"EndBizDate,omitempty" name:"EndBizDate"`
}

type QueryRecordListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。QueryRecordList
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 偏移位移，最大20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 被叫号码
	CalledPhone *string `json:"CalledPhone,omitempty" name:"CalledPhone"`

	// 开始日期
	StartBizDate *string `json:"StartBizDate,omitempty" name:"StartBizDate"`

	// 结束日期
	EndBizDate *string `json:"EndBizDate,omitempty" name:"EndBizDate"`
}

func (r *QueryRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "BotId")
	delete(f, "BotName")
	delete(f, "CalledPhone")
	delete(f, "StartBizDate")
	delete(f, "EndBizDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRecordListResponseParams struct {
	// 录音列表。
	RecordList []*RecordInfo `json:"RecordList,omitempty" name:"RecordList"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryRecordListResponse struct {
	*tchttp.BaseResponse
	Response *QueryRecordListResponseParams `json:"Response"`
}

func (r *QueryRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordInfo struct {
	// 任务Id
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 任务日期
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 被叫号码
	CalledPhone *string `json:"CalledPhone,omitempty" name:"CalledPhone"`

	// 开始通话时间
	CallStartTime *string `json:"CallStartTime,omitempty" name:"CallStartTime"`

	// 通话时长
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 录音文件地址
	CosUrl *string `json:"CosUrl,omitempty" name:"CosUrl"`

	// 对话日志。JSON格式
	DialogueLog *string `json:"DialogueLog,omitempty" name:"DialogueLog"`

	// 录音文件名
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`
}

type SingleBlackApply struct {
	// 黑名单类型，01代表手机号码。
	BlackType *string `json:"BlackType,omitempty" name:"BlackType"`

	// 操作类型，A为新增，D为删除。
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 黑名单值，BlackType为01时，填写11位手机号码。
	BlackValue *string `json:"BlackValue,omitempty" name:"BlackValue"`

	// 备注。
	BlackDescription *string `json:"BlackDescription,omitempty" name:"BlackDescription"`

	// 黑名单生效截止日期，格式为YYYY-MM-DD，不填默认为永久。
	BlackValidDate *string `json:"BlackValidDate,omitempty" name:"BlackValidDate"`
}

type SingleRecord struct {
	// 案件编号。
	AccountNum *string `json:"AccountNum,omitempty" name:"AccountNum"`

	// 外呼日期。
	BizDate *string `json:"BizDate,omitempty" name:"BizDate"`

	// 开始呼叫时间。
	CallStartTime *string `json:"CallStartTime,omitempty" name:"CallStartTime"`

	// 主叫号码。
	CallerPhone *string `json:"CallerPhone,omitempty" name:"CallerPhone"`

	// 呼叫方向，O为呼出，I为呼入。
	Direction *string `json:"Direction,omitempty" name:"Direction"`

	// 通话时长。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 产品ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 录音下载链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCosUrl *string `json:"RecordCosUrl,omitempty" name:"RecordCosUrl"`
}

type SmsSign struct {
	// 短信签名ID
	SignId *string `json:"SignId,omitempty" name:"SignId"`

	// 短信签名名称
	SignName *string `json:"SignName,omitempty" name:"SignName"`
}

type SmsTemplate struct {
	// 短信模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 短信模板名称
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
}

// Predefined struct for user
type UpdateBotTaskRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：UpdateTask
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务名称
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 任务ID
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 产品拨打时间集合
	CallTimeCollection *CallTimeDict `json:"CallTimeCollection,omitempty" name:"CallTimeCollection"`

	// 是否禁止拨打，默认Y
	BanCall *string `json:"BanCall,omitempty" name:"BanCall"`

	// 禁止拨打起始时间。默认130000
	StartTimeBan *string `json:"StartTimeBan,omitempty" name:"StartTimeBan"`

	// 禁止拨打结束时间。默认140000
	EndTimeBan *string `json:"EndTimeBan,omitempty" name:"EndTimeBan"`

	// 拨打线路集合
	PhoneCollection *string `json:"PhoneCollection,omitempty" name:"PhoneCollection"`

	// 重播方式，NON：未接通、LABEL：意向分级，可多选，用竖线分隔：NON|LABEL
	CodeType *string `json:"CodeType,omitempty" name:"CodeType"`

	// 重播值集合，A：强意向、B：中意向、C：低意向、D：无意向、E：在忙、F：未接通、G：无效号码，可多选，用竖线分隔：A|B|C|D|E|F|G
	CodeCollection *string `json:"CodeCollection,omitempty" name:"CodeCollection"`

	// 继续拨打次数
	CallCount *int64 `json:"CallCount,omitempty" name:"CallCount"`

	// 拨打间隔
	CallInterval *int64 `json:"CallInterval,omitempty" name:"CallInterval"`

	// 未接通引用短信签名ID
	SmsSignId *string `json:"SmsSignId,omitempty" name:"SmsSignId"`

	// 未接通引用短信模板ID
	SmsTemplateId *string `json:"SmsTemplateId,omitempty" name:"SmsTemplateId"`
}

type UpdateBotTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：UpdateTask
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务名称
	BotName *string `json:"BotName,omitempty" name:"BotName"`

	// 任务ID
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 产品拨打时间集合
	CallTimeCollection *CallTimeDict `json:"CallTimeCollection,omitempty" name:"CallTimeCollection"`

	// 是否禁止拨打，默认Y
	BanCall *string `json:"BanCall,omitempty" name:"BanCall"`

	// 禁止拨打起始时间。默认130000
	StartTimeBan *string `json:"StartTimeBan,omitempty" name:"StartTimeBan"`

	// 禁止拨打结束时间。默认140000
	EndTimeBan *string `json:"EndTimeBan,omitempty" name:"EndTimeBan"`

	// 拨打线路集合
	PhoneCollection *string `json:"PhoneCollection,omitempty" name:"PhoneCollection"`

	// 重播方式，NON：未接通、LABEL：意向分级，可多选，用竖线分隔：NON|LABEL
	CodeType *string `json:"CodeType,omitempty" name:"CodeType"`

	// 重播值集合，A：强意向、B：中意向、C：低意向、D：无意向、E：在忙、F：未接通、G：无效号码，可多选，用竖线分隔：A|B|C|D|E|F|G
	CodeCollection *string `json:"CodeCollection,omitempty" name:"CodeCollection"`

	// 继续拨打次数
	CallCount *int64 `json:"CallCount,omitempty" name:"CallCount"`

	// 拨打间隔
	CallInterval *int64 `json:"CallInterval,omitempty" name:"CallInterval"`

	// 未接通引用短信签名ID
	SmsSignId *string `json:"SmsSignId,omitempty" name:"SmsSignId"`

	// 未接通引用短信模板ID
	SmsTemplateId *string `json:"SmsTemplateId,omitempty" name:"SmsTemplateId"`
}

func (r *UpdateBotTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBotTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "BotName")
	delete(f, "BotId")
	delete(f, "CallTimeCollection")
	delete(f, "BanCall")
	delete(f, "StartTimeBan")
	delete(f, "EndTimeBan")
	delete(f, "PhoneCollection")
	delete(f, "CodeType")
	delete(f, "CodeCollection")
	delete(f, "CallCount")
	delete(f, "CallInterval")
	delete(f, "SmsSignId")
	delete(f, "SmsTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateBotTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateBotTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateBotTaskResponse struct {
	*tchttp.BaseResponse
	Response *UpdateBotTaskResponseParams `json:"Response"`
}

func (r *UpdateBotTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateBotTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadBotDataRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：UploadData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务数据。JSON格式
	Data *string `json:"Data,omitempty" name:"Data"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

type UploadBotDataRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：UploadData
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务数据。JSON格式
	Data *string `json:"Data,omitempty" name:"Data"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

func (r *UploadBotDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadBotDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Data")
	delete(f, "BotId")
	delete(f, "BotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadBotDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadBotDataResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadBotDataResponse struct {
	*tchttp.BaseResponse
	Response *UploadBotDataResponseParams `json:"Response"`
}

func (r *UploadBotDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadBotDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadBotFileRequestParams struct {
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：Upload
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件类型，输入input，停拨stop
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件链接
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

type UploadBotFileRequest struct {
	*tchttp.BaseRequest
	
	// 模块名。默认值（固定）：AiApi
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名。默认值（固定）：Upload
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件类型，输入input，停拨stop
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件链接
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 任务ID，二者必填一个
	BotId *string `json:"BotId,omitempty" name:"BotId"`

	// 任务名称，二者必填一个
	BotName *string `json:"BotName,omitempty" name:"BotName"`
}

func (r *UploadBotFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadBotFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "FileType")
	delete(f, "FileUrl")
	delete(f, "FileName")
	delete(f, "BotId")
	delete(f, "BotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadBotFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadBotFileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadBotFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadBotFileResponseParams `json:"Response"`
}

func (r *UploadBotFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadBotFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDataFileRequestParams struct {
	// 模块名，本接口取值：Data
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Upload
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// <p>上传类型，不填默认到期/逾期提醒文件，取值范围：</p><ul style="margin-bottom:0px;"><li>data：到期/逾期提醒文件</li><li>repay：到期/逾期提醒停拨文件</li><li>callback：回访文件</li><li>callstop：回访停拨文件</li><li>blacklist：黑名单文件</li></ul>
	UploadModel *string `json:"UploadModel,omitempty" name:"UploadModel"`

	// 文件，文件与文件地址上传只可选用一种，必须使用multipart/form-data协议来上传二进制流文件，建议使用xlsx格式，大小不超过5MB。
	File *string `json:"File,omitempty" name:"File"`

	// 文件上传地址，文件与文件地址上传只可选用一种，大小不超过50MB。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

type UploadDataFileRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Data
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：Upload
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// <p>上传类型，不填默认到期/逾期提醒文件，取值范围：</p><ul style="margin-bottom:0px;"><li>data：到期/逾期提醒文件</li><li>repay：到期/逾期提醒停拨文件</li><li>callback：回访文件</li><li>callstop：回访停拨文件</li><li>blacklist：黑名单文件</li></ul>
	UploadModel *string `json:"UploadModel,omitempty" name:"UploadModel"`

	// 文件，文件与文件地址上传只可选用一种，必须使用multipart/form-data协议来上传二进制流文件，建议使用xlsx格式，大小不超过5MB。
	File *string `json:"File,omitempty" name:"File"`

	// 文件上传地址，文件与文件地址上传只可选用一种，大小不超过50MB。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstId *string `json:"InstId,omitempty" name:"InstId"`
}

func (r *UploadDataFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDataFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "FileName")
	delete(f, "UploadModel")
	delete(f, "File")
	delete(f, "FileUrl")
	delete(f, "InstId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadDataFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDataFileResponseParams struct {
	// 数据ID
	DataResId *string `json:"DataResId,omitempty" name:"DataResId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadDataFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadDataFileResponseParams `json:"Response"`
}

func (r *UploadDataFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDataFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDataJsonRequestParams struct {
	// 模块名，本接口取值：Data
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：UploadJson
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 报文信息
	Data *string `json:"Data,omitempty" name:"Data"`

	// <p>上传类型，不填默认到期/逾期提醒数据，取值范围：</p><ul style="margin-bottom:0px;"><li>data：到期/逾期提醒数据</li><li>repay：到期/逾期提醒停拨数据</li></ul>
	UploadModel *string `json:"UploadModel,omitempty" name:"UploadModel"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type UploadDataJsonRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，本接口取值：Data
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名，本接口取值：UploadJson
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 报文信息
	Data *string `json:"Data,omitempty" name:"Data"`

	// <p>上传类型，不填默认到期/逾期提醒数据，取值范围：</p><ul style="margin-bottom:0px;"><li>data：到期/逾期提醒数据</li><li>repay：到期/逾期提醒停拨数据</li></ul>
	UploadModel *string `json:"UploadModel,omitempty" name:"UploadModel"`

	// 实例ID，不传默认为系统分配的初始实例。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *UploadDataJsonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDataJsonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Data")
	delete(f, "UploadModel")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadDataJsonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadDataJsonResponseParams struct {
	// 响应报文信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadDataJsonResponse struct {
	*tchttp.BaseResponse
	Response *UploadDataJsonResponseParams `json:"Response"`
}

func (r *UploadDataJsonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadDataJsonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFileRequestParams struct {
	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件上传地址，要求地址协议为HTTPS，且URL端口必须为443
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件日期
	FileDate *string `json:"FileDate,omitempty" name:"FileDate"`
}

type UploadFileRequest struct {
	*tchttp.BaseRequest
	
	// 模块名
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 文件上传地址，要求地址协议为HTTPS，且URL端口必须为443
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件日期
	FileDate *string `json:"FileDate,omitempty" name:"FileDate"`
}

func (r *UploadFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "FileUrl")
	delete(f, "FileName")
	delete(f, "FileDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFileResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadFileResponseParams `json:"Response"`
}

func (r *UploadFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}