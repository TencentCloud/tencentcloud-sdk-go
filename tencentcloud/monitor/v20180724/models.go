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

package v20180724

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlarmEvent struct {

	// 事件名
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 展示的事件名
	Description *string `json:"Description,omitempty" name:"Description"`

	// 告警策略类型
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type AlarmHistory struct {

	// 告警历史Id
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 策略类型
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 告警对象
	AlarmObject *string `json:"AlarmObject,omitempty" name:"AlarmObject"`

	// 告警内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 时间戳，首次出现时间
	FirstOccurTime *int64 `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// 时间戳，最后出现时间
	LastOccurTime *int64 `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// 告警状态，ALARM=未恢复 OK=已恢复 NO_CONF=已失效 NO_DATA=数据不足
	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 告警策略 Id
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 基础产品告警的告警对象所属网络
	VPC *string `json:"VPC,omitempty" name:"VPC"`

	// 项目 Id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名字
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 告警对象所属实例组
	InstanceGroup []*InstanceGroups `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

	// 接收人列表
	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids"`

	// 接收组列表
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`

	// 告警渠道列表 SMS=短信 EMAIL=邮件 CALL=电话 WECHAT=微信
	NoticeWays []*string `json:"NoticeWays,omitempty" name:"NoticeWays"`

	// 可用于实例、实例组的绑定和解绑接口（[BindingPolicyObject](https://cloud.tencent.com/document/product/248/40421)、[UnBindingAllPolicyObject](https://cloud.tencent.com/document/product/248/40568)、[UnBindingPolicyObject](https://cloud.tencent.com/document/product/248/40567)）的策略 ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 告警类型
	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`

	// 事件Id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 策略是否存在 0=不存在 1=存在
	PolicyExists *int64 `json:"PolicyExists,omitempty" name:"PolicyExists"`

	// 指标信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricsInfo []*AlarmHistoryMetric `json:"MetricsInfo,omitempty" name:"MetricsInfo"`

	// 告警实例的维度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type AlarmHistoryMetric struct {

	// 云产品监控类型查询数据使用的命名空间
	QceNamespace *string `json:"QceNamespace,omitempty" name:"QceNamespace"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 统计周期
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 触发告警的数值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 指标的展示名
	Description *string `json:"Description,omitempty" name:"Description"`
}

type AlarmNotice struct {

	// 告警通知模板 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 告警通知模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 上次修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 上次修改人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`

	// 告警通知类型 ALARM=未恢复通知 OK=已恢复通知 ALL=全部通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`

	// 用户通知列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices"`

	// 回调通知列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`

	// 是否是系统预设通知模板 0=否 1=是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPreset *int64 `json:"IsPreset,omitempty" name:"IsPreset"`

	// 通知语言 zh-CN=中文 en-US=英文
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`

	// 告警通知模板绑定的告警策略ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

type AlarmPolicy struct {

	// 告警策略 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 监控类型 MT_QCE=云产品监控
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 启停状态 0=停用 1=启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 策略组绑定的实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`

	// 项目 Id -1=无项目 0=默认项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 告警策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 触发条件模板 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionTemplateId *string `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`

	// 指标触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// 事件触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`

	// 通知规则 id 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 通知规则 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notices []*AlarmNotice `json:"Notices,omitempty" name:"Notices"`

	// 触发任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`

	// 模板策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionsTemp *ConditionsTemp `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

	// 最后编辑的用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*string `json:"Region,omitempty" name:"Region"`

	// namespace显示名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" name:"NamespaceShowName"`

	// 是否默认策略，1是，0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 能否设置默认策略，1是，0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanSetDefault *int64 `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// 实例分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 实例分组总实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 实例分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`

	// 触发条件类型 STATIC=静态阈值 DYNAMIC=动态类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagInstances []*TagInstance `json:"TagInstances,omitempty" name:"TagInstances"`
}

type AlarmPolicyCondition struct {

	// 指标触发与或条件，0=或，1=与
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 告警触发条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*AlarmPolicyRule `json:"Rules,omitempty" name:"Rules"`
}

type AlarmPolicyEventCondition struct {

	// 告警触发条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*AlarmPolicyRule `json:"Rules,omitempty" name:"Rules"`
}

type AlarmPolicyFilter struct {

	// 过滤条件类型 DIMENSION=使用 Dimensions 做过滤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// AlarmPolicyDimension 二维数组序列化后的json字符串，一维数组之间互为或关系，一维数组内的元素互为与关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type AlarmPolicyRule struct {

	// 指标名或事件名，支持的指标可以从 [DescribeAlarmMetrics](https://cloud.tencent.com/document/product/248/51283) 查询，支持的事件可以从 [DescribeAlarmEvents](https://cloud.tencent.com/document/product/248/51284) 查询 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 秒数 统计周期，支持的值可以从 [DescribeAlarmMetrics](https://cloud.tencent.com/document/product/248/51283) 查询。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 英文运算符
	// intelligent=无阈值智能检测
	// eq=等于
	// ge=大于等于
	// gt=大于
	// le=小于等于
	// lt=小于
	// ne=不等于
	// day_increase=天同比增长
	// day_decrease=天同比下降
	// day_wave=天同比波动
	// week_increase=周同比增长
	// week_decrease=周同比下降
	// week_wave=周同比波动
	// cycle_increase=环比增长
	// cycle_decrease=环比下降
	// cycle_wave=环比波动
	// re=正则匹配
	// 支持的值可以从 [DescribeAlarmMetrics](https://cloud.tencent.com/document/product/248/51283) 查询。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 阈值，支持的范围可以从 [DescribeAlarmMetrics](https://cloud.tencent.com/document/product/248/51283) 查询。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 周期数 持续通知周期 1=持续1个周期 2=持续2个周期...，支持的值可以从 [DescribeAlarmMetrics](https://cloud.tencent.com/document/product/248/51283) 查询
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// 秒数 告警间隔  0=不重复 300=每5分钟告警一次 600=每10分钟告警一次 900=每15分钟告警一次 1800=每30分钟告警一次 3600=每1小时告警一次 7200=每2小时告警一次 10800=每3小时告警一次 21600=每6小时告警一次 43200=每12小时告警一次 86400=每1天告警一次
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeFrequency *int64 `json:"NoticeFrequency,omitempty" name:"NoticeFrequency"`

	// 告警频率是否指数增长 0=否 1=是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPowerNotice *int64 `json:"IsPowerNotice,omitempty" name:"IsPowerNotice"`

	// 对于单个触发规则的过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filter *AlarmPolicyFilter `json:"Filter,omitempty" name:"Filter"`

	// 指标展示名，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 单位，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 触发条件类型 STATIC=静态阈值 DYNAMIC=动态阈值。创建或编辑策略时，如不填则默认为 STATIC。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

type AlarmPolicyTriggerTask struct {

	// 触发任务类型 AS=弹性伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 用 json 表示配置信息 {"Key1":"Value1","Key2":"Value2"}
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskConfig *string `json:"TaskConfig,omitempty" name:"TaskConfig"`
}

type BindingPolicyObjectDimension struct {

	// 地域名
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 实例的维度信息，格式为
	// {"unInstanceId":"ins-00jvv9mo"}。不同云产品的维度信息不同，详见
	// [指标维度信息Dimensions列表](https://cloud.tencent.com/document/product/248/50397)
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 事件维度信息
	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type BindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 必填。固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，如传入 PolicyId 则该字段会被忽略可传入任意值如 0
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警策略ID，使用此字段时 GroupId 会被忽略
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例分组ID
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 需要绑定的对象维度信息
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *BindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindingPolicyObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "PolicyId")
	delete(f, "InstanceGroupId")
	delete(f, "Dimensions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindingPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindingPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonNamespace struct {

	// 命名空间标示
	Id *string `json:"Id,omitempty" name:"Id"`

	// 命名空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 配置信息
	Config *string `json:"Config,omitempty" name:"Config"`

	// 支持地域列表
	AvailableRegions []*string `json:"AvailableRegions,omitempty" name:"AvailableRegions"`

	// 排序Id
	SortId *int64 `json:"SortId,omitempty" name:"SortId"`

	// Dashboard中的唯一表示
	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

type ConditionsTemp struct {

	// 模版名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// 指标触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// 事件触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 通知模板名称 60字符以内
	Name *string `json:"Name,omitempty" name:"Name"`

	// 通知类型 ALARM=未恢复通知 OK=已恢复通知 ALL=都通知
	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`

	// 通知语言 zh-CN=中文 en-US=英文
	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`

	// 用户通知 最多5个
	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices"`

	// 回调通知 最多3个
	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Name")
	delete(f, "NoticeType")
	delete(f, "NoticeLanguage")
	delete(f, "UserNotices")
	delete(f, "URLNotices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警通知模板ID
		NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略名称，不超过20字符
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 监控类型 MT_QCE=云产品监控
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 告警策略类型，由 [DescribeAllNamespaces](https://cloud.tencent.com/document/product/248/48683) 获得。对于云产品监控，取接口出参的 QceNamespacesNew.N.Id，例如 cvm_device
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 备注，不超过100字符，仅支持中英文、数字、下划线、-
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否启用 0=停用 1=启用，可不传 默认为1
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 项目 Id，对于区分项目的产品必须传入非 -1 的值。 -1=无项目 0=默认项目，如不传 默认为 -1。支持的项目 Id 可以在控制台 [账号中心-项目管理](https://console.cloud.tencent.com/project) 中查看。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 触发条件模板 Id ，可不传
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`

	// 指标触发条件，支持的指标可以从 [DescribeAlarmMetrics](https://cloud.tencent.com/document/product/248/51283) 查询。
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// 事件触发条件，支持的事件可以从 [DescribeAlarmEvents](https://cloud.tencent.com/document/product/248/51284) 查询。
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`

	// 通知规则 Id 列表，由 [DescribeAlarmNotices](https://cloud.tencent.com/document/product/248/51280) 获得
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 触发任务列表
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`

	// 全局过滤条件
	Filter *AlarmPolicyFilter `json:"Filter,omitempty" name:"Filter"`

	// 聚合维度列表，指定按哪些维度 key 来做 group by
	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *CreateAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyName")
	delete(f, "MonitorType")
	delete(f, "Namespace")
	delete(f, "Remark")
	delete(f, "Enable")
	delete(f, "ProjectId")
	delete(f, "ConditionTemplateId")
	delete(f, "Condition")
	delete(f, "EventCondition")
	delete(f, "NoticeIds")
	delete(f, "TriggerTasks")
	delete(f, "Filter")
	delete(f, "GroupBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略 ID
		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

		// 可用于实例、实例组的绑定和解绑接口（[BindingPolicyObject](https://cloud.tencent.com/document/product/248/40421)、[UnBindingAllPolicyObject](https://cloud.tencent.com/document/product/248/40568)、[UnBindingPolicyObject](https://cloud.tencent.com/document/product/248/40567)）的策略 ID
		OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertRuleRequest struct {
	*tchttp.BaseRequest

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则表达式
	Expr *string `json:"Expr,omitempty" name:"Expr"`

	// 告警通知模板 ID 列表
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 规则状态码，取值如下：
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`

	// 规则报警持续时间
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 标签列表
	Labels []*PrometheusRuleKV `json:"Labels,omitempty" name:"Labels"`

	// 注释列表。
	// 
	// 告警对象和告警消息是 Prometheus Rule Annotations 的特殊字段，需要通过 annotations 来传递，对应的 Key 分别为summary/description。
	Annotations []*PrometheusRuleKV `json:"Annotations,omitempty" name:"Annotations"`

	// 报警策略模板分类
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	delete(f, "Expr")
	delete(f, "Receivers")
	delete(f, "RuleState")
	delete(f, "Duration")
	delete(f, "Labels")
	delete(f, "Annotations")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupCondition struct {

	// 指标Id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等。如果指标有配置默认比较类型值可以不填。
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// 比较的值，如果指标不必须CalcValue可不填
	CalcValue *float64 `json:"CalcValue,omitempty" name:"CalcValue"`

	// 数据聚合周期(单位秒)，若指标有默认值可不填
	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// 持续几个检测周期触发规则会告警
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// 如果通过模版创建，需要传入模版中该指标的对应RuleId
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupEventCondition struct {

	// 告警事件的Id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 如果通过模版创建，需要传入模版中该指标的对应RuleId
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 组策略名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组所属视图的名称，若通过模版创建，可不传入
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组所属项目Id，会进行鉴权操作
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模版策略组Id, 通过模版创建时才需要传
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`

	// 是否屏蔽策略组，0表示不屏蔽，1表示屏蔽。不填默认为0
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// 策略组的备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 插入时间，戳格式为Unix时间戳，不填则按后台处理时间填充
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 策略组中的阈值告警规则
	Conditions []*CreatePolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 策略组中的事件告警规则
	EventConditions []*CreatePolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`

	// 是否为后端调用。当且仅当值为1时，后台拉取策略模版中的规则填充入Conditions以及EventConditions字段
	BackEndCall *int64 `json:"BackEndCall,omitempty" name:"BackEndCall"`

	// 指标告警规则的且或关系，0表示或规则(满足任意规则就告警)，1表示且规则(满足所有规则才告警)
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

func (r *CreatePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Module")
	delete(f, "ViewName")
	delete(f, "ProjectId")
	delete(f, "ConditionTempGroupId")
	delete(f, "IsShielded")
	delete(f, "Remark")
	delete(f, "InsertTime")
	delete(f, "Conditions")
	delete(f, "EventConditions")
	delete(f, "BackEndCall")
	delete(f, "IsUnionRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功的策略组Id
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceDiscoveryRequest struct {
	*tchttp.BaseRequest

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// <li>类型为TKE：对应集成的腾讯云容器服务集群 ID</li>
	KubeClusterId *string `json:"KubeClusterId,omitempty" name:"KubeClusterId"`

	// 用户 Kubernetes 集群类型：
	// <li> 1 = 容器服务集群(TKE) </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 服务发现类型，取值如下：
	// <li> 1 = ServiceMonitor</li>
	// <li> 2 = PodMonitor</li>
	// <li> 3 = JobMonitor</li>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 服务发现配置信息
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

func (r *CreateServiceDiscoveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceDiscoveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeClusterId")
	delete(f, "KubeType")
	delete(f, "Type")
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceDiscoveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功之后，返回对应服务发现信息
		ServiceDiscovery *ServiceDiscoveryItem `json:"ServiceDiscovery,omitempty" name:"ServiceDiscovery"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceDiscoveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceDiscoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataPoint struct {

	// 实例对象维度组合
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`

	// 时间戳数组，表示那些时间点有数据，缺失的时间戳，没有数据点，可以理解为掉点了
	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps"`

	// 监控值数组，该数组和Timestamps一一对应
	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type DeleteAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警通知模板id列表
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
}

func (r *DeleteAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "NoticeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID 列表
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

func (r *DeleteAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlertRulesRequest struct {
	*tchttp.BaseRequest

	// 规则 ID 列表
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteAlertRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlertRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlertRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlertRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlertRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlertRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId []*int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeletePolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePolicyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceDiscoveryRequest struct {
	*tchttp.BaseRequest

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// <li>类型是 TKE，为对应的腾讯云容器服务集群 ID</li>
	KubeClusterId *string `json:"KubeClusterId,omitempty" name:"KubeClusterId"`

	// 用户 Kubernetes 集群类型：
	// <li> 1 = 容器服务集群(TKE) </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 服务发现类型，取值如下：
	// <li> 1 = ServiceMonitor</li>
	// <li> 2 = PodMonitor</li>
	// <li> 3 = PodMonitor</li>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 服务发现配置信息
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

func (r *DeleteServiceDiscoveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceDiscoveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeClusterId")
	delete(f, "KubeType")
	delete(f, "Type")
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceDiscoveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceDiscoveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceDiscoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListAlarms struct {

	// 事件分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessTypeDesc *string `json:"BusinessTypeDesc,omitempty" name:"BusinessTypeDesc"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccidentTypeDesc *string `json:"AccidentTypeDesc,omitempty" name:"AccidentTypeDesc"`

	// 事件分类的ID，1表示服务问题，2表示其他订阅
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessID *int64 `json:"BusinessID,omitempty" name:"BusinessID"`

	// 事件状态的ID，0表示已恢复，1表示未恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventStatus *int64 `json:"EventStatus,omitempty" name:"EventStatus"`

	// 影响的对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`

	// 事件的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 事件发生的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAccidentEventListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前接口取值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 起始时间，默认一天前的时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 根据UpdateTime排序的规则，取值asc或desc
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// 根据OccurTime排序的规则，取值asc或desc（优先根据UpdateTimeOrder排序）
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// 根据事件类型过滤，1表示服务问题，2表示其他订阅
	AccidentType []*int64 `json:"AccidentType,omitempty" name:"AccidentType"`

	// 根据事件过滤，1表示云服务器存储问题，2表示云服务器网络连接问题，3表示云服务器运行异常，202表示运营商网络抖动
	AccidentEvent []*int64 `json:"AccidentEvent,omitempty" name:"AccidentEvent"`

	// 根据事件状态过滤，0表示已恢复，1表示未恢复
	AccidentStatus []*int64 `json:"AccidentStatus,omitempty" name:"AccidentStatus"`

	// 根据事件地域过滤，gz表示广州，sh表示上海等
	AccidentRegion []*string `json:"AccidentRegion,omitempty" name:"AccidentRegion"`

	// 根据影响资源过滤，比如ins-19a06bka
	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`
}

func (r *DescribeAccidentEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccidentEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "UpdateTimeOrder")
	delete(f, "OccurTimeOrder")
	delete(f, "AccidentType")
	delete(f, "AccidentEvent")
	delete(f, "AccidentStatus")
	delete(f, "AccidentRegion")
	delete(f, "AffectResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccidentEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 平台事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Alarms []*DescribeAccidentEventListAlarms `json:"Alarms,omitempty" name:"Alarms"`

		// 平台事件的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccidentEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccidentEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmEventsRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略类型，由 DescribeAllNamespaces 获得，例如 cvm_device
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 监控类型，如 MT_QCE。如果不填默认为 MT_QCE。
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
}

func (r *DescribeAlarmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Namespace")
	delete(f, "MonitorType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警事件列表
		Events []*AlarmEvent `json:"Events,omitempty" name:"Events"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmHistoriesRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 页数，从 1 开始计数，默认 1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页的数量，取值1~100，默认20
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 默认按首次出现时间倒序排列 "ASC"=正序 "DESC"=逆序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 起始时间，默认一天前的时间戳。对应 `FirstOccurTime` 告警首次出现时间，告警历史的 `FirstOccurTime` 晚于 `StartTime` 才可能被搜索到。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间戳。对应 `FirstOccurTime` 告警首次出现时间，告警历史的 `FirstOccurTime` 早于 `EndTime` 才可能被搜索到。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 根据监控类型过滤 不选默认查所有类型 "MT_QCE"=云产品监控
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`

	// 根据告警对象过滤 字符串模糊搜索
	AlarmObject *string `json:"AlarmObject,omitempty" name:"AlarmObject"`

	// 根据告警状态过滤 ALARM=未恢复 OK=已恢复 NO_CONF=已失效 NO_DATA=数据不足，不选默认查所有
	AlarmStatus []*string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 根据项目ID过滤，-1=无项目 0=默认项目
	// 可在此页面查询 [项目管理](https://console.cloud.tencent.com/project)
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 根据实例组ID过滤
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitempty" name:"InstanceGroupIds"`

	// 根据策略类型过滤，策略类型是监控类型之下的概念，在这里两者都需要传入，例如 `[{"MonitorType": "MT_QCE", "Namespace": "cvm_device"}]`
	// 可使用 [查询所有名字空间 DescribeAllNamespaces](https://cloud.tencent.com/document/product/248/48683) 接口查询
	Namespaces []*MonitorTypeNamespace `json:"Namespaces,omitempty" name:"Namespaces"`

	// 根据指标名过滤
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 根据策略名称模糊搜索
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 根据告警内容模糊搜索
	Content *string `json:"Content,omitempty" name:"Content"`

	// 根据接收人搜索，可以使用“访问管理”的 [拉取子用户 ListUsers](https://cloud.tencent.com/document/product/598/34587) 接口获取用户列表 或 [查询子用户 GetUser](https://cloud.tencent.com/document/product/598/34590) 接口查询子用户详情，此处填入返回结果中的 `Uid` 字段
	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids"`

	// 根据接收组搜索，可以使用“访问管理”的 [查询用户组列表 ListGroups](https://cloud.tencent.com/document/product/598/34589) 接口获取用户组列表 或 [列出用户关联的用户组 ListGroupsForUser](https://cloud.tencent.com/document/product/598/34588) 查询某个子用户所在的用户组列表 ，此处填入返回结果中的 `GroupId ` 字段
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`

	// 根据告警策略 Id 列表搜索
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

func (r *DescribeAlarmHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Order")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MonitorTypes")
	delete(f, "AlarmObject")
	delete(f, "AlarmStatus")
	delete(f, "ProjectIds")
	delete(f, "InstanceGroupIds")
	delete(f, "Namespaces")
	delete(f, "MetricNames")
	delete(f, "PolicyName")
	delete(f, "Content")
	delete(f, "ReceiverUids")
	delete(f, "ReceiverGroups")
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 告警历史列表
		Histories []*AlarmHistory `json:"Histories,omitempty" name:"Histories"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmMetricsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控类型过滤 "MT_QCE"=云产品监控
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 告警策略类型，由 DescribeAllNamespaces 获得，例如 cvm_device
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeAlarmMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "MonitorType")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警指标列表
		Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeCallbacksRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmNoticeCallbacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeCallbacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticeCallbacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeCallbacksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警回调通知
	// 注意：此字段可能返回 null，表示取不到有效值。
		URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticeCallbacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeCallbacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警通知模板 id
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`
}

func (r *DescribeAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "NoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警通知模板详细信息
		Notice *AlarmNotice `json:"Notice,omitempty" name:"Notice"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 页码 最小为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小 1～200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 按更新时间排序方式 ASC=正序 DESC=倒序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 主账号 uid 用于创建预设通知
	OwnerUid *int64 `json:"OwnerUid,omitempty" name:"OwnerUid"`

	// 告警通知模板名称 用来模糊搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 根据接收人过滤告警通知模板需要选定通知用户类型 USER=用户 GROUP=用户组 传空=不按接收人过滤
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 接收对象列表
	UserIds []*int64 `json:"UserIds,omitempty" name:"UserIds"`

	// 接收组列表
	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`

	// 根据通知模板 id 过滤，空数组/不传则不过滤
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Order")
	delete(f, "OwnerUid")
	delete(f, "Name")
	delete(f, "ReceiverType")
	delete(f, "UserIds")
	delete(f, "GroupIds")
	delete(f, "NoticeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警通知模板总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 告警通知模板列表
		Notices []*AlarmNotice `json:"Notices,omitempty" name:"Notices"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 页数，从 1 开始计数，默认 1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页的数量，取值1~100，默认20
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 按策略名称模糊搜索
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 根据监控类型过滤 不选默认查所有类型 "MT_QCE"=云产品监控
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`

	// 根据命名空间过滤，不同策略类型的值详见
	// [策略类型列表](https://cloud.tencent.com/document/product/248/50397)
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`

	// 告警对象列表，JSON 字符串。外层数组，对应多个实例，内层为对象的维度。例如“云服务器-基础监控”可写为：
	// `[ {"Dimensions": {"unInstanceId": "ins-qr8d555g"}}, {"Dimensions": {"unInstanceId": "ins-qr8d555h"}} ]`
	// 具体也可以参考下方的示例 2。
	// 
	// 不同云产品参数示例详见 [维度信息Dimensions列表](https://cloud.tencent.com/document/product/248/50397)
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 根据接收人搜索，可以使用“访问管理”的 [拉取子用户 ListUsers](https://cloud.tencent.com/document/product/598/34587) 接口获取用户列表 或 [查询子用户 GetUser](https://cloud.tencent.com/document/product/598/34590) 接口查询子用户详情，此处填入返回结果中的 `Uid` 字段
	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids"`

	// 根据接收组搜索，可以使用“访问管理”的 [查询用户组列表 ListGroups](https://cloud.tencent.com/document/product/598/34589) 接口获取用户组列表 或 [列出用户关联的用户组 ListGroupsForUser](https://cloud.tencent.com/document/product/598/34588) 查询某个子用户所在的用户组列表 ，此处填入返回结果中的 `GroupId ` 字段
	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`

	// 根据默认策略筛选 不传展示全部策略 DEFAULT=展示默认策略 NOT_DEFAULT=展示非默认策略
	PolicyType []*string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 排序字段，例如按照最后修改时间排序，Field: "UpdateTime"
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序顺序：升序：ASC  降序：DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 策略所属项目的id数组，可在此页面查看
	// [项目管理](https://console.cloud.tencent.com/project)
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 通知模版的id列表，可查询通知模版列表获取。
	// 可使用 [查询通知模板列表](https://cloud.tencent.com/document/product/248/51280) 接口查询。
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 根据触发条件筛选 不传展示全部策略 STATIC=展示静态阈值策略 DYNAMIC=展示动态阈值策略
	RuleTypes []*string `json:"RuleTypes,omitempty" name:"RuleTypes"`

	// 告警启停筛选，[1]：启用   [0]：停止，全部[0, 1]
	Enable []*int64 `json:"Enable,omitempty" name:"Enable"`

	// 传 1 查询未配置通知规则的告警策略；不传或传其他数值，查询所有策略。
	NotBindingNoticeRule *int64 `json:"NotBindingNoticeRule,omitempty" name:"NotBindingNoticeRule"`

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
}

func (r *DescribeAlarmPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "PolicyName")
	delete(f, "MonitorTypes")
	delete(f, "Namespaces")
	delete(f, "Dimensions")
	delete(f, "ReceiverUids")
	delete(f, "ReceiverGroups")
	delete(f, "PolicyType")
	delete(f, "Field")
	delete(f, "Order")
	delete(f, "ProjectIds")
	delete(f, "NoticeIds")
	delete(f, "RuleTypes")
	delete(f, "Enable")
	delete(f, "NotBindingNoticeRule")
	delete(f, "InstanceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 策略数组
		Policies []*AlarmPolicy `json:"Policies,omitempty" name:"Policies"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略详情
		Policy *AlarmPolicy `json:"Policy,omitempty" name:"Policy"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertRulesRequest struct {
	*tchttp.BaseRequest

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则状态码，取值如下：
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 报警策略模板分类
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeAlertRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleId")
	delete(f, "RuleState")
	delete(f, "RuleName")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlertRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 报警规则数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 报警规则详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		AlertRuleSet []*PrometheusRuleSet `json:"AlertRuleSet,omitempty" name:"AlertRuleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlertRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllNamespacesRequest struct {
	*tchttp.BaseRequest

	// 根据使用场景过滤 目前仅有"ST_ALARM"=告警类型
	SceneType *string `json:"SceneType,omitempty" name:"SceneType"`

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 根据监控类型过滤 不填默认查所有类型 "MT_QCE"=云产品监控
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`

	// 根据namespace的Id过滤 不填默认查询所有
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
}

func (r *DescribeAllNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SceneType")
	delete(f, "Module")
	delete(f, "MonitorTypes")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云产品的告警策略类型，已废弃
		QceNamespaces *CommonNamespace `json:"QceNamespaces,omitempty" name:"QceNamespaces"`

		// 其他告警策略类型，已废弃
		CustomNamespaces *CommonNamespace `json:"CustomNamespaces,omitempty" name:"CustomNamespaces"`

		// 云产品的告警策略类型
		QceNamespacesNew []*CommonNamespace `json:"QceNamespacesNew,omitempty" name:"QceNamespacesNew"`

		// 其他告警策略类型，暂不支持
		CustomNamespacesNew []*CommonNamespace `json:"CustomNamespacesNew,omitempty" name:"CustomNamespacesNew"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest

	// 业务命名空间，各个云产品的业务命名空间不同。如需获取业务命名空间，请前往各产品监控指标文档，例如云服务器的命名空间，可参见 [云服务器监控指标](https://cloud.tencent.com/document/product/248/6843)
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名，各个云产品的指标名不同。如需获取指标名，请前往各产品监控指标文档，例如云服务器的指标名，可参见 [云服务器监控指标](https://cloud.tencent.com/document/product/248/6843)
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeBaseMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaseMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询得到的指标描述列表
		MetricSet []*MetricSet `json:"MetricSet,omitempty" name:"MetricSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListAlarms struct {

	// 该条告警的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 告警状态ID，0表示未恢复；1表示已恢复；2,3,5表示数据不足；4表示已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 告警状态，ALARM表示未恢复；OK表示已恢复；NO_DATA表示数据不足；NO_CONF表示已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 策略组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 发生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`

	// 持续时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`

	// 告警内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 告警对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`

	// 告警对象ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`

	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// VPC，只有CVM有
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`

	// 指标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 告警类型，0表示指标告警，2表示产品事件告警，3表示平台事件告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 告警对象维度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 通知方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`

	// 所属实例组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroup []*InstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`
}

type DescribeBasicAlarmListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前取值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 起始时间，默认一天前的时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 根据发生时间排序，取值ASC或DESC
	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`

	// 根据项目ID过滤
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 根据策略类型过滤
	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames"`

	// 根据告警状态过滤
	AlarmStatus []*int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`

	// 根据告警对象过滤
	ObjLike *string `json:"ObjLike,omitempty" name:"ObjLike"`

	// 根据实例组ID过滤
	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitempty" name:"InstanceGroupIds"`

	// 根据指标名过滤
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
}

func (r *DescribeBasicAlarmListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicAlarmListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OccurTimeOrder")
	delete(f, "ProjectIds")
	delete(f, "ViewNames")
	delete(f, "AlarmStatus")
	delete(f, "ObjLike")
	delete(f, "InstanceGroupIds")
	delete(f, "MetricNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicAlarmListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Alarms []*DescribeBasicAlarmListAlarms `json:"Alarms,omitempty" name:"Alarms"`

		// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicAlarmListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicAlarmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListDimension struct {

	// 地域id
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域简称
	Region *string `json:"Region,omitempty" name:"Region"`

	// 维度组合json字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 事件维度组合json字符串
	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type DescribeBindingPolicyObjectListInstance struct {

	// 对象唯一id
	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 表示对象实例的维度集合，jsonObj字符串
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 对象是否被屏蔽，0表示未屏蔽，1表示被屏蔽
	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`

	// 对象所在的地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeBindingPolicyObjectListInstanceGroup struct {

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 告警策略类型名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最后编辑uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 实例分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 实例所在的地域集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

type DescribeBindingPolicyObjectListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，如果有形如 policy-xxxx 的 id，请填到 PolicyId 字段中，本字段填 0
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警策略id，形如 policy-xxxx，如果填入，则GroupId可以填0
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 每次返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，从0开始计数，默认0。举例来说，参数 Offset=0&Limit=20 返回第 0 到 19 项，Offset=20&Limit=20 返回第 20 到 39 项，以此类推
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 筛选对象的维度信息
	Dimensions []*DescribeBindingPolicyObjectListDimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DescribeBindingPolicyObjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindingPolicyObjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "PolicyId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Dimensions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindingPolicyObjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 绑定的对象实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*DescribeBindingPolicyObjectListInstance `json:"List,omitempty" name:"List"`

		// 绑定的对象实例总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 未屏蔽的对象实例数
		NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

		// 绑定的实例分组信息，没有绑定实例分组则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
		InstanceGroup *DescribeBindingPolicyObjectListInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindingPolicyObjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindingPolicyObjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorTypesRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeMonitorTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 监控类型，云产品监控为 MT_QCE
		MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListCondition struct {

	// 策略视图名称
	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`

	// 事件告警条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventMetrics []*DescribePolicyConditionListEventMetric `json:"EventMetrics,omitempty" name:"EventMetrics"`

	// 是否支持多地域
	IsSupportMultiRegion *bool `json:"IsSupportMultiRegion,omitempty" name:"IsSupportMultiRegion"`

	// 指标告警条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*DescribePolicyConditionListMetric `json:"Metrics,omitempty" name:"Metrics"`

	// 策略类型名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序id
	SortId *int64 `json:"SortId,omitempty" name:"SortId"`

	// 是否支持默认策略
	SupportDefault *bool `json:"SupportDefault,omitempty" name:"SupportDefault"`

	// 支持该策略类型的地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportRegions []*string `json:"SupportRegions,omitempty" name:"SupportRegions"`
}

type DescribePolicyConditionListConfigManual struct {

	// 检测方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcType *DescribePolicyConditionListConfigManualCalcType `json:"CalcType,omitempty" name:"CalcType"`

	// 检测阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcValue *DescribePolicyConditionListConfigManualCalcValue `json:"CalcValue,omitempty" name:"CalcValue"`

	// 持续时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinueTime *DescribePolicyConditionListConfigManualContinueTime `json:"ContinueTime,omitempty" name:"ContinueTime"`

	// 数据周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *DescribePolicyConditionListConfigManualPeriod `json:"Period,omitempty" name:"Period"`

	// 持续周期个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodNum *DescribePolicyConditionListConfigManualPeriodNum `json:"PeriodNum,omitempty" name:"PeriodNum"`

	// 聚合方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatType *DescribePolicyConditionListConfigManualStatType `json:"StatType,omitempty" name:"StatType"`
}

type DescribePolicyConditionListConfigManualCalcType struct {

	// CalcType 取值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualCalcValue struct {

	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *string `json:"Default,omitempty" name:"Default"`

	// 固定值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fixed *string `json:"Fixed,omitempty" name:"Fixed"`

	// 最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *string `json:"Max,omitempty" name:"Max"`

	// 最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *string `json:"Min,omitempty" name:"Min"`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualContinueTime struct {

	// 默认持续时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// 可选持续时间，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriod struct {

	// 默认周期，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// 可选周期，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualPeriodNum struct {

	// 默认周期数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *int64 `json:"Default,omitempty" name:"Default"`

	// 可选周期数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`

	// 是否必须
	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribePolicyConditionListConfigManualStatType struct {

	// 数据聚合方式，周期5秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	P5 *string `json:"P5,omitempty" name:"P5"`

	// 数据聚合方式，周期10秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	P10 *string `json:"P10,omitempty" name:"P10"`

	// 数据聚合方式，周期1分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P60 *string `json:"P60,omitempty" name:"P60"`

	// 数据聚合方式，周期5分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P300 *string `json:"P300,omitempty" name:"P300"`

	// 数据聚合方式，周期10分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P600 *string `json:"P600,omitempty" name:"P600"`

	// 数据聚合方式，周期30分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	P1800 *string `json:"P1800,omitempty" name:"P1800"`

	// 数据聚合方式，周期1小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	P3600 *string `json:"P3600,omitempty" name:"P3600"`

	// 数据聚合方式，周期1天
	// 注意：此字段可能返回 null，表示取不到有效值。
	P86400 *string `json:"P86400,omitempty" name:"P86400"`
}

type DescribePolicyConditionListEventMetric struct {

	// 事件id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 事件名称
	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`

	// 是否需要恢复
	NeedRecovered *bool `json:"NeedRecovered,omitempty" name:"NeedRecovered"`

	// 事件类型，预留字段，当前固定取值为2
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribePolicyConditionListMetric struct {

	// 指标配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigManual *DescribePolicyConditionListConfigManual `json:"ConfigManual,omitempty" name:"ConfigManual"`

	// 指标id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 指标名称
	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`

	// 指标单位
	MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`
}

type DescribePolicyConditionListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribePolicyConditionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConditionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyConditionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略条件列表
		Conditions []*DescribePolicyConditionListCondition `json:"Conditions,omitempty" name:"Conditions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyConditionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConditionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoCallback struct {

	// 用户回调接口地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 用户回调接口状态，0表示未验证，1表示已验证，2表示存在url但没有通过验证
	ValidFlag *int64 `json:"ValidFlag,omitempty" name:"ValidFlag"`

	// 用户回调接口验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type DescribePolicyGroupInfoCondition struct {

	// 指标名称
	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`

	// 数据聚合周期(单位秒)
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 指标id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 阈值规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 指标单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等，7表示日同比上涨，8表示日同比下降，9表示周同比上涨，10表示周同比下降，11表示周期环比上涨，12表示周期环比下降
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// 检测阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// 持续多长时间触发规则会告警(单位秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinueTime *int64 `json:"ContinueTime,omitempty" name:"ContinueTime"`

	// 告警指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type DescribePolicyGroupInfoConditionTpl struct {

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 策略类型
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 最后编辑的用户uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 是否且规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupInfoEventCondition struct {

	// 事件id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 事件告警规则id
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

	// 事件名称
	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
}

type DescribePolicyGroupInfoReceiverInfo struct {

	// 告警接收组id列表
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList"`

	// 告警接收人id列表
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`

	// 告警时间段开始时间。范围[0,86400)，作为unix时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 告警时间段结束时间。含义同StartTime
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 接收类型。“group”(接收组)或“user”(接收人)
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 告警通知方式。可选 "SMS","SITE","EMAIL","CALL","WECHAT"
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`

	// 电话告警接收者uid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UidList []*int64 `json:"UidList,omitempty" name:"UidList"`

	// 电话告警轮数
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// 电话告警每轮间隔（秒）
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 电话告警对个人间隔（秒）
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 是否需要电话告警触达提示。0不需要，1需要
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor"`

	// 恢复通知方式。可选"SMS"
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`

	// 告警发送语言
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiveLanguage *string `json:"ReceiveLanguage,omitempty" name:"ReceiveLanguage"`
}

type DescribePolicyGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribePolicyGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略组名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 策略组所属的项目id
		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

		// 是否为默认策略，0表示非默认策略，1表示默认策略
		IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

		// 策略类型
		ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

		// 策略说明
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 策略类型名称
		ShowName *string `json:"ShowName,omitempty" name:"ShowName"`

		// 最近编辑的用户uin
		LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

		// 最近编辑时间
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 该策略支持的地域
		Region []*string `json:"Region,omitempty" name:"Region"`

		// 策略类型的维度列表
		DimensionGroup []*string `json:"DimensionGroup,omitempty" name:"DimensionGroup"`

		// 阈值规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		ConditionsConfig []*DescribePolicyGroupInfoCondition `json:"ConditionsConfig,omitempty" name:"ConditionsConfig"`

		// 产品事件规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		EventConfig []*DescribePolicyGroupInfoEventCondition `json:"EventConfig,omitempty" name:"EventConfig"`

		// 用户接收人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`

		// 用户回调信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Callback *DescribePolicyGroupInfoCallback `json:"Callback,omitempty" name:"Callback"`

		// 模板策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
		ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

		// 是否可以设置成默认策略
		CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

		// 是否且规则
	// 注意：此字段可能返回 null，表示取不到有效值。
		IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListGroup struct {

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 是否开启
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`

	// 策略视图名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最近编辑的用户uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 最后修改时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 策略组绑定的实例数
	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`

	// 策略组绑定的未屏蔽实例数
	NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`

	// 是否为默认策略，0表示非默认策略，1表示默认策略
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 是否可以设置成默认策略
	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// 父策略组id
	ParentGroupId *int64 `json:"ParentGroupId,omitempty" name:"ParentGroupId"`

	// 策略组备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 策略组所属项目id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 阈值规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*DescribePolicyGroupInfoCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 产品事件规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventConditions []*DescribePolicyGroupInfoEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`

	// 用户接收人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`

	// 模板策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`

	// 策略组绑定的实例组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroup *DescribePolicyGroupListGroupInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

	// 且或规则标识, 0表示或规则(任意一条规则满足阈值条件就告警), 1表示且规则(所有规则都满足阈值条件才告警)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribePolicyGroupListGroupInstanceGroup struct {

	// 实例分组名称id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 策略类型视图名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 最近编辑的用户uin
	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 实例分组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DescribePolicyGroupListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 分页参数，每页返回的数量，取值1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按策略名搜索
	Like *string `json:"Like,omitempty" name:"Like"`

	// 实例分组id
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 按更新时间排序, asc 或者 desc
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// 项目id列表
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 告警策略类型列表
	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames"`

	// 是否过滤无接收人策略组, 1表示过滤, 0表示不过滤
	FilterUnuseReceiver *int64 `json:"FilterUnuseReceiver,omitempty" name:"FilterUnuseReceiver"`

	// 过滤条件, 接收组列表
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 过滤条件, 接收人列表
	ReceiverUserList []*string `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`

	// 维度组合字段(json字符串), 例如[[{"name":"unInstanceId","value":"ins-6e4b2aaa"}]]
	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 模板策略组id, 多个id用逗号分隔
	ConditionTempGroupId *string `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`

	// 过滤条件, 接收人或者接收组, user表示接收人, group表示接收组
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 过滤条件，告警策略是否已启动或停止
	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
}

func (r *DescribePolicyGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Like")
	delete(f, "InstanceGroupId")
	delete(f, "UpdateTimeOrder")
	delete(f, "ProjectIds")
	delete(f, "ViewNames")
	delete(f, "FilterUnuseReceiver")
	delete(f, "Receivers")
	delete(f, "ReceiverUserList")
	delete(f, "Dimensions")
	delete(f, "ConditionTempGroupId")
	delete(f, "ReceiverType")
	delete(f, "IsOpen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GroupList []*DescribePolicyGroupListGroup `json:"GroupList,omitempty" name:"GroupList"`

		// 策略组总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListDimensions struct {

	// 维度名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeProductEventListEvents struct {

	// 事件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 事件中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`

	// 事件英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`

	// 事件简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 产品中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`

	// 产品英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`

	// 产品简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否支持告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportAlarm *int64 `json:"SupportAlarm,omitempty" name:"SupportAlarm"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*DescribeProductEventListEventsDimensions `json:"Dimensions,omitempty" name:"Dimensions"`

	// 实例对象附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdditionMsg []*DescribeProductEventListEventsDimensions `json:"AdditionMsg,omitempty" name:"AdditionMsg"`

	// 是否配置告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`

	// 策略信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupInfo []*DescribeProductEventListEventsGroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`

	// 显示名称ViewName
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

type DescribeProductEventListEventsDimensions struct {

	// 维度名（英文）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 维度名（中文）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 维度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeProductEventListEventsGroupInfo struct {

	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type DescribeProductEventListOverView struct {

	// 状态变更的事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusChangeAmount *int64 `json:"StatusChangeAmount,omitempty" name:"StatusChangeAmount"`

	// 告警状态未配置的事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnConfigAlarmAmount *int64 `json:"UnConfigAlarmAmount,omitempty" name:"UnConfigAlarmAmount"`

	// 异常事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnNormalEventAmount *int64 `json:"UnNormalEventAmount,omitempty" name:"UnNormalEventAmount"`

	// 未恢复的事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnRecoverAmount *int64 `json:"UnRecoverAmount,omitempty" name:"UnRecoverAmount"`
}

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 产品类型过滤，比如"cvm"表示云服务器
	ProductName []*string `json:"ProductName,omitempty" name:"ProductName"`

	// 事件名称过滤，比如"guest_reboot"表示机器重启
	EventName []*string `json:"EventName,omitempty" name:"EventName"`

	// 影响对象，比如"ins-19708ino"
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维度过滤，比如外网IP:10.0.0.1
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions"`

	// 产品事件地域过滤参数，比如gz，各地域缩写可参见[地域列表](https://cloud.tencent.com/document/product/248/50863)
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// 事件类型过滤，取值范围["status_change","abnormal"]，分别表示状态变更、异常事件
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 事件状态过滤，取值范围["recover","alarm","-"]，分别表示已恢复、未恢复、无状态
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 项目ID过滤
	Project []*string `json:"Project,omitempty" name:"Project"`

	// 告警状态配置过滤，1表示已配置，0表示未配置
	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`

	// 按更新时间排序，ASC表示升序，DESC表示降序，默认DESC
	TimeOrder *string `json:"TimeOrder,omitempty" name:"TimeOrder"`

	// 起始时间，默认一天前的时间戳
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认当前时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 页偏移量，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页返回的数量，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "ProductName")
	delete(f, "EventName")
	delete(f, "InstanceId")
	delete(f, "Dimensions")
	delete(f, "RegionList")
	delete(f, "Type")
	delete(f, "Status")
	delete(f, "Project")
	delete(f, "IsAlarmConfig")
	delete(f, "TimeOrder")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Events []*DescribeProductEventListEvents `json:"Events,omitempty" name:"Events"`

		// 事件统计
		OverView *DescribeProductEventListOverView `json:"OverView,omitempty" name:"OverView"`

		// 事件总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductListRequest struct {
	*tchttp.BaseRequest

	// 固定传值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 排序方式：DESC/ASC（区分大小写），默认值DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 分页查询的偏移量，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询的每页数据量，默认值20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Order")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		ProductList []*ProductSimple `json:"ProductList,omitempty" name:"ProductList"`

		// 产品总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceDiscoveryRequest struct {
	*tchttp.BaseRequest

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// <li>类型是 TKE，为对应的腾讯云容器服务集群 ID</li>
	KubeClusterId *string `json:"KubeClusterId,omitempty" name:"KubeClusterId"`

	// 用户 Kubernetes 集群类型：
	// <li> 1 = 容器服务集群(TKE) </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`
}

func (r *DescribeServiceDiscoveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceDiscoveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeClusterId")
	delete(f, "KubeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceDiscoveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回服务发现列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServiceDiscoverySet []*ServiceDiscoveryItem `json:"ServiceDiscoverySet,omitempty" name:"ServiceDiscoverySet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceDiscoveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceDiscoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDataRequest struct {
	*tchttp.BaseRequest

	// 所属模块，固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 命名空间，目前只支持QCE/TKE
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名列表
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 维度条件，操作符支持=、in
	Conditions []*MidQueryCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 统计粒度。默认取值为300，单位为s
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 起始时间，默认为当前时间，如2020-12-08T19:51:23+08:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认为当前时间，如2020-12-08T19:51:23+08:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 按指定维度groupBy
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys"`
}

func (r *DescribeStatisticDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Namespace")
	delete(f, "MetricNames")
	delete(f, "Conditions")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "GroupBys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStatisticDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计周期
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 监控数据
		Data []*MetricData `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {

	// 实例维度名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DimensionsDesc struct {

	// 维度名数组
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 命名空间，如QCE/CVM。各个云产品的详细命名空间说明请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名称，如CPUUsage，仅支持单指标拉取。各个云产品的详细指标说明请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的指标英文名即为MetricName
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 实例对象的维度组合，格式为key-value键值对形式的集合。不同类型的实例字段完全不同，如CVM为[{"Name":"InstanceId","Value":"ins-j0hk02zo"}]，Ckafka为[{"Name":"instanceId","Value":"ckafka-l49k54dd"}]，COS为[{"Name":"appid","Value":"1258344699"},{"Name":"bucket","Value":"rig-1258344699"}]。各个云产品的维度请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的维度列即为维度组合的key，value为key对应的值。单请求最多支持批量拉取10个实例的监控数据。
	Instances []*Instance `json:"Instances,omitempty" name:"Instances"`

	// 监控统计周期，如60。默认为取值为300，单位为s。每个指标支持的统计周期不一定相同，各个云产品支持的统计周期请参阅各个产品[监控指标](https://cloud.tencent.com/document/product/248/6140)文档，对应的统计周期列即为支持的统计周期。单请求的数据点数限制为1440个。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 起始时间，如2018-09-22T19:51:23+08:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如2018-09-22T20:51:23+08:00，默认为当前时间。 EndTime不能小于StartTime
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	delete(f, "MetricName")
	delete(f, "Instances")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 统计周期
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// 指标名
		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

		// 数据点数组
		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints"`

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 结束时间
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例的维度组合
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

type InstanceGroup struct {

	// 实例组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 实例组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`
}

type InstanceGroups struct {

	// 实例组 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 实例组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type Metric struct {

	// 告警策略类型
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标展示名
	Description *string `json:"Description,omitempty" name:"Description"`

	// 最小值
	Min *float64 `json:"Min,omitempty" name:"Min"`

	// 最大值
	Max *float64 `json:"Max,omitempty" name:"Max"`

	// 维度列表
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`

	// 单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricConfig *MetricConfig `json:"MetricConfig,omitempty" name:"MetricConfig"`
}

type MetricConfig struct {

	// 允许使用的运算符
	Operator []*string `json:"Operator,omitempty" name:"Operator"`

	// 允许配置的数据周期，以秒为单位
	Period []*int64 `json:"Period,omitempty" name:"Period"`

	// 允许配置的持续周期个数
	ContinuePeriod []*int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`
}

type MetricData struct {

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 监控数据点
	Points []*MetricDataPoint `json:"Points,omitempty" name:"Points"`
}

type MetricDataPoint struct {

	// 实例对象维度组合
	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`

	// 数据点列表
	Values []*Point `json:"Values,omitempty" name:"Values"`
}

type MetricDatum struct {

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标的值
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type MetricObjectMeaning struct {

	// 指标英文解释
	En *string `json:"En,omitempty" name:"En"`

	// 指标中文解释
	Zh *string `json:"Zh,omitempty" name:"Zh"`
}

type MetricSet struct {

	// 命名空间，每个云产品会有一个命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标使用的单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标使用的单位
	UnitCname *string `json:"UnitCname,omitempty" name:"UnitCname"`

	// 指标支持的统计周期，单位是秒，如60、300
	Period []*int64 `json:"Period,omitempty" name:"Period"`

	// 统计周期内指标方式
	Periods []*PeriodsSt `json:"Periods,omitempty" name:"Periods"`

	// 统计指标含义解释
	Meaning *MetricObjectMeaning `json:"Meaning,omitempty" name:"Meaning"`

	// 维度描述信息
	Dimensions []*DimensionsDesc `json:"Dimensions,omitempty" name:"Dimensions"`

	// 指标中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricCName *string `json:"MetricCName,omitempty" name:"MetricCName"`

	// 指标英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricEName *string `json:"MetricEName,omitempty" name:"MetricEName"`
}

type MidQueryCondition struct {

	// 维度
	Key *string `json:"Key,omitempty" name:"Key"`

	// 操作符，支持等于(eq)、不等于(ne)，以及in
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 维度值，当Op是eq、ne时，只使用第一个元素
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警通知规则名称 60字符以内
	Name *string `json:"Name,omitempty" name:"Name"`

	// 通知类型 ALARM=未恢复通知 OK=已恢复通知 ALL=都通知
	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`

	// 通知语言 zh-CN=中文 en-US=英文
	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`

	// 告警通知模板 ID
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`

	// 用户通知 最多5个
	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices"`

	// 回调通知 最多3个
	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Name")
	delete(f, "NoticeType")
	delete(f, "NoticeLanguage")
	delete(f, "NoticeId")
	delete(f, "UserNotices")
	delete(f, "URLNotices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyConditionRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 触发条件模板 Id，可不传
	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`

	// 指标触发条件
	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// 事件触发条件
	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`

	// 全局过滤条件
	Filter *AlarmPolicyFilter `json:"Filter,omitempty" name:"Filter"`

	// 聚合维度列表，指定按哪些维度 key 来做 group by
	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
}

func (r *ModifyAlarmPolicyConditionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyConditionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "ConditionTemplateId")
	delete(f, "Condition")
	delete(f, "EventCondition")
	delete(f, "Filter")
	delete(f, "GroupBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyConditionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyConditionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyConditionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyConditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyInfoRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 要修改的字段 NAME=策略名称 REMARK=策略备注
	Key *string `json:"Key,omitempty" name:"Key"`

	// 修改后的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyAlarmPolicyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "Key")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警通知模板 ID 列表
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
}

func (r *ModifyAlarmPolicyNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "NoticeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 启停状态 0=停用 1=启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyAlarmPolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyTasksRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警策略触发任务列表，空数据代表解绑
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`
}

func (r *ModifyAlarmPolicyTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "TriggerTasks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversRequest struct {
	*tchttp.BaseRequest

	// 需要修改接收人的策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 必填。固定为“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 新接收人信息, 没有填写则删除所有接收人
	ReceiverInfos []*ReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`
}

func (r *ModifyAlarmReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmReceiversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Module")
	delete(f, "ReceiverInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmReceiversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupCondition struct {

	// 指标id
	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`

	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等
	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`

	// 检测阈值
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// 检测指标的数据周期
	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`

	// 持续周期个数
	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 规则id，不填表示新增，填写了ruleId表示在已存在的规则基础上进行修改
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyPolicyGroupEventCondition struct {

	// 事件id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 规则id，不填表示新增，填写了ruleId表示在已存在的规则基础上进行修改
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type ModifyPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警类型
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 指标告警条件的且或关系，1表示且告警，所有指标告警条件都达到才告警，0表示或告警，任意指标告警条件达到都告警
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 指标告警条件规则，不填表示删除已有的所有指标告警条件规则
	Conditions []*ModifyPolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 事件告警条件，不填表示删除已有的事件告警条件
	EventConditions []*ModifyPolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`

	// 模板策略组id
	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`
}

func (r *ModifyPolicyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "ViewName")
	delete(f, "GroupName")
	delete(f, "IsUnionRule")
	delete(f, "Conditions")
	delete(f, "EventConditions")
	delete(f, "ConditionTempGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPolicyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略组id
		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTypeNamespace struct {

	// 监控类型
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 策略类型值
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type PeriodsSt struct {

	// 周期
	Period *string `json:"Period,omitempty" name:"Period"`

	// 统计方式
	StatType []*string `json:"StatType,omitempty" name:"StatType"`
}

type Point struct {

	// 该监控数据点生成的时间点
	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 监控数据点的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type ProductSimple struct {

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 产品中文名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductEnName *string `json:"ProductEnName,omitempty" name:"ProductEnName"`
}

type PrometheusRuleKV struct {

	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PrometheusRuleSet struct {

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则状态码
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`

	// 规则类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 规则标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*PrometheusRuleKV `json:"Labels,omitempty" name:"Labels"`

	// 规则注释列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotations []*PrometheusRuleKV `json:"Annotations,omitempty" name:"Annotations"`

	// 规则表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expr *string `json:"Expr,omitempty" name:"Expr"`

	// 规则报警持续时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 报警接收组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 规则运行健康状态，取值如下：
	// <li>unknown 未知状态</li>
	// <li>pending 加载中</li>
	// <li>ok 运行正常</li>
	// <li>err 运行错误</li>
	Health *string `json:"Health,omitempty" name:"Health"`

	// 规则创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 规则更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PutMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 一组指标和数据
	Metrics []*MetricDatum `json:"Metrics,omitempty" name:"Metrics"`

	// 上报时自行指定的 IP
	AnnounceIp *string `json:"AnnounceIp,omitempty" name:"AnnounceIp"`

	// 上报时自行指定的时间戳
	AnnounceTimestamp *uint64 `json:"AnnounceTimestamp,omitempty" name:"AnnounceTimestamp"`

	// 上报时自行指定的 IP 或 产品实例ID
	AnnounceInstance *string `json:"AnnounceInstance,omitempty" name:"AnnounceInstance"`
}

func (r *PutMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Metrics")
	delete(f, "AnnounceIp")
	delete(f, "AnnounceTimestamp")
	delete(f, "AnnounceInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PutMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReceiverInfo struct {

	// 告警时间段开始时间。范围[0,86400)，作为unix时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 告警时间段结束时间。含义同StartTime
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 告警通知方式。可选 "SMS","SITE","EMAIL","CALL","WECHAT"
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`

	// 接收人类型。“group” 或 “user”
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// ReceiverId
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor"`

	// 电话告警接收者uid
	UidList []*int64 `json:"UidList,omitempty" name:"UidList"`

	// 电话告警轮数
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// 电话告警对个人间隔（秒）
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 电话告警每轮间隔（秒）
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 恢复通知方式。可选"SMS"
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`

	// 是否需要电话告警触达提示。0不需要，1需要
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// 接收组列表。通过平台接口查询到的接收组id列表
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList"`

	// 接收人列表。通过平台接口查询到的接收人id列表
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`

	// 告警接收语言，枚举值（zh-CN，en-US）
	ReceiveLanguage *string `json:"ReceiveLanguage,omitempty" name:"ReceiveLanguage"`
}

type SendCustomAlarmMsgRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前取值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 消息策略ID，在云监控自定义消息页面配置
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用户想要发送的自定义消息内容
	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

func (r *SendCustomAlarmMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCustomAlarmMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "Msg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendCustomAlarmMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendCustomAlarmMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCustomAlarmMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCustomAlarmMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceDiscoveryItem struct {

	// 服务发现名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务发现属于的 Namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 服务发现类型: ServiceMonitor/PodMonitor
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// Namespace 选取方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceSelector *string `json:"NamespaceSelector,omitempty" name:"NamespaceSelector"`

	// Label 选取方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Selector *string `json:"Selector,omitempty" name:"Selector"`

	// Endpoints 信息（PodMonitor 不含该参数）
	Endpoints *string `json:"Endpoints,omitempty" name:"Endpoints"`

	// 服务发现对应的配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

type SetDefaultAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *SetDefaultAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDefaultAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInstance struct {

	// 标签Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签Value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 产品类型，如：cvm
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 绑定状态，2：绑定成功，1：绑定中
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingStatus *int64 `json:"BindingStatus,omitempty" name:"BindingStatus"`

	// 标签状态，2：标签存在，1：标签不存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagStatus *int64 `json:"TagStatus,omitempty" name:"TagStatus"`
}

type URLNotice struct {

	// 回调 url（限长256字符）
	// 注意：此字段可能返回 null，表示取不到有效值。
	URL *string `json:"URL,omitempty" name:"URL"`

	// 是否通过验证 0=否 1=是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsValid *int64 `json:"IsValid,omitempty" name:"IsValid"`

	// 验证码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidationCode *string `json:"ValidationCode,omitempty" name:"ValidationCode"`
}

type UnBindingAllPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，如传入 PolicyId 则该字段被忽略可传入任意值如 0
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警策略ID，使用此字段时 GroupId 会被忽略
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *UnBindingAllPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingAllPolicyObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindingAllPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingAllPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingAllPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingAllPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，如传入 PolicyId 则该字段被忽略可传入任意值如 0
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 待删除对象实例的唯一id列表，UniqueId从调用[获取已绑定对象列表接口](https://cloud.tencent.com/document/api/248/40570)的出参的List中得到
	UniqueId []*string `json:"UniqueId,omitempty" name:"UniqueId"`

	// 实例分组id，如果按实例分组删除的话UniqueId参数是无效的
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 告警策略ID，使用此字段时 GroupId 会被忽略
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *UnBindingPolicyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingPolicyObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "GroupId")
	delete(f, "UniqueId")
	delete(f, "InstanceGroupId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindingPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingPolicyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindingPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAlertRuleRequest struct {
	*tchttp.BaseRequest

	// Prometheus 报警规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则状态码，取值如下：
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// 默认状态码为 2 启用。
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`

	// 报警规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 报警规则表达式
	Expr *string `json:"Expr,omitempty" name:"Expr"`

	// 报警规则持续时间
	Duration *string `json:"Duration,omitempty" name:"Duration"`

	// 报警规则接收组列表
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 报警规则标签列表
	Labels []*PrometheusRuleKV `json:"Labels,omitempty" name:"Labels"`

	// 报警规则注释列表。
	// 
	// 告警对象和告警消息是 Prometheus Rule Annotations 的特殊字段，需要通过 annotations 来传递，对应的 Key 分别为summary/description。
	Annotations []*PrometheusRuleKV `json:"Annotations,omitempty" name:"Annotations"`

	// 报警策略模板分类
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *UpdateAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "InstanceId")
	delete(f, "RuleState")
	delete(f, "RuleName")
	delete(f, "Expr")
	delete(f, "Duration")
	delete(f, "Receivers")
	delete(f, "Labels")
	delete(f, "Annotations")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则 ID
		RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAlertRuleStateRequest struct {
	*tchttp.BaseRequest

	// 规则 ID 列表
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则状态码，取值如下：
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// 默认状态码为 2 启用。
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`
}

func (r *UpdateAlertRuleStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	delete(f, "InstanceId")
	delete(f, "RuleState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAlertRuleStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAlertRuleStateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAlertRuleStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlertRuleStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceDiscoveryRequest struct {
	*tchttp.BaseRequest

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// <li>类型是 TKE，为对应的腾讯云容器服务集群 ID</li>
	KubeClusterId *string `json:"KubeClusterId,omitempty" name:"KubeClusterId"`

	// 用户 Kubernetes 集群类型：
	// <li> 1 = 容器服务集群(TKE) </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 服务发现类型，取值如下：
	// <li> 1 = ServiceMonitor</li>
	// <li> 2 = PodMonitor</li>
	// <li> 3 = JobMonitor</li>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 服务发现配置信息
	Yaml *string `json:"Yaml,omitempty" name:"Yaml"`
}

func (r *UpdateServiceDiscoveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceDiscoveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeClusterId")
	delete(f, "KubeType")
	delete(f, "Type")
	delete(f, "Yaml")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateServiceDiscoveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新成功之后，返回对应服务发现的信息
		ServiceDiscovery *ServiceDiscoveryItem `json:"ServiceDiscovery,omitempty" name:"ServiceDiscovery"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateServiceDiscoveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServiceDiscoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserNotice struct {

	// 接收者类型 USER=用户 GROUP=用户组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 通知开始时间 00:00:00 开始的秒数（取值范围0-86399）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 通知结束时间 00:00:00 开始的秒数（取值范围0-86399）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 通知渠道列表 EMAIL=邮件 SMS=短信 CALL=电话 WECHAT=微信
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeWay []*string `json:"NoticeWay,omitempty" name:"NoticeWay"`

	// 用户 uid 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIds []*int64 `json:"UserIds,omitempty" name:"UserIds"`

	// 用户组 group id 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`

	// 电话轮询列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneOrder []*int64 `json:"PhoneOrder,omitempty" name:"PhoneOrder"`

	// 电话轮询次数 （取值范围1-5）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`

	// 单次轮询内拨打间隔 秒数 （取值范围60-900）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`

	// 两次轮询间隔 秒数（取值范围60-900）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`

	// 是否需要触达通知 0=否 1=是
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedPhoneArriveNotice *int64 `json:"NeedPhoneArriveNotice,omitempty" name:"NeedPhoneArriveNotice"`
}
