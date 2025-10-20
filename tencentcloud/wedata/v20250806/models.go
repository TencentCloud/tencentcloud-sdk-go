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

package v20250806

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddCalcEnginesToProjectRequestParams struct {
	// 修改的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// DLC集群信息
	DLCInfo []*DLCClusterInfo `json:"DLCInfo,omitnil,omitempty" name:"DLCInfo"`
}

type AddCalcEnginesToProjectRequest struct {
	*tchttp.BaseRequest
	
	// 修改的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// DLC集群信息
	DLCInfo []*DLCClusterInfo `json:"DLCInfo,omitnil,omitempty" name:"DLCInfo"`
}

func (r *AddCalcEnginesToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCalcEnginesToProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DLCInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCalcEnginesToProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCalcEnginesToProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCalcEnginesToProjectResponse struct {
	*tchttp.BaseResponse
	Response *AddCalcEnginesToProjectResponseParams `json:"Response"`
}

func (r *AddCalcEnginesToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCalcEnginesToProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmGroup struct {
	// 告警升级人ID列表
	// 若告警接收人或上级升级人未在告警间隔时间内确认告警，则会发送告警给下一级升级人。
	AlarmEscalationRecipientIds []*string `json:"AlarmEscalationRecipientIds,omitnil,omitempty" name:"AlarmEscalationRecipientIds"`

	// 告警升级间隔
	AlarmEscalationInterval *int64 `json:"AlarmEscalationInterval,omitnil,omitempty" name:"AlarmEscalationInterval"`

	// 告警通知疲劳配置
	NotificationFatigue *NotificationFatigue `json:"NotificationFatigue,omitnil,omitempty" name:"NotificationFatigue"`

	// 告警渠道 1.邮件，2.短信，3.微信，4.语音，5.企业微信，6.Http，7.企业微信群 8 飞书群 9 钉钉群 10 Slack群 11 Teams群（默认1.邮件） 7.企业微信群 8 飞书群 9 钉钉群 10 Slack群 11 Teams群 只能选择一个渠道
	AlarmWays []*string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// 企业微信群/飞书群/钉钉群 /Slack群/Teams群的webhook地址列表
	WebHooks []*AlarmWayWebHook `json:"WebHooks,omitnil,omitempty" name:"WebHooks"`

	// 告警接收人类型：1.指定人员，2.任务责任人，3.值班表（默认1.指定人员）
	AlarmRecipientType *int64 `json:"AlarmRecipientType,omitnil,omitempty" name:"AlarmRecipientType"`

	// 根据AlarmRecipientType的类型该列表具有不同的业务id 1（指定人员）: 告警接收人id列表 2（任务责任人）：无需配置 3（值班表）：值班表id列表
	AlarmRecipientIds []*string `json:"AlarmRecipientIds,omitnil,omitempty" name:"AlarmRecipientIds"`
}

type AlarmMessage struct {
	// 告警消息Id
	AlarmMessageId *uint64 `json:"AlarmMessageId,omitnil,omitempty" name:"AlarmMessageId"`

	// 告警时间，同一条告警可能发送多次，只显示最新的告警时间
	AlarmTime *string `json:"AlarmTime,omitnil,omitempty" name:"AlarmTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务的实例数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 告警原因
	AlarmReason *string `json:"AlarmReason,omitnil,omitempty" name:"AlarmReason"`

	// 告警级别，1.普通， 2.重要，3.紧急
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警规则Id
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// 告警渠道 1.邮件，2.短信，3.微信，4.语音，5.企业微信，6.Http，7.企业微信群， 8.飞书群，9.钉钉群，10.Slack群,11.Teams群（默认1.邮件），7.企业微信群，8.飞书群，9.钉钉群，10.Slack群，11.Teams群 
	AlarmWays []*string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// 告警接收人
	AlarmRecipients []*string `json:"AlarmRecipients,omitnil,omitempty" name:"AlarmRecipients"`
}

type AlarmQuietInterval struct {
	// ISO标准，1表示周一，7表示周日。
	DaysOfWeek []*uint64 `json:"DaysOfWeek,omitnil,omitempty" name:"DaysOfWeek"`

	// 开始时间，精度时分秒，格式 HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，精度时分秒，格式 HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type AlarmRuleData struct {
	// 告警规则id
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// 告警规则名称
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// 告警规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 监控对象类型,
	// 任务维度监控： 可按照任务/工作流/项目来配置：1.任务、2.工作流、3.项目（默认为1.任务） 
	// 项目维度监控： 项目整体任务波动告警， 7：项目波动监控告警
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 根据MonitorType 的设置传入不同的业务id，如下1（任务）： MonitorObjectIds为任务id列表2（工作流）： MonitorObjectIds 为工作流id列表(工作流id可从接口ListWorkflows获取)3（项目）： MonitorObjectIds为项目id列表
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// 告警规则监控类型
	// failure：失败告警 ；overtime：超时告警 success：成功告警; backTrackingOrRerunSuccess: 补录重跑成功告警 backTrackingOrRerunFailure：补录重跑失败告警；
	// 项目波动告警
	// projectFailureInstanceUpwardFluctuationAlarm： 当天失败实例数向上波动率超过阀值告警；
	// projectSuccessInstanceDownwardFluctuationAlarm：当天成功实例数向下波动率超过阀值告警；
	// 
	// 离线集成任务对账告警：
	// reconciliationFailure： 离线对账任务失败告警
	// reconciliationOvertime： 离线对账任务运行超时告警
	// reconciliationMismatch： 数据对账任务不一致条数超过阀值告警
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// 告警规则是否启用
	// 0-- 禁用 1--启用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警规则配置信息 成功告警无需配置；失败告警可以配置首次失败告警或者所有重试失败告警；超时配置需要配置超时类型及超时阀值；项目波动告警需要配置波动率及防抖周期配置
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// 告警级别 1.普通、2.重要、3.紧急
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警规则创建人uid
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// bundle 客户端绑定的告警规则:  为空是正常的告警规则，不为空则是对应bundle客户端绑定的规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundleId不为空 则表示绑定的bundle客户端名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`

	// 告警接收人配置列表
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`
}

type AlarmRuleDetail struct {
	// 失败触发时机 
	// 
	// 1 – 首次失败触发
	// 2 --所有重试完成触发 (默认)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trigger *int64 `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 补录重跑触发时机
	// 
	// 1 –  首次失败触发
	// 2 – 所有重试完成触发
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataBackfillOrRerunTrigger *int64 `json:"DataBackfillOrRerunTrigger,omitnil,omitempty" name:"DataBackfillOrRerunTrigger"`

	// 周期实例超时配置明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOutExtInfo []*TimeOutStrategyInfo `json:"TimeOutExtInfo,omitnil,omitempty" name:"TimeOutExtInfo"`

	// 重跑补录实例超时配置明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataBackfillOrRerunTimeOutExtInfo []*TimeOutStrategyInfo `json:"DataBackfillOrRerunTimeOutExtInfo,omitnil,omitempty" name:"DataBackfillOrRerunTimeOutExtInfo"`

	// 项目波动告警配置明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectInstanceStatisticsAlarmInfoList []*ProjectInstanceStatisticsAlarmInfo `json:"ProjectInstanceStatisticsAlarmInfoList,omitnil,omitempty" name:"ProjectInstanceStatisticsAlarmInfoList"`

	// 离线集成对账告警配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReconciliationExtInfo []*ReconciliationStrategyInfo `json:"ReconciliationExtInfo,omitnil,omitempty" name:"ReconciliationExtInfo"`

	// 监控对象的白名单配置
	MonitorWhiteTasks []*MonitorWhiteTask `json:"MonitorWhiteTasks,omitnil,omitempty" name:"MonitorWhiteTasks"`
}

type AlarmWayWebHook struct {
	// 告警渠道值
	// 7.企业微信群,8 飞书群 9 钉钉群 10 Slack群 11 Teams群
	AlarmWay *string `json:"AlarmWay,omitnil,omitempty" name:"AlarmWay"`

	// 告警群的webhook地址列表
	WebHooks []*string `json:"WebHooks,omitnil,omitempty" name:"WebHooks"`
}

// Predefined struct for user
type AssociateResourceGroupToProjectRequestParams struct {
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type AssociateResourceGroupToProjectRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *AssociateResourceGroupToProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateResourceGroupToProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateResourceGroupToProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateResourceGroupToProjectResponseParams struct {
	// 输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ResourceStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateResourceGroupToProjectResponse struct {
	*tchttp.BaseResponse
	Response *AssociateResourceGroupToProjectResponseParams `json:"Response"`
}

func (r *AssociateResourceGroupToProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateResourceGroupToProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackfillInstance struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 执行状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 执行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *string `json:"CostTime,omitnil,omitempty" name:"CostTime"`
}

type BackfillInstanceCollection struct {
	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 补录实例列表
	Items []*BackfillInstance `json:"Items,omitnil,omitempty" name:"Items"`
}

type BindProject struct {
	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目展示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectDisplayName *string `json:"ProjectDisplayName,omitnil,omitempty" name:"ProjectDisplayName"`
}

type BriefTask struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务责任人id，一个任务可能有多个责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUinList []*string `json:"OwnerUinList,omitnil,omitempty" name:"OwnerUinList"`
}

type BusinessMetadata struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagNames []*string `json:"TagNames,omitnil,omitempty" name:"TagNames"`
}

type CatalogInfo struct {
	// 目录名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ChildDependencyConfigPage struct {
	// 结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OpsTaskDepend `json:"Items,omitnil,omitempty" name:"Items"`
}

type CodeFile struct {
	// 脚本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// 脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeFileName *string `json:"CodeFileName,omitnil,omitempty" name:"CodeFileName"`

	// 脚本所有者 uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 脚本配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// 脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`

	// 最近一次操作人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 更新时间 yyyy-MM-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间 yyyy-MM-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 权限范围：SHARED, PRIVATE
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// 节点全路径，/aaa/bbb/ccc.ipynb，由各个节点的名称组成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 父文件夹路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`
}

type CodeFileConfig struct {
	// 高级运行参数,变量替换，map-json String,String
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// notebook kernel session信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookSessionInfo *NotebookSessionInfo `json:"NotebookSessionInfo,omitnil,omitempty" name:"NotebookSessionInfo"`
}

type CodeFolderNode struct {
	// 唯一标识
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 类型 folder，script
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否叶子节点
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 业务参数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 权限范围: SHARED, PRIVATE
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// 节点路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 目录/文件责任人uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建人
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 当前用户对节点拥有的权限	
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePermission *string `json:"NodePermission,omitnil,omitempty" name:"NodePermission"`

	// 子节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*CodeFolderNode `json:"Children,omitnil,omitempty" name:"Children"`

	// 父文件夹路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`
}

type CodeStudioFileActionResult struct {
	// 成功true，失败false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`
}

type CodeStudioFolderActionResult struct {
	// 成功true，失败false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type CodeStudioFolderResult struct {
	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type ColumnInfo struct {
	// 字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 字段的顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 是否为分区字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitnil,omitempty" name:"IsPartition"`
}

type CreateAlarmRuleData struct {
	// 告警规则唯一id
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

// Predefined struct for user
type CreateCodeFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件名称
	CodeFileName *string `json:"CodeFileName,omitnil,omitempty" name:"CodeFileName"`

	// 父文件夹path，例如/aaa/bbb/ccc，路径头需带斜杠，根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 代码文件配置
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// 代码文件内容
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

type CreateCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件名称
	CodeFileName *string `json:"CodeFileName,omitnil,omitempty" name:"CodeFileName"`

	// 父文件夹path，例如/aaa/bbb/ccc，路径头需带斜杠，根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 代码文件配置
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// 代码文件内容
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

func (r *CreateCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileName")
	delete(f, "ParentFolderPath")
	delete(f, "CodeFileConfig")
	delete(f, "CodeFileContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeFileResponseParams struct {
	// 结果
	Data *CodeFile `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateCodeFileResponseParams `json:"Response"`
}

func (r *CreateCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 父文件夹path，例如/aaa/bbb/ccc，路径头需带斜杠，根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`
}

type CreateCodeFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 父文件夹path，例如/aaa/bbb/ccc，路径头需带斜杠，根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`
}

func (r *CreateCodeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "ParentFolderPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCodeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCodeFolderResponseParams struct {
	// 成功true，失败false
	Data *CodeStudioFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCodeFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateCodeFolderResponseParams `json:"Response"`
}

func (r *CreateCodeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCodeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataBackfillPlanRequestParams struct {
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录任务集合
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 补录任务的数据时间配置
	DataBackfillRangeList []*DataBackfillRange `json:"DataBackfillRangeList,omitnil,omitempty" name:"DataBackfillRangeList"`

	// 时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 数据补录计划名称，不填则由系统随机生成一串字符
	DataBackfillPlanName *string `json:"DataBackfillPlanName,omitnil,omitempty" name:"DataBackfillPlanName"`

	// 检查父任务类型，取值范围：- NONE-全部不检查- ALL-检查全部上游父任务- MAKE_SCOPE-只在（当前补录计划）选中任务中检查,默认NONE不检查
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// 补录是否忽略事件依赖,默认true
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 自定义的工作流自依赖，yes或者no；如果不配置，则使用工作流原有自依赖
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// 自定义实例运行并发度, 如果不配置，则使用任务原有自依赖
	RedefineParallelNum *uint64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// 调度资源组id，为空则表示使用任务原有调度执行资源组
	SchedulerResourceGroupId *string `json:"SchedulerResourceGroupId,omitnil,omitempty" name:"SchedulerResourceGroupId"`

	// 集成任务资源组id，为空则表示使用任务原有调度执行资源组
	IntegrationResourceGroupId *string `json:"IntegrationResourceGroupId,omitnil,omitempty" name:"IntegrationResourceGroupId"`

	// 自定义参数，可以重新指定任务的参数，方便补录实例执行新的逻辑
	RedefineParamList []*KVPair `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`

	// 补录是实例数据时间顺序，生效必须满足2个条件:
	// 1. 必须同周期任务
	// 2. 优先按依赖关系执行，无依赖关系影响的情况下按配置执行顺序执行
	//  
	// 可选值
	// - NORMAL: 不设置
	// - ORDER: 顺序
	// - REVERSE: 逆序
	// 不设置默认为NORMAL
	DataTimeOrder *string `json:"DataTimeOrder,omitnil,omitempty" name:"DataTimeOrder"`

	// 补录实例重新生成周期，如果设置会重新指定补录任务实例的生成周期，目前只会将天实例转换成每月1号生成的实例
	// * MONTH_CYCLE: 月
	RedefineCycleType *string `json:"RedefineCycleType,omitnil,omitempty" name:"RedefineCycleType"`
}

type CreateDataBackfillPlanRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录任务集合
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 补录任务的数据时间配置
	DataBackfillRangeList []*DataBackfillRange `json:"DataBackfillRangeList,omitnil,omitempty" name:"DataBackfillRangeList"`

	// 时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 数据补录计划名称，不填则由系统随机生成一串字符
	DataBackfillPlanName *string `json:"DataBackfillPlanName,omitnil,omitempty" name:"DataBackfillPlanName"`

	// 检查父任务类型，取值范围：- NONE-全部不检查- ALL-检查全部上游父任务- MAKE_SCOPE-只在（当前补录计划）选中任务中检查,默认NONE不检查
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// 补录是否忽略事件依赖,默认true
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 自定义的工作流自依赖，yes或者no；如果不配置，则使用工作流原有自依赖
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// 自定义实例运行并发度, 如果不配置，则使用任务原有自依赖
	RedefineParallelNum *uint64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// 调度资源组id，为空则表示使用任务原有调度执行资源组
	SchedulerResourceGroupId *string `json:"SchedulerResourceGroupId,omitnil,omitempty" name:"SchedulerResourceGroupId"`

	// 集成任务资源组id，为空则表示使用任务原有调度执行资源组
	IntegrationResourceGroupId *string `json:"IntegrationResourceGroupId,omitnil,omitempty" name:"IntegrationResourceGroupId"`

	// 自定义参数，可以重新指定任务的参数，方便补录实例执行新的逻辑
	RedefineParamList []*KVPair `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`

	// 补录是实例数据时间顺序，生效必须满足2个条件:
	// 1. 必须同周期任务
	// 2. 优先按依赖关系执行，无依赖关系影响的情况下按配置执行顺序执行
	//  
	// 可选值
	// - NORMAL: 不设置
	// - ORDER: 顺序
	// - REVERSE: 逆序
	// 不设置默认为NORMAL
	DataTimeOrder *string `json:"DataTimeOrder,omitnil,omitempty" name:"DataTimeOrder"`

	// 补录实例重新生成周期，如果设置会重新指定补录任务实例的生成周期，目前只会将天实例转换成每月1号生成的实例
	// * MONTH_CYCLE: 月
	RedefineCycleType *string `json:"RedefineCycleType,omitnil,omitempty" name:"RedefineCycleType"`
}

func (r *CreateDataBackfillPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataBackfillPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "DataBackfillRangeList")
	delete(f, "TimeZone")
	delete(f, "DataBackfillPlanName")
	delete(f, "CheckParentType")
	delete(f, "SkipEventListening")
	delete(f, "RedefineSelfWorkflowDependency")
	delete(f, "RedefineParallelNum")
	delete(f, "SchedulerResourceGroupId")
	delete(f, "IntegrationResourceGroupId")
	delete(f, "RedefineParamList")
	delete(f, "DataTimeOrder")
	delete(f, "RedefineCycleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataBackfillPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataBackfillPlanResponseParams struct {
	// 数据补录计划创建结果
	Data *CreateDataReplenishmentPlan `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataBackfillPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataBackfillPlanResponseParams `json:"Response"`
}

func (r *CreateDataBackfillPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataBackfillPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataReplenishmentPlan struct {
	// 补录计划Id
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`
}

// Predefined struct for user
type CreateDataSourceRequestParams struct {
	// 数据源项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类型:枚举值
	// 
	// - MYSQL
	// - TENCENT_MYSQL
	// - POSTGRE
	// - ORACLE
	// - SQLSERVER
	// - FTP
	// - HIVE
	// - HUDI
	// - HDFS
	// - ICEBERG
	// - KAFKA
	// - DTS_KAFKA
	// - HBASE
	// - SPARK
	// - TBASE
	// - DB2
	// - DM
	// - GAUSSDB
	// - GBASE
	// - IMPALA
	// - ES
	// - TENCENT_ES
	// - GREENPLUM
	// - SAP_HANA
	// - SFTP
	// - OCEANBASE
	// - CLICKHOUSE
	// - KUDU
	// - VERTICA
	// - REDIS
	// - COS
	// - DLC
	// - DORIS
	// - CKAFKA
	// - S3_DATAINSIGHT
	// - TDSQL
	// - TDSQL_MYSQL
	// - MONGODB
	// - TENCENT_MONGODB
	// - REST_API
	// - TiDB
	// - StarRocks
	// - Trino
	// - Kyuubi
	// - TCHOUSE_X
	// - TCHOUSE_P
	// - TCHOUSE_C
	// - TCHOUSE_D
	// - INFLUXDB
	// - BIG_QUERY
	// - SSH
	// - BLOB
	// - TDSQL_POSTGRE
	// - GDB
	// - TDENGINE
	// - TDSQLC
	// - FileSystem
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	// 
	// > deployType: 
	// CONNSTR_PUBLICDB(公网实例) 
	// CONNSTR_CVMDB(自建实例)
	// INSTANCE(云实例)
	// 
	// ```
	// mysql: 自建实例
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:mysql://1.1.1.1:1111/database",
	//     "username": "root",
	//     "password": "root",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "MYSQL"
	// }
	// mysql: 云实例
	// {
	//     "instanceid": "cdb-12uxdo5e",
	//     "db": "db",
	//     "region": "ap-shanghai",
	//     "username": "msyql",
	//     "password": "mysql",
	//     "deployType": "INSTANCE",
	//     "type": "TENCENT_MYSQL"
	// }
	// sql_server: 
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:sqlserver://1.1.1.1:223;DatabaseName=database",
	//     "username": "user_1",
	//     "password": "pass_2",
	//     "type": "SQLSERVER"
	// }
	// redis:
	//     redisType:
	//     -NO_ACCOUT(免账号)
	//     -SELF_ACCOUNT(自定义账号)
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username":""
	//     "password": "pass",
	//     "ip": "1.1.1.1",
	//     "port": "6379",
	//     "redisType": "NO_ACCOUT",
	//     "type": "REDIS"
	// }
	// oracle: 
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:oracle:thin:@1.1.1.1:1521:prod",
	//     "username": "oracle",
	//     "password": "pass",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "ORACLE"
	// }
	// mongodb:
	//     advanceParams(自定义参数，会拼接至url后)
	// {
	//     "advanceParams": [
	//         {
	//             "key": "authSource",
	//             "value": "auth"
	//         }
	//     ],
	//     "db": "admin",
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "MONGODB",
	//     "host": "1.1.1.1:9200"
	// }
	// postgresql:
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:postgresql://1.1.1.1:1921/database",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "POSTGRE"
	// }
	// kafka:
	//     authType:
	//         - sasl
	//         - jaas
	//         - sasl_plaintext
	//         - sasl_ssl
	//         - GSSAPI
	//     ssl:
	//         -PLAIN
	//         -GSSAPI
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "host": "1.1.1.1:9092",
	//     "ssl": "GSSAPI",
	//     "authType": "sasl",
	//     "type": "KAFKA",
	//     "principal": "aaaa",
	//     "serviceName": "kafka"
	// }
	// 
	// cos:
	// {
	//     "region": "ap-shanghai",
	//     "deployType": "INSTANCE",
	//     "secretId": "aaaaa",
	//     "secretKey": "sssssss",
	//     "bucket": "aaa",
	//     "type": "COS"
	// }
	// 
	// ```
	ProdConProperties *string `json:"ProdConProperties,omitnil,omitempty" name:"ProdConProperties"`

	// 开发环境数据源配置信息，若项目为标准模式，则此字段必填
	DevConProperties *string `json:"DevConProperties,omitnil,omitempty" name:"DevConProperties"`

	// 生产环境数据源文件上传
	ProdFileUpload *DataSourceFileUpload `json:"ProdFileUpload,omitnil,omitempty" name:"ProdFileUpload"`

	// 开发环境数据源文件上传
	DevFileUpload *DataSourceFileUpload `json:"DevFileUpload,omitnil,omitempty" name:"DevFileUpload"`

	// 数据源展示名，为了可视化查看
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类型:枚举值
	// 
	// - MYSQL
	// - TENCENT_MYSQL
	// - POSTGRE
	// - ORACLE
	// - SQLSERVER
	// - FTP
	// - HIVE
	// - HUDI
	// - HDFS
	// - ICEBERG
	// - KAFKA
	// - DTS_KAFKA
	// - HBASE
	// - SPARK
	// - TBASE
	// - DB2
	// - DM
	// - GAUSSDB
	// - GBASE
	// - IMPALA
	// - ES
	// - TENCENT_ES
	// - GREENPLUM
	// - SAP_HANA
	// - SFTP
	// - OCEANBASE
	// - CLICKHOUSE
	// - KUDU
	// - VERTICA
	// - REDIS
	// - COS
	// - DLC
	// - DORIS
	// - CKAFKA
	// - S3_DATAINSIGHT
	// - TDSQL
	// - TDSQL_MYSQL
	// - MONGODB
	// - TENCENT_MONGODB
	// - REST_API
	// - TiDB
	// - StarRocks
	// - Trino
	// - Kyuubi
	// - TCHOUSE_X
	// - TCHOUSE_P
	// - TCHOUSE_C
	// - TCHOUSE_D
	// - INFLUXDB
	// - BIG_QUERY
	// - SSH
	// - BLOB
	// - TDSQL_POSTGRE
	// - GDB
	// - TDENGINE
	// - TDSQLC
	// - FileSystem
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	// 
	// > deployType: 
	// CONNSTR_PUBLICDB(公网实例) 
	// CONNSTR_CVMDB(自建实例)
	// INSTANCE(云实例)
	// 
	// ```
	// mysql: 自建实例
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:mysql://1.1.1.1:1111/database",
	//     "username": "root",
	//     "password": "root",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "MYSQL"
	// }
	// mysql: 云实例
	// {
	//     "instanceid": "cdb-12uxdo5e",
	//     "db": "db",
	//     "region": "ap-shanghai",
	//     "username": "msyql",
	//     "password": "mysql",
	//     "deployType": "INSTANCE",
	//     "type": "TENCENT_MYSQL"
	// }
	// sql_server: 
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:sqlserver://1.1.1.1:223;DatabaseName=database",
	//     "username": "user_1",
	//     "password": "pass_2",
	//     "type": "SQLSERVER"
	// }
	// redis:
	//     redisType:
	//     -NO_ACCOUT(免账号)
	//     -SELF_ACCOUNT(自定义账号)
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username":""
	//     "password": "pass",
	//     "ip": "1.1.1.1",
	//     "port": "6379",
	//     "redisType": "NO_ACCOUT",
	//     "type": "REDIS"
	// }
	// oracle: 
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:oracle:thin:@1.1.1.1:1521:prod",
	//     "username": "oracle",
	//     "password": "pass",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "ORACLE"
	// }
	// mongodb:
	//     advanceParams(自定义参数，会拼接至url后)
	// {
	//     "advanceParams": [
	//         {
	//             "key": "authSource",
	//             "value": "auth"
	//         }
	//     ],
	//     "db": "admin",
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "MONGODB",
	//     "host": "1.1.1.1:9200"
	// }
	// postgresql:
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:postgresql://1.1.1.1:1921/database",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "POSTGRE"
	// }
	// kafka:
	//     authType:
	//         - sasl
	//         - jaas
	//         - sasl_plaintext
	//         - sasl_ssl
	//         - GSSAPI
	//     ssl:
	//         -PLAIN
	//         -GSSAPI
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "host": "1.1.1.1:9092",
	//     "ssl": "GSSAPI",
	//     "authType": "sasl",
	//     "type": "KAFKA",
	//     "principal": "aaaa",
	//     "serviceName": "kafka"
	// }
	// 
	// cos:
	// {
	//     "region": "ap-shanghai",
	//     "deployType": "INSTANCE",
	//     "secretId": "aaaaa",
	//     "secretKey": "sssssss",
	//     "bucket": "aaa",
	//     "type": "COS"
	// }
	// 
	// ```
	ProdConProperties *string `json:"ProdConProperties,omitnil,omitempty" name:"ProdConProperties"`

	// 开发环境数据源配置信息，若项目为标准模式，则此字段必填
	DevConProperties *string `json:"DevConProperties,omitnil,omitempty" name:"DevConProperties"`

	// 生产环境数据源文件上传
	ProdFileUpload *DataSourceFileUpload `json:"ProdFileUpload,omitnil,omitempty" name:"ProdFileUpload"`

	// 开发环境数据源文件上传
	DevFileUpload *DataSourceFileUpload `json:"DevFileUpload,omitnil,omitempty" name:"DevFileUpload"`

	// 数据源展示名，为了可视化查看
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "ProdConProperties")
	delete(f, "DevConProperties")
	delete(f, "ProdFileUpload")
	delete(f, "DevFileUpload")
	delete(f, "DisplayName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataSourceResponseParams struct {
	// 主键ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataSourceResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataSourceResponseParams `json:"Response"`
}

func (r *CreateDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFolderResult struct {
	// 创建成功的文件夹ID。如果创建失败则报错。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

// Predefined struct for user
type CreateOpsAlarmRuleRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则名称
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// 监控对象业务id列表，根据MonitorType 的设置传入不同的业务id，如下1（任务）： MonitorObjectIds为任务id列表2（工作流）： MonitorObjectIds 为工作流id列表(工作流id可从接口ListWorkflows获取)3（项目）：  MonitorObjectIds为项目id列表
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// 告警规则监控类型 failure：失败告警 ；overtime：超时告警 success：成功告警; backTrackingOrRerunSuccess: 补录重跑成功告警 backTrackingOrRerunFailure：补录重跑失败告警； 
	// 项目波动告警 projectFailureInstanceUpwardFluctuationAlarm： 当天失败实例数向上波动率超过阀值告警； projectSuccessInstanceDownwardFluctuationAlarm：当天成功实例数向下波动率超过阀值告警； 
	// 离线集成任务对账告警： reconciliationFailure： 离线对账任务失败告警 reconciliationOvertime： 离线对账任务运行超时告警 reconciliationMismatch： 数据对账任务不一致条数超过阀值告警
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// 告警接收人配置信息
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// 监控对象类型, 
	// 任务维度监控： 可按照任务/工作流/项目来配置：1.任务、2.工作流、3.项目（默认为1.任务）
	// 项目维度监控： 项目整体任务波动告警，  7：项目波动监控告警
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 告警规则配置信息
	// 成功告警无需配置；失败告警可以配置首次失败告警或者所有重试失败告警；超时配置需要配置超时类型及超时阀值；项目波动告警需要配置波动率及防抖周期配置
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// 告警级别 1.普通、2.重要、3.紧急（默认1.普通）
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则名称
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// 监控对象业务id列表，根据MonitorType 的设置传入不同的业务id，如下1（任务）： MonitorObjectIds为任务id列表2（工作流）： MonitorObjectIds 为工作流id列表(工作流id可从接口ListWorkflows获取)3（项目）：  MonitorObjectIds为项目id列表
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// 告警规则监控类型 failure：失败告警 ；overtime：超时告警 success：成功告警; backTrackingOrRerunSuccess: 补录重跑成功告警 backTrackingOrRerunFailure：补录重跑失败告警； 
	// 项目波动告警 projectFailureInstanceUpwardFluctuationAlarm： 当天失败实例数向上波动率超过阀值告警； projectSuccessInstanceDownwardFluctuationAlarm：当天成功实例数向下波动率超过阀值告警； 
	// 离线集成任务对账告警： reconciliationFailure： 离线对账任务失败告警 reconciliationOvertime： 离线对账任务运行超时告警 reconciliationMismatch： 数据对账任务不一致条数超过阀值告警
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// 告警接收人配置信息
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// 监控对象类型, 
	// 任务维度监控： 可按照任务/工作流/项目来配置：1.任务、2.工作流、3.项目（默认为1.任务）
	// 项目维度监控： 项目整体任务波动告警，  7：项目波动监控告警
	MonitorObjectType *uint64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 告警规则配置信息
	// 成功告警无需配置；失败告警可以配置首次失败告警或者所有重试失败告警；超时配置需要配置超时类型及超时阀值；项目波动告警需要配置波动率及防抖周期配置
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// 告警级别 1.普通、2.重要、3.紧急（默认1.普通）
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleName")
	delete(f, "MonitorObjectIds")
	delete(f, "AlarmTypes")
	delete(f, "AlarmGroups")
	delete(f, "MonitorObjectType")
	delete(f, "AlarmRuleDetail")
	delete(f, "AlarmLevel")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpsAlarmRuleResponseParams struct {
	// 告警规则唯一id
	Data *CreateAlarmRuleData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *CreateOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectMemberRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// 角色id
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type CreateProjectMemberRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`

	// 角色id
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *CreateProjectMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserUins")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectMemberResponseParams struct {
	// 返回数据
	Data *ProjectResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectMemberResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectMemberResponseParams `json:"Response"`
}

func (r *CreateProjectMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// 项目基本信息
	Project *ProjectRequest `json:"Project,omitnil,omitempty" name:"Project"`

	// DLC绑定集群信息
	DLCInfo *DLCClusterInfo `json:"DLCInfo,omitnil,omitempty" name:"DLCInfo"`

	// 绑定资源组的id列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目基本信息
	Project *ProjectRequest `json:"Project,omitnil,omitempty" name:"Project"`

	// DLC绑定集群信息
	DLCInfo *DLCClusterInfo `json:"DLCInfo,omitnil,omitempty" name:"DLCInfo"`

	// 绑定资源组的id列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Project")
	delete(f, "DLCInfo")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// 创建项目结果
	Data *CreateProjectResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResult struct {
	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

// Predefined struct for user
type CreateResourceFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件名称, 尽可能和上传文件名保持一致
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// cos存储桶名称, 可从GetResourceCosPath接口获取
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// BucketName桶对应的cos存储桶区域
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 项目中资源文件上传的路径, 取值示例: /wedata/qxxxm/, 根目录,请使用/即可
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// - 上传文件及手填两种方式只能选择其一，如果两者均提供，取值顺序为文件>手填值
	// -   手填值必须是存在的cos路径, /datastudio/resource/ 为固定前缀, projectId 为项目ID,需传入具体值, parentFolderPath为父文件夹路径, name为文件名, 手填值取值示例:     /datastudio/resource/projectId/parentFolderPath/name 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// bundle客户端ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundle客户端信息
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type CreateResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件名称, 尽可能和上传文件名保持一致
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// cos存储桶名称, 可从GetResourceCosPath接口获取
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// BucketName桶对应的cos存储桶区域
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 项目中资源文件上传的路径, 取值示例: /wedata/qxxxm/, 根目录,请使用/即可
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// - 上传文件及手填两种方式只能选择其一，如果两者均提供，取值顺序为文件>手填值
	// -   手填值必须是存在的cos路径, /datastudio/resource/ 为固定前缀, projectId 为项目ID,需传入具体值, parentFolderPath为父文件夹路径, name为文件名, 手填值取值示例:     /datastudio/resource/projectId/parentFolderPath/name 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// bundle客户端ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundle客户端信息
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *CreateResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceName")
	delete(f, "BucketName")
	delete(f, "CosRegion")
	delete(f, "ParentFolderPath")
	delete(f, "ResourceFile")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceFileResponseParams struct {
	// 创建资源文件结果
	Data *CreateResourceFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceFileResponseParams `json:"Response"`
}

func (r *CreateResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceFileResult struct {
	// 资源文件ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

// Predefined struct for user
type CreateResourceFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹绝对路径, 取值示例 /wedata/test, 根目录,请使用/即可
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type CreateResourceFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹绝对路径, 取值示例 /wedata/test, 根目录,请使用/即可
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *CreateResourceFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceFolderResponseParams struct {
	// 创建文件夹结果，如果创建失败则报错。
	Data *CreateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourceFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceFolderResponseParams `json:"Response"`
}

func (r *CreateResourceFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceGroupRequestParams struct {
	// 资源组名称。创建通用资源组的名称，必须以字母开头，可包含字母、数字、下划线（_），最多 64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开通的资源组信息
	Type *ResourceType `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否自动续费
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitnil,omitempty" name:"AutoRenewEnabled"`

	// 购买时长，单位月
	PurchasePeriod *int64 `json:"PurchasePeriod,omitnil,omitempty" name:"PurchasePeriod"`

	// vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网
	SubNet *string `json:"SubNet,omitnil,omitempty" name:"SubNet"`

	// 资源购买地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 关联项目空间项目id
	AssociatedProjectId *string `json:"AssociatedProjectId,omitnil,omitempty" name:"AssociatedProjectId"`

	// 资源组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资源组名称。创建通用资源组的名称，必须以字母开头，可包含字母、数字、下划线（_），最多 64 个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开通的资源组信息
	Type *ResourceType `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否自动续费
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitnil,omitempty" name:"AutoRenewEnabled"`

	// 购买时长，单位月
	PurchasePeriod *int64 `json:"PurchasePeriod,omitnil,omitempty" name:"PurchasePeriod"`

	// vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网
	SubNet *string `json:"SubNet,omitnil,omitempty" name:"SubNet"`

	// 资源购买地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 关联项目空间项目id
	AssociatedProjectId *string `json:"AssociatedProjectId,omitnil,omitempty" name:"AssociatedProjectId"`

	// 资源组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "AutoRenewEnabled")
	delete(f, "PurchasePeriod")
	delete(f, "VpcId")
	delete(f, "SubNet")
	delete(f, "ResourceRegion")
	delete(f, "AssociatedProjectId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceGroupResponseParams struct {
	// 是否开启成功
	Data *ResourceResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceGroupResponseParams `json:"Response"`
}

func (r *CreateResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLFolderRequestParams struct {
	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，/aaa/bbb/ccc，路径头需带斜杠，查询根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type CreateSQLFolderRequest struct {
	*tchttp.BaseRequest
	
	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，/aaa/bbb/ccc，路径头需带斜杠，查询根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *CreateSQLFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderName")
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSQLFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLFolderResponseParams struct {
	// 成功true，失败false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SqlCreateResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSQLFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateSQLFolderResponseParams `json:"Response"`
}

func (r *CreateSQLFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLScriptRequestParams struct {
	// 脚本名称
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，/aaa/bbb/ccc，根目录为空字符串或/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 数据探索脚本配置
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// 脚本内容，如有值，则要将内容进行base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type CreateSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本名称
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，/aaa/bbb/ccc，根目录为空字符串或/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 数据探索脚本配置
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// 脚本内容，如有值，则要将内容进行base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *CreateSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptName")
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "ScriptConfig")
	delete(f, "ScriptContent")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSQLScriptResponseParams struct {
	// 结果
	Data *SQLScript `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *CreateSQLScriptResponseParams `json:"Response"`
}

func (r *CreateSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskBaseAttribute struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型ID：
	// 
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 32:DLC SQL
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务负责人ID，默认为当前用户
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 任务描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`
}

type CreateTaskConfiguration struct {
	// 资源组ID： 需要通过 DescribeNormalSchedulerExecutorGroups 获取 ExecutorGroupId
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 代码内容的Base64编码
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// 任务扩展属性配置列表
	TaskExtConfigurationList []*TaskExtParameter `json:"TaskExtConfigurationList,omitnil,omitempty" name:"TaskExtConfigurationList"`

	// 集群ID
	DataCluster *string `json:"DataCluster,omitnil,omitempty" name:"DataCluster"`

	// 指定的运行节点
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 资源池队列名称，需要通过 DescribeProjectClusterQueues 获取
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// 来源数据源ID, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 目标数据源ID, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// 调度参数
	TaskSchedulingParameterList []*TaskSchedulingParameter `json:"TaskSchedulingParameterList,omitnil,omitempty" name:"TaskSchedulingParameterList"`

	// Bundle使用的ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle信息
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务基本属性
	TaskBaseAttribute *CreateTaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// 任务配置
	TaskConfiguration *CreateTaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// 任务调度配置
	TaskSchedulerConfiguration *CreateTaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务基本属性
	TaskBaseAttribute *CreateTaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// 任务配置
	TaskConfiguration *CreateTaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// 任务调度配置
	TaskSchedulerConfiguration *CreateTaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskBaseAttribute")
	delete(f, "TaskConfiguration")
	delete(f, "TaskSchedulerConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 任务ID
	Data *CreateTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResult struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CreateTaskSchedulerConfiguration struct {
	// 周期类型：默认为 DAY_CYCLE
	// 
	// 支持的类型为 
	// 
	// * ONEOFF_CYCLE: 一次性
	// * YEAR_CYCLE: 年
	// * MONTH_CYCLE: 月
	// * WEEK_CYCLE: 周
	// * DAY_CYCLE: 天
	// * HOUR_CYCLE: 小时
	// * MINUTE_CYCLE: 分钟
	// * CRONTAB_CYCLE: crontab表达式类型
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 时区，默认为 UTC+8
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// Cron表达式，默认为 0 0 0 * * ? * 
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 生效日期，默认为当前日期的 00:00:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期，默认为 2099-12-31 23:59:59
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 执行时间 左闭区间，默认 00:00
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间 右闭区间，默认 23:59
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 调度类型: 0 正常调度 1 空跑调度，默认为 0
	ScheduleRunType *string `json:"ScheduleRunType,omitnil,omitempty" name:"ScheduleRunType"`

	// 日历调度 取值为 0 和 1， 1为打开，0为关闭，默认为0
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// 日历调度 日历 ID
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`

	// 自依赖, 默认值 serial, 取值为：parallel(并行), serial(串行), orderly(有序)
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 上游依赖数组
	UpstreamDependencyConfigList []*DependencyTaskBrief `json:"UpstreamDependencyConfigList,omitnil,omitempty" name:"UpstreamDependencyConfigList"`

	// 事件数组
	EventListenerList []*EventListener `json:"EventListenerList,omitnil,omitempty" name:"EventListenerList"`

	// 任务调度优先级 运行优先级 4高 5中 6低 , 默认:6
	RunPriority *string `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// 重试策略 重试等待时间,单位分钟: 默认: 5
	RetryWait *string `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// 重试策略 最大尝试次数, 默认: 4
	MaxRetryAttempts *string `json:"MaxRetryAttempts,omitnil,omitempty" name:"MaxRetryAttempts"`

	// 超时处理策略 运行耗时超时（单位：分钟）默认为 -1
	ExecutionTTL *string `json:"ExecutionTTL,omitnil,omitempty" name:"ExecutionTTL"`

	// 超时处理策略 等待总时长耗时超时（单位：分钟）默认为 -1
	WaitExecutionTotalTTL *string `json:"WaitExecutionTotalTTL,omitnil,omitempty" name:"WaitExecutionTotalTTL"`

	// 重跑&补录配置, 默认为 ALL; , ALL 运行成功或失败后皆可重跑或补录, FAILURE 运行成功后不可重跑或补录，运行失败后可重跑或补录, NONE 运行成功或失败后皆不可重跑或补录;
	AllowRedoType *string `json:"AllowRedoType,omitnil,omitempty" name:"AllowRedoType"`

	// 输出参数数组
	ParamTaskOutList []*OutTaskParameter `json:"ParamTaskOutList,omitnil,omitempty" name:"ParamTaskOutList"`

	// 输入参数数组
	ParamTaskInList []*InTaskParameter `json:"ParamTaskInList,omitnil,omitempty" name:"ParamTaskInList"`

	// 产出登记
	TaskOutputRegistryList []*TaskDataRegistry `json:"TaskOutputRegistryList,omitnil,omitempty" name:"TaskOutputRegistryList"`

	// **实例生成策略**
	// * T_PLUS_0: T+0生成,默认策略
	// * T_PLUS_1: T+1生成
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`
}

// Predefined struct for user
type CreateWorkflowFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹绝对路径，如/abc/de，如果是根目录则传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 要创建的文件夹名字
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type CreateWorkflowFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹绝对路径，如/abc/de，如果是根目录则传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 要创建的文件夹名字
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *CreateWorkflowFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowFolderResponseParams struct {
	// 创建文件夹结果，如果创建失败则报错。
	Data *CreateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowFolderResponseParams `json:"Response"`
}

func (r *CreateWorkflowFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 所属文件夹路径
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 工作流类型,取值示例：cycle 周期工作流；manual 手动工作流，默认传入cycle
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 工作流描述
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流负责人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 工作流参数
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 统一调度信息
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId项
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle信息
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 所属文件夹路径
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 工作流类型,取值示例：cycle 周期工作流；manual 手动工作流，默认传入cycle
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 工作流描述
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流负责人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 工作流参数
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 统一调度信息
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId项
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle信息
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowName")
	delete(f, "ParentFolderPath")
	delete(f, "WorkflowType")
	delete(f, "WorkflowDesc")
	delete(f, "OwnerUin")
	delete(f, "WorkflowParams")
	delete(f, "WorkflowSchedulerConfiguration")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// 返回工作流ID
	Data *CreateWorkflowResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowResponseParams `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWorkflowResult struct {
	// 创建成功后的工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DLCClusterInfo struct {
	// dlc资源名称（需要添加角色Uin到dlc中，否则可能获取不到资源）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeResources []*string `json:"ComputeResources,omitnil,omitempty" name:"ComputeResources"`

	// dlc地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 指定DLC集群的默认数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDatabase *string `json:"DefaultDatabase,omitnil,omitempty" name:"DefaultDatabase"`

	// 集群配置标记（ 仅对标准模式的项目生效并且标准模式必填），枚举值：
	// - Prod  (生产环境)
	// - Dev  (开发环境)
	StandardModeEnvTag *string `json:"StandardModeEnvTag,omitnil,omitempty" name:"StandardModeEnvTag"`

	// 访问账号（ 仅对标准模式的项目生效并且标准模式必填），用于提交dlc任务的账号
	// 建议使用指定子账号，给子账号设置对应库表的权限；任务负责人模式在负责人离职后容易造成任务失败；主账号模式在多个项目权限不同的情况下不易做权限控制。
	// 
	// 枚举值：
	// - TASK_RUNNER （任务负责人）
	// - OWNER （主账号模式）
	// - SUB （子账号模式）
	AccessAccount *string `json:"AccessAccount,omitnil,omitempty" name:"AccessAccount"`

	// 子账号id（ 仅对标准模式的项目生效），AccessAccount为子账号模式时，需要指定子账号的id信息，其他模式不需要指定
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`
}

type DataBackfill struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据补录计划id
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`

	// 数据补录计划名称
	DataBackfillPlanName *string `json:"DataBackfillPlanName,omitnil,omitempty" name:"DataBackfillPlanName"`

	// 补录任务集合
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 补录任务的数据配置列表
	DataBackfillRangeList []*DataBackfillRange `json:"DataBackfillRangeList,omitnil,omitempty" name:"DataBackfillRangeList"`

	// 检查父任务类型，取值范围：- NONE-全部不检查- ALL-检查全部上游父任务- MAKE_SCOPE-只在（当前补录计划）选中任务中检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// 补录是否忽略事件依赖	
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 自定义实例运行并发度, 返回为null或者不返回，则表示任务原有自依赖
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedefineParallelNum *uint64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// 自定义的工作流自依赖，yes或者no；如果不配置，则使用工作流原有自依赖
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// 调度资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerResourceGroupId *string `json:"SchedulerResourceGroupId,omitnil,omitempty" name:"SchedulerResourceGroupId"`

	// 集成资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntegrationResourceGroupId *string `json:"IntegrationResourceGroupId,omitnil,omitempty" name:"IntegrationResourceGroupId"`

	// 补录自定义的生成周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedefineCycleType *string `json:"RedefineCycleType,omitnil,omitempty" name:"RedefineCycleType"`

	// 自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedefineParamList []*KVPair `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`

	// 补录任务的执行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 补录任务的执行结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 创建用户id
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 补录计划实例完成百分数
	CompletePercent *uint64 `json:"CompletePercent,omitnil,omitempty" name:"CompletePercent"`

	// 补录计划实例成功百分数
	SuccessPercent *uint64 `json:"SuccessPercent,omitnil,omitempty" name:"SuccessPercent"`

	// 补录是实例数据时间顺序，生效必须满足2个条件:1. 必须同周期任务2. 优先按依赖关系执行，无依赖关系影响的情况下按配置执行顺序执行 可选值- NORMAL: 不设置- ORDER: 顺序- REVERSE: 逆序不设置默认为NORMAL
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTimeOrder *string `json:"DataTimeOrder,omitnil,omitempty" name:"DataTimeOrder"`
}

type DataBackfillRange struct {
	// 开始日期，格式yyyy-MM-dd 表示从指定日期的00:00:00开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd，表示从指定日期的 23:59:59结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 在[StartDate, EndDate]之间每天的开始时间点，格式HH:mm,只针对小时及周期小于小时的任务生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 在[StartDate, EndDate]之间每天的结束时间点，格式HH:mm,只针对小时及周期小于小时的任务生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`
}

type DataSource struct {
	// 归属项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据源类型:枚举值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源展示名，为了可视化查看
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据源描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 归属项目Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 数据源创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUser *string `json:"CreateUser,omitnil,omitempty" name:"CreateUser"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyUser *string `json:"ModifyUser,omitnil,omitempty" name:"ModifyUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProdConProperties *string `json:"ProdConProperties,omitnil,omitempty" name:"ProdConProperties"`

	// 同params 内容为开发数据源的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevConProperties *string `json:"DevConProperties,omitnil,omitempty" name:"DevConProperties"`

	// 数据源类别：
	// 
	// - DB ---自定义源
	// - CLUSTER --- 系统源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`
}

type DataSourceFileUpload struct {
	// Truststore认证文件，默认文件名truststore.jks
	TrustStore *string `json:"TrustStore,omitnil,omitempty" name:"TrustStore"`

	// Keystore认证文件，默认文件名keystore.jks
	KeyStore *string `json:"KeyStore,omitnil,omitempty" name:"KeyStore"`

	// core-site.xml文件
	CoreSite *string `json:"CoreSite,omitnil,omitempty" name:"CoreSite"`

	// hdfs-site.xml文件
	HdfsSite *string `json:"HdfsSite,omitnil,omitempty" name:"HdfsSite"`

	// hive-site.xml文件
	HiveSite *string `json:"HiveSite,omitnil,omitempty" name:"HiveSite"`

	// hbase-site文件
	HBASESite *string `json:"HBASESite,omitnil,omitempty" name:"HBASESite"`

	// keytab文件，默认文件名[数据源名].keytab
	KeyTab *string `json:"KeyTab,omitnil,omitempty" name:"KeyTab"`

	// krb5.conf文件
	KRB5Conf *string `json:"KRB5Conf,omitnil,omitempty" name:"KRB5Conf"`

	// 私钥,默认文件名private_key.pem
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 公钥,默认文件名public_key.pem
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`
}

type DataSourceInfo struct {
	// 分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DataSource `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type DataSourceResult struct {
	// 操作是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *int64 `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`
}

type DataSourceStatus struct {
	// 数据源操作状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type DatabaseInfo struct {
	// 数据库GUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guid *string `json:"Guid,omitnil,omitempty" name:"Guid"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据库位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据库存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`
}

type DatasourceRelationTaskInfo struct {
	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 数据源关联任务信息，本期仅支持数据开发任务接口返回数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfo []*RelateTask `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`
}

type DeleteAlarmRuleResult struct {
	// 是否删除成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteCodeFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件ID，参数值来自CreateCodeFile接口的返回
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`
}

type DeleteCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件ID，参数值来自CreateCodeFile接口的返回
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`
}

func (r *DeleteCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeFileResponseParams struct {
	// 执行结果
	Data *CodeStudioFileActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCodeFileResponseParams `json:"Response"`
}

func (r *DeleteCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，参数值来自CreateCodeFolder接口的返回
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type DeleteCodeFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，参数值来自CreateCodeFolder接口的返回
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *DeleteCodeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCodeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCodeFolderResponseParams struct {
	// 执行结果
	Data *CodeStudioFolderActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCodeFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCodeFolderResponseParams `json:"Response"`
}

func (r *DeleteCodeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCodeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataSourceRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataSourceResponseParams struct {
	// 是否删除成功
	Data *DataSourceStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataSourceResponseParams `json:"Response"`
}

func (r *DeleteDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFolderResult struct {
	// 删除状态,true表示成功，false表示失败
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteLineageRequestParams struct {
	// 需要删除的血缘关系列表
	Relations []*LineagePair `json:"Relations,omitnil,omitempty" name:"Relations"`
}

type DeleteLineageRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的血缘关系列表
	Relations []*LineagePair `json:"Relations,omitnil,omitempty" name:"Relations"`
}

func (r *DeleteLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Relations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLineageResponseParams struct {
	// 删除结果
	Data *OperateResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLineageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLineageResponseParams `json:"Response"`
}

func (r *DeleteLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOpsAlarmRuleRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则唯一id，接口CreateAlarmRule返回
	// 与AlarmRuleName二选一
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

type DeleteOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则唯一id，接口CreateAlarmRule返回
	// 与AlarmRuleName二选一
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

func (r *DeleteOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOpsAlarmRuleResponseParams struct {
	// 是否删除成功
	Data *DeleteAlarmRuleResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *DeleteOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectMemberRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID列表
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`
}

type DeleteProjectMemberRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID列表
	UserUins []*string `json:"UserUins,omitnil,omitempty" name:"UserUins"`
}

func (r *DeleteProjectMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectMemberResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectMemberResponseParams `json:"Response"`
}

func (r *DeleteProjectMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源ID, 可通过ListResourceFiles接口获取
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DeleteResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源ID, 可通过ListResourceFiles接口获取
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DeleteResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFileResponseParams struct {
	// 资源删除结果
	Data *DeleteResourceFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceFileResponseParams `json:"Response"`
}

func (r *DeleteResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResourceFileResult struct {
	// true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteResourceFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID, 可通过ListResourceFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type DeleteResourceFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID, 可通过ListResourceFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *DeleteResourceFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFolderResponseParams struct {
	// true代表删除成功，false代表删除失败
	Data *DeleteFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceFolderResponseParams `json:"Response"`
}

func (r *DeleteResourceFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupRequestParams struct {
	// 资源组id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupResponseParams struct {
	// 是否销毁成功
	Data *ResourceStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceGroupResponseParams `json:"Response"`
}

func (r *DeleteResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLFolderRequestParams struct {
	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteSQLFolderRequest struct {
	*tchttp.BaseRequest
	
	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteSQLFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSQLFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLFolderResponseParams struct {
	// 操作结果
	Data *SQLContentActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSQLFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSQLFolderResponseParams `json:"Response"`
}

func (r *DeleteSQLFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLScriptRequestParams struct {
	// 探索脚本Id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// 探索脚本Id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSQLScriptResponseParams struct {
	// 执行结果
	Data *SQLContentActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSQLScriptResponseParams `json:"Response"`
}

func (r *DeleteSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	// 和VirtualTaskId选填一个
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务操作是否消息通知下游任务责任人true：通知
	// false：不通知
	// 不传默认false
	OperateInform *bool `json:"OperateInform,omitnil,omitempty" name:"OperateInform"`

	// 任务删除方式
	// true：不针对下游任务实例进行强制失败
	// false：针对下游任务实例进行强制失败
	// 不传默认false
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	// 和VirtualTaskId选填一个
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务操作是否消息通知下游任务责任人true：通知
	// false：不通知
	// 不传默认false
	OperateInform *bool `json:"OperateInform,omitnil,omitempty" name:"OperateInform"`

	// 任务删除方式
	// true：不针对下游任务实例进行强制失败
	// false：针对下游任务实例进行强制失败
	// 不传默认false
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "OperateInform")
	delete(f, "DeleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// 是否删除成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DeleteTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskResponseParams `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTaskResult struct {
	// 删除状态,true表示成功，false表示失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteWorkflowFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，可通过ListWorkflowFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type DeleteWorkflowFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，可通过ListWorkflowFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *DeleteWorkflowFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowFolderResponseParams struct {
	// 删除结果
	Data *DeleteFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowFolderResponseParams `json:"Response"`
}

func (r *DeleteWorkflowFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DeleteWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DeleteWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowResponseParams struct {
	// 是否删除成功
	Data *DeleteWorkflowResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowResponseParams `json:"Response"`
}

func (r *DeleteWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWorkflowResult struct {
	// 删除工作流是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type DependencyConfigPage struct {
	// 满足查询条件的数据总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 满足查询条件的数据总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 当前请求的数据页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 当前请求的数据页条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TaskDependDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type DependencyStrategyTask struct {
	// 等待上游任务实例策略：EXECUTING（执行）；WAITING（等待）
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	PollingNullStrategy *string `json:"PollingNullStrategy,omitnil,omitempty" name:"PollingNullStrategy"`

	// 仅当PollingNullStrategy为EXECUTING时才需要填本字段，List类型：NOT_EXIST（默认，在分钟依赖分钟/小时依赖小时的情况下，父实例不在下游实例调度时间范围内）；PARENT_EXPIRED（父实例失败）；PARENT_TIMEOUT（父实例超时）。以上场景满足任一条件即可通过该父任务实例依赖判断，除以上场景外均需等待父实例。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskDependencyExecutingStrategies []*string `json:"TaskDependencyExecutingStrategies,omitnil,omitempty" name:"TaskDependencyExecutingStrategies"`

	// 仅当TaskDependencyExecutingStrategies中包含PARENT_TIMEOUT时才需要填本字段，下游任务依赖父实例执行超时时间，单位：分钟。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskDependencyExecutingTimeoutValue *int64 `json:"TaskDependencyExecutingTimeoutValue,omitnil,omitempty" name:"TaskDependencyExecutingTimeoutValue"`
}

type DependencyTaskBrief struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主依赖配置，取值为：
	// 
	// * CRONTAB
	// * DAY
	// * HOUR
	// * LIST_DAY
	// * LIST_HOUR
	// * LIST_MINUTE
	// * MINUTE
	// * MONTH
	// * RANGE_DAY
	// * RANGE_HOUR
	// * RANGE_MINUTE
	// * WEEK
	// * YEAR
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainCyclicConfig *string `json:"MainCyclicConfig,omitnil,omitempty" name:"MainCyclicConfig"`

	// 次依赖配置，取值为：
	// * ALL_DAY_OF_YEAR
	// * ALL_MONTH_OF_YEAR
	// * CURRENT
	// * CURRENT_DAY
	// * CURRENT_HOUR
	// * CURRENT_MINUTE
	// * CURRENT_MONTH
	// * CURRENT_WEEK
	// * CURRENT_YEAR
	// * PREVIOUS_BEGIN_OF_MONTH
	// * PREVIOUS_DAY
	// * PREVIOUS_DAY_LATER_OFFSET_HOUR
	// * PREVIOUS_DAY_LATER_OFFSET_MINUTE
	// * PREVIOUS_END_OF_MONTH
	// * PREVIOUS_FRIDAY
	// * PREVIOUS_HOUR
	// * PREVIOUS_HOUR_CYCLE
	// * PREVIOUS_HOUR_LATER_OFFSET_MINUTE
	// * PREVIOUS_MINUTE_CYCLE
	// * PREVIOUS_MONTH
	// * PREVIOUS_WEEK
	// * PREVIOUS_WEEKEND
	// * RECENT_DATE
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubordinateCyclicConfig *string `json:"SubordinateCyclicConfig,omitnil,omitempty" name:"SubordinateCyclicConfig"`

	// 区间、列表模式下的偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 依赖执行策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyStrategy *DependencyStrategyTask `json:"DependencyStrategy,omitnil,omitempty" name:"DependencyStrategy"`
}

// Predefined struct for user
type DisableProjectRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DisableProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DisableProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableProjectResponseParams struct {
	// 无
	Data *ProjectResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableProjectResponse struct {
	*tchttp.BaseResponse
	Response *DisableProjectResponseParams `json:"Response"`
}

func (r *DisableProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DissociateResourceGroupFromProjectRequestParams struct {
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DissociateResourceGroupFromProjectRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DissociateResourceGroupFromProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DissociateResourceGroupFromProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DissociateResourceGroupFromProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DissociateResourceGroupFromProjectResponseParams struct {
	// 是否绑定成功，失败返回异常
	Data *ResourceStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DissociateResourceGroupFromProjectResponse struct {
	*tchttp.BaseResponse
	Response *DissociateResourceGroupFromProjectResponseParams `json:"Response"`
}

func (r *DissociateResourceGroupFromProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DissociateResourceGroupFromProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableProjectRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type EnableProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *EnableProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableProjectResponseParams struct {
	// 无
	Data *ProjectResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableProjectResponse struct {
	*tchttp.BaseResponse
	Response *EnableProjectResponseParams `json:"Response"`
}

func (r *EnableProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventListener struct {
	// 事件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件周期：SECOND, MIN, HOUR, DAY
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 事件广播类型：SINGLE, BROADCAST
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// 扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PropertiesList []*ParamInfo `json:"PropertiesList,omitnil,omitempty" name:"PropertiesList"`
}

type ExecutorResourceGroupData struct {
	// 结果list
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ExecutorResourceGroupInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type ExecutorResourceGroupInfo struct {
	// 资源组唯一标识
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 执行资源组类型，不能为空
	// 
	// - Schedule --- 调度资源组
	// - Integration --- 集成资源组
	// - DataService -- 数据服务资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupType *string `json:"ResourceGroupType,omitnil,omitempty" name:"ResourceGroupType"`

	// 项目集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateProjects []*BindProject `json:"AssociateProjects,omitnil,omitempty" name:"AssociateProjects"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// vpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubNet *string `json:"SubNet,omitnil,omitempty" name:"SubNet"`

	// 是否自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitnil,omitempty" name:"AutoRenewEnabled"`
}

// Predefined struct for user
type GetAlarmMessageRequestParams struct {
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警消息Id
	AlarmMessageId *string `json:"AlarmMessageId,omitnil,omitempty" name:"AlarmMessageId"`

	// 返回日期的时区, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type GetAlarmMessageRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警消息Id
	AlarmMessageId *string `json:"AlarmMessageId,omitnil,omitempty" name:"AlarmMessageId"`

	// 返回日期的时区, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *GetAlarmMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmMessageId")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAlarmMessageResponseParams struct {
	// 告警信息
	Data *AlarmMessage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAlarmMessageResponse struct {
	*tchttp.BaseResponse
	Response *GetAlarmMessageResponseParams `json:"Response"`
}

func (r *GetAlarmMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCodeFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件ID，参数值来自CreateCodeFile接口的返回
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// true：返回文件内容+配置，false：不返回文件内容，只返回配置信息；默认为false
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`
}

type GetCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件ID，参数值来自CreateCodeFile接口的返回
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// true：返回文件内容+配置，false：不返回文件内容，只返回配置信息；默认为false
	IncludeContent *bool `json:"IncludeContent,omitnil,omitempty" name:"IncludeContent"`
}

func (r *GetCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileId")
	delete(f, "IncludeContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCodeFileResponseParams struct {
	// 代码文件详情
	Data *CodeFile `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *GetCodeFileResponseParams `json:"Response"`
}

func (r *GetCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCodeFolderRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type GetCodeFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *GetCodeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCodeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCodeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCodeFolderResponseParams struct {
	// codestudio文件夹
	Data *CodeFolderNode `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCodeFolderResponse struct {
	*tchttp.BaseResponse
	Response *GetCodeFolderResponseParams `json:"Response"`
}

func (r *GetCodeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCodeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataBackfillPlanRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划id
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`

	// 展示时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type GetDataBackfillPlanRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划id
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`

	// 展示时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *GetDataBackfillPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataBackfillPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DataBackfillPlanId")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataBackfillPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataBackfillPlanResponseParams struct {
	// 补录详情
	Data *DataBackfill `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDataBackfillPlanResponse struct {
	*tchttp.BaseResponse
	Response *GetDataBackfillPlanResponseParams `json:"Response"`
}

func (r *GetDataBackfillPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataBackfillPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataSourceRelatedTasksRequestParams struct {
	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type GetDataSourceRelatedTasksRequest struct {
	*tchttp.BaseRequest
	
	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *GetDataSourceRelatedTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataSourceRelatedTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataSourceRelatedTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataSourceRelatedTasksResponseParams struct {
	// 无
	Data []*DatasourceRelationTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDataSourceRelatedTasksResponse struct {
	*tchttp.BaseResponse
	Response *GetDataSourceRelatedTasksResponseParams `json:"Response"`
}

func (r *GetDataSourceRelatedTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataSourceRelatedTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataSourceRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type GetDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *GetDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataSourceResponseParams struct {
	// 数据源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataSource `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *GetDataSourceResponseParams `json:"Response"`
}

func (r *GetDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAlarmRuleRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则唯一id
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

type GetOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则唯一id
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`
}

func (r *GetOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAlarmRuleResponseParams struct {
	// 告警规则详细信息
	Data *AlarmRuleData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *GetOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAsyncJobRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 异步操作id
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`
}

type GetOpsAsyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 异步操作id
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`
}

func (r *GetOpsAsyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAsyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AsyncId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsAsyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsAsyncJobResponseParams struct {
	// 异步操作详情结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OpsAsyncJobDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsAsyncJobResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsAsyncJobResponseParams `json:"Response"`
}

func (r *GetOpsAsyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsAsyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskCodeRequestParams struct {
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetOpsTaskCodeRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetOpsTaskCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsTaskCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskCodeResponseParams struct {
	// 获取任务代码结果
	Data *TaskCode `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsTaskCodeResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsTaskCodeResponseParams `json:"Response"`
}

func (r *GetOpsTaskCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskRequestParams struct {
	// 任务Id	
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetOpsTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id	
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *GetOpsTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsTaskResponseParams struct {
	// 任务详情
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsTaskResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsTaskResponseParams `json:"Response"`
}

func (r *GetOpsTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsWorkflowRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流Id，可以从ListOpsWorkflows接口获取
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type GetOpsWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流Id，可以从ListOpsWorkflows接口获取
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *GetOpsWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOpsWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOpsWorkflowResponseParams struct {
	// 工作流调度详情
	Data *OpsWorkflowDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOpsWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *GetOpsWorkflowResponseParams `json:"Response"`
}

func (r *GetOpsWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOpsWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProjectRequestParams struct {
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *GetProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProjectResponseParams struct {
	// 项目信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *Project `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetProjectResponse struct {
	*tchttp.BaseResponse
	Response *GetProjectResponseParams `json:"Response"`
}

func (r *GetProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourceFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件ID,可通过ListResourceFiles接口获取
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type GetResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件ID,可通过ListResourceFiles接口获取
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *GetResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourceFileResponseParams struct {
	// 资源文件详情
	Data *ResourceFile `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *GetResourceFileResponseParams `json:"Response"`
}

func (r *GetResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourceGroupMetricsRequestParams struct {
	// 执行资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 使用趋势开始时间(毫秒)，默认最近一小时
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 使用趋势结束时间(毫秒)，默认当前
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标维度
	// 
	// - all --- 全部
	// - task --- 任务指标
	// - system --- 系统指标
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 指标采集粒度，单位分钟，默认 1分钟
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

type GetResourceGroupMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 执行资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 使用趋势开始时间(毫秒)，默认最近一小时
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 使用趋势结束时间(毫秒)，默认当前
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标维度
	// 
	// - all --- 全部
	// - task --- 任务指标
	// - system --- 系统指标
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 指标采集粒度，单位分钟，默认 1分钟
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`
}

func (r *GetResourceGroupMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceGroupMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricType")
	delete(f, "Granularity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetResourceGroupMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetResourceGroupMetricsResponseParams struct {
	// 执行组指标信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ResourceGroupMetrics `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetResourceGroupMetricsResponse struct {
	*tchttp.BaseResponse
	Response *GetResourceGroupMetricsResponseParams `json:"Response"`
}

func (r *GetResourceGroupMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetResourceGroupMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSQLFolderRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type GetSQLFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *GetSQLFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSQLFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSQLFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSQLFolderResponseParams struct {
	// sql文件夹
	Data *SQLFolderNode `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSQLFolderResponse struct {
	*tchttp.BaseResponse
	Response *GetSQLFolderResponseParams `json:"Response"`
}

func (r *GetSQLFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSQLFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSQLScriptRequestParams struct {
	// 探索脚本Id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type GetSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// 探索脚本Id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *GetSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSQLScriptResponseParams struct {
	// 脚本详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SQLScript `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *GetSQLScriptResponseParams `json:"Response"`
}

func (r *GetSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTableColumnsRequestParams struct {
	// 数据表GUID
	TableGuid *string `json:"TableGuid,omitnil,omitempty" name:"TableGuid"`
}

type GetTableColumnsRequest struct {
	*tchttp.BaseRequest
	
	// 数据表GUID
	TableGuid *string `json:"TableGuid,omitnil,omitempty" name:"TableGuid"`
}

func (r *GetTableColumnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTableColumnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableGuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTableColumnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTableColumnsResponseParams struct {
	// 表字段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ColumnInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTableColumnsResponse struct {
	*tchttp.BaseResponse
	Response *GetTableColumnsResponseParams `json:"Response"`
}

func (r *GetTableColumnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTableColumnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTableRequestParams struct {
	// 表GUID
	TableGuid *string `json:"TableGuid,omitnil,omitempty" name:"TableGuid"`
}

type GetTableRequest struct {
	*tchttp.BaseRequest
	
	// 表GUID
	TableGuid *string `json:"TableGuid,omitnil,omitempty" name:"TableGuid"`
}

func (r *GetTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableGuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTableResponseParams struct {
	// 数据表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TableInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTableResponse struct {
	*tchttp.BaseResponse
	Response *GetTableResponseParams `json:"Response"`
}

func (r *GetTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskCodeRequestParams struct {
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetTaskCodeRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetTaskCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskCodeResponseParams struct {
	// 获取任务代码结果
	Data *TaskCodeResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskCodeResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskCodeResponseParams `json:"Response"`
}

func (r *GetTaskCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceLogRequestParams struct {
	// **项目ID**
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **实例生命周期编号，标识实例的某一次执行**例如：周期实例第一次运行的编号为0，用户后期又重跑了该实例，第二次执行的编号为1; 默认最新一次
	LifeRoundNum *uint64 `json:"LifeRoundNum,omitnil,omitempty" name:"LifeRoundNum"`

	// **日志级别** 默认All - Info - Debug - Warn - Error - All
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// **分页查询日志时使用，无具体业务含义** 第一次查询时值为null 第二次及以后查询时使用上一次查询返回信息中的NextCursor字段值即可
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`
}

type GetTaskInstanceLogRequest struct {
	*tchttp.BaseRequest
	
	// **项目ID**
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **实例生命周期编号，标识实例的某一次执行**例如：周期实例第一次运行的编号为0，用户后期又重跑了该实例，第二次执行的编号为1; 默认最新一次
	LifeRoundNum *uint64 `json:"LifeRoundNum,omitnil,omitempty" name:"LifeRoundNum"`

	// **日志级别** 默认All - Info - Debug - Warn - Error - All
	LogLevel *string `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// **分页查询日志时使用，无具体业务含义** 第一次查询时值为null 第二次及以后查询时使用上一次查询返回信息中的NextCursor字段值即可
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`
}

func (r *GetTaskInstanceLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "LifeRoundNum")
	delete(f, "LogLevel")
	delete(f, "NextCursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskInstanceLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceLogResponseParams struct {
	// 调度实例详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *InstanceLog `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskInstanceLogResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskInstanceLogResponseParams `json:"Response"`
}

func (r *GetTaskInstanceLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceRequestParams struct {
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例唯一标识，可以通过ListInstances获取
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区**timeZone, 传入的时间字符串的所在时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type GetTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例唯一标识，可以通过ListInstances获取
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区**timeZone, 传入的时间字符串的所在时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *GetTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskInstanceResponseParams struct {
	// 实例详情
	Data *TaskInstanceDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskInstanceResponseParams `json:"Response"`
}

func (r *GetTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskResponseParams struct {
	// 任务详情
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskResponseParams `json:"Response"`
}

func (r *GetTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskVersionRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 提交版本ID，不填默认拿最新提交版本
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type GetTaskVersionRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 提交版本ID，不填默认拿最新提交版本
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *GetTaskVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskVersionResponseParams struct {
	// 版本详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TaskVersionDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskVersionResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskVersionResponseParams `json:"Response"`
}

func (r *GetTaskVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID 通过ListWorkflows接口获取
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type GetWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID 通过ListWorkflows接口获取
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *GetWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWorkflowResponseParams struct {
	// 工作流详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkflowDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *GetWorkflowResponseParams `json:"Response"`
}

func (r *GetWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantMemberProjectRoleRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 角色id列表，目前支持的项目角色有
	// - 308335260274237440 (项目管理员)
	// - 308335260676890624 (数据工程师)
	// - 308335260844662784 (运维工程师)
	// - 308335260945326080 (普通成员)
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type GrantMemberProjectRoleRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 角色id列表，目前支持的项目角色有
	// - 308335260274237440 (项目管理员)
	// - 308335260676890624 (数据工程师)
	// - 308335260844662784 (运维工程师)
	// - 308335260945326080 (普通成员)
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *GrantMemberProjectRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantMemberProjectRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserUin")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantMemberProjectRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantMemberProjectRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GrantMemberProjectRoleResponse struct {
	*tchttp.BaseResponse
	Response *GrantMemberProjectRoleResponseParams `json:"Response"`
}

func (r *GrantMemberProjectRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantMemberProjectRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InTaskParameter struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数描述：格式为 项目标识.任务名称.参数名；例：project_wedata_1.sh_250820_104107.pp_out
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamDesc *string `json:"ParamDesc,omitnil,omitempty" name:"ParamDesc"`

	// 父任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromTaskId *string `json:"FromTaskId,omitnil,omitempty" name:"FromTaskId"`

	// 父任务参数key
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromParamKey *string `json:"FromParamKey,omitnil,omitempty" name:"FromParamKey"`
}

type InstanceExecution struct {
	// 实例唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **实例生命周期编号，标识实例的某一次执行**
	// 
	// 例如：周期实例第一次运行的编号为0，用户后期又重跑了该实例，第二次执行的编号为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeRoundNum *uint64 `json:"LifeRoundNum,omitnil,omitempty" name:"LifeRoundNum"`

	// **实例状态**
	// - WAIT_EVENT: 等待事件
	// - WAIT_UPSTREAM: 等待上游
	// - WAIT_RUN: 等待运行
	// - RUNNING: 运行中
	// - SKIP_RUNNING: 跳过运行
	// - FAILED_RETRY: 失败重试
	// - EXPIRED: 失败
	// - COMPLETED: 成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// **实例运行触发类型**
	// 
	// - RERUN 表示重跑
	// - ADDITION 表示补录
	// - PERIODIC 表示周期
	// - APERIODIC 表示非周期
	// - RERUN_SKIP_RUN 表示重跑 - 空跑
	// - ADDITION_SKIP_RUN 表示补录 - 空跑
	// - PERIODIC_SKIP_RUN 表示周期 - 空跑
	// - APERIODIC_SKIP_RUN 表示非周期 - 空跑
	// - MANUAL_TRIGGER 表示手动触发
	// - RERUN_MANUAL_TRIGGER 表示手动触发 - 重跑
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunType *string `json:"RunType,omitnil,omitempty" name:"RunType"`

	// 失败重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// **实例执行生命周期列表**
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionPhaseList []*InstanceExecutionPhase `json:"ExecutionPhaseList,omitnil,omitempty" name:"ExecutionPhaseList"`

	// 耗费时间, 单位ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`
}

type InstanceExecutionPhase struct {
	// 该状态开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// **实例生命周期阶段状态**
	// 
	// - WAIT_UPSTREAM 表示 等待事件/上游状态
	// - WAIT_RUN 表示 等待运行状态
	// - RUNNING 表示 运行中状态
	// - COMPLETE 表示 终态-完成
	// - FAILED 表示 终态-失败重试
	// - EXPIRED 表示 终态-失败
	// - SKIP_RUNNING 表示 终态-被上游分支节点跳过的分支
	// - HISTORY 表示 兼容2024-03-30之前的历史实例，之后实例无需关注次枚举类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailState *string `json:"DetailState,omitnil,omitempty" name:"DetailState"`

	// 该状态结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type InstanceLog struct {
	// 实例唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **运行代码内容**
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// **日志内容**
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogInfo *string `json:"LogInfo,omitnil,omitempty" name:"LogInfo"`

	// **分页查询日志时使用，无具体业务含义**
	// 
	// 第一次查询时值为null 
	// 第二次及以后查询时使用上一次查询返回信息中的NextCursor字段值即可
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`
}

type IntegrationResource struct {
	// 实时集成资源组
	// 
	// - i32c(实时数据同步-16C64G)
	RealTimeDataSync *ResourceGroupSpecification `json:"RealTimeDataSync,omitnil,omitempty" name:"RealTimeDataSync"`

	// 离线集成资源组
	// 
	// - integrated(离线数据同步-8C16G)
	// - i16(离线数据同步-8C32G)
	OfflineDataSync *ResourceGroupSpecification `json:"OfflineDataSync,omitnil,omitempty" name:"OfflineDataSync"`
}

type JobDto struct {
	// 数据探索任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 数据探索任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 脚本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 子任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobExecutionList []*JobExecutionDto `json:"JobExecutionList,omitnil,omitempty" name:"JobExecutionList"`

	// 脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 云主账号UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 账号UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeCost *int64 `json:"TimeCost,omitnil,omitempty" name:"TimeCost"`

	// 是否脚本内容被截断
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContentTruncate *bool `json:"ScriptContentTruncate,omitnil,omitempty" name:"ScriptContentTruncate"`
}

type JobExecutionDto struct {
	// 数据探索任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 子查询任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobExecutionId *string `json:"JobExecutionId,omitnil,omitempty" name:"JobExecutionId"`

	// 子查询名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobExecutionName *string `json:"JobExecutionName,omitnil,omitempty" name:"JobExecutionName"`

	// 子查询sql内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 子查询状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 执行阶段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteStageInfo *string `json:"ExecuteStageInfo,omitnil,omitempty" name:"ExecuteStageInfo"`

	// 日志路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogFilePath *string `json:"LogFilePath,omitnil,omitempty" name:"LogFilePath"`

	// 下载结果路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultFilePath *string `json:"ResultFilePath,omitnil,omitempty" name:"ResultFilePath"`

	// 预览结果路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultPreviewFilePath *string `json:"ResultPreviewFilePath,omitnil,omitempty" name:"ResultPreviewFilePath"`

	// 任务执行的结果总行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultTotalCount *int64 `json:"ResultTotalCount,omitnil,omitempty" name:"ResultTotalCount"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeCost *int64 `json:"TimeCost,omitnil,omitempty" name:"TimeCost"`

	// 上下文SQL内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextScriptContent []*string `json:"ContextScriptContent,omitnil,omitempty" name:"ContextScriptContent"`

	// 任务执行的结果预览行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultPreviewCount *int64 `json:"ResultPreviewCount,omitnil,omitempty" name:"ResultPreviewCount"`

	// 任务执行的结果影响行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultEffectCount *int64 `json:"ResultEffectCount,omitnil,omitempty" name:"ResultEffectCount"`

	// 是否正在收集全量结果：默认false，true表示正在收集全量结果，用于前端判断是否需要继续轮询
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectingTotalResult *bool `json:"CollectingTotalResult,omitnil,omitempty" name:"CollectingTotalResult"`

	// 是否需要截断脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContentTruncate *bool `json:"ScriptContentTruncate,omitnil,omitempty" name:"ScriptContentTruncate"`
}

type KVMap struct {
	// k
	// 注意：此字段可能返回 null，表示取不到有效值。
	K *string `json:"K,omitnil,omitempty" name:"K"`

	// v
	// 注意：此字段可能返回 null，表示取不到有效值。
	V *string `json:"V,omitnil,omitempty" name:"V"`
}

type KVPair struct {
	// 键名
	// 注意：此字段可能返回 null，表示取不到有效值。
	K *string `json:"K,omitnil,omitempty" name:"K"`

	// 值，请勿传SQL(请求会被视为攻击接口)，如果有需要，请将SQL进行Base64转码并解码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	V *string `json:"V,omitnil,omitempty" name:"V"`
}

// Predefined struct for user
type KillTaskInstancesAsyncRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例id列表,可以从ListInstances中获取
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

type KillTaskInstancesAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例id列表,可以从ListInstances中获取
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

func (r *KillTaskInstancesAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillTaskInstancesAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKeyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillTaskInstancesAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillTaskInstancesAsyncResponseParams struct {
	// 批量中止操作的返回的异步id, 可以在接口GetAsyncJob获取具体执行详情
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillTaskInstancesAsyncResponse struct {
	*tchttp.BaseResponse
	Response *KillTaskInstancesAsyncResponseParams `json:"Response"`
}

func (r *KillTaskInstancesAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillTaskInstancesAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LineageNodeInfo struct {
	// 当前资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *LineageResource `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *LineageRelation `json:"Relation,omitnil,omitempty" name:"Relation"`
}

type LineagePair struct {
	// 来源
	Source *LineageResouce `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标
	Target *LineageResouce `json:"Target,omitnil,omitempty" name:"Target"`

	// 血缘加工过程
	Processes []*LineageProcess `json:"Processes,omitnil,omitempty" name:"Processes"`
}

type LineageProcess struct {
	// 原始唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessId *string `json:"ProcessId,omitnil,omitempty" name:"ProcessId"`

	// 任务类型
	//     //调度任务
	//     SCHEDULE_TASK,
	//     //集成任务
	//     INTEGRATION_TASK,
	//     //第三方上报
	//     THIRD_REPORT,
	//     //数据建模
	//     TABLE_MODEL,
	//     //模型创建指标
	//     MODEL_METRIC,
	//     //原子指标创建衍生指标
	//     METRIC_METRIC,
	//     //数据服务
	//     DATA_SERVICE
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// WEDATA, THIRD;
	// 注意：此字段可能返回 null，表示取不到有效值。
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 任务子类型
	//  SQL_TASK,
	//     //集成实时任务血缘
	//     INTEGRATED_STREAM,
	//     //集成离线任务血缘
	//     INTEGRATED_OFFLINE;
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessSubType *string `json:"ProcessSubType,omitnil,omitempty" name:"ProcessSubType"`

	// 额外扩展参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessProperties []*LineageProperty `json:"ProcessProperties,omitnil,omitempty" name:"ProcessProperties"`

	// 血缘任务唯一节点ID
	LineageNodeId *string `json:"LineageNodeId,omitnil,omitempty" name:"LineageNodeId"`
}

type LineageProperty struct {
	// 属性名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LineageRelation struct {
	// 关联ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationId *string `json:"RelationId,omitnil,omitempty" name:"RelationId"`

	// 源端唯一血缘ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceUniqueId *string `json:"SourceUniqueId,omitnil,omitempty" name:"SourceUniqueId"`

	// 目标端唯一血缘ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetUniqueId *string `json:"TargetUniqueId,omitnil,omitempty" name:"TargetUniqueId"`

	// 血缘加工过程
	// 注意：此字段可能返回 null，表示取不到有效值。
	Processes []*LineageProcess `json:"Processes,omitnil,omitempty" name:"Processes"`
}

type LineageResouce struct {
	// 实体原始唯一ID\n
	// 备注：当血缘为表的列时候 唯一ID传表ResourceUniqueId::字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceUniqueId *string `json:"ResourceUniqueId,omitnil,omitempty" name:"ResourceUniqueId"`

	// 实体类型
	// TABLE|METRIC|MODEL|SERVICE|COLUMN
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 来源：WEDATA|THIRD
	// 默认wedata
	// 注意：此字段可能返回 null，表示取不到有效值。
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 业务名称：库名.表名｜指标名称｜模型名称|字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 描述：表类型｜指标描述｜模型描述|字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// resource 额外扩展参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceProperties []*LineageProperty `json:"ResourceProperties,omitnil,omitempty" name:"ResourceProperties"`

	// 血缘节点唯一标识符号
	LineageNodeId *string `json:"LineageNodeId,omitnil,omitempty" name:"LineageNodeId"`
}

type LineageResource struct {
	// 实体原始唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceUniqueId *string `json:"ResourceUniqueId,omitnil,omitempty" name:"ResourceUniqueId"`

	// 业务名称：库名.表名｜指标名称｜模型名称|字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 实体类型
	// TABLE|METRIC|MODEL|SERVICE|COLUMN
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 血缘节点唯一标识符号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LineageNodeId *string `json:"LineageNodeId,omitnil,omitempty" name:"LineageNodeId"`

	// 描述：表类型｜指标描述｜模型描述|字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 来源：WEDATA|THIRD
	// 默认wedata
	// 注意：此字段可能返回 null，表示取不到有效值。
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// resource 额外扩展参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceProperties []*LineageProperty `json:"ResourceProperties,omitnil,omitempty" name:"ResourceProperties"`
}

type ListAlarmMessages struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 告警信息列表
	Items []*AlarmMessage `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ListAlarmMessagesRequestParams struct {
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码数，用于翻页，最小值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的条数，最大 100 条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 起始告警时间, 格式为yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止告警时间, 格式yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 告警级别
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警接收人Id
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// 对于传入和返回的过滤时区, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type ListAlarmMessagesRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码数，用于翻页，最小值为 1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的条数，最大 100 条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 起始告警时间, 格式为yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止告警时间, 格式yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 告警级别
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警接收人Id
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// 对于传入和返回的过滤时区, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *ListAlarmMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlarmMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AlarmLevel")
	delete(f, "AlarmRecipientId")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAlarmMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAlarmMessagesResponseParams struct {
	// 告警信息列表
	Data *ListAlarmMessages `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAlarmMessagesResponse struct {
	*tchttp.BaseResponse
	Response *ListAlarmMessagesResponseParams `json:"Response"`
}

func (r *ListAlarmMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlarmMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAlarmRulesResult struct {
	// 分页的页数，当前页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的条数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页总页数
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 所有的告警规则个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 告警规则信息列表
	Items []*AlarmRuleData `json:"Items,omitnil,omitempty" name:"Items"`
}

type ListCatalogPage struct {
	// 目录记录列表
	Items []*CatalogInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页总页数
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListCatalogRequestParams struct {
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 父目录ID
	ParentCatalogId *string `json:"ParentCatalogId,omitnil,omitempty" name:"ParentCatalogId"`
}

type ListCatalogRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 父目录ID
	ParentCatalogId *string `json:"ParentCatalogId,omitnil,omitempty" name:"ParentCatalogId"`
}

func (r *ListCatalogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCatalogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ParentCatalogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCatalogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCatalogResponseParams struct {
	// 分页数据
	Data *ListCatalogPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListCatalogResponse struct {
	*tchttp.BaseResponse
	Response *ListCatalogResponseParams `json:"Response"`
}

func (r *ListCatalogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCatalogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCodeFolderContentsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，例如/aaa/bbb/ccc，路径头需带斜杠，根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 文件夹名称/代码文件名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 只查询文件夹
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// 是否只查询用户自己创建的代码文件
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`
}

type ListCodeFolderContentsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，例如/aaa/bbb/ccc，路径头需带斜杠，根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 文件夹名称/代码文件名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 只查询文件夹
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// 是否只查询用户自己创建的代码文件
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`
}

func (r *ListCodeFolderContentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCodeFolderContentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "Keyword")
	delete(f, "OnlyFolderNode")
	delete(f, "OnlyUserSelfScript")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCodeFolderContentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCodeFolderContentsResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CodeFolderNode `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListCodeFolderContentsResponse struct {
	*tchttp.BaseResponse
	Response *ListCodeFolderContentsResponseParams `json:"Response"`
}

func (r *ListCodeFolderContentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCodeFolderContentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListColumnLineageRequestParams struct {
	// 表唯一ID
	TableUniqueId *string `json:"TableUniqueId,omitnil,omitempty" name:"TableUniqueId"`

	// 血缘方向 INPUT｜OUTPUT
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 列名称
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 来源：WEDATA|THIRD 默认WEDATA
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type ListColumnLineageRequest struct {
	*tchttp.BaseRequest
	
	// 表唯一ID
	TableUniqueId *string `json:"TableUniqueId,omitnil,omitempty" name:"TableUniqueId"`

	// 血缘方向 INPUT｜OUTPUT
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 列名称
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 来源：WEDATA|THIRD 默认WEDATA
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *ListColumnLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListColumnLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableUniqueId")
	delete(f, "Direction")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ColumnName")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListColumnLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListColumnLineageResponseParams struct {
	// 分页数据
	Data *ListLineagePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListColumnLineageResponse struct {
	*tchttp.BaseResponse
	Response *ListColumnLineageResponseParams `json:"Response"`
}

func (r *ListColumnLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListColumnLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataBackfillInstancesRequestParams struct {
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划Id
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDataBackfillInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划Id
	DataBackfillPlanId *string `json:"DataBackfillPlanId,omitnil,omitempty" name:"DataBackfillPlanId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDataBackfillInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataBackfillInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DataBackfillPlanId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDataBackfillInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataBackfillInstancesResponseParams struct {
	// 单个补录计划下的所有补录实例
	Data *BackfillInstanceCollection `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDataBackfillInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListDataBackfillInstancesResponseParams `json:"Response"`
}

func (r *ListDataBackfillInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataBackfillInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataSourcesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源展示名
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据源类型:枚举值
	// 
	// - MYSQL
	// - TENCENT_MYSQL
	// - POSTGRE
	// - ORACLE
	// - SQLSERVER
	// - FTP
	// - HIVE
	// - HUDI
	// - HDFS
	// - ICEBERG
	// - KAFKA
	// - HBASE
	// - SPARK
	// - VIRTUAL
	// - TBASE
	// - DB2
	// - DM
	// - GAUSSDB
	// - GBASE
	// - IMPALA
	// - ES
	// - TENCENT_ES
	// - GREENPLUM
	// - PHOENIX
	// - SAP_HANA
	// - SFTP
	// - OCEANBASE
	// - CLICKHOUSE
	// - KUDU
	// - VERTICA
	// - REDIS
	// - COS
	// - DLC
	// - DORIS
	// - CKAFKA
	// - S3
	// - TDSQL
	// - TDSQL_MYSQL
	// - MONGODB
	// - TENCENT_MONGODB
	// - REST_API
	// - SuperSQL
	// - PRESTO
	// - TiDB
	// - StarRocks
	// - Trino
	// - Kyuubi
	// - TCHOUSE_X
	// - TCHOUSE_P
	// - TCHOUSE_C
	// - TCHOUSE_D
	// - INFLUXDB
	// - BIG_QUERY
	// - SSH
	// - BLOB
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

type ListDataSourcesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 返回数量，默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 数据源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源展示名
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据源类型:枚举值
	// 
	// - MYSQL
	// - TENCENT_MYSQL
	// - POSTGRE
	// - ORACLE
	// - SQLSERVER
	// - FTP
	// - HIVE
	// - HUDI
	// - HDFS
	// - ICEBERG
	// - KAFKA
	// - HBASE
	// - SPARK
	// - VIRTUAL
	// - TBASE
	// - DB2
	// - DM
	// - GAUSSDB
	// - GBASE
	// - IMPALA
	// - ES
	// - TENCENT_ES
	// - GREENPLUM
	// - PHOENIX
	// - SAP_HANA
	// - SFTP
	// - OCEANBASE
	// - CLICKHOUSE
	// - KUDU
	// - VERTICA
	// - REDIS
	// - COS
	// - DLC
	// - DORIS
	// - CKAFKA
	// - S3
	// - TDSQL
	// - TDSQL_MYSQL
	// - MONGODB
	// - TENCENT_MONGODB
	// - REST_API
	// - SuperSQL
	// - PRESTO
	// - TiDB
	// - StarRocks
	// - Trino
	// - Kyuubi
	// - TCHOUSE_X
	// - TCHOUSE_P
	// - TCHOUSE_C
	// - TCHOUSE_D
	// - INFLUXDB
	// - BIG_QUERY
	// - SSH
	// - BLOB
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

func (r *ListDataSourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataSourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Name")
	delete(f, "DisplayName")
	delete(f, "Type")
	delete(f, "Creator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDataSourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDataSourcesResponseParams struct {
	// 数据源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataSourceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDataSourcesResponse struct {
	*tchttp.BaseResponse
	Response *ListDataSourcesResponseParams `json:"Response"`
}

func (r *ListDataSourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDatabasePage struct {
	// 数据库记录列表
	Items []*DatabaseInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页总页数
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListDatabaseRequestParams struct {
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 目录名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据源ID
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 目录名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据源ID
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "CatalogName")
	delete(f, "DatasourceId")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDatabaseResponseParams struct {
	// 分页数据
	Data *ListDatabasePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *ListDatabaseResponseParams `json:"Response"`
}

func (r *ListDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamOpsTasksRequestParams struct {
	// 任务Id	
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id	
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页页码
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDownstreamOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id	
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id	
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页页码
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDownstreamOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDownstreamOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamOpsTasksResponseParams struct {
	// 下游依赖详情
	Data *ChildDependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDownstreamOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListDownstreamOpsTasksResponseParams `json:"Response"`
}

func (r *ListDownstreamOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTaskInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区** timeZone, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **页码，整型**配合pageSize使用且不能小于1， 默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// **每页显示的条数，默认为10，最小值为1、最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDownstreamTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区** timeZone, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **页码，整型**配合pageSize使用且不能小于1， 默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// **每页显示的条数，默认为10，最小值为1、最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDownstreamTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDownstreamTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTaskInstancesResponseParams struct {
	// 直接下游实例列表
	Data *TaskInstancePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDownstreamTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListDownstreamTaskInstancesResponseParams `json:"Response"`
}

func (r *ListDownstreamTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTasksRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页大小
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页页码
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDownstreamTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页大小
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页页码
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDownstreamTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDownstreamTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDownstreamTasksResponseParams struct {
	// 下游依赖详情
	Data *DependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDownstreamTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListDownstreamTasksResponseParams `json:"Response"`
}

func (r *ListDownstreamTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDownstreamTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLineagePage struct {
	// 血缘记录列表
	Items []*LineageNodeInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页总页数
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListLineageRequestParams struct {
	// 实体唯一ID
	ResourceUniqueId *string `json:"ResourceUniqueId,omitnil,omitempty" name:"ResourceUniqueId"`

	// 实体类型 TABLE|METRIC|MODEL|SERVICE|COLUMN
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 血缘方向 INPUT｜OUTPUT
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 来源：WEDATA|THIRD 默认WEDATA
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type ListLineageRequest struct {
	*tchttp.BaseRequest
	
	// 实体唯一ID
	ResourceUniqueId *string `json:"ResourceUniqueId,omitnil,omitempty" name:"ResourceUniqueId"`

	// 实体类型 TABLE|METRIC|MODEL|SERVICE|COLUMN
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 血缘方向 INPUT｜OUTPUT
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 来源：WEDATA|THIRD 默认WEDATA
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *ListLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceUniqueId")
	delete(f, "ResourceType")
	delete(f, "Direction")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLineageResponseParams struct {
	// 分页数据
	Data *ListLineagePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListLineageResponse struct {
	*tchttp.BaseResponse
	Response *ListLineageResponseParams `json:"Response"`
}

func (r *ListLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsAlarmRulesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页的页数，默认为1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的条数，默认为20，最小值为1、最大值为200
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 监控对象类型, 任务维度监控： 可按照任务/工作流/项目来配置：1.任务、2.工作流、3.项目（默认为1.任务） 项目维度监控： 项目整体任务波动告警， 7：项目波动监控告警
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 根据任务id查询告警规则
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询配置对应告警类型的告警规则
	// 告警规则监控类型 failure：失败告警 ；overtime：超时告警 success：成功告警; backTrackingOrRerunSuccess: 补录重跑成功告警 backTrackingOrRerunFailure：补录重跑失败告警；
	// 项目波动告警
	// projectFailureInstanceUpwardFluctuationAlarm： 当天失败实例数向上波动率超过阀值告警； projectSuccessInstanceDownwardFluctuationAlarm：当天成功实例数向下波动率超过阀值告警；
	// 离线集成任务对账告警： reconciliationFailure： 离线对账任务失败告警 reconciliationOvertime： 离线对账任务运行超时告警 reconciliationMismatch： 数据对账任务不一致条数超过阀值告警
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// 查询配置了对应告警级别的告警规则
	// 告警级别 1.普通、2.重要、3.紧急
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 查询配置对应告警接收人的告警规则
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// 根据告警规则id/规则名称查询对应的告警规则
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 告警规则创建人过滤
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 告警规则创建时间范围起始时间, 格式如2025-08-17 00:00:00
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// 告警规则创建时间范围结束时间，格式如"2025-08-26 23:59:59"
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`

	// 最后更新时间过滤告警规则， 格式如"2025-08-26 00:00:00"
	UpdateTimeFrom *string `json:"UpdateTimeFrom,omitnil,omitempty" name:"UpdateTimeFrom"`

	// 最后更新时间过滤告警规则 格式如： "2025-08-26 23:59:59"
	UpdateTimeTo *string `json:"UpdateTimeTo,omitnil,omitempty" name:"UpdateTimeTo"`
}

type ListOpsAlarmRulesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页的页数，默认为1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的条数，默认为20，最小值为1、最大值为200
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 监控对象类型, 任务维度监控： 可按照任务/工作流/项目来配置：1.任务、2.工作流、3.项目（默认为1.任务） 项目维度监控： 项目整体任务波动告警， 7：项目波动监控告警
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 根据任务id查询告警规则
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询配置对应告警类型的告警规则
	// 告警规则监控类型 failure：失败告警 ；overtime：超时告警 success：成功告警; backTrackingOrRerunSuccess: 补录重跑成功告警 backTrackingOrRerunFailure：补录重跑失败告警；
	// 项目波动告警
	// projectFailureInstanceUpwardFluctuationAlarm： 当天失败实例数向上波动率超过阀值告警； projectSuccessInstanceDownwardFluctuationAlarm：当天成功实例数向下波动率超过阀值告警；
	// 离线集成任务对账告警： reconciliationFailure： 离线对账任务失败告警 reconciliationOvertime： 离线对账任务运行超时告警 reconciliationMismatch： 数据对账任务不一致条数超过阀值告警
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// 查询配置了对应告警级别的告警规则
	// 告警级别 1.普通、2.重要、3.紧急
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 查询配置对应告警接收人的告警规则
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// 根据告警规则id/规则名称查询对应的告警规则
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 告警规则创建人过滤
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 告警规则创建时间范围起始时间, 格式如2025-08-17 00:00:00
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// 告警规则创建时间范围结束时间，格式如"2025-08-26 23:59:59"
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`

	// 最后更新时间过滤告警规则， 格式如"2025-08-26 00:00:00"
	UpdateTimeFrom *string `json:"UpdateTimeFrom,omitnil,omitempty" name:"UpdateTimeFrom"`

	// 最后更新时间过滤告警规则 格式如： "2025-08-26 23:59:59"
	UpdateTimeTo *string `json:"UpdateTimeTo,omitnil,omitempty" name:"UpdateTimeTo"`
}

func (r *ListOpsAlarmRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsAlarmRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "MonitorObjectType")
	delete(f, "TaskId")
	delete(f, "AlarmType")
	delete(f, "AlarmLevel")
	delete(f, "AlarmRecipientId")
	delete(f, "Keyword")
	delete(f, "CreateUserUin")
	delete(f, "CreateTimeFrom")
	delete(f, "CreateTimeTo")
	delete(f, "UpdateTimeFrom")
	delete(f, "UpdateTimeTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOpsAlarmRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsAlarmRulesResponseParams struct {
	// 告警信息信息响应
	Data *ListAlarmRulesResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOpsAlarmRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListOpsAlarmRulesResponseParams `json:"Response"`
}

func (r *ListOpsAlarmRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsAlarmRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListOpsTasksPage struct {
	// 结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 记录列表	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TaskOpsInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

// Predefined struct for user
type ListOpsTasksRequestParams struct {
	// 项目Id	
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 任务类型Id 
	//  - 20：通用数据同步
	//  - 25：ETLTaskType
	//  - 26：ETLTaskType
	//  - 30：python
	//  - 31：pyspark
	//  - 34：HiveSQLTaskType
	//  - 35：shell
	//  - 36：SparkSQLTaskType
	//  - 21：JDBCSQLTaskType
	//  - 32：DLCTaskType
	//  - 33：ImpalaTaskType
	//  - 40：CDWTaskType
	//  - 41：kettle
	//  - 46：DLCSparkTaskType
	//  - 47：TiOne机器学习
	//  - 48：TrinoTaskType
	//  - 50：DLCPyspark39：spark
	//  - 92：mr
	//  - 38：shell脚本
	//  - 70：hivesql脚本
	//  - 1000：自定义业务通用
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 工作流Id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 责任人id
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 数据源id
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 目标数据源id
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 任务周期类型
	// * ONEOFF_CYCLE: 一次性
	// * YEAR_CYCLE: 年
	// * MONTH_CYCLE: 月
	// * WEEK_CYCLE: 周
	// * DAY_CYCLE: 天
	// * HOUR_CYCLE: 小时
	// * MINUTE_CYCLE: 分钟
	// * CRONTAB_CYCLE: crontab表达式类型
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 任务状态:
	// - Y: 运行
	// - F: 停止
	// - O: 冻结
	// - T: 停止中
	// - INVALID: 已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 时区, 默认默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type ListOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id	
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 任务类型Id 
	//  - 20：通用数据同步
	//  - 25：ETLTaskType
	//  - 26：ETLTaskType
	//  - 30：python
	//  - 31：pyspark
	//  - 34：HiveSQLTaskType
	//  - 35：shell
	//  - 36：SparkSQLTaskType
	//  - 21：JDBCSQLTaskType
	//  - 32：DLCTaskType
	//  - 33：ImpalaTaskType
	//  - 40：CDWTaskType
	//  - 41：kettle
	//  - 46：DLCSparkTaskType
	//  - 47：TiOne机器学习
	//  - 48：TrinoTaskType
	//  - 50：DLCPyspark39：spark
	//  - 92：mr
	//  - 38：shell脚本
	//  - 70：hivesql脚本
	//  - 1000：自定义业务通用
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 工作流Id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 责任人id
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 数据源id
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 目标数据源id
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 任务周期类型
	// * ONEOFF_CYCLE: 一次性
	// * YEAR_CYCLE: 年
	// * MONTH_CYCLE: 月
	// * WEEK_CYCLE: 周
	// * DAY_CYCLE: 天
	// * HOUR_CYCLE: 小时
	// * MINUTE_CYCLE: 分钟
	// * CRONTAB_CYCLE: crontab表达式类型
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 任务状态:
	// - Y: 运行
	// - F: 停止
	// - O: 冻结
	// - T: 停止中
	// - INVALID: 已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 时区, 默认默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

func (r *ListOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "TaskTypeId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowName")
	delete(f, "OwnerUin")
	delete(f, "FolderId")
	delete(f, "SourceServiceId")
	delete(f, "TargetServiceId")
	delete(f, "ExecutorGroupId")
	delete(f, "CycleType")
	delete(f, "Status")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsTasksResponseParams struct {
	// 任务列表
	Data *ListOpsTasksPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListOpsTasksResponseParams `json:"Response"`
}

func (r *ListOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsWorkflowsRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 文件Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流状态筛选
	// * ALL_RUNNING : 全部调度中
	// * ALL_FREEZED : 全部已暂停
	// * ALL_STOPPTED : 全部已下线
	// * PART_RUNNING : 部分调度中
	// * ALL_NO_RUNNING : 全部未调度
	// * ALL_INVALID : 全部已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 负责人Id
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 工作流类型筛选, 支持值 Cycle或Manual. 默认只查询 Cycle
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 工作流关键词过滤，支持工作流 Id/name 模糊匹配
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 排序项，可选CreateTime、TaskCount
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序方式，DESC或ASC, 大写
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 创建人Id
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 更新时间，格式yyyy-MM-dd HH:mm:ss
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间，格式yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ListOpsWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 文件Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流状态筛选
	// * ALL_RUNNING : 全部调度中
	// * ALL_FREEZED : 全部已暂停
	// * ALL_STOPPTED : 全部已下线
	// * PART_RUNNING : 部分调度中
	// * ALL_NO_RUNNING : 全部未调度
	// * ALL_INVALID : 全部已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 负责人Id
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 工作流类型筛选, 支持值 Cycle或Manual. 默认只查询 Cycle
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 工作流关键词过滤，支持工作流 Id/name 模糊匹配
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 排序项，可选CreateTime、TaskCount
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序方式，DESC或ASC, 大写
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 创建人Id
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 更新时间，格式yyyy-MM-dd HH:mm:ss
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间，格式yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

func (r *ListOpsWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "FolderId")
	delete(f, "Status")
	delete(f, "OwnerUin")
	delete(f, "WorkflowType")
	delete(f, "KeyWord")
	delete(f, "SortItem")
	delete(f, "SortType")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTime")
	delete(f, "CreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOpsWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOpsWorkflowsResponseParams struct {
	// 工作流列表
	Data *OpsWorkflows `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOpsWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *ListOpsWorkflowsResponseParams `json:"Response"`
}

func (r *ListOpsWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOpsWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListProcessLineagePage struct {
	// 血缘pair列表
	Items []*LineagePair `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页总页数
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListProcessLineageRequestParams struct {
	// 任务唯一ID
	ProcessId *string `json:"ProcessId,omitnil,omitempty" name:"ProcessId"`

	// 任务类型    //调度任务     SCHEDULE_TASK,     //集成任务     INTEGRATION_TASK,     //第三方上报     THIRD_REPORT,     //数据建模     TABLE_MODEL,     //模型创建指标     MODEL_METRIC,     //原子指标创建衍生指标     METRIC_METRIC,     //数据服务     DATA_SERVICE
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 来源：WEDATA|THIRD 默认WEDATA
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type ListProcessLineageRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一ID
	ProcessId *string `json:"ProcessId,omitnil,omitempty" name:"ProcessId"`

	// 任务类型    //调度任务     SCHEDULE_TASK,     //集成任务     INTEGRATION_TASK,     //第三方上报     THIRD_REPORT,     //数据建模     TABLE_MODEL,     //模型创建指标     MODEL_METRIC,     //原子指标创建衍生指标     METRIC_METRIC,     //数据服务     DATA_SERVICE
	ProcessType *string `json:"ProcessType,omitnil,omitempty" name:"ProcessType"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 来源：WEDATA|THIRD 默认WEDATA
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *ListProcessLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProcessLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProcessId")
	delete(f, "ProcessType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListProcessLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListProcessLineageResponseParams struct {
	// 分页数据
	Data *ListProcessLineagePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListProcessLineageResponse struct {
	*tchttp.BaseResponse
	Response *ListProcessLineageResponseParams `json:"Response"`
}

func (r *ListProcessLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProcessLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListProjectMembersRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 使用成员名过滤，支持模糊查询
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 使用成员id过滤，支持模糊查询
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 分页大小，默认第一页
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页条数，默认10条
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type ListProjectMembersRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 使用成员名过滤，支持模糊查询
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 使用成员id过滤，支持模糊查询
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 分页大小，默认第一页
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页条数，默认10条
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *ListProjectMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProjectMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserName")
	delete(f, "UserUin")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListProjectMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListProjectMembersResponseParams struct {
	// 项目列表
	Data *ProjectUsersBrief `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListProjectMembersResponse struct {
	*tchttp.BaseResponse
	Response *ListProjectMembersResponseParams `json:"Response"`
}

func (r *ListProjectMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProjectMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListProjectRolesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 角色中文展示名模糊查询，只能传一个值
	RoleDisplayName *string `json:"RoleDisplayName,omitnil,omitempty" name:"RoleDisplayName"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页信息
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListProjectRolesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 角色中文展示名模糊查询，只能传一个值
	RoleDisplayName *string `json:"RoleDisplayName,omitnil,omitempty" name:"RoleDisplayName"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页信息
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListProjectRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProjectRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RoleDisplayName")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListProjectRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListProjectRolesResponseParams struct {
	// 角色列表
	Data *PageRoles `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListProjectRolesResponse struct {
	*tchttp.BaseResponse
	Response *ListProjectRolesResponseParams `json:"Response"`
}

func (r *ListProjectRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProjectRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListProjectsRequestParams struct {
	// 项目id列表
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 项目名或项目唯一标识名，支持模糊搜索
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目状态，可选值：0（禁用）、1（正常）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目模式，可选值：SIMPLE、STANDARD
	ProjectModel *string `json:"ProjectModel,omitnil,omitempty" name:"ProjectModel"`

	// 请求的数据页数，用于翻页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的条数，默认为 10 条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListProjectsRequest struct {
	*tchttp.BaseRequest
	
	// 项目id列表
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 项目名或项目唯一标识名，支持模糊搜索
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目状态，可选值：0（禁用）、1（正常）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目模式，可选值：SIMPLE、STANDARD
	ProjectModel *string `json:"ProjectModel,omitnil,omitempty" name:"ProjectModel"`

	// 请求的数据页数，用于翻页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的条数，默认为 10 条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectIds")
	delete(f, "ProjectName")
	delete(f, "Status")
	delete(f, "ProjectModel")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListProjectsResponseParams struct {
	// 项目列表
	Data *ProjectBrief `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListProjectsResponse struct {
	*tchttp.BaseResponse
	Response *ListProjectsResponseParams `json:"Response"`
}

func (r *ListProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFilesRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据页数，大于等于1。默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数，最小为10条，最大为200 条。默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 资源文件名称(模糊搜索关键词)
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源文件所属文件夹路径(如/a/b/c，查询c文件夹下的资源文件)
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 创建人ID, 可通过DescribeCurrentUserInfo接口获取
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 更新时间范围,开始时间, 格式yyyy-MM-dd HH:mm:ss
	ModifyTimeStart *string `json:"ModifyTimeStart,omitnil,omitempty" name:"ModifyTimeStart"`

	// 更新时间范围,结束时间, 格式yyyy-MM-dd HH:mm:ss
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitnil,omitempty" name:"ModifyTimeEnd"`

	// 创建时间范围,开始时间, 格式yyyy-MM-dd HH:mm:ss
	CreateTimeStart *string `json:"CreateTimeStart,omitnil,omitempty" name:"CreateTimeStart"`

	// 创建时间范围,结束时间, 格式yyyy-MM-dd HH:mm:ss
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`
}

type ListResourceFilesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据页数，大于等于1。默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数，最小为10条，最大为200 条。默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 资源文件名称(模糊搜索关键词)
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源文件所属文件夹路径(如/a/b/c，查询c文件夹下的资源文件)
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 创建人ID, 可通过DescribeCurrentUserInfo接口获取
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 更新时间范围,开始时间, 格式yyyy-MM-dd HH:mm:ss
	ModifyTimeStart *string `json:"ModifyTimeStart,omitnil,omitempty" name:"ModifyTimeStart"`

	// 更新时间范围,结束时间, 格式yyyy-MM-dd HH:mm:ss
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitnil,omitempty" name:"ModifyTimeEnd"`

	// 创建时间范围,开始时间, 格式yyyy-MM-dd HH:mm:ss
	CreateTimeStart *string `json:"CreateTimeStart,omitnil,omitempty" name:"CreateTimeStart"`

	// 创建时间范围,结束时间, 格式yyyy-MM-dd HH:mm:ss
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`
}

func (r *ListResourceFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ResourceName")
	delete(f, "ParentFolderPath")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTimeStart")
	delete(f, "ModifyTimeEnd")
	delete(f, "CreateTimeStart")
	delete(f, "CreateTimeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListResourceFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFilesResponseParams struct {
	// 获取资源文件列表
	Data *ResourceFilePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListResourceFilesResponse struct {
	*tchttp.BaseResponse
	Response *ListResourceFilesResponseParams `json:"Response"`
}

func (r *ListResourceFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFoldersRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件夹绝对路径，取值示例
	// /wedata/test
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 数据页数，大于等于1。默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数，最小为10条，最大为200 条。默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListResourceFoldersRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件夹绝对路径，取值示例
	// /wedata/test
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 数据页数，大于等于1。默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数，最小为10条，最大为200 条。默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListResourceFoldersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFoldersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListResourceFoldersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceFoldersResponseParams struct {
	// 分页的资源文件夹查询结果
	Data *ResourceFolderPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListResourceFoldersResponse struct {
	*tchttp.BaseResponse
	Response *ListResourceFoldersResponseParams `json:"Response"`
}

func (r *ListResourceFoldersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceGroupsRequestParams struct {
	// 执行资源组类型
	// 
	// - Schedule --- 调度资源组
	// - Integration --- 集成资源组
	// - DataService -- 数据服务资源组
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 资源组id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 搜索的执行资源组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目空间id查询列表
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 执行资源组类型
	// 
	// - Schedule --- 调度资源组
	// - Integration --- 集成资源组
	// - DataService -- 数据服务资源组
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 资源组id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 搜索的执行资源组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目空间id查询列表
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 页数
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "ProjectIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceGroupsResponseParams struct {
	// 分页结果
	Data *ExecutorResourceGroupData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListResourceGroupsResponseParams `json:"Response"`
}

func (r *ListResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLFolderContentsRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，/aaa/bbb/ccc，路径头需带斜杠，查询根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 文件夹名称/脚本名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 只查询文件夹
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// 是否只查询用户自己创建的脚本
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type ListSQLFolderContentsRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹path，/aaa/bbb/ccc，路径头需带斜杠，查询根目录传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 文件夹名称/脚本名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 只查询文件夹
	OnlyFolderNode *bool `json:"OnlyFolderNode,omitnil,omitempty" name:"OnlyFolderNode"`

	// 是否只查询用户自己创建的脚本
	OnlyUserSelfScript *bool `json:"OnlyUserSelfScript,omitnil,omitempty" name:"OnlyUserSelfScript"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *ListSQLFolderContentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLFolderContentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "Keyword")
	delete(f, "OnlyFolderNode")
	delete(f, "OnlyUserSelfScript")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSQLFolderContentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLFolderContentsResponseParams struct {
	// 结果列表
	Data []*SQLFolderNode `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSQLFolderContentsResponse struct {
	*tchttp.BaseResponse
	Response *ListSQLFolderContentsResponseParams `json:"Response"`
}

func (r *ListSQLFolderContentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLFolderContentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLScriptRunsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 脚本id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 任务id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 执行人
	ExecuteUserUin *string `json:"ExecuteUserUin,omitnil,omitempty" name:"ExecuteUserUin"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ListSQLScriptRunsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 脚本id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 任务id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 执行人
	ExecuteUserUin *string `json:"ExecuteUserUin,omitnil,omitempty" name:"ExecuteUserUin"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ListSQLScriptRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLScriptRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ScriptId")
	delete(f, "JobId")
	delete(f, "SearchWord")
	delete(f, "ExecuteUserUin")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSQLScriptRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSQLScriptRunsResponseParams struct {
	// 数据探索任务
	Data []*JobDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSQLScriptRunsResponse struct {
	*tchttp.BaseResponse
	Response *ListSQLScriptRunsResponseParams `json:"Response"`
}

func (r *ListSQLScriptRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSQLScriptRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSchemaPage struct {
	// Schema记录列表
	Items []*SchemaInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页总页数
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListSchemaRequestParams struct {
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 目录名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据源ID
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库模式搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListSchemaRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 目录名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据源ID
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库模式搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSchemaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "CatalogName")
	delete(f, "DatasourceId")
	delete(f, "DatabaseName")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSchemaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSchemaResponseParams struct {
	// 分页数据
	Data *ListSchemaPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSchemaResponse struct {
	*tchttp.BaseResponse
	Response *ListSchemaResponseParams `json:"Response"`
}

func (r *ListSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTablePage struct {
	// Schema记录列表
	Items []*TableInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页总页数
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 记录总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListTableRequestParams struct {
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 目录名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据源ID
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListTableRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，最大500
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 目录名称
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// 数据源ID
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "CatalogName")
	delete(f, "DatasourceId")
	delete(f, "DatabaseName")
	delete(f, "SchemaName")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTableResponseParams struct {
	// 分页数据
	Data *ListTablePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTableResponse struct {
	*tchttp.BaseResponse
	Response *ListTableResponseParams `json:"Response"`
}

func (r *ListTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskInfo struct {
	// 任务数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TaskBaseAttribute `json:"Items,omitnil,omitempty" name:"Items"`

	// 当前请求的数据页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 当前请求的数据页条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 满足查询条件的数据总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 满足查询条件的数据总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

// Predefined struct for user
type ListTaskInstanceExecutionsRequestParams struct {
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例唯一标识，可以通过ListInstances获取
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区**timeZone, 传入的时间字符串的所在时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 每页大小，默认10, 最大200
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type ListTaskInstanceExecutionsRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例唯一标识，可以通过ListInstances获取
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区**timeZone, 传入的时间字符串的所在时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 每页大小，默认10, 最大200
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码，默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *ListTaskInstanceExecutionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstanceExecutionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskInstanceExecutionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskInstanceExecutionsResponseParams struct {
	// 实例详情
	Data *TaskInstanceExecutions `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskInstanceExecutionsResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskInstanceExecutionsResponseParams `json:"Response"`
}

func (r *ListTaskInstanceExecutionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstanceExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskInstancesRequestParams struct {
	// **项目ID**
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **页码，整型**
	// 配合pageSize使用且不能小于1， 默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// **每页显示的条数，默认为10，最小值为1、最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// **任务名称 或 任务ID**
	// 支持模糊搜索过滤, 多个用 英文逗号, 分割
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// **时区**timeZone, 传入的时间字符串的所在时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **实例类型** 
	// - 0 表示补录类型 
	// - 1 表示周期实例 
	// - 2 表示非周期实例
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// **实例状态** - WAIT_EVENT: 等待事件 - WAIT_UPSTREAM: 等待上游 - WAIT_RUN: 等待运行 - RUNNING: 运行中 - SKIP_RUNNING: 跳过运行 - FAILED_RETRY: 失败重试 - EXPIRED: 失败 - COMPLETED: 成功
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// **任务类型Id**
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// **任务周期类型** * ONEOFF_CYCLE: 一次性 * YEAR_CYCLE: 年 * MONTH_CYCLE: 月 * WEEK_CYCLE: 周 * DAY_CYCLE: 天 * HOUR_CYCLE: 小时 * MINUTE_CYCLE: 分钟 * CRONTAB_CYCLE: crontab表达式类型
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// **任务负责人id**
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// **任务所属文件id**
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// **任务所属工作流id**
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// **执行资源组Id**
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// **实例计划调度时间过滤条件**过滤起始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ScheduleTimeFrom *string `json:"ScheduleTimeFrom,omitnil,omitempty" name:"ScheduleTimeFrom"`

	// **实例计划调度时间过滤条件**过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ScheduleTimeTo *string `json:"ScheduleTimeTo,omitnil,omitempty" name:"ScheduleTimeTo"`

	// **实例执行开始时间过滤条件**过滤起始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// **实例执行开始时间过滤条件**
	// 过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// **实例最近更新时间过滤条件**
	// 过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	LastUpdateTimeFrom *string `json:"LastUpdateTimeFrom,omitnil,omitempty" name:"LastUpdateTimeFrom"`

	// **实例最近更新时间过滤条件**
	// 过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	LastUpdateTimeTo *string `json:"LastUpdateTimeTo,omitnil,omitempty" name:"LastUpdateTimeTo"`

	// **查询结果排序字段**- SCHEDULE_DATE 表示 根据计划调度时间排序- START_TIME 表示 根据实例开始执行时间排序- END_TIME 表示 根据实例结束执行时间排序- COST_TIME 表示 根据实例执行时长排序
	SortColumn *string `json:"SortColumn,omitnil,omitempty" name:"SortColumn"`

	// **实例排序方式**
	// 
	// - ASC 
	// - DESC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type ListTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// **项目ID**
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **页码，整型**
	// 配合pageSize使用且不能小于1， 默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// **每页显示的条数，默认为10，最小值为1、最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// **任务名称 或 任务ID**
	// 支持模糊搜索过滤, 多个用 英文逗号, 分割
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// **时区**timeZone, 传入的时间字符串的所在时区，默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **实例类型** 
	// - 0 表示补录类型 
	// - 1 表示周期实例 
	// - 2 表示非周期实例
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// **实例状态** - WAIT_EVENT: 等待事件 - WAIT_UPSTREAM: 等待上游 - WAIT_RUN: 等待运行 - RUNNING: 运行中 - SKIP_RUNNING: 跳过运行 - FAILED_RETRY: 失败重试 - EXPIRED: 失败 - COMPLETED: 成功
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// **任务类型Id**
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// **任务周期类型** * ONEOFF_CYCLE: 一次性 * YEAR_CYCLE: 年 * MONTH_CYCLE: 月 * WEEK_CYCLE: 周 * DAY_CYCLE: 天 * HOUR_CYCLE: 小时 * MINUTE_CYCLE: 分钟 * CRONTAB_CYCLE: crontab表达式类型
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// **任务负责人id**
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// **任务所属文件id**
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// **任务所属工作流id**
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// **执行资源组Id**
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// **实例计划调度时间过滤条件**过滤起始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ScheduleTimeFrom *string `json:"ScheduleTimeFrom,omitnil,omitempty" name:"ScheduleTimeFrom"`

	// **实例计划调度时间过滤条件**过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ScheduleTimeTo *string `json:"ScheduleTimeTo,omitnil,omitempty" name:"ScheduleTimeTo"`

	// **实例执行开始时间过滤条件**过滤起始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// **实例执行开始时间过滤条件**
	// 过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// **实例最近更新时间过滤条件**
	// 过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	LastUpdateTimeFrom *string `json:"LastUpdateTimeFrom,omitnil,omitempty" name:"LastUpdateTimeFrom"`

	// **实例最近更新时间过滤条件**
	// 过滤截止时间，时间格式为 yyyy-MM-dd HH:mm:ss
	LastUpdateTimeTo *string `json:"LastUpdateTimeTo,omitnil,omitempty" name:"LastUpdateTimeTo"`

	// **查询结果排序字段**- SCHEDULE_DATE 表示 根据计划调度时间排序- START_TIME 表示 根据实例开始执行时间排序- END_TIME 表示 根据实例结束执行时间排序- COST_TIME 表示 根据实例执行时长排序
	SortColumn *string `json:"SortColumn,omitnil,omitempty" name:"SortColumn"`

	// **实例排序方式**
	// 
	// - ASC 
	// - DESC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

func (r *ListTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "TimeZone")
	delete(f, "InstanceType")
	delete(f, "InstanceState")
	delete(f, "TaskTypeId")
	delete(f, "CycleType")
	delete(f, "OwnerUin")
	delete(f, "FolderId")
	delete(f, "WorkflowId")
	delete(f, "ExecutorGroupId")
	delete(f, "ScheduleTimeFrom")
	delete(f, "ScheduleTimeTo")
	delete(f, "StartTimeFrom")
	delete(f, "StartTimeTo")
	delete(f, "LastUpdateTimeFrom")
	delete(f, "LastUpdateTimeTo")
	delete(f, "SortColumn")
	delete(f, "SortType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskInstancesResponseParams struct {
	// 实例结果集
	Data *TaskInstancePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskInstancesResponseParams `json:"Response"`
}

func (r *ListTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTaskVersions struct {
	// 记录列表	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TaskVersion `json:"Items,omitnil,omitempty" name:"Items"`

	// 满足查询条件的数据总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 满足查询条件的数据总页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 当前页记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 当前请求的数据页条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前请求的数据页数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

// Predefined struct for user
type ListTaskVersionsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 保存版本：SAVE
	// 提交版本：SUBMIT
	// 默认为SAVE
	TaskVersionType *string `json:"TaskVersionType,omitnil,omitempty" name:"TaskVersionType"`

	// 请求的数据页数。默认值为1，取值大于等于1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数。默认值为10 ，最小值为10，最大值为200。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListTaskVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 保存版本：SAVE
	// 提交版本：SUBMIT
	// 默认为SAVE
	TaskVersionType *string `json:"TaskVersionType,omitnil,omitempty" name:"TaskVersionType"`

	// 请求的数据页数。默认值为1，取值大于等于1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数。默认值为10 ，最小值为10，最大值为200。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListTaskVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "TaskVersionType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskVersionsResponseParams struct {
	// 版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ListTaskVersions `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskVersionsResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskVersionsResponseParams `json:"Response"`
}

func (r *ListTaskVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求的数据页数。默认值为1，取值大于等于1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数。默认值为10 ，最小值为10，最大值为200
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 所属工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 责任人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 任务类型
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务状态
	// * N: 新建 
	// * Y: 调度中 
	// * F: 已下线 
	// * O: 已暂停 
	// * T: 下线中 
	// * INVALID: 已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 提交状态
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// BundleId信息
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 创建人ID
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 修改时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ListTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求的数据页数。默认值为1，取值大于等于1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数。默认值为10 ，最小值为10，最大值为200
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 所属工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 责任人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 任务类型
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务状态
	// * N: 新建 
	// * Y: 调度中 
	// * F: 已下线 
	// * O: 已暂停 
	// * T: 下线中 
	// * INVALID: 已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 提交状态
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// BundleId信息
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 创建人ID
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 修改时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

func (r *ListTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "TaskName")
	delete(f, "WorkflowId")
	delete(f, "OwnerUin")
	delete(f, "TaskTypeId")
	delete(f, "Status")
	delete(f, "Submit")
	delete(f, "BundleId")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTime")
	delete(f, "CreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksResponseParams struct {
	// 任务分页信息
	Data *ListTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListTasksResponseParams `json:"Response"`
}

func (r *ListTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTenantRolesRequestParams struct {
	// 角色中文展示名模糊查询，只能传一个值
	RoleDisplayName *string `json:"RoleDisplayName,omitnil,omitempty" name:"RoleDisplayName"`
}

type ListTenantRolesRequest struct {
	*tchttp.BaseRequest
	
	// 角色中文展示名模糊查询，只能传一个值
	RoleDisplayName *string `json:"RoleDisplayName,omitnil,omitempty" name:"RoleDisplayName"`
}

func (r *ListTenantRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTenantRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleDisplayName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTenantRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTenantRolesResponseParams struct {
	// 主账号角色列表
	Data []*SystemRole `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTenantRolesResponse struct {
	*tchttp.BaseResponse
	Response *ListTenantRolesResponseParams `json:"Response"`
}

func (r *ListTenantRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTenantRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamOpsTasksRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页页码
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListUpstreamOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页页码
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListUpstreamOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUpstreamOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamOpsTasksResponseParams struct {
	// 上游任务详情
	Data *ParentDependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUpstreamOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListUpstreamOpsTasksResponseParams `json:"Response"`
}

func (r *ListUpstreamOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTaskInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区** timeZone, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **页码，整型**配合pageSize使用且不能小于1， 默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// **每页显示的条数，默认为10，最小值为1、最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListUpstreamTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// **时区** timeZone, 默认UTC+8
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// **页码，整型**配合pageSize使用且不能小于1， 默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// **每页显示的条数，默认为10，最小值为1、最大值为100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListUpstreamTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKey")
	delete(f, "TimeZone")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUpstreamTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTaskInstancesResponseParams struct {
	// 上游实例列表
	Data *TaskInstancePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUpstreamTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListUpstreamTaskInstancesResponseParams `json:"Response"`
}

func (r *ListUpstreamTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTasksRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 请求的数据页数。默认值为1，取值大于等于1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 请求的数据页数。默认值为1，取值大于等于1。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListUpstreamTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 请求的数据页数。默认值为1，取值大于等于1。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 请求的数据页数。默认值为1，取值大于等于1。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListUpstreamTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUpstreamTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUpstreamTasksResponseParams struct {
	// 上游任务详情
	Data *DependencyConfigPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUpstreamTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListUpstreamTasksResponseParams `json:"Response"`
}

func (r *ListUpstreamTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUpstreamTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowFoldersRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹绝对路径，如/abc/de，如果是根目录则传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 数据页数，大于等于1。默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数，最小为10条，最大为200 条。默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListWorkflowFoldersRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹绝对路径，如/abc/de，如果是根目录则传/
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 数据页数，大于等于1。默认1
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数，最小为10条，最大为200 条。默认10
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListWorkflowFoldersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowFoldersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentFolderPath")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowFoldersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowFoldersResponseParams struct {
	// 分页的文件夹查询结果
	Data *WorkflowFolderPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowFoldersResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowFoldersResponseParams `json:"Response"`
}

func (r *ListWorkflowFoldersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListWorkflowInfo struct {
	// 列表item
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 满足查询条件的数据总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 当前请求的数据页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 当前请求的数据页条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 满足查询条件的数据总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListWorkflowsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求的数据页数。默认值为1，取值大于等于1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数。默认值为10 ，最小值为10，最大值为200
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 工作流所属文件夹
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 工作流类型，cycle和manual
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// bundleId项
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 负责人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建人ID
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 修改时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ListWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求的数据页数。默认值为1，取值大于等于1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数。默认值为10 ，最小值为10，最大值为200
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 工作流所属文件夹
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 工作流类型，cycle和manual
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// bundleId项
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 负责人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建人ID
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 修改时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	ModifyTime []*string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间区间 yyyy-MM-dd HH:mm:ss，需要在数组填入两个时间
	CreateTime []*string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

func (r *ListWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "ParentFolderPath")
	delete(f, "WorkflowType")
	delete(f, "BundleId")
	delete(f, "OwnerUin")
	delete(f, "CreateUserUin")
	delete(f, "ModifyTime")
	delete(f, "CreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowsResponseParams struct {
	// 查询工作流分页信息
	Data *ListWorkflowInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowsResponseParams `json:"Response"`
}

func (r *ListWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricData struct {
	// 指标名称
	// 
	// - ConcurrencyUsage --- 并发使用率
	// - CpuCoreUsage --- cpu使用率
	// - CpuLoad --- cpu负载
	// - DevelopQueueTask --- 正在队列中的开发任务数量
	// - DevelopRunningTask --- 正在运行的开发任务数量
	// - DevelopSchedulingTask --- 正在调度的开发任务数量
	// - DiskUsage --- 磁盘使用情况
	// - DiskUsed --- 磁盘已用量
	// - MaximumConcurrency --- 最大并发
	// - MemoryLoad --- 内存负载
	// - MemoryUsage --- 内存使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 当前值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotValue *uint64 `json:"SnapshotValue,omitnil,omitempty" name:"SnapshotValue"`

	// 指标趋势
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrendList []*TrendData `json:"TrendList,omitnil,omitempty" name:"TrendList"`
}

type ModifyAlarmRuleResult struct {
	// 是否更新成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type MonitorWhiteTask struct {
	// 配置白名单的对应的工作流/项目的id
	MonitorObjectId *string `json:"MonitorObjectId,omitnil,omitempty" name:"MonitorObjectId"`

	// 白名单任务列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type NotebookSessionInfo struct {
	// 会话ID
	NotebookSessionId *string `json:"NotebookSessionId,omitnil,omitempty" name:"NotebookSessionId"`

	// 会话名称
	NotebookSessionName *string `json:"NotebookSessionName,omitnil,omitempty" name:"NotebookSessionName"`
}

type NotificationFatigue struct {
	// 告警次数
	NotifyCount *uint64 `json:"NotifyCount,omitnil,omitempty" name:"NotifyCount"`

	// 告警间隔，分钟
	NotifyInterval *uint64 `json:"NotifyInterval,omitnil,omitempty" name:"NotifyInterval"`

	// 免打扰时间，例如示例值
	// [{DaysOfWeek: [1, 2], StartTime: "00:00:00", EndTime: "09:00:00"}]	
	// 每周一、周二的00:00到09:00免打扰
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuietIntervals []*AlarmQuietInterval `json:"QuietIntervals,omitnil,omitempty" name:"QuietIntervals"`
}

type OperateResult struct {
	// 操作结果1 成功 其他失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type OpsAsyncJobDetail struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 操作id
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`

	// 异步操作类型
	AsyncType *string `json:"AsyncType,omitnil,omitempty" name:"AsyncType"`

	// 异步操作状态：初始状态: INIT; 运行中: RUNNING; 成功: SUCCESS; 失败: FAIL; 部分成功: PART_SUCCESS
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`

	// 子操作总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSubProcessCount *uint64 `json:"TotalSubProcessCount,omitnil,omitempty" name:"TotalSubProcessCount"`

	// 已完成的子操作个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedSubProcessCount *uint64 `json:"FinishedSubProcessCount,omitnil,omitempty" name:"FinishedSubProcessCount"`

	// 已成功的子操作个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessSubProcessCount *uint64 `json:"SuccessSubProcessCount,omitnil,omitempty" name:"SuccessSubProcessCount"`

	// 操作人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 操作创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OpsAsyncResponse struct {
	// 异步执行记录Id
	AsyncId *string `json:"AsyncId,omitnil,omitempty" name:"AsyncId"`
}

type OpsTaskDepend struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 任务状态:
	// - Y: 调度中
	// - F: 已下线
	// - O: 已暂停
	// - T: 下线中
	// - INVALID: 已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型Id：
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务类型描述
	//  - 20 :  通用数据同步
	//  - 25 :  ETLTaskType
	//  - 26 :  ETLTaskType
	//  - 30 :  python
	//  - 31 :  pyspark
	//  - 34 :  hivesql
	//  - 35 :  shell
	//  - 36 :  sparksql
	//  - 21 :  jdbcsql
	//  - 32 :  dlc
	//  - 33 :  ImpalaTaskType
	//  - 40 :  CDWTaskType
	//  - 41 :  kettle
	//  - 42 :  TCHouse-X
	//  - 43 :  TCHouse-X SQL
	//  - 46 :  dlcsparkTaskType
	//  - 47 :  TiOneMachineLearningTaskType
	//  - 48 :  Trino
	//  - 50 :  DLCPyspark
	//  - 23 :  TencentDistributedSQL
	//  - 39 :  spark
	//  - 92 :  MRTaskType
	//  - 38 :  ShellScript
	//  - 70 :  HiveSQLScrip
	//  - 130 :  分支
	//  - 131 :  归并
	//  - 132 :  Notebook探索
	//  - 133 :  SSH节点
	//  - 134 :  StarRocks
	//  - 137 :  For-each
	//  - 10000 :  自定义业务通用
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 首次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstRunTime *string `json:"FirstRunTime,omitnil,omitempty" name:"FirstRunTime"`

	// 调度计划展示描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// 任务周期类型
	// * ONEOFF_CYCLE: 一次性
	// * YEAR_CYCLE: 年
	// * MONTH_CYCLE: 月
	// * WEEK_CYCLE: 周
	// * DAY_CYCLE: 天
	// * HOUR_CYCLE: 小时
	// * MINUTE_CYCLE: 分钟
	// * CRONTAB_CYCLE: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 执行开始时间, 格式HH:mm, 如00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行结束时间, 格式HH:mm, 如23:59
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`
}

type OpsWorkflow struct {
	// 任务数量
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 工作流文件id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流类型
	//  - cycle周期
	//  - manual手动
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 工作流描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 负责人userId,多个‘；’隔开
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 工作流状态
	// * ALL_RUNNING : 全部调度中
	// * ALL_FREEZED : 全部已暂停
	// * ALL_STOPPTED : 全部已下线
	// * PART_RUNNING : 部分调度中
	// * ALL_NO_RUNNING : 全部未调度
	// * ALL_INVALID : 全部已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 工作流创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近更新时间, 包含开发、生产变更
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最近更新人，包含开发、生产变更
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`
}

type OpsWorkflowDetail struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流类型：
	//  - cycle 周期；
	//  - manual手动
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 延时执行时间,unit=minute
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupTime *uint64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 配置生效日期 开始日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 配置结束日期 结束日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务周期类型
	// * ONEOFF_CYCLE: 一次性
	// * YEAR_CYCLE: 年
	// * MONTH_CYCLE: 月
	// * WEEK_CYCLE: 周
	// * DAY_CYCLE: 天
	// * HOUR_CYCLE: 小时
	// * MINUTE_CYCLE: 分钟
	// * CRONTAB_CYCLE: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 文件夹Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 任务实例初始化策略
	//  - T_PLUS_1（T+1）：延迟一天初始化
	//  - T_PLUS_0（T+0）：当天初始化
	//  - T_MINUS_1（T-1）：提前一天初始化
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`

	// 调度计划释义
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitnil,omitempty" name:"SchedulerDesc"`

	// 工作流首次提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 工作流最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestSubmitTime *string `json:"LatestSubmitTime,omitnil,omitempty" name:"LatestSubmitTime"`

	// 工作流状态
	// * ALL_RUNNING : 全部调度中
	// * ALL_FREEZED : 全部已暂停
	// * ALL_STOPPTED : 全部已下线
	// * PART_RUNNING : 部分调度中
	// * ALL_NO_RUNNING : 全部未调度
	// * ALL_INVALID : 全部已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 负责人, 多个以‘；’隔开
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`
}

type OpsWorkflows struct {
	// 记录列表	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OpsWorkflow `json:"Items,omitnil,omitempty" name:"Items"`

	// 结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type OutTaskParameter struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type PageRoles struct {
	// 角色信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SystemRole `json:"Items,omitnil,omitempty" name:"Items"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type ParamInfo struct {
	// 参数名
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type ParentDependencyConfigPage struct {
	// 结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OpsTaskDepend `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type PauseOpsTasksAsyncRequestParams struct {
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 是否需要终止已生成实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type PauseOpsTasksAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 是否需要终止已生成实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *PauseOpsTasksAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOpsTasksAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseOpsTasksAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseOpsTasksAsyncResponseParams struct {
	// 异步操作结果
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseOpsTasksAsyncResponse struct {
	*tchttp.BaseResponse
	Response *PauseOpsTasksAsyncResponseParams `json:"Response"`
}

func (r *PauseOpsTasksAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseOpsTasksAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Project struct {
	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识，英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目显示名称，可以为中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目创建人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 项目责任人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectOwnerUin *string `json:"ProjectOwnerUin,omitnil,omitempty" name:"ProjectOwnerUin"`

	// 项目状态：0：禁用，1：启用，-3:禁用中，2：启用中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目模式，SIMPLE：简单模式 STANDARD：标准模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectModel *string `json:"ProjectModel,omitnil,omitempty" name:"ProjectModel"`
}

type ProjectBrief struct {
	// 分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Project `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type ProjectInstanceStatisticsAlarmInfo struct {
	// 告警类型 
	// 
	// projectFailureInstanceUpwardFluctuationAlarm: 失败实例向上波动告警
	// 
	// projectSuccessInstanceDownwardFluctuationAlarm： 成功实例向下波动告警
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// 实例成功数向下波动比例告警阀值；实例失败数向上波动比例告警阀值
	InstanceThresholdCountPercent *uint64 `json:"InstanceThresholdCountPercent,omitnil,omitempty" name:"InstanceThresholdCountPercent"`

	// 累计实例数波动阀值
	InstanceThresholdCount *uint64 `json:"InstanceThresholdCount,omitnil,omitempty" name:"InstanceThresholdCount"`

	// 稳定性次数阈值(防抖动配置统计周期数)
	StabilizeThreshold *uint64 `json:"StabilizeThreshold,omitnil,omitempty" name:"StabilizeThreshold"`

	// 稳定性统计周期(防抖动配置统计周期数)
	StabilizeStatisticsCycle *uint64 `json:"StabilizeStatisticsCycle,omitnil,omitempty" name:"StabilizeStatisticsCycle"`

	// 是否累计计算,false:连续,true:累计
	IsCumulant *bool `json:"IsCumulant,omitnil,omitempty" name:"IsCumulant"`

	// 当日累计实例数;
	// 当天失败实例数向下波动量
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`
}

type ProjectRequest struct {
	// 项目标识，英文名，以字母开头，可包含字母、数字和下划线，不能超过32个字符
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目显示名称，可以为中文名，以字母开头，可包含字母、数字和下划线，不能超过32个字符
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 项目模式，SIMPLE（默认）：简单模式 STANDARD：标准模式
	ProjectModel *string `json:"ProjectModel,omitnil,omitempty" name:"ProjectModel"`
}

type ProjectResult struct {
	// 返回的结果 true/false
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type ProjectUserRole struct {
	// 租户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootAccountId *string `json:"RootAccountId,omitnil,omitempty" name:"RootAccountId"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 显示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户角色对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*SystemRole `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 是否创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCreator *bool `json:"IsCreator,omitnil,omitempty" name:"IsCreator"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否项目负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsProjectOwner *bool `json:"IsProjectOwner,omitnil,omitempty" name:"IsProjectOwner"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

type ProjectUsersBrief struct {
	// 用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ProjectUserRole `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type ReconciliationStrategyInfo struct {
	// 离线告警规则类型
	// reconciliationFailure： 离线对账失败告警
	// reconciliationOvertime： 离线对账任务运行超时告警(需配置超时时间)
	// reconciliationMismatch： 离线对账不一致条数告警(需配置不一致条数阀值)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 对账不一致条数阀值， RuleType=reconciliationMismatch对账不一致条数类型，需要配置该字段，无默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MismatchCount *uint64 `json:"MismatchCount,omitnil,omitempty" name:"MismatchCount"`

	// 对账任务运行超时阀值： 小时， 默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hour *int64 `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 对账任务运行超时阀值： 分钟， 默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`
}

// Predefined struct for user
type RegisterLineageRequestParams struct {
	// 需要注册的血缘关系列表
	Relations []*LineagePair `json:"Relations,omitnil,omitempty" name:"Relations"`
}

type RegisterLineageRequest struct {
	*tchttp.BaseRequest
	
	// 需要注册的血缘关系列表
	Relations []*LineagePair `json:"Relations,omitnil,omitempty" name:"Relations"`
}

func (r *RegisterLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Relations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterLineageResponseParams struct {
	// 注册结果
	Data *OperateResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterLineageResponse struct {
	*tchttp.BaseResponse
	Response *RegisterLineageResponseParams `json:"Response"`
}

func (r *RegisterLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RelateTask struct {
	// 任务类型
	// 
	// - OfflineIntegration --- 离线集成任务
	// - RealtimeIntegration --- 实时集成任务
	// - DataDevelopment --- 数据开发任务
	// - DataQuality --- 数据质量任务
	// - DataService --- 数据服务任务
	// - MetadataCollection --- 元数据采集任务
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务数量
	TaskNum *int64 `json:"TaskNum,omitnil,omitempty" name:"TaskNum"`

	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskList []*BriefTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`
}

// Predefined struct for user
type RemoveMemberProjectRoleRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 角色id列表，目前支持的项目角色有
	// - 308335260274237440 (项目管理员)
	// - 308335260676890624 (数据工程师)
	// - 308335260844662784 (运维工程师)
	// - 308335260945326080 (普通成员)
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type RemoveMemberProjectRoleRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户id
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 角色id列表，目前支持的项目角色有
	// - 308335260274237440 (项目管理员)
	// - 308335260676890624 (数据工程师)
	// - 308335260844662784 (运维工程师)
	// - 308335260945326080 (普通成员)
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *RemoveMemberProjectRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMemberProjectRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserUin")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveMemberProjectRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveMemberProjectRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveMemberProjectRoleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveMemberProjectRoleResponseParams `json:"Response"`
}

func (r *RemoveMemberProjectRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMemberProjectRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunTaskInstancesAsyncRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例id列表,可以从ListInstances中获取
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子，默认1
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 是否检查上游任务： ALL（全部）、 MAKE_SCOPE（选中）、NONE （全部不检查），默认NONE
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// 下游实例范围 WORKFLOW: 所在工作流 PROJECT: 所在项目 ALL: 所有跨工作流依赖的项目，默认WORKFLOW
	SonRangeType *string `json:"SonRangeType,omitnil,omitempty" name:"SonRangeType"`

	// 重跑是否忽略事件监听
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 自定义实例运行并发度, 如果不配置，则使用任务原有自依赖
	RedefineParallelNum *int64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// 自定义的工作流自依赖: yes开启，no关闭，如果不配置，则使用工作流原有自依赖
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// 重跑实例自定义参数
	RedefineParamList *KVMap `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`
}

type RerunTaskInstancesAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例id列表,可以从ListInstances中获取
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子，默认1
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 是否检查上游任务： ALL（全部）、 MAKE_SCOPE（选中）、NONE （全部不检查），默认NONE
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// 下游实例范围 WORKFLOW: 所在工作流 PROJECT: 所在项目 ALL: 所有跨工作流依赖的项目，默认WORKFLOW
	SonRangeType *string `json:"SonRangeType,omitnil,omitempty" name:"SonRangeType"`

	// 重跑是否忽略事件监听
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 自定义实例运行并发度, 如果不配置，则使用任务原有自依赖
	RedefineParallelNum *int64 `json:"RedefineParallelNum,omitnil,omitempty" name:"RedefineParallelNum"`

	// 自定义的工作流自依赖: yes开启，no关闭，如果不配置，则使用工作流原有自依赖
	RedefineSelfWorkflowDependency *string `json:"RedefineSelfWorkflowDependency,omitnil,omitempty" name:"RedefineSelfWorkflowDependency"`

	// 重跑实例自定义参数
	RedefineParamList *KVMap `json:"RedefineParamList,omitnil,omitempty" name:"RedefineParamList"`
}

func (r *RerunTaskInstancesAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunTaskInstancesAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKeyList")
	delete(f, "RerunType")
	delete(f, "CheckParentType")
	delete(f, "SonRangeType")
	delete(f, "SkipEventListening")
	delete(f, "RedefineParallelNum")
	delete(f, "RedefineSelfWorkflowDependency")
	delete(f, "RedefineParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RerunTaskInstancesAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunTaskInstancesAsyncResponseParams struct {
	// 批量重跑操作的返回的异步id, 可以在接口GetAsyncJob获取具体执行详情
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RerunTaskInstancesAsyncResponse struct {
	*tchttp.BaseResponse
	Response *RerunTaskInstancesAsyncResponseParams `json:"Response"`
}

func (r *RerunTaskInstancesAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunTaskInstancesAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceFile struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源文件名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源文件路径
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// 资源对象COS存储路径
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`

	// 资源文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// 资源大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 创建用户ID
	CreatorUserUin *string `json:"CreatorUserUin,omitnil,omitempty" name:"CreatorUserUin"`

	// 创建用户名称
	CreatorUserName *string `json:"CreatorUserName,omitnil,omitempty" name:"CreatorUserName"`

	// 最后更新用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserName *string `json:"UpdateUserName,omitnil,omitempty" name:"UpdateUserName"`

	// 最后更新用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// COS 桶
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// COS 地域
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 资源来源模式
	ResourceSourceMode *string `json:"ResourceSourceMode,omitnil,omitempty" name:"ResourceSourceMode"`

	// 本地工程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 本地工程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type ResourceFileItem struct {
	// 资源文件ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源文件名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源文件类型
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// 资源路径
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`
}

type ResourceFilePage struct {
	// 任务集合信息
	Items []*ResourceFileItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 总页数
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 当前页
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ResourceFolder struct {
	// 资源文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 创建人ID
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 创建人名称
	CreateUserName *string `json:"CreateUserName,omitnil,omitempty" name:"CreateUserName"`

	// 文件夹路径
	FolderPath *string `json:"FolderPath,omitnil,omitempty" name:"FolderPath"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type ResourceFolderPage struct {
	// 资源文件夹集合信息
	Items []*ResourceFolder `json:"Items,omitnil,omitempty" name:"Items"`

	// 总页数
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 当前页
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ResourceGroupMetrics struct {
	// 资源组规格相关：cpu个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 资源组规格相关：磁盘规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskVolume *uint64 `json:"DiskVolume,omitnil,omitempty" name:"DiskVolume"`

	// 资源组规格相关：内存大小，单位:G
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 资源组生命周期, 单位：天
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeCycle *uint64 `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// 资源组规格相关：最高并发
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumConcurrency *uint64 `json:"MaximumConcurrency,omitnil,omitempty" name:"MaximumConcurrency"`

	// 资源组状态
	// 
	// - 0 --- 初始化中
	// - 1 --- 运行中
	// - 2 --- 运行异常
	// - 3 --- 释放中
	// - 4 --- 已释放
	// - 5 --- 创建中
	// - 6 --- 创建失败
	// - 7 --- 更新中
	// - 8 --- 更新失败
	// - 9 --- 已到期
	// - 10 --- 释放失败
	// - 11 --- 使用中
	// - 12 --- 未使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 指标详情
	MetricSnapshots []*MetricData `json:"MetricSnapshots,omitnil,omitempty" name:"MetricSnapshots"`
}

type ResourceGroupSpecification struct {
	// 资源组规格
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 数量
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type ResourceNumber struct {
	// 增加/减少枚举
	// 
	// - ADD -- 增加
	// - DELETE -- 减少
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 增加/减少资源包的数量
	Quantity *int64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`
}

type ResourceResult struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`
}

type ResourceStatus struct {
	// 资源组操作状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type ResourceType struct {
	// 资源组类型
	// 
	// - Schedule --- 调度资源组
	// - Integration --- 集成资源组
	// - DataService -- 数据服务资源组
	ResourceGroupType *string `json:"ResourceGroupType,omitnil,omitempty" name:"ResourceGroupType"`

	// 集成资源组，细分实时资源组和离线资源组(集成、调度、数据服务资源组不可以同时购买)
	Integration *IntegrationResource `json:"Integration,omitnil,omitempty" name:"Integration"`

	// 调度资源组(集成、调度、数据服务资源组不可以同时购买)
	// 
	// - s_test(测试规格)
	// - s_small(基础规格)
	// - s_medium(普及规格)
	// - s_large(专业规格)
	// 
	// 
	Schedule *ResourceGroupSpecification `json:"Schedule,omitnil,omitempty" name:"Schedule"`

	// 数据服务资源组(集成、调度、数据服务资源组不可以同时购买)
	// 
	// - ds_t(测试规格)
	// - ds_s(基础规格)
	// - ds_m(普及规格)
	// - ds_l(专业规格)
	DataService *ResourceGroupSpecification `json:"DataService,omitnil,omitempty" name:"DataService"`
}

// Predefined struct for user
type RunSQLScriptRequestParams struct {
	// 脚本id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 脚本内容，不传则默认执行已保存的全量脚本内容；若传递则要用Base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 高级运行参数，JSON格式base64编码
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type RunSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 脚本内容，不传则默认执行已保存的全量脚本内容；若传递则要用Base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 高级运行参数，JSON格式base64编码
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

func (r *RunSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	delete(f, "ScriptContent")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSQLScriptResponseParams struct {
	// 数据探索任务
	Data *JobDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *RunSQLScriptResponseParams `json:"Response"`
}

func (r *RunSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SQLContentActionResult struct {
	// 操作是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type SQLFolderNode struct {
	// 唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型 folder，script
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 父文件夹path，/aaa/bbb/ccc
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 是否叶子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 业务参数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 权限范围: SHARED, PRIVATE
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// 节点路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 当前用户对节点拥有的权限	
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePermission *string `json:"NodePermission,omitnil,omitempty" name:"NodePermission"`

	// 子节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*SQLFolderNode `json:"Children,omitnil,omitempty" name:"Children"`

	// 文件责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type SQLScript struct {
	// 脚本id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// 脚本所有者 uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 父文件夹path，/aaa/bbb/ccc
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentFolderPath *string `json:"ParentFolderPath,omitnil,omitempty" name:"ParentFolderPath"`

	// 脚本配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// 脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 最近一次操作人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 更新时间 yyyy-MM-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建时间 yyyy-MM-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 权限范围：SHARED, PRIVATE
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`

	// 节点全路径，/aaa/bbb/ccc.ipynb，由各个节点的名称组成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type SQLScriptConfig struct {
	// 数据源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据源环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceEnv *string `json:"DatasourceEnv,omitnil,omitempty" name:"DatasourceEnv"`

	// 计算资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeResource *string `json:"ComputeResource,omitnil,omitempty" name:"ComputeResource"`

	// 执行资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 高级运行参数,变量替换，map-json String,String
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 高级设置，执行配置参数，map-json String,String. 采用Base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceConfig *string `json:"AdvanceConfig,omitnil,omitempty" name:"AdvanceConfig"`
}

type SQLStopResult struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type SchemaInfo struct {
	// Schema GUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guid *string `json:"Guid,omitnil,omitempty" name:"Guid"`

	// Schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`
}

// Predefined struct for user
type SetSuccessTaskInstancesAsyncRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例id列表,可以从ListInstances中获取
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

type SetSuccessTaskInstancesAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例id列表,可以从ListInstances中获取
	InstanceKeyList []*string `json:"InstanceKeyList,omitnil,omitempty" name:"InstanceKeyList"`
}

func (r *SetSuccessTaskInstancesAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSuccessTaskInstancesAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InstanceKeyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetSuccessTaskInstancesAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSuccessTaskInstancesAsyncResponseParams struct {
	// 批量置成功操作的返回的异步id, 可以在接口GetAsyncJob获取具体执行详情
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetSuccessTaskInstancesAsyncResponse struct {
	*tchttp.BaseResponse
	Response *SetSuccessTaskInstancesAsyncResponseParams `json:"Response"`
}

func (r *SetSuccessTaskInstancesAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSuccessTaskInstancesAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SqlCreateResult struct {
	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

// Predefined struct for user
type StartOpsTasksRequestParams struct {
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 启动时是否补录上次暂停到当前的中间实例，默认false即不补录
	EnableDataBackfill *bool `json:"EnableDataBackfill,omitnil,omitempty" name:"EnableDataBackfill"`
}

type StartOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 启动时是否补录上次暂停到当前的中间实例，默认false即不补录
	EnableDataBackfill *bool `json:"EnableDataBackfill,omitnil,omitempty" name:"EnableDataBackfill"`
}

func (r *StartOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "EnableDataBackfill")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartOpsTasksResponseParams struct {
	// 异步操作结果
	Data *StartTasks `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *StartOpsTasksResponseParams `json:"Response"`
}

func (r *StartOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StartTasks struct {
	// 任务启动是否成功
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type StopOpsTasksAsyncRequestParams struct {
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 是否终止已生成实例，默认false
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type StopOpsTasksAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 是否终止已生成实例，默认false
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *StopOpsTasksAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOpsTasksAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopOpsTasksAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopOpsTasksAsyncResponseParams struct {
	// AsyncId
	Data *OpsAsyncResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopOpsTasksAsyncResponse struct {
	*tchttp.BaseResponse
	Response *StopOpsTasksAsyncResponseParams `json:"Response"`
}

func (r *StopOpsTasksAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopOpsTasksAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSQLScriptRunRequestParams struct {
	// 查询id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type StopSQLScriptRunRequest struct {
	*tchttp.BaseRequest
	
	// 查询id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *StopSQLScriptRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSQLScriptRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSQLScriptRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSQLScriptRunResponseParams struct {
	// 执行结果
	Data *SQLStopResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopSQLScriptRunResponse struct {
	*tchttp.BaseResponse
	Response *StopSQLScriptRunResponseParams `json:"Response"`
}

func (r *StopSQLScriptRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSQLScriptRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`
}

type SubmitTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`
}

func (r *SubmitTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "VersionRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskResponseParams struct {
	// 成功或者失败
	Data *SubmitTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTaskResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTaskResponseParams `json:"Response"`
}

func (r *SubmitTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitTaskResult struct {
	// 生成的任务版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 提交状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type SystemRole struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色展示名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleDisplayName *string `json:"RoleDisplayName,omitnil,omitempty" name:"RoleDisplayName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type TableInfo struct {
	// 数据表GUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guid *string `json:"Guid,omitnil,omitempty" name:"Guid"`

	// 数据表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据表描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库Schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 表的技术元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TechnicalMetadata *TechnicalMetadata `json:"TechnicalMetadata,omitnil,omitempty" name:"TechnicalMetadata"`

	// 表的业务元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessMetadata *BusinessMetadata `json:"BusinessMetadata,omitnil,omitempty" name:"BusinessMetadata"`
}

type Task struct {
	// 任务基本属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskBaseAttribute *TaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// 任务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskConfiguration *TaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// 任务调度配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSchedulerConfiguration *TaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

type TaskBaseAttribute struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型ID：
	// 
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 最近一次保存版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskLatestVersionNo *string `json:"TaskLatestVersionNo,omitnil,omitempty" name:"TaskLatestVersionNo"`

	// 最近一次提交版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskLatestSubmitVersionNo *string `json:"TaskLatestSubmitVersionNo,omitnil,omitempty" name:"TaskLatestSubmitVersionNo"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 任务状态：
	// * N: 新建
	// * Y: 调度中
	// * F: 已下线
	// * O: 已暂停
	// * T: 下线中
	// * INVALID: 已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务最新提交状态，任务是否已经提交：true/false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// 任务创建时间，示例：2022-02-12 11:13:41
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后更新时间，示例：2025-08-13 16:34:06
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 最后更新人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateUserName *string `json:"LastUpdateUserName,omitnil,omitempty" name:"LastUpdateUserName"`

	// 最后运维时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOpsTime *string `json:"LastOpsTime,omitnil,omitempty" name:"LastOpsTime"`

	// 最后运维人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOpsUserName *string `json:"LastOpsUserName,omitnil,omitempty" name:"LastOpsUserName"`

	// 任务负责人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`

	// 最近一次更新用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// 创建用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type TaskCode struct {
	// 代码内容
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// 代码文件大小，单位bytes
	CodeFileSize *uint64 `json:"CodeFileSize,omitnil,omitempty" name:"CodeFileSize"`
}

type TaskCodeResult struct {
	// 代码内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeInfo *string `json:"CodeInfo,omitnil,omitempty" name:"CodeInfo"`

	// 代码文件大小，单位KB
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeFileSize *string `json:"CodeFileSize,omitnil,omitempty" name:"CodeFileSize"`
}

type TaskConfiguration struct {
	// 代码内容的Base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeContent *string `json:"CodeContent,omitnil,omitempty" name:"CodeContent"`

	// 任务扩展属性配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskExtConfigurationList []*TaskExtParameter `json:"TaskExtConfigurationList,omitnil,omitempty" name:"TaskExtConfigurationList"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataCluster *string `json:"DataCluster,omitnil,omitempty" name:"DataCluster"`

	// 指定的运行节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 资源池队列名称，需要通过 DescribeProjectClusterQueues 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// 来源数据源ID, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 来源数据源类型, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceType *string `json:"SourceServiceType,omitnil,omitempty" name:"SourceServiceType"`

	// 来源数据源名称, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceName *string `json:"SourceServiceName,omitnil,omitempty" name:"SourceServiceName"`

	// 目标数据源ID, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// 目标数据源类型, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceType *string `json:"TargetServiceType,omitnil,omitempty" name:"TargetServiceType"`

	// 目标数据源名称, 使用 ; 分隔, 需要通过 DescribeDataSourceWithoutInfo 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceName *string `json:"TargetServiceName,omitnil,omitempty" name:"TargetServiceName"`

	// 资源组ID： 需要通过 DescribeNormalSchedulerExecutorGroups 获取 ExecutorGroupId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 资源组名称： 需要通过 DescribeNormalSchedulerExecutorGroups 获取 ExecutorGroupName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 调度参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSchedulingParameterList []*TaskSchedulingParameter `json:"TaskSchedulingParameterList,omitnil,omitempty" name:"TaskSchedulingParameterList"`

	// Bundle使用的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type TaskDataRegistry struct {
	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 分区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionName *string `json:"PartitionName,omitnil,omitempty" name:"PartitionName"`

	// 输入输出表类型
	//       输入流
	//  UPSTREAM,
	//       输出流
	//   DOWNSTREAM;
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataFlowType *string `json:"DataFlowType,omitnil,omitempty" name:"DataFlowType"`

	// 表物理唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TablePhysicalId *string `json:"TablePhysicalId,omitnil,omitempty" name:"TablePhysicalId"`

	// 库唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbGuid *string `json:"DbGuid,omitnil,omitempty" name:"DbGuid"`

	// 表唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGuid *string `json:"TableGuid,omitnil,omitempty" name:"TableGuid"`
}

type TaskDependDto struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务状态:
	// - Y: 运行
	// - F: 停止
	// - O: 冻结
	// - T: 停止中
	// - INVALID: 已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务类型描述
	//  - 20 :  通用数据同步
	//  - 25 :  ETLTaskType
	//  - 26 :  ETLTaskType
	//  - 30 :  python
	//  - 31 :  pyspark
	//  - 34 :  hivesql
	//  - 35 :  shell
	//  - 36 :  sparksql
	//  - 21 :  jdbcsql
	//  - 32 :  dlc
	//  - 33 :  ImpalaTaskType
	//  - 40 :  CDWTaskType
	//  - 41 :  kettle
	//  - 42 :  TCHouse-X
	//  - 43 :  TCHouse-X SQL
	//  - 46 :  dlcsparkTaskType
	//  - 47 :  TiOneMachineLearningTaskType
	//  - 48 :  Trino
	//  - 50 :  DLCPyspark
	//  - 23 :  TencentDistributedSQL
	//  - 39 :  spark
	//  - 92 :  MRTaskType
	//  - 38 :  ShellScript
	//  - 70 :  HiveSQLScrip
	//  - 130 :  分支
	//  - 131 :  归并
	//  - 132 :  Notebook探索
	//  - 133 :  SSH节点
	//  - 134 :  StarRocks
	//  - 137 :  For-each
	//  - 10000 :  自定义业务通用
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// 调度计划展示描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 延迟时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 周期类型：默认为 D
	// 
	// 支持的类型为 
	// 
	// * O: 一次性
	// * Y: 年
	// * M: 月
	// * W: 周
	// * D: 天
	// * H: 小时
	// * I: 分钟
	// * C: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 弹性周期配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 调度初始化策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`

	// crontab表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`
}

type TaskExtParameter struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type TaskInstance struct {
	// 所属项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 实例数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// **实例状态**
	// - WAIT_EVENT: 等待事件
	// - WAIT_UPSTREAM: 等待上游
	// - WAIT_RUN: 等待运行
	// - RUNNING: 运行中
	// - SKIP_RUNNING: 跳过运行
	// - FAILED_RETRY: 失败重试
	// - EXPIRED: 失败
	// - COMPLETED: 成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// **实例类型**
	// 
	// - 0 表示补录类型
	// - 1 表示周期实例
	// - 2 表示非周期实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 负责人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUinList []*string `json:"OwnerUinList,omitnil,omitempty" name:"OwnerUinList"`

	// 累计运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRunNum *uint64 `json:"TotalRunNum,omitnil,omitempty" name:"TotalRunNum"`

	// 任务类型描述
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// **任务周期类型**
	// 支持过滤多个，条件间为 或 的过滤关系
	// * O: ONEOFF_CYCLE
	// * Y: YEAR_CYCLE
	// * M: MONTH_CYCLE
	// * W: WEEK_CYCLE
	// * D: DAY_CYCLE
	// * H: HOUR_CYCLE
	// * I: MINUTE_CYCLE
	// * C: CRONTAB_CYCLE
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 每次运行失败，下发重试次数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *uint64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// **失败重试次数**
	// 再次使用 手动重跑 或 补录实例等方式触发运行时，会被重置为 0 后重新计数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 运行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 运行完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 耗费时间, 单位ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 计划调度时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerTime *string `json:"SchedulerTime,omitnil,omitempty" name:"SchedulerTime"`

	// 实例最近更新时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 执行资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`
}

type TaskInstanceDetail struct {
	// 所属项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// **实例唯一标识**
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// taskType对应的id
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// **任务周期类型**
	// * ONEOFF_CYCLE: 一次性
	// * YEAR_CYCLE: 年
	// * MONTH_CYCLE: 月
	// * WEEK_CYCLE: 周
	// * DAY_CYCLE: 天
	// * HOUR_CYCLE: 小时
	// * MINUTE_CYCLE: 分钟
	// * CRONTAB_CYCLE: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 实例数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// **实例状态**
	// - WAIT_EVENT: 等待事件
	// - WAIT_UPSTREAM: 等待上游
	// - WAIT_RUN: 等待运行
	// - RUNNING: 运行中
	// - SKIP_RUNNING: 跳过运行
	// - FAILED_RETRY: 失败重试
	// - EXPIRED: 失败
	// - COMPLETED: 成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// **实例类型**
	// 
	// - 0 表示补录类型
	// - 1 表示周期实例
	// - 2 表示非周期实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 负责人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUinList []*string `json:"OwnerUinList,omitnil,omitempty" name:"OwnerUinList"`

	// 累计运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRunNum *uint64 `json:"TotalRunNum,omitnil,omitempty" name:"TotalRunNum"`

	// 每次运行失败，下发重试次数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *uint64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// **失败重试次数**
	// 再次使用 手动重跑 或 补录实例等方式触发运行时，会被重置为 0 后重新计数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 耗费时间, 单位ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *int64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 运行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 运行完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 计划调度时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerTime *string `json:"SchedulerTime,omitnil,omitempty" name:"SchedulerTime"`

	// 实例最近更新时间, 时间格式为 yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 执行资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// 简要的任务失败信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`
}

type TaskInstanceExecutions struct {
	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 记录列表
	Items []*InstanceExecution `json:"Items,omitnil,omitempty" name:"Items"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type TaskInstancePage struct {
	// **总条数**
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// **总分页数**
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页条目数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TaskInstance `json:"Items,omitnil,omitempty" name:"Items"`
}

type TaskOpsInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 负责人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 任务状态:
	// - Y: 调度中
	// - F: 已下线
	// - O: 已暂停
	// - T: 下线中
	// - INVALID: 已失效
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名字
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 更新人名称
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// 任务类型Id：
	// * 21:JDBC SQL
	// * 23:TDSQL-PostgreSQL
	// * 26:OfflineSynchronization
	// * 30:Python
	// * 31:PySpark
	// * 33:Impala
	// * 34:Hive SQL
	// * 35:Shell
	// * 36:Spark SQL
	// * 38:Shell Form Mode
	// * 39:Spark
	// * 40:TCHouse-P
	// * 41:Kettle
	// * 42:Tchouse-X
	// * 43:TCHouse-X SQL
	// * 46:DLC Spark
	// * 47:TiOne
	// * 48:Trino
	// * 50:DLC PySpark
	// * 92:MapReduce
	// * 130:Branch Node
	// * 131:Merged Node
	// * 132:Notebook
	// * 133:SSH
	// * 134:StarRocks
	// * 137:For-each
	// * 138:Setats SQL
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务类型描述
	//  - 通用数据同步
	//  - ETLTaskType
	//  - ETLTaskType
	//  - python
	//  - pyspark
	//  - HiveSQLTaskType
	//  - shell
	//  - SparkSQLTaskType
	//  - JDBCSQLTaskType
	//  - DLCTaskType
	//  - ImpalaTaskType
	//  - CDWTaskType
	//  - kettle
	//  - DLCSparkTaskType
	//  - TiOne机器学习
	//  - TrinoTaskType
	//  - DLCPyspark
	//  - spark
	//  - mr
	//  - shell脚本
	//  - hivesql脚本
	//  - 自定义业务通用
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// 任务周期类型
	// * ONEOFF_CYCLE: 一次性
	// * YEAR_CYCLE: 年
	// * MONTH_CYCLE: 月
	// * WEEK_CYCLE: 周
	// * DAY_CYCLE: 天
	// * HOUR_CYCLE: 小时
	// * MINUTE_CYCLE: 分钟
	// * CRONTAB_CYCLE: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 调度描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// 资源组名称
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// 最新调度提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSchedulerCommitTime *string `json:"LastSchedulerCommitTime,omitnil,omitempty" name:"LastSchedulerCommitTime"`

	// 首次执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstRunTime *string `json:"FirstRunTime,omitnil,omitempty" name:"FirstRunTime"`

	// 最近一次提交时间
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

type TaskSchedulerConfiguration struct {
	// 周期类型：支持的类型为
	// 
	// ONEOFF_CYCLE: 一次性
	// YEAR_CYCLE: 年
	// MONTH_CYCLE: 月
	// WEEK_CYCLE: 周
	// DAY_CYCLE: 天
	// HOUR_CYCLE: 小时
	// MINUTE_CYCLE: 分钟
	// CRONTAB_CYCLE: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 时区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// 0 2 3 1,L,2 * ?	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 生效日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 执行时间 左闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间 右闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 调度类型: 0 正常调度 1 空跑调度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleRunType *int64 `json:"ScheduleRunType,omitnil,omitempty" name:"ScheduleRunType"`

	// 日历调度 取值为 0 和 1， 1为打开，0为关闭，默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// 日历调度 日历 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`

	// 日历调度 日历名称, 需要从 DescribeScheduleCalendarPageList 中获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalendarName *string `json:"CalendarName,omitnil,omitempty" name:"CalendarName"`

	// 自依赖, 默认值 serial, 取值为：parallel(并行), serial(串行), orderly(有序)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 上游依赖数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamDependencyConfigList []*DependencyTaskBrief `json:"UpstreamDependencyConfigList,omitnil,omitempty" name:"UpstreamDependencyConfigList"`

	// 下游依赖数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownStreamDependencyConfigList []*DependencyTaskBrief `json:"DownStreamDependencyConfigList,omitnil,omitempty" name:"DownStreamDependencyConfigList"`

	// 事件数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventListenerList []*EventListener `json:"EventListenerList,omitnil,omitempty" name:"EventListenerList"`

	// 任务调度优先级 运行优先级 4高 5中 6低 , 默认:6
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunPriority *uint64 `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// 重试策略 重试等待时间,单位分钟: 默认: 5
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryWait *int64 `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// 重试策略 最大尝试次数, 默认: 4
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetryAttempts *int64 `json:"MaxRetryAttempts,omitnil,omitempty" name:"MaxRetryAttempts"`

	// 超时处理策略 运行耗时超时（单位：分钟）默认为 -1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionTTL *int64 `json:"ExecutionTTL,omitnil,omitempty" name:"ExecutionTTL"`

	// 超时处理策略 等待总时长耗时超时（单位：分钟）默认为 -1
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitExecutionTotalTTL *string `json:"WaitExecutionTotalTTL,omitnil,omitempty" name:"WaitExecutionTotalTTL"`

	// 重跑&补录配置, 默认为 ALL; , ALL 运行成功或失败后皆可重跑或补录, FAILURE 运行成功后不可重跑或补录，运行失败后可重跑或补录, NONE 运行成功或失败后皆不可重跑或补录;
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowRedoType *string `json:"AllowRedoType,omitnil,omitempty" name:"AllowRedoType"`

	// 输出参数数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTaskOutList []*OutTaskParameter `json:"ParamTaskOutList,omitnil,omitempty" name:"ParamTaskOutList"`

	// 输入参数数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTaskInList []*InTaskParameter `json:"ParamTaskInList,omitnil,omitempty" name:"ParamTaskInList"`

	// 产出登记
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskOutputRegistryList []*TaskDataRegistry `json:"TaskOutputRegistryList,omitnil,omitempty" name:"TaskOutputRegistryList"`

	// **实例生成策略**
	// * T_PLUS_0: T+0生成,默认策略
	// * T_PLUS_1: T+1生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`
}

type TaskSchedulingParameter struct {
	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type TaskVersion struct {
	// 保存时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 保存版本id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 版本描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// 审批状态（只有提交版本有）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveStatus *string `json:"ApproveStatus,omitnil,omitempty" name:"ApproveStatus"`

	// 生产状态（只有提交版本有）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 审批人（只有提交版本有）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveUserUin *string `json:"ApproveUserUin,omitnil,omitempty" name:"ApproveUserUin"`
}

type TaskVersionDetail struct {
	// 保存时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionNum *string `json:"VersionNum,omitnil,omitempty" name:"VersionNum"`

	// 版本创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 保存版本Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 版本描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// 审批状态（只有提交版本有）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveStatus *string `json:"ApproveStatus,omitnil,omitempty" name:"ApproveStatus"`

	// 生产状态（只有提交版本有）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveTime *string `json:"ApproveTime,omitnil,omitempty" name:"ApproveTime"`

	// 版本的任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Task *Task `json:"Task,omitnil,omitempty" name:"Task"`

	// 审批人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveUserUin *string `json:"ApproveUserUin,omitnil,omitempty" name:"ApproveUserUin"`
}

type TechnicalMetadata struct {
	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 数据表位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`
}

type TimeOutStrategyInfo struct {
	// 超时告警超时配置：
	// 
	// 1.预计运行耗时超时，2.预计完成时间超时，3.预计等待调度耗时超时，4.预计周期内完成但实际未完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 超时值配置类型
	// 
	// 1--指定值
	// 2--平均值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 超时指定值小时， 默认 为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hour *uint64 `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 超时指定值分钟， 默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *int64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 超时时间对应的时区配置， 如 UTC+7, 默认为UTC+8
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`
}

type TrendData struct {
	// 时间戳
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 指标值
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type UpdateCodeFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件ID，参数值来自CreateCodeFile接口的返回
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// 代码文件配置
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// 代码文件内容
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

type UpdateCodeFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 代码文件ID，参数值来自CreateCodeFile接口的返回
	CodeFileId *string `json:"CodeFileId,omitnil,omitempty" name:"CodeFileId"`

	// 代码文件配置
	CodeFileConfig *CodeFileConfig `json:"CodeFileConfig,omitnil,omitempty" name:"CodeFileConfig"`

	// 代码文件内容
	CodeFileContent *string `json:"CodeFileContent,omitnil,omitempty" name:"CodeFileContent"`
}

func (r *UpdateCodeFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CodeFileId")
	delete(f, "CodeFileConfig")
	delete(f, "CodeFileContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCodeFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCodeFileResponseParams struct {
	// 结果
	Data *CodeFile `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCodeFileResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCodeFileResponseParams `json:"Response"`
}

func (r *UpdateCodeFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCodeFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，参数值来自CreateCodeFolder接口的返回
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type UpdateCodeFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，参数值来自CreateCodeFolder接口的返回
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *UpdateCodeFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCodeFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCodeFolderResponseParams struct {
	// 执行结果
	Data *CodeStudioFolderActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCodeFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCodeFolderResponseParams `json:"Response"`
}

func (r *UpdateCodeFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCodeFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataSourceRequestParams struct {
	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	// 
	// > deployType: 
	// CONNSTR_PUBLICDB(公网实例) 
	// CONNSTR_CVMDB(自建实例)
	// INSTANCE(云实例)
	// 
	// ```
	// mysql: 自建实例
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:mysql://1.1.1.1:1111/database",
	//     "username": "root",
	//     "password": "root",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "MYSQL"
	// }
	// mysql: 云实例
	// {
	//     "instanceid": "cdb-12uxdo5e",
	//     "db": "db",
	//     "region": "ap-shanghai",
	//     "username": "msyql",
	//     "password": "mysql",
	//     "deployType": "INSTANCE",
	//     "type": "TENCENT_MYSQL"
	// }
	// sql_server: 
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:sqlserver://1.1.1.1:223;DatabaseName=database",
	//     "username": "user_1",
	//     "password": "pass_2",
	//     "type": "SQLSERVER"
	// }
	// redis:
	//     redisType:
	//     -NO_ACCOUT(免账号)
	//     -SELF_ACCOUNT(自定义账号)
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username":""
	//     "password": "pass",
	//     "ip": "1.1.1.1",
	//     "port": "6379",
	//     "redisType": "NO_ACCOUT",
	//     "type": "REDIS"
	// }
	// oracle: 
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:oracle:thin:@1.1.1.1:1521:prod",
	//     "username": "oracle",
	//     "password": "pass",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "ORACLE"
	// }
	// mongodb:
	//     advanceParams(自定义参数，会拼接至url后)
	// {
	//     "advanceParams": [
	//         {
	//             "key": "authSource",
	//             "value": "auth"
	//         }
	//     ],
	//     "db": "admin",
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "MONGODB",
	//     "host": "1.1.1.1:9200"
	// }
	// postgresql:
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:postgresql://1.1.1.1:1921/database",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "POSTGRE"
	// }
	// kafka:
	//     authType:
	//         - sasl
	//         - jaas
	//         - sasl_plaintext
	//         - sasl_ssl
	//         - GSSAPI
	//     ssl:
	//         -PLAIN
	//         -GSSAPI
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "host": "1.1.1.1:9092",
	//     "ssl": "GSSAPI",
	//     "authType": "sasl",
	//     "type": "KAFKA",
	//     "principal": "aaaa",
	//     "serviceName": "kafka"
	// }
	// 
	// cos:
	// {
	//     "region": "ap-shanghai",
	//     "deployType": "INSTANCE",
	//     "secretId": "aaaaa",
	//     "secretKey": "sssssss",
	//     "bucket": "aaa",
	//     "type": "COS"
	// }
	// 
	// ```
	ProdConProperties *string `json:"ProdConProperties,omitnil,omitempty" name:"ProdConProperties"`

	// 若项目为标准模式，则此字段必填
	DevConProperties *string `json:"DevConProperties,omitnil,omitempty" name:"DevConProperties"`

	// 生产环境数据源文件上传
	ProdFileUpload *DataSourceFileUpload `json:"ProdFileUpload,omitnil,omitempty" name:"ProdFileUpload"`

	// 开发环境数据源文件上传
	DevFileUpload *DataSourceFileUpload `json:"DevFileUpload,omitnil,omitempty" name:"DevFileUpload"`

	// 数据源展示名，为了可视化查看
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	// 
	// > deployType: 
	// CONNSTR_PUBLICDB(公网实例) 
	// CONNSTR_CVMDB(自建实例)
	// INSTANCE(云实例)
	// 
	// ```
	// mysql: 自建实例
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:mysql://1.1.1.1:1111/database",
	//     "username": "root",
	//     "password": "root",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "MYSQL"
	// }
	// mysql: 云实例
	// {
	//     "instanceid": "cdb-12uxdo5e",
	//     "db": "db",
	//     "region": "ap-shanghai",
	//     "username": "msyql",
	//     "password": "mysql",
	//     "deployType": "INSTANCE",
	//     "type": "TENCENT_MYSQL"
	// }
	// sql_server: 
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:sqlserver://1.1.1.1:223;DatabaseName=database",
	//     "username": "user_1",
	//     "password": "pass_2",
	//     "type": "SQLSERVER"
	// }
	// redis:
	//     redisType:
	//     -NO_ACCOUT(免账号)
	//     -SELF_ACCOUNT(自定义账号)
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username":""
	//     "password": "pass",
	//     "ip": "1.1.1.1",
	//     "port": "6379",
	//     "redisType": "NO_ACCOUT",
	//     "type": "REDIS"
	// }
	// oracle: 
	// {
	//     "deployType": "CONNSTR_CVMDB",
	//     "url": "jdbc:oracle:thin:@1.1.1.1:1521:prod",
	//     "username": "oracle",
	//     "password": "pass",
	//     "region": "ap-shanghai",
	//     "vpcId": "vpc-kprq42yo",
	//     "type": "ORACLE"
	// }
	// mongodb:
	//     advanceParams(自定义参数，会拼接至url后)
	// {
	//     "advanceParams": [
	//         {
	//             "key": "authSource",
	//             "value": "auth"
	//         }
	//     ],
	//     "db": "admin",
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "MONGODB",
	//     "host": "1.1.1.1:9200"
	// }
	// postgresql:
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "url": "jdbc:postgresql://1.1.1.1:1921/database",
	//     "username": "user",
	//     "password": "pass",
	//     "type": "POSTGRE"
	// }
	// kafka:
	//     authType:
	//         - sasl
	//         - jaas
	//         - sasl_plaintext
	//         - sasl_ssl
	//         - GSSAPI
	//     ssl:
	//         -PLAIN
	//         -GSSAPI
	// {
	//     "deployType": "CONNSTR_PUBLICDB",
	//     "host": "1.1.1.1:9092",
	//     "ssl": "GSSAPI",
	//     "authType": "sasl",
	//     "type": "KAFKA",
	//     "principal": "aaaa",
	//     "serviceName": "kafka"
	// }
	// 
	// cos:
	// {
	//     "region": "ap-shanghai",
	//     "deployType": "INSTANCE",
	//     "secretId": "aaaaa",
	//     "secretKey": "sssssss",
	//     "bucket": "aaa",
	//     "type": "COS"
	// }
	// 
	// ```
	ProdConProperties *string `json:"ProdConProperties,omitnil,omitempty" name:"ProdConProperties"`

	// 若项目为标准模式，则此字段必填
	DevConProperties *string `json:"DevConProperties,omitnil,omitempty" name:"DevConProperties"`

	// 生产环境数据源文件上传
	ProdFileUpload *DataSourceFileUpload `json:"ProdFileUpload,omitnil,omitempty" name:"ProdFileUpload"`

	// 开发环境数据源文件上传
	DevFileUpload *DataSourceFileUpload `json:"DevFileUpload,omitnil,omitempty" name:"DevFileUpload"`

	// 数据源展示名，为了可视化查看
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	delete(f, "ProdConProperties")
	delete(f, "DevConProperties")
	delete(f, "ProdFileUpload")
	delete(f, "DevFileUpload")
	delete(f, "DisplayName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataSourceResponseParams struct {
	// 操作是否成功
	Data *DataSourceStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataSourceResponseParams `json:"Response"`
}

func (r *UpdateDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateFolderResult struct {
	// 更新状态,true表示更新成功，false表示更新失败
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateOpsAlarmRuleRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则唯一id，通过接口GetAlarmRule/ListAlarmRule获取
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// 告警规则新的规则名称
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// 监控对象类型, 
	// 任务维度监控： 可按照任务/工作流/项目来配置：1.任务、2.工作流、3.项目（默认为1.任务） 
	// 项目维度监控： 项目整体任务波动告警， 7：项目波动监控告警
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 根据MonitorType 的设置传入不同的业务id，如下1（任务）： MonitorObjectIds为任务id列表2（工作流）： MonitorObjectIds 为工作流id列表(工作流id可从接口ListWorkflows获取)3（项目）： MonitorObjectIds为项目id列表
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// 告警规则监控类型 failure：失败告警 ；overtime：超时告警 success：成功告警; backTrackingOrRerunSuccess: 补录重跑成功告警 backTrackingOrRerunFailure：补录重跑失败告警； 项目波动告警 projectFailureInstanceUpwardFluctuationAlarm： 当天失败实例数向上波动率超过阀值告警； projectFailureInstanceUpwardVolatilityAlarm：当天失败实例数向上波动量超过阀值告警 projectSuccessInstanceDownwardFluctuationAlarm：当天成功实例数向下波动率超过阀值告警； projectSuccessInstanceDownwardVolatilityAlarm： 当天成功实例数向下波动量超过阀值告警 projectFailureInstanceCountAlarm: 当天失败实例数超过阀值告警 projectFailureInstanceProportionAlarm： 当天失败实例数占比超过阀值告警 离线集成任务对账告警： reconciliationFailure： 离线对账任务失败告警 reconciliationOvertime： 离线对账任务运行超时告警 reconciliationMismatch： 数据对账任务不一致条数超过阀值告警
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// 告警规则配置信息 成功告警无需配置；失败告警可以配置首次失败告警或者所有重试失败告警；超时配置需要配置超时类型及超时阀值；项目波动告警需要配置波动率及防抖周期配置
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// 告警规则的启用状态0--禁用1--启用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警级别 1.普通、2.重要、3.紧急
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警接收人配置信息
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// 告警描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateOpsAlarmRuleRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警规则唯一id，通过接口GetAlarmRule/ListAlarmRule获取
	AlarmRuleId *string `json:"AlarmRuleId,omitnil,omitempty" name:"AlarmRuleId"`

	// 告警规则新的规则名称
	AlarmRuleName *string `json:"AlarmRuleName,omitnil,omitempty" name:"AlarmRuleName"`

	// 监控对象类型, 
	// 任务维度监控： 可按照任务/工作流/项目来配置：1.任务、2.工作流、3.项目（默认为1.任务） 
	// 项目维度监控： 项目整体任务波动告警， 7：项目波动监控告警
	MonitorObjectType *int64 `json:"MonitorObjectType,omitnil,omitempty" name:"MonitorObjectType"`

	// 根据MonitorType 的设置传入不同的业务id，如下1（任务）： MonitorObjectIds为任务id列表2（工作流）： MonitorObjectIds 为工作流id列表(工作流id可从接口ListWorkflows获取)3（项目）： MonitorObjectIds为项目id列表
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// 告警规则监控类型 failure：失败告警 ；overtime：超时告警 success：成功告警; backTrackingOrRerunSuccess: 补录重跑成功告警 backTrackingOrRerunFailure：补录重跑失败告警； 项目波动告警 projectFailureInstanceUpwardFluctuationAlarm： 当天失败实例数向上波动率超过阀值告警； projectFailureInstanceUpwardVolatilityAlarm：当天失败实例数向上波动量超过阀值告警 projectSuccessInstanceDownwardFluctuationAlarm：当天成功实例数向下波动率超过阀值告警； projectSuccessInstanceDownwardVolatilityAlarm： 当天成功实例数向下波动量超过阀值告警 projectFailureInstanceCountAlarm: 当天失败实例数超过阀值告警 projectFailureInstanceProportionAlarm： 当天失败实例数占比超过阀值告警 离线集成任务对账告警： reconciliationFailure： 离线对账任务失败告警 reconciliationOvertime： 离线对账任务运行超时告警 reconciliationMismatch： 数据对账任务不一致条数超过阀值告警
	AlarmTypes []*string `json:"AlarmTypes,omitnil,omitempty" name:"AlarmTypes"`

	// 告警规则配置信息 成功告警无需配置；失败告警可以配置首次失败告警或者所有重试失败告警；超时配置需要配置超时类型及超时阀值；项目波动告警需要配置波动率及防抖周期配置
	AlarmRuleDetail *AlarmRuleDetail `json:"AlarmRuleDetail,omitnil,omitempty" name:"AlarmRuleDetail"`

	// 告警规则的启用状态0--禁用1--启用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 告警级别 1.普通、2.重要、3.紧急
	AlarmLevel *int64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警接收人配置信息
	AlarmGroups []*AlarmGroup `json:"AlarmGroups,omitnil,omitempty" name:"AlarmGroups"`

	// 告警描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateOpsAlarmRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsAlarmRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AlarmRuleId")
	delete(f, "AlarmRuleName")
	delete(f, "MonitorObjectType")
	delete(f, "MonitorObjectIds")
	delete(f, "AlarmTypes")
	delete(f, "AlarmRuleDetail")
	delete(f, "Status")
	delete(f, "AlarmLevel")
	delete(f, "AlarmGroups")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOpsAlarmRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOpsAlarmRuleResponseParams struct {
	// 更新结果是否成功
	// true: 更新成功 false：更新失败/未更新
	Data *ModifyAlarmRuleResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOpsAlarmRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOpsAlarmRuleResponseParams `json:"Response"`
}

func (r *UpdateOpsAlarmRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOpsTasksOwnerRequestParams struct {
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务负责人Id
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type UpdateOpsTasksOwnerRequest struct {
	*tchttp.BaseRequest
	
	// 所属项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务负责人Id
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

func (r *UpdateOpsTasksOwnerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsTasksOwnerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOpsTasksOwnerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOpsTasksOwnerResponseParams struct {
	// 操作结果
	Data *UpdateTasksOwner `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOpsTasksOwnerResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOpsTasksOwnerResponseParams `json:"Response"`
}

func (r *UpdateOpsTasksOwnerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOpsTasksOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProjectRequestParams struct {
	// 目标修改的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目显示名称，可以为中文名,需要租户范围内唯一
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 项目负责人id
	ProjectOwnerUin *string `json:"ProjectOwnerUin,omitnil,omitempty" name:"ProjectOwnerUin"`
}

type UpdateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目标修改的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目显示名称，可以为中文名,需要租户范围内唯一
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 项目负责人id
	ProjectOwnerUin *string `json:"ProjectOwnerUin,omitnil,omitempty" name:"ProjectOwnerUin"`
}

func (r *UpdateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "ProjectOwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateProjectResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProjectResponseParams `json:"Response"`
}

func (r *UpdateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件ID,可通过ListResourceFiles接口获取
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// - 上传文件及手填两种方式只能选择其一，如果两者均提供，取值顺序为文件>手填值
	// -  手填值必须是存在的cos路径, /datastudio/resource/ 为固定前缀, projectId 为项目ID,需传入具体值, parentFolderPath为父文件夹路径, name为文件名, 手填值取值示例:
	//      /datastudio/resource/projectId/parentFolderPath/name 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// 资源名称, 尽可能和文件名保持一致
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// bundle客户端ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundle客户端名称
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type UpdateResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源文件ID,可通过ListResourceFiles接口获取
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// - 上传文件及手填两种方式只能选择其一，如果两者均提供，取值顺序为文件>手填值
	// -  手填值必须是存在的cos路径, /datastudio/resource/ 为固定前缀, projectId 为项目ID,需传入具体值, parentFolderPath为父文件夹路径, name为文件名, 手填值取值示例:
	//      /datastudio/resource/projectId/parentFolderPath/name 
	ResourceFile *string `json:"ResourceFile,omitnil,omitempty" name:"ResourceFile"`

	// 资源名称, 尽可能和文件名保持一致
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// bundle客户端ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// bundle客户端名称
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *UpdateResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	delete(f, "ResourceFile")
	delete(f, "ResourceName")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceFileResponseParams struct {
	// 更新状态
	Data *UpdateResourceFileResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *UpdateResourceFileResponseParams `json:"Response"`
}

func (r *UpdateResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceFileResult struct {
	// true
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateResourceFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID, 可通过ListResourceFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type UpdateResourceFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID, 可通过ListResourceFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *UpdateResourceFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateResourceFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceFolderResponseParams struct {
	// 更新文件夹结果，如果更新失败则报错。
	Data *UpdateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateResourceFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateResourceFolderResponseParams `json:"Response"`
}

func (r *UpdateResourceFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceGroupRequestParams struct {
	// 资源组id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// **变更配置(变配、续费、修改资源包数量 不能同时操作), 变配无法修改资源组类型**
	// 
	// 实时集成资源组  
	// - i32c(实时数据同步-16C64G)
	// 
	// 离线集成资源组
	// - integrated(离线数据同步-8C16G)
	// - i16(离线数据同步-8C32G)
	// 
	// 调度资源组
	// - s_test(测试规格)
	// - s_small(基础规格)
	// - s_medium(普及规格)
	// - s_large(专业规格)
	// 
	// 数据服务资源组
	// - ds_t(测试规格)
	// - ds_s(基础规格)
	// - ds_m(普及规格)
	// - ds_l(专业规格)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 续费时长，单位月(变配、续费、修改资源包数量 不能同时操作)
	PurchasePeriod *int64 `json:"PurchasePeriod,omitnil,omitempty" name:"PurchasePeriod"`

	// 增加/减少资源包的数量(变配、续费、修改资源包数量 不能同时操作)
	Number *ResourceNumber `json:"Number,omitnil,omitempty" name:"Number"`

	// 是否自动续费，续费参数PurchasePeriod不为空时可以生效
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitnil,omitempty" name:"AutoRenewEnabled"`
}

type UpdateResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// **变更配置(变配、续费、修改资源包数量 不能同时操作), 变配无法修改资源组类型**
	// 
	// 实时集成资源组  
	// - i32c(实时数据同步-16C64G)
	// 
	// 离线集成资源组
	// - integrated(离线数据同步-8C16G)
	// - i16(离线数据同步-8C32G)
	// 
	// 调度资源组
	// - s_test(测试规格)
	// - s_small(基础规格)
	// - s_medium(普及规格)
	// - s_large(专业规格)
	// 
	// 数据服务资源组
	// - ds_t(测试规格)
	// - ds_s(基础规格)
	// - ds_m(普及规格)
	// - ds_l(专业规格)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 续费时长，单位月(变配、续费、修改资源包数量 不能同时操作)
	PurchasePeriod *int64 `json:"PurchasePeriod,omitnil,omitempty" name:"PurchasePeriod"`

	// 增加/减少资源包的数量(变配、续费、修改资源包数量 不能同时操作)
	Number *ResourceNumber `json:"Number,omitnil,omitempty" name:"Number"`

	// 是否自动续费，续费参数PurchasePeriod不为空时可以生效
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitnil,omitempty" name:"AutoRenewEnabled"`
}

func (r *UpdateResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Type")
	delete(f, "PurchasePeriod")
	delete(f, "Number")
	delete(f, "AutoRenewEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateResourceGroupResponseParams struct {
	// 是否修改成功
	Data *ResourceStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateResourceGroupResponseParams `json:"Response"`
}

func (r *UpdateResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLFolderRequestParams struct {
	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

type UpdateSQLFolderRequest struct {
	*tchttp.BaseRequest
	
	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 权限范围：SHARED, PRIVATE
	AccessScope *string `json:"AccessScope,omitnil,omitempty" name:"AccessScope"`
}

func (r *UpdateSQLFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FolderId")
	delete(f, "FolderName")
	delete(f, "ProjectId")
	delete(f, "AccessScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSQLFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLFolderResponseParams struct {
	// 成功true，失败false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SQLContentActionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSQLFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSQLFolderResponseParams `json:"Response"`
}

func (r *UpdateSQLFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLScriptRequestParams struct {
	// 探索脚本Id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据探索脚本配置
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// 脚本内容, 需要用Base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`
}

type UpdateSQLScriptRequest struct {
	*tchttp.BaseRequest
	
	// 探索脚本Id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据探索脚本配置
	ScriptConfig *SQLScriptConfig `json:"ScriptConfig,omitnil,omitempty" name:"ScriptConfig"`

	// 脚本内容, 需要用Base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`
}

func (r *UpdateSQLScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	delete(f, "ScriptConfig")
	delete(f, "ScriptContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSQLScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSQLScriptResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SQLScript `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSQLScriptResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSQLScriptResponseParams `json:"Response"`
}

func (r *UpdateSQLScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSQLScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskBaseAttribute struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务负责人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 任务描述
	TaskDescription *string `json:"TaskDescription,omitnil,omitempty" name:"TaskDescription"`
}

type UpdateTaskBrief struct {
	// 任务基本属性
	TaskBaseAttribute *UpdateTaskBaseAttribute `json:"TaskBaseAttribute,omitnil,omitempty" name:"TaskBaseAttribute"`

	// 任务配置
	TaskConfiguration *TaskConfiguration `json:"TaskConfiguration,omitnil,omitempty" name:"TaskConfiguration"`

	// 任务调度配置
	TaskSchedulerConfiguration *TaskSchedulerConfiguration `json:"TaskSchedulerConfiguration,omitnil,omitempty" name:"TaskSchedulerConfiguration"`
}

// Predefined struct for user
type UpdateTaskRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务基本属性
	Task *UpdateTaskBrief `json:"Task,omitnil,omitempty" name:"Task"`
}

type UpdateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务基本属性
	Task *UpdateTaskBrief `json:"Task,omitnil,omitempty" name:"Task"`
}

func (r *UpdateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "Task")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTaskResponseParams struct {
	// 任务ID
	Data *UpdateTaskResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTaskResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTaskResponseParams `json:"Response"`
}

func (r *UpdateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTaskResult struct {
	// 处理结果，成功返回 true，不成功返回 false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateTasksOwner struct {
	// 修改任务责任人结果
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UpdateWorkflowFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，可通过ListWorkflowFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 更新后的文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

type UpdateWorkflowFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID，可通过ListWorkflowFolders接口获取
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 更新后的文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`
}

func (r *UpdateWorkflowFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	delete(f, "FolderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkflowFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowFolderResponseParams struct {
	// 更新文件夹结果，如果更新失败则报错。
	Data *UpdateFolderResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkflowFolderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkflowFolderResponseParams `json:"Response"`
}

func (r *UpdateWorkflowFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 责任人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 统一调度参数
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId项
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle信息
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type UpdateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 责任人ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 统一调度参数
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfigurationInfo `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// BundleId项
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Bundle信息
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

func (r *UpdateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "WorkflowName")
	delete(f, "OwnerUin")
	delete(f, "WorkflowDesc")
	delete(f, "WorkflowParams")
	delete(f, "WorkflowSchedulerConfiguration")
	delete(f, "BundleId")
	delete(f, "BundleInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowResponseParams struct {
	// true代表成功，false代表失败
	Data *UpdateWorkflowResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkflowResponseParams `json:"Response"`
}

func (r *UpdateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateWorkflowResult struct {
	// 更新工作流结果
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type WorkflowDetail struct {
	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 责任人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`

	// 工作流类型，cycle和manual
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 工作流参数数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 统一调度参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowSchedulerConfiguration *WorkflowSchedulerConfiguration `json:"WorkflowSchedulerConfiguration,omitnil,omitempty" name:"WorkflowSchedulerConfiguration"`

	// 工作流描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流所属路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// BundleId项
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// BundleInfo项
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleInfo *string `json:"BundleInfo,omitnil,omitempty" name:"BundleInfo"`
}

type WorkflowFolder struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹绝对路径
	FolderPath *string `json:"FolderPath,omitnil,omitempty" name:"FolderPath"`

	// 创建人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowFolderPage struct {
	// 数据页数，大于等于1
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示的数据条数，最小为10条，最大为200 条
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 文件夹总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 文件夹列表
	Items []*WorkflowFolder `json:"Items,omitnil,omitempty" name:"Items"`
}

type WorkflowInfo struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流类型，cycle及manual
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowType *string `json:"WorkflowType,omitnil,omitempty" name:"WorkflowType"`

	// 负责人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最新修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 最后更新人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserUin *string `json:"UpdateUserUin,omitnil,omitempty" name:"UpdateUserUin"`

	// 工作流描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 创建人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUserUin *string `json:"CreateUserUin,omitnil,omitempty" name:"CreateUserUin"`
}

type WorkflowSchedulerConfiguration struct {
	// 时区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// 周期类型：支持的类型为
	// ONEOFF_CYCLE: 一次性
	// YEAR_CYCLE: 年
	// MONTH_CYCLE: 月
	// WEEK_CYCLE: 周
	// DAY_CYCLE: 天
	// HOUR_CYCLE: 小时
	// MINUTE_CYCLE: 分钟
	// CRONTAB_CYCLE: crontab表达式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 自依赖, 默认值 serial, 取值为：parallel(并行), serial(串行), orderly(有序)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 生效开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 生效结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 工作流依赖，yes or no
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 执行时间左闭区间，示例：00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，示例：23:59
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// cron表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 是否开启日历调度 1 开启 0关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// 日历名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalendarName *string `json:"CalendarName,omitnil,omitempty" name:"CalendarName"`

	// 日历id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`
}

type WorkflowSchedulerConfigurationInfo struct {
	// 时区
	ScheduleTimeZone *string `json:"ScheduleTimeZone,omitnil,omitempty" name:"ScheduleTimeZone"`

	// 周期类型：支持的类型为
	// ONEOFF_CYCLE: 一次性
	// YEAR_CYCLE: 年
	// MONTH_CYCLE: 月
	// WEEK_CYCLE: 周
	// DAY_CYCLE: 天
	// HOUR_CYCLE: 小时
	// MINUTE_CYCLE: 分钟
	// CRONTAB_CYCLE: crontab表达式类型
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 自依赖, 默认值 serial, 取值为：parallel(并行), serial(串行), orderly(有序)
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 生效开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 生效结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// cron表达式
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 工作流依赖，yes or no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 0：不修改 1：将任务的上游依赖配置改为默认值
	ModifyCycleValue *string `json:"ModifyCycleValue,omitnil,omitempty" name:"ModifyCycleValue"`

	// 工作流存在跨工作流依赖且使用cron表达式调度。如果保存统一调度，会断开不支持的依赖关系
	ClearLink *bool `json:"ClearLink,omitnil,omitempty" name:"ClearLink"`

	// ModifyCycleValue为1的时候生效，表示默认修改的上游依赖-时间维度，取值为：
	// * CRONTAB
	// * DAY
	// * HOUR
	// * LIST_DAY
	// * LIST_HOUR
	// * LIST_MINUTE
	// * MINUTE
	// * MONTH
	// * RANGE_DAY
	// * RANGE_HOUR
	// * RANGE_MINUTE
	// * WEEK
	// * YEAR
	// 
	// https://capi.woa.com/object/detail?product=wedata&env=api_dev&version=2025-08-06&name=WorkflowSchedulerConfigurationInfo
	MainCyclicConfig *string `json:"MainCyclicConfig,omitnil,omitempty" name:"MainCyclicConfig"`

	// ModifyCycleValue为1的时候生效，表示默认修改的上游依赖-实例范围
	// 取值为：
	// * ALL_DAY_OF_YEAR
	// * ALL_MONTH_OF_YEAR
	// * CURRENT
	// * CURRENT_DAY
	// * CURRENT_HOUR
	// * CURRENT_MINUTE
	// * CURRENT_MONTH
	// * CURRENT_WEEK
	// * CURRENT_YEAR
	// * PREVIOUS_BEGIN_OF_MONTH
	// * PREVIOUS_DAY
	// * PREVIOUS_DAY_LATER_OFFSET_HOUR
	// * PREVIOUS_DAY_LATER_OFFSET_MINUTE
	// * PREVIOUS_END_OF_MONTH
	// * PREVIOUS_FRIDAY
	// * PREVIOUS_HOUR
	// * PREVIOUS_HOUR_CYCLE
	// * PREVIOUS_HOUR_LATER_OFFSET_MINUTE
	// * PREVIOUS_MINUTE_CYCLE
	// * PREVIOUS_MONTH
	// * PREVIOUS_WEEK
	// * PREVIOUS_WEEKEND
	// * RECENT_DATE
	// 
	// https://capi.woa.com/object/detail?product=wedata&env=api_dev&version=2025-08-06&name=WorkflowSchedulerConfigurationInfo
	SubordinateCyclicConfig *string `json:"SubordinateCyclicConfig,omitnil,omitempty" name:"SubordinateCyclicConfig"`

	// 执行时间左闭区间，示例：00:00，只有周期类型为MINUTE_CYCLE才需要填入
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，示例：23:59，只有周期类型为MINUTE_CYCLE才需要填入
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 是否开启日历调度 1 开启 0关闭
	CalendarOpen *string `json:"CalendarOpen,omitnil,omitempty" name:"CalendarOpen"`

	// 日历id
	CalendarId *string `json:"CalendarId,omitnil,omitempty" name:"CalendarId"`
}