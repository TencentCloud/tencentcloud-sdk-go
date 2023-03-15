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

type AlarmHierarchicalNotice struct {
	// 通知模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`

	// 通知等级列表，["Remind","Serious"]表示该通知模板仅接收提醒和严重类别的告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	Classification []*string `json:"Classification,omitempty" name:"Classification"`
}

type AlarmHierarchicalValue struct {
	// 提醒等级阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remind *string `json:"Remind,omitempty" name:"Remind"`

	// 警告等级阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warn *string `json:"Warn,omitempty" name:"Warn"`

	// 严重等级阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Serious *string `json:"Serious,omitempty" name:"Serious"`
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

	// 后台 amp consumer id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AMPConsumerId *string `json:"AMPConsumerId,omitempty" name:"AMPConsumerId"`

	// 推送cls渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`

	// 通知模板绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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

	// 策略关联的过滤维度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterDimensionsParam *string `json:"FilterDimensionsParam,omitempty" name:"FilterDimensionsParam"`

	// 是否为一键告警策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOneClick *int64 `json:"IsOneClick,omitempty" name:"IsOneClick"`

	// 一键告警策略是否开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	OneClickStatus *int64 `json:"OneClickStatus,omitempty" name:"OneClickStatus"`

	// 高级指标数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedMetricNumber *int64 `json:"AdvancedMetricNumber,omitempty" name:"AdvancedMetricNumber"`

	// 策略是否是全部对象策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBindAll *int64 `json:"IsBindAll,omitempty" name:"IsBindAll"`

	// 策略标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type AlarmPolicyCondition struct {
	// 告警触发条件的判断方式. 0: 任意; 1: 全部; 2: 复合. 当取值为2的时候为复合告警，与参数 ComplexExpression 配合使用.
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`

	// 告警触发条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*AlarmPolicyRule `json:"Rules,omitempty" name:"Rules"`

	// 复合告警触发条件的判断表达式，当 IsUnionRule 取值为2的时候有效. 其作用是描述多个触发条件需要满足表达式求值为True时才算是满足告警条件.
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplexExpression *string `json:"ComplexExpression,omitempty" name:"ComplexExpression"`
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

	// 是否为高级指标，0否，1是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAdvanced *int64 `json:"IsAdvanced,omitempty" name:"IsAdvanced"`

	// 高级指标是否开通，0否，1是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`

	// 集成中心产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueMax *float64 `json:"ValueMax,omitempty" name:"ValueMax"`

	// 最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueMin *float64 `json:"ValueMin,omitempty" name:"ValueMin"`

	// 告警分级阈值配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HierarchicalValue *AlarmHierarchicalValue `json:"HierarchicalValue,omitempty" name:"HierarchicalValue"`
}

type AlarmPolicyTriggerTask struct {
	// 触发任务类型 AS=弹性伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 用 json 表示配置信息 {"Key1":"Value1","Key2":"Value2"}
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskConfig *string `json:"TaskConfig,omitempty" name:"TaskConfig"`
}

// Predefined struct for user
type BindPrometheusManagedGrafanaRequestParams struct {
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Grafana 可视化服务实例 ID
	GrafanaId *string `json:"GrafanaId,omitempty" name:"GrafanaId"`
}

type BindPrometheusManagedGrafanaRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Grafana 可视化服务实例 ID
	GrafanaId *string `json:"GrafanaId,omitempty" name:"GrafanaId"`
}

func (r *BindPrometheusManagedGrafanaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPrometheusManagedGrafanaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GrafanaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindPrometheusManagedGrafanaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindPrometheusManagedGrafanaResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindPrometheusManagedGrafanaResponse struct {
	*tchttp.BaseResponse
	Response *BindPrometheusManagedGrafanaResponseParams `json:"Response"`
}

func (r *BindPrometheusManagedGrafanaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindPrometheusManagedGrafanaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// Predefined struct for user
type BindingPolicyObjectRequestParams struct {
	// 必填。固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，例如 4739573。逐渐弃用，建议使用 PolicyId 参数
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警策略ID，例如“policy-gh892hg0”。PolicyId 参数与 GroupId 参数至少要填一个，否则会报参数错误，建议使用该参数。若两者同时存在则以该参数为准
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例分组ID
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 需要绑定的对象维度信息
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`

	// 是否配置了事件告警
	EbEventFlag *int64 `json:"EbEventFlag,omitempty" name:"EbEventFlag"`
}

type BindingPolicyObjectRequest struct {
	*tchttp.BaseRequest
	
	// 必填。固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，例如 4739573。逐渐弃用，建议使用 PolicyId 参数
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警策略ID，例如“policy-gh892hg0”。PolicyId 参数与 GroupId 参数至少要填一个，否则会报参数错误，建议使用该参数。若两者同时存在则以该参数为准
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 实例分组ID
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 需要绑定的对象维度信息
	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`

	// 是否配置了事件告警
	EbEventFlag *int64 `json:"EbEventFlag,omitempty" name:"EbEventFlag"`
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
	delete(f, "EbSubject")
	delete(f, "EbEventFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindingPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindingPolicyObjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *BindingPolicyObjectResponseParams `json:"Response"`
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

// Predefined struct for user
type BindingPolicyTagRequestParams struct {
	// 固定取值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 产品类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 策略标签
	Tag *PolicyTag `json:"Tag,omitempty" name:"Tag"`

	// 实例分组ID
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 批量绑定标签
	BatchTag []*PolicyTag `json:"BatchTag,omitempty" name:"BatchTag"`
}

type BindingPolicyTagRequest struct {
	*tchttp.BaseRequest
	
	// 固定取值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 产品类型
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

	// 策略标签
	Tag *PolicyTag `json:"Tag,omitempty" name:"Tag"`

	// 实例分组ID
	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 批量绑定标签
	BatchTag []*PolicyTag `json:"BatchTag,omitempty" name:"BatchTag"`
}

func (r *BindingPolicyTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindingPolicyTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "PolicyId")
	delete(f, "GroupId")
	delete(f, "ServiceType")
	delete(f, "Tag")
	delete(f, "InstanceGroupId")
	delete(f, "BatchTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindingPolicyTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindingPolicyTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindingPolicyTagResponse struct {
	*tchttp.BaseResponse
	Response *BindingPolicyTagResponseParams `json:"Response"`
}

func (r *BindingPolicyTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindingPolicyTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CLSNotice struct {
	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 日志集Id
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 主题Id
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 启停状态，可不传，默认启用。0=停用，1=启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

// Predefined struct for user
type CheckIsPrometheusNewUserRequestParams struct {

}

type CheckIsPrometheusNewUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CheckIsPrometheusNewUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIsPrometheusNewUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIsPrometheusNewUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIsPrometheusNewUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckIsPrometheusNewUserResponse struct {
	*tchttp.BaseResponse
	Response *CheckIsPrometheusNewUserResponseParams `json:"Response"`
}

func (r *CheckIsPrometheusNewUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIsPrometheusNewUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanGrafanaInstanceRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CleanGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CleanGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CleanGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanGrafanaInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CleanGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CleanGrafanaInstanceResponseParams `json:"Response"`
}

func (r *CleanGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanGrafanaInstanceResponse) FromJsonString(s string) error {
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

type CommonNamespaceNew struct {
	// 命名空间标示
	Id *string `json:"Id,omitempty" name:"Id"`

	// 命名空间名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 维度信息
	Dimensions []*DimensionNew `json:"Dimensions,omitempty" name:"Dimensions"`
}

type Condition struct {
	// 告警通知频率
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 重复通知策略预定义（0 - 只告警一次， 1 - 指数告警，2 - 连接告警）
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 检测方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcType *string `json:"CalcType,omitempty" name:"CalcType"`

	// 检测值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`

	// 持续时间，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContinueTime *string `json:"ContinueTime,omitempty" name:"ContinueTime"`

	// 指标ID
	MetricID *int64 `json:"MetricID,omitempty" name:"MetricID"`

	// 指标展示名称（对外）
	MetricDisplayName *string `json:"MetricDisplayName,omitempty" name:"MetricDisplayName"`

	// 周期
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 规则ID
	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`

	// 指标单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 是否为高级指标，0：否；1：是
	IsAdvanced *int64 `json:"IsAdvanced,omitempty" name:"IsAdvanced"`

	// 是否开通高级指标，0：否；1：是
	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`

	// 产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
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

// Predefined struct for user
type CreateAlarmNoticeRequestParams struct {
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

	// 推送CLS日志服务的操作 最多1个
	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`

	// 模板绑定的标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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

	// 推送CLS日志服务的操作 最多1个
	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`

	// 模板绑定的标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "CLSNotices")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmNoticeResponseParams struct {
	// 告警通知模板ID
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmNoticeResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAlarmPolicyRequestParams struct {
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

	// 触发条件模板 Id，该参数与 Condition 参数二选一。如果策略绑定触发条件模板，则传该参数；否则不传该参数，而是传 Condition 参数。触发条件模板 Id 可以从 [DescribeConditionsTemplateList](https://cloud.tencent.com/document/api/248/70250) 接口获取。
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

	// 模板绑定的标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 日志告警信息
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitempty" name:"LogAlarmReqInfo"`

	// 告警分级通知规则配置
	HierarchicalNotices []*AlarmHierarchicalNotice `json:"HierarchicalNotices,omitempty" name:"HierarchicalNotices"`

	// 迁移策略专用字段，0-走鉴权逻辑，1-跳过鉴权逻辑
	MigrateFlag *int64 `json:"MigrateFlag,omitempty" name:"MigrateFlag"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`
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

	// 触发条件模板 Id，该参数与 Condition 参数二选一。如果策略绑定触发条件模板，则传该参数；否则不传该参数，而是传 Condition 参数。触发条件模板 Id 可以从 [DescribeConditionsTemplateList](https://cloud.tencent.com/document/api/248/70250) 接口获取。
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

	// 模板绑定的标签
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 日志告警信息
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitempty" name:"LogAlarmReqInfo"`

	// 告警分级通知规则配置
	HierarchicalNotices []*AlarmHierarchicalNotice `json:"HierarchicalNotices,omitempty" name:"HierarchicalNotices"`

	// 迁移策略专用字段，0-走鉴权逻辑，1-跳过鉴权逻辑
	MigrateFlag *int64 `json:"MigrateFlag,omitempty" name:"MigrateFlag"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`
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
	delete(f, "Tags")
	delete(f, "LogAlarmReqInfo")
	delete(f, "HierarchicalNotices")
	delete(f, "MigrateFlag")
	delete(f, "EbSubject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlarmPolicyResponseParams struct {
	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 可用于实例、实例组的绑定和解绑接口（[BindingPolicyObject](https://cloud.tencent.com/document/product/248/40421)、[UnBindingAllPolicyObject](https://cloud.tencent.com/document/product/248/40568)、[UnBindingPolicyObject](https://cloud.tencent.com/document/product/248/40567)）的策略 ID
	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlarmPolicyResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAlertRuleRequestParams struct {
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则表达式，可参考<a href="https://cloud.tencent.com/document/product/1416/56008">告警规则说明</a>
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

type CreateAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`

	// 规则表达式，可参考<a href="https://cloud.tencent.com/document/product/1416/56008">告警规则说明</a>
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

// Predefined struct for user
type CreateAlertRuleResponseParams struct {
	// 规则 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlertRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateExporterIntegrationRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集成配置
	Content *string `json:"Content,omitempty" name:"Content"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type CreateExporterIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集成配置
	Content *string `json:"Content,omitempty" name:"Content"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateExporterIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExporterIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Content")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExporterIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExporterIntegrationResponseParams struct {
	// 返回创建成功的集成名称列表
	Names []*string `json:"Names,omitempty" name:"Names"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateExporterIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateExporterIntegrationResponseParams `json:"Response"`
}

func (r *CreateExporterIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExporterIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaInstanceRequestParams struct {
	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID 数组
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Grafana 初始密码
	GrafanaInitPassword *string `json:"GrafanaInitPassword,omitempty" name:"GrafanaInitPassword"`

	// 是否启用外网
	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`

	// 标签
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

type CreateGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID 数组
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Grafana 初始密码
	GrafanaInitPassword *string `json:"GrafanaInitPassword,omitempty" name:"GrafanaInitPassword"`

	// 是否启用外网
	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`

	// 标签
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitempty" name:"TagSpecification"`
}

func (r *CreateGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "VpcId")
	delete(f, "SubnetIds")
	delete(f, "GrafanaInitPassword")
	delete(f, "EnableInternet")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaInstanceResponseParams struct {
	// 实例名
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateGrafanaInstanceResponseParams `json:"Response"`
}

func (r *CreateGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaIntegrationRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成类型，可在实例详情-云产品集成-集成列表查看。例如：tencent-cloud-prometheus
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集成配置
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreateGrafanaIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成类型，可在实例详情-云产品集成-集成列表查看。例如：tencent-cloud-prometheus
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集成配置
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *CreateGrafanaIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGrafanaIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaIntegrationResponseParams struct {
	// 集成 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGrafanaIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateGrafanaIntegrationResponseParams `json:"Response"`
}

func (r *CreateGrafanaIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaNotificationChannelRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 默认为1，已废弃，请使用 OrganizationIds
	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`

	// 接受告警通道 ID 数组
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 额外组织 ID 数组，已废弃，请使用 OrganizationIds
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitempty" name:"ExtraOrgIds"`

	// 生效的所有组织 ID 数组，默认为 ["1"]
	OrganizationIds []*string `json:"OrganizationIds,omitempty" name:"OrganizationIds"`
}

type CreateGrafanaNotificationChannelRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 默认为1，已废弃，请使用 OrganizationIds
	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`

	// 接受告警通道 ID 数组
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 额外组织 ID 数组，已废弃，请使用 OrganizationIds
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitempty" name:"ExtraOrgIds"`

	// 生效的所有组织 ID 数组，默认为 ["1"]
	OrganizationIds []*string `json:"OrganizationIds,omitempty" name:"OrganizationIds"`
}

func (r *CreateGrafanaNotificationChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaNotificationChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ChannelName")
	delete(f, "OrgId")
	delete(f, "Receivers")
	delete(f, "ExtraOrgIds")
	delete(f, "OrganizationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGrafanaNotificationChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGrafanaNotificationChannelResponseParams struct {
	// 通道 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateGrafanaNotificationChannelResponse struct {
	*tchttp.BaseResponse
	Response *CreateGrafanaNotificationChannelResponseParams `json:"Response"`
}

func (r *CreateGrafanaNotificationChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGrafanaNotificationChannelResponse) FromJsonString(s string) error {
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

	// 如果通过模板创建，需要传入模板中该指标的对应RuleId
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type CreatePolicyGroupEventCondition struct {
	// 告警事件的Id
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 告警发送收敛类型。0连续告警，1指数告警
	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次
	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 如果通过模板创建，需要传入模板中该指标的对应RuleId
	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

// Predefined struct for user
type CreatePolicyGroupRequestParams struct {
	// 组策略名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组所属视图的名称，若通过模板创建，可不传入
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组所属项目Id，会进行鉴权操作
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模板策略组Id, 通过模板创建时才需要传
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

	// 是否为后端调用。当且仅当值为1时，后台拉取策略模板中的规则填充入Conditions以及EventConditions字段
	BackEndCall *int64 `json:"BackEndCall,omitempty" name:"BackEndCall"`

	// 指标告警规则的且或关系，0表示或规则(满足任意规则就告警)，1表示且规则(满足所有规则才告警)
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type CreatePolicyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 组策略名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组所属视图的名称，若通过模板创建，可不传入
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 策略组所属项目Id，会进行鉴权操作
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模板策略组Id, 通过模板创建时才需要传
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

	// 是否为后端调用。当且仅当值为1时，后台拉取策略模板中的规则填充入Conditions以及EventConditions字段
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

// Predefined struct for user
type CreatePolicyGroupResponseParams struct {
	// 创建成功的策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePrometheusAgentRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent 名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreatePrometheusAgentRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent 名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreatePrometheusAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAgentResponseParams struct {
	// 创建成功的 Agent Id
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAgentResponseParams `json:"Response"`
}

func (r *CreatePrometheusAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitempty" name:"AlertRule"`
}

type CreatePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitempty" name:"AlertRule"`
}

func (r *CreatePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertPolicyResponseParams struct {
	// 告警id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *CreatePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusClusterAgentRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// agent列表
	Agents []*PrometheusClusterAgentBasic `json:"Agents,omitempty" name:"Agents"`
}

type CreatePrometheusClusterAgentRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// agent列表
	Agents []*PrometheusClusterAgentBasic `json:"Agents,omitempty" name:"Agents"`
}

func (r *CreatePrometheusClusterAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusClusterAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Agents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusClusterAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusClusterAgentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusClusterAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusClusterAgentResponseParams `json:"Response"`
}

func (r *CreatePrometheusClusterAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusClusterAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`
}

type CreatePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`
}

func (r *CreatePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusConfigResponseParams `json:"Response"`
}

func (r *CreatePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusGlobalNotificationRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitempty" name:"Notification"`
}

type CreatePrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitempty" name:"Notification"`
}

func (r *CreatePrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Notification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusGlobalNotificationResponseParams struct {
	// 全局告警通知渠道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *CreatePrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusMultiTenantInstancePostPayModeRequestParams struct {
	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据存储时间（单位天），限制值为15，30，45之一
	DataRetentionTime *int64 `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例的标签
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 需要关联的 Grafana 实例
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" name:"GrafanaInstanceId"`
}

type CreatePrometheusMultiTenantInstancePostPayModeRequest struct {
	*tchttp.BaseRequest
	
	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据存储时间（单位天），限制值为15，30，45之一
	DataRetentionTime *int64 `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例的标签
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 需要关联的 Grafana 实例
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" name:"GrafanaInstanceId"`
}

func (r *CreatePrometheusMultiTenantInstancePostPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusMultiTenantInstancePostPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DataRetentionTime")
	delete(f, "Zone")
	delete(f, "TagSpecification")
	delete(f, "GrafanaInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusMultiTenantInstancePostPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusMultiTenantInstancePostPayModeResponseParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusMultiTenantInstancePostPayModeResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusMultiTenantInstancePostPayModeResponseParams `json:"Response"`
}

func (r *CreatePrometheusMultiTenantInstancePostPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusMultiTenantInstancePostPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusRecordRuleYamlRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// yaml的内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreatePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// yaml的内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *CreatePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusRecordRuleYamlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *CreatePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusScrapeJobRequestParams struct {
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID，例如：agent-abcd1234，可在控制台 Agent 管理中获取
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 抓取任务配置，格式：job_name:xx
	Config *string `json:"Config,omitempty" name:"Config"`
}

type CreatePrometheusScrapeJobRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID，例如：agent-abcd1234，可在控制台 Agent 管理中获取
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 抓取任务配置，格式：job_name:xx
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *CreatePrometheusScrapeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusScrapeJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusScrapeJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusScrapeJobResponseParams struct {
	// 成功创建抓取任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusScrapeJobResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusScrapeJobResponseParams `json:"Response"`
}

func (r *CreatePrometheusScrapeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusScrapeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTempRequestParams struct {
	// 模板设置
	Template *PrometheusTemp `json:"Template,omitempty" name:"Template"`
}

type CreatePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模板设置
	Template *PrometheusTemp `json:"Template,omitempty" name:"Template"`
}

func (r *CreatePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTempResponseParams struct {
	// 模板Id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusTempResponseParams `json:"Response"`
}

func (r *CreatePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordingRuleRequestParams struct {
	// 聚合规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 聚合规则组内容，格式为 yaml
	Group *string `json:"Group,omitempty" name:"Group"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则状态码，取值如下：
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// 默认状态码为 2 启用。
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`
}

type CreateRecordingRuleRequest struct {
	*tchttp.BaseRequest
	
	// 聚合规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 聚合规则组内容，格式为 yaml
	Group *string `json:"Group,omitempty" name:"Group"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 规则状态码，取值如下：
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// 默认状态码为 2 启用。
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`
}

func (r *CreateRecordingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Group")
	delete(f, "InstanceId")
	delete(f, "RuleState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordingRuleResponseParams struct {
	// 规则 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRecordingRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordingRuleResponseParams `json:"Response"`
}

func (r *CreateRecordingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSSOAccountRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户账号 ID ，例如：10000000
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 权限
	Role []*GrafanaAccountRole `json:"Role,omitempty" name:"Role"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

type CreateSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户账号 ID ，例如：10000000
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 权限
	Role []*GrafanaAccountRole `json:"Role,omitempty" name:"Role"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

func (r *CreateSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	delete(f, "Role")
	delete(f, "Notes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSSOAccountResponseParams struct {
	// 已添加的用户 UIN
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateSSOAccountResponseParams `json:"Response"`
}

func (r *CreateSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceDiscoveryRequestParams struct {
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

// Predefined struct for user
type CreateServiceDiscoveryResponseParams struct {
	// 创建成功之后，返回对应服务发现信息
	ServiceDiscovery *ServiceDiscoveryItem `json:"ServiceDiscovery,omitempty" name:"ServiceDiscovery"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceDiscoveryResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAlarmNoticesRequestParams struct {
	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警通知模板id列表
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 通知模板与策略绑定关系
	NoticeBindPolicys []*NoticeBindPolicys `json:"NoticeBindPolicys,omitempty" name:"NoticeBindPolicys"`
}

type DeleteAlarmNoticesRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警通知模板id列表
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 通知模板与策略绑定关系
	NoticeBindPolicys []*NoticeBindPolicys `json:"NoticeBindPolicys,omitempty" name:"NoticeBindPolicys"`
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
	delete(f, "NoticeBindPolicys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmNoticesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmNoticesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAlarmPolicyRequestParams struct {
	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID 列表
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
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

// Predefined struct for user
type DeleteAlarmPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmPolicyResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAlertRulesRequestParams struct {
	// 规则 ID 列表
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DeleteAlertRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAlertRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlertRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteExporterIntegrationRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteExporterIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteExporterIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExporterIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Name")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExporterIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteExporterIntegrationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteExporterIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteExporterIntegrationResponseParams `json:"Response"`
}

func (r *DeleteExporterIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExporterIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaInstanceRequestParams struct {
	// 实例名数组
	InstanceIDs []*string `json:"InstanceIDs,omitempty" name:"InstanceIDs"`
}

type DeleteGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例名数组
	InstanceIDs []*string `json:"InstanceIDs,omitempty" name:"InstanceIDs"`
}

func (r *DeleteGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGrafanaInstanceResponseParams `json:"Response"`
}

func (r *DeleteGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaIntegrationRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成 ID，可在实例详情-云产品集成-集成列表查看。例如：integration-abcd1234
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`
}

type DeleteGrafanaIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成 ID，可在实例详情-云产品集成-集成列表查看。例如：integration-abcd1234
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`
}

func (r *DeleteGrafanaIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGrafanaIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaIntegrationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGrafanaIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGrafanaIntegrationResponseParams `json:"Response"`
}

func (r *DeleteGrafanaIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaNotificationChannelRequestParams struct {
	// 通道 ID 数组。例如：nchannel-abcd1234
	ChannelIDs []*string `json:"ChannelIDs,omitempty" name:"ChannelIDs"`

	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeleteGrafanaNotificationChannelRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID 数组。例如：nchannel-abcd1234
	ChannelIDs []*string `json:"ChannelIDs,omitempty" name:"ChannelIDs"`

	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteGrafanaNotificationChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaNotificationChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelIDs")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGrafanaNotificationChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGrafanaNotificationChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteGrafanaNotificationChannelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGrafanaNotificationChannelResponseParams `json:"Response"`
}

func (r *DeleteGrafanaNotificationChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGrafanaNotificationChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyGroupRequestParams struct {
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId []*int64 `json:"GroupId,omitempty" name:"GroupId"`
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

// Predefined struct for user
type DeletePolicyGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeletePolicyGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DeletePrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警策略id列表
	AlertIds []*string `json:"AlertIds,omitempty" name:"AlertIds"`

	// 告警策略名称
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type DeletePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警策略id列表
	AlertIds []*string `json:"AlertIds,omitempty" name:"AlertIds"`

	// 告警策略名称
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DeletePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertIds")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *DeletePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusClusterAgentRequestParams struct {
	// agent列表
	Agents []*PrometheusAgentInfo `json:"Agents,omitempty" name:"Agents"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeletePrometheusClusterAgentRequest struct {
	*tchttp.BaseRequest
	
	// agent列表
	Agents []*PrometheusAgentInfo `json:"Agents,omitempty" name:"Agents"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeletePrometheusClusterAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusClusterAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agents")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusClusterAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusClusterAgentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusClusterAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusClusterAgentResponseParams `json:"Response"`
}

func (r *DeletePrometheusClusterAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusClusterAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 要删除的ServiceMonitor名字列表
	ServiceMonitors []*string `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// 要删除的PodMonitor名字列表
	PodMonitors []*string `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// 要删除的RawJobs名字列表
	RawJobs []*string `json:"RawJobs,omitempty" name:"RawJobs"`
}

type DeletePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 要删除的ServiceMonitor名字列表
	ServiceMonitors []*string `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// 要删除的PodMonitor名字列表
	PodMonitors []*string `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// 要删除的RawJobs名字列表
	RawJobs []*string `json:"RawJobs,omitempty" name:"RawJobs"`
}

func (r *DeletePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusConfigResponseParams `json:"Response"`
}

func (r *DeletePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusRecordRuleYamlRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 聚合规则列表
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type DeletePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 聚合规则列表
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DeletePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusRecordRuleYamlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *DeletePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusScrapeJobsRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 任务 ID 列表
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`
}

type DeletePrometheusScrapeJobsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 任务 ID 列表
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`
}

func (r *DeletePrometheusScrapeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusScrapeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "JobIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusScrapeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusScrapeJobsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusScrapeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusScrapeJobsResponseParams `json:"Response"`
}

func (r *DeletePrometheusScrapeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusScrapeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempRequestParams struct {
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DeletePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeletePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTempResponseParams `json:"Response"`
}

func (r *DeletePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempSyncRequestParams struct {
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 取消同步的对象列表
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitempty" name:"Targets"`
}

type DeletePrometheusTempSyncRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 取消同步的对象列表
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *DeletePrometheusTempSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTempSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempSyncResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePrometheusTempSyncResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTempSyncResponseParams `json:"Response"`
}

func (r *DeletePrometheusTempSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordingRulesRequestParams struct {
	// 规则 ID 列表
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeleteRecordingRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则 ID 列表
	RuleIds []*string `json:"RuleIds,omitempty" name:"RuleIds"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteRecordingRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordingRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordingRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordingRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRecordingRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordingRulesResponseParams `json:"Response"`
}

func (r *DeleteRecordingRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordingRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSSOAccountRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户账号 ID ，例如：10000000
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DeleteSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户账号 ID ，例如：10000000
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DeleteSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSSOAccountResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSSOAccountResponseParams `json:"Response"`
}

func (r *DeleteSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceDiscoveryRequestParams struct {
	// Prometheus 实例 ID，例如：prom-sdfk2342a
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

type DeleteServiceDiscoveryRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID，例如：prom-sdfk2342a
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

// Predefined struct for user
type DeleteServiceDiscoveryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceDiscoveryResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAccidentEventListRequestParams struct {
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

// Predefined struct for user
type DescribeAccidentEventListResponseParams struct {
	// 平台事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarms []*DescribeAccidentEventListAlarms `json:"Alarms,omitempty" name:"Alarms"`

	// 平台事件的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccidentEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccidentEventListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmEventsRequestParams struct {
	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略类型，由 DescribeAllNamespaces 获得，例如 cvm_device
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 监控类型，如 MT_QCE。如果不填默认为 MT_QCE。
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
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

// Predefined struct for user
type DescribeAlarmEventsResponseParams struct {
	// 告警事件列表
	Events []*AlarmEvent `json:"Events,omitempty" name:"Events"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmEventsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmHistoriesRequestParams struct {
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

	// 根据监控类型过滤，不选默认查所有类型。"MT_QCE"=云产品监控，支持的枚举值有："MT_QCE"=云产品监控；"MT_TAW"=应用性能观测；"MT_RUM"=前端性能监控；"MT_PROBE"=云拨测
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

	// 根据监控类型过滤，不选默认查所有类型。"MT_QCE"=云产品监控，支持的枚举值有："MT_QCE"=云产品监控；"MT_TAW"=应用性能观测；"MT_RUM"=前端性能监控；"MT_PROBE"=云拨测
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

// Predefined struct for user
type DescribeAlarmHistoriesResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 告警历史列表
	Histories []*AlarmHistory `json:"Histories,omitempty" name:"Histories"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmHistoriesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmMetricsRequestParams struct {
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 监控类型过滤 "MT_QCE"=云产品监控
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 告警策略类型，由 DescribeAllNamespaces 获得，例如 cvm_device
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

// Predefined struct for user
type DescribeAlarmMetricsResponseParams struct {
	// 告警指标列表
	Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmMetricsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmNoticeCallbacksRequestParams struct {
	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`
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

// Predefined struct for user
type DescribeAlarmNoticeCallbacksResponseParams struct {
	// 告警回调通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmNoticeCallbacksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNoticeCallbacksResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmNoticeRequestParams struct {
	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警通知模板 id
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`
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

// Predefined struct for user
type DescribeAlarmNoticeResponseParams struct {
	// 告警通知模板详细信息
	Notice *AlarmNotice `json:"Notice,omitempty" name:"Notice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNoticeResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmNoticesRequestParams struct {
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

	// 模板根据标签过滤
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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

	// 模板根据标签过滤
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNoticesResponseParams struct {
	// 告警通知模板总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 告警通知模板列表
	Notices []*AlarmNotice `json:"Notices,omitempty" name:"Notices"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNoticesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmPoliciesRequestParams struct {
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
	// 
	// 注意：如果NeedCorrespondence传入1，即需要返回策略与实例对应关系，请传入不多于20个告警对象维度，否则容易请求超时
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

	// 是否需要策略与入参过滤维度参数的对应关系，1：是  0：否，默认为0
	NeedCorrespondence *int64 `json:"NeedCorrespondence,omitempty" name:"NeedCorrespondence"`

	// 按照触发任务（例如弹性伸缩）过滤策略。最多10个
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`

	// 根据一键告警策略筛选 不传展示全部策略 ONECLICK=展示一键告警策略 NOT_ONECLICK=展示非一键告警策略
	OneClickPolicyType []*string `json:"OneClickPolicyType,omitempty" name:"OneClickPolicyType"`

	// 根据全部对象过滤，1代表需要过滤掉全部对象，0则无需过滤
	NotBindAll *int64 `json:"NotBindAll,omitempty" name:"NotBindAll"`

	// 根据实例对象过滤，1代表需要过滤掉有实例对象，0则无需过滤
	NotInstanceGroup *int64 `json:"NotInstanceGroup,omitempty" name:"NotInstanceGroup"`

	// 策略根据标签过滤
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	// 
	// 注意：如果NeedCorrespondence传入1，即需要返回策略与实例对应关系，请传入不多于20个告警对象维度，否则容易请求超时
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

	// 是否需要策略与入参过滤维度参数的对应关系，1：是  0：否，默认为0
	NeedCorrespondence *int64 `json:"NeedCorrespondence,omitempty" name:"NeedCorrespondence"`

	// 按照触发任务（例如弹性伸缩）过滤策略。最多10个
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`

	// 根据一键告警策略筛选 不传展示全部策略 ONECLICK=展示一键告警策略 NOT_ONECLICK=展示非一键告警策略
	OneClickPolicyType []*string `json:"OneClickPolicyType,omitempty" name:"OneClickPolicyType"`

	// 根据全部对象过滤，1代表需要过滤掉全部对象，0则无需过滤
	NotBindAll *int64 `json:"NotBindAll,omitempty" name:"NotBindAll"`

	// 根据实例对象过滤，1代表需要过滤掉有实例对象，0则无需过滤
	NotInstanceGroup *int64 `json:"NotInstanceGroup,omitempty" name:"NotInstanceGroup"`

	// 策略根据标签过滤
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
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
	delete(f, "NeedCorrespondence")
	delete(f, "TriggerTasks")
	delete(f, "OneClickPolicyType")
	delete(f, "NotBindAll")
	delete(f, "NotInstanceGroup")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmPoliciesResponseParams struct {
	// 策略总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 策略数组
	Policies []*AlarmPolicy `json:"Policies,omitempty" name:"Policies"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmPoliciesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlarmPolicyRequestParams struct {
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
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

// Predefined struct for user
type DescribeAlarmPolicyResponseParams struct {
	// 策略详情
	Policy *AlarmPolicy `json:"Policy,omitempty" name:"Policy"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmPolicyResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAlertRulesRequestParams struct {
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

// Predefined struct for user
type DescribeAlertRulesResponseParams struct {
	// 报警规则数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 报警规则详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertRuleSet []*PrometheusRuleSet `json:"AlertRuleSet,omitempty" name:"AlertRuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlertRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlertRulesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAllNamespacesRequestParams struct {
	// 根据使用场景过滤 目前仅有"ST_ALARM"=告警类型
	SceneType *string `json:"SceneType,omitempty" name:"SceneType"`

	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 根据监控类型过滤 不填默认查所有类型 "MT_QCE"=云产品监控
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`

	// 根据namespace的Id过滤 不填默认查询所有
	Ids []*string `json:"Ids,omitempty" name:"Ids"`
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

// Predefined struct for user
type DescribeAllNamespacesResponseParams struct {
	// 云产品的告警策略类型，已废弃
	QceNamespaces *CommonNamespace `json:"QceNamespaces,omitempty" name:"QceNamespaces"`

	// 其他告警策略类型，已废弃
	CustomNamespaces *CommonNamespace `json:"CustomNamespaces,omitempty" name:"CustomNamespaces"`

	// 云产品的告警策略类型
	QceNamespacesNew []*CommonNamespace `json:"QceNamespacesNew,omitempty" name:"QceNamespacesNew"`

	// 其他告警策略类型，暂不支持
	CustomNamespacesNew []*CommonNamespace `json:"CustomNamespacesNew,omitempty" name:"CustomNamespacesNew"`

	// 通用告警策略类型(包括：应用性能监控，前端性能监控，云拨测)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonNamespaces []*CommonNamespaceNew `json:"CommonNamespaces,omitempty" name:"CommonNamespaces"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAllNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllNamespacesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBaseMetricsRequestParams struct {
	// 业务命名空间，各个云产品的业务命名空间不同。如需获取业务命名空间，请前往各产品监控指标文档，例如云服务器的命名空间，可参见 [云服务器监控指标](https://cloud.tencent.com/document/product/248/6843)
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名，各个云产品的指标名不同。如需获取指标名，请前往各产品监控指标文档，例如云服务器的指标名，可参见 [云服务器监控指标](https://cloud.tencent.com/document/product/248/6843)
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 可选参数，按照维度过滤
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 业务命名空间，各个云产品的业务命名空间不同。如需获取业务命名空间，请前往各产品监控指标文档，例如云服务器的命名空间，可参见 [云服务器监控指标](https://cloud.tencent.com/document/product/248/6843)
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名，各个云产品的指标名不同。如需获取指标名，请前往各产品监控指标文档，例如云服务器的指标名，可参见 [云服务器监控指标](https://cloud.tencent.com/document/product/248/6843)
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 可选参数，按照维度过滤
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
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
	delete(f, "Dimensions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaseMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseMetricsResponseParams struct {
	// 查询得到的指标描述列表
	MetricSet []*MetricSet `json:"MetricSet,omitempty" name:"MetricSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaseMetricsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBasicAlarmListRequestParams struct {
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

// Predefined struct for user
type DescribeBasicAlarmListResponseParams struct {
	// 告警列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alarms []*DescribeBasicAlarmListAlarms `json:"Alarms,omitempty" name:"Alarms"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warning *string `json:"Warning,omitempty" name:"Warning"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBasicAlarmListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicAlarmListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBindingPolicyObjectListRequestParams struct {
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

// Predefined struct for user
type DescribeBindingPolicyObjectListResponseParams struct {
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
}

type DescribeBindingPolicyObjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindingPolicyObjectListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeConditionsTemplateListRequestParams struct {
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图名，由 [DescribeAllNamespaces](https://cloud.tencent.com/document/product/248/48683) 获得。对于云产品监控，取接口出参的 QceNamespacesNew.N.Id，例如 cvm_device
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 根据触发条件模板名称过滤查询
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 根据触发条件模板ID过滤查询
	GroupID *string `json:"GroupID,omitempty" name:"GroupID"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定按更新时间的排序方式，asc=升序, desc=降序
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// 指定按绑定策略数目的排序方式，asc=升序, desc=降序
	PolicyCountOrder *string `json:"PolicyCountOrder,omitempty" name:"PolicyCountOrder"`
}

type DescribeConditionsTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 视图名，由 [DescribeAllNamespaces](https://cloud.tencent.com/document/product/248/48683) 获得。对于云产品监控，取接口出参的 QceNamespacesNew.N.Id，例如 cvm_device
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 根据触发条件模板名称过滤查询
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 根据触发条件模板ID过滤查询
	GroupID *string `json:"GroupID,omitempty" name:"GroupID"`

	// 分页参数，每页返回的数量，取值1~100，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页参数，页偏移量，从0开始计数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 指定按更新时间的排序方式，asc=升序, desc=降序
	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`

	// 指定按绑定策略数目的排序方式，asc=升序, desc=降序
	PolicyCountOrder *string `json:"PolicyCountOrder,omitempty" name:"PolicyCountOrder"`
}

func (r *DescribeConditionsTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConditionsTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "ViewName")
	delete(f, "GroupName")
	delete(f, "GroupID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "UpdateTimeOrder")
	delete(f, "PolicyCountOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConditionsTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConditionsTemplateListResponseParams struct {
	// 模板总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateGroupList []*TemplateGroup `json:"TemplateGroupList,omitempty" name:"TemplateGroupList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConditionsTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConditionsTemplateListResponseParams `json:"Response"`
}

func (r *DescribeConditionsTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConditionsTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDNSConfigRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDNSConfigRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDNSConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDNSConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDNSConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDNSConfigResponseParams struct {
	// DNS 服务器数组
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDNSConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDNSConfigResponseParams `json:"Response"`
}

func (r *DescribeDNSConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDNSConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExporterIntegrationsRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeExporterIntegrationsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeExporterIntegrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExporterIntegrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	delete(f, "Kind")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExporterIntegrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExporterIntegrationsResponseParams struct {
	// 集成配置列表
	IntegrationSet []*IntegrationConfiguration `json:"IntegrationSet,omitempty" name:"IntegrationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeExporterIntegrationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExporterIntegrationsResponseParams `json:"Response"`
}

func (r *DescribeExporterIntegrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExporterIntegrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaChannelsRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 告警通道 ID，例如：nchannel-abcd1234
	ChannelIds []*string `json:"ChannelIds,omitempty" name:"ChannelIds"`

	// 告警通道状态
	ChannelState *int64 `json:"ChannelState,omitempty" name:"ChannelState"`
}

type DescribeGrafanaChannelsRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 告警通道 ID，例如：nchannel-abcd1234
	ChannelIds []*string `json:"ChannelIds,omitempty" name:"ChannelIds"`

	// 告警通道状态
	ChannelState *int64 `json:"ChannelState,omitempty" name:"ChannelState"`
}

func (r *DescribeGrafanaChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ChannelName")
	delete(f, "ChannelIds")
	delete(f, "ChannelState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaChannelsResponseParams struct {
	// 告警通道数组
	NotificationChannelSet []*GrafanaChannel `json:"NotificationChannelSet,omitempty" name:"NotificationChannelSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGrafanaChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaChannelsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaConfigRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeGrafanaConfigRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGrafanaConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaConfigResponseParams struct {
	// JSON 编码后的字符串
	Config *string `json:"Config,omitempty" name:"Config"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGrafanaConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaConfigResponseParams `json:"Response"`
}

func (r *DescribeGrafanaConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaEnvironmentsRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeGrafanaEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGrafanaEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaEnvironmentsResponseParams struct {
	// 环境变量字符串
	Envs *string `json:"Envs,omitempty" name:"Envs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGrafanaEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaInstancesRequestParams struct {
	// 查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Grafana 实例 ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Grafana 实例名，支持前缀模糊搜索
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 查询状态
	InstanceStatus []*int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 标签过滤数组
	TagFilters []*PrometheusTag `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeGrafanaInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Grafana 实例 ID 数组
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// Grafana 实例名，支持前缀模糊搜索
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 查询状态
	InstanceStatus []*int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 标签过滤数组
	TagFilters []*PrometheusTag `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeGrafanaInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIds")
	delete(f, "InstanceName")
	delete(f, "InstanceStatus")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaInstancesResponseParams struct {
	// 已废弃，请使用 Instances
	InstanceSet []*GrafanaInstanceInfo `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 符合查询条件的实例总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例列表
	Instances []*GrafanaInstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGrafanaInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaInstancesResponseParams `json:"Response"`
}

func (r *DescribeGrafanaInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaIntegrationsRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成 ID
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

type DescribeGrafanaIntegrationsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成 ID
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`
}

func (r *DescribeGrafanaIntegrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaIntegrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationId")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaIntegrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaIntegrationsResponseParams struct {
	// 集成数组
	IntegrationSet []*GrafanaIntegrationConfig `json:"IntegrationSet,omitempty" name:"IntegrationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGrafanaIntegrationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaIntegrationsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaIntegrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaIntegrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaNotificationChannelsRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 告警通道 ID，例如：nchannel-abcd1234
	ChannelIDs []*string `json:"ChannelIDs,omitempty" name:"ChannelIDs"`

	// 告警通道状态
	ChannelState *int64 `json:"ChannelState,omitempty" name:"ChannelState"`
}

type DescribeGrafanaNotificationChannelsRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 告警通道 ID，例如：nchannel-abcd1234
	ChannelIDs []*string `json:"ChannelIDs,omitempty" name:"ChannelIDs"`

	// 告警通道状态
	ChannelState *int64 `json:"ChannelState,omitempty" name:"ChannelState"`
}

func (r *DescribeGrafanaNotificationChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaNotificationChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ChannelName")
	delete(f, "ChannelIDs")
	delete(f, "ChannelState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaNotificationChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaNotificationChannelsResponseParams struct {
	// 告警通道数组
	NotificationChannelSet []*GrafanaNotificationChannel `json:"NotificationChannelSet,omitempty" name:"NotificationChannelSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGrafanaNotificationChannelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaNotificationChannelsResponseParams `json:"Response"`
}

func (r *DescribeGrafanaNotificationChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaNotificationChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaWhiteListRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeGrafanaWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGrafanaWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGrafanaWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGrafanaWhiteListResponseParams struct {
	// 数组
	WhiteList []*string `json:"WhiteList,omitempty" name:"WhiteList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeGrafanaWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGrafanaWhiteListResponseParams `json:"Response"`
}

func (r *DescribeGrafanaWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGrafanaWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstalledPluginsRequestParams struct {
	// Grafana 实例 ID，例如：grafana-kleu3gt0
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按插件 ID 过滤，例如：grafana-piechart-panel，可通过接口 DescribeInstalledPlugins 查看已安装的插件 ID
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

type DescribeInstalledPluginsRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-kleu3gt0
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按插件 ID 过滤，例如：grafana-piechart-panel，可通过接口 DescribeInstalledPlugins 查看已安装的插件 ID
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`
}

func (r *DescribeInstalledPluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstalledPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PluginId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstalledPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstalledPluginsResponseParams struct {
	// 插件列表
	PluginSet []*GrafanaPlugin `json:"PluginSet,omitempty" name:"PluginSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstalledPluginsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstalledPluginsResponseParams `json:"Response"`
}

func (r *DescribeInstalledPluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstalledPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorTypesRequestParams struct {
	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`
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

// Predefined struct for user
type DescribeMonitorTypesResponseParams struct {
	// 监控类型，云产品监控为 MT_QCE
	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`

	// 监控类型详情
	MonitorTypeInfos []*MonitorTypeInfo `json:"MonitorTypeInfos,omitempty" name:"MonitorTypeInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMonitorTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorTypesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePluginOverviewsRequestParams struct {

}

type DescribePluginOverviewsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePluginOverviewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginOverviewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginOverviewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginOverviewsResponseParams struct {
	// 插件列表
	PluginSet []*GrafanaPlugin `json:"PluginSet,omitempty" name:"PluginSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePluginOverviewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginOverviewsResponseParams `json:"Response"`
}

func (r *DescribePluginOverviewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginOverviewsResponse) FromJsonString(s string) error {
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

	// 弃用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeprecatingInfo *DescribePolicyConditionListResponseDeprecatingInfo `json:"DeprecatingInfo,omitempty" name:"DeprecatingInfo"`
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

// Predefined struct for user
type DescribePolicyConditionListRequestParams struct {
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`
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

// Predefined struct for user
type DescribePolicyConditionListResponseParams struct {
	// 告警策略条件列表
	Conditions []*DescribePolicyConditionListCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePolicyConditionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyConditionListResponseParams `json:"Response"`
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

type DescribePolicyConditionListResponseDeprecatingInfo struct {
	// 是否隐藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hidden *bool `json:"Hidden,omitempty" name:"Hidden"`

	// 新视图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewViewNames []*string `json:"NewViewNames,omitempty" name:"NewViewNames"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
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

	// 告警时间段开始时间。范围[0,86400)，作为 UNIX 时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"
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

// Predefined struct for user
type DescribePolicyGroupInfoRequestParams struct {
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
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

// Predefined struct for user
type DescribePolicyGroupInfoResponseParams struct {
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
}

type DescribePolicyGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyGroupInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePolicyGroupListRequestParams struct {
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

// Predefined struct for user
type DescribePolicyGroupListResponseParams struct {
	// 策略组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupList []*DescribePolicyGroupListGroup `json:"GroupList,omitempty" name:"GroupList"`

	// 策略组总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warning *string `json:"Warning,omitempty" name:"Warning"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePolicyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyGroupListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeProductEventListRequestParams struct {
	// 接口模块名，固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 产品类型过滤，例如"cvm"表示云服务器
	ProductName []*string `json:"ProductName,omitempty" name:"ProductName"`

	// 事件名称过滤，例如"guest_reboot"表示机器重启
	EventName []*string `json:"EventName,omitempty" name:"EventName"`

	// 影响对象，例如"ins-19708ino"
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维度过滤，例如外网IP:10.0.0.1
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions"`

	// 产品事件地域过滤参数，例如gz，各地域缩写可参见[地域列表](https://cloud.tencent.com/document/product/248/50863)
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

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest
	
	// 接口模块名，固定值"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 产品类型过滤，例如"cvm"表示云服务器
	ProductName []*string `json:"ProductName,omitempty" name:"ProductName"`

	// 事件名称过滤，例如"guest_reboot"表示机器重启
	EventName []*string `json:"EventName,omitempty" name:"EventName"`

	// 影响对象，例如"ins-19708ino"
	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维度过滤，例如外网IP:10.0.0.1
	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions"`

	// 产品事件地域过滤参数，例如gz，各地域缩写可参见[地域列表](https://cloud.tencent.com/document/product/248/50863)
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

// Predefined struct for user
type DescribeProductEventListResponseParams struct {
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
}

type DescribeProductEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductEventListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeProductListRequestParams struct {
	// 固定传值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 排序方式：DESC/ASC（区分大小写），默认值DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 分页查询的偏移量，默认值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询的每页数据量，默认值20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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

// Predefined struct for user
type DescribeProductListResponseParams struct {
	// 产品信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductList []*ProductSimple `json:"ProductList,omitempty" name:"ProductList"`

	// 产品总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePrometheusAgentInstancesRequestParams struct {
	// 集群id
	// 可以是tke, eks, edge的集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribePrometheusAgentInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	// 可以是tke, eks, edge的集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribePrometheusAgentInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAgentInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentInstancesResponseParams struct {
	// 关联该集群的实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*string `json:"Instances,omitempty" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusAgentInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAgentInstancesResponseParams `json:"Response"`
}

func (r *DescribePrometheusAgentInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentsRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Agent ID 列表
	AgentIds []*string `json:"AgentIds,omitempty" name:"AgentIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePrometheusAgentsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Agent ID 列表
	AgentIds []*string `json:"AgentIds,omitempty" name:"AgentIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrometheusAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "AgentIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentsResponseParams struct {
	// Agent 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentSet []*PrometheusAgent `json:"AgentSet,omitempty" name:"AgentSet"`

	// Agent 总量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusAgentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAgentsResponseParams `json:"Response"`
}

func (r *DescribePrometheusAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤
	// 支持ID，Name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤
	// 支持ID，Name
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertPolicyResponseParams struct {
	// 告警详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertRules []*PrometheusAlertPolicyItem `json:"AlertRules,omitempty" name:"AlertRules"`

	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *DescribePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusClusterAgentsRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用于分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePrometheusClusterAgentsRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用于分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrometheusClusterAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusClusterAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusClusterAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusClusterAgentsResponseParams struct {
	// 被关联集群信息
	Agents []*PrometheusAgentOverview `json:"Agents,omitempty" name:"Agents"`

	// 被关联集群总量
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusClusterAgentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusClusterAgentsResponseParams `json:"Response"`
}

func (r *DescribePrometheusClusterAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusClusterAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type DescribePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

func (r *DescribePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusConfigResponseParams struct {
	// 全局配置
	Config *string `json:"Config,omitempty" name:"Config"`

	// ServiceMonitor配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// PodMonitor配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// 原生Job
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`

	// Probes
	Probes []*PrometheusConfigItem `json:"Probes,omitempty" name:"Probes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusConfigResponseParams `json:"Response"`
}

func (r *DescribePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalConfigRequestParams struct {
	// 实例级别抓取配置
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否禁用统计
	DisableStatistics *bool `json:"DisableStatistics,omitempty" name:"DisableStatistics"`
}

type DescribePrometheusGlobalConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例级别抓取配置
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否禁用统计
	DisableStatistics *bool `json:"DisableStatistics,omitempty" name:"DisableStatistics"`
}

func (r *DescribePrometheusGlobalConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DisableStatistics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusGlobalConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalConfigResponseParams struct {
	// 配置内容
	Config *string `json:"Config,omitempty" name:"Config"`

	// ServiceMonitors列表以及对应targets信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// PodMonitors列表以及对应targets信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// RawJobs列表以及对应targets信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusGlobalConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusGlobalConfigResponseParams `json:"Response"`
}

func (r *DescribePrometheusGlobalConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalNotificationRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribePrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalNotificationResponseParams struct {
	// 全局告警通知渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notification *PrometheusNotificationItem `json:"Notification,omitempty" name:"Notification"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *DescribePrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceDetailRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribePrometheusInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceDetailResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例业务状态。取值范围：
	// 
	// 1：正在创建
	// 2：运行中
	// 3：异常
	// 4：重建中
	// 5：销毁中
	// 6：已停服
	// 8：欠费停服中
	// 9：欠费已停服
	InstanceStatus *int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 计费状态
	// 
	// 1：正常
	// 2：过期
	// 3：销毁
	// 4：分配中
	// 5：分配失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *int64 `json:"ChargeStatus,omitempty" name:"ChargeStatus"`

	// 是否开启 Grafana
	// 0：不开启
	// 1：开启
	EnableGrafana *int64 `json:"EnableGrafana,omitempty" name:"EnableGrafana"`

	// Grafana 面板 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaURL *string `json:"GrafanaURL,omitempty" name:"GrafanaURL"`

	// 实例计费模式。取值范围：
	// 
	// 2：包年包月
	// 3：按量
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 规格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 存储周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRetentionTime *int64 `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`

	// 购买的实例过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 自动续费标记
	// 
	// 0：不自动续费
	// 1：开启自动续费
	// 2：禁止自动续费
	// -1：无效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceInitStatusRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribePrometheusInstanceInitStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusInstanceInitStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceInitStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceInitStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceInitStatusResponseParams struct {
	// 实例初始化状态，取值：
	// uninitialized 未初始化 
	// initializing 初始化中
	// running 初始化完成，运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 初始化任务步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Steps []*TaskStepInfo `json:"Steps,omitempty" name:"Steps"`

	// 实例eks集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EksClusterId *string `json:"EksClusterId,omitempty" name:"EksClusterId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceInitStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceInitStatusResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceInitStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceInitStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceUsageRequestParams struct {
	// 按照一个或者多个实例ID查询。实例ID形如：prom-xxxxxxxx。请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 开始时间
	StartCalcDate *string `json:"StartCalcDate,omitempty" name:"StartCalcDate"`

	// 结束时间
	EndCalcDate *string `json:"EndCalcDate,omitempty" name:"EndCalcDate"`
}

type DescribePrometheusInstanceUsageRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询。实例ID形如：prom-xxxxxxxx。请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 开始时间
	StartCalcDate *string `json:"StartCalcDate,omitempty" name:"StartCalcDate"`

	// 结束时间
	EndCalcDate *string `json:"EndCalcDate,omitempty" name:"EndCalcDate"`
}

func (r *DescribePrometheusInstanceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "StartCalcDate")
	delete(f, "EndCalcDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceUsageResponseParams struct {
	// 用量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageSet []*PrometheusInstanceTenantUsage `json:"UsageSet,omitempty" name:"UsageSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceUsageResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesOverviewRequestParams struct {
	// 用于分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤实例，目前支持：
	// ID: 通过实例ID来过滤 
	// Name: 通过实例名称来过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrometheusInstancesOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤实例，目前支持：
	// ID: 通过实例ID来过滤 
	// Name: 通过实例名称来过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrometheusInstancesOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstancesOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesOverviewResponseParams struct {
	// 实例列表
	Instances []*PrometheusInstancesOverview `json:"Instances,omitempty" name:"Instances"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusInstancesOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstancesOverviewResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstancesOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesRequestParams struct {
	// 按照一个或者多个实例ID查询。实例ID形如：prom-xxxxxxxx。请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 按照【实例状态】进行过滤。
	// <ul>
	// <li>1：正在创建</li>
	// <li>2：运行中</li>
	// <li>3：异常</li>
	// <li>4：重建中</li>
	// <li>5：销毁中</li>
	// <li>6：已停服</li>
	// <li>8：欠费停服中</li>
	// <li>9：欠费已停服</li>
	// </ul>
	InstanceStatus []*int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 按照【实例名称】进行过滤。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 按照【可用区】进行过滤。可用区形如：ap-guangzhou-1。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。
	TagFilters []*PrometheusTag `json:"TagFilters,omitempty" name:"TagFilters"`

	// 按照【实例的IPv4地址】进行过滤。
	IPv4Address []*string `json:"IPv4Address,omitempty" name:"IPv4Address"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照【计费类型】进行过滤。
	// <li>2：包年包月</li>
	// <li>3：按量</li>
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

type DescribePrometheusInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个实例ID查询。实例ID形如：prom-xxxxxxxx。请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 按照【实例状态】进行过滤。
	// <ul>
	// <li>1：正在创建</li>
	// <li>2：运行中</li>
	// <li>3：异常</li>
	// <li>4：重建中</li>
	// <li>5：销毁中</li>
	// <li>6：已停服</li>
	// <li>8：欠费停服中</li>
	// <li>9：欠费已停服</li>
	// </ul>
	InstanceStatus []*int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 按照【实例名称】进行过滤。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 按照【可用区】进行过滤。可用区形如：ap-guangzhou-1。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。
	TagFilters []*PrometheusTag `json:"TagFilters,omitempty" name:"TagFilters"`

	// 按照【实例的IPv4地址】进行过滤。
	IPv4Address []*string `json:"IPv4Address,omitempty" name:"IPv4Address"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照【计费类型】进行过滤。
	// <li>2：包年包月</li>
	// <li>3：按量</li>
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

func (r *DescribePrometheusInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceStatus")
	delete(f, "InstanceName")
	delete(f, "Zones")
	delete(f, "TagFilters")
	delete(f, "IPv4Address")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesResponseParams struct {
	// 实例详细信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*PrometheusInstancesItem `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstancesResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRuleYamlRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤，当前支持
	// Name = Name
	// Values = 目标名称列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤，当前支持
	// Name = Name
	// Values = 目标名称列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRuleYamlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *DescribePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRulesRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrometheusRecordRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrometheusRecordRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusRecordRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRulesResponseParams struct {
	// 聚合规则
	Records []*PrometheusRecordRuleYamlItem `json:"Records,omitempty" name:"Records"`

	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusRecordRulesResponseParams `json:"Response"`
}

func (r *DescribePrometheusRecordRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusScrapeJobsRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务 ID 列表
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePrometheusScrapeJobsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 任务名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 任务 ID 列表
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrometheusScrapeJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusScrapeJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "Name")
	delete(f, "JobIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusScrapeJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusScrapeJobsResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScrapeJobSet []*PrometheusScrapeJob `json:"ScrapeJobSet,omitempty" name:"ScrapeJobSet"`

	// 任务总量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusScrapeJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusScrapeJobsResponseParams `json:"Response"`
}

func (r *DescribePrometheusScrapeJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusScrapeJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTargetsTMPRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 过滤条件，当前支持
	// Name=state
	// Value=up, down, unknown
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribePrometheusTargetsTMPRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 过滤条件，当前支持
	// Name=state
	// Value=up, down, unknown
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrometheusTargetsTMPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTargetsTMPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTargetsTMPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTargetsTMPResponseParams struct {
	// 所有Job的targets信息
	Jobs []*PrometheusJobTargets `json:"Jobs,omitempty" name:"Jobs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusTargetsTMPResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTargetsTMPResponseParams `json:"Response"`
}

func (r *DescribePrometheusTargetsTMPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTargetsTMPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempRequestParams struct {
	// 模糊过滤条件，支持
	// Level 按模板级别过滤
	// Name 按名称过滤
	// Describe 按描述过滤
	// ID 按templateId过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 总数限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模糊过滤条件，支持
	// Level 按模板级别过滤
	// Name 按名称过滤
	// Describe 按描述过滤
	// ID 按templateId过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 总数限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempResponseParams struct {
	// 模板列表
	Templates []*PrometheusTemp `json:"Templates,omitempty" name:"Templates"`

	// 总数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTempResponseParams `json:"Response"`
}

func (r *DescribePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempSyncRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribePrometheusTempSyncRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribePrometheusTempSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTempSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempSyncResponseParams struct {
	// 同步目标详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitempty" name:"Targets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusTempSyncResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTempSyncResponseParams `json:"Response"`
}

func (r *DescribePrometheusTempSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusZonesRequestParams struct {
	// 地域 ID（RegionId 和 RegionName 只需要填一个）
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名（RegionId 和 RegionName 只需要填一个）
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type DescribePrometheusZonesRequest struct {
	*tchttp.BaseRequest
	
	// 地域 ID（RegionId 和 RegionName 只需要填一个）
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名（RegionId 和 RegionName 只需要填一个）
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

func (r *DescribePrometheusZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionId")
	delete(f, "RegionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusZonesResponseParams struct {
	// 区域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneSet []*PrometheusZoneItem `json:"ZoneSet,omitempty" name:"ZoneSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrometheusZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusZonesResponseParams `json:"Response"`
}

func (r *DescribePrometheusZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingRulesRequestParams struct {
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回数量，默认为 20，最大值为 100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则状态码，取值如下：
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeRecordingRulesRequest struct {
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
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeRecordingRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "RuleId")
	delete(f, "RuleState")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordingRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingRulesResponseParams struct {
	// 规则组数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 规则组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordingRuleSet []*RecordingRuleSet `json:"RecordingRuleSet,omitempty" name:"RecordingRuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordingRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordingRulesResponseParams `json:"Response"`
}

func (r *DescribeRecordingRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSOAccountRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 填写对应的账号 ID，将会按账号 ID 进行过滤，例如：10000
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 填写对应的账号 ID，将会按账号 ID 进行过滤，例如：10000
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSOAccountResponseParams struct {
	// 授权账号列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountSet []*GrafanaAccountInfo `json:"AccountSet,omitempty" name:"AccountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSSOAccountResponseParams `json:"Response"`
}

func (r *DescribeSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceDiscoveryRequestParams struct {
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// <li>类型是 TKE，为对应的腾讯云容器服务集群 ID</li>
	KubeClusterId *string `json:"KubeClusterId,omitempty" name:"KubeClusterId"`

	// 用户 Kubernetes 集群类型：
	// <li> 1 = 容器服务集群(TKE) </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`
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

// Predefined struct for user
type DescribeServiceDiscoveryResponseParams struct {
	// 返回服务发现列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDiscoverySet []*ServiceDiscoveryItem `json:"ServiceDiscoverySet,omitempty" name:"ServiceDiscoverySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceDiscoveryResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeStatisticDataRequestParams struct {
	// 所属模块，固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 命名空间，目前只支持QCE/TKE2
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名列表
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 维度条件，操作符支持=、in
	Conditions []*MidQueryCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 统计粒度。默认取值为300，单位为s；可选的值为60、300、3600、86400
	// 受存储时长限制，统计粒度与统计的时间范围有关：
	// 60s：EndTime-StartTime<12小时，且StartTime距当前时间不能超过15天；
	// 300s：EndTime-StartTime<3天，且StartTime距当前时间不能超过31天；
	// 3600s：EndTime-StartTime<30天，且StartTime距当前时间不能超过93天；
	// 86400s：EndTime-StartTime<186天，且StartTime距当前时间不能超过186天。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 起始时间，默认为当前时间，如2020-12-08T19:51:23+08:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认为当前时间，如2020-12-08T19:51:23+08:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 按指定维度groupBy
	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys"`
}

type DescribeStatisticDataRequest struct {
	*tchttp.BaseRequest
	
	// 所属模块，固定值，为monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 命名空间，目前只支持QCE/TKE2
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 指标名列表
	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`

	// 维度条件，操作符支持=、in
	Conditions []*MidQueryCondition `json:"Conditions,omitempty" name:"Conditions"`

	// 统计粒度。默认取值为300，单位为s；可选的值为60、300、3600、86400
	// 受存储时长限制，统计粒度与统计的时间范围有关：
	// 60s：EndTime-StartTime<12小时，且StartTime距当前时间不能超过15天；
	// 300s：EndTime-StartTime<3天，且StartTime距当前时间不能超过31天；
	// 3600s：EndTime-StartTime<30天，且StartTime距当前时间不能超过93天；
	// 86400s：EndTime-StartTime<186天，且StartTime距当前时间不能超过186天。
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

// Predefined struct for user
type DescribeStatisticDataResponseParams struct {
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
}

type DescribeStatisticDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStatisticDataResponseParams `json:"Response"`
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

// Predefined struct for user
type DestroyPrometheusInstanceRequestParams struct {
	// 实例 ID，该实例必须先被 terminate
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DestroyPrometheusInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，该实例必须先被 terminate
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyPrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrometheusInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPrometheusInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrometheusInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyPrometheusInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPrometheusInstanceResponseParams `json:"Response"`
}

func (r *DestroyPrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrometheusInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {
	// 实例维度名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例维度值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DimensionNew struct {
	// 维度 key 标示，后台英文名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 维度 key 名称，中英文前台展示名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否必选
	IsRequired *bool `json:"IsRequired,omitempty" name:"IsRequired"`

	// 支持的操作符列表
	Operators []*Operator `json:"Operators,omitempty" name:"Operators"`

	// 是否支持多选
	IsMultiple *bool `json:"IsMultiple,omitempty" name:"IsMultiple"`

	// 创建后是否可以修改
	IsMutable *bool `json:"IsMutable,omitempty" name:"IsMutable"`

	// 是否展示给用户
	IsVisible *bool `json:"IsVisible,omitempty" name:"IsVisible"`

	// 能否用来过滤策略列表
	CanFilterPolicy *bool `json:"CanFilterPolicy,omitempty" name:"CanFilterPolicy"`

	// 能否用来过滤告警历史
	CanFilterHistory *bool `json:"CanFilterHistory,omitempty" name:"CanFilterHistory"`

	// 能否作为聚合维度
	CanGroupBy *bool `json:"CanGroupBy,omitempty" name:"CanGroupBy"`

	// 是否必须作为聚合维度
	MustGroupBy *bool `json:"MustGroupBy,omitempty" name:"MustGroupBy"`

	// 前端翻译要替换的 key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowValueReplace *string `json:"ShowValueReplace,omitempty" name:"ShowValueReplace"`
}

type DimensionsDesc struct {
	// 维度名数组
	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

// Predefined struct for user
type EnableGrafanaInternetRequestParams struct {
	// Grafana 实例 ID，例如：grafana-kleu3gt0
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 开启或关闭公网访问，true为开启，false 为不开启
	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`
}

type EnableGrafanaInternetRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-kleu3gt0
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 开启或关闭公网访问，true为开启，false 为不开启
	EnableInternet *bool `json:"EnableInternet,omitempty" name:"EnableInternet"`
}

func (r *EnableGrafanaInternetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaInternetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "EnableInternet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableGrafanaInternetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGrafanaInternetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableGrafanaInternetResponse struct {
	*tchttp.BaseResponse
	Response *EnableGrafanaInternetResponseParams `json:"Response"`
}

func (r *EnableGrafanaInternetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaInternetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGrafanaSSORequestParams struct {
	// 是否开启 SSO，true为开启，false 为不开启
	EnableSSO *bool `json:"EnableSSO,omitempty" name:"EnableSSO"`

	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type EnableGrafanaSSORequest struct {
	*tchttp.BaseRequest
	
	// 是否开启 SSO，true为开启，false 为不开启
	EnableSSO *bool `json:"EnableSSO,omitempty" name:"EnableSSO"`

	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *EnableGrafanaSSORequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaSSORequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnableSSO")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableGrafanaSSORequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGrafanaSSOResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableGrafanaSSOResponse struct {
	*tchttp.BaseResponse
	Response *EnableGrafanaSSOResponseParams `json:"Response"`
}

func (r *EnableGrafanaSSOResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGrafanaSSOResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSSOCamCheckRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否开启 cam 鉴权，true为开启，false 为不开启
	EnableSSOCamCheck *bool `json:"EnableSSOCamCheck,omitempty" name:"EnableSSOCamCheck"`
}

type EnableSSOCamCheckRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否开启 cam 鉴权，true为开启，false 为不开启
	EnableSSOCamCheck *bool `json:"EnableSSOCamCheck,omitempty" name:"EnableSSOCamCheck"`
}

func (r *EnableSSOCamCheckRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSSOCamCheckRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EnableSSOCamCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableSSOCamCheckRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSSOCamCheckResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableSSOCamCheckResponse struct {
	*tchttp.BaseResponse
	Response *EnableSSOCamCheckResponseParams `json:"Response"`
}

func (r *EnableSSOCamCheckResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSSOCamCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventCondition struct {
	// 告警通知频率
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`

	// 重复通知策略预定义（0 - 只告警一次， 1 - 指数告警，2 - 连接告警）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmNotifyType *string `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`

	// 事件ID
	EventID *string `json:"EventID,omitempty" name:"EventID"`

	// 事件展示名称（对外）
	EventDisplayName *string `json:"EventDisplayName,omitempty" name:"EventDisplayName"`

	// 规则ID
	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

type Filter struct {
	// 过滤方式（=, !=, in）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 过滤维度名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 过滤值，in过滤方式用逗号分割多个值
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type GetMonitorDataRequestParams struct {
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

// Predefined struct for user
type GetMonitorDataResponseParams struct {
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

	// 返回信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *GetMonitorDataResponseParams `json:"Response"`
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

// Predefined struct for user
type GetPrometheusAgentManagementCommandRequestParams struct {
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
}

type GetPrometheusAgentManagementCommandRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
}

func (r *GetPrometheusAgentManagementCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPrometheusAgentManagementCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPrometheusAgentManagementCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPrometheusAgentManagementCommandResponseParams struct {
	// Agent 管理命令
	Command *ManagementCommand `json:"Command,omitempty" name:"Command"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPrometheusAgentManagementCommandResponse struct {
	*tchttp.BaseResponse
	Response *GetPrometheusAgentManagementCommandResponseParams `json:"Response"`
}

func (r *GetPrometheusAgentManagementCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPrometheusAgentManagementCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GrafanaAccountInfo struct {
	// 用户账号ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户权限
	Role []*GrafanaAccountRole `json:"Role,omitempty" name:"Role"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitempty" name:"CreateAt"`

	// 实例 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户主账号 UIN
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type GrafanaAccountRole struct {
	// 组织
	Organization *string `json:"Organization,omitempty" name:"Organization"`

	// 权限
	Role *string `json:"Role,omitempty" name:"Role"`
}

type GrafanaChannel struct {
	// 渠道 ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 渠道名
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 告警通道模板 ID 数组
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 告警渠道的所有生效组织
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationIds []*string `json:"OrganizationIds,omitempty" name:"OrganizationIds"`
}

type GrafanaInstanceInfo struct {
	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID 数组
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// Grafana 内网地址
	InternetUrl *string `json:"InternetUrl,omitempty" name:"InternetUrl"`

	// Grafana 公网地址
	InternalUrl *string `json:"InternalUrl,omitempty" name:"InternalUrl"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 运行状态（1:正在创建；2:运行中；3:异常；4:重启中；5:停机中； 6:已停机； 7: 已删除）
	InstanceStatus *int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 实例的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 实例的可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计费模式（1:包年包月）
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// VPC 名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 地域 ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可访问此实例的完整 URL
	RootUrl *string `json:"RootUrl,omitempty" name:"RootUrl"`

	// 是否开启 SSO
	EnableSSO *bool `json:"EnableSSO,omitempty" name:"EnableSSO"`

	// 版本号
	Version *string `json:"Version,omitempty" name:"Version"`

	// SSO登录时是否开启cam鉴权
	EnableSSOCamCheck *bool `json:"EnableSSOCamCheck,omitempty" name:"EnableSSOCamCheck"`
}

type GrafanaIntegrationConfig struct {
	// 集成 ID
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`

	// 集成类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集成内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 集成描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// Grafana 跳转地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaURL *string `json:"GrafanaURL,omitempty" name:"GrafanaURL"`
}

type GrafanaNotificationChannel struct {
	// 渠道 ID
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// 渠道名
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 告警通道模板 ID 数组
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 默认生效组织，已废弃，请使用 OrganizationIds
	OrgId *string `json:"OrgId,omitempty" name:"OrgId"`

	// 额外生效组织，已废弃，请使用 OrganizationIds
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitempty" name:"ExtraOrgIds"`

	// 生效组织，已废弃，请使用 OrganizationIds
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgIds []*string `json:"OrgIds,omitempty" name:"OrgIds"`

	// 告警渠道的所有生效组织
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationIds []*string `json:"OrganizationIds,omitempty" name:"OrganizationIds"`
}

type GrafanaPlugin struct {
	// Grafana 插件 ID
	PluginId *string `json:"PluginId,omitempty" name:"PluginId"`

	// Grafana 插件版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

// Predefined struct for user
type InstallPluginsRequestParams struct {
	// 插件信息
	Plugins []*GrafanaPlugin `json:"Plugins,omitempty" name:"Plugins"`

	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type InstallPluginsRequest struct {
	*tchttp.BaseRequest
	
	// 插件信息
	Plugins []*GrafanaPlugin `json:"Plugins,omitempty" name:"Plugins"`

	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *InstallPluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Plugins")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallPluginsResponseParams struct {
	// 已安装插件 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginIds []*string `json:"PluginIds,omitempty" name:"PluginIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InstallPluginsResponse struct {
	*tchttp.BaseResponse
	Response *InstallPluginsResponseParams `json:"Response"`
}

func (r *InstallPluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallPluginsResponse) FromJsonString(s string) error {
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

type IntegrationConfiguration struct {
	// 名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例类型
	Category *string `json:"Category,omitempty" name:"Category"`

	// 实例描述
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// dashboard 的 URL
	GrafanaDashboardURL *string `json:"GrafanaDashboardURL,omitempty" name:"GrafanaDashboardURL"`
}

type Label struct {
	// map表中的Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// map表中的Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogAlarmReq struct {
	// apm实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 检索条件信息
	Filter []*LogFilterInfo `json:"Filter,omitempty" name:"Filter"`

	// 告警合并开启/暂停
	AlarmMerge *string `json:"AlarmMerge,omitempty" name:"AlarmMerge"`

	// 告警合并时间
	AlarmMergeTime *string `json:"AlarmMergeTime,omitempty" name:"AlarmMergeTime"`
}

type LogFilterInfo struct {
	// 字段名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 比较符号
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 字段值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ManagementCommand struct {
	// Agent 安装命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Install *string `json:"Install,omitempty" name:"Install"`

	// Agent 重启命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Restart *string `json:"Restart,omitempty" name:"Restart"`

	// Agent 停止命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stop *string `json:"Stop,omitempty" name:"Stop"`

	// Agent 状态检测命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCheck *string `json:"StatusCheck,omitempty" name:"StatusCheck"`

	// Agent 日志检测命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogCheck *string `json:"LogCheck,omitempty" name:"LogCheck"`
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

	// 是否为高级指标。1是 0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAdvanced *int64 `json:"IsAdvanced,omitempty" name:"IsAdvanced"`

	// 高级指标是否开通。1是 0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`

	// 集成中心产品ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`
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

// Predefined struct for user
type ModifyAlarmNoticeRequestParams struct {
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

	// 告警通知推送到CLS服务 最多1个
	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`

	// 告警通知模板绑定的告警策略ID列表
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
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

	// 告警通知推送到CLS服务 最多1个
	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`

	// 告警通知模板绑定的告警策略ID列表
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
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
	delete(f, "CLSNotices")
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmNoticeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmNoticeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAlarmPolicyConditionRequestParams struct {
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

	// 日志告警创建请求参数信息
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitempty" name:"LogAlarmReqInfo"`

	// 模版id，专供prom使用
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 启停状态，0=停用，1=启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 专供prom策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`
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

	// 日志告警创建请求参数信息
	LogAlarmReqInfo *LogAlarmReq `json:"LogAlarmReqInfo,omitempty" name:"LogAlarmReqInfo"`

	// 模版id，专供prom使用
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 启停状态，0=停用，1=启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 专供prom策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`
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
	delete(f, "LogAlarmReqInfo")
	delete(f, "NoticeIds")
	delete(f, "Enable")
	delete(f, "PolicyName")
	delete(f, "EbSubject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyConditionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyConditionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyConditionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyConditionResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAlarmPolicyInfoRequestParams struct {
	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 要修改的字段 NAME=策略名称 REMARK=策略备注
	Key *string `json:"Key,omitempty" name:"Key"`

	// 修改后的值
	Value *string `json:"Value,omitempty" name:"Value"`
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

// Predefined struct for user
type ModifyAlarmPolicyInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAlarmPolicyNoticeRequestParams struct {
	// 模块名，这里填“monitor”。
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID，如果该参数与PolicyIds参数同时存在，则以PolicyIds为准。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警通知模板 ID 列表。
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 告警策略ID数组，支持给多个告警策略批量绑定通知模板。最多30个。
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

type ModifyAlarmPolicyNoticeRequest struct {
	*tchttp.BaseRequest
	
	// 模块名，这里填“monitor”。
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID，如果该参数与PolicyIds参数同时存在，则以PolicyIds为准。
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警通知模板 ID 列表。
	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 告警策略ID数组，支持给多个告警策略批量绑定通知模板。最多30个。
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
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
	delete(f, "PolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyNoticeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyNoticeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyNoticeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAlarmPolicyStatusRequestParams struct {
	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 启停状态 0=停用 1=启用
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
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

// Predefined struct for user
type ModifyAlarmPolicyStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAlarmPolicyTasksRequestParams struct {
	// 模块名，这里填“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 告警策略触发任务列表，空数据代表解绑
	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`
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

// Predefined struct for user
type ModifyAlarmPolicyTasksResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyTasksResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyTasksResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAlarmReceiversRequestParams struct {
	// 需要修改接收人的策略组Id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 必填。固定为“monitor”
	Module *string `json:"Module,omitempty" name:"Module"`

	// 新接收人信息, 没有填写则删除所有接收人
	ReceiverInfos []*ReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`
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

// Predefined struct for user
type ModifyAlarmReceiversResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAlarmReceiversResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmReceiversResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyGrafanaInstanceRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Grafana 实例名称，例如：test
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Grafana 实例名称，例如：test
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGrafanaInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGrafanaInstanceResponseParams `json:"Response"`
}

func (r *ModifyGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGrafanaInstanceResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type ModifyPolicyGroupRequestParams struct {
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

// Predefined struct for user
type ModifyPolicyGroupResponseParams struct {
	// 策略组id
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPolicyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPolicyGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPrometheusAgentExternalLabelsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 新的external_labels
	ExternalLabels []*Label `json:"ExternalLabels,omitempty" name:"ExternalLabels"`
}

type ModifyPrometheusAgentExternalLabelsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 新的external_labels
	ExternalLabels []*Label `json:"ExternalLabels,omitempty" name:"ExternalLabels"`
}

func (r *ModifyPrometheusAgentExternalLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAgentExternalLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterId")
	delete(f, "ExternalLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAgentExternalLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAgentExternalLabelsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusAgentExternalLabelsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAgentExternalLabelsResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAgentExternalLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAgentExternalLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitempty" name:"AlertRule"`
}

type ModifyPrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitempty" name:"AlertRule"`
}

func (r *ModifyPrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`
}

type ModifyPrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`
}

func (r *ModifyPrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusConfigResponseParams `json:"Response"`
}

func (r *ModifyPrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusGlobalNotificationRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitempty" name:"Notification"`
}

type ModifyPrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitempty" name:"Notification"`
}

func (r *ModifyPrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Notification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusGlobalNotificationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *ModifyPrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusInstanceAttributesRequestParams struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 存储时长（取值为 15、30、45。此参数不适用于包年包月实例）
	DataRetentionTime *int64 `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`
}

type ModifyPrometheusInstanceAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 存储时长（取值为 15、30、45。此参数不适用于包年包月实例）
	DataRetentionTime *int64 `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`
}

func (r *ModifyPrometheusInstanceAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusInstanceAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "InstanceId")
	delete(f, "DataRetentionTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusInstanceAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusInstanceAttributesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusInstanceAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusInstanceAttributesResponseParams `json:"Response"`
}

func (r *ModifyPrometheusInstanceAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusRecordRuleYamlRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 聚合实例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 新的内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type ModifyPrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 聚合实例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 新的内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *ModifyPrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusRecordRuleYamlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *ModifyPrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTempRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 修改内容
	Template *PrometheusTempModify `json:"Template,omitempty" name:"Template"`
}

type ModifyPrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 修改内容
	Template *PrometheusTempModify `json:"Template,omitempty" name:"Template"`
}

func (r *ModifyPrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTempResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusTempResponseParams `json:"Response"`
}

func (r *ModifyPrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTypeInfo struct {
	// 监控类型ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 监控类型
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排列顺序
	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
}

type MonitorTypeNamespace struct {
	// 监控类型
	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 策略类型值
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type NoticeBindPolicys struct {
	// 告警通知模板 ID
	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`

	// 告警通知模板绑定的告警策略ID列表
	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

type Operator struct {
	// 运算符标识
	Id *string `json:"Id,omitempty" name:"Id"`

	// 运算符展示名
	Name *string `json:"Name,omitempty" name:"Name"`
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

type PolicyGroup struct {
	// 是否可设为默认告警策略
	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`

	// 告警策略组ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 告警策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 是否为默认告警策略
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// 告警策略启用状态
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 最后修改人UIN
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 未屏蔽的实例数
	NoShieldedInstanceCount *int64 `json:"NoShieldedInstanceCount,omitempty" name:"NoShieldedInstanceCount"`

	// 父策略组ID
	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`

	// 所属项目ID
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// 告警接收对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverInfos []*PolicyGroupReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 修改时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 总绑定实例数
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 是否为与关系规则
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type PolicyGroupReceiverInfo struct {
	// 有效时段结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 是否需要发送通知
	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`

	// 告警接收渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`

	// 电话告警对个人间隔（秒）
	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`

	// 消息接收组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList"`

	// 接受者类型
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 接收人列表。通过平台接口查询到的接收人id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`

	// 告警恢复通知方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`

	// 电话告警每轮间隔（秒）
	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`

	// 电话告警轮数
	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`

	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendFor []*string `json:"SendFor,omitempty" name:"SendFor"`

	// 有效时段开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 电话告警接收者uid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UIDList []*int64 `json:"UIDList,omitempty" name:"UIDList"`
}

type PolicyTag struct {
	// 标签Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签Value
	Value *string `json:"Value,omitempty" name:"Value"`
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

type PrometheusAgent struct {
	// Agent 名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv4 *string `json:"Ipv4,omitempty" name:"Ipv4"`

	// 心跳时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeartbeatTime *string `json:"HeartbeatTime,omitempty" name:"HeartbeatTime"`

	// 最近一次错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastError *string `json:"LastError,omitempty" name:"LastError"`

	// Agent 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`

	// Agent 状态
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type PrometheusAgentInfo struct {
	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 备注
	Describe *string `json:"Describe,omitempty" name:"Describe"`
}

type PrometheusAgentOverview struct {
	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// agent状态
	// normal = 正常
	// abnormal = 异常
	Status *string `json:"Status,omitempty" name:"Status"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 额外labels
	// 本集群的所有指标都会带上这几个label
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalLabels []*Label `json:"ExternalLabels,omitempty" name:"ExternalLabels"`

	// 集群所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 集群所在VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 记录关联等操作的失败信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`

	// agent名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PrometheusAlertManagerConfig struct {
	// alertmanager url
	Url *string `json:"Url,omitempty" name:"Url"`

	// alertmanager部署所在集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// alertmanager部署所在集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type PrometheusAlertPolicyItem struct {
	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则列表
	Rules []*PrometheusAlertRule `json:"Rules,omitempty" name:"Rules"`

	// 告警策略 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 如果该告警来自模板下发，则TemplateId为模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 告警渠道，模板中使用可能返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notification *PrometheusNotificationItem `json:"Notification,omitempty" name:"Notification"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 如果告警策略来源于用户集群CRD资源定义，则ClusterId为所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type PrometheusAlertRule struct {
	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// prometheus语句
	Rule *string `json:"Rule,omitempty" name:"Rule"`

	// 额外标签
	Labels []*Label `json:"Labels,omitempty" name:"Labels"`

	// 告警发送模板
	Template *string `json:"Template,omitempty" name:"Template"`

	// 持续时间
	For *string `json:"For,omitempty" name:"For"`

	// 该条规则的描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// 参考prometheus rule中的annotations
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotations []*Label `json:"Annotations,omitempty" name:"Annotations"`

	// 告警规则状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`
}

type PrometheusClusterAgentBasic struct {
	// 集群ID
	Region *string `json:"Region,omitempty" name:"Region"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否开启公网CLB
	EnableExternal *bool `json:"EnableExternal,omitempty" name:"EnableExternal"`

	// 集群内部署组件的pod配置
	InClusterPodConfig *PrometheusClusterAgentPodConfig `json:"InClusterPodConfig,omitempty" name:"InClusterPodConfig"`

	// 该集群采集的所有指标都会带上这些labels
	ExternalLabels []*Label `json:"ExternalLabels,omitempty" name:"ExternalLabels"`

	// 是否安装默认采集配置
	NotInstallBasicScrape *bool `json:"NotInstallBasicScrape,omitempty" name:"NotInstallBasicScrape"`

	// 是否采集指标，true代表drop所有指标，false代表采集默认指标
	NotScrape *bool `json:"NotScrape,omitempty" name:"NotScrape"`
}

type PrometheusClusterAgentPodConfig struct {
	// 是否使用HostNetWork
	HostNet *bool `json:"HostNet,omitempty" name:"HostNet"`

	// 指定pod运行节点
	NodeSelector []*Label `json:"NodeSelector,omitempty" name:"NodeSelector"`

	// 容忍污点
	Tolerations []*Toleration `json:"Tolerations,omitempty" name:"Tolerations"`
}

type PrometheusConfigItem struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 配置内容
	Config *string `json:"Config,omitempty" name:"Config"`

	// 用于出参，如果该配置来至模板，则为模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type PrometheusInstanceGrantInfo struct {
	// 是否有计费操作权限(1=有，2=无)
	HasChargeOperation *int64 `json:"HasChargeOperation,omitempty" name:"HasChargeOperation"`

	// 是否显示VPC信息的权限(1=有，2=无)
	HasVpcDisplay *int64 `json:"HasVpcDisplay,omitempty" name:"HasVpcDisplay"`

	// 是否可修改Grafana的状态(1=有，2=无)
	HasGrafanaStatusChange *int64 `json:"HasGrafanaStatusChange,omitempty" name:"HasGrafanaStatusChange"`

	// 是否有管理agent的权限(1=有，2=无)
	HasAgentManage *int64 `json:"HasAgentManage,omitempty" name:"HasAgentManage"`

	// 是否有管理TKE集成的权限(1=有，2=无)
	HasTkeManage *int64 `json:"HasTkeManage,omitempty" name:"HasTkeManage"`

	// 是否显示API等信息(1=有, 2=无)
	HasApiOperation *int64 `json:"HasApiOperation,omitempty" name:"HasApiOperation"`
}

type PrometheusInstanceTenantUsage struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 计费周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	CalcDate *string `json:"CalcDate,omitempty" name:"CalcDate"`

	// 总用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *float64 `json:"Total,omitempty" name:"Total"`

	// 基础指标用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Basic *float64 `json:"Basic,omitempty" name:"Basic"`

	// 付费指标用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fee *float64 `json:"Fee,omitempty" name:"Fee"`
}

type PrometheusInstancesItem struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例计费模式。取值范围：
	// <ul>
	// <li>2：包年包月</li>
	// <li>3：按量</li>
	// </ul>
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 地域 ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 存储周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRetentionTime *int64 `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`

	// 实例业务状态。取值范围：
	// <ul>
	// <li>1：正在创建</li>
	// <li>2：运行中</li>
	// <li>3：异常</li>
	// <li>4：重建中</li>
	// <li>5：销毁中</li>
	// <li>6：已停服</li>
	// <li>8：欠费停服中</li>
	// <li>9：欠费已停服</li>
	// </ul>
	InstanceStatus *int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// Grafana 面板 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaURL *string `json:"GrafanaURL,omitempty" name:"GrafanaURL"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 是否开启 Grafana
	// <li>0：不开启</li>
	// <li>1：开启</li>
	EnableGrafana *int64 `json:"EnableGrafana,omitempty" name:"EnableGrafana"`

	// 实例IPV4地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv4Address *string `json:"IPv4Address,omitempty" name:"IPv4Address"`

	// 实例关联的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification []*PrometheusTag `json:"TagSpecification,omitempty" name:"TagSpecification"`

	// 购买的实例过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 计费状态
	// <ul>
	// <li>1：正常</li>
	// <li>2：过期</li>
	// <li>3：销毁</li>
	// <li>4：分配中</li>
	// <li>5：分配失败</li>
	// </ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *int64 `json:"ChargeStatus,omitempty" name:"ChargeStatus"`

	// 规格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 自动续费标记
	// <ul>
	// <li>0：不自动续费</li>
	// <li>1：开启自动续费</li>
	// <li>2：禁止自动续费</li>
	// <li>-1：无效</ii>
	// </ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 是否快过期
	// <ul>
	// <li>0：否</li>
	// <li>1：快过期</li>
	// </ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNearExpire *int64 `json:"IsNearExpire,omitempty" name:"IsNearExpire"`

	// 数据写入需要的 Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthToken *string `json:"AuthToken,omitempty" name:"AuthToken"`

	// Prometheus Remote Write 的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteWrite *string `json:"RemoteWrite,omitempty" name:"RemoteWrite"`

	// Prometheus HTTP Api 根地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiRootPath *string `json:"ApiRootPath,omitempty" name:"ApiRootPath"`

	// Proxy 的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyAddress *string `json:"ProxyAddress,omitempty" name:"ProxyAddress"`

	// Grafana 运行状态
	// <ul>
	// <li>1：正在创建</li>
	// <li>2：运行中</li>
	// <li>3：异常</li>
	// <li>4：重启中</li>
	// <li>5：销毁中</li>
	// <li>6：已停机</li>
	// <li>7：已删除</li>
	// </ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaStatus *int64 `json:"GrafanaStatus,omitempty" name:"GrafanaStatus"`

	// Grafana IP 白名单列表，使用英文分号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaIpWhiteList *string `json:"GrafanaIpWhiteList,omitempty" name:"GrafanaIpWhiteList"`

	// 实例的授权信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grant *PrometheusInstanceGrantInfo `json:"Grant,omitempty" name:"Grant"`

	// 绑定的 Grafana 实例 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" name:"GrafanaInstanceId"`

	// 告警规则限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertRuleLimit *int64 `json:"AlertRuleLimit,omitempty" name:"AlertRuleLimit"`

	// 预聚合规则限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordingRuleLimit *int64 `json:"RecordingRuleLimit,omitempty" name:"RecordingRuleLimit"`

	// 迁移状态，0-不在迁移中，1-迁移中、原实例，2-迁移中、目标实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrationType *int64 `json:"MigrationType,omitempty" name:"MigrationType"`
}

type PrometheusInstancesOverview struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 运行状态（1:正在创建；2:运行中；3:异常；4:重启中；5:销毁中； 6:已停机； 7: 已删除）
	InstanceStatus *int64 `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 计费状态（1:正常；2:过期; 3:销毁; 4:分配中; 5:分配失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *int64 `json:"ChargeStatus,omitempty" name:"ChargeStatus"`

	// 是否开启 Grafana（0:不开启，1:开启）
	EnableGrafana *int64 `json:"EnableGrafana,omitempty" name:"EnableGrafana"`

	// Grafana 面板 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaURL *string `json:"GrafanaURL,omitempty" name:"GrafanaURL"`

	// 实例付费类型（1:试用版；2:预付费）
	InstanceChargeType *int64 `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 规格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 存储周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRetentionTime *int64 `json:"DataRetentionTime,omitempty" name:"DataRetentionTime"`

	// 购买的实例过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 自动续费标记(0:不自动续费；1:开启自动续费；2:禁止自动续费；-1:无效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 绑定集群总数
	BoundTotal *int64 `json:"BoundTotal,omitempty" name:"BoundTotal"`

	// 绑定集群正常状态总数
	BoundNormal *int64 `json:"BoundNormal,omitempty" name:"BoundNormal"`
}

type PrometheusJobTargets struct {

}

type PrometheusNotificationItem struct {
	// 是否启用
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 通道类型，默认为amp，支持以下
	// amp
	// webhook
	// alertmanager
	Type *string `json:"Type,omitempty" name:"Type"`

	// 如果Type为webhook, 则该字段为必填项
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebHook *string `json:"WebHook,omitempty" name:"WebHook"`

	// 如果Type为alertmanager, 则该字段为必填项
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertManager *PrometheusAlertManagerConfig `json:"AlertManager,omitempty" name:"AlertManager"`

	// 收敛时间
	RepeatInterval *string `json:"RepeatInterval,omitempty" name:"RepeatInterval"`

	// 生效起始时间
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" name:"TimeRangeStart"`

	// 生效结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

	// 告警通知方式。目前有SMS、EMAIL、CALL、WECHAT方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`

	// 告警接收组（用户组）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverGroups []*string `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`

	// 电话告警顺序。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNotifyOrder []*uint64 `json:"PhoneNotifyOrder,omitempty" name:"PhoneNotifyOrder"`

	// 电话告警次数。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`

	// 电话告警轮内间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`

	// 电话告警轮外间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`

	// 电话告警触达通知
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneArriveNotice *bool `json:"PhoneArriveNotice,omitempty" name:"PhoneArriveNotice"`
}

type PrometheusRecordRuleYamlItem struct {
	// 实例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 最近更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Yaml内容
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 如果该聚合规则来至模板，则TemplateId为模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 该聚合规则如果来源于用户集群crd资源定义，则ClusterId为所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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

type PrometheusScrapeJob struct {
	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 任务 ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *string `json:"Config,omitempty" name:"Config"`
}

type PrometheusTag struct {
	// 标签的健值
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签对应的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PrometheusTemp struct {
	// 模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板维度，支持以下类型
	// instance 实例级别
	// cluster 集群级别
	Level *string `json:"Level,omitempty" name:"Level"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// 当Level为instance时有效，
	// 模板中的聚合规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitempty" name:"RecordRules"`

	// 当Level为cluster时有效，
	// 模板中的ServiceMonitor规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// 当Level为cluster时有效，
	// 模板中的PodMonitors规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// 当Level为cluster时有效，
	// 模板中的RawJobs规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`

	// 模板的ID, 用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 最近更新时间，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 当前版本，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 是否系统提供的默认模板，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 当Level为instance时有效，
	// 模板中的告警配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertDetailRules []*PrometheusAlertPolicyItem `json:"AlertDetailRules,omitempty" name:"AlertDetailRules"`

	// 关联实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetsTotal *int64 `json:"TargetsTotal,omitempty" name:"TargetsTotal"`
}

type PrometheusTempModify struct {
	// 修改名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 修改描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// 当Level为cluster时有效，
	// 模板中的ServiceMonitor规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitempty" name:"ServiceMonitors"`

	// 当Level为cluster时有效，
	// 模板中的PodMonitors规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitempty" name:"PodMonitors"`

	// 当Level为cluster时有效，
	// 模板中的RawJobs规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitempty" name:"RawJobs"`

	// 当Level为instance时有效，
	// 模板中的聚合规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitempty" name:"RecordRules"`

	// 修改内容，只有当模板类型是Alert时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertDetailRules []*PrometheusAlertPolicyItem `json:"AlertDetailRules,omitempty" name:"AlertDetailRules"`
}

type PrometheusTemplateSyncTarget struct {
	// 目标所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 目标实例
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群id，只有当采集模板的Level为cluster的时候需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最后一次同步时间， 用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	SyncTime *string `json:"SyncTime,omitempty" name:"SyncTime"`

	// 当前使用的模板版本，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 集群类型，只有当采集模板的Level为cluster的时候需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 用于出参，实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 用于出参，集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type PrometheusZoneItem struct {
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区 ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区状态( 0: 不可用；1: 可用)
	ZoneState *int64 `json:"ZoneState,omitempty" name:"ZoneState"`

	// 地域 ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区名（目前为中文）
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

// Predefined struct for user
type PutMonitorDataRequestParams struct {
	// 一组指标和数据
	Metrics []*MetricDatum `json:"Metrics,omitempty" name:"Metrics"`

	// 上报时自行指定的 IP
	AnnounceIp *string `json:"AnnounceIp,omitempty" name:"AnnounceIp"`

	// 上报时自行指定的时间戳
	AnnounceTimestamp *uint64 `json:"AnnounceTimestamp,omitempty" name:"AnnounceTimestamp"`

	// 上报时自行指定的 IP 或 产品实例ID
	AnnounceInstance *string `json:"AnnounceInstance,omitempty" name:"AnnounceInstance"`
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

// Predefined struct for user
type PutMonitorDataResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PutMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *PutMonitorDataResponseParams `json:"Response"`
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
	// 告警时间段开始时间。范围[0,86400)，作为 UNIX 时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"
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

	// 电话告警接收者 UID
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

	// 接收组列表。通过平台接口查询到的接收组 ID 列表
	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList"`

	// 接收人列表。通过平台接口查询到的接收人 ID 列表
	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`

	// 告警接收语言，枚举值（zh-CN，en-US）
	ReceiveLanguage *string `json:"ReceiveLanguage,omitempty" name:"ReceiveLanguage"`
}

type RecordingRuleSet struct {
	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则状态码
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`

	// 分组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 规则内容组
	Group *string `json:"Group,omitempty" name:"Group"`

	// 规则数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 规则创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 规则最近更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitempty" name:"RuleName"`
}

// Predefined struct for user
type ResumeGrafanaInstanceRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ResumeGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ResumeGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeGrafanaInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResumeGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeGrafanaInstanceResponseParams `json:"Response"`
}

func (r *ResumeGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeGrafanaInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunPrometheusInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网ID，默认使用实例所用子网初始化，也可通过该参数传递新的子网ID初始化
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type RunPrometheusInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子网ID，默认使用实例所用子网初始化，也可通过该参数传递新的子网ID初始化
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *RunPrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunPrometheusInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunPrometheusInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunPrometheusInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunPrometheusInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RunPrometheusInstanceResponseParams `json:"Response"`
}

func (r *RunPrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunPrometheusInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendCustomAlarmMsgRequestParams struct {
	// 接口模块名，当前取值monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 消息策略ID，在云监控自定义消息页面配置
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用户想要发送的自定义消息内容
	Msg *string `json:"Msg,omitempty" name:"Msg"`
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

// Predefined struct for user
type SendCustomAlarmMsgResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendCustomAlarmMsgResponse struct {
	*tchttp.BaseResponse
	Response *SendCustomAlarmMsgResponseParams `json:"Response"`
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

// Predefined struct for user
type SetDefaultAlarmPolicyRequestParams struct {
	// 模块名，固定值 monitor
	Module *string `json:"Module,omitempty" name:"Module"`

	// 告警策略 ID
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
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

// Predefined struct for user
type SetDefaultAlarmPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetDefaultAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *SetDefaultAlarmPolicyResponseParams `json:"Response"`
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

// Predefined struct for user
type SyncPrometheusTempRequestParams struct {
	// 实例id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 同步目标
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitempty" name:"Targets"`
}

type SyncPrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 同步目标
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitempty" name:"Targets"`
}

func (r *SyncPrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncPrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPrometheusTempResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SyncPrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *SyncPrometheusTempResponseParams `json:"Response"`
}

func (r *SyncPrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签value
	Value *string `json:"Value,omitempty" name:"Value"`
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

type TaskStepInfo struct {
	// 步骤名称
	Step *string `json:"Step,omitempty" name:"Step"`

	// 生命周期
	// pending : 步骤未开始
	// running: 步骤执行中
	// success: 步骤成功完成
	// failed: 步骤失败
	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`

	// 步骤开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// 步骤结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// 若步骤生命周期为failed,则此字段显示错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedMsg *string `json:"FailedMsg,omitempty" name:"FailedMsg"`
}

type TemplateGroup struct {
	// 指标告警规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions"`

	// 事件告警规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventConditions []*EventCondition `json:"EventConditions,omitempty" name:"EventConditions"`

	// 关联告警策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyGroups []*PolicyGroup `json:"PolicyGroups,omitempty" name:"PolicyGroups"`

	// 模板策略组ID
	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`

	// 模板策略组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 最后修改人UIN
	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 视图
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 是否为与关系
	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

// Predefined struct for user
type TerminatePrometheusInstancesRequestParams struct {
	// 实例 ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type TerminatePrometheusInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *TerminatePrometheusInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminatePrometheusInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminatePrometheusInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminatePrometheusInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminatePrometheusInstancesResponse struct {
	*tchttp.BaseResponse
	Response *TerminatePrometheusInstancesResponseParams `json:"Response"`
}

func (r *TerminatePrometheusInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminatePrometheusInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Toleration struct {
	// 容忍应用到的 taint key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 键与值的关系
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 要匹配的污点效果
	Effect *string `json:"Effect,omitempty" name:"Effect"`
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

	// 通知开始时间 一天开始的秒数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 通知结束时间 一天开始的秒数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 通知周期 1-7表示周一到周日
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weekday []*int64 `json:"Weekday,omitempty" name:"Weekday"`
}

// Predefined struct for user
type UnBindingAllPolicyObjectRequestParams struct {
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，如传入 PolicyId 则该字段被忽略可传入任意值如 0
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警策略ID，使用此字段时 GroupId 会被忽略
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`

	// 是否配置了事件告警
	EbEventFlag *int64 `json:"EbEventFlag,omitempty" name:"EbEventFlag"`
}

type UnBindingAllPolicyObjectRequest struct {
	*tchttp.BaseRequest
	
	// 固定值，为"monitor"
	Module *string `json:"Module,omitempty" name:"Module"`

	// 策略组id，如传入 PolicyId 则该字段被忽略可传入任意值如 0
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 告警策略ID，使用此字段时 GroupId 会被忽略
	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`

	// 是否配置了事件告警
	EbEventFlag *int64 `json:"EbEventFlag,omitempty" name:"EbEventFlag"`
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
	delete(f, "EbSubject")
	delete(f, "EbEventFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindingAllPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindingAllPolicyObjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindingAllPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *UnBindingAllPolicyObjectResponseParams `json:"Response"`
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

// Predefined struct for user
type UnBindingPolicyObjectRequestParams struct {
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

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`

	// 是否配置了事件告警
	EbEventFlag *int64 `json:"EbEventFlag,omitempty" name:"EbEventFlag"`
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

	// 事件配置的告警
	EbSubject *string `json:"EbSubject,omitempty" name:"EbSubject"`

	// 是否配置了事件告警
	EbEventFlag *int64 `json:"EbEventFlag,omitempty" name:"EbEventFlag"`
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
	delete(f, "EbSubject")
	delete(f, "EbEventFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindingPolicyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnBindingPolicyObjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindingPolicyObjectResponse struct {
	*tchttp.BaseResponse
	Response *UnBindingPolicyObjectResponseParams `json:"Response"`
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

// Predefined struct for user
type UnbindPrometheusManagedGrafanaRequestParams struct {
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Grafana 实例 ID
	GrafanaId *string `json:"GrafanaId,omitempty" name:"GrafanaId"`
}

type UnbindPrometheusManagedGrafanaRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Grafana 实例 ID
	GrafanaId *string `json:"GrafanaId,omitempty" name:"GrafanaId"`
}

func (r *UnbindPrometheusManagedGrafanaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindPrometheusManagedGrafanaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GrafanaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindPrometheusManagedGrafanaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindPrometheusManagedGrafanaResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindPrometheusManagedGrafanaResponse struct {
	*tchttp.BaseResponse
	Response *UnbindPrometheusManagedGrafanaResponseParams `json:"Response"`
}

func (r *UnbindPrometheusManagedGrafanaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindPrometheusManagedGrafanaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaDashboardRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus 集成项 Code，删除对应的 Dashboard，Code 如下：
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitempty" name:"IntegrationCodes"`
}

type UninstallGrafanaDashboardRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus 集成项 Code，删除对应的 Dashboard，Code 如下：
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitempty" name:"IntegrationCodes"`
}

func (r *UninstallGrafanaDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallGrafanaDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaDashboardResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UninstallGrafanaDashboardResponse struct {
	*tchttp.BaseResponse
	Response *UninstallGrafanaDashboardResponseParams `json:"Response"`
}

func (r *UninstallGrafanaDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaPluginsRequestParams struct {
	// 插件 ID 数组，例如"PluginIds": [ "grafana-clock-panel" ]，可通过 DescribePluginOverviews 获取 PluginId
	PluginIds []*string `json:"PluginIds,omitempty" name:"PluginIds"`

	// Grafana 实例 ID，例如：grafana-abcdefg
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type UninstallGrafanaPluginsRequest struct {
	*tchttp.BaseRequest
	
	// 插件 ID 数组，例如"PluginIds": [ "grafana-clock-panel" ]，可通过 DescribePluginOverviews 获取 PluginId
	PluginIds []*string `json:"PluginIds,omitempty" name:"PluginIds"`

	// Grafana 实例 ID，例如：grafana-abcdefg
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *UninstallGrafanaPluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaPluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallGrafanaPluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallGrafanaPluginsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UninstallGrafanaPluginsResponse struct {
	*tchttp.BaseResponse
	Response *UninstallGrafanaPluginsResponseParams `json:"Response"`
}

func (r *UninstallGrafanaPluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallGrafanaPluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlertRuleRequestParams struct {
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

// Predefined struct for user
type UpdateAlertRuleResponseParams struct {
	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAlertRuleResponseParams `json:"Response"`
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

// Predefined struct for user
type UpdateAlertRuleStateRequestParams struct {
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

// Predefined struct for user
type UpdateAlertRuleStateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAlertRuleStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAlertRuleStateResponseParams `json:"Response"`
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

// Predefined struct for user
type UpdateDNSConfigRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DNS 数组
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`
}

type UpdateDNSConfigRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// DNS 数组
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`
}

func (r *UpdateDNSConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDNSConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NameServers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDNSConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDNSConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDNSConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDNSConfigResponseParams `json:"Response"`
}

func (r *UpdateDNSConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDNSConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateExporterIntegrationRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 配置内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type UpdateExporterIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 类型
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 配置内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// Kubernetes 集群类型，取值如下：
	// <li> 1= 容器集群(TKE) </li>
	// <li> 2=弹性集群<EKS> </li>
	// <li> 3= Prometheus管理的弹性集群<MEKS> </li>
	KubeType *int64 `json:"KubeType,omitempty" name:"KubeType"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *UpdateExporterIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateExporterIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Content")
	delete(f, "KubeType")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateExporterIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateExporterIntegrationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateExporterIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateExporterIntegrationResponseParams `json:"Response"`
}

func (r *UpdateExporterIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateExporterIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaConfigRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// JSON 编码后的字符串
	Config *string `json:"Config,omitempty" name:"Config"`
}

type UpdateGrafanaConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// JSON 编码后的字符串
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *UpdateGrafanaConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGrafanaConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaConfigResponseParams `json:"Response"`
}

func (r *UpdateGrafanaConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaEnvironmentsRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 环境变量字符串
	Envs *string `json:"Envs,omitempty" name:"Envs"`
}

type UpdateGrafanaEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 环境变量字符串
	Envs *string `json:"Envs,omitempty" name:"Envs"`
}

func (r *UpdateGrafanaEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Envs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaEnvironmentsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGrafanaEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaEnvironmentsResponseParams `json:"Response"`
}

func (r *UpdateGrafanaEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaIntegrationRequestParams struct {
	// 集成 ID，可在实例详情-云产品集成-集成列表查看。例如：integration-abcd1234
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`

	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成类型，可在实例详情-云产品集成-集成列表查看。例如：tencent-cloud-prometheus
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集成内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type UpdateGrafanaIntegrationRequest struct {
	*tchttp.BaseRequest
	
	// 集成 ID，可在实例详情-云产品集成-集成列表查看。例如：integration-abcd1234
	IntegrationId *string `json:"IntegrationId,omitempty" name:"IntegrationId"`

	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集成类型，可在实例详情-云产品集成-集成列表查看。例如：tencent-cloud-prometheus
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集成内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *UpdateGrafanaIntegrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaIntegrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntegrationId")
	delete(f, "InstanceId")
	delete(f, "Kind")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaIntegrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaIntegrationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGrafanaIntegrationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaIntegrationResponseParams `json:"Response"`
}

func (r *UpdateGrafanaIntegrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaIntegrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaNotificationChannelRequestParams struct {
	// 通道 ID，例如：nchannel-abcd1234
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 接受告警通道 ID 数组
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 已废弃，请使用 OrganizationIds
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitempty" name:"ExtraOrgIds"`

	// 生效的组织 ID 数组
	OrganizationIds []*string `json:"OrganizationIds,omitempty" name:"OrganizationIds"`
}

type UpdateGrafanaNotificationChannelRequest struct {
	*tchttp.BaseRequest
	
	// 通道 ID，例如：nchannel-abcd1234
	ChannelId *string `json:"ChannelId,omitempty" name:"ChannelId"`

	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 告警通道名称，例如：test
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 接受告警通道 ID 数组
	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`

	// 已废弃，请使用 OrganizationIds
	ExtraOrgIds []*string `json:"ExtraOrgIds,omitempty" name:"ExtraOrgIds"`

	// 生效的组织 ID 数组
	OrganizationIds []*string `json:"OrganizationIds,omitempty" name:"OrganizationIds"`
}

func (r *UpdateGrafanaNotificationChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaNotificationChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "InstanceId")
	delete(f, "ChannelName")
	delete(f, "Receivers")
	delete(f, "ExtraOrgIds")
	delete(f, "OrganizationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaNotificationChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaNotificationChannelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGrafanaNotificationChannelResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaNotificationChannelResponseParams `json:"Response"`
}

func (r *UpdateGrafanaNotificationChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaNotificationChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaWhiteListRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 白名单数组，输入公网域名 IP ，例如：127.0.0.1，可通过接口 DescribeGrafanaWhiteList 查看
	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist"`
}

type UpdateGrafanaWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 白名单数组，输入公网域名 IP ，例如：127.0.0.1，可通过接口 DescribeGrafanaWhiteList 查看
	Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist"`
}

func (r *UpdateGrafanaWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Whitelist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGrafanaWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGrafanaWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateGrafanaWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGrafanaWhiteListResponseParams `json:"Response"`
}

func (r *UpdateGrafanaWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGrafanaWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAgentStatusRequestParams struct {
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID 列表，例如：agent-abcd1234，可在控制台 Agent 管理中获取
	AgentIds []*string `json:"AgentIds,omitempty" name:"AgentIds"`

	// 要更新的状态
	// <li> 1= 开启 </li>
	// <li> 2= 关闭 </li>
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type UpdatePrometheusAgentStatusRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID 列表，例如：agent-abcd1234，可在控制台 Agent 管理中获取
	AgentIds []*string `json:"AgentIds,omitempty" name:"AgentIds"`

	// 要更新的状态
	// <li> 1= 开启 </li>
	// <li> 2= 关闭 </li>
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *UpdatePrometheusAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAgentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentIds")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrometheusAgentStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusAgentStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdatePrometheusAgentStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePrometheusAgentStatusResponseParams `json:"Response"`
}

func (r *UpdatePrometheusAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusScrapeJobRequestParams struct {
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID，例如：agent-abcd1234，可在控制台 Agent 管理中获取
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 抓取任务 ID，例如：job-abcd1234，可在控制台 Agent 管理-抓取任务配置中获取
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 抓取任务配置，格式：job_name:xx
	Config *string `json:"Config,omitempty" name:"Config"`
}

type UpdatePrometheusScrapeJobRequest struct {
	*tchttp.BaseRequest
	
	// Prometheus 实例 ID，例如：prom-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Agent ID，例如：agent-abcd1234，可在控制台 Agent 管理中获取
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 抓取任务 ID，例如：job-abcd1234，可在控制台 Agent 管理-抓取任务配置中获取
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 抓取任务配置，格式：job_name:xx
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *UpdatePrometheusScrapeJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusScrapeJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentId")
	delete(f, "JobId")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePrometheusScrapeJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePrometheusScrapeJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdatePrometheusScrapeJobResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePrometheusScrapeJobResponseParams `json:"Response"`
}

func (r *UpdatePrometheusScrapeJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePrometheusScrapeJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordingRuleRequestParams struct {
	// 聚合规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 聚合规则组内容，格式为 yaml，通过 base64 进行编码。
	Group *string `json:"Group,omitempty" name:"Group"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus 聚合规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则状态码，取值如下：
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// 默认状态码为 2 启用。
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`
}

type UpdateRecordingRuleRequest struct {
	*tchttp.BaseRequest
	
	// 聚合规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 聚合规则组内容，格式为 yaml，通过 base64 进行编码。
	Group *string `json:"Group,omitempty" name:"Group"`

	// Prometheus 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus 聚合规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 规则状态码，取值如下：
	// <li>1=RuleDeleted</li>
	// <li>2=RuleEnabled</li>
	// <li>3=RuleDisabled</li>
	// 默认状态码为 2 启用。
	RuleState *int64 `json:"RuleState,omitempty" name:"RuleState"`
}

func (r *UpdateRecordingRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordingRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Group")
	delete(f, "InstanceId")
	delete(f, "RuleId")
	delete(f, "RuleState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordingRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordingRuleResponseParams struct {
	// 规则 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateRecordingRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordingRuleResponseParams `json:"Response"`
}

func (r *UpdateRecordingRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordingRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSSOAccountRequestParams struct {
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户账号 ID ，例如：10000000
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 权限
	Role []*GrafanaAccountRole `json:"Role,omitempty" name:"Role"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

type UpdateSSOAccountRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-abcdefgh
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户账号 ID ，例如：10000000
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 权限
	Role []*GrafanaAccountRole `json:"Role,omitempty" name:"Role"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

func (r *UpdateSSOAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSSOAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserId")
	delete(f, "Role")
	delete(f, "Notes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSSOAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSSOAccountResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateSSOAccountResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSSOAccountResponseParams `json:"Response"`
}

func (r *UpdateSSOAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSSOAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServiceDiscoveryRequestParams struct {
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

// Predefined struct for user
type UpdateServiceDiscoveryResponseParams struct {
	// 更新成功之后，返回对应服务发现的信息
	ServiceDiscovery *ServiceDiscoveryItem `json:"ServiceDiscovery,omitempty" name:"ServiceDiscovery"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateServiceDiscoveryResponse struct {
	*tchttp.BaseResponse
	Response *UpdateServiceDiscoveryResponseParams `json:"Response"`
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

// Predefined struct for user
type UpgradeGrafanaDashboardRequestParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus 集成项 Code，升级对应的 Dashboard，取值如下：
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitempty" name:"IntegrationCodes"`
}

type UpgradeGrafanaDashboardRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// Prometheus 集成项 Code，升级对应的 Dashboard，取值如下：
	// <li>spring_mvc</li>
	// <li>mysql</li>
	// <li>go</li>
	// <li>redis</li>
	// <li>jvm</li>
	// <li>pgsql</li>
	// <li>mongo</li>
	// <li>kafka</li>
	// <li>es</li>
	// <li>flink</li>
	// <li>blackbox</li>
	// <li>consule</li>
	// <li>memcached</li>
	// <li>zk</li>
	// <li>tps</li>
	// <li>istio</li>
	// <li>etcd</li>
	IntegrationCodes []*string `json:"IntegrationCodes,omitempty" name:"IntegrationCodes"`
}

func (r *UpgradeGrafanaDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IntegrationCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeGrafanaDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGrafanaDashboardResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeGrafanaDashboardResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeGrafanaDashboardResponseParams `json:"Response"`
}

func (r *UpgradeGrafanaDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGrafanaInstanceRequestParams struct {
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 版本别名，例如：v7.4.2
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type UpgradeGrafanaInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Grafana 实例 ID，例如：grafana-12345678
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 版本别名，例如：v7.4.2
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *UpgradeGrafanaInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeGrafanaInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeGrafanaInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeGrafanaInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeGrafanaInstanceResponseParams `json:"Response"`
}

func (r *UpgradeGrafanaInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeGrafanaInstanceResponse) FromJsonString(s string) error {
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

	// 通知渠道列表 EMAIL=邮件 SMS=短信 CALL=电话 WECHAT=微信 RTX=企业微信
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

	// 电话拨打类型 SYNC=同时拨打 CIRCLE=轮询拨打 不指定时默认是轮询
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCallType *string `json:"PhoneCallType,omitempty" name:"PhoneCallType"`

	// 通知周期 1-7表示周一到周日
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weekday []*int64 `json:"Weekday,omitempty" name:"Weekday"`
}